package model

type TaxonomyIndex struct {
	Nid     int64 `xorm:"not null pk default 0 INT(10)"`
	Tid     int   `xorm:"not null pk default 0 index(term_node) INT(10)"`
	Status  int   `xorm:"not null default 1 index(term_node) INT(11)"`
	Sticky  int   `xorm:"default 0 index(term_node) TINYINT(4)"`
	Created int   `xorm:"not null default 0 index(term_node) INT(11)"`
}

// GetTaxonomyIndexsCount TaxonomyIndexs' Count
func GetTaxonomyIndexsCount(offset int, limit int) (int64, error) {
	total, err := Engine.Limit(limit, offset).Count(&TaxonomyIndex{})
	return total, err
}

// GetTaxonomyIndexCountViaNid Get TaxonomyIndex via Nid
func GetTaxonomyIndexCountViaNid(iNid int64) int64 {
	n, _ := Engine.Where("nid = ?", iNid).Count(&TaxonomyIndex{Nid: iNid})
	return n
}

// GetTaxonomyIndexCountViaTid Get TaxonomyIndex via Tid
func GetTaxonomyIndexCountViaTid(iTid int) int64 {
	n, _ := Engine.Where("tid = ?", iTid).Count(&TaxonomyIndex{Tid: iTid})
	return n
}

// GetTaxonomyIndexCountViaStatus Get TaxonomyIndex via Status
func GetTaxonomyIndexCountViaStatus(iStatus int) int64 {
	n, _ := Engine.Where("status = ?", iStatus).Count(&TaxonomyIndex{Status: iStatus})
	return n
}

// GetTaxonomyIndexCountViaSticky Get TaxonomyIndex via Sticky
func GetTaxonomyIndexCountViaSticky(iSticky int) int64 {
	n, _ := Engine.Where("sticky = ?", iSticky).Count(&TaxonomyIndex{Sticky: iSticky})
	return n
}

// GetTaxonomyIndexCountViaCreated Get TaxonomyIndex via Created
func GetTaxonomyIndexCountViaCreated(iCreated int) int64 {
	n, _ := Engine.Where("created = ?", iCreated).Count(&TaxonomyIndex{Created: iCreated})
	return n
}

// GetTaxonomyIndexsByNidAndTidAndStatus Get TaxonomyIndexs via NidAndTidAndStatus
func GetTaxonomyIndexsByNidAndTidAndStatus(offset int, limit int, Nid_ int64, Tid_ int, Status_ int) (*[]*TaxonomyIndex, error) {
	var _TaxonomyIndex = new([]*TaxonomyIndex)
	err := Engine.Table("taxonomy_index").Where("nid = ? and tid = ? and status = ?", Nid_, Tid_, Status_).Limit(limit, offset).Find(_TaxonomyIndex)
	return _TaxonomyIndex, err
}

// GetTaxonomyIndexsByNidAndTidAndSticky Get TaxonomyIndexs via NidAndTidAndSticky
func GetTaxonomyIndexsByNidAndTidAndSticky(offset int, limit int, Nid_ int64, Tid_ int, Sticky_ int) (*[]*TaxonomyIndex, error) {
	var _TaxonomyIndex = new([]*TaxonomyIndex)
	err := Engine.Table("taxonomy_index").Where("nid = ? and tid = ? and sticky = ?", Nid_, Tid_, Sticky_).Limit(limit, offset).Find(_TaxonomyIndex)
	return _TaxonomyIndex, err
}

// GetTaxonomyIndexsByNidAndTidAndCreated Get TaxonomyIndexs via NidAndTidAndCreated
func GetTaxonomyIndexsByNidAndTidAndCreated(offset int, limit int, Nid_ int64, Tid_ int, Created_ int) (*[]*TaxonomyIndex, error) {
	var _TaxonomyIndex = new([]*TaxonomyIndex)
	err := Engine.Table("taxonomy_index").Where("nid = ? and tid = ? and created = ?", Nid_, Tid_, Created_).Limit(limit, offset).Find(_TaxonomyIndex)
	return _TaxonomyIndex, err
}

// GetTaxonomyIndexsByNidAndStatusAndSticky Get TaxonomyIndexs via NidAndStatusAndSticky
func GetTaxonomyIndexsByNidAndStatusAndSticky(offset int, limit int, Nid_ int64, Status_ int, Sticky_ int) (*[]*TaxonomyIndex, error) {
	var _TaxonomyIndex = new([]*TaxonomyIndex)
	err := Engine.Table("taxonomy_index").Where("nid = ? and status = ? and sticky = ?", Nid_, Status_, Sticky_).Limit(limit, offset).Find(_TaxonomyIndex)
	return _TaxonomyIndex, err
}

// GetTaxonomyIndexsByNidAndStatusAndCreated Get TaxonomyIndexs via NidAndStatusAndCreated
func GetTaxonomyIndexsByNidAndStatusAndCreated(offset int, limit int, Nid_ int64, Status_ int, Created_ int) (*[]*TaxonomyIndex, error) {
	var _TaxonomyIndex = new([]*TaxonomyIndex)
	err := Engine.Table("taxonomy_index").Where("nid = ? and status = ? and created = ?", Nid_, Status_, Created_).Limit(limit, offset).Find(_TaxonomyIndex)
	return _TaxonomyIndex, err
}

// GetTaxonomyIndexsByNidAndStickyAndCreated Get TaxonomyIndexs via NidAndStickyAndCreated
func GetTaxonomyIndexsByNidAndStickyAndCreated(offset int, limit int, Nid_ int64, Sticky_ int, Created_ int) (*[]*TaxonomyIndex, error) {
	var _TaxonomyIndex = new([]*TaxonomyIndex)
	err := Engine.Table("taxonomy_index").Where("nid = ? and sticky = ? and created = ?", Nid_, Sticky_, Created_).Limit(limit, offset).Find(_TaxonomyIndex)
	return _TaxonomyIndex, err
}

// GetTaxonomyIndexsByTidAndStatusAndSticky Get TaxonomyIndexs via TidAndStatusAndSticky
func GetTaxonomyIndexsByTidAndStatusAndSticky(offset int, limit int, Tid_ int, Status_ int, Sticky_ int) (*[]*TaxonomyIndex, error) {
	var _TaxonomyIndex = new([]*TaxonomyIndex)
	err := Engine.Table("taxonomy_index").Where("tid = ? and status = ? and sticky = ?", Tid_, Status_, Sticky_).Limit(limit, offset).Find(_TaxonomyIndex)
	return _TaxonomyIndex, err
}

// GetTaxonomyIndexsByTidAndStatusAndCreated Get TaxonomyIndexs via TidAndStatusAndCreated
func GetTaxonomyIndexsByTidAndStatusAndCreated(offset int, limit int, Tid_ int, Status_ int, Created_ int) (*[]*TaxonomyIndex, error) {
	var _TaxonomyIndex = new([]*TaxonomyIndex)
	err := Engine.Table("taxonomy_index").Where("tid = ? and status = ? and created = ?", Tid_, Status_, Created_).Limit(limit, offset).Find(_TaxonomyIndex)
	return _TaxonomyIndex, err
}

// GetTaxonomyIndexsByTidAndStickyAndCreated Get TaxonomyIndexs via TidAndStickyAndCreated
func GetTaxonomyIndexsByTidAndStickyAndCreated(offset int, limit int, Tid_ int, Sticky_ int, Created_ int) (*[]*TaxonomyIndex, error) {
	var _TaxonomyIndex = new([]*TaxonomyIndex)
	err := Engine.Table("taxonomy_index").Where("tid = ? and sticky = ? and created = ?", Tid_, Sticky_, Created_).Limit(limit, offset).Find(_TaxonomyIndex)
	return _TaxonomyIndex, err
}

// GetTaxonomyIndexsByStatusAndStickyAndCreated Get TaxonomyIndexs via StatusAndStickyAndCreated
func GetTaxonomyIndexsByStatusAndStickyAndCreated(offset int, limit int, Status_ int, Sticky_ int, Created_ int) (*[]*TaxonomyIndex, error) {
	var _TaxonomyIndex = new([]*TaxonomyIndex)
	err := Engine.Table("taxonomy_index").Where("status = ? and sticky = ? and created = ?", Status_, Sticky_, Created_).Limit(limit, offset).Find(_TaxonomyIndex)
	return _TaxonomyIndex, err
}

// GetTaxonomyIndexsByNidAndTid Get TaxonomyIndexs via NidAndTid
func GetTaxonomyIndexsByNidAndTid(offset int, limit int, Nid_ int64, Tid_ int) (*[]*TaxonomyIndex, error) {
	var _TaxonomyIndex = new([]*TaxonomyIndex)
	err := Engine.Table("taxonomy_index").Where("nid = ? and tid = ?", Nid_, Tid_).Limit(limit, offset).Find(_TaxonomyIndex)
	return _TaxonomyIndex, err
}

// GetTaxonomyIndexsByNidAndStatus Get TaxonomyIndexs via NidAndStatus
func GetTaxonomyIndexsByNidAndStatus(offset int, limit int, Nid_ int64, Status_ int) (*[]*TaxonomyIndex, error) {
	var _TaxonomyIndex = new([]*TaxonomyIndex)
	err := Engine.Table("taxonomy_index").Where("nid = ? and status = ?", Nid_, Status_).Limit(limit, offset).Find(_TaxonomyIndex)
	return _TaxonomyIndex, err
}

// GetTaxonomyIndexsByNidAndSticky Get TaxonomyIndexs via NidAndSticky
func GetTaxonomyIndexsByNidAndSticky(offset int, limit int, Nid_ int64, Sticky_ int) (*[]*TaxonomyIndex, error) {
	var _TaxonomyIndex = new([]*TaxonomyIndex)
	err := Engine.Table("taxonomy_index").Where("nid = ? and sticky = ?", Nid_, Sticky_).Limit(limit, offset).Find(_TaxonomyIndex)
	return _TaxonomyIndex, err
}

// GetTaxonomyIndexsByNidAndCreated Get TaxonomyIndexs via NidAndCreated
func GetTaxonomyIndexsByNidAndCreated(offset int, limit int, Nid_ int64, Created_ int) (*[]*TaxonomyIndex, error) {
	var _TaxonomyIndex = new([]*TaxonomyIndex)
	err := Engine.Table("taxonomy_index").Where("nid = ? and created = ?", Nid_, Created_).Limit(limit, offset).Find(_TaxonomyIndex)
	return _TaxonomyIndex, err
}

// GetTaxonomyIndexsByTidAndStatus Get TaxonomyIndexs via TidAndStatus
func GetTaxonomyIndexsByTidAndStatus(offset int, limit int, Tid_ int, Status_ int) (*[]*TaxonomyIndex, error) {
	var _TaxonomyIndex = new([]*TaxonomyIndex)
	err := Engine.Table("taxonomy_index").Where("tid = ? and status = ?", Tid_, Status_).Limit(limit, offset).Find(_TaxonomyIndex)
	return _TaxonomyIndex, err
}

// GetTaxonomyIndexsByTidAndSticky Get TaxonomyIndexs via TidAndSticky
func GetTaxonomyIndexsByTidAndSticky(offset int, limit int, Tid_ int, Sticky_ int) (*[]*TaxonomyIndex, error) {
	var _TaxonomyIndex = new([]*TaxonomyIndex)
	err := Engine.Table("taxonomy_index").Where("tid = ? and sticky = ?", Tid_, Sticky_).Limit(limit, offset).Find(_TaxonomyIndex)
	return _TaxonomyIndex, err
}

// GetTaxonomyIndexsByTidAndCreated Get TaxonomyIndexs via TidAndCreated
func GetTaxonomyIndexsByTidAndCreated(offset int, limit int, Tid_ int, Created_ int) (*[]*TaxonomyIndex, error) {
	var _TaxonomyIndex = new([]*TaxonomyIndex)
	err := Engine.Table("taxonomy_index").Where("tid = ? and created = ?", Tid_, Created_).Limit(limit, offset).Find(_TaxonomyIndex)
	return _TaxonomyIndex, err
}

// GetTaxonomyIndexsByStatusAndSticky Get TaxonomyIndexs via StatusAndSticky
func GetTaxonomyIndexsByStatusAndSticky(offset int, limit int, Status_ int, Sticky_ int) (*[]*TaxonomyIndex, error) {
	var _TaxonomyIndex = new([]*TaxonomyIndex)
	err := Engine.Table("taxonomy_index").Where("status = ? and sticky = ?", Status_, Sticky_).Limit(limit, offset).Find(_TaxonomyIndex)
	return _TaxonomyIndex, err
}

// GetTaxonomyIndexsByStatusAndCreated Get TaxonomyIndexs via StatusAndCreated
func GetTaxonomyIndexsByStatusAndCreated(offset int, limit int, Status_ int, Created_ int) (*[]*TaxonomyIndex, error) {
	var _TaxonomyIndex = new([]*TaxonomyIndex)
	err := Engine.Table("taxonomy_index").Where("status = ? and created = ?", Status_, Created_).Limit(limit, offset).Find(_TaxonomyIndex)
	return _TaxonomyIndex, err
}

// GetTaxonomyIndexsByStickyAndCreated Get TaxonomyIndexs via StickyAndCreated
func GetTaxonomyIndexsByStickyAndCreated(offset int, limit int, Sticky_ int, Created_ int) (*[]*TaxonomyIndex, error) {
	var _TaxonomyIndex = new([]*TaxonomyIndex)
	err := Engine.Table("taxonomy_index").Where("sticky = ? and created = ?", Sticky_, Created_).Limit(limit, offset).Find(_TaxonomyIndex)
	return _TaxonomyIndex, err
}

// GetTaxonomyIndexs Get TaxonomyIndexs via field
func GetTaxonomyIndexs(offset int, limit int, field string) (*[]*TaxonomyIndex, error) {
	var _TaxonomyIndex = new([]*TaxonomyIndex)
	err := Engine.Table("taxonomy_index").Limit(limit, offset).Desc(field).Find(_TaxonomyIndex)
	return _TaxonomyIndex, err
}

// GetTaxonomyIndexsViaNid Get TaxonomyIndexs via Nid
func GetTaxonomyIndexsViaNid(offset int, limit int, Nid_ int64, field string) (*[]*TaxonomyIndex, error) {
	var _TaxonomyIndex = new([]*TaxonomyIndex)
	err := Engine.Table("taxonomy_index").Where("nid = ?", Nid_).Limit(limit, offset).Desc(field).Find(_TaxonomyIndex)
	return _TaxonomyIndex, err
}

// GetTaxonomyIndexsViaTid Get TaxonomyIndexs via Tid
func GetTaxonomyIndexsViaTid(offset int, limit int, Tid_ int, field string) (*[]*TaxonomyIndex, error) {
	var _TaxonomyIndex = new([]*TaxonomyIndex)
	err := Engine.Table("taxonomy_index").Where("tid = ?", Tid_).Limit(limit, offset).Desc(field).Find(_TaxonomyIndex)
	return _TaxonomyIndex, err
}

// GetTaxonomyIndexsViaStatus Get TaxonomyIndexs via Status
func GetTaxonomyIndexsViaStatus(offset int, limit int, Status_ int, field string) (*[]*TaxonomyIndex, error) {
	var _TaxonomyIndex = new([]*TaxonomyIndex)
	err := Engine.Table("taxonomy_index").Where("status = ?", Status_).Limit(limit, offset).Desc(field).Find(_TaxonomyIndex)
	return _TaxonomyIndex, err
}

// GetTaxonomyIndexsViaSticky Get TaxonomyIndexs via Sticky
func GetTaxonomyIndexsViaSticky(offset int, limit int, Sticky_ int, field string) (*[]*TaxonomyIndex, error) {
	var _TaxonomyIndex = new([]*TaxonomyIndex)
	err := Engine.Table("taxonomy_index").Where("sticky = ?", Sticky_).Limit(limit, offset).Desc(field).Find(_TaxonomyIndex)
	return _TaxonomyIndex, err
}

// GetTaxonomyIndexsViaCreated Get TaxonomyIndexs via Created
func GetTaxonomyIndexsViaCreated(offset int, limit int, Created_ int, field string) (*[]*TaxonomyIndex, error) {
	var _TaxonomyIndex = new([]*TaxonomyIndex)
	err := Engine.Table("taxonomy_index").Where("created = ?", Created_).Limit(limit, offset).Desc(field).Find(_TaxonomyIndex)
	return _TaxonomyIndex, err
}

// HasTaxonomyIndexViaNid Has TaxonomyIndex via Nid
func HasTaxonomyIndexViaNid(iNid int64) bool {
	if has, err := Engine.Where("nid = ?", iNid).Get(new(TaxonomyIndex)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasTaxonomyIndexViaTid Has TaxonomyIndex via Tid
func HasTaxonomyIndexViaTid(iTid int) bool {
	if has, err := Engine.Where("tid = ?", iTid).Get(new(TaxonomyIndex)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasTaxonomyIndexViaStatus Has TaxonomyIndex via Status
func HasTaxonomyIndexViaStatus(iStatus int) bool {
	if has, err := Engine.Where("status = ?", iStatus).Get(new(TaxonomyIndex)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasTaxonomyIndexViaSticky Has TaxonomyIndex via Sticky
func HasTaxonomyIndexViaSticky(iSticky int) bool {
	if has, err := Engine.Where("sticky = ?", iSticky).Get(new(TaxonomyIndex)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasTaxonomyIndexViaCreated Has TaxonomyIndex via Created
func HasTaxonomyIndexViaCreated(iCreated int) bool {
	if has, err := Engine.Where("created = ?", iCreated).Get(new(TaxonomyIndex)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// GetTaxonomyIndexViaNid Get TaxonomyIndex via Nid
func GetTaxonomyIndexViaNid(iNid int64) (*TaxonomyIndex, error) {
	var _TaxonomyIndex = &TaxonomyIndex{Nid: iNid}
	has, err := Engine.Get(_TaxonomyIndex)
	if has {
		return _TaxonomyIndex, err
	} else {
		return nil, err
	}
}

// GetTaxonomyIndexViaTid Get TaxonomyIndex via Tid
func GetTaxonomyIndexViaTid(iTid int) (*TaxonomyIndex, error) {
	var _TaxonomyIndex = &TaxonomyIndex{Tid: iTid}
	has, err := Engine.Get(_TaxonomyIndex)
	if has {
		return _TaxonomyIndex, err
	} else {
		return nil, err
	}
}

// GetTaxonomyIndexViaStatus Get TaxonomyIndex via Status
func GetTaxonomyIndexViaStatus(iStatus int) (*TaxonomyIndex, error) {
	var _TaxonomyIndex = &TaxonomyIndex{Status: iStatus}
	has, err := Engine.Get(_TaxonomyIndex)
	if has {
		return _TaxonomyIndex, err
	} else {
		return nil, err
	}
}

// GetTaxonomyIndexViaSticky Get TaxonomyIndex via Sticky
func GetTaxonomyIndexViaSticky(iSticky int) (*TaxonomyIndex, error) {
	var _TaxonomyIndex = &TaxonomyIndex{Sticky: iSticky}
	has, err := Engine.Get(_TaxonomyIndex)
	if has {
		return _TaxonomyIndex, err
	} else {
		return nil, err
	}
}

// GetTaxonomyIndexViaCreated Get TaxonomyIndex via Created
func GetTaxonomyIndexViaCreated(iCreated int) (*TaxonomyIndex, error) {
	var _TaxonomyIndex = &TaxonomyIndex{Created: iCreated}
	has, err := Engine.Get(_TaxonomyIndex)
	if has {
		return _TaxonomyIndex, err
	} else {
		return nil, err
	}
}

// SetTaxonomyIndexViaNid Set TaxonomyIndex via Nid
func SetTaxonomyIndexViaNid(iNid int64, taxonomy_index *TaxonomyIndex) (int64, error) {
	taxonomy_index.Nid = iNid
	return Engine.Insert(taxonomy_index)
}

// SetTaxonomyIndexViaTid Set TaxonomyIndex via Tid
func SetTaxonomyIndexViaTid(iTid int, taxonomy_index *TaxonomyIndex) (int64, error) {
	taxonomy_index.Tid = iTid
	return Engine.Insert(taxonomy_index)
}

// SetTaxonomyIndexViaStatus Set TaxonomyIndex via Status
func SetTaxonomyIndexViaStatus(iStatus int, taxonomy_index *TaxonomyIndex) (int64, error) {
	taxonomy_index.Status = iStatus
	return Engine.Insert(taxonomy_index)
}

// SetTaxonomyIndexViaSticky Set TaxonomyIndex via Sticky
func SetTaxonomyIndexViaSticky(iSticky int, taxonomy_index *TaxonomyIndex) (int64, error) {
	taxonomy_index.Sticky = iSticky
	return Engine.Insert(taxonomy_index)
}

// SetTaxonomyIndexViaCreated Set TaxonomyIndex via Created
func SetTaxonomyIndexViaCreated(iCreated int, taxonomy_index *TaxonomyIndex) (int64, error) {
	taxonomy_index.Created = iCreated
	return Engine.Insert(taxonomy_index)
}

// AddTaxonomyIndex Add TaxonomyIndex via all columns
func AddTaxonomyIndex(iNid int64, iTid int, iStatus int, iSticky int, iCreated int) error {
	_TaxonomyIndex := &TaxonomyIndex{Nid: iNid, Tid: iTid, Status: iStatus, Sticky: iSticky, Created: iCreated}
	if _, err := Engine.Insert(_TaxonomyIndex); err != nil {
		return err
	} else {
		return nil
	}
}

// PostTaxonomyIndex Post TaxonomyIndex via iTaxonomyIndex
func PostTaxonomyIndex(iTaxonomyIndex *TaxonomyIndex) (int64, error) {
	_, err := Engine.Insert(iTaxonomyIndex)
	return iTaxonomyIndex.Nid, err
}

// PutTaxonomyIndex Put TaxonomyIndex
func PutTaxonomyIndex(iTaxonomyIndex *TaxonomyIndex) (int64, error) {
	_, err := Engine.Id(iTaxonomyIndex.Nid).Update(iTaxonomyIndex)
	return iTaxonomyIndex.Nid, err
}

// PutTaxonomyIndexViaNid Put TaxonomyIndex via Nid
func PutTaxonomyIndexViaNid(Nid_ int64, iTaxonomyIndex *TaxonomyIndex) (int64, error) {
	row, err := Engine.Update(iTaxonomyIndex, &TaxonomyIndex{Nid: Nid_})
	return row, err
}

// PutTaxonomyIndexViaTid Put TaxonomyIndex via Tid
func PutTaxonomyIndexViaTid(Tid_ int, iTaxonomyIndex *TaxonomyIndex) (int64, error) {
	row, err := Engine.Update(iTaxonomyIndex, &TaxonomyIndex{Tid: Tid_})
	return row, err
}

// PutTaxonomyIndexViaStatus Put TaxonomyIndex via Status
func PutTaxonomyIndexViaStatus(Status_ int, iTaxonomyIndex *TaxonomyIndex) (int64, error) {
	row, err := Engine.Update(iTaxonomyIndex, &TaxonomyIndex{Status: Status_})
	return row, err
}

// PutTaxonomyIndexViaSticky Put TaxonomyIndex via Sticky
func PutTaxonomyIndexViaSticky(Sticky_ int, iTaxonomyIndex *TaxonomyIndex) (int64, error) {
	row, err := Engine.Update(iTaxonomyIndex, &TaxonomyIndex{Sticky: Sticky_})
	return row, err
}

// PutTaxonomyIndexViaCreated Put TaxonomyIndex via Created
func PutTaxonomyIndexViaCreated(Created_ int, iTaxonomyIndex *TaxonomyIndex) (int64, error) {
	row, err := Engine.Update(iTaxonomyIndex, &TaxonomyIndex{Created: Created_})
	return row, err
}

// UpdateTaxonomyIndexViaNid via map[string]interface{}{}
func UpdateTaxonomyIndexViaNid(iNid int64, iTaxonomyIndexMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(TaxonomyIndex)).Where("nid = ?", iNid).Update(iTaxonomyIndexMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateTaxonomyIndexViaTid via map[string]interface{}{}
func UpdateTaxonomyIndexViaTid(iTid int, iTaxonomyIndexMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(TaxonomyIndex)).Where("tid = ?", iTid).Update(iTaxonomyIndexMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateTaxonomyIndexViaStatus via map[string]interface{}{}
func UpdateTaxonomyIndexViaStatus(iStatus int, iTaxonomyIndexMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(TaxonomyIndex)).Where("status = ?", iStatus).Update(iTaxonomyIndexMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateTaxonomyIndexViaSticky via map[string]interface{}{}
func UpdateTaxonomyIndexViaSticky(iSticky int, iTaxonomyIndexMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(TaxonomyIndex)).Where("sticky = ?", iSticky).Update(iTaxonomyIndexMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateTaxonomyIndexViaCreated via map[string]interface{}{}
func UpdateTaxonomyIndexViaCreated(iCreated int, iTaxonomyIndexMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(TaxonomyIndex)).Where("created = ?", iCreated).Update(iTaxonomyIndexMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// DeleteTaxonomyIndex Delete TaxonomyIndex
func DeleteTaxonomyIndex(iNid int64) (int64, error) {
	row, err := Engine.Id(iNid).Delete(new(TaxonomyIndex))
	return row, err
}

// DeleteTaxonomyIndexViaNid Delete TaxonomyIndex via Nid
func DeleteTaxonomyIndexViaNid(iNid int64) (err error) {
	var has bool
	var _TaxonomyIndex = &TaxonomyIndex{Nid: iNid}
	if has, err = Engine.Get(_TaxonomyIndex); (has == true) && (err == nil) {
		if row, err := Engine.Where("nid = ?", iNid).Delete(new(TaxonomyIndex)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteTaxonomyIndexViaTid Delete TaxonomyIndex via Tid
func DeleteTaxonomyIndexViaTid(iTid int) (err error) {
	var has bool
	var _TaxonomyIndex = &TaxonomyIndex{Tid: iTid}
	if has, err = Engine.Get(_TaxonomyIndex); (has == true) && (err == nil) {
		if row, err := Engine.Where("tid = ?", iTid).Delete(new(TaxonomyIndex)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteTaxonomyIndexViaStatus Delete TaxonomyIndex via Status
func DeleteTaxonomyIndexViaStatus(iStatus int) (err error) {
	var has bool
	var _TaxonomyIndex = &TaxonomyIndex{Status: iStatus}
	if has, err = Engine.Get(_TaxonomyIndex); (has == true) && (err == nil) {
		if row, err := Engine.Where("status = ?", iStatus).Delete(new(TaxonomyIndex)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteTaxonomyIndexViaSticky Delete TaxonomyIndex via Sticky
func DeleteTaxonomyIndexViaSticky(iSticky int) (err error) {
	var has bool
	var _TaxonomyIndex = &TaxonomyIndex{Sticky: iSticky}
	if has, err = Engine.Get(_TaxonomyIndex); (has == true) && (err == nil) {
		if row, err := Engine.Where("sticky = ?", iSticky).Delete(new(TaxonomyIndex)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteTaxonomyIndexViaCreated Delete TaxonomyIndex via Created
func DeleteTaxonomyIndexViaCreated(iCreated int) (err error) {
	var has bool
	var _TaxonomyIndex = &TaxonomyIndex{Created: iCreated}
	if has, err = Engine.Get(_TaxonomyIndex); (has == true) && (err == nil) {
		if row, err := Engine.Where("created = ?", iCreated).Delete(new(TaxonomyIndex)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}
