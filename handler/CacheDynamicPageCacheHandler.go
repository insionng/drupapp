package handler

import (
	"github.com/insionng/drupapp/helper"
	"github.com/insionng/drupapp/model"
	"github.com/insionng/macross"
)

func GetCacheDynamicPageCachesCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()

	if offset > 0 {
		var m = map[string]interface{}{}
		_int64, _error := model.GetCacheDynamicPageCachesCount(offset, limit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		m["cache_dynamic_page_cachesCount"] = _int64
		return self.JSON(m)
	}
	herr.Message = "Can't get to the GetCacheDynamicPageCachesCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDynamicPageCacheCountViaCidHandler(self *macross.Context) error {
	Cid_ := self.Args("cid").String()
	_int64 := model.GetCacheDynamicPageCacheCountViaCid(Cid_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["cache_dynamic_page_cacheCount"] = 0
	}
	m["cache_dynamic_page_cacheCount"] = _int64
	return self.JSON(m)
}

func GetCacheDynamicPageCacheCountViaDataHandler(self *macross.Context) error {
	Data_ := self.Args("data").Bytes()
	_int64 := model.GetCacheDynamicPageCacheCountViaData(Data_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["cache_dynamic_page_cacheCount"] = 0
	}
	m["cache_dynamic_page_cacheCount"] = _int64
	return self.JSON(m)
}

func GetCacheDynamicPageCacheCountViaExpireHandler(self *macross.Context) error {
	Expire_ := self.Args("expire").MustInt()
	_int64 := model.GetCacheDynamicPageCacheCountViaExpire(Expire_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["cache_dynamic_page_cacheCount"] = 0
	}
	m["cache_dynamic_page_cacheCount"] = _int64
	return self.JSON(m)
}

func GetCacheDynamicPageCacheCountViaCreatedHandler(self *macross.Context) error {
	Created_ := self.Args("created").String()
	_int64 := model.GetCacheDynamicPageCacheCountViaCreated(Created_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["cache_dynamic_page_cacheCount"] = 0
	}
	m["cache_dynamic_page_cacheCount"] = _int64
	return self.JSON(m)
}

func GetCacheDynamicPageCacheCountViaSerializedHandler(self *macross.Context) error {
	Serialized_ := self.Args("serialized").MustInt()
	_int64 := model.GetCacheDynamicPageCacheCountViaSerialized(Serialized_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["cache_dynamic_page_cacheCount"] = 0
	}
	m["cache_dynamic_page_cacheCount"] = _int64
	return self.JSON(m)
}

func GetCacheDynamicPageCacheCountViaTagsHandler(self *macross.Context) error {
	Tags_ := self.Args("tags").String()
	_int64 := model.GetCacheDynamicPageCacheCountViaTags(Tags_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["cache_dynamic_page_cacheCount"] = 0
	}
	m["cache_dynamic_page_cacheCount"] = _int64
	return self.JSON(m)
}

func GetCacheDynamicPageCacheCountViaChecksumHandler(self *macross.Context) error {
	Checksum_ := self.Args("checksum").String()
	_int64 := model.GetCacheDynamicPageCacheCountViaChecksum(Checksum_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["cache_dynamic_page_cacheCount"] = 0
	}
	m["cache_dynamic_page_cacheCount"] = _int64
	return self.JSON(m)
}

func GetCacheDynamicPageCachesViaCidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iCid := self.Args("cid").String()
	if (offset > 0) && helper.IsHas(iCid) {
		_CacheDynamicPageCache, _error := model.GetCacheDynamicPageCachesViaCid(offset, limit, iCid, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDynamicPageCache)
	}
	herr.Message = "Can't get to the GetCacheDynamicPageCachesViaCid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDynamicPageCachesViaDataHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iData := self.Args("data").Bytes()
	if (offset > 0) && helper.IsHas(iData) {
		_CacheDynamicPageCache, _error := model.GetCacheDynamicPageCachesViaData(offset, limit, iData, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDynamicPageCache)
	}
	herr.Message = "Can't get to the GetCacheDynamicPageCachesViaData's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDynamicPageCachesViaExpireHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iExpire := self.Args("expire").MustInt()
	if (offset > 0) && helper.IsHas(iExpire) {
		_CacheDynamicPageCache, _error := model.GetCacheDynamicPageCachesViaExpire(offset, limit, iExpire, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDynamicPageCache)
	}
	herr.Message = "Can't get to the GetCacheDynamicPageCachesViaExpire's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDynamicPageCachesViaCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iCreated := self.Args("created").String()
	if (offset > 0) && helper.IsHas(iCreated) {
		_CacheDynamicPageCache, _error := model.GetCacheDynamicPageCachesViaCreated(offset, limit, iCreated, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDynamicPageCache)
	}
	herr.Message = "Can't get to the GetCacheDynamicPageCachesViaCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDynamicPageCachesViaSerializedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iSerialized := self.Args("serialized").MustInt()
	if (offset > 0) && helper.IsHas(iSerialized) {
		_CacheDynamicPageCache, _error := model.GetCacheDynamicPageCachesViaSerialized(offset, limit, iSerialized, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDynamicPageCache)
	}
	herr.Message = "Can't get to the GetCacheDynamicPageCachesViaSerialized's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDynamicPageCachesViaTagsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iTags := self.Args("tags").String()
	if (offset > 0) && helper.IsHas(iTags) {
		_CacheDynamicPageCache, _error := model.GetCacheDynamicPageCachesViaTags(offset, limit, iTags, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDynamicPageCache)
	}
	herr.Message = "Can't get to the GetCacheDynamicPageCachesViaTags's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDynamicPageCachesViaChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iChecksum := self.Args("checksum").String()
	if (offset > 0) && helper.IsHas(iChecksum) {
		_CacheDynamicPageCache, _error := model.GetCacheDynamicPageCachesViaChecksum(offset, limit, iChecksum, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDynamicPageCache)
	}
	herr.Message = "Can't get to the GetCacheDynamicPageCachesViaChecksum's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDynamicPageCachesByCidAndDataAndExpireHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").String()
	iData := self.Args("data").Bytes()
	iExpire := self.Args("expire").MustInt()

	if helper.IsHas(iCid) {
		_CacheDynamicPageCache, _error := model.GetCacheDynamicPageCachesByCidAndDataAndExpire(offset, limit, iCid,iData,iExpire)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDynamicPageCache)
	}
	herr.Message = "Can't get to the GetCacheDynamicPageCachesByCidAndDataAndExpire's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDynamicPageCachesByCidAndDataAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").String()
	iData := self.Args("data").Bytes()
	iCreated := self.Args("created").String()

	if helper.IsHas(iCid) {
		_CacheDynamicPageCache, _error := model.GetCacheDynamicPageCachesByCidAndDataAndCreated(offset, limit, iCid,iData,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDynamicPageCache)
	}
	herr.Message = "Can't get to the GetCacheDynamicPageCachesByCidAndDataAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDynamicPageCachesByCidAndDataAndSerializedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").String()
	iData := self.Args("data").Bytes()
	iSerialized := self.Args("serialized").MustInt()

	if helper.IsHas(iCid) {
		_CacheDynamicPageCache, _error := model.GetCacheDynamicPageCachesByCidAndDataAndSerialized(offset, limit, iCid,iData,iSerialized)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDynamicPageCache)
	}
	herr.Message = "Can't get to the GetCacheDynamicPageCachesByCidAndDataAndSerialized's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDynamicPageCachesByCidAndDataAndTagsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").String()
	iData := self.Args("data").Bytes()
	iTags := self.Args("tags").String()

	if helper.IsHas(iCid) {
		_CacheDynamicPageCache, _error := model.GetCacheDynamicPageCachesByCidAndDataAndTags(offset, limit, iCid,iData,iTags)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDynamicPageCache)
	}
	herr.Message = "Can't get to the GetCacheDynamicPageCachesByCidAndDataAndTags's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDynamicPageCachesByCidAndDataAndChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").String()
	iData := self.Args("data").Bytes()
	iChecksum := self.Args("checksum").String()

	if helper.IsHas(iCid) {
		_CacheDynamicPageCache, _error := model.GetCacheDynamicPageCachesByCidAndDataAndChecksum(offset, limit, iCid,iData,iChecksum)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDynamicPageCache)
	}
	herr.Message = "Can't get to the GetCacheDynamicPageCachesByCidAndDataAndChecksum's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDynamicPageCachesByCidAndExpireAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").String()
	iExpire := self.Args("expire").MustInt()
	iCreated := self.Args("created").String()

	if helper.IsHas(iCid) {
		_CacheDynamicPageCache, _error := model.GetCacheDynamicPageCachesByCidAndExpireAndCreated(offset, limit, iCid,iExpire,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDynamicPageCache)
	}
	herr.Message = "Can't get to the GetCacheDynamicPageCachesByCidAndExpireAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDynamicPageCachesByCidAndExpireAndSerializedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").String()
	iExpire := self.Args("expire").MustInt()
	iSerialized := self.Args("serialized").MustInt()

	if helper.IsHas(iCid) {
		_CacheDynamicPageCache, _error := model.GetCacheDynamicPageCachesByCidAndExpireAndSerialized(offset, limit, iCid,iExpire,iSerialized)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDynamicPageCache)
	}
	herr.Message = "Can't get to the GetCacheDynamicPageCachesByCidAndExpireAndSerialized's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDynamicPageCachesByCidAndExpireAndTagsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").String()
	iExpire := self.Args("expire").MustInt()
	iTags := self.Args("tags").String()

	if helper.IsHas(iCid) {
		_CacheDynamicPageCache, _error := model.GetCacheDynamicPageCachesByCidAndExpireAndTags(offset, limit, iCid,iExpire,iTags)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDynamicPageCache)
	}
	herr.Message = "Can't get to the GetCacheDynamicPageCachesByCidAndExpireAndTags's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDynamicPageCachesByCidAndExpireAndChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").String()
	iExpire := self.Args("expire").MustInt()
	iChecksum := self.Args("checksum").String()

	if helper.IsHas(iCid) {
		_CacheDynamicPageCache, _error := model.GetCacheDynamicPageCachesByCidAndExpireAndChecksum(offset, limit, iCid,iExpire,iChecksum)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDynamicPageCache)
	}
	herr.Message = "Can't get to the GetCacheDynamicPageCachesByCidAndExpireAndChecksum's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDynamicPageCachesByCidAndCreatedAndSerializedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").String()
	iCreated := self.Args("created").String()
	iSerialized := self.Args("serialized").MustInt()

	if helper.IsHas(iCid) {
		_CacheDynamicPageCache, _error := model.GetCacheDynamicPageCachesByCidAndCreatedAndSerialized(offset, limit, iCid,iCreated,iSerialized)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDynamicPageCache)
	}
	herr.Message = "Can't get to the GetCacheDynamicPageCachesByCidAndCreatedAndSerialized's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDynamicPageCachesByCidAndCreatedAndTagsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").String()
	iCreated := self.Args("created").String()
	iTags := self.Args("tags").String()

	if helper.IsHas(iCid) {
		_CacheDynamicPageCache, _error := model.GetCacheDynamicPageCachesByCidAndCreatedAndTags(offset, limit, iCid,iCreated,iTags)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDynamicPageCache)
	}
	herr.Message = "Can't get to the GetCacheDynamicPageCachesByCidAndCreatedAndTags's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDynamicPageCachesByCidAndCreatedAndChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").String()
	iCreated := self.Args("created").String()
	iChecksum := self.Args("checksum").String()

	if helper.IsHas(iCid) {
		_CacheDynamicPageCache, _error := model.GetCacheDynamicPageCachesByCidAndCreatedAndChecksum(offset, limit, iCid,iCreated,iChecksum)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDynamicPageCache)
	}
	herr.Message = "Can't get to the GetCacheDynamicPageCachesByCidAndCreatedAndChecksum's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDynamicPageCachesByCidAndSerializedAndTagsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").String()
	iSerialized := self.Args("serialized").MustInt()
	iTags := self.Args("tags").String()

	if helper.IsHas(iCid) {
		_CacheDynamicPageCache, _error := model.GetCacheDynamicPageCachesByCidAndSerializedAndTags(offset, limit, iCid,iSerialized,iTags)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDynamicPageCache)
	}
	herr.Message = "Can't get to the GetCacheDynamicPageCachesByCidAndSerializedAndTags's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDynamicPageCachesByCidAndSerializedAndChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").String()
	iSerialized := self.Args("serialized").MustInt()
	iChecksum := self.Args("checksum").String()

	if helper.IsHas(iCid) {
		_CacheDynamicPageCache, _error := model.GetCacheDynamicPageCachesByCidAndSerializedAndChecksum(offset, limit, iCid,iSerialized,iChecksum)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDynamicPageCache)
	}
	herr.Message = "Can't get to the GetCacheDynamicPageCachesByCidAndSerializedAndChecksum's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDynamicPageCachesByCidAndTagsAndChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").String()
	iTags := self.Args("tags").String()
	iChecksum := self.Args("checksum").String()

	if helper.IsHas(iCid) {
		_CacheDynamicPageCache, _error := model.GetCacheDynamicPageCachesByCidAndTagsAndChecksum(offset, limit, iCid,iTags,iChecksum)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDynamicPageCache)
	}
	herr.Message = "Can't get to the GetCacheDynamicPageCachesByCidAndTagsAndChecksum's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDynamicPageCachesByDataAndExpireAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iData := self.Args("data").Bytes()
	iExpire := self.Args("expire").MustInt()
	iCreated := self.Args("created").String()

	if helper.IsHas(iData) {
		_CacheDynamicPageCache, _error := model.GetCacheDynamicPageCachesByDataAndExpireAndCreated(offset, limit, iData,iExpire,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDynamicPageCache)
	}
	herr.Message = "Can't get to the GetCacheDynamicPageCachesByDataAndExpireAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDynamicPageCachesByDataAndExpireAndSerializedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iData := self.Args("data").Bytes()
	iExpire := self.Args("expire").MustInt()
	iSerialized := self.Args("serialized").MustInt()

	if helper.IsHas(iData) {
		_CacheDynamicPageCache, _error := model.GetCacheDynamicPageCachesByDataAndExpireAndSerialized(offset, limit, iData,iExpire,iSerialized)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDynamicPageCache)
	}
	herr.Message = "Can't get to the GetCacheDynamicPageCachesByDataAndExpireAndSerialized's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDynamicPageCachesByDataAndExpireAndTagsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iData := self.Args("data").Bytes()
	iExpire := self.Args("expire").MustInt()
	iTags := self.Args("tags").String()

	if helper.IsHas(iData) {
		_CacheDynamicPageCache, _error := model.GetCacheDynamicPageCachesByDataAndExpireAndTags(offset, limit, iData,iExpire,iTags)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDynamicPageCache)
	}
	herr.Message = "Can't get to the GetCacheDynamicPageCachesByDataAndExpireAndTags's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDynamicPageCachesByDataAndExpireAndChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iData := self.Args("data").Bytes()
	iExpire := self.Args("expire").MustInt()
	iChecksum := self.Args("checksum").String()

	if helper.IsHas(iData) {
		_CacheDynamicPageCache, _error := model.GetCacheDynamicPageCachesByDataAndExpireAndChecksum(offset, limit, iData,iExpire,iChecksum)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDynamicPageCache)
	}
	herr.Message = "Can't get to the GetCacheDynamicPageCachesByDataAndExpireAndChecksum's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDynamicPageCachesByDataAndCreatedAndSerializedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iData := self.Args("data").Bytes()
	iCreated := self.Args("created").String()
	iSerialized := self.Args("serialized").MustInt()

	if helper.IsHas(iData) {
		_CacheDynamicPageCache, _error := model.GetCacheDynamicPageCachesByDataAndCreatedAndSerialized(offset, limit, iData,iCreated,iSerialized)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDynamicPageCache)
	}
	herr.Message = "Can't get to the GetCacheDynamicPageCachesByDataAndCreatedAndSerialized's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDynamicPageCachesByDataAndCreatedAndTagsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iData := self.Args("data").Bytes()
	iCreated := self.Args("created").String()
	iTags := self.Args("tags").String()

	if helper.IsHas(iData) {
		_CacheDynamicPageCache, _error := model.GetCacheDynamicPageCachesByDataAndCreatedAndTags(offset, limit, iData,iCreated,iTags)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDynamicPageCache)
	}
	herr.Message = "Can't get to the GetCacheDynamicPageCachesByDataAndCreatedAndTags's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDynamicPageCachesByDataAndCreatedAndChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iData := self.Args("data").Bytes()
	iCreated := self.Args("created").String()
	iChecksum := self.Args("checksum").String()

	if helper.IsHas(iData) {
		_CacheDynamicPageCache, _error := model.GetCacheDynamicPageCachesByDataAndCreatedAndChecksum(offset, limit, iData,iCreated,iChecksum)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDynamicPageCache)
	}
	herr.Message = "Can't get to the GetCacheDynamicPageCachesByDataAndCreatedAndChecksum's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDynamicPageCachesByDataAndSerializedAndTagsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iData := self.Args("data").Bytes()
	iSerialized := self.Args("serialized").MustInt()
	iTags := self.Args("tags").String()

	if helper.IsHas(iData) {
		_CacheDynamicPageCache, _error := model.GetCacheDynamicPageCachesByDataAndSerializedAndTags(offset, limit, iData,iSerialized,iTags)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDynamicPageCache)
	}
	herr.Message = "Can't get to the GetCacheDynamicPageCachesByDataAndSerializedAndTags's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDynamicPageCachesByDataAndSerializedAndChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iData := self.Args("data").Bytes()
	iSerialized := self.Args("serialized").MustInt()
	iChecksum := self.Args("checksum").String()

	if helper.IsHas(iData) {
		_CacheDynamicPageCache, _error := model.GetCacheDynamicPageCachesByDataAndSerializedAndChecksum(offset, limit, iData,iSerialized,iChecksum)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDynamicPageCache)
	}
	herr.Message = "Can't get to the GetCacheDynamicPageCachesByDataAndSerializedAndChecksum's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDynamicPageCachesByDataAndTagsAndChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iData := self.Args("data").Bytes()
	iTags := self.Args("tags").String()
	iChecksum := self.Args("checksum").String()

	if helper.IsHas(iData) {
		_CacheDynamicPageCache, _error := model.GetCacheDynamicPageCachesByDataAndTagsAndChecksum(offset, limit, iData,iTags,iChecksum)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDynamicPageCache)
	}
	herr.Message = "Can't get to the GetCacheDynamicPageCachesByDataAndTagsAndChecksum's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDynamicPageCachesByExpireAndCreatedAndSerializedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iExpire := self.Args("expire").MustInt()
	iCreated := self.Args("created").String()
	iSerialized := self.Args("serialized").MustInt()

	if helper.IsHas(iExpire) {
		_CacheDynamicPageCache, _error := model.GetCacheDynamicPageCachesByExpireAndCreatedAndSerialized(offset, limit, iExpire,iCreated,iSerialized)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDynamicPageCache)
	}
	herr.Message = "Can't get to the GetCacheDynamicPageCachesByExpireAndCreatedAndSerialized's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDynamicPageCachesByExpireAndCreatedAndTagsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iExpire := self.Args("expire").MustInt()
	iCreated := self.Args("created").String()
	iTags := self.Args("tags").String()

	if helper.IsHas(iExpire) {
		_CacheDynamicPageCache, _error := model.GetCacheDynamicPageCachesByExpireAndCreatedAndTags(offset, limit, iExpire,iCreated,iTags)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDynamicPageCache)
	}
	herr.Message = "Can't get to the GetCacheDynamicPageCachesByExpireAndCreatedAndTags's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDynamicPageCachesByExpireAndCreatedAndChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iExpire := self.Args("expire").MustInt()
	iCreated := self.Args("created").String()
	iChecksum := self.Args("checksum").String()

	if helper.IsHas(iExpire) {
		_CacheDynamicPageCache, _error := model.GetCacheDynamicPageCachesByExpireAndCreatedAndChecksum(offset, limit, iExpire,iCreated,iChecksum)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDynamicPageCache)
	}
	herr.Message = "Can't get to the GetCacheDynamicPageCachesByExpireAndCreatedAndChecksum's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDynamicPageCachesByExpireAndSerializedAndTagsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iExpire := self.Args("expire").MustInt()
	iSerialized := self.Args("serialized").MustInt()
	iTags := self.Args("tags").String()

	if helper.IsHas(iExpire) {
		_CacheDynamicPageCache, _error := model.GetCacheDynamicPageCachesByExpireAndSerializedAndTags(offset, limit, iExpire,iSerialized,iTags)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDynamicPageCache)
	}
	herr.Message = "Can't get to the GetCacheDynamicPageCachesByExpireAndSerializedAndTags's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDynamicPageCachesByExpireAndSerializedAndChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iExpire := self.Args("expire").MustInt()
	iSerialized := self.Args("serialized").MustInt()
	iChecksum := self.Args("checksum").String()

	if helper.IsHas(iExpire) {
		_CacheDynamicPageCache, _error := model.GetCacheDynamicPageCachesByExpireAndSerializedAndChecksum(offset, limit, iExpire,iSerialized,iChecksum)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDynamicPageCache)
	}
	herr.Message = "Can't get to the GetCacheDynamicPageCachesByExpireAndSerializedAndChecksum's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDynamicPageCachesByExpireAndTagsAndChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iExpire := self.Args("expire").MustInt()
	iTags := self.Args("tags").String()
	iChecksum := self.Args("checksum").String()

	if helper.IsHas(iExpire) {
		_CacheDynamicPageCache, _error := model.GetCacheDynamicPageCachesByExpireAndTagsAndChecksum(offset, limit, iExpire,iTags,iChecksum)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDynamicPageCache)
	}
	herr.Message = "Can't get to the GetCacheDynamicPageCachesByExpireAndTagsAndChecksum's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDynamicPageCachesByCreatedAndSerializedAndTagsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCreated := self.Args("created").String()
	iSerialized := self.Args("serialized").MustInt()
	iTags := self.Args("tags").String()

	if helper.IsHas(iCreated) {
		_CacheDynamicPageCache, _error := model.GetCacheDynamicPageCachesByCreatedAndSerializedAndTags(offset, limit, iCreated,iSerialized,iTags)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDynamicPageCache)
	}
	herr.Message = "Can't get to the GetCacheDynamicPageCachesByCreatedAndSerializedAndTags's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDynamicPageCachesByCreatedAndSerializedAndChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCreated := self.Args("created").String()
	iSerialized := self.Args("serialized").MustInt()
	iChecksum := self.Args("checksum").String()

	if helper.IsHas(iCreated) {
		_CacheDynamicPageCache, _error := model.GetCacheDynamicPageCachesByCreatedAndSerializedAndChecksum(offset, limit, iCreated,iSerialized,iChecksum)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDynamicPageCache)
	}
	herr.Message = "Can't get to the GetCacheDynamicPageCachesByCreatedAndSerializedAndChecksum's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDynamicPageCachesByCreatedAndTagsAndChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCreated := self.Args("created").String()
	iTags := self.Args("tags").String()
	iChecksum := self.Args("checksum").String()

	if helper.IsHas(iCreated) {
		_CacheDynamicPageCache, _error := model.GetCacheDynamicPageCachesByCreatedAndTagsAndChecksum(offset, limit, iCreated,iTags,iChecksum)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDynamicPageCache)
	}
	herr.Message = "Can't get to the GetCacheDynamicPageCachesByCreatedAndTagsAndChecksum's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDynamicPageCachesBySerializedAndTagsAndChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSerialized := self.Args("serialized").MustInt()
	iTags := self.Args("tags").String()
	iChecksum := self.Args("checksum").String()

	if helper.IsHas(iSerialized) {
		_CacheDynamicPageCache, _error := model.GetCacheDynamicPageCachesBySerializedAndTagsAndChecksum(offset, limit, iSerialized,iTags,iChecksum)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDynamicPageCache)
	}
	herr.Message = "Can't get to the GetCacheDynamicPageCachesBySerializedAndTagsAndChecksum's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDynamicPageCachesByCidAndDataHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").String()
	iData := self.Args("data").Bytes()

	if helper.IsHas(iCid) {
		_CacheDynamicPageCache, _error := model.GetCacheDynamicPageCachesByCidAndData(offset, limit, iCid,iData)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDynamicPageCache)
	}
	herr.Message = "Can't get to the GetCacheDynamicPageCachesByCidAndData's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDynamicPageCachesByCidAndExpireHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").String()
	iExpire := self.Args("expire").MustInt()

	if helper.IsHas(iCid) {
		_CacheDynamicPageCache, _error := model.GetCacheDynamicPageCachesByCidAndExpire(offset, limit, iCid,iExpire)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDynamicPageCache)
	}
	herr.Message = "Can't get to the GetCacheDynamicPageCachesByCidAndExpire's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDynamicPageCachesByCidAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").String()
	iCreated := self.Args("created").String()

	if helper.IsHas(iCid) {
		_CacheDynamicPageCache, _error := model.GetCacheDynamicPageCachesByCidAndCreated(offset, limit, iCid,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDynamicPageCache)
	}
	herr.Message = "Can't get to the GetCacheDynamicPageCachesByCidAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDynamicPageCachesByCidAndSerializedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").String()
	iSerialized := self.Args("serialized").MustInt()

	if helper.IsHas(iCid) {
		_CacheDynamicPageCache, _error := model.GetCacheDynamicPageCachesByCidAndSerialized(offset, limit, iCid,iSerialized)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDynamicPageCache)
	}
	herr.Message = "Can't get to the GetCacheDynamicPageCachesByCidAndSerialized's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDynamicPageCachesByCidAndTagsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").String()
	iTags := self.Args("tags").String()

	if helper.IsHas(iCid) {
		_CacheDynamicPageCache, _error := model.GetCacheDynamicPageCachesByCidAndTags(offset, limit, iCid,iTags)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDynamicPageCache)
	}
	herr.Message = "Can't get to the GetCacheDynamicPageCachesByCidAndTags's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDynamicPageCachesByCidAndChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").String()
	iChecksum := self.Args("checksum").String()

	if helper.IsHas(iCid) {
		_CacheDynamicPageCache, _error := model.GetCacheDynamicPageCachesByCidAndChecksum(offset, limit, iCid,iChecksum)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDynamicPageCache)
	}
	herr.Message = "Can't get to the GetCacheDynamicPageCachesByCidAndChecksum's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDynamicPageCachesByDataAndExpireHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iData := self.Args("data").Bytes()
	iExpire := self.Args("expire").MustInt()

	if helper.IsHas(iData) {
		_CacheDynamicPageCache, _error := model.GetCacheDynamicPageCachesByDataAndExpire(offset, limit, iData,iExpire)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDynamicPageCache)
	}
	herr.Message = "Can't get to the GetCacheDynamicPageCachesByDataAndExpire's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDynamicPageCachesByDataAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iData := self.Args("data").Bytes()
	iCreated := self.Args("created").String()

	if helper.IsHas(iData) {
		_CacheDynamicPageCache, _error := model.GetCacheDynamicPageCachesByDataAndCreated(offset, limit, iData,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDynamicPageCache)
	}
	herr.Message = "Can't get to the GetCacheDynamicPageCachesByDataAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDynamicPageCachesByDataAndSerializedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iData := self.Args("data").Bytes()
	iSerialized := self.Args("serialized").MustInt()

	if helper.IsHas(iData) {
		_CacheDynamicPageCache, _error := model.GetCacheDynamicPageCachesByDataAndSerialized(offset, limit, iData,iSerialized)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDynamicPageCache)
	}
	herr.Message = "Can't get to the GetCacheDynamicPageCachesByDataAndSerialized's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDynamicPageCachesByDataAndTagsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iData := self.Args("data").Bytes()
	iTags := self.Args("tags").String()

	if helper.IsHas(iData) {
		_CacheDynamicPageCache, _error := model.GetCacheDynamicPageCachesByDataAndTags(offset, limit, iData,iTags)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDynamicPageCache)
	}
	herr.Message = "Can't get to the GetCacheDynamicPageCachesByDataAndTags's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDynamicPageCachesByDataAndChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iData := self.Args("data").Bytes()
	iChecksum := self.Args("checksum").String()

	if helper.IsHas(iData) {
		_CacheDynamicPageCache, _error := model.GetCacheDynamicPageCachesByDataAndChecksum(offset, limit, iData,iChecksum)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDynamicPageCache)
	}
	herr.Message = "Can't get to the GetCacheDynamicPageCachesByDataAndChecksum's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDynamicPageCachesByExpireAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iExpire := self.Args("expire").MustInt()
	iCreated := self.Args("created").String()

	if helper.IsHas(iExpire) {
		_CacheDynamicPageCache, _error := model.GetCacheDynamicPageCachesByExpireAndCreated(offset, limit, iExpire,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDynamicPageCache)
	}
	herr.Message = "Can't get to the GetCacheDynamicPageCachesByExpireAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDynamicPageCachesByExpireAndSerializedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iExpire := self.Args("expire").MustInt()
	iSerialized := self.Args("serialized").MustInt()

	if helper.IsHas(iExpire) {
		_CacheDynamicPageCache, _error := model.GetCacheDynamicPageCachesByExpireAndSerialized(offset, limit, iExpire,iSerialized)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDynamicPageCache)
	}
	herr.Message = "Can't get to the GetCacheDynamicPageCachesByExpireAndSerialized's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDynamicPageCachesByExpireAndTagsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iExpire := self.Args("expire").MustInt()
	iTags := self.Args("tags").String()

	if helper.IsHas(iExpire) {
		_CacheDynamicPageCache, _error := model.GetCacheDynamicPageCachesByExpireAndTags(offset, limit, iExpire,iTags)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDynamicPageCache)
	}
	herr.Message = "Can't get to the GetCacheDynamicPageCachesByExpireAndTags's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDynamicPageCachesByExpireAndChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iExpire := self.Args("expire").MustInt()
	iChecksum := self.Args("checksum").String()

	if helper.IsHas(iExpire) {
		_CacheDynamicPageCache, _error := model.GetCacheDynamicPageCachesByExpireAndChecksum(offset, limit, iExpire,iChecksum)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDynamicPageCache)
	}
	herr.Message = "Can't get to the GetCacheDynamicPageCachesByExpireAndChecksum's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDynamicPageCachesByCreatedAndSerializedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCreated := self.Args("created").String()
	iSerialized := self.Args("serialized").MustInt()

	if helper.IsHas(iCreated) {
		_CacheDynamicPageCache, _error := model.GetCacheDynamicPageCachesByCreatedAndSerialized(offset, limit, iCreated,iSerialized)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDynamicPageCache)
	}
	herr.Message = "Can't get to the GetCacheDynamicPageCachesByCreatedAndSerialized's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDynamicPageCachesByCreatedAndTagsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCreated := self.Args("created").String()
	iTags := self.Args("tags").String()

	if helper.IsHas(iCreated) {
		_CacheDynamicPageCache, _error := model.GetCacheDynamicPageCachesByCreatedAndTags(offset, limit, iCreated,iTags)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDynamicPageCache)
	}
	herr.Message = "Can't get to the GetCacheDynamicPageCachesByCreatedAndTags's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDynamicPageCachesByCreatedAndChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCreated := self.Args("created").String()
	iChecksum := self.Args("checksum").String()

	if helper.IsHas(iCreated) {
		_CacheDynamicPageCache, _error := model.GetCacheDynamicPageCachesByCreatedAndChecksum(offset, limit, iCreated,iChecksum)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDynamicPageCache)
	}
	herr.Message = "Can't get to the GetCacheDynamicPageCachesByCreatedAndChecksum's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDynamicPageCachesBySerializedAndTagsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSerialized := self.Args("serialized").MustInt()
	iTags := self.Args("tags").String()

	if helper.IsHas(iSerialized) {
		_CacheDynamicPageCache, _error := model.GetCacheDynamicPageCachesBySerializedAndTags(offset, limit, iSerialized,iTags)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDynamicPageCache)
	}
	herr.Message = "Can't get to the GetCacheDynamicPageCachesBySerializedAndTags's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDynamicPageCachesBySerializedAndChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSerialized := self.Args("serialized").MustInt()
	iChecksum := self.Args("checksum").String()

	if helper.IsHas(iSerialized) {
		_CacheDynamicPageCache, _error := model.GetCacheDynamicPageCachesBySerializedAndChecksum(offset, limit, iSerialized,iChecksum)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDynamicPageCache)
	}
	herr.Message = "Can't get to the GetCacheDynamicPageCachesBySerializedAndChecksum's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDynamicPageCachesByTagsAndChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTags := self.Args("tags").String()
	iChecksum := self.Args("checksum").String()

	if helper.IsHas(iTags) {
		_CacheDynamicPageCache, _error := model.GetCacheDynamicPageCachesByTagsAndChecksum(offset, limit, iTags,iChecksum)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDynamicPageCache)
	}
	herr.Message = "Can't get to the GetCacheDynamicPageCachesByTagsAndChecksum's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDynamicPageCachesHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	if (offset > 0) && (len(field) > 0) {
		_CacheDynamicPageCache, _error := model.GetCacheDynamicPageCaches(offset, limit, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDynamicPageCache)
	}
	herr.Message = "Can't get to the GetCacheDynamicPageCaches' args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasCacheDynamicPageCacheViaCidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iCid := self.Args("cid").String()
	if helper.IsHas(iCid) {
		_CacheDynamicPageCache := model.HasCacheDynamicPageCacheViaCid(iCid)
		var m = map[string]interface{}{}
		m["cache_dynamic_page_cache"] = _CacheDynamicPageCache
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasCacheDynamicPageCacheViaCid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasCacheDynamicPageCacheViaDataHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iData := self.Args("data").Bytes()
	if helper.IsHas(iData) {
		_CacheDynamicPageCache := model.HasCacheDynamicPageCacheViaData(iData)
		var m = map[string]interface{}{}
		m["cache_dynamic_page_cache"] = _CacheDynamicPageCache
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasCacheDynamicPageCacheViaData's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasCacheDynamicPageCacheViaExpireHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iExpire := self.Args("expire").MustInt()
	if helper.IsHas(iExpire) {
		_CacheDynamicPageCache := model.HasCacheDynamicPageCacheViaExpire(iExpire)
		var m = map[string]interface{}{}
		m["cache_dynamic_page_cache"] = _CacheDynamicPageCache
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasCacheDynamicPageCacheViaExpire's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasCacheDynamicPageCacheViaCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iCreated := self.Args("created").String()
	if helper.IsHas(iCreated) {
		_CacheDynamicPageCache := model.HasCacheDynamicPageCacheViaCreated(iCreated)
		var m = map[string]interface{}{}
		m["cache_dynamic_page_cache"] = _CacheDynamicPageCache
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasCacheDynamicPageCacheViaCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasCacheDynamicPageCacheViaSerializedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iSerialized := self.Args("serialized").MustInt()
	if helper.IsHas(iSerialized) {
		_CacheDynamicPageCache := model.HasCacheDynamicPageCacheViaSerialized(iSerialized)
		var m = map[string]interface{}{}
		m["cache_dynamic_page_cache"] = _CacheDynamicPageCache
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasCacheDynamicPageCacheViaSerialized's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasCacheDynamicPageCacheViaTagsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iTags := self.Args("tags").String()
	if helper.IsHas(iTags) {
		_CacheDynamicPageCache := model.HasCacheDynamicPageCacheViaTags(iTags)
		var m = map[string]interface{}{}
		m["cache_dynamic_page_cache"] = _CacheDynamicPageCache
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasCacheDynamicPageCacheViaTags's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasCacheDynamicPageCacheViaChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iChecksum := self.Args("checksum").String()
	if helper.IsHas(iChecksum) {
		_CacheDynamicPageCache := model.HasCacheDynamicPageCacheViaChecksum(iChecksum)
		var m = map[string]interface{}{}
		m["cache_dynamic_page_cache"] = _CacheDynamicPageCache
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasCacheDynamicPageCacheViaChecksum's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDynamicPageCacheViaCidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iCid := self.Args("cid").String()
	if helper.IsHas(iCid) {
		_CacheDynamicPageCache, _error := model.GetCacheDynamicPageCacheViaCid(iCid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDynamicPageCache)
	}
	herr.Message = "Can't get to the GetCacheDynamicPageCacheViaCid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDynamicPageCacheViaDataHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iData := self.Args("data").Bytes()
	if helper.IsHas(iData) {
		_CacheDynamicPageCache, _error := model.GetCacheDynamicPageCacheViaData(iData)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDynamicPageCache)
	}
	herr.Message = "Can't get to the GetCacheDynamicPageCacheViaData's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDynamicPageCacheViaExpireHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iExpire := self.Args("expire").MustInt()
	if helper.IsHas(iExpire) {
		_CacheDynamicPageCache, _error := model.GetCacheDynamicPageCacheViaExpire(iExpire)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDynamicPageCache)
	}
	herr.Message = "Can't get to the GetCacheDynamicPageCacheViaExpire's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDynamicPageCacheViaCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iCreated := self.Args("created").String()
	if helper.IsHas(iCreated) {
		_CacheDynamicPageCache, _error := model.GetCacheDynamicPageCacheViaCreated(iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDynamicPageCache)
	}
	herr.Message = "Can't get to the GetCacheDynamicPageCacheViaCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDynamicPageCacheViaSerializedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iSerialized := self.Args("serialized").MustInt()
	if helper.IsHas(iSerialized) {
		_CacheDynamicPageCache, _error := model.GetCacheDynamicPageCacheViaSerialized(iSerialized)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDynamicPageCache)
	}
	herr.Message = "Can't get to the GetCacheDynamicPageCacheViaSerialized's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDynamicPageCacheViaTagsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iTags := self.Args("tags").String()
	if helper.IsHas(iTags) {
		_CacheDynamicPageCache, _error := model.GetCacheDynamicPageCacheViaTags(iTags)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDynamicPageCache)
	}
	herr.Message = "Can't get to the GetCacheDynamicPageCacheViaTags's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDynamicPageCacheViaChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iChecksum := self.Args("checksum").String()
	if helper.IsHas(iChecksum) {
		_CacheDynamicPageCache, _error := model.GetCacheDynamicPageCacheViaChecksum(iChecksum)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDynamicPageCache)
	}
	herr.Message = "Can't get to the GetCacheDynamicPageCacheViaChecksum's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetCacheDynamicPageCacheViaCidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Cid_ := self.Args("cid").String()
	if helper.IsHas(Cid_) {
		var iCacheDynamicPageCache model.CacheDynamicPageCache
		self.Bind(&iCacheDynamicPageCache)
		_CacheDynamicPageCache, _error := model.SetCacheDynamicPageCacheViaCid(Cid_, &iCacheDynamicPageCache)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDynamicPageCache)
	}
	herr.Message = "Can't get to the SetCacheDynamicPageCacheViaCid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetCacheDynamicPageCacheViaDataHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Data_ := self.Args("data").Bytes()
	if helper.IsHas(Data_) {
		var iCacheDynamicPageCache model.CacheDynamicPageCache
		self.Bind(&iCacheDynamicPageCache)
		_CacheDynamicPageCache, _error := model.SetCacheDynamicPageCacheViaData(Data_, &iCacheDynamicPageCache)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDynamicPageCache)
	}
	herr.Message = "Can't get to the SetCacheDynamicPageCacheViaData's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetCacheDynamicPageCacheViaExpireHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Expire_ := self.Args("expire").MustInt()
	if helper.IsHas(Expire_) {
		var iCacheDynamicPageCache model.CacheDynamicPageCache
		self.Bind(&iCacheDynamicPageCache)
		_CacheDynamicPageCache, _error := model.SetCacheDynamicPageCacheViaExpire(Expire_, &iCacheDynamicPageCache)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDynamicPageCache)
	}
	herr.Message = "Can't get to the SetCacheDynamicPageCacheViaExpire's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetCacheDynamicPageCacheViaCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Created_ := self.Args("created").String()
	if helper.IsHas(Created_) {
		var iCacheDynamicPageCache model.CacheDynamicPageCache
		self.Bind(&iCacheDynamicPageCache)
		_CacheDynamicPageCache, _error := model.SetCacheDynamicPageCacheViaCreated(Created_, &iCacheDynamicPageCache)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDynamicPageCache)
	}
	herr.Message = "Can't get to the SetCacheDynamicPageCacheViaCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetCacheDynamicPageCacheViaSerializedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Serialized_ := self.Args("serialized").MustInt()
	if helper.IsHas(Serialized_) {
		var iCacheDynamicPageCache model.CacheDynamicPageCache
		self.Bind(&iCacheDynamicPageCache)
		_CacheDynamicPageCache, _error := model.SetCacheDynamicPageCacheViaSerialized(Serialized_, &iCacheDynamicPageCache)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDynamicPageCache)
	}
	herr.Message = "Can't get to the SetCacheDynamicPageCacheViaSerialized's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetCacheDynamicPageCacheViaTagsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Tags_ := self.Args("tags").String()
	if helper.IsHas(Tags_) {
		var iCacheDynamicPageCache model.CacheDynamicPageCache
		self.Bind(&iCacheDynamicPageCache)
		_CacheDynamicPageCache, _error := model.SetCacheDynamicPageCacheViaTags(Tags_, &iCacheDynamicPageCache)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDynamicPageCache)
	}
	herr.Message = "Can't get to the SetCacheDynamicPageCacheViaTags's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetCacheDynamicPageCacheViaChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Checksum_ := self.Args("checksum").String()
	if helper.IsHas(Checksum_) {
		var iCacheDynamicPageCache model.CacheDynamicPageCache
		self.Bind(&iCacheDynamicPageCache)
		_CacheDynamicPageCache, _error := model.SetCacheDynamicPageCacheViaChecksum(Checksum_, &iCacheDynamicPageCache)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDynamicPageCache)
	}
	herr.Message = "Can't get to the SetCacheDynamicPageCacheViaChecksum's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostAddCacheDynamicPageCacheHandler(self *macross.Context) error {
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
		_error := model.AddCacheDynamicPageCache(Cid_,Data_,Expire_,Created_,Serialized_,Tags_,Checksum_)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		herr.Message = "StatusOK"
		herr.Status = macross.StatusOK
		return self.JSON(herr)
	}
	herr.Message = "Can't get to the AddCacheDynamicPageCache's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostCacheDynamicPageCacheHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iCacheDynamicPageCache model.CacheDynamicPageCache
	self.Bind(&iCacheDynamicPageCache)
	_string, _error := model.PostCacheDynamicPageCache(&iCacheDynamicPageCache)
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

func PutCacheDynamicPageCacheHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iCacheDynamicPageCache model.CacheDynamicPageCache
	self.Bind(&iCacheDynamicPageCache)
	_string, _error := model.PutCacheDynamicPageCache(&iCacheDynamicPageCache)
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

func PutCacheDynamicPageCacheViaCidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Cid_ := self.Args("cid").String()
	var iCacheDynamicPageCache model.CacheDynamicPageCache
	self.Bind(&iCacheDynamicPageCache)
	_int64, _error := model.PutCacheDynamicPageCacheViaCid(Cid_, &iCacheDynamicPageCache)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutCacheDynamicPageCacheViaDataHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Data_ := self.Args("data").Bytes()
	var iCacheDynamicPageCache model.CacheDynamicPageCache
	self.Bind(&iCacheDynamicPageCache)
	_int64, _error := model.PutCacheDynamicPageCacheViaData(Data_, &iCacheDynamicPageCache)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutCacheDynamicPageCacheViaExpireHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Expire_ := self.Args("expire").MustInt()
	var iCacheDynamicPageCache model.CacheDynamicPageCache
	self.Bind(&iCacheDynamicPageCache)
	_int64, _error := model.PutCacheDynamicPageCacheViaExpire(Expire_, &iCacheDynamicPageCache)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutCacheDynamicPageCacheViaCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Created_ := self.Args("created").String()
	var iCacheDynamicPageCache model.CacheDynamicPageCache
	self.Bind(&iCacheDynamicPageCache)
	_int64, _error := model.PutCacheDynamicPageCacheViaCreated(Created_, &iCacheDynamicPageCache)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutCacheDynamicPageCacheViaSerializedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Serialized_ := self.Args("serialized").MustInt()
	var iCacheDynamicPageCache model.CacheDynamicPageCache
	self.Bind(&iCacheDynamicPageCache)
	_int64, _error := model.PutCacheDynamicPageCacheViaSerialized(Serialized_, &iCacheDynamicPageCache)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutCacheDynamicPageCacheViaTagsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Tags_ := self.Args("tags").String()
	var iCacheDynamicPageCache model.CacheDynamicPageCache
	self.Bind(&iCacheDynamicPageCache)
	_int64, _error := model.PutCacheDynamicPageCacheViaTags(Tags_, &iCacheDynamicPageCache)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutCacheDynamicPageCacheViaChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Checksum_ := self.Args("checksum").String()
	var iCacheDynamicPageCache model.CacheDynamicPageCache
	self.Bind(&iCacheDynamicPageCache)
	_int64, _error := model.PutCacheDynamicPageCacheViaChecksum(Checksum_, &iCacheDynamicPageCache)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateCacheDynamicPageCacheViaCidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Cid_ := self.Args("cid").String()
	var iCacheDynamicPageCache model.CacheDynamicPageCache
	self.Bind(&iCacheDynamicPageCache)
	var iMap = helper.StructToMap(iCacheDynamicPageCache)
	_error := model.UpdateCacheDynamicPageCacheViaCid(Cid_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateCacheDynamicPageCacheViaDataHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Data_ := self.Args("data").Bytes()
	var iCacheDynamicPageCache model.CacheDynamicPageCache
	self.Bind(&iCacheDynamicPageCache)
	var iMap = helper.StructToMap(iCacheDynamicPageCache)
	_error := model.UpdateCacheDynamicPageCacheViaData(Data_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateCacheDynamicPageCacheViaExpireHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Expire_ := self.Args("expire").MustInt()
	var iCacheDynamicPageCache model.CacheDynamicPageCache
	self.Bind(&iCacheDynamicPageCache)
	var iMap = helper.StructToMap(iCacheDynamicPageCache)
	_error := model.UpdateCacheDynamicPageCacheViaExpire(Expire_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateCacheDynamicPageCacheViaCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Created_ := self.Args("created").String()
	var iCacheDynamicPageCache model.CacheDynamicPageCache
	self.Bind(&iCacheDynamicPageCache)
	var iMap = helper.StructToMap(iCacheDynamicPageCache)
	_error := model.UpdateCacheDynamicPageCacheViaCreated(Created_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateCacheDynamicPageCacheViaSerializedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Serialized_ := self.Args("serialized").MustInt()
	var iCacheDynamicPageCache model.CacheDynamicPageCache
	self.Bind(&iCacheDynamicPageCache)
	var iMap = helper.StructToMap(iCacheDynamicPageCache)
	_error := model.UpdateCacheDynamicPageCacheViaSerialized(Serialized_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateCacheDynamicPageCacheViaTagsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Tags_ := self.Args("tags").String()
	var iCacheDynamicPageCache model.CacheDynamicPageCache
	self.Bind(&iCacheDynamicPageCache)
	var iMap = helper.StructToMap(iCacheDynamicPageCache)
	_error := model.UpdateCacheDynamicPageCacheViaTags(Tags_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateCacheDynamicPageCacheViaChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Checksum_ := self.Args("checksum").String()
	var iCacheDynamicPageCache model.CacheDynamicPageCache
	self.Bind(&iCacheDynamicPageCache)
	var iMap = helper.StructToMap(iCacheDynamicPageCache)
	_error := model.UpdateCacheDynamicPageCacheViaChecksum(Checksum_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteCacheDynamicPageCacheHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Cid_ := self.Args("cid").String()
	_int64, _error := model.DeleteCacheDynamicPageCache(Cid_)
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

func DeleteCacheDynamicPageCacheViaCidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Cid_ := self.Args("cid").String()
	_error := model.DeleteCacheDynamicPageCacheViaCid(Cid_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteCacheDynamicPageCacheViaDataHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Data_ := self.Args("data").Bytes()
	_error := model.DeleteCacheDynamicPageCacheViaData(Data_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteCacheDynamicPageCacheViaExpireHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Expire_ := self.Args("expire").MustInt()
	_error := model.DeleteCacheDynamicPageCacheViaExpire(Expire_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteCacheDynamicPageCacheViaCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Created_ := self.Args("created").String()
	_error := model.DeleteCacheDynamicPageCacheViaCreated(Created_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteCacheDynamicPageCacheViaSerializedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Serialized_ := self.Args("serialized").MustInt()
	_error := model.DeleteCacheDynamicPageCacheViaSerialized(Serialized_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteCacheDynamicPageCacheViaTagsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Tags_ := self.Args("tags").String()
	_error := model.DeleteCacheDynamicPageCacheViaTags(Tags_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteCacheDynamicPageCacheViaChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Checksum_ := self.Args("checksum").String()
	_error := model.DeleteCacheDynamicPageCacheViaChecksum(Checksum_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}
