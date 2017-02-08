package model

type Node_body struct {
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

// GetNode_bodiesCount Node_bodys' Count
func GetNode_bodiesCount(offset int, limit int) (int64, error) {
	total, err := Engine.Limit(limit, offset).Count(&Node_body{})
	return total, err
}

// GetNode_bodyCountViaBundle Get Node_body via Bundle
func GetNode_bodyCountViaBundle(iBundle string) int64 {
	n, _ := Engine.Where("bundle = ?", iBundle).Count(&Node_body{Bundle: iBundle})
	return n
}

// GetNode_bodyCountViaDeleted Get Node_body via Deleted
func GetNode_bodyCountViaDeleted(iDeleted int) int64 {
	n, _ := Engine.Where("deleted = ?", iDeleted).Count(&Node_body{Deleted: iDeleted})
	return n
}

// GetNode_bodyCountViaEntityId Get Node_body via EntityId
func GetNode_bodyCountViaEntityId(iEntityId int) int64 {
	n, _ := Engine.Where("entity_id = ?", iEntityId).Count(&Node_body{EntityId: iEntityId})
	return n
}

// GetNode_bodyCountViaRevisionId Get Node_body via RevisionId
func GetNode_bodyCountViaRevisionId(iRevisionId int) int64 {
	n, _ := Engine.Where("revision_id = ?", iRevisionId).Count(&Node_body{RevisionId: iRevisionId})
	return n
}

// GetNode_bodyCountViaLangcode Get Node_body via Langcode
func GetNode_bodyCountViaLangcode(iLangcode string) int64 {
	n, _ := Engine.Where("langcode = ?", iLangcode).Count(&Node_body{Langcode: iLangcode})
	return n
}

// GetNode_bodyCountViaDelta Get Node_body via Delta
func GetNode_bodyCountViaDelta(iDelta int) int64 {
	n, _ := Engine.Where("delta = ?", iDelta).Count(&Node_body{Delta: iDelta})
	return n
}

// GetNode_bodyCountViaBodyValue Get Node_body via BodyValue
func GetNode_bodyCountViaBodyValue(iBodyValue string) int64 {
	n, _ := Engine.Where("body_value = ?", iBodyValue).Count(&Node_body{BodyValue: iBodyValue})
	return n
}

// GetNode_bodyCountViaBodySummary Get Node_body via BodySummary
func GetNode_bodyCountViaBodySummary(iBodySummary string) int64 {
	n, _ := Engine.Where("body_summary = ?", iBodySummary).Count(&Node_body{BodySummary: iBodySummary})
	return n
}

// GetNode_bodyCountViaBodyFormat Get Node_body via BodyFormat
func GetNode_bodyCountViaBodyFormat(iBodyFormat string) int64 {
	n, _ := Engine.Where("body_format = ?", iBodyFormat).Count(&Node_body{BodyFormat: iBodyFormat})
	return n
}

// GetNode_bodiesByBundleAndDeletedAndEntityId Get Node_bodies via BundleAndDeletedAndEntityId
func GetNode_bodiesByBundleAndDeletedAndEntityId(offset int, limit int, Bundle_ string, Deleted_ int, EntityId_ int) (*[]*Node_body, error) {
	var _Node_body = new([]*Node_body)
	err := Engine.Table("node__body").Where("bundle = ? and deleted = ? and entity_id = ?", Bundle_, Deleted_, EntityId_).Limit(limit, offset).Find(_Node_body)
	return _Node_body, err
}

// GetNode_bodiesByBundleAndDeletedAndRevisionId Get Node_bodies via BundleAndDeletedAndRevisionId
func GetNode_bodiesByBundleAndDeletedAndRevisionId(offset int, limit int, Bundle_ string, Deleted_ int, RevisionId_ int) (*[]*Node_body, error) {
	var _Node_body = new([]*Node_body)
	err := Engine.Table("node__body").Where("bundle = ? and deleted = ? and revision_id = ?", Bundle_, Deleted_, RevisionId_).Limit(limit, offset).Find(_Node_body)
	return _Node_body, err
}

// GetNode_bodiesByBundleAndDeletedAndLangcode Get Node_bodies via BundleAndDeletedAndLangcode
func GetNode_bodiesByBundleAndDeletedAndLangcode(offset int, limit int, Bundle_ string, Deleted_ int, Langcode_ string) (*[]*Node_body, error) {
	var _Node_body = new([]*Node_body)
	err := Engine.Table("node__body").Where("bundle = ? and deleted = ? and langcode = ?", Bundle_, Deleted_, Langcode_).Limit(limit, offset).Find(_Node_body)
	return _Node_body, err
}

// GetNode_bodiesByBundleAndDeletedAndDelta Get Node_bodies via BundleAndDeletedAndDelta
func GetNode_bodiesByBundleAndDeletedAndDelta(offset int, limit int, Bundle_ string, Deleted_ int, Delta_ int) (*[]*Node_body, error) {
	var _Node_body = new([]*Node_body)
	err := Engine.Table("node__body").Where("bundle = ? and deleted = ? and delta = ?", Bundle_, Deleted_, Delta_).Limit(limit, offset).Find(_Node_body)
	return _Node_body, err
}

// GetNode_bodiesByBundleAndDeletedAndBodyValue Get Node_bodies via BundleAndDeletedAndBodyValue
func GetNode_bodiesByBundleAndDeletedAndBodyValue(offset int, limit int, Bundle_ string, Deleted_ int, BodyValue_ string) (*[]*Node_body, error) {
	var _Node_body = new([]*Node_body)
	err := Engine.Table("node__body").Where("bundle = ? and deleted = ? and body_value = ?", Bundle_, Deleted_, BodyValue_).Limit(limit, offset).Find(_Node_body)
	return _Node_body, err
}

// GetNode_bodiesByBundleAndDeletedAndBodySummary Get Node_bodies via BundleAndDeletedAndBodySummary
func GetNode_bodiesByBundleAndDeletedAndBodySummary(offset int, limit int, Bundle_ string, Deleted_ int, BodySummary_ string) (*[]*Node_body, error) {
	var _Node_body = new([]*Node_body)
	err := Engine.Table("node__body").Where("bundle = ? and deleted = ? and body_summary = ?", Bundle_, Deleted_, BodySummary_).Limit(limit, offset).Find(_Node_body)
	return _Node_body, err
}

// GetNode_bodiesByBundleAndDeletedAndBodyFormat Get Node_bodies via BundleAndDeletedAndBodyFormat
func GetNode_bodiesByBundleAndDeletedAndBodyFormat(offset int, limit int, Bundle_ string, Deleted_ int, BodyFormat_ string) (*[]*Node_body, error) {
	var _Node_body = new([]*Node_body)
	err := Engine.Table("node__body").Where("bundle = ? and deleted = ? and body_format = ?", Bundle_, Deleted_, BodyFormat_).Limit(limit, offset).Find(_Node_body)
	return _Node_body, err
}

// GetNode_bodiesByBundleAndEntityIdAndRevisionId Get Node_bodies via BundleAndEntityIdAndRevisionId
func GetNode_bodiesByBundleAndEntityIdAndRevisionId(offset int, limit int, Bundle_ string, EntityId_ int, RevisionId_ int) (*[]*Node_body, error) {
	var _Node_body = new([]*Node_body)
	err := Engine.Table("node__body").Where("bundle = ? and entity_id = ? and revision_id = ?", Bundle_, EntityId_, RevisionId_).Limit(limit, offset).Find(_Node_body)
	return _Node_body, err
}

// GetNode_bodiesByBundleAndEntityIdAndLangcode Get Node_bodies via BundleAndEntityIdAndLangcode
func GetNode_bodiesByBundleAndEntityIdAndLangcode(offset int, limit int, Bundle_ string, EntityId_ int, Langcode_ string) (*[]*Node_body, error) {
	var _Node_body = new([]*Node_body)
	err := Engine.Table("node__body").Where("bundle = ? and entity_id = ? and langcode = ?", Bundle_, EntityId_, Langcode_).Limit(limit, offset).Find(_Node_body)
	return _Node_body, err
}

// GetNode_bodiesByBundleAndEntityIdAndDelta Get Node_bodies via BundleAndEntityIdAndDelta
func GetNode_bodiesByBundleAndEntityIdAndDelta(offset int, limit int, Bundle_ string, EntityId_ int, Delta_ int) (*[]*Node_body, error) {
	var _Node_body = new([]*Node_body)
	err := Engine.Table("node__body").Where("bundle = ? and entity_id = ? and delta = ?", Bundle_, EntityId_, Delta_).Limit(limit, offset).Find(_Node_body)
	return _Node_body, err
}

// GetNode_bodiesByBundleAndEntityIdAndBodyValue Get Node_bodies via BundleAndEntityIdAndBodyValue
func GetNode_bodiesByBundleAndEntityIdAndBodyValue(offset int, limit int, Bundle_ string, EntityId_ int, BodyValue_ string) (*[]*Node_body, error) {
	var _Node_body = new([]*Node_body)
	err := Engine.Table("node__body").Where("bundle = ? and entity_id = ? and body_value = ?", Bundle_, EntityId_, BodyValue_).Limit(limit, offset).Find(_Node_body)
	return _Node_body, err
}

// GetNode_bodiesByBundleAndEntityIdAndBodySummary Get Node_bodies via BundleAndEntityIdAndBodySummary
func GetNode_bodiesByBundleAndEntityIdAndBodySummary(offset int, limit int, Bundle_ string, EntityId_ int, BodySummary_ string) (*[]*Node_body, error) {
	var _Node_body = new([]*Node_body)
	err := Engine.Table("node__body").Where("bundle = ? and entity_id = ? and body_summary = ?", Bundle_, EntityId_, BodySummary_).Limit(limit, offset).Find(_Node_body)
	return _Node_body, err
}

// GetNode_bodiesByBundleAndEntityIdAndBodyFormat Get Node_bodies via BundleAndEntityIdAndBodyFormat
func GetNode_bodiesByBundleAndEntityIdAndBodyFormat(offset int, limit int, Bundle_ string, EntityId_ int, BodyFormat_ string) (*[]*Node_body, error) {
	var _Node_body = new([]*Node_body)
	err := Engine.Table("node__body").Where("bundle = ? and entity_id = ? and body_format = ?", Bundle_, EntityId_, BodyFormat_).Limit(limit, offset).Find(_Node_body)
	return _Node_body, err
}

// GetNode_bodiesByBundleAndRevisionIdAndLangcode Get Node_bodies via BundleAndRevisionIdAndLangcode
func GetNode_bodiesByBundleAndRevisionIdAndLangcode(offset int, limit int, Bundle_ string, RevisionId_ int, Langcode_ string) (*[]*Node_body, error) {
	var _Node_body = new([]*Node_body)
	err := Engine.Table("node__body").Where("bundle = ? and revision_id = ? and langcode = ?", Bundle_, RevisionId_, Langcode_).Limit(limit, offset).Find(_Node_body)
	return _Node_body, err
}

// GetNode_bodiesByBundleAndRevisionIdAndDelta Get Node_bodies via BundleAndRevisionIdAndDelta
func GetNode_bodiesByBundleAndRevisionIdAndDelta(offset int, limit int, Bundle_ string, RevisionId_ int, Delta_ int) (*[]*Node_body, error) {
	var _Node_body = new([]*Node_body)
	err := Engine.Table("node__body").Where("bundle = ? and revision_id = ? and delta = ?", Bundle_, RevisionId_, Delta_).Limit(limit, offset).Find(_Node_body)
	return _Node_body, err
}

// GetNode_bodiesByBundleAndRevisionIdAndBodyValue Get Node_bodies via BundleAndRevisionIdAndBodyValue
func GetNode_bodiesByBundleAndRevisionIdAndBodyValue(offset int, limit int, Bundle_ string, RevisionId_ int, BodyValue_ string) (*[]*Node_body, error) {
	var _Node_body = new([]*Node_body)
	err := Engine.Table("node__body").Where("bundle = ? and revision_id = ? and body_value = ?", Bundle_, RevisionId_, BodyValue_).Limit(limit, offset).Find(_Node_body)
	return _Node_body, err
}

// GetNode_bodiesByBundleAndRevisionIdAndBodySummary Get Node_bodies via BundleAndRevisionIdAndBodySummary
func GetNode_bodiesByBundleAndRevisionIdAndBodySummary(offset int, limit int, Bundle_ string, RevisionId_ int, BodySummary_ string) (*[]*Node_body, error) {
	var _Node_body = new([]*Node_body)
	err := Engine.Table("node__body").Where("bundle = ? and revision_id = ? and body_summary = ?", Bundle_, RevisionId_, BodySummary_).Limit(limit, offset).Find(_Node_body)
	return _Node_body, err
}

// GetNode_bodiesByBundleAndRevisionIdAndBodyFormat Get Node_bodies via BundleAndRevisionIdAndBodyFormat
func GetNode_bodiesByBundleAndRevisionIdAndBodyFormat(offset int, limit int, Bundle_ string, RevisionId_ int, BodyFormat_ string) (*[]*Node_body, error) {
	var _Node_body = new([]*Node_body)
	err := Engine.Table("node__body").Where("bundle = ? and revision_id = ? and body_format = ?", Bundle_, RevisionId_, BodyFormat_).Limit(limit, offset).Find(_Node_body)
	return _Node_body, err
}

// GetNode_bodiesByBundleAndLangcodeAndDelta Get Node_bodies via BundleAndLangcodeAndDelta
func GetNode_bodiesByBundleAndLangcodeAndDelta(offset int, limit int, Bundle_ string, Langcode_ string, Delta_ int) (*[]*Node_body, error) {
	var _Node_body = new([]*Node_body)
	err := Engine.Table("node__body").Where("bundle = ? and langcode = ? and delta = ?", Bundle_, Langcode_, Delta_).Limit(limit, offset).Find(_Node_body)
	return _Node_body, err
}

// GetNode_bodiesByBundleAndLangcodeAndBodyValue Get Node_bodies via BundleAndLangcodeAndBodyValue
func GetNode_bodiesByBundleAndLangcodeAndBodyValue(offset int, limit int, Bundle_ string, Langcode_ string, BodyValue_ string) (*[]*Node_body, error) {
	var _Node_body = new([]*Node_body)
	err := Engine.Table("node__body").Where("bundle = ? and langcode = ? and body_value = ?", Bundle_, Langcode_, BodyValue_).Limit(limit, offset).Find(_Node_body)
	return _Node_body, err
}

// GetNode_bodiesByBundleAndLangcodeAndBodySummary Get Node_bodies via BundleAndLangcodeAndBodySummary
func GetNode_bodiesByBundleAndLangcodeAndBodySummary(offset int, limit int, Bundle_ string, Langcode_ string, BodySummary_ string) (*[]*Node_body, error) {
	var _Node_body = new([]*Node_body)
	err := Engine.Table("node__body").Where("bundle = ? and langcode = ? and body_summary = ?", Bundle_, Langcode_, BodySummary_).Limit(limit, offset).Find(_Node_body)
	return _Node_body, err
}

// GetNode_bodiesByBundleAndLangcodeAndBodyFormat Get Node_bodies via BundleAndLangcodeAndBodyFormat
func GetNode_bodiesByBundleAndLangcodeAndBodyFormat(offset int, limit int, Bundle_ string, Langcode_ string, BodyFormat_ string) (*[]*Node_body, error) {
	var _Node_body = new([]*Node_body)
	err := Engine.Table("node__body").Where("bundle = ? and langcode = ? and body_format = ?", Bundle_, Langcode_, BodyFormat_).Limit(limit, offset).Find(_Node_body)
	return _Node_body, err
}

// GetNode_bodiesByBundleAndDeltaAndBodyValue Get Node_bodies via BundleAndDeltaAndBodyValue
func GetNode_bodiesByBundleAndDeltaAndBodyValue(offset int, limit int, Bundle_ string, Delta_ int, BodyValue_ string) (*[]*Node_body, error) {
	var _Node_body = new([]*Node_body)
	err := Engine.Table("node__body").Where("bundle = ? and delta = ? and body_value = ?", Bundle_, Delta_, BodyValue_).Limit(limit, offset).Find(_Node_body)
	return _Node_body, err
}

// GetNode_bodiesByBundleAndDeltaAndBodySummary Get Node_bodies via BundleAndDeltaAndBodySummary
func GetNode_bodiesByBundleAndDeltaAndBodySummary(offset int, limit int, Bundle_ string, Delta_ int, BodySummary_ string) (*[]*Node_body, error) {
	var _Node_body = new([]*Node_body)
	err := Engine.Table("node__body").Where("bundle = ? and delta = ? and body_summary = ?", Bundle_, Delta_, BodySummary_).Limit(limit, offset).Find(_Node_body)
	return _Node_body, err
}

// GetNode_bodiesByBundleAndDeltaAndBodyFormat Get Node_bodies via BundleAndDeltaAndBodyFormat
func GetNode_bodiesByBundleAndDeltaAndBodyFormat(offset int, limit int, Bundle_ string, Delta_ int, BodyFormat_ string) (*[]*Node_body, error) {
	var _Node_body = new([]*Node_body)
	err := Engine.Table("node__body").Where("bundle = ? and delta = ? and body_format = ?", Bundle_, Delta_, BodyFormat_).Limit(limit, offset).Find(_Node_body)
	return _Node_body, err
}

// GetNode_bodiesByBundleAndBodyValueAndBodySummary Get Node_bodies via BundleAndBodyValueAndBodySummary
func GetNode_bodiesByBundleAndBodyValueAndBodySummary(offset int, limit int, Bundle_ string, BodyValue_ string, BodySummary_ string) (*[]*Node_body, error) {
	var _Node_body = new([]*Node_body)
	err := Engine.Table("node__body").Where("bundle = ? and body_value = ? and body_summary = ?", Bundle_, BodyValue_, BodySummary_).Limit(limit, offset).Find(_Node_body)
	return _Node_body, err
}

// GetNode_bodiesByBundleAndBodyValueAndBodyFormat Get Node_bodies via BundleAndBodyValueAndBodyFormat
func GetNode_bodiesByBundleAndBodyValueAndBodyFormat(offset int, limit int, Bundle_ string, BodyValue_ string, BodyFormat_ string) (*[]*Node_body, error) {
	var _Node_body = new([]*Node_body)
	err := Engine.Table("node__body").Where("bundle = ? and body_value = ? and body_format = ?", Bundle_, BodyValue_, BodyFormat_).Limit(limit, offset).Find(_Node_body)
	return _Node_body, err
}

// GetNode_bodiesByBundleAndBodySummaryAndBodyFormat Get Node_bodies via BundleAndBodySummaryAndBodyFormat
func GetNode_bodiesByBundleAndBodySummaryAndBodyFormat(offset int, limit int, Bundle_ string, BodySummary_ string, BodyFormat_ string) (*[]*Node_body, error) {
	var _Node_body = new([]*Node_body)
	err := Engine.Table("node__body").Where("bundle = ? and body_summary = ? and body_format = ?", Bundle_, BodySummary_, BodyFormat_).Limit(limit, offset).Find(_Node_body)
	return _Node_body, err
}

// GetNode_bodiesByDeletedAndEntityIdAndRevisionId Get Node_bodies via DeletedAndEntityIdAndRevisionId
func GetNode_bodiesByDeletedAndEntityIdAndRevisionId(offset int, limit int, Deleted_ int, EntityId_ int, RevisionId_ int) (*[]*Node_body, error) {
	var _Node_body = new([]*Node_body)
	err := Engine.Table("node__body").Where("deleted = ? and entity_id = ? and revision_id = ?", Deleted_, EntityId_, RevisionId_).Limit(limit, offset).Find(_Node_body)
	return _Node_body, err
}

// GetNode_bodiesByDeletedAndEntityIdAndLangcode Get Node_bodies via DeletedAndEntityIdAndLangcode
func GetNode_bodiesByDeletedAndEntityIdAndLangcode(offset int, limit int, Deleted_ int, EntityId_ int, Langcode_ string) (*[]*Node_body, error) {
	var _Node_body = new([]*Node_body)
	err := Engine.Table("node__body").Where("deleted = ? and entity_id = ? and langcode = ?", Deleted_, EntityId_, Langcode_).Limit(limit, offset).Find(_Node_body)
	return _Node_body, err
}

// GetNode_bodiesByDeletedAndEntityIdAndDelta Get Node_bodies via DeletedAndEntityIdAndDelta
func GetNode_bodiesByDeletedAndEntityIdAndDelta(offset int, limit int, Deleted_ int, EntityId_ int, Delta_ int) (*[]*Node_body, error) {
	var _Node_body = new([]*Node_body)
	err := Engine.Table("node__body").Where("deleted = ? and entity_id = ? and delta = ?", Deleted_, EntityId_, Delta_).Limit(limit, offset).Find(_Node_body)
	return _Node_body, err
}

// GetNode_bodiesByDeletedAndEntityIdAndBodyValue Get Node_bodies via DeletedAndEntityIdAndBodyValue
func GetNode_bodiesByDeletedAndEntityIdAndBodyValue(offset int, limit int, Deleted_ int, EntityId_ int, BodyValue_ string) (*[]*Node_body, error) {
	var _Node_body = new([]*Node_body)
	err := Engine.Table("node__body").Where("deleted = ? and entity_id = ? and body_value = ?", Deleted_, EntityId_, BodyValue_).Limit(limit, offset).Find(_Node_body)
	return _Node_body, err
}

// GetNode_bodiesByDeletedAndEntityIdAndBodySummary Get Node_bodies via DeletedAndEntityIdAndBodySummary
func GetNode_bodiesByDeletedAndEntityIdAndBodySummary(offset int, limit int, Deleted_ int, EntityId_ int, BodySummary_ string) (*[]*Node_body, error) {
	var _Node_body = new([]*Node_body)
	err := Engine.Table("node__body").Where("deleted = ? and entity_id = ? and body_summary = ?", Deleted_, EntityId_, BodySummary_).Limit(limit, offset).Find(_Node_body)
	return _Node_body, err
}

// GetNode_bodiesByDeletedAndEntityIdAndBodyFormat Get Node_bodies via DeletedAndEntityIdAndBodyFormat
func GetNode_bodiesByDeletedAndEntityIdAndBodyFormat(offset int, limit int, Deleted_ int, EntityId_ int, BodyFormat_ string) (*[]*Node_body, error) {
	var _Node_body = new([]*Node_body)
	err := Engine.Table("node__body").Where("deleted = ? and entity_id = ? and body_format = ?", Deleted_, EntityId_, BodyFormat_).Limit(limit, offset).Find(_Node_body)
	return _Node_body, err
}

// GetNode_bodiesByDeletedAndRevisionIdAndLangcode Get Node_bodies via DeletedAndRevisionIdAndLangcode
func GetNode_bodiesByDeletedAndRevisionIdAndLangcode(offset int, limit int, Deleted_ int, RevisionId_ int, Langcode_ string) (*[]*Node_body, error) {
	var _Node_body = new([]*Node_body)
	err := Engine.Table("node__body").Where("deleted = ? and revision_id = ? and langcode = ?", Deleted_, RevisionId_, Langcode_).Limit(limit, offset).Find(_Node_body)
	return _Node_body, err
}

// GetNode_bodiesByDeletedAndRevisionIdAndDelta Get Node_bodies via DeletedAndRevisionIdAndDelta
func GetNode_bodiesByDeletedAndRevisionIdAndDelta(offset int, limit int, Deleted_ int, RevisionId_ int, Delta_ int) (*[]*Node_body, error) {
	var _Node_body = new([]*Node_body)
	err := Engine.Table("node__body").Where("deleted = ? and revision_id = ? and delta = ?", Deleted_, RevisionId_, Delta_).Limit(limit, offset).Find(_Node_body)
	return _Node_body, err
}

// GetNode_bodiesByDeletedAndRevisionIdAndBodyValue Get Node_bodies via DeletedAndRevisionIdAndBodyValue
func GetNode_bodiesByDeletedAndRevisionIdAndBodyValue(offset int, limit int, Deleted_ int, RevisionId_ int, BodyValue_ string) (*[]*Node_body, error) {
	var _Node_body = new([]*Node_body)
	err := Engine.Table("node__body").Where("deleted = ? and revision_id = ? and body_value = ?", Deleted_, RevisionId_, BodyValue_).Limit(limit, offset).Find(_Node_body)
	return _Node_body, err
}

// GetNode_bodiesByDeletedAndRevisionIdAndBodySummary Get Node_bodies via DeletedAndRevisionIdAndBodySummary
func GetNode_bodiesByDeletedAndRevisionIdAndBodySummary(offset int, limit int, Deleted_ int, RevisionId_ int, BodySummary_ string) (*[]*Node_body, error) {
	var _Node_body = new([]*Node_body)
	err := Engine.Table("node__body").Where("deleted = ? and revision_id = ? and body_summary = ?", Deleted_, RevisionId_, BodySummary_).Limit(limit, offset).Find(_Node_body)
	return _Node_body, err
}

// GetNode_bodiesByDeletedAndRevisionIdAndBodyFormat Get Node_bodies via DeletedAndRevisionIdAndBodyFormat
func GetNode_bodiesByDeletedAndRevisionIdAndBodyFormat(offset int, limit int, Deleted_ int, RevisionId_ int, BodyFormat_ string) (*[]*Node_body, error) {
	var _Node_body = new([]*Node_body)
	err := Engine.Table("node__body").Where("deleted = ? and revision_id = ? and body_format = ?", Deleted_, RevisionId_, BodyFormat_).Limit(limit, offset).Find(_Node_body)
	return _Node_body, err
}

// GetNode_bodiesByDeletedAndLangcodeAndDelta Get Node_bodies via DeletedAndLangcodeAndDelta
func GetNode_bodiesByDeletedAndLangcodeAndDelta(offset int, limit int, Deleted_ int, Langcode_ string, Delta_ int) (*[]*Node_body, error) {
	var _Node_body = new([]*Node_body)
	err := Engine.Table("node__body").Where("deleted = ? and langcode = ? and delta = ?", Deleted_, Langcode_, Delta_).Limit(limit, offset).Find(_Node_body)
	return _Node_body, err
}

// GetNode_bodiesByDeletedAndLangcodeAndBodyValue Get Node_bodies via DeletedAndLangcodeAndBodyValue
func GetNode_bodiesByDeletedAndLangcodeAndBodyValue(offset int, limit int, Deleted_ int, Langcode_ string, BodyValue_ string) (*[]*Node_body, error) {
	var _Node_body = new([]*Node_body)
	err := Engine.Table("node__body").Where("deleted = ? and langcode = ? and body_value = ?", Deleted_, Langcode_, BodyValue_).Limit(limit, offset).Find(_Node_body)
	return _Node_body, err
}

// GetNode_bodiesByDeletedAndLangcodeAndBodySummary Get Node_bodies via DeletedAndLangcodeAndBodySummary
func GetNode_bodiesByDeletedAndLangcodeAndBodySummary(offset int, limit int, Deleted_ int, Langcode_ string, BodySummary_ string) (*[]*Node_body, error) {
	var _Node_body = new([]*Node_body)
	err := Engine.Table("node__body").Where("deleted = ? and langcode = ? and body_summary = ?", Deleted_, Langcode_, BodySummary_).Limit(limit, offset).Find(_Node_body)
	return _Node_body, err
}

// GetNode_bodiesByDeletedAndLangcodeAndBodyFormat Get Node_bodies via DeletedAndLangcodeAndBodyFormat
func GetNode_bodiesByDeletedAndLangcodeAndBodyFormat(offset int, limit int, Deleted_ int, Langcode_ string, BodyFormat_ string) (*[]*Node_body, error) {
	var _Node_body = new([]*Node_body)
	err := Engine.Table("node__body").Where("deleted = ? and langcode = ? and body_format = ?", Deleted_, Langcode_, BodyFormat_).Limit(limit, offset).Find(_Node_body)
	return _Node_body, err
}

// GetNode_bodiesByDeletedAndDeltaAndBodyValue Get Node_bodies via DeletedAndDeltaAndBodyValue
func GetNode_bodiesByDeletedAndDeltaAndBodyValue(offset int, limit int, Deleted_ int, Delta_ int, BodyValue_ string) (*[]*Node_body, error) {
	var _Node_body = new([]*Node_body)
	err := Engine.Table("node__body").Where("deleted = ? and delta = ? and body_value = ?", Deleted_, Delta_, BodyValue_).Limit(limit, offset).Find(_Node_body)
	return _Node_body, err
}

// GetNode_bodiesByDeletedAndDeltaAndBodySummary Get Node_bodies via DeletedAndDeltaAndBodySummary
func GetNode_bodiesByDeletedAndDeltaAndBodySummary(offset int, limit int, Deleted_ int, Delta_ int, BodySummary_ string) (*[]*Node_body, error) {
	var _Node_body = new([]*Node_body)
	err := Engine.Table("node__body").Where("deleted = ? and delta = ? and body_summary = ?", Deleted_, Delta_, BodySummary_).Limit(limit, offset).Find(_Node_body)
	return _Node_body, err
}

// GetNode_bodiesByDeletedAndDeltaAndBodyFormat Get Node_bodies via DeletedAndDeltaAndBodyFormat
func GetNode_bodiesByDeletedAndDeltaAndBodyFormat(offset int, limit int, Deleted_ int, Delta_ int, BodyFormat_ string) (*[]*Node_body, error) {
	var _Node_body = new([]*Node_body)
	err := Engine.Table("node__body").Where("deleted = ? and delta = ? and body_format = ?", Deleted_, Delta_, BodyFormat_).Limit(limit, offset).Find(_Node_body)
	return _Node_body, err
}

// GetNode_bodiesByDeletedAndBodyValueAndBodySummary Get Node_bodies via DeletedAndBodyValueAndBodySummary
func GetNode_bodiesByDeletedAndBodyValueAndBodySummary(offset int, limit int, Deleted_ int, BodyValue_ string, BodySummary_ string) (*[]*Node_body, error) {
	var _Node_body = new([]*Node_body)
	err := Engine.Table("node__body").Where("deleted = ? and body_value = ? and body_summary = ?", Deleted_, BodyValue_, BodySummary_).Limit(limit, offset).Find(_Node_body)
	return _Node_body, err
}

// GetNode_bodiesByDeletedAndBodyValueAndBodyFormat Get Node_bodies via DeletedAndBodyValueAndBodyFormat
func GetNode_bodiesByDeletedAndBodyValueAndBodyFormat(offset int, limit int, Deleted_ int, BodyValue_ string, BodyFormat_ string) (*[]*Node_body, error) {
	var _Node_body = new([]*Node_body)
	err := Engine.Table("node__body").Where("deleted = ? and body_value = ? and body_format = ?", Deleted_, BodyValue_, BodyFormat_).Limit(limit, offset).Find(_Node_body)
	return _Node_body, err
}

// GetNode_bodiesByDeletedAndBodySummaryAndBodyFormat Get Node_bodies via DeletedAndBodySummaryAndBodyFormat
func GetNode_bodiesByDeletedAndBodySummaryAndBodyFormat(offset int, limit int, Deleted_ int, BodySummary_ string, BodyFormat_ string) (*[]*Node_body, error) {
	var _Node_body = new([]*Node_body)
	err := Engine.Table("node__body").Where("deleted = ? and body_summary = ? and body_format = ?", Deleted_, BodySummary_, BodyFormat_).Limit(limit, offset).Find(_Node_body)
	return _Node_body, err
}

// GetNode_bodiesByEntityIdAndRevisionIdAndLangcode Get Node_bodies via EntityIdAndRevisionIdAndLangcode
func GetNode_bodiesByEntityIdAndRevisionIdAndLangcode(offset int, limit int, EntityId_ int, RevisionId_ int, Langcode_ string) (*[]*Node_body, error) {
	var _Node_body = new([]*Node_body)
	err := Engine.Table("node__body").Where("entity_id = ? and revision_id = ? and langcode = ?", EntityId_, RevisionId_, Langcode_).Limit(limit, offset).Find(_Node_body)
	return _Node_body, err
}

// GetNode_bodiesByEntityIdAndRevisionIdAndDelta Get Node_bodies via EntityIdAndRevisionIdAndDelta
func GetNode_bodiesByEntityIdAndRevisionIdAndDelta(offset int, limit int, EntityId_ int, RevisionId_ int, Delta_ int) (*[]*Node_body, error) {
	var _Node_body = new([]*Node_body)
	err := Engine.Table("node__body").Where("entity_id = ? and revision_id = ? and delta = ?", EntityId_, RevisionId_, Delta_).Limit(limit, offset).Find(_Node_body)
	return _Node_body, err
}

// GetNode_bodiesByEntityIdAndRevisionIdAndBodyValue Get Node_bodies via EntityIdAndRevisionIdAndBodyValue
func GetNode_bodiesByEntityIdAndRevisionIdAndBodyValue(offset int, limit int, EntityId_ int, RevisionId_ int, BodyValue_ string) (*[]*Node_body, error) {
	var _Node_body = new([]*Node_body)
	err := Engine.Table("node__body").Where("entity_id = ? and revision_id = ? and body_value = ?", EntityId_, RevisionId_, BodyValue_).Limit(limit, offset).Find(_Node_body)
	return _Node_body, err
}

// GetNode_bodiesByEntityIdAndRevisionIdAndBodySummary Get Node_bodies via EntityIdAndRevisionIdAndBodySummary
func GetNode_bodiesByEntityIdAndRevisionIdAndBodySummary(offset int, limit int, EntityId_ int, RevisionId_ int, BodySummary_ string) (*[]*Node_body, error) {
	var _Node_body = new([]*Node_body)
	err := Engine.Table("node__body").Where("entity_id = ? and revision_id = ? and body_summary = ?", EntityId_, RevisionId_, BodySummary_).Limit(limit, offset).Find(_Node_body)
	return _Node_body, err
}

// GetNode_bodiesByEntityIdAndRevisionIdAndBodyFormat Get Node_bodies via EntityIdAndRevisionIdAndBodyFormat
func GetNode_bodiesByEntityIdAndRevisionIdAndBodyFormat(offset int, limit int, EntityId_ int, RevisionId_ int, BodyFormat_ string) (*[]*Node_body, error) {
	var _Node_body = new([]*Node_body)
	err := Engine.Table("node__body").Where("entity_id = ? and revision_id = ? and body_format = ?", EntityId_, RevisionId_, BodyFormat_).Limit(limit, offset).Find(_Node_body)
	return _Node_body, err
}

// GetNode_bodiesByEntityIdAndLangcodeAndDelta Get Node_bodies via EntityIdAndLangcodeAndDelta
func GetNode_bodiesByEntityIdAndLangcodeAndDelta(offset int, limit int, EntityId_ int, Langcode_ string, Delta_ int) (*[]*Node_body, error) {
	var _Node_body = new([]*Node_body)
	err := Engine.Table("node__body").Where("entity_id = ? and langcode = ? and delta = ?", EntityId_, Langcode_, Delta_).Limit(limit, offset).Find(_Node_body)
	return _Node_body, err
}

// GetNode_bodiesByEntityIdAndLangcodeAndBodyValue Get Node_bodies via EntityIdAndLangcodeAndBodyValue
func GetNode_bodiesByEntityIdAndLangcodeAndBodyValue(offset int, limit int, EntityId_ int, Langcode_ string, BodyValue_ string) (*[]*Node_body, error) {
	var _Node_body = new([]*Node_body)
	err := Engine.Table("node__body").Where("entity_id = ? and langcode = ? and body_value = ?", EntityId_, Langcode_, BodyValue_).Limit(limit, offset).Find(_Node_body)
	return _Node_body, err
}

// GetNode_bodiesByEntityIdAndLangcodeAndBodySummary Get Node_bodies via EntityIdAndLangcodeAndBodySummary
func GetNode_bodiesByEntityIdAndLangcodeAndBodySummary(offset int, limit int, EntityId_ int, Langcode_ string, BodySummary_ string) (*[]*Node_body, error) {
	var _Node_body = new([]*Node_body)
	err := Engine.Table("node__body").Where("entity_id = ? and langcode = ? and body_summary = ?", EntityId_, Langcode_, BodySummary_).Limit(limit, offset).Find(_Node_body)
	return _Node_body, err
}

// GetNode_bodiesByEntityIdAndLangcodeAndBodyFormat Get Node_bodies via EntityIdAndLangcodeAndBodyFormat
func GetNode_bodiesByEntityIdAndLangcodeAndBodyFormat(offset int, limit int, EntityId_ int, Langcode_ string, BodyFormat_ string) (*[]*Node_body, error) {
	var _Node_body = new([]*Node_body)
	err := Engine.Table("node__body").Where("entity_id = ? and langcode = ? and body_format = ?", EntityId_, Langcode_, BodyFormat_).Limit(limit, offset).Find(_Node_body)
	return _Node_body, err
}

// GetNode_bodiesByEntityIdAndDeltaAndBodyValue Get Node_bodies via EntityIdAndDeltaAndBodyValue
func GetNode_bodiesByEntityIdAndDeltaAndBodyValue(offset int, limit int, EntityId_ int, Delta_ int, BodyValue_ string) (*[]*Node_body, error) {
	var _Node_body = new([]*Node_body)
	err := Engine.Table("node__body").Where("entity_id = ? and delta = ? and body_value = ?", EntityId_, Delta_, BodyValue_).Limit(limit, offset).Find(_Node_body)
	return _Node_body, err
}

// GetNode_bodiesByEntityIdAndDeltaAndBodySummary Get Node_bodies via EntityIdAndDeltaAndBodySummary
func GetNode_bodiesByEntityIdAndDeltaAndBodySummary(offset int, limit int, EntityId_ int, Delta_ int, BodySummary_ string) (*[]*Node_body, error) {
	var _Node_body = new([]*Node_body)
	err := Engine.Table("node__body").Where("entity_id = ? and delta = ? and body_summary = ?", EntityId_, Delta_, BodySummary_).Limit(limit, offset).Find(_Node_body)
	return _Node_body, err
}

// GetNode_bodiesByEntityIdAndDeltaAndBodyFormat Get Node_bodies via EntityIdAndDeltaAndBodyFormat
func GetNode_bodiesByEntityIdAndDeltaAndBodyFormat(offset int, limit int, EntityId_ int, Delta_ int, BodyFormat_ string) (*[]*Node_body, error) {
	var _Node_body = new([]*Node_body)
	err := Engine.Table("node__body").Where("entity_id = ? and delta = ? and body_format = ?", EntityId_, Delta_, BodyFormat_).Limit(limit, offset).Find(_Node_body)
	return _Node_body, err
}

// GetNode_bodiesByEntityIdAndBodyValueAndBodySummary Get Node_bodies via EntityIdAndBodyValueAndBodySummary
func GetNode_bodiesByEntityIdAndBodyValueAndBodySummary(offset int, limit int, EntityId_ int, BodyValue_ string, BodySummary_ string) (*[]*Node_body, error) {
	var _Node_body = new([]*Node_body)
	err := Engine.Table("node__body").Where("entity_id = ? and body_value = ? and body_summary = ?", EntityId_, BodyValue_, BodySummary_).Limit(limit, offset).Find(_Node_body)
	return _Node_body, err
}

// GetNode_bodiesByEntityIdAndBodyValueAndBodyFormat Get Node_bodies via EntityIdAndBodyValueAndBodyFormat
func GetNode_bodiesByEntityIdAndBodyValueAndBodyFormat(offset int, limit int, EntityId_ int, BodyValue_ string, BodyFormat_ string) (*[]*Node_body, error) {
	var _Node_body = new([]*Node_body)
	err := Engine.Table("node__body").Where("entity_id = ? and body_value = ? and body_format = ?", EntityId_, BodyValue_, BodyFormat_).Limit(limit, offset).Find(_Node_body)
	return _Node_body, err
}

// GetNode_bodiesByEntityIdAndBodySummaryAndBodyFormat Get Node_bodies via EntityIdAndBodySummaryAndBodyFormat
func GetNode_bodiesByEntityIdAndBodySummaryAndBodyFormat(offset int, limit int, EntityId_ int, BodySummary_ string, BodyFormat_ string) (*[]*Node_body, error) {
	var _Node_body = new([]*Node_body)
	err := Engine.Table("node__body").Where("entity_id = ? and body_summary = ? and body_format = ?", EntityId_, BodySummary_, BodyFormat_).Limit(limit, offset).Find(_Node_body)
	return _Node_body, err
}

// GetNode_bodiesByRevisionIdAndLangcodeAndDelta Get Node_bodies via RevisionIdAndLangcodeAndDelta
func GetNode_bodiesByRevisionIdAndLangcodeAndDelta(offset int, limit int, RevisionId_ int, Langcode_ string, Delta_ int) (*[]*Node_body, error) {
	var _Node_body = new([]*Node_body)
	err := Engine.Table("node__body").Where("revision_id = ? and langcode = ? and delta = ?", RevisionId_, Langcode_, Delta_).Limit(limit, offset).Find(_Node_body)
	return _Node_body, err
}

// GetNode_bodiesByRevisionIdAndLangcodeAndBodyValue Get Node_bodies via RevisionIdAndLangcodeAndBodyValue
func GetNode_bodiesByRevisionIdAndLangcodeAndBodyValue(offset int, limit int, RevisionId_ int, Langcode_ string, BodyValue_ string) (*[]*Node_body, error) {
	var _Node_body = new([]*Node_body)
	err := Engine.Table("node__body").Where("revision_id = ? and langcode = ? and body_value = ?", RevisionId_, Langcode_, BodyValue_).Limit(limit, offset).Find(_Node_body)
	return _Node_body, err
}

// GetNode_bodiesByRevisionIdAndLangcodeAndBodySummary Get Node_bodies via RevisionIdAndLangcodeAndBodySummary
func GetNode_bodiesByRevisionIdAndLangcodeAndBodySummary(offset int, limit int, RevisionId_ int, Langcode_ string, BodySummary_ string) (*[]*Node_body, error) {
	var _Node_body = new([]*Node_body)
	err := Engine.Table("node__body").Where("revision_id = ? and langcode = ? and body_summary = ?", RevisionId_, Langcode_, BodySummary_).Limit(limit, offset).Find(_Node_body)
	return _Node_body, err
}

// GetNode_bodiesByRevisionIdAndLangcodeAndBodyFormat Get Node_bodies via RevisionIdAndLangcodeAndBodyFormat
func GetNode_bodiesByRevisionIdAndLangcodeAndBodyFormat(offset int, limit int, RevisionId_ int, Langcode_ string, BodyFormat_ string) (*[]*Node_body, error) {
	var _Node_body = new([]*Node_body)
	err := Engine.Table("node__body").Where("revision_id = ? and langcode = ? and body_format = ?", RevisionId_, Langcode_, BodyFormat_).Limit(limit, offset).Find(_Node_body)
	return _Node_body, err
}

// GetNode_bodiesByRevisionIdAndDeltaAndBodyValue Get Node_bodies via RevisionIdAndDeltaAndBodyValue
func GetNode_bodiesByRevisionIdAndDeltaAndBodyValue(offset int, limit int, RevisionId_ int, Delta_ int, BodyValue_ string) (*[]*Node_body, error) {
	var _Node_body = new([]*Node_body)
	err := Engine.Table("node__body").Where("revision_id = ? and delta = ? and body_value = ?", RevisionId_, Delta_, BodyValue_).Limit(limit, offset).Find(_Node_body)
	return _Node_body, err
}

// GetNode_bodiesByRevisionIdAndDeltaAndBodySummary Get Node_bodies via RevisionIdAndDeltaAndBodySummary
func GetNode_bodiesByRevisionIdAndDeltaAndBodySummary(offset int, limit int, RevisionId_ int, Delta_ int, BodySummary_ string) (*[]*Node_body, error) {
	var _Node_body = new([]*Node_body)
	err := Engine.Table("node__body").Where("revision_id = ? and delta = ? and body_summary = ?", RevisionId_, Delta_, BodySummary_).Limit(limit, offset).Find(_Node_body)
	return _Node_body, err
}

// GetNode_bodiesByRevisionIdAndDeltaAndBodyFormat Get Node_bodies via RevisionIdAndDeltaAndBodyFormat
func GetNode_bodiesByRevisionIdAndDeltaAndBodyFormat(offset int, limit int, RevisionId_ int, Delta_ int, BodyFormat_ string) (*[]*Node_body, error) {
	var _Node_body = new([]*Node_body)
	err := Engine.Table("node__body").Where("revision_id = ? and delta = ? and body_format = ?", RevisionId_, Delta_, BodyFormat_).Limit(limit, offset).Find(_Node_body)
	return _Node_body, err
}

// GetNode_bodiesByRevisionIdAndBodyValueAndBodySummary Get Node_bodies via RevisionIdAndBodyValueAndBodySummary
func GetNode_bodiesByRevisionIdAndBodyValueAndBodySummary(offset int, limit int, RevisionId_ int, BodyValue_ string, BodySummary_ string) (*[]*Node_body, error) {
	var _Node_body = new([]*Node_body)
	err := Engine.Table("node__body").Where("revision_id = ? and body_value = ? and body_summary = ?", RevisionId_, BodyValue_, BodySummary_).Limit(limit, offset).Find(_Node_body)
	return _Node_body, err
}

// GetNode_bodiesByRevisionIdAndBodyValueAndBodyFormat Get Node_bodies via RevisionIdAndBodyValueAndBodyFormat
func GetNode_bodiesByRevisionIdAndBodyValueAndBodyFormat(offset int, limit int, RevisionId_ int, BodyValue_ string, BodyFormat_ string) (*[]*Node_body, error) {
	var _Node_body = new([]*Node_body)
	err := Engine.Table("node__body").Where("revision_id = ? and body_value = ? and body_format = ?", RevisionId_, BodyValue_, BodyFormat_).Limit(limit, offset).Find(_Node_body)
	return _Node_body, err
}

// GetNode_bodiesByRevisionIdAndBodySummaryAndBodyFormat Get Node_bodies via RevisionIdAndBodySummaryAndBodyFormat
func GetNode_bodiesByRevisionIdAndBodySummaryAndBodyFormat(offset int, limit int, RevisionId_ int, BodySummary_ string, BodyFormat_ string) (*[]*Node_body, error) {
	var _Node_body = new([]*Node_body)
	err := Engine.Table("node__body").Where("revision_id = ? and body_summary = ? and body_format = ?", RevisionId_, BodySummary_, BodyFormat_).Limit(limit, offset).Find(_Node_body)
	return _Node_body, err
}

// GetNode_bodiesByLangcodeAndDeltaAndBodyValue Get Node_bodies via LangcodeAndDeltaAndBodyValue
func GetNode_bodiesByLangcodeAndDeltaAndBodyValue(offset int, limit int, Langcode_ string, Delta_ int, BodyValue_ string) (*[]*Node_body, error) {
	var _Node_body = new([]*Node_body)
	err := Engine.Table("node__body").Where("langcode = ? and delta = ? and body_value = ?", Langcode_, Delta_, BodyValue_).Limit(limit, offset).Find(_Node_body)
	return _Node_body, err
}

// GetNode_bodiesByLangcodeAndDeltaAndBodySummary Get Node_bodies via LangcodeAndDeltaAndBodySummary
func GetNode_bodiesByLangcodeAndDeltaAndBodySummary(offset int, limit int, Langcode_ string, Delta_ int, BodySummary_ string) (*[]*Node_body, error) {
	var _Node_body = new([]*Node_body)
	err := Engine.Table("node__body").Where("langcode = ? and delta = ? and body_summary = ?", Langcode_, Delta_, BodySummary_).Limit(limit, offset).Find(_Node_body)
	return _Node_body, err
}

// GetNode_bodiesByLangcodeAndDeltaAndBodyFormat Get Node_bodies via LangcodeAndDeltaAndBodyFormat
func GetNode_bodiesByLangcodeAndDeltaAndBodyFormat(offset int, limit int, Langcode_ string, Delta_ int, BodyFormat_ string) (*[]*Node_body, error) {
	var _Node_body = new([]*Node_body)
	err := Engine.Table("node__body").Where("langcode = ? and delta = ? and body_format = ?", Langcode_, Delta_, BodyFormat_).Limit(limit, offset).Find(_Node_body)
	return _Node_body, err
}

// GetNode_bodiesByLangcodeAndBodyValueAndBodySummary Get Node_bodies via LangcodeAndBodyValueAndBodySummary
func GetNode_bodiesByLangcodeAndBodyValueAndBodySummary(offset int, limit int, Langcode_ string, BodyValue_ string, BodySummary_ string) (*[]*Node_body, error) {
	var _Node_body = new([]*Node_body)
	err := Engine.Table("node__body").Where("langcode = ? and body_value = ? and body_summary = ?", Langcode_, BodyValue_, BodySummary_).Limit(limit, offset).Find(_Node_body)
	return _Node_body, err
}

// GetNode_bodiesByLangcodeAndBodyValueAndBodyFormat Get Node_bodies via LangcodeAndBodyValueAndBodyFormat
func GetNode_bodiesByLangcodeAndBodyValueAndBodyFormat(offset int, limit int, Langcode_ string, BodyValue_ string, BodyFormat_ string) (*[]*Node_body, error) {
	var _Node_body = new([]*Node_body)
	err := Engine.Table("node__body").Where("langcode = ? and body_value = ? and body_format = ?", Langcode_, BodyValue_, BodyFormat_).Limit(limit, offset).Find(_Node_body)
	return _Node_body, err
}

// GetNode_bodiesByLangcodeAndBodySummaryAndBodyFormat Get Node_bodies via LangcodeAndBodySummaryAndBodyFormat
func GetNode_bodiesByLangcodeAndBodySummaryAndBodyFormat(offset int, limit int, Langcode_ string, BodySummary_ string, BodyFormat_ string) (*[]*Node_body, error) {
	var _Node_body = new([]*Node_body)
	err := Engine.Table("node__body").Where("langcode = ? and body_summary = ? and body_format = ?", Langcode_, BodySummary_, BodyFormat_).Limit(limit, offset).Find(_Node_body)
	return _Node_body, err
}

// GetNode_bodiesByDeltaAndBodyValueAndBodySummary Get Node_bodies via DeltaAndBodyValueAndBodySummary
func GetNode_bodiesByDeltaAndBodyValueAndBodySummary(offset int, limit int, Delta_ int, BodyValue_ string, BodySummary_ string) (*[]*Node_body, error) {
	var _Node_body = new([]*Node_body)
	err := Engine.Table("node__body").Where("delta = ? and body_value = ? and body_summary = ?", Delta_, BodyValue_, BodySummary_).Limit(limit, offset).Find(_Node_body)
	return _Node_body, err
}

// GetNode_bodiesByDeltaAndBodyValueAndBodyFormat Get Node_bodies via DeltaAndBodyValueAndBodyFormat
func GetNode_bodiesByDeltaAndBodyValueAndBodyFormat(offset int, limit int, Delta_ int, BodyValue_ string, BodyFormat_ string) (*[]*Node_body, error) {
	var _Node_body = new([]*Node_body)
	err := Engine.Table("node__body").Where("delta = ? and body_value = ? and body_format = ?", Delta_, BodyValue_, BodyFormat_).Limit(limit, offset).Find(_Node_body)
	return _Node_body, err
}

// GetNode_bodiesByDeltaAndBodySummaryAndBodyFormat Get Node_bodies via DeltaAndBodySummaryAndBodyFormat
func GetNode_bodiesByDeltaAndBodySummaryAndBodyFormat(offset int, limit int, Delta_ int, BodySummary_ string, BodyFormat_ string) (*[]*Node_body, error) {
	var _Node_body = new([]*Node_body)
	err := Engine.Table("node__body").Where("delta = ? and body_summary = ? and body_format = ?", Delta_, BodySummary_, BodyFormat_).Limit(limit, offset).Find(_Node_body)
	return _Node_body, err
}

// GetNode_bodiesByBodyValueAndBodySummaryAndBodyFormat Get Node_bodies via BodyValueAndBodySummaryAndBodyFormat
func GetNode_bodiesByBodyValueAndBodySummaryAndBodyFormat(offset int, limit int, BodyValue_ string, BodySummary_ string, BodyFormat_ string) (*[]*Node_body, error) {
	var _Node_body = new([]*Node_body)
	err := Engine.Table("node__body").Where("body_value = ? and body_summary = ? and body_format = ?", BodyValue_, BodySummary_, BodyFormat_).Limit(limit, offset).Find(_Node_body)
	return _Node_body, err
}

// GetNode_bodiesByBundleAndDeleted Get Node_bodies via BundleAndDeleted
func GetNode_bodiesByBundleAndDeleted(offset int, limit int, Bundle_ string, Deleted_ int) (*[]*Node_body, error) {
	var _Node_body = new([]*Node_body)
	err := Engine.Table("node__body").Where("bundle = ? and deleted = ?", Bundle_, Deleted_).Limit(limit, offset).Find(_Node_body)
	return _Node_body, err
}

// GetNode_bodiesByBundleAndEntityId Get Node_bodies via BundleAndEntityId
func GetNode_bodiesByBundleAndEntityId(offset int, limit int, Bundle_ string, EntityId_ int) (*[]*Node_body, error) {
	var _Node_body = new([]*Node_body)
	err := Engine.Table("node__body").Where("bundle = ? and entity_id = ?", Bundle_, EntityId_).Limit(limit, offset).Find(_Node_body)
	return _Node_body, err
}

// GetNode_bodiesByBundleAndRevisionId Get Node_bodies via BundleAndRevisionId
func GetNode_bodiesByBundleAndRevisionId(offset int, limit int, Bundle_ string, RevisionId_ int) (*[]*Node_body, error) {
	var _Node_body = new([]*Node_body)
	err := Engine.Table("node__body").Where("bundle = ? and revision_id = ?", Bundle_, RevisionId_).Limit(limit, offset).Find(_Node_body)
	return _Node_body, err
}

// GetNode_bodiesByBundleAndLangcode Get Node_bodies via BundleAndLangcode
func GetNode_bodiesByBundleAndLangcode(offset int, limit int, Bundle_ string, Langcode_ string) (*[]*Node_body, error) {
	var _Node_body = new([]*Node_body)
	err := Engine.Table("node__body").Where("bundle = ? and langcode = ?", Bundle_, Langcode_).Limit(limit, offset).Find(_Node_body)
	return _Node_body, err
}

// GetNode_bodiesByBundleAndDelta Get Node_bodies via BundleAndDelta
func GetNode_bodiesByBundleAndDelta(offset int, limit int, Bundle_ string, Delta_ int) (*[]*Node_body, error) {
	var _Node_body = new([]*Node_body)
	err := Engine.Table("node__body").Where("bundle = ? and delta = ?", Bundle_, Delta_).Limit(limit, offset).Find(_Node_body)
	return _Node_body, err
}

// GetNode_bodiesByBundleAndBodyValue Get Node_bodies via BundleAndBodyValue
func GetNode_bodiesByBundleAndBodyValue(offset int, limit int, Bundle_ string, BodyValue_ string) (*[]*Node_body, error) {
	var _Node_body = new([]*Node_body)
	err := Engine.Table("node__body").Where("bundle = ? and body_value = ?", Bundle_, BodyValue_).Limit(limit, offset).Find(_Node_body)
	return _Node_body, err
}

// GetNode_bodiesByBundleAndBodySummary Get Node_bodies via BundleAndBodySummary
func GetNode_bodiesByBundleAndBodySummary(offset int, limit int, Bundle_ string, BodySummary_ string) (*[]*Node_body, error) {
	var _Node_body = new([]*Node_body)
	err := Engine.Table("node__body").Where("bundle = ? and body_summary = ?", Bundle_, BodySummary_).Limit(limit, offset).Find(_Node_body)
	return _Node_body, err
}

// GetNode_bodiesByBundleAndBodyFormat Get Node_bodies via BundleAndBodyFormat
func GetNode_bodiesByBundleAndBodyFormat(offset int, limit int, Bundle_ string, BodyFormat_ string) (*[]*Node_body, error) {
	var _Node_body = new([]*Node_body)
	err := Engine.Table("node__body").Where("bundle = ? and body_format = ?", Bundle_, BodyFormat_).Limit(limit, offset).Find(_Node_body)
	return _Node_body, err
}

// GetNode_bodiesByDeletedAndEntityId Get Node_bodies via DeletedAndEntityId
func GetNode_bodiesByDeletedAndEntityId(offset int, limit int, Deleted_ int, EntityId_ int) (*[]*Node_body, error) {
	var _Node_body = new([]*Node_body)
	err := Engine.Table("node__body").Where("deleted = ? and entity_id = ?", Deleted_, EntityId_).Limit(limit, offset).Find(_Node_body)
	return _Node_body, err
}

// GetNode_bodiesByDeletedAndRevisionId Get Node_bodies via DeletedAndRevisionId
func GetNode_bodiesByDeletedAndRevisionId(offset int, limit int, Deleted_ int, RevisionId_ int) (*[]*Node_body, error) {
	var _Node_body = new([]*Node_body)
	err := Engine.Table("node__body").Where("deleted = ? and revision_id = ?", Deleted_, RevisionId_).Limit(limit, offset).Find(_Node_body)
	return _Node_body, err
}

// GetNode_bodiesByDeletedAndLangcode Get Node_bodies via DeletedAndLangcode
func GetNode_bodiesByDeletedAndLangcode(offset int, limit int, Deleted_ int, Langcode_ string) (*[]*Node_body, error) {
	var _Node_body = new([]*Node_body)
	err := Engine.Table("node__body").Where("deleted = ? and langcode = ?", Deleted_, Langcode_).Limit(limit, offset).Find(_Node_body)
	return _Node_body, err
}

// GetNode_bodiesByDeletedAndDelta Get Node_bodies via DeletedAndDelta
func GetNode_bodiesByDeletedAndDelta(offset int, limit int, Deleted_ int, Delta_ int) (*[]*Node_body, error) {
	var _Node_body = new([]*Node_body)
	err := Engine.Table("node__body").Where("deleted = ? and delta = ?", Deleted_, Delta_).Limit(limit, offset).Find(_Node_body)
	return _Node_body, err
}

// GetNode_bodiesByDeletedAndBodyValue Get Node_bodies via DeletedAndBodyValue
func GetNode_bodiesByDeletedAndBodyValue(offset int, limit int, Deleted_ int, BodyValue_ string) (*[]*Node_body, error) {
	var _Node_body = new([]*Node_body)
	err := Engine.Table("node__body").Where("deleted = ? and body_value = ?", Deleted_, BodyValue_).Limit(limit, offset).Find(_Node_body)
	return _Node_body, err
}

// GetNode_bodiesByDeletedAndBodySummary Get Node_bodies via DeletedAndBodySummary
func GetNode_bodiesByDeletedAndBodySummary(offset int, limit int, Deleted_ int, BodySummary_ string) (*[]*Node_body, error) {
	var _Node_body = new([]*Node_body)
	err := Engine.Table("node__body").Where("deleted = ? and body_summary = ?", Deleted_, BodySummary_).Limit(limit, offset).Find(_Node_body)
	return _Node_body, err
}

// GetNode_bodiesByDeletedAndBodyFormat Get Node_bodies via DeletedAndBodyFormat
func GetNode_bodiesByDeletedAndBodyFormat(offset int, limit int, Deleted_ int, BodyFormat_ string) (*[]*Node_body, error) {
	var _Node_body = new([]*Node_body)
	err := Engine.Table("node__body").Where("deleted = ? and body_format = ?", Deleted_, BodyFormat_).Limit(limit, offset).Find(_Node_body)
	return _Node_body, err
}

// GetNode_bodiesByEntityIdAndRevisionId Get Node_bodies via EntityIdAndRevisionId
func GetNode_bodiesByEntityIdAndRevisionId(offset int, limit int, EntityId_ int, RevisionId_ int) (*[]*Node_body, error) {
	var _Node_body = new([]*Node_body)
	err := Engine.Table("node__body").Where("entity_id = ? and revision_id = ?", EntityId_, RevisionId_).Limit(limit, offset).Find(_Node_body)
	return _Node_body, err
}

// GetNode_bodiesByEntityIdAndLangcode Get Node_bodies via EntityIdAndLangcode
func GetNode_bodiesByEntityIdAndLangcode(offset int, limit int, EntityId_ int, Langcode_ string) (*[]*Node_body, error) {
	var _Node_body = new([]*Node_body)
	err := Engine.Table("node__body").Where("entity_id = ? and langcode = ?", EntityId_, Langcode_).Limit(limit, offset).Find(_Node_body)
	return _Node_body, err
}

// GetNode_bodiesByEntityIdAndDelta Get Node_bodies via EntityIdAndDelta
func GetNode_bodiesByEntityIdAndDelta(offset int, limit int, EntityId_ int, Delta_ int) (*[]*Node_body, error) {
	var _Node_body = new([]*Node_body)
	err := Engine.Table("node__body").Where("entity_id = ? and delta = ?", EntityId_, Delta_).Limit(limit, offset).Find(_Node_body)
	return _Node_body, err
}

// GetNode_bodiesByEntityIdAndBodyValue Get Node_bodies via EntityIdAndBodyValue
func GetNode_bodiesByEntityIdAndBodyValue(offset int, limit int, EntityId_ int, BodyValue_ string) (*[]*Node_body, error) {
	var _Node_body = new([]*Node_body)
	err := Engine.Table("node__body").Where("entity_id = ? and body_value = ?", EntityId_, BodyValue_).Limit(limit, offset).Find(_Node_body)
	return _Node_body, err
}

// GetNode_bodiesByEntityIdAndBodySummary Get Node_bodies via EntityIdAndBodySummary
func GetNode_bodiesByEntityIdAndBodySummary(offset int, limit int, EntityId_ int, BodySummary_ string) (*[]*Node_body, error) {
	var _Node_body = new([]*Node_body)
	err := Engine.Table("node__body").Where("entity_id = ? and body_summary = ?", EntityId_, BodySummary_).Limit(limit, offset).Find(_Node_body)
	return _Node_body, err
}

// GetNode_bodiesByEntityIdAndBodyFormat Get Node_bodies via EntityIdAndBodyFormat
func GetNode_bodiesByEntityIdAndBodyFormat(offset int, limit int, EntityId_ int, BodyFormat_ string) (*[]*Node_body, error) {
	var _Node_body = new([]*Node_body)
	err := Engine.Table("node__body").Where("entity_id = ? and body_format = ?", EntityId_, BodyFormat_).Limit(limit, offset).Find(_Node_body)
	return _Node_body, err
}

// GetNode_bodiesByRevisionIdAndLangcode Get Node_bodies via RevisionIdAndLangcode
func GetNode_bodiesByRevisionIdAndLangcode(offset int, limit int, RevisionId_ int, Langcode_ string) (*[]*Node_body, error) {
	var _Node_body = new([]*Node_body)
	err := Engine.Table("node__body").Where("revision_id = ? and langcode = ?", RevisionId_, Langcode_).Limit(limit, offset).Find(_Node_body)
	return _Node_body, err
}

// GetNode_bodiesByRevisionIdAndDelta Get Node_bodies via RevisionIdAndDelta
func GetNode_bodiesByRevisionIdAndDelta(offset int, limit int, RevisionId_ int, Delta_ int) (*[]*Node_body, error) {
	var _Node_body = new([]*Node_body)
	err := Engine.Table("node__body").Where("revision_id = ? and delta = ?", RevisionId_, Delta_).Limit(limit, offset).Find(_Node_body)
	return _Node_body, err
}

// GetNode_bodiesByRevisionIdAndBodyValue Get Node_bodies via RevisionIdAndBodyValue
func GetNode_bodiesByRevisionIdAndBodyValue(offset int, limit int, RevisionId_ int, BodyValue_ string) (*[]*Node_body, error) {
	var _Node_body = new([]*Node_body)
	err := Engine.Table("node__body").Where("revision_id = ? and body_value = ?", RevisionId_, BodyValue_).Limit(limit, offset).Find(_Node_body)
	return _Node_body, err
}

// GetNode_bodiesByRevisionIdAndBodySummary Get Node_bodies via RevisionIdAndBodySummary
func GetNode_bodiesByRevisionIdAndBodySummary(offset int, limit int, RevisionId_ int, BodySummary_ string) (*[]*Node_body, error) {
	var _Node_body = new([]*Node_body)
	err := Engine.Table("node__body").Where("revision_id = ? and body_summary = ?", RevisionId_, BodySummary_).Limit(limit, offset).Find(_Node_body)
	return _Node_body, err
}

// GetNode_bodiesByRevisionIdAndBodyFormat Get Node_bodies via RevisionIdAndBodyFormat
func GetNode_bodiesByRevisionIdAndBodyFormat(offset int, limit int, RevisionId_ int, BodyFormat_ string) (*[]*Node_body, error) {
	var _Node_body = new([]*Node_body)
	err := Engine.Table("node__body").Where("revision_id = ? and body_format = ?", RevisionId_, BodyFormat_).Limit(limit, offset).Find(_Node_body)
	return _Node_body, err
}

// GetNode_bodiesByLangcodeAndDelta Get Node_bodies via LangcodeAndDelta
func GetNode_bodiesByLangcodeAndDelta(offset int, limit int, Langcode_ string, Delta_ int) (*[]*Node_body, error) {
	var _Node_body = new([]*Node_body)
	err := Engine.Table("node__body").Where("langcode = ? and delta = ?", Langcode_, Delta_).Limit(limit, offset).Find(_Node_body)
	return _Node_body, err
}

// GetNode_bodiesByLangcodeAndBodyValue Get Node_bodies via LangcodeAndBodyValue
func GetNode_bodiesByLangcodeAndBodyValue(offset int, limit int, Langcode_ string, BodyValue_ string) (*[]*Node_body, error) {
	var _Node_body = new([]*Node_body)
	err := Engine.Table("node__body").Where("langcode = ? and body_value = ?", Langcode_, BodyValue_).Limit(limit, offset).Find(_Node_body)
	return _Node_body, err
}

// GetNode_bodiesByLangcodeAndBodySummary Get Node_bodies via LangcodeAndBodySummary
func GetNode_bodiesByLangcodeAndBodySummary(offset int, limit int, Langcode_ string, BodySummary_ string) (*[]*Node_body, error) {
	var _Node_body = new([]*Node_body)
	err := Engine.Table("node__body").Where("langcode = ? and body_summary = ?", Langcode_, BodySummary_).Limit(limit, offset).Find(_Node_body)
	return _Node_body, err
}

// GetNode_bodiesByLangcodeAndBodyFormat Get Node_bodies via LangcodeAndBodyFormat
func GetNode_bodiesByLangcodeAndBodyFormat(offset int, limit int, Langcode_ string, BodyFormat_ string) (*[]*Node_body, error) {
	var _Node_body = new([]*Node_body)
	err := Engine.Table("node__body").Where("langcode = ? and body_format = ?", Langcode_, BodyFormat_).Limit(limit, offset).Find(_Node_body)
	return _Node_body, err
}

// GetNode_bodiesByDeltaAndBodyValue Get Node_bodies via DeltaAndBodyValue
func GetNode_bodiesByDeltaAndBodyValue(offset int, limit int, Delta_ int, BodyValue_ string) (*[]*Node_body, error) {
	var _Node_body = new([]*Node_body)
	err := Engine.Table("node__body").Where("delta = ? and body_value = ?", Delta_, BodyValue_).Limit(limit, offset).Find(_Node_body)
	return _Node_body, err
}

// GetNode_bodiesByDeltaAndBodySummary Get Node_bodies via DeltaAndBodySummary
func GetNode_bodiesByDeltaAndBodySummary(offset int, limit int, Delta_ int, BodySummary_ string) (*[]*Node_body, error) {
	var _Node_body = new([]*Node_body)
	err := Engine.Table("node__body").Where("delta = ? and body_summary = ?", Delta_, BodySummary_).Limit(limit, offset).Find(_Node_body)
	return _Node_body, err
}

// GetNode_bodiesByDeltaAndBodyFormat Get Node_bodies via DeltaAndBodyFormat
func GetNode_bodiesByDeltaAndBodyFormat(offset int, limit int, Delta_ int, BodyFormat_ string) (*[]*Node_body, error) {
	var _Node_body = new([]*Node_body)
	err := Engine.Table("node__body").Where("delta = ? and body_format = ?", Delta_, BodyFormat_).Limit(limit, offset).Find(_Node_body)
	return _Node_body, err
}

// GetNode_bodiesByBodyValueAndBodySummary Get Node_bodies via BodyValueAndBodySummary
func GetNode_bodiesByBodyValueAndBodySummary(offset int, limit int, BodyValue_ string, BodySummary_ string) (*[]*Node_body, error) {
	var _Node_body = new([]*Node_body)
	err := Engine.Table("node__body").Where("body_value = ? and body_summary = ?", BodyValue_, BodySummary_).Limit(limit, offset).Find(_Node_body)
	return _Node_body, err
}

// GetNode_bodiesByBodyValueAndBodyFormat Get Node_bodies via BodyValueAndBodyFormat
func GetNode_bodiesByBodyValueAndBodyFormat(offset int, limit int, BodyValue_ string, BodyFormat_ string) (*[]*Node_body, error) {
	var _Node_body = new([]*Node_body)
	err := Engine.Table("node__body").Where("body_value = ? and body_format = ?", BodyValue_, BodyFormat_).Limit(limit, offset).Find(_Node_body)
	return _Node_body, err
}

// GetNode_bodiesByBodySummaryAndBodyFormat Get Node_bodies via BodySummaryAndBodyFormat
func GetNode_bodiesByBodySummaryAndBodyFormat(offset int, limit int, BodySummary_ string, BodyFormat_ string) (*[]*Node_body, error) {
	var _Node_body = new([]*Node_body)
	err := Engine.Table("node__body").Where("body_summary = ? and body_format = ?", BodySummary_, BodyFormat_).Limit(limit, offset).Find(_Node_body)
	return _Node_body, err
}

// GetNode_bodies Get Node_bodies via field
func GetNode_bodies(offset int, limit int, field string) (*[]*Node_body, error) {
	var _Node_body = new([]*Node_body)
	err := Engine.Table("node__body").Limit(limit, offset).Desc(field).Find(_Node_body)
	return _Node_body, err
}

// GetNode_bodiesViaBundle Get Node_bodys via Bundle
func GetNode_bodiesViaBundle(offset int, limit int, Bundle_ string, field string) (*[]*Node_body, error) {
	var _Node_body = new([]*Node_body)
	err := Engine.Table("node__body").Where("bundle = ?", Bundle_).Limit(limit, offset).Desc(field).Find(_Node_body)
	return _Node_body, err
}

// GetNode_bodiesViaDeleted Get Node_bodys via Deleted
func GetNode_bodiesViaDeleted(offset int, limit int, Deleted_ int, field string) (*[]*Node_body, error) {
	var _Node_body = new([]*Node_body)
	err := Engine.Table("node__body").Where("deleted = ?", Deleted_).Limit(limit, offset).Desc(field).Find(_Node_body)
	return _Node_body, err
}

// GetNode_bodiesViaEntityId Get Node_bodys via EntityId
func GetNode_bodiesViaEntityId(offset int, limit int, EntityId_ int, field string) (*[]*Node_body, error) {
	var _Node_body = new([]*Node_body)
	err := Engine.Table("node__body").Where("entity_id = ?", EntityId_).Limit(limit, offset).Desc(field).Find(_Node_body)
	return _Node_body, err
}

// GetNode_bodiesViaRevisionId Get Node_bodys via RevisionId
func GetNode_bodiesViaRevisionId(offset int, limit int, RevisionId_ int, field string) (*[]*Node_body, error) {
	var _Node_body = new([]*Node_body)
	err := Engine.Table("node__body").Where("revision_id = ?", RevisionId_).Limit(limit, offset).Desc(field).Find(_Node_body)
	return _Node_body, err
}

// GetNode_bodiesViaLangcode Get Node_bodys via Langcode
func GetNode_bodiesViaLangcode(offset int, limit int, Langcode_ string, field string) (*[]*Node_body, error) {
	var _Node_body = new([]*Node_body)
	err := Engine.Table("node__body").Where("langcode = ?", Langcode_).Limit(limit, offset).Desc(field).Find(_Node_body)
	return _Node_body, err
}

// GetNode_bodiesViaDelta Get Node_bodys via Delta
func GetNode_bodiesViaDelta(offset int, limit int, Delta_ int, field string) (*[]*Node_body, error) {
	var _Node_body = new([]*Node_body)
	err := Engine.Table("node__body").Where("delta = ?", Delta_).Limit(limit, offset).Desc(field).Find(_Node_body)
	return _Node_body, err
}

// GetNode_bodiesViaBodyValue Get Node_bodys via BodyValue
func GetNode_bodiesViaBodyValue(offset int, limit int, BodyValue_ string, field string) (*[]*Node_body, error) {
	var _Node_body = new([]*Node_body)
	err := Engine.Table("node__body").Where("body_value = ?", BodyValue_).Limit(limit, offset).Desc(field).Find(_Node_body)
	return _Node_body, err
}

// GetNode_bodiesViaBodySummary Get Node_bodys via BodySummary
func GetNode_bodiesViaBodySummary(offset int, limit int, BodySummary_ string, field string) (*[]*Node_body, error) {
	var _Node_body = new([]*Node_body)
	err := Engine.Table("node__body").Where("body_summary = ?", BodySummary_).Limit(limit, offset).Desc(field).Find(_Node_body)
	return _Node_body, err
}

// GetNode_bodiesViaBodyFormat Get Node_bodys via BodyFormat
func GetNode_bodiesViaBodyFormat(offset int, limit int, BodyFormat_ string, field string) (*[]*Node_body, error) {
	var _Node_body = new([]*Node_body)
	err := Engine.Table("node__body").Where("body_format = ?", BodyFormat_).Limit(limit, offset).Desc(field).Find(_Node_body)
	return _Node_body, err
}

// HasNode_bodyViaBundle Has Node_body via Bundle
func HasNode_bodyViaBundle(iBundle string) bool {
	if has, err := Engine.Where("bundle = ?", iBundle).Get(new(Node_body)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasNode_bodyViaDeleted Has Node_body via Deleted
func HasNode_bodyViaDeleted(iDeleted int) bool {
	if has, err := Engine.Where("deleted = ?", iDeleted).Get(new(Node_body)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasNode_bodyViaEntityId Has Node_body via EntityId
func HasNode_bodyViaEntityId(iEntityId int) bool {
	if has, err := Engine.Where("entity_id = ?", iEntityId).Get(new(Node_body)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasNode_bodyViaRevisionId Has Node_body via RevisionId
func HasNode_bodyViaRevisionId(iRevisionId int) bool {
	if has, err := Engine.Where("revision_id = ?", iRevisionId).Get(new(Node_body)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasNode_bodyViaLangcode Has Node_body via Langcode
func HasNode_bodyViaLangcode(iLangcode string) bool {
	if has, err := Engine.Where("langcode = ?", iLangcode).Get(new(Node_body)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasNode_bodyViaDelta Has Node_body via Delta
func HasNode_bodyViaDelta(iDelta int) bool {
	if has, err := Engine.Where("delta = ?", iDelta).Get(new(Node_body)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasNode_bodyViaBodyValue Has Node_body via BodyValue
func HasNode_bodyViaBodyValue(iBodyValue string) bool {
	if has, err := Engine.Where("body_value = ?", iBodyValue).Get(new(Node_body)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasNode_bodyViaBodySummary Has Node_body via BodySummary
func HasNode_bodyViaBodySummary(iBodySummary string) bool {
	if has, err := Engine.Where("body_summary = ?", iBodySummary).Get(new(Node_body)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasNode_bodyViaBodyFormat Has Node_body via BodyFormat
func HasNode_bodyViaBodyFormat(iBodyFormat string) bool {
	if has, err := Engine.Where("body_format = ?", iBodyFormat).Get(new(Node_body)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// GetNode_bodyViaBundle Get Node_body via Bundle
func GetNode_bodyViaBundle(iBundle string) (*Node_body, error) {
	var _Node_body = &Node_body{Bundle: iBundle}
	has, err := Engine.Get(_Node_body)
	if has {
		return _Node_body, err
	} else {
		return nil, err
	}
}

// GetNode_bodyViaDeleted Get Node_body via Deleted
func GetNode_bodyViaDeleted(iDeleted int) (*Node_body, error) {
	var _Node_body = &Node_body{Deleted: iDeleted}
	has, err := Engine.Get(_Node_body)
	if has {
		return _Node_body, err
	} else {
		return nil, err
	}
}

// GetNode_bodyViaEntityId Get Node_body via EntityId
func GetNode_bodyViaEntityId(iEntityId int) (*Node_body, error) {
	var _Node_body = &Node_body{EntityId: iEntityId}
	has, err := Engine.Get(_Node_body)
	if has {
		return _Node_body, err
	} else {
		return nil, err
	}
}

// GetNode_bodyViaRevisionId Get Node_body via RevisionId
func GetNode_bodyViaRevisionId(iRevisionId int) (*Node_body, error) {
	var _Node_body = &Node_body{RevisionId: iRevisionId}
	has, err := Engine.Get(_Node_body)
	if has {
		return _Node_body, err
	} else {
		return nil, err
	}
}

// GetNode_bodyViaLangcode Get Node_body via Langcode
func GetNode_bodyViaLangcode(iLangcode string) (*Node_body, error) {
	var _Node_body = &Node_body{Langcode: iLangcode}
	has, err := Engine.Get(_Node_body)
	if has {
		return _Node_body, err
	} else {
		return nil, err
	}
}

// GetNode_bodyViaDelta Get Node_body via Delta
func GetNode_bodyViaDelta(iDelta int) (*Node_body, error) {
	var _Node_body = &Node_body{Delta: iDelta}
	has, err := Engine.Get(_Node_body)
	if has {
		return _Node_body, err
	} else {
		return nil, err
	}
}

// GetNode_bodyViaBodyValue Get Node_body via BodyValue
func GetNode_bodyViaBodyValue(iBodyValue string) (*Node_body, error) {
	var _Node_body = &Node_body{BodyValue: iBodyValue}
	has, err := Engine.Get(_Node_body)
	if has {
		return _Node_body, err
	} else {
		return nil, err
	}
}

// GetNode_bodyViaBodySummary Get Node_body via BodySummary
func GetNode_bodyViaBodySummary(iBodySummary string) (*Node_body, error) {
	var _Node_body = &Node_body{BodySummary: iBodySummary}
	has, err := Engine.Get(_Node_body)
	if has {
		return _Node_body, err
	} else {
		return nil, err
	}
}

// GetNode_bodyViaBodyFormat Get Node_body via BodyFormat
func GetNode_bodyViaBodyFormat(iBodyFormat string) (*Node_body, error) {
	var _Node_body = &Node_body{BodyFormat: iBodyFormat}
	has, err := Engine.Get(_Node_body)
	if has {
		return _Node_body, err
	} else {
		return nil, err
	}
}

// SetNode_bodyViaBundle Set Node_body via Bundle
func SetNode_bodyViaBundle(iBundle string, node__body *Node_body) (int64, error) {
	node__body.Bundle = iBundle
	return Engine.Insert(node__body)
}

// SetNode_bodyViaDeleted Set Node_body via Deleted
func SetNode_bodyViaDeleted(iDeleted int, node__body *Node_body) (int64, error) {
	node__body.Deleted = iDeleted
	return Engine.Insert(node__body)
}

// SetNode_bodyViaEntityId Set Node_body via EntityId
func SetNode_bodyViaEntityId(iEntityId int, node__body *Node_body) (int64, error) {
	node__body.EntityId = iEntityId
	return Engine.Insert(node__body)
}

// SetNode_bodyViaRevisionId Set Node_body via RevisionId
func SetNode_bodyViaRevisionId(iRevisionId int, node__body *Node_body) (int64, error) {
	node__body.RevisionId = iRevisionId
	return Engine.Insert(node__body)
}

// SetNode_bodyViaLangcode Set Node_body via Langcode
func SetNode_bodyViaLangcode(iLangcode string, node__body *Node_body) (int64, error) {
	node__body.Langcode = iLangcode
	return Engine.Insert(node__body)
}

// SetNode_bodyViaDelta Set Node_body via Delta
func SetNode_bodyViaDelta(iDelta int, node__body *Node_body) (int64, error) {
	node__body.Delta = iDelta
	return Engine.Insert(node__body)
}

// SetNode_bodyViaBodyValue Set Node_body via BodyValue
func SetNode_bodyViaBodyValue(iBodyValue string, node__body *Node_body) (int64, error) {
	node__body.BodyValue = iBodyValue
	return Engine.Insert(node__body)
}

// SetNode_bodyViaBodySummary Set Node_body via BodySummary
func SetNode_bodyViaBodySummary(iBodySummary string, node__body *Node_body) (int64, error) {
	node__body.BodySummary = iBodySummary
	return Engine.Insert(node__body)
}

// SetNode_bodyViaBodyFormat Set Node_body via BodyFormat
func SetNode_bodyViaBodyFormat(iBodyFormat string, node__body *Node_body) (int64, error) {
	node__body.BodyFormat = iBodyFormat
	return Engine.Insert(node__body)
}

// AddNode_body Add Node_body via all columns
func AddNode_body(iBundle string, iDeleted int, iEntityId int, iRevisionId int, iLangcode string, iDelta int, iBodyValue string, iBodySummary string, iBodyFormat string) error {
	_Node_body := &Node_body{Bundle: iBundle, Deleted: iDeleted, EntityId: iEntityId, RevisionId: iRevisionId, Langcode: iLangcode, Delta: iDelta, BodyValue: iBodyValue, BodySummary: iBodySummary, BodyFormat: iBodyFormat}
	if _, err := Engine.Insert(_Node_body); err != nil {
		return err
	} else {
		return nil
	}
}

// PostNode_body Post Node_body via iNode_body
func PostNode_body(iNode_body *Node_body) (string, error) {
	_, err := Engine.Insert(iNode_body)
	return iNode_body.Bundle, err
}

// PutNode_body Put Node_body
func PutNode_body(iNode_body *Node_body) (string, error) {
	_, err := Engine.Id(iNode_body.Bundle).Update(iNode_body)
	return iNode_body.Bundle, err
}

// PutNode_bodyViaBundle Put Node_body via Bundle
func PutNode_bodyViaBundle(Bundle_ string, iNode_body *Node_body) (int64, error) {
	row, err := Engine.Update(iNode_body, &Node_body{Bundle: Bundle_})
	return row, err
}

// PutNode_bodyViaDeleted Put Node_body via Deleted
func PutNode_bodyViaDeleted(Deleted_ int, iNode_body *Node_body) (int64, error) {
	row, err := Engine.Update(iNode_body, &Node_body{Deleted: Deleted_})
	return row, err
}

// PutNode_bodyViaEntityId Put Node_body via EntityId
func PutNode_bodyViaEntityId(EntityId_ int, iNode_body *Node_body) (int64, error) {
	row, err := Engine.Update(iNode_body, &Node_body{EntityId: EntityId_})
	return row, err
}

// PutNode_bodyViaRevisionId Put Node_body via RevisionId
func PutNode_bodyViaRevisionId(RevisionId_ int, iNode_body *Node_body) (int64, error) {
	row, err := Engine.Update(iNode_body, &Node_body{RevisionId: RevisionId_})
	return row, err
}

// PutNode_bodyViaLangcode Put Node_body via Langcode
func PutNode_bodyViaLangcode(Langcode_ string, iNode_body *Node_body) (int64, error) {
	row, err := Engine.Update(iNode_body, &Node_body{Langcode: Langcode_})
	return row, err
}

// PutNode_bodyViaDelta Put Node_body via Delta
func PutNode_bodyViaDelta(Delta_ int, iNode_body *Node_body) (int64, error) {
	row, err := Engine.Update(iNode_body, &Node_body{Delta: Delta_})
	return row, err
}

// PutNode_bodyViaBodyValue Put Node_body via BodyValue
func PutNode_bodyViaBodyValue(BodyValue_ string, iNode_body *Node_body) (int64, error) {
	row, err := Engine.Update(iNode_body, &Node_body{BodyValue: BodyValue_})
	return row, err
}

// PutNode_bodyViaBodySummary Put Node_body via BodySummary
func PutNode_bodyViaBodySummary(BodySummary_ string, iNode_body *Node_body) (int64, error) {
	row, err := Engine.Update(iNode_body, &Node_body{BodySummary: BodySummary_})
	return row, err
}

// PutNode_bodyViaBodyFormat Put Node_body via BodyFormat
func PutNode_bodyViaBodyFormat(BodyFormat_ string, iNode_body *Node_body) (int64, error) {
	row, err := Engine.Update(iNode_body, &Node_body{BodyFormat: BodyFormat_})
	return row, err
}

// UpdateNode_bodyViaBundle via map[string]interface{}{}
func UpdateNode_bodyViaBundle(iBundle string, iNode_bodyMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Node_body)).Where("bundle = ?", iBundle).Update(iNode_bodyMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateNode_bodyViaDeleted via map[string]interface{}{}
func UpdateNode_bodyViaDeleted(iDeleted int, iNode_bodyMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Node_body)).Where("deleted = ?", iDeleted).Update(iNode_bodyMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateNode_bodyViaEntityId via map[string]interface{}{}
func UpdateNode_bodyViaEntityId(iEntityId int, iNode_bodyMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Node_body)).Where("entity_id = ?", iEntityId).Update(iNode_bodyMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateNode_bodyViaRevisionId via map[string]interface{}{}
func UpdateNode_bodyViaRevisionId(iRevisionId int, iNode_bodyMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Node_body)).Where("revision_id = ?", iRevisionId).Update(iNode_bodyMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateNode_bodyViaLangcode via map[string]interface{}{}
func UpdateNode_bodyViaLangcode(iLangcode string, iNode_bodyMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Node_body)).Where("langcode = ?", iLangcode).Update(iNode_bodyMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateNode_bodyViaDelta via map[string]interface{}{}
func UpdateNode_bodyViaDelta(iDelta int, iNode_bodyMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Node_body)).Where("delta = ?", iDelta).Update(iNode_bodyMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateNode_bodyViaBodyValue via map[string]interface{}{}
func UpdateNode_bodyViaBodyValue(iBodyValue string, iNode_bodyMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Node_body)).Where("body_value = ?", iBodyValue).Update(iNode_bodyMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateNode_bodyViaBodySummary via map[string]interface{}{}
func UpdateNode_bodyViaBodySummary(iBodySummary string, iNode_bodyMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Node_body)).Where("body_summary = ?", iBodySummary).Update(iNode_bodyMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateNode_bodyViaBodyFormat via map[string]interface{}{}
func UpdateNode_bodyViaBodyFormat(iBodyFormat string, iNode_bodyMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Node_body)).Where("body_format = ?", iBodyFormat).Update(iNode_bodyMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// DeleteNode_body Delete Node_body
func DeleteNode_body(iBundle string) (int64, error) {
	row, err := Engine.Id(iBundle).Delete(new(Node_body))
	return row, err
}

// DeleteNode_bodyViaBundle Delete Node_body via Bundle
func DeleteNode_bodyViaBundle(iBundle string) (err error) {
	var has bool
	var _Node_body = &Node_body{Bundle: iBundle}
	if has, err = Engine.Get(_Node_body); (has == true) && (err == nil) {
		if row, err := Engine.Where("bundle = ?", iBundle).Delete(new(Node_body)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteNode_bodyViaDeleted Delete Node_body via Deleted
func DeleteNode_bodyViaDeleted(iDeleted int) (err error) {
	var has bool
	var _Node_body = &Node_body{Deleted: iDeleted}
	if has, err = Engine.Get(_Node_body); (has == true) && (err == nil) {
		if row, err := Engine.Where("deleted = ?", iDeleted).Delete(new(Node_body)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteNode_bodyViaEntityId Delete Node_body via EntityId
func DeleteNode_bodyViaEntityId(iEntityId int) (err error) {
	var has bool
	var _Node_body = &Node_body{EntityId: iEntityId}
	if has, err = Engine.Get(_Node_body); (has == true) && (err == nil) {
		if row, err := Engine.Where("entity_id = ?", iEntityId).Delete(new(Node_body)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteNode_bodyViaRevisionId Delete Node_body via RevisionId
func DeleteNode_bodyViaRevisionId(iRevisionId int) (err error) {
	var has bool
	var _Node_body = &Node_body{RevisionId: iRevisionId}
	if has, err = Engine.Get(_Node_body); (has == true) && (err == nil) {
		if row, err := Engine.Where("revision_id = ?", iRevisionId).Delete(new(Node_body)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteNode_bodyViaLangcode Delete Node_body via Langcode
func DeleteNode_bodyViaLangcode(iLangcode string) (err error) {
	var has bool
	var _Node_body = &Node_body{Langcode: iLangcode}
	if has, err = Engine.Get(_Node_body); (has == true) && (err == nil) {
		if row, err := Engine.Where("langcode = ?", iLangcode).Delete(new(Node_body)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteNode_bodyViaDelta Delete Node_body via Delta
func DeleteNode_bodyViaDelta(iDelta int) (err error) {
	var has bool
	var _Node_body = &Node_body{Delta: iDelta}
	if has, err = Engine.Get(_Node_body); (has == true) && (err == nil) {
		if row, err := Engine.Where("delta = ?", iDelta).Delete(new(Node_body)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteNode_bodyViaBodyValue Delete Node_body via BodyValue
func DeleteNode_bodyViaBodyValue(iBodyValue string) (err error) {
	var has bool
	var _Node_body = &Node_body{BodyValue: iBodyValue}
	if has, err = Engine.Get(_Node_body); (has == true) && (err == nil) {
		if row, err := Engine.Where("body_value = ?", iBodyValue).Delete(new(Node_body)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteNode_bodyViaBodySummary Delete Node_body via BodySummary
func DeleteNode_bodyViaBodySummary(iBodySummary string) (err error) {
	var has bool
	var _Node_body = &Node_body{BodySummary: iBodySummary}
	if has, err = Engine.Get(_Node_body); (has == true) && (err == nil) {
		if row, err := Engine.Where("body_summary = ?", iBodySummary).Delete(new(Node_body)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteNode_bodyViaBodyFormat Delete Node_body via BodyFormat
func DeleteNode_bodyViaBodyFormat(iBodyFormat string) (err error) {
	var has bool
	var _Node_body = &Node_body{BodyFormat: iBodyFormat}
	if has, err = Engine.Get(_Node_body); (has == true) && (err == nil) {
		if row, err := Engine.Where("body_format = ?", iBodyFormat).Delete(new(Node_body)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}
