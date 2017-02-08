package model

type BlockContentRevision_body struct {
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

// GetBlockContentRevision_bodiesCount BlockContentRevision_bodys' Count
func GetBlockContentRevision_bodiesCount(offset int, limit int) (int64, error) {
	total, err := Engine.Limit(limit, offset).Count(&BlockContentRevision_body{})
	return total, err
}

// GetBlockContentRevision_bodyCountViaBundle Get BlockContentRevision_body via Bundle
func GetBlockContentRevision_bodyCountViaBundle(iBundle string) int64 {
	n, _ := Engine.Where("bundle = ?", iBundle).Count(&BlockContentRevision_body{Bundle: iBundle})
	return n
}

// GetBlockContentRevision_bodyCountViaDeleted Get BlockContentRevision_body via Deleted
func GetBlockContentRevision_bodyCountViaDeleted(iDeleted int) int64 {
	n, _ := Engine.Where("deleted = ?", iDeleted).Count(&BlockContentRevision_body{Deleted: iDeleted})
	return n
}

// GetBlockContentRevision_bodyCountViaEntityId Get BlockContentRevision_body via EntityId
func GetBlockContentRevision_bodyCountViaEntityId(iEntityId int) int64 {
	n, _ := Engine.Where("entity_id = ?", iEntityId).Count(&BlockContentRevision_body{EntityId: iEntityId})
	return n
}

// GetBlockContentRevision_bodyCountViaRevisionId Get BlockContentRevision_body via RevisionId
func GetBlockContentRevision_bodyCountViaRevisionId(iRevisionId int) int64 {
	n, _ := Engine.Where("revision_id = ?", iRevisionId).Count(&BlockContentRevision_body{RevisionId: iRevisionId})
	return n
}

// GetBlockContentRevision_bodyCountViaLangcode Get BlockContentRevision_body via Langcode
func GetBlockContentRevision_bodyCountViaLangcode(iLangcode string) int64 {
	n, _ := Engine.Where("langcode = ?", iLangcode).Count(&BlockContentRevision_body{Langcode: iLangcode})
	return n
}

// GetBlockContentRevision_bodyCountViaDelta Get BlockContentRevision_body via Delta
func GetBlockContentRevision_bodyCountViaDelta(iDelta int) int64 {
	n, _ := Engine.Where("delta = ?", iDelta).Count(&BlockContentRevision_body{Delta: iDelta})
	return n
}

// GetBlockContentRevision_bodyCountViaBodyValue Get BlockContentRevision_body via BodyValue
func GetBlockContentRevision_bodyCountViaBodyValue(iBodyValue string) int64 {
	n, _ := Engine.Where("body_value = ?", iBodyValue).Count(&BlockContentRevision_body{BodyValue: iBodyValue})
	return n
}

// GetBlockContentRevision_bodyCountViaBodySummary Get BlockContentRevision_body via BodySummary
func GetBlockContentRevision_bodyCountViaBodySummary(iBodySummary string) int64 {
	n, _ := Engine.Where("body_summary = ?", iBodySummary).Count(&BlockContentRevision_body{BodySummary: iBodySummary})
	return n
}

// GetBlockContentRevision_bodyCountViaBodyFormat Get BlockContentRevision_body via BodyFormat
func GetBlockContentRevision_bodyCountViaBodyFormat(iBodyFormat string) int64 {
	n, _ := Engine.Where("body_format = ?", iBodyFormat).Count(&BlockContentRevision_body{BodyFormat: iBodyFormat})
	return n
}

// GetBlockContentRevision_bodiesByBundleAndDeletedAndEntityId Get BlockContentRevision_bodies via BundleAndDeletedAndEntityId
func GetBlockContentRevision_bodiesByBundleAndDeletedAndEntityId(offset int, limit int, Bundle_ string, Deleted_ int, EntityId_ int) (*[]*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = new([]*BlockContentRevision_body)
	err := Engine.Table("block_content_revision__body").Where("bundle = ? and deleted = ? and entity_id = ?", Bundle_, Deleted_, EntityId_).Limit(limit, offset).Find(_BlockContentRevision_body)
	return _BlockContentRevision_body, err
}

// GetBlockContentRevision_bodiesByBundleAndDeletedAndRevisionId Get BlockContentRevision_bodies via BundleAndDeletedAndRevisionId
func GetBlockContentRevision_bodiesByBundleAndDeletedAndRevisionId(offset int, limit int, Bundle_ string, Deleted_ int, RevisionId_ int) (*[]*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = new([]*BlockContentRevision_body)
	err := Engine.Table("block_content_revision__body").Where("bundle = ? and deleted = ? and revision_id = ?", Bundle_, Deleted_, RevisionId_).Limit(limit, offset).Find(_BlockContentRevision_body)
	return _BlockContentRevision_body, err
}

// GetBlockContentRevision_bodiesByBundleAndDeletedAndLangcode Get BlockContentRevision_bodies via BundleAndDeletedAndLangcode
func GetBlockContentRevision_bodiesByBundleAndDeletedAndLangcode(offset int, limit int, Bundle_ string, Deleted_ int, Langcode_ string) (*[]*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = new([]*BlockContentRevision_body)
	err := Engine.Table("block_content_revision__body").Where("bundle = ? and deleted = ? and langcode = ?", Bundle_, Deleted_, Langcode_).Limit(limit, offset).Find(_BlockContentRevision_body)
	return _BlockContentRevision_body, err
}

// GetBlockContentRevision_bodiesByBundleAndDeletedAndDelta Get BlockContentRevision_bodies via BundleAndDeletedAndDelta
func GetBlockContentRevision_bodiesByBundleAndDeletedAndDelta(offset int, limit int, Bundle_ string, Deleted_ int, Delta_ int) (*[]*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = new([]*BlockContentRevision_body)
	err := Engine.Table("block_content_revision__body").Where("bundle = ? and deleted = ? and delta = ?", Bundle_, Deleted_, Delta_).Limit(limit, offset).Find(_BlockContentRevision_body)
	return _BlockContentRevision_body, err
}

// GetBlockContentRevision_bodiesByBundleAndDeletedAndBodyValue Get BlockContentRevision_bodies via BundleAndDeletedAndBodyValue
func GetBlockContentRevision_bodiesByBundleAndDeletedAndBodyValue(offset int, limit int, Bundle_ string, Deleted_ int, BodyValue_ string) (*[]*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = new([]*BlockContentRevision_body)
	err := Engine.Table("block_content_revision__body").Where("bundle = ? and deleted = ? and body_value = ?", Bundle_, Deleted_, BodyValue_).Limit(limit, offset).Find(_BlockContentRevision_body)
	return _BlockContentRevision_body, err
}

// GetBlockContentRevision_bodiesByBundleAndDeletedAndBodySummary Get BlockContentRevision_bodies via BundleAndDeletedAndBodySummary
func GetBlockContentRevision_bodiesByBundleAndDeletedAndBodySummary(offset int, limit int, Bundle_ string, Deleted_ int, BodySummary_ string) (*[]*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = new([]*BlockContentRevision_body)
	err := Engine.Table("block_content_revision__body").Where("bundle = ? and deleted = ? and body_summary = ?", Bundle_, Deleted_, BodySummary_).Limit(limit, offset).Find(_BlockContentRevision_body)
	return _BlockContentRevision_body, err
}

// GetBlockContentRevision_bodiesByBundleAndDeletedAndBodyFormat Get BlockContentRevision_bodies via BundleAndDeletedAndBodyFormat
func GetBlockContentRevision_bodiesByBundleAndDeletedAndBodyFormat(offset int, limit int, Bundle_ string, Deleted_ int, BodyFormat_ string) (*[]*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = new([]*BlockContentRevision_body)
	err := Engine.Table("block_content_revision__body").Where("bundle = ? and deleted = ? and body_format = ?", Bundle_, Deleted_, BodyFormat_).Limit(limit, offset).Find(_BlockContentRevision_body)
	return _BlockContentRevision_body, err
}

// GetBlockContentRevision_bodiesByBundleAndEntityIdAndRevisionId Get BlockContentRevision_bodies via BundleAndEntityIdAndRevisionId
func GetBlockContentRevision_bodiesByBundleAndEntityIdAndRevisionId(offset int, limit int, Bundle_ string, EntityId_ int, RevisionId_ int) (*[]*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = new([]*BlockContentRevision_body)
	err := Engine.Table("block_content_revision__body").Where("bundle = ? and entity_id = ? and revision_id = ?", Bundle_, EntityId_, RevisionId_).Limit(limit, offset).Find(_BlockContentRevision_body)
	return _BlockContentRevision_body, err
}

// GetBlockContentRevision_bodiesByBundleAndEntityIdAndLangcode Get BlockContentRevision_bodies via BundleAndEntityIdAndLangcode
func GetBlockContentRevision_bodiesByBundleAndEntityIdAndLangcode(offset int, limit int, Bundle_ string, EntityId_ int, Langcode_ string) (*[]*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = new([]*BlockContentRevision_body)
	err := Engine.Table("block_content_revision__body").Where("bundle = ? and entity_id = ? and langcode = ?", Bundle_, EntityId_, Langcode_).Limit(limit, offset).Find(_BlockContentRevision_body)
	return _BlockContentRevision_body, err
}

// GetBlockContentRevision_bodiesByBundleAndEntityIdAndDelta Get BlockContentRevision_bodies via BundleAndEntityIdAndDelta
func GetBlockContentRevision_bodiesByBundleAndEntityIdAndDelta(offset int, limit int, Bundle_ string, EntityId_ int, Delta_ int) (*[]*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = new([]*BlockContentRevision_body)
	err := Engine.Table("block_content_revision__body").Where("bundle = ? and entity_id = ? and delta = ?", Bundle_, EntityId_, Delta_).Limit(limit, offset).Find(_BlockContentRevision_body)
	return _BlockContentRevision_body, err
}

// GetBlockContentRevision_bodiesByBundleAndEntityIdAndBodyValue Get BlockContentRevision_bodies via BundleAndEntityIdAndBodyValue
func GetBlockContentRevision_bodiesByBundleAndEntityIdAndBodyValue(offset int, limit int, Bundle_ string, EntityId_ int, BodyValue_ string) (*[]*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = new([]*BlockContentRevision_body)
	err := Engine.Table("block_content_revision__body").Where("bundle = ? and entity_id = ? and body_value = ?", Bundle_, EntityId_, BodyValue_).Limit(limit, offset).Find(_BlockContentRevision_body)
	return _BlockContentRevision_body, err
}

// GetBlockContentRevision_bodiesByBundleAndEntityIdAndBodySummary Get BlockContentRevision_bodies via BundleAndEntityIdAndBodySummary
func GetBlockContentRevision_bodiesByBundleAndEntityIdAndBodySummary(offset int, limit int, Bundle_ string, EntityId_ int, BodySummary_ string) (*[]*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = new([]*BlockContentRevision_body)
	err := Engine.Table("block_content_revision__body").Where("bundle = ? and entity_id = ? and body_summary = ?", Bundle_, EntityId_, BodySummary_).Limit(limit, offset).Find(_BlockContentRevision_body)
	return _BlockContentRevision_body, err
}

// GetBlockContentRevision_bodiesByBundleAndEntityIdAndBodyFormat Get BlockContentRevision_bodies via BundleAndEntityIdAndBodyFormat
func GetBlockContentRevision_bodiesByBundleAndEntityIdAndBodyFormat(offset int, limit int, Bundle_ string, EntityId_ int, BodyFormat_ string) (*[]*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = new([]*BlockContentRevision_body)
	err := Engine.Table("block_content_revision__body").Where("bundle = ? and entity_id = ? and body_format = ?", Bundle_, EntityId_, BodyFormat_).Limit(limit, offset).Find(_BlockContentRevision_body)
	return _BlockContentRevision_body, err
}

// GetBlockContentRevision_bodiesByBundleAndRevisionIdAndLangcode Get BlockContentRevision_bodies via BundleAndRevisionIdAndLangcode
func GetBlockContentRevision_bodiesByBundleAndRevisionIdAndLangcode(offset int, limit int, Bundle_ string, RevisionId_ int, Langcode_ string) (*[]*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = new([]*BlockContentRevision_body)
	err := Engine.Table("block_content_revision__body").Where("bundle = ? and revision_id = ? and langcode = ?", Bundle_, RevisionId_, Langcode_).Limit(limit, offset).Find(_BlockContentRevision_body)
	return _BlockContentRevision_body, err
}

// GetBlockContentRevision_bodiesByBundleAndRevisionIdAndDelta Get BlockContentRevision_bodies via BundleAndRevisionIdAndDelta
func GetBlockContentRevision_bodiesByBundleAndRevisionIdAndDelta(offset int, limit int, Bundle_ string, RevisionId_ int, Delta_ int) (*[]*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = new([]*BlockContentRevision_body)
	err := Engine.Table("block_content_revision__body").Where("bundle = ? and revision_id = ? and delta = ?", Bundle_, RevisionId_, Delta_).Limit(limit, offset).Find(_BlockContentRevision_body)
	return _BlockContentRevision_body, err
}

// GetBlockContentRevision_bodiesByBundleAndRevisionIdAndBodyValue Get BlockContentRevision_bodies via BundleAndRevisionIdAndBodyValue
func GetBlockContentRevision_bodiesByBundleAndRevisionIdAndBodyValue(offset int, limit int, Bundle_ string, RevisionId_ int, BodyValue_ string) (*[]*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = new([]*BlockContentRevision_body)
	err := Engine.Table("block_content_revision__body").Where("bundle = ? and revision_id = ? and body_value = ?", Bundle_, RevisionId_, BodyValue_).Limit(limit, offset).Find(_BlockContentRevision_body)
	return _BlockContentRevision_body, err
}

// GetBlockContentRevision_bodiesByBundleAndRevisionIdAndBodySummary Get BlockContentRevision_bodies via BundleAndRevisionIdAndBodySummary
func GetBlockContentRevision_bodiesByBundleAndRevisionIdAndBodySummary(offset int, limit int, Bundle_ string, RevisionId_ int, BodySummary_ string) (*[]*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = new([]*BlockContentRevision_body)
	err := Engine.Table("block_content_revision__body").Where("bundle = ? and revision_id = ? and body_summary = ?", Bundle_, RevisionId_, BodySummary_).Limit(limit, offset).Find(_BlockContentRevision_body)
	return _BlockContentRevision_body, err
}

// GetBlockContentRevision_bodiesByBundleAndRevisionIdAndBodyFormat Get BlockContentRevision_bodies via BundleAndRevisionIdAndBodyFormat
func GetBlockContentRevision_bodiesByBundleAndRevisionIdAndBodyFormat(offset int, limit int, Bundle_ string, RevisionId_ int, BodyFormat_ string) (*[]*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = new([]*BlockContentRevision_body)
	err := Engine.Table("block_content_revision__body").Where("bundle = ? and revision_id = ? and body_format = ?", Bundle_, RevisionId_, BodyFormat_).Limit(limit, offset).Find(_BlockContentRevision_body)
	return _BlockContentRevision_body, err
}

// GetBlockContentRevision_bodiesByBundleAndLangcodeAndDelta Get BlockContentRevision_bodies via BundleAndLangcodeAndDelta
func GetBlockContentRevision_bodiesByBundleAndLangcodeAndDelta(offset int, limit int, Bundle_ string, Langcode_ string, Delta_ int) (*[]*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = new([]*BlockContentRevision_body)
	err := Engine.Table("block_content_revision__body").Where("bundle = ? and langcode = ? and delta = ?", Bundle_, Langcode_, Delta_).Limit(limit, offset).Find(_BlockContentRevision_body)
	return _BlockContentRevision_body, err
}

// GetBlockContentRevision_bodiesByBundleAndLangcodeAndBodyValue Get BlockContentRevision_bodies via BundleAndLangcodeAndBodyValue
func GetBlockContentRevision_bodiesByBundleAndLangcodeAndBodyValue(offset int, limit int, Bundle_ string, Langcode_ string, BodyValue_ string) (*[]*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = new([]*BlockContentRevision_body)
	err := Engine.Table("block_content_revision__body").Where("bundle = ? and langcode = ? and body_value = ?", Bundle_, Langcode_, BodyValue_).Limit(limit, offset).Find(_BlockContentRevision_body)
	return _BlockContentRevision_body, err
}

// GetBlockContentRevision_bodiesByBundleAndLangcodeAndBodySummary Get BlockContentRevision_bodies via BundleAndLangcodeAndBodySummary
func GetBlockContentRevision_bodiesByBundleAndLangcodeAndBodySummary(offset int, limit int, Bundle_ string, Langcode_ string, BodySummary_ string) (*[]*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = new([]*BlockContentRevision_body)
	err := Engine.Table("block_content_revision__body").Where("bundle = ? and langcode = ? and body_summary = ?", Bundle_, Langcode_, BodySummary_).Limit(limit, offset).Find(_BlockContentRevision_body)
	return _BlockContentRevision_body, err
}

// GetBlockContentRevision_bodiesByBundleAndLangcodeAndBodyFormat Get BlockContentRevision_bodies via BundleAndLangcodeAndBodyFormat
func GetBlockContentRevision_bodiesByBundleAndLangcodeAndBodyFormat(offset int, limit int, Bundle_ string, Langcode_ string, BodyFormat_ string) (*[]*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = new([]*BlockContentRevision_body)
	err := Engine.Table("block_content_revision__body").Where("bundle = ? and langcode = ? and body_format = ?", Bundle_, Langcode_, BodyFormat_).Limit(limit, offset).Find(_BlockContentRevision_body)
	return _BlockContentRevision_body, err
}

// GetBlockContentRevision_bodiesByBundleAndDeltaAndBodyValue Get BlockContentRevision_bodies via BundleAndDeltaAndBodyValue
func GetBlockContentRevision_bodiesByBundleAndDeltaAndBodyValue(offset int, limit int, Bundle_ string, Delta_ int, BodyValue_ string) (*[]*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = new([]*BlockContentRevision_body)
	err := Engine.Table("block_content_revision__body").Where("bundle = ? and delta = ? and body_value = ?", Bundle_, Delta_, BodyValue_).Limit(limit, offset).Find(_BlockContentRevision_body)
	return _BlockContentRevision_body, err
}

// GetBlockContentRevision_bodiesByBundleAndDeltaAndBodySummary Get BlockContentRevision_bodies via BundleAndDeltaAndBodySummary
func GetBlockContentRevision_bodiesByBundleAndDeltaAndBodySummary(offset int, limit int, Bundle_ string, Delta_ int, BodySummary_ string) (*[]*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = new([]*BlockContentRevision_body)
	err := Engine.Table("block_content_revision__body").Where("bundle = ? and delta = ? and body_summary = ?", Bundle_, Delta_, BodySummary_).Limit(limit, offset).Find(_BlockContentRevision_body)
	return _BlockContentRevision_body, err
}

// GetBlockContentRevision_bodiesByBundleAndDeltaAndBodyFormat Get BlockContentRevision_bodies via BundleAndDeltaAndBodyFormat
func GetBlockContentRevision_bodiesByBundleAndDeltaAndBodyFormat(offset int, limit int, Bundle_ string, Delta_ int, BodyFormat_ string) (*[]*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = new([]*BlockContentRevision_body)
	err := Engine.Table("block_content_revision__body").Where("bundle = ? and delta = ? and body_format = ?", Bundle_, Delta_, BodyFormat_).Limit(limit, offset).Find(_BlockContentRevision_body)
	return _BlockContentRevision_body, err
}

// GetBlockContentRevision_bodiesByBundleAndBodyValueAndBodySummary Get BlockContentRevision_bodies via BundleAndBodyValueAndBodySummary
func GetBlockContentRevision_bodiesByBundleAndBodyValueAndBodySummary(offset int, limit int, Bundle_ string, BodyValue_ string, BodySummary_ string) (*[]*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = new([]*BlockContentRevision_body)
	err := Engine.Table("block_content_revision__body").Where("bundle = ? and body_value = ? and body_summary = ?", Bundle_, BodyValue_, BodySummary_).Limit(limit, offset).Find(_BlockContentRevision_body)
	return _BlockContentRevision_body, err
}

// GetBlockContentRevision_bodiesByBundleAndBodyValueAndBodyFormat Get BlockContentRevision_bodies via BundleAndBodyValueAndBodyFormat
func GetBlockContentRevision_bodiesByBundleAndBodyValueAndBodyFormat(offset int, limit int, Bundle_ string, BodyValue_ string, BodyFormat_ string) (*[]*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = new([]*BlockContentRevision_body)
	err := Engine.Table("block_content_revision__body").Where("bundle = ? and body_value = ? and body_format = ?", Bundle_, BodyValue_, BodyFormat_).Limit(limit, offset).Find(_BlockContentRevision_body)
	return _BlockContentRevision_body, err
}

// GetBlockContentRevision_bodiesByBundleAndBodySummaryAndBodyFormat Get BlockContentRevision_bodies via BundleAndBodySummaryAndBodyFormat
func GetBlockContentRevision_bodiesByBundleAndBodySummaryAndBodyFormat(offset int, limit int, Bundle_ string, BodySummary_ string, BodyFormat_ string) (*[]*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = new([]*BlockContentRevision_body)
	err := Engine.Table("block_content_revision__body").Where("bundle = ? and body_summary = ? and body_format = ?", Bundle_, BodySummary_, BodyFormat_).Limit(limit, offset).Find(_BlockContentRevision_body)
	return _BlockContentRevision_body, err
}

// GetBlockContentRevision_bodiesByDeletedAndEntityIdAndRevisionId Get BlockContentRevision_bodies via DeletedAndEntityIdAndRevisionId
func GetBlockContentRevision_bodiesByDeletedAndEntityIdAndRevisionId(offset int, limit int, Deleted_ int, EntityId_ int, RevisionId_ int) (*[]*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = new([]*BlockContentRevision_body)
	err := Engine.Table("block_content_revision__body").Where("deleted = ? and entity_id = ? and revision_id = ?", Deleted_, EntityId_, RevisionId_).Limit(limit, offset).Find(_BlockContentRevision_body)
	return _BlockContentRevision_body, err
}

// GetBlockContentRevision_bodiesByDeletedAndEntityIdAndLangcode Get BlockContentRevision_bodies via DeletedAndEntityIdAndLangcode
func GetBlockContentRevision_bodiesByDeletedAndEntityIdAndLangcode(offset int, limit int, Deleted_ int, EntityId_ int, Langcode_ string) (*[]*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = new([]*BlockContentRevision_body)
	err := Engine.Table("block_content_revision__body").Where("deleted = ? and entity_id = ? and langcode = ?", Deleted_, EntityId_, Langcode_).Limit(limit, offset).Find(_BlockContentRevision_body)
	return _BlockContentRevision_body, err
}

// GetBlockContentRevision_bodiesByDeletedAndEntityIdAndDelta Get BlockContentRevision_bodies via DeletedAndEntityIdAndDelta
func GetBlockContentRevision_bodiesByDeletedAndEntityIdAndDelta(offset int, limit int, Deleted_ int, EntityId_ int, Delta_ int) (*[]*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = new([]*BlockContentRevision_body)
	err := Engine.Table("block_content_revision__body").Where("deleted = ? and entity_id = ? and delta = ?", Deleted_, EntityId_, Delta_).Limit(limit, offset).Find(_BlockContentRevision_body)
	return _BlockContentRevision_body, err
}

// GetBlockContentRevision_bodiesByDeletedAndEntityIdAndBodyValue Get BlockContentRevision_bodies via DeletedAndEntityIdAndBodyValue
func GetBlockContentRevision_bodiesByDeletedAndEntityIdAndBodyValue(offset int, limit int, Deleted_ int, EntityId_ int, BodyValue_ string) (*[]*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = new([]*BlockContentRevision_body)
	err := Engine.Table("block_content_revision__body").Where("deleted = ? and entity_id = ? and body_value = ?", Deleted_, EntityId_, BodyValue_).Limit(limit, offset).Find(_BlockContentRevision_body)
	return _BlockContentRevision_body, err
}

// GetBlockContentRevision_bodiesByDeletedAndEntityIdAndBodySummary Get BlockContentRevision_bodies via DeletedAndEntityIdAndBodySummary
func GetBlockContentRevision_bodiesByDeletedAndEntityIdAndBodySummary(offset int, limit int, Deleted_ int, EntityId_ int, BodySummary_ string) (*[]*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = new([]*BlockContentRevision_body)
	err := Engine.Table("block_content_revision__body").Where("deleted = ? and entity_id = ? and body_summary = ?", Deleted_, EntityId_, BodySummary_).Limit(limit, offset).Find(_BlockContentRevision_body)
	return _BlockContentRevision_body, err
}

// GetBlockContentRevision_bodiesByDeletedAndEntityIdAndBodyFormat Get BlockContentRevision_bodies via DeletedAndEntityIdAndBodyFormat
func GetBlockContentRevision_bodiesByDeletedAndEntityIdAndBodyFormat(offset int, limit int, Deleted_ int, EntityId_ int, BodyFormat_ string) (*[]*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = new([]*BlockContentRevision_body)
	err := Engine.Table("block_content_revision__body").Where("deleted = ? and entity_id = ? and body_format = ?", Deleted_, EntityId_, BodyFormat_).Limit(limit, offset).Find(_BlockContentRevision_body)
	return _BlockContentRevision_body, err
}

// GetBlockContentRevision_bodiesByDeletedAndRevisionIdAndLangcode Get BlockContentRevision_bodies via DeletedAndRevisionIdAndLangcode
func GetBlockContentRevision_bodiesByDeletedAndRevisionIdAndLangcode(offset int, limit int, Deleted_ int, RevisionId_ int, Langcode_ string) (*[]*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = new([]*BlockContentRevision_body)
	err := Engine.Table("block_content_revision__body").Where("deleted = ? and revision_id = ? and langcode = ?", Deleted_, RevisionId_, Langcode_).Limit(limit, offset).Find(_BlockContentRevision_body)
	return _BlockContentRevision_body, err
}

// GetBlockContentRevision_bodiesByDeletedAndRevisionIdAndDelta Get BlockContentRevision_bodies via DeletedAndRevisionIdAndDelta
func GetBlockContentRevision_bodiesByDeletedAndRevisionIdAndDelta(offset int, limit int, Deleted_ int, RevisionId_ int, Delta_ int) (*[]*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = new([]*BlockContentRevision_body)
	err := Engine.Table("block_content_revision__body").Where("deleted = ? and revision_id = ? and delta = ?", Deleted_, RevisionId_, Delta_).Limit(limit, offset).Find(_BlockContentRevision_body)
	return _BlockContentRevision_body, err
}

// GetBlockContentRevision_bodiesByDeletedAndRevisionIdAndBodyValue Get BlockContentRevision_bodies via DeletedAndRevisionIdAndBodyValue
func GetBlockContentRevision_bodiesByDeletedAndRevisionIdAndBodyValue(offset int, limit int, Deleted_ int, RevisionId_ int, BodyValue_ string) (*[]*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = new([]*BlockContentRevision_body)
	err := Engine.Table("block_content_revision__body").Where("deleted = ? and revision_id = ? and body_value = ?", Deleted_, RevisionId_, BodyValue_).Limit(limit, offset).Find(_BlockContentRevision_body)
	return _BlockContentRevision_body, err
}

// GetBlockContentRevision_bodiesByDeletedAndRevisionIdAndBodySummary Get BlockContentRevision_bodies via DeletedAndRevisionIdAndBodySummary
func GetBlockContentRevision_bodiesByDeletedAndRevisionIdAndBodySummary(offset int, limit int, Deleted_ int, RevisionId_ int, BodySummary_ string) (*[]*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = new([]*BlockContentRevision_body)
	err := Engine.Table("block_content_revision__body").Where("deleted = ? and revision_id = ? and body_summary = ?", Deleted_, RevisionId_, BodySummary_).Limit(limit, offset).Find(_BlockContentRevision_body)
	return _BlockContentRevision_body, err
}

// GetBlockContentRevision_bodiesByDeletedAndRevisionIdAndBodyFormat Get BlockContentRevision_bodies via DeletedAndRevisionIdAndBodyFormat
func GetBlockContentRevision_bodiesByDeletedAndRevisionIdAndBodyFormat(offset int, limit int, Deleted_ int, RevisionId_ int, BodyFormat_ string) (*[]*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = new([]*BlockContentRevision_body)
	err := Engine.Table("block_content_revision__body").Where("deleted = ? and revision_id = ? and body_format = ?", Deleted_, RevisionId_, BodyFormat_).Limit(limit, offset).Find(_BlockContentRevision_body)
	return _BlockContentRevision_body, err
}

// GetBlockContentRevision_bodiesByDeletedAndLangcodeAndDelta Get BlockContentRevision_bodies via DeletedAndLangcodeAndDelta
func GetBlockContentRevision_bodiesByDeletedAndLangcodeAndDelta(offset int, limit int, Deleted_ int, Langcode_ string, Delta_ int) (*[]*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = new([]*BlockContentRevision_body)
	err := Engine.Table("block_content_revision__body").Where("deleted = ? and langcode = ? and delta = ?", Deleted_, Langcode_, Delta_).Limit(limit, offset).Find(_BlockContentRevision_body)
	return _BlockContentRevision_body, err
}

// GetBlockContentRevision_bodiesByDeletedAndLangcodeAndBodyValue Get BlockContentRevision_bodies via DeletedAndLangcodeAndBodyValue
func GetBlockContentRevision_bodiesByDeletedAndLangcodeAndBodyValue(offset int, limit int, Deleted_ int, Langcode_ string, BodyValue_ string) (*[]*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = new([]*BlockContentRevision_body)
	err := Engine.Table("block_content_revision__body").Where("deleted = ? and langcode = ? and body_value = ?", Deleted_, Langcode_, BodyValue_).Limit(limit, offset).Find(_BlockContentRevision_body)
	return _BlockContentRevision_body, err
}

// GetBlockContentRevision_bodiesByDeletedAndLangcodeAndBodySummary Get BlockContentRevision_bodies via DeletedAndLangcodeAndBodySummary
func GetBlockContentRevision_bodiesByDeletedAndLangcodeAndBodySummary(offset int, limit int, Deleted_ int, Langcode_ string, BodySummary_ string) (*[]*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = new([]*BlockContentRevision_body)
	err := Engine.Table("block_content_revision__body").Where("deleted = ? and langcode = ? and body_summary = ?", Deleted_, Langcode_, BodySummary_).Limit(limit, offset).Find(_BlockContentRevision_body)
	return _BlockContentRevision_body, err
}

// GetBlockContentRevision_bodiesByDeletedAndLangcodeAndBodyFormat Get BlockContentRevision_bodies via DeletedAndLangcodeAndBodyFormat
func GetBlockContentRevision_bodiesByDeletedAndLangcodeAndBodyFormat(offset int, limit int, Deleted_ int, Langcode_ string, BodyFormat_ string) (*[]*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = new([]*BlockContentRevision_body)
	err := Engine.Table("block_content_revision__body").Where("deleted = ? and langcode = ? and body_format = ?", Deleted_, Langcode_, BodyFormat_).Limit(limit, offset).Find(_BlockContentRevision_body)
	return _BlockContentRevision_body, err
}

// GetBlockContentRevision_bodiesByDeletedAndDeltaAndBodyValue Get BlockContentRevision_bodies via DeletedAndDeltaAndBodyValue
func GetBlockContentRevision_bodiesByDeletedAndDeltaAndBodyValue(offset int, limit int, Deleted_ int, Delta_ int, BodyValue_ string) (*[]*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = new([]*BlockContentRevision_body)
	err := Engine.Table("block_content_revision__body").Where("deleted = ? and delta = ? and body_value = ?", Deleted_, Delta_, BodyValue_).Limit(limit, offset).Find(_BlockContentRevision_body)
	return _BlockContentRevision_body, err
}

// GetBlockContentRevision_bodiesByDeletedAndDeltaAndBodySummary Get BlockContentRevision_bodies via DeletedAndDeltaAndBodySummary
func GetBlockContentRevision_bodiesByDeletedAndDeltaAndBodySummary(offset int, limit int, Deleted_ int, Delta_ int, BodySummary_ string) (*[]*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = new([]*BlockContentRevision_body)
	err := Engine.Table("block_content_revision__body").Where("deleted = ? and delta = ? and body_summary = ?", Deleted_, Delta_, BodySummary_).Limit(limit, offset).Find(_BlockContentRevision_body)
	return _BlockContentRevision_body, err
}

// GetBlockContentRevision_bodiesByDeletedAndDeltaAndBodyFormat Get BlockContentRevision_bodies via DeletedAndDeltaAndBodyFormat
func GetBlockContentRevision_bodiesByDeletedAndDeltaAndBodyFormat(offset int, limit int, Deleted_ int, Delta_ int, BodyFormat_ string) (*[]*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = new([]*BlockContentRevision_body)
	err := Engine.Table("block_content_revision__body").Where("deleted = ? and delta = ? and body_format = ?", Deleted_, Delta_, BodyFormat_).Limit(limit, offset).Find(_BlockContentRevision_body)
	return _BlockContentRevision_body, err
}

// GetBlockContentRevision_bodiesByDeletedAndBodyValueAndBodySummary Get BlockContentRevision_bodies via DeletedAndBodyValueAndBodySummary
func GetBlockContentRevision_bodiesByDeletedAndBodyValueAndBodySummary(offset int, limit int, Deleted_ int, BodyValue_ string, BodySummary_ string) (*[]*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = new([]*BlockContentRevision_body)
	err := Engine.Table("block_content_revision__body").Where("deleted = ? and body_value = ? and body_summary = ?", Deleted_, BodyValue_, BodySummary_).Limit(limit, offset).Find(_BlockContentRevision_body)
	return _BlockContentRevision_body, err
}

// GetBlockContentRevision_bodiesByDeletedAndBodyValueAndBodyFormat Get BlockContentRevision_bodies via DeletedAndBodyValueAndBodyFormat
func GetBlockContentRevision_bodiesByDeletedAndBodyValueAndBodyFormat(offset int, limit int, Deleted_ int, BodyValue_ string, BodyFormat_ string) (*[]*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = new([]*BlockContentRevision_body)
	err := Engine.Table("block_content_revision__body").Where("deleted = ? and body_value = ? and body_format = ?", Deleted_, BodyValue_, BodyFormat_).Limit(limit, offset).Find(_BlockContentRevision_body)
	return _BlockContentRevision_body, err
}

// GetBlockContentRevision_bodiesByDeletedAndBodySummaryAndBodyFormat Get BlockContentRevision_bodies via DeletedAndBodySummaryAndBodyFormat
func GetBlockContentRevision_bodiesByDeletedAndBodySummaryAndBodyFormat(offset int, limit int, Deleted_ int, BodySummary_ string, BodyFormat_ string) (*[]*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = new([]*BlockContentRevision_body)
	err := Engine.Table("block_content_revision__body").Where("deleted = ? and body_summary = ? and body_format = ?", Deleted_, BodySummary_, BodyFormat_).Limit(limit, offset).Find(_BlockContentRevision_body)
	return _BlockContentRevision_body, err
}

// GetBlockContentRevision_bodiesByEntityIdAndRevisionIdAndLangcode Get BlockContentRevision_bodies via EntityIdAndRevisionIdAndLangcode
func GetBlockContentRevision_bodiesByEntityIdAndRevisionIdAndLangcode(offset int, limit int, EntityId_ int, RevisionId_ int, Langcode_ string) (*[]*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = new([]*BlockContentRevision_body)
	err := Engine.Table("block_content_revision__body").Where("entity_id = ? and revision_id = ? and langcode = ?", EntityId_, RevisionId_, Langcode_).Limit(limit, offset).Find(_BlockContentRevision_body)
	return _BlockContentRevision_body, err
}

// GetBlockContentRevision_bodiesByEntityIdAndRevisionIdAndDelta Get BlockContentRevision_bodies via EntityIdAndRevisionIdAndDelta
func GetBlockContentRevision_bodiesByEntityIdAndRevisionIdAndDelta(offset int, limit int, EntityId_ int, RevisionId_ int, Delta_ int) (*[]*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = new([]*BlockContentRevision_body)
	err := Engine.Table("block_content_revision__body").Where("entity_id = ? and revision_id = ? and delta = ?", EntityId_, RevisionId_, Delta_).Limit(limit, offset).Find(_BlockContentRevision_body)
	return _BlockContentRevision_body, err
}

// GetBlockContentRevision_bodiesByEntityIdAndRevisionIdAndBodyValue Get BlockContentRevision_bodies via EntityIdAndRevisionIdAndBodyValue
func GetBlockContentRevision_bodiesByEntityIdAndRevisionIdAndBodyValue(offset int, limit int, EntityId_ int, RevisionId_ int, BodyValue_ string) (*[]*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = new([]*BlockContentRevision_body)
	err := Engine.Table("block_content_revision__body").Where("entity_id = ? and revision_id = ? and body_value = ?", EntityId_, RevisionId_, BodyValue_).Limit(limit, offset).Find(_BlockContentRevision_body)
	return _BlockContentRevision_body, err
}

// GetBlockContentRevision_bodiesByEntityIdAndRevisionIdAndBodySummary Get BlockContentRevision_bodies via EntityIdAndRevisionIdAndBodySummary
func GetBlockContentRevision_bodiesByEntityIdAndRevisionIdAndBodySummary(offset int, limit int, EntityId_ int, RevisionId_ int, BodySummary_ string) (*[]*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = new([]*BlockContentRevision_body)
	err := Engine.Table("block_content_revision__body").Where("entity_id = ? and revision_id = ? and body_summary = ?", EntityId_, RevisionId_, BodySummary_).Limit(limit, offset).Find(_BlockContentRevision_body)
	return _BlockContentRevision_body, err
}

// GetBlockContentRevision_bodiesByEntityIdAndRevisionIdAndBodyFormat Get BlockContentRevision_bodies via EntityIdAndRevisionIdAndBodyFormat
func GetBlockContentRevision_bodiesByEntityIdAndRevisionIdAndBodyFormat(offset int, limit int, EntityId_ int, RevisionId_ int, BodyFormat_ string) (*[]*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = new([]*BlockContentRevision_body)
	err := Engine.Table("block_content_revision__body").Where("entity_id = ? and revision_id = ? and body_format = ?", EntityId_, RevisionId_, BodyFormat_).Limit(limit, offset).Find(_BlockContentRevision_body)
	return _BlockContentRevision_body, err
}

// GetBlockContentRevision_bodiesByEntityIdAndLangcodeAndDelta Get BlockContentRevision_bodies via EntityIdAndLangcodeAndDelta
func GetBlockContentRevision_bodiesByEntityIdAndLangcodeAndDelta(offset int, limit int, EntityId_ int, Langcode_ string, Delta_ int) (*[]*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = new([]*BlockContentRevision_body)
	err := Engine.Table("block_content_revision__body").Where("entity_id = ? and langcode = ? and delta = ?", EntityId_, Langcode_, Delta_).Limit(limit, offset).Find(_BlockContentRevision_body)
	return _BlockContentRevision_body, err
}

// GetBlockContentRevision_bodiesByEntityIdAndLangcodeAndBodyValue Get BlockContentRevision_bodies via EntityIdAndLangcodeAndBodyValue
func GetBlockContentRevision_bodiesByEntityIdAndLangcodeAndBodyValue(offset int, limit int, EntityId_ int, Langcode_ string, BodyValue_ string) (*[]*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = new([]*BlockContentRevision_body)
	err := Engine.Table("block_content_revision__body").Where("entity_id = ? and langcode = ? and body_value = ?", EntityId_, Langcode_, BodyValue_).Limit(limit, offset).Find(_BlockContentRevision_body)
	return _BlockContentRevision_body, err
}

// GetBlockContentRevision_bodiesByEntityIdAndLangcodeAndBodySummary Get BlockContentRevision_bodies via EntityIdAndLangcodeAndBodySummary
func GetBlockContentRevision_bodiesByEntityIdAndLangcodeAndBodySummary(offset int, limit int, EntityId_ int, Langcode_ string, BodySummary_ string) (*[]*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = new([]*BlockContentRevision_body)
	err := Engine.Table("block_content_revision__body").Where("entity_id = ? and langcode = ? and body_summary = ?", EntityId_, Langcode_, BodySummary_).Limit(limit, offset).Find(_BlockContentRevision_body)
	return _BlockContentRevision_body, err
}

// GetBlockContentRevision_bodiesByEntityIdAndLangcodeAndBodyFormat Get BlockContentRevision_bodies via EntityIdAndLangcodeAndBodyFormat
func GetBlockContentRevision_bodiesByEntityIdAndLangcodeAndBodyFormat(offset int, limit int, EntityId_ int, Langcode_ string, BodyFormat_ string) (*[]*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = new([]*BlockContentRevision_body)
	err := Engine.Table("block_content_revision__body").Where("entity_id = ? and langcode = ? and body_format = ?", EntityId_, Langcode_, BodyFormat_).Limit(limit, offset).Find(_BlockContentRevision_body)
	return _BlockContentRevision_body, err
}

// GetBlockContentRevision_bodiesByEntityIdAndDeltaAndBodyValue Get BlockContentRevision_bodies via EntityIdAndDeltaAndBodyValue
func GetBlockContentRevision_bodiesByEntityIdAndDeltaAndBodyValue(offset int, limit int, EntityId_ int, Delta_ int, BodyValue_ string) (*[]*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = new([]*BlockContentRevision_body)
	err := Engine.Table("block_content_revision__body").Where("entity_id = ? and delta = ? and body_value = ?", EntityId_, Delta_, BodyValue_).Limit(limit, offset).Find(_BlockContentRevision_body)
	return _BlockContentRevision_body, err
}

// GetBlockContentRevision_bodiesByEntityIdAndDeltaAndBodySummary Get BlockContentRevision_bodies via EntityIdAndDeltaAndBodySummary
func GetBlockContentRevision_bodiesByEntityIdAndDeltaAndBodySummary(offset int, limit int, EntityId_ int, Delta_ int, BodySummary_ string) (*[]*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = new([]*BlockContentRevision_body)
	err := Engine.Table("block_content_revision__body").Where("entity_id = ? and delta = ? and body_summary = ?", EntityId_, Delta_, BodySummary_).Limit(limit, offset).Find(_BlockContentRevision_body)
	return _BlockContentRevision_body, err
}

// GetBlockContentRevision_bodiesByEntityIdAndDeltaAndBodyFormat Get BlockContentRevision_bodies via EntityIdAndDeltaAndBodyFormat
func GetBlockContentRevision_bodiesByEntityIdAndDeltaAndBodyFormat(offset int, limit int, EntityId_ int, Delta_ int, BodyFormat_ string) (*[]*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = new([]*BlockContentRevision_body)
	err := Engine.Table("block_content_revision__body").Where("entity_id = ? and delta = ? and body_format = ?", EntityId_, Delta_, BodyFormat_).Limit(limit, offset).Find(_BlockContentRevision_body)
	return _BlockContentRevision_body, err
}

// GetBlockContentRevision_bodiesByEntityIdAndBodyValueAndBodySummary Get BlockContentRevision_bodies via EntityIdAndBodyValueAndBodySummary
func GetBlockContentRevision_bodiesByEntityIdAndBodyValueAndBodySummary(offset int, limit int, EntityId_ int, BodyValue_ string, BodySummary_ string) (*[]*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = new([]*BlockContentRevision_body)
	err := Engine.Table("block_content_revision__body").Where("entity_id = ? and body_value = ? and body_summary = ?", EntityId_, BodyValue_, BodySummary_).Limit(limit, offset).Find(_BlockContentRevision_body)
	return _BlockContentRevision_body, err
}

// GetBlockContentRevision_bodiesByEntityIdAndBodyValueAndBodyFormat Get BlockContentRevision_bodies via EntityIdAndBodyValueAndBodyFormat
func GetBlockContentRevision_bodiesByEntityIdAndBodyValueAndBodyFormat(offset int, limit int, EntityId_ int, BodyValue_ string, BodyFormat_ string) (*[]*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = new([]*BlockContentRevision_body)
	err := Engine.Table("block_content_revision__body").Where("entity_id = ? and body_value = ? and body_format = ?", EntityId_, BodyValue_, BodyFormat_).Limit(limit, offset).Find(_BlockContentRevision_body)
	return _BlockContentRevision_body, err
}

// GetBlockContentRevision_bodiesByEntityIdAndBodySummaryAndBodyFormat Get BlockContentRevision_bodies via EntityIdAndBodySummaryAndBodyFormat
func GetBlockContentRevision_bodiesByEntityIdAndBodySummaryAndBodyFormat(offset int, limit int, EntityId_ int, BodySummary_ string, BodyFormat_ string) (*[]*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = new([]*BlockContentRevision_body)
	err := Engine.Table("block_content_revision__body").Where("entity_id = ? and body_summary = ? and body_format = ?", EntityId_, BodySummary_, BodyFormat_).Limit(limit, offset).Find(_BlockContentRevision_body)
	return _BlockContentRevision_body, err
}

// GetBlockContentRevision_bodiesByRevisionIdAndLangcodeAndDelta Get BlockContentRevision_bodies via RevisionIdAndLangcodeAndDelta
func GetBlockContentRevision_bodiesByRevisionIdAndLangcodeAndDelta(offset int, limit int, RevisionId_ int, Langcode_ string, Delta_ int) (*[]*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = new([]*BlockContentRevision_body)
	err := Engine.Table("block_content_revision__body").Where("revision_id = ? and langcode = ? and delta = ?", RevisionId_, Langcode_, Delta_).Limit(limit, offset).Find(_BlockContentRevision_body)
	return _BlockContentRevision_body, err
}

// GetBlockContentRevision_bodiesByRevisionIdAndLangcodeAndBodyValue Get BlockContentRevision_bodies via RevisionIdAndLangcodeAndBodyValue
func GetBlockContentRevision_bodiesByRevisionIdAndLangcodeAndBodyValue(offset int, limit int, RevisionId_ int, Langcode_ string, BodyValue_ string) (*[]*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = new([]*BlockContentRevision_body)
	err := Engine.Table("block_content_revision__body").Where("revision_id = ? and langcode = ? and body_value = ?", RevisionId_, Langcode_, BodyValue_).Limit(limit, offset).Find(_BlockContentRevision_body)
	return _BlockContentRevision_body, err
}

// GetBlockContentRevision_bodiesByRevisionIdAndLangcodeAndBodySummary Get BlockContentRevision_bodies via RevisionIdAndLangcodeAndBodySummary
func GetBlockContentRevision_bodiesByRevisionIdAndLangcodeAndBodySummary(offset int, limit int, RevisionId_ int, Langcode_ string, BodySummary_ string) (*[]*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = new([]*BlockContentRevision_body)
	err := Engine.Table("block_content_revision__body").Where("revision_id = ? and langcode = ? and body_summary = ?", RevisionId_, Langcode_, BodySummary_).Limit(limit, offset).Find(_BlockContentRevision_body)
	return _BlockContentRevision_body, err
}

// GetBlockContentRevision_bodiesByRevisionIdAndLangcodeAndBodyFormat Get BlockContentRevision_bodies via RevisionIdAndLangcodeAndBodyFormat
func GetBlockContentRevision_bodiesByRevisionIdAndLangcodeAndBodyFormat(offset int, limit int, RevisionId_ int, Langcode_ string, BodyFormat_ string) (*[]*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = new([]*BlockContentRevision_body)
	err := Engine.Table("block_content_revision__body").Where("revision_id = ? and langcode = ? and body_format = ?", RevisionId_, Langcode_, BodyFormat_).Limit(limit, offset).Find(_BlockContentRevision_body)
	return _BlockContentRevision_body, err
}

// GetBlockContentRevision_bodiesByRevisionIdAndDeltaAndBodyValue Get BlockContentRevision_bodies via RevisionIdAndDeltaAndBodyValue
func GetBlockContentRevision_bodiesByRevisionIdAndDeltaAndBodyValue(offset int, limit int, RevisionId_ int, Delta_ int, BodyValue_ string) (*[]*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = new([]*BlockContentRevision_body)
	err := Engine.Table("block_content_revision__body").Where("revision_id = ? and delta = ? and body_value = ?", RevisionId_, Delta_, BodyValue_).Limit(limit, offset).Find(_BlockContentRevision_body)
	return _BlockContentRevision_body, err
}

// GetBlockContentRevision_bodiesByRevisionIdAndDeltaAndBodySummary Get BlockContentRevision_bodies via RevisionIdAndDeltaAndBodySummary
func GetBlockContentRevision_bodiesByRevisionIdAndDeltaAndBodySummary(offset int, limit int, RevisionId_ int, Delta_ int, BodySummary_ string) (*[]*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = new([]*BlockContentRevision_body)
	err := Engine.Table("block_content_revision__body").Where("revision_id = ? and delta = ? and body_summary = ?", RevisionId_, Delta_, BodySummary_).Limit(limit, offset).Find(_BlockContentRevision_body)
	return _BlockContentRevision_body, err
}

// GetBlockContentRevision_bodiesByRevisionIdAndDeltaAndBodyFormat Get BlockContentRevision_bodies via RevisionIdAndDeltaAndBodyFormat
func GetBlockContentRevision_bodiesByRevisionIdAndDeltaAndBodyFormat(offset int, limit int, RevisionId_ int, Delta_ int, BodyFormat_ string) (*[]*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = new([]*BlockContentRevision_body)
	err := Engine.Table("block_content_revision__body").Where("revision_id = ? and delta = ? and body_format = ?", RevisionId_, Delta_, BodyFormat_).Limit(limit, offset).Find(_BlockContentRevision_body)
	return _BlockContentRevision_body, err
}

// GetBlockContentRevision_bodiesByRevisionIdAndBodyValueAndBodySummary Get BlockContentRevision_bodies via RevisionIdAndBodyValueAndBodySummary
func GetBlockContentRevision_bodiesByRevisionIdAndBodyValueAndBodySummary(offset int, limit int, RevisionId_ int, BodyValue_ string, BodySummary_ string) (*[]*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = new([]*BlockContentRevision_body)
	err := Engine.Table("block_content_revision__body").Where("revision_id = ? and body_value = ? and body_summary = ?", RevisionId_, BodyValue_, BodySummary_).Limit(limit, offset).Find(_BlockContentRevision_body)
	return _BlockContentRevision_body, err
}

// GetBlockContentRevision_bodiesByRevisionIdAndBodyValueAndBodyFormat Get BlockContentRevision_bodies via RevisionIdAndBodyValueAndBodyFormat
func GetBlockContentRevision_bodiesByRevisionIdAndBodyValueAndBodyFormat(offset int, limit int, RevisionId_ int, BodyValue_ string, BodyFormat_ string) (*[]*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = new([]*BlockContentRevision_body)
	err := Engine.Table("block_content_revision__body").Where("revision_id = ? and body_value = ? and body_format = ?", RevisionId_, BodyValue_, BodyFormat_).Limit(limit, offset).Find(_BlockContentRevision_body)
	return _BlockContentRevision_body, err
}

// GetBlockContentRevision_bodiesByRevisionIdAndBodySummaryAndBodyFormat Get BlockContentRevision_bodies via RevisionIdAndBodySummaryAndBodyFormat
func GetBlockContentRevision_bodiesByRevisionIdAndBodySummaryAndBodyFormat(offset int, limit int, RevisionId_ int, BodySummary_ string, BodyFormat_ string) (*[]*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = new([]*BlockContentRevision_body)
	err := Engine.Table("block_content_revision__body").Where("revision_id = ? and body_summary = ? and body_format = ?", RevisionId_, BodySummary_, BodyFormat_).Limit(limit, offset).Find(_BlockContentRevision_body)
	return _BlockContentRevision_body, err
}

// GetBlockContentRevision_bodiesByLangcodeAndDeltaAndBodyValue Get BlockContentRevision_bodies via LangcodeAndDeltaAndBodyValue
func GetBlockContentRevision_bodiesByLangcodeAndDeltaAndBodyValue(offset int, limit int, Langcode_ string, Delta_ int, BodyValue_ string) (*[]*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = new([]*BlockContentRevision_body)
	err := Engine.Table("block_content_revision__body").Where("langcode = ? and delta = ? and body_value = ?", Langcode_, Delta_, BodyValue_).Limit(limit, offset).Find(_BlockContentRevision_body)
	return _BlockContentRevision_body, err
}

// GetBlockContentRevision_bodiesByLangcodeAndDeltaAndBodySummary Get BlockContentRevision_bodies via LangcodeAndDeltaAndBodySummary
func GetBlockContentRevision_bodiesByLangcodeAndDeltaAndBodySummary(offset int, limit int, Langcode_ string, Delta_ int, BodySummary_ string) (*[]*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = new([]*BlockContentRevision_body)
	err := Engine.Table("block_content_revision__body").Where("langcode = ? and delta = ? and body_summary = ?", Langcode_, Delta_, BodySummary_).Limit(limit, offset).Find(_BlockContentRevision_body)
	return _BlockContentRevision_body, err
}

// GetBlockContentRevision_bodiesByLangcodeAndDeltaAndBodyFormat Get BlockContentRevision_bodies via LangcodeAndDeltaAndBodyFormat
func GetBlockContentRevision_bodiesByLangcodeAndDeltaAndBodyFormat(offset int, limit int, Langcode_ string, Delta_ int, BodyFormat_ string) (*[]*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = new([]*BlockContentRevision_body)
	err := Engine.Table("block_content_revision__body").Where("langcode = ? and delta = ? and body_format = ?", Langcode_, Delta_, BodyFormat_).Limit(limit, offset).Find(_BlockContentRevision_body)
	return _BlockContentRevision_body, err
}

// GetBlockContentRevision_bodiesByLangcodeAndBodyValueAndBodySummary Get BlockContentRevision_bodies via LangcodeAndBodyValueAndBodySummary
func GetBlockContentRevision_bodiesByLangcodeAndBodyValueAndBodySummary(offset int, limit int, Langcode_ string, BodyValue_ string, BodySummary_ string) (*[]*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = new([]*BlockContentRevision_body)
	err := Engine.Table("block_content_revision__body").Where("langcode = ? and body_value = ? and body_summary = ?", Langcode_, BodyValue_, BodySummary_).Limit(limit, offset).Find(_BlockContentRevision_body)
	return _BlockContentRevision_body, err
}

// GetBlockContentRevision_bodiesByLangcodeAndBodyValueAndBodyFormat Get BlockContentRevision_bodies via LangcodeAndBodyValueAndBodyFormat
func GetBlockContentRevision_bodiesByLangcodeAndBodyValueAndBodyFormat(offset int, limit int, Langcode_ string, BodyValue_ string, BodyFormat_ string) (*[]*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = new([]*BlockContentRevision_body)
	err := Engine.Table("block_content_revision__body").Where("langcode = ? and body_value = ? and body_format = ?", Langcode_, BodyValue_, BodyFormat_).Limit(limit, offset).Find(_BlockContentRevision_body)
	return _BlockContentRevision_body, err
}

// GetBlockContentRevision_bodiesByLangcodeAndBodySummaryAndBodyFormat Get BlockContentRevision_bodies via LangcodeAndBodySummaryAndBodyFormat
func GetBlockContentRevision_bodiesByLangcodeAndBodySummaryAndBodyFormat(offset int, limit int, Langcode_ string, BodySummary_ string, BodyFormat_ string) (*[]*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = new([]*BlockContentRevision_body)
	err := Engine.Table("block_content_revision__body").Where("langcode = ? and body_summary = ? and body_format = ?", Langcode_, BodySummary_, BodyFormat_).Limit(limit, offset).Find(_BlockContentRevision_body)
	return _BlockContentRevision_body, err
}

// GetBlockContentRevision_bodiesByDeltaAndBodyValueAndBodySummary Get BlockContentRevision_bodies via DeltaAndBodyValueAndBodySummary
func GetBlockContentRevision_bodiesByDeltaAndBodyValueAndBodySummary(offset int, limit int, Delta_ int, BodyValue_ string, BodySummary_ string) (*[]*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = new([]*BlockContentRevision_body)
	err := Engine.Table("block_content_revision__body").Where("delta = ? and body_value = ? and body_summary = ?", Delta_, BodyValue_, BodySummary_).Limit(limit, offset).Find(_BlockContentRevision_body)
	return _BlockContentRevision_body, err
}

// GetBlockContentRevision_bodiesByDeltaAndBodyValueAndBodyFormat Get BlockContentRevision_bodies via DeltaAndBodyValueAndBodyFormat
func GetBlockContentRevision_bodiesByDeltaAndBodyValueAndBodyFormat(offset int, limit int, Delta_ int, BodyValue_ string, BodyFormat_ string) (*[]*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = new([]*BlockContentRevision_body)
	err := Engine.Table("block_content_revision__body").Where("delta = ? and body_value = ? and body_format = ?", Delta_, BodyValue_, BodyFormat_).Limit(limit, offset).Find(_BlockContentRevision_body)
	return _BlockContentRevision_body, err
}

// GetBlockContentRevision_bodiesByDeltaAndBodySummaryAndBodyFormat Get BlockContentRevision_bodies via DeltaAndBodySummaryAndBodyFormat
func GetBlockContentRevision_bodiesByDeltaAndBodySummaryAndBodyFormat(offset int, limit int, Delta_ int, BodySummary_ string, BodyFormat_ string) (*[]*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = new([]*BlockContentRevision_body)
	err := Engine.Table("block_content_revision__body").Where("delta = ? and body_summary = ? and body_format = ?", Delta_, BodySummary_, BodyFormat_).Limit(limit, offset).Find(_BlockContentRevision_body)
	return _BlockContentRevision_body, err
}

// GetBlockContentRevision_bodiesByBodyValueAndBodySummaryAndBodyFormat Get BlockContentRevision_bodies via BodyValueAndBodySummaryAndBodyFormat
func GetBlockContentRevision_bodiesByBodyValueAndBodySummaryAndBodyFormat(offset int, limit int, BodyValue_ string, BodySummary_ string, BodyFormat_ string) (*[]*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = new([]*BlockContentRevision_body)
	err := Engine.Table("block_content_revision__body").Where("body_value = ? and body_summary = ? and body_format = ?", BodyValue_, BodySummary_, BodyFormat_).Limit(limit, offset).Find(_BlockContentRevision_body)
	return _BlockContentRevision_body, err
}

// GetBlockContentRevision_bodiesByBundleAndDeleted Get BlockContentRevision_bodies via BundleAndDeleted
func GetBlockContentRevision_bodiesByBundleAndDeleted(offset int, limit int, Bundle_ string, Deleted_ int) (*[]*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = new([]*BlockContentRevision_body)
	err := Engine.Table("block_content_revision__body").Where("bundle = ? and deleted = ?", Bundle_, Deleted_).Limit(limit, offset).Find(_BlockContentRevision_body)
	return _BlockContentRevision_body, err
}

// GetBlockContentRevision_bodiesByBundleAndEntityId Get BlockContentRevision_bodies via BundleAndEntityId
func GetBlockContentRevision_bodiesByBundleAndEntityId(offset int, limit int, Bundle_ string, EntityId_ int) (*[]*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = new([]*BlockContentRevision_body)
	err := Engine.Table("block_content_revision__body").Where("bundle = ? and entity_id = ?", Bundle_, EntityId_).Limit(limit, offset).Find(_BlockContentRevision_body)
	return _BlockContentRevision_body, err
}

// GetBlockContentRevision_bodiesByBundleAndRevisionId Get BlockContentRevision_bodies via BundleAndRevisionId
func GetBlockContentRevision_bodiesByBundleAndRevisionId(offset int, limit int, Bundle_ string, RevisionId_ int) (*[]*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = new([]*BlockContentRevision_body)
	err := Engine.Table("block_content_revision__body").Where("bundle = ? and revision_id = ?", Bundle_, RevisionId_).Limit(limit, offset).Find(_BlockContentRevision_body)
	return _BlockContentRevision_body, err
}

// GetBlockContentRevision_bodiesByBundleAndLangcode Get BlockContentRevision_bodies via BundleAndLangcode
func GetBlockContentRevision_bodiesByBundleAndLangcode(offset int, limit int, Bundle_ string, Langcode_ string) (*[]*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = new([]*BlockContentRevision_body)
	err := Engine.Table("block_content_revision__body").Where("bundle = ? and langcode = ?", Bundle_, Langcode_).Limit(limit, offset).Find(_BlockContentRevision_body)
	return _BlockContentRevision_body, err
}

// GetBlockContentRevision_bodiesByBundleAndDelta Get BlockContentRevision_bodies via BundleAndDelta
func GetBlockContentRevision_bodiesByBundleAndDelta(offset int, limit int, Bundle_ string, Delta_ int) (*[]*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = new([]*BlockContentRevision_body)
	err := Engine.Table("block_content_revision__body").Where("bundle = ? and delta = ?", Bundle_, Delta_).Limit(limit, offset).Find(_BlockContentRevision_body)
	return _BlockContentRevision_body, err
}

// GetBlockContentRevision_bodiesByBundleAndBodyValue Get BlockContentRevision_bodies via BundleAndBodyValue
func GetBlockContentRevision_bodiesByBundleAndBodyValue(offset int, limit int, Bundle_ string, BodyValue_ string) (*[]*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = new([]*BlockContentRevision_body)
	err := Engine.Table("block_content_revision__body").Where("bundle = ? and body_value = ?", Bundle_, BodyValue_).Limit(limit, offset).Find(_BlockContentRevision_body)
	return _BlockContentRevision_body, err
}

// GetBlockContentRevision_bodiesByBundleAndBodySummary Get BlockContentRevision_bodies via BundleAndBodySummary
func GetBlockContentRevision_bodiesByBundleAndBodySummary(offset int, limit int, Bundle_ string, BodySummary_ string) (*[]*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = new([]*BlockContentRevision_body)
	err := Engine.Table("block_content_revision__body").Where("bundle = ? and body_summary = ?", Bundle_, BodySummary_).Limit(limit, offset).Find(_BlockContentRevision_body)
	return _BlockContentRevision_body, err
}

// GetBlockContentRevision_bodiesByBundleAndBodyFormat Get BlockContentRevision_bodies via BundleAndBodyFormat
func GetBlockContentRevision_bodiesByBundleAndBodyFormat(offset int, limit int, Bundle_ string, BodyFormat_ string) (*[]*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = new([]*BlockContentRevision_body)
	err := Engine.Table("block_content_revision__body").Where("bundle = ? and body_format = ?", Bundle_, BodyFormat_).Limit(limit, offset).Find(_BlockContentRevision_body)
	return _BlockContentRevision_body, err
}

// GetBlockContentRevision_bodiesByDeletedAndEntityId Get BlockContentRevision_bodies via DeletedAndEntityId
func GetBlockContentRevision_bodiesByDeletedAndEntityId(offset int, limit int, Deleted_ int, EntityId_ int) (*[]*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = new([]*BlockContentRevision_body)
	err := Engine.Table("block_content_revision__body").Where("deleted = ? and entity_id = ?", Deleted_, EntityId_).Limit(limit, offset).Find(_BlockContentRevision_body)
	return _BlockContentRevision_body, err
}

// GetBlockContentRevision_bodiesByDeletedAndRevisionId Get BlockContentRevision_bodies via DeletedAndRevisionId
func GetBlockContentRevision_bodiesByDeletedAndRevisionId(offset int, limit int, Deleted_ int, RevisionId_ int) (*[]*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = new([]*BlockContentRevision_body)
	err := Engine.Table("block_content_revision__body").Where("deleted = ? and revision_id = ?", Deleted_, RevisionId_).Limit(limit, offset).Find(_BlockContentRevision_body)
	return _BlockContentRevision_body, err
}

// GetBlockContentRevision_bodiesByDeletedAndLangcode Get BlockContentRevision_bodies via DeletedAndLangcode
func GetBlockContentRevision_bodiesByDeletedAndLangcode(offset int, limit int, Deleted_ int, Langcode_ string) (*[]*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = new([]*BlockContentRevision_body)
	err := Engine.Table("block_content_revision__body").Where("deleted = ? and langcode = ?", Deleted_, Langcode_).Limit(limit, offset).Find(_BlockContentRevision_body)
	return _BlockContentRevision_body, err
}

// GetBlockContentRevision_bodiesByDeletedAndDelta Get BlockContentRevision_bodies via DeletedAndDelta
func GetBlockContentRevision_bodiesByDeletedAndDelta(offset int, limit int, Deleted_ int, Delta_ int) (*[]*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = new([]*BlockContentRevision_body)
	err := Engine.Table("block_content_revision__body").Where("deleted = ? and delta = ?", Deleted_, Delta_).Limit(limit, offset).Find(_BlockContentRevision_body)
	return _BlockContentRevision_body, err
}

// GetBlockContentRevision_bodiesByDeletedAndBodyValue Get BlockContentRevision_bodies via DeletedAndBodyValue
func GetBlockContentRevision_bodiesByDeletedAndBodyValue(offset int, limit int, Deleted_ int, BodyValue_ string) (*[]*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = new([]*BlockContentRevision_body)
	err := Engine.Table("block_content_revision__body").Where("deleted = ? and body_value = ?", Deleted_, BodyValue_).Limit(limit, offset).Find(_BlockContentRevision_body)
	return _BlockContentRevision_body, err
}

// GetBlockContentRevision_bodiesByDeletedAndBodySummary Get BlockContentRevision_bodies via DeletedAndBodySummary
func GetBlockContentRevision_bodiesByDeletedAndBodySummary(offset int, limit int, Deleted_ int, BodySummary_ string) (*[]*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = new([]*BlockContentRevision_body)
	err := Engine.Table("block_content_revision__body").Where("deleted = ? and body_summary = ?", Deleted_, BodySummary_).Limit(limit, offset).Find(_BlockContentRevision_body)
	return _BlockContentRevision_body, err
}

// GetBlockContentRevision_bodiesByDeletedAndBodyFormat Get BlockContentRevision_bodies via DeletedAndBodyFormat
func GetBlockContentRevision_bodiesByDeletedAndBodyFormat(offset int, limit int, Deleted_ int, BodyFormat_ string) (*[]*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = new([]*BlockContentRevision_body)
	err := Engine.Table("block_content_revision__body").Where("deleted = ? and body_format = ?", Deleted_, BodyFormat_).Limit(limit, offset).Find(_BlockContentRevision_body)
	return _BlockContentRevision_body, err
}

// GetBlockContentRevision_bodiesByEntityIdAndRevisionId Get BlockContentRevision_bodies via EntityIdAndRevisionId
func GetBlockContentRevision_bodiesByEntityIdAndRevisionId(offset int, limit int, EntityId_ int, RevisionId_ int) (*[]*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = new([]*BlockContentRevision_body)
	err := Engine.Table("block_content_revision__body").Where("entity_id = ? and revision_id = ?", EntityId_, RevisionId_).Limit(limit, offset).Find(_BlockContentRevision_body)
	return _BlockContentRevision_body, err
}

// GetBlockContentRevision_bodiesByEntityIdAndLangcode Get BlockContentRevision_bodies via EntityIdAndLangcode
func GetBlockContentRevision_bodiesByEntityIdAndLangcode(offset int, limit int, EntityId_ int, Langcode_ string) (*[]*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = new([]*BlockContentRevision_body)
	err := Engine.Table("block_content_revision__body").Where("entity_id = ? and langcode = ?", EntityId_, Langcode_).Limit(limit, offset).Find(_BlockContentRevision_body)
	return _BlockContentRevision_body, err
}

// GetBlockContentRevision_bodiesByEntityIdAndDelta Get BlockContentRevision_bodies via EntityIdAndDelta
func GetBlockContentRevision_bodiesByEntityIdAndDelta(offset int, limit int, EntityId_ int, Delta_ int) (*[]*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = new([]*BlockContentRevision_body)
	err := Engine.Table("block_content_revision__body").Where("entity_id = ? and delta = ?", EntityId_, Delta_).Limit(limit, offset).Find(_BlockContentRevision_body)
	return _BlockContentRevision_body, err
}

// GetBlockContentRevision_bodiesByEntityIdAndBodyValue Get BlockContentRevision_bodies via EntityIdAndBodyValue
func GetBlockContentRevision_bodiesByEntityIdAndBodyValue(offset int, limit int, EntityId_ int, BodyValue_ string) (*[]*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = new([]*BlockContentRevision_body)
	err := Engine.Table("block_content_revision__body").Where("entity_id = ? and body_value = ?", EntityId_, BodyValue_).Limit(limit, offset).Find(_BlockContentRevision_body)
	return _BlockContentRevision_body, err
}

// GetBlockContentRevision_bodiesByEntityIdAndBodySummary Get BlockContentRevision_bodies via EntityIdAndBodySummary
func GetBlockContentRevision_bodiesByEntityIdAndBodySummary(offset int, limit int, EntityId_ int, BodySummary_ string) (*[]*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = new([]*BlockContentRevision_body)
	err := Engine.Table("block_content_revision__body").Where("entity_id = ? and body_summary = ?", EntityId_, BodySummary_).Limit(limit, offset).Find(_BlockContentRevision_body)
	return _BlockContentRevision_body, err
}

// GetBlockContentRevision_bodiesByEntityIdAndBodyFormat Get BlockContentRevision_bodies via EntityIdAndBodyFormat
func GetBlockContentRevision_bodiesByEntityIdAndBodyFormat(offset int, limit int, EntityId_ int, BodyFormat_ string) (*[]*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = new([]*BlockContentRevision_body)
	err := Engine.Table("block_content_revision__body").Where("entity_id = ? and body_format = ?", EntityId_, BodyFormat_).Limit(limit, offset).Find(_BlockContentRevision_body)
	return _BlockContentRevision_body, err
}

// GetBlockContentRevision_bodiesByRevisionIdAndLangcode Get BlockContentRevision_bodies via RevisionIdAndLangcode
func GetBlockContentRevision_bodiesByRevisionIdAndLangcode(offset int, limit int, RevisionId_ int, Langcode_ string) (*[]*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = new([]*BlockContentRevision_body)
	err := Engine.Table("block_content_revision__body").Where("revision_id = ? and langcode = ?", RevisionId_, Langcode_).Limit(limit, offset).Find(_BlockContentRevision_body)
	return _BlockContentRevision_body, err
}

// GetBlockContentRevision_bodiesByRevisionIdAndDelta Get BlockContentRevision_bodies via RevisionIdAndDelta
func GetBlockContentRevision_bodiesByRevisionIdAndDelta(offset int, limit int, RevisionId_ int, Delta_ int) (*[]*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = new([]*BlockContentRevision_body)
	err := Engine.Table("block_content_revision__body").Where("revision_id = ? and delta = ?", RevisionId_, Delta_).Limit(limit, offset).Find(_BlockContentRevision_body)
	return _BlockContentRevision_body, err
}

// GetBlockContentRevision_bodiesByRevisionIdAndBodyValue Get BlockContentRevision_bodies via RevisionIdAndBodyValue
func GetBlockContentRevision_bodiesByRevisionIdAndBodyValue(offset int, limit int, RevisionId_ int, BodyValue_ string) (*[]*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = new([]*BlockContentRevision_body)
	err := Engine.Table("block_content_revision__body").Where("revision_id = ? and body_value = ?", RevisionId_, BodyValue_).Limit(limit, offset).Find(_BlockContentRevision_body)
	return _BlockContentRevision_body, err
}

// GetBlockContentRevision_bodiesByRevisionIdAndBodySummary Get BlockContentRevision_bodies via RevisionIdAndBodySummary
func GetBlockContentRevision_bodiesByRevisionIdAndBodySummary(offset int, limit int, RevisionId_ int, BodySummary_ string) (*[]*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = new([]*BlockContentRevision_body)
	err := Engine.Table("block_content_revision__body").Where("revision_id = ? and body_summary = ?", RevisionId_, BodySummary_).Limit(limit, offset).Find(_BlockContentRevision_body)
	return _BlockContentRevision_body, err
}

// GetBlockContentRevision_bodiesByRevisionIdAndBodyFormat Get BlockContentRevision_bodies via RevisionIdAndBodyFormat
func GetBlockContentRevision_bodiesByRevisionIdAndBodyFormat(offset int, limit int, RevisionId_ int, BodyFormat_ string) (*[]*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = new([]*BlockContentRevision_body)
	err := Engine.Table("block_content_revision__body").Where("revision_id = ? and body_format = ?", RevisionId_, BodyFormat_).Limit(limit, offset).Find(_BlockContentRevision_body)
	return _BlockContentRevision_body, err
}

// GetBlockContentRevision_bodiesByLangcodeAndDelta Get BlockContentRevision_bodies via LangcodeAndDelta
func GetBlockContentRevision_bodiesByLangcodeAndDelta(offset int, limit int, Langcode_ string, Delta_ int) (*[]*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = new([]*BlockContentRevision_body)
	err := Engine.Table("block_content_revision__body").Where("langcode = ? and delta = ?", Langcode_, Delta_).Limit(limit, offset).Find(_BlockContentRevision_body)
	return _BlockContentRevision_body, err
}

// GetBlockContentRevision_bodiesByLangcodeAndBodyValue Get BlockContentRevision_bodies via LangcodeAndBodyValue
func GetBlockContentRevision_bodiesByLangcodeAndBodyValue(offset int, limit int, Langcode_ string, BodyValue_ string) (*[]*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = new([]*BlockContentRevision_body)
	err := Engine.Table("block_content_revision__body").Where("langcode = ? and body_value = ?", Langcode_, BodyValue_).Limit(limit, offset).Find(_BlockContentRevision_body)
	return _BlockContentRevision_body, err
}

// GetBlockContentRevision_bodiesByLangcodeAndBodySummary Get BlockContentRevision_bodies via LangcodeAndBodySummary
func GetBlockContentRevision_bodiesByLangcodeAndBodySummary(offset int, limit int, Langcode_ string, BodySummary_ string) (*[]*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = new([]*BlockContentRevision_body)
	err := Engine.Table("block_content_revision__body").Where("langcode = ? and body_summary = ?", Langcode_, BodySummary_).Limit(limit, offset).Find(_BlockContentRevision_body)
	return _BlockContentRevision_body, err
}

// GetBlockContentRevision_bodiesByLangcodeAndBodyFormat Get BlockContentRevision_bodies via LangcodeAndBodyFormat
func GetBlockContentRevision_bodiesByLangcodeAndBodyFormat(offset int, limit int, Langcode_ string, BodyFormat_ string) (*[]*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = new([]*BlockContentRevision_body)
	err := Engine.Table("block_content_revision__body").Where("langcode = ? and body_format = ?", Langcode_, BodyFormat_).Limit(limit, offset).Find(_BlockContentRevision_body)
	return _BlockContentRevision_body, err
}

// GetBlockContentRevision_bodiesByDeltaAndBodyValue Get BlockContentRevision_bodies via DeltaAndBodyValue
func GetBlockContentRevision_bodiesByDeltaAndBodyValue(offset int, limit int, Delta_ int, BodyValue_ string) (*[]*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = new([]*BlockContentRevision_body)
	err := Engine.Table("block_content_revision__body").Where("delta = ? and body_value = ?", Delta_, BodyValue_).Limit(limit, offset).Find(_BlockContentRevision_body)
	return _BlockContentRevision_body, err
}

// GetBlockContentRevision_bodiesByDeltaAndBodySummary Get BlockContentRevision_bodies via DeltaAndBodySummary
func GetBlockContentRevision_bodiesByDeltaAndBodySummary(offset int, limit int, Delta_ int, BodySummary_ string) (*[]*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = new([]*BlockContentRevision_body)
	err := Engine.Table("block_content_revision__body").Where("delta = ? and body_summary = ?", Delta_, BodySummary_).Limit(limit, offset).Find(_BlockContentRevision_body)
	return _BlockContentRevision_body, err
}

// GetBlockContentRevision_bodiesByDeltaAndBodyFormat Get BlockContentRevision_bodies via DeltaAndBodyFormat
func GetBlockContentRevision_bodiesByDeltaAndBodyFormat(offset int, limit int, Delta_ int, BodyFormat_ string) (*[]*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = new([]*BlockContentRevision_body)
	err := Engine.Table("block_content_revision__body").Where("delta = ? and body_format = ?", Delta_, BodyFormat_).Limit(limit, offset).Find(_BlockContentRevision_body)
	return _BlockContentRevision_body, err
}

// GetBlockContentRevision_bodiesByBodyValueAndBodySummary Get BlockContentRevision_bodies via BodyValueAndBodySummary
func GetBlockContentRevision_bodiesByBodyValueAndBodySummary(offset int, limit int, BodyValue_ string, BodySummary_ string) (*[]*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = new([]*BlockContentRevision_body)
	err := Engine.Table("block_content_revision__body").Where("body_value = ? and body_summary = ?", BodyValue_, BodySummary_).Limit(limit, offset).Find(_BlockContentRevision_body)
	return _BlockContentRevision_body, err
}

// GetBlockContentRevision_bodiesByBodyValueAndBodyFormat Get BlockContentRevision_bodies via BodyValueAndBodyFormat
func GetBlockContentRevision_bodiesByBodyValueAndBodyFormat(offset int, limit int, BodyValue_ string, BodyFormat_ string) (*[]*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = new([]*BlockContentRevision_body)
	err := Engine.Table("block_content_revision__body").Where("body_value = ? and body_format = ?", BodyValue_, BodyFormat_).Limit(limit, offset).Find(_BlockContentRevision_body)
	return _BlockContentRevision_body, err
}

// GetBlockContentRevision_bodiesByBodySummaryAndBodyFormat Get BlockContentRevision_bodies via BodySummaryAndBodyFormat
func GetBlockContentRevision_bodiesByBodySummaryAndBodyFormat(offset int, limit int, BodySummary_ string, BodyFormat_ string) (*[]*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = new([]*BlockContentRevision_body)
	err := Engine.Table("block_content_revision__body").Where("body_summary = ? and body_format = ?", BodySummary_, BodyFormat_).Limit(limit, offset).Find(_BlockContentRevision_body)
	return _BlockContentRevision_body, err
}

// GetBlockContentRevision_bodies Get BlockContentRevision_bodies via field
func GetBlockContentRevision_bodies(offset int, limit int, field string) (*[]*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = new([]*BlockContentRevision_body)
	err := Engine.Table("block_content_revision__body").Limit(limit, offset).Desc(field).Find(_BlockContentRevision_body)
	return _BlockContentRevision_body, err
}

// GetBlockContentRevision_bodiesViaBundle Get BlockContentRevision_bodys via Bundle
func GetBlockContentRevision_bodiesViaBundle(offset int, limit int, Bundle_ string, field string) (*[]*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = new([]*BlockContentRevision_body)
	err := Engine.Table("block_content_revision__body").Where("bundle = ?", Bundle_).Limit(limit, offset).Desc(field).Find(_BlockContentRevision_body)
	return _BlockContentRevision_body, err
}

// GetBlockContentRevision_bodiesViaDeleted Get BlockContentRevision_bodys via Deleted
func GetBlockContentRevision_bodiesViaDeleted(offset int, limit int, Deleted_ int, field string) (*[]*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = new([]*BlockContentRevision_body)
	err := Engine.Table("block_content_revision__body").Where("deleted = ?", Deleted_).Limit(limit, offset).Desc(field).Find(_BlockContentRevision_body)
	return _BlockContentRevision_body, err
}

// GetBlockContentRevision_bodiesViaEntityId Get BlockContentRevision_bodys via EntityId
func GetBlockContentRevision_bodiesViaEntityId(offset int, limit int, EntityId_ int, field string) (*[]*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = new([]*BlockContentRevision_body)
	err := Engine.Table("block_content_revision__body").Where("entity_id = ?", EntityId_).Limit(limit, offset).Desc(field).Find(_BlockContentRevision_body)
	return _BlockContentRevision_body, err
}

// GetBlockContentRevision_bodiesViaRevisionId Get BlockContentRevision_bodys via RevisionId
func GetBlockContentRevision_bodiesViaRevisionId(offset int, limit int, RevisionId_ int, field string) (*[]*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = new([]*BlockContentRevision_body)
	err := Engine.Table("block_content_revision__body").Where("revision_id = ?", RevisionId_).Limit(limit, offset).Desc(field).Find(_BlockContentRevision_body)
	return _BlockContentRevision_body, err
}

// GetBlockContentRevision_bodiesViaLangcode Get BlockContentRevision_bodys via Langcode
func GetBlockContentRevision_bodiesViaLangcode(offset int, limit int, Langcode_ string, field string) (*[]*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = new([]*BlockContentRevision_body)
	err := Engine.Table("block_content_revision__body").Where("langcode = ?", Langcode_).Limit(limit, offset).Desc(field).Find(_BlockContentRevision_body)
	return _BlockContentRevision_body, err
}

// GetBlockContentRevision_bodiesViaDelta Get BlockContentRevision_bodys via Delta
func GetBlockContentRevision_bodiesViaDelta(offset int, limit int, Delta_ int, field string) (*[]*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = new([]*BlockContentRevision_body)
	err := Engine.Table("block_content_revision__body").Where("delta = ?", Delta_).Limit(limit, offset).Desc(field).Find(_BlockContentRevision_body)
	return _BlockContentRevision_body, err
}

// GetBlockContentRevision_bodiesViaBodyValue Get BlockContentRevision_bodys via BodyValue
func GetBlockContentRevision_bodiesViaBodyValue(offset int, limit int, BodyValue_ string, field string) (*[]*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = new([]*BlockContentRevision_body)
	err := Engine.Table("block_content_revision__body").Where("body_value = ?", BodyValue_).Limit(limit, offset).Desc(field).Find(_BlockContentRevision_body)
	return _BlockContentRevision_body, err
}

// GetBlockContentRevision_bodiesViaBodySummary Get BlockContentRevision_bodys via BodySummary
func GetBlockContentRevision_bodiesViaBodySummary(offset int, limit int, BodySummary_ string, field string) (*[]*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = new([]*BlockContentRevision_body)
	err := Engine.Table("block_content_revision__body").Where("body_summary = ?", BodySummary_).Limit(limit, offset).Desc(field).Find(_BlockContentRevision_body)
	return _BlockContentRevision_body, err
}

// GetBlockContentRevision_bodiesViaBodyFormat Get BlockContentRevision_bodys via BodyFormat
func GetBlockContentRevision_bodiesViaBodyFormat(offset int, limit int, BodyFormat_ string, field string) (*[]*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = new([]*BlockContentRevision_body)
	err := Engine.Table("block_content_revision__body").Where("body_format = ?", BodyFormat_).Limit(limit, offset).Desc(field).Find(_BlockContentRevision_body)
	return _BlockContentRevision_body, err
}

// HasBlockContentRevision_bodyViaBundle Has BlockContentRevision_body via Bundle
func HasBlockContentRevision_bodyViaBundle(iBundle string) bool {
	if has, err := Engine.Where("bundle = ?", iBundle).Get(new(BlockContentRevision_body)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasBlockContentRevision_bodyViaDeleted Has BlockContentRevision_body via Deleted
func HasBlockContentRevision_bodyViaDeleted(iDeleted int) bool {
	if has, err := Engine.Where("deleted = ?", iDeleted).Get(new(BlockContentRevision_body)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasBlockContentRevision_bodyViaEntityId Has BlockContentRevision_body via EntityId
func HasBlockContentRevision_bodyViaEntityId(iEntityId int) bool {
	if has, err := Engine.Where("entity_id = ?", iEntityId).Get(new(BlockContentRevision_body)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasBlockContentRevision_bodyViaRevisionId Has BlockContentRevision_body via RevisionId
func HasBlockContentRevision_bodyViaRevisionId(iRevisionId int) bool {
	if has, err := Engine.Where("revision_id = ?", iRevisionId).Get(new(BlockContentRevision_body)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasBlockContentRevision_bodyViaLangcode Has BlockContentRevision_body via Langcode
func HasBlockContentRevision_bodyViaLangcode(iLangcode string) bool {
	if has, err := Engine.Where("langcode = ?", iLangcode).Get(new(BlockContentRevision_body)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasBlockContentRevision_bodyViaDelta Has BlockContentRevision_body via Delta
func HasBlockContentRevision_bodyViaDelta(iDelta int) bool {
	if has, err := Engine.Where("delta = ?", iDelta).Get(new(BlockContentRevision_body)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasBlockContentRevision_bodyViaBodyValue Has BlockContentRevision_body via BodyValue
func HasBlockContentRevision_bodyViaBodyValue(iBodyValue string) bool {
	if has, err := Engine.Where("body_value = ?", iBodyValue).Get(new(BlockContentRevision_body)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasBlockContentRevision_bodyViaBodySummary Has BlockContentRevision_body via BodySummary
func HasBlockContentRevision_bodyViaBodySummary(iBodySummary string) bool {
	if has, err := Engine.Where("body_summary = ?", iBodySummary).Get(new(BlockContentRevision_body)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasBlockContentRevision_bodyViaBodyFormat Has BlockContentRevision_body via BodyFormat
func HasBlockContentRevision_bodyViaBodyFormat(iBodyFormat string) bool {
	if has, err := Engine.Where("body_format = ?", iBodyFormat).Get(new(BlockContentRevision_body)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// GetBlockContentRevision_bodyViaBundle Get BlockContentRevision_body via Bundle
func GetBlockContentRevision_bodyViaBundle(iBundle string) (*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = &BlockContentRevision_body{Bundle: iBundle}
	has, err := Engine.Get(_BlockContentRevision_body)
	if has {
		return _BlockContentRevision_body, err
	} else {
		return nil, err
	}
}

// GetBlockContentRevision_bodyViaDeleted Get BlockContentRevision_body via Deleted
func GetBlockContentRevision_bodyViaDeleted(iDeleted int) (*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = &BlockContentRevision_body{Deleted: iDeleted}
	has, err := Engine.Get(_BlockContentRevision_body)
	if has {
		return _BlockContentRevision_body, err
	} else {
		return nil, err
	}
}

// GetBlockContentRevision_bodyViaEntityId Get BlockContentRevision_body via EntityId
func GetBlockContentRevision_bodyViaEntityId(iEntityId int) (*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = &BlockContentRevision_body{EntityId: iEntityId}
	has, err := Engine.Get(_BlockContentRevision_body)
	if has {
		return _BlockContentRevision_body, err
	} else {
		return nil, err
	}
}

// GetBlockContentRevision_bodyViaRevisionId Get BlockContentRevision_body via RevisionId
func GetBlockContentRevision_bodyViaRevisionId(iRevisionId int) (*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = &BlockContentRevision_body{RevisionId: iRevisionId}
	has, err := Engine.Get(_BlockContentRevision_body)
	if has {
		return _BlockContentRevision_body, err
	} else {
		return nil, err
	}
}

// GetBlockContentRevision_bodyViaLangcode Get BlockContentRevision_body via Langcode
func GetBlockContentRevision_bodyViaLangcode(iLangcode string) (*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = &BlockContentRevision_body{Langcode: iLangcode}
	has, err := Engine.Get(_BlockContentRevision_body)
	if has {
		return _BlockContentRevision_body, err
	} else {
		return nil, err
	}
}

// GetBlockContentRevision_bodyViaDelta Get BlockContentRevision_body via Delta
func GetBlockContentRevision_bodyViaDelta(iDelta int) (*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = &BlockContentRevision_body{Delta: iDelta}
	has, err := Engine.Get(_BlockContentRevision_body)
	if has {
		return _BlockContentRevision_body, err
	} else {
		return nil, err
	}
}

// GetBlockContentRevision_bodyViaBodyValue Get BlockContentRevision_body via BodyValue
func GetBlockContentRevision_bodyViaBodyValue(iBodyValue string) (*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = &BlockContentRevision_body{BodyValue: iBodyValue}
	has, err := Engine.Get(_BlockContentRevision_body)
	if has {
		return _BlockContentRevision_body, err
	} else {
		return nil, err
	}
}

// GetBlockContentRevision_bodyViaBodySummary Get BlockContentRevision_body via BodySummary
func GetBlockContentRevision_bodyViaBodySummary(iBodySummary string) (*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = &BlockContentRevision_body{BodySummary: iBodySummary}
	has, err := Engine.Get(_BlockContentRevision_body)
	if has {
		return _BlockContentRevision_body, err
	} else {
		return nil, err
	}
}

// GetBlockContentRevision_bodyViaBodyFormat Get BlockContentRevision_body via BodyFormat
func GetBlockContentRevision_bodyViaBodyFormat(iBodyFormat string) (*BlockContentRevision_body, error) {
	var _BlockContentRevision_body = &BlockContentRevision_body{BodyFormat: iBodyFormat}
	has, err := Engine.Get(_BlockContentRevision_body)
	if has {
		return _BlockContentRevision_body, err
	} else {
		return nil, err
	}
}

// SetBlockContentRevision_bodyViaBundle Set BlockContentRevision_body via Bundle
func SetBlockContentRevision_bodyViaBundle(iBundle string, block_content_revision__body *BlockContentRevision_body) (int64, error) {
	block_content_revision__body.Bundle = iBundle
	return Engine.Insert(block_content_revision__body)
}

// SetBlockContentRevision_bodyViaDeleted Set BlockContentRevision_body via Deleted
func SetBlockContentRevision_bodyViaDeleted(iDeleted int, block_content_revision__body *BlockContentRevision_body) (int64, error) {
	block_content_revision__body.Deleted = iDeleted
	return Engine.Insert(block_content_revision__body)
}

// SetBlockContentRevision_bodyViaEntityId Set BlockContentRevision_body via EntityId
func SetBlockContentRevision_bodyViaEntityId(iEntityId int, block_content_revision__body *BlockContentRevision_body) (int64, error) {
	block_content_revision__body.EntityId = iEntityId
	return Engine.Insert(block_content_revision__body)
}

// SetBlockContentRevision_bodyViaRevisionId Set BlockContentRevision_body via RevisionId
func SetBlockContentRevision_bodyViaRevisionId(iRevisionId int, block_content_revision__body *BlockContentRevision_body) (int64, error) {
	block_content_revision__body.RevisionId = iRevisionId
	return Engine.Insert(block_content_revision__body)
}

// SetBlockContentRevision_bodyViaLangcode Set BlockContentRevision_body via Langcode
func SetBlockContentRevision_bodyViaLangcode(iLangcode string, block_content_revision__body *BlockContentRevision_body) (int64, error) {
	block_content_revision__body.Langcode = iLangcode
	return Engine.Insert(block_content_revision__body)
}

// SetBlockContentRevision_bodyViaDelta Set BlockContentRevision_body via Delta
func SetBlockContentRevision_bodyViaDelta(iDelta int, block_content_revision__body *BlockContentRevision_body) (int64, error) {
	block_content_revision__body.Delta = iDelta
	return Engine.Insert(block_content_revision__body)
}

// SetBlockContentRevision_bodyViaBodyValue Set BlockContentRevision_body via BodyValue
func SetBlockContentRevision_bodyViaBodyValue(iBodyValue string, block_content_revision__body *BlockContentRevision_body) (int64, error) {
	block_content_revision__body.BodyValue = iBodyValue
	return Engine.Insert(block_content_revision__body)
}

// SetBlockContentRevision_bodyViaBodySummary Set BlockContentRevision_body via BodySummary
func SetBlockContentRevision_bodyViaBodySummary(iBodySummary string, block_content_revision__body *BlockContentRevision_body) (int64, error) {
	block_content_revision__body.BodySummary = iBodySummary
	return Engine.Insert(block_content_revision__body)
}

// SetBlockContentRevision_bodyViaBodyFormat Set BlockContentRevision_body via BodyFormat
func SetBlockContentRevision_bodyViaBodyFormat(iBodyFormat string, block_content_revision__body *BlockContentRevision_body) (int64, error) {
	block_content_revision__body.BodyFormat = iBodyFormat
	return Engine.Insert(block_content_revision__body)
}

// AddBlockContentRevision_body Add BlockContentRevision_body via all columns
func AddBlockContentRevision_body(iBundle string, iDeleted int, iEntityId int, iRevisionId int, iLangcode string, iDelta int, iBodyValue string, iBodySummary string, iBodyFormat string) error {
	_BlockContentRevision_body := &BlockContentRevision_body{Bundle: iBundle, Deleted: iDeleted, EntityId: iEntityId, RevisionId: iRevisionId, Langcode: iLangcode, Delta: iDelta, BodyValue: iBodyValue, BodySummary: iBodySummary, BodyFormat: iBodyFormat}
	if _, err := Engine.Insert(_BlockContentRevision_body); err != nil {
		return err
	} else {
		return nil
	}
}

// PostBlockContentRevision_body Post BlockContentRevision_body via iBlockContentRevision_body
func PostBlockContentRevision_body(iBlockContentRevision_body *BlockContentRevision_body) (string, error) {
	_, err := Engine.Insert(iBlockContentRevision_body)
	return iBlockContentRevision_body.Bundle, err
}

// PutBlockContentRevision_body Put BlockContentRevision_body
func PutBlockContentRevision_body(iBlockContentRevision_body *BlockContentRevision_body) (string, error) {
	_, err := Engine.Id(iBlockContentRevision_body.Bundle).Update(iBlockContentRevision_body)
	return iBlockContentRevision_body.Bundle, err
}

// PutBlockContentRevision_bodyViaBundle Put BlockContentRevision_body via Bundle
func PutBlockContentRevision_bodyViaBundle(Bundle_ string, iBlockContentRevision_body *BlockContentRevision_body) (int64, error) {
	row, err := Engine.Update(iBlockContentRevision_body, &BlockContentRevision_body{Bundle: Bundle_})
	return row, err
}

// PutBlockContentRevision_bodyViaDeleted Put BlockContentRevision_body via Deleted
func PutBlockContentRevision_bodyViaDeleted(Deleted_ int, iBlockContentRevision_body *BlockContentRevision_body) (int64, error) {
	row, err := Engine.Update(iBlockContentRevision_body, &BlockContentRevision_body{Deleted: Deleted_})
	return row, err
}

// PutBlockContentRevision_bodyViaEntityId Put BlockContentRevision_body via EntityId
func PutBlockContentRevision_bodyViaEntityId(EntityId_ int, iBlockContentRevision_body *BlockContentRevision_body) (int64, error) {
	row, err := Engine.Update(iBlockContentRevision_body, &BlockContentRevision_body{EntityId: EntityId_})
	return row, err
}

// PutBlockContentRevision_bodyViaRevisionId Put BlockContentRevision_body via RevisionId
func PutBlockContentRevision_bodyViaRevisionId(RevisionId_ int, iBlockContentRevision_body *BlockContentRevision_body) (int64, error) {
	row, err := Engine.Update(iBlockContentRevision_body, &BlockContentRevision_body{RevisionId: RevisionId_})
	return row, err
}

// PutBlockContentRevision_bodyViaLangcode Put BlockContentRevision_body via Langcode
func PutBlockContentRevision_bodyViaLangcode(Langcode_ string, iBlockContentRevision_body *BlockContentRevision_body) (int64, error) {
	row, err := Engine.Update(iBlockContentRevision_body, &BlockContentRevision_body{Langcode: Langcode_})
	return row, err
}

// PutBlockContentRevision_bodyViaDelta Put BlockContentRevision_body via Delta
func PutBlockContentRevision_bodyViaDelta(Delta_ int, iBlockContentRevision_body *BlockContentRevision_body) (int64, error) {
	row, err := Engine.Update(iBlockContentRevision_body, &BlockContentRevision_body{Delta: Delta_})
	return row, err
}

// PutBlockContentRevision_bodyViaBodyValue Put BlockContentRevision_body via BodyValue
func PutBlockContentRevision_bodyViaBodyValue(BodyValue_ string, iBlockContentRevision_body *BlockContentRevision_body) (int64, error) {
	row, err := Engine.Update(iBlockContentRevision_body, &BlockContentRevision_body{BodyValue: BodyValue_})
	return row, err
}

// PutBlockContentRevision_bodyViaBodySummary Put BlockContentRevision_body via BodySummary
func PutBlockContentRevision_bodyViaBodySummary(BodySummary_ string, iBlockContentRevision_body *BlockContentRevision_body) (int64, error) {
	row, err := Engine.Update(iBlockContentRevision_body, &BlockContentRevision_body{BodySummary: BodySummary_})
	return row, err
}

// PutBlockContentRevision_bodyViaBodyFormat Put BlockContentRevision_body via BodyFormat
func PutBlockContentRevision_bodyViaBodyFormat(BodyFormat_ string, iBlockContentRevision_body *BlockContentRevision_body) (int64, error) {
	row, err := Engine.Update(iBlockContentRevision_body, &BlockContentRevision_body{BodyFormat: BodyFormat_})
	return row, err
}

// UpdateBlockContentRevision_bodyViaBundle via map[string]interface{}{}
func UpdateBlockContentRevision_bodyViaBundle(iBundle string, iBlockContentRevision_bodyMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(BlockContentRevision_body)).Where("bundle = ?", iBundle).Update(iBlockContentRevision_bodyMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateBlockContentRevision_bodyViaDeleted via map[string]interface{}{}
func UpdateBlockContentRevision_bodyViaDeleted(iDeleted int, iBlockContentRevision_bodyMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(BlockContentRevision_body)).Where("deleted = ?", iDeleted).Update(iBlockContentRevision_bodyMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateBlockContentRevision_bodyViaEntityId via map[string]interface{}{}
func UpdateBlockContentRevision_bodyViaEntityId(iEntityId int, iBlockContentRevision_bodyMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(BlockContentRevision_body)).Where("entity_id = ?", iEntityId).Update(iBlockContentRevision_bodyMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateBlockContentRevision_bodyViaRevisionId via map[string]interface{}{}
func UpdateBlockContentRevision_bodyViaRevisionId(iRevisionId int, iBlockContentRevision_bodyMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(BlockContentRevision_body)).Where("revision_id = ?", iRevisionId).Update(iBlockContentRevision_bodyMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateBlockContentRevision_bodyViaLangcode via map[string]interface{}{}
func UpdateBlockContentRevision_bodyViaLangcode(iLangcode string, iBlockContentRevision_bodyMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(BlockContentRevision_body)).Where("langcode = ?", iLangcode).Update(iBlockContentRevision_bodyMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateBlockContentRevision_bodyViaDelta via map[string]interface{}{}
func UpdateBlockContentRevision_bodyViaDelta(iDelta int, iBlockContentRevision_bodyMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(BlockContentRevision_body)).Where("delta = ?", iDelta).Update(iBlockContentRevision_bodyMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateBlockContentRevision_bodyViaBodyValue via map[string]interface{}{}
func UpdateBlockContentRevision_bodyViaBodyValue(iBodyValue string, iBlockContentRevision_bodyMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(BlockContentRevision_body)).Where("body_value = ?", iBodyValue).Update(iBlockContentRevision_bodyMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateBlockContentRevision_bodyViaBodySummary via map[string]interface{}{}
func UpdateBlockContentRevision_bodyViaBodySummary(iBodySummary string, iBlockContentRevision_bodyMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(BlockContentRevision_body)).Where("body_summary = ?", iBodySummary).Update(iBlockContentRevision_bodyMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateBlockContentRevision_bodyViaBodyFormat via map[string]interface{}{}
func UpdateBlockContentRevision_bodyViaBodyFormat(iBodyFormat string, iBlockContentRevision_bodyMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(BlockContentRevision_body)).Where("body_format = ?", iBodyFormat).Update(iBlockContentRevision_bodyMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// DeleteBlockContentRevision_body Delete BlockContentRevision_body
func DeleteBlockContentRevision_body(iBundle string) (int64, error) {
	row, err := Engine.Id(iBundle).Delete(new(BlockContentRevision_body))
	return row, err
}

// DeleteBlockContentRevision_bodyViaBundle Delete BlockContentRevision_body via Bundle
func DeleteBlockContentRevision_bodyViaBundle(iBundle string) (err error) {
	var has bool
	var _BlockContentRevision_body = &BlockContentRevision_body{Bundle: iBundle}
	if has, err = Engine.Get(_BlockContentRevision_body); (has == true) && (err == nil) {
		if row, err := Engine.Where("bundle = ?", iBundle).Delete(new(BlockContentRevision_body)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteBlockContentRevision_bodyViaDeleted Delete BlockContentRevision_body via Deleted
func DeleteBlockContentRevision_bodyViaDeleted(iDeleted int) (err error) {
	var has bool
	var _BlockContentRevision_body = &BlockContentRevision_body{Deleted: iDeleted}
	if has, err = Engine.Get(_BlockContentRevision_body); (has == true) && (err == nil) {
		if row, err := Engine.Where("deleted = ?", iDeleted).Delete(new(BlockContentRevision_body)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteBlockContentRevision_bodyViaEntityId Delete BlockContentRevision_body via EntityId
func DeleteBlockContentRevision_bodyViaEntityId(iEntityId int) (err error) {
	var has bool
	var _BlockContentRevision_body = &BlockContentRevision_body{EntityId: iEntityId}
	if has, err = Engine.Get(_BlockContentRevision_body); (has == true) && (err == nil) {
		if row, err := Engine.Where("entity_id = ?", iEntityId).Delete(new(BlockContentRevision_body)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteBlockContentRevision_bodyViaRevisionId Delete BlockContentRevision_body via RevisionId
func DeleteBlockContentRevision_bodyViaRevisionId(iRevisionId int) (err error) {
	var has bool
	var _BlockContentRevision_body = &BlockContentRevision_body{RevisionId: iRevisionId}
	if has, err = Engine.Get(_BlockContentRevision_body); (has == true) && (err == nil) {
		if row, err := Engine.Where("revision_id = ?", iRevisionId).Delete(new(BlockContentRevision_body)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteBlockContentRevision_bodyViaLangcode Delete BlockContentRevision_body via Langcode
func DeleteBlockContentRevision_bodyViaLangcode(iLangcode string) (err error) {
	var has bool
	var _BlockContentRevision_body = &BlockContentRevision_body{Langcode: iLangcode}
	if has, err = Engine.Get(_BlockContentRevision_body); (has == true) && (err == nil) {
		if row, err := Engine.Where("langcode = ?", iLangcode).Delete(new(BlockContentRevision_body)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteBlockContentRevision_bodyViaDelta Delete BlockContentRevision_body via Delta
func DeleteBlockContentRevision_bodyViaDelta(iDelta int) (err error) {
	var has bool
	var _BlockContentRevision_body = &BlockContentRevision_body{Delta: iDelta}
	if has, err = Engine.Get(_BlockContentRevision_body); (has == true) && (err == nil) {
		if row, err := Engine.Where("delta = ?", iDelta).Delete(new(BlockContentRevision_body)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteBlockContentRevision_bodyViaBodyValue Delete BlockContentRevision_body via BodyValue
func DeleteBlockContentRevision_bodyViaBodyValue(iBodyValue string) (err error) {
	var has bool
	var _BlockContentRevision_body = &BlockContentRevision_body{BodyValue: iBodyValue}
	if has, err = Engine.Get(_BlockContentRevision_body); (has == true) && (err == nil) {
		if row, err := Engine.Where("body_value = ?", iBodyValue).Delete(new(BlockContentRevision_body)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteBlockContentRevision_bodyViaBodySummary Delete BlockContentRevision_body via BodySummary
func DeleteBlockContentRevision_bodyViaBodySummary(iBodySummary string) (err error) {
	var has bool
	var _BlockContentRevision_body = &BlockContentRevision_body{BodySummary: iBodySummary}
	if has, err = Engine.Get(_BlockContentRevision_body); (has == true) && (err == nil) {
		if row, err := Engine.Where("body_summary = ?", iBodySummary).Delete(new(BlockContentRevision_body)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteBlockContentRevision_bodyViaBodyFormat Delete BlockContentRevision_body via BodyFormat
func DeleteBlockContentRevision_bodyViaBodyFormat(iBodyFormat string) (err error) {
	var has bool
	var _BlockContentRevision_body = &BlockContentRevision_body{BodyFormat: iBodyFormat}
	if has, err = Engine.Get(_BlockContentRevision_body); (has == true) && (err == nil) {
		if row, err := Engine.Where("body_format = ?", iBodyFormat).Delete(new(BlockContentRevision_body)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}
