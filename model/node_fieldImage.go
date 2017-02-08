package model

type Node_fieldImage struct {
	Bundle             string `xorm:"not null default '' index VARCHAR(128)"`
	Deleted            int    `xorm:"not null pk default 0 TINYINT(4)"`
	EntityId           int    `xorm:"not null pk INT(10)"`
	RevisionId         int    `xorm:"not null index INT(10)"`
	Langcode           string `xorm:"not null pk default '' VARCHAR(32)"`
	Delta              int    `xorm:"not null pk INT(10)"`
	FieldImageTargetId int    `xorm:"not null index INT(10)"`
	FieldImageAlt      string `xorm:"VARCHAR(512)"`
	FieldImageTitle    string `xorm:"VARCHAR(1024)"`
	FieldImageWidth    int    `xorm:"INT(10)"`
	FieldImageHeight   int    `xorm:"INT(10)"`
}

// GetNode_fieldImagesCount Node_fieldImages' Count
func GetNode_fieldImagesCount(offset int, limit int) (int64, error) {
	total, err := Engine.Limit(limit, offset).Count(&Node_fieldImage{})
	return total, err
}

// GetNode_fieldImageCountViaBundle Get Node_fieldImage via Bundle
func GetNode_fieldImageCountViaBundle(iBundle string) int64 {
	n, _ := Engine.Where("bundle = ?", iBundle).Count(&Node_fieldImage{Bundle: iBundle})
	return n
}

// GetNode_fieldImageCountViaDeleted Get Node_fieldImage via Deleted
func GetNode_fieldImageCountViaDeleted(iDeleted int) int64 {
	n, _ := Engine.Where("deleted = ?", iDeleted).Count(&Node_fieldImage{Deleted: iDeleted})
	return n
}

// GetNode_fieldImageCountViaEntityId Get Node_fieldImage via EntityId
func GetNode_fieldImageCountViaEntityId(iEntityId int) int64 {
	n, _ := Engine.Where("entity_id = ?", iEntityId).Count(&Node_fieldImage{EntityId: iEntityId})
	return n
}

// GetNode_fieldImageCountViaRevisionId Get Node_fieldImage via RevisionId
func GetNode_fieldImageCountViaRevisionId(iRevisionId int) int64 {
	n, _ := Engine.Where("revision_id = ?", iRevisionId).Count(&Node_fieldImage{RevisionId: iRevisionId})
	return n
}

// GetNode_fieldImageCountViaLangcode Get Node_fieldImage via Langcode
func GetNode_fieldImageCountViaLangcode(iLangcode string) int64 {
	n, _ := Engine.Where("langcode = ?", iLangcode).Count(&Node_fieldImage{Langcode: iLangcode})
	return n
}

// GetNode_fieldImageCountViaDelta Get Node_fieldImage via Delta
func GetNode_fieldImageCountViaDelta(iDelta int) int64 {
	n, _ := Engine.Where("delta = ?", iDelta).Count(&Node_fieldImage{Delta: iDelta})
	return n
}

// GetNode_fieldImageCountViaFieldImageTargetId Get Node_fieldImage via FieldImageTargetId
func GetNode_fieldImageCountViaFieldImageTargetId(iFieldImageTargetId int) int64 {
	n, _ := Engine.Where("field_image_target_id = ?", iFieldImageTargetId).Count(&Node_fieldImage{FieldImageTargetId: iFieldImageTargetId})
	return n
}

// GetNode_fieldImageCountViaFieldImageAlt Get Node_fieldImage via FieldImageAlt
func GetNode_fieldImageCountViaFieldImageAlt(iFieldImageAlt string) int64 {
	n, _ := Engine.Where("field_image_alt = ?", iFieldImageAlt).Count(&Node_fieldImage{FieldImageAlt: iFieldImageAlt})
	return n
}

// GetNode_fieldImageCountViaFieldImageTitle Get Node_fieldImage via FieldImageTitle
func GetNode_fieldImageCountViaFieldImageTitle(iFieldImageTitle string) int64 {
	n, _ := Engine.Where("field_image_title = ?", iFieldImageTitle).Count(&Node_fieldImage{FieldImageTitle: iFieldImageTitle})
	return n
}

// GetNode_fieldImageCountViaFieldImageWidth Get Node_fieldImage via FieldImageWidth
func GetNode_fieldImageCountViaFieldImageWidth(iFieldImageWidth int) int64 {
	n, _ := Engine.Where("field_image_width = ?", iFieldImageWidth).Count(&Node_fieldImage{FieldImageWidth: iFieldImageWidth})
	return n
}

// GetNode_fieldImageCountViaFieldImageHeight Get Node_fieldImage via FieldImageHeight
func GetNode_fieldImageCountViaFieldImageHeight(iFieldImageHeight int) int64 {
	n, _ := Engine.Where("field_image_height = ?", iFieldImageHeight).Count(&Node_fieldImage{FieldImageHeight: iFieldImageHeight})
	return n
}

// GetNode_fieldImagesByBundleAndDeletedAndEntityId Get Node_fieldImages via BundleAndDeletedAndEntityId
func GetNode_fieldImagesByBundleAndDeletedAndEntityId(offset int, limit int, Bundle_ string, Deleted_ int, EntityId_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("bundle = ? and deleted = ? and entity_id = ?", Bundle_, Deleted_, EntityId_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByBundleAndDeletedAndRevisionId Get Node_fieldImages via BundleAndDeletedAndRevisionId
func GetNode_fieldImagesByBundleAndDeletedAndRevisionId(offset int, limit int, Bundle_ string, Deleted_ int, RevisionId_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("bundle = ? and deleted = ? and revision_id = ?", Bundle_, Deleted_, RevisionId_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByBundleAndDeletedAndLangcode Get Node_fieldImages via BundleAndDeletedAndLangcode
func GetNode_fieldImagesByBundleAndDeletedAndLangcode(offset int, limit int, Bundle_ string, Deleted_ int, Langcode_ string) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("bundle = ? and deleted = ? and langcode = ?", Bundle_, Deleted_, Langcode_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByBundleAndDeletedAndDelta Get Node_fieldImages via BundleAndDeletedAndDelta
func GetNode_fieldImagesByBundleAndDeletedAndDelta(offset int, limit int, Bundle_ string, Deleted_ int, Delta_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("bundle = ? and deleted = ? and delta = ?", Bundle_, Deleted_, Delta_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByBundleAndDeletedAndFieldImageTargetId Get Node_fieldImages via BundleAndDeletedAndFieldImageTargetId
func GetNode_fieldImagesByBundleAndDeletedAndFieldImageTargetId(offset int, limit int, Bundle_ string, Deleted_ int, FieldImageTargetId_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("bundle = ? and deleted = ? and field_image_target_id = ?", Bundle_, Deleted_, FieldImageTargetId_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByBundleAndDeletedAndFieldImageAlt Get Node_fieldImages via BundleAndDeletedAndFieldImageAlt
func GetNode_fieldImagesByBundleAndDeletedAndFieldImageAlt(offset int, limit int, Bundle_ string, Deleted_ int, FieldImageAlt_ string) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("bundle = ? and deleted = ? and field_image_alt = ?", Bundle_, Deleted_, FieldImageAlt_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByBundleAndDeletedAndFieldImageTitle Get Node_fieldImages via BundleAndDeletedAndFieldImageTitle
func GetNode_fieldImagesByBundleAndDeletedAndFieldImageTitle(offset int, limit int, Bundle_ string, Deleted_ int, FieldImageTitle_ string) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("bundle = ? and deleted = ? and field_image_title = ?", Bundle_, Deleted_, FieldImageTitle_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByBundleAndDeletedAndFieldImageWidth Get Node_fieldImages via BundleAndDeletedAndFieldImageWidth
func GetNode_fieldImagesByBundleAndDeletedAndFieldImageWidth(offset int, limit int, Bundle_ string, Deleted_ int, FieldImageWidth_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("bundle = ? and deleted = ? and field_image_width = ?", Bundle_, Deleted_, FieldImageWidth_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByBundleAndDeletedAndFieldImageHeight Get Node_fieldImages via BundleAndDeletedAndFieldImageHeight
func GetNode_fieldImagesByBundleAndDeletedAndFieldImageHeight(offset int, limit int, Bundle_ string, Deleted_ int, FieldImageHeight_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("bundle = ? and deleted = ? and field_image_height = ?", Bundle_, Deleted_, FieldImageHeight_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByBundleAndEntityIdAndRevisionId Get Node_fieldImages via BundleAndEntityIdAndRevisionId
func GetNode_fieldImagesByBundleAndEntityIdAndRevisionId(offset int, limit int, Bundle_ string, EntityId_ int, RevisionId_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("bundle = ? and entity_id = ? and revision_id = ?", Bundle_, EntityId_, RevisionId_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByBundleAndEntityIdAndLangcode Get Node_fieldImages via BundleAndEntityIdAndLangcode
func GetNode_fieldImagesByBundleAndEntityIdAndLangcode(offset int, limit int, Bundle_ string, EntityId_ int, Langcode_ string) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("bundle = ? and entity_id = ? and langcode = ?", Bundle_, EntityId_, Langcode_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByBundleAndEntityIdAndDelta Get Node_fieldImages via BundleAndEntityIdAndDelta
func GetNode_fieldImagesByBundleAndEntityIdAndDelta(offset int, limit int, Bundle_ string, EntityId_ int, Delta_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("bundle = ? and entity_id = ? and delta = ?", Bundle_, EntityId_, Delta_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByBundleAndEntityIdAndFieldImageTargetId Get Node_fieldImages via BundleAndEntityIdAndFieldImageTargetId
func GetNode_fieldImagesByBundleAndEntityIdAndFieldImageTargetId(offset int, limit int, Bundle_ string, EntityId_ int, FieldImageTargetId_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("bundle = ? and entity_id = ? and field_image_target_id = ?", Bundle_, EntityId_, FieldImageTargetId_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByBundleAndEntityIdAndFieldImageAlt Get Node_fieldImages via BundleAndEntityIdAndFieldImageAlt
func GetNode_fieldImagesByBundleAndEntityIdAndFieldImageAlt(offset int, limit int, Bundle_ string, EntityId_ int, FieldImageAlt_ string) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("bundle = ? and entity_id = ? and field_image_alt = ?", Bundle_, EntityId_, FieldImageAlt_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByBundleAndEntityIdAndFieldImageTitle Get Node_fieldImages via BundleAndEntityIdAndFieldImageTitle
func GetNode_fieldImagesByBundleAndEntityIdAndFieldImageTitle(offset int, limit int, Bundle_ string, EntityId_ int, FieldImageTitle_ string) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("bundle = ? and entity_id = ? and field_image_title = ?", Bundle_, EntityId_, FieldImageTitle_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByBundleAndEntityIdAndFieldImageWidth Get Node_fieldImages via BundleAndEntityIdAndFieldImageWidth
func GetNode_fieldImagesByBundleAndEntityIdAndFieldImageWidth(offset int, limit int, Bundle_ string, EntityId_ int, FieldImageWidth_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("bundle = ? and entity_id = ? and field_image_width = ?", Bundle_, EntityId_, FieldImageWidth_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByBundleAndEntityIdAndFieldImageHeight Get Node_fieldImages via BundleAndEntityIdAndFieldImageHeight
func GetNode_fieldImagesByBundleAndEntityIdAndFieldImageHeight(offset int, limit int, Bundle_ string, EntityId_ int, FieldImageHeight_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("bundle = ? and entity_id = ? and field_image_height = ?", Bundle_, EntityId_, FieldImageHeight_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByBundleAndRevisionIdAndLangcode Get Node_fieldImages via BundleAndRevisionIdAndLangcode
func GetNode_fieldImagesByBundleAndRevisionIdAndLangcode(offset int, limit int, Bundle_ string, RevisionId_ int, Langcode_ string) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("bundle = ? and revision_id = ? and langcode = ?", Bundle_, RevisionId_, Langcode_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByBundleAndRevisionIdAndDelta Get Node_fieldImages via BundleAndRevisionIdAndDelta
func GetNode_fieldImagesByBundleAndRevisionIdAndDelta(offset int, limit int, Bundle_ string, RevisionId_ int, Delta_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("bundle = ? and revision_id = ? and delta = ?", Bundle_, RevisionId_, Delta_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByBundleAndRevisionIdAndFieldImageTargetId Get Node_fieldImages via BundleAndRevisionIdAndFieldImageTargetId
func GetNode_fieldImagesByBundleAndRevisionIdAndFieldImageTargetId(offset int, limit int, Bundle_ string, RevisionId_ int, FieldImageTargetId_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("bundle = ? and revision_id = ? and field_image_target_id = ?", Bundle_, RevisionId_, FieldImageTargetId_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByBundleAndRevisionIdAndFieldImageAlt Get Node_fieldImages via BundleAndRevisionIdAndFieldImageAlt
func GetNode_fieldImagesByBundleAndRevisionIdAndFieldImageAlt(offset int, limit int, Bundle_ string, RevisionId_ int, FieldImageAlt_ string) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("bundle = ? and revision_id = ? and field_image_alt = ?", Bundle_, RevisionId_, FieldImageAlt_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByBundleAndRevisionIdAndFieldImageTitle Get Node_fieldImages via BundleAndRevisionIdAndFieldImageTitle
func GetNode_fieldImagesByBundleAndRevisionIdAndFieldImageTitle(offset int, limit int, Bundle_ string, RevisionId_ int, FieldImageTitle_ string) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("bundle = ? and revision_id = ? and field_image_title = ?", Bundle_, RevisionId_, FieldImageTitle_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByBundleAndRevisionIdAndFieldImageWidth Get Node_fieldImages via BundleAndRevisionIdAndFieldImageWidth
func GetNode_fieldImagesByBundleAndRevisionIdAndFieldImageWidth(offset int, limit int, Bundle_ string, RevisionId_ int, FieldImageWidth_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("bundle = ? and revision_id = ? and field_image_width = ?", Bundle_, RevisionId_, FieldImageWidth_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByBundleAndRevisionIdAndFieldImageHeight Get Node_fieldImages via BundleAndRevisionIdAndFieldImageHeight
func GetNode_fieldImagesByBundleAndRevisionIdAndFieldImageHeight(offset int, limit int, Bundle_ string, RevisionId_ int, FieldImageHeight_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("bundle = ? and revision_id = ? and field_image_height = ?", Bundle_, RevisionId_, FieldImageHeight_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByBundleAndLangcodeAndDelta Get Node_fieldImages via BundleAndLangcodeAndDelta
func GetNode_fieldImagesByBundleAndLangcodeAndDelta(offset int, limit int, Bundle_ string, Langcode_ string, Delta_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("bundle = ? and langcode = ? and delta = ?", Bundle_, Langcode_, Delta_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByBundleAndLangcodeAndFieldImageTargetId Get Node_fieldImages via BundleAndLangcodeAndFieldImageTargetId
func GetNode_fieldImagesByBundleAndLangcodeAndFieldImageTargetId(offset int, limit int, Bundle_ string, Langcode_ string, FieldImageTargetId_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("bundle = ? and langcode = ? and field_image_target_id = ?", Bundle_, Langcode_, FieldImageTargetId_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByBundleAndLangcodeAndFieldImageAlt Get Node_fieldImages via BundleAndLangcodeAndFieldImageAlt
func GetNode_fieldImagesByBundleAndLangcodeAndFieldImageAlt(offset int, limit int, Bundle_ string, Langcode_ string, FieldImageAlt_ string) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("bundle = ? and langcode = ? and field_image_alt = ?", Bundle_, Langcode_, FieldImageAlt_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByBundleAndLangcodeAndFieldImageTitle Get Node_fieldImages via BundleAndLangcodeAndFieldImageTitle
func GetNode_fieldImagesByBundleAndLangcodeAndFieldImageTitle(offset int, limit int, Bundle_ string, Langcode_ string, FieldImageTitle_ string) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("bundle = ? and langcode = ? and field_image_title = ?", Bundle_, Langcode_, FieldImageTitle_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByBundleAndLangcodeAndFieldImageWidth Get Node_fieldImages via BundleAndLangcodeAndFieldImageWidth
func GetNode_fieldImagesByBundleAndLangcodeAndFieldImageWidth(offset int, limit int, Bundle_ string, Langcode_ string, FieldImageWidth_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("bundle = ? and langcode = ? and field_image_width = ?", Bundle_, Langcode_, FieldImageWidth_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByBundleAndLangcodeAndFieldImageHeight Get Node_fieldImages via BundleAndLangcodeAndFieldImageHeight
func GetNode_fieldImagesByBundleAndLangcodeAndFieldImageHeight(offset int, limit int, Bundle_ string, Langcode_ string, FieldImageHeight_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("bundle = ? and langcode = ? and field_image_height = ?", Bundle_, Langcode_, FieldImageHeight_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByBundleAndDeltaAndFieldImageTargetId Get Node_fieldImages via BundleAndDeltaAndFieldImageTargetId
func GetNode_fieldImagesByBundleAndDeltaAndFieldImageTargetId(offset int, limit int, Bundle_ string, Delta_ int, FieldImageTargetId_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("bundle = ? and delta = ? and field_image_target_id = ?", Bundle_, Delta_, FieldImageTargetId_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByBundleAndDeltaAndFieldImageAlt Get Node_fieldImages via BundleAndDeltaAndFieldImageAlt
func GetNode_fieldImagesByBundleAndDeltaAndFieldImageAlt(offset int, limit int, Bundle_ string, Delta_ int, FieldImageAlt_ string) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("bundle = ? and delta = ? and field_image_alt = ?", Bundle_, Delta_, FieldImageAlt_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByBundleAndDeltaAndFieldImageTitle Get Node_fieldImages via BundleAndDeltaAndFieldImageTitle
func GetNode_fieldImagesByBundleAndDeltaAndFieldImageTitle(offset int, limit int, Bundle_ string, Delta_ int, FieldImageTitle_ string) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("bundle = ? and delta = ? and field_image_title = ?", Bundle_, Delta_, FieldImageTitle_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByBundleAndDeltaAndFieldImageWidth Get Node_fieldImages via BundleAndDeltaAndFieldImageWidth
func GetNode_fieldImagesByBundleAndDeltaAndFieldImageWidth(offset int, limit int, Bundle_ string, Delta_ int, FieldImageWidth_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("bundle = ? and delta = ? and field_image_width = ?", Bundle_, Delta_, FieldImageWidth_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByBundleAndDeltaAndFieldImageHeight Get Node_fieldImages via BundleAndDeltaAndFieldImageHeight
func GetNode_fieldImagesByBundleAndDeltaAndFieldImageHeight(offset int, limit int, Bundle_ string, Delta_ int, FieldImageHeight_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("bundle = ? and delta = ? and field_image_height = ?", Bundle_, Delta_, FieldImageHeight_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByBundleAndFieldImageTargetIdAndFieldImageAlt Get Node_fieldImages via BundleAndFieldImageTargetIdAndFieldImageAlt
func GetNode_fieldImagesByBundleAndFieldImageTargetIdAndFieldImageAlt(offset int, limit int, Bundle_ string, FieldImageTargetId_ int, FieldImageAlt_ string) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("bundle = ? and field_image_target_id = ? and field_image_alt = ?", Bundle_, FieldImageTargetId_, FieldImageAlt_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByBundleAndFieldImageTargetIdAndFieldImageTitle Get Node_fieldImages via BundleAndFieldImageTargetIdAndFieldImageTitle
func GetNode_fieldImagesByBundleAndFieldImageTargetIdAndFieldImageTitle(offset int, limit int, Bundle_ string, FieldImageTargetId_ int, FieldImageTitle_ string) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("bundle = ? and field_image_target_id = ? and field_image_title = ?", Bundle_, FieldImageTargetId_, FieldImageTitle_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByBundleAndFieldImageTargetIdAndFieldImageWidth Get Node_fieldImages via BundleAndFieldImageTargetIdAndFieldImageWidth
func GetNode_fieldImagesByBundleAndFieldImageTargetIdAndFieldImageWidth(offset int, limit int, Bundle_ string, FieldImageTargetId_ int, FieldImageWidth_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("bundle = ? and field_image_target_id = ? and field_image_width = ?", Bundle_, FieldImageTargetId_, FieldImageWidth_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByBundleAndFieldImageTargetIdAndFieldImageHeight Get Node_fieldImages via BundleAndFieldImageTargetIdAndFieldImageHeight
func GetNode_fieldImagesByBundleAndFieldImageTargetIdAndFieldImageHeight(offset int, limit int, Bundle_ string, FieldImageTargetId_ int, FieldImageHeight_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("bundle = ? and field_image_target_id = ? and field_image_height = ?", Bundle_, FieldImageTargetId_, FieldImageHeight_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByBundleAndFieldImageAltAndFieldImageTitle Get Node_fieldImages via BundleAndFieldImageAltAndFieldImageTitle
func GetNode_fieldImagesByBundleAndFieldImageAltAndFieldImageTitle(offset int, limit int, Bundle_ string, FieldImageAlt_ string, FieldImageTitle_ string) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("bundle = ? and field_image_alt = ? and field_image_title = ?", Bundle_, FieldImageAlt_, FieldImageTitle_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByBundleAndFieldImageAltAndFieldImageWidth Get Node_fieldImages via BundleAndFieldImageAltAndFieldImageWidth
func GetNode_fieldImagesByBundleAndFieldImageAltAndFieldImageWidth(offset int, limit int, Bundle_ string, FieldImageAlt_ string, FieldImageWidth_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("bundle = ? and field_image_alt = ? and field_image_width = ?", Bundle_, FieldImageAlt_, FieldImageWidth_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByBundleAndFieldImageAltAndFieldImageHeight Get Node_fieldImages via BundleAndFieldImageAltAndFieldImageHeight
func GetNode_fieldImagesByBundleAndFieldImageAltAndFieldImageHeight(offset int, limit int, Bundle_ string, FieldImageAlt_ string, FieldImageHeight_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("bundle = ? and field_image_alt = ? and field_image_height = ?", Bundle_, FieldImageAlt_, FieldImageHeight_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByBundleAndFieldImageTitleAndFieldImageWidth Get Node_fieldImages via BundleAndFieldImageTitleAndFieldImageWidth
func GetNode_fieldImagesByBundleAndFieldImageTitleAndFieldImageWidth(offset int, limit int, Bundle_ string, FieldImageTitle_ string, FieldImageWidth_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("bundle = ? and field_image_title = ? and field_image_width = ?", Bundle_, FieldImageTitle_, FieldImageWidth_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByBundleAndFieldImageTitleAndFieldImageHeight Get Node_fieldImages via BundleAndFieldImageTitleAndFieldImageHeight
func GetNode_fieldImagesByBundleAndFieldImageTitleAndFieldImageHeight(offset int, limit int, Bundle_ string, FieldImageTitle_ string, FieldImageHeight_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("bundle = ? and field_image_title = ? and field_image_height = ?", Bundle_, FieldImageTitle_, FieldImageHeight_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByBundleAndFieldImageWidthAndFieldImageHeight Get Node_fieldImages via BundleAndFieldImageWidthAndFieldImageHeight
func GetNode_fieldImagesByBundleAndFieldImageWidthAndFieldImageHeight(offset int, limit int, Bundle_ string, FieldImageWidth_ int, FieldImageHeight_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("bundle = ? and field_image_width = ? and field_image_height = ?", Bundle_, FieldImageWidth_, FieldImageHeight_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByDeletedAndEntityIdAndRevisionId Get Node_fieldImages via DeletedAndEntityIdAndRevisionId
func GetNode_fieldImagesByDeletedAndEntityIdAndRevisionId(offset int, limit int, Deleted_ int, EntityId_ int, RevisionId_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("deleted = ? and entity_id = ? and revision_id = ?", Deleted_, EntityId_, RevisionId_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByDeletedAndEntityIdAndLangcode Get Node_fieldImages via DeletedAndEntityIdAndLangcode
func GetNode_fieldImagesByDeletedAndEntityIdAndLangcode(offset int, limit int, Deleted_ int, EntityId_ int, Langcode_ string) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("deleted = ? and entity_id = ? and langcode = ?", Deleted_, EntityId_, Langcode_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByDeletedAndEntityIdAndDelta Get Node_fieldImages via DeletedAndEntityIdAndDelta
func GetNode_fieldImagesByDeletedAndEntityIdAndDelta(offset int, limit int, Deleted_ int, EntityId_ int, Delta_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("deleted = ? and entity_id = ? and delta = ?", Deleted_, EntityId_, Delta_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByDeletedAndEntityIdAndFieldImageTargetId Get Node_fieldImages via DeletedAndEntityIdAndFieldImageTargetId
func GetNode_fieldImagesByDeletedAndEntityIdAndFieldImageTargetId(offset int, limit int, Deleted_ int, EntityId_ int, FieldImageTargetId_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("deleted = ? and entity_id = ? and field_image_target_id = ?", Deleted_, EntityId_, FieldImageTargetId_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByDeletedAndEntityIdAndFieldImageAlt Get Node_fieldImages via DeletedAndEntityIdAndFieldImageAlt
func GetNode_fieldImagesByDeletedAndEntityIdAndFieldImageAlt(offset int, limit int, Deleted_ int, EntityId_ int, FieldImageAlt_ string) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("deleted = ? and entity_id = ? and field_image_alt = ?", Deleted_, EntityId_, FieldImageAlt_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByDeletedAndEntityIdAndFieldImageTitle Get Node_fieldImages via DeletedAndEntityIdAndFieldImageTitle
func GetNode_fieldImagesByDeletedAndEntityIdAndFieldImageTitle(offset int, limit int, Deleted_ int, EntityId_ int, FieldImageTitle_ string) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("deleted = ? and entity_id = ? and field_image_title = ?", Deleted_, EntityId_, FieldImageTitle_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByDeletedAndEntityIdAndFieldImageWidth Get Node_fieldImages via DeletedAndEntityIdAndFieldImageWidth
func GetNode_fieldImagesByDeletedAndEntityIdAndFieldImageWidth(offset int, limit int, Deleted_ int, EntityId_ int, FieldImageWidth_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("deleted = ? and entity_id = ? and field_image_width = ?", Deleted_, EntityId_, FieldImageWidth_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByDeletedAndEntityIdAndFieldImageHeight Get Node_fieldImages via DeletedAndEntityIdAndFieldImageHeight
func GetNode_fieldImagesByDeletedAndEntityIdAndFieldImageHeight(offset int, limit int, Deleted_ int, EntityId_ int, FieldImageHeight_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("deleted = ? and entity_id = ? and field_image_height = ?", Deleted_, EntityId_, FieldImageHeight_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByDeletedAndRevisionIdAndLangcode Get Node_fieldImages via DeletedAndRevisionIdAndLangcode
func GetNode_fieldImagesByDeletedAndRevisionIdAndLangcode(offset int, limit int, Deleted_ int, RevisionId_ int, Langcode_ string) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("deleted = ? and revision_id = ? and langcode = ?", Deleted_, RevisionId_, Langcode_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByDeletedAndRevisionIdAndDelta Get Node_fieldImages via DeletedAndRevisionIdAndDelta
func GetNode_fieldImagesByDeletedAndRevisionIdAndDelta(offset int, limit int, Deleted_ int, RevisionId_ int, Delta_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("deleted = ? and revision_id = ? and delta = ?", Deleted_, RevisionId_, Delta_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByDeletedAndRevisionIdAndFieldImageTargetId Get Node_fieldImages via DeletedAndRevisionIdAndFieldImageTargetId
func GetNode_fieldImagesByDeletedAndRevisionIdAndFieldImageTargetId(offset int, limit int, Deleted_ int, RevisionId_ int, FieldImageTargetId_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("deleted = ? and revision_id = ? and field_image_target_id = ?", Deleted_, RevisionId_, FieldImageTargetId_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByDeletedAndRevisionIdAndFieldImageAlt Get Node_fieldImages via DeletedAndRevisionIdAndFieldImageAlt
func GetNode_fieldImagesByDeletedAndRevisionIdAndFieldImageAlt(offset int, limit int, Deleted_ int, RevisionId_ int, FieldImageAlt_ string) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("deleted = ? and revision_id = ? and field_image_alt = ?", Deleted_, RevisionId_, FieldImageAlt_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByDeletedAndRevisionIdAndFieldImageTitle Get Node_fieldImages via DeletedAndRevisionIdAndFieldImageTitle
func GetNode_fieldImagesByDeletedAndRevisionIdAndFieldImageTitle(offset int, limit int, Deleted_ int, RevisionId_ int, FieldImageTitle_ string) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("deleted = ? and revision_id = ? and field_image_title = ?", Deleted_, RevisionId_, FieldImageTitle_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByDeletedAndRevisionIdAndFieldImageWidth Get Node_fieldImages via DeletedAndRevisionIdAndFieldImageWidth
func GetNode_fieldImagesByDeletedAndRevisionIdAndFieldImageWidth(offset int, limit int, Deleted_ int, RevisionId_ int, FieldImageWidth_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("deleted = ? and revision_id = ? and field_image_width = ?", Deleted_, RevisionId_, FieldImageWidth_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByDeletedAndRevisionIdAndFieldImageHeight Get Node_fieldImages via DeletedAndRevisionIdAndFieldImageHeight
func GetNode_fieldImagesByDeletedAndRevisionIdAndFieldImageHeight(offset int, limit int, Deleted_ int, RevisionId_ int, FieldImageHeight_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("deleted = ? and revision_id = ? and field_image_height = ?", Deleted_, RevisionId_, FieldImageHeight_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByDeletedAndLangcodeAndDelta Get Node_fieldImages via DeletedAndLangcodeAndDelta
func GetNode_fieldImagesByDeletedAndLangcodeAndDelta(offset int, limit int, Deleted_ int, Langcode_ string, Delta_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("deleted = ? and langcode = ? and delta = ?", Deleted_, Langcode_, Delta_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByDeletedAndLangcodeAndFieldImageTargetId Get Node_fieldImages via DeletedAndLangcodeAndFieldImageTargetId
func GetNode_fieldImagesByDeletedAndLangcodeAndFieldImageTargetId(offset int, limit int, Deleted_ int, Langcode_ string, FieldImageTargetId_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("deleted = ? and langcode = ? and field_image_target_id = ?", Deleted_, Langcode_, FieldImageTargetId_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByDeletedAndLangcodeAndFieldImageAlt Get Node_fieldImages via DeletedAndLangcodeAndFieldImageAlt
func GetNode_fieldImagesByDeletedAndLangcodeAndFieldImageAlt(offset int, limit int, Deleted_ int, Langcode_ string, FieldImageAlt_ string) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("deleted = ? and langcode = ? and field_image_alt = ?", Deleted_, Langcode_, FieldImageAlt_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByDeletedAndLangcodeAndFieldImageTitle Get Node_fieldImages via DeletedAndLangcodeAndFieldImageTitle
func GetNode_fieldImagesByDeletedAndLangcodeAndFieldImageTitle(offset int, limit int, Deleted_ int, Langcode_ string, FieldImageTitle_ string) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("deleted = ? and langcode = ? and field_image_title = ?", Deleted_, Langcode_, FieldImageTitle_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByDeletedAndLangcodeAndFieldImageWidth Get Node_fieldImages via DeletedAndLangcodeAndFieldImageWidth
func GetNode_fieldImagesByDeletedAndLangcodeAndFieldImageWidth(offset int, limit int, Deleted_ int, Langcode_ string, FieldImageWidth_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("deleted = ? and langcode = ? and field_image_width = ?", Deleted_, Langcode_, FieldImageWidth_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByDeletedAndLangcodeAndFieldImageHeight Get Node_fieldImages via DeletedAndLangcodeAndFieldImageHeight
func GetNode_fieldImagesByDeletedAndLangcodeAndFieldImageHeight(offset int, limit int, Deleted_ int, Langcode_ string, FieldImageHeight_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("deleted = ? and langcode = ? and field_image_height = ?", Deleted_, Langcode_, FieldImageHeight_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByDeletedAndDeltaAndFieldImageTargetId Get Node_fieldImages via DeletedAndDeltaAndFieldImageTargetId
func GetNode_fieldImagesByDeletedAndDeltaAndFieldImageTargetId(offset int, limit int, Deleted_ int, Delta_ int, FieldImageTargetId_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("deleted = ? and delta = ? and field_image_target_id = ?", Deleted_, Delta_, FieldImageTargetId_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByDeletedAndDeltaAndFieldImageAlt Get Node_fieldImages via DeletedAndDeltaAndFieldImageAlt
func GetNode_fieldImagesByDeletedAndDeltaAndFieldImageAlt(offset int, limit int, Deleted_ int, Delta_ int, FieldImageAlt_ string) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("deleted = ? and delta = ? and field_image_alt = ?", Deleted_, Delta_, FieldImageAlt_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByDeletedAndDeltaAndFieldImageTitle Get Node_fieldImages via DeletedAndDeltaAndFieldImageTitle
func GetNode_fieldImagesByDeletedAndDeltaAndFieldImageTitle(offset int, limit int, Deleted_ int, Delta_ int, FieldImageTitle_ string) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("deleted = ? and delta = ? and field_image_title = ?", Deleted_, Delta_, FieldImageTitle_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByDeletedAndDeltaAndFieldImageWidth Get Node_fieldImages via DeletedAndDeltaAndFieldImageWidth
func GetNode_fieldImagesByDeletedAndDeltaAndFieldImageWidth(offset int, limit int, Deleted_ int, Delta_ int, FieldImageWidth_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("deleted = ? and delta = ? and field_image_width = ?", Deleted_, Delta_, FieldImageWidth_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByDeletedAndDeltaAndFieldImageHeight Get Node_fieldImages via DeletedAndDeltaAndFieldImageHeight
func GetNode_fieldImagesByDeletedAndDeltaAndFieldImageHeight(offset int, limit int, Deleted_ int, Delta_ int, FieldImageHeight_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("deleted = ? and delta = ? and field_image_height = ?", Deleted_, Delta_, FieldImageHeight_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByDeletedAndFieldImageTargetIdAndFieldImageAlt Get Node_fieldImages via DeletedAndFieldImageTargetIdAndFieldImageAlt
func GetNode_fieldImagesByDeletedAndFieldImageTargetIdAndFieldImageAlt(offset int, limit int, Deleted_ int, FieldImageTargetId_ int, FieldImageAlt_ string) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("deleted = ? and field_image_target_id = ? and field_image_alt = ?", Deleted_, FieldImageTargetId_, FieldImageAlt_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByDeletedAndFieldImageTargetIdAndFieldImageTitle Get Node_fieldImages via DeletedAndFieldImageTargetIdAndFieldImageTitle
func GetNode_fieldImagesByDeletedAndFieldImageTargetIdAndFieldImageTitle(offset int, limit int, Deleted_ int, FieldImageTargetId_ int, FieldImageTitle_ string) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("deleted = ? and field_image_target_id = ? and field_image_title = ?", Deleted_, FieldImageTargetId_, FieldImageTitle_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByDeletedAndFieldImageTargetIdAndFieldImageWidth Get Node_fieldImages via DeletedAndFieldImageTargetIdAndFieldImageWidth
func GetNode_fieldImagesByDeletedAndFieldImageTargetIdAndFieldImageWidth(offset int, limit int, Deleted_ int, FieldImageTargetId_ int, FieldImageWidth_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("deleted = ? and field_image_target_id = ? and field_image_width = ?", Deleted_, FieldImageTargetId_, FieldImageWidth_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByDeletedAndFieldImageTargetIdAndFieldImageHeight Get Node_fieldImages via DeletedAndFieldImageTargetIdAndFieldImageHeight
func GetNode_fieldImagesByDeletedAndFieldImageTargetIdAndFieldImageHeight(offset int, limit int, Deleted_ int, FieldImageTargetId_ int, FieldImageHeight_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("deleted = ? and field_image_target_id = ? and field_image_height = ?", Deleted_, FieldImageTargetId_, FieldImageHeight_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByDeletedAndFieldImageAltAndFieldImageTitle Get Node_fieldImages via DeletedAndFieldImageAltAndFieldImageTitle
func GetNode_fieldImagesByDeletedAndFieldImageAltAndFieldImageTitle(offset int, limit int, Deleted_ int, FieldImageAlt_ string, FieldImageTitle_ string) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("deleted = ? and field_image_alt = ? and field_image_title = ?", Deleted_, FieldImageAlt_, FieldImageTitle_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByDeletedAndFieldImageAltAndFieldImageWidth Get Node_fieldImages via DeletedAndFieldImageAltAndFieldImageWidth
func GetNode_fieldImagesByDeletedAndFieldImageAltAndFieldImageWidth(offset int, limit int, Deleted_ int, FieldImageAlt_ string, FieldImageWidth_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("deleted = ? and field_image_alt = ? and field_image_width = ?", Deleted_, FieldImageAlt_, FieldImageWidth_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByDeletedAndFieldImageAltAndFieldImageHeight Get Node_fieldImages via DeletedAndFieldImageAltAndFieldImageHeight
func GetNode_fieldImagesByDeletedAndFieldImageAltAndFieldImageHeight(offset int, limit int, Deleted_ int, FieldImageAlt_ string, FieldImageHeight_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("deleted = ? and field_image_alt = ? and field_image_height = ?", Deleted_, FieldImageAlt_, FieldImageHeight_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByDeletedAndFieldImageTitleAndFieldImageWidth Get Node_fieldImages via DeletedAndFieldImageTitleAndFieldImageWidth
func GetNode_fieldImagesByDeletedAndFieldImageTitleAndFieldImageWidth(offset int, limit int, Deleted_ int, FieldImageTitle_ string, FieldImageWidth_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("deleted = ? and field_image_title = ? and field_image_width = ?", Deleted_, FieldImageTitle_, FieldImageWidth_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByDeletedAndFieldImageTitleAndFieldImageHeight Get Node_fieldImages via DeletedAndFieldImageTitleAndFieldImageHeight
func GetNode_fieldImagesByDeletedAndFieldImageTitleAndFieldImageHeight(offset int, limit int, Deleted_ int, FieldImageTitle_ string, FieldImageHeight_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("deleted = ? and field_image_title = ? and field_image_height = ?", Deleted_, FieldImageTitle_, FieldImageHeight_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByDeletedAndFieldImageWidthAndFieldImageHeight Get Node_fieldImages via DeletedAndFieldImageWidthAndFieldImageHeight
func GetNode_fieldImagesByDeletedAndFieldImageWidthAndFieldImageHeight(offset int, limit int, Deleted_ int, FieldImageWidth_ int, FieldImageHeight_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("deleted = ? and field_image_width = ? and field_image_height = ?", Deleted_, FieldImageWidth_, FieldImageHeight_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByEntityIdAndRevisionIdAndLangcode Get Node_fieldImages via EntityIdAndRevisionIdAndLangcode
func GetNode_fieldImagesByEntityIdAndRevisionIdAndLangcode(offset int, limit int, EntityId_ int, RevisionId_ int, Langcode_ string) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("entity_id = ? and revision_id = ? and langcode = ?", EntityId_, RevisionId_, Langcode_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByEntityIdAndRevisionIdAndDelta Get Node_fieldImages via EntityIdAndRevisionIdAndDelta
func GetNode_fieldImagesByEntityIdAndRevisionIdAndDelta(offset int, limit int, EntityId_ int, RevisionId_ int, Delta_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("entity_id = ? and revision_id = ? and delta = ?", EntityId_, RevisionId_, Delta_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByEntityIdAndRevisionIdAndFieldImageTargetId Get Node_fieldImages via EntityIdAndRevisionIdAndFieldImageTargetId
func GetNode_fieldImagesByEntityIdAndRevisionIdAndFieldImageTargetId(offset int, limit int, EntityId_ int, RevisionId_ int, FieldImageTargetId_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("entity_id = ? and revision_id = ? and field_image_target_id = ?", EntityId_, RevisionId_, FieldImageTargetId_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByEntityIdAndRevisionIdAndFieldImageAlt Get Node_fieldImages via EntityIdAndRevisionIdAndFieldImageAlt
func GetNode_fieldImagesByEntityIdAndRevisionIdAndFieldImageAlt(offset int, limit int, EntityId_ int, RevisionId_ int, FieldImageAlt_ string) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("entity_id = ? and revision_id = ? and field_image_alt = ?", EntityId_, RevisionId_, FieldImageAlt_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByEntityIdAndRevisionIdAndFieldImageTitle Get Node_fieldImages via EntityIdAndRevisionIdAndFieldImageTitle
func GetNode_fieldImagesByEntityIdAndRevisionIdAndFieldImageTitle(offset int, limit int, EntityId_ int, RevisionId_ int, FieldImageTitle_ string) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("entity_id = ? and revision_id = ? and field_image_title = ?", EntityId_, RevisionId_, FieldImageTitle_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByEntityIdAndRevisionIdAndFieldImageWidth Get Node_fieldImages via EntityIdAndRevisionIdAndFieldImageWidth
func GetNode_fieldImagesByEntityIdAndRevisionIdAndFieldImageWidth(offset int, limit int, EntityId_ int, RevisionId_ int, FieldImageWidth_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("entity_id = ? and revision_id = ? and field_image_width = ?", EntityId_, RevisionId_, FieldImageWidth_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByEntityIdAndRevisionIdAndFieldImageHeight Get Node_fieldImages via EntityIdAndRevisionIdAndFieldImageHeight
func GetNode_fieldImagesByEntityIdAndRevisionIdAndFieldImageHeight(offset int, limit int, EntityId_ int, RevisionId_ int, FieldImageHeight_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("entity_id = ? and revision_id = ? and field_image_height = ?", EntityId_, RevisionId_, FieldImageHeight_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByEntityIdAndLangcodeAndDelta Get Node_fieldImages via EntityIdAndLangcodeAndDelta
func GetNode_fieldImagesByEntityIdAndLangcodeAndDelta(offset int, limit int, EntityId_ int, Langcode_ string, Delta_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("entity_id = ? and langcode = ? and delta = ?", EntityId_, Langcode_, Delta_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByEntityIdAndLangcodeAndFieldImageTargetId Get Node_fieldImages via EntityIdAndLangcodeAndFieldImageTargetId
func GetNode_fieldImagesByEntityIdAndLangcodeAndFieldImageTargetId(offset int, limit int, EntityId_ int, Langcode_ string, FieldImageTargetId_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("entity_id = ? and langcode = ? and field_image_target_id = ?", EntityId_, Langcode_, FieldImageTargetId_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByEntityIdAndLangcodeAndFieldImageAlt Get Node_fieldImages via EntityIdAndLangcodeAndFieldImageAlt
func GetNode_fieldImagesByEntityIdAndLangcodeAndFieldImageAlt(offset int, limit int, EntityId_ int, Langcode_ string, FieldImageAlt_ string) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("entity_id = ? and langcode = ? and field_image_alt = ?", EntityId_, Langcode_, FieldImageAlt_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByEntityIdAndLangcodeAndFieldImageTitle Get Node_fieldImages via EntityIdAndLangcodeAndFieldImageTitle
func GetNode_fieldImagesByEntityIdAndLangcodeAndFieldImageTitle(offset int, limit int, EntityId_ int, Langcode_ string, FieldImageTitle_ string) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("entity_id = ? and langcode = ? and field_image_title = ?", EntityId_, Langcode_, FieldImageTitle_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByEntityIdAndLangcodeAndFieldImageWidth Get Node_fieldImages via EntityIdAndLangcodeAndFieldImageWidth
func GetNode_fieldImagesByEntityIdAndLangcodeAndFieldImageWidth(offset int, limit int, EntityId_ int, Langcode_ string, FieldImageWidth_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("entity_id = ? and langcode = ? and field_image_width = ?", EntityId_, Langcode_, FieldImageWidth_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByEntityIdAndLangcodeAndFieldImageHeight Get Node_fieldImages via EntityIdAndLangcodeAndFieldImageHeight
func GetNode_fieldImagesByEntityIdAndLangcodeAndFieldImageHeight(offset int, limit int, EntityId_ int, Langcode_ string, FieldImageHeight_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("entity_id = ? and langcode = ? and field_image_height = ?", EntityId_, Langcode_, FieldImageHeight_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByEntityIdAndDeltaAndFieldImageTargetId Get Node_fieldImages via EntityIdAndDeltaAndFieldImageTargetId
func GetNode_fieldImagesByEntityIdAndDeltaAndFieldImageTargetId(offset int, limit int, EntityId_ int, Delta_ int, FieldImageTargetId_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("entity_id = ? and delta = ? and field_image_target_id = ?", EntityId_, Delta_, FieldImageTargetId_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByEntityIdAndDeltaAndFieldImageAlt Get Node_fieldImages via EntityIdAndDeltaAndFieldImageAlt
func GetNode_fieldImagesByEntityIdAndDeltaAndFieldImageAlt(offset int, limit int, EntityId_ int, Delta_ int, FieldImageAlt_ string) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("entity_id = ? and delta = ? and field_image_alt = ?", EntityId_, Delta_, FieldImageAlt_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByEntityIdAndDeltaAndFieldImageTitle Get Node_fieldImages via EntityIdAndDeltaAndFieldImageTitle
func GetNode_fieldImagesByEntityIdAndDeltaAndFieldImageTitle(offset int, limit int, EntityId_ int, Delta_ int, FieldImageTitle_ string) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("entity_id = ? and delta = ? and field_image_title = ?", EntityId_, Delta_, FieldImageTitle_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByEntityIdAndDeltaAndFieldImageWidth Get Node_fieldImages via EntityIdAndDeltaAndFieldImageWidth
func GetNode_fieldImagesByEntityIdAndDeltaAndFieldImageWidth(offset int, limit int, EntityId_ int, Delta_ int, FieldImageWidth_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("entity_id = ? and delta = ? and field_image_width = ?", EntityId_, Delta_, FieldImageWidth_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByEntityIdAndDeltaAndFieldImageHeight Get Node_fieldImages via EntityIdAndDeltaAndFieldImageHeight
func GetNode_fieldImagesByEntityIdAndDeltaAndFieldImageHeight(offset int, limit int, EntityId_ int, Delta_ int, FieldImageHeight_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("entity_id = ? and delta = ? and field_image_height = ?", EntityId_, Delta_, FieldImageHeight_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByEntityIdAndFieldImageTargetIdAndFieldImageAlt Get Node_fieldImages via EntityIdAndFieldImageTargetIdAndFieldImageAlt
func GetNode_fieldImagesByEntityIdAndFieldImageTargetIdAndFieldImageAlt(offset int, limit int, EntityId_ int, FieldImageTargetId_ int, FieldImageAlt_ string) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("entity_id = ? and field_image_target_id = ? and field_image_alt = ?", EntityId_, FieldImageTargetId_, FieldImageAlt_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByEntityIdAndFieldImageTargetIdAndFieldImageTitle Get Node_fieldImages via EntityIdAndFieldImageTargetIdAndFieldImageTitle
func GetNode_fieldImagesByEntityIdAndFieldImageTargetIdAndFieldImageTitle(offset int, limit int, EntityId_ int, FieldImageTargetId_ int, FieldImageTitle_ string) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("entity_id = ? and field_image_target_id = ? and field_image_title = ?", EntityId_, FieldImageTargetId_, FieldImageTitle_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByEntityIdAndFieldImageTargetIdAndFieldImageWidth Get Node_fieldImages via EntityIdAndFieldImageTargetIdAndFieldImageWidth
func GetNode_fieldImagesByEntityIdAndFieldImageTargetIdAndFieldImageWidth(offset int, limit int, EntityId_ int, FieldImageTargetId_ int, FieldImageWidth_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("entity_id = ? and field_image_target_id = ? and field_image_width = ?", EntityId_, FieldImageTargetId_, FieldImageWidth_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByEntityIdAndFieldImageTargetIdAndFieldImageHeight Get Node_fieldImages via EntityIdAndFieldImageTargetIdAndFieldImageHeight
func GetNode_fieldImagesByEntityIdAndFieldImageTargetIdAndFieldImageHeight(offset int, limit int, EntityId_ int, FieldImageTargetId_ int, FieldImageHeight_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("entity_id = ? and field_image_target_id = ? and field_image_height = ?", EntityId_, FieldImageTargetId_, FieldImageHeight_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByEntityIdAndFieldImageAltAndFieldImageTitle Get Node_fieldImages via EntityIdAndFieldImageAltAndFieldImageTitle
func GetNode_fieldImagesByEntityIdAndFieldImageAltAndFieldImageTitle(offset int, limit int, EntityId_ int, FieldImageAlt_ string, FieldImageTitle_ string) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("entity_id = ? and field_image_alt = ? and field_image_title = ?", EntityId_, FieldImageAlt_, FieldImageTitle_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByEntityIdAndFieldImageAltAndFieldImageWidth Get Node_fieldImages via EntityIdAndFieldImageAltAndFieldImageWidth
func GetNode_fieldImagesByEntityIdAndFieldImageAltAndFieldImageWidth(offset int, limit int, EntityId_ int, FieldImageAlt_ string, FieldImageWidth_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("entity_id = ? and field_image_alt = ? and field_image_width = ?", EntityId_, FieldImageAlt_, FieldImageWidth_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByEntityIdAndFieldImageAltAndFieldImageHeight Get Node_fieldImages via EntityIdAndFieldImageAltAndFieldImageHeight
func GetNode_fieldImagesByEntityIdAndFieldImageAltAndFieldImageHeight(offset int, limit int, EntityId_ int, FieldImageAlt_ string, FieldImageHeight_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("entity_id = ? and field_image_alt = ? and field_image_height = ?", EntityId_, FieldImageAlt_, FieldImageHeight_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByEntityIdAndFieldImageTitleAndFieldImageWidth Get Node_fieldImages via EntityIdAndFieldImageTitleAndFieldImageWidth
func GetNode_fieldImagesByEntityIdAndFieldImageTitleAndFieldImageWidth(offset int, limit int, EntityId_ int, FieldImageTitle_ string, FieldImageWidth_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("entity_id = ? and field_image_title = ? and field_image_width = ?", EntityId_, FieldImageTitle_, FieldImageWidth_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByEntityIdAndFieldImageTitleAndFieldImageHeight Get Node_fieldImages via EntityIdAndFieldImageTitleAndFieldImageHeight
func GetNode_fieldImagesByEntityIdAndFieldImageTitleAndFieldImageHeight(offset int, limit int, EntityId_ int, FieldImageTitle_ string, FieldImageHeight_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("entity_id = ? and field_image_title = ? and field_image_height = ?", EntityId_, FieldImageTitle_, FieldImageHeight_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByEntityIdAndFieldImageWidthAndFieldImageHeight Get Node_fieldImages via EntityIdAndFieldImageWidthAndFieldImageHeight
func GetNode_fieldImagesByEntityIdAndFieldImageWidthAndFieldImageHeight(offset int, limit int, EntityId_ int, FieldImageWidth_ int, FieldImageHeight_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("entity_id = ? and field_image_width = ? and field_image_height = ?", EntityId_, FieldImageWidth_, FieldImageHeight_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByRevisionIdAndLangcodeAndDelta Get Node_fieldImages via RevisionIdAndLangcodeAndDelta
func GetNode_fieldImagesByRevisionIdAndLangcodeAndDelta(offset int, limit int, RevisionId_ int, Langcode_ string, Delta_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("revision_id = ? and langcode = ? and delta = ?", RevisionId_, Langcode_, Delta_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByRevisionIdAndLangcodeAndFieldImageTargetId Get Node_fieldImages via RevisionIdAndLangcodeAndFieldImageTargetId
func GetNode_fieldImagesByRevisionIdAndLangcodeAndFieldImageTargetId(offset int, limit int, RevisionId_ int, Langcode_ string, FieldImageTargetId_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("revision_id = ? and langcode = ? and field_image_target_id = ?", RevisionId_, Langcode_, FieldImageTargetId_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByRevisionIdAndLangcodeAndFieldImageAlt Get Node_fieldImages via RevisionIdAndLangcodeAndFieldImageAlt
func GetNode_fieldImagesByRevisionIdAndLangcodeAndFieldImageAlt(offset int, limit int, RevisionId_ int, Langcode_ string, FieldImageAlt_ string) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("revision_id = ? and langcode = ? and field_image_alt = ?", RevisionId_, Langcode_, FieldImageAlt_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByRevisionIdAndLangcodeAndFieldImageTitle Get Node_fieldImages via RevisionIdAndLangcodeAndFieldImageTitle
func GetNode_fieldImagesByRevisionIdAndLangcodeAndFieldImageTitle(offset int, limit int, RevisionId_ int, Langcode_ string, FieldImageTitle_ string) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("revision_id = ? and langcode = ? and field_image_title = ?", RevisionId_, Langcode_, FieldImageTitle_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByRevisionIdAndLangcodeAndFieldImageWidth Get Node_fieldImages via RevisionIdAndLangcodeAndFieldImageWidth
func GetNode_fieldImagesByRevisionIdAndLangcodeAndFieldImageWidth(offset int, limit int, RevisionId_ int, Langcode_ string, FieldImageWidth_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("revision_id = ? and langcode = ? and field_image_width = ?", RevisionId_, Langcode_, FieldImageWidth_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByRevisionIdAndLangcodeAndFieldImageHeight Get Node_fieldImages via RevisionIdAndLangcodeAndFieldImageHeight
func GetNode_fieldImagesByRevisionIdAndLangcodeAndFieldImageHeight(offset int, limit int, RevisionId_ int, Langcode_ string, FieldImageHeight_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("revision_id = ? and langcode = ? and field_image_height = ?", RevisionId_, Langcode_, FieldImageHeight_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByRevisionIdAndDeltaAndFieldImageTargetId Get Node_fieldImages via RevisionIdAndDeltaAndFieldImageTargetId
func GetNode_fieldImagesByRevisionIdAndDeltaAndFieldImageTargetId(offset int, limit int, RevisionId_ int, Delta_ int, FieldImageTargetId_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("revision_id = ? and delta = ? and field_image_target_id = ?", RevisionId_, Delta_, FieldImageTargetId_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByRevisionIdAndDeltaAndFieldImageAlt Get Node_fieldImages via RevisionIdAndDeltaAndFieldImageAlt
func GetNode_fieldImagesByRevisionIdAndDeltaAndFieldImageAlt(offset int, limit int, RevisionId_ int, Delta_ int, FieldImageAlt_ string) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("revision_id = ? and delta = ? and field_image_alt = ?", RevisionId_, Delta_, FieldImageAlt_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByRevisionIdAndDeltaAndFieldImageTitle Get Node_fieldImages via RevisionIdAndDeltaAndFieldImageTitle
func GetNode_fieldImagesByRevisionIdAndDeltaAndFieldImageTitle(offset int, limit int, RevisionId_ int, Delta_ int, FieldImageTitle_ string) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("revision_id = ? and delta = ? and field_image_title = ?", RevisionId_, Delta_, FieldImageTitle_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByRevisionIdAndDeltaAndFieldImageWidth Get Node_fieldImages via RevisionIdAndDeltaAndFieldImageWidth
func GetNode_fieldImagesByRevisionIdAndDeltaAndFieldImageWidth(offset int, limit int, RevisionId_ int, Delta_ int, FieldImageWidth_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("revision_id = ? and delta = ? and field_image_width = ?", RevisionId_, Delta_, FieldImageWidth_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByRevisionIdAndDeltaAndFieldImageHeight Get Node_fieldImages via RevisionIdAndDeltaAndFieldImageHeight
func GetNode_fieldImagesByRevisionIdAndDeltaAndFieldImageHeight(offset int, limit int, RevisionId_ int, Delta_ int, FieldImageHeight_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("revision_id = ? and delta = ? and field_image_height = ?", RevisionId_, Delta_, FieldImageHeight_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByRevisionIdAndFieldImageTargetIdAndFieldImageAlt Get Node_fieldImages via RevisionIdAndFieldImageTargetIdAndFieldImageAlt
func GetNode_fieldImagesByRevisionIdAndFieldImageTargetIdAndFieldImageAlt(offset int, limit int, RevisionId_ int, FieldImageTargetId_ int, FieldImageAlt_ string) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("revision_id = ? and field_image_target_id = ? and field_image_alt = ?", RevisionId_, FieldImageTargetId_, FieldImageAlt_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByRevisionIdAndFieldImageTargetIdAndFieldImageTitle Get Node_fieldImages via RevisionIdAndFieldImageTargetIdAndFieldImageTitle
func GetNode_fieldImagesByRevisionIdAndFieldImageTargetIdAndFieldImageTitle(offset int, limit int, RevisionId_ int, FieldImageTargetId_ int, FieldImageTitle_ string) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("revision_id = ? and field_image_target_id = ? and field_image_title = ?", RevisionId_, FieldImageTargetId_, FieldImageTitle_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByRevisionIdAndFieldImageTargetIdAndFieldImageWidth Get Node_fieldImages via RevisionIdAndFieldImageTargetIdAndFieldImageWidth
func GetNode_fieldImagesByRevisionIdAndFieldImageTargetIdAndFieldImageWidth(offset int, limit int, RevisionId_ int, FieldImageTargetId_ int, FieldImageWidth_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("revision_id = ? and field_image_target_id = ? and field_image_width = ?", RevisionId_, FieldImageTargetId_, FieldImageWidth_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByRevisionIdAndFieldImageTargetIdAndFieldImageHeight Get Node_fieldImages via RevisionIdAndFieldImageTargetIdAndFieldImageHeight
func GetNode_fieldImagesByRevisionIdAndFieldImageTargetIdAndFieldImageHeight(offset int, limit int, RevisionId_ int, FieldImageTargetId_ int, FieldImageHeight_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("revision_id = ? and field_image_target_id = ? and field_image_height = ?", RevisionId_, FieldImageTargetId_, FieldImageHeight_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByRevisionIdAndFieldImageAltAndFieldImageTitle Get Node_fieldImages via RevisionIdAndFieldImageAltAndFieldImageTitle
func GetNode_fieldImagesByRevisionIdAndFieldImageAltAndFieldImageTitle(offset int, limit int, RevisionId_ int, FieldImageAlt_ string, FieldImageTitle_ string) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("revision_id = ? and field_image_alt = ? and field_image_title = ?", RevisionId_, FieldImageAlt_, FieldImageTitle_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByRevisionIdAndFieldImageAltAndFieldImageWidth Get Node_fieldImages via RevisionIdAndFieldImageAltAndFieldImageWidth
func GetNode_fieldImagesByRevisionIdAndFieldImageAltAndFieldImageWidth(offset int, limit int, RevisionId_ int, FieldImageAlt_ string, FieldImageWidth_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("revision_id = ? and field_image_alt = ? and field_image_width = ?", RevisionId_, FieldImageAlt_, FieldImageWidth_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByRevisionIdAndFieldImageAltAndFieldImageHeight Get Node_fieldImages via RevisionIdAndFieldImageAltAndFieldImageHeight
func GetNode_fieldImagesByRevisionIdAndFieldImageAltAndFieldImageHeight(offset int, limit int, RevisionId_ int, FieldImageAlt_ string, FieldImageHeight_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("revision_id = ? and field_image_alt = ? and field_image_height = ?", RevisionId_, FieldImageAlt_, FieldImageHeight_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByRevisionIdAndFieldImageTitleAndFieldImageWidth Get Node_fieldImages via RevisionIdAndFieldImageTitleAndFieldImageWidth
func GetNode_fieldImagesByRevisionIdAndFieldImageTitleAndFieldImageWidth(offset int, limit int, RevisionId_ int, FieldImageTitle_ string, FieldImageWidth_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("revision_id = ? and field_image_title = ? and field_image_width = ?", RevisionId_, FieldImageTitle_, FieldImageWidth_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByRevisionIdAndFieldImageTitleAndFieldImageHeight Get Node_fieldImages via RevisionIdAndFieldImageTitleAndFieldImageHeight
func GetNode_fieldImagesByRevisionIdAndFieldImageTitleAndFieldImageHeight(offset int, limit int, RevisionId_ int, FieldImageTitle_ string, FieldImageHeight_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("revision_id = ? and field_image_title = ? and field_image_height = ?", RevisionId_, FieldImageTitle_, FieldImageHeight_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByRevisionIdAndFieldImageWidthAndFieldImageHeight Get Node_fieldImages via RevisionIdAndFieldImageWidthAndFieldImageHeight
func GetNode_fieldImagesByRevisionIdAndFieldImageWidthAndFieldImageHeight(offset int, limit int, RevisionId_ int, FieldImageWidth_ int, FieldImageHeight_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("revision_id = ? and field_image_width = ? and field_image_height = ?", RevisionId_, FieldImageWidth_, FieldImageHeight_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByLangcodeAndDeltaAndFieldImageTargetId Get Node_fieldImages via LangcodeAndDeltaAndFieldImageTargetId
func GetNode_fieldImagesByLangcodeAndDeltaAndFieldImageTargetId(offset int, limit int, Langcode_ string, Delta_ int, FieldImageTargetId_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("langcode = ? and delta = ? and field_image_target_id = ?", Langcode_, Delta_, FieldImageTargetId_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByLangcodeAndDeltaAndFieldImageAlt Get Node_fieldImages via LangcodeAndDeltaAndFieldImageAlt
func GetNode_fieldImagesByLangcodeAndDeltaAndFieldImageAlt(offset int, limit int, Langcode_ string, Delta_ int, FieldImageAlt_ string) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("langcode = ? and delta = ? and field_image_alt = ?", Langcode_, Delta_, FieldImageAlt_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByLangcodeAndDeltaAndFieldImageTitle Get Node_fieldImages via LangcodeAndDeltaAndFieldImageTitle
func GetNode_fieldImagesByLangcodeAndDeltaAndFieldImageTitle(offset int, limit int, Langcode_ string, Delta_ int, FieldImageTitle_ string) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("langcode = ? and delta = ? and field_image_title = ?", Langcode_, Delta_, FieldImageTitle_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByLangcodeAndDeltaAndFieldImageWidth Get Node_fieldImages via LangcodeAndDeltaAndFieldImageWidth
func GetNode_fieldImagesByLangcodeAndDeltaAndFieldImageWidth(offset int, limit int, Langcode_ string, Delta_ int, FieldImageWidth_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("langcode = ? and delta = ? and field_image_width = ?", Langcode_, Delta_, FieldImageWidth_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByLangcodeAndDeltaAndFieldImageHeight Get Node_fieldImages via LangcodeAndDeltaAndFieldImageHeight
func GetNode_fieldImagesByLangcodeAndDeltaAndFieldImageHeight(offset int, limit int, Langcode_ string, Delta_ int, FieldImageHeight_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("langcode = ? and delta = ? and field_image_height = ?", Langcode_, Delta_, FieldImageHeight_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByLangcodeAndFieldImageTargetIdAndFieldImageAlt Get Node_fieldImages via LangcodeAndFieldImageTargetIdAndFieldImageAlt
func GetNode_fieldImagesByLangcodeAndFieldImageTargetIdAndFieldImageAlt(offset int, limit int, Langcode_ string, FieldImageTargetId_ int, FieldImageAlt_ string) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("langcode = ? and field_image_target_id = ? and field_image_alt = ?", Langcode_, FieldImageTargetId_, FieldImageAlt_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByLangcodeAndFieldImageTargetIdAndFieldImageTitle Get Node_fieldImages via LangcodeAndFieldImageTargetIdAndFieldImageTitle
func GetNode_fieldImagesByLangcodeAndFieldImageTargetIdAndFieldImageTitle(offset int, limit int, Langcode_ string, FieldImageTargetId_ int, FieldImageTitle_ string) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("langcode = ? and field_image_target_id = ? and field_image_title = ?", Langcode_, FieldImageTargetId_, FieldImageTitle_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByLangcodeAndFieldImageTargetIdAndFieldImageWidth Get Node_fieldImages via LangcodeAndFieldImageTargetIdAndFieldImageWidth
func GetNode_fieldImagesByLangcodeAndFieldImageTargetIdAndFieldImageWidth(offset int, limit int, Langcode_ string, FieldImageTargetId_ int, FieldImageWidth_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("langcode = ? and field_image_target_id = ? and field_image_width = ?", Langcode_, FieldImageTargetId_, FieldImageWidth_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByLangcodeAndFieldImageTargetIdAndFieldImageHeight Get Node_fieldImages via LangcodeAndFieldImageTargetIdAndFieldImageHeight
func GetNode_fieldImagesByLangcodeAndFieldImageTargetIdAndFieldImageHeight(offset int, limit int, Langcode_ string, FieldImageTargetId_ int, FieldImageHeight_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("langcode = ? and field_image_target_id = ? and field_image_height = ?", Langcode_, FieldImageTargetId_, FieldImageHeight_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByLangcodeAndFieldImageAltAndFieldImageTitle Get Node_fieldImages via LangcodeAndFieldImageAltAndFieldImageTitle
func GetNode_fieldImagesByLangcodeAndFieldImageAltAndFieldImageTitle(offset int, limit int, Langcode_ string, FieldImageAlt_ string, FieldImageTitle_ string) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("langcode = ? and field_image_alt = ? and field_image_title = ?", Langcode_, FieldImageAlt_, FieldImageTitle_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByLangcodeAndFieldImageAltAndFieldImageWidth Get Node_fieldImages via LangcodeAndFieldImageAltAndFieldImageWidth
func GetNode_fieldImagesByLangcodeAndFieldImageAltAndFieldImageWidth(offset int, limit int, Langcode_ string, FieldImageAlt_ string, FieldImageWidth_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("langcode = ? and field_image_alt = ? and field_image_width = ?", Langcode_, FieldImageAlt_, FieldImageWidth_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByLangcodeAndFieldImageAltAndFieldImageHeight Get Node_fieldImages via LangcodeAndFieldImageAltAndFieldImageHeight
func GetNode_fieldImagesByLangcodeAndFieldImageAltAndFieldImageHeight(offset int, limit int, Langcode_ string, FieldImageAlt_ string, FieldImageHeight_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("langcode = ? and field_image_alt = ? and field_image_height = ?", Langcode_, FieldImageAlt_, FieldImageHeight_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByLangcodeAndFieldImageTitleAndFieldImageWidth Get Node_fieldImages via LangcodeAndFieldImageTitleAndFieldImageWidth
func GetNode_fieldImagesByLangcodeAndFieldImageTitleAndFieldImageWidth(offset int, limit int, Langcode_ string, FieldImageTitle_ string, FieldImageWidth_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("langcode = ? and field_image_title = ? and field_image_width = ?", Langcode_, FieldImageTitle_, FieldImageWidth_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByLangcodeAndFieldImageTitleAndFieldImageHeight Get Node_fieldImages via LangcodeAndFieldImageTitleAndFieldImageHeight
func GetNode_fieldImagesByLangcodeAndFieldImageTitleAndFieldImageHeight(offset int, limit int, Langcode_ string, FieldImageTitle_ string, FieldImageHeight_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("langcode = ? and field_image_title = ? and field_image_height = ?", Langcode_, FieldImageTitle_, FieldImageHeight_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByLangcodeAndFieldImageWidthAndFieldImageHeight Get Node_fieldImages via LangcodeAndFieldImageWidthAndFieldImageHeight
func GetNode_fieldImagesByLangcodeAndFieldImageWidthAndFieldImageHeight(offset int, limit int, Langcode_ string, FieldImageWidth_ int, FieldImageHeight_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("langcode = ? and field_image_width = ? and field_image_height = ?", Langcode_, FieldImageWidth_, FieldImageHeight_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByDeltaAndFieldImageTargetIdAndFieldImageAlt Get Node_fieldImages via DeltaAndFieldImageTargetIdAndFieldImageAlt
func GetNode_fieldImagesByDeltaAndFieldImageTargetIdAndFieldImageAlt(offset int, limit int, Delta_ int, FieldImageTargetId_ int, FieldImageAlt_ string) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("delta = ? and field_image_target_id = ? and field_image_alt = ?", Delta_, FieldImageTargetId_, FieldImageAlt_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByDeltaAndFieldImageTargetIdAndFieldImageTitle Get Node_fieldImages via DeltaAndFieldImageTargetIdAndFieldImageTitle
func GetNode_fieldImagesByDeltaAndFieldImageTargetIdAndFieldImageTitle(offset int, limit int, Delta_ int, FieldImageTargetId_ int, FieldImageTitle_ string) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("delta = ? and field_image_target_id = ? and field_image_title = ?", Delta_, FieldImageTargetId_, FieldImageTitle_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByDeltaAndFieldImageTargetIdAndFieldImageWidth Get Node_fieldImages via DeltaAndFieldImageTargetIdAndFieldImageWidth
func GetNode_fieldImagesByDeltaAndFieldImageTargetIdAndFieldImageWidth(offset int, limit int, Delta_ int, FieldImageTargetId_ int, FieldImageWidth_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("delta = ? and field_image_target_id = ? and field_image_width = ?", Delta_, FieldImageTargetId_, FieldImageWidth_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByDeltaAndFieldImageTargetIdAndFieldImageHeight Get Node_fieldImages via DeltaAndFieldImageTargetIdAndFieldImageHeight
func GetNode_fieldImagesByDeltaAndFieldImageTargetIdAndFieldImageHeight(offset int, limit int, Delta_ int, FieldImageTargetId_ int, FieldImageHeight_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("delta = ? and field_image_target_id = ? and field_image_height = ?", Delta_, FieldImageTargetId_, FieldImageHeight_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByDeltaAndFieldImageAltAndFieldImageTitle Get Node_fieldImages via DeltaAndFieldImageAltAndFieldImageTitle
func GetNode_fieldImagesByDeltaAndFieldImageAltAndFieldImageTitle(offset int, limit int, Delta_ int, FieldImageAlt_ string, FieldImageTitle_ string) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("delta = ? and field_image_alt = ? and field_image_title = ?", Delta_, FieldImageAlt_, FieldImageTitle_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByDeltaAndFieldImageAltAndFieldImageWidth Get Node_fieldImages via DeltaAndFieldImageAltAndFieldImageWidth
func GetNode_fieldImagesByDeltaAndFieldImageAltAndFieldImageWidth(offset int, limit int, Delta_ int, FieldImageAlt_ string, FieldImageWidth_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("delta = ? and field_image_alt = ? and field_image_width = ?", Delta_, FieldImageAlt_, FieldImageWidth_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByDeltaAndFieldImageAltAndFieldImageHeight Get Node_fieldImages via DeltaAndFieldImageAltAndFieldImageHeight
func GetNode_fieldImagesByDeltaAndFieldImageAltAndFieldImageHeight(offset int, limit int, Delta_ int, FieldImageAlt_ string, FieldImageHeight_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("delta = ? and field_image_alt = ? and field_image_height = ?", Delta_, FieldImageAlt_, FieldImageHeight_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByDeltaAndFieldImageTitleAndFieldImageWidth Get Node_fieldImages via DeltaAndFieldImageTitleAndFieldImageWidth
func GetNode_fieldImagesByDeltaAndFieldImageTitleAndFieldImageWidth(offset int, limit int, Delta_ int, FieldImageTitle_ string, FieldImageWidth_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("delta = ? and field_image_title = ? and field_image_width = ?", Delta_, FieldImageTitle_, FieldImageWidth_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByDeltaAndFieldImageTitleAndFieldImageHeight Get Node_fieldImages via DeltaAndFieldImageTitleAndFieldImageHeight
func GetNode_fieldImagesByDeltaAndFieldImageTitleAndFieldImageHeight(offset int, limit int, Delta_ int, FieldImageTitle_ string, FieldImageHeight_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("delta = ? and field_image_title = ? and field_image_height = ?", Delta_, FieldImageTitle_, FieldImageHeight_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByDeltaAndFieldImageWidthAndFieldImageHeight Get Node_fieldImages via DeltaAndFieldImageWidthAndFieldImageHeight
func GetNode_fieldImagesByDeltaAndFieldImageWidthAndFieldImageHeight(offset int, limit int, Delta_ int, FieldImageWidth_ int, FieldImageHeight_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("delta = ? and field_image_width = ? and field_image_height = ?", Delta_, FieldImageWidth_, FieldImageHeight_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByFieldImageTargetIdAndFieldImageAltAndFieldImageTitle Get Node_fieldImages via FieldImageTargetIdAndFieldImageAltAndFieldImageTitle
func GetNode_fieldImagesByFieldImageTargetIdAndFieldImageAltAndFieldImageTitle(offset int, limit int, FieldImageTargetId_ int, FieldImageAlt_ string, FieldImageTitle_ string) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("field_image_target_id = ? and field_image_alt = ? and field_image_title = ?", FieldImageTargetId_, FieldImageAlt_, FieldImageTitle_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByFieldImageTargetIdAndFieldImageAltAndFieldImageWidth Get Node_fieldImages via FieldImageTargetIdAndFieldImageAltAndFieldImageWidth
func GetNode_fieldImagesByFieldImageTargetIdAndFieldImageAltAndFieldImageWidth(offset int, limit int, FieldImageTargetId_ int, FieldImageAlt_ string, FieldImageWidth_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("field_image_target_id = ? and field_image_alt = ? and field_image_width = ?", FieldImageTargetId_, FieldImageAlt_, FieldImageWidth_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByFieldImageTargetIdAndFieldImageAltAndFieldImageHeight Get Node_fieldImages via FieldImageTargetIdAndFieldImageAltAndFieldImageHeight
func GetNode_fieldImagesByFieldImageTargetIdAndFieldImageAltAndFieldImageHeight(offset int, limit int, FieldImageTargetId_ int, FieldImageAlt_ string, FieldImageHeight_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("field_image_target_id = ? and field_image_alt = ? and field_image_height = ?", FieldImageTargetId_, FieldImageAlt_, FieldImageHeight_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByFieldImageTargetIdAndFieldImageTitleAndFieldImageWidth Get Node_fieldImages via FieldImageTargetIdAndFieldImageTitleAndFieldImageWidth
func GetNode_fieldImagesByFieldImageTargetIdAndFieldImageTitleAndFieldImageWidth(offset int, limit int, FieldImageTargetId_ int, FieldImageTitle_ string, FieldImageWidth_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("field_image_target_id = ? and field_image_title = ? and field_image_width = ?", FieldImageTargetId_, FieldImageTitle_, FieldImageWidth_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByFieldImageTargetIdAndFieldImageTitleAndFieldImageHeight Get Node_fieldImages via FieldImageTargetIdAndFieldImageTitleAndFieldImageHeight
func GetNode_fieldImagesByFieldImageTargetIdAndFieldImageTitleAndFieldImageHeight(offset int, limit int, FieldImageTargetId_ int, FieldImageTitle_ string, FieldImageHeight_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("field_image_target_id = ? and field_image_title = ? and field_image_height = ?", FieldImageTargetId_, FieldImageTitle_, FieldImageHeight_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByFieldImageTargetIdAndFieldImageWidthAndFieldImageHeight Get Node_fieldImages via FieldImageTargetIdAndFieldImageWidthAndFieldImageHeight
func GetNode_fieldImagesByFieldImageTargetIdAndFieldImageWidthAndFieldImageHeight(offset int, limit int, FieldImageTargetId_ int, FieldImageWidth_ int, FieldImageHeight_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("field_image_target_id = ? and field_image_width = ? and field_image_height = ?", FieldImageTargetId_, FieldImageWidth_, FieldImageHeight_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByFieldImageAltAndFieldImageTitleAndFieldImageWidth Get Node_fieldImages via FieldImageAltAndFieldImageTitleAndFieldImageWidth
func GetNode_fieldImagesByFieldImageAltAndFieldImageTitleAndFieldImageWidth(offset int, limit int, FieldImageAlt_ string, FieldImageTitle_ string, FieldImageWidth_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("field_image_alt = ? and field_image_title = ? and field_image_width = ?", FieldImageAlt_, FieldImageTitle_, FieldImageWidth_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByFieldImageAltAndFieldImageTitleAndFieldImageHeight Get Node_fieldImages via FieldImageAltAndFieldImageTitleAndFieldImageHeight
func GetNode_fieldImagesByFieldImageAltAndFieldImageTitleAndFieldImageHeight(offset int, limit int, FieldImageAlt_ string, FieldImageTitle_ string, FieldImageHeight_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("field_image_alt = ? and field_image_title = ? and field_image_height = ?", FieldImageAlt_, FieldImageTitle_, FieldImageHeight_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByFieldImageAltAndFieldImageWidthAndFieldImageHeight Get Node_fieldImages via FieldImageAltAndFieldImageWidthAndFieldImageHeight
func GetNode_fieldImagesByFieldImageAltAndFieldImageWidthAndFieldImageHeight(offset int, limit int, FieldImageAlt_ string, FieldImageWidth_ int, FieldImageHeight_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("field_image_alt = ? and field_image_width = ? and field_image_height = ?", FieldImageAlt_, FieldImageWidth_, FieldImageHeight_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByFieldImageTitleAndFieldImageWidthAndFieldImageHeight Get Node_fieldImages via FieldImageTitleAndFieldImageWidthAndFieldImageHeight
func GetNode_fieldImagesByFieldImageTitleAndFieldImageWidthAndFieldImageHeight(offset int, limit int, FieldImageTitle_ string, FieldImageWidth_ int, FieldImageHeight_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("field_image_title = ? and field_image_width = ? and field_image_height = ?", FieldImageTitle_, FieldImageWidth_, FieldImageHeight_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByBundleAndDeleted Get Node_fieldImages via BundleAndDeleted
func GetNode_fieldImagesByBundleAndDeleted(offset int, limit int, Bundle_ string, Deleted_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("bundle = ? and deleted = ?", Bundle_, Deleted_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByBundleAndEntityId Get Node_fieldImages via BundleAndEntityId
func GetNode_fieldImagesByBundleAndEntityId(offset int, limit int, Bundle_ string, EntityId_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("bundle = ? and entity_id = ?", Bundle_, EntityId_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByBundleAndRevisionId Get Node_fieldImages via BundleAndRevisionId
func GetNode_fieldImagesByBundleAndRevisionId(offset int, limit int, Bundle_ string, RevisionId_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("bundle = ? and revision_id = ?", Bundle_, RevisionId_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByBundleAndLangcode Get Node_fieldImages via BundleAndLangcode
func GetNode_fieldImagesByBundleAndLangcode(offset int, limit int, Bundle_ string, Langcode_ string) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("bundle = ? and langcode = ?", Bundle_, Langcode_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByBundleAndDelta Get Node_fieldImages via BundleAndDelta
func GetNode_fieldImagesByBundleAndDelta(offset int, limit int, Bundle_ string, Delta_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("bundle = ? and delta = ?", Bundle_, Delta_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByBundleAndFieldImageTargetId Get Node_fieldImages via BundleAndFieldImageTargetId
func GetNode_fieldImagesByBundleAndFieldImageTargetId(offset int, limit int, Bundle_ string, FieldImageTargetId_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("bundle = ? and field_image_target_id = ?", Bundle_, FieldImageTargetId_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByBundleAndFieldImageAlt Get Node_fieldImages via BundleAndFieldImageAlt
func GetNode_fieldImagesByBundleAndFieldImageAlt(offset int, limit int, Bundle_ string, FieldImageAlt_ string) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("bundle = ? and field_image_alt = ?", Bundle_, FieldImageAlt_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByBundleAndFieldImageTitle Get Node_fieldImages via BundleAndFieldImageTitle
func GetNode_fieldImagesByBundleAndFieldImageTitle(offset int, limit int, Bundle_ string, FieldImageTitle_ string) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("bundle = ? and field_image_title = ?", Bundle_, FieldImageTitle_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByBundleAndFieldImageWidth Get Node_fieldImages via BundleAndFieldImageWidth
func GetNode_fieldImagesByBundleAndFieldImageWidth(offset int, limit int, Bundle_ string, FieldImageWidth_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("bundle = ? and field_image_width = ?", Bundle_, FieldImageWidth_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByBundleAndFieldImageHeight Get Node_fieldImages via BundleAndFieldImageHeight
func GetNode_fieldImagesByBundleAndFieldImageHeight(offset int, limit int, Bundle_ string, FieldImageHeight_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("bundle = ? and field_image_height = ?", Bundle_, FieldImageHeight_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByDeletedAndEntityId Get Node_fieldImages via DeletedAndEntityId
func GetNode_fieldImagesByDeletedAndEntityId(offset int, limit int, Deleted_ int, EntityId_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("deleted = ? and entity_id = ?", Deleted_, EntityId_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByDeletedAndRevisionId Get Node_fieldImages via DeletedAndRevisionId
func GetNode_fieldImagesByDeletedAndRevisionId(offset int, limit int, Deleted_ int, RevisionId_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("deleted = ? and revision_id = ?", Deleted_, RevisionId_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByDeletedAndLangcode Get Node_fieldImages via DeletedAndLangcode
func GetNode_fieldImagesByDeletedAndLangcode(offset int, limit int, Deleted_ int, Langcode_ string) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("deleted = ? and langcode = ?", Deleted_, Langcode_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByDeletedAndDelta Get Node_fieldImages via DeletedAndDelta
func GetNode_fieldImagesByDeletedAndDelta(offset int, limit int, Deleted_ int, Delta_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("deleted = ? and delta = ?", Deleted_, Delta_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByDeletedAndFieldImageTargetId Get Node_fieldImages via DeletedAndFieldImageTargetId
func GetNode_fieldImagesByDeletedAndFieldImageTargetId(offset int, limit int, Deleted_ int, FieldImageTargetId_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("deleted = ? and field_image_target_id = ?", Deleted_, FieldImageTargetId_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByDeletedAndFieldImageAlt Get Node_fieldImages via DeletedAndFieldImageAlt
func GetNode_fieldImagesByDeletedAndFieldImageAlt(offset int, limit int, Deleted_ int, FieldImageAlt_ string) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("deleted = ? and field_image_alt = ?", Deleted_, FieldImageAlt_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByDeletedAndFieldImageTitle Get Node_fieldImages via DeletedAndFieldImageTitle
func GetNode_fieldImagesByDeletedAndFieldImageTitle(offset int, limit int, Deleted_ int, FieldImageTitle_ string) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("deleted = ? and field_image_title = ?", Deleted_, FieldImageTitle_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByDeletedAndFieldImageWidth Get Node_fieldImages via DeletedAndFieldImageWidth
func GetNode_fieldImagesByDeletedAndFieldImageWidth(offset int, limit int, Deleted_ int, FieldImageWidth_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("deleted = ? and field_image_width = ?", Deleted_, FieldImageWidth_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByDeletedAndFieldImageHeight Get Node_fieldImages via DeletedAndFieldImageHeight
func GetNode_fieldImagesByDeletedAndFieldImageHeight(offset int, limit int, Deleted_ int, FieldImageHeight_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("deleted = ? and field_image_height = ?", Deleted_, FieldImageHeight_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByEntityIdAndRevisionId Get Node_fieldImages via EntityIdAndRevisionId
func GetNode_fieldImagesByEntityIdAndRevisionId(offset int, limit int, EntityId_ int, RevisionId_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("entity_id = ? and revision_id = ?", EntityId_, RevisionId_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByEntityIdAndLangcode Get Node_fieldImages via EntityIdAndLangcode
func GetNode_fieldImagesByEntityIdAndLangcode(offset int, limit int, EntityId_ int, Langcode_ string) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("entity_id = ? and langcode = ?", EntityId_, Langcode_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByEntityIdAndDelta Get Node_fieldImages via EntityIdAndDelta
func GetNode_fieldImagesByEntityIdAndDelta(offset int, limit int, EntityId_ int, Delta_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("entity_id = ? and delta = ?", EntityId_, Delta_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByEntityIdAndFieldImageTargetId Get Node_fieldImages via EntityIdAndFieldImageTargetId
func GetNode_fieldImagesByEntityIdAndFieldImageTargetId(offset int, limit int, EntityId_ int, FieldImageTargetId_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("entity_id = ? and field_image_target_id = ?", EntityId_, FieldImageTargetId_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByEntityIdAndFieldImageAlt Get Node_fieldImages via EntityIdAndFieldImageAlt
func GetNode_fieldImagesByEntityIdAndFieldImageAlt(offset int, limit int, EntityId_ int, FieldImageAlt_ string) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("entity_id = ? and field_image_alt = ?", EntityId_, FieldImageAlt_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByEntityIdAndFieldImageTitle Get Node_fieldImages via EntityIdAndFieldImageTitle
func GetNode_fieldImagesByEntityIdAndFieldImageTitle(offset int, limit int, EntityId_ int, FieldImageTitle_ string) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("entity_id = ? and field_image_title = ?", EntityId_, FieldImageTitle_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByEntityIdAndFieldImageWidth Get Node_fieldImages via EntityIdAndFieldImageWidth
func GetNode_fieldImagesByEntityIdAndFieldImageWidth(offset int, limit int, EntityId_ int, FieldImageWidth_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("entity_id = ? and field_image_width = ?", EntityId_, FieldImageWidth_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByEntityIdAndFieldImageHeight Get Node_fieldImages via EntityIdAndFieldImageHeight
func GetNode_fieldImagesByEntityIdAndFieldImageHeight(offset int, limit int, EntityId_ int, FieldImageHeight_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("entity_id = ? and field_image_height = ?", EntityId_, FieldImageHeight_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByRevisionIdAndLangcode Get Node_fieldImages via RevisionIdAndLangcode
func GetNode_fieldImagesByRevisionIdAndLangcode(offset int, limit int, RevisionId_ int, Langcode_ string) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("revision_id = ? and langcode = ?", RevisionId_, Langcode_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByRevisionIdAndDelta Get Node_fieldImages via RevisionIdAndDelta
func GetNode_fieldImagesByRevisionIdAndDelta(offset int, limit int, RevisionId_ int, Delta_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("revision_id = ? and delta = ?", RevisionId_, Delta_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByRevisionIdAndFieldImageTargetId Get Node_fieldImages via RevisionIdAndFieldImageTargetId
func GetNode_fieldImagesByRevisionIdAndFieldImageTargetId(offset int, limit int, RevisionId_ int, FieldImageTargetId_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("revision_id = ? and field_image_target_id = ?", RevisionId_, FieldImageTargetId_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByRevisionIdAndFieldImageAlt Get Node_fieldImages via RevisionIdAndFieldImageAlt
func GetNode_fieldImagesByRevisionIdAndFieldImageAlt(offset int, limit int, RevisionId_ int, FieldImageAlt_ string) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("revision_id = ? and field_image_alt = ?", RevisionId_, FieldImageAlt_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByRevisionIdAndFieldImageTitle Get Node_fieldImages via RevisionIdAndFieldImageTitle
func GetNode_fieldImagesByRevisionIdAndFieldImageTitle(offset int, limit int, RevisionId_ int, FieldImageTitle_ string) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("revision_id = ? and field_image_title = ?", RevisionId_, FieldImageTitle_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByRevisionIdAndFieldImageWidth Get Node_fieldImages via RevisionIdAndFieldImageWidth
func GetNode_fieldImagesByRevisionIdAndFieldImageWidth(offset int, limit int, RevisionId_ int, FieldImageWidth_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("revision_id = ? and field_image_width = ?", RevisionId_, FieldImageWidth_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByRevisionIdAndFieldImageHeight Get Node_fieldImages via RevisionIdAndFieldImageHeight
func GetNode_fieldImagesByRevisionIdAndFieldImageHeight(offset int, limit int, RevisionId_ int, FieldImageHeight_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("revision_id = ? and field_image_height = ?", RevisionId_, FieldImageHeight_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByLangcodeAndDelta Get Node_fieldImages via LangcodeAndDelta
func GetNode_fieldImagesByLangcodeAndDelta(offset int, limit int, Langcode_ string, Delta_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("langcode = ? and delta = ?", Langcode_, Delta_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByLangcodeAndFieldImageTargetId Get Node_fieldImages via LangcodeAndFieldImageTargetId
func GetNode_fieldImagesByLangcodeAndFieldImageTargetId(offset int, limit int, Langcode_ string, FieldImageTargetId_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("langcode = ? and field_image_target_id = ?", Langcode_, FieldImageTargetId_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByLangcodeAndFieldImageAlt Get Node_fieldImages via LangcodeAndFieldImageAlt
func GetNode_fieldImagesByLangcodeAndFieldImageAlt(offset int, limit int, Langcode_ string, FieldImageAlt_ string) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("langcode = ? and field_image_alt = ?", Langcode_, FieldImageAlt_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByLangcodeAndFieldImageTitle Get Node_fieldImages via LangcodeAndFieldImageTitle
func GetNode_fieldImagesByLangcodeAndFieldImageTitle(offset int, limit int, Langcode_ string, FieldImageTitle_ string) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("langcode = ? and field_image_title = ?", Langcode_, FieldImageTitle_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByLangcodeAndFieldImageWidth Get Node_fieldImages via LangcodeAndFieldImageWidth
func GetNode_fieldImagesByLangcodeAndFieldImageWidth(offset int, limit int, Langcode_ string, FieldImageWidth_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("langcode = ? and field_image_width = ?", Langcode_, FieldImageWidth_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByLangcodeAndFieldImageHeight Get Node_fieldImages via LangcodeAndFieldImageHeight
func GetNode_fieldImagesByLangcodeAndFieldImageHeight(offset int, limit int, Langcode_ string, FieldImageHeight_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("langcode = ? and field_image_height = ?", Langcode_, FieldImageHeight_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByDeltaAndFieldImageTargetId Get Node_fieldImages via DeltaAndFieldImageTargetId
func GetNode_fieldImagesByDeltaAndFieldImageTargetId(offset int, limit int, Delta_ int, FieldImageTargetId_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("delta = ? and field_image_target_id = ?", Delta_, FieldImageTargetId_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByDeltaAndFieldImageAlt Get Node_fieldImages via DeltaAndFieldImageAlt
func GetNode_fieldImagesByDeltaAndFieldImageAlt(offset int, limit int, Delta_ int, FieldImageAlt_ string) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("delta = ? and field_image_alt = ?", Delta_, FieldImageAlt_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByDeltaAndFieldImageTitle Get Node_fieldImages via DeltaAndFieldImageTitle
func GetNode_fieldImagesByDeltaAndFieldImageTitle(offset int, limit int, Delta_ int, FieldImageTitle_ string) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("delta = ? and field_image_title = ?", Delta_, FieldImageTitle_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByDeltaAndFieldImageWidth Get Node_fieldImages via DeltaAndFieldImageWidth
func GetNode_fieldImagesByDeltaAndFieldImageWidth(offset int, limit int, Delta_ int, FieldImageWidth_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("delta = ? and field_image_width = ?", Delta_, FieldImageWidth_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByDeltaAndFieldImageHeight Get Node_fieldImages via DeltaAndFieldImageHeight
func GetNode_fieldImagesByDeltaAndFieldImageHeight(offset int, limit int, Delta_ int, FieldImageHeight_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("delta = ? and field_image_height = ?", Delta_, FieldImageHeight_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByFieldImageTargetIdAndFieldImageAlt Get Node_fieldImages via FieldImageTargetIdAndFieldImageAlt
func GetNode_fieldImagesByFieldImageTargetIdAndFieldImageAlt(offset int, limit int, FieldImageTargetId_ int, FieldImageAlt_ string) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("field_image_target_id = ? and field_image_alt = ?", FieldImageTargetId_, FieldImageAlt_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByFieldImageTargetIdAndFieldImageTitle Get Node_fieldImages via FieldImageTargetIdAndFieldImageTitle
func GetNode_fieldImagesByFieldImageTargetIdAndFieldImageTitle(offset int, limit int, FieldImageTargetId_ int, FieldImageTitle_ string) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("field_image_target_id = ? and field_image_title = ?", FieldImageTargetId_, FieldImageTitle_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByFieldImageTargetIdAndFieldImageWidth Get Node_fieldImages via FieldImageTargetIdAndFieldImageWidth
func GetNode_fieldImagesByFieldImageTargetIdAndFieldImageWidth(offset int, limit int, FieldImageTargetId_ int, FieldImageWidth_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("field_image_target_id = ? and field_image_width = ?", FieldImageTargetId_, FieldImageWidth_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByFieldImageTargetIdAndFieldImageHeight Get Node_fieldImages via FieldImageTargetIdAndFieldImageHeight
func GetNode_fieldImagesByFieldImageTargetIdAndFieldImageHeight(offset int, limit int, FieldImageTargetId_ int, FieldImageHeight_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("field_image_target_id = ? and field_image_height = ?", FieldImageTargetId_, FieldImageHeight_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByFieldImageAltAndFieldImageTitle Get Node_fieldImages via FieldImageAltAndFieldImageTitle
func GetNode_fieldImagesByFieldImageAltAndFieldImageTitle(offset int, limit int, FieldImageAlt_ string, FieldImageTitle_ string) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("field_image_alt = ? and field_image_title = ?", FieldImageAlt_, FieldImageTitle_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByFieldImageAltAndFieldImageWidth Get Node_fieldImages via FieldImageAltAndFieldImageWidth
func GetNode_fieldImagesByFieldImageAltAndFieldImageWidth(offset int, limit int, FieldImageAlt_ string, FieldImageWidth_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("field_image_alt = ? and field_image_width = ?", FieldImageAlt_, FieldImageWidth_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByFieldImageAltAndFieldImageHeight Get Node_fieldImages via FieldImageAltAndFieldImageHeight
func GetNode_fieldImagesByFieldImageAltAndFieldImageHeight(offset int, limit int, FieldImageAlt_ string, FieldImageHeight_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("field_image_alt = ? and field_image_height = ?", FieldImageAlt_, FieldImageHeight_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByFieldImageTitleAndFieldImageWidth Get Node_fieldImages via FieldImageTitleAndFieldImageWidth
func GetNode_fieldImagesByFieldImageTitleAndFieldImageWidth(offset int, limit int, FieldImageTitle_ string, FieldImageWidth_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("field_image_title = ? and field_image_width = ?", FieldImageTitle_, FieldImageWidth_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByFieldImageTitleAndFieldImageHeight Get Node_fieldImages via FieldImageTitleAndFieldImageHeight
func GetNode_fieldImagesByFieldImageTitleAndFieldImageHeight(offset int, limit int, FieldImageTitle_ string, FieldImageHeight_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("field_image_title = ? and field_image_height = ?", FieldImageTitle_, FieldImageHeight_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesByFieldImageWidthAndFieldImageHeight Get Node_fieldImages via FieldImageWidthAndFieldImageHeight
func GetNode_fieldImagesByFieldImageWidthAndFieldImageHeight(offset int, limit int, FieldImageWidth_ int, FieldImageHeight_ int) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("field_image_width = ? and field_image_height = ?", FieldImageWidth_, FieldImageHeight_).Limit(limit, offset).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImages Get Node_fieldImages via field
func GetNode_fieldImages(offset int, limit int, field string) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Limit(limit, offset).Desc(field).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesViaBundle Get Node_fieldImages via Bundle
func GetNode_fieldImagesViaBundle(offset int, limit int, Bundle_ string, field string) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("bundle = ?", Bundle_).Limit(limit, offset).Desc(field).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesViaDeleted Get Node_fieldImages via Deleted
func GetNode_fieldImagesViaDeleted(offset int, limit int, Deleted_ int, field string) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("deleted = ?", Deleted_).Limit(limit, offset).Desc(field).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesViaEntityId Get Node_fieldImages via EntityId
func GetNode_fieldImagesViaEntityId(offset int, limit int, EntityId_ int, field string) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("entity_id = ?", EntityId_).Limit(limit, offset).Desc(field).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesViaRevisionId Get Node_fieldImages via RevisionId
func GetNode_fieldImagesViaRevisionId(offset int, limit int, RevisionId_ int, field string) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("revision_id = ?", RevisionId_).Limit(limit, offset).Desc(field).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesViaLangcode Get Node_fieldImages via Langcode
func GetNode_fieldImagesViaLangcode(offset int, limit int, Langcode_ string, field string) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("langcode = ?", Langcode_).Limit(limit, offset).Desc(field).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesViaDelta Get Node_fieldImages via Delta
func GetNode_fieldImagesViaDelta(offset int, limit int, Delta_ int, field string) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("delta = ?", Delta_).Limit(limit, offset).Desc(field).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesViaFieldImageTargetId Get Node_fieldImages via FieldImageTargetId
func GetNode_fieldImagesViaFieldImageTargetId(offset int, limit int, FieldImageTargetId_ int, field string) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("field_image_target_id = ?", FieldImageTargetId_).Limit(limit, offset).Desc(field).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesViaFieldImageAlt Get Node_fieldImages via FieldImageAlt
func GetNode_fieldImagesViaFieldImageAlt(offset int, limit int, FieldImageAlt_ string, field string) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("field_image_alt = ?", FieldImageAlt_).Limit(limit, offset).Desc(field).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesViaFieldImageTitle Get Node_fieldImages via FieldImageTitle
func GetNode_fieldImagesViaFieldImageTitle(offset int, limit int, FieldImageTitle_ string, field string) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("field_image_title = ?", FieldImageTitle_).Limit(limit, offset).Desc(field).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesViaFieldImageWidth Get Node_fieldImages via FieldImageWidth
func GetNode_fieldImagesViaFieldImageWidth(offset int, limit int, FieldImageWidth_ int, field string) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("field_image_width = ?", FieldImageWidth_).Limit(limit, offset).Desc(field).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// GetNode_fieldImagesViaFieldImageHeight Get Node_fieldImages via FieldImageHeight
func GetNode_fieldImagesViaFieldImageHeight(offset int, limit int, FieldImageHeight_ int, field string) (*[]*Node_fieldImage, error) {
	var _Node_fieldImage = new([]*Node_fieldImage)
	err := Engine.Table("node__field_image").Where("field_image_height = ?", FieldImageHeight_).Limit(limit, offset).Desc(field).Find(_Node_fieldImage)
	return _Node_fieldImage, err
}

// HasNode_fieldImageViaBundle Has Node_fieldImage via Bundle
func HasNode_fieldImageViaBundle(iBundle string) bool {
	if has, err := Engine.Where("bundle = ?", iBundle).Get(new(Node_fieldImage)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasNode_fieldImageViaDeleted Has Node_fieldImage via Deleted
func HasNode_fieldImageViaDeleted(iDeleted int) bool {
	if has, err := Engine.Where("deleted = ?", iDeleted).Get(new(Node_fieldImage)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasNode_fieldImageViaEntityId Has Node_fieldImage via EntityId
func HasNode_fieldImageViaEntityId(iEntityId int) bool {
	if has, err := Engine.Where("entity_id = ?", iEntityId).Get(new(Node_fieldImage)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasNode_fieldImageViaRevisionId Has Node_fieldImage via RevisionId
func HasNode_fieldImageViaRevisionId(iRevisionId int) bool {
	if has, err := Engine.Where("revision_id = ?", iRevisionId).Get(new(Node_fieldImage)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasNode_fieldImageViaLangcode Has Node_fieldImage via Langcode
func HasNode_fieldImageViaLangcode(iLangcode string) bool {
	if has, err := Engine.Where("langcode = ?", iLangcode).Get(new(Node_fieldImage)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasNode_fieldImageViaDelta Has Node_fieldImage via Delta
func HasNode_fieldImageViaDelta(iDelta int) bool {
	if has, err := Engine.Where("delta = ?", iDelta).Get(new(Node_fieldImage)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasNode_fieldImageViaFieldImageTargetId Has Node_fieldImage via FieldImageTargetId
func HasNode_fieldImageViaFieldImageTargetId(iFieldImageTargetId int) bool {
	if has, err := Engine.Where("field_image_target_id = ?", iFieldImageTargetId).Get(new(Node_fieldImage)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasNode_fieldImageViaFieldImageAlt Has Node_fieldImage via FieldImageAlt
func HasNode_fieldImageViaFieldImageAlt(iFieldImageAlt string) bool {
	if has, err := Engine.Where("field_image_alt = ?", iFieldImageAlt).Get(new(Node_fieldImage)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasNode_fieldImageViaFieldImageTitle Has Node_fieldImage via FieldImageTitle
func HasNode_fieldImageViaFieldImageTitle(iFieldImageTitle string) bool {
	if has, err := Engine.Where("field_image_title = ?", iFieldImageTitle).Get(new(Node_fieldImage)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasNode_fieldImageViaFieldImageWidth Has Node_fieldImage via FieldImageWidth
func HasNode_fieldImageViaFieldImageWidth(iFieldImageWidth int) bool {
	if has, err := Engine.Where("field_image_width = ?", iFieldImageWidth).Get(new(Node_fieldImage)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasNode_fieldImageViaFieldImageHeight Has Node_fieldImage via FieldImageHeight
func HasNode_fieldImageViaFieldImageHeight(iFieldImageHeight int) bool {
	if has, err := Engine.Where("field_image_height = ?", iFieldImageHeight).Get(new(Node_fieldImage)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// GetNode_fieldImageViaBundle Get Node_fieldImage via Bundle
func GetNode_fieldImageViaBundle(iBundle string) (*Node_fieldImage, error) {
	var _Node_fieldImage = &Node_fieldImage{Bundle: iBundle}
	has, err := Engine.Get(_Node_fieldImage)
	if has {
		return _Node_fieldImage, err
	} else {
		return nil, err
	}
}

// GetNode_fieldImageViaDeleted Get Node_fieldImage via Deleted
func GetNode_fieldImageViaDeleted(iDeleted int) (*Node_fieldImage, error) {
	var _Node_fieldImage = &Node_fieldImage{Deleted: iDeleted}
	has, err := Engine.Get(_Node_fieldImage)
	if has {
		return _Node_fieldImage, err
	} else {
		return nil, err
	}
}

// GetNode_fieldImageViaEntityId Get Node_fieldImage via EntityId
func GetNode_fieldImageViaEntityId(iEntityId int) (*Node_fieldImage, error) {
	var _Node_fieldImage = &Node_fieldImage{EntityId: iEntityId}
	has, err := Engine.Get(_Node_fieldImage)
	if has {
		return _Node_fieldImage, err
	} else {
		return nil, err
	}
}

// GetNode_fieldImageViaRevisionId Get Node_fieldImage via RevisionId
func GetNode_fieldImageViaRevisionId(iRevisionId int) (*Node_fieldImage, error) {
	var _Node_fieldImage = &Node_fieldImage{RevisionId: iRevisionId}
	has, err := Engine.Get(_Node_fieldImage)
	if has {
		return _Node_fieldImage, err
	} else {
		return nil, err
	}
}

// GetNode_fieldImageViaLangcode Get Node_fieldImage via Langcode
func GetNode_fieldImageViaLangcode(iLangcode string) (*Node_fieldImage, error) {
	var _Node_fieldImage = &Node_fieldImage{Langcode: iLangcode}
	has, err := Engine.Get(_Node_fieldImage)
	if has {
		return _Node_fieldImage, err
	} else {
		return nil, err
	}
}

// GetNode_fieldImageViaDelta Get Node_fieldImage via Delta
func GetNode_fieldImageViaDelta(iDelta int) (*Node_fieldImage, error) {
	var _Node_fieldImage = &Node_fieldImage{Delta: iDelta}
	has, err := Engine.Get(_Node_fieldImage)
	if has {
		return _Node_fieldImage, err
	} else {
		return nil, err
	}
}

// GetNode_fieldImageViaFieldImageTargetId Get Node_fieldImage via FieldImageTargetId
func GetNode_fieldImageViaFieldImageTargetId(iFieldImageTargetId int) (*Node_fieldImage, error) {
	var _Node_fieldImage = &Node_fieldImage{FieldImageTargetId: iFieldImageTargetId}
	has, err := Engine.Get(_Node_fieldImage)
	if has {
		return _Node_fieldImage, err
	} else {
		return nil, err
	}
}

// GetNode_fieldImageViaFieldImageAlt Get Node_fieldImage via FieldImageAlt
func GetNode_fieldImageViaFieldImageAlt(iFieldImageAlt string) (*Node_fieldImage, error) {
	var _Node_fieldImage = &Node_fieldImage{FieldImageAlt: iFieldImageAlt}
	has, err := Engine.Get(_Node_fieldImage)
	if has {
		return _Node_fieldImage, err
	} else {
		return nil, err
	}
}

// GetNode_fieldImageViaFieldImageTitle Get Node_fieldImage via FieldImageTitle
func GetNode_fieldImageViaFieldImageTitle(iFieldImageTitle string) (*Node_fieldImage, error) {
	var _Node_fieldImage = &Node_fieldImage{FieldImageTitle: iFieldImageTitle}
	has, err := Engine.Get(_Node_fieldImage)
	if has {
		return _Node_fieldImage, err
	} else {
		return nil, err
	}
}

// GetNode_fieldImageViaFieldImageWidth Get Node_fieldImage via FieldImageWidth
func GetNode_fieldImageViaFieldImageWidth(iFieldImageWidth int) (*Node_fieldImage, error) {
	var _Node_fieldImage = &Node_fieldImage{FieldImageWidth: iFieldImageWidth}
	has, err := Engine.Get(_Node_fieldImage)
	if has {
		return _Node_fieldImage, err
	} else {
		return nil, err
	}
}

// GetNode_fieldImageViaFieldImageHeight Get Node_fieldImage via FieldImageHeight
func GetNode_fieldImageViaFieldImageHeight(iFieldImageHeight int) (*Node_fieldImage, error) {
	var _Node_fieldImage = &Node_fieldImage{FieldImageHeight: iFieldImageHeight}
	has, err := Engine.Get(_Node_fieldImage)
	if has {
		return _Node_fieldImage, err
	} else {
		return nil, err
	}
}

// SetNode_fieldImageViaBundle Set Node_fieldImage via Bundle
func SetNode_fieldImageViaBundle(iBundle string, node__field_image *Node_fieldImage) (int64, error) {
	node__field_image.Bundle = iBundle
	return Engine.Insert(node__field_image)
}

// SetNode_fieldImageViaDeleted Set Node_fieldImage via Deleted
func SetNode_fieldImageViaDeleted(iDeleted int, node__field_image *Node_fieldImage) (int64, error) {
	node__field_image.Deleted = iDeleted
	return Engine.Insert(node__field_image)
}

// SetNode_fieldImageViaEntityId Set Node_fieldImage via EntityId
func SetNode_fieldImageViaEntityId(iEntityId int, node__field_image *Node_fieldImage) (int64, error) {
	node__field_image.EntityId = iEntityId
	return Engine.Insert(node__field_image)
}

// SetNode_fieldImageViaRevisionId Set Node_fieldImage via RevisionId
func SetNode_fieldImageViaRevisionId(iRevisionId int, node__field_image *Node_fieldImage) (int64, error) {
	node__field_image.RevisionId = iRevisionId
	return Engine.Insert(node__field_image)
}

// SetNode_fieldImageViaLangcode Set Node_fieldImage via Langcode
func SetNode_fieldImageViaLangcode(iLangcode string, node__field_image *Node_fieldImage) (int64, error) {
	node__field_image.Langcode = iLangcode
	return Engine.Insert(node__field_image)
}

// SetNode_fieldImageViaDelta Set Node_fieldImage via Delta
func SetNode_fieldImageViaDelta(iDelta int, node__field_image *Node_fieldImage) (int64, error) {
	node__field_image.Delta = iDelta
	return Engine.Insert(node__field_image)
}

// SetNode_fieldImageViaFieldImageTargetId Set Node_fieldImage via FieldImageTargetId
func SetNode_fieldImageViaFieldImageTargetId(iFieldImageTargetId int, node__field_image *Node_fieldImage) (int64, error) {
	node__field_image.FieldImageTargetId = iFieldImageTargetId
	return Engine.Insert(node__field_image)
}

// SetNode_fieldImageViaFieldImageAlt Set Node_fieldImage via FieldImageAlt
func SetNode_fieldImageViaFieldImageAlt(iFieldImageAlt string, node__field_image *Node_fieldImage) (int64, error) {
	node__field_image.FieldImageAlt = iFieldImageAlt
	return Engine.Insert(node__field_image)
}

// SetNode_fieldImageViaFieldImageTitle Set Node_fieldImage via FieldImageTitle
func SetNode_fieldImageViaFieldImageTitle(iFieldImageTitle string, node__field_image *Node_fieldImage) (int64, error) {
	node__field_image.FieldImageTitle = iFieldImageTitle
	return Engine.Insert(node__field_image)
}

// SetNode_fieldImageViaFieldImageWidth Set Node_fieldImage via FieldImageWidth
func SetNode_fieldImageViaFieldImageWidth(iFieldImageWidth int, node__field_image *Node_fieldImage) (int64, error) {
	node__field_image.FieldImageWidth = iFieldImageWidth
	return Engine.Insert(node__field_image)
}

// SetNode_fieldImageViaFieldImageHeight Set Node_fieldImage via FieldImageHeight
func SetNode_fieldImageViaFieldImageHeight(iFieldImageHeight int, node__field_image *Node_fieldImage) (int64, error) {
	node__field_image.FieldImageHeight = iFieldImageHeight
	return Engine.Insert(node__field_image)
}

// AddNode_fieldImage Add Node_fieldImage via all columns
func AddNode_fieldImage(iBundle string, iDeleted int, iEntityId int, iRevisionId int, iLangcode string, iDelta int, iFieldImageTargetId int, iFieldImageAlt string, iFieldImageTitle string, iFieldImageWidth int, iFieldImageHeight int) error {
	_Node_fieldImage := &Node_fieldImage{Bundle: iBundle, Deleted: iDeleted, EntityId: iEntityId, RevisionId: iRevisionId, Langcode: iLangcode, Delta: iDelta, FieldImageTargetId: iFieldImageTargetId, FieldImageAlt: iFieldImageAlt, FieldImageTitle: iFieldImageTitle, FieldImageWidth: iFieldImageWidth, FieldImageHeight: iFieldImageHeight}
	if _, err := Engine.Insert(_Node_fieldImage); err != nil {
		return err
	} else {
		return nil
	}
}

// PostNode_fieldImage Post Node_fieldImage via iNode_fieldImage
func PostNode_fieldImage(iNode_fieldImage *Node_fieldImage) (string, error) {
	_, err := Engine.Insert(iNode_fieldImage)
	return iNode_fieldImage.Bundle, err
}

// PutNode_fieldImage Put Node_fieldImage
func PutNode_fieldImage(iNode_fieldImage *Node_fieldImage) (string, error) {
	_, err := Engine.Id(iNode_fieldImage.Bundle).Update(iNode_fieldImage)
	return iNode_fieldImage.Bundle, err
}

// PutNode_fieldImageViaBundle Put Node_fieldImage via Bundle
func PutNode_fieldImageViaBundle(Bundle_ string, iNode_fieldImage *Node_fieldImage) (int64, error) {
	row, err := Engine.Update(iNode_fieldImage, &Node_fieldImage{Bundle: Bundle_})
	return row, err
}

// PutNode_fieldImageViaDeleted Put Node_fieldImage via Deleted
func PutNode_fieldImageViaDeleted(Deleted_ int, iNode_fieldImage *Node_fieldImage) (int64, error) {
	row, err := Engine.Update(iNode_fieldImage, &Node_fieldImage{Deleted: Deleted_})
	return row, err
}

// PutNode_fieldImageViaEntityId Put Node_fieldImage via EntityId
func PutNode_fieldImageViaEntityId(EntityId_ int, iNode_fieldImage *Node_fieldImage) (int64, error) {
	row, err := Engine.Update(iNode_fieldImage, &Node_fieldImage{EntityId: EntityId_})
	return row, err
}

// PutNode_fieldImageViaRevisionId Put Node_fieldImage via RevisionId
func PutNode_fieldImageViaRevisionId(RevisionId_ int, iNode_fieldImage *Node_fieldImage) (int64, error) {
	row, err := Engine.Update(iNode_fieldImage, &Node_fieldImage{RevisionId: RevisionId_})
	return row, err
}

// PutNode_fieldImageViaLangcode Put Node_fieldImage via Langcode
func PutNode_fieldImageViaLangcode(Langcode_ string, iNode_fieldImage *Node_fieldImage) (int64, error) {
	row, err := Engine.Update(iNode_fieldImage, &Node_fieldImage{Langcode: Langcode_})
	return row, err
}

// PutNode_fieldImageViaDelta Put Node_fieldImage via Delta
func PutNode_fieldImageViaDelta(Delta_ int, iNode_fieldImage *Node_fieldImage) (int64, error) {
	row, err := Engine.Update(iNode_fieldImage, &Node_fieldImage{Delta: Delta_})
	return row, err
}

// PutNode_fieldImageViaFieldImageTargetId Put Node_fieldImage via FieldImageTargetId
func PutNode_fieldImageViaFieldImageTargetId(FieldImageTargetId_ int, iNode_fieldImage *Node_fieldImage) (int64, error) {
	row, err := Engine.Update(iNode_fieldImage, &Node_fieldImage{FieldImageTargetId: FieldImageTargetId_})
	return row, err
}

// PutNode_fieldImageViaFieldImageAlt Put Node_fieldImage via FieldImageAlt
func PutNode_fieldImageViaFieldImageAlt(FieldImageAlt_ string, iNode_fieldImage *Node_fieldImage) (int64, error) {
	row, err := Engine.Update(iNode_fieldImage, &Node_fieldImage{FieldImageAlt: FieldImageAlt_})
	return row, err
}

// PutNode_fieldImageViaFieldImageTitle Put Node_fieldImage via FieldImageTitle
func PutNode_fieldImageViaFieldImageTitle(FieldImageTitle_ string, iNode_fieldImage *Node_fieldImage) (int64, error) {
	row, err := Engine.Update(iNode_fieldImage, &Node_fieldImage{FieldImageTitle: FieldImageTitle_})
	return row, err
}

// PutNode_fieldImageViaFieldImageWidth Put Node_fieldImage via FieldImageWidth
func PutNode_fieldImageViaFieldImageWidth(FieldImageWidth_ int, iNode_fieldImage *Node_fieldImage) (int64, error) {
	row, err := Engine.Update(iNode_fieldImage, &Node_fieldImage{FieldImageWidth: FieldImageWidth_})
	return row, err
}

// PutNode_fieldImageViaFieldImageHeight Put Node_fieldImage via FieldImageHeight
func PutNode_fieldImageViaFieldImageHeight(FieldImageHeight_ int, iNode_fieldImage *Node_fieldImage) (int64, error) {
	row, err := Engine.Update(iNode_fieldImage, &Node_fieldImage{FieldImageHeight: FieldImageHeight_})
	return row, err
}

// UpdateNode_fieldImageViaBundle via map[string]interface{}{}
func UpdateNode_fieldImageViaBundle(iBundle string, iNode_fieldImageMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Node_fieldImage)).Where("bundle = ?", iBundle).Update(iNode_fieldImageMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateNode_fieldImageViaDeleted via map[string]interface{}{}
func UpdateNode_fieldImageViaDeleted(iDeleted int, iNode_fieldImageMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Node_fieldImage)).Where("deleted = ?", iDeleted).Update(iNode_fieldImageMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateNode_fieldImageViaEntityId via map[string]interface{}{}
func UpdateNode_fieldImageViaEntityId(iEntityId int, iNode_fieldImageMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Node_fieldImage)).Where("entity_id = ?", iEntityId).Update(iNode_fieldImageMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateNode_fieldImageViaRevisionId via map[string]interface{}{}
func UpdateNode_fieldImageViaRevisionId(iRevisionId int, iNode_fieldImageMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Node_fieldImage)).Where("revision_id = ?", iRevisionId).Update(iNode_fieldImageMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateNode_fieldImageViaLangcode via map[string]interface{}{}
func UpdateNode_fieldImageViaLangcode(iLangcode string, iNode_fieldImageMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Node_fieldImage)).Where("langcode = ?", iLangcode).Update(iNode_fieldImageMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateNode_fieldImageViaDelta via map[string]interface{}{}
func UpdateNode_fieldImageViaDelta(iDelta int, iNode_fieldImageMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Node_fieldImage)).Where("delta = ?", iDelta).Update(iNode_fieldImageMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateNode_fieldImageViaFieldImageTargetId via map[string]interface{}{}
func UpdateNode_fieldImageViaFieldImageTargetId(iFieldImageTargetId int, iNode_fieldImageMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Node_fieldImage)).Where("field_image_target_id = ?", iFieldImageTargetId).Update(iNode_fieldImageMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateNode_fieldImageViaFieldImageAlt via map[string]interface{}{}
func UpdateNode_fieldImageViaFieldImageAlt(iFieldImageAlt string, iNode_fieldImageMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Node_fieldImage)).Where("field_image_alt = ?", iFieldImageAlt).Update(iNode_fieldImageMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateNode_fieldImageViaFieldImageTitle via map[string]interface{}{}
func UpdateNode_fieldImageViaFieldImageTitle(iFieldImageTitle string, iNode_fieldImageMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Node_fieldImage)).Where("field_image_title = ?", iFieldImageTitle).Update(iNode_fieldImageMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateNode_fieldImageViaFieldImageWidth via map[string]interface{}{}
func UpdateNode_fieldImageViaFieldImageWidth(iFieldImageWidth int, iNode_fieldImageMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Node_fieldImage)).Where("field_image_width = ?", iFieldImageWidth).Update(iNode_fieldImageMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateNode_fieldImageViaFieldImageHeight via map[string]interface{}{}
func UpdateNode_fieldImageViaFieldImageHeight(iFieldImageHeight int, iNode_fieldImageMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Node_fieldImage)).Where("field_image_height = ?", iFieldImageHeight).Update(iNode_fieldImageMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// DeleteNode_fieldImage Delete Node_fieldImage
func DeleteNode_fieldImage(iBundle string) (int64, error) {
	row, err := Engine.Id(iBundle).Delete(new(Node_fieldImage))
	return row, err
}

// DeleteNode_fieldImageViaBundle Delete Node_fieldImage via Bundle
func DeleteNode_fieldImageViaBundle(iBundle string) (err error) {
	var has bool
	var _Node_fieldImage = &Node_fieldImage{Bundle: iBundle}
	if has, err = Engine.Get(_Node_fieldImage); (has == true) && (err == nil) {
		if row, err := Engine.Where("bundle = ?", iBundle).Delete(new(Node_fieldImage)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteNode_fieldImageViaDeleted Delete Node_fieldImage via Deleted
func DeleteNode_fieldImageViaDeleted(iDeleted int) (err error) {
	var has bool
	var _Node_fieldImage = &Node_fieldImage{Deleted: iDeleted}
	if has, err = Engine.Get(_Node_fieldImage); (has == true) && (err == nil) {
		if row, err := Engine.Where("deleted = ?", iDeleted).Delete(new(Node_fieldImage)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteNode_fieldImageViaEntityId Delete Node_fieldImage via EntityId
func DeleteNode_fieldImageViaEntityId(iEntityId int) (err error) {
	var has bool
	var _Node_fieldImage = &Node_fieldImage{EntityId: iEntityId}
	if has, err = Engine.Get(_Node_fieldImage); (has == true) && (err == nil) {
		if row, err := Engine.Where("entity_id = ?", iEntityId).Delete(new(Node_fieldImage)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteNode_fieldImageViaRevisionId Delete Node_fieldImage via RevisionId
func DeleteNode_fieldImageViaRevisionId(iRevisionId int) (err error) {
	var has bool
	var _Node_fieldImage = &Node_fieldImage{RevisionId: iRevisionId}
	if has, err = Engine.Get(_Node_fieldImage); (has == true) && (err == nil) {
		if row, err := Engine.Where("revision_id = ?", iRevisionId).Delete(new(Node_fieldImage)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteNode_fieldImageViaLangcode Delete Node_fieldImage via Langcode
func DeleteNode_fieldImageViaLangcode(iLangcode string) (err error) {
	var has bool
	var _Node_fieldImage = &Node_fieldImage{Langcode: iLangcode}
	if has, err = Engine.Get(_Node_fieldImage); (has == true) && (err == nil) {
		if row, err := Engine.Where("langcode = ?", iLangcode).Delete(new(Node_fieldImage)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteNode_fieldImageViaDelta Delete Node_fieldImage via Delta
func DeleteNode_fieldImageViaDelta(iDelta int) (err error) {
	var has bool
	var _Node_fieldImage = &Node_fieldImage{Delta: iDelta}
	if has, err = Engine.Get(_Node_fieldImage); (has == true) && (err == nil) {
		if row, err := Engine.Where("delta = ?", iDelta).Delete(new(Node_fieldImage)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteNode_fieldImageViaFieldImageTargetId Delete Node_fieldImage via FieldImageTargetId
func DeleteNode_fieldImageViaFieldImageTargetId(iFieldImageTargetId int) (err error) {
	var has bool
	var _Node_fieldImage = &Node_fieldImage{FieldImageTargetId: iFieldImageTargetId}
	if has, err = Engine.Get(_Node_fieldImage); (has == true) && (err == nil) {
		if row, err := Engine.Where("field_image_target_id = ?", iFieldImageTargetId).Delete(new(Node_fieldImage)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteNode_fieldImageViaFieldImageAlt Delete Node_fieldImage via FieldImageAlt
func DeleteNode_fieldImageViaFieldImageAlt(iFieldImageAlt string) (err error) {
	var has bool
	var _Node_fieldImage = &Node_fieldImage{FieldImageAlt: iFieldImageAlt}
	if has, err = Engine.Get(_Node_fieldImage); (has == true) && (err == nil) {
		if row, err := Engine.Where("field_image_alt = ?", iFieldImageAlt).Delete(new(Node_fieldImage)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteNode_fieldImageViaFieldImageTitle Delete Node_fieldImage via FieldImageTitle
func DeleteNode_fieldImageViaFieldImageTitle(iFieldImageTitle string) (err error) {
	var has bool
	var _Node_fieldImage = &Node_fieldImage{FieldImageTitle: iFieldImageTitle}
	if has, err = Engine.Get(_Node_fieldImage); (has == true) && (err == nil) {
		if row, err := Engine.Where("field_image_title = ?", iFieldImageTitle).Delete(new(Node_fieldImage)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteNode_fieldImageViaFieldImageWidth Delete Node_fieldImage via FieldImageWidth
func DeleteNode_fieldImageViaFieldImageWidth(iFieldImageWidth int) (err error) {
	var has bool
	var _Node_fieldImage = &Node_fieldImage{FieldImageWidth: iFieldImageWidth}
	if has, err = Engine.Get(_Node_fieldImage); (has == true) && (err == nil) {
		if row, err := Engine.Where("field_image_width = ?", iFieldImageWidth).Delete(new(Node_fieldImage)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteNode_fieldImageViaFieldImageHeight Delete Node_fieldImage via FieldImageHeight
func DeleteNode_fieldImageViaFieldImageHeight(iFieldImageHeight int) (err error) {
	var has bool
	var _Node_fieldImage = &Node_fieldImage{FieldImageHeight: iFieldImageHeight}
	if has, err = Engine.Get(_Node_fieldImage); (has == true) && (err == nil) {
		if row, err := Engine.Where("field_image_height = ?", iFieldImageHeight).Delete(new(Node_fieldImage)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}
