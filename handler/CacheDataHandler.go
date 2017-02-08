package handler

import (
	"drupapp/helper"
	"drupapp/model"
	"github.com/insionng/macross"
)

func GetCacheDatasCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()

	if offset > 0 {
		var m = map[string]interface{}{}
		_int64, _error := model.GetCacheDatasCount(offset, limit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		m["cache_datasCount"] = _int64
		return self.JSON(m)
	}
	herr.Message = "Can't get to the GetCacheDatasCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDataCountViaCidHandler(self *macross.Context) error {
	Cid_ := self.Args("cid").String()
	_int64 := model.GetCacheDataCountViaCid(Cid_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["cache_dataCount"] = 0
	}
	m["cache_dataCount"] = _int64
	return self.JSON(m)
}

func GetCacheDataCountViaDataHandler(self *macross.Context) error {
	Data_ := self.Args("data").Bytes()
	_int64 := model.GetCacheDataCountViaData(Data_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["cache_dataCount"] = 0
	}
	m["cache_dataCount"] = _int64
	return self.JSON(m)
}

func GetCacheDataCountViaExpireHandler(self *macross.Context) error {
	Expire_ := self.Args("expire").MustInt()
	_int64 := model.GetCacheDataCountViaExpire(Expire_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["cache_dataCount"] = 0
	}
	m["cache_dataCount"] = _int64
	return self.JSON(m)
}

func GetCacheDataCountViaCreatedHandler(self *macross.Context) error {
	Created_ := self.Args("created").String()
	_int64 := model.GetCacheDataCountViaCreated(Created_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["cache_dataCount"] = 0
	}
	m["cache_dataCount"] = _int64
	return self.JSON(m)
}

func GetCacheDataCountViaSerializedHandler(self *macross.Context) error {
	Serialized_ := self.Args("serialized").MustInt()
	_int64 := model.GetCacheDataCountViaSerialized(Serialized_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["cache_dataCount"] = 0
	}
	m["cache_dataCount"] = _int64
	return self.JSON(m)
}

func GetCacheDataCountViaTagsHandler(self *macross.Context) error {
	Tags_ := self.Args("tags").String()
	_int64 := model.GetCacheDataCountViaTags(Tags_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["cache_dataCount"] = 0
	}
	m["cache_dataCount"] = _int64
	return self.JSON(m)
}

func GetCacheDataCountViaChecksumHandler(self *macross.Context) error {
	Checksum_ := self.Args("checksum").String()
	_int64 := model.GetCacheDataCountViaChecksum(Checksum_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["cache_dataCount"] = 0
	}
	m["cache_dataCount"] = _int64
	return self.JSON(m)
}

func GetCacheDatasViaCidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iCid := self.Args("cid").String()
	if (offset > 0) && helper.IsHas(iCid) {
		_CacheData, _error := model.GetCacheDatasViaCid(offset, limit, iCid, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheData)
	}
	herr.Message = "Can't get to the GetCacheDatasViaCid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDatasViaDataHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iData := self.Args("data").Bytes()
	if (offset > 0) && helper.IsHas(iData) {
		_CacheData, _error := model.GetCacheDatasViaData(offset, limit, iData, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheData)
	}
	herr.Message = "Can't get to the GetCacheDatasViaData's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDatasViaExpireHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iExpire := self.Args("expire").MustInt()
	if (offset > 0) && helper.IsHas(iExpire) {
		_CacheData, _error := model.GetCacheDatasViaExpire(offset, limit, iExpire, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheData)
	}
	herr.Message = "Can't get to the GetCacheDatasViaExpire's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDatasViaCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iCreated := self.Args("created").String()
	if (offset > 0) && helper.IsHas(iCreated) {
		_CacheData, _error := model.GetCacheDatasViaCreated(offset, limit, iCreated, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheData)
	}
	herr.Message = "Can't get to the GetCacheDatasViaCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDatasViaSerializedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iSerialized := self.Args("serialized").MustInt()
	if (offset > 0) && helper.IsHas(iSerialized) {
		_CacheData, _error := model.GetCacheDatasViaSerialized(offset, limit, iSerialized, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheData)
	}
	herr.Message = "Can't get to the GetCacheDatasViaSerialized's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDatasViaTagsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iTags := self.Args("tags").String()
	if (offset > 0) && helper.IsHas(iTags) {
		_CacheData, _error := model.GetCacheDatasViaTags(offset, limit, iTags, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheData)
	}
	herr.Message = "Can't get to the GetCacheDatasViaTags's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDatasViaChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iChecksum := self.Args("checksum").String()
	if (offset > 0) && helper.IsHas(iChecksum) {
		_CacheData, _error := model.GetCacheDatasViaChecksum(offset, limit, iChecksum, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheData)
	}
	herr.Message = "Can't get to the GetCacheDatasViaChecksum's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDatasByCidAndDataAndExpireHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").String()
	iData := self.Args("data").Bytes()
	iExpire := self.Args("expire").MustInt()

	if helper.IsHas(iCid) {
		_CacheData, _error := model.GetCacheDatasByCidAndDataAndExpire(offset, limit, iCid,iData,iExpire)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheData)
	}
	herr.Message = "Can't get to the GetCacheDatasByCidAndDataAndExpire's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDatasByCidAndDataAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").String()
	iData := self.Args("data").Bytes()
	iCreated := self.Args("created").String()

	if helper.IsHas(iCid) {
		_CacheData, _error := model.GetCacheDatasByCidAndDataAndCreated(offset, limit, iCid,iData,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheData)
	}
	herr.Message = "Can't get to the GetCacheDatasByCidAndDataAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDatasByCidAndDataAndSerializedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").String()
	iData := self.Args("data").Bytes()
	iSerialized := self.Args("serialized").MustInt()

	if helper.IsHas(iCid) {
		_CacheData, _error := model.GetCacheDatasByCidAndDataAndSerialized(offset, limit, iCid,iData,iSerialized)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheData)
	}
	herr.Message = "Can't get to the GetCacheDatasByCidAndDataAndSerialized's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDatasByCidAndDataAndTagsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").String()
	iData := self.Args("data").Bytes()
	iTags := self.Args("tags").String()

	if helper.IsHas(iCid) {
		_CacheData, _error := model.GetCacheDatasByCidAndDataAndTags(offset, limit, iCid,iData,iTags)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheData)
	}
	herr.Message = "Can't get to the GetCacheDatasByCidAndDataAndTags's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDatasByCidAndDataAndChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").String()
	iData := self.Args("data").Bytes()
	iChecksum := self.Args("checksum").String()

	if helper.IsHas(iCid) {
		_CacheData, _error := model.GetCacheDatasByCidAndDataAndChecksum(offset, limit, iCid,iData,iChecksum)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheData)
	}
	herr.Message = "Can't get to the GetCacheDatasByCidAndDataAndChecksum's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDatasByCidAndExpireAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").String()
	iExpire := self.Args("expire").MustInt()
	iCreated := self.Args("created").String()

	if helper.IsHas(iCid) {
		_CacheData, _error := model.GetCacheDatasByCidAndExpireAndCreated(offset, limit, iCid,iExpire,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheData)
	}
	herr.Message = "Can't get to the GetCacheDatasByCidAndExpireAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDatasByCidAndExpireAndSerializedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").String()
	iExpire := self.Args("expire").MustInt()
	iSerialized := self.Args("serialized").MustInt()

	if helper.IsHas(iCid) {
		_CacheData, _error := model.GetCacheDatasByCidAndExpireAndSerialized(offset, limit, iCid,iExpire,iSerialized)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheData)
	}
	herr.Message = "Can't get to the GetCacheDatasByCidAndExpireAndSerialized's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDatasByCidAndExpireAndTagsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").String()
	iExpire := self.Args("expire").MustInt()
	iTags := self.Args("tags").String()

	if helper.IsHas(iCid) {
		_CacheData, _error := model.GetCacheDatasByCidAndExpireAndTags(offset, limit, iCid,iExpire,iTags)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheData)
	}
	herr.Message = "Can't get to the GetCacheDatasByCidAndExpireAndTags's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDatasByCidAndExpireAndChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").String()
	iExpire := self.Args("expire").MustInt()
	iChecksum := self.Args("checksum").String()

	if helper.IsHas(iCid) {
		_CacheData, _error := model.GetCacheDatasByCidAndExpireAndChecksum(offset, limit, iCid,iExpire,iChecksum)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheData)
	}
	herr.Message = "Can't get to the GetCacheDatasByCidAndExpireAndChecksum's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDatasByCidAndCreatedAndSerializedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").String()
	iCreated := self.Args("created").String()
	iSerialized := self.Args("serialized").MustInt()

	if helper.IsHas(iCid) {
		_CacheData, _error := model.GetCacheDatasByCidAndCreatedAndSerialized(offset, limit, iCid,iCreated,iSerialized)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheData)
	}
	herr.Message = "Can't get to the GetCacheDatasByCidAndCreatedAndSerialized's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDatasByCidAndCreatedAndTagsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").String()
	iCreated := self.Args("created").String()
	iTags := self.Args("tags").String()

	if helper.IsHas(iCid) {
		_CacheData, _error := model.GetCacheDatasByCidAndCreatedAndTags(offset, limit, iCid,iCreated,iTags)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheData)
	}
	herr.Message = "Can't get to the GetCacheDatasByCidAndCreatedAndTags's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDatasByCidAndCreatedAndChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").String()
	iCreated := self.Args("created").String()
	iChecksum := self.Args("checksum").String()

	if helper.IsHas(iCid) {
		_CacheData, _error := model.GetCacheDatasByCidAndCreatedAndChecksum(offset, limit, iCid,iCreated,iChecksum)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheData)
	}
	herr.Message = "Can't get to the GetCacheDatasByCidAndCreatedAndChecksum's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDatasByCidAndSerializedAndTagsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").String()
	iSerialized := self.Args("serialized").MustInt()
	iTags := self.Args("tags").String()

	if helper.IsHas(iCid) {
		_CacheData, _error := model.GetCacheDatasByCidAndSerializedAndTags(offset, limit, iCid,iSerialized,iTags)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheData)
	}
	herr.Message = "Can't get to the GetCacheDatasByCidAndSerializedAndTags's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDatasByCidAndSerializedAndChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").String()
	iSerialized := self.Args("serialized").MustInt()
	iChecksum := self.Args("checksum").String()

	if helper.IsHas(iCid) {
		_CacheData, _error := model.GetCacheDatasByCidAndSerializedAndChecksum(offset, limit, iCid,iSerialized,iChecksum)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheData)
	}
	herr.Message = "Can't get to the GetCacheDatasByCidAndSerializedAndChecksum's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDatasByCidAndTagsAndChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").String()
	iTags := self.Args("tags").String()
	iChecksum := self.Args("checksum").String()

	if helper.IsHas(iCid) {
		_CacheData, _error := model.GetCacheDatasByCidAndTagsAndChecksum(offset, limit, iCid,iTags,iChecksum)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheData)
	}
	herr.Message = "Can't get to the GetCacheDatasByCidAndTagsAndChecksum's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDatasByDataAndExpireAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iData := self.Args("data").Bytes()
	iExpire := self.Args("expire").MustInt()
	iCreated := self.Args("created").String()

	if helper.IsHas(iData) {
		_CacheData, _error := model.GetCacheDatasByDataAndExpireAndCreated(offset, limit, iData,iExpire,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheData)
	}
	herr.Message = "Can't get to the GetCacheDatasByDataAndExpireAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDatasByDataAndExpireAndSerializedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iData := self.Args("data").Bytes()
	iExpire := self.Args("expire").MustInt()
	iSerialized := self.Args("serialized").MustInt()

	if helper.IsHas(iData) {
		_CacheData, _error := model.GetCacheDatasByDataAndExpireAndSerialized(offset, limit, iData,iExpire,iSerialized)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheData)
	}
	herr.Message = "Can't get to the GetCacheDatasByDataAndExpireAndSerialized's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDatasByDataAndExpireAndTagsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iData := self.Args("data").Bytes()
	iExpire := self.Args("expire").MustInt()
	iTags := self.Args("tags").String()

	if helper.IsHas(iData) {
		_CacheData, _error := model.GetCacheDatasByDataAndExpireAndTags(offset, limit, iData,iExpire,iTags)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheData)
	}
	herr.Message = "Can't get to the GetCacheDatasByDataAndExpireAndTags's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDatasByDataAndExpireAndChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iData := self.Args("data").Bytes()
	iExpire := self.Args("expire").MustInt()
	iChecksum := self.Args("checksum").String()

	if helper.IsHas(iData) {
		_CacheData, _error := model.GetCacheDatasByDataAndExpireAndChecksum(offset, limit, iData,iExpire,iChecksum)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheData)
	}
	herr.Message = "Can't get to the GetCacheDatasByDataAndExpireAndChecksum's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDatasByDataAndCreatedAndSerializedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iData := self.Args("data").Bytes()
	iCreated := self.Args("created").String()
	iSerialized := self.Args("serialized").MustInt()

	if helper.IsHas(iData) {
		_CacheData, _error := model.GetCacheDatasByDataAndCreatedAndSerialized(offset, limit, iData,iCreated,iSerialized)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheData)
	}
	herr.Message = "Can't get to the GetCacheDatasByDataAndCreatedAndSerialized's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDatasByDataAndCreatedAndTagsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iData := self.Args("data").Bytes()
	iCreated := self.Args("created").String()
	iTags := self.Args("tags").String()

	if helper.IsHas(iData) {
		_CacheData, _error := model.GetCacheDatasByDataAndCreatedAndTags(offset, limit, iData,iCreated,iTags)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheData)
	}
	herr.Message = "Can't get to the GetCacheDatasByDataAndCreatedAndTags's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDatasByDataAndCreatedAndChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iData := self.Args("data").Bytes()
	iCreated := self.Args("created").String()
	iChecksum := self.Args("checksum").String()

	if helper.IsHas(iData) {
		_CacheData, _error := model.GetCacheDatasByDataAndCreatedAndChecksum(offset, limit, iData,iCreated,iChecksum)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheData)
	}
	herr.Message = "Can't get to the GetCacheDatasByDataAndCreatedAndChecksum's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDatasByDataAndSerializedAndTagsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iData := self.Args("data").Bytes()
	iSerialized := self.Args("serialized").MustInt()
	iTags := self.Args("tags").String()

	if helper.IsHas(iData) {
		_CacheData, _error := model.GetCacheDatasByDataAndSerializedAndTags(offset, limit, iData,iSerialized,iTags)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheData)
	}
	herr.Message = "Can't get to the GetCacheDatasByDataAndSerializedAndTags's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDatasByDataAndSerializedAndChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iData := self.Args("data").Bytes()
	iSerialized := self.Args("serialized").MustInt()
	iChecksum := self.Args("checksum").String()

	if helper.IsHas(iData) {
		_CacheData, _error := model.GetCacheDatasByDataAndSerializedAndChecksum(offset, limit, iData,iSerialized,iChecksum)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheData)
	}
	herr.Message = "Can't get to the GetCacheDatasByDataAndSerializedAndChecksum's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDatasByDataAndTagsAndChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iData := self.Args("data").Bytes()
	iTags := self.Args("tags").String()
	iChecksum := self.Args("checksum").String()

	if helper.IsHas(iData) {
		_CacheData, _error := model.GetCacheDatasByDataAndTagsAndChecksum(offset, limit, iData,iTags,iChecksum)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheData)
	}
	herr.Message = "Can't get to the GetCacheDatasByDataAndTagsAndChecksum's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDatasByExpireAndCreatedAndSerializedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iExpire := self.Args("expire").MustInt()
	iCreated := self.Args("created").String()
	iSerialized := self.Args("serialized").MustInt()

	if helper.IsHas(iExpire) {
		_CacheData, _error := model.GetCacheDatasByExpireAndCreatedAndSerialized(offset, limit, iExpire,iCreated,iSerialized)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheData)
	}
	herr.Message = "Can't get to the GetCacheDatasByExpireAndCreatedAndSerialized's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDatasByExpireAndCreatedAndTagsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iExpire := self.Args("expire").MustInt()
	iCreated := self.Args("created").String()
	iTags := self.Args("tags").String()

	if helper.IsHas(iExpire) {
		_CacheData, _error := model.GetCacheDatasByExpireAndCreatedAndTags(offset, limit, iExpire,iCreated,iTags)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheData)
	}
	herr.Message = "Can't get to the GetCacheDatasByExpireAndCreatedAndTags's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDatasByExpireAndCreatedAndChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iExpire := self.Args("expire").MustInt()
	iCreated := self.Args("created").String()
	iChecksum := self.Args("checksum").String()

	if helper.IsHas(iExpire) {
		_CacheData, _error := model.GetCacheDatasByExpireAndCreatedAndChecksum(offset, limit, iExpire,iCreated,iChecksum)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheData)
	}
	herr.Message = "Can't get to the GetCacheDatasByExpireAndCreatedAndChecksum's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDatasByExpireAndSerializedAndTagsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iExpire := self.Args("expire").MustInt()
	iSerialized := self.Args("serialized").MustInt()
	iTags := self.Args("tags").String()

	if helper.IsHas(iExpire) {
		_CacheData, _error := model.GetCacheDatasByExpireAndSerializedAndTags(offset, limit, iExpire,iSerialized,iTags)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheData)
	}
	herr.Message = "Can't get to the GetCacheDatasByExpireAndSerializedAndTags's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDatasByExpireAndSerializedAndChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iExpire := self.Args("expire").MustInt()
	iSerialized := self.Args("serialized").MustInt()
	iChecksum := self.Args("checksum").String()

	if helper.IsHas(iExpire) {
		_CacheData, _error := model.GetCacheDatasByExpireAndSerializedAndChecksum(offset, limit, iExpire,iSerialized,iChecksum)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheData)
	}
	herr.Message = "Can't get to the GetCacheDatasByExpireAndSerializedAndChecksum's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDatasByExpireAndTagsAndChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iExpire := self.Args("expire").MustInt()
	iTags := self.Args("tags").String()
	iChecksum := self.Args("checksum").String()

	if helper.IsHas(iExpire) {
		_CacheData, _error := model.GetCacheDatasByExpireAndTagsAndChecksum(offset, limit, iExpire,iTags,iChecksum)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheData)
	}
	herr.Message = "Can't get to the GetCacheDatasByExpireAndTagsAndChecksum's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDatasByCreatedAndSerializedAndTagsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCreated := self.Args("created").String()
	iSerialized := self.Args("serialized").MustInt()
	iTags := self.Args("tags").String()

	if helper.IsHas(iCreated) {
		_CacheData, _error := model.GetCacheDatasByCreatedAndSerializedAndTags(offset, limit, iCreated,iSerialized,iTags)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheData)
	}
	herr.Message = "Can't get to the GetCacheDatasByCreatedAndSerializedAndTags's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDatasByCreatedAndSerializedAndChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCreated := self.Args("created").String()
	iSerialized := self.Args("serialized").MustInt()
	iChecksum := self.Args("checksum").String()

	if helper.IsHas(iCreated) {
		_CacheData, _error := model.GetCacheDatasByCreatedAndSerializedAndChecksum(offset, limit, iCreated,iSerialized,iChecksum)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheData)
	}
	herr.Message = "Can't get to the GetCacheDatasByCreatedAndSerializedAndChecksum's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDatasByCreatedAndTagsAndChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCreated := self.Args("created").String()
	iTags := self.Args("tags").String()
	iChecksum := self.Args("checksum").String()

	if helper.IsHas(iCreated) {
		_CacheData, _error := model.GetCacheDatasByCreatedAndTagsAndChecksum(offset, limit, iCreated,iTags,iChecksum)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheData)
	}
	herr.Message = "Can't get to the GetCacheDatasByCreatedAndTagsAndChecksum's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDatasBySerializedAndTagsAndChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSerialized := self.Args("serialized").MustInt()
	iTags := self.Args("tags").String()
	iChecksum := self.Args("checksum").String()

	if helper.IsHas(iSerialized) {
		_CacheData, _error := model.GetCacheDatasBySerializedAndTagsAndChecksum(offset, limit, iSerialized,iTags,iChecksum)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheData)
	}
	herr.Message = "Can't get to the GetCacheDatasBySerializedAndTagsAndChecksum's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDatasByCidAndDataHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").String()
	iData := self.Args("data").Bytes()

	if helper.IsHas(iCid) {
		_CacheData, _error := model.GetCacheDatasByCidAndData(offset, limit, iCid,iData)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheData)
	}
	herr.Message = "Can't get to the GetCacheDatasByCidAndData's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDatasByCidAndExpireHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").String()
	iExpire := self.Args("expire").MustInt()

	if helper.IsHas(iCid) {
		_CacheData, _error := model.GetCacheDatasByCidAndExpire(offset, limit, iCid,iExpire)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheData)
	}
	herr.Message = "Can't get to the GetCacheDatasByCidAndExpire's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDatasByCidAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").String()
	iCreated := self.Args("created").String()

	if helper.IsHas(iCid) {
		_CacheData, _error := model.GetCacheDatasByCidAndCreated(offset, limit, iCid,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheData)
	}
	herr.Message = "Can't get to the GetCacheDatasByCidAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDatasByCidAndSerializedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").String()
	iSerialized := self.Args("serialized").MustInt()

	if helper.IsHas(iCid) {
		_CacheData, _error := model.GetCacheDatasByCidAndSerialized(offset, limit, iCid,iSerialized)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheData)
	}
	herr.Message = "Can't get to the GetCacheDatasByCidAndSerialized's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDatasByCidAndTagsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").String()
	iTags := self.Args("tags").String()

	if helper.IsHas(iCid) {
		_CacheData, _error := model.GetCacheDatasByCidAndTags(offset, limit, iCid,iTags)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheData)
	}
	herr.Message = "Can't get to the GetCacheDatasByCidAndTags's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDatasByCidAndChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").String()
	iChecksum := self.Args("checksum").String()

	if helper.IsHas(iCid) {
		_CacheData, _error := model.GetCacheDatasByCidAndChecksum(offset, limit, iCid,iChecksum)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheData)
	}
	herr.Message = "Can't get to the GetCacheDatasByCidAndChecksum's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDatasByDataAndExpireHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iData := self.Args("data").Bytes()
	iExpire := self.Args("expire").MustInt()

	if helper.IsHas(iData) {
		_CacheData, _error := model.GetCacheDatasByDataAndExpire(offset, limit, iData,iExpire)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheData)
	}
	herr.Message = "Can't get to the GetCacheDatasByDataAndExpire's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDatasByDataAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iData := self.Args("data").Bytes()
	iCreated := self.Args("created").String()

	if helper.IsHas(iData) {
		_CacheData, _error := model.GetCacheDatasByDataAndCreated(offset, limit, iData,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheData)
	}
	herr.Message = "Can't get to the GetCacheDatasByDataAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDatasByDataAndSerializedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iData := self.Args("data").Bytes()
	iSerialized := self.Args("serialized").MustInt()

	if helper.IsHas(iData) {
		_CacheData, _error := model.GetCacheDatasByDataAndSerialized(offset, limit, iData,iSerialized)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheData)
	}
	herr.Message = "Can't get to the GetCacheDatasByDataAndSerialized's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDatasByDataAndTagsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iData := self.Args("data").Bytes()
	iTags := self.Args("tags").String()

	if helper.IsHas(iData) {
		_CacheData, _error := model.GetCacheDatasByDataAndTags(offset, limit, iData,iTags)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheData)
	}
	herr.Message = "Can't get to the GetCacheDatasByDataAndTags's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDatasByDataAndChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iData := self.Args("data").Bytes()
	iChecksum := self.Args("checksum").String()

	if helper.IsHas(iData) {
		_CacheData, _error := model.GetCacheDatasByDataAndChecksum(offset, limit, iData,iChecksum)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheData)
	}
	herr.Message = "Can't get to the GetCacheDatasByDataAndChecksum's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDatasByExpireAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iExpire := self.Args("expire").MustInt()
	iCreated := self.Args("created").String()

	if helper.IsHas(iExpire) {
		_CacheData, _error := model.GetCacheDatasByExpireAndCreated(offset, limit, iExpire,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheData)
	}
	herr.Message = "Can't get to the GetCacheDatasByExpireAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDatasByExpireAndSerializedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iExpire := self.Args("expire").MustInt()
	iSerialized := self.Args("serialized").MustInt()

	if helper.IsHas(iExpire) {
		_CacheData, _error := model.GetCacheDatasByExpireAndSerialized(offset, limit, iExpire,iSerialized)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheData)
	}
	herr.Message = "Can't get to the GetCacheDatasByExpireAndSerialized's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDatasByExpireAndTagsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iExpire := self.Args("expire").MustInt()
	iTags := self.Args("tags").String()

	if helper.IsHas(iExpire) {
		_CacheData, _error := model.GetCacheDatasByExpireAndTags(offset, limit, iExpire,iTags)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheData)
	}
	herr.Message = "Can't get to the GetCacheDatasByExpireAndTags's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDatasByExpireAndChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iExpire := self.Args("expire").MustInt()
	iChecksum := self.Args("checksum").String()

	if helper.IsHas(iExpire) {
		_CacheData, _error := model.GetCacheDatasByExpireAndChecksum(offset, limit, iExpire,iChecksum)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheData)
	}
	herr.Message = "Can't get to the GetCacheDatasByExpireAndChecksum's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDatasByCreatedAndSerializedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCreated := self.Args("created").String()
	iSerialized := self.Args("serialized").MustInt()

	if helper.IsHas(iCreated) {
		_CacheData, _error := model.GetCacheDatasByCreatedAndSerialized(offset, limit, iCreated,iSerialized)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheData)
	}
	herr.Message = "Can't get to the GetCacheDatasByCreatedAndSerialized's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDatasByCreatedAndTagsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCreated := self.Args("created").String()
	iTags := self.Args("tags").String()

	if helper.IsHas(iCreated) {
		_CacheData, _error := model.GetCacheDatasByCreatedAndTags(offset, limit, iCreated,iTags)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheData)
	}
	herr.Message = "Can't get to the GetCacheDatasByCreatedAndTags's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDatasByCreatedAndChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCreated := self.Args("created").String()
	iChecksum := self.Args("checksum").String()

	if helper.IsHas(iCreated) {
		_CacheData, _error := model.GetCacheDatasByCreatedAndChecksum(offset, limit, iCreated,iChecksum)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheData)
	}
	herr.Message = "Can't get to the GetCacheDatasByCreatedAndChecksum's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDatasBySerializedAndTagsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSerialized := self.Args("serialized").MustInt()
	iTags := self.Args("tags").String()

	if helper.IsHas(iSerialized) {
		_CacheData, _error := model.GetCacheDatasBySerializedAndTags(offset, limit, iSerialized,iTags)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheData)
	}
	herr.Message = "Can't get to the GetCacheDatasBySerializedAndTags's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDatasBySerializedAndChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSerialized := self.Args("serialized").MustInt()
	iChecksum := self.Args("checksum").String()

	if helper.IsHas(iSerialized) {
		_CacheData, _error := model.GetCacheDatasBySerializedAndChecksum(offset, limit, iSerialized,iChecksum)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheData)
	}
	herr.Message = "Can't get to the GetCacheDatasBySerializedAndChecksum's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDatasByTagsAndChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTags := self.Args("tags").String()
	iChecksum := self.Args("checksum").String()

	if helper.IsHas(iTags) {
		_CacheData, _error := model.GetCacheDatasByTagsAndChecksum(offset, limit, iTags,iChecksum)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheData)
	}
	herr.Message = "Can't get to the GetCacheDatasByTagsAndChecksum's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDatasHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	if (offset > 0) && (len(field) > 0) {
		_CacheData, _error := model.GetCacheDatas(offset, limit, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheData)
	}
	herr.Message = "Can't get to the GetCacheDatas' args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasCacheDataViaCidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iCid := self.Args("cid").String()
	if helper.IsHas(iCid) {
		_CacheData := model.HasCacheDataViaCid(iCid)
		var m = map[string]interface{}{}
		m["cache_data"] = _CacheData
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasCacheDataViaCid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasCacheDataViaDataHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iData := self.Args("data").Bytes()
	if helper.IsHas(iData) {
		_CacheData := model.HasCacheDataViaData(iData)
		var m = map[string]interface{}{}
		m["cache_data"] = _CacheData
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasCacheDataViaData's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasCacheDataViaExpireHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iExpire := self.Args("expire").MustInt()
	if helper.IsHas(iExpire) {
		_CacheData := model.HasCacheDataViaExpire(iExpire)
		var m = map[string]interface{}{}
		m["cache_data"] = _CacheData
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasCacheDataViaExpire's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasCacheDataViaCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iCreated := self.Args("created").String()
	if helper.IsHas(iCreated) {
		_CacheData := model.HasCacheDataViaCreated(iCreated)
		var m = map[string]interface{}{}
		m["cache_data"] = _CacheData
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasCacheDataViaCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasCacheDataViaSerializedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iSerialized := self.Args("serialized").MustInt()
	if helper.IsHas(iSerialized) {
		_CacheData := model.HasCacheDataViaSerialized(iSerialized)
		var m = map[string]interface{}{}
		m["cache_data"] = _CacheData
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasCacheDataViaSerialized's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasCacheDataViaTagsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iTags := self.Args("tags").String()
	if helper.IsHas(iTags) {
		_CacheData := model.HasCacheDataViaTags(iTags)
		var m = map[string]interface{}{}
		m["cache_data"] = _CacheData
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasCacheDataViaTags's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasCacheDataViaChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iChecksum := self.Args("checksum").String()
	if helper.IsHas(iChecksum) {
		_CacheData := model.HasCacheDataViaChecksum(iChecksum)
		var m = map[string]interface{}{}
		m["cache_data"] = _CacheData
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasCacheDataViaChecksum's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDataViaCidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iCid := self.Args("cid").String()
	if helper.IsHas(iCid) {
		_CacheData, _error := model.GetCacheDataViaCid(iCid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheData)
	}
	herr.Message = "Can't get to the GetCacheDataViaCid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDataViaDataHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iData := self.Args("data").Bytes()
	if helper.IsHas(iData) {
		_CacheData, _error := model.GetCacheDataViaData(iData)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheData)
	}
	herr.Message = "Can't get to the GetCacheDataViaData's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDataViaExpireHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iExpire := self.Args("expire").MustInt()
	if helper.IsHas(iExpire) {
		_CacheData, _error := model.GetCacheDataViaExpire(iExpire)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheData)
	}
	herr.Message = "Can't get to the GetCacheDataViaExpire's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDataViaCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iCreated := self.Args("created").String()
	if helper.IsHas(iCreated) {
		_CacheData, _error := model.GetCacheDataViaCreated(iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheData)
	}
	herr.Message = "Can't get to the GetCacheDataViaCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDataViaSerializedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iSerialized := self.Args("serialized").MustInt()
	if helper.IsHas(iSerialized) {
		_CacheData, _error := model.GetCacheDataViaSerialized(iSerialized)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheData)
	}
	herr.Message = "Can't get to the GetCacheDataViaSerialized's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDataViaTagsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iTags := self.Args("tags").String()
	if helper.IsHas(iTags) {
		_CacheData, _error := model.GetCacheDataViaTags(iTags)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheData)
	}
	herr.Message = "Can't get to the GetCacheDataViaTags's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDataViaChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iChecksum := self.Args("checksum").String()
	if helper.IsHas(iChecksum) {
		_CacheData, _error := model.GetCacheDataViaChecksum(iChecksum)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheData)
	}
	herr.Message = "Can't get to the GetCacheDataViaChecksum's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetCacheDataViaCidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Cid_ := self.Args("cid").String()
	if helper.IsHas(Cid_) {
		var iCacheData model.CacheData
		self.Bind(&iCacheData)
		_CacheData, _error := model.SetCacheDataViaCid(Cid_, &iCacheData)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheData)
	}
	herr.Message = "Can't get to the SetCacheDataViaCid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetCacheDataViaDataHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Data_ := self.Args("data").Bytes()
	if helper.IsHas(Data_) {
		var iCacheData model.CacheData
		self.Bind(&iCacheData)
		_CacheData, _error := model.SetCacheDataViaData(Data_, &iCacheData)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheData)
	}
	herr.Message = "Can't get to the SetCacheDataViaData's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetCacheDataViaExpireHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Expire_ := self.Args("expire").MustInt()
	if helper.IsHas(Expire_) {
		var iCacheData model.CacheData
		self.Bind(&iCacheData)
		_CacheData, _error := model.SetCacheDataViaExpire(Expire_, &iCacheData)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheData)
	}
	herr.Message = "Can't get to the SetCacheDataViaExpire's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetCacheDataViaCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Created_ := self.Args("created").String()
	if helper.IsHas(Created_) {
		var iCacheData model.CacheData
		self.Bind(&iCacheData)
		_CacheData, _error := model.SetCacheDataViaCreated(Created_, &iCacheData)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheData)
	}
	herr.Message = "Can't get to the SetCacheDataViaCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetCacheDataViaSerializedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Serialized_ := self.Args("serialized").MustInt()
	if helper.IsHas(Serialized_) {
		var iCacheData model.CacheData
		self.Bind(&iCacheData)
		_CacheData, _error := model.SetCacheDataViaSerialized(Serialized_, &iCacheData)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheData)
	}
	herr.Message = "Can't get to the SetCacheDataViaSerialized's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetCacheDataViaTagsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Tags_ := self.Args("tags").String()
	if helper.IsHas(Tags_) {
		var iCacheData model.CacheData
		self.Bind(&iCacheData)
		_CacheData, _error := model.SetCacheDataViaTags(Tags_, &iCacheData)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheData)
	}
	herr.Message = "Can't get to the SetCacheDataViaTags's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetCacheDataViaChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Checksum_ := self.Args("checksum").String()
	if helper.IsHas(Checksum_) {
		var iCacheData model.CacheData
		self.Bind(&iCacheData)
		_CacheData, _error := model.SetCacheDataViaChecksum(Checksum_, &iCacheData)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheData)
	}
	herr.Message = "Can't get to the SetCacheDataViaChecksum's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostAddCacheDataHandler(self *macross.Context) error {
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
		_error := model.AddCacheData(Cid_,Data_,Expire_,Created_,Serialized_,Tags_,Checksum_)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		herr.Message = "StatusOK"
		herr.Status = macross.StatusOK
		return self.JSON(herr)
	}
	herr.Message = "Can't get to the AddCacheData's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostCacheDataHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iCacheData model.CacheData
	self.Bind(&iCacheData)
	_string, _error := model.PostCacheData(&iCacheData)
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

func PutCacheDataHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iCacheData model.CacheData
	self.Bind(&iCacheData)
	_string, _error := model.PutCacheData(&iCacheData)
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

func PutCacheDataViaCidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Cid_ := self.Args("cid").String()
	var iCacheData model.CacheData
	self.Bind(&iCacheData)
	_int64, _error := model.PutCacheDataViaCid(Cid_, &iCacheData)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutCacheDataViaDataHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Data_ := self.Args("data").Bytes()
	var iCacheData model.CacheData
	self.Bind(&iCacheData)
	_int64, _error := model.PutCacheDataViaData(Data_, &iCacheData)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutCacheDataViaExpireHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Expire_ := self.Args("expire").MustInt()
	var iCacheData model.CacheData
	self.Bind(&iCacheData)
	_int64, _error := model.PutCacheDataViaExpire(Expire_, &iCacheData)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutCacheDataViaCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Created_ := self.Args("created").String()
	var iCacheData model.CacheData
	self.Bind(&iCacheData)
	_int64, _error := model.PutCacheDataViaCreated(Created_, &iCacheData)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutCacheDataViaSerializedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Serialized_ := self.Args("serialized").MustInt()
	var iCacheData model.CacheData
	self.Bind(&iCacheData)
	_int64, _error := model.PutCacheDataViaSerialized(Serialized_, &iCacheData)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutCacheDataViaTagsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Tags_ := self.Args("tags").String()
	var iCacheData model.CacheData
	self.Bind(&iCacheData)
	_int64, _error := model.PutCacheDataViaTags(Tags_, &iCacheData)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutCacheDataViaChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Checksum_ := self.Args("checksum").String()
	var iCacheData model.CacheData
	self.Bind(&iCacheData)
	_int64, _error := model.PutCacheDataViaChecksum(Checksum_, &iCacheData)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateCacheDataViaCidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Cid_ := self.Args("cid").String()
	var iCacheData model.CacheData
	self.Bind(&iCacheData)
	var iMap = helper.StructToMap(iCacheData)
	_error := model.UpdateCacheDataViaCid(Cid_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateCacheDataViaDataHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Data_ := self.Args("data").Bytes()
	var iCacheData model.CacheData
	self.Bind(&iCacheData)
	var iMap = helper.StructToMap(iCacheData)
	_error := model.UpdateCacheDataViaData(Data_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateCacheDataViaExpireHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Expire_ := self.Args("expire").MustInt()
	var iCacheData model.CacheData
	self.Bind(&iCacheData)
	var iMap = helper.StructToMap(iCacheData)
	_error := model.UpdateCacheDataViaExpire(Expire_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateCacheDataViaCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Created_ := self.Args("created").String()
	var iCacheData model.CacheData
	self.Bind(&iCacheData)
	var iMap = helper.StructToMap(iCacheData)
	_error := model.UpdateCacheDataViaCreated(Created_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateCacheDataViaSerializedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Serialized_ := self.Args("serialized").MustInt()
	var iCacheData model.CacheData
	self.Bind(&iCacheData)
	var iMap = helper.StructToMap(iCacheData)
	_error := model.UpdateCacheDataViaSerialized(Serialized_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateCacheDataViaTagsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Tags_ := self.Args("tags").String()
	var iCacheData model.CacheData
	self.Bind(&iCacheData)
	var iMap = helper.StructToMap(iCacheData)
	_error := model.UpdateCacheDataViaTags(Tags_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateCacheDataViaChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Checksum_ := self.Args("checksum").String()
	var iCacheData model.CacheData
	self.Bind(&iCacheData)
	var iMap = helper.StructToMap(iCacheData)
	_error := model.UpdateCacheDataViaChecksum(Checksum_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteCacheDataHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Cid_ := self.Args("cid").String()
	_int64, _error := model.DeleteCacheData(Cid_)
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

func DeleteCacheDataViaCidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Cid_ := self.Args("cid").String()
	_error := model.DeleteCacheDataViaCid(Cid_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteCacheDataViaDataHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Data_ := self.Args("data").Bytes()
	_error := model.DeleteCacheDataViaData(Data_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteCacheDataViaExpireHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Expire_ := self.Args("expire").MustInt()
	_error := model.DeleteCacheDataViaExpire(Expire_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteCacheDataViaCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Created_ := self.Args("created").String()
	_error := model.DeleteCacheDataViaCreated(Created_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteCacheDataViaSerializedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Serialized_ := self.Args("serialized").MustInt()
	_error := model.DeleteCacheDataViaSerialized(Serialized_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteCacheDataViaTagsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Tags_ := self.Args("tags").String()
	_error := model.DeleteCacheDataViaTags(Tags_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteCacheDataViaChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Checksum_ := self.Args("checksum").String()
	_error := model.DeleteCacheDataViaChecksum(Checksum_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}
