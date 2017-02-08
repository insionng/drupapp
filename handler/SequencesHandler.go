package handler

import (
	"drupapp/helper"
	"drupapp/model"
	"github.com/insionng/macross"
)

func GetSequencesesCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()

	if offset > 0 {
		var m = map[string]interface{}{}
		_int64, _error := model.GetSequencesesCount(offset, limit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		m["sequencessCount"] = _int64
		return self.JSON(m)
	}
	herr.Message = "Can't get to the GetSequencesesCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSequencesCountViaValueHandler(self *macross.Context) error {
	Value_ := self.Args("value").MustInt64()
	_int64 := model.GetSequencesCountViaValue(Value_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["sequencesCount"] = 0
	}
	m["sequencesCount"] = _int64
	return self.JSON(m)
}

func GetSequencesesViaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iValue := self.Args("value").MustInt64()
	if (offset > 0) && helper.IsHas(iValue) {
		_Sequences, _error := model.GetSequencesesViaValue(offset, limit, iValue, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Sequences)
	}
	herr.Message = "Can't get to the GetSequencesesViaValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSequencesesByValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iValue := self.Args("value").MustInt64()

	if helper.IsHas(iValue) {
		_Sequences, _error := model.GetSequencesesByValue(offset, limit, iValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Sequences)
	}
	herr.Message = "Can't get to the GetSequencesesByValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSequencesesHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	if (offset > 0) && (len(field) > 0) {
		_Sequences, _error := model.GetSequenceses(offset, limit, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Sequences)
	}
	herr.Message = "Can't get to the GetSequenceses' args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasSequencesViaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iValue := self.Args("value").MustInt64()
	if helper.IsHas(iValue) {
		_Sequences := model.HasSequencesViaValue(iValue)
		var m = map[string]interface{}{}
		m["sequences"] = _Sequences
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasSequencesViaValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSequencesViaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iValue := self.Args("value").MustInt64()
	if helper.IsHas(iValue) {
		_Sequences, _error := model.GetSequencesViaValue(iValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Sequences)
	}
	herr.Message = "Can't get to the GetSequencesViaValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetSequencesViaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Value_ := self.Args("value").MustInt64()
	if helper.IsHas(Value_) {
		var iSequences model.Sequences
		self.Bind(&iSequences)
		_Sequences, _error := model.SetSequencesViaValue(Value_, &iSequences)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Sequences)
	}
	herr.Message = "Can't get to the SetSequencesViaValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostAddSequencesHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Value_ := self.Args("value").MustInt64()

	if helper.IsHas(Value_) {
		_error := model.AddSequences(Value_)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		herr.Message = "StatusOK"
		herr.Status = macross.StatusOK
		return self.JSON(herr)
	}
	herr.Message = "Can't get to the AddSequences's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSequencesHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iSequences model.Sequences
	self.Bind(&iSequences)
	_int64, _error := model.PostSequences(&iSequences)
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

func PutSequencesHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iSequences model.Sequences
	self.Bind(&iSequences)
	_int64, _error := model.PutSequences(&iSequences)
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

func PutSequencesViaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Value_ := self.Args("value").MustInt64()
	var iSequences model.Sequences
	self.Bind(&iSequences)
	_int64, _error := model.PutSequencesViaValue(Value_, &iSequences)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateSequencesViaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Value_ := self.Args("value").MustInt64()
	var iSequences model.Sequences
	self.Bind(&iSequences)
	var iMap = helper.StructToMap(iSequences)
	_error := model.UpdateSequencesViaValue(Value_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteSequencesHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Value_ := self.Args("value").MustInt64()
	_int64, _error := model.DeleteSequences(Value_)
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

func DeleteSequencesViaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Value_ := self.Args("value").MustInt64()
	_error := model.DeleteSequencesViaValue(Value_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}
