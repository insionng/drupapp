package model

type BlockContent_body struct {
	Bundle      string `xorm:"not null default '' index VARCHAR(128)"`
	Deleted     int    `xorm:"not null pk default 0 TINYINT(4)"`
	EntityId    int    `xorm:"not null pk INT(10)"`
	RevisionId  int    `xorm:"not null index INT(10)"`
	Langcode    string `xorm:"not null pk default '' VARCHAR(32)"`
	Delta       int    `xorm:"not null pk INT(10)"`
	BodyValue   string `xorm:"not null LONGTEXT"`
	BodySummary string `xorm:"LONGTEXT"`
	BodyFormat  string `xorm:"index VARCHAR(255)"`
}

// GetBlockContent_bodiesCount BlockContent_bodys' Count
func GetBlockContent_bodiesCount(offset int, limit int) (int64, error) {
	total, err := Engine.Limit(limit, offset).Count(&BlockContent_body{})
	return total, err
}

// GetBlockContent_bodyCountViaBundle Get BlockContent_body via Bundle
func GetBlockContent_bodyCountViaBundle(iBundle string) int64 {
	n, _ := Engine.Where("bundle = ?", iBundle).Count(&BlockContent_body{Bundle: iBundle})
	return n
}

// GetBlockContent_bodyCountViaDeleted Get BlockContent_body via Deleted
func GetBlockContent_bodyCountViaDeleted(iDeleted int) int64 {
	n, _ := Engine.Where("deleted = ?", iDeleted).Count(&BlockContent_body{Deleted: iDeleted})
	return n
}

// GetBlockContent_bodyCountViaEntityId Get BlockContent_body via EntityId
func GetBlockContent_bodyCountViaEntityId(iEntityId int) int64 {
	n, _ := Engine.Where("entity_id = ?", iEntityId).Count(&BlockContent_body{EntityId: iEntityId})
	return n
}

// GetBlockContent_bodyCountViaRevisionId Get BlockContent_body via RevisionId
func GetBlockContent_bodyCountViaRevisionId(iRevisionId int) int64 {
	n, _ := Engine.Where("revision_id = ?", iRevisionId).Count(&BlockContent_body{RevisionId: iRevisionId})
	return n
}

// GetBlockContent_bodyCountViaLangcode Get BlockContent_body via Langcode
func GetBlockContent_bodyCountViaLangcode(iLangcode string) int64 {
	n, _ := Engine.Where("langcode = ?", iLangcode).Count(&BlockContent_body{Langcode: iLangcode})
	return n
}

// GetBlockContent_bodyCountViaDelta Get BlockContent_body via Delta
func GetBlockContent_bodyCountViaDelta(iDelta int) int64 {
	n, _ := Engine.Where("delta = ?", iDelta).Count(&BlockContent_body{Delta: iDelta})
	return n
}

// GetBlockContent_bodyCountViaBodyValue Get BlockContent_body via BodyValue
func GetBlockContent_bodyCountViaBodyValue(iBodyValue string) int64 {
	n, _ := Engine.Where("body_value = ?", iBodyValue).Count(&BlockContent_body{BodyValue: iBodyValue})
	return n
}

// GetBlockContent_bodyCountViaBodySummary Get BlockContent_body via BodySummary
func GetBlockContent_bodyCountViaBodySummary(iBodySummary string) int64 {
	n, _ := Engine.Where("body_summary = ?", iBodySummary).Count(&BlockContent_body{BodySummary: iBodySummary})
	return n
}

// GetBlockContent_bodyCountViaBodyFormat Get BlockContent_body via BodyFormat
func GetBlockContent_bodyCountViaBodyFormat(iBodyFormat string) int64 {
	n, _ := Engine.Where("body_format = ?", iBodyFormat).Count(&BlockContent_body{BodyFormat: iBodyFormat})
	return n
}

// GetBlockContent_bodiesByBundleAndDeletedAndEntityId Get BlockContent_bodies via BundleAndDeletedAndEntityId
func GetBlockContent_bodiesByBundleAndDeletedAndEntityId(offset int, limit int, Bundle_ string, Deleted_ int, EntityId_ int) (*[]*BlockContent_body, error) {
	var _BlockContent_body = new([]*BlockContent_body)
	err := Engine.Table("block_content__body").Where("bundle = ? and deleted = ? and entity_id = ?", Bundle_, Deleted_, EntityId_).Limit(limit, offset).Find(_BlockContent_body)
	return _BlockContent_body, err
}

// GetBlockContent_bodiesByBundleAndDeletedAndRevisionId Get BlockContent_bodies via BundleAndDeletedAndRevisionId
func GetBlockContent_bodiesByBundleAndDeletedAndRevisionId(offset int, limit int, Bundle_ string, Deleted_ int, RevisionId_ int) (*[]*BlockContent_body, error) {
	var _BlockContent_body = new([]*BlockContent_body)
	err := Engine.Table("block_content__body").Where("bundle = ? and deleted = ? and revision_id = ?", Bundle_, Deleted_, RevisionId_).Limit(limit, offset).Find(_BlockContent_body)
	return _BlockContent_body, err
}

// GetBlockContent_bodiesByBundleAndDeletedAndLangcode Get BlockContent_bodies via BundleAndDeletedAndLangcode
func GetBlockContent_bodiesByBundleAndDeletedAndLangcode(offset int, limit int, Bundle_ string, Deleted_ int, Langcode_ string) (*[]*BlockContent_body, error) {
	var _BlockContent_body = new([]*BlockContent_body)
	err := Engine.Table("block_content__body").Where("bundle = ? and deleted = ? and langcode = ?", Bundle_, Deleted_, Langcode_).Limit(limit, offset).Find(_BlockContent_body)
	return _BlockContent_body, err
}

// GetBlockContent_bodiesByBundleAndDeletedAndDelta Get BlockContent_bodies via BundleAndDeletedAndDelta
func GetBlockContent_bodiesByBundleAndDeletedAndDelta(offset int, limit int, Bundle_ string, Deleted_ int, Delta_ int) (*[]*BlockContent_body, error) {
	var _BlockContent_body = new([]*BlockContent_body)
	err := Engine.Table("block_content__body").Where("bundle = ? and deleted = ? and delta = ?", Bundle_, Deleted_, Delta_).Limit(limit, offset).Find(_BlockContent_body)
	return _BlockContent_body, err
}

// GetBlockContent_bodiesByBundleAndDeletedAndBodyValue Get BlockContent_bodies via BundleAndDeletedAndBodyValue
func GetBlockContent_bodiesByBundleAndDeletedAndBodyValue(offset int, limit int, Bundle_ string, Deleted_ int, BodyValue_ string) (*[]*BlockContent_body, error) {
	var _BlockContent_body = new([]*BlockContent_body)
	err := Engine.Table("block_content__body").Where("bundle = ? and deleted = ? and body_value = ?", Bundle_, Deleted_, BodyValue_).Limit(limit, offset).Find(_BlockContent_body)
	return _BlockContent_body, err
}

// GetBlockContent_bodiesByBundleAndDeletedAndBodySummary Get BlockContent_bodies via BundleAndDeletedAndBodySummary
func GetBlockContent_bodiesByBundleAndDeletedAndBodySummary(offset int, limit int, Bundle_ string, Deleted_ int, BodySummary_ string) (*[]*BlockContent_body, error) {
	var _BlockContent_body = new([]*BlockContent_body)
	err := Engine.Table("block_content__body").Where("bundle = ? and deleted = ? and body_summary = ?", Bundle_, Deleted_, BodySummary_).Limit(limit, offset).Find(_BlockContent_body)
	return _BlockContent_body, err
}

// GetBlockContent_bodiesByBundleAndDeletedAndBodyFormat Get BlockContent_bodies via BundleAndDeletedAndBodyFormat
func GetBlockContent_bodiesByBundleAndDeletedAndBodyFormat(offset int, limit int, Bundle_ string, Deleted_ int, BodyFormat_ string) (*[]*BlockContent_body, error) {
	var _BlockContent_body = new([]*BlockContent_body)
	err := Engine.Table("block_content__body").Where("bundle = ? and deleted = ? and body_format = ?", Bundle_, Deleted_, BodyFormat_).Limit(limit, offset).Find(_BlockContent_body)
	return _BlockContent_body, err
}

// GetBlockContent_bodiesByBundleAndEntityIdAndRevisionId Get BlockContent_bodies via BundleAndEntityIdAndRevisionId
func GetBlockContent_bodiesByBundleAndEntityIdAndRevisionId(offset int, limit int, Bundle_ string, EntityId_ int, RevisionId_ int) (*[]*BlockContent_body, error) {
	var _BlockContent_body = new([]*BlockContent_body)
	err := Engine.Table("block_content__body").Where("bundle = ? and entity_id = ? and revision_id = ?", Bundle_, EntityId_, RevisionId_).Limit(limit, offset).Find(_BlockContent_body)
	return _BlockContent_body, err
}

// GetBlockContent_bodiesByBundleAndEntityIdAndLangcode Get BlockContent_bodies via BundleAndEntityIdAndLangcode
func GetBlockContent_bodiesByBundleAndEntityIdAndLangcode(offset int, limit int, Bundle_ string, EntityId_ int, Langcode_ string) (*[]*BlockContent_body, error) {
	var _BlockContent_body = new([]*BlockContent_body)
	err := Engine.Table("block_content__body").Where("bundle = ? and entity_id = ? and langcode = ?", Bundle_, EntityId_, Langcode_).Limit(limit, offset).Find(_BlockContent_body)
	return _BlockContent_body, err
}

// GetBlockContent_bodiesByBundleAndEntityIdAndDelta Get BlockContent_bodies via BundleAndEntityIdAndDelta
func GetBlockContent_bodiesByBundleAndEntityIdAndDelta(offset int, limit int, Bundle_ string, EntityId_ int, Delta_ int) (*[]*BlockContent_body, error) {
	var _BlockContent_body = new([]*BlockContent_body)
	err := Engine.Table("block_content__body").Where("bundle = ? and entity_id = ? and delta = ?", Bundle_, EntityId_, Delta_).Limit(limit, offset).Find(_BlockContent_body)
	return _BlockContent_body, err
}

// GetBlockContent_bodiesByBundleAndEntityIdAndBodyValue Get BlockContent_bodies via BundleAndEntityIdAndBodyValue
func GetBlockContent_bodiesByBundleAndEntityIdAndBodyValue(offset int, limit int, Bundle_ string, EntityId_ int, BodyValue_ string) (*[]*BlockContent_body, error) {
	var _BlockContent_body = new([]*BlockContent_body)
	err := Engine.Table("block_content__body").Where("bundle = ? and entity_id = ? and body_value = ?", Bundle_, EntityId_, BodyValue_).Limit(limit, offset).Find(_BlockContent_body)
	return _BlockContent_body, err
}

// GetBlockContent_bodiesByBundleAndEntityIdAndBodySummary Get BlockContent_bodies via BundleAndEntityIdAndBodySummary
func GetBlockContent_bodiesByBundleAndEntityIdAndBodySummary(offset int, limit int, Bundle_ string, EntityId_ int, BodySummary_ string) (*[]*BlockContent_body, error) {
	var _BlockContent_body = new([]*BlockContent_body)
	err := Engine.Table("block_content__body").Where("bundle = ? and entity_id = ? and body_summary = ?", Bundle_, EntityId_, BodySummary_).Limit(limit, offset).Find(_BlockContent_body)
	return _BlockContent_body, err
}

// GetBlockContent_bodiesByBundleAndEntityIdAndBodyFormat Get BlockContent_bodies via BundleAndEntityIdAndBodyFormat
func GetBlockContent_bodiesByBundleAndEntityIdAndBodyFormat(offset int, limit int, Bundle_ string, EntityId_ int, BodyFormat_ string) (*[]*BlockContent_body, error) {
	var _BlockContent_body = new([]*BlockContent_body)
	err := Engine.Table("block_content__body").Where("bundle = ? and entity_id = ? and body_format = ?", Bundle_, EntityId_, BodyFormat_).Limit(limit, offset).Find(_BlockContent_body)
	return _BlockContent_body, err
}

// GetBlockContent_bodiesByBundleAndRevisionIdAndLangcode Get BlockContent_bodies via BundleAndRevisionIdAndLangcode
func GetBlockContent_bodiesByBundleAndRevisionIdAndLangcode(offset int, limit int, Bundle_ string, RevisionId_ int, Langcode_ string) (*[]*BlockContent_body, error) {
	var _BlockContent_body = new([]*BlockContent_body)
	err := Engine.Table("block_content__body").Where("bundle = ? and revision_id = ? and langcode = ?", Bundle_, RevisionId_, Langcode_).Limit(limit, offset).Find(_BlockContent_body)
	return _BlockContent_body, err
}

// GetBlockContent_bodiesByBundleAndRevisionIdAndDelta Get BlockContent_bodies via BundleAndRevisionIdAndDelta
func GetBlockContent_bodiesByBundleAndRevisionIdAndDelta(offset int, limit int, Bundle_ string, RevisionId_ int, Delta_ int) (*[]*BlockContent_body, error) {
	var _BlockContent_body = new([]*BlockContent_body)
	err := Engine.Table("block_content__body").Where("bundle = ? and revision_id = ? and delta = ?", Bundle_, RevisionId_, Delta_).Limit(limit, offset).Find(_BlockContent_body)
	return _BlockContent_body, err
}

// GetBlockContent_bodiesByBundleAndRevisionIdAndBodyValue Get BlockContent_bodies via BundleAndRevisionIdAndBodyValue
func GetBlockContent_bodiesByBundleAndRevisionIdAndBodyValue(offset int, limit int, Bundle_ string, RevisionId_ int, BodyValue_ string) (*[]*BlockContent_body, error) {
	var _BlockContent_body = new([]*BlockContent_body)
	err := Engine.Table("block_content__body").Where("bundle = ? and revision_id = ? and body_value = ?", Bundle_, RevisionId_, BodyValue_).Limit(limit, offset).Find(_BlockContent_body)
	return _BlockContent_body, err
}

// GetBlockContent_bodiesByBundleAndRevisionIdAndBodySummary Get BlockContent_bodies via BundleAndRevisionIdAndBodySummary
func GetBlockContent_bodiesByBundleAndRevisionIdAndBodySummary(offset int, limit int, Bundle_ string, RevisionId_ int, BodySummary_ string) (*[]*BlockContent_body, error) {
	var _BlockContent_body = new([]*BlockContent_body)
	err := Engine.Table("block_content__body").Where("bundle = ? and revision_id = ? and body_summary = ?", Bundle_, RevisionId_, BodySummary_).Limit(limit, offset).Find(_BlockContent_body)
	return _BlockContent_body, err
}

// GetBlockContent_bodiesByBundleAndRevisionIdAndBodyFormat Get BlockContent_bodies via BundleAndRevisionIdAndBodyFormat
func GetBlockContent_bodiesByBundleAndRevisionIdAndBodyFormat(offset int, limit int, Bundle_ string, RevisionId_ int, BodyFormat_ string) (*[]*BlockContent_body, error) {
	var _BlockContent_body = new([]*BlockContent_body)
	err := Engine.Table("block_content__body").Where("bundle = ? and revision_id = ? and body_format = ?", Bundle_, RevisionId_, BodyFormat_).Limit(limit, offset).Find(_BlockContent_body)
	return _BlockContent_body, err
}

// GetBlockContent_bodiesByBundleAndLangcodeAndDelta Get BlockContent_bodies via BundleAndLangcodeAndDelta
func GetBlockContent_bodiesByBundleAndLangcodeAndDelta(offset int, limit int, Bundle_ string, Langcode_ string, Delta_ int) (*[]*BlockContent_body, error) {
	var _BlockContent_body = new([]*BlockContent_body)
	err := Engine.Table("block_content__body").Where("bundle = ? and langcode = ? and delta = ?", Bundle_, Langcode_, Delta_).Limit(limit, offset).Find(_BlockContent_body)
	return _BlockContent_body, err
}

// GetBlockContent_bodiesByBundleAndLangcodeAndBodyValue Get BlockContent_bodies via BundleAndLangcodeAndBodyValue
func GetBlockContent_bodiesByBundleAndLangcodeAndBodyValue(offset int, limit int, Bundle_ string, Langcode_ string, BodyValue_ string) (*[]*BlockContent_body, error) {
	var _BlockContent_body = new([]*BlockContent_body)
	err := Engine.Table("block_content__body").Where("bundle = ? and langcode = ? and body_value = ?", Bundle_, Langcode_, BodyValue_).Limit(limit, offset).Find(_BlockContent_body)
	return _BlockContent_body, err
}

// GetBlockContent_bodiesByBundleAndLangcodeAndBodySummary Get BlockContent_bodies via BundleAndLangcodeAndBodySummary
func GetBlockContent_bodiesByBundleAndLangcodeAndBodySummary(offset int, limit int, Bundle_ string, Langcode_ string, BodySummary_ string) (*[]*BlockContent_body, error) {
	var _BlockContent_body = new([]*BlockContent_body)
	err := Engine.Table("block_content__body").Where("bundle = ? and langcode = ? and body_summary = ?", Bundle_, Langcode_, BodySummary_).Limit(limit, offset).Find(_BlockContent_body)
	return _BlockContent_body, err
}

// GetBlockContent_bodiesByBundleAndLangcodeAndBodyFormat Get BlockContent_bodies via BundleAndLangcodeAndBodyFormat
func GetBlockContent_bodiesByBundleAndLangcodeAndBodyFormat(offset int, limit int, Bundle_ string, Langcode_ string, BodyFormat_ string) (*[]*BlockContent_body, error) {
	var _BlockContent_body = new([]*BlockContent_body)
	err := Engine.Table("block_content__body").Where("bundle = ? and langcode = ? and body_format = ?", Bundle_, Langcode_, BodyFormat_).Limit(limit, offset).Find(_BlockContent_body)
	return _BlockContent_body, err
}

// GetBlockContent_bodiesByBundleAndDeltaAndBodyValue Get BlockContent_bodies via BundleAndDeltaAndBodyValue
func GetBlockContent_bodiesByBundleAndDeltaAndBodyValue(offset int, limit int, Bundle_ string, Delta_ int, BodyValue_ string) (*[]*BlockContent_body, error) {
	var _BlockContent_body = new([]*BlockContent_body)
	err := Engine.Table("block_content__body").Where("bundle = ? and delta = ? and body_value = ?", Bundle_, Delta_, BodyValue_).Limit(limit, offset).Find(_BlockContent_body)
	return _BlockContent_body, err
}

// GetBlockContent_bodiesByBundleAndDeltaAndBodySummary Get BlockContent_bodies via BundleAndDeltaAndBodySummary
func GetBlockContent_bodiesByBundleAndDeltaAndBodySummary(offset int, limit int, Bundle_ string, Delta_ int, BodySummary_ string) (*[]*BlockContent_body, error) {
	var _BlockContent_body = new([]*BlockContent_body)
	err := Engine.Table("block_content__body").Where("bundle = ? and delta = ? and body_summary = ?", Bundle_, Delta_, BodySummary_).Limit(limit, offset).Find(_BlockContent_body)
	return _BlockContent_body, err
}

// GetBlockContent_bodiesByBundleAndDeltaAndBodyFormat Get BlockContent_bodies via BundleAndDeltaAndBodyFormat
func GetBlockContent_bodiesByBundleAndDeltaAndBodyFormat(offset int, limit int, Bundle_ string, Delta_ int, BodyFormat_ string) (*[]*BlockContent_body, error) {
	var _BlockContent_body = new([]*BlockContent_body)
	err := Engine.Table("block_content__body").Where("bundle = ? and delta = ? and body_format = ?", Bundle_, Delta_, BodyFormat_).Limit(limit, offset).Find(_BlockContent_body)
	return _BlockContent_body, err
}

// GetBlockContent_bodiesByBundleAndBodyValueAndBodySummary Get BlockContent_bodies via BundleAndBodyValueAndBodySummary
func GetBlockContent_bodiesByBundleAndBodyValueAndBodySummary(offset int, limit int, Bundle_ string, BodyValue_ string, BodySummary_ string) (*[]*BlockContent_body, error) {
	var _BlockContent_body = new([]*BlockContent_body)
	err := Engine.Table("block_content__body").Where("bundle = ? and body_value = ? and body_summary = ?", Bundle_, BodyValue_, BodySummary_).Limit(limit, offset).Find(_BlockContent_body)
	return _BlockContent_body, err
}

// GetBlockContent_bodiesByBundleAndBodyValueAndBodyFormat Get BlockContent_bodies via BundleAndBodyValueAndBodyFormat
func GetBlockContent_bodiesByBundleAndBodyValueAndBodyFormat(offset int, limit int, Bundle_ string, BodyValue_ string, BodyFormat_ string) (*[]*BlockContent_body, error) {
	var _BlockContent_body = new([]*BlockContent_body)
	err := Engine.Table("block_content__body").Where("bundle = ? and body_value = ? and body_format = ?", Bundle_, BodyValue_, BodyFormat_).Limit(limit, offset).Find(_BlockContent_body)
	return _BlockContent_body, err
}

// GetBlockContent_bodiesByBundleAndBodySummaryAndBodyFormat Get BlockContent_bodies via BundleAndBodySummaryAndBodyFormat
func GetBlockContent_bodiesByBundleAndBodySummaryAndBodyFormat(offset int, limit int, Bundle_ string, BodySummary_ string, BodyFormat_ string) (*[]*BlockContent_body, error) {
	var _BlockContent_body = new([]*BlockContent_body)
	err := Engine.Table("block_content__body").Where("bundle = ? and body_summary = ? and body_format = ?", Bundle_, BodySummary_, BodyFormat_).Limit(limit, offset).Find(_BlockContent_body)
	return _BlockContent_body, err
}

// GetBlockContent_bodiesByDeletedAndEntityIdAndRevisionId Get BlockContent_bodies via DeletedAndEntityIdAndRevisionId
func GetBlockContent_bodiesByDeletedAndEntityIdAndRevisionId(offset int, limit int, Deleted_ int, EntityId_ int, RevisionId_ int) (*[]*BlockContent_body, error) {
	var _BlockContent_body = new([]*BlockContent_body)
	err := Engine.Table("block_content__body").Where("deleted = ? and entity_id = ? and revision_id = ?", Deleted_, EntityId_, RevisionId_).Limit(limit, offset).Find(_BlockContent_body)
	return _BlockContent_body, err
}

// GetBlockContent_bodiesByDeletedAndEntityIdAndLangcode Get BlockContent_bodies via DeletedAndEntityIdAndLangcode
func GetBlockContent_bodiesByDeletedAndEntityIdAndLangcode(offset int, limit int, Deleted_ int, EntityId_ int, Langcode_ string) (*[]*BlockContent_body, error) {
	var _BlockContent_body = new([]*BlockContent_body)
	err := Engine.Table("block_content__body").Where("deleted = ? and entity_id = ? and langcode = ?", Deleted_, EntityId_, Langcode_).Limit(limit, offset).Find(_BlockContent_body)
	return _BlockContent_body, err
}

// GetBlockContent_bodiesByDeletedAndEntityIdAndDelta Get BlockContent_bodies via DeletedAndEntityIdAndDelta
func GetBlockContent_bodiesByDeletedAndEntityIdAndDelta(offset int, limit int, Deleted_ int, EntityId_ int, Delta_ int) (*[]*BlockContent_body, error) {
	var _BlockContent_body = new([]*BlockContent_body)
	err := Engine.Table("block_content__body").Where("deleted = ? and entity_id = ? and delta = ?", Deleted_, EntityId_, Delta_).Limit(limit, offset).Find(_BlockContent_body)
	return _BlockContent_body, err
}

// GetBlockContent_bodiesByDeletedAndEntityIdAndBodyValue Get BlockContent_bodies via DeletedAndEntityIdAndBodyValue
func GetBlockContent_bodiesByDeletedAndEntityIdAndBodyValue(offset int, limit int, Deleted_ int, EntityId_ int, BodyValue_ string) (*[]*BlockContent_body, error) {
	var _BlockContent_body = new([]*BlockContent_body)
	err := Engine.Table("block_content__body").Where("deleted = ? and entity_id = ? and body_value = ?", Deleted_, EntityId_, BodyValue_).Limit(limit, offset).Find(_BlockContent_body)
	return _BlockContent_body, err
}

// GetBlockContent_bodiesByDeletedAndEntityIdAndBodySummary Get BlockContent_bodies via DeletedAndEntityIdAndBodySummary
func GetBlockContent_bodiesByDeletedAndEntityIdAndBodySummary(offset int, limit int, Deleted_ int, EntityId_ int, BodySummary_ string) (*[]*BlockContent_body, error) {
	var _BlockContent_body = new([]*BlockContent_body)
	err := Engine.Table("block_content__body").Where("deleted = ? and entity_id = ? and body_summary = ?", Deleted_, EntityId_, BodySummary_).Limit(limit, offset).Find(_BlockContent_body)
	return _BlockContent_body, err
}

// GetBlockContent_bodiesByDeletedAndEntityIdAndBodyFormat Get BlockContent_bodies via DeletedAndEntityIdAndBodyFormat
func GetBlockContent_bodiesByDeletedAndEntityIdAndBodyFormat(offset int, limit int, Deleted_ int, EntityId_ int, BodyFormat_ string) (*[]*BlockContent_body, error) {
	var _BlockContent_body = new([]*BlockContent_body)
	err := Engine.Table("block_content__body").Where("deleted = ? and entity_id = ? and body_format = ?", Deleted_, EntityId_, BodyFormat_).Limit(limit, offset).Find(_BlockContent_body)
	return _BlockContent_body, err
}

// GetBlockContent_bodiesByDeletedAndRevisionIdAndLangcode Get BlockContent_bodies via DeletedAndRevisionIdAndLangcode
func GetBlockContent_bodiesByDeletedAndRevisionIdAndLangcode(offset int, limit int, Deleted_ int, RevisionId_ int, Langcode_ string) (*[]*BlockContent_body, error) {
	var _BlockContent_body = new([]*BlockContent_body)
	err := Engine.Table("block_content__body").Where("deleted = ? and revision_id = ? and langcode = ?", Deleted_, RevisionId_, Langcode_).Limit(limit, offset).Find(_BlockContent_body)
	return _BlockContent_body, err
}

// GetBlockContent_bodiesByDeletedAndRevisionIdAndDelta Get BlockContent_bodies via DeletedAndRevisionIdAndDelta
func GetBlockContent_bodiesByDeletedAndRevisionIdAndDelta(offset int, limit int, Deleted_ int, RevisionId_ int, Delta_ int) (*[]*BlockContent_body, error) {
	var _BlockContent_body = new([]*BlockContent_body)
	err := Engine.Table("block_content__body").Where("deleted = ? and revision_id = ? and delta = ?", Deleted_, RevisionId_, Delta_).Limit(limit, offset).Find(_BlockContent_body)
	return _BlockContent_body, err
}

// GetBlockContent_bodiesByDeletedAndRevisionIdAndBodyValue Get BlockContent_bodies via DeletedAndRevisionIdAndBodyValue
func GetBlockContent_bodiesByDeletedAndRevisionIdAndBodyValue(offset int, limit int, Deleted_ int, RevisionId_ int, BodyValue_ string) (*[]*BlockContent_body, error) {
	var _BlockContent_body = new([]*BlockContent_body)
	err := Engine.Table("block_content__body").Where("deleted = ? and revision_id = ? and body_value = ?", Deleted_, RevisionId_, BodyValue_).Limit(limit, offset).Find(_BlockContent_body)
	return _BlockContent_body, err
}

// GetBlockContent_bodiesByDeletedAndRevisionIdAndBodySummary Get BlockContent_bodies via DeletedAndRevisionIdAndBodySummary
func GetBlockContent_bodiesByDeletedAndRevisionIdAndBodySummary(offset int, limit int, Deleted_ int, RevisionId_ int, BodySummary_ string) (*[]*BlockContent_body, error) {
	var _BlockContent_body = new([]*BlockContent_body)
	err := Engine.Table("block_content__body").Where("deleted = ? and revision_id = ? and body_summary = ?", Deleted_, RevisionId_, BodySummary_).Limit(limit, offset).Find(_BlockContent_body)
	return _BlockContent_body, err
}

// GetBlockContent_bodiesByDeletedAndRevisionIdAndBodyFormat Get BlockContent_bodies via DeletedAndRevisionIdAndBodyFormat
func GetBlockContent_bodiesByDeletedAndRevisionIdAndBodyFormat(offset int, limit int, Deleted_ int, RevisionId_ int, BodyFormat_ string) (*[]*BlockContent_body, error) {
	var _BlockContent_body = new([]*BlockContent_body)
	err := Engine.Table("block_content__body").Where("deleted = ? and revision_id = ? and body_format = ?", Deleted_, RevisionId_, BodyFormat_).Limit(limit, offset).Find(_BlockContent_body)
	return _BlockContent_body, err
}

// GetBlockContent_bodiesByDeletedAndLangcodeAndDelta Get BlockContent_bodies via DeletedAndLangcodeAndDelta
func GetBlockContent_bodiesByDeletedAndLangcodeAndDelta(offset int, limit int, Deleted_ int, Langcode_ string, Delta_ int) (*[]*BlockContent_body, error) {
	var _BlockContent_body = new([]*BlockContent_body)
	err := Engine.Table("block_content__body").Where("deleted = ? and langcode = ? and delta = ?", Deleted_, Langcode_, Delta_).Limit(limit, offset).Find(_BlockContent_body)
	return _BlockContent_body, err
}

// GetBlockContent_bodiesByDeletedAndLangcodeAndBodyValue Get BlockContent_bodies via DeletedAndLangcodeAndBodyValue
func GetBlockContent_bodiesByDeletedAndLangcodeAndBodyValue(offset int, limit int, Deleted_ int, Langcode_ string, BodyValue_ string) (*[]*BlockContent_body, error) {
	var _BlockContent_body = new([]*BlockContent_body)
	err := Engine.Table("block_content__body").Where("deleted = ? and langcode = ? and body_value = ?", Deleted_, Langcode_, BodyValue_).Limit(limit, offset).Find(_BlockContent_body)
	return _BlockContent_body, err
}

// GetBlockContent_bodiesByDeletedAndLangcodeAndBodySummary Get BlockContent_bodies via DeletedAndLangcodeAndBodySummary
func GetBlockContent_bodiesByDeletedAndLangcodeAndBodySummary(offset int, limit int, Deleted_ int, Langcode_ string, BodySummary_ string) (*[]*BlockContent_body, error) {
	var _BlockContent_body = new([]*BlockContent_body)
	err := Engine.Table("block_content__body").Where("deleted = ? and langcode = ? and body_summary = ?", Deleted_, Langcode_, BodySummary_).Limit(limit, offset).Find(_BlockContent_body)
	return _BlockContent_body, err
}

// GetBlockContent_bodiesByDeletedAndLangcodeAndBodyFormat Get BlockContent_bodies via DeletedAndLangcodeAndBodyFormat
func GetBlockContent_bodiesByDeletedAndLangcodeAndBodyFormat(offset int, limit int, Deleted_ int, Langcode_ string, BodyFormat_ string) (*[]*BlockContent_body, error) {
	var _BlockContent_body = new([]*BlockContent_body)
	err := Engine.Table("block_content__body").Where("deleted = ? and langcode = ? and body_format = ?", Deleted_, Langcode_, BodyFormat_).Limit(limit, offset).Find(_BlockContent_body)
	return _BlockContent_body, err
}

// GetBlockContent_bodiesByDeletedAndDeltaAndBodyValue Get BlockContent_bodies via DeletedAndDeltaAndBodyValue
func GetBlockContent_bodiesByDeletedAndDeltaAndBodyValue(offset int, limit int, Deleted_ int, Delta_ int, BodyValue_ string) (*[]*BlockContent_body, error) {
	var _BlockContent_body = new([]*BlockContent_body)
	err := Engine.Table("block_content__body").Where("deleted = ? and delta = ? and body_value = ?", Deleted_, Delta_, BodyValue_).Limit(limit, offset).Find(_BlockContent_body)
	return _BlockContent_body, err
}

// GetBlockContent_bodiesByDeletedAndDeltaAndBodySummary Get BlockContent_bodies via DeletedAndDeltaAndBodySummary
func GetBlockContent_bodiesByDeletedAndDeltaAndBodySummary(offset int, limit int, Deleted_ int, Delta_ int, BodySummary_ string) (*[]*BlockContent_body, error) {
	var _BlockContent_body = new([]*BlockContent_body)
	err := Engine.Table("block_content__body").Where("deleted = ? and delta = ? and body_summary = ?", Deleted_, Delta_, BodySummary_).Limit(limit, offset).Find(_BlockContent_body)
	return _BlockContent_body, err
}

// GetBlockContent_bodiesByDeletedAndDeltaAndBodyFormat Get BlockContent_bodies via DeletedAndDeltaAndBodyFormat
func GetBlockContent_bodiesByDeletedAndDeltaAndBodyFormat(offset int, limit int, Deleted_ int, Delta_ int, BodyFormat_ string) (*[]*BlockContent_body, error) {
	var _BlockContent_body = new([]*BlockContent_body)
	err := Engine.Table("block_content__body").Where("deleted = ? and delta = ? and body_format = ?", Deleted_, Delta_, BodyFormat_).Limit(limit, offset).Find(_BlockContent_body)
	return _BlockContent_body, err
}

// GetBlockContent_bodiesByDeletedAndBodyValueAndBodySummary Get BlockContent_bodies via DeletedAndBodyValueAndBodySummary
func GetBlockContent_bodiesByDeletedAndBodyValueAndBodySummary(offset int, limit int, Deleted_ int, BodyValue_ string, BodySummary_ string) (*[]*BlockContent_body, error) {
	var _BlockContent_body = new([]*BlockContent_body)
	err := Engine.Table("block_content__body").Where("deleted = ? and body_value = ? and body_summary = ?", Deleted_, BodyValue_, BodySummary_).Limit(limit, offset).Find(_BlockContent_body)
	return _BlockContent_body, err
}

// GetBlockContent_bodiesByDeletedAndBodyValueAndBodyFormat Get BlockContent_bodies via DeletedAndBodyValueAndBodyFormat
func GetBlockContent_bodiesByDeletedAndBodyValueAndBodyFormat(offset int, limit int, Deleted_ int, BodyValue_ string, BodyFormat_ string) (*[]*BlockContent_body, error) {
	var _BlockContent_body = new([]*BlockContent_body)
	err := Engine.Table("block_content__body").Where("deleted = ? and body_value = ? and body_format = ?", Deleted_, BodyValue_, BodyFormat_).Limit(limit, offset).Find(_BlockContent_body)
	return _BlockContent_body, err
}

// GetBlockContent_bodiesByDeletedAndBodySummaryAndBodyFormat Get BlockContent_bodies via DeletedAndBodySummaryAndBodyFormat
func GetBlockContent_bodiesByDeletedAndBodySummaryAndBodyFormat(offset int, limit int, Deleted_ int, BodySummary_ string, BodyFormat_ string) (*[]*BlockContent_body, error) {
	var _BlockContent_body = new([]*BlockContent_body)
	err := Engine.Table("block_content__body").Where("deleted = ? and body_summary = ? and body_format = ?", Deleted_, BodySummary_, BodyFormat_).Limit(limit, offset).Find(_BlockContent_body)
	return _BlockContent_body, err
}

// GetBlockContent_bodiesByEntityIdAndRevisionIdAndLangcode Get BlockContent_bodies via EntityIdAndRevisionIdAndLangcode
func GetBlockContent_bodiesByEntityIdAndRevisionIdAndLangcode(offset int, limit int, EntityId_ int, RevisionId_ int, Langcode_ string) (*[]*BlockContent_body, error) {
	var _BlockContent_body = new([]*BlockContent_body)
	err := Engine.Table("block_content__body").Where("entity_id = ? and revision_id = ? and langcode = ?", EntityId_, RevisionId_, Langcode_).Limit(limit, offset).Find(_BlockContent_body)
	return _BlockContent_body, err
}

// GetBlockContent_bodiesByEntityIdAndRevisionIdAndDelta Get BlockContent_bodies via EntityIdAndRevisionIdAndDelta
func GetBlockContent_bodiesByEntityIdAndRevisionIdAndDelta(offset int, limit int, EntityId_ int, RevisionId_ int, Delta_ int) (*[]*BlockContent_body, error) {
	var _BlockContent_body = new([]*BlockContent_body)
	err := Engine.Table("block_content__body").Where("entity_id = ? and revision_id = ? and delta = ?", EntityId_, RevisionId_, Delta_).Limit(limit, offset).Find(_BlockContent_body)
	return _BlockContent_body, err
}

// GetBlockContent_bodiesByEntityIdAndRevisionIdAndBodyValue Get BlockContent_bodies via EntityIdAndRevisionIdAndBodyValue
func GetBlockContent_bodiesByEntityIdAndRevisionIdAndBodyValue(offset int, limit int, EntityId_ int, RevisionId_ int, BodyValue_ string) (*[]*BlockContent_body, error) {
	var _BlockContent_body = new([]*BlockContent_body)
	err := Engine.Table("block_content__body").Where("entity_id = ? and revision_id = ? and body_value = ?", EntityId_, RevisionId_, BodyValue_).Limit(limit, offset).Find(_BlockContent_body)
	return _BlockContent_body, err
}

// GetBlockContent_bodiesByEntityIdAndRevisionIdAndBodySummary Get BlockContent_bodies via EntityIdAndRevisionIdAndBodySummary
func GetBlockContent_bodiesByEntityIdAndRevisionIdAndBodySummary(offset int, limit int, EntityId_ int, RevisionId_ int, BodySummary_ string) (*[]*BlockContent_body, error) {
	var _BlockContent_body = new([]*BlockContent_body)
	err := Engine.Table("block_content__body").Where("entity_id = ? and revision_id = ? and body_summary = ?", EntityId_, RevisionId_, BodySummary_).Limit(limit, offset).Find(_BlockContent_body)
	return _BlockContent_body, err
}

// GetBlockContent_bodiesByEntityIdAndRevisionIdAndBodyFormat Get BlockContent_bodies via EntityIdAndRevisionIdAndBodyFormat
func GetBlockContent_bodiesByEntityIdAndRevisionIdAndBodyFormat(offset int, limit int, EntityId_ int, RevisionId_ int, BodyFormat_ string) (*[]*BlockContent_body, error) {
	var _BlockContent_body = new([]*BlockContent_body)
	err := Engine.Table("block_content__body").Where("entity_id = ? and revision_id = ? and body_format = ?", EntityId_, RevisionId_, BodyFormat_).Limit(limit, offset).Find(_BlockContent_body)
	return _BlockContent_body, err
}

// GetBlockContent_bodiesByEntityIdAndLangcodeAndDelta Get BlockContent_bodies via EntityIdAndLangcodeAndDelta
func GetBlockContent_bodiesByEntityIdAndLangcodeAndDelta(offset int, limit int, EntityId_ int, Langcode_ string, Delta_ int) (*[]*BlockContent_body, error) {
	var _BlockContent_body = new([]*BlockContent_body)
	err := Engine.Table("block_content__body").Where("entity_id = ? and langcode = ? and delta = ?", EntityId_, Langcode_, Delta_).Limit(limit, offset).Find(_BlockContent_body)
	return _BlockContent_body, err
}

// GetBlockContent_bodiesByEntityIdAndLangcodeAndBodyValue Get BlockContent_bodies via EntityIdAndLangcodeAndBodyValue
func GetBlockContent_bodiesByEntityIdAndLangcodeAndBodyValue(offset int, limit int, EntityId_ int, Langcode_ string, BodyValue_ string) (*[]*BlockContent_body, error) {
	var _BlockContent_body = new([]*BlockContent_body)
	err := Engine.Table("block_content__body").Where("entity_id = ? and langcode = ? and body_value = ?", EntityId_, Langcode_, BodyValue_).Limit(limit, offset).Find(_BlockContent_body)
	return _BlockContent_body, err
}

// GetBlockContent_bodiesByEntityIdAndLangcodeAndBodySummary Get BlockContent_bodies via EntityIdAndLangcodeAndBodySummary
func GetBlockContent_bodiesByEntityIdAndLangcodeAndBodySummary(offset int, limit int, EntityId_ int, Langcode_ string, BodySummary_ string) (*[]*BlockContent_body, error) {
	var _BlockContent_body = new([]*BlockContent_body)
	err := Engine.Table("block_content__body").Where("entity_id = ? and langcode = ? and body_summary = ?", EntityId_, Langcode_, BodySummary_).Limit(limit, offset).Find(_BlockContent_body)
	return _BlockContent_body, err
}

// GetBlockContent_bodiesByEntityIdAndLangcodeAndBodyFormat Get BlockContent_bodies via EntityIdAndLangcodeAndBodyFormat
func GetBlockContent_bodiesByEntityIdAndLangcodeAndBodyFormat(offset int, limit int, EntityId_ int, Langcode_ string, BodyFormat_ string) (*[]*BlockContent_body, error) {
	var _BlockContent_body = new([]*BlockContent_body)
	err := Engine.Table("block_content__body").Where("entity_id = ? and langcode = ? and body_format = ?", EntityId_, Langcode_, BodyFormat_).Limit(limit, offset).Find(_BlockContent_body)
	return _BlockContent_body, err
}

// GetBlockContent_bodiesByEntityIdAndDeltaAndBodyValue Get BlockContent_bodies via EntityIdAndDeltaAndBodyValue
func GetBlockContent_bodiesByEntityIdAndDeltaAndBodyValue(offset int, limit int, EntityId_ int, Delta_ int, BodyValue_ string) (*[]*BlockContent_body, error) {
	var _BlockContent_body = new([]*BlockContent_body)
	err := Engine.Table("block_content__body").Where("entity_id = ? and delta = ? and body_value = ?", EntityId_, Delta_, BodyValue_).Limit(limit, offset).Find(_BlockContent_body)
	return _BlockContent_body, err
}

// GetBlockContent_bodiesByEntityIdAndDeltaAndBodySummary Get BlockContent_bodies via EntityIdAndDeltaAndBodySummary
func GetBlockContent_bodiesByEntityIdAndDeltaAndBodySummary(offset int, limit int, EntityId_ int, Delta_ int, BodySummary_ string) (*[]*BlockContent_body, error) {
	var _BlockContent_body = new([]*BlockContent_body)
	err := Engine.Table("block_content__body").Where("entity_id = ? and delta = ? and body_summary = ?", EntityId_, Delta_, BodySummary_).Limit(limit, offset).Find(_BlockContent_body)
	return _BlockContent_body, err
}

// GetBlockContent_bodiesByEntityIdAndDeltaAndBodyFormat Get BlockContent_bodies via EntityIdAndDeltaAndBodyFormat
func GetBlockContent_bodiesByEntityIdAndDeltaAndBodyFormat(offset int, limit int, EntityId_ int, Delta_ int, BodyFormat_ string) (*[]*BlockContent_body, error) {
	var _BlockContent_body = new([]*BlockContent_body)
	err := Engine.Table("block_content__body").Where("entity_id = ? and delta = ? and body_format = ?", EntityId_, Delta_, BodyFormat_).Limit(limit, offset).Find(_BlockContent_body)
	return _BlockContent_body, err
}

// GetBlockContent_bodiesByEntityIdAndBodyValueAndBodySummary Get BlockContent_bodies via EntityIdAndBodyValueAndBodySummary
func GetBlockContent_bodiesByEntityIdAndBodyValueAndBodySummary(offset int, limit int, EntityId_ int, BodyValue_ string, BodySummary_ string) (*[]*BlockContent_body, error) {
	var _BlockContent_body = new([]*BlockContent_body)
	err := Engine.Table("block_content__body").Where("entity_id = ? and body_value = ? and body_summary = ?", EntityId_, BodyValue_, BodySummary_).Limit(limit, offset).Find(_BlockContent_body)
	return _BlockContent_body, err
}

// GetBlockContent_bodiesByEntityIdAndBodyValueAndBodyFormat Get BlockContent_bodies via EntityIdAndBodyValueAndBodyFormat
func GetBlockContent_bodiesByEntityIdAndBodyValueAndBodyFormat(offset int, limit int, EntityId_ int, BodyValue_ string, BodyFormat_ string) (*[]*BlockContent_body, error) {
	var _BlockContent_body = new([]*BlockContent_body)
	err := Engine.Table("block_content__body").Where("entity_id = ? and body_value = ? and body_format = ?", EntityId_, BodyValue_, BodyFormat_).Limit(limit, offset).Find(_BlockContent_body)
	return _BlockContent_body, err
}

// GetBlockContent_bodiesByEntityIdAndBodySummaryAndBodyFormat Get BlockContent_bodies via EntityIdAndBodySummaryAndBodyFormat
func GetBlockContent_bodiesByEntityIdAndBodySummaryAndBodyFormat(offset int, limit int, EntityId_ int, BodySummary_ string, BodyFormat_ string) (*[]*BlockContent_body, error) {
	var _BlockContent_body = new([]*BlockContent_body)
	err := Engine.Table("block_content__body").Where("entity_id = ? and body_summary = ? and body_format = ?", EntityId_, BodySummary_, BodyFormat_).Limit(limit, offset).Find(_BlockContent_body)
	return _BlockContent_body, err
}

// GetBlockContent_bodiesByRevisionIdAndLangcodeAndDelta Get BlockContent_bodies via RevisionIdAndLangcodeAndDelta
func GetBlockContent_bodiesByRevisionIdAndLangcodeAndDelta(offset int, limit int, RevisionId_ int, Langcode_ string, Delta_ int) (*[]*BlockContent_body, error) {
	var _BlockContent_body = new([]*BlockContent_body)
	err := Engine.Table("block_content__body").Where("revision_id = ? and langcode = ? and delta = ?", RevisionId_, Langcode_, Delta_).Limit(limit, offset).Find(_BlockContent_body)
	return _BlockContent_body, err
}

// GetBlockContent_bodiesByRevisionIdAndLangcodeAndBodyValue Get BlockContent_bodies via RevisionIdAndLangcodeAndBodyValue
func GetBlockContent_bodiesByRevisionIdAndLangcodeAndBodyValue(offset int, limit int, RevisionId_ int, Langcode_ string, BodyValue_ string) (*[]*BlockContent_body, error) {
	var _BlockContent_body = new([]*BlockContent_body)
	err := Engine.Table("block_content__body").Where("revision_id = ? and langcode = ? and body_value = ?", RevisionId_, Langcode_, BodyValue_).Limit(limit, offset).Find(_BlockContent_body)
	return _BlockContent_body, err
}

// GetBlockContent_bodiesByRevisionIdAndLangcodeAndBodySummary Get BlockContent_bodies via RevisionIdAndLangcodeAndBodySummary
func GetBlockContent_bodiesByRevisionIdAndLangcodeAndBodySummary(offset int, limit int, RevisionId_ int, Langcode_ string, BodySummary_ string) (*[]*BlockContent_body, error) {
	var _BlockContent_body = new([]*BlockContent_body)
	err := Engine.Table("block_content__body").Where("revision_id = ? and langcode = ? and body_summary = ?", RevisionId_, Langcode_, BodySummary_).Limit(limit, offset).Find(_BlockContent_body)
	return _BlockContent_body, err
}

// GetBlockContent_bodiesByRevisionIdAndLangcodeAndBodyFormat Get BlockContent_bodies via RevisionIdAndLangcodeAndBodyFormat
func GetBlockContent_bodiesByRevisionIdAndLangcodeAndBodyFormat(offset int, limit int, RevisionId_ int, Langcode_ string, BodyFormat_ string) (*[]*BlockContent_body, error) {
	var _BlockContent_body = new([]*BlockContent_body)
	err := Engine.Table("block_content__body").Where("revision_id = ? and langcode = ? and body_format = ?", RevisionId_, Langcode_, BodyFormat_).Limit(limit, offset).Find(_BlockContent_body)
	return _BlockContent_body, err
}

// GetBlockContent_bodiesByRevisionIdAndDeltaAndBodyValue Get BlockContent_bodies via RevisionIdAndDeltaAndBodyValue
func GetBlockContent_bodiesByRevisionIdAndDeltaAndBodyValue(offset int, limit int, RevisionId_ int, Delta_ int, BodyValue_ string) (*[]*BlockContent_body, error) {
	var _BlockContent_body = new([]*BlockContent_body)
	err := Engine.Table("block_content__body").Where("revision_id = ? and delta = ? and body_value = ?", RevisionId_, Delta_, BodyValue_).Limit(limit, offset).Find(_BlockContent_body)
	return _BlockContent_body, err
}

// GetBlockContent_bodiesByRevisionIdAndDeltaAndBodySummary Get BlockContent_bodies via RevisionIdAndDeltaAndBodySummary
func GetBlockContent_bodiesByRevisionIdAndDeltaAndBodySummary(offset int, limit int, RevisionId_ int, Delta_ int, BodySummary_ string) (*[]*BlockContent_body, error) {
	var _BlockContent_body = new([]*BlockContent_body)
	err := Engine.Table("block_content__body").Where("revision_id = ? and delta = ? and body_summary = ?", RevisionId_, Delta_, BodySummary_).Limit(limit, offset).Find(_BlockContent_body)
	return _BlockContent_body, err
}

// GetBlockContent_bodiesByRevisionIdAndDeltaAndBodyFormat Get BlockContent_bodies via RevisionIdAndDeltaAndBodyFormat
func GetBlockContent_bodiesByRevisionIdAndDeltaAndBodyFormat(offset int, limit int, RevisionId_ int, Delta_ int, BodyFormat_ string) (*[]*BlockContent_body, error) {
	var _BlockContent_body = new([]*BlockContent_body)
	err := Engine.Table("block_content__body").Where("revision_id = ? and delta = ? and body_format = ?", RevisionId_, Delta_, BodyFormat_).Limit(limit, offset).Find(_BlockContent_body)
	return _BlockContent_body, err
}

// GetBlockContent_bodiesByRevisionIdAndBodyValueAndBodySummary Get BlockContent_bodies via RevisionIdAndBodyValueAndBodySummary
func GetBlockContent_bodiesByRevisionIdAndBodyValueAndBodySummary(offset int, limit int, RevisionId_ int, BodyValue_ string, BodySummary_ string) (*[]*BlockContent_body, error) {
	var _BlockContent_body = new([]*BlockContent_body)
	err := Engine.Table("block_content__body").Where("revision_id = ? and body_value = ? and body_summary = ?", RevisionId_, BodyValue_, BodySummary_).Limit(limit, offset).Find(_BlockContent_body)
	return _BlockContent_body, err
}

// GetBlockContent_bodiesByRevisionIdAndBodyValueAndBodyFormat Get BlockContent_bodies via RevisionIdAndBodyValueAndBodyFormat
func GetBlockContent_bodiesByRevisionIdAndBodyValueAndBodyFormat(offset int, limit int, RevisionId_ int, BodyValue_ string, BodyFormat_ string) (*[]*BlockContent_body, error) {
	var _BlockContent_body = new([]*BlockContent_body)
	err := Engine.Table("block_content__body").Where("revision_id = ? and body_value = ? and body_format = ?", RevisionId_, BodyValue_, BodyFormat_).Limit(limit, offset).Find(_BlockContent_body)
	return _BlockContent_body, err
}

// GetBlockContent_bodiesByRevisionIdAndBodySummaryAndBodyFormat Get BlockContent_bodies via RevisionIdAndBodySummaryAndBodyFormat
func GetBlockContent_bodiesByRevisionIdAndBodySummaryAndBodyFormat(offset int, limit int, RevisionId_ int, BodySummary_ string, BodyFormat_ string) (*[]*BlockContent_body, error) {
	var _BlockContent_body = new([]*BlockContent_body)
	err := Engine.Table("block_content__body").Where("revision_id = ? and body_summary = ? and body_format = ?", RevisionId_, BodySummary_, BodyFormat_).Limit(limit, offset).Find(_BlockContent_body)
	return _BlockContent_body, err
}

// GetBlockContent_bodiesByLangcodeAndDeltaAndBodyValue Get BlockContent_bodies via LangcodeAndDeltaAndBodyValue
func GetBlockContent_bodiesByLangcodeAndDeltaAndBodyValue(offset int, limit int, Langcode_ string, Delta_ int, BodyValue_ string) (*[]*BlockContent_body, error) {
	var _BlockContent_body = new([]*BlockContent_body)
	err := Engine.Table("block_content__body").Where("langcode = ? and delta = ? and body_value = ?", Langcode_, Delta_, BodyValue_).Limit(limit, offset).Find(_BlockContent_body)
	return _BlockContent_body, err
}

// GetBlockContent_bodiesByLangcodeAndDeltaAndBodySummary Get BlockContent_bodies via LangcodeAndDeltaAndBodySummary
func GetBlockContent_bodiesByLangcodeAndDeltaAndBodySummary(offset int, limit int, Langcode_ string, Delta_ int, BodySummary_ string) (*[]*BlockContent_body, error) {
	var _BlockContent_body = new([]*BlockContent_body)
	err := Engine.Table("block_content__body").Where("langcode = ? and delta = ? and body_summary = ?", Langcode_, Delta_, BodySummary_).Limit(limit, offset).Find(_BlockContent_body)
	return _BlockContent_body, err
}

// GetBlockContent_bodiesByLangcodeAndDeltaAndBodyFormat Get BlockContent_bodies via LangcodeAndDeltaAndBodyFormat
func GetBlockContent_bodiesByLangcodeAndDeltaAndBodyFormat(offset int, limit int, Langcode_ string, Delta_ int, BodyFormat_ string) (*[]*BlockContent_body, error) {
	var _BlockContent_body = new([]*BlockContent_body)
	err := Engine.Table("block_content__body").Where("langcode = ? and delta = ? and body_format = ?", Langcode_, Delta_, BodyFormat_).Limit(limit, offset).Find(_BlockContent_body)
	return _BlockContent_body, err
}

// GetBlockContent_bodiesByLangcodeAndBodyValueAndBodySummary Get BlockContent_bodies via LangcodeAndBodyValueAndBodySummary
func GetBlockContent_bodiesByLangcodeAndBodyValueAndBodySummary(offset int, limit int, Langcode_ string, BodyValue_ string, BodySummary_ string) (*[]*BlockContent_body, error) {
	var _BlockContent_body = new([]*BlockContent_body)
	err := Engine.Table("block_content__body").Where("langcode = ? and body_value = ? and body_summary = ?", Langcode_, BodyValue_, BodySummary_).Limit(limit, offset).Find(_BlockContent_body)
	return _BlockContent_body, err
}

// GetBlockContent_bodiesByLangcodeAndBodyValueAndBodyFormat Get BlockContent_bodies via LangcodeAndBodyValueAndBodyFormat
func GetBlockContent_bodiesByLangcodeAndBodyValueAndBodyFormat(offset int, limit int, Langcode_ string, BodyValue_ string, BodyFormat_ string) (*[]*BlockContent_body, error) {
	var _BlockContent_body = new([]*BlockContent_body)
	err := Engine.Table("block_content__body").Where("langcode = ? and body_value = ? and body_format = ?", Langcode_, BodyValue_, BodyFormat_).Limit(limit, offset).Find(_BlockContent_body)
	return _BlockContent_body, err
}

// GetBlockContent_bodiesByLangcodeAndBodySummaryAndBodyFormat Get BlockContent_bodies via LangcodeAndBodySummaryAndBodyFormat
func GetBlockContent_bodiesByLangcodeAndBodySummaryAndBodyFormat(offset int, limit int, Langcode_ string, BodySummary_ string, BodyFormat_ string) (*[]*BlockContent_body, error) {
	var _BlockContent_body = new([]*BlockContent_body)
	err := Engine.Table("block_content__body").Where("langcode = ? and body_summary = ? and body_format = ?", Langcode_, BodySummary_, BodyFormat_).Limit(limit, offset).Find(_BlockContent_body)
	return _BlockContent_body, err
}

// GetBlockContent_bodiesByDeltaAndBodyValueAndBodySummary Get BlockContent_bodies via DeltaAndBodyValueAndBodySummary
func GetBlockContent_bodiesByDeltaAndBodyValueAndBodySummary(offset int, limit int, Delta_ int, BodyValue_ string, BodySummary_ string) (*[]*BlockContent_body, error) {
	var _BlockContent_body = new([]*BlockContent_body)
	err := Engine.Table("block_content__body").Where("delta = ? and body_value = ? and body_summary = ?", Delta_, BodyValue_, BodySummary_).Limit(limit, offset).Find(_BlockContent_body)
	return _BlockContent_body, err
}

// GetBlockContent_bodiesByDeltaAndBodyValueAndBodyFormat Get BlockContent_bodies via DeltaAndBodyValueAndBodyFormat
func GetBlockContent_bodiesByDeltaAndBodyValueAndBodyFormat(offset int, limit int, Delta_ int, BodyValue_ string, BodyFormat_ string) (*[]*BlockContent_body, error) {
	var _BlockContent_body = new([]*BlockContent_body)
	err := Engine.Table("block_content__body").Where("delta = ? and body_value = ? and body_format = ?", Delta_, BodyValue_, BodyFormat_).Limit(limit, offset).Find(_BlockContent_body)
	return _BlockContent_body, err
}

// GetBlockContent_bodiesByDeltaAndBodySummaryAndBodyFormat Get BlockContent_bodies via DeltaAndBodySummaryAndBodyFormat
func GetBlockContent_bodiesByDeltaAndBodySummaryAndBodyFormat(offset int, limit int, Delta_ int, BodySummary_ string, BodyFormat_ string) (*[]*BlockContent_body, error) {
	var _BlockContent_body = new([]*BlockContent_body)
	err := Engine.Table("block_content__body").Where("delta = ? and body_summary = ? and body_format = ?", Delta_, BodySummary_, BodyFormat_).Limit(limit, offset).Find(_BlockContent_body)
	return _BlockContent_body, err
}

// GetBlockContent_bodiesByBodyValueAndBodySummaryAndBodyFormat Get BlockContent_bodies via BodyValueAndBodySummaryAndBodyFormat
func GetBlockContent_bodiesByBodyValueAndBodySummaryAndBodyFormat(offset int, limit int, BodyValue_ string, BodySummary_ string, BodyFormat_ string) (*[]*BlockContent_body, error) {
	var _BlockContent_body = new([]*BlockContent_body)
	err := Engine.Table("block_content__body").Where("body_value = ? and body_summary = ? and body_format = ?", BodyValue_, BodySummary_, BodyFormat_).Limit(limit, offset).Find(_BlockContent_body)
	return _BlockContent_body, err
}

// GetBlockContent_bodiesByBundleAndDeleted Get BlockContent_bodies via BundleAndDeleted
func GetBlockContent_bodiesByBundleAndDeleted(offset int, limit int, Bundle_ string, Deleted_ int) (*[]*BlockContent_body, error) {
	var _BlockContent_body = new([]*BlockContent_body)
	err := Engine.Table("block_content__body").Where("bundle = ? and deleted = ?", Bundle_, Deleted_).Limit(limit, offset).Find(_BlockContent_body)
	return _BlockContent_body, err
}

// GetBlockContent_bodiesByBundleAndEntityId Get BlockContent_bodies via BundleAndEntityId
func GetBlockContent_bodiesByBundleAndEntityId(offset int, limit int, Bundle_ string, EntityId_ int) (*[]*BlockContent_body, error) {
	var _BlockContent_body = new([]*BlockContent_body)
	err := Engine.Table("block_content__body").Where("bundle = ? and entity_id = ?", Bundle_, EntityId_).Limit(limit, offset).Find(_BlockContent_body)
	return _BlockContent_body, err
}

// GetBlockContent_bodiesByBundleAndRevisionId Get BlockContent_bodies via BundleAndRevisionId
func GetBlockContent_bodiesByBundleAndRevisionId(offset int, limit int, Bundle_ string, RevisionId_ int) (*[]*BlockContent_body, error) {
	var _BlockContent_body = new([]*BlockContent_body)
	err := Engine.Table("block_content__body").Where("bundle = ? and revision_id = ?", Bundle_, RevisionId_).Limit(limit, offset).Find(_BlockContent_body)
	return _BlockContent_body, err
}

// GetBlockContent_bodiesByBundleAndLangcode Get BlockContent_bodies via BundleAndLangcode
func GetBlockContent_bodiesByBundleAndLangcode(offset int, limit int, Bundle_ string, Langcode_ string) (*[]*BlockContent_body, error) {
	var _BlockContent_body = new([]*BlockContent_body)
	err := Engine.Table("block_content__body").Where("bundle = ? and langcode = ?", Bundle_, Langcode_).Limit(limit, offset).Find(_BlockContent_body)
	return _BlockContent_body, err
}

// GetBlockContent_bodiesByBundleAndDelta Get BlockContent_bodies via BundleAndDelta
func GetBlockContent_bodiesByBundleAndDelta(offset int, limit int, Bundle_ string, Delta_ int) (*[]*BlockContent_body, error) {
	var _BlockContent_body = new([]*BlockContent_body)
	err := Engine.Table("block_content__body").Where("bundle = ? and delta = ?", Bundle_, Delta_).Limit(limit, offset).Find(_BlockContent_body)
	return _BlockContent_body, err
}

// GetBlockContent_bodiesByBundleAndBodyValue Get BlockContent_bodies via BundleAndBodyValue
func GetBlockContent_bodiesByBundleAndBodyValue(offset int, limit int, Bundle_ string, BodyValue_ string) (*[]*BlockContent_body, error) {
	var _BlockContent_body = new([]*BlockContent_body)
	err := Engine.Table("block_content__body").Where("bundle = ? and body_value = ?", Bundle_, BodyValue_).Limit(limit, offset).Find(_BlockContent_body)
	return _BlockContent_body, err
}

// GetBlockContent_bodiesByBundleAndBodySummary Get BlockContent_bodies via BundleAndBodySummary
func GetBlockContent_bodiesByBundleAndBodySummary(offset int, limit int, Bundle_ string, BodySummary_ string) (*[]*BlockContent_body, error) {
	var _BlockContent_body = new([]*BlockContent_body)
	err := Engine.Table("block_content__body").Where("bundle = ? and body_summary = ?", Bundle_, BodySummary_).Limit(limit, offset).Find(_BlockContent_body)
	return _BlockContent_body, err
}

// GetBlockContent_bodiesByBundleAndBodyFormat Get BlockContent_bodies via BundleAndBodyFormat
func GetBlockContent_bodiesByBundleAndBodyFormat(offset int, limit int, Bundle_ string, BodyFormat_ string) (*[]*BlockContent_body, error) {
	var _BlockContent_body = new([]*BlockContent_body)
	err := Engine.Table("block_content__body").Where("bundle = ? and body_format = ?", Bundle_, BodyFormat_).Limit(limit, offset).Find(_BlockContent_body)
	return _BlockContent_body, err
}

// GetBlockContent_bodiesByDeletedAndEntityId Get BlockContent_bodies via DeletedAndEntityId
func GetBlockContent_bodiesByDeletedAndEntityId(offset int, limit int, Deleted_ int, EntityId_ int) (*[]*BlockContent_body, error) {
	var _BlockContent_body = new([]*BlockContent_body)
	err := Engine.Table("block_content__body").Where("deleted = ? and entity_id = ?", Deleted_, EntityId_).Limit(limit, offset).Find(_BlockContent_body)
	return _BlockContent_body, err
}

// GetBlockContent_bodiesByDeletedAndRevisionId Get BlockContent_bodies via DeletedAndRevisionId
func GetBlockContent_bodiesByDeletedAndRevisionId(offset int, limit int, Deleted_ int, RevisionId_ int) (*[]*BlockContent_body, error) {
	var _BlockContent_body = new([]*BlockContent_body)
	err := Engine.Table("block_content__body").Where("deleted = ? and revision_id = ?", Deleted_, RevisionId_).Limit(limit, offset).Find(_BlockContent_body)
	return _BlockContent_body, err
}

// GetBlockContent_bodiesByDeletedAndLangcode Get BlockContent_bodies via DeletedAndLangcode
func GetBlockContent_bodiesByDeletedAndLangcode(offset int, limit int, Deleted_ int, Langcode_ string) (*[]*BlockContent_body, error) {
	var _BlockContent_body = new([]*BlockContent_body)
	err := Engine.Table("block_content__body").Where("deleted = ? and langcode = ?", Deleted_, Langcode_).Limit(limit, offset).Find(_BlockContent_body)
	return _BlockContent_body, err
}

// GetBlockContent_bodiesByDeletedAndDelta Get BlockContent_bodies via DeletedAndDelta
func GetBlockContent_bodiesByDeletedAndDelta(offset int, limit int, Deleted_ int, Delta_ int) (*[]*BlockContent_body, error) {
	var _BlockContent_body = new([]*BlockContent_body)
	err := Engine.Table("block_content__body").Where("deleted = ? and delta = ?", Deleted_, Delta_).Limit(limit, offset).Find(_BlockContent_body)
	return _BlockContent_body, err
}

// GetBlockContent_bodiesByDeletedAndBodyValue Get BlockContent_bodies via DeletedAndBodyValue
func GetBlockContent_bodiesByDeletedAndBodyValue(offset int, limit int, Deleted_ int, BodyValue_ string) (*[]*BlockContent_body, error) {
	var _BlockContent_body = new([]*BlockContent_body)
	err := Engine.Table("block_content__body").Where("deleted = ? and body_value = ?", Deleted_, BodyValue_).Limit(limit, offset).Find(_BlockContent_body)
	return _BlockContent_body, err
}

// GetBlockContent_bodiesByDeletedAndBodySummary Get BlockContent_bodies via DeletedAndBodySummary
func GetBlockContent_bodiesByDeletedAndBodySummary(offset int, limit int, Deleted_ int, BodySummary_ string) (*[]*BlockContent_body, error) {
	var _BlockContent_body = new([]*BlockContent_body)
	err := Engine.Table("block_content__body").Where("deleted = ? and body_summary = ?", Deleted_, BodySummary_).Limit(limit, offset).Find(_BlockContent_body)
	return _BlockContent_body, err
}

// GetBlockContent_bodiesByDeletedAndBodyFormat Get BlockContent_bodies via DeletedAndBodyFormat
func GetBlockContent_bodiesByDeletedAndBodyFormat(offset int, limit int, Deleted_ int, BodyFormat_ string) (*[]*BlockContent_body, error) {
	var _BlockContent_body = new([]*BlockContent_body)
	err := Engine.Table("block_content__body").Where("deleted = ? and body_format = ?", Deleted_, BodyFormat_).Limit(limit, offset).Find(_BlockContent_body)
	return _BlockContent_body, err
}

// GetBlockContent_bodiesByEntityIdAndRevisionId Get BlockContent_bodies via EntityIdAndRevisionId
func GetBlockContent_bodiesByEntityIdAndRevisionId(offset int, limit int, EntityId_ int, RevisionId_ int) (*[]*BlockContent_body, error) {
	var _BlockContent_body = new([]*BlockContent_body)
	err := Engine.Table("block_content__body").Where("entity_id = ? and revision_id = ?", EntityId_, RevisionId_).Limit(limit, offset).Find(_BlockContent_body)
	return _BlockContent_body, err
}

// GetBlockContent_bodiesByEntityIdAndLangcode Get BlockContent_bodies via EntityIdAndLangcode
func GetBlockContent_bodiesByEntityIdAndLangcode(offset int, limit int, EntityId_ int, Langcode_ string) (*[]*BlockContent_body, error) {
	var _BlockContent_body = new([]*BlockContent_body)
	err := Engine.Table("block_content__body").Where("entity_id = ? and langcode = ?", EntityId_, Langcode_).Limit(limit, offset).Find(_BlockContent_body)
	return _BlockContent_body, err
}

// GetBlockContent_bodiesByEntityIdAndDelta Get BlockContent_bodies via EntityIdAndDelta
func GetBlockContent_bodiesByEntityIdAndDelta(offset int, limit int, EntityId_ int, Delta_ int) (*[]*BlockContent_body, error) {
	var _BlockContent_body = new([]*BlockContent_body)
	err := Engine.Table("block_content__body").Where("entity_id = ? and delta = ?", EntityId_, Delta_).Limit(limit, offset).Find(_BlockContent_body)
	return _BlockContent_body, err
}

// GetBlockContent_bodiesByEntityIdAndBodyValue Get BlockContent_bodies via EntityIdAndBodyValue
func GetBlockContent_bodiesByEntityIdAndBodyValue(offset int, limit int, EntityId_ int, BodyValue_ string) (*[]*BlockContent_body, error) {
	var _BlockContent_body = new([]*BlockContent_body)
	err := Engine.Table("block_content__body").Where("entity_id = ? and body_value = ?", EntityId_, BodyValue_).Limit(limit, offset).Find(_BlockContent_body)
	return _BlockContent_body, err
}

// GetBlockContent_bodiesByEntityIdAndBodySummary Get BlockContent_bodies via EntityIdAndBodySummary
func GetBlockContent_bodiesByEntityIdAndBodySummary(offset int, limit int, EntityId_ int, BodySummary_ string) (*[]*BlockContent_body, error) {
	var _BlockContent_body = new([]*BlockContent_body)
	err := Engine.Table("block_content__body").Where("entity_id = ? and body_summary = ?", EntityId_, BodySummary_).Limit(limit, offset).Find(_BlockContent_body)
	return _BlockContent_body, err
}

// GetBlockContent_bodiesByEntityIdAndBodyFormat Get BlockContent_bodies via EntityIdAndBodyFormat
func GetBlockContent_bodiesByEntityIdAndBodyFormat(offset int, limit int, EntityId_ int, BodyFormat_ string) (*[]*BlockContent_body, error) {
	var _BlockContent_body = new([]*BlockContent_body)
	err := Engine.Table("block_content__body").Where("entity_id = ? and body_format = ?", EntityId_, BodyFormat_).Limit(limit, offset).Find(_BlockContent_body)
	return _BlockContent_body, err
}

// GetBlockContent_bodiesByRevisionIdAndLangcode Get BlockContent_bodies via RevisionIdAndLangcode
func GetBlockContent_bodiesByRevisionIdAndLangcode(offset int, limit int, RevisionId_ int, Langcode_ string) (*[]*BlockContent_body, error) {
	var _BlockContent_body = new([]*BlockContent_body)
	err := Engine.Table("block_content__body").Where("revision_id = ? and langcode = ?", RevisionId_, Langcode_).Limit(limit, offset).Find(_BlockContent_body)
	return _BlockContent_body, err
}

// GetBlockContent_bodiesByRevisionIdAndDelta Get BlockContent_bodies via RevisionIdAndDelta
func GetBlockContent_bodiesByRevisionIdAndDelta(offset int, limit int, RevisionId_ int, Delta_ int) (*[]*BlockContent_body, error) {
	var _BlockContent_body = new([]*BlockContent_body)
	err := Engine.Table("block_content__body").Where("revision_id = ? and delta = ?", RevisionId_, Delta_).Limit(limit, offset).Find(_BlockContent_body)
	return _BlockContent_body, err
}

// GetBlockContent_bodiesByRevisionIdAndBodyValue Get BlockContent_bodies via RevisionIdAndBodyValue
func GetBlockContent_bodiesByRevisionIdAndBodyValue(offset int, limit int, RevisionId_ int, BodyValue_ string) (*[]*BlockContent_body, error) {
	var _BlockContent_body = new([]*BlockContent_body)
	err := Engine.Table("block_content__body").Where("revision_id = ? and body_value = ?", RevisionId_, BodyValue_).Limit(limit, offset).Find(_BlockContent_body)
	return _BlockContent_body, err
}

// GetBlockContent_bodiesByRevisionIdAndBodySummary Get BlockContent_bodies via RevisionIdAndBodySummary
func GetBlockContent_bodiesByRevisionIdAndBodySummary(offset int, limit int, RevisionId_ int, BodySummary_ string) (*[]*BlockContent_body, error) {
	var _BlockContent_body = new([]*BlockContent_body)
	err := Engine.Table("block_content__body").Where("revision_id = ? and body_summary = ?", RevisionId_, BodySummary_).Limit(limit, offset).Find(_BlockContent_body)
	return _BlockContent_body, err
}

// GetBlockContent_bodiesByRevisionIdAndBodyFormat Get BlockContent_bodies via RevisionIdAndBodyFormat
func GetBlockContent_bodiesByRevisionIdAndBodyFormat(offset int, limit int, RevisionId_ int, BodyFormat_ string) (*[]*BlockContent_body, error) {
	var _BlockContent_body = new([]*BlockContent_body)
	err := Engine.Table("block_content__body").Where("revision_id = ? and body_format = ?", RevisionId_, BodyFormat_).Limit(limit, offset).Find(_BlockContent_body)
	return _BlockContent_body, err
}

// GetBlockContent_bodiesByLangcodeAndDelta Get BlockContent_bodies via LangcodeAndDelta
func GetBlockContent_bodiesByLangcodeAndDelta(offset int, limit int, Langcode_ string, Delta_ int) (*[]*BlockContent_body, error) {
	var _BlockContent_body = new([]*BlockContent_body)
	err := Engine.Table("block_content__body").Where("langcode = ? and delta = ?", Langcode_, Delta_).Limit(limit, offset).Find(_BlockContent_body)
	return _BlockContent_body, err
}

// GetBlockContent_bodiesByLangcodeAndBodyValue Get BlockContent_bodies via LangcodeAndBodyValue
func GetBlockContent_bodiesByLangcodeAndBodyValue(offset int, limit int, Langcode_ string, BodyValue_ string) (*[]*BlockContent_body, error) {
	var _BlockContent_body = new([]*BlockContent_body)
	err := Engine.Table("block_content__body").Where("langcode = ? and body_value = ?", Langcode_, BodyValue_).Limit(limit, offset).Find(_BlockContent_body)
	return _BlockContent_body, err
}

// GetBlockContent_bodiesByLangcodeAndBodySummary Get BlockContent_bodies via LangcodeAndBodySummary
func GetBlockContent_bodiesByLangcodeAndBodySummary(offset int, limit int, Langcode_ string, BodySummary_ string) (*[]*BlockContent_body, error) {
	var _BlockContent_body = new([]*BlockContent_body)
	err := Engine.Table("block_content__body").Where("langcode = ? and body_summary = ?", Langcode_, BodySummary_).Limit(limit, offset).Find(_BlockContent_body)
	return _BlockContent_body, err
}

// GetBlockContent_bodiesByLangcodeAndBodyFormat Get BlockContent_bodies via LangcodeAndBodyFormat
func GetBlockContent_bodiesByLangcodeAndBodyFormat(offset int, limit int, Langcode_ string, BodyFormat_ string) (*[]*BlockContent_body, error) {
	var _BlockContent_body = new([]*BlockContent_body)
	err := Engine.Table("block_content__body").Where("langcode = ? and body_format = ?", Langcode_, BodyFormat_).Limit(limit, offset).Find(_BlockContent_body)
	return _BlockContent_body, err
}

// GetBlockContent_bodiesByDeltaAndBodyValue Get BlockContent_bodies via DeltaAndBodyValue
func GetBlockContent_bodiesByDeltaAndBodyValue(offset int, limit int, Delta_ int, BodyValue_ string) (*[]*BlockContent_body, error) {
	var _BlockContent_body = new([]*BlockContent_body)
	err := Engine.Table("block_content__body").Where("delta = ? and body_value = ?", Delta_, BodyValue_).Limit(limit, offset).Find(_BlockContent_body)
	return _BlockContent_body, err
}

// GetBlockContent_bodiesByDeltaAndBodySummary Get BlockContent_bodies via DeltaAndBodySummary
func GetBlockContent_bodiesByDeltaAndBodySummary(offset int, limit int, Delta_ int, BodySummary_ string) (*[]*BlockContent_body, error) {
	var _BlockContent_body = new([]*BlockContent_body)
	err := Engine.Table("block_content__body").Where("delta = ? and body_summary = ?", Delta_, BodySummary_).Limit(limit, offset).Find(_BlockContent_body)
	return _BlockContent_body, err
}

// GetBlockContent_bodiesByDeltaAndBodyFormat Get BlockContent_bodies via DeltaAndBodyFormat
func GetBlockContent_bodiesByDeltaAndBodyFormat(offset int, limit int, Delta_ int, BodyFormat_ string) (*[]*BlockContent_body, error) {
	var _BlockContent_body = new([]*BlockContent_body)
	err := Engine.Table("block_content__body").Where("delta = ? and body_format = ?", Delta_, BodyFormat_).Limit(limit, offset).Find(_BlockContent_body)
	return _BlockContent_body, err
}

// GetBlockContent_bodiesByBodyValueAndBodySummary Get BlockContent_bodies via BodyValueAndBodySummary
func GetBlockContent_bodiesByBodyValueAndBodySummary(offset int, limit int, BodyValue_ string, BodySummary_ string) (*[]*BlockContent_body, error) {
	var _BlockContent_body = new([]*BlockContent_body)
	err := Engine.Table("block_content__body").Where("body_value = ? and body_summary = ?", BodyValue_, BodySummary_).Limit(limit, offset).Find(_BlockContent_body)
	return _BlockContent_body, err
}

// GetBlockContent_bodiesByBodyValueAndBodyFormat Get BlockContent_bodies via BodyValueAndBodyFormat
func GetBlockContent_bodiesByBodyValueAndBodyFormat(offset int, limit int, BodyValue_ string, BodyFormat_ string) (*[]*BlockContent_body, error) {
	var _BlockContent_body = new([]*BlockContent_body)
	err := Engine.Table("block_content__body").Where("body_value = ? and body_format = ?", BodyValue_, BodyFormat_).Limit(limit, offset).Find(_BlockContent_body)
	return _BlockContent_body, err
}

// GetBlockContent_bodiesByBodySummaryAndBodyFormat Get BlockContent_bodies via BodySummaryAndBodyFormat
func GetBlockContent_bodiesByBodySummaryAndBodyFormat(offset int, limit int, BodySummary_ string, BodyFormat_ string) (*[]*BlockContent_body, error) {
	var _BlockContent_body = new([]*BlockContent_body)
	err := Engine.Table("block_content__body").Where("body_summary = ? and body_format = ?", BodySummary_, BodyFormat_).Limit(limit, offset).Find(_BlockContent_body)
	return _BlockContent_body, err
}

// GetBlockContent_bodies Get BlockContent_bodies via field
func GetBlockContent_bodies(offset int, limit int, field string) (*[]*BlockContent_body, error) {
	var _BlockContent_body = new([]*BlockContent_body)
	err := Engine.Table("block_content__body").Limit(limit, offset).Desc(field).Find(_BlockContent_body)
	return _BlockContent_body, err
}

// GetBlockContent_bodiesViaBundle Get BlockContent_bodys via Bundle
func GetBlockContent_bodiesViaBundle(offset int, limit int, Bundle_ string, field string) (*[]*BlockContent_body, error) {
	var _BlockContent_body = new([]*BlockContent_body)
	err := Engine.Table("block_content__body").Where("bundle = ?", Bundle_).Limit(limit, offset).Desc(field).Find(_BlockContent_body)
	return _BlockContent_body, err
}

// GetBlockContent_bodiesViaDeleted Get BlockContent_bodys via Deleted
func GetBlockContent_bodiesViaDeleted(offset int, limit int, Deleted_ int, field string) (*[]*BlockContent_body, error) {
	var _BlockContent_body = new([]*BlockContent_body)
	err := Engine.Table("block_content__body").Where("deleted = ?", Deleted_).Limit(limit, offset).Desc(field).Find(_BlockContent_body)
	return _BlockContent_body, err
}

// GetBlockContent_bodiesViaEntityId Get BlockContent_bodys via EntityId
func GetBlockContent_bodiesViaEntityId(offset int, limit int, EntityId_ int, field string) (*[]*BlockContent_body, error) {
	var _BlockContent_body = new([]*BlockContent_body)
	err := Engine.Table("block_content__body").Where("entity_id = ?", EntityId_).Limit(limit, offset).Desc(field).Find(_BlockContent_body)
	return _BlockContent_body, err
}

// GetBlockContent_bodiesViaRevisionId Get BlockContent_bodys via RevisionId
func GetBlockContent_bodiesViaRevisionId(offset int, limit int, RevisionId_ int, field string) (*[]*BlockContent_body, error) {
	var _BlockContent_body = new([]*BlockContent_body)
	err := Engine.Table("block_content__body").Where("revision_id = ?", RevisionId_).Limit(limit, offset).Desc(field).Find(_BlockContent_body)
	return _BlockContent_body, err
}

// GetBlockContent_bodiesViaLangcode Get BlockContent_bodys via Langcode
func GetBlockContent_bodiesViaLangcode(offset int, limit int, Langcode_ string, field string) (*[]*BlockContent_body, error) {
	var _BlockContent_body = new([]*BlockContent_body)
	err := Engine.Table("block_content__body").Where("langcode = ?", Langcode_).Limit(limit, offset).Desc(field).Find(_BlockContent_body)
	return _BlockContent_body, err
}

// GetBlockContent_bodiesViaDelta Get BlockContent_bodys via Delta
func GetBlockContent_bodiesViaDelta(offset int, limit int, Delta_ int, field string) (*[]*BlockContent_body, error) {
	var _BlockContent_body = new([]*BlockContent_body)
	err := Engine.Table("block_content__body").Where("delta = ?", Delta_).Limit(limit, offset).Desc(field).Find(_BlockContent_body)
	return _BlockContent_body, err
}

// GetBlockContent_bodiesViaBodyValue Get BlockContent_bodys via BodyValue
func GetBlockContent_bodiesViaBodyValue(offset int, limit int, BodyValue_ string, field string) (*[]*BlockContent_body, error) {
	var _BlockContent_body = new([]*BlockContent_body)
	err := Engine.Table("block_content__body").Where("body_value = ?", BodyValue_).Limit(limit, offset).Desc(field).Find(_BlockContent_body)
	return _BlockContent_body, err
}

// GetBlockContent_bodiesViaBodySummary Get BlockContent_bodys via BodySummary
func GetBlockContent_bodiesViaBodySummary(offset int, limit int, BodySummary_ string, field string) (*[]*BlockContent_body, error) {
	var _BlockContent_body = new([]*BlockContent_body)
	err := Engine.Table("block_content__body").Where("body_summary = ?", BodySummary_).Limit(limit, offset).Desc(field).Find(_BlockContent_body)
	return _BlockContent_body, err
}

// GetBlockContent_bodiesViaBodyFormat Get BlockContent_bodys via BodyFormat
func GetBlockContent_bodiesViaBodyFormat(offset int, limit int, BodyFormat_ string, field string) (*[]*BlockContent_body, error) {
	var _BlockContent_body = new([]*BlockContent_body)
	err := Engine.Table("block_content__body").Where("body_format = ?", BodyFormat_).Limit(limit, offset).Desc(field).Find(_BlockContent_body)
	return _BlockContent_body, err
}

// HasBlockContent_bodyViaBundle Has BlockContent_body via Bundle
func HasBlockContent_bodyViaBundle(iBundle string) bool {
	if has, err := Engine.Where("bundle = ?", iBundle).Get(new(BlockContent_body)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasBlockContent_bodyViaDeleted Has BlockContent_body via Deleted
func HasBlockContent_bodyViaDeleted(iDeleted int) bool {
	if has, err := Engine.Where("deleted = ?", iDeleted).Get(new(BlockContent_body)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasBlockContent_bodyViaEntityId Has BlockContent_body via EntityId
func HasBlockContent_bodyViaEntityId(iEntityId int) bool {
	if has, err := Engine.Where("entity_id = ?", iEntityId).Get(new(BlockContent_body)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasBlockContent_bodyViaRevisionId Has BlockContent_body via RevisionId
func HasBlockContent_bodyViaRevisionId(iRevisionId int) bool {
	if has, err := Engine.Where("revision_id = ?", iRevisionId).Get(new(BlockContent_body)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasBlockContent_bodyViaLangcode Has BlockContent_body via Langcode
func HasBlockContent_bodyViaLangcode(iLangcode string) bool {
	if has, err := Engine.Where("langcode = ?", iLangcode).Get(new(BlockContent_body)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasBlockContent_bodyViaDelta Has BlockContent_body via Delta
func HasBlockContent_bodyViaDelta(iDelta int) bool {
	if has, err := Engine.Where("delta = ?", iDelta).Get(new(BlockContent_body)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasBlockContent_bodyViaBodyValue Has BlockContent_body via BodyValue
func HasBlockContent_bodyViaBodyValue(iBodyValue string) bool {
	if has, err := Engine.Where("body_value = ?", iBodyValue).Get(new(BlockContent_body)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasBlockContent_bodyViaBodySummary Has BlockContent_body via BodySummary
func HasBlockContent_bodyViaBodySummary(iBodySummary string) bool {
	if has, err := Engine.Where("body_summary = ?", iBodySummary).Get(new(BlockContent_body)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasBlockContent_bodyViaBodyFormat Has BlockContent_body via BodyFormat
func HasBlockContent_bodyViaBodyFormat(iBodyFormat string) bool {
	if has, err := Engine.Where("body_format = ?", iBodyFormat).Get(new(BlockContent_body)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// GetBlockContent_bodyViaBundle Get BlockContent_body via Bundle
func GetBlockContent_bodyViaBundle(iBundle string) (*BlockContent_body, error) {
	var _BlockContent_body = &BlockContent_body{Bundle: iBundle}
	has, err := Engine.Get(_BlockContent_body)
	if has {
		return _BlockContent_body, err
	} else {
		return nil, err
	}
}

// GetBlockContent_bodyViaDeleted Get BlockContent_body via Deleted
func GetBlockContent_bodyViaDeleted(iDeleted int) (*BlockContent_body, error) {
	var _BlockContent_body = &BlockContent_body{Deleted: iDeleted}
	has, err := Engine.Get(_BlockContent_body)
	if has {
		return _BlockContent_body, err
	} else {
		return nil, err
	}
}

// GetBlockContent_bodyViaEntityId Get BlockContent_body via EntityId
func GetBlockContent_bodyViaEntityId(iEntityId int) (*BlockContent_body, error) {
	var _BlockContent_body = &BlockContent_body{EntityId: iEntityId}
	has, err := Engine.Get(_BlockContent_body)
	if has {
		return _BlockContent_body, err
	} else {
		return nil, err
	}
}

// GetBlockContent_bodyViaRevisionId Get BlockContent_body via RevisionId
func GetBlockContent_bodyViaRevisionId(iRevisionId int) (*BlockContent_body, error) {
	var _BlockContent_body = &BlockContent_body{RevisionId: iRevisionId}
	has, err := Engine.Get(_BlockContent_body)
	if has {
		return _BlockContent_body, err
	} else {
		return nil, err
	}
}

// GetBlockContent_bodyViaLangcode Get BlockContent_body via Langcode
func GetBlockContent_bodyViaLangcode(iLangcode string) (*BlockContent_body, error) {
	var _BlockContent_body = &BlockContent_body{Langcode: iLangcode}
	has, err := Engine.Get(_BlockContent_body)
	if has {
		return _BlockContent_body, err
	} else {
		return nil, err
	}
}

// GetBlockContent_bodyViaDelta Get BlockContent_body via Delta
func GetBlockContent_bodyViaDelta(iDelta int) (*BlockContent_body, error) {
	var _BlockContent_body = &BlockContent_body{Delta: iDelta}
	has, err := Engine.Get(_BlockContent_body)
	if has {
		return _BlockContent_body, err
	} else {
		return nil, err
	}
}

// GetBlockContent_bodyViaBodyValue Get BlockContent_body via BodyValue
func GetBlockContent_bodyViaBodyValue(iBodyValue string) (*BlockContent_body, error) {
	var _BlockContent_body = &BlockContent_body{BodyValue: iBodyValue}
	has, err := Engine.Get(_BlockContent_body)
	if has {
		return _BlockContent_body, err
	} else {
		return nil, err
	}
}

// GetBlockContent_bodyViaBodySummary Get BlockContent_body via BodySummary
func GetBlockContent_bodyViaBodySummary(iBodySummary string) (*BlockContent_body, error) {
	var _BlockContent_body = &BlockContent_body{BodySummary: iBodySummary}
	has, err := Engine.Get(_BlockContent_body)
	if has {
		return _BlockContent_body, err
	} else {
		return nil, err
	}
}

// GetBlockContent_bodyViaBodyFormat Get BlockContent_body via BodyFormat
func GetBlockContent_bodyViaBodyFormat(iBodyFormat string) (*BlockContent_body, error) {
	var _BlockContent_body = &BlockContent_body{BodyFormat: iBodyFormat}
	has, err := Engine.Get(_BlockContent_body)
	if has {
		return _BlockContent_body, err
	} else {
		return nil, err
	}
}

// SetBlockContent_bodyViaBundle Set BlockContent_body via Bundle
func SetBlockContent_bodyViaBundle(iBundle string, block_content__body *BlockContent_body) (int64, error) {
	block_content__body.Bundle = iBundle
	return Engine.Insert(block_content__body)
}

// SetBlockContent_bodyViaDeleted Set BlockContent_body via Deleted
func SetBlockContent_bodyViaDeleted(iDeleted int, block_content__body *BlockContent_body) (int64, error) {
	block_content__body.Deleted = iDeleted
	return Engine.Insert(block_content__body)
}

// SetBlockContent_bodyViaEntityId Set BlockContent_body via EntityId
func SetBlockContent_bodyViaEntityId(iEntityId int, block_content__body *BlockContent_body) (int64, error) {
	block_content__body.EntityId = iEntityId
	return Engine.Insert(block_content__body)
}

// SetBlockContent_bodyViaRevisionId Set BlockContent_body via RevisionId
func SetBlockContent_bodyViaRevisionId(iRevisionId int, block_content__body *BlockContent_body) (int64, error) {
	block_content__body.RevisionId = iRevisionId
	return Engine.Insert(block_content__body)
}

// SetBlockContent_bodyViaLangcode Set BlockContent_body via Langcode
func SetBlockContent_bodyViaLangcode(iLangcode string, block_content__body *BlockContent_body) (int64, error) {
	block_content__body.Langcode = iLangcode
	return Engine.Insert(block_content__body)
}

// SetBlockContent_bodyViaDelta Set BlockContent_body via Delta
func SetBlockContent_bodyViaDelta(iDelta int, block_content__body *BlockContent_body) (int64, error) {
	block_content__body.Delta = iDelta
	return Engine.Insert(block_content__body)
}

// SetBlockContent_bodyViaBodyValue Set BlockContent_body via BodyValue
func SetBlockContent_bodyViaBodyValue(iBodyValue string, block_content__body *BlockContent_body) (int64, error) {
	block_content__body.BodyValue = iBodyValue
	return Engine.Insert(block_content__body)
}

// SetBlockContent_bodyViaBodySummary Set BlockContent_body via BodySummary
func SetBlockContent_bodyViaBodySummary(iBodySummary string, block_content__body *BlockContent_body) (int64, error) {
	block_content__body.BodySummary = iBodySummary
	return Engine.Insert(block_content__body)
}

// SetBlockContent_bodyViaBodyFormat Set BlockContent_body via BodyFormat
func SetBlockContent_bodyViaBodyFormat(iBodyFormat string, block_content__body *BlockContent_body) (int64, error) {
	block_content__body.BodyFormat = iBodyFormat
	return Engine.Insert(block_content__body)
}

// AddBlockContent_body Add BlockContent_body via all columns
func AddBlockContent_body(iBundle string, iDeleted int, iEntityId int, iRevisionId int, iLangcode string, iDelta int, iBodyValue string, iBodySummary string, iBodyFormat string) error {
	_BlockContent_body := &BlockContent_body{Bundle: iBundle, Deleted: iDeleted, EntityId: iEntityId, RevisionId: iRevisionId, Langcode: iLangcode, Delta: iDelta, BodyValue: iBodyValue, BodySummary: iBodySummary, BodyFormat: iBodyFormat}
	if _, err := Engine.Insert(_BlockContent_body); err != nil {
		return err
	} else {
		return nil
	}
}

// PostBlockContent_body Post BlockContent_body via iBlockContent_body
func PostBlockContent_body(iBlockContent_body *BlockContent_body) (string, error) {
	_, err := Engine.Insert(iBlockContent_body)
	return iBlockContent_body.Bundle, err
}

// PutBlockContent_body Put BlockContent_body
func PutBlockContent_body(iBlockContent_body *BlockContent_body) (string, error) {
	_, err := Engine.Id(iBlockContent_body.Bundle).Update(iBlockContent_body)
	return iBlockContent_body.Bundle, err
}

// PutBlockContent_bodyViaBundle Put BlockContent_body via Bundle
func PutBlockContent_bodyViaBundle(Bundle_ string, iBlockContent_body *BlockContent_body) (int64, error) {
	row, err := Engine.Update(iBlockContent_body, &BlockContent_body{Bundle: Bundle_})
	return row, err
}

// PutBlockContent_bodyViaDeleted Put BlockContent_body via Deleted
func PutBlockContent_bodyViaDeleted(Deleted_ int, iBlockContent_body *BlockContent_body) (int64, error) {
	row, err := Engine.Update(iBlockContent_body, &BlockContent_body{Deleted: Deleted_})
	return row, err
}

// PutBlockContent_bodyViaEntityId Put BlockContent_body via EntityId
func PutBlockContent_bodyViaEntityId(EntityId_ int, iBlockContent_body *BlockContent_body) (int64, error) {
	row, err := Engine.Update(iBlockContent_body, &BlockContent_body{EntityId: EntityId_})
	return row, err
}

// PutBlockContent_bodyViaRevisionId Put BlockContent_body via RevisionId
func PutBlockContent_bodyViaRevisionId(RevisionId_ int, iBlockContent_body *BlockContent_body) (int64, error) {
	row, err := Engine.Update(iBlockContent_body, &BlockContent_body{RevisionId: RevisionId_})
	return row, err
}

// PutBlockContent_bodyViaLangcode Put BlockContent_body via Langcode
func PutBlockContent_bodyViaLangcode(Langcode_ string, iBlockContent_body *BlockContent_body) (int64, error) {
	row, err := Engine.Update(iBlockContent_body, &BlockContent_body{Langcode: Langcode_})
	return row, err
}

// PutBlockContent_bodyViaDelta Put BlockContent_body via Delta
func PutBlockContent_bodyViaDelta(Delta_ int, iBlockContent_body *BlockContent_body) (int64, error) {
	row, err := Engine.Update(iBlockContent_body, &BlockContent_body{Delta: Delta_})
	return row, err
}

// PutBlockContent_bodyViaBodyValue Put BlockContent_body via BodyValue
func PutBlockContent_bodyViaBodyValue(BodyValue_ string, iBlockContent_body *BlockContent_body) (int64, error) {
	row, err := Engine.Update(iBlockContent_body, &BlockContent_body{BodyValue: BodyValue_})
	return row, err
}

// PutBlockContent_bodyViaBodySummary Put BlockContent_body via BodySummary
func PutBlockContent_bodyViaBodySummary(BodySummary_ string, iBlockContent_body *BlockContent_body) (int64, error) {
	row, err := Engine.Update(iBlockContent_body, &BlockContent_body{BodySummary: BodySummary_})
	return row, err
}

// PutBlockContent_bodyViaBodyFormat Put BlockContent_body via BodyFormat
func PutBlockContent_bodyViaBodyFormat(BodyFormat_ string, iBlockContent_body *BlockContent_body) (int64, error) {
	row, err := Engine.Update(iBlockContent_body, &BlockContent_body{BodyFormat: BodyFormat_})
	return row, err
}

// UpdateBlockContent_bodyViaBundle via map[string]interface{}{}
func UpdateBlockContent_bodyViaBundle(iBundle string, iBlockContent_bodyMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(BlockContent_body)).Where("bundle = ?", iBundle).Update(iBlockContent_bodyMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateBlockContent_bodyViaDeleted via map[string]interface{}{}
func UpdateBlockContent_bodyViaDeleted(iDeleted int, iBlockContent_bodyMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(BlockContent_body)).Where("deleted = ?", iDeleted).Update(iBlockContent_bodyMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateBlockContent_bodyViaEntityId via map[string]interface{}{}
func UpdateBlockContent_bodyViaEntityId(iEntityId int, iBlockContent_bodyMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(BlockContent_body)).Where("entity_id = ?", iEntityId).Update(iBlockContent_bodyMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateBlockContent_bodyViaRevisionId via map[string]interface{}{}
func UpdateBlockContent_bodyViaRevisionId(iRevisionId int, iBlockContent_bodyMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(BlockContent_body)).Where("revision_id = ?", iRevisionId).Update(iBlockContent_bodyMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateBlockContent_bodyViaLangcode via map[string]interface{}{}
func UpdateBlockContent_bodyViaLangcode(iLangcode string, iBlockContent_bodyMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(BlockContent_body)).Where("langcode = ?", iLangcode).Update(iBlockContent_bodyMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateBlockContent_bodyViaDelta via map[string]interface{}{}
func UpdateBlockContent_bodyViaDelta(iDelta int, iBlockContent_bodyMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(BlockContent_body)).Where("delta = ?", iDelta).Update(iBlockContent_bodyMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateBlockContent_bodyViaBodyValue via map[string]interface{}{}
func UpdateBlockContent_bodyViaBodyValue(iBodyValue string, iBlockContent_bodyMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(BlockContent_body)).Where("body_value = ?", iBodyValue).Update(iBlockContent_bodyMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateBlockContent_bodyViaBodySummary via map[string]interface{}{}
func UpdateBlockContent_bodyViaBodySummary(iBodySummary string, iBlockContent_bodyMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(BlockContent_body)).Where("body_summary = ?", iBodySummary).Update(iBlockContent_bodyMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateBlockContent_bodyViaBodyFormat via map[string]interface{}{}
func UpdateBlockContent_bodyViaBodyFormat(iBodyFormat string, iBlockContent_bodyMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(BlockContent_body)).Where("body_format = ?", iBodyFormat).Update(iBlockContent_bodyMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// DeleteBlockContent_body Delete BlockContent_body
func DeleteBlockContent_body(iBundle string) (int64, error) {
	row, err := Engine.Id(iBundle).Delete(new(BlockContent_body))
	return row, err
}

// DeleteBlockContent_bodyViaBundle Delete BlockContent_body via Bundle
func DeleteBlockContent_bodyViaBundle(iBundle string) (err error) {
	var has bool
	var _BlockContent_body = &BlockContent_body{Bundle: iBundle}
	if has, err = Engine.Get(_BlockContent_body); (has == true) && (err == nil) {
		if row, err := Engine.Where("bundle = ?", iBundle).Delete(new(BlockContent_body)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteBlockContent_bodyViaDeleted Delete BlockContent_body via Deleted
func DeleteBlockContent_bodyViaDeleted(iDeleted int) (err error) {
	var has bool
	var _BlockContent_body = &BlockContent_body{Deleted: iDeleted}
	if has, err = Engine.Get(_BlockContent_body); (has == true) && (err == nil) {
		if row, err := Engine.Where("deleted = ?", iDeleted).Delete(new(BlockContent_body)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteBlockContent_bodyViaEntityId Delete BlockContent_body via EntityId
func DeleteBlockContent_bodyViaEntityId(iEntityId int) (err error) {
	var has bool
	var _BlockContent_body = &BlockContent_body{EntityId: iEntityId}
	if has, err = Engine.Get(_BlockContent_body); (has == true) && (err == nil) {
		if row, err := Engine.Where("entity_id = ?", iEntityId).Delete(new(BlockContent_body)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteBlockContent_bodyViaRevisionId Delete BlockContent_body via RevisionId
func DeleteBlockContent_bodyViaRevisionId(iRevisionId int) (err error) {
	var has bool
	var _BlockContent_body = &BlockContent_body{RevisionId: iRevisionId}
	if has, err = Engine.Get(_BlockContent_body); (has == true) && (err == nil) {
		if row, err := Engine.Where("revision_id = ?", iRevisionId).Delete(new(BlockContent_body)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteBlockContent_bodyViaLangcode Delete BlockContent_body via Langcode
func DeleteBlockContent_bodyViaLangcode(iLangcode string) (err error) {
	var has bool
	var _BlockContent_body = &BlockContent_body{Langcode: iLangcode}
	if has, err = Engine.Get(_BlockContent_body); (has == true) && (err == nil) {
		if row, err := Engine.Where("langcode = ?", iLangcode).Delete(new(BlockContent_body)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteBlockContent_bodyViaDelta Delete BlockContent_body via Delta
func DeleteBlockContent_bodyViaDelta(iDelta int) (err error) {
	var has bool
	var _BlockContent_body = &BlockContent_body{Delta: iDelta}
	if has, err = Engine.Get(_BlockContent_body); (has == true) && (err == nil) {
		if row, err := Engine.Where("delta = ?", iDelta).Delete(new(BlockContent_body)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteBlockContent_bodyViaBodyValue Delete BlockContent_body via BodyValue
func DeleteBlockContent_bodyViaBodyValue(iBodyValue string) (err error) {
	var has bool
	var _BlockContent_body = &BlockContent_body{BodyValue: iBodyValue}
	if has, err = Engine.Get(_BlockContent_body); (has == true) && (err == nil) {
		if row, err := Engine.Where("body_value = ?", iBodyValue).Delete(new(BlockContent_body)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteBlockContent_bodyViaBodySummary Delete BlockContent_body via BodySummary
func DeleteBlockContent_bodyViaBodySummary(iBodySummary string) (err error) {
	var has bool
	var _BlockContent_body = &BlockContent_body{BodySummary: iBodySummary}
	if has, err = Engine.Get(_BlockContent_body); (has == true) && (err == nil) {
		if row, err := Engine.Where("body_summary = ?", iBodySummary).Delete(new(BlockContent_body)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteBlockContent_bodyViaBodyFormat Delete BlockContent_body via BodyFormat
func DeleteBlockContent_bodyViaBodyFormat(iBodyFormat string) (err error) {
	var has bool
	var _BlockContent_body = &BlockContent_body{BodyFormat: iBodyFormat}
	if has, err = Engine.Get(_BlockContent_body); (has == true) && (err == nil) {
		if row, err := Engine.Where("body_format = ?", iBodyFormat).Delete(new(BlockContent_body)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}
