package model

type NodeRevision_comment struct {
	Bundle        string `xorm:"not null default '' index VARCHAR(128)"`
	Deleted       int    `xorm:"not null pk default 0 TINYINT(4)"`
	EntityId      int    `xorm:"not null pk INT(10)"`
	RevisionId    int    `xorm:"not null pk index INT(10)"`
	Langcode      string `xorm:"not null pk default '' VARCHAR(32)"`
	Delta         int    `xorm:"not null pk INT(10)"`
	CommentStatus int    `xorm:"not null default 0 INT(11)"`
}

// GetNodeRevision_commentsCount NodeRevision_comments' Count
func GetNodeRevision_commentsCount(offset int, limit int) (int64, error) {
	total, err := Engine.Limit(limit, offset).Count(&NodeRevision_comment{})
	return total, err
}

// GetNodeRevision_commentCountViaBundle Get NodeRevision_comment via Bundle
func GetNodeRevision_commentCountViaBundle(iBundle string) int64 {
	n, _ := Engine.Where("bundle = ?", iBundle).Count(&NodeRevision_comment{Bundle: iBundle})
	return n
}

// GetNodeRevision_commentCountViaDeleted Get NodeRevision_comment via Deleted
func GetNodeRevision_commentCountViaDeleted(iDeleted int) int64 {
	n, _ := Engine.Where("deleted = ?", iDeleted).Count(&NodeRevision_comment{Deleted: iDeleted})
	return n
}

// GetNodeRevision_commentCountViaEntityId Get NodeRevision_comment via EntityId
func GetNodeRevision_commentCountViaEntityId(iEntityId int) int64 {
	n, _ := Engine.Where("entity_id = ?", iEntityId).Count(&NodeRevision_comment{EntityId: iEntityId})
	return n
}

// GetNodeRevision_commentCountViaRevisionId Get NodeRevision_comment via RevisionId
func GetNodeRevision_commentCountViaRevisionId(iRevisionId int) int64 {
	n, _ := Engine.Where("revision_id = ?", iRevisionId).Count(&NodeRevision_comment{RevisionId: iRevisionId})
	return n
}

// GetNodeRevision_commentCountViaLangcode Get NodeRevision_comment via Langcode
func GetNodeRevision_commentCountViaLangcode(iLangcode string) int64 {
	n, _ := Engine.Where("langcode = ?", iLangcode).Count(&NodeRevision_comment{Langcode: iLangcode})
	return n
}

// GetNodeRevision_commentCountViaDelta Get NodeRevision_comment via Delta
func GetNodeRevision_commentCountViaDelta(iDelta int) int64 {
	n, _ := Engine.Where("delta = ?", iDelta).Count(&NodeRevision_comment{Delta: iDelta})
	return n
}

// GetNodeRevision_commentCountViaCommentStatus Get NodeRevision_comment via CommentStatus
func GetNodeRevision_commentCountViaCommentStatus(iCommentStatus int) int64 {
	n, _ := Engine.Where("comment_status = ?", iCommentStatus).Count(&NodeRevision_comment{CommentStatus: iCommentStatus})
	return n
}

// GetNodeRevision_commentsByBundleAndDeletedAndEntityId Get NodeRevision_comments via BundleAndDeletedAndEntityId
func GetNodeRevision_commentsByBundleAndDeletedAndEntityId(offset int, limit int, Bundle_ string, Deleted_ int, EntityId_ int) (*[]*NodeRevision_comment, error) {
	var _NodeRevision_comment = new([]*NodeRevision_comment)
	err := Engine.Table("node_revision__comment").Where("bundle = ? and deleted = ? and entity_id = ?", Bundle_, Deleted_, EntityId_).Limit(limit, offset).Find(_NodeRevision_comment)
	return _NodeRevision_comment, err
}

// GetNodeRevision_commentsByBundleAndDeletedAndRevisionId Get NodeRevision_comments via BundleAndDeletedAndRevisionId
func GetNodeRevision_commentsByBundleAndDeletedAndRevisionId(offset int, limit int, Bundle_ string, Deleted_ int, RevisionId_ int) (*[]*NodeRevision_comment, error) {
	var _NodeRevision_comment = new([]*NodeRevision_comment)
	err := Engine.Table("node_revision__comment").Where("bundle = ? and deleted = ? and revision_id = ?", Bundle_, Deleted_, RevisionId_).Limit(limit, offset).Find(_NodeRevision_comment)
	return _NodeRevision_comment, err
}

// GetNodeRevision_commentsByBundleAndDeletedAndLangcode Get NodeRevision_comments via BundleAndDeletedAndLangcode
func GetNodeRevision_commentsByBundleAndDeletedAndLangcode(offset int, limit int, Bundle_ string, Deleted_ int, Langcode_ string) (*[]*NodeRevision_comment, error) {
	var _NodeRevision_comment = new([]*NodeRevision_comment)
	err := Engine.Table("node_revision__comment").Where("bundle = ? and deleted = ? and langcode = ?", Bundle_, Deleted_, Langcode_).Limit(limit, offset).Find(_NodeRevision_comment)
	return _NodeRevision_comment, err
}

// GetNodeRevision_commentsByBundleAndDeletedAndDelta Get NodeRevision_comments via BundleAndDeletedAndDelta
func GetNodeRevision_commentsByBundleAndDeletedAndDelta(offset int, limit int, Bundle_ string, Deleted_ int, Delta_ int) (*[]*NodeRevision_comment, error) {
	var _NodeRevision_comment = new([]*NodeRevision_comment)
	err := Engine.Table("node_revision__comment").Where("bundle = ? and deleted = ? and delta = ?", Bundle_, Deleted_, Delta_).Limit(limit, offset).Find(_NodeRevision_comment)
	return _NodeRevision_comment, err
}

// GetNodeRevision_commentsByBundleAndDeletedAndCommentStatus Get NodeRevision_comments via BundleAndDeletedAndCommentStatus
func GetNodeRevision_commentsByBundleAndDeletedAndCommentStatus(offset int, limit int, Bundle_ string, Deleted_ int, CommentStatus_ int) (*[]*NodeRevision_comment, error) {
	var _NodeRevision_comment = new([]*NodeRevision_comment)
	err := Engine.Table("node_revision__comment").Where("bundle = ? and deleted = ? and comment_status = ?", Bundle_, Deleted_, CommentStatus_).Limit(limit, offset).Find(_NodeRevision_comment)
	return _NodeRevision_comment, err
}

// GetNodeRevision_commentsByBundleAndEntityIdAndRevisionId Get NodeRevision_comments via BundleAndEntityIdAndRevisionId
func GetNodeRevision_commentsByBundleAndEntityIdAndRevisionId(offset int, limit int, Bundle_ string, EntityId_ int, RevisionId_ int) (*[]*NodeRevision_comment, error) {
	var _NodeRevision_comment = new([]*NodeRevision_comment)
	err := Engine.Table("node_revision__comment").Where("bundle = ? and entity_id = ? and revision_id = ?", Bundle_, EntityId_, RevisionId_).Limit(limit, offset).Find(_NodeRevision_comment)
	return _NodeRevision_comment, err
}

// GetNodeRevision_commentsByBundleAndEntityIdAndLangcode Get NodeRevision_comments via BundleAndEntityIdAndLangcode
func GetNodeRevision_commentsByBundleAndEntityIdAndLangcode(offset int, limit int, Bundle_ string, EntityId_ int, Langcode_ string) (*[]*NodeRevision_comment, error) {
	var _NodeRevision_comment = new([]*NodeRevision_comment)
	err := Engine.Table("node_revision__comment").Where("bundle = ? and entity_id = ? and langcode = ?", Bundle_, EntityId_, Langcode_).Limit(limit, offset).Find(_NodeRevision_comment)
	return _NodeRevision_comment, err
}

// GetNodeRevision_commentsByBundleAndEntityIdAndDelta Get NodeRevision_comments via BundleAndEntityIdAndDelta
func GetNodeRevision_commentsByBundleAndEntityIdAndDelta(offset int, limit int, Bundle_ string, EntityId_ int, Delta_ int) (*[]*NodeRevision_comment, error) {
	var _NodeRevision_comment = new([]*NodeRevision_comment)
	err := Engine.Table("node_revision__comment").Where("bundle = ? and entity_id = ? and delta = ?", Bundle_, EntityId_, Delta_).Limit(limit, offset).Find(_NodeRevision_comment)
	return _NodeRevision_comment, err
}

// GetNodeRevision_commentsByBundleAndEntityIdAndCommentStatus Get NodeRevision_comments via BundleAndEntityIdAndCommentStatus
func GetNodeRevision_commentsByBundleAndEntityIdAndCommentStatus(offset int, limit int, Bundle_ string, EntityId_ int, CommentStatus_ int) (*[]*NodeRevision_comment, error) {
	var _NodeRevision_comment = new([]*NodeRevision_comment)
	err := Engine.Table("node_revision__comment").Where("bundle = ? and entity_id = ? and comment_status = ?", Bundle_, EntityId_, CommentStatus_).Limit(limit, offset).Find(_NodeRevision_comment)
	return _NodeRevision_comment, err
}

// GetNodeRevision_commentsByBundleAndRevisionIdAndLangcode Get NodeRevision_comments via BundleAndRevisionIdAndLangcode
func GetNodeRevision_commentsByBundleAndRevisionIdAndLangcode(offset int, limit int, Bundle_ string, RevisionId_ int, Langcode_ string) (*[]*NodeRevision_comment, error) {
	var _NodeRevision_comment = new([]*NodeRevision_comment)
	err := Engine.Table("node_revision__comment").Where("bundle = ? and revision_id = ? and langcode = ?", Bundle_, RevisionId_, Langcode_).Limit(limit, offset).Find(_NodeRevision_comment)
	return _NodeRevision_comment, err
}

// GetNodeRevision_commentsByBundleAndRevisionIdAndDelta Get NodeRevision_comments via BundleAndRevisionIdAndDelta
func GetNodeRevision_commentsByBundleAndRevisionIdAndDelta(offset int, limit int, Bundle_ string, RevisionId_ int, Delta_ int) (*[]*NodeRevision_comment, error) {
	var _NodeRevision_comment = new([]*NodeRevision_comment)
	err := Engine.Table("node_revision__comment").Where("bundle = ? and revision_id = ? and delta = ?", Bundle_, RevisionId_, Delta_).Limit(limit, offset).Find(_NodeRevision_comment)
	return _NodeRevision_comment, err
}

// GetNodeRevision_commentsByBundleAndRevisionIdAndCommentStatus Get NodeRevision_comments via BundleAndRevisionIdAndCommentStatus
func GetNodeRevision_commentsByBundleAndRevisionIdAndCommentStatus(offset int, limit int, Bundle_ string, RevisionId_ int, CommentStatus_ int) (*[]*NodeRevision_comment, error) {
	var _NodeRevision_comment = new([]*NodeRevision_comment)
	err := Engine.Table("node_revision__comment").Where("bundle = ? and revision_id = ? and comment_status = ?", Bundle_, RevisionId_, CommentStatus_).Limit(limit, offset).Find(_NodeRevision_comment)
	return _NodeRevision_comment, err
}

// GetNodeRevision_commentsByBundleAndLangcodeAndDelta Get NodeRevision_comments via BundleAndLangcodeAndDelta
func GetNodeRevision_commentsByBundleAndLangcodeAndDelta(offset int, limit int, Bundle_ string, Langcode_ string, Delta_ int) (*[]*NodeRevision_comment, error) {
	var _NodeRevision_comment = new([]*NodeRevision_comment)
	err := Engine.Table("node_revision__comment").Where("bundle = ? and langcode = ? and delta = ?", Bundle_, Langcode_, Delta_).Limit(limit, offset).Find(_NodeRevision_comment)
	return _NodeRevision_comment, err
}

// GetNodeRevision_commentsByBundleAndLangcodeAndCommentStatus Get NodeRevision_comments via BundleAndLangcodeAndCommentStatus
func GetNodeRevision_commentsByBundleAndLangcodeAndCommentStatus(offset int, limit int, Bundle_ string, Langcode_ string, CommentStatus_ int) (*[]*NodeRevision_comment, error) {
	var _NodeRevision_comment = new([]*NodeRevision_comment)
	err := Engine.Table("node_revision__comment").Where("bundle = ? and langcode = ? and comment_status = ?", Bundle_, Langcode_, CommentStatus_).Limit(limit, offset).Find(_NodeRevision_comment)
	return _NodeRevision_comment, err
}

// GetNodeRevision_commentsByBundleAndDeltaAndCommentStatus Get NodeRevision_comments via BundleAndDeltaAndCommentStatus
func GetNodeRevision_commentsByBundleAndDeltaAndCommentStatus(offset int, limit int, Bundle_ string, Delta_ int, CommentStatus_ int) (*[]*NodeRevision_comment, error) {
	var _NodeRevision_comment = new([]*NodeRevision_comment)
	err := Engine.Table("node_revision__comment").Where("bundle = ? and delta = ? and comment_status = ?", Bundle_, Delta_, CommentStatus_).Limit(limit, offset).Find(_NodeRevision_comment)
	return _NodeRevision_comment, err
}

// GetNodeRevision_commentsByDeletedAndEntityIdAndRevisionId Get NodeRevision_comments via DeletedAndEntityIdAndRevisionId
func GetNodeRevision_commentsByDeletedAndEntityIdAndRevisionId(offset int, limit int, Deleted_ int, EntityId_ int, RevisionId_ int) (*[]*NodeRevision_comment, error) {
	var _NodeRevision_comment = new([]*NodeRevision_comment)
	err := Engine.Table("node_revision__comment").Where("deleted = ? and entity_id = ? and revision_id = ?", Deleted_, EntityId_, RevisionId_).Limit(limit, offset).Find(_NodeRevision_comment)
	return _NodeRevision_comment, err
}

// GetNodeRevision_commentsByDeletedAndEntityIdAndLangcode Get NodeRevision_comments via DeletedAndEntityIdAndLangcode
func GetNodeRevision_commentsByDeletedAndEntityIdAndLangcode(offset int, limit int, Deleted_ int, EntityId_ int, Langcode_ string) (*[]*NodeRevision_comment, error) {
	var _NodeRevision_comment = new([]*NodeRevision_comment)
	err := Engine.Table("node_revision__comment").Where("deleted = ? and entity_id = ? and langcode = ?", Deleted_, EntityId_, Langcode_).Limit(limit, offset).Find(_NodeRevision_comment)
	return _NodeRevision_comment, err
}

// GetNodeRevision_commentsByDeletedAndEntityIdAndDelta Get NodeRevision_comments via DeletedAndEntityIdAndDelta
func GetNodeRevision_commentsByDeletedAndEntityIdAndDelta(offset int, limit int, Deleted_ int, EntityId_ int, Delta_ int) (*[]*NodeRevision_comment, error) {
	var _NodeRevision_comment = new([]*NodeRevision_comment)
	err := Engine.Table("node_revision__comment").Where("deleted = ? and entity_id = ? and delta = ?", Deleted_, EntityId_, Delta_).Limit(limit, offset).Find(_NodeRevision_comment)
	return _NodeRevision_comment, err
}

// GetNodeRevision_commentsByDeletedAndEntityIdAndCommentStatus Get NodeRevision_comments via DeletedAndEntityIdAndCommentStatus
func GetNodeRevision_commentsByDeletedAndEntityIdAndCommentStatus(offset int, limit int, Deleted_ int, EntityId_ int, CommentStatus_ int) (*[]*NodeRevision_comment, error) {
	var _NodeRevision_comment = new([]*NodeRevision_comment)
	err := Engine.Table("node_revision__comment").Where("deleted = ? and entity_id = ? and comment_status = ?", Deleted_, EntityId_, CommentStatus_).Limit(limit, offset).Find(_NodeRevision_comment)
	return _NodeRevision_comment, err
}

// GetNodeRevision_commentsByDeletedAndRevisionIdAndLangcode Get NodeRevision_comments via DeletedAndRevisionIdAndLangcode
func GetNodeRevision_commentsByDeletedAndRevisionIdAndLangcode(offset int, limit int, Deleted_ int, RevisionId_ int, Langcode_ string) (*[]*NodeRevision_comment, error) {
	var _NodeRevision_comment = new([]*NodeRevision_comment)
	err := Engine.Table("node_revision__comment").Where("deleted = ? and revision_id = ? and langcode = ?", Deleted_, RevisionId_, Langcode_).Limit(limit, offset).Find(_NodeRevision_comment)
	return _NodeRevision_comment, err
}

// GetNodeRevision_commentsByDeletedAndRevisionIdAndDelta Get NodeRevision_comments via DeletedAndRevisionIdAndDelta
func GetNodeRevision_commentsByDeletedAndRevisionIdAndDelta(offset int, limit int, Deleted_ int, RevisionId_ int, Delta_ int) (*[]*NodeRevision_comment, error) {
	var _NodeRevision_comment = new([]*NodeRevision_comment)
	err := Engine.Table("node_revision__comment").Where("deleted = ? and revision_id = ? and delta = ?", Deleted_, RevisionId_, Delta_).Limit(limit, offset).Find(_NodeRevision_comment)
	return _NodeRevision_comment, err
}

// GetNodeRevision_commentsByDeletedAndRevisionIdAndCommentStatus Get NodeRevision_comments via DeletedAndRevisionIdAndCommentStatus
func GetNodeRevision_commentsByDeletedAndRevisionIdAndCommentStatus(offset int, limit int, Deleted_ int, RevisionId_ int, CommentStatus_ int) (*[]*NodeRevision_comment, error) {
	var _NodeRevision_comment = new([]*NodeRevision_comment)
	err := Engine.Table("node_revision__comment").Where("deleted = ? and revision_id = ? and comment_status = ?", Deleted_, RevisionId_, CommentStatus_).Limit(limit, offset).Find(_NodeRevision_comment)
	return _NodeRevision_comment, err
}

// GetNodeRevision_commentsByDeletedAndLangcodeAndDelta Get NodeRevision_comments via DeletedAndLangcodeAndDelta
func GetNodeRevision_commentsByDeletedAndLangcodeAndDelta(offset int, limit int, Deleted_ int, Langcode_ string, Delta_ int) (*[]*NodeRevision_comment, error) {
	var _NodeRevision_comment = new([]*NodeRevision_comment)
	err := Engine.Table("node_revision__comment").Where("deleted = ? and langcode = ? and delta = ?", Deleted_, Langcode_, Delta_).Limit(limit, offset).Find(_NodeRevision_comment)
	return _NodeRevision_comment, err
}

// GetNodeRevision_commentsByDeletedAndLangcodeAndCommentStatus Get NodeRevision_comments via DeletedAndLangcodeAndCommentStatus
func GetNodeRevision_commentsByDeletedAndLangcodeAndCommentStatus(offset int, limit int, Deleted_ int, Langcode_ string, CommentStatus_ int) (*[]*NodeRevision_comment, error) {
	var _NodeRevision_comment = new([]*NodeRevision_comment)
	err := Engine.Table("node_revision__comment").Where("deleted = ? and langcode = ? and comment_status = ?", Deleted_, Langcode_, CommentStatus_).Limit(limit, offset).Find(_NodeRevision_comment)
	return _NodeRevision_comment, err
}

// GetNodeRevision_commentsByDeletedAndDeltaAndCommentStatus Get NodeRevision_comments via DeletedAndDeltaAndCommentStatus
func GetNodeRevision_commentsByDeletedAndDeltaAndCommentStatus(offset int, limit int, Deleted_ int, Delta_ int, CommentStatus_ int) (*[]*NodeRevision_comment, error) {
	var _NodeRevision_comment = new([]*NodeRevision_comment)
	err := Engine.Table("node_revision__comment").Where("deleted = ? and delta = ? and comment_status = ?", Deleted_, Delta_, CommentStatus_).Limit(limit, offset).Find(_NodeRevision_comment)
	return _NodeRevision_comment, err
}

// GetNodeRevision_commentsByEntityIdAndRevisionIdAndLangcode Get NodeRevision_comments via EntityIdAndRevisionIdAndLangcode
func GetNodeRevision_commentsByEntityIdAndRevisionIdAndLangcode(offset int, limit int, EntityId_ int, RevisionId_ int, Langcode_ string) (*[]*NodeRevision_comment, error) {
	var _NodeRevision_comment = new([]*NodeRevision_comment)
	err := Engine.Table("node_revision__comment").Where("entity_id = ? and revision_id = ? and langcode = ?", EntityId_, RevisionId_, Langcode_).Limit(limit, offset).Find(_NodeRevision_comment)
	return _NodeRevision_comment, err
}

// GetNodeRevision_commentsByEntityIdAndRevisionIdAndDelta Get NodeRevision_comments via EntityIdAndRevisionIdAndDelta
func GetNodeRevision_commentsByEntityIdAndRevisionIdAndDelta(offset int, limit int, EntityId_ int, RevisionId_ int, Delta_ int) (*[]*NodeRevision_comment, error) {
	var _NodeRevision_comment = new([]*NodeRevision_comment)
	err := Engine.Table("node_revision__comment").Where("entity_id = ? and revision_id = ? and delta = ?", EntityId_, RevisionId_, Delta_).Limit(limit, offset).Find(_NodeRevision_comment)
	return _NodeRevision_comment, err
}

// GetNodeRevision_commentsByEntityIdAndRevisionIdAndCommentStatus Get NodeRevision_comments via EntityIdAndRevisionIdAndCommentStatus
func GetNodeRevision_commentsByEntityIdAndRevisionIdAndCommentStatus(offset int, limit int, EntityId_ int, RevisionId_ int, CommentStatus_ int) (*[]*NodeRevision_comment, error) {
	var _NodeRevision_comment = new([]*NodeRevision_comment)
	err := Engine.Table("node_revision__comment").Where("entity_id = ? and revision_id = ? and comment_status = ?", EntityId_, RevisionId_, CommentStatus_).Limit(limit, offset).Find(_NodeRevision_comment)
	return _NodeRevision_comment, err
}

// GetNodeRevision_commentsByEntityIdAndLangcodeAndDelta Get NodeRevision_comments via EntityIdAndLangcodeAndDelta
func GetNodeRevision_commentsByEntityIdAndLangcodeAndDelta(offset int, limit int, EntityId_ int, Langcode_ string, Delta_ int) (*[]*NodeRevision_comment, error) {
	var _NodeRevision_comment = new([]*NodeRevision_comment)
	err := Engine.Table("node_revision__comment").Where("entity_id = ? and langcode = ? and delta = ?", EntityId_, Langcode_, Delta_).Limit(limit, offset).Find(_NodeRevision_comment)
	return _NodeRevision_comment, err
}

// GetNodeRevision_commentsByEntityIdAndLangcodeAndCommentStatus Get NodeRevision_comments via EntityIdAndLangcodeAndCommentStatus
func GetNodeRevision_commentsByEntityIdAndLangcodeAndCommentStatus(offset int, limit int, EntityId_ int, Langcode_ string, CommentStatus_ int) (*[]*NodeRevision_comment, error) {
	var _NodeRevision_comment = new([]*NodeRevision_comment)
	err := Engine.Table("node_revision__comment").Where("entity_id = ? and langcode = ? and comment_status = ?", EntityId_, Langcode_, CommentStatus_).Limit(limit, offset).Find(_NodeRevision_comment)
	return _NodeRevision_comment, err
}

// GetNodeRevision_commentsByEntityIdAndDeltaAndCommentStatus Get NodeRevision_comments via EntityIdAndDeltaAndCommentStatus
func GetNodeRevision_commentsByEntityIdAndDeltaAndCommentStatus(offset int, limit int, EntityId_ int, Delta_ int, CommentStatus_ int) (*[]*NodeRevision_comment, error) {
	var _NodeRevision_comment = new([]*NodeRevision_comment)
	err := Engine.Table("node_revision__comment").Where("entity_id = ? and delta = ? and comment_status = ?", EntityId_, Delta_, CommentStatus_).Limit(limit, offset).Find(_NodeRevision_comment)
	return _NodeRevision_comment, err
}

// GetNodeRevision_commentsByRevisionIdAndLangcodeAndDelta Get NodeRevision_comments via RevisionIdAndLangcodeAndDelta
func GetNodeRevision_commentsByRevisionIdAndLangcodeAndDelta(offset int, limit int, RevisionId_ int, Langcode_ string, Delta_ int) (*[]*NodeRevision_comment, error) {
	var _NodeRevision_comment = new([]*NodeRevision_comment)
	err := Engine.Table("node_revision__comment").Where("revision_id = ? and langcode = ? and delta = ?", RevisionId_, Langcode_, Delta_).Limit(limit, offset).Find(_NodeRevision_comment)
	return _NodeRevision_comment, err
}

// GetNodeRevision_commentsByRevisionIdAndLangcodeAndCommentStatus Get NodeRevision_comments via RevisionIdAndLangcodeAndCommentStatus
func GetNodeRevision_commentsByRevisionIdAndLangcodeAndCommentStatus(offset int, limit int, RevisionId_ int, Langcode_ string, CommentStatus_ int) (*[]*NodeRevision_comment, error) {
	var _NodeRevision_comment = new([]*NodeRevision_comment)
	err := Engine.Table("node_revision__comment").Where("revision_id = ? and langcode = ? and comment_status = ?", RevisionId_, Langcode_, CommentStatus_).Limit(limit, offset).Find(_NodeRevision_comment)
	return _NodeRevision_comment, err
}

// GetNodeRevision_commentsByRevisionIdAndDeltaAndCommentStatus Get NodeRevision_comments via RevisionIdAndDeltaAndCommentStatus
func GetNodeRevision_commentsByRevisionIdAndDeltaAndCommentStatus(offset int, limit int, RevisionId_ int, Delta_ int, CommentStatus_ int) (*[]*NodeRevision_comment, error) {
	var _NodeRevision_comment = new([]*NodeRevision_comment)
	err := Engine.Table("node_revision__comment").Where("revision_id = ? and delta = ? and comment_status = ?", RevisionId_, Delta_, CommentStatus_).Limit(limit, offset).Find(_NodeRevision_comment)
	return _NodeRevision_comment, err
}

// GetNodeRevision_commentsByLangcodeAndDeltaAndCommentStatus Get NodeRevision_comments via LangcodeAndDeltaAndCommentStatus
func GetNodeRevision_commentsByLangcodeAndDeltaAndCommentStatus(offset int, limit int, Langcode_ string, Delta_ int, CommentStatus_ int) (*[]*NodeRevision_comment, error) {
	var _NodeRevision_comment = new([]*NodeRevision_comment)
	err := Engine.Table("node_revision__comment").Where("langcode = ? and delta = ? and comment_status = ?", Langcode_, Delta_, CommentStatus_).Limit(limit, offset).Find(_NodeRevision_comment)
	return _NodeRevision_comment, err
}

// GetNodeRevision_commentsByBundleAndDeleted Get NodeRevision_comments via BundleAndDeleted
func GetNodeRevision_commentsByBundleAndDeleted(offset int, limit int, Bundle_ string, Deleted_ int) (*[]*NodeRevision_comment, error) {
	var _NodeRevision_comment = new([]*NodeRevision_comment)
	err := Engine.Table("node_revision__comment").Where("bundle = ? and deleted = ?", Bundle_, Deleted_).Limit(limit, offset).Find(_NodeRevision_comment)
	return _NodeRevision_comment, err
}

// GetNodeRevision_commentsByBundleAndEntityId Get NodeRevision_comments via BundleAndEntityId
func GetNodeRevision_commentsByBundleAndEntityId(offset int, limit int, Bundle_ string, EntityId_ int) (*[]*NodeRevision_comment, error) {
	var _NodeRevision_comment = new([]*NodeRevision_comment)
	err := Engine.Table("node_revision__comment").Where("bundle = ? and entity_id = ?", Bundle_, EntityId_).Limit(limit, offset).Find(_NodeRevision_comment)
	return _NodeRevision_comment, err
}

// GetNodeRevision_commentsByBundleAndRevisionId Get NodeRevision_comments via BundleAndRevisionId
func GetNodeRevision_commentsByBundleAndRevisionId(offset int, limit int, Bundle_ string, RevisionId_ int) (*[]*NodeRevision_comment, error) {
	var _NodeRevision_comment = new([]*NodeRevision_comment)
	err := Engine.Table("node_revision__comment").Where("bundle = ? and revision_id = ?", Bundle_, RevisionId_).Limit(limit, offset).Find(_NodeRevision_comment)
	return _NodeRevision_comment, err
}

// GetNodeRevision_commentsByBundleAndLangcode Get NodeRevision_comments via BundleAndLangcode
func GetNodeRevision_commentsByBundleAndLangcode(offset int, limit int, Bundle_ string, Langcode_ string) (*[]*NodeRevision_comment, error) {
	var _NodeRevision_comment = new([]*NodeRevision_comment)
	err := Engine.Table("node_revision__comment").Where("bundle = ? and langcode = ?", Bundle_, Langcode_).Limit(limit, offset).Find(_NodeRevision_comment)
	return _NodeRevision_comment, err
}

// GetNodeRevision_commentsByBundleAndDelta Get NodeRevision_comments via BundleAndDelta
func GetNodeRevision_commentsByBundleAndDelta(offset int, limit int, Bundle_ string, Delta_ int) (*[]*NodeRevision_comment, error) {
	var _NodeRevision_comment = new([]*NodeRevision_comment)
	err := Engine.Table("node_revision__comment").Where("bundle = ? and delta = ?", Bundle_, Delta_).Limit(limit, offset).Find(_NodeRevision_comment)
	return _NodeRevision_comment, err
}

// GetNodeRevision_commentsByBundleAndCommentStatus Get NodeRevision_comments via BundleAndCommentStatus
func GetNodeRevision_commentsByBundleAndCommentStatus(offset int, limit int, Bundle_ string, CommentStatus_ int) (*[]*NodeRevision_comment, error) {
	var _NodeRevision_comment = new([]*NodeRevision_comment)
	err := Engine.Table("node_revision__comment").Where("bundle = ? and comment_status = ?", Bundle_, CommentStatus_).Limit(limit, offset).Find(_NodeRevision_comment)
	return _NodeRevision_comment, err
}

// GetNodeRevision_commentsByDeletedAndEntityId Get NodeRevision_comments via DeletedAndEntityId
func GetNodeRevision_commentsByDeletedAndEntityId(offset int, limit int, Deleted_ int, EntityId_ int) (*[]*NodeRevision_comment, error) {
	var _NodeRevision_comment = new([]*NodeRevision_comment)
	err := Engine.Table("node_revision__comment").Where("deleted = ? and entity_id = ?", Deleted_, EntityId_).Limit(limit, offset).Find(_NodeRevision_comment)
	return _NodeRevision_comment, err
}

// GetNodeRevision_commentsByDeletedAndRevisionId Get NodeRevision_comments via DeletedAndRevisionId
func GetNodeRevision_commentsByDeletedAndRevisionId(offset int, limit int, Deleted_ int, RevisionId_ int) (*[]*NodeRevision_comment, error) {
	var _NodeRevision_comment = new([]*NodeRevision_comment)
	err := Engine.Table("node_revision__comment").Where("deleted = ? and revision_id = ?", Deleted_, RevisionId_).Limit(limit, offset).Find(_NodeRevision_comment)
	return _NodeRevision_comment, err
}

// GetNodeRevision_commentsByDeletedAndLangcode Get NodeRevision_comments via DeletedAndLangcode
func GetNodeRevision_commentsByDeletedAndLangcode(offset int, limit int, Deleted_ int, Langcode_ string) (*[]*NodeRevision_comment, error) {
	var _NodeRevision_comment = new([]*NodeRevision_comment)
	err := Engine.Table("node_revision__comment").Where("deleted = ? and langcode = ?", Deleted_, Langcode_).Limit(limit, offset).Find(_NodeRevision_comment)
	return _NodeRevision_comment, err
}

// GetNodeRevision_commentsByDeletedAndDelta Get NodeRevision_comments via DeletedAndDelta
func GetNodeRevision_commentsByDeletedAndDelta(offset int, limit int, Deleted_ int, Delta_ int) (*[]*NodeRevision_comment, error) {
	var _NodeRevision_comment = new([]*NodeRevision_comment)
	err := Engine.Table("node_revision__comment").Where("deleted = ? and delta = ?", Deleted_, Delta_).Limit(limit, offset).Find(_NodeRevision_comment)
	return _NodeRevision_comment, err
}

// GetNodeRevision_commentsByDeletedAndCommentStatus Get NodeRevision_comments via DeletedAndCommentStatus
func GetNodeRevision_commentsByDeletedAndCommentStatus(offset int, limit int, Deleted_ int, CommentStatus_ int) (*[]*NodeRevision_comment, error) {
	var _NodeRevision_comment = new([]*NodeRevision_comment)
	err := Engine.Table("node_revision__comment").Where("deleted = ? and comment_status = ?", Deleted_, CommentStatus_).Limit(limit, offset).Find(_NodeRevision_comment)
	return _NodeRevision_comment, err
}

// GetNodeRevision_commentsByEntityIdAndRevisionId Get NodeRevision_comments via EntityIdAndRevisionId
func GetNodeRevision_commentsByEntityIdAndRevisionId(offset int, limit int, EntityId_ int, RevisionId_ int) (*[]*NodeRevision_comment, error) {
	var _NodeRevision_comment = new([]*NodeRevision_comment)
	err := Engine.Table("node_revision__comment").Where("entity_id = ? and revision_id = ?", EntityId_, RevisionId_).Limit(limit, offset).Find(_NodeRevision_comment)
	return _NodeRevision_comment, err
}

// GetNodeRevision_commentsByEntityIdAndLangcode Get NodeRevision_comments via EntityIdAndLangcode
func GetNodeRevision_commentsByEntityIdAndLangcode(offset int, limit int, EntityId_ int, Langcode_ string) (*[]*NodeRevision_comment, error) {
	var _NodeRevision_comment = new([]*NodeRevision_comment)
	err := Engine.Table("node_revision__comment").Where("entity_id = ? and langcode = ?", EntityId_, Langcode_).Limit(limit, offset).Find(_NodeRevision_comment)
	return _NodeRevision_comment, err
}

// GetNodeRevision_commentsByEntityIdAndDelta Get NodeRevision_comments via EntityIdAndDelta
func GetNodeRevision_commentsByEntityIdAndDelta(offset int, limit int, EntityId_ int, Delta_ int) (*[]*NodeRevision_comment, error) {
	var _NodeRevision_comment = new([]*NodeRevision_comment)
	err := Engine.Table("node_revision__comment").Where("entity_id = ? and delta = ?", EntityId_, Delta_).Limit(limit, offset).Find(_NodeRevision_comment)
	return _NodeRevision_comment, err
}

// GetNodeRevision_commentsByEntityIdAndCommentStatus Get NodeRevision_comments via EntityIdAndCommentStatus
func GetNodeRevision_commentsByEntityIdAndCommentStatus(offset int, limit int, EntityId_ int, CommentStatus_ int) (*[]*NodeRevision_comment, error) {
	var _NodeRevision_comment = new([]*NodeRevision_comment)
	err := Engine.Table("node_revision__comment").Where("entity_id = ? and comment_status = ?", EntityId_, CommentStatus_).Limit(limit, offset).Find(_NodeRevision_comment)
	return _NodeRevision_comment, err
}

// GetNodeRevision_commentsByRevisionIdAndLangcode Get NodeRevision_comments via RevisionIdAndLangcode
func GetNodeRevision_commentsByRevisionIdAndLangcode(offset int, limit int, RevisionId_ int, Langcode_ string) (*[]*NodeRevision_comment, error) {
	var _NodeRevision_comment = new([]*NodeRevision_comment)
	err := Engine.Table("node_revision__comment").Where("revision_id = ? and langcode = ?", RevisionId_, Langcode_).Limit(limit, offset).Find(_NodeRevision_comment)
	return _NodeRevision_comment, err
}

// GetNodeRevision_commentsByRevisionIdAndDelta Get NodeRevision_comments via RevisionIdAndDelta
func GetNodeRevision_commentsByRevisionIdAndDelta(offset int, limit int, RevisionId_ int, Delta_ int) (*[]*NodeRevision_comment, error) {
	var _NodeRevision_comment = new([]*NodeRevision_comment)
	err := Engine.Table("node_revision__comment").Where("revision_id = ? and delta = ?", RevisionId_, Delta_).Limit(limit, offset).Find(_NodeRevision_comment)
	return _NodeRevision_comment, err
}

// GetNodeRevision_commentsByRevisionIdAndCommentStatus Get NodeRevision_comments via RevisionIdAndCommentStatus
func GetNodeRevision_commentsByRevisionIdAndCommentStatus(offset int, limit int, RevisionId_ int, CommentStatus_ int) (*[]*NodeRevision_comment, error) {
	var _NodeRevision_comment = new([]*NodeRevision_comment)
	err := Engine.Table("node_revision__comment").Where("revision_id = ? and comment_status = ?", RevisionId_, CommentStatus_).Limit(limit, offset).Find(_NodeRevision_comment)
	return _NodeRevision_comment, err
}

// GetNodeRevision_commentsByLangcodeAndDelta Get NodeRevision_comments via LangcodeAndDelta
func GetNodeRevision_commentsByLangcodeAndDelta(offset int, limit int, Langcode_ string, Delta_ int) (*[]*NodeRevision_comment, error) {
	var _NodeRevision_comment = new([]*NodeRevision_comment)
	err := Engine.Table("node_revision__comment").Where("langcode = ? and delta = ?", Langcode_, Delta_).Limit(limit, offset).Find(_NodeRevision_comment)
	return _NodeRevision_comment, err
}

// GetNodeRevision_commentsByLangcodeAndCommentStatus Get NodeRevision_comments via LangcodeAndCommentStatus
func GetNodeRevision_commentsByLangcodeAndCommentStatus(offset int, limit int, Langcode_ string, CommentStatus_ int) (*[]*NodeRevision_comment, error) {
	var _NodeRevision_comment = new([]*NodeRevision_comment)
	err := Engine.Table("node_revision__comment").Where("langcode = ? and comment_status = ?", Langcode_, CommentStatus_).Limit(limit, offset).Find(_NodeRevision_comment)
	return _NodeRevision_comment, err
}

// GetNodeRevision_commentsByDeltaAndCommentStatus Get NodeRevision_comments via DeltaAndCommentStatus
func GetNodeRevision_commentsByDeltaAndCommentStatus(offset int, limit int, Delta_ int, CommentStatus_ int) (*[]*NodeRevision_comment, error) {
	var _NodeRevision_comment = new([]*NodeRevision_comment)
	err := Engine.Table("node_revision__comment").Where("delta = ? and comment_status = ?", Delta_, CommentStatus_).Limit(limit, offset).Find(_NodeRevision_comment)
	return _NodeRevision_comment, err
}

// GetNodeRevision_comments Get NodeRevision_comments via field
func GetNodeRevision_comments(offset int, limit int, field string) (*[]*NodeRevision_comment, error) {
	var _NodeRevision_comment = new([]*NodeRevision_comment)
	err := Engine.Table("node_revision__comment").Limit(limit, offset).Desc(field).Find(_NodeRevision_comment)
	return _NodeRevision_comment, err
}

// GetNodeRevision_commentsViaBundle Get NodeRevision_comments via Bundle
func GetNodeRevision_commentsViaBundle(offset int, limit int, Bundle_ string, field string) (*[]*NodeRevision_comment, error) {
	var _NodeRevision_comment = new([]*NodeRevision_comment)
	err := Engine.Table("node_revision__comment").Where("bundle = ?", Bundle_).Limit(limit, offset).Desc(field).Find(_NodeRevision_comment)
	return _NodeRevision_comment, err
}

// GetNodeRevision_commentsViaDeleted Get NodeRevision_comments via Deleted
func GetNodeRevision_commentsViaDeleted(offset int, limit int, Deleted_ int, field string) (*[]*NodeRevision_comment, error) {
	var _NodeRevision_comment = new([]*NodeRevision_comment)
	err := Engine.Table("node_revision__comment").Where("deleted = ?", Deleted_).Limit(limit, offset).Desc(field).Find(_NodeRevision_comment)
	return _NodeRevision_comment, err
}

// GetNodeRevision_commentsViaEntityId Get NodeRevision_comments via EntityId
func GetNodeRevision_commentsViaEntityId(offset int, limit int, EntityId_ int, field string) (*[]*NodeRevision_comment, error) {
	var _NodeRevision_comment = new([]*NodeRevision_comment)
	err := Engine.Table("node_revision__comment").Where("entity_id = ?", EntityId_).Limit(limit, offset).Desc(field).Find(_NodeRevision_comment)
	return _NodeRevision_comment, err
}

// GetNodeRevision_commentsViaRevisionId Get NodeRevision_comments via RevisionId
func GetNodeRevision_commentsViaRevisionId(offset int, limit int, RevisionId_ int, field string) (*[]*NodeRevision_comment, error) {
	var _NodeRevision_comment = new([]*NodeRevision_comment)
	err := Engine.Table("node_revision__comment").Where("revision_id = ?", RevisionId_).Limit(limit, offset).Desc(field).Find(_NodeRevision_comment)
	return _NodeRevision_comment, err
}

// GetNodeRevision_commentsViaLangcode Get NodeRevision_comments via Langcode
func GetNodeRevision_commentsViaLangcode(offset int, limit int, Langcode_ string, field string) (*[]*NodeRevision_comment, error) {
	var _NodeRevision_comment = new([]*NodeRevision_comment)
	err := Engine.Table("node_revision__comment").Where("langcode = ?", Langcode_).Limit(limit, offset).Desc(field).Find(_NodeRevision_comment)
	return _NodeRevision_comment, err
}

// GetNodeRevision_commentsViaDelta Get NodeRevision_comments via Delta
func GetNodeRevision_commentsViaDelta(offset int, limit int, Delta_ int, field string) (*[]*NodeRevision_comment, error) {
	var _NodeRevision_comment = new([]*NodeRevision_comment)
	err := Engine.Table("node_revision__comment").Where("delta = ?", Delta_).Limit(limit, offset).Desc(field).Find(_NodeRevision_comment)
	return _NodeRevision_comment, err
}

// GetNodeRevision_commentsViaCommentStatus Get NodeRevision_comments via CommentStatus
func GetNodeRevision_commentsViaCommentStatus(offset int, limit int, CommentStatus_ int, field string) (*[]*NodeRevision_comment, error) {
	var _NodeRevision_comment = new([]*NodeRevision_comment)
	err := Engine.Table("node_revision__comment").Where("comment_status = ?", CommentStatus_).Limit(limit, offset).Desc(field).Find(_NodeRevision_comment)
	return _NodeRevision_comment, err
}

// HasNodeRevision_commentViaBundle Has NodeRevision_comment via Bundle
func HasNodeRevision_commentViaBundle(iBundle string) bool {
	if has, err := Engine.Where("bundle = ?", iBundle).Get(new(NodeRevision_comment)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasNodeRevision_commentViaDeleted Has NodeRevision_comment via Deleted
func HasNodeRevision_commentViaDeleted(iDeleted int) bool {
	if has, err := Engine.Where("deleted = ?", iDeleted).Get(new(NodeRevision_comment)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasNodeRevision_commentViaEntityId Has NodeRevision_comment via EntityId
func HasNodeRevision_commentViaEntityId(iEntityId int) bool {
	if has, err := Engine.Where("entity_id = ?", iEntityId).Get(new(NodeRevision_comment)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasNodeRevision_commentViaRevisionId Has NodeRevision_comment via RevisionId
func HasNodeRevision_commentViaRevisionId(iRevisionId int) bool {
	if has, err := Engine.Where("revision_id = ?", iRevisionId).Get(new(NodeRevision_comment)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasNodeRevision_commentViaLangcode Has NodeRevision_comment via Langcode
func HasNodeRevision_commentViaLangcode(iLangcode string) bool {
	if has, err := Engine.Where("langcode = ?", iLangcode).Get(new(NodeRevision_comment)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasNodeRevision_commentViaDelta Has NodeRevision_comment via Delta
func HasNodeRevision_commentViaDelta(iDelta int) bool {
	if has, err := Engine.Where("delta = ?", iDelta).Get(new(NodeRevision_comment)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasNodeRevision_commentViaCommentStatus Has NodeRevision_comment via CommentStatus
func HasNodeRevision_commentViaCommentStatus(iCommentStatus int) bool {
	if has, err := Engine.Where("comment_status = ?", iCommentStatus).Get(new(NodeRevision_comment)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// GetNodeRevision_commentViaBundle Get NodeRevision_comment via Bundle
func GetNodeRevision_commentViaBundle(iBundle string) (*NodeRevision_comment, error) {
	var _NodeRevision_comment = &NodeRevision_comment{Bundle: iBundle}
	has, err := Engine.Get(_NodeRevision_comment)
	if has {
		return _NodeRevision_comment, err
	} else {
		return nil, err
	}
}

// GetNodeRevision_commentViaDeleted Get NodeRevision_comment via Deleted
func GetNodeRevision_commentViaDeleted(iDeleted int) (*NodeRevision_comment, error) {
	var _NodeRevision_comment = &NodeRevision_comment{Deleted: iDeleted}
	has, err := Engine.Get(_NodeRevision_comment)
	if has {
		return _NodeRevision_comment, err
	} else {
		return nil, err
	}
}

// GetNodeRevision_commentViaEntityId Get NodeRevision_comment via EntityId
func GetNodeRevision_commentViaEntityId(iEntityId int) (*NodeRevision_comment, error) {
	var _NodeRevision_comment = &NodeRevision_comment{EntityId: iEntityId}
	has, err := Engine.Get(_NodeRevision_comment)
	if has {
		return _NodeRevision_comment, err
	} else {
		return nil, err
	}
}

// GetNodeRevision_commentViaRevisionId Get NodeRevision_comment via RevisionId
func GetNodeRevision_commentViaRevisionId(iRevisionId int) (*NodeRevision_comment, error) {
	var _NodeRevision_comment = &NodeRevision_comment{RevisionId: iRevisionId}
	has, err := Engine.Get(_NodeRevision_comment)
	if has {
		return _NodeRevision_comment, err
	} else {
		return nil, err
	}
}

// GetNodeRevision_commentViaLangcode Get NodeRevision_comment via Langcode
func GetNodeRevision_commentViaLangcode(iLangcode string) (*NodeRevision_comment, error) {
	var _NodeRevision_comment = &NodeRevision_comment{Langcode: iLangcode}
	has, err := Engine.Get(_NodeRevision_comment)
	if has {
		return _NodeRevision_comment, err
	} else {
		return nil, err
	}
}

// GetNodeRevision_commentViaDelta Get NodeRevision_comment via Delta
func GetNodeRevision_commentViaDelta(iDelta int) (*NodeRevision_comment, error) {
	var _NodeRevision_comment = &NodeRevision_comment{Delta: iDelta}
	has, err := Engine.Get(_NodeRevision_comment)
	if has {
		return _NodeRevision_comment, err
	} else {
		return nil, err
	}
}

// GetNodeRevision_commentViaCommentStatus Get NodeRevision_comment via CommentStatus
func GetNodeRevision_commentViaCommentStatus(iCommentStatus int) (*NodeRevision_comment, error) {
	var _NodeRevision_comment = &NodeRevision_comment{CommentStatus: iCommentStatus}
	has, err := Engine.Get(_NodeRevision_comment)
	if has {
		return _NodeRevision_comment, err
	} else {
		return nil, err
	}
}

// SetNodeRevision_commentViaBundle Set NodeRevision_comment via Bundle
func SetNodeRevision_commentViaBundle(iBundle string, node_revision__comment *NodeRevision_comment) (int64, error) {
	node_revision__comment.Bundle = iBundle
	return Engine.Insert(node_revision__comment)
}

// SetNodeRevision_commentViaDeleted Set NodeRevision_comment via Deleted
func SetNodeRevision_commentViaDeleted(iDeleted int, node_revision__comment *NodeRevision_comment) (int64, error) {
	node_revision__comment.Deleted = iDeleted
	return Engine.Insert(node_revision__comment)
}

// SetNodeRevision_commentViaEntityId Set NodeRevision_comment via EntityId
func SetNodeRevision_commentViaEntityId(iEntityId int, node_revision__comment *NodeRevision_comment) (int64, error) {
	node_revision__comment.EntityId = iEntityId
	return Engine.Insert(node_revision__comment)
}

// SetNodeRevision_commentViaRevisionId Set NodeRevision_comment via RevisionId
func SetNodeRevision_commentViaRevisionId(iRevisionId int, node_revision__comment *NodeRevision_comment) (int64, error) {
	node_revision__comment.RevisionId = iRevisionId
	return Engine.Insert(node_revision__comment)
}

// SetNodeRevision_commentViaLangcode Set NodeRevision_comment via Langcode
func SetNodeRevision_commentViaLangcode(iLangcode string, node_revision__comment *NodeRevision_comment) (int64, error) {
	node_revision__comment.Langcode = iLangcode
	return Engine.Insert(node_revision__comment)
}

// SetNodeRevision_commentViaDelta Set NodeRevision_comment via Delta
func SetNodeRevision_commentViaDelta(iDelta int, node_revision__comment *NodeRevision_comment) (int64, error) {
	node_revision__comment.Delta = iDelta
	return Engine.Insert(node_revision__comment)
}

// SetNodeRevision_commentViaCommentStatus Set NodeRevision_comment via CommentStatus
func SetNodeRevision_commentViaCommentStatus(iCommentStatus int, node_revision__comment *NodeRevision_comment) (int64, error) {
	node_revision__comment.CommentStatus = iCommentStatus
	return Engine.Insert(node_revision__comment)
}

// AddNodeRevision_comment Add NodeRevision_comment via all columns
func AddNodeRevision_comment(iBundle string, iDeleted int, iEntityId int, iRevisionId int, iLangcode string, iDelta int, iCommentStatus int) error {
	_NodeRevision_comment := &NodeRevision_comment{Bundle: iBundle, Deleted: iDeleted, EntityId: iEntityId, RevisionId: iRevisionId, Langcode: iLangcode, Delta: iDelta, CommentStatus: iCommentStatus}
	if _, err := Engine.Insert(_NodeRevision_comment); err != nil {
		return err
	} else {
		return nil
	}
}

// PostNodeRevision_comment Post NodeRevision_comment via iNodeRevision_comment
func PostNodeRevision_comment(iNodeRevision_comment *NodeRevision_comment) (string, error) {
	_, err := Engine.Insert(iNodeRevision_comment)
	return iNodeRevision_comment.Bundle, err
}

// PutNodeRevision_comment Put NodeRevision_comment
func PutNodeRevision_comment(iNodeRevision_comment *NodeRevision_comment) (string, error) {
	_, err := Engine.Id(iNodeRevision_comment.Bundle).Update(iNodeRevision_comment)
	return iNodeRevision_comment.Bundle, err
}

// PutNodeRevision_commentViaBundle Put NodeRevision_comment via Bundle
func PutNodeRevision_commentViaBundle(Bundle_ string, iNodeRevision_comment *NodeRevision_comment) (int64, error) {
	row, err := Engine.Update(iNodeRevision_comment, &NodeRevision_comment{Bundle: Bundle_})
	return row, err
}

// PutNodeRevision_commentViaDeleted Put NodeRevision_comment via Deleted
func PutNodeRevision_commentViaDeleted(Deleted_ int, iNodeRevision_comment *NodeRevision_comment) (int64, error) {
	row, err := Engine.Update(iNodeRevision_comment, &NodeRevision_comment{Deleted: Deleted_})
	return row, err
}

// PutNodeRevision_commentViaEntityId Put NodeRevision_comment via EntityId
func PutNodeRevision_commentViaEntityId(EntityId_ int, iNodeRevision_comment *NodeRevision_comment) (int64, error) {
	row, err := Engine.Update(iNodeRevision_comment, &NodeRevision_comment{EntityId: EntityId_})
	return row, err
}

// PutNodeRevision_commentViaRevisionId Put NodeRevision_comment via RevisionId
func PutNodeRevision_commentViaRevisionId(RevisionId_ int, iNodeRevision_comment *NodeRevision_comment) (int64, error) {
	row, err := Engine.Update(iNodeRevision_comment, &NodeRevision_comment{RevisionId: RevisionId_})
	return row, err
}

// PutNodeRevision_commentViaLangcode Put NodeRevision_comment via Langcode
func PutNodeRevision_commentViaLangcode(Langcode_ string, iNodeRevision_comment *NodeRevision_comment) (int64, error) {
	row, err := Engine.Update(iNodeRevision_comment, &NodeRevision_comment{Langcode: Langcode_})
	return row, err
}

// PutNodeRevision_commentViaDelta Put NodeRevision_comment via Delta
func PutNodeRevision_commentViaDelta(Delta_ int, iNodeRevision_comment *NodeRevision_comment) (int64, error) {
	row, err := Engine.Update(iNodeRevision_comment, &NodeRevision_comment{Delta: Delta_})
	return row, err
}

// PutNodeRevision_commentViaCommentStatus Put NodeRevision_comment via CommentStatus
func PutNodeRevision_commentViaCommentStatus(CommentStatus_ int, iNodeRevision_comment *NodeRevision_comment) (int64, error) {
	row, err := Engine.Update(iNodeRevision_comment, &NodeRevision_comment{CommentStatus: CommentStatus_})
	return row, err
}

// UpdateNodeRevision_commentViaBundle via map[string]interface{}{}
func UpdateNodeRevision_commentViaBundle(iBundle string, iNodeRevision_commentMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(NodeRevision_comment)).Where("bundle = ?", iBundle).Update(iNodeRevision_commentMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateNodeRevision_commentViaDeleted via map[string]interface{}{}
func UpdateNodeRevision_commentViaDeleted(iDeleted int, iNodeRevision_commentMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(NodeRevision_comment)).Where("deleted = ?", iDeleted).Update(iNodeRevision_commentMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateNodeRevision_commentViaEntityId via map[string]interface{}{}
func UpdateNodeRevision_commentViaEntityId(iEntityId int, iNodeRevision_commentMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(NodeRevision_comment)).Where("entity_id = ?", iEntityId).Update(iNodeRevision_commentMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateNodeRevision_commentViaRevisionId via map[string]interface{}{}
func UpdateNodeRevision_commentViaRevisionId(iRevisionId int, iNodeRevision_commentMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(NodeRevision_comment)).Where("revision_id = ?", iRevisionId).Update(iNodeRevision_commentMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateNodeRevision_commentViaLangcode via map[string]interface{}{}
func UpdateNodeRevision_commentViaLangcode(iLangcode string, iNodeRevision_commentMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(NodeRevision_comment)).Where("langcode = ?", iLangcode).Update(iNodeRevision_commentMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateNodeRevision_commentViaDelta via map[string]interface{}{}
func UpdateNodeRevision_commentViaDelta(iDelta int, iNodeRevision_commentMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(NodeRevision_comment)).Where("delta = ?", iDelta).Update(iNodeRevision_commentMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateNodeRevision_commentViaCommentStatus via map[string]interface{}{}
func UpdateNodeRevision_commentViaCommentStatus(iCommentStatus int, iNodeRevision_commentMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(NodeRevision_comment)).Where("comment_status = ?", iCommentStatus).Update(iNodeRevision_commentMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// DeleteNodeRevision_comment Delete NodeRevision_comment
func DeleteNodeRevision_comment(iBundle string) (int64, error) {
	row, err := Engine.Id(iBundle).Delete(new(NodeRevision_comment))
	return row, err
}

// DeleteNodeRevision_commentViaBundle Delete NodeRevision_comment via Bundle
func DeleteNodeRevision_commentViaBundle(iBundle string) (err error) {
	var has bool
	var _NodeRevision_comment = &NodeRevision_comment{Bundle: iBundle}
	if has, err = Engine.Get(_NodeRevision_comment); (has == true) && (err == nil) {
		if row, err := Engine.Where("bundle = ?", iBundle).Delete(new(NodeRevision_comment)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteNodeRevision_commentViaDeleted Delete NodeRevision_comment via Deleted
func DeleteNodeRevision_commentViaDeleted(iDeleted int) (err error) {
	var has bool
	var _NodeRevision_comment = &NodeRevision_comment{Deleted: iDeleted}
	if has, err = Engine.Get(_NodeRevision_comment); (has == true) && (err == nil) {
		if row, err := Engine.Where("deleted = ?", iDeleted).Delete(new(NodeRevision_comment)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteNodeRevision_commentViaEntityId Delete NodeRevision_comment via EntityId
func DeleteNodeRevision_commentViaEntityId(iEntityId int) (err error) {
	var has bool
	var _NodeRevision_comment = &NodeRevision_comment{EntityId: iEntityId}
	if has, err = Engine.Get(_NodeRevision_comment); (has == true) && (err == nil) {
		if row, err := Engine.Where("entity_id = ?", iEntityId).Delete(new(NodeRevision_comment)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteNodeRevision_commentViaRevisionId Delete NodeRevision_comment via RevisionId
func DeleteNodeRevision_commentViaRevisionId(iRevisionId int) (err error) {
	var has bool
	var _NodeRevision_comment = &NodeRevision_comment{RevisionId: iRevisionId}
	if has, err = Engine.Get(_NodeRevision_comment); (has == true) && (err == nil) {
		if row, err := Engine.Where("revision_id = ?", iRevisionId).Delete(new(NodeRevision_comment)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteNodeRevision_commentViaLangcode Delete NodeRevision_comment via Langcode
func DeleteNodeRevision_commentViaLangcode(iLangcode string) (err error) {
	var has bool
	var _NodeRevision_comment = &NodeRevision_comment{Langcode: iLangcode}
	if has, err = Engine.Get(_NodeRevision_comment); (has == true) && (err == nil) {
		if row, err := Engine.Where("langcode = ?", iLangcode).Delete(new(NodeRevision_comment)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteNodeRevision_commentViaDelta Delete NodeRevision_comment via Delta
func DeleteNodeRevision_commentViaDelta(iDelta int) (err error) {
	var has bool
	var _NodeRevision_comment = &NodeRevision_comment{Delta: iDelta}
	if has, err = Engine.Get(_NodeRevision_comment); (has == true) && (err == nil) {
		if row, err := Engine.Where("delta = ?", iDelta).Delete(new(NodeRevision_comment)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteNodeRevision_commentViaCommentStatus Delete NodeRevision_comment via CommentStatus
func DeleteNodeRevision_commentViaCommentStatus(iCommentStatus int) (err error) {
	var has bool
	var _NodeRevision_comment = &NodeRevision_comment{CommentStatus: iCommentStatus}
	if has, err = Engine.Get(_NodeRevision_comment); (has == true) && (err == nil) {
		if row, err := Engine.Where("comment_status = ?", iCommentStatus).Delete(new(NodeRevision_comment)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}
