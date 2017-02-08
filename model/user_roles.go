package model

type User_roles struct {
	Bundle        string `xorm:"not null default '' index VARCHAR(128)"`
	Deleted       int    `xorm:"not null pk default 0 TINYINT(4)"`
	EntityId      int    `xorm:"not null pk INT(10)"`
	RevisionId    int    `xorm:"not null index INT(10)"`
	Langcode      string `xorm:"not null pk default '' VARCHAR(32)"`
	Delta         int    `xorm:"not null pk INT(10)"`
	RolesTargetId string `xorm:"not null index VARCHAR(255)"`
}

// GetUser_rolesesCount User_roless' Count
func GetUser_rolesesCount(offset int, limit int) (int64, error) {
	total, err := Engine.Limit(limit, offset).Count(&User_roles{})
	return total, err
}

// GetUser_rolesCountViaBundle Get User_roles via Bundle
func GetUser_rolesCountViaBundle(iBundle string) int64 {
	n, _ := Engine.Where("bundle = ?", iBundle).Count(&User_roles{Bundle: iBundle})
	return n
}

// GetUser_rolesCountViaDeleted Get User_roles via Deleted
func GetUser_rolesCountViaDeleted(iDeleted int) int64 {
	n, _ := Engine.Where("deleted = ?", iDeleted).Count(&User_roles{Deleted: iDeleted})
	return n
}

// GetUser_rolesCountViaEntityId Get User_roles via EntityId
func GetUser_rolesCountViaEntityId(iEntityId int) int64 {
	n, _ := Engine.Where("entity_id = ?", iEntityId).Count(&User_roles{EntityId: iEntityId})
	return n
}

// GetUser_rolesCountViaRevisionId Get User_roles via RevisionId
func GetUser_rolesCountViaRevisionId(iRevisionId int) int64 {
	n, _ := Engine.Where("revision_id = ?", iRevisionId).Count(&User_roles{RevisionId: iRevisionId})
	return n
}

// GetUser_rolesCountViaLangcode Get User_roles via Langcode
func GetUser_rolesCountViaLangcode(iLangcode string) int64 {
	n, _ := Engine.Where("langcode = ?", iLangcode).Count(&User_roles{Langcode: iLangcode})
	return n
}

// GetUser_rolesCountViaDelta Get User_roles via Delta
func GetUser_rolesCountViaDelta(iDelta int) int64 {
	n, _ := Engine.Where("delta = ?", iDelta).Count(&User_roles{Delta: iDelta})
	return n
}

// GetUser_rolesCountViaRolesTargetId Get User_roles via RolesTargetId
func GetUser_rolesCountViaRolesTargetId(iRolesTargetId string) int64 {
	n, _ := Engine.Where("roles_target_id = ?", iRolesTargetId).Count(&User_roles{RolesTargetId: iRolesTargetId})
	return n
}

// GetUser_rolesesByBundleAndDeletedAndEntityId Get User_roleses via BundleAndDeletedAndEntityId
func GetUser_rolesesByBundleAndDeletedAndEntityId(offset int, limit int, Bundle_ string, Deleted_ int, EntityId_ int) (*[]*User_roles, error) {
	var _User_roles = new([]*User_roles)
	err := Engine.Table("user__roles").Where("bundle = ? and deleted = ? and entity_id = ?", Bundle_, Deleted_, EntityId_).Limit(limit, offset).Find(_User_roles)
	return _User_roles, err
}

// GetUser_rolesesByBundleAndDeletedAndRevisionId Get User_roleses via BundleAndDeletedAndRevisionId
func GetUser_rolesesByBundleAndDeletedAndRevisionId(offset int, limit int, Bundle_ string, Deleted_ int, RevisionId_ int) (*[]*User_roles, error) {
	var _User_roles = new([]*User_roles)
	err := Engine.Table("user__roles").Where("bundle = ? and deleted = ? and revision_id = ?", Bundle_, Deleted_, RevisionId_).Limit(limit, offset).Find(_User_roles)
	return _User_roles, err
}

// GetUser_rolesesByBundleAndDeletedAndLangcode Get User_roleses via BundleAndDeletedAndLangcode
func GetUser_rolesesByBundleAndDeletedAndLangcode(offset int, limit int, Bundle_ string, Deleted_ int, Langcode_ string) (*[]*User_roles, error) {
	var _User_roles = new([]*User_roles)
	err := Engine.Table("user__roles").Where("bundle = ? and deleted = ? and langcode = ?", Bundle_, Deleted_, Langcode_).Limit(limit, offset).Find(_User_roles)
	return _User_roles, err
}

// GetUser_rolesesByBundleAndDeletedAndDelta Get User_roleses via BundleAndDeletedAndDelta
func GetUser_rolesesByBundleAndDeletedAndDelta(offset int, limit int, Bundle_ string, Deleted_ int, Delta_ int) (*[]*User_roles, error) {
	var _User_roles = new([]*User_roles)
	err := Engine.Table("user__roles").Where("bundle = ? and deleted = ? and delta = ?", Bundle_, Deleted_, Delta_).Limit(limit, offset).Find(_User_roles)
	return _User_roles, err
}

// GetUser_rolesesByBundleAndDeletedAndRolesTargetId Get User_roleses via BundleAndDeletedAndRolesTargetId
func GetUser_rolesesByBundleAndDeletedAndRolesTargetId(offset int, limit int, Bundle_ string, Deleted_ int, RolesTargetId_ string) (*[]*User_roles, error) {
	var _User_roles = new([]*User_roles)
	err := Engine.Table("user__roles").Where("bundle = ? and deleted = ? and roles_target_id = ?", Bundle_, Deleted_, RolesTargetId_).Limit(limit, offset).Find(_User_roles)
	return _User_roles, err
}

// GetUser_rolesesByBundleAndEntityIdAndRevisionId Get User_roleses via BundleAndEntityIdAndRevisionId
func GetUser_rolesesByBundleAndEntityIdAndRevisionId(offset int, limit int, Bundle_ string, EntityId_ int, RevisionId_ int) (*[]*User_roles, error) {
	var _User_roles = new([]*User_roles)
	err := Engine.Table("user__roles").Where("bundle = ? and entity_id = ? and revision_id = ?", Bundle_, EntityId_, RevisionId_).Limit(limit, offset).Find(_User_roles)
	return _User_roles, err
}

// GetUser_rolesesByBundleAndEntityIdAndLangcode Get User_roleses via BundleAndEntityIdAndLangcode
func GetUser_rolesesByBundleAndEntityIdAndLangcode(offset int, limit int, Bundle_ string, EntityId_ int, Langcode_ string) (*[]*User_roles, error) {
	var _User_roles = new([]*User_roles)
	err := Engine.Table("user__roles").Where("bundle = ? and entity_id = ? and langcode = ?", Bundle_, EntityId_, Langcode_).Limit(limit, offset).Find(_User_roles)
	return _User_roles, err
}

// GetUser_rolesesByBundleAndEntityIdAndDelta Get User_roleses via BundleAndEntityIdAndDelta
func GetUser_rolesesByBundleAndEntityIdAndDelta(offset int, limit int, Bundle_ string, EntityId_ int, Delta_ int) (*[]*User_roles, error) {
	var _User_roles = new([]*User_roles)
	err := Engine.Table("user__roles").Where("bundle = ? and entity_id = ? and delta = ?", Bundle_, EntityId_, Delta_).Limit(limit, offset).Find(_User_roles)
	return _User_roles, err
}

// GetUser_rolesesByBundleAndEntityIdAndRolesTargetId Get User_roleses via BundleAndEntityIdAndRolesTargetId
func GetUser_rolesesByBundleAndEntityIdAndRolesTargetId(offset int, limit int, Bundle_ string, EntityId_ int, RolesTargetId_ string) (*[]*User_roles, error) {
	var _User_roles = new([]*User_roles)
	err := Engine.Table("user__roles").Where("bundle = ? and entity_id = ? and roles_target_id = ?", Bundle_, EntityId_, RolesTargetId_).Limit(limit, offset).Find(_User_roles)
	return _User_roles, err
}

// GetUser_rolesesByBundleAndRevisionIdAndLangcode Get User_roleses via BundleAndRevisionIdAndLangcode
func GetUser_rolesesByBundleAndRevisionIdAndLangcode(offset int, limit int, Bundle_ string, RevisionId_ int, Langcode_ string) (*[]*User_roles, error) {
	var _User_roles = new([]*User_roles)
	err := Engine.Table("user__roles").Where("bundle = ? and revision_id = ? and langcode = ?", Bundle_, RevisionId_, Langcode_).Limit(limit, offset).Find(_User_roles)
	return _User_roles, err
}

// GetUser_rolesesByBundleAndRevisionIdAndDelta Get User_roleses via BundleAndRevisionIdAndDelta
func GetUser_rolesesByBundleAndRevisionIdAndDelta(offset int, limit int, Bundle_ string, RevisionId_ int, Delta_ int) (*[]*User_roles, error) {
	var _User_roles = new([]*User_roles)
	err := Engine.Table("user__roles").Where("bundle = ? and revision_id = ? and delta = ?", Bundle_, RevisionId_, Delta_).Limit(limit, offset).Find(_User_roles)
	return _User_roles, err
}

// GetUser_rolesesByBundleAndRevisionIdAndRolesTargetId Get User_roleses via BundleAndRevisionIdAndRolesTargetId
func GetUser_rolesesByBundleAndRevisionIdAndRolesTargetId(offset int, limit int, Bundle_ string, RevisionId_ int, RolesTargetId_ string) (*[]*User_roles, error) {
	var _User_roles = new([]*User_roles)
	err := Engine.Table("user__roles").Where("bundle = ? and revision_id = ? and roles_target_id = ?", Bundle_, RevisionId_, RolesTargetId_).Limit(limit, offset).Find(_User_roles)
	return _User_roles, err
}

// GetUser_rolesesByBundleAndLangcodeAndDelta Get User_roleses via BundleAndLangcodeAndDelta
func GetUser_rolesesByBundleAndLangcodeAndDelta(offset int, limit int, Bundle_ string, Langcode_ string, Delta_ int) (*[]*User_roles, error) {
	var _User_roles = new([]*User_roles)
	err := Engine.Table("user__roles").Where("bundle = ? and langcode = ? and delta = ?", Bundle_, Langcode_, Delta_).Limit(limit, offset).Find(_User_roles)
	return _User_roles, err
}

// GetUser_rolesesByBundleAndLangcodeAndRolesTargetId Get User_roleses via BundleAndLangcodeAndRolesTargetId
func GetUser_rolesesByBundleAndLangcodeAndRolesTargetId(offset int, limit int, Bundle_ string, Langcode_ string, RolesTargetId_ string) (*[]*User_roles, error) {
	var _User_roles = new([]*User_roles)
	err := Engine.Table("user__roles").Where("bundle = ? and langcode = ? and roles_target_id = ?", Bundle_, Langcode_, RolesTargetId_).Limit(limit, offset).Find(_User_roles)
	return _User_roles, err
}

// GetUser_rolesesByBundleAndDeltaAndRolesTargetId Get User_roleses via BundleAndDeltaAndRolesTargetId
func GetUser_rolesesByBundleAndDeltaAndRolesTargetId(offset int, limit int, Bundle_ string, Delta_ int, RolesTargetId_ string) (*[]*User_roles, error) {
	var _User_roles = new([]*User_roles)
	err := Engine.Table("user__roles").Where("bundle = ? and delta = ? and roles_target_id = ?", Bundle_, Delta_, RolesTargetId_).Limit(limit, offset).Find(_User_roles)
	return _User_roles, err
}

// GetUser_rolesesByDeletedAndEntityIdAndRevisionId Get User_roleses via DeletedAndEntityIdAndRevisionId
func GetUser_rolesesByDeletedAndEntityIdAndRevisionId(offset int, limit int, Deleted_ int, EntityId_ int, RevisionId_ int) (*[]*User_roles, error) {
	var _User_roles = new([]*User_roles)
	err := Engine.Table("user__roles").Where("deleted = ? and entity_id = ? and revision_id = ?", Deleted_, EntityId_, RevisionId_).Limit(limit, offset).Find(_User_roles)
	return _User_roles, err
}

// GetUser_rolesesByDeletedAndEntityIdAndLangcode Get User_roleses via DeletedAndEntityIdAndLangcode
func GetUser_rolesesByDeletedAndEntityIdAndLangcode(offset int, limit int, Deleted_ int, EntityId_ int, Langcode_ string) (*[]*User_roles, error) {
	var _User_roles = new([]*User_roles)
	err := Engine.Table("user__roles").Where("deleted = ? and entity_id = ? and langcode = ?", Deleted_, EntityId_, Langcode_).Limit(limit, offset).Find(_User_roles)
	return _User_roles, err
}

// GetUser_rolesesByDeletedAndEntityIdAndDelta Get User_roleses via DeletedAndEntityIdAndDelta
func GetUser_rolesesByDeletedAndEntityIdAndDelta(offset int, limit int, Deleted_ int, EntityId_ int, Delta_ int) (*[]*User_roles, error) {
	var _User_roles = new([]*User_roles)
	err := Engine.Table("user__roles").Where("deleted = ? and entity_id = ? and delta = ?", Deleted_, EntityId_, Delta_).Limit(limit, offset).Find(_User_roles)
	return _User_roles, err
}

// GetUser_rolesesByDeletedAndEntityIdAndRolesTargetId Get User_roleses via DeletedAndEntityIdAndRolesTargetId
func GetUser_rolesesByDeletedAndEntityIdAndRolesTargetId(offset int, limit int, Deleted_ int, EntityId_ int, RolesTargetId_ string) (*[]*User_roles, error) {
	var _User_roles = new([]*User_roles)
	err := Engine.Table("user__roles").Where("deleted = ? and entity_id = ? and roles_target_id = ?", Deleted_, EntityId_, RolesTargetId_).Limit(limit, offset).Find(_User_roles)
	return _User_roles, err
}

// GetUser_rolesesByDeletedAndRevisionIdAndLangcode Get User_roleses via DeletedAndRevisionIdAndLangcode
func GetUser_rolesesByDeletedAndRevisionIdAndLangcode(offset int, limit int, Deleted_ int, RevisionId_ int, Langcode_ string) (*[]*User_roles, error) {
	var _User_roles = new([]*User_roles)
	err := Engine.Table("user__roles").Where("deleted = ? and revision_id = ? and langcode = ?", Deleted_, RevisionId_, Langcode_).Limit(limit, offset).Find(_User_roles)
	return _User_roles, err
}

// GetUser_rolesesByDeletedAndRevisionIdAndDelta Get User_roleses via DeletedAndRevisionIdAndDelta
func GetUser_rolesesByDeletedAndRevisionIdAndDelta(offset int, limit int, Deleted_ int, RevisionId_ int, Delta_ int) (*[]*User_roles, error) {
	var _User_roles = new([]*User_roles)
	err := Engine.Table("user__roles").Where("deleted = ? and revision_id = ? and delta = ?", Deleted_, RevisionId_, Delta_).Limit(limit, offset).Find(_User_roles)
	return _User_roles, err
}

// GetUser_rolesesByDeletedAndRevisionIdAndRolesTargetId Get User_roleses via DeletedAndRevisionIdAndRolesTargetId
func GetUser_rolesesByDeletedAndRevisionIdAndRolesTargetId(offset int, limit int, Deleted_ int, RevisionId_ int, RolesTargetId_ string) (*[]*User_roles, error) {
	var _User_roles = new([]*User_roles)
	err := Engine.Table("user__roles").Where("deleted = ? and revision_id = ? and roles_target_id = ?", Deleted_, RevisionId_, RolesTargetId_).Limit(limit, offset).Find(_User_roles)
	return _User_roles, err
}

// GetUser_rolesesByDeletedAndLangcodeAndDelta Get User_roleses via DeletedAndLangcodeAndDelta
func GetUser_rolesesByDeletedAndLangcodeAndDelta(offset int, limit int, Deleted_ int, Langcode_ string, Delta_ int) (*[]*User_roles, error) {
	var _User_roles = new([]*User_roles)
	err := Engine.Table("user__roles").Where("deleted = ? and langcode = ? and delta = ?", Deleted_, Langcode_, Delta_).Limit(limit, offset).Find(_User_roles)
	return _User_roles, err
}

// GetUser_rolesesByDeletedAndLangcodeAndRolesTargetId Get User_roleses via DeletedAndLangcodeAndRolesTargetId
func GetUser_rolesesByDeletedAndLangcodeAndRolesTargetId(offset int, limit int, Deleted_ int, Langcode_ string, RolesTargetId_ string) (*[]*User_roles, error) {
	var _User_roles = new([]*User_roles)
	err := Engine.Table("user__roles").Where("deleted = ? and langcode = ? and roles_target_id = ?", Deleted_, Langcode_, RolesTargetId_).Limit(limit, offset).Find(_User_roles)
	return _User_roles, err
}

// GetUser_rolesesByDeletedAndDeltaAndRolesTargetId Get User_roleses via DeletedAndDeltaAndRolesTargetId
func GetUser_rolesesByDeletedAndDeltaAndRolesTargetId(offset int, limit int, Deleted_ int, Delta_ int, RolesTargetId_ string) (*[]*User_roles, error) {
	var _User_roles = new([]*User_roles)
	err := Engine.Table("user__roles").Where("deleted = ? and delta = ? and roles_target_id = ?", Deleted_, Delta_, RolesTargetId_).Limit(limit, offset).Find(_User_roles)
	return _User_roles, err
}

// GetUser_rolesesByEntityIdAndRevisionIdAndLangcode Get User_roleses via EntityIdAndRevisionIdAndLangcode
func GetUser_rolesesByEntityIdAndRevisionIdAndLangcode(offset int, limit int, EntityId_ int, RevisionId_ int, Langcode_ string) (*[]*User_roles, error) {
	var _User_roles = new([]*User_roles)
	err := Engine.Table("user__roles").Where("entity_id = ? and revision_id = ? and langcode = ?", EntityId_, RevisionId_, Langcode_).Limit(limit, offset).Find(_User_roles)
	return _User_roles, err
}

// GetUser_rolesesByEntityIdAndRevisionIdAndDelta Get User_roleses via EntityIdAndRevisionIdAndDelta
func GetUser_rolesesByEntityIdAndRevisionIdAndDelta(offset int, limit int, EntityId_ int, RevisionId_ int, Delta_ int) (*[]*User_roles, error) {
	var _User_roles = new([]*User_roles)
	err := Engine.Table("user__roles").Where("entity_id = ? and revision_id = ? and delta = ?", EntityId_, RevisionId_, Delta_).Limit(limit, offset).Find(_User_roles)
	return _User_roles, err
}

// GetUser_rolesesByEntityIdAndRevisionIdAndRolesTargetId Get User_roleses via EntityIdAndRevisionIdAndRolesTargetId
func GetUser_rolesesByEntityIdAndRevisionIdAndRolesTargetId(offset int, limit int, EntityId_ int, RevisionId_ int, RolesTargetId_ string) (*[]*User_roles, error) {
	var _User_roles = new([]*User_roles)
	err := Engine.Table("user__roles").Where("entity_id = ? and revision_id = ? and roles_target_id = ?", EntityId_, RevisionId_, RolesTargetId_).Limit(limit, offset).Find(_User_roles)
	return _User_roles, err
}

// GetUser_rolesesByEntityIdAndLangcodeAndDelta Get User_roleses via EntityIdAndLangcodeAndDelta
func GetUser_rolesesByEntityIdAndLangcodeAndDelta(offset int, limit int, EntityId_ int, Langcode_ string, Delta_ int) (*[]*User_roles, error) {
	var _User_roles = new([]*User_roles)
	err := Engine.Table("user__roles").Where("entity_id = ? and langcode = ? and delta = ?", EntityId_, Langcode_, Delta_).Limit(limit, offset).Find(_User_roles)
	return _User_roles, err
}

// GetUser_rolesesByEntityIdAndLangcodeAndRolesTargetId Get User_roleses via EntityIdAndLangcodeAndRolesTargetId
func GetUser_rolesesByEntityIdAndLangcodeAndRolesTargetId(offset int, limit int, EntityId_ int, Langcode_ string, RolesTargetId_ string) (*[]*User_roles, error) {
	var _User_roles = new([]*User_roles)
	err := Engine.Table("user__roles").Where("entity_id = ? and langcode = ? and roles_target_id = ?", EntityId_, Langcode_, RolesTargetId_).Limit(limit, offset).Find(_User_roles)
	return _User_roles, err
}

// GetUser_rolesesByEntityIdAndDeltaAndRolesTargetId Get User_roleses via EntityIdAndDeltaAndRolesTargetId
func GetUser_rolesesByEntityIdAndDeltaAndRolesTargetId(offset int, limit int, EntityId_ int, Delta_ int, RolesTargetId_ string) (*[]*User_roles, error) {
	var _User_roles = new([]*User_roles)
	err := Engine.Table("user__roles").Where("entity_id = ? and delta = ? and roles_target_id = ?", EntityId_, Delta_, RolesTargetId_).Limit(limit, offset).Find(_User_roles)
	return _User_roles, err
}

// GetUser_rolesesByRevisionIdAndLangcodeAndDelta Get User_roleses via RevisionIdAndLangcodeAndDelta
func GetUser_rolesesByRevisionIdAndLangcodeAndDelta(offset int, limit int, RevisionId_ int, Langcode_ string, Delta_ int) (*[]*User_roles, error) {
	var _User_roles = new([]*User_roles)
	err := Engine.Table("user__roles").Where("revision_id = ? and langcode = ? and delta = ?", RevisionId_, Langcode_, Delta_).Limit(limit, offset).Find(_User_roles)
	return _User_roles, err
}

// GetUser_rolesesByRevisionIdAndLangcodeAndRolesTargetId Get User_roleses via RevisionIdAndLangcodeAndRolesTargetId
func GetUser_rolesesByRevisionIdAndLangcodeAndRolesTargetId(offset int, limit int, RevisionId_ int, Langcode_ string, RolesTargetId_ string) (*[]*User_roles, error) {
	var _User_roles = new([]*User_roles)
	err := Engine.Table("user__roles").Where("revision_id = ? and langcode = ? and roles_target_id = ?", RevisionId_, Langcode_, RolesTargetId_).Limit(limit, offset).Find(_User_roles)
	return _User_roles, err
}

// GetUser_rolesesByRevisionIdAndDeltaAndRolesTargetId Get User_roleses via RevisionIdAndDeltaAndRolesTargetId
func GetUser_rolesesByRevisionIdAndDeltaAndRolesTargetId(offset int, limit int, RevisionId_ int, Delta_ int, RolesTargetId_ string) (*[]*User_roles, error) {
	var _User_roles = new([]*User_roles)
	err := Engine.Table("user__roles").Where("revision_id = ? and delta = ? and roles_target_id = ?", RevisionId_, Delta_, RolesTargetId_).Limit(limit, offset).Find(_User_roles)
	return _User_roles, err
}

// GetUser_rolesesByLangcodeAndDeltaAndRolesTargetId Get User_roleses via LangcodeAndDeltaAndRolesTargetId
func GetUser_rolesesByLangcodeAndDeltaAndRolesTargetId(offset int, limit int, Langcode_ string, Delta_ int, RolesTargetId_ string) (*[]*User_roles, error) {
	var _User_roles = new([]*User_roles)
	err := Engine.Table("user__roles").Where("langcode = ? and delta = ? and roles_target_id = ?", Langcode_, Delta_, RolesTargetId_).Limit(limit, offset).Find(_User_roles)
	return _User_roles, err
}

// GetUser_rolesesByBundleAndDeleted Get User_roleses via BundleAndDeleted
func GetUser_rolesesByBundleAndDeleted(offset int, limit int, Bundle_ string, Deleted_ int) (*[]*User_roles, error) {
	var _User_roles = new([]*User_roles)
	err := Engine.Table("user__roles").Where("bundle = ? and deleted = ?", Bundle_, Deleted_).Limit(limit, offset).Find(_User_roles)
	return _User_roles, err
}

// GetUser_rolesesByBundleAndEntityId Get User_roleses via BundleAndEntityId
func GetUser_rolesesByBundleAndEntityId(offset int, limit int, Bundle_ string, EntityId_ int) (*[]*User_roles, error) {
	var _User_roles = new([]*User_roles)
	err := Engine.Table("user__roles").Where("bundle = ? and entity_id = ?", Bundle_, EntityId_).Limit(limit, offset).Find(_User_roles)
	return _User_roles, err
}

// GetUser_rolesesByBundleAndRevisionId Get User_roleses via BundleAndRevisionId
func GetUser_rolesesByBundleAndRevisionId(offset int, limit int, Bundle_ string, RevisionId_ int) (*[]*User_roles, error) {
	var _User_roles = new([]*User_roles)
	err := Engine.Table("user__roles").Where("bundle = ? and revision_id = ?", Bundle_, RevisionId_).Limit(limit, offset).Find(_User_roles)
	return _User_roles, err
}

// GetUser_rolesesByBundleAndLangcode Get User_roleses via BundleAndLangcode
func GetUser_rolesesByBundleAndLangcode(offset int, limit int, Bundle_ string, Langcode_ string) (*[]*User_roles, error) {
	var _User_roles = new([]*User_roles)
	err := Engine.Table("user__roles").Where("bundle = ? and langcode = ?", Bundle_, Langcode_).Limit(limit, offset).Find(_User_roles)
	return _User_roles, err
}

// GetUser_rolesesByBundleAndDelta Get User_roleses via BundleAndDelta
func GetUser_rolesesByBundleAndDelta(offset int, limit int, Bundle_ string, Delta_ int) (*[]*User_roles, error) {
	var _User_roles = new([]*User_roles)
	err := Engine.Table("user__roles").Where("bundle = ? and delta = ?", Bundle_, Delta_).Limit(limit, offset).Find(_User_roles)
	return _User_roles, err
}

// GetUser_rolesesByBundleAndRolesTargetId Get User_roleses via BundleAndRolesTargetId
func GetUser_rolesesByBundleAndRolesTargetId(offset int, limit int, Bundle_ string, RolesTargetId_ string) (*[]*User_roles, error) {
	var _User_roles = new([]*User_roles)
	err := Engine.Table("user__roles").Where("bundle = ? and roles_target_id = ?", Bundle_, RolesTargetId_).Limit(limit, offset).Find(_User_roles)
	return _User_roles, err
}

// GetUser_rolesesByDeletedAndEntityId Get User_roleses via DeletedAndEntityId
func GetUser_rolesesByDeletedAndEntityId(offset int, limit int, Deleted_ int, EntityId_ int) (*[]*User_roles, error) {
	var _User_roles = new([]*User_roles)
	err := Engine.Table("user__roles").Where("deleted = ? and entity_id = ?", Deleted_, EntityId_).Limit(limit, offset).Find(_User_roles)
	return _User_roles, err
}

// GetUser_rolesesByDeletedAndRevisionId Get User_roleses via DeletedAndRevisionId
func GetUser_rolesesByDeletedAndRevisionId(offset int, limit int, Deleted_ int, RevisionId_ int) (*[]*User_roles, error) {
	var _User_roles = new([]*User_roles)
	err := Engine.Table("user__roles").Where("deleted = ? and revision_id = ?", Deleted_, RevisionId_).Limit(limit, offset).Find(_User_roles)
	return _User_roles, err
}

// GetUser_rolesesByDeletedAndLangcode Get User_roleses via DeletedAndLangcode
func GetUser_rolesesByDeletedAndLangcode(offset int, limit int, Deleted_ int, Langcode_ string) (*[]*User_roles, error) {
	var _User_roles = new([]*User_roles)
	err := Engine.Table("user__roles").Where("deleted = ? and langcode = ?", Deleted_, Langcode_).Limit(limit, offset).Find(_User_roles)
	return _User_roles, err
}

// GetUser_rolesesByDeletedAndDelta Get User_roleses via DeletedAndDelta
func GetUser_rolesesByDeletedAndDelta(offset int, limit int, Deleted_ int, Delta_ int) (*[]*User_roles, error) {
	var _User_roles = new([]*User_roles)
	err := Engine.Table("user__roles").Where("deleted = ? and delta = ?", Deleted_, Delta_).Limit(limit, offset).Find(_User_roles)
	return _User_roles, err
}

// GetUser_rolesesByDeletedAndRolesTargetId Get User_roleses via DeletedAndRolesTargetId
func GetUser_rolesesByDeletedAndRolesTargetId(offset int, limit int, Deleted_ int, RolesTargetId_ string) (*[]*User_roles, error) {
	var _User_roles = new([]*User_roles)
	err := Engine.Table("user__roles").Where("deleted = ? and roles_target_id = ?", Deleted_, RolesTargetId_).Limit(limit, offset).Find(_User_roles)
	return _User_roles, err
}

// GetUser_rolesesByEntityIdAndRevisionId Get User_roleses via EntityIdAndRevisionId
func GetUser_rolesesByEntityIdAndRevisionId(offset int, limit int, EntityId_ int, RevisionId_ int) (*[]*User_roles, error) {
	var _User_roles = new([]*User_roles)
	err := Engine.Table("user__roles").Where("entity_id = ? and revision_id = ?", EntityId_, RevisionId_).Limit(limit, offset).Find(_User_roles)
	return _User_roles, err
}

// GetUser_rolesesByEntityIdAndLangcode Get User_roleses via EntityIdAndLangcode
func GetUser_rolesesByEntityIdAndLangcode(offset int, limit int, EntityId_ int, Langcode_ string) (*[]*User_roles, error) {
	var _User_roles = new([]*User_roles)
	err := Engine.Table("user__roles").Where("entity_id = ? and langcode = ?", EntityId_, Langcode_).Limit(limit, offset).Find(_User_roles)
	return _User_roles, err
}

// GetUser_rolesesByEntityIdAndDelta Get User_roleses via EntityIdAndDelta
func GetUser_rolesesByEntityIdAndDelta(offset int, limit int, EntityId_ int, Delta_ int) (*[]*User_roles, error) {
	var _User_roles = new([]*User_roles)
	err := Engine.Table("user__roles").Where("entity_id = ? and delta = ?", EntityId_, Delta_).Limit(limit, offset).Find(_User_roles)
	return _User_roles, err
}

// GetUser_rolesesByEntityIdAndRolesTargetId Get User_roleses via EntityIdAndRolesTargetId
func GetUser_rolesesByEntityIdAndRolesTargetId(offset int, limit int, EntityId_ int, RolesTargetId_ string) (*[]*User_roles, error) {
	var _User_roles = new([]*User_roles)
	err := Engine.Table("user__roles").Where("entity_id = ? and roles_target_id = ?", EntityId_, RolesTargetId_).Limit(limit, offset).Find(_User_roles)
	return _User_roles, err
}

// GetUser_rolesesByRevisionIdAndLangcode Get User_roleses via RevisionIdAndLangcode
func GetUser_rolesesByRevisionIdAndLangcode(offset int, limit int, RevisionId_ int, Langcode_ string) (*[]*User_roles, error) {
	var _User_roles = new([]*User_roles)
	err := Engine.Table("user__roles").Where("revision_id = ? and langcode = ?", RevisionId_, Langcode_).Limit(limit, offset).Find(_User_roles)
	return _User_roles, err
}

// GetUser_rolesesByRevisionIdAndDelta Get User_roleses via RevisionIdAndDelta
func GetUser_rolesesByRevisionIdAndDelta(offset int, limit int, RevisionId_ int, Delta_ int) (*[]*User_roles, error) {
	var _User_roles = new([]*User_roles)
	err := Engine.Table("user__roles").Where("revision_id = ? and delta = ?", RevisionId_, Delta_).Limit(limit, offset).Find(_User_roles)
	return _User_roles, err
}

// GetUser_rolesesByRevisionIdAndRolesTargetId Get User_roleses via RevisionIdAndRolesTargetId
func GetUser_rolesesByRevisionIdAndRolesTargetId(offset int, limit int, RevisionId_ int, RolesTargetId_ string) (*[]*User_roles, error) {
	var _User_roles = new([]*User_roles)
	err := Engine.Table("user__roles").Where("revision_id = ? and roles_target_id = ?", RevisionId_, RolesTargetId_).Limit(limit, offset).Find(_User_roles)
	return _User_roles, err
}

// GetUser_rolesesByLangcodeAndDelta Get User_roleses via LangcodeAndDelta
func GetUser_rolesesByLangcodeAndDelta(offset int, limit int, Langcode_ string, Delta_ int) (*[]*User_roles, error) {
	var _User_roles = new([]*User_roles)
	err := Engine.Table("user__roles").Where("langcode = ? and delta = ?", Langcode_, Delta_).Limit(limit, offset).Find(_User_roles)
	return _User_roles, err
}

// GetUser_rolesesByLangcodeAndRolesTargetId Get User_roleses via LangcodeAndRolesTargetId
func GetUser_rolesesByLangcodeAndRolesTargetId(offset int, limit int, Langcode_ string, RolesTargetId_ string) (*[]*User_roles, error) {
	var _User_roles = new([]*User_roles)
	err := Engine.Table("user__roles").Where("langcode = ? and roles_target_id = ?", Langcode_, RolesTargetId_).Limit(limit, offset).Find(_User_roles)
	return _User_roles, err
}

// GetUser_rolesesByDeltaAndRolesTargetId Get User_roleses via DeltaAndRolesTargetId
func GetUser_rolesesByDeltaAndRolesTargetId(offset int, limit int, Delta_ int, RolesTargetId_ string) (*[]*User_roles, error) {
	var _User_roles = new([]*User_roles)
	err := Engine.Table("user__roles").Where("delta = ? and roles_target_id = ?", Delta_, RolesTargetId_).Limit(limit, offset).Find(_User_roles)
	return _User_roles, err
}

// GetUser_roleses Get User_roleses via field
func GetUser_roleses(offset int, limit int, field string) (*[]*User_roles, error) {
	var _User_roles = new([]*User_roles)
	err := Engine.Table("user__roles").Limit(limit, offset).Desc(field).Find(_User_roles)
	return _User_roles, err
}

// GetUser_rolesesViaBundle Get User_roless via Bundle
func GetUser_rolesesViaBundle(offset int, limit int, Bundle_ string, field string) (*[]*User_roles, error) {
	var _User_roles = new([]*User_roles)
	err := Engine.Table("user__roles").Where("bundle = ?", Bundle_).Limit(limit, offset).Desc(field).Find(_User_roles)
	return _User_roles, err
}

// GetUser_rolesesViaDeleted Get User_roless via Deleted
func GetUser_rolesesViaDeleted(offset int, limit int, Deleted_ int, field string) (*[]*User_roles, error) {
	var _User_roles = new([]*User_roles)
	err := Engine.Table("user__roles").Where("deleted = ?", Deleted_).Limit(limit, offset).Desc(field).Find(_User_roles)
	return _User_roles, err
}

// GetUser_rolesesViaEntityId Get User_roless via EntityId
func GetUser_rolesesViaEntityId(offset int, limit int, EntityId_ int, field string) (*[]*User_roles, error) {
	var _User_roles = new([]*User_roles)
	err := Engine.Table("user__roles").Where("entity_id = ?", EntityId_).Limit(limit, offset).Desc(field).Find(_User_roles)
	return _User_roles, err
}

// GetUser_rolesesViaRevisionId Get User_roless via RevisionId
func GetUser_rolesesViaRevisionId(offset int, limit int, RevisionId_ int, field string) (*[]*User_roles, error) {
	var _User_roles = new([]*User_roles)
	err := Engine.Table("user__roles").Where("revision_id = ?", RevisionId_).Limit(limit, offset).Desc(field).Find(_User_roles)
	return _User_roles, err
}

// GetUser_rolesesViaLangcode Get User_roless via Langcode
func GetUser_rolesesViaLangcode(offset int, limit int, Langcode_ string, field string) (*[]*User_roles, error) {
	var _User_roles = new([]*User_roles)
	err := Engine.Table("user__roles").Where("langcode = ?", Langcode_).Limit(limit, offset).Desc(field).Find(_User_roles)
	return _User_roles, err
}

// GetUser_rolesesViaDelta Get User_roless via Delta
func GetUser_rolesesViaDelta(offset int, limit int, Delta_ int, field string) (*[]*User_roles, error) {
	var _User_roles = new([]*User_roles)
	err := Engine.Table("user__roles").Where("delta = ?", Delta_).Limit(limit, offset).Desc(field).Find(_User_roles)
	return _User_roles, err
}

// GetUser_rolesesViaRolesTargetId Get User_roless via RolesTargetId
func GetUser_rolesesViaRolesTargetId(offset int, limit int, RolesTargetId_ string, field string) (*[]*User_roles, error) {
	var _User_roles = new([]*User_roles)
	err := Engine.Table("user__roles").Where("roles_target_id = ?", RolesTargetId_).Limit(limit, offset).Desc(field).Find(_User_roles)
	return _User_roles, err
}

// HasUser_rolesViaBundle Has User_roles via Bundle
func HasUser_rolesViaBundle(iBundle string) bool {
	if has, err := Engine.Where("bundle = ?", iBundle).Get(new(User_roles)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasUser_rolesViaDeleted Has User_roles via Deleted
func HasUser_rolesViaDeleted(iDeleted int) bool {
	if has, err := Engine.Where("deleted = ?", iDeleted).Get(new(User_roles)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasUser_rolesViaEntityId Has User_roles via EntityId
func HasUser_rolesViaEntityId(iEntityId int) bool {
	if has, err := Engine.Where("entity_id = ?", iEntityId).Get(new(User_roles)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasUser_rolesViaRevisionId Has User_roles via RevisionId
func HasUser_rolesViaRevisionId(iRevisionId int) bool {
	if has, err := Engine.Where("revision_id = ?", iRevisionId).Get(new(User_roles)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasUser_rolesViaLangcode Has User_roles via Langcode
func HasUser_rolesViaLangcode(iLangcode string) bool {
	if has, err := Engine.Where("langcode = ?", iLangcode).Get(new(User_roles)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasUser_rolesViaDelta Has User_roles via Delta
func HasUser_rolesViaDelta(iDelta int) bool {
	if has, err := Engine.Where("delta = ?", iDelta).Get(new(User_roles)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasUser_rolesViaRolesTargetId Has User_roles via RolesTargetId
func HasUser_rolesViaRolesTargetId(iRolesTargetId string) bool {
	if has, err := Engine.Where("roles_target_id = ?", iRolesTargetId).Get(new(User_roles)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// GetUser_rolesViaBundle Get User_roles via Bundle
func GetUser_rolesViaBundle(iBundle string) (*User_roles, error) {
	var _User_roles = &User_roles{Bundle: iBundle}
	has, err := Engine.Get(_User_roles)
	if has {
		return _User_roles, err
	} else {
		return nil, err
	}
}

// GetUser_rolesViaDeleted Get User_roles via Deleted
func GetUser_rolesViaDeleted(iDeleted int) (*User_roles, error) {
	var _User_roles = &User_roles{Deleted: iDeleted}
	has, err := Engine.Get(_User_roles)
	if has {
		return _User_roles, err
	} else {
		return nil, err
	}
}

// GetUser_rolesViaEntityId Get User_roles via EntityId
func GetUser_rolesViaEntityId(iEntityId int) (*User_roles, error) {
	var _User_roles = &User_roles{EntityId: iEntityId}
	has, err := Engine.Get(_User_roles)
	if has {
		return _User_roles, err
	} else {
		return nil, err
	}
}

// GetUser_rolesViaRevisionId Get User_roles via RevisionId
func GetUser_rolesViaRevisionId(iRevisionId int) (*User_roles, error) {
	var _User_roles = &User_roles{RevisionId: iRevisionId}
	has, err := Engine.Get(_User_roles)
	if has {
		return _User_roles, err
	} else {
		return nil, err
	}
}

// GetUser_rolesViaLangcode Get User_roles via Langcode
func GetUser_rolesViaLangcode(iLangcode string) (*User_roles, error) {
	var _User_roles = &User_roles{Langcode: iLangcode}
	has, err := Engine.Get(_User_roles)
	if has {
		return _User_roles, err
	} else {
		return nil, err
	}
}

// GetUser_rolesViaDelta Get User_roles via Delta
func GetUser_rolesViaDelta(iDelta int) (*User_roles, error) {
	var _User_roles = &User_roles{Delta: iDelta}
	has, err := Engine.Get(_User_roles)
	if has {
		return _User_roles, err
	} else {
		return nil, err
	}
}

// GetUser_rolesViaRolesTargetId Get User_roles via RolesTargetId
func GetUser_rolesViaRolesTargetId(iRolesTargetId string) (*User_roles, error) {
	var _User_roles = &User_roles{RolesTargetId: iRolesTargetId}
	has, err := Engine.Get(_User_roles)
	if has {
		return _User_roles, err
	} else {
		return nil, err
	}
}

// SetUser_rolesViaBundle Set User_roles via Bundle
func SetUser_rolesViaBundle(iBundle string, user__roles *User_roles) (int64, error) {
	user__roles.Bundle = iBundle
	return Engine.Insert(user__roles)
}

// SetUser_rolesViaDeleted Set User_roles via Deleted
func SetUser_rolesViaDeleted(iDeleted int, user__roles *User_roles) (int64, error) {
	user__roles.Deleted = iDeleted
	return Engine.Insert(user__roles)
}

// SetUser_rolesViaEntityId Set User_roles via EntityId
func SetUser_rolesViaEntityId(iEntityId int, user__roles *User_roles) (int64, error) {
	user__roles.EntityId = iEntityId
	return Engine.Insert(user__roles)
}

// SetUser_rolesViaRevisionId Set User_roles via RevisionId
func SetUser_rolesViaRevisionId(iRevisionId int, user__roles *User_roles) (int64, error) {
	user__roles.RevisionId = iRevisionId
	return Engine.Insert(user__roles)
}

// SetUser_rolesViaLangcode Set User_roles via Langcode
func SetUser_rolesViaLangcode(iLangcode string, user__roles *User_roles) (int64, error) {
	user__roles.Langcode = iLangcode
	return Engine.Insert(user__roles)
}

// SetUser_rolesViaDelta Set User_roles via Delta
func SetUser_rolesViaDelta(iDelta int, user__roles *User_roles) (int64, error) {
	user__roles.Delta = iDelta
	return Engine.Insert(user__roles)
}

// SetUser_rolesViaRolesTargetId Set User_roles via RolesTargetId
func SetUser_rolesViaRolesTargetId(iRolesTargetId string, user__roles *User_roles) (int64, error) {
	user__roles.RolesTargetId = iRolesTargetId
	return Engine.Insert(user__roles)
}

// AddUser_roles Add User_roles via all columns
func AddUser_roles(iBundle string, iDeleted int, iEntityId int, iRevisionId int, iLangcode string, iDelta int, iRolesTargetId string) error {
	_User_roles := &User_roles{Bundle: iBundle, Deleted: iDeleted, EntityId: iEntityId, RevisionId: iRevisionId, Langcode: iLangcode, Delta: iDelta, RolesTargetId: iRolesTargetId}
	if _, err := Engine.Insert(_User_roles); err != nil {
		return err
	} else {
		return nil
	}
}

// PostUser_roles Post User_roles via iUser_roles
func PostUser_roles(iUser_roles *User_roles) (string, error) {
	_, err := Engine.Insert(iUser_roles)
	return iUser_roles.Bundle, err
}

// PutUser_roles Put User_roles
func PutUser_roles(iUser_roles *User_roles) (string, error) {
	_, err := Engine.Id(iUser_roles.Bundle).Update(iUser_roles)
	return iUser_roles.Bundle, err
}

// PutUser_rolesViaBundle Put User_roles via Bundle
func PutUser_rolesViaBundle(Bundle_ string, iUser_roles *User_roles) (int64, error) {
	row, err := Engine.Update(iUser_roles, &User_roles{Bundle: Bundle_})
	return row, err
}

// PutUser_rolesViaDeleted Put User_roles via Deleted
func PutUser_rolesViaDeleted(Deleted_ int, iUser_roles *User_roles) (int64, error) {
	row, err := Engine.Update(iUser_roles, &User_roles{Deleted: Deleted_})
	return row, err
}

// PutUser_rolesViaEntityId Put User_roles via EntityId
func PutUser_rolesViaEntityId(EntityId_ int, iUser_roles *User_roles) (int64, error) {
	row, err := Engine.Update(iUser_roles, &User_roles{EntityId: EntityId_})
	return row, err
}

// PutUser_rolesViaRevisionId Put User_roles via RevisionId
func PutUser_rolesViaRevisionId(RevisionId_ int, iUser_roles *User_roles) (int64, error) {
	row, err := Engine.Update(iUser_roles, &User_roles{RevisionId: RevisionId_})
	return row, err
}

// PutUser_rolesViaLangcode Put User_roles via Langcode
func PutUser_rolesViaLangcode(Langcode_ string, iUser_roles *User_roles) (int64, error) {
	row, err := Engine.Update(iUser_roles, &User_roles{Langcode: Langcode_})
	return row, err
}

// PutUser_rolesViaDelta Put User_roles via Delta
func PutUser_rolesViaDelta(Delta_ int, iUser_roles *User_roles) (int64, error) {
	row, err := Engine.Update(iUser_roles, &User_roles{Delta: Delta_})
	return row, err
}

// PutUser_rolesViaRolesTargetId Put User_roles via RolesTargetId
func PutUser_rolesViaRolesTargetId(RolesTargetId_ string, iUser_roles *User_roles) (int64, error) {
	row, err := Engine.Update(iUser_roles, &User_roles{RolesTargetId: RolesTargetId_})
	return row, err
}

// UpdateUser_rolesViaBundle via map[string]interface{}{}
func UpdateUser_rolesViaBundle(iBundle string, iUser_rolesMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(User_roles)).Where("bundle = ?", iBundle).Update(iUser_rolesMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateUser_rolesViaDeleted via map[string]interface{}{}
func UpdateUser_rolesViaDeleted(iDeleted int, iUser_rolesMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(User_roles)).Where("deleted = ?", iDeleted).Update(iUser_rolesMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateUser_rolesViaEntityId via map[string]interface{}{}
func UpdateUser_rolesViaEntityId(iEntityId int, iUser_rolesMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(User_roles)).Where("entity_id = ?", iEntityId).Update(iUser_rolesMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateUser_rolesViaRevisionId via map[string]interface{}{}
func UpdateUser_rolesViaRevisionId(iRevisionId int, iUser_rolesMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(User_roles)).Where("revision_id = ?", iRevisionId).Update(iUser_rolesMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateUser_rolesViaLangcode via map[string]interface{}{}
func UpdateUser_rolesViaLangcode(iLangcode string, iUser_rolesMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(User_roles)).Where("langcode = ?", iLangcode).Update(iUser_rolesMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateUser_rolesViaDelta via map[string]interface{}{}
func UpdateUser_rolesViaDelta(iDelta int, iUser_rolesMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(User_roles)).Where("delta = ?", iDelta).Update(iUser_rolesMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateUser_rolesViaRolesTargetId via map[string]interface{}{}
func UpdateUser_rolesViaRolesTargetId(iRolesTargetId string, iUser_rolesMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(User_roles)).Where("roles_target_id = ?", iRolesTargetId).Update(iUser_rolesMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// DeleteUser_roles Delete User_roles
func DeleteUser_roles(iBundle string) (int64, error) {
	row, err := Engine.Id(iBundle).Delete(new(User_roles))
	return row, err
}

// DeleteUser_rolesViaBundle Delete User_roles via Bundle
func DeleteUser_rolesViaBundle(iBundle string) (err error) {
	var has bool
	var _User_roles = &User_roles{Bundle: iBundle}
	if has, err = Engine.Get(_User_roles); (has == true) && (err == nil) {
		if row, err := Engine.Where("bundle = ?", iBundle).Delete(new(User_roles)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteUser_rolesViaDeleted Delete User_roles via Deleted
func DeleteUser_rolesViaDeleted(iDeleted int) (err error) {
	var has bool
	var _User_roles = &User_roles{Deleted: iDeleted}
	if has, err = Engine.Get(_User_roles); (has == true) && (err == nil) {
		if row, err := Engine.Where("deleted = ?", iDeleted).Delete(new(User_roles)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteUser_rolesViaEntityId Delete User_roles via EntityId
func DeleteUser_rolesViaEntityId(iEntityId int) (err error) {
	var has bool
	var _User_roles = &User_roles{EntityId: iEntityId}
	if has, err = Engine.Get(_User_roles); (has == true) && (err == nil) {
		if row, err := Engine.Where("entity_id = ?", iEntityId).Delete(new(User_roles)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteUser_rolesViaRevisionId Delete User_roles via RevisionId
func DeleteUser_rolesViaRevisionId(iRevisionId int) (err error) {
	var has bool
	var _User_roles = &User_roles{RevisionId: iRevisionId}
	if has, err = Engine.Get(_User_roles); (has == true) && (err == nil) {
		if row, err := Engine.Where("revision_id = ?", iRevisionId).Delete(new(User_roles)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteUser_rolesViaLangcode Delete User_roles via Langcode
func DeleteUser_rolesViaLangcode(iLangcode string) (err error) {
	var has bool
	var _User_roles = &User_roles{Langcode: iLangcode}
	if has, err = Engine.Get(_User_roles); (has == true) && (err == nil) {
		if row, err := Engine.Where("langcode = ?", iLangcode).Delete(new(User_roles)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteUser_rolesViaDelta Delete User_roles via Delta
func DeleteUser_rolesViaDelta(iDelta int) (err error) {
	var has bool
	var _User_roles = &User_roles{Delta: iDelta}
	if has, err = Engine.Get(_User_roles); (has == true) && (err == nil) {
		if row, err := Engine.Where("delta = ?", iDelta).Delete(new(User_roles)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteUser_rolesViaRolesTargetId Delete User_roles via RolesTargetId
func DeleteUser_rolesViaRolesTargetId(iRolesTargetId string) (err error) {
	var has bool
	var _User_roles = &User_roles{RolesTargetId: iRolesTargetId}
	if has, err = Engine.Get(_User_roles); (has == true) && (err == nil) {
		if row, err := Engine.Where("roles_target_id = ?", iRolesTargetId).Delete(new(User_roles)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}
