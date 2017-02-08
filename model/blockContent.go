package model

type BlockContent struct {
	Id         int64  `xorm:"not null pk autoincr INT(10)"`
	RevisionId int    `xorm:"unique INT(10)"`
	Type       string `xorm:"not null index VARCHAR(32)"`
	Uuid       string `xorm:"not null unique VARCHAR(128)"`
	Langcode   string `xorm:"not null VARCHAR(12)"`
}

// GetBlockContentsCount BlockContents' Count
func GetBlockContentsCount(offset int, limit int) (int64, error) {
	total, err := Engine.Limit(limit, offset).Count(&BlockContent{})
	return total, err
}

// GetBlockContentCountViaId Get BlockContent via Id
func GetBlockContentCountViaId(iId int64) int64 {
	n, _ := Engine.Where("id = ?", iId).Count(&BlockContent{Id: iId})
	return n
}

// GetBlockContentCountViaRevisionId Get BlockContent via RevisionId
func GetBlockContentCountViaRevisionId(iRevisionId int) int64 {
	n, _ := Engine.Where("revision_id = ?", iRevisionId).Count(&BlockContent{RevisionId: iRevisionId})
	return n
}

// GetBlockContentCountViaType Get BlockContent via Type
func GetBlockContentCountViaType(iType string) int64 {
	n, _ := Engine.Where("type = ?", iType).Count(&BlockContent{Type: iType})
	return n
}

// GetBlockContentCountViaUuid Get BlockContent via Uuid
func GetBlockContentCountViaUuid(iUuid string) int64 {
	n, _ := Engine.Where("uuid = ?", iUuid).Count(&BlockContent{Uuid: iUuid})
	return n
}

// GetBlockContentCountViaLangcode Get BlockContent via Langcode
func GetBlockContentCountViaLangcode(iLangcode string) int64 {
	n, _ := Engine.Where("langcode = ?", iLangcode).Count(&BlockContent{Langcode: iLangcode})
	return n
}

// GetBlockContentsByIdAndRevisionIdAndType Get BlockContents via IdAndRevisionIdAndType
func GetBlockContentsByIdAndRevisionIdAndType(offset int, limit int, Id_ int64, RevisionId_ int, Type_ string) (*[]*BlockContent, error) {
	var _BlockContent = new([]*BlockContent)
	err := Engine.Table("block_content").Where("id = ? and revision_id = ? and type = ?", Id_, RevisionId_, Type_).Limit(limit, offset).Find(_BlockContent)
	return _BlockContent, err
}

// GetBlockContentsByIdAndRevisionIdAndUuid Get BlockContents via IdAndRevisionIdAndUuid
func GetBlockContentsByIdAndRevisionIdAndUuid(offset int, limit int, Id_ int64, RevisionId_ int, Uuid_ string) (*[]*BlockContent, error) {
	var _BlockContent = new([]*BlockContent)
	err := Engine.Table("block_content").Where("id = ? and revision_id = ? and uuid = ?", Id_, RevisionId_, Uuid_).Limit(limit, offset).Find(_BlockContent)
	return _BlockContent, err
}

// GetBlockContentsByIdAndRevisionIdAndLangcode Get BlockContents via IdAndRevisionIdAndLangcode
func GetBlockContentsByIdAndRevisionIdAndLangcode(offset int, limit int, Id_ int64, RevisionId_ int, Langcode_ string) (*[]*BlockContent, error) {
	var _BlockContent = new([]*BlockContent)
	err := Engine.Table("block_content").Where("id = ? and revision_id = ? and langcode = ?", Id_, RevisionId_, Langcode_).Limit(limit, offset).Find(_BlockContent)
	return _BlockContent, err
}

// GetBlockContentsByIdAndTypeAndUuid Get BlockContents via IdAndTypeAndUuid
func GetBlockContentsByIdAndTypeAndUuid(offset int, limit int, Id_ int64, Type_ string, Uuid_ string) (*[]*BlockContent, error) {
	var _BlockContent = new([]*BlockContent)
	err := Engine.Table("block_content").Where("id = ? and type = ? and uuid = ?", Id_, Type_, Uuid_).Limit(limit, offset).Find(_BlockContent)
	return _BlockContent, err
}

// GetBlockContentsByIdAndTypeAndLangcode Get BlockContents via IdAndTypeAndLangcode
func GetBlockContentsByIdAndTypeAndLangcode(offset int, limit int, Id_ int64, Type_ string, Langcode_ string) (*[]*BlockContent, error) {
	var _BlockContent = new([]*BlockContent)
	err := Engine.Table("block_content").Where("id = ? and type = ? and langcode = ?", Id_, Type_, Langcode_).Limit(limit, offset).Find(_BlockContent)
	return _BlockContent, err
}

// GetBlockContentsByIdAndUuidAndLangcode Get BlockContents via IdAndUuidAndLangcode
func GetBlockContentsByIdAndUuidAndLangcode(offset int, limit int, Id_ int64, Uuid_ string, Langcode_ string) (*[]*BlockContent, error) {
	var _BlockContent = new([]*BlockContent)
	err := Engine.Table("block_content").Where("id = ? and uuid = ? and langcode = ?", Id_, Uuid_, Langcode_).Limit(limit, offset).Find(_BlockContent)
	return _BlockContent, err
}

// GetBlockContentsByRevisionIdAndTypeAndUuid Get BlockContents via RevisionIdAndTypeAndUuid
func GetBlockContentsByRevisionIdAndTypeAndUuid(offset int, limit int, RevisionId_ int, Type_ string, Uuid_ string) (*[]*BlockContent, error) {
	var _BlockContent = new([]*BlockContent)
	err := Engine.Table("block_content").Where("revision_id = ? and type = ? and uuid = ?", RevisionId_, Type_, Uuid_).Limit(limit, offset).Find(_BlockContent)
	return _BlockContent, err
}

// GetBlockContentsByRevisionIdAndTypeAndLangcode Get BlockContents via RevisionIdAndTypeAndLangcode
func GetBlockContentsByRevisionIdAndTypeAndLangcode(offset int, limit int, RevisionId_ int, Type_ string, Langcode_ string) (*[]*BlockContent, error) {
	var _BlockContent = new([]*BlockContent)
	err := Engine.Table("block_content").Where("revision_id = ? and type = ? and langcode = ?", RevisionId_, Type_, Langcode_).Limit(limit, offset).Find(_BlockContent)
	return _BlockContent, err
}

// GetBlockContentsByRevisionIdAndUuidAndLangcode Get BlockContents via RevisionIdAndUuidAndLangcode
func GetBlockContentsByRevisionIdAndUuidAndLangcode(offset int, limit int, RevisionId_ int, Uuid_ string, Langcode_ string) (*[]*BlockContent, error) {
	var _BlockContent = new([]*BlockContent)
	err := Engine.Table("block_content").Where("revision_id = ? and uuid = ? and langcode = ?", RevisionId_, Uuid_, Langcode_).Limit(limit, offset).Find(_BlockContent)
	return _BlockContent, err
}

// GetBlockContentsByTypeAndUuidAndLangcode Get BlockContents via TypeAndUuidAndLangcode
func GetBlockContentsByTypeAndUuidAndLangcode(offset int, limit int, Type_ string, Uuid_ string, Langcode_ string) (*[]*BlockContent, error) {
	var _BlockContent = new([]*BlockContent)
	err := Engine.Table("block_content").Where("type = ? and uuid = ? and langcode = ?", Type_, Uuid_, Langcode_).Limit(limit, offset).Find(_BlockContent)
	return _BlockContent, err
}

// GetBlockContentsByIdAndRevisionId Get BlockContents via IdAndRevisionId
func GetBlockContentsByIdAndRevisionId(offset int, limit int, Id_ int64, RevisionId_ int) (*[]*BlockContent, error) {
	var _BlockContent = new([]*BlockContent)
	err := Engine.Table("block_content").Where("id = ? and revision_id = ?", Id_, RevisionId_).Limit(limit, offset).Find(_BlockContent)
	return _BlockContent, err
}

// GetBlockContentsByIdAndType Get BlockContents via IdAndType
func GetBlockContentsByIdAndType(offset int, limit int, Id_ int64, Type_ string) (*[]*BlockContent, error) {
	var _BlockContent = new([]*BlockContent)
	err := Engine.Table("block_content").Where("id = ? and type = ?", Id_, Type_).Limit(limit, offset).Find(_BlockContent)
	return _BlockContent, err
}

// GetBlockContentsByIdAndUuid Get BlockContents via IdAndUuid
func GetBlockContentsByIdAndUuid(offset int, limit int, Id_ int64, Uuid_ string) (*[]*BlockContent, error) {
	var _BlockContent = new([]*BlockContent)
	err := Engine.Table("block_content").Where("id = ? and uuid = ?", Id_, Uuid_).Limit(limit, offset).Find(_BlockContent)
	return _BlockContent, err
}

// GetBlockContentsByIdAndLangcode Get BlockContents via IdAndLangcode
func GetBlockContentsByIdAndLangcode(offset int, limit int, Id_ int64, Langcode_ string) (*[]*BlockContent, error) {
	var _BlockContent = new([]*BlockContent)
	err := Engine.Table("block_content").Where("id = ? and langcode = ?", Id_, Langcode_).Limit(limit, offset).Find(_BlockContent)
	return _BlockContent, err
}

// GetBlockContentsByRevisionIdAndType Get BlockContents via RevisionIdAndType
func GetBlockContentsByRevisionIdAndType(offset int, limit int, RevisionId_ int, Type_ string) (*[]*BlockContent, error) {
	var _BlockContent = new([]*BlockContent)
	err := Engine.Table("block_content").Where("revision_id = ? and type = ?", RevisionId_, Type_).Limit(limit, offset).Find(_BlockContent)
	return _BlockContent, err
}

// GetBlockContentsByRevisionIdAndUuid Get BlockContents via RevisionIdAndUuid
func GetBlockContentsByRevisionIdAndUuid(offset int, limit int, RevisionId_ int, Uuid_ string) (*[]*BlockContent, error) {
	var _BlockContent = new([]*BlockContent)
	err := Engine.Table("block_content").Where("revision_id = ? and uuid = ?", RevisionId_, Uuid_).Limit(limit, offset).Find(_BlockContent)
	return _BlockContent, err
}

// GetBlockContentsByRevisionIdAndLangcode Get BlockContents via RevisionIdAndLangcode
func GetBlockContentsByRevisionIdAndLangcode(offset int, limit int, RevisionId_ int, Langcode_ string) (*[]*BlockContent, error) {
	var _BlockContent = new([]*BlockContent)
	err := Engine.Table("block_content").Where("revision_id = ? and langcode = ?", RevisionId_, Langcode_).Limit(limit, offset).Find(_BlockContent)
	return _BlockContent, err
}

// GetBlockContentsByTypeAndUuid Get BlockContents via TypeAndUuid
func GetBlockContentsByTypeAndUuid(offset int, limit int, Type_ string, Uuid_ string) (*[]*BlockContent, error) {
	var _BlockContent = new([]*BlockContent)
	err := Engine.Table("block_content").Where("type = ? and uuid = ?", Type_, Uuid_).Limit(limit, offset).Find(_BlockContent)
	return _BlockContent, err
}

// GetBlockContentsByTypeAndLangcode Get BlockContents via TypeAndLangcode
func GetBlockContentsByTypeAndLangcode(offset int, limit int, Type_ string, Langcode_ string) (*[]*BlockContent, error) {
	var _BlockContent = new([]*BlockContent)
	err := Engine.Table("block_content").Where("type = ? and langcode = ?", Type_, Langcode_).Limit(limit, offset).Find(_BlockContent)
	return _BlockContent, err
}

// GetBlockContentsByUuidAndLangcode Get BlockContents via UuidAndLangcode
func GetBlockContentsByUuidAndLangcode(offset int, limit int, Uuid_ string, Langcode_ string) (*[]*BlockContent, error) {
	var _BlockContent = new([]*BlockContent)
	err := Engine.Table("block_content").Where("uuid = ? and langcode = ?", Uuid_, Langcode_).Limit(limit, offset).Find(_BlockContent)
	return _BlockContent, err
}

// GetBlockContents Get BlockContents via field
func GetBlockContents(offset int, limit int, field string) (*[]*BlockContent, error) {
	var _BlockContent = new([]*BlockContent)
	err := Engine.Table("block_content").Limit(limit, offset).Desc(field).Find(_BlockContent)
	return _BlockContent, err
}

// GetBlockContentsViaId Get BlockContents via Id
func GetBlockContentsViaId(offset int, limit int, Id_ int64, field string) (*[]*BlockContent, error) {
	var _BlockContent = new([]*BlockContent)
	err := Engine.Table("block_content").Where("id = ?", Id_).Limit(limit, offset).Desc(field).Find(_BlockContent)
	return _BlockContent, err
}

// GetBlockContentsViaRevisionId Get BlockContents via RevisionId
func GetBlockContentsViaRevisionId(offset int, limit int, RevisionId_ int, field string) (*[]*BlockContent, error) {
	var _BlockContent = new([]*BlockContent)
	err := Engine.Table("block_content").Where("revision_id = ?", RevisionId_).Limit(limit, offset).Desc(field).Find(_BlockContent)
	return _BlockContent, err
}

// GetBlockContentsViaType Get BlockContents via Type
func GetBlockContentsViaType(offset int, limit int, Type_ string, field string) (*[]*BlockContent, error) {
	var _BlockContent = new([]*BlockContent)
	err := Engine.Table("block_content").Where("type = ?", Type_).Limit(limit, offset).Desc(field).Find(_BlockContent)
	return _BlockContent, err
}

// GetBlockContentsViaUuid Get BlockContents via Uuid
func GetBlockContentsViaUuid(offset int, limit int, Uuid_ string, field string) (*[]*BlockContent, error) {
	var _BlockContent = new([]*BlockContent)
	err := Engine.Table("block_content").Where("uuid = ?", Uuid_).Limit(limit, offset).Desc(field).Find(_BlockContent)
	return _BlockContent, err
}

// GetBlockContentsViaLangcode Get BlockContents via Langcode
func GetBlockContentsViaLangcode(offset int, limit int, Langcode_ string, field string) (*[]*BlockContent, error) {
	var _BlockContent = new([]*BlockContent)
	err := Engine.Table("block_content").Where("langcode = ?", Langcode_).Limit(limit, offset).Desc(field).Find(_BlockContent)
	return _BlockContent, err
}

// HasBlockContentViaId Has BlockContent via Id
func HasBlockContentViaId(iId int64) bool {
	if has, err := Engine.Where("id = ?", iId).Get(new(BlockContent)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasBlockContentViaRevisionId Has BlockContent via RevisionId
func HasBlockContentViaRevisionId(iRevisionId int) bool {
	if has, err := Engine.Where("revision_id = ?", iRevisionId).Get(new(BlockContent)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasBlockContentViaType Has BlockContent via Type
func HasBlockContentViaType(iType string) bool {
	if has, err := Engine.Where("type = ?", iType).Get(new(BlockContent)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasBlockContentViaUuid Has BlockContent via Uuid
func HasBlockContentViaUuid(iUuid string) bool {
	if has, err := Engine.Where("uuid = ?", iUuid).Get(new(BlockContent)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasBlockContentViaLangcode Has BlockContent via Langcode
func HasBlockContentViaLangcode(iLangcode string) bool {
	if has, err := Engine.Where("langcode = ?", iLangcode).Get(new(BlockContent)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// GetBlockContentViaId Get BlockContent via Id
func GetBlockContentViaId(iId int64) (*BlockContent, error) {
	var _BlockContent = &BlockContent{Id: iId}
	has, err := Engine.Get(_BlockContent)
	if has {
		return _BlockContent, err
	} else {
		return nil, err
	}
}

// GetBlockContentViaRevisionId Get BlockContent via RevisionId
func GetBlockContentViaRevisionId(iRevisionId int) (*BlockContent, error) {
	var _BlockContent = &BlockContent{RevisionId: iRevisionId}
	has, err := Engine.Get(_BlockContent)
	if has {
		return _BlockContent, err
	} else {
		return nil, err
	}
}

// GetBlockContentViaType Get BlockContent via Type
func GetBlockContentViaType(iType string) (*BlockContent, error) {
	var _BlockContent = &BlockContent{Type: iType}
	has, err := Engine.Get(_BlockContent)
	if has {
		return _BlockContent, err
	} else {
		return nil, err
	}
}

// GetBlockContentViaUuid Get BlockContent via Uuid
func GetBlockContentViaUuid(iUuid string) (*BlockContent, error) {
	var _BlockContent = &BlockContent{Uuid: iUuid}
	has, err := Engine.Get(_BlockContent)
	if has {
		return _BlockContent, err
	} else {
		return nil, err
	}
}

// GetBlockContentViaLangcode Get BlockContent via Langcode
func GetBlockContentViaLangcode(iLangcode string) (*BlockContent, error) {
	var _BlockContent = &BlockContent{Langcode: iLangcode}
	has, err := Engine.Get(_BlockContent)
	if has {
		return _BlockContent, err
	} else {
		return nil, err
	}
}

// SetBlockContentViaId Set BlockContent via Id
func SetBlockContentViaId(iId int64, block_content *BlockContent) (int64, error) {
	block_content.Id = iId
	return Engine.Insert(block_content)
}

// SetBlockContentViaRevisionId Set BlockContent via RevisionId
func SetBlockContentViaRevisionId(iRevisionId int, block_content *BlockContent) (int64, error) {
	block_content.RevisionId = iRevisionId
	return Engine.Insert(block_content)
}

// SetBlockContentViaType Set BlockContent via Type
func SetBlockContentViaType(iType string, block_content *BlockContent) (int64, error) {
	block_content.Type = iType
	return Engine.Insert(block_content)
}

// SetBlockContentViaUuid Set BlockContent via Uuid
func SetBlockContentViaUuid(iUuid string, block_content *BlockContent) (int64, error) {
	block_content.Uuid = iUuid
	return Engine.Insert(block_content)
}

// SetBlockContentViaLangcode Set BlockContent via Langcode
func SetBlockContentViaLangcode(iLangcode string, block_content *BlockContent) (int64, error) {
	block_content.Langcode = iLangcode
	return Engine.Insert(block_content)
}

// AddBlockContent Add BlockContent via all columns
func AddBlockContent(iId int64, iRevisionId int, iType string, iUuid string, iLangcode string) error {
	_BlockContent := &BlockContent{Id: iId, RevisionId: iRevisionId, Type: iType, Uuid: iUuid, Langcode: iLangcode}
	if _, err := Engine.Insert(_BlockContent); err != nil {
		return err
	} else {
		return nil
	}
}

// PostBlockContent Post BlockContent via iBlockContent
func PostBlockContent(iBlockContent *BlockContent) (int64, error) {
	_, err := Engine.Insert(iBlockContent)
	return iBlockContent.Id, err
}

// PutBlockContent Put BlockContent
func PutBlockContent(iBlockContent *BlockContent) (int64, error) {
	_, err := Engine.Id(iBlockContent.Id).Update(iBlockContent)
	return iBlockContent.Id, err
}

// PutBlockContentViaId Put BlockContent via Id
func PutBlockContentViaId(Id_ int64, iBlockContent *BlockContent) (int64, error) {
	row, err := Engine.Update(iBlockContent, &BlockContent{Id: Id_})
	return row, err
}

// PutBlockContentViaRevisionId Put BlockContent via RevisionId
func PutBlockContentViaRevisionId(RevisionId_ int, iBlockContent *BlockContent) (int64, error) {
	row, err := Engine.Update(iBlockContent, &BlockContent{RevisionId: RevisionId_})
	return row, err
}

// PutBlockContentViaType Put BlockContent via Type
func PutBlockContentViaType(Type_ string, iBlockContent *BlockContent) (int64, error) {
	row, err := Engine.Update(iBlockContent, &BlockContent{Type: Type_})
	return row, err
}

// PutBlockContentViaUuid Put BlockContent via Uuid
func PutBlockContentViaUuid(Uuid_ string, iBlockContent *BlockContent) (int64, error) {
	row, err := Engine.Update(iBlockContent, &BlockContent{Uuid: Uuid_})
	return row, err
}

// PutBlockContentViaLangcode Put BlockContent via Langcode
func PutBlockContentViaLangcode(Langcode_ string, iBlockContent *BlockContent) (int64, error) {
	row, err := Engine.Update(iBlockContent, &BlockContent{Langcode: Langcode_})
	return row, err
}

// UpdateBlockContentViaId via map[string]interface{}{}
func UpdateBlockContentViaId(iId int64, iBlockContentMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(BlockContent)).Where("id = ?", iId).Update(iBlockContentMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateBlockContentViaRevisionId via map[string]interface{}{}
func UpdateBlockContentViaRevisionId(iRevisionId int, iBlockContentMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(BlockContent)).Where("revision_id = ?", iRevisionId).Update(iBlockContentMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateBlockContentViaType via map[string]interface{}{}
func UpdateBlockContentViaType(iType string, iBlockContentMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(BlockContent)).Where("type = ?", iType).Update(iBlockContentMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateBlockContentViaUuid via map[string]interface{}{}
func UpdateBlockContentViaUuid(iUuid string, iBlockContentMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(BlockContent)).Where("uuid = ?", iUuid).Update(iBlockContentMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateBlockContentViaLangcode via map[string]interface{}{}
func UpdateBlockContentViaLangcode(iLangcode string, iBlockContentMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(BlockContent)).Where("langcode = ?", iLangcode).Update(iBlockContentMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// DeleteBlockContent Delete BlockContent
func DeleteBlockContent(iId int64) (int64, error) {
	row, err := Engine.Id(iId).Delete(new(BlockContent))
	return row, err
}

// DeleteBlockContentViaId Delete BlockContent via Id
func DeleteBlockContentViaId(iId int64) (err error) {
	var has bool
	var _BlockContent = &BlockContent{Id: iId}
	if has, err = Engine.Get(_BlockContent); (has == true) && (err == nil) {
		if row, err := Engine.Where("id = ?", iId).Delete(new(BlockContent)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteBlockContentViaRevisionId Delete BlockContent via RevisionId
func DeleteBlockContentViaRevisionId(iRevisionId int) (err error) {
	var has bool
	var _BlockContent = &BlockContent{RevisionId: iRevisionId}
	if has, err = Engine.Get(_BlockContent); (has == true) && (err == nil) {
		if row, err := Engine.Where("revision_id = ?", iRevisionId).Delete(new(BlockContent)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteBlockContentViaType Delete BlockContent via Type
func DeleteBlockContentViaType(iType string) (err error) {
	var has bool
	var _BlockContent = &BlockContent{Type: iType}
	if has, err = Engine.Get(_BlockContent); (has == true) && (err == nil) {
		if row, err := Engine.Where("type = ?", iType).Delete(new(BlockContent)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteBlockContentViaUuid Delete BlockContent via Uuid
func DeleteBlockContentViaUuid(iUuid string) (err error) {
	var has bool
	var _BlockContent = &BlockContent{Uuid: iUuid}
	if has, err = Engine.Get(_BlockContent); (has == true) && (err == nil) {
		if row, err := Engine.Where("uuid = ?", iUuid).Delete(new(BlockContent)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteBlockContentViaLangcode Delete BlockContent via Langcode
func DeleteBlockContentViaLangcode(iLangcode string) (err error) {
	var has bool
	var _BlockContent = &BlockContent{Langcode: iLangcode}
	if has, err = Engine.Get(_BlockContent); (has == true) && (err == nil) {
		if row, err := Engine.Where("langcode = ?", iLangcode).Delete(new(BlockContent)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}
