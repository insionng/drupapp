package model

type NodeRevision_fieldImage struct {
	Bundle             string `xorm:"not null default '' index VARCHAR(128)"`
	Deleted            int    `xorm:"not null pk default 0 TINYINT(4)"`
	EntityId           int    `xorm:"not null pk INT(10)"`
	RevisionId         int    `xorm:"not null pk index INT(10)"`
	Langcode           string `xorm:"not null pk default '' VARCHAR(32)"`
	Delta              int    `xorm:"not null pk INT(10)"`
	FieldImageTargetId int    `xorm:"not null index INT(10)"`
	FieldImageAlt      string `xorm:"VARCHAR(512)"`
	FieldImageTitle    string `xorm:"VARCHAR(1024)"`
	FieldImageWidth    int    `xorm:"INT(10)"`
	FieldImageHeight   int    `xorm:"INT(10)"`
}

// GetNodeRevision_fieldImagesCount NodeRevision_fieldImages' Count
func GetNodeRevision_fieldImagesCount(offset int, limit int) (int64, error) {
	total, err := Engine.Limit(limit, offset).Count(&NodeRevision_fieldImage{})
	return total, err
}

// GetNodeRevision_fieldImageCountViaBundle Get NodeRevision_fieldImage via Bundle
func GetNodeRevision_fieldImageCountViaBundle(iBundle string) int64 {
	n, _ := Engine.Where("bundle = ?", iBundle).Count(&NodeRevision_fieldImage{Bundle: iBundle})
	return n
}

// GetNodeRevision_fieldImageCountViaDeleted Get NodeRevision_fieldImage via Deleted
func GetNodeRevision_fieldImageCountViaDeleted(iDeleted int) int64 {
	n, _ := Engine.Where("deleted = ?", iDeleted).Count(&NodeRevision_fieldImage{Deleted: iDeleted})
	return n
}

// GetNodeRevision_fieldImageCountViaEntityId Get NodeRevision_fieldImage via EntityId
func GetNodeRevision_fieldImageCountViaEntityId(iEntityId int) int64 {
	n, _ := Engine.Where("entity_id = ?", iEntityId).Count(&NodeRevision_fieldImage{EntityId: iEntityId})
	return n
}

// GetNodeRevision_fieldImageCountViaRevisionId Get NodeRevision_fieldImage via RevisionId
func GetNodeRevision_fieldImageCountViaRevisionId(iRevisionId int) int64 {
	n, _ := Engine.Where("revision_id = ?", iRevisionId).Count(&NodeRevision_fieldImage{RevisionId: iRevisionId})
	return n
}

// GetNodeRevision_fieldImageCountViaLangcode Get NodeRevision_fieldImage via Langcode
func GetNodeRevision_fieldImageCountViaLangcode(iLangcode string) int64 {
	n, _ := Engine.Where("langcode = ?", iLangcode).Count(&NodeRevision_fieldImage{Langcode: iLangcode})
	return n
}

// GetNodeRevision_fieldImageCountViaDelta Get NodeRevision_fieldImage via Delta
func GetNodeRevision_fieldImageCountViaDelta(iDelta int) int64 {
	n, _ := Engine.Where("delta = ?", iDelta).Count(&NodeRevision_fieldImage{Delta: iDelta})
	return n
}

// GetNodeRevision_fieldImageCountViaFieldImageTargetId Get NodeRevision_fieldImage via FieldImageTargetId
func GetNodeRevision_fieldImageCountViaFieldImageTargetId(iFieldImageTargetId int) int64 {
	n, _ := Engine.Where("field_image_target_id = ?", iFieldImageTargetId).Count(&NodeRevision_fieldImage{FieldImageTargetId: iFieldImageTargetId})
	return n
}

// GetNodeRevision_fieldImageCountViaFieldImageAlt Get NodeRevision_fieldImage via FieldImageAlt
func GetNodeRevision_fieldImageCountViaFieldImageAlt(iFieldImageAlt string) int64 {
	n, _ := Engine.Where("field_image_alt = ?", iFieldImageAlt).Count(&NodeRevision_fieldImage{FieldImageAlt: iFieldImageAlt})
	return n
}

// GetNodeRevision_fieldImageCountViaFieldImageTitle Get NodeRevision_fieldImage via FieldImageTitle
func GetNodeRevision_fieldImageCountViaFieldImageTitle(iFieldImageTitle string) int64 {
	n, _ := Engine.Where("field_image_title = ?", iFieldImageTitle).Count(&NodeRevision_fieldImage{FieldImageTitle: iFieldImageTitle})
	return n
}

// GetNodeRevision_fieldImageCountViaFieldImageWidth Get NodeRevision_fieldImage via FieldImageWidth
func GetNodeRevision_fieldImageCountViaFieldImageWidth(iFieldImageWidth int) int64 {
	n, _ := Engine.Where("field_image_width = ?", iFieldImageWidth).Count(&NodeRevision_fieldImage{FieldImageWidth: iFieldImageWidth})
	return n
}

// GetNodeRevision_fieldImageCountViaFieldImageHeight Get NodeRevision_fieldImage via FieldImageHeight
func GetNodeRevision_fieldImageCountViaFieldImageHeight(iFieldImageHeight int) int64 {
	n, _ := Engine.Where("field_image_height = ?", iFieldImageHeight).Count(&NodeRevision_fieldImage{FieldImageHeight: iFieldImageHeight})
	return n
}

// GetNodeRevision_fieldImagesByBundleAndDeletedAndEntityId Get NodeRevision_fieldImages via BundleAndDeletedAndEntityId
func GetNodeRevision_fieldImagesByBundleAndDeletedAndEntityId(offset int, limit int, Bundle_ string, Deleted_ int, EntityId_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("bundle = ? and deleted = ? and entity_id = ?", Bundle_, Deleted_, EntityId_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByBundleAndDeletedAndRevisionId Get NodeRevision_fieldImages via BundleAndDeletedAndRevisionId
func GetNodeRevision_fieldImagesByBundleAndDeletedAndRevisionId(offset int, limit int, Bundle_ string, Deleted_ int, RevisionId_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("bundle = ? and deleted = ? and revision_id = ?", Bundle_, Deleted_, RevisionId_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByBundleAndDeletedAndLangcode Get NodeRevision_fieldImages via BundleAndDeletedAndLangcode
func GetNodeRevision_fieldImagesByBundleAndDeletedAndLangcode(offset int, limit int, Bundle_ string, Deleted_ int, Langcode_ string) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("bundle = ? and deleted = ? and langcode = ?", Bundle_, Deleted_, Langcode_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByBundleAndDeletedAndDelta Get NodeRevision_fieldImages via BundleAndDeletedAndDelta
func GetNodeRevision_fieldImagesByBundleAndDeletedAndDelta(offset int, limit int, Bundle_ string, Deleted_ int, Delta_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("bundle = ? and deleted = ? and delta = ?", Bundle_, Deleted_, Delta_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByBundleAndDeletedAndFieldImageTargetId Get NodeRevision_fieldImages via BundleAndDeletedAndFieldImageTargetId
func GetNodeRevision_fieldImagesByBundleAndDeletedAndFieldImageTargetId(offset int, limit int, Bundle_ string, Deleted_ int, FieldImageTargetId_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("bundle = ? and deleted = ? and field_image_target_id = ?", Bundle_, Deleted_, FieldImageTargetId_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByBundleAndDeletedAndFieldImageAlt Get NodeRevision_fieldImages via BundleAndDeletedAndFieldImageAlt
func GetNodeRevision_fieldImagesByBundleAndDeletedAndFieldImageAlt(offset int, limit int, Bundle_ string, Deleted_ int, FieldImageAlt_ string) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("bundle = ? and deleted = ? and field_image_alt = ?", Bundle_, Deleted_, FieldImageAlt_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByBundleAndDeletedAndFieldImageTitle Get NodeRevision_fieldImages via BundleAndDeletedAndFieldImageTitle
func GetNodeRevision_fieldImagesByBundleAndDeletedAndFieldImageTitle(offset int, limit int, Bundle_ string, Deleted_ int, FieldImageTitle_ string) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("bundle = ? and deleted = ? and field_image_title = ?", Bundle_, Deleted_, FieldImageTitle_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByBundleAndDeletedAndFieldImageWidth Get NodeRevision_fieldImages via BundleAndDeletedAndFieldImageWidth
func GetNodeRevision_fieldImagesByBundleAndDeletedAndFieldImageWidth(offset int, limit int, Bundle_ string, Deleted_ int, FieldImageWidth_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("bundle = ? and deleted = ? and field_image_width = ?", Bundle_, Deleted_, FieldImageWidth_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByBundleAndDeletedAndFieldImageHeight Get NodeRevision_fieldImages via BundleAndDeletedAndFieldImageHeight
func GetNodeRevision_fieldImagesByBundleAndDeletedAndFieldImageHeight(offset int, limit int, Bundle_ string, Deleted_ int, FieldImageHeight_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("bundle = ? and deleted = ? and field_image_height = ?", Bundle_, Deleted_, FieldImageHeight_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByBundleAndEntityIdAndRevisionId Get NodeRevision_fieldImages via BundleAndEntityIdAndRevisionId
func GetNodeRevision_fieldImagesByBundleAndEntityIdAndRevisionId(offset int, limit int, Bundle_ string, EntityId_ int, RevisionId_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("bundle = ? and entity_id = ? and revision_id = ?", Bundle_, EntityId_, RevisionId_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByBundleAndEntityIdAndLangcode Get NodeRevision_fieldImages via BundleAndEntityIdAndLangcode
func GetNodeRevision_fieldImagesByBundleAndEntityIdAndLangcode(offset int, limit int, Bundle_ string, EntityId_ int, Langcode_ string) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("bundle = ? and entity_id = ? and langcode = ?", Bundle_, EntityId_, Langcode_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByBundleAndEntityIdAndDelta Get NodeRevision_fieldImages via BundleAndEntityIdAndDelta
func GetNodeRevision_fieldImagesByBundleAndEntityIdAndDelta(offset int, limit int, Bundle_ string, EntityId_ int, Delta_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("bundle = ? and entity_id = ? and delta = ?", Bundle_, EntityId_, Delta_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByBundleAndEntityIdAndFieldImageTargetId Get NodeRevision_fieldImages via BundleAndEntityIdAndFieldImageTargetId
func GetNodeRevision_fieldImagesByBundleAndEntityIdAndFieldImageTargetId(offset int, limit int, Bundle_ string, EntityId_ int, FieldImageTargetId_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("bundle = ? and entity_id = ? and field_image_target_id = ?", Bundle_, EntityId_, FieldImageTargetId_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByBundleAndEntityIdAndFieldImageAlt Get NodeRevision_fieldImages via BundleAndEntityIdAndFieldImageAlt
func GetNodeRevision_fieldImagesByBundleAndEntityIdAndFieldImageAlt(offset int, limit int, Bundle_ string, EntityId_ int, FieldImageAlt_ string) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("bundle = ? and entity_id = ? and field_image_alt = ?", Bundle_, EntityId_, FieldImageAlt_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByBundleAndEntityIdAndFieldImageTitle Get NodeRevision_fieldImages via BundleAndEntityIdAndFieldImageTitle
func GetNodeRevision_fieldImagesByBundleAndEntityIdAndFieldImageTitle(offset int, limit int, Bundle_ string, EntityId_ int, FieldImageTitle_ string) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("bundle = ? and entity_id = ? and field_image_title = ?", Bundle_, EntityId_, FieldImageTitle_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByBundleAndEntityIdAndFieldImageWidth Get NodeRevision_fieldImages via BundleAndEntityIdAndFieldImageWidth
func GetNodeRevision_fieldImagesByBundleAndEntityIdAndFieldImageWidth(offset int, limit int, Bundle_ string, EntityId_ int, FieldImageWidth_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("bundle = ? and entity_id = ? and field_image_width = ?", Bundle_, EntityId_, FieldImageWidth_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByBundleAndEntityIdAndFieldImageHeight Get NodeRevision_fieldImages via BundleAndEntityIdAndFieldImageHeight
func GetNodeRevision_fieldImagesByBundleAndEntityIdAndFieldImageHeight(offset int, limit int, Bundle_ string, EntityId_ int, FieldImageHeight_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("bundle = ? and entity_id = ? and field_image_height = ?", Bundle_, EntityId_, FieldImageHeight_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByBundleAndRevisionIdAndLangcode Get NodeRevision_fieldImages via BundleAndRevisionIdAndLangcode
func GetNodeRevision_fieldImagesByBundleAndRevisionIdAndLangcode(offset int, limit int, Bundle_ string, RevisionId_ int, Langcode_ string) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("bundle = ? and revision_id = ? and langcode = ?", Bundle_, RevisionId_, Langcode_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByBundleAndRevisionIdAndDelta Get NodeRevision_fieldImages via BundleAndRevisionIdAndDelta
func GetNodeRevision_fieldImagesByBundleAndRevisionIdAndDelta(offset int, limit int, Bundle_ string, RevisionId_ int, Delta_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("bundle = ? and revision_id = ? and delta = ?", Bundle_, RevisionId_, Delta_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByBundleAndRevisionIdAndFieldImageTargetId Get NodeRevision_fieldImages via BundleAndRevisionIdAndFieldImageTargetId
func GetNodeRevision_fieldImagesByBundleAndRevisionIdAndFieldImageTargetId(offset int, limit int, Bundle_ string, RevisionId_ int, FieldImageTargetId_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("bundle = ? and revision_id = ? and field_image_target_id = ?", Bundle_, RevisionId_, FieldImageTargetId_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByBundleAndRevisionIdAndFieldImageAlt Get NodeRevision_fieldImages via BundleAndRevisionIdAndFieldImageAlt
func GetNodeRevision_fieldImagesByBundleAndRevisionIdAndFieldImageAlt(offset int, limit int, Bundle_ string, RevisionId_ int, FieldImageAlt_ string) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("bundle = ? and revision_id = ? and field_image_alt = ?", Bundle_, RevisionId_, FieldImageAlt_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByBundleAndRevisionIdAndFieldImageTitle Get NodeRevision_fieldImages via BundleAndRevisionIdAndFieldImageTitle
func GetNodeRevision_fieldImagesByBundleAndRevisionIdAndFieldImageTitle(offset int, limit int, Bundle_ string, RevisionId_ int, FieldImageTitle_ string) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("bundle = ? and revision_id = ? and field_image_title = ?", Bundle_, RevisionId_, FieldImageTitle_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByBundleAndRevisionIdAndFieldImageWidth Get NodeRevision_fieldImages via BundleAndRevisionIdAndFieldImageWidth
func GetNodeRevision_fieldImagesByBundleAndRevisionIdAndFieldImageWidth(offset int, limit int, Bundle_ string, RevisionId_ int, FieldImageWidth_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("bundle = ? and revision_id = ? and field_image_width = ?", Bundle_, RevisionId_, FieldImageWidth_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByBundleAndRevisionIdAndFieldImageHeight Get NodeRevision_fieldImages via BundleAndRevisionIdAndFieldImageHeight
func GetNodeRevision_fieldImagesByBundleAndRevisionIdAndFieldImageHeight(offset int, limit int, Bundle_ string, RevisionId_ int, FieldImageHeight_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("bundle = ? and revision_id = ? and field_image_height = ?", Bundle_, RevisionId_, FieldImageHeight_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByBundleAndLangcodeAndDelta Get NodeRevision_fieldImages via BundleAndLangcodeAndDelta
func GetNodeRevision_fieldImagesByBundleAndLangcodeAndDelta(offset int, limit int, Bundle_ string, Langcode_ string, Delta_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("bundle = ? and langcode = ? and delta = ?", Bundle_, Langcode_, Delta_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByBundleAndLangcodeAndFieldImageTargetId Get NodeRevision_fieldImages via BundleAndLangcodeAndFieldImageTargetId
func GetNodeRevision_fieldImagesByBundleAndLangcodeAndFieldImageTargetId(offset int, limit int, Bundle_ string, Langcode_ string, FieldImageTargetId_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("bundle = ? and langcode = ? and field_image_target_id = ?", Bundle_, Langcode_, FieldImageTargetId_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByBundleAndLangcodeAndFieldImageAlt Get NodeRevision_fieldImages via BundleAndLangcodeAndFieldImageAlt
func GetNodeRevision_fieldImagesByBundleAndLangcodeAndFieldImageAlt(offset int, limit int, Bundle_ string, Langcode_ string, FieldImageAlt_ string) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("bundle = ? and langcode = ? and field_image_alt = ?", Bundle_, Langcode_, FieldImageAlt_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByBundleAndLangcodeAndFieldImageTitle Get NodeRevision_fieldImages via BundleAndLangcodeAndFieldImageTitle
func GetNodeRevision_fieldImagesByBundleAndLangcodeAndFieldImageTitle(offset int, limit int, Bundle_ string, Langcode_ string, FieldImageTitle_ string) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("bundle = ? and langcode = ? and field_image_title = ?", Bundle_, Langcode_, FieldImageTitle_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByBundleAndLangcodeAndFieldImageWidth Get NodeRevision_fieldImages via BundleAndLangcodeAndFieldImageWidth
func GetNodeRevision_fieldImagesByBundleAndLangcodeAndFieldImageWidth(offset int, limit int, Bundle_ string, Langcode_ string, FieldImageWidth_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("bundle = ? and langcode = ? and field_image_width = ?", Bundle_, Langcode_, FieldImageWidth_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByBundleAndLangcodeAndFieldImageHeight Get NodeRevision_fieldImages via BundleAndLangcodeAndFieldImageHeight
func GetNodeRevision_fieldImagesByBundleAndLangcodeAndFieldImageHeight(offset int, limit int, Bundle_ string, Langcode_ string, FieldImageHeight_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("bundle = ? and langcode = ? and field_image_height = ?", Bundle_, Langcode_, FieldImageHeight_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByBundleAndDeltaAndFieldImageTargetId Get NodeRevision_fieldImages via BundleAndDeltaAndFieldImageTargetId
func GetNodeRevision_fieldImagesByBundleAndDeltaAndFieldImageTargetId(offset int, limit int, Bundle_ string, Delta_ int, FieldImageTargetId_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("bundle = ? and delta = ? and field_image_target_id = ?", Bundle_, Delta_, FieldImageTargetId_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByBundleAndDeltaAndFieldImageAlt Get NodeRevision_fieldImages via BundleAndDeltaAndFieldImageAlt
func GetNodeRevision_fieldImagesByBundleAndDeltaAndFieldImageAlt(offset int, limit int, Bundle_ string, Delta_ int, FieldImageAlt_ string) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("bundle = ? and delta = ? and field_image_alt = ?", Bundle_, Delta_, FieldImageAlt_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByBundleAndDeltaAndFieldImageTitle Get NodeRevision_fieldImages via BundleAndDeltaAndFieldImageTitle
func GetNodeRevision_fieldImagesByBundleAndDeltaAndFieldImageTitle(offset int, limit int, Bundle_ string, Delta_ int, FieldImageTitle_ string) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("bundle = ? and delta = ? and field_image_title = ?", Bundle_, Delta_, FieldImageTitle_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByBundleAndDeltaAndFieldImageWidth Get NodeRevision_fieldImages via BundleAndDeltaAndFieldImageWidth
func GetNodeRevision_fieldImagesByBundleAndDeltaAndFieldImageWidth(offset int, limit int, Bundle_ string, Delta_ int, FieldImageWidth_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("bundle = ? and delta = ? and field_image_width = ?", Bundle_, Delta_, FieldImageWidth_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByBundleAndDeltaAndFieldImageHeight Get NodeRevision_fieldImages via BundleAndDeltaAndFieldImageHeight
func GetNodeRevision_fieldImagesByBundleAndDeltaAndFieldImageHeight(offset int, limit int, Bundle_ string, Delta_ int, FieldImageHeight_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("bundle = ? and delta = ? and field_image_height = ?", Bundle_, Delta_, FieldImageHeight_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByBundleAndFieldImageTargetIdAndFieldImageAlt Get NodeRevision_fieldImages via BundleAndFieldImageTargetIdAndFieldImageAlt
func GetNodeRevision_fieldImagesByBundleAndFieldImageTargetIdAndFieldImageAlt(offset int, limit int, Bundle_ string, FieldImageTargetId_ int, FieldImageAlt_ string) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("bundle = ? and field_image_target_id = ? and field_image_alt = ?", Bundle_, FieldImageTargetId_, FieldImageAlt_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByBundleAndFieldImageTargetIdAndFieldImageTitle Get NodeRevision_fieldImages via BundleAndFieldImageTargetIdAndFieldImageTitle
func GetNodeRevision_fieldImagesByBundleAndFieldImageTargetIdAndFieldImageTitle(offset int, limit int, Bundle_ string, FieldImageTargetId_ int, FieldImageTitle_ string) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("bundle = ? and field_image_target_id = ? and field_image_title = ?", Bundle_, FieldImageTargetId_, FieldImageTitle_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByBundleAndFieldImageTargetIdAndFieldImageWidth Get NodeRevision_fieldImages via BundleAndFieldImageTargetIdAndFieldImageWidth
func GetNodeRevision_fieldImagesByBundleAndFieldImageTargetIdAndFieldImageWidth(offset int, limit int, Bundle_ string, FieldImageTargetId_ int, FieldImageWidth_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("bundle = ? and field_image_target_id = ? and field_image_width = ?", Bundle_, FieldImageTargetId_, FieldImageWidth_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByBundleAndFieldImageTargetIdAndFieldImageHeight Get NodeRevision_fieldImages via BundleAndFieldImageTargetIdAndFieldImageHeight
func GetNodeRevision_fieldImagesByBundleAndFieldImageTargetIdAndFieldImageHeight(offset int, limit int, Bundle_ string, FieldImageTargetId_ int, FieldImageHeight_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("bundle = ? and field_image_target_id = ? and field_image_height = ?", Bundle_, FieldImageTargetId_, FieldImageHeight_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByBundleAndFieldImageAltAndFieldImageTitle Get NodeRevision_fieldImages via BundleAndFieldImageAltAndFieldImageTitle
func GetNodeRevision_fieldImagesByBundleAndFieldImageAltAndFieldImageTitle(offset int, limit int, Bundle_ string, FieldImageAlt_ string, FieldImageTitle_ string) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("bundle = ? and field_image_alt = ? and field_image_title = ?", Bundle_, FieldImageAlt_, FieldImageTitle_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByBundleAndFieldImageAltAndFieldImageWidth Get NodeRevision_fieldImages via BundleAndFieldImageAltAndFieldImageWidth
func GetNodeRevision_fieldImagesByBundleAndFieldImageAltAndFieldImageWidth(offset int, limit int, Bundle_ string, FieldImageAlt_ string, FieldImageWidth_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("bundle = ? and field_image_alt = ? and field_image_width = ?", Bundle_, FieldImageAlt_, FieldImageWidth_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByBundleAndFieldImageAltAndFieldImageHeight Get NodeRevision_fieldImages via BundleAndFieldImageAltAndFieldImageHeight
func GetNodeRevision_fieldImagesByBundleAndFieldImageAltAndFieldImageHeight(offset int, limit int, Bundle_ string, FieldImageAlt_ string, FieldImageHeight_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("bundle = ? and field_image_alt = ? and field_image_height = ?", Bundle_, FieldImageAlt_, FieldImageHeight_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByBundleAndFieldImageTitleAndFieldImageWidth Get NodeRevision_fieldImages via BundleAndFieldImageTitleAndFieldImageWidth
func GetNodeRevision_fieldImagesByBundleAndFieldImageTitleAndFieldImageWidth(offset int, limit int, Bundle_ string, FieldImageTitle_ string, FieldImageWidth_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("bundle = ? and field_image_title = ? and field_image_width = ?", Bundle_, FieldImageTitle_, FieldImageWidth_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByBundleAndFieldImageTitleAndFieldImageHeight Get NodeRevision_fieldImages via BundleAndFieldImageTitleAndFieldImageHeight
func GetNodeRevision_fieldImagesByBundleAndFieldImageTitleAndFieldImageHeight(offset int, limit int, Bundle_ string, FieldImageTitle_ string, FieldImageHeight_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("bundle = ? and field_image_title = ? and field_image_height = ?", Bundle_, FieldImageTitle_, FieldImageHeight_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByBundleAndFieldImageWidthAndFieldImageHeight Get NodeRevision_fieldImages via BundleAndFieldImageWidthAndFieldImageHeight
func GetNodeRevision_fieldImagesByBundleAndFieldImageWidthAndFieldImageHeight(offset int, limit int, Bundle_ string, FieldImageWidth_ int, FieldImageHeight_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("bundle = ? and field_image_width = ? and field_image_height = ?", Bundle_, FieldImageWidth_, FieldImageHeight_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByDeletedAndEntityIdAndRevisionId Get NodeRevision_fieldImages via DeletedAndEntityIdAndRevisionId
func GetNodeRevision_fieldImagesByDeletedAndEntityIdAndRevisionId(offset int, limit int, Deleted_ int, EntityId_ int, RevisionId_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("deleted = ? and entity_id = ? and revision_id = ?", Deleted_, EntityId_, RevisionId_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByDeletedAndEntityIdAndLangcode Get NodeRevision_fieldImages via DeletedAndEntityIdAndLangcode
func GetNodeRevision_fieldImagesByDeletedAndEntityIdAndLangcode(offset int, limit int, Deleted_ int, EntityId_ int, Langcode_ string) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("deleted = ? and entity_id = ? and langcode = ?", Deleted_, EntityId_, Langcode_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByDeletedAndEntityIdAndDelta Get NodeRevision_fieldImages via DeletedAndEntityIdAndDelta
func GetNodeRevision_fieldImagesByDeletedAndEntityIdAndDelta(offset int, limit int, Deleted_ int, EntityId_ int, Delta_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("deleted = ? and entity_id = ? and delta = ?", Deleted_, EntityId_, Delta_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByDeletedAndEntityIdAndFieldImageTargetId Get NodeRevision_fieldImages via DeletedAndEntityIdAndFieldImageTargetId
func GetNodeRevision_fieldImagesByDeletedAndEntityIdAndFieldImageTargetId(offset int, limit int, Deleted_ int, EntityId_ int, FieldImageTargetId_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("deleted = ? and entity_id = ? and field_image_target_id = ?", Deleted_, EntityId_, FieldImageTargetId_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByDeletedAndEntityIdAndFieldImageAlt Get NodeRevision_fieldImages via DeletedAndEntityIdAndFieldImageAlt
func GetNodeRevision_fieldImagesByDeletedAndEntityIdAndFieldImageAlt(offset int, limit int, Deleted_ int, EntityId_ int, FieldImageAlt_ string) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("deleted = ? and entity_id = ? and field_image_alt = ?", Deleted_, EntityId_, FieldImageAlt_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByDeletedAndEntityIdAndFieldImageTitle Get NodeRevision_fieldImages via DeletedAndEntityIdAndFieldImageTitle
func GetNodeRevision_fieldImagesByDeletedAndEntityIdAndFieldImageTitle(offset int, limit int, Deleted_ int, EntityId_ int, FieldImageTitle_ string) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("deleted = ? and entity_id = ? and field_image_title = ?", Deleted_, EntityId_, FieldImageTitle_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByDeletedAndEntityIdAndFieldImageWidth Get NodeRevision_fieldImages via DeletedAndEntityIdAndFieldImageWidth
func GetNodeRevision_fieldImagesByDeletedAndEntityIdAndFieldImageWidth(offset int, limit int, Deleted_ int, EntityId_ int, FieldImageWidth_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("deleted = ? and entity_id = ? and field_image_width = ?", Deleted_, EntityId_, FieldImageWidth_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByDeletedAndEntityIdAndFieldImageHeight Get NodeRevision_fieldImages via DeletedAndEntityIdAndFieldImageHeight
func GetNodeRevision_fieldImagesByDeletedAndEntityIdAndFieldImageHeight(offset int, limit int, Deleted_ int, EntityId_ int, FieldImageHeight_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("deleted = ? and entity_id = ? and field_image_height = ?", Deleted_, EntityId_, FieldImageHeight_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByDeletedAndRevisionIdAndLangcode Get NodeRevision_fieldImages via DeletedAndRevisionIdAndLangcode
func GetNodeRevision_fieldImagesByDeletedAndRevisionIdAndLangcode(offset int, limit int, Deleted_ int, RevisionId_ int, Langcode_ string) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("deleted = ? and revision_id = ? and langcode = ?", Deleted_, RevisionId_, Langcode_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByDeletedAndRevisionIdAndDelta Get NodeRevision_fieldImages via DeletedAndRevisionIdAndDelta
func GetNodeRevision_fieldImagesByDeletedAndRevisionIdAndDelta(offset int, limit int, Deleted_ int, RevisionId_ int, Delta_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("deleted = ? and revision_id = ? and delta = ?", Deleted_, RevisionId_, Delta_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByDeletedAndRevisionIdAndFieldImageTargetId Get NodeRevision_fieldImages via DeletedAndRevisionIdAndFieldImageTargetId
func GetNodeRevision_fieldImagesByDeletedAndRevisionIdAndFieldImageTargetId(offset int, limit int, Deleted_ int, RevisionId_ int, FieldImageTargetId_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("deleted = ? and revision_id = ? and field_image_target_id = ?", Deleted_, RevisionId_, FieldImageTargetId_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByDeletedAndRevisionIdAndFieldImageAlt Get NodeRevision_fieldImages via DeletedAndRevisionIdAndFieldImageAlt
func GetNodeRevision_fieldImagesByDeletedAndRevisionIdAndFieldImageAlt(offset int, limit int, Deleted_ int, RevisionId_ int, FieldImageAlt_ string) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("deleted = ? and revision_id = ? and field_image_alt = ?", Deleted_, RevisionId_, FieldImageAlt_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByDeletedAndRevisionIdAndFieldImageTitle Get NodeRevision_fieldImages via DeletedAndRevisionIdAndFieldImageTitle
func GetNodeRevision_fieldImagesByDeletedAndRevisionIdAndFieldImageTitle(offset int, limit int, Deleted_ int, RevisionId_ int, FieldImageTitle_ string) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("deleted = ? and revision_id = ? and field_image_title = ?", Deleted_, RevisionId_, FieldImageTitle_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByDeletedAndRevisionIdAndFieldImageWidth Get NodeRevision_fieldImages via DeletedAndRevisionIdAndFieldImageWidth
func GetNodeRevision_fieldImagesByDeletedAndRevisionIdAndFieldImageWidth(offset int, limit int, Deleted_ int, RevisionId_ int, FieldImageWidth_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("deleted = ? and revision_id = ? and field_image_width = ?", Deleted_, RevisionId_, FieldImageWidth_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByDeletedAndRevisionIdAndFieldImageHeight Get NodeRevision_fieldImages via DeletedAndRevisionIdAndFieldImageHeight
func GetNodeRevision_fieldImagesByDeletedAndRevisionIdAndFieldImageHeight(offset int, limit int, Deleted_ int, RevisionId_ int, FieldImageHeight_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("deleted = ? and revision_id = ? and field_image_height = ?", Deleted_, RevisionId_, FieldImageHeight_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByDeletedAndLangcodeAndDelta Get NodeRevision_fieldImages via DeletedAndLangcodeAndDelta
func GetNodeRevision_fieldImagesByDeletedAndLangcodeAndDelta(offset int, limit int, Deleted_ int, Langcode_ string, Delta_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("deleted = ? and langcode = ? and delta = ?", Deleted_, Langcode_, Delta_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByDeletedAndLangcodeAndFieldImageTargetId Get NodeRevision_fieldImages via DeletedAndLangcodeAndFieldImageTargetId
func GetNodeRevision_fieldImagesByDeletedAndLangcodeAndFieldImageTargetId(offset int, limit int, Deleted_ int, Langcode_ string, FieldImageTargetId_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("deleted = ? and langcode = ? and field_image_target_id = ?", Deleted_, Langcode_, FieldImageTargetId_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByDeletedAndLangcodeAndFieldImageAlt Get NodeRevision_fieldImages via DeletedAndLangcodeAndFieldImageAlt
func GetNodeRevision_fieldImagesByDeletedAndLangcodeAndFieldImageAlt(offset int, limit int, Deleted_ int, Langcode_ string, FieldImageAlt_ string) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("deleted = ? and langcode = ? and field_image_alt = ?", Deleted_, Langcode_, FieldImageAlt_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByDeletedAndLangcodeAndFieldImageTitle Get NodeRevision_fieldImages via DeletedAndLangcodeAndFieldImageTitle
func GetNodeRevision_fieldImagesByDeletedAndLangcodeAndFieldImageTitle(offset int, limit int, Deleted_ int, Langcode_ string, FieldImageTitle_ string) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("deleted = ? and langcode = ? and field_image_title = ?", Deleted_, Langcode_, FieldImageTitle_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByDeletedAndLangcodeAndFieldImageWidth Get NodeRevision_fieldImages via DeletedAndLangcodeAndFieldImageWidth
func GetNodeRevision_fieldImagesByDeletedAndLangcodeAndFieldImageWidth(offset int, limit int, Deleted_ int, Langcode_ string, FieldImageWidth_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("deleted = ? and langcode = ? and field_image_width = ?", Deleted_, Langcode_, FieldImageWidth_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByDeletedAndLangcodeAndFieldImageHeight Get NodeRevision_fieldImages via DeletedAndLangcodeAndFieldImageHeight
func GetNodeRevision_fieldImagesByDeletedAndLangcodeAndFieldImageHeight(offset int, limit int, Deleted_ int, Langcode_ string, FieldImageHeight_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("deleted = ? and langcode = ? and field_image_height = ?", Deleted_, Langcode_, FieldImageHeight_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByDeletedAndDeltaAndFieldImageTargetId Get NodeRevision_fieldImages via DeletedAndDeltaAndFieldImageTargetId
func GetNodeRevision_fieldImagesByDeletedAndDeltaAndFieldImageTargetId(offset int, limit int, Deleted_ int, Delta_ int, FieldImageTargetId_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("deleted = ? and delta = ? and field_image_target_id = ?", Deleted_, Delta_, FieldImageTargetId_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByDeletedAndDeltaAndFieldImageAlt Get NodeRevision_fieldImages via DeletedAndDeltaAndFieldImageAlt
func GetNodeRevision_fieldImagesByDeletedAndDeltaAndFieldImageAlt(offset int, limit int, Deleted_ int, Delta_ int, FieldImageAlt_ string) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("deleted = ? and delta = ? and field_image_alt = ?", Deleted_, Delta_, FieldImageAlt_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByDeletedAndDeltaAndFieldImageTitle Get NodeRevision_fieldImages via DeletedAndDeltaAndFieldImageTitle
func GetNodeRevision_fieldImagesByDeletedAndDeltaAndFieldImageTitle(offset int, limit int, Deleted_ int, Delta_ int, FieldImageTitle_ string) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("deleted = ? and delta = ? and field_image_title = ?", Deleted_, Delta_, FieldImageTitle_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByDeletedAndDeltaAndFieldImageWidth Get NodeRevision_fieldImages via DeletedAndDeltaAndFieldImageWidth
func GetNodeRevision_fieldImagesByDeletedAndDeltaAndFieldImageWidth(offset int, limit int, Deleted_ int, Delta_ int, FieldImageWidth_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("deleted = ? and delta = ? and field_image_width = ?", Deleted_, Delta_, FieldImageWidth_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByDeletedAndDeltaAndFieldImageHeight Get NodeRevision_fieldImages via DeletedAndDeltaAndFieldImageHeight
func GetNodeRevision_fieldImagesByDeletedAndDeltaAndFieldImageHeight(offset int, limit int, Deleted_ int, Delta_ int, FieldImageHeight_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("deleted = ? and delta = ? and field_image_height = ?", Deleted_, Delta_, FieldImageHeight_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByDeletedAndFieldImageTargetIdAndFieldImageAlt Get NodeRevision_fieldImages via DeletedAndFieldImageTargetIdAndFieldImageAlt
func GetNodeRevision_fieldImagesByDeletedAndFieldImageTargetIdAndFieldImageAlt(offset int, limit int, Deleted_ int, FieldImageTargetId_ int, FieldImageAlt_ string) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("deleted = ? and field_image_target_id = ? and field_image_alt = ?", Deleted_, FieldImageTargetId_, FieldImageAlt_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByDeletedAndFieldImageTargetIdAndFieldImageTitle Get NodeRevision_fieldImages via DeletedAndFieldImageTargetIdAndFieldImageTitle
func GetNodeRevision_fieldImagesByDeletedAndFieldImageTargetIdAndFieldImageTitle(offset int, limit int, Deleted_ int, FieldImageTargetId_ int, FieldImageTitle_ string) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("deleted = ? and field_image_target_id = ? and field_image_title = ?", Deleted_, FieldImageTargetId_, FieldImageTitle_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByDeletedAndFieldImageTargetIdAndFieldImageWidth Get NodeRevision_fieldImages via DeletedAndFieldImageTargetIdAndFieldImageWidth
func GetNodeRevision_fieldImagesByDeletedAndFieldImageTargetIdAndFieldImageWidth(offset int, limit int, Deleted_ int, FieldImageTargetId_ int, FieldImageWidth_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("deleted = ? and field_image_target_id = ? and field_image_width = ?", Deleted_, FieldImageTargetId_, FieldImageWidth_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByDeletedAndFieldImageTargetIdAndFieldImageHeight Get NodeRevision_fieldImages via DeletedAndFieldImageTargetIdAndFieldImageHeight
func GetNodeRevision_fieldImagesByDeletedAndFieldImageTargetIdAndFieldImageHeight(offset int, limit int, Deleted_ int, FieldImageTargetId_ int, FieldImageHeight_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("deleted = ? and field_image_target_id = ? and field_image_height = ?", Deleted_, FieldImageTargetId_, FieldImageHeight_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByDeletedAndFieldImageAltAndFieldImageTitle Get NodeRevision_fieldImages via DeletedAndFieldImageAltAndFieldImageTitle
func GetNodeRevision_fieldImagesByDeletedAndFieldImageAltAndFieldImageTitle(offset int, limit int, Deleted_ int, FieldImageAlt_ string, FieldImageTitle_ string) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("deleted = ? and field_image_alt = ? and field_image_title = ?", Deleted_, FieldImageAlt_, FieldImageTitle_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByDeletedAndFieldImageAltAndFieldImageWidth Get NodeRevision_fieldImages via DeletedAndFieldImageAltAndFieldImageWidth
func GetNodeRevision_fieldImagesByDeletedAndFieldImageAltAndFieldImageWidth(offset int, limit int, Deleted_ int, FieldImageAlt_ string, FieldImageWidth_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("deleted = ? and field_image_alt = ? and field_image_width = ?", Deleted_, FieldImageAlt_, FieldImageWidth_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByDeletedAndFieldImageAltAndFieldImageHeight Get NodeRevision_fieldImages via DeletedAndFieldImageAltAndFieldImageHeight
func GetNodeRevision_fieldImagesByDeletedAndFieldImageAltAndFieldImageHeight(offset int, limit int, Deleted_ int, FieldImageAlt_ string, FieldImageHeight_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("deleted = ? and field_image_alt = ? and field_image_height = ?", Deleted_, FieldImageAlt_, FieldImageHeight_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByDeletedAndFieldImageTitleAndFieldImageWidth Get NodeRevision_fieldImages via DeletedAndFieldImageTitleAndFieldImageWidth
func GetNodeRevision_fieldImagesByDeletedAndFieldImageTitleAndFieldImageWidth(offset int, limit int, Deleted_ int, FieldImageTitle_ string, FieldImageWidth_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("deleted = ? and field_image_title = ? and field_image_width = ?", Deleted_, FieldImageTitle_, FieldImageWidth_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByDeletedAndFieldImageTitleAndFieldImageHeight Get NodeRevision_fieldImages via DeletedAndFieldImageTitleAndFieldImageHeight
func GetNodeRevision_fieldImagesByDeletedAndFieldImageTitleAndFieldImageHeight(offset int, limit int, Deleted_ int, FieldImageTitle_ string, FieldImageHeight_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("deleted = ? and field_image_title = ? and field_image_height = ?", Deleted_, FieldImageTitle_, FieldImageHeight_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByDeletedAndFieldImageWidthAndFieldImageHeight Get NodeRevision_fieldImages via DeletedAndFieldImageWidthAndFieldImageHeight
func GetNodeRevision_fieldImagesByDeletedAndFieldImageWidthAndFieldImageHeight(offset int, limit int, Deleted_ int, FieldImageWidth_ int, FieldImageHeight_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("deleted = ? and field_image_width = ? and field_image_height = ?", Deleted_, FieldImageWidth_, FieldImageHeight_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByEntityIdAndRevisionIdAndLangcode Get NodeRevision_fieldImages via EntityIdAndRevisionIdAndLangcode
func GetNodeRevision_fieldImagesByEntityIdAndRevisionIdAndLangcode(offset int, limit int, EntityId_ int, RevisionId_ int, Langcode_ string) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("entity_id = ? and revision_id = ? and langcode = ?", EntityId_, RevisionId_, Langcode_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByEntityIdAndRevisionIdAndDelta Get NodeRevision_fieldImages via EntityIdAndRevisionIdAndDelta
func GetNodeRevision_fieldImagesByEntityIdAndRevisionIdAndDelta(offset int, limit int, EntityId_ int, RevisionId_ int, Delta_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("entity_id = ? and revision_id = ? and delta = ?", EntityId_, RevisionId_, Delta_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByEntityIdAndRevisionIdAndFieldImageTargetId Get NodeRevision_fieldImages via EntityIdAndRevisionIdAndFieldImageTargetId
func GetNodeRevision_fieldImagesByEntityIdAndRevisionIdAndFieldImageTargetId(offset int, limit int, EntityId_ int, RevisionId_ int, FieldImageTargetId_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("entity_id = ? and revision_id = ? and field_image_target_id = ?", EntityId_, RevisionId_, FieldImageTargetId_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByEntityIdAndRevisionIdAndFieldImageAlt Get NodeRevision_fieldImages via EntityIdAndRevisionIdAndFieldImageAlt
func GetNodeRevision_fieldImagesByEntityIdAndRevisionIdAndFieldImageAlt(offset int, limit int, EntityId_ int, RevisionId_ int, FieldImageAlt_ string) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("entity_id = ? and revision_id = ? and field_image_alt = ?", EntityId_, RevisionId_, FieldImageAlt_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByEntityIdAndRevisionIdAndFieldImageTitle Get NodeRevision_fieldImages via EntityIdAndRevisionIdAndFieldImageTitle
func GetNodeRevision_fieldImagesByEntityIdAndRevisionIdAndFieldImageTitle(offset int, limit int, EntityId_ int, RevisionId_ int, FieldImageTitle_ string) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("entity_id = ? and revision_id = ? and field_image_title = ?", EntityId_, RevisionId_, FieldImageTitle_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByEntityIdAndRevisionIdAndFieldImageWidth Get NodeRevision_fieldImages via EntityIdAndRevisionIdAndFieldImageWidth
func GetNodeRevision_fieldImagesByEntityIdAndRevisionIdAndFieldImageWidth(offset int, limit int, EntityId_ int, RevisionId_ int, FieldImageWidth_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("entity_id = ? and revision_id = ? and field_image_width = ?", EntityId_, RevisionId_, FieldImageWidth_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByEntityIdAndRevisionIdAndFieldImageHeight Get NodeRevision_fieldImages via EntityIdAndRevisionIdAndFieldImageHeight
func GetNodeRevision_fieldImagesByEntityIdAndRevisionIdAndFieldImageHeight(offset int, limit int, EntityId_ int, RevisionId_ int, FieldImageHeight_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("entity_id = ? and revision_id = ? and field_image_height = ?", EntityId_, RevisionId_, FieldImageHeight_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByEntityIdAndLangcodeAndDelta Get NodeRevision_fieldImages via EntityIdAndLangcodeAndDelta
func GetNodeRevision_fieldImagesByEntityIdAndLangcodeAndDelta(offset int, limit int, EntityId_ int, Langcode_ string, Delta_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("entity_id = ? and langcode = ? and delta = ?", EntityId_, Langcode_, Delta_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByEntityIdAndLangcodeAndFieldImageTargetId Get NodeRevision_fieldImages via EntityIdAndLangcodeAndFieldImageTargetId
func GetNodeRevision_fieldImagesByEntityIdAndLangcodeAndFieldImageTargetId(offset int, limit int, EntityId_ int, Langcode_ string, FieldImageTargetId_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("entity_id = ? and langcode = ? and field_image_target_id = ?", EntityId_, Langcode_, FieldImageTargetId_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByEntityIdAndLangcodeAndFieldImageAlt Get NodeRevision_fieldImages via EntityIdAndLangcodeAndFieldImageAlt
func GetNodeRevision_fieldImagesByEntityIdAndLangcodeAndFieldImageAlt(offset int, limit int, EntityId_ int, Langcode_ string, FieldImageAlt_ string) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("entity_id = ? and langcode = ? and field_image_alt = ?", EntityId_, Langcode_, FieldImageAlt_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByEntityIdAndLangcodeAndFieldImageTitle Get NodeRevision_fieldImages via EntityIdAndLangcodeAndFieldImageTitle
func GetNodeRevision_fieldImagesByEntityIdAndLangcodeAndFieldImageTitle(offset int, limit int, EntityId_ int, Langcode_ string, FieldImageTitle_ string) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("entity_id = ? and langcode = ? and field_image_title = ?", EntityId_, Langcode_, FieldImageTitle_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByEntityIdAndLangcodeAndFieldImageWidth Get NodeRevision_fieldImages via EntityIdAndLangcodeAndFieldImageWidth
func GetNodeRevision_fieldImagesByEntityIdAndLangcodeAndFieldImageWidth(offset int, limit int, EntityId_ int, Langcode_ string, FieldImageWidth_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("entity_id = ? and langcode = ? and field_image_width = ?", EntityId_, Langcode_, FieldImageWidth_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByEntityIdAndLangcodeAndFieldImageHeight Get NodeRevision_fieldImages via EntityIdAndLangcodeAndFieldImageHeight
func GetNodeRevision_fieldImagesByEntityIdAndLangcodeAndFieldImageHeight(offset int, limit int, EntityId_ int, Langcode_ string, FieldImageHeight_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("entity_id = ? and langcode = ? and field_image_height = ?", EntityId_, Langcode_, FieldImageHeight_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByEntityIdAndDeltaAndFieldImageTargetId Get NodeRevision_fieldImages via EntityIdAndDeltaAndFieldImageTargetId
func GetNodeRevision_fieldImagesByEntityIdAndDeltaAndFieldImageTargetId(offset int, limit int, EntityId_ int, Delta_ int, FieldImageTargetId_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("entity_id = ? and delta = ? and field_image_target_id = ?", EntityId_, Delta_, FieldImageTargetId_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByEntityIdAndDeltaAndFieldImageAlt Get NodeRevision_fieldImages via EntityIdAndDeltaAndFieldImageAlt
func GetNodeRevision_fieldImagesByEntityIdAndDeltaAndFieldImageAlt(offset int, limit int, EntityId_ int, Delta_ int, FieldImageAlt_ string) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("entity_id = ? and delta = ? and field_image_alt = ?", EntityId_, Delta_, FieldImageAlt_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByEntityIdAndDeltaAndFieldImageTitle Get NodeRevision_fieldImages via EntityIdAndDeltaAndFieldImageTitle
func GetNodeRevision_fieldImagesByEntityIdAndDeltaAndFieldImageTitle(offset int, limit int, EntityId_ int, Delta_ int, FieldImageTitle_ string) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("entity_id = ? and delta = ? and field_image_title = ?", EntityId_, Delta_, FieldImageTitle_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByEntityIdAndDeltaAndFieldImageWidth Get NodeRevision_fieldImages via EntityIdAndDeltaAndFieldImageWidth
func GetNodeRevision_fieldImagesByEntityIdAndDeltaAndFieldImageWidth(offset int, limit int, EntityId_ int, Delta_ int, FieldImageWidth_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("entity_id = ? and delta = ? and field_image_width = ?", EntityId_, Delta_, FieldImageWidth_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByEntityIdAndDeltaAndFieldImageHeight Get NodeRevision_fieldImages via EntityIdAndDeltaAndFieldImageHeight
func GetNodeRevision_fieldImagesByEntityIdAndDeltaAndFieldImageHeight(offset int, limit int, EntityId_ int, Delta_ int, FieldImageHeight_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("entity_id = ? and delta = ? and field_image_height = ?", EntityId_, Delta_, FieldImageHeight_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByEntityIdAndFieldImageTargetIdAndFieldImageAlt Get NodeRevision_fieldImages via EntityIdAndFieldImageTargetIdAndFieldImageAlt
func GetNodeRevision_fieldImagesByEntityIdAndFieldImageTargetIdAndFieldImageAlt(offset int, limit int, EntityId_ int, FieldImageTargetId_ int, FieldImageAlt_ string) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("entity_id = ? and field_image_target_id = ? and field_image_alt = ?", EntityId_, FieldImageTargetId_, FieldImageAlt_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByEntityIdAndFieldImageTargetIdAndFieldImageTitle Get NodeRevision_fieldImages via EntityIdAndFieldImageTargetIdAndFieldImageTitle
func GetNodeRevision_fieldImagesByEntityIdAndFieldImageTargetIdAndFieldImageTitle(offset int, limit int, EntityId_ int, FieldImageTargetId_ int, FieldImageTitle_ string) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("entity_id = ? and field_image_target_id = ? and field_image_title = ?", EntityId_, FieldImageTargetId_, FieldImageTitle_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByEntityIdAndFieldImageTargetIdAndFieldImageWidth Get NodeRevision_fieldImages via EntityIdAndFieldImageTargetIdAndFieldImageWidth
func GetNodeRevision_fieldImagesByEntityIdAndFieldImageTargetIdAndFieldImageWidth(offset int, limit int, EntityId_ int, FieldImageTargetId_ int, FieldImageWidth_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("entity_id = ? and field_image_target_id = ? and field_image_width = ?", EntityId_, FieldImageTargetId_, FieldImageWidth_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByEntityIdAndFieldImageTargetIdAndFieldImageHeight Get NodeRevision_fieldImages via EntityIdAndFieldImageTargetIdAndFieldImageHeight
func GetNodeRevision_fieldImagesByEntityIdAndFieldImageTargetIdAndFieldImageHeight(offset int, limit int, EntityId_ int, FieldImageTargetId_ int, FieldImageHeight_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("entity_id = ? and field_image_target_id = ? and field_image_height = ?", EntityId_, FieldImageTargetId_, FieldImageHeight_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByEntityIdAndFieldImageAltAndFieldImageTitle Get NodeRevision_fieldImages via EntityIdAndFieldImageAltAndFieldImageTitle
func GetNodeRevision_fieldImagesByEntityIdAndFieldImageAltAndFieldImageTitle(offset int, limit int, EntityId_ int, FieldImageAlt_ string, FieldImageTitle_ string) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("entity_id = ? and field_image_alt = ? and field_image_title = ?", EntityId_, FieldImageAlt_, FieldImageTitle_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByEntityIdAndFieldImageAltAndFieldImageWidth Get NodeRevision_fieldImages via EntityIdAndFieldImageAltAndFieldImageWidth
func GetNodeRevision_fieldImagesByEntityIdAndFieldImageAltAndFieldImageWidth(offset int, limit int, EntityId_ int, FieldImageAlt_ string, FieldImageWidth_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("entity_id = ? and field_image_alt = ? and field_image_width = ?", EntityId_, FieldImageAlt_, FieldImageWidth_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByEntityIdAndFieldImageAltAndFieldImageHeight Get NodeRevision_fieldImages via EntityIdAndFieldImageAltAndFieldImageHeight
func GetNodeRevision_fieldImagesByEntityIdAndFieldImageAltAndFieldImageHeight(offset int, limit int, EntityId_ int, FieldImageAlt_ string, FieldImageHeight_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("entity_id = ? and field_image_alt = ? and field_image_height = ?", EntityId_, FieldImageAlt_, FieldImageHeight_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByEntityIdAndFieldImageTitleAndFieldImageWidth Get NodeRevision_fieldImages via EntityIdAndFieldImageTitleAndFieldImageWidth
func GetNodeRevision_fieldImagesByEntityIdAndFieldImageTitleAndFieldImageWidth(offset int, limit int, EntityId_ int, FieldImageTitle_ string, FieldImageWidth_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("entity_id = ? and field_image_title = ? and field_image_width = ?", EntityId_, FieldImageTitle_, FieldImageWidth_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByEntityIdAndFieldImageTitleAndFieldImageHeight Get NodeRevision_fieldImages via EntityIdAndFieldImageTitleAndFieldImageHeight
func GetNodeRevision_fieldImagesByEntityIdAndFieldImageTitleAndFieldImageHeight(offset int, limit int, EntityId_ int, FieldImageTitle_ string, FieldImageHeight_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("entity_id = ? and field_image_title = ? and field_image_height = ?", EntityId_, FieldImageTitle_, FieldImageHeight_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByEntityIdAndFieldImageWidthAndFieldImageHeight Get NodeRevision_fieldImages via EntityIdAndFieldImageWidthAndFieldImageHeight
func GetNodeRevision_fieldImagesByEntityIdAndFieldImageWidthAndFieldImageHeight(offset int, limit int, EntityId_ int, FieldImageWidth_ int, FieldImageHeight_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("entity_id = ? and field_image_width = ? and field_image_height = ?", EntityId_, FieldImageWidth_, FieldImageHeight_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByRevisionIdAndLangcodeAndDelta Get NodeRevision_fieldImages via RevisionIdAndLangcodeAndDelta
func GetNodeRevision_fieldImagesByRevisionIdAndLangcodeAndDelta(offset int, limit int, RevisionId_ int, Langcode_ string, Delta_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("revision_id = ? and langcode = ? and delta = ?", RevisionId_, Langcode_, Delta_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByRevisionIdAndLangcodeAndFieldImageTargetId Get NodeRevision_fieldImages via RevisionIdAndLangcodeAndFieldImageTargetId
func GetNodeRevision_fieldImagesByRevisionIdAndLangcodeAndFieldImageTargetId(offset int, limit int, RevisionId_ int, Langcode_ string, FieldImageTargetId_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("revision_id = ? and langcode = ? and field_image_target_id = ?", RevisionId_, Langcode_, FieldImageTargetId_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByRevisionIdAndLangcodeAndFieldImageAlt Get NodeRevision_fieldImages via RevisionIdAndLangcodeAndFieldImageAlt
func GetNodeRevision_fieldImagesByRevisionIdAndLangcodeAndFieldImageAlt(offset int, limit int, RevisionId_ int, Langcode_ string, FieldImageAlt_ string) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("revision_id = ? and langcode = ? and field_image_alt = ?", RevisionId_, Langcode_, FieldImageAlt_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByRevisionIdAndLangcodeAndFieldImageTitle Get NodeRevision_fieldImages via RevisionIdAndLangcodeAndFieldImageTitle
func GetNodeRevision_fieldImagesByRevisionIdAndLangcodeAndFieldImageTitle(offset int, limit int, RevisionId_ int, Langcode_ string, FieldImageTitle_ string) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("revision_id = ? and langcode = ? and field_image_title = ?", RevisionId_, Langcode_, FieldImageTitle_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByRevisionIdAndLangcodeAndFieldImageWidth Get NodeRevision_fieldImages via RevisionIdAndLangcodeAndFieldImageWidth
func GetNodeRevision_fieldImagesByRevisionIdAndLangcodeAndFieldImageWidth(offset int, limit int, RevisionId_ int, Langcode_ string, FieldImageWidth_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("revision_id = ? and langcode = ? and field_image_width = ?", RevisionId_, Langcode_, FieldImageWidth_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByRevisionIdAndLangcodeAndFieldImageHeight Get NodeRevision_fieldImages via RevisionIdAndLangcodeAndFieldImageHeight
func GetNodeRevision_fieldImagesByRevisionIdAndLangcodeAndFieldImageHeight(offset int, limit int, RevisionId_ int, Langcode_ string, FieldImageHeight_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("revision_id = ? and langcode = ? and field_image_height = ?", RevisionId_, Langcode_, FieldImageHeight_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByRevisionIdAndDeltaAndFieldImageTargetId Get NodeRevision_fieldImages via RevisionIdAndDeltaAndFieldImageTargetId
func GetNodeRevision_fieldImagesByRevisionIdAndDeltaAndFieldImageTargetId(offset int, limit int, RevisionId_ int, Delta_ int, FieldImageTargetId_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("revision_id = ? and delta = ? and field_image_target_id = ?", RevisionId_, Delta_, FieldImageTargetId_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByRevisionIdAndDeltaAndFieldImageAlt Get NodeRevision_fieldImages via RevisionIdAndDeltaAndFieldImageAlt
func GetNodeRevision_fieldImagesByRevisionIdAndDeltaAndFieldImageAlt(offset int, limit int, RevisionId_ int, Delta_ int, FieldImageAlt_ string) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("revision_id = ? and delta = ? and field_image_alt = ?", RevisionId_, Delta_, FieldImageAlt_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByRevisionIdAndDeltaAndFieldImageTitle Get NodeRevision_fieldImages via RevisionIdAndDeltaAndFieldImageTitle
func GetNodeRevision_fieldImagesByRevisionIdAndDeltaAndFieldImageTitle(offset int, limit int, RevisionId_ int, Delta_ int, FieldImageTitle_ string) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("revision_id = ? and delta = ? and field_image_title = ?", RevisionId_, Delta_, FieldImageTitle_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByRevisionIdAndDeltaAndFieldImageWidth Get NodeRevision_fieldImages via RevisionIdAndDeltaAndFieldImageWidth
func GetNodeRevision_fieldImagesByRevisionIdAndDeltaAndFieldImageWidth(offset int, limit int, RevisionId_ int, Delta_ int, FieldImageWidth_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("revision_id = ? and delta = ? and field_image_width = ?", RevisionId_, Delta_, FieldImageWidth_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByRevisionIdAndDeltaAndFieldImageHeight Get NodeRevision_fieldImages via RevisionIdAndDeltaAndFieldImageHeight
func GetNodeRevision_fieldImagesByRevisionIdAndDeltaAndFieldImageHeight(offset int, limit int, RevisionId_ int, Delta_ int, FieldImageHeight_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("revision_id = ? and delta = ? and field_image_height = ?", RevisionId_, Delta_, FieldImageHeight_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByRevisionIdAndFieldImageTargetIdAndFieldImageAlt Get NodeRevision_fieldImages via RevisionIdAndFieldImageTargetIdAndFieldImageAlt
func GetNodeRevision_fieldImagesByRevisionIdAndFieldImageTargetIdAndFieldImageAlt(offset int, limit int, RevisionId_ int, FieldImageTargetId_ int, FieldImageAlt_ string) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("revision_id = ? and field_image_target_id = ? and field_image_alt = ?", RevisionId_, FieldImageTargetId_, FieldImageAlt_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByRevisionIdAndFieldImageTargetIdAndFieldImageTitle Get NodeRevision_fieldImages via RevisionIdAndFieldImageTargetIdAndFieldImageTitle
func GetNodeRevision_fieldImagesByRevisionIdAndFieldImageTargetIdAndFieldImageTitle(offset int, limit int, RevisionId_ int, FieldImageTargetId_ int, FieldImageTitle_ string) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("revision_id = ? and field_image_target_id = ? and field_image_title = ?", RevisionId_, FieldImageTargetId_, FieldImageTitle_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByRevisionIdAndFieldImageTargetIdAndFieldImageWidth Get NodeRevision_fieldImages via RevisionIdAndFieldImageTargetIdAndFieldImageWidth
func GetNodeRevision_fieldImagesByRevisionIdAndFieldImageTargetIdAndFieldImageWidth(offset int, limit int, RevisionId_ int, FieldImageTargetId_ int, FieldImageWidth_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("revision_id = ? and field_image_target_id = ? and field_image_width = ?", RevisionId_, FieldImageTargetId_, FieldImageWidth_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByRevisionIdAndFieldImageTargetIdAndFieldImageHeight Get NodeRevision_fieldImages via RevisionIdAndFieldImageTargetIdAndFieldImageHeight
func GetNodeRevision_fieldImagesByRevisionIdAndFieldImageTargetIdAndFieldImageHeight(offset int, limit int, RevisionId_ int, FieldImageTargetId_ int, FieldImageHeight_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("revision_id = ? and field_image_target_id = ? and field_image_height = ?", RevisionId_, FieldImageTargetId_, FieldImageHeight_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByRevisionIdAndFieldImageAltAndFieldImageTitle Get NodeRevision_fieldImages via RevisionIdAndFieldImageAltAndFieldImageTitle
func GetNodeRevision_fieldImagesByRevisionIdAndFieldImageAltAndFieldImageTitle(offset int, limit int, RevisionId_ int, FieldImageAlt_ string, FieldImageTitle_ string) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("revision_id = ? and field_image_alt = ? and field_image_title = ?", RevisionId_, FieldImageAlt_, FieldImageTitle_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByRevisionIdAndFieldImageAltAndFieldImageWidth Get NodeRevision_fieldImages via RevisionIdAndFieldImageAltAndFieldImageWidth
func GetNodeRevision_fieldImagesByRevisionIdAndFieldImageAltAndFieldImageWidth(offset int, limit int, RevisionId_ int, FieldImageAlt_ string, FieldImageWidth_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("revision_id = ? and field_image_alt = ? and field_image_width = ?", RevisionId_, FieldImageAlt_, FieldImageWidth_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByRevisionIdAndFieldImageAltAndFieldImageHeight Get NodeRevision_fieldImages via RevisionIdAndFieldImageAltAndFieldImageHeight
func GetNodeRevision_fieldImagesByRevisionIdAndFieldImageAltAndFieldImageHeight(offset int, limit int, RevisionId_ int, FieldImageAlt_ string, FieldImageHeight_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("revision_id = ? and field_image_alt = ? and field_image_height = ?", RevisionId_, FieldImageAlt_, FieldImageHeight_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByRevisionIdAndFieldImageTitleAndFieldImageWidth Get NodeRevision_fieldImages via RevisionIdAndFieldImageTitleAndFieldImageWidth
func GetNodeRevision_fieldImagesByRevisionIdAndFieldImageTitleAndFieldImageWidth(offset int, limit int, RevisionId_ int, FieldImageTitle_ string, FieldImageWidth_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("revision_id = ? and field_image_title = ? and field_image_width = ?", RevisionId_, FieldImageTitle_, FieldImageWidth_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByRevisionIdAndFieldImageTitleAndFieldImageHeight Get NodeRevision_fieldImages via RevisionIdAndFieldImageTitleAndFieldImageHeight
func GetNodeRevision_fieldImagesByRevisionIdAndFieldImageTitleAndFieldImageHeight(offset int, limit int, RevisionId_ int, FieldImageTitle_ string, FieldImageHeight_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("revision_id = ? and field_image_title = ? and field_image_height = ?", RevisionId_, FieldImageTitle_, FieldImageHeight_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByRevisionIdAndFieldImageWidthAndFieldImageHeight Get NodeRevision_fieldImages via RevisionIdAndFieldImageWidthAndFieldImageHeight
func GetNodeRevision_fieldImagesByRevisionIdAndFieldImageWidthAndFieldImageHeight(offset int, limit int, RevisionId_ int, FieldImageWidth_ int, FieldImageHeight_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("revision_id = ? and field_image_width = ? and field_image_height = ?", RevisionId_, FieldImageWidth_, FieldImageHeight_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByLangcodeAndDeltaAndFieldImageTargetId Get NodeRevision_fieldImages via LangcodeAndDeltaAndFieldImageTargetId
func GetNodeRevision_fieldImagesByLangcodeAndDeltaAndFieldImageTargetId(offset int, limit int, Langcode_ string, Delta_ int, FieldImageTargetId_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("langcode = ? and delta = ? and field_image_target_id = ?", Langcode_, Delta_, FieldImageTargetId_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByLangcodeAndDeltaAndFieldImageAlt Get NodeRevision_fieldImages via LangcodeAndDeltaAndFieldImageAlt
func GetNodeRevision_fieldImagesByLangcodeAndDeltaAndFieldImageAlt(offset int, limit int, Langcode_ string, Delta_ int, FieldImageAlt_ string) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("langcode = ? and delta = ? and field_image_alt = ?", Langcode_, Delta_, FieldImageAlt_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByLangcodeAndDeltaAndFieldImageTitle Get NodeRevision_fieldImages via LangcodeAndDeltaAndFieldImageTitle
func GetNodeRevision_fieldImagesByLangcodeAndDeltaAndFieldImageTitle(offset int, limit int, Langcode_ string, Delta_ int, FieldImageTitle_ string) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("langcode = ? and delta = ? and field_image_title = ?", Langcode_, Delta_, FieldImageTitle_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByLangcodeAndDeltaAndFieldImageWidth Get NodeRevision_fieldImages via LangcodeAndDeltaAndFieldImageWidth
func GetNodeRevision_fieldImagesByLangcodeAndDeltaAndFieldImageWidth(offset int, limit int, Langcode_ string, Delta_ int, FieldImageWidth_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("langcode = ? and delta = ? and field_image_width = ?", Langcode_, Delta_, FieldImageWidth_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByLangcodeAndDeltaAndFieldImageHeight Get NodeRevision_fieldImages via LangcodeAndDeltaAndFieldImageHeight
func GetNodeRevision_fieldImagesByLangcodeAndDeltaAndFieldImageHeight(offset int, limit int, Langcode_ string, Delta_ int, FieldImageHeight_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("langcode = ? and delta = ? and field_image_height = ?", Langcode_, Delta_, FieldImageHeight_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByLangcodeAndFieldImageTargetIdAndFieldImageAlt Get NodeRevision_fieldImages via LangcodeAndFieldImageTargetIdAndFieldImageAlt
func GetNodeRevision_fieldImagesByLangcodeAndFieldImageTargetIdAndFieldImageAlt(offset int, limit int, Langcode_ string, FieldImageTargetId_ int, FieldImageAlt_ string) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("langcode = ? and field_image_target_id = ? and field_image_alt = ?", Langcode_, FieldImageTargetId_, FieldImageAlt_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByLangcodeAndFieldImageTargetIdAndFieldImageTitle Get NodeRevision_fieldImages via LangcodeAndFieldImageTargetIdAndFieldImageTitle
func GetNodeRevision_fieldImagesByLangcodeAndFieldImageTargetIdAndFieldImageTitle(offset int, limit int, Langcode_ string, FieldImageTargetId_ int, FieldImageTitle_ string) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("langcode = ? and field_image_target_id = ? and field_image_title = ?", Langcode_, FieldImageTargetId_, FieldImageTitle_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByLangcodeAndFieldImageTargetIdAndFieldImageWidth Get NodeRevision_fieldImages via LangcodeAndFieldImageTargetIdAndFieldImageWidth
func GetNodeRevision_fieldImagesByLangcodeAndFieldImageTargetIdAndFieldImageWidth(offset int, limit int, Langcode_ string, FieldImageTargetId_ int, FieldImageWidth_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("langcode = ? and field_image_target_id = ? and field_image_width = ?", Langcode_, FieldImageTargetId_, FieldImageWidth_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByLangcodeAndFieldImageTargetIdAndFieldImageHeight Get NodeRevision_fieldImages via LangcodeAndFieldImageTargetIdAndFieldImageHeight
func GetNodeRevision_fieldImagesByLangcodeAndFieldImageTargetIdAndFieldImageHeight(offset int, limit int, Langcode_ string, FieldImageTargetId_ int, FieldImageHeight_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("langcode = ? and field_image_target_id = ? and field_image_height = ?", Langcode_, FieldImageTargetId_, FieldImageHeight_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByLangcodeAndFieldImageAltAndFieldImageTitle Get NodeRevision_fieldImages via LangcodeAndFieldImageAltAndFieldImageTitle
func GetNodeRevision_fieldImagesByLangcodeAndFieldImageAltAndFieldImageTitle(offset int, limit int, Langcode_ string, FieldImageAlt_ string, FieldImageTitle_ string) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("langcode = ? and field_image_alt = ? and field_image_title = ?", Langcode_, FieldImageAlt_, FieldImageTitle_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByLangcodeAndFieldImageAltAndFieldImageWidth Get NodeRevision_fieldImages via LangcodeAndFieldImageAltAndFieldImageWidth
func GetNodeRevision_fieldImagesByLangcodeAndFieldImageAltAndFieldImageWidth(offset int, limit int, Langcode_ string, FieldImageAlt_ string, FieldImageWidth_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("langcode = ? and field_image_alt = ? and field_image_width = ?", Langcode_, FieldImageAlt_, FieldImageWidth_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByLangcodeAndFieldImageAltAndFieldImageHeight Get NodeRevision_fieldImages via LangcodeAndFieldImageAltAndFieldImageHeight
func GetNodeRevision_fieldImagesByLangcodeAndFieldImageAltAndFieldImageHeight(offset int, limit int, Langcode_ string, FieldImageAlt_ string, FieldImageHeight_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("langcode = ? and field_image_alt = ? and field_image_height = ?", Langcode_, FieldImageAlt_, FieldImageHeight_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByLangcodeAndFieldImageTitleAndFieldImageWidth Get NodeRevision_fieldImages via LangcodeAndFieldImageTitleAndFieldImageWidth
func GetNodeRevision_fieldImagesByLangcodeAndFieldImageTitleAndFieldImageWidth(offset int, limit int, Langcode_ string, FieldImageTitle_ string, FieldImageWidth_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("langcode = ? and field_image_title = ? and field_image_width = ?", Langcode_, FieldImageTitle_, FieldImageWidth_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByLangcodeAndFieldImageTitleAndFieldImageHeight Get NodeRevision_fieldImages via LangcodeAndFieldImageTitleAndFieldImageHeight
func GetNodeRevision_fieldImagesByLangcodeAndFieldImageTitleAndFieldImageHeight(offset int, limit int, Langcode_ string, FieldImageTitle_ string, FieldImageHeight_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("langcode = ? and field_image_title = ? and field_image_height = ?", Langcode_, FieldImageTitle_, FieldImageHeight_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByLangcodeAndFieldImageWidthAndFieldImageHeight Get NodeRevision_fieldImages via LangcodeAndFieldImageWidthAndFieldImageHeight
func GetNodeRevision_fieldImagesByLangcodeAndFieldImageWidthAndFieldImageHeight(offset int, limit int, Langcode_ string, FieldImageWidth_ int, FieldImageHeight_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("langcode = ? and field_image_width = ? and field_image_height = ?", Langcode_, FieldImageWidth_, FieldImageHeight_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByDeltaAndFieldImageTargetIdAndFieldImageAlt Get NodeRevision_fieldImages via DeltaAndFieldImageTargetIdAndFieldImageAlt
func GetNodeRevision_fieldImagesByDeltaAndFieldImageTargetIdAndFieldImageAlt(offset int, limit int, Delta_ int, FieldImageTargetId_ int, FieldImageAlt_ string) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("delta = ? and field_image_target_id = ? and field_image_alt = ?", Delta_, FieldImageTargetId_, FieldImageAlt_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByDeltaAndFieldImageTargetIdAndFieldImageTitle Get NodeRevision_fieldImages via DeltaAndFieldImageTargetIdAndFieldImageTitle
func GetNodeRevision_fieldImagesByDeltaAndFieldImageTargetIdAndFieldImageTitle(offset int, limit int, Delta_ int, FieldImageTargetId_ int, FieldImageTitle_ string) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("delta = ? and field_image_target_id = ? and field_image_title = ?", Delta_, FieldImageTargetId_, FieldImageTitle_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByDeltaAndFieldImageTargetIdAndFieldImageWidth Get NodeRevision_fieldImages via DeltaAndFieldImageTargetIdAndFieldImageWidth
func GetNodeRevision_fieldImagesByDeltaAndFieldImageTargetIdAndFieldImageWidth(offset int, limit int, Delta_ int, FieldImageTargetId_ int, FieldImageWidth_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("delta = ? and field_image_target_id = ? and field_image_width = ?", Delta_, FieldImageTargetId_, FieldImageWidth_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByDeltaAndFieldImageTargetIdAndFieldImageHeight Get NodeRevision_fieldImages via DeltaAndFieldImageTargetIdAndFieldImageHeight
func GetNodeRevision_fieldImagesByDeltaAndFieldImageTargetIdAndFieldImageHeight(offset int, limit int, Delta_ int, FieldImageTargetId_ int, FieldImageHeight_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("delta = ? and field_image_target_id = ? and field_image_height = ?", Delta_, FieldImageTargetId_, FieldImageHeight_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByDeltaAndFieldImageAltAndFieldImageTitle Get NodeRevision_fieldImages via DeltaAndFieldImageAltAndFieldImageTitle
func GetNodeRevision_fieldImagesByDeltaAndFieldImageAltAndFieldImageTitle(offset int, limit int, Delta_ int, FieldImageAlt_ string, FieldImageTitle_ string) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("delta = ? and field_image_alt = ? and field_image_title = ?", Delta_, FieldImageAlt_, FieldImageTitle_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByDeltaAndFieldImageAltAndFieldImageWidth Get NodeRevision_fieldImages via DeltaAndFieldImageAltAndFieldImageWidth
func GetNodeRevision_fieldImagesByDeltaAndFieldImageAltAndFieldImageWidth(offset int, limit int, Delta_ int, FieldImageAlt_ string, FieldImageWidth_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("delta = ? and field_image_alt = ? and field_image_width = ?", Delta_, FieldImageAlt_, FieldImageWidth_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByDeltaAndFieldImageAltAndFieldImageHeight Get NodeRevision_fieldImages via DeltaAndFieldImageAltAndFieldImageHeight
func GetNodeRevision_fieldImagesByDeltaAndFieldImageAltAndFieldImageHeight(offset int, limit int, Delta_ int, FieldImageAlt_ string, FieldImageHeight_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("delta = ? and field_image_alt = ? and field_image_height = ?", Delta_, FieldImageAlt_, FieldImageHeight_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByDeltaAndFieldImageTitleAndFieldImageWidth Get NodeRevision_fieldImages via DeltaAndFieldImageTitleAndFieldImageWidth
func GetNodeRevision_fieldImagesByDeltaAndFieldImageTitleAndFieldImageWidth(offset int, limit int, Delta_ int, FieldImageTitle_ string, FieldImageWidth_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("delta = ? and field_image_title = ? and field_image_width = ?", Delta_, FieldImageTitle_, FieldImageWidth_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByDeltaAndFieldImageTitleAndFieldImageHeight Get NodeRevision_fieldImages via DeltaAndFieldImageTitleAndFieldImageHeight
func GetNodeRevision_fieldImagesByDeltaAndFieldImageTitleAndFieldImageHeight(offset int, limit int, Delta_ int, FieldImageTitle_ string, FieldImageHeight_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("delta = ? and field_image_title = ? and field_image_height = ?", Delta_, FieldImageTitle_, FieldImageHeight_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByDeltaAndFieldImageWidthAndFieldImageHeight Get NodeRevision_fieldImages via DeltaAndFieldImageWidthAndFieldImageHeight
func GetNodeRevision_fieldImagesByDeltaAndFieldImageWidthAndFieldImageHeight(offset int, limit int, Delta_ int, FieldImageWidth_ int, FieldImageHeight_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("delta = ? and field_image_width = ? and field_image_height = ?", Delta_, FieldImageWidth_, FieldImageHeight_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByFieldImageTargetIdAndFieldImageAltAndFieldImageTitle Get NodeRevision_fieldImages via FieldImageTargetIdAndFieldImageAltAndFieldImageTitle
func GetNodeRevision_fieldImagesByFieldImageTargetIdAndFieldImageAltAndFieldImageTitle(offset int, limit int, FieldImageTargetId_ int, FieldImageAlt_ string, FieldImageTitle_ string) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("field_image_target_id = ? and field_image_alt = ? and field_image_title = ?", FieldImageTargetId_, FieldImageAlt_, FieldImageTitle_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByFieldImageTargetIdAndFieldImageAltAndFieldImageWidth Get NodeRevision_fieldImages via FieldImageTargetIdAndFieldImageAltAndFieldImageWidth
func GetNodeRevision_fieldImagesByFieldImageTargetIdAndFieldImageAltAndFieldImageWidth(offset int, limit int, FieldImageTargetId_ int, FieldImageAlt_ string, FieldImageWidth_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("field_image_target_id = ? and field_image_alt = ? and field_image_width = ?", FieldImageTargetId_, FieldImageAlt_, FieldImageWidth_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByFieldImageTargetIdAndFieldImageAltAndFieldImageHeight Get NodeRevision_fieldImages via FieldImageTargetIdAndFieldImageAltAndFieldImageHeight
func GetNodeRevision_fieldImagesByFieldImageTargetIdAndFieldImageAltAndFieldImageHeight(offset int, limit int, FieldImageTargetId_ int, FieldImageAlt_ string, FieldImageHeight_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("field_image_target_id = ? and field_image_alt = ? and field_image_height = ?", FieldImageTargetId_, FieldImageAlt_, FieldImageHeight_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByFieldImageTargetIdAndFieldImageTitleAndFieldImageWidth Get NodeRevision_fieldImages via FieldImageTargetIdAndFieldImageTitleAndFieldImageWidth
func GetNodeRevision_fieldImagesByFieldImageTargetIdAndFieldImageTitleAndFieldImageWidth(offset int, limit int, FieldImageTargetId_ int, FieldImageTitle_ string, FieldImageWidth_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("field_image_target_id = ? and field_image_title = ? and field_image_width = ?", FieldImageTargetId_, FieldImageTitle_, FieldImageWidth_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByFieldImageTargetIdAndFieldImageTitleAndFieldImageHeight Get NodeRevision_fieldImages via FieldImageTargetIdAndFieldImageTitleAndFieldImageHeight
func GetNodeRevision_fieldImagesByFieldImageTargetIdAndFieldImageTitleAndFieldImageHeight(offset int, limit int, FieldImageTargetId_ int, FieldImageTitle_ string, FieldImageHeight_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("field_image_target_id = ? and field_image_title = ? and field_image_height = ?", FieldImageTargetId_, FieldImageTitle_, FieldImageHeight_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByFieldImageTargetIdAndFieldImageWidthAndFieldImageHeight Get NodeRevision_fieldImages via FieldImageTargetIdAndFieldImageWidthAndFieldImageHeight
func GetNodeRevision_fieldImagesByFieldImageTargetIdAndFieldImageWidthAndFieldImageHeight(offset int, limit int, FieldImageTargetId_ int, FieldImageWidth_ int, FieldImageHeight_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("field_image_target_id = ? and field_image_width = ? and field_image_height = ?", FieldImageTargetId_, FieldImageWidth_, FieldImageHeight_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByFieldImageAltAndFieldImageTitleAndFieldImageWidth Get NodeRevision_fieldImages via FieldImageAltAndFieldImageTitleAndFieldImageWidth
func GetNodeRevision_fieldImagesByFieldImageAltAndFieldImageTitleAndFieldImageWidth(offset int, limit int, FieldImageAlt_ string, FieldImageTitle_ string, FieldImageWidth_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("field_image_alt = ? and field_image_title = ? and field_image_width = ?", FieldImageAlt_, FieldImageTitle_, FieldImageWidth_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByFieldImageAltAndFieldImageTitleAndFieldImageHeight Get NodeRevision_fieldImages via FieldImageAltAndFieldImageTitleAndFieldImageHeight
func GetNodeRevision_fieldImagesByFieldImageAltAndFieldImageTitleAndFieldImageHeight(offset int, limit int, FieldImageAlt_ string, FieldImageTitle_ string, FieldImageHeight_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("field_image_alt = ? and field_image_title = ? and field_image_height = ?", FieldImageAlt_, FieldImageTitle_, FieldImageHeight_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByFieldImageAltAndFieldImageWidthAndFieldImageHeight Get NodeRevision_fieldImages via FieldImageAltAndFieldImageWidthAndFieldImageHeight
func GetNodeRevision_fieldImagesByFieldImageAltAndFieldImageWidthAndFieldImageHeight(offset int, limit int, FieldImageAlt_ string, FieldImageWidth_ int, FieldImageHeight_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("field_image_alt = ? and field_image_width = ? and field_image_height = ?", FieldImageAlt_, FieldImageWidth_, FieldImageHeight_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByFieldImageTitleAndFieldImageWidthAndFieldImageHeight Get NodeRevision_fieldImages via FieldImageTitleAndFieldImageWidthAndFieldImageHeight
func GetNodeRevision_fieldImagesByFieldImageTitleAndFieldImageWidthAndFieldImageHeight(offset int, limit int, FieldImageTitle_ string, FieldImageWidth_ int, FieldImageHeight_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("field_image_title = ? and field_image_width = ? and field_image_height = ?", FieldImageTitle_, FieldImageWidth_, FieldImageHeight_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByBundleAndDeleted Get NodeRevision_fieldImages via BundleAndDeleted
func GetNodeRevision_fieldImagesByBundleAndDeleted(offset int, limit int, Bundle_ string, Deleted_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("bundle = ? and deleted = ?", Bundle_, Deleted_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByBundleAndEntityId Get NodeRevision_fieldImages via BundleAndEntityId
func GetNodeRevision_fieldImagesByBundleAndEntityId(offset int, limit int, Bundle_ string, EntityId_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("bundle = ? and entity_id = ?", Bundle_, EntityId_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByBundleAndRevisionId Get NodeRevision_fieldImages via BundleAndRevisionId
func GetNodeRevision_fieldImagesByBundleAndRevisionId(offset int, limit int, Bundle_ string, RevisionId_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("bundle = ? and revision_id = ?", Bundle_, RevisionId_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByBundleAndLangcode Get NodeRevision_fieldImages via BundleAndLangcode
func GetNodeRevision_fieldImagesByBundleAndLangcode(offset int, limit int, Bundle_ string, Langcode_ string) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("bundle = ? and langcode = ?", Bundle_, Langcode_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByBundleAndDelta Get NodeRevision_fieldImages via BundleAndDelta
func GetNodeRevision_fieldImagesByBundleAndDelta(offset int, limit int, Bundle_ string, Delta_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("bundle = ? and delta = ?", Bundle_, Delta_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByBundleAndFieldImageTargetId Get NodeRevision_fieldImages via BundleAndFieldImageTargetId
func GetNodeRevision_fieldImagesByBundleAndFieldImageTargetId(offset int, limit int, Bundle_ string, FieldImageTargetId_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("bundle = ? and field_image_target_id = ?", Bundle_, FieldImageTargetId_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByBundleAndFieldImageAlt Get NodeRevision_fieldImages via BundleAndFieldImageAlt
func GetNodeRevision_fieldImagesByBundleAndFieldImageAlt(offset int, limit int, Bundle_ string, FieldImageAlt_ string) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("bundle = ? and field_image_alt = ?", Bundle_, FieldImageAlt_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByBundleAndFieldImageTitle Get NodeRevision_fieldImages via BundleAndFieldImageTitle
func GetNodeRevision_fieldImagesByBundleAndFieldImageTitle(offset int, limit int, Bundle_ string, FieldImageTitle_ string) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("bundle = ? and field_image_title = ?", Bundle_, FieldImageTitle_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByBundleAndFieldImageWidth Get NodeRevision_fieldImages via BundleAndFieldImageWidth
func GetNodeRevision_fieldImagesByBundleAndFieldImageWidth(offset int, limit int, Bundle_ string, FieldImageWidth_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("bundle = ? and field_image_width = ?", Bundle_, FieldImageWidth_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByBundleAndFieldImageHeight Get NodeRevision_fieldImages via BundleAndFieldImageHeight
func GetNodeRevision_fieldImagesByBundleAndFieldImageHeight(offset int, limit int, Bundle_ string, FieldImageHeight_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("bundle = ? and field_image_height = ?", Bundle_, FieldImageHeight_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByDeletedAndEntityId Get NodeRevision_fieldImages via DeletedAndEntityId
func GetNodeRevision_fieldImagesByDeletedAndEntityId(offset int, limit int, Deleted_ int, EntityId_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("deleted = ? and entity_id = ?", Deleted_, EntityId_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByDeletedAndRevisionId Get NodeRevision_fieldImages via DeletedAndRevisionId
func GetNodeRevision_fieldImagesByDeletedAndRevisionId(offset int, limit int, Deleted_ int, RevisionId_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("deleted = ? and revision_id = ?", Deleted_, RevisionId_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByDeletedAndLangcode Get NodeRevision_fieldImages via DeletedAndLangcode
func GetNodeRevision_fieldImagesByDeletedAndLangcode(offset int, limit int, Deleted_ int, Langcode_ string) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("deleted = ? and langcode = ?", Deleted_, Langcode_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByDeletedAndDelta Get NodeRevision_fieldImages via DeletedAndDelta
func GetNodeRevision_fieldImagesByDeletedAndDelta(offset int, limit int, Deleted_ int, Delta_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("deleted = ? and delta = ?", Deleted_, Delta_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByDeletedAndFieldImageTargetId Get NodeRevision_fieldImages via DeletedAndFieldImageTargetId
func GetNodeRevision_fieldImagesByDeletedAndFieldImageTargetId(offset int, limit int, Deleted_ int, FieldImageTargetId_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("deleted = ? and field_image_target_id = ?", Deleted_, FieldImageTargetId_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByDeletedAndFieldImageAlt Get NodeRevision_fieldImages via DeletedAndFieldImageAlt
func GetNodeRevision_fieldImagesByDeletedAndFieldImageAlt(offset int, limit int, Deleted_ int, FieldImageAlt_ string) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("deleted = ? and field_image_alt = ?", Deleted_, FieldImageAlt_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByDeletedAndFieldImageTitle Get NodeRevision_fieldImages via DeletedAndFieldImageTitle
func GetNodeRevision_fieldImagesByDeletedAndFieldImageTitle(offset int, limit int, Deleted_ int, FieldImageTitle_ string) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("deleted = ? and field_image_title = ?", Deleted_, FieldImageTitle_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByDeletedAndFieldImageWidth Get NodeRevision_fieldImages via DeletedAndFieldImageWidth
func GetNodeRevision_fieldImagesByDeletedAndFieldImageWidth(offset int, limit int, Deleted_ int, FieldImageWidth_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("deleted = ? and field_image_width = ?", Deleted_, FieldImageWidth_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByDeletedAndFieldImageHeight Get NodeRevision_fieldImages via DeletedAndFieldImageHeight
func GetNodeRevision_fieldImagesByDeletedAndFieldImageHeight(offset int, limit int, Deleted_ int, FieldImageHeight_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("deleted = ? and field_image_height = ?", Deleted_, FieldImageHeight_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByEntityIdAndRevisionId Get NodeRevision_fieldImages via EntityIdAndRevisionId
func GetNodeRevision_fieldImagesByEntityIdAndRevisionId(offset int, limit int, EntityId_ int, RevisionId_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("entity_id = ? and revision_id = ?", EntityId_, RevisionId_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByEntityIdAndLangcode Get NodeRevision_fieldImages via EntityIdAndLangcode
func GetNodeRevision_fieldImagesByEntityIdAndLangcode(offset int, limit int, EntityId_ int, Langcode_ string) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("entity_id = ? and langcode = ?", EntityId_, Langcode_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByEntityIdAndDelta Get NodeRevision_fieldImages via EntityIdAndDelta
func GetNodeRevision_fieldImagesByEntityIdAndDelta(offset int, limit int, EntityId_ int, Delta_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("entity_id = ? and delta = ?", EntityId_, Delta_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByEntityIdAndFieldImageTargetId Get NodeRevision_fieldImages via EntityIdAndFieldImageTargetId
func GetNodeRevision_fieldImagesByEntityIdAndFieldImageTargetId(offset int, limit int, EntityId_ int, FieldImageTargetId_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("entity_id = ? and field_image_target_id = ?", EntityId_, FieldImageTargetId_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByEntityIdAndFieldImageAlt Get NodeRevision_fieldImages via EntityIdAndFieldImageAlt
func GetNodeRevision_fieldImagesByEntityIdAndFieldImageAlt(offset int, limit int, EntityId_ int, FieldImageAlt_ string) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("entity_id = ? and field_image_alt = ?", EntityId_, FieldImageAlt_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByEntityIdAndFieldImageTitle Get NodeRevision_fieldImages via EntityIdAndFieldImageTitle
func GetNodeRevision_fieldImagesByEntityIdAndFieldImageTitle(offset int, limit int, EntityId_ int, FieldImageTitle_ string) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("entity_id = ? and field_image_title = ?", EntityId_, FieldImageTitle_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByEntityIdAndFieldImageWidth Get NodeRevision_fieldImages via EntityIdAndFieldImageWidth
func GetNodeRevision_fieldImagesByEntityIdAndFieldImageWidth(offset int, limit int, EntityId_ int, FieldImageWidth_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("entity_id = ? and field_image_width = ?", EntityId_, FieldImageWidth_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByEntityIdAndFieldImageHeight Get NodeRevision_fieldImages via EntityIdAndFieldImageHeight
func GetNodeRevision_fieldImagesByEntityIdAndFieldImageHeight(offset int, limit int, EntityId_ int, FieldImageHeight_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("entity_id = ? and field_image_height = ?", EntityId_, FieldImageHeight_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByRevisionIdAndLangcode Get NodeRevision_fieldImages via RevisionIdAndLangcode
func GetNodeRevision_fieldImagesByRevisionIdAndLangcode(offset int, limit int, RevisionId_ int, Langcode_ string) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("revision_id = ? and langcode = ?", RevisionId_, Langcode_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByRevisionIdAndDelta Get NodeRevision_fieldImages via RevisionIdAndDelta
func GetNodeRevision_fieldImagesByRevisionIdAndDelta(offset int, limit int, RevisionId_ int, Delta_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("revision_id = ? and delta = ?", RevisionId_, Delta_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByRevisionIdAndFieldImageTargetId Get NodeRevision_fieldImages via RevisionIdAndFieldImageTargetId
func GetNodeRevision_fieldImagesByRevisionIdAndFieldImageTargetId(offset int, limit int, RevisionId_ int, FieldImageTargetId_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("revision_id = ? and field_image_target_id = ?", RevisionId_, FieldImageTargetId_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByRevisionIdAndFieldImageAlt Get NodeRevision_fieldImages via RevisionIdAndFieldImageAlt
func GetNodeRevision_fieldImagesByRevisionIdAndFieldImageAlt(offset int, limit int, RevisionId_ int, FieldImageAlt_ string) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("revision_id = ? and field_image_alt = ?", RevisionId_, FieldImageAlt_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByRevisionIdAndFieldImageTitle Get NodeRevision_fieldImages via RevisionIdAndFieldImageTitle
func GetNodeRevision_fieldImagesByRevisionIdAndFieldImageTitle(offset int, limit int, RevisionId_ int, FieldImageTitle_ string) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("revision_id = ? and field_image_title = ?", RevisionId_, FieldImageTitle_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByRevisionIdAndFieldImageWidth Get NodeRevision_fieldImages via RevisionIdAndFieldImageWidth
func GetNodeRevision_fieldImagesByRevisionIdAndFieldImageWidth(offset int, limit int, RevisionId_ int, FieldImageWidth_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("revision_id = ? and field_image_width = ?", RevisionId_, FieldImageWidth_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByRevisionIdAndFieldImageHeight Get NodeRevision_fieldImages via RevisionIdAndFieldImageHeight
func GetNodeRevision_fieldImagesByRevisionIdAndFieldImageHeight(offset int, limit int, RevisionId_ int, FieldImageHeight_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("revision_id = ? and field_image_height = ?", RevisionId_, FieldImageHeight_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByLangcodeAndDelta Get NodeRevision_fieldImages via LangcodeAndDelta
func GetNodeRevision_fieldImagesByLangcodeAndDelta(offset int, limit int, Langcode_ string, Delta_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("langcode = ? and delta = ?", Langcode_, Delta_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByLangcodeAndFieldImageTargetId Get NodeRevision_fieldImages via LangcodeAndFieldImageTargetId
func GetNodeRevision_fieldImagesByLangcodeAndFieldImageTargetId(offset int, limit int, Langcode_ string, FieldImageTargetId_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("langcode = ? and field_image_target_id = ?", Langcode_, FieldImageTargetId_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByLangcodeAndFieldImageAlt Get NodeRevision_fieldImages via LangcodeAndFieldImageAlt
func GetNodeRevision_fieldImagesByLangcodeAndFieldImageAlt(offset int, limit int, Langcode_ string, FieldImageAlt_ string) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("langcode = ? and field_image_alt = ?", Langcode_, FieldImageAlt_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByLangcodeAndFieldImageTitle Get NodeRevision_fieldImages via LangcodeAndFieldImageTitle
func GetNodeRevision_fieldImagesByLangcodeAndFieldImageTitle(offset int, limit int, Langcode_ string, FieldImageTitle_ string) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("langcode = ? and field_image_title = ?", Langcode_, FieldImageTitle_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByLangcodeAndFieldImageWidth Get NodeRevision_fieldImages via LangcodeAndFieldImageWidth
func GetNodeRevision_fieldImagesByLangcodeAndFieldImageWidth(offset int, limit int, Langcode_ string, FieldImageWidth_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("langcode = ? and field_image_width = ?", Langcode_, FieldImageWidth_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByLangcodeAndFieldImageHeight Get NodeRevision_fieldImages via LangcodeAndFieldImageHeight
func GetNodeRevision_fieldImagesByLangcodeAndFieldImageHeight(offset int, limit int, Langcode_ string, FieldImageHeight_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("langcode = ? and field_image_height = ?", Langcode_, FieldImageHeight_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByDeltaAndFieldImageTargetId Get NodeRevision_fieldImages via DeltaAndFieldImageTargetId
func GetNodeRevision_fieldImagesByDeltaAndFieldImageTargetId(offset int, limit int, Delta_ int, FieldImageTargetId_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("delta = ? and field_image_target_id = ?", Delta_, FieldImageTargetId_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByDeltaAndFieldImageAlt Get NodeRevision_fieldImages via DeltaAndFieldImageAlt
func GetNodeRevision_fieldImagesByDeltaAndFieldImageAlt(offset int, limit int, Delta_ int, FieldImageAlt_ string) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("delta = ? and field_image_alt = ?", Delta_, FieldImageAlt_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByDeltaAndFieldImageTitle Get NodeRevision_fieldImages via DeltaAndFieldImageTitle
func GetNodeRevision_fieldImagesByDeltaAndFieldImageTitle(offset int, limit int, Delta_ int, FieldImageTitle_ string) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("delta = ? and field_image_title = ?", Delta_, FieldImageTitle_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByDeltaAndFieldImageWidth Get NodeRevision_fieldImages via DeltaAndFieldImageWidth
func GetNodeRevision_fieldImagesByDeltaAndFieldImageWidth(offset int, limit int, Delta_ int, FieldImageWidth_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("delta = ? and field_image_width = ?", Delta_, FieldImageWidth_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByDeltaAndFieldImageHeight Get NodeRevision_fieldImages via DeltaAndFieldImageHeight
func GetNodeRevision_fieldImagesByDeltaAndFieldImageHeight(offset int, limit int, Delta_ int, FieldImageHeight_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("delta = ? and field_image_height = ?", Delta_, FieldImageHeight_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByFieldImageTargetIdAndFieldImageAlt Get NodeRevision_fieldImages via FieldImageTargetIdAndFieldImageAlt
func GetNodeRevision_fieldImagesByFieldImageTargetIdAndFieldImageAlt(offset int, limit int, FieldImageTargetId_ int, FieldImageAlt_ string) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("field_image_target_id = ? and field_image_alt = ?", FieldImageTargetId_, FieldImageAlt_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByFieldImageTargetIdAndFieldImageTitle Get NodeRevision_fieldImages via FieldImageTargetIdAndFieldImageTitle
func GetNodeRevision_fieldImagesByFieldImageTargetIdAndFieldImageTitle(offset int, limit int, FieldImageTargetId_ int, FieldImageTitle_ string) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("field_image_target_id = ? and field_image_title = ?", FieldImageTargetId_, FieldImageTitle_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByFieldImageTargetIdAndFieldImageWidth Get NodeRevision_fieldImages via FieldImageTargetIdAndFieldImageWidth
func GetNodeRevision_fieldImagesByFieldImageTargetIdAndFieldImageWidth(offset int, limit int, FieldImageTargetId_ int, FieldImageWidth_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("field_image_target_id = ? and field_image_width = ?", FieldImageTargetId_, FieldImageWidth_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByFieldImageTargetIdAndFieldImageHeight Get NodeRevision_fieldImages via FieldImageTargetIdAndFieldImageHeight
func GetNodeRevision_fieldImagesByFieldImageTargetIdAndFieldImageHeight(offset int, limit int, FieldImageTargetId_ int, FieldImageHeight_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("field_image_target_id = ? and field_image_height = ?", FieldImageTargetId_, FieldImageHeight_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByFieldImageAltAndFieldImageTitle Get NodeRevision_fieldImages via FieldImageAltAndFieldImageTitle
func GetNodeRevision_fieldImagesByFieldImageAltAndFieldImageTitle(offset int, limit int, FieldImageAlt_ string, FieldImageTitle_ string) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("field_image_alt = ? and field_image_title = ?", FieldImageAlt_, FieldImageTitle_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByFieldImageAltAndFieldImageWidth Get NodeRevision_fieldImages via FieldImageAltAndFieldImageWidth
func GetNodeRevision_fieldImagesByFieldImageAltAndFieldImageWidth(offset int, limit int, FieldImageAlt_ string, FieldImageWidth_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("field_image_alt = ? and field_image_width = ?", FieldImageAlt_, FieldImageWidth_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByFieldImageAltAndFieldImageHeight Get NodeRevision_fieldImages via FieldImageAltAndFieldImageHeight
func GetNodeRevision_fieldImagesByFieldImageAltAndFieldImageHeight(offset int, limit int, FieldImageAlt_ string, FieldImageHeight_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("field_image_alt = ? and field_image_height = ?", FieldImageAlt_, FieldImageHeight_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByFieldImageTitleAndFieldImageWidth Get NodeRevision_fieldImages via FieldImageTitleAndFieldImageWidth
func GetNodeRevision_fieldImagesByFieldImageTitleAndFieldImageWidth(offset int, limit int, FieldImageTitle_ string, FieldImageWidth_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("field_image_title = ? and field_image_width = ?", FieldImageTitle_, FieldImageWidth_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByFieldImageTitleAndFieldImageHeight Get NodeRevision_fieldImages via FieldImageTitleAndFieldImageHeight
func GetNodeRevision_fieldImagesByFieldImageTitleAndFieldImageHeight(offset int, limit int, FieldImageTitle_ string, FieldImageHeight_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("field_image_title = ? and field_image_height = ?", FieldImageTitle_, FieldImageHeight_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesByFieldImageWidthAndFieldImageHeight Get NodeRevision_fieldImages via FieldImageWidthAndFieldImageHeight
func GetNodeRevision_fieldImagesByFieldImageWidthAndFieldImageHeight(offset int, limit int, FieldImageWidth_ int, FieldImageHeight_ int) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("field_image_width = ? and field_image_height = ?", FieldImageWidth_, FieldImageHeight_).Limit(limit, offset).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImages Get NodeRevision_fieldImages via field
func GetNodeRevision_fieldImages(offset int, limit int, field string) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Limit(limit, offset).Desc(field).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesViaBundle Get NodeRevision_fieldImages via Bundle
func GetNodeRevision_fieldImagesViaBundle(offset int, limit int, Bundle_ string, field string) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("bundle = ?", Bundle_).Limit(limit, offset).Desc(field).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesViaDeleted Get NodeRevision_fieldImages via Deleted
func GetNodeRevision_fieldImagesViaDeleted(offset int, limit int, Deleted_ int, field string) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("deleted = ?", Deleted_).Limit(limit, offset).Desc(field).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesViaEntityId Get NodeRevision_fieldImages via EntityId
func GetNodeRevision_fieldImagesViaEntityId(offset int, limit int, EntityId_ int, field string) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("entity_id = ?", EntityId_).Limit(limit, offset).Desc(field).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesViaRevisionId Get NodeRevision_fieldImages via RevisionId
func GetNodeRevision_fieldImagesViaRevisionId(offset int, limit int, RevisionId_ int, field string) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("revision_id = ?", RevisionId_).Limit(limit, offset).Desc(field).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesViaLangcode Get NodeRevision_fieldImages via Langcode
func GetNodeRevision_fieldImagesViaLangcode(offset int, limit int, Langcode_ string, field string) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("langcode = ?", Langcode_).Limit(limit, offset).Desc(field).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesViaDelta Get NodeRevision_fieldImages via Delta
func GetNodeRevision_fieldImagesViaDelta(offset int, limit int, Delta_ int, field string) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("delta = ?", Delta_).Limit(limit, offset).Desc(field).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesViaFieldImageTargetId Get NodeRevision_fieldImages via FieldImageTargetId
func GetNodeRevision_fieldImagesViaFieldImageTargetId(offset int, limit int, FieldImageTargetId_ int, field string) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("field_image_target_id = ?", FieldImageTargetId_).Limit(limit, offset).Desc(field).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesViaFieldImageAlt Get NodeRevision_fieldImages via FieldImageAlt
func GetNodeRevision_fieldImagesViaFieldImageAlt(offset int, limit int, FieldImageAlt_ string, field string) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("field_image_alt = ?", FieldImageAlt_).Limit(limit, offset).Desc(field).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesViaFieldImageTitle Get NodeRevision_fieldImages via FieldImageTitle
func GetNodeRevision_fieldImagesViaFieldImageTitle(offset int, limit int, FieldImageTitle_ string, field string) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("field_image_title = ?", FieldImageTitle_).Limit(limit, offset).Desc(field).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesViaFieldImageWidth Get NodeRevision_fieldImages via FieldImageWidth
func GetNodeRevision_fieldImagesViaFieldImageWidth(offset int, limit int, FieldImageWidth_ int, field string) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("field_image_width = ?", FieldImageWidth_).Limit(limit, offset).Desc(field).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// GetNodeRevision_fieldImagesViaFieldImageHeight Get NodeRevision_fieldImages via FieldImageHeight
func GetNodeRevision_fieldImagesViaFieldImageHeight(offset int, limit int, FieldImageHeight_ int, field string) (*[]*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = new([]*NodeRevision_fieldImage)
	err := Engine.Table("node_revision__field_image").Where("field_image_height = ?", FieldImageHeight_).Limit(limit, offset).Desc(field).Find(_NodeRevision_fieldImage)
	return _NodeRevision_fieldImage, err
}

// HasNodeRevision_fieldImageViaBundle Has NodeRevision_fieldImage via Bundle
func HasNodeRevision_fieldImageViaBundle(iBundle string) bool {
	if has, err := Engine.Where("bundle = ?", iBundle).Get(new(NodeRevision_fieldImage)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasNodeRevision_fieldImageViaDeleted Has NodeRevision_fieldImage via Deleted
func HasNodeRevision_fieldImageViaDeleted(iDeleted int) bool {
	if has, err := Engine.Where("deleted = ?", iDeleted).Get(new(NodeRevision_fieldImage)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasNodeRevision_fieldImageViaEntityId Has NodeRevision_fieldImage via EntityId
func HasNodeRevision_fieldImageViaEntityId(iEntityId int) bool {
	if has, err := Engine.Where("entity_id = ?", iEntityId).Get(new(NodeRevision_fieldImage)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasNodeRevision_fieldImageViaRevisionId Has NodeRevision_fieldImage via RevisionId
func HasNodeRevision_fieldImageViaRevisionId(iRevisionId int) bool {
	if has, err := Engine.Where("revision_id = ?", iRevisionId).Get(new(NodeRevision_fieldImage)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasNodeRevision_fieldImageViaLangcode Has NodeRevision_fieldImage via Langcode
func HasNodeRevision_fieldImageViaLangcode(iLangcode string) bool {
	if has, err := Engine.Where("langcode = ?", iLangcode).Get(new(NodeRevision_fieldImage)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasNodeRevision_fieldImageViaDelta Has NodeRevision_fieldImage via Delta
func HasNodeRevision_fieldImageViaDelta(iDelta int) bool {
	if has, err := Engine.Where("delta = ?", iDelta).Get(new(NodeRevision_fieldImage)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasNodeRevision_fieldImageViaFieldImageTargetId Has NodeRevision_fieldImage via FieldImageTargetId
func HasNodeRevision_fieldImageViaFieldImageTargetId(iFieldImageTargetId int) bool {
	if has, err := Engine.Where("field_image_target_id = ?", iFieldImageTargetId).Get(new(NodeRevision_fieldImage)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasNodeRevision_fieldImageViaFieldImageAlt Has NodeRevision_fieldImage via FieldImageAlt
func HasNodeRevision_fieldImageViaFieldImageAlt(iFieldImageAlt string) bool {
	if has, err := Engine.Where("field_image_alt = ?", iFieldImageAlt).Get(new(NodeRevision_fieldImage)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasNodeRevision_fieldImageViaFieldImageTitle Has NodeRevision_fieldImage via FieldImageTitle
func HasNodeRevision_fieldImageViaFieldImageTitle(iFieldImageTitle string) bool {
	if has, err := Engine.Where("field_image_title = ?", iFieldImageTitle).Get(new(NodeRevision_fieldImage)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasNodeRevision_fieldImageViaFieldImageWidth Has NodeRevision_fieldImage via FieldImageWidth
func HasNodeRevision_fieldImageViaFieldImageWidth(iFieldImageWidth int) bool {
	if has, err := Engine.Where("field_image_width = ?", iFieldImageWidth).Get(new(NodeRevision_fieldImage)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasNodeRevision_fieldImageViaFieldImageHeight Has NodeRevision_fieldImage via FieldImageHeight
func HasNodeRevision_fieldImageViaFieldImageHeight(iFieldImageHeight int) bool {
	if has, err := Engine.Where("field_image_height = ?", iFieldImageHeight).Get(new(NodeRevision_fieldImage)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// GetNodeRevision_fieldImageViaBundle Get NodeRevision_fieldImage via Bundle
func GetNodeRevision_fieldImageViaBundle(iBundle string) (*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = &NodeRevision_fieldImage{Bundle: iBundle}
	has, err := Engine.Get(_NodeRevision_fieldImage)
	if has {
		return _NodeRevision_fieldImage, err
	} else {
		return nil, err
	}
}

// GetNodeRevision_fieldImageViaDeleted Get NodeRevision_fieldImage via Deleted
func GetNodeRevision_fieldImageViaDeleted(iDeleted int) (*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = &NodeRevision_fieldImage{Deleted: iDeleted}
	has, err := Engine.Get(_NodeRevision_fieldImage)
	if has {
		return _NodeRevision_fieldImage, err
	} else {
		return nil, err
	}
}

// GetNodeRevision_fieldImageViaEntityId Get NodeRevision_fieldImage via EntityId
func GetNodeRevision_fieldImageViaEntityId(iEntityId int) (*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = &NodeRevision_fieldImage{EntityId: iEntityId}
	has, err := Engine.Get(_NodeRevision_fieldImage)
	if has {
		return _NodeRevision_fieldImage, err
	} else {
		return nil, err
	}
}

// GetNodeRevision_fieldImageViaRevisionId Get NodeRevision_fieldImage via RevisionId
func GetNodeRevision_fieldImageViaRevisionId(iRevisionId int) (*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = &NodeRevision_fieldImage{RevisionId: iRevisionId}
	has, err := Engine.Get(_NodeRevision_fieldImage)
	if has {
		return _NodeRevision_fieldImage, err
	} else {
		return nil, err
	}
}

// GetNodeRevision_fieldImageViaLangcode Get NodeRevision_fieldImage via Langcode
func GetNodeRevision_fieldImageViaLangcode(iLangcode string) (*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = &NodeRevision_fieldImage{Langcode: iLangcode}
	has, err := Engine.Get(_NodeRevision_fieldImage)
	if has {
		return _NodeRevision_fieldImage, err
	} else {
		return nil, err
	}
}

// GetNodeRevision_fieldImageViaDelta Get NodeRevision_fieldImage via Delta
func GetNodeRevision_fieldImageViaDelta(iDelta int) (*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = &NodeRevision_fieldImage{Delta: iDelta}
	has, err := Engine.Get(_NodeRevision_fieldImage)
	if has {
		return _NodeRevision_fieldImage, err
	} else {
		return nil, err
	}
}

// GetNodeRevision_fieldImageViaFieldImageTargetId Get NodeRevision_fieldImage via FieldImageTargetId
func GetNodeRevision_fieldImageViaFieldImageTargetId(iFieldImageTargetId int) (*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = &NodeRevision_fieldImage{FieldImageTargetId: iFieldImageTargetId}
	has, err := Engine.Get(_NodeRevision_fieldImage)
	if has {
		return _NodeRevision_fieldImage, err
	} else {
		return nil, err
	}
}

// GetNodeRevision_fieldImageViaFieldImageAlt Get NodeRevision_fieldImage via FieldImageAlt
func GetNodeRevision_fieldImageViaFieldImageAlt(iFieldImageAlt string) (*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = &NodeRevision_fieldImage{FieldImageAlt: iFieldImageAlt}
	has, err := Engine.Get(_NodeRevision_fieldImage)
	if has {
		return _NodeRevision_fieldImage, err
	} else {
		return nil, err
	}
}

// GetNodeRevision_fieldImageViaFieldImageTitle Get NodeRevision_fieldImage via FieldImageTitle
func GetNodeRevision_fieldImageViaFieldImageTitle(iFieldImageTitle string) (*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = &NodeRevision_fieldImage{FieldImageTitle: iFieldImageTitle}
	has, err := Engine.Get(_NodeRevision_fieldImage)
	if has {
		return _NodeRevision_fieldImage, err
	} else {
		return nil, err
	}
}

// GetNodeRevision_fieldImageViaFieldImageWidth Get NodeRevision_fieldImage via FieldImageWidth
func GetNodeRevision_fieldImageViaFieldImageWidth(iFieldImageWidth int) (*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = &NodeRevision_fieldImage{FieldImageWidth: iFieldImageWidth}
	has, err := Engine.Get(_NodeRevision_fieldImage)
	if has {
		return _NodeRevision_fieldImage, err
	} else {
		return nil, err
	}
}

// GetNodeRevision_fieldImageViaFieldImageHeight Get NodeRevision_fieldImage via FieldImageHeight
func GetNodeRevision_fieldImageViaFieldImageHeight(iFieldImageHeight int) (*NodeRevision_fieldImage, error) {
	var _NodeRevision_fieldImage = &NodeRevision_fieldImage{FieldImageHeight: iFieldImageHeight}
	has, err := Engine.Get(_NodeRevision_fieldImage)
	if has {
		return _NodeRevision_fieldImage, err
	} else {
		return nil, err
	}
}

// SetNodeRevision_fieldImageViaBundle Set NodeRevision_fieldImage via Bundle
func SetNodeRevision_fieldImageViaBundle(iBundle string, node_revision__field_image *NodeRevision_fieldImage) (int64, error) {
	node_revision__field_image.Bundle = iBundle
	return Engine.Insert(node_revision__field_image)
}

// SetNodeRevision_fieldImageViaDeleted Set NodeRevision_fieldImage via Deleted
func SetNodeRevision_fieldImageViaDeleted(iDeleted int, node_revision__field_image *NodeRevision_fieldImage) (int64, error) {
	node_revision__field_image.Deleted = iDeleted
	return Engine.Insert(node_revision__field_image)
}

// SetNodeRevision_fieldImageViaEntityId Set NodeRevision_fieldImage via EntityId
func SetNodeRevision_fieldImageViaEntityId(iEntityId int, node_revision__field_image *NodeRevision_fieldImage) (int64, error) {
	node_revision__field_image.EntityId = iEntityId
	return Engine.Insert(node_revision__field_image)
}

// SetNodeRevision_fieldImageViaRevisionId Set NodeRevision_fieldImage via RevisionId
func SetNodeRevision_fieldImageViaRevisionId(iRevisionId int, node_revision__field_image *NodeRevision_fieldImage) (int64, error) {
	node_revision__field_image.RevisionId = iRevisionId
	return Engine.Insert(node_revision__field_image)
}

// SetNodeRevision_fieldImageViaLangcode Set NodeRevision_fieldImage via Langcode
func SetNodeRevision_fieldImageViaLangcode(iLangcode string, node_revision__field_image *NodeRevision_fieldImage) (int64, error) {
	node_revision__field_image.Langcode = iLangcode
	return Engine.Insert(node_revision__field_image)
}

// SetNodeRevision_fieldImageViaDelta Set NodeRevision_fieldImage via Delta
func SetNodeRevision_fieldImageViaDelta(iDelta int, node_revision__field_image *NodeRevision_fieldImage) (int64, error) {
	node_revision__field_image.Delta = iDelta
	return Engine.Insert(node_revision__field_image)
}

// SetNodeRevision_fieldImageViaFieldImageTargetId Set NodeRevision_fieldImage via FieldImageTargetId
func SetNodeRevision_fieldImageViaFieldImageTargetId(iFieldImageTargetId int, node_revision__field_image *NodeRevision_fieldImage) (int64, error) {
	node_revision__field_image.FieldImageTargetId = iFieldImageTargetId
	return Engine.Insert(node_revision__field_image)
}

// SetNodeRevision_fieldImageViaFieldImageAlt Set NodeRevision_fieldImage via FieldImageAlt
func SetNodeRevision_fieldImageViaFieldImageAlt(iFieldImageAlt string, node_revision__field_image *NodeRevision_fieldImage) (int64, error) {
	node_revision__field_image.FieldImageAlt = iFieldImageAlt
	return Engine.Insert(node_revision__field_image)
}

// SetNodeRevision_fieldImageViaFieldImageTitle Set NodeRevision_fieldImage via FieldImageTitle
func SetNodeRevision_fieldImageViaFieldImageTitle(iFieldImageTitle string, node_revision__field_image *NodeRevision_fieldImage) (int64, error) {
	node_revision__field_image.FieldImageTitle = iFieldImageTitle
	return Engine.Insert(node_revision__field_image)
}

// SetNodeRevision_fieldImageViaFieldImageWidth Set NodeRevision_fieldImage via FieldImageWidth
func SetNodeRevision_fieldImageViaFieldImageWidth(iFieldImageWidth int, node_revision__field_image *NodeRevision_fieldImage) (int64, error) {
	node_revision__field_image.FieldImageWidth = iFieldImageWidth
	return Engine.Insert(node_revision__field_image)
}

// SetNodeRevision_fieldImageViaFieldImageHeight Set NodeRevision_fieldImage via FieldImageHeight
func SetNodeRevision_fieldImageViaFieldImageHeight(iFieldImageHeight int, node_revision__field_image *NodeRevision_fieldImage) (int64, error) {
	node_revision__field_image.FieldImageHeight = iFieldImageHeight
	return Engine.Insert(node_revision__field_image)
}

// AddNodeRevision_fieldImage Add NodeRevision_fieldImage via all columns
func AddNodeRevision_fieldImage(iBundle string, iDeleted int, iEntityId int, iRevisionId int, iLangcode string, iDelta int, iFieldImageTargetId int, iFieldImageAlt string, iFieldImageTitle string, iFieldImageWidth int, iFieldImageHeight int) error {
	_NodeRevision_fieldImage := &NodeRevision_fieldImage{Bundle: iBundle, Deleted: iDeleted, EntityId: iEntityId, RevisionId: iRevisionId, Langcode: iLangcode, Delta: iDelta, FieldImageTargetId: iFieldImageTargetId, FieldImageAlt: iFieldImageAlt, FieldImageTitle: iFieldImageTitle, FieldImageWidth: iFieldImageWidth, FieldImageHeight: iFieldImageHeight}
	if _, err := Engine.Insert(_NodeRevision_fieldImage); err != nil {
		return err
	} else {
		return nil
	}
}

// PostNodeRevision_fieldImage Post NodeRevision_fieldImage via iNodeRevision_fieldImage
func PostNodeRevision_fieldImage(iNodeRevision_fieldImage *NodeRevision_fieldImage) (string, error) {
	_, err := Engine.Insert(iNodeRevision_fieldImage)
	return iNodeRevision_fieldImage.Bundle, err
}

// PutNodeRevision_fieldImage Put NodeRevision_fieldImage
func PutNodeRevision_fieldImage(iNodeRevision_fieldImage *NodeRevision_fieldImage) (string, error) {
	_, err := Engine.Id(iNodeRevision_fieldImage.Bundle).Update(iNodeRevision_fieldImage)
	return iNodeRevision_fieldImage.Bundle, err
}

// PutNodeRevision_fieldImageViaBundle Put NodeRevision_fieldImage via Bundle
func PutNodeRevision_fieldImageViaBundle(Bundle_ string, iNodeRevision_fieldImage *NodeRevision_fieldImage) (int64, error) {
	row, err := Engine.Update(iNodeRevision_fieldImage, &NodeRevision_fieldImage{Bundle: Bundle_})
	return row, err
}

// PutNodeRevision_fieldImageViaDeleted Put NodeRevision_fieldImage via Deleted
func PutNodeRevision_fieldImageViaDeleted(Deleted_ int, iNodeRevision_fieldImage *NodeRevision_fieldImage) (int64, error) {
	row, err := Engine.Update(iNodeRevision_fieldImage, &NodeRevision_fieldImage{Deleted: Deleted_})
	return row, err
}

// PutNodeRevision_fieldImageViaEntityId Put NodeRevision_fieldImage via EntityId
func PutNodeRevision_fieldImageViaEntityId(EntityId_ int, iNodeRevision_fieldImage *NodeRevision_fieldImage) (int64, error) {
	row, err := Engine.Update(iNodeRevision_fieldImage, &NodeRevision_fieldImage{EntityId: EntityId_})
	return row, err
}

// PutNodeRevision_fieldImageViaRevisionId Put NodeRevision_fieldImage via RevisionId
func PutNodeRevision_fieldImageViaRevisionId(RevisionId_ int, iNodeRevision_fieldImage *NodeRevision_fieldImage) (int64, error) {
	row, err := Engine.Update(iNodeRevision_fieldImage, &NodeRevision_fieldImage{RevisionId: RevisionId_})
	return row, err
}

// PutNodeRevision_fieldImageViaLangcode Put NodeRevision_fieldImage via Langcode
func PutNodeRevision_fieldImageViaLangcode(Langcode_ string, iNodeRevision_fieldImage *NodeRevision_fieldImage) (int64, error) {
	row, err := Engine.Update(iNodeRevision_fieldImage, &NodeRevision_fieldImage{Langcode: Langcode_})
	return row, err
}

// PutNodeRevision_fieldImageViaDelta Put NodeRevision_fieldImage via Delta
func PutNodeRevision_fieldImageViaDelta(Delta_ int, iNodeRevision_fieldImage *NodeRevision_fieldImage) (int64, error) {
	row, err := Engine.Update(iNodeRevision_fieldImage, &NodeRevision_fieldImage{Delta: Delta_})
	return row, err
}

// PutNodeRevision_fieldImageViaFieldImageTargetId Put NodeRevision_fieldImage via FieldImageTargetId
func PutNodeRevision_fieldImageViaFieldImageTargetId(FieldImageTargetId_ int, iNodeRevision_fieldImage *NodeRevision_fieldImage) (int64, error) {
	row, err := Engine.Update(iNodeRevision_fieldImage, &NodeRevision_fieldImage{FieldImageTargetId: FieldImageTargetId_})
	return row, err
}

// PutNodeRevision_fieldImageViaFieldImageAlt Put NodeRevision_fieldImage via FieldImageAlt
func PutNodeRevision_fieldImageViaFieldImageAlt(FieldImageAlt_ string, iNodeRevision_fieldImage *NodeRevision_fieldImage) (int64, error) {
	row, err := Engine.Update(iNodeRevision_fieldImage, &NodeRevision_fieldImage{FieldImageAlt: FieldImageAlt_})
	return row, err
}

// PutNodeRevision_fieldImageViaFieldImageTitle Put NodeRevision_fieldImage via FieldImageTitle
func PutNodeRevision_fieldImageViaFieldImageTitle(FieldImageTitle_ string, iNodeRevision_fieldImage *NodeRevision_fieldImage) (int64, error) {
	row, err := Engine.Update(iNodeRevision_fieldImage, &NodeRevision_fieldImage{FieldImageTitle: FieldImageTitle_})
	return row, err
}

// PutNodeRevision_fieldImageViaFieldImageWidth Put NodeRevision_fieldImage via FieldImageWidth
func PutNodeRevision_fieldImageViaFieldImageWidth(FieldImageWidth_ int, iNodeRevision_fieldImage *NodeRevision_fieldImage) (int64, error) {
	row, err := Engine.Update(iNodeRevision_fieldImage, &NodeRevision_fieldImage{FieldImageWidth: FieldImageWidth_})
	return row, err
}

// PutNodeRevision_fieldImageViaFieldImageHeight Put NodeRevision_fieldImage via FieldImageHeight
func PutNodeRevision_fieldImageViaFieldImageHeight(FieldImageHeight_ int, iNodeRevision_fieldImage *NodeRevision_fieldImage) (int64, error) {
	row, err := Engine.Update(iNodeRevision_fieldImage, &NodeRevision_fieldImage{FieldImageHeight: FieldImageHeight_})
	return row, err
}

// UpdateNodeRevision_fieldImageViaBundle via map[string]interface{}{}
func UpdateNodeRevision_fieldImageViaBundle(iBundle string, iNodeRevision_fieldImageMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(NodeRevision_fieldImage)).Where("bundle = ?", iBundle).Update(iNodeRevision_fieldImageMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateNodeRevision_fieldImageViaDeleted via map[string]interface{}{}
func UpdateNodeRevision_fieldImageViaDeleted(iDeleted int, iNodeRevision_fieldImageMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(NodeRevision_fieldImage)).Where("deleted = ?", iDeleted).Update(iNodeRevision_fieldImageMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateNodeRevision_fieldImageViaEntityId via map[string]interface{}{}
func UpdateNodeRevision_fieldImageViaEntityId(iEntityId int, iNodeRevision_fieldImageMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(NodeRevision_fieldImage)).Where("entity_id = ?", iEntityId).Update(iNodeRevision_fieldImageMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateNodeRevision_fieldImageViaRevisionId via map[string]interface{}{}
func UpdateNodeRevision_fieldImageViaRevisionId(iRevisionId int, iNodeRevision_fieldImageMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(NodeRevision_fieldImage)).Where("revision_id = ?", iRevisionId).Update(iNodeRevision_fieldImageMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateNodeRevision_fieldImageViaLangcode via map[string]interface{}{}
func UpdateNodeRevision_fieldImageViaLangcode(iLangcode string, iNodeRevision_fieldImageMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(NodeRevision_fieldImage)).Where("langcode = ?", iLangcode).Update(iNodeRevision_fieldImageMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateNodeRevision_fieldImageViaDelta via map[string]interface{}{}
func UpdateNodeRevision_fieldImageViaDelta(iDelta int, iNodeRevision_fieldImageMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(NodeRevision_fieldImage)).Where("delta = ?", iDelta).Update(iNodeRevision_fieldImageMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateNodeRevision_fieldImageViaFieldImageTargetId via map[string]interface{}{}
func UpdateNodeRevision_fieldImageViaFieldImageTargetId(iFieldImageTargetId int, iNodeRevision_fieldImageMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(NodeRevision_fieldImage)).Where("field_image_target_id = ?", iFieldImageTargetId).Update(iNodeRevision_fieldImageMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateNodeRevision_fieldImageViaFieldImageAlt via map[string]interface{}{}
func UpdateNodeRevision_fieldImageViaFieldImageAlt(iFieldImageAlt string, iNodeRevision_fieldImageMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(NodeRevision_fieldImage)).Where("field_image_alt = ?", iFieldImageAlt).Update(iNodeRevision_fieldImageMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateNodeRevision_fieldImageViaFieldImageTitle via map[string]interface{}{}
func UpdateNodeRevision_fieldImageViaFieldImageTitle(iFieldImageTitle string, iNodeRevision_fieldImageMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(NodeRevision_fieldImage)).Where("field_image_title = ?", iFieldImageTitle).Update(iNodeRevision_fieldImageMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateNodeRevision_fieldImageViaFieldImageWidth via map[string]interface{}{}
func UpdateNodeRevision_fieldImageViaFieldImageWidth(iFieldImageWidth int, iNodeRevision_fieldImageMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(NodeRevision_fieldImage)).Where("field_image_width = ?", iFieldImageWidth).Update(iNodeRevision_fieldImageMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateNodeRevision_fieldImageViaFieldImageHeight via map[string]interface{}{}
func UpdateNodeRevision_fieldImageViaFieldImageHeight(iFieldImageHeight int, iNodeRevision_fieldImageMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(NodeRevision_fieldImage)).Where("field_image_height = ?", iFieldImageHeight).Update(iNodeRevision_fieldImageMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// DeleteNodeRevision_fieldImage Delete NodeRevision_fieldImage
func DeleteNodeRevision_fieldImage(iBundle string) (int64, error) {
	row, err := Engine.Id(iBundle).Delete(new(NodeRevision_fieldImage))
	return row, err
}

// DeleteNodeRevision_fieldImageViaBundle Delete NodeRevision_fieldImage via Bundle
func DeleteNodeRevision_fieldImageViaBundle(iBundle string) (err error) {
	var has bool
	var _NodeRevision_fieldImage = &NodeRevision_fieldImage{Bundle: iBundle}
	if has, err = Engine.Get(_NodeRevision_fieldImage); (has == true) && (err == nil) {
		if row, err := Engine.Where("bundle = ?", iBundle).Delete(new(NodeRevision_fieldImage)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteNodeRevision_fieldImageViaDeleted Delete NodeRevision_fieldImage via Deleted
func DeleteNodeRevision_fieldImageViaDeleted(iDeleted int) (err error) {
	var has bool
	var _NodeRevision_fieldImage = &NodeRevision_fieldImage{Deleted: iDeleted}
	if has, err = Engine.Get(_NodeRevision_fieldImage); (has == true) && (err == nil) {
		if row, err := Engine.Where("deleted = ?", iDeleted).Delete(new(NodeRevision_fieldImage)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteNodeRevision_fieldImageViaEntityId Delete NodeRevision_fieldImage via EntityId
func DeleteNodeRevision_fieldImageViaEntityId(iEntityId int) (err error) {
	var has bool
	var _NodeRevision_fieldImage = &NodeRevision_fieldImage{EntityId: iEntityId}
	if has, err = Engine.Get(_NodeRevision_fieldImage); (has == true) && (err == nil) {
		if row, err := Engine.Where("entity_id = ?", iEntityId).Delete(new(NodeRevision_fieldImage)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteNodeRevision_fieldImageViaRevisionId Delete NodeRevision_fieldImage via RevisionId
func DeleteNodeRevision_fieldImageViaRevisionId(iRevisionId int) (err error) {
	var has bool
	var _NodeRevision_fieldImage = &NodeRevision_fieldImage{RevisionId: iRevisionId}
	if has, err = Engine.Get(_NodeRevision_fieldImage); (has == true) && (err == nil) {
		if row, err := Engine.Where("revision_id = ?", iRevisionId).Delete(new(NodeRevision_fieldImage)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteNodeRevision_fieldImageViaLangcode Delete NodeRevision_fieldImage via Langcode
func DeleteNodeRevision_fieldImageViaLangcode(iLangcode string) (err error) {
	var has bool
	var _NodeRevision_fieldImage = &NodeRevision_fieldImage{Langcode: iLangcode}
	if has, err = Engine.Get(_NodeRevision_fieldImage); (has == true) && (err == nil) {
		if row, err := Engine.Where("langcode = ?", iLangcode).Delete(new(NodeRevision_fieldImage)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteNodeRevision_fieldImageViaDelta Delete NodeRevision_fieldImage via Delta
func DeleteNodeRevision_fieldImageViaDelta(iDelta int) (err error) {
	var has bool
	var _NodeRevision_fieldImage = &NodeRevision_fieldImage{Delta: iDelta}
	if has, err = Engine.Get(_NodeRevision_fieldImage); (has == true) && (err == nil) {
		if row, err := Engine.Where("delta = ?", iDelta).Delete(new(NodeRevision_fieldImage)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteNodeRevision_fieldImageViaFieldImageTargetId Delete NodeRevision_fieldImage via FieldImageTargetId
func DeleteNodeRevision_fieldImageViaFieldImageTargetId(iFieldImageTargetId int) (err error) {
	var has bool
	var _NodeRevision_fieldImage = &NodeRevision_fieldImage{FieldImageTargetId: iFieldImageTargetId}
	if has, err = Engine.Get(_NodeRevision_fieldImage); (has == true) && (err == nil) {
		if row, err := Engine.Where("field_image_target_id = ?", iFieldImageTargetId).Delete(new(NodeRevision_fieldImage)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteNodeRevision_fieldImageViaFieldImageAlt Delete NodeRevision_fieldImage via FieldImageAlt
func DeleteNodeRevision_fieldImageViaFieldImageAlt(iFieldImageAlt string) (err error) {
	var has bool
	var _NodeRevision_fieldImage = &NodeRevision_fieldImage{FieldImageAlt: iFieldImageAlt}
	if has, err = Engine.Get(_NodeRevision_fieldImage); (has == true) && (err == nil) {
		if row, err := Engine.Where("field_image_alt = ?", iFieldImageAlt).Delete(new(NodeRevision_fieldImage)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteNodeRevision_fieldImageViaFieldImageTitle Delete NodeRevision_fieldImage via FieldImageTitle
func DeleteNodeRevision_fieldImageViaFieldImageTitle(iFieldImageTitle string) (err error) {
	var has bool
	var _NodeRevision_fieldImage = &NodeRevision_fieldImage{FieldImageTitle: iFieldImageTitle}
	if has, err = Engine.Get(_NodeRevision_fieldImage); (has == true) && (err == nil) {
		if row, err := Engine.Where("field_image_title = ?", iFieldImageTitle).Delete(new(NodeRevision_fieldImage)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteNodeRevision_fieldImageViaFieldImageWidth Delete NodeRevision_fieldImage via FieldImageWidth
func DeleteNodeRevision_fieldImageViaFieldImageWidth(iFieldImageWidth int) (err error) {
	var has bool
	var _NodeRevision_fieldImage = &NodeRevision_fieldImage{FieldImageWidth: iFieldImageWidth}
	if has, err = Engine.Get(_NodeRevision_fieldImage); (has == true) && (err == nil) {
		if row, err := Engine.Where("field_image_width = ?", iFieldImageWidth).Delete(new(NodeRevision_fieldImage)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteNodeRevision_fieldImageViaFieldImageHeight Delete NodeRevision_fieldImage via FieldImageHeight
func DeleteNodeRevision_fieldImageViaFieldImageHeight(iFieldImageHeight int) (err error) {
	var has bool
	var _NodeRevision_fieldImage = &NodeRevision_fieldImage{FieldImageHeight: iFieldImageHeight}
	if has, err = Engine.Get(_NodeRevision_fieldImage); (has == true) && (err == nil) {
		if row, err := Engine.Where("field_image_height = ?", iFieldImageHeight).Delete(new(NodeRevision_fieldImage)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}
