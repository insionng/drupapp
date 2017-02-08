package handler

import (
	"drupapp/helper"
	"drupapp/model"
	"github.com/insionng/macross"
)

func GetKeyValuesCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()

	if offset > 0 {
		var m = map[string]interface{}{}
		_int64, _error := model.GetKeyValuesCount(offset, limit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		m["key_valuesCount"] = _int64
		return self.JSON(m)
	}
	herr.Message = "Can't get to the GetKeyValuesCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetKeyValueCountViaCollectionHandler(self *macross.Context) error {
	Collection_ := self.Args("collection").String()
	_int64 := model.GetKeyValueCountViaCollection(Collection_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["key_valueCount"] = 0
	}
	m["key_valueCount"] = _int64
	return self.JSON(m)
}

func GetKeyValueCountViaNameHandler(self *macross.Context) error {
	Name_ := self.Args("name").String()
	_int64 := model.GetKeyValueCountViaName(Name_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["key_valueCount"] = 0
	}
	m["key_valueCount"] = _int64
	return self.JSON(m)
}

func GetKeyValueCountViaValueHandler(self *macross.Context) error {
	Value_ := self.Args("value").Bytes()
	_int64 := model.GetKeyValueCountViaValue(Value_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["key_valueCount"] = 0
	}
	m["key_valueCount"] = _int64
	return self.JSON(m)
}

func GetKeyValuesViaCollectionHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iCollection := self.Args("collection").String()
	if (offset > 0) && helper.IsHas(iCollection) {
		_KeyValue, _error := model.GetKeyValuesViaCollection(offset, limit, iCollection, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_KeyValue)
	}
	herr.Message = "Can't get to the GetKeyValuesViaCollection's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetKeyValuesViaNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iName := self.Args("name").String()
	if (offset > 0) && helper.IsHas(iName) {
		_KeyValue, _error := model.GetKeyValuesViaName(offset, limit, iName, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_KeyValue)
	}
	herr.Message = "Can't get to the GetKeyValuesViaName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetKeyValuesViaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iValue := self.Args("value").Bytes()
	if (offset > 0) && helper.IsHas(iValue) {
		_KeyValue, _error := model.GetKeyValuesViaValue(offset, limit, iValue, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_KeyValue)
	}
	herr.Message = "Can't get to the GetKeyValuesViaValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetKeyValuesByCollectionAndNameAndValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCollection := self.Args("collection").String()
	iName := self.Args("name").String()
	iValue := self.Args("value").Bytes()

	if helper.IsHas(iCollection) {
		_KeyValue, _error := model.GetKeyValuesByCollectionAndNameAndValue(offset, limit, iCollection,iName,iValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_KeyValue)
	}
	herr.Message = "Can't get to the GetKeyValuesByCollectionAndNameAndValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetKeyValuesByCollectionAndNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCollection := self.Args("collection").String()
	iName := self.Args("name").String()

	if helper.IsHas(iCollection) {
		_KeyValue, _error := model.GetKeyValuesByCollectionAndName(offset, limit, iCollection,iName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_KeyValue)
	}
	herr.Message = "Can't get to the GetKeyValuesByCollectionAndName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetKeyValuesByCollectionAndValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCollection := self.Args("collection").String()
	iValue := self.Args("value").Bytes()

	if helper.IsHas(iCollection) {
		_KeyValue, _error := model.GetKeyValuesByCollectionAndValue(offset, limit, iCollection,iValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_KeyValue)
	}
	herr.Message = "Can't get to the GetKeyValuesByCollectionAndValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetKeyValuesByNameAndValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iValue := self.Args("value").Bytes()

	if helper.IsHas(iName) {
		_KeyValue, _error := model.GetKeyValuesByNameAndValue(offset, limit, iName,iValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_KeyValue)
	}
	herr.Message = "Can't get to the GetKeyValuesByNameAndValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetKeyValuesHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	if (offset > 0) && (len(field) > 0) {
		_KeyValue, _error := model.GetKeyValues(offset, limit, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_KeyValue)
	}
	herr.Message = "Can't get to the GetKeyValues' args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasKeyValueViaCollectionHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iCollection := self.Args("collection").String()
	if helper.IsHas(iCollection) {
		_KeyValue := model.HasKeyValueViaCollection(iCollection)
		var m = map[string]interface{}{}
		m["key_value"] = _KeyValue
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasKeyValueViaCollection's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasKeyValueViaNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iName := self.Args("name").String()
	if helper.IsHas(iName) {
		_KeyValue := model.HasKeyValueViaName(iName)
		var m = map[string]interface{}{}
		m["key_value"] = _KeyValue
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasKeyValueViaName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasKeyValueViaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iValue := self.Args("value").Bytes()
	if helper.IsHas(iValue) {
		_KeyValue := model.HasKeyValueViaValue(iValue)
		var m = map[string]interface{}{}
		m["key_value"] = _KeyValue
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasKeyValueViaValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetKeyValueViaCollectionHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iCollection := self.Args("collection").String()
	if helper.IsHas(iCollection) {
		_KeyValue, _error := model.GetKeyValueViaCollection(iCollection)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_KeyValue)
	}
	herr.Message = "Can't get to the GetKeyValueViaCollection's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetKeyValueViaNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iName := self.Args("name").String()
	if helper.IsHas(iName) {
		_KeyValue, _error := model.GetKeyValueViaName(iName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_KeyValue)
	}
	herr.Message = "Can't get to the GetKeyValueViaName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetKeyValueViaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iValue := self.Args("value").Bytes()
	if helper.IsHas(iValue) {
		_KeyValue, _error := model.GetKeyValueViaValue(iValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_KeyValue)
	}
	herr.Message = "Can't get to the GetKeyValueViaValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetKeyValueViaCollectionHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Collection_ := self.Args("collection").String()
	if helper.IsHas(Collection_) {
		var iKeyValue model.KeyValue
		self.Bind(&iKeyValue)
		_KeyValue, _error := model.SetKeyValueViaCollection(Collection_, &iKeyValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_KeyValue)
	}
	herr.Message = "Can't get to the SetKeyValueViaCollection's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetKeyValueViaNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Name_ := self.Args("name").String()
	if helper.IsHas(Name_) {
		var iKeyValue model.KeyValue
		self.Bind(&iKeyValue)
		_KeyValue, _error := model.SetKeyValueViaName(Name_, &iKeyValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_KeyValue)
	}
	herr.Message = "Can't get to the SetKeyValueViaName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetKeyValueViaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Value_ := self.Args("value").Bytes()
	if helper.IsHas(Value_) {
		var iKeyValue model.KeyValue
		self.Bind(&iKeyValue)
		_KeyValue, _error := model.SetKeyValueViaValue(Value_, &iKeyValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_KeyValue)
	}
	herr.Message = "Can't get to the SetKeyValueViaValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostAddKeyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Collection_ := self.Args("collection").String()
	Name_ := self.Args("name").String()
	Value_ := self.Args("value").Bytes()

	if helper.IsHas(Collection_) {
		_error := model.AddKeyValue(Collection_,Name_,Value_)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		herr.Message = "StatusOK"
		herr.Status = macross.StatusOK
		return self.JSON(herr)
	}
	herr.Message = "Can't get to the AddKeyValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostKeyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iKeyValue model.KeyValue
	self.Bind(&iKeyValue)
	_string, _error := model.PostKeyValue(&iKeyValue)
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

func PutKeyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iKeyValue model.KeyValue
	self.Bind(&iKeyValue)
	_string, _error := model.PutKeyValue(&iKeyValue)
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

func PutKeyValueViaCollectionHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Collection_ := self.Args("collection").String()
	var iKeyValue model.KeyValue
	self.Bind(&iKeyValue)
	_int64, _error := model.PutKeyValueViaCollection(Collection_, &iKeyValue)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutKeyValueViaNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Name_ := self.Args("name").String()
	var iKeyValue model.KeyValue
	self.Bind(&iKeyValue)
	_int64, _error := model.PutKeyValueViaName(Name_, &iKeyValue)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutKeyValueViaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Value_ := self.Args("value").Bytes()
	var iKeyValue model.KeyValue
	self.Bind(&iKeyValue)
	_int64, _error := model.PutKeyValueViaValue(Value_, &iKeyValue)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateKeyValueViaCollectionHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Collection_ := self.Args("collection").String()
	var iKeyValue model.KeyValue
	self.Bind(&iKeyValue)
	var iMap = helper.StructToMap(iKeyValue)
	_error := model.UpdateKeyValueViaCollection(Collection_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateKeyValueViaNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Name_ := self.Args("name").String()
	var iKeyValue model.KeyValue
	self.Bind(&iKeyValue)
	var iMap = helper.StructToMap(iKeyValue)
	_error := model.UpdateKeyValueViaName(Name_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateKeyValueViaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Value_ := self.Args("value").Bytes()
	var iKeyValue model.KeyValue
	self.Bind(&iKeyValue)
	var iMap = helper.StructToMap(iKeyValue)
	_error := model.UpdateKeyValueViaValue(Value_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteKeyValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Collection_ := self.Args("collection").String()
	_int64, _error := model.DeleteKeyValue(Collection_)
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

func DeleteKeyValueViaCollectionHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Collection_ := self.Args("collection").String()
	_error := model.DeleteKeyValueViaCollection(Collection_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteKeyValueViaNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Name_ := self.Args("name").String()
	_error := model.DeleteKeyValueViaName(Name_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteKeyValueViaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Value_ := self.Args("value").Bytes()
	_error := model.DeleteKeyValueViaValue(Value_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}
