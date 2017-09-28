package handler

import (
	"github.com/insionng/drupapp/helper"
	"github.com/insionng/drupapp/model"
	"github.com/insionng/macross"
)

func GetConfigsCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()

	if offset > 0 {
		var m = map[string]interface{}{}
		_int64, _error := model.GetConfigsCount(offset, limit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		m["configsCount"] = _int64
		return self.JSON(m)
	}
	herr.Message = "Can't get to the GetConfigsCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetConfigCountViaCollectionHandler(self *macross.Context) error {
	Collection_ := self.Args("collection").String()
	_int64 := model.GetConfigCountViaCollection(Collection_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["configCount"] = 0
	}
	m["configCount"] = _int64
	return self.JSON(m)
}

func GetConfigCountViaNameHandler(self *macross.Context) error {
	Name_ := self.Args("name").String()
	_int64 := model.GetConfigCountViaName(Name_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["configCount"] = 0
	}
	m["configCount"] = _int64
	return self.JSON(m)
}

func GetConfigCountViaDataHandler(self *macross.Context) error {
	Data_ := self.Args("data").Bytes()
	_int64 := model.GetConfigCountViaData(Data_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["configCount"] = 0
	}
	m["configCount"] = _int64
	return self.JSON(m)
}

func GetConfigsViaCollectionHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iCollection := self.Args("collection").String()
	if (offset > 0) && helper.IsHas(iCollection) {
		_Config, _error := model.GetConfigsViaCollection(offset, limit, iCollection, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Config)
	}
	herr.Message = "Can't get to the GetConfigsViaCollection's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetConfigsViaNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iName := self.Args("name").String()
	if (offset > 0) && helper.IsHas(iName) {
		_Config, _error := model.GetConfigsViaName(offset, limit, iName, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Config)
	}
	herr.Message = "Can't get to the GetConfigsViaName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetConfigsViaDataHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iData := self.Args("data").Bytes()
	if (offset > 0) && helper.IsHas(iData) {
		_Config, _error := model.GetConfigsViaData(offset, limit, iData, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Config)
	}
	herr.Message = "Can't get to the GetConfigsViaData's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetConfigsByCollectionAndNameAndDataHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCollection := self.Args("collection").String()
	iName := self.Args("name").String()
	iData := self.Args("data").Bytes()

	if helper.IsHas(iCollection) {
		_Config, _error := model.GetConfigsByCollectionAndNameAndData(offset, limit, iCollection,iName,iData)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Config)
	}
	herr.Message = "Can't get to the GetConfigsByCollectionAndNameAndData's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetConfigsByCollectionAndNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCollection := self.Args("collection").String()
	iName := self.Args("name").String()

	if helper.IsHas(iCollection) {
		_Config, _error := model.GetConfigsByCollectionAndName(offset, limit, iCollection,iName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Config)
	}
	herr.Message = "Can't get to the GetConfigsByCollectionAndName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetConfigsByCollectionAndDataHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCollection := self.Args("collection").String()
	iData := self.Args("data").Bytes()

	if helper.IsHas(iCollection) {
		_Config, _error := model.GetConfigsByCollectionAndData(offset, limit, iCollection,iData)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Config)
	}
	herr.Message = "Can't get to the GetConfigsByCollectionAndData's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetConfigsByNameAndDataHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iData := self.Args("data").Bytes()

	if helper.IsHas(iName) {
		_Config, _error := model.GetConfigsByNameAndData(offset, limit, iName,iData)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Config)
	}
	herr.Message = "Can't get to the GetConfigsByNameAndData's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetConfigsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	if (offset > 0) && (len(field) > 0) {
		_Config, _error := model.GetConfigs(offset, limit, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Config)
	}
	herr.Message = "Can't get to the GetConfigs' args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasConfigViaCollectionHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iCollection := self.Args("collection").String()
	if helper.IsHas(iCollection) {
		_Config := model.HasConfigViaCollection(iCollection)
		var m = map[string]interface{}{}
		m["config"] = _Config
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasConfigViaCollection's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasConfigViaNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iName := self.Args("name").String()
	if helper.IsHas(iName) {
		_Config := model.HasConfigViaName(iName)
		var m = map[string]interface{}{}
		m["config"] = _Config
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasConfigViaName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasConfigViaDataHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iData := self.Args("data").Bytes()
	if helper.IsHas(iData) {
		_Config := model.HasConfigViaData(iData)
		var m = map[string]interface{}{}
		m["config"] = _Config
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasConfigViaData's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetConfigViaCollectionHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iCollection := self.Args("collection").String()
	if helper.IsHas(iCollection) {
		_Config, _error := model.GetConfigViaCollection(iCollection)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Config)
	}
	herr.Message = "Can't get to the GetConfigViaCollection's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetConfigViaNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iName := self.Args("name").String()
	if helper.IsHas(iName) {
		_Config, _error := model.GetConfigViaName(iName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Config)
	}
	herr.Message = "Can't get to the GetConfigViaName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetConfigViaDataHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iData := self.Args("data").Bytes()
	if helper.IsHas(iData) {
		_Config, _error := model.GetConfigViaData(iData)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Config)
	}
	herr.Message = "Can't get to the GetConfigViaData's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetConfigViaCollectionHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Collection_ := self.Args("collection").String()
	if helper.IsHas(Collection_) {
		var iConfig model.Config
		self.Bind(&iConfig)
		_Config, _error := model.SetConfigViaCollection(Collection_, &iConfig)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Config)
	}
	herr.Message = "Can't get to the SetConfigViaCollection's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetConfigViaNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Name_ := self.Args("name").String()
	if helper.IsHas(Name_) {
		var iConfig model.Config
		self.Bind(&iConfig)
		_Config, _error := model.SetConfigViaName(Name_, &iConfig)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Config)
	}
	herr.Message = "Can't get to the SetConfigViaName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetConfigViaDataHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Data_ := self.Args("data").Bytes()
	if helper.IsHas(Data_) {
		var iConfig model.Config
		self.Bind(&iConfig)
		_Config, _error := model.SetConfigViaData(Data_, &iConfig)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Config)
	}
	herr.Message = "Can't get to the SetConfigViaData's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostAddConfigHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Collection_ := self.Args("collection").String()
	Name_ := self.Args("name").String()
	Data_ := self.Args("data").Bytes()

	if helper.IsHas(Collection_) {
		_error := model.AddConfig(Collection_,Name_,Data_)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		herr.Message = "StatusOK"
		herr.Status = macross.StatusOK
		return self.JSON(herr)
	}
	herr.Message = "Can't get to the AddConfig's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostConfigHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iConfig model.Config
	self.Bind(&iConfig)
	_string, _error := model.PostConfig(&iConfig)
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

func PutConfigHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iConfig model.Config
	self.Bind(&iConfig)
	_string, _error := model.PutConfig(&iConfig)
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

func PutConfigViaCollectionHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Collection_ := self.Args("collection").String()
	var iConfig model.Config
	self.Bind(&iConfig)
	_int64, _error := model.PutConfigViaCollection(Collection_, &iConfig)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutConfigViaNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Name_ := self.Args("name").String()
	var iConfig model.Config
	self.Bind(&iConfig)
	_int64, _error := model.PutConfigViaName(Name_, &iConfig)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutConfigViaDataHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Data_ := self.Args("data").Bytes()
	var iConfig model.Config
	self.Bind(&iConfig)
	_int64, _error := model.PutConfigViaData(Data_, &iConfig)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateConfigViaCollectionHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Collection_ := self.Args("collection").String()
	var iConfig model.Config
	self.Bind(&iConfig)
	var iMap = helper.StructToMap(iConfig)
	_error := model.UpdateConfigViaCollection(Collection_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateConfigViaNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Name_ := self.Args("name").String()
	var iConfig model.Config
	self.Bind(&iConfig)
	var iMap = helper.StructToMap(iConfig)
	_error := model.UpdateConfigViaName(Name_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateConfigViaDataHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Data_ := self.Args("data").Bytes()
	var iConfig model.Config
	self.Bind(&iConfig)
	var iMap = helper.StructToMap(iConfig)
	_error := model.UpdateConfigViaData(Data_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteConfigHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Collection_ := self.Args("collection").String()
	_int64, _error := model.DeleteConfig(Collection_)
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

func DeleteConfigViaCollectionHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Collection_ := self.Args("collection").String()
	_error := model.DeleteConfigViaCollection(Collection_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteConfigViaNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Name_ := self.Args("name").String()
	_error := model.DeleteConfigViaName(Name_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteConfigViaDataHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Data_ := self.Args("data").Bytes()
	_error := model.DeleteConfigViaData(Data_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}
