package model

type History struct {
	Uid       int64 `xorm:"not null pk default 0 INT(11)"`
	Nid       int   `xorm:"not null pk default 0 index INT(10)"`
	Timestamp int   `xorm:"not null default 0 INT(11)"`
}

// GetHistoriesCount Historys' Count
func GetHistoriesCount(offset int, limit int) (int64, error) {
	total, err := Engine.Limit(limit, offset).Count(&History{})
	return total, err
}

// GetHistoryCountViaUid Get History via Uid
func GetHistoryCountViaUid(iUid int64) int64 {
	n, _ := Engine.Where("uid = ?", iUid).Count(&History{Uid: iUid})
	return n
}

// GetHistoryCountViaNid Get History via Nid
func GetHistoryCountViaNid(iNid int) int64 {
	n, _ := Engine.Where("nid = ?", iNid).Count(&History{Nid: iNid})
	return n
}

// GetHistoryCountViaTimestamp Get History via Timestamp
func GetHistoryCountViaTimestamp(iTimestamp int) int64 {
	n, _ := Engine.Where("timestamp = ?", iTimestamp).Count(&History{Timestamp: iTimestamp})
	return n
}

// GetHistoriesByUidAndNidAndTimestamp Get Histories via UidAndNidAndTimestamp
func GetHistoriesByUidAndNidAndTimestamp(offset int, limit int, Uid_ int64, Nid_ int, Timestamp_ int) (*[]*History, error) {
	var _History = new([]*History)
	err := Engine.Table("history").Where("uid = ? and nid = ? and timestamp = ?", Uid_, Nid_, Timestamp_).Limit(limit, offset).Find(_History)
	return _History, err
}

// GetHistoriesByUidAndNid Get Histories via UidAndNid
func GetHistoriesByUidAndNid(offset int, limit int, Uid_ int64, Nid_ int) (*[]*History, error) {
	var _History = new([]*History)
	err := Engine.Table("history").Where("uid = ? and nid = ?", Uid_, Nid_).Limit(limit, offset).Find(_History)
	return _History, err
}

// GetHistoriesByUidAndTimestamp Get Histories via UidAndTimestamp
func GetHistoriesByUidAndTimestamp(offset int, limit int, Uid_ int64, Timestamp_ int) (*[]*History, error) {
	var _History = new([]*History)
	err := Engine.Table("history").Where("uid = ? and timestamp = ?", Uid_, Timestamp_).Limit(limit, offset).Find(_History)
	return _History, err
}

// GetHistoriesByNidAndTimestamp Get Histories via NidAndTimestamp
func GetHistoriesByNidAndTimestamp(offset int, limit int, Nid_ int, Timestamp_ int) (*[]*History, error) {
	var _History = new([]*History)
	err := Engine.Table("history").Where("nid = ? and timestamp = ?", Nid_, Timestamp_).Limit(limit, offset).Find(_History)
	return _History, err
}

// GetHistories Get Histories via field
func GetHistories(offset int, limit int, field string) (*[]*History, error) {
	var _History = new([]*History)
	err := Engine.Table("history").Limit(limit, offset).Desc(field).Find(_History)
	return _History, err
}

// GetHistoriesViaUid Get Historys via Uid
func GetHistoriesViaUid(offset int, limit int, Uid_ int64, field string) (*[]*History, error) {
	var _History = new([]*History)
	err := Engine.Table("history").Where("uid = ?", Uid_).Limit(limit, offset).Desc(field).Find(_History)
	return _History, err
}

// GetHistoriesViaNid Get Historys via Nid
func GetHistoriesViaNid(offset int, limit int, Nid_ int, field string) (*[]*History, error) {
	var _History = new([]*History)
	err := Engine.Table("history").Where("nid = ?", Nid_).Limit(limit, offset).Desc(field).Find(_History)
	return _History, err
}

// GetHistoriesViaTimestamp Get Historys via Timestamp
func GetHistoriesViaTimestamp(offset int, limit int, Timestamp_ int, field string) (*[]*History, error) {
	var _History = new([]*History)
	err := Engine.Table("history").Where("timestamp = ?", Timestamp_).Limit(limit, offset).Desc(field).Find(_History)
	return _History, err
}

// HasHistoryViaUid Has History via Uid
func HasHistoryViaUid(iUid int64) bool {
	if has, err := Engine.Where("uid = ?", iUid).Get(new(History)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasHistoryViaNid Has History via Nid
func HasHistoryViaNid(iNid int) bool {
	if has, err := Engine.Where("nid = ?", iNid).Get(new(History)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasHistoryViaTimestamp Has History via Timestamp
func HasHistoryViaTimestamp(iTimestamp int) bool {
	if has, err := Engine.Where("timestamp = ?", iTimestamp).Get(new(History)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// GetHistoryViaUid Get History via Uid
func GetHistoryViaUid(iUid int64) (*History, error) {
	var _History = &History{Uid: iUid}
	has, err := Engine.Get(_History)
	if has {
		return _History, err
	} else {
		return nil, err
	}
}

// GetHistoryViaNid Get History via Nid
func GetHistoryViaNid(iNid int) (*History, error) {
	var _History = &History{Nid: iNid}
	has, err := Engine.Get(_History)
	if has {
		return _History, err
	} else {
		return nil, err
	}
}

// GetHistoryViaTimestamp Get History via Timestamp
func GetHistoryViaTimestamp(iTimestamp int) (*History, error) {
	var _History = &History{Timestamp: iTimestamp}
	has, err := Engine.Get(_History)
	if has {
		return _History, err
	} else {
		return nil, err
	}
}

// SetHistoryViaUid Set History via Uid
func SetHistoryViaUid(iUid int64, history *History) (int64, error) {
	history.Uid = iUid
	return Engine.Insert(history)
}

// SetHistoryViaNid Set History via Nid
func SetHistoryViaNid(iNid int, history *History) (int64, error) {
	history.Nid = iNid
	return Engine.Insert(history)
}

// SetHistoryViaTimestamp Set History via Timestamp
func SetHistoryViaTimestamp(iTimestamp int, history *History) (int64, error) {
	history.Timestamp = iTimestamp
	return Engine.Insert(history)
}

// AddHistory Add History via all columns
func AddHistory(iUid int64, iNid int, iTimestamp int) error {
	_History := &History{Uid: iUid, Nid: iNid, Timestamp: iTimestamp}
	if _, err := Engine.Insert(_History); err != nil {
		return err
	} else {
		return nil
	}
}

// PostHistory Post History via iHistory
func PostHistory(iHistory *History) (int64, error) {
	_, err := Engine.Insert(iHistory)
	return iHistory.Uid, err
}

// PutHistory Put History
func PutHistory(iHistory *History) (int64, error) {
	_, err := Engine.Id(iHistory.Uid).Update(iHistory)
	return iHistory.Uid, err
}

// PutHistoryViaUid Put History via Uid
func PutHistoryViaUid(Uid_ int64, iHistory *History) (int64, error) {
	row, err := Engine.Update(iHistory, &History{Uid: Uid_})
	return row, err
}

// PutHistoryViaNid Put History via Nid
func PutHistoryViaNid(Nid_ int, iHistory *History) (int64, error) {
	row, err := Engine.Update(iHistory, &History{Nid: Nid_})
	return row, err
}

// PutHistoryViaTimestamp Put History via Timestamp
func PutHistoryViaTimestamp(Timestamp_ int, iHistory *History) (int64, error) {
	row, err := Engine.Update(iHistory, &History{Timestamp: Timestamp_})
	return row, err
}

// UpdateHistoryViaUid via map[string]interface{}{}
func UpdateHistoryViaUid(iUid int64, iHistoryMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(History)).Where("uid = ?", iUid).Update(iHistoryMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateHistoryViaNid via map[string]interface{}{}
func UpdateHistoryViaNid(iNid int, iHistoryMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(History)).Where("nid = ?", iNid).Update(iHistoryMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateHistoryViaTimestamp via map[string]interface{}{}
func UpdateHistoryViaTimestamp(iTimestamp int, iHistoryMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(History)).Where("timestamp = ?", iTimestamp).Update(iHistoryMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// DeleteHistory Delete History
func DeleteHistory(iUid int64) (int64, error) {
	row, err := Engine.Id(iUid).Delete(new(History))
	return row, err
}

// DeleteHistoryViaUid Delete History via Uid
func DeleteHistoryViaUid(iUid int64) (err error) {
	var has bool
	var _History = &History{Uid: iUid}
	if has, err = Engine.Get(_History); (has == true) && (err == nil) {
		if row, err := Engine.Where("uid = ?", iUid).Delete(new(History)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteHistoryViaNid Delete History via Nid
func DeleteHistoryViaNid(iNid int) (err error) {
	var has bool
	var _History = &History{Nid: iNid}
	if has, err = Engine.Get(_History); (has == true) && (err == nil) {
		if row, err := Engine.Where("nid = ?", iNid).Delete(new(History)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteHistoryViaTimestamp Delete History via Timestamp
func DeleteHistoryViaTimestamp(iTimestamp int) (err error) {
	var has bool
	var _History = &History{Timestamp: iTimestamp}
	if has, err = Engine.Get(_History); (has == true) && (err == nil) {
		if row, err := Engine.Where("timestamp = ?", iTimestamp).Delete(new(History)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}
