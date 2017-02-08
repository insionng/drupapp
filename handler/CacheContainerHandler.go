package handler

import (
	"drupapp/helper"
	"drupapp/model"
	"github.com/insionng/macross"
)

func GetCacheContainersCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()

	if offset > 0 {
		var m = map[string]interface{}{}
		_int64, _error := model.GetCacheContainersCount(offset, limit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		m["cache_containersCount"] = _int64
		return self.JSON(m)
	}
	herr.Message = "Can't get to the GetCacheContainersCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheContainerCountViaCidHandler(self *macross.Context) error {
	Cid_ := self.Args("cid").String()
	_int64 := model.GetCacheContainerCountViaCid(Cid_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["cache_containerCount"] = 0
	}
	m["cache_containerCount"] = _int64
	return self.JSON(m)
}

func GetCacheContainerCountViaDataHandler(self *macross.Context) error {
	Data_ := self.Args("data").Bytes()
	_int64 := model.GetCacheContainerCountViaData(Data_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["cache_containerCount"] = 0
	}
	m["cache_containerCount"] = _int64
	return self.JSON(m)
}

func GetCacheContainerCountViaExpireHandler(self *macross.Context) error {
	Expire_ := self.Args("expire").MustInt()
	_int64 := model.GetCacheContainerCountViaExpire(Expire_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["cache_containerCount"] = 0
	}
	m["cache_containerCount"] = _int64
	return self.JSON(m)
}

func GetCacheContainerCountViaCreatedHandler(self *macross.Context) error {
	Created_ := self.Args("created").String()
	_int64 := model.GetCacheContainerCountViaCreated(Created_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["cache_containerCount"] = 0
	}
	m["cache_containerCount"] = _int64
	return self.JSON(m)
}

func GetCacheContainerCountViaSerializedHandler(self *macross.Context) error {
	Serialized_ := self.Args("serialized").MustInt()
	_int64 := model.GetCacheContainerCountViaSerialized(Serialized_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["cache_containerCount"] = 0
	}
	m["cache_containerCount"] = _int64
	return self.JSON(m)
}

func GetCacheContainerCountViaTagsHandler(self *macross.Context) error {
	Tags_ := self.Args("tags").String()
	_int64 := model.GetCacheContainerCountViaTags(Tags_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["cache_containerCount"] = 0
	}
	m["cache_containerCount"] = _int64
	return self.JSON(m)
}

func GetCacheContainerCountViaChecksumHandler(self *macross.Context) error {
	Checksum_ := self.Args("checksum").String()
	_int64 := model.GetCacheContainerCountViaChecksum(Checksum_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["cache_containerCount"] = 0
	}
	m["cache_containerCount"] = _int64
	return self.JSON(m)
}

func GetCacheContainersViaCidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iCid := self.Args("cid").String()
	if (offset > 0) && helper.IsHas(iCid) {
		_CacheContainer, _error := model.GetCacheContainersViaCid(offset, limit, iCid, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheContainer)
	}
	herr.Message = "Can't get to the GetCacheContainersViaCid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheContainersViaDataHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iData := self.Args("data").Bytes()
	if (offset > 0) && helper.IsHas(iData) {
		_CacheContainer, _error := model.GetCacheContainersViaData(offset, limit, iData, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheContainer)
	}
	herr.Message = "Can't get to the GetCacheContainersViaData's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheContainersViaExpireHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iExpire := self.Args("expire").MustInt()
	if (offset > 0) && helper.IsHas(iExpire) {
		_CacheContainer, _error := model.GetCacheContainersViaExpire(offset, limit, iExpire, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheContainer)
	}
	herr.Message = "Can't get to the GetCacheContainersViaExpire's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheContainersViaCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iCreated := self.Args("created").String()
	if (offset > 0) && helper.IsHas(iCreated) {
		_CacheContainer, _error := model.GetCacheContainersViaCreated(offset, limit, iCreated, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheContainer)
	}
	herr.Message = "Can't get to the GetCacheContainersViaCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheContainersViaSerializedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iSerialized := self.Args("serialized").MustInt()
	if (offset > 0) && helper.IsHas(iSerialized) {
		_CacheContainer, _error := model.GetCacheContainersViaSerialized(offset, limit, iSerialized, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheContainer)
	}
	herr.Message = "Can't get to the GetCacheContainersViaSerialized's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheContainersViaTagsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iTags := self.Args("tags").String()
	if (offset > 0) && helper.IsHas(iTags) {
		_CacheContainer, _error := model.GetCacheContainersViaTags(offset, limit, iTags, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheContainer)
	}
	herr.Message = "Can't get to the GetCacheContainersViaTags's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheContainersViaChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iChecksum := self.Args("checksum").String()
	if (offset > 0) && helper.IsHas(iChecksum) {
		_CacheContainer, _error := model.GetCacheContainersViaChecksum(offset, limit, iChecksum, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheContainer)
	}
	herr.Message = "Can't get to the GetCacheContainersViaChecksum's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheContainersByCidAndDataAndExpireHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").String()
	iData := self.Args("data").Bytes()
	iExpire := self.Args("expire").MustInt()

	if helper.IsHas(iCid) {
		_CacheContainer, _error := model.GetCacheContainersByCidAndDataAndExpire(offset, limit, iCid,iData,iExpire)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheContainer)
	}
	herr.Message = "Can't get to the GetCacheContainersByCidAndDataAndExpire's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheContainersByCidAndDataAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").String()
	iData := self.Args("data").Bytes()
	iCreated := self.Args("created").String()

	if helper.IsHas(iCid) {
		_CacheContainer, _error := model.GetCacheContainersByCidAndDataAndCreated(offset, limit, iCid,iData,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheContainer)
	}
	herr.Message = "Can't get to the GetCacheContainersByCidAndDataAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheContainersByCidAndDataAndSerializedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").String()
	iData := self.Args("data").Bytes()
	iSerialized := self.Args("serialized").MustInt()

	if helper.IsHas(iCid) {
		_CacheContainer, _error := model.GetCacheContainersByCidAndDataAndSerialized(offset, limit, iCid,iData,iSerialized)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheContainer)
	}
	herr.Message = "Can't get to the GetCacheContainersByCidAndDataAndSerialized's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheContainersByCidAndDataAndTagsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").String()
	iData := self.Args("data").Bytes()
	iTags := self.Args("tags").String()

	if helper.IsHas(iCid) {
		_CacheContainer, _error := model.GetCacheContainersByCidAndDataAndTags(offset, limit, iCid,iData,iTags)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheContainer)
	}
	herr.Message = "Can't get to the GetCacheContainersByCidAndDataAndTags's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheContainersByCidAndDataAndChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").String()
	iData := self.Args("data").Bytes()
	iChecksum := self.Args("checksum").String()

	if helper.IsHas(iCid) {
		_CacheContainer, _error := model.GetCacheContainersByCidAndDataAndChecksum(offset, limit, iCid,iData,iChecksum)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheContainer)
	}
	herr.Message = "Can't get to the GetCacheContainersByCidAndDataAndChecksum's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheContainersByCidAndExpireAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").String()
	iExpire := self.Args("expire").MustInt()
	iCreated := self.Args("created").String()

	if helper.IsHas(iCid) {
		_CacheContainer, _error := model.GetCacheContainersByCidAndExpireAndCreated(offset, limit, iCid,iExpire,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheContainer)
	}
	herr.Message = "Can't get to the GetCacheContainersByCidAndExpireAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheContainersByCidAndExpireAndSerializedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").String()
	iExpire := self.Args("expire").MustInt()
	iSerialized := self.Args("serialized").MustInt()

	if helper.IsHas(iCid) {
		_CacheContainer, _error := model.GetCacheContainersByCidAndExpireAndSerialized(offset, limit, iCid,iExpire,iSerialized)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheContainer)
	}
	herr.Message = "Can't get to the GetCacheContainersByCidAndExpireAndSerialized's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheContainersByCidAndExpireAndTagsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").String()
	iExpire := self.Args("expire").MustInt()
	iTags := self.Args("tags").String()

	if helper.IsHas(iCid) {
		_CacheContainer, _error := model.GetCacheContainersByCidAndExpireAndTags(offset, limit, iCid,iExpire,iTags)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheContainer)
	}
	herr.Message = "Can't get to the GetCacheContainersByCidAndExpireAndTags's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheContainersByCidAndExpireAndChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").String()
	iExpire := self.Args("expire").MustInt()
	iChecksum := self.Args("checksum").String()

	if helper.IsHas(iCid) {
		_CacheContainer, _error := model.GetCacheContainersByCidAndExpireAndChecksum(offset, limit, iCid,iExpire,iChecksum)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheContainer)
	}
	herr.Message = "Can't get to the GetCacheContainersByCidAndExpireAndChecksum's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheContainersByCidAndCreatedAndSerializedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").String()
	iCreated := self.Args("created").String()
	iSerialized := self.Args("serialized").MustInt()

	if helper.IsHas(iCid) {
		_CacheContainer, _error := model.GetCacheContainersByCidAndCreatedAndSerialized(offset, limit, iCid,iCreated,iSerialized)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheContainer)
	}
	herr.Message = "Can't get to the GetCacheContainersByCidAndCreatedAndSerialized's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheContainersByCidAndCreatedAndTagsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").String()
	iCreated := self.Args("created").String()
	iTags := self.Args("tags").String()

	if helper.IsHas(iCid) {
		_CacheContainer, _error := model.GetCacheContainersByCidAndCreatedAndTags(offset, limit, iCid,iCreated,iTags)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheContainer)
	}
	herr.Message = "Can't get to the GetCacheContainersByCidAndCreatedAndTags's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheContainersByCidAndCreatedAndChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").String()
	iCreated := self.Args("created").String()
	iChecksum := self.Args("checksum").String()

	if helper.IsHas(iCid) {
		_CacheContainer, _error := model.GetCacheContainersByCidAndCreatedAndChecksum(offset, limit, iCid,iCreated,iChecksum)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheContainer)
	}
	herr.Message = "Can't get to the GetCacheContainersByCidAndCreatedAndChecksum's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheContainersByCidAndSerializedAndTagsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").String()
	iSerialized := self.Args("serialized").MustInt()
	iTags := self.Args("tags").String()

	if helper.IsHas(iCid) {
		_CacheContainer, _error := model.GetCacheContainersByCidAndSerializedAndTags(offset, limit, iCid,iSerialized,iTags)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheContainer)
	}
	herr.Message = "Can't get to the GetCacheContainersByCidAndSerializedAndTags's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheContainersByCidAndSerializedAndChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").String()
	iSerialized := self.Args("serialized").MustInt()
	iChecksum := self.Args("checksum").String()

	if helper.IsHas(iCid) {
		_CacheContainer, _error := model.GetCacheContainersByCidAndSerializedAndChecksum(offset, limit, iCid,iSerialized,iChecksum)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheContainer)
	}
	herr.Message = "Can't get to the GetCacheContainersByCidAndSerializedAndChecksum's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheContainersByCidAndTagsAndChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").String()
	iTags := self.Args("tags").String()
	iChecksum := self.Args("checksum").String()

	if helper.IsHas(iCid) {
		_CacheContainer, _error := model.GetCacheContainersByCidAndTagsAndChecksum(offset, limit, iCid,iTags,iChecksum)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheContainer)
	}
	herr.Message = "Can't get to the GetCacheContainersByCidAndTagsAndChecksum's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheContainersByDataAndExpireAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iData := self.Args("data").Bytes()
	iExpire := self.Args("expire").MustInt()
	iCreated := self.Args("created").String()

	if helper.IsHas(iData) {
		_CacheContainer, _error := model.GetCacheContainersByDataAndExpireAndCreated(offset, limit, iData,iExpire,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheContainer)
	}
	herr.Message = "Can't get to the GetCacheContainersByDataAndExpireAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheContainersByDataAndExpireAndSerializedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iData := self.Args("data").Bytes()
	iExpire := self.Args("expire").MustInt()
	iSerialized := self.Args("serialized").MustInt()

	if helper.IsHas(iData) {
		_CacheContainer, _error := model.GetCacheContainersByDataAndExpireAndSerialized(offset, limit, iData,iExpire,iSerialized)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheContainer)
	}
	herr.Message = "Can't get to the GetCacheContainersByDataAndExpireAndSerialized's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheContainersByDataAndExpireAndTagsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iData := self.Args("data").Bytes()
	iExpire := self.Args("expire").MustInt()
	iTags := self.Args("tags").String()

	if helper.IsHas(iData) {
		_CacheContainer, _error := model.GetCacheContainersByDataAndExpireAndTags(offset, limit, iData,iExpire,iTags)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheContainer)
	}
	herr.Message = "Can't get to the GetCacheContainersByDataAndExpireAndTags's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheContainersByDataAndExpireAndChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iData := self.Args("data").Bytes()
	iExpire := self.Args("expire").MustInt()
	iChecksum := self.Args("checksum").String()

	if helper.IsHas(iData) {
		_CacheContainer, _error := model.GetCacheContainersByDataAndExpireAndChecksum(offset, limit, iData,iExpire,iChecksum)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheContainer)
	}
	herr.Message = "Can't get to the GetCacheContainersByDataAndExpireAndChecksum's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheContainersByDataAndCreatedAndSerializedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iData := self.Args("data").Bytes()
	iCreated := self.Args("created").String()
	iSerialized := self.Args("serialized").MustInt()

	if helper.IsHas(iData) {
		_CacheContainer, _error := model.GetCacheContainersByDataAndCreatedAndSerialized(offset, limit, iData,iCreated,iSerialized)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheContainer)
	}
	herr.Message = "Can't get to the GetCacheContainersByDataAndCreatedAndSerialized's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheContainersByDataAndCreatedAndTagsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iData := self.Args("data").Bytes()
	iCreated := self.Args("created").String()
	iTags := self.Args("tags").String()

	if helper.IsHas(iData) {
		_CacheContainer, _error := model.GetCacheContainersByDataAndCreatedAndTags(offset, limit, iData,iCreated,iTags)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheContainer)
	}
	herr.Message = "Can't get to the GetCacheContainersByDataAndCreatedAndTags's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheContainersByDataAndCreatedAndChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iData := self.Args("data").Bytes()
	iCreated := self.Args("created").String()
	iChecksum := self.Args("checksum").String()

	if helper.IsHas(iData) {
		_CacheContainer, _error := model.GetCacheContainersByDataAndCreatedAndChecksum(offset, limit, iData,iCreated,iChecksum)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheContainer)
	}
	herr.Message = "Can't get to the GetCacheContainersByDataAndCreatedAndChecksum's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheContainersByDataAndSerializedAndTagsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iData := self.Args("data").Bytes()
	iSerialized := self.Args("serialized").MustInt()
	iTags := self.Args("tags").String()

	if helper.IsHas(iData) {
		_CacheContainer, _error := model.GetCacheContainersByDataAndSerializedAndTags(offset, limit, iData,iSerialized,iTags)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheContainer)
	}
	herr.Message = "Can't get to the GetCacheContainersByDataAndSerializedAndTags's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheContainersByDataAndSerializedAndChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iData := self.Args("data").Bytes()
	iSerialized := self.Args("serialized").MustInt()
	iChecksum := self.Args("checksum").String()

	if helper.IsHas(iData) {
		_CacheContainer, _error := model.GetCacheContainersByDataAndSerializedAndChecksum(offset, limit, iData,iSerialized,iChecksum)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheContainer)
	}
	herr.Message = "Can't get to the GetCacheContainersByDataAndSerializedAndChecksum's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheContainersByDataAndTagsAndChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iData := self.Args("data").Bytes()
	iTags := self.Args("tags").String()
	iChecksum := self.Args("checksum").String()

	if helper.IsHas(iData) {
		_CacheContainer, _error := model.GetCacheContainersByDataAndTagsAndChecksum(offset, limit, iData,iTags,iChecksum)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheContainer)
	}
	herr.Message = "Can't get to the GetCacheContainersByDataAndTagsAndChecksum's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheContainersByExpireAndCreatedAndSerializedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iExpire := self.Args("expire").MustInt()
	iCreated := self.Args("created").String()
	iSerialized := self.Args("serialized").MustInt()

	if helper.IsHas(iExpire) {
		_CacheContainer, _error := model.GetCacheContainersByExpireAndCreatedAndSerialized(offset, limit, iExpire,iCreated,iSerialized)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheContainer)
	}
	herr.Message = "Can't get to the GetCacheContainersByExpireAndCreatedAndSerialized's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheContainersByExpireAndCreatedAndTagsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iExpire := self.Args("expire").MustInt()
	iCreated := self.Args("created").String()
	iTags := self.Args("tags").String()

	if helper.IsHas(iExpire) {
		_CacheContainer, _error := model.GetCacheContainersByExpireAndCreatedAndTags(offset, limit, iExpire,iCreated,iTags)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheContainer)
	}
	herr.Message = "Can't get to the GetCacheContainersByExpireAndCreatedAndTags's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheContainersByExpireAndCreatedAndChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iExpire := self.Args("expire").MustInt()
	iCreated := self.Args("created").String()
	iChecksum := self.Args("checksum").String()

	if helper.IsHas(iExpire) {
		_CacheContainer, _error := model.GetCacheContainersByExpireAndCreatedAndChecksum(offset, limit, iExpire,iCreated,iChecksum)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheContainer)
	}
	herr.Message = "Can't get to the GetCacheContainersByExpireAndCreatedAndChecksum's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheContainersByExpireAndSerializedAndTagsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iExpire := self.Args("expire").MustInt()
	iSerialized := self.Args("serialized").MustInt()
	iTags := self.Args("tags").String()

	if helper.IsHas(iExpire) {
		_CacheContainer, _error := model.GetCacheContainersByExpireAndSerializedAndTags(offset, limit, iExpire,iSerialized,iTags)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheContainer)
	}
	herr.Message = "Can't get to the GetCacheContainersByExpireAndSerializedAndTags's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheContainersByExpireAndSerializedAndChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iExpire := self.Args("expire").MustInt()
	iSerialized := self.Args("serialized").MustInt()
	iChecksum := self.Args("checksum").String()

	if helper.IsHas(iExpire) {
		_CacheContainer, _error := model.GetCacheContainersByExpireAndSerializedAndChecksum(offset, limit, iExpire,iSerialized,iChecksum)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheContainer)
	}
	herr.Message = "Can't get to the GetCacheContainersByExpireAndSerializedAndChecksum's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheContainersByExpireAndTagsAndChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iExpire := self.Args("expire").MustInt()
	iTags := self.Args("tags").String()
	iChecksum := self.Args("checksum").String()

	if helper.IsHas(iExpire) {
		_CacheContainer, _error := model.GetCacheContainersByExpireAndTagsAndChecksum(offset, limit, iExpire,iTags,iChecksum)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheContainer)
	}
	herr.Message = "Can't get to the GetCacheContainersByExpireAndTagsAndChecksum's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheContainersByCreatedAndSerializedAndTagsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCreated := self.Args("created").String()
	iSerialized := self.Args("serialized").MustInt()
	iTags := self.Args("tags").String()

	if helper.IsHas(iCreated) {
		_CacheContainer, _error := model.GetCacheContainersByCreatedAndSerializedAndTags(offset, limit, iCreated,iSerialized,iTags)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheContainer)
	}
	herr.Message = "Can't get to the GetCacheContainersByCreatedAndSerializedAndTags's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheContainersByCreatedAndSerializedAndChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCreated := self.Args("created").String()
	iSerialized := self.Args("serialized").MustInt()
	iChecksum := self.Args("checksum").String()

	if helper.IsHas(iCreated) {
		_CacheContainer, _error := model.GetCacheContainersByCreatedAndSerializedAndChecksum(offset, limit, iCreated,iSerialized,iChecksum)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheContainer)
	}
	herr.Message = "Can't get to the GetCacheContainersByCreatedAndSerializedAndChecksum's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheContainersByCreatedAndTagsAndChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCreated := self.Args("created").String()
	iTags := self.Args("tags").String()
	iChecksum := self.Args("checksum").String()

	if helper.IsHas(iCreated) {
		_CacheContainer, _error := model.GetCacheContainersByCreatedAndTagsAndChecksum(offset, limit, iCreated,iTags,iChecksum)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheContainer)
	}
	herr.Message = "Can't get to the GetCacheContainersByCreatedAndTagsAndChecksum's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheContainersBySerializedAndTagsAndChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSerialized := self.Args("serialized").MustInt()
	iTags := self.Args("tags").String()
	iChecksum := self.Args("checksum").String()

	if helper.IsHas(iSerialized) {
		_CacheContainer, _error := model.GetCacheContainersBySerializedAndTagsAndChecksum(offset, limit, iSerialized,iTags,iChecksum)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheContainer)
	}
	herr.Message = "Can't get to the GetCacheContainersBySerializedAndTagsAndChecksum's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheContainersByCidAndDataHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").String()
	iData := self.Args("data").Bytes()

	if helper.IsHas(iCid) {
		_CacheContainer, _error := model.GetCacheContainersByCidAndData(offset, limit, iCid,iData)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheContainer)
	}
	herr.Message = "Can't get to the GetCacheContainersByCidAndData's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheContainersByCidAndExpireHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").String()
	iExpire := self.Args("expire").MustInt()

	if helper.IsHas(iCid) {
		_CacheContainer, _error := model.GetCacheContainersByCidAndExpire(offset, limit, iCid,iExpire)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheContainer)
	}
	herr.Message = "Can't get to the GetCacheContainersByCidAndExpire's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheContainersByCidAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").String()
	iCreated := self.Args("created").String()

	if helper.IsHas(iCid) {
		_CacheContainer, _error := model.GetCacheContainersByCidAndCreated(offset, limit, iCid,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheContainer)
	}
	herr.Message = "Can't get to the GetCacheContainersByCidAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheContainersByCidAndSerializedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").String()
	iSerialized := self.Args("serialized").MustInt()

	if helper.IsHas(iCid) {
		_CacheContainer, _error := model.GetCacheContainersByCidAndSerialized(offset, limit, iCid,iSerialized)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheContainer)
	}
	herr.Message = "Can't get to the GetCacheContainersByCidAndSerialized's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheContainersByCidAndTagsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").String()
	iTags := self.Args("tags").String()

	if helper.IsHas(iCid) {
		_CacheContainer, _error := model.GetCacheContainersByCidAndTags(offset, limit, iCid,iTags)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheContainer)
	}
	herr.Message = "Can't get to the GetCacheContainersByCidAndTags's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheContainersByCidAndChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").String()
	iChecksum := self.Args("checksum").String()

	if helper.IsHas(iCid) {
		_CacheContainer, _error := model.GetCacheContainersByCidAndChecksum(offset, limit, iCid,iChecksum)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheContainer)
	}
	herr.Message = "Can't get to the GetCacheContainersByCidAndChecksum's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheContainersByDataAndExpireHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iData := self.Args("data").Bytes()
	iExpire := self.Args("expire").MustInt()

	if helper.IsHas(iData) {
		_CacheContainer, _error := model.GetCacheContainersByDataAndExpire(offset, limit, iData,iExpire)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheContainer)
	}
	herr.Message = "Can't get to the GetCacheContainersByDataAndExpire's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheContainersByDataAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iData := self.Args("data").Bytes()
	iCreated := self.Args("created").String()

	if helper.IsHas(iData) {
		_CacheContainer, _error := model.GetCacheContainersByDataAndCreated(offset, limit, iData,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheContainer)
	}
	herr.Message = "Can't get to the GetCacheContainersByDataAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheContainersByDataAndSerializedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iData := self.Args("data").Bytes()
	iSerialized := self.Args("serialized").MustInt()

	if helper.IsHas(iData) {
		_CacheContainer, _error := model.GetCacheContainersByDataAndSerialized(offset, limit, iData,iSerialized)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheContainer)
	}
	herr.Message = "Can't get to the GetCacheContainersByDataAndSerialized's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheContainersByDataAndTagsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iData := self.Args("data").Bytes()
	iTags := self.Args("tags").String()

	if helper.IsHas(iData) {
		_CacheContainer, _error := model.GetCacheContainersByDataAndTags(offset, limit, iData,iTags)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheContainer)
	}
	herr.Message = "Can't get to the GetCacheContainersByDataAndTags's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheContainersByDataAndChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iData := self.Args("data").Bytes()
	iChecksum := self.Args("checksum").String()

	if helper.IsHas(iData) {
		_CacheContainer, _error := model.GetCacheContainersByDataAndChecksum(offset, limit, iData,iChecksum)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheContainer)
	}
	herr.Message = "Can't get to the GetCacheContainersByDataAndChecksum's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheContainersByExpireAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iExpire := self.Args("expire").MustInt()
	iCreated := self.Args("created").String()

	if helper.IsHas(iExpire) {
		_CacheContainer, _error := model.GetCacheContainersByExpireAndCreated(offset, limit, iExpire,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheContainer)
	}
	herr.Message = "Can't get to the GetCacheContainersByExpireAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheContainersByExpireAndSerializedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iExpire := self.Args("expire").MustInt()
	iSerialized := self.Args("serialized").MustInt()

	if helper.IsHas(iExpire) {
		_CacheContainer, _error := model.GetCacheContainersByExpireAndSerialized(offset, limit, iExpire,iSerialized)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheContainer)
	}
	herr.Message = "Can't get to the GetCacheContainersByExpireAndSerialized's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheContainersByExpireAndTagsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iExpire := self.Args("expire").MustInt()
	iTags := self.Args("tags").String()

	if helper.IsHas(iExpire) {
		_CacheContainer, _error := model.GetCacheContainersByExpireAndTags(offset, limit, iExpire,iTags)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheContainer)
	}
	herr.Message = "Can't get to the GetCacheContainersByExpireAndTags's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheContainersByExpireAndChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iExpire := self.Args("expire").MustInt()
	iChecksum := self.Args("checksum").String()

	if helper.IsHas(iExpire) {
		_CacheContainer, _error := model.GetCacheContainersByExpireAndChecksum(offset, limit, iExpire,iChecksum)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheContainer)
	}
	herr.Message = "Can't get to the GetCacheContainersByExpireAndChecksum's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheContainersByCreatedAndSerializedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCreated := self.Args("created").String()
	iSerialized := self.Args("serialized").MustInt()

	if helper.IsHas(iCreated) {
		_CacheContainer, _error := model.GetCacheContainersByCreatedAndSerialized(offset, limit, iCreated,iSerialized)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheContainer)
	}
	herr.Message = "Can't get to the GetCacheContainersByCreatedAndSerialized's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheContainersByCreatedAndTagsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCreated := self.Args("created").String()
	iTags := self.Args("tags").String()

	if helper.IsHas(iCreated) {
		_CacheContainer, _error := model.GetCacheContainersByCreatedAndTags(offset, limit, iCreated,iTags)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheContainer)
	}
	herr.Message = "Can't get to the GetCacheContainersByCreatedAndTags's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheContainersByCreatedAndChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCreated := self.Args("created").String()
	iChecksum := self.Args("checksum").String()

	if helper.IsHas(iCreated) {
		_CacheContainer, _error := model.GetCacheContainersByCreatedAndChecksum(offset, limit, iCreated,iChecksum)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheContainer)
	}
	herr.Message = "Can't get to the GetCacheContainersByCreatedAndChecksum's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheContainersBySerializedAndTagsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSerialized := self.Args("serialized").MustInt()
	iTags := self.Args("tags").String()

	if helper.IsHas(iSerialized) {
		_CacheContainer, _error := model.GetCacheContainersBySerializedAndTags(offset, limit, iSerialized,iTags)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheContainer)
	}
	herr.Message = "Can't get to the GetCacheContainersBySerializedAndTags's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheContainersBySerializedAndChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSerialized := self.Args("serialized").MustInt()
	iChecksum := self.Args("checksum").String()

	if helper.IsHas(iSerialized) {
		_CacheContainer, _error := model.GetCacheContainersBySerializedAndChecksum(offset, limit, iSerialized,iChecksum)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheContainer)
	}
	herr.Message = "Can't get to the GetCacheContainersBySerializedAndChecksum's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheContainersByTagsAndChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTags := self.Args("tags").String()
	iChecksum := self.Args("checksum").String()

	if helper.IsHas(iTags) {
		_CacheContainer, _error := model.GetCacheContainersByTagsAndChecksum(offset, limit, iTags,iChecksum)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheContainer)
	}
	herr.Message = "Can't get to the GetCacheContainersByTagsAndChecksum's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheContainersHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	if (offset > 0) && (len(field) > 0) {
		_CacheContainer, _error := model.GetCacheContainers(offset, limit, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheContainer)
	}
	herr.Message = "Can't get to the GetCacheContainers' args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasCacheContainerViaCidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iCid := self.Args("cid").String()
	if helper.IsHas(iCid) {
		_CacheContainer := model.HasCacheContainerViaCid(iCid)
		var m = map[string]interface{}{}
		m["cache_container"] = _CacheContainer
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasCacheContainerViaCid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasCacheContainerViaDataHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iData := self.Args("data").Bytes()
	if helper.IsHas(iData) {
		_CacheContainer := model.HasCacheContainerViaData(iData)
		var m = map[string]interface{}{}
		m["cache_container"] = _CacheContainer
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasCacheContainerViaData's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasCacheContainerViaExpireHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iExpire := self.Args("expire").MustInt()
	if helper.IsHas(iExpire) {
		_CacheContainer := model.HasCacheContainerViaExpire(iExpire)
		var m = map[string]interface{}{}
		m["cache_container"] = _CacheContainer
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasCacheContainerViaExpire's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasCacheContainerViaCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iCreated := self.Args("created").String()
	if helper.IsHas(iCreated) {
		_CacheContainer := model.HasCacheContainerViaCreated(iCreated)
		var m = map[string]interface{}{}
		m["cache_container"] = _CacheContainer
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasCacheContainerViaCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasCacheContainerViaSerializedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iSerialized := self.Args("serialized").MustInt()
	if helper.IsHas(iSerialized) {
		_CacheContainer := model.HasCacheContainerViaSerialized(iSerialized)
		var m = map[string]interface{}{}
		m["cache_container"] = _CacheContainer
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasCacheContainerViaSerialized's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasCacheContainerViaTagsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iTags := self.Args("tags").String()
	if helper.IsHas(iTags) {
		_CacheContainer := model.HasCacheContainerViaTags(iTags)
		var m = map[string]interface{}{}
		m["cache_container"] = _CacheContainer
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasCacheContainerViaTags's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasCacheContainerViaChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iChecksum := self.Args("checksum").String()
	if helper.IsHas(iChecksum) {
		_CacheContainer := model.HasCacheContainerViaChecksum(iChecksum)
		var m = map[string]interface{}{}
		m["cache_container"] = _CacheContainer
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasCacheContainerViaChecksum's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheContainerViaCidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iCid := self.Args("cid").String()
	if helper.IsHas(iCid) {
		_CacheContainer, _error := model.GetCacheContainerViaCid(iCid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheContainer)
	}
	herr.Message = "Can't get to the GetCacheContainerViaCid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheContainerViaDataHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iData := self.Args("data").Bytes()
	if helper.IsHas(iData) {
		_CacheContainer, _error := model.GetCacheContainerViaData(iData)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheContainer)
	}
	herr.Message = "Can't get to the GetCacheContainerViaData's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheContainerViaExpireHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iExpire := self.Args("expire").MustInt()
	if helper.IsHas(iExpire) {
		_CacheContainer, _error := model.GetCacheContainerViaExpire(iExpire)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheContainer)
	}
	herr.Message = "Can't get to the GetCacheContainerViaExpire's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheContainerViaCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iCreated := self.Args("created").String()
	if helper.IsHas(iCreated) {
		_CacheContainer, _error := model.GetCacheContainerViaCreated(iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheContainer)
	}
	herr.Message = "Can't get to the GetCacheContainerViaCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheContainerViaSerializedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iSerialized := self.Args("serialized").MustInt()
	if helper.IsHas(iSerialized) {
		_CacheContainer, _error := model.GetCacheContainerViaSerialized(iSerialized)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheContainer)
	}
	herr.Message = "Can't get to the GetCacheContainerViaSerialized's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheContainerViaTagsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iTags := self.Args("tags").String()
	if helper.IsHas(iTags) {
		_CacheContainer, _error := model.GetCacheContainerViaTags(iTags)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheContainer)
	}
	herr.Message = "Can't get to the GetCacheContainerViaTags's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCacheContainerViaChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iChecksum := self.Args("checksum").String()
	if helper.IsHas(iChecksum) {
		_CacheContainer, _error := model.GetCacheContainerViaChecksum(iChecksum)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheContainer)
	}
	herr.Message = "Can't get to the GetCacheContainerViaChecksum's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetCacheContainerViaCidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Cid_ := self.Args("cid").String()
	if helper.IsHas(Cid_) {
		var iCacheContainer model.CacheContainer
		self.Bind(&iCacheContainer)
		_CacheContainer, _error := model.SetCacheContainerViaCid(Cid_, &iCacheContainer)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheContainer)
	}
	herr.Message = "Can't get to the SetCacheContainerViaCid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetCacheContainerViaDataHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Data_ := self.Args("data").Bytes()
	if helper.IsHas(Data_) {
		var iCacheContainer model.CacheContainer
		self.Bind(&iCacheContainer)
		_CacheContainer, _error := model.SetCacheContainerViaData(Data_, &iCacheContainer)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheContainer)
	}
	herr.Message = "Can't get to the SetCacheContainerViaData's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetCacheContainerViaExpireHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Expire_ := self.Args("expire").MustInt()
	if helper.IsHas(Expire_) {
		var iCacheContainer model.CacheContainer
		self.Bind(&iCacheContainer)
		_CacheContainer, _error := model.SetCacheContainerViaExpire(Expire_, &iCacheContainer)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheContainer)
	}
	herr.Message = "Can't get to the SetCacheContainerViaExpire's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetCacheContainerViaCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Created_ := self.Args("created").String()
	if helper.IsHas(Created_) {
		var iCacheContainer model.CacheContainer
		self.Bind(&iCacheContainer)
		_CacheContainer, _error := model.SetCacheContainerViaCreated(Created_, &iCacheContainer)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheContainer)
	}
	herr.Message = "Can't get to the SetCacheContainerViaCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetCacheContainerViaSerializedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Serialized_ := self.Args("serialized").MustInt()
	if helper.IsHas(Serialized_) {
		var iCacheContainer model.CacheContainer
		self.Bind(&iCacheContainer)
		_CacheContainer, _error := model.SetCacheContainerViaSerialized(Serialized_, &iCacheContainer)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheContainer)
	}
	herr.Message = "Can't get to the SetCacheContainerViaSerialized's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetCacheContainerViaTagsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Tags_ := self.Args("tags").String()
	if helper.IsHas(Tags_) {
		var iCacheContainer model.CacheContainer
		self.Bind(&iCacheContainer)
		_CacheContainer, _error := model.SetCacheContainerViaTags(Tags_, &iCacheContainer)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheContainer)
	}
	herr.Message = "Can't get to the SetCacheContainerViaTags's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetCacheContainerViaChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Checksum_ := self.Args("checksum").String()
	if helper.IsHas(Checksum_) {
		var iCacheContainer model.CacheContainer
		self.Bind(&iCacheContainer)
		_CacheContainer, _error := model.SetCacheContainerViaChecksum(Checksum_, &iCacheContainer)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CacheContainer)
	}
	herr.Message = "Can't get to the SetCacheContainerViaChecksum's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostAddCacheContainerHandler(self *macross.Context) error {
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
		_error := model.AddCacheContainer(Cid_,Data_,Expire_,Created_,Serialized_,Tags_,Checksum_)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		herr.Message = "StatusOK"
		herr.Status = macross.StatusOK
		return self.JSON(herr)
	}
	herr.Message = "Can't get to the AddCacheContainer's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostCacheContainerHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iCacheContainer model.CacheContainer
	self.Bind(&iCacheContainer)
	_string, _error := model.PostCacheContainer(&iCacheContainer)
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

func PutCacheContainerHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iCacheContainer model.CacheContainer
	self.Bind(&iCacheContainer)
	_string, _error := model.PutCacheContainer(&iCacheContainer)
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

func PutCacheContainerViaCidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Cid_ := self.Args("cid").String()
	var iCacheContainer model.CacheContainer
	self.Bind(&iCacheContainer)
	_int64, _error := model.PutCacheContainerViaCid(Cid_, &iCacheContainer)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutCacheContainerViaDataHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Data_ := self.Args("data").Bytes()
	var iCacheContainer model.CacheContainer
	self.Bind(&iCacheContainer)
	_int64, _error := model.PutCacheContainerViaData(Data_, &iCacheContainer)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutCacheContainerViaExpireHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Expire_ := self.Args("expire").MustInt()
	var iCacheContainer model.CacheContainer
	self.Bind(&iCacheContainer)
	_int64, _error := model.PutCacheContainerViaExpire(Expire_, &iCacheContainer)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutCacheContainerViaCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Created_ := self.Args("created").String()
	var iCacheContainer model.CacheContainer
	self.Bind(&iCacheContainer)
	_int64, _error := model.PutCacheContainerViaCreated(Created_, &iCacheContainer)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutCacheContainerViaSerializedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Serialized_ := self.Args("serialized").MustInt()
	var iCacheContainer model.CacheContainer
	self.Bind(&iCacheContainer)
	_int64, _error := model.PutCacheContainerViaSerialized(Serialized_, &iCacheContainer)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutCacheContainerViaTagsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Tags_ := self.Args("tags").String()
	var iCacheContainer model.CacheContainer
	self.Bind(&iCacheContainer)
	_int64, _error := model.PutCacheContainerViaTags(Tags_, &iCacheContainer)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutCacheContainerViaChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Checksum_ := self.Args("checksum").String()
	var iCacheContainer model.CacheContainer
	self.Bind(&iCacheContainer)
	_int64, _error := model.PutCacheContainerViaChecksum(Checksum_, &iCacheContainer)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateCacheContainerViaCidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Cid_ := self.Args("cid").String()
	var iCacheContainer model.CacheContainer
	self.Bind(&iCacheContainer)
	var iMap = helper.StructToMap(iCacheContainer)
	_error := model.UpdateCacheContainerViaCid(Cid_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateCacheContainerViaDataHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Data_ := self.Args("data").Bytes()
	var iCacheContainer model.CacheContainer
	self.Bind(&iCacheContainer)
	var iMap = helper.StructToMap(iCacheContainer)
	_error := model.UpdateCacheContainerViaData(Data_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateCacheContainerViaExpireHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Expire_ := self.Args("expire").MustInt()
	var iCacheContainer model.CacheContainer
	self.Bind(&iCacheContainer)
	var iMap = helper.StructToMap(iCacheContainer)
	_error := model.UpdateCacheContainerViaExpire(Expire_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateCacheContainerViaCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Created_ := self.Args("created").String()
	var iCacheContainer model.CacheContainer
	self.Bind(&iCacheContainer)
	var iMap = helper.StructToMap(iCacheContainer)
	_error := model.UpdateCacheContainerViaCreated(Created_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateCacheContainerViaSerializedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Serialized_ := self.Args("serialized").MustInt()
	var iCacheContainer model.CacheContainer
	self.Bind(&iCacheContainer)
	var iMap = helper.StructToMap(iCacheContainer)
	_error := model.UpdateCacheContainerViaSerialized(Serialized_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateCacheContainerViaTagsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Tags_ := self.Args("tags").String()
	var iCacheContainer model.CacheContainer
	self.Bind(&iCacheContainer)
	var iMap = helper.StructToMap(iCacheContainer)
	_error := model.UpdateCacheContainerViaTags(Tags_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateCacheContainerViaChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Checksum_ := self.Args("checksum").String()
	var iCacheContainer model.CacheContainer
	self.Bind(&iCacheContainer)
	var iMap = helper.StructToMap(iCacheContainer)
	_error := model.UpdateCacheContainerViaChecksum(Checksum_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteCacheContainerHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Cid_ := self.Args("cid").String()
	_int64, _error := model.DeleteCacheContainer(Cid_)
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

func DeleteCacheContainerViaCidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Cid_ := self.Args("cid").String()
	_error := model.DeleteCacheContainerViaCid(Cid_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteCacheContainerViaDataHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Data_ := self.Args("data").Bytes()
	_error := model.DeleteCacheContainerViaData(Data_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteCacheContainerViaExpireHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Expire_ := self.Args("expire").MustInt()
	_error := model.DeleteCacheContainerViaExpire(Expire_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteCacheContainerViaCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Created_ := self.Args("created").String()
	_error := model.DeleteCacheContainerViaCreated(Created_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteCacheContainerViaSerializedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Serialized_ := self.Args("serialized").MustInt()
	_error := model.DeleteCacheContainerViaSerialized(Serialized_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteCacheContainerViaTagsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Tags_ := self.Args("tags").String()
	_error := model.DeleteCacheContainerViaTags(Tags_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteCacheContainerViaChecksumHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Checksum_ := self.Args("checksum").String()
	_error := model.DeleteCacheContainerViaChecksum(Checksum_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}
