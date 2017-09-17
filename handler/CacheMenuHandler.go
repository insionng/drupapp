package handler

import (
	"github.com/insionng/drupapp/helper"
	"github.com/insionng/drupapp/model"
	"github.com/insionng/macross"
)

func GetCacheMenusCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()

	if offset > 0 {
		var m = map[string]interface{}{}
		_int64, _error := model.GetCacheMenusCount(offset, limit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		m["cache_menusCount"] = _int64
		return self.JSON(m)
	}
	herr.Message = "Can't get to the GetCacheMenusCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheMenuCountViaCidHandler(self *macross.Context) error {
	Cid_ := self.Args("cid").String()
	_int64 := model.GetCacheMenuCountViaCid(Cid_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["cache_menuCount"] = 0
	}
	m["cache_menuCount"] = _int64
	return self.JSON(m)
}

func GetCacheMenuCountViaDataHandler(self *macross.Context) error {
	Data_ := self.Args("data").Bytes()
	_int64 := model.GetCacheMenuCountViaData(Data_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["cache_menuCount"] = 0
	}
	m["cache_menuCount"] = _int64
	return self.JSON(m)
}

func GetCacheMenuCountViaExpireHandler(self *macross.Context) error {
	Expire_ := self.Args("expire").MustInt()
	_int64 := model.GetCacheMenuCountViaExpire(Expire_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["cache_menuCount"] = 0
	}
	m["cache_menuCount"] = _int64
	return self.JSON(m)
}

func GetCacheMenuCountViaCreatedHandler(self *macross.Context) error {
	Created_ := self.Args("created").String()
	_int64 := model.GetCacheMenuCountViaCreated(Created_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["cache_menuCount"] = 0
	}
	m["cache_menuCount"] = _int64
	return self.JSON(m)
}

func GetCacheMenuCountViaSerializedHandler(self *macross.Context) error {
	Serialized_ := self.Args("serialized").MustInt()
	_int64 := model.GetCacheMenuCountViaSerialized(Serialized_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["cache_menuCount"] = 0
	}
	m["cache_menuCount"] = _int64
	return self.JSON(m)
}

func GetCacheMenuCountViaTagsHandler(self *macross.Context) error {
	Tags_ := self.Args("tags").String()
	_int64 := model.GetCacheMenuCountViaTags(Tags_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["cache_menuCount"] = 0
	}
	m["cache_menuCount"] = _int64
	return self.JSON(m)
}

func GetCacheMenuCountViaChecksumHandler(self *macross.Context) error {
	Checksum_ := self.Args("checksum").String()
	_int64 := model.GetCacheMenuCountViaChecksum(Checksum_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["cache_menuCount"] = 0
	}
	m["cache_menuCount"] = _int64
	return self.JSON(m)
}

func GetCacheMenusViaCidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iCid := self.Args("cid").String()
	if (offset > 0) && helper.IsHas(iCid) {
		_CacheMenu, _error := model.GetCacheMenusViaCid(offset, limit, iCid, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheMenu)
	}
	herr.Message = "Can't get to the GetCacheMenusViaCid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheMenusViaDataHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iData := self.Args("data").Bytes()
	if (offset > 0) && helper.IsHas(iData) {
		_CacheMenu, _error := model.GetCacheMenusViaData(offset, limit, iData, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheMenu)
	}
	herr.Message = "Can't get to the GetCacheMenusViaData's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheMenusViaExpireHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iExpire := self.Args("expire").MustInt()
	if (offset > 0) && helper.IsHas(iExpire) {
		_CacheMenu, _error := model.GetCacheMenusViaExpire(offset, limit, iExpire, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheMenu)
	}
	herr.Message = "Can't get to the GetCacheMenusViaExpire's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheMenusViaCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iCreated := self.Args("created").String()
	if (offset > 0) && helper.IsHas(iCreated) {
		_CacheMenu, _error := model.GetCacheMenusViaCreated(offset, limit, iCreated, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheMenu)
	}
	herr.Message = "Can't get to the GetCacheMenusViaCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheMenusViaSerializedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iSerialized := self.Args("serialized").MustInt()
	if (offset > 0) && helper.IsHas(iSerialized) {
		_CacheMenu, _error := model.GetCacheMenusViaSerialized(offset, limit, iSerialized, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheMenu)
	}
	herr.Message = "Can't get to the GetCacheMenusViaSerialized's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheMenusViaTagsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iTags := self.Args("tags").String()
	if (offset > 0) && helper.IsHas(iTags) {
		_CacheMenu, _error := model.GetCacheMenusViaTags(offset, limit, iTags, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheMenu)
	}
	herr.Message = "Can't get to the GetCacheMenusViaTags's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheMenusViaChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iChecksum := self.Args("checksum").String()
	if (offset > 0) && helper.IsHas(iChecksum) {
		_CacheMenu, _error := model.GetCacheMenusViaChecksum(offset, limit, iChecksum, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheMenu)
	}
	herr.Message = "Can't get to the GetCacheMenusViaChecksum's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheMenusByCidAndDataAndExpireHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").String()
	iData := self.Args("data").Bytes()
	iExpire := self.Args("expire").MustInt()

	if helper.IsHas(iCid) {
		_CacheMenu, _error := model.GetCacheMenusByCidAndDataAndExpire(offset, limit, iCid,iData,iExpire)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheMenu)
	}
	herr.Message = "Can't get to the GetCacheMenusByCidAndDataAndExpire's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheMenusByCidAndDataAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").String()
	iData := self.Args("data").Bytes()
	iCreated := self.Args("created").String()

	if helper.IsHas(iCid) {
		_CacheMenu, _error := model.GetCacheMenusByCidAndDataAndCreated(offset, limit, iCid,iData,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheMenu)
	}
	herr.Message = "Can't get to the GetCacheMenusByCidAndDataAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheMenusByCidAndDataAndSerializedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").String()
	iData := self.Args("data").Bytes()
	iSerialized := self.Args("serialized").MustInt()

	if helper.IsHas(iCid) {
		_CacheMenu, _error := model.GetCacheMenusByCidAndDataAndSerialized(offset, limit, iCid,iData,iSerialized)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheMenu)
	}
	herr.Message = "Can't get to the GetCacheMenusByCidAndDataAndSerialized's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheMenusByCidAndDataAndTagsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").String()
	iData := self.Args("data").Bytes()
	iTags := self.Args("tags").String()

	if helper.IsHas(iCid) {
		_CacheMenu, _error := model.GetCacheMenusByCidAndDataAndTags(offset, limit, iCid,iData,iTags)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheMenu)
	}
	herr.Message = "Can't get to the GetCacheMenusByCidAndDataAndTags's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheMenusByCidAndDataAndChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").String()
	iData := self.Args("data").Bytes()
	iChecksum := self.Args("checksum").String()

	if helper.IsHas(iCid) {
		_CacheMenu, _error := model.GetCacheMenusByCidAndDataAndChecksum(offset, limit, iCid,iData,iChecksum)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheMenu)
	}
	herr.Message = "Can't get to the GetCacheMenusByCidAndDataAndChecksum's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheMenusByCidAndExpireAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").String()
	iExpire := self.Args("expire").MustInt()
	iCreated := self.Args("created").String()

	if helper.IsHas(iCid) {
		_CacheMenu, _error := model.GetCacheMenusByCidAndExpireAndCreated(offset, limit, iCid,iExpire,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheMenu)
	}
	herr.Message = "Can't get to the GetCacheMenusByCidAndExpireAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheMenusByCidAndExpireAndSerializedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").String()
	iExpire := self.Args("expire").MustInt()
	iSerialized := self.Args("serialized").MustInt()

	if helper.IsHas(iCid) {
		_CacheMenu, _error := model.GetCacheMenusByCidAndExpireAndSerialized(offset, limit, iCid,iExpire,iSerialized)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheMenu)
	}
	herr.Message = "Can't get to the GetCacheMenusByCidAndExpireAndSerialized's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheMenusByCidAndExpireAndTagsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").String()
	iExpire := self.Args("expire").MustInt()
	iTags := self.Args("tags").String()

	if helper.IsHas(iCid) {
		_CacheMenu, _error := model.GetCacheMenusByCidAndExpireAndTags(offset, limit, iCid,iExpire,iTags)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheMenu)
	}
	herr.Message = "Can't get to the GetCacheMenusByCidAndExpireAndTags's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheMenusByCidAndExpireAndChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").String()
	iExpire := self.Args("expire").MustInt()
	iChecksum := self.Args("checksum").String()

	if helper.IsHas(iCid) {
		_CacheMenu, _error := model.GetCacheMenusByCidAndExpireAndChecksum(offset, limit, iCid,iExpire,iChecksum)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheMenu)
	}
	herr.Message = "Can't get to the GetCacheMenusByCidAndExpireAndChecksum's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheMenusByCidAndCreatedAndSerializedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").String()
	iCreated := self.Args("created").String()
	iSerialized := self.Args("serialized").MustInt()

	if helper.IsHas(iCid) {
		_CacheMenu, _error := model.GetCacheMenusByCidAndCreatedAndSerialized(offset, limit, iCid,iCreated,iSerialized)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheMenu)
	}
	herr.Message = "Can't get to the GetCacheMenusByCidAndCreatedAndSerialized's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheMenusByCidAndCreatedAndTagsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").String()
	iCreated := self.Args("created").String()
	iTags := self.Args("tags").String()

	if helper.IsHas(iCid) {
		_CacheMenu, _error := model.GetCacheMenusByCidAndCreatedAndTags(offset, limit, iCid,iCreated,iTags)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheMenu)
	}
	herr.Message = "Can't get to the GetCacheMenusByCidAndCreatedAndTags's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheMenusByCidAndCreatedAndChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").String()
	iCreated := self.Args("created").String()
	iChecksum := self.Args("checksum").String()

	if helper.IsHas(iCid) {
		_CacheMenu, _error := model.GetCacheMenusByCidAndCreatedAndChecksum(offset, limit, iCid,iCreated,iChecksum)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheMenu)
	}
	herr.Message = "Can't get to the GetCacheMenusByCidAndCreatedAndChecksum's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheMenusByCidAndSerializedAndTagsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").String()
	iSerialized := self.Args("serialized").MustInt()
	iTags := self.Args("tags").String()

	if helper.IsHas(iCid) {
		_CacheMenu, _error := model.GetCacheMenusByCidAndSerializedAndTags(offset, limit, iCid,iSerialized,iTags)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheMenu)
	}
	herr.Message = "Can't get to the GetCacheMenusByCidAndSerializedAndTags's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheMenusByCidAndSerializedAndChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").String()
	iSerialized := self.Args("serialized").MustInt()
	iChecksum := self.Args("checksum").String()

	if helper.IsHas(iCid) {
		_CacheMenu, _error := model.GetCacheMenusByCidAndSerializedAndChecksum(offset, limit, iCid,iSerialized,iChecksum)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheMenu)
	}
	herr.Message = "Can't get to the GetCacheMenusByCidAndSerializedAndChecksum's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheMenusByCidAndTagsAndChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").String()
	iTags := self.Args("tags").String()
	iChecksum := self.Args("checksum").String()

	if helper.IsHas(iCid) {
		_CacheMenu, _error := model.GetCacheMenusByCidAndTagsAndChecksum(offset, limit, iCid,iTags,iChecksum)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheMenu)
	}
	herr.Message = "Can't get to the GetCacheMenusByCidAndTagsAndChecksum's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheMenusByDataAndExpireAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iData := self.Args("data").Bytes()
	iExpire := self.Args("expire").MustInt()
	iCreated := self.Args("created").String()

	if helper.IsHas(iData) {
		_CacheMenu, _error := model.GetCacheMenusByDataAndExpireAndCreated(offset, limit, iData,iExpire,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheMenu)
	}
	herr.Message = "Can't get to the GetCacheMenusByDataAndExpireAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheMenusByDataAndExpireAndSerializedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iData := self.Args("data").Bytes()
	iExpire := self.Args("expire").MustInt()
	iSerialized := self.Args("serialized").MustInt()

	if helper.IsHas(iData) {
		_CacheMenu, _error := model.GetCacheMenusByDataAndExpireAndSerialized(offset, limit, iData,iExpire,iSerialized)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheMenu)
	}
	herr.Message = "Can't get to the GetCacheMenusByDataAndExpireAndSerialized's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheMenusByDataAndExpireAndTagsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iData := self.Args("data").Bytes()
	iExpire := self.Args("expire").MustInt()
	iTags := self.Args("tags").String()

	if helper.IsHas(iData) {
		_CacheMenu, _error := model.GetCacheMenusByDataAndExpireAndTags(offset, limit, iData,iExpire,iTags)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheMenu)
	}
	herr.Message = "Can't get to the GetCacheMenusByDataAndExpireAndTags's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheMenusByDataAndExpireAndChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iData := self.Args("data").Bytes()
	iExpire := self.Args("expire").MustInt()
	iChecksum := self.Args("checksum").String()

	if helper.IsHas(iData) {
		_CacheMenu, _error := model.GetCacheMenusByDataAndExpireAndChecksum(offset, limit, iData,iExpire,iChecksum)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheMenu)
	}
	herr.Message = "Can't get to the GetCacheMenusByDataAndExpireAndChecksum's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheMenusByDataAndCreatedAndSerializedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iData := self.Args("data").Bytes()
	iCreated := self.Args("created").String()
	iSerialized := self.Args("serialized").MustInt()

	if helper.IsHas(iData) {
		_CacheMenu, _error := model.GetCacheMenusByDataAndCreatedAndSerialized(offset, limit, iData,iCreated,iSerialized)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheMenu)
	}
	herr.Message = "Can't get to the GetCacheMenusByDataAndCreatedAndSerialized's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheMenusByDataAndCreatedAndTagsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iData := self.Args("data").Bytes()
	iCreated := self.Args("created").String()
	iTags := self.Args("tags").String()

	if helper.IsHas(iData) {
		_CacheMenu, _error := model.GetCacheMenusByDataAndCreatedAndTags(offset, limit, iData,iCreated,iTags)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheMenu)
	}
	herr.Message = "Can't get to the GetCacheMenusByDataAndCreatedAndTags's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheMenusByDataAndCreatedAndChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iData := self.Args("data").Bytes()
	iCreated := self.Args("created").String()
	iChecksum := self.Args("checksum").String()

	if helper.IsHas(iData) {
		_CacheMenu, _error := model.GetCacheMenusByDataAndCreatedAndChecksum(offset, limit, iData,iCreated,iChecksum)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheMenu)
	}
	herr.Message = "Can't get to the GetCacheMenusByDataAndCreatedAndChecksum's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheMenusByDataAndSerializedAndTagsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iData := self.Args("data").Bytes()
	iSerialized := self.Args("serialized").MustInt()
	iTags := self.Args("tags").String()

	if helper.IsHas(iData) {
		_CacheMenu, _error := model.GetCacheMenusByDataAndSerializedAndTags(offset, limit, iData,iSerialized,iTags)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheMenu)
	}
	herr.Message = "Can't get to the GetCacheMenusByDataAndSerializedAndTags's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheMenusByDataAndSerializedAndChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iData := self.Args("data").Bytes()
	iSerialized := self.Args("serialized").MustInt()
	iChecksum := self.Args("checksum").String()

	if helper.IsHas(iData) {
		_CacheMenu, _error := model.GetCacheMenusByDataAndSerializedAndChecksum(offset, limit, iData,iSerialized,iChecksum)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheMenu)
	}
	herr.Message = "Can't get to the GetCacheMenusByDataAndSerializedAndChecksum's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheMenusByDataAndTagsAndChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iData := self.Args("data").Bytes()
	iTags := self.Args("tags").String()
	iChecksum := self.Args("checksum").String()

	if helper.IsHas(iData) {
		_CacheMenu, _error := model.GetCacheMenusByDataAndTagsAndChecksum(offset, limit, iData,iTags,iChecksum)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheMenu)
	}
	herr.Message = "Can't get to the GetCacheMenusByDataAndTagsAndChecksum's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheMenusByExpireAndCreatedAndSerializedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iExpire := self.Args("expire").MustInt()
	iCreated := self.Args("created").String()
	iSerialized := self.Args("serialized").MustInt()

	if helper.IsHas(iExpire) {
		_CacheMenu, _error := model.GetCacheMenusByExpireAndCreatedAndSerialized(offset, limit, iExpire,iCreated,iSerialized)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheMenu)
	}
	herr.Message = "Can't get to the GetCacheMenusByExpireAndCreatedAndSerialized's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheMenusByExpireAndCreatedAndTagsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iExpire := self.Args("expire").MustInt()
	iCreated := self.Args("created").String()
	iTags := self.Args("tags").String()

	if helper.IsHas(iExpire) {
		_CacheMenu, _error := model.GetCacheMenusByExpireAndCreatedAndTags(offset, limit, iExpire,iCreated,iTags)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheMenu)
	}
	herr.Message = "Can't get to the GetCacheMenusByExpireAndCreatedAndTags's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheMenusByExpireAndCreatedAndChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iExpire := self.Args("expire").MustInt()
	iCreated := self.Args("created").String()
	iChecksum := self.Args("checksum").String()

	if helper.IsHas(iExpire) {
		_CacheMenu, _error := model.GetCacheMenusByExpireAndCreatedAndChecksum(offset, limit, iExpire,iCreated,iChecksum)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheMenu)
	}
	herr.Message = "Can't get to the GetCacheMenusByExpireAndCreatedAndChecksum's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheMenusByExpireAndSerializedAndTagsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iExpire := self.Args("expire").MustInt()
	iSerialized := self.Args("serialized").MustInt()
	iTags := self.Args("tags").String()

	if helper.IsHas(iExpire) {
		_CacheMenu, _error := model.GetCacheMenusByExpireAndSerializedAndTags(offset, limit, iExpire,iSerialized,iTags)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheMenu)
	}
	herr.Message = "Can't get to the GetCacheMenusByExpireAndSerializedAndTags's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheMenusByExpireAndSerializedAndChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iExpire := self.Args("expire").MustInt()
	iSerialized := self.Args("serialized").MustInt()
	iChecksum := self.Args("checksum").String()

	if helper.IsHas(iExpire) {
		_CacheMenu, _error := model.GetCacheMenusByExpireAndSerializedAndChecksum(offset, limit, iExpire,iSerialized,iChecksum)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheMenu)
	}
	herr.Message = "Can't get to the GetCacheMenusByExpireAndSerializedAndChecksum's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheMenusByExpireAndTagsAndChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iExpire := self.Args("expire").MustInt()
	iTags := self.Args("tags").String()
	iChecksum := self.Args("checksum").String()

	if helper.IsHas(iExpire) {
		_CacheMenu, _error := model.GetCacheMenusByExpireAndTagsAndChecksum(offset, limit, iExpire,iTags,iChecksum)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheMenu)
	}
	herr.Message = "Can't get to the GetCacheMenusByExpireAndTagsAndChecksum's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheMenusByCreatedAndSerializedAndTagsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCreated := self.Args("created").String()
	iSerialized := self.Args("serialized").MustInt()
	iTags := self.Args("tags").String()

	if helper.IsHas(iCreated) {
		_CacheMenu, _error := model.GetCacheMenusByCreatedAndSerializedAndTags(offset, limit, iCreated,iSerialized,iTags)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheMenu)
	}
	herr.Message = "Can't get to the GetCacheMenusByCreatedAndSerializedAndTags's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheMenusByCreatedAndSerializedAndChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCreated := self.Args("created").String()
	iSerialized := self.Args("serialized").MustInt()
	iChecksum := self.Args("checksum").String()

	if helper.IsHas(iCreated) {
		_CacheMenu, _error := model.GetCacheMenusByCreatedAndSerializedAndChecksum(offset, limit, iCreated,iSerialized,iChecksum)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheMenu)
	}
	herr.Message = "Can't get to the GetCacheMenusByCreatedAndSerializedAndChecksum's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheMenusByCreatedAndTagsAndChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCreated := self.Args("created").String()
	iTags := self.Args("tags").String()
	iChecksum := self.Args("checksum").String()

	if helper.IsHas(iCreated) {
		_CacheMenu, _error := model.GetCacheMenusByCreatedAndTagsAndChecksum(offset, limit, iCreated,iTags,iChecksum)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheMenu)
	}
	herr.Message = "Can't get to the GetCacheMenusByCreatedAndTagsAndChecksum's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheMenusBySerializedAndTagsAndChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSerialized := self.Args("serialized").MustInt()
	iTags := self.Args("tags").String()
	iChecksum := self.Args("checksum").String()

	if helper.IsHas(iSerialized) {
		_CacheMenu, _error := model.GetCacheMenusBySerializedAndTagsAndChecksum(offset, limit, iSerialized,iTags,iChecksum)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheMenu)
	}
	herr.Message = "Can't get to the GetCacheMenusBySerializedAndTagsAndChecksum's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheMenusByCidAndDataHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").String()
	iData := self.Args("data").Bytes()

	if helper.IsHas(iCid) {
		_CacheMenu, _error := model.GetCacheMenusByCidAndData(offset, limit, iCid,iData)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheMenu)
	}
	herr.Message = "Can't get to the GetCacheMenusByCidAndData's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheMenusByCidAndExpireHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").String()
	iExpire := self.Args("expire").MustInt()

	if helper.IsHas(iCid) {
		_CacheMenu, _error := model.GetCacheMenusByCidAndExpire(offset, limit, iCid,iExpire)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheMenu)
	}
	herr.Message = "Can't get to the GetCacheMenusByCidAndExpire's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheMenusByCidAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").String()
	iCreated := self.Args("created").String()

	if helper.IsHas(iCid) {
		_CacheMenu, _error := model.GetCacheMenusByCidAndCreated(offset, limit, iCid,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheMenu)
	}
	herr.Message = "Can't get to the GetCacheMenusByCidAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheMenusByCidAndSerializedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").String()
	iSerialized := self.Args("serialized").MustInt()

	if helper.IsHas(iCid) {
		_CacheMenu, _error := model.GetCacheMenusByCidAndSerialized(offset, limit, iCid,iSerialized)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheMenu)
	}
	herr.Message = "Can't get to the GetCacheMenusByCidAndSerialized's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheMenusByCidAndTagsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").String()
	iTags := self.Args("tags").String()

	if helper.IsHas(iCid) {
		_CacheMenu, _error := model.GetCacheMenusByCidAndTags(offset, limit, iCid,iTags)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheMenu)
	}
	herr.Message = "Can't get to the GetCacheMenusByCidAndTags's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheMenusByCidAndChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").String()
	iChecksum := self.Args("checksum").String()

	if helper.IsHas(iCid) {
		_CacheMenu, _error := model.GetCacheMenusByCidAndChecksum(offset, limit, iCid,iChecksum)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheMenu)
	}
	herr.Message = "Can't get to the GetCacheMenusByCidAndChecksum's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheMenusByDataAndExpireHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iData := self.Args("data").Bytes()
	iExpire := self.Args("expire").MustInt()

	if helper.IsHas(iData) {
		_CacheMenu, _error := model.GetCacheMenusByDataAndExpire(offset, limit, iData,iExpire)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheMenu)
	}
	herr.Message = "Can't get to the GetCacheMenusByDataAndExpire's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheMenusByDataAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iData := self.Args("data").Bytes()
	iCreated := self.Args("created").String()

	if helper.IsHas(iData) {
		_CacheMenu, _error := model.GetCacheMenusByDataAndCreated(offset, limit, iData,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheMenu)
	}
	herr.Message = "Can't get to the GetCacheMenusByDataAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheMenusByDataAndSerializedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iData := self.Args("data").Bytes()
	iSerialized := self.Args("serialized").MustInt()

	if helper.IsHas(iData) {
		_CacheMenu, _error := model.GetCacheMenusByDataAndSerialized(offset, limit, iData,iSerialized)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheMenu)
	}
	herr.Message = "Can't get to the GetCacheMenusByDataAndSerialized's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheMenusByDataAndTagsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iData := self.Args("data").Bytes()
	iTags := self.Args("tags").String()

	if helper.IsHas(iData) {
		_CacheMenu, _error := model.GetCacheMenusByDataAndTags(offset, limit, iData,iTags)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheMenu)
	}
	herr.Message = "Can't get to the GetCacheMenusByDataAndTags's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheMenusByDataAndChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iData := self.Args("data").Bytes()
	iChecksum := self.Args("checksum").String()

	if helper.IsHas(iData) {
		_CacheMenu, _error := model.GetCacheMenusByDataAndChecksum(offset, limit, iData,iChecksum)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheMenu)
	}
	herr.Message = "Can't get to the GetCacheMenusByDataAndChecksum's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheMenusByExpireAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iExpire := self.Args("expire").MustInt()
	iCreated := self.Args("created").String()

	if helper.IsHas(iExpire) {
		_CacheMenu, _error := model.GetCacheMenusByExpireAndCreated(offset, limit, iExpire,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheMenu)
	}
	herr.Message = "Can't get to the GetCacheMenusByExpireAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheMenusByExpireAndSerializedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iExpire := self.Args("expire").MustInt()
	iSerialized := self.Args("serialized").MustInt()

	if helper.IsHas(iExpire) {
		_CacheMenu, _error := model.GetCacheMenusByExpireAndSerialized(offset, limit, iExpire,iSerialized)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheMenu)
	}
	herr.Message = "Can't get to the GetCacheMenusByExpireAndSerialized's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheMenusByExpireAndTagsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iExpire := self.Args("expire").MustInt()
	iTags := self.Args("tags").String()

	if helper.IsHas(iExpire) {
		_CacheMenu, _error := model.GetCacheMenusByExpireAndTags(offset, limit, iExpire,iTags)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheMenu)
	}
	herr.Message = "Can't get to the GetCacheMenusByExpireAndTags's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheMenusByExpireAndChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iExpire := self.Args("expire").MustInt()
	iChecksum := self.Args("checksum").String()

	if helper.IsHas(iExpire) {
		_CacheMenu, _error := model.GetCacheMenusByExpireAndChecksum(offset, limit, iExpire,iChecksum)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheMenu)
	}
	herr.Message = "Can't get to the GetCacheMenusByExpireAndChecksum's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheMenusByCreatedAndSerializedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCreated := self.Args("created").String()
	iSerialized := self.Args("serialized").MustInt()

	if helper.IsHas(iCreated) {
		_CacheMenu, _error := model.GetCacheMenusByCreatedAndSerialized(offset, limit, iCreated,iSerialized)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheMenu)
	}
	herr.Message = "Can't get to the GetCacheMenusByCreatedAndSerialized's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheMenusByCreatedAndTagsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCreated := self.Args("created").String()
	iTags := self.Args("tags").String()

	if helper.IsHas(iCreated) {
		_CacheMenu, _error := model.GetCacheMenusByCreatedAndTags(offset, limit, iCreated,iTags)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheMenu)
	}
	herr.Message = "Can't get to the GetCacheMenusByCreatedAndTags's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheMenusByCreatedAndChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCreated := self.Args("created").String()
	iChecksum := self.Args("checksum").String()

	if helper.IsHas(iCreated) {
		_CacheMenu, _error := model.GetCacheMenusByCreatedAndChecksum(offset, limit, iCreated,iChecksum)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheMenu)
	}
	herr.Message = "Can't get to the GetCacheMenusByCreatedAndChecksum's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheMenusBySerializedAndTagsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSerialized := self.Args("serialized").MustInt()
	iTags := self.Args("tags").String()

	if helper.IsHas(iSerialized) {
		_CacheMenu, _error := model.GetCacheMenusBySerializedAndTags(offset, limit, iSerialized,iTags)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheMenu)
	}
	herr.Message = "Can't get to the GetCacheMenusBySerializedAndTags's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheMenusBySerializedAndChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSerialized := self.Args("serialized").MustInt()
	iChecksum := self.Args("checksum").String()

	if helper.IsHas(iSerialized) {
		_CacheMenu, _error := model.GetCacheMenusBySerializedAndChecksum(offset, limit, iSerialized,iChecksum)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheMenu)
	}
	herr.Message = "Can't get to the GetCacheMenusBySerializedAndChecksum's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheMenusByTagsAndChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTags := self.Args("tags").String()
	iChecksum := self.Args("checksum").String()

	if helper.IsHas(iTags) {
		_CacheMenu, _error := model.GetCacheMenusByTagsAndChecksum(offset, limit, iTags,iChecksum)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheMenu)
	}
	herr.Message = "Can't get to the GetCacheMenusByTagsAndChecksum's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheMenusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	if (offset > 0) && (len(field) > 0) {
		_CacheMenu, _error := model.GetCacheMenus(offset, limit, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheMenu)
	}
	herr.Message = "Can't get to the GetCacheMenus' args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasCacheMenuViaCidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iCid := self.Args("cid").String()
	if helper.IsHas(iCid) {
		_CacheMenu := model.HasCacheMenuViaCid(iCid)
		var m = map[string]interface{}{}
		m["cache_menu"] = _CacheMenu
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasCacheMenuViaCid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasCacheMenuViaDataHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iData := self.Args("data").Bytes()
	if helper.IsHas(iData) {
		_CacheMenu := model.HasCacheMenuViaData(iData)
		var m = map[string]interface{}{}
		m["cache_menu"] = _CacheMenu
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasCacheMenuViaData's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasCacheMenuViaExpireHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iExpire := self.Args("expire").MustInt()
	if helper.IsHas(iExpire) {
		_CacheMenu := model.HasCacheMenuViaExpire(iExpire)
		var m = map[string]interface{}{}
		m["cache_menu"] = _CacheMenu
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasCacheMenuViaExpire's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasCacheMenuViaCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iCreated := self.Args("created").String()
	if helper.IsHas(iCreated) {
		_CacheMenu := model.HasCacheMenuViaCreated(iCreated)
		var m = map[string]interface{}{}
		m["cache_menu"] = _CacheMenu
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasCacheMenuViaCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasCacheMenuViaSerializedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iSerialized := self.Args("serialized").MustInt()
	if helper.IsHas(iSerialized) {
		_CacheMenu := model.HasCacheMenuViaSerialized(iSerialized)
		var m = map[string]interface{}{}
		m["cache_menu"] = _CacheMenu
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasCacheMenuViaSerialized's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasCacheMenuViaTagsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iTags := self.Args("tags").String()
	if helper.IsHas(iTags) {
		_CacheMenu := model.HasCacheMenuViaTags(iTags)
		var m = map[string]interface{}{}
		m["cache_menu"] = _CacheMenu
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasCacheMenuViaTags's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasCacheMenuViaChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iChecksum := self.Args("checksum").String()
	if helper.IsHas(iChecksum) {
		_CacheMenu := model.HasCacheMenuViaChecksum(iChecksum)
		var m = map[string]interface{}{}
		m["cache_menu"] = _CacheMenu
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasCacheMenuViaChecksum's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheMenuViaCidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iCid := self.Args("cid").String()
	if helper.IsHas(iCid) {
		_CacheMenu, _error := model.GetCacheMenuViaCid(iCid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheMenu)
	}
	herr.Message = "Can't get to the GetCacheMenuViaCid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheMenuViaDataHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iData := self.Args("data").Bytes()
	if helper.IsHas(iData) {
		_CacheMenu, _error := model.GetCacheMenuViaData(iData)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheMenu)
	}
	herr.Message = "Can't get to the GetCacheMenuViaData's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheMenuViaExpireHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iExpire := self.Args("expire").MustInt()
	if helper.IsHas(iExpire) {
		_CacheMenu, _error := model.GetCacheMenuViaExpire(iExpire)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheMenu)
	}
	herr.Message = "Can't get to the GetCacheMenuViaExpire's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheMenuViaCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iCreated := self.Args("created").String()
	if helper.IsHas(iCreated) {
		_CacheMenu, _error := model.GetCacheMenuViaCreated(iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheMenu)
	}
	herr.Message = "Can't get to the GetCacheMenuViaCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheMenuViaSerializedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iSerialized := self.Args("serialized").MustInt()
	if helper.IsHas(iSerialized) {
		_CacheMenu, _error := model.GetCacheMenuViaSerialized(iSerialized)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheMenu)
	}
	herr.Message = "Can't get to the GetCacheMenuViaSerialized's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheMenuViaTagsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iTags := self.Args("tags").String()
	if helper.IsHas(iTags) {
		_CacheMenu, _error := model.GetCacheMenuViaTags(iTags)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheMenu)
	}
	herr.Message = "Can't get to the GetCacheMenuViaTags's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheMenuViaChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iChecksum := self.Args("checksum").String()
	if helper.IsHas(iChecksum) {
		_CacheMenu, _error := model.GetCacheMenuViaChecksum(iChecksum)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheMenu)
	}
	herr.Message = "Can't get to the GetCacheMenuViaChecksum's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetCacheMenuViaCidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Cid_ := self.Args("cid").String()
	if helper.IsHas(Cid_) {
		var iCacheMenu model.CacheMenu
		self.Bind(&iCacheMenu)
		_CacheMenu, _error := model.SetCacheMenuViaCid(Cid_, &iCacheMenu)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheMenu)
	}
	herr.Message = "Can't get to the SetCacheMenuViaCid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetCacheMenuViaDataHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Data_ := self.Args("data").Bytes()
	if helper.IsHas(Data_) {
		var iCacheMenu model.CacheMenu
		self.Bind(&iCacheMenu)
		_CacheMenu, _error := model.SetCacheMenuViaData(Data_, &iCacheMenu)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheMenu)
	}
	herr.Message = "Can't get to the SetCacheMenuViaData's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetCacheMenuViaExpireHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Expire_ := self.Args("expire").MustInt()
	if helper.IsHas(Expire_) {
		var iCacheMenu model.CacheMenu
		self.Bind(&iCacheMenu)
		_CacheMenu, _error := model.SetCacheMenuViaExpire(Expire_, &iCacheMenu)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheMenu)
	}
	herr.Message = "Can't get to the SetCacheMenuViaExpire's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetCacheMenuViaCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Created_ := self.Args("created").String()
	if helper.IsHas(Created_) {
		var iCacheMenu model.CacheMenu
		self.Bind(&iCacheMenu)
		_CacheMenu, _error := model.SetCacheMenuViaCreated(Created_, &iCacheMenu)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheMenu)
	}
	herr.Message = "Can't get to the SetCacheMenuViaCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetCacheMenuViaSerializedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Serialized_ := self.Args("serialized").MustInt()
	if helper.IsHas(Serialized_) {
		var iCacheMenu model.CacheMenu
		self.Bind(&iCacheMenu)
		_CacheMenu, _error := model.SetCacheMenuViaSerialized(Serialized_, &iCacheMenu)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheMenu)
	}
	herr.Message = "Can't get to the SetCacheMenuViaSerialized's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetCacheMenuViaTagsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Tags_ := self.Args("tags").String()
	if helper.IsHas(Tags_) {
		var iCacheMenu model.CacheMenu
		self.Bind(&iCacheMenu)
		_CacheMenu, _error := model.SetCacheMenuViaTags(Tags_, &iCacheMenu)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheMenu)
	}
	herr.Message = "Can't get to the SetCacheMenuViaTags's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetCacheMenuViaChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Checksum_ := self.Args("checksum").String()
	if helper.IsHas(Checksum_) {
		var iCacheMenu model.CacheMenu
		self.Bind(&iCacheMenu)
		_CacheMenu, _error := model.SetCacheMenuViaChecksum(Checksum_, &iCacheMenu)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheMenu)
	}
	herr.Message = "Can't get to the SetCacheMenuViaChecksum's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostAddCacheMenuHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Cid_ := self.Args("cid").String()
	Data_ := self.Args("data").Bytes()
	Expire_ := self.Args("expire").MustInt()
	Created_ := self.Args("created").String()
	Serialized_ := self.Args("serialized").MustInt()
	Tags_ := self.Args("tags").String()
	Checksum_ := self.Args("checksum").String()

	if helper.IsHas(Cid_) {
		_error := model.AddCacheMenu(Cid_,Data_,Expire_,Created_,Serialized_,Tags_,Checksum_)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		herr.Message = "StatusOK"
		herr.Status = macross.StatusOK
		return self.JSON(herr)
	}
	herr.Message = "Can't get to the AddCacheMenu's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostCacheMenuHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iCacheMenu model.CacheMenu
	self.Bind(&iCacheMenu)
	_string, _error := model.PostCacheMenu(&iCacheMenu)
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

func PutCacheMenuHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iCacheMenu model.CacheMenu
	self.Bind(&iCacheMenu)
	_string, _error := model.PutCacheMenu(&iCacheMenu)
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

func PutCacheMenuViaCidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Cid_ := self.Args("cid").String()
	var iCacheMenu model.CacheMenu
	self.Bind(&iCacheMenu)
	_int64, _error := model.PutCacheMenuViaCid(Cid_, &iCacheMenu)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutCacheMenuViaDataHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Data_ := self.Args("data").Bytes()
	var iCacheMenu model.CacheMenu
	self.Bind(&iCacheMenu)
	_int64, _error := model.PutCacheMenuViaData(Data_, &iCacheMenu)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutCacheMenuViaExpireHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Expire_ := self.Args("expire").MustInt()
	var iCacheMenu model.CacheMenu
	self.Bind(&iCacheMenu)
	_int64, _error := model.PutCacheMenuViaExpire(Expire_, &iCacheMenu)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutCacheMenuViaCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Created_ := self.Args("created").String()
	var iCacheMenu model.CacheMenu
	self.Bind(&iCacheMenu)
	_int64, _error := model.PutCacheMenuViaCreated(Created_, &iCacheMenu)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutCacheMenuViaSerializedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Serialized_ := self.Args("serialized").MustInt()
	var iCacheMenu model.CacheMenu
	self.Bind(&iCacheMenu)
	_int64, _error := model.PutCacheMenuViaSerialized(Serialized_, &iCacheMenu)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutCacheMenuViaTagsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Tags_ := self.Args("tags").String()
	var iCacheMenu model.CacheMenu
	self.Bind(&iCacheMenu)
	_int64, _error := model.PutCacheMenuViaTags(Tags_, &iCacheMenu)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutCacheMenuViaChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Checksum_ := self.Args("checksum").String()
	var iCacheMenu model.CacheMenu
	self.Bind(&iCacheMenu)
	_int64, _error := model.PutCacheMenuViaChecksum(Checksum_, &iCacheMenu)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateCacheMenuViaCidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Cid_ := self.Args("cid").String()
	var iCacheMenu model.CacheMenu
	self.Bind(&iCacheMenu)
	var iMap = helper.StructToMap(iCacheMenu)
	_error := model.UpdateCacheMenuViaCid(Cid_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateCacheMenuViaDataHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Data_ := self.Args("data").Bytes()
	var iCacheMenu model.CacheMenu
	self.Bind(&iCacheMenu)
	var iMap = helper.StructToMap(iCacheMenu)
	_error := model.UpdateCacheMenuViaData(Data_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateCacheMenuViaExpireHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Expire_ := self.Args("expire").MustInt()
	var iCacheMenu model.CacheMenu
	self.Bind(&iCacheMenu)
	var iMap = helper.StructToMap(iCacheMenu)
	_error := model.UpdateCacheMenuViaExpire(Expire_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateCacheMenuViaCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Created_ := self.Args("created").String()
	var iCacheMenu model.CacheMenu
	self.Bind(&iCacheMenu)
	var iMap = helper.StructToMap(iCacheMenu)
	_error := model.UpdateCacheMenuViaCreated(Created_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateCacheMenuViaSerializedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Serialized_ := self.Args("serialized").MustInt()
	var iCacheMenu model.CacheMenu
	self.Bind(&iCacheMenu)
	var iMap = helper.StructToMap(iCacheMenu)
	_error := model.UpdateCacheMenuViaSerialized(Serialized_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateCacheMenuViaTagsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Tags_ := self.Args("tags").String()
	var iCacheMenu model.CacheMenu
	self.Bind(&iCacheMenu)
	var iMap = helper.StructToMap(iCacheMenu)
	_error := model.UpdateCacheMenuViaTags(Tags_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateCacheMenuViaChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Checksum_ := self.Args("checksum").String()
	var iCacheMenu model.CacheMenu
	self.Bind(&iCacheMenu)
	var iMap = helper.StructToMap(iCacheMenu)
	_error := model.UpdateCacheMenuViaChecksum(Checksum_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteCacheMenuHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Cid_ := self.Args("cid").String()
	_int64, _error := model.DeleteCacheMenu(Cid_)
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

func DeleteCacheMenuViaCidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Cid_ := self.Args("cid").String()
	_error := model.DeleteCacheMenuViaCid(Cid_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteCacheMenuViaDataHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Data_ := self.Args("data").Bytes()
	_error := model.DeleteCacheMenuViaData(Data_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteCacheMenuViaExpireHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Expire_ := self.Args("expire").MustInt()
	_error := model.DeleteCacheMenuViaExpire(Expire_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteCacheMenuViaCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Created_ := self.Args("created").String()
	_error := model.DeleteCacheMenuViaCreated(Created_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteCacheMenuViaSerializedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Serialized_ := self.Args("serialized").MustInt()
	_error := model.DeleteCacheMenuViaSerialized(Serialized_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteCacheMenuViaTagsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Tags_ := self.Args("tags").String()
	_error := model.DeleteCacheMenuViaTags(Tags_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteCacheMenuViaChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Checksum_ := self.Args("checksum").String()
	_error := model.DeleteCacheMenuViaChecksum(Checksum_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}
