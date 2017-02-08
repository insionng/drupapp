package model

type Node_fieldTags struct {
	Bundle            string `xorm:"not null default '' index VARCHAR(128)"`
	Deleted           int    `xorm:"not null pk default 0 TINYINT(4)"`
	EntityId          int    `xorm:"not null pk INT(10)"`
	RevisionId        int    `xorm:"not null index INT(10)"`
	Langcode          string `xorm:"not null pk default '' VARCHAR(32)"`
	Delta             int    `xorm:"not null pk INT(10)"`
	FieldTagsTargetId int    `xorm:"not null index INT(10)"`
}

// GetNode_fieldTagsesCount Node_fieldTagss' Count
func GetNode_fieldTagsesCount(offset int, limit int) (int64, error) {
	total, err := Engine.Limit(limit, offset).Count(&Node_fieldTags{})
	return total, err
}

// GetNode_fieldTagsCountViaBundle Get Node_fieldTags via Bundle
func GetNode_fieldTagsCountViaBundle(iBundle string) int64 {
	n, _ := Engine.Where("bundle = ?", iBundle).Count(&Node_fieldTags{Bundle: iBundle})
	return n
}

// GetNode_fieldTagsCountViaDeleted Get Node_fieldTags via Deleted
func GetNode_fieldTagsCountViaDeleted(iDeleted int) int64 {
	n, _ := Engine.Where("deleted = ?", iDeleted).Count(&Node_fieldTags{Deleted: iDeleted})
	return n
}

// GetNode_fieldTagsCountViaEntityId Get Node_fieldTags via EntityId
func GetNode_fieldTagsCountViaEntityId(iEntityId int) int64 {
	n, _ := Engine.Where("entity_id = ?", iEntityId).Count(&Node_fieldTags{EntityId: iEntityId})
	return n
}

// GetNode_fieldTagsCountViaRevisionId Get Node_fieldTags via RevisionId
func GetNode_fieldTagsCountViaRevisionId(iRevisionId int) int64 {
	n, _ := Engine.Where("revision_id = ?", iRevisionId).Count(&Node_fieldTags{RevisionId: iRevisionId})
	return n
}

// GetNode_fieldTagsCountViaLangcode Get Node_fieldTags via Langcode
func GetNode_fieldTagsCountViaLangcode(iLangcode string) int64 {
	n, _ := Engine.Where("langcode = ?", iLangcode).Count(&Node_fieldTags{Langcode: iLangcode})
	return n
}

// GetNode_fieldTagsCountViaDelta Get Node_fieldTags via Delta
func GetNode_fieldTagsCountViaDelta(iDelta int) int64 {
	n, _ := Engine.Where("delta = ?", iDelta).Count(&Node_fieldTags{Delta: iDelta})
	return n
}

// GetNode_fieldTagsCountViaFieldTagsTargetId Get Node_fieldTags via FieldTagsTargetId
func GetNode_fieldTagsCountViaFieldTagsTargetId(iFieldTagsTargetId int) int64 {
	n, _ := Engine.Where("field_tags_target_id = ?", iFieldTagsTargetId).Count(&Node_fieldTags{FieldTagsTargetId: iFieldTagsTargetId})
	return n
}

// GetNode_fieldTagsesByBundleAndDeletedAndEntityId Get Node_fieldTagses via BundleAndDeletedAndEntityId
func GetNode_fieldTagsesByBundleAndDeletedAndEntityId(offset int, limit int, Bundle_ string, Deleted_ int, EntityId_ int) (*[]*Node_fieldTags, error) {
	var _Node_fieldTags = new([]*Node_fieldTags)
	err := Engine.Table("node__field_tags").Where("bundle = ? and deleted = ? and entity_id = ?", Bundle_, Deleted_, EntityId_).Limit(limit, offset).Find(_Node_fieldTags)
	return _Node_fieldTags, err
}

// GetNode_fieldTagsesByBundleAndDeletedAndRevisionId Get Node_fieldTagses via BundleAndDeletedAndRevisionId
func GetNode_fieldTagsesByBundleAndDeletedAndRevisionId(offset int, limit int, Bundle_ string, Deleted_ int, RevisionId_ int) (*[]*Node_fieldTags, error) {
	var _Node_fieldTags = new([]*Node_fieldTags)
	err := Engine.Table("node__field_tags").Where("bundle = ? and deleted = ? and revision_id = ?", Bundle_, Deleted_, RevisionId_).Limit(limit, offset).Find(_Node_fieldTags)
	return _Node_fieldTags, err
}

// GetNode_fieldTagsesByBundleAndDeletedAndLangcode Get Node_fieldTagses via BundleAndDeletedAndLangcode
func GetNode_fieldTagsesByBundleAndDeletedAndLangcode(offset int, limit int, Bundle_ string, Deleted_ int, Langcode_ string) (*[]*Node_fieldTags, error) {
	var _Node_fieldTags = new([]*Node_fieldTags)
	err := Engine.Table("node__field_tags").Where("bundle = ? and deleted = ? and langcode = ?", Bundle_, Deleted_, Langcode_).Limit(limit, offset).Find(_Node_fieldTags)
	return _Node_fieldTags, err
}

// GetNode_fieldTagsesByBundleAndDeletedAndDelta Get Node_fieldTagses via BundleAndDeletedAndDelta
func GetNode_fieldTagsesByBundleAndDeletedAndDelta(offset int, limit int, Bundle_ string, Deleted_ int, Delta_ int) (*[]*Node_fieldTags, error) {
	var _Node_fieldTags = new([]*Node_fieldTags)
	err := Engine.Table("node__field_tags").Where("bundle = ? and deleted = ? and delta = ?", Bundle_, Deleted_, Delta_).Limit(limit, offset).Find(_Node_fieldTags)
	return _Node_fieldTags, err
}

// GetNode_fieldTagsesByBundleAndDeletedAndFieldTagsTargetId Get Node_fieldTagses via BundleAndDeletedAndFieldTagsTargetId
func GetNode_fieldTagsesByBundleAndDeletedAndFieldTagsTargetId(offset int, limit int, Bundle_ string, Deleted_ int, FieldTagsTargetId_ int) (*[]*Node_fieldTags, error) {
	var _Node_fieldTags = new([]*Node_fieldTags)
	err := Engine.Table("node__field_tags").Where("bundle = ? and deleted = ? and field_tags_target_id = ?", Bundle_, Deleted_, FieldTagsTargetId_).Limit(limit, offset).Find(_Node_fieldTags)
	return _Node_fieldTags, err
}

// GetNode_fieldTagsesByBundleAndEntityIdAndRevisionId Get Node_fieldTagses via BundleAndEntityIdAndRevisionId
func GetNode_fieldTagsesByBundleAndEntityIdAndRevisionId(offset int, limit int, Bundle_ string, EntityId_ int, RevisionId_ int) (*[]*Node_fieldTags, error) {
	var _Node_fieldTags = new([]*Node_fieldTags)
	err := Engine.Table("node__field_tags").Where("bundle = ? and entity_id = ? and revision_id = ?", Bundle_, EntityId_, RevisionId_).Limit(limit, offset).Find(_Node_fieldTags)
	return _Node_fieldTags, err
}

// GetNode_fieldTagsesByBundleAndEntityIdAndLangcode Get Node_fieldTagses via BundleAndEntityIdAndLangcode
func GetNode_fieldTagsesByBundleAndEntityIdAndLangcode(offset int, limit int, Bundle_ string, EntityId_ int, Langcode_ string) (*[]*Node_fieldTags, error) {
	var _Node_fieldTags = new([]*Node_fieldTags)
	err := Engine.Table("node__field_tags").Where("bundle = ? and entity_id = ? and langcode = ?", Bundle_, EntityId_, Langcode_).Limit(limit, offset).Find(_Node_fieldTags)
	return _Node_fieldTags, err
}

// GetNode_fieldTagsesByBundleAndEntityIdAndDelta Get Node_fieldTagses via BundleAndEntityIdAndDelta
func GetNode_fieldTagsesByBundleAndEntityIdAndDelta(offset int, limit int, Bundle_ string, EntityId_ int, Delta_ int) (*[]*Node_fieldTags, error) {
	var _Node_fieldTags = new([]*Node_fieldTags)
	err := Engine.Table("node__field_tags").Where("bundle = ? and entity_id = ? and delta = ?", Bundle_, EntityId_, Delta_).Limit(limit, offset).Find(_Node_fieldTags)
	return _Node_fieldTags, err
}

// GetNode_fieldTagsesByBundleAndEntityIdAndFieldTagsTargetId Get Node_fieldTagses via BundleAndEntityIdAndFieldTagsTargetId
func GetNode_fieldTagsesByBundleAndEntityIdAndFieldTagsTargetId(offset int, limit int, Bundle_ string, EntityId_ int, FieldTagsTargetId_ int) (*[]*Node_fieldTags, error) {
	var _Node_fieldTags = new([]*Node_fieldTags)
	err := Engine.Table("node__field_tags").Where("bundle = ? and entity_id = ? and field_tags_target_id = ?", Bundle_, EntityId_, FieldTagsTargetId_).Limit(limit, offset).Find(_Node_fieldTags)
	return _Node_fieldTags, err
}

// GetNode_fieldTagsesByBundleAndRevisionIdAndLangcode Get Node_fieldTagses via BundleAndRevisionIdAndLangcode
func GetNode_fieldTagsesByBundleAndRevisionIdAndLangcode(offset int, limit int, Bundle_ string, RevisionId_ int, Langcode_ string) (*[]*Node_fieldTags, error) {
	var _Node_fieldTags = new([]*Node_fieldTags)
	err := Engine.Table("node__field_tags").Where("bundle = ? and revision_id = ? and langcode = ?", Bundle_, RevisionId_, Langcode_).Limit(limit, offset).Find(_Node_fieldTags)
	return _Node_fieldTags, err
}

// GetNode_fieldTagsesByBundleAndRevisionIdAndDelta Get Node_fieldTagses via BundleAndRevisionIdAndDelta
func GetNode_fieldTagsesByBundleAndRevisionIdAndDelta(offset int, limit int, Bundle_ string, RevisionId_ int, Delta_ int) (*[]*Node_fieldTags, error) {
	var _Node_fieldTags = new([]*Node_fieldTags)
	err := Engine.Table("node__field_tags").Where("bundle = ? and revision_id = ? and delta = ?", Bundle_, RevisionId_, Delta_).Limit(limit, offset).Find(_Node_fieldTags)
	return _Node_fieldTags, err
}

// GetNode_fieldTagsesByBundleAndRevisionIdAndFieldTagsTargetId Get Node_fieldTagses via BundleAndRevisionIdAndFieldTagsTargetId
func GetNode_fieldTagsesByBundleAndRevisionIdAndFieldTagsTargetId(offset int, limit int, Bundle_ string, RevisionId_ int, FieldTagsTargetId_ int) (*[]*Node_fieldTags, error) {
	var _Node_fieldTags = new([]*Node_fieldTags)
	err := Engine.Table("node__field_tags").Where("bundle = ? and revision_id = ? and field_tags_target_id = ?", Bundle_, RevisionId_, FieldTagsTargetId_).Limit(limit, offset).Find(_Node_fieldTags)
	return _Node_fieldTags, err
}

// GetNode_fieldTagsesByBundleAndLangcodeAndDelta Get Node_fieldTagses via BundleAndLangcodeAndDelta
func GetNode_fieldTagsesByBundleAndLangcodeAndDelta(offset int, limit int, Bundle_ string, Langcode_ string, Delta_ int) (*[]*Node_fieldTags, error) {
	var _Node_fieldTags = new([]*Node_fieldTags)
	err := Engine.Table("node__field_tags").Where("bundle = ? and langcode = ? and delta = ?", Bundle_, Langcode_, Delta_).Limit(limit, offset).Find(_Node_fieldTags)
	return _Node_fieldTags, err
}

// GetNode_fieldTagsesByBundleAndLangcodeAndFieldTagsTargetId Get Node_fieldTagses via BundleAndLangcodeAndFieldTagsTargetId
func GetNode_fieldTagsesByBundleAndLangcodeAndFieldTagsTargetId(offset int, limit int, Bundle_ string, Langcode_ string, FieldTagsTargetId_ int) (*[]*Node_fieldTags, error) {
	var _Node_fieldTags = new([]*Node_fieldTags)
	err := Engine.Table("node__field_tags").Where("bundle = ? and langcode = ? and field_tags_target_id = ?", Bundle_, Langcode_, FieldTagsTargetId_).Limit(limit, offset).Find(_Node_fieldTags)
	return _Node_fieldTags, err
}

// GetNode_fieldTagsesByBundleAndDeltaAndFieldTagsTargetId Get Node_fieldTagses via BundleAndDeltaAndFieldTagsTargetId
func GetNode_fieldTagsesByBundleAndDeltaAndFieldTagsTargetId(offset int, limit int, Bundle_ string, Delta_ int, FieldTagsTargetId_ int) (*[]*Node_fieldTags, error) {
	var _Node_fieldTags = new([]*Node_fieldTags)
	err := Engine.Table("node__field_tags").Where("bundle = ? and delta = ? and field_tags_target_id = ?", Bundle_, Delta_, FieldTagsTargetId_).Limit(limit, offset).Find(_Node_fieldTags)
	return _Node_fieldTags, err
}

// GetNode_fieldTagsesByDeletedAndEntityIdAndRevisionId Get Node_fieldTagses via DeletedAndEntityIdAndRevisionId
func GetNode_fieldTagsesByDeletedAndEntityIdAndRevisionId(offset int, limit int, Deleted_ int, EntityId_ int, RevisionId_ int) (*[]*Node_fieldTags, error) {
	var _Node_fieldTags = new([]*Node_fieldTags)
	err := Engine.Table("node__field_tags").Where("deleted = ? and entity_id = ? and revision_id = ?", Deleted_, EntityId_, RevisionId_).Limit(limit, offset).Find(_Node_fieldTags)
	return _Node_fieldTags, err
}

// GetNode_fieldTagsesByDeletedAndEntityIdAndLangcode Get Node_fieldTagses via DeletedAndEntityIdAndLangcode
func GetNode_fieldTagsesByDeletedAndEntityIdAndLangcode(offset int, limit int, Deleted_ int, EntityId_ int, Langcode_ string) (*[]*Node_fieldTags, error) {
	var _Node_fieldTags = new([]*Node_fieldTags)
	err := Engine.Table("node__field_tags").Where("deleted = ? and entity_id = ? and langcode = ?", Deleted_, EntityId_, Langcode_).Limit(limit, offset).Find(_Node_fieldTags)
	return _Node_fieldTags, err
}

// GetNode_fieldTagsesByDeletedAndEntityIdAndDelta Get Node_fieldTagses via DeletedAndEntityIdAndDelta
func GetNode_fieldTagsesByDeletedAndEntityIdAndDelta(offset int, limit int, Deleted_ int, EntityId_ int, Delta_ int) (*[]*Node_fieldTags, error) {
	var _Node_fieldTags = new([]*Node_fieldTags)
	err := Engine.Table("node__field_tags").Where("deleted = ? and entity_id = ? and delta = ?", Deleted_, EntityId_, Delta_).Limit(limit, offset).Find(_Node_fieldTags)
	return _Node_fieldTags, err
}

// GetNode_fieldTagsesByDeletedAndEntityIdAndFieldTagsTargetId Get Node_fieldTagses via DeletedAndEntityIdAndFieldTagsTargetId
func GetNode_fieldTagsesByDeletedAndEntityIdAndFieldTagsTargetId(offset int, limit int, Deleted_ int, EntityId_ int, FieldTagsTargetId_ int) (*[]*Node_fieldTags, error) {
	var _Node_fieldTags = new([]*Node_fieldTags)
	err := Engine.Table("node__field_tags").Where("deleted = ? and entity_id = ? and field_tags_target_id = ?", Deleted_, EntityId_, FieldTagsTargetId_).Limit(limit, offset).Find(_Node_fieldTags)
	return _Node_fieldTags, err
}

// GetNode_fieldTagsesByDeletedAndRevisionIdAndLangcode Get Node_fieldTagses via DeletedAndRevisionIdAndLangcode
func GetNode_fieldTagsesByDeletedAndRevisionIdAndLangcode(offset int, limit int, Deleted_ int, RevisionId_ int, Langcode_ string) (*[]*Node_fieldTags, error) {
	var _Node_fieldTags = new([]*Node_fieldTags)
	err := Engine.Table("node__field_tags").Where("deleted = ? and revision_id = ? and langcode = ?", Deleted_, RevisionId_, Langcode_).Limit(limit, offset).Find(_Node_fieldTags)
	return _Node_fieldTags, err
}

// GetNode_fieldTagsesByDeletedAndRevisionIdAndDelta Get Node_fieldTagses via DeletedAndRevisionIdAndDelta
func GetNode_fieldTagsesByDeletedAndRevisionIdAndDelta(offset int, limit int, Deleted_ int, RevisionId_ int, Delta_ int) (*[]*Node_fieldTags, error) {
	var _Node_fieldTags = new([]*Node_fieldTags)
	err := Engine.Table("node__field_tags").Where("deleted = ? and revision_id = ? and delta = ?", Deleted_, RevisionId_, Delta_).Limit(limit, offset).Find(_Node_fieldTags)
	return _Node_fieldTags, err
}

// GetNode_fieldTagsesByDeletedAndRevisionIdAndFieldTagsTargetId Get Node_fieldTagses via DeletedAndRevisionIdAndFieldTagsTargetId
func GetNode_fieldTagsesByDeletedAndRevisionIdAndFieldTagsTargetId(offset int, limit int, Deleted_ int, RevisionId_ int, FieldTagsTargetId_ int) (*[]*Node_fieldTags, error) {
	var _Node_fieldTags = new([]*Node_fieldTags)
	err := Engine.Table("node__field_tags").Where("deleted = ? and revision_id = ? and field_tags_target_id = ?", Deleted_, RevisionId_, FieldTagsTargetId_).Limit(limit, offset).Find(_Node_fieldTags)
	return _Node_fieldTags, err
}

// GetNode_fieldTagsesByDeletedAndLangcodeAndDelta Get Node_fieldTagses via DeletedAndLangcodeAndDelta
func GetNode_fieldTagsesByDeletedAndLangcodeAndDelta(offset int, limit int, Deleted_ int, Langcode_ string, Delta_ int) (*[]*Node_fieldTags, error) {
	var _Node_fieldTags = new([]*Node_fieldTags)
	err := Engine.Table("node__field_tags").Where("deleted = ? and langcode = ? and delta = ?", Deleted_, Langcode_, Delta_).Limit(limit, offset).Find(_Node_fieldTags)
	return _Node_fieldTags, err
}

// GetNode_fieldTagsesByDeletedAndLangcodeAndFieldTagsTargetId Get Node_fieldTagses via DeletedAndLangcodeAndFieldTagsTargetId
func GetNode_fieldTagsesByDeletedAndLangcodeAndFieldTagsTargetId(offset int, limit int, Deleted_ int, Langcode_ string, FieldTagsTargetId_ int) (*[]*Node_fieldTags, error) {
	var _Node_fieldTags = new([]*Node_fieldTags)
	err := Engine.Table("node__field_tags").Where("deleted = ? and langcode = ? and field_tags_target_id = ?", Deleted_, Langcode_, FieldTagsTargetId_).Limit(limit, offset).Find(_Node_fieldTags)
	return _Node_fieldTags, err
}

// GetNode_fieldTagsesByDeletedAndDeltaAndFieldTagsTargetId Get Node_fieldTagses via DeletedAndDeltaAndFieldTagsTargetId
func GetNode_fieldTagsesByDeletedAndDeltaAndFieldTagsTargetId(offset int, limit int, Deleted_ int, Delta_ int, FieldTagsTargetId_ int) (*[]*Node_fieldTags, error) {
	var _Node_fieldTags = new([]*Node_fieldTags)
	err := Engine.Table("node__field_tags").Where("deleted = ? and delta = ? and field_tags_target_id = ?", Deleted_, Delta_, FieldTagsTargetId_).Limit(limit, offset).Find(_Node_fieldTags)
	return _Node_fieldTags, err
}

// GetNode_fieldTagsesByEntityIdAndRevisionIdAndLangcode Get Node_fieldTagses via EntityIdAndRevisionIdAndLangcode
func GetNode_fieldTagsesByEntityIdAndRevisionIdAndLangcode(offset int, limit int, EntityId_ int, RevisionId_ int, Langcode_ string) (*[]*Node_fieldTags, error) {
	var _Node_fieldTags = new([]*Node_fieldTags)
	err := Engine.Table("node__field_tags").Where("entity_id = ? and revision_id = ? and langcode = ?", EntityId_, RevisionId_, Langcode_).Limit(limit, offset).Find(_Node_fieldTags)
	return _Node_fieldTags, err
}

// GetNode_fieldTagsesByEntityIdAndRevisionIdAndDelta Get Node_fieldTagses via EntityIdAndRevisionIdAndDelta
func GetNode_fieldTagsesByEntityIdAndRevisionIdAndDelta(offset int, limit int, EntityId_ int, RevisionId_ int, Delta_ int) (*[]*Node_fieldTags, error) {
	var _Node_fieldTags = new([]*Node_fieldTags)
	err := Engine.Table("node__field_tags").Where("entity_id = ? and revision_id = ? and delta = ?", EntityId_, RevisionId_, Delta_).Limit(limit, offset).Find(_Node_fieldTags)
	return _Node_fieldTags, err
}

// GetNode_fieldTagsesByEntityIdAndRevisionIdAndFieldTagsTargetId Get Node_fieldTagses via EntityIdAndRevisionIdAndFieldTagsTargetId
func GetNode_fieldTagsesByEntityIdAndRevisionIdAndFieldTagsTargetId(offset int, limit int, EntityId_ int, RevisionId_ int, FieldTagsTargetId_ int) (*[]*Node_fieldTags, error) {
	var _Node_fieldTags = new([]*Node_fieldTags)
	err := Engine.Table("node__field_tags").Where("entity_id = ? and revision_id = ? and field_tags_target_id = ?", EntityId_, RevisionId_, FieldTagsTargetId_).Limit(limit, offset).Find(_Node_fieldTags)
	return _Node_fieldTags, err
}

// GetNode_fieldTagsesByEntityIdAndLangcodeAndDelta Get Node_fieldTagses via EntityIdAndLangcodeAndDelta
func GetNode_fieldTagsesByEntityIdAndLangcodeAndDelta(offset int, limit int, EntityId_ int, Langcode_ string, Delta_ int) (*[]*Node_fieldTags, error) {
	var _Node_fieldTags = new([]*Node_fieldTags)
	err := Engine.Table("node__field_tags").Where("entity_id = ? and langcode = ? and delta = ?", EntityId_, Langcode_, Delta_).Limit(limit, offset).Find(_Node_fieldTags)
	return _Node_fieldTags, err
}

// GetNode_fieldTagsesByEntityIdAndLangcodeAndFieldTagsTargetId Get Node_fieldTagses via EntityIdAndLangcodeAndFieldTagsTargetId
func GetNode_fieldTagsesByEntityIdAndLangcodeAndFieldTagsTargetId(offset int, limit int, EntityId_ int, Langcode_ string, FieldTagsTargetId_ int) (*[]*Node_fieldTags, error) {
	var _Node_fieldTags = new([]*Node_fieldTags)
	err := Engine.Table("node__field_tags").Where("entity_id = ? and langcode = ? and field_tags_target_id = ?", EntityId_, Langcode_, FieldTagsTargetId_).Limit(limit, offset).Find(_Node_fieldTags)
	return _Node_fieldTags, err
}

// GetNode_fieldTagsesByEntityIdAndDeltaAndFieldTagsTargetId Get Node_fieldTagses via EntityIdAndDeltaAndFieldTagsTargetId
func GetNode_fieldTagsesByEntityIdAndDeltaAndFieldTagsTargetId(offset int, limit int, EntityId_ int, Delta_ int, FieldTagsTargetId_ int) (*[]*Node_fieldTags, error) {
	var _Node_fieldTags = new([]*Node_fieldTags)
	err := Engine.Table("node__field_tags").Where("entity_id = ? and delta = ? and field_tags_target_id = ?", EntityId_, Delta_, FieldTagsTargetId_).Limit(limit, offset).Find(_Node_fieldTags)
	return _Node_fieldTags, err
}

// GetNode_fieldTagsesByRevisionIdAndLangcodeAndDelta Get Node_fieldTagses via RevisionIdAndLangcodeAndDelta
func GetNode_fieldTagsesByRevisionIdAndLangcodeAndDelta(offset int, limit int, RevisionId_ int, Langcode_ string, Delta_ int) (*[]*Node_fieldTags, error) {
	var _Node_fieldTags = new([]*Node_fieldTags)
	err := Engine.Table("node__field_tags").Where("revision_id = ? and langcode = ? and delta = ?", RevisionId_, Langcode_, Delta_).Limit(limit, offset).Find(_Node_fieldTags)
	return _Node_fieldTags, err
}

// GetNode_fieldTagsesByRevisionIdAndLangcodeAndFieldTagsTargetId Get Node_fieldTagses via RevisionIdAndLangcodeAndFieldTagsTargetId
func GetNode_fieldTagsesByRevisionIdAndLangcodeAndFieldTagsTargetId(offset int, limit int, RevisionId_ int, Langcode_ string, FieldTagsTargetId_ int) (*[]*Node_fieldTags, error) {
	var _Node_fieldTags = new([]*Node_fieldTags)
	err := Engine.Table("node__field_tags").Where("revision_id = ? and langcode = ? and field_tags_target_id = ?", RevisionId_, Langcode_, FieldTagsTargetId_).Limit(limit, offset).Find(_Node_fieldTags)
	return _Node_fieldTags, err
}

// GetNode_fieldTagsesByRevisionIdAndDeltaAndFieldTagsTargetId Get Node_fieldTagses via RevisionIdAndDeltaAndFieldTagsTargetId
func GetNode_fieldTagsesByRevisionIdAndDeltaAndFieldTagsTargetId(offset int, limit int, RevisionId_ int, Delta_ int, FieldTagsTargetId_ int) (*[]*Node_fieldTags, error) {
	var _Node_fieldTags = new([]*Node_fieldTags)
	err := Engine.Table("node__field_tags").Where("revision_id = ? and delta = ? and field_tags_target_id = ?", RevisionId_, Delta_, FieldTagsTargetId_).Limit(limit, offset).Find(_Node_fieldTags)
	return _Node_fieldTags, err
}

// GetNode_fieldTagsesByLangcodeAndDeltaAndFieldTagsTargetId Get Node_fieldTagses via LangcodeAndDeltaAndFieldTagsTargetId
func GetNode_fieldTagsesByLangcodeAndDeltaAndFieldTagsTargetId(offset int, limit int, Langcode_ string, Delta_ int, FieldTagsTargetId_ int) (*[]*Node_fieldTags, error) {
	var _Node_fieldTags = new([]*Node_fieldTags)
	err := Engine.Table("node__field_tags").Where("langcode = ? and delta = ? and field_tags_target_id = ?", Langcode_, Delta_, FieldTagsTargetId_).Limit(limit, offset).Find(_Node_fieldTags)
	return _Node_fieldTags, err
}

// GetNode_fieldTagsesByBundleAndDeleted Get Node_fieldTagses via BundleAndDeleted
func GetNode_fieldTagsesByBundleAndDeleted(offset int, limit int, Bundle_ string, Deleted_ int) (*[]*Node_fieldTags, error) {
	var _Node_fieldTags = new([]*Node_fieldTags)
	err := Engine.Table("node__field_tags").Where("bundle = ? and deleted = ?", Bundle_, Deleted_).Limit(limit, offset).Find(_Node_fieldTags)
	return _Node_fieldTags, err
}

// GetNode_fieldTagsesByBundleAndEntityId Get Node_fieldTagses via BundleAndEntityId
func GetNode_fieldTagsesByBundleAndEntityId(offset int, limit int, Bundle_ string, EntityId_ int) (*[]*Node_fieldTags, error) {
	var _Node_fieldTags = new([]*Node_fieldTags)
	err := Engine.Table("node__field_tags").Where("bundle = ? and entity_id = ?", Bundle_, EntityId_).Limit(limit, offset).Find(_Node_fieldTags)
	return _Node_fieldTags, err
}

// GetNode_fieldTagsesByBundleAndRevisionId Get Node_fieldTagses via BundleAndRevisionId
func GetNode_fieldTagsesByBundleAndRevisionId(offset int, limit int, Bundle_ string, RevisionId_ int) (*[]*Node_fieldTags, error) {
	var _Node_fieldTags = new([]*Node_fieldTags)
	err := Engine.Table("node__field_tags").Where("bundle = ? and revision_id = ?", Bundle_, RevisionId_).Limit(limit, offset).Find(_Node_fieldTags)
	return _Node_fieldTags, err
}

// GetNode_fieldTagsesByBundleAndLangcode Get Node_fieldTagses via BundleAndLangcode
func GetNode_fieldTagsesByBundleAndLangcode(offset int, limit int, Bundle_ string, Langcode_ string) (*[]*Node_fieldTags, error) {
	var _Node_fieldTags = new([]*Node_fieldTags)
	err := Engine.Table("node__field_tags").Where("bundle = ? and langcode = ?", Bundle_, Langcode_).Limit(limit, offset).Find(_Node_fieldTags)
	return _Node_fieldTags, err
}

// GetNode_fieldTagsesByBundleAndDelta Get Node_fieldTagses via BundleAndDelta
func GetNode_fieldTagsesByBundleAndDelta(offset int, limit int, Bundle_ string, Delta_ int) (*[]*Node_fieldTags, error) {
	var _Node_fieldTags = new([]*Node_fieldTags)
	err := Engine.Table("node__field_tags").Where("bundle = ? and delta = ?", Bundle_, Delta_).Limit(limit, offset).Find(_Node_fieldTags)
	return _Node_fieldTags, err
}

// GetNode_fieldTagsesByBundleAndFieldTagsTargetId Get Node_fieldTagses via BundleAndFieldTagsTargetId
func GetNode_fieldTagsesByBundleAndFieldTagsTargetId(offset int, limit int, Bundle_ string, FieldTagsTargetId_ int) (*[]*Node_fieldTags, error) {
	var _Node_fieldTags = new([]*Node_fieldTags)
	err := Engine.Table("node__field_tags").Where("bundle = ? and field_tags_target_id = ?", Bundle_, FieldTagsTargetId_).Limit(limit, offset).Find(_Node_fieldTags)
	return _Node_fieldTags, err
}

// GetNode_fieldTagsesByDeletedAndEntityId Get Node_fieldTagses via DeletedAndEntityId
func GetNode_fieldTagsesByDeletedAndEntityId(offset int, limit int, Deleted_ int, EntityId_ int) (*[]*Node_fieldTags, error) {
	var _Node_fieldTags = new([]*Node_fieldTags)
	err := Engine.Table("node__field_tags").Where("deleted = ? and entity_id = ?", Deleted_, EntityId_).Limit(limit, offset).Find(_Node_fieldTags)
	return _Node_fieldTags, err
}

// GetNode_fieldTagsesByDeletedAndRevisionId Get Node_fieldTagses via DeletedAndRevisionId
func GetNode_fieldTagsesByDeletedAndRevisionId(offset int, limit int, Deleted_ int, RevisionId_ int) (*[]*Node_fieldTags, error) {
	var _Node_fieldTags = new([]*Node_fieldTags)
	err := Engine.Table("node__field_tags").Where("deleted = ? and revision_id = ?", Deleted_, RevisionId_).Limit(limit, offset).Find(_Node_fieldTags)
	return _Node_fieldTags, err
}

// GetNode_fieldTagsesByDeletedAndLangcode Get Node_fieldTagses via DeletedAndLangcode
func GetNode_fieldTagsesByDeletedAndLangcode(offset int, limit int, Deleted_ int, Langcode_ string) (*[]*Node_fieldTags, error) {
	var _Node_fieldTags = new([]*Node_fieldTags)
	err := Engine.Table("node__field_tags").Where("deleted = ? and langcode = ?", Deleted_, Langcode_).Limit(limit, offset).Find(_Node_fieldTags)
	return _Node_fieldTags, err
}

// GetNode_fieldTagsesByDeletedAndDelta Get Node_fieldTagses via DeletedAndDelta
func GetNode_fieldTagsesByDeletedAndDelta(offset int, limit int, Deleted_ int, Delta_ int) (*[]*Node_fieldTags, error) {
	var _Node_fieldTags = new([]*Node_fieldTags)
	err := Engine.Table("node__field_tags").Where("deleted = ? and delta = ?", Deleted_, Delta_).Limit(limit, offset).Find(_Node_fieldTags)
	return _Node_fieldTags, err
}

// GetNode_fieldTagsesByDeletedAndFieldTagsTargetId Get Node_fieldTagses via DeletedAndFieldTagsTargetId
func GetNode_fieldTagsesByDeletedAndFieldTagsTargetId(offset int, limit int, Deleted_ int, FieldTagsTargetId_ int) (*[]*Node_fieldTags, error) {
	var _Node_fieldTags = new([]*Node_fieldTags)
	err := Engine.Table("node__field_tags").Where("deleted = ? and field_tags_target_id = ?", Deleted_, FieldTagsTargetId_).Limit(limit, offset).Find(_Node_fieldTags)
	return _Node_fieldTags, err
}

// GetNode_fieldTagsesByEntityIdAndRevisionId Get Node_fieldTagses via EntityIdAndRevisionId
func GetNode_fieldTagsesByEntityIdAndRevisionId(offset int, limit int, EntityId_ int, RevisionId_ int) (*[]*Node_fieldTags, error) {
	var _Node_fieldTags = new([]*Node_fieldTags)
	err := Engine.Table("node__field_tags").Where("entity_id = ? and revision_id = ?", EntityId_, RevisionId_).Limit(limit, offset).Find(_Node_fieldTags)
	return _Node_fieldTags, err
}

// GetNode_fieldTagsesByEntityIdAndLangcode Get Node_fieldTagses via EntityIdAndLangcode
func GetNode_fieldTagsesByEntityIdAndLangcode(offset int, limit int, EntityId_ int, Langcode_ string) (*[]*Node_fieldTags, error) {
	var _Node_fieldTags = new([]*Node_fieldTags)
	err := Engine.Table("node__field_tags").Where("entity_id = ? and langcode = ?", EntityId_, Langcode_).Limit(limit, offset).Find(_Node_fieldTags)
	return _Node_fieldTags, err
}

// GetNode_fieldTagsesByEntityIdAndDelta Get Node_fieldTagses via EntityIdAndDelta
func GetNode_fieldTagsesByEntityIdAndDelta(offset int, limit int, EntityId_ int, Delta_ int) (*[]*Node_fieldTags, error) {
	var _Node_fieldTags = new([]*Node_fieldTags)
	err := Engine.Table("node__field_tags").Where("entity_id = ? and delta = ?", EntityId_, Delta_).Limit(limit, offset).Find(_Node_fieldTags)
	return _Node_fieldTags, err
}

// GetNode_fieldTagsesByEntityIdAndFieldTagsTargetId Get Node_fieldTagses via EntityIdAndFieldTagsTargetId
func GetNode_fieldTagsesByEntityIdAndFieldTagsTargetId(offset int, limit int, EntityId_ int, FieldTagsTargetId_ int) (*[]*Node_fieldTags, error) {
	var _Node_fieldTags = new([]*Node_fieldTags)
	err := Engine.Table("node__field_tags").Where("entity_id = ? and field_tags_target_id = ?", EntityId_, FieldTagsTargetId_).Limit(limit, offset).Find(_Node_fieldTags)
	return _Node_fieldTags, err
}

// GetNode_fieldTagsesByRevisionIdAndLangcode Get Node_fieldTagses via RevisionIdAndLangcode
func GetNode_fieldTagsesByRevisionIdAndLangcode(offset int, limit int, RevisionId_ int, Langcode_ string) (*[]*Node_fieldTags, error) {
	var _Node_fieldTags = new([]*Node_fieldTags)
	err := Engine.Table("node__field_tags").Where("revision_id = ? and langcode = ?", RevisionId_, Langcode_).Limit(limit, offset).Find(_Node_fieldTags)
	return _Node_fieldTags, err
}

// GetNode_fieldTagsesByRevisionIdAndDelta Get Node_fieldTagses via RevisionIdAndDelta
func GetNode_fieldTagsesByRevisionIdAndDelta(offset int, limit int, RevisionId_ int, Delta_ int) (*[]*Node_fieldTags, error) {
	var _Node_fieldTags = new([]*Node_fieldTags)
	err := Engine.Table("node__field_tags").Where("revision_id = ? and delta = ?", RevisionId_, Delta_).Limit(limit, offset).Find(_Node_fieldTags)
	return _Node_fieldTags, err
}

// GetNode_fieldTagsesByRevisionIdAndFieldTagsTargetId Get Node_fieldTagses via RevisionIdAndFieldTagsTargetId
func GetNode_fieldTagsesByRevisionIdAndFieldTagsTargetId(offset int, limit int, RevisionId_ int, FieldTagsTargetId_ int) (*[]*Node_fieldTags, error) {
	var _Node_fieldTags = new([]*Node_fieldTags)
	err := Engine.Table("node__field_tags").Where("revision_id = ? and field_tags_target_id = ?", RevisionId_, FieldTagsTargetId_).Limit(limit, offset).Find(_Node_fieldTags)
	return _Node_fieldTags, err
}

// GetNode_fieldTagsesByLangcodeAndDelta Get Node_fieldTagses via LangcodeAndDelta
func GetNode_fieldTagsesByLangcodeAndDelta(offset int, limit int, Langcode_ string, Delta_ int) (*[]*Node_fieldTags, error) {
	var _Node_fieldTags = new([]*Node_fieldTags)
	err := Engine.Table("node__field_tags").Where("langcode = ? and delta = ?", Langcode_, Delta_).Limit(limit, offset).Find(_Node_fieldTags)
	return _Node_fieldTags, err
}

// GetNode_fieldTagsesByLangcodeAndFieldTagsTargetId Get Node_fieldTagses via LangcodeAndFieldTagsTargetId
func GetNode_fieldTagsesByLangcodeAndFieldTagsTargetId(offset int, limit int, Langcode_ string, FieldTagsTargetId_ int) (*[]*Node_fieldTags, error) {
	var _Node_fieldTags = new([]*Node_fieldTags)
	err := Engine.Table("node__field_tags").Where("langcode = ? and field_tags_target_id = ?", Langcode_, FieldTagsTargetId_).Limit(limit, offset).Find(_Node_fieldTags)
	return _Node_fieldTags, err
}

// GetNode_fieldTagsesByDeltaAndFieldTagsTargetId Get Node_fieldTagses via DeltaAndFieldTagsTargetId
func GetNode_fieldTagsesByDeltaAndFieldTagsTargetId(offset int, limit int, Delta_ int, FieldTagsTargetId_ int) (*[]*Node_fieldTags, error) {
	var _Node_fieldTags = new([]*Node_fieldTags)
	err := Engine.Table("node__field_tags").Where("delta = ? and field_tags_target_id = ?", Delta_, FieldTagsTargetId_).Limit(limit, offset).Find(_Node_fieldTags)
	return _Node_fieldTags, err
}

// GetNode_fieldTagses Get Node_fieldTagses via field
func GetNode_fieldTagses(offset int, limit int, field string) (*[]*Node_fieldTags, error) {
	var _Node_fieldTags = new([]*Node_fieldTags)
	err := Engine.Table("node__field_tags").Limit(limit, offset).Desc(field).Find(_Node_fieldTags)
	return _Node_fieldTags, err
}

// GetNode_fieldTagsesViaBundle Get Node_fieldTagss via Bundle
func GetNode_fieldTagsesViaBundle(offset int, limit int, Bundle_ string, field string) (*[]*Node_fieldTags, error) {
	var _Node_fieldTags = new([]*Node_fieldTags)
	err := Engine.Table("node__field_tags").Where("bundle = ?", Bundle_).Limit(limit, offset).Desc(field).Find(_Node_fieldTags)
	return _Node_fieldTags, err
}

// GetNode_fieldTagsesViaDeleted Get Node_fieldTagss via Deleted
func GetNode_fieldTagsesViaDeleted(offset int, limit int, Deleted_ int, field string) (*[]*Node_fieldTags, error) {
	var _Node_fieldTags = new([]*Node_fieldTags)
	err := Engine.Table("node__field_tags").Where("deleted = ?", Deleted_).Limit(limit, offset).Desc(field).Find(_Node_fieldTags)
	return _Node_fieldTags, err
}

// GetNode_fieldTagsesViaEntityId Get Node_fieldTagss via EntityId
func GetNode_fieldTagsesViaEntityId(offset int, limit int, EntityId_ int, field string) (*[]*Node_fieldTags, error) {
	var _Node_fieldTags = new([]*Node_fieldTags)
	err := Engine.Table("node__field_tags").Where("entity_id = ?", EntityId_).Limit(limit, offset).Desc(field).Find(_Node_fieldTags)
	return _Node_fieldTags, err
}

// GetNode_fieldTagsesViaRevisionId Get Node_fieldTagss via RevisionId
func GetNode_fieldTagsesViaRevisionId(offset int, limit int, RevisionId_ int, field string) (*[]*Node_fieldTags, error) {
	var _Node_fieldTags = new([]*Node_fieldTags)
	err := Engine.Table("node__field_tags").Where("revision_id = ?", RevisionId_).Limit(limit, offset).Desc(field).Find(_Node_fieldTags)
	return _Node_fieldTags, err
}

// GetNode_fieldTagsesViaLangcode Get Node_fieldTagss via Langcode
func GetNode_fieldTagsesViaLangcode(offset int, limit int, Langcode_ string, field string) (*[]*Node_fieldTags, error) {
	var _Node_fieldTags = new([]*Node_fieldTags)
	err := Engine.Table("node__field_tags").Where("langcode = ?", Langcode_).Limit(limit, offset).Desc(field).Find(_Node_fieldTags)
	return _Node_fieldTags, err
}

// GetNode_fieldTagsesViaDelta Get Node_fieldTagss via Delta
func GetNode_fieldTagsesViaDelta(offset int, limit int, Delta_ int, field string) (*[]*Node_fieldTags, error) {
	var _Node_fieldTags = new([]*Node_fieldTags)
	err := Engine.Table("node__field_tags").Where("delta = ?", Delta_).Limit(limit, offset).Desc(field).Find(_Node_fieldTags)
	return _Node_fieldTags, err
}

// GetNode_fieldTagsesViaFieldTagsTargetId Get Node_fieldTagss via FieldTagsTargetId
func GetNode_fieldTagsesViaFieldTagsTargetId(offset int, limit int, FieldTagsTargetId_ int, field string) (*[]*Node_fieldTags, error) {
	var _Node_fieldTags = new([]*Node_fieldTags)
	err := Engine.Table("node__field_tags").Where("field_tags_target_id = ?", FieldTagsTargetId_).Limit(limit, offset).Desc(field).Find(_Node_fieldTags)
	return _Node_fieldTags, err
}

// HasNode_fieldTagsViaBundle Has Node_fieldTags via Bundle
func HasNode_fieldTagsViaBundle(iBundle string) bool {
	if has, err := Engine.Where("bundle = ?", iBundle).Get(new(Node_fieldTags)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasNode_fieldTagsViaDeleted Has Node_fieldTags via Deleted
func HasNode_fieldTagsViaDeleted(iDeleted int) bool {
	if has, err := Engine.Where("deleted = ?", iDeleted).Get(new(Node_fieldTags)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasNode_fieldTagsViaEntityId Has Node_fieldTags via EntityId
func HasNode_fieldTagsViaEntityId(iEntityId int) bool {
	if has, err := Engine.Where("entity_id = ?", iEntityId).Get(new(Node_fieldTags)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasNode_fieldTagsViaRevisionId Has Node_fieldTags via RevisionId
func HasNode_fieldTagsViaRevisionId(iRevisionId int) bool {
	if has, err := Engine.Where("revision_id = ?", iRevisionId).Get(new(Node_fieldTags)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasNode_fieldTagsViaLangcode Has Node_fieldTags via Langcode
func HasNode_fieldTagsViaLangcode(iLangcode string) bool {
	if has, err := Engine.Where("langcode = ?", iLangcode).Get(new(Node_fieldTags)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasNode_fieldTagsViaDelta Has Node_fieldTags via Delta
func HasNode_fieldTagsViaDelta(iDelta int) bool {
	if has, err := Engine.Where("delta = ?", iDelta).Get(new(Node_fieldTags)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasNode_fieldTagsViaFieldTagsTargetId Has Node_fieldTags via FieldTagsTargetId
func HasNode_fieldTagsViaFieldTagsTargetId(iFieldTagsTargetId int) bool {
	if has, err := Engine.Where("field_tags_target_id = ?", iFieldTagsTargetId).Get(new(Node_fieldTags)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// GetNode_fieldTagsViaBundle Get Node_fieldTags via Bundle
func GetNode_fieldTagsViaBundle(iBundle string) (*Node_fieldTags, error) {
	var _Node_fieldTags = &Node_fieldTags{Bundle: iBundle}
	has, err := Engine.Get(_Node_fieldTags)
	if has {
		return _Node_fieldTags, err
	} else {
		return nil, err
	}
}

// GetNode_fieldTagsViaDeleted Get Node_fieldTags via Deleted
func GetNode_fieldTagsViaDeleted(iDeleted int) (*Node_fieldTags, error) {
	var _Node_fieldTags = &Node_fieldTags{Deleted: iDeleted}
	has, err := Engine.Get(_Node_fieldTags)
	if has {
		return _Node_fieldTags, err
	} else {
		return nil, err
	}
}

// GetNode_fieldTagsViaEntityId Get Node_fieldTags via EntityId
func GetNode_fieldTagsViaEntityId(iEntityId int) (*Node_fieldTags, error) {
	var _Node_fieldTags = &Node_fieldTags{EntityId: iEntityId}
	has, err := Engine.Get(_Node_fieldTags)
	if has {
		return _Node_fieldTags, err
	} else {
		return nil, err
	}
}

// GetNode_fieldTagsViaRevisionId Get Node_fieldTags via RevisionId
func GetNode_fieldTagsViaRevisionId(iRevisionId int) (*Node_fieldTags, error) {
	var _Node_fieldTags = &Node_fieldTags{RevisionId: iRevisionId}
	has, err := Engine.Get(_Node_fieldTags)
	if has {
		return _Node_fieldTags, err
	} else {
		return nil, err
	}
}

// GetNode_fieldTagsViaLangcode Get Node_fieldTags via Langcode
func GetNode_fieldTagsViaLangcode(iLangcode string) (*Node_fieldTags, error) {
	var _Node_fieldTags = &Node_fieldTags{Langcode: iLangcode}
	has, err := Engine.Get(_Node_fieldTags)
	if has {
		return _Node_fieldTags, err
	} else {
		return nil, err
	}
}

// GetNode_fieldTagsViaDelta Get Node_fieldTags via Delta
func GetNode_fieldTagsViaDelta(iDelta int) (*Node_fieldTags, error) {
	var _Node_fieldTags = &Node_fieldTags{Delta: iDelta}
	has, err := Engine.Get(_Node_fieldTags)
	if has {
		return _Node_fieldTags, err
	} else {
		return nil, err
	}
}

// GetNode_fieldTagsViaFieldTagsTargetId Get Node_fieldTags via FieldTagsTargetId
func GetNode_fieldTagsViaFieldTagsTargetId(iFieldTagsTargetId int) (*Node_fieldTags, error) {
	var _Node_fieldTags = &Node_fieldTags{FieldTagsTargetId: iFieldTagsTargetId}
	has, err := Engine.Get(_Node_fieldTags)
	if has {
		return _Node_fieldTags, err
	} else {
		return nil, err
	}
}

// SetNode_fieldTagsViaBundle Set Node_fieldTags via Bundle
func SetNode_fieldTagsViaBundle(iBundle string, node__field_tags *Node_fieldTags) (int64, error) {
	node__field_tags.Bundle = iBundle
	return Engine.Insert(node__field_tags)
}

// SetNode_fieldTagsViaDeleted Set Node_fieldTags via Deleted
func SetNode_fieldTagsViaDeleted(iDeleted int, node__field_tags *Node_fieldTags) (int64, error) {
	node__field_tags.Deleted = iDeleted
	return Engine.Insert(node__field_tags)
}

// SetNode_fieldTagsViaEntityId Set Node_fieldTags via EntityId
func SetNode_fieldTagsViaEntityId(iEntityId int, node__field_tags *Node_fieldTags) (int64, error) {
	node__field_tags.EntityId = iEntityId
	return Engine.Insert(node__field_tags)
}

// SetNode_fieldTagsViaRevisionId Set Node_fieldTags via RevisionId
func SetNode_fieldTagsViaRevisionId(iRevisionId int, node__field_tags *Node_fieldTags) (int64, error) {
	node__field_tags.RevisionId = iRevisionId
	return Engine.Insert(node__field_tags)
}

// SetNode_fieldTagsViaLangcode Set Node_fieldTags via Langcode
func SetNode_fieldTagsViaLangcode(iLangcode string, node__field_tags *Node_fieldTags) (int64, error) {
	node__field_tags.Langcode = iLangcode
	return Engine.Insert(node__field_tags)
}

// SetNode_fieldTagsViaDelta Set Node_fieldTags via Delta
func SetNode_fieldTagsViaDelta(iDelta int, node__field_tags *Node_fieldTags) (int64, error) {
	node__field_tags.Delta = iDelta
	return Engine.Insert(node__field_tags)
}

// SetNode_fieldTagsViaFieldTagsTargetId Set Node_fieldTags via FieldTagsTargetId
func SetNode_fieldTagsViaFieldTagsTargetId(iFieldTagsTargetId int, node__field_tags *Node_fieldTags) (int64, error) {
	node__field_tags.FieldTagsTargetId = iFieldTagsTargetId
	return Engine.Insert(node__field_tags)
}

// AddNode_fieldTags Add Node_fieldTags via all columns
func AddNode_fieldTags(iBundle string, iDeleted int, iEntityId int, iRevisionId int, iLangcode string, iDelta int, iFieldTagsTargetId int) error {
	_Node_fieldTags := &Node_fieldTags{Bundle: iBundle, Deleted: iDeleted, EntityId: iEntityId, RevisionId: iRevisionId, Langcode: iLangcode, Delta: iDelta, FieldTagsTargetId: iFieldTagsTargetId}
	if _, err := Engine.Insert(_Node_fieldTags); err != nil {
		return err
	} else {
		return nil
	}
}

// PostNode_fieldTags Post Node_fieldTags via iNode_fieldTags
func PostNode_fieldTags(iNode_fieldTags *Node_fieldTags) (string, error) {
	_, err := Engine.Insert(iNode_fieldTags)
	return iNode_fieldTags.Bundle, err
}

// PutNode_fieldTags Put Node_fieldTags
func PutNode_fieldTags(iNode_fieldTags *Node_fieldTags) (string, error) {
	_, err := Engine.Id(iNode_fieldTags.Bundle).Update(iNode_fieldTags)
	return iNode_fieldTags.Bundle, err
}

// PutNode_fieldTagsViaBundle Put Node_fieldTags via Bundle
func PutNode_fieldTagsViaBundle(Bundle_ string, iNode_fieldTags *Node_fieldTags) (int64, error) {
	row, err := Engine.Update(iNode_fieldTags, &Node_fieldTags{Bundle: Bundle_})
	return row, err
}

// PutNode_fieldTagsViaDeleted Put Node_fieldTags via Deleted
func PutNode_fieldTagsViaDeleted(Deleted_ int, iNode_fieldTags *Node_fieldTags) (int64, error) {
	row, err := Engine.Update(iNode_fieldTags, &Node_fieldTags{Deleted: Deleted_})
	return row, err
}

// PutNode_fieldTagsViaEntityId Put Node_fieldTags via EntityId
func PutNode_fieldTagsViaEntityId(EntityId_ int, iNode_fieldTags *Node_fieldTags) (int64, error) {
	row, err := Engine.Update(iNode_fieldTags, &Node_fieldTags{EntityId: EntityId_})
	return row, err
}

// PutNode_fieldTagsViaRevisionId Put Node_fieldTags via RevisionId
func PutNode_fieldTagsViaRevisionId(RevisionId_ int, iNode_fieldTags *Node_fieldTags) (int64, error) {
	row, err := Engine.Update(iNode_fieldTags, &Node_fieldTags{RevisionId: RevisionId_})
	return row, err
}

// PutNode_fieldTagsViaLangcode Put Node_fieldTags via Langcode
func PutNode_fieldTagsViaLangcode(Langcode_ string, iNode_fieldTags *Node_fieldTags) (int64, error) {
	row, err := Engine.Update(iNode_fieldTags, &Node_fieldTags{Langcode: Langcode_})
	return row, err
}

// PutNode_fieldTagsViaDelta Put Node_fieldTags via Delta
func PutNode_fieldTagsViaDelta(Delta_ int, iNode_fieldTags *Node_fieldTags) (int64, error) {
	row, err := Engine.Update(iNode_fieldTags, &Node_fieldTags{Delta: Delta_})
	return row, err
}

// PutNode_fieldTagsViaFieldTagsTargetId Put Node_fieldTags via FieldTagsTargetId
func PutNode_fieldTagsViaFieldTagsTargetId(FieldTagsTargetId_ int, iNode_fieldTags *Node_fieldTags) (int64, error) {
	row, err := Engine.Update(iNode_fieldTags, &Node_fieldTags{FieldTagsTargetId: FieldTagsTargetId_})
	return row, err
}

// UpdateNode_fieldTagsViaBundle via map[string]interface{}{}
func UpdateNode_fieldTagsViaBundle(iBundle string, iNode_fieldTagsMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Node_fieldTags)).Where("bundle = ?", iBundle).Update(iNode_fieldTagsMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateNode_fieldTagsViaDeleted via map[string]interface{}{}
func UpdateNode_fieldTagsViaDeleted(iDeleted int, iNode_fieldTagsMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Node_fieldTags)).Where("deleted = ?", iDeleted).Update(iNode_fieldTagsMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateNode_fieldTagsViaEntityId via map[string]interface{}{}
func UpdateNode_fieldTagsViaEntityId(iEntityId int, iNode_fieldTagsMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Node_fieldTags)).Where("entity_id = ?", iEntityId).Update(iNode_fieldTagsMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateNode_fieldTagsViaRevisionId via map[string]interface{}{}
func UpdateNode_fieldTagsViaRevisionId(iRevisionId int, iNode_fieldTagsMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Node_fieldTags)).Where("revision_id = ?", iRevisionId).Update(iNode_fieldTagsMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateNode_fieldTagsViaLangcode via map[string]interface{}{}
func UpdateNode_fieldTagsViaLangcode(iLangcode string, iNode_fieldTagsMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Node_fieldTags)).Where("langcode = ?", iLangcode).Update(iNode_fieldTagsMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateNode_fieldTagsViaDelta via map[string]interface{}{}
func UpdateNode_fieldTagsViaDelta(iDelta int, iNode_fieldTagsMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Node_fieldTags)).Where("delta = ?", iDelta).Update(iNode_fieldTagsMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateNode_fieldTagsViaFieldTagsTargetId via map[string]interface{}{}
func UpdateNode_fieldTagsViaFieldTagsTargetId(iFieldTagsTargetId int, iNode_fieldTagsMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Node_fieldTags)).Where("field_tags_target_id = ?", iFieldTagsTargetId).Update(iNode_fieldTagsMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// DeleteNode_fieldTags Delete Node_fieldTags
func DeleteNode_fieldTags(iBundle string) (int64, error) {
	row, err := Engine.Id(iBundle).Delete(new(Node_fieldTags))
	return row, err
}

// DeleteNode_fieldTagsViaBundle Delete Node_fieldTags via Bundle
func DeleteNode_fieldTagsViaBundle(iBundle string) (err error) {
	var has bool
	var _Node_fieldTags = &Node_fieldTags{Bundle: iBundle}
	if has, err = Engine.Get(_Node_fieldTags); (has == true) && (err == nil) {
		if row, err := Engine.Where("bundle = ?", iBundle).Delete(new(Node_fieldTags)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteNode_fieldTagsViaDeleted Delete Node_fieldTags via Deleted
func DeleteNode_fieldTagsViaDeleted(iDeleted int) (err error) {
	var has bool
	var _Node_fieldTags = &Node_fieldTags{Deleted: iDeleted}
	if has, err = Engine.Get(_Node_fieldTags); (has == true) && (err == nil) {
		if row, err := Engine.Where("deleted = ?", iDeleted).Delete(new(Node_fieldTags)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteNode_fieldTagsViaEntityId Delete Node_fieldTags via EntityId
func DeleteNode_fieldTagsViaEntityId(iEntityId int) (err error) {
	var has bool
	var _Node_fieldTags = &Node_fieldTags{EntityId: iEntityId}
	if has, err = Engine.Get(_Node_fieldTags); (has == true) && (err == nil) {
		if row, err := Engine.Where("entity_id = ?", iEntityId).Delete(new(Node_fieldTags)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteNode_fieldTagsViaRevisionId Delete Node_fieldTags via RevisionId
func DeleteNode_fieldTagsViaRevisionId(iRevisionId int) (err error) {
	var has bool
	var _Node_fieldTags = &Node_fieldTags{RevisionId: iRevisionId}
	if has, err = Engine.Get(_Node_fieldTags); (has == true) && (err == nil) {
		if row, err := Engine.Where("revision_id = ?", iRevisionId).Delete(new(Node_fieldTags)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteNode_fieldTagsViaLangcode Delete Node_fieldTags via Langcode
func DeleteNode_fieldTagsViaLangcode(iLangcode string) (err error) {
	var has bool
	var _Node_fieldTags = &Node_fieldTags{Langcode: iLangcode}
	if has, err = Engine.Get(_Node_fieldTags); (has == true) && (err == nil) {
		if row, err := Engine.Where("langcode = ?", iLangcode).Delete(new(Node_fieldTags)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteNode_fieldTagsViaDelta Delete Node_fieldTags via Delta
func DeleteNode_fieldTagsViaDelta(iDelta int) (err error) {
	var has bool
	var _Node_fieldTags = &Node_fieldTags{Delta: iDelta}
	if has, err = Engine.Get(_Node_fieldTags); (has == true) && (err == nil) {
		if row, err := Engine.Where("delta = ?", iDelta).Delete(new(Node_fieldTags)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteNode_fieldTagsViaFieldTagsTargetId Delete Node_fieldTags via FieldTagsTargetId
func DeleteNode_fieldTagsViaFieldTagsTargetId(iFieldTagsTargetId int) (err error) {
	var has bool
	var _Node_fieldTags = &Node_fieldTags{FieldTagsTargetId: iFieldTagsTargetId}
	if has, err = Engine.Get(_Node_fieldTags); (has == true) && (err == nil) {
		if row, err := Engine.Where("field_tags_target_id = ?", iFieldTagsTargetId).Delete(new(Node_fieldTags)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}
