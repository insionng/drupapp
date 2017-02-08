package handler

import (
	"drupapp/helper"
	"drupapp/model"
	"github.com/insionng/macross"
)

func GetKeyValueExpiresCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()

	if offset > 0 {
		var m = map[string]interface{}{}
		_int64, _error := model.GetKeyValueExpiresCount(offset, limit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		m["key_value_expiresCount"] = _int64
		return self.JSON(m)
	}
	herr.Message = "Can't get to the GetKeyValueExpiresCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetKeyValueExpireCountViaCollectionHandler(self *macross.Context) error {
	Collection_ := self.Args("collection").String()
	_int64 := model.GetKeyValueExpireCountViaCollection(Collection_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["key_value_expireCount"] = 0
	}
	m["key_value_expireCount"] = _int64
	return self.JSON(m)
}

func GetKeyValueExpireCountViaNameHandler(self *macross.Context) error {
	Name_ := self.Args("name").String()
	_int64 := model.GetKeyValueExpireCountViaName(Name_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["key_value_expireCount"] = 0
	}
	m["key_value_expireCount"] = _int64
	return self.JSON(m)
}

func GetKeyValueExpireCountViaValueHandler(self *macross.Context) error {
	Value_ := self.Args("value").Bytes()
	_int64 := model.GetKeyValueExpireCountViaValue(Value_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["key_value_expireCount"] = 0
	}
	m["key_value_expireCount"] = _int64
	return self.JSON(m)
}

func GetKeyValueExpireCountViaExpireHandler(self *macross.Context) error {
	Expire_ := self.Args("expire").MustInt()
	_int64 := model.GetKeyValueExpireCountViaExpire(Expire_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["key_value_expireCount"] = 0
	}
	m["key_value_expireCount"] = _int64
	return self.JSON(m)
}

func GetKeyValueExpiresViaCollectionHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iCollection := self.Args("collection").String()
	if (offset > 0) && helper.IsHas(iCollection) {
		_KeyValueExpire, _error := model.GetKeyValueExpiresViaCollection(offset, limit, iCollection, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_KeyValueExpire)
	}
	herr.Message = "Can't get to the GetKeyValueExpiresViaCollection's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetKeyValueExpiresViaNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iName := self.Args("name").String()
	if (offset > 0) && helper.IsHas(iName) {
		_KeyValueExpire, _error := model.GetKeyValueExpiresViaName(offset, limit, iName, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_KeyValueExpire)
	}
	herr.Message = "Can't get to the GetKeyValueExpiresViaName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetKeyValueExpiresViaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iValue := self.Args("value").Bytes()
	if (offset > 0) && helper.IsHas(iValue) {
		_KeyValueExpire, _error := model.GetKeyValueExpiresViaValue(offset, limit, iValue, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_KeyValueExpire)
	}
	herr.Message = "Can't get to the GetKeyValueExpiresViaValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetKeyValueExpiresViaExpireHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iExpire := self.Args("expire").MustInt()
	if (offset > 0) && helper.IsHas(iExpire) {
		_KeyValueExpire, _error := model.GetKeyValueExpiresViaExpire(offset, limit, iExpire, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_KeyValueExpire)
	}
	herr.Message = "Can't get to the GetKeyValueExpiresViaExpire's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetKeyValueExpiresByCollectionAndNameAndValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCollection := self.Args("collection").String()
	iName := self.Args("name").String()
	iValue := self.Args("value").Bytes()

	if helper.IsHas(iCollection) {
		_KeyValueExpire, _error := model.GetKeyValueExpiresByCollectionAndNameAndValue(offset, limit, iCollection,iName,iValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_KeyValueExpire)
	}
	herr.Message = "Can't get to the GetKeyValueExpiresByCollectionAndNameAndValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetKeyValueExpiresByCollectionAndNameAndExpireHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCollection := self.Args("collection").String()
	iName := self.Args("name").String()
	iExpire := self.Args("expire").MustInt()

	if helper.IsHas(iCollection) {
		_KeyValueExpire, _error := model.GetKeyValueExpiresByCollectionAndNameAndExpire(offset, limit, iCollection,iName,iExpire)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_KeyValueExpire)
	}
	herr.Message = "Can't get to the GetKeyValueExpiresByCollectionAndNameAndExpire's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetKeyValueExpiresByCollectionAndValueAndExpireHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCollection := self.Args("collection").String()
	iValue := self.Args("value").Bytes()
	iExpire := self.Args("expire").MustInt()

	if helper.IsHas(iCollection) {
		_KeyValueExpire, _error := model.GetKeyValueExpiresByCollectionAndValueAndExpire(offset, limit, iCollection,iValue,iExpire)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_KeyValueExpire)
	}
	herr.Message = "Can't get to the GetKeyValueExpiresByCollectionAndValueAndExpire's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetKeyValueExpiresByNameAndValueAndExpireHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iValue := self.Args("value").Bytes()
	iExpire := self.Args("expire").MustInt()

	if helper.IsHas(iName) {
		_KeyValueExpire, _error := model.GetKeyValueExpiresByNameAndValueAndExpire(offset, limit, iName,iValue,iExpire)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_KeyValueExpire)
	}
	herr.Message = "Can't get to the GetKeyValueExpiresByNameAndValueAndExpire's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetKeyValueExpiresByCollectionAndNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCollection := self.Args("collection").String()
	iName := self.Args("name").String()

	if helper.IsHas(iCollection) {
		_KeyValueExpire, _error := model.GetKeyValueExpiresByCollectionAndName(offset, limit, iCollection,iName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_KeyValueExpire)
	}
	herr.Message = "Can't get to the GetKeyValueExpiresByCollectionAndName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetKeyValueExpiresByCollectionAndValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCollection := self.Args("collection").String()
	iValue := self.Args("value").Bytes()

	if helper.IsHas(iCollection) {
		_KeyValueExpire, _error := model.GetKeyValueExpiresByCollectionAndValue(offset, limit, iCollection,iValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_KeyValueExpire)
	}
	herr.Message = "Can't get to the GetKeyValueExpiresByCollectionAndValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetKeyValueExpiresByCollectionAndExpireHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCollection := self.Args("collection").String()
	iExpire := self.Args("expire").MustInt()

	if helper.IsHas(iCollection) {
		_KeyValueExpire, _error := model.GetKeyValueExpiresByCollectionAndExpire(offset, limit, iCollection,iExpire)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_KeyValueExpire)
	}
	herr.Message = "Can't get to the GetKeyValueExpiresByCollectionAndExpire's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetKeyValueExpiresByNameAndValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iValue := self.Args("value").Bytes()

	if helper.IsHas(iName) {
		_KeyValueExpire, _error := model.GetKeyValueExpiresByNameAndValue(offset, limit, iName,iValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_KeyValueExpire)
	}
	herr.Message = "Can't get to the GetKeyValueExpiresByNameAndValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetKeyValueExpiresByNameAndExpireHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iExpire := self.Args("expire").MustInt()

	if helper.IsHas(iName) {
		_KeyValueExpire, _error := model.GetKeyValueExpiresByNameAndExpire(offset, limit, iName,iExpire)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_KeyValueExpire)
	}
	herr.Message = "Can't get to the GetKeyValueExpiresByNameAndExpire's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetKeyValueExpiresByValueAndExpireHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iValue := self.Args("value").Bytes()
	iExpire := self.Args("expire").MustInt()

	if helper.IsHas(iValue) {
		_KeyValueExpire, _error := model.GetKeyValueExpiresByValueAndExpire(offset, limit, iValue,iExpire)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_KeyValueExpire)
	}
	herr.Message = "Can't get to the GetKeyValueExpiresByValueAndExpire's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetKeyValueExpiresHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	if (offset > 0) && (len(field) > 0) {
		_KeyValueExpire, _error := model.GetKeyValueExpires(offset, limit, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_KeyValueExpire)
	}
	herr.Message = "Can't get to the GetKeyValueExpires' args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasKeyValueExpireViaCollectionHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iCollection := self.Args("collection").String()
	if helper.IsHas(iCollection) {
		_KeyValueExpire := model.HasKeyValueExpireViaCollection(iCollection)
		var m = map[string]interface{}{}
		m["key_value_expire"] = _KeyValueExpire
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasKeyValueExpireViaCollection's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasKeyValueExpireViaNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iName := self.Args("name").String()
	if helper.IsHas(iName) {
		_KeyValueExpire := model.HasKeyValueExpireViaName(iName)
		var m = map[string]interface{}{}
		m["key_value_expire"] = _KeyValueExpire
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasKeyValueExpireViaName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasKeyValueExpireViaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iValue := self.Args("value").Bytes()
	if helper.IsHas(iValue) {
		_KeyValueExpire := model.HasKeyValueExpireViaValue(iValue)
		var m = map[string]interface{}{}
		m["key_value_expire"] = _KeyValueExpire
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasKeyValueExpireViaValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasKeyValueExpireViaExpireHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iExpire := self.Args("expire").MustInt()
	if helper.IsHas(iExpire) {
		_KeyValueExpire := model.HasKeyValueExpireViaExpire(iExpire)
		var m = map[string]interface{}{}
		m["key_value_expire"] = _KeyValueExpire
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasKeyValueExpireViaExpire's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetKeyValueExpireViaCollectionHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iCollection := self.Args("collection").String()
	if helper.IsHas(iCollection) {
		_KeyValueExpire, _error := model.GetKeyValueExpireViaCollection(iCollection)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_KeyValueExpire)
	}
	herr.Message = "Can't get to the GetKeyValueExpireViaCollection's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetKeyValueExpireViaNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iName := self.Args("name").String()
	if helper.IsHas(iName) {
		_KeyValueExpire, _error := model.GetKeyValueExpireViaName(iName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_KeyValueExpire)
	}
	herr.Message = "Can't get to the GetKeyValueExpireViaName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetKeyValueExpireViaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iValue := self.Args("value").Bytes()
	if helper.IsHas(iValue) {
		_KeyValueExpire, _error := model.GetKeyValueExpireViaValue(iValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_KeyValueExpire)
	}
	herr.Message = "Can't get to the GetKeyValueExpireViaValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetKeyValueExpireViaExpireHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iExpire := self.Args("expire").MustInt()
	if helper.IsHas(iExpire) {
		_KeyValueExpire, _error := model.GetKeyValueExpireViaExpire(iExpire)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_KeyValueExpire)
	}
	herr.Message = "Can't get to the GetKeyValueExpireViaExpire's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetKeyValueExpireViaCollectionHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Collection_ := self.Args("collection").String()
	if helper.IsHas(Collection_) {
		var iKeyValueExpire model.KeyValueExpire
		self.Bind(&iKeyValueExpire)
		_KeyValueExpire, _error := model.SetKeyValueExpireViaCollection(Collection_, &iKeyValueExpire)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_KeyValueExpire)
	}
	herr.Message = "Can't get to the SetKeyValueExpireViaCollection's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetKeyValueExpireViaNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Name_ := self.Args("name").String()
	if helper.IsHas(Name_) {
		var iKeyValueExpire model.KeyValueExpire
		self.Bind(&iKeyValueExpire)
		_KeyValueExpire, _error := model.SetKeyValueExpireViaName(Name_, &iKeyValueExpire)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_KeyValueExpire)
	}
	herr.Message = "Can't get to the SetKeyValueExpireViaName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetKeyValueExpireViaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Value_ := self.Args("value").Bytes()
	if helper.IsHas(Value_) {
		var iKeyValueExpire model.KeyValueExpire
		self.Bind(&iKeyValueExpire)
		_KeyValueExpire, _error := model.SetKeyValueExpireViaValue(Value_, &iKeyValueExpire)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_KeyValueExpire)
	}
	herr.Message = "Can't get to the SetKeyValueExpireViaValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetKeyValueExpireViaExpireHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Expire_ := self.Args("expire").MustInt()
	if helper.IsHas(Expire_) {
		var iKeyValueExpire model.KeyValueExpire
		self.Bind(&iKeyValueExpire)
		_KeyValueExpire, _error := model.SetKeyValueExpireViaExpire(Expire_, &iKeyValueExpire)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_KeyValueExpire)
	}
	herr.Message = "Can't get to the SetKeyValueExpireViaExpire's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostAddKeyValueExpireHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Collection_ := self.Args("collection").String()
	Name_ := self.Args("name").String()
	Value_ := self.Args("value").Bytes()
	Expire_ := self.Args("expire").MustInt()

	if helper.IsHas(Collection_) {
		_error := model.AddKeyValueExpire(Collection_,Name_,Value_,Expire_)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		herr.Message = "StatusOK"
		herr.Status = macross.StatusOK
		return self.JSON(herr)
	}
	herr.Message = "Can't get to the AddKeyValueExpire's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostKeyValueExpireHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iKeyValueExpire model.KeyValueExpire
	self.Bind(&iKeyValueExpire)
	_string, _error := model.PostKeyValueExpire(&iKeyValueExpire)
	if (helper.IsHas(_string)) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	} else {
		var m = map[string]interface{}{}
		m["created"] = _string
		return self.JSON(m, macross.StatusCreated)
	}
	return self.JSON(herr)
}

func PutKeyValueExpireHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iKeyValueExpire model.KeyValueExpire
	self.Bind(&iKeyValueExpire)
	_string, _error := model.PutKeyValueExpire(&iKeyValueExpire)
	if (helper.IsHas(_string)) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	} else {
		var m = map[string]interface{}{}
		m["updated"] = _string
		return self.JSON(m)
	}
	return self.JSON(herr)
}

func PutKeyValueExpireViaCollectionHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Collection_ := self.Args("collection").String()
	var iKeyValueExpire model.KeyValueExpire
	self.Bind(&iKeyValueExpire)
	_int64, _error := model.PutKeyValueExpireViaCollection(Collection_, &iKeyValueExpire)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutKeyValueExpireViaNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Name_ := self.Args("name").String()
	var iKeyValueExpire model.KeyValueExpire
	self.Bind(&iKeyValueExpire)
	_int64, _error := model.PutKeyValueExpireViaName(Name_, &iKeyValueExpire)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutKeyValueExpireViaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Value_ := self.Args("value").Bytes()
	var iKeyValueExpire model.KeyValueExpire
	self.Bind(&iKeyValueExpire)
	_int64, _error := model.PutKeyValueExpireViaValue(Value_, &iKeyValueExpire)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutKeyValueExpireViaExpireHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Expire_ := self.Args("expire").MustInt()
	var iKeyValueExpire model.KeyValueExpire
	self.Bind(&iKeyValueExpire)
	_int64, _error := model.PutKeyValueExpireViaExpire(Expire_, &iKeyValueExpire)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateKeyValueExpireViaCollectionHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Collection_ := self.Args("collection").String()
	var iKeyValueExpire model.KeyValueExpire
	self.Bind(&iKeyValueExpire)
	var iMap = helper.StructToMap(iKeyValueExpire)
	_error := model.UpdateKeyValueExpireViaCollection(Collection_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateKeyValueExpireViaNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Name_ := self.Args("name").String()
	var iKeyValueExpire model.KeyValueExpire
	self.Bind(&iKeyValueExpire)
	var iMap = helper.StructToMap(iKeyValueExpire)
	_error := model.UpdateKeyValueExpireViaName(Name_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateKeyValueExpireViaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Value_ := self.Args("value").Bytes()
	var iKeyValueExpire model.KeyValueExpire
	self.Bind(&iKeyValueExpire)
	var iMap = helper.StructToMap(iKeyValueExpire)
	_error := model.UpdateKeyValueExpireViaValue(Value_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateKeyValueExpireViaExpireHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Expire_ := self.Args("expire").MustInt()
	var iKeyValueExpire model.KeyValueExpire
	self.Bind(&iKeyValueExpire)
	var iMap = helper.StructToMap(iKeyValueExpire)
	_error := model.UpdateKeyValueExpireViaExpire(Expire_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteKeyValueExpireHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Collection_ := self.Args("collection").String()
	_int64, _error := model.DeleteKeyValueExpire(Collection_)
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

func DeleteKeyValueExpireViaCollectionHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Collection_ := self.Args("collection").String()
	_error := model.DeleteKeyValueExpireViaCollection(Collection_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteKeyValueExpireViaNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Name_ := self.Args("name").String()
	_error := model.DeleteKeyValueExpireViaName(Name_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteKeyValueExpireViaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Value_ := self.Args("value").Bytes()
	_error := model.DeleteKeyValueExpireViaValue(Value_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteKeyValueExpireViaExpireHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Expire_ := self.Args("expire").MustInt()
	_error := model.DeleteKeyValueExpireViaExpire(Expire_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}
