package model

type Comment_commentBody struct {
	Bundle            string `xorm:"not null default '' index VARCHAR(128)"`
	Deleted           int    `xorm:"not null pk default 0 TINYINT(4)"`
	EntityId          int    `xorm:"not null pk INT(10)"`
	RevisionId        int    `xorm:"not null index INT(10)"`
	Langcode          string `xorm:"not null pk default '' VARCHAR(32)"`
	Delta             int    `xorm:"not null pk INT(10)"`
	CommentBodyValue  string `xorm:"not null LONGTEXT"`
	CommentBodyFormat string `xorm:"index VARCHAR(255)"`
}

// GetComment_commentBodiesCount Comment_commentBodys' Count
func GetComment_commentBodiesCount(offset int, limit int) (int64, error) {
	total, err := Engine.Limit(limit, offset).Count(&Comment_commentBody{})
	return total, err
}

// GetComment_commentBodyCountViaBundle Get Comment_commentBody via Bundle
func GetComment_commentBodyCountViaBundle(iBundle string) int64 {
	n, _ := Engine.Where("bundle = ?", iBundle).Count(&Comment_commentBody{Bundle: iBundle})
	return n
}

// GetComment_commentBodyCountViaDeleted Get Comment_commentBody via Deleted
func GetComment_commentBodyCountViaDeleted(iDeleted int) int64 {
	n, _ := Engine.Where("deleted = ?", iDeleted).Count(&Comment_commentBody{Deleted: iDeleted})
	return n
}

// GetComment_commentBodyCountViaEntityId Get Comment_commentBody via EntityId
func GetComment_commentBodyCountViaEntityId(iEntityId int) int64 {
	n, _ := Engine.Where("entity_id = ?", iEntityId).Count(&Comment_commentBody{EntityId: iEntityId})
	return n
}

// GetComment_commentBodyCountViaRevisionId Get Comment_commentBody via RevisionId
func GetComment_commentBodyCountViaRevisionId(iRevisionId int) int64 {
	n, _ := Engine.Where("revision_id = ?", iRevisionId).Count(&Comment_commentBody{RevisionId: iRevisionId})
	return n
}

// GetComment_commentBodyCountViaLangcode Get Comment_commentBody via Langcode
func GetComment_commentBodyCountViaLangcode(iLangcode string) int64 {
	n, _ := Engine.Where("langcode = ?", iLangcode).Count(&Comment_commentBody{Langcode: iLangcode})
	return n
}

// GetComment_commentBodyCountViaDelta Get Comment_commentBody via Delta
func GetComment_commentBodyCountViaDelta(iDelta int) int64 {
	n, _ := Engine.Where("delta = ?", iDelta).Count(&Comment_commentBody{Delta: iDelta})
	return n
}

// GetComment_commentBodyCountViaCommentBodyValue Get Comment_commentBody via CommentBodyValue
func GetComment_commentBodyCountViaCommentBodyValue(iCommentBodyValue string) int64 {
	n, _ := Engine.Where("comment_body_value = ?", iCommentBodyValue).Count(&Comment_commentBody{CommentBodyValue: iCommentBodyValue})
	return n
}

// GetComment_commentBodyCountViaCommentBodyFormat Get Comment_commentBody via CommentBodyFormat
func GetComment_commentBodyCountViaCommentBodyFormat(iCommentBodyFormat string) int64 {
	n, _ := Engine.Where("comment_body_format = ?", iCommentBodyFormat).Count(&Comment_commentBody{CommentBodyFormat: iCommentBodyFormat})
	return n
}

// GetComment_commentBodiesByBundleAndDeletedAndEntityId Get Comment_commentBodies via BundleAndDeletedAndEntityId
func GetComment_commentBodiesByBundleAndDeletedAndEntityId(offset int, limit int, Bundle_ string, Deleted_ int, EntityId_ int) (*[]*Comment_commentBody, error) {
	var _Comment_commentBody = new([]*Comment_commentBody)
	err := Engine.Table("comment__comment_body").Where("bundle = ? and deleted = ? and entity_id = ?", Bundle_, Deleted_, EntityId_).Limit(limit, offset).Find(_Comment_commentBody)
	return _Comment_commentBody, err
}

// GetComment_commentBodiesByBundleAndDeletedAndRevisionId Get Comment_commentBodies via BundleAndDeletedAndRevisionId
func GetComment_commentBodiesByBundleAndDeletedAndRevisionId(offset int, limit int, Bundle_ string, Deleted_ int, RevisionId_ int) (*[]*Comment_commentBody, error) {
	var _Comment_commentBody = new([]*Comment_commentBody)
	err := Engine.Table("comment__comment_body").Where("bundle = ? and deleted = ? and revision_id = ?", Bundle_, Deleted_, RevisionId_).Limit(limit, offset).Find(_Comment_commentBody)
	return _Comment_commentBody, err
}

// GetComment_commentBodiesByBundleAndDeletedAndLangcode Get Comment_commentBodies via BundleAndDeletedAndLangcode
func GetComment_commentBodiesByBundleAndDeletedAndLangcode(offset int, limit int, Bundle_ string, Deleted_ int, Langcode_ string) (*[]*Comment_commentBody, error) {
	var _Comment_commentBody = new([]*Comment_commentBody)
	err := Engine.Table("comment__comment_body").Where("bundle = ? and deleted = ? and langcode = ?", Bundle_, Deleted_, Langcode_).Limit(limit, offset).Find(_Comment_commentBody)
	return _Comment_commentBody, err
}

// GetComment_commentBodiesByBundleAndDeletedAndDelta Get Comment_commentBodies via BundleAndDeletedAndDelta
func GetComment_commentBodiesByBundleAndDeletedAndDelta(offset int, limit int, Bundle_ string, Deleted_ int, Delta_ int) (*[]*Comment_commentBody, error) {
	var _Comment_commentBody = new([]*Comment_commentBody)
	err := Engine.Table("comment__comment_body").Where("bundle = ? and deleted = ? and delta = ?", Bundle_, Deleted_, Delta_).Limit(limit, offset).Find(_Comment_commentBody)
	return _Comment_commentBody, err
}

// GetComment_commentBodiesByBundleAndDeletedAndCommentBodyValue Get Comment_commentBodies via BundleAndDeletedAndCommentBodyValue
func GetComment_commentBodiesByBundleAndDeletedAndCommentBodyValue(offset int, limit int, Bundle_ string, Deleted_ int, CommentBodyValue_ string) (*[]*Comment_commentBody, error) {
	var _Comment_commentBody = new([]*Comment_commentBody)
	err := Engine.Table("comment__comment_body").Where("bundle = ? and deleted = ? and comment_body_value = ?", Bundle_, Deleted_, CommentBodyValue_).Limit(limit, offset).Find(_Comment_commentBody)
	return _Comment_commentBody, err
}

// GetComment_commentBodiesByBundleAndDeletedAndCommentBodyFormat Get Comment_commentBodies via BundleAndDeletedAndCommentBodyFormat
func GetComment_commentBodiesByBundleAndDeletedAndCommentBodyFormat(offset int, limit int, Bundle_ string, Deleted_ int, CommentBodyFormat_ string) (*[]*Comment_commentBody, error) {
	var _Comment_commentBody = new([]*Comment_commentBody)
	err := Engine.Table("comment__comment_body").Where("bundle = ? and deleted = ? and comment_body_format = ?", Bundle_, Deleted_, CommentBodyFormat_).Limit(limit, offset).Find(_Comment_commentBody)
	return _Comment_commentBody, err
}

// GetComment_commentBodiesByBundleAndEntityIdAndRevisionId Get Comment_commentBodies via BundleAndEntityIdAndRevisionId
func GetComment_commentBodiesByBundleAndEntityIdAndRevisionId(offset int, limit int, Bundle_ string, EntityId_ int, RevisionId_ int) (*[]*Comment_commentBody, error) {
	var _Comment_commentBody = new([]*Comment_commentBody)
	err := Engine.Table("comment__comment_body").Where("bundle = ? and entity_id = ? and revision_id = ?", Bundle_, EntityId_, RevisionId_).Limit(limit, offset).Find(_Comment_commentBody)
	return _Comment_commentBody, err
}

// GetComment_commentBodiesByBundleAndEntityIdAndLangcode Get Comment_commentBodies via BundleAndEntityIdAndLangcode
func GetComment_commentBodiesByBundleAndEntityIdAndLangcode(offset int, limit int, Bundle_ string, EntityId_ int, Langcode_ string) (*[]*Comment_commentBody, error) {
	var _Comment_commentBody = new([]*Comment_commentBody)
	err := Engine.Table("comment__comment_body").Where("bundle = ? and entity_id = ? and langcode = ?", Bundle_, EntityId_, Langcode_).Limit(limit, offset).Find(_Comment_commentBody)
	return _Comment_commentBody, err
}

// GetComment_commentBodiesByBundleAndEntityIdAndDelta Get Comment_commentBodies via BundleAndEntityIdAndDelta
func GetComment_commentBodiesByBundleAndEntityIdAndDelta(offset int, limit int, Bundle_ string, EntityId_ int, Delta_ int) (*[]*Comment_commentBody, error) {
	var _Comment_commentBody = new([]*Comment_commentBody)
	err := Engine.Table("comment__comment_body").Where("bundle = ? and entity_id = ? and delta = ?", Bundle_, EntityId_, Delta_).Limit(limit, offset).Find(_Comment_commentBody)
	return _Comment_commentBody, err
}

// GetComment_commentBodiesByBundleAndEntityIdAndCommentBodyValue Get Comment_commentBodies via BundleAndEntityIdAndCommentBodyValue
func GetComment_commentBodiesByBundleAndEntityIdAndCommentBodyValue(offset int, limit int, Bundle_ string, EntityId_ int, CommentBodyValue_ string) (*[]*Comment_commentBody, error) {
	var _Comment_commentBody = new([]*Comment_commentBody)
	err := Engine.Table("comment__comment_body").Where("bundle = ? and entity_id = ? and comment_body_value = ?", Bundle_, EntityId_, CommentBodyValue_).Limit(limit, offset).Find(_Comment_commentBody)
	return _Comment_commentBody, err
}

// GetComment_commentBodiesByBundleAndEntityIdAndCommentBodyFormat Get Comment_commentBodies via BundleAndEntityIdAndCommentBodyFormat
func GetComment_commentBodiesByBundleAndEntityIdAndCommentBodyFormat(offset int, limit int, Bundle_ string, EntityId_ int, CommentBodyFormat_ string) (*[]*Comment_commentBody, error) {
	var _Comment_commentBody = new([]*Comment_commentBody)
	err := Engine.Table("comment__comment_body").Where("bundle = ? and entity_id = ? and comment_body_format = ?", Bundle_, EntityId_, CommentBodyFormat_).Limit(limit, offset).Find(_Comment_commentBody)
	return _Comment_commentBody, err
}

// GetComment_commentBodiesByBundleAndRevisionIdAndLangcode Get Comment_commentBodies via BundleAndRevisionIdAndLangcode
func GetComment_commentBodiesByBundleAndRevisionIdAndLangcode(offset int, limit int, Bundle_ string, RevisionId_ int, Langcode_ string) (*[]*Comment_commentBody, error) {
	var _Comment_commentBody = new([]*Comment_commentBody)
	err := Engine.Table("comment__comment_body").Where("bundle = ? and revision_id = ? and langcode = ?", Bundle_, RevisionId_, Langcode_).Limit(limit, offset).Find(_Comment_commentBody)
	return _Comment_commentBody, err
}

// GetComment_commentBodiesByBundleAndRevisionIdAndDelta Get Comment_commentBodies via BundleAndRevisionIdAndDelta
func GetComment_commentBodiesByBundleAndRevisionIdAndDelta(offset int, limit int, Bundle_ string, RevisionId_ int, Delta_ int) (*[]*Comment_commentBody, error) {
	var _Comment_commentBody = new([]*Comment_commentBody)
	err := Engine.Table("comment__comment_body").Where("bundle = ? and revision_id = ? and delta = ?", Bundle_, RevisionId_, Delta_).Limit(limit, offset).Find(_Comment_commentBody)
	return _Comment_commentBody, err
}

// GetComment_commentBodiesByBundleAndRevisionIdAndCommentBodyValue Get Comment_commentBodies via BundleAndRevisionIdAndCommentBodyValue
func GetComment_commentBodiesByBundleAndRevisionIdAndCommentBodyValue(offset int, limit int, Bundle_ string, RevisionId_ int, CommentBodyValue_ string) (*[]*Comment_commentBody, error) {
	var _Comment_commentBody = new([]*Comment_commentBody)
	err := Engine.Table("comment__comment_body").Where("bundle = ? and revision_id = ? and comment_body_value = ?", Bundle_, RevisionId_, CommentBodyValue_).Limit(limit, offset).Find(_Comment_commentBody)
	return _Comment_commentBody, err
}

// GetComment_commentBodiesByBundleAndRevisionIdAndCommentBodyFormat Get Comment_commentBodies via BundleAndRevisionIdAndCommentBodyFormat
func GetComment_commentBodiesByBundleAndRevisionIdAndCommentBodyFormat(offset int, limit int, Bundle_ string, RevisionId_ int, CommentBodyFormat_ string) (*[]*Comment_commentBody, error) {
	var _Comment_commentBody = new([]*Comment_commentBody)
	err := Engine.Table("comment__comment_body").Where("bundle = ? and revision_id = ? and comment_body_format = ?", Bundle_, RevisionId_, CommentBodyFormat_).Limit(limit, offset).Find(_Comment_commentBody)
	return _Comment_commentBody, err
}

// GetComment_commentBodiesByBundleAndLangcodeAndDelta Get Comment_commentBodies via BundleAndLangcodeAndDelta
func GetComment_commentBodiesByBundleAndLangcodeAndDelta(offset int, limit int, Bundle_ string, Langcode_ string, Delta_ int) (*[]*Comment_commentBody, error) {
	var _Comment_commentBody = new([]*Comment_commentBody)
	err := Engine.Table("comment__comment_body").Where("bundle = ? and langcode = ? and delta = ?", Bundle_, Langcode_, Delta_).Limit(limit, offset).Find(_Comment_commentBody)
	return _Comment_commentBody, err
}

// GetComment_commentBodiesByBundleAndLangcodeAndCommentBodyValue Get Comment_commentBodies via BundleAndLangcodeAndCommentBodyValue
func GetComment_commentBodiesByBundleAndLangcodeAndCommentBodyValue(offset int, limit int, Bundle_ string, Langcode_ string, CommentBodyValue_ string) (*[]*Comment_commentBody, error) {
	var _Comment_commentBody = new([]*Comment_commentBody)
	err := Engine.Table("comment__comment_body").Where("bundle = ? and langcode = ? and comment_body_value = ?", Bundle_, Langcode_, CommentBodyValue_).Limit(limit, offset).Find(_Comment_commentBody)
	return _Comment_commentBody, err
}

// GetComment_commentBodiesByBundleAndLangcodeAndCommentBodyFormat Get Comment_commentBodies via BundleAndLangcodeAndCommentBodyFormat
func GetComment_commentBodiesByBundleAndLangcodeAndCommentBodyFormat(offset int, limit int, Bundle_ string, Langcode_ string, CommentBodyFormat_ string) (*[]*Comment_commentBody, error) {
	var _Comment_commentBody = new([]*Comment_commentBody)
	err := Engine.Table("comment__comment_body").Where("bundle = ? and langcode = ? and comment_body_format = ?", Bundle_, Langcode_, CommentBodyFormat_).Limit(limit, offset).Find(_Comment_commentBody)
	return _Comment_commentBody, err
}

// GetComment_commentBodiesByBundleAndDeltaAndCommentBodyValue Get Comment_commentBodies via BundleAndDeltaAndCommentBodyValue
func GetComment_commentBodiesByBundleAndDeltaAndCommentBodyValue(offset int, limit int, Bundle_ string, Delta_ int, CommentBodyValue_ string) (*[]*Comment_commentBody, error) {
	var _Comment_commentBody = new([]*Comment_commentBody)
	err := Engine.Table("comment__comment_body").Where("bundle = ? and delta = ? and comment_body_value = ?", Bundle_, Delta_, CommentBodyValue_).Limit(limit, offset).Find(_Comment_commentBody)
	return _Comment_commentBody, err
}

// GetComment_commentBodiesByBundleAndDeltaAndCommentBodyFormat Get Comment_commentBodies via BundleAndDeltaAndCommentBodyFormat
func GetComment_commentBodiesByBundleAndDeltaAndCommentBodyFormat(offset int, limit int, Bundle_ string, Delta_ int, CommentBodyFormat_ string) (*[]*Comment_commentBody, error) {
	var _Comment_commentBody = new([]*Comment_commentBody)
	err := Engine.Table("comment__comment_body").Where("bundle = ? and delta = ? and comment_body_format = ?", Bundle_, Delta_, CommentBodyFormat_).Limit(limit, offset).Find(_Comment_commentBody)
	return _Comment_commentBody, err
}

// GetComment_commentBodiesByBundleAndCommentBodyValueAndCommentBodyFormat Get Comment_commentBodies via BundleAndCommentBodyValueAndCommentBodyFormat
func GetComment_commentBodiesByBundleAndCommentBodyValueAndCommentBodyFormat(offset int, limit int, Bundle_ string, CommentBodyValue_ string, CommentBodyFormat_ string) (*[]*Comment_commentBody, error) {
	var _Comment_commentBody = new([]*Comment_commentBody)
	err := Engine.Table("comment__comment_body").Where("bundle = ? and comment_body_value = ? and comment_body_format = ?", Bundle_, CommentBodyValue_, CommentBodyFormat_).Limit(limit, offset).Find(_Comment_commentBody)
	return _Comment_commentBody, err
}

// GetComment_commentBodiesByDeletedAndEntityIdAndRevisionId Get Comment_commentBodies via DeletedAndEntityIdAndRevisionId
func GetComment_commentBodiesByDeletedAndEntityIdAndRevisionId(offset int, limit int, Deleted_ int, EntityId_ int, RevisionId_ int) (*[]*Comment_commentBody, error) {
	var _Comment_commentBody = new([]*Comment_commentBody)
	err := Engine.Table("comment__comment_body").Where("deleted = ? and entity_id = ? and revision_id = ?", Deleted_, EntityId_, RevisionId_).Limit(limit, offset).Find(_Comment_commentBody)
	return _Comment_commentBody, err
}

// GetComment_commentBodiesByDeletedAndEntityIdAndLangcode Get Comment_commentBodies via DeletedAndEntityIdAndLangcode
func GetComment_commentBodiesByDeletedAndEntityIdAndLangcode(offset int, limit int, Deleted_ int, EntityId_ int, Langcode_ string) (*[]*Comment_commentBody, error) {
	var _Comment_commentBody = new([]*Comment_commentBody)
	err := Engine.Table("comment__comment_body").Where("deleted = ? and entity_id = ? and langcode = ?", Deleted_, EntityId_, Langcode_).Limit(limit, offset).Find(_Comment_commentBody)
	return _Comment_commentBody, err
}

// GetComment_commentBodiesByDeletedAndEntityIdAndDelta Get Comment_commentBodies via DeletedAndEntityIdAndDelta
func GetComment_commentBodiesByDeletedAndEntityIdAndDelta(offset int, limit int, Deleted_ int, EntityId_ int, Delta_ int) (*[]*Comment_commentBody, error) {
	var _Comment_commentBody = new([]*Comment_commentBody)
	err := Engine.Table("comment__comment_body").Where("deleted = ? and entity_id = ? and delta = ?", Deleted_, EntityId_, Delta_).Limit(limit, offset).Find(_Comment_commentBody)
	return _Comment_commentBody, err
}

// GetComment_commentBodiesByDeletedAndEntityIdAndCommentBodyValue Get Comment_commentBodies via DeletedAndEntityIdAndCommentBodyValue
func GetComment_commentBodiesByDeletedAndEntityIdAndCommentBodyValue(offset int, limit int, Deleted_ int, EntityId_ int, CommentBodyValue_ string) (*[]*Comment_commentBody, error) {
	var _Comment_commentBody = new([]*Comment_commentBody)
	err := Engine.Table("comment__comment_body").Where("deleted = ? and entity_id = ? and comment_body_value = ?", Deleted_, EntityId_, CommentBodyValue_).Limit(limit, offset).Find(_Comment_commentBody)
	return _Comment_commentBody, err
}

// GetComment_commentBodiesByDeletedAndEntityIdAndCommentBodyFormat Get Comment_commentBodies via DeletedAndEntityIdAndCommentBodyFormat
func GetComment_commentBodiesByDeletedAndEntityIdAndCommentBodyFormat(offset int, limit int, Deleted_ int, EntityId_ int, CommentBodyFormat_ string) (*[]*Comment_commentBody, error) {
	var _Comment_commentBody = new([]*Comment_commentBody)
	err := Engine.Table("comment__comment_body").Where("deleted = ? and entity_id = ? and comment_body_format = ?", Deleted_, EntityId_, CommentBodyFormat_).Limit(limit, offset).Find(_Comment_commentBody)
	return _Comment_commentBody, err
}

// GetComment_commentBodiesByDeletedAndRevisionIdAndLangcode Get Comment_commentBodies via DeletedAndRevisionIdAndLangcode
func GetComment_commentBodiesByDeletedAndRevisionIdAndLangcode(offset int, limit int, Deleted_ int, RevisionId_ int, Langcode_ string) (*[]*Comment_commentBody, error) {
	var _Comment_commentBody = new([]*Comment_commentBody)
	err := Engine.Table("comment__comment_body").Where("deleted = ? and revision_id = ? and langcode = ?", Deleted_, RevisionId_, Langcode_).Limit(limit, offset).Find(_Comment_commentBody)
	return _Comment_commentBody, err
}

// GetComment_commentBodiesByDeletedAndRevisionIdAndDelta Get Comment_commentBodies via DeletedAndRevisionIdAndDelta
func GetComment_commentBodiesByDeletedAndRevisionIdAndDelta(offset int, limit int, Deleted_ int, RevisionId_ int, Delta_ int) (*[]*Comment_commentBody, error) {
	var _Comment_commentBody = new([]*Comment_commentBody)
	err := Engine.Table("comment__comment_body").Where("deleted = ? and revision_id = ? and delta = ?", Deleted_, RevisionId_, Delta_).Limit(limit, offset).Find(_Comment_commentBody)
	return _Comment_commentBody, err
}

// GetComment_commentBodiesByDeletedAndRevisionIdAndCommentBodyValue Get Comment_commentBodies via DeletedAndRevisionIdAndCommentBodyValue
func GetComment_commentBodiesByDeletedAndRevisionIdAndCommentBodyValue(offset int, limit int, Deleted_ int, RevisionId_ int, CommentBodyValue_ string) (*[]*Comment_commentBody, error) {
	var _Comment_commentBody = new([]*Comment_commentBody)
	err := Engine.Table("comment__comment_body").Where("deleted = ? and revision_id = ? and comment_body_value = ?", Deleted_, RevisionId_, CommentBodyValue_).Limit(limit, offset).Find(_Comment_commentBody)
	return _Comment_commentBody, err
}

// GetComment_commentBodiesByDeletedAndRevisionIdAndCommentBodyFormat Get Comment_commentBodies via DeletedAndRevisionIdAndCommentBodyFormat
func GetComment_commentBodiesByDeletedAndRevisionIdAndCommentBodyFormat(offset int, limit int, Deleted_ int, RevisionId_ int, CommentBodyFormat_ string) (*[]*Comment_commentBody, error) {
	var _Comment_commentBody = new([]*Comment_commentBody)
	err := Engine.Table("comment__comment_body").Where("deleted = ? and revision_id = ? and comment_body_format = ?", Deleted_, RevisionId_, CommentBodyFormat_).Limit(limit, offset).Find(_Comment_commentBody)
	return _Comment_commentBody, err
}

// GetComment_commentBodiesByDeletedAndLangcodeAndDelta Get Comment_commentBodies via DeletedAndLangcodeAndDelta
func GetComment_commentBodiesByDeletedAndLangcodeAndDelta(offset int, limit int, Deleted_ int, Langcode_ string, Delta_ int) (*[]*Comment_commentBody, error) {
	var _Comment_commentBody = new([]*Comment_commentBody)
	err := Engine.Table("comment__comment_body").Where("deleted = ? and langcode = ? and delta = ?", Deleted_, Langcode_, Delta_).Limit(limit, offset).Find(_Comment_commentBody)
	return _Comment_commentBody, err
}

// GetComment_commentBodiesByDeletedAndLangcodeAndCommentBodyValue Get Comment_commentBodies via DeletedAndLangcodeAndCommentBodyValue
func GetComment_commentBodiesByDeletedAndLangcodeAndCommentBodyValue(offset int, limit int, Deleted_ int, Langcode_ string, CommentBodyValue_ string) (*[]*Comment_commentBody, error) {
	var _Comment_commentBody = new([]*Comment_commentBody)
	err := Engine.Table("comment__comment_body").Where("deleted = ? and langcode = ? and comment_body_value = ?", Deleted_, Langcode_, CommentBodyValue_).Limit(limit, offset).Find(_Comment_commentBody)
	return _Comment_commentBody, err
}

// GetComment_commentBodiesByDeletedAndLangcodeAndCommentBodyFormat Get Comment_commentBodies via DeletedAndLangcodeAndCommentBodyFormat
func GetComment_commentBodiesByDeletedAndLangcodeAndCommentBodyFormat(offset int, limit int, Deleted_ int, Langcode_ string, CommentBodyFormat_ string) (*[]*Comment_commentBody, error) {
	var _Comment_commentBody = new([]*Comment_commentBody)
	err := Engine.Table("comment__comment_body").Where("deleted = ? and langcode = ? and comment_body_format = ?", Deleted_, Langcode_, CommentBodyFormat_).Limit(limit, offset).Find(_Comment_commentBody)
	return _Comment_commentBody, err
}

// GetComment_commentBodiesByDeletedAndDeltaAndCommentBodyValue Get Comment_commentBodies via DeletedAndDeltaAndCommentBodyValue
func GetComment_commentBodiesByDeletedAndDeltaAndCommentBodyValue(offset int, limit int, Deleted_ int, Delta_ int, CommentBodyValue_ string) (*[]*Comment_commentBody, error) {
	var _Comment_commentBody = new([]*Comment_commentBody)
	err := Engine.Table("comment__comment_body").Where("deleted = ? and delta = ? and comment_body_value = ?", Deleted_, Delta_, CommentBodyValue_).Limit(limit, offset).Find(_Comment_commentBody)
	return _Comment_commentBody, err
}

// GetComment_commentBodiesByDeletedAndDeltaAndCommentBodyFormat Get Comment_commentBodies via DeletedAndDeltaAndCommentBodyFormat
func GetComment_commentBodiesByDeletedAndDeltaAndCommentBodyFormat(offset int, limit int, Deleted_ int, Delta_ int, CommentBodyFormat_ string) (*[]*Comment_commentBody, error) {
	var _Comment_commentBody = new([]*Comment_commentBody)
	err := Engine.Table("comment__comment_body").Where("deleted = ? and delta = ? and comment_body_format = ?", Deleted_, Delta_, CommentBodyFormat_).Limit(limit, offset).Find(_Comment_commentBody)
	return _Comment_commentBody, err
}

// GetComment_commentBodiesByDeletedAndCommentBodyValueAndCommentBodyFormat Get Comment_commentBodies via DeletedAndCommentBodyValueAndCommentBodyFormat
func GetComment_commentBodiesByDeletedAndCommentBodyValueAndCommentBodyFormat(offset int, limit int, Deleted_ int, CommentBodyValue_ string, CommentBodyFormat_ string) (*[]*Comment_commentBody, error) {
	var _Comment_commentBody = new([]*Comment_commentBody)
	err := Engine.Table("comment__comment_body").Where("deleted = ? and comment_body_value = ? and comment_body_format = ?", Deleted_, CommentBodyValue_, CommentBodyFormat_).Limit(limit, offset).Find(_Comment_commentBody)
	return _Comment_commentBody, err
}

// GetComment_commentBodiesByEntityIdAndRevisionIdAndLangcode Get Comment_commentBodies via EntityIdAndRevisionIdAndLangcode
func GetComment_commentBodiesByEntityIdAndRevisionIdAndLangcode(offset int, limit int, EntityId_ int, RevisionId_ int, Langcode_ string) (*[]*Comment_commentBody, error) {
	var _Comment_commentBody = new([]*Comment_commentBody)
	err := Engine.Table("comment__comment_body").Where("entity_id = ? and revision_id = ? and langcode = ?", EntityId_, RevisionId_, Langcode_).Limit(limit, offset).Find(_Comment_commentBody)
	return _Comment_commentBody, err
}

// GetComment_commentBodiesByEntityIdAndRevisionIdAndDelta Get Comment_commentBodies via EntityIdAndRevisionIdAndDelta
func GetComment_commentBodiesByEntityIdAndRevisionIdAndDelta(offset int, limit int, EntityId_ int, RevisionId_ int, Delta_ int) (*[]*Comment_commentBody, error) {
	var _Comment_commentBody = new([]*Comment_commentBody)
	err := Engine.Table("comment__comment_body").Where("entity_id = ? and revision_id = ? and delta = ?", EntityId_, RevisionId_, Delta_).Limit(limit, offset).Find(_Comment_commentBody)
	return _Comment_commentBody, err
}

// GetComment_commentBodiesByEntityIdAndRevisionIdAndCommentBodyValue Get Comment_commentBodies via EntityIdAndRevisionIdAndCommentBodyValue
func GetComment_commentBodiesByEntityIdAndRevisionIdAndCommentBodyValue(offset int, limit int, EntityId_ int, RevisionId_ int, CommentBodyValue_ string) (*[]*Comment_commentBody, error) {
	var _Comment_commentBody = new([]*Comment_commentBody)
	err := Engine.Table("comment__comment_body").Where("entity_id = ? and revision_id = ? and comment_body_value = ?", EntityId_, RevisionId_, CommentBodyValue_).Limit(limit, offset).Find(_Comment_commentBody)
	return _Comment_commentBody, err
}

// GetComment_commentBodiesByEntityIdAndRevisionIdAndCommentBodyFormat Get Comment_commentBodies via EntityIdAndRevisionIdAndCommentBodyFormat
func GetComment_commentBodiesByEntityIdAndRevisionIdAndCommentBodyFormat(offset int, limit int, EntityId_ int, RevisionId_ int, CommentBodyFormat_ string) (*[]*Comment_commentBody, error) {
	var _Comment_commentBody = new([]*Comment_commentBody)
	err := Engine.Table("comment__comment_body").Where("entity_id = ? and revision_id = ? and comment_body_format = ?", EntityId_, RevisionId_, CommentBodyFormat_).Limit(limit, offset).Find(_Comment_commentBody)
	return _Comment_commentBody, err
}

// GetComment_commentBodiesByEntityIdAndLangcodeAndDelta Get Comment_commentBodies via EntityIdAndLangcodeAndDelta
func GetComment_commentBodiesByEntityIdAndLangcodeAndDelta(offset int, limit int, EntityId_ int, Langcode_ string, Delta_ int) (*[]*Comment_commentBody, error) {
	var _Comment_commentBody = new([]*Comment_commentBody)
	err := Engine.Table("comment__comment_body").Where("entity_id = ? and langcode = ? and delta = ?", EntityId_, Langcode_, Delta_).Limit(limit, offset).Find(_Comment_commentBody)
	return _Comment_commentBody, err
}

// GetComment_commentBodiesByEntityIdAndLangcodeAndCommentBodyValue Get Comment_commentBodies via EntityIdAndLangcodeAndCommentBodyValue
func GetComment_commentBodiesByEntityIdAndLangcodeAndCommentBodyValue(offset int, limit int, EntityId_ int, Langcode_ string, CommentBodyValue_ string) (*[]*Comment_commentBody, error) {
	var _Comment_commentBody = new([]*Comment_commentBody)
	err := Engine.Table("comment__comment_body").Where("entity_id = ? and langcode = ? and comment_body_value = ?", EntityId_, Langcode_, CommentBodyValue_).Limit(limit, offset).Find(_Comment_commentBody)
	return _Comment_commentBody, err
}

// GetComment_commentBodiesByEntityIdAndLangcodeAndCommentBodyFormat Get Comment_commentBodies via EntityIdAndLangcodeAndCommentBodyFormat
func GetComment_commentBodiesByEntityIdAndLangcodeAndCommentBodyFormat(offset int, limit int, EntityId_ int, Langcode_ string, CommentBodyFormat_ string) (*[]*Comment_commentBody, error) {
	var _Comment_commentBody = new([]*Comment_commentBody)
	err := Engine.Table("comment__comment_body").Where("entity_id = ? and langcode = ? and comment_body_format = ?", EntityId_, Langcode_, CommentBodyFormat_).Limit(limit, offset).Find(_Comment_commentBody)
	return _Comment_commentBody, err
}

// GetComment_commentBodiesByEntityIdAndDeltaAndCommentBodyValue Get Comment_commentBodies via EntityIdAndDeltaAndCommentBodyValue
func GetComment_commentBodiesByEntityIdAndDeltaAndCommentBodyValue(offset int, limit int, EntityId_ int, Delta_ int, CommentBodyValue_ string) (*[]*Comment_commentBody, error) {
	var _Comment_commentBody = new([]*Comment_commentBody)
	err := Engine.Table("comment__comment_body").Where("entity_id = ? and delta = ? and comment_body_value = ?", EntityId_, Delta_, CommentBodyValue_).Limit(limit, offset).Find(_Comment_commentBody)
	return _Comment_commentBody, err
}

// GetComment_commentBodiesByEntityIdAndDeltaAndCommentBodyFormat Get Comment_commentBodies via EntityIdAndDeltaAndCommentBodyFormat
func GetComment_commentBodiesByEntityIdAndDeltaAndCommentBodyFormat(offset int, limit int, EntityId_ int, Delta_ int, CommentBodyFormat_ string) (*[]*Comment_commentBody, error) {
	var _Comment_commentBody = new([]*Comment_commentBody)
	err := Engine.Table("comment__comment_body").Where("entity_id = ? and delta = ? and comment_body_format = ?", EntityId_, Delta_, CommentBodyFormat_).Limit(limit, offset).Find(_Comment_commentBody)
	return _Comment_commentBody, err
}

// GetComment_commentBodiesByEntityIdAndCommentBodyValueAndCommentBodyFormat Get Comment_commentBodies via EntityIdAndCommentBodyValueAndCommentBodyFormat
func GetComment_commentBodiesByEntityIdAndCommentBodyValueAndCommentBodyFormat(offset int, limit int, EntityId_ int, CommentBodyValue_ string, CommentBodyFormat_ string) (*[]*Comment_commentBody, error) {
	var _Comment_commentBody = new([]*Comment_commentBody)
	err := Engine.Table("comment__comment_body").Where("entity_id = ? and comment_body_value = ? and comment_body_format = ?", EntityId_, CommentBodyValue_, CommentBodyFormat_).Limit(limit, offset).Find(_Comment_commentBody)
	return _Comment_commentBody, err
}

// GetComment_commentBodiesByRevisionIdAndLangcodeAndDelta Get Comment_commentBodies via RevisionIdAndLangcodeAndDelta
func GetComment_commentBodiesByRevisionIdAndLangcodeAndDelta(offset int, limit int, RevisionId_ int, Langcode_ string, Delta_ int) (*[]*Comment_commentBody, error) {
	var _Comment_commentBody = new([]*Comment_commentBody)
	err := Engine.Table("comment__comment_body").Where("revision_id = ? and langcode = ? and delta = ?", RevisionId_, Langcode_, Delta_).Limit(limit, offset).Find(_Comment_commentBody)
	return _Comment_commentBody, err
}

// GetComment_commentBodiesByRevisionIdAndLangcodeAndCommentBodyValue Get Comment_commentBodies via RevisionIdAndLangcodeAndCommentBodyValue
func GetComment_commentBodiesByRevisionIdAndLangcodeAndCommentBodyValue(offset int, limit int, RevisionId_ int, Langcode_ string, CommentBodyValue_ string) (*[]*Comment_commentBody, error) {
	var _Comment_commentBody = new([]*Comment_commentBody)
	err := Engine.Table("comment__comment_body").Where("revision_id = ? and langcode = ? and comment_body_value = ?", RevisionId_, Langcode_, CommentBodyValue_).Limit(limit, offset).Find(_Comment_commentBody)
	return _Comment_commentBody, err
}

// GetComment_commentBodiesByRevisionIdAndLangcodeAndCommentBodyFormat Get Comment_commentBodies via RevisionIdAndLangcodeAndCommentBodyFormat
func GetComment_commentBodiesByRevisionIdAndLangcodeAndCommentBodyFormat(offset int, limit int, RevisionId_ int, Langcode_ string, CommentBodyFormat_ string) (*[]*Comment_commentBody, error) {
	var _Comment_commentBody = new([]*Comment_commentBody)
	err := Engine.Table("comment__comment_body").Where("revision_id = ? and langcode = ? and comment_body_format = ?", RevisionId_, Langcode_, CommentBodyFormat_).Limit(limit, offset).Find(_Comment_commentBody)
	return _Comment_commentBody, err
}

// GetComment_commentBodiesByRevisionIdAndDeltaAndCommentBodyValue Get Comment_commentBodies via RevisionIdAndDeltaAndCommentBodyValue
func GetComment_commentBodiesByRevisionIdAndDeltaAndCommentBodyValue(offset int, limit int, RevisionId_ int, Delta_ int, CommentBodyValue_ string) (*[]*Comment_commentBody, error) {
	var _Comment_commentBody = new([]*Comment_commentBody)
	err := Engine.Table("comment__comment_body").Where("revision_id = ? and delta = ? and comment_body_value = ?", RevisionId_, Delta_, CommentBodyValue_).Limit(limit, offset).Find(_Comment_commentBody)
	return _Comment_commentBody, err
}

// GetComment_commentBodiesByRevisionIdAndDeltaAndCommentBodyFormat Get Comment_commentBodies via RevisionIdAndDeltaAndCommentBodyFormat
func GetComment_commentBodiesByRevisionIdAndDeltaAndCommentBodyFormat(offset int, limit int, RevisionId_ int, Delta_ int, CommentBodyFormat_ string) (*[]*Comment_commentBody, error) {
	var _Comment_commentBody = new([]*Comment_commentBody)
	err := Engine.Table("comment__comment_body").Where("revision_id = ? and delta = ? and comment_body_format = ?", RevisionId_, Delta_, CommentBodyFormat_).Limit(limit, offset).Find(_Comment_commentBody)
	return _Comment_commentBody, err
}

// GetComment_commentBodiesByRevisionIdAndCommentBodyValueAndCommentBodyFormat Get Comment_commentBodies via RevisionIdAndCommentBodyValueAndCommentBodyFormat
func GetComment_commentBodiesByRevisionIdAndCommentBodyValueAndCommentBodyFormat(offset int, limit int, RevisionId_ int, CommentBodyValue_ string, CommentBodyFormat_ string) (*[]*Comment_commentBody, error) {
	var _Comment_commentBody = new([]*Comment_commentBody)
	err := Engine.Table("comment__comment_body").Where("revision_id = ? and comment_body_value = ? and comment_body_format = ?", RevisionId_, CommentBodyValue_, CommentBodyFormat_).Limit(limit, offset).Find(_Comment_commentBody)
	return _Comment_commentBody, err
}

// GetComment_commentBodiesByLangcodeAndDeltaAndCommentBodyValue Get Comment_commentBodies via LangcodeAndDeltaAndCommentBodyValue
func GetComment_commentBodiesByLangcodeAndDeltaAndCommentBodyValue(offset int, limit int, Langcode_ string, Delta_ int, CommentBodyValue_ string) (*[]*Comment_commentBody, error) {
	var _Comment_commentBody = new([]*Comment_commentBody)
	err := Engine.Table("comment__comment_body").Where("langcode = ? and delta = ? and comment_body_value = ?", Langcode_, Delta_, CommentBodyValue_).Limit(limit, offset).Find(_Comment_commentBody)
	return _Comment_commentBody, err
}

// GetComment_commentBodiesByLangcodeAndDeltaAndCommentBodyFormat Get Comment_commentBodies via LangcodeAndDeltaAndCommentBodyFormat
func GetComment_commentBodiesByLangcodeAndDeltaAndCommentBodyFormat(offset int, limit int, Langcode_ string, Delta_ int, CommentBodyFormat_ string) (*[]*Comment_commentBody, error) {
	var _Comment_commentBody = new([]*Comment_commentBody)
	err := Engine.Table("comment__comment_body").Where("langcode = ? and delta = ? and comment_body_format = ?", Langcode_, Delta_, CommentBodyFormat_).Limit(limit, offset).Find(_Comment_commentBody)
	return _Comment_commentBody, err
}

// GetComment_commentBodiesByLangcodeAndCommentBodyValueAndCommentBodyFormat Get Comment_commentBodies via LangcodeAndCommentBodyValueAndCommentBodyFormat
func GetComment_commentBodiesByLangcodeAndCommentBodyValueAndCommentBodyFormat(offset int, limit int, Langcode_ string, CommentBodyValue_ string, CommentBodyFormat_ string) (*[]*Comment_commentBody, error) {
	var _Comment_commentBody = new([]*Comment_commentBody)
	err := Engine.Table("comment__comment_body").Where("langcode = ? and comment_body_value = ? and comment_body_format = ?", Langcode_, CommentBodyValue_, CommentBodyFormat_).Limit(limit, offset).Find(_Comment_commentBody)
	return _Comment_commentBody, err
}

// GetComment_commentBodiesByDeltaAndCommentBodyValueAndCommentBodyFormat Get Comment_commentBodies via DeltaAndCommentBodyValueAndCommentBodyFormat
func GetComment_commentBodiesByDeltaAndCommentBodyValueAndCommentBodyFormat(offset int, limit int, Delta_ int, CommentBodyValue_ string, CommentBodyFormat_ string) (*[]*Comment_commentBody, error) {
	var _Comment_commentBody = new([]*Comment_commentBody)
	err := Engine.Table("comment__comment_body").Where("delta = ? and comment_body_value = ? and comment_body_format = ?", Delta_, CommentBodyValue_, CommentBodyFormat_).Limit(limit, offset).Find(_Comment_commentBody)
	return _Comment_commentBody, err
}

// GetComment_commentBodiesByBundleAndDeleted Get Comment_commentBodies via BundleAndDeleted
func GetComment_commentBodiesByBundleAndDeleted(offset int, limit int, Bundle_ string, Deleted_ int) (*[]*Comment_commentBody, error) {
	var _Comment_commentBody = new([]*Comment_commentBody)
	err := Engine.Table("comment__comment_body").Where("bundle = ? and deleted = ?", Bundle_, Deleted_).Limit(limit, offset).Find(_Comment_commentBody)
	return _Comment_commentBody, err
}

// GetComment_commentBodiesByBundleAndEntityId Get Comment_commentBodies via BundleAndEntityId
func GetComment_commentBodiesByBundleAndEntityId(offset int, limit int, Bundle_ string, EntityId_ int) (*[]*Comment_commentBody, error) {
	var _Comment_commentBody = new([]*Comment_commentBody)
	err := Engine.Table("comment__comment_body").Where("bundle = ? and entity_id = ?", Bundle_, EntityId_).Limit(limit, offset).Find(_Comment_commentBody)
	return _Comment_commentBody, err
}

// GetComment_commentBodiesByBundleAndRevisionId Get Comment_commentBodies via BundleAndRevisionId
func GetComment_commentBodiesByBundleAndRevisionId(offset int, limit int, Bundle_ string, RevisionId_ int) (*[]*Comment_commentBody, error) {
	var _Comment_commentBody = new([]*Comment_commentBody)
	err := Engine.Table("comment__comment_body").Where("bundle = ? and revision_id = ?", Bundle_, RevisionId_).Limit(limit, offset).Find(_Comment_commentBody)
	return _Comment_commentBody, err
}

// GetComment_commentBodiesByBundleAndLangcode Get Comment_commentBodies via BundleAndLangcode
func GetComment_commentBodiesByBundleAndLangcode(offset int, limit int, Bundle_ string, Langcode_ string) (*[]*Comment_commentBody, error) {
	var _Comment_commentBody = new([]*Comment_commentBody)
	err := Engine.Table("comment__comment_body").Where("bundle = ? and langcode = ?", Bundle_, Langcode_).Limit(limit, offset).Find(_Comment_commentBody)
	return _Comment_commentBody, err
}

// GetComment_commentBodiesByBundleAndDelta Get Comment_commentBodies via BundleAndDelta
func GetComment_commentBodiesByBundleAndDelta(offset int, limit int, Bundle_ string, Delta_ int) (*[]*Comment_commentBody, error) {
	var _Comment_commentBody = new([]*Comment_commentBody)
	err := Engine.Table("comment__comment_body").Where("bundle = ? and delta = ?", Bundle_, Delta_).Limit(limit, offset).Find(_Comment_commentBody)
	return _Comment_commentBody, err
}

// GetComment_commentBodiesByBundleAndCommentBodyValue Get Comment_commentBodies via BundleAndCommentBodyValue
func GetComment_commentBodiesByBundleAndCommentBodyValue(offset int, limit int, Bundle_ string, CommentBodyValue_ string) (*[]*Comment_commentBody, error) {
	var _Comment_commentBody = new([]*Comment_commentBody)
	err := Engine.Table("comment__comment_body").Where("bundle = ? and comment_body_value = ?", Bundle_, CommentBodyValue_).Limit(limit, offset).Find(_Comment_commentBody)
	return _Comment_commentBody, err
}

// GetComment_commentBodiesByBundleAndCommentBodyFormat Get Comment_commentBodies via BundleAndCommentBodyFormat
func GetComment_commentBodiesByBundleAndCommentBodyFormat(offset int, limit int, Bundle_ string, CommentBodyFormat_ string) (*[]*Comment_commentBody, error) {
	var _Comment_commentBody = new([]*Comment_commentBody)
	err := Engine.Table("comment__comment_body").Where("bundle = ? and comment_body_format = ?", Bundle_, CommentBodyFormat_).Limit(limit, offset).Find(_Comment_commentBody)
	return _Comment_commentBody, err
}

// GetComment_commentBodiesByDeletedAndEntityId Get Comment_commentBodies via DeletedAndEntityId
func GetComment_commentBodiesByDeletedAndEntityId(offset int, limit int, Deleted_ int, EntityId_ int) (*[]*Comment_commentBody, error) {
	var _Comment_commentBody = new([]*Comment_commentBody)
	err := Engine.Table("comment__comment_body").Where("deleted = ? and entity_id = ?", Deleted_, EntityId_).Limit(limit, offset).Find(_Comment_commentBody)
	return _Comment_commentBody, err
}

// GetComment_commentBodiesByDeletedAndRevisionId Get Comment_commentBodies via DeletedAndRevisionId
func GetComment_commentBodiesByDeletedAndRevisionId(offset int, limit int, Deleted_ int, RevisionId_ int) (*[]*Comment_commentBody, error) {
	var _Comment_commentBody = new([]*Comment_commentBody)
	err := Engine.Table("comment__comment_body").Where("deleted = ? and revision_id = ?", Deleted_, RevisionId_).Limit(limit, offset).Find(_Comment_commentBody)
	return _Comment_commentBody, err
}

// GetComment_commentBodiesByDeletedAndLangcode Get Comment_commentBodies via DeletedAndLangcode
func GetComment_commentBodiesByDeletedAndLangcode(offset int, limit int, Deleted_ int, Langcode_ string) (*[]*Comment_commentBody, error) {
	var _Comment_commentBody = new([]*Comment_commentBody)
	err := Engine.Table("comment__comment_body").Where("deleted = ? and langcode = ?", Deleted_, Langcode_).Limit(limit, offset).Find(_Comment_commentBody)
	return _Comment_commentBody, err
}

// GetComment_commentBodiesByDeletedAndDelta Get Comment_commentBodies via DeletedAndDelta
func GetComment_commentBodiesByDeletedAndDelta(offset int, limit int, Deleted_ int, Delta_ int) (*[]*Comment_commentBody, error) {
	var _Comment_commentBody = new([]*Comment_commentBody)
	err := Engine.Table("comment__comment_body").Where("deleted = ? and delta = ?", Deleted_, Delta_).Limit(limit, offset).Find(_Comment_commentBody)
	return _Comment_commentBody, err
}

// GetComment_commentBodiesByDeletedAndCommentBodyValue Get Comment_commentBodies via DeletedAndCommentBodyValue
func GetComment_commentBodiesByDeletedAndCommentBodyValue(offset int, limit int, Deleted_ int, CommentBodyValue_ string) (*[]*Comment_commentBody, error) {
	var _Comment_commentBody = new([]*Comment_commentBody)
	err := Engine.Table("comment__comment_body").Where("deleted = ? and comment_body_value = ?", Deleted_, CommentBodyValue_).Limit(limit, offset).Find(_Comment_commentBody)
	return _Comment_commentBody, err
}

// GetComment_commentBodiesByDeletedAndCommentBodyFormat Get Comment_commentBodies via DeletedAndCommentBodyFormat
func GetComment_commentBodiesByDeletedAndCommentBodyFormat(offset int, limit int, Deleted_ int, CommentBodyFormat_ string) (*[]*Comment_commentBody, error) {
	var _Comment_commentBody = new([]*Comment_commentBody)
	err := Engine.Table("comment__comment_body").Where("deleted = ? and comment_body_format = ?", Deleted_, CommentBodyFormat_).Limit(limit, offset).Find(_Comment_commentBody)
	return _Comment_commentBody, err
}

// GetComment_commentBodiesByEntityIdAndRevisionId Get Comment_commentBodies via EntityIdAndRevisionId
func GetComment_commentBodiesByEntityIdAndRevisionId(offset int, limit int, EntityId_ int, RevisionId_ int) (*[]*Comment_commentBody, error) {
	var _Comment_commentBody = new([]*Comment_commentBody)
	err := Engine.Table("comment__comment_body").Where("entity_id = ? and revision_id = ?", EntityId_, RevisionId_).Limit(limit, offset).Find(_Comment_commentBody)
	return _Comment_commentBody, err
}

// GetComment_commentBodiesByEntityIdAndLangcode Get Comment_commentBodies via EntityIdAndLangcode
func GetComment_commentBodiesByEntityIdAndLangcode(offset int, limit int, EntityId_ int, Langcode_ string) (*[]*Comment_commentBody, error) {
	var _Comment_commentBody = new([]*Comment_commentBody)
	err := Engine.Table("comment__comment_body").Where("entity_id = ? and langcode = ?", EntityId_, Langcode_).Limit(limit, offset).Find(_Comment_commentBody)
	return _Comment_commentBody, err
}

// GetComment_commentBodiesByEntityIdAndDelta Get Comment_commentBodies via EntityIdAndDelta
func GetComment_commentBodiesByEntityIdAndDelta(offset int, limit int, EntityId_ int, Delta_ int) (*[]*Comment_commentBody, error) {
	var _Comment_commentBody = new([]*Comment_commentBody)
	err := Engine.Table("comment__comment_body").Where("entity_id = ? and delta = ?", EntityId_, Delta_).Limit(limit, offset).Find(_Comment_commentBody)
	return _Comment_commentBody, err
}

// GetComment_commentBodiesByEntityIdAndCommentBodyValue Get Comment_commentBodies via EntityIdAndCommentBodyValue
func GetComment_commentBodiesByEntityIdAndCommentBodyValue(offset int, limit int, EntityId_ int, CommentBodyValue_ string) (*[]*Comment_commentBody, error) {
	var _Comment_commentBody = new([]*Comment_commentBody)
	err := Engine.Table("comment__comment_body").Where("entity_id = ? and comment_body_value = ?", EntityId_, CommentBodyValue_).Limit(limit, offset).Find(_Comment_commentBody)
	return _Comment_commentBody, err
}

// GetComment_commentBodiesByEntityIdAndCommentBodyFormat Get Comment_commentBodies via EntityIdAndCommentBodyFormat
func GetComment_commentBodiesByEntityIdAndCommentBodyFormat(offset int, limit int, EntityId_ int, CommentBodyFormat_ string) (*[]*Comment_commentBody, error) {
	var _Comment_commentBody = new([]*Comment_commentBody)
	err := Engine.Table("comment__comment_body").Where("entity_id = ? and comment_body_format = ?", EntityId_, CommentBodyFormat_).Limit(limit, offset).Find(_Comment_commentBody)
	return _Comment_commentBody, err
}

// GetComment_commentBodiesByRevisionIdAndLangcode Get Comment_commentBodies via RevisionIdAndLangcode
func GetComment_commentBodiesByRevisionIdAndLangcode(offset int, limit int, RevisionId_ int, Langcode_ string) (*[]*Comment_commentBody, error) {
	var _Comment_commentBody = new([]*Comment_commentBody)
	err := Engine.Table("comment__comment_body").Where("revision_id = ? and langcode = ?", RevisionId_, Langcode_).Limit(limit, offset).Find(_Comment_commentBody)
	return _Comment_commentBody, err
}

// GetComment_commentBodiesByRevisionIdAndDelta Get Comment_commentBodies via RevisionIdAndDelta
func GetComment_commentBodiesByRevisionIdAndDelta(offset int, limit int, RevisionId_ int, Delta_ int) (*[]*Comment_commentBody, error) {
	var _Comment_commentBody = new([]*Comment_commentBody)
	err := Engine.Table("comment__comment_body").Where("revision_id = ? and delta = ?", RevisionId_, Delta_).Limit(limit, offset).Find(_Comment_commentBody)
	return _Comment_commentBody, err
}

// GetComment_commentBodiesByRevisionIdAndCommentBodyValue Get Comment_commentBodies via RevisionIdAndCommentBodyValue
func GetComment_commentBodiesByRevisionIdAndCommentBodyValue(offset int, limit int, RevisionId_ int, CommentBodyValue_ string) (*[]*Comment_commentBody, error) {
	var _Comment_commentBody = new([]*Comment_commentBody)
	err := Engine.Table("comment__comment_body").Where("revision_id = ? and comment_body_value = ?", RevisionId_, CommentBodyValue_).Limit(limit, offset).Find(_Comment_commentBody)
	return _Comment_commentBody, err
}

// GetComment_commentBodiesByRevisionIdAndCommentBodyFormat Get Comment_commentBodies via RevisionIdAndCommentBodyFormat
func GetComment_commentBodiesByRevisionIdAndCommentBodyFormat(offset int, limit int, RevisionId_ int, CommentBodyFormat_ string) (*[]*Comment_commentBody, error) {
	var _Comment_commentBody = new([]*Comment_commentBody)
	err := Engine.Table("comment__comment_body").Where("revision_id = ? and comment_body_format = ?", RevisionId_, CommentBodyFormat_).Limit(limit, offset).Find(_Comment_commentBody)
	return _Comment_commentBody, err
}

// GetComment_commentBodiesByLangcodeAndDelta Get Comment_commentBodies via LangcodeAndDelta
func GetComment_commentBodiesByLangcodeAndDelta(offset int, limit int, Langcode_ string, Delta_ int) (*[]*Comment_commentBody, error) {
	var _Comment_commentBody = new([]*Comment_commentBody)
	err := Engine.Table("comment__comment_body").Where("langcode = ? and delta = ?", Langcode_, Delta_).Limit(limit, offset).Find(_Comment_commentBody)
	return _Comment_commentBody, err
}

// GetComment_commentBodiesByLangcodeAndCommentBodyValue Get Comment_commentBodies via LangcodeAndCommentBodyValue
func GetComment_commentBodiesByLangcodeAndCommentBodyValue(offset int, limit int, Langcode_ string, CommentBodyValue_ string) (*[]*Comment_commentBody, error) {
	var _Comment_commentBody = new([]*Comment_commentBody)
	err := Engine.Table("comment__comment_body").Where("langcode = ? and comment_body_value = ?", Langcode_, CommentBodyValue_).Limit(limit, offset).Find(_Comment_commentBody)
	return _Comment_commentBody, err
}

// GetComment_commentBodiesByLangcodeAndCommentBodyFormat Get Comment_commentBodies via LangcodeAndCommentBodyFormat
func GetComment_commentBodiesByLangcodeAndCommentBodyFormat(offset int, limit int, Langcode_ string, CommentBodyFormat_ string) (*[]*Comment_commentBody, error) {
	var _Comment_commentBody = new([]*Comment_commentBody)
	err := Engine.Table("comment__comment_body").Where("langcode = ? and comment_body_format = ?", Langcode_, CommentBodyFormat_).Limit(limit, offset).Find(_Comment_commentBody)
	return _Comment_commentBody, err
}

// GetComment_commentBodiesByDeltaAndCommentBodyValue Get Comment_commentBodies via DeltaAndCommentBodyValue
func GetComment_commentBodiesByDeltaAndCommentBodyValue(offset int, limit int, Delta_ int, CommentBodyValue_ string) (*[]*Comment_commentBody, error) {
	var _Comment_commentBody = new([]*Comment_commentBody)
	err := Engine.Table("comment__comment_body").Where("delta = ? and comment_body_value = ?", Delta_, CommentBodyValue_).Limit(limit, offset).Find(_Comment_commentBody)
	return _Comment_commentBody, err
}

// GetComment_commentBodiesByDeltaAndCommentBodyFormat Get Comment_commentBodies via DeltaAndCommentBodyFormat
func GetComment_commentBodiesByDeltaAndCommentBodyFormat(offset int, limit int, Delta_ int, CommentBodyFormat_ string) (*[]*Comment_commentBody, error) {
	var _Comment_commentBody = new([]*Comment_commentBody)
	err := Engine.Table("comment__comment_body").Where("delta = ? and comment_body_format = ?", Delta_, CommentBodyFormat_).Limit(limit, offset).Find(_Comment_commentBody)
	return _Comment_commentBody, err
}

// GetComment_commentBodiesByCommentBodyValueAndCommentBodyFormat Get Comment_commentBodies via CommentBodyValueAndCommentBodyFormat
func GetComment_commentBodiesByCommentBodyValueAndCommentBodyFormat(offset int, limit int, CommentBodyValue_ string, CommentBodyFormat_ string) (*[]*Comment_commentBody, error) {
	var _Comment_commentBody = new([]*Comment_commentBody)
	err := Engine.Table("comment__comment_body").Where("comment_body_value = ? and comment_body_format = ?", CommentBodyValue_, CommentBodyFormat_).Limit(limit, offset).Find(_Comment_commentBody)
	return _Comment_commentBody, err
}

// GetComment_commentBodies Get Comment_commentBodies via field
func GetComment_commentBodies(offset int, limit int, field string) (*[]*Comment_commentBody, error) {
	var _Comment_commentBody = new([]*Comment_commentBody)
	err := Engine.Table("comment__comment_body").Limit(limit, offset).Desc(field).Find(_Comment_commentBody)
	return _Comment_commentBody, err
}

// GetComment_commentBodiesViaBundle Get Comment_commentBodys via Bundle
func GetComment_commentBodiesViaBundle(offset int, limit int, Bundle_ string, field string) (*[]*Comment_commentBody, error) {
	var _Comment_commentBody = new([]*Comment_commentBody)
	err := Engine.Table("comment__comment_body").Where("bundle = ?", Bundle_).Limit(limit, offset).Desc(field).Find(_Comment_commentBody)
	return _Comment_commentBody, err
}

// GetComment_commentBodiesViaDeleted Get Comment_commentBodys via Deleted
func GetComment_commentBodiesViaDeleted(offset int, limit int, Deleted_ int, field string) (*[]*Comment_commentBody, error) {
	var _Comment_commentBody = new([]*Comment_commentBody)
	err := Engine.Table("comment__comment_body").Where("deleted = ?", Deleted_).Limit(limit, offset).Desc(field).Find(_Comment_commentBody)
	return _Comment_commentBody, err
}

// GetComment_commentBodiesViaEntityId Get Comment_commentBodys via EntityId
func GetComment_commentBodiesViaEntityId(offset int, limit int, EntityId_ int, field string) (*[]*Comment_commentBody, error) {
	var _Comment_commentBody = new([]*Comment_commentBody)
	err := Engine.Table("comment__comment_body").Where("entity_id = ?", EntityId_).Limit(limit, offset).Desc(field).Find(_Comment_commentBody)
	return _Comment_commentBody, err
}

// GetComment_commentBodiesViaRevisionId Get Comment_commentBodys via RevisionId
func GetComment_commentBodiesViaRevisionId(offset int, limit int, RevisionId_ int, field string) (*[]*Comment_commentBody, error) {
	var _Comment_commentBody = new([]*Comment_commentBody)
	err := Engine.Table("comment__comment_body").Where("revision_id = ?", RevisionId_).Limit(limit, offset).Desc(field).Find(_Comment_commentBody)
	return _Comment_commentBody, err
}

// GetComment_commentBodiesViaLangcode Get Comment_commentBodys via Langcode
func GetComment_commentBodiesViaLangcode(offset int, limit int, Langcode_ string, field string) (*[]*Comment_commentBody, error) {
	var _Comment_commentBody = new([]*Comment_commentBody)
	err := Engine.Table("comment__comment_body").Where("langcode = ?", Langcode_).Limit(limit, offset).Desc(field).Find(_Comment_commentBody)
	return _Comment_commentBody, err
}

// GetComment_commentBodiesViaDelta Get Comment_commentBodys via Delta
func GetComment_commentBodiesViaDelta(offset int, limit int, Delta_ int, field string) (*[]*Comment_commentBody, error) {
	var _Comment_commentBody = new([]*Comment_commentBody)
	err := Engine.Table("comment__comment_body").Where("delta = ?", Delta_).Limit(limit, offset).Desc(field).Find(_Comment_commentBody)
	return _Comment_commentBody, err
}

// GetComment_commentBodiesViaCommentBodyValue Get Comment_commentBodys via CommentBodyValue
func GetComment_commentBodiesViaCommentBodyValue(offset int, limit int, CommentBodyValue_ string, field string) (*[]*Comment_commentBody, error) {
	var _Comment_commentBody = new([]*Comment_commentBody)
	err := Engine.Table("comment__comment_body").Where("comment_body_value = ?", CommentBodyValue_).Limit(limit, offset).Desc(field).Find(_Comment_commentBody)
	return _Comment_commentBody, err
}

// GetComment_commentBodiesViaCommentBodyFormat Get Comment_commentBodys via CommentBodyFormat
func GetComment_commentBodiesViaCommentBodyFormat(offset int, limit int, CommentBodyFormat_ string, field string) (*[]*Comment_commentBody, error) {
	var _Comment_commentBody = new([]*Comment_commentBody)
	err := Engine.Table("comment__comment_body").Where("comment_body_format = ?", CommentBodyFormat_).Limit(limit, offset).Desc(field).Find(_Comment_commentBody)
	return _Comment_commentBody, err
}

// HasComment_commentBodyViaBundle Has Comment_commentBody via Bundle
func HasComment_commentBodyViaBundle(iBundle string) bool {
	if has, err := Engine.Where("bundle = ?", iBundle).Get(new(Comment_commentBody)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasComment_commentBodyViaDeleted Has Comment_commentBody via Deleted
func HasComment_commentBodyViaDeleted(iDeleted int) bool {
	if has, err := Engine.Where("deleted = ?", iDeleted).Get(new(Comment_commentBody)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasComment_commentBodyViaEntityId Has Comment_commentBody via EntityId
func HasComment_commentBodyViaEntityId(iEntityId int) bool {
	if has, err := Engine.Where("entity_id = ?", iEntityId).Get(new(Comment_commentBody)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasComment_commentBodyViaRevisionId Has Comment_commentBody via RevisionId
func HasComment_commentBodyViaRevisionId(iRevisionId int) bool {
	if has, err := Engine.Where("revision_id = ?", iRevisionId).Get(new(Comment_commentBody)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasComment_commentBodyViaLangcode Has Comment_commentBody via Langcode
func HasComment_commentBodyViaLangcode(iLangcode string) bool {
	if has, err := Engine.Where("langcode = ?", iLangcode).Get(new(Comment_commentBody)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasComment_commentBodyViaDelta Has Comment_commentBody via Delta
func HasComment_commentBodyViaDelta(iDelta int) bool {
	if has, err := Engine.Where("delta = ?", iDelta).Get(new(Comment_commentBody)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasComment_commentBodyViaCommentBodyValue Has Comment_commentBody via CommentBodyValue
func HasComment_commentBodyViaCommentBodyValue(iCommentBodyValue string) bool {
	if has, err := Engine.Where("comment_body_value = ?", iCommentBodyValue).Get(new(Comment_commentBody)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasComment_commentBodyViaCommentBodyFormat Has Comment_commentBody via CommentBodyFormat
func HasComment_commentBodyViaCommentBodyFormat(iCommentBodyFormat string) bool {
	if has, err := Engine.Where("comment_body_format = ?", iCommentBodyFormat).Get(new(Comment_commentBody)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// GetComment_commentBodyViaBundle Get Comment_commentBody via Bundle
func GetComment_commentBodyViaBundle(iBundle string) (*Comment_commentBody, error) {
	var _Comment_commentBody = &Comment_commentBody{Bundle: iBundle}
	has, err := Engine.Get(_Comment_commentBody)
	if has {
		return _Comment_commentBody, err
	} else {
		return nil, err
	}
}

// GetComment_commentBodyViaDeleted Get Comment_commentBody via Deleted
func GetComment_commentBodyViaDeleted(iDeleted int) (*Comment_commentBody, error) {
	var _Comment_commentBody = &Comment_commentBody{Deleted: iDeleted}
	has, err := Engine.Get(_Comment_commentBody)
	if has {
		return _Comment_commentBody, err
	} else {
		return nil, err
	}
}

// GetComment_commentBodyViaEntityId Get Comment_commentBody via EntityId
func GetComment_commentBodyViaEntityId(iEntityId int) (*Comment_commentBody, error) {
	var _Comment_commentBody = &Comment_commentBody{EntityId: iEntityId}
	has, err := Engine.Get(_Comment_commentBody)
	if has {
		return _Comment_commentBody, err
	} else {
		return nil, err
	}
}

// GetComment_commentBodyViaRevisionId Get Comment_commentBody via RevisionId
func GetComment_commentBodyViaRevisionId(iRevisionId int) (*Comment_commentBody, error) {
	var _Comment_commentBody = &Comment_commentBody{RevisionId: iRevisionId}
	has, err := Engine.Get(_Comment_commentBody)
	if has {
		return _Comment_commentBody, err
	} else {
		return nil, err
	}
}

// GetComment_commentBodyViaLangcode Get Comment_commentBody via Langcode
func GetComment_commentBodyViaLangcode(iLangcode string) (*Comment_commentBody, error) {
	var _Comment_commentBody = &Comment_commentBody{Langcode: iLangcode}
	has, err := Engine.Get(_Comment_commentBody)
	if has {
		return _Comment_commentBody, err
	} else {
		return nil, err
	}
}

// GetComment_commentBodyViaDelta Get Comment_commentBody via Delta
func GetComment_commentBodyViaDelta(iDelta int) (*Comment_commentBody, error) {
	var _Comment_commentBody = &Comment_commentBody{Delta: iDelta}
	has, err := Engine.Get(_Comment_commentBody)
	if has {
		return _Comment_commentBody, err
	} else {
		return nil, err
	}
}

// GetComment_commentBodyViaCommentBodyValue Get Comment_commentBody via CommentBodyValue
func GetComment_commentBodyViaCommentBodyValue(iCommentBodyValue string) (*Comment_commentBody, error) {
	var _Comment_commentBody = &Comment_commentBody{CommentBodyValue: iCommentBodyValue}
	has, err := Engine.Get(_Comment_commentBody)
	if has {
		return _Comment_commentBody, err
	} else {
		return nil, err
	}
}

// GetComment_commentBodyViaCommentBodyFormat Get Comment_commentBody via CommentBodyFormat
func GetComment_commentBodyViaCommentBodyFormat(iCommentBodyFormat string) (*Comment_commentBody, error) {
	var _Comment_commentBody = &Comment_commentBody{CommentBodyFormat: iCommentBodyFormat}
	has, err := Engine.Get(_Comment_commentBody)
	if has {
		return _Comment_commentBody, err
	} else {
		return nil, err
	}
}

// SetComment_commentBodyViaBundle Set Comment_commentBody via Bundle
func SetComment_commentBodyViaBundle(iBundle string, comment__comment_body *Comment_commentBody) (int64, error) {
	comment__comment_body.Bundle = iBundle
	return Engine.Insert(comment__comment_body)
}

// SetComment_commentBodyViaDeleted Set Comment_commentBody via Deleted
func SetComment_commentBodyViaDeleted(iDeleted int, comment__comment_body *Comment_commentBody) (int64, error) {
	comment__comment_body.Deleted = iDeleted
	return Engine.Insert(comment__comment_body)
}

// SetComment_commentBodyViaEntityId Set Comment_commentBody via EntityId
func SetComment_commentBodyViaEntityId(iEntityId int, comment__comment_body *Comment_commentBody) (int64, error) {
	comment__comment_body.EntityId = iEntityId
	return Engine.Insert(comment__comment_body)
}

// SetComment_commentBodyViaRevisionId Set Comment_commentBody via RevisionId
func SetComment_commentBodyViaRevisionId(iRevisionId int, comment__comment_body *Comment_commentBody) (int64, error) {
	comment__comment_body.RevisionId = iRevisionId
	return Engine.Insert(comment__comment_body)
}

// SetComment_commentBodyViaLangcode Set Comment_commentBody via Langcode
func SetComment_commentBodyViaLangcode(iLangcode string, comment__comment_body *Comment_commentBody) (int64, error) {
	comment__comment_body.Langcode = iLangcode
	return Engine.Insert(comment__comment_body)
}

// SetComment_commentBodyViaDelta Set Comment_commentBody via Delta
func SetComment_commentBodyViaDelta(iDelta int, comment__comment_body *Comment_commentBody) (int64, error) {
	comment__comment_body.Delta = iDelta
	return Engine.Insert(comment__comment_body)
}

// SetComment_commentBodyViaCommentBodyValue Set Comment_commentBody via CommentBodyValue
func SetComment_commentBodyViaCommentBodyValue(iCommentBodyValue string, comment__comment_body *Comment_commentBody) (int64, error) {
	comment__comment_body.CommentBodyValue = iCommentBodyValue
	return Engine.Insert(comment__comment_body)
}

// SetComment_commentBodyViaCommentBodyFormat Set Comment_commentBody via CommentBodyFormat
func SetComment_commentBodyViaCommentBodyFormat(iCommentBodyFormat string, comment__comment_body *Comment_commentBody) (int64, error) {
	comment__comment_body.CommentBodyFormat = iCommentBodyFormat
	return Engine.Insert(comment__comment_body)
}

// AddComment_commentBody Add Comment_commentBody via all columns
func AddComment_commentBody(iBundle string, iDeleted int, iEntityId int, iRevisionId int, iLangcode string, iDelta int, iCommentBodyValue string, iCommentBodyFormat string) error {
	_Comment_commentBody := &Comment_commentBody{Bundle: iBundle, Deleted: iDeleted, EntityId: iEntityId, RevisionId: iRevisionId, Langcode: iLangcode, Delta: iDelta, CommentBodyValue: iCommentBodyValue, CommentBodyFormat: iCommentBodyFormat}
	if _, err := Engine.Insert(_Comment_commentBody); err != nil {
		return err
	} else {
		return nil
	}
}

// PostComment_commentBody Post Comment_commentBody via iComment_commentBody
func PostComment_commentBody(iComment_commentBody *Comment_commentBody) (string, error) {
	_, err := Engine.Insert(iComment_commentBody)
	return iComment_commentBody.Bundle, err
}

// PutComment_commentBody Put Comment_commentBody
func PutComment_commentBody(iComment_commentBody *Comment_commentBody) (string, error) {
	_, err := Engine.Id(iComment_commentBody.Bundle).Update(iComment_commentBody)
	return iComment_commentBody.Bundle, err
}

// PutComment_commentBodyViaBundle Put Comment_commentBody via Bundle
func PutComment_commentBodyViaBundle(Bundle_ string, iComment_commentBody *Comment_commentBody) (int64, error) {
	row, err := Engine.Update(iComment_commentBody, &Comment_commentBody{Bundle: Bundle_})
	return row, err
}

// PutComment_commentBodyViaDeleted Put Comment_commentBody via Deleted
func PutComment_commentBodyViaDeleted(Deleted_ int, iComment_commentBody *Comment_commentBody) (int64, error) {
	row, err := Engine.Update(iComment_commentBody, &Comment_commentBody{Deleted: Deleted_})
	return row, err
}

// PutComment_commentBodyViaEntityId Put Comment_commentBody via EntityId
func PutComment_commentBodyViaEntityId(EntityId_ int, iComment_commentBody *Comment_commentBody) (int64, error) {
	row, err := Engine.Update(iComment_commentBody, &Comment_commentBody{EntityId: EntityId_})
	return row, err
}

// PutComment_commentBodyViaRevisionId Put Comment_commentBody via RevisionId
func PutComment_commentBodyViaRevisionId(RevisionId_ int, iComment_commentBody *Comment_commentBody) (int64, error) {
	row, err := Engine.Update(iComment_commentBody, &Comment_commentBody{RevisionId: RevisionId_})
	return row, err
}

// PutComment_commentBodyViaLangcode Put Comment_commentBody via Langcode
func PutComment_commentBodyViaLangcode(Langcode_ string, iComment_commentBody *Comment_commentBody) (int64, error) {
	row, err := Engine.Update(iComment_commentBody, &Comment_commentBody{Langcode: Langcode_})
	return row, err
}

// PutComment_commentBodyViaDelta Put Comment_commentBody via Delta
func PutComment_commentBodyViaDelta(Delta_ int, iComment_commentBody *Comment_commentBody) (int64, error) {
	row, err := Engine.Update(iComment_commentBody, &Comment_commentBody{Delta: Delta_})
	return row, err
}

// PutComment_commentBodyViaCommentBodyValue Put Comment_commentBody via CommentBodyValue
func PutComment_commentBodyViaCommentBodyValue(CommentBodyValue_ string, iComment_commentBody *Comment_commentBody) (int64, error) {
	row, err := Engine.Update(iComment_commentBody, &Comment_commentBody{CommentBodyValue: CommentBodyValue_})
	return row, err
}

// PutComment_commentBodyViaCommentBodyFormat Put Comment_commentBody via CommentBodyFormat
func PutComment_commentBodyViaCommentBodyFormat(CommentBodyFormat_ string, iComment_commentBody *Comment_commentBody) (int64, error) {
	row, err := Engine.Update(iComment_commentBody, &Comment_commentBody{CommentBodyFormat: CommentBodyFormat_})
	return row, err
}

// UpdateComment_commentBodyViaBundle via map[string]interface{}{}
func UpdateComment_commentBodyViaBundle(iBundle string, iComment_commentBodyMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Comment_commentBody)).Where("bundle = ?", iBundle).Update(iComment_commentBodyMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateComment_commentBodyViaDeleted via map[string]interface{}{}
func UpdateComment_commentBodyViaDeleted(iDeleted int, iComment_commentBodyMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Comment_commentBody)).Where("deleted = ?", iDeleted).Update(iComment_commentBodyMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateComment_commentBodyViaEntityId via map[string]interface{}{}
func UpdateComment_commentBodyViaEntityId(iEntityId int, iComment_commentBodyMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Comment_commentBody)).Where("entity_id = ?", iEntityId).Update(iComment_commentBodyMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateComment_commentBodyViaRevisionId via map[string]interface{}{}
func UpdateComment_commentBodyViaRevisionId(iRevisionId int, iComment_commentBodyMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Comment_commentBody)).Where("revision_id = ?", iRevisionId).Update(iComment_commentBodyMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateComment_commentBodyViaLangcode via map[string]interface{}{}
func UpdateComment_commentBodyViaLangcode(iLangcode string, iComment_commentBodyMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Comment_commentBody)).Where("langcode = ?", iLangcode).Update(iComment_commentBodyMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateComment_commentBodyViaDelta via map[string]interface{}{}
func UpdateComment_commentBodyViaDelta(iDelta int, iComment_commentBodyMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Comment_commentBody)).Where("delta = ?", iDelta).Update(iComment_commentBodyMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateComment_commentBodyViaCommentBodyValue via map[string]interface{}{}
func UpdateComment_commentBodyViaCommentBodyValue(iCommentBodyValue string, iComment_commentBodyMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Comment_commentBody)).Where("comment_body_value = ?", iCommentBodyValue).Update(iComment_commentBodyMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateComment_commentBodyViaCommentBodyFormat via map[string]interface{}{}
func UpdateComment_commentBodyViaCommentBodyFormat(iCommentBodyFormat string, iComment_commentBodyMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Comment_commentBody)).Where("comment_body_format = ?", iCommentBodyFormat).Update(iComment_commentBodyMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// DeleteComment_commentBody Delete Comment_commentBody
func DeleteComment_commentBody(iBundle string) (int64, error) {
	row, err := Engine.Id(iBundle).Delete(new(Comment_commentBody))
	return row, err
}

// DeleteComment_commentBodyViaBundle Delete Comment_commentBody via Bundle
func DeleteComment_commentBodyViaBundle(iBundle string) (err error) {
	var has bool
	var _Comment_commentBody = &Comment_commentBody{Bundle: iBundle}
	if has, err = Engine.Get(_Comment_commentBody); (has == true) && (err == nil) {
		if row, err := Engine.Where("bundle = ?", iBundle).Delete(new(Comment_commentBody)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteComment_commentBodyViaDeleted Delete Comment_commentBody via Deleted
func DeleteComment_commentBodyViaDeleted(iDeleted int) (err error) {
	var has bool
	var _Comment_commentBody = &Comment_commentBody{Deleted: iDeleted}
	if has, err = Engine.Get(_Comment_commentBody); (has == true) && (err == nil) {
		if row, err := Engine.Where("deleted = ?", iDeleted).Delete(new(Comment_commentBody)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteComment_commentBodyViaEntityId Delete Comment_commentBody via EntityId
func DeleteComment_commentBodyViaEntityId(iEntityId int) (err error) {
	var has bool
	var _Comment_commentBody = &Comment_commentBody{EntityId: iEntityId}
	if has, err = Engine.Get(_Comment_commentBody); (has == true) && (err == nil) {
		if row, err := Engine.Where("entity_id = ?", iEntityId).Delete(new(Comment_commentBody)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteComment_commentBodyViaRevisionId Delete Comment_commentBody via RevisionId
func DeleteComment_commentBodyViaRevisionId(iRevisionId int) (err error) {
	var has bool
	var _Comment_commentBody = &Comment_commentBody{RevisionId: iRevisionId}
	if has, err = Engine.Get(_Comment_commentBody); (has == true) && (err == nil) {
		if row, err := Engine.Where("revision_id = ?", iRevisionId).Delete(new(Comment_commentBody)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteComment_commentBodyViaLangcode Delete Comment_commentBody via Langcode
func DeleteComment_commentBodyViaLangcode(iLangcode string) (err error) {
	var has bool
	var _Comment_commentBody = &Comment_commentBody{Langcode: iLangcode}
	if has, err = Engine.Get(_Comment_commentBody); (has == true) && (err == nil) {
		if row, err := Engine.Where("langcode = ?", iLangcode).Delete(new(Comment_commentBody)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteComment_commentBodyViaDelta Delete Comment_commentBody via Delta
func DeleteComment_commentBodyViaDelta(iDelta int) (err error) {
	var has bool
	var _Comment_commentBody = &Comment_commentBody{Delta: iDelta}
	if has, err = Engine.Get(_Comment_commentBody); (has == true) && (err == nil) {
		if row, err := Engine.Where("delta = ?", iDelta).Delete(new(Comment_commentBody)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteComment_commentBodyViaCommentBodyValue Delete Comment_commentBody via CommentBodyValue
func DeleteComment_commentBodyViaCommentBodyValue(iCommentBodyValue string) (err error) {
	var has bool
	var _Comment_commentBody = &Comment_commentBody{CommentBodyValue: iCommentBodyValue}
	if has, err = Engine.Get(_Comment_commentBody); (has == true) && (err == nil) {
		if row, err := Engine.Where("comment_body_value = ?", iCommentBodyValue).Delete(new(Comment_commentBody)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteComment_commentBodyViaCommentBodyFormat Delete Comment_commentBody via CommentBodyFormat
func DeleteComment_commentBodyViaCommentBodyFormat(iCommentBodyFormat string) (err error) {
	var has bool
	var _Comment_commentBody = &Comment_commentBody{CommentBodyFormat: iCommentBodyFormat}
	if has, err = Engine.Get(_Comment_commentBody); (has == true) && (err == nil) {
		if row, err := Engine.Where("comment_body_format = ?", iCommentBodyFormat).Delete(new(Comment_commentBody)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}
