package model

type NodeFieldData struct {
	Nid                         int64  `xorm:"not null pk index(node__id__default_langcode__langcode) index(node__status_type) INT(10)"`
	Vid                         int    `xorm:"not null index INT(10)"`
	Type                        string `xorm:"not null index index(node__status_type) index(node__title_type) VARCHAR(32)"`
	Langcode                    string `xorm:"not null pk index(node__id__default_langcode__langcode) VARCHAR(12)"`
	Title                       string `xorm:"not null index(node__title_type) VARCHAR(255)"`
	Uid                         int    `xorm:"not null index INT(10)"`
	Status                      int    `xorm:"not null index(node__frontpage) index(node__status_type) TINYINT(4)"`
	Created                     int    `xorm:"not null index index(node__frontpage) INT(11)"`
	Changed                     int    `xorm:"not null index INT(11)"`
	Promote                     int    `xorm:"not null index(node__frontpage) TINYINT(4)"`
	Sticky                      int    `xorm:"not null index(node__frontpage) TINYINT(4)"`
	RevisionTranslationAffected int    `xorm:"TINYINT(4)"`
	DefaultLangcode             int    `xorm:"not null index(node__id__default_langcode__langcode) TINYINT(4)"`
}

// GetNodeFieldDatasCount NodeFieldDatas' Count
func GetNodeFieldDatasCount(offset int, limit int) (int64, error) {
	total, err := Engine.Limit(limit, offset).Count(&NodeFieldData{})
	return total, err
}

// GetNodeFieldDataCountViaNid Get NodeFieldData via Nid
func GetNodeFieldDataCountViaNid(iNid int64) int64 {
	n, _ := Engine.Where("nid = ?", iNid).Count(&NodeFieldData{Nid: iNid})
	return n
}

// GetNodeFieldDataCountViaVid Get NodeFieldData via Vid
func GetNodeFieldDataCountViaVid(iVid int) int64 {
	n, _ := Engine.Where("vid = ?", iVid).Count(&NodeFieldData{Vid: iVid})
	return n
}

// GetNodeFieldDataCountViaType Get NodeFieldData via Type
func GetNodeFieldDataCountViaType(iType string) int64 {
	n, _ := Engine.Where("type = ?", iType).Count(&NodeFieldData{Type: iType})
	return n
}

// GetNodeFieldDataCountViaLangcode Get NodeFieldData via Langcode
func GetNodeFieldDataCountViaLangcode(iLangcode string) int64 {
	n, _ := Engine.Where("langcode = ?", iLangcode).Count(&NodeFieldData{Langcode: iLangcode})
	return n
}

// GetNodeFieldDataCountViaTitle Get NodeFieldData via Title
func GetNodeFieldDataCountViaTitle(iTitle string) int64 {
	n, _ := Engine.Where("title = ?", iTitle).Count(&NodeFieldData{Title: iTitle})
	return n
}

// GetNodeFieldDataCountViaUid Get NodeFieldData via Uid
func GetNodeFieldDataCountViaUid(iUid int) int64 {
	n, _ := Engine.Where("uid = ?", iUid).Count(&NodeFieldData{Uid: iUid})
	return n
}

// GetNodeFieldDataCountViaStatus Get NodeFieldData via Status
func GetNodeFieldDataCountViaStatus(iStatus int) int64 {
	n, _ := Engine.Where("status = ?", iStatus).Count(&NodeFieldData{Status: iStatus})
	return n
}

// GetNodeFieldDataCountViaCreated Get NodeFieldData via Created
func GetNodeFieldDataCountViaCreated(iCreated int) int64 {
	n, _ := Engine.Where("created = ?", iCreated).Count(&NodeFieldData{Created: iCreated})
	return n
}

// GetNodeFieldDataCountViaChanged Get NodeFieldData via Changed
func GetNodeFieldDataCountViaChanged(iChanged int) int64 {
	n, _ := Engine.Where("changed = ?", iChanged).Count(&NodeFieldData{Changed: iChanged})
	return n
}

// GetNodeFieldDataCountViaPromote Get NodeFieldData via Promote
func GetNodeFieldDataCountViaPromote(iPromote int) int64 {
	n, _ := Engine.Where("promote = ?", iPromote).Count(&NodeFieldData{Promote: iPromote})
	return n
}

// GetNodeFieldDataCountViaSticky Get NodeFieldData via Sticky
func GetNodeFieldDataCountViaSticky(iSticky int) int64 {
	n, _ := Engine.Where("sticky = ?", iSticky).Count(&NodeFieldData{Sticky: iSticky})
	return n
}

// GetNodeFieldDataCountViaRevisionTranslationAffected Get NodeFieldData via RevisionTranslationAffected
func GetNodeFieldDataCountViaRevisionTranslationAffected(iRevisionTranslationAffected int) int64 {
	n, _ := Engine.Where("revision_translation_affected = ?", iRevisionTranslationAffected).Count(&NodeFieldData{RevisionTranslationAffected: iRevisionTranslationAffected})
	return n
}

// GetNodeFieldDataCountViaDefaultLangcode Get NodeFieldData via DefaultLangcode
func GetNodeFieldDataCountViaDefaultLangcode(iDefaultLangcode int) int64 {
	n, _ := Engine.Where("default_langcode = ?", iDefaultLangcode).Count(&NodeFieldData{DefaultLangcode: iDefaultLangcode})
	return n
}

// GetNodeFieldDatasByNidAndVidAndType Get NodeFieldDatas via NidAndVidAndType
func GetNodeFieldDatasByNidAndVidAndType(offset int, limit int, Nid_ int64, Vid_ int, Type_ string) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("nid = ? and vid = ? and type = ?", Nid_, Vid_, Type_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByNidAndVidAndLangcode Get NodeFieldDatas via NidAndVidAndLangcode
func GetNodeFieldDatasByNidAndVidAndLangcode(offset int, limit int, Nid_ int64, Vid_ int, Langcode_ string) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("nid = ? and vid = ? and langcode = ?", Nid_, Vid_, Langcode_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByNidAndVidAndTitle Get NodeFieldDatas via NidAndVidAndTitle
func GetNodeFieldDatasByNidAndVidAndTitle(offset int, limit int, Nid_ int64, Vid_ int, Title_ string) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("nid = ? and vid = ? and title = ?", Nid_, Vid_, Title_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByNidAndVidAndUid Get NodeFieldDatas via NidAndVidAndUid
func GetNodeFieldDatasByNidAndVidAndUid(offset int, limit int, Nid_ int64, Vid_ int, Uid_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("nid = ? and vid = ? and uid = ?", Nid_, Vid_, Uid_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByNidAndVidAndStatus Get NodeFieldDatas via NidAndVidAndStatus
func GetNodeFieldDatasByNidAndVidAndStatus(offset int, limit int, Nid_ int64, Vid_ int, Status_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("nid = ? and vid = ? and status = ?", Nid_, Vid_, Status_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByNidAndVidAndCreated Get NodeFieldDatas via NidAndVidAndCreated
func GetNodeFieldDatasByNidAndVidAndCreated(offset int, limit int, Nid_ int64, Vid_ int, Created_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("nid = ? and vid = ? and created = ?", Nid_, Vid_, Created_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByNidAndVidAndChanged Get NodeFieldDatas via NidAndVidAndChanged
func GetNodeFieldDatasByNidAndVidAndChanged(offset int, limit int, Nid_ int64, Vid_ int, Changed_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("nid = ? and vid = ? and changed = ?", Nid_, Vid_, Changed_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByNidAndVidAndPromote Get NodeFieldDatas via NidAndVidAndPromote
func GetNodeFieldDatasByNidAndVidAndPromote(offset int, limit int, Nid_ int64, Vid_ int, Promote_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("nid = ? and vid = ? and promote = ?", Nid_, Vid_, Promote_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByNidAndVidAndSticky Get NodeFieldDatas via NidAndVidAndSticky
func GetNodeFieldDatasByNidAndVidAndSticky(offset int, limit int, Nid_ int64, Vid_ int, Sticky_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("nid = ? and vid = ? and sticky = ?", Nid_, Vid_, Sticky_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByNidAndVidAndRevisionTranslationAffected Get NodeFieldDatas via NidAndVidAndRevisionTranslationAffected
func GetNodeFieldDatasByNidAndVidAndRevisionTranslationAffected(offset int, limit int, Nid_ int64, Vid_ int, RevisionTranslationAffected_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("nid = ? and vid = ? and revision_translation_affected = ?", Nid_, Vid_, RevisionTranslationAffected_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByNidAndVidAndDefaultLangcode Get NodeFieldDatas via NidAndVidAndDefaultLangcode
func GetNodeFieldDatasByNidAndVidAndDefaultLangcode(offset int, limit int, Nid_ int64, Vid_ int, DefaultLangcode_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("nid = ? and vid = ? and default_langcode = ?", Nid_, Vid_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByNidAndTypeAndLangcode Get NodeFieldDatas via NidAndTypeAndLangcode
func GetNodeFieldDatasByNidAndTypeAndLangcode(offset int, limit int, Nid_ int64, Type_ string, Langcode_ string) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("nid = ? and type = ? and langcode = ?", Nid_, Type_, Langcode_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByNidAndTypeAndTitle Get NodeFieldDatas via NidAndTypeAndTitle
func GetNodeFieldDatasByNidAndTypeAndTitle(offset int, limit int, Nid_ int64, Type_ string, Title_ string) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("nid = ? and type = ? and title = ?", Nid_, Type_, Title_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByNidAndTypeAndUid Get NodeFieldDatas via NidAndTypeAndUid
func GetNodeFieldDatasByNidAndTypeAndUid(offset int, limit int, Nid_ int64, Type_ string, Uid_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("nid = ? and type = ? and uid = ?", Nid_, Type_, Uid_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByNidAndTypeAndStatus Get NodeFieldDatas via NidAndTypeAndStatus
func GetNodeFieldDatasByNidAndTypeAndStatus(offset int, limit int, Nid_ int64, Type_ string, Status_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("nid = ? and type = ? and status = ?", Nid_, Type_, Status_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByNidAndTypeAndCreated Get NodeFieldDatas via NidAndTypeAndCreated
func GetNodeFieldDatasByNidAndTypeAndCreated(offset int, limit int, Nid_ int64, Type_ string, Created_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("nid = ? and type = ? and created = ?", Nid_, Type_, Created_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByNidAndTypeAndChanged Get NodeFieldDatas via NidAndTypeAndChanged
func GetNodeFieldDatasByNidAndTypeAndChanged(offset int, limit int, Nid_ int64, Type_ string, Changed_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("nid = ? and type = ? and changed = ?", Nid_, Type_, Changed_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByNidAndTypeAndPromote Get NodeFieldDatas via NidAndTypeAndPromote
func GetNodeFieldDatasByNidAndTypeAndPromote(offset int, limit int, Nid_ int64, Type_ string, Promote_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("nid = ? and type = ? and promote = ?", Nid_, Type_, Promote_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByNidAndTypeAndSticky Get NodeFieldDatas via NidAndTypeAndSticky
func GetNodeFieldDatasByNidAndTypeAndSticky(offset int, limit int, Nid_ int64, Type_ string, Sticky_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("nid = ? and type = ? and sticky = ?", Nid_, Type_, Sticky_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByNidAndTypeAndRevisionTranslationAffected Get NodeFieldDatas via NidAndTypeAndRevisionTranslationAffected
func GetNodeFieldDatasByNidAndTypeAndRevisionTranslationAffected(offset int, limit int, Nid_ int64, Type_ string, RevisionTranslationAffected_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("nid = ? and type = ? and revision_translation_affected = ?", Nid_, Type_, RevisionTranslationAffected_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByNidAndTypeAndDefaultLangcode Get NodeFieldDatas via NidAndTypeAndDefaultLangcode
func GetNodeFieldDatasByNidAndTypeAndDefaultLangcode(offset int, limit int, Nid_ int64, Type_ string, DefaultLangcode_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("nid = ? and type = ? and default_langcode = ?", Nid_, Type_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByNidAndLangcodeAndTitle Get NodeFieldDatas via NidAndLangcodeAndTitle
func GetNodeFieldDatasByNidAndLangcodeAndTitle(offset int, limit int, Nid_ int64, Langcode_ string, Title_ string) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("nid = ? and langcode = ? and title = ?", Nid_, Langcode_, Title_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByNidAndLangcodeAndUid Get NodeFieldDatas via NidAndLangcodeAndUid
func GetNodeFieldDatasByNidAndLangcodeAndUid(offset int, limit int, Nid_ int64, Langcode_ string, Uid_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("nid = ? and langcode = ? and uid = ?", Nid_, Langcode_, Uid_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByNidAndLangcodeAndStatus Get NodeFieldDatas via NidAndLangcodeAndStatus
func GetNodeFieldDatasByNidAndLangcodeAndStatus(offset int, limit int, Nid_ int64, Langcode_ string, Status_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("nid = ? and langcode = ? and status = ?", Nid_, Langcode_, Status_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByNidAndLangcodeAndCreated Get NodeFieldDatas via NidAndLangcodeAndCreated
func GetNodeFieldDatasByNidAndLangcodeAndCreated(offset int, limit int, Nid_ int64, Langcode_ string, Created_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("nid = ? and langcode = ? and created = ?", Nid_, Langcode_, Created_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByNidAndLangcodeAndChanged Get NodeFieldDatas via NidAndLangcodeAndChanged
func GetNodeFieldDatasByNidAndLangcodeAndChanged(offset int, limit int, Nid_ int64, Langcode_ string, Changed_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("nid = ? and langcode = ? and changed = ?", Nid_, Langcode_, Changed_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByNidAndLangcodeAndPromote Get NodeFieldDatas via NidAndLangcodeAndPromote
func GetNodeFieldDatasByNidAndLangcodeAndPromote(offset int, limit int, Nid_ int64, Langcode_ string, Promote_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("nid = ? and langcode = ? and promote = ?", Nid_, Langcode_, Promote_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByNidAndLangcodeAndSticky Get NodeFieldDatas via NidAndLangcodeAndSticky
func GetNodeFieldDatasByNidAndLangcodeAndSticky(offset int, limit int, Nid_ int64, Langcode_ string, Sticky_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("nid = ? and langcode = ? and sticky = ?", Nid_, Langcode_, Sticky_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByNidAndLangcodeAndRevisionTranslationAffected Get NodeFieldDatas via NidAndLangcodeAndRevisionTranslationAffected
func GetNodeFieldDatasByNidAndLangcodeAndRevisionTranslationAffected(offset int, limit int, Nid_ int64, Langcode_ string, RevisionTranslationAffected_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("nid = ? and langcode = ? and revision_translation_affected = ?", Nid_, Langcode_, RevisionTranslationAffected_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByNidAndLangcodeAndDefaultLangcode Get NodeFieldDatas via NidAndLangcodeAndDefaultLangcode
func GetNodeFieldDatasByNidAndLangcodeAndDefaultLangcode(offset int, limit int, Nid_ int64, Langcode_ string, DefaultLangcode_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("nid = ? and langcode = ? and default_langcode = ?", Nid_, Langcode_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByNidAndTitleAndUid Get NodeFieldDatas via NidAndTitleAndUid
func GetNodeFieldDatasByNidAndTitleAndUid(offset int, limit int, Nid_ int64, Title_ string, Uid_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("nid = ? and title = ? and uid = ?", Nid_, Title_, Uid_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByNidAndTitleAndStatus Get NodeFieldDatas via NidAndTitleAndStatus
func GetNodeFieldDatasByNidAndTitleAndStatus(offset int, limit int, Nid_ int64, Title_ string, Status_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("nid = ? and title = ? and status = ?", Nid_, Title_, Status_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByNidAndTitleAndCreated Get NodeFieldDatas via NidAndTitleAndCreated
func GetNodeFieldDatasByNidAndTitleAndCreated(offset int, limit int, Nid_ int64, Title_ string, Created_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("nid = ? and title = ? and created = ?", Nid_, Title_, Created_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByNidAndTitleAndChanged Get NodeFieldDatas via NidAndTitleAndChanged
func GetNodeFieldDatasByNidAndTitleAndChanged(offset int, limit int, Nid_ int64, Title_ string, Changed_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("nid = ? and title = ? and changed = ?", Nid_, Title_, Changed_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByNidAndTitleAndPromote Get NodeFieldDatas via NidAndTitleAndPromote
func GetNodeFieldDatasByNidAndTitleAndPromote(offset int, limit int, Nid_ int64, Title_ string, Promote_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("nid = ? and title = ? and promote = ?", Nid_, Title_, Promote_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByNidAndTitleAndSticky Get NodeFieldDatas via NidAndTitleAndSticky
func GetNodeFieldDatasByNidAndTitleAndSticky(offset int, limit int, Nid_ int64, Title_ string, Sticky_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("nid = ? and title = ? and sticky = ?", Nid_, Title_, Sticky_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByNidAndTitleAndRevisionTranslationAffected Get NodeFieldDatas via NidAndTitleAndRevisionTranslationAffected
func GetNodeFieldDatasByNidAndTitleAndRevisionTranslationAffected(offset int, limit int, Nid_ int64, Title_ string, RevisionTranslationAffected_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("nid = ? and title = ? and revision_translation_affected = ?", Nid_, Title_, RevisionTranslationAffected_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByNidAndTitleAndDefaultLangcode Get NodeFieldDatas via NidAndTitleAndDefaultLangcode
func GetNodeFieldDatasByNidAndTitleAndDefaultLangcode(offset int, limit int, Nid_ int64, Title_ string, DefaultLangcode_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("nid = ? and title = ? and default_langcode = ?", Nid_, Title_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByNidAndUidAndStatus Get NodeFieldDatas via NidAndUidAndStatus
func GetNodeFieldDatasByNidAndUidAndStatus(offset int, limit int, Nid_ int64, Uid_ int, Status_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("nid = ? and uid = ? and status = ?", Nid_, Uid_, Status_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByNidAndUidAndCreated Get NodeFieldDatas via NidAndUidAndCreated
func GetNodeFieldDatasByNidAndUidAndCreated(offset int, limit int, Nid_ int64, Uid_ int, Created_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("nid = ? and uid = ? and created = ?", Nid_, Uid_, Created_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByNidAndUidAndChanged Get NodeFieldDatas via NidAndUidAndChanged
func GetNodeFieldDatasByNidAndUidAndChanged(offset int, limit int, Nid_ int64, Uid_ int, Changed_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("nid = ? and uid = ? and changed = ?", Nid_, Uid_, Changed_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByNidAndUidAndPromote Get NodeFieldDatas via NidAndUidAndPromote
func GetNodeFieldDatasByNidAndUidAndPromote(offset int, limit int, Nid_ int64, Uid_ int, Promote_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("nid = ? and uid = ? and promote = ?", Nid_, Uid_, Promote_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByNidAndUidAndSticky Get NodeFieldDatas via NidAndUidAndSticky
func GetNodeFieldDatasByNidAndUidAndSticky(offset int, limit int, Nid_ int64, Uid_ int, Sticky_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("nid = ? and uid = ? and sticky = ?", Nid_, Uid_, Sticky_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByNidAndUidAndRevisionTranslationAffected Get NodeFieldDatas via NidAndUidAndRevisionTranslationAffected
func GetNodeFieldDatasByNidAndUidAndRevisionTranslationAffected(offset int, limit int, Nid_ int64, Uid_ int, RevisionTranslationAffected_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("nid = ? and uid = ? and revision_translation_affected = ?", Nid_, Uid_, RevisionTranslationAffected_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByNidAndUidAndDefaultLangcode Get NodeFieldDatas via NidAndUidAndDefaultLangcode
func GetNodeFieldDatasByNidAndUidAndDefaultLangcode(offset int, limit int, Nid_ int64, Uid_ int, DefaultLangcode_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("nid = ? and uid = ? and default_langcode = ?", Nid_, Uid_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByNidAndStatusAndCreated Get NodeFieldDatas via NidAndStatusAndCreated
func GetNodeFieldDatasByNidAndStatusAndCreated(offset int, limit int, Nid_ int64, Status_ int, Created_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("nid = ? and status = ? and created = ?", Nid_, Status_, Created_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByNidAndStatusAndChanged Get NodeFieldDatas via NidAndStatusAndChanged
func GetNodeFieldDatasByNidAndStatusAndChanged(offset int, limit int, Nid_ int64, Status_ int, Changed_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("nid = ? and status = ? and changed = ?", Nid_, Status_, Changed_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByNidAndStatusAndPromote Get NodeFieldDatas via NidAndStatusAndPromote
func GetNodeFieldDatasByNidAndStatusAndPromote(offset int, limit int, Nid_ int64, Status_ int, Promote_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("nid = ? and status = ? and promote = ?", Nid_, Status_, Promote_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByNidAndStatusAndSticky Get NodeFieldDatas via NidAndStatusAndSticky
func GetNodeFieldDatasByNidAndStatusAndSticky(offset int, limit int, Nid_ int64, Status_ int, Sticky_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("nid = ? and status = ? and sticky = ?", Nid_, Status_, Sticky_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByNidAndStatusAndRevisionTranslationAffected Get NodeFieldDatas via NidAndStatusAndRevisionTranslationAffected
func GetNodeFieldDatasByNidAndStatusAndRevisionTranslationAffected(offset int, limit int, Nid_ int64, Status_ int, RevisionTranslationAffected_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("nid = ? and status = ? and revision_translation_affected = ?", Nid_, Status_, RevisionTranslationAffected_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByNidAndStatusAndDefaultLangcode Get NodeFieldDatas via NidAndStatusAndDefaultLangcode
func GetNodeFieldDatasByNidAndStatusAndDefaultLangcode(offset int, limit int, Nid_ int64, Status_ int, DefaultLangcode_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("nid = ? and status = ? and default_langcode = ?", Nid_, Status_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByNidAndCreatedAndChanged Get NodeFieldDatas via NidAndCreatedAndChanged
func GetNodeFieldDatasByNidAndCreatedAndChanged(offset int, limit int, Nid_ int64, Created_ int, Changed_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("nid = ? and created = ? and changed = ?", Nid_, Created_, Changed_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByNidAndCreatedAndPromote Get NodeFieldDatas via NidAndCreatedAndPromote
func GetNodeFieldDatasByNidAndCreatedAndPromote(offset int, limit int, Nid_ int64, Created_ int, Promote_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("nid = ? and created = ? and promote = ?", Nid_, Created_, Promote_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByNidAndCreatedAndSticky Get NodeFieldDatas via NidAndCreatedAndSticky
func GetNodeFieldDatasByNidAndCreatedAndSticky(offset int, limit int, Nid_ int64, Created_ int, Sticky_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("nid = ? and created = ? and sticky = ?", Nid_, Created_, Sticky_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByNidAndCreatedAndRevisionTranslationAffected Get NodeFieldDatas via NidAndCreatedAndRevisionTranslationAffected
func GetNodeFieldDatasByNidAndCreatedAndRevisionTranslationAffected(offset int, limit int, Nid_ int64, Created_ int, RevisionTranslationAffected_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("nid = ? and created = ? and revision_translation_affected = ?", Nid_, Created_, RevisionTranslationAffected_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByNidAndCreatedAndDefaultLangcode Get NodeFieldDatas via NidAndCreatedAndDefaultLangcode
func GetNodeFieldDatasByNidAndCreatedAndDefaultLangcode(offset int, limit int, Nid_ int64, Created_ int, DefaultLangcode_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("nid = ? and created = ? and default_langcode = ?", Nid_, Created_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByNidAndChangedAndPromote Get NodeFieldDatas via NidAndChangedAndPromote
func GetNodeFieldDatasByNidAndChangedAndPromote(offset int, limit int, Nid_ int64, Changed_ int, Promote_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("nid = ? and changed = ? and promote = ?", Nid_, Changed_, Promote_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByNidAndChangedAndSticky Get NodeFieldDatas via NidAndChangedAndSticky
func GetNodeFieldDatasByNidAndChangedAndSticky(offset int, limit int, Nid_ int64, Changed_ int, Sticky_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("nid = ? and changed = ? and sticky = ?", Nid_, Changed_, Sticky_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByNidAndChangedAndRevisionTranslationAffected Get NodeFieldDatas via NidAndChangedAndRevisionTranslationAffected
func GetNodeFieldDatasByNidAndChangedAndRevisionTranslationAffected(offset int, limit int, Nid_ int64, Changed_ int, RevisionTranslationAffected_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("nid = ? and changed = ? and revision_translation_affected = ?", Nid_, Changed_, RevisionTranslationAffected_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByNidAndChangedAndDefaultLangcode Get NodeFieldDatas via NidAndChangedAndDefaultLangcode
func GetNodeFieldDatasByNidAndChangedAndDefaultLangcode(offset int, limit int, Nid_ int64, Changed_ int, DefaultLangcode_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("nid = ? and changed = ? and default_langcode = ?", Nid_, Changed_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByNidAndPromoteAndSticky Get NodeFieldDatas via NidAndPromoteAndSticky
func GetNodeFieldDatasByNidAndPromoteAndSticky(offset int, limit int, Nid_ int64, Promote_ int, Sticky_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("nid = ? and promote = ? and sticky = ?", Nid_, Promote_, Sticky_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByNidAndPromoteAndRevisionTranslationAffected Get NodeFieldDatas via NidAndPromoteAndRevisionTranslationAffected
func GetNodeFieldDatasByNidAndPromoteAndRevisionTranslationAffected(offset int, limit int, Nid_ int64, Promote_ int, RevisionTranslationAffected_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("nid = ? and promote = ? and revision_translation_affected = ?", Nid_, Promote_, RevisionTranslationAffected_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByNidAndPromoteAndDefaultLangcode Get NodeFieldDatas via NidAndPromoteAndDefaultLangcode
func GetNodeFieldDatasByNidAndPromoteAndDefaultLangcode(offset int, limit int, Nid_ int64, Promote_ int, DefaultLangcode_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("nid = ? and promote = ? and default_langcode = ?", Nid_, Promote_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByNidAndStickyAndRevisionTranslationAffected Get NodeFieldDatas via NidAndStickyAndRevisionTranslationAffected
func GetNodeFieldDatasByNidAndStickyAndRevisionTranslationAffected(offset int, limit int, Nid_ int64, Sticky_ int, RevisionTranslationAffected_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("nid = ? and sticky = ? and revision_translation_affected = ?", Nid_, Sticky_, RevisionTranslationAffected_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByNidAndStickyAndDefaultLangcode Get NodeFieldDatas via NidAndStickyAndDefaultLangcode
func GetNodeFieldDatasByNidAndStickyAndDefaultLangcode(offset int, limit int, Nid_ int64, Sticky_ int, DefaultLangcode_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("nid = ? and sticky = ? and default_langcode = ?", Nid_, Sticky_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByNidAndRevisionTranslationAffectedAndDefaultLangcode Get NodeFieldDatas via NidAndRevisionTranslationAffectedAndDefaultLangcode
func GetNodeFieldDatasByNidAndRevisionTranslationAffectedAndDefaultLangcode(offset int, limit int, Nid_ int64, RevisionTranslationAffected_ int, DefaultLangcode_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("nid = ? and revision_translation_affected = ? and default_langcode = ?", Nid_, RevisionTranslationAffected_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByVidAndTypeAndLangcode Get NodeFieldDatas via VidAndTypeAndLangcode
func GetNodeFieldDatasByVidAndTypeAndLangcode(offset int, limit int, Vid_ int, Type_ string, Langcode_ string) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("vid = ? and type = ? and langcode = ?", Vid_, Type_, Langcode_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByVidAndTypeAndTitle Get NodeFieldDatas via VidAndTypeAndTitle
func GetNodeFieldDatasByVidAndTypeAndTitle(offset int, limit int, Vid_ int, Type_ string, Title_ string) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("vid = ? and type = ? and title = ?", Vid_, Type_, Title_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByVidAndTypeAndUid Get NodeFieldDatas via VidAndTypeAndUid
func GetNodeFieldDatasByVidAndTypeAndUid(offset int, limit int, Vid_ int, Type_ string, Uid_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("vid = ? and type = ? and uid = ?", Vid_, Type_, Uid_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByVidAndTypeAndStatus Get NodeFieldDatas via VidAndTypeAndStatus
func GetNodeFieldDatasByVidAndTypeAndStatus(offset int, limit int, Vid_ int, Type_ string, Status_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("vid = ? and type = ? and status = ?", Vid_, Type_, Status_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByVidAndTypeAndCreated Get NodeFieldDatas via VidAndTypeAndCreated
func GetNodeFieldDatasByVidAndTypeAndCreated(offset int, limit int, Vid_ int, Type_ string, Created_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("vid = ? and type = ? and created = ?", Vid_, Type_, Created_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByVidAndTypeAndChanged Get NodeFieldDatas via VidAndTypeAndChanged
func GetNodeFieldDatasByVidAndTypeAndChanged(offset int, limit int, Vid_ int, Type_ string, Changed_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("vid = ? and type = ? and changed = ?", Vid_, Type_, Changed_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByVidAndTypeAndPromote Get NodeFieldDatas via VidAndTypeAndPromote
func GetNodeFieldDatasByVidAndTypeAndPromote(offset int, limit int, Vid_ int, Type_ string, Promote_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("vid = ? and type = ? and promote = ?", Vid_, Type_, Promote_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByVidAndTypeAndSticky Get NodeFieldDatas via VidAndTypeAndSticky
func GetNodeFieldDatasByVidAndTypeAndSticky(offset int, limit int, Vid_ int, Type_ string, Sticky_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("vid = ? and type = ? and sticky = ?", Vid_, Type_, Sticky_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByVidAndTypeAndRevisionTranslationAffected Get NodeFieldDatas via VidAndTypeAndRevisionTranslationAffected
func GetNodeFieldDatasByVidAndTypeAndRevisionTranslationAffected(offset int, limit int, Vid_ int, Type_ string, RevisionTranslationAffected_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("vid = ? and type = ? and revision_translation_affected = ?", Vid_, Type_, RevisionTranslationAffected_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByVidAndTypeAndDefaultLangcode Get NodeFieldDatas via VidAndTypeAndDefaultLangcode
func GetNodeFieldDatasByVidAndTypeAndDefaultLangcode(offset int, limit int, Vid_ int, Type_ string, DefaultLangcode_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("vid = ? and type = ? and default_langcode = ?", Vid_, Type_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByVidAndLangcodeAndTitle Get NodeFieldDatas via VidAndLangcodeAndTitle
func GetNodeFieldDatasByVidAndLangcodeAndTitle(offset int, limit int, Vid_ int, Langcode_ string, Title_ string) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("vid = ? and langcode = ? and title = ?", Vid_, Langcode_, Title_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByVidAndLangcodeAndUid Get NodeFieldDatas via VidAndLangcodeAndUid
func GetNodeFieldDatasByVidAndLangcodeAndUid(offset int, limit int, Vid_ int, Langcode_ string, Uid_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("vid = ? and langcode = ? and uid = ?", Vid_, Langcode_, Uid_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByVidAndLangcodeAndStatus Get NodeFieldDatas via VidAndLangcodeAndStatus
func GetNodeFieldDatasByVidAndLangcodeAndStatus(offset int, limit int, Vid_ int, Langcode_ string, Status_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("vid = ? and langcode = ? and status = ?", Vid_, Langcode_, Status_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByVidAndLangcodeAndCreated Get NodeFieldDatas via VidAndLangcodeAndCreated
func GetNodeFieldDatasByVidAndLangcodeAndCreated(offset int, limit int, Vid_ int, Langcode_ string, Created_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("vid = ? and langcode = ? and created = ?", Vid_, Langcode_, Created_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByVidAndLangcodeAndChanged Get NodeFieldDatas via VidAndLangcodeAndChanged
func GetNodeFieldDatasByVidAndLangcodeAndChanged(offset int, limit int, Vid_ int, Langcode_ string, Changed_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("vid = ? and langcode = ? and changed = ?", Vid_, Langcode_, Changed_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByVidAndLangcodeAndPromote Get NodeFieldDatas via VidAndLangcodeAndPromote
func GetNodeFieldDatasByVidAndLangcodeAndPromote(offset int, limit int, Vid_ int, Langcode_ string, Promote_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("vid = ? and langcode = ? and promote = ?", Vid_, Langcode_, Promote_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByVidAndLangcodeAndSticky Get NodeFieldDatas via VidAndLangcodeAndSticky
func GetNodeFieldDatasByVidAndLangcodeAndSticky(offset int, limit int, Vid_ int, Langcode_ string, Sticky_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("vid = ? and langcode = ? and sticky = ?", Vid_, Langcode_, Sticky_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByVidAndLangcodeAndRevisionTranslationAffected Get NodeFieldDatas via VidAndLangcodeAndRevisionTranslationAffected
func GetNodeFieldDatasByVidAndLangcodeAndRevisionTranslationAffected(offset int, limit int, Vid_ int, Langcode_ string, RevisionTranslationAffected_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("vid = ? and langcode = ? and revision_translation_affected = ?", Vid_, Langcode_, RevisionTranslationAffected_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByVidAndLangcodeAndDefaultLangcode Get NodeFieldDatas via VidAndLangcodeAndDefaultLangcode
func GetNodeFieldDatasByVidAndLangcodeAndDefaultLangcode(offset int, limit int, Vid_ int, Langcode_ string, DefaultLangcode_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("vid = ? and langcode = ? and default_langcode = ?", Vid_, Langcode_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByVidAndTitleAndUid Get NodeFieldDatas via VidAndTitleAndUid
func GetNodeFieldDatasByVidAndTitleAndUid(offset int, limit int, Vid_ int, Title_ string, Uid_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("vid = ? and title = ? and uid = ?", Vid_, Title_, Uid_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByVidAndTitleAndStatus Get NodeFieldDatas via VidAndTitleAndStatus
func GetNodeFieldDatasByVidAndTitleAndStatus(offset int, limit int, Vid_ int, Title_ string, Status_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("vid = ? and title = ? and status = ?", Vid_, Title_, Status_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByVidAndTitleAndCreated Get NodeFieldDatas via VidAndTitleAndCreated
func GetNodeFieldDatasByVidAndTitleAndCreated(offset int, limit int, Vid_ int, Title_ string, Created_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("vid = ? and title = ? and created = ?", Vid_, Title_, Created_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByVidAndTitleAndChanged Get NodeFieldDatas via VidAndTitleAndChanged
func GetNodeFieldDatasByVidAndTitleAndChanged(offset int, limit int, Vid_ int, Title_ string, Changed_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("vid = ? and title = ? and changed = ?", Vid_, Title_, Changed_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByVidAndTitleAndPromote Get NodeFieldDatas via VidAndTitleAndPromote
func GetNodeFieldDatasByVidAndTitleAndPromote(offset int, limit int, Vid_ int, Title_ string, Promote_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("vid = ? and title = ? and promote = ?", Vid_, Title_, Promote_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByVidAndTitleAndSticky Get NodeFieldDatas via VidAndTitleAndSticky
func GetNodeFieldDatasByVidAndTitleAndSticky(offset int, limit int, Vid_ int, Title_ string, Sticky_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("vid = ? and title = ? and sticky = ?", Vid_, Title_, Sticky_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByVidAndTitleAndRevisionTranslationAffected Get NodeFieldDatas via VidAndTitleAndRevisionTranslationAffected
func GetNodeFieldDatasByVidAndTitleAndRevisionTranslationAffected(offset int, limit int, Vid_ int, Title_ string, RevisionTranslationAffected_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("vid = ? and title = ? and revision_translation_affected = ?", Vid_, Title_, RevisionTranslationAffected_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByVidAndTitleAndDefaultLangcode Get NodeFieldDatas via VidAndTitleAndDefaultLangcode
func GetNodeFieldDatasByVidAndTitleAndDefaultLangcode(offset int, limit int, Vid_ int, Title_ string, DefaultLangcode_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("vid = ? and title = ? and default_langcode = ?", Vid_, Title_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByVidAndUidAndStatus Get NodeFieldDatas via VidAndUidAndStatus
func GetNodeFieldDatasByVidAndUidAndStatus(offset int, limit int, Vid_ int, Uid_ int, Status_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("vid = ? and uid = ? and status = ?", Vid_, Uid_, Status_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByVidAndUidAndCreated Get NodeFieldDatas via VidAndUidAndCreated
func GetNodeFieldDatasByVidAndUidAndCreated(offset int, limit int, Vid_ int, Uid_ int, Created_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("vid = ? and uid = ? and created = ?", Vid_, Uid_, Created_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByVidAndUidAndChanged Get NodeFieldDatas via VidAndUidAndChanged
func GetNodeFieldDatasByVidAndUidAndChanged(offset int, limit int, Vid_ int, Uid_ int, Changed_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("vid = ? and uid = ? and changed = ?", Vid_, Uid_, Changed_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByVidAndUidAndPromote Get NodeFieldDatas via VidAndUidAndPromote
func GetNodeFieldDatasByVidAndUidAndPromote(offset int, limit int, Vid_ int, Uid_ int, Promote_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("vid = ? and uid = ? and promote = ?", Vid_, Uid_, Promote_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByVidAndUidAndSticky Get NodeFieldDatas via VidAndUidAndSticky
func GetNodeFieldDatasByVidAndUidAndSticky(offset int, limit int, Vid_ int, Uid_ int, Sticky_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("vid = ? and uid = ? and sticky = ?", Vid_, Uid_, Sticky_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByVidAndUidAndRevisionTranslationAffected Get NodeFieldDatas via VidAndUidAndRevisionTranslationAffected
func GetNodeFieldDatasByVidAndUidAndRevisionTranslationAffected(offset int, limit int, Vid_ int, Uid_ int, RevisionTranslationAffected_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("vid = ? and uid = ? and revision_translation_affected = ?", Vid_, Uid_, RevisionTranslationAffected_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByVidAndUidAndDefaultLangcode Get NodeFieldDatas via VidAndUidAndDefaultLangcode
func GetNodeFieldDatasByVidAndUidAndDefaultLangcode(offset int, limit int, Vid_ int, Uid_ int, DefaultLangcode_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("vid = ? and uid = ? and default_langcode = ?", Vid_, Uid_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByVidAndStatusAndCreated Get NodeFieldDatas via VidAndStatusAndCreated
func GetNodeFieldDatasByVidAndStatusAndCreated(offset int, limit int, Vid_ int, Status_ int, Created_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("vid = ? and status = ? and created = ?", Vid_, Status_, Created_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByVidAndStatusAndChanged Get NodeFieldDatas via VidAndStatusAndChanged
func GetNodeFieldDatasByVidAndStatusAndChanged(offset int, limit int, Vid_ int, Status_ int, Changed_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("vid = ? and status = ? and changed = ?", Vid_, Status_, Changed_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByVidAndStatusAndPromote Get NodeFieldDatas via VidAndStatusAndPromote
func GetNodeFieldDatasByVidAndStatusAndPromote(offset int, limit int, Vid_ int, Status_ int, Promote_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("vid = ? and status = ? and promote = ?", Vid_, Status_, Promote_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByVidAndStatusAndSticky Get NodeFieldDatas via VidAndStatusAndSticky
func GetNodeFieldDatasByVidAndStatusAndSticky(offset int, limit int, Vid_ int, Status_ int, Sticky_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("vid = ? and status = ? and sticky = ?", Vid_, Status_, Sticky_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByVidAndStatusAndRevisionTranslationAffected Get NodeFieldDatas via VidAndStatusAndRevisionTranslationAffected
func GetNodeFieldDatasByVidAndStatusAndRevisionTranslationAffected(offset int, limit int, Vid_ int, Status_ int, RevisionTranslationAffected_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("vid = ? and status = ? and revision_translation_affected = ?", Vid_, Status_, RevisionTranslationAffected_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByVidAndStatusAndDefaultLangcode Get NodeFieldDatas via VidAndStatusAndDefaultLangcode
func GetNodeFieldDatasByVidAndStatusAndDefaultLangcode(offset int, limit int, Vid_ int, Status_ int, DefaultLangcode_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("vid = ? and status = ? and default_langcode = ?", Vid_, Status_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByVidAndCreatedAndChanged Get NodeFieldDatas via VidAndCreatedAndChanged
func GetNodeFieldDatasByVidAndCreatedAndChanged(offset int, limit int, Vid_ int, Created_ int, Changed_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("vid = ? and created = ? and changed = ?", Vid_, Created_, Changed_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByVidAndCreatedAndPromote Get NodeFieldDatas via VidAndCreatedAndPromote
func GetNodeFieldDatasByVidAndCreatedAndPromote(offset int, limit int, Vid_ int, Created_ int, Promote_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("vid = ? and created = ? and promote = ?", Vid_, Created_, Promote_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByVidAndCreatedAndSticky Get NodeFieldDatas via VidAndCreatedAndSticky
func GetNodeFieldDatasByVidAndCreatedAndSticky(offset int, limit int, Vid_ int, Created_ int, Sticky_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("vid = ? and created = ? and sticky = ?", Vid_, Created_, Sticky_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByVidAndCreatedAndRevisionTranslationAffected Get NodeFieldDatas via VidAndCreatedAndRevisionTranslationAffected
func GetNodeFieldDatasByVidAndCreatedAndRevisionTranslationAffected(offset int, limit int, Vid_ int, Created_ int, RevisionTranslationAffected_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("vid = ? and created = ? and revision_translation_affected = ?", Vid_, Created_, RevisionTranslationAffected_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByVidAndCreatedAndDefaultLangcode Get NodeFieldDatas via VidAndCreatedAndDefaultLangcode
func GetNodeFieldDatasByVidAndCreatedAndDefaultLangcode(offset int, limit int, Vid_ int, Created_ int, DefaultLangcode_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("vid = ? and created = ? and default_langcode = ?", Vid_, Created_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByVidAndChangedAndPromote Get NodeFieldDatas via VidAndChangedAndPromote
func GetNodeFieldDatasByVidAndChangedAndPromote(offset int, limit int, Vid_ int, Changed_ int, Promote_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("vid = ? and changed = ? and promote = ?", Vid_, Changed_, Promote_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByVidAndChangedAndSticky Get NodeFieldDatas via VidAndChangedAndSticky
func GetNodeFieldDatasByVidAndChangedAndSticky(offset int, limit int, Vid_ int, Changed_ int, Sticky_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("vid = ? and changed = ? and sticky = ?", Vid_, Changed_, Sticky_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByVidAndChangedAndRevisionTranslationAffected Get NodeFieldDatas via VidAndChangedAndRevisionTranslationAffected
func GetNodeFieldDatasByVidAndChangedAndRevisionTranslationAffected(offset int, limit int, Vid_ int, Changed_ int, RevisionTranslationAffected_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("vid = ? and changed = ? and revision_translation_affected = ?", Vid_, Changed_, RevisionTranslationAffected_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByVidAndChangedAndDefaultLangcode Get NodeFieldDatas via VidAndChangedAndDefaultLangcode
func GetNodeFieldDatasByVidAndChangedAndDefaultLangcode(offset int, limit int, Vid_ int, Changed_ int, DefaultLangcode_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("vid = ? and changed = ? and default_langcode = ?", Vid_, Changed_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByVidAndPromoteAndSticky Get NodeFieldDatas via VidAndPromoteAndSticky
func GetNodeFieldDatasByVidAndPromoteAndSticky(offset int, limit int, Vid_ int, Promote_ int, Sticky_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("vid = ? and promote = ? and sticky = ?", Vid_, Promote_, Sticky_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByVidAndPromoteAndRevisionTranslationAffected Get NodeFieldDatas via VidAndPromoteAndRevisionTranslationAffected
func GetNodeFieldDatasByVidAndPromoteAndRevisionTranslationAffected(offset int, limit int, Vid_ int, Promote_ int, RevisionTranslationAffected_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("vid = ? and promote = ? and revision_translation_affected = ?", Vid_, Promote_, RevisionTranslationAffected_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByVidAndPromoteAndDefaultLangcode Get NodeFieldDatas via VidAndPromoteAndDefaultLangcode
func GetNodeFieldDatasByVidAndPromoteAndDefaultLangcode(offset int, limit int, Vid_ int, Promote_ int, DefaultLangcode_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("vid = ? and promote = ? and default_langcode = ?", Vid_, Promote_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByVidAndStickyAndRevisionTranslationAffected Get NodeFieldDatas via VidAndStickyAndRevisionTranslationAffected
func GetNodeFieldDatasByVidAndStickyAndRevisionTranslationAffected(offset int, limit int, Vid_ int, Sticky_ int, RevisionTranslationAffected_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("vid = ? and sticky = ? and revision_translation_affected = ?", Vid_, Sticky_, RevisionTranslationAffected_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByVidAndStickyAndDefaultLangcode Get NodeFieldDatas via VidAndStickyAndDefaultLangcode
func GetNodeFieldDatasByVidAndStickyAndDefaultLangcode(offset int, limit int, Vid_ int, Sticky_ int, DefaultLangcode_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("vid = ? and sticky = ? and default_langcode = ?", Vid_, Sticky_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByVidAndRevisionTranslationAffectedAndDefaultLangcode Get NodeFieldDatas via VidAndRevisionTranslationAffectedAndDefaultLangcode
func GetNodeFieldDatasByVidAndRevisionTranslationAffectedAndDefaultLangcode(offset int, limit int, Vid_ int, RevisionTranslationAffected_ int, DefaultLangcode_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("vid = ? and revision_translation_affected = ? and default_langcode = ?", Vid_, RevisionTranslationAffected_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByTypeAndLangcodeAndTitle Get NodeFieldDatas via TypeAndLangcodeAndTitle
func GetNodeFieldDatasByTypeAndLangcodeAndTitle(offset int, limit int, Type_ string, Langcode_ string, Title_ string) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("type = ? and langcode = ? and title = ?", Type_, Langcode_, Title_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByTypeAndLangcodeAndUid Get NodeFieldDatas via TypeAndLangcodeAndUid
func GetNodeFieldDatasByTypeAndLangcodeAndUid(offset int, limit int, Type_ string, Langcode_ string, Uid_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("type = ? and langcode = ? and uid = ?", Type_, Langcode_, Uid_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByTypeAndLangcodeAndStatus Get NodeFieldDatas via TypeAndLangcodeAndStatus
func GetNodeFieldDatasByTypeAndLangcodeAndStatus(offset int, limit int, Type_ string, Langcode_ string, Status_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("type = ? and langcode = ? and status = ?", Type_, Langcode_, Status_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByTypeAndLangcodeAndCreated Get NodeFieldDatas via TypeAndLangcodeAndCreated
func GetNodeFieldDatasByTypeAndLangcodeAndCreated(offset int, limit int, Type_ string, Langcode_ string, Created_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("type = ? and langcode = ? and created = ?", Type_, Langcode_, Created_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByTypeAndLangcodeAndChanged Get NodeFieldDatas via TypeAndLangcodeAndChanged
func GetNodeFieldDatasByTypeAndLangcodeAndChanged(offset int, limit int, Type_ string, Langcode_ string, Changed_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("type = ? and langcode = ? and changed = ?", Type_, Langcode_, Changed_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByTypeAndLangcodeAndPromote Get NodeFieldDatas via TypeAndLangcodeAndPromote
func GetNodeFieldDatasByTypeAndLangcodeAndPromote(offset int, limit int, Type_ string, Langcode_ string, Promote_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("type = ? and langcode = ? and promote = ?", Type_, Langcode_, Promote_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByTypeAndLangcodeAndSticky Get NodeFieldDatas via TypeAndLangcodeAndSticky
func GetNodeFieldDatasByTypeAndLangcodeAndSticky(offset int, limit int, Type_ string, Langcode_ string, Sticky_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("type = ? and langcode = ? and sticky = ?", Type_, Langcode_, Sticky_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByTypeAndLangcodeAndRevisionTranslationAffected Get NodeFieldDatas via TypeAndLangcodeAndRevisionTranslationAffected
func GetNodeFieldDatasByTypeAndLangcodeAndRevisionTranslationAffected(offset int, limit int, Type_ string, Langcode_ string, RevisionTranslationAffected_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("type = ? and langcode = ? and revision_translation_affected = ?", Type_, Langcode_, RevisionTranslationAffected_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByTypeAndLangcodeAndDefaultLangcode Get NodeFieldDatas via TypeAndLangcodeAndDefaultLangcode
func GetNodeFieldDatasByTypeAndLangcodeAndDefaultLangcode(offset int, limit int, Type_ string, Langcode_ string, DefaultLangcode_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("type = ? and langcode = ? and default_langcode = ?", Type_, Langcode_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByTypeAndTitleAndUid Get NodeFieldDatas via TypeAndTitleAndUid
func GetNodeFieldDatasByTypeAndTitleAndUid(offset int, limit int, Type_ string, Title_ string, Uid_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("type = ? and title = ? and uid = ?", Type_, Title_, Uid_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByTypeAndTitleAndStatus Get NodeFieldDatas via TypeAndTitleAndStatus
func GetNodeFieldDatasByTypeAndTitleAndStatus(offset int, limit int, Type_ string, Title_ string, Status_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("type = ? and title = ? and status = ?", Type_, Title_, Status_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByTypeAndTitleAndCreated Get NodeFieldDatas via TypeAndTitleAndCreated
func GetNodeFieldDatasByTypeAndTitleAndCreated(offset int, limit int, Type_ string, Title_ string, Created_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("type = ? and title = ? and created = ?", Type_, Title_, Created_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByTypeAndTitleAndChanged Get NodeFieldDatas via TypeAndTitleAndChanged
func GetNodeFieldDatasByTypeAndTitleAndChanged(offset int, limit int, Type_ string, Title_ string, Changed_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("type = ? and title = ? and changed = ?", Type_, Title_, Changed_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByTypeAndTitleAndPromote Get NodeFieldDatas via TypeAndTitleAndPromote
func GetNodeFieldDatasByTypeAndTitleAndPromote(offset int, limit int, Type_ string, Title_ string, Promote_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("type = ? and title = ? and promote = ?", Type_, Title_, Promote_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByTypeAndTitleAndSticky Get NodeFieldDatas via TypeAndTitleAndSticky
func GetNodeFieldDatasByTypeAndTitleAndSticky(offset int, limit int, Type_ string, Title_ string, Sticky_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("type = ? and title = ? and sticky = ?", Type_, Title_, Sticky_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByTypeAndTitleAndRevisionTranslationAffected Get NodeFieldDatas via TypeAndTitleAndRevisionTranslationAffected
func GetNodeFieldDatasByTypeAndTitleAndRevisionTranslationAffected(offset int, limit int, Type_ string, Title_ string, RevisionTranslationAffected_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("type = ? and title = ? and revision_translation_affected = ?", Type_, Title_, RevisionTranslationAffected_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByTypeAndTitleAndDefaultLangcode Get NodeFieldDatas via TypeAndTitleAndDefaultLangcode
func GetNodeFieldDatasByTypeAndTitleAndDefaultLangcode(offset int, limit int, Type_ string, Title_ string, DefaultLangcode_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("type = ? and title = ? and default_langcode = ?", Type_, Title_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByTypeAndUidAndStatus Get NodeFieldDatas via TypeAndUidAndStatus
func GetNodeFieldDatasByTypeAndUidAndStatus(offset int, limit int, Type_ string, Uid_ int, Status_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("type = ? and uid = ? and status = ?", Type_, Uid_, Status_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByTypeAndUidAndCreated Get NodeFieldDatas via TypeAndUidAndCreated
func GetNodeFieldDatasByTypeAndUidAndCreated(offset int, limit int, Type_ string, Uid_ int, Created_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("type = ? and uid = ? and created = ?", Type_, Uid_, Created_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByTypeAndUidAndChanged Get NodeFieldDatas via TypeAndUidAndChanged
func GetNodeFieldDatasByTypeAndUidAndChanged(offset int, limit int, Type_ string, Uid_ int, Changed_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("type = ? and uid = ? and changed = ?", Type_, Uid_, Changed_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByTypeAndUidAndPromote Get NodeFieldDatas via TypeAndUidAndPromote
func GetNodeFieldDatasByTypeAndUidAndPromote(offset int, limit int, Type_ string, Uid_ int, Promote_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("type = ? and uid = ? and promote = ?", Type_, Uid_, Promote_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByTypeAndUidAndSticky Get NodeFieldDatas via TypeAndUidAndSticky
func GetNodeFieldDatasByTypeAndUidAndSticky(offset int, limit int, Type_ string, Uid_ int, Sticky_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("type = ? and uid = ? and sticky = ?", Type_, Uid_, Sticky_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByTypeAndUidAndRevisionTranslationAffected Get NodeFieldDatas via TypeAndUidAndRevisionTranslationAffected
func GetNodeFieldDatasByTypeAndUidAndRevisionTranslationAffected(offset int, limit int, Type_ string, Uid_ int, RevisionTranslationAffected_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("type = ? and uid = ? and revision_translation_affected = ?", Type_, Uid_, RevisionTranslationAffected_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByTypeAndUidAndDefaultLangcode Get NodeFieldDatas via TypeAndUidAndDefaultLangcode
func GetNodeFieldDatasByTypeAndUidAndDefaultLangcode(offset int, limit int, Type_ string, Uid_ int, DefaultLangcode_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("type = ? and uid = ? and default_langcode = ?", Type_, Uid_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByTypeAndStatusAndCreated Get NodeFieldDatas via TypeAndStatusAndCreated
func GetNodeFieldDatasByTypeAndStatusAndCreated(offset int, limit int, Type_ string, Status_ int, Created_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("type = ? and status = ? and created = ?", Type_, Status_, Created_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByTypeAndStatusAndChanged Get NodeFieldDatas via TypeAndStatusAndChanged
func GetNodeFieldDatasByTypeAndStatusAndChanged(offset int, limit int, Type_ string, Status_ int, Changed_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("type = ? and status = ? and changed = ?", Type_, Status_, Changed_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByTypeAndStatusAndPromote Get NodeFieldDatas via TypeAndStatusAndPromote
func GetNodeFieldDatasByTypeAndStatusAndPromote(offset int, limit int, Type_ string, Status_ int, Promote_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("type = ? and status = ? and promote = ?", Type_, Status_, Promote_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByTypeAndStatusAndSticky Get NodeFieldDatas via TypeAndStatusAndSticky
func GetNodeFieldDatasByTypeAndStatusAndSticky(offset int, limit int, Type_ string, Status_ int, Sticky_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("type = ? and status = ? and sticky = ?", Type_, Status_, Sticky_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByTypeAndStatusAndRevisionTranslationAffected Get NodeFieldDatas via TypeAndStatusAndRevisionTranslationAffected
func GetNodeFieldDatasByTypeAndStatusAndRevisionTranslationAffected(offset int, limit int, Type_ string, Status_ int, RevisionTranslationAffected_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("type = ? and status = ? and revision_translation_affected = ?", Type_, Status_, RevisionTranslationAffected_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByTypeAndStatusAndDefaultLangcode Get NodeFieldDatas via TypeAndStatusAndDefaultLangcode
func GetNodeFieldDatasByTypeAndStatusAndDefaultLangcode(offset int, limit int, Type_ string, Status_ int, DefaultLangcode_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("type = ? and status = ? and default_langcode = ?", Type_, Status_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByTypeAndCreatedAndChanged Get NodeFieldDatas via TypeAndCreatedAndChanged
func GetNodeFieldDatasByTypeAndCreatedAndChanged(offset int, limit int, Type_ string, Created_ int, Changed_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("type = ? and created = ? and changed = ?", Type_, Created_, Changed_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByTypeAndCreatedAndPromote Get NodeFieldDatas via TypeAndCreatedAndPromote
func GetNodeFieldDatasByTypeAndCreatedAndPromote(offset int, limit int, Type_ string, Created_ int, Promote_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("type = ? and created = ? and promote = ?", Type_, Created_, Promote_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByTypeAndCreatedAndSticky Get NodeFieldDatas via TypeAndCreatedAndSticky
func GetNodeFieldDatasByTypeAndCreatedAndSticky(offset int, limit int, Type_ string, Created_ int, Sticky_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("type = ? and created = ? and sticky = ?", Type_, Created_, Sticky_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByTypeAndCreatedAndRevisionTranslationAffected Get NodeFieldDatas via TypeAndCreatedAndRevisionTranslationAffected
func GetNodeFieldDatasByTypeAndCreatedAndRevisionTranslationAffected(offset int, limit int, Type_ string, Created_ int, RevisionTranslationAffected_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("type = ? and created = ? and revision_translation_affected = ?", Type_, Created_, RevisionTranslationAffected_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByTypeAndCreatedAndDefaultLangcode Get NodeFieldDatas via TypeAndCreatedAndDefaultLangcode
func GetNodeFieldDatasByTypeAndCreatedAndDefaultLangcode(offset int, limit int, Type_ string, Created_ int, DefaultLangcode_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("type = ? and created = ? and default_langcode = ?", Type_, Created_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByTypeAndChangedAndPromote Get NodeFieldDatas via TypeAndChangedAndPromote
func GetNodeFieldDatasByTypeAndChangedAndPromote(offset int, limit int, Type_ string, Changed_ int, Promote_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("type = ? and changed = ? and promote = ?", Type_, Changed_, Promote_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByTypeAndChangedAndSticky Get NodeFieldDatas via TypeAndChangedAndSticky
func GetNodeFieldDatasByTypeAndChangedAndSticky(offset int, limit int, Type_ string, Changed_ int, Sticky_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("type = ? and changed = ? and sticky = ?", Type_, Changed_, Sticky_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByTypeAndChangedAndRevisionTranslationAffected Get NodeFieldDatas via TypeAndChangedAndRevisionTranslationAffected
func GetNodeFieldDatasByTypeAndChangedAndRevisionTranslationAffected(offset int, limit int, Type_ string, Changed_ int, RevisionTranslationAffected_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("type = ? and changed = ? and revision_translation_affected = ?", Type_, Changed_, RevisionTranslationAffected_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByTypeAndChangedAndDefaultLangcode Get NodeFieldDatas via TypeAndChangedAndDefaultLangcode
func GetNodeFieldDatasByTypeAndChangedAndDefaultLangcode(offset int, limit int, Type_ string, Changed_ int, DefaultLangcode_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("type = ? and changed = ? and default_langcode = ?", Type_, Changed_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByTypeAndPromoteAndSticky Get NodeFieldDatas via TypeAndPromoteAndSticky
func GetNodeFieldDatasByTypeAndPromoteAndSticky(offset int, limit int, Type_ string, Promote_ int, Sticky_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("type = ? and promote = ? and sticky = ?", Type_, Promote_, Sticky_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByTypeAndPromoteAndRevisionTranslationAffected Get NodeFieldDatas via TypeAndPromoteAndRevisionTranslationAffected
func GetNodeFieldDatasByTypeAndPromoteAndRevisionTranslationAffected(offset int, limit int, Type_ string, Promote_ int, RevisionTranslationAffected_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("type = ? and promote = ? and revision_translation_affected = ?", Type_, Promote_, RevisionTranslationAffected_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByTypeAndPromoteAndDefaultLangcode Get NodeFieldDatas via TypeAndPromoteAndDefaultLangcode
func GetNodeFieldDatasByTypeAndPromoteAndDefaultLangcode(offset int, limit int, Type_ string, Promote_ int, DefaultLangcode_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("type = ? and promote = ? and default_langcode = ?", Type_, Promote_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByTypeAndStickyAndRevisionTranslationAffected Get NodeFieldDatas via TypeAndStickyAndRevisionTranslationAffected
func GetNodeFieldDatasByTypeAndStickyAndRevisionTranslationAffected(offset int, limit int, Type_ string, Sticky_ int, RevisionTranslationAffected_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("type = ? and sticky = ? and revision_translation_affected = ?", Type_, Sticky_, RevisionTranslationAffected_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByTypeAndStickyAndDefaultLangcode Get NodeFieldDatas via TypeAndStickyAndDefaultLangcode
func GetNodeFieldDatasByTypeAndStickyAndDefaultLangcode(offset int, limit int, Type_ string, Sticky_ int, DefaultLangcode_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("type = ? and sticky = ? and default_langcode = ?", Type_, Sticky_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByTypeAndRevisionTranslationAffectedAndDefaultLangcode Get NodeFieldDatas via TypeAndRevisionTranslationAffectedAndDefaultLangcode
func GetNodeFieldDatasByTypeAndRevisionTranslationAffectedAndDefaultLangcode(offset int, limit int, Type_ string, RevisionTranslationAffected_ int, DefaultLangcode_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("type = ? and revision_translation_affected = ? and default_langcode = ?", Type_, RevisionTranslationAffected_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByLangcodeAndTitleAndUid Get NodeFieldDatas via LangcodeAndTitleAndUid
func GetNodeFieldDatasByLangcodeAndTitleAndUid(offset int, limit int, Langcode_ string, Title_ string, Uid_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("langcode = ? and title = ? and uid = ?", Langcode_, Title_, Uid_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByLangcodeAndTitleAndStatus Get NodeFieldDatas via LangcodeAndTitleAndStatus
func GetNodeFieldDatasByLangcodeAndTitleAndStatus(offset int, limit int, Langcode_ string, Title_ string, Status_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("langcode = ? and title = ? and status = ?", Langcode_, Title_, Status_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByLangcodeAndTitleAndCreated Get NodeFieldDatas via LangcodeAndTitleAndCreated
func GetNodeFieldDatasByLangcodeAndTitleAndCreated(offset int, limit int, Langcode_ string, Title_ string, Created_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("langcode = ? and title = ? and created = ?", Langcode_, Title_, Created_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByLangcodeAndTitleAndChanged Get NodeFieldDatas via LangcodeAndTitleAndChanged
func GetNodeFieldDatasByLangcodeAndTitleAndChanged(offset int, limit int, Langcode_ string, Title_ string, Changed_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("langcode = ? and title = ? and changed = ?", Langcode_, Title_, Changed_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByLangcodeAndTitleAndPromote Get NodeFieldDatas via LangcodeAndTitleAndPromote
func GetNodeFieldDatasByLangcodeAndTitleAndPromote(offset int, limit int, Langcode_ string, Title_ string, Promote_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("langcode = ? and title = ? and promote = ?", Langcode_, Title_, Promote_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByLangcodeAndTitleAndSticky Get NodeFieldDatas via LangcodeAndTitleAndSticky
func GetNodeFieldDatasByLangcodeAndTitleAndSticky(offset int, limit int, Langcode_ string, Title_ string, Sticky_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("langcode = ? and title = ? and sticky = ?", Langcode_, Title_, Sticky_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByLangcodeAndTitleAndRevisionTranslationAffected Get NodeFieldDatas via LangcodeAndTitleAndRevisionTranslationAffected
func GetNodeFieldDatasByLangcodeAndTitleAndRevisionTranslationAffected(offset int, limit int, Langcode_ string, Title_ string, RevisionTranslationAffected_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("langcode = ? and title = ? and revision_translation_affected = ?", Langcode_, Title_, RevisionTranslationAffected_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByLangcodeAndTitleAndDefaultLangcode Get NodeFieldDatas via LangcodeAndTitleAndDefaultLangcode
func GetNodeFieldDatasByLangcodeAndTitleAndDefaultLangcode(offset int, limit int, Langcode_ string, Title_ string, DefaultLangcode_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("langcode = ? and title = ? and default_langcode = ?", Langcode_, Title_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByLangcodeAndUidAndStatus Get NodeFieldDatas via LangcodeAndUidAndStatus
func GetNodeFieldDatasByLangcodeAndUidAndStatus(offset int, limit int, Langcode_ string, Uid_ int, Status_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("langcode = ? and uid = ? and status = ?", Langcode_, Uid_, Status_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByLangcodeAndUidAndCreated Get NodeFieldDatas via LangcodeAndUidAndCreated
func GetNodeFieldDatasByLangcodeAndUidAndCreated(offset int, limit int, Langcode_ string, Uid_ int, Created_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("langcode = ? and uid = ? and created = ?", Langcode_, Uid_, Created_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByLangcodeAndUidAndChanged Get NodeFieldDatas via LangcodeAndUidAndChanged
func GetNodeFieldDatasByLangcodeAndUidAndChanged(offset int, limit int, Langcode_ string, Uid_ int, Changed_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("langcode = ? and uid = ? and changed = ?", Langcode_, Uid_, Changed_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByLangcodeAndUidAndPromote Get NodeFieldDatas via LangcodeAndUidAndPromote
func GetNodeFieldDatasByLangcodeAndUidAndPromote(offset int, limit int, Langcode_ string, Uid_ int, Promote_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("langcode = ? and uid = ? and promote = ?", Langcode_, Uid_, Promote_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByLangcodeAndUidAndSticky Get NodeFieldDatas via LangcodeAndUidAndSticky
func GetNodeFieldDatasByLangcodeAndUidAndSticky(offset int, limit int, Langcode_ string, Uid_ int, Sticky_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("langcode = ? and uid = ? and sticky = ?", Langcode_, Uid_, Sticky_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByLangcodeAndUidAndRevisionTranslationAffected Get NodeFieldDatas via LangcodeAndUidAndRevisionTranslationAffected
func GetNodeFieldDatasByLangcodeAndUidAndRevisionTranslationAffected(offset int, limit int, Langcode_ string, Uid_ int, RevisionTranslationAffected_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("langcode = ? and uid = ? and revision_translation_affected = ?", Langcode_, Uid_, RevisionTranslationAffected_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByLangcodeAndUidAndDefaultLangcode Get NodeFieldDatas via LangcodeAndUidAndDefaultLangcode
func GetNodeFieldDatasByLangcodeAndUidAndDefaultLangcode(offset int, limit int, Langcode_ string, Uid_ int, DefaultLangcode_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("langcode = ? and uid = ? and default_langcode = ?", Langcode_, Uid_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByLangcodeAndStatusAndCreated Get NodeFieldDatas via LangcodeAndStatusAndCreated
func GetNodeFieldDatasByLangcodeAndStatusAndCreated(offset int, limit int, Langcode_ string, Status_ int, Created_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("langcode = ? and status = ? and created = ?", Langcode_, Status_, Created_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByLangcodeAndStatusAndChanged Get NodeFieldDatas via LangcodeAndStatusAndChanged
func GetNodeFieldDatasByLangcodeAndStatusAndChanged(offset int, limit int, Langcode_ string, Status_ int, Changed_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("langcode = ? and status = ? and changed = ?", Langcode_, Status_, Changed_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByLangcodeAndStatusAndPromote Get NodeFieldDatas via LangcodeAndStatusAndPromote
func GetNodeFieldDatasByLangcodeAndStatusAndPromote(offset int, limit int, Langcode_ string, Status_ int, Promote_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("langcode = ? and status = ? and promote = ?", Langcode_, Status_, Promote_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByLangcodeAndStatusAndSticky Get NodeFieldDatas via LangcodeAndStatusAndSticky
func GetNodeFieldDatasByLangcodeAndStatusAndSticky(offset int, limit int, Langcode_ string, Status_ int, Sticky_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("langcode = ? and status = ? and sticky = ?", Langcode_, Status_, Sticky_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByLangcodeAndStatusAndRevisionTranslationAffected Get NodeFieldDatas via LangcodeAndStatusAndRevisionTranslationAffected
func GetNodeFieldDatasByLangcodeAndStatusAndRevisionTranslationAffected(offset int, limit int, Langcode_ string, Status_ int, RevisionTranslationAffected_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("langcode = ? and status = ? and revision_translation_affected = ?", Langcode_, Status_, RevisionTranslationAffected_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByLangcodeAndStatusAndDefaultLangcode Get NodeFieldDatas via LangcodeAndStatusAndDefaultLangcode
func GetNodeFieldDatasByLangcodeAndStatusAndDefaultLangcode(offset int, limit int, Langcode_ string, Status_ int, DefaultLangcode_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("langcode = ? and status = ? and default_langcode = ?", Langcode_, Status_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByLangcodeAndCreatedAndChanged Get NodeFieldDatas via LangcodeAndCreatedAndChanged
func GetNodeFieldDatasByLangcodeAndCreatedAndChanged(offset int, limit int, Langcode_ string, Created_ int, Changed_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("langcode = ? and created = ? and changed = ?", Langcode_, Created_, Changed_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByLangcodeAndCreatedAndPromote Get NodeFieldDatas via LangcodeAndCreatedAndPromote
func GetNodeFieldDatasByLangcodeAndCreatedAndPromote(offset int, limit int, Langcode_ string, Created_ int, Promote_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("langcode = ? and created = ? and promote = ?", Langcode_, Created_, Promote_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByLangcodeAndCreatedAndSticky Get NodeFieldDatas via LangcodeAndCreatedAndSticky
func GetNodeFieldDatasByLangcodeAndCreatedAndSticky(offset int, limit int, Langcode_ string, Created_ int, Sticky_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("langcode = ? and created = ? and sticky = ?", Langcode_, Created_, Sticky_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByLangcodeAndCreatedAndRevisionTranslationAffected Get NodeFieldDatas via LangcodeAndCreatedAndRevisionTranslationAffected
func GetNodeFieldDatasByLangcodeAndCreatedAndRevisionTranslationAffected(offset int, limit int, Langcode_ string, Created_ int, RevisionTranslationAffected_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("langcode = ? and created = ? and revision_translation_affected = ?", Langcode_, Created_, RevisionTranslationAffected_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByLangcodeAndCreatedAndDefaultLangcode Get NodeFieldDatas via LangcodeAndCreatedAndDefaultLangcode
func GetNodeFieldDatasByLangcodeAndCreatedAndDefaultLangcode(offset int, limit int, Langcode_ string, Created_ int, DefaultLangcode_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("langcode = ? and created = ? and default_langcode = ?", Langcode_, Created_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByLangcodeAndChangedAndPromote Get NodeFieldDatas via LangcodeAndChangedAndPromote
func GetNodeFieldDatasByLangcodeAndChangedAndPromote(offset int, limit int, Langcode_ string, Changed_ int, Promote_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("langcode = ? and changed = ? and promote = ?", Langcode_, Changed_, Promote_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByLangcodeAndChangedAndSticky Get NodeFieldDatas via LangcodeAndChangedAndSticky
func GetNodeFieldDatasByLangcodeAndChangedAndSticky(offset int, limit int, Langcode_ string, Changed_ int, Sticky_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("langcode = ? and changed = ? and sticky = ?", Langcode_, Changed_, Sticky_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByLangcodeAndChangedAndRevisionTranslationAffected Get NodeFieldDatas via LangcodeAndChangedAndRevisionTranslationAffected
func GetNodeFieldDatasByLangcodeAndChangedAndRevisionTranslationAffected(offset int, limit int, Langcode_ string, Changed_ int, RevisionTranslationAffected_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("langcode = ? and changed = ? and revision_translation_affected = ?", Langcode_, Changed_, RevisionTranslationAffected_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByLangcodeAndChangedAndDefaultLangcode Get NodeFieldDatas via LangcodeAndChangedAndDefaultLangcode
func GetNodeFieldDatasByLangcodeAndChangedAndDefaultLangcode(offset int, limit int, Langcode_ string, Changed_ int, DefaultLangcode_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("langcode = ? and changed = ? and default_langcode = ?", Langcode_, Changed_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByLangcodeAndPromoteAndSticky Get NodeFieldDatas via LangcodeAndPromoteAndSticky
func GetNodeFieldDatasByLangcodeAndPromoteAndSticky(offset int, limit int, Langcode_ string, Promote_ int, Sticky_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("langcode = ? and promote = ? and sticky = ?", Langcode_, Promote_, Sticky_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByLangcodeAndPromoteAndRevisionTranslationAffected Get NodeFieldDatas via LangcodeAndPromoteAndRevisionTranslationAffected
func GetNodeFieldDatasByLangcodeAndPromoteAndRevisionTranslationAffected(offset int, limit int, Langcode_ string, Promote_ int, RevisionTranslationAffected_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("langcode = ? and promote = ? and revision_translation_affected = ?", Langcode_, Promote_, RevisionTranslationAffected_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByLangcodeAndPromoteAndDefaultLangcode Get NodeFieldDatas via LangcodeAndPromoteAndDefaultLangcode
func GetNodeFieldDatasByLangcodeAndPromoteAndDefaultLangcode(offset int, limit int, Langcode_ string, Promote_ int, DefaultLangcode_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("langcode = ? and promote = ? and default_langcode = ?", Langcode_, Promote_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByLangcodeAndStickyAndRevisionTranslationAffected Get NodeFieldDatas via LangcodeAndStickyAndRevisionTranslationAffected
func GetNodeFieldDatasByLangcodeAndStickyAndRevisionTranslationAffected(offset int, limit int, Langcode_ string, Sticky_ int, RevisionTranslationAffected_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("langcode = ? and sticky = ? and revision_translation_affected = ?", Langcode_, Sticky_, RevisionTranslationAffected_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByLangcodeAndStickyAndDefaultLangcode Get NodeFieldDatas via LangcodeAndStickyAndDefaultLangcode
func GetNodeFieldDatasByLangcodeAndStickyAndDefaultLangcode(offset int, limit int, Langcode_ string, Sticky_ int, DefaultLangcode_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("langcode = ? and sticky = ? and default_langcode = ?", Langcode_, Sticky_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByLangcodeAndRevisionTranslationAffectedAndDefaultLangcode Get NodeFieldDatas via LangcodeAndRevisionTranslationAffectedAndDefaultLangcode
func GetNodeFieldDatasByLangcodeAndRevisionTranslationAffectedAndDefaultLangcode(offset int, limit int, Langcode_ string, RevisionTranslationAffected_ int, DefaultLangcode_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("langcode = ? and revision_translation_affected = ? and default_langcode = ?", Langcode_, RevisionTranslationAffected_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByTitleAndUidAndStatus Get NodeFieldDatas via TitleAndUidAndStatus
func GetNodeFieldDatasByTitleAndUidAndStatus(offset int, limit int, Title_ string, Uid_ int, Status_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("title = ? and uid = ? and status = ?", Title_, Uid_, Status_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByTitleAndUidAndCreated Get NodeFieldDatas via TitleAndUidAndCreated
func GetNodeFieldDatasByTitleAndUidAndCreated(offset int, limit int, Title_ string, Uid_ int, Created_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("title = ? and uid = ? and created = ?", Title_, Uid_, Created_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByTitleAndUidAndChanged Get NodeFieldDatas via TitleAndUidAndChanged
func GetNodeFieldDatasByTitleAndUidAndChanged(offset int, limit int, Title_ string, Uid_ int, Changed_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("title = ? and uid = ? and changed = ?", Title_, Uid_, Changed_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByTitleAndUidAndPromote Get NodeFieldDatas via TitleAndUidAndPromote
func GetNodeFieldDatasByTitleAndUidAndPromote(offset int, limit int, Title_ string, Uid_ int, Promote_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("title = ? and uid = ? and promote = ?", Title_, Uid_, Promote_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByTitleAndUidAndSticky Get NodeFieldDatas via TitleAndUidAndSticky
func GetNodeFieldDatasByTitleAndUidAndSticky(offset int, limit int, Title_ string, Uid_ int, Sticky_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("title = ? and uid = ? and sticky = ?", Title_, Uid_, Sticky_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByTitleAndUidAndRevisionTranslationAffected Get NodeFieldDatas via TitleAndUidAndRevisionTranslationAffected
func GetNodeFieldDatasByTitleAndUidAndRevisionTranslationAffected(offset int, limit int, Title_ string, Uid_ int, RevisionTranslationAffected_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("title = ? and uid = ? and revision_translation_affected = ?", Title_, Uid_, RevisionTranslationAffected_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByTitleAndUidAndDefaultLangcode Get NodeFieldDatas via TitleAndUidAndDefaultLangcode
func GetNodeFieldDatasByTitleAndUidAndDefaultLangcode(offset int, limit int, Title_ string, Uid_ int, DefaultLangcode_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("title = ? and uid = ? and default_langcode = ?", Title_, Uid_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByTitleAndStatusAndCreated Get NodeFieldDatas via TitleAndStatusAndCreated
func GetNodeFieldDatasByTitleAndStatusAndCreated(offset int, limit int, Title_ string, Status_ int, Created_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("title = ? and status = ? and created = ?", Title_, Status_, Created_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByTitleAndStatusAndChanged Get NodeFieldDatas via TitleAndStatusAndChanged
func GetNodeFieldDatasByTitleAndStatusAndChanged(offset int, limit int, Title_ string, Status_ int, Changed_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("title = ? and status = ? and changed = ?", Title_, Status_, Changed_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByTitleAndStatusAndPromote Get NodeFieldDatas via TitleAndStatusAndPromote
func GetNodeFieldDatasByTitleAndStatusAndPromote(offset int, limit int, Title_ string, Status_ int, Promote_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("title = ? and status = ? and promote = ?", Title_, Status_, Promote_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByTitleAndStatusAndSticky Get NodeFieldDatas via TitleAndStatusAndSticky
func GetNodeFieldDatasByTitleAndStatusAndSticky(offset int, limit int, Title_ string, Status_ int, Sticky_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("title = ? and status = ? and sticky = ?", Title_, Status_, Sticky_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByTitleAndStatusAndRevisionTranslationAffected Get NodeFieldDatas via TitleAndStatusAndRevisionTranslationAffected
func GetNodeFieldDatasByTitleAndStatusAndRevisionTranslationAffected(offset int, limit int, Title_ string, Status_ int, RevisionTranslationAffected_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("title = ? and status = ? and revision_translation_affected = ?", Title_, Status_, RevisionTranslationAffected_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByTitleAndStatusAndDefaultLangcode Get NodeFieldDatas via TitleAndStatusAndDefaultLangcode
func GetNodeFieldDatasByTitleAndStatusAndDefaultLangcode(offset int, limit int, Title_ string, Status_ int, DefaultLangcode_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("title = ? and status = ? and default_langcode = ?", Title_, Status_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByTitleAndCreatedAndChanged Get NodeFieldDatas via TitleAndCreatedAndChanged
func GetNodeFieldDatasByTitleAndCreatedAndChanged(offset int, limit int, Title_ string, Created_ int, Changed_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("title = ? and created = ? and changed = ?", Title_, Created_, Changed_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByTitleAndCreatedAndPromote Get NodeFieldDatas via TitleAndCreatedAndPromote
func GetNodeFieldDatasByTitleAndCreatedAndPromote(offset int, limit int, Title_ string, Created_ int, Promote_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("title = ? and created = ? and promote = ?", Title_, Created_, Promote_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByTitleAndCreatedAndSticky Get NodeFieldDatas via TitleAndCreatedAndSticky
func GetNodeFieldDatasByTitleAndCreatedAndSticky(offset int, limit int, Title_ string, Created_ int, Sticky_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("title = ? and created = ? and sticky = ?", Title_, Created_, Sticky_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByTitleAndCreatedAndRevisionTranslationAffected Get NodeFieldDatas via TitleAndCreatedAndRevisionTranslationAffected
func GetNodeFieldDatasByTitleAndCreatedAndRevisionTranslationAffected(offset int, limit int, Title_ string, Created_ int, RevisionTranslationAffected_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("title = ? and created = ? and revision_translation_affected = ?", Title_, Created_, RevisionTranslationAffected_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByTitleAndCreatedAndDefaultLangcode Get NodeFieldDatas via TitleAndCreatedAndDefaultLangcode
func GetNodeFieldDatasByTitleAndCreatedAndDefaultLangcode(offset int, limit int, Title_ string, Created_ int, DefaultLangcode_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("title = ? and created = ? and default_langcode = ?", Title_, Created_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByTitleAndChangedAndPromote Get NodeFieldDatas via TitleAndChangedAndPromote
func GetNodeFieldDatasByTitleAndChangedAndPromote(offset int, limit int, Title_ string, Changed_ int, Promote_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("title = ? and changed = ? and promote = ?", Title_, Changed_, Promote_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByTitleAndChangedAndSticky Get NodeFieldDatas via TitleAndChangedAndSticky
func GetNodeFieldDatasByTitleAndChangedAndSticky(offset int, limit int, Title_ string, Changed_ int, Sticky_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("title = ? and changed = ? and sticky = ?", Title_, Changed_, Sticky_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByTitleAndChangedAndRevisionTranslationAffected Get NodeFieldDatas via TitleAndChangedAndRevisionTranslationAffected
func GetNodeFieldDatasByTitleAndChangedAndRevisionTranslationAffected(offset int, limit int, Title_ string, Changed_ int, RevisionTranslationAffected_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("title = ? and changed = ? and revision_translation_affected = ?", Title_, Changed_, RevisionTranslationAffected_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByTitleAndChangedAndDefaultLangcode Get NodeFieldDatas via TitleAndChangedAndDefaultLangcode
func GetNodeFieldDatasByTitleAndChangedAndDefaultLangcode(offset int, limit int, Title_ string, Changed_ int, DefaultLangcode_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("title = ? and changed = ? and default_langcode = ?", Title_, Changed_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByTitleAndPromoteAndSticky Get NodeFieldDatas via TitleAndPromoteAndSticky
func GetNodeFieldDatasByTitleAndPromoteAndSticky(offset int, limit int, Title_ string, Promote_ int, Sticky_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("title = ? and promote = ? and sticky = ?", Title_, Promote_, Sticky_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByTitleAndPromoteAndRevisionTranslationAffected Get NodeFieldDatas via TitleAndPromoteAndRevisionTranslationAffected
func GetNodeFieldDatasByTitleAndPromoteAndRevisionTranslationAffected(offset int, limit int, Title_ string, Promote_ int, RevisionTranslationAffected_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("title = ? and promote = ? and revision_translation_affected = ?", Title_, Promote_, RevisionTranslationAffected_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByTitleAndPromoteAndDefaultLangcode Get NodeFieldDatas via TitleAndPromoteAndDefaultLangcode
func GetNodeFieldDatasByTitleAndPromoteAndDefaultLangcode(offset int, limit int, Title_ string, Promote_ int, DefaultLangcode_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("title = ? and promote = ? and default_langcode = ?", Title_, Promote_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByTitleAndStickyAndRevisionTranslationAffected Get NodeFieldDatas via TitleAndStickyAndRevisionTranslationAffected
func GetNodeFieldDatasByTitleAndStickyAndRevisionTranslationAffected(offset int, limit int, Title_ string, Sticky_ int, RevisionTranslationAffected_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("title = ? and sticky = ? and revision_translation_affected = ?", Title_, Sticky_, RevisionTranslationAffected_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByTitleAndStickyAndDefaultLangcode Get NodeFieldDatas via TitleAndStickyAndDefaultLangcode
func GetNodeFieldDatasByTitleAndStickyAndDefaultLangcode(offset int, limit int, Title_ string, Sticky_ int, DefaultLangcode_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("title = ? and sticky = ? and default_langcode = ?", Title_, Sticky_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByTitleAndRevisionTranslationAffectedAndDefaultLangcode Get NodeFieldDatas via TitleAndRevisionTranslationAffectedAndDefaultLangcode
func GetNodeFieldDatasByTitleAndRevisionTranslationAffectedAndDefaultLangcode(offset int, limit int, Title_ string, RevisionTranslationAffected_ int, DefaultLangcode_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("title = ? and revision_translation_affected = ? and default_langcode = ?", Title_, RevisionTranslationAffected_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByUidAndStatusAndCreated Get NodeFieldDatas via UidAndStatusAndCreated
func GetNodeFieldDatasByUidAndStatusAndCreated(offset int, limit int, Uid_ int, Status_ int, Created_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("uid = ? and status = ? and created = ?", Uid_, Status_, Created_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByUidAndStatusAndChanged Get NodeFieldDatas via UidAndStatusAndChanged
func GetNodeFieldDatasByUidAndStatusAndChanged(offset int, limit int, Uid_ int, Status_ int, Changed_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("uid = ? and status = ? and changed = ?", Uid_, Status_, Changed_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByUidAndStatusAndPromote Get NodeFieldDatas via UidAndStatusAndPromote
func GetNodeFieldDatasByUidAndStatusAndPromote(offset int, limit int, Uid_ int, Status_ int, Promote_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("uid = ? and status = ? and promote = ?", Uid_, Status_, Promote_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByUidAndStatusAndSticky Get NodeFieldDatas via UidAndStatusAndSticky
func GetNodeFieldDatasByUidAndStatusAndSticky(offset int, limit int, Uid_ int, Status_ int, Sticky_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("uid = ? and status = ? and sticky = ?", Uid_, Status_, Sticky_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByUidAndStatusAndRevisionTranslationAffected Get NodeFieldDatas via UidAndStatusAndRevisionTranslationAffected
func GetNodeFieldDatasByUidAndStatusAndRevisionTranslationAffected(offset int, limit int, Uid_ int, Status_ int, RevisionTranslationAffected_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("uid = ? and status = ? and revision_translation_affected = ?", Uid_, Status_, RevisionTranslationAffected_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByUidAndStatusAndDefaultLangcode Get NodeFieldDatas via UidAndStatusAndDefaultLangcode
func GetNodeFieldDatasByUidAndStatusAndDefaultLangcode(offset int, limit int, Uid_ int, Status_ int, DefaultLangcode_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("uid = ? and status = ? and default_langcode = ?", Uid_, Status_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByUidAndCreatedAndChanged Get NodeFieldDatas via UidAndCreatedAndChanged
func GetNodeFieldDatasByUidAndCreatedAndChanged(offset int, limit int, Uid_ int, Created_ int, Changed_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("uid = ? and created = ? and changed = ?", Uid_, Created_, Changed_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByUidAndCreatedAndPromote Get NodeFieldDatas via UidAndCreatedAndPromote
func GetNodeFieldDatasByUidAndCreatedAndPromote(offset int, limit int, Uid_ int, Created_ int, Promote_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("uid = ? and created = ? and promote = ?", Uid_, Created_, Promote_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByUidAndCreatedAndSticky Get NodeFieldDatas via UidAndCreatedAndSticky
func GetNodeFieldDatasByUidAndCreatedAndSticky(offset int, limit int, Uid_ int, Created_ int, Sticky_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("uid = ? and created = ? and sticky = ?", Uid_, Created_, Sticky_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByUidAndCreatedAndRevisionTranslationAffected Get NodeFieldDatas via UidAndCreatedAndRevisionTranslationAffected
func GetNodeFieldDatasByUidAndCreatedAndRevisionTranslationAffected(offset int, limit int, Uid_ int, Created_ int, RevisionTranslationAffected_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("uid = ? and created = ? and revision_translation_affected = ?", Uid_, Created_, RevisionTranslationAffected_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByUidAndCreatedAndDefaultLangcode Get NodeFieldDatas via UidAndCreatedAndDefaultLangcode
func GetNodeFieldDatasByUidAndCreatedAndDefaultLangcode(offset int, limit int, Uid_ int, Created_ int, DefaultLangcode_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("uid = ? and created = ? and default_langcode = ?", Uid_, Created_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByUidAndChangedAndPromote Get NodeFieldDatas via UidAndChangedAndPromote
func GetNodeFieldDatasByUidAndChangedAndPromote(offset int, limit int, Uid_ int, Changed_ int, Promote_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("uid = ? and changed = ? and promote = ?", Uid_, Changed_, Promote_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByUidAndChangedAndSticky Get NodeFieldDatas via UidAndChangedAndSticky
func GetNodeFieldDatasByUidAndChangedAndSticky(offset int, limit int, Uid_ int, Changed_ int, Sticky_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("uid = ? and changed = ? and sticky = ?", Uid_, Changed_, Sticky_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByUidAndChangedAndRevisionTranslationAffected Get NodeFieldDatas via UidAndChangedAndRevisionTranslationAffected
func GetNodeFieldDatasByUidAndChangedAndRevisionTranslationAffected(offset int, limit int, Uid_ int, Changed_ int, RevisionTranslationAffected_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("uid = ? and changed = ? and revision_translation_affected = ?", Uid_, Changed_, RevisionTranslationAffected_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByUidAndChangedAndDefaultLangcode Get NodeFieldDatas via UidAndChangedAndDefaultLangcode
func GetNodeFieldDatasByUidAndChangedAndDefaultLangcode(offset int, limit int, Uid_ int, Changed_ int, DefaultLangcode_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("uid = ? and changed = ? and default_langcode = ?", Uid_, Changed_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByUidAndPromoteAndSticky Get NodeFieldDatas via UidAndPromoteAndSticky
func GetNodeFieldDatasByUidAndPromoteAndSticky(offset int, limit int, Uid_ int, Promote_ int, Sticky_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("uid = ? and promote = ? and sticky = ?", Uid_, Promote_, Sticky_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByUidAndPromoteAndRevisionTranslationAffected Get NodeFieldDatas via UidAndPromoteAndRevisionTranslationAffected
func GetNodeFieldDatasByUidAndPromoteAndRevisionTranslationAffected(offset int, limit int, Uid_ int, Promote_ int, RevisionTranslationAffected_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("uid = ? and promote = ? and revision_translation_affected = ?", Uid_, Promote_, RevisionTranslationAffected_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByUidAndPromoteAndDefaultLangcode Get NodeFieldDatas via UidAndPromoteAndDefaultLangcode
func GetNodeFieldDatasByUidAndPromoteAndDefaultLangcode(offset int, limit int, Uid_ int, Promote_ int, DefaultLangcode_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("uid = ? and promote = ? and default_langcode = ?", Uid_, Promote_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByUidAndStickyAndRevisionTranslationAffected Get NodeFieldDatas via UidAndStickyAndRevisionTranslationAffected
func GetNodeFieldDatasByUidAndStickyAndRevisionTranslationAffected(offset int, limit int, Uid_ int, Sticky_ int, RevisionTranslationAffected_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("uid = ? and sticky = ? and revision_translation_affected = ?", Uid_, Sticky_, RevisionTranslationAffected_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByUidAndStickyAndDefaultLangcode Get NodeFieldDatas via UidAndStickyAndDefaultLangcode
func GetNodeFieldDatasByUidAndStickyAndDefaultLangcode(offset int, limit int, Uid_ int, Sticky_ int, DefaultLangcode_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("uid = ? and sticky = ? and default_langcode = ?", Uid_, Sticky_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByUidAndRevisionTranslationAffectedAndDefaultLangcode Get NodeFieldDatas via UidAndRevisionTranslationAffectedAndDefaultLangcode
func GetNodeFieldDatasByUidAndRevisionTranslationAffectedAndDefaultLangcode(offset int, limit int, Uid_ int, RevisionTranslationAffected_ int, DefaultLangcode_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("uid = ? and revision_translation_affected = ? and default_langcode = ?", Uid_, RevisionTranslationAffected_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByStatusAndCreatedAndChanged Get NodeFieldDatas via StatusAndCreatedAndChanged
func GetNodeFieldDatasByStatusAndCreatedAndChanged(offset int, limit int, Status_ int, Created_ int, Changed_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("status = ? and created = ? and changed = ?", Status_, Created_, Changed_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByStatusAndCreatedAndPromote Get NodeFieldDatas via StatusAndCreatedAndPromote
func GetNodeFieldDatasByStatusAndCreatedAndPromote(offset int, limit int, Status_ int, Created_ int, Promote_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("status = ? and created = ? and promote = ?", Status_, Created_, Promote_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByStatusAndCreatedAndSticky Get NodeFieldDatas via StatusAndCreatedAndSticky
func GetNodeFieldDatasByStatusAndCreatedAndSticky(offset int, limit int, Status_ int, Created_ int, Sticky_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("status = ? and created = ? and sticky = ?", Status_, Created_, Sticky_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByStatusAndCreatedAndRevisionTranslationAffected Get NodeFieldDatas via StatusAndCreatedAndRevisionTranslationAffected
func GetNodeFieldDatasByStatusAndCreatedAndRevisionTranslationAffected(offset int, limit int, Status_ int, Created_ int, RevisionTranslationAffected_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("status = ? and created = ? and revision_translation_affected = ?", Status_, Created_, RevisionTranslationAffected_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByStatusAndCreatedAndDefaultLangcode Get NodeFieldDatas via StatusAndCreatedAndDefaultLangcode
func GetNodeFieldDatasByStatusAndCreatedAndDefaultLangcode(offset int, limit int, Status_ int, Created_ int, DefaultLangcode_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("status = ? and created = ? and default_langcode = ?", Status_, Created_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByStatusAndChangedAndPromote Get NodeFieldDatas via StatusAndChangedAndPromote
func GetNodeFieldDatasByStatusAndChangedAndPromote(offset int, limit int, Status_ int, Changed_ int, Promote_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("status = ? and changed = ? and promote = ?", Status_, Changed_, Promote_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByStatusAndChangedAndSticky Get NodeFieldDatas via StatusAndChangedAndSticky
func GetNodeFieldDatasByStatusAndChangedAndSticky(offset int, limit int, Status_ int, Changed_ int, Sticky_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("status = ? and changed = ? and sticky = ?", Status_, Changed_, Sticky_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByStatusAndChangedAndRevisionTranslationAffected Get NodeFieldDatas via StatusAndChangedAndRevisionTranslationAffected
func GetNodeFieldDatasByStatusAndChangedAndRevisionTranslationAffected(offset int, limit int, Status_ int, Changed_ int, RevisionTranslationAffected_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("status = ? and changed = ? and revision_translation_affected = ?", Status_, Changed_, RevisionTranslationAffected_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByStatusAndChangedAndDefaultLangcode Get NodeFieldDatas via StatusAndChangedAndDefaultLangcode
func GetNodeFieldDatasByStatusAndChangedAndDefaultLangcode(offset int, limit int, Status_ int, Changed_ int, DefaultLangcode_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("status = ? and changed = ? and default_langcode = ?", Status_, Changed_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByStatusAndPromoteAndSticky Get NodeFieldDatas via StatusAndPromoteAndSticky
func GetNodeFieldDatasByStatusAndPromoteAndSticky(offset int, limit int, Status_ int, Promote_ int, Sticky_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("status = ? and promote = ? and sticky = ?", Status_, Promote_, Sticky_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByStatusAndPromoteAndRevisionTranslationAffected Get NodeFieldDatas via StatusAndPromoteAndRevisionTranslationAffected
func GetNodeFieldDatasByStatusAndPromoteAndRevisionTranslationAffected(offset int, limit int, Status_ int, Promote_ int, RevisionTranslationAffected_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("status = ? and promote = ? and revision_translation_affected = ?", Status_, Promote_, RevisionTranslationAffected_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByStatusAndPromoteAndDefaultLangcode Get NodeFieldDatas via StatusAndPromoteAndDefaultLangcode
func GetNodeFieldDatasByStatusAndPromoteAndDefaultLangcode(offset int, limit int, Status_ int, Promote_ int, DefaultLangcode_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("status = ? and promote = ? and default_langcode = ?", Status_, Promote_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByStatusAndStickyAndRevisionTranslationAffected Get NodeFieldDatas via StatusAndStickyAndRevisionTranslationAffected
func GetNodeFieldDatasByStatusAndStickyAndRevisionTranslationAffected(offset int, limit int, Status_ int, Sticky_ int, RevisionTranslationAffected_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("status = ? and sticky = ? and revision_translation_affected = ?", Status_, Sticky_, RevisionTranslationAffected_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByStatusAndStickyAndDefaultLangcode Get NodeFieldDatas via StatusAndStickyAndDefaultLangcode
func GetNodeFieldDatasByStatusAndStickyAndDefaultLangcode(offset int, limit int, Status_ int, Sticky_ int, DefaultLangcode_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("status = ? and sticky = ? and default_langcode = ?", Status_, Sticky_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByStatusAndRevisionTranslationAffectedAndDefaultLangcode Get NodeFieldDatas via StatusAndRevisionTranslationAffectedAndDefaultLangcode
func GetNodeFieldDatasByStatusAndRevisionTranslationAffectedAndDefaultLangcode(offset int, limit int, Status_ int, RevisionTranslationAffected_ int, DefaultLangcode_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("status = ? and revision_translation_affected = ? and default_langcode = ?", Status_, RevisionTranslationAffected_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByCreatedAndChangedAndPromote Get NodeFieldDatas via CreatedAndChangedAndPromote
func GetNodeFieldDatasByCreatedAndChangedAndPromote(offset int, limit int, Created_ int, Changed_ int, Promote_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("created = ? and changed = ? and promote = ?", Created_, Changed_, Promote_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByCreatedAndChangedAndSticky Get NodeFieldDatas via CreatedAndChangedAndSticky
func GetNodeFieldDatasByCreatedAndChangedAndSticky(offset int, limit int, Created_ int, Changed_ int, Sticky_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("created = ? and changed = ? and sticky = ?", Created_, Changed_, Sticky_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByCreatedAndChangedAndRevisionTranslationAffected Get NodeFieldDatas via CreatedAndChangedAndRevisionTranslationAffected
func GetNodeFieldDatasByCreatedAndChangedAndRevisionTranslationAffected(offset int, limit int, Created_ int, Changed_ int, RevisionTranslationAffected_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("created = ? and changed = ? and revision_translation_affected = ?", Created_, Changed_, RevisionTranslationAffected_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByCreatedAndChangedAndDefaultLangcode Get NodeFieldDatas via CreatedAndChangedAndDefaultLangcode
func GetNodeFieldDatasByCreatedAndChangedAndDefaultLangcode(offset int, limit int, Created_ int, Changed_ int, DefaultLangcode_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("created = ? and changed = ? and default_langcode = ?", Created_, Changed_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByCreatedAndPromoteAndSticky Get NodeFieldDatas via CreatedAndPromoteAndSticky
func GetNodeFieldDatasByCreatedAndPromoteAndSticky(offset int, limit int, Created_ int, Promote_ int, Sticky_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("created = ? and promote = ? and sticky = ?", Created_, Promote_, Sticky_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByCreatedAndPromoteAndRevisionTranslationAffected Get NodeFieldDatas via CreatedAndPromoteAndRevisionTranslationAffected
func GetNodeFieldDatasByCreatedAndPromoteAndRevisionTranslationAffected(offset int, limit int, Created_ int, Promote_ int, RevisionTranslationAffected_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("created = ? and promote = ? and revision_translation_affected = ?", Created_, Promote_, RevisionTranslationAffected_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByCreatedAndPromoteAndDefaultLangcode Get NodeFieldDatas via CreatedAndPromoteAndDefaultLangcode
func GetNodeFieldDatasByCreatedAndPromoteAndDefaultLangcode(offset int, limit int, Created_ int, Promote_ int, DefaultLangcode_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("created = ? and promote = ? and default_langcode = ?", Created_, Promote_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByCreatedAndStickyAndRevisionTranslationAffected Get NodeFieldDatas via CreatedAndStickyAndRevisionTranslationAffected
func GetNodeFieldDatasByCreatedAndStickyAndRevisionTranslationAffected(offset int, limit int, Created_ int, Sticky_ int, RevisionTranslationAffected_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("created = ? and sticky = ? and revision_translation_affected = ?", Created_, Sticky_, RevisionTranslationAffected_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByCreatedAndStickyAndDefaultLangcode Get NodeFieldDatas via CreatedAndStickyAndDefaultLangcode
func GetNodeFieldDatasByCreatedAndStickyAndDefaultLangcode(offset int, limit int, Created_ int, Sticky_ int, DefaultLangcode_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("created = ? and sticky = ? and default_langcode = ?", Created_, Sticky_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByCreatedAndRevisionTranslationAffectedAndDefaultLangcode Get NodeFieldDatas via CreatedAndRevisionTranslationAffectedAndDefaultLangcode
func GetNodeFieldDatasByCreatedAndRevisionTranslationAffectedAndDefaultLangcode(offset int, limit int, Created_ int, RevisionTranslationAffected_ int, DefaultLangcode_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("created = ? and revision_translation_affected = ? and default_langcode = ?", Created_, RevisionTranslationAffected_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByChangedAndPromoteAndSticky Get NodeFieldDatas via ChangedAndPromoteAndSticky
func GetNodeFieldDatasByChangedAndPromoteAndSticky(offset int, limit int, Changed_ int, Promote_ int, Sticky_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("changed = ? and promote = ? and sticky = ?", Changed_, Promote_, Sticky_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByChangedAndPromoteAndRevisionTranslationAffected Get NodeFieldDatas via ChangedAndPromoteAndRevisionTranslationAffected
func GetNodeFieldDatasByChangedAndPromoteAndRevisionTranslationAffected(offset int, limit int, Changed_ int, Promote_ int, RevisionTranslationAffected_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("changed = ? and promote = ? and revision_translation_affected = ?", Changed_, Promote_, RevisionTranslationAffected_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByChangedAndPromoteAndDefaultLangcode Get NodeFieldDatas via ChangedAndPromoteAndDefaultLangcode
func GetNodeFieldDatasByChangedAndPromoteAndDefaultLangcode(offset int, limit int, Changed_ int, Promote_ int, DefaultLangcode_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("changed = ? and promote = ? and default_langcode = ?", Changed_, Promote_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByChangedAndStickyAndRevisionTranslationAffected Get NodeFieldDatas via ChangedAndStickyAndRevisionTranslationAffected
func GetNodeFieldDatasByChangedAndStickyAndRevisionTranslationAffected(offset int, limit int, Changed_ int, Sticky_ int, RevisionTranslationAffected_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("changed = ? and sticky = ? and revision_translation_affected = ?", Changed_, Sticky_, RevisionTranslationAffected_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByChangedAndStickyAndDefaultLangcode Get NodeFieldDatas via ChangedAndStickyAndDefaultLangcode
func GetNodeFieldDatasByChangedAndStickyAndDefaultLangcode(offset int, limit int, Changed_ int, Sticky_ int, DefaultLangcode_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("changed = ? and sticky = ? and default_langcode = ?", Changed_, Sticky_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByChangedAndRevisionTranslationAffectedAndDefaultLangcode Get NodeFieldDatas via ChangedAndRevisionTranslationAffectedAndDefaultLangcode
func GetNodeFieldDatasByChangedAndRevisionTranslationAffectedAndDefaultLangcode(offset int, limit int, Changed_ int, RevisionTranslationAffected_ int, DefaultLangcode_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("changed = ? and revision_translation_affected = ? and default_langcode = ?", Changed_, RevisionTranslationAffected_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByPromoteAndStickyAndRevisionTranslationAffected Get NodeFieldDatas via PromoteAndStickyAndRevisionTranslationAffected
func GetNodeFieldDatasByPromoteAndStickyAndRevisionTranslationAffected(offset int, limit int, Promote_ int, Sticky_ int, RevisionTranslationAffected_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("promote = ? and sticky = ? and revision_translation_affected = ?", Promote_, Sticky_, RevisionTranslationAffected_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByPromoteAndStickyAndDefaultLangcode Get NodeFieldDatas via PromoteAndStickyAndDefaultLangcode
func GetNodeFieldDatasByPromoteAndStickyAndDefaultLangcode(offset int, limit int, Promote_ int, Sticky_ int, DefaultLangcode_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("promote = ? and sticky = ? and default_langcode = ?", Promote_, Sticky_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByPromoteAndRevisionTranslationAffectedAndDefaultLangcode Get NodeFieldDatas via PromoteAndRevisionTranslationAffectedAndDefaultLangcode
func GetNodeFieldDatasByPromoteAndRevisionTranslationAffectedAndDefaultLangcode(offset int, limit int, Promote_ int, RevisionTranslationAffected_ int, DefaultLangcode_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("promote = ? and revision_translation_affected = ? and default_langcode = ?", Promote_, RevisionTranslationAffected_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByStickyAndRevisionTranslationAffectedAndDefaultLangcode Get NodeFieldDatas via StickyAndRevisionTranslationAffectedAndDefaultLangcode
func GetNodeFieldDatasByStickyAndRevisionTranslationAffectedAndDefaultLangcode(offset int, limit int, Sticky_ int, RevisionTranslationAffected_ int, DefaultLangcode_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("sticky = ? and revision_translation_affected = ? and default_langcode = ?", Sticky_, RevisionTranslationAffected_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByNidAndVid Get NodeFieldDatas via NidAndVid
func GetNodeFieldDatasByNidAndVid(offset int, limit int, Nid_ int64, Vid_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("nid = ? and vid = ?", Nid_, Vid_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByNidAndType Get NodeFieldDatas via NidAndType
func GetNodeFieldDatasByNidAndType(offset int, limit int, Nid_ int64, Type_ string) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("nid = ? and type = ?", Nid_, Type_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByNidAndLangcode Get NodeFieldDatas via NidAndLangcode
func GetNodeFieldDatasByNidAndLangcode(offset int, limit int, Nid_ int64, Langcode_ string) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("nid = ? and langcode = ?", Nid_, Langcode_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByNidAndTitle Get NodeFieldDatas via NidAndTitle
func GetNodeFieldDatasByNidAndTitle(offset int, limit int, Nid_ int64, Title_ string) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("nid = ? and title = ?", Nid_, Title_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByNidAndUid Get NodeFieldDatas via NidAndUid
func GetNodeFieldDatasByNidAndUid(offset int, limit int, Nid_ int64, Uid_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("nid = ? and uid = ?", Nid_, Uid_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByNidAndStatus Get NodeFieldDatas via NidAndStatus
func GetNodeFieldDatasByNidAndStatus(offset int, limit int, Nid_ int64, Status_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("nid = ? and status = ?", Nid_, Status_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByNidAndCreated Get NodeFieldDatas via NidAndCreated
func GetNodeFieldDatasByNidAndCreated(offset int, limit int, Nid_ int64, Created_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("nid = ? and created = ?", Nid_, Created_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByNidAndChanged Get NodeFieldDatas via NidAndChanged
func GetNodeFieldDatasByNidAndChanged(offset int, limit int, Nid_ int64, Changed_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("nid = ? and changed = ?", Nid_, Changed_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByNidAndPromote Get NodeFieldDatas via NidAndPromote
func GetNodeFieldDatasByNidAndPromote(offset int, limit int, Nid_ int64, Promote_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("nid = ? and promote = ?", Nid_, Promote_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByNidAndSticky Get NodeFieldDatas via NidAndSticky
func GetNodeFieldDatasByNidAndSticky(offset int, limit int, Nid_ int64, Sticky_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("nid = ? and sticky = ?", Nid_, Sticky_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByNidAndRevisionTranslationAffected Get NodeFieldDatas via NidAndRevisionTranslationAffected
func GetNodeFieldDatasByNidAndRevisionTranslationAffected(offset int, limit int, Nid_ int64, RevisionTranslationAffected_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("nid = ? and revision_translation_affected = ?", Nid_, RevisionTranslationAffected_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByNidAndDefaultLangcode Get NodeFieldDatas via NidAndDefaultLangcode
func GetNodeFieldDatasByNidAndDefaultLangcode(offset int, limit int, Nid_ int64, DefaultLangcode_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("nid = ? and default_langcode = ?", Nid_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByVidAndType Get NodeFieldDatas via VidAndType
func GetNodeFieldDatasByVidAndType(offset int, limit int, Vid_ int, Type_ string) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("vid = ? and type = ?", Vid_, Type_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByVidAndLangcode Get NodeFieldDatas via VidAndLangcode
func GetNodeFieldDatasByVidAndLangcode(offset int, limit int, Vid_ int, Langcode_ string) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("vid = ? and langcode = ?", Vid_, Langcode_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByVidAndTitle Get NodeFieldDatas via VidAndTitle
func GetNodeFieldDatasByVidAndTitle(offset int, limit int, Vid_ int, Title_ string) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("vid = ? and title = ?", Vid_, Title_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByVidAndUid Get NodeFieldDatas via VidAndUid
func GetNodeFieldDatasByVidAndUid(offset int, limit int, Vid_ int, Uid_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("vid = ? and uid = ?", Vid_, Uid_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByVidAndStatus Get NodeFieldDatas via VidAndStatus
func GetNodeFieldDatasByVidAndStatus(offset int, limit int, Vid_ int, Status_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("vid = ? and status = ?", Vid_, Status_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByVidAndCreated Get NodeFieldDatas via VidAndCreated
func GetNodeFieldDatasByVidAndCreated(offset int, limit int, Vid_ int, Created_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("vid = ? and created = ?", Vid_, Created_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByVidAndChanged Get NodeFieldDatas via VidAndChanged
func GetNodeFieldDatasByVidAndChanged(offset int, limit int, Vid_ int, Changed_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("vid = ? and changed = ?", Vid_, Changed_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByVidAndPromote Get NodeFieldDatas via VidAndPromote
func GetNodeFieldDatasByVidAndPromote(offset int, limit int, Vid_ int, Promote_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("vid = ? and promote = ?", Vid_, Promote_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByVidAndSticky Get NodeFieldDatas via VidAndSticky
func GetNodeFieldDatasByVidAndSticky(offset int, limit int, Vid_ int, Sticky_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("vid = ? and sticky = ?", Vid_, Sticky_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByVidAndRevisionTranslationAffected Get NodeFieldDatas via VidAndRevisionTranslationAffected
func GetNodeFieldDatasByVidAndRevisionTranslationAffected(offset int, limit int, Vid_ int, RevisionTranslationAffected_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("vid = ? and revision_translation_affected = ?", Vid_, RevisionTranslationAffected_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByVidAndDefaultLangcode Get NodeFieldDatas via VidAndDefaultLangcode
func GetNodeFieldDatasByVidAndDefaultLangcode(offset int, limit int, Vid_ int, DefaultLangcode_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("vid = ? and default_langcode = ?", Vid_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByTypeAndLangcode Get NodeFieldDatas via TypeAndLangcode
func GetNodeFieldDatasByTypeAndLangcode(offset int, limit int, Type_ string, Langcode_ string) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("type = ? and langcode = ?", Type_, Langcode_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByTypeAndTitle Get NodeFieldDatas via TypeAndTitle
func GetNodeFieldDatasByTypeAndTitle(offset int, limit int, Type_ string, Title_ string) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("type = ? and title = ?", Type_, Title_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByTypeAndUid Get NodeFieldDatas via TypeAndUid
func GetNodeFieldDatasByTypeAndUid(offset int, limit int, Type_ string, Uid_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("type = ? and uid = ?", Type_, Uid_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByTypeAndStatus Get NodeFieldDatas via TypeAndStatus
func GetNodeFieldDatasByTypeAndStatus(offset int, limit int, Type_ string, Status_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("type = ? and status = ?", Type_, Status_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByTypeAndCreated Get NodeFieldDatas via TypeAndCreated
func GetNodeFieldDatasByTypeAndCreated(offset int, limit int, Type_ string, Created_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("type = ? and created = ?", Type_, Created_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByTypeAndChanged Get NodeFieldDatas via TypeAndChanged
func GetNodeFieldDatasByTypeAndChanged(offset int, limit int, Type_ string, Changed_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("type = ? and changed = ?", Type_, Changed_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByTypeAndPromote Get NodeFieldDatas via TypeAndPromote
func GetNodeFieldDatasByTypeAndPromote(offset int, limit int, Type_ string, Promote_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("type = ? and promote = ?", Type_, Promote_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByTypeAndSticky Get NodeFieldDatas via TypeAndSticky
func GetNodeFieldDatasByTypeAndSticky(offset int, limit int, Type_ string, Sticky_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("type = ? and sticky = ?", Type_, Sticky_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByTypeAndRevisionTranslationAffected Get NodeFieldDatas via TypeAndRevisionTranslationAffected
func GetNodeFieldDatasByTypeAndRevisionTranslationAffected(offset int, limit int, Type_ string, RevisionTranslationAffected_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("type = ? and revision_translation_affected = ?", Type_, RevisionTranslationAffected_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByTypeAndDefaultLangcode Get NodeFieldDatas via TypeAndDefaultLangcode
func GetNodeFieldDatasByTypeAndDefaultLangcode(offset int, limit int, Type_ string, DefaultLangcode_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("type = ? and default_langcode = ?", Type_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByLangcodeAndTitle Get NodeFieldDatas via LangcodeAndTitle
func GetNodeFieldDatasByLangcodeAndTitle(offset int, limit int, Langcode_ string, Title_ string) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("langcode = ? and title = ?", Langcode_, Title_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByLangcodeAndUid Get NodeFieldDatas via LangcodeAndUid
func GetNodeFieldDatasByLangcodeAndUid(offset int, limit int, Langcode_ string, Uid_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("langcode = ? and uid = ?", Langcode_, Uid_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByLangcodeAndStatus Get NodeFieldDatas via LangcodeAndStatus
func GetNodeFieldDatasByLangcodeAndStatus(offset int, limit int, Langcode_ string, Status_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("langcode = ? and status = ?", Langcode_, Status_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByLangcodeAndCreated Get NodeFieldDatas via LangcodeAndCreated
func GetNodeFieldDatasByLangcodeAndCreated(offset int, limit int, Langcode_ string, Created_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("langcode = ? and created = ?", Langcode_, Created_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByLangcodeAndChanged Get NodeFieldDatas via LangcodeAndChanged
func GetNodeFieldDatasByLangcodeAndChanged(offset int, limit int, Langcode_ string, Changed_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("langcode = ? and changed = ?", Langcode_, Changed_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByLangcodeAndPromote Get NodeFieldDatas via LangcodeAndPromote
func GetNodeFieldDatasByLangcodeAndPromote(offset int, limit int, Langcode_ string, Promote_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("langcode = ? and promote = ?", Langcode_, Promote_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByLangcodeAndSticky Get NodeFieldDatas via LangcodeAndSticky
func GetNodeFieldDatasByLangcodeAndSticky(offset int, limit int, Langcode_ string, Sticky_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("langcode = ? and sticky = ?", Langcode_, Sticky_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByLangcodeAndRevisionTranslationAffected Get NodeFieldDatas via LangcodeAndRevisionTranslationAffected
func GetNodeFieldDatasByLangcodeAndRevisionTranslationAffected(offset int, limit int, Langcode_ string, RevisionTranslationAffected_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("langcode = ? and revision_translation_affected = ?", Langcode_, RevisionTranslationAffected_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByLangcodeAndDefaultLangcode Get NodeFieldDatas via LangcodeAndDefaultLangcode
func GetNodeFieldDatasByLangcodeAndDefaultLangcode(offset int, limit int, Langcode_ string, DefaultLangcode_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("langcode = ? and default_langcode = ?", Langcode_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByTitleAndUid Get NodeFieldDatas via TitleAndUid
func GetNodeFieldDatasByTitleAndUid(offset int, limit int, Title_ string, Uid_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("title = ? and uid = ?", Title_, Uid_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByTitleAndStatus Get NodeFieldDatas via TitleAndStatus
func GetNodeFieldDatasByTitleAndStatus(offset int, limit int, Title_ string, Status_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("title = ? and status = ?", Title_, Status_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByTitleAndCreated Get NodeFieldDatas via TitleAndCreated
func GetNodeFieldDatasByTitleAndCreated(offset int, limit int, Title_ string, Created_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("title = ? and created = ?", Title_, Created_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByTitleAndChanged Get NodeFieldDatas via TitleAndChanged
func GetNodeFieldDatasByTitleAndChanged(offset int, limit int, Title_ string, Changed_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("title = ? and changed = ?", Title_, Changed_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByTitleAndPromote Get NodeFieldDatas via TitleAndPromote
func GetNodeFieldDatasByTitleAndPromote(offset int, limit int, Title_ string, Promote_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("title = ? and promote = ?", Title_, Promote_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByTitleAndSticky Get NodeFieldDatas via TitleAndSticky
func GetNodeFieldDatasByTitleAndSticky(offset int, limit int, Title_ string, Sticky_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("title = ? and sticky = ?", Title_, Sticky_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByTitleAndRevisionTranslationAffected Get NodeFieldDatas via TitleAndRevisionTranslationAffected
func GetNodeFieldDatasByTitleAndRevisionTranslationAffected(offset int, limit int, Title_ string, RevisionTranslationAffected_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("title = ? and revision_translation_affected = ?", Title_, RevisionTranslationAffected_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByTitleAndDefaultLangcode Get NodeFieldDatas via TitleAndDefaultLangcode
func GetNodeFieldDatasByTitleAndDefaultLangcode(offset int, limit int, Title_ string, DefaultLangcode_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("title = ? and default_langcode = ?", Title_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByUidAndStatus Get NodeFieldDatas via UidAndStatus
func GetNodeFieldDatasByUidAndStatus(offset int, limit int, Uid_ int, Status_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("uid = ? and status = ?", Uid_, Status_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByUidAndCreated Get NodeFieldDatas via UidAndCreated
func GetNodeFieldDatasByUidAndCreated(offset int, limit int, Uid_ int, Created_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("uid = ? and created = ?", Uid_, Created_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByUidAndChanged Get NodeFieldDatas via UidAndChanged
func GetNodeFieldDatasByUidAndChanged(offset int, limit int, Uid_ int, Changed_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("uid = ? and changed = ?", Uid_, Changed_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByUidAndPromote Get NodeFieldDatas via UidAndPromote
func GetNodeFieldDatasByUidAndPromote(offset int, limit int, Uid_ int, Promote_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("uid = ? and promote = ?", Uid_, Promote_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByUidAndSticky Get NodeFieldDatas via UidAndSticky
func GetNodeFieldDatasByUidAndSticky(offset int, limit int, Uid_ int, Sticky_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("uid = ? and sticky = ?", Uid_, Sticky_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByUidAndRevisionTranslationAffected Get NodeFieldDatas via UidAndRevisionTranslationAffected
func GetNodeFieldDatasByUidAndRevisionTranslationAffected(offset int, limit int, Uid_ int, RevisionTranslationAffected_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("uid = ? and revision_translation_affected = ?", Uid_, RevisionTranslationAffected_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByUidAndDefaultLangcode Get NodeFieldDatas via UidAndDefaultLangcode
func GetNodeFieldDatasByUidAndDefaultLangcode(offset int, limit int, Uid_ int, DefaultLangcode_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("uid = ? and default_langcode = ?", Uid_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByStatusAndCreated Get NodeFieldDatas via StatusAndCreated
func GetNodeFieldDatasByStatusAndCreated(offset int, limit int, Status_ int, Created_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("status = ? and created = ?", Status_, Created_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByStatusAndChanged Get NodeFieldDatas via StatusAndChanged
func GetNodeFieldDatasByStatusAndChanged(offset int, limit int, Status_ int, Changed_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("status = ? and changed = ?", Status_, Changed_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByStatusAndPromote Get NodeFieldDatas via StatusAndPromote
func GetNodeFieldDatasByStatusAndPromote(offset int, limit int, Status_ int, Promote_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("status = ? and promote = ?", Status_, Promote_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByStatusAndSticky Get NodeFieldDatas via StatusAndSticky
func GetNodeFieldDatasByStatusAndSticky(offset int, limit int, Status_ int, Sticky_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("status = ? and sticky = ?", Status_, Sticky_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByStatusAndRevisionTranslationAffected Get NodeFieldDatas via StatusAndRevisionTranslationAffected
func GetNodeFieldDatasByStatusAndRevisionTranslationAffected(offset int, limit int, Status_ int, RevisionTranslationAffected_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("status = ? and revision_translation_affected = ?", Status_, RevisionTranslationAffected_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByStatusAndDefaultLangcode Get NodeFieldDatas via StatusAndDefaultLangcode
func GetNodeFieldDatasByStatusAndDefaultLangcode(offset int, limit int, Status_ int, DefaultLangcode_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("status = ? and default_langcode = ?", Status_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByCreatedAndChanged Get NodeFieldDatas via CreatedAndChanged
func GetNodeFieldDatasByCreatedAndChanged(offset int, limit int, Created_ int, Changed_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("created = ? and changed = ?", Created_, Changed_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByCreatedAndPromote Get NodeFieldDatas via CreatedAndPromote
func GetNodeFieldDatasByCreatedAndPromote(offset int, limit int, Created_ int, Promote_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("created = ? and promote = ?", Created_, Promote_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByCreatedAndSticky Get NodeFieldDatas via CreatedAndSticky
func GetNodeFieldDatasByCreatedAndSticky(offset int, limit int, Created_ int, Sticky_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("created = ? and sticky = ?", Created_, Sticky_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByCreatedAndRevisionTranslationAffected Get NodeFieldDatas via CreatedAndRevisionTranslationAffected
func GetNodeFieldDatasByCreatedAndRevisionTranslationAffected(offset int, limit int, Created_ int, RevisionTranslationAffected_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("created = ? and revision_translation_affected = ?", Created_, RevisionTranslationAffected_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByCreatedAndDefaultLangcode Get NodeFieldDatas via CreatedAndDefaultLangcode
func GetNodeFieldDatasByCreatedAndDefaultLangcode(offset int, limit int, Created_ int, DefaultLangcode_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("created = ? and default_langcode = ?", Created_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByChangedAndPromote Get NodeFieldDatas via ChangedAndPromote
func GetNodeFieldDatasByChangedAndPromote(offset int, limit int, Changed_ int, Promote_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("changed = ? and promote = ?", Changed_, Promote_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByChangedAndSticky Get NodeFieldDatas via ChangedAndSticky
func GetNodeFieldDatasByChangedAndSticky(offset int, limit int, Changed_ int, Sticky_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("changed = ? and sticky = ?", Changed_, Sticky_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByChangedAndRevisionTranslationAffected Get NodeFieldDatas via ChangedAndRevisionTranslationAffected
func GetNodeFieldDatasByChangedAndRevisionTranslationAffected(offset int, limit int, Changed_ int, RevisionTranslationAffected_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("changed = ? and revision_translation_affected = ?", Changed_, RevisionTranslationAffected_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByChangedAndDefaultLangcode Get NodeFieldDatas via ChangedAndDefaultLangcode
func GetNodeFieldDatasByChangedAndDefaultLangcode(offset int, limit int, Changed_ int, DefaultLangcode_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("changed = ? and default_langcode = ?", Changed_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByPromoteAndSticky Get NodeFieldDatas via PromoteAndSticky
func GetNodeFieldDatasByPromoteAndSticky(offset int, limit int, Promote_ int, Sticky_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("promote = ? and sticky = ?", Promote_, Sticky_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByPromoteAndRevisionTranslationAffected Get NodeFieldDatas via PromoteAndRevisionTranslationAffected
func GetNodeFieldDatasByPromoteAndRevisionTranslationAffected(offset int, limit int, Promote_ int, RevisionTranslationAffected_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("promote = ? and revision_translation_affected = ?", Promote_, RevisionTranslationAffected_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByPromoteAndDefaultLangcode Get NodeFieldDatas via PromoteAndDefaultLangcode
func GetNodeFieldDatasByPromoteAndDefaultLangcode(offset int, limit int, Promote_ int, DefaultLangcode_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("promote = ? and default_langcode = ?", Promote_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByStickyAndRevisionTranslationAffected Get NodeFieldDatas via StickyAndRevisionTranslationAffected
func GetNodeFieldDatasByStickyAndRevisionTranslationAffected(offset int, limit int, Sticky_ int, RevisionTranslationAffected_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("sticky = ? and revision_translation_affected = ?", Sticky_, RevisionTranslationAffected_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByStickyAndDefaultLangcode Get NodeFieldDatas via StickyAndDefaultLangcode
func GetNodeFieldDatasByStickyAndDefaultLangcode(offset int, limit int, Sticky_ int, DefaultLangcode_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("sticky = ? and default_langcode = ?", Sticky_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasByRevisionTranslationAffectedAndDefaultLangcode Get NodeFieldDatas via RevisionTranslationAffectedAndDefaultLangcode
func GetNodeFieldDatasByRevisionTranslationAffectedAndDefaultLangcode(offset int, limit int, RevisionTranslationAffected_ int, DefaultLangcode_ int) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("revision_translation_affected = ? and default_langcode = ?", RevisionTranslationAffected_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatas Get NodeFieldDatas via field
func GetNodeFieldDatas(offset int, limit int, field string) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Limit(limit, offset).Desc(field).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasViaNid Get NodeFieldDatas via Nid
func GetNodeFieldDatasViaNid(offset int, limit int, Nid_ int64, field string) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("nid = ?", Nid_).Limit(limit, offset).Desc(field).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasViaVid Get NodeFieldDatas via Vid
func GetNodeFieldDatasViaVid(offset int, limit int, Vid_ int, field string) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("vid = ?", Vid_).Limit(limit, offset).Desc(field).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasViaType Get NodeFieldDatas via Type
func GetNodeFieldDatasViaType(offset int, limit int, Type_ string, field string) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("type = ?", Type_).Limit(limit, offset).Desc(field).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasViaLangcode Get NodeFieldDatas via Langcode
func GetNodeFieldDatasViaLangcode(offset int, limit int, Langcode_ string, field string) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("langcode = ?", Langcode_).Limit(limit, offset).Desc(field).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasViaTitle Get NodeFieldDatas via Title
func GetNodeFieldDatasViaTitle(offset int, limit int, Title_ string, field string) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("title = ?", Title_).Limit(limit, offset).Desc(field).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasViaUid Get NodeFieldDatas via Uid
func GetNodeFieldDatasViaUid(offset int, limit int, Uid_ int, field string) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("uid = ?", Uid_).Limit(limit, offset).Desc(field).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasViaStatus Get NodeFieldDatas via Status
func GetNodeFieldDatasViaStatus(offset int, limit int, Status_ int, field string) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("status = ?", Status_).Limit(limit, offset).Desc(field).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasViaCreated Get NodeFieldDatas via Created
func GetNodeFieldDatasViaCreated(offset int, limit int, Created_ int, field string) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("created = ?", Created_).Limit(limit, offset).Desc(field).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasViaChanged Get NodeFieldDatas via Changed
func GetNodeFieldDatasViaChanged(offset int, limit int, Changed_ int, field string) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("changed = ?", Changed_).Limit(limit, offset).Desc(field).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasViaPromote Get NodeFieldDatas via Promote
func GetNodeFieldDatasViaPromote(offset int, limit int, Promote_ int, field string) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("promote = ?", Promote_).Limit(limit, offset).Desc(field).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasViaSticky Get NodeFieldDatas via Sticky
func GetNodeFieldDatasViaSticky(offset int, limit int, Sticky_ int, field string) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("sticky = ?", Sticky_).Limit(limit, offset).Desc(field).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasViaRevisionTranslationAffected Get NodeFieldDatas via RevisionTranslationAffected
func GetNodeFieldDatasViaRevisionTranslationAffected(offset int, limit int, RevisionTranslationAffected_ int, field string) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("revision_translation_affected = ?", RevisionTranslationAffected_).Limit(limit, offset).Desc(field).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// GetNodeFieldDatasViaDefaultLangcode Get NodeFieldDatas via DefaultLangcode
func GetNodeFieldDatasViaDefaultLangcode(offset int, limit int, DefaultLangcode_ int, field string) (*[]*NodeFieldData, error) {
	var _NodeFieldData = new([]*NodeFieldData)
	err := Engine.Table("node_field_data").Where("default_langcode = ?", DefaultLangcode_).Limit(limit, offset).Desc(field).Find(_NodeFieldData)
	return _NodeFieldData, err
}

// HasNodeFieldDataViaNid Has NodeFieldData via Nid
func HasNodeFieldDataViaNid(iNid int64) bool {
	if has, err := Engine.Where("nid = ?", iNid).Get(new(NodeFieldData)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasNodeFieldDataViaVid Has NodeFieldData via Vid
func HasNodeFieldDataViaVid(iVid int) bool {
	if has, err := Engine.Where("vid = ?", iVid).Get(new(NodeFieldData)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasNodeFieldDataViaType Has NodeFieldData via Type
func HasNodeFieldDataViaType(iType string) bool {
	if has, err := Engine.Where("type = ?", iType).Get(new(NodeFieldData)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasNodeFieldDataViaLangcode Has NodeFieldData via Langcode
func HasNodeFieldDataViaLangcode(iLangcode string) bool {
	if has, err := Engine.Where("langcode = ?", iLangcode).Get(new(NodeFieldData)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasNodeFieldDataViaTitle Has NodeFieldData via Title
func HasNodeFieldDataViaTitle(iTitle string) bool {
	if has, err := Engine.Where("title = ?", iTitle).Get(new(NodeFieldData)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasNodeFieldDataViaUid Has NodeFieldData via Uid
func HasNodeFieldDataViaUid(iUid int) bool {
	if has, err := Engine.Where("uid = ?", iUid).Get(new(NodeFieldData)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasNodeFieldDataViaStatus Has NodeFieldData via Status
func HasNodeFieldDataViaStatus(iStatus int) bool {
	if has, err := Engine.Where("status = ?", iStatus).Get(new(NodeFieldData)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasNodeFieldDataViaCreated Has NodeFieldData via Created
func HasNodeFieldDataViaCreated(iCreated int) bool {
	if has, err := Engine.Where("created = ?", iCreated).Get(new(NodeFieldData)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasNodeFieldDataViaChanged Has NodeFieldData via Changed
func HasNodeFieldDataViaChanged(iChanged int) bool {
	if has, err := Engine.Where("changed = ?", iChanged).Get(new(NodeFieldData)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasNodeFieldDataViaPromote Has NodeFieldData via Promote
func HasNodeFieldDataViaPromote(iPromote int) bool {
	if has, err := Engine.Where("promote = ?", iPromote).Get(new(NodeFieldData)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasNodeFieldDataViaSticky Has NodeFieldData via Sticky
func HasNodeFieldDataViaSticky(iSticky int) bool {
	if has, err := Engine.Where("sticky = ?", iSticky).Get(new(NodeFieldData)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasNodeFieldDataViaRevisionTranslationAffected Has NodeFieldData via RevisionTranslationAffected
func HasNodeFieldDataViaRevisionTranslationAffected(iRevisionTranslationAffected int) bool {
	if has, err := Engine.Where("revision_translation_affected = ?", iRevisionTranslationAffected).Get(new(NodeFieldData)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasNodeFieldDataViaDefaultLangcode Has NodeFieldData via DefaultLangcode
func HasNodeFieldDataViaDefaultLangcode(iDefaultLangcode int) bool {
	if has, err := Engine.Where("default_langcode = ?", iDefaultLangcode).Get(new(NodeFieldData)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// GetNodeFieldDataViaNid Get NodeFieldData via Nid
func GetNodeFieldDataViaNid(iNid int64) (*NodeFieldData, error) {
	var _NodeFieldData = &NodeFieldData{Nid: iNid}
	has, err := Engine.Get(_NodeFieldData)
	if has {
		return _NodeFieldData, err
	} else {
		return nil, err
	}
}

// GetNodeFieldDataViaVid Get NodeFieldData via Vid
func GetNodeFieldDataViaVid(iVid int) (*NodeFieldData, error) {
	var _NodeFieldData = &NodeFieldData{Vid: iVid}
	has, err := Engine.Get(_NodeFieldData)
	if has {
		return _NodeFieldData, err
	} else {
		return nil, err
	}
}

// GetNodeFieldDataViaType Get NodeFieldData via Type
func GetNodeFieldDataViaType(iType string) (*NodeFieldData, error) {
	var _NodeFieldData = &NodeFieldData{Type: iType}
	has, err := Engine.Get(_NodeFieldData)
	if has {
		return _NodeFieldData, err
	} else {
		return nil, err
	}
}

// GetNodeFieldDataViaLangcode Get NodeFieldData via Langcode
func GetNodeFieldDataViaLangcode(iLangcode string) (*NodeFieldData, error) {
	var _NodeFieldData = &NodeFieldData{Langcode: iLangcode}
	has, err := Engine.Get(_NodeFieldData)
	if has {
		return _NodeFieldData, err
	} else {
		return nil, err
	}
}

// GetNodeFieldDataViaTitle Get NodeFieldData via Title
func GetNodeFieldDataViaTitle(iTitle string) (*NodeFieldData, error) {
	var _NodeFieldData = &NodeFieldData{Title: iTitle}
	has, err := Engine.Get(_NodeFieldData)
	if has {
		return _NodeFieldData, err
	} else {
		return nil, err
	}
}

// GetNodeFieldDataViaUid Get NodeFieldData via Uid
func GetNodeFieldDataViaUid(iUid int) (*NodeFieldData, error) {
	var _NodeFieldData = &NodeFieldData{Uid: iUid}
	has, err := Engine.Get(_NodeFieldData)
	if has {
		return _NodeFieldData, err
	} else {
		return nil, err
	}
}

// GetNodeFieldDataViaStatus Get NodeFieldData via Status
func GetNodeFieldDataViaStatus(iStatus int) (*NodeFieldData, error) {
	var _NodeFieldData = &NodeFieldData{Status: iStatus}
	has, err := Engine.Get(_NodeFieldData)
	if has {
		return _NodeFieldData, err
	} else {
		return nil, err
	}
}

// GetNodeFieldDataViaCreated Get NodeFieldData via Created
func GetNodeFieldDataViaCreated(iCreated int) (*NodeFieldData, error) {
	var _NodeFieldData = &NodeFieldData{Created: iCreated}
	has, err := Engine.Get(_NodeFieldData)
	if has {
		return _NodeFieldData, err
	} else {
		return nil, err
	}
}

// GetNodeFieldDataViaChanged Get NodeFieldData via Changed
func GetNodeFieldDataViaChanged(iChanged int) (*NodeFieldData, error) {
	var _NodeFieldData = &NodeFieldData{Changed: iChanged}
	has, err := Engine.Get(_NodeFieldData)
	if has {
		return _NodeFieldData, err
	} else {
		return nil, err
	}
}

// GetNodeFieldDataViaPromote Get NodeFieldData via Promote
func GetNodeFieldDataViaPromote(iPromote int) (*NodeFieldData, error) {
	var _NodeFieldData = &NodeFieldData{Promote: iPromote}
	has, err := Engine.Get(_NodeFieldData)
	if has {
		return _NodeFieldData, err
	} else {
		return nil, err
	}
}

// GetNodeFieldDataViaSticky Get NodeFieldData via Sticky
func GetNodeFieldDataViaSticky(iSticky int) (*NodeFieldData, error) {
	var _NodeFieldData = &NodeFieldData{Sticky: iSticky}
	has, err := Engine.Get(_NodeFieldData)
	if has {
		return _NodeFieldData, err
	} else {
		return nil, err
	}
}

// GetNodeFieldDataViaRevisionTranslationAffected Get NodeFieldData via RevisionTranslationAffected
func GetNodeFieldDataViaRevisionTranslationAffected(iRevisionTranslationAffected int) (*NodeFieldData, error) {
	var _NodeFieldData = &NodeFieldData{RevisionTranslationAffected: iRevisionTranslationAffected}
	has, err := Engine.Get(_NodeFieldData)
	if has {
		return _NodeFieldData, err
	} else {
		return nil, err
	}
}

// GetNodeFieldDataViaDefaultLangcode Get NodeFieldData via DefaultLangcode
func GetNodeFieldDataViaDefaultLangcode(iDefaultLangcode int) (*NodeFieldData, error) {
	var _NodeFieldData = &NodeFieldData{DefaultLangcode: iDefaultLangcode}
	has, err := Engine.Get(_NodeFieldData)
	if has {
		return _NodeFieldData, err
	} else {
		return nil, err
	}
}

// SetNodeFieldDataViaNid Set NodeFieldData via Nid
func SetNodeFieldDataViaNid(iNid int64, node_field_data *NodeFieldData) (int64, error) {
	node_field_data.Nid = iNid
	return Engine.Insert(node_field_data)
}

// SetNodeFieldDataViaVid Set NodeFieldData via Vid
func SetNodeFieldDataViaVid(iVid int, node_field_data *NodeFieldData) (int64, error) {
	node_field_data.Vid = iVid
	return Engine.Insert(node_field_data)
}

// SetNodeFieldDataViaType Set NodeFieldData via Type
func SetNodeFieldDataViaType(iType string, node_field_data *NodeFieldData) (int64, error) {
	node_field_data.Type = iType
	return Engine.Insert(node_field_data)
}

// SetNodeFieldDataViaLangcode Set NodeFieldData via Langcode
func SetNodeFieldDataViaLangcode(iLangcode string, node_field_data *NodeFieldData) (int64, error) {
	node_field_data.Langcode = iLangcode
	return Engine.Insert(node_field_data)
}

// SetNodeFieldDataViaTitle Set NodeFieldData via Title
func SetNodeFieldDataViaTitle(iTitle string, node_field_data *NodeFieldData) (int64, error) {
	node_field_data.Title = iTitle
	return Engine.Insert(node_field_data)
}

// SetNodeFieldDataViaUid Set NodeFieldData via Uid
func SetNodeFieldDataViaUid(iUid int, node_field_data *NodeFieldData) (int64, error) {
	node_field_data.Uid = iUid
	return Engine.Insert(node_field_data)
}

// SetNodeFieldDataViaStatus Set NodeFieldData via Status
func SetNodeFieldDataViaStatus(iStatus int, node_field_data *NodeFieldData) (int64, error) {
	node_field_data.Status = iStatus
	return Engine.Insert(node_field_data)
}

// SetNodeFieldDataViaCreated Set NodeFieldData via Created
func SetNodeFieldDataViaCreated(iCreated int, node_field_data *NodeFieldData) (int64, error) {
	node_field_data.Created = iCreated
	return Engine.Insert(node_field_data)
}

// SetNodeFieldDataViaChanged Set NodeFieldData via Changed
func SetNodeFieldDataViaChanged(iChanged int, node_field_data *NodeFieldData) (int64, error) {
	node_field_data.Changed = iChanged
	return Engine.Insert(node_field_data)
}

// SetNodeFieldDataViaPromote Set NodeFieldData via Promote
func SetNodeFieldDataViaPromote(iPromote int, node_field_data *NodeFieldData) (int64, error) {
	node_field_data.Promote = iPromote
	return Engine.Insert(node_field_data)
}

// SetNodeFieldDataViaSticky Set NodeFieldData via Sticky
func SetNodeFieldDataViaSticky(iSticky int, node_field_data *NodeFieldData) (int64, error) {
	node_field_data.Sticky = iSticky
	return Engine.Insert(node_field_data)
}

// SetNodeFieldDataViaRevisionTranslationAffected Set NodeFieldData via RevisionTranslationAffected
func SetNodeFieldDataViaRevisionTranslationAffected(iRevisionTranslationAffected int, node_field_data *NodeFieldData) (int64, error) {
	node_field_data.RevisionTranslationAffected = iRevisionTranslationAffected
	return Engine.Insert(node_field_data)
}

// SetNodeFieldDataViaDefaultLangcode Set NodeFieldData via DefaultLangcode
func SetNodeFieldDataViaDefaultLangcode(iDefaultLangcode int, node_field_data *NodeFieldData) (int64, error) {
	node_field_data.DefaultLangcode = iDefaultLangcode
	return Engine.Insert(node_field_data)
}

// AddNodeFieldData Add NodeFieldData via all columns
func AddNodeFieldData(iNid int64, iVid int, iType string, iLangcode string, iTitle string, iUid int, iStatus int, iCreated int, iChanged int, iPromote int, iSticky int, iRevisionTranslationAffected int, iDefaultLangcode int) error {
	_NodeFieldData := &NodeFieldData{Nid: iNid, Vid: iVid, Type: iType, Langcode: iLangcode, Title: iTitle, Uid: iUid, Status: iStatus, Created: iCreated, Changed: iChanged, Promote: iPromote, Sticky: iSticky, RevisionTranslationAffected: iRevisionTranslationAffected, DefaultLangcode: iDefaultLangcode}
	if _, err := Engine.Insert(_NodeFieldData); err != nil {
		return err
	} else {
		return nil
	}
}

// PostNodeFieldData Post NodeFieldData via iNodeFieldData
func PostNodeFieldData(iNodeFieldData *NodeFieldData) (int64, error) {
	_, err := Engine.Insert(iNodeFieldData)
	return iNodeFieldData.Nid, err
}

// PutNodeFieldData Put NodeFieldData
func PutNodeFieldData(iNodeFieldData *NodeFieldData) (int64, error) {
	_, err := Engine.Id(iNodeFieldData.Nid).Update(iNodeFieldData)
	return iNodeFieldData.Nid, err
}

// PutNodeFieldDataViaNid Put NodeFieldData via Nid
func PutNodeFieldDataViaNid(Nid_ int64, iNodeFieldData *NodeFieldData) (int64, error) {
	row, err := Engine.Update(iNodeFieldData, &NodeFieldData{Nid: Nid_})
	return row, err
}

// PutNodeFieldDataViaVid Put NodeFieldData via Vid
func PutNodeFieldDataViaVid(Vid_ int, iNodeFieldData *NodeFieldData) (int64, error) {
	row, err := Engine.Update(iNodeFieldData, &NodeFieldData{Vid: Vid_})
	return row, err
}

// PutNodeFieldDataViaType Put NodeFieldData via Type
func PutNodeFieldDataViaType(Type_ string, iNodeFieldData *NodeFieldData) (int64, error) {
	row, err := Engine.Update(iNodeFieldData, &NodeFieldData{Type: Type_})
	return row, err
}

// PutNodeFieldDataViaLangcode Put NodeFieldData via Langcode
func PutNodeFieldDataViaLangcode(Langcode_ string, iNodeFieldData *NodeFieldData) (int64, error) {
	row, err := Engine.Update(iNodeFieldData, &NodeFieldData{Langcode: Langcode_})
	return row, err
}

// PutNodeFieldDataViaTitle Put NodeFieldData via Title
func PutNodeFieldDataViaTitle(Title_ string, iNodeFieldData *NodeFieldData) (int64, error) {
	row, err := Engine.Update(iNodeFieldData, &NodeFieldData{Title: Title_})
	return row, err
}

// PutNodeFieldDataViaUid Put NodeFieldData via Uid
func PutNodeFieldDataViaUid(Uid_ int, iNodeFieldData *NodeFieldData) (int64, error) {
	row, err := Engine.Update(iNodeFieldData, &NodeFieldData{Uid: Uid_})
	return row, err
}

// PutNodeFieldDataViaStatus Put NodeFieldData via Status
func PutNodeFieldDataViaStatus(Status_ int, iNodeFieldData *NodeFieldData) (int64, error) {
	row, err := Engine.Update(iNodeFieldData, &NodeFieldData{Status: Status_})
	return row, err
}

// PutNodeFieldDataViaCreated Put NodeFieldData via Created
func PutNodeFieldDataViaCreated(Created_ int, iNodeFieldData *NodeFieldData) (int64, error) {
	row, err := Engine.Update(iNodeFieldData, &NodeFieldData{Created: Created_})
	return row, err
}

// PutNodeFieldDataViaChanged Put NodeFieldData via Changed
func PutNodeFieldDataViaChanged(Changed_ int, iNodeFieldData *NodeFieldData) (int64, error) {
	row, err := Engine.Update(iNodeFieldData, &NodeFieldData{Changed: Changed_})
	return row, err
}

// PutNodeFieldDataViaPromote Put NodeFieldData via Promote
func PutNodeFieldDataViaPromote(Promote_ int, iNodeFieldData *NodeFieldData) (int64, error) {
	row, err := Engine.Update(iNodeFieldData, &NodeFieldData{Promote: Promote_})
	return row, err
}

// PutNodeFieldDataViaSticky Put NodeFieldData via Sticky
func PutNodeFieldDataViaSticky(Sticky_ int, iNodeFieldData *NodeFieldData) (int64, error) {
	row, err := Engine.Update(iNodeFieldData, &NodeFieldData{Sticky: Sticky_})
	return row, err
}

// PutNodeFieldDataViaRevisionTranslationAffected Put NodeFieldData via RevisionTranslationAffected
func PutNodeFieldDataViaRevisionTranslationAffected(RevisionTranslationAffected_ int, iNodeFieldData *NodeFieldData) (int64, error) {
	row, err := Engine.Update(iNodeFieldData, &NodeFieldData{RevisionTranslationAffected: RevisionTranslationAffected_})
	return row, err
}

// PutNodeFieldDataViaDefaultLangcode Put NodeFieldData via DefaultLangcode
func PutNodeFieldDataViaDefaultLangcode(DefaultLangcode_ int, iNodeFieldData *NodeFieldData) (int64, error) {
	row, err := Engine.Update(iNodeFieldData, &NodeFieldData{DefaultLangcode: DefaultLangcode_})
	return row, err
}

// UpdateNodeFieldDataViaNid via map[string]interface{}{}
func UpdateNodeFieldDataViaNid(iNid int64, iNodeFieldDataMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(NodeFieldData)).Where("nid = ?", iNid).Update(iNodeFieldDataMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateNodeFieldDataViaVid via map[string]interface{}{}
func UpdateNodeFieldDataViaVid(iVid int, iNodeFieldDataMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(NodeFieldData)).Where("vid = ?", iVid).Update(iNodeFieldDataMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateNodeFieldDataViaType via map[string]interface{}{}
func UpdateNodeFieldDataViaType(iType string, iNodeFieldDataMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(NodeFieldData)).Where("type = ?", iType).Update(iNodeFieldDataMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateNodeFieldDataViaLangcode via map[string]interface{}{}
func UpdateNodeFieldDataViaLangcode(iLangcode string, iNodeFieldDataMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(NodeFieldData)).Where("langcode = ?", iLangcode).Update(iNodeFieldDataMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateNodeFieldDataViaTitle via map[string]interface{}{}
func UpdateNodeFieldDataViaTitle(iTitle string, iNodeFieldDataMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(NodeFieldData)).Where("title = ?", iTitle).Update(iNodeFieldDataMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateNodeFieldDataViaUid via map[string]interface{}{}
func UpdateNodeFieldDataViaUid(iUid int, iNodeFieldDataMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(NodeFieldData)).Where("uid = ?", iUid).Update(iNodeFieldDataMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateNodeFieldDataViaStatus via map[string]interface{}{}
func UpdateNodeFieldDataViaStatus(iStatus int, iNodeFieldDataMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(NodeFieldData)).Where("status = ?", iStatus).Update(iNodeFieldDataMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateNodeFieldDataViaCreated via map[string]interface{}{}
func UpdateNodeFieldDataViaCreated(iCreated int, iNodeFieldDataMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(NodeFieldData)).Where("created = ?", iCreated).Update(iNodeFieldDataMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateNodeFieldDataViaChanged via map[string]interface{}{}
func UpdateNodeFieldDataViaChanged(iChanged int, iNodeFieldDataMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(NodeFieldData)).Where("changed = ?", iChanged).Update(iNodeFieldDataMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateNodeFieldDataViaPromote via map[string]interface{}{}
func UpdateNodeFieldDataViaPromote(iPromote int, iNodeFieldDataMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(NodeFieldData)).Where("promote = ?", iPromote).Update(iNodeFieldDataMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateNodeFieldDataViaSticky via map[string]interface{}{}
func UpdateNodeFieldDataViaSticky(iSticky int, iNodeFieldDataMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(NodeFieldData)).Where("sticky = ?", iSticky).Update(iNodeFieldDataMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateNodeFieldDataViaRevisionTranslationAffected via map[string]interface{}{}
func UpdateNodeFieldDataViaRevisionTranslationAffected(iRevisionTranslationAffected int, iNodeFieldDataMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(NodeFieldData)).Where("revision_translation_affected = ?", iRevisionTranslationAffected).Update(iNodeFieldDataMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateNodeFieldDataViaDefaultLangcode via map[string]interface{}{}
func UpdateNodeFieldDataViaDefaultLangcode(iDefaultLangcode int, iNodeFieldDataMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(NodeFieldData)).Where("default_langcode = ?", iDefaultLangcode).Update(iNodeFieldDataMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// DeleteNodeFieldData Delete NodeFieldData
func DeleteNodeFieldData(iNid int64) (int64, error) {
	row, err := Engine.Id(iNid).Delete(new(NodeFieldData))
	return row, err
}

// DeleteNodeFieldDataViaNid Delete NodeFieldData via Nid
func DeleteNodeFieldDataViaNid(iNid int64) (err error) {
	var has bool
	var _NodeFieldData = &NodeFieldData{Nid: iNid}
	if has, err = Engine.Get(_NodeFieldData); (has == true) && (err == nil) {
		if row, err := Engine.Where("nid = ?", iNid).Delete(new(NodeFieldData)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteNodeFieldDataViaVid Delete NodeFieldData via Vid
func DeleteNodeFieldDataViaVid(iVid int) (err error) {
	var has bool
	var _NodeFieldData = &NodeFieldData{Vid: iVid}
	if has, err = Engine.Get(_NodeFieldData); (has == true) && (err == nil) {
		if row, err := Engine.Where("vid = ?", iVid).Delete(new(NodeFieldData)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteNodeFieldDataViaType Delete NodeFieldData via Type
func DeleteNodeFieldDataViaType(iType string) (err error) {
	var has bool
	var _NodeFieldData = &NodeFieldData{Type: iType}
	if has, err = Engine.Get(_NodeFieldData); (has == true) && (err == nil) {
		if row, err := Engine.Where("type = ?", iType).Delete(new(NodeFieldData)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteNodeFieldDataViaLangcode Delete NodeFieldData via Langcode
func DeleteNodeFieldDataViaLangcode(iLangcode string) (err error) {
	var has bool
	var _NodeFieldData = &NodeFieldData{Langcode: iLangcode}
	if has, err = Engine.Get(_NodeFieldData); (has == true) && (err == nil) {
		if row, err := Engine.Where("langcode = ?", iLangcode).Delete(new(NodeFieldData)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteNodeFieldDataViaTitle Delete NodeFieldData via Title
func DeleteNodeFieldDataViaTitle(iTitle string) (err error) {
	var has bool
	var _NodeFieldData = &NodeFieldData{Title: iTitle}
	if has, err = Engine.Get(_NodeFieldData); (has == true) && (err == nil) {
		if row, err := Engine.Where("title = ?", iTitle).Delete(new(NodeFieldData)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteNodeFieldDataViaUid Delete NodeFieldData via Uid
func DeleteNodeFieldDataViaUid(iUid int) (err error) {
	var has bool
	var _NodeFieldData = &NodeFieldData{Uid: iUid}
	if has, err = Engine.Get(_NodeFieldData); (has == true) && (err == nil) {
		if row, err := Engine.Where("uid = ?", iUid).Delete(new(NodeFieldData)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteNodeFieldDataViaStatus Delete NodeFieldData via Status
func DeleteNodeFieldDataViaStatus(iStatus int) (err error) {
	var has bool
	var _NodeFieldData = &NodeFieldData{Status: iStatus}
	if has, err = Engine.Get(_NodeFieldData); (has == true) && (err == nil) {
		if row, err := Engine.Where("status = ?", iStatus).Delete(new(NodeFieldData)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteNodeFieldDataViaCreated Delete NodeFieldData via Created
func DeleteNodeFieldDataViaCreated(iCreated int) (err error) {
	var has bool
	var _NodeFieldData = &NodeFieldData{Created: iCreated}
	if has, err = Engine.Get(_NodeFieldData); (has == true) && (err == nil) {
		if row, err := Engine.Where("created = ?", iCreated).Delete(new(NodeFieldData)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteNodeFieldDataViaChanged Delete NodeFieldData via Changed
func DeleteNodeFieldDataViaChanged(iChanged int) (err error) {
	var has bool
	var _NodeFieldData = &NodeFieldData{Changed: iChanged}
	if has, err = Engine.Get(_NodeFieldData); (has == true) && (err == nil) {
		if row, err := Engine.Where("changed = ?", iChanged).Delete(new(NodeFieldData)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteNodeFieldDataViaPromote Delete NodeFieldData via Promote
func DeleteNodeFieldDataViaPromote(iPromote int) (err error) {
	var has bool
	var _NodeFieldData = &NodeFieldData{Promote: iPromote}
	if has, err = Engine.Get(_NodeFieldData); (has == true) && (err == nil) {
		if row, err := Engine.Where("promote = ?", iPromote).Delete(new(NodeFieldData)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteNodeFieldDataViaSticky Delete NodeFieldData via Sticky
func DeleteNodeFieldDataViaSticky(iSticky int) (err error) {
	var has bool
	var _NodeFieldData = &NodeFieldData{Sticky: iSticky}
	if has, err = Engine.Get(_NodeFieldData); (has == true) && (err == nil) {
		if row, err := Engine.Where("sticky = ?", iSticky).Delete(new(NodeFieldData)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteNodeFieldDataViaRevisionTranslationAffected Delete NodeFieldData via RevisionTranslationAffected
func DeleteNodeFieldDataViaRevisionTranslationAffected(iRevisionTranslationAffected int) (err error) {
	var has bool
	var _NodeFieldData = &NodeFieldData{RevisionTranslationAffected: iRevisionTranslationAffected}
	if has, err = Engine.Get(_NodeFieldData); (has == true) && (err == nil) {
		if row, err := Engine.Where("revision_translation_affected = ?", iRevisionTranslationAffected).Delete(new(NodeFieldData)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteNodeFieldDataViaDefaultLangcode Delete NodeFieldData via DefaultLangcode
func DeleteNodeFieldDataViaDefaultLangcode(iDefaultLangcode int) (err error) {
	var has bool
	var _NodeFieldData = &NodeFieldData{DefaultLangcode: iDefaultLangcode}
	if has, err = Engine.Get(_NodeFieldData); (has == true) && (err == nil) {
		if row, err := Engine.Where("default_langcode = ?", iDefaultLangcode).Delete(new(NodeFieldData)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}
