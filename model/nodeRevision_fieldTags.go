package model

type NodeRevision_fieldTags struct {
	Bundle            string `xorm:"not null default '' index VARCHAR(128)"`
	Deleted           int    `xorm:"not null pk default 0 TINYINT(4)"`
	EntityId          int    `xorm:"not null pk INT(10)"`
	RevisionId        int    `xorm:"not null pk index INT(10)"`
	Langcode          string `xorm:"not null pk default '' VARCHAR(32)"`
	Delta             int    `xorm:"not null pk INT(10)"`
	FieldTagsTargetId int    `xorm:"not null index INT(10)"`
}

// GetNodeRevision_fieldTagsesCount NodeRevision_fieldTagss' Count
func GetNodeRevision_fieldTagsesCount(offset int, limit int) (int64, error) {
	total, err := Engine.Limit(limit, offset).Count(&NodeRevision_fieldTags{})
	return total, err
}

// GetNodeRevision_fieldTagsCountViaBundle Get NodeRevision_fieldTags via Bundle
func GetNodeRevision_fieldTagsCountViaBundle(iBundle string) int64 {
	n, _ := Engine.Where("bundle = ?", iBundle).Count(&NodeRevision_fieldTags{Bundle: iBundle})
	return n
}

// GetNodeRevision_fieldTagsCountViaDeleted Get NodeRevision_fieldTags via Deleted
func GetNodeRevision_fieldTagsCountViaDeleted(iDeleted int) int64 {
	n, _ := Engine.Where("deleted = ?", iDeleted).Count(&NodeRevision_fieldTags{Deleted: iDeleted})
	return n
}

// GetNodeRevision_fieldTagsCountViaEntityId Get NodeRevision_fieldTags via EntityId
func GetNodeRevision_fieldTagsCountViaEntityId(iEntityId int) int64 {
	n, _ := Engine.Where("entity_id = ?", iEntityId).Count(&NodeRevision_fieldTags{EntityId: iEntityId})
	return n
}

// GetNodeRevision_fieldTagsCountViaRevisionId Get NodeRevision_fieldTags via RevisionId
func GetNodeRevision_fieldTagsCountViaRevisionId(iRevisionId int) int64 {
	n, _ := Engine.Where("revision_id = ?", iRevisionId).Count(&NodeRevision_fieldTags{RevisionId: iRevisionId})
	return n
}

// GetNodeRevision_fieldTagsCountViaLangcode Get NodeRevision_fieldTags via Langcode
func GetNodeRevision_fieldTagsCountViaLangcode(iLangcode string) int64 {
	n, _ := Engine.Where("langcode = ?", iLangcode).Count(&NodeRevision_fieldTags{Langcode: iLangcode})
	return n
}

// GetNodeRevision_fieldTagsCountViaDelta Get NodeRevision_fieldTags via Delta
func GetNodeRevision_fieldTagsCountViaDelta(iDelta int) int64 {
	n, _ := Engine.Where("delta = ?", iDelta).Count(&NodeRevision_fieldTags{Delta: iDelta})
	return n
}

// GetNodeRevision_fieldTagsCountViaFieldTagsTargetId Get NodeRevision_fieldTags via FieldTagsTargetId
func GetNodeRevision_fieldTagsCountViaFieldTagsTargetId(iFieldTagsTargetId int) int64 {
	n, _ := Engine.Where("field_tags_target_id = ?", iFieldTagsTargetId).Count(&NodeRevision_fieldTags{FieldTagsTargetId: iFieldTagsTargetId})
	return n
}

// GetNodeRevision_fieldTagsesByBundleAndDeletedAndEntityId Get NodeRevision_fieldTagses via BundleAndDeletedAndEntityId
func GetNodeRevision_fieldTagsesByBundleAndDeletedAndEntityId(offset int, limit int, Bundle_ string, Deleted_ int, EntityId_ int) (*[]*NodeRevision_fieldTags, error) {
	var _NodeRevision_fieldTags = new([]*NodeRevision_fieldTags)
	err := Engine.Table("node_revision__field_tags").Where("bundle = ? and deleted = ? and entity_id = ?", Bundle_, Deleted_, EntityId_).Limit(limit, offset).Find(_NodeRevision_fieldTags)
	return _NodeRevision_fieldTags, err
}

// GetNodeRevision_fieldTagsesByBundleAndDeletedAndRevisionId Get NodeRevision_fieldTagses via BundleAndDeletedAndRevisionId
func GetNodeRevision_fieldTagsesByBundleAndDeletedAndRevisionId(offset int, limit int, Bundle_ string, Deleted_ int, RevisionId_ int) (*[]*NodeRevision_fieldTags, error) {
	var _NodeRevision_fieldTags = new([]*NodeRevision_fieldTags)
	err := Engine.Table("node_revision__field_tags").Where("bundle = ? and deleted = ? and revision_id = ?", Bundle_, Deleted_, RevisionId_).Limit(limit, offset).Find(_NodeRevision_fieldTags)
	return _NodeRevision_fieldTags, err
}

// GetNodeRevision_fieldTagsesByBundleAndDeletedAndLangcode Get NodeRevision_fieldTagses via BundleAndDeletedAndLangcode
func GetNodeRevision_fieldTagsesByBundleAndDeletedAndLangcode(offset int, limit int, Bundle_ string, Deleted_ int, Langcode_ string) (*[]*NodeRevision_fieldTags, error) {
	var _NodeRevision_fieldTags = new([]*NodeRevision_fieldTags)
	err := Engine.Table("node_revision__field_tags").Where("bundle = ? and deleted = ? and langcode = ?", Bundle_, Deleted_, Langcode_).Limit(limit, offset).Find(_NodeRevision_fieldTags)
	return _NodeRevision_fieldTags, err
}

// GetNodeRevision_fieldTagsesByBundleAndDeletedAndDelta Get NodeRevision_fieldTagses via BundleAndDeletedAndDelta
func GetNodeRevision_fieldTagsesByBundleAndDeletedAndDelta(offset int, limit int, Bundle_ string, Deleted_ int, Delta_ int) (*[]*NodeRevision_fieldTags, error) {
	var _NodeRevision_fieldTags = new([]*NodeRevision_fieldTags)
	err := Engine.Table("node_revision__field_tags").Where("bundle = ? and deleted = ? and delta = ?", Bundle_, Deleted_, Delta_).Limit(limit, offset).Find(_NodeRevision_fieldTags)
	return _NodeRevision_fieldTags, err
}

// GetNodeRevision_fieldTagsesByBundleAndDeletedAndFieldTagsTargetId Get NodeRevision_fieldTagses via BundleAndDeletedAndFieldTagsTargetId
func GetNodeRevision_fieldTagsesByBundleAndDeletedAndFieldTagsTargetId(offset int, limit int, Bundle_ string, Deleted_ int, FieldTagsTargetId_ int) (*[]*NodeRevision_fieldTags, error) {
	var _NodeRevision_fieldTags = new([]*NodeRevision_fieldTags)
	err := Engine.Table("node_revision__field_tags").Where("bundle = ? and deleted = ? and field_tags_target_id = ?", Bundle_, Deleted_, FieldTagsTargetId_).Limit(limit, offset).Find(_NodeRevision_fieldTags)
	return _NodeRevision_fieldTags, err
}

// GetNodeRevision_fieldTagsesByBundleAndEntityIdAndRevisionId Get NodeRevision_fieldTagses via BundleAndEntityIdAndRevisionId
func GetNodeRevision_fieldTagsesByBundleAndEntityIdAndRevisionId(offset int, limit int, Bundle_ string, EntityId_ int, RevisionId_ int) (*[]*NodeRevision_fieldTags, error) {
	var _NodeRevision_fieldTags = new([]*NodeRevision_fieldTags)
	err := Engine.Table("node_revision__field_tags").Where("bundle = ? and entity_id = ? and revision_id = ?", Bundle_, EntityId_, RevisionId_).Limit(limit, offset).Find(_NodeRevision_fieldTags)
	return _NodeRevision_fieldTags, err
}

// GetNodeRevision_fieldTagsesByBundleAndEntityIdAndLangcode Get NodeRevision_fieldTagses via BundleAndEntityIdAndLangcode
func GetNodeRevision_fieldTagsesByBundleAndEntityIdAndLangcode(offset int, limit int, Bundle_ string, EntityId_ int, Langcode_ string) (*[]*NodeRevision_fieldTags, error) {
	var _NodeRevision_fieldTags = new([]*NodeRevision_fieldTags)
	err := Engine.Table("node_revision__field_tags").Where("bundle = ? and entity_id = ? and langcode = ?", Bundle_, EntityId_, Langcode_).Limit(limit, offset).Find(_NodeRevision_fieldTags)
	return _NodeRevision_fieldTags, err
}

// GetNodeRevision_fieldTagsesByBundleAndEntityIdAndDelta Get NodeRevision_fieldTagses via BundleAndEntityIdAndDelta
func GetNodeRevision_fieldTagsesByBundleAndEntityIdAndDelta(offset int, limit int, Bundle_ string, EntityId_ int, Delta_ int) (*[]*NodeRevision_fieldTags, error) {
	var _NodeRevision_fieldTags = new([]*NodeRevision_fieldTags)
	err := Engine.Table("node_revision__field_tags").Where("bundle = ? and entity_id = ? and delta = ?", Bundle_, EntityId_, Delta_).Limit(limit, offset).Find(_NodeRevision_fieldTags)
	return _NodeRevision_fieldTags, err
}

// GetNodeRevision_fieldTagsesByBundleAndEntityIdAndFieldTagsTargetId Get NodeRevision_fieldTagses via BundleAndEntityIdAndFieldTagsTargetId
func GetNodeRevision_fieldTagsesByBundleAndEntityIdAndFieldTagsTargetId(offset int, limit int, Bundle_ string, EntityId_ int, FieldTagsTargetId_ int) (*[]*NodeRevision_fieldTags, error) {
	var _NodeRevision_fieldTags = new([]*NodeRevision_fieldTags)
	err := Engine.Table("node_revision__field_tags").Where("bundle = ? and entity_id = ? and field_tags_target_id = ?", Bundle_, EntityId_, FieldTagsTargetId_).Limit(limit, offset).Find(_NodeRevision_fieldTags)
	return _NodeRevision_fieldTags, err
}

// GetNodeRevision_fieldTagsesByBundleAndRevisionIdAndLangcode Get NodeRevision_fieldTagses via BundleAndRevisionIdAndLangcode
func GetNodeRevision_fieldTagsesByBundleAndRevisionIdAndLangcode(offset int, limit int, Bundle_ string, RevisionId_ int, Langcode_ string) (*[]*NodeRevision_fieldTags, error) {
	var _NodeRevision_fieldTags = new([]*NodeRevision_fieldTags)
	err := Engine.Table("node_revision__field_tags").Where("bundle = ? and revision_id = ? and langcode = ?", Bundle_, RevisionId_, Langcode_).Limit(limit, offset).Find(_NodeRevision_fieldTags)
	return _NodeRevision_fieldTags, err
}

// GetNodeRevision_fieldTagsesByBundleAndRevisionIdAndDelta Get NodeRevision_fieldTagses via BundleAndRevisionIdAndDelta
func GetNodeRevision_fieldTagsesByBundleAndRevisionIdAndDelta(offset int, limit int, Bundle_ string, RevisionId_ int, Delta_ int) (*[]*NodeRevision_fieldTags, error) {
	var _NodeRevision_fieldTags = new([]*NodeRevision_fieldTags)
	err := Engine.Table("node_revision__field_tags").Where("bundle = ? and revision_id = ? and delta = ?", Bundle_, RevisionId_, Delta_).Limit(limit, offset).Find(_NodeRevision_fieldTags)
	return _NodeRevision_fieldTags, err
}

// GetNodeRevision_fieldTagsesByBundleAndRevisionIdAndFieldTagsTargetId Get NodeRevision_fieldTagses via BundleAndRevisionIdAndFieldTagsTargetId
func GetNodeRevision_fieldTagsesByBundleAndRevisionIdAndFieldTagsTargetId(offset int, limit int, Bundle_ string, RevisionId_ int, FieldTagsTargetId_ int) (*[]*NodeRevision_fieldTags, error) {
	var _NodeRevision_fieldTags = new([]*NodeRevision_fieldTags)
	err := Engine.Table("node_revision__field_tags").Where("bundle = ? and revision_id = ? and field_tags_target_id = ?", Bundle_, RevisionId_, FieldTagsTargetId_).Limit(limit, offset).Find(_NodeRevision_fieldTags)
	return _NodeRevision_fieldTags, err
}

// GetNodeRevision_fieldTagsesByBundleAndLangcodeAndDelta Get NodeRevision_fieldTagses via BundleAndLangcodeAndDelta
func GetNodeRevision_fieldTagsesByBundleAndLangcodeAndDelta(offset int, limit int, Bundle_ string, Langcode_ string, Delta_ int) (*[]*NodeRevision_fieldTags, error) {
	var _NodeRevision_fieldTags = new([]*NodeRevision_fieldTags)
	err := Engine.Table("node_revision__field_tags").Where("bundle = ? and langcode = ? and delta = ?", Bundle_, Langcode_, Delta_).Limit(limit, offset).Find(_NodeRevision_fieldTags)
	return _NodeRevision_fieldTags, err
}

// GetNodeRevision_fieldTagsesByBundleAndLangcodeAndFieldTagsTargetId Get NodeRevision_fieldTagses via BundleAndLangcodeAndFieldTagsTargetId
func GetNodeRevision_fieldTagsesByBundleAndLangcodeAndFieldTagsTargetId(offset int, limit int, Bundle_ string, Langcode_ string, FieldTagsTargetId_ int) (*[]*NodeRevision_fieldTags, error) {
	var _NodeRevision_fieldTags = new([]*NodeRevision_fieldTags)
	err := Engine.Table("node_revision__field_tags").Where("bundle = ? and langcode = ? and field_tags_target_id = ?", Bundle_, Langcode_, FieldTagsTargetId_).Limit(limit, offset).Find(_NodeRevision_fieldTags)
	return _NodeRevision_fieldTags, err
}

// GetNodeRevision_fieldTagsesByBundleAndDeltaAndFieldTagsTargetId Get NodeRevision_fieldTagses via BundleAndDeltaAndFieldTagsTargetId
func GetNodeRevision_fieldTagsesByBundleAndDeltaAndFieldTagsTargetId(offset int, limit int, Bundle_ string, Delta_ int, FieldTagsTargetId_ int) (*[]*NodeRevision_fieldTags, error) {
	var _NodeRevision_fieldTags = new([]*NodeRevision_fieldTags)
	err := Engine.Table("node_revision__field_tags").Where("bundle = ? and delta = ? and field_tags_target_id = ?", Bundle_, Delta_, FieldTagsTargetId_).Limit(limit, offset).Find(_NodeRevision_fieldTags)
	return _NodeRevision_fieldTags, err
}

// GetNodeRevision_fieldTagsesByDeletedAndEntityIdAndRevisionId Get NodeRevision_fieldTagses via DeletedAndEntityIdAndRevisionId
func GetNodeRevision_fieldTagsesByDeletedAndEntityIdAndRevisionId(offset int, limit int, Deleted_ int, EntityId_ int, RevisionId_ int) (*[]*NodeRevision_fieldTags, error) {
	var _NodeRevision_fieldTags = new([]*NodeRevision_fieldTags)
	err := Engine.Table("node_revision__field_tags").Where("deleted = ? and entity_id = ? and revision_id = ?", Deleted_, EntityId_, RevisionId_).Limit(limit, offset).Find(_NodeRevision_fieldTags)
	return _NodeRevision_fieldTags, err
}

// GetNodeRevision_fieldTagsesByDeletedAndEntityIdAndLangcode Get NodeRevision_fieldTagses via DeletedAndEntityIdAndLangcode
func GetNodeRevision_fieldTagsesByDeletedAndEntityIdAndLangcode(offset int, limit int, Deleted_ int, EntityId_ int, Langcode_ string) (*[]*NodeRevision_fieldTags, error) {
	var _NodeRevision_fieldTags = new([]*NodeRevision_fieldTags)
	err := Engine.Table("node_revision__field_tags").Where("deleted = ? and entity_id = ? and langcode = ?", Deleted_, EntityId_, Langcode_).Limit(limit, offset).Find(_NodeRevision_fieldTags)
	return _NodeRevision_fieldTags, err
}

// GetNodeRevision_fieldTagsesByDeletedAndEntityIdAndDelta Get NodeRevision_fieldTagses via DeletedAndEntityIdAndDelta
func GetNodeRevision_fieldTagsesByDeletedAndEntityIdAndDelta(offset int, limit int, Deleted_ int, EntityId_ int, Delta_ int) (*[]*NodeRevision_fieldTags, error) {
	var _NodeRevision_fieldTags = new([]*NodeRevision_fieldTags)
	err := Engine.Table("node_revision__field_tags").Where("deleted = ? and entity_id = ? and delta = ?", Deleted_, EntityId_, Delta_).Limit(limit, offset).Find(_NodeRevision_fieldTags)
	return _NodeRevision_fieldTags, err
}

// GetNodeRevision_fieldTagsesByDeletedAndEntityIdAndFieldTagsTargetId Get NodeRevision_fieldTagses via DeletedAndEntityIdAndFieldTagsTargetId
func GetNodeRevision_fieldTagsesByDeletedAndEntityIdAndFieldTagsTargetId(offset int, limit int, Deleted_ int, EntityId_ int, FieldTagsTargetId_ int) (*[]*NodeRevision_fieldTags, error) {
	var _NodeRevision_fieldTags = new([]*NodeRevision_fieldTags)
	err := Engine.Table("node_revision__field_tags").Where("deleted = ? and entity_id = ? and field_tags_target_id = ?", Deleted_, EntityId_, FieldTagsTargetId_).Limit(limit, offset).Find(_NodeRevision_fieldTags)
	return _NodeRevision_fieldTags, err
}

// GetNodeRevision_fieldTagsesByDeletedAndRevisionIdAndLangcode Get NodeRevision_fieldTagses via DeletedAndRevisionIdAndLangcode
func GetNodeRevision_fieldTagsesByDeletedAndRevisionIdAndLangcode(offset int, limit int, Deleted_ int, RevisionId_ int, Langcode_ string) (*[]*NodeRevision_fieldTags, error) {
	var _NodeRevision_fieldTags = new([]*NodeRevision_fieldTags)
	err := Engine.Table("node_revision__field_tags").Where("deleted = ? and revision_id = ? and langcode = ?", Deleted_, RevisionId_, Langcode_).Limit(limit, offset).Find(_NodeRevision_fieldTags)
	return _NodeRevision_fieldTags, err
}

// GetNodeRevision_fieldTagsesByDeletedAndRevisionIdAndDelta Get NodeRevision_fieldTagses via DeletedAndRevisionIdAndDelta
func GetNodeRevision_fieldTagsesByDeletedAndRevisionIdAndDelta(offset int, limit int, Deleted_ int, RevisionId_ int, Delta_ int) (*[]*NodeRevision_fieldTags, error) {
	var _NodeRevision_fieldTags = new([]*NodeRevision_fieldTags)
	err := Engine.Table("node_revision__field_tags").Where("deleted = ? and revision_id = ? and delta = ?", Deleted_, RevisionId_, Delta_).Limit(limit, offset).Find(_NodeRevision_fieldTags)
	return _NodeRevision_fieldTags, err
}

// GetNodeRevision_fieldTagsesByDeletedAndRevisionIdAndFieldTagsTargetId Get NodeRevision_fieldTagses via DeletedAndRevisionIdAndFieldTagsTargetId
func GetNodeRevision_fieldTagsesByDeletedAndRevisionIdAndFieldTagsTargetId(offset int, limit int, Deleted_ int, RevisionId_ int, FieldTagsTargetId_ int) (*[]*NodeRevision_fieldTags, error) {
	var _NodeRevision_fieldTags = new([]*NodeRevision_fieldTags)
	err := Engine.Table("node_revision__field_tags").Where("deleted = ? and revision_id = ? and field_tags_target_id = ?", Deleted_, RevisionId_, FieldTagsTargetId_).Limit(limit, offset).Find(_NodeRevision_fieldTags)
	return _NodeRevision_fieldTags, err
}

// GetNodeRevision_fieldTagsesByDeletedAndLangcodeAndDelta Get NodeRevision_fieldTagses via DeletedAndLangcodeAndDelta
func GetNodeRevision_fieldTagsesByDeletedAndLangcodeAndDelta(offset int, limit int, Deleted_ int, Langcode_ string, Delta_ int) (*[]*NodeRevision_fieldTags, error) {
	var _NodeRevision_fieldTags = new([]*NodeRevision_fieldTags)
	err := Engine.Table("node_revision__field_tags").Where("deleted = ? and langcode = ? and delta = ?", Deleted_, Langcode_, Delta_).Limit(limit, offset).Find(_NodeRevision_fieldTags)
	return _NodeRevision_fieldTags, err
}

// GetNodeRevision_fieldTagsesByDeletedAndLangcodeAndFieldTagsTargetId Get NodeRevision_fieldTagses via DeletedAndLangcodeAndFieldTagsTargetId
func GetNodeRevision_fieldTagsesByDeletedAndLangcodeAndFieldTagsTargetId(offset int, limit int, Deleted_ int, Langcode_ string, FieldTagsTargetId_ int) (*[]*NodeRevision_fieldTags, error) {
	var _NodeRevision_fieldTags = new([]*NodeRevision_fieldTags)
	err := Engine.Table("node_revision__field_tags").Where("deleted = ? and langcode = ? and field_tags_target_id = ?", Deleted_, Langcode_, FieldTagsTargetId_).Limit(limit, offset).Find(_NodeRevision_fieldTags)
	return _NodeRevision_fieldTags, err
}

// GetNodeRevision_fieldTagsesByDeletedAndDeltaAndFieldTagsTargetId Get NodeRevision_fieldTagses via DeletedAndDeltaAndFieldTagsTargetId
func GetNodeRevision_fieldTagsesByDeletedAndDeltaAndFieldTagsTargetId(offset int, limit int, Deleted_ int, Delta_ int, FieldTagsTargetId_ int) (*[]*NodeRevision_fieldTags, error) {
	var _NodeRevision_fieldTags = new([]*NodeRevision_fieldTags)
	err := Engine.Table("node_revision__field_tags").Where("deleted = ? and delta = ? and field_tags_target_id = ?", Deleted_, Delta_, FieldTagsTargetId_).Limit(limit, offset).Find(_NodeRevision_fieldTags)
	return _NodeRevision_fieldTags, err
}

// GetNodeRevision_fieldTagsesByEntityIdAndRevisionIdAndLangcode Get NodeRevision_fieldTagses via EntityIdAndRevisionIdAndLangcode
func GetNodeRevision_fieldTagsesByEntityIdAndRevisionIdAndLangcode(offset int, limit int, EntityId_ int, RevisionId_ int, Langcode_ string) (*[]*NodeRevision_fieldTags, error) {
	var _NodeRevision_fieldTags = new([]*NodeRevision_fieldTags)
	err := Engine.Table("node_revision__field_tags").Where("entity_id = ? and revision_id = ? and langcode = ?", EntityId_, RevisionId_, Langcode_).Limit(limit, offset).Find(_NodeRevision_fieldTags)
	return _NodeRevision_fieldTags, err
}

// GetNodeRevision_fieldTagsesByEntityIdAndRevisionIdAndDelta Get NodeRevision_fieldTagses via EntityIdAndRevisionIdAndDelta
func GetNodeRevision_fieldTagsesByEntityIdAndRevisionIdAndDelta(offset int, limit int, EntityId_ int, RevisionId_ int, Delta_ int) (*[]*NodeRevision_fieldTags, error) {
	var _NodeRevision_fieldTags = new([]*NodeRevision_fieldTags)
	err := Engine.Table("node_revision__field_tags").Where("entity_id = ? and revision_id = ? and delta = ?", EntityId_, RevisionId_, Delta_).Limit(limit, offset).Find(_NodeRevision_fieldTags)
	return _NodeRevision_fieldTags, err
}

// GetNodeRevision_fieldTagsesByEntityIdAndRevisionIdAndFieldTagsTargetId Get NodeRevision_fieldTagses via EntityIdAndRevisionIdAndFieldTagsTargetId
func GetNodeRevision_fieldTagsesByEntityIdAndRevisionIdAndFieldTagsTargetId(offset int, limit int, EntityId_ int, RevisionId_ int, FieldTagsTargetId_ int) (*[]*NodeRevision_fieldTags, error) {
	var _NodeRevision_fieldTags = new([]*NodeRevision_fieldTags)
	err := Engine.Table("node_revision__field_tags").Where("entity_id = ? and revision_id = ? and field_tags_target_id = ?", EntityId_, RevisionId_, FieldTagsTargetId_).Limit(limit, offset).Find(_NodeRevision_fieldTags)
	return _NodeRevision_fieldTags, err
}

// GetNodeRevision_fieldTagsesByEntityIdAndLangcodeAndDelta Get NodeRevision_fieldTagses via EntityIdAndLangcodeAndDelta
func GetNodeRevision_fieldTagsesByEntityIdAndLangcodeAndDelta(offset int, limit int, EntityId_ int, Langcode_ string, Delta_ int) (*[]*NodeRevision_fieldTags, error) {
	var _NodeRevision_fieldTags = new([]*NodeRevision_fieldTags)
	err := Engine.Table("node_revision__field_tags").Where("entity_id = ? and langcode = ? and delta = ?", EntityId_, Langcode_, Delta_).Limit(limit, offset).Find(_NodeRevision_fieldTags)
	return _NodeRevision_fieldTags, err
}

// GetNodeRevision_fieldTagsesByEntityIdAndLangcodeAndFieldTagsTargetId Get NodeRevision_fieldTagses via EntityIdAndLangcodeAndFieldTagsTargetId
func GetNodeRevision_fieldTagsesByEntityIdAndLangcodeAndFieldTagsTargetId(offset int, limit int, EntityId_ int, Langcode_ string, FieldTagsTargetId_ int) (*[]*NodeRevision_fieldTags, error) {
	var _NodeRevision_fieldTags = new([]*NodeRevision_fieldTags)
	err := Engine.Table("node_revision__field_tags").Where("entity_id = ? and langcode = ? and field_tags_target_id = ?", EntityId_, Langcode_, FieldTagsTargetId_).Limit(limit, offset).Find(_NodeRevision_fieldTags)
	return _NodeRevision_fieldTags, err
}

// GetNodeRevision_fieldTagsesByEntityIdAndDeltaAndFieldTagsTargetId Get NodeRevision_fieldTagses via EntityIdAndDeltaAndFieldTagsTargetId
func GetNodeRevision_fieldTagsesByEntityIdAndDeltaAndFieldTagsTargetId(offset int, limit int, EntityId_ int, Delta_ int, FieldTagsTargetId_ int) (*[]*NodeRevision_fieldTags, error) {
	var _NodeRevision_fieldTags = new([]*NodeRevision_fieldTags)
	err := Engine.Table("node_revision__field_tags").Where("entity_id = ? and delta = ? and field_tags_target_id = ?", EntityId_, Delta_, FieldTagsTargetId_).Limit(limit, offset).Find(_NodeRevision_fieldTags)
	return _NodeRevision_fieldTags, err
}

// GetNodeRevision_fieldTagsesByRevisionIdAndLangcodeAndDelta Get NodeRevision_fieldTagses via RevisionIdAndLangcodeAndDelta
func GetNodeRevision_fieldTagsesByRevisionIdAndLangcodeAndDelta(offset int, limit int, RevisionId_ int, Langcode_ string, Delta_ int) (*[]*NodeRevision_fieldTags, error) {
	var _NodeRevision_fieldTags = new([]*NodeRevision_fieldTags)
	err := Engine.Table("node_revision__field_tags").Where("revision_id = ? and langcode = ? and delta = ?", RevisionId_, Langcode_, Delta_).Limit(limit, offset).Find(_NodeRevision_fieldTags)
	return _NodeRevision_fieldTags, err
}

// GetNodeRevision_fieldTagsesByRevisionIdAndLangcodeAndFieldTagsTargetId Get NodeRevision_fieldTagses via RevisionIdAndLangcodeAndFieldTagsTargetId
func GetNodeRevision_fieldTagsesByRevisionIdAndLangcodeAndFieldTagsTargetId(offset int, limit int, RevisionId_ int, Langcode_ string, FieldTagsTargetId_ int) (*[]*NodeRevision_fieldTags, error) {
	var _NodeRevision_fieldTags = new([]*NodeRevision_fieldTags)
	err := Engine.Table("node_revision__field_tags").Where("revision_id = ? and langcode = ? and field_tags_target_id = ?", RevisionId_, Langcode_, FieldTagsTargetId_).Limit(limit, offset).Find(_NodeRevision_fieldTags)
	return _NodeRevision_fieldTags, err
}

// GetNodeRevision_fieldTagsesByRevisionIdAndDeltaAndFieldTagsTargetId Get NodeRevision_fieldTagses via RevisionIdAndDeltaAndFieldTagsTargetId
func GetNodeRevision_fieldTagsesByRevisionIdAndDeltaAndFieldTagsTargetId(offset int, limit int, RevisionId_ int, Delta_ int, FieldTagsTargetId_ int) (*[]*NodeRevision_fieldTags, error) {
	var _NodeRevision_fieldTags = new([]*NodeRevision_fieldTags)
	err := Engine.Table("node_revision__field_tags").Where("revision_id = ? and delta = ? and field_tags_target_id = ?", RevisionId_, Delta_, FieldTagsTargetId_).Limit(limit, offset).Find(_NodeRevision_fieldTags)
	return _NodeRevision_fieldTags, err
}

// GetNodeRevision_fieldTagsesByLangcodeAndDeltaAndFieldTagsTargetId Get NodeRevision_fieldTagses via LangcodeAndDeltaAndFieldTagsTargetId
func GetNodeRevision_fieldTagsesByLangcodeAndDeltaAndFieldTagsTargetId(offset int, limit int, Langcode_ string, Delta_ int, FieldTagsTargetId_ int) (*[]*NodeRevision_fieldTags, error) {
	var _NodeRevision_fieldTags = new([]*NodeRevision_fieldTags)
	err := Engine.Table("node_revision__field_tags").Where("langcode = ? and delta = ? and field_tags_target_id = ?", Langcode_, Delta_, FieldTagsTargetId_).Limit(limit, offset).Find(_NodeRevision_fieldTags)
	return _NodeRevision_fieldTags, err
}

// GetNodeRevision_fieldTagsesByBundleAndDeleted Get NodeRevision_fieldTagses via BundleAndDeleted
func GetNodeRevision_fieldTagsesByBundleAndDeleted(offset int, limit int, Bundle_ string, Deleted_ int) (*[]*NodeRevision_fieldTags, error) {
	var _NodeRevision_fieldTags = new([]*NodeRevision_fieldTags)
	err := Engine.Table("node_revision__field_tags").Where("bundle = ? and deleted = ?", Bundle_, Deleted_).Limit(limit, offset).Find(_NodeRevision_fieldTags)
	return _NodeRevision_fieldTags, err
}

// GetNodeRevision_fieldTagsesByBundleAndEntityId Get NodeRevision_fieldTagses via BundleAndEntityId
func GetNodeRevision_fieldTagsesByBundleAndEntityId(offset int, limit int, Bundle_ string, EntityId_ int) (*[]*NodeRevision_fieldTags, error) {
	var _NodeRevision_fieldTags = new([]*NodeRevision_fieldTags)
	err := Engine.Table("node_revision__field_tags").Where("bundle = ? and entity_id = ?", Bundle_, EntityId_).Limit(limit, offset).Find(_NodeRevision_fieldTags)
	return _NodeRevision_fieldTags, err
}

// GetNodeRevision_fieldTagsesByBundleAndRevisionId Get NodeRevision_fieldTagses via BundleAndRevisionId
func GetNodeRevision_fieldTagsesByBundleAndRevisionId(offset int, limit int, Bundle_ string, RevisionId_ int) (*[]*NodeRevision_fieldTags, error) {
	var _NodeRevision_fieldTags = new([]*NodeRevision_fieldTags)
	err := Engine.Table("node_revision__field_tags").Where("bundle = ? and revision_id = ?", Bundle_, RevisionId_).Limit(limit, offset).Find(_NodeRevision_fieldTags)
	return _NodeRevision_fieldTags, err
}

// GetNodeRevision_fieldTagsesByBundleAndLangcode Get NodeRevision_fieldTagses via BundleAndLangcode
func GetNodeRevision_fieldTagsesByBundleAndLangcode(offset int, limit int, Bundle_ string, Langcode_ string) (*[]*NodeRevision_fieldTags, error) {
	var _NodeRevision_fieldTags = new([]*NodeRevision_fieldTags)
	err := Engine.Table("node_revision__field_tags").Where("bundle = ? and langcode = ?", Bundle_, Langcode_).Limit(limit, offset).Find(_NodeRevision_fieldTags)
	return _NodeRevision_fieldTags, err
}

// GetNodeRevision_fieldTagsesByBundleAndDelta Get NodeRevision_fieldTagses via BundleAndDelta
func GetNodeRevision_fieldTagsesByBundleAndDelta(offset int, limit int, Bundle_ string, Delta_ int) (*[]*NodeRevision_fieldTags, error) {
	var _NodeRevision_fieldTags = new([]*NodeRevision_fieldTags)
	err := Engine.Table("node_revision__field_tags").Where("bundle = ? and delta = ?", Bundle_, Delta_).Limit(limit, offset).Find(_NodeRevision_fieldTags)
	return _NodeRevision_fieldTags, err
}

// GetNodeRevision_fieldTagsesByBundleAndFieldTagsTargetId Get NodeRevision_fieldTagses via BundleAndFieldTagsTargetId
func GetNodeRevision_fieldTagsesByBundleAndFieldTagsTargetId(offset int, limit int, Bundle_ string, FieldTagsTargetId_ int) (*[]*NodeRevision_fieldTags, error) {
	var _NodeRevision_fieldTags = new([]*NodeRevision_fieldTags)
	err := Engine.Table("node_revision__field_tags").Where("bundle = ? and field_tags_target_id = ?", Bundle_, FieldTagsTargetId_).Limit(limit, offset).Find(_NodeRevision_fieldTags)
	return _NodeRevision_fieldTags, err
}

// GetNodeRevision_fieldTagsesByDeletedAndEntityId Get NodeRevision_fieldTagses via DeletedAndEntityId
func GetNodeRevision_fieldTagsesByDeletedAndEntityId(offset int, limit int, Deleted_ int, EntityId_ int) (*[]*NodeRevision_fieldTags, error) {
	var _NodeRevision_fieldTags = new([]*NodeRevision_fieldTags)
	err := Engine.Table("node_revision__field_tags").Where("deleted = ? and entity_id = ?", Deleted_, EntityId_).Limit(limit, offset).Find(_NodeRevision_fieldTags)
	return _NodeRevision_fieldTags, err
}

// GetNodeRevision_fieldTagsesByDeletedAndRevisionId Get NodeRevision_fieldTagses via DeletedAndRevisionId
func GetNodeRevision_fieldTagsesByDeletedAndRevisionId(offset int, limit int, Deleted_ int, RevisionId_ int) (*[]*NodeRevision_fieldTags, error) {
	var _NodeRevision_fieldTags = new([]*NodeRevision_fieldTags)
	err := Engine.Table("node_revision__field_tags").Where("deleted = ? and revision_id = ?", Deleted_, RevisionId_).Limit(limit, offset).Find(_NodeRevision_fieldTags)
	return _NodeRevision_fieldTags, err
}

// GetNodeRevision_fieldTagsesByDeletedAndLangcode Get NodeRevision_fieldTagses via DeletedAndLangcode
func GetNodeRevision_fieldTagsesByDeletedAndLangcode(offset int, limit int, Deleted_ int, Langcode_ string) (*[]*NodeRevision_fieldTags, error) {
	var _NodeRevision_fieldTags = new([]*NodeRevision_fieldTags)
	err := Engine.Table("node_revision__field_tags").Where("deleted = ? and langcode = ?", Deleted_, Langcode_).Limit(limit, offset).Find(_NodeRevision_fieldTags)
	return _NodeRevision_fieldTags, err
}

// GetNodeRevision_fieldTagsesByDeletedAndDelta Get NodeRevision_fieldTagses via DeletedAndDelta
func GetNodeRevision_fieldTagsesByDeletedAndDelta(offset int, limit int, Deleted_ int, Delta_ int) (*[]*NodeRevision_fieldTags, error) {
	var _NodeRevision_fieldTags = new([]*NodeRevision_fieldTags)
	err := Engine.Table("node_revision__field_tags").Where("deleted = ? and delta = ?", Deleted_, Delta_).Limit(limit, offset).Find(_NodeRevision_fieldTags)
	return _NodeRevision_fieldTags, err
}

// GetNodeRevision_fieldTagsesByDeletedAndFieldTagsTargetId Get NodeRevision_fieldTagses via DeletedAndFieldTagsTargetId
func GetNodeRevision_fieldTagsesByDeletedAndFieldTagsTargetId(offset int, limit int, Deleted_ int, FieldTagsTargetId_ int) (*[]*NodeRevision_fieldTags, error) {
	var _NodeRevision_fieldTags = new([]*NodeRevision_fieldTags)
	err := Engine.Table("node_revision__field_tags").Where("deleted = ? and field_tags_target_id = ?", Deleted_, FieldTagsTargetId_).Limit(limit, offset).Find(_NodeRevision_fieldTags)
	return _NodeRevision_fieldTags, err
}

// GetNodeRevision_fieldTagsesByEntityIdAndRevisionId Get NodeRevision_fieldTagses via EntityIdAndRevisionId
func GetNodeRevision_fieldTagsesByEntityIdAndRevisionId(offset int, limit int, EntityId_ int, RevisionId_ int) (*[]*NodeRevision_fieldTags, error) {
	var _NodeRevision_fieldTags = new([]*NodeRevision_fieldTags)
	err := Engine.Table("node_revision__field_tags").Where("entity_id = ? and revision_id = ?", EntityId_, RevisionId_).Limit(limit, offset).Find(_NodeRevision_fieldTags)
	return _NodeRevision_fieldTags, err
}

// GetNodeRevision_fieldTagsesByEntityIdAndLangcode Get NodeRevision_fieldTagses via EntityIdAndLangcode
func GetNodeRevision_fieldTagsesByEntityIdAndLangcode(offset int, limit int, EntityId_ int, Langcode_ string) (*[]*NodeRevision_fieldTags, error) {
	var _NodeRevision_fieldTags = new([]*NodeRevision_fieldTags)
	err := Engine.Table("node_revision__field_tags").Where("entity_id = ? and langcode = ?", EntityId_, Langcode_).Limit(limit, offset).Find(_NodeRevision_fieldTags)
	return _NodeRevision_fieldTags, err
}

// GetNodeRevision_fieldTagsesByEntityIdAndDelta Get NodeRevision_fieldTagses via EntityIdAndDelta
func GetNodeRevision_fieldTagsesByEntityIdAndDelta(offset int, limit int, EntityId_ int, Delta_ int) (*[]*NodeRevision_fieldTags, error) {
	var _NodeRevision_fieldTags = new([]*NodeRevision_fieldTags)
	err := Engine.Table("node_revision__field_tags").Where("entity_id = ? and delta = ?", EntityId_, Delta_).Limit(limit, offset).Find(_NodeRevision_fieldTags)
	return _NodeRevision_fieldTags, err
}

// GetNodeRevision_fieldTagsesByEntityIdAndFieldTagsTargetId Get NodeRevision_fieldTagses via EntityIdAndFieldTagsTargetId
func GetNodeRevision_fieldTagsesByEntityIdAndFieldTagsTargetId(offset int, limit int, EntityId_ int, FieldTagsTargetId_ int) (*[]*NodeRevision_fieldTags, error) {
	var _NodeRevision_fieldTags = new([]*NodeRevision_fieldTags)
	err := Engine.Table("node_revision__field_tags").Where("entity_id = ? and field_tags_target_id = ?", EntityId_, FieldTagsTargetId_).Limit(limit, offset).Find(_NodeRevision_fieldTags)
	return _NodeRevision_fieldTags, err
}

// GetNodeRevision_fieldTagsesByRevisionIdAndLangcode Get NodeRevision_fieldTagses via RevisionIdAndLangcode
func GetNodeRevision_fieldTagsesByRevisionIdAndLangcode(offset int, limit int, RevisionId_ int, Langcode_ string) (*[]*NodeRevision_fieldTags, error) {
	var _NodeRevision_fieldTags = new([]*NodeRevision_fieldTags)
	err := Engine.Table("node_revision__field_tags").Where("revision_id = ? and langcode = ?", RevisionId_, Langcode_).Limit(limit, offset).Find(_NodeRevision_fieldTags)
	return _NodeRevision_fieldTags, err
}

// GetNodeRevision_fieldTagsesByRevisionIdAndDelta Get NodeRevision_fieldTagses via RevisionIdAndDelta
func GetNodeRevision_fieldTagsesByRevisionIdAndDelta(offset int, limit int, RevisionId_ int, Delta_ int) (*[]*NodeRevision_fieldTags, error) {
	var _NodeRevision_fieldTags = new([]*NodeRevision_fieldTags)
	err := Engine.Table("node_revision__field_tags").Where("revision_id = ? and delta = ?", RevisionId_, Delta_).Limit(limit, offset).Find(_NodeRevision_fieldTags)
	return _NodeRevision_fieldTags, err
}

// GetNodeRevision_fieldTagsesByRevisionIdAndFieldTagsTargetId Get NodeRevision_fieldTagses via RevisionIdAndFieldTagsTargetId
func GetNodeRevision_fieldTagsesByRevisionIdAndFieldTagsTargetId(offset int, limit int, RevisionId_ int, FieldTagsTargetId_ int) (*[]*NodeRevision_fieldTags, error) {
	var _NodeRevision_fieldTags = new([]*NodeRevision_fieldTags)
	err := Engine.Table("node_revision__field_tags").Where("revision_id = ? and field_tags_target_id = ?", RevisionId_, FieldTagsTargetId_).Limit(limit, offset).Find(_NodeRevision_fieldTags)
	return _NodeRevision_fieldTags, err
}

// GetNodeRevision_fieldTagsesByLangcodeAndDelta Get NodeRevision_fieldTagses via LangcodeAndDelta
func GetNodeRevision_fieldTagsesByLangcodeAndDelta(offset int, limit int, Langcode_ string, Delta_ int) (*[]*NodeRevision_fieldTags, error) {
	var _NodeRevision_fieldTags = new([]*NodeRevision_fieldTags)
	err := Engine.Table("node_revision__field_tags").Where("langcode = ? and delta = ?", Langcode_, Delta_).Limit(limit, offset).Find(_NodeRevision_fieldTags)
	return _NodeRevision_fieldTags, err
}

// GetNodeRevision_fieldTagsesByLangcodeAndFieldTagsTargetId Get NodeRevision_fieldTagses via LangcodeAndFieldTagsTargetId
func GetNodeRevision_fieldTagsesByLangcodeAndFieldTagsTargetId(offset int, limit int, Langcode_ string, FieldTagsTargetId_ int) (*[]*NodeRevision_fieldTags, error) {
	var _NodeRevision_fieldTags = new([]*NodeRevision_fieldTags)
	err := Engine.Table("node_revision__field_tags").Where("langcode = ? and field_tags_target_id = ?", Langcode_, FieldTagsTargetId_).Limit(limit, offset).Find(_NodeRevision_fieldTags)
	return _NodeRevision_fieldTags, err
}

// GetNodeRevision_fieldTagsesByDeltaAndFieldTagsTargetId Get NodeRevision_fieldTagses via DeltaAndFieldTagsTargetId
func GetNodeRevision_fieldTagsesByDeltaAndFieldTagsTargetId(offset int, limit int, Delta_ int, FieldTagsTargetId_ int) (*[]*NodeRevision_fieldTags, error) {
	var _NodeRevision_fieldTags = new([]*NodeRevision_fieldTags)
	err := Engine.Table("node_revision__field_tags").Where("delta = ? and field_tags_target_id = ?", Delta_, FieldTagsTargetId_).Limit(limit, offset).Find(_NodeRevision_fieldTags)
	return _NodeRevision_fieldTags, err
}

// GetNodeRevision_fieldTagses Get NodeRevision_fieldTagses via field
func GetNodeRevision_fieldTagses(offset int, limit int, field string) (*[]*NodeRevision_fieldTags, error) {
	var _NodeRevision_fieldTags = new([]*NodeRevision_fieldTags)
	err := Engine.Table("node_revision__field_tags").Limit(limit, offset).Desc(field).Find(_NodeRevision_fieldTags)
	return _NodeRevision_fieldTags, err
}

// GetNodeRevision_fieldTagsesViaBundle Get NodeRevision_fieldTagss via Bundle
func GetNodeRevision_fieldTagsesViaBundle(offset int, limit int, Bundle_ string, field string) (*[]*NodeRevision_fieldTags, error) {
	var _NodeRevision_fieldTags = new([]*NodeRevision_fieldTags)
	err := Engine.Table("node_revision__field_tags").Where("bundle = ?", Bundle_).Limit(limit, offset).Desc(field).Find(_NodeRevision_fieldTags)
	return _NodeRevision_fieldTags, err
}

// GetNodeRevision_fieldTagsesViaDeleted Get NodeRevision_fieldTagss via Deleted
func GetNodeRevision_fieldTagsesViaDeleted(offset int, limit int, Deleted_ int, field string) (*[]*NodeRevision_fieldTags, error) {
	var _NodeRevision_fieldTags = new([]*NodeRevision_fieldTags)
	err := Engine.Table("node_revision__field_tags").Where("deleted = ?", Deleted_).Limit(limit, offset).Desc(field).Find(_NodeRevision_fieldTags)
	return _NodeRevision_fieldTags, err
}

// GetNodeRevision_fieldTagsesViaEntityId Get NodeRevision_fieldTagss via EntityId
func GetNodeRevision_fieldTagsesViaEntityId(offset int, limit int, EntityId_ int, field string) (*[]*NodeRevision_fieldTags, error) {
	var _NodeRevision_fieldTags = new([]*NodeRevision_fieldTags)
	err := Engine.Table("node_revision__field_tags").Where("entity_id = ?", EntityId_).Limit(limit, offset).Desc(field).Find(_NodeRevision_fieldTags)
	return _NodeRevision_fieldTags, err
}

// GetNodeRevision_fieldTagsesViaRevisionId Get NodeRevision_fieldTagss via RevisionId
func GetNodeRevision_fieldTagsesViaRevisionId(offset int, limit int, RevisionId_ int, field string) (*[]*NodeRevision_fieldTags, error) {
	var _NodeRevision_fieldTags = new([]*NodeRevision_fieldTags)
	err := Engine.Table("node_revision__field_tags").Where("revision_id = ?", RevisionId_).Limit(limit, offset).Desc(field).Find(_NodeRevision_fieldTags)
	return _NodeRevision_fieldTags, err
}

// GetNodeRevision_fieldTagsesViaLangcode Get NodeRevision_fieldTagss via Langcode
func GetNodeRevision_fieldTagsesViaLangcode(offset int, limit int, Langcode_ string, field string) (*[]*NodeRevision_fieldTags, error) {
	var _NodeRevision_fieldTags = new([]*NodeRevision_fieldTags)
	err := Engine.Table("node_revision__field_tags").Where("langcode = ?", Langcode_).Limit(limit, offset).Desc(field).Find(_NodeRevision_fieldTags)
	return _NodeRevision_fieldTags, err
}

// GetNodeRevision_fieldTagsesViaDelta Get NodeRevision_fieldTagss via Delta
func GetNodeRevision_fieldTagsesViaDelta(offset int, limit int, Delta_ int, field string) (*[]*NodeRevision_fieldTags, error) {
	var _NodeRevision_fieldTags = new([]*NodeRevision_fieldTags)
	err := Engine.Table("node_revision__field_tags").Where("delta = ?", Delta_).Limit(limit, offset).Desc(field).Find(_NodeRevision_fieldTags)
	return _NodeRevision_fieldTags, err
}

// GetNodeRevision_fieldTagsesViaFieldTagsTargetId Get NodeRevision_fieldTagss via FieldTagsTargetId
func GetNodeRevision_fieldTagsesViaFieldTagsTargetId(offset int, limit int, FieldTagsTargetId_ int, field string) (*[]*NodeRevision_fieldTags, error) {
	var _NodeRevision_fieldTags = new([]*NodeRevision_fieldTags)
	err := Engine.Table("node_revision__field_tags").Where("field_tags_target_id = ?", FieldTagsTargetId_).Limit(limit, offset).Desc(field).Find(_NodeRevision_fieldTags)
	return _NodeRevision_fieldTags, err
}

// HasNodeRevision_fieldTagsViaBundle Has NodeRevision_fieldTags via Bundle
func HasNodeRevision_fieldTagsViaBundle(iBundle string) bool {
	if has, err := Engine.Where("bundle = ?", iBundle).Get(new(NodeRevision_fieldTags)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasNodeRevision_fieldTagsViaDeleted Has NodeRevision_fieldTags via Deleted
func HasNodeRevision_fieldTagsViaDeleted(iDeleted int) bool {
	if has, err := Engine.Where("deleted = ?", iDeleted).Get(new(NodeRevision_fieldTags)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasNodeRevision_fieldTagsViaEntityId Has NodeRevision_fieldTags via EntityId
func HasNodeRevision_fieldTagsViaEntityId(iEntityId int) bool {
	if has, err := Engine.Where("entity_id = ?", iEntityId).Get(new(NodeRevision_fieldTags)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasNodeRevision_fieldTagsViaRevisionId Has NodeRevision_fieldTags via RevisionId
func HasNodeRevision_fieldTagsViaRevisionId(iRevisionId int) bool {
	if has, err := Engine.Where("revision_id = ?", iRevisionId).Get(new(NodeRevision_fieldTags)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasNodeRevision_fieldTagsViaLangcode Has NodeRevision_fieldTags via Langcode
func HasNodeRevision_fieldTagsViaLangcode(iLangcode string) bool {
	if has, err := Engine.Where("langcode = ?", iLangcode).Get(new(NodeRevision_fieldTags)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasNodeRevision_fieldTagsViaDelta Has NodeRevision_fieldTags via Delta
func HasNodeRevision_fieldTagsViaDelta(iDelta int) bool {
	if has, err := Engine.Where("delta = ?", iDelta).Get(new(NodeRevision_fieldTags)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasNodeRevision_fieldTagsViaFieldTagsTargetId Has NodeRevision_fieldTags via FieldTagsTargetId
func HasNodeRevision_fieldTagsViaFieldTagsTargetId(iFieldTagsTargetId int) bool {
	if has, err := Engine.Where("field_tags_target_id = ?", iFieldTagsTargetId).Get(new(NodeRevision_fieldTags)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// GetNodeRevision_fieldTagsViaBundle Get NodeRevision_fieldTags via Bundle
func GetNodeRevision_fieldTagsViaBundle(iBundle string) (*NodeRevision_fieldTags, error) {
	var _NodeRevision_fieldTags = &NodeRevision_fieldTags{Bundle: iBundle}
	has, err := Engine.Get(_NodeRevision_fieldTags)
	if has {
		return _NodeRevision_fieldTags, err
	} else {
		return nil, err
	}
}

// GetNodeRevision_fieldTagsViaDeleted Get NodeRevision_fieldTags via Deleted
func GetNodeRevision_fieldTagsViaDeleted(iDeleted int) (*NodeRevision_fieldTags, error) {
	var _NodeRevision_fieldTags = &NodeRevision_fieldTags{Deleted: iDeleted}
	has, err := Engine.Get(_NodeRevision_fieldTags)
	if has {
		return _NodeRevision_fieldTags, err
	} else {
		return nil, err
	}
}

// GetNodeRevision_fieldTagsViaEntityId Get NodeRevision_fieldTags via EntityId
func GetNodeRevision_fieldTagsViaEntityId(iEntityId int) (*NodeRevision_fieldTags, error) {
	var _NodeRevision_fieldTags = &NodeRevision_fieldTags{EntityId: iEntityId}
	has, err := Engine.Get(_NodeRevision_fieldTags)
	if has {
		return _NodeRevision_fieldTags, err
	} else {
		return nil, err
	}
}

// GetNodeRevision_fieldTagsViaRevisionId Get NodeRevision_fieldTags via RevisionId
func GetNodeRevision_fieldTagsViaRevisionId(iRevisionId int) (*NodeRevision_fieldTags, error) {
	var _NodeRevision_fieldTags = &NodeRevision_fieldTags{RevisionId: iRevisionId}
	has, err := Engine.Get(_NodeRevision_fieldTags)
	if has {
		return _NodeRevision_fieldTags, err
	} else {
		return nil, err
	}
}

// GetNodeRevision_fieldTagsViaLangcode Get NodeRevision_fieldTags via Langcode
func GetNodeRevision_fieldTagsViaLangcode(iLangcode string) (*NodeRevision_fieldTags, error) {
	var _NodeRevision_fieldTags = &NodeRevision_fieldTags{Langcode: iLangcode}
	has, err := Engine.Get(_NodeRevision_fieldTags)
	if has {
		return _NodeRevision_fieldTags, err
	} else {
		return nil, err
	}
}

// GetNodeRevision_fieldTagsViaDelta Get NodeRevision_fieldTags via Delta
func GetNodeRevision_fieldTagsViaDelta(iDelta int) (*NodeRevision_fieldTags, error) {
	var _NodeRevision_fieldTags = &NodeRevision_fieldTags{Delta: iDelta}
	has, err := Engine.Get(_NodeRevision_fieldTags)
	if has {
		return _NodeRevision_fieldTags, err
	} else {
		return nil, err
	}
}

// GetNodeRevision_fieldTagsViaFieldTagsTargetId Get NodeRevision_fieldTags via FieldTagsTargetId
func GetNodeRevision_fieldTagsViaFieldTagsTargetId(iFieldTagsTargetId int) (*NodeRevision_fieldTags, error) {
	var _NodeRevision_fieldTags = &NodeRevision_fieldTags{FieldTagsTargetId: iFieldTagsTargetId}
	has, err := Engine.Get(_NodeRevision_fieldTags)
	if has {
		return _NodeRevision_fieldTags, err
	} else {
		return nil, err
	}
}

// SetNodeRevision_fieldTagsViaBundle Set NodeRevision_fieldTags via Bundle
func SetNodeRevision_fieldTagsViaBundle(iBundle string, node_revision__field_tags *NodeRevision_fieldTags) (int64, error) {
	node_revision__field_tags.Bundle = iBundle
	return Engine.Insert(node_revision__field_tags)
}

// SetNodeRevision_fieldTagsViaDeleted Set NodeRevision_fieldTags via Deleted
func SetNodeRevision_fieldTagsViaDeleted(iDeleted int, node_revision__field_tags *NodeRevision_fieldTags) (int64, error) {
	node_revision__field_tags.Deleted = iDeleted
	return Engine.Insert(node_revision__field_tags)
}

// SetNodeRevision_fieldTagsViaEntityId Set NodeRevision_fieldTags via EntityId
func SetNodeRevision_fieldTagsViaEntityId(iEntityId int, node_revision__field_tags *NodeRevision_fieldTags) (int64, error) {
	node_revision__field_tags.EntityId = iEntityId
	return Engine.Insert(node_revision__field_tags)
}

// SetNodeRevision_fieldTagsViaRevisionId Set NodeRevision_fieldTags via RevisionId
func SetNodeRevision_fieldTagsViaRevisionId(iRevisionId int, node_revision__field_tags *NodeRevision_fieldTags) (int64, error) {
	node_revision__field_tags.RevisionId = iRevisionId
	return Engine.Insert(node_revision__field_tags)
}

// SetNodeRevision_fieldTagsViaLangcode Set NodeRevision_fieldTags via Langcode
func SetNodeRevision_fieldTagsViaLangcode(iLangcode string, node_revision__field_tags *NodeRevision_fieldTags) (int64, error) {
	node_revision__field_tags.Langcode = iLangcode
	return Engine.Insert(node_revision__field_tags)
}

// SetNodeRevision_fieldTagsViaDelta Set NodeRevision_fieldTags via Delta
func SetNodeRevision_fieldTagsViaDelta(iDelta int, node_revision__field_tags *NodeRevision_fieldTags) (int64, error) {
	node_revision__field_tags.Delta = iDelta
	return Engine.Insert(node_revision__field_tags)
}

// SetNodeRevision_fieldTagsViaFieldTagsTargetId Set NodeRevision_fieldTags via FieldTagsTargetId
func SetNodeRevision_fieldTagsViaFieldTagsTargetId(iFieldTagsTargetId int, node_revision__field_tags *NodeRevision_fieldTags) (int64, error) {
	node_revision__field_tags.FieldTagsTargetId = iFieldTagsTargetId
	return Engine.Insert(node_revision__field_tags)
}

// AddNodeRevision_fieldTags Add NodeRevision_fieldTags via all columns
func AddNodeRevision_fieldTags(iBundle string, iDeleted int, iEntityId int, iRevisionId int, iLangcode string, iDelta int, iFieldTagsTargetId int) error {
	_NodeRevision_fieldTags := &NodeRevision_fieldTags{Bundle: iBundle, Deleted: iDeleted, EntityId: iEntityId, RevisionId: iRevisionId, Langcode: iLangcode, Delta: iDelta, FieldTagsTargetId: iFieldTagsTargetId}
	if _, err := Engine.Insert(_NodeRevision_fieldTags); err != nil {
		return err
	} else {
		return nil
	}
}

// PostNodeRevision_fieldTags Post NodeRevision_fieldTags via iNodeRevision_fieldTags
func PostNodeRevision_fieldTags(iNodeRevision_fieldTags *NodeRevision_fieldTags) (string, error) {
	_, err := Engine.Insert(iNodeRevision_fieldTags)
	return iNodeRevision_fieldTags.Bundle, err
}

// PutNodeRevision_fieldTags Put NodeRevision_fieldTags
func PutNodeRevision_fieldTags(iNodeRevision_fieldTags *NodeRevision_fieldTags) (string, error) {
	_, err := Engine.Id(iNodeRevision_fieldTags.Bundle).Update(iNodeRevision_fieldTags)
	return iNodeRevision_fieldTags.Bundle, err
}

// PutNodeRevision_fieldTagsViaBundle Put NodeRevision_fieldTags via Bundle
func PutNodeRevision_fieldTagsViaBundle(Bundle_ string, iNodeRevision_fieldTags *NodeRevision_fieldTags) (int64, error) {
	row, err := Engine.Update(iNodeRevision_fieldTags, &NodeRevision_fieldTags{Bundle: Bundle_})
	return row, err
}

// PutNodeRevision_fieldTagsViaDeleted Put NodeRevision_fieldTags via Deleted
func PutNodeRevision_fieldTagsViaDeleted(Deleted_ int, iNodeRevision_fieldTags *NodeRevision_fieldTags) (int64, error) {
	row, err := Engine.Update(iNodeRevision_fieldTags, &NodeRevision_fieldTags{Deleted: Deleted_})
	return row, err
}

// PutNodeRevision_fieldTagsViaEntityId Put NodeRevision_fieldTags via EntityId
func PutNodeRevision_fieldTagsViaEntityId(EntityId_ int, iNodeRevision_fieldTags *NodeRevision_fieldTags) (int64, error) {
	row, err := Engine.Update(iNodeRevision_fieldTags, &NodeRevision_fieldTags{EntityId: EntityId_})
	return row, err
}

// PutNodeRevision_fieldTagsViaRevisionId Put NodeRevision_fieldTags via RevisionId
func PutNodeRevision_fieldTagsViaRevisionId(RevisionId_ int, iNodeRevision_fieldTags *NodeRevision_fieldTags) (int64, error) {
	row, err := Engine.Update(iNodeRevision_fieldTags, &NodeRevision_fieldTags{RevisionId: RevisionId_})
	return row, err
}

// PutNodeRevision_fieldTagsViaLangcode Put NodeRevision_fieldTags via Langcode
func PutNodeRevision_fieldTagsViaLangcode(Langcode_ string, iNodeRevision_fieldTags *NodeRevision_fieldTags) (int64, error) {
	row, err := Engine.Update(iNodeRevision_fieldTags, &NodeRevision_fieldTags{Langcode: Langcode_})
	return row, err
}

// PutNodeRevision_fieldTagsViaDelta Put NodeRevision_fieldTags via Delta
func PutNodeRevision_fieldTagsViaDelta(Delta_ int, iNodeRevision_fieldTags *NodeRevision_fieldTags) (int64, error) {
	row, err := Engine.Update(iNodeRevision_fieldTags, &NodeRevision_fieldTags{Delta: Delta_})
	return row, err
}

// PutNodeRevision_fieldTagsViaFieldTagsTargetId Put NodeRevision_fieldTags via FieldTagsTargetId
func PutNodeRevision_fieldTagsViaFieldTagsTargetId(FieldTagsTargetId_ int, iNodeRevision_fieldTags *NodeRevision_fieldTags) (int64, error) {
	row, err := Engine.Update(iNodeRevision_fieldTags, &NodeRevision_fieldTags{FieldTagsTargetId: FieldTagsTargetId_})
	return row, err
}

// UpdateNodeRevision_fieldTagsViaBundle via map[string]interface{}{}
func UpdateNodeRevision_fieldTagsViaBundle(iBundle string, iNodeRevision_fieldTagsMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(NodeRevision_fieldTags)).Where("bundle = ?", iBundle).Update(iNodeRevision_fieldTagsMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateNodeRevision_fieldTagsViaDeleted via map[string]interface{}{}
func UpdateNodeRevision_fieldTagsViaDeleted(iDeleted int, iNodeRevision_fieldTagsMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(NodeRevision_fieldTags)).Where("deleted = ?", iDeleted).Update(iNodeRevision_fieldTagsMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateNodeRevision_fieldTagsViaEntityId via map[string]interface{}{}
func UpdateNodeRevision_fieldTagsViaEntityId(iEntityId int, iNodeRevision_fieldTagsMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(NodeRevision_fieldTags)).Where("entity_id = ?", iEntityId).Update(iNodeRevision_fieldTagsMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateNodeRevision_fieldTagsViaRevisionId via map[string]interface{}{}
func UpdateNodeRevision_fieldTagsViaRevisionId(iRevisionId int, iNodeRevision_fieldTagsMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(NodeRevision_fieldTags)).Where("revision_id = ?", iRevisionId).Update(iNodeRevision_fieldTagsMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateNodeRevision_fieldTagsViaLangcode via map[string]interface{}{}
func UpdateNodeRevision_fieldTagsViaLangcode(iLangcode string, iNodeRevision_fieldTagsMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(NodeRevision_fieldTags)).Where("langcode = ?", iLangcode).Update(iNodeRevision_fieldTagsMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateNodeRevision_fieldTagsViaDelta via map[string]interface{}{}
func UpdateNodeRevision_fieldTagsViaDelta(iDelta int, iNodeRevision_fieldTagsMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(NodeRevision_fieldTags)).Where("delta = ?", iDelta).Update(iNodeRevision_fieldTagsMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateNodeRevision_fieldTagsViaFieldTagsTargetId via map[string]interface{}{}
func UpdateNodeRevision_fieldTagsViaFieldTagsTargetId(iFieldTagsTargetId int, iNodeRevision_fieldTagsMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(NodeRevision_fieldTags)).Where("field_tags_target_id = ?", iFieldTagsTargetId).Update(iNodeRevision_fieldTagsMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// DeleteNodeRevision_fieldTags Delete NodeRevision_fieldTags
func DeleteNodeRevision_fieldTags(iBundle string) (int64, error) {
	row, err := Engine.Id(iBundle).Delete(new(NodeRevision_fieldTags))
	return row, err
}

// DeleteNodeRevision_fieldTagsViaBundle Delete NodeRevision_fieldTags via Bundle
func DeleteNodeRevision_fieldTagsViaBundle(iBundle string) (err error) {
	var has bool
	var _NodeRevision_fieldTags = &NodeRevision_fieldTags{Bundle: iBundle}
	if has, err = Engine.Get(_NodeRevision_fieldTags); (has == true) && (err == nil) {
		if row, err := Engine.Where("bundle = ?", iBundle).Delete(new(NodeRevision_fieldTags)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteNodeRevision_fieldTagsViaDeleted Delete NodeRevision_fieldTags via Deleted
func DeleteNodeRevision_fieldTagsViaDeleted(iDeleted int) (err error) {
	var has bool
	var _NodeRevision_fieldTags = &NodeRevision_fieldTags{Deleted: iDeleted}
	if has, err = Engine.Get(_NodeRevision_fieldTags); (has == true) && (err == nil) {
		if row, err := Engine.Where("deleted = ?", iDeleted).Delete(new(NodeRevision_fieldTags)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteNodeRevision_fieldTagsViaEntityId Delete NodeRevision_fieldTags via EntityId
func DeleteNodeRevision_fieldTagsViaEntityId(iEntityId int) (err error) {
	var has bool
	var _NodeRevision_fieldTags = &NodeRevision_fieldTags{EntityId: iEntityId}
	if has, err = Engine.Get(_NodeRevision_fieldTags); (has == true) && (err == nil) {
		if row, err := Engine.Where("entity_id = ?", iEntityId).Delete(new(NodeRevision_fieldTags)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteNodeRevision_fieldTagsViaRevisionId Delete NodeRevision_fieldTags via RevisionId
func DeleteNodeRevision_fieldTagsViaRevisionId(iRevisionId int) (err error) {
	var has bool
	var _NodeRevision_fieldTags = &NodeRevision_fieldTags{RevisionId: iRevisionId}
	if has, err = Engine.Get(_NodeRevision_fieldTags); (has == true) && (err == nil) {
		if row, err := Engine.Where("revision_id = ?", iRevisionId).Delete(new(NodeRevision_fieldTags)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteNodeRevision_fieldTagsViaLangcode Delete NodeRevision_fieldTags via Langcode
func DeleteNodeRevision_fieldTagsViaLangcode(iLangcode string) (err error) {
	var has bool
	var _NodeRevision_fieldTags = &NodeRevision_fieldTags{Langcode: iLangcode}
	if has, err = Engine.Get(_NodeRevision_fieldTags); (has == true) && (err == nil) {
		if row, err := Engine.Where("langcode = ?", iLangcode).Delete(new(NodeRevision_fieldTags)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteNodeRevision_fieldTagsViaDelta Delete NodeRevision_fieldTags via Delta
func DeleteNodeRevision_fieldTagsViaDelta(iDelta int) (err error) {
	var has bool
	var _NodeRevision_fieldTags = &NodeRevision_fieldTags{Delta: iDelta}
	if has, err = Engine.Get(_NodeRevision_fieldTags); (has == true) && (err == nil) {
		if row, err := Engine.Where("delta = ?", iDelta).Delete(new(NodeRevision_fieldTags)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteNodeRevision_fieldTagsViaFieldTagsTargetId Delete NodeRevision_fieldTags via FieldTagsTargetId
func DeleteNodeRevision_fieldTagsViaFieldTagsTargetId(iFieldTagsTargetId int) (err error) {
	var has bool
	var _NodeRevision_fieldTags = &NodeRevision_fieldTags{FieldTagsTargetId: iFieldTagsTargetId}
	if has, err = Engine.Get(_NodeRevision_fieldTags); (has == true) && (err == nil) {
		if row, err := Engine.Where("field_tags_target_id = ?", iFieldTagsTargetId).Delete(new(NodeRevision_fieldTags)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}
