package model

type NodeRevision struct {
	Nid               int64  `xorm:"not null index INT(10)"`
	Vid               int    `xorm:"not null pk autoincr INT(10)"`
	Langcode          string `xorm:"not null index VARCHAR(12)"`
	RevisionTimestamp int    `xorm:"INT(11)"`
	RevisionUid       int    `xorm:"index INT(10)"`
	RevisionLog       string `xorm:"LONGTEXT"`
}

// GetNodeRevisionsCount NodeRevisions' Count
func GetNodeRevisionsCount(offset int, limit int) (int64, error) {
	total, err := Engine.Limit(limit, offset).Count(&NodeRevision{})
	return total, err
}

// GetNodeRevisionCountViaNid Get NodeRevision via Nid
func GetNodeRevisionCountViaNid(iNid int64) int64 {
	n, _ := Engine.Where("nid = ?", iNid).Count(&NodeRevision{Nid: iNid})
	return n
}

// GetNodeRevisionCountViaVid Get NodeRevision via Vid
func GetNodeRevisionCountViaVid(iVid int) int64 {
	n, _ := Engine.Where("vid = ?", iVid).Count(&NodeRevision{Vid: iVid})
	return n
}

// GetNodeRevisionCountViaLangcode Get NodeRevision via Langcode
func GetNodeRevisionCountViaLangcode(iLangcode string) int64 {
	n, _ := Engine.Where("langcode = ?", iLangcode).Count(&NodeRevision{Langcode: iLangcode})
	return n
}

// GetNodeRevisionCountViaRevisionTimestamp Get NodeRevision via RevisionTimestamp
func GetNodeRevisionCountViaRevisionTimestamp(iRevisionTimestamp int) int64 {
	n, _ := Engine.Where("revision_timestamp = ?", iRevisionTimestamp).Count(&NodeRevision{RevisionTimestamp: iRevisionTimestamp})
	return n
}

// GetNodeRevisionCountViaRevisionUid Get NodeRevision via RevisionUid
func GetNodeRevisionCountViaRevisionUid(iRevisionUid int) int64 {
	n, _ := Engine.Where("revision_uid = ?", iRevisionUid).Count(&NodeRevision{RevisionUid: iRevisionUid})
	return n
}

// GetNodeRevisionCountViaRevisionLog Get NodeRevision via RevisionLog
func GetNodeRevisionCountViaRevisionLog(iRevisionLog string) int64 {
	n, _ := Engine.Where("revision_log = ?", iRevisionLog).Count(&NodeRevision{RevisionLog: iRevisionLog})
	return n
}

// GetNodeRevisionsByNidAndVidAndLangcode Get NodeRevisions via NidAndVidAndLangcode
func GetNodeRevisionsByNidAndVidAndLangcode(offset int, limit int, Nid_ int64, Vid_ int, Langcode_ string) (*[]*NodeRevision, error) {
	var _NodeRevision = new([]*NodeRevision)
	err := Engine.Table("node_revision").Where("nid = ? and vid = ? and langcode = ?", Nid_, Vid_, Langcode_).Limit(limit, offset).Find(_NodeRevision)
	return _NodeRevision, err
}

// GetNodeRevisionsByNidAndVidAndRevisionTimestamp Get NodeRevisions via NidAndVidAndRevisionTimestamp
func GetNodeRevisionsByNidAndVidAndRevisionTimestamp(offset int, limit int, Nid_ int64, Vid_ int, RevisionTimestamp_ int) (*[]*NodeRevision, error) {
	var _NodeRevision = new([]*NodeRevision)
	err := Engine.Table("node_revision").Where("nid = ? and vid = ? and revision_timestamp = ?", Nid_, Vid_, RevisionTimestamp_).Limit(limit, offset).Find(_NodeRevision)
	return _NodeRevision, err
}

// GetNodeRevisionsByNidAndVidAndRevisionUid Get NodeRevisions via NidAndVidAndRevisionUid
func GetNodeRevisionsByNidAndVidAndRevisionUid(offset int, limit int, Nid_ int64, Vid_ int, RevisionUid_ int) (*[]*NodeRevision, error) {
	var _NodeRevision = new([]*NodeRevision)
	err := Engine.Table("node_revision").Where("nid = ? and vid = ? and revision_uid = ?", Nid_, Vid_, RevisionUid_).Limit(limit, offset).Find(_NodeRevision)
	return _NodeRevision, err
}

// GetNodeRevisionsByNidAndVidAndRevisionLog Get NodeRevisions via NidAndVidAndRevisionLog
func GetNodeRevisionsByNidAndVidAndRevisionLog(offset int, limit int, Nid_ int64, Vid_ int, RevisionLog_ string) (*[]*NodeRevision, error) {
	var _NodeRevision = new([]*NodeRevision)
	err := Engine.Table("node_revision").Where("nid = ? and vid = ? and revision_log = ?", Nid_, Vid_, RevisionLog_).Limit(limit, offset).Find(_NodeRevision)
	return _NodeRevision, err
}

// GetNodeRevisionsByNidAndLangcodeAndRevisionTimestamp Get NodeRevisions via NidAndLangcodeAndRevisionTimestamp
func GetNodeRevisionsByNidAndLangcodeAndRevisionTimestamp(offset int, limit int, Nid_ int64, Langcode_ string, RevisionTimestamp_ int) (*[]*NodeRevision, error) {
	var _NodeRevision = new([]*NodeRevision)
	err := Engine.Table("node_revision").Where("nid = ? and langcode = ? and revision_timestamp = ?", Nid_, Langcode_, RevisionTimestamp_).Limit(limit, offset).Find(_NodeRevision)
	return _NodeRevision, err
}

// GetNodeRevisionsByNidAndLangcodeAndRevisionUid Get NodeRevisions via NidAndLangcodeAndRevisionUid
func GetNodeRevisionsByNidAndLangcodeAndRevisionUid(offset int, limit int, Nid_ int64, Langcode_ string, RevisionUid_ int) (*[]*NodeRevision, error) {
	var _NodeRevision = new([]*NodeRevision)
	err := Engine.Table("node_revision").Where("nid = ? and langcode = ? and revision_uid = ?", Nid_, Langcode_, RevisionUid_).Limit(limit, offset).Find(_NodeRevision)
	return _NodeRevision, err
}

// GetNodeRevisionsByNidAndLangcodeAndRevisionLog Get NodeRevisions via NidAndLangcodeAndRevisionLog
func GetNodeRevisionsByNidAndLangcodeAndRevisionLog(offset int, limit int, Nid_ int64, Langcode_ string, RevisionLog_ string) (*[]*NodeRevision, error) {
	var _NodeRevision = new([]*NodeRevision)
	err := Engine.Table("node_revision").Where("nid = ? and langcode = ? and revision_log = ?", Nid_, Langcode_, RevisionLog_).Limit(limit, offset).Find(_NodeRevision)
	return _NodeRevision, err
}

// GetNodeRevisionsByNidAndRevisionTimestampAndRevisionUid Get NodeRevisions via NidAndRevisionTimestampAndRevisionUid
func GetNodeRevisionsByNidAndRevisionTimestampAndRevisionUid(offset int, limit int, Nid_ int64, RevisionTimestamp_ int, RevisionUid_ int) (*[]*NodeRevision, error) {
	var _NodeRevision = new([]*NodeRevision)
	err := Engine.Table("node_revision").Where("nid = ? and revision_timestamp = ? and revision_uid = ?", Nid_, RevisionTimestamp_, RevisionUid_).Limit(limit, offset).Find(_NodeRevision)
	return _NodeRevision, err
}

// GetNodeRevisionsByNidAndRevisionTimestampAndRevisionLog Get NodeRevisions via NidAndRevisionTimestampAndRevisionLog
func GetNodeRevisionsByNidAndRevisionTimestampAndRevisionLog(offset int, limit int, Nid_ int64, RevisionTimestamp_ int, RevisionLog_ string) (*[]*NodeRevision, error) {
	var _NodeRevision = new([]*NodeRevision)
	err := Engine.Table("node_revision").Where("nid = ? and revision_timestamp = ? and revision_log = ?", Nid_, RevisionTimestamp_, RevisionLog_).Limit(limit, offset).Find(_NodeRevision)
	return _NodeRevision, err
}

// GetNodeRevisionsByNidAndRevisionUidAndRevisionLog Get NodeRevisions via NidAndRevisionUidAndRevisionLog
func GetNodeRevisionsByNidAndRevisionUidAndRevisionLog(offset int, limit int, Nid_ int64, RevisionUid_ int, RevisionLog_ string) (*[]*NodeRevision, error) {
	var _NodeRevision = new([]*NodeRevision)
	err := Engine.Table("node_revision").Where("nid = ? and revision_uid = ? and revision_log = ?", Nid_, RevisionUid_, RevisionLog_).Limit(limit, offset).Find(_NodeRevision)
	return _NodeRevision, err
}

// GetNodeRevisionsByVidAndLangcodeAndRevisionTimestamp Get NodeRevisions via VidAndLangcodeAndRevisionTimestamp
func GetNodeRevisionsByVidAndLangcodeAndRevisionTimestamp(offset int, limit int, Vid_ int, Langcode_ string, RevisionTimestamp_ int) (*[]*NodeRevision, error) {
	var _NodeRevision = new([]*NodeRevision)
	err := Engine.Table("node_revision").Where("vid = ? and langcode = ? and revision_timestamp = ?", Vid_, Langcode_, RevisionTimestamp_).Limit(limit, offset).Find(_NodeRevision)
	return _NodeRevision, err
}

// GetNodeRevisionsByVidAndLangcodeAndRevisionUid Get NodeRevisions via VidAndLangcodeAndRevisionUid
func GetNodeRevisionsByVidAndLangcodeAndRevisionUid(offset int, limit int, Vid_ int, Langcode_ string, RevisionUid_ int) (*[]*NodeRevision, error) {
	var _NodeRevision = new([]*NodeRevision)
	err := Engine.Table("node_revision").Where("vid = ? and langcode = ? and revision_uid = ?", Vid_, Langcode_, RevisionUid_).Limit(limit, offset).Find(_NodeRevision)
	return _NodeRevision, err
}

// GetNodeRevisionsByVidAndLangcodeAndRevisionLog Get NodeRevisions via VidAndLangcodeAndRevisionLog
func GetNodeRevisionsByVidAndLangcodeAndRevisionLog(offset int, limit int, Vid_ int, Langcode_ string, RevisionLog_ string) (*[]*NodeRevision, error) {
	var _NodeRevision = new([]*NodeRevision)
	err := Engine.Table("node_revision").Where("vid = ? and langcode = ? and revision_log = ?", Vid_, Langcode_, RevisionLog_).Limit(limit, offset).Find(_NodeRevision)
	return _NodeRevision, err
}

// GetNodeRevisionsByVidAndRevisionTimestampAndRevisionUid Get NodeRevisions via VidAndRevisionTimestampAndRevisionUid
func GetNodeRevisionsByVidAndRevisionTimestampAndRevisionUid(offset int, limit int, Vid_ int, RevisionTimestamp_ int, RevisionUid_ int) (*[]*NodeRevision, error) {
	var _NodeRevision = new([]*NodeRevision)
	err := Engine.Table("node_revision").Where("vid = ? and revision_timestamp = ? and revision_uid = ?", Vid_, RevisionTimestamp_, RevisionUid_).Limit(limit, offset).Find(_NodeRevision)
	return _NodeRevision, err
}

// GetNodeRevisionsByVidAndRevisionTimestampAndRevisionLog Get NodeRevisions via VidAndRevisionTimestampAndRevisionLog
func GetNodeRevisionsByVidAndRevisionTimestampAndRevisionLog(offset int, limit int, Vid_ int, RevisionTimestamp_ int, RevisionLog_ string) (*[]*NodeRevision, error) {
	var _NodeRevision = new([]*NodeRevision)
	err := Engine.Table("node_revision").Where("vid = ? and revision_timestamp = ? and revision_log = ?", Vid_, RevisionTimestamp_, RevisionLog_).Limit(limit, offset).Find(_NodeRevision)
	return _NodeRevision, err
}

// GetNodeRevisionsByVidAndRevisionUidAndRevisionLog Get NodeRevisions via VidAndRevisionUidAndRevisionLog
func GetNodeRevisionsByVidAndRevisionUidAndRevisionLog(offset int, limit int, Vid_ int, RevisionUid_ int, RevisionLog_ string) (*[]*NodeRevision, error) {
	var _NodeRevision = new([]*NodeRevision)
	err := Engine.Table("node_revision").Where("vid = ? and revision_uid = ? and revision_log = ?", Vid_, RevisionUid_, RevisionLog_).Limit(limit, offset).Find(_NodeRevision)
	return _NodeRevision, err
}

// GetNodeRevisionsByLangcodeAndRevisionTimestampAndRevisionUid Get NodeRevisions via LangcodeAndRevisionTimestampAndRevisionUid
func GetNodeRevisionsByLangcodeAndRevisionTimestampAndRevisionUid(offset int, limit int, Langcode_ string, RevisionTimestamp_ int, RevisionUid_ int) (*[]*NodeRevision, error) {
	var _NodeRevision = new([]*NodeRevision)
	err := Engine.Table("node_revision").Where("langcode = ? and revision_timestamp = ? and revision_uid = ?", Langcode_, RevisionTimestamp_, RevisionUid_).Limit(limit, offset).Find(_NodeRevision)
	return _NodeRevision, err
}

// GetNodeRevisionsByLangcodeAndRevisionTimestampAndRevisionLog Get NodeRevisions via LangcodeAndRevisionTimestampAndRevisionLog
func GetNodeRevisionsByLangcodeAndRevisionTimestampAndRevisionLog(offset int, limit int, Langcode_ string, RevisionTimestamp_ int, RevisionLog_ string) (*[]*NodeRevision, error) {
	var _NodeRevision = new([]*NodeRevision)
	err := Engine.Table("node_revision").Where("langcode = ? and revision_timestamp = ? and revision_log = ?", Langcode_, RevisionTimestamp_, RevisionLog_).Limit(limit, offset).Find(_NodeRevision)
	return _NodeRevision, err
}

// GetNodeRevisionsByLangcodeAndRevisionUidAndRevisionLog Get NodeRevisions via LangcodeAndRevisionUidAndRevisionLog
func GetNodeRevisionsByLangcodeAndRevisionUidAndRevisionLog(offset int, limit int, Langcode_ string, RevisionUid_ int, RevisionLog_ string) (*[]*NodeRevision, error) {
	var _NodeRevision = new([]*NodeRevision)
	err := Engine.Table("node_revision").Where("langcode = ? and revision_uid = ? and revision_log = ?", Langcode_, RevisionUid_, RevisionLog_).Limit(limit, offset).Find(_NodeRevision)
	return _NodeRevision, err
}

// GetNodeRevisionsByRevisionTimestampAndRevisionUidAndRevisionLog Get NodeRevisions via RevisionTimestampAndRevisionUidAndRevisionLog
func GetNodeRevisionsByRevisionTimestampAndRevisionUidAndRevisionLog(offset int, limit int, RevisionTimestamp_ int, RevisionUid_ int, RevisionLog_ string) (*[]*NodeRevision, error) {
	var _NodeRevision = new([]*NodeRevision)
	err := Engine.Table("node_revision").Where("revision_timestamp = ? and revision_uid = ? and revision_log = ?", RevisionTimestamp_, RevisionUid_, RevisionLog_).Limit(limit, offset).Find(_NodeRevision)
	return _NodeRevision, err
}

// GetNodeRevisionsByNidAndVid Get NodeRevisions via NidAndVid
func GetNodeRevisionsByNidAndVid(offset int, limit int, Nid_ int64, Vid_ int) (*[]*NodeRevision, error) {
	var _NodeRevision = new([]*NodeRevision)
	err := Engine.Table("node_revision").Where("nid = ? and vid = ?", Nid_, Vid_).Limit(limit, offset).Find(_NodeRevision)
	return _NodeRevision, err
}

// GetNodeRevisionsByNidAndLangcode Get NodeRevisions via NidAndLangcode
func GetNodeRevisionsByNidAndLangcode(offset int, limit int, Nid_ int64, Langcode_ string) (*[]*NodeRevision, error) {
	var _NodeRevision = new([]*NodeRevision)
	err := Engine.Table("node_revision").Where("nid = ? and langcode = ?", Nid_, Langcode_).Limit(limit, offset).Find(_NodeRevision)
	return _NodeRevision, err
}

// GetNodeRevisionsByNidAndRevisionTimestamp Get NodeRevisions via NidAndRevisionTimestamp
func GetNodeRevisionsByNidAndRevisionTimestamp(offset int, limit int, Nid_ int64, RevisionTimestamp_ int) (*[]*NodeRevision, error) {
	var _NodeRevision = new([]*NodeRevision)
	err := Engine.Table("node_revision").Where("nid = ? and revision_timestamp = ?", Nid_, RevisionTimestamp_).Limit(limit, offset).Find(_NodeRevision)
	return _NodeRevision, err
}

// GetNodeRevisionsByNidAndRevisionUid Get NodeRevisions via NidAndRevisionUid
func GetNodeRevisionsByNidAndRevisionUid(offset int, limit int, Nid_ int64, RevisionUid_ int) (*[]*NodeRevision, error) {
	var _NodeRevision = new([]*NodeRevision)
	err := Engine.Table("node_revision").Where("nid = ? and revision_uid = ?", Nid_, RevisionUid_).Limit(limit, offset).Find(_NodeRevision)
	return _NodeRevision, err
}

// GetNodeRevisionsByNidAndRevisionLog Get NodeRevisions via NidAndRevisionLog
func GetNodeRevisionsByNidAndRevisionLog(offset int, limit int, Nid_ int64, RevisionLog_ string) (*[]*NodeRevision, error) {
	var _NodeRevision = new([]*NodeRevision)
	err := Engine.Table("node_revision").Where("nid = ? and revision_log = ?", Nid_, RevisionLog_).Limit(limit, offset).Find(_NodeRevision)
	return _NodeRevision, err
}

// GetNodeRevisionsByVidAndLangcode Get NodeRevisions via VidAndLangcode
func GetNodeRevisionsByVidAndLangcode(offset int, limit int, Vid_ int, Langcode_ string) (*[]*NodeRevision, error) {
	var _NodeRevision = new([]*NodeRevision)
	err := Engine.Table("node_revision").Where("vid = ? and langcode = ?", Vid_, Langcode_).Limit(limit, offset).Find(_NodeRevision)
	return _NodeRevision, err
}

// GetNodeRevisionsByVidAndRevisionTimestamp Get NodeRevisions via VidAndRevisionTimestamp
func GetNodeRevisionsByVidAndRevisionTimestamp(offset int, limit int, Vid_ int, RevisionTimestamp_ int) (*[]*NodeRevision, error) {
	var _NodeRevision = new([]*NodeRevision)
	err := Engine.Table("node_revision").Where("vid = ? and revision_timestamp = ?", Vid_, RevisionTimestamp_).Limit(limit, offset).Find(_NodeRevision)
	return _NodeRevision, err
}

// GetNodeRevisionsByVidAndRevisionUid Get NodeRevisions via VidAndRevisionUid
func GetNodeRevisionsByVidAndRevisionUid(offset int, limit int, Vid_ int, RevisionUid_ int) (*[]*NodeRevision, error) {
	var _NodeRevision = new([]*NodeRevision)
	err := Engine.Table("node_revision").Where("vid = ? and revision_uid = ?", Vid_, RevisionUid_).Limit(limit, offset).Find(_NodeRevision)
	return _NodeRevision, err
}

// GetNodeRevisionsByVidAndRevisionLog Get NodeRevisions via VidAndRevisionLog
func GetNodeRevisionsByVidAndRevisionLog(offset int, limit int, Vid_ int, RevisionLog_ string) (*[]*NodeRevision, error) {
	var _NodeRevision = new([]*NodeRevision)
	err := Engine.Table("node_revision").Where("vid = ? and revision_log = ?", Vid_, RevisionLog_).Limit(limit, offset).Find(_NodeRevision)
	return _NodeRevision, err
}

// GetNodeRevisionsByLangcodeAndRevisionTimestamp Get NodeRevisions via LangcodeAndRevisionTimestamp
func GetNodeRevisionsByLangcodeAndRevisionTimestamp(offset int, limit int, Langcode_ string, RevisionTimestamp_ int) (*[]*NodeRevision, error) {
	var _NodeRevision = new([]*NodeRevision)
	err := Engine.Table("node_revision").Where("langcode = ? and revision_timestamp = ?", Langcode_, RevisionTimestamp_).Limit(limit, offset).Find(_NodeRevision)
	return _NodeRevision, err
}

// GetNodeRevisionsByLangcodeAndRevisionUid Get NodeRevisions via LangcodeAndRevisionUid
func GetNodeRevisionsByLangcodeAndRevisionUid(offset int, limit int, Langcode_ string, RevisionUid_ int) (*[]*NodeRevision, error) {
	var _NodeRevision = new([]*NodeRevision)
	err := Engine.Table("node_revision").Where("langcode = ? and revision_uid = ?", Langcode_, RevisionUid_).Limit(limit, offset).Find(_NodeRevision)
	return _NodeRevision, err
}

// GetNodeRevisionsByLangcodeAndRevisionLog Get NodeRevisions via LangcodeAndRevisionLog
func GetNodeRevisionsByLangcodeAndRevisionLog(offset int, limit int, Langcode_ string, RevisionLog_ string) (*[]*NodeRevision, error) {
	var _NodeRevision = new([]*NodeRevision)
	err := Engine.Table("node_revision").Where("langcode = ? and revision_log = ?", Langcode_, RevisionLog_).Limit(limit, offset).Find(_NodeRevision)
	return _NodeRevision, err
}

// GetNodeRevisionsByRevisionTimestampAndRevisionUid Get NodeRevisions via RevisionTimestampAndRevisionUid
func GetNodeRevisionsByRevisionTimestampAndRevisionUid(offset int, limit int, RevisionTimestamp_ int, RevisionUid_ int) (*[]*NodeRevision, error) {
	var _NodeRevision = new([]*NodeRevision)
	err := Engine.Table("node_revision").Where("revision_timestamp = ? and revision_uid = ?", RevisionTimestamp_, RevisionUid_).Limit(limit, offset).Find(_NodeRevision)
	return _NodeRevision, err
}

// GetNodeRevisionsByRevisionTimestampAndRevisionLog Get NodeRevisions via RevisionTimestampAndRevisionLog
func GetNodeRevisionsByRevisionTimestampAndRevisionLog(offset int, limit int, RevisionTimestamp_ int, RevisionLog_ string) (*[]*NodeRevision, error) {
	var _NodeRevision = new([]*NodeRevision)
	err := Engine.Table("node_revision").Where("revision_timestamp = ? and revision_log = ?", RevisionTimestamp_, RevisionLog_).Limit(limit, offset).Find(_NodeRevision)
	return _NodeRevision, err
}

// GetNodeRevisionsByRevisionUidAndRevisionLog Get NodeRevisions via RevisionUidAndRevisionLog
func GetNodeRevisionsByRevisionUidAndRevisionLog(offset int, limit int, RevisionUid_ int, RevisionLog_ string) (*[]*NodeRevision, error) {
	var _NodeRevision = new([]*NodeRevision)
	err := Engine.Table("node_revision").Where("revision_uid = ? and revision_log = ?", RevisionUid_, RevisionLog_).Limit(limit, offset).Find(_NodeRevision)
	return _NodeRevision, err
}

// GetNodeRevisions Get NodeRevisions via field
func GetNodeRevisions(offset int, limit int, field string) (*[]*NodeRevision, error) {
	var _NodeRevision = new([]*NodeRevision)
	err := Engine.Table("node_revision").Limit(limit, offset).Desc(field).Find(_NodeRevision)
	return _NodeRevision, err
}

// GetNodeRevisionsViaNid Get NodeRevisions via Nid
func GetNodeRevisionsViaNid(offset int, limit int, Nid_ int64, field string) (*[]*NodeRevision, error) {
	var _NodeRevision = new([]*NodeRevision)
	err := Engine.Table("node_revision").Where("nid = ?", Nid_).Limit(limit, offset).Desc(field).Find(_NodeRevision)
	return _NodeRevision, err
}

// GetNodeRevisionsViaVid Get NodeRevisions via Vid
func GetNodeRevisionsViaVid(offset int, limit int, Vid_ int, field string) (*[]*NodeRevision, error) {
	var _NodeRevision = new([]*NodeRevision)
	err := Engine.Table("node_revision").Where("vid = ?", Vid_).Limit(limit, offset).Desc(field).Find(_NodeRevision)
	return _NodeRevision, err
}

// GetNodeRevisionsViaLangcode Get NodeRevisions via Langcode
func GetNodeRevisionsViaLangcode(offset int, limit int, Langcode_ string, field string) (*[]*NodeRevision, error) {
	var _NodeRevision = new([]*NodeRevision)
	err := Engine.Table("node_revision").Where("langcode = ?", Langcode_).Limit(limit, offset).Desc(field).Find(_NodeRevision)
	return _NodeRevision, err
}

// GetNodeRevisionsViaRevisionTimestamp Get NodeRevisions via RevisionTimestamp
func GetNodeRevisionsViaRevisionTimestamp(offset int, limit int, RevisionTimestamp_ int, field string) (*[]*NodeRevision, error) {
	var _NodeRevision = new([]*NodeRevision)
	err := Engine.Table("node_revision").Where("revision_timestamp = ?", RevisionTimestamp_).Limit(limit, offset).Desc(field).Find(_NodeRevision)
	return _NodeRevision, err
}

// GetNodeRevisionsViaRevisionUid Get NodeRevisions via RevisionUid
func GetNodeRevisionsViaRevisionUid(offset int, limit int, RevisionUid_ int, field string) (*[]*NodeRevision, error) {
	var _NodeRevision = new([]*NodeRevision)
	err := Engine.Table("node_revision").Where("revision_uid = ?", RevisionUid_).Limit(limit, offset).Desc(field).Find(_NodeRevision)
	return _NodeRevision, err
}

// GetNodeRevisionsViaRevisionLog Get NodeRevisions via RevisionLog
func GetNodeRevisionsViaRevisionLog(offset int, limit int, RevisionLog_ string, field string) (*[]*NodeRevision, error) {
	var _NodeRevision = new([]*NodeRevision)
	err := Engine.Table("node_revision").Where("revision_log = ?", RevisionLog_).Limit(limit, offset).Desc(field).Find(_NodeRevision)
	return _NodeRevision, err
}

// HasNodeRevisionViaNid Has NodeRevision via Nid
func HasNodeRevisionViaNid(iNid int64) bool {
	if has, err := Engine.Where("nid = ?", iNid).Get(new(NodeRevision)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasNodeRevisionViaVid Has NodeRevision via Vid
func HasNodeRevisionViaVid(iVid int) bool {
	if has, err := Engine.Where("vid = ?", iVid).Get(new(NodeRevision)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasNodeRevisionViaLangcode Has NodeRevision via Langcode
func HasNodeRevisionViaLangcode(iLangcode string) bool {
	if has, err := Engine.Where("langcode = ?", iLangcode).Get(new(NodeRevision)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasNodeRevisionViaRevisionTimestamp Has NodeRevision via RevisionTimestamp
func HasNodeRevisionViaRevisionTimestamp(iRevisionTimestamp int) bool {
	if has, err := Engine.Where("revision_timestamp = ?", iRevisionTimestamp).Get(new(NodeRevision)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasNodeRevisionViaRevisionUid Has NodeRevision via RevisionUid
func HasNodeRevisionViaRevisionUid(iRevisionUid int) bool {
	if has, err := Engine.Where("revision_uid = ?", iRevisionUid).Get(new(NodeRevision)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasNodeRevisionViaRevisionLog Has NodeRevision via RevisionLog
func HasNodeRevisionViaRevisionLog(iRevisionLog string) bool {
	if has, err := Engine.Where("revision_log = ?", iRevisionLog).Get(new(NodeRevision)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// GetNodeRevisionViaNid Get NodeRevision via Nid
func GetNodeRevisionViaNid(iNid int64) (*NodeRevision, error) {
	var _NodeRevision = &NodeRevision{Nid: iNid}
	has, err := Engine.Get(_NodeRevision)
	if has {
		return _NodeRevision, err
	} else {
		return nil, err
	}
}

// GetNodeRevisionViaVid Get NodeRevision via Vid
func GetNodeRevisionViaVid(iVid int) (*NodeRevision, error) {
	var _NodeRevision = &NodeRevision{Vid: iVid}
	has, err := Engine.Get(_NodeRevision)
	if has {
		return _NodeRevision, err
	} else {
		return nil, err
	}
}

// GetNodeRevisionViaLangcode Get NodeRevision via Langcode
func GetNodeRevisionViaLangcode(iLangcode string) (*NodeRevision, error) {
	var _NodeRevision = &NodeRevision{Langcode: iLangcode}
	has, err := Engine.Get(_NodeRevision)
	if has {
		return _NodeRevision, err
	} else {
		return nil, err
	}
}

// GetNodeRevisionViaRevisionTimestamp Get NodeRevision via RevisionTimestamp
func GetNodeRevisionViaRevisionTimestamp(iRevisionTimestamp int) (*NodeRevision, error) {
	var _NodeRevision = &NodeRevision{RevisionTimestamp: iRevisionTimestamp}
	has, err := Engine.Get(_NodeRevision)
	if has {
		return _NodeRevision, err
	} else {
		return nil, err
	}
}

// GetNodeRevisionViaRevisionUid Get NodeRevision via RevisionUid
func GetNodeRevisionViaRevisionUid(iRevisionUid int) (*NodeRevision, error) {
	var _NodeRevision = &NodeRevision{RevisionUid: iRevisionUid}
	has, err := Engine.Get(_NodeRevision)
	if has {
		return _NodeRevision, err
	} else {
		return nil, err
	}
}

// GetNodeRevisionViaRevisionLog Get NodeRevision via RevisionLog
func GetNodeRevisionViaRevisionLog(iRevisionLog string) (*NodeRevision, error) {
	var _NodeRevision = &NodeRevision{RevisionLog: iRevisionLog}
	has, err := Engine.Get(_NodeRevision)
	if has {
		return _NodeRevision, err
	} else {
		return nil, err
	}
}

// SetNodeRevisionViaNid Set NodeRevision via Nid
func SetNodeRevisionViaNid(iNid int64, node_revision *NodeRevision) (int64, error) {
	node_revision.Nid = iNid
	return Engine.Insert(node_revision)
}

// SetNodeRevisionViaVid Set NodeRevision via Vid
func SetNodeRevisionViaVid(iVid int, node_revision *NodeRevision) (int64, error) {
	node_revision.Vid = iVid
	return Engine.Insert(node_revision)
}

// SetNodeRevisionViaLangcode Set NodeRevision via Langcode
func SetNodeRevisionViaLangcode(iLangcode string, node_revision *NodeRevision) (int64, error) {
	node_revision.Langcode = iLangcode
	return Engine.Insert(node_revision)
}

// SetNodeRevisionViaRevisionTimestamp Set NodeRevision via RevisionTimestamp
func SetNodeRevisionViaRevisionTimestamp(iRevisionTimestamp int, node_revision *NodeRevision) (int64, error) {
	node_revision.RevisionTimestamp = iRevisionTimestamp
	return Engine.Insert(node_revision)
}

// SetNodeRevisionViaRevisionUid Set NodeRevision via RevisionUid
func SetNodeRevisionViaRevisionUid(iRevisionUid int, node_revision *NodeRevision) (int64, error) {
	node_revision.RevisionUid = iRevisionUid
	return Engine.Insert(node_revision)
}

// SetNodeRevisionViaRevisionLog Set NodeRevision via RevisionLog
func SetNodeRevisionViaRevisionLog(iRevisionLog string, node_revision *NodeRevision) (int64, error) {
	node_revision.RevisionLog = iRevisionLog
	return Engine.Insert(node_revision)
}

// AddNodeRevision Add NodeRevision via all columns
func AddNodeRevision(iNid int64, iVid int, iLangcode string, iRevisionTimestamp int, iRevisionUid int, iRevisionLog string) error {
	_NodeRevision := &NodeRevision{Nid: iNid, Vid: iVid, Langcode: iLangcode, RevisionTimestamp: iRevisionTimestamp, RevisionUid: iRevisionUid, RevisionLog: iRevisionLog}
	if _, err := Engine.Insert(_NodeRevision); err != nil {
		return err
	} else {
		return nil
	}
}

// PostNodeRevision Post NodeRevision via iNodeRevision
func PostNodeRevision(iNodeRevision *NodeRevision) (int64, error) {
	_, err := Engine.Insert(iNodeRevision)
	return iNodeRevision.Nid, err
}

// PutNodeRevision Put NodeRevision
func PutNodeRevision(iNodeRevision *NodeRevision) (int64, error) {
	_, err := Engine.Id(iNodeRevision.Nid).Update(iNodeRevision)
	return iNodeRevision.Nid, err
}

// PutNodeRevisionViaNid Put NodeRevision via Nid
func PutNodeRevisionViaNid(Nid_ int64, iNodeRevision *NodeRevision) (int64, error) {
	row, err := Engine.Update(iNodeRevision, &NodeRevision{Nid: Nid_})
	return row, err
}

// PutNodeRevisionViaVid Put NodeRevision via Vid
func PutNodeRevisionViaVid(Vid_ int, iNodeRevision *NodeRevision) (int64, error) {
	row, err := Engine.Update(iNodeRevision, &NodeRevision{Vid: Vid_})
	return row, err
}

// PutNodeRevisionViaLangcode Put NodeRevision via Langcode
func PutNodeRevisionViaLangcode(Langcode_ string, iNodeRevision *NodeRevision) (int64, error) {
	row, err := Engine.Update(iNodeRevision, &NodeRevision{Langcode: Langcode_})
	return row, err
}

// PutNodeRevisionViaRevisionTimestamp Put NodeRevision via RevisionTimestamp
func PutNodeRevisionViaRevisionTimestamp(RevisionTimestamp_ int, iNodeRevision *NodeRevision) (int64, error) {
	row, err := Engine.Update(iNodeRevision, &NodeRevision{RevisionTimestamp: RevisionTimestamp_})
	return row, err
}

// PutNodeRevisionViaRevisionUid Put NodeRevision via RevisionUid
func PutNodeRevisionViaRevisionUid(RevisionUid_ int, iNodeRevision *NodeRevision) (int64, error) {
	row, err := Engine.Update(iNodeRevision, &NodeRevision{RevisionUid: RevisionUid_})
	return row, err
}

// PutNodeRevisionViaRevisionLog Put NodeRevision via RevisionLog
func PutNodeRevisionViaRevisionLog(RevisionLog_ string, iNodeRevision *NodeRevision) (int64, error) {
	row, err := Engine.Update(iNodeRevision, &NodeRevision{RevisionLog: RevisionLog_})
	return row, err
}

// UpdateNodeRevisionViaNid via map[string]interface{}{}
func UpdateNodeRevisionViaNid(iNid int64, iNodeRevisionMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(NodeRevision)).Where("nid = ?", iNid).Update(iNodeRevisionMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateNodeRevisionViaVid via map[string]interface{}{}
func UpdateNodeRevisionViaVid(iVid int, iNodeRevisionMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(NodeRevision)).Where("vid = ?", iVid).Update(iNodeRevisionMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateNodeRevisionViaLangcode via map[string]interface{}{}
func UpdateNodeRevisionViaLangcode(iLangcode string, iNodeRevisionMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(NodeRevision)).Where("langcode = ?", iLangcode).Update(iNodeRevisionMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateNodeRevisionViaRevisionTimestamp via map[string]interface{}{}
func UpdateNodeRevisionViaRevisionTimestamp(iRevisionTimestamp int, iNodeRevisionMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(NodeRevision)).Where("revision_timestamp = ?", iRevisionTimestamp).Update(iNodeRevisionMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateNodeRevisionViaRevisionUid via map[string]interface{}{}
func UpdateNodeRevisionViaRevisionUid(iRevisionUid int, iNodeRevisionMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(NodeRevision)).Where("revision_uid = ?", iRevisionUid).Update(iNodeRevisionMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateNodeRevisionViaRevisionLog via map[string]interface{}{}
func UpdateNodeRevisionViaRevisionLog(iRevisionLog string, iNodeRevisionMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(NodeRevision)).Where("revision_log = ?", iRevisionLog).Update(iNodeRevisionMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// DeleteNodeRevision Delete NodeRevision
func DeleteNodeRevision(iNid int64) (int64, error) {
	row, err := Engine.Id(iNid).Delete(new(NodeRevision))
	return row, err
}

// DeleteNodeRevisionViaNid Delete NodeRevision via Nid
func DeleteNodeRevisionViaNid(iNid int64) (err error) {
	var has bool
	var _NodeRevision = &NodeRevision{Nid: iNid}
	if has, err = Engine.Get(_NodeRevision); (has == true) && (err == nil) {
		if row, err := Engine.Where("nid = ?", iNid).Delete(new(NodeRevision)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteNodeRevisionViaVid Delete NodeRevision via Vid
func DeleteNodeRevisionViaVid(iVid int) (err error) {
	var has bool
	var _NodeRevision = &NodeRevision{Vid: iVid}
	if has, err = Engine.Get(_NodeRevision); (has == true) && (err == nil) {
		if row, err := Engine.Where("vid = ?", iVid).Delete(new(NodeRevision)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteNodeRevisionViaLangcode Delete NodeRevision via Langcode
func DeleteNodeRevisionViaLangcode(iLangcode string) (err error) {
	var has bool
	var _NodeRevision = &NodeRevision{Langcode: iLangcode}
	if has, err = Engine.Get(_NodeRevision); (has == true) && (err == nil) {
		if row, err := Engine.Where("langcode = ?", iLangcode).Delete(new(NodeRevision)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteNodeRevisionViaRevisionTimestamp Delete NodeRevision via RevisionTimestamp
func DeleteNodeRevisionViaRevisionTimestamp(iRevisionTimestamp int) (err error) {
	var has bool
	var _NodeRevision = &NodeRevision{RevisionTimestamp: iRevisionTimestamp}
	if has, err = Engine.Get(_NodeRevision); (has == true) && (err == nil) {
		if row, err := Engine.Where("revision_timestamp = ?", iRevisionTimestamp).Delete(new(NodeRevision)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteNodeRevisionViaRevisionUid Delete NodeRevision via RevisionUid
func DeleteNodeRevisionViaRevisionUid(iRevisionUid int) (err error) {
	var has bool
	var _NodeRevision = &NodeRevision{RevisionUid: iRevisionUid}
	if has, err = Engine.Get(_NodeRevision); (has == true) && (err == nil) {
		if row, err := Engine.Where("revision_uid = ?", iRevisionUid).Delete(new(NodeRevision)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteNodeRevisionViaRevisionLog Delete NodeRevision via RevisionLog
func DeleteNodeRevisionViaRevisionLog(iRevisionLog string) (err error) {
	var has bool
	var _NodeRevision = &NodeRevision{RevisionLog: iRevisionLog}
	if has, err = Engine.Get(_NodeRevision); (has == true) && (err == nil) {
		if row, err := Engine.Where("revision_log = ?", iRevisionLog).Delete(new(NodeRevision)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}
