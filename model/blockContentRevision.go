package model

type BlockContentRevision struct {
	Id          int64  `xorm:"not null index INT(10)"`
	RevisionId  int    `xorm:"not null pk autoincr INT(10)"`
	Langcode    string `xorm:"not null VARCHAR(12)"`
	RevisionLog string `xorm:"LONGTEXT"`
}

// GetBlockContentRevisionsCount BlockContentRevisions' Count
func GetBlockContentRevisionsCount(offset int, limit int) (int64, error) {
	total, err := Engine.Limit(limit, offset).Count(&BlockContentRevision{})
	return total, err
}

// GetBlockContentRevisionCountViaId Get BlockContentRevision via Id
func GetBlockContentRevisionCountViaId(iId int64) int64 {
	n, _ := Engine.Where("id = ?", iId).Count(&BlockContentRevision{Id: iId})
	return n
}

// GetBlockContentRevisionCountViaRevisionId Get BlockContentRevision via RevisionId
func GetBlockContentRevisionCountViaRevisionId(iRevisionId int) int64 {
	n, _ := Engine.Where("revision_id = ?", iRevisionId).Count(&BlockContentRevision{RevisionId: iRevisionId})
	return n
}

// GetBlockContentRevisionCountViaLangcode Get BlockContentRevision via Langcode
func GetBlockContentRevisionCountViaLangcode(iLangcode string) int64 {
	n, _ := Engine.Where("langcode = ?", iLangcode).Count(&BlockContentRevision{Langcode: iLangcode})
	return n
}

// GetBlockContentRevisionCountViaRevisionLog Get BlockContentRevision via RevisionLog
func GetBlockContentRevisionCountViaRevisionLog(iRevisionLog string) int64 {
	n, _ := Engine.Where("revision_log = ?", iRevisionLog).Count(&BlockContentRevision{RevisionLog: iRevisionLog})
	return n
}

// GetBlockContentRevisionsByIdAndRevisionIdAndLangcode Get BlockContentRevisions via IdAndRevisionIdAndLangcode
func GetBlockContentRevisionsByIdAndRevisionIdAndLangcode(offset int, limit int, Id_ int64, RevisionId_ int, Langcode_ string) (*[]*BlockContentRevision, error) {
	var _BlockContentRevision = new([]*BlockContentRevision)
	err := Engine.Table("block_content_revision").Where("id = ? and revision_id = ? and langcode = ?", Id_, RevisionId_, Langcode_).Limit(limit, offset).Find(_BlockContentRevision)
	return _BlockContentRevision, err
}

// GetBlockContentRevisionsByIdAndRevisionIdAndRevisionLog Get BlockContentRevisions via IdAndRevisionIdAndRevisionLog
func GetBlockContentRevisionsByIdAndRevisionIdAndRevisionLog(offset int, limit int, Id_ int64, RevisionId_ int, RevisionLog_ string) (*[]*BlockContentRevision, error) {
	var _BlockContentRevision = new([]*BlockContentRevision)
	err := Engine.Table("block_content_revision").Where("id = ? and revision_id = ? and revision_log = ?", Id_, RevisionId_, RevisionLog_).Limit(limit, offset).Find(_BlockContentRevision)
	return _BlockContentRevision, err
}

// GetBlockContentRevisionsByIdAndLangcodeAndRevisionLog Get BlockContentRevisions via IdAndLangcodeAndRevisionLog
func GetBlockContentRevisionsByIdAndLangcodeAndRevisionLog(offset int, limit int, Id_ int64, Langcode_ string, RevisionLog_ string) (*[]*BlockContentRevision, error) {
	var _BlockContentRevision = new([]*BlockContentRevision)
	err := Engine.Table("block_content_revision").Where("id = ? and langcode = ? and revision_log = ?", Id_, Langcode_, RevisionLog_).Limit(limit, offset).Find(_BlockContentRevision)
	return _BlockContentRevision, err
}

// GetBlockContentRevisionsByRevisionIdAndLangcodeAndRevisionLog Get BlockContentRevisions via RevisionIdAndLangcodeAndRevisionLog
func GetBlockContentRevisionsByRevisionIdAndLangcodeAndRevisionLog(offset int, limit int, RevisionId_ int, Langcode_ string, RevisionLog_ string) (*[]*BlockContentRevision, error) {
	var _BlockContentRevision = new([]*BlockContentRevision)
	err := Engine.Table("block_content_revision").Where("revision_id = ? and langcode = ? and revision_log = ?", RevisionId_, Langcode_, RevisionLog_).Limit(limit, offset).Find(_BlockContentRevision)
	return _BlockContentRevision, err
}

// GetBlockContentRevisionsByIdAndRevisionId Get BlockContentRevisions via IdAndRevisionId
func GetBlockContentRevisionsByIdAndRevisionId(offset int, limit int, Id_ int64, RevisionId_ int) (*[]*BlockContentRevision, error) {
	var _BlockContentRevision = new([]*BlockContentRevision)
	err := Engine.Table("block_content_revision").Where("id = ? and revision_id = ?", Id_, RevisionId_).Limit(limit, offset).Find(_BlockContentRevision)
	return _BlockContentRevision, err
}

// GetBlockContentRevisionsByIdAndLangcode Get BlockContentRevisions via IdAndLangcode
func GetBlockContentRevisionsByIdAndLangcode(offset int, limit int, Id_ int64, Langcode_ string) (*[]*BlockContentRevision, error) {
	var _BlockContentRevision = new([]*BlockContentRevision)
	err := Engine.Table("block_content_revision").Where("id = ? and langcode = ?", Id_, Langcode_).Limit(limit, offset).Find(_BlockContentRevision)
	return _BlockContentRevision, err
}

// GetBlockContentRevisionsByIdAndRevisionLog Get BlockContentRevisions via IdAndRevisionLog
func GetBlockContentRevisionsByIdAndRevisionLog(offset int, limit int, Id_ int64, RevisionLog_ string) (*[]*BlockContentRevision, error) {
	var _BlockContentRevision = new([]*BlockContentRevision)
	err := Engine.Table("block_content_revision").Where("id = ? and revision_log = ?", Id_, RevisionLog_).Limit(limit, offset).Find(_BlockContentRevision)
	return _BlockContentRevision, err
}

// GetBlockContentRevisionsByRevisionIdAndLangcode Get BlockContentRevisions via RevisionIdAndLangcode
func GetBlockContentRevisionsByRevisionIdAndLangcode(offset int, limit int, RevisionId_ int, Langcode_ string) (*[]*BlockContentRevision, error) {
	var _BlockContentRevision = new([]*BlockContentRevision)
	err := Engine.Table("block_content_revision").Where("revision_id = ? and langcode = ?", RevisionId_, Langcode_).Limit(limit, offset).Find(_BlockContentRevision)
	return _BlockContentRevision, err
}

// GetBlockContentRevisionsByRevisionIdAndRevisionLog Get BlockContentRevisions via RevisionIdAndRevisionLog
func GetBlockContentRevisionsByRevisionIdAndRevisionLog(offset int, limit int, RevisionId_ int, RevisionLog_ string) (*[]*BlockContentRevision, error) {
	var _BlockContentRevision = new([]*BlockContentRevision)
	err := Engine.Table("block_content_revision").Where("revision_id = ? and revision_log = ?", RevisionId_, RevisionLog_).Limit(limit, offset).Find(_BlockContentRevision)
	return _BlockContentRevision, err
}

// GetBlockContentRevisionsByLangcodeAndRevisionLog Get BlockContentRevisions via LangcodeAndRevisionLog
func GetBlockContentRevisionsByLangcodeAndRevisionLog(offset int, limit int, Langcode_ string, RevisionLog_ string) (*[]*BlockContentRevision, error) {
	var _BlockContentRevision = new([]*BlockContentRevision)
	err := Engine.Table("block_content_revision").Where("langcode = ? and revision_log = ?", Langcode_, RevisionLog_).Limit(limit, offset).Find(_BlockContentRevision)
	return _BlockContentRevision, err
}

// GetBlockContentRevisions Get BlockContentRevisions via field
func GetBlockContentRevisions(offset int, limit int, field string) (*[]*BlockContentRevision, error) {
	var _BlockContentRevision = new([]*BlockContentRevision)
	err := Engine.Table("block_content_revision").Limit(limit, offset).Desc(field).Find(_BlockContentRevision)
	return _BlockContentRevision, err
}

// GetBlockContentRevisionsViaId Get BlockContentRevisions via Id
func GetBlockContentRevisionsViaId(offset int, limit int, Id_ int64, field string) (*[]*BlockContentRevision, error) {
	var _BlockContentRevision = new([]*BlockContentRevision)
	err := Engine.Table("block_content_revision").Where("id = ?", Id_).Limit(limit, offset).Desc(field).Find(_BlockContentRevision)
	return _BlockContentRevision, err
}

// GetBlockContentRevisionsViaRevisionId Get BlockContentRevisions via RevisionId
func GetBlockContentRevisionsViaRevisionId(offset int, limit int, RevisionId_ int, field string) (*[]*BlockContentRevision, error) {
	var _BlockContentRevision = new([]*BlockContentRevision)
	err := Engine.Table("block_content_revision").Where("revision_id = ?", RevisionId_).Limit(limit, offset).Desc(field).Find(_BlockContentRevision)
	return _BlockContentRevision, err
}

// GetBlockContentRevisionsViaLangcode Get BlockContentRevisions via Langcode
func GetBlockContentRevisionsViaLangcode(offset int, limit int, Langcode_ string, field string) (*[]*BlockContentRevision, error) {
	var _BlockContentRevision = new([]*BlockContentRevision)
	err := Engine.Table("block_content_revision").Where("langcode = ?", Langcode_).Limit(limit, offset).Desc(field).Find(_BlockContentRevision)
	return _BlockContentRevision, err
}

// GetBlockContentRevisionsViaRevisionLog Get BlockContentRevisions via RevisionLog
func GetBlockContentRevisionsViaRevisionLog(offset int, limit int, RevisionLog_ string, field string) (*[]*BlockContentRevision, error) {
	var _BlockContentRevision = new([]*BlockContentRevision)
	err := Engine.Table("block_content_revision").Where("revision_log = ?", RevisionLog_).Limit(limit, offset).Desc(field).Find(_BlockContentRevision)
	return _BlockContentRevision, err
}

// HasBlockContentRevisionViaId Has BlockContentRevision via Id
func HasBlockContentRevisionViaId(iId int64) bool {
	if has, err := Engine.Where("id = ?", iId).Get(new(BlockContentRevision)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasBlockContentRevisionViaRevisionId Has BlockContentRevision via RevisionId
func HasBlockContentRevisionViaRevisionId(iRevisionId int) bool {
	if has, err := Engine.Where("revision_id = ?", iRevisionId).Get(new(BlockContentRevision)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasBlockContentRevisionViaLangcode Has BlockContentRevision via Langcode
func HasBlockContentRevisionViaLangcode(iLangcode string) bool {
	if has, err := Engine.Where("langcode = ?", iLangcode).Get(new(BlockContentRevision)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasBlockContentRevisionViaRevisionLog Has BlockContentRevision via RevisionLog
func HasBlockContentRevisionViaRevisionLog(iRevisionLog string) bool {
	if has, err := Engine.Where("revision_log = ?", iRevisionLog).Get(new(BlockContentRevision)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// GetBlockContentRevisionViaId Get BlockContentRevision via Id
func GetBlockContentRevisionViaId(iId int64) (*BlockContentRevision, error) {
	var _BlockContentRevision = &BlockContentRevision{Id: iId}
	has, err := Engine.Get(_BlockContentRevision)
	if has {
		return _BlockContentRevision, err
	} else {
		return nil, err
	}
}

// GetBlockContentRevisionViaRevisionId Get BlockContentRevision via RevisionId
func GetBlockContentRevisionViaRevisionId(iRevisionId int) (*BlockContentRevision, error) {
	var _BlockContentRevision = &BlockContentRevision{RevisionId: iRevisionId}
	has, err := Engine.Get(_BlockContentRevision)
	if has {
		return _BlockContentRevision, err
	} else {
		return nil, err
	}
}

// GetBlockContentRevisionViaLangcode Get BlockContentRevision via Langcode
func GetBlockContentRevisionViaLangcode(iLangcode string) (*BlockContentRevision, error) {
	var _BlockContentRevision = &BlockContentRevision{Langcode: iLangcode}
	has, err := Engine.Get(_BlockContentRevision)
	if has {
		return _BlockContentRevision, err
	} else {
		return nil, err
	}
}

// GetBlockContentRevisionViaRevisionLog Get BlockContentRevision via RevisionLog
func GetBlockContentRevisionViaRevisionLog(iRevisionLog string) (*BlockContentRevision, error) {
	var _BlockContentRevision = &BlockContentRevision{RevisionLog: iRevisionLog}
	has, err := Engine.Get(_BlockContentRevision)
	if has {
		return _BlockContentRevision, err
	} else {
		return nil, err
	}
}

// SetBlockContentRevisionViaId Set BlockContentRevision via Id
func SetBlockContentRevisionViaId(iId int64, block_content_revision *BlockContentRevision) (int64, error) {
	block_content_revision.Id = iId
	return Engine.Insert(block_content_revision)
}

// SetBlockContentRevisionViaRevisionId Set BlockContentRevision via RevisionId
func SetBlockContentRevisionViaRevisionId(iRevisionId int, block_content_revision *BlockContentRevision) (int64, error) {
	block_content_revision.RevisionId = iRevisionId
	return Engine.Insert(block_content_revision)
}

// SetBlockContentRevisionViaLangcode Set BlockContentRevision via Langcode
func SetBlockContentRevisionViaLangcode(iLangcode string, block_content_revision *BlockContentRevision) (int64, error) {
	block_content_revision.Langcode = iLangcode
	return Engine.Insert(block_content_revision)
}

// SetBlockContentRevisionViaRevisionLog Set BlockContentRevision via RevisionLog
func SetBlockContentRevisionViaRevisionLog(iRevisionLog string, block_content_revision *BlockContentRevision) (int64, error) {
	block_content_revision.RevisionLog = iRevisionLog
	return Engine.Insert(block_content_revision)
}

// AddBlockContentRevision Add BlockContentRevision via all columns
func AddBlockContentRevision(iId int64, iRevisionId int, iLangcode string, iRevisionLog string) error {
	_BlockContentRevision := &BlockContentRevision{Id: iId, RevisionId: iRevisionId, Langcode: iLangcode, RevisionLog: iRevisionLog}
	if _, err := Engine.Insert(_BlockContentRevision); err != nil {
		return err
	} else {
		return nil
	}
}

// PostBlockContentRevision Post BlockContentRevision via iBlockContentRevision
func PostBlockContentRevision(iBlockContentRevision *BlockContentRevision) (int64, error) {
	_, err := Engine.Insert(iBlockContentRevision)
	return iBlockContentRevision.Id, err
}

// PutBlockContentRevision Put BlockContentRevision
func PutBlockContentRevision(iBlockContentRevision *BlockContentRevision) (int64, error) {
	_, err := Engine.Id(iBlockContentRevision.Id).Update(iBlockContentRevision)
	return iBlockContentRevision.Id, err
}

// PutBlockContentRevisionViaId Put BlockContentRevision via Id
func PutBlockContentRevisionViaId(Id_ int64, iBlockContentRevision *BlockContentRevision) (int64, error) {
	row, err := Engine.Update(iBlockContentRevision, &BlockContentRevision{Id: Id_})
	return row, err
}

// PutBlockContentRevisionViaRevisionId Put BlockContentRevision via RevisionId
func PutBlockContentRevisionViaRevisionId(RevisionId_ int, iBlockContentRevision *BlockContentRevision) (int64, error) {
	row, err := Engine.Update(iBlockContentRevision, &BlockContentRevision{RevisionId: RevisionId_})
	return row, err
}

// PutBlockContentRevisionViaLangcode Put BlockContentRevision via Langcode
func PutBlockContentRevisionViaLangcode(Langcode_ string, iBlockContentRevision *BlockContentRevision) (int64, error) {
	row, err := Engine.Update(iBlockContentRevision, &BlockContentRevision{Langcode: Langcode_})
	return row, err
}

// PutBlockContentRevisionViaRevisionLog Put BlockContentRevision via RevisionLog
func PutBlockContentRevisionViaRevisionLog(RevisionLog_ string, iBlockContentRevision *BlockContentRevision) (int64, error) {
	row, err := Engine.Update(iBlockContentRevision, &BlockContentRevision{RevisionLog: RevisionLog_})
	return row, err
}

// UpdateBlockContentRevisionViaId via map[string]interface{}{}
func UpdateBlockContentRevisionViaId(iId int64, iBlockContentRevisionMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(BlockContentRevision)).Where("id = ?", iId).Update(iBlockContentRevisionMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateBlockContentRevisionViaRevisionId via map[string]interface{}{}
func UpdateBlockContentRevisionViaRevisionId(iRevisionId int, iBlockContentRevisionMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(BlockContentRevision)).Where("revision_id = ?", iRevisionId).Update(iBlockContentRevisionMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateBlockContentRevisionViaLangcode via map[string]interface{}{}
func UpdateBlockContentRevisionViaLangcode(iLangcode string, iBlockContentRevisionMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(BlockContentRevision)).Where("langcode = ?", iLangcode).Update(iBlockContentRevisionMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateBlockContentRevisionViaRevisionLog via map[string]interface{}{}
func UpdateBlockContentRevisionViaRevisionLog(iRevisionLog string, iBlockContentRevisionMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(BlockContentRevision)).Where("revision_log = ?", iRevisionLog).Update(iBlockContentRevisionMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// DeleteBlockContentRevision Delete BlockContentRevision
func DeleteBlockContentRevision(iId int64) (int64, error) {
	row, err := Engine.Id(iId).Delete(new(BlockContentRevision))
	return row, err
}

// DeleteBlockContentRevisionViaId Delete BlockContentRevision via Id
func DeleteBlockContentRevisionViaId(iId int64) (err error) {
	var has bool
	var _BlockContentRevision = &BlockContentRevision{Id: iId}
	if has, err = Engine.Get(_BlockContentRevision); (has == true) && (err == nil) {
		if row, err := Engine.Where("id = ?", iId).Delete(new(BlockContentRevision)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteBlockContentRevisionViaRevisionId Delete BlockContentRevision via RevisionId
func DeleteBlockContentRevisionViaRevisionId(iRevisionId int) (err error) {
	var has bool
	var _BlockContentRevision = &BlockContentRevision{RevisionId: iRevisionId}
	if has, err = Engine.Get(_BlockContentRevision); (has == true) && (err == nil) {
		if row, err := Engine.Where("revision_id = ?", iRevisionId).Delete(new(BlockContentRevision)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteBlockContentRevisionViaLangcode Delete BlockContentRevision via Langcode
func DeleteBlockContentRevisionViaLangcode(iLangcode string) (err error) {
	var has bool
	var _BlockContentRevision = &BlockContentRevision{Langcode: iLangcode}
	if has, err = Engine.Get(_BlockContentRevision); (has == true) && (err == nil) {
		if row, err := Engine.Where("langcode = ?", iLangcode).Delete(new(BlockContentRevision)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteBlockContentRevisionViaRevisionLog Delete BlockContentRevision via RevisionLog
func DeleteBlockContentRevisionViaRevisionLog(iRevisionLog string) (err error) {
	var has bool
	var _BlockContentRevision = &BlockContentRevision{RevisionLog: iRevisionLog}
	if has, err = Engine.Get(_BlockContentRevision); (has == true) && (err == nil) {
		if row, err := Engine.Where("revision_log = ?", iRevisionLog).Delete(new(BlockContentRevision)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}
