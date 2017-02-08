package model

type Flood struct {
	Fid        int64  `xorm:"not null pk autoincr INT(11)"`
	Event      string `xorm:"not null default '' index(allow) VARCHAR(64)"`
	Identifier string `xorm:"not null default '' index(allow) VARCHAR(128)"`
	Timestamp  int    `xorm:"not null default 0 index(allow) INT(11)"`
	Expiration int    `xorm:"not null default 0 index INT(11)"`
}

// GetFloodsCount Floods' Count
func GetFloodsCount(offset int, limit int) (int64, error) {
	total, err := Engine.Limit(limit, offset).Count(&Flood{})
	return total, err
}

// GetFloodCountViaFid Get Flood via Fid
func GetFloodCountViaFid(iFid int64) int64 {
	n, _ := Engine.Where("fid = ?", iFid).Count(&Flood{Fid: iFid})
	return n
}

// GetFloodCountViaEvent Get Flood via Event
func GetFloodCountViaEvent(iEvent string) int64 {
	n, _ := Engine.Where("event = ?", iEvent).Count(&Flood{Event: iEvent})
	return n
}

// GetFloodCountViaIdentifier Get Flood via Identifier
func GetFloodCountViaIdentifier(iIdentifier string) int64 {
	n, _ := Engine.Where("identifier = ?", iIdentifier).Count(&Flood{Identifier: iIdentifier})
	return n
}

// GetFloodCountViaTimestamp Get Flood via Timestamp
func GetFloodCountViaTimestamp(iTimestamp int) int64 {
	n, _ := Engine.Where("timestamp = ?", iTimestamp).Count(&Flood{Timestamp: iTimestamp})
	return n
}

// GetFloodCountViaExpiration Get Flood via Expiration
func GetFloodCountViaExpiration(iExpiration int) int64 {
	n, _ := Engine.Where("expiration = ?", iExpiration).Count(&Flood{Expiration: iExpiration})
	return n
}

// GetFloodsByFidAndEventAndIdentifier Get Floods via FidAndEventAndIdentifier
func GetFloodsByFidAndEventAndIdentifier(offset int, limit int, Fid_ int64, Event_ string, Identifier_ string) (*[]*Flood, error) {
	var _Flood = new([]*Flood)
	err := Engine.Table("flood").Where("fid = ? and event = ? and identifier = ?", Fid_, Event_, Identifier_).Limit(limit, offset).Find(_Flood)
	return _Flood, err
}

// GetFloodsByFidAndEventAndTimestamp Get Floods via FidAndEventAndTimestamp
func GetFloodsByFidAndEventAndTimestamp(offset int, limit int, Fid_ int64, Event_ string, Timestamp_ int) (*[]*Flood, error) {
	var _Flood = new([]*Flood)
	err := Engine.Table("flood").Where("fid = ? and event = ? and timestamp = ?", Fid_, Event_, Timestamp_).Limit(limit, offset).Find(_Flood)
	return _Flood, err
}

// GetFloodsByFidAndEventAndExpiration Get Floods via FidAndEventAndExpiration
func GetFloodsByFidAndEventAndExpiration(offset int, limit int, Fid_ int64, Event_ string, Expiration_ int) (*[]*Flood, error) {
	var _Flood = new([]*Flood)
	err := Engine.Table("flood").Where("fid = ? and event = ? and expiration = ?", Fid_, Event_, Expiration_).Limit(limit, offset).Find(_Flood)
	return _Flood, err
}

// GetFloodsByFidAndIdentifierAndTimestamp Get Floods via FidAndIdentifierAndTimestamp
func GetFloodsByFidAndIdentifierAndTimestamp(offset int, limit int, Fid_ int64, Identifier_ string, Timestamp_ int) (*[]*Flood, error) {
	var _Flood = new([]*Flood)
	err := Engine.Table("flood").Where("fid = ? and identifier = ? and timestamp = ?", Fid_, Identifier_, Timestamp_).Limit(limit, offset).Find(_Flood)
	return _Flood, err
}

// GetFloodsByFidAndIdentifierAndExpiration Get Floods via FidAndIdentifierAndExpiration
func GetFloodsByFidAndIdentifierAndExpiration(offset int, limit int, Fid_ int64, Identifier_ string, Expiration_ int) (*[]*Flood, error) {
	var _Flood = new([]*Flood)
	err := Engine.Table("flood").Where("fid = ? and identifier = ? and expiration = ?", Fid_, Identifier_, Expiration_).Limit(limit, offset).Find(_Flood)
	return _Flood, err
}

// GetFloodsByFidAndTimestampAndExpiration Get Floods via FidAndTimestampAndExpiration
func GetFloodsByFidAndTimestampAndExpiration(offset int, limit int, Fid_ int64, Timestamp_ int, Expiration_ int) (*[]*Flood, error) {
	var _Flood = new([]*Flood)
	err := Engine.Table("flood").Where("fid = ? and timestamp = ? and expiration = ?", Fid_, Timestamp_, Expiration_).Limit(limit, offset).Find(_Flood)
	return _Flood, err
}

// GetFloodsByEventAndIdentifierAndTimestamp Get Floods via EventAndIdentifierAndTimestamp
func GetFloodsByEventAndIdentifierAndTimestamp(offset int, limit int, Event_ string, Identifier_ string, Timestamp_ int) (*[]*Flood, error) {
	var _Flood = new([]*Flood)
	err := Engine.Table("flood").Where("event = ? and identifier = ? and timestamp = ?", Event_, Identifier_, Timestamp_).Limit(limit, offset).Find(_Flood)
	return _Flood, err
}

// GetFloodsByEventAndIdentifierAndExpiration Get Floods via EventAndIdentifierAndExpiration
func GetFloodsByEventAndIdentifierAndExpiration(offset int, limit int, Event_ string, Identifier_ string, Expiration_ int) (*[]*Flood, error) {
	var _Flood = new([]*Flood)
	err := Engine.Table("flood").Where("event = ? and identifier = ? and expiration = ?", Event_, Identifier_, Expiration_).Limit(limit, offset).Find(_Flood)
	return _Flood, err
}

// GetFloodsByEventAndTimestampAndExpiration Get Floods via EventAndTimestampAndExpiration
func GetFloodsByEventAndTimestampAndExpiration(offset int, limit int, Event_ string, Timestamp_ int, Expiration_ int) (*[]*Flood, error) {
	var _Flood = new([]*Flood)
	err := Engine.Table("flood").Where("event = ? and timestamp = ? and expiration = ?", Event_, Timestamp_, Expiration_).Limit(limit, offset).Find(_Flood)
	return _Flood, err
}

// GetFloodsByIdentifierAndTimestampAndExpiration Get Floods via IdentifierAndTimestampAndExpiration
func GetFloodsByIdentifierAndTimestampAndExpiration(offset int, limit int, Identifier_ string, Timestamp_ int, Expiration_ int) (*[]*Flood, error) {
	var _Flood = new([]*Flood)
	err := Engine.Table("flood").Where("identifier = ? and timestamp = ? and expiration = ?", Identifier_, Timestamp_, Expiration_).Limit(limit, offset).Find(_Flood)
	return _Flood, err
}

// GetFloodsByFidAndEvent Get Floods via FidAndEvent
func GetFloodsByFidAndEvent(offset int, limit int, Fid_ int64, Event_ string) (*[]*Flood, error) {
	var _Flood = new([]*Flood)
	err := Engine.Table("flood").Where("fid = ? and event = ?", Fid_, Event_).Limit(limit, offset).Find(_Flood)
	return _Flood, err
}

// GetFloodsByFidAndIdentifier Get Floods via FidAndIdentifier
func GetFloodsByFidAndIdentifier(offset int, limit int, Fid_ int64, Identifier_ string) (*[]*Flood, error) {
	var _Flood = new([]*Flood)
	err := Engine.Table("flood").Where("fid = ? and identifier = ?", Fid_, Identifier_).Limit(limit, offset).Find(_Flood)
	return _Flood, err
}

// GetFloodsByFidAndTimestamp Get Floods via FidAndTimestamp
func GetFloodsByFidAndTimestamp(offset int, limit int, Fid_ int64, Timestamp_ int) (*[]*Flood, error) {
	var _Flood = new([]*Flood)
	err := Engine.Table("flood").Where("fid = ? and timestamp = ?", Fid_, Timestamp_).Limit(limit, offset).Find(_Flood)
	return _Flood, err
}

// GetFloodsByFidAndExpiration Get Floods via FidAndExpiration
func GetFloodsByFidAndExpiration(offset int, limit int, Fid_ int64, Expiration_ int) (*[]*Flood, error) {
	var _Flood = new([]*Flood)
	err := Engine.Table("flood").Where("fid = ? and expiration = ?", Fid_, Expiration_).Limit(limit, offset).Find(_Flood)
	return _Flood, err
}

// GetFloodsByEventAndIdentifier Get Floods via EventAndIdentifier
func GetFloodsByEventAndIdentifier(offset int, limit int, Event_ string, Identifier_ string) (*[]*Flood, error) {
	var _Flood = new([]*Flood)
	err := Engine.Table("flood").Where("event = ? and identifier = ?", Event_, Identifier_).Limit(limit, offset).Find(_Flood)
	return _Flood, err
}

// GetFloodsByEventAndTimestamp Get Floods via EventAndTimestamp
func GetFloodsByEventAndTimestamp(offset int, limit int, Event_ string, Timestamp_ int) (*[]*Flood, error) {
	var _Flood = new([]*Flood)
	err := Engine.Table("flood").Where("event = ? and timestamp = ?", Event_, Timestamp_).Limit(limit, offset).Find(_Flood)
	return _Flood, err
}

// GetFloodsByEventAndExpiration Get Floods via EventAndExpiration
func GetFloodsByEventAndExpiration(offset int, limit int, Event_ string, Expiration_ int) (*[]*Flood, error) {
	var _Flood = new([]*Flood)
	err := Engine.Table("flood").Where("event = ? and expiration = ?", Event_, Expiration_).Limit(limit, offset).Find(_Flood)
	return _Flood, err
}

// GetFloodsByIdentifierAndTimestamp Get Floods via IdentifierAndTimestamp
func GetFloodsByIdentifierAndTimestamp(offset int, limit int, Identifier_ string, Timestamp_ int) (*[]*Flood, error) {
	var _Flood = new([]*Flood)
	err := Engine.Table("flood").Where("identifier = ? and timestamp = ?", Identifier_, Timestamp_).Limit(limit, offset).Find(_Flood)
	return _Flood, err
}

// GetFloodsByIdentifierAndExpiration Get Floods via IdentifierAndExpiration
func GetFloodsByIdentifierAndExpiration(offset int, limit int, Identifier_ string, Expiration_ int) (*[]*Flood, error) {
	var _Flood = new([]*Flood)
	err := Engine.Table("flood").Where("identifier = ? and expiration = ?", Identifier_, Expiration_).Limit(limit, offset).Find(_Flood)
	return _Flood, err
}

// GetFloodsByTimestampAndExpiration Get Floods via TimestampAndExpiration
func GetFloodsByTimestampAndExpiration(offset int, limit int, Timestamp_ int, Expiration_ int) (*[]*Flood, error) {
	var _Flood = new([]*Flood)
	err := Engine.Table("flood").Where("timestamp = ? and expiration = ?", Timestamp_, Expiration_).Limit(limit, offset).Find(_Flood)
	return _Flood, err
}

// GetFloods Get Floods via field
func GetFloods(offset int, limit int, field string) (*[]*Flood, error) {
	var _Flood = new([]*Flood)
	err := Engine.Table("flood").Limit(limit, offset).Desc(field).Find(_Flood)
	return _Flood, err
}

// GetFloodsViaFid Get Floods via Fid
func GetFloodsViaFid(offset int, limit int, Fid_ int64, field string) (*[]*Flood, error) {
	var _Flood = new([]*Flood)
	err := Engine.Table("flood").Where("fid = ?", Fid_).Limit(limit, offset).Desc(field).Find(_Flood)
	return _Flood, err
}

// GetFloodsViaEvent Get Floods via Event
func GetFloodsViaEvent(offset int, limit int, Event_ string, field string) (*[]*Flood, error) {
	var _Flood = new([]*Flood)
	err := Engine.Table("flood").Where("event = ?", Event_).Limit(limit, offset).Desc(field).Find(_Flood)
	return _Flood, err
}

// GetFloodsViaIdentifier Get Floods via Identifier
func GetFloodsViaIdentifier(offset int, limit int, Identifier_ string, field string) (*[]*Flood, error) {
	var _Flood = new([]*Flood)
	err := Engine.Table("flood").Where("identifier = ?", Identifier_).Limit(limit, offset).Desc(field).Find(_Flood)
	return _Flood, err
}

// GetFloodsViaTimestamp Get Floods via Timestamp
func GetFloodsViaTimestamp(offset int, limit int, Timestamp_ int, field string) (*[]*Flood, error) {
	var _Flood = new([]*Flood)
	err := Engine.Table("flood").Where("timestamp = ?", Timestamp_).Limit(limit, offset).Desc(field).Find(_Flood)
	return _Flood, err
}

// GetFloodsViaExpiration Get Floods via Expiration
func GetFloodsViaExpiration(offset int, limit int, Expiration_ int, field string) (*[]*Flood, error) {
	var _Flood = new([]*Flood)
	err := Engine.Table("flood").Where("expiration = ?", Expiration_).Limit(limit, offset).Desc(field).Find(_Flood)
	return _Flood, err
}

// HasFloodViaFid Has Flood via Fid
func HasFloodViaFid(iFid int64) bool {
	if has, err := Engine.Where("fid = ?", iFid).Get(new(Flood)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasFloodViaEvent Has Flood via Event
func HasFloodViaEvent(iEvent string) bool {
	if has, err := Engine.Where("event = ?", iEvent).Get(new(Flood)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasFloodViaIdentifier Has Flood via Identifier
func HasFloodViaIdentifier(iIdentifier string) bool {
	if has, err := Engine.Where("identifier = ?", iIdentifier).Get(new(Flood)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasFloodViaTimestamp Has Flood via Timestamp
func HasFloodViaTimestamp(iTimestamp int) bool {
	if has, err := Engine.Where("timestamp = ?", iTimestamp).Get(new(Flood)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasFloodViaExpiration Has Flood via Expiration
func HasFloodViaExpiration(iExpiration int) bool {
	if has, err := Engine.Where("expiration = ?", iExpiration).Get(new(Flood)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// GetFloodViaFid Get Flood via Fid
func GetFloodViaFid(iFid int64) (*Flood, error) {
	var _Flood = &Flood{Fid: iFid}
	has, err := Engine.Get(_Flood)
	if has {
		return _Flood, err
	} else {
		return nil, err
	}
}

// GetFloodViaEvent Get Flood via Event
func GetFloodViaEvent(iEvent string) (*Flood, error) {
	var _Flood = &Flood{Event: iEvent}
	has, err := Engine.Get(_Flood)
	if has {
		return _Flood, err
	} else {
		return nil, err
	}
}

// GetFloodViaIdentifier Get Flood via Identifier
func GetFloodViaIdentifier(iIdentifier string) (*Flood, error) {
	var _Flood = &Flood{Identifier: iIdentifier}
	has, err := Engine.Get(_Flood)
	if has {
		return _Flood, err
	} else {
		return nil, err
	}
}

// GetFloodViaTimestamp Get Flood via Timestamp
func GetFloodViaTimestamp(iTimestamp int) (*Flood, error) {
	var _Flood = &Flood{Timestamp: iTimestamp}
	has, err := Engine.Get(_Flood)
	if has {
		return _Flood, err
	} else {
		return nil, err
	}
}

// GetFloodViaExpiration Get Flood via Expiration
func GetFloodViaExpiration(iExpiration int) (*Flood, error) {
	var _Flood = &Flood{Expiration: iExpiration}
	has, err := Engine.Get(_Flood)
	if has {
		return _Flood, err
	} else {
		return nil, err
	}
}

// SetFloodViaFid Set Flood via Fid
func SetFloodViaFid(iFid int64, flood *Flood) (int64, error) {
	flood.Fid = iFid
	return Engine.Insert(flood)
}

// SetFloodViaEvent Set Flood via Event
func SetFloodViaEvent(iEvent string, flood *Flood) (int64, error) {
	flood.Event = iEvent
	return Engine.Insert(flood)
}

// SetFloodViaIdentifier Set Flood via Identifier
func SetFloodViaIdentifier(iIdentifier string, flood *Flood) (int64, error) {
	flood.Identifier = iIdentifier
	return Engine.Insert(flood)
}

// SetFloodViaTimestamp Set Flood via Timestamp
func SetFloodViaTimestamp(iTimestamp int, flood *Flood) (int64, error) {
	flood.Timestamp = iTimestamp
	return Engine.Insert(flood)
}

// SetFloodViaExpiration Set Flood via Expiration
func SetFloodViaExpiration(iExpiration int, flood *Flood) (int64, error) {
	flood.Expiration = iExpiration
	return Engine.Insert(flood)
}

// AddFlood Add Flood via all columns
func AddFlood(iFid int64, iEvent string, iIdentifier string, iTimestamp int, iExpiration int) error {
	_Flood := &Flood{Fid: iFid, Event: iEvent, Identifier: iIdentifier, Timestamp: iTimestamp, Expiration: iExpiration}
	if _, err := Engine.Insert(_Flood); err != nil {
		return err
	} else {
		return nil
	}
}

// PostFlood Post Flood via iFlood
func PostFlood(iFlood *Flood) (int64, error) {
	_, err := Engine.Insert(iFlood)
	return iFlood.Fid, err
}

// PutFlood Put Flood
func PutFlood(iFlood *Flood) (int64, error) {
	_, err := Engine.Id(iFlood.Fid).Update(iFlood)
	return iFlood.Fid, err
}

// PutFloodViaFid Put Flood via Fid
func PutFloodViaFid(Fid_ int64, iFlood *Flood) (int64, error) {
	row, err := Engine.Update(iFlood, &Flood{Fid: Fid_})
	return row, err
}

// PutFloodViaEvent Put Flood via Event
func PutFloodViaEvent(Event_ string, iFlood *Flood) (int64, error) {
	row, err := Engine.Update(iFlood, &Flood{Event: Event_})
	return row, err
}

// PutFloodViaIdentifier Put Flood via Identifier
func PutFloodViaIdentifier(Identifier_ string, iFlood *Flood) (int64, error) {
	row, err := Engine.Update(iFlood, &Flood{Identifier: Identifier_})
	return row, err
}

// PutFloodViaTimestamp Put Flood via Timestamp
func PutFloodViaTimestamp(Timestamp_ int, iFlood *Flood) (int64, error) {
	row, err := Engine.Update(iFlood, &Flood{Timestamp: Timestamp_})
	return row, err
}

// PutFloodViaExpiration Put Flood via Expiration
func PutFloodViaExpiration(Expiration_ int, iFlood *Flood) (int64, error) {
	row, err := Engine.Update(iFlood, &Flood{Expiration: Expiration_})
	return row, err
}

// UpdateFloodViaFid via map[string]interface{}{}
func UpdateFloodViaFid(iFid int64, iFloodMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Flood)).Where("fid = ?", iFid).Update(iFloodMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateFloodViaEvent via map[string]interface{}{}
func UpdateFloodViaEvent(iEvent string, iFloodMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Flood)).Where("event = ?", iEvent).Update(iFloodMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateFloodViaIdentifier via map[string]interface{}{}
func UpdateFloodViaIdentifier(iIdentifier string, iFloodMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Flood)).Where("identifier = ?", iIdentifier).Update(iFloodMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateFloodViaTimestamp via map[string]interface{}{}
func UpdateFloodViaTimestamp(iTimestamp int, iFloodMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Flood)).Where("timestamp = ?", iTimestamp).Update(iFloodMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateFloodViaExpiration via map[string]interface{}{}
func UpdateFloodViaExpiration(iExpiration int, iFloodMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Flood)).Where("expiration = ?", iExpiration).Update(iFloodMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// DeleteFlood Delete Flood
func DeleteFlood(iFid int64) (int64, error) {
	row, err := Engine.Id(iFid).Delete(new(Flood))
	return row, err
}

// DeleteFloodViaFid Delete Flood via Fid
func DeleteFloodViaFid(iFid int64) (err error) {
	var has bool
	var _Flood = &Flood{Fid: iFid}
	if has, err = Engine.Get(_Flood); (has == true) && (err == nil) {
		if row, err := Engine.Where("fid = ?", iFid).Delete(new(Flood)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteFloodViaEvent Delete Flood via Event
func DeleteFloodViaEvent(iEvent string) (err error) {
	var has bool
	var _Flood = &Flood{Event: iEvent}
	if has, err = Engine.Get(_Flood); (has == true) && (err == nil) {
		if row, err := Engine.Where("event = ?", iEvent).Delete(new(Flood)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteFloodViaIdentifier Delete Flood via Identifier
func DeleteFloodViaIdentifier(iIdentifier string) (err error) {
	var has bool
	var _Flood = &Flood{Identifier: iIdentifier}
	if has, err = Engine.Get(_Flood); (has == true) && (err == nil) {
		if row, err := Engine.Where("identifier = ?", iIdentifier).Delete(new(Flood)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteFloodViaTimestamp Delete Flood via Timestamp
func DeleteFloodViaTimestamp(iTimestamp int) (err error) {
	var has bool
	var _Flood = &Flood{Timestamp: iTimestamp}
	if has, err = Engine.Get(_Flood); (has == true) && (err == nil) {
		if row, err := Engine.Where("timestamp = ?", iTimestamp).Delete(new(Flood)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteFloodViaExpiration Delete Flood via Expiration
func DeleteFloodViaExpiration(iExpiration int) (err error) {
	var has bool
	var _Flood = &Flood{Expiration: iExpiration}
	if has, err = Engine.Get(_Flood); (has == true) && (err == nil) {
		if row, err := Engine.Where("expiration = ?", iExpiration).Delete(new(Flood)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}
