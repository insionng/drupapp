package handler

import (
	"github.com/insionng/drupapp/helper"
	"github.com/insionng/drupapp/model"
	"github.com/insionng/macross"
)

func GetBatchsCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()

	if offset > 0 {
		var m = map[string]interface{}{}
		_int64, _error := model.GetBatchsCount(offset, limit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		m["batchsCount"] = _int64
		return self.JSON(m)
	}
	herr.Message = "Can't get to the GetBatchsCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBatchCountViaBidHandler(self *macross.Context) error {
	Bid_ := self.Args("bid").MustInt64()
	_int64 := model.GetBatchCountViaBid(Bid_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["batchCount"] = 0
	}
	m["batchCount"] = _int64
	return self.JSON(m)
}

func GetBatchCountViaTokenHandler(self *macross.Context) error {
	Token_ := self.Args("token").String()
	_int64 := model.GetBatchCountViaToken(Token_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["batchCount"] = 0
	}
	m["batchCount"] = _int64
	return self.JSON(m)
}

func GetBatchCountViaTimestampHandler(self *macross.Context) error {
	Timestamp_ := self.Args("timestamp").MustInt()
	_int64 := model.GetBatchCountViaTimestamp(Timestamp_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["batchCount"] = 0
	}
	m["batchCount"] = _int64
	return self.JSON(m)
}

func GetBatchCountViaBatchHandler(self *macross.Context) error {
	Batch_ := self.Args("batch").Bytes()
	_int64 := model.GetBatchCountViaBatch(Batch_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["batchCount"] = 0
	}
	m["batchCount"] = _int64
	return self.JSON(m)
}

func GetBatchsViaBidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iBid := self.Args("bid").MustInt64()
	if (offset > 0) && helper.IsHas(iBid) {
		_Batch, _error := model.GetBatchsViaBid(offset, limit, iBid, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Batch)
	}
	herr.Message = "Can't get to the GetBatchsViaBid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBatchsViaTokenHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iToken := self.Args("token").String()
	if (offset > 0) && helper.IsHas(iToken) {
		_Batch, _error := model.GetBatchsViaToken(offset, limit, iToken, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Batch)
	}
	herr.Message = "Can't get to the GetBatchsViaToken's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBatchsViaTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iTimestamp := self.Args("timestamp").MustInt()
	if (offset > 0) && helper.IsHas(iTimestamp) {
		_Batch, _error := model.GetBatchsViaTimestamp(offset, limit, iTimestamp, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Batch)
	}
	herr.Message = "Can't get to the GetBatchsViaTimestamp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBatchsViaBatchHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iBatch := self.Args("batch").Bytes()
	if (offset > 0) && helper.IsHas(iBatch) {
		_Batch, _error := model.GetBatchsViaBatch(offset, limit, iBatch, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Batch)
	}
	herr.Message = "Can't get to the GetBatchsViaBatch's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBatchsByBidAndTokenAndTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBid := self.Args("bid").MustInt64()
	iToken := self.Args("token").String()
	iTimestamp := self.Args("timestamp").MustInt()

	if helper.IsHas(iBid) {
		_Batch, _error := model.GetBatchsByBidAndTokenAndTimestamp(offset, limit, iBid,iToken,iTimestamp)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Batch)
	}
	herr.Message = "Can't get to the GetBatchsByBidAndTokenAndTimestamp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBatchsByBidAndTokenAndBatchHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBid := self.Args("bid").MustInt64()
	iToken := self.Args("token").String()
	iBatch := self.Args("batch").Bytes()

	if helper.IsHas(iBid) {
		_Batch, _error := model.GetBatchsByBidAndTokenAndBatch(offset, limit, iBid,iToken,iBatch)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Batch)
	}
	herr.Message = "Can't get to the GetBatchsByBidAndTokenAndBatch's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBatchsByBidAndTimestampAndBatchHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBid := self.Args("bid").MustInt64()
	iTimestamp := self.Args("timestamp").MustInt()
	iBatch := self.Args("batch").Bytes()

	if helper.IsHas(iBid) {
		_Batch, _error := model.GetBatchsByBidAndTimestampAndBatch(offset, limit, iBid,iTimestamp,iBatch)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Batch)
	}
	herr.Message = "Can't get to the GetBatchsByBidAndTimestampAndBatch's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBatchsByTokenAndTimestampAndBatchHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iToken := self.Args("token").String()
	iTimestamp := self.Args("timestamp").MustInt()
	iBatch := self.Args("batch").Bytes()

	if helper.IsHas(iToken) {
		_Batch, _error := model.GetBatchsByTokenAndTimestampAndBatch(offset, limit, iToken,iTimestamp,iBatch)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Batch)
	}
	herr.Message = "Can't get to the GetBatchsByTokenAndTimestampAndBatch's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBatchsByBidAndTokenHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBid := self.Args("bid").MustInt64()
	iToken := self.Args("token").String()

	if helper.IsHas(iBid) {
		_Batch, _error := model.GetBatchsByBidAndToken(offset, limit, iBid,iToken)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Batch)
	}
	herr.Message = "Can't get to the GetBatchsByBidAndToken's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBatchsByBidAndTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBid := self.Args("bid").MustInt64()
	iTimestamp := self.Args("timestamp").MustInt()

	if helper.IsHas(iBid) {
		_Batch, _error := model.GetBatchsByBidAndTimestamp(offset, limit, iBid,iTimestamp)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Batch)
	}
	herr.Message = "Can't get to the GetBatchsByBidAndTimestamp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBatchsByBidAndBatchHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBid := self.Args("bid").MustInt64()
	iBatch := self.Args("batch").Bytes()

	if helper.IsHas(iBid) {
		_Batch, _error := model.GetBatchsByBidAndBatch(offset, limit, iBid,iBatch)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Batch)
	}
	herr.Message = "Can't get to the GetBatchsByBidAndBatch's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBatchsByTokenAndTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iToken := self.Args("token").String()
	iTimestamp := self.Args("timestamp").MustInt()

	if helper.IsHas(iToken) {
		_Batch, _error := model.GetBatchsByTokenAndTimestamp(offset, limit, iToken,iTimestamp)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Batch)
	}
	herr.Message = "Can't get to the GetBatchsByTokenAndTimestamp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBatchsByTokenAndBatchHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iToken := self.Args("token").String()
	iBatch := self.Args("batch").Bytes()

	if helper.IsHas(iToken) {
		_Batch, _error := model.GetBatchsByTokenAndBatch(offset, limit, iToken,iBatch)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Batch)
	}
	herr.Message = "Can't get to the GetBatchsByTokenAndBatch's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBatchsByTimestampAndBatchHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTimestamp := self.Args("timestamp").MustInt()
	iBatch := self.Args("batch").Bytes()

	if helper.IsHas(iTimestamp) {
		_Batch, _error := model.GetBatchsByTimestampAndBatch(offset, limit, iTimestamp,iBatch)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Batch)
	}
	herr.Message = "Can't get to the GetBatchsByTimestampAndBatch's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBatchsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	if (offset > 0) && (len(field) > 0) {
		_Batch, _error := model.GetBatchs(offset, limit, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Batch)
	}
	herr.Message = "Can't get to the GetBatchs' args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasBatchViaBidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iBid := self.Args("bid").MustInt64()
	if helper.IsHas(iBid) {
		_Batch := model.HasBatchViaBid(iBid)
		var m = map[string]interface{}{}
		m["batch"] = _Batch
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasBatchViaBid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasBatchViaTokenHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iToken := self.Args("token").String()
	if helper.IsHas(iToken) {
		_Batch := model.HasBatchViaToken(iToken)
		var m = map[string]interface{}{}
		m["batch"] = _Batch
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasBatchViaToken's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasBatchViaTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iTimestamp := self.Args("timestamp").MustInt()
	if helper.IsHas(iTimestamp) {
		_Batch := model.HasBatchViaTimestamp(iTimestamp)
		var m = map[string]interface{}{}
		m["batch"] = _Batch
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasBatchViaTimestamp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasBatchViaBatchHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iBatch := self.Args("batch").Bytes()
	if helper.IsHas(iBatch) {
		_Batch := model.HasBatchViaBatch(iBatch)
		var m = map[string]interface{}{}
		m["batch"] = _Batch
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasBatchViaBatch's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBatchViaBidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iBid := self.Args("bid").MustInt64()
	if helper.IsHas(iBid) {
		_Batch, _error := model.GetBatchViaBid(iBid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Batch)
	}
	herr.Message = "Can't get to the GetBatchViaBid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBatchViaTokenHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iToken := self.Args("token").String()
	if helper.IsHas(iToken) {
		_Batch, _error := model.GetBatchViaToken(iToken)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Batch)
	}
	herr.Message = "Can't get to the GetBatchViaToken's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBatchViaTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iTimestamp := self.Args("timestamp").MustInt()
	if helper.IsHas(iTimestamp) {
		_Batch, _error := model.GetBatchViaTimestamp(iTimestamp)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Batch)
	}
	herr.Message = "Can't get to the GetBatchViaTimestamp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBatchViaBatchHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iBatch := self.Args("batch").Bytes()
	if helper.IsHas(iBatch) {
		_Batch, _error := model.GetBatchViaBatch(iBatch)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Batch)
	}
	herr.Message = "Can't get to the GetBatchViaBatch's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetBatchViaBidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Bid_ := self.Args("bid").MustInt64()
	if helper.IsHas(Bid_) {
		var iBatch model.Batch
		self.Bind(&iBatch)
		_Batch, _error := model.SetBatchViaBid(Bid_, &iBatch)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Batch)
	}
	herr.Message = "Can't get to the SetBatchViaBid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetBatchViaTokenHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Token_ := self.Args("token").String()
	if helper.IsHas(Token_) {
		var iBatch model.Batch
		self.Bind(&iBatch)
		_Batch, _error := model.SetBatchViaToken(Token_, &iBatch)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Batch)
	}
	herr.Message = "Can't get to the SetBatchViaToken's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetBatchViaTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Timestamp_ := self.Args("timestamp").MustInt()
	if helper.IsHas(Timestamp_) {
		var iBatch model.Batch
		self.Bind(&iBatch)
		_Batch, _error := model.SetBatchViaTimestamp(Timestamp_, &iBatch)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Batch)
	}
	herr.Message = "Can't get to the SetBatchViaTimestamp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetBatchViaBatchHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Batch_ := self.Args("batch").Bytes()
	if helper.IsHas(Batch_) {
		var iBatch model.Batch
		self.Bind(&iBatch)
		_Batch, _error := model.SetBatchViaBatch(Batch_, &iBatch)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Batch)
	}
	herr.Message = "Can't get to the SetBatchViaBatch's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostAddBatchHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Bid_ := self.Args("bid").MustInt64()
	Token_ := self.Args("token").String()
	Timestamp_ := self.Args("timestamp").MustInt()
	Batch_ := self.Args("batch").Bytes()

	if helper.IsHas(Bid_) {
		_error := model.AddBatch(Bid_,Token_,Timestamp_,Batch_)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		herr.Message = "StatusOK"
		herr.Status = macross.StatusOK
		return self.JSON(herr)
	}
	herr.Message = "Can't get to the AddBatch's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostBatchHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iBatch model.Batch
	self.Bind(&iBatch)
	_int64, _error := model.PostBatch(&iBatch)
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

func PutBatchHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iBatch model.Batch
	self.Bind(&iBatch)
	_int64, _error := model.PutBatch(&iBatch)
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

func PutBatchViaBidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Bid_ := self.Args("bid").MustInt64()
	var iBatch model.Batch
	self.Bind(&iBatch)
	_int64, _error := model.PutBatchViaBid(Bid_, &iBatch)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutBatchViaTokenHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Token_ := self.Args("token").String()
	var iBatch model.Batch
	self.Bind(&iBatch)
	_int64, _error := model.PutBatchViaToken(Token_, &iBatch)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutBatchViaTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Timestamp_ := self.Args("timestamp").MustInt()
	var iBatch model.Batch
	self.Bind(&iBatch)
	_int64, _error := model.PutBatchViaTimestamp(Timestamp_, &iBatch)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutBatchViaBatchHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Batch_ := self.Args("batch").Bytes()
	var iBatch model.Batch
	self.Bind(&iBatch)
	_int64, _error := model.PutBatchViaBatch(Batch_, &iBatch)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateBatchViaBidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Bid_ := self.Args("bid").MustInt64()
	var iBatch model.Batch
	self.Bind(&iBatch)
	var iMap = helper.StructToMap(iBatch)
	_error := model.UpdateBatchViaBid(Bid_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateBatchViaTokenHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Token_ := self.Args("token").String()
	var iBatch model.Batch
	self.Bind(&iBatch)
	var iMap = helper.StructToMap(iBatch)
	_error := model.UpdateBatchViaToken(Token_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateBatchViaTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Timestamp_ := self.Args("timestamp").MustInt()
	var iBatch model.Batch
	self.Bind(&iBatch)
	var iMap = helper.StructToMap(iBatch)
	_error := model.UpdateBatchViaTimestamp(Timestamp_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateBatchViaBatchHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Batch_ := self.Args("batch").Bytes()
	var iBatch model.Batch
	self.Bind(&iBatch)
	var iMap = helper.StructToMap(iBatch)
	_error := model.UpdateBatchViaBatch(Batch_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteBatchHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Bid_ := self.Args("bid").MustInt64()
	_int64, _error := model.DeleteBatch(Bid_)
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

func DeleteBatchViaBidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Bid_ := self.Args("bid").MustInt64()
	_error := model.DeleteBatchViaBid(Bid_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteBatchViaTokenHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Token_ := self.Args("token").String()
	_error := model.DeleteBatchViaToken(Token_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteBatchViaTimestampHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Timestamp_ := self.Args("timestamp").MustInt()
	_error := model.DeleteBatchViaTimestamp(Timestamp_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteBatchViaBatchHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Batch_ := self.Args("batch").Bytes()
	_error := model.DeleteBatchViaBatch(Batch_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}
