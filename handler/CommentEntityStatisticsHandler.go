package handler

import (
	"drupapp/helper"
	"drupapp/model"
	"github.com/insionng/macross"
)

func GetCommentEntityStatisticsesCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()

	if offset > 0 {
		var m = map[string]interface{}{}
		_int64, _error := model.GetCommentEntityStatisticsesCount(offset, limit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		m["comment_entity_statisticssCount"] = _int64
		return self.JSON(m)
	}
	herr.Message = "Can't get to the GetCommentEntityStatisticsesCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentEntityStatisticsCountViaEntityIdHandler(self *macross.Context) error {
	EntityId_ := self.Args("entity_id").MustInt64()
	_int64 := model.GetCommentEntityStatisticsCountViaEntityId(EntityId_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["comment_entity_statisticsCount"] = 0
	}
	m["comment_entity_statisticsCount"] = _int64
	return self.JSON(m)
}

func GetCommentEntityStatisticsCountViaEntityTypeHandler(self *macross.Context) error {
	EntityType_ := self.Args("entity_type").String()
	_int64 := model.GetCommentEntityStatisticsCountViaEntityType(EntityType_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["comment_entity_statisticsCount"] = 0
	}
	m["comment_entity_statisticsCount"] = _int64
	return self.JSON(m)
}

func GetCommentEntityStatisticsCountViaFieldNameHandler(self *macross.Context) error {
	FieldName_ := self.Args("field_name").String()
	_int64 := model.GetCommentEntityStatisticsCountViaFieldName(FieldName_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["comment_entity_statisticsCount"] = 0
	}
	m["comment_entity_statisticsCount"] = _int64
	return self.JSON(m)
}

func GetCommentEntityStatisticsCountViaCidHandler(self *macross.Context) error {
	Cid_ := self.Args("cid").MustInt()
	_int64 := model.GetCommentEntityStatisticsCountViaCid(Cid_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["comment_entity_statisticsCount"] = 0
	}
	m["comment_entity_statisticsCount"] = _int64
	return self.JSON(m)
}

func GetCommentEntityStatisticsCountViaLastCommentTimestampHandler(self *macross.Context) error {
	LastCommentTimestamp_ := self.Args("last_comment_timestamp").MustInt()
	_int64 := model.GetCommentEntityStatisticsCountViaLastCommentTimestamp(LastCommentTimestamp_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["comment_entity_statisticsCount"] = 0
	}
	m["comment_entity_statisticsCount"] = _int64
	return self.JSON(m)
}

func GetCommentEntityStatisticsCountViaLastCommentNameHandler(self *macross.Context) error {
	LastCommentName_ := self.Args("last_comment_name").String()
	_int64 := model.GetCommentEntityStatisticsCountViaLastCommentName(LastCommentName_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["comment_entity_statisticsCount"] = 0
	}
	m["comment_entity_statisticsCount"] = _int64
	return self.JSON(m)
}

func GetCommentEntityStatisticsCountViaLastCommentUidHandler(self *macross.Context) error {
	LastCommentUid_ := self.Args("last_comment_uid").MustInt()
	_int64 := model.GetCommentEntityStatisticsCountViaLastCommentUid(LastCommentUid_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["comment_entity_statisticsCount"] = 0
	}
	m["comment_entity_statisticsCount"] = _int64
	return self.JSON(m)
}

func GetCommentEntityStatisticsCountViaCommentCountHandler(self *macross.Context) error {
	CommentCount_ := self.Args("comment_count").MustInt()
	_int64 := model.GetCommentEntityStatisticsCountViaCommentCount(CommentCount_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["comment_entity_statisticsCount"] = 0
	}
	m["comment_entity_statisticsCount"] = _int64
	return self.JSON(m)
}

func GetCommentEntityStatisticsesViaEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iEntityId := self.Args("entity_id").MustInt64()
	if (offset > 0) && helper.IsHas(iEntityId) {
		_CommentEntityStatistics, _error := model.GetCommentEntityStatisticsesViaEntityId(offset, limit, iEntityId, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentEntityStatistics)
	}
	herr.Message = "Can't get to the GetCommentEntityStatisticsesViaEntityId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentEntityStatisticsesViaEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iEntityType := self.Args("entity_type").String()
	if (offset > 0) && helper.IsHas(iEntityType) {
		_CommentEntityStatistics, _error := model.GetCommentEntityStatisticsesViaEntityType(offset, limit, iEntityType, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentEntityStatistics)
	}
	herr.Message = "Can't get to the GetCommentEntityStatisticsesViaEntityType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentEntityStatisticsesViaFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iFieldName := self.Args("field_name").String()
	if (offset > 0) && helper.IsHas(iFieldName) {
		_CommentEntityStatistics, _error := model.GetCommentEntityStatisticsesViaFieldName(offset, limit, iFieldName, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentEntityStatistics)
	}
	herr.Message = "Can't get to the GetCommentEntityStatisticsesViaFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentEntityStatisticsesViaCidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iCid := self.Args("cid").MustInt()
	if (offset > 0) && helper.IsHas(iCid) {
		_CommentEntityStatistics, _error := model.GetCommentEntityStatisticsesViaCid(offset, limit, iCid, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentEntityStatistics)
	}
	herr.Message = "Can't get to the GetCommentEntityStatisticsesViaCid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentEntityStatisticsesViaLastCommentTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iLastCommentTimestamp := self.Args("last_comment_timestamp").MustInt()
	if (offset > 0) && helper.IsHas(iLastCommentTimestamp) {
		_CommentEntityStatistics, _error := model.GetCommentEntityStatisticsesViaLastCommentTimestamp(offset, limit, iLastCommentTimestamp, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentEntityStatistics)
	}
	herr.Message = "Can't get to the GetCommentEntityStatisticsesViaLastCommentTimestamp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentEntityStatisticsesViaLastCommentNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iLastCommentName := self.Args("last_comment_name").String()
	if (offset > 0) && helper.IsHas(iLastCommentName) {
		_CommentEntityStatistics, _error := model.GetCommentEntityStatisticsesViaLastCommentName(offset, limit, iLastCommentName, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentEntityStatistics)
	}
	herr.Message = "Can't get to the GetCommentEntityStatisticsesViaLastCommentName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentEntityStatisticsesViaLastCommentUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iLastCommentUid := self.Args("last_comment_uid").MustInt()
	if (offset > 0) && helper.IsHas(iLastCommentUid) {
		_CommentEntityStatistics, _error := model.GetCommentEntityStatisticsesViaLastCommentUid(offset, limit, iLastCommentUid, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentEntityStatistics)
	}
	herr.Message = "Can't get to the GetCommentEntityStatisticsesViaLastCommentUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentEntityStatisticsesViaCommentCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iCommentCount := self.Args("comment_count").MustInt()
	if (offset > 0) && helper.IsHas(iCommentCount) {
		_CommentEntityStatistics, _error := model.GetCommentEntityStatisticsesViaCommentCount(offset, limit, iCommentCount, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentEntityStatistics)
	}
	herr.Message = "Can't get to the GetCommentEntityStatisticsesViaCommentCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentEntityStatisticsesByEntityIdAndEntityTypeAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt64()
	iEntityType := self.Args("entity_type").String()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iEntityId) {
		_CommentEntityStatistics, _error := model.GetCommentEntityStatisticsesByEntityIdAndEntityTypeAndFieldName(offset, limit, iEntityId,iEntityType,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentEntityStatistics)
	}
	herr.Message = "Can't get to the GetCommentEntityStatisticsesByEntityIdAndEntityTypeAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentEntityStatisticsesByEntityIdAndEntityTypeAndCidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt64()
	iEntityType := self.Args("entity_type").String()
	iCid := self.Args("cid").MustInt()

	if helper.IsHas(iEntityId) {
		_CommentEntityStatistics, _error := model.GetCommentEntityStatisticsesByEntityIdAndEntityTypeAndCid(offset, limit, iEntityId,iEntityType,iCid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentEntityStatistics)
	}
	herr.Message = "Can't get to the GetCommentEntityStatisticsesByEntityIdAndEntityTypeAndCid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentEntityStatisticsesByEntityIdAndEntityTypeAndLastCommentTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt64()
	iEntityType := self.Args("entity_type").String()
	iLastCommentTimestamp := self.Args("last_comment_timestamp").MustInt()

	if helper.IsHas(iEntityId) {
		_CommentEntityStatistics, _error := model.GetCommentEntityStatisticsesByEntityIdAndEntityTypeAndLastCommentTimestamp(offset, limit, iEntityId,iEntityType,iLastCommentTimestamp)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentEntityStatistics)
	}
	herr.Message = "Can't get to the GetCommentEntityStatisticsesByEntityIdAndEntityTypeAndLastCommentTimestamp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentEntityStatisticsesByEntityIdAndEntityTypeAndLastCommentNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt64()
	iEntityType := self.Args("entity_type").String()
	iLastCommentName := self.Args("last_comment_name").String()

	if helper.IsHas(iEntityId) {
		_CommentEntityStatistics, _error := model.GetCommentEntityStatisticsesByEntityIdAndEntityTypeAndLastCommentName(offset, limit, iEntityId,iEntityType,iLastCommentName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentEntityStatistics)
	}
	herr.Message = "Can't get to the GetCommentEntityStatisticsesByEntityIdAndEntityTypeAndLastCommentName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentEntityStatisticsesByEntityIdAndEntityTypeAndLastCommentUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt64()
	iEntityType := self.Args("entity_type").String()
	iLastCommentUid := self.Args("last_comment_uid").MustInt()

	if helper.IsHas(iEntityId) {
		_CommentEntityStatistics, _error := model.GetCommentEntityStatisticsesByEntityIdAndEntityTypeAndLastCommentUid(offset, limit, iEntityId,iEntityType,iLastCommentUid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentEntityStatistics)
	}
	herr.Message = "Can't get to the GetCommentEntityStatisticsesByEntityIdAndEntityTypeAndLastCommentUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentEntityStatisticsesByEntityIdAndEntityTypeAndCommentCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt64()
	iEntityType := self.Args("entity_type").String()
	iCommentCount := self.Args("comment_count").MustInt()

	if helper.IsHas(iEntityId) {
		_CommentEntityStatistics, _error := model.GetCommentEntityStatisticsesByEntityIdAndEntityTypeAndCommentCount(offset, limit, iEntityId,iEntityType,iCommentCount)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentEntityStatistics)
	}
	herr.Message = "Can't get to the GetCommentEntityStatisticsesByEntityIdAndEntityTypeAndCommentCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentEntityStatisticsesByEntityIdAndFieldNameAndCidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt64()
	iFieldName := self.Args("field_name").String()
	iCid := self.Args("cid").MustInt()

	if helper.IsHas(iEntityId) {
		_CommentEntityStatistics, _error := model.GetCommentEntityStatisticsesByEntityIdAndFieldNameAndCid(offset, limit, iEntityId,iFieldName,iCid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentEntityStatistics)
	}
	herr.Message = "Can't get to the GetCommentEntityStatisticsesByEntityIdAndFieldNameAndCid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentEntityStatisticsesByEntityIdAndFieldNameAndLastCommentTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt64()
	iFieldName := self.Args("field_name").String()
	iLastCommentTimestamp := self.Args("last_comment_timestamp").MustInt()

	if helper.IsHas(iEntityId) {
		_CommentEntityStatistics, _error := model.GetCommentEntityStatisticsesByEntityIdAndFieldNameAndLastCommentTimestamp(offset, limit, iEntityId,iFieldName,iLastCommentTimestamp)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentEntityStatistics)
	}
	herr.Message = "Can't get to the GetCommentEntityStatisticsesByEntityIdAndFieldNameAndLastCommentTimestamp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentEntityStatisticsesByEntityIdAndFieldNameAndLastCommentNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt64()
	iFieldName := self.Args("field_name").String()
	iLastCommentName := self.Args("last_comment_name").String()

	if helper.IsHas(iEntityId) {
		_CommentEntityStatistics, _error := model.GetCommentEntityStatisticsesByEntityIdAndFieldNameAndLastCommentName(offset, limit, iEntityId,iFieldName,iLastCommentName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentEntityStatistics)
	}
	herr.Message = "Can't get to the GetCommentEntityStatisticsesByEntityIdAndFieldNameAndLastCommentName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentEntityStatisticsesByEntityIdAndFieldNameAndLastCommentUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt64()
	iFieldName := self.Args("field_name").String()
	iLastCommentUid := self.Args("last_comment_uid").MustInt()

	if helper.IsHas(iEntityId) {
		_CommentEntityStatistics, _error := model.GetCommentEntityStatisticsesByEntityIdAndFieldNameAndLastCommentUid(offset, limit, iEntityId,iFieldName,iLastCommentUid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentEntityStatistics)
	}
	herr.Message = "Can't get to the GetCommentEntityStatisticsesByEntityIdAndFieldNameAndLastCommentUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentEntityStatisticsesByEntityIdAndFieldNameAndCommentCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt64()
	iFieldName := self.Args("field_name").String()
	iCommentCount := self.Args("comment_count").MustInt()

	if helper.IsHas(iEntityId) {
		_CommentEntityStatistics, _error := model.GetCommentEntityStatisticsesByEntityIdAndFieldNameAndCommentCount(offset, limit, iEntityId,iFieldName,iCommentCount)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentEntityStatistics)
	}
	herr.Message = "Can't get to the GetCommentEntityStatisticsesByEntityIdAndFieldNameAndCommentCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentEntityStatisticsesByEntityIdAndCidAndLastCommentTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt64()
	iCid := self.Args("cid").MustInt()
	iLastCommentTimestamp := self.Args("last_comment_timestamp").MustInt()

	if helper.IsHas(iEntityId) {
		_CommentEntityStatistics, _error := model.GetCommentEntityStatisticsesByEntityIdAndCidAndLastCommentTimestamp(offset, limit, iEntityId,iCid,iLastCommentTimestamp)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentEntityStatistics)
	}
	herr.Message = "Can't get to the GetCommentEntityStatisticsesByEntityIdAndCidAndLastCommentTimestamp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentEntityStatisticsesByEntityIdAndCidAndLastCommentNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt64()
	iCid := self.Args("cid").MustInt()
	iLastCommentName := self.Args("last_comment_name").String()

	if helper.IsHas(iEntityId) {
		_CommentEntityStatistics, _error := model.GetCommentEntityStatisticsesByEntityIdAndCidAndLastCommentName(offset, limit, iEntityId,iCid,iLastCommentName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentEntityStatistics)
	}
	herr.Message = "Can't get to the GetCommentEntityStatisticsesByEntityIdAndCidAndLastCommentName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentEntityStatisticsesByEntityIdAndCidAndLastCommentUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt64()
	iCid := self.Args("cid").MustInt()
	iLastCommentUid := self.Args("last_comment_uid").MustInt()

	if helper.IsHas(iEntityId) {
		_CommentEntityStatistics, _error := model.GetCommentEntityStatisticsesByEntityIdAndCidAndLastCommentUid(offset, limit, iEntityId,iCid,iLastCommentUid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentEntityStatistics)
	}
	herr.Message = "Can't get to the GetCommentEntityStatisticsesByEntityIdAndCidAndLastCommentUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentEntityStatisticsesByEntityIdAndCidAndCommentCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt64()
	iCid := self.Args("cid").MustInt()
	iCommentCount := self.Args("comment_count").MustInt()

	if helper.IsHas(iEntityId) {
		_CommentEntityStatistics, _error := model.GetCommentEntityStatisticsesByEntityIdAndCidAndCommentCount(offset, limit, iEntityId,iCid,iCommentCount)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentEntityStatistics)
	}
	herr.Message = "Can't get to the GetCommentEntityStatisticsesByEntityIdAndCidAndCommentCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentEntityStatisticsesByEntityIdAndLastCommentTimestampAndLastCommentNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt64()
	iLastCommentTimestamp := self.Args("last_comment_timestamp").MustInt()
	iLastCommentName := self.Args("last_comment_name").String()

	if helper.IsHas(iEntityId) {
		_CommentEntityStatistics, _error := model.GetCommentEntityStatisticsesByEntityIdAndLastCommentTimestampAndLastCommentName(offset, limit, iEntityId,iLastCommentTimestamp,iLastCommentName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentEntityStatistics)
	}
	herr.Message = "Can't get to the GetCommentEntityStatisticsesByEntityIdAndLastCommentTimestampAndLastCommentName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentEntityStatisticsesByEntityIdAndLastCommentTimestampAndLastCommentUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt64()
	iLastCommentTimestamp := self.Args("last_comment_timestamp").MustInt()
	iLastCommentUid := self.Args("last_comment_uid").MustInt()

	if helper.IsHas(iEntityId) {
		_CommentEntityStatistics, _error := model.GetCommentEntityStatisticsesByEntityIdAndLastCommentTimestampAndLastCommentUid(offset, limit, iEntityId,iLastCommentTimestamp,iLastCommentUid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentEntityStatistics)
	}
	herr.Message = "Can't get to the GetCommentEntityStatisticsesByEntityIdAndLastCommentTimestampAndLastCommentUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentEntityStatisticsesByEntityIdAndLastCommentTimestampAndCommentCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt64()
	iLastCommentTimestamp := self.Args("last_comment_timestamp").MustInt()
	iCommentCount := self.Args("comment_count").MustInt()

	if helper.IsHas(iEntityId) {
		_CommentEntityStatistics, _error := model.GetCommentEntityStatisticsesByEntityIdAndLastCommentTimestampAndCommentCount(offset, limit, iEntityId,iLastCommentTimestamp,iCommentCount)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentEntityStatistics)
	}
	herr.Message = "Can't get to the GetCommentEntityStatisticsesByEntityIdAndLastCommentTimestampAndCommentCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentEntityStatisticsesByEntityIdAndLastCommentNameAndLastCommentUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt64()
	iLastCommentName := self.Args("last_comment_name").String()
	iLastCommentUid := self.Args("last_comment_uid").MustInt()

	if helper.IsHas(iEntityId) {
		_CommentEntityStatistics, _error := model.GetCommentEntityStatisticsesByEntityIdAndLastCommentNameAndLastCommentUid(offset, limit, iEntityId,iLastCommentName,iLastCommentUid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentEntityStatistics)
	}
	herr.Message = "Can't get to the GetCommentEntityStatisticsesByEntityIdAndLastCommentNameAndLastCommentUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentEntityStatisticsesByEntityIdAndLastCommentNameAndCommentCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt64()
	iLastCommentName := self.Args("last_comment_name").String()
	iCommentCount := self.Args("comment_count").MustInt()

	if helper.IsHas(iEntityId) {
		_CommentEntityStatistics, _error := model.GetCommentEntityStatisticsesByEntityIdAndLastCommentNameAndCommentCount(offset, limit, iEntityId,iLastCommentName,iCommentCount)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentEntityStatistics)
	}
	herr.Message = "Can't get to the GetCommentEntityStatisticsesByEntityIdAndLastCommentNameAndCommentCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentEntityStatisticsesByEntityIdAndLastCommentUidAndCommentCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt64()
	iLastCommentUid := self.Args("last_comment_uid").MustInt()
	iCommentCount := self.Args("comment_count").MustInt()

	if helper.IsHas(iEntityId) {
		_CommentEntityStatistics, _error := model.GetCommentEntityStatisticsesByEntityIdAndLastCommentUidAndCommentCount(offset, limit, iEntityId,iLastCommentUid,iCommentCount)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentEntityStatistics)
	}
	herr.Message = "Can't get to the GetCommentEntityStatisticsesByEntityIdAndLastCommentUidAndCommentCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentEntityStatisticsesByEntityTypeAndFieldNameAndCidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityType := self.Args("entity_type").String()
	iFieldName := self.Args("field_name").String()
	iCid := self.Args("cid").MustInt()

	if helper.IsHas(iEntityType) {
		_CommentEntityStatistics, _error := model.GetCommentEntityStatisticsesByEntityTypeAndFieldNameAndCid(offset, limit, iEntityType,iFieldName,iCid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentEntityStatistics)
	}
	herr.Message = "Can't get to the GetCommentEntityStatisticsesByEntityTypeAndFieldNameAndCid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentEntityStatisticsesByEntityTypeAndFieldNameAndLastCommentTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityType := self.Args("entity_type").String()
	iFieldName := self.Args("field_name").String()
	iLastCommentTimestamp := self.Args("last_comment_timestamp").MustInt()

	if helper.IsHas(iEntityType) {
		_CommentEntityStatistics, _error := model.GetCommentEntityStatisticsesByEntityTypeAndFieldNameAndLastCommentTimestamp(offset, limit, iEntityType,iFieldName,iLastCommentTimestamp)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentEntityStatistics)
	}
	herr.Message = "Can't get to the GetCommentEntityStatisticsesByEntityTypeAndFieldNameAndLastCommentTimestamp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentEntityStatisticsesByEntityTypeAndFieldNameAndLastCommentNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityType := self.Args("entity_type").String()
	iFieldName := self.Args("field_name").String()
	iLastCommentName := self.Args("last_comment_name").String()

	if helper.IsHas(iEntityType) {
		_CommentEntityStatistics, _error := model.GetCommentEntityStatisticsesByEntityTypeAndFieldNameAndLastCommentName(offset, limit, iEntityType,iFieldName,iLastCommentName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentEntityStatistics)
	}
	herr.Message = "Can't get to the GetCommentEntityStatisticsesByEntityTypeAndFieldNameAndLastCommentName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentEntityStatisticsesByEntityTypeAndFieldNameAndLastCommentUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityType := self.Args("entity_type").String()
	iFieldName := self.Args("field_name").String()
	iLastCommentUid := self.Args("last_comment_uid").MustInt()

	if helper.IsHas(iEntityType) {
		_CommentEntityStatistics, _error := model.GetCommentEntityStatisticsesByEntityTypeAndFieldNameAndLastCommentUid(offset, limit, iEntityType,iFieldName,iLastCommentUid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentEntityStatistics)
	}
	herr.Message = "Can't get to the GetCommentEntityStatisticsesByEntityTypeAndFieldNameAndLastCommentUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentEntityStatisticsesByEntityTypeAndFieldNameAndCommentCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityType := self.Args("entity_type").String()
	iFieldName := self.Args("field_name").String()
	iCommentCount := self.Args("comment_count").MustInt()

	if helper.IsHas(iEntityType) {
		_CommentEntityStatistics, _error := model.GetCommentEntityStatisticsesByEntityTypeAndFieldNameAndCommentCount(offset, limit, iEntityType,iFieldName,iCommentCount)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentEntityStatistics)
	}
	herr.Message = "Can't get to the GetCommentEntityStatisticsesByEntityTypeAndFieldNameAndCommentCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentEntityStatisticsesByEntityTypeAndCidAndLastCommentTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityType := self.Args("entity_type").String()
	iCid := self.Args("cid").MustInt()
	iLastCommentTimestamp := self.Args("last_comment_timestamp").MustInt()

	if helper.IsHas(iEntityType) {
		_CommentEntityStatistics, _error := model.GetCommentEntityStatisticsesByEntityTypeAndCidAndLastCommentTimestamp(offset, limit, iEntityType,iCid,iLastCommentTimestamp)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentEntityStatistics)
	}
	herr.Message = "Can't get to the GetCommentEntityStatisticsesByEntityTypeAndCidAndLastCommentTimestamp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentEntityStatisticsesByEntityTypeAndCidAndLastCommentNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityType := self.Args("entity_type").String()
	iCid := self.Args("cid").MustInt()
	iLastCommentName := self.Args("last_comment_name").String()

	if helper.IsHas(iEntityType) {
		_CommentEntityStatistics, _error := model.GetCommentEntityStatisticsesByEntityTypeAndCidAndLastCommentName(offset, limit, iEntityType,iCid,iLastCommentName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentEntityStatistics)
	}
	herr.Message = "Can't get to the GetCommentEntityStatisticsesByEntityTypeAndCidAndLastCommentName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentEntityStatisticsesByEntityTypeAndCidAndLastCommentUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityType := self.Args("entity_type").String()
	iCid := self.Args("cid").MustInt()
	iLastCommentUid := self.Args("last_comment_uid").MustInt()

	if helper.IsHas(iEntityType) {
		_CommentEntityStatistics, _error := model.GetCommentEntityStatisticsesByEntityTypeAndCidAndLastCommentUid(offset, limit, iEntityType,iCid,iLastCommentUid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentEntityStatistics)
	}
	herr.Message = "Can't get to the GetCommentEntityStatisticsesByEntityTypeAndCidAndLastCommentUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentEntityStatisticsesByEntityTypeAndCidAndCommentCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityType := self.Args("entity_type").String()
	iCid := self.Args("cid").MustInt()
	iCommentCount := self.Args("comment_count").MustInt()

	if helper.IsHas(iEntityType) {
		_CommentEntityStatistics, _error := model.GetCommentEntityStatisticsesByEntityTypeAndCidAndCommentCount(offset, limit, iEntityType,iCid,iCommentCount)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentEntityStatistics)
	}
	herr.Message = "Can't get to the GetCommentEntityStatisticsesByEntityTypeAndCidAndCommentCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentEntityStatisticsesByEntityTypeAndLastCommentTimestampAndLastCommentNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityType := self.Args("entity_type").String()
	iLastCommentTimestamp := self.Args("last_comment_timestamp").MustInt()
	iLastCommentName := self.Args("last_comment_name").String()

	if helper.IsHas(iEntityType) {
		_CommentEntityStatistics, _error := model.GetCommentEntityStatisticsesByEntityTypeAndLastCommentTimestampAndLastCommentName(offset, limit, iEntityType,iLastCommentTimestamp,iLastCommentName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentEntityStatistics)
	}
	herr.Message = "Can't get to the GetCommentEntityStatisticsesByEntityTypeAndLastCommentTimestampAndLastCommentName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentEntityStatisticsesByEntityTypeAndLastCommentTimestampAndLastCommentUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityType := self.Args("entity_type").String()
	iLastCommentTimestamp := self.Args("last_comment_timestamp").MustInt()
	iLastCommentUid := self.Args("last_comment_uid").MustInt()

	if helper.IsHas(iEntityType) {
		_CommentEntityStatistics, _error := model.GetCommentEntityStatisticsesByEntityTypeAndLastCommentTimestampAndLastCommentUid(offset, limit, iEntityType,iLastCommentTimestamp,iLastCommentUid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentEntityStatistics)
	}
	herr.Message = "Can't get to the GetCommentEntityStatisticsesByEntityTypeAndLastCommentTimestampAndLastCommentUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentEntityStatisticsesByEntityTypeAndLastCommentTimestampAndCommentCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityType := self.Args("entity_type").String()
	iLastCommentTimestamp := self.Args("last_comment_timestamp").MustInt()
	iCommentCount := self.Args("comment_count").MustInt()

	if helper.IsHas(iEntityType) {
		_CommentEntityStatistics, _error := model.GetCommentEntityStatisticsesByEntityTypeAndLastCommentTimestampAndCommentCount(offset, limit, iEntityType,iLastCommentTimestamp,iCommentCount)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentEntityStatistics)
	}
	herr.Message = "Can't get to the GetCommentEntityStatisticsesByEntityTypeAndLastCommentTimestampAndCommentCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentEntityStatisticsesByEntityTypeAndLastCommentNameAndLastCommentUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityType := self.Args("entity_type").String()
	iLastCommentName := self.Args("last_comment_name").String()
	iLastCommentUid := self.Args("last_comment_uid").MustInt()

	if helper.IsHas(iEntityType) {
		_CommentEntityStatistics, _error := model.GetCommentEntityStatisticsesByEntityTypeAndLastCommentNameAndLastCommentUid(offset, limit, iEntityType,iLastCommentName,iLastCommentUid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentEntityStatistics)
	}
	herr.Message = "Can't get to the GetCommentEntityStatisticsesByEntityTypeAndLastCommentNameAndLastCommentUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentEntityStatisticsesByEntityTypeAndLastCommentNameAndCommentCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityType := self.Args("entity_type").String()
	iLastCommentName := self.Args("last_comment_name").String()
	iCommentCount := self.Args("comment_count").MustInt()

	if helper.IsHas(iEntityType) {
		_CommentEntityStatistics, _error := model.GetCommentEntityStatisticsesByEntityTypeAndLastCommentNameAndCommentCount(offset, limit, iEntityType,iLastCommentName,iCommentCount)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentEntityStatistics)
	}
	herr.Message = "Can't get to the GetCommentEntityStatisticsesByEntityTypeAndLastCommentNameAndCommentCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentEntityStatisticsesByEntityTypeAndLastCommentUidAndCommentCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityType := self.Args("entity_type").String()
	iLastCommentUid := self.Args("last_comment_uid").MustInt()
	iCommentCount := self.Args("comment_count").MustInt()

	if helper.IsHas(iEntityType) {
		_CommentEntityStatistics, _error := model.GetCommentEntityStatisticsesByEntityTypeAndLastCommentUidAndCommentCount(offset, limit, iEntityType,iLastCommentUid,iCommentCount)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentEntityStatistics)
	}
	herr.Message = "Can't get to the GetCommentEntityStatisticsesByEntityTypeAndLastCommentUidAndCommentCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentEntityStatisticsesByFieldNameAndCidAndLastCommentTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFieldName := self.Args("field_name").String()
	iCid := self.Args("cid").MustInt()
	iLastCommentTimestamp := self.Args("last_comment_timestamp").MustInt()

	if helper.IsHas(iFieldName) {
		_CommentEntityStatistics, _error := model.GetCommentEntityStatisticsesByFieldNameAndCidAndLastCommentTimestamp(offset, limit, iFieldName,iCid,iLastCommentTimestamp)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentEntityStatistics)
	}
	herr.Message = "Can't get to the GetCommentEntityStatisticsesByFieldNameAndCidAndLastCommentTimestamp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentEntityStatisticsesByFieldNameAndCidAndLastCommentNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFieldName := self.Args("field_name").String()
	iCid := self.Args("cid").MustInt()
	iLastCommentName := self.Args("last_comment_name").String()

	if helper.IsHas(iFieldName) {
		_CommentEntityStatistics, _error := model.GetCommentEntityStatisticsesByFieldNameAndCidAndLastCommentName(offset, limit, iFieldName,iCid,iLastCommentName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentEntityStatistics)
	}
	herr.Message = "Can't get to the GetCommentEntityStatisticsesByFieldNameAndCidAndLastCommentName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentEntityStatisticsesByFieldNameAndCidAndLastCommentUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFieldName := self.Args("field_name").String()
	iCid := self.Args("cid").MustInt()
	iLastCommentUid := self.Args("last_comment_uid").MustInt()

	if helper.IsHas(iFieldName) {
		_CommentEntityStatistics, _error := model.GetCommentEntityStatisticsesByFieldNameAndCidAndLastCommentUid(offset, limit, iFieldName,iCid,iLastCommentUid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentEntityStatistics)
	}
	herr.Message = "Can't get to the GetCommentEntityStatisticsesByFieldNameAndCidAndLastCommentUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentEntityStatisticsesByFieldNameAndCidAndCommentCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFieldName := self.Args("field_name").String()
	iCid := self.Args("cid").MustInt()
	iCommentCount := self.Args("comment_count").MustInt()

	if helper.IsHas(iFieldName) {
		_CommentEntityStatistics, _error := model.GetCommentEntityStatisticsesByFieldNameAndCidAndCommentCount(offset, limit, iFieldName,iCid,iCommentCount)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentEntityStatistics)
	}
	herr.Message = "Can't get to the GetCommentEntityStatisticsesByFieldNameAndCidAndCommentCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentEntityStatisticsesByFieldNameAndLastCommentTimestampAndLastCommentNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFieldName := self.Args("field_name").String()
	iLastCommentTimestamp := self.Args("last_comment_timestamp").MustInt()
	iLastCommentName := self.Args("last_comment_name").String()

	if helper.IsHas(iFieldName) {
		_CommentEntityStatistics, _error := model.GetCommentEntityStatisticsesByFieldNameAndLastCommentTimestampAndLastCommentName(offset, limit, iFieldName,iLastCommentTimestamp,iLastCommentName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentEntityStatistics)
	}
	herr.Message = "Can't get to the GetCommentEntityStatisticsesByFieldNameAndLastCommentTimestampAndLastCommentName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentEntityStatisticsesByFieldNameAndLastCommentTimestampAndLastCommentUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFieldName := self.Args("field_name").String()
	iLastCommentTimestamp := self.Args("last_comment_timestamp").MustInt()
	iLastCommentUid := self.Args("last_comment_uid").MustInt()

	if helper.IsHas(iFieldName) {
		_CommentEntityStatistics, _error := model.GetCommentEntityStatisticsesByFieldNameAndLastCommentTimestampAndLastCommentUid(offset, limit, iFieldName,iLastCommentTimestamp,iLastCommentUid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentEntityStatistics)
	}
	herr.Message = "Can't get to the GetCommentEntityStatisticsesByFieldNameAndLastCommentTimestampAndLastCommentUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentEntityStatisticsesByFieldNameAndLastCommentTimestampAndCommentCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFieldName := self.Args("field_name").String()
	iLastCommentTimestamp := self.Args("last_comment_timestamp").MustInt()
	iCommentCount := self.Args("comment_count").MustInt()

	if helper.IsHas(iFieldName) {
		_CommentEntityStatistics, _error := model.GetCommentEntityStatisticsesByFieldNameAndLastCommentTimestampAndCommentCount(offset, limit, iFieldName,iLastCommentTimestamp,iCommentCount)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentEntityStatistics)
	}
	herr.Message = "Can't get to the GetCommentEntityStatisticsesByFieldNameAndLastCommentTimestampAndCommentCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentEntityStatisticsesByFieldNameAndLastCommentNameAndLastCommentUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFieldName := self.Args("field_name").String()
	iLastCommentName := self.Args("last_comment_name").String()
	iLastCommentUid := self.Args("last_comment_uid").MustInt()

	if helper.IsHas(iFieldName) {
		_CommentEntityStatistics, _error := model.GetCommentEntityStatisticsesByFieldNameAndLastCommentNameAndLastCommentUid(offset, limit, iFieldName,iLastCommentName,iLastCommentUid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentEntityStatistics)
	}
	herr.Message = "Can't get to the GetCommentEntityStatisticsesByFieldNameAndLastCommentNameAndLastCommentUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentEntityStatisticsesByFieldNameAndLastCommentNameAndCommentCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFieldName := self.Args("field_name").String()
	iLastCommentName := self.Args("last_comment_name").String()
	iCommentCount := self.Args("comment_count").MustInt()

	if helper.IsHas(iFieldName) {
		_CommentEntityStatistics, _error := model.GetCommentEntityStatisticsesByFieldNameAndLastCommentNameAndCommentCount(offset, limit, iFieldName,iLastCommentName,iCommentCount)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentEntityStatistics)
	}
	herr.Message = "Can't get to the GetCommentEntityStatisticsesByFieldNameAndLastCommentNameAndCommentCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentEntityStatisticsesByFieldNameAndLastCommentUidAndCommentCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFieldName := self.Args("field_name").String()
	iLastCommentUid := self.Args("last_comment_uid").MustInt()
	iCommentCount := self.Args("comment_count").MustInt()

	if helper.IsHas(iFieldName) {
		_CommentEntityStatistics, _error := model.GetCommentEntityStatisticsesByFieldNameAndLastCommentUidAndCommentCount(offset, limit, iFieldName,iLastCommentUid,iCommentCount)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentEntityStatistics)
	}
	herr.Message = "Can't get to the GetCommentEntityStatisticsesByFieldNameAndLastCommentUidAndCommentCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentEntityStatisticsesByCidAndLastCommentTimestampAndLastCommentNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt()
	iLastCommentTimestamp := self.Args("last_comment_timestamp").MustInt()
	iLastCommentName := self.Args("last_comment_name").String()

	if helper.IsHas(iCid) {
		_CommentEntityStatistics, _error := model.GetCommentEntityStatisticsesByCidAndLastCommentTimestampAndLastCommentName(offset, limit, iCid,iLastCommentTimestamp,iLastCommentName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentEntityStatistics)
	}
	herr.Message = "Can't get to the GetCommentEntityStatisticsesByCidAndLastCommentTimestampAndLastCommentName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentEntityStatisticsesByCidAndLastCommentTimestampAndLastCommentUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt()
	iLastCommentTimestamp := self.Args("last_comment_timestamp").MustInt()
	iLastCommentUid := self.Args("last_comment_uid").MustInt()

	if helper.IsHas(iCid) {
		_CommentEntityStatistics, _error := model.GetCommentEntityStatisticsesByCidAndLastCommentTimestampAndLastCommentUid(offset, limit, iCid,iLastCommentTimestamp,iLastCommentUid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentEntityStatistics)
	}
	herr.Message = "Can't get to the GetCommentEntityStatisticsesByCidAndLastCommentTimestampAndLastCommentUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentEntityStatisticsesByCidAndLastCommentTimestampAndCommentCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt()
	iLastCommentTimestamp := self.Args("last_comment_timestamp").MustInt()
	iCommentCount := self.Args("comment_count").MustInt()

	if helper.IsHas(iCid) {
		_CommentEntityStatistics, _error := model.GetCommentEntityStatisticsesByCidAndLastCommentTimestampAndCommentCount(offset, limit, iCid,iLastCommentTimestamp,iCommentCount)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentEntityStatistics)
	}
	herr.Message = "Can't get to the GetCommentEntityStatisticsesByCidAndLastCommentTimestampAndCommentCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentEntityStatisticsesByCidAndLastCommentNameAndLastCommentUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt()
	iLastCommentName := self.Args("last_comment_name").String()
	iLastCommentUid := self.Args("last_comment_uid").MustInt()

	if helper.IsHas(iCid) {
		_CommentEntityStatistics, _error := model.GetCommentEntityStatisticsesByCidAndLastCommentNameAndLastCommentUid(offset, limit, iCid,iLastCommentName,iLastCommentUid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentEntityStatistics)
	}
	herr.Message = "Can't get to the GetCommentEntityStatisticsesByCidAndLastCommentNameAndLastCommentUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentEntityStatisticsesByCidAndLastCommentNameAndCommentCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt()
	iLastCommentName := self.Args("last_comment_name").String()
	iCommentCount := self.Args("comment_count").MustInt()

	if helper.IsHas(iCid) {
		_CommentEntityStatistics, _error := model.GetCommentEntityStatisticsesByCidAndLastCommentNameAndCommentCount(offset, limit, iCid,iLastCommentName,iCommentCount)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentEntityStatistics)
	}
	herr.Message = "Can't get to the GetCommentEntityStatisticsesByCidAndLastCommentNameAndCommentCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentEntityStatisticsesByCidAndLastCommentUidAndCommentCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt()
	iLastCommentUid := self.Args("last_comment_uid").MustInt()
	iCommentCount := self.Args("comment_count").MustInt()

	if helper.IsHas(iCid) {
		_CommentEntityStatistics, _error := model.GetCommentEntityStatisticsesByCidAndLastCommentUidAndCommentCount(offset, limit, iCid,iLastCommentUid,iCommentCount)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentEntityStatistics)
	}
	herr.Message = "Can't get to the GetCommentEntityStatisticsesByCidAndLastCommentUidAndCommentCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentEntityStatisticsesByLastCommentTimestampAndLastCommentNameAndLastCommentUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLastCommentTimestamp := self.Args("last_comment_timestamp").MustInt()
	iLastCommentName := self.Args("last_comment_name").String()
	iLastCommentUid := self.Args("last_comment_uid").MustInt()

	if helper.IsHas(iLastCommentTimestamp) {
		_CommentEntityStatistics, _error := model.GetCommentEntityStatisticsesByLastCommentTimestampAndLastCommentNameAndLastCommentUid(offset, limit, iLastCommentTimestamp,iLastCommentName,iLastCommentUid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentEntityStatistics)
	}
	herr.Message = "Can't get to the GetCommentEntityStatisticsesByLastCommentTimestampAndLastCommentNameAndLastCommentUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentEntityStatisticsesByLastCommentTimestampAndLastCommentNameAndCommentCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLastCommentTimestamp := self.Args("last_comment_timestamp").MustInt()
	iLastCommentName := self.Args("last_comment_name").String()
	iCommentCount := self.Args("comment_count").MustInt()

	if helper.IsHas(iLastCommentTimestamp) {
		_CommentEntityStatistics, _error := model.GetCommentEntityStatisticsesByLastCommentTimestampAndLastCommentNameAndCommentCount(offset, limit, iLastCommentTimestamp,iLastCommentName,iCommentCount)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentEntityStatistics)
	}
	herr.Message = "Can't get to the GetCommentEntityStatisticsesByLastCommentTimestampAndLastCommentNameAndCommentCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentEntityStatisticsesByLastCommentTimestampAndLastCommentUidAndCommentCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLastCommentTimestamp := self.Args("last_comment_timestamp").MustInt()
	iLastCommentUid := self.Args("last_comment_uid").MustInt()
	iCommentCount := self.Args("comment_count").MustInt()

	if helper.IsHas(iLastCommentTimestamp) {
		_CommentEntityStatistics, _error := model.GetCommentEntityStatisticsesByLastCommentTimestampAndLastCommentUidAndCommentCount(offset, limit, iLastCommentTimestamp,iLastCommentUid,iCommentCount)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentEntityStatistics)
	}
	herr.Message = "Can't get to the GetCommentEntityStatisticsesByLastCommentTimestampAndLastCommentUidAndCommentCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentEntityStatisticsesByLastCommentNameAndLastCommentUidAndCommentCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLastCommentName := self.Args("last_comment_name").String()
	iLastCommentUid := self.Args("last_comment_uid").MustInt()
	iCommentCount := self.Args("comment_count").MustInt()

	if helper.IsHas(iLastCommentName) {
		_CommentEntityStatistics, _error := model.GetCommentEntityStatisticsesByLastCommentNameAndLastCommentUidAndCommentCount(offset, limit, iLastCommentName,iLastCommentUid,iCommentCount)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentEntityStatistics)
	}
	herr.Message = "Can't get to the GetCommentEntityStatisticsesByLastCommentNameAndLastCommentUidAndCommentCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentEntityStatisticsesByEntityIdAndEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt64()
	iEntityType := self.Args("entity_type").String()

	if helper.IsHas(iEntityId) {
		_CommentEntityStatistics, _error := model.GetCommentEntityStatisticsesByEntityIdAndEntityType(offset, limit, iEntityId,iEntityType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentEntityStatistics)
	}
	herr.Message = "Can't get to the GetCommentEntityStatisticsesByEntityIdAndEntityType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentEntityStatisticsesByEntityIdAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt64()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iEntityId) {
		_CommentEntityStatistics, _error := model.GetCommentEntityStatisticsesByEntityIdAndFieldName(offset, limit, iEntityId,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentEntityStatistics)
	}
	herr.Message = "Can't get to the GetCommentEntityStatisticsesByEntityIdAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentEntityStatisticsesByEntityIdAndCidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt64()
	iCid := self.Args("cid").MustInt()

	if helper.IsHas(iEntityId) {
		_CommentEntityStatistics, _error := model.GetCommentEntityStatisticsesByEntityIdAndCid(offset, limit, iEntityId,iCid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentEntityStatistics)
	}
	herr.Message = "Can't get to the GetCommentEntityStatisticsesByEntityIdAndCid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentEntityStatisticsesByEntityIdAndLastCommentTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt64()
	iLastCommentTimestamp := self.Args("last_comment_timestamp").MustInt()

	if helper.IsHas(iEntityId) {
		_CommentEntityStatistics, _error := model.GetCommentEntityStatisticsesByEntityIdAndLastCommentTimestamp(offset, limit, iEntityId,iLastCommentTimestamp)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentEntityStatistics)
	}
	herr.Message = "Can't get to the GetCommentEntityStatisticsesByEntityIdAndLastCommentTimestamp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentEntityStatisticsesByEntityIdAndLastCommentNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt64()
	iLastCommentName := self.Args("last_comment_name").String()

	if helper.IsHas(iEntityId) {
		_CommentEntityStatistics, _error := model.GetCommentEntityStatisticsesByEntityIdAndLastCommentName(offset, limit, iEntityId,iLastCommentName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentEntityStatistics)
	}
	herr.Message = "Can't get to the GetCommentEntityStatisticsesByEntityIdAndLastCommentName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentEntityStatisticsesByEntityIdAndLastCommentUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt64()
	iLastCommentUid := self.Args("last_comment_uid").MustInt()

	if helper.IsHas(iEntityId) {
		_CommentEntityStatistics, _error := model.GetCommentEntityStatisticsesByEntityIdAndLastCommentUid(offset, limit, iEntityId,iLastCommentUid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentEntityStatistics)
	}
	herr.Message = "Can't get to the GetCommentEntityStatisticsesByEntityIdAndLastCommentUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentEntityStatisticsesByEntityIdAndCommentCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt64()
	iCommentCount := self.Args("comment_count").MustInt()

	if helper.IsHas(iEntityId) {
		_CommentEntityStatistics, _error := model.GetCommentEntityStatisticsesByEntityIdAndCommentCount(offset, limit, iEntityId,iCommentCount)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentEntityStatistics)
	}
	herr.Message = "Can't get to the GetCommentEntityStatisticsesByEntityIdAndCommentCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentEntityStatisticsesByEntityTypeAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityType := self.Args("entity_type").String()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iEntityType) {
		_CommentEntityStatistics, _error := model.GetCommentEntityStatisticsesByEntityTypeAndFieldName(offset, limit, iEntityType,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentEntityStatistics)
	}
	herr.Message = "Can't get to the GetCommentEntityStatisticsesByEntityTypeAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentEntityStatisticsesByEntityTypeAndCidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityType := self.Args("entity_type").String()
	iCid := self.Args("cid").MustInt()

	if helper.IsHas(iEntityType) {
		_CommentEntityStatistics, _error := model.GetCommentEntityStatisticsesByEntityTypeAndCid(offset, limit, iEntityType,iCid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentEntityStatistics)
	}
	herr.Message = "Can't get to the GetCommentEntityStatisticsesByEntityTypeAndCid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentEntityStatisticsesByEntityTypeAndLastCommentTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityType := self.Args("entity_type").String()
	iLastCommentTimestamp := self.Args("last_comment_timestamp").MustInt()

	if helper.IsHas(iEntityType) {
		_CommentEntityStatistics, _error := model.GetCommentEntityStatisticsesByEntityTypeAndLastCommentTimestamp(offset, limit, iEntityType,iLastCommentTimestamp)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentEntityStatistics)
	}
	herr.Message = "Can't get to the GetCommentEntityStatisticsesByEntityTypeAndLastCommentTimestamp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentEntityStatisticsesByEntityTypeAndLastCommentNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityType := self.Args("entity_type").String()
	iLastCommentName := self.Args("last_comment_name").String()

	if helper.IsHas(iEntityType) {
		_CommentEntityStatistics, _error := model.GetCommentEntityStatisticsesByEntityTypeAndLastCommentName(offset, limit, iEntityType,iLastCommentName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentEntityStatistics)
	}
	herr.Message = "Can't get to the GetCommentEntityStatisticsesByEntityTypeAndLastCommentName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentEntityStatisticsesByEntityTypeAndLastCommentUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityType := self.Args("entity_type").String()
	iLastCommentUid := self.Args("last_comment_uid").MustInt()

	if helper.IsHas(iEntityType) {
		_CommentEntityStatistics, _error := model.GetCommentEntityStatisticsesByEntityTypeAndLastCommentUid(offset, limit, iEntityType,iLastCommentUid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentEntityStatistics)
	}
	herr.Message = "Can't get to the GetCommentEntityStatisticsesByEntityTypeAndLastCommentUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentEntityStatisticsesByEntityTypeAndCommentCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityType := self.Args("entity_type").String()
	iCommentCount := self.Args("comment_count").MustInt()

	if helper.IsHas(iEntityType) {
		_CommentEntityStatistics, _error := model.GetCommentEntityStatisticsesByEntityTypeAndCommentCount(offset, limit, iEntityType,iCommentCount)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentEntityStatistics)
	}
	herr.Message = "Can't get to the GetCommentEntityStatisticsesByEntityTypeAndCommentCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentEntityStatisticsesByFieldNameAndCidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFieldName := self.Args("field_name").String()
	iCid := self.Args("cid").MustInt()

	if helper.IsHas(iFieldName) {
		_CommentEntityStatistics, _error := model.GetCommentEntityStatisticsesByFieldNameAndCid(offset, limit, iFieldName,iCid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentEntityStatistics)
	}
	herr.Message = "Can't get to the GetCommentEntityStatisticsesByFieldNameAndCid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentEntityStatisticsesByFieldNameAndLastCommentTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFieldName := self.Args("field_name").String()
	iLastCommentTimestamp := self.Args("last_comment_timestamp").MustInt()

	if helper.IsHas(iFieldName) {
		_CommentEntityStatistics, _error := model.GetCommentEntityStatisticsesByFieldNameAndLastCommentTimestamp(offset, limit, iFieldName,iLastCommentTimestamp)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentEntityStatistics)
	}
	herr.Message = "Can't get to the GetCommentEntityStatisticsesByFieldNameAndLastCommentTimestamp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentEntityStatisticsesByFieldNameAndLastCommentNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFieldName := self.Args("field_name").String()
	iLastCommentName := self.Args("last_comment_name").String()

	if helper.IsHas(iFieldName) {
		_CommentEntityStatistics, _error := model.GetCommentEntityStatisticsesByFieldNameAndLastCommentName(offset, limit, iFieldName,iLastCommentName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentEntityStatistics)
	}
	herr.Message = "Can't get to the GetCommentEntityStatisticsesByFieldNameAndLastCommentName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentEntityStatisticsesByFieldNameAndLastCommentUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFieldName := self.Args("field_name").String()
	iLastCommentUid := self.Args("last_comment_uid").MustInt()

	if helper.IsHas(iFieldName) {
		_CommentEntityStatistics, _error := model.GetCommentEntityStatisticsesByFieldNameAndLastCommentUid(offset, limit, iFieldName,iLastCommentUid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentEntityStatistics)
	}
	herr.Message = "Can't get to the GetCommentEntityStatisticsesByFieldNameAndLastCommentUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentEntityStatisticsesByFieldNameAndCommentCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFieldName := self.Args("field_name").String()
	iCommentCount := self.Args("comment_count").MustInt()

	if helper.IsHas(iFieldName) {
		_CommentEntityStatistics, _error := model.GetCommentEntityStatisticsesByFieldNameAndCommentCount(offset, limit, iFieldName,iCommentCount)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentEntityStatistics)
	}
	herr.Message = "Can't get to the GetCommentEntityStatisticsesByFieldNameAndCommentCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentEntityStatisticsesByCidAndLastCommentTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt()
	iLastCommentTimestamp := self.Args("last_comment_timestamp").MustInt()

	if helper.IsHas(iCid) {
		_CommentEntityStatistics, _error := model.GetCommentEntityStatisticsesByCidAndLastCommentTimestamp(offset, limit, iCid,iLastCommentTimestamp)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentEntityStatistics)
	}
	herr.Message = "Can't get to the GetCommentEntityStatisticsesByCidAndLastCommentTimestamp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentEntityStatisticsesByCidAndLastCommentNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt()
	iLastCommentName := self.Args("last_comment_name").String()

	if helper.IsHas(iCid) {
		_CommentEntityStatistics, _error := model.GetCommentEntityStatisticsesByCidAndLastCommentName(offset, limit, iCid,iLastCommentName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentEntityStatistics)
	}
	herr.Message = "Can't get to the GetCommentEntityStatisticsesByCidAndLastCommentName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentEntityStatisticsesByCidAndLastCommentUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt()
	iLastCommentUid := self.Args("last_comment_uid").MustInt()

	if helper.IsHas(iCid) {
		_CommentEntityStatistics, _error := model.GetCommentEntityStatisticsesByCidAndLastCommentUid(offset, limit, iCid,iLastCommentUid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentEntityStatistics)
	}
	herr.Message = "Can't get to the GetCommentEntityStatisticsesByCidAndLastCommentUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentEntityStatisticsesByCidAndCommentCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt()
	iCommentCount := self.Args("comment_count").MustInt()

	if helper.IsHas(iCid) {
		_CommentEntityStatistics, _error := model.GetCommentEntityStatisticsesByCidAndCommentCount(offset, limit, iCid,iCommentCount)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentEntityStatistics)
	}
	herr.Message = "Can't get to the GetCommentEntityStatisticsesByCidAndCommentCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentEntityStatisticsesByLastCommentTimestampAndLastCommentNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLastCommentTimestamp := self.Args("last_comment_timestamp").MustInt()
	iLastCommentName := self.Args("last_comment_name").String()

	if helper.IsHas(iLastCommentTimestamp) {
		_CommentEntityStatistics, _error := model.GetCommentEntityStatisticsesByLastCommentTimestampAndLastCommentName(offset, limit, iLastCommentTimestamp,iLastCommentName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentEntityStatistics)
	}
	herr.Message = "Can't get to the GetCommentEntityStatisticsesByLastCommentTimestampAndLastCommentName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentEntityStatisticsesByLastCommentTimestampAndLastCommentUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLastCommentTimestamp := self.Args("last_comment_timestamp").MustInt()
	iLastCommentUid := self.Args("last_comment_uid").MustInt()

	if helper.IsHas(iLastCommentTimestamp) {
		_CommentEntityStatistics, _error := model.GetCommentEntityStatisticsesByLastCommentTimestampAndLastCommentUid(offset, limit, iLastCommentTimestamp,iLastCommentUid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentEntityStatistics)
	}
	herr.Message = "Can't get to the GetCommentEntityStatisticsesByLastCommentTimestampAndLastCommentUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentEntityStatisticsesByLastCommentTimestampAndCommentCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLastCommentTimestamp := self.Args("last_comment_timestamp").MustInt()
	iCommentCount := self.Args("comment_count").MustInt()

	if helper.IsHas(iLastCommentTimestamp) {
		_CommentEntityStatistics, _error := model.GetCommentEntityStatisticsesByLastCommentTimestampAndCommentCount(offset, limit, iLastCommentTimestamp,iCommentCount)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentEntityStatistics)
	}
	herr.Message = "Can't get to the GetCommentEntityStatisticsesByLastCommentTimestampAndCommentCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentEntityStatisticsesByLastCommentNameAndLastCommentUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLastCommentName := self.Args("last_comment_name").String()
	iLastCommentUid := self.Args("last_comment_uid").MustInt()

	if helper.IsHas(iLastCommentName) {
		_CommentEntityStatistics, _error := model.GetCommentEntityStatisticsesByLastCommentNameAndLastCommentUid(offset, limit, iLastCommentName,iLastCommentUid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentEntityStatistics)
	}
	herr.Message = "Can't get to the GetCommentEntityStatisticsesByLastCommentNameAndLastCommentUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentEntityStatisticsesByLastCommentNameAndCommentCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLastCommentName := self.Args("last_comment_name").String()
	iCommentCount := self.Args("comment_count").MustInt()

	if helper.IsHas(iLastCommentName) {
		_CommentEntityStatistics, _error := model.GetCommentEntityStatisticsesByLastCommentNameAndCommentCount(offset, limit, iLastCommentName,iCommentCount)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentEntityStatistics)
	}
	herr.Message = "Can't get to the GetCommentEntityStatisticsesByLastCommentNameAndCommentCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentEntityStatisticsesByLastCommentUidAndCommentCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLastCommentUid := self.Args("last_comment_uid").MustInt()
	iCommentCount := self.Args("comment_count").MustInt()

	if helper.IsHas(iLastCommentUid) {
		_CommentEntityStatistics, _error := model.GetCommentEntityStatisticsesByLastCommentUidAndCommentCount(offset, limit, iLastCommentUid,iCommentCount)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentEntityStatistics)
	}
	herr.Message = "Can't get to the GetCommentEntityStatisticsesByLastCommentUidAndCommentCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentEntityStatisticsesHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	if (offset > 0) && (len(field) > 0) {
		_CommentEntityStatistics, _error := model.GetCommentEntityStatisticses(offset, limit, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentEntityStatistics)
	}
	herr.Message = "Can't get to the GetCommentEntityStatisticses' args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasCommentEntityStatisticsViaEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iEntityId := self.Args("entity_id").MustInt64()
	if helper.IsHas(iEntityId) {
		_CommentEntityStatistics := model.HasCommentEntityStatisticsViaEntityId(iEntityId)
		var m = map[string]interface{}{}
		m["comment_entity_statistics"] = _CommentEntityStatistics
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasCommentEntityStatisticsViaEntityId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasCommentEntityStatisticsViaEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iEntityType := self.Args("entity_type").String()
	if helper.IsHas(iEntityType) {
		_CommentEntityStatistics := model.HasCommentEntityStatisticsViaEntityType(iEntityType)
		var m = map[string]interface{}{}
		m["comment_entity_statistics"] = _CommentEntityStatistics
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasCommentEntityStatisticsViaEntityType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasCommentEntityStatisticsViaFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iFieldName := self.Args("field_name").String()
	if helper.IsHas(iFieldName) {
		_CommentEntityStatistics := model.HasCommentEntityStatisticsViaFieldName(iFieldName)
		var m = map[string]interface{}{}
		m["comment_entity_statistics"] = _CommentEntityStatistics
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasCommentEntityStatisticsViaFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasCommentEntityStatisticsViaCidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iCid := self.Args("cid").MustInt()
	if helper.IsHas(iCid) {
		_CommentEntityStatistics := model.HasCommentEntityStatisticsViaCid(iCid)
		var m = map[string]interface{}{}
		m["comment_entity_statistics"] = _CommentEntityStatistics
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasCommentEntityStatisticsViaCid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasCommentEntityStatisticsViaLastCommentTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iLastCommentTimestamp := self.Args("last_comment_timestamp").MustInt()
	if helper.IsHas(iLastCommentTimestamp) {
		_CommentEntityStatistics := model.HasCommentEntityStatisticsViaLastCommentTimestamp(iLastCommentTimestamp)
		var m = map[string]interface{}{}
		m["comment_entity_statistics"] = _CommentEntityStatistics
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasCommentEntityStatisticsViaLastCommentTimestamp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasCommentEntityStatisticsViaLastCommentNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iLastCommentName := self.Args("last_comment_name").String()
	if helper.IsHas(iLastCommentName) {
		_CommentEntityStatistics := model.HasCommentEntityStatisticsViaLastCommentName(iLastCommentName)
		var m = map[string]interface{}{}
		m["comment_entity_statistics"] = _CommentEntityStatistics
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasCommentEntityStatisticsViaLastCommentName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasCommentEntityStatisticsViaLastCommentUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iLastCommentUid := self.Args("last_comment_uid").MustInt()
	if helper.IsHas(iLastCommentUid) {
		_CommentEntityStatistics := model.HasCommentEntityStatisticsViaLastCommentUid(iLastCommentUid)
		var m = map[string]interface{}{}
		m["comment_entity_statistics"] = _CommentEntityStatistics
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasCommentEntityStatisticsViaLastCommentUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasCommentEntityStatisticsViaCommentCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iCommentCount := self.Args("comment_count").MustInt()
	if helper.IsHas(iCommentCount) {
		_CommentEntityStatistics := model.HasCommentEntityStatisticsViaCommentCount(iCommentCount)
		var m = map[string]interface{}{}
		m["comment_entity_statistics"] = _CommentEntityStatistics
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasCommentEntityStatisticsViaCommentCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentEntityStatisticsViaEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iEntityId := self.Args("entity_id").MustInt64()
	if helper.IsHas(iEntityId) {
		_CommentEntityStatistics, _error := model.GetCommentEntityStatisticsViaEntityId(iEntityId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentEntityStatistics)
	}
	herr.Message = "Can't get to the GetCommentEntityStatisticsViaEntityId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentEntityStatisticsViaEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iEntityType := self.Args("entity_type").String()
	if helper.IsHas(iEntityType) {
		_CommentEntityStatistics, _error := model.GetCommentEntityStatisticsViaEntityType(iEntityType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentEntityStatistics)
	}
	herr.Message = "Can't get to the GetCommentEntityStatisticsViaEntityType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentEntityStatisticsViaFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iFieldName := self.Args("field_name").String()
	if helper.IsHas(iFieldName) {
		_CommentEntityStatistics, _error := model.GetCommentEntityStatisticsViaFieldName(iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentEntityStatistics)
	}
	herr.Message = "Can't get to the GetCommentEntityStatisticsViaFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentEntityStatisticsViaCidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iCid := self.Args("cid").MustInt()
	if helper.IsHas(iCid) {
		_CommentEntityStatistics, _error := model.GetCommentEntityStatisticsViaCid(iCid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentEntityStatistics)
	}
	herr.Message = "Can't get to the GetCommentEntityStatisticsViaCid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentEntityStatisticsViaLastCommentTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iLastCommentTimestamp := self.Args("last_comment_timestamp").MustInt()
	if helper.IsHas(iLastCommentTimestamp) {
		_CommentEntityStatistics, _error := model.GetCommentEntityStatisticsViaLastCommentTimestamp(iLastCommentTimestamp)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentEntityStatistics)
	}
	herr.Message = "Can't get to the GetCommentEntityStatisticsViaLastCommentTimestamp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentEntityStatisticsViaLastCommentNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iLastCommentName := self.Args("last_comment_name").String()
	if helper.IsHas(iLastCommentName) {
		_CommentEntityStatistics, _error := model.GetCommentEntityStatisticsViaLastCommentName(iLastCommentName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentEntityStatistics)
	}
	herr.Message = "Can't get to the GetCommentEntityStatisticsViaLastCommentName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentEntityStatisticsViaLastCommentUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iLastCommentUid := self.Args("last_comment_uid").MustInt()
	if helper.IsHas(iLastCommentUid) {
		_CommentEntityStatistics, _error := model.GetCommentEntityStatisticsViaLastCommentUid(iLastCommentUid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentEntityStatistics)
	}
	herr.Message = "Can't get to the GetCommentEntityStatisticsViaLastCommentUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentEntityStatisticsViaCommentCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iCommentCount := self.Args("comment_count").MustInt()
	if helper.IsHas(iCommentCount) {
		_CommentEntityStatistics, _error := model.GetCommentEntityStatisticsViaCommentCount(iCommentCount)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentEntityStatistics)
	}
	herr.Message = "Can't get to the GetCommentEntityStatisticsViaCommentCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetCommentEntityStatisticsViaEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	EntityId_ := self.Args("entity_id").MustInt64()
	if helper.IsHas(EntityId_) {
		var iCommentEntityStatistics model.CommentEntityStatistics
		self.Bind(&iCommentEntityStatistics)
		_CommentEntityStatistics, _error := model.SetCommentEntityStatisticsViaEntityId(EntityId_, &iCommentEntityStatistics)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentEntityStatistics)
	}
	herr.Message = "Can't get to the SetCommentEntityStatisticsViaEntityId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetCommentEntityStatisticsViaEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	EntityType_ := self.Args("entity_type").String()
	if helper.IsHas(EntityType_) {
		var iCommentEntityStatistics model.CommentEntityStatistics
		self.Bind(&iCommentEntityStatistics)
		_CommentEntityStatistics, _error := model.SetCommentEntityStatisticsViaEntityType(EntityType_, &iCommentEntityStatistics)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentEntityStatistics)
	}
	herr.Message = "Can't get to the SetCommentEntityStatisticsViaEntityType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetCommentEntityStatisticsViaFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	FieldName_ := self.Args("field_name").String()
	if helper.IsHas(FieldName_) {
		var iCommentEntityStatistics model.CommentEntityStatistics
		self.Bind(&iCommentEntityStatistics)
		_CommentEntityStatistics, _error := model.SetCommentEntityStatisticsViaFieldName(FieldName_, &iCommentEntityStatistics)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentEntityStatistics)
	}
	herr.Message = "Can't get to the SetCommentEntityStatisticsViaFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetCommentEntityStatisticsViaCidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Cid_ := self.Args("cid").MustInt()
	if helper.IsHas(Cid_) {
		var iCommentEntityStatistics model.CommentEntityStatistics
		self.Bind(&iCommentEntityStatistics)
		_CommentEntityStatistics, _error := model.SetCommentEntityStatisticsViaCid(Cid_, &iCommentEntityStatistics)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentEntityStatistics)
	}
	herr.Message = "Can't get to the SetCommentEntityStatisticsViaCid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetCommentEntityStatisticsViaLastCommentTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	LastCommentTimestamp_ := self.Args("last_comment_timestamp").MustInt()
	if helper.IsHas(LastCommentTimestamp_) {
		var iCommentEntityStatistics model.CommentEntityStatistics
		self.Bind(&iCommentEntityStatistics)
		_CommentEntityStatistics, _error := model.SetCommentEntityStatisticsViaLastCommentTimestamp(LastCommentTimestamp_, &iCommentEntityStatistics)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentEntityStatistics)
	}
	herr.Message = "Can't get to the SetCommentEntityStatisticsViaLastCommentTimestamp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetCommentEntityStatisticsViaLastCommentNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	LastCommentName_ := self.Args("last_comment_name").String()
	if helper.IsHas(LastCommentName_) {
		var iCommentEntityStatistics model.CommentEntityStatistics
		self.Bind(&iCommentEntityStatistics)
		_CommentEntityStatistics, _error := model.SetCommentEntityStatisticsViaLastCommentName(LastCommentName_, &iCommentEntityStatistics)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentEntityStatistics)
	}
	herr.Message = "Can't get to the SetCommentEntityStatisticsViaLastCommentName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetCommentEntityStatisticsViaLastCommentUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	LastCommentUid_ := self.Args("last_comment_uid").MustInt()
	if helper.IsHas(LastCommentUid_) {
		var iCommentEntityStatistics model.CommentEntityStatistics
		self.Bind(&iCommentEntityStatistics)
		_CommentEntityStatistics, _error := model.SetCommentEntityStatisticsViaLastCommentUid(LastCommentUid_, &iCommentEntityStatistics)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentEntityStatistics)
	}
	herr.Message = "Can't get to the SetCommentEntityStatisticsViaLastCommentUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetCommentEntityStatisticsViaCommentCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	CommentCount_ := self.Args("comment_count").MustInt()
	if helper.IsHas(CommentCount_) {
		var iCommentEntityStatistics model.CommentEntityStatistics
		self.Bind(&iCommentEntityStatistics)
		_CommentEntityStatistics, _error := model.SetCommentEntityStatisticsViaCommentCount(CommentCount_, &iCommentEntityStatistics)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentEntityStatistics)
	}
	herr.Message = "Can't get to the SetCommentEntityStatisticsViaCommentCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostAddCommentEntityStatisticsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	EntityId_ := self.Args("entity_id").MustInt64()
	EntityType_ := self.Args("entity_type").String()
	FieldName_ := self.Args("field_name").String()
	Cid_ := self.Args("cid").MustInt()
	LastCommentTimestamp_ := self.Args("last_comment_timestamp").MustInt()
	LastCommentName_ := self.Args("last_comment_name").String()
	LastCommentUid_ := self.Args("last_comment_uid").MustInt()
	CommentCount_ := self.Args("comment_count").MustInt()

	if helper.IsHas(EntityId_) {
		_error := model.AddCommentEntityStatistics(EntityId_,EntityType_,FieldName_,Cid_,LastCommentTimestamp_,LastCommentName_,LastCommentUid_,CommentCount_)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		herr.Message = "StatusOK"
		herr.Status = macross.StatusOK
		return self.JSON(herr)
	}
	herr.Message = "Can't get to the AddCommentEntityStatistics's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostCommentEntityStatisticsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iCommentEntityStatistics model.CommentEntityStatistics
	self.Bind(&iCommentEntityStatistics)
	_int64, _error := model.PostCommentEntityStatistics(&iCommentEntityStatistics)
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

func PutCommentEntityStatisticsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iCommentEntityStatistics model.CommentEntityStatistics
	self.Bind(&iCommentEntityStatistics)
	_int64, _error := model.PutCommentEntityStatistics(&iCommentEntityStatistics)
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

func PutCommentEntityStatisticsViaEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	EntityId_ := self.Args("entity_id").MustInt64()
	var iCommentEntityStatistics model.CommentEntityStatistics
	self.Bind(&iCommentEntityStatistics)
	_int64, _error := model.PutCommentEntityStatisticsViaEntityId(EntityId_, &iCommentEntityStatistics)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutCommentEntityStatisticsViaEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	EntityType_ := self.Args("entity_type").String()
	var iCommentEntityStatistics model.CommentEntityStatistics
	self.Bind(&iCommentEntityStatistics)
	_int64, _error := model.PutCommentEntityStatisticsViaEntityType(EntityType_, &iCommentEntityStatistics)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutCommentEntityStatisticsViaFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	FieldName_ := self.Args("field_name").String()
	var iCommentEntityStatistics model.CommentEntityStatistics
	self.Bind(&iCommentEntityStatistics)
	_int64, _error := model.PutCommentEntityStatisticsViaFieldName(FieldName_, &iCommentEntityStatistics)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutCommentEntityStatisticsViaCidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Cid_ := self.Args("cid").MustInt()
	var iCommentEntityStatistics model.CommentEntityStatistics
	self.Bind(&iCommentEntityStatistics)
	_int64, _error := model.PutCommentEntityStatisticsViaCid(Cid_, &iCommentEntityStatistics)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutCommentEntityStatisticsViaLastCommentTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	LastCommentTimestamp_ := self.Args("last_comment_timestamp").MustInt()
	var iCommentEntityStatistics model.CommentEntityStatistics
	self.Bind(&iCommentEntityStatistics)
	_int64, _error := model.PutCommentEntityStatisticsViaLastCommentTimestamp(LastCommentTimestamp_, &iCommentEntityStatistics)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutCommentEntityStatisticsViaLastCommentNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	LastCommentName_ := self.Args("last_comment_name").String()
	var iCommentEntityStatistics model.CommentEntityStatistics
	self.Bind(&iCommentEntityStatistics)
	_int64, _error := model.PutCommentEntityStatisticsViaLastCommentName(LastCommentName_, &iCommentEntityStatistics)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutCommentEntityStatisticsViaLastCommentUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	LastCommentUid_ := self.Args("last_comment_uid").MustInt()
	var iCommentEntityStatistics model.CommentEntityStatistics
	self.Bind(&iCommentEntityStatistics)
	_int64, _error := model.PutCommentEntityStatisticsViaLastCommentUid(LastCommentUid_, &iCommentEntityStatistics)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutCommentEntityStatisticsViaCommentCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	CommentCount_ := self.Args("comment_count").MustInt()
	var iCommentEntityStatistics model.CommentEntityStatistics
	self.Bind(&iCommentEntityStatistics)
	_int64, _error := model.PutCommentEntityStatisticsViaCommentCount(CommentCount_, &iCommentEntityStatistics)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateCommentEntityStatisticsViaEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	EntityId_ := self.Args("entity_id").MustInt64()
	var iCommentEntityStatistics model.CommentEntityStatistics
	self.Bind(&iCommentEntityStatistics)
	var iMap = helper.StructToMap(iCommentEntityStatistics)
	_error := model.UpdateCommentEntityStatisticsViaEntityId(EntityId_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateCommentEntityStatisticsViaEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	EntityType_ := self.Args("entity_type").String()
	var iCommentEntityStatistics model.CommentEntityStatistics
	self.Bind(&iCommentEntityStatistics)
	var iMap = helper.StructToMap(iCommentEntityStatistics)
	_error := model.UpdateCommentEntityStatisticsViaEntityType(EntityType_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateCommentEntityStatisticsViaFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	FieldName_ := self.Args("field_name").String()
	var iCommentEntityStatistics model.CommentEntityStatistics
	self.Bind(&iCommentEntityStatistics)
	var iMap = helper.StructToMap(iCommentEntityStatistics)
	_error := model.UpdateCommentEntityStatisticsViaFieldName(FieldName_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateCommentEntityStatisticsViaCidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Cid_ := self.Args("cid").MustInt()
	var iCommentEntityStatistics model.CommentEntityStatistics
	self.Bind(&iCommentEntityStatistics)
	var iMap = helper.StructToMap(iCommentEntityStatistics)
	_error := model.UpdateCommentEntityStatisticsViaCid(Cid_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateCommentEntityStatisticsViaLastCommentTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	LastCommentTimestamp_ := self.Args("last_comment_timestamp").MustInt()
	var iCommentEntityStatistics model.CommentEntityStatistics
	self.Bind(&iCommentEntityStatistics)
	var iMap = helper.StructToMap(iCommentEntityStatistics)
	_error := model.UpdateCommentEntityStatisticsViaLastCommentTimestamp(LastCommentTimestamp_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateCommentEntityStatisticsViaLastCommentNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	LastCommentName_ := self.Args("last_comment_name").String()
	var iCommentEntityStatistics model.CommentEntityStatistics
	self.Bind(&iCommentEntityStatistics)
	var iMap = helper.StructToMap(iCommentEntityStatistics)
	_error := model.UpdateCommentEntityStatisticsViaLastCommentName(LastCommentName_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateCommentEntityStatisticsViaLastCommentUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	LastCommentUid_ := self.Args("last_comment_uid").MustInt()
	var iCommentEntityStatistics model.CommentEntityStatistics
	self.Bind(&iCommentEntityStatistics)
	var iMap = helper.StructToMap(iCommentEntityStatistics)
	_error := model.UpdateCommentEntityStatisticsViaLastCommentUid(LastCommentUid_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateCommentEntityStatisticsViaCommentCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	CommentCount_ := self.Args("comment_count").MustInt()
	var iCommentEntityStatistics model.CommentEntityStatistics
	self.Bind(&iCommentEntityStatistics)
	var iMap = helper.StructToMap(iCommentEntityStatistics)
	_error := model.UpdateCommentEntityStatisticsViaCommentCount(CommentCount_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteCommentEntityStatisticsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	EntityId_ := self.Args("entity_id").MustInt64()
	_int64, _error := model.DeleteCommentEntityStatistics(EntityId_)
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

func DeleteCommentEntityStatisticsViaEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	EntityId_ := self.Args("entity_id").MustInt64()
	_error := model.DeleteCommentEntityStatisticsViaEntityId(EntityId_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteCommentEntityStatisticsViaEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	EntityType_ := self.Args("entity_type").String()
	_error := model.DeleteCommentEntityStatisticsViaEntityType(EntityType_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteCommentEntityStatisticsViaFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	FieldName_ := self.Args("field_name").String()
	_error := model.DeleteCommentEntityStatisticsViaFieldName(FieldName_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteCommentEntityStatisticsViaCidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Cid_ := self.Args("cid").MustInt()
	_error := model.DeleteCommentEntityStatisticsViaCid(Cid_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteCommentEntityStatisticsViaLastCommentTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	LastCommentTimestamp_ := self.Args("last_comment_timestamp").MustInt()
	_error := model.DeleteCommentEntityStatisticsViaLastCommentTimestamp(LastCommentTimestamp_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteCommentEntityStatisticsViaLastCommentNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	LastCommentName_ := self.Args("last_comment_name").String()
	_error := model.DeleteCommentEntityStatisticsViaLastCommentName(LastCommentName_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteCommentEntityStatisticsViaLastCommentUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	LastCommentUid_ := self.Args("last_comment_uid").MustInt()
	_error := model.DeleteCommentEntityStatisticsViaLastCommentUid(LastCommentUid_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteCommentEntityStatisticsViaCommentCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	CommentCount_ := self.Args("comment_count").MustInt()
	_error := model.DeleteCommentEntityStatisticsViaCommentCount(CommentCount_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}
