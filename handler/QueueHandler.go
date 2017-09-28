package handler

import (
	"github.com/insionng/drupapp/helper"
	"github.com/insionng/drupapp/model"
	"github.com/insionng/macross"
)

func GetQueuesCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()

	if offset > 0 {
		var m = map[string]interface{}{}
		_int64, _error := model.GetQueuesCount(offset, limit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		m["queuesCount"] = _int64
		return self.JSON(m)
	}
	herr.Message = "Can't get to the GetQueuesCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetQueueCountViaItemIdHandler(self *macross.Context) error {
	ItemId_ := self.Args("item_id").MustInt64()
	_int64 := model.GetQueueCountViaItemId(ItemId_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["queueCount"] = 0
	}
	m["queueCount"] = _int64
	return self.JSON(m)
}

func GetQueueCountViaNameHandler(self *macross.Context) error {
	Name_ := self.Args("name").String()
	_int64 := model.GetQueueCountViaName(Name_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["queueCount"] = 0
	}
	m["queueCount"] = _int64
	return self.JSON(m)
}

func GetQueueCountViaDataHandler(self *macross.Context) error {
	Data_ := self.Args("data").Bytes()
	_int64 := model.GetQueueCountViaData(Data_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["queueCount"] = 0
	}
	m["queueCount"] = _int64
	return self.JSON(m)
}

func GetQueueCountViaExpireHandler(self *macross.Context) error {
	Expire_ := self.Args("expire").MustInt()
	_int64 := model.GetQueueCountViaExpire(Expire_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["queueCount"] = 0
	}
	m["queueCount"] = _int64
	return self.JSON(m)
}

func GetQueueCountViaCreatedHandler(self *macross.Context) error {
	Created_ := self.Args("created").MustInt()
	_int64 := model.GetQueueCountViaCreated(Created_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["queueCount"] = 0
	}
	m["queueCount"] = _int64
	return self.JSON(m)
}

func GetQueuesViaItemIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iItemId := self.Args("item_id").MustInt64()
	if (offset > 0) && helper.IsHas(iItemId) {
		_Queue, _error := model.GetQueuesViaItemId(offset, limit, iItemId, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Queue)
	}
	herr.Message = "Can't get to the GetQueuesViaItemId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetQueuesViaNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iName := self.Args("name").String()
	if (offset > 0) && helper.IsHas(iName) {
		_Queue, _error := model.GetQueuesViaName(offset, limit, iName, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Queue)
	}
	herr.Message = "Can't get to the GetQueuesViaName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetQueuesViaDataHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iData := self.Args("data").Bytes()
	if (offset > 0) && helper.IsHas(iData) {
		_Queue, _error := model.GetQueuesViaData(offset, limit, iData, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Queue)
	}
	herr.Message = "Can't get to the GetQueuesViaData's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetQueuesViaExpireHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iExpire := self.Args("expire").MustInt()
	if (offset > 0) && helper.IsHas(iExpire) {
		_Queue, _error := model.GetQueuesViaExpire(offset, limit, iExpire, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Queue)
	}
	herr.Message = "Can't get to the GetQueuesViaExpire's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetQueuesViaCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iCreated := self.Args("created").MustInt()
	if (offset > 0) && helper.IsHas(iCreated) {
		_Queue, _error := model.GetQueuesViaCreated(offset, limit, iCreated, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Queue)
	}
	herr.Message = "Can't get to the GetQueuesViaCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetQueuesByItemIdAndNameAndDataHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iItemId := self.Args("item_id").MustInt64()
	iName := self.Args("name").String()
	iData := self.Args("data").Bytes()

	if helper.IsHas(iItemId) {
		_Queue, _error := model.GetQueuesByItemIdAndNameAndData(offset, limit, iItemId,iName,iData)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Queue)
	}
	herr.Message = "Can't get to the GetQueuesByItemIdAndNameAndData's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetQueuesByItemIdAndNameAndExpireHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iItemId := self.Args("item_id").MustInt64()
	iName := self.Args("name").String()
	iExpire := self.Args("expire").MustInt()

	if helper.IsHas(iItemId) {
		_Queue, _error := model.GetQueuesByItemIdAndNameAndExpire(offset, limit, iItemId,iName,iExpire)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Queue)
	}
	herr.Message = "Can't get to the GetQueuesByItemIdAndNameAndExpire's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetQueuesByItemIdAndNameAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iItemId := self.Args("item_id").MustInt64()
	iName := self.Args("name").String()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iItemId) {
		_Queue, _error := model.GetQueuesByItemIdAndNameAndCreated(offset, limit, iItemId,iName,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Queue)
	}
	herr.Message = "Can't get to the GetQueuesByItemIdAndNameAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetQueuesByItemIdAndDataAndExpireHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iItemId := self.Args("item_id").MustInt64()
	iData := self.Args("data").Bytes()
	iExpire := self.Args("expire").MustInt()

	if helper.IsHas(iItemId) {
		_Queue, _error := model.GetQueuesByItemIdAndDataAndExpire(offset, limit, iItemId,iData,iExpire)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Queue)
	}
	herr.Message = "Can't get to the GetQueuesByItemIdAndDataAndExpire's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetQueuesByItemIdAndDataAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iItemId := self.Args("item_id").MustInt64()
	iData := self.Args("data").Bytes()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iItemId) {
		_Queue, _error := model.GetQueuesByItemIdAndDataAndCreated(offset, limit, iItemId,iData,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Queue)
	}
	herr.Message = "Can't get to the GetQueuesByItemIdAndDataAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetQueuesByItemIdAndExpireAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iItemId := self.Args("item_id").MustInt64()
	iExpire := self.Args("expire").MustInt()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iItemId) {
		_Queue, _error := model.GetQueuesByItemIdAndExpireAndCreated(offset, limit, iItemId,iExpire,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Queue)
	}
	herr.Message = "Can't get to the GetQueuesByItemIdAndExpireAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetQueuesByNameAndDataAndExpireHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iData := self.Args("data").Bytes()
	iExpire := self.Args("expire").MustInt()

	if helper.IsHas(iName) {
		_Queue, _error := model.GetQueuesByNameAndDataAndExpire(offset, limit, iName,iData,iExpire)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Queue)
	}
	herr.Message = "Can't get to the GetQueuesByNameAndDataAndExpire's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetQueuesByNameAndDataAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iData := self.Args("data").Bytes()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iName) {
		_Queue, _error := model.GetQueuesByNameAndDataAndCreated(offset, limit, iName,iData,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Queue)
	}
	herr.Message = "Can't get to the GetQueuesByNameAndDataAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetQueuesByNameAndExpireAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iExpire := self.Args("expire").MustInt()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iName) {
		_Queue, _error := model.GetQueuesByNameAndExpireAndCreated(offset, limit, iName,iExpire,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Queue)
	}
	herr.Message = "Can't get to the GetQueuesByNameAndExpireAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetQueuesByDataAndExpireAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iData := self.Args("data").Bytes()
	iExpire := self.Args("expire").MustInt()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iData) {
		_Queue, _error := model.GetQueuesByDataAndExpireAndCreated(offset, limit, iData,iExpire,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Queue)
	}
	herr.Message = "Can't get to the GetQueuesByDataAndExpireAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetQueuesByItemIdAndNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iItemId := self.Args("item_id").MustInt64()
	iName := self.Args("name").String()

	if helper.IsHas(iItemId) {
		_Queue, _error := model.GetQueuesByItemIdAndName(offset, limit, iItemId,iName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Queue)
	}
	herr.Message = "Can't get to the GetQueuesByItemIdAndName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetQueuesByItemIdAndDataHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iItemId := self.Args("item_id").MustInt64()
	iData := self.Args("data").Bytes()

	if helper.IsHas(iItemId) {
		_Queue, _error := model.GetQueuesByItemIdAndData(offset, limit, iItemId,iData)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Queue)
	}
	herr.Message = "Can't get to the GetQueuesByItemIdAndData's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetQueuesByItemIdAndExpireHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iItemId := self.Args("item_id").MustInt64()
	iExpire := self.Args("expire").MustInt()

	if helper.IsHas(iItemId) {
		_Queue, _error := model.GetQueuesByItemIdAndExpire(offset, limit, iItemId,iExpire)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Queue)
	}
	herr.Message = "Can't get to the GetQueuesByItemIdAndExpire's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetQueuesByItemIdAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iItemId := self.Args("item_id").MustInt64()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iItemId) {
		_Queue, _error := model.GetQueuesByItemIdAndCreated(offset, limit, iItemId,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Queue)
	}
	herr.Message = "Can't get to the GetQueuesByItemIdAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetQueuesByNameAndDataHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iData := self.Args("data").Bytes()

	if helper.IsHas(iName) {
		_Queue, _error := model.GetQueuesByNameAndData(offset, limit, iName,iData)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Queue)
	}
	herr.Message = "Can't get to the GetQueuesByNameAndData's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetQueuesByNameAndExpireHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iExpire := self.Args("expire").MustInt()

	if helper.IsHas(iName) {
		_Queue, _error := model.GetQueuesByNameAndExpire(offset, limit, iName,iExpire)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Queue)
	}
	herr.Message = "Can't get to the GetQueuesByNameAndExpire's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetQueuesByNameAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iName) {
		_Queue, _error := model.GetQueuesByNameAndCreated(offset, limit, iName,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Queue)
	}
	herr.Message = "Can't get to the GetQueuesByNameAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetQueuesByDataAndExpireHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iData := self.Args("data").Bytes()
	iExpire := self.Args("expire").MustInt()

	if helper.IsHas(iData) {
		_Queue, _error := model.GetQueuesByDataAndExpire(offset, limit, iData,iExpire)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Queue)
	}
	herr.Message = "Can't get to the GetQueuesByDataAndExpire's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetQueuesByDataAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iData := self.Args("data").Bytes()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iData) {
		_Queue, _error := model.GetQueuesByDataAndCreated(offset, limit, iData,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Queue)
	}
	herr.Message = "Can't get to the GetQueuesByDataAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetQueuesByExpireAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iExpire := self.Args("expire").MustInt()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iExpire) {
		_Queue, _error := model.GetQueuesByExpireAndCreated(offset, limit, iExpire,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Queue)
	}
	herr.Message = "Can't get to the GetQueuesByExpireAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetQueuesHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	if (offset > 0) && (len(field) > 0) {
		_Queue, _error := model.GetQueues(offset, limit, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Queue)
	}
	herr.Message = "Can't get to the GetQueues' args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasQueueViaItemIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iItemId := self.Args("item_id").MustInt64()
	if helper.IsHas(iItemId) {
		_Queue := model.HasQueueViaItemId(iItemId)
		var m = map[string]interface{}{}
		m["queue"] = _Queue
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasQueueViaItemId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasQueueViaNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iName := self.Args("name").String()
	if helper.IsHas(iName) {
		_Queue := model.HasQueueViaName(iName)
		var m = map[string]interface{}{}
		m["queue"] = _Queue
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasQueueViaName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasQueueViaDataHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iData := self.Args("data").Bytes()
	if helper.IsHas(iData) {
		_Queue := model.HasQueueViaData(iData)
		var m = map[string]interface{}{}
		m["queue"] = _Queue
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasQueueViaData's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasQueueViaExpireHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iExpire := self.Args("expire").MustInt()
	if helper.IsHas(iExpire) {
		_Queue := model.HasQueueViaExpire(iExpire)
		var m = map[string]interface{}{}
		m["queue"] = _Queue
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasQueueViaExpire's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasQueueViaCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iCreated := self.Args("created").MustInt()
	if helper.IsHas(iCreated) {
		_Queue := model.HasQueueViaCreated(iCreated)
		var m = map[string]interface{}{}
		m["queue"] = _Queue
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasQueueViaCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetQueueViaItemIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iItemId := self.Args("item_id").MustInt64()
	if helper.IsHas(iItemId) {
		_Queue, _error := model.GetQueueViaItemId(iItemId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Queue)
	}
	herr.Message = "Can't get to the GetQueueViaItemId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetQueueViaNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iName := self.Args("name").String()
	if helper.IsHas(iName) {
		_Queue, _error := model.GetQueueViaName(iName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Queue)
	}
	herr.Message = "Can't get to the GetQueueViaName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetQueueViaDataHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iData := self.Args("data").Bytes()
	if helper.IsHas(iData) {
		_Queue, _error := model.GetQueueViaData(iData)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Queue)
	}
	herr.Message = "Can't get to the GetQueueViaData's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetQueueViaExpireHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iExpire := self.Args("expire").MustInt()
	if helper.IsHas(iExpire) {
		_Queue, _error := model.GetQueueViaExpire(iExpire)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Queue)
	}
	herr.Message = "Can't get to the GetQueueViaExpire's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetQueueViaCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iCreated := self.Args("created").MustInt()
	if helper.IsHas(iCreated) {
		_Queue, _error := model.GetQueueViaCreated(iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Queue)
	}
	herr.Message = "Can't get to the GetQueueViaCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetQueueViaItemIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	ItemId_ := self.Args("item_id").MustInt64()
	if helper.IsHas(ItemId_) {
		var iQueue model.Queue
		self.Bind(&iQueue)
		_Queue, _error := model.SetQueueViaItemId(ItemId_, &iQueue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Queue)
	}
	herr.Message = "Can't get to the SetQueueViaItemId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetQueueViaNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Name_ := self.Args("name").String()
	if helper.IsHas(Name_) {
		var iQueue model.Queue
		self.Bind(&iQueue)
		_Queue, _error := model.SetQueueViaName(Name_, &iQueue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Queue)
	}
	herr.Message = "Can't get to the SetQueueViaName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetQueueViaDataHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Data_ := self.Args("data").Bytes()
	if helper.IsHas(Data_) {
		var iQueue model.Queue
		self.Bind(&iQueue)
		_Queue, _error := model.SetQueueViaData(Data_, &iQueue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Queue)
	}
	herr.Message = "Can't get to the SetQueueViaData's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetQueueViaExpireHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Expire_ := self.Args("expire").MustInt()
	if helper.IsHas(Expire_) {
		var iQueue model.Queue
		self.Bind(&iQueue)
		_Queue, _error := model.SetQueueViaExpire(Expire_, &iQueue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Queue)
	}
	herr.Message = "Can't get to the SetQueueViaExpire's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetQueueViaCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Created_ := self.Args("created").MustInt()
	if helper.IsHas(Created_) {
		var iQueue model.Queue
		self.Bind(&iQueue)
		_Queue, _error := model.SetQueueViaCreated(Created_, &iQueue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Queue)
	}
	herr.Message = "Can't get to the SetQueueViaCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostAddQueueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	ItemId_ := self.Args("item_id").MustInt64()
	Name_ := self.Args("name").String()
	Data_ := self.Args("data").Bytes()
	Expire_ := self.Args("expire").MustInt()
	Created_ := self.Args("created").MustInt()

	if helper.IsHas(ItemId_) {
		_error := model.AddQueue(ItemId_,Name_,Data_,Expire_,Created_)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		herr.Message = "StatusOK"
		herr.Status = macross.StatusOK
		return self.JSON(herr)
	}
	herr.Message = "Can't get to the AddQueue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostQueueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iQueue model.Queue
	self.Bind(&iQueue)
	_int64, _error := model.PostQueue(&iQueue)
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

func PutQueueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iQueue model.Queue
	self.Bind(&iQueue)
	_int64, _error := model.PutQueue(&iQueue)
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

func PutQueueViaItemIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	ItemId_ := self.Args("item_id").MustInt64()
	var iQueue model.Queue
	self.Bind(&iQueue)
	_int64, _error := model.PutQueueViaItemId(ItemId_, &iQueue)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutQueueViaNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Name_ := self.Args("name").String()
	var iQueue model.Queue
	self.Bind(&iQueue)
	_int64, _error := model.PutQueueViaName(Name_, &iQueue)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutQueueViaDataHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Data_ := self.Args("data").Bytes()
	var iQueue model.Queue
	self.Bind(&iQueue)
	_int64, _error := model.PutQueueViaData(Data_, &iQueue)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutQueueViaExpireHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Expire_ := self.Args("expire").MustInt()
	var iQueue model.Queue
	self.Bind(&iQueue)
	_int64, _error := model.PutQueueViaExpire(Expire_, &iQueue)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutQueueViaCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Created_ := self.Args("created").MustInt()
	var iQueue model.Queue
	self.Bind(&iQueue)
	_int64, _error := model.PutQueueViaCreated(Created_, &iQueue)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateQueueViaItemIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	ItemId_ := self.Args("item_id").MustInt64()
	var iQueue model.Queue
	self.Bind(&iQueue)
	var iMap = helper.StructToMap(iQueue)
	_error := model.UpdateQueueViaItemId(ItemId_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateQueueViaNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Name_ := self.Args("name").String()
	var iQueue model.Queue
	self.Bind(&iQueue)
	var iMap = helper.StructToMap(iQueue)
	_error := model.UpdateQueueViaName(Name_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateQueueViaDataHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Data_ := self.Args("data").Bytes()
	var iQueue model.Queue
	self.Bind(&iQueue)
	var iMap = helper.StructToMap(iQueue)
	_error := model.UpdateQueueViaData(Data_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateQueueViaExpireHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Expire_ := self.Args("expire").MustInt()
	var iQueue model.Queue
	self.Bind(&iQueue)
	var iMap = helper.StructToMap(iQueue)
	_error := model.UpdateQueueViaExpire(Expire_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateQueueViaCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Created_ := self.Args("created").MustInt()
	var iQueue model.Queue
	self.Bind(&iQueue)
	var iMap = helper.StructToMap(iQueue)
	_error := model.UpdateQueueViaCreated(Created_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteQueueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	ItemId_ := self.Args("item_id").MustInt64()
	_int64, _error := model.DeleteQueue(ItemId_)
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

func DeleteQueueViaItemIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	ItemId_ := self.Args("item_id").MustInt64()
	_error := model.DeleteQueueViaItemId(ItemId_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteQueueViaNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Name_ := self.Args("name").String()
	_error := model.DeleteQueueViaName(Name_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteQueueViaDataHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Data_ := self.Args("data").Bytes()
	_error := model.DeleteQueueViaData(Data_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteQueueViaExpireHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Expire_ := self.Args("expire").MustInt()
	_error := model.DeleteQueueViaExpire(Expire_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteQueueViaCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Created_ := self.Args("created").MustInt()
	_error := model.DeleteQueueViaCreated(Created_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}
