package model

type User_userPicture struct {
	Bundle              string `xorm:"not null default '' index VARCHAR(128)"`
	Deleted             int    `xorm:"not null pk default 0 TINYINT(4)"`
	EntityId            int    `xorm:"not null pk INT(10)"`
	RevisionId          int    `xorm:"not null index INT(10)"`
	Langcode            string `xorm:"not null pk default '' VARCHAR(32)"`
	Delta               int    `xorm:"not null pk INT(10)"`
	UserPictureTargetId int    `xorm:"not null index INT(10)"`
	UserPictureAlt      string `xorm:"VARCHAR(512)"`
	UserPictureTitle    string `xorm:"VARCHAR(1024)"`
	UserPictureWidth    int    `xorm:"INT(10)"`
	UserPictureHeight   int    `xorm:"INT(10)"`
}

// GetUser_userPicturesCount User_userPictures' Count
func GetUser_userPicturesCount(offset int, limit int) (int64, error) {
	total, err := Engine.Limit(limit, offset).Count(&User_userPicture{})
	return total, err
}

// GetUser_userPictureCountViaBundle Get User_userPicture via Bundle
func GetUser_userPictureCountViaBundle(iBundle string) int64 {
	n, _ := Engine.Where("bundle = ?", iBundle).Count(&User_userPicture{Bundle: iBundle})
	return n
}

// GetUser_userPictureCountViaDeleted Get User_userPicture via Deleted
func GetUser_userPictureCountViaDeleted(iDeleted int) int64 {
	n, _ := Engine.Where("deleted = ?", iDeleted).Count(&User_userPicture{Deleted: iDeleted})
	return n
}

// GetUser_userPictureCountViaEntityId Get User_userPicture via EntityId
func GetUser_userPictureCountViaEntityId(iEntityId int) int64 {
	n, _ := Engine.Where("entity_id = ?", iEntityId).Count(&User_userPicture{EntityId: iEntityId})
	return n
}

// GetUser_userPictureCountViaRevisionId Get User_userPicture via RevisionId
func GetUser_userPictureCountViaRevisionId(iRevisionId int) int64 {
	n, _ := Engine.Where("revision_id = ?", iRevisionId).Count(&User_userPicture{RevisionId: iRevisionId})
	return n
}

// GetUser_userPictureCountViaLangcode Get User_userPicture via Langcode
func GetUser_userPictureCountViaLangcode(iLangcode string) int64 {
	n, _ := Engine.Where("langcode = ?", iLangcode).Count(&User_userPicture{Langcode: iLangcode})
	return n
}

// GetUser_userPictureCountViaDelta Get User_userPicture via Delta
func GetUser_userPictureCountViaDelta(iDelta int) int64 {
	n, _ := Engine.Where("delta = ?", iDelta).Count(&User_userPicture{Delta: iDelta})
	return n
}

// GetUser_userPictureCountViaUserPictureTargetId Get User_userPicture via UserPictureTargetId
func GetUser_userPictureCountViaUserPictureTargetId(iUserPictureTargetId int) int64 {
	n, _ := Engine.Where("user_picture_target_id = ?", iUserPictureTargetId).Count(&User_userPicture{UserPictureTargetId: iUserPictureTargetId})
	return n
}

// GetUser_userPictureCountViaUserPictureAlt Get User_userPicture via UserPictureAlt
func GetUser_userPictureCountViaUserPictureAlt(iUserPictureAlt string) int64 {
	n, _ := Engine.Where("user_picture_alt = ?", iUserPictureAlt).Count(&User_userPicture{UserPictureAlt: iUserPictureAlt})
	return n
}

// GetUser_userPictureCountViaUserPictureTitle Get User_userPicture via UserPictureTitle
func GetUser_userPictureCountViaUserPictureTitle(iUserPictureTitle string) int64 {
	n, _ := Engine.Where("user_picture_title = ?", iUserPictureTitle).Count(&User_userPicture{UserPictureTitle: iUserPictureTitle})
	return n
}

// GetUser_userPictureCountViaUserPictureWidth Get User_userPicture via UserPictureWidth
func GetUser_userPictureCountViaUserPictureWidth(iUserPictureWidth int) int64 {
	n, _ := Engine.Where("user_picture_width = ?", iUserPictureWidth).Count(&User_userPicture{UserPictureWidth: iUserPictureWidth})
	return n
}

// GetUser_userPictureCountViaUserPictureHeight Get User_userPicture via UserPictureHeight
func GetUser_userPictureCountViaUserPictureHeight(iUserPictureHeight int) int64 {
	n, _ := Engine.Where("user_picture_height = ?", iUserPictureHeight).Count(&User_userPicture{UserPictureHeight: iUserPictureHeight})
	return n
}

// GetUser_userPicturesByBundleAndDeletedAndEntityId Get User_userPictures via BundleAndDeletedAndEntityId
func GetUser_userPicturesByBundleAndDeletedAndEntityId(offset int, limit int, Bundle_ string, Deleted_ int, EntityId_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("bundle = ? and deleted = ? and entity_id = ?", Bundle_, Deleted_, EntityId_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByBundleAndDeletedAndRevisionId Get User_userPictures via BundleAndDeletedAndRevisionId
func GetUser_userPicturesByBundleAndDeletedAndRevisionId(offset int, limit int, Bundle_ string, Deleted_ int, RevisionId_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("bundle = ? and deleted = ? and revision_id = ?", Bundle_, Deleted_, RevisionId_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByBundleAndDeletedAndLangcode Get User_userPictures via BundleAndDeletedAndLangcode
func GetUser_userPicturesByBundleAndDeletedAndLangcode(offset int, limit int, Bundle_ string, Deleted_ int, Langcode_ string) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("bundle = ? and deleted = ? and langcode = ?", Bundle_, Deleted_, Langcode_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByBundleAndDeletedAndDelta Get User_userPictures via BundleAndDeletedAndDelta
func GetUser_userPicturesByBundleAndDeletedAndDelta(offset int, limit int, Bundle_ string, Deleted_ int, Delta_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("bundle = ? and deleted = ? and delta = ?", Bundle_, Deleted_, Delta_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByBundleAndDeletedAndUserPictureTargetId Get User_userPictures via BundleAndDeletedAndUserPictureTargetId
func GetUser_userPicturesByBundleAndDeletedAndUserPictureTargetId(offset int, limit int, Bundle_ string, Deleted_ int, UserPictureTargetId_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("bundle = ? and deleted = ? and user_picture_target_id = ?", Bundle_, Deleted_, UserPictureTargetId_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByBundleAndDeletedAndUserPictureAlt Get User_userPictures via BundleAndDeletedAndUserPictureAlt
func GetUser_userPicturesByBundleAndDeletedAndUserPictureAlt(offset int, limit int, Bundle_ string, Deleted_ int, UserPictureAlt_ string) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("bundle = ? and deleted = ? and user_picture_alt = ?", Bundle_, Deleted_, UserPictureAlt_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByBundleAndDeletedAndUserPictureTitle Get User_userPictures via BundleAndDeletedAndUserPictureTitle
func GetUser_userPicturesByBundleAndDeletedAndUserPictureTitle(offset int, limit int, Bundle_ string, Deleted_ int, UserPictureTitle_ string) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("bundle = ? and deleted = ? and user_picture_title = ?", Bundle_, Deleted_, UserPictureTitle_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByBundleAndDeletedAndUserPictureWidth Get User_userPictures via BundleAndDeletedAndUserPictureWidth
func GetUser_userPicturesByBundleAndDeletedAndUserPictureWidth(offset int, limit int, Bundle_ string, Deleted_ int, UserPictureWidth_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("bundle = ? and deleted = ? and user_picture_width = ?", Bundle_, Deleted_, UserPictureWidth_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByBundleAndDeletedAndUserPictureHeight Get User_userPictures via BundleAndDeletedAndUserPictureHeight
func GetUser_userPicturesByBundleAndDeletedAndUserPictureHeight(offset int, limit int, Bundle_ string, Deleted_ int, UserPictureHeight_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("bundle = ? and deleted = ? and user_picture_height = ?", Bundle_, Deleted_, UserPictureHeight_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByBundleAndEntityIdAndRevisionId Get User_userPictures via BundleAndEntityIdAndRevisionId
func GetUser_userPicturesByBundleAndEntityIdAndRevisionId(offset int, limit int, Bundle_ string, EntityId_ int, RevisionId_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("bundle = ? and entity_id = ? and revision_id = ?", Bundle_, EntityId_, RevisionId_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByBundleAndEntityIdAndLangcode Get User_userPictures via BundleAndEntityIdAndLangcode
func GetUser_userPicturesByBundleAndEntityIdAndLangcode(offset int, limit int, Bundle_ string, EntityId_ int, Langcode_ string) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("bundle = ? and entity_id = ? and langcode = ?", Bundle_, EntityId_, Langcode_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByBundleAndEntityIdAndDelta Get User_userPictures via BundleAndEntityIdAndDelta
func GetUser_userPicturesByBundleAndEntityIdAndDelta(offset int, limit int, Bundle_ string, EntityId_ int, Delta_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("bundle = ? and entity_id = ? and delta = ?", Bundle_, EntityId_, Delta_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByBundleAndEntityIdAndUserPictureTargetId Get User_userPictures via BundleAndEntityIdAndUserPictureTargetId
func GetUser_userPicturesByBundleAndEntityIdAndUserPictureTargetId(offset int, limit int, Bundle_ string, EntityId_ int, UserPictureTargetId_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("bundle = ? and entity_id = ? and user_picture_target_id = ?", Bundle_, EntityId_, UserPictureTargetId_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByBundleAndEntityIdAndUserPictureAlt Get User_userPictures via BundleAndEntityIdAndUserPictureAlt
func GetUser_userPicturesByBundleAndEntityIdAndUserPictureAlt(offset int, limit int, Bundle_ string, EntityId_ int, UserPictureAlt_ string) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("bundle = ? and entity_id = ? and user_picture_alt = ?", Bundle_, EntityId_, UserPictureAlt_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByBundleAndEntityIdAndUserPictureTitle Get User_userPictures via BundleAndEntityIdAndUserPictureTitle
func GetUser_userPicturesByBundleAndEntityIdAndUserPictureTitle(offset int, limit int, Bundle_ string, EntityId_ int, UserPictureTitle_ string) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("bundle = ? and entity_id = ? and user_picture_title = ?", Bundle_, EntityId_, UserPictureTitle_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByBundleAndEntityIdAndUserPictureWidth Get User_userPictures via BundleAndEntityIdAndUserPictureWidth
func GetUser_userPicturesByBundleAndEntityIdAndUserPictureWidth(offset int, limit int, Bundle_ string, EntityId_ int, UserPictureWidth_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("bundle = ? and entity_id = ? and user_picture_width = ?", Bundle_, EntityId_, UserPictureWidth_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByBundleAndEntityIdAndUserPictureHeight Get User_userPictures via BundleAndEntityIdAndUserPictureHeight
func GetUser_userPicturesByBundleAndEntityIdAndUserPictureHeight(offset int, limit int, Bundle_ string, EntityId_ int, UserPictureHeight_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("bundle = ? and entity_id = ? and user_picture_height = ?", Bundle_, EntityId_, UserPictureHeight_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByBundleAndRevisionIdAndLangcode Get User_userPictures via BundleAndRevisionIdAndLangcode
func GetUser_userPicturesByBundleAndRevisionIdAndLangcode(offset int, limit int, Bundle_ string, RevisionId_ int, Langcode_ string) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("bundle = ? and revision_id = ? and langcode = ?", Bundle_, RevisionId_, Langcode_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByBundleAndRevisionIdAndDelta Get User_userPictures via BundleAndRevisionIdAndDelta
func GetUser_userPicturesByBundleAndRevisionIdAndDelta(offset int, limit int, Bundle_ string, RevisionId_ int, Delta_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("bundle = ? and revision_id = ? and delta = ?", Bundle_, RevisionId_, Delta_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByBundleAndRevisionIdAndUserPictureTargetId Get User_userPictures via BundleAndRevisionIdAndUserPictureTargetId
func GetUser_userPicturesByBundleAndRevisionIdAndUserPictureTargetId(offset int, limit int, Bundle_ string, RevisionId_ int, UserPictureTargetId_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("bundle = ? and revision_id = ? and user_picture_target_id = ?", Bundle_, RevisionId_, UserPictureTargetId_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByBundleAndRevisionIdAndUserPictureAlt Get User_userPictures via BundleAndRevisionIdAndUserPictureAlt
func GetUser_userPicturesByBundleAndRevisionIdAndUserPictureAlt(offset int, limit int, Bundle_ string, RevisionId_ int, UserPictureAlt_ string) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("bundle = ? and revision_id = ? and user_picture_alt = ?", Bundle_, RevisionId_, UserPictureAlt_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByBundleAndRevisionIdAndUserPictureTitle Get User_userPictures via BundleAndRevisionIdAndUserPictureTitle
func GetUser_userPicturesByBundleAndRevisionIdAndUserPictureTitle(offset int, limit int, Bundle_ string, RevisionId_ int, UserPictureTitle_ string) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("bundle = ? and revision_id = ? and user_picture_title = ?", Bundle_, RevisionId_, UserPictureTitle_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByBundleAndRevisionIdAndUserPictureWidth Get User_userPictures via BundleAndRevisionIdAndUserPictureWidth
func GetUser_userPicturesByBundleAndRevisionIdAndUserPictureWidth(offset int, limit int, Bundle_ string, RevisionId_ int, UserPictureWidth_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("bundle = ? and revision_id = ? and user_picture_width = ?", Bundle_, RevisionId_, UserPictureWidth_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByBundleAndRevisionIdAndUserPictureHeight Get User_userPictures via BundleAndRevisionIdAndUserPictureHeight
func GetUser_userPicturesByBundleAndRevisionIdAndUserPictureHeight(offset int, limit int, Bundle_ string, RevisionId_ int, UserPictureHeight_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("bundle = ? and revision_id = ? and user_picture_height = ?", Bundle_, RevisionId_, UserPictureHeight_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByBundleAndLangcodeAndDelta Get User_userPictures via BundleAndLangcodeAndDelta
func GetUser_userPicturesByBundleAndLangcodeAndDelta(offset int, limit int, Bundle_ string, Langcode_ string, Delta_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("bundle = ? and langcode = ? and delta = ?", Bundle_, Langcode_, Delta_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByBundleAndLangcodeAndUserPictureTargetId Get User_userPictures via BundleAndLangcodeAndUserPictureTargetId
func GetUser_userPicturesByBundleAndLangcodeAndUserPictureTargetId(offset int, limit int, Bundle_ string, Langcode_ string, UserPictureTargetId_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("bundle = ? and langcode = ? and user_picture_target_id = ?", Bundle_, Langcode_, UserPictureTargetId_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByBundleAndLangcodeAndUserPictureAlt Get User_userPictures via BundleAndLangcodeAndUserPictureAlt
func GetUser_userPicturesByBundleAndLangcodeAndUserPictureAlt(offset int, limit int, Bundle_ string, Langcode_ string, UserPictureAlt_ string) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("bundle = ? and langcode = ? and user_picture_alt = ?", Bundle_, Langcode_, UserPictureAlt_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByBundleAndLangcodeAndUserPictureTitle Get User_userPictures via BundleAndLangcodeAndUserPictureTitle
func GetUser_userPicturesByBundleAndLangcodeAndUserPictureTitle(offset int, limit int, Bundle_ string, Langcode_ string, UserPictureTitle_ string) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("bundle = ? and langcode = ? and user_picture_title = ?", Bundle_, Langcode_, UserPictureTitle_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByBundleAndLangcodeAndUserPictureWidth Get User_userPictures via BundleAndLangcodeAndUserPictureWidth
func GetUser_userPicturesByBundleAndLangcodeAndUserPictureWidth(offset int, limit int, Bundle_ string, Langcode_ string, UserPictureWidth_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("bundle = ? and langcode = ? and user_picture_width = ?", Bundle_, Langcode_, UserPictureWidth_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByBundleAndLangcodeAndUserPictureHeight Get User_userPictures via BundleAndLangcodeAndUserPictureHeight
func GetUser_userPicturesByBundleAndLangcodeAndUserPictureHeight(offset int, limit int, Bundle_ string, Langcode_ string, UserPictureHeight_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("bundle = ? and langcode = ? and user_picture_height = ?", Bundle_, Langcode_, UserPictureHeight_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByBundleAndDeltaAndUserPictureTargetId Get User_userPictures via BundleAndDeltaAndUserPictureTargetId
func GetUser_userPicturesByBundleAndDeltaAndUserPictureTargetId(offset int, limit int, Bundle_ string, Delta_ int, UserPictureTargetId_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("bundle = ? and delta = ? and user_picture_target_id = ?", Bundle_, Delta_, UserPictureTargetId_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByBundleAndDeltaAndUserPictureAlt Get User_userPictures via BundleAndDeltaAndUserPictureAlt
func GetUser_userPicturesByBundleAndDeltaAndUserPictureAlt(offset int, limit int, Bundle_ string, Delta_ int, UserPictureAlt_ string) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("bundle = ? and delta = ? and user_picture_alt = ?", Bundle_, Delta_, UserPictureAlt_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByBundleAndDeltaAndUserPictureTitle Get User_userPictures via BundleAndDeltaAndUserPictureTitle
func GetUser_userPicturesByBundleAndDeltaAndUserPictureTitle(offset int, limit int, Bundle_ string, Delta_ int, UserPictureTitle_ string) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("bundle = ? and delta = ? and user_picture_title = ?", Bundle_, Delta_, UserPictureTitle_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByBundleAndDeltaAndUserPictureWidth Get User_userPictures via BundleAndDeltaAndUserPictureWidth
func GetUser_userPicturesByBundleAndDeltaAndUserPictureWidth(offset int, limit int, Bundle_ string, Delta_ int, UserPictureWidth_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("bundle = ? and delta = ? and user_picture_width = ?", Bundle_, Delta_, UserPictureWidth_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByBundleAndDeltaAndUserPictureHeight Get User_userPictures via BundleAndDeltaAndUserPictureHeight
func GetUser_userPicturesByBundleAndDeltaAndUserPictureHeight(offset int, limit int, Bundle_ string, Delta_ int, UserPictureHeight_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("bundle = ? and delta = ? and user_picture_height = ?", Bundle_, Delta_, UserPictureHeight_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByBundleAndUserPictureTargetIdAndUserPictureAlt Get User_userPictures via BundleAndUserPictureTargetIdAndUserPictureAlt
func GetUser_userPicturesByBundleAndUserPictureTargetIdAndUserPictureAlt(offset int, limit int, Bundle_ string, UserPictureTargetId_ int, UserPictureAlt_ string) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("bundle = ? and user_picture_target_id = ? and user_picture_alt = ?", Bundle_, UserPictureTargetId_, UserPictureAlt_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByBundleAndUserPictureTargetIdAndUserPictureTitle Get User_userPictures via BundleAndUserPictureTargetIdAndUserPictureTitle
func GetUser_userPicturesByBundleAndUserPictureTargetIdAndUserPictureTitle(offset int, limit int, Bundle_ string, UserPictureTargetId_ int, UserPictureTitle_ string) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("bundle = ? and user_picture_target_id = ? and user_picture_title = ?", Bundle_, UserPictureTargetId_, UserPictureTitle_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByBundleAndUserPictureTargetIdAndUserPictureWidth Get User_userPictures via BundleAndUserPictureTargetIdAndUserPictureWidth
func GetUser_userPicturesByBundleAndUserPictureTargetIdAndUserPictureWidth(offset int, limit int, Bundle_ string, UserPictureTargetId_ int, UserPictureWidth_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("bundle = ? and user_picture_target_id = ? and user_picture_width = ?", Bundle_, UserPictureTargetId_, UserPictureWidth_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByBundleAndUserPictureTargetIdAndUserPictureHeight Get User_userPictures via BundleAndUserPictureTargetIdAndUserPictureHeight
func GetUser_userPicturesByBundleAndUserPictureTargetIdAndUserPictureHeight(offset int, limit int, Bundle_ string, UserPictureTargetId_ int, UserPictureHeight_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("bundle = ? and user_picture_target_id = ? and user_picture_height = ?", Bundle_, UserPictureTargetId_, UserPictureHeight_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByBundleAndUserPictureAltAndUserPictureTitle Get User_userPictures via BundleAndUserPictureAltAndUserPictureTitle
func GetUser_userPicturesByBundleAndUserPictureAltAndUserPictureTitle(offset int, limit int, Bundle_ string, UserPictureAlt_ string, UserPictureTitle_ string) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("bundle = ? and user_picture_alt = ? and user_picture_title = ?", Bundle_, UserPictureAlt_, UserPictureTitle_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByBundleAndUserPictureAltAndUserPictureWidth Get User_userPictures via BundleAndUserPictureAltAndUserPictureWidth
func GetUser_userPicturesByBundleAndUserPictureAltAndUserPictureWidth(offset int, limit int, Bundle_ string, UserPictureAlt_ string, UserPictureWidth_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("bundle = ? and user_picture_alt = ? and user_picture_width = ?", Bundle_, UserPictureAlt_, UserPictureWidth_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByBundleAndUserPictureAltAndUserPictureHeight Get User_userPictures via BundleAndUserPictureAltAndUserPictureHeight
func GetUser_userPicturesByBundleAndUserPictureAltAndUserPictureHeight(offset int, limit int, Bundle_ string, UserPictureAlt_ string, UserPictureHeight_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("bundle = ? and user_picture_alt = ? and user_picture_height = ?", Bundle_, UserPictureAlt_, UserPictureHeight_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByBundleAndUserPictureTitleAndUserPictureWidth Get User_userPictures via BundleAndUserPictureTitleAndUserPictureWidth
func GetUser_userPicturesByBundleAndUserPictureTitleAndUserPictureWidth(offset int, limit int, Bundle_ string, UserPictureTitle_ string, UserPictureWidth_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("bundle = ? and user_picture_title = ? and user_picture_width = ?", Bundle_, UserPictureTitle_, UserPictureWidth_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByBundleAndUserPictureTitleAndUserPictureHeight Get User_userPictures via BundleAndUserPictureTitleAndUserPictureHeight
func GetUser_userPicturesByBundleAndUserPictureTitleAndUserPictureHeight(offset int, limit int, Bundle_ string, UserPictureTitle_ string, UserPictureHeight_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("bundle = ? and user_picture_title = ? and user_picture_height = ?", Bundle_, UserPictureTitle_, UserPictureHeight_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByBundleAndUserPictureWidthAndUserPictureHeight Get User_userPictures via BundleAndUserPictureWidthAndUserPictureHeight
func GetUser_userPicturesByBundleAndUserPictureWidthAndUserPictureHeight(offset int, limit int, Bundle_ string, UserPictureWidth_ int, UserPictureHeight_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("bundle = ? and user_picture_width = ? and user_picture_height = ?", Bundle_, UserPictureWidth_, UserPictureHeight_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByDeletedAndEntityIdAndRevisionId Get User_userPictures via DeletedAndEntityIdAndRevisionId
func GetUser_userPicturesByDeletedAndEntityIdAndRevisionId(offset int, limit int, Deleted_ int, EntityId_ int, RevisionId_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("deleted = ? and entity_id = ? and revision_id = ?", Deleted_, EntityId_, RevisionId_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByDeletedAndEntityIdAndLangcode Get User_userPictures via DeletedAndEntityIdAndLangcode
func GetUser_userPicturesByDeletedAndEntityIdAndLangcode(offset int, limit int, Deleted_ int, EntityId_ int, Langcode_ string) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("deleted = ? and entity_id = ? and langcode = ?", Deleted_, EntityId_, Langcode_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByDeletedAndEntityIdAndDelta Get User_userPictures via DeletedAndEntityIdAndDelta
func GetUser_userPicturesByDeletedAndEntityIdAndDelta(offset int, limit int, Deleted_ int, EntityId_ int, Delta_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("deleted = ? and entity_id = ? and delta = ?", Deleted_, EntityId_, Delta_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByDeletedAndEntityIdAndUserPictureTargetId Get User_userPictures via DeletedAndEntityIdAndUserPictureTargetId
func GetUser_userPicturesByDeletedAndEntityIdAndUserPictureTargetId(offset int, limit int, Deleted_ int, EntityId_ int, UserPictureTargetId_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("deleted = ? and entity_id = ? and user_picture_target_id = ?", Deleted_, EntityId_, UserPictureTargetId_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByDeletedAndEntityIdAndUserPictureAlt Get User_userPictures via DeletedAndEntityIdAndUserPictureAlt
func GetUser_userPicturesByDeletedAndEntityIdAndUserPictureAlt(offset int, limit int, Deleted_ int, EntityId_ int, UserPictureAlt_ string) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("deleted = ? and entity_id = ? and user_picture_alt = ?", Deleted_, EntityId_, UserPictureAlt_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByDeletedAndEntityIdAndUserPictureTitle Get User_userPictures via DeletedAndEntityIdAndUserPictureTitle
func GetUser_userPicturesByDeletedAndEntityIdAndUserPictureTitle(offset int, limit int, Deleted_ int, EntityId_ int, UserPictureTitle_ string) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("deleted = ? and entity_id = ? and user_picture_title = ?", Deleted_, EntityId_, UserPictureTitle_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByDeletedAndEntityIdAndUserPictureWidth Get User_userPictures via DeletedAndEntityIdAndUserPictureWidth
func GetUser_userPicturesByDeletedAndEntityIdAndUserPictureWidth(offset int, limit int, Deleted_ int, EntityId_ int, UserPictureWidth_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("deleted = ? and entity_id = ? and user_picture_width = ?", Deleted_, EntityId_, UserPictureWidth_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByDeletedAndEntityIdAndUserPictureHeight Get User_userPictures via DeletedAndEntityIdAndUserPictureHeight
func GetUser_userPicturesByDeletedAndEntityIdAndUserPictureHeight(offset int, limit int, Deleted_ int, EntityId_ int, UserPictureHeight_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("deleted = ? and entity_id = ? and user_picture_height = ?", Deleted_, EntityId_, UserPictureHeight_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByDeletedAndRevisionIdAndLangcode Get User_userPictures via DeletedAndRevisionIdAndLangcode
func GetUser_userPicturesByDeletedAndRevisionIdAndLangcode(offset int, limit int, Deleted_ int, RevisionId_ int, Langcode_ string) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("deleted = ? and revision_id = ? and langcode = ?", Deleted_, RevisionId_, Langcode_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByDeletedAndRevisionIdAndDelta Get User_userPictures via DeletedAndRevisionIdAndDelta
func GetUser_userPicturesByDeletedAndRevisionIdAndDelta(offset int, limit int, Deleted_ int, RevisionId_ int, Delta_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("deleted = ? and revision_id = ? and delta = ?", Deleted_, RevisionId_, Delta_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByDeletedAndRevisionIdAndUserPictureTargetId Get User_userPictures via DeletedAndRevisionIdAndUserPictureTargetId
func GetUser_userPicturesByDeletedAndRevisionIdAndUserPictureTargetId(offset int, limit int, Deleted_ int, RevisionId_ int, UserPictureTargetId_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("deleted = ? and revision_id = ? and user_picture_target_id = ?", Deleted_, RevisionId_, UserPictureTargetId_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByDeletedAndRevisionIdAndUserPictureAlt Get User_userPictures via DeletedAndRevisionIdAndUserPictureAlt
func GetUser_userPicturesByDeletedAndRevisionIdAndUserPictureAlt(offset int, limit int, Deleted_ int, RevisionId_ int, UserPictureAlt_ string) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("deleted = ? and revision_id = ? and user_picture_alt = ?", Deleted_, RevisionId_, UserPictureAlt_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByDeletedAndRevisionIdAndUserPictureTitle Get User_userPictures via DeletedAndRevisionIdAndUserPictureTitle
func GetUser_userPicturesByDeletedAndRevisionIdAndUserPictureTitle(offset int, limit int, Deleted_ int, RevisionId_ int, UserPictureTitle_ string) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("deleted = ? and revision_id = ? and user_picture_title = ?", Deleted_, RevisionId_, UserPictureTitle_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByDeletedAndRevisionIdAndUserPictureWidth Get User_userPictures via DeletedAndRevisionIdAndUserPictureWidth
func GetUser_userPicturesByDeletedAndRevisionIdAndUserPictureWidth(offset int, limit int, Deleted_ int, RevisionId_ int, UserPictureWidth_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("deleted = ? and revision_id = ? and user_picture_width = ?", Deleted_, RevisionId_, UserPictureWidth_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByDeletedAndRevisionIdAndUserPictureHeight Get User_userPictures via DeletedAndRevisionIdAndUserPictureHeight
func GetUser_userPicturesByDeletedAndRevisionIdAndUserPictureHeight(offset int, limit int, Deleted_ int, RevisionId_ int, UserPictureHeight_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("deleted = ? and revision_id = ? and user_picture_height = ?", Deleted_, RevisionId_, UserPictureHeight_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByDeletedAndLangcodeAndDelta Get User_userPictures via DeletedAndLangcodeAndDelta
func GetUser_userPicturesByDeletedAndLangcodeAndDelta(offset int, limit int, Deleted_ int, Langcode_ string, Delta_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("deleted = ? and langcode = ? and delta = ?", Deleted_, Langcode_, Delta_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByDeletedAndLangcodeAndUserPictureTargetId Get User_userPictures via DeletedAndLangcodeAndUserPictureTargetId
func GetUser_userPicturesByDeletedAndLangcodeAndUserPictureTargetId(offset int, limit int, Deleted_ int, Langcode_ string, UserPictureTargetId_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("deleted = ? and langcode = ? and user_picture_target_id = ?", Deleted_, Langcode_, UserPictureTargetId_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByDeletedAndLangcodeAndUserPictureAlt Get User_userPictures via DeletedAndLangcodeAndUserPictureAlt
func GetUser_userPicturesByDeletedAndLangcodeAndUserPictureAlt(offset int, limit int, Deleted_ int, Langcode_ string, UserPictureAlt_ string) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("deleted = ? and langcode = ? and user_picture_alt = ?", Deleted_, Langcode_, UserPictureAlt_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByDeletedAndLangcodeAndUserPictureTitle Get User_userPictures via DeletedAndLangcodeAndUserPictureTitle
func GetUser_userPicturesByDeletedAndLangcodeAndUserPictureTitle(offset int, limit int, Deleted_ int, Langcode_ string, UserPictureTitle_ string) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("deleted = ? and langcode = ? and user_picture_title = ?", Deleted_, Langcode_, UserPictureTitle_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByDeletedAndLangcodeAndUserPictureWidth Get User_userPictures via DeletedAndLangcodeAndUserPictureWidth
func GetUser_userPicturesByDeletedAndLangcodeAndUserPictureWidth(offset int, limit int, Deleted_ int, Langcode_ string, UserPictureWidth_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("deleted = ? and langcode = ? and user_picture_width = ?", Deleted_, Langcode_, UserPictureWidth_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByDeletedAndLangcodeAndUserPictureHeight Get User_userPictures via DeletedAndLangcodeAndUserPictureHeight
func GetUser_userPicturesByDeletedAndLangcodeAndUserPictureHeight(offset int, limit int, Deleted_ int, Langcode_ string, UserPictureHeight_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("deleted = ? and langcode = ? and user_picture_height = ?", Deleted_, Langcode_, UserPictureHeight_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByDeletedAndDeltaAndUserPictureTargetId Get User_userPictures via DeletedAndDeltaAndUserPictureTargetId
func GetUser_userPicturesByDeletedAndDeltaAndUserPictureTargetId(offset int, limit int, Deleted_ int, Delta_ int, UserPictureTargetId_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("deleted = ? and delta = ? and user_picture_target_id = ?", Deleted_, Delta_, UserPictureTargetId_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByDeletedAndDeltaAndUserPictureAlt Get User_userPictures via DeletedAndDeltaAndUserPictureAlt
func GetUser_userPicturesByDeletedAndDeltaAndUserPictureAlt(offset int, limit int, Deleted_ int, Delta_ int, UserPictureAlt_ string) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("deleted = ? and delta = ? and user_picture_alt = ?", Deleted_, Delta_, UserPictureAlt_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByDeletedAndDeltaAndUserPictureTitle Get User_userPictures via DeletedAndDeltaAndUserPictureTitle
func GetUser_userPicturesByDeletedAndDeltaAndUserPictureTitle(offset int, limit int, Deleted_ int, Delta_ int, UserPictureTitle_ string) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("deleted = ? and delta = ? and user_picture_title = ?", Deleted_, Delta_, UserPictureTitle_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByDeletedAndDeltaAndUserPictureWidth Get User_userPictures via DeletedAndDeltaAndUserPictureWidth
func GetUser_userPicturesByDeletedAndDeltaAndUserPictureWidth(offset int, limit int, Deleted_ int, Delta_ int, UserPictureWidth_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("deleted = ? and delta = ? and user_picture_width = ?", Deleted_, Delta_, UserPictureWidth_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByDeletedAndDeltaAndUserPictureHeight Get User_userPictures via DeletedAndDeltaAndUserPictureHeight
func GetUser_userPicturesByDeletedAndDeltaAndUserPictureHeight(offset int, limit int, Deleted_ int, Delta_ int, UserPictureHeight_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("deleted = ? and delta = ? and user_picture_height = ?", Deleted_, Delta_, UserPictureHeight_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByDeletedAndUserPictureTargetIdAndUserPictureAlt Get User_userPictures via DeletedAndUserPictureTargetIdAndUserPictureAlt
func GetUser_userPicturesByDeletedAndUserPictureTargetIdAndUserPictureAlt(offset int, limit int, Deleted_ int, UserPictureTargetId_ int, UserPictureAlt_ string) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("deleted = ? and user_picture_target_id = ? and user_picture_alt = ?", Deleted_, UserPictureTargetId_, UserPictureAlt_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByDeletedAndUserPictureTargetIdAndUserPictureTitle Get User_userPictures via DeletedAndUserPictureTargetIdAndUserPictureTitle
func GetUser_userPicturesByDeletedAndUserPictureTargetIdAndUserPictureTitle(offset int, limit int, Deleted_ int, UserPictureTargetId_ int, UserPictureTitle_ string) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("deleted = ? and user_picture_target_id = ? and user_picture_title = ?", Deleted_, UserPictureTargetId_, UserPictureTitle_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByDeletedAndUserPictureTargetIdAndUserPictureWidth Get User_userPictures via DeletedAndUserPictureTargetIdAndUserPictureWidth
func GetUser_userPicturesByDeletedAndUserPictureTargetIdAndUserPictureWidth(offset int, limit int, Deleted_ int, UserPictureTargetId_ int, UserPictureWidth_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("deleted = ? and user_picture_target_id = ? and user_picture_width = ?", Deleted_, UserPictureTargetId_, UserPictureWidth_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByDeletedAndUserPictureTargetIdAndUserPictureHeight Get User_userPictures via DeletedAndUserPictureTargetIdAndUserPictureHeight
func GetUser_userPicturesByDeletedAndUserPictureTargetIdAndUserPictureHeight(offset int, limit int, Deleted_ int, UserPictureTargetId_ int, UserPictureHeight_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("deleted = ? and user_picture_target_id = ? and user_picture_height = ?", Deleted_, UserPictureTargetId_, UserPictureHeight_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByDeletedAndUserPictureAltAndUserPictureTitle Get User_userPictures via DeletedAndUserPictureAltAndUserPictureTitle
func GetUser_userPicturesByDeletedAndUserPictureAltAndUserPictureTitle(offset int, limit int, Deleted_ int, UserPictureAlt_ string, UserPictureTitle_ string) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("deleted = ? and user_picture_alt = ? and user_picture_title = ?", Deleted_, UserPictureAlt_, UserPictureTitle_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByDeletedAndUserPictureAltAndUserPictureWidth Get User_userPictures via DeletedAndUserPictureAltAndUserPictureWidth
func GetUser_userPicturesByDeletedAndUserPictureAltAndUserPictureWidth(offset int, limit int, Deleted_ int, UserPictureAlt_ string, UserPictureWidth_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("deleted = ? and user_picture_alt = ? and user_picture_width = ?", Deleted_, UserPictureAlt_, UserPictureWidth_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByDeletedAndUserPictureAltAndUserPictureHeight Get User_userPictures via DeletedAndUserPictureAltAndUserPictureHeight
func GetUser_userPicturesByDeletedAndUserPictureAltAndUserPictureHeight(offset int, limit int, Deleted_ int, UserPictureAlt_ string, UserPictureHeight_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("deleted = ? and user_picture_alt = ? and user_picture_height = ?", Deleted_, UserPictureAlt_, UserPictureHeight_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByDeletedAndUserPictureTitleAndUserPictureWidth Get User_userPictures via DeletedAndUserPictureTitleAndUserPictureWidth
func GetUser_userPicturesByDeletedAndUserPictureTitleAndUserPictureWidth(offset int, limit int, Deleted_ int, UserPictureTitle_ string, UserPictureWidth_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("deleted = ? and user_picture_title = ? and user_picture_width = ?", Deleted_, UserPictureTitle_, UserPictureWidth_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByDeletedAndUserPictureTitleAndUserPictureHeight Get User_userPictures via DeletedAndUserPictureTitleAndUserPictureHeight
func GetUser_userPicturesByDeletedAndUserPictureTitleAndUserPictureHeight(offset int, limit int, Deleted_ int, UserPictureTitle_ string, UserPictureHeight_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("deleted = ? and user_picture_title = ? and user_picture_height = ?", Deleted_, UserPictureTitle_, UserPictureHeight_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByDeletedAndUserPictureWidthAndUserPictureHeight Get User_userPictures via DeletedAndUserPictureWidthAndUserPictureHeight
func GetUser_userPicturesByDeletedAndUserPictureWidthAndUserPictureHeight(offset int, limit int, Deleted_ int, UserPictureWidth_ int, UserPictureHeight_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("deleted = ? and user_picture_width = ? and user_picture_height = ?", Deleted_, UserPictureWidth_, UserPictureHeight_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByEntityIdAndRevisionIdAndLangcode Get User_userPictures via EntityIdAndRevisionIdAndLangcode
func GetUser_userPicturesByEntityIdAndRevisionIdAndLangcode(offset int, limit int, EntityId_ int, RevisionId_ int, Langcode_ string) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("entity_id = ? and revision_id = ? and langcode = ?", EntityId_, RevisionId_, Langcode_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByEntityIdAndRevisionIdAndDelta Get User_userPictures via EntityIdAndRevisionIdAndDelta
func GetUser_userPicturesByEntityIdAndRevisionIdAndDelta(offset int, limit int, EntityId_ int, RevisionId_ int, Delta_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("entity_id = ? and revision_id = ? and delta = ?", EntityId_, RevisionId_, Delta_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByEntityIdAndRevisionIdAndUserPictureTargetId Get User_userPictures via EntityIdAndRevisionIdAndUserPictureTargetId
func GetUser_userPicturesByEntityIdAndRevisionIdAndUserPictureTargetId(offset int, limit int, EntityId_ int, RevisionId_ int, UserPictureTargetId_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("entity_id = ? and revision_id = ? and user_picture_target_id = ?", EntityId_, RevisionId_, UserPictureTargetId_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByEntityIdAndRevisionIdAndUserPictureAlt Get User_userPictures via EntityIdAndRevisionIdAndUserPictureAlt
func GetUser_userPicturesByEntityIdAndRevisionIdAndUserPictureAlt(offset int, limit int, EntityId_ int, RevisionId_ int, UserPictureAlt_ string) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("entity_id = ? and revision_id = ? and user_picture_alt = ?", EntityId_, RevisionId_, UserPictureAlt_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByEntityIdAndRevisionIdAndUserPictureTitle Get User_userPictures via EntityIdAndRevisionIdAndUserPictureTitle
func GetUser_userPicturesByEntityIdAndRevisionIdAndUserPictureTitle(offset int, limit int, EntityId_ int, RevisionId_ int, UserPictureTitle_ string) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("entity_id = ? and revision_id = ? and user_picture_title = ?", EntityId_, RevisionId_, UserPictureTitle_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByEntityIdAndRevisionIdAndUserPictureWidth Get User_userPictures via EntityIdAndRevisionIdAndUserPictureWidth
func GetUser_userPicturesByEntityIdAndRevisionIdAndUserPictureWidth(offset int, limit int, EntityId_ int, RevisionId_ int, UserPictureWidth_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("entity_id = ? and revision_id = ? and user_picture_width = ?", EntityId_, RevisionId_, UserPictureWidth_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByEntityIdAndRevisionIdAndUserPictureHeight Get User_userPictures via EntityIdAndRevisionIdAndUserPictureHeight
func GetUser_userPicturesByEntityIdAndRevisionIdAndUserPictureHeight(offset int, limit int, EntityId_ int, RevisionId_ int, UserPictureHeight_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("entity_id = ? and revision_id = ? and user_picture_height = ?", EntityId_, RevisionId_, UserPictureHeight_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByEntityIdAndLangcodeAndDelta Get User_userPictures via EntityIdAndLangcodeAndDelta
func GetUser_userPicturesByEntityIdAndLangcodeAndDelta(offset int, limit int, EntityId_ int, Langcode_ string, Delta_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("entity_id = ? and langcode = ? and delta = ?", EntityId_, Langcode_, Delta_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByEntityIdAndLangcodeAndUserPictureTargetId Get User_userPictures via EntityIdAndLangcodeAndUserPictureTargetId
func GetUser_userPicturesByEntityIdAndLangcodeAndUserPictureTargetId(offset int, limit int, EntityId_ int, Langcode_ string, UserPictureTargetId_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("entity_id = ? and langcode = ? and user_picture_target_id = ?", EntityId_, Langcode_, UserPictureTargetId_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByEntityIdAndLangcodeAndUserPictureAlt Get User_userPictures via EntityIdAndLangcodeAndUserPictureAlt
func GetUser_userPicturesByEntityIdAndLangcodeAndUserPictureAlt(offset int, limit int, EntityId_ int, Langcode_ string, UserPictureAlt_ string) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("entity_id = ? and langcode = ? and user_picture_alt = ?", EntityId_, Langcode_, UserPictureAlt_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByEntityIdAndLangcodeAndUserPictureTitle Get User_userPictures via EntityIdAndLangcodeAndUserPictureTitle
func GetUser_userPicturesByEntityIdAndLangcodeAndUserPictureTitle(offset int, limit int, EntityId_ int, Langcode_ string, UserPictureTitle_ string) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("entity_id = ? and langcode = ? and user_picture_title = ?", EntityId_, Langcode_, UserPictureTitle_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByEntityIdAndLangcodeAndUserPictureWidth Get User_userPictures via EntityIdAndLangcodeAndUserPictureWidth
func GetUser_userPicturesByEntityIdAndLangcodeAndUserPictureWidth(offset int, limit int, EntityId_ int, Langcode_ string, UserPictureWidth_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("entity_id = ? and langcode = ? and user_picture_width = ?", EntityId_, Langcode_, UserPictureWidth_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByEntityIdAndLangcodeAndUserPictureHeight Get User_userPictures via EntityIdAndLangcodeAndUserPictureHeight
func GetUser_userPicturesByEntityIdAndLangcodeAndUserPictureHeight(offset int, limit int, EntityId_ int, Langcode_ string, UserPictureHeight_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("entity_id = ? and langcode = ? and user_picture_height = ?", EntityId_, Langcode_, UserPictureHeight_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByEntityIdAndDeltaAndUserPictureTargetId Get User_userPictures via EntityIdAndDeltaAndUserPictureTargetId
func GetUser_userPicturesByEntityIdAndDeltaAndUserPictureTargetId(offset int, limit int, EntityId_ int, Delta_ int, UserPictureTargetId_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("entity_id = ? and delta = ? and user_picture_target_id = ?", EntityId_, Delta_, UserPictureTargetId_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByEntityIdAndDeltaAndUserPictureAlt Get User_userPictures via EntityIdAndDeltaAndUserPictureAlt
func GetUser_userPicturesByEntityIdAndDeltaAndUserPictureAlt(offset int, limit int, EntityId_ int, Delta_ int, UserPictureAlt_ string) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("entity_id = ? and delta = ? and user_picture_alt = ?", EntityId_, Delta_, UserPictureAlt_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByEntityIdAndDeltaAndUserPictureTitle Get User_userPictures via EntityIdAndDeltaAndUserPictureTitle
func GetUser_userPicturesByEntityIdAndDeltaAndUserPictureTitle(offset int, limit int, EntityId_ int, Delta_ int, UserPictureTitle_ string) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("entity_id = ? and delta = ? and user_picture_title = ?", EntityId_, Delta_, UserPictureTitle_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByEntityIdAndDeltaAndUserPictureWidth Get User_userPictures via EntityIdAndDeltaAndUserPictureWidth
func GetUser_userPicturesByEntityIdAndDeltaAndUserPictureWidth(offset int, limit int, EntityId_ int, Delta_ int, UserPictureWidth_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("entity_id = ? and delta = ? and user_picture_width = ?", EntityId_, Delta_, UserPictureWidth_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByEntityIdAndDeltaAndUserPictureHeight Get User_userPictures via EntityIdAndDeltaAndUserPictureHeight
func GetUser_userPicturesByEntityIdAndDeltaAndUserPictureHeight(offset int, limit int, EntityId_ int, Delta_ int, UserPictureHeight_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("entity_id = ? and delta = ? and user_picture_height = ?", EntityId_, Delta_, UserPictureHeight_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByEntityIdAndUserPictureTargetIdAndUserPictureAlt Get User_userPictures via EntityIdAndUserPictureTargetIdAndUserPictureAlt
func GetUser_userPicturesByEntityIdAndUserPictureTargetIdAndUserPictureAlt(offset int, limit int, EntityId_ int, UserPictureTargetId_ int, UserPictureAlt_ string) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("entity_id = ? and user_picture_target_id = ? and user_picture_alt = ?", EntityId_, UserPictureTargetId_, UserPictureAlt_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByEntityIdAndUserPictureTargetIdAndUserPictureTitle Get User_userPictures via EntityIdAndUserPictureTargetIdAndUserPictureTitle
func GetUser_userPicturesByEntityIdAndUserPictureTargetIdAndUserPictureTitle(offset int, limit int, EntityId_ int, UserPictureTargetId_ int, UserPictureTitle_ string) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("entity_id = ? and user_picture_target_id = ? and user_picture_title = ?", EntityId_, UserPictureTargetId_, UserPictureTitle_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByEntityIdAndUserPictureTargetIdAndUserPictureWidth Get User_userPictures via EntityIdAndUserPictureTargetIdAndUserPictureWidth
func GetUser_userPicturesByEntityIdAndUserPictureTargetIdAndUserPictureWidth(offset int, limit int, EntityId_ int, UserPictureTargetId_ int, UserPictureWidth_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("entity_id = ? and user_picture_target_id = ? and user_picture_width = ?", EntityId_, UserPictureTargetId_, UserPictureWidth_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByEntityIdAndUserPictureTargetIdAndUserPictureHeight Get User_userPictures via EntityIdAndUserPictureTargetIdAndUserPictureHeight
func GetUser_userPicturesByEntityIdAndUserPictureTargetIdAndUserPictureHeight(offset int, limit int, EntityId_ int, UserPictureTargetId_ int, UserPictureHeight_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("entity_id = ? and user_picture_target_id = ? and user_picture_height = ?", EntityId_, UserPictureTargetId_, UserPictureHeight_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByEntityIdAndUserPictureAltAndUserPictureTitle Get User_userPictures via EntityIdAndUserPictureAltAndUserPictureTitle
func GetUser_userPicturesByEntityIdAndUserPictureAltAndUserPictureTitle(offset int, limit int, EntityId_ int, UserPictureAlt_ string, UserPictureTitle_ string) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("entity_id = ? and user_picture_alt = ? and user_picture_title = ?", EntityId_, UserPictureAlt_, UserPictureTitle_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByEntityIdAndUserPictureAltAndUserPictureWidth Get User_userPictures via EntityIdAndUserPictureAltAndUserPictureWidth
func GetUser_userPicturesByEntityIdAndUserPictureAltAndUserPictureWidth(offset int, limit int, EntityId_ int, UserPictureAlt_ string, UserPictureWidth_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("entity_id = ? and user_picture_alt = ? and user_picture_width = ?", EntityId_, UserPictureAlt_, UserPictureWidth_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByEntityIdAndUserPictureAltAndUserPictureHeight Get User_userPictures via EntityIdAndUserPictureAltAndUserPictureHeight
func GetUser_userPicturesByEntityIdAndUserPictureAltAndUserPictureHeight(offset int, limit int, EntityId_ int, UserPictureAlt_ string, UserPictureHeight_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("entity_id = ? and user_picture_alt = ? and user_picture_height = ?", EntityId_, UserPictureAlt_, UserPictureHeight_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByEntityIdAndUserPictureTitleAndUserPictureWidth Get User_userPictures via EntityIdAndUserPictureTitleAndUserPictureWidth
func GetUser_userPicturesByEntityIdAndUserPictureTitleAndUserPictureWidth(offset int, limit int, EntityId_ int, UserPictureTitle_ string, UserPictureWidth_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("entity_id = ? and user_picture_title = ? and user_picture_width = ?", EntityId_, UserPictureTitle_, UserPictureWidth_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByEntityIdAndUserPictureTitleAndUserPictureHeight Get User_userPictures via EntityIdAndUserPictureTitleAndUserPictureHeight
func GetUser_userPicturesByEntityIdAndUserPictureTitleAndUserPictureHeight(offset int, limit int, EntityId_ int, UserPictureTitle_ string, UserPictureHeight_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("entity_id = ? and user_picture_title = ? and user_picture_height = ?", EntityId_, UserPictureTitle_, UserPictureHeight_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByEntityIdAndUserPictureWidthAndUserPictureHeight Get User_userPictures via EntityIdAndUserPictureWidthAndUserPictureHeight
func GetUser_userPicturesByEntityIdAndUserPictureWidthAndUserPictureHeight(offset int, limit int, EntityId_ int, UserPictureWidth_ int, UserPictureHeight_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("entity_id = ? and user_picture_width = ? and user_picture_height = ?", EntityId_, UserPictureWidth_, UserPictureHeight_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByRevisionIdAndLangcodeAndDelta Get User_userPictures via RevisionIdAndLangcodeAndDelta
func GetUser_userPicturesByRevisionIdAndLangcodeAndDelta(offset int, limit int, RevisionId_ int, Langcode_ string, Delta_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("revision_id = ? and langcode = ? and delta = ?", RevisionId_, Langcode_, Delta_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByRevisionIdAndLangcodeAndUserPictureTargetId Get User_userPictures via RevisionIdAndLangcodeAndUserPictureTargetId
func GetUser_userPicturesByRevisionIdAndLangcodeAndUserPictureTargetId(offset int, limit int, RevisionId_ int, Langcode_ string, UserPictureTargetId_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("revision_id = ? and langcode = ? and user_picture_target_id = ?", RevisionId_, Langcode_, UserPictureTargetId_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByRevisionIdAndLangcodeAndUserPictureAlt Get User_userPictures via RevisionIdAndLangcodeAndUserPictureAlt
func GetUser_userPicturesByRevisionIdAndLangcodeAndUserPictureAlt(offset int, limit int, RevisionId_ int, Langcode_ string, UserPictureAlt_ string) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("revision_id = ? and langcode = ? and user_picture_alt = ?", RevisionId_, Langcode_, UserPictureAlt_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByRevisionIdAndLangcodeAndUserPictureTitle Get User_userPictures via RevisionIdAndLangcodeAndUserPictureTitle
func GetUser_userPicturesByRevisionIdAndLangcodeAndUserPictureTitle(offset int, limit int, RevisionId_ int, Langcode_ string, UserPictureTitle_ string) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("revision_id = ? and langcode = ? and user_picture_title = ?", RevisionId_, Langcode_, UserPictureTitle_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByRevisionIdAndLangcodeAndUserPictureWidth Get User_userPictures via RevisionIdAndLangcodeAndUserPictureWidth
func GetUser_userPicturesByRevisionIdAndLangcodeAndUserPictureWidth(offset int, limit int, RevisionId_ int, Langcode_ string, UserPictureWidth_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("revision_id = ? and langcode = ? and user_picture_width = ?", RevisionId_, Langcode_, UserPictureWidth_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByRevisionIdAndLangcodeAndUserPictureHeight Get User_userPictures via RevisionIdAndLangcodeAndUserPictureHeight
func GetUser_userPicturesByRevisionIdAndLangcodeAndUserPictureHeight(offset int, limit int, RevisionId_ int, Langcode_ string, UserPictureHeight_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("revision_id = ? and langcode = ? and user_picture_height = ?", RevisionId_, Langcode_, UserPictureHeight_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByRevisionIdAndDeltaAndUserPictureTargetId Get User_userPictures via RevisionIdAndDeltaAndUserPictureTargetId
func GetUser_userPicturesByRevisionIdAndDeltaAndUserPictureTargetId(offset int, limit int, RevisionId_ int, Delta_ int, UserPictureTargetId_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("revision_id = ? and delta = ? and user_picture_target_id = ?", RevisionId_, Delta_, UserPictureTargetId_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByRevisionIdAndDeltaAndUserPictureAlt Get User_userPictures via RevisionIdAndDeltaAndUserPictureAlt
func GetUser_userPicturesByRevisionIdAndDeltaAndUserPictureAlt(offset int, limit int, RevisionId_ int, Delta_ int, UserPictureAlt_ string) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("revision_id = ? and delta = ? and user_picture_alt = ?", RevisionId_, Delta_, UserPictureAlt_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByRevisionIdAndDeltaAndUserPictureTitle Get User_userPictures via RevisionIdAndDeltaAndUserPictureTitle
func GetUser_userPicturesByRevisionIdAndDeltaAndUserPictureTitle(offset int, limit int, RevisionId_ int, Delta_ int, UserPictureTitle_ string) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("revision_id = ? and delta = ? and user_picture_title = ?", RevisionId_, Delta_, UserPictureTitle_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByRevisionIdAndDeltaAndUserPictureWidth Get User_userPictures via RevisionIdAndDeltaAndUserPictureWidth
func GetUser_userPicturesByRevisionIdAndDeltaAndUserPictureWidth(offset int, limit int, RevisionId_ int, Delta_ int, UserPictureWidth_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("revision_id = ? and delta = ? and user_picture_width = ?", RevisionId_, Delta_, UserPictureWidth_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByRevisionIdAndDeltaAndUserPictureHeight Get User_userPictures via RevisionIdAndDeltaAndUserPictureHeight
func GetUser_userPicturesByRevisionIdAndDeltaAndUserPictureHeight(offset int, limit int, RevisionId_ int, Delta_ int, UserPictureHeight_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("revision_id = ? and delta = ? and user_picture_height = ?", RevisionId_, Delta_, UserPictureHeight_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByRevisionIdAndUserPictureTargetIdAndUserPictureAlt Get User_userPictures via RevisionIdAndUserPictureTargetIdAndUserPictureAlt
func GetUser_userPicturesByRevisionIdAndUserPictureTargetIdAndUserPictureAlt(offset int, limit int, RevisionId_ int, UserPictureTargetId_ int, UserPictureAlt_ string) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("revision_id = ? and user_picture_target_id = ? and user_picture_alt = ?", RevisionId_, UserPictureTargetId_, UserPictureAlt_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByRevisionIdAndUserPictureTargetIdAndUserPictureTitle Get User_userPictures via RevisionIdAndUserPictureTargetIdAndUserPictureTitle
func GetUser_userPicturesByRevisionIdAndUserPictureTargetIdAndUserPictureTitle(offset int, limit int, RevisionId_ int, UserPictureTargetId_ int, UserPictureTitle_ string) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("revision_id = ? and user_picture_target_id = ? and user_picture_title = ?", RevisionId_, UserPictureTargetId_, UserPictureTitle_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByRevisionIdAndUserPictureTargetIdAndUserPictureWidth Get User_userPictures via RevisionIdAndUserPictureTargetIdAndUserPictureWidth
func GetUser_userPicturesByRevisionIdAndUserPictureTargetIdAndUserPictureWidth(offset int, limit int, RevisionId_ int, UserPictureTargetId_ int, UserPictureWidth_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("revision_id = ? and user_picture_target_id = ? and user_picture_width = ?", RevisionId_, UserPictureTargetId_, UserPictureWidth_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByRevisionIdAndUserPictureTargetIdAndUserPictureHeight Get User_userPictures via RevisionIdAndUserPictureTargetIdAndUserPictureHeight
func GetUser_userPicturesByRevisionIdAndUserPictureTargetIdAndUserPictureHeight(offset int, limit int, RevisionId_ int, UserPictureTargetId_ int, UserPictureHeight_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("revision_id = ? and user_picture_target_id = ? and user_picture_height = ?", RevisionId_, UserPictureTargetId_, UserPictureHeight_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByRevisionIdAndUserPictureAltAndUserPictureTitle Get User_userPictures via RevisionIdAndUserPictureAltAndUserPictureTitle
func GetUser_userPicturesByRevisionIdAndUserPictureAltAndUserPictureTitle(offset int, limit int, RevisionId_ int, UserPictureAlt_ string, UserPictureTitle_ string) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("revision_id = ? and user_picture_alt = ? and user_picture_title = ?", RevisionId_, UserPictureAlt_, UserPictureTitle_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByRevisionIdAndUserPictureAltAndUserPictureWidth Get User_userPictures via RevisionIdAndUserPictureAltAndUserPictureWidth
func GetUser_userPicturesByRevisionIdAndUserPictureAltAndUserPictureWidth(offset int, limit int, RevisionId_ int, UserPictureAlt_ string, UserPictureWidth_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("revision_id = ? and user_picture_alt = ? and user_picture_width = ?", RevisionId_, UserPictureAlt_, UserPictureWidth_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByRevisionIdAndUserPictureAltAndUserPictureHeight Get User_userPictures via RevisionIdAndUserPictureAltAndUserPictureHeight
func GetUser_userPicturesByRevisionIdAndUserPictureAltAndUserPictureHeight(offset int, limit int, RevisionId_ int, UserPictureAlt_ string, UserPictureHeight_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("revision_id = ? and user_picture_alt = ? and user_picture_height = ?", RevisionId_, UserPictureAlt_, UserPictureHeight_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByRevisionIdAndUserPictureTitleAndUserPictureWidth Get User_userPictures via RevisionIdAndUserPictureTitleAndUserPictureWidth
func GetUser_userPicturesByRevisionIdAndUserPictureTitleAndUserPictureWidth(offset int, limit int, RevisionId_ int, UserPictureTitle_ string, UserPictureWidth_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("revision_id = ? and user_picture_title = ? and user_picture_width = ?", RevisionId_, UserPictureTitle_, UserPictureWidth_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByRevisionIdAndUserPictureTitleAndUserPictureHeight Get User_userPictures via RevisionIdAndUserPictureTitleAndUserPictureHeight
func GetUser_userPicturesByRevisionIdAndUserPictureTitleAndUserPictureHeight(offset int, limit int, RevisionId_ int, UserPictureTitle_ string, UserPictureHeight_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("revision_id = ? and user_picture_title = ? and user_picture_height = ?", RevisionId_, UserPictureTitle_, UserPictureHeight_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByRevisionIdAndUserPictureWidthAndUserPictureHeight Get User_userPictures via RevisionIdAndUserPictureWidthAndUserPictureHeight
func GetUser_userPicturesByRevisionIdAndUserPictureWidthAndUserPictureHeight(offset int, limit int, RevisionId_ int, UserPictureWidth_ int, UserPictureHeight_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("revision_id = ? and user_picture_width = ? and user_picture_height = ?", RevisionId_, UserPictureWidth_, UserPictureHeight_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByLangcodeAndDeltaAndUserPictureTargetId Get User_userPictures via LangcodeAndDeltaAndUserPictureTargetId
func GetUser_userPicturesByLangcodeAndDeltaAndUserPictureTargetId(offset int, limit int, Langcode_ string, Delta_ int, UserPictureTargetId_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("langcode = ? and delta = ? and user_picture_target_id = ?", Langcode_, Delta_, UserPictureTargetId_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByLangcodeAndDeltaAndUserPictureAlt Get User_userPictures via LangcodeAndDeltaAndUserPictureAlt
func GetUser_userPicturesByLangcodeAndDeltaAndUserPictureAlt(offset int, limit int, Langcode_ string, Delta_ int, UserPictureAlt_ string) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("langcode = ? and delta = ? and user_picture_alt = ?", Langcode_, Delta_, UserPictureAlt_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByLangcodeAndDeltaAndUserPictureTitle Get User_userPictures via LangcodeAndDeltaAndUserPictureTitle
func GetUser_userPicturesByLangcodeAndDeltaAndUserPictureTitle(offset int, limit int, Langcode_ string, Delta_ int, UserPictureTitle_ string) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("langcode = ? and delta = ? and user_picture_title = ?", Langcode_, Delta_, UserPictureTitle_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByLangcodeAndDeltaAndUserPictureWidth Get User_userPictures via LangcodeAndDeltaAndUserPictureWidth
func GetUser_userPicturesByLangcodeAndDeltaAndUserPictureWidth(offset int, limit int, Langcode_ string, Delta_ int, UserPictureWidth_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("langcode = ? and delta = ? and user_picture_width = ?", Langcode_, Delta_, UserPictureWidth_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByLangcodeAndDeltaAndUserPictureHeight Get User_userPictures via LangcodeAndDeltaAndUserPictureHeight
func GetUser_userPicturesByLangcodeAndDeltaAndUserPictureHeight(offset int, limit int, Langcode_ string, Delta_ int, UserPictureHeight_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("langcode = ? and delta = ? and user_picture_height = ?", Langcode_, Delta_, UserPictureHeight_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByLangcodeAndUserPictureTargetIdAndUserPictureAlt Get User_userPictures via LangcodeAndUserPictureTargetIdAndUserPictureAlt
func GetUser_userPicturesByLangcodeAndUserPictureTargetIdAndUserPictureAlt(offset int, limit int, Langcode_ string, UserPictureTargetId_ int, UserPictureAlt_ string) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("langcode = ? and user_picture_target_id = ? and user_picture_alt = ?", Langcode_, UserPictureTargetId_, UserPictureAlt_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByLangcodeAndUserPictureTargetIdAndUserPictureTitle Get User_userPictures via LangcodeAndUserPictureTargetIdAndUserPictureTitle
func GetUser_userPicturesByLangcodeAndUserPictureTargetIdAndUserPictureTitle(offset int, limit int, Langcode_ string, UserPictureTargetId_ int, UserPictureTitle_ string) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("langcode = ? and user_picture_target_id = ? and user_picture_title = ?", Langcode_, UserPictureTargetId_, UserPictureTitle_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByLangcodeAndUserPictureTargetIdAndUserPictureWidth Get User_userPictures via LangcodeAndUserPictureTargetIdAndUserPictureWidth
func GetUser_userPicturesByLangcodeAndUserPictureTargetIdAndUserPictureWidth(offset int, limit int, Langcode_ string, UserPictureTargetId_ int, UserPictureWidth_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("langcode = ? and user_picture_target_id = ? and user_picture_width = ?", Langcode_, UserPictureTargetId_, UserPictureWidth_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByLangcodeAndUserPictureTargetIdAndUserPictureHeight Get User_userPictures via LangcodeAndUserPictureTargetIdAndUserPictureHeight
func GetUser_userPicturesByLangcodeAndUserPictureTargetIdAndUserPictureHeight(offset int, limit int, Langcode_ string, UserPictureTargetId_ int, UserPictureHeight_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("langcode = ? and user_picture_target_id = ? and user_picture_height = ?", Langcode_, UserPictureTargetId_, UserPictureHeight_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByLangcodeAndUserPictureAltAndUserPictureTitle Get User_userPictures via LangcodeAndUserPictureAltAndUserPictureTitle
func GetUser_userPicturesByLangcodeAndUserPictureAltAndUserPictureTitle(offset int, limit int, Langcode_ string, UserPictureAlt_ string, UserPictureTitle_ string) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("langcode = ? and user_picture_alt = ? and user_picture_title = ?", Langcode_, UserPictureAlt_, UserPictureTitle_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByLangcodeAndUserPictureAltAndUserPictureWidth Get User_userPictures via LangcodeAndUserPictureAltAndUserPictureWidth
func GetUser_userPicturesByLangcodeAndUserPictureAltAndUserPictureWidth(offset int, limit int, Langcode_ string, UserPictureAlt_ string, UserPictureWidth_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("langcode = ? and user_picture_alt = ? and user_picture_width = ?", Langcode_, UserPictureAlt_, UserPictureWidth_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByLangcodeAndUserPictureAltAndUserPictureHeight Get User_userPictures via LangcodeAndUserPictureAltAndUserPictureHeight
func GetUser_userPicturesByLangcodeAndUserPictureAltAndUserPictureHeight(offset int, limit int, Langcode_ string, UserPictureAlt_ string, UserPictureHeight_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("langcode = ? and user_picture_alt = ? and user_picture_height = ?", Langcode_, UserPictureAlt_, UserPictureHeight_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByLangcodeAndUserPictureTitleAndUserPictureWidth Get User_userPictures via LangcodeAndUserPictureTitleAndUserPictureWidth
func GetUser_userPicturesByLangcodeAndUserPictureTitleAndUserPictureWidth(offset int, limit int, Langcode_ string, UserPictureTitle_ string, UserPictureWidth_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("langcode = ? and user_picture_title = ? and user_picture_width = ?", Langcode_, UserPictureTitle_, UserPictureWidth_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByLangcodeAndUserPictureTitleAndUserPictureHeight Get User_userPictures via LangcodeAndUserPictureTitleAndUserPictureHeight
func GetUser_userPicturesByLangcodeAndUserPictureTitleAndUserPictureHeight(offset int, limit int, Langcode_ string, UserPictureTitle_ string, UserPictureHeight_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("langcode = ? and user_picture_title = ? and user_picture_height = ?", Langcode_, UserPictureTitle_, UserPictureHeight_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByLangcodeAndUserPictureWidthAndUserPictureHeight Get User_userPictures via LangcodeAndUserPictureWidthAndUserPictureHeight
func GetUser_userPicturesByLangcodeAndUserPictureWidthAndUserPictureHeight(offset int, limit int, Langcode_ string, UserPictureWidth_ int, UserPictureHeight_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("langcode = ? and user_picture_width = ? and user_picture_height = ?", Langcode_, UserPictureWidth_, UserPictureHeight_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByDeltaAndUserPictureTargetIdAndUserPictureAlt Get User_userPictures via DeltaAndUserPictureTargetIdAndUserPictureAlt
func GetUser_userPicturesByDeltaAndUserPictureTargetIdAndUserPictureAlt(offset int, limit int, Delta_ int, UserPictureTargetId_ int, UserPictureAlt_ string) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("delta = ? and user_picture_target_id = ? and user_picture_alt = ?", Delta_, UserPictureTargetId_, UserPictureAlt_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByDeltaAndUserPictureTargetIdAndUserPictureTitle Get User_userPictures via DeltaAndUserPictureTargetIdAndUserPictureTitle
func GetUser_userPicturesByDeltaAndUserPictureTargetIdAndUserPictureTitle(offset int, limit int, Delta_ int, UserPictureTargetId_ int, UserPictureTitle_ string) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("delta = ? and user_picture_target_id = ? and user_picture_title = ?", Delta_, UserPictureTargetId_, UserPictureTitle_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByDeltaAndUserPictureTargetIdAndUserPictureWidth Get User_userPictures via DeltaAndUserPictureTargetIdAndUserPictureWidth
func GetUser_userPicturesByDeltaAndUserPictureTargetIdAndUserPictureWidth(offset int, limit int, Delta_ int, UserPictureTargetId_ int, UserPictureWidth_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("delta = ? and user_picture_target_id = ? and user_picture_width = ?", Delta_, UserPictureTargetId_, UserPictureWidth_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByDeltaAndUserPictureTargetIdAndUserPictureHeight Get User_userPictures via DeltaAndUserPictureTargetIdAndUserPictureHeight
func GetUser_userPicturesByDeltaAndUserPictureTargetIdAndUserPictureHeight(offset int, limit int, Delta_ int, UserPictureTargetId_ int, UserPictureHeight_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("delta = ? and user_picture_target_id = ? and user_picture_height = ?", Delta_, UserPictureTargetId_, UserPictureHeight_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByDeltaAndUserPictureAltAndUserPictureTitle Get User_userPictures via DeltaAndUserPictureAltAndUserPictureTitle
func GetUser_userPicturesByDeltaAndUserPictureAltAndUserPictureTitle(offset int, limit int, Delta_ int, UserPictureAlt_ string, UserPictureTitle_ string) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("delta = ? and user_picture_alt = ? and user_picture_title = ?", Delta_, UserPictureAlt_, UserPictureTitle_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByDeltaAndUserPictureAltAndUserPictureWidth Get User_userPictures via DeltaAndUserPictureAltAndUserPictureWidth
func GetUser_userPicturesByDeltaAndUserPictureAltAndUserPictureWidth(offset int, limit int, Delta_ int, UserPictureAlt_ string, UserPictureWidth_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("delta = ? and user_picture_alt = ? and user_picture_width = ?", Delta_, UserPictureAlt_, UserPictureWidth_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByDeltaAndUserPictureAltAndUserPictureHeight Get User_userPictures via DeltaAndUserPictureAltAndUserPictureHeight
func GetUser_userPicturesByDeltaAndUserPictureAltAndUserPictureHeight(offset int, limit int, Delta_ int, UserPictureAlt_ string, UserPictureHeight_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("delta = ? and user_picture_alt = ? and user_picture_height = ?", Delta_, UserPictureAlt_, UserPictureHeight_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByDeltaAndUserPictureTitleAndUserPictureWidth Get User_userPictures via DeltaAndUserPictureTitleAndUserPictureWidth
func GetUser_userPicturesByDeltaAndUserPictureTitleAndUserPictureWidth(offset int, limit int, Delta_ int, UserPictureTitle_ string, UserPictureWidth_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("delta = ? and user_picture_title = ? and user_picture_width = ?", Delta_, UserPictureTitle_, UserPictureWidth_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByDeltaAndUserPictureTitleAndUserPictureHeight Get User_userPictures via DeltaAndUserPictureTitleAndUserPictureHeight
func GetUser_userPicturesByDeltaAndUserPictureTitleAndUserPictureHeight(offset int, limit int, Delta_ int, UserPictureTitle_ string, UserPictureHeight_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("delta = ? and user_picture_title = ? and user_picture_height = ?", Delta_, UserPictureTitle_, UserPictureHeight_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByDeltaAndUserPictureWidthAndUserPictureHeight Get User_userPictures via DeltaAndUserPictureWidthAndUserPictureHeight
func GetUser_userPicturesByDeltaAndUserPictureWidthAndUserPictureHeight(offset int, limit int, Delta_ int, UserPictureWidth_ int, UserPictureHeight_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("delta = ? and user_picture_width = ? and user_picture_height = ?", Delta_, UserPictureWidth_, UserPictureHeight_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByUserPictureTargetIdAndUserPictureAltAndUserPictureTitle Get User_userPictures via UserPictureTargetIdAndUserPictureAltAndUserPictureTitle
func GetUser_userPicturesByUserPictureTargetIdAndUserPictureAltAndUserPictureTitle(offset int, limit int, UserPictureTargetId_ int, UserPictureAlt_ string, UserPictureTitle_ string) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("user_picture_target_id = ? and user_picture_alt = ? and user_picture_title = ?", UserPictureTargetId_, UserPictureAlt_, UserPictureTitle_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByUserPictureTargetIdAndUserPictureAltAndUserPictureWidth Get User_userPictures via UserPictureTargetIdAndUserPictureAltAndUserPictureWidth
func GetUser_userPicturesByUserPictureTargetIdAndUserPictureAltAndUserPictureWidth(offset int, limit int, UserPictureTargetId_ int, UserPictureAlt_ string, UserPictureWidth_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("user_picture_target_id = ? and user_picture_alt = ? and user_picture_width = ?", UserPictureTargetId_, UserPictureAlt_, UserPictureWidth_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByUserPictureTargetIdAndUserPictureAltAndUserPictureHeight Get User_userPictures via UserPictureTargetIdAndUserPictureAltAndUserPictureHeight
func GetUser_userPicturesByUserPictureTargetIdAndUserPictureAltAndUserPictureHeight(offset int, limit int, UserPictureTargetId_ int, UserPictureAlt_ string, UserPictureHeight_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("user_picture_target_id = ? and user_picture_alt = ? and user_picture_height = ?", UserPictureTargetId_, UserPictureAlt_, UserPictureHeight_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByUserPictureTargetIdAndUserPictureTitleAndUserPictureWidth Get User_userPictures via UserPictureTargetIdAndUserPictureTitleAndUserPictureWidth
func GetUser_userPicturesByUserPictureTargetIdAndUserPictureTitleAndUserPictureWidth(offset int, limit int, UserPictureTargetId_ int, UserPictureTitle_ string, UserPictureWidth_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("user_picture_target_id = ? and user_picture_title = ? and user_picture_width = ?", UserPictureTargetId_, UserPictureTitle_, UserPictureWidth_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByUserPictureTargetIdAndUserPictureTitleAndUserPictureHeight Get User_userPictures via UserPictureTargetIdAndUserPictureTitleAndUserPictureHeight
func GetUser_userPicturesByUserPictureTargetIdAndUserPictureTitleAndUserPictureHeight(offset int, limit int, UserPictureTargetId_ int, UserPictureTitle_ string, UserPictureHeight_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("user_picture_target_id = ? and user_picture_title = ? and user_picture_height = ?", UserPictureTargetId_, UserPictureTitle_, UserPictureHeight_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByUserPictureTargetIdAndUserPictureWidthAndUserPictureHeight Get User_userPictures via UserPictureTargetIdAndUserPictureWidthAndUserPictureHeight
func GetUser_userPicturesByUserPictureTargetIdAndUserPictureWidthAndUserPictureHeight(offset int, limit int, UserPictureTargetId_ int, UserPictureWidth_ int, UserPictureHeight_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("user_picture_target_id = ? and user_picture_width = ? and user_picture_height = ?", UserPictureTargetId_, UserPictureWidth_, UserPictureHeight_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByUserPictureAltAndUserPictureTitleAndUserPictureWidth Get User_userPictures via UserPictureAltAndUserPictureTitleAndUserPictureWidth
func GetUser_userPicturesByUserPictureAltAndUserPictureTitleAndUserPictureWidth(offset int, limit int, UserPictureAlt_ string, UserPictureTitle_ string, UserPictureWidth_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("user_picture_alt = ? and user_picture_title = ? and user_picture_width = ?", UserPictureAlt_, UserPictureTitle_, UserPictureWidth_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByUserPictureAltAndUserPictureTitleAndUserPictureHeight Get User_userPictures via UserPictureAltAndUserPictureTitleAndUserPictureHeight
func GetUser_userPicturesByUserPictureAltAndUserPictureTitleAndUserPictureHeight(offset int, limit int, UserPictureAlt_ string, UserPictureTitle_ string, UserPictureHeight_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("user_picture_alt = ? and user_picture_title = ? and user_picture_height = ?", UserPictureAlt_, UserPictureTitle_, UserPictureHeight_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByUserPictureAltAndUserPictureWidthAndUserPictureHeight Get User_userPictures via UserPictureAltAndUserPictureWidthAndUserPictureHeight
func GetUser_userPicturesByUserPictureAltAndUserPictureWidthAndUserPictureHeight(offset int, limit int, UserPictureAlt_ string, UserPictureWidth_ int, UserPictureHeight_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("user_picture_alt = ? and user_picture_width = ? and user_picture_height = ?", UserPictureAlt_, UserPictureWidth_, UserPictureHeight_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByUserPictureTitleAndUserPictureWidthAndUserPictureHeight Get User_userPictures via UserPictureTitleAndUserPictureWidthAndUserPictureHeight
func GetUser_userPicturesByUserPictureTitleAndUserPictureWidthAndUserPictureHeight(offset int, limit int, UserPictureTitle_ string, UserPictureWidth_ int, UserPictureHeight_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("user_picture_title = ? and user_picture_width = ? and user_picture_height = ?", UserPictureTitle_, UserPictureWidth_, UserPictureHeight_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByBundleAndDeleted Get User_userPictures via BundleAndDeleted
func GetUser_userPicturesByBundleAndDeleted(offset int, limit int, Bundle_ string, Deleted_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("bundle = ? and deleted = ?", Bundle_, Deleted_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByBundleAndEntityId Get User_userPictures via BundleAndEntityId
func GetUser_userPicturesByBundleAndEntityId(offset int, limit int, Bundle_ string, EntityId_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("bundle = ? and entity_id = ?", Bundle_, EntityId_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByBundleAndRevisionId Get User_userPictures via BundleAndRevisionId
func GetUser_userPicturesByBundleAndRevisionId(offset int, limit int, Bundle_ string, RevisionId_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("bundle = ? and revision_id = ?", Bundle_, RevisionId_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByBundleAndLangcode Get User_userPictures via BundleAndLangcode
func GetUser_userPicturesByBundleAndLangcode(offset int, limit int, Bundle_ string, Langcode_ string) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("bundle = ? and langcode = ?", Bundle_, Langcode_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByBundleAndDelta Get User_userPictures via BundleAndDelta
func GetUser_userPicturesByBundleAndDelta(offset int, limit int, Bundle_ string, Delta_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("bundle = ? and delta = ?", Bundle_, Delta_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByBundleAndUserPictureTargetId Get User_userPictures via BundleAndUserPictureTargetId
func GetUser_userPicturesByBundleAndUserPictureTargetId(offset int, limit int, Bundle_ string, UserPictureTargetId_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("bundle = ? and user_picture_target_id = ?", Bundle_, UserPictureTargetId_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByBundleAndUserPictureAlt Get User_userPictures via BundleAndUserPictureAlt
func GetUser_userPicturesByBundleAndUserPictureAlt(offset int, limit int, Bundle_ string, UserPictureAlt_ string) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("bundle = ? and user_picture_alt = ?", Bundle_, UserPictureAlt_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByBundleAndUserPictureTitle Get User_userPictures via BundleAndUserPictureTitle
func GetUser_userPicturesByBundleAndUserPictureTitle(offset int, limit int, Bundle_ string, UserPictureTitle_ string) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("bundle = ? and user_picture_title = ?", Bundle_, UserPictureTitle_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByBundleAndUserPictureWidth Get User_userPictures via BundleAndUserPictureWidth
func GetUser_userPicturesByBundleAndUserPictureWidth(offset int, limit int, Bundle_ string, UserPictureWidth_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("bundle = ? and user_picture_width = ?", Bundle_, UserPictureWidth_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByBundleAndUserPictureHeight Get User_userPictures via BundleAndUserPictureHeight
func GetUser_userPicturesByBundleAndUserPictureHeight(offset int, limit int, Bundle_ string, UserPictureHeight_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("bundle = ? and user_picture_height = ?", Bundle_, UserPictureHeight_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByDeletedAndEntityId Get User_userPictures via DeletedAndEntityId
func GetUser_userPicturesByDeletedAndEntityId(offset int, limit int, Deleted_ int, EntityId_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("deleted = ? and entity_id = ?", Deleted_, EntityId_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByDeletedAndRevisionId Get User_userPictures via DeletedAndRevisionId
func GetUser_userPicturesByDeletedAndRevisionId(offset int, limit int, Deleted_ int, RevisionId_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("deleted = ? and revision_id = ?", Deleted_, RevisionId_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByDeletedAndLangcode Get User_userPictures via DeletedAndLangcode
func GetUser_userPicturesByDeletedAndLangcode(offset int, limit int, Deleted_ int, Langcode_ string) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("deleted = ? and langcode = ?", Deleted_, Langcode_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByDeletedAndDelta Get User_userPictures via DeletedAndDelta
func GetUser_userPicturesByDeletedAndDelta(offset int, limit int, Deleted_ int, Delta_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("deleted = ? and delta = ?", Deleted_, Delta_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByDeletedAndUserPictureTargetId Get User_userPictures via DeletedAndUserPictureTargetId
func GetUser_userPicturesByDeletedAndUserPictureTargetId(offset int, limit int, Deleted_ int, UserPictureTargetId_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("deleted = ? and user_picture_target_id = ?", Deleted_, UserPictureTargetId_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByDeletedAndUserPictureAlt Get User_userPictures via DeletedAndUserPictureAlt
func GetUser_userPicturesByDeletedAndUserPictureAlt(offset int, limit int, Deleted_ int, UserPictureAlt_ string) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("deleted = ? and user_picture_alt = ?", Deleted_, UserPictureAlt_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByDeletedAndUserPictureTitle Get User_userPictures via DeletedAndUserPictureTitle
func GetUser_userPicturesByDeletedAndUserPictureTitle(offset int, limit int, Deleted_ int, UserPictureTitle_ string) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("deleted = ? and user_picture_title = ?", Deleted_, UserPictureTitle_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByDeletedAndUserPictureWidth Get User_userPictures via DeletedAndUserPictureWidth
func GetUser_userPicturesByDeletedAndUserPictureWidth(offset int, limit int, Deleted_ int, UserPictureWidth_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("deleted = ? and user_picture_width = ?", Deleted_, UserPictureWidth_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByDeletedAndUserPictureHeight Get User_userPictures via DeletedAndUserPictureHeight
func GetUser_userPicturesByDeletedAndUserPictureHeight(offset int, limit int, Deleted_ int, UserPictureHeight_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("deleted = ? and user_picture_height = ?", Deleted_, UserPictureHeight_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByEntityIdAndRevisionId Get User_userPictures via EntityIdAndRevisionId
func GetUser_userPicturesByEntityIdAndRevisionId(offset int, limit int, EntityId_ int, RevisionId_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("entity_id = ? and revision_id = ?", EntityId_, RevisionId_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByEntityIdAndLangcode Get User_userPictures via EntityIdAndLangcode
func GetUser_userPicturesByEntityIdAndLangcode(offset int, limit int, EntityId_ int, Langcode_ string) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("entity_id = ? and langcode = ?", EntityId_, Langcode_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByEntityIdAndDelta Get User_userPictures via EntityIdAndDelta
func GetUser_userPicturesByEntityIdAndDelta(offset int, limit int, EntityId_ int, Delta_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("entity_id = ? and delta = ?", EntityId_, Delta_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByEntityIdAndUserPictureTargetId Get User_userPictures via EntityIdAndUserPictureTargetId
func GetUser_userPicturesByEntityIdAndUserPictureTargetId(offset int, limit int, EntityId_ int, UserPictureTargetId_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("entity_id = ? and user_picture_target_id = ?", EntityId_, UserPictureTargetId_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByEntityIdAndUserPictureAlt Get User_userPictures via EntityIdAndUserPictureAlt
func GetUser_userPicturesByEntityIdAndUserPictureAlt(offset int, limit int, EntityId_ int, UserPictureAlt_ string) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("entity_id = ? and user_picture_alt = ?", EntityId_, UserPictureAlt_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByEntityIdAndUserPictureTitle Get User_userPictures via EntityIdAndUserPictureTitle
func GetUser_userPicturesByEntityIdAndUserPictureTitle(offset int, limit int, EntityId_ int, UserPictureTitle_ string) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("entity_id = ? and user_picture_title = ?", EntityId_, UserPictureTitle_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByEntityIdAndUserPictureWidth Get User_userPictures via EntityIdAndUserPictureWidth
func GetUser_userPicturesByEntityIdAndUserPictureWidth(offset int, limit int, EntityId_ int, UserPictureWidth_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("entity_id = ? and user_picture_width = ?", EntityId_, UserPictureWidth_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByEntityIdAndUserPictureHeight Get User_userPictures via EntityIdAndUserPictureHeight
func GetUser_userPicturesByEntityIdAndUserPictureHeight(offset int, limit int, EntityId_ int, UserPictureHeight_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("entity_id = ? and user_picture_height = ?", EntityId_, UserPictureHeight_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByRevisionIdAndLangcode Get User_userPictures via RevisionIdAndLangcode
func GetUser_userPicturesByRevisionIdAndLangcode(offset int, limit int, RevisionId_ int, Langcode_ string) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("revision_id = ? and langcode = ?", RevisionId_, Langcode_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByRevisionIdAndDelta Get User_userPictures via RevisionIdAndDelta
func GetUser_userPicturesByRevisionIdAndDelta(offset int, limit int, RevisionId_ int, Delta_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("revision_id = ? and delta = ?", RevisionId_, Delta_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByRevisionIdAndUserPictureTargetId Get User_userPictures via RevisionIdAndUserPictureTargetId
func GetUser_userPicturesByRevisionIdAndUserPictureTargetId(offset int, limit int, RevisionId_ int, UserPictureTargetId_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("revision_id = ? and user_picture_target_id = ?", RevisionId_, UserPictureTargetId_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByRevisionIdAndUserPictureAlt Get User_userPictures via RevisionIdAndUserPictureAlt
func GetUser_userPicturesByRevisionIdAndUserPictureAlt(offset int, limit int, RevisionId_ int, UserPictureAlt_ string) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("revision_id = ? and user_picture_alt = ?", RevisionId_, UserPictureAlt_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByRevisionIdAndUserPictureTitle Get User_userPictures via RevisionIdAndUserPictureTitle
func GetUser_userPicturesByRevisionIdAndUserPictureTitle(offset int, limit int, RevisionId_ int, UserPictureTitle_ string) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("revision_id = ? and user_picture_title = ?", RevisionId_, UserPictureTitle_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByRevisionIdAndUserPictureWidth Get User_userPictures via RevisionIdAndUserPictureWidth
func GetUser_userPicturesByRevisionIdAndUserPictureWidth(offset int, limit int, RevisionId_ int, UserPictureWidth_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("revision_id = ? and user_picture_width = ?", RevisionId_, UserPictureWidth_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByRevisionIdAndUserPictureHeight Get User_userPictures via RevisionIdAndUserPictureHeight
func GetUser_userPicturesByRevisionIdAndUserPictureHeight(offset int, limit int, RevisionId_ int, UserPictureHeight_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("revision_id = ? and user_picture_height = ?", RevisionId_, UserPictureHeight_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByLangcodeAndDelta Get User_userPictures via LangcodeAndDelta
func GetUser_userPicturesByLangcodeAndDelta(offset int, limit int, Langcode_ string, Delta_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("langcode = ? and delta = ?", Langcode_, Delta_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByLangcodeAndUserPictureTargetId Get User_userPictures via LangcodeAndUserPictureTargetId
func GetUser_userPicturesByLangcodeAndUserPictureTargetId(offset int, limit int, Langcode_ string, UserPictureTargetId_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("langcode = ? and user_picture_target_id = ?", Langcode_, UserPictureTargetId_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByLangcodeAndUserPictureAlt Get User_userPictures via LangcodeAndUserPictureAlt
func GetUser_userPicturesByLangcodeAndUserPictureAlt(offset int, limit int, Langcode_ string, UserPictureAlt_ string) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("langcode = ? and user_picture_alt = ?", Langcode_, UserPictureAlt_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByLangcodeAndUserPictureTitle Get User_userPictures via LangcodeAndUserPictureTitle
func GetUser_userPicturesByLangcodeAndUserPictureTitle(offset int, limit int, Langcode_ string, UserPictureTitle_ string) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("langcode = ? and user_picture_title = ?", Langcode_, UserPictureTitle_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByLangcodeAndUserPictureWidth Get User_userPictures via LangcodeAndUserPictureWidth
func GetUser_userPicturesByLangcodeAndUserPictureWidth(offset int, limit int, Langcode_ string, UserPictureWidth_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("langcode = ? and user_picture_width = ?", Langcode_, UserPictureWidth_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByLangcodeAndUserPictureHeight Get User_userPictures via LangcodeAndUserPictureHeight
func GetUser_userPicturesByLangcodeAndUserPictureHeight(offset int, limit int, Langcode_ string, UserPictureHeight_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("langcode = ? and user_picture_height = ?", Langcode_, UserPictureHeight_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByDeltaAndUserPictureTargetId Get User_userPictures via DeltaAndUserPictureTargetId
func GetUser_userPicturesByDeltaAndUserPictureTargetId(offset int, limit int, Delta_ int, UserPictureTargetId_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("delta = ? and user_picture_target_id = ?", Delta_, UserPictureTargetId_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByDeltaAndUserPictureAlt Get User_userPictures via DeltaAndUserPictureAlt
func GetUser_userPicturesByDeltaAndUserPictureAlt(offset int, limit int, Delta_ int, UserPictureAlt_ string) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("delta = ? and user_picture_alt = ?", Delta_, UserPictureAlt_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByDeltaAndUserPictureTitle Get User_userPictures via DeltaAndUserPictureTitle
func GetUser_userPicturesByDeltaAndUserPictureTitle(offset int, limit int, Delta_ int, UserPictureTitle_ string) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("delta = ? and user_picture_title = ?", Delta_, UserPictureTitle_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByDeltaAndUserPictureWidth Get User_userPictures via DeltaAndUserPictureWidth
func GetUser_userPicturesByDeltaAndUserPictureWidth(offset int, limit int, Delta_ int, UserPictureWidth_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("delta = ? and user_picture_width = ?", Delta_, UserPictureWidth_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByDeltaAndUserPictureHeight Get User_userPictures via DeltaAndUserPictureHeight
func GetUser_userPicturesByDeltaAndUserPictureHeight(offset int, limit int, Delta_ int, UserPictureHeight_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("delta = ? and user_picture_height = ?", Delta_, UserPictureHeight_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByUserPictureTargetIdAndUserPictureAlt Get User_userPictures via UserPictureTargetIdAndUserPictureAlt
func GetUser_userPicturesByUserPictureTargetIdAndUserPictureAlt(offset int, limit int, UserPictureTargetId_ int, UserPictureAlt_ string) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("user_picture_target_id = ? and user_picture_alt = ?", UserPictureTargetId_, UserPictureAlt_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByUserPictureTargetIdAndUserPictureTitle Get User_userPictures via UserPictureTargetIdAndUserPictureTitle
func GetUser_userPicturesByUserPictureTargetIdAndUserPictureTitle(offset int, limit int, UserPictureTargetId_ int, UserPictureTitle_ string) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("user_picture_target_id = ? and user_picture_title = ?", UserPictureTargetId_, UserPictureTitle_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByUserPictureTargetIdAndUserPictureWidth Get User_userPictures via UserPictureTargetIdAndUserPictureWidth
func GetUser_userPicturesByUserPictureTargetIdAndUserPictureWidth(offset int, limit int, UserPictureTargetId_ int, UserPictureWidth_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("user_picture_target_id = ? and user_picture_width = ?", UserPictureTargetId_, UserPictureWidth_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByUserPictureTargetIdAndUserPictureHeight Get User_userPictures via UserPictureTargetIdAndUserPictureHeight
func GetUser_userPicturesByUserPictureTargetIdAndUserPictureHeight(offset int, limit int, UserPictureTargetId_ int, UserPictureHeight_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("user_picture_target_id = ? and user_picture_height = ?", UserPictureTargetId_, UserPictureHeight_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByUserPictureAltAndUserPictureTitle Get User_userPictures via UserPictureAltAndUserPictureTitle
func GetUser_userPicturesByUserPictureAltAndUserPictureTitle(offset int, limit int, UserPictureAlt_ string, UserPictureTitle_ string) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("user_picture_alt = ? and user_picture_title = ?", UserPictureAlt_, UserPictureTitle_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByUserPictureAltAndUserPictureWidth Get User_userPictures via UserPictureAltAndUserPictureWidth
func GetUser_userPicturesByUserPictureAltAndUserPictureWidth(offset int, limit int, UserPictureAlt_ string, UserPictureWidth_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("user_picture_alt = ? and user_picture_width = ?", UserPictureAlt_, UserPictureWidth_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByUserPictureAltAndUserPictureHeight Get User_userPictures via UserPictureAltAndUserPictureHeight
func GetUser_userPicturesByUserPictureAltAndUserPictureHeight(offset int, limit int, UserPictureAlt_ string, UserPictureHeight_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("user_picture_alt = ? and user_picture_height = ?", UserPictureAlt_, UserPictureHeight_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByUserPictureTitleAndUserPictureWidth Get User_userPictures via UserPictureTitleAndUserPictureWidth
func GetUser_userPicturesByUserPictureTitleAndUserPictureWidth(offset int, limit int, UserPictureTitle_ string, UserPictureWidth_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("user_picture_title = ? and user_picture_width = ?", UserPictureTitle_, UserPictureWidth_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByUserPictureTitleAndUserPictureHeight Get User_userPictures via UserPictureTitleAndUserPictureHeight
func GetUser_userPicturesByUserPictureTitleAndUserPictureHeight(offset int, limit int, UserPictureTitle_ string, UserPictureHeight_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("user_picture_title = ? and user_picture_height = ?", UserPictureTitle_, UserPictureHeight_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesByUserPictureWidthAndUserPictureHeight Get User_userPictures via UserPictureWidthAndUserPictureHeight
func GetUser_userPicturesByUserPictureWidthAndUserPictureHeight(offset int, limit int, UserPictureWidth_ int, UserPictureHeight_ int) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("user_picture_width = ? and user_picture_height = ?", UserPictureWidth_, UserPictureHeight_).Limit(limit, offset).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPictures Get User_userPictures via field
func GetUser_userPictures(offset int, limit int, field string) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Limit(limit, offset).Desc(field).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesViaBundle Get User_userPictures via Bundle
func GetUser_userPicturesViaBundle(offset int, limit int, Bundle_ string, field string) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("bundle = ?", Bundle_).Limit(limit, offset).Desc(field).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesViaDeleted Get User_userPictures via Deleted
func GetUser_userPicturesViaDeleted(offset int, limit int, Deleted_ int, field string) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("deleted = ?", Deleted_).Limit(limit, offset).Desc(field).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesViaEntityId Get User_userPictures via EntityId
func GetUser_userPicturesViaEntityId(offset int, limit int, EntityId_ int, field string) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("entity_id = ?", EntityId_).Limit(limit, offset).Desc(field).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesViaRevisionId Get User_userPictures via RevisionId
func GetUser_userPicturesViaRevisionId(offset int, limit int, RevisionId_ int, field string) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("revision_id = ?", RevisionId_).Limit(limit, offset).Desc(field).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesViaLangcode Get User_userPictures via Langcode
func GetUser_userPicturesViaLangcode(offset int, limit int, Langcode_ string, field string) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("langcode = ?", Langcode_).Limit(limit, offset).Desc(field).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesViaDelta Get User_userPictures via Delta
func GetUser_userPicturesViaDelta(offset int, limit int, Delta_ int, field string) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("delta = ?", Delta_).Limit(limit, offset).Desc(field).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesViaUserPictureTargetId Get User_userPictures via UserPictureTargetId
func GetUser_userPicturesViaUserPictureTargetId(offset int, limit int, UserPictureTargetId_ int, field string) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("user_picture_target_id = ?", UserPictureTargetId_).Limit(limit, offset).Desc(field).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesViaUserPictureAlt Get User_userPictures via UserPictureAlt
func GetUser_userPicturesViaUserPictureAlt(offset int, limit int, UserPictureAlt_ string, field string) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("user_picture_alt = ?", UserPictureAlt_).Limit(limit, offset).Desc(field).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesViaUserPictureTitle Get User_userPictures via UserPictureTitle
func GetUser_userPicturesViaUserPictureTitle(offset int, limit int, UserPictureTitle_ string, field string) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("user_picture_title = ?", UserPictureTitle_).Limit(limit, offset).Desc(field).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesViaUserPictureWidth Get User_userPictures via UserPictureWidth
func GetUser_userPicturesViaUserPictureWidth(offset int, limit int, UserPictureWidth_ int, field string) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("user_picture_width = ?", UserPictureWidth_).Limit(limit, offset).Desc(field).Find(_User_userPicture)
	return _User_userPicture, err
}

// GetUser_userPicturesViaUserPictureHeight Get User_userPictures via UserPictureHeight
func GetUser_userPicturesViaUserPictureHeight(offset int, limit int, UserPictureHeight_ int, field string) (*[]*User_userPicture, error) {
	var _User_userPicture = new([]*User_userPicture)
	err := Engine.Table("user__user_picture").Where("user_picture_height = ?", UserPictureHeight_).Limit(limit, offset).Desc(field).Find(_User_userPicture)
	return _User_userPicture, err
}

// HasUser_userPictureViaBundle Has User_userPicture via Bundle
func HasUser_userPictureViaBundle(iBundle string) bool {
	if has, err := Engine.Where("bundle = ?", iBundle).Get(new(User_userPicture)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasUser_userPictureViaDeleted Has User_userPicture via Deleted
func HasUser_userPictureViaDeleted(iDeleted int) bool {
	if has, err := Engine.Where("deleted = ?", iDeleted).Get(new(User_userPicture)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasUser_userPictureViaEntityId Has User_userPicture via EntityId
func HasUser_userPictureViaEntityId(iEntityId int) bool {
	if has, err := Engine.Where("entity_id = ?", iEntityId).Get(new(User_userPicture)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasUser_userPictureViaRevisionId Has User_userPicture via RevisionId
func HasUser_userPictureViaRevisionId(iRevisionId int) bool {
	if has, err := Engine.Where("revision_id = ?", iRevisionId).Get(new(User_userPicture)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasUser_userPictureViaLangcode Has User_userPicture via Langcode
func HasUser_userPictureViaLangcode(iLangcode string) bool {
	if has, err := Engine.Where("langcode = ?", iLangcode).Get(new(User_userPicture)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasUser_userPictureViaDelta Has User_userPicture via Delta
func HasUser_userPictureViaDelta(iDelta int) bool {
	if has, err := Engine.Where("delta = ?", iDelta).Get(new(User_userPicture)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasUser_userPictureViaUserPictureTargetId Has User_userPicture via UserPictureTargetId
func HasUser_userPictureViaUserPictureTargetId(iUserPictureTargetId int) bool {
	if has, err := Engine.Where("user_picture_target_id = ?", iUserPictureTargetId).Get(new(User_userPicture)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasUser_userPictureViaUserPictureAlt Has User_userPicture via UserPictureAlt
func HasUser_userPictureViaUserPictureAlt(iUserPictureAlt string) bool {
	if has, err := Engine.Where("user_picture_alt = ?", iUserPictureAlt).Get(new(User_userPicture)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasUser_userPictureViaUserPictureTitle Has User_userPicture via UserPictureTitle
func HasUser_userPictureViaUserPictureTitle(iUserPictureTitle string) bool {
	if has, err := Engine.Where("user_picture_title = ?", iUserPictureTitle).Get(new(User_userPicture)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasUser_userPictureViaUserPictureWidth Has User_userPicture via UserPictureWidth
func HasUser_userPictureViaUserPictureWidth(iUserPictureWidth int) bool {
	if has, err := Engine.Where("user_picture_width = ?", iUserPictureWidth).Get(new(User_userPicture)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasUser_userPictureViaUserPictureHeight Has User_userPicture via UserPictureHeight
func HasUser_userPictureViaUserPictureHeight(iUserPictureHeight int) bool {
	if has, err := Engine.Where("user_picture_height = ?", iUserPictureHeight).Get(new(User_userPicture)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// GetUser_userPictureViaBundle Get User_userPicture via Bundle
func GetUser_userPictureViaBundle(iBundle string) (*User_userPicture, error) {
	var _User_userPicture = &User_userPicture{Bundle: iBundle}
	has, err := Engine.Get(_User_userPicture)
	if has {
		return _User_userPicture, err
	} else {
		return nil, err
	}
}

// GetUser_userPictureViaDeleted Get User_userPicture via Deleted
func GetUser_userPictureViaDeleted(iDeleted int) (*User_userPicture, error) {
	var _User_userPicture = &User_userPicture{Deleted: iDeleted}
	has, err := Engine.Get(_User_userPicture)
	if has {
		return _User_userPicture, err
	} else {
		return nil, err
	}
}

// GetUser_userPictureViaEntityId Get User_userPicture via EntityId
func GetUser_userPictureViaEntityId(iEntityId int) (*User_userPicture, error) {
	var _User_userPicture = &User_userPicture{EntityId: iEntityId}
	has, err := Engine.Get(_User_userPicture)
	if has {
		return _User_userPicture, err
	} else {
		return nil, err
	}
}

// GetUser_userPictureViaRevisionId Get User_userPicture via RevisionId
func GetUser_userPictureViaRevisionId(iRevisionId int) (*User_userPicture, error) {
	var _User_userPicture = &User_userPicture{RevisionId: iRevisionId}
	has, err := Engine.Get(_User_userPicture)
	if has {
		return _User_userPicture, err
	} else {
		return nil, err
	}
}

// GetUser_userPictureViaLangcode Get User_userPicture via Langcode
func GetUser_userPictureViaLangcode(iLangcode string) (*User_userPicture, error) {
	var _User_userPicture = &User_userPicture{Langcode: iLangcode}
	has, err := Engine.Get(_User_userPicture)
	if has {
		return _User_userPicture, err
	} else {
		return nil, err
	}
}

// GetUser_userPictureViaDelta Get User_userPicture via Delta
func GetUser_userPictureViaDelta(iDelta int) (*User_userPicture, error) {
	var _User_userPicture = &User_userPicture{Delta: iDelta}
	has, err := Engine.Get(_User_userPicture)
	if has {
		return _User_userPicture, err
	} else {
		return nil, err
	}
}

// GetUser_userPictureViaUserPictureTargetId Get User_userPicture via UserPictureTargetId
func GetUser_userPictureViaUserPictureTargetId(iUserPictureTargetId int) (*User_userPicture, error) {
	var _User_userPicture = &User_userPicture{UserPictureTargetId: iUserPictureTargetId}
	has, err := Engine.Get(_User_userPicture)
	if has {
		return _User_userPicture, err
	} else {
		return nil, err
	}
}

// GetUser_userPictureViaUserPictureAlt Get User_userPicture via UserPictureAlt
func GetUser_userPictureViaUserPictureAlt(iUserPictureAlt string) (*User_userPicture, error) {
	var _User_userPicture = &User_userPicture{UserPictureAlt: iUserPictureAlt}
	has, err := Engine.Get(_User_userPicture)
	if has {
		return _User_userPicture, err
	} else {
		return nil, err
	}
}

// GetUser_userPictureViaUserPictureTitle Get User_userPicture via UserPictureTitle
func GetUser_userPictureViaUserPictureTitle(iUserPictureTitle string) (*User_userPicture, error) {
	var _User_userPicture = &User_userPicture{UserPictureTitle: iUserPictureTitle}
	has, err := Engine.Get(_User_userPicture)
	if has {
		return _User_userPicture, err
	} else {
		return nil, err
	}
}

// GetUser_userPictureViaUserPictureWidth Get User_userPicture via UserPictureWidth
func GetUser_userPictureViaUserPictureWidth(iUserPictureWidth int) (*User_userPicture, error) {
	var _User_userPicture = &User_userPicture{UserPictureWidth: iUserPictureWidth}
	has, err := Engine.Get(_User_userPicture)
	if has {
		return _User_userPicture, err
	} else {
		return nil, err
	}
}

// GetUser_userPictureViaUserPictureHeight Get User_userPicture via UserPictureHeight
func GetUser_userPictureViaUserPictureHeight(iUserPictureHeight int) (*User_userPicture, error) {
	var _User_userPicture = &User_userPicture{UserPictureHeight: iUserPictureHeight}
	has, err := Engine.Get(_User_userPicture)
	if has {
		return _User_userPicture, err
	} else {
		return nil, err
	}
}

// SetUser_userPictureViaBundle Set User_userPicture via Bundle
func SetUser_userPictureViaBundle(iBundle string, user__user_picture *User_userPicture) (int64, error) {
	user__user_picture.Bundle = iBundle
	return Engine.Insert(user__user_picture)
}

// SetUser_userPictureViaDeleted Set User_userPicture via Deleted
func SetUser_userPictureViaDeleted(iDeleted int, user__user_picture *User_userPicture) (int64, error) {
	user__user_picture.Deleted = iDeleted
	return Engine.Insert(user__user_picture)
}

// SetUser_userPictureViaEntityId Set User_userPicture via EntityId
func SetUser_userPictureViaEntityId(iEntityId int, user__user_picture *User_userPicture) (int64, error) {
	user__user_picture.EntityId = iEntityId
	return Engine.Insert(user__user_picture)
}

// SetUser_userPictureViaRevisionId Set User_userPicture via RevisionId
func SetUser_userPictureViaRevisionId(iRevisionId int, user__user_picture *User_userPicture) (int64, error) {
	user__user_picture.RevisionId = iRevisionId
	return Engine.Insert(user__user_picture)
}

// SetUser_userPictureViaLangcode Set User_userPicture via Langcode
func SetUser_userPictureViaLangcode(iLangcode string, user__user_picture *User_userPicture) (int64, error) {
	user__user_picture.Langcode = iLangcode
	return Engine.Insert(user__user_picture)
}

// SetUser_userPictureViaDelta Set User_userPicture via Delta
func SetUser_userPictureViaDelta(iDelta int, user__user_picture *User_userPicture) (int64, error) {
	user__user_picture.Delta = iDelta
	return Engine.Insert(user__user_picture)
}

// SetUser_userPictureViaUserPictureTargetId Set User_userPicture via UserPictureTargetId
func SetUser_userPictureViaUserPictureTargetId(iUserPictureTargetId int, user__user_picture *User_userPicture) (int64, error) {
	user__user_picture.UserPictureTargetId = iUserPictureTargetId
	return Engine.Insert(user__user_picture)
}

// SetUser_userPictureViaUserPictureAlt Set User_userPicture via UserPictureAlt
func SetUser_userPictureViaUserPictureAlt(iUserPictureAlt string, user__user_picture *User_userPicture) (int64, error) {
	user__user_picture.UserPictureAlt = iUserPictureAlt
	return Engine.Insert(user__user_picture)
}

// SetUser_userPictureViaUserPictureTitle Set User_userPicture via UserPictureTitle
func SetUser_userPictureViaUserPictureTitle(iUserPictureTitle string, user__user_picture *User_userPicture) (int64, error) {
	user__user_picture.UserPictureTitle = iUserPictureTitle
	return Engine.Insert(user__user_picture)
}

// SetUser_userPictureViaUserPictureWidth Set User_userPicture via UserPictureWidth
func SetUser_userPictureViaUserPictureWidth(iUserPictureWidth int, user__user_picture *User_userPicture) (int64, error) {
	user__user_picture.UserPictureWidth = iUserPictureWidth
	return Engine.Insert(user__user_picture)
}

// SetUser_userPictureViaUserPictureHeight Set User_userPicture via UserPictureHeight
func SetUser_userPictureViaUserPictureHeight(iUserPictureHeight int, user__user_picture *User_userPicture) (int64, error) {
	user__user_picture.UserPictureHeight = iUserPictureHeight
	return Engine.Insert(user__user_picture)
}

// AddUser_userPicture Add User_userPicture via all columns
func AddUser_userPicture(iBundle string, iDeleted int, iEntityId int, iRevisionId int, iLangcode string, iDelta int, iUserPictureTargetId int, iUserPictureAlt string, iUserPictureTitle string, iUserPictureWidth int, iUserPictureHeight int) error {
	_User_userPicture := &User_userPicture{Bundle: iBundle, Deleted: iDeleted, EntityId: iEntityId, RevisionId: iRevisionId, Langcode: iLangcode, Delta: iDelta, UserPictureTargetId: iUserPictureTargetId, UserPictureAlt: iUserPictureAlt, UserPictureTitle: iUserPictureTitle, UserPictureWidth: iUserPictureWidth, UserPictureHeight: iUserPictureHeight}
	if _, err := Engine.Insert(_User_userPicture); err != nil {
		return err
	} else {
		return nil
	}
}

// PostUser_userPicture Post User_userPicture via iUser_userPicture
func PostUser_userPicture(iUser_userPicture *User_userPicture) (string, error) {
	_, err := Engine.Insert(iUser_userPicture)
	return iUser_userPicture.Bundle, err
}

// PutUser_userPicture Put User_userPicture
func PutUser_userPicture(iUser_userPicture *User_userPicture) (string, error) {
	_, err := Engine.Id(iUser_userPicture.Bundle).Update(iUser_userPicture)
	return iUser_userPicture.Bundle, err
}

// PutUser_userPictureViaBundle Put User_userPicture via Bundle
func PutUser_userPictureViaBundle(Bundle_ string, iUser_userPicture *User_userPicture) (int64, error) {
	row, err := Engine.Update(iUser_userPicture, &User_userPicture{Bundle: Bundle_})
	return row, err
}

// PutUser_userPictureViaDeleted Put User_userPicture via Deleted
func PutUser_userPictureViaDeleted(Deleted_ int, iUser_userPicture *User_userPicture) (int64, error) {
	row, err := Engine.Update(iUser_userPicture, &User_userPicture{Deleted: Deleted_})
	return row, err
}

// PutUser_userPictureViaEntityId Put User_userPicture via EntityId
func PutUser_userPictureViaEntityId(EntityId_ int, iUser_userPicture *User_userPicture) (int64, error) {
	row, err := Engine.Update(iUser_userPicture, &User_userPicture{EntityId: EntityId_})
	return row, err
}

// PutUser_userPictureViaRevisionId Put User_userPicture via RevisionId
func PutUser_userPictureViaRevisionId(RevisionId_ int, iUser_userPicture *User_userPicture) (int64, error) {
	row, err := Engine.Update(iUser_userPicture, &User_userPicture{RevisionId: RevisionId_})
	return row, err
}

// PutUser_userPictureViaLangcode Put User_userPicture via Langcode
func PutUser_userPictureViaLangcode(Langcode_ string, iUser_userPicture *User_userPicture) (int64, error) {
	row, err := Engine.Update(iUser_userPicture, &User_userPicture{Langcode: Langcode_})
	return row, err
}

// PutUser_userPictureViaDelta Put User_userPicture via Delta
func PutUser_userPictureViaDelta(Delta_ int, iUser_userPicture *User_userPicture) (int64, error) {
	row, err := Engine.Update(iUser_userPicture, &User_userPicture{Delta: Delta_})
	return row, err
}

// PutUser_userPictureViaUserPictureTargetId Put User_userPicture via UserPictureTargetId
func PutUser_userPictureViaUserPictureTargetId(UserPictureTargetId_ int, iUser_userPicture *User_userPicture) (int64, error) {
	row, err := Engine.Update(iUser_userPicture, &User_userPicture{UserPictureTargetId: UserPictureTargetId_})
	return row, err
}

// PutUser_userPictureViaUserPictureAlt Put User_userPicture via UserPictureAlt
func PutUser_userPictureViaUserPictureAlt(UserPictureAlt_ string, iUser_userPicture *User_userPicture) (int64, error) {
	row, err := Engine.Update(iUser_userPicture, &User_userPicture{UserPictureAlt: UserPictureAlt_})
	return row, err
}

// PutUser_userPictureViaUserPictureTitle Put User_userPicture via UserPictureTitle
func PutUser_userPictureViaUserPictureTitle(UserPictureTitle_ string, iUser_userPicture *User_userPicture) (int64, error) {
	row, err := Engine.Update(iUser_userPicture, &User_userPicture{UserPictureTitle: UserPictureTitle_})
	return row, err
}

// PutUser_userPictureViaUserPictureWidth Put User_userPicture via UserPictureWidth
func PutUser_userPictureViaUserPictureWidth(UserPictureWidth_ int, iUser_userPicture *User_userPicture) (int64, error) {
	row, err := Engine.Update(iUser_userPicture, &User_userPicture{UserPictureWidth: UserPictureWidth_})
	return row, err
}

// PutUser_userPictureViaUserPictureHeight Put User_userPicture via UserPictureHeight
func PutUser_userPictureViaUserPictureHeight(UserPictureHeight_ int, iUser_userPicture *User_userPicture) (int64, error) {
	row, err := Engine.Update(iUser_userPicture, &User_userPicture{UserPictureHeight: UserPictureHeight_})
	return row, err
}

// UpdateUser_userPictureViaBundle via map[string]interface{}{}
func UpdateUser_userPictureViaBundle(iBundle string, iUser_userPictureMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(User_userPicture)).Where("bundle = ?", iBundle).Update(iUser_userPictureMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateUser_userPictureViaDeleted via map[string]interface{}{}
func UpdateUser_userPictureViaDeleted(iDeleted int, iUser_userPictureMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(User_userPicture)).Where("deleted = ?", iDeleted).Update(iUser_userPictureMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateUser_userPictureViaEntityId via map[string]interface{}{}
func UpdateUser_userPictureViaEntityId(iEntityId int, iUser_userPictureMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(User_userPicture)).Where("entity_id = ?", iEntityId).Update(iUser_userPictureMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateUser_userPictureViaRevisionId via map[string]interface{}{}
func UpdateUser_userPictureViaRevisionId(iRevisionId int, iUser_userPictureMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(User_userPicture)).Where("revision_id = ?", iRevisionId).Update(iUser_userPictureMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateUser_userPictureViaLangcode via map[string]interface{}{}
func UpdateUser_userPictureViaLangcode(iLangcode string, iUser_userPictureMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(User_userPicture)).Where("langcode = ?", iLangcode).Update(iUser_userPictureMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateUser_userPictureViaDelta via map[string]interface{}{}
func UpdateUser_userPictureViaDelta(iDelta int, iUser_userPictureMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(User_userPicture)).Where("delta = ?", iDelta).Update(iUser_userPictureMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateUser_userPictureViaUserPictureTargetId via map[string]interface{}{}
func UpdateUser_userPictureViaUserPictureTargetId(iUserPictureTargetId int, iUser_userPictureMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(User_userPicture)).Where("user_picture_target_id = ?", iUserPictureTargetId).Update(iUser_userPictureMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateUser_userPictureViaUserPictureAlt via map[string]interface{}{}
func UpdateUser_userPictureViaUserPictureAlt(iUserPictureAlt string, iUser_userPictureMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(User_userPicture)).Where("user_picture_alt = ?", iUserPictureAlt).Update(iUser_userPictureMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateUser_userPictureViaUserPictureTitle via map[string]interface{}{}
func UpdateUser_userPictureViaUserPictureTitle(iUserPictureTitle string, iUser_userPictureMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(User_userPicture)).Where("user_picture_title = ?", iUserPictureTitle).Update(iUser_userPictureMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateUser_userPictureViaUserPictureWidth via map[string]interface{}{}
func UpdateUser_userPictureViaUserPictureWidth(iUserPictureWidth int, iUser_userPictureMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(User_userPicture)).Where("user_picture_width = ?", iUserPictureWidth).Update(iUser_userPictureMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateUser_userPictureViaUserPictureHeight via map[string]interface{}{}
func UpdateUser_userPictureViaUserPictureHeight(iUserPictureHeight int, iUser_userPictureMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(User_userPicture)).Where("user_picture_height = ?", iUserPictureHeight).Update(iUser_userPictureMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// DeleteUser_userPicture Delete User_userPicture
func DeleteUser_userPicture(iBundle string) (int64, error) {
	row, err := Engine.Id(iBundle).Delete(new(User_userPicture))
	return row, err
}

// DeleteUser_userPictureViaBundle Delete User_userPicture via Bundle
func DeleteUser_userPictureViaBundle(iBundle string) (err error) {
	var has bool
	var _User_userPicture = &User_userPicture{Bundle: iBundle}
	if has, err = Engine.Get(_User_userPicture); (has == true) && (err == nil) {
		if row, err := Engine.Where("bundle = ?", iBundle).Delete(new(User_userPicture)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteUser_userPictureViaDeleted Delete User_userPicture via Deleted
func DeleteUser_userPictureViaDeleted(iDeleted int) (err error) {
	var has bool
	var _User_userPicture = &User_userPicture{Deleted: iDeleted}
	if has, err = Engine.Get(_User_userPicture); (has == true) && (err == nil) {
		if row, err := Engine.Where("deleted = ?", iDeleted).Delete(new(User_userPicture)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteUser_userPictureViaEntityId Delete User_userPicture via EntityId
func DeleteUser_userPictureViaEntityId(iEntityId int) (err error) {
	var has bool
	var _User_userPicture = &User_userPicture{EntityId: iEntityId}
	if has, err = Engine.Get(_User_userPicture); (has == true) && (err == nil) {
		if row, err := Engine.Where("entity_id = ?", iEntityId).Delete(new(User_userPicture)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteUser_userPictureViaRevisionId Delete User_userPicture via RevisionId
func DeleteUser_userPictureViaRevisionId(iRevisionId int) (err error) {
	var has bool
	var _User_userPicture = &User_userPicture{RevisionId: iRevisionId}
	if has, err = Engine.Get(_User_userPicture); (has == true) && (err == nil) {
		if row, err := Engine.Where("revision_id = ?", iRevisionId).Delete(new(User_userPicture)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteUser_userPictureViaLangcode Delete User_userPicture via Langcode
func DeleteUser_userPictureViaLangcode(iLangcode string) (err error) {
	var has bool
	var _User_userPicture = &User_userPicture{Langcode: iLangcode}
	if has, err = Engine.Get(_User_userPicture); (has == true) && (err == nil) {
		if row, err := Engine.Where("langcode = ?", iLangcode).Delete(new(User_userPicture)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteUser_userPictureViaDelta Delete User_userPicture via Delta
func DeleteUser_userPictureViaDelta(iDelta int) (err error) {
	var has bool
	var _User_userPicture = &User_userPicture{Delta: iDelta}
	if has, err = Engine.Get(_User_userPicture); (has == true) && (err == nil) {
		if row, err := Engine.Where("delta = ?", iDelta).Delete(new(User_userPicture)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteUser_userPictureViaUserPictureTargetId Delete User_userPicture via UserPictureTargetId
func DeleteUser_userPictureViaUserPictureTargetId(iUserPictureTargetId int) (err error) {
	var has bool
	var _User_userPicture = &User_userPicture{UserPictureTargetId: iUserPictureTargetId}
	if has, err = Engine.Get(_User_userPicture); (has == true) && (err == nil) {
		if row, err := Engine.Where("user_picture_target_id = ?", iUserPictureTargetId).Delete(new(User_userPicture)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteUser_userPictureViaUserPictureAlt Delete User_userPicture via UserPictureAlt
func DeleteUser_userPictureViaUserPictureAlt(iUserPictureAlt string) (err error) {
	var has bool
	var _User_userPicture = &User_userPicture{UserPictureAlt: iUserPictureAlt}
	if has, err = Engine.Get(_User_userPicture); (has == true) && (err == nil) {
		if row, err := Engine.Where("user_picture_alt = ?", iUserPictureAlt).Delete(new(User_userPicture)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteUser_userPictureViaUserPictureTitle Delete User_userPicture via UserPictureTitle
func DeleteUser_userPictureViaUserPictureTitle(iUserPictureTitle string) (err error) {
	var has bool
	var _User_userPicture = &User_userPicture{UserPictureTitle: iUserPictureTitle}
	if has, err = Engine.Get(_User_userPicture); (has == true) && (err == nil) {
		if row, err := Engine.Where("user_picture_title = ?", iUserPictureTitle).Delete(new(User_userPicture)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteUser_userPictureViaUserPictureWidth Delete User_userPicture via UserPictureWidth
func DeleteUser_userPictureViaUserPictureWidth(iUserPictureWidth int) (err error) {
	var has bool
	var _User_userPicture = &User_userPicture{UserPictureWidth: iUserPictureWidth}
	if has, err = Engine.Get(_User_userPicture); (has == true) && (err == nil) {
		if row, err := Engine.Where("user_picture_width = ?", iUserPictureWidth).Delete(new(User_userPicture)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteUser_userPictureViaUserPictureHeight Delete User_userPicture via UserPictureHeight
func DeleteUser_userPictureViaUserPictureHeight(iUserPictureHeight int) (err error) {
	var has bool
	var _User_userPicture = &User_userPicture{UserPictureHeight: iUserPictureHeight}
	if has, err = Engine.Get(_User_userPicture); (has == true) && (err == nil) {
		if row, err := Engine.Where("user_picture_height = ?", iUserPictureHeight).Delete(new(User_userPicture)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}
