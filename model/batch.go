package model

type Batch struct {
	Bid       int64  `xorm:"not null pk INT(10)"`
	Token     string `xorm:"not null index VARCHAR(64)"`
	Timestamp int    `xorm:"not null INT(11)"`
	Batch     []byte `xorm:"LONGBLOB"`
}

// GetBatchsCount Batchs' Count
func GetBatchsCount(offset int, limit int) (int64, error) {
	total, err := Engine.Limit(limit, offset).Count(&Batch{})
	return total, err
}

// GetBatchCountViaBid Get Batch via Bid
func GetBatchCountViaBid(iBid int64) int64 {
	n, _ := Engine.Where("bid = ?", iBid).Count(&Batch{Bid: iBid})
	return n
}

// GetBatchCountViaToken Get Batch via Token
func GetBatchCountViaToken(iToken string) int64 {
	n, _ := Engine.Where("token = ?", iToken).Count(&Batch{Token: iToken})
	return n
}

// GetBatchCountViaTimestamp Get Batch via Timestamp
func GetBatchCountViaTimestamp(iTimestamp int) int64 {
	n, _ := Engine.Where("timestamp = ?", iTimestamp).Count(&Batch{Timestamp: iTimestamp})
	return n
}

// GetBatchCountViaBatch Get Batch via Batch
func GetBatchCountViaBatch(iBatch []byte) int64 {
	n, _ := Engine.Where("batch = ?", iBatch).Count(&Batch{Batch: iBatch})
	return n
}

// GetBatchsByBidAndTokenAndTimestamp Get Batchs via BidAndTokenAndTimestamp
func GetBatchsByBidAndTokenAndTimestamp(offset int, limit int, Bid_ int64, Token_ string, Timestamp_ int) (*[]*Batch, error) {
	var _Batch = new([]*Batch)
	err := Engine.Table("batch").Where("bid = ? and token = ? and timestamp = ?", Bid_, Token_, Timestamp_).Limit(limit, offset).Find(_Batch)
	return _Batch, err
}

// GetBatchsByBidAndTokenAndBatch Get Batchs via BidAndTokenAndBatch
func GetBatchsByBidAndTokenAndBatch(offset int, limit int, Bid_ int64, Token_ string, Batch_ []byte) (*[]*Batch, error) {
	var _Batch = new([]*Batch)
	err := Engine.Table("batch").Where("bid = ? and token = ? and batch = ?", Bid_, Token_, Batch_).Limit(limit, offset).Find(_Batch)
	return _Batch, err
}

// GetBatchsByBidAndTimestampAndBatch Get Batchs via BidAndTimestampAndBatch
func GetBatchsByBidAndTimestampAndBatch(offset int, limit int, Bid_ int64, Timestamp_ int, Batch_ []byte) (*[]*Batch, error) {
	var _Batch = new([]*Batch)
	err := Engine.Table("batch").Where("bid = ? and timestamp = ? and batch = ?", Bid_, Timestamp_, Batch_).Limit(limit, offset).Find(_Batch)
	return _Batch, err
}

// GetBatchsByTokenAndTimestampAndBatch Get Batchs via TokenAndTimestampAndBatch
func GetBatchsByTokenAndTimestampAndBatch(offset int, limit int, Token_ string, Timestamp_ int, Batch_ []byte) (*[]*Batch, error) {
	var _Batch = new([]*Batch)
	err := Engine.Table("batch").Where("token = ? and timestamp = ? and batch = ?", Token_, Timestamp_, Batch_).Limit(limit, offset).Find(_Batch)
	return _Batch, err
}

// GetBatchsByBidAndToken Get Batchs via BidAndToken
func GetBatchsByBidAndToken(offset int, limit int, Bid_ int64, Token_ string) (*[]*Batch, error) {
	var _Batch = new([]*Batch)
	err := Engine.Table("batch").Where("bid = ? and token = ?", Bid_, Token_).Limit(limit, offset).Find(_Batch)
	return _Batch, err
}

// GetBatchsByBidAndTimestamp Get Batchs via BidAndTimestamp
func GetBatchsByBidAndTimestamp(offset int, limit int, Bid_ int64, Timestamp_ int) (*[]*Batch, error) {
	var _Batch = new([]*Batch)
	err := Engine.Table("batch").Where("bid = ? and timestamp = ?", Bid_, Timestamp_).Limit(limit, offset).Find(_Batch)
	return _Batch, err
}

// GetBatchsByBidAndBatch Get Batchs via BidAndBatch
func GetBatchsByBidAndBatch(offset int, limit int, Bid_ int64, Batch_ []byte) (*[]*Batch, error) {
	var _Batch = new([]*Batch)
	err := Engine.Table("batch").Where("bid = ? and batch = ?", Bid_, Batch_).Limit(limit, offset).Find(_Batch)
	return _Batch, err
}

// GetBatchsByTokenAndTimestamp Get Batchs via TokenAndTimestamp
func GetBatchsByTokenAndTimestamp(offset int, limit int, Token_ string, Timestamp_ int) (*[]*Batch, error) {
	var _Batch = new([]*Batch)
	err := Engine.Table("batch").Where("token = ? and timestamp = ?", Token_, Timestamp_).Limit(limit, offset).Find(_Batch)
	return _Batch, err
}

// GetBatchsByTokenAndBatch Get Batchs via TokenAndBatch
func GetBatchsByTokenAndBatch(offset int, limit int, Token_ string, Batch_ []byte) (*[]*Batch, error) {
	var _Batch = new([]*Batch)
	err := Engine.Table("batch").Where("token = ? and batch = ?", Token_, Batch_).Limit(limit, offset).Find(_Batch)
	return _Batch, err
}

// GetBatchsByTimestampAndBatch Get Batchs via TimestampAndBatch
func GetBatchsByTimestampAndBatch(offset int, limit int, Timestamp_ int, Batch_ []byte) (*[]*Batch, error) {
	var _Batch = new([]*Batch)
	err := Engine.Table("batch").Where("timestamp = ? and batch = ?", Timestamp_, Batch_).Limit(limit, offset).Find(_Batch)
	return _Batch, err
}

// GetBatchs Get Batchs via field
func GetBatchs(offset int, limit int, field string) (*[]*Batch, error) {
	var _Batch = new([]*Batch)
	err := Engine.Table("batch").Limit(limit, offset).Desc(field).Find(_Batch)
	return _Batch, err
}

// GetBatchsViaBid Get Batchs via Bid
func GetBatchsViaBid(offset int, limit int, Bid_ int64, field string) (*[]*Batch, error) {
	var _Batch = new([]*Batch)
	err := Engine.Table("batch").Where("bid = ?", Bid_).Limit(limit, offset).Desc(field).Find(_Batch)
	return _Batch, err
}

// GetBatchsViaToken Get Batchs via Token
func GetBatchsViaToken(offset int, limit int, Token_ string, field string) (*[]*Batch, error) {
	var _Batch = new([]*Batch)
	err := Engine.Table("batch").Where("token = ?", Token_).Limit(limit, offset).Desc(field).Find(_Batch)
	return _Batch, err
}

// GetBatchsViaTimestamp Get Batchs via Timestamp
func GetBatchsViaTimestamp(offset int, limit int, Timestamp_ int, field string) (*[]*Batch, error) {
	var _Batch = new([]*Batch)
	err := Engine.Table("batch").Where("timestamp = ?", Timestamp_).Limit(limit, offset).Desc(field).Find(_Batch)
	return _Batch, err
}

// GetBatchsViaBatch Get Batchs via Batch
func GetBatchsViaBatch(offset int, limit int, Batch_ []byte, field string) (*[]*Batch, error) {
	var _Batch = new([]*Batch)
	err := Engine.Table("batch").Where("batch = ?", Batch_).Limit(limit, offset).Desc(field).Find(_Batch)
	return _Batch, err
}

// HasBatchViaBid Has Batch via Bid
func HasBatchViaBid(iBid int64) bool {
	if has, err := Engine.Where("bid = ?", iBid).Get(new(Batch)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasBatchViaToken Has Batch via Token
func HasBatchViaToken(iToken string) bool {
	if has, err := Engine.Where("token = ?", iToken).Get(new(Batch)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasBatchViaTimestamp Has Batch via Timestamp
func HasBatchViaTimestamp(iTimestamp int) bool {
	if has, err := Engine.Where("timestamp = ?", iTimestamp).Get(new(Batch)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasBatchViaBatch Has Batch via Batch
func HasBatchViaBatch(iBatch []byte) bool {
	if has, err := Engine.Where("batch = ?", iBatch).Get(new(Batch)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// GetBatchViaBid Get Batch via Bid
func GetBatchViaBid(iBid int64) (*Batch, error) {
	var _Batch = &Batch{Bid: iBid}
	has, err := Engine.Get(_Batch)
	if has {
		return _Batch, err
	} else {
		return nil, err
	}
}

// GetBatchViaToken Get Batch via Token
func GetBatchViaToken(iToken string) (*Batch, error) {
	var _Batch = &Batch{Token: iToken}
	has, err := Engine.Get(_Batch)
	if has {
		return _Batch, err
	} else {
		return nil, err
	}
}

// GetBatchViaTimestamp Get Batch via Timestamp
func GetBatchViaTimestamp(iTimestamp int) (*Batch, error) {
	var _Batch = &Batch{Timestamp: iTimestamp}
	has, err := Engine.Get(_Batch)
	if has {
		return _Batch, err
	} else {
		return nil, err
	}
}

// GetBatchViaBatch Get Batch via Batch
func GetBatchViaBatch(iBatch []byte) (*Batch, error) {
	var _Batch = &Batch{Batch: iBatch}
	has, err := Engine.Get(_Batch)
	if has {
		return _Batch, err
	} else {
		return nil, err
	}
}

// SetBatchViaBid Set Batch via Bid
func SetBatchViaBid(iBid int64, batch *Batch) (int64, error) {
	batch.Bid = iBid
	return Engine.Insert(batch)
}

// SetBatchViaToken Set Batch via Token
func SetBatchViaToken(iToken string, batch *Batch) (int64, error) {
	batch.Token = iToken
	return Engine.Insert(batch)
}

// SetBatchViaTimestamp Set Batch via Timestamp
func SetBatchViaTimestamp(iTimestamp int, batch *Batch) (int64, error) {
	batch.Timestamp = iTimestamp
	return Engine.Insert(batch)
}

// SetBatchViaBatch Set Batch via Batch
func SetBatchViaBatch(iBatch []byte, batch *Batch) (int64, error) {
	batch.Batch = iBatch
	return Engine.Insert(batch)
}

// AddBatch Add Batch via all columns
func AddBatch(iBid int64, iToken string, iTimestamp int, iBatch []byte) error {
	_Batch := &Batch{Bid: iBid, Token: iToken, Timestamp: iTimestamp, Batch: iBatch}
	if _, err := Engine.Insert(_Batch); err != nil {
		return err
	} else {
		return nil
	}
}

// PostBatch Post Batch via iBatch
func PostBatch(iBatch *Batch) (int64, error) {
	_, err := Engine.Insert(iBatch)
	return iBatch.Bid, err
}

// PutBatch Put Batch
func PutBatch(iBatch *Batch) (int64, error) {
	_, err := Engine.Id(iBatch.Bid).Update(iBatch)
	return iBatch.Bid, err
}

// PutBatchViaBid Put Batch via Bid
func PutBatchViaBid(Bid_ int64, iBatch *Batch) (int64, error) {
	row, err := Engine.Update(iBatch, &Batch{Bid: Bid_})
	return row, err
}

// PutBatchViaToken Put Batch via Token
func PutBatchViaToken(Token_ string, iBatch *Batch) (int64, error) {
	row, err := Engine.Update(iBatch, &Batch{Token: Token_})
	return row, err
}

// PutBatchViaTimestamp Put Batch via Timestamp
func PutBatchViaTimestamp(Timestamp_ int, iBatch *Batch) (int64, error) {
	row, err := Engine.Update(iBatch, &Batch{Timestamp: Timestamp_})
	return row, err
}

// PutBatchViaBatch Put Batch via Batch
func PutBatchViaBatch(Batch_ []byte, iBatch *Batch) (int64, error) {
	row, err := Engine.Update(iBatch, &Batch{Batch: Batch_})
	return row, err
}

// UpdateBatchViaBid via map[string]interface{}{}
func UpdateBatchViaBid(iBid int64, iBatchMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Batch)).Where("bid = ?", iBid).Update(iBatchMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateBatchViaToken via map[string]interface{}{}
func UpdateBatchViaToken(iToken string, iBatchMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Batch)).Where("token = ?", iToken).Update(iBatchMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateBatchViaTimestamp via map[string]interface{}{}
func UpdateBatchViaTimestamp(iTimestamp int, iBatchMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Batch)).Where("timestamp = ?", iTimestamp).Update(iBatchMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateBatchViaBatch via map[string]interface{}{}
func UpdateBatchViaBatch(iBatch []byte, iBatchMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Batch)).Where("batch = ?", iBatch).Update(iBatchMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// DeleteBatch Delete Batch
func DeleteBatch(iBid int64) (int64, error) {
	row, err := Engine.Id(iBid).Delete(new(Batch))
	return row, err
}

// DeleteBatchViaBid Delete Batch via Bid
func DeleteBatchViaBid(iBid int64) (err error) {
	var has bool
	var _Batch = &Batch{Bid: iBid}
	if has, err = Engine.Get(_Batch); (has == true) && (err == nil) {
		if row, err := Engine.Where("bid = ?", iBid).Delete(new(Batch)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteBatchViaToken Delete Batch via Token
func DeleteBatchViaToken(iToken string) (err error) {
	var has bool
	var _Batch = &Batch{Token: iToken}
	if has, err = Engine.Get(_Batch); (has == true) && (err == nil) {
		if row, err := Engine.Where("token = ?", iToken).Delete(new(Batch)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteBatchViaTimestamp Delete Batch via Timestamp
func DeleteBatchViaTimestamp(iTimestamp int) (err error) {
	var has bool
	var _Batch = &Batch{Timestamp: iTimestamp}
	if has, err = Engine.Get(_Batch); (has == true) && (err == nil) {
		if row, err := Engine.Where("timestamp = ?", iTimestamp).Delete(new(Batch)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteBatchViaBatch Delete Batch via Batch
func DeleteBatchViaBatch(iBatch []byte) (err error) {
	var has bool
	var _Batch = &Batch{Batch: iBatch}
	if has, err = Engine.Get(_Batch); (has == true) && (err == nil) {
		if row, err := Engine.Where("batch = ?", iBatch).Delete(new(Batch)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}
