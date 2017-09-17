package handler

import (
	"github.com/insionng/drupapp/helper"
	"github.com/insionng/drupapp/model"
	"github.com/insionng/macross"
)

func GetCacheDiscoveriesCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()

	if offset > 0 {
		var m = map[string]interface{}{}
		_int64, _error := model.GetCacheDiscoveriesCount(offset, limit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		m["cache_discoverysCount"] = _int64
		return self.JSON(m)
	}
	herr.Message = "Can't get to the GetCacheDiscoveriesCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDiscoveryCountViaCidHandler(self *macross.Context) error {
	Cid_ := self.Args("cid").String()
	_int64 := model.GetCacheDiscoveryCountViaCid(Cid_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["cache_discoveryCount"] = 0
	}
	m["cache_discoveryCount"] = _int64
	return self.JSON(m)
}

func GetCacheDiscoveryCountViaDataHandler(self *macross.Context) error {
	Data_ := self.Args("data").Bytes()
	_int64 := model.GetCacheDiscoveryCountViaData(Data_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["cache_discoveryCount"] = 0
	}
	m["cache_discoveryCount"] = _int64
	return self.JSON(m)
}

func GetCacheDiscoveryCountViaExpireHandler(self *macross.Context) error {
	Expire_ := self.Args("expire").MustInt()
	_int64 := model.GetCacheDiscoveryCountViaExpire(Expire_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["cache_discoveryCount"] = 0
	}
	m["cache_discoveryCount"] = _int64
	return self.JSON(m)
}

func GetCacheDiscoveryCountViaCreatedHandler(self *macross.Context) error {
	Created_ := self.Args("created").String()
	_int64 := model.GetCacheDiscoveryCountViaCreated(Created_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["cache_discoveryCount"] = 0
	}
	m["cache_discoveryCount"] = _int64
	return self.JSON(m)
}

func GetCacheDiscoveryCountViaSerializedHandler(self *macross.Context) error {
	Serialized_ := self.Args("serialized").MustInt()
	_int64 := model.GetCacheDiscoveryCountViaSerialized(Serialized_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["cache_discoveryCount"] = 0
	}
	m["cache_discoveryCount"] = _int64
	return self.JSON(m)
}

func GetCacheDiscoveryCountViaTagsHandler(self *macross.Context) error {
	Tags_ := self.Args("tags").String()
	_int64 := model.GetCacheDiscoveryCountViaTags(Tags_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["cache_discoveryCount"] = 0
	}
	m["cache_discoveryCount"] = _int64
	return self.JSON(m)
}

func GetCacheDiscoveryCountViaChecksumHandler(self *macross.Context) error {
	Checksum_ := self.Args("checksum").String()
	_int64 := model.GetCacheDiscoveryCountViaChecksum(Checksum_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["cache_discoveryCount"] = 0
	}
	m["cache_discoveryCount"] = _int64
	return self.JSON(m)
}

func GetCacheDiscoveriesViaCidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iCid := self.Args("cid").String()
	if (offset > 0) && helper.IsHas(iCid) {
		_CacheDiscovery, _error := model.GetCacheDiscoveriesViaCid(offset, limit, iCid, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDiscovery)
	}
	herr.Message = "Can't get to the GetCacheDiscoveriesViaCid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDiscoveriesViaDataHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iData := self.Args("data").Bytes()
	if (offset > 0) && helper.IsHas(iData) {
		_CacheDiscovery, _error := model.GetCacheDiscoveriesViaData(offset, limit, iData, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDiscovery)
	}
	herr.Message = "Can't get to the GetCacheDiscoveriesViaData's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDiscoveriesViaExpireHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iExpire := self.Args("expire").MustInt()
	if (offset > 0) && helper.IsHas(iExpire) {
		_CacheDiscovery, _error := model.GetCacheDiscoveriesViaExpire(offset, limit, iExpire, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDiscovery)
	}
	herr.Message = "Can't get to the GetCacheDiscoveriesViaExpire's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDiscoveriesViaCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iCreated := self.Args("created").String()
	if (offset > 0) && helper.IsHas(iCreated) {
		_CacheDiscovery, _error := model.GetCacheDiscoveriesViaCreated(offset, limit, iCreated, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDiscovery)
	}
	herr.Message = "Can't get to the GetCacheDiscoveriesViaCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDiscoveriesViaSerializedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iSerialized := self.Args("serialized").MustInt()
	if (offset > 0) && helper.IsHas(iSerialized) {
		_CacheDiscovery, _error := model.GetCacheDiscoveriesViaSerialized(offset, limit, iSerialized, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDiscovery)
	}
	herr.Message = "Can't get to the GetCacheDiscoveriesViaSerialized's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDiscoveriesViaTagsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iTags := self.Args("tags").String()
	if (offset > 0) && helper.IsHas(iTags) {
		_CacheDiscovery, _error := model.GetCacheDiscoveriesViaTags(offset, limit, iTags, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDiscovery)
	}
	herr.Message = "Can't get to the GetCacheDiscoveriesViaTags's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDiscoveriesViaChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iChecksum := self.Args("checksum").String()
	if (offset > 0) && helper.IsHas(iChecksum) {
		_CacheDiscovery, _error := model.GetCacheDiscoveriesViaChecksum(offset, limit, iChecksum, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDiscovery)
	}
	herr.Message = "Can't get to the GetCacheDiscoveriesViaChecksum's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDiscoveriesByCidAndDataAndExpireHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").String()
	iData := self.Args("data").Bytes()
	iExpire := self.Args("expire").MustInt()

	if helper.IsHas(iCid) {
		_CacheDiscovery, _error := model.GetCacheDiscoveriesByCidAndDataAndExpire(offset, limit, iCid,iData,iExpire)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDiscovery)
	}
	herr.Message = "Can't get to the GetCacheDiscoveriesByCidAndDataAndExpire's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDiscoveriesByCidAndDataAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").String()
	iData := self.Args("data").Bytes()
	iCreated := self.Args("created").String()

	if helper.IsHas(iCid) {
		_CacheDiscovery, _error := model.GetCacheDiscoveriesByCidAndDataAndCreated(offset, limit, iCid,iData,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDiscovery)
	}
	herr.Message = "Can't get to the GetCacheDiscoveriesByCidAndDataAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDiscoveriesByCidAndDataAndSerializedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").String()
	iData := self.Args("data").Bytes()
	iSerialized := self.Args("serialized").MustInt()

	if helper.IsHas(iCid) {
		_CacheDiscovery, _error := model.GetCacheDiscoveriesByCidAndDataAndSerialized(offset, limit, iCid,iData,iSerialized)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDiscovery)
	}
	herr.Message = "Can't get to the GetCacheDiscoveriesByCidAndDataAndSerialized's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDiscoveriesByCidAndDataAndTagsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").String()
	iData := self.Args("data").Bytes()
	iTags := self.Args("tags").String()

	if helper.IsHas(iCid) {
		_CacheDiscovery, _error := model.GetCacheDiscoveriesByCidAndDataAndTags(offset, limit, iCid,iData,iTags)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDiscovery)
	}
	herr.Message = "Can't get to the GetCacheDiscoveriesByCidAndDataAndTags's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDiscoveriesByCidAndDataAndChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").String()
	iData := self.Args("data").Bytes()
	iChecksum := self.Args("checksum").String()

	if helper.IsHas(iCid) {
		_CacheDiscovery, _error := model.GetCacheDiscoveriesByCidAndDataAndChecksum(offset, limit, iCid,iData,iChecksum)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDiscovery)
	}
	herr.Message = "Can't get to the GetCacheDiscoveriesByCidAndDataAndChecksum's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDiscoveriesByCidAndExpireAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").String()
	iExpire := self.Args("expire").MustInt()
	iCreated := self.Args("created").String()

	if helper.IsHas(iCid) {
		_CacheDiscovery, _error := model.GetCacheDiscoveriesByCidAndExpireAndCreated(offset, limit, iCid,iExpire,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDiscovery)
	}
	herr.Message = "Can't get to the GetCacheDiscoveriesByCidAndExpireAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDiscoveriesByCidAndExpireAndSerializedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").String()
	iExpire := self.Args("expire").MustInt()
	iSerialized := self.Args("serialized").MustInt()

	if helper.IsHas(iCid) {
		_CacheDiscovery, _error := model.GetCacheDiscoveriesByCidAndExpireAndSerialized(offset, limit, iCid,iExpire,iSerialized)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDiscovery)
	}
	herr.Message = "Can't get to the GetCacheDiscoveriesByCidAndExpireAndSerialized's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDiscoveriesByCidAndExpireAndTagsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").String()
	iExpire := self.Args("expire").MustInt()
	iTags := self.Args("tags").String()

	if helper.IsHas(iCid) {
		_CacheDiscovery, _error := model.GetCacheDiscoveriesByCidAndExpireAndTags(offset, limit, iCid,iExpire,iTags)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDiscovery)
	}
	herr.Message = "Can't get to the GetCacheDiscoveriesByCidAndExpireAndTags's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDiscoveriesByCidAndExpireAndChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").String()
	iExpire := self.Args("expire").MustInt()
	iChecksum := self.Args("checksum").String()

	if helper.IsHas(iCid) {
		_CacheDiscovery, _error := model.GetCacheDiscoveriesByCidAndExpireAndChecksum(offset, limit, iCid,iExpire,iChecksum)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDiscovery)
	}
	herr.Message = "Can't get to the GetCacheDiscoveriesByCidAndExpireAndChecksum's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDiscoveriesByCidAndCreatedAndSerializedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").String()
	iCreated := self.Args("created").String()
	iSerialized := self.Args("serialized").MustInt()

	if helper.IsHas(iCid) {
		_CacheDiscovery, _error := model.GetCacheDiscoveriesByCidAndCreatedAndSerialized(offset, limit, iCid,iCreated,iSerialized)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDiscovery)
	}
	herr.Message = "Can't get to the GetCacheDiscoveriesByCidAndCreatedAndSerialized's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDiscoveriesByCidAndCreatedAndTagsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").String()
	iCreated := self.Args("created").String()
	iTags := self.Args("tags").String()

	if helper.IsHas(iCid) {
		_CacheDiscovery, _error := model.GetCacheDiscoveriesByCidAndCreatedAndTags(offset, limit, iCid,iCreated,iTags)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDiscovery)
	}
	herr.Message = "Can't get to the GetCacheDiscoveriesByCidAndCreatedAndTags's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDiscoveriesByCidAndCreatedAndChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").String()
	iCreated := self.Args("created").String()
	iChecksum := self.Args("checksum").String()

	if helper.IsHas(iCid) {
		_CacheDiscovery, _error := model.GetCacheDiscoveriesByCidAndCreatedAndChecksum(offset, limit, iCid,iCreated,iChecksum)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDiscovery)
	}
	herr.Message = "Can't get to the GetCacheDiscoveriesByCidAndCreatedAndChecksum's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDiscoveriesByCidAndSerializedAndTagsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").String()
	iSerialized := self.Args("serialized").MustInt()
	iTags := self.Args("tags").String()

	if helper.IsHas(iCid) {
		_CacheDiscovery, _error := model.GetCacheDiscoveriesByCidAndSerializedAndTags(offset, limit, iCid,iSerialized,iTags)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDiscovery)
	}
	herr.Message = "Can't get to the GetCacheDiscoveriesByCidAndSerializedAndTags's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDiscoveriesByCidAndSerializedAndChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").String()
	iSerialized := self.Args("serialized").MustInt()
	iChecksum := self.Args("checksum").String()

	if helper.IsHas(iCid) {
		_CacheDiscovery, _error := model.GetCacheDiscoveriesByCidAndSerializedAndChecksum(offset, limit, iCid,iSerialized,iChecksum)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDiscovery)
	}
	herr.Message = "Can't get to the GetCacheDiscoveriesByCidAndSerializedAndChecksum's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDiscoveriesByCidAndTagsAndChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").String()
	iTags := self.Args("tags").String()
	iChecksum := self.Args("checksum").String()

	if helper.IsHas(iCid) {
		_CacheDiscovery, _error := model.GetCacheDiscoveriesByCidAndTagsAndChecksum(offset, limit, iCid,iTags,iChecksum)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDiscovery)
	}
	herr.Message = "Can't get to the GetCacheDiscoveriesByCidAndTagsAndChecksum's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDiscoveriesByDataAndExpireAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iData := self.Args("data").Bytes()
	iExpire := self.Args("expire").MustInt()
	iCreated := self.Args("created").String()

	if helper.IsHas(iData) {
		_CacheDiscovery, _error := model.GetCacheDiscoveriesByDataAndExpireAndCreated(offset, limit, iData,iExpire,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDiscovery)
	}
	herr.Message = "Can't get to the GetCacheDiscoveriesByDataAndExpireAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDiscoveriesByDataAndExpireAndSerializedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iData := self.Args("data").Bytes()
	iExpire := self.Args("expire").MustInt()
	iSerialized := self.Args("serialized").MustInt()

	if helper.IsHas(iData) {
		_CacheDiscovery, _error := model.GetCacheDiscoveriesByDataAndExpireAndSerialized(offset, limit, iData,iExpire,iSerialized)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDiscovery)
	}
	herr.Message = "Can't get to the GetCacheDiscoveriesByDataAndExpireAndSerialized's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDiscoveriesByDataAndExpireAndTagsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iData := self.Args("data").Bytes()
	iExpire := self.Args("expire").MustInt()
	iTags := self.Args("tags").String()

	if helper.IsHas(iData) {
		_CacheDiscovery, _error := model.GetCacheDiscoveriesByDataAndExpireAndTags(offset, limit, iData,iExpire,iTags)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDiscovery)
	}
	herr.Message = "Can't get to the GetCacheDiscoveriesByDataAndExpireAndTags's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDiscoveriesByDataAndExpireAndChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iData := self.Args("data").Bytes()
	iExpire := self.Args("expire").MustInt()
	iChecksum := self.Args("checksum").String()

	if helper.IsHas(iData) {
		_CacheDiscovery, _error := model.GetCacheDiscoveriesByDataAndExpireAndChecksum(offset, limit, iData,iExpire,iChecksum)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDiscovery)
	}
	herr.Message = "Can't get to the GetCacheDiscoveriesByDataAndExpireAndChecksum's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDiscoveriesByDataAndCreatedAndSerializedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iData := self.Args("data").Bytes()
	iCreated := self.Args("created").String()
	iSerialized := self.Args("serialized").MustInt()

	if helper.IsHas(iData) {
		_CacheDiscovery, _error := model.GetCacheDiscoveriesByDataAndCreatedAndSerialized(offset, limit, iData,iCreated,iSerialized)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDiscovery)
	}
	herr.Message = "Can't get to the GetCacheDiscoveriesByDataAndCreatedAndSerialized's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDiscoveriesByDataAndCreatedAndTagsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iData := self.Args("data").Bytes()
	iCreated := self.Args("created").String()
	iTags := self.Args("tags").String()

	if helper.IsHas(iData) {
		_CacheDiscovery, _error := model.GetCacheDiscoveriesByDataAndCreatedAndTags(offset, limit, iData,iCreated,iTags)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDiscovery)
	}
	herr.Message = "Can't get to the GetCacheDiscoveriesByDataAndCreatedAndTags's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDiscoveriesByDataAndCreatedAndChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iData := self.Args("data").Bytes()
	iCreated := self.Args("created").String()
	iChecksum := self.Args("checksum").String()

	if helper.IsHas(iData) {
		_CacheDiscovery, _error := model.GetCacheDiscoveriesByDataAndCreatedAndChecksum(offset, limit, iData,iCreated,iChecksum)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDiscovery)
	}
	herr.Message = "Can't get to the GetCacheDiscoveriesByDataAndCreatedAndChecksum's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDiscoveriesByDataAndSerializedAndTagsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iData := self.Args("data").Bytes()
	iSerialized := self.Args("serialized").MustInt()
	iTags := self.Args("tags").String()

	if helper.IsHas(iData) {
		_CacheDiscovery, _error := model.GetCacheDiscoveriesByDataAndSerializedAndTags(offset, limit, iData,iSerialized,iTags)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDiscovery)
	}
	herr.Message = "Can't get to the GetCacheDiscoveriesByDataAndSerializedAndTags's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDiscoveriesByDataAndSerializedAndChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iData := self.Args("data").Bytes()
	iSerialized := self.Args("serialized").MustInt()
	iChecksum := self.Args("checksum").String()

	if helper.IsHas(iData) {
		_CacheDiscovery, _error := model.GetCacheDiscoveriesByDataAndSerializedAndChecksum(offset, limit, iData,iSerialized,iChecksum)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDiscovery)
	}
	herr.Message = "Can't get to the GetCacheDiscoveriesByDataAndSerializedAndChecksum's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDiscoveriesByDataAndTagsAndChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iData := self.Args("data").Bytes()
	iTags := self.Args("tags").String()
	iChecksum := self.Args("checksum").String()

	if helper.IsHas(iData) {
		_CacheDiscovery, _error := model.GetCacheDiscoveriesByDataAndTagsAndChecksum(offset, limit, iData,iTags,iChecksum)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDiscovery)
	}
	herr.Message = "Can't get to the GetCacheDiscoveriesByDataAndTagsAndChecksum's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDiscoveriesByExpireAndCreatedAndSerializedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iExpire := self.Args("expire").MustInt()
	iCreated := self.Args("created").String()
	iSerialized := self.Args("serialized").MustInt()

	if helper.IsHas(iExpire) {
		_CacheDiscovery, _error := model.GetCacheDiscoveriesByExpireAndCreatedAndSerialized(offset, limit, iExpire,iCreated,iSerialized)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDiscovery)
	}
	herr.Message = "Can't get to the GetCacheDiscoveriesByExpireAndCreatedAndSerialized's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDiscoveriesByExpireAndCreatedAndTagsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iExpire := self.Args("expire").MustInt()
	iCreated := self.Args("created").String()
	iTags := self.Args("tags").String()

	if helper.IsHas(iExpire) {
		_CacheDiscovery, _error := model.GetCacheDiscoveriesByExpireAndCreatedAndTags(offset, limit, iExpire,iCreated,iTags)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDiscovery)
	}
	herr.Message = "Can't get to the GetCacheDiscoveriesByExpireAndCreatedAndTags's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDiscoveriesByExpireAndCreatedAndChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iExpire := self.Args("expire").MustInt()
	iCreated := self.Args("created").String()
	iChecksum := self.Args("checksum").String()

	if helper.IsHas(iExpire) {
		_CacheDiscovery, _error := model.GetCacheDiscoveriesByExpireAndCreatedAndChecksum(offset, limit, iExpire,iCreated,iChecksum)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDiscovery)
	}
	herr.Message = "Can't get to the GetCacheDiscoveriesByExpireAndCreatedAndChecksum's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDiscoveriesByExpireAndSerializedAndTagsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iExpire := self.Args("expire").MustInt()
	iSerialized := self.Args("serialized").MustInt()
	iTags := self.Args("tags").String()

	if helper.IsHas(iExpire) {
		_CacheDiscovery, _error := model.GetCacheDiscoveriesByExpireAndSerializedAndTags(offset, limit, iExpire,iSerialized,iTags)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDiscovery)
	}
	herr.Message = "Can't get to the GetCacheDiscoveriesByExpireAndSerializedAndTags's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDiscoveriesByExpireAndSerializedAndChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iExpire := self.Args("expire").MustInt()
	iSerialized := self.Args("serialized").MustInt()
	iChecksum := self.Args("checksum").String()

	if helper.IsHas(iExpire) {
		_CacheDiscovery, _error := model.GetCacheDiscoveriesByExpireAndSerializedAndChecksum(offset, limit, iExpire,iSerialized,iChecksum)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDiscovery)
	}
	herr.Message = "Can't get to the GetCacheDiscoveriesByExpireAndSerializedAndChecksum's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDiscoveriesByExpireAndTagsAndChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iExpire := self.Args("expire").MustInt()
	iTags := self.Args("tags").String()
	iChecksum := self.Args("checksum").String()

	if helper.IsHas(iExpire) {
		_CacheDiscovery, _error := model.GetCacheDiscoveriesByExpireAndTagsAndChecksum(offset, limit, iExpire,iTags,iChecksum)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDiscovery)
	}
	herr.Message = "Can't get to the GetCacheDiscoveriesByExpireAndTagsAndChecksum's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDiscoveriesByCreatedAndSerializedAndTagsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCreated := self.Args("created").String()
	iSerialized := self.Args("serialized").MustInt()
	iTags := self.Args("tags").String()

	if helper.IsHas(iCreated) {
		_CacheDiscovery, _error := model.GetCacheDiscoveriesByCreatedAndSerializedAndTags(offset, limit, iCreated,iSerialized,iTags)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDiscovery)
	}
	herr.Message = "Can't get to the GetCacheDiscoveriesByCreatedAndSerializedAndTags's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDiscoveriesByCreatedAndSerializedAndChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCreated := self.Args("created").String()
	iSerialized := self.Args("serialized").MustInt()
	iChecksum := self.Args("checksum").String()

	if helper.IsHas(iCreated) {
		_CacheDiscovery, _error := model.GetCacheDiscoveriesByCreatedAndSerializedAndChecksum(offset, limit, iCreated,iSerialized,iChecksum)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDiscovery)
	}
	herr.Message = "Can't get to the GetCacheDiscoveriesByCreatedAndSerializedAndChecksum's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDiscoveriesByCreatedAndTagsAndChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCreated := self.Args("created").String()
	iTags := self.Args("tags").String()
	iChecksum := self.Args("checksum").String()

	if helper.IsHas(iCreated) {
		_CacheDiscovery, _error := model.GetCacheDiscoveriesByCreatedAndTagsAndChecksum(offset, limit, iCreated,iTags,iChecksum)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDiscovery)
	}
	herr.Message = "Can't get to the GetCacheDiscoveriesByCreatedAndTagsAndChecksum's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDiscoveriesBySerializedAndTagsAndChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSerialized := self.Args("serialized").MustInt()
	iTags := self.Args("tags").String()
	iChecksum := self.Args("checksum").String()

	if helper.IsHas(iSerialized) {
		_CacheDiscovery, _error := model.GetCacheDiscoveriesBySerializedAndTagsAndChecksum(offset, limit, iSerialized,iTags,iChecksum)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDiscovery)
	}
	herr.Message = "Can't get to the GetCacheDiscoveriesBySerializedAndTagsAndChecksum's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDiscoveriesByCidAndDataHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").String()
	iData := self.Args("data").Bytes()

	if helper.IsHas(iCid) {
		_CacheDiscovery, _error := model.GetCacheDiscoveriesByCidAndData(offset, limit, iCid,iData)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDiscovery)
	}
	herr.Message = "Can't get to the GetCacheDiscoveriesByCidAndData's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDiscoveriesByCidAndExpireHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").String()
	iExpire := self.Args("expire").MustInt()

	if helper.IsHas(iCid) {
		_CacheDiscovery, _error := model.GetCacheDiscoveriesByCidAndExpire(offset, limit, iCid,iExpire)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDiscovery)
	}
	herr.Message = "Can't get to the GetCacheDiscoveriesByCidAndExpire's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDiscoveriesByCidAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").String()
	iCreated := self.Args("created").String()

	if helper.IsHas(iCid) {
		_CacheDiscovery, _error := model.GetCacheDiscoveriesByCidAndCreated(offset, limit, iCid,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDiscovery)
	}
	herr.Message = "Can't get to the GetCacheDiscoveriesByCidAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDiscoveriesByCidAndSerializedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").String()
	iSerialized := self.Args("serialized").MustInt()

	if helper.IsHas(iCid) {
		_CacheDiscovery, _error := model.GetCacheDiscoveriesByCidAndSerialized(offset, limit, iCid,iSerialized)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDiscovery)
	}
	herr.Message = "Can't get to the GetCacheDiscoveriesByCidAndSerialized's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDiscoveriesByCidAndTagsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").String()
	iTags := self.Args("tags").String()

	if helper.IsHas(iCid) {
		_CacheDiscovery, _error := model.GetCacheDiscoveriesByCidAndTags(offset, limit, iCid,iTags)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDiscovery)
	}
	herr.Message = "Can't get to the GetCacheDiscoveriesByCidAndTags's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDiscoveriesByCidAndChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").String()
	iChecksum := self.Args("checksum").String()

	if helper.IsHas(iCid) {
		_CacheDiscovery, _error := model.GetCacheDiscoveriesByCidAndChecksum(offset, limit, iCid,iChecksum)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDiscovery)
	}
	herr.Message = "Can't get to the GetCacheDiscoveriesByCidAndChecksum's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDiscoveriesByDataAndExpireHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iData := self.Args("data").Bytes()
	iExpire := self.Args("expire").MustInt()

	if helper.IsHas(iData) {
		_CacheDiscovery, _error := model.GetCacheDiscoveriesByDataAndExpire(offset, limit, iData,iExpire)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDiscovery)
	}
	herr.Message = "Can't get to the GetCacheDiscoveriesByDataAndExpire's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDiscoveriesByDataAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iData := self.Args("data").Bytes()
	iCreated := self.Args("created").String()

	if helper.IsHas(iData) {
		_CacheDiscovery, _error := model.GetCacheDiscoveriesByDataAndCreated(offset, limit, iData,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDiscovery)
	}
	herr.Message = "Can't get to the GetCacheDiscoveriesByDataAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDiscoveriesByDataAndSerializedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iData := self.Args("data").Bytes()
	iSerialized := self.Args("serialized").MustInt()

	if helper.IsHas(iData) {
		_CacheDiscovery, _error := model.GetCacheDiscoveriesByDataAndSerialized(offset, limit, iData,iSerialized)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDiscovery)
	}
	herr.Message = "Can't get to the GetCacheDiscoveriesByDataAndSerialized's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDiscoveriesByDataAndTagsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iData := self.Args("data").Bytes()
	iTags := self.Args("tags").String()

	if helper.IsHas(iData) {
		_CacheDiscovery, _error := model.GetCacheDiscoveriesByDataAndTags(offset, limit, iData,iTags)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDiscovery)
	}
	herr.Message = "Can't get to the GetCacheDiscoveriesByDataAndTags's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDiscoveriesByDataAndChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iData := self.Args("data").Bytes()
	iChecksum := self.Args("checksum").String()

	if helper.IsHas(iData) {
		_CacheDiscovery, _error := model.GetCacheDiscoveriesByDataAndChecksum(offset, limit, iData,iChecksum)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDiscovery)
	}
	herr.Message = "Can't get to the GetCacheDiscoveriesByDataAndChecksum's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDiscoveriesByExpireAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iExpire := self.Args("expire").MustInt()
	iCreated := self.Args("created").String()

	if helper.IsHas(iExpire) {
		_CacheDiscovery, _error := model.GetCacheDiscoveriesByExpireAndCreated(offset, limit, iExpire,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDiscovery)
	}
	herr.Message = "Can't get to the GetCacheDiscoveriesByExpireAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDiscoveriesByExpireAndSerializedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iExpire := self.Args("expire").MustInt()
	iSerialized := self.Args("serialized").MustInt()

	if helper.IsHas(iExpire) {
		_CacheDiscovery, _error := model.GetCacheDiscoveriesByExpireAndSerialized(offset, limit, iExpire,iSerialized)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDiscovery)
	}
	herr.Message = "Can't get to the GetCacheDiscoveriesByExpireAndSerialized's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDiscoveriesByExpireAndTagsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iExpire := self.Args("expire").MustInt()
	iTags := self.Args("tags").String()

	if helper.IsHas(iExpire) {
		_CacheDiscovery, _error := model.GetCacheDiscoveriesByExpireAndTags(offset, limit, iExpire,iTags)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDiscovery)
	}
	herr.Message = "Can't get to the GetCacheDiscoveriesByExpireAndTags's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDiscoveriesByExpireAndChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iExpire := self.Args("expire").MustInt()
	iChecksum := self.Args("checksum").String()

	if helper.IsHas(iExpire) {
		_CacheDiscovery, _error := model.GetCacheDiscoveriesByExpireAndChecksum(offset, limit, iExpire,iChecksum)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDiscovery)
	}
	herr.Message = "Can't get to the GetCacheDiscoveriesByExpireAndChecksum's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDiscoveriesByCreatedAndSerializedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCreated := self.Args("created").String()
	iSerialized := self.Args("serialized").MustInt()

	if helper.IsHas(iCreated) {
		_CacheDiscovery, _error := model.GetCacheDiscoveriesByCreatedAndSerialized(offset, limit, iCreated,iSerialized)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDiscovery)
	}
	herr.Message = "Can't get to the GetCacheDiscoveriesByCreatedAndSerialized's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDiscoveriesByCreatedAndTagsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCreated := self.Args("created").String()
	iTags := self.Args("tags").String()

	if helper.IsHas(iCreated) {
		_CacheDiscovery, _error := model.GetCacheDiscoveriesByCreatedAndTags(offset, limit, iCreated,iTags)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDiscovery)
	}
	herr.Message = "Can't get to the GetCacheDiscoveriesByCreatedAndTags's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDiscoveriesByCreatedAndChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCreated := self.Args("created").String()
	iChecksum := self.Args("checksum").String()

	if helper.IsHas(iCreated) {
		_CacheDiscovery, _error := model.GetCacheDiscoveriesByCreatedAndChecksum(offset, limit, iCreated,iChecksum)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDiscovery)
	}
	herr.Message = "Can't get to the GetCacheDiscoveriesByCreatedAndChecksum's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDiscoveriesBySerializedAndTagsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSerialized := self.Args("serialized").MustInt()
	iTags := self.Args("tags").String()

	if helper.IsHas(iSerialized) {
		_CacheDiscovery, _error := model.GetCacheDiscoveriesBySerializedAndTags(offset, limit, iSerialized,iTags)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDiscovery)
	}
	herr.Message = "Can't get to the GetCacheDiscoveriesBySerializedAndTags's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDiscoveriesBySerializedAndChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSerialized := self.Args("serialized").MustInt()
	iChecksum := self.Args("checksum").String()

	if helper.IsHas(iSerialized) {
		_CacheDiscovery, _error := model.GetCacheDiscoveriesBySerializedAndChecksum(offset, limit, iSerialized,iChecksum)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDiscovery)
	}
	herr.Message = "Can't get to the GetCacheDiscoveriesBySerializedAndChecksum's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDiscoveriesByTagsAndChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTags := self.Args("tags").String()
	iChecksum := self.Args("checksum").String()

	if helper.IsHas(iTags) {
		_CacheDiscovery, _error := model.GetCacheDiscoveriesByTagsAndChecksum(offset, limit, iTags,iChecksum)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDiscovery)
	}
	herr.Message = "Can't get to the GetCacheDiscoveriesByTagsAndChecksum's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDiscoveriesHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	if (offset > 0) && (len(field) > 0) {
		_CacheDiscovery, _error := model.GetCacheDiscoveries(offset, limit, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDiscovery)
	}
	herr.Message = "Can't get to the GetCacheDiscoveries' args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasCacheDiscoveryViaCidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iCid := self.Args("cid").String()
	if helper.IsHas(iCid) {
		_CacheDiscovery := model.HasCacheDiscoveryViaCid(iCid)
		var m = map[string]interface{}{}
		m["cache_discovery"] = _CacheDiscovery
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasCacheDiscoveryViaCid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasCacheDiscoveryViaDataHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iData := self.Args("data").Bytes()
	if helper.IsHas(iData) {
		_CacheDiscovery := model.HasCacheDiscoveryViaData(iData)
		var m = map[string]interface{}{}
		m["cache_discovery"] = _CacheDiscovery
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasCacheDiscoveryViaData's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasCacheDiscoveryViaExpireHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iExpire := self.Args("expire").MustInt()
	if helper.IsHas(iExpire) {
		_CacheDiscovery := model.HasCacheDiscoveryViaExpire(iExpire)
		var m = map[string]interface{}{}
		m["cache_discovery"] = _CacheDiscovery
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasCacheDiscoveryViaExpire's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasCacheDiscoveryViaCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iCreated := self.Args("created").String()
	if helper.IsHas(iCreated) {
		_CacheDiscovery := model.HasCacheDiscoveryViaCreated(iCreated)
		var m = map[string]interface{}{}
		m["cache_discovery"] = _CacheDiscovery
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasCacheDiscoveryViaCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasCacheDiscoveryViaSerializedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iSerialized := self.Args("serialized").MustInt()
	if helper.IsHas(iSerialized) {
		_CacheDiscovery := model.HasCacheDiscoveryViaSerialized(iSerialized)
		var m = map[string]interface{}{}
		m["cache_discovery"] = _CacheDiscovery
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasCacheDiscoveryViaSerialized's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasCacheDiscoveryViaTagsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iTags := self.Args("tags").String()
	if helper.IsHas(iTags) {
		_CacheDiscovery := model.HasCacheDiscoveryViaTags(iTags)
		var m = map[string]interface{}{}
		m["cache_discovery"] = _CacheDiscovery
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasCacheDiscoveryViaTags's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasCacheDiscoveryViaChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iChecksum := self.Args("checksum").String()
	if helper.IsHas(iChecksum) {
		_CacheDiscovery := model.HasCacheDiscoveryViaChecksum(iChecksum)
		var m = map[string]interface{}{}
		m["cache_discovery"] = _CacheDiscovery
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasCacheDiscoveryViaChecksum's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDiscoveryViaCidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iCid := self.Args("cid").String()
	if helper.IsHas(iCid) {
		_CacheDiscovery, _error := model.GetCacheDiscoveryViaCid(iCid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDiscovery)
	}
	herr.Message = "Can't get to the GetCacheDiscoveryViaCid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDiscoveryViaDataHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iData := self.Args("data").Bytes()
	if helper.IsHas(iData) {
		_CacheDiscovery, _error := model.GetCacheDiscoveryViaData(iData)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDiscovery)
	}
	herr.Message = "Can't get to the GetCacheDiscoveryViaData's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDiscoveryViaExpireHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iExpire := self.Args("expire").MustInt()
	if helper.IsHas(iExpire) {
		_CacheDiscovery, _error := model.GetCacheDiscoveryViaExpire(iExpire)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDiscovery)
	}
	herr.Message = "Can't get to the GetCacheDiscoveryViaExpire's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDiscoveryViaCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iCreated := self.Args("created").String()
	if helper.IsHas(iCreated) {
		_CacheDiscovery, _error := model.GetCacheDiscoveryViaCreated(iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDiscovery)
	}
	herr.Message = "Can't get to the GetCacheDiscoveryViaCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDiscoveryViaSerializedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iSerialized := self.Args("serialized").MustInt()
	if helper.IsHas(iSerialized) {
		_CacheDiscovery, _error := model.GetCacheDiscoveryViaSerialized(iSerialized)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDiscovery)
	}
	herr.Message = "Can't get to the GetCacheDiscoveryViaSerialized's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDiscoveryViaTagsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iTags := self.Args("tags").String()
	if helper.IsHas(iTags) {
		_CacheDiscovery, _error := model.GetCacheDiscoveryViaTags(iTags)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDiscovery)
	}
	herr.Message = "Can't get to the GetCacheDiscoveryViaTags's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheDiscoveryViaChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iChecksum := self.Args("checksum").String()
	if helper.IsHas(iChecksum) {
		_CacheDiscovery, _error := model.GetCacheDiscoveryViaChecksum(iChecksum)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDiscovery)
	}
	herr.Message = "Can't get to the GetCacheDiscoveryViaChecksum's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetCacheDiscoveryViaCidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Cid_ := self.Args("cid").String()
	if helper.IsHas(Cid_) {
		var iCacheDiscovery model.CacheDiscovery
		self.Bind(&iCacheDiscovery)
		_CacheDiscovery, _error := model.SetCacheDiscoveryViaCid(Cid_, &iCacheDiscovery)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDiscovery)
	}
	herr.Message = "Can't get to the SetCacheDiscoveryViaCid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetCacheDiscoveryViaDataHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Data_ := self.Args("data").Bytes()
	if helper.IsHas(Data_) {
		var iCacheDiscovery model.CacheDiscovery
		self.Bind(&iCacheDiscovery)
		_CacheDiscovery, _error := model.SetCacheDiscoveryViaData(Data_, &iCacheDiscovery)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDiscovery)
	}
	herr.Message = "Can't get to the SetCacheDiscoveryViaData's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetCacheDiscoveryViaExpireHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Expire_ := self.Args("expire").MustInt()
	if helper.IsHas(Expire_) {
		var iCacheDiscovery model.CacheDiscovery
		self.Bind(&iCacheDiscovery)
		_CacheDiscovery, _error := model.SetCacheDiscoveryViaExpire(Expire_, &iCacheDiscovery)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDiscovery)
	}
	herr.Message = "Can't get to the SetCacheDiscoveryViaExpire's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetCacheDiscoveryViaCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Created_ := self.Args("created").String()
	if helper.IsHas(Created_) {
		var iCacheDiscovery model.CacheDiscovery
		self.Bind(&iCacheDiscovery)
		_CacheDiscovery, _error := model.SetCacheDiscoveryViaCreated(Created_, &iCacheDiscovery)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDiscovery)
	}
	herr.Message = "Can't get to the SetCacheDiscoveryViaCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetCacheDiscoveryViaSerializedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Serialized_ := self.Args("serialized").MustInt()
	if helper.IsHas(Serialized_) {
		var iCacheDiscovery model.CacheDiscovery
		self.Bind(&iCacheDiscovery)
		_CacheDiscovery, _error := model.SetCacheDiscoveryViaSerialized(Serialized_, &iCacheDiscovery)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDiscovery)
	}
	herr.Message = "Can't get to the SetCacheDiscoveryViaSerialized's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetCacheDiscoveryViaTagsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Tags_ := self.Args("tags").String()
	if helper.IsHas(Tags_) {
		var iCacheDiscovery model.CacheDiscovery
		self.Bind(&iCacheDiscovery)
		_CacheDiscovery, _error := model.SetCacheDiscoveryViaTags(Tags_, &iCacheDiscovery)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDiscovery)
	}
	herr.Message = "Can't get to the SetCacheDiscoveryViaTags's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetCacheDiscoveryViaChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Checksum_ := self.Args("checksum").String()
	if helper.IsHas(Checksum_) {
		var iCacheDiscovery model.CacheDiscovery
		self.Bind(&iCacheDiscovery)
		_CacheDiscovery, _error := model.SetCacheDiscoveryViaChecksum(Checksum_, &iCacheDiscovery)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheDiscovery)
	}
	herr.Message = "Can't get to the SetCacheDiscoveryViaChecksum's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostAddCacheDiscoveryHandler(self *macross.Context) error {
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
		_error := model.AddCacheDiscovery(Cid_,Data_,Expire_,Created_,Serialized_,Tags_,Checksum_)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		herr.Message = "StatusOK"
		herr.Status = macross.StatusOK
		return self.JSON(herr)
	}
	herr.Message = "Can't get to the AddCacheDiscovery's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostCacheDiscoveryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iCacheDiscovery model.CacheDiscovery
	self.Bind(&iCacheDiscovery)
	_string, _error := model.PostCacheDiscovery(&iCacheDiscovery)
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

func PutCacheDiscoveryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iCacheDiscovery model.CacheDiscovery
	self.Bind(&iCacheDiscovery)
	_string, _error := model.PutCacheDiscovery(&iCacheDiscovery)
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

func PutCacheDiscoveryViaCidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Cid_ := self.Args("cid").String()
	var iCacheDiscovery model.CacheDiscovery
	self.Bind(&iCacheDiscovery)
	_int64, _error := model.PutCacheDiscoveryViaCid(Cid_, &iCacheDiscovery)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutCacheDiscoveryViaDataHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Data_ := self.Args("data").Bytes()
	var iCacheDiscovery model.CacheDiscovery
	self.Bind(&iCacheDiscovery)
	_int64, _error := model.PutCacheDiscoveryViaData(Data_, &iCacheDiscovery)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutCacheDiscoveryViaExpireHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Expire_ := self.Args("expire").MustInt()
	var iCacheDiscovery model.CacheDiscovery
	self.Bind(&iCacheDiscovery)
	_int64, _error := model.PutCacheDiscoveryViaExpire(Expire_, &iCacheDiscovery)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutCacheDiscoveryViaCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Created_ := self.Args("created").String()
	var iCacheDiscovery model.CacheDiscovery
	self.Bind(&iCacheDiscovery)
	_int64, _error := model.PutCacheDiscoveryViaCreated(Created_, &iCacheDiscovery)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutCacheDiscoveryViaSerializedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Serialized_ := self.Args("serialized").MustInt()
	var iCacheDiscovery model.CacheDiscovery
	self.Bind(&iCacheDiscovery)
	_int64, _error := model.PutCacheDiscoveryViaSerialized(Serialized_, &iCacheDiscovery)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutCacheDiscoveryViaTagsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Tags_ := self.Args("tags").String()
	var iCacheDiscovery model.CacheDiscovery
	self.Bind(&iCacheDiscovery)
	_int64, _error := model.PutCacheDiscoveryViaTags(Tags_, &iCacheDiscovery)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutCacheDiscoveryViaChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Checksum_ := self.Args("checksum").String()
	var iCacheDiscovery model.CacheDiscovery
	self.Bind(&iCacheDiscovery)
	_int64, _error := model.PutCacheDiscoveryViaChecksum(Checksum_, &iCacheDiscovery)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateCacheDiscoveryViaCidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Cid_ := self.Args("cid").String()
	var iCacheDiscovery model.CacheDiscovery
	self.Bind(&iCacheDiscovery)
	var iMap = helper.StructToMap(iCacheDiscovery)
	_error := model.UpdateCacheDiscoveryViaCid(Cid_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateCacheDiscoveryViaDataHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Data_ := self.Args("data").Bytes()
	var iCacheDiscovery model.CacheDiscovery
	self.Bind(&iCacheDiscovery)
	var iMap = helper.StructToMap(iCacheDiscovery)
	_error := model.UpdateCacheDiscoveryViaData(Data_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateCacheDiscoveryViaExpireHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Expire_ := self.Args("expire").MustInt()
	var iCacheDiscovery model.CacheDiscovery
	self.Bind(&iCacheDiscovery)
	var iMap = helper.StructToMap(iCacheDiscovery)
	_error := model.UpdateCacheDiscoveryViaExpire(Expire_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateCacheDiscoveryViaCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Created_ := self.Args("created").String()
	var iCacheDiscovery model.CacheDiscovery
	self.Bind(&iCacheDiscovery)
	var iMap = helper.StructToMap(iCacheDiscovery)
	_error := model.UpdateCacheDiscoveryViaCreated(Created_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateCacheDiscoveryViaSerializedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Serialized_ := self.Args("serialized").MustInt()
	var iCacheDiscovery model.CacheDiscovery
	self.Bind(&iCacheDiscovery)
	var iMap = helper.StructToMap(iCacheDiscovery)
	_error := model.UpdateCacheDiscoveryViaSerialized(Serialized_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateCacheDiscoveryViaTagsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Tags_ := self.Args("tags").String()
	var iCacheDiscovery model.CacheDiscovery
	self.Bind(&iCacheDiscovery)
	var iMap = helper.StructToMap(iCacheDiscovery)
	_error := model.UpdateCacheDiscoveryViaTags(Tags_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateCacheDiscoveryViaChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Checksum_ := self.Args("checksum").String()
	var iCacheDiscovery model.CacheDiscovery
	self.Bind(&iCacheDiscovery)
	var iMap = helper.StructToMap(iCacheDiscovery)
	_error := model.UpdateCacheDiscoveryViaChecksum(Checksum_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteCacheDiscoveryHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Cid_ := self.Args("cid").String()
	_int64, _error := model.DeleteCacheDiscovery(Cid_)
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

func DeleteCacheDiscoveryViaCidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Cid_ := self.Args("cid").String()
	_error := model.DeleteCacheDiscoveryViaCid(Cid_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteCacheDiscoveryViaDataHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Data_ := self.Args("data").Bytes()
	_error := model.DeleteCacheDiscoveryViaData(Data_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteCacheDiscoveryViaExpireHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Expire_ := self.Args("expire").MustInt()
	_error := model.DeleteCacheDiscoveryViaExpire(Expire_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteCacheDiscoveryViaCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Created_ := self.Args("created").String()
	_error := model.DeleteCacheDiscoveryViaCreated(Created_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteCacheDiscoveryViaSerializedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Serialized_ := self.Args("serialized").MustInt()
	_error := model.DeleteCacheDiscoveryViaSerialized(Serialized_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteCacheDiscoveryViaTagsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Tags_ := self.Args("tags").String()
	_error := model.DeleteCacheDiscoveryViaTags(Tags_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteCacheDiscoveryViaChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Checksum_ := self.Args("checksum").String()
	_error := model.DeleteCacheDiscoveryViaChecksum(Checksum_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}
