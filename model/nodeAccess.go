package model

type NodeAccess struct {
	Nid         int64  `xorm:"not null pk default 0 INT(10)"`
	Langcode    string `xorm:"not null pk default '' VARCHAR(12)"`
	Fallback    int    `xorm:"not null default 1 INT(10)"`
	Gid         int    `xorm:"not null pk default 0 INT(10)"`
	Realm       string `xorm:"not null pk default '' VARCHAR(255)"`
	GrantView   int    `xorm:"not null default 0 TINYINT(3)"`
	GrantUpdate int    `xorm:"not null default 0 TINYINT(3)"`
	GrantDelete int    `xorm:"not null default 0 TINYINT(3)"`
}

// GetNodeAccessesCount NodeAccesss' Count
func GetNodeAccessesCount(offset int, limit int) (int64, error) {
	total, err := Engine.Limit(limit, offset).Count(&NodeAccess{})
	return total, err
}

// GetNodeAccessCountViaNid Get NodeAccess via Nid
func GetNodeAccessCountViaNid(iNid int64) int64 {
	n, _ := Engine.Where("nid = ?", iNid).Count(&NodeAccess{Nid: iNid})
	return n
}

// GetNodeAccessCountViaLangcode Get NodeAccess via Langcode
func GetNodeAccessCountViaLangcode(iLangcode string) int64 {
	n, _ := Engine.Where("langcode = ?", iLangcode).Count(&NodeAccess{Langcode: iLangcode})
	return n
}

// GetNodeAccessCountViaFallback Get NodeAccess via Fallback
func GetNodeAccessCountViaFallback(iFallback int) int64 {
	n, _ := Engine.Where("fallback = ?", iFallback).Count(&NodeAccess{Fallback: iFallback})
	return n
}

// GetNodeAccessCountViaGid Get NodeAccess via Gid
func GetNodeAccessCountViaGid(iGid int) int64 {
	n, _ := Engine.Where("gid = ?", iGid).Count(&NodeAccess{Gid: iGid})
	return n
}

// GetNodeAccessCountViaRealm Get NodeAccess via Realm
func GetNodeAccessCountViaRealm(iRealm string) int64 {
	n, _ := Engine.Where("realm = ?", iRealm).Count(&NodeAccess{Realm: iRealm})
	return n
}

// GetNodeAccessCountViaGrantView Get NodeAccess via GrantView
func GetNodeAccessCountViaGrantView(iGrantView int) int64 {
	n, _ := Engine.Where("grant_view = ?", iGrantView).Count(&NodeAccess{GrantView: iGrantView})
	return n
}

// GetNodeAccessCountViaGrantUpdate Get NodeAccess via GrantUpdate
func GetNodeAccessCountViaGrantUpdate(iGrantUpdate int) int64 {
	n, _ := Engine.Where("grant_update = ?", iGrantUpdate).Count(&NodeAccess{GrantUpdate: iGrantUpdate})
	return n
}

// GetNodeAccessCountViaGrantDelete Get NodeAccess via GrantDelete
func GetNodeAccessCountViaGrantDelete(iGrantDelete int) int64 {
	n, _ := Engine.Where("grant_delete = ?", iGrantDelete).Count(&NodeAccess{GrantDelete: iGrantDelete})
	return n
}

// GetNodeAccessesByNidAndLangcodeAndFallback Get NodeAccesses via NidAndLangcodeAndFallback
func GetNodeAccessesByNidAndLangcodeAndFallback(offset int, limit int, Nid_ int64, Langcode_ string, Fallback_ int) (*[]*NodeAccess, error) {
	var _NodeAccess = new([]*NodeAccess)
	err := Engine.Table("node_access").Where("nid = ? and langcode = ? and fallback = ?", Nid_, Langcode_, Fallback_).Limit(limit, offset).Find(_NodeAccess)
	return _NodeAccess, err
}

// GetNodeAccessesByNidAndLangcodeAndGid Get NodeAccesses via NidAndLangcodeAndGid
func GetNodeAccessesByNidAndLangcodeAndGid(offset int, limit int, Nid_ int64, Langcode_ string, Gid_ int) (*[]*NodeAccess, error) {
	var _NodeAccess = new([]*NodeAccess)
	err := Engine.Table("node_access").Where("nid = ? and langcode = ? and gid = ?", Nid_, Langcode_, Gid_).Limit(limit, offset).Find(_NodeAccess)
	return _NodeAccess, err
}

// GetNodeAccessesByNidAndLangcodeAndRealm Get NodeAccesses via NidAndLangcodeAndRealm
func GetNodeAccessesByNidAndLangcodeAndRealm(offset int, limit int, Nid_ int64, Langcode_ string, Realm_ string) (*[]*NodeAccess, error) {
	var _NodeAccess = new([]*NodeAccess)
	err := Engine.Table("node_access").Where("nid = ? and langcode = ? and realm = ?", Nid_, Langcode_, Realm_).Limit(limit, offset).Find(_NodeAccess)
	return _NodeAccess, err
}

// GetNodeAccessesByNidAndLangcodeAndGrantView Get NodeAccesses via NidAndLangcodeAndGrantView
func GetNodeAccessesByNidAndLangcodeAndGrantView(offset int, limit int, Nid_ int64, Langcode_ string, GrantView_ int) (*[]*NodeAccess, error) {
	var _NodeAccess = new([]*NodeAccess)
	err := Engine.Table("node_access").Where("nid = ? and langcode = ? and grant_view = ?", Nid_, Langcode_, GrantView_).Limit(limit, offset).Find(_NodeAccess)
	return _NodeAccess, err
}

// GetNodeAccessesByNidAndLangcodeAndGrantUpdate Get NodeAccesses via NidAndLangcodeAndGrantUpdate
func GetNodeAccessesByNidAndLangcodeAndGrantUpdate(offset int, limit int, Nid_ int64, Langcode_ string, GrantUpdate_ int) (*[]*NodeAccess, error) {
	var _NodeAccess = new([]*NodeAccess)
	err := Engine.Table("node_access").Where("nid = ? and langcode = ? and grant_update = ?", Nid_, Langcode_, GrantUpdate_).Limit(limit, offset).Find(_NodeAccess)
	return _NodeAccess, err
}

// GetNodeAccessesByNidAndLangcodeAndGrantDelete Get NodeAccesses via NidAndLangcodeAndGrantDelete
func GetNodeAccessesByNidAndLangcodeAndGrantDelete(offset int, limit int, Nid_ int64, Langcode_ string, GrantDelete_ int) (*[]*NodeAccess, error) {
	var _NodeAccess = new([]*NodeAccess)
	err := Engine.Table("node_access").Where("nid = ? and langcode = ? and grant_delete = ?", Nid_, Langcode_, GrantDelete_).Limit(limit, offset).Find(_NodeAccess)
	return _NodeAccess, err
}

// GetNodeAccessesByNidAndFallbackAndGid Get NodeAccesses via NidAndFallbackAndGid
func GetNodeAccessesByNidAndFallbackAndGid(offset int, limit int, Nid_ int64, Fallback_ int, Gid_ int) (*[]*NodeAccess, error) {
	var _NodeAccess = new([]*NodeAccess)
	err := Engine.Table("node_access").Where("nid = ? and fallback = ? and gid = ?", Nid_, Fallback_, Gid_).Limit(limit, offset).Find(_NodeAccess)
	return _NodeAccess, err
}

// GetNodeAccessesByNidAndFallbackAndRealm Get NodeAccesses via NidAndFallbackAndRealm
func GetNodeAccessesByNidAndFallbackAndRealm(offset int, limit int, Nid_ int64, Fallback_ int, Realm_ string) (*[]*NodeAccess, error) {
	var _NodeAccess = new([]*NodeAccess)
	err := Engine.Table("node_access").Where("nid = ? and fallback = ? and realm = ?", Nid_, Fallback_, Realm_).Limit(limit, offset).Find(_NodeAccess)
	return _NodeAccess, err
}

// GetNodeAccessesByNidAndFallbackAndGrantView Get NodeAccesses via NidAndFallbackAndGrantView
func GetNodeAccessesByNidAndFallbackAndGrantView(offset int, limit int, Nid_ int64, Fallback_ int, GrantView_ int) (*[]*NodeAccess, error) {
	var _NodeAccess = new([]*NodeAccess)
	err := Engine.Table("node_access").Where("nid = ? and fallback = ? and grant_view = ?", Nid_, Fallback_, GrantView_).Limit(limit, offset).Find(_NodeAccess)
	return _NodeAccess, err
}

// GetNodeAccessesByNidAndFallbackAndGrantUpdate Get NodeAccesses via NidAndFallbackAndGrantUpdate
func GetNodeAccessesByNidAndFallbackAndGrantUpdate(offset int, limit int, Nid_ int64, Fallback_ int, GrantUpdate_ int) (*[]*NodeAccess, error) {
	var _NodeAccess = new([]*NodeAccess)
	err := Engine.Table("node_access").Where("nid = ? and fallback = ? and grant_update = ?", Nid_, Fallback_, GrantUpdate_).Limit(limit, offset).Find(_NodeAccess)
	return _NodeAccess, err
}

// GetNodeAccessesByNidAndFallbackAndGrantDelete Get NodeAccesses via NidAndFallbackAndGrantDelete
func GetNodeAccessesByNidAndFallbackAndGrantDelete(offset int, limit int, Nid_ int64, Fallback_ int, GrantDelete_ int) (*[]*NodeAccess, error) {
	var _NodeAccess = new([]*NodeAccess)
	err := Engine.Table("node_access").Where("nid = ? and fallback = ? and grant_delete = ?", Nid_, Fallback_, GrantDelete_).Limit(limit, offset).Find(_NodeAccess)
	return _NodeAccess, err
}

// GetNodeAccessesByNidAndGidAndRealm Get NodeAccesses via NidAndGidAndRealm
func GetNodeAccessesByNidAndGidAndRealm(offset int, limit int, Nid_ int64, Gid_ int, Realm_ string) (*[]*NodeAccess, error) {
	var _NodeAccess = new([]*NodeAccess)
	err := Engine.Table("node_access").Where("nid = ? and gid = ? and realm = ?", Nid_, Gid_, Realm_).Limit(limit, offset).Find(_NodeAccess)
	return _NodeAccess, err
}

// GetNodeAccessesByNidAndGidAndGrantView Get NodeAccesses via NidAndGidAndGrantView
func GetNodeAccessesByNidAndGidAndGrantView(offset int, limit int, Nid_ int64, Gid_ int, GrantView_ int) (*[]*NodeAccess, error) {
	var _NodeAccess = new([]*NodeAccess)
	err := Engine.Table("node_access").Where("nid = ? and gid = ? and grant_view = ?", Nid_, Gid_, GrantView_).Limit(limit, offset).Find(_NodeAccess)
	return _NodeAccess, err
}

// GetNodeAccessesByNidAndGidAndGrantUpdate Get NodeAccesses via NidAndGidAndGrantUpdate
func GetNodeAccessesByNidAndGidAndGrantUpdate(offset int, limit int, Nid_ int64, Gid_ int, GrantUpdate_ int) (*[]*NodeAccess, error) {
	var _NodeAccess = new([]*NodeAccess)
	err := Engine.Table("node_access").Where("nid = ? and gid = ? and grant_update = ?", Nid_, Gid_, GrantUpdate_).Limit(limit, offset).Find(_NodeAccess)
	return _NodeAccess, err
}

// GetNodeAccessesByNidAndGidAndGrantDelete Get NodeAccesses via NidAndGidAndGrantDelete
func GetNodeAccessesByNidAndGidAndGrantDelete(offset int, limit int, Nid_ int64, Gid_ int, GrantDelete_ int) (*[]*NodeAccess, error) {
	var _NodeAccess = new([]*NodeAccess)
	err := Engine.Table("node_access").Where("nid = ? and gid = ? and grant_delete = ?", Nid_, Gid_, GrantDelete_).Limit(limit, offset).Find(_NodeAccess)
	return _NodeAccess, err
}

// GetNodeAccessesByNidAndRealmAndGrantView Get NodeAccesses via NidAndRealmAndGrantView
func GetNodeAccessesByNidAndRealmAndGrantView(offset int, limit int, Nid_ int64, Realm_ string, GrantView_ int) (*[]*NodeAccess, error) {
	var _NodeAccess = new([]*NodeAccess)
	err := Engine.Table("node_access").Where("nid = ? and realm = ? and grant_view = ?", Nid_, Realm_, GrantView_).Limit(limit, offset).Find(_NodeAccess)
	return _NodeAccess, err
}

// GetNodeAccessesByNidAndRealmAndGrantUpdate Get NodeAccesses via NidAndRealmAndGrantUpdate
func GetNodeAccessesByNidAndRealmAndGrantUpdate(offset int, limit int, Nid_ int64, Realm_ string, GrantUpdate_ int) (*[]*NodeAccess, error) {
	var _NodeAccess = new([]*NodeAccess)
	err := Engine.Table("node_access").Where("nid = ? and realm = ? and grant_update = ?", Nid_, Realm_, GrantUpdate_).Limit(limit, offset).Find(_NodeAccess)
	return _NodeAccess, err
}

// GetNodeAccessesByNidAndRealmAndGrantDelete Get NodeAccesses via NidAndRealmAndGrantDelete
func GetNodeAccessesByNidAndRealmAndGrantDelete(offset int, limit int, Nid_ int64, Realm_ string, GrantDelete_ int) (*[]*NodeAccess, error) {
	var _NodeAccess = new([]*NodeAccess)
	err := Engine.Table("node_access").Where("nid = ? and realm = ? and grant_delete = ?", Nid_, Realm_, GrantDelete_).Limit(limit, offset).Find(_NodeAccess)
	return _NodeAccess, err
}

// GetNodeAccessesByNidAndGrantViewAndGrantUpdate Get NodeAccesses via NidAndGrantViewAndGrantUpdate
func GetNodeAccessesByNidAndGrantViewAndGrantUpdate(offset int, limit int, Nid_ int64, GrantView_ int, GrantUpdate_ int) (*[]*NodeAccess, error) {
	var _NodeAccess = new([]*NodeAccess)
	err := Engine.Table("node_access").Where("nid = ? and grant_view = ? and grant_update = ?", Nid_, GrantView_, GrantUpdate_).Limit(limit, offset).Find(_NodeAccess)
	return _NodeAccess, err
}

// GetNodeAccessesByNidAndGrantViewAndGrantDelete Get NodeAccesses via NidAndGrantViewAndGrantDelete
func GetNodeAccessesByNidAndGrantViewAndGrantDelete(offset int, limit int, Nid_ int64, GrantView_ int, GrantDelete_ int) (*[]*NodeAccess, error) {
	var _NodeAccess = new([]*NodeAccess)
	err := Engine.Table("node_access").Where("nid = ? and grant_view = ? and grant_delete = ?", Nid_, GrantView_, GrantDelete_).Limit(limit, offset).Find(_NodeAccess)
	return _NodeAccess, err
}

// GetNodeAccessesByNidAndGrantUpdateAndGrantDelete Get NodeAccesses via NidAndGrantUpdateAndGrantDelete
func GetNodeAccessesByNidAndGrantUpdateAndGrantDelete(offset int, limit int, Nid_ int64, GrantUpdate_ int, GrantDelete_ int) (*[]*NodeAccess, error) {
	var _NodeAccess = new([]*NodeAccess)
	err := Engine.Table("node_access").Where("nid = ? and grant_update = ? and grant_delete = ?", Nid_, GrantUpdate_, GrantDelete_).Limit(limit, offset).Find(_NodeAccess)
	return _NodeAccess, err
}

// GetNodeAccessesByLangcodeAndFallbackAndGid Get NodeAccesses via LangcodeAndFallbackAndGid
func GetNodeAccessesByLangcodeAndFallbackAndGid(offset int, limit int, Langcode_ string, Fallback_ int, Gid_ int) (*[]*NodeAccess, error) {
	var _NodeAccess = new([]*NodeAccess)
	err := Engine.Table("node_access").Where("langcode = ? and fallback = ? and gid = ?", Langcode_, Fallback_, Gid_).Limit(limit, offset).Find(_NodeAccess)
	return _NodeAccess, err
}

// GetNodeAccessesByLangcodeAndFallbackAndRealm Get NodeAccesses via LangcodeAndFallbackAndRealm
func GetNodeAccessesByLangcodeAndFallbackAndRealm(offset int, limit int, Langcode_ string, Fallback_ int, Realm_ string) (*[]*NodeAccess, error) {
	var _NodeAccess = new([]*NodeAccess)
	err := Engine.Table("node_access").Where("langcode = ? and fallback = ? and realm = ?", Langcode_, Fallback_, Realm_).Limit(limit, offset).Find(_NodeAccess)
	return _NodeAccess, err
}

// GetNodeAccessesByLangcodeAndFallbackAndGrantView Get NodeAccesses via LangcodeAndFallbackAndGrantView
func GetNodeAccessesByLangcodeAndFallbackAndGrantView(offset int, limit int, Langcode_ string, Fallback_ int, GrantView_ int) (*[]*NodeAccess, error) {
	var _NodeAccess = new([]*NodeAccess)
	err := Engine.Table("node_access").Where("langcode = ? and fallback = ? and grant_view = ?", Langcode_, Fallback_, GrantView_).Limit(limit, offset).Find(_NodeAccess)
	return _NodeAccess, err
}

// GetNodeAccessesByLangcodeAndFallbackAndGrantUpdate Get NodeAccesses via LangcodeAndFallbackAndGrantUpdate
func GetNodeAccessesByLangcodeAndFallbackAndGrantUpdate(offset int, limit int, Langcode_ string, Fallback_ int, GrantUpdate_ int) (*[]*NodeAccess, error) {
	var _NodeAccess = new([]*NodeAccess)
	err := Engine.Table("node_access").Where("langcode = ? and fallback = ? and grant_update = ?", Langcode_, Fallback_, GrantUpdate_).Limit(limit, offset).Find(_NodeAccess)
	return _NodeAccess, err
}

// GetNodeAccessesByLangcodeAndFallbackAndGrantDelete Get NodeAccesses via LangcodeAndFallbackAndGrantDelete
func GetNodeAccessesByLangcodeAndFallbackAndGrantDelete(offset int, limit int, Langcode_ string, Fallback_ int, GrantDelete_ int) (*[]*NodeAccess, error) {
	var _NodeAccess = new([]*NodeAccess)
	err := Engine.Table("node_access").Where("langcode = ? and fallback = ? and grant_delete = ?", Langcode_, Fallback_, GrantDelete_).Limit(limit, offset).Find(_NodeAccess)
	return _NodeAccess, err
}

// GetNodeAccessesByLangcodeAndGidAndRealm Get NodeAccesses via LangcodeAndGidAndRealm
func GetNodeAccessesByLangcodeAndGidAndRealm(offset int, limit int, Langcode_ string, Gid_ int, Realm_ string) (*[]*NodeAccess, error) {
	var _NodeAccess = new([]*NodeAccess)
	err := Engine.Table("node_access").Where("langcode = ? and gid = ? and realm = ?", Langcode_, Gid_, Realm_).Limit(limit, offset).Find(_NodeAccess)
	return _NodeAccess, err
}

// GetNodeAccessesByLangcodeAndGidAndGrantView Get NodeAccesses via LangcodeAndGidAndGrantView
func GetNodeAccessesByLangcodeAndGidAndGrantView(offset int, limit int, Langcode_ string, Gid_ int, GrantView_ int) (*[]*NodeAccess, error) {
	var _NodeAccess = new([]*NodeAccess)
	err := Engine.Table("node_access").Where("langcode = ? and gid = ? and grant_view = ?", Langcode_, Gid_, GrantView_).Limit(limit, offset).Find(_NodeAccess)
	return _NodeAccess, err
}

// GetNodeAccessesByLangcodeAndGidAndGrantUpdate Get NodeAccesses via LangcodeAndGidAndGrantUpdate
func GetNodeAccessesByLangcodeAndGidAndGrantUpdate(offset int, limit int, Langcode_ string, Gid_ int, GrantUpdate_ int) (*[]*NodeAccess, error) {
	var _NodeAccess = new([]*NodeAccess)
	err := Engine.Table("node_access").Where("langcode = ? and gid = ? and grant_update = ?", Langcode_, Gid_, GrantUpdate_).Limit(limit, offset).Find(_NodeAccess)
	return _NodeAccess, err
}

// GetNodeAccessesByLangcodeAndGidAndGrantDelete Get NodeAccesses via LangcodeAndGidAndGrantDelete
func GetNodeAccessesByLangcodeAndGidAndGrantDelete(offset int, limit int, Langcode_ string, Gid_ int, GrantDelete_ int) (*[]*NodeAccess, error) {
	var _NodeAccess = new([]*NodeAccess)
	err := Engine.Table("node_access").Where("langcode = ? and gid = ? and grant_delete = ?", Langcode_, Gid_, GrantDelete_).Limit(limit, offset).Find(_NodeAccess)
	return _NodeAccess, err
}

// GetNodeAccessesByLangcodeAndRealmAndGrantView Get NodeAccesses via LangcodeAndRealmAndGrantView
func GetNodeAccessesByLangcodeAndRealmAndGrantView(offset int, limit int, Langcode_ string, Realm_ string, GrantView_ int) (*[]*NodeAccess, error) {
	var _NodeAccess = new([]*NodeAccess)
	err := Engine.Table("node_access").Where("langcode = ? and realm = ? and grant_view = ?", Langcode_, Realm_, GrantView_).Limit(limit, offset).Find(_NodeAccess)
	return _NodeAccess, err
}

// GetNodeAccessesByLangcodeAndRealmAndGrantUpdate Get NodeAccesses via LangcodeAndRealmAndGrantUpdate
func GetNodeAccessesByLangcodeAndRealmAndGrantUpdate(offset int, limit int, Langcode_ string, Realm_ string, GrantUpdate_ int) (*[]*NodeAccess, error) {
	var _NodeAccess = new([]*NodeAccess)
	err := Engine.Table("node_access").Where("langcode = ? and realm = ? and grant_update = ?", Langcode_, Realm_, GrantUpdate_).Limit(limit, offset).Find(_NodeAccess)
	return _NodeAccess, err
}

// GetNodeAccessesByLangcodeAndRealmAndGrantDelete Get NodeAccesses via LangcodeAndRealmAndGrantDelete
func GetNodeAccessesByLangcodeAndRealmAndGrantDelete(offset int, limit int, Langcode_ string, Realm_ string, GrantDelete_ int) (*[]*NodeAccess, error) {
	var _NodeAccess = new([]*NodeAccess)
	err := Engine.Table("node_access").Where("langcode = ? and realm = ? and grant_delete = ?", Langcode_, Realm_, GrantDelete_).Limit(limit, offset).Find(_NodeAccess)
	return _NodeAccess, err
}

// GetNodeAccessesByLangcodeAndGrantViewAndGrantUpdate Get NodeAccesses via LangcodeAndGrantViewAndGrantUpdate
func GetNodeAccessesByLangcodeAndGrantViewAndGrantUpdate(offset int, limit int, Langcode_ string, GrantView_ int, GrantUpdate_ int) (*[]*NodeAccess, error) {
	var _NodeAccess = new([]*NodeAccess)
	err := Engine.Table("node_access").Where("langcode = ? and grant_view = ? and grant_update = ?", Langcode_, GrantView_, GrantUpdate_).Limit(limit, offset).Find(_NodeAccess)
	return _NodeAccess, err
}

// GetNodeAccessesByLangcodeAndGrantViewAndGrantDelete Get NodeAccesses via LangcodeAndGrantViewAndGrantDelete
func GetNodeAccessesByLangcodeAndGrantViewAndGrantDelete(offset int, limit int, Langcode_ string, GrantView_ int, GrantDelete_ int) (*[]*NodeAccess, error) {
	var _NodeAccess = new([]*NodeAccess)
	err := Engine.Table("node_access").Where("langcode = ? and grant_view = ? and grant_delete = ?", Langcode_, GrantView_, GrantDelete_).Limit(limit, offset).Find(_NodeAccess)
	return _NodeAccess, err
}

// GetNodeAccessesByLangcodeAndGrantUpdateAndGrantDelete Get NodeAccesses via LangcodeAndGrantUpdateAndGrantDelete
func GetNodeAccessesByLangcodeAndGrantUpdateAndGrantDelete(offset int, limit int, Langcode_ string, GrantUpdate_ int, GrantDelete_ int) (*[]*NodeAccess, error) {
	var _NodeAccess = new([]*NodeAccess)
	err := Engine.Table("node_access").Where("langcode = ? and grant_update = ? and grant_delete = ?", Langcode_, GrantUpdate_, GrantDelete_).Limit(limit, offset).Find(_NodeAccess)
	return _NodeAccess, err
}

// GetNodeAccessesByFallbackAndGidAndRealm Get NodeAccesses via FallbackAndGidAndRealm
func GetNodeAccessesByFallbackAndGidAndRealm(offset int, limit int, Fallback_ int, Gid_ int, Realm_ string) (*[]*NodeAccess, error) {
	var _NodeAccess = new([]*NodeAccess)
	err := Engine.Table("node_access").Where("fallback = ? and gid = ? and realm = ?", Fallback_, Gid_, Realm_).Limit(limit, offset).Find(_NodeAccess)
	return _NodeAccess, err
}

// GetNodeAccessesByFallbackAndGidAndGrantView Get NodeAccesses via FallbackAndGidAndGrantView
func GetNodeAccessesByFallbackAndGidAndGrantView(offset int, limit int, Fallback_ int, Gid_ int, GrantView_ int) (*[]*NodeAccess, error) {
	var _NodeAccess = new([]*NodeAccess)
	err := Engine.Table("node_access").Where("fallback = ? and gid = ? and grant_view = ?", Fallback_, Gid_, GrantView_).Limit(limit, offset).Find(_NodeAccess)
	return _NodeAccess, err
}

// GetNodeAccessesByFallbackAndGidAndGrantUpdate Get NodeAccesses via FallbackAndGidAndGrantUpdate
func GetNodeAccessesByFallbackAndGidAndGrantUpdate(offset int, limit int, Fallback_ int, Gid_ int, GrantUpdate_ int) (*[]*NodeAccess, error) {
	var _NodeAccess = new([]*NodeAccess)
	err := Engine.Table("node_access").Where("fallback = ? and gid = ? and grant_update = ?", Fallback_, Gid_, GrantUpdate_).Limit(limit, offset).Find(_NodeAccess)
	return _NodeAccess, err
}

// GetNodeAccessesByFallbackAndGidAndGrantDelete Get NodeAccesses via FallbackAndGidAndGrantDelete
func GetNodeAccessesByFallbackAndGidAndGrantDelete(offset int, limit int, Fallback_ int, Gid_ int, GrantDelete_ int) (*[]*NodeAccess, error) {
	var _NodeAccess = new([]*NodeAccess)
	err := Engine.Table("node_access").Where("fallback = ? and gid = ? and grant_delete = ?", Fallback_, Gid_, GrantDelete_).Limit(limit, offset).Find(_NodeAccess)
	return _NodeAccess, err
}

// GetNodeAccessesByFallbackAndRealmAndGrantView Get NodeAccesses via FallbackAndRealmAndGrantView
func GetNodeAccessesByFallbackAndRealmAndGrantView(offset int, limit int, Fallback_ int, Realm_ string, GrantView_ int) (*[]*NodeAccess, error) {
	var _NodeAccess = new([]*NodeAccess)
	err := Engine.Table("node_access").Where("fallback = ? and realm = ? and grant_view = ?", Fallback_, Realm_, GrantView_).Limit(limit, offset).Find(_NodeAccess)
	return _NodeAccess, err
}

// GetNodeAccessesByFallbackAndRealmAndGrantUpdate Get NodeAccesses via FallbackAndRealmAndGrantUpdate
func GetNodeAccessesByFallbackAndRealmAndGrantUpdate(offset int, limit int, Fallback_ int, Realm_ string, GrantUpdate_ int) (*[]*NodeAccess, error) {
	var _NodeAccess = new([]*NodeAccess)
	err := Engine.Table("node_access").Where("fallback = ? and realm = ? and grant_update = ?", Fallback_, Realm_, GrantUpdate_).Limit(limit, offset).Find(_NodeAccess)
	return _NodeAccess, err
}

// GetNodeAccessesByFallbackAndRealmAndGrantDelete Get NodeAccesses via FallbackAndRealmAndGrantDelete
func GetNodeAccessesByFallbackAndRealmAndGrantDelete(offset int, limit int, Fallback_ int, Realm_ string, GrantDelete_ int) (*[]*NodeAccess, error) {
	var _NodeAccess = new([]*NodeAccess)
	err := Engine.Table("node_access").Where("fallback = ? and realm = ? and grant_delete = ?", Fallback_, Realm_, GrantDelete_).Limit(limit, offset).Find(_NodeAccess)
	return _NodeAccess, err
}

// GetNodeAccessesByFallbackAndGrantViewAndGrantUpdate Get NodeAccesses via FallbackAndGrantViewAndGrantUpdate
func GetNodeAccessesByFallbackAndGrantViewAndGrantUpdate(offset int, limit int, Fallback_ int, GrantView_ int, GrantUpdate_ int) (*[]*NodeAccess, error) {
	var _NodeAccess = new([]*NodeAccess)
	err := Engine.Table("node_access").Where("fallback = ? and grant_view = ? and grant_update = ?", Fallback_, GrantView_, GrantUpdate_).Limit(limit, offset).Find(_NodeAccess)
	return _NodeAccess, err
}

// GetNodeAccessesByFallbackAndGrantViewAndGrantDelete Get NodeAccesses via FallbackAndGrantViewAndGrantDelete
func GetNodeAccessesByFallbackAndGrantViewAndGrantDelete(offset int, limit int, Fallback_ int, GrantView_ int, GrantDelete_ int) (*[]*NodeAccess, error) {
	var _NodeAccess = new([]*NodeAccess)
	err := Engine.Table("node_access").Where("fallback = ? and grant_view = ? and grant_delete = ?", Fallback_, GrantView_, GrantDelete_).Limit(limit, offset).Find(_NodeAccess)
	return _NodeAccess, err
}

// GetNodeAccessesByFallbackAndGrantUpdateAndGrantDelete Get NodeAccesses via FallbackAndGrantUpdateAndGrantDelete
func GetNodeAccessesByFallbackAndGrantUpdateAndGrantDelete(offset int, limit int, Fallback_ int, GrantUpdate_ int, GrantDelete_ int) (*[]*NodeAccess, error) {
	var _NodeAccess = new([]*NodeAccess)
	err := Engine.Table("node_access").Where("fallback = ? and grant_update = ? and grant_delete = ?", Fallback_, GrantUpdate_, GrantDelete_).Limit(limit, offset).Find(_NodeAccess)
	return _NodeAccess, err
}

// GetNodeAccessesByGidAndRealmAndGrantView Get NodeAccesses via GidAndRealmAndGrantView
func GetNodeAccessesByGidAndRealmAndGrantView(offset int, limit int, Gid_ int, Realm_ string, GrantView_ int) (*[]*NodeAccess, error) {
	var _NodeAccess = new([]*NodeAccess)
	err := Engine.Table("node_access").Where("gid = ? and realm = ? and grant_view = ?", Gid_, Realm_, GrantView_).Limit(limit, offset).Find(_NodeAccess)
	return _NodeAccess, err
}

// GetNodeAccessesByGidAndRealmAndGrantUpdate Get NodeAccesses via GidAndRealmAndGrantUpdate
func GetNodeAccessesByGidAndRealmAndGrantUpdate(offset int, limit int, Gid_ int, Realm_ string, GrantUpdate_ int) (*[]*NodeAccess, error) {
	var _NodeAccess = new([]*NodeAccess)
	err := Engine.Table("node_access").Where("gid = ? and realm = ? and grant_update = ?", Gid_, Realm_, GrantUpdate_).Limit(limit, offset).Find(_NodeAccess)
	return _NodeAccess, err
}

// GetNodeAccessesByGidAndRealmAndGrantDelete Get NodeAccesses via GidAndRealmAndGrantDelete
func GetNodeAccessesByGidAndRealmAndGrantDelete(offset int, limit int, Gid_ int, Realm_ string, GrantDelete_ int) (*[]*NodeAccess, error) {
	var _NodeAccess = new([]*NodeAccess)
	err := Engine.Table("node_access").Where("gid = ? and realm = ? and grant_delete = ?", Gid_, Realm_, GrantDelete_).Limit(limit, offset).Find(_NodeAccess)
	return _NodeAccess, err
}

// GetNodeAccessesByGidAndGrantViewAndGrantUpdate Get NodeAccesses via GidAndGrantViewAndGrantUpdate
func GetNodeAccessesByGidAndGrantViewAndGrantUpdate(offset int, limit int, Gid_ int, GrantView_ int, GrantUpdate_ int) (*[]*NodeAccess, error) {
	var _NodeAccess = new([]*NodeAccess)
	err := Engine.Table("node_access").Where("gid = ? and grant_view = ? and grant_update = ?", Gid_, GrantView_, GrantUpdate_).Limit(limit, offset).Find(_NodeAccess)
	return _NodeAccess, err
}

// GetNodeAccessesByGidAndGrantViewAndGrantDelete Get NodeAccesses via GidAndGrantViewAndGrantDelete
func GetNodeAccessesByGidAndGrantViewAndGrantDelete(offset int, limit int, Gid_ int, GrantView_ int, GrantDelete_ int) (*[]*NodeAccess, error) {
	var _NodeAccess = new([]*NodeAccess)
	err := Engine.Table("node_access").Where("gid = ? and grant_view = ? and grant_delete = ?", Gid_, GrantView_, GrantDelete_).Limit(limit, offset).Find(_NodeAccess)
	return _NodeAccess, err
}

// GetNodeAccessesByGidAndGrantUpdateAndGrantDelete Get NodeAccesses via GidAndGrantUpdateAndGrantDelete
func GetNodeAccessesByGidAndGrantUpdateAndGrantDelete(offset int, limit int, Gid_ int, GrantUpdate_ int, GrantDelete_ int) (*[]*NodeAccess, error) {
	var _NodeAccess = new([]*NodeAccess)
	err := Engine.Table("node_access").Where("gid = ? and grant_update = ? and grant_delete = ?", Gid_, GrantUpdate_, GrantDelete_).Limit(limit, offset).Find(_NodeAccess)
	return _NodeAccess, err
}

// GetNodeAccessesByRealmAndGrantViewAndGrantUpdate Get NodeAccesses via RealmAndGrantViewAndGrantUpdate
func GetNodeAccessesByRealmAndGrantViewAndGrantUpdate(offset int, limit int, Realm_ string, GrantView_ int, GrantUpdate_ int) (*[]*NodeAccess, error) {
	var _NodeAccess = new([]*NodeAccess)
	err := Engine.Table("node_access").Where("realm = ? and grant_view = ? and grant_update = ?", Realm_, GrantView_, GrantUpdate_).Limit(limit, offset).Find(_NodeAccess)
	return _NodeAccess, err
}

// GetNodeAccessesByRealmAndGrantViewAndGrantDelete Get NodeAccesses via RealmAndGrantViewAndGrantDelete
func GetNodeAccessesByRealmAndGrantViewAndGrantDelete(offset int, limit int, Realm_ string, GrantView_ int, GrantDelete_ int) (*[]*NodeAccess, error) {
	var _NodeAccess = new([]*NodeAccess)
	err := Engine.Table("node_access").Where("realm = ? and grant_view = ? and grant_delete = ?", Realm_, GrantView_, GrantDelete_).Limit(limit, offset).Find(_NodeAccess)
	return _NodeAccess, err
}

// GetNodeAccessesByRealmAndGrantUpdateAndGrantDelete Get NodeAccesses via RealmAndGrantUpdateAndGrantDelete
func GetNodeAccessesByRealmAndGrantUpdateAndGrantDelete(offset int, limit int, Realm_ string, GrantUpdate_ int, GrantDelete_ int) (*[]*NodeAccess, error) {
	var _NodeAccess = new([]*NodeAccess)
	err := Engine.Table("node_access").Where("realm = ? and grant_update = ? and grant_delete = ?", Realm_, GrantUpdate_, GrantDelete_).Limit(limit, offset).Find(_NodeAccess)
	return _NodeAccess, err
}

// GetNodeAccessesByGrantViewAndGrantUpdateAndGrantDelete Get NodeAccesses via GrantViewAndGrantUpdateAndGrantDelete
func GetNodeAccessesByGrantViewAndGrantUpdateAndGrantDelete(offset int, limit int, GrantView_ int, GrantUpdate_ int, GrantDelete_ int) (*[]*NodeAccess, error) {
	var _NodeAccess = new([]*NodeAccess)
	err := Engine.Table("node_access").Where("grant_view = ? and grant_update = ? and grant_delete = ?", GrantView_, GrantUpdate_, GrantDelete_).Limit(limit, offset).Find(_NodeAccess)
	return _NodeAccess, err
}

// GetNodeAccessesByNidAndLangcode Get NodeAccesses via NidAndLangcode
func GetNodeAccessesByNidAndLangcode(offset int, limit int, Nid_ int64, Langcode_ string) (*[]*NodeAccess, error) {
	var _NodeAccess = new([]*NodeAccess)
	err := Engine.Table("node_access").Where("nid = ? and langcode = ?", Nid_, Langcode_).Limit(limit, offset).Find(_NodeAccess)
	return _NodeAccess, err
}

// GetNodeAccessesByNidAndFallback Get NodeAccesses via NidAndFallback
func GetNodeAccessesByNidAndFallback(offset int, limit int, Nid_ int64, Fallback_ int) (*[]*NodeAccess, error) {
	var _NodeAccess = new([]*NodeAccess)
	err := Engine.Table("node_access").Where("nid = ? and fallback = ?", Nid_, Fallback_).Limit(limit, offset).Find(_NodeAccess)
	return _NodeAccess, err
}

// GetNodeAccessesByNidAndGid Get NodeAccesses via NidAndGid
func GetNodeAccessesByNidAndGid(offset int, limit int, Nid_ int64, Gid_ int) (*[]*NodeAccess, error) {
	var _NodeAccess = new([]*NodeAccess)
	err := Engine.Table("node_access").Where("nid = ? and gid = ?", Nid_, Gid_).Limit(limit, offset).Find(_NodeAccess)
	return _NodeAccess, err
}

// GetNodeAccessesByNidAndRealm Get NodeAccesses via NidAndRealm
func GetNodeAccessesByNidAndRealm(offset int, limit int, Nid_ int64, Realm_ string) (*[]*NodeAccess, error) {
	var _NodeAccess = new([]*NodeAccess)
	err := Engine.Table("node_access").Where("nid = ? and realm = ?", Nid_, Realm_).Limit(limit, offset).Find(_NodeAccess)
	return _NodeAccess, err
}

// GetNodeAccessesByNidAndGrantView Get NodeAccesses via NidAndGrantView
func GetNodeAccessesByNidAndGrantView(offset int, limit int, Nid_ int64, GrantView_ int) (*[]*NodeAccess, error) {
	var _NodeAccess = new([]*NodeAccess)
	err := Engine.Table("node_access").Where("nid = ? and grant_view = ?", Nid_, GrantView_).Limit(limit, offset).Find(_NodeAccess)
	return _NodeAccess, err
}

// GetNodeAccessesByNidAndGrantUpdate Get NodeAccesses via NidAndGrantUpdate
func GetNodeAccessesByNidAndGrantUpdate(offset int, limit int, Nid_ int64, GrantUpdate_ int) (*[]*NodeAccess, error) {
	var _NodeAccess = new([]*NodeAccess)
	err := Engine.Table("node_access").Where("nid = ? and grant_update = ?", Nid_, GrantUpdate_).Limit(limit, offset).Find(_NodeAccess)
	return _NodeAccess, err
}

// GetNodeAccessesByNidAndGrantDelete Get NodeAccesses via NidAndGrantDelete
func GetNodeAccessesByNidAndGrantDelete(offset int, limit int, Nid_ int64, GrantDelete_ int) (*[]*NodeAccess, error) {
	var _NodeAccess = new([]*NodeAccess)
	err := Engine.Table("node_access").Where("nid = ? and grant_delete = ?", Nid_, GrantDelete_).Limit(limit, offset).Find(_NodeAccess)
	return _NodeAccess, err
}

// GetNodeAccessesByLangcodeAndFallback Get NodeAccesses via LangcodeAndFallback
func GetNodeAccessesByLangcodeAndFallback(offset int, limit int, Langcode_ string, Fallback_ int) (*[]*NodeAccess, error) {
	var _NodeAccess = new([]*NodeAccess)
	err := Engine.Table("node_access").Where("langcode = ? and fallback = ?", Langcode_, Fallback_).Limit(limit, offset).Find(_NodeAccess)
	return _NodeAccess, err
}

// GetNodeAccessesByLangcodeAndGid Get NodeAccesses via LangcodeAndGid
func GetNodeAccessesByLangcodeAndGid(offset int, limit int, Langcode_ string, Gid_ int) (*[]*NodeAccess, error) {
	var _NodeAccess = new([]*NodeAccess)
	err := Engine.Table("node_access").Where("langcode = ? and gid = ?", Langcode_, Gid_).Limit(limit, offset).Find(_NodeAccess)
	return _NodeAccess, err
}

// GetNodeAccessesByLangcodeAndRealm Get NodeAccesses via LangcodeAndRealm
func GetNodeAccessesByLangcodeAndRealm(offset int, limit int, Langcode_ string, Realm_ string) (*[]*NodeAccess, error) {
	var _NodeAccess = new([]*NodeAccess)
	err := Engine.Table("node_access").Where("langcode = ? and realm = ?", Langcode_, Realm_).Limit(limit, offset).Find(_NodeAccess)
	return _NodeAccess, err
}

// GetNodeAccessesByLangcodeAndGrantView Get NodeAccesses via LangcodeAndGrantView
func GetNodeAccessesByLangcodeAndGrantView(offset int, limit int, Langcode_ string, GrantView_ int) (*[]*NodeAccess, error) {
	var _NodeAccess = new([]*NodeAccess)
	err := Engine.Table("node_access").Where("langcode = ? and grant_view = ?", Langcode_, GrantView_).Limit(limit, offset).Find(_NodeAccess)
	return _NodeAccess, err
}

// GetNodeAccessesByLangcodeAndGrantUpdate Get NodeAccesses via LangcodeAndGrantUpdate
func GetNodeAccessesByLangcodeAndGrantUpdate(offset int, limit int, Langcode_ string, GrantUpdate_ int) (*[]*NodeAccess, error) {
	var _NodeAccess = new([]*NodeAccess)
	err := Engine.Table("node_access").Where("langcode = ? and grant_update = ?", Langcode_, GrantUpdate_).Limit(limit, offset).Find(_NodeAccess)
	return _NodeAccess, err
}

// GetNodeAccessesByLangcodeAndGrantDelete Get NodeAccesses via LangcodeAndGrantDelete
func GetNodeAccessesByLangcodeAndGrantDelete(offset int, limit int, Langcode_ string, GrantDelete_ int) (*[]*NodeAccess, error) {
	var _NodeAccess = new([]*NodeAccess)
	err := Engine.Table("node_access").Where("langcode = ? and grant_delete = ?", Langcode_, GrantDelete_).Limit(limit, offset).Find(_NodeAccess)
	return _NodeAccess, err
}

// GetNodeAccessesByFallbackAndGid Get NodeAccesses via FallbackAndGid
func GetNodeAccessesByFallbackAndGid(offset int, limit int, Fallback_ int, Gid_ int) (*[]*NodeAccess, error) {
	var _NodeAccess = new([]*NodeAccess)
	err := Engine.Table("node_access").Where("fallback = ? and gid = ?", Fallback_, Gid_).Limit(limit, offset).Find(_NodeAccess)
	return _NodeAccess, err
}

// GetNodeAccessesByFallbackAndRealm Get NodeAccesses via FallbackAndRealm
func GetNodeAccessesByFallbackAndRealm(offset int, limit int, Fallback_ int, Realm_ string) (*[]*NodeAccess, error) {
	var _NodeAccess = new([]*NodeAccess)
	err := Engine.Table("node_access").Where("fallback = ? and realm = ?", Fallback_, Realm_).Limit(limit, offset).Find(_NodeAccess)
	return _NodeAccess, err
}

// GetNodeAccessesByFallbackAndGrantView Get NodeAccesses via FallbackAndGrantView
func GetNodeAccessesByFallbackAndGrantView(offset int, limit int, Fallback_ int, GrantView_ int) (*[]*NodeAccess, error) {
	var _NodeAccess = new([]*NodeAccess)
	err := Engine.Table("node_access").Where("fallback = ? and grant_view = ?", Fallback_, GrantView_).Limit(limit, offset).Find(_NodeAccess)
	return _NodeAccess, err
}

// GetNodeAccessesByFallbackAndGrantUpdate Get NodeAccesses via FallbackAndGrantUpdate
func GetNodeAccessesByFallbackAndGrantUpdate(offset int, limit int, Fallback_ int, GrantUpdate_ int) (*[]*NodeAccess, error) {
	var _NodeAccess = new([]*NodeAccess)
	err := Engine.Table("node_access").Where("fallback = ? and grant_update = ?", Fallback_, GrantUpdate_).Limit(limit, offset).Find(_NodeAccess)
	return _NodeAccess, err
}

// GetNodeAccessesByFallbackAndGrantDelete Get NodeAccesses via FallbackAndGrantDelete
func GetNodeAccessesByFallbackAndGrantDelete(offset int, limit int, Fallback_ int, GrantDelete_ int) (*[]*NodeAccess, error) {
	var _NodeAccess = new([]*NodeAccess)
	err := Engine.Table("node_access").Where("fallback = ? and grant_delete = ?", Fallback_, GrantDelete_).Limit(limit, offset).Find(_NodeAccess)
	return _NodeAccess, err
}

// GetNodeAccessesByGidAndRealm Get NodeAccesses via GidAndRealm
func GetNodeAccessesByGidAndRealm(offset int, limit int, Gid_ int, Realm_ string) (*[]*NodeAccess, error) {
	var _NodeAccess = new([]*NodeAccess)
	err := Engine.Table("node_access").Where("gid = ? and realm = ?", Gid_, Realm_).Limit(limit, offset).Find(_NodeAccess)
	return _NodeAccess, err
}

// GetNodeAccessesByGidAndGrantView Get NodeAccesses via GidAndGrantView
func GetNodeAccessesByGidAndGrantView(offset int, limit int, Gid_ int, GrantView_ int) (*[]*NodeAccess, error) {
	var _NodeAccess = new([]*NodeAccess)
	err := Engine.Table("node_access").Where("gid = ? and grant_view = ?", Gid_, GrantView_).Limit(limit, offset).Find(_NodeAccess)
	return _NodeAccess, err
}

// GetNodeAccessesByGidAndGrantUpdate Get NodeAccesses via GidAndGrantUpdate
func GetNodeAccessesByGidAndGrantUpdate(offset int, limit int, Gid_ int, GrantUpdate_ int) (*[]*NodeAccess, error) {
	var _NodeAccess = new([]*NodeAccess)
	err := Engine.Table("node_access").Where("gid = ? and grant_update = ?", Gid_, GrantUpdate_).Limit(limit, offset).Find(_NodeAccess)
	return _NodeAccess, err
}

// GetNodeAccessesByGidAndGrantDelete Get NodeAccesses via GidAndGrantDelete
func GetNodeAccessesByGidAndGrantDelete(offset int, limit int, Gid_ int, GrantDelete_ int) (*[]*NodeAccess, error) {
	var _NodeAccess = new([]*NodeAccess)
	err := Engine.Table("node_access").Where("gid = ? and grant_delete = ?", Gid_, GrantDelete_).Limit(limit, offset).Find(_NodeAccess)
	return _NodeAccess, err
}

// GetNodeAccessesByRealmAndGrantView Get NodeAccesses via RealmAndGrantView
func GetNodeAccessesByRealmAndGrantView(offset int, limit int, Realm_ string, GrantView_ int) (*[]*NodeAccess, error) {
	var _NodeAccess = new([]*NodeAccess)
	err := Engine.Table("node_access").Where("realm = ? and grant_view = ?", Realm_, GrantView_).Limit(limit, offset).Find(_NodeAccess)
	return _NodeAccess, err
}

// GetNodeAccessesByRealmAndGrantUpdate Get NodeAccesses via RealmAndGrantUpdate
func GetNodeAccessesByRealmAndGrantUpdate(offset int, limit int, Realm_ string, GrantUpdate_ int) (*[]*NodeAccess, error) {
	var _NodeAccess = new([]*NodeAccess)
	err := Engine.Table("node_access").Where("realm = ? and grant_update = ?", Realm_, GrantUpdate_).Limit(limit, offset).Find(_NodeAccess)
	return _NodeAccess, err
}

// GetNodeAccessesByRealmAndGrantDelete Get NodeAccesses via RealmAndGrantDelete
func GetNodeAccessesByRealmAndGrantDelete(offset int, limit int, Realm_ string, GrantDelete_ int) (*[]*NodeAccess, error) {
	var _NodeAccess = new([]*NodeAccess)
	err := Engine.Table("node_access").Where("realm = ? and grant_delete = ?", Realm_, GrantDelete_).Limit(limit, offset).Find(_NodeAccess)
	return _NodeAccess, err
}

// GetNodeAccessesByGrantViewAndGrantUpdate Get NodeAccesses via GrantViewAndGrantUpdate
func GetNodeAccessesByGrantViewAndGrantUpdate(offset int, limit int, GrantView_ int, GrantUpdate_ int) (*[]*NodeAccess, error) {
	var _NodeAccess = new([]*NodeAccess)
	err := Engine.Table("node_access").Where("grant_view = ? and grant_update = ?", GrantView_, GrantUpdate_).Limit(limit, offset).Find(_NodeAccess)
	return _NodeAccess, err
}

// GetNodeAccessesByGrantViewAndGrantDelete Get NodeAccesses via GrantViewAndGrantDelete
func GetNodeAccessesByGrantViewAndGrantDelete(offset int, limit int, GrantView_ int, GrantDelete_ int) (*[]*NodeAccess, error) {
	var _NodeAccess = new([]*NodeAccess)
	err := Engine.Table("node_access").Where("grant_view = ? and grant_delete = ?", GrantView_, GrantDelete_).Limit(limit, offset).Find(_NodeAccess)
	return _NodeAccess, err
}

// GetNodeAccessesByGrantUpdateAndGrantDelete Get NodeAccesses via GrantUpdateAndGrantDelete
func GetNodeAccessesByGrantUpdateAndGrantDelete(offset int, limit int, GrantUpdate_ int, GrantDelete_ int) (*[]*NodeAccess, error) {
	var _NodeAccess = new([]*NodeAccess)
	err := Engine.Table("node_access").Where("grant_update = ? and grant_delete = ?", GrantUpdate_, GrantDelete_).Limit(limit, offset).Find(_NodeAccess)
	return _NodeAccess, err
}

// GetNodeAccesses Get NodeAccesses via field
func GetNodeAccesses(offset int, limit int, field string) (*[]*NodeAccess, error) {
	var _NodeAccess = new([]*NodeAccess)
	err := Engine.Table("node_access").Limit(limit, offset).Desc(field).Find(_NodeAccess)
	return _NodeAccess, err
}

// GetNodeAccessesViaNid Get NodeAccesss via Nid
func GetNodeAccessesViaNid(offset int, limit int, Nid_ int64, field string) (*[]*NodeAccess, error) {
	var _NodeAccess = new([]*NodeAccess)
	err := Engine.Table("node_access").Where("nid = ?", Nid_).Limit(limit, offset).Desc(field).Find(_NodeAccess)
	return _NodeAccess, err
}

// GetNodeAccessesViaLangcode Get NodeAccesss via Langcode
func GetNodeAccessesViaLangcode(offset int, limit int, Langcode_ string, field string) (*[]*NodeAccess, error) {
	var _NodeAccess = new([]*NodeAccess)
	err := Engine.Table("node_access").Where("langcode = ?", Langcode_).Limit(limit, offset).Desc(field).Find(_NodeAccess)
	return _NodeAccess, err
}

// GetNodeAccessesViaFallback Get NodeAccesss via Fallback
func GetNodeAccessesViaFallback(offset int, limit int, Fallback_ int, field string) (*[]*NodeAccess, error) {
	var _NodeAccess = new([]*NodeAccess)
	err := Engine.Table("node_access").Where("fallback = ?", Fallback_).Limit(limit, offset).Desc(field).Find(_NodeAccess)
	return _NodeAccess, err
}

// GetNodeAccessesViaGid Get NodeAccesss via Gid
func GetNodeAccessesViaGid(offset int, limit int, Gid_ int, field string) (*[]*NodeAccess, error) {
	var _NodeAccess = new([]*NodeAccess)
	err := Engine.Table("node_access").Where("gid = ?", Gid_).Limit(limit, offset).Desc(field).Find(_NodeAccess)
	return _NodeAccess, err
}

// GetNodeAccessesViaRealm Get NodeAccesss via Realm
func GetNodeAccessesViaRealm(offset int, limit int, Realm_ string, field string) (*[]*NodeAccess, error) {
	var _NodeAccess = new([]*NodeAccess)
	err := Engine.Table("node_access").Where("realm = ?", Realm_).Limit(limit, offset).Desc(field).Find(_NodeAccess)
	return _NodeAccess, err
}

// GetNodeAccessesViaGrantView Get NodeAccesss via GrantView
func GetNodeAccessesViaGrantView(offset int, limit int, GrantView_ int, field string) (*[]*NodeAccess, error) {
	var _NodeAccess = new([]*NodeAccess)
	err := Engine.Table("node_access").Where("grant_view = ?", GrantView_).Limit(limit, offset).Desc(field).Find(_NodeAccess)
	return _NodeAccess, err
}

// GetNodeAccessesViaGrantUpdate Get NodeAccesss via GrantUpdate
func GetNodeAccessesViaGrantUpdate(offset int, limit int, GrantUpdate_ int, field string) (*[]*NodeAccess, error) {
	var _NodeAccess = new([]*NodeAccess)
	err := Engine.Table("node_access").Where("grant_update = ?", GrantUpdate_).Limit(limit, offset).Desc(field).Find(_NodeAccess)
	return _NodeAccess, err
}

// GetNodeAccessesViaGrantDelete Get NodeAccesss via GrantDelete
func GetNodeAccessesViaGrantDelete(offset int, limit int, GrantDelete_ int, field string) (*[]*NodeAccess, error) {
	var _NodeAccess = new([]*NodeAccess)
	err := Engine.Table("node_access").Where("grant_delete = ?", GrantDelete_).Limit(limit, offset).Desc(field).Find(_NodeAccess)
	return _NodeAccess, err
}

// HasNodeAccessViaNid Has NodeAccess via Nid
func HasNodeAccessViaNid(iNid int64) bool {
	if has, err := Engine.Where("nid = ?", iNid).Get(new(NodeAccess)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasNodeAccessViaLangcode Has NodeAccess via Langcode
func HasNodeAccessViaLangcode(iLangcode string) bool {
	if has, err := Engine.Where("langcode = ?", iLangcode).Get(new(NodeAccess)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasNodeAccessViaFallback Has NodeAccess via Fallback
func HasNodeAccessViaFallback(iFallback int) bool {
	if has, err := Engine.Where("fallback = ?", iFallback).Get(new(NodeAccess)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasNodeAccessViaGid Has NodeAccess via Gid
func HasNodeAccessViaGid(iGid int) bool {
	if has, err := Engine.Where("gid = ?", iGid).Get(new(NodeAccess)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasNodeAccessViaRealm Has NodeAccess via Realm
func HasNodeAccessViaRealm(iRealm string) bool {
	if has, err := Engine.Where("realm = ?", iRealm).Get(new(NodeAccess)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasNodeAccessViaGrantView Has NodeAccess via GrantView
func HasNodeAccessViaGrantView(iGrantView int) bool {
	if has, err := Engine.Where("grant_view = ?", iGrantView).Get(new(NodeAccess)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasNodeAccessViaGrantUpdate Has NodeAccess via GrantUpdate
func HasNodeAccessViaGrantUpdate(iGrantUpdate int) bool {
	if has, err := Engine.Where("grant_update = ?", iGrantUpdate).Get(new(NodeAccess)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasNodeAccessViaGrantDelete Has NodeAccess via GrantDelete
func HasNodeAccessViaGrantDelete(iGrantDelete int) bool {
	if has, err := Engine.Where("grant_delete = ?", iGrantDelete).Get(new(NodeAccess)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// GetNodeAccessViaNid Get NodeAccess via Nid
func GetNodeAccessViaNid(iNid int64) (*NodeAccess, error) {
	var _NodeAccess = &NodeAccess{Nid: iNid}
	has, err := Engine.Get(_NodeAccess)
	if has {
		return _NodeAccess, err
	} else {
		return nil, err
	}
}

// GetNodeAccessViaLangcode Get NodeAccess via Langcode
func GetNodeAccessViaLangcode(iLangcode string) (*NodeAccess, error) {
	var _NodeAccess = &NodeAccess{Langcode: iLangcode}
	has, err := Engine.Get(_NodeAccess)
	if has {
		return _NodeAccess, err
	} else {
		return nil, err
	}
}

// GetNodeAccessViaFallback Get NodeAccess via Fallback
func GetNodeAccessViaFallback(iFallback int) (*NodeAccess, error) {
	var _NodeAccess = &NodeAccess{Fallback: iFallback}
	has, err := Engine.Get(_NodeAccess)
	if has {
		return _NodeAccess, err
	} else {
		return nil, err
	}
}

// GetNodeAccessViaGid Get NodeAccess via Gid
func GetNodeAccessViaGid(iGid int) (*NodeAccess, error) {
	var _NodeAccess = &NodeAccess{Gid: iGid}
	has, err := Engine.Get(_NodeAccess)
	if has {
		return _NodeAccess, err
	} else {
		return nil, err
	}
}

// GetNodeAccessViaRealm Get NodeAccess via Realm
func GetNodeAccessViaRealm(iRealm string) (*NodeAccess, error) {
	var _NodeAccess = &NodeAccess{Realm: iRealm}
	has, err := Engine.Get(_NodeAccess)
	if has {
		return _NodeAccess, err
	} else {
		return nil, err
	}
}

// GetNodeAccessViaGrantView Get NodeAccess via GrantView
func GetNodeAccessViaGrantView(iGrantView int) (*NodeAccess, error) {
	var _NodeAccess = &NodeAccess{GrantView: iGrantView}
	has, err := Engine.Get(_NodeAccess)
	if has {
		return _NodeAccess, err
	} else {
		return nil, err
	}
}

// GetNodeAccessViaGrantUpdate Get NodeAccess via GrantUpdate
func GetNodeAccessViaGrantUpdate(iGrantUpdate int) (*NodeAccess, error) {
	var _NodeAccess = &NodeAccess{GrantUpdate: iGrantUpdate}
	has, err := Engine.Get(_NodeAccess)
	if has {
		return _NodeAccess, err
	} else {
		return nil, err
	}
}

// GetNodeAccessViaGrantDelete Get NodeAccess via GrantDelete
func GetNodeAccessViaGrantDelete(iGrantDelete int) (*NodeAccess, error) {
	var _NodeAccess = &NodeAccess{GrantDelete: iGrantDelete}
	has, err := Engine.Get(_NodeAccess)
	if has {
		return _NodeAccess, err
	} else {
		return nil, err
	}
}

// SetNodeAccessViaNid Set NodeAccess via Nid
func SetNodeAccessViaNid(iNid int64, node_access *NodeAccess) (int64, error) {
	node_access.Nid = iNid
	return Engine.Insert(node_access)
}

// SetNodeAccessViaLangcode Set NodeAccess via Langcode
func SetNodeAccessViaLangcode(iLangcode string, node_access *NodeAccess) (int64, error) {
	node_access.Langcode = iLangcode
	return Engine.Insert(node_access)
}

// SetNodeAccessViaFallback Set NodeAccess via Fallback
func SetNodeAccessViaFallback(iFallback int, node_access *NodeAccess) (int64, error) {
	node_access.Fallback = iFallback
	return Engine.Insert(node_access)
}

// SetNodeAccessViaGid Set NodeAccess via Gid
func SetNodeAccessViaGid(iGid int, node_access *NodeAccess) (int64, error) {
	node_access.Gid = iGid
	return Engine.Insert(node_access)
}

// SetNodeAccessViaRealm Set NodeAccess via Realm
func SetNodeAccessViaRealm(iRealm string, node_access *NodeAccess) (int64, error) {
	node_access.Realm = iRealm
	return Engine.Insert(node_access)
}

// SetNodeAccessViaGrantView Set NodeAccess via GrantView
func SetNodeAccessViaGrantView(iGrantView int, node_access *NodeAccess) (int64, error) {
	node_access.GrantView = iGrantView
	return Engine.Insert(node_access)
}

// SetNodeAccessViaGrantUpdate Set NodeAccess via GrantUpdate
func SetNodeAccessViaGrantUpdate(iGrantUpdate int, node_access *NodeAccess) (int64, error) {
	node_access.GrantUpdate = iGrantUpdate
	return Engine.Insert(node_access)
}

// SetNodeAccessViaGrantDelete Set NodeAccess via GrantDelete
func SetNodeAccessViaGrantDelete(iGrantDelete int, node_access *NodeAccess) (int64, error) {
	node_access.GrantDelete = iGrantDelete
	return Engine.Insert(node_access)
}

// AddNodeAccess Add NodeAccess via all columns
func AddNodeAccess(iNid int64, iLangcode string, iFallback int, iGid int, iRealm string, iGrantView int, iGrantUpdate int, iGrantDelete int) error {
	_NodeAccess := &NodeAccess{Nid: iNid, Langcode: iLangcode, Fallback: iFallback, Gid: iGid, Realm: iRealm, GrantView: iGrantView, GrantUpdate: iGrantUpdate, GrantDelete: iGrantDelete}
	if _, err := Engine.Insert(_NodeAccess); err != nil {
		return err
	} else {
		return nil
	}
}

// PostNodeAccess Post NodeAccess via iNodeAccess
func PostNodeAccess(iNodeAccess *NodeAccess) (int64, error) {
	_, err := Engine.Insert(iNodeAccess)
	return iNodeAccess.Nid, err
}

// PutNodeAccess Put NodeAccess
func PutNodeAccess(iNodeAccess *NodeAccess) (int64, error) {
	_, err := Engine.Id(iNodeAccess.Nid).Update(iNodeAccess)
	return iNodeAccess.Nid, err
}

// PutNodeAccessViaNid Put NodeAccess via Nid
func PutNodeAccessViaNid(Nid_ int64, iNodeAccess *NodeAccess) (int64, error) {
	row, err := Engine.Update(iNodeAccess, &NodeAccess{Nid: Nid_})
	return row, err
}

// PutNodeAccessViaLangcode Put NodeAccess via Langcode
func PutNodeAccessViaLangcode(Langcode_ string, iNodeAccess *NodeAccess) (int64, error) {
	row, err := Engine.Update(iNodeAccess, &NodeAccess{Langcode: Langcode_})
	return row, err
}

// PutNodeAccessViaFallback Put NodeAccess via Fallback
func PutNodeAccessViaFallback(Fallback_ int, iNodeAccess *NodeAccess) (int64, error) {
	row, err := Engine.Update(iNodeAccess, &NodeAccess{Fallback: Fallback_})
	return row, err
}

// PutNodeAccessViaGid Put NodeAccess via Gid
func PutNodeAccessViaGid(Gid_ int, iNodeAccess *NodeAccess) (int64, error) {
	row, err := Engine.Update(iNodeAccess, &NodeAccess{Gid: Gid_})
	return row, err
}

// PutNodeAccessViaRealm Put NodeAccess via Realm
func PutNodeAccessViaRealm(Realm_ string, iNodeAccess *NodeAccess) (int64, error) {
	row, err := Engine.Update(iNodeAccess, &NodeAccess{Realm: Realm_})
	return row, err
}

// PutNodeAccessViaGrantView Put NodeAccess via GrantView
func PutNodeAccessViaGrantView(GrantView_ int, iNodeAccess *NodeAccess) (int64, error) {
	row, err := Engine.Update(iNodeAccess, &NodeAccess{GrantView: GrantView_})
	return row, err
}

// PutNodeAccessViaGrantUpdate Put NodeAccess via GrantUpdate
func PutNodeAccessViaGrantUpdate(GrantUpdate_ int, iNodeAccess *NodeAccess) (int64, error) {
	row, err := Engine.Update(iNodeAccess, &NodeAccess{GrantUpdate: GrantUpdate_})
	return row, err
}

// PutNodeAccessViaGrantDelete Put NodeAccess via GrantDelete
func PutNodeAccessViaGrantDelete(GrantDelete_ int, iNodeAccess *NodeAccess) (int64, error) {
	row, err := Engine.Update(iNodeAccess, &NodeAccess{GrantDelete: GrantDelete_})
	return row, err
}

// UpdateNodeAccessViaNid via map[string]interface{}{}
func UpdateNodeAccessViaNid(iNid int64, iNodeAccessMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(NodeAccess)).Where("nid = ?", iNid).Update(iNodeAccessMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateNodeAccessViaLangcode via map[string]interface{}{}
func UpdateNodeAccessViaLangcode(iLangcode string, iNodeAccessMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(NodeAccess)).Where("langcode = ?", iLangcode).Update(iNodeAccessMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateNodeAccessViaFallback via map[string]interface{}{}
func UpdateNodeAccessViaFallback(iFallback int, iNodeAccessMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(NodeAccess)).Where("fallback = ?", iFallback).Update(iNodeAccessMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateNodeAccessViaGid via map[string]interface{}{}
func UpdateNodeAccessViaGid(iGid int, iNodeAccessMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(NodeAccess)).Where("gid = ?", iGid).Update(iNodeAccessMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateNodeAccessViaRealm via map[string]interface{}{}
func UpdateNodeAccessViaRealm(iRealm string, iNodeAccessMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(NodeAccess)).Where("realm = ?", iRealm).Update(iNodeAccessMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateNodeAccessViaGrantView via map[string]interface{}{}
func UpdateNodeAccessViaGrantView(iGrantView int, iNodeAccessMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(NodeAccess)).Where("grant_view = ?", iGrantView).Update(iNodeAccessMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateNodeAccessViaGrantUpdate via map[string]interface{}{}
func UpdateNodeAccessViaGrantUpdate(iGrantUpdate int, iNodeAccessMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(NodeAccess)).Where("grant_update = ?", iGrantUpdate).Update(iNodeAccessMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateNodeAccessViaGrantDelete via map[string]interface{}{}
func UpdateNodeAccessViaGrantDelete(iGrantDelete int, iNodeAccessMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(NodeAccess)).Where("grant_delete = ?", iGrantDelete).Update(iNodeAccessMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// DeleteNodeAccess Delete NodeAccess
func DeleteNodeAccess(iNid int64) (int64, error) {
	row, err := Engine.Id(iNid).Delete(new(NodeAccess))
	return row, err
}

// DeleteNodeAccessViaNid Delete NodeAccess via Nid
func DeleteNodeAccessViaNid(iNid int64) (err error) {
	var has bool
	var _NodeAccess = &NodeAccess{Nid: iNid}
	if has, err = Engine.Get(_NodeAccess); (has == true) && (err == nil) {
		if row, err := Engine.Where("nid = ?", iNid).Delete(new(NodeAccess)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteNodeAccessViaLangcode Delete NodeAccess via Langcode
func DeleteNodeAccessViaLangcode(iLangcode string) (err error) {
	var has bool
	var _NodeAccess = &NodeAccess{Langcode: iLangcode}
	if has, err = Engine.Get(_NodeAccess); (has == true) && (err == nil) {
		if row, err := Engine.Where("langcode = ?", iLangcode).Delete(new(NodeAccess)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteNodeAccessViaFallback Delete NodeAccess via Fallback
func DeleteNodeAccessViaFallback(iFallback int) (err error) {
	var has bool
	var _NodeAccess = &NodeAccess{Fallback: iFallback}
	if has, err = Engine.Get(_NodeAccess); (has == true) && (err == nil) {
		if row, err := Engine.Where("fallback = ?", iFallback).Delete(new(NodeAccess)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteNodeAccessViaGid Delete NodeAccess via Gid
func DeleteNodeAccessViaGid(iGid int) (err error) {
	var has bool
	var _NodeAccess = &NodeAccess{Gid: iGid}
	if has, err = Engine.Get(_NodeAccess); (has == true) && (err == nil) {
		if row, err := Engine.Where("gid = ?", iGid).Delete(new(NodeAccess)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteNodeAccessViaRealm Delete NodeAccess via Realm
func DeleteNodeAccessViaRealm(iRealm string) (err error) {
	var has bool
	var _NodeAccess = &NodeAccess{Realm: iRealm}
	if has, err = Engine.Get(_NodeAccess); (has == true) && (err == nil) {
		if row, err := Engine.Where("realm = ?", iRealm).Delete(new(NodeAccess)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteNodeAccessViaGrantView Delete NodeAccess via GrantView
func DeleteNodeAccessViaGrantView(iGrantView int) (err error) {
	var has bool
	var _NodeAccess = &NodeAccess{GrantView: iGrantView}
	if has, err = Engine.Get(_NodeAccess); (has == true) && (err == nil) {
		if row, err := Engine.Where("grant_view = ?", iGrantView).Delete(new(NodeAccess)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteNodeAccessViaGrantUpdate Delete NodeAccess via GrantUpdate
func DeleteNodeAccessViaGrantUpdate(iGrantUpdate int) (err error) {
	var has bool
	var _NodeAccess = &NodeAccess{GrantUpdate: iGrantUpdate}
	if has, err = Engine.Get(_NodeAccess); (has == true) && (err == nil) {
		if row, err := Engine.Where("grant_update = ?", iGrantUpdate).Delete(new(NodeAccess)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteNodeAccessViaGrantDelete Delete NodeAccess via GrantDelete
func DeleteNodeAccessViaGrantDelete(iGrantDelete int) (err error) {
	var has bool
	var _NodeAccess = &NodeAccess{GrantDelete: iGrantDelete}
	if has, err = Engine.Get(_NodeAccess); (has == true) && (err == nil) {
		if row, err := Engine.Where("grant_delete = ?", iGrantDelete).Delete(new(NodeAccess)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}
