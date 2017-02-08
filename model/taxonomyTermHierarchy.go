package model

type TaxonomyTermHierarchy struct {
	Tid    int64 `xorm:"not null pk default 0 INT(10)"`
	Parent int   `xorm:"not null pk default 0 index INT(10)"`
}

// GetTaxonomyTermHierarchiesCount TaxonomyTermHierarchys' Count
func GetTaxonomyTermHierarchiesCount(offset int, limit int) (int64, error) {
	total, err := Engine.Limit(limit, offset).Count(&TaxonomyTermHierarchy{})
	return total, err
}

// GetTaxonomyTermHierarchyCountViaTid Get TaxonomyTermHierarchy via Tid
func GetTaxonomyTermHierarchyCountViaTid(iTid int64) int64 {
	n, _ := Engine.Where("tid = ?", iTid).Count(&TaxonomyTermHierarchy{Tid: iTid})
	return n
}

// GetTaxonomyTermHierarchyCountViaParent Get TaxonomyTermHierarchy via Parent
func GetTaxonomyTermHierarchyCountViaParent(iParent int) int64 {
	n, _ := Engine.Where("parent = ?", iParent).Count(&TaxonomyTermHierarchy{Parent: iParent})
	return n
}

// GetTaxonomyTermHierarchiesByTidAndParent Get TaxonomyTermHierarchies via TidAndParent
func GetTaxonomyTermHierarchiesByTidAndParent(offset int, limit int, Tid_ int64, Parent_ int) (*[]*TaxonomyTermHierarchy, error) {
	var _TaxonomyTermHierarchy = new([]*TaxonomyTermHierarchy)
	err := Engine.Table("taxonomy_term_hierarchy").Where("tid = ? and parent = ?", Tid_, Parent_).Limit(limit, offset).Find(_TaxonomyTermHierarchy)
	return _TaxonomyTermHierarchy, err
}

// GetTaxonomyTermHierarchies Get TaxonomyTermHierarchies via field
func GetTaxonomyTermHierarchies(offset int, limit int, field string) (*[]*TaxonomyTermHierarchy, error) {
	var _TaxonomyTermHierarchy = new([]*TaxonomyTermHierarchy)
	err := Engine.Table("taxonomy_term_hierarchy").Limit(limit, offset).Desc(field).Find(_TaxonomyTermHierarchy)
	return _TaxonomyTermHierarchy, err
}

// GetTaxonomyTermHierarchiesViaTid Get TaxonomyTermHierarchys via Tid
func GetTaxonomyTermHierarchiesViaTid(offset int, limit int, Tid_ int64, field string) (*[]*TaxonomyTermHierarchy, error) {
	var _TaxonomyTermHierarchy = new([]*TaxonomyTermHierarchy)
	err := Engine.Table("taxonomy_term_hierarchy").Where("tid = ?", Tid_).Limit(limit, offset).Desc(field).Find(_TaxonomyTermHierarchy)
	return _TaxonomyTermHierarchy, err
}

// GetTaxonomyTermHierarchiesViaParent Get TaxonomyTermHierarchys via Parent
func GetTaxonomyTermHierarchiesViaParent(offset int, limit int, Parent_ int, field string) (*[]*TaxonomyTermHierarchy, error) {
	var _TaxonomyTermHierarchy = new([]*TaxonomyTermHierarchy)
	err := Engine.Table("taxonomy_term_hierarchy").Where("parent = ?", Parent_).Limit(limit, offset).Desc(field).Find(_TaxonomyTermHierarchy)
	return _TaxonomyTermHierarchy, err
}

// HasTaxonomyTermHierarchyViaTid Has TaxonomyTermHierarchy via Tid
func HasTaxonomyTermHierarchyViaTid(iTid int64) bool {
	if has, err := Engine.Where("tid = ?", iTid).Get(new(TaxonomyTermHierarchy)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasTaxonomyTermHierarchyViaParent Has TaxonomyTermHierarchy via Parent
func HasTaxonomyTermHierarchyViaParent(iParent int) bool {
	if has, err := Engine.Where("parent = ?", iParent).Get(new(TaxonomyTermHierarchy)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// GetTaxonomyTermHierarchyViaTid Get TaxonomyTermHierarchy via Tid
func GetTaxonomyTermHierarchyViaTid(iTid int64) (*TaxonomyTermHierarchy, error) {
	var _TaxonomyTermHierarchy = &TaxonomyTermHierarchy{Tid: iTid}
	has, err := Engine.Get(_TaxonomyTermHierarchy)
	if has {
		return _TaxonomyTermHierarchy, err
	} else {
		return nil, err
	}
}

// GetTaxonomyTermHierarchyViaParent Get TaxonomyTermHierarchy via Parent
func GetTaxonomyTermHierarchyViaParent(iParent int) (*TaxonomyTermHierarchy, error) {
	var _TaxonomyTermHierarchy = &TaxonomyTermHierarchy{Parent: iParent}
	has, err := Engine.Get(_TaxonomyTermHierarchy)
	if has {
		return _TaxonomyTermHierarchy, err
	} else {
		return nil, err
	}
}

// SetTaxonomyTermHierarchyViaTid Set TaxonomyTermHierarchy via Tid
func SetTaxonomyTermHierarchyViaTid(iTid int64, taxonomy_term_hierarchy *TaxonomyTermHierarchy) (int64, error) {
	taxonomy_term_hierarchy.Tid = iTid
	return Engine.Insert(taxonomy_term_hierarchy)
}

// SetTaxonomyTermHierarchyViaParent Set TaxonomyTermHierarchy via Parent
func SetTaxonomyTermHierarchyViaParent(iParent int, taxonomy_term_hierarchy *TaxonomyTermHierarchy) (int64, error) {
	taxonomy_term_hierarchy.Parent = iParent
	return Engine.Insert(taxonomy_term_hierarchy)
}

// AddTaxonomyTermHierarchy Add TaxonomyTermHierarchy via all columns
func AddTaxonomyTermHierarchy(iTid int64, iParent int) error {
	_TaxonomyTermHierarchy := &TaxonomyTermHierarchy{Tid: iTid, Parent: iParent}
	if _, err := Engine.Insert(_TaxonomyTermHierarchy); err != nil {
		return err
	} else {
		return nil
	}
}

// PostTaxonomyTermHierarchy Post TaxonomyTermHierarchy via iTaxonomyTermHierarchy
func PostTaxonomyTermHierarchy(iTaxonomyTermHierarchy *TaxonomyTermHierarchy) (int64, error) {
	_, err := Engine.Insert(iTaxonomyTermHierarchy)
	return iTaxonomyTermHierarchy.Tid, err
}

// PutTaxonomyTermHierarchy Put TaxonomyTermHierarchy
func PutTaxonomyTermHierarchy(iTaxonomyTermHierarchy *TaxonomyTermHierarchy) (int64, error) {
	_, err := Engine.Id(iTaxonomyTermHierarchy.Tid).Update(iTaxonomyTermHierarchy)
	return iTaxonomyTermHierarchy.Tid, err
}

// PutTaxonomyTermHierarchyViaTid Put TaxonomyTermHierarchy via Tid
func PutTaxonomyTermHierarchyViaTid(Tid_ int64, iTaxonomyTermHierarchy *TaxonomyTermHierarchy) (int64, error) {
	row, err := Engine.Update(iTaxonomyTermHierarchy, &TaxonomyTermHierarchy{Tid: Tid_})
	return row, err
}

// PutTaxonomyTermHierarchyViaParent Put TaxonomyTermHierarchy via Parent
func PutTaxonomyTermHierarchyViaParent(Parent_ int, iTaxonomyTermHierarchy *TaxonomyTermHierarchy) (int64, error) {
	row, err := Engine.Update(iTaxonomyTermHierarchy, &TaxonomyTermHierarchy{Parent: Parent_})
	return row, err
}

// UpdateTaxonomyTermHierarchyViaTid via map[string]interface{}{}
func UpdateTaxonomyTermHierarchyViaTid(iTid int64, iTaxonomyTermHierarchyMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(TaxonomyTermHierarchy)).Where("tid = ?", iTid).Update(iTaxonomyTermHierarchyMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateTaxonomyTermHierarchyViaParent via map[string]interface{}{}
func UpdateTaxonomyTermHierarchyViaParent(iParent int, iTaxonomyTermHierarchyMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(TaxonomyTermHierarchy)).Where("parent = ?", iParent).Update(iTaxonomyTermHierarchyMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// DeleteTaxonomyTermHierarchy Delete TaxonomyTermHierarchy
func DeleteTaxonomyTermHierarchy(iTid int64) (int64, error) {
	row, err := Engine.Id(iTid).Delete(new(TaxonomyTermHierarchy))
	return row, err
}

// DeleteTaxonomyTermHierarchyViaTid Delete TaxonomyTermHierarchy via Tid
func DeleteTaxonomyTermHierarchyViaTid(iTid int64) (err error) {
	var has bool
	var _TaxonomyTermHierarchy = &TaxonomyTermHierarchy{Tid: iTid}
	if has, err = Engine.Get(_TaxonomyTermHierarchy); (has == true) && (err == nil) {
		if row, err := Engine.Where("tid = ?", iTid).Delete(new(TaxonomyTermHierarchy)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteTaxonomyTermHierarchyViaParent Delete TaxonomyTermHierarchy via Parent
func DeleteTaxonomyTermHierarchyViaParent(iParent int) (err error) {
	var has bool
	var _TaxonomyTermHierarchy = &TaxonomyTermHierarchy{Parent: iParent}
	if has, err = Engine.Get(_TaxonomyTermHierarchy); (has == true) && (err == nil) {
		if row, err := Engine.Where("parent = ?", iParent).Delete(new(TaxonomyTermHierarchy)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}
