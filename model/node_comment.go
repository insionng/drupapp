package model

type Node_comment struct {
	Bundle        string `xorm:"not null default '' index VARCHAR(128)"`
	Deleted       int    `xorm:"not null pk default 0 TINYINT(4)"`
	EntityId      int    `xorm:"not null pk INT(10)"`
	RevisionId    int    `xorm:"not null index INT(10)"`
	Langcode      string `xorm:"not null pk default '' VARCHAR(32)"`
	Delta         int    `xorm:"not null pk INT(10)"`
	CommentStatus int    `xorm:"not null default 0 INT(11)"`
}

// GetNode_commentsCount Node_comments' Count
func GetNode_commentsCount(offset int, limit int) (int64, error) {
	total, err := Engine.Limit(limit, offset).Count(&Node_comment{})
	return total, err
}

// GetNode_commentCountViaBundle Get Node_comment via Bundle
func GetNode_commentCountViaBundle(iBundle string) int64 {
	n, _ := Engine.Where("bundle = ?", iBundle).Count(&Node_comment{Bundle: iBundle})
	return n
}

// GetNode_commentCountViaDeleted Get Node_comment via Deleted
func GetNode_commentCountViaDeleted(iDeleted int) int64 {
	n, _ := Engine.Where("deleted = ?", iDeleted).Count(&Node_comment{Deleted: iDeleted})
	return n
}

// GetNode_commentCountViaEntityId Get Node_comment via EntityId
func GetNode_commentCountViaEntityId(iEntityId int) int64 {
	n, _ := Engine.Where("entity_id = ?", iEntityId).Count(&Node_comment{EntityId: iEntityId})
	return n
}

// GetNode_commentCountViaRevisionId Get Node_comment via RevisionId
func GetNode_commentCountViaRevisionId(iRevisionId int) int64 {
	n, _ := Engine.Where("revision_id = ?", iRevisionId).Count(&Node_comment{RevisionId: iRevisionId})
	return n
}

// GetNode_commentCountViaLangcode Get Node_comment via Langcode
func GetNode_commentCountViaLangcode(iLangcode string) int64 {
	n, _ := Engine.Where("langcode = ?", iLangcode).Count(&Node_comment{Langcode: iLangcode})
	return n
}

// GetNode_commentCountViaDelta Get Node_comment via Delta
func GetNode_commentCountViaDelta(iDelta int) int64 {
	n, _ := Engine.Where("delta = ?", iDelta).Count(&Node_comment{Delta: iDelta})
	return n
}

// GetNode_commentCountViaCommentStatus Get Node_comment via CommentStatus
func GetNode_commentCountViaCommentStatus(iCommentStatus int) int64 {
	n, _ := Engine.Where("comment_status = ?", iCommentStatus).Count(&Node_comment{CommentStatus: iCommentStatus})
	return n
}

// GetNode_commentsByBundleAndDeletedAndEntityId Get Node_comments via BundleAndDeletedAndEntityId
func GetNode_commentsByBundleAndDeletedAndEntityId(offset int, limit int, Bundle_ string, Deleted_ int, EntityId_ int) (*[]*Node_comment, error) {
	var _Node_comment = new([]*Node_comment)
	err := Engine.Table("node__comment").Where("bundle = ? and deleted = ? and entity_id = ?", Bundle_, Deleted_, EntityId_).Limit(limit, offset).Find(_Node_comment)
	return _Node_comment, err
}

// GetNode_commentsByBundleAndDeletedAndRevisionId Get Node_comments via BundleAndDeletedAndRevisionId
func GetNode_commentsByBundleAndDeletedAndRevisionId(offset int, limit int, Bundle_ string, Deleted_ int, RevisionId_ int) (*[]*Node_comment, error) {
	var _Node_comment = new([]*Node_comment)
	err := Engine.Table("node__comment").Where("bundle = ? and deleted = ? and revision_id = ?", Bundle_, Deleted_, RevisionId_).Limit(limit, offset).Find(_Node_comment)
	return _Node_comment, err
}

// GetNode_commentsByBundleAndDeletedAndLangcode Get Node_comments via BundleAndDeletedAndLangcode
func GetNode_commentsByBundleAndDeletedAndLangcode(offset int, limit int, Bundle_ string, Deleted_ int, Langcode_ string) (*[]*Node_comment, error) {
	var _Node_comment = new([]*Node_comment)
	err := Engine.Table("node__comment").Where("bundle = ? and deleted = ? and langcode = ?", Bundle_, Deleted_, Langcode_).Limit(limit, offset).Find(_Node_comment)
	return _Node_comment, err
}

// GetNode_commentsByBundleAndDeletedAndDelta Get Node_comments via BundleAndDeletedAndDelta
func GetNode_commentsByBundleAndDeletedAndDelta(offset int, limit int, Bundle_ string, Deleted_ int, Delta_ int) (*[]*Node_comment, error) {
	var _Node_comment = new([]*Node_comment)
	err := Engine.Table("node__comment").Where("bundle = ? and deleted = ? and delta = ?", Bundle_, Deleted_, Delta_).Limit(limit, offset).Find(_Node_comment)
	return _Node_comment, err
}

// GetNode_commentsByBundleAndDeletedAndCommentStatus Get Node_comments via BundleAndDeletedAndCommentStatus
func GetNode_commentsByBundleAndDeletedAndCommentStatus(offset int, limit int, Bundle_ string, Deleted_ int, CommentStatus_ int) (*[]*Node_comment, error) {
	var _Node_comment = new([]*Node_comment)
	err := Engine.Table("node__comment").Where("bundle = ? and deleted = ? and comment_status = ?", Bundle_, Deleted_, CommentStatus_).Limit(limit, offset).Find(_Node_comment)
	return _Node_comment, err
}

// GetNode_commentsByBundleAndEntityIdAndRevisionId Get Node_comments via BundleAndEntityIdAndRevisionId
func GetNode_commentsByBundleAndEntityIdAndRevisionId(offset int, limit int, Bundle_ string, EntityId_ int, RevisionId_ int) (*[]*Node_comment, error) {
	var _Node_comment = new([]*Node_comment)
	err := Engine.Table("node__comment").Where("bundle = ? and entity_id = ? and revision_id = ?", Bundle_, EntityId_, RevisionId_).Limit(limit, offset).Find(_Node_comment)
	return _Node_comment, err
}

// GetNode_commentsByBundleAndEntityIdAndLangcode Get Node_comments via BundleAndEntityIdAndLangcode
func GetNode_commentsByBundleAndEntityIdAndLangcode(offset int, limit int, Bundle_ string, EntityId_ int, Langcode_ string) (*[]*Node_comment, error) {
	var _Node_comment = new([]*Node_comment)
	err := Engine.Table("node__comment").Where("bundle = ? and entity_id = ? and langcode = ?", Bundle_, EntityId_, Langcode_).Limit(limit, offset).Find(_Node_comment)
	return _Node_comment, err
}

// GetNode_commentsByBundleAndEntityIdAndDelta Get Node_comments via BundleAndEntityIdAndDelta
func GetNode_commentsByBundleAndEntityIdAndDelta(offset int, limit int, Bundle_ string, EntityId_ int, Delta_ int) (*[]*Node_comment, error) {
	var _Node_comment = new([]*Node_comment)
	err := Engine.Table("node__comment").Where("bundle = ? and entity_id = ? and delta = ?", Bundle_, EntityId_, Delta_).Limit(limit, offset).Find(_Node_comment)
	return _Node_comment, err
}

// GetNode_commentsByBundleAndEntityIdAndCommentStatus Get Node_comments via BundleAndEntityIdAndCommentStatus
func GetNode_commentsByBundleAndEntityIdAndCommentStatus(offset int, limit int, Bundle_ string, EntityId_ int, CommentStatus_ int) (*[]*Node_comment, error) {
	var _Node_comment = new([]*Node_comment)
	err := Engine.Table("node__comment").Where("bundle = ? and entity_id = ? and comment_status = ?", Bundle_, EntityId_, CommentStatus_).Limit(limit, offset).Find(_Node_comment)
	return _Node_comment, err
}

// GetNode_commentsByBundleAndRevisionIdAndLangcode Get Node_comments via BundleAndRevisionIdAndLangcode
func GetNode_commentsByBundleAndRevisionIdAndLangcode(offset int, limit int, Bundle_ string, RevisionId_ int, Langcode_ string) (*[]*Node_comment, error) {
	var _Node_comment = new([]*Node_comment)
	err := Engine.Table("node__comment").Where("bundle = ? and revision_id = ? and langcode = ?", Bundle_, RevisionId_, Langcode_).Limit(limit, offset).Find(_Node_comment)
	return _Node_comment, err
}

// GetNode_commentsByBundleAndRevisionIdAndDelta Get Node_comments via BundleAndRevisionIdAndDelta
func GetNode_commentsByBundleAndRevisionIdAndDelta(offset int, limit int, Bundle_ string, RevisionId_ int, Delta_ int) (*[]*Node_comment, error) {
	var _Node_comment = new([]*Node_comment)
	err := Engine.Table("node__comment").Where("bundle = ? and revision_id = ? and delta = ?", Bundle_, RevisionId_, Delta_).Limit(limit, offset).Find(_Node_comment)
	return _Node_comment, err
}

// GetNode_commentsByBundleAndRevisionIdAndCommentStatus Get Node_comments via BundleAndRevisionIdAndCommentStatus
func GetNode_commentsByBundleAndRevisionIdAndCommentStatus(offset int, limit int, Bundle_ string, RevisionId_ int, CommentStatus_ int) (*[]*Node_comment, error) {
	var _Node_comment = new([]*Node_comment)
	err := Engine.Table("node__comment").Where("bundle = ? and revision_id = ? and comment_status = ?", Bundle_, RevisionId_, CommentStatus_).Limit(limit, offset).Find(_Node_comment)
	return _Node_comment, err
}

// GetNode_commentsByBundleAndLangcodeAndDelta Get Node_comments via BundleAndLangcodeAndDelta
func GetNode_commentsByBundleAndLangcodeAndDelta(offset int, limit int, Bundle_ string, Langcode_ string, Delta_ int) (*[]*Node_comment, error) {
	var _Node_comment = new([]*Node_comment)
	err := Engine.Table("node__comment").Where("bundle = ? and langcode = ? and delta = ?", Bundle_, Langcode_, Delta_).Limit(limit, offset).Find(_Node_comment)
	return _Node_comment, err
}

// GetNode_commentsByBundleAndLangcodeAndCommentStatus Get Node_comments via BundleAndLangcodeAndCommentStatus
func GetNode_commentsByBundleAndLangcodeAndCommentStatus(offset int, limit int, Bundle_ string, Langcode_ string, CommentStatus_ int) (*[]*Node_comment, error) {
	var _Node_comment = new([]*Node_comment)
	err := Engine.Table("node__comment").Where("bundle = ? and langcode = ? and comment_status = ?", Bundle_, Langcode_, CommentStatus_).Limit(limit, offset).Find(_Node_comment)
	return _Node_comment, err
}

// GetNode_commentsByBundleAndDeltaAndCommentStatus Get Node_comments via BundleAndDeltaAndCommentStatus
func GetNode_commentsByBundleAndDeltaAndCommentStatus(offset int, limit int, Bundle_ string, Delta_ int, CommentStatus_ int) (*[]*Node_comment, error) {
	var _Node_comment = new([]*Node_comment)
	err := Engine.Table("node__comment").Where("bundle = ? and delta = ? and comment_status = ?", Bundle_, Delta_, CommentStatus_).Limit(limit, offset).Find(_Node_comment)
	return _Node_comment, err
}

// GetNode_commentsByDeletedAndEntityIdAndRevisionId Get Node_comments via DeletedAndEntityIdAndRevisionId
func GetNode_commentsByDeletedAndEntityIdAndRevisionId(offset int, limit int, Deleted_ int, EntityId_ int, RevisionId_ int) (*[]*Node_comment, error) {
	var _Node_comment = new([]*Node_comment)
	err := Engine.Table("node__comment").Where("deleted = ? and entity_id = ? and revision_id = ?", Deleted_, EntityId_, RevisionId_).Limit(limit, offset).Find(_Node_comment)
	return _Node_comment, err
}

// GetNode_commentsByDeletedAndEntityIdAndLangcode Get Node_comments via DeletedAndEntityIdAndLangcode
func GetNode_commentsByDeletedAndEntityIdAndLangcode(offset int, limit int, Deleted_ int, EntityId_ int, Langcode_ string) (*[]*Node_comment, error) {
	var _Node_comment = new([]*Node_comment)
	err := Engine.Table("node__comment").Where("deleted = ? and entity_id = ? and langcode = ?", Deleted_, EntityId_, Langcode_).Limit(limit, offset).Find(_Node_comment)
	return _Node_comment, err
}

// GetNode_commentsByDeletedAndEntityIdAndDelta Get Node_comments via DeletedAndEntityIdAndDelta
func GetNode_commentsByDeletedAndEntityIdAndDelta(offset int, limit int, Deleted_ int, EntityId_ int, Delta_ int) (*[]*Node_comment, error) {
	var _Node_comment = new([]*Node_comment)
	err := Engine.Table("node__comment").Where("deleted = ? and entity_id = ? and delta = ?", Deleted_, EntityId_, Delta_).Limit(limit, offset).Find(_Node_comment)
	return _Node_comment, err
}

// GetNode_commentsByDeletedAndEntityIdAndCommentStatus Get Node_comments via DeletedAndEntityIdAndCommentStatus
func GetNode_commentsByDeletedAndEntityIdAndCommentStatus(offset int, limit int, Deleted_ int, EntityId_ int, CommentStatus_ int) (*[]*Node_comment, error) {
	var _Node_comment = new([]*Node_comment)
	err := Engine.Table("node__comment").Where("deleted = ? and entity_id = ? and comment_status = ?", Deleted_, EntityId_, CommentStatus_).Limit(limit, offset).Find(_Node_comment)
	return _Node_comment, err
}

// GetNode_commentsByDeletedAndRevisionIdAndLangcode Get Node_comments via DeletedAndRevisionIdAndLangcode
func GetNode_commentsByDeletedAndRevisionIdAndLangcode(offset int, limit int, Deleted_ int, RevisionId_ int, Langcode_ string) (*[]*Node_comment, error) {
	var _Node_comment = new([]*Node_comment)
	err := Engine.Table("node__comment").Where("deleted = ? and revision_id = ? and langcode = ?", Deleted_, RevisionId_, Langcode_).Limit(limit, offset).Find(_Node_comment)
	return _Node_comment, err
}

// GetNode_commentsByDeletedAndRevisionIdAndDelta Get Node_comments via DeletedAndRevisionIdAndDelta
func GetNode_commentsByDeletedAndRevisionIdAndDelta(offset int, limit int, Deleted_ int, RevisionId_ int, Delta_ int) (*[]*Node_comment, error) {
	var _Node_comment = new([]*Node_comment)
	err := Engine.Table("node__comment").Where("deleted = ? and revision_id = ? and delta = ?", Deleted_, RevisionId_, Delta_).Limit(limit, offset).Find(_Node_comment)
	return _Node_comment, err
}

// GetNode_commentsByDeletedAndRevisionIdAndCommentStatus Get Node_comments via DeletedAndRevisionIdAndCommentStatus
func GetNode_commentsByDeletedAndRevisionIdAndCommentStatus(offset int, limit int, Deleted_ int, RevisionId_ int, CommentStatus_ int) (*[]*Node_comment, error) {
	var _Node_comment = new([]*Node_comment)
	err := Engine.Table("node__comment").Where("deleted = ? and revision_id = ? and comment_status = ?", Deleted_, RevisionId_, CommentStatus_).Limit(limit, offset).Find(_Node_comment)
	return _Node_comment, err
}

// GetNode_commentsByDeletedAndLangcodeAndDelta Get Node_comments via DeletedAndLangcodeAndDelta
func GetNode_commentsByDeletedAndLangcodeAndDelta(offset int, limit int, Deleted_ int, Langcode_ string, Delta_ int) (*[]*Node_comment, error) {
	var _Node_comment = new([]*Node_comment)
	err := Engine.Table("node__comment").Where("deleted = ? and langcode = ? and delta = ?", Deleted_, Langcode_, Delta_).Limit(limit, offset).Find(_Node_comment)
	return _Node_comment, err
}

// GetNode_commentsByDeletedAndLangcodeAndCommentStatus Get Node_comments via DeletedAndLangcodeAndCommentStatus
func GetNode_commentsByDeletedAndLangcodeAndCommentStatus(offset int, limit int, Deleted_ int, Langcode_ string, CommentStatus_ int) (*[]*Node_comment, error) {
	var _Node_comment = new([]*Node_comment)
	err := Engine.Table("node__comment").Where("deleted = ? and langcode = ? and comment_status = ?", Deleted_, Langcode_, CommentStatus_).Limit(limit, offset).Find(_Node_comment)
	return _Node_comment, err
}

// GetNode_commentsByDeletedAndDeltaAndCommentStatus Get Node_comments via DeletedAndDeltaAndCommentStatus
func GetNode_commentsByDeletedAndDeltaAndCommentStatus(offset int, limit int, Deleted_ int, Delta_ int, CommentStatus_ int) (*[]*Node_comment, error) {
	var _Node_comment = new([]*Node_comment)
	err := Engine.Table("node__comment").Where("deleted = ? and delta = ? and comment_status = ?", Deleted_, Delta_, CommentStatus_).Limit(limit, offset).Find(_Node_comment)
	return _Node_comment, err
}

// GetNode_commentsByEntityIdAndRevisionIdAndLangcode Get Node_comments via EntityIdAndRevisionIdAndLangcode
func GetNode_commentsByEntityIdAndRevisionIdAndLangcode(offset int, limit int, EntityId_ int, RevisionId_ int, Langcode_ string) (*[]*Node_comment, error) {
	var _Node_comment = new([]*Node_comment)
	err := Engine.Table("node__comment").Where("entity_id = ? and revision_id = ? and langcode = ?", EntityId_, RevisionId_, Langcode_).Limit(limit, offset).Find(_Node_comment)
	return _Node_comment, err
}

// GetNode_commentsByEntityIdAndRevisionIdAndDelta Get Node_comments via EntityIdAndRevisionIdAndDelta
func GetNode_commentsByEntityIdAndRevisionIdAndDelta(offset int, limit int, EntityId_ int, RevisionId_ int, Delta_ int) (*[]*Node_comment, error) {
	var _Node_comment = new([]*Node_comment)
	err := Engine.Table("node__comment").Where("entity_id = ? and revision_id = ? and delta = ?", EntityId_, RevisionId_, Delta_).Limit(limit, offset).Find(_Node_comment)
	return _Node_comment, err
}

// GetNode_commentsByEntityIdAndRevisionIdAndCommentStatus Get Node_comments via EntityIdAndRevisionIdAndCommentStatus
func GetNode_commentsByEntityIdAndRevisionIdAndCommentStatus(offset int, limit int, EntityId_ int, RevisionId_ int, CommentStatus_ int) (*[]*Node_comment, error) {
	var _Node_comment = new([]*Node_comment)
	err := Engine.Table("node__comment").Where("entity_id = ? and revision_id = ? and comment_status = ?", EntityId_, RevisionId_, CommentStatus_).Limit(limit, offset).Find(_Node_comment)
	return _Node_comment, err
}

// GetNode_commentsByEntityIdAndLangcodeAndDelta Get Node_comments via EntityIdAndLangcodeAndDelta
func GetNode_commentsByEntityIdAndLangcodeAndDelta(offset int, limit int, EntityId_ int, Langcode_ string, Delta_ int) (*[]*Node_comment, error) {
	var _Node_comment = new([]*Node_comment)
	err := Engine.Table("node__comment").Where("entity_id = ? and langcode = ? and delta = ?", EntityId_, Langcode_, Delta_).Limit(limit, offset).Find(_Node_comment)
	return _Node_comment, err
}

// GetNode_commentsByEntityIdAndLangcodeAndCommentStatus Get Node_comments via EntityIdAndLangcodeAndCommentStatus
func GetNode_commentsByEntityIdAndLangcodeAndCommentStatus(offset int, limit int, EntityId_ int, Langcode_ string, CommentStatus_ int) (*[]*Node_comment, error) {
	var _Node_comment = new([]*Node_comment)
	err := Engine.Table("node__comment").Where("entity_id = ? and langcode = ? and comment_status = ?", EntityId_, Langcode_, CommentStatus_).Limit(limit, offset).Find(_Node_comment)
	return _Node_comment, err
}

// GetNode_commentsByEntityIdAndDeltaAndCommentStatus Get Node_comments via EntityIdAndDeltaAndCommentStatus
func GetNode_commentsByEntityIdAndDeltaAndCommentStatus(offset int, limit int, EntityId_ int, Delta_ int, CommentStatus_ int) (*[]*Node_comment, error) {
	var _Node_comment = new([]*Node_comment)
	err := Engine.Table("node__comment").Where("entity_id = ? and delta = ? and comment_status = ?", EntityId_, Delta_, CommentStatus_).Limit(limit, offset).Find(_Node_comment)
	return _Node_comment, err
}

// GetNode_commentsByRevisionIdAndLangcodeAndDelta Get Node_comments via RevisionIdAndLangcodeAndDelta
func GetNode_commentsByRevisionIdAndLangcodeAndDelta(offset int, limit int, RevisionId_ int, Langcode_ string, Delta_ int) (*[]*Node_comment, error) {
	var _Node_comment = new([]*Node_comment)
	err := Engine.Table("node__comment").Where("revision_id = ? and langcode = ? and delta = ?", RevisionId_, Langcode_, Delta_).Limit(limit, offset).Find(_Node_comment)
	return _Node_comment, err
}

// GetNode_commentsByRevisionIdAndLangcodeAndCommentStatus Get Node_comments via RevisionIdAndLangcodeAndCommentStatus
func GetNode_commentsByRevisionIdAndLangcodeAndCommentStatus(offset int, limit int, RevisionId_ int, Langcode_ string, CommentStatus_ int) (*[]*Node_comment, error) {
	var _Node_comment = new([]*Node_comment)
	err := Engine.Table("node__comment").Where("revision_id = ? and langcode = ? and comment_status = ?", RevisionId_, Langcode_, CommentStatus_).Limit(limit, offset).Find(_Node_comment)
	return _Node_comment, err
}

// GetNode_commentsByRevisionIdAndDeltaAndCommentStatus Get Node_comments via RevisionIdAndDeltaAndCommentStatus
func GetNode_commentsByRevisionIdAndDeltaAndCommentStatus(offset int, limit int, RevisionId_ int, Delta_ int, CommentStatus_ int) (*[]*Node_comment, error) {
	var _Node_comment = new([]*Node_comment)
	err := Engine.Table("node__comment").Where("revision_id = ? and delta = ? and comment_status = ?", RevisionId_, Delta_, CommentStatus_).Limit(limit, offset).Find(_Node_comment)
	return _Node_comment, err
}

// GetNode_commentsByLangcodeAndDeltaAndCommentStatus Get Node_comments via LangcodeAndDeltaAndCommentStatus
func GetNode_commentsByLangcodeAndDeltaAndCommentStatus(offset int, limit int, Langcode_ string, Delta_ int, CommentStatus_ int) (*[]*Node_comment, error) {
	var _Node_comment = new([]*Node_comment)
	err := Engine.Table("node__comment").Where("langcode = ? and delta = ? and comment_status = ?", Langcode_, Delta_, CommentStatus_).Limit(limit, offset).Find(_Node_comment)
	return _Node_comment, err
}

// GetNode_commentsByBundleAndDeleted Get Node_comments via BundleAndDeleted
func GetNode_commentsByBundleAndDeleted(offset int, limit int, Bundle_ string, Deleted_ int) (*[]*Node_comment, error) {
	var _Node_comment = new([]*Node_comment)
	err := Engine.Table("node__comment").Where("bundle = ? and deleted = ?", Bundle_, Deleted_).Limit(limit, offset).Find(_Node_comment)
	return _Node_comment, err
}

// GetNode_commentsByBundleAndEntityId Get Node_comments via BundleAndEntityId
func GetNode_commentsByBundleAndEntityId(offset int, limit int, Bundle_ string, EntityId_ int) (*[]*Node_comment, error) {
	var _Node_comment = new([]*Node_comment)
	err := Engine.Table("node__comment").Where("bundle = ? and entity_id = ?", Bundle_, EntityId_).Limit(limit, offset).Find(_Node_comment)
	return _Node_comment, err
}

// GetNode_commentsByBundleAndRevisionId Get Node_comments via BundleAndRevisionId
func GetNode_commentsByBundleAndRevisionId(offset int, limit int, Bundle_ string, RevisionId_ int) (*[]*Node_comment, error) {
	var _Node_comment = new([]*Node_comment)
	err := Engine.Table("node__comment").Where("bundle = ? and revision_id = ?", Bundle_, RevisionId_).Limit(limit, offset).Find(_Node_comment)
	return _Node_comment, err
}

// GetNode_commentsByBundleAndLangcode Get Node_comments via BundleAndLangcode
func GetNode_commentsByBundleAndLangcode(offset int, limit int, Bundle_ string, Langcode_ string) (*[]*Node_comment, error) {
	var _Node_comment = new([]*Node_comment)
	err := Engine.Table("node__comment").Where("bundle = ? and langcode = ?", Bundle_, Langcode_).Limit(limit, offset).Find(_Node_comment)
	return _Node_comment, err
}

// GetNode_commentsByBundleAndDelta Get Node_comments via BundleAndDelta
func GetNode_commentsByBundleAndDelta(offset int, limit int, Bundle_ string, Delta_ int) (*[]*Node_comment, error) {
	var _Node_comment = new([]*Node_comment)
	err := Engine.Table("node__comment").Where("bundle = ? and delta = ?", Bundle_, Delta_).Limit(limit, offset).Find(_Node_comment)
	return _Node_comment, err
}

// GetNode_commentsByBundleAndCommentStatus Get Node_comments via BundleAndCommentStatus
func GetNode_commentsByBundleAndCommentStatus(offset int, limit int, Bundle_ string, CommentStatus_ int) (*[]*Node_comment, error) {
	var _Node_comment = new([]*Node_comment)
	err := Engine.Table("node__comment").Where("bundle = ? and comment_status = ?", Bundle_, CommentStatus_).Limit(limit, offset).Find(_Node_comment)
	return _Node_comment, err
}

// GetNode_commentsByDeletedAndEntityId Get Node_comments via DeletedAndEntityId
func GetNode_commentsByDeletedAndEntityId(offset int, limit int, Deleted_ int, EntityId_ int) (*[]*Node_comment, error) {
	var _Node_comment = new([]*Node_comment)
	err := Engine.Table("node__comment").Where("deleted = ? and entity_id = ?", Deleted_, EntityId_).Limit(limit, offset).Find(_Node_comment)
	return _Node_comment, err
}

// GetNode_commentsByDeletedAndRevisionId Get Node_comments via DeletedAndRevisionId
func GetNode_commentsByDeletedAndRevisionId(offset int, limit int, Deleted_ int, RevisionId_ int) (*[]*Node_comment, error) {
	var _Node_comment = new([]*Node_comment)
	err := Engine.Table("node__comment").Where("deleted = ? and revision_id = ?", Deleted_, RevisionId_).Limit(limit, offset).Find(_Node_comment)
	return _Node_comment, err
}

// GetNode_commentsByDeletedAndLangcode Get Node_comments via DeletedAndLangcode
func GetNode_commentsByDeletedAndLangcode(offset int, limit int, Deleted_ int, Langcode_ string) (*[]*Node_comment, error) {
	var _Node_comment = new([]*Node_comment)
	err := Engine.Table("node__comment").Where("deleted = ? and langcode = ?", Deleted_, Langcode_).Limit(limit, offset).Find(_Node_comment)
	return _Node_comment, err
}

// GetNode_commentsByDeletedAndDelta Get Node_comments via DeletedAndDelta
func GetNode_commentsByDeletedAndDelta(offset int, limit int, Deleted_ int, Delta_ int) (*[]*Node_comment, error) {
	var _Node_comment = new([]*Node_comment)
	err := Engine.Table("node__comment").Where("deleted = ? and delta = ?", Deleted_, Delta_).Limit(limit, offset).Find(_Node_comment)
	return _Node_comment, err
}

// GetNode_commentsByDeletedAndCommentStatus Get Node_comments via DeletedAndCommentStatus
func GetNode_commentsByDeletedAndCommentStatus(offset int, limit int, Deleted_ int, CommentStatus_ int) (*[]*Node_comment, error) {
	var _Node_comment = new([]*Node_comment)
	err := Engine.Table("node__comment").Where("deleted = ? and comment_status = ?", Deleted_, CommentStatus_).Limit(limit, offset).Find(_Node_comment)
	return _Node_comment, err
}

// GetNode_commentsByEntityIdAndRevisionId Get Node_comments via EntityIdAndRevisionId
func GetNode_commentsByEntityIdAndRevisionId(offset int, limit int, EntityId_ int, RevisionId_ int) (*[]*Node_comment, error) {
	var _Node_comment = new([]*Node_comment)
	err := Engine.Table("node__comment").Where("entity_id = ? and revision_id = ?", EntityId_, RevisionId_).Limit(limit, offset).Find(_Node_comment)
	return _Node_comment, err
}

// GetNode_commentsByEntityIdAndLangcode Get Node_comments via EntityIdAndLangcode
func GetNode_commentsByEntityIdAndLangcode(offset int, limit int, EntityId_ int, Langcode_ string) (*[]*Node_comment, error) {
	var _Node_comment = new([]*Node_comment)
	err := Engine.Table("node__comment").Where("entity_id = ? and langcode = ?", EntityId_, Langcode_).Limit(limit, offset).Find(_Node_comment)
	return _Node_comment, err
}

// GetNode_commentsByEntityIdAndDelta Get Node_comments via EntityIdAndDelta
func GetNode_commentsByEntityIdAndDelta(offset int, limit int, EntityId_ int, Delta_ int) (*[]*Node_comment, error) {
	var _Node_comment = new([]*Node_comment)
	err := Engine.Table("node__comment").Where("entity_id = ? and delta = ?", EntityId_, Delta_).Limit(limit, offset).Find(_Node_comment)
	return _Node_comment, err
}

// GetNode_commentsByEntityIdAndCommentStatus Get Node_comments via EntityIdAndCommentStatus
func GetNode_commentsByEntityIdAndCommentStatus(offset int, limit int, EntityId_ int, CommentStatus_ int) (*[]*Node_comment, error) {
	var _Node_comment = new([]*Node_comment)
	err := Engine.Table("node__comment").Where("entity_id = ? and comment_status = ?", EntityId_, CommentStatus_).Limit(limit, offset).Find(_Node_comment)
	return _Node_comment, err
}

// GetNode_commentsByRevisionIdAndLangcode Get Node_comments via RevisionIdAndLangcode
func GetNode_commentsByRevisionIdAndLangcode(offset int, limit int, RevisionId_ int, Langcode_ string) (*[]*Node_comment, error) {
	var _Node_comment = new([]*Node_comment)
	err := Engine.Table("node__comment").Where("revision_id = ? and langcode = ?", RevisionId_, Langcode_).Limit(limit, offset).Find(_Node_comment)
	return _Node_comment, err
}

// GetNode_commentsByRevisionIdAndDelta Get Node_comments via RevisionIdAndDelta
func GetNode_commentsByRevisionIdAndDelta(offset int, limit int, RevisionId_ int, Delta_ int) (*[]*Node_comment, error) {
	var _Node_comment = new([]*Node_comment)
	err := Engine.Table("node__comment").Where("revision_id = ? and delta = ?", RevisionId_, Delta_).Limit(limit, offset).Find(_Node_comment)
	return _Node_comment, err
}

// GetNode_commentsByRevisionIdAndCommentStatus Get Node_comments via RevisionIdAndCommentStatus
func GetNode_commentsByRevisionIdAndCommentStatus(offset int, limit int, RevisionId_ int, CommentStatus_ int) (*[]*Node_comment, error) {
	var _Node_comment = new([]*Node_comment)
	err := Engine.Table("node__comment").Where("revision_id = ? and comment_status = ?", RevisionId_, CommentStatus_).Limit(limit, offset).Find(_Node_comment)
	return _Node_comment, err
}

// GetNode_commentsByLangcodeAndDelta Get Node_comments via LangcodeAndDelta
func GetNode_commentsByLangcodeAndDelta(offset int, limit int, Langcode_ string, Delta_ int) (*[]*Node_comment, error) {
	var _Node_comment = new([]*Node_comment)
	err := Engine.Table("node__comment").Where("langcode = ? and delta = ?", Langcode_, Delta_).Limit(limit, offset).Find(_Node_comment)
	return _Node_comment, err
}

// GetNode_commentsByLangcodeAndCommentStatus Get Node_comments via LangcodeAndCommentStatus
func GetNode_commentsByLangcodeAndCommentStatus(offset int, limit int, Langcode_ string, CommentStatus_ int) (*[]*Node_comment, error) {
	var _Node_comment = new([]*Node_comment)
	err := Engine.Table("node__comment").Where("langcode = ? and comment_status = ?", Langcode_, CommentStatus_).Limit(limit, offset).Find(_Node_comment)
	return _Node_comment, err
}

// GetNode_commentsByDeltaAndCommentStatus Get Node_comments via DeltaAndCommentStatus
func GetNode_commentsByDeltaAndCommentStatus(offset int, limit int, Delta_ int, CommentStatus_ int) (*[]*Node_comment, error) {
	var _Node_comment = new([]*Node_comment)
	err := Engine.Table("node__comment").Where("delta = ? and comment_status = ?", Delta_, CommentStatus_).Limit(limit, offset).Find(_Node_comment)
	return _Node_comment, err
}

// GetNode_comments Get Node_comments via field
func GetNode_comments(offset int, limit int, field string) (*[]*Node_comment, error) {
	var _Node_comment = new([]*Node_comment)
	err := Engine.Table("node__comment").Limit(limit, offset).Desc(field).Find(_Node_comment)
	return _Node_comment, err
}

// GetNode_commentsViaBundle Get Node_comments via Bundle
func GetNode_commentsViaBundle(offset int, limit int, Bundle_ string, field string) (*[]*Node_comment, error) {
	var _Node_comment = new([]*Node_comment)
	err := Engine.Table("node__comment").Where("bundle = ?", Bundle_).Limit(limit, offset).Desc(field).Find(_Node_comment)
	return _Node_comment, err
}

// GetNode_commentsViaDeleted Get Node_comments via Deleted
func GetNode_commentsViaDeleted(offset int, limit int, Deleted_ int, field string) (*[]*Node_comment, error) {
	var _Node_comment = new([]*Node_comment)
	err := Engine.Table("node__comment").Where("deleted = ?", Deleted_).Limit(limit, offset).Desc(field).Find(_Node_comment)
	return _Node_comment, err
}

// GetNode_commentsViaEntityId Get Node_comments via EntityId
func GetNode_commentsViaEntityId(offset int, limit int, EntityId_ int, field string) (*[]*Node_comment, error) {
	var _Node_comment = new([]*Node_comment)
	err := Engine.Table("node__comment").Where("entity_id = ?", EntityId_).Limit(limit, offset).Desc(field).Find(_Node_comment)
	return _Node_comment, err
}

// GetNode_commentsViaRevisionId Get Node_comments via RevisionId
func GetNode_commentsViaRevisionId(offset int, limit int, RevisionId_ int, field string) (*[]*Node_comment, error) {
	var _Node_comment = new([]*Node_comment)
	err := Engine.Table("node__comment").Where("revision_id = ?", RevisionId_).Limit(limit, offset).Desc(field).Find(_Node_comment)
	return _Node_comment, err
}

// GetNode_commentsViaLangcode Get Node_comments via Langcode
func GetNode_commentsViaLangcode(offset int, limit int, Langcode_ string, field string) (*[]*Node_comment, error) {
	var _Node_comment = new([]*Node_comment)
	err := Engine.Table("node__comment").Where("langcode = ?", Langcode_).Limit(limit, offset).Desc(field).Find(_Node_comment)
	return _Node_comment, err
}

// GetNode_commentsViaDelta Get Node_comments via Delta
func GetNode_commentsViaDelta(offset int, limit int, Delta_ int, field string) (*[]*Node_comment, error) {
	var _Node_comment = new([]*Node_comment)
	err := Engine.Table("node__comment").Where("delta = ?", Delta_).Limit(limit, offset).Desc(field).Find(_Node_comment)
	return _Node_comment, err
}

// GetNode_commentsViaCommentStatus Get Node_comments via CommentStatus
func GetNode_commentsViaCommentStatus(offset int, limit int, CommentStatus_ int, field string) (*[]*Node_comment, error) {
	var _Node_comment = new([]*Node_comment)
	err := Engine.Table("node__comment").Where("comment_status = ?", CommentStatus_).Limit(limit, offset).Desc(field).Find(_Node_comment)
	return _Node_comment, err
}

// HasNode_commentViaBundle Has Node_comment via Bundle
func HasNode_commentViaBundle(iBundle string) bool {
	if has, err := Engine.Where("bundle = ?", iBundle).Get(new(Node_comment)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasNode_commentViaDeleted Has Node_comment via Deleted
func HasNode_commentViaDeleted(iDeleted int) bool {
	if has, err := Engine.Where("deleted = ?", iDeleted).Get(new(Node_comment)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasNode_commentViaEntityId Has Node_comment via EntityId
func HasNode_commentViaEntityId(iEntityId int) bool {
	if has, err := Engine.Where("entity_id = ?", iEntityId).Get(new(Node_comment)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasNode_commentViaRevisionId Has Node_comment via RevisionId
func HasNode_commentViaRevisionId(iRevisionId int) bool {
	if has, err := Engine.Where("revision_id = ?", iRevisionId).Get(new(Node_comment)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasNode_commentViaLangcode Has Node_comment via Langcode
func HasNode_commentViaLangcode(iLangcode string) bool {
	if has, err := Engine.Where("langcode = ?", iLangcode).Get(new(Node_comment)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasNode_commentViaDelta Has Node_comment via Delta
func HasNode_commentViaDelta(iDelta int) bool {
	if has, err := Engine.Where("delta = ?", iDelta).Get(new(Node_comment)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasNode_commentViaCommentStatus Has Node_comment via CommentStatus
func HasNode_commentViaCommentStatus(iCommentStatus int) bool {
	if has, err := Engine.Where("comment_status = ?", iCommentStatus).Get(new(Node_comment)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// GetNode_commentViaBundle Get Node_comment via Bundle
func GetNode_commentViaBundle(iBundle string) (*Node_comment, error) {
	var _Node_comment = &Node_comment{Bundle: iBundle}
	has, err := Engine.Get(_Node_comment)
	if has {
		return _Node_comment, err
	} else {
		return nil, err
	}
}

// GetNode_commentViaDeleted Get Node_comment via Deleted
func GetNode_commentViaDeleted(iDeleted int) (*Node_comment, error) {
	var _Node_comment = &Node_comment{Deleted: iDeleted}
	has, err := Engine.Get(_Node_comment)
	if has {
		return _Node_comment, err
	} else {
		return nil, err
	}
}

// GetNode_commentViaEntityId Get Node_comment via EntityId
func GetNode_commentViaEntityId(iEntityId int) (*Node_comment, error) {
	var _Node_comment = &Node_comment{EntityId: iEntityId}
	has, err := Engine.Get(_Node_comment)
	if has {
		return _Node_comment, err
	} else {
		return nil, err
	}
}

// GetNode_commentViaRevisionId Get Node_comment via RevisionId
func GetNode_commentViaRevisionId(iRevisionId int) (*Node_comment, error) {
	var _Node_comment = &Node_comment{RevisionId: iRevisionId}
	has, err := Engine.Get(_Node_comment)
	if has {
		return _Node_comment, err
	} else {
		return nil, err
	}
}

// GetNode_commentViaLangcode Get Node_comment via Langcode
func GetNode_commentViaLangcode(iLangcode string) (*Node_comment, error) {
	var _Node_comment = &Node_comment{Langcode: iLangcode}
	has, err := Engine.Get(_Node_comment)
	if has {
		return _Node_comment, err
	} else {
		return nil, err
	}
}

// GetNode_commentViaDelta Get Node_comment via Delta
func GetNode_commentViaDelta(iDelta int) (*Node_comment, error) {
	var _Node_comment = &Node_comment{Delta: iDelta}
	has, err := Engine.Get(_Node_comment)
	if has {
		return _Node_comment, err
	} else {
		return nil, err
	}
}

// GetNode_commentViaCommentStatus Get Node_comment via CommentStatus
func GetNode_commentViaCommentStatus(iCommentStatus int) (*Node_comment, error) {
	var _Node_comment = &Node_comment{CommentStatus: iCommentStatus}
	has, err := Engine.Get(_Node_comment)
	if has {
		return _Node_comment, err
	} else {
		return nil, err
	}
}

// SetNode_commentViaBundle Set Node_comment via Bundle
func SetNode_commentViaBundle(iBundle string, node__comment *Node_comment) (int64, error) {
	node__comment.Bundle = iBundle
	return Engine.Insert(node__comment)
}

// SetNode_commentViaDeleted Set Node_comment via Deleted
func SetNode_commentViaDeleted(iDeleted int, node__comment *Node_comment) (int64, error) {
	node__comment.Deleted = iDeleted
	return Engine.Insert(node__comment)
}

// SetNode_commentViaEntityId Set Node_comment via EntityId
func SetNode_commentViaEntityId(iEntityId int, node__comment *Node_comment) (int64, error) {
	node__comment.EntityId = iEntityId
	return Engine.Insert(node__comment)
}

// SetNode_commentViaRevisionId Set Node_comment via RevisionId
func SetNode_commentViaRevisionId(iRevisionId int, node__comment *Node_comment) (int64, error) {
	node__comment.RevisionId = iRevisionId
	return Engine.Insert(node__comment)
}

// SetNode_commentViaLangcode Set Node_comment via Langcode
func SetNode_commentViaLangcode(iLangcode string, node__comment *Node_comment) (int64, error) {
	node__comment.Langcode = iLangcode
	return Engine.Insert(node__comment)
}

// SetNode_commentViaDelta Set Node_comment via Delta
func SetNode_commentViaDelta(iDelta int, node__comment *Node_comment) (int64, error) {
	node__comment.Delta = iDelta
	return Engine.Insert(node__comment)
}

// SetNode_commentViaCommentStatus Set Node_comment via CommentStatus
func SetNode_commentViaCommentStatus(iCommentStatus int, node__comment *Node_comment) (int64, error) {
	node__comment.CommentStatus = iCommentStatus
	return Engine.Insert(node__comment)
}

// AddNode_comment Add Node_comment via all columns
func AddNode_comment(iBundle string, iDeleted int, iEntityId int, iRevisionId int, iLangcode string, iDelta int, iCommentStatus int) error {
	_Node_comment := &Node_comment{Bundle: iBundle, Deleted: iDeleted, EntityId: iEntityId, RevisionId: iRevisionId, Langcode: iLangcode, Delta: iDelta, CommentStatus: iCommentStatus}
	if _, err := Engine.Insert(_Node_comment); err != nil {
		return err
	} else {
		return nil
	}
}

// PostNode_comment Post Node_comment via iNode_comment
func PostNode_comment(iNode_comment *Node_comment) (string, error) {
	_, err := Engine.Insert(iNode_comment)
	return iNode_comment.Bundle, err
}

// PutNode_comment Put Node_comment
func PutNode_comment(iNode_comment *Node_comment) (string, error) {
	_, err := Engine.Id(iNode_comment.Bundle).Update(iNode_comment)
	return iNode_comment.Bundle, err
}

// PutNode_commentViaBundle Put Node_comment via Bundle
func PutNode_commentViaBundle(Bundle_ string, iNode_comment *Node_comment) (int64, error) {
	row, err := Engine.Update(iNode_comment, &Node_comment{Bundle: Bundle_})
	return row, err
}

// PutNode_commentViaDeleted Put Node_comment via Deleted
func PutNode_commentViaDeleted(Deleted_ int, iNode_comment *Node_comment) (int64, error) {
	row, err := Engine.Update(iNode_comment, &Node_comment{Deleted: Deleted_})
	return row, err
}

// PutNode_commentViaEntityId Put Node_comment via EntityId
func PutNode_commentViaEntityId(EntityId_ int, iNode_comment *Node_comment) (int64, error) {
	row, err := Engine.Update(iNode_comment, &Node_comment{EntityId: EntityId_})
	return row, err
}

// PutNode_commentViaRevisionId Put Node_comment via RevisionId
func PutNode_commentViaRevisionId(RevisionId_ int, iNode_comment *Node_comment) (int64, error) {
	row, err := Engine.Update(iNode_comment, &Node_comment{RevisionId: RevisionId_})
	return row, err
}

// PutNode_commentViaLangcode Put Node_comment via Langcode
func PutNode_commentViaLangcode(Langcode_ string, iNode_comment *Node_comment) (int64, error) {
	row, err := Engine.Update(iNode_comment, &Node_comment{Langcode: Langcode_})
	return row, err
}

// PutNode_commentViaDelta Put Node_comment via Delta
func PutNode_commentViaDelta(Delta_ int, iNode_comment *Node_comment) (int64, error) {
	row, err := Engine.Update(iNode_comment, &Node_comment{Delta: Delta_})
	return row, err
}

// PutNode_commentViaCommentStatus Put Node_comment via CommentStatus
func PutNode_commentViaCommentStatus(CommentStatus_ int, iNode_comment *Node_comment) (int64, error) {
	row, err := Engine.Update(iNode_comment, &Node_comment{CommentStatus: CommentStatus_})
	return row, err
}

// UpdateNode_commentViaBundle via map[string]interface{}{}
func UpdateNode_commentViaBundle(iBundle string, iNode_commentMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Node_comment)).Where("bundle = ?", iBundle).Update(iNode_commentMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateNode_commentViaDeleted via map[string]interface{}{}
func UpdateNode_commentViaDeleted(iDeleted int, iNode_commentMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Node_comment)).Where("deleted = ?", iDeleted).Update(iNode_commentMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateNode_commentViaEntityId via map[string]interface{}{}
func UpdateNode_commentViaEntityId(iEntityId int, iNode_commentMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Node_comment)).Where("entity_id = ?", iEntityId).Update(iNode_commentMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateNode_commentViaRevisionId via map[string]interface{}{}
func UpdateNode_commentViaRevisionId(iRevisionId int, iNode_commentMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Node_comment)).Where("revision_id = ?", iRevisionId).Update(iNode_commentMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateNode_commentViaLangcode via map[string]interface{}{}
func UpdateNode_commentViaLangcode(iLangcode string, iNode_commentMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Node_comment)).Where("langcode = ?", iLangcode).Update(iNode_commentMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateNode_commentViaDelta via map[string]interface{}{}
func UpdateNode_commentViaDelta(iDelta int, iNode_commentMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Node_comment)).Where("delta = ?", iDelta).Update(iNode_commentMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateNode_commentViaCommentStatus via map[string]interface{}{}
func UpdateNode_commentViaCommentStatus(iCommentStatus int, iNode_commentMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Node_comment)).Where("comment_status = ?", iCommentStatus).Update(iNode_commentMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// DeleteNode_comment Delete Node_comment
func DeleteNode_comment(iBundle string) (int64, error) {
	row, err := Engine.Id(iBundle).Delete(new(Node_comment))
	return row, err
}

// DeleteNode_commentViaBundle Delete Node_comment via Bundle
func DeleteNode_commentViaBundle(iBundle string) (err error) {
	var has bool
	var _Node_comment = &Node_comment{Bundle: iBundle}
	if has, err = Engine.Get(_Node_comment); (has == true) && (err == nil) {
		if row, err := Engine.Where("bundle = ?", iBundle).Delete(new(Node_comment)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteNode_commentViaDeleted Delete Node_comment via Deleted
func DeleteNode_commentViaDeleted(iDeleted int) (err error) {
	var has bool
	var _Node_comment = &Node_comment{Deleted: iDeleted}
	if has, err = Engine.Get(_Node_comment); (has == true) && (err == nil) {
		if row, err := Engine.Where("deleted = ?", iDeleted).Delete(new(Node_comment)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteNode_commentViaEntityId Delete Node_comment via EntityId
func DeleteNode_commentViaEntityId(iEntityId int) (err error) {
	var has bool
	var _Node_comment = &Node_comment{EntityId: iEntityId}
	if has, err = Engine.Get(_Node_comment); (has == true) && (err == nil) {
		if row, err := Engine.Where("entity_id = ?", iEntityId).Delete(new(Node_comment)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteNode_commentViaRevisionId Delete Node_comment via RevisionId
func DeleteNode_commentViaRevisionId(iRevisionId int) (err error) {
	var has bool
	var _Node_comment = &Node_comment{RevisionId: iRevisionId}
	if has, err = Engine.Get(_Node_comment); (has == true) && (err == nil) {
		if row, err := Engine.Where("revision_id = ?", iRevisionId).Delete(new(Node_comment)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteNode_commentViaLangcode Delete Node_comment via Langcode
func DeleteNode_commentViaLangcode(iLangcode string) (err error) {
	var has bool
	var _Node_comment = &Node_comment{Langcode: iLangcode}
	if has, err = Engine.Get(_Node_comment); (has == true) && (err == nil) {
		if row, err := Engine.Where("langcode = ?", iLangcode).Delete(new(Node_comment)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteNode_commentViaDelta Delete Node_comment via Delta
func DeleteNode_commentViaDelta(iDelta int) (err error) {
	var has bool
	var _Node_comment = &Node_comment{Delta: iDelta}
	if has, err = Engine.Get(_Node_comment); (has == true) && (err == nil) {
		if row, err := Engine.Where("delta = ?", iDelta).Delete(new(Node_comment)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteNode_commentViaCommentStatus Delete Node_comment via CommentStatus
func DeleteNode_commentViaCommentStatus(iCommentStatus int) (err error) {
	var has bool
	var _Node_comment = &Node_comment{CommentStatus: iCommentStatus}
	if has, err = Engine.Get(_Node_comment); (has == true) && (err == nil) {
		if row, err := Engine.Where("comment_status = ?", iCommentStatus).Delete(new(Node_comment)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}
