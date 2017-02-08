package model

type Sessions struct {
	Uid       int64  `xorm:"not null index INT(10)"`
	Sid       string `xorm:"not null pk VARCHAR(128)"`
	Hostname  string `xorm:"not null default '' VARCHAR(128)"`
	Timestamp int    `xorm:"not null default 0 index INT(11)"`
	Session   []byte `xorm:"LONGBLOB"`
}

// GetSessionsesCount Sessionss' Count
func GetSessionsesCount(offset int, limit int) (int64, error) {
	total, err := Engine.Limit(limit, offset).Count(&Sessions{})
	return total, err
}

// GetSessionsCountViaUid Get Sessions via Uid
func GetSessionsCountViaUid(iUid int64) int64 {
	n, _ := Engine.Where("uid = ?", iUid).Count(&Sessions{Uid: iUid})
	return n
}

// GetSessionsCountViaSid Get Sessions via Sid
func GetSessionsCountViaSid(iSid string) int64 {
	n, _ := Engine.Where("sid = ?", iSid).Count(&Sessions{Sid: iSid})
	return n
}

// GetSessionsCountViaHostname Get Sessions via Hostname
func GetSessionsCountViaHostname(iHostname string) int64 {
	n, _ := Engine.Where("hostname = ?", iHostname).Count(&Sessions{Hostname: iHostname})
	return n
}

// GetSessionsCountViaTimestamp Get Sessions via Timestamp
func GetSessionsCountViaTimestamp(iTimestamp int) int64 {
	n, _ := Engine.Where("timestamp = ?", iTimestamp).Count(&Sessions{Timestamp: iTimestamp})
	return n
}

// GetSessionsCountViaSession Get Sessions via Session
func GetSessionsCountViaSession(iSession []byte) int64 {
	n, _ := Engine.Where("session = ?", iSession).Count(&Sessions{Session: iSession})
	return n
}

// GetSessionsesByUidAndSidAndHostname Get Sessionses via UidAndSidAndHostname
func GetSessionsesByUidAndSidAndHostname(offset int, limit int, Uid_ int64, Sid_ string, Hostname_ string) (*[]*Sessions, error) {
	var _Sessions = new([]*Sessions)
	err := Engine.Table("sessions").Where("uid = ? and sid = ? and hostname = ?", Uid_, Sid_, Hostname_).Limit(limit, offset).Find(_Sessions)
	return _Sessions, err
}

// GetSessionsesByUidAndSidAndTimestamp Get Sessionses via UidAndSidAndTimestamp
func GetSessionsesByUidAndSidAndTimestamp(offset int, limit int, Uid_ int64, Sid_ string, Timestamp_ int) (*[]*Sessions, error) {
	var _Sessions = new([]*Sessions)
	err := Engine.Table("sessions").Where("uid = ? and sid = ? and timestamp = ?", Uid_, Sid_, Timestamp_).Limit(limit, offset).Find(_Sessions)
	return _Sessions, err
}

// GetSessionsesByUidAndSidAndSession Get Sessionses via UidAndSidAndSession
func GetSessionsesByUidAndSidAndSession(offset int, limit int, Uid_ int64, Sid_ string, Session_ []byte) (*[]*Sessions, error) {
	var _Sessions = new([]*Sessions)
	err := Engine.Table("sessions").Where("uid = ? and sid = ? and session = ?", Uid_, Sid_, Session_).Limit(limit, offset).Find(_Sessions)
	return _Sessions, err
}

// GetSessionsesByUidAndHostnameAndTimestamp Get Sessionses via UidAndHostnameAndTimestamp
func GetSessionsesByUidAndHostnameAndTimestamp(offset int, limit int, Uid_ int64, Hostname_ string, Timestamp_ int) (*[]*Sessions, error) {
	var _Sessions = new([]*Sessions)
	err := Engine.Table("sessions").Where("uid = ? and hostname = ? and timestamp = ?", Uid_, Hostname_, Timestamp_).Limit(limit, offset).Find(_Sessions)
	return _Sessions, err
}

// GetSessionsesByUidAndHostnameAndSession Get Sessionses via UidAndHostnameAndSession
func GetSessionsesByUidAndHostnameAndSession(offset int, limit int, Uid_ int64, Hostname_ string, Session_ []byte) (*[]*Sessions, error) {
	var _Sessions = new([]*Sessions)
	err := Engine.Table("sessions").Where("uid = ? and hostname = ? and session = ?", Uid_, Hostname_, Session_).Limit(limit, offset).Find(_Sessions)
	return _Sessions, err
}

// GetSessionsesByUidAndTimestampAndSession Get Sessionses via UidAndTimestampAndSession
func GetSessionsesByUidAndTimestampAndSession(offset int, limit int, Uid_ int64, Timestamp_ int, Session_ []byte) (*[]*Sessions, error) {
	var _Sessions = new([]*Sessions)
	err := Engine.Table("sessions").Where("uid = ? and timestamp = ? and session = ?", Uid_, Timestamp_, Session_).Limit(limit, offset).Find(_Sessions)
	return _Sessions, err
}

// GetSessionsesBySidAndHostnameAndTimestamp Get Sessionses via SidAndHostnameAndTimestamp
func GetSessionsesBySidAndHostnameAndTimestamp(offset int, limit int, Sid_ string, Hostname_ string, Timestamp_ int) (*[]*Sessions, error) {
	var _Sessions = new([]*Sessions)
	err := Engine.Table("sessions").Where("sid = ? and hostname = ? and timestamp = ?", Sid_, Hostname_, Timestamp_).Limit(limit, offset).Find(_Sessions)
	return _Sessions, err
}

// GetSessionsesBySidAndHostnameAndSession Get Sessionses via SidAndHostnameAndSession
func GetSessionsesBySidAndHostnameAndSession(offset int, limit int, Sid_ string, Hostname_ string, Session_ []byte) (*[]*Sessions, error) {
	var _Sessions = new([]*Sessions)
	err := Engine.Table("sessions").Where("sid = ? and hostname = ? and session = ?", Sid_, Hostname_, Session_).Limit(limit, offset).Find(_Sessions)
	return _Sessions, err
}

// GetSessionsesBySidAndTimestampAndSession Get Sessionses via SidAndTimestampAndSession
func GetSessionsesBySidAndTimestampAndSession(offset int, limit int, Sid_ string, Timestamp_ int, Session_ []byte) (*[]*Sessions, error) {
	var _Sessions = new([]*Sessions)
	err := Engine.Table("sessions").Where("sid = ? and timestamp = ? and session = ?", Sid_, Timestamp_, Session_).Limit(limit, offset).Find(_Sessions)
	return _Sessions, err
}

// GetSessionsesByHostnameAndTimestampAndSession Get Sessionses via HostnameAndTimestampAndSession
func GetSessionsesByHostnameAndTimestampAndSession(offset int, limit int, Hostname_ string, Timestamp_ int, Session_ []byte) (*[]*Sessions, error) {
	var _Sessions = new([]*Sessions)
	err := Engine.Table("sessions").Where("hostname = ? and timestamp = ? and session = ?", Hostname_, Timestamp_, Session_).Limit(limit, offset).Find(_Sessions)
	return _Sessions, err
}

// GetSessionsesByUidAndSid Get Sessionses via UidAndSid
func GetSessionsesByUidAndSid(offset int, limit int, Uid_ int64, Sid_ string) (*[]*Sessions, error) {
	var _Sessions = new([]*Sessions)
	err := Engine.Table("sessions").Where("uid = ? and sid = ?", Uid_, Sid_).Limit(limit, offset).Find(_Sessions)
	return _Sessions, err
}

// GetSessionsesByUidAndHostname Get Sessionses via UidAndHostname
func GetSessionsesByUidAndHostname(offset int, limit int, Uid_ int64, Hostname_ string) (*[]*Sessions, error) {
	var _Sessions = new([]*Sessions)
	err := Engine.Table("sessions").Where("uid = ? and hostname = ?", Uid_, Hostname_).Limit(limit, offset).Find(_Sessions)
	return _Sessions, err
}

// GetSessionsesByUidAndTimestamp Get Sessionses via UidAndTimestamp
func GetSessionsesByUidAndTimestamp(offset int, limit int, Uid_ int64, Timestamp_ int) (*[]*Sessions, error) {
	var _Sessions = new([]*Sessions)
	err := Engine.Table("sessions").Where("uid = ? and timestamp = ?", Uid_, Timestamp_).Limit(limit, offset).Find(_Sessions)
	return _Sessions, err
}

// GetSessionsesByUidAndSession Get Sessionses via UidAndSession
func GetSessionsesByUidAndSession(offset int, limit int, Uid_ int64, Session_ []byte) (*[]*Sessions, error) {
	var _Sessions = new([]*Sessions)
	err := Engine.Table("sessions").Where("uid = ? and session = ?", Uid_, Session_).Limit(limit, offset).Find(_Sessions)
	return _Sessions, err
}

// GetSessionsesBySidAndHostname Get Sessionses via SidAndHostname
func GetSessionsesBySidAndHostname(offset int, limit int, Sid_ string, Hostname_ string) (*[]*Sessions, error) {
	var _Sessions = new([]*Sessions)
	err := Engine.Table("sessions").Where("sid = ? and hostname = ?", Sid_, Hostname_).Limit(limit, offset).Find(_Sessions)
	return _Sessions, err
}

// GetSessionsesBySidAndTimestamp Get Sessionses via SidAndTimestamp
func GetSessionsesBySidAndTimestamp(offset int, limit int, Sid_ string, Timestamp_ int) (*[]*Sessions, error) {
	var _Sessions = new([]*Sessions)
	err := Engine.Table("sessions").Where("sid = ? and timestamp = ?", Sid_, Timestamp_).Limit(limit, offset).Find(_Sessions)
	return _Sessions, err
}

// GetSessionsesBySidAndSession Get Sessionses via SidAndSession
func GetSessionsesBySidAndSession(offset int, limit int, Sid_ string, Session_ []byte) (*[]*Sessions, error) {
	var _Sessions = new([]*Sessions)
	err := Engine.Table("sessions").Where("sid = ? and session = ?", Sid_, Session_).Limit(limit, offset).Find(_Sessions)
	return _Sessions, err
}

// GetSessionsesByHostnameAndTimestamp Get Sessionses via HostnameAndTimestamp
func GetSessionsesByHostnameAndTimestamp(offset int, limit int, Hostname_ string, Timestamp_ int) (*[]*Sessions, error) {
	var _Sessions = new([]*Sessions)
	err := Engine.Table("sessions").Where("hostname = ? and timestamp = ?", Hostname_, Timestamp_).Limit(limit, offset).Find(_Sessions)
	return _Sessions, err
}

// GetSessionsesByHostnameAndSession Get Sessionses via HostnameAndSession
func GetSessionsesByHostnameAndSession(offset int, limit int, Hostname_ string, Session_ []byte) (*[]*Sessions, error) {
	var _Sessions = new([]*Sessions)
	err := Engine.Table("sessions").Where("hostname = ? and session = ?", Hostname_, Session_).Limit(limit, offset).Find(_Sessions)
	return _Sessions, err
}

// GetSessionsesByTimestampAndSession Get Sessionses via TimestampAndSession
func GetSessionsesByTimestampAndSession(offset int, limit int, Timestamp_ int, Session_ []byte) (*[]*Sessions, error) {
	var _Sessions = new([]*Sessions)
	err := Engine.Table("sessions").Where("timestamp = ? and session = ?", Timestamp_, Session_).Limit(limit, offset).Find(_Sessions)
	return _Sessions, err
}

// GetSessionses Get Sessionses via field
func GetSessionses(offset int, limit int, field string) (*[]*Sessions, error) {
	var _Sessions = new([]*Sessions)
	err := Engine.Table("sessions").Limit(limit, offset).Desc(field).Find(_Sessions)
	return _Sessions, err
}

// GetSessionsesViaUid Get Sessionss via Uid
func GetSessionsesViaUid(offset int, limit int, Uid_ int64, field string) (*[]*Sessions, error) {
	var _Sessions = new([]*Sessions)
	err := Engine.Table("sessions").Where("uid = ?", Uid_).Limit(limit, offset).Desc(field).Find(_Sessions)
	return _Sessions, err
}

// GetSessionsesViaSid Get Sessionss via Sid
func GetSessionsesViaSid(offset int, limit int, Sid_ string, field string) (*[]*Sessions, error) {
	var _Sessions = new([]*Sessions)
	err := Engine.Table("sessions").Where("sid = ?", Sid_).Limit(limit, offset).Desc(field).Find(_Sessions)
	return _Sessions, err
}

// GetSessionsesViaHostname Get Sessionss via Hostname
func GetSessionsesViaHostname(offset int, limit int, Hostname_ string, field string) (*[]*Sessions, error) {
	var _Sessions = new([]*Sessions)
	err := Engine.Table("sessions").Where("hostname = ?", Hostname_).Limit(limit, offset).Desc(field).Find(_Sessions)
	return _Sessions, err
}

// GetSessionsesViaTimestamp Get Sessionss via Timestamp
func GetSessionsesViaTimestamp(offset int, limit int, Timestamp_ int, field string) (*[]*Sessions, error) {
	var _Sessions = new([]*Sessions)
	err := Engine.Table("sessions").Where("timestamp = ?", Timestamp_).Limit(limit, offset).Desc(field).Find(_Sessions)
	return _Sessions, err
}

// GetSessionsesViaSession Get Sessionss via Session
func GetSessionsesViaSession(offset int, limit int, Session_ []byte, field string) (*[]*Sessions, error) {
	var _Sessions = new([]*Sessions)
	err := Engine.Table("sessions").Where("session = ?", Session_).Limit(limit, offset).Desc(field).Find(_Sessions)
	return _Sessions, err
}

// HasSessionsViaUid Has Sessions via Uid
func HasSessionsViaUid(iUid int64) bool {
	if has, err := Engine.Where("uid = ?", iUid).Get(new(Sessions)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasSessionsViaSid Has Sessions via Sid
func HasSessionsViaSid(iSid string) bool {
	if has, err := Engine.Where("sid = ?", iSid).Get(new(Sessions)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasSessionsViaHostname Has Sessions via Hostname
func HasSessionsViaHostname(iHostname string) bool {
	if has, err := Engine.Where("hostname = ?", iHostname).Get(new(Sessions)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasSessionsViaTimestamp Has Sessions via Timestamp
func HasSessionsViaTimestamp(iTimestamp int) bool {
	if has, err := Engine.Where("timestamp = ?", iTimestamp).Get(new(Sessions)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasSessionsViaSession Has Sessions via Session
func HasSessionsViaSession(iSession []byte) bool {
	if has, err := Engine.Where("session = ?", iSession).Get(new(Sessions)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// GetSessionsViaUid Get Sessions via Uid
func GetSessionsViaUid(iUid int64) (*Sessions, error) {
	var _Sessions = &Sessions{Uid: iUid}
	has, err := Engine.Get(_Sessions)
	if has {
		return _Sessions, err
	} else {
		return nil, err
	}
}

// GetSessionsViaSid Get Sessions via Sid
func GetSessionsViaSid(iSid string) (*Sessions, error) {
	var _Sessions = &Sessions{Sid: iSid}
	has, err := Engine.Get(_Sessions)
	if has {
		return _Sessions, err
	} else {
		return nil, err
	}
}

// GetSessionsViaHostname Get Sessions via Hostname
func GetSessionsViaHostname(iHostname string) (*Sessions, error) {
	var _Sessions = &Sessions{Hostname: iHostname}
	has, err := Engine.Get(_Sessions)
	if has {
		return _Sessions, err
	} else {
		return nil, err
	}
}

// GetSessionsViaTimestamp Get Sessions via Timestamp
func GetSessionsViaTimestamp(iTimestamp int) (*Sessions, error) {
	var _Sessions = &Sessions{Timestamp: iTimestamp}
	has, err := Engine.Get(_Sessions)
	if has {
		return _Sessions, err
	} else {
		return nil, err
	}
}

// GetSessionsViaSession Get Sessions via Session
func GetSessionsViaSession(iSession []byte) (*Sessions, error) {
	var _Sessions = &Sessions{Session: iSession}
	has, err := Engine.Get(_Sessions)
	if has {
		return _Sessions, err
	} else {
		return nil, err
	}
}

// SetSessionsViaUid Set Sessions via Uid
func SetSessionsViaUid(iUid int64, sessions *Sessions) (int64, error) {
	sessions.Uid = iUid
	return Engine.Insert(sessions)
}

// SetSessionsViaSid Set Sessions via Sid
func SetSessionsViaSid(iSid string, sessions *Sessions) (int64, error) {
	sessions.Sid = iSid
	return Engine.Insert(sessions)
}

// SetSessionsViaHostname Set Sessions via Hostname
func SetSessionsViaHostname(iHostname string, sessions *Sessions) (int64, error) {
	sessions.Hostname = iHostname
	return Engine.Insert(sessions)
}

// SetSessionsViaTimestamp Set Sessions via Timestamp
func SetSessionsViaTimestamp(iTimestamp int, sessions *Sessions) (int64, error) {
	sessions.Timestamp = iTimestamp
	return Engine.Insert(sessions)
}

// SetSessionsViaSession Set Sessions via Session
func SetSessionsViaSession(iSession []byte, sessions *Sessions) (int64, error) {
	sessions.Session = iSession
	return Engine.Insert(sessions)
}

// AddSessions Add Sessions via all columns
func AddSessions(iUid int64, iSid string, iHostname string, iTimestamp int, iSession []byte) error {
	_Sessions := &Sessions{Uid: iUid, Sid: iSid, Hostname: iHostname, Timestamp: iTimestamp, Session: iSession}
	if _, err := Engine.Insert(_Sessions); err != nil {
		return err
	} else {
		return nil
	}
}

// PostSessions Post Sessions via iSessions
func PostSessions(iSessions *Sessions) (int64, error) {
	_, err := Engine.Insert(iSessions)
	return iSessions.Uid, err
}

// PutSessions Put Sessions
func PutSessions(iSessions *Sessions) (int64, error) {
	_, err := Engine.Id(iSessions.Uid).Update(iSessions)
	return iSessions.Uid, err
}

// PutSessionsViaUid Put Sessions via Uid
func PutSessionsViaUid(Uid_ int64, iSessions *Sessions) (int64, error) {
	row, err := Engine.Update(iSessions, &Sessions{Uid: Uid_})
	return row, err
}

// PutSessionsViaSid Put Sessions via Sid
func PutSessionsViaSid(Sid_ string, iSessions *Sessions) (int64, error) {
	row, err := Engine.Update(iSessions, &Sessions{Sid: Sid_})
	return row, err
}

// PutSessionsViaHostname Put Sessions via Hostname
func PutSessionsViaHostname(Hostname_ string, iSessions *Sessions) (int64, error) {
	row, err := Engine.Update(iSessions, &Sessions{Hostname: Hostname_})
	return row, err
}

// PutSessionsViaTimestamp Put Sessions via Timestamp
func PutSessionsViaTimestamp(Timestamp_ int, iSessions *Sessions) (int64, error) {
	row, err := Engine.Update(iSessions, &Sessions{Timestamp: Timestamp_})
	return row, err
}

// PutSessionsViaSession Put Sessions via Session
func PutSessionsViaSession(Session_ []byte, iSessions *Sessions) (int64, error) {
	row, err := Engine.Update(iSessions, &Sessions{Session: Session_})
	return row, err
}

// UpdateSessionsViaUid via map[string]interface{}{}
func UpdateSessionsViaUid(iUid int64, iSessionsMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Sessions)).Where("uid = ?", iUid).Update(iSessionsMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateSessionsViaSid via map[string]interface{}{}
func UpdateSessionsViaSid(iSid string, iSessionsMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Sessions)).Where("sid = ?", iSid).Update(iSessionsMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateSessionsViaHostname via map[string]interface{}{}
func UpdateSessionsViaHostname(iHostname string, iSessionsMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Sessions)).Where("hostname = ?", iHostname).Update(iSessionsMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateSessionsViaTimestamp via map[string]interface{}{}
func UpdateSessionsViaTimestamp(iTimestamp int, iSessionsMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Sessions)).Where("timestamp = ?", iTimestamp).Update(iSessionsMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateSessionsViaSession via map[string]interface{}{}
func UpdateSessionsViaSession(iSession []byte, iSessionsMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Sessions)).Where("session = ?", iSession).Update(iSessionsMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// DeleteSessions Delete Sessions
func DeleteSessions(iUid int64) (int64, error) {
	row, err := Engine.Id(iUid).Delete(new(Sessions))
	return row, err
}

// DeleteSessionsViaUid Delete Sessions via Uid
func DeleteSessionsViaUid(iUid int64) (err error) {
	var has bool
	var _Sessions = &Sessions{Uid: iUid}
	if has, err = Engine.Get(_Sessions); (has == true) && (err == nil) {
		if row, err := Engine.Where("uid = ?", iUid).Delete(new(Sessions)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteSessionsViaSid Delete Sessions via Sid
func DeleteSessionsViaSid(iSid string) (err error) {
	var has bool
	var _Sessions = &Sessions{Sid: iSid}
	if has, err = Engine.Get(_Sessions); (has == true) && (err == nil) {
		if row, err := Engine.Where("sid = ?", iSid).Delete(new(Sessions)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteSessionsViaHostname Delete Sessions via Hostname
func DeleteSessionsViaHostname(iHostname string) (err error) {
	var has bool
	var _Sessions = &Sessions{Hostname: iHostname}
	if has, err = Engine.Get(_Sessions); (has == true) && (err == nil) {
		if row, err := Engine.Where("hostname = ?", iHostname).Delete(new(Sessions)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteSessionsViaTimestamp Delete Sessions via Timestamp
func DeleteSessionsViaTimestamp(iTimestamp int) (err error) {
	var has bool
	var _Sessions = &Sessions{Timestamp: iTimestamp}
	if has, err = Engine.Get(_Sessions); (has == true) && (err == nil) {
		if row, err := Engine.Where("timestamp = ?", iTimestamp).Delete(new(Sessions)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteSessionsViaSession Delete Sessions via Session
func DeleteSessionsViaSession(iSession []byte) (err error) {
	var has bool
	var _Sessions = &Sessions{Session: iSession}
	if has, err = Engine.Get(_Sessions); (has == true) && (err == nil) {
		if row, err := Engine.Where("session = ?", iSession).Delete(new(Sessions)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}
