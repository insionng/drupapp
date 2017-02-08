package model

type NodeRevision_body struct {
	Bundle      string `xorm:"not null default '' index VARCHAR(128)"`
	Deleted     int    `xorm:"not null pk default 0 TINYINT(4)"`
	EntityId    int    `xorm:"not null pk INT(10)"`
	RevisionId  int    `xorm:"not null pk index INT(10)"`
	Langcode    string `xorm:"not null pk default '' VARCHAR(32)"`
	Delta       int    `xorm:"not null pk INT(10)"`
	BodyValue   string `xorm:"not null LONGTEXT"`
	BodySummary string `xorm:"LONGTEXT"`
	BodyFormat  string `xorm:"index VARCHAR(255)"`
}

// GetNodeRevision_bodiesCount NodeRevision_bodys' Count
func GetNodeRevision_bodiesCount(offset int, limit int) (int64, error) {
	total, err := Engine.Limit(limit, offset).Count(&NodeRevision_body{})
	return total, err
}

// GetNodeRevision_bodyCountViaBundle Get NodeRevision_body via Bundle
func GetNodeRevision_bodyCountViaBundle(iBundle string) int64 {
	n, _ := Engine.Where("bundle = ?", iBundle).Count(&NodeRevision_body{Bundle: iBundle})
	return n
}

// GetNodeRevision_bodyCountViaDeleted Get NodeRevision_body via Deleted
func GetNodeRevision_bodyCountViaDeleted(iDeleted int) int64 {
	n, _ := Engine.Where("deleted = ?", iDeleted).Count(&NodeRevision_body{Deleted: iDeleted})
	return n
}

// GetNodeRevision_bodyCountViaEntityId Get NodeRevision_body via EntityId
func GetNodeRevision_bodyCountViaEntityId(iEntityId int) int64 {
	n, _ := Engine.Where("entity_id = ?", iEntityId).Count(&NodeRevision_body{EntityId: iEntityId})
	return n
}

// GetNodeRevision_bodyCountViaRevisionId Get NodeRevision_body via RevisionId
func GetNodeRevision_bodyCountViaRevisionId(iRevisionId int) int64 {
	n, _ := Engine.Where("revision_id = ?", iRevisionId).Count(&NodeRevision_body{RevisionId: iRevisionId})
	return n
}

// GetNodeRevision_bodyCountViaLangcode Get NodeRevision_body via Langcode
func GetNodeRevision_bodyCountViaLangcode(iLangcode string) int64 {
	n, _ := Engine.Where("langcode = ?", iLangcode).Count(&NodeRevision_body{Langcode: iLangcode})
	return n
}

// GetNodeRevision_bodyCountViaDelta Get NodeRevision_body via Delta
func GetNodeRevision_bodyCountViaDelta(iDelta int) int64 {
	n, _ := Engine.Where("delta = ?", iDelta).Count(&NodeRevision_body{Delta: iDelta})
	return n
}

// GetNodeRevision_bodyCountViaBodyValue Get NodeRevision_body via BodyValue
func GetNodeRevision_bodyCountViaBodyValue(iBodyValue string) int64 {
	n, _ := Engine.Where("body_value = ?", iBodyValue).Count(&NodeRevision_body{BodyValue: iBodyValue})
	return n
}

// GetNodeRevision_bodyCountViaBodySummary Get NodeRevision_body via BodySummary
func GetNodeRevision_bodyCountViaBodySummary(iBodySummary string) int64 {
	n, _ := Engine.Where("body_summary = ?", iBodySummary).Count(&NodeRevision_body{BodySummary: iBodySummary})
	return n
}

// GetNodeRevision_bodyCountViaBodyFormat Get NodeRevision_body via BodyFormat
func GetNodeRevision_bodyCountViaBodyFormat(iBodyFormat string) int64 {
	n, _ := Engine.Where("body_format = ?", iBodyFormat).Count(&NodeRevision_body{BodyFormat: iBodyFormat})
	return n
}

// GetNodeRevision_bodiesByBundleAndDeletedAndEntityId Get NodeRevision_bodies via BundleAndDeletedAndEntityId
func GetNodeRevision_bodiesByBundleAndDeletedAndEntityId(offset int, limit int, Bundle_ string, Deleted_ int, EntityId_ int) (*[]*NodeRevision_body, error) {
	var _NodeRevision_body = new([]*NodeRevision_body)
	err := Engine.Table("node_revision__body").Where("bundle = ? and deleted = ? and entity_id = ?", Bundle_, Deleted_, EntityId_).Limit(limit, offset).Find(_NodeRevision_body)
	return _NodeRevision_body, err
}

// GetNodeRevision_bodiesByBundleAndDeletedAndRevisionId Get NodeRevision_bodies via BundleAndDeletedAndRevisionId
func GetNodeRevision_bodiesByBundleAndDeletedAndRevisionId(offset int, limit int, Bundle_ string, Deleted_ int, RevisionId_ int) (*[]*NodeRevision_body, error) {
	var _NodeRevision_body = new([]*NodeRevision_body)
	err := Engine.Table("node_revision__body").Where("bundle = ? and deleted = ? and revision_id = ?", Bundle_, Deleted_, RevisionId_).Limit(limit, offset).Find(_NodeRevision_body)
	return _NodeRevision_body, err
}

// GetNodeRevision_bodiesByBundleAndDeletedAndLangcode Get NodeRevision_bodies via BundleAndDeletedAndLangcode
func GetNodeRevision_bodiesByBundleAndDeletedAndLangcode(offset int, limit int, Bundle_ string, Deleted_ int, Langcode_ string) (*[]*NodeRevision_body, error) {
	var _NodeRevision_body = new([]*NodeRevision_body)
	err := Engine.Table("node_revision__body").Where("bundle = ? and deleted = ? and langcode = ?", Bundle_, Deleted_, Langcode_).Limit(limit, offset).Find(_NodeRevision_body)
	return _NodeRevision_body, err
}

// GetNodeRevision_bodiesByBundleAndDeletedAndDelta Get NodeRevision_bodies via BundleAndDeletedAndDelta
func GetNodeRevision_bodiesByBundleAndDeletedAndDelta(offset int, limit int, Bundle_ string, Deleted_ int, Delta_ int) (*[]*NodeRevision_body, error) {
	var _NodeRevision_body = new([]*NodeRevision_body)
	err := Engine.Table("node_revision__body").Where("bundle = ? and deleted = ? and delta = ?", Bundle_, Deleted_, Delta_).Limit(limit, offset).Find(_NodeRevision_body)
	return _NodeRevision_body, err
}

// GetNodeRevision_bodiesByBundleAndDeletedAndBodyValue Get NodeRevision_bodies via BundleAndDeletedAndBodyValue
func GetNodeRevision_bodiesByBundleAndDeletedAndBodyValue(offset int, limit int, Bundle_ string, Deleted_ int, BodyValue_ string) (*[]*NodeRevision_body, error) {
	var _NodeRevision_body = new([]*NodeRevision_body)
	err := Engine.Table("node_revision__body").Where("bundle = ? and deleted = ? and body_value = ?", Bundle_, Deleted_, BodyValue_).Limit(limit, offset).Find(_NodeRevision_body)
	return _NodeRevision_body, err
}

// GetNodeRevision_bodiesByBundleAndDeletedAndBodySummary Get NodeRevision_bodies via BundleAndDeletedAndBodySummary
func GetNodeRevision_bodiesByBundleAndDeletedAndBodySummary(offset int, limit int, Bundle_ string, Deleted_ int, BodySummary_ string) (*[]*NodeRevision_body, error) {
	var _NodeRevision_body = new([]*NodeRevision_body)
	err := Engine.Table("node_revision__body").Where("bundle = ? and deleted = ? and body_summary = ?", Bundle_, Deleted_, BodySummary_).Limit(limit, offset).Find(_NodeRevision_body)
	return _NodeRevision_body, err
}

// GetNodeRevision_bodiesByBundleAndDeletedAndBodyFormat Get NodeRevision_bodies via BundleAndDeletedAndBodyFormat
func GetNodeRevision_bodiesByBundleAndDeletedAndBodyFormat(offset int, limit int, Bundle_ string, Deleted_ int, BodyFormat_ string) (*[]*NodeRevision_body, error) {
	var _NodeRevision_body = new([]*NodeRevision_body)
	err := Engine.Table("node_revision__body").Where("bundle = ? and deleted = ? and body_format = ?", Bundle_, Deleted_, BodyFormat_).Limit(limit, offset).Find(_NodeRevision_body)
	return _NodeRevision_body, err
}

// GetNodeRevision_bodiesByBundleAndEntityIdAndRevisionId Get NodeRevision_bodies via BundleAndEntityIdAndRevisionId
func GetNodeRevision_bodiesByBundleAndEntityIdAndRevisionId(offset int, limit int, Bundle_ string, EntityId_ int, RevisionId_ int) (*[]*NodeRevision_body, error) {
	var _NodeRevision_body = new([]*NodeRevision_body)
	err := Engine.Table("node_revision__body").Where("bundle = ? and entity_id = ? and revision_id = ?", Bundle_, EntityId_, RevisionId_).Limit(limit, offset).Find(_NodeRevision_body)
	return _NodeRevision_body, err
}

// GetNodeRevision_bodiesByBundleAndEntityIdAndLangcode Get NodeRevision_bodies via BundleAndEntityIdAndLangcode
func GetNodeRevision_bodiesByBundleAndEntityIdAndLangcode(offset int, limit int, Bundle_ string, EntityId_ int, Langcode_ string) (*[]*NodeRevision_body, error) {
	var _NodeRevision_body = new([]*NodeRevision_body)
	err := Engine.Table("node_revision__body").Where("bundle = ? and entity_id = ? and langcode = ?", Bundle_, EntityId_, Langcode_).Limit(limit, offset).Find(_NodeRevision_body)
	return _NodeRevision_body, err
}

// GetNodeRevision_bodiesByBundleAndEntityIdAndDelta Get NodeRevision_bodies via BundleAndEntityIdAndDelta
func GetNodeRevision_bodiesByBundleAndEntityIdAndDelta(offset int, limit int, Bundle_ string, EntityId_ int, Delta_ int) (*[]*NodeRevision_body, error) {
	var _NodeRevision_body = new([]*NodeRevision_body)
	err := Engine.Table("node_revision__body").Where("bundle = ? and entity_id = ? and delta = ?", Bundle_, EntityId_, Delta_).Limit(limit, offset).Find(_NodeRevision_body)
	return _NodeRevision_body, err
}

// GetNodeRevision_bodiesByBundleAndEntityIdAndBodyValue Get NodeRevision_bodies via BundleAndEntityIdAndBodyValue
func GetNodeRevision_bodiesByBundleAndEntityIdAndBodyValue(offset int, limit int, Bundle_ string, EntityId_ int, BodyValue_ string) (*[]*NodeRevision_body, error) {
	var _NodeRevision_body = new([]*NodeRevision_body)
	err := Engine.Table("node_revision__body").Where("bundle = ? and entity_id = ? and body_value = ?", Bundle_, EntityId_, BodyValue_).Limit(limit, offset).Find(_NodeRevision_body)
	return _NodeRevision_body, err
}

// GetNodeRevision_bodiesByBundleAndEntityIdAndBodySummary Get NodeRevision_bodies via BundleAndEntityIdAndBodySummary
func GetNodeRevision_bodiesByBundleAndEntityIdAndBodySummary(offset int, limit int, Bundle_ string, EntityId_ int, BodySummary_ string) (*[]*NodeRevision_body, error) {
	var _NodeRevision_body = new([]*NodeRevision_body)
	err := Engine.Table("node_revision__body").Where("bundle = ? and entity_id = ? and body_summary = ?", Bundle_, EntityId_, BodySummary_).Limit(limit, offset).Find(_NodeRevision_body)
	return _NodeRevision_body, err
}

// GetNodeRevision_bodiesByBundleAndEntityIdAndBodyFormat Get NodeRevision_bodies via BundleAndEntityIdAndBodyFormat
func GetNodeRevision_bodiesByBundleAndEntityIdAndBodyFormat(offset int, limit int, Bundle_ string, EntityId_ int, BodyFormat_ string) (*[]*NodeRevision_body, error) {
	var _NodeRevision_body = new([]*NodeRevision_body)
	err := Engine.Table("node_revision__body").Where("bundle = ? and entity_id = ? and body_format = ?", Bundle_, EntityId_, BodyFormat_).Limit(limit, offset).Find(_NodeRevision_body)
	return _NodeRevision_body, err
}

// GetNodeRevision_bodiesByBundleAndRevisionIdAndLangcode Get NodeRevision_bodies via BundleAndRevisionIdAndLangcode
func GetNodeRevision_bodiesByBundleAndRevisionIdAndLangcode(offset int, limit int, Bundle_ string, RevisionId_ int, Langcode_ string) (*[]*NodeRevision_body, error) {
	var _NodeRevision_body = new([]*NodeRevision_body)
	err := Engine.Table("node_revision__body").Where("bundle = ? and revision_id = ? and langcode = ?", Bundle_, RevisionId_, Langcode_).Limit(limit, offset).Find(_NodeRevision_body)
	return _NodeRevision_body, err
}

// GetNodeRevision_bodiesByBundleAndRevisionIdAndDelta Get NodeRevision_bodies via BundleAndRevisionIdAndDelta
func GetNodeRevision_bodiesByBundleAndRevisionIdAndDelta(offset int, limit int, Bundle_ string, RevisionId_ int, Delta_ int) (*[]*NodeRevision_body, error) {
	var _NodeRevision_body = new([]*NodeRevision_body)
	err := Engine.Table("node_revision__body").Where("bundle = ? and revision_id = ? and delta = ?", Bundle_, RevisionId_, Delta_).Limit(limit, offset).Find(_NodeRevision_body)
	return _NodeRevision_body, err
}

// GetNodeRevision_bodiesByBundleAndRevisionIdAndBodyValue Get NodeRevision_bodies via BundleAndRevisionIdAndBodyValue
func GetNodeRevision_bodiesByBundleAndRevisionIdAndBodyValue(offset int, limit int, Bundle_ string, RevisionId_ int, BodyValue_ string) (*[]*NodeRevision_body, error) {
	var _NodeRevision_body = new([]*NodeRevision_body)
	err := Engine.Table("node_revision__body").Where("bundle = ? and revision_id = ? and body_value = ?", Bundle_, RevisionId_, BodyValue_).Limit(limit, offset).Find(_NodeRevision_body)
	return _NodeRevision_body, err
}

// GetNodeRevision_bodiesByBundleAndRevisionIdAndBodySummary Get NodeRevision_bodies via BundleAndRevisionIdAndBodySummary
func GetNodeRevision_bodiesByBundleAndRevisionIdAndBodySummary(offset int, limit int, Bundle_ string, RevisionId_ int, BodySummary_ string) (*[]*NodeRevision_body, error) {
	var _NodeRevision_body = new([]*NodeRevision_body)
	err := Engine.Table("node_revision__body").Where("bundle = ? and revision_id = ? and body_summary = ?", Bundle_, RevisionId_, BodySummary_).Limit(limit, offset).Find(_NodeRevision_body)
	return _NodeRevision_body, err
}

// GetNodeRevision_bodiesByBundleAndRevisionIdAndBodyFormat Get NodeRevision_bodies via BundleAndRevisionIdAndBodyFormat
func GetNodeRevision_bodiesByBundleAndRevisionIdAndBodyFormat(offset int, limit int, Bundle_ string, RevisionId_ int, BodyFormat_ string) (*[]*NodeRevision_body, error) {
	var _NodeRevision_body = new([]*NodeRevision_body)
	err := Engine.Table("node_revision__body").Where("bundle = ? and revision_id = ? and body_format = ?", Bundle_, RevisionId_, BodyFormat_).Limit(limit, offset).Find(_NodeRevision_body)
	return _NodeRevision_body, err
}

// GetNodeRevision_bodiesByBundleAndLangcodeAndDelta Get NodeRevision_bodies via BundleAndLangcodeAndDelta
func GetNodeRevision_bodiesByBundleAndLangcodeAndDelta(offset int, limit int, Bundle_ string, Langcode_ string, Delta_ int) (*[]*NodeRevision_body, error) {
	var _NodeRevision_body = new([]*NodeRevision_body)
	err := Engine.Table("node_revision__body").Where("bundle = ? and langcode = ? and delta = ?", Bundle_, Langcode_, Delta_).Limit(limit, offset).Find(_NodeRevision_body)
	return _NodeRevision_body, err
}

// GetNodeRevision_bodiesByBundleAndLangcodeAndBodyValue Get NodeRevision_bodies via BundleAndLangcodeAndBodyValue
func GetNodeRevision_bodiesByBundleAndLangcodeAndBodyValue(offset int, limit int, Bundle_ string, Langcode_ string, BodyValue_ string) (*[]*NodeRevision_body, error) {
	var _NodeRevision_body = new([]*NodeRevision_body)
	err := Engine.Table("node_revision__body").Where("bundle = ? and langcode = ? and body_value = ?", Bundle_, Langcode_, BodyValue_).Limit(limit, offset).Find(_NodeRevision_body)
	return _NodeRevision_body, err
}

// GetNodeRevision_bodiesByBundleAndLangcodeAndBodySummary Get NodeRevision_bodies via BundleAndLangcodeAndBodySummary
func GetNodeRevision_bodiesByBundleAndLangcodeAndBodySummary(offset int, limit int, Bundle_ string, Langcode_ string, BodySummary_ string) (*[]*NodeRevision_body, error) {
	var _NodeRevision_body = new([]*NodeRevision_body)
	err := Engine.Table("node_revision__body").Where("bundle = ? and langcode = ? and body_summary = ?", Bundle_, Langcode_, BodySummary_).Limit(limit, offset).Find(_NodeRevision_body)
	return _NodeRevision_body, err
}

// GetNodeRevision_bodiesByBundleAndLangcodeAndBodyFormat Get NodeRevision_bodies via BundleAndLangcodeAndBodyFormat
func GetNodeRevision_bodiesByBundleAndLangcodeAndBodyFormat(offset int, limit int, Bundle_ string, Langcode_ string, BodyFormat_ string) (*[]*NodeRevision_body, error) {
	var _NodeRevision_body = new([]*NodeRevision_body)
	err := Engine.Table("node_revision__body").Where("bundle = ? and langcode = ? and body_format = ?", Bundle_, Langcode_, BodyFormat_).Limit(limit, offset).Find(_NodeRevision_body)
	return _NodeRevision_body, err
}

// GetNodeRevision_bodiesByBundleAndDeltaAndBodyValue Get NodeRevision_bodies via BundleAndDeltaAndBodyValue
func GetNodeRevision_bodiesByBundleAndDeltaAndBodyValue(offset int, limit int, Bundle_ string, Delta_ int, BodyValue_ string) (*[]*NodeRevision_body, error) {
	var _NodeRevision_body = new([]*NodeRevision_body)
	err := Engine.Table("node_revision__body").Where("bundle = ? and delta = ? and body_value = ?", Bundle_, Delta_, BodyValue_).Limit(limit, offset).Find(_NodeRevision_body)
	return _NodeRevision_body, err
}

// GetNodeRevision_bodiesByBundleAndDeltaAndBodySummary Get NodeRevision_bodies via BundleAndDeltaAndBodySummary
func GetNodeRevision_bodiesByBundleAndDeltaAndBodySummary(offset int, limit int, Bundle_ string, Delta_ int, BodySummary_ string) (*[]*NodeRevision_body, error) {
	var _NodeRevision_body = new([]*NodeRevision_body)
	err := Engine.Table("node_revision__body").Where("bundle = ? and delta = ? and body_summary = ?", Bundle_, Delta_, BodySummary_).Limit(limit, offset).Find(_NodeRevision_body)
	return _NodeRevision_body, err
}

// GetNodeRevision_bodiesByBundleAndDeltaAndBodyFormat Get NodeRevision_bodies via BundleAndDeltaAndBodyFormat
func GetNodeRevision_bodiesByBundleAndDeltaAndBodyFormat(offset int, limit int, Bundle_ string, Delta_ int, BodyFormat_ string) (*[]*NodeRevision_body, error) {
	var _NodeRevision_body = new([]*NodeRevision_body)
	err := Engine.Table("node_revision__body").Where("bundle = ? and delta = ? and body_format = ?", Bundle_, Delta_, BodyFormat_).Limit(limit, offset).Find(_NodeRevision_body)
	return _NodeRevision_body, err
}

// GetNodeRevision_bodiesByBundleAndBodyValueAndBodySummary Get NodeRevision_bodies via BundleAndBodyValueAndBodySummary
func GetNodeRevision_bodiesByBundleAndBodyValueAndBodySummary(offset int, limit int, Bundle_ string, BodyValue_ string, BodySummary_ string) (*[]*NodeRevision_body, error) {
	var _NodeRevision_body = new([]*NodeRevision_body)
	err := Engine.Table("node_revision__body").Where("bundle = ? and body_value = ? and body_summary = ?", Bundle_, BodyValue_, BodySummary_).Limit(limit, offset).Find(_NodeRevision_body)
	return _NodeRevision_body, err
}

// GetNodeRevision_bodiesByBundleAndBodyValueAndBodyFormat Get NodeRevision_bodies via BundleAndBodyValueAndBodyFormat
func GetNodeRevision_bodiesByBundleAndBodyValueAndBodyFormat(offset int, limit int, Bundle_ string, BodyValue_ string, BodyFormat_ string) (*[]*NodeRevision_body, error) {
	var _NodeRevision_body = new([]*NodeRevision_body)
	err := Engine.Table("node_revision__body").Where("bundle = ? and body_value = ? and body_format = ?", Bundle_, BodyValue_, BodyFormat_).Limit(limit, offset).Find(_NodeRevision_body)
	return _NodeRevision_body, err
}

// GetNodeRevision_bodiesByBundleAndBodySummaryAndBodyFormat Get NodeRevision_bodies via BundleAndBodySummaryAndBodyFormat
func GetNodeRevision_bodiesByBundleAndBodySummaryAndBodyFormat(offset int, limit int, Bundle_ string, BodySummary_ string, BodyFormat_ string) (*[]*NodeRevision_body, error) {
	var _NodeRevision_body = new([]*NodeRevision_body)
	err := Engine.Table("node_revision__body").Where("bundle = ? and body_summary = ? and body_format = ?", Bundle_, BodySummary_, BodyFormat_).Limit(limit, offset).Find(_NodeRevision_body)
	return _NodeRevision_body, err
}

// GetNodeRevision_bodiesByDeletedAndEntityIdAndRevisionId Get NodeRevision_bodies via DeletedAndEntityIdAndRevisionId
func GetNodeRevision_bodiesByDeletedAndEntityIdAndRevisionId(offset int, limit int, Deleted_ int, EntityId_ int, RevisionId_ int) (*[]*NodeRevision_body, error) {
	var _NodeRevision_body = new([]*NodeRevision_body)
	err := Engine.Table("node_revision__body").Where("deleted = ? and entity_id = ? and revision_id = ?", Deleted_, EntityId_, RevisionId_).Limit(limit, offset).Find(_NodeRevision_body)
	return _NodeRevision_body, err
}

// GetNodeRevision_bodiesByDeletedAndEntityIdAndLangcode Get NodeRevision_bodies via DeletedAndEntityIdAndLangcode
func GetNodeRevision_bodiesByDeletedAndEntityIdAndLangcode(offset int, limit int, Deleted_ int, EntityId_ int, Langcode_ string) (*[]*NodeRevision_body, error) {
	var _NodeRevision_body = new([]*NodeRevision_body)
	err := Engine.Table("node_revision__body").Where("deleted = ? and entity_id = ? and langcode = ?", Deleted_, EntityId_, Langcode_).Limit(limit, offset).Find(_NodeRevision_body)
	return _NodeRevision_body, err
}

// GetNodeRevision_bodiesByDeletedAndEntityIdAndDelta Get NodeRevision_bodies via DeletedAndEntityIdAndDelta
func GetNodeRevision_bodiesByDeletedAndEntityIdAndDelta(offset int, limit int, Deleted_ int, EntityId_ int, Delta_ int) (*[]*NodeRevision_body, error) {
	var _NodeRevision_body = new([]*NodeRevision_body)
	err := Engine.Table("node_revision__body").Where("deleted = ? and entity_id = ? and delta = ?", Deleted_, EntityId_, Delta_).Limit(limit, offset).Find(_NodeRevision_body)
	return _NodeRevision_body, err
}

// GetNodeRevision_bodiesByDeletedAndEntityIdAndBodyValue Get NodeRevision_bodies via DeletedAndEntityIdAndBodyValue
func GetNodeRevision_bodiesByDeletedAndEntityIdAndBodyValue(offset int, limit int, Deleted_ int, EntityId_ int, BodyValue_ string) (*[]*NodeRevision_body, error) {
	var _NodeRevision_body = new([]*NodeRevision_body)
	err := Engine.Table("node_revision__body").Where("deleted = ? and entity_id = ? and body_value = ?", Deleted_, EntityId_, BodyValue_).Limit(limit, offset).Find(_NodeRevision_body)
	return _NodeRevision_body, err
}

// GetNodeRevision_bodiesByDeletedAndEntityIdAndBodySummary Get NodeRevision_bodies via DeletedAndEntityIdAndBodySummary
func GetNodeRevision_bodiesByDeletedAndEntityIdAndBodySummary(offset int, limit int, Deleted_ int, EntityId_ int, BodySummary_ string) (*[]*NodeRevision_body, error) {
	var _NodeRevision_body = new([]*NodeRevision_body)
	err := Engine.Table("node_revision__body").Where("deleted = ? and entity_id = ? and body_summary = ?", Deleted_, EntityId_, BodySummary_).Limit(limit, offset).Find(_NodeRevision_body)
	return _NodeRevision_body, err
}

// GetNodeRevision_bodiesByDeletedAndEntityIdAndBodyFormat Get NodeRevision_bodies via DeletedAndEntityIdAndBodyFormat
func GetNodeRevision_bodiesByDeletedAndEntityIdAndBodyFormat(offset int, limit int, Deleted_ int, EntityId_ int, BodyFormat_ string) (*[]*NodeRevision_body, error) {
	var _NodeRevision_body = new([]*NodeRevision_body)
	err := Engine.Table("node_revision__body").Where("deleted = ? and entity_id = ? and body_format = ?", Deleted_, EntityId_, BodyFormat_).Limit(limit, offset).Find(_NodeRevision_body)
	return _NodeRevision_body, err
}

// GetNodeRevision_bodiesByDeletedAndRevisionIdAndLangcode Get NodeRevision_bodies via DeletedAndRevisionIdAndLangcode
func GetNodeRevision_bodiesByDeletedAndRevisionIdAndLangcode(offset int, limit int, Deleted_ int, RevisionId_ int, Langcode_ string) (*[]*NodeRevision_body, error) {
	var _NodeRevision_body = new([]*NodeRevision_body)
	err := Engine.Table("node_revision__body").Where("deleted = ? and revision_id = ? and langcode = ?", Deleted_, RevisionId_, Langcode_).Limit(limit, offset).Find(_NodeRevision_body)
	return _NodeRevision_body, err
}

// GetNodeRevision_bodiesByDeletedAndRevisionIdAndDelta Get NodeRevision_bodies via DeletedAndRevisionIdAndDelta
func GetNodeRevision_bodiesByDeletedAndRevisionIdAndDelta(offset int, limit int, Deleted_ int, RevisionId_ int, Delta_ int) (*[]*NodeRevision_body, error) {
	var _NodeRevision_body = new([]*NodeRevision_body)
	err := Engine.Table("node_revision__body").Where("deleted = ? and revision_id = ? and delta = ?", Deleted_, RevisionId_, Delta_).Limit(limit, offset).Find(_NodeRevision_body)
	return _NodeRevision_body, err
}

// GetNodeRevision_bodiesByDeletedAndRevisionIdAndBodyValue Get NodeRevision_bodies via DeletedAndRevisionIdAndBodyValue
func GetNodeRevision_bodiesByDeletedAndRevisionIdAndBodyValue(offset int, limit int, Deleted_ int, RevisionId_ int, BodyValue_ string) (*[]*NodeRevision_body, error) {
	var _NodeRevision_body = new([]*NodeRevision_body)
	err := Engine.Table("node_revision__body").Where("deleted = ? and revision_id = ? and body_value = ?", Deleted_, RevisionId_, BodyValue_).Limit(limit, offset).Find(_NodeRevision_body)
	return _NodeRevision_body, err
}

// GetNodeRevision_bodiesByDeletedAndRevisionIdAndBodySummary Get NodeRevision_bodies via DeletedAndRevisionIdAndBodySummary
func GetNodeRevision_bodiesByDeletedAndRevisionIdAndBodySummary(offset int, limit int, Deleted_ int, RevisionId_ int, BodySummary_ string) (*[]*NodeRevision_body, error) {
	var _NodeRevision_body = new([]*NodeRevision_body)
	err := Engine.Table("node_revision__body").Where("deleted = ? and revision_id = ? and body_summary = ?", Deleted_, RevisionId_, BodySummary_).Limit(limit, offset).Find(_NodeRevision_body)
	return _NodeRevision_body, err
}

// GetNodeRevision_bodiesByDeletedAndRevisionIdAndBodyFormat Get NodeRevision_bodies via DeletedAndRevisionIdAndBodyFormat
func GetNodeRevision_bodiesByDeletedAndRevisionIdAndBodyFormat(offset int, limit int, Deleted_ int, RevisionId_ int, BodyFormat_ string) (*[]*NodeRevision_body, error) {
	var _NodeRevision_body = new([]*NodeRevision_body)
	err := Engine.Table("node_revision__body").Where("deleted = ? and revision_id = ? and body_format = ?", Deleted_, RevisionId_, BodyFormat_).Limit(limit, offset).Find(_NodeRevision_body)
	return _NodeRevision_body, err
}

// GetNodeRevision_bodiesByDeletedAndLangcodeAndDelta Get NodeRevision_bodies via DeletedAndLangcodeAndDelta
func GetNodeRevision_bodiesByDeletedAndLangcodeAndDelta(offset int, limit int, Deleted_ int, Langcode_ string, Delta_ int) (*[]*NodeRevision_body, error) {
	var _NodeRevision_body = new([]*NodeRevision_body)
	err := Engine.Table("node_revision__body").Where("deleted = ? and langcode = ? and delta = ?", Deleted_, Langcode_, Delta_).Limit(limit, offset).Find(_NodeRevision_body)
	return _NodeRevision_body, err
}

// GetNodeRevision_bodiesByDeletedAndLangcodeAndBodyValue Get NodeRevision_bodies via DeletedAndLangcodeAndBodyValue
func GetNodeRevision_bodiesByDeletedAndLangcodeAndBodyValue(offset int, limit int, Deleted_ int, Langcode_ string, BodyValue_ string) (*[]*NodeRevision_body, error) {
	var _NodeRevision_body = new([]*NodeRevision_body)
	err := Engine.Table("node_revision__body").Where("deleted = ? and langcode = ? and body_value = ?", Deleted_, Langcode_, BodyValue_).Limit(limit, offset).Find(_NodeRevision_body)
	return _NodeRevision_body, err
}

// GetNodeRevision_bodiesByDeletedAndLangcodeAndBodySummary Get NodeRevision_bodies via DeletedAndLangcodeAndBodySummary
func GetNodeRevision_bodiesByDeletedAndLangcodeAndBodySummary(offset int, limit int, Deleted_ int, Langcode_ string, BodySummary_ string) (*[]*NodeRevision_body, error) {
	var _NodeRevision_body = new([]*NodeRevision_body)
	err := Engine.Table("node_revision__body").Where("deleted = ? and langcode = ? and body_summary = ?", Deleted_, Langcode_, BodySummary_).Limit(limit, offset).Find(_NodeRevision_body)
	return _NodeRevision_body, err
}

// GetNodeRevision_bodiesByDeletedAndLangcodeAndBodyFormat Get NodeRevision_bodies via DeletedAndLangcodeAndBodyFormat
func GetNodeRevision_bodiesByDeletedAndLangcodeAndBodyFormat(offset int, limit int, Deleted_ int, Langcode_ string, BodyFormat_ string) (*[]*NodeRevision_body, error) {
	var _NodeRevision_body = new([]*NodeRevision_body)
	err := Engine.Table("node_revision__body").Where("deleted = ? and langcode = ? and body_format = ?", Deleted_, Langcode_, BodyFormat_).Limit(limit, offset).Find(_NodeRevision_body)
	return _NodeRevision_body, err
}

// GetNodeRevision_bodiesByDeletedAndDeltaAndBodyValue Get NodeRevision_bodies via DeletedAndDeltaAndBodyValue
func GetNodeRevision_bodiesByDeletedAndDeltaAndBodyValue(offset int, limit int, Deleted_ int, Delta_ int, BodyValue_ string) (*[]*NodeRevision_body, error) {
	var _NodeRevision_body = new([]*NodeRevision_body)
	err := Engine.Table("node_revision__body").Where("deleted = ? and delta = ? and body_value = ?", Deleted_, Delta_, BodyValue_).Limit(limit, offset).Find(_NodeRevision_body)
	return _NodeRevision_body, err
}

// GetNodeRevision_bodiesByDeletedAndDeltaAndBodySummary Get NodeRevision_bodies via DeletedAndDeltaAndBodySummary
func GetNodeRevision_bodiesByDeletedAndDeltaAndBodySummary(offset int, limit int, Deleted_ int, Delta_ int, BodySummary_ string) (*[]*NodeRevision_body, error) {
	var _NodeRevision_body = new([]*NodeRevision_body)
	err := Engine.Table("node_revision__body").Where("deleted = ? and delta = ? and body_summary = ?", Deleted_, Delta_, BodySummary_).Limit(limit, offset).Find(_NodeRevision_body)
	return _NodeRevision_body, err
}

// GetNodeRevision_bodiesByDeletedAndDeltaAndBodyFormat Get NodeRevision_bodies via DeletedAndDeltaAndBodyFormat
func GetNodeRevision_bodiesByDeletedAndDeltaAndBodyFormat(offset int, limit int, Deleted_ int, Delta_ int, BodyFormat_ string) (*[]*NodeRevision_body, error) {
	var _NodeRevision_body = new([]*NodeRevision_body)
	err := Engine.Table("node_revision__body").Where("deleted = ? and delta = ? and body_format = ?", Deleted_, Delta_, BodyFormat_).Limit(limit, offset).Find(_NodeRevision_body)
	return _NodeRevision_body, err
}

// GetNodeRevision_bodiesByDeletedAndBodyValueAndBodySummary Get NodeRevision_bodies via DeletedAndBodyValueAndBodySummary
func GetNodeRevision_bodiesByDeletedAndBodyValueAndBodySummary(offset int, limit int, Deleted_ int, BodyValue_ string, BodySummary_ string) (*[]*NodeRevision_body, error) {
	var _NodeRevision_body = new([]*NodeRevision_body)
	err := Engine.Table("node_revision__body").Where("deleted = ? and body_value = ? and body_summary = ?", Deleted_, BodyValue_, BodySummary_).Limit(limit, offset).Find(_NodeRevision_body)
	return _NodeRevision_body, err
}

// GetNodeRevision_bodiesByDeletedAndBodyValueAndBodyFormat Get NodeRevision_bodies via DeletedAndBodyValueAndBodyFormat
func GetNodeRevision_bodiesByDeletedAndBodyValueAndBodyFormat(offset int, limit int, Deleted_ int, BodyValue_ string, BodyFormat_ string) (*[]*NodeRevision_body, error) {
	var _NodeRevision_body = new([]*NodeRevision_body)
	err := Engine.Table("node_revision__body").Where("deleted = ? and body_value = ? and body_format = ?", Deleted_, BodyValue_, BodyFormat_).Limit(limit, offset).Find(_NodeRevision_body)
	return _NodeRevision_body, err
}

// GetNodeRevision_bodiesByDeletedAndBodySummaryAndBodyFormat Get NodeRevision_bodies via DeletedAndBodySummaryAndBodyFormat
func GetNodeRevision_bodiesByDeletedAndBodySummaryAndBodyFormat(offset int, limit int, Deleted_ int, BodySummary_ string, BodyFormat_ string) (*[]*NodeRevision_body, error) {
	var _NodeRevision_body = new([]*NodeRevision_body)
	err := Engine.Table("node_revision__body").Where("deleted = ? and body_summary = ? and body_format = ?", Deleted_, BodySummary_, BodyFormat_).Limit(limit, offset).Find(_NodeRevision_body)
	return _NodeRevision_body, err
}

// GetNodeRevision_bodiesByEntityIdAndRevisionIdAndLangcode Get NodeRevision_bodies via EntityIdAndRevisionIdAndLangcode
func GetNodeRevision_bodiesByEntityIdAndRevisionIdAndLangcode(offset int, limit int, EntityId_ int, RevisionId_ int, Langcode_ string) (*[]*NodeRevision_body, error) {
	var _NodeRevision_body = new([]*NodeRevision_body)
	err := Engine.Table("node_revision__body").Where("entity_id = ? and revision_id = ? and langcode = ?", EntityId_, RevisionId_, Langcode_).Limit(limit, offset).Find(_NodeRevision_body)
	return _NodeRevision_body, err
}

// GetNodeRevision_bodiesByEntityIdAndRevisionIdAndDelta Get NodeRevision_bodies via EntityIdAndRevisionIdAndDelta
func GetNodeRevision_bodiesByEntityIdAndRevisionIdAndDelta(offset int, limit int, EntityId_ int, RevisionId_ int, Delta_ int) (*[]*NodeRevision_body, error) {
	var _NodeRevision_body = new([]*NodeRevision_body)
	err := Engine.Table("node_revision__body").Where("entity_id = ? and revision_id = ? and delta = ?", EntityId_, RevisionId_, Delta_).Limit(limit, offset).Find(_NodeRevision_body)
	return _NodeRevision_body, err
}

// GetNodeRevision_bodiesByEntityIdAndRevisionIdAndBodyValue Get NodeRevision_bodies via EntityIdAndRevisionIdAndBodyValue
func GetNodeRevision_bodiesByEntityIdAndRevisionIdAndBodyValue(offset int, limit int, EntityId_ int, RevisionId_ int, BodyValue_ string) (*[]*NodeRevision_body, error) {
	var _NodeRevision_body = new([]*NodeRevision_body)
	err := Engine.Table("node_revision__body").Where("entity_id = ? and revision_id = ? and body_value = ?", EntityId_, RevisionId_, BodyValue_).Limit(limit, offset).Find(_NodeRevision_body)
	return _NodeRevision_body, err
}

// GetNodeRevision_bodiesByEntityIdAndRevisionIdAndBodySummary Get NodeRevision_bodies via EntityIdAndRevisionIdAndBodySummary
func GetNodeRevision_bodiesByEntityIdAndRevisionIdAndBodySummary(offset int, limit int, EntityId_ int, RevisionId_ int, BodySummary_ string) (*[]*NodeRevision_body, error) {
	var _NodeRevision_body = new([]*NodeRevision_body)
	err := Engine.Table("node_revision__body").Where("entity_id = ? and revision_id = ? and body_summary = ?", EntityId_, RevisionId_, BodySummary_).Limit(limit, offset).Find(_NodeRevision_body)
	return _NodeRevision_body, err
}

// GetNodeRevision_bodiesByEntityIdAndRevisionIdAndBodyFormat Get NodeRevision_bodies via EntityIdAndRevisionIdAndBodyFormat
func GetNodeRevision_bodiesByEntityIdAndRevisionIdAndBodyFormat(offset int, limit int, EntityId_ int, RevisionId_ int, BodyFormat_ string) (*[]*NodeRevision_body, error) {
	var _NodeRevision_body = new([]*NodeRevision_body)
	err := Engine.Table("node_revision__body").Where("entity_id = ? and revision_id = ? and body_format = ?", EntityId_, RevisionId_, BodyFormat_).Limit(limit, offset).Find(_NodeRevision_body)
	return _NodeRevision_body, err
}

// GetNodeRevision_bodiesByEntityIdAndLangcodeAndDelta Get NodeRevision_bodies via EntityIdAndLangcodeAndDelta
func GetNodeRevision_bodiesByEntityIdAndLangcodeAndDelta(offset int, limit int, EntityId_ int, Langcode_ string, Delta_ int) (*[]*NodeRevision_body, error) {
	var _NodeRevision_body = new([]*NodeRevision_body)
	err := Engine.Table("node_revision__body").Where("entity_id = ? and langcode = ? and delta = ?", EntityId_, Langcode_, Delta_).Limit(limit, offset).Find(_NodeRevision_body)
	return _NodeRevision_body, err
}

// GetNodeRevision_bodiesByEntityIdAndLangcodeAndBodyValue Get NodeRevision_bodies via EntityIdAndLangcodeAndBodyValue
func GetNodeRevision_bodiesByEntityIdAndLangcodeAndBodyValue(offset int, limit int, EntityId_ int, Langcode_ string, BodyValue_ string) (*[]*NodeRevision_body, error) {
	var _NodeRevision_body = new([]*NodeRevision_body)
	err := Engine.Table("node_revision__body").Where("entity_id = ? and langcode = ? and body_value = ?", EntityId_, Langcode_, BodyValue_).Limit(limit, offset).Find(_NodeRevision_body)
	return _NodeRevision_body, err
}

// GetNodeRevision_bodiesByEntityIdAndLangcodeAndBodySummary Get NodeRevision_bodies via EntityIdAndLangcodeAndBodySummary
func GetNodeRevision_bodiesByEntityIdAndLangcodeAndBodySummary(offset int, limit int, EntityId_ int, Langcode_ string, BodySummary_ string) (*[]*NodeRevision_body, error) {
	var _NodeRevision_body = new([]*NodeRevision_body)
	err := Engine.Table("node_revision__body").Where("entity_id = ? and langcode = ? and body_summary = ?", EntityId_, Langcode_, BodySummary_).Limit(limit, offset).Find(_NodeRevision_body)
	return _NodeRevision_body, err
}

// GetNodeRevision_bodiesByEntityIdAndLangcodeAndBodyFormat Get NodeRevision_bodies via EntityIdAndLangcodeAndBodyFormat
func GetNodeRevision_bodiesByEntityIdAndLangcodeAndBodyFormat(offset int, limit int, EntityId_ int, Langcode_ string, BodyFormat_ string) (*[]*NodeRevision_body, error) {
	var _NodeRevision_body = new([]*NodeRevision_body)
	err := Engine.Table("node_revision__body").Where("entity_id = ? and langcode = ? and body_format = ?", EntityId_, Langcode_, BodyFormat_).Limit(limit, offset).Find(_NodeRevision_body)
	return _NodeRevision_body, err
}

// GetNodeRevision_bodiesByEntityIdAndDeltaAndBodyValue Get NodeRevision_bodies via EntityIdAndDeltaAndBodyValue
func GetNodeRevision_bodiesByEntityIdAndDeltaAndBodyValue(offset int, limit int, EntityId_ int, Delta_ int, BodyValue_ string) (*[]*NodeRevision_body, error) {
	var _NodeRevision_body = new([]*NodeRevision_body)
	err := Engine.Table("node_revision__body").Where("entity_id = ? and delta = ? and body_value = ?", EntityId_, Delta_, BodyValue_).Limit(limit, offset).Find(_NodeRevision_body)
	return _NodeRevision_body, err
}

// GetNodeRevision_bodiesByEntityIdAndDeltaAndBodySummary Get NodeRevision_bodies via EntityIdAndDeltaAndBodySummary
func GetNodeRevision_bodiesByEntityIdAndDeltaAndBodySummary(offset int, limit int, EntityId_ int, Delta_ int, BodySummary_ string) (*[]*NodeRevision_body, error) {
	var _NodeRevision_body = new([]*NodeRevision_body)
	err := Engine.Table("node_revision__body").Where("entity_id = ? and delta = ? and body_summary = ?", EntityId_, Delta_, BodySummary_).Limit(limit, offset).Find(_NodeRevision_body)
	return _NodeRevision_body, err
}

// GetNodeRevision_bodiesByEntityIdAndDeltaAndBodyFormat Get NodeRevision_bodies via EntityIdAndDeltaAndBodyFormat
func GetNodeRevision_bodiesByEntityIdAndDeltaAndBodyFormat(offset int, limit int, EntityId_ int, Delta_ int, BodyFormat_ string) (*[]*NodeRevision_body, error) {
	var _NodeRevision_body = new([]*NodeRevision_body)
	err := Engine.Table("node_revision__body").Where("entity_id = ? and delta = ? and body_format = ?", EntityId_, Delta_, BodyFormat_).Limit(limit, offset).Find(_NodeRevision_body)
	return _NodeRevision_body, err
}

// GetNodeRevision_bodiesByEntityIdAndBodyValueAndBodySummary Get NodeRevision_bodies via EntityIdAndBodyValueAndBodySummary
func GetNodeRevision_bodiesByEntityIdAndBodyValueAndBodySummary(offset int, limit int, EntityId_ int, BodyValue_ string, BodySummary_ string) (*[]*NodeRevision_body, error) {
	var _NodeRevision_body = new([]*NodeRevision_body)
	err := Engine.Table("node_revision__body").Where("entity_id = ? and body_value = ? and body_summary = ?", EntityId_, BodyValue_, BodySummary_).Limit(limit, offset).Find(_NodeRevision_body)
	return _NodeRevision_body, err
}

// GetNodeRevision_bodiesByEntityIdAndBodyValueAndBodyFormat Get NodeRevision_bodies via EntityIdAndBodyValueAndBodyFormat
func GetNodeRevision_bodiesByEntityIdAndBodyValueAndBodyFormat(offset int, limit int, EntityId_ int, BodyValue_ string, BodyFormat_ string) (*[]*NodeRevision_body, error) {
	var _NodeRevision_body = new([]*NodeRevision_body)
	err := Engine.Table("node_revision__body").Where("entity_id = ? and body_value = ? and body_format = ?", EntityId_, BodyValue_, BodyFormat_).Limit(limit, offset).Find(_NodeRevision_body)
	return _NodeRevision_body, err
}

// GetNodeRevision_bodiesByEntityIdAndBodySummaryAndBodyFormat Get NodeRevision_bodies via EntityIdAndBodySummaryAndBodyFormat
func GetNodeRevision_bodiesByEntityIdAndBodySummaryAndBodyFormat(offset int, limit int, EntityId_ int, BodySummary_ string, BodyFormat_ string) (*[]*NodeRevision_body, error) {
	var _NodeRevision_body = new([]*NodeRevision_body)
	err := Engine.Table("node_revision__body").Where("entity_id = ? and body_summary = ? and body_format = ?", EntityId_, BodySummary_, BodyFormat_).Limit(limit, offset).Find(_NodeRevision_body)
	return _NodeRevision_body, err
}

// GetNodeRevision_bodiesByRevisionIdAndLangcodeAndDelta Get NodeRevision_bodies via RevisionIdAndLangcodeAndDelta
func GetNodeRevision_bodiesByRevisionIdAndLangcodeAndDelta(offset int, limit int, RevisionId_ int, Langcode_ string, Delta_ int) (*[]*NodeRevision_body, error) {
	var _NodeRevision_body = new([]*NodeRevision_body)
	err := Engine.Table("node_revision__body").Where("revision_id = ? and langcode = ? and delta = ?", RevisionId_, Langcode_, Delta_).Limit(limit, offset).Find(_NodeRevision_body)
	return _NodeRevision_body, err
}

// GetNodeRevision_bodiesByRevisionIdAndLangcodeAndBodyValue Get NodeRevision_bodies via RevisionIdAndLangcodeAndBodyValue
func GetNodeRevision_bodiesByRevisionIdAndLangcodeAndBodyValue(offset int, limit int, RevisionId_ int, Langcode_ string, BodyValue_ string) (*[]*NodeRevision_body, error) {
	var _NodeRevision_body = new([]*NodeRevision_body)
	err := Engine.Table("node_revision__body").Where("revision_id = ? and langcode = ? and body_value = ?", RevisionId_, Langcode_, BodyValue_).Limit(limit, offset).Find(_NodeRevision_body)
	return _NodeRevision_body, err
}

// GetNodeRevision_bodiesByRevisionIdAndLangcodeAndBodySummary Get NodeRevision_bodies via RevisionIdAndLangcodeAndBodySummary
func GetNodeRevision_bodiesByRevisionIdAndLangcodeAndBodySummary(offset int, limit int, RevisionId_ int, Langcode_ string, BodySummary_ string) (*[]*NodeRevision_body, error) {
	var _NodeRevision_body = new([]*NodeRevision_body)
	err := Engine.Table("node_revision__body").Where("revision_id = ? and langcode = ? and body_summary = ?", RevisionId_, Langcode_, BodySummary_).Limit(limit, offset).Find(_NodeRevision_body)
	return _NodeRevision_body, err
}

// GetNodeRevision_bodiesByRevisionIdAndLangcodeAndBodyFormat Get NodeRevision_bodies via RevisionIdAndLangcodeAndBodyFormat
func GetNodeRevision_bodiesByRevisionIdAndLangcodeAndBodyFormat(offset int, limit int, RevisionId_ int, Langcode_ string, BodyFormat_ string) (*[]*NodeRevision_body, error) {
	var _NodeRevision_body = new([]*NodeRevision_body)
	err := Engine.Table("node_revision__body").Where("revision_id = ? and langcode = ? and body_format = ?", RevisionId_, Langcode_, BodyFormat_).Limit(limit, offset).Find(_NodeRevision_body)
	return _NodeRevision_body, err
}

// GetNodeRevision_bodiesByRevisionIdAndDeltaAndBodyValue Get NodeRevision_bodies via RevisionIdAndDeltaAndBodyValue
func GetNodeRevision_bodiesByRevisionIdAndDeltaAndBodyValue(offset int, limit int, RevisionId_ int, Delta_ int, BodyValue_ string) (*[]*NodeRevision_body, error) {
	var _NodeRevision_body = new([]*NodeRevision_body)
	err := Engine.Table("node_revision__body").Where("revision_id = ? and delta = ? and body_value = ?", RevisionId_, Delta_, BodyValue_).Limit(limit, offset).Find(_NodeRevision_body)
	return _NodeRevision_body, err
}

// GetNodeRevision_bodiesByRevisionIdAndDeltaAndBodySummary Get NodeRevision_bodies via RevisionIdAndDeltaAndBodySummary
func GetNodeRevision_bodiesByRevisionIdAndDeltaAndBodySummary(offset int, limit int, RevisionId_ int, Delta_ int, BodySummary_ string) (*[]*NodeRevision_body, error) {
	var _NodeRevision_body = new([]*NodeRevision_body)
	err := Engine.Table("node_revision__body").Where("revision_id = ? and delta = ? and body_summary = ?", RevisionId_, Delta_, BodySummary_).Limit(limit, offset).Find(_NodeRevision_body)
	return _NodeRevision_body, err
}

// GetNodeRevision_bodiesByRevisionIdAndDeltaAndBodyFormat Get NodeRevision_bodies via RevisionIdAndDeltaAndBodyFormat
func GetNodeRevision_bodiesByRevisionIdAndDeltaAndBodyFormat(offset int, limit int, RevisionId_ int, Delta_ int, BodyFormat_ string) (*[]*NodeRevision_body, error) {
	var _NodeRevision_body = new([]*NodeRevision_body)
	err := Engine.Table("node_revision__body").Where("revision_id = ? and delta = ? and body_format = ?", RevisionId_, Delta_, BodyFormat_).Limit(limit, offset).Find(_NodeRevision_body)
	return _NodeRevision_body, err
}

// GetNodeRevision_bodiesByRevisionIdAndBodyValueAndBodySummary Get NodeRevision_bodies via RevisionIdAndBodyValueAndBodySummary
func GetNodeRevision_bodiesByRevisionIdAndBodyValueAndBodySummary(offset int, limit int, RevisionId_ int, BodyValue_ string, BodySummary_ string) (*[]*NodeRevision_body, error) {
	var _NodeRevision_body = new([]*NodeRevision_body)
	err := Engine.Table("node_revision__body").Where("revision_id = ? and body_value = ? and body_summary = ?", RevisionId_, BodyValue_, BodySummary_).Limit(limit, offset).Find(_NodeRevision_body)
	return _NodeRevision_body, err
}

// GetNodeRevision_bodiesByRevisionIdAndBodyValueAndBodyFormat Get NodeRevision_bodies via RevisionIdAndBodyValueAndBodyFormat
func GetNodeRevision_bodiesByRevisionIdAndBodyValueAndBodyFormat(offset int, limit int, RevisionId_ int, BodyValue_ string, BodyFormat_ string) (*[]*NodeRevision_body, error) {
	var _NodeRevision_body = new([]*NodeRevision_body)
	err := Engine.Table("node_revision__body").Where("revision_id = ? and body_value = ? and body_format = ?", RevisionId_, BodyValue_, BodyFormat_).Limit(limit, offset).Find(_NodeRevision_body)
	return _NodeRevision_body, err
}

// GetNodeRevision_bodiesByRevisionIdAndBodySummaryAndBodyFormat Get NodeRevision_bodies via RevisionIdAndBodySummaryAndBodyFormat
func GetNodeRevision_bodiesByRevisionIdAndBodySummaryAndBodyFormat(offset int, limit int, RevisionId_ int, BodySummary_ string, BodyFormat_ string) (*[]*NodeRevision_body, error) {
	var _NodeRevision_body = new([]*NodeRevision_body)
	err := Engine.Table("node_revision__body").Where("revision_id = ? and body_summary = ? and body_format = ?", RevisionId_, BodySummary_, BodyFormat_).Limit(limit, offset).Find(_NodeRevision_body)
	return _NodeRevision_body, err
}

// GetNodeRevision_bodiesByLangcodeAndDeltaAndBodyValue Get NodeRevision_bodies via LangcodeAndDeltaAndBodyValue
func GetNodeRevision_bodiesByLangcodeAndDeltaAndBodyValue(offset int, limit int, Langcode_ string, Delta_ int, BodyValue_ string) (*[]*NodeRevision_body, error) {
	var _NodeRevision_body = new([]*NodeRevision_body)
	err := Engine.Table("node_revision__body").Where("langcode = ? and delta = ? and body_value = ?", Langcode_, Delta_, BodyValue_).Limit(limit, offset).Find(_NodeRevision_body)
	return _NodeRevision_body, err
}

// GetNodeRevision_bodiesByLangcodeAndDeltaAndBodySummary Get NodeRevision_bodies via LangcodeAndDeltaAndBodySummary
func GetNodeRevision_bodiesByLangcodeAndDeltaAndBodySummary(offset int, limit int, Langcode_ string, Delta_ int, BodySummary_ string) (*[]*NodeRevision_body, error) {
	var _NodeRevision_body = new([]*NodeRevision_body)
	err := Engine.Table("node_revision__body").Where("langcode = ? and delta = ? and body_summary = ?", Langcode_, Delta_, BodySummary_).Limit(limit, offset).Find(_NodeRevision_body)
	return _NodeRevision_body, err
}

// GetNodeRevision_bodiesByLangcodeAndDeltaAndBodyFormat Get NodeRevision_bodies via LangcodeAndDeltaAndBodyFormat
func GetNodeRevision_bodiesByLangcodeAndDeltaAndBodyFormat(offset int, limit int, Langcode_ string, Delta_ int, BodyFormat_ string) (*[]*NodeRevision_body, error) {
	var _NodeRevision_body = new([]*NodeRevision_body)
	err := Engine.Table("node_revision__body").Where("langcode = ? and delta = ? and body_format = ?", Langcode_, Delta_, BodyFormat_).Limit(limit, offset).Find(_NodeRevision_body)
	return _NodeRevision_body, err
}

// GetNodeRevision_bodiesByLangcodeAndBodyValueAndBodySummary Get NodeRevision_bodies via LangcodeAndBodyValueAndBodySummary
func GetNodeRevision_bodiesByLangcodeAndBodyValueAndBodySummary(offset int, limit int, Langcode_ string, BodyValue_ string, BodySummary_ string) (*[]*NodeRevision_body, error) {
	var _NodeRevision_body = new([]*NodeRevision_body)
	err := Engine.Table("node_revision__body").Where("langcode = ? and body_value = ? and body_summary = ?", Langcode_, BodyValue_, BodySummary_).Limit(limit, offset).Find(_NodeRevision_body)
	return _NodeRevision_body, err
}

// GetNodeRevision_bodiesByLangcodeAndBodyValueAndBodyFormat Get NodeRevision_bodies via LangcodeAndBodyValueAndBodyFormat
func GetNodeRevision_bodiesByLangcodeAndBodyValueAndBodyFormat(offset int, limit int, Langcode_ string, BodyValue_ string, BodyFormat_ string) (*[]*NodeRevision_body, error) {
	var _NodeRevision_body = new([]*NodeRevision_body)
	err := Engine.Table("node_revision__body").Where("langcode = ? and body_value = ? and body_format = ?", Langcode_, BodyValue_, BodyFormat_).Limit(limit, offset).Find(_NodeRevision_body)
	return _NodeRevision_body, err
}

// GetNodeRevision_bodiesByLangcodeAndBodySummaryAndBodyFormat Get NodeRevision_bodies via LangcodeAndBodySummaryAndBodyFormat
func GetNodeRevision_bodiesByLangcodeAndBodySummaryAndBodyFormat(offset int, limit int, Langcode_ string, BodySummary_ string, BodyFormat_ string) (*[]*NodeRevision_body, error) {
	var _NodeRevision_body = new([]*NodeRevision_body)
	err := Engine.Table("node_revision__body").Where("langcode = ? and body_summary = ? and body_format = ?", Langcode_, BodySummary_, BodyFormat_).Limit(limit, offset).Find(_NodeRevision_body)
	return _NodeRevision_body, err
}

// GetNodeRevision_bodiesByDeltaAndBodyValueAndBodySummary Get NodeRevision_bodies via DeltaAndBodyValueAndBodySummary
func GetNodeRevision_bodiesByDeltaAndBodyValueAndBodySummary(offset int, limit int, Delta_ int, BodyValue_ string, BodySummary_ string) (*[]*NodeRevision_body, error) {
	var _NodeRevision_body = new([]*NodeRevision_body)
	err := Engine.Table("node_revision__body").Where("delta = ? and body_value = ? and body_summary = ?", Delta_, BodyValue_, BodySummary_).Limit(limit, offset).Find(_NodeRevision_body)
	return _NodeRevision_body, err
}

// GetNodeRevision_bodiesByDeltaAndBodyValueAndBodyFormat Get NodeRevision_bodies via DeltaAndBodyValueAndBodyFormat
func GetNodeRevision_bodiesByDeltaAndBodyValueAndBodyFormat(offset int, limit int, Delta_ int, BodyValue_ string, BodyFormat_ string) (*[]*NodeRevision_body, error) {
	var _NodeRevision_body = new([]*NodeRevision_body)
	err := Engine.Table("node_revision__body").Where("delta = ? and body_value = ? and body_format = ?", Delta_, BodyValue_, BodyFormat_).Limit(limit, offset).Find(_NodeRevision_body)
	return _NodeRevision_body, err
}

// GetNodeRevision_bodiesByDeltaAndBodySummaryAndBodyFormat Get NodeRevision_bodies via DeltaAndBodySummaryAndBodyFormat
func GetNodeRevision_bodiesByDeltaAndBodySummaryAndBodyFormat(offset int, limit int, Delta_ int, BodySummary_ string, BodyFormat_ string) (*[]*NodeRevision_body, error) {
	var _NodeRevision_body = new([]*NodeRevision_body)
	err := Engine.Table("node_revision__body").Where("delta = ? and body_summary = ? and body_format = ?", Delta_, BodySummary_, BodyFormat_).Limit(limit, offset).Find(_NodeRevision_body)
	return _NodeRevision_body, err
}

// GetNodeRevision_bodiesByBodyValueAndBodySummaryAndBodyFormat Get NodeRevision_bodies via BodyValueAndBodySummaryAndBodyFormat
func GetNodeRevision_bodiesByBodyValueAndBodySummaryAndBodyFormat(offset int, limit int, BodyValue_ string, BodySummary_ string, BodyFormat_ string) (*[]*NodeRevision_body, error) {
	var _NodeRevision_body = new([]*NodeRevision_body)
	err := Engine.Table("node_revision__body").Where("body_value = ? and body_summary = ? and body_format = ?", BodyValue_, BodySummary_, BodyFormat_).Limit(limit, offset).Find(_NodeRevision_body)
	return _NodeRevision_body, err
}

// GetNodeRevision_bodiesByBundleAndDeleted Get NodeRevision_bodies via BundleAndDeleted
func GetNodeRevision_bodiesByBundleAndDeleted(offset int, limit int, Bundle_ string, Deleted_ int) (*[]*NodeRevision_body, error) {
	var _NodeRevision_body = new([]*NodeRevision_body)
	err := Engine.Table("node_revision__body").Where("bundle = ? and deleted = ?", Bundle_, Deleted_).Limit(limit, offset).Find(_NodeRevision_body)
	return _NodeRevision_body, err
}

// GetNodeRevision_bodiesByBundleAndEntityId Get NodeRevision_bodies via BundleAndEntityId
func GetNodeRevision_bodiesByBundleAndEntityId(offset int, limit int, Bundle_ string, EntityId_ int) (*[]*NodeRevision_body, error) {
	var _NodeRevision_body = new([]*NodeRevision_body)
	err := Engine.Table("node_revision__body").Where("bundle = ? and entity_id = ?", Bundle_, EntityId_).Limit(limit, offset).Find(_NodeRevision_body)
	return _NodeRevision_body, err
}

// GetNodeRevision_bodiesByBundleAndRevisionId Get NodeRevision_bodies via BundleAndRevisionId
func GetNodeRevision_bodiesByBundleAndRevisionId(offset int, limit int, Bundle_ string, RevisionId_ int) (*[]*NodeRevision_body, error) {
	var _NodeRevision_body = new([]*NodeRevision_body)
	err := Engine.Table("node_revision__body").Where("bundle = ? and revision_id = ?", Bundle_, RevisionId_).Limit(limit, offset).Find(_NodeRevision_body)
	return _NodeRevision_body, err
}

// GetNodeRevision_bodiesByBundleAndLangcode Get NodeRevision_bodies via BundleAndLangcode
func GetNodeRevision_bodiesByBundleAndLangcode(offset int, limit int, Bundle_ string, Langcode_ string) (*[]*NodeRevision_body, error) {
	var _NodeRevision_body = new([]*NodeRevision_body)
	err := Engine.Table("node_revision__body").Where("bundle = ? and langcode = ?", Bundle_, Langcode_).Limit(limit, offset).Find(_NodeRevision_body)
	return _NodeRevision_body, err
}

// GetNodeRevision_bodiesByBundleAndDelta Get NodeRevision_bodies via BundleAndDelta
func GetNodeRevision_bodiesByBundleAndDelta(offset int, limit int, Bundle_ string, Delta_ int) (*[]*NodeRevision_body, error) {
	var _NodeRevision_body = new([]*NodeRevision_body)
	err := Engine.Table("node_revision__body").Where("bundle = ? and delta = ?", Bundle_, Delta_).Limit(limit, offset).Find(_NodeRevision_body)
	return _NodeRevision_body, err
}

// GetNodeRevision_bodiesByBundleAndBodyValue Get NodeRevision_bodies via BundleAndBodyValue
func GetNodeRevision_bodiesByBundleAndBodyValue(offset int, limit int, Bundle_ string, BodyValue_ string) (*[]*NodeRevision_body, error) {
	var _NodeRevision_body = new([]*NodeRevision_body)
	err := Engine.Table("node_revision__body").Where("bundle = ? and body_value = ?", Bundle_, BodyValue_).Limit(limit, offset).Find(_NodeRevision_body)
	return _NodeRevision_body, err
}

// GetNodeRevision_bodiesByBundleAndBodySummary Get NodeRevision_bodies via BundleAndBodySummary
func GetNodeRevision_bodiesByBundleAndBodySummary(offset int, limit int, Bundle_ string, BodySummary_ string) (*[]*NodeRevision_body, error) {
	var _NodeRevision_body = new([]*NodeRevision_body)
	err := Engine.Table("node_revision__body").Where("bundle = ? and body_summary = ?", Bundle_, BodySummary_).Limit(limit, offset).Find(_NodeRevision_body)
	return _NodeRevision_body, err
}

// GetNodeRevision_bodiesByBundleAndBodyFormat Get NodeRevision_bodies via BundleAndBodyFormat
func GetNodeRevision_bodiesByBundleAndBodyFormat(offset int, limit int, Bundle_ string, BodyFormat_ string) (*[]*NodeRevision_body, error) {
	var _NodeRevision_body = new([]*NodeRevision_body)
	err := Engine.Table("node_revision__body").Where("bundle = ? and body_format = ?", Bundle_, BodyFormat_).Limit(limit, offset).Find(_NodeRevision_body)
	return _NodeRevision_body, err
}

// GetNodeRevision_bodiesByDeletedAndEntityId Get NodeRevision_bodies via DeletedAndEntityId
func GetNodeRevision_bodiesByDeletedAndEntityId(offset int, limit int, Deleted_ int, EntityId_ int) (*[]*NodeRevision_body, error) {
	var _NodeRevision_body = new([]*NodeRevision_body)
	err := Engine.Table("node_revision__body").Where("deleted = ? and entity_id = ?", Deleted_, EntityId_).Limit(limit, offset).Find(_NodeRevision_body)
	return _NodeRevision_body, err
}

// GetNodeRevision_bodiesByDeletedAndRevisionId Get NodeRevision_bodies via DeletedAndRevisionId
func GetNodeRevision_bodiesByDeletedAndRevisionId(offset int, limit int, Deleted_ int, RevisionId_ int) (*[]*NodeRevision_body, error) {
	var _NodeRevision_body = new([]*NodeRevision_body)
	err := Engine.Table("node_revision__body").Where("deleted = ? and revision_id = ?", Deleted_, RevisionId_).Limit(limit, offset).Find(_NodeRevision_body)
	return _NodeRevision_body, err
}

// GetNodeRevision_bodiesByDeletedAndLangcode Get NodeRevision_bodies via DeletedAndLangcode
func GetNodeRevision_bodiesByDeletedAndLangcode(offset int, limit int, Deleted_ int, Langcode_ string) (*[]*NodeRevision_body, error) {
	var _NodeRevision_body = new([]*NodeRevision_body)
	err := Engine.Table("node_revision__body").Where("deleted = ? and langcode = ?", Deleted_, Langcode_).Limit(limit, offset).Find(_NodeRevision_body)
	return _NodeRevision_body, err
}

// GetNodeRevision_bodiesByDeletedAndDelta Get NodeRevision_bodies via DeletedAndDelta
func GetNodeRevision_bodiesByDeletedAndDelta(offset int, limit int, Deleted_ int, Delta_ int) (*[]*NodeRevision_body, error) {
	var _NodeRevision_body = new([]*NodeRevision_body)
	err := Engine.Table("node_revision__body").Where("deleted = ? and delta = ?", Deleted_, Delta_).Limit(limit, offset).Find(_NodeRevision_body)
	return _NodeRevision_body, err
}

// GetNodeRevision_bodiesByDeletedAndBodyValue Get NodeRevision_bodies via DeletedAndBodyValue
func GetNodeRevision_bodiesByDeletedAndBodyValue(offset int, limit int, Deleted_ int, BodyValue_ string) (*[]*NodeRevision_body, error) {
	var _NodeRevision_body = new([]*NodeRevision_body)
	err := Engine.Table("node_revision__body").Where("deleted = ? and body_value = ?", Deleted_, BodyValue_).Limit(limit, offset).Find(_NodeRevision_body)
	return _NodeRevision_body, err
}

// GetNodeRevision_bodiesByDeletedAndBodySummary Get NodeRevision_bodies via DeletedAndBodySummary
func GetNodeRevision_bodiesByDeletedAndBodySummary(offset int, limit int, Deleted_ int, BodySummary_ string) (*[]*NodeRevision_body, error) {
	var _NodeRevision_body = new([]*NodeRevision_body)
	err := Engine.Table("node_revision__body").Where("deleted = ? and body_summary = ?", Deleted_, BodySummary_).Limit(limit, offset).Find(_NodeRevision_body)
	return _NodeRevision_body, err
}

// GetNodeRevision_bodiesByDeletedAndBodyFormat Get NodeRevision_bodies via DeletedAndBodyFormat
func GetNodeRevision_bodiesByDeletedAndBodyFormat(offset int, limit int, Deleted_ int, BodyFormat_ string) (*[]*NodeRevision_body, error) {
	var _NodeRevision_body = new([]*NodeRevision_body)
	err := Engine.Table("node_revision__body").Where("deleted = ? and body_format = ?", Deleted_, BodyFormat_).Limit(limit, offset).Find(_NodeRevision_body)
	return _NodeRevision_body, err
}

// GetNodeRevision_bodiesByEntityIdAndRevisionId Get NodeRevision_bodies via EntityIdAndRevisionId
func GetNodeRevision_bodiesByEntityIdAndRevisionId(offset int, limit int, EntityId_ int, RevisionId_ int) (*[]*NodeRevision_body, error) {
	var _NodeRevision_body = new([]*NodeRevision_body)
	err := Engine.Table("node_revision__body").Where("entity_id = ? and revision_id = ?", EntityId_, RevisionId_).Limit(limit, offset).Find(_NodeRevision_body)
	return _NodeRevision_body, err
}

// GetNodeRevision_bodiesByEntityIdAndLangcode Get NodeRevision_bodies via EntityIdAndLangcode
func GetNodeRevision_bodiesByEntityIdAndLangcode(offset int, limit int, EntityId_ int, Langcode_ string) (*[]*NodeRevision_body, error) {
	var _NodeRevision_body = new([]*NodeRevision_body)
	err := Engine.Table("node_revision__body").Where("entity_id = ? and langcode = ?", EntityId_, Langcode_).Limit(limit, offset).Find(_NodeRevision_body)
	return _NodeRevision_body, err
}

// GetNodeRevision_bodiesByEntityIdAndDelta Get NodeRevision_bodies via EntityIdAndDelta
func GetNodeRevision_bodiesByEntityIdAndDelta(offset int, limit int, EntityId_ int, Delta_ int) (*[]*NodeRevision_body, error) {
	var _NodeRevision_body = new([]*NodeRevision_body)
	err := Engine.Table("node_revision__body").Where("entity_id = ? and delta = ?", EntityId_, Delta_).Limit(limit, offset).Find(_NodeRevision_body)
	return _NodeRevision_body, err
}

// GetNodeRevision_bodiesByEntityIdAndBodyValue Get NodeRevision_bodies via EntityIdAndBodyValue
func GetNodeRevision_bodiesByEntityIdAndBodyValue(offset int, limit int, EntityId_ int, BodyValue_ string) (*[]*NodeRevision_body, error) {
	var _NodeRevision_body = new([]*NodeRevision_body)
	err := Engine.Table("node_revision__body").Where("entity_id = ? and body_value = ?", EntityId_, BodyValue_).Limit(limit, offset).Find(_NodeRevision_body)
	return _NodeRevision_body, err
}

// GetNodeRevision_bodiesByEntityIdAndBodySummary Get NodeRevision_bodies via EntityIdAndBodySummary
func GetNodeRevision_bodiesByEntityIdAndBodySummary(offset int, limit int, EntityId_ int, BodySummary_ string) (*[]*NodeRevision_body, error) {
	var _NodeRevision_body = new([]*NodeRevision_body)
	err := Engine.Table("node_revision__body").Where("entity_id = ? and body_summary = ?", EntityId_, BodySummary_).Limit(limit, offset).Find(_NodeRevision_body)
	return _NodeRevision_body, err
}

// GetNodeRevision_bodiesByEntityIdAndBodyFormat Get NodeRevision_bodies via EntityIdAndBodyFormat
func GetNodeRevision_bodiesByEntityIdAndBodyFormat(offset int, limit int, EntityId_ int, BodyFormat_ string) (*[]*NodeRevision_body, error) {
	var _NodeRevision_body = new([]*NodeRevision_body)
	err := Engine.Table("node_revision__body").Where("entity_id = ? and body_format = ?", EntityId_, BodyFormat_).Limit(limit, offset).Find(_NodeRevision_body)
	return _NodeRevision_body, err
}

// GetNodeRevision_bodiesByRevisionIdAndLangcode Get NodeRevision_bodies via RevisionIdAndLangcode
func GetNodeRevision_bodiesByRevisionIdAndLangcode(offset int, limit int, RevisionId_ int, Langcode_ string) (*[]*NodeRevision_body, error) {
	var _NodeRevision_body = new([]*NodeRevision_body)
	err := Engine.Table("node_revision__body").Where("revision_id = ? and langcode = ?", RevisionId_, Langcode_).Limit(limit, offset).Find(_NodeRevision_body)
	return _NodeRevision_body, err
}

// GetNodeRevision_bodiesByRevisionIdAndDelta Get NodeRevision_bodies via RevisionIdAndDelta
func GetNodeRevision_bodiesByRevisionIdAndDelta(offset int, limit int, RevisionId_ int, Delta_ int) (*[]*NodeRevision_body, error) {
	var _NodeRevision_body = new([]*NodeRevision_body)
	err := Engine.Table("node_revision__body").Where("revision_id = ? and delta = ?", RevisionId_, Delta_).Limit(limit, offset).Find(_NodeRevision_body)
	return _NodeRevision_body, err
}

// GetNodeRevision_bodiesByRevisionIdAndBodyValue Get NodeRevision_bodies via RevisionIdAndBodyValue
func GetNodeRevision_bodiesByRevisionIdAndBodyValue(offset int, limit int, RevisionId_ int, BodyValue_ string) (*[]*NodeRevision_body, error) {
	var _NodeRevision_body = new([]*NodeRevision_body)
	err := Engine.Table("node_revision__body").Where("revision_id = ? and body_value = ?", RevisionId_, BodyValue_).Limit(limit, offset).Find(_NodeRevision_body)
	return _NodeRevision_body, err
}

// GetNodeRevision_bodiesByRevisionIdAndBodySummary Get NodeRevision_bodies via RevisionIdAndBodySummary
func GetNodeRevision_bodiesByRevisionIdAndBodySummary(offset int, limit int, RevisionId_ int, BodySummary_ string) (*[]*NodeRevision_body, error) {
	var _NodeRevision_body = new([]*NodeRevision_body)
	err := Engine.Table("node_revision__body").Where("revision_id = ? and body_summary = ?", RevisionId_, BodySummary_).Limit(limit, offset).Find(_NodeRevision_body)
	return _NodeRevision_body, err
}

// GetNodeRevision_bodiesByRevisionIdAndBodyFormat Get NodeRevision_bodies via RevisionIdAndBodyFormat
func GetNodeRevision_bodiesByRevisionIdAndBodyFormat(offset int, limit int, RevisionId_ int, BodyFormat_ string) (*[]*NodeRevision_body, error) {
	var _NodeRevision_body = new([]*NodeRevision_body)
	err := Engine.Table("node_revision__body").Where("revision_id = ? and body_format = ?", RevisionId_, BodyFormat_).Limit(limit, offset).Find(_NodeRevision_body)
	return _NodeRevision_body, err
}

// GetNodeRevision_bodiesByLangcodeAndDelta Get NodeRevision_bodies via LangcodeAndDelta
func GetNodeRevision_bodiesByLangcodeAndDelta(offset int, limit int, Langcode_ string, Delta_ int) (*[]*NodeRevision_body, error) {
	var _NodeRevision_body = new([]*NodeRevision_body)
	err := Engine.Table("node_revision__body").Where("langcode = ? and delta = ?", Langcode_, Delta_).Limit(limit, offset).Find(_NodeRevision_body)
	return _NodeRevision_body, err
}

// GetNodeRevision_bodiesByLangcodeAndBodyValue Get NodeRevision_bodies via LangcodeAndBodyValue
func GetNodeRevision_bodiesByLangcodeAndBodyValue(offset int, limit int, Langcode_ string, BodyValue_ string) (*[]*NodeRevision_body, error) {
	var _NodeRevision_body = new([]*NodeRevision_body)
	err := Engine.Table("node_revision__body").Where("langcode = ? and body_value = ?", Langcode_, BodyValue_).Limit(limit, offset).Find(_NodeRevision_body)
	return _NodeRevision_body, err
}

// GetNodeRevision_bodiesByLangcodeAndBodySummary Get NodeRevision_bodies via LangcodeAndBodySummary
func GetNodeRevision_bodiesByLangcodeAndBodySummary(offset int, limit int, Langcode_ string, BodySummary_ string) (*[]*NodeRevision_body, error) {
	var _NodeRevision_body = new([]*NodeRevision_body)
	err := Engine.Table("node_revision__body").Where("langcode = ? and body_summary = ?", Langcode_, BodySummary_).Limit(limit, offset).Find(_NodeRevision_body)
	return _NodeRevision_body, err
}

// GetNodeRevision_bodiesByLangcodeAndBodyFormat Get NodeRevision_bodies via LangcodeAndBodyFormat
func GetNodeRevision_bodiesByLangcodeAndBodyFormat(offset int, limit int, Langcode_ string, BodyFormat_ string) (*[]*NodeRevision_body, error) {
	var _NodeRevision_body = new([]*NodeRevision_body)
	err := Engine.Table("node_revision__body").Where("langcode = ? and body_format = ?", Langcode_, BodyFormat_).Limit(limit, offset).Find(_NodeRevision_body)
	return _NodeRevision_body, err
}

// GetNodeRevision_bodiesByDeltaAndBodyValue Get NodeRevision_bodies via DeltaAndBodyValue
func GetNodeRevision_bodiesByDeltaAndBodyValue(offset int, limit int, Delta_ int, BodyValue_ string) (*[]*NodeRevision_body, error) {
	var _NodeRevision_body = new([]*NodeRevision_body)
	err := Engine.Table("node_revision__body").Where("delta = ? and body_value = ?", Delta_, BodyValue_).Limit(limit, offset).Find(_NodeRevision_body)
	return _NodeRevision_body, err
}

// GetNodeRevision_bodiesByDeltaAndBodySummary Get NodeRevision_bodies via DeltaAndBodySummary
func GetNodeRevision_bodiesByDeltaAndBodySummary(offset int, limit int, Delta_ int, BodySummary_ string) (*[]*NodeRevision_body, error) {
	var _NodeRevision_body = new([]*NodeRevision_body)
	err := Engine.Table("node_revision__body").Where("delta = ? and body_summary = ?", Delta_, BodySummary_).Limit(limit, offset).Find(_NodeRevision_body)
	return _NodeRevision_body, err
}

// GetNodeRevision_bodiesByDeltaAndBodyFormat Get NodeRevision_bodies via DeltaAndBodyFormat
func GetNodeRevision_bodiesByDeltaAndBodyFormat(offset int, limit int, Delta_ int, BodyFormat_ string) (*[]*NodeRevision_body, error) {
	var _NodeRevision_body = new([]*NodeRevision_body)
	err := Engine.Table("node_revision__body").Where("delta = ? and body_format = ?", Delta_, BodyFormat_).Limit(limit, offset).Find(_NodeRevision_body)
	return _NodeRevision_body, err
}

// GetNodeRevision_bodiesByBodyValueAndBodySummary Get NodeRevision_bodies via BodyValueAndBodySummary
func GetNodeRevision_bodiesByBodyValueAndBodySummary(offset int, limit int, BodyValue_ string, BodySummary_ string) (*[]*NodeRevision_body, error) {
	var _NodeRevision_body = new([]*NodeRevision_body)
	err := Engine.Table("node_revision__body").Where("body_value = ? and body_summary = ?", BodyValue_, BodySummary_).Limit(limit, offset).Find(_NodeRevision_body)
	return _NodeRevision_body, err
}

// GetNodeRevision_bodiesByBodyValueAndBodyFormat Get NodeRevision_bodies via BodyValueAndBodyFormat
func GetNodeRevision_bodiesByBodyValueAndBodyFormat(offset int, limit int, BodyValue_ string, BodyFormat_ string) (*[]*NodeRevision_body, error) {
	var _NodeRevision_body = new([]*NodeRevision_body)
	err := Engine.Table("node_revision__body").Where("body_value = ? and body_format = ?", BodyValue_, BodyFormat_).Limit(limit, offset).Find(_NodeRevision_body)
	return _NodeRevision_body, err
}

// GetNodeRevision_bodiesByBodySummaryAndBodyFormat Get NodeRevision_bodies via BodySummaryAndBodyFormat
func GetNodeRevision_bodiesByBodySummaryAndBodyFormat(offset int, limit int, BodySummary_ string, BodyFormat_ string) (*[]*NodeRevision_body, error) {
	var _NodeRevision_body = new([]*NodeRevision_body)
	err := Engine.Table("node_revision__body").Where("body_summary = ? and body_format = ?", BodySummary_, BodyFormat_).Limit(limit, offset).Find(_NodeRevision_body)
	return _NodeRevision_body, err
}

// GetNodeRevision_bodies Get NodeRevision_bodies via field
func GetNodeRevision_bodies(offset int, limit int, field string) (*[]*NodeRevision_body, error) {
	var _NodeRevision_body = new([]*NodeRevision_body)
	err := Engine.Table("node_revision__body").Limit(limit, offset).Desc(field).Find(_NodeRevision_body)
	return _NodeRevision_body, err
}

// GetNodeRevision_bodiesViaBundle Get NodeRevision_bodys via Bundle
func GetNodeRevision_bodiesViaBundle(offset int, limit int, Bundle_ string, field string) (*[]*NodeRevision_body, error) {
	var _NodeRevision_body = new([]*NodeRevision_body)
	err := Engine.Table("node_revision__body").Where("bundle = ?", Bundle_).Limit(limit, offset).Desc(field).Find(_NodeRevision_body)
	return _NodeRevision_body, err
}

// GetNodeRevision_bodiesViaDeleted Get NodeRevision_bodys via Deleted
func GetNodeRevision_bodiesViaDeleted(offset int, limit int, Deleted_ int, field string) (*[]*NodeRevision_body, error) {
	var _NodeRevision_body = new([]*NodeRevision_body)
	err := Engine.Table("node_revision__body").Where("deleted = ?", Deleted_).Limit(limit, offset).Desc(field).Find(_NodeRevision_body)
	return _NodeRevision_body, err
}

// GetNodeRevision_bodiesViaEntityId Get NodeRevision_bodys via EntityId
func GetNodeRevision_bodiesViaEntityId(offset int, limit int, EntityId_ int, field string) (*[]*NodeRevision_body, error) {
	var _NodeRevision_body = new([]*NodeRevision_body)
	err := Engine.Table("node_revision__body").Where("entity_id = ?", EntityId_).Limit(limit, offset).Desc(field).Find(_NodeRevision_body)
	return _NodeRevision_body, err
}

// GetNodeRevision_bodiesViaRevisionId Get NodeRevision_bodys via RevisionId
func GetNodeRevision_bodiesViaRevisionId(offset int, limit int, RevisionId_ int, field string) (*[]*NodeRevision_body, error) {
	var _NodeRevision_body = new([]*NodeRevision_body)
	err := Engine.Table("node_revision__body").Where("revision_id = ?", RevisionId_).Limit(limit, offset).Desc(field).Find(_NodeRevision_body)
	return _NodeRevision_body, err
}

// GetNodeRevision_bodiesViaLangcode Get NodeRevision_bodys via Langcode
func GetNodeRevision_bodiesViaLangcode(offset int, limit int, Langcode_ string, field string) (*[]*NodeRevision_body, error) {
	var _NodeRevision_body = new([]*NodeRevision_body)
	err := Engine.Table("node_revision__body").Where("langcode = ?", Langcode_).Limit(limit, offset).Desc(field).Find(_NodeRevision_body)
	return _NodeRevision_body, err
}

// GetNodeRevision_bodiesViaDelta Get NodeRevision_bodys via Delta
func GetNodeRevision_bodiesViaDelta(offset int, limit int, Delta_ int, field string) (*[]*NodeRevision_body, error) {
	var _NodeRevision_body = new([]*NodeRevision_body)
	err := Engine.Table("node_revision__body").Where("delta = ?", Delta_).Limit(limit, offset).Desc(field).Find(_NodeRevision_body)
	return _NodeRevision_body, err
}

// GetNodeRevision_bodiesViaBodyValue Get NodeRevision_bodys via BodyValue
func GetNodeRevision_bodiesViaBodyValue(offset int, limit int, BodyValue_ string, field string) (*[]*NodeRevision_body, error) {
	var _NodeRevision_body = new([]*NodeRevision_body)
	err := Engine.Table("node_revision__body").Where("body_value = ?", BodyValue_).Limit(limit, offset).Desc(field).Find(_NodeRevision_body)
	return _NodeRevision_body, err
}

// GetNodeRevision_bodiesViaBodySummary Get NodeRevision_bodys via BodySummary
func GetNodeRevision_bodiesViaBodySummary(offset int, limit int, BodySummary_ string, field string) (*[]*NodeRevision_body, error) {
	var _NodeRevision_body = new([]*NodeRevision_body)
	err := Engine.Table("node_revision__body").Where("body_summary = ?", BodySummary_).Limit(limit, offset).Desc(field).Find(_NodeRevision_body)
	return _NodeRevision_body, err
}

// GetNodeRevision_bodiesViaBodyFormat Get NodeRevision_bodys via BodyFormat
func GetNodeRevision_bodiesViaBodyFormat(offset int, limit int, BodyFormat_ string, field string) (*[]*NodeRevision_body, error) {
	var _NodeRevision_body = new([]*NodeRevision_body)
	err := Engine.Table("node_revision__body").Where("body_format = ?", BodyFormat_).Limit(limit, offset).Desc(field).Find(_NodeRevision_body)
	return _NodeRevision_body, err
}

// HasNodeRevision_bodyViaBundle Has NodeRevision_body via Bundle
func HasNodeRevision_bodyViaBundle(iBundle string) bool {
	if has, err := Engine.Where("bundle = ?", iBundle).Get(new(NodeRevision_body)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasNodeRevision_bodyViaDeleted Has NodeRevision_body via Deleted
func HasNodeRevision_bodyViaDeleted(iDeleted int) bool {
	if has, err := Engine.Where("deleted = ?", iDeleted).Get(new(NodeRevision_body)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasNodeRevision_bodyViaEntityId Has NodeRevision_body via EntityId
func HasNodeRevision_bodyViaEntityId(iEntityId int) bool {
	if has, err := Engine.Where("entity_id = ?", iEntityId).Get(new(NodeRevision_body)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasNodeRevision_bodyViaRevisionId Has NodeRevision_body via RevisionId
func HasNodeRevision_bodyViaRevisionId(iRevisionId int) bool {
	if has, err := Engine.Where("revision_id = ?", iRevisionId).Get(new(NodeRevision_body)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasNodeRevision_bodyViaLangcode Has NodeRevision_body via Langcode
func HasNodeRevision_bodyViaLangcode(iLangcode string) bool {
	if has, err := Engine.Where("langcode = ?", iLangcode).Get(new(NodeRevision_body)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasNodeRevision_bodyViaDelta Has NodeRevision_body via Delta
func HasNodeRevision_bodyViaDelta(iDelta int) bool {
	if has, err := Engine.Where("delta = ?", iDelta).Get(new(NodeRevision_body)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasNodeRevision_bodyViaBodyValue Has NodeRevision_body via BodyValue
func HasNodeRevision_bodyViaBodyValue(iBodyValue string) bool {
	if has, err := Engine.Where("body_value = ?", iBodyValue).Get(new(NodeRevision_body)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasNodeRevision_bodyViaBodySummary Has NodeRevision_body via BodySummary
func HasNodeRevision_bodyViaBodySummary(iBodySummary string) bool {
	if has, err := Engine.Where("body_summary = ?", iBodySummary).Get(new(NodeRevision_body)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasNodeRevision_bodyViaBodyFormat Has NodeRevision_body via BodyFormat
func HasNodeRevision_bodyViaBodyFormat(iBodyFormat string) bool {
	if has, err := Engine.Where("body_format = ?", iBodyFormat).Get(new(NodeRevision_body)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// GetNodeRevision_bodyViaBundle Get NodeRevision_body via Bundle
func GetNodeRevision_bodyViaBundle(iBundle string) (*NodeRevision_body, error) {
	var _NodeRevision_body = &NodeRevision_body{Bundle: iBundle}
	has, err := Engine.Get(_NodeRevision_body)
	if has {
		return _NodeRevision_body, err
	} else {
		return nil, err
	}
}

// GetNodeRevision_bodyViaDeleted Get NodeRevision_body via Deleted
func GetNodeRevision_bodyViaDeleted(iDeleted int) (*NodeRevision_body, error) {
	var _NodeRevision_body = &NodeRevision_body{Deleted: iDeleted}
	has, err := Engine.Get(_NodeRevision_body)
	if has {
		return _NodeRevision_body, err
	} else {
		return nil, err
	}
}

// GetNodeRevision_bodyViaEntityId Get NodeRevision_body via EntityId
func GetNodeRevision_bodyViaEntityId(iEntityId int) (*NodeRevision_body, error) {
	var _NodeRevision_body = &NodeRevision_body{EntityId: iEntityId}
	has, err := Engine.Get(_NodeRevision_body)
	if has {
		return _NodeRevision_body, err
	} else {
		return nil, err
	}
}

// GetNodeRevision_bodyViaRevisionId Get NodeRevision_body via RevisionId
func GetNodeRevision_bodyViaRevisionId(iRevisionId int) (*NodeRevision_body, error) {
	var _NodeRevision_body = &NodeRevision_body{RevisionId: iRevisionId}
	has, err := Engine.Get(_NodeRevision_body)
	if has {
		return _NodeRevision_body, err
	} else {
		return nil, err
	}
}

// GetNodeRevision_bodyViaLangcode Get NodeRevision_body via Langcode
func GetNodeRevision_bodyViaLangcode(iLangcode string) (*NodeRevision_body, error) {
	var _NodeRevision_body = &NodeRevision_body{Langcode: iLangcode}
	has, err := Engine.Get(_NodeRevision_body)
	if has {
		return _NodeRevision_body, err
	} else {
		return nil, err
	}
}

// GetNodeRevision_bodyViaDelta Get NodeRevision_body via Delta
func GetNodeRevision_bodyViaDelta(iDelta int) (*NodeRevision_body, error) {
	var _NodeRevision_body = &NodeRevision_body{Delta: iDelta}
	has, err := Engine.Get(_NodeRevision_body)
	if has {
		return _NodeRevision_body, err
	} else {
		return nil, err
	}
}

// GetNodeRevision_bodyViaBodyValue Get NodeRevision_body via BodyValue
func GetNodeRevision_bodyViaBodyValue(iBodyValue string) (*NodeRevision_body, error) {
	var _NodeRevision_body = &NodeRevision_body{BodyValue: iBodyValue}
	has, err := Engine.Get(_NodeRevision_body)
	if has {
		return _NodeRevision_body, err
	} else {
		return nil, err
	}
}

// GetNodeRevision_bodyViaBodySummary Get NodeRevision_body via BodySummary
func GetNodeRevision_bodyViaBodySummary(iBodySummary string) (*NodeRevision_body, error) {
	var _NodeRevision_body = &NodeRevision_body{BodySummary: iBodySummary}
	has, err := Engine.Get(_NodeRevision_body)
	if has {
		return _NodeRevision_body, err
	} else {
		return nil, err
	}
}

// GetNodeRevision_bodyViaBodyFormat Get NodeRevision_body via BodyFormat
func GetNodeRevision_bodyViaBodyFormat(iBodyFormat string) (*NodeRevision_body, error) {
	var _NodeRevision_body = &NodeRevision_body{BodyFormat: iBodyFormat}
	has, err := Engine.Get(_NodeRevision_body)
	if has {
		return _NodeRevision_body, err
	} else {
		return nil, err
	}
}

// SetNodeRevision_bodyViaBundle Set NodeRevision_body via Bundle
func SetNodeRevision_bodyViaBundle(iBundle string, node_revision__body *NodeRevision_body) (int64, error) {
	node_revision__body.Bundle = iBundle
	return Engine.Insert(node_revision__body)
}

// SetNodeRevision_bodyViaDeleted Set NodeRevision_body via Deleted
func SetNodeRevision_bodyViaDeleted(iDeleted int, node_revision__body *NodeRevision_body) (int64, error) {
	node_revision__body.Deleted = iDeleted
	return Engine.Insert(node_revision__body)
}

// SetNodeRevision_bodyViaEntityId Set NodeRevision_body via EntityId
func SetNodeRevision_bodyViaEntityId(iEntityId int, node_revision__body *NodeRevision_body) (int64, error) {
	node_revision__body.EntityId = iEntityId
	return Engine.Insert(node_revision__body)
}

// SetNodeRevision_bodyViaRevisionId Set NodeRevision_body via RevisionId
func SetNodeRevision_bodyViaRevisionId(iRevisionId int, node_revision__body *NodeRevision_body) (int64, error) {
	node_revision__body.RevisionId = iRevisionId
	return Engine.Insert(node_revision__body)
}

// SetNodeRevision_bodyViaLangcode Set NodeRevision_body via Langcode
func SetNodeRevision_bodyViaLangcode(iLangcode string, node_revision__body *NodeRevision_body) (int64, error) {
	node_revision__body.Langcode = iLangcode
	return Engine.Insert(node_revision__body)
}

// SetNodeRevision_bodyViaDelta Set NodeRevision_body via Delta
func SetNodeRevision_bodyViaDelta(iDelta int, node_revision__body *NodeRevision_body) (int64, error) {
	node_revision__body.Delta = iDelta
	return Engine.Insert(node_revision__body)
}

// SetNodeRevision_bodyViaBodyValue Set NodeRevision_body via BodyValue
func SetNodeRevision_bodyViaBodyValue(iBodyValue string, node_revision__body *NodeRevision_body) (int64, error) {
	node_revision__body.BodyValue = iBodyValue
	return Engine.Insert(node_revision__body)
}

// SetNodeRevision_bodyViaBodySummary Set NodeRevision_body via BodySummary
func SetNodeRevision_bodyViaBodySummary(iBodySummary string, node_revision__body *NodeRevision_body) (int64, error) {
	node_revision__body.BodySummary = iBodySummary
	return Engine.Insert(node_revision__body)
}

// SetNodeRevision_bodyViaBodyFormat Set NodeRevision_body via BodyFormat
func SetNodeRevision_bodyViaBodyFormat(iBodyFormat string, node_revision__body *NodeRevision_body) (int64, error) {
	node_revision__body.BodyFormat = iBodyFormat
	return Engine.Insert(node_revision__body)
}

// AddNodeRevision_body Add NodeRevision_body via all columns
func AddNodeRevision_body(iBundle string, iDeleted int, iEntityId int, iRevisionId int, iLangcode string, iDelta int, iBodyValue string, iBodySummary string, iBodyFormat string) error {
	_NodeRevision_body := &NodeRevision_body{Bundle: iBundle, Deleted: iDeleted, EntityId: iEntityId, RevisionId: iRevisionId, Langcode: iLangcode, Delta: iDelta, BodyValue: iBodyValue, BodySummary: iBodySummary, BodyFormat: iBodyFormat}
	if _, err := Engine.Insert(_NodeRevision_body); err != nil {
		return err
	} else {
		return nil
	}
}

// PostNodeRevision_body Post NodeRevision_body via iNodeRevision_body
func PostNodeRevision_body(iNodeRevision_body *NodeRevision_body) (string, error) {
	_, err := Engine.Insert(iNodeRevision_body)
	return iNodeRevision_body.Bundle, err
}

// PutNodeRevision_body Put NodeRevision_body
func PutNodeRevision_body(iNodeRevision_body *NodeRevision_body) (string, error) {
	_, err := Engine.Id(iNodeRevision_body.Bundle).Update(iNodeRevision_body)
	return iNodeRevision_body.Bundle, err
}

// PutNodeRevision_bodyViaBundle Put NodeRevision_body via Bundle
func PutNodeRevision_bodyViaBundle(Bundle_ string, iNodeRevision_body *NodeRevision_body) (int64, error) {
	row, err := Engine.Update(iNodeRevision_body, &NodeRevision_body{Bundle: Bundle_})
	return row, err
}

// PutNodeRevision_bodyViaDeleted Put NodeRevision_body via Deleted
func PutNodeRevision_bodyViaDeleted(Deleted_ int, iNodeRevision_body *NodeRevision_body) (int64, error) {
	row, err := Engine.Update(iNodeRevision_body, &NodeRevision_body{Deleted: Deleted_})
	return row, err
}

// PutNodeRevision_bodyViaEntityId Put NodeRevision_body via EntityId
func PutNodeRevision_bodyViaEntityId(EntityId_ int, iNodeRevision_body *NodeRevision_body) (int64, error) {
	row, err := Engine.Update(iNodeRevision_body, &NodeRevision_body{EntityId: EntityId_})
	return row, err
}

// PutNodeRevision_bodyViaRevisionId Put NodeRevision_body via RevisionId
func PutNodeRevision_bodyViaRevisionId(RevisionId_ int, iNodeRevision_body *NodeRevision_body) (int64, error) {
	row, err := Engine.Update(iNodeRevision_body, &NodeRevision_body{RevisionId: RevisionId_})
	return row, err
}

// PutNodeRevision_bodyViaLangcode Put NodeRevision_body via Langcode
func PutNodeRevision_bodyViaLangcode(Langcode_ string, iNodeRevision_body *NodeRevision_body) (int64, error) {
	row, err := Engine.Update(iNodeRevision_body, &NodeRevision_body{Langcode: Langcode_})
	return row, err
}

// PutNodeRevision_bodyViaDelta Put NodeRevision_body via Delta
func PutNodeRevision_bodyViaDelta(Delta_ int, iNodeRevision_body *NodeRevision_body) (int64, error) {
	row, err := Engine.Update(iNodeRevision_body, &NodeRevision_body{Delta: Delta_})
	return row, err
}

// PutNodeRevision_bodyViaBodyValue Put NodeRevision_body via BodyValue
func PutNodeRevision_bodyViaBodyValue(BodyValue_ string, iNodeRevision_body *NodeRevision_body) (int64, error) {
	row, err := Engine.Update(iNodeRevision_body, &NodeRevision_body{BodyValue: BodyValue_})
	return row, err
}

// PutNodeRevision_bodyViaBodySummary Put NodeRevision_body via BodySummary
func PutNodeRevision_bodyViaBodySummary(BodySummary_ string, iNodeRevision_body *NodeRevision_body) (int64, error) {
	row, err := Engine.Update(iNodeRevision_body, &NodeRevision_body{BodySummary: BodySummary_})
	return row, err
}

// PutNodeRevision_bodyViaBodyFormat Put NodeRevision_body via BodyFormat
func PutNodeRevision_bodyViaBodyFormat(BodyFormat_ string, iNodeRevision_body *NodeRevision_body) (int64, error) {
	row, err := Engine.Update(iNodeRevision_body, &NodeRevision_body{BodyFormat: BodyFormat_})
	return row, err
}

// UpdateNodeRevision_bodyViaBundle via map[string]interface{}{}
func UpdateNodeRevision_bodyViaBundle(iBundle string, iNodeRevision_bodyMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(NodeRevision_body)).Where("bundle = ?", iBundle).Update(iNodeRevision_bodyMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateNodeRevision_bodyViaDeleted via map[string]interface{}{}
func UpdateNodeRevision_bodyViaDeleted(iDeleted int, iNodeRevision_bodyMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(NodeRevision_body)).Where("deleted = ?", iDeleted).Update(iNodeRevision_bodyMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateNodeRevision_bodyViaEntityId via map[string]interface{}{}
func UpdateNodeRevision_bodyViaEntityId(iEntityId int, iNodeRevision_bodyMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(NodeRevision_body)).Where("entity_id = ?", iEntityId).Update(iNodeRevision_bodyMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateNodeRevision_bodyViaRevisionId via map[string]interface{}{}
func UpdateNodeRevision_bodyViaRevisionId(iRevisionId int, iNodeRevision_bodyMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(NodeRevision_body)).Where("revision_id = ?", iRevisionId).Update(iNodeRevision_bodyMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateNodeRevision_bodyViaLangcode via map[string]interface{}{}
func UpdateNodeRevision_bodyViaLangcode(iLangcode string, iNodeRevision_bodyMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(NodeRevision_body)).Where("langcode = ?", iLangcode).Update(iNodeRevision_bodyMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateNodeRevision_bodyViaDelta via map[string]interface{}{}
func UpdateNodeRevision_bodyViaDelta(iDelta int, iNodeRevision_bodyMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(NodeRevision_body)).Where("delta = ?", iDelta).Update(iNodeRevision_bodyMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateNodeRevision_bodyViaBodyValue via map[string]interface{}{}
func UpdateNodeRevision_bodyViaBodyValue(iBodyValue string, iNodeRevision_bodyMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(NodeRevision_body)).Where("body_value = ?", iBodyValue).Update(iNodeRevision_bodyMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateNodeRevision_bodyViaBodySummary via map[string]interface{}{}
func UpdateNodeRevision_bodyViaBodySummary(iBodySummary string, iNodeRevision_bodyMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(NodeRevision_body)).Where("body_summary = ?", iBodySummary).Update(iNodeRevision_bodyMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateNodeRevision_bodyViaBodyFormat via map[string]interface{}{}
func UpdateNodeRevision_bodyViaBodyFormat(iBodyFormat string, iNodeRevision_bodyMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(NodeRevision_body)).Where("body_format = ?", iBodyFormat).Update(iNodeRevision_bodyMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// DeleteNodeRevision_body Delete NodeRevision_body
func DeleteNodeRevision_body(iBundle string) (int64, error) {
	row, err := Engine.Id(iBundle).Delete(new(NodeRevision_body))
	return row, err
}

// DeleteNodeRevision_bodyViaBundle Delete NodeRevision_body via Bundle
func DeleteNodeRevision_bodyViaBundle(iBundle string) (err error) {
	var has bool
	var _NodeRevision_body = &NodeRevision_body{Bundle: iBundle}
	if has, err = Engine.Get(_NodeRevision_body); (has == true) && (err == nil) {
		if row, err := Engine.Where("bundle = ?", iBundle).Delete(new(NodeRevision_body)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteNodeRevision_bodyViaDeleted Delete NodeRevision_body via Deleted
func DeleteNodeRevision_bodyViaDeleted(iDeleted int) (err error) {
	var has bool
	var _NodeRevision_body = &NodeRevision_body{Deleted: iDeleted}
	if has, err = Engine.Get(_NodeRevision_body); (has == true) && (err == nil) {
		if row, err := Engine.Where("deleted = ?", iDeleted).Delete(new(NodeRevision_body)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteNodeRevision_bodyViaEntityId Delete NodeRevision_body via EntityId
func DeleteNodeRevision_bodyViaEntityId(iEntityId int) (err error) {
	var has bool
	var _NodeRevision_body = &NodeRevision_body{EntityId: iEntityId}
	if has, err = Engine.Get(_NodeRevision_body); (has == true) && (err == nil) {
		if row, err := Engine.Where("entity_id = ?", iEntityId).Delete(new(NodeRevision_body)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteNodeRevision_bodyViaRevisionId Delete NodeRevision_body via RevisionId
func DeleteNodeRevision_bodyViaRevisionId(iRevisionId int) (err error) {
	var has bool
	var _NodeRevision_body = &NodeRevision_body{RevisionId: iRevisionId}
	if has, err = Engine.Get(_NodeRevision_body); (has == true) && (err == nil) {
		if row, err := Engine.Where("revision_id = ?", iRevisionId).Delete(new(NodeRevision_body)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteNodeRevision_bodyViaLangcode Delete NodeRevision_body via Langcode
func DeleteNodeRevision_bodyViaLangcode(iLangcode string) (err error) {
	var has bool
	var _NodeRevision_body = &NodeRevision_body{Langcode: iLangcode}
	if has, err = Engine.Get(_NodeRevision_body); (has == true) && (err == nil) {
		if row, err := Engine.Where("langcode = ?", iLangcode).Delete(new(NodeRevision_body)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteNodeRevision_bodyViaDelta Delete NodeRevision_body via Delta
func DeleteNodeRevision_bodyViaDelta(iDelta int) (err error) {
	var has bool
	var _NodeRevision_body = &NodeRevision_body{Delta: iDelta}
	if has, err = Engine.Get(_NodeRevision_body); (has == true) && (err == nil) {
		if row, err := Engine.Where("delta = ?", iDelta).Delete(new(NodeRevision_body)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteNodeRevision_bodyViaBodyValue Delete NodeRevision_body via BodyValue
func DeleteNodeRevision_bodyViaBodyValue(iBodyValue string) (err error) {
	var has bool
	var _NodeRevision_body = &NodeRevision_body{BodyValue: iBodyValue}
	if has, err = Engine.Get(_NodeRevision_body); (has == true) && (err == nil) {
		if row, err := Engine.Where("body_value = ?", iBodyValue).Delete(new(NodeRevision_body)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteNodeRevision_bodyViaBodySummary Delete NodeRevision_body via BodySummary
func DeleteNodeRevision_bodyViaBodySummary(iBodySummary string) (err error) {
	var has bool
	var _NodeRevision_body = &NodeRevision_body{BodySummary: iBodySummary}
	if has, err = Engine.Get(_NodeRevision_body); (has == true) && (err == nil) {
		if row, err := Engine.Where("body_summary = ?", iBodySummary).Delete(new(NodeRevision_body)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteNodeRevision_bodyViaBodyFormat Delete NodeRevision_body via BodyFormat
func DeleteNodeRevision_bodyViaBodyFormat(iBodyFormat string) (err error) {
	var has bool
	var _NodeRevision_body = &NodeRevision_body{BodyFormat: iBodyFormat}
	if has, err = Engine.Get(_NodeRevision_body); (has == true) && (err == nil) {
		if row, err := Engine.Where("body_format = ?", iBodyFormat).Delete(new(NodeRevision_body)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}
