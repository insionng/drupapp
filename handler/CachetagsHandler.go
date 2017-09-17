package handler

import (
	"github.com/insionng/drupapp/helper"
	"github.com/insionng/drupapp/model"
	"github.com/insionng/macross"
)

func GetCachetagsesCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()

	if offset > 0 {
		var m = map[string]interface{}{}
		_int64, _error := model.GetCachetagsesCount(offset, limit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		m["cachetagssCount"] = _int64
		return self.JSON(m)
	}
	herr.Message = "Can't get to the GetCachetagsesCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCachetagsCountViaTagHandler(self *macross.Context) error {
	Tag_ := self.Args("tag").String()
	_int64 := model.GetCachetagsCountViaTag(Tag_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["cachetagsCount"] = 0
	}
	m["cachetagsCount"] = _int64
	return self.JSON(m)
}

func GetCachetagsCountViaInvalidationsHandler(self *macross.Context) error {
	Invalidations_ := self.Args("invalidations").MustInt()
	_int64 := model.GetCachetagsCountViaInvalidations(Invalidations_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["cachetagsCount"] = 0
	}
	m["cachetagsCount"] = _int64
	return self.JSON(m)
}

func GetCachetagsesViaTagHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iTag := self.Args("tag").String()
	if (offset > 0) && helper.IsHas(iTag) {
		_Cachetags, _error := model.GetCachetagsesViaTag(offset, limit, iTag, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Cachetags)
	}
	herr.Message = "Can't get to the GetCachetagsesViaTag's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCachetagsesViaInvalidationsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iInvalidations := self.Args("invalidations").MustInt()
	if (offset > 0) && helper.IsHas(iInvalidations) {
		_Cachetags, _error := model.GetCachetagsesViaInvalidations(offset, limit, iInvalidations, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Cachetags)
	}
	herr.Message = "Can't get to the GetCachetagsesViaInvalidations's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCachetagsesByTagAndInvalidationsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTag := self.Args("tag").String()
	iInvalidations := self.Args("invalidations").MustInt()

	if helper.IsHas(iTag) {
		_Cachetags, _error := model.GetCachetagsesByTagAndInvalidations(offset, limit, iTag,iInvalidations)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Cachetags)
	}
	herr.Message = "Can't get to the GetCachetagsesByTagAndInvalidations's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCachetagsesHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	if (offset > 0) && (len(field) > 0) {
		_Cachetags, _error := model.GetCachetagses(offset, limit, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Cachetags)
	}
	herr.Message = "Can't get to the GetCachetagses' args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasCachetagsViaTagHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iTag := self.Args("tag").String()
	if helper.IsHas(iTag) {
		_Cachetags := model.HasCachetagsViaTag(iTag)
		var m = map[string]interface{}{}
		m["cachetags"] = _Cachetags
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasCachetagsViaTag's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasCachetagsViaInvalidationsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iInvalidations := self.Args("invalidations").MustInt()
	if helper.IsHas(iInvalidations) {
		_Cachetags := model.HasCachetagsViaInvalidations(iInvalidations)
		var m = map[string]interface{}{}
		m["cachetags"] = _Cachetags
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasCachetagsViaInvalidations's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCachetagsViaTagHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iTag := self.Args("tag").String()
	if helper.IsHas(iTag) {
		_Cachetags, _error := model.GetCachetagsViaTag(iTag)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Cachetags)
	}
	herr.Message = "Can't get to the GetCachetagsViaTag's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCachetagsViaInvalidationsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iInvalidations := self.Args("invalidations").MustInt()
	if helper.IsHas(iInvalidations) {
		_Cachetags, _error := model.GetCachetagsViaInvalidations(iInvalidations)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Cachetags)
	}
	herr.Message = "Can't get to the GetCachetagsViaInvalidations's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetCachetagsViaTagHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Tag_ := self.Args("tag").String()
	if helper.IsHas(Tag_) {
		var iCachetags model.Cachetags
		self.Bind(&iCachetags)
		_Cachetags, _error := model.SetCachetagsViaTag(Tag_, &iCachetags)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Cachetags)
	}
	herr.Message = "Can't get to the SetCachetagsViaTag's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetCachetagsViaInvalidationsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Invalidations_ := self.Args("invalidations").MustInt()
	if helper.IsHas(Invalidations_) {
		var iCachetags model.Cachetags
		self.Bind(&iCachetags)
		_Cachetags, _error := model.SetCachetagsViaInvalidations(Invalidations_, &iCachetags)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Cachetags)
	}
	herr.Message = "Can't get to the SetCachetagsViaInvalidations's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostAddCachetagsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Tag_ := self.Args("tag").String()
	Invalidations_ := self.Args("invalidations").MustInt()

	if helper.IsHas(Tag_) {
		_error := model.AddCachetags(Tag_,Invalidations_)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		herr.Message = "StatusOK"
		herr.Status = macross.StatusOK
		return self.JSON(herr)
	}
	herr.Message = "Can't get to the AddCachetags's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostCachetagsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iCachetags model.Cachetags
	self.Bind(&iCachetags)
	_string, _error := model.PostCachetags(&iCachetags)
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

func PutCachetagsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iCachetags model.Cachetags
	self.Bind(&iCachetags)
	_string, _error := model.PutCachetags(&iCachetags)
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

func PutCachetagsViaTagHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Tag_ := self.Args("tag").String()
	var iCachetags model.Cachetags
	self.Bind(&iCachetags)
	_int64, _error := model.PutCachetagsViaTag(Tag_, &iCachetags)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutCachetagsViaInvalidationsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Invalidations_ := self.Args("invalidations").MustInt()
	var iCachetags model.Cachetags
	self.Bind(&iCachetags)
	_int64, _error := model.PutCachetagsViaInvalidations(Invalidations_, &iCachetags)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateCachetagsViaTagHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Tag_ := self.Args("tag").String()
	var iCachetags model.Cachetags
	self.Bind(&iCachetags)
	var iMap = helper.StructToMap(iCachetags)
	_error := model.UpdateCachetagsViaTag(Tag_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateCachetagsViaInvalidationsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Invalidations_ := self.Args("invalidations").MustInt()
	var iCachetags model.Cachetags
	self.Bind(&iCachetags)
	var iMap = helper.StructToMap(iCachetags)
	_error := model.UpdateCachetagsViaInvalidations(Invalidations_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteCachetagsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Tag_ := self.Args("tag").String()
	_int64, _error := model.DeleteCachetags(Tag_)
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

func DeleteCachetagsViaTagHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Tag_ := self.Args("tag").String()
	_error := model.DeleteCachetagsViaTag(Tag_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteCachetagsViaInvalidationsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Invalidations_ := self.Args("invalidations").MustInt()
	_error := model.DeleteCachetagsViaInvalidations(Invalidations_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}
