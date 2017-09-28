package handler

import (
	"github.com/insionng/drupapp/helper"
	"github.com/insionng/drupapp/model"
	"github.com/insionng/macross"
)

func GetSemaphoresCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()

	if offset > 0 {
		var m = map[string]interface{}{}
		_int64, _error := model.GetSemaphoresCount(offset, limit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		m["semaphoresCount"] = _int64
		return self.JSON(m)
	}
	herr.Message = "Can't get to the GetSemaphoresCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSemaphoreCountViaNameHandler(self *macross.Context) error {
	Name_ := self.Args("name").String()
	_int64 := model.GetSemaphoreCountViaName(Name_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["semaphoreCount"] = 0
	}
	m["semaphoreCount"] = _int64
	return self.JSON(m)
}

func GetSemaphoreCountViaValueHandler(self *macross.Context) error {
	Value_ := self.Args("value").String()
	_int64 := model.GetSemaphoreCountViaValue(Value_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["semaphoreCount"] = 0
	}
	m["semaphoreCount"] = _int64
	return self.JSON(m)
}

func GetSemaphoreCountViaExpireHandler(self *macross.Context) error {
	Expire_ := self.Args("expire").MustFloat64()
	_int64 := model.GetSemaphoreCountViaExpire(Expire_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["semaphoreCount"] = 0
	}
	m["semaphoreCount"] = _int64
	return self.JSON(m)
}

func GetSemaphoresViaNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iName := self.Args("name").String()
	if (offset > 0) && helper.IsHas(iName) {
		_Semaphore, _error := model.GetSemaphoresViaName(offset, limit, iName, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Semaphore)
	}
	herr.Message = "Can't get to the GetSemaphoresViaName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSemaphoresViaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iValue := self.Args("value").String()
	if (offset > 0) && helper.IsHas(iValue) {
		_Semaphore, _error := model.GetSemaphoresViaValue(offset, limit, iValue, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Semaphore)
	}
	herr.Message = "Can't get to the GetSemaphoresViaValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSemaphoresViaExpireHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iExpire := self.Args("expire").MustFloat64()
	if (offset > 0) && helper.IsHas(iExpire) {
		_Semaphore, _error := model.GetSemaphoresViaExpire(offset, limit, iExpire, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Semaphore)
	}
	herr.Message = "Can't get to the GetSemaphoresViaExpire's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSemaphoresByNameAndValueAndExpireHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iValue := self.Args("value").String()
	iExpire := self.Args("expire").MustFloat64()

	if helper.IsHas(iName) {
		_Semaphore, _error := model.GetSemaphoresByNameAndValueAndExpire(offset, limit, iName,iValue,iExpire)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Semaphore)
	}
	herr.Message = "Can't get to the GetSemaphoresByNameAndValueAndExpire's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSemaphoresByNameAndValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iValue := self.Args("value").String()

	if helper.IsHas(iName) {
		_Semaphore, _error := model.GetSemaphoresByNameAndValue(offset, limit, iName,iValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Semaphore)
	}
	herr.Message = "Can't get to the GetSemaphoresByNameAndValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSemaphoresByNameAndExpireHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iExpire := self.Args("expire").MustFloat64()

	if helper.IsHas(iName) {
		_Semaphore, _error := model.GetSemaphoresByNameAndExpire(offset, limit, iName,iExpire)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Semaphore)
	}
	herr.Message = "Can't get to the GetSemaphoresByNameAndExpire's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSemaphoresByValueAndExpireHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iValue := self.Args("value").String()
	iExpire := self.Args("expire").MustFloat64()

	if helper.IsHas(iValue) {
		_Semaphore, _error := model.GetSemaphoresByValueAndExpire(offset, limit, iValue,iExpire)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Semaphore)
	}
	herr.Message = "Can't get to the GetSemaphoresByValueAndExpire's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSemaphoresHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	if (offset > 0) && (len(field) > 0) {
		_Semaphore, _error := model.GetSemaphores(offset, limit, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Semaphore)
	}
	herr.Message = "Can't get to the GetSemaphores' args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasSemaphoreViaNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iName := self.Args("name").String()
	if helper.IsHas(iName) {
		_Semaphore := model.HasSemaphoreViaName(iName)
		var m = map[string]interface{}{}
		m["semaphore"] = _Semaphore
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasSemaphoreViaName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasSemaphoreViaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iValue := self.Args("value").String()
	if helper.IsHas(iValue) {
		_Semaphore := model.HasSemaphoreViaValue(iValue)
		var m = map[string]interface{}{}
		m["semaphore"] = _Semaphore
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasSemaphoreViaValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasSemaphoreViaExpireHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iExpire := self.Args("expire").MustFloat64()
	if helper.IsHas(iExpire) {
		_Semaphore := model.HasSemaphoreViaExpire(iExpire)
		var m = map[string]interface{}{}
		m["semaphore"] = _Semaphore
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasSemaphoreViaExpire's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSemaphoreViaNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iName := self.Args("name").String()
	if helper.IsHas(iName) {
		_Semaphore, _error := model.GetSemaphoreViaName(iName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Semaphore)
	}
	herr.Message = "Can't get to the GetSemaphoreViaName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSemaphoreViaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iValue := self.Args("value").String()
	if helper.IsHas(iValue) {
		_Semaphore, _error := model.GetSemaphoreViaValue(iValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Semaphore)
	}
	herr.Message = "Can't get to the GetSemaphoreViaValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSemaphoreViaExpireHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iExpire := self.Args("expire").MustFloat64()
	if helper.IsHas(iExpire) {
		_Semaphore, _error := model.GetSemaphoreViaExpire(iExpire)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Semaphore)
	}
	herr.Message = "Can't get to the GetSemaphoreViaExpire's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetSemaphoreViaNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Name_ := self.Args("name").String()
	if helper.IsHas(Name_) {
		var iSemaphore model.Semaphore
		self.Bind(&iSemaphore)
		_Semaphore, _error := model.SetSemaphoreViaName(Name_, &iSemaphore)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Semaphore)
	}
	herr.Message = "Can't get to the SetSemaphoreViaName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetSemaphoreViaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Value_ := self.Args("value").String()
	if helper.IsHas(Value_) {
		var iSemaphore model.Semaphore
		self.Bind(&iSemaphore)
		_Semaphore, _error := model.SetSemaphoreViaValue(Value_, &iSemaphore)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Semaphore)
	}
	herr.Message = "Can't get to the SetSemaphoreViaValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetSemaphoreViaExpireHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Expire_ := self.Args("expire").MustFloat64()
	if helper.IsHas(Expire_) {
		var iSemaphore model.Semaphore
		self.Bind(&iSemaphore)
		_Semaphore, _error := model.SetSemaphoreViaExpire(Expire_, &iSemaphore)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Semaphore)
	}
	herr.Message = "Can't get to the SetSemaphoreViaExpire's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostAddSemaphoreHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Name_ := self.Args("name").String()
	Value_ := self.Args("value").String()
	Expire_ := self.Args("expire").MustFloat64()

	if helper.IsHas(Name_) {
		_error := model.AddSemaphore(Name_,Value_,Expire_)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		herr.Message = "StatusOK"
		herr.Status = macross.StatusOK
		return self.JSON(herr)
	}
	herr.Message = "Can't get to the AddSemaphore's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSemaphoreHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iSemaphore model.Semaphore
	self.Bind(&iSemaphore)
	_string, _error := model.PostSemaphore(&iSemaphore)
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

func PutSemaphoreHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iSemaphore model.Semaphore
	self.Bind(&iSemaphore)
	_string, _error := model.PutSemaphore(&iSemaphore)
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

func PutSemaphoreViaNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Name_ := self.Args("name").String()
	var iSemaphore model.Semaphore
	self.Bind(&iSemaphore)
	_int64, _error := model.PutSemaphoreViaName(Name_, &iSemaphore)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutSemaphoreViaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Value_ := self.Args("value").String()
	var iSemaphore model.Semaphore
	self.Bind(&iSemaphore)
	_int64, _error := model.PutSemaphoreViaValue(Value_, &iSemaphore)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutSemaphoreViaExpireHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Expire_ := self.Args("expire").MustFloat64()
	var iSemaphore model.Semaphore
	self.Bind(&iSemaphore)
	_int64, _error := model.PutSemaphoreViaExpire(Expire_, &iSemaphore)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateSemaphoreViaNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Name_ := self.Args("name").String()
	var iSemaphore model.Semaphore
	self.Bind(&iSemaphore)
	var iMap = helper.StructToMap(iSemaphore)
	_error := model.UpdateSemaphoreViaName(Name_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateSemaphoreViaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Value_ := self.Args("value").String()
	var iSemaphore model.Semaphore
	self.Bind(&iSemaphore)
	var iMap = helper.StructToMap(iSemaphore)
	_error := model.UpdateSemaphoreViaValue(Value_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateSemaphoreViaExpireHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Expire_ := self.Args("expire").MustFloat64()
	var iSemaphore model.Semaphore
	self.Bind(&iSemaphore)
	var iMap = helper.StructToMap(iSemaphore)
	_error := model.UpdateSemaphoreViaExpire(Expire_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteSemaphoreHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Name_ := self.Args("name").String()
	_int64, _error := model.DeleteSemaphore(Name_)
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

func DeleteSemaphoreViaNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Name_ := self.Args("name").String()
	_error := model.DeleteSemaphoreViaName(Name_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteSemaphoreViaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Value_ := self.Args("value").String()
	_error := model.DeleteSemaphoreViaValue(Value_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteSemaphoreViaExpireHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Expire_ := self.Args("expire").MustFloat64()
	_error := model.DeleteSemaphoreViaExpire(Expire_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}
