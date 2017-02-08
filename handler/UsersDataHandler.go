package handler

import (
	"drupapp/helper"
	"drupapp/model"
	"github.com/insionng/macross"
)

func GetUsersDatasCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()

	if offset > 0 {
		var m = map[string]interface{}{}
		_int64, _error := model.GetUsersDatasCount(offset, limit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		m["users_datasCount"] = _int64
		return self.JSON(m)
	}
	herr.Message = "Can't get to the GetUsersDatasCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersDataCountViaUidHandler(self *macross.Context) error {
	Uid_ := self.Args("uid").MustInt64()
	_int64 := model.GetUsersDataCountViaUid(Uid_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["users_dataCount"] = 0
	}
	m["users_dataCount"] = _int64
	return self.JSON(m)
}

func GetUsersDataCountViaModuleHandler(self *macross.Context) error {
	Module_ := self.Args("module").String()
	_int64 := model.GetUsersDataCountViaModule(Module_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["users_dataCount"] = 0
	}
	m["users_dataCount"] = _int64
	return self.JSON(m)
}

func GetUsersDataCountViaNameHandler(self *macross.Context) error {
	Name_ := self.Args("name").String()
	_int64 := model.GetUsersDataCountViaName(Name_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["users_dataCount"] = 0
	}
	m["users_dataCount"] = _int64
	return self.JSON(m)
}

func GetUsersDataCountViaValueHandler(self *macross.Context) error {
	Value_ := self.Args("value").Bytes()
	_int64 := model.GetUsersDataCountViaValue(Value_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["users_dataCount"] = 0
	}
	m["users_dataCount"] = _int64
	return self.JSON(m)
}

func GetUsersDataCountViaSerializedHandler(self *macross.Context) error {
	Serialized_ := self.Args("serialized").MustInt()
	_int64 := model.GetUsersDataCountViaSerialized(Serialized_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["users_dataCount"] = 0
	}
	m["users_dataCount"] = _int64
	return self.JSON(m)
}

func GetUsersDatasViaUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iUid := self.Args("uid").MustInt64()
	if (offset > 0) && helper.IsHas(iUid) {
		_UsersData, _error := model.GetUsersDatasViaUid(offset, limit, iUid, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersData)
	}
	herr.Message = "Can't get to the GetUsersDatasViaUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersDatasViaModuleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iModule := self.Args("module").String()
	if (offset > 0) && helper.IsHas(iModule) {
		_UsersData, _error := model.GetUsersDatasViaModule(offset, limit, iModule, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersData)
	}
	herr.Message = "Can't get to the GetUsersDatasViaModule's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersDatasViaNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iName := self.Args("name").String()
	if (offset > 0) && helper.IsHas(iName) {
		_UsersData, _error := model.GetUsersDatasViaName(offset, limit, iName, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersData)
	}
	herr.Message = "Can't get to the GetUsersDatasViaName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersDatasViaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iValue := self.Args("value").Bytes()
	if (offset > 0) && helper.IsHas(iValue) {
		_UsersData, _error := model.GetUsersDatasViaValue(offset, limit, iValue, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersData)
	}
	herr.Message = "Can't get to the GetUsersDatasViaValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersDatasViaSerializedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iSerialized := self.Args("serialized").MustInt()
	if (offset > 0) && helper.IsHas(iSerialized) {
		_UsersData, _error := model.GetUsersDatasViaSerialized(offset, limit, iSerialized, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersData)
	}
	herr.Message = "Can't get to the GetUsersDatasViaSerialized's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersDatasByUidAndModuleAndNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt64()
	iModule := self.Args("module").String()
	iName := self.Args("name").String()

	if helper.IsHas(iUid) {
		_UsersData, _error := model.GetUsersDatasByUidAndModuleAndName(offset, limit, iUid,iModule,iName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersData)
	}
	herr.Message = "Can't get to the GetUsersDatasByUidAndModuleAndName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersDatasByUidAndModuleAndValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt64()
	iModule := self.Args("module").String()
	iValue := self.Args("value").Bytes()

	if helper.IsHas(iUid) {
		_UsersData, _error := model.GetUsersDatasByUidAndModuleAndValue(offset, limit, iUid,iModule,iValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersData)
	}
	herr.Message = "Can't get to the GetUsersDatasByUidAndModuleAndValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersDatasByUidAndModuleAndSerializedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt64()
	iModule := self.Args("module").String()
	iSerialized := self.Args("serialized").MustInt()

	if helper.IsHas(iUid) {
		_UsersData, _error := model.GetUsersDatasByUidAndModuleAndSerialized(offset, limit, iUid,iModule,iSerialized)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersData)
	}
	herr.Message = "Can't get to the GetUsersDatasByUidAndModuleAndSerialized's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersDatasByUidAndNameAndValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt64()
	iName := self.Args("name").String()
	iValue := self.Args("value").Bytes()

	if helper.IsHas(iUid) {
		_UsersData, _error := model.GetUsersDatasByUidAndNameAndValue(offset, limit, iUid,iName,iValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersData)
	}
	herr.Message = "Can't get to the GetUsersDatasByUidAndNameAndValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersDatasByUidAndNameAndSerializedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt64()
	iName := self.Args("name").String()
	iSerialized := self.Args("serialized").MustInt()

	if helper.IsHas(iUid) {
		_UsersData, _error := model.GetUsersDatasByUidAndNameAndSerialized(offset, limit, iUid,iName,iSerialized)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersData)
	}
	herr.Message = "Can't get to the GetUsersDatasByUidAndNameAndSerialized's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersDatasByUidAndValueAndSerializedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt64()
	iValue := self.Args("value").Bytes()
	iSerialized := self.Args("serialized").MustInt()

	if helper.IsHas(iUid) {
		_UsersData, _error := model.GetUsersDatasByUidAndValueAndSerialized(offset, limit, iUid,iValue,iSerialized)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersData)
	}
	herr.Message = "Can't get to the GetUsersDatasByUidAndValueAndSerialized's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersDatasByModuleAndNameAndValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iModule := self.Args("module").String()
	iName := self.Args("name").String()
	iValue := self.Args("value").Bytes()

	if helper.IsHas(iModule) {
		_UsersData, _error := model.GetUsersDatasByModuleAndNameAndValue(offset, limit, iModule,iName,iValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersData)
	}
	herr.Message = "Can't get to the GetUsersDatasByModuleAndNameAndValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersDatasByModuleAndNameAndSerializedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iModule := self.Args("module").String()
	iName := self.Args("name").String()
	iSerialized := self.Args("serialized").MustInt()

	if helper.IsHas(iModule) {
		_UsersData, _error := model.GetUsersDatasByModuleAndNameAndSerialized(offset, limit, iModule,iName,iSerialized)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersData)
	}
	herr.Message = "Can't get to the GetUsersDatasByModuleAndNameAndSerialized's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersDatasByModuleAndValueAndSerializedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iModule := self.Args("module").String()
	iValue := self.Args("value").Bytes()
	iSerialized := self.Args("serialized").MustInt()

	if helper.IsHas(iModule) {
		_UsersData, _error := model.GetUsersDatasByModuleAndValueAndSerialized(offset, limit, iModule,iValue,iSerialized)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersData)
	}
	herr.Message = "Can't get to the GetUsersDatasByModuleAndValueAndSerialized's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersDatasByNameAndValueAndSerializedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iValue := self.Args("value").Bytes()
	iSerialized := self.Args("serialized").MustInt()

	if helper.IsHas(iName) {
		_UsersData, _error := model.GetUsersDatasByNameAndValueAndSerialized(offset, limit, iName,iValue,iSerialized)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersData)
	}
	herr.Message = "Can't get to the GetUsersDatasByNameAndValueAndSerialized's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersDatasByUidAndModuleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt64()
	iModule := self.Args("module").String()

	if helper.IsHas(iUid) {
		_UsersData, _error := model.GetUsersDatasByUidAndModule(offset, limit, iUid,iModule)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersData)
	}
	herr.Message = "Can't get to the GetUsersDatasByUidAndModule's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersDatasByUidAndNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt64()
	iName := self.Args("name").String()

	if helper.IsHas(iUid) {
		_UsersData, _error := model.GetUsersDatasByUidAndName(offset, limit, iUid,iName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersData)
	}
	herr.Message = "Can't get to the GetUsersDatasByUidAndName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersDatasByUidAndValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt64()
	iValue := self.Args("value").Bytes()

	if helper.IsHas(iUid) {
		_UsersData, _error := model.GetUsersDatasByUidAndValue(offset, limit, iUid,iValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersData)
	}
	herr.Message = "Can't get to the GetUsersDatasByUidAndValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersDatasByUidAndSerializedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt64()
	iSerialized := self.Args("serialized").MustInt()

	if helper.IsHas(iUid) {
		_UsersData, _error := model.GetUsersDatasByUidAndSerialized(offset, limit, iUid,iSerialized)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersData)
	}
	herr.Message = "Can't get to the GetUsersDatasByUidAndSerialized's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersDatasByModuleAndNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iModule := self.Args("module").String()
	iName := self.Args("name").String()

	if helper.IsHas(iModule) {
		_UsersData, _error := model.GetUsersDatasByModuleAndName(offset, limit, iModule,iName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersData)
	}
	herr.Message = "Can't get to the GetUsersDatasByModuleAndName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersDatasByModuleAndValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iModule := self.Args("module").String()
	iValue := self.Args("value").Bytes()

	if helper.IsHas(iModule) {
		_UsersData, _error := model.GetUsersDatasByModuleAndValue(offset, limit, iModule,iValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersData)
	}
	herr.Message = "Can't get to the GetUsersDatasByModuleAndValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersDatasByModuleAndSerializedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iModule := self.Args("module").String()
	iSerialized := self.Args("serialized").MustInt()

	if helper.IsHas(iModule) {
		_UsersData, _error := model.GetUsersDatasByModuleAndSerialized(offset, limit, iModule,iSerialized)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersData)
	}
	herr.Message = "Can't get to the GetUsersDatasByModuleAndSerialized's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersDatasByNameAndValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iValue := self.Args("value").Bytes()

	if helper.IsHas(iName) {
		_UsersData, _error := model.GetUsersDatasByNameAndValue(offset, limit, iName,iValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersData)
	}
	herr.Message = "Can't get to the GetUsersDatasByNameAndValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersDatasByNameAndSerializedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iSerialized := self.Args("serialized").MustInt()

	if helper.IsHas(iName) {
		_UsersData, _error := model.GetUsersDatasByNameAndSerialized(offset, limit, iName,iSerialized)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersData)
	}
	herr.Message = "Can't get to the GetUsersDatasByNameAndSerialized's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersDatasByValueAndSerializedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iValue := self.Args("value").Bytes()
	iSerialized := self.Args("serialized").MustInt()

	if helper.IsHas(iValue) {
		_UsersData, _error := model.GetUsersDatasByValueAndSerialized(offset, limit, iValue,iSerialized)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersData)
	}
	herr.Message = "Can't get to the GetUsersDatasByValueAndSerialized's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersDatasHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	if (offset > 0) && (len(field) > 0) {
		_UsersData, _error := model.GetUsersDatas(offset, limit, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersData)
	}
	herr.Message = "Can't get to the GetUsersDatas' args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasUsersDataViaUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iUid := self.Args("uid").MustInt64()
	if helper.IsHas(iUid) {
		_UsersData := model.HasUsersDataViaUid(iUid)
		var m = map[string]interface{}{}
		m["users_data"] = _UsersData
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasUsersDataViaUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasUsersDataViaModuleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iModule := self.Args("module").String()
	if helper.IsHas(iModule) {
		_UsersData := model.HasUsersDataViaModule(iModule)
		var m = map[string]interface{}{}
		m["users_data"] = _UsersData
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasUsersDataViaModule's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasUsersDataViaNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iName := self.Args("name").String()
	if helper.IsHas(iName) {
		_UsersData := model.HasUsersDataViaName(iName)
		var m = map[string]interface{}{}
		m["users_data"] = _UsersData
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasUsersDataViaName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasUsersDataViaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iValue := self.Args("value").Bytes()
	if helper.IsHas(iValue) {
		_UsersData := model.HasUsersDataViaValue(iValue)
		var m = map[string]interface{}{}
		m["users_data"] = _UsersData
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasUsersDataViaValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasUsersDataViaSerializedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iSerialized := self.Args("serialized").MustInt()
	if helper.IsHas(iSerialized) {
		_UsersData := model.HasUsersDataViaSerialized(iSerialized)
		var m = map[string]interface{}{}
		m["users_data"] = _UsersData
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasUsersDataViaSerialized's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersDataViaUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iUid := self.Args("uid").MustInt64()
	if helper.IsHas(iUid) {
		_UsersData, _error := model.GetUsersDataViaUid(iUid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersData)
	}
	herr.Message = "Can't get to the GetUsersDataViaUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersDataViaModuleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iModule := self.Args("module").String()
	if helper.IsHas(iModule) {
		_UsersData, _error := model.GetUsersDataViaModule(iModule)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersData)
	}
	herr.Message = "Can't get to the GetUsersDataViaModule's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersDataViaNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iName := self.Args("name").String()
	if helper.IsHas(iName) {
		_UsersData, _error := model.GetUsersDataViaName(iName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersData)
	}
	herr.Message = "Can't get to the GetUsersDataViaName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersDataViaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iValue := self.Args("value").Bytes()
	if helper.IsHas(iValue) {
		_UsersData, _error := model.GetUsersDataViaValue(iValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersData)
	}
	herr.Message = "Can't get to the GetUsersDataViaValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersDataViaSerializedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iSerialized := self.Args("serialized").MustInt()
	if helper.IsHas(iSerialized) {
		_UsersData, _error := model.GetUsersDataViaSerialized(iSerialized)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersData)
	}
	herr.Message = "Can't get to the GetUsersDataViaSerialized's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetUsersDataViaUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Uid_ := self.Args("uid").MustInt64()
	if helper.IsHas(Uid_) {
		var iUsersData model.UsersData
		self.Bind(&iUsersData)
		_UsersData, _error := model.SetUsersDataViaUid(Uid_, &iUsersData)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersData)
	}
	herr.Message = "Can't get to the SetUsersDataViaUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetUsersDataViaModuleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Module_ := self.Args("module").String()
	if helper.IsHas(Module_) {
		var iUsersData model.UsersData
		self.Bind(&iUsersData)
		_UsersData, _error := model.SetUsersDataViaModule(Module_, &iUsersData)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersData)
	}
	herr.Message = "Can't get to the SetUsersDataViaModule's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetUsersDataViaNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Name_ := self.Args("name").String()
	if helper.IsHas(Name_) {
		var iUsersData model.UsersData
		self.Bind(&iUsersData)
		_UsersData, _error := model.SetUsersDataViaName(Name_, &iUsersData)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersData)
	}
	herr.Message = "Can't get to the SetUsersDataViaName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetUsersDataViaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Value_ := self.Args("value").Bytes()
	if helper.IsHas(Value_) {
		var iUsersData model.UsersData
		self.Bind(&iUsersData)
		_UsersData, _error := model.SetUsersDataViaValue(Value_, &iUsersData)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersData)
	}
	herr.Message = "Can't get to the SetUsersDataViaValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetUsersDataViaSerializedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Serialized_ := self.Args("serialized").MustInt()
	if helper.IsHas(Serialized_) {
		var iUsersData model.UsersData
		self.Bind(&iUsersData)
		_UsersData, _error := model.SetUsersDataViaSerialized(Serialized_, &iUsersData)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UsersData)
	}
	herr.Message = "Can't get to the SetUsersDataViaSerialized's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostAddUsersDataHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Uid_ := self.Args("uid").MustInt64()
	Module_ := self.Args("module").String()
	Name_ := self.Args("name").String()
	Value_ := self.Args("value").Bytes()
	Serialized_ := self.Args("serialized").MustInt()

	if helper.IsHas(Uid_) {
		_error := model.AddUsersData(Uid_,Module_,Name_,Value_,Serialized_)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		herr.Message = "StatusOK"
		herr.Status = macross.StatusOK
		return self.JSON(herr)
	}
	herr.Message = "Can't get to the AddUsersData's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostUsersDataHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iUsersData model.UsersData
	self.Bind(&iUsersData)
	_int64, _error := model.PostUsersData(&iUsersData)
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

func PutUsersDataHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iUsersData model.UsersData
	self.Bind(&iUsersData)
	_int64, _error := model.PutUsersData(&iUsersData)
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

func PutUsersDataViaUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Uid_ := self.Args("uid").MustInt64()
	var iUsersData model.UsersData
	self.Bind(&iUsersData)
	_int64, _error := model.PutUsersDataViaUid(Uid_, &iUsersData)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUsersDataViaModuleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Module_ := self.Args("module").String()
	var iUsersData model.UsersData
	self.Bind(&iUsersData)
	_int64, _error := model.PutUsersDataViaModule(Module_, &iUsersData)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUsersDataViaNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Name_ := self.Args("name").String()
	var iUsersData model.UsersData
	self.Bind(&iUsersData)
	_int64, _error := model.PutUsersDataViaName(Name_, &iUsersData)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUsersDataViaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Value_ := self.Args("value").Bytes()
	var iUsersData model.UsersData
	self.Bind(&iUsersData)
	_int64, _error := model.PutUsersDataViaValue(Value_, &iUsersData)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUsersDataViaSerializedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Serialized_ := self.Args("serialized").MustInt()
	var iUsersData model.UsersData
	self.Bind(&iUsersData)
	_int64, _error := model.PutUsersDataViaSerialized(Serialized_, &iUsersData)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateUsersDataViaUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Uid_ := self.Args("uid").MustInt64()
	var iUsersData model.UsersData
	self.Bind(&iUsersData)
	var iMap = helper.StructToMap(iUsersData)
	_error := model.UpdateUsersDataViaUid(Uid_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateUsersDataViaModuleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Module_ := self.Args("module").String()
	var iUsersData model.UsersData
	self.Bind(&iUsersData)
	var iMap = helper.StructToMap(iUsersData)
	_error := model.UpdateUsersDataViaModule(Module_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateUsersDataViaNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Name_ := self.Args("name").String()
	var iUsersData model.UsersData
	self.Bind(&iUsersData)
	var iMap = helper.StructToMap(iUsersData)
	_error := model.UpdateUsersDataViaName(Name_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateUsersDataViaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Value_ := self.Args("value").Bytes()
	var iUsersData model.UsersData
	self.Bind(&iUsersData)
	var iMap = helper.StructToMap(iUsersData)
	_error := model.UpdateUsersDataViaValue(Value_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateUsersDataViaSerializedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Serialized_ := self.Args("serialized").MustInt()
	var iUsersData model.UsersData
	self.Bind(&iUsersData)
	var iMap = helper.StructToMap(iUsersData)
	_error := model.UpdateUsersDataViaSerialized(Serialized_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteUsersDataHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Uid_ := self.Args("uid").MustInt64()
	_int64, _error := model.DeleteUsersData(Uid_)
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

func DeleteUsersDataViaUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Uid_ := self.Args("uid").MustInt64()
	_error := model.DeleteUsersDataViaUid(Uid_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteUsersDataViaModuleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Module_ := self.Args("module").String()
	_error := model.DeleteUsersDataViaModule(Module_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteUsersDataViaNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Name_ := self.Args("name").String()
	_error := model.DeleteUsersDataViaName(Name_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteUsersDataViaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Value_ := self.Args("value").Bytes()
	_error := model.DeleteUsersDataViaValue(Value_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteUsersDataViaSerializedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Serialized_ := self.Args("serialized").MustInt()
	_error := model.DeleteUsersDataViaSerialized(Serialized_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}
