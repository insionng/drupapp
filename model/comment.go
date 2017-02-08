package model

type Comment struct {
	Cid         int64  `xorm:"not null pk autoincr INT(10)"`
	CommentType string `xorm:"not null index VARCHAR(32)"`
	Uuid        string `xorm:"not null unique VARCHAR(128)"`
	Langcode    string `xorm:"not null VARCHAR(12)"`
}

// GetCommentsCount Comments' Count
func GetCommentsCount(offset int, limit int) (int64, error) {
	total, err := Engine.Limit(limit, offset).Count(&Comment{})
	return total, err
}

// GetCommentCountViaCid Get Comment via Cid
func GetCommentCountViaCid(iCid int64) int64 {
	n, _ := Engine.Where("cid = ?", iCid).Count(&Comment{Cid: iCid})
	return n
}

// GetCommentCountViaCommentType Get Comment via CommentType
func GetCommentCountViaCommentType(iCommentType string) int64 {
	n, _ := Engine.Where("comment_type = ?", iCommentType).Count(&Comment{CommentType: iCommentType})
	return n
}

// GetCommentCountViaUuid Get Comment via Uuid
func GetCommentCountViaUuid(iUuid string) int64 {
	n, _ := Engine.Where("uuid = ?", iUuid).Count(&Comment{Uuid: iUuid})
	return n
}

// GetCommentCountViaLangcode Get Comment via Langcode
func GetCommentCountViaLangcode(iLangcode string) int64 {
	n, _ := Engine.Where("langcode = ?", iLangcode).Count(&Comment{Langcode: iLangcode})
	return n
}

// GetCommentsByCidAndCommentTypeAndUuid Get Comments via CidAndCommentTypeAndUuid
func GetCommentsByCidAndCommentTypeAndUuid(offset int, limit int, Cid_ int64, CommentType_ string, Uuid_ string) (*[]*Comment, error) {
	var _Comment = new([]*Comment)
	err := Engine.Table("comment").Where("cid = ? and comment_type = ? and uuid = ?", Cid_, CommentType_, Uuid_).Limit(limit, offset).Find(_Comment)
	return _Comment, err
}

// GetCommentsByCidAndCommentTypeAndLangcode Get Comments via CidAndCommentTypeAndLangcode
func GetCommentsByCidAndCommentTypeAndLangcode(offset int, limit int, Cid_ int64, CommentType_ string, Langcode_ string) (*[]*Comment, error) {
	var _Comment = new([]*Comment)
	err := Engine.Table("comment").Where("cid = ? and comment_type = ? and langcode = ?", Cid_, CommentType_, Langcode_).Limit(limit, offset).Find(_Comment)
	return _Comment, err
}

// GetCommentsByCidAndUuidAndLangcode Get Comments via CidAndUuidAndLangcode
func GetCommentsByCidAndUuidAndLangcode(offset int, limit int, Cid_ int64, Uuid_ string, Langcode_ string) (*[]*Comment, error) {
	var _Comment = new([]*Comment)
	err := Engine.Table("comment").Where("cid = ? and uuid = ? and langcode = ?", Cid_, Uuid_, Langcode_).Limit(limit, offset).Find(_Comment)
	return _Comment, err
}

// GetCommentsByCommentTypeAndUuidAndLangcode Get Comments via CommentTypeAndUuidAndLangcode
func GetCommentsByCommentTypeAndUuidAndLangcode(offset int, limit int, CommentType_ string, Uuid_ string, Langcode_ string) (*[]*Comment, error) {
	var _Comment = new([]*Comment)
	err := Engine.Table("comment").Where("comment_type = ? and uuid = ? and langcode = ?", CommentType_, Uuid_, Langcode_).Limit(limit, offset).Find(_Comment)
	return _Comment, err
}

// GetCommentsByCidAndCommentType Get Comments via CidAndCommentType
func GetCommentsByCidAndCommentType(offset int, limit int, Cid_ int64, CommentType_ string) (*[]*Comment, error) {
	var _Comment = new([]*Comment)
	err := Engine.Table("comment").Where("cid = ? and comment_type = ?", Cid_, CommentType_).Limit(limit, offset).Find(_Comment)
	return _Comment, err
}

// GetCommentsByCidAndUuid Get Comments via CidAndUuid
func GetCommentsByCidAndUuid(offset int, limit int, Cid_ int64, Uuid_ string) (*[]*Comment, error) {
	var _Comment = new([]*Comment)
	err := Engine.Table("comment").Where("cid = ? and uuid = ?", Cid_, Uuid_).Limit(limit, offset).Find(_Comment)
	return _Comment, err
}

// GetCommentsByCidAndLangcode Get Comments via CidAndLangcode
func GetCommentsByCidAndLangcode(offset int, limit int, Cid_ int64, Langcode_ string) (*[]*Comment, error) {
	var _Comment = new([]*Comment)
	err := Engine.Table("comment").Where("cid = ? and langcode = ?", Cid_, Langcode_).Limit(limit, offset).Find(_Comment)
	return _Comment, err
}

// GetCommentsByCommentTypeAndUuid Get Comments via CommentTypeAndUuid
func GetCommentsByCommentTypeAndUuid(offset int, limit int, CommentType_ string, Uuid_ string) (*[]*Comment, error) {
	var _Comment = new([]*Comment)
	err := Engine.Table("comment").Where("comment_type = ? and uuid = ?", CommentType_, Uuid_).Limit(limit, offset).Find(_Comment)
	return _Comment, err
}

// GetCommentsByCommentTypeAndLangcode Get Comments via CommentTypeAndLangcode
func GetCommentsByCommentTypeAndLangcode(offset int, limit int, CommentType_ string, Langcode_ string) (*[]*Comment, error) {
	var _Comment = new([]*Comment)
	err := Engine.Table("comment").Where("comment_type = ? and langcode = ?", CommentType_, Langcode_).Limit(limit, offset).Find(_Comment)
	return _Comment, err
}

// GetCommentsByUuidAndLangcode Get Comments via UuidAndLangcode
func GetCommentsByUuidAndLangcode(offset int, limit int, Uuid_ string, Langcode_ string) (*[]*Comment, error) {
	var _Comment = new([]*Comment)
	err := Engine.Table("comment").Where("uuid = ? and langcode = ?", Uuid_, Langcode_).Limit(limit, offset).Find(_Comment)
	return _Comment, err
}

// GetComments Get Comments via field
func GetComments(offset int, limit int, field string) (*[]*Comment, error) {
	var _Comment = new([]*Comment)
	err := Engine.Table("comment").Limit(limit, offset).Desc(field).Find(_Comment)
	return _Comment, err
}

// GetCommentsViaCid Get Comments via Cid
func GetCommentsViaCid(offset int, limit int, Cid_ int64, field string) (*[]*Comment, error) {
	var _Comment = new([]*Comment)
	err := Engine.Table("comment").Where("cid = ?", Cid_).Limit(limit, offset).Desc(field).Find(_Comment)
	return _Comment, err
}

// GetCommentsViaCommentType Get Comments via CommentType
func GetCommentsViaCommentType(offset int, limit int, CommentType_ string, field string) (*[]*Comment, error) {
	var _Comment = new([]*Comment)
	err := Engine.Table("comment").Where("comment_type = ?", CommentType_).Limit(limit, offset).Desc(field).Find(_Comment)
	return _Comment, err
}

// GetCommentsViaUuid Get Comments via Uuid
func GetCommentsViaUuid(offset int, limit int, Uuid_ string, field string) (*[]*Comment, error) {
	var _Comment = new([]*Comment)
	err := Engine.Table("comment").Where("uuid = ?", Uuid_).Limit(limit, offset).Desc(field).Find(_Comment)
	return _Comment, err
}

// GetCommentsViaLangcode Get Comments via Langcode
func GetCommentsViaLangcode(offset int, limit int, Langcode_ string, field string) (*[]*Comment, error) {
	var _Comment = new([]*Comment)
	err := Engine.Table("comment").Where("langcode = ?", Langcode_).Limit(limit, offset).Desc(field).Find(_Comment)
	return _Comment, err
}

// HasCommentViaCid Has Comment via Cid
func HasCommentViaCid(iCid int64) bool {
	if has, err := Engine.Where("cid = ?", iCid).Get(new(Comment)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasCommentViaCommentType Has Comment via CommentType
func HasCommentViaCommentType(iCommentType string) bool {
	if has, err := Engine.Where("comment_type = ?", iCommentType).Get(new(Comment)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasCommentViaUuid Has Comment via Uuid
func HasCommentViaUuid(iUuid string) bool {
	if has, err := Engine.Where("uuid = ?", iUuid).Get(new(Comment)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasCommentViaLangcode Has Comment via Langcode
func HasCommentViaLangcode(iLangcode string) bool {
	if has, err := Engine.Where("langcode = ?", iLangcode).Get(new(Comment)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// GetCommentViaCid Get Comment via Cid
func GetCommentViaCid(iCid int64) (*Comment, error) {
	var _Comment = &Comment{Cid: iCid}
	has, err := Engine.Get(_Comment)
	if has {
		return _Comment, err
	} else {
		return nil, err
	}
}

// GetCommentViaCommentType Get Comment via CommentType
func GetCommentViaCommentType(iCommentType string) (*Comment, error) {
	var _Comment = &Comment{CommentType: iCommentType}
	has, err := Engine.Get(_Comment)
	if has {
		return _Comment, err
	} else {
		return nil, err
	}
}

// GetCommentViaUuid Get Comment via Uuid
func GetCommentViaUuid(iUuid string) (*Comment, error) {
	var _Comment = &Comment{Uuid: iUuid}
	has, err := Engine.Get(_Comment)
	if has {
		return _Comment, err
	} else {
		return nil, err
	}
}

// GetCommentViaLangcode Get Comment via Langcode
func GetCommentViaLangcode(iLangcode string) (*Comment, error) {
	var _Comment = &Comment{Langcode: iLangcode}
	has, err := Engine.Get(_Comment)
	if has {
		return _Comment, err
	} else {
		return nil, err
	}
}

// SetCommentViaCid Set Comment via Cid
func SetCommentViaCid(iCid int64, comment *Comment) (int64, error) {
	comment.Cid = iCid
	return Engine.Insert(comment)
}

// SetCommentViaCommentType Set Comment via CommentType
func SetCommentViaCommentType(iCommentType string, comment *Comment) (int64, error) {
	comment.CommentType = iCommentType
	return Engine.Insert(comment)
}

// SetCommentViaUuid Set Comment via Uuid
func SetCommentViaUuid(iUuid string, comment *Comment) (int64, error) {
	comment.Uuid = iUuid
	return Engine.Insert(comment)
}

// SetCommentViaLangcode Set Comment via Langcode
func SetCommentViaLangcode(iLangcode string, comment *Comment) (int64, error) {
	comment.Langcode = iLangcode
	return Engine.Insert(comment)
}

// AddComment Add Comment via all columns
func AddComment(iCid int64, iCommentType string, iUuid string, iLangcode string) error {
	_Comment := &Comment{Cid: iCid, CommentType: iCommentType, Uuid: iUuid, Langcode: iLangcode}
	if _, err := Engine.Insert(_Comment); err != nil {
		return err
	} else {
		return nil
	}
}

// PostComment Post Comment via iComment
func PostComment(iComment *Comment) (int64, error) {
	_, err := Engine.Insert(iComment)
	return iComment.Cid, err
}

// PutComment Put Comment
func PutComment(iComment *Comment) (int64, error) {
	_, err := Engine.Id(iComment.Cid).Update(iComment)
	return iComment.Cid, err
}

// PutCommentViaCid Put Comment via Cid
func PutCommentViaCid(Cid_ int64, iComment *Comment) (int64, error) {
	row, err := Engine.Update(iComment, &Comment{Cid: Cid_})
	return row, err
}

// PutCommentViaCommentType Put Comment via CommentType
func PutCommentViaCommentType(CommentType_ string, iComment *Comment) (int64, error) {
	row, err := Engine.Update(iComment, &Comment{CommentType: CommentType_})
	return row, err
}

// PutCommentViaUuid Put Comment via Uuid
func PutCommentViaUuid(Uuid_ string, iComment *Comment) (int64, error) {
	row, err := Engine.Update(iComment, &Comment{Uuid: Uuid_})
	return row, err
}

// PutCommentViaLangcode Put Comment via Langcode
func PutCommentViaLangcode(Langcode_ string, iComment *Comment) (int64, error) {
	row, err := Engine.Update(iComment, &Comment{Langcode: Langcode_})
	return row, err
}

// UpdateCommentViaCid via map[string]interface{}{}
func UpdateCommentViaCid(iCid int64, iCommentMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Comment)).Where("cid = ?", iCid).Update(iCommentMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateCommentViaCommentType via map[string]interface{}{}
func UpdateCommentViaCommentType(iCommentType string, iCommentMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Comment)).Where("comment_type = ?", iCommentType).Update(iCommentMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateCommentViaUuid via map[string]interface{}{}
func UpdateCommentViaUuid(iUuid string, iCommentMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Comment)).Where("uuid = ?", iUuid).Update(iCommentMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateCommentViaLangcode via map[string]interface{}{}
func UpdateCommentViaLangcode(iLangcode string, iCommentMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Comment)).Where("langcode = ?", iLangcode).Update(iCommentMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// DeleteComment Delete Comment
func DeleteComment(iCid int64) (int64, error) {
	row, err := Engine.Id(iCid).Delete(new(Comment))
	return row, err
}

// DeleteCommentViaCid Delete Comment via Cid
func DeleteCommentViaCid(iCid int64) (err error) {
	var has bool
	var _Comment = &Comment{Cid: iCid}
	if has, err = Engine.Get(_Comment); (has == true) && (err == nil) {
		if row, err := Engine.Where("cid = ?", iCid).Delete(new(Comment)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteCommentViaCommentType Delete Comment via CommentType
func DeleteCommentViaCommentType(iCommentType string) (err error) {
	var has bool
	var _Comment = &Comment{CommentType: iCommentType}
	if has, err = Engine.Get(_Comment); (has == true) && (err == nil) {
		if row, err := Engine.Where("comment_type = ?", iCommentType).Delete(new(Comment)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteCommentViaUuid Delete Comment via Uuid
func DeleteCommentViaUuid(iUuid string) (err error) {
	var has bool
	var _Comment = &Comment{Uuid: iUuid}
	if has, err = Engine.Get(_Comment); (has == true) && (err == nil) {
		if row, err := Engine.Where("uuid = ?", iUuid).Delete(new(Comment)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteCommentViaLangcode Delete Comment via Langcode
func DeleteCommentViaLangcode(iLangcode string) (err error) {
	var has bool
	var _Comment = &Comment{Langcode: iLangcode}
	if has, err = Engine.Get(_Comment); (has == true) && (err == nil) {
		if row, err := Engine.Where("langcode = ?", iLangcode).Delete(new(Comment)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}
