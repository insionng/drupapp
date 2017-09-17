package handler

import (
	"github.com/insionng/drupapp/helper"
	"github.com/insionng/drupapp/model"
	"github.com/insionng/macross"
)

func GetUsersFieldDatasCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()

	if offset > 0 {
		var m = map[string]interface{}{}
		_int64, _error := model.GetUsersFieldDatasCount(offset, limit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		m["users_field_datasCount"] = _int64
		return self.JSON(m)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDataCountViaUidHandler(self *macross.Context) error {
	Uid_ := self.Args("uid").MustInt64()
	_int64 := model.GetUsersFieldDataCountViaUid(Uid_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["users_field_dataCount"] = 0
	}
	m["users_field_dataCount"] = _int64
	return self.JSON(m)
}

func GetUsersFieldDataCountViaLangcodeHandler(self *macross.Context) error {
	Langcode_ := self.Args("langcode").String()
	_int64 := model.GetUsersFieldDataCountViaLangcode(Langcode_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["users_field_dataCount"] = 0
	}
	m["users_field_dataCount"] = _int64
	return self.JSON(m)
}

func GetUsersFieldDataCountViaPreferredLangcodeHandler(self *macross.Context) error {
	PreferredLangcode_ := self.Args("preferred_langcode").String()
	_int64 := model.GetUsersFieldDataCountViaPreferredLangcode(PreferredLangcode_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["users_field_dataCount"] = 0
	}
	m["users_field_dataCount"] = _int64
	return self.JSON(m)
}

func GetUsersFieldDataCountViaPreferredAdminLangcodeHandler(self *macross.Context) error {
	PreferredAdminLangcode_ := self.Args("preferred_admin_langcode").String()
	_int64 := model.GetUsersFieldDataCountViaPreferredAdminLangcode(PreferredAdminLangcode_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["users_field_dataCount"] = 0
	}
	m["users_field_dataCount"] = _int64
	return self.JSON(m)
}

func GetUsersFieldDataCountViaNameHandler(self *macross.Context) error {
	Name_ := self.Args("name").String()
	_int64 := model.GetUsersFieldDataCountViaName(Name_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["users_field_dataCount"] = 0
	}
	m["users_field_dataCount"] = _int64
	return self.JSON(m)
}

func GetUsersFieldDataCountViaPassHandler(self *macross.Context) error {
	Pass_ := self.Args("pass").String()
	_int64 := model.GetUsersFieldDataCountViaPass(Pass_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["users_field_dataCount"] = 0
	}
	m["users_field_dataCount"] = _int64
	return self.JSON(m)
}

func GetUsersFieldDataCountViaMailHandler(self *macross.Context) error {
	Mail_ := self.Args("mail").String()
	_int64 := model.GetUsersFieldDataCountViaMail(Mail_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["users_field_dataCount"] = 0
	}
	m["users_field_dataCount"] = _int64
	return self.JSON(m)
}

func GetUsersFieldDataCountViaTimezoneHandler(self *macross.Context) error {
	Timezone_ := self.Args("timezone").String()
	_int64 := model.GetUsersFieldDataCountViaTimezone(Timezone_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["users_field_dataCount"] = 0
	}
	m["users_field_dataCount"] = _int64
	return self.JSON(m)
}

func GetUsersFieldDataCountViaStatusHandler(self *macross.Context) error {
	Status_ := self.Args("status").MustInt()
	_int64 := model.GetUsersFieldDataCountViaStatus(Status_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["users_field_dataCount"] = 0
	}
	m["users_field_dataCount"] = _int64
	return self.JSON(m)
}

func GetUsersFieldDataCountViaCreatedHandler(self *macross.Context) error {
	Created_ := self.Args("created").MustInt()
	_int64 := model.GetUsersFieldDataCountViaCreated(Created_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["users_field_dataCount"] = 0
	}
	m["users_field_dataCount"] = _int64
	return self.JSON(m)
}

func GetUsersFieldDataCountViaChangedHandler(self *macross.Context) error {
	Changed_ := self.Args("changed").MustInt()
	_int64 := model.GetUsersFieldDataCountViaChanged(Changed_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["users_field_dataCount"] = 0
	}
	m["users_field_dataCount"] = _int64
	return self.JSON(m)
}

func GetUsersFieldDataCountViaAccessHandler(self *macross.Context) error {
	Access_ := self.Args("access").MustInt()
	_int64 := model.GetUsersFieldDataCountViaAccess(Access_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["users_field_dataCount"] = 0
	}
	m["users_field_dataCount"] = _int64
	return self.JSON(m)
}

func GetUsersFieldDataCountViaLoginHandler(self *macross.Context) error {
	Login_ := self.Args("login").MustInt()
	_int64 := model.GetUsersFieldDataCountViaLogin(Login_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["users_field_dataCount"] = 0
	}
	m["users_field_dataCount"] = _int64
	return self.JSON(m)
}

func GetUsersFieldDataCountViaInitHandler(self *macross.Context) error {
	Init_ := self.Args("init").String()
	_int64 := model.GetUsersFieldDataCountViaInit(Init_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["users_field_dataCount"] = 0
	}
	m["users_field_dataCount"] = _int64
	return self.JSON(m)
}

func GetUsersFieldDataCountViaDefaultLangcodeHandler(self *macross.Context) error {
	DefaultLangcode_ := self.Args("default_langcode").MustInt()
	_int64 := model.GetUsersFieldDataCountViaDefaultLangcode(DefaultLangcode_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["users_field_dataCount"] = 0
	}
	m["users_field_dataCount"] = _int64
	return self.JSON(m)
}

func GetUsersFieldDatasViaUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iUid := self.Args("uid").MustInt64()
	if (offset > 0) && helper.IsHas(iUid) {
		_UsersFieldData, _error := model.GetUsersFieldDatasViaUid(offset, limit, iUid, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasViaUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iLangcode := self.Args("langcode").String()
	if (offset > 0) && helper.IsHas(iLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasViaLangcode(offset, limit, iLangcode, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasViaPreferredLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iPreferredLangcode := self.Args("preferred_langcode").String()
	if (offset > 0) && helper.IsHas(iPreferredLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasViaPreferredLangcode(offset, limit, iPreferredLangcode, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasViaPreferredLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasViaPreferredAdminLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iPreferredAdminLangcode := self.Args("preferred_admin_langcode").String()
	if (offset > 0) && helper.IsHas(iPreferredAdminLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasViaPreferredAdminLangcode(offset, limit, iPreferredAdminLangcode, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasViaPreferredAdminLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasViaNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iName := self.Args("name").String()
	if (offset > 0) && helper.IsHas(iName) {
		_UsersFieldData, _error := model.GetUsersFieldDatasViaName(offset, limit, iName, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasViaName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasViaPassHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iPass := self.Args("pass").String()
	if (offset > 0) && helper.IsHas(iPass) {
		_UsersFieldData, _error := model.GetUsersFieldDatasViaPass(offset, limit, iPass, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasViaPass's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasViaMailHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iMail := self.Args("mail").String()
	if (offset > 0) && helper.IsHas(iMail) {
		_UsersFieldData, _error := model.GetUsersFieldDatasViaMail(offset, limit, iMail, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasViaMail's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasViaTimezoneHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iTimezone := self.Args("timezone").String()
	if (offset > 0) && helper.IsHas(iTimezone) {
		_UsersFieldData, _error := model.GetUsersFieldDatasViaTimezone(offset, limit, iTimezone, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasViaTimezone's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasViaStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iStatus := self.Args("status").MustInt()
	if (offset > 0) && helper.IsHas(iStatus) {
		_UsersFieldData, _error := model.GetUsersFieldDatasViaStatus(offset, limit, iStatus, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasViaStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasViaCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iCreated := self.Args("created").MustInt()
	if (offset > 0) && helper.IsHas(iCreated) {
		_UsersFieldData, _error := model.GetUsersFieldDatasViaCreated(offset, limit, iCreated, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasViaCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasViaChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iChanged := self.Args("changed").MustInt()
	if (offset > 0) && helper.IsHas(iChanged) {
		_UsersFieldData, _error := model.GetUsersFieldDatasViaChanged(offset, limit, iChanged, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasViaChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasViaAccessHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iAccess := self.Args("access").MustInt()
	if (offset > 0) && helper.IsHas(iAccess) {
		_UsersFieldData, _error := model.GetUsersFieldDatasViaAccess(offset, limit, iAccess, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasViaAccess's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasViaLoginHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iLogin := self.Args("login").MustInt()
	if (offset > 0) && helper.IsHas(iLogin) {
		_UsersFieldData, _error := model.GetUsersFieldDatasViaLogin(offset, limit, iLogin, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasViaLogin's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasViaInitHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iInit := self.Args("init").String()
	if (offset > 0) && helper.IsHas(iInit) {
		_UsersFieldData, _error := model.GetUsersFieldDatasViaInit(offset, limit, iInit, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasViaInit's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasViaDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()
	if (offset > 0) && helper.IsHas(iDefaultLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasViaDefaultLangcode(offset, limit, iDefaultLangcode, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasViaDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByUidAndLangcodeAndPreferredLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt64()
	iLangcode := self.Args("langcode").String()
	iPreferredLangcode := self.Args("preferred_langcode").String()

	if helper.IsHas(iUid) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByUidAndLangcodeAndPreferredLangcode(offset, limit, iUid,iLangcode,iPreferredLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByUidAndLangcodeAndPreferredLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByUidAndLangcodeAndPreferredAdminLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt64()
	iLangcode := self.Args("langcode").String()
	iPreferredAdminLangcode := self.Args("preferred_admin_langcode").String()

	if helper.IsHas(iUid) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByUidAndLangcodeAndPreferredAdminLangcode(offset, limit, iUid,iLangcode,iPreferredAdminLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByUidAndLangcodeAndPreferredAdminLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByUidAndLangcodeAndNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt64()
	iLangcode := self.Args("langcode").String()
	iName := self.Args("name").String()

	if helper.IsHas(iUid) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByUidAndLangcodeAndName(offset, limit, iUid,iLangcode,iName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByUidAndLangcodeAndName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByUidAndLangcodeAndPassHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt64()
	iLangcode := self.Args("langcode").String()
	iPass := self.Args("pass").String()

	if helper.IsHas(iUid) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByUidAndLangcodeAndPass(offset, limit, iUid,iLangcode,iPass)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByUidAndLangcodeAndPass's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByUidAndLangcodeAndMailHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt64()
	iLangcode := self.Args("langcode").String()
	iMail := self.Args("mail").String()

	if helper.IsHas(iUid) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByUidAndLangcodeAndMail(offset, limit, iUid,iLangcode,iMail)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByUidAndLangcodeAndMail's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByUidAndLangcodeAndTimezoneHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt64()
	iLangcode := self.Args("langcode").String()
	iTimezone := self.Args("timezone").String()

	if helper.IsHas(iUid) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByUidAndLangcodeAndTimezone(offset, limit, iUid,iLangcode,iTimezone)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByUidAndLangcodeAndTimezone's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByUidAndLangcodeAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt64()
	iLangcode := self.Args("langcode").String()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iUid) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByUidAndLangcodeAndStatus(offset, limit, iUid,iLangcode,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByUidAndLangcodeAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByUidAndLangcodeAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt64()
	iLangcode := self.Args("langcode").String()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iUid) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByUidAndLangcodeAndCreated(offset, limit, iUid,iLangcode,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByUidAndLangcodeAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByUidAndLangcodeAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt64()
	iLangcode := self.Args("langcode").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iUid) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByUidAndLangcodeAndChanged(offset, limit, iUid,iLangcode,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByUidAndLangcodeAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByUidAndLangcodeAndAccessHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt64()
	iLangcode := self.Args("langcode").String()
	iAccess := self.Args("access").MustInt()

	if helper.IsHas(iUid) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByUidAndLangcodeAndAccess(offset, limit, iUid,iLangcode,iAccess)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByUidAndLangcodeAndAccess's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByUidAndLangcodeAndLoginHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt64()
	iLangcode := self.Args("langcode").String()
	iLogin := self.Args("login").MustInt()

	if helper.IsHas(iUid) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByUidAndLangcodeAndLogin(offset, limit, iUid,iLangcode,iLogin)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByUidAndLangcodeAndLogin's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByUidAndLangcodeAndInitHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt64()
	iLangcode := self.Args("langcode").String()
	iInit := self.Args("init").String()

	if helper.IsHas(iUid) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByUidAndLangcodeAndInit(offset, limit, iUid,iLangcode,iInit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByUidAndLangcodeAndInit's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByUidAndLangcodeAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt64()
	iLangcode := self.Args("langcode").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iUid) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByUidAndLangcodeAndDefaultLangcode(offset, limit, iUid,iLangcode,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByUidAndLangcodeAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByUidAndPreferredLangcodeAndPreferredAdminLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt64()
	iPreferredLangcode := self.Args("preferred_langcode").String()
	iPreferredAdminLangcode := self.Args("preferred_admin_langcode").String()

	if helper.IsHas(iUid) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByUidAndPreferredLangcodeAndPreferredAdminLangcode(offset, limit, iUid,iPreferredLangcode,iPreferredAdminLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByUidAndPreferredLangcodeAndPreferredAdminLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByUidAndPreferredLangcodeAndNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt64()
	iPreferredLangcode := self.Args("preferred_langcode").String()
	iName := self.Args("name").String()

	if helper.IsHas(iUid) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByUidAndPreferredLangcodeAndName(offset, limit, iUid,iPreferredLangcode,iName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByUidAndPreferredLangcodeAndName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByUidAndPreferredLangcodeAndPassHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt64()
	iPreferredLangcode := self.Args("preferred_langcode").String()
	iPass := self.Args("pass").String()

	if helper.IsHas(iUid) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByUidAndPreferredLangcodeAndPass(offset, limit, iUid,iPreferredLangcode,iPass)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByUidAndPreferredLangcodeAndPass's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByUidAndPreferredLangcodeAndMailHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt64()
	iPreferredLangcode := self.Args("preferred_langcode").String()
	iMail := self.Args("mail").String()

	if helper.IsHas(iUid) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByUidAndPreferredLangcodeAndMail(offset, limit, iUid,iPreferredLangcode,iMail)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByUidAndPreferredLangcodeAndMail's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByUidAndPreferredLangcodeAndTimezoneHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt64()
	iPreferredLangcode := self.Args("preferred_langcode").String()
	iTimezone := self.Args("timezone").String()

	if helper.IsHas(iUid) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByUidAndPreferredLangcodeAndTimezone(offset, limit, iUid,iPreferredLangcode,iTimezone)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByUidAndPreferredLangcodeAndTimezone's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByUidAndPreferredLangcodeAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt64()
	iPreferredLangcode := self.Args("preferred_langcode").String()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iUid) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByUidAndPreferredLangcodeAndStatus(offset, limit, iUid,iPreferredLangcode,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByUidAndPreferredLangcodeAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByUidAndPreferredLangcodeAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt64()
	iPreferredLangcode := self.Args("preferred_langcode").String()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iUid) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByUidAndPreferredLangcodeAndCreated(offset, limit, iUid,iPreferredLangcode,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByUidAndPreferredLangcodeAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByUidAndPreferredLangcodeAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt64()
	iPreferredLangcode := self.Args("preferred_langcode").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iUid) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByUidAndPreferredLangcodeAndChanged(offset, limit, iUid,iPreferredLangcode,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByUidAndPreferredLangcodeAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByUidAndPreferredLangcodeAndAccessHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt64()
	iPreferredLangcode := self.Args("preferred_langcode").String()
	iAccess := self.Args("access").MustInt()

	if helper.IsHas(iUid) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByUidAndPreferredLangcodeAndAccess(offset, limit, iUid,iPreferredLangcode,iAccess)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByUidAndPreferredLangcodeAndAccess's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByUidAndPreferredLangcodeAndLoginHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt64()
	iPreferredLangcode := self.Args("preferred_langcode").String()
	iLogin := self.Args("login").MustInt()

	if helper.IsHas(iUid) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByUidAndPreferredLangcodeAndLogin(offset, limit, iUid,iPreferredLangcode,iLogin)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByUidAndPreferredLangcodeAndLogin's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByUidAndPreferredLangcodeAndInitHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt64()
	iPreferredLangcode := self.Args("preferred_langcode").String()
	iInit := self.Args("init").String()

	if helper.IsHas(iUid) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByUidAndPreferredLangcodeAndInit(offset, limit, iUid,iPreferredLangcode,iInit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByUidAndPreferredLangcodeAndInit's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByUidAndPreferredLangcodeAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt64()
	iPreferredLangcode := self.Args("preferred_langcode").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iUid) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByUidAndPreferredLangcodeAndDefaultLangcode(offset, limit, iUid,iPreferredLangcode,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByUidAndPreferredLangcodeAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByUidAndPreferredAdminLangcodeAndNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt64()
	iPreferredAdminLangcode := self.Args("preferred_admin_langcode").String()
	iName := self.Args("name").String()

	if helper.IsHas(iUid) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByUidAndPreferredAdminLangcodeAndName(offset, limit, iUid,iPreferredAdminLangcode,iName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByUidAndPreferredAdminLangcodeAndName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByUidAndPreferredAdminLangcodeAndPassHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt64()
	iPreferredAdminLangcode := self.Args("preferred_admin_langcode").String()
	iPass := self.Args("pass").String()

	if helper.IsHas(iUid) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByUidAndPreferredAdminLangcodeAndPass(offset, limit, iUid,iPreferredAdminLangcode,iPass)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByUidAndPreferredAdminLangcodeAndPass's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByUidAndPreferredAdminLangcodeAndMailHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt64()
	iPreferredAdminLangcode := self.Args("preferred_admin_langcode").String()
	iMail := self.Args("mail").String()

	if helper.IsHas(iUid) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByUidAndPreferredAdminLangcodeAndMail(offset, limit, iUid,iPreferredAdminLangcode,iMail)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByUidAndPreferredAdminLangcodeAndMail's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByUidAndPreferredAdminLangcodeAndTimezoneHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt64()
	iPreferredAdminLangcode := self.Args("preferred_admin_langcode").String()
	iTimezone := self.Args("timezone").String()

	if helper.IsHas(iUid) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByUidAndPreferredAdminLangcodeAndTimezone(offset, limit, iUid,iPreferredAdminLangcode,iTimezone)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByUidAndPreferredAdminLangcodeAndTimezone's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByUidAndPreferredAdminLangcodeAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt64()
	iPreferredAdminLangcode := self.Args("preferred_admin_langcode").String()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iUid) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByUidAndPreferredAdminLangcodeAndStatus(offset, limit, iUid,iPreferredAdminLangcode,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByUidAndPreferredAdminLangcodeAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByUidAndPreferredAdminLangcodeAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt64()
	iPreferredAdminLangcode := self.Args("preferred_admin_langcode").String()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iUid) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByUidAndPreferredAdminLangcodeAndCreated(offset, limit, iUid,iPreferredAdminLangcode,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByUidAndPreferredAdminLangcodeAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByUidAndPreferredAdminLangcodeAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt64()
	iPreferredAdminLangcode := self.Args("preferred_admin_langcode").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iUid) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByUidAndPreferredAdminLangcodeAndChanged(offset, limit, iUid,iPreferredAdminLangcode,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByUidAndPreferredAdminLangcodeAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByUidAndPreferredAdminLangcodeAndAccessHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt64()
	iPreferredAdminLangcode := self.Args("preferred_admin_langcode").String()
	iAccess := self.Args("access").MustInt()

	if helper.IsHas(iUid) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByUidAndPreferredAdminLangcodeAndAccess(offset, limit, iUid,iPreferredAdminLangcode,iAccess)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByUidAndPreferredAdminLangcodeAndAccess's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByUidAndPreferredAdminLangcodeAndLoginHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt64()
	iPreferredAdminLangcode := self.Args("preferred_admin_langcode").String()
	iLogin := self.Args("login").MustInt()

	if helper.IsHas(iUid) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByUidAndPreferredAdminLangcodeAndLogin(offset, limit, iUid,iPreferredAdminLangcode,iLogin)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByUidAndPreferredAdminLangcodeAndLogin's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByUidAndPreferredAdminLangcodeAndInitHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt64()
	iPreferredAdminLangcode := self.Args("preferred_admin_langcode").String()
	iInit := self.Args("init").String()

	if helper.IsHas(iUid) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByUidAndPreferredAdminLangcodeAndInit(offset, limit, iUid,iPreferredAdminLangcode,iInit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByUidAndPreferredAdminLangcodeAndInit's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByUidAndPreferredAdminLangcodeAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt64()
	iPreferredAdminLangcode := self.Args("preferred_admin_langcode").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iUid) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByUidAndPreferredAdminLangcodeAndDefaultLangcode(offset, limit, iUid,iPreferredAdminLangcode,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByUidAndPreferredAdminLangcodeAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByUidAndNameAndPassHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt64()
	iName := self.Args("name").String()
	iPass := self.Args("pass").String()

	if helper.IsHas(iUid) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByUidAndNameAndPass(offset, limit, iUid,iName,iPass)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByUidAndNameAndPass's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByUidAndNameAndMailHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt64()
	iName := self.Args("name").String()
	iMail := self.Args("mail").String()

	if helper.IsHas(iUid) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByUidAndNameAndMail(offset, limit, iUid,iName,iMail)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByUidAndNameAndMail's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByUidAndNameAndTimezoneHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt64()
	iName := self.Args("name").String()
	iTimezone := self.Args("timezone").String()

	if helper.IsHas(iUid) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByUidAndNameAndTimezone(offset, limit, iUid,iName,iTimezone)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByUidAndNameAndTimezone's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByUidAndNameAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt64()
	iName := self.Args("name").String()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iUid) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByUidAndNameAndStatus(offset, limit, iUid,iName,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByUidAndNameAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByUidAndNameAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt64()
	iName := self.Args("name").String()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iUid) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByUidAndNameAndCreated(offset, limit, iUid,iName,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByUidAndNameAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByUidAndNameAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt64()
	iName := self.Args("name").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iUid) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByUidAndNameAndChanged(offset, limit, iUid,iName,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByUidAndNameAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByUidAndNameAndAccessHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt64()
	iName := self.Args("name").String()
	iAccess := self.Args("access").MustInt()

	if helper.IsHas(iUid) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByUidAndNameAndAccess(offset, limit, iUid,iName,iAccess)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByUidAndNameAndAccess's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByUidAndNameAndLoginHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt64()
	iName := self.Args("name").String()
	iLogin := self.Args("login").MustInt()

	if helper.IsHas(iUid) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByUidAndNameAndLogin(offset, limit, iUid,iName,iLogin)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByUidAndNameAndLogin's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByUidAndNameAndInitHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt64()
	iName := self.Args("name").String()
	iInit := self.Args("init").String()

	if helper.IsHas(iUid) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByUidAndNameAndInit(offset, limit, iUid,iName,iInit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByUidAndNameAndInit's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByUidAndNameAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt64()
	iName := self.Args("name").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iUid) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByUidAndNameAndDefaultLangcode(offset, limit, iUid,iName,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByUidAndNameAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByUidAndPassAndMailHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt64()
	iPass := self.Args("pass").String()
	iMail := self.Args("mail").String()

	if helper.IsHas(iUid) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByUidAndPassAndMail(offset, limit, iUid,iPass,iMail)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByUidAndPassAndMail's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByUidAndPassAndTimezoneHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt64()
	iPass := self.Args("pass").String()
	iTimezone := self.Args("timezone").String()

	if helper.IsHas(iUid) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByUidAndPassAndTimezone(offset, limit, iUid,iPass,iTimezone)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByUidAndPassAndTimezone's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByUidAndPassAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt64()
	iPass := self.Args("pass").String()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iUid) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByUidAndPassAndStatus(offset, limit, iUid,iPass,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByUidAndPassAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByUidAndPassAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt64()
	iPass := self.Args("pass").String()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iUid) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByUidAndPassAndCreated(offset, limit, iUid,iPass,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByUidAndPassAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByUidAndPassAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt64()
	iPass := self.Args("pass").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iUid) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByUidAndPassAndChanged(offset, limit, iUid,iPass,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByUidAndPassAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByUidAndPassAndAccessHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt64()
	iPass := self.Args("pass").String()
	iAccess := self.Args("access").MustInt()

	if helper.IsHas(iUid) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByUidAndPassAndAccess(offset, limit, iUid,iPass,iAccess)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByUidAndPassAndAccess's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByUidAndPassAndLoginHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt64()
	iPass := self.Args("pass").String()
	iLogin := self.Args("login").MustInt()

	if helper.IsHas(iUid) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByUidAndPassAndLogin(offset, limit, iUid,iPass,iLogin)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByUidAndPassAndLogin's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByUidAndPassAndInitHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt64()
	iPass := self.Args("pass").String()
	iInit := self.Args("init").String()

	if helper.IsHas(iUid) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByUidAndPassAndInit(offset, limit, iUid,iPass,iInit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByUidAndPassAndInit's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByUidAndPassAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt64()
	iPass := self.Args("pass").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iUid) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByUidAndPassAndDefaultLangcode(offset, limit, iUid,iPass,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByUidAndPassAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByUidAndMailAndTimezoneHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt64()
	iMail := self.Args("mail").String()
	iTimezone := self.Args("timezone").String()

	if helper.IsHas(iUid) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByUidAndMailAndTimezone(offset, limit, iUid,iMail,iTimezone)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByUidAndMailAndTimezone's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByUidAndMailAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt64()
	iMail := self.Args("mail").String()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iUid) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByUidAndMailAndStatus(offset, limit, iUid,iMail,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByUidAndMailAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByUidAndMailAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt64()
	iMail := self.Args("mail").String()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iUid) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByUidAndMailAndCreated(offset, limit, iUid,iMail,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByUidAndMailAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByUidAndMailAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt64()
	iMail := self.Args("mail").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iUid) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByUidAndMailAndChanged(offset, limit, iUid,iMail,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByUidAndMailAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByUidAndMailAndAccessHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt64()
	iMail := self.Args("mail").String()
	iAccess := self.Args("access").MustInt()

	if helper.IsHas(iUid) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByUidAndMailAndAccess(offset, limit, iUid,iMail,iAccess)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByUidAndMailAndAccess's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByUidAndMailAndLoginHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt64()
	iMail := self.Args("mail").String()
	iLogin := self.Args("login").MustInt()

	if helper.IsHas(iUid) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByUidAndMailAndLogin(offset, limit, iUid,iMail,iLogin)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByUidAndMailAndLogin's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByUidAndMailAndInitHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt64()
	iMail := self.Args("mail").String()
	iInit := self.Args("init").String()

	if helper.IsHas(iUid) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByUidAndMailAndInit(offset, limit, iUid,iMail,iInit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByUidAndMailAndInit's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByUidAndMailAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt64()
	iMail := self.Args("mail").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iUid) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByUidAndMailAndDefaultLangcode(offset, limit, iUid,iMail,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByUidAndMailAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByUidAndTimezoneAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt64()
	iTimezone := self.Args("timezone").String()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iUid) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByUidAndTimezoneAndStatus(offset, limit, iUid,iTimezone,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByUidAndTimezoneAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByUidAndTimezoneAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt64()
	iTimezone := self.Args("timezone").String()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iUid) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByUidAndTimezoneAndCreated(offset, limit, iUid,iTimezone,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByUidAndTimezoneAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByUidAndTimezoneAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt64()
	iTimezone := self.Args("timezone").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iUid) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByUidAndTimezoneAndChanged(offset, limit, iUid,iTimezone,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByUidAndTimezoneAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByUidAndTimezoneAndAccessHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt64()
	iTimezone := self.Args("timezone").String()
	iAccess := self.Args("access").MustInt()

	if helper.IsHas(iUid) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByUidAndTimezoneAndAccess(offset, limit, iUid,iTimezone,iAccess)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByUidAndTimezoneAndAccess's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByUidAndTimezoneAndLoginHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt64()
	iTimezone := self.Args("timezone").String()
	iLogin := self.Args("login").MustInt()

	if helper.IsHas(iUid) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByUidAndTimezoneAndLogin(offset, limit, iUid,iTimezone,iLogin)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByUidAndTimezoneAndLogin's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByUidAndTimezoneAndInitHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt64()
	iTimezone := self.Args("timezone").String()
	iInit := self.Args("init").String()

	if helper.IsHas(iUid) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByUidAndTimezoneAndInit(offset, limit, iUid,iTimezone,iInit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByUidAndTimezoneAndInit's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByUidAndTimezoneAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt64()
	iTimezone := self.Args("timezone").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iUid) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByUidAndTimezoneAndDefaultLangcode(offset, limit, iUid,iTimezone,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByUidAndTimezoneAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByUidAndStatusAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt64()
	iStatus := self.Args("status").MustInt()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iUid) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByUidAndStatusAndCreated(offset, limit, iUid,iStatus,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByUidAndStatusAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByUidAndStatusAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt64()
	iStatus := self.Args("status").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iUid) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByUidAndStatusAndChanged(offset, limit, iUid,iStatus,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByUidAndStatusAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByUidAndStatusAndAccessHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt64()
	iStatus := self.Args("status").MustInt()
	iAccess := self.Args("access").MustInt()

	if helper.IsHas(iUid) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByUidAndStatusAndAccess(offset, limit, iUid,iStatus,iAccess)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByUidAndStatusAndAccess's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByUidAndStatusAndLoginHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt64()
	iStatus := self.Args("status").MustInt()
	iLogin := self.Args("login").MustInt()

	if helper.IsHas(iUid) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByUidAndStatusAndLogin(offset, limit, iUid,iStatus,iLogin)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByUidAndStatusAndLogin's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByUidAndStatusAndInitHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt64()
	iStatus := self.Args("status").MustInt()
	iInit := self.Args("init").String()

	if helper.IsHas(iUid) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByUidAndStatusAndInit(offset, limit, iUid,iStatus,iInit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByUidAndStatusAndInit's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByUidAndStatusAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt64()
	iStatus := self.Args("status").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iUid) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByUidAndStatusAndDefaultLangcode(offset, limit, iUid,iStatus,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByUidAndStatusAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByUidAndCreatedAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt64()
	iCreated := self.Args("created").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iUid) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByUidAndCreatedAndChanged(offset, limit, iUid,iCreated,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByUidAndCreatedAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByUidAndCreatedAndAccessHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt64()
	iCreated := self.Args("created").MustInt()
	iAccess := self.Args("access").MustInt()

	if helper.IsHas(iUid) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByUidAndCreatedAndAccess(offset, limit, iUid,iCreated,iAccess)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByUidAndCreatedAndAccess's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByUidAndCreatedAndLoginHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt64()
	iCreated := self.Args("created").MustInt()
	iLogin := self.Args("login").MustInt()

	if helper.IsHas(iUid) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByUidAndCreatedAndLogin(offset, limit, iUid,iCreated,iLogin)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByUidAndCreatedAndLogin's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByUidAndCreatedAndInitHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt64()
	iCreated := self.Args("created").MustInt()
	iInit := self.Args("init").String()

	if helper.IsHas(iUid) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByUidAndCreatedAndInit(offset, limit, iUid,iCreated,iInit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByUidAndCreatedAndInit's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByUidAndCreatedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt64()
	iCreated := self.Args("created").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iUid) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByUidAndCreatedAndDefaultLangcode(offset, limit, iUid,iCreated,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByUidAndCreatedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByUidAndChangedAndAccessHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt64()
	iChanged := self.Args("changed").MustInt()
	iAccess := self.Args("access").MustInt()

	if helper.IsHas(iUid) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByUidAndChangedAndAccess(offset, limit, iUid,iChanged,iAccess)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByUidAndChangedAndAccess's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByUidAndChangedAndLoginHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt64()
	iChanged := self.Args("changed").MustInt()
	iLogin := self.Args("login").MustInt()

	if helper.IsHas(iUid) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByUidAndChangedAndLogin(offset, limit, iUid,iChanged,iLogin)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByUidAndChangedAndLogin's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByUidAndChangedAndInitHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt64()
	iChanged := self.Args("changed").MustInt()
	iInit := self.Args("init").String()

	if helper.IsHas(iUid) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByUidAndChangedAndInit(offset, limit, iUid,iChanged,iInit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByUidAndChangedAndInit's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByUidAndChangedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt64()
	iChanged := self.Args("changed").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iUid) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByUidAndChangedAndDefaultLangcode(offset, limit, iUid,iChanged,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByUidAndChangedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByUidAndAccessAndLoginHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt64()
	iAccess := self.Args("access").MustInt()
	iLogin := self.Args("login").MustInt()

	if helper.IsHas(iUid) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByUidAndAccessAndLogin(offset, limit, iUid,iAccess,iLogin)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByUidAndAccessAndLogin's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByUidAndAccessAndInitHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt64()
	iAccess := self.Args("access").MustInt()
	iInit := self.Args("init").String()

	if helper.IsHas(iUid) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByUidAndAccessAndInit(offset, limit, iUid,iAccess,iInit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByUidAndAccessAndInit's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByUidAndAccessAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt64()
	iAccess := self.Args("access").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iUid) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByUidAndAccessAndDefaultLangcode(offset, limit, iUid,iAccess,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByUidAndAccessAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByUidAndLoginAndInitHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt64()
	iLogin := self.Args("login").MustInt()
	iInit := self.Args("init").String()

	if helper.IsHas(iUid) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByUidAndLoginAndInit(offset, limit, iUid,iLogin,iInit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByUidAndLoginAndInit's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByUidAndLoginAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt64()
	iLogin := self.Args("login").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iUid) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByUidAndLoginAndDefaultLangcode(offset, limit, iUid,iLogin,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByUidAndLoginAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByUidAndInitAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt64()
	iInit := self.Args("init").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iUid) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByUidAndInitAndDefaultLangcode(offset, limit, iUid,iInit,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByUidAndInitAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByLangcodeAndPreferredLangcodeAndPreferredAdminLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iPreferredLangcode := self.Args("preferred_langcode").String()
	iPreferredAdminLangcode := self.Args("preferred_admin_langcode").String()

	if helper.IsHas(iLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByLangcodeAndPreferredLangcodeAndPreferredAdminLangcode(offset, limit, iLangcode,iPreferredLangcode,iPreferredAdminLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByLangcodeAndPreferredLangcodeAndPreferredAdminLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByLangcodeAndPreferredLangcodeAndNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iPreferredLangcode := self.Args("preferred_langcode").String()
	iName := self.Args("name").String()

	if helper.IsHas(iLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByLangcodeAndPreferredLangcodeAndName(offset, limit, iLangcode,iPreferredLangcode,iName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByLangcodeAndPreferredLangcodeAndName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByLangcodeAndPreferredLangcodeAndPassHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iPreferredLangcode := self.Args("preferred_langcode").String()
	iPass := self.Args("pass").String()

	if helper.IsHas(iLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByLangcodeAndPreferredLangcodeAndPass(offset, limit, iLangcode,iPreferredLangcode,iPass)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByLangcodeAndPreferredLangcodeAndPass's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByLangcodeAndPreferredLangcodeAndMailHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iPreferredLangcode := self.Args("preferred_langcode").String()
	iMail := self.Args("mail").String()

	if helper.IsHas(iLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByLangcodeAndPreferredLangcodeAndMail(offset, limit, iLangcode,iPreferredLangcode,iMail)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByLangcodeAndPreferredLangcodeAndMail's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByLangcodeAndPreferredLangcodeAndTimezoneHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iPreferredLangcode := self.Args("preferred_langcode").String()
	iTimezone := self.Args("timezone").String()

	if helper.IsHas(iLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByLangcodeAndPreferredLangcodeAndTimezone(offset, limit, iLangcode,iPreferredLangcode,iTimezone)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByLangcodeAndPreferredLangcodeAndTimezone's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByLangcodeAndPreferredLangcodeAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iPreferredLangcode := self.Args("preferred_langcode").String()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByLangcodeAndPreferredLangcodeAndStatus(offset, limit, iLangcode,iPreferredLangcode,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByLangcodeAndPreferredLangcodeAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByLangcodeAndPreferredLangcodeAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iPreferredLangcode := self.Args("preferred_langcode").String()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByLangcodeAndPreferredLangcodeAndCreated(offset, limit, iLangcode,iPreferredLangcode,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByLangcodeAndPreferredLangcodeAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByLangcodeAndPreferredLangcodeAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iPreferredLangcode := self.Args("preferred_langcode").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByLangcodeAndPreferredLangcodeAndChanged(offset, limit, iLangcode,iPreferredLangcode,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByLangcodeAndPreferredLangcodeAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByLangcodeAndPreferredLangcodeAndAccessHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iPreferredLangcode := self.Args("preferred_langcode").String()
	iAccess := self.Args("access").MustInt()

	if helper.IsHas(iLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByLangcodeAndPreferredLangcodeAndAccess(offset, limit, iLangcode,iPreferredLangcode,iAccess)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByLangcodeAndPreferredLangcodeAndAccess's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByLangcodeAndPreferredLangcodeAndLoginHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iPreferredLangcode := self.Args("preferred_langcode").String()
	iLogin := self.Args("login").MustInt()

	if helper.IsHas(iLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByLangcodeAndPreferredLangcodeAndLogin(offset, limit, iLangcode,iPreferredLangcode,iLogin)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByLangcodeAndPreferredLangcodeAndLogin's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByLangcodeAndPreferredLangcodeAndInitHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iPreferredLangcode := self.Args("preferred_langcode").String()
	iInit := self.Args("init").String()

	if helper.IsHas(iLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByLangcodeAndPreferredLangcodeAndInit(offset, limit, iLangcode,iPreferredLangcode,iInit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByLangcodeAndPreferredLangcodeAndInit's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByLangcodeAndPreferredLangcodeAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iPreferredLangcode := self.Args("preferred_langcode").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByLangcodeAndPreferredLangcodeAndDefaultLangcode(offset, limit, iLangcode,iPreferredLangcode,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByLangcodeAndPreferredLangcodeAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByLangcodeAndPreferredAdminLangcodeAndNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iPreferredAdminLangcode := self.Args("preferred_admin_langcode").String()
	iName := self.Args("name").String()

	if helper.IsHas(iLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByLangcodeAndPreferredAdminLangcodeAndName(offset, limit, iLangcode,iPreferredAdminLangcode,iName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByLangcodeAndPreferredAdminLangcodeAndName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByLangcodeAndPreferredAdminLangcodeAndPassHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iPreferredAdminLangcode := self.Args("preferred_admin_langcode").String()
	iPass := self.Args("pass").String()

	if helper.IsHas(iLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByLangcodeAndPreferredAdminLangcodeAndPass(offset, limit, iLangcode,iPreferredAdminLangcode,iPass)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByLangcodeAndPreferredAdminLangcodeAndPass's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByLangcodeAndPreferredAdminLangcodeAndMailHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iPreferredAdminLangcode := self.Args("preferred_admin_langcode").String()
	iMail := self.Args("mail").String()

	if helper.IsHas(iLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByLangcodeAndPreferredAdminLangcodeAndMail(offset, limit, iLangcode,iPreferredAdminLangcode,iMail)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByLangcodeAndPreferredAdminLangcodeAndMail's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByLangcodeAndPreferredAdminLangcodeAndTimezoneHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iPreferredAdminLangcode := self.Args("preferred_admin_langcode").String()
	iTimezone := self.Args("timezone").String()

	if helper.IsHas(iLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByLangcodeAndPreferredAdminLangcodeAndTimezone(offset, limit, iLangcode,iPreferredAdminLangcode,iTimezone)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByLangcodeAndPreferredAdminLangcodeAndTimezone's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByLangcodeAndPreferredAdminLangcodeAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iPreferredAdminLangcode := self.Args("preferred_admin_langcode").String()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByLangcodeAndPreferredAdminLangcodeAndStatus(offset, limit, iLangcode,iPreferredAdminLangcode,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByLangcodeAndPreferredAdminLangcodeAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByLangcodeAndPreferredAdminLangcodeAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iPreferredAdminLangcode := self.Args("preferred_admin_langcode").String()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByLangcodeAndPreferredAdminLangcodeAndCreated(offset, limit, iLangcode,iPreferredAdminLangcode,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByLangcodeAndPreferredAdminLangcodeAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByLangcodeAndPreferredAdminLangcodeAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iPreferredAdminLangcode := self.Args("preferred_admin_langcode").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByLangcodeAndPreferredAdminLangcodeAndChanged(offset, limit, iLangcode,iPreferredAdminLangcode,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByLangcodeAndPreferredAdminLangcodeAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByLangcodeAndPreferredAdminLangcodeAndAccessHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iPreferredAdminLangcode := self.Args("preferred_admin_langcode").String()
	iAccess := self.Args("access").MustInt()

	if helper.IsHas(iLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByLangcodeAndPreferredAdminLangcodeAndAccess(offset, limit, iLangcode,iPreferredAdminLangcode,iAccess)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByLangcodeAndPreferredAdminLangcodeAndAccess's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByLangcodeAndPreferredAdminLangcodeAndLoginHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iPreferredAdminLangcode := self.Args("preferred_admin_langcode").String()
	iLogin := self.Args("login").MustInt()

	if helper.IsHas(iLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByLangcodeAndPreferredAdminLangcodeAndLogin(offset, limit, iLangcode,iPreferredAdminLangcode,iLogin)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByLangcodeAndPreferredAdminLangcodeAndLogin's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByLangcodeAndPreferredAdminLangcodeAndInitHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iPreferredAdminLangcode := self.Args("preferred_admin_langcode").String()
	iInit := self.Args("init").String()

	if helper.IsHas(iLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByLangcodeAndPreferredAdminLangcodeAndInit(offset, limit, iLangcode,iPreferredAdminLangcode,iInit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByLangcodeAndPreferredAdminLangcodeAndInit's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByLangcodeAndPreferredAdminLangcodeAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iPreferredAdminLangcode := self.Args("preferred_admin_langcode").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByLangcodeAndPreferredAdminLangcodeAndDefaultLangcode(offset, limit, iLangcode,iPreferredAdminLangcode,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByLangcodeAndPreferredAdminLangcodeAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByLangcodeAndNameAndPassHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iName := self.Args("name").String()
	iPass := self.Args("pass").String()

	if helper.IsHas(iLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByLangcodeAndNameAndPass(offset, limit, iLangcode,iName,iPass)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByLangcodeAndNameAndPass's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByLangcodeAndNameAndMailHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iName := self.Args("name").String()
	iMail := self.Args("mail").String()

	if helper.IsHas(iLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByLangcodeAndNameAndMail(offset, limit, iLangcode,iName,iMail)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByLangcodeAndNameAndMail's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByLangcodeAndNameAndTimezoneHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iName := self.Args("name").String()
	iTimezone := self.Args("timezone").String()

	if helper.IsHas(iLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByLangcodeAndNameAndTimezone(offset, limit, iLangcode,iName,iTimezone)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByLangcodeAndNameAndTimezone's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByLangcodeAndNameAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iName := self.Args("name").String()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByLangcodeAndNameAndStatus(offset, limit, iLangcode,iName,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByLangcodeAndNameAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByLangcodeAndNameAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iName := self.Args("name").String()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByLangcodeAndNameAndCreated(offset, limit, iLangcode,iName,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByLangcodeAndNameAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByLangcodeAndNameAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iName := self.Args("name").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByLangcodeAndNameAndChanged(offset, limit, iLangcode,iName,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByLangcodeAndNameAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByLangcodeAndNameAndAccessHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iName := self.Args("name").String()
	iAccess := self.Args("access").MustInt()

	if helper.IsHas(iLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByLangcodeAndNameAndAccess(offset, limit, iLangcode,iName,iAccess)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByLangcodeAndNameAndAccess's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByLangcodeAndNameAndLoginHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iName := self.Args("name").String()
	iLogin := self.Args("login").MustInt()

	if helper.IsHas(iLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByLangcodeAndNameAndLogin(offset, limit, iLangcode,iName,iLogin)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByLangcodeAndNameAndLogin's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByLangcodeAndNameAndInitHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iName := self.Args("name").String()
	iInit := self.Args("init").String()

	if helper.IsHas(iLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByLangcodeAndNameAndInit(offset, limit, iLangcode,iName,iInit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByLangcodeAndNameAndInit's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByLangcodeAndNameAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iName := self.Args("name").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByLangcodeAndNameAndDefaultLangcode(offset, limit, iLangcode,iName,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByLangcodeAndNameAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByLangcodeAndPassAndMailHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iPass := self.Args("pass").String()
	iMail := self.Args("mail").String()

	if helper.IsHas(iLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByLangcodeAndPassAndMail(offset, limit, iLangcode,iPass,iMail)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByLangcodeAndPassAndMail's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByLangcodeAndPassAndTimezoneHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iPass := self.Args("pass").String()
	iTimezone := self.Args("timezone").String()

	if helper.IsHas(iLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByLangcodeAndPassAndTimezone(offset, limit, iLangcode,iPass,iTimezone)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByLangcodeAndPassAndTimezone's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByLangcodeAndPassAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iPass := self.Args("pass").String()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByLangcodeAndPassAndStatus(offset, limit, iLangcode,iPass,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByLangcodeAndPassAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByLangcodeAndPassAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iPass := self.Args("pass").String()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByLangcodeAndPassAndCreated(offset, limit, iLangcode,iPass,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByLangcodeAndPassAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByLangcodeAndPassAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iPass := self.Args("pass").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByLangcodeAndPassAndChanged(offset, limit, iLangcode,iPass,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByLangcodeAndPassAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByLangcodeAndPassAndAccessHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iPass := self.Args("pass").String()
	iAccess := self.Args("access").MustInt()

	if helper.IsHas(iLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByLangcodeAndPassAndAccess(offset, limit, iLangcode,iPass,iAccess)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByLangcodeAndPassAndAccess's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByLangcodeAndPassAndLoginHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iPass := self.Args("pass").String()
	iLogin := self.Args("login").MustInt()

	if helper.IsHas(iLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByLangcodeAndPassAndLogin(offset, limit, iLangcode,iPass,iLogin)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByLangcodeAndPassAndLogin's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByLangcodeAndPassAndInitHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iPass := self.Args("pass").String()
	iInit := self.Args("init").String()

	if helper.IsHas(iLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByLangcodeAndPassAndInit(offset, limit, iLangcode,iPass,iInit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByLangcodeAndPassAndInit's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByLangcodeAndPassAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iPass := self.Args("pass").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByLangcodeAndPassAndDefaultLangcode(offset, limit, iLangcode,iPass,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByLangcodeAndPassAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByLangcodeAndMailAndTimezoneHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iMail := self.Args("mail").String()
	iTimezone := self.Args("timezone").String()

	if helper.IsHas(iLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByLangcodeAndMailAndTimezone(offset, limit, iLangcode,iMail,iTimezone)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByLangcodeAndMailAndTimezone's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByLangcodeAndMailAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iMail := self.Args("mail").String()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByLangcodeAndMailAndStatus(offset, limit, iLangcode,iMail,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByLangcodeAndMailAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByLangcodeAndMailAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iMail := self.Args("mail").String()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByLangcodeAndMailAndCreated(offset, limit, iLangcode,iMail,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByLangcodeAndMailAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByLangcodeAndMailAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iMail := self.Args("mail").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByLangcodeAndMailAndChanged(offset, limit, iLangcode,iMail,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByLangcodeAndMailAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByLangcodeAndMailAndAccessHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iMail := self.Args("mail").String()
	iAccess := self.Args("access").MustInt()

	if helper.IsHas(iLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByLangcodeAndMailAndAccess(offset, limit, iLangcode,iMail,iAccess)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByLangcodeAndMailAndAccess's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByLangcodeAndMailAndLoginHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iMail := self.Args("mail").String()
	iLogin := self.Args("login").MustInt()

	if helper.IsHas(iLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByLangcodeAndMailAndLogin(offset, limit, iLangcode,iMail,iLogin)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByLangcodeAndMailAndLogin's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByLangcodeAndMailAndInitHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iMail := self.Args("mail").String()
	iInit := self.Args("init").String()

	if helper.IsHas(iLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByLangcodeAndMailAndInit(offset, limit, iLangcode,iMail,iInit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByLangcodeAndMailAndInit's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByLangcodeAndMailAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iMail := self.Args("mail").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByLangcodeAndMailAndDefaultLangcode(offset, limit, iLangcode,iMail,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByLangcodeAndMailAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByLangcodeAndTimezoneAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iTimezone := self.Args("timezone").String()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByLangcodeAndTimezoneAndStatus(offset, limit, iLangcode,iTimezone,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByLangcodeAndTimezoneAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByLangcodeAndTimezoneAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iTimezone := self.Args("timezone").String()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByLangcodeAndTimezoneAndCreated(offset, limit, iLangcode,iTimezone,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByLangcodeAndTimezoneAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByLangcodeAndTimezoneAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iTimezone := self.Args("timezone").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByLangcodeAndTimezoneAndChanged(offset, limit, iLangcode,iTimezone,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByLangcodeAndTimezoneAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByLangcodeAndTimezoneAndAccessHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iTimezone := self.Args("timezone").String()
	iAccess := self.Args("access").MustInt()

	if helper.IsHas(iLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByLangcodeAndTimezoneAndAccess(offset, limit, iLangcode,iTimezone,iAccess)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByLangcodeAndTimezoneAndAccess's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByLangcodeAndTimezoneAndLoginHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iTimezone := self.Args("timezone").String()
	iLogin := self.Args("login").MustInt()

	if helper.IsHas(iLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByLangcodeAndTimezoneAndLogin(offset, limit, iLangcode,iTimezone,iLogin)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByLangcodeAndTimezoneAndLogin's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByLangcodeAndTimezoneAndInitHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iTimezone := self.Args("timezone").String()
	iInit := self.Args("init").String()

	if helper.IsHas(iLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByLangcodeAndTimezoneAndInit(offset, limit, iLangcode,iTimezone,iInit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByLangcodeAndTimezoneAndInit's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByLangcodeAndTimezoneAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iTimezone := self.Args("timezone").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByLangcodeAndTimezoneAndDefaultLangcode(offset, limit, iLangcode,iTimezone,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByLangcodeAndTimezoneAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByLangcodeAndStatusAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iStatus := self.Args("status").MustInt()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByLangcodeAndStatusAndCreated(offset, limit, iLangcode,iStatus,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByLangcodeAndStatusAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByLangcodeAndStatusAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iStatus := self.Args("status").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByLangcodeAndStatusAndChanged(offset, limit, iLangcode,iStatus,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByLangcodeAndStatusAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByLangcodeAndStatusAndAccessHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iStatus := self.Args("status").MustInt()
	iAccess := self.Args("access").MustInt()

	if helper.IsHas(iLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByLangcodeAndStatusAndAccess(offset, limit, iLangcode,iStatus,iAccess)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByLangcodeAndStatusAndAccess's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByLangcodeAndStatusAndLoginHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iStatus := self.Args("status").MustInt()
	iLogin := self.Args("login").MustInt()

	if helper.IsHas(iLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByLangcodeAndStatusAndLogin(offset, limit, iLangcode,iStatus,iLogin)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByLangcodeAndStatusAndLogin's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByLangcodeAndStatusAndInitHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iStatus := self.Args("status").MustInt()
	iInit := self.Args("init").String()

	if helper.IsHas(iLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByLangcodeAndStatusAndInit(offset, limit, iLangcode,iStatus,iInit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByLangcodeAndStatusAndInit's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByLangcodeAndStatusAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iStatus := self.Args("status").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByLangcodeAndStatusAndDefaultLangcode(offset, limit, iLangcode,iStatus,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByLangcodeAndStatusAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByLangcodeAndCreatedAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iCreated := self.Args("created").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByLangcodeAndCreatedAndChanged(offset, limit, iLangcode,iCreated,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByLangcodeAndCreatedAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByLangcodeAndCreatedAndAccessHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iCreated := self.Args("created").MustInt()
	iAccess := self.Args("access").MustInt()

	if helper.IsHas(iLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByLangcodeAndCreatedAndAccess(offset, limit, iLangcode,iCreated,iAccess)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByLangcodeAndCreatedAndAccess's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByLangcodeAndCreatedAndLoginHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iCreated := self.Args("created").MustInt()
	iLogin := self.Args("login").MustInt()

	if helper.IsHas(iLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByLangcodeAndCreatedAndLogin(offset, limit, iLangcode,iCreated,iLogin)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByLangcodeAndCreatedAndLogin's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByLangcodeAndCreatedAndInitHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iCreated := self.Args("created").MustInt()
	iInit := self.Args("init").String()

	if helper.IsHas(iLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByLangcodeAndCreatedAndInit(offset, limit, iLangcode,iCreated,iInit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByLangcodeAndCreatedAndInit's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByLangcodeAndCreatedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iCreated := self.Args("created").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByLangcodeAndCreatedAndDefaultLangcode(offset, limit, iLangcode,iCreated,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByLangcodeAndCreatedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByLangcodeAndChangedAndAccessHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iChanged := self.Args("changed").MustInt()
	iAccess := self.Args("access").MustInt()

	if helper.IsHas(iLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByLangcodeAndChangedAndAccess(offset, limit, iLangcode,iChanged,iAccess)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByLangcodeAndChangedAndAccess's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByLangcodeAndChangedAndLoginHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iChanged := self.Args("changed").MustInt()
	iLogin := self.Args("login").MustInt()

	if helper.IsHas(iLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByLangcodeAndChangedAndLogin(offset, limit, iLangcode,iChanged,iLogin)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByLangcodeAndChangedAndLogin's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByLangcodeAndChangedAndInitHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iChanged := self.Args("changed").MustInt()
	iInit := self.Args("init").String()

	if helper.IsHas(iLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByLangcodeAndChangedAndInit(offset, limit, iLangcode,iChanged,iInit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByLangcodeAndChangedAndInit's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByLangcodeAndChangedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iChanged := self.Args("changed").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByLangcodeAndChangedAndDefaultLangcode(offset, limit, iLangcode,iChanged,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByLangcodeAndChangedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByLangcodeAndAccessAndLoginHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iAccess := self.Args("access").MustInt()
	iLogin := self.Args("login").MustInt()

	if helper.IsHas(iLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByLangcodeAndAccessAndLogin(offset, limit, iLangcode,iAccess,iLogin)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByLangcodeAndAccessAndLogin's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByLangcodeAndAccessAndInitHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iAccess := self.Args("access").MustInt()
	iInit := self.Args("init").String()

	if helper.IsHas(iLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByLangcodeAndAccessAndInit(offset, limit, iLangcode,iAccess,iInit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByLangcodeAndAccessAndInit's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByLangcodeAndAccessAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iAccess := self.Args("access").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByLangcodeAndAccessAndDefaultLangcode(offset, limit, iLangcode,iAccess,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByLangcodeAndAccessAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByLangcodeAndLoginAndInitHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iLogin := self.Args("login").MustInt()
	iInit := self.Args("init").String()

	if helper.IsHas(iLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByLangcodeAndLoginAndInit(offset, limit, iLangcode,iLogin,iInit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByLangcodeAndLoginAndInit's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByLangcodeAndLoginAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iLogin := self.Args("login").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByLangcodeAndLoginAndDefaultLangcode(offset, limit, iLangcode,iLogin,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByLangcodeAndLoginAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByLangcodeAndInitAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iInit := self.Args("init").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByLangcodeAndInitAndDefaultLangcode(offset, limit, iLangcode,iInit,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByLangcodeAndInitAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredLangcodeAndPreferredAdminLangcodeAndNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredLangcode := self.Args("preferred_langcode").String()
	iPreferredAdminLangcode := self.Args("preferred_admin_langcode").String()
	iName := self.Args("name").String()

	if helper.IsHas(iPreferredLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredLangcodeAndPreferredAdminLangcodeAndName(offset, limit, iPreferredLangcode,iPreferredAdminLangcode,iName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredLangcodeAndPreferredAdminLangcodeAndName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredLangcodeAndPreferredAdminLangcodeAndPassHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredLangcode := self.Args("preferred_langcode").String()
	iPreferredAdminLangcode := self.Args("preferred_admin_langcode").String()
	iPass := self.Args("pass").String()

	if helper.IsHas(iPreferredLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredLangcodeAndPreferredAdminLangcodeAndPass(offset, limit, iPreferredLangcode,iPreferredAdminLangcode,iPass)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredLangcodeAndPreferredAdminLangcodeAndPass's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredLangcodeAndPreferredAdminLangcodeAndMailHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredLangcode := self.Args("preferred_langcode").String()
	iPreferredAdminLangcode := self.Args("preferred_admin_langcode").String()
	iMail := self.Args("mail").String()

	if helper.IsHas(iPreferredLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredLangcodeAndPreferredAdminLangcodeAndMail(offset, limit, iPreferredLangcode,iPreferredAdminLangcode,iMail)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredLangcodeAndPreferredAdminLangcodeAndMail's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredLangcodeAndPreferredAdminLangcodeAndTimezoneHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredLangcode := self.Args("preferred_langcode").String()
	iPreferredAdminLangcode := self.Args("preferred_admin_langcode").String()
	iTimezone := self.Args("timezone").String()

	if helper.IsHas(iPreferredLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredLangcodeAndPreferredAdminLangcodeAndTimezone(offset, limit, iPreferredLangcode,iPreferredAdminLangcode,iTimezone)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredLangcodeAndPreferredAdminLangcodeAndTimezone's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredLangcodeAndPreferredAdminLangcodeAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredLangcode := self.Args("preferred_langcode").String()
	iPreferredAdminLangcode := self.Args("preferred_admin_langcode").String()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iPreferredLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredLangcodeAndPreferredAdminLangcodeAndStatus(offset, limit, iPreferredLangcode,iPreferredAdminLangcode,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredLangcodeAndPreferredAdminLangcodeAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredLangcodeAndPreferredAdminLangcodeAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredLangcode := self.Args("preferred_langcode").String()
	iPreferredAdminLangcode := self.Args("preferred_admin_langcode").String()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iPreferredLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredLangcodeAndPreferredAdminLangcodeAndCreated(offset, limit, iPreferredLangcode,iPreferredAdminLangcode,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredLangcodeAndPreferredAdminLangcodeAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredLangcodeAndPreferredAdminLangcodeAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredLangcode := self.Args("preferred_langcode").String()
	iPreferredAdminLangcode := self.Args("preferred_admin_langcode").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iPreferredLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredLangcodeAndPreferredAdminLangcodeAndChanged(offset, limit, iPreferredLangcode,iPreferredAdminLangcode,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredLangcodeAndPreferredAdminLangcodeAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredLangcodeAndPreferredAdminLangcodeAndAccessHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredLangcode := self.Args("preferred_langcode").String()
	iPreferredAdminLangcode := self.Args("preferred_admin_langcode").String()
	iAccess := self.Args("access").MustInt()

	if helper.IsHas(iPreferredLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredLangcodeAndPreferredAdminLangcodeAndAccess(offset, limit, iPreferredLangcode,iPreferredAdminLangcode,iAccess)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredLangcodeAndPreferredAdminLangcodeAndAccess's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredLangcodeAndPreferredAdminLangcodeAndLoginHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredLangcode := self.Args("preferred_langcode").String()
	iPreferredAdminLangcode := self.Args("preferred_admin_langcode").String()
	iLogin := self.Args("login").MustInt()

	if helper.IsHas(iPreferredLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredLangcodeAndPreferredAdminLangcodeAndLogin(offset, limit, iPreferredLangcode,iPreferredAdminLangcode,iLogin)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredLangcodeAndPreferredAdminLangcodeAndLogin's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredLangcodeAndPreferredAdminLangcodeAndInitHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredLangcode := self.Args("preferred_langcode").String()
	iPreferredAdminLangcode := self.Args("preferred_admin_langcode").String()
	iInit := self.Args("init").String()

	if helper.IsHas(iPreferredLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredLangcodeAndPreferredAdminLangcodeAndInit(offset, limit, iPreferredLangcode,iPreferredAdminLangcode,iInit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredLangcodeAndPreferredAdminLangcodeAndInit's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredLangcodeAndPreferredAdminLangcodeAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredLangcode := self.Args("preferred_langcode").String()
	iPreferredAdminLangcode := self.Args("preferred_admin_langcode").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iPreferredLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredLangcodeAndPreferredAdminLangcodeAndDefaultLangcode(offset, limit, iPreferredLangcode,iPreferredAdminLangcode,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredLangcodeAndPreferredAdminLangcodeAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredLangcodeAndNameAndPassHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredLangcode := self.Args("preferred_langcode").String()
	iName := self.Args("name").String()
	iPass := self.Args("pass").String()

	if helper.IsHas(iPreferredLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredLangcodeAndNameAndPass(offset, limit, iPreferredLangcode,iName,iPass)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredLangcodeAndNameAndPass's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredLangcodeAndNameAndMailHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredLangcode := self.Args("preferred_langcode").String()
	iName := self.Args("name").String()
	iMail := self.Args("mail").String()

	if helper.IsHas(iPreferredLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredLangcodeAndNameAndMail(offset, limit, iPreferredLangcode,iName,iMail)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredLangcodeAndNameAndMail's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredLangcodeAndNameAndTimezoneHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredLangcode := self.Args("preferred_langcode").String()
	iName := self.Args("name").String()
	iTimezone := self.Args("timezone").String()

	if helper.IsHas(iPreferredLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredLangcodeAndNameAndTimezone(offset, limit, iPreferredLangcode,iName,iTimezone)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredLangcodeAndNameAndTimezone's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredLangcodeAndNameAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredLangcode := self.Args("preferred_langcode").String()
	iName := self.Args("name").String()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iPreferredLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredLangcodeAndNameAndStatus(offset, limit, iPreferredLangcode,iName,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredLangcodeAndNameAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredLangcodeAndNameAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredLangcode := self.Args("preferred_langcode").String()
	iName := self.Args("name").String()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iPreferredLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredLangcodeAndNameAndCreated(offset, limit, iPreferredLangcode,iName,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredLangcodeAndNameAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredLangcodeAndNameAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredLangcode := self.Args("preferred_langcode").String()
	iName := self.Args("name").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iPreferredLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredLangcodeAndNameAndChanged(offset, limit, iPreferredLangcode,iName,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredLangcodeAndNameAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredLangcodeAndNameAndAccessHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredLangcode := self.Args("preferred_langcode").String()
	iName := self.Args("name").String()
	iAccess := self.Args("access").MustInt()

	if helper.IsHas(iPreferredLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredLangcodeAndNameAndAccess(offset, limit, iPreferredLangcode,iName,iAccess)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredLangcodeAndNameAndAccess's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredLangcodeAndNameAndLoginHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredLangcode := self.Args("preferred_langcode").String()
	iName := self.Args("name").String()
	iLogin := self.Args("login").MustInt()

	if helper.IsHas(iPreferredLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredLangcodeAndNameAndLogin(offset, limit, iPreferredLangcode,iName,iLogin)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredLangcodeAndNameAndLogin's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredLangcodeAndNameAndInitHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredLangcode := self.Args("preferred_langcode").String()
	iName := self.Args("name").String()
	iInit := self.Args("init").String()

	if helper.IsHas(iPreferredLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredLangcodeAndNameAndInit(offset, limit, iPreferredLangcode,iName,iInit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredLangcodeAndNameAndInit's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredLangcodeAndNameAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredLangcode := self.Args("preferred_langcode").String()
	iName := self.Args("name").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iPreferredLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredLangcodeAndNameAndDefaultLangcode(offset, limit, iPreferredLangcode,iName,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredLangcodeAndNameAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredLangcodeAndPassAndMailHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredLangcode := self.Args("preferred_langcode").String()
	iPass := self.Args("pass").String()
	iMail := self.Args("mail").String()

	if helper.IsHas(iPreferredLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredLangcodeAndPassAndMail(offset, limit, iPreferredLangcode,iPass,iMail)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredLangcodeAndPassAndMail's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredLangcodeAndPassAndTimezoneHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredLangcode := self.Args("preferred_langcode").String()
	iPass := self.Args("pass").String()
	iTimezone := self.Args("timezone").String()

	if helper.IsHas(iPreferredLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredLangcodeAndPassAndTimezone(offset, limit, iPreferredLangcode,iPass,iTimezone)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredLangcodeAndPassAndTimezone's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredLangcodeAndPassAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredLangcode := self.Args("preferred_langcode").String()
	iPass := self.Args("pass").String()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iPreferredLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredLangcodeAndPassAndStatus(offset, limit, iPreferredLangcode,iPass,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredLangcodeAndPassAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredLangcodeAndPassAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredLangcode := self.Args("preferred_langcode").String()
	iPass := self.Args("pass").String()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iPreferredLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredLangcodeAndPassAndCreated(offset, limit, iPreferredLangcode,iPass,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredLangcodeAndPassAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredLangcodeAndPassAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredLangcode := self.Args("preferred_langcode").String()
	iPass := self.Args("pass").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iPreferredLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredLangcodeAndPassAndChanged(offset, limit, iPreferredLangcode,iPass,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredLangcodeAndPassAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredLangcodeAndPassAndAccessHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredLangcode := self.Args("preferred_langcode").String()
	iPass := self.Args("pass").String()
	iAccess := self.Args("access").MustInt()

	if helper.IsHas(iPreferredLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredLangcodeAndPassAndAccess(offset, limit, iPreferredLangcode,iPass,iAccess)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredLangcodeAndPassAndAccess's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredLangcodeAndPassAndLoginHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredLangcode := self.Args("preferred_langcode").String()
	iPass := self.Args("pass").String()
	iLogin := self.Args("login").MustInt()

	if helper.IsHas(iPreferredLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredLangcodeAndPassAndLogin(offset, limit, iPreferredLangcode,iPass,iLogin)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredLangcodeAndPassAndLogin's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredLangcodeAndPassAndInitHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredLangcode := self.Args("preferred_langcode").String()
	iPass := self.Args("pass").String()
	iInit := self.Args("init").String()

	if helper.IsHas(iPreferredLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredLangcodeAndPassAndInit(offset, limit, iPreferredLangcode,iPass,iInit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredLangcodeAndPassAndInit's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredLangcodeAndPassAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredLangcode := self.Args("preferred_langcode").String()
	iPass := self.Args("pass").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iPreferredLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredLangcodeAndPassAndDefaultLangcode(offset, limit, iPreferredLangcode,iPass,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredLangcodeAndPassAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredLangcodeAndMailAndTimezoneHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredLangcode := self.Args("preferred_langcode").String()
	iMail := self.Args("mail").String()
	iTimezone := self.Args("timezone").String()

	if helper.IsHas(iPreferredLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredLangcodeAndMailAndTimezone(offset, limit, iPreferredLangcode,iMail,iTimezone)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredLangcodeAndMailAndTimezone's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredLangcodeAndMailAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredLangcode := self.Args("preferred_langcode").String()
	iMail := self.Args("mail").String()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iPreferredLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredLangcodeAndMailAndStatus(offset, limit, iPreferredLangcode,iMail,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredLangcodeAndMailAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredLangcodeAndMailAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredLangcode := self.Args("preferred_langcode").String()
	iMail := self.Args("mail").String()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iPreferredLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredLangcodeAndMailAndCreated(offset, limit, iPreferredLangcode,iMail,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredLangcodeAndMailAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredLangcodeAndMailAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredLangcode := self.Args("preferred_langcode").String()
	iMail := self.Args("mail").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iPreferredLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredLangcodeAndMailAndChanged(offset, limit, iPreferredLangcode,iMail,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredLangcodeAndMailAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredLangcodeAndMailAndAccessHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredLangcode := self.Args("preferred_langcode").String()
	iMail := self.Args("mail").String()
	iAccess := self.Args("access").MustInt()

	if helper.IsHas(iPreferredLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredLangcodeAndMailAndAccess(offset, limit, iPreferredLangcode,iMail,iAccess)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredLangcodeAndMailAndAccess's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredLangcodeAndMailAndLoginHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredLangcode := self.Args("preferred_langcode").String()
	iMail := self.Args("mail").String()
	iLogin := self.Args("login").MustInt()

	if helper.IsHas(iPreferredLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredLangcodeAndMailAndLogin(offset, limit, iPreferredLangcode,iMail,iLogin)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredLangcodeAndMailAndLogin's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredLangcodeAndMailAndInitHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredLangcode := self.Args("preferred_langcode").String()
	iMail := self.Args("mail").String()
	iInit := self.Args("init").String()

	if helper.IsHas(iPreferredLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredLangcodeAndMailAndInit(offset, limit, iPreferredLangcode,iMail,iInit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredLangcodeAndMailAndInit's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredLangcodeAndMailAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredLangcode := self.Args("preferred_langcode").String()
	iMail := self.Args("mail").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iPreferredLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredLangcodeAndMailAndDefaultLangcode(offset, limit, iPreferredLangcode,iMail,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredLangcodeAndMailAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredLangcodeAndTimezoneAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredLangcode := self.Args("preferred_langcode").String()
	iTimezone := self.Args("timezone").String()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iPreferredLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredLangcodeAndTimezoneAndStatus(offset, limit, iPreferredLangcode,iTimezone,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredLangcodeAndTimezoneAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredLangcodeAndTimezoneAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredLangcode := self.Args("preferred_langcode").String()
	iTimezone := self.Args("timezone").String()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iPreferredLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredLangcodeAndTimezoneAndCreated(offset, limit, iPreferredLangcode,iTimezone,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredLangcodeAndTimezoneAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredLangcodeAndTimezoneAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredLangcode := self.Args("preferred_langcode").String()
	iTimezone := self.Args("timezone").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iPreferredLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredLangcodeAndTimezoneAndChanged(offset, limit, iPreferredLangcode,iTimezone,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredLangcodeAndTimezoneAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredLangcodeAndTimezoneAndAccessHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredLangcode := self.Args("preferred_langcode").String()
	iTimezone := self.Args("timezone").String()
	iAccess := self.Args("access").MustInt()

	if helper.IsHas(iPreferredLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredLangcodeAndTimezoneAndAccess(offset, limit, iPreferredLangcode,iTimezone,iAccess)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredLangcodeAndTimezoneAndAccess's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredLangcodeAndTimezoneAndLoginHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredLangcode := self.Args("preferred_langcode").String()
	iTimezone := self.Args("timezone").String()
	iLogin := self.Args("login").MustInt()

	if helper.IsHas(iPreferredLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredLangcodeAndTimezoneAndLogin(offset, limit, iPreferredLangcode,iTimezone,iLogin)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredLangcodeAndTimezoneAndLogin's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredLangcodeAndTimezoneAndInitHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredLangcode := self.Args("preferred_langcode").String()
	iTimezone := self.Args("timezone").String()
	iInit := self.Args("init").String()

	if helper.IsHas(iPreferredLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredLangcodeAndTimezoneAndInit(offset, limit, iPreferredLangcode,iTimezone,iInit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredLangcodeAndTimezoneAndInit's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredLangcodeAndTimezoneAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredLangcode := self.Args("preferred_langcode").String()
	iTimezone := self.Args("timezone").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iPreferredLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredLangcodeAndTimezoneAndDefaultLangcode(offset, limit, iPreferredLangcode,iTimezone,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredLangcodeAndTimezoneAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredLangcodeAndStatusAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredLangcode := self.Args("preferred_langcode").String()
	iStatus := self.Args("status").MustInt()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iPreferredLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredLangcodeAndStatusAndCreated(offset, limit, iPreferredLangcode,iStatus,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredLangcodeAndStatusAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredLangcodeAndStatusAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredLangcode := self.Args("preferred_langcode").String()
	iStatus := self.Args("status").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iPreferredLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredLangcodeAndStatusAndChanged(offset, limit, iPreferredLangcode,iStatus,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredLangcodeAndStatusAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredLangcodeAndStatusAndAccessHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredLangcode := self.Args("preferred_langcode").String()
	iStatus := self.Args("status").MustInt()
	iAccess := self.Args("access").MustInt()

	if helper.IsHas(iPreferredLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredLangcodeAndStatusAndAccess(offset, limit, iPreferredLangcode,iStatus,iAccess)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredLangcodeAndStatusAndAccess's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredLangcodeAndStatusAndLoginHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredLangcode := self.Args("preferred_langcode").String()
	iStatus := self.Args("status").MustInt()
	iLogin := self.Args("login").MustInt()

	if helper.IsHas(iPreferredLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredLangcodeAndStatusAndLogin(offset, limit, iPreferredLangcode,iStatus,iLogin)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredLangcodeAndStatusAndLogin's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredLangcodeAndStatusAndInitHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredLangcode := self.Args("preferred_langcode").String()
	iStatus := self.Args("status").MustInt()
	iInit := self.Args("init").String()

	if helper.IsHas(iPreferredLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredLangcodeAndStatusAndInit(offset, limit, iPreferredLangcode,iStatus,iInit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredLangcodeAndStatusAndInit's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredLangcodeAndStatusAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredLangcode := self.Args("preferred_langcode").String()
	iStatus := self.Args("status").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iPreferredLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredLangcodeAndStatusAndDefaultLangcode(offset, limit, iPreferredLangcode,iStatus,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredLangcodeAndStatusAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredLangcodeAndCreatedAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredLangcode := self.Args("preferred_langcode").String()
	iCreated := self.Args("created").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iPreferredLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredLangcodeAndCreatedAndChanged(offset, limit, iPreferredLangcode,iCreated,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredLangcodeAndCreatedAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredLangcodeAndCreatedAndAccessHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredLangcode := self.Args("preferred_langcode").String()
	iCreated := self.Args("created").MustInt()
	iAccess := self.Args("access").MustInt()

	if helper.IsHas(iPreferredLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredLangcodeAndCreatedAndAccess(offset, limit, iPreferredLangcode,iCreated,iAccess)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredLangcodeAndCreatedAndAccess's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredLangcodeAndCreatedAndLoginHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredLangcode := self.Args("preferred_langcode").String()
	iCreated := self.Args("created").MustInt()
	iLogin := self.Args("login").MustInt()

	if helper.IsHas(iPreferredLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredLangcodeAndCreatedAndLogin(offset, limit, iPreferredLangcode,iCreated,iLogin)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredLangcodeAndCreatedAndLogin's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredLangcodeAndCreatedAndInitHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredLangcode := self.Args("preferred_langcode").String()
	iCreated := self.Args("created").MustInt()
	iInit := self.Args("init").String()

	if helper.IsHas(iPreferredLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredLangcodeAndCreatedAndInit(offset, limit, iPreferredLangcode,iCreated,iInit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredLangcodeAndCreatedAndInit's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredLangcodeAndCreatedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredLangcode := self.Args("preferred_langcode").String()
	iCreated := self.Args("created").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iPreferredLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredLangcodeAndCreatedAndDefaultLangcode(offset, limit, iPreferredLangcode,iCreated,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredLangcodeAndCreatedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredLangcodeAndChangedAndAccessHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredLangcode := self.Args("preferred_langcode").String()
	iChanged := self.Args("changed").MustInt()
	iAccess := self.Args("access").MustInt()

	if helper.IsHas(iPreferredLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredLangcodeAndChangedAndAccess(offset, limit, iPreferredLangcode,iChanged,iAccess)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredLangcodeAndChangedAndAccess's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredLangcodeAndChangedAndLoginHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredLangcode := self.Args("preferred_langcode").String()
	iChanged := self.Args("changed").MustInt()
	iLogin := self.Args("login").MustInt()

	if helper.IsHas(iPreferredLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredLangcodeAndChangedAndLogin(offset, limit, iPreferredLangcode,iChanged,iLogin)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredLangcodeAndChangedAndLogin's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredLangcodeAndChangedAndInitHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredLangcode := self.Args("preferred_langcode").String()
	iChanged := self.Args("changed").MustInt()
	iInit := self.Args("init").String()

	if helper.IsHas(iPreferredLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredLangcodeAndChangedAndInit(offset, limit, iPreferredLangcode,iChanged,iInit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredLangcodeAndChangedAndInit's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredLangcodeAndChangedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredLangcode := self.Args("preferred_langcode").String()
	iChanged := self.Args("changed").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iPreferredLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredLangcodeAndChangedAndDefaultLangcode(offset, limit, iPreferredLangcode,iChanged,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredLangcodeAndChangedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredLangcodeAndAccessAndLoginHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredLangcode := self.Args("preferred_langcode").String()
	iAccess := self.Args("access").MustInt()
	iLogin := self.Args("login").MustInt()

	if helper.IsHas(iPreferredLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredLangcodeAndAccessAndLogin(offset, limit, iPreferredLangcode,iAccess,iLogin)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredLangcodeAndAccessAndLogin's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredLangcodeAndAccessAndInitHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredLangcode := self.Args("preferred_langcode").String()
	iAccess := self.Args("access").MustInt()
	iInit := self.Args("init").String()

	if helper.IsHas(iPreferredLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredLangcodeAndAccessAndInit(offset, limit, iPreferredLangcode,iAccess,iInit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredLangcodeAndAccessAndInit's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredLangcodeAndAccessAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredLangcode := self.Args("preferred_langcode").String()
	iAccess := self.Args("access").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iPreferredLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredLangcodeAndAccessAndDefaultLangcode(offset, limit, iPreferredLangcode,iAccess,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredLangcodeAndAccessAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredLangcodeAndLoginAndInitHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredLangcode := self.Args("preferred_langcode").String()
	iLogin := self.Args("login").MustInt()
	iInit := self.Args("init").String()

	if helper.IsHas(iPreferredLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredLangcodeAndLoginAndInit(offset, limit, iPreferredLangcode,iLogin,iInit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredLangcodeAndLoginAndInit's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredLangcodeAndLoginAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredLangcode := self.Args("preferred_langcode").String()
	iLogin := self.Args("login").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iPreferredLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredLangcodeAndLoginAndDefaultLangcode(offset, limit, iPreferredLangcode,iLogin,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredLangcodeAndLoginAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredLangcodeAndInitAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredLangcode := self.Args("preferred_langcode").String()
	iInit := self.Args("init").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iPreferredLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredLangcodeAndInitAndDefaultLangcode(offset, limit, iPreferredLangcode,iInit,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredLangcodeAndInitAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredAdminLangcodeAndNameAndPassHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredAdminLangcode := self.Args("preferred_admin_langcode").String()
	iName := self.Args("name").String()
	iPass := self.Args("pass").String()

	if helper.IsHas(iPreferredAdminLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredAdminLangcodeAndNameAndPass(offset, limit, iPreferredAdminLangcode,iName,iPass)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredAdminLangcodeAndNameAndPass's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredAdminLangcodeAndNameAndMailHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredAdminLangcode := self.Args("preferred_admin_langcode").String()
	iName := self.Args("name").String()
	iMail := self.Args("mail").String()

	if helper.IsHas(iPreferredAdminLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredAdminLangcodeAndNameAndMail(offset, limit, iPreferredAdminLangcode,iName,iMail)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredAdminLangcodeAndNameAndMail's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredAdminLangcodeAndNameAndTimezoneHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredAdminLangcode := self.Args("preferred_admin_langcode").String()
	iName := self.Args("name").String()
	iTimezone := self.Args("timezone").String()

	if helper.IsHas(iPreferredAdminLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredAdminLangcodeAndNameAndTimezone(offset, limit, iPreferredAdminLangcode,iName,iTimezone)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredAdminLangcodeAndNameAndTimezone's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredAdminLangcodeAndNameAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredAdminLangcode := self.Args("preferred_admin_langcode").String()
	iName := self.Args("name").String()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iPreferredAdminLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredAdminLangcodeAndNameAndStatus(offset, limit, iPreferredAdminLangcode,iName,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredAdminLangcodeAndNameAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredAdminLangcodeAndNameAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredAdminLangcode := self.Args("preferred_admin_langcode").String()
	iName := self.Args("name").String()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iPreferredAdminLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredAdminLangcodeAndNameAndCreated(offset, limit, iPreferredAdminLangcode,iName,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredAdminLangcodeAndNameAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredAdminLangcodeAndNameAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredAdminLangcode := self.Args("preferred_admin_langcode").String()
	iName := self.Args("name").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iPreferredAdminLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredAdminLangcodeAndNameAndChanged(offset, limit, iPreferredAdminLangcode,iName,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredAdminLangcodeAndNameAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredAdminLangcodeAndNameAndAccessHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredAdminLangcode := self.Args("preferred_admin_langcode").String()
	iName := self.Args("name").String()
	iAccess := self.Args("access").MustInt()

	if helper.IsHas(iPreferredAdminLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredAdminLangcodeAndNameAndAccess(offset, limit, iPreferredAdminLangcode,iName,iAccess)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredAdminLangcodeAndNameAndAccess's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredAdminLangcodeAndNameAndLoginHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredAdminLangcode := self.Args("preferred_admin_langcode").String()
	iName := self.Args("name").String()
	iLogin := self.Args("login").MustInt()

	if helper.IsHas(iPreferredAdminLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredAdminLangcodeAndNameAndLogin(offset, limit, iPreferredAdminLangcode,iName,iLogin)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredAdminLangcodeAndNameAndLogin's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredAdminLangcodeAndNameAndInitHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredAdminLangcode := self.Args("preferred_admin_langcode").String()
	iName := self.Args("name").String()
	iInit := self.Args("init").String()

	if helper.IsHas(iPreferredAdminLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredAdminLangcodeAndNameAndInit(offset, limit, iPreferredAdminLangcode,iName,iInit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredAdminLangcodeAndNameAndInit's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredAdminLangcodeAndNameAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredAdminLangcode := self.Args("preferred_admin_langcode").String()
	iName := self.Args("name").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iPreferredAdminLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredAdminLangcodeAndNameAndDefaultLangcode(offset, limit, iPreferredAdminLangcode,iName,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredAdminLangcodeAndNameAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredAdminLangcodeAndPassAndMailHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredAdminLangcode := self.Args("preferred_admin_langcode").String()
	iPass := self.Args("pass").String()
	iMail := self.Args("mail").String()

	if helper.IsHas(iPreferredAdminLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredAdminLangcodeAndPassAndMail(offset, limit, iPreferredAdminLangcode,iPass,iMail)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredAdminLangcodeAndPassAndMail's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredAdminLangcodeAndPassAndTimezoneHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredAdminLangcode := self.Args("preferred_admin_langcode").String()
	iPass := self.Args("pass").String()
	iTimezone := self.Args("timezone").String()

	if helper.IsHas(iPreferredAdminLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredAdminLangcodeAndPassAndTimezone(offset, limit, iPreferredAdminLangcode,iPass,iTimezone)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredAdminLangcodeAndPassAndTimezone's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredAdminLangcodeAndPassAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredAdminLangcode := self.Args("preferred_admin_langcode").String()
	iPass := self.Args("pass").String()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iPreferredAdminLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredAdminLangcodeAndPassAndStatus(offset, limit, iPreferredAdminLangcode,iPass,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredAdminLangcodeAndPassAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredAdminLangcodeAndPassAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredAdminLangcode := self.Args("preferred_admin_langcode").String()
	iPass := self.Args("pass").String()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iPreferredAdminLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredAdminLangcodeAndPassAndCreated(offset, limit, iPreferredAdminLangcode,iPass,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredAdminLangcodeAndPassAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredAdminLangcodeAndPassAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredAdminLangcode := self.Args("preferred_admin_langcode").String()
	iPass := self.Args("pass").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iPreferredAdminLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredAdminLangcodeAndPassAndChanged(offset, limit, iPreferredAdminLangcode,iPass,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredAdminLangcodeAndPassAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredAdminLangcodeAndPassAndAccessHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredAdminLangcode := self.Args("preferred_admin_langcode").String()
	iPass := self.Args("pass").String()
	iAccess := self.Args("access").MustInt()

	if helper.IsHas(iPreferredAdminLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredAdminLangcodeAndPassAndAccess(offset, limit, iPreferredAdminLangcode,iPass,iAccess)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredAdminLangcodeAndPassAndAccess's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredAdminLangcodeAndPassAndLoginHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredAdminLangcode := self.Args("preferred_admin_langcode").String()
	iPass := self.Args("pass").String()
	iLogin := self.Args("login").MustInt()

	if helper.IsHas(iPreferredAdminLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredAdminLangcodeAndPassAndLogin(offset, limit, iPreferredAdminLangcode,iPass,iLogin)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredAdminLangcodeAndPassAndLogin's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredAdminLangcodeAndPassAndInitHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredAdminLangcode := self.Args("preferred_admin_langcode").String()
	iPass := self.Args("pass").String()
	iInit := self.Args("init").String()

	if helper.IsHas(iPreferredAdminLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredAdminLangcodeAndPassAndInit(offset, limit, iPreferredAdminLangcode,iPass,iInit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredAdminLangcodeAndPassAndInit's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredAdminLangcodeAndPassAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredAdminLangcode := self.Args("preferred_admin_langcode").String()
	iPass := self.Args("pass").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iPreferredAdminLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredAdminLangcodeAndPassAndDefaultLangcode(offset, limit, iPreferredAdminLangcode,iPass,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredAdminLangcodeAndPassAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredAdminLangcodeAndMailAndTimezoneHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredAdminLangcode := self.Args("preferred_admin_langcode").String()
	iMail := self.Args("mail").String()
	iTimezone := self.Args("timezone").String()

	if helper.IsHas(iPreferredAdminLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredAdminLangcodeAndMailAndTimezone(offset, limit, iPreferredAdminLangcode,iMail,iTimezone)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredAdminLangcodeAndMailAndTimezone's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredAdminLangcodeAndMailAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredAdminLangcode := self.Args("preferred_admin_langcode").String()
	iMail := self.Args("mail").String()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iPreferredAdminLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredAdminLangcodeAndMailAndStatus(offset, limit, iPreferredAdminLangcode,iMail,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredAdminLangcodeAndMailAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredAdminLangcodeAndMailAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredAdminLangcode := self.Args("preferred_admin_langcode").String()
	iMail := self.Args("mail").String()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iPreferredAdminLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredAdminLangcodeAndMailAndCreated(offset, limit, iPreferredAdminLangcode,iMail,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredAdminLangcodeAndMailAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredAdminLangcodeAndMailAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredAdminLangcode := self.Args("preferred_admin_langcode").String()
	iMail := self.Args("mail").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iPreferredAdminLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredAdminLangcodeAndMailAndChanged(offset, limit, iPreferredAdminLangcode,iMail,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredAdminLangcodeAndMailAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredAdminLangcodeAndMailAndAccessHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredAdminLangcode := self.Args("preferred_admin_langcode").String()
	iMail := self.Args("mail").String()
	iAccess := self.Args("access").MustInt()

	if helper.IsHas(iPreferredAdminLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredAdminLangcodeAndMailAndAccess(offset, limit, iPreferredAdminLangcode,iMail,iAccess)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredAdminLangcodeAndMailAndAccess's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredAdminLangcodeAndMailAndLoginHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredAdminLangcode := self.Args("preferred_admin_langcode").String()
	iMail := self.Args("mail").String()
	iLogin := self.Args("login").MustInt()

	if helper.IsHas(iPreferredAdminLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredAdminLangcodeAndMailAndLogin(offset, limit, iPreferredAdminLangcode,iMail,iLogin)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredAdminLangcodeAndMailAndLogin's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredAdminLangcodeAndMailAndInitHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredAdminLangcode := self.Args("preferred_admin_langcode").String()
	iMail := self.Args("mail").String()
	iInit := self.Args("init").String()

	if helper.IsHas(iPreferredAdminLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredAdminLangcodeAndMailAndInit(offset, limit, iPreferredAdminLangcode,iMail,iInit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredAdminLangcodeAndMailAndInit's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredAdminLangcodeAndMailAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredAdminLangcode := self.Args("preferred_admin_langcode").String()
	iMail := self.Args("mail").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iPreferredAdminLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredAdminLangcodeAndMailAndDefaultLangcode(offset, limit, iPreferredAdminLangcode,iMail,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredAdminLangcodeAndMailAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredAdminLangcodeAndTimezoneAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredAdminLangcode := self.Args("preferred_admin_langcode").String()
	iTimezone := self.Args("timezone").String()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iPreferredAdminLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredAdminLangcodeAndTimezoneAndStatus(offset, limit, iPreferredAdminLangcode,iTimezone,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredAdminLangcodeAndTimezoneAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredAdminLangcodeAndTimezoneAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredAdminLangcode := self.Args("preferred_admin_langcode").String()
	iTimezone := self.Args("timezone").String()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iPreferredAdminLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredAdminLangcodeAndTimezoneAndCreated(offset, limit, iPreferredAdminLangcode,iTimezone,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredAdminLangcodeAndTimezoneAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredAdminLangcodeAndTimezoneAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredAdminLangcode := self.Args("preferred_admin_langcode").String()
	iTimezone := self.Args("timezone").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iPreferredAdminLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredAdminLangcodeAndTimezoneAndChanged(offset, limit, iPreferredAdminLangcode,iTimezone,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredAdminLangcodeAndTimezoneAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredAdminLangcodeAndTimezoneAndAccessHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredAdminLangcode := self.Args("preferred_admin_langcode").String()
	iTimezone := self.Args("timezone").String()
	iAccess := self.Args("access").MustInt()

	if helper.IsHas(iPreferredAdminLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredAdminLangcodeAndTimezoneAndAccess(offset, limit, iPreferredAdminLangcode,iTimezone,iAccess)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredAdminLangcodeAndTimezoneAndAccess's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredAdminLangcodeAndTimezoneAndLoginHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredAdminLangcode := self.Args("preferred_admin_langcode").String()
	iTimezone := self.Args("timezone").String()
	iLogin := self.Args("login").MustInt()

	if helper.IsHas(iPreferredAdminLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredAdminLangcodeAndTimezoneAndLogin(offset, limit, iPreferredAdminLangcode,iTimezone,iLogin)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredAdminLangcodeAndTimezoneAndLogin's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredAdminLangcodeAndTimezoneAndInitHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredAdminLangcode := self.Args("preferred_admin_langcode").String()
	iTimezone := self.Args("timezone").String()
	iInit := self.Args("init").String()

	if helper.IsHas(iPreferredAdminLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredAdminLangcodeAndTimezoneAndInit(offset, limit, iPreferredAdminLangcode,iTimezone,iInit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredAdminLangcodeAndTimezoneAndInit's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredAdminLangcodeAndTimezoneAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredAdminLangcode := self.Args("preferred_admin_langcode").String()
	iTimezone := self.Args("timezone").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iPreferredAdminLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredAdminLangcodeAndTimezoneAndDefaultLangcode(offset, limit, iPreferredAdminLangcode,iTimezone,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredAdminLangcodeAndTimezoneAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredAdminLangcodeAndStatusAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredAdminLangcode := self.Args("preferred_admin_langcode").String()
	iStatus := self.Args("status").MustInt()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iPreferredAdminLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredAdminLangcodeAndStatusAndCreated(offset, limit, iPreferredAdminLangcode,iStatus,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredAdminLangcodeAndStatusAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredAdminLangcodeAndStatusAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredAdminLangcode := self.Args("preferred_admin_langcode").String()
	iStatus := self.Args("status").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iPreferredAdminLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredAdminLangcodeAndStatusAndChanged(offset, limit, iPreferredAdminLangcode,iStatus,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredAdminLangcodeAndStatusAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredAdminLangcodeAndStatusAndAccessHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredAdminLangcode := self.Args("preferred_admin_langcode").String()
	iStatus := self.Args("status").MustInt()
	iAccess := self.Args("access").MustInt()

	if helper.IsHas(iPreferredAdminLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredAdminLangcodeAndStatusAndAccess(offset, limit, iPreferredAdminLangcode,iStatus,iAccess)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredAdminLangcodeAndStatusAndAccess's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredAdminLangcodeAndStatusAndLoginHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredAdminLangcode := self.Args("preferred_admin_langcode").String()
	iStatus := self.Args("status").MustInt()
	iLogin := self.Args("login").MustInt()

	if helper.IsHas(iPreferredAdminLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredAdminLangcodeAndStatusAndLogin(offset, limit, iPreferredAdminLangcode,iStatus,iLogin)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredAdminLangcodeAndStatusAndLogin's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredAdminLangcodeAndStatusAndInitHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredAdminLangcode := self.Args("preferred_admin_langcode").String()
	iStatus := self.Args("status").MustInt()
	iInit := self.Args("init").String()

	if helper.IsHas(iPreferredAdminLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredAdminLangcodeAndStatusAndInit(offset, limit, iPreferredAdminLangcode,iStatus,iInit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredAdminLangcodeAndStatusAndInit's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredAdminLangcodeAndStatusAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredAdminLangcode := self.Args("preferred_admin_langcode").String()
	iStatus := self.Args("status").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iPreferredAdminLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredAdminLangcodeAndStatusAndDefaultLangcode(offset, limit, iPreferredAdminLangcode,iStatus,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredAdminLangcodeAndStatusAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredAdminLangcodeAndCreatedAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredAdminLangcode := self.Args("preferred_admin_langcode").String()
	iCreated := self.Args("created").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iPreferredAdminLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredAdminLangcodeAndCreatedAndChanged(offset, limit, iPreferredAdminLangcode,iCreated,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredAdminLangcodeAndCreatedAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredAdminLangcodeAndCreatedAndAccessHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredAdminLangcode := self.Args("preferred_admin_langcode").String()
	iCreated := self.Args("created").MustInt()
	iAccess := self.Args("access").MustInt()

	if helper.IsHas(iPreferredAdminLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredAdminLangcodeAndCreatedAndAccess(offset, limit, iPreferredAdminLangcode,iCreated,iAccess)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredAdminLangcodeAndCreatedAndAccess's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredAdminLangcodeAndCreatedAndLoginHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredAdminLangcode := self.Args("preferred_admin_langcode").String()
	iCreated := self.Args("created").MustInt()
	iLogin := self.Args("login").MustInt()

	if helper.IsHas(iPreferredAdminLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredAdminLangcodeAndCreatedAndLogin(offset, limit, iPreferredAdminLangcode,iCreated,iLogin)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredAdminLangcodeAndCreatedAndLogin's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredAdminLangcodeAndCreatedAndInitHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredAdminLangcode := self.Args("preferred_admin_langcode").String()
	iCreated := self.Args("created").MustInt()
	iInit := self.Args("init").String()

	if helper.IsHas(iPreferredAdminLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredAdminLangcodeAndCreatedAndInit(offset, limit, iPreferredAdminLangcode,iCreated,iInit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredAdminLangcodeAndCreatedAndInit's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredAdminLangcodeAndCreatedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredAdminLangcode := self.Args("preferred_admin_langcode").String()
	iCreated := self.Args("created").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iPreferredAdminLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredAdminLangcodeAndCreatedAndDefaultLangcode(offset, limit, iPreferredAdminLangcode,iCreated,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredAdminLangcodeAndCreatedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredAdminLangcodeAndChangedAndAccessHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredAdminLangcode := self.Args("preferred_admin_langcode").String()
	iChanged := self.Args("changed").MustInt()
	iAccess := self.Args("access").MustInt()

	if helper.IsHas(iPreferredAdminLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredAdminLangcodeAndChangedAndAccess(offset, limit, iPreferredAdminLangcode,iChanged,iAccess)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredAdminLangcodeAndChangedAndAccess's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredAdminLangcodeAndChangedAndLoginHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredAdminLangcode := self.Args("preferred_admin_langcode").String()
	iChanged := self.Args("changed").MustInt()
	iLogin := self.Args("login").MustInt()

	if helper.IsHas(iPreferredAdminLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredAdminLangcodeAndChangedAndLogin(offset, limit, iPreferredAdminLangcode,iChanged,iLogin)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredAdminLangcodeAndChangedAndLogin's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredAdminLangcodeAndChangedAndInitHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredAdminLangcode := self.Args("preferred_admin_langcode").String()
	iChanged := self.Args("changed").MustInt()
	iInit := self.Args("init").String()

	if helper.IsHas(iPreferredAdminLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredAdminLangcodeAndChangedAndInit(offset, limit, iPreferredAdminLangcode,iChanged,iInit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredAdminLangcodeAndChangedAndInit's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredAdminLangcodeAndChangedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredAdminLangcode := self.Args("preferred_admin_langcode").String()
	iChanged := self.Args("changed").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iPreferredAdminLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredAdminLangcodeAndChangedAndDefaultLangcode(offset, limit, iPreferredAdminLangcode,iChanged,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredAdminLangcodeAndChangedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredAdminLangcodeAndAccessAndLoginHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredAdminLangcode := self.Args("preferred_admin_langcode").String()
	iAccess := self.Args("access").MustInt()
	iLogin := self.Args("login").MustInt()

	if helper.IsHas(iPreferredAdminLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredAdminLangcodeAndAccessAndLogin(offset, limit, iPreferredAdminLangcode,iAccess,iLogin)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredAdminLangcodeAndAccessAndLogin's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredAdminLangcodeAndAccessAndInitHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredAdminLangcode := self.Args("preferred_admin_langcode").String()
	iAccess := self.Args("access").MustInt()
	iInit := self.Args("init").String()

	if helper.IsHas(iPreferredAdminLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredAdminLangcodeAndAccessAndInit(offset, limit, iPreferredAdminLangcode,iAccess,iInit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredAdminLangcodeAndAccessAndInit's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredAdminLangcodeAndAccessAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredAdminLangcode := self.Args("preferred_admin_langcode").String()
	iAccess := self.Args("access").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iPreferredAdminLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredAdminLangcodeAndAccessAndDefaultLangcode(offset, limit, iPreferredAdminLangcode,iAccess,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredAdminLangcodeAndAccessAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredAdminLangcodeAndLoginAndInitHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredAdminLangcode := self.Args("preferred_admin_langcode").String()
	iLogin := self.Args("login").MustInt()
	iInit := self.Args("init").String()

	if helper.IsHas(iPreferredAdminLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredAdminLangcodeAndLoginAndInit(offset, limit, iPreferredAdminLangcode,iLogin,iInit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredAdminLangcodeAndLoginAndInit's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredAdminLangcodeAndLoginAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredAdminLangcode := self.Args("preferred_admin_langcode").String()
	iLogin := self.Args("login").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iPreferredAdminLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredAdminLangcodeAndLoginAndDefaultLangcode(offset, limit, iPreferredAdminLangcode,iLogin,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredAdminLangcodeAndLoginAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredAdminLangcodeAndInitAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredAdminLangcode := self.Args("preferred_admin_langcode").String()
	iInit := self.Args("init").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iPreferredAdminLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredAdminLangcodeAndInitAndDefaultLangcode(offset, limit, iPreferredAdminLangcode,iInit,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredAdminLangcodeAndInitAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByNameAndPassAndMailHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iPass := self.Args("pass").String()
	iMail := self.Args("mail").String()

	if helper.IsHas(iName) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByNameAndPassAndMail(offset, limit, iName,iPass,iMail)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByNameAndPassAndMail's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByNameAndPassAndTimezoneHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iPass := self.Args("pass").String()
	iTimezone := self.Args("timezone").String()

	if helper.IsHas(iName) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByNameAndPassAndTimezone(offset, limit, iName,iPass,iTimezone)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByNameAndPassAndTimezone's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByNameAndPassAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iPass := self.Args("pass").String()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iName) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByNameAndPassAndStatus(offset, limit, iName,iPass,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByNameAndPassAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByNameAndPassAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iPass := self.Args("pass").String()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iName) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByNameAndPassAndCreated(offset, limit, iName,iPass,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByNameAndPassAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByNameAndPassAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iPass := self.Args("pass").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iName) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByNameAndPassAndChanged(offset, limit, iName,iPass,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByNameAndPassAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByNameAndPassAndAccessHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iPass := self.Args("pass").String()
	iAccess := self.Args("access").MustInt()

	if helper.IsHas(iName) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByNameAndPassAndAccess(offset, limit, iName,iPass,iAccess)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByNameAndPassAndAccess's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByNameAndPassAndLoginHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iPass := self.Args("pass").String()
	iLogin := self.Args("login").MustInt()

	if helper.IsHas(iName) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByNameAndPassAndLogin(offset, limit, iName,iPass,iLogin)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByNameAndPassAndLogin's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByNameAndPassAndInitHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iPass := self.Args("pass").String()
	iInit := self.Args("init").String()

	if helper.IsHas(iName) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByNameAndPassAndInit(offset, limit, iName,iPass,iInit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByNameAndPassAndInit's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByNameAndPassAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iPass := self.Args("pass").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iName) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByNameAndPassAndDefaultLangcode(offset, limit, iName,iPass,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByNameAndPassAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByNameAndMailAndTimezoneHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iMail := self.Args("mail").String()
	iTimezone := self.Args("timezone").String()

	if helper.IsHas(iName) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByNameAndMailAndTimezone(offset, limit, iName,iMail,iTimezone)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByNameAndMailAndTimezone's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByNameAndMailAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iMail := self.Args("mail").String()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iName) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByNameAndMailAndStatus(offset, limit, iName,iMail,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByNameAndMailAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByNameAndMailAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iMail := self.Args("mail").String()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iName) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByNameAndMailAndCreated(offset, limit, iName,iMail,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByNameAndMailAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByNameAndMailAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iMail := self.Args("mail").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iName) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByNameAndMailAndChanged(offset, limit, iName,iMail,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByNameAndMailAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByNameAndMailAndAccessHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iMail := self.Args("mail").String()
	iAccess := self.Args("access").MustInt()

	if helper.IsHas(iName) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByNameAndMailAndAccess(offset, limit, iName,iMail,iAccess)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByNameAndMailAndAccess's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByNameAndMailAndLoginHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iMail := self.Args("mail").String()
	iLogin := self.Args("login").MustInt()

	if helper.IsHas(iName) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByNameAndMailAndLogin(offset, limit, iName,iMail,iLogin)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByNameAndMailAndLogin's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByNameAndMailAndInitHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iMail := self.Args("mail").String()
	iInit := self.Args("init").String()

	if helper.IsHas(iName) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByNameAndMailAndInit(offset, limit, iName,iMail,iInit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByNameAndMailAndInit's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByNameAndMailAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iMail := self.Args("mail").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iName) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByNameAndMailAndDefaultLangcode(offset, limit, iName,iMail,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByNameAndMailAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByNameAndTimezoneAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iTimezone := self.Args("timezone").String()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iName) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByNameAndTimezoneAndStatus(offset, limit, iName,iTimezone,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByNameAndTimezoneAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByNameAndTimezoneAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iTimezone := self.Args("timezone").String()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iName) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByNameAndTimezoneAndCreated(offset, limit, iName,iTimezone,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByNameAndTimezoneAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByNameAndTimezoneAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iTimezone := self.Args("timezone").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iName) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByNameAndTimezoneAndChanged(offset, limit, iName,iTimezone,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByNameAndTimezoneAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByNameAndTimezoneAndAccessHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iTimezone := self.Args("timezone").String()
	iAccess := self.Args("access").MustInt()

	if helper.IsHas(iName) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByNameAndTimezoneAndAccess(offset, limit, iName,iTimezone,iAccess)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByNameAndTimezoneAndAccess's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByNameAndTimezoneAndLoginHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iTimezone := self.Args("timezone").String()
	iLogin := self.Args("login").MustInt()

	if helper.IsHas(iName) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByNameAndTimezoneAndLogin(offset, limit, iName,iTimezone,iLogin)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByNameAndTimezoneAndLogin's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByNameAndTimezoneAndInitHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iTimezone := self.Args("timezone").String()
	iInit := self.Args("init").String()

	if helper.IsHas(iName) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByNameAndTimezoneAndInit(offset, limit, iName,iTimezone,iInit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByNameAndTimezoneAndInit's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByNameAndTimezoneAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iTimezone := self.Args("timezone").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iName) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByNameAndTimezoneAndDefaultLangcode(offset, limit, iName,iTimezone,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByNameAndTimezoneAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByNameAndStatusAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iStatus := self.Args("status").MustInt()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iName) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByNameAndStatusAndCreated(offset, limit, iName,iStatus,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByNameAndStatusAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByNameAndStatusAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iStatus := self.Args("status").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iName) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByNameAndStatusAndChanged(offset, limit, iName,iStatus,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByNameAndStatusAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByNameAndStatusAndAccessHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iStatus := self.Args("status").MustInt()
	iAccess := self.Args("access").MustInt()

	if helper.IsHas(iName) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByNameAndStatusAndAccess(offset, limit, iName,iStatus,iAccess)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByNameAndStatusAndAccess's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByNameAndStatusAndLoginHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iStatus := self.Args("status").MustInt()
	iLogin := self.Args("login").MustInt()

	if helper.IsHas(iName) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByNameAndStatusAndLogin(offset, limit, iName,iStatus,iLogin)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByNameAndStatusAndLogin's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByNameAndStatusAndInitHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iStatus := self.Args("status").MustInt()
	iInit := self.Args("init").String()

	if helper.IsHas(iName) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByNameAndStatusAndInit(offset, limit, iName,iStatus,iInit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByNameAndStatusAndInit's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByNameAndStatusAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iStatus := self.Args("status").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iName) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByNameAndStatusAndDefaultLangcode(offset, limit, iName,iStatus,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByNameAndStatusAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByNameAndCreatedAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iCreated := self.Args("created").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iName) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByNameAndCreatedAndChanged(offset, limit, iName,iCreated,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByNameAndCreatedAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByNameAndCreatedAndAccessHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iCreated := self.Args("created").MustInt()
	iAccess := self.Args("access").MustInt()

	if helper.IsHas(iName) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByNameAndCreatedAndAccess(offset, limit, iName,iCreated,iAccess)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByNameAndCreatedAndAccess's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByNameAndCreatedAndLoginHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iCreated := self.Args("created").MustInt()
	iLogin := self.Args("login").MustInt()

	if helper.IsHas(iName) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByNameAndCreatedAndLogin(offset, limit, iName,iCreated,iLogin)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByNameAndCreatedAndLogin's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByNameAndCreatedAndInitHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iCreated := self.Args("created").MustInt()
	iInit := self.Args("init").String()

	if helper.IsHas(iName) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByNameAndCreatedAndInit(offset, limit, iName,iCreated,iInit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByNameAndCreatedAndInit's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByNameAndCreatedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iCreated := self.Args("created").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iName) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByNameAndCreatedAndDefaultLangcode(offset, limit, iName,iCreated,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByNameAndCreatedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByNameAndChangedAndAccessHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iChanged := self.Args("changed").MustInt()
	iAccess := self.Args("access").MustInt()

	if helper.IsHas(iName) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByNameAndChangedAndAccess(offset, limit, iName,iChanged,iAccess)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByNameAndChangedAndAccess's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByNameAndChangedAndLoginHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iChanged := self.Args("changed").MustInt()
	iLogin := self.Args("login").MustInt()

	if helper.IsHas(iName) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByNameAndChangedAndLogin(offset, limit, iName,iChanged,iLogin)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByNameAndChangedAndLogin's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByNameAndChangedAndInitHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iChanged := self.Args("changed").MustInt()
	iInit := self.Args("init").String()

	if helper.IsHas(iName) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByNameAndChangedAndInit(offset, limit, iName,iChanged,iInit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByNameAndChangedAndInit's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByNameAndChangedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iChanged := self.Args("changed").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iName) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByNameAndChangedAndDefaultLangcode(offset, limit, iName,iChanged,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByNameAndChangedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByNameAndAccessAndLoginHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iAccess := self.Args("access").MustInt()
	iLogin := self.Args("login").MustInt()

	if helper.IsHas(iName) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByNameAndAccessAndLogin(offset, limit, iName,iAccess,iLogin)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByNameAndAccessAndLogin's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByNameAndAccessAndInitHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iAccess := self.Args("access").MustInt()
	iInit := self.Args("init").String()

	if helper.IsHas(iName) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByNameAndAccessAndInit(offset, limit, iName,iAccess,iInit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByNameAndAccessAndInit's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByNameAndAccessAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iAccess := self.Args("access").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iName) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByNameAndAccessAndDefaultLangcode(offset, limit, iName,iAccess,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByNameAndAccessAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByNameAndLoginAndInitHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iLogin := self.Args("login").MustInt()
	iInit := self.Args("init").String()

	if helper.IsHas(iName) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByNameAndLoginAndInit(offset, limit, iName,iLogin,iInit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByNameAndLoginAndInit's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByNameAndLoginAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iLogin := self.Args("login").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iName) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByNameAndLoginAndDefaultLangcode(offset, limit, iName,iLogin,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByNameAndLoginAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByNameAndInitAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iInit := self.Args("init").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iName) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByNameAndInitAndDefaultLangcode(offset, limit, iName,iInit,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByNameAndInitAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPassAndMailAndTimezoneHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPass := self.Args("pass").String()
	iMail := self.Args("mail").String()
	iTimezone := self.Args("timezone").String()

	if helper.IsHas(iPass) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPassAndMailAndTimezone(offset, limit, iPass,iMail,iTimezone)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPassAndMailAndTimezone's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPassAndMailAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPass := self.Args("pass").String()
	iMail := self.Args("mail").String()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iPass) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPassAndMailAndStatus(offset, limit, iPass,iMail,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPassAndMailAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPassAndMailAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPass := self.Args("pass").String()
	iMail := self.Args("mail").String()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iPass) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPassAndMailAndCreated(offset, limit, iPass,iMail,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPassAndMailAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPassAndMailAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPass := self.Args("pass").String()
	iMail := self.Args("mail").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iPass) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPassAndMailAndChanged(offset, limit, iPass,iMail,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPassAndMailAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPassAndMailAndAccessHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPass := self.Args("pass").String()
	iMail := self.Args("mail").String()
	iAccess := self.Args("access").MustInt()

	if helper.IsHas(iPass) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPassAndMailAndAccess(offset, limit, iPass,iMail,iAccess)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPassAndMailAndAccess's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPassAndMailAndLoginHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPass := self.Args("pass").String()
	iMail := self.Args("mail").String()
	iLogin := self.Args("login").MustInt()

	if helper.IsHas(iPass) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPassAndMailAndLogin(offset, limit, iPass,iMail,iLogin)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPassAndMailAndLogin's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPassAndMailAndInitHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPass := self.Args("pass").String()
	iMail := self.Args("mail").String()
	iInit := self.Args("init").String()

	if helper.IsHas(iPass) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPassAndMailAndInit(offset, limit, iPass,iMail,iInit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPassAndMailAndInit's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPassAndMailAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPass := self.Args("pass").String()
	iMail := self.Args("mail").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iPass) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPassAndMailAndDefaultLangcode(offset, limit, iPass,iMail,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPassAndMailAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPassAndTimezoneAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPass := self.Args("pass").String()
	iTimezone := self.Args("timezone").String()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iPass) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPassAndTimezoneAndStatus(offset, limit, iPass,iTimezone,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPassAndTimezoneAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPassAndTimezoneAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPass := self.Args("pass").String()
	iTimezone := self.Args("timezone").String()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iPass) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPassAndTimezoneAndCreated(offset, limit, iPass,iTimezone,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPassAndTimezoneAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPassAndTimezoneAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPass := self.Args("pass").String()
	iTimezone := self.Args("timezone").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iPass) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPassAndTimezoneAndChanged(offset, limit, iPass,iTimezone,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPassAndTimezoneAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPassAndTimezoneAndAccessHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPass := self.Args("pass").String()
	iTimezone := self.Args("timezone").String()
	iAccess := self.Args("access").MustInt()

	if helper.IsHas(iPass) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPassAndTimezoneAndAccess(offset, limit, iPass,iTimezone,iAccess)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPassAndTimezoneAndAccess's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPassAndTimezoneAndLoginHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPass := self.Args("pass").String()
	iTimezone := self.Args("timezone").String()
	iLogin := self.Args("login").MustInt()

	if helper.IsHas(iPass) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPassAndTimezoneAndLogin(offset, limit, iPass,iTimezone,iLogin)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPassAndTimezoneAndLogin's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPassAndTimezoneAndInitHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPass := self.Args("pass").String()
	iTimezone := self.Args("timezone").String()
	iInit := self.Args("init").String()

	if helper.IsHas(iPass) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPassAndTimezoneAndInit(offset, limit, iPass,iTimezone,iInit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPassAndTimezoneAndInit's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPassAndTimezoneAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPass := self.Args("pass").String()
	iTimezone := self.Args("timezone").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iPass) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPassAndTimezoneAndDefaultLangcode(offset, limit, iPass,iTimezone,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPassAndTimezoneAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPassAndStatusAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPass := self.Args("pass").String()
	iStatus := self.Args("status").MustInt()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iPass) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPassAndStatusAndCreated(offset, limit, iPass,iStatus,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPassAndStatusAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPassAndStatusAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPass := self.Args("pass").String()
	iStatus := self.Args("status").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iPass) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPassAndStatusAndChanged(offset, limit, iPass,iStatus,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPassAndStatusAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPassAndStatusAndAccessHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPass := self.Args("pass").String()
	iStatus := self.Args("status").MustInt()
	iAccess := self.Args("access").MustInt()

	if helper.IsHas(iPass) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPassAndStatusAndAccess(offset, limit, iPass,iStatus,iAccess)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPassAndStatusAndAccess's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPassAndStatusAndLoginHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPass := self.Args("pass").String()
	iStatus := self.Args("status").MustInt()
	iLogin := self.Args("login").MustInt()

	if helper.IsHas(iPass) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPassAndStatusAndLogin(offset, limit, iPass,iStatus,iLogin)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPassAndStatusAndLogin's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPassAndStatusAndInitHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPass := self.Args("pass").String()
	iStatus := self.Args("status").MustInt()
	iInit := self.Args("init").String()

	if helper.IsHas(iPass) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPassAndStatusAndInit(offset, limit, iPass,iStatus,iInit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPassAndStatusAndInit's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPassAndStatusAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPass := self.Args("pass").String()
	iStatus := self.Args("status").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iPass) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPassAndStatusAndDefaultLangcode(offset, limit, iPass,iStatus,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPassAndStatusAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPassAndCreatedAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPass := self.Args("pass").String()
	iCreated := self.Args("created").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iPass) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPassAndCreatedAndChanged(offset, limit, iPass,iCreated,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPassAndCreatedAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPassAndCreatedAndAccessHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPass := self.Args("pass").String()
	iCreated := self.Args("created").MustInt()
	iAccess := self.Args("access").MustInt()

	if helper.IsHas(iPass) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPassAndCreatedAndAccess(offset, limit, iPass,iCreated,iAccess)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPassAndCreatedAndAccess's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPassAndCreatedAndLoginHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPass := self.Args("pass").String()
	iCreated := self.Args("created").MustInt()
	iLogin := self.Args("login").MustInt()

	if helper.IsHas(iPass) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPassAndCreatedAndLogin(offset, limit, iPass,iCreated,iLogin)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPassAndCreatedAndLogin's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPassAndCreatedAndInitHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPass := self.Args("pass").String()
	iCreated := self.Args("created").MustInt()
	iInit := self.Args("init").String()

	if helper.IsHas(iPass) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPassAndCreatedAndInit(offset, limit, iPass,iCreated,iInit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPassAndCreatedAndInit's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPassAndCreatedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPass := self.Args("pass").String()
	iCreated := self.Args("created").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iPass) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPassAndCreatedAndDefaultLangcode(offset, limit, iPass,iCreated,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPassAndCreatedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPassAndChangedAndAccessHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPass := self.Args("pass").String()
	iChanged := self.Args("changed").MustInt()
	iAccess := self.Args("access").MustInt()

	if helper.IsHas(iPass) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPassAndChangedAndAccess(offset, limit, iPass,iChanged,iAccess)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPassAndChangedAndAccess's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPassAndChangedAndLoginHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPass := self.Args("pass").String()
	iChanged := self.Args("changed").MustInt()
	iLogin := self.Args("login").MustInt()

	if helper.IsHas(iPass) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPassAndChangedAndLogin(offset, limit, iPass,iChanged,iLogin)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPassAndChangedAndLogin's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPassAndChangedAndInitHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPass := self.Args("pass").String()
	iChanged := self.Args("changed").MustInt()
	iInit := self.Args("init").String()

	if helper.IsHas(iPass) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPassAndChangedAndInit(offset, limit, iPass,iChanged,iInit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPassAndChangedAndInit's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPassAndChangedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPass := self.Args("pass").String()
	iChanged := self.Args("changed").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iPass) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPassAndChangedAndDefaultLangcode(offset, limit, iPass,iChanged,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPassAndChangedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPassAndAccessAndLoginHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPass := self.Args("pass").String()
	iAccess := self.Args("access").MustInt()
	iLogin := self.Args("login").MustInt()

	if helper.IsHas(iPass) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPassAndAccessAndLogin(offset, limit, iPass,iAccess,iLogin)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPassAndAccessAndLogin's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPassAndAccessAndInitHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPass := self.Args("pass").String()
	iAccess := self.Args("access").MustInt()
	iInit := self.Args("init").String()

	if helper.IsHas(iPass) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPassAndAccessAndInit(offset, limit, iPass,iAccess,iInit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPassAndAccessAndInit's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPassAndAccessAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPass := self.Args("pass").String()
	iAccess := self.Args("access").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iPass) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPassAndAccessAndDefaultLangcode(offset, limit, iPass,iAccess,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPassAndAccessAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPassAndLoginAndInitHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPass := self.Args("pass").String()
	iLogin := self.Args("login").MustInt()
	iInit := self.Args("init").String()

	if helper.IsHas(iPass) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPassAndLoginAndInit(offset, limit, iPass,iLogin,iInit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPassAndLoginAndInit's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPassAndLoginAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPass := self.Args("pass").String()
	iLogin := self.Args("login").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iPass) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPassAndLoginAndDefaultLangcode(offset, limit, iPass,iLogin,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPassAndLoginAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPassAndInitAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPass := self.Args("pass").String()
	iInit := self.Args("init").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iPass) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPassAndInitAndDefaultLangcode(offset, limit, iPass,iInit,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPassAndInitAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByMailAndTimezoneAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMail := self.Args("mail").String()
	iTimezone := self.Args("timezone").String()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iMail) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByMailAndTimezoneAndStatus(offset, limit, iMail,iTimezone,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByMailAndTimezoneAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByMailAndTimezoneAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMail := self.Args("mail").String()
	iTimezone := self.Args("timezone").String()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iMail) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByMailAndTimezoneAndCreated(offset, limit, iMail,iTimezone,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByMailAndTimezoneAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByMailAndTimezoneAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMail := self.Args("mail").String()
	iTimezone := self.Args("timezone").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iMail) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByMailAndTimezoneAndChanged(offset, limit, iMail,iTimezone,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByMailAndTimezoneAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByMailAndTimezoneAndAccessHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMail := self.Args("mail").String()
	iTimezone := self.Args("timezone").String()
	iAccess := self.Args("access").MustInt()

	if helper.IsHas(iMail) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByMailAndTimezoneAndAccess(offset, limit, iMail,iTimezone,iAccess)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByMailAndTimezoneAndAccess's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByMailAndTimezoneAndLoginHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMail := self.Args("mail").String()
	iTimezone := self.Args("timezone").String()
	iLogin := self.Args("login").MustInt()

	if helper.IsHas(iMail) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByMailAndTimezoneAndLogin(offset, limit, iMail,iTimezone,iLogin)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByMailAndTimezoneAndLogin's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByMailAndTimezoneAndInitHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMail := self.Args("mail").String()
	iTimezone := self.Args("timezone").String()
	iInit := self.Args("init").String()

	if helper.IsHas(iMail) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByMailAndTimezoneAndInit(offset, limit, iMail,iTimezone,iInit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByMailAndTimezoneAndInit's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByMailAndTimezoneAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMail := self.Args("mail").String()
	iTimezone := self.Args("timezone").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iMail) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByMailAndTimezoneAndDefaultLangcode(offset, limit, iMail,iTimezone,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByMailAndTimezoneAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByMailAndStatusAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMail := self.Args("mail").String()
	iStatus := self.Args("status").MustInt()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iMail) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByMailAndStatusAndCreated(offset, limit, iMail,iStatus,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByMailAndStatusAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByMailAndStatusAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMail := self.Args("mail").String()
	iStatus := self.Args("status").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iMail) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByMailAndStatusAndChanged(offset, limit, iMail,iStatus,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByMailAndStatusAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByMailAndStatusAndAccessHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMail := self.Args("mail").String()
	iStatus := self.Args("status").MustInt()
	iAccess := self.Args("access").MustInt()

	if helper.IsHas(iMail) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByMailAndStatusAndAccess(offset, limit, iMail,iStatus,iAccess)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByMailAndStatusAndAccess's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByMailAndStatusAndLoginHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMail := self.Args("mail").String()
	iStatus := self.Args("status").MustInt()
	iLogin := self.Args("login").MustInt()

	if helper.IsHas(iMail) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByMailAndStatusAndLogin(offset, limit, iMail,iStatus,iLogin)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByMailAndStatusAndLogin's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByMailAndStatusAndInitHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMail := self.Args("mail").String()
	iStatus := self.Args("status").MustInt()
	iInit := self.Args("init").String()

	if helper.IsHas(iMail) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByMailAndStatusAndInit(offset, limit, iMail,iStatus,iInit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByMailAndStatusAndInit's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByMailAndStatusAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMail := self.Args("mail").String()
	iStatus := self.Args("status").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iMail) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByMailAndStatusAndDefaultLangcode(offset, limit, iMail,iStatus,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByMailAndStatusAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByMailAndCreatedAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMail := self.Args("mail").String()
	iCreated := self.Args("created").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iMail) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByMailAndCreatedAndChanged(offset, limit, iMail,iCreated,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByMailAndCreatedAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByMailAndCreatedAndAccessHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMail := self.Args("mail").String()
	iCreated := self.Args("created").MustInt()
	iAccess := self.Args("access").MustInt()

	if helper.IsHas(iMail) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByMailAndCreatedAndAccess(offset, limit, iMail,iCreated,iAccess)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByMailAndCreatedAndAccess's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByMailAndCreatedAndLoginHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMail := self.Args("mail").String()
	iCreated := self.Args("created").MustInt()
	iLogin := self.Args("login").MustInt()

	if helper.IsHas(iMail) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByMailAndCreatedAndLogin(offset, limit, iMail,iCreated,iLogin)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByMailAndCreatedAndLogin's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByMailAndCreatedAndInitHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMail := self.Args("mail").String()
	iCreated := self.Args("created").MustInt()
	iInit := self.Args("init").String()

	if helper.IsHas(iMail) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByMailAndCreatedAndInit(offset, limit, iMail,iCreated,iInit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByMailAndCreatedAndInit's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByMailAndCreatedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMail := self.Args("mail").String()
	iCreated := self.Args("created").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iMail) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByMailAndCreatedAndDefaultLangcode(offset, limit, iMail,iCreated,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByMailAndCreatedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByMailAndChangedAndAccessHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMail := self.Args("mail").String()
	iChanged := self.Args("changed").MustInt()
	iAccess := self.Args("access").MustInt()

	if helper.IsHas(iMail) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByMailAndChangedAndAccess(offset, limit, iMail,iChanged,iAccess)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByMailAndChangedAndAccess's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByMailAndChangedAndLoginHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMail := self.Args("mail").String()
	iChanged := self.Args("changed").MustInt()
	iLogin := self.Args("login").MustInt()

	if helper.IsHas(iMail) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByMailAndChangedAndLogin(offset, limit, iMail,iChanged,iLogin)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByMailAndChangedAndLogin's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByMailAndChangedAndInitHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMail := self.Args("mail").String()
	iChanged := self.Args("changed").MustInt()
	iInit := self.Args("init").String()

	if helper.IsHas(iMail) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByMailAndChangedAndInit(offset, limit, iMail,iChanged,iInit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByMailAndChangedAndInit's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByMailAndChangedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMail := self.Args("mail").String()
	iChanged := self.Args("changed").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iMail) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByMailAndChangedAndDefaultLangcode(offset, limit, iMail,iChanged,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByMailAndChangedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByMailAndAccessAndLoginHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMail := self.Args("mail").String()
	iAccess := self.Args("access").MustInt()
	iLogin := self.Args("login").MustInt()

	if helper.IsHas(iMail) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByMailAndAccessAndLogin(offset, limit, iMail,iAccess,iLogin)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByMailAndAccessAndLogin's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByMailAndAccessAndInitHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMail := self.Args("mail").String()
	iAccess := self.Args("access").MustInt()
	iInit := self.Args("init").String()

	if helper.IsHas(iMail) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByMailAndAccessAndInit(offset, limit, iMail,iAccess,iInit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByMailAndAccessAndInit's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByMailAndAccessAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMail := self.Args("mail").String()
	iAccess := self.Args("access").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iMail) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByMailAndAccessAndDefaultLangcode(offset, limit, iMail,iAccess,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByMailAndAccessAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByMailAndLoginAndInitHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMail := self.Args("mail").String()
	iLogin := self.Args("login").MustInt()
	iInit := self.Args("init").String()

	if helper.IsHas(iMail) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByMailAndLoginAndInit(offset, limit, iMail,iLogin,iInit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByMailAndLoginAndInit's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByMailAndLoginAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMail := self.Args("mail").String()
	iLogin := self.Args("login").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iMail) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByMailAndLoginAndDefaultLangcode(offset, limit, iMail,iLogin,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByMailAndLoginAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByMailAndInitAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMail := self.Args("mail").String()
	iInit := self.Args("init").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iMail) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByMailAndInitAndDefaultLangcode(offset, limit, iMail,iInit,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByMailAndInitAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByTimezoneAndStatusAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTimezone := self.Args("timezone").String()
	iStatus := self.Args("status").MustInt()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iTimezone) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByTimezoneAndStatusAndCreated(offset, limit, iTimezone,iStatus,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByTimezoneAndStatusAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByTimezoneAndStatusAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTimezone := self.Args("timezone").String()
	iStatus := self.Args("status").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iTimezone) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByTimezoneAndStatusAndChanged(offset, limit, iTimezone,iStatus,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByTimezoneAndStatusAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByTimezoneAndStatusAndAccessHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTimezone := self.Args("timezone").String()
	iStatus := self.Args("status").MustInt()
	iAccess := self.Args("access").MustInt()

	if helper.IsHas(iTimezone) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByTimezoneAndStatusAndAccess(offset, limit, iTimezone,iStatus,iAccess)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByTimezoneAndStatusAndAccess's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByTimezoneAndStatusAndLoginHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTimezone := self.Args("timezone").String()
	iStatus := self.Args("status").MustInt()
	iLogin := self.Args("login").MustInt()

	if helper.IsHas(iTimezone) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByTimezoneAndStatusAndLogin(offset, limit, iTimezone,iStatus,iLogin)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByTimezoneAndStatusAndLogin's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByTimezoneAndStatusAndInitHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTimezone := self.Args("timezone").String()
	iStatus := self.Args("status").MustInt()
	iInit := self.Args("init").String()

	if helper.IsHas(iTimezone) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByTimezoneAndStatusAndInit(offset, limit, iTimezone,iStatus,iInit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByTimezoneAndStatusAndInit's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByTimezoneAndStatusAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTimezone := self.Args("timezone").String()
	iStatus := self.Args("status").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iTimezone) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByTimezoneAndStatusAndDefaultLangcode(offset, limit, iTimezone,iStatus,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByTimezoneAndStatusAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByTimezoneAndCreatedAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTimezone := self.Args("timezone").String()
	iCreated := self.Args("created").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iTimezone) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByTimezoneAndCreatedAndChanged(offset, limit, iTimezone,iCreated,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByTimezoneAndCreatedAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByTimezoneAndCreatedAndAccessHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTimezone := self.Args("timezone").String()
	iCreated := self.Args("created").MustInt()
	iAccess := self.Args("access").MustInt()

	if helper.IsHas(iTimezone) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByTimezoneAndCreatedAndAccess(offset, limit, iTimezone,iCreated,iAccess)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByTimezoneAndCreatedAndAccess's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByTimezoneAndCreatedAndLoginHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTimezone := self.Args("timezone").String()
	iCreated := self.Args("created").MustInt()
	iLogin := self.Args("login").MustInt()

	if helper.IsHas(iTimezone) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByTimezoneAndCreatedAndLogin(offset, limit, iTimezone,iCreated,iLogin)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByTimezoneAndCreatedAndLogin's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByTimezoneAndCreatedAndInitHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTimezone := self.Args("timezone").String()
	iCreated := self.Args("created").MustInt()
	iInit := self.Args("init").String()

	if helper.IsHas(iTimezone) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByTimezoneAndCreatedAndInit(offset, limit, iTimezone,iCreated,iInit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByTimezoneAndCreatedAndInit's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByTimezoneAndCreatedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTimezone := self.Args("timezone").String()
	iCreated := self.Args("created").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iTimezone) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByTimezoneAndCreatedAndDefaultLangcode(offset, limit, iTimezone,iCreated,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByTimezoneAndCreatedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByTimezoneAndChangedAndAccessHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTimezone := self.Args("timezone").String()
	iChanged := self.Args("changed").MustInt()
	iAccess := self.Args("access").MustInt()

	if helper.IsHas(iTimezone) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByTimezoneAndChangedAndAccess(offset, limit, iTimezone,iChanged,iAccess)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByTimezoneAndChangedAndAccess's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByTimezoneAndChangedAndLoginHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTimezone := self.Args("timezone").String()
	iChanged := self.Args("changed").MustInt()
	iLogin := self.Args("login").MustInt()

	if helper.IsHas(iTimezone) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByTimezoneAndChangedAndLogin(offset, limit, iTimezone,iChanged,iLogin)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByTimezoneAndChangedAndLogin's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByTimezoneAndChangedAndInitHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTimezone := self.Args("timezone").String()
	iChanged := self.Args("changed").MustInt()
	iInit := self.Args("init").String()

	if helper.IsHas(iTimezone) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByTimezoneAndChangedAndInit(offset, limit, iTimezone,iChanged,iInit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByTimezoneAndChangedAndInit's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByTimezoneAndChangedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTimezone := self.Args("timezone").String()
	iChanged := self.Args("changed").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iTimezone) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByTimezoneAndChangedAndDefaultLangcode(offset, limit, iTimezone,iChanged,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByTimezoneAndChangedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByTimezoneAndAccessAndLoginHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTimezone := self.Args("timezone").String()
	iAccess := self.Args("access").MustInt()
	iLogin := self.Args("login").MustInt()

	if helper.IsHas(iTimezone) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByTimezoneAndAccessAndLogin(offset, limit, iTimezone,iAccess,iLogin)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByTimezoneAndAccessAndLogin's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByTimezoneAndAccessAndInitHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTimezone := self.Args("timezone").String()
	iAccess := self.Args("access").MustInt()
	iInit := self.Args("init").String()

	if helper.IsHas(iTimezone) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByTimezoneAndAccessAndInit(offset, limit, iTimezone,iAccess,iInit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByTimezoneAndAccessAndInit's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByTimezoneAndAccessAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTimezone := self.Args("timezone").String()
	iAccess := self.Args("access").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iTimezone) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByTimezoneAndAccessAndDefaultLangcode(offset, limit, iTimezone,iAccess,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByTimezoneAndAccessAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByTimezoneAndLoginAndInitHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTimezone := self.Args("timezone").String()
	iLogin := self.Args("login").MustInt()
	iInit := self.Args("init").String()

	if helper.IsHas(iTimezone) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByTimezoneAndLoginAndInit(offset, limit, iTimezone,iLogin,iInit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByTimezoneAndLoginAndInit's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByTimezoneAndLoginAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTimezone := self.Args("timezone").String()
	iLogin := self.Args("login").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iTimezone) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByTimezoneAndLoginAndDefaultLangcode(offset, limit, iTimezone,iLogin,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByTimezoneAndLoginAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByTimezoneAndInitAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTimezone := self.Args("timezone").String()
	iInit := self.Args("init").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iTimezone) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByTimezoneAndInitAndDefaultLangcode(offset, limit, iTimezone,iInit,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByTimezoneAndInitAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByStatusAndCreatedAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iStatus := self.Args("status").MustInt()
	iCreated := self.Args("created").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iStatus) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByStatusAndCreatedAndChanged(offset, limit, iStatus,iCreated,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByStatusAndCreatedAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByStatusAndCreatedAndAccessHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iStatus := self.Args("status").MustInt()
	iCreated := self.Args("created").MustInt()
	iAccess := self.Args("access").MustInt()

	if helper.IsHas(iStatus) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByStatusAndCreatedAndAccess(offset, limit, iStatus,iCreated,iAccess)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByStatusAndCreatedAndAccess's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByStatusAndCreatedAndLoginHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iStatus := self.Args("status").MustInt()
	iCreated := self.Args("created").MustInt()
	iLogin := self.Args("login").MustInt()

	if helper.IsHas(iStatus) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByStatusAndCreatedAndLogin(offset, limit, iStatus,iCreated,iLogin)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByStatusAndCreatedAndLogin's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByStatusAndCreatedAndInitHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iStatus := self.Args("status").MustInt()
	iCreated := self.Args("created").MustInt()
	iInit := self.Args("init").String()

	if helper.IsHas(iStatus) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByStatusAndCreatedAndInit(offset, limit, iStatus,iCreated,iInit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByStatusAndCreatedAndInit's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByStatusAndCreatedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iStatus := self.Args("status").MustInt()
	iCreated := self.Args("created").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iStatus) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByStatusAndCreatedAndDefaultLangcode(offset, limit, iStatus,iCreated,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByStatusAndCreatedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByStatusAndChangedAndAccessHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iStatus := self.Args("status").MustInt()
	iChanged := self.Args("changed").MustInt()
	iAccess := self.Args("access").MustInt()

	if helper.IsHas(iStatus) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByStatusAndChangedAndAccess(offset, limit, iStatus,iChanged,iAccess)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByStatusAndChangedAndAccess's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByStatusAndChangedAndLoginHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iStatus := self.Args("status").MustInt()
	iChanged := self.Args("changed").MustInt()
	iLogin := self.Args("login").MustInt()

	if helper.IsHas(iStatus) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByStatusAndChangedAndLogin(offset, limit, iStatus,iChanged,iLogin)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByStatusAndChangedAndLogin's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByStatusAndChangedAndInitHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iStatus := self.Args("status").MustInt()
	iChanged := self.Args("changed").MustInt()
	iInit := self.Args("init").String()

	if helper.IsHas(iStatus) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByStatusAndChangedAndInit(offset, limit, iStatus,iChanged,iInit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByStatusAndChangedAndInit's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByStatusAndChangedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iStatus := self.Args("status").MustInt()
	iChanged := self.Args("changed").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iStatus) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByStatusAndChangedAndDefaultLangcode(offset, limit, iStatus,iChanged,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByStatusAndChangedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByStatusAndAccessAndLoginHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iStatus := self.Args("status").MustInt()
	iAccess := self.Args("access").MustInt()
	iLogin := self.Args("login").MustInt()

	if helper.IsHas(iStatus) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByStatusAndAccessAndLogin(offset, limit, iStatus,iAccess,iLogin)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByStatusAndAccessAndLogin's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByStatusAndAccessAndInitHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iStatus := self.Args("status").MustInt()
	iAccess := self.Args("access").MustInt()
	iInit := self.Args("init").String()

	if helper.IsHas(iStatus) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByStatusAndAccessAndInit(offset, limit, iStatus,iAccess,iInit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByStatusAndAccessAndInit's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByStatusAndAccessAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iStatus := self.Args("status").MustInt()
	iAccess := self.Args("access").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iStatus) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByStatusAndAccessAndDefaultLangcode(offset, limit, iStatus,iAccess,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByStatusAndAccessAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByStatusAndLoginAndInitHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iStatus := self.Args("status").MustInt()
	iLogin := self.Args("login").MustInt()
	iInit := self.Args("init").String()

	if helper.IsHas(iStatus) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByStatusAndLoginAndInit(offset, limit, iStatus,iLogin,iInit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByStatusAndLoginAndInit's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByStatusAndLoginAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iStatus := self.Args("status").MustInt()
	iLogin := self.Args("login").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iStatus) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByStatusAndLoginAndDefaultLangcode(offset, limit, iStatus,iLogin,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByStatusAndLoginAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByStatusAndInitAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iStatus := self.Args("status").MustInt()
	iInit := self.Args("init").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iStatus) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByStatusAndInitAndDefaultLangcode(offset, limit, iStatus,iInit,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByStatusAndInitAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByCreatedAndChangedAndAccessHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCreated := self.Args("created").MustInt()
	iChanged := self.Args("changed").MustInt()
	iAccess := self.Args("access").MustInt()

	if helper.IsHas(iCreated) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByCreatedAndChangedAndAccess(offset, limit, iCreated,iChanged,iAccess)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByCreatedAndChangedAndAccess's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByCreatedAndChangedAndLoginHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCreated := self.Args("created").MustInt()
	iChanged := self.Args("changed").MustInt()
	iLogin := self.Args("login").MustInt()

	if helper.IsHas(iCreated) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByCreatedAndChangedAndLogin(offset, limit, iCreated,iChanged,iLogin)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByCreatedAndChangedAndLogin's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByCreatedAndChangedAndInitHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCreated := self.Args("created").MustInt()
	iChanged := self.Args("changed").MustInt()
	iInit := self.Args("init").String()

	if helper.IsHas(iCreated) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByCreatedAndChangedAndInit(offset, limit, iCreated,iChanged,iInit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByCreatedAndChangedAndInit's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByCreatedAndChangedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCreated := self.Args("created").MustInt()
	iChanged := self.Args("changed").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iCreated) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByCreatedAndChangedAndDefaultLangcode(offset, limit, iCreated,iChanged,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByCreatedAndChangedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByCreatedAndAccessAndLoginHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCreated := self.Args("created").MustInt()
	iAccess := self.Args("access").MustInt()
	iLogin := self.Args("login").MustInt()

	if helper.IsHas(iCreated) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByCreatedAndAccessAndLogin(offset, limit, iCreated,iAccess,iLogin)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByCreatedAndAccessAndLogin's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByCreatedAndAccessAndInitHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCreated := self.Args("created").MustInt()
	iAccess := self.Args("access").MustInt()
	iInit := self.Args("init").String()

	if helper.IsHas(iCreated) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByCreatedAndAccessAndInit(offset, limit, iCreated,iAccess,iInit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByCreatedAndAccessAndInit's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByCreatedAndAccessAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCreated := self.Args("created").MustInt()
	iAccess := self.Args("access").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iCreated) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByCreatedAndAccessAndDefaultLangcode(offset, limit, iCreated,iAccess,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByCreatedAndAccessAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByCreatedAndLoginAndInitHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCreated := self.Args("created").MustInt()
	iLogin := self.Args("login").MustInt()
	iInit := self.Args("init").String()

	if helper.IsHas(iCreated) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByCreatedAndLoginAndInit(offset, limit, iCreated,iLogin,iInit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByCreatedAndLoginAndInit's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByCreatedAndLoginAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCreated := self.Args("created").MustInt()
	iLogin := self.Args("login").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iCreated) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByCreatedAndLoginAndDefaultLangcode(offset, limit, iCreated,iLogin,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByCreatedAndLoginAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByCreatedAndInitAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCreated := self.Args("created").MustInt()
	iInit := self.Args("init").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iCreated) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByCreatedAndInitAndDefaultLangcode(offset, limit, iCreated,iInit,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByCreatedAndInitAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByChangedAndAccessAndLoginHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iChanged := self.Args("changed").MustInt()
	iAccess := self.Args("access").MustInt()
	iLogin := self.Args("login").MustInt()

	if helper.IsHas(iChanged) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByChangedAndAccessAndLogin(offset, limit, iChanged,iAccess,iLogin)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByChangedAndAccessAndLogin's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByChangedAndAccessAndInitHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iChanged := self.Args("changed").MustInt()
	iAccess := self.Args("access").MustInt()
	iInit := self.Args("init").String()

	if helper.IsHas(iChanged) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByChangedAndAccessAndInit(offset, limit, iChanged,iAccess,iInit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByChangedAndAccessAndInit's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByChangedAndAccessAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iChanged := self.Args("changed").MustInt()
	iAccess := self.Args("access").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iChanged) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByChangedAndAccessAndDefaultLangcode(offset, limit, iChanged,iAccess,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByChangedAndAccessAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByChangedAndLoginAndInitHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iChanged := self.Args("changed").MustInt()
	iLogin := self.Args("login").MustInt()
	iInit := self.Args("init").String()

	if helper.IsHas(iChanged) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByChangedAndLoginAndInit(offset, limit, iChanged,iLogin,iInit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByChangedAndLoginAndInit's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByChangedAndLoginAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iChanged := self.Args("changed").MustInt()
	iLogin := self.Args("login").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iChanged) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByChangedAndLoginAndDefaultLangcode(offset, limit, iChanged,iLogin,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByChangedAndLoginAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByChangedAndInitAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iChanged := self.Args("changed").MustInt()
	iInit := self.Args("init").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iChanged) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByChangedAndInitAndDefaultLangcode(offset, limit, iChanged,iInit,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByChangedAndInitAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByAccessAndLoginAndInitHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iAccess := self.Args("access").MustInt()
	iLogin := self.Args("login").MustInt()
	iInit := self.Args("init").String()

	if helper.IsHas(iAccess) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByAccessAndLoginAndInit(offset, limit, iAccess,iLogin,iInit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByAccessAndLoginAndInit's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByAccessAndLoginAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iAccess := self.Args("access").MustInt()
	iLogin := self.Args("login").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iAccess) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByAccessAndLoginAndDefaultLangcode(offset, limit, iAccess,iLogin,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByAccessAndLoginAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByAccessAndInitAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iAccess := self.Args("access").MustInt()
	iInit := self.Args("init").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iAccess) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByAccessAndInitAndDefaultLangcode(offset, limit, iAccess,iInit,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByAccessAndInitAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByLoginAndInitAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLogin := self.Args("login").MustInt()
	iInit := self.Args("init").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iLogin) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByLoginAndInitAndDefaultLangcode(offset, limit, iLogin,iInit,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByLoginAndInitAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByUidAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt64()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iUid) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByUidAndLangcode(offset, limit, iUid,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByUidAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByUidAndPreferredLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt64()
	iPreferredLangcode := self.Args("preferred_langcode").String()

	if helper.IsHas(iUid) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByUidAndPreferredLangcode(offset, limit, iUid,iPreferredLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByUidAndPreferredLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByUidAndPreferredAdminLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt64()
	iPreferredAdminLangcode := self.Args("preferred_admin_langcode").String()

	if helper.IsHas(iUid) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByUidAndPreferredAdminLangcode(offset, limit, iUid,iPreferredAdminLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByUidAndPreferredAdminLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByUidAndNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt64()
	iName := self.Args("name").String()

	if helper.IsHas(iUid) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByUidAndName(offset, limit, iUid,iName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByUidAndName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByUidAndPassHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt64()
	iPass := self.Args("pass").String()

	if helper.IsHas(iUid) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByUidAndPass(offset, limit, iUid,iPass)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByUidAndPass's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByUidAndMailHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt64()
	iMail := self.Args("mail").String()

	if helper.IsHas(iUid) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByUidAndMail(offset, limit, iUid,iMail)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByUidAndMail's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByUidAndTimezoneHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt64()
	iTimezone := self.Args("timezone").String()

	if helper.IsHas(iUid) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByUidAndTimezone(offset, limit, iUid,iTimezone)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByUidAndTimezone's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByUidAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt64()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iUid) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByUidAndStatus(offset, limit, iUid,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByUidAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByUidAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt64()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iUid) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByUidAndCreated(offset, limit, iUid,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByUidAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByUidAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt64()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iUid) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByUidAndChanged(offset, limit, iUid,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByUidAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByUidAndAccessHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt64()
	iAccess := self.Args("access").MustInt()

	if helper.IsHas(iUid) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByUidAndAccess(offset, limit, iUid,iAccess)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByUidAndAccess's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByUidAndLoginHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt64()
	iLogin := self.Args("login").MustInt()

	if helper.IsHas(iUid) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByUidAndLogin(offset, limit, iUid,iLogin)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByUidAndLogin's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByUidAndInitHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt64()
	iInit := self.Args("init").String()

	if helper.IsHas(iUid) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByUidAndInit(offset, limit, iUid,iInit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByUidAndInit's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByUidAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt64()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iUid) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByUidAndDefaultLangcode(offset, limit, iUid,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByUidAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByLangcodeAndPreferredLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iPreferredLangcode := self.Args("preferred_langcode").String()

	if helper.IsHas(iLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByLangcodeAndPreferredLangcode(offset, limit, iLangcode,iPreferredLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByLangcodeAndPreferredLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByLangcodeAndPreferredAdminLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iPreferredAdminLangcode := self.Args("preferred_admin_langcode").String()

	if helper.IsHas(iLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByLangcodeAndPreferredAdminLangcode(offset, limit, iLangcode,iPreferredAdminLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByLangcodeAndPreferredAdminLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByLangcodeAndNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iName := self.Args("name").String()

	if helper.IsHas(iLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByLangcodeAndName(offset, limit, iLangcode,iName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByLangcodeAndName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByLangcodeAndPassHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iPass := self.Args("pass").String()

	if helper.IsHas(iLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByLangcodeAndPass(offset, limit, iLangcode,iPass)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByLangcodeAndPass's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByLangcodeAndMailHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iMail := self.Args("mail").String()

	if helper.IsHas(iLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByLangcodeAndMail(offset, limit, iLangcode,iMail)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByLangcodeAndMail's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByLangcodeAndTimezoneHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iTimezone := self.Args("timezone").String()

	if helper.IsHas(iLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByLangcodeAndTimezone(offset, limit, iLangcode,iTimezone)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByLangcodeAndTimezone's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByLangcodeAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByLangcodeAndStatus(offset, limit, iLangcode,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByLangcodeAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByLangcodeAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByLangcodeAndCreated(offset, limit, iLangcode,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByLangcodeAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByLangcodeAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByLangcodeAndChanged(offset, limit, iLangcode,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByLangcodeAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByLangcodeAndAccessHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iAccess := self.Args("access").MustInt()

	if helper.IsHas(iLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByLangcodeAndAccess(offset, limit, iLangcode,iAccess)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByLangcodeAndAccess's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByLangcodeAndLoginHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iLogin := self.Args("login").MustInt()

	if helper.IsHas(iLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByLangcodeAndLogin(offset, limit, iLangcode,iLogin)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByLangcodeAndLogin's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByLangcodeAndInitHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iInit := self.Args("init").String()

	if helper.IsHas(iLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByLangcodeAndInit(offset, limit, iLangcode,iInit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByLangcodeAndInit's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByLangcodeAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByLangcodeAndDefaultLangcode(offset, limit, iLangcode,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByLangcodeAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredLangcodeAndPreferredAdminLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredLangcode := self.Args("preferred_langcode").String()
	iPreferredAdminLangcode := self.Args("preferred_admin_langcode").String()

	if helper.IsHas(iPreferredLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredLangcodeAndPreferredAdminLangcode(offset, limit, iPreferredLangcode,iPreferredAdminLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredLangcodeAndPreferredAdminLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredLangcodeAndNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredLangcode := self.Args("preferred_langcode").String()
	iName := self.Args("name").String()

	if helper.IsHas(iPreferredLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredLangcodeAndName(offset, limit, iPreferredLangcode,iName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredLangcodeAndName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredLangcodeAndPassHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredLangcode := self.Args("preferred_langcode").String()
	iPass := self.Args("pass").String()

	if helper.IsHas(iPreferredLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredLangcodeAndPass(offset, limit, iPreferredLangcode,iPass)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredLangcodeAndPass's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredLangcodeAndMailHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredLangcode := self.Args("preferred_langcode").String()
	iMail := self.Args("mail").String()

	if helper.IsHas(iPreferredLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredLangcodeAndMail(offset, limit, iPreferredLangcode,iMail)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredLangcodeAndMail's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredLangcodeAndTimezoneHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredLangcode := self.Args("preferred_langcode").String()
	iTimezone := self.Args("timezone").String()

	if helper.IsHas(iPreferredLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredLangcodeAndTimezone(offset, limit, iPreferredLangcode,iTimezone)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredLangcodeAndTimezone's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredLangcodeAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredLangcode := self.Args("preferred_langcode").String()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iPreferredLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredLangcodeAndStatus(offset, limit, iPreferredLangcode,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredLangcodeAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredLangcodeAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredLangcode := self.Args("preferred_langcode").String()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iPreferredLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredLangcodeAndCreated(offset, limit, iPreferredLangcode,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredLangcodeAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredLangcodeAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredLangcode := self.Args("preferred_langcode").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iPreferredLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredLangcodeAndChanged(offset, limit, iPreferredLangcode,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredLangcodeAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredLangcodeAndAccessHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredLangcode := self.Args("preferred_langcode").String()
	iAccess := self.Args("access").MustInt()

	if helper.IsHas(iPreferredLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredLangcodeAndAccess(offset, limit, iPreferredLangcode,iAccess)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredLangcodeAndAccess's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredLangcodeAndLoginHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredLangcode := self.Args("preferred_langcode").String()
	iLogin := self.Args("login").MustInt()

	if helper.IsHas(iPreferredLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredLangcodeAndLogin(offset, limit, iPreferredLangcode,iLogin)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredLangcodeAndLogin's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredLangcodeAndInitHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredLangcode := self.Args("preferred_langcode").String()
	iInit := self.Args("init").String()

	if helper.IsHas(iPreferredLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredLangcodeAndInit(offset, limit, iPreferredLangcode,iInit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredLangcodeAndInit's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredLangcodeAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredLangcode := self.Args("preferred_langcode").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iPreferredLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredLangcodeAndDefaultLangcode(offset, limit, iPreferredLangcode,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredLangcodeAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredAdminLangcodeAndNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredAdminLangcode := self.Args("preferred_admin_langcode").String()
	iName := self.Args("name").String()

	if helper.IsHas(iPreferredAdminLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredAdminLangcodeAndName(offset, limit, iPreferredAdminLangcode,iName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredAdminLangcodeAndName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredAdminLangcodeAndPassHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredAdminLangcode := self.Args("preferred_admin_langcode").String()
	iPass := self.Args("pass").String()

	if helper.IsHas(iPreferredAdminLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredAdminLangcodeAndPass(offset, limit, iPreferredAdminLangcode,iPass)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredAdminLangcodeAndPass's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredAdminLangcodeAndMailHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredAdminLangcode := self.Args("preferred_admin_langcode").String()
	iMail := self.Args("mail").String()

	if helper.IsHas(iPreferredAdminLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredAdminLangcodeAndMail(offset, limit, iPreferredAdminLangcode,iMail)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredAdminLangcodeAndMail's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredAdminLangcodeAndTimezoneHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredAdminLangcode := self.Args("preferred_admin_langcode").String()
	iTimezone := self.Args("timezone").String()

	if helper.IsHas(iPreferredAdminLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredAdminLangcodeAndTimezone(offset, limit, iPreferredAdminLangcode,iTimezone)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredAdminLangcodeAndTimezone's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredAdminLangcodeAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredAdminLangcode := self.Args("preferred_admin_langcode").String()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iPreferredAdminLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredAdminLangcodeAndStatus(offset, limit, iPreferredAdminLangcode,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredAdminLangcodeAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredAdminLangcodeAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredAdminLangcode := self.Args("preferred_admin_langcode").String()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iPreferredAdminLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredAdminLangcodeAndCreated(offset, limit, iPreferredAdminLangcode,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredAdminLangcodeAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredAdminLangcodeAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredAdminLangcode := self.Args("preferred_admin_langcode").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iPreferredAdminLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredAdminLangcodeAndChanged(offset, limit, iPreferredAdminLangcode,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredAdminLangcodeAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredAdminLangcodeAndAccessHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredAdminLangcode := self.Args("preferred_admin_langcode").String()
	iAccess := self.Args("access").MustInt()

	if helper.IsHas(iPreferredAdminLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredAdminLangcodeAndAccess(offset, limit, iPreferredAdminLangcode,iAccess)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredAdminLangcodeAndAccess's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredAdminLangcodeAndLoginHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredAdminLangcode := self.Args("preferred_admin_langcode").String()
	iLogin := self.Args("login").MustInt()

	if helper.IsHas(iPreferredAdminLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredAdminLangcodeAndLogin(offset, limit, iPreferredAdminLangcode,iLogin)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredAdminLangcodeAndLogin's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredAdminLangcodeAndInitHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredAdminLangcode := self.Args("preferred_admin_langcode").String()
	iInit := self.Args("init").String()

	if helper.IsHas(iPreferredAdminLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredAdminLangcodeAndInit(offset, limit, iPreferredAdminLangcode,iInit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredAdminLangcodeAndInit's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPreferredAdminLangcodeAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPreferredAdminLangcode := self.Args("preferred_admin_langcode").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iPreferredAdminLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPreferredAdminLangcodeAndDefaultLangcode(offset, limit, iPreferredAdminLangcode,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPreferredAdminLangcodeAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByNameAndPassHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iPass := self.Args("pass").String()

	if helper.IsHas(iName) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByNameAndPass(offset, limit, iName,iPass)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByNameAndPass's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByNameAndMailHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iMail := self.Args("mail").String()

	if helper.IsHas(iName) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByNameAndMail(offset, limit, iName,iMail)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByNameAndMail's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByNameAndTimezoneHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iTimezone := self.Args("timezone").String()

	if helper.IsHas(iName) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByNameAndTimezone(offset, limit, iName,iTimezone)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByNameAndTimezone's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByNameAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iName) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByNameAndStatus(offset, limit, iName,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByNameAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByNameAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iName) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByNameAndCreated(offset, limit, iName,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByNameAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByNameAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iName) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByNameAndChanged(offset, limit, iName,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByNameAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByNameAndAccessHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iAccess := self.Args("access").MustInt()

	if helper.IsHas(iName) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByNameAndAccess(offset, limit, iName,iAccess)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByNameAndAccess's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByNameAndLoginHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iLogin := self.Args("login").MustInt()

	if helper.IsHas(iName) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByNameAndLogin(offset, limit, iName,iLogin)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByNameAndLogin's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByNameAndInitHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iInit := self.Args("init").String()

	if helper.IsHas(iName) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByNameAndInit(offset, limit, iName,iInit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByNameAndInit's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByNameAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iName) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByNameAndDefaultLangcode(offset, limit, iName,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByNameAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPassAndMailHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPass := self.Args("pass").String()
	iMail := self.Args("mail").String()

	if helper.IsHas(iPass) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPassAndMail(offset, limit, iPass,iMail)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPassAndMail's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPassAndTimezoneHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPass := self.Args("pass").String()
	iTimezone := self.Args("timezone").String()

	if helper.IsHas(iPass) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPassAndTimezone(offset, limit, iPass,iTimezone)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPassAndTimezone's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPassAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPass := self.Args("pass").String()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iPass) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPassAndStatus(offset, limit, iPass,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPassAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPassAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPass := self.Args("pass").String()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iPass) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPassAndCreated(offset, limit, iPass,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPassAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPassAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPass := self.Args("pass").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iPass) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPassAndChanged(offset, limit, iPass,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPassAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPassAndAccessHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPass := self.Args("pass").String()
	iAccess := self.Args("access").MustInt()

	if helper.IsHas(iPass) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPassAndAccess(offset, limit, iPass,iAccess)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPassAndAccess's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPassAndLoginHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPass := self.Args("pass").String()
	iLogin := self.Args("login").MustInt()

	if helper.IsHas(iPass) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPassAndLogin(offset, limit, iPass,iLogin)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPassAndLogin's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPassAndInitHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPass := self.Args("pass").String()
	iInit := self.Args("init").String()

	if helper.IsHas(iPass) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPassAndInit(offset, limit, iPass,iInit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPassAndInit's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByPassAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPass := self.Args("pass").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iPass) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByPassAndDefaultLangcode(offset, limit, iPass,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByPassAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByMailAndTimezoneHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMail := self.Args("mail").String()
	iTimezone := self.Args("timezone").String()

	if helper.IsHas(iMail) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByMailAndTimezone(offset, limit, iMail,iTimezone)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByMailAndTimezone's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByMailAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMail := self.Args("mail").String()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iMail) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByMailAndStatus(offset, limit, iMail,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByMailAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByMailAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMail := self.Args("mail").String()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iMail) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByMailAndCreated(offset, limit, iMail,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByMailAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByMailAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMail := self.Args("mail").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iMail) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByMailAndChanged(offset, limit, iMail,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByMailAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByMailAndAccessHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMail := self.Args("mail").String()
	iAccess := self.Args("access").MustInt()

	if helper.IsHas(iMail) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByMailAndAccess(offset, limit, iMail,iAccess)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByMailAndAccess's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByMailAndLoginHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMail := self.Args("mail").String()
	iLogin := self.Args("login").MustInt()

	if helper.IsHas(iMail) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByMailAndLogin(offset, limit, iMail,iLogin)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByMailAndLogin's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByMailAndInitHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMail := self.Args("mail").String()
	iInit := self.Args("init").String()

	if helper.IsHas(iMail) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByMailAndInit(offset, limit, iMail,iInit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByMailAndInit's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByMailAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMail := self.Args("mail").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iMail) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByMailAndDefaultLangcode(offset, limit, iMail,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByMailAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByTimezoneAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTimezone := self.Args("timezone").String()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iTimezone) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByTimezoneAndStatus(offset, limit, iTimezone,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByTimezoneAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByTimezoneAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTimezone := self.Args("timezone").String()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iTimezone) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByTimezoneAndCreated(offset, limit, iTimezone,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByTimezoneAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByTimezoneAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTimezone := self.Args("timezone").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iTimezone) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByTimezoneAndChanged(offset, limit, iTimezone,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByTimezoneAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByTimezoneAndAccessHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTimezone := self.Args("timezone").String()
	iAccess := self.Args("access").MustInt()

	if helper.IsHas(iTimezone) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByTimezoneAndAccess(offset, limit, iTimezone,iAccess)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByTimezoneAndAccess's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByTimezoneAndLoginHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTimezone := self.Args("timezone").String()
	iLogin := self.Args("login").MustInt()

	if helper.IsHas(iTimezone) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByTimezoneAndLogin(offset, limit, iTimezone,iLogin)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByTimezoneAndLogin's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByTimezoneAndInitHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTimezone := self.Args("timezone").String()
	iInit := self.Args("init").String()

	if helper.IsHas(iTimezone) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByTimezoneAndInit(offset, limit, iTimezone,iInit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByTimezoneAndInit's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByTimezoneAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTimezone := self.Args("timezone").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iTimezone) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByTimezoneAndDefaultLangcode(offset, limit, iTimezone,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByTimezoneAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByStatusAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iStatus := self.Args("status").MustInt()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iStatus) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByStatusAndCreated(offset, limit, iStatus,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByStatusAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByStatusAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iStatus := self.Args("status").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iStatus) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByStatusAndChanged(offset, limit, iStatus,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByStatusAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByStatusAndAccessHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iStatus := self.Args("status").MustInt()
	iAccess := self.Args("access").MustInt()

	if helper.IsHas(iStatus) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByStatusAndAccess(offset, limit, iStatus,iAccess)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByStatusAndAccess's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByStatusAndLoginHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iStatus := self.Args("status").MustInt()
	iLogin := self.Args("login").MustInt()

	if helper.IsHas(iStatus) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByStatusAndLogin(offset, limit, iStatus,iLogin)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByStatusAndLogin's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByStatusAndInitHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iStatus := self.Args("status").MustInt()
	iInit := self.Args("init").String()

	if helper.IsHas(iStatus) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByStatusAndInit(offset, limit, iStatus,iInit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByStatusAndInit's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByStatusAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iStatus := self.Args("status").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iStatus) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByStatusAndDefaultLangcode(offset, limit, iStatus,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByStatusAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByCreatedAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCreated := self.Args("created").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iCreated) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByCreatedAndChanged(offset, limit, iCreated,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByCreatedAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByCreatedAndAccessHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCreated := self.Args("created").MustInt()
	iAccess := self.Args("access").MustInt()

	if helper.IsHas(iCreated) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByCreatedAndAccess(offset, limit, iCreated,iAccess)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByCreatedAndAccess's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByCreatedAndLoginHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCreated := self.Args("created").MustInt()
	iLogin := self.Args("login").MustInt()

	if helper.IsHas(iCreated) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByCreatedAndLogin(offset, limit, iCreated,iLogin)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByCreatedAndLogin's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByCreatedAndInitHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCreated := self.Args("created").MustInt()
	iInit := self.Args("init").String()

	if helper.IsHas(iCreated) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByCreatedAndInit(offset, limit, iCreated,iInit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByCreatedAndInit's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByCreatedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCreated := self.Args("created").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iCreated) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByCreatedAndDefaultLangcode(offset, limit, iCreated,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByCreatedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByChangedAndAccessHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iChanged := self.Args("changed").MustInt()
	iAccess := self.Args("access").MustInt()

	if helper.IsHas(iChanged) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByChangedAndAccess(offset, limit, iChanged,iAccess)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByChangedAndAccess's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByChangedAndLoginHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iChanged := self.Args("changed").MustInt()
	iLogin := self.Args("login").MustInt()

	if helper.IsHas(iChanged) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByChangedAndLogin(offset, limit, iChanged,iLogin)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByChangedAndLogin's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByChangedAndInitHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iChanged := self.Args("changed").MustInt()
	iInit := self.Args("init").String()

	if helper.IsHas(iChanged) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByChangedAndInit(offset, limit, iChanged,iInit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByChangedAndInit's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByChangedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iChanged := self.Args("changed").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iChanged) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByChangedAndDefaultLangcode(offset, limit, iChanged,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByChangedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByAccessAndLoginHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iAccess := self.Args("access").MustInt()
	iLogin := self.Args("login").MustInt()

	if helper.IsHas(iAccess) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByAccessAndLogin(offset, limit, iAccess,iLogin)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByAccessAndLogin's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByAccessAndInitHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iAccess := self.Args("access").MustInt()
	iInit := self.Args("init").String()

	if helper.IsHas(iAccess) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByAccessAndInit(offset, limit, iAccess,iInit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByAccessAndInit's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByAccessAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iAccess := self.Args("access").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iAccess) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByAccessAndDefaultLangcode(offset, limit, iAccess,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByAccessAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByLoginAndInitHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLogin := self.Args("login").MustInt()
	iInit := self.Args("init").String()

	if helper.IsHas(iLogin) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByLoginAndInit(offset, limit, iLogin,iInit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByLoginAndInit's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByLoginAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLogin := self.Args("login").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iLogin) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByLoginAndDefaultLangcode(offset, limit, iLogin,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByLoginAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasByInitAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iInit := self.Args("init").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iInit) {
		_UsersFieldData, _error := model.GetUsersFieldDatasByInitAndDefaultLangcode(offset, limit, iInit,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatasByInitAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDatasHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	if (offset > 0) && (len(field) > 0) {
		_UsersFieldData, _error := model.GetUsersFieldDatas(offset, limit, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDatas' args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasUsersFieldDataViaUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iUid := self.Args("uid").MustInt64()
	if helper.IsHas(iUid) {
		_UsersFieldData := model.HasUsersFieldDataViaUid(iUid)
		var m = map[string]interface{}{}
		m["users_field_data"] = _UsersFieldData
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasUsersFieldDataViaUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasUsersFieldDataViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iLangcode := self.Args("langcode").String()
	if helper.IsHas(iLangcode) {
		_UsersFieldData := model.HasUsersFieldDataViaLangcode(iLangcode)
		var m = map[string]interface{}{}
		m["users_field_data"] = _UsersFieldData
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasUsersFieldDataViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasUsersFieldDataViaPreferredLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iPreferredLangcode := self.Args("preferred_langcode").String()
	if helper.IsHas(iPreferredLangcode) {
		_UsersFieldData := model.HasUsersFieldDataViaPreferredLangcode(iPreferredLangcode)
		var m = map[string]interface{}{}
		m["users_field_data"] = _UsersFieldData
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasUsersFieldDataViaPreferredLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasUsersFieldDataViaPreferredAdminLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iPreferredAdminLangcode := self.Args("preferred_admin_langcode").String()
	if helper.IsHas(iPreferredAdminLangcode) {
		_UsersFieldData := model.HasUsersFieldDataViaPreferredAdminLangcode(iPreferredAdminLangcode)
		var m = map[string]interface{}{}
		m["users_field_data"] = _UsersFieldData
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasUsersFieldDataViaPreferredAdminLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasUsersFieldDataViaNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iName := self.Args("name").String()
	if helper.IsHas(iName) {
		_UsersFieldData := model.HasUsersFieldDataViaName(iName)
		var m = map[string]interface{}{}
		m["users_field_data"] = _UsersFieldData
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasUsersFieldDataViaName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasUsersFieldDataViaPassHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iPass := self.Args("pass").String()
	if helper.IsHas(iPass) {
		_UsersFieldData := model.HasUsersFieldDataViaPass(iPass)
		var m = map[string]interface{}{}
		m["users_field_data"] = _UsersFieldData
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasUsersFieldDataViaPass's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasUsersFieldDataViaMailHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iMail := self.Args("mail").String()
	if helper.IsHas(iMail) {
		_UsersFieldData := model.HasUsersFieldDataViaMail(iMail)
		var m = map[string]interface{}{}
		m["users_field_data"] = _UsersFieldData
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasUsersFieldDataViaMail's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasUsersFieldDataViaTimezoneHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iTimezone := self.Args("timezone").String()
	if helper.IsHas(iTimezone) {
		_UsersFieldData := model.HasUsersFieldDataViaTimezone(iTimezone)
		var m = map[string]interface{}{}
		m["users_field_data"] = _UsersFieldData
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasUsersFieldDataViaTimezone's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasUsersFieldDataViaStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iStatus := self.Args("status").MustInt()
	if helper.IsHas(iStatus) {
		_UsersFieldData := model.HasUsersFieldDataViaStatus(iStatus)
		var m = map[string]interface{}{}
		m["users_field_data"] = _UsersFieldData
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasUsersFieldDataViaStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasUsersFieldDataViaCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iCreated := self.Args("created").MustInt()
	if helper.IsHas(iCreated) {
		_UsersFieldData := model.HasUsersFieldDataViaCreated(iCreated)
		var m = map[string]interface{}{}
		m["users_field_data"] = _UsersFieldData
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasUsersFieldDataViaCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasUsersFieldDataViaChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iChanged := self.Args("changed").MustInt()
	if helper.IsHas(iChanged) {
		_UsersFieldData := model.HasUsersFieldDataViaChanged(iChanged)
		var m = map[string]interface{}{}
		m["users_field_data"] = _UsersFieldData
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasUsersFieldDataViaChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasUsersFieldDataViaAccessHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iAccess := self.Args("access").MustInt()
	if helper.IsHas(iAccess) {
		_UsersFieldData := model.HasUsersFieldDataViaAccess(iAccess)
		var m = map[string]interface{}{}
		m["users_field_data"] = _UsersFieldData
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasUsersFieldDataViaAccess's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasUsersFieldDataViaLoginHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iLogin := self.Args("login").MustInt()
	if helper.IsHas(iLogin) {
		_UsersFieldData := model.HasUsersFieldDataViaLogin(iLogin)
		var m = map[string]interface{}{}
		m["users_field_data"] = _UsersFieldData
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasUsersFieldDataViaLogin's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasUsersFieldDataViaInitHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iInit := self.Args("init").String()
	if helper.IsHas(iInit) {
		_UsersFieldData := model.HasUsersFieldDataViaInit(iInit)
		var m = map[string]interface{}{}
		m["users_field_data"] = _UsersFieldData
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasUsersFieldDataViaInit's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasUsersFieldDataViaDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iDefaultLangcode := self.Args("default_langcode").MustInt()
	if helper.IsHas(iDefaultLangcode) {
		_UsersFieldData := model.HasUsersFieldDataViaDefaultLangcode(iDefaultLangcode)
		var m = map[string]interface{}{}
		m["users_field_data"] = _UsersFieldData
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasUsersFieldDataViaDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDataViaUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iUid := self.Args("uid").MustInt64()
	if helper.IsHas(iUid) {
		_UsersFieldData, _error := model.GetUsersFieldDataViaUid(iUid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDataViaUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDataViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iLangcode := self.Args("langcode").String()
	if helper.IsHas(iLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDataViaLangcode(iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDataViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDataViaPreferredLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iPreferredLangcode := self.Args("preferred_langcode").String()
	if helper.IsHas(iPreferredLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDataViaPreferredLangcode(iPreferredLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDataViaPreferredLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDataViaPreferredAdminLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iPreferredAdminLangcode := self.Args("preferred_admin_langcode").String()
	if helper.IsHas(iPreferredAdminLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDataViaPreferredAdminLangcode(iPreferredAdminLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDataViaPreferredAdminLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDataViaNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iName := self.Args("name").String()
	if helper.IsHas(iName) {
		_UsersFieldData, _error := model.GetUsersFieldDataViaName(iName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDataViaName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDataViaPassHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iPass := self.Args("pass").String()
	if helper.IsHas(iPass) {
		_UsersFieldData, _error := model.GetUsersFieldDataViaPass(iPass)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDataViaPass's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDataViaMailHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iMail := self.Args("mail").String()
	if helper.IsHas(iMail) {
		_UsersFieldData, _error := model.GetUsersFieldDataViaMail(iMail)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDataViaMail's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDataViaTimezoneHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iTimezone := self.Args("timezone").String()
	if helper.IsHas(iTimezone) {
		_UsersFieldData, _error := model.GetUsersFieldDataViaTimezone(iTimezone)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDataViaTimezone's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDataViaStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iStatus := self.Args("status").MustInt()
	if helper.IsHas(iStatus) {
		_UsersFieldData, _error := model.GetUsersFieldDataViaStatus(iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDataViaStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDataViaCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iCreated := self.Args("created").MustInt()
	if helper.IsHas(iCreated) {
		_UsersFieldData, _error := model.GetUsersFieldDataViaCreated(iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDataViaCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDataViaChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iChanged := self.Args("changed").MustInt()
	if helper.IsHas(iChanged) {
		_UsersFieldData, _error := model.GetUsersFieldDataViaChanged(iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDataViaChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDataViaAccessHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iAccess := self.Args("access").MustInt()
	if helper.IsHas(iAccess) {
		_UsersFieldData, _error := model.GetUsersFieldDataViaAccess(iAccess)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDataViaAccess's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDataViaLoginHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iLogin := self.Args("login").MustInt()
	if helper.IsHas(iLogin) {
		_UsersFieldData, _error := model.GetUsersFieldDataViaLogin(iLogin)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDataViaLogin's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDataViaInitHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iInit := self.Args("init").String()
	if helper.IsHas(iInit) {
		_UsersFieldData, _error := model.GetUsersFieldDataViaInit(iInit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDataViaInit's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersFieldDataViaDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iDefaultLangcode := self.Args("default_langcode").MustInt()
	if helper.IsHas(iDefaultLangcode) {
		_UsersFieldData, _error := model.GetUsersFieldDataViaDefaultLangcode(iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the GetUsersFieldDataViaDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetUsersFieldDataViaUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Uid_ := self.Args("uid").MustInt64()
	if helper.IsHas(Uid_) {
		var iUsersFieldData model.UsersFieldData
		self.Bind(&iUsersFieldData)
		_UsersFieldData, _error := model.SetUsersFieldDataViaUid(Uid_, &iUsersFieldData)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the SetUsersFieldDataViaUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetUsersFieldDataViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	if helper.IsHas(Langcode_) {
		var iUsersFieldData model.UsersFieldData
		self.Bind(&iUsersFieldData)
		_UsersFieldData, _error := model.SetUsersFieldDataViaLangcode(Langcode_, &iUsersFieldData)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the SetUsersFieldDataViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetUsersFieldDataViaPreferredLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	PreferredLangcode_ := self.Args("preferred_langcode").String()
	if helper.IsHas(PreferredLangcode_) {
		var iUsersFieldData model.UsersFieldData
		self.Bind(&iUsersFieldData)
		_UsersFieldData, _error := model.SetUsersFieldDataViaPreferredLangcode(PreferredLangcode_, &iUsersFieldData)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the SetUsersFieldDataViaPreferredLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetUsersFieldDataViaPreferredAdminLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	PreferredAdminLangcode_ := self.Args("preferred_admin_langcode").String()
	if helper.IsHas(PreferredAdminLangcode_) {
		var iUsersFieldData model.UsersFieldData
		self.Bind(&iUsersFieldData)
		_UsersFieldData, _error := model.SetUsersFieldDataViaPreferredAdminLangcode(PreferredAdminLangcode_, &iUsersFieldData)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the SetUsersFieldDataViaPreferredAdminLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetUsersFieldDataViaNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Name_ := self.Args("name").String()
	if helper.IsHas(Name_) {
		var iUsersFieldData model.UsersFieldData
		self.Bind(&iUsersFieldData)
		_UsersFieldData, _error := model.SetUsersFieldDataViaName(Name_, &iUsersFieldData)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the SetUsersFieldDataViaName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetUsersFieldDataViaPassHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Pass_ := self.Args("pass").String()
	if helper.IsHas(Pass_) {
		var iUsersFieldData model.UsersFieldData
		self.Bind(&iUsersFieldData)
		_UsersFieldData, _error := model.SetUsersFieldDataViaPass(Pass_, &iUsersFieldData)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the SetUsersFieldDataViaPass's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetUsersFieldDataViaMailHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Mail_ := self.Args("mail").String()
	if helper.IsHas(Mail_) {
		var iUsersFieldData model.UsersFieldData
		self.Bind(&iUsersFieldData)
		_UsersFieldData, _error := model.SetUsersFieldDataViaMail(Mail_, &iUsersFieldData)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the SetUsersFieldDataViaMail's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetUsersFieldDataViaTimezoneHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Timezone_ := self.Args("timezone").String()
	if helper.IsHas(Timezone_) {
		var iUsersFieldData model.UsersFieldData
		self.Bind(&iUsersFieldData)
		_UsersFieldData, _error := model.SetUsersFieldDataViaTimezone(Timezone_, &iUsersFieldData)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the SetUsersFieldDataViaTimezone's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetUsersFieldDataViaStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Status_ := self.Args("status").MustInt()
	if helper.IsHas(Status_) {
		var iUsersFieldData model.UsersFieldData
		self.Bind(&iUsersFieldData)
		_UsersFieldData, _error := model.SetUsersFieldDataViaStatus(Status_, &iUsersFieldData)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the SetUsersFieldDataViaStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetUsersFieldDataViaCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Created_ := self.Args("created").MustInt()
	if helper.IsHas(Created_) {
		var iUsersFieldData model.UsersFieldData
		self.Bind(&iUsersFieldData)
		_UsersFieldData, _error := model.SetUsersFieldDataViaCreated(Created_, &iUsersFieldData)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the SetUsersFieldDataViaCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetUsersFieldDataViaChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Changed_ := self.Args("changed").MustInt()
	if helper.IsHas(Changed_) {
		var iUsersFieldData model.UsersFieldData
		self.Bind(&iUsersFieldData)
		_UsersFieldData, _error := model.SetUsersFieldDataViaChanged(Changed_, &iUsersFieldData)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the SetUsersFieldDataViaChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetUsersFieldDataViaAccessHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Access_ := self.Args("access").MustInt()
	if helper.IsHas(Access_) {
		var iUsersFieldData model.UsersFieldData
		self.Bind(&iUsersFieldData)
		_UsersFieldData, _error := model.SetUsersFieldDataViaAccess(Access_, &iUsersFieldData)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the SetUsersFieldDataViaAccess's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetUsersFieldDataViaLoginHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Login_ := self.Args("login").MustInt()
	if helper.IsHas(Login_) {
		var iUsersFieldData model.UsersFieldData
		self.Bind(&iUsersFieldData)
		_UsersFieldData, _error := model.SetUsersFieldDataViaLogin(Login_, &iUsersFieldData)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the SetUsersFieldDataViaLogin's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetUsersFieldDataViaInitHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Init_ := self.Args("init").String()
	if helper.IsHas(Init_) {
		var iUsersFieldData model.UsersFieldData
		self.Bind(&iUsersFieldData)
		_UsersFieldData, _error := model.SetUsersFieldDataViaInit(Init_, &iUsersFieldData)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the SetUsersFieldDataViaInit's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetUsersFieldDataViaDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	DefaultLangcode_ := self.Args("default_langcode").MustInt()
	if helper.IsHas(DefaultLangcode_) {
		var iUsersFieldData model.UsersFieldData
		self.Bind(&iUsersFieldData)
		_UsersFieldData, _error := model.SetUsersFieldDataViaDefaultLangcode(DefaultLangcode_, &iUsersFieldData)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersFieldData)
	}
	herr.Message = "Can't get to the SetUsersFieldDataViaDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostAddUsersFieldDataHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Uid_ := self.Args("uid").MustInt64()
	Langcode_ := self.Args("langcode").String()
	PreferredLangcode_ := self.Args("preferred_langcode").String()
	PreferredAdminLangcode_ := self.Args("preferred_admin_langcode").String()
	Name_ := self.Args("name").String()
	Pass_ := self.Args("pass").String()
	Mail_ := self.Args("mail").String()
	Timezone_ := self.Args("timezone").String()
	Status_ := self.Args("status").MustInt()
	Created_ := self.Args("created").MustInt()
	Changed_ := self.Args("changed").MustInt()
	Access_ := self.Args("access").MustInt()
	Login_ := self.Args("login").MustInt()
	Init_ := self.Args("init").String()
	DefaultLangcode_ := self.Args("default_langcode").MustInt()

	if helper.IsHas(Uid_) {
		_error := model.AddUsersFieldData(Uid_,Langcode_,PreferredLangcode_,PreferredAdminLangcode_,Name_,Pass_,Mail_,Timezone_,Status_,Created_,Changed_,Access_,Login_,Init_,DefaultLangcode_)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		herr.Message = "StatusOK"
		herr.Status = macross.StatusOK
		return self.JSON(herr)
	}
	herr.Message = "Can't get to the AddUsersFieldData's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostUsersFieldDataHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iUsersFieldData model.UsersFieldData
	self.Bind(&iUsersFieldData)
	_int64, _error := model.PostUsersFieldData(&iUsersFieldData)
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

func PutUsersFieldDataHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iUsersFieldData model.UsersFieldData
	self.Bind(&iUsersFieldData)
	_int64, _error := model.PutUsersFieldData(&iUsersFieldData)
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

func PutUsersFieldDataViaUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Uid_ := self.Args("uid").MustInt64()
	var iUsersFieldData model.UsersFieldData
	self.Bind(&iUsersFieldData)
	_int64, _error := model.PutUsersFieldDataViaUid(Uid_, &iUsersFieldData)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUsersFieldDataViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	var iUsersFieldData model.UsersFieldData
	self.Bind(&iUsersFieldData)
	_int64, _error := model.PutUsersFieldDataViaLangcode(Langcode_, &iUsersFieldData)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUsersFieldDataViaPreferredLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	PreferredLangcode_ := self.Args("preferred_langcode").String()
	var iUsersFieldData model.UsersFieldData
	self.Bind(&iUsersFieldData)
	_int64, _error := model.PutUsersFieldDataViaPreferredLangcode(PreferredLangcode_, &iUsersFieldData)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUsersFieldDataViaPreferredAdminLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	PreferredAdminLangcode_ := self.Args("preferred_admin_langcode").String()
	var iUsersFieldData model.UsersFieldData
	self.Bind(&iUsersFieldData)
	_int64, _error := model.PutUsersFieldDataViaPreferredAdminLangcode(PreferredAdminLangcode_, &iUsersFieldData)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUsersFieldDataViaNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Name_ := self.Args("name").String()
	var iUsersFieldData model.UsersFieldData
	self.Bind(&iUsersFieldData)
	_int64, _error := model.PutUsersFieldDataViaName(Name_, &iUsersFieldData)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUsersFieldDataViaPassHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Pass_ := self.Args("pass").String()
	var iUsersFieldData model.UsersFieldData
	self.Bind(&iUsersFieldData)
	_int64, _error := model.PutUsersFieldDataViaPass(Pass_, &iUsersFieldData)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUsersFieldDataViaMailHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Mail_ := self.Args("mail").String()
	var iUsersFieldData model.UsersFieldData
	self.Bind(&iUsersFieldData)
	_int64, _error := model.PutUsersFieldDataViaMail(Mail_, &iUsersFieldData)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUsersFieldDataViaTimezoneHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Timezone_ := self.Args("timezone").String()
	var iUsersFieldData model.UsersFieldData
	self.Bind(&iUsersFieldData)
	_int64, _error := model.PutUsersFieldDataViaTimezone(Timezone_, &iUsersFieldData)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUsersFieldDataViaStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Status_ := self.Args("status").MustInt()
	var iUsersFieldData model.UsersFieldData
	self.Bind(&iUsersFieldData)
	_int64, _error := model.PutUsersFieldDataViaStatus(Status_, &iUsersFieldData)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUsersFieldDataViaCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Created_ := self.Args("created").MustInt()
	var iUsersFieldData model.UsersFieldData
	self.Bind(&iUsersFieldData)
	_int64, _error := model.PutUsersFieldDataViaCreated(Created_, &iUsersFieldData)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUsersFieldDataViaChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Changed_ := self.Args("changed").MustInt()
	var iUsersFieldData model.UsersFieldData
	self.Bind(&iUsersFieldData)
	_int64, _error := model.PutUsersFieldDataViaChanged(Changed_, &iUsersFieldData)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUsersFieldDataViaAccessHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Access_ := self.Args("access").MustInt()
	var iUsersFieldData model.UsersFieldData
	self.Bind(&iUsersFieldData)
	_int64, _error := model.PutUsersFieldDataViaAccess(Access_, &iUsersFieldData)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUsersFieldDataViaLoginHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Login_ := self.Args("login").MustInt()
	var iUsersFieldData model.UsersFieldData
	self.Bind(&iUsersFieldData)
	_int64, _error := model.PutUsersFieldDataViaLogin(Login_, &iUsersFieldData)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUsersFieldDataViaInitHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Init_ := self.Args("init").String()
	var iUsersFieldData model.UsersFieldData
	self.Bind(&iUsersFieldData)
	_int64, _error := model.PutUsersFieldDataViaInit(Init_, &iUsersFieldData)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUsersFieldDataViaDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	DefaultLangcode_ := self.Args("default_langcode").MustInt()
	var iUsersFieldData model.UsersFieldData
	self.Bind(&iUsersFieldData)
	_int64, _error := model.PutUsersFieldDataViaDefaultLangcode(DefaultLangcode_, &iUsersFieldData)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateUsersFieldDataViaUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Uid_ := self.Args("uid").MustInt64()
	var iUsersFieldData model.UsersFieldData
	self.Bind(&iUsersFieldData)
	var iMap = helper.StructToMap(iUsersFieldData)
	_error := model.UpdateUsersFieldDataViaUid(Uid_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateUsersFieldDataViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	var iUsersFieldData model.UsersFieldData
	self.Bind(&iUsersFieldData)
	var iMap = helper.StructToMap(iUsersFieldData)
	_error := model.UpdateUsersFieldDataViaLangcode(Langcode_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateUsersFieldDataViaPreferredLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	PreferredLangcode_ := self.Args("preferred_langcode").String()
	var iUsersFieldData model.UsersFieldData
	self.Bind(&iUsersFieldData)
	var iMap = helper.StructToMap(iUsersFieldData)
	_error := model.UpdateUsersFieldDataViaPreferredLangcode(PreferredLangcode_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateUsersFieldDataViaPreferredAdminLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	PreferredAdminLangcode_ := self.Args("preferred_admin_langcode").String()
	var iUsersFieldData model.UsersFieldData
	self.Bind(&iUsersFieldData)
	var iMap = helper.StructToMap(iUsersFieldData)
	_error := model.UpdateUsersFieldDataViaPreferredAdminLangcode(PreferredAdminLangcode_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateUsersFieldDataViaNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Name_ := self.Args("name").String()
	var iUsersFieldData model.UsersFieldData
	self.Bind(&iUsersFieldData)
	var iMap = helper.StructToMap(iUsersFieldData)
	_error := model.UpdateUsersFieldDataViaName(Name_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateUsersFieldDataViaPassHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Pass_ := self.Args("pass").String()
	var iUsersFieldData model.UsersFieldData
	self.Bind(&iUsersFieldData)
	var iMap = helper.StructToMap(iUsersFieldData)
	_error := model.UpdateUsersFieldDataViaPass(Pass_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateUsersFieldDataViaMailHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Mail_ := self.Args("mail").String()
	var iUsersFieldData model.UsersFieldData
	self.Bind(&iUsersFieldData)
	var iMap = helper.StructToMap(iUsersFieldData)
	_error := model.UpdateUsersFieldDataViaMail(Mail_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateUsersFieldDataViaTimezoneHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Timezone_ := self.Args("timezone").String()
	var iUsersFieldData model.UsersFieldData
	self.Bind(&iUsersFieldData)
	var iMap = helper.StructToMap(iUsersFieldData)
	_error := model.UpdateUsersFieldDataViaTimezone(Timezone_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateUsersFieldDataViaStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Status_ := self.Args("status").MustInt()
	var iUsersFieldData model.UsersFieldData
	self.Bind(&iUsersFieldData)
	var iMap = helper.StructToMap(iUsersFieldData)
	_error := model.UpdateUsersFieldDataViaStatus(Status_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateUsersFieldDataViaCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Created_ := self.Args("created").MustInt()
	var iUsersFieldData model.UsersFieldData
	self.Bind(&iUsersFieldData)
	var iMap = helper.StructToMap(iUsersFieldData)
	_error := model.UpdateUsersFieldDataViaCreated(Created_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateUsersFieldDataViaChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Changed_ := self.Args("changed").MustInt()
	var iUsersFieldData model.UsersFieldData
	self.Bind(&iUsersFieldData)
	var iMap = helper.StructToMap(iUsersFieldData)
	_error := model.UpdateUsersFieldDataViaChanged(Changed_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateUsersFieldDataViaAccessHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Access_ := self.Args("access").MustInt()
	var iUsersFieldData model.UsersFieldData
	self.Bind(&iUsersFieldData)
	var iMap = helper.StructToMap(iUsersFieldData)
	_error := model.UpdateUsersFieldDataViaAccess(Access_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateUsersFieldDataViaLoginHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Login_ := self.Args("login").MustInt()
	var iUsersFieldData model.UsersFieldData
	self.Bind(&iUsersFieldData)
	var iMap = helper.StructToMap(iUsersFieldData)
	_error := model.UpdateUsersFieldDataViaLogin(Login_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateUsersFieldDataViaInitHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Init_ := self.Args("init").String()
	var iUsersFieldData model.UsersFieldData
	self.Bind(&iUsersFieldData)
	var iMap = helper.StructToMap(iUsersFieldData)
	_error := model.UpdateUsersFieldDataViaInit(Init_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateUsersFieldDataViaDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	DefaultLangcode_ := self.Args("default_langcode").MustInt()
	var iUsersFieldData model.UsersFieldData
	self.Bind(&iUsersFieldData)
	var iMap = helper.StructToMap(iUsersFieldData)
	_error := model.UpdateUsersFieldDataViaDefaultLangcode(DefaultLangcode_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteUsersFieldDataHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Uid_ := self.Args("uid").MustInt64()
	_int64, _error := model.DeleteUsersFieldData(Uid_)
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

func DeleteUsersFieldDataViaUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Uid_ := self.Args("uid").MustInt64()
	_error := model.DeleteUsersFieldDataViaUid(Uid_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteUsersFieldDataViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	_error := model.DeleteUsersFieldDataViaLangcode(Langcode_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteUsersFieldDataViaPreferredLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	PreferredLangcode_ := self.Args("preferred_langcode").String()
	_error := model.DeleteUsersFieldDataViaPreferredLangcode(PreferredLangcode_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteUsersFieldDataViaPreferredAdminLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	PreferredAdminLangcode_ := self.Args("preferred_admin_langcode").String()
	_error := model.DeleteUsersFieldDataViaPreferredAdminLangcode(PreferredAdminLangcode_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteUsersFieldDataViaNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Name_ := self.Args("name").String()
	_error := model.DeleteUsersFieldDataViaName(Name_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteUsersFieldDataViaPassHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Pass_ := self.Args("pass").String()
	_error := model.DeleteUsersFieldDataViaPass(Pass_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteUsersFieldDataViaMailHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Mail_ := self.Args("mail").String()
	_error := model.DeleteUsersFieldDataViaMail(Mail_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteUsersFieldDataViaTimezoneHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Timezone_ := self.Args("timezone").String()
	_error := model.DeleteUsersFieldDataViaTimezone(Timezone_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteUsersFieldDataViaStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Status_ := self.Args("status").MustInt()
	_error := model.DeleteUsersFieldDataViaStatus(Status_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteUsersFieldDataViaCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Created_ := self.Args("created").MustInt()
	_error := model.DeleteUsersFieldDataViaCreated(Created_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteUsersFieldDataViaChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Changed_ := self.Args("changed").MustInt()
	_error := model.DeleteUsersFieldDataViaChanged(Changed_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteUsersFieldDataViaAccessHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Access_ := self.Args("access").MustInt()
	_error := model.DeleteUsersFieldDataViaAccess(Access_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteUsersFieldDataViaLoginHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Login_ := self.Args("login").MustInt()
	_error := model.DeleteUsersFieldDataViaLogin(Login_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteUsersFieldDataViaInitHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Init_ := self.Args("init").String()
	_error := model.DeleteUsersFieldDataViaInit(Init_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteUsersFieldDataViaDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	DefaultLangcode_ := self.Args("default_langcode").MustInt()
	_error := model.DeleteUsersFieldDataViaDefaultLangcode(DefaultLangcode_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}
