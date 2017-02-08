package model

type NodeFieldRevision struct {
	Nid                         int64  `xorm:"not null index(node__id__default_langcode__langcode) INT(10)"`
	Vid                         int    `xorm:"not null pk INT(10)"`
	Langcode                    string `xorm:"not null pk index(node__id__default_langcode__langcode) VARCHAR(12)"`
	Title                       string `xorm:"VARCHAR(255)"`
	Uid                         int    `xorm:"not null index INT(10)"`
	Status                      int    `xorm:"not null TINYINT(4)"`
	Created                     int    `xorm:"INT(11)"`
	Changed                     int    `xorm:"INT(11)"`
	Promote                     int    `xorm:"TINYINT(4)"`
	Sticky                      int    `xorm:"TINYINT(4)"`
	RevisionTranslationAffected int    `xorm:"TINYINT(4)"`
	DefaultLangcode             int    `xorm:"not null index(node__id__default_langcode__langcode) TINYINT(4)"`
}

// GetNodeFieldRevisionsCount NodeFieldRevisions' Count
func GetNodeFieldRevisionsCount(offset int, limit int) (int64, error) {
	total, err := Engine.Limit(limit, offset).Count(&NodeFieldRevision{})
	return total, err
}

// GetNodeFieldRevisionCountViaNid Get NodeFieldRevision via Nid
func GetNodeFieldRevisionCountViaNid(iNid int64) int64 {
	n, _ := Engine.Where("nid = ?", iNid).Count(&NodeFieldRevision{Nid: iNid})
	return n
}

// GetNodeFieldRevisionCountViaVid Get NodeFieldRevision via Vid
func GetNodeFieldRevisionCountViaVid(iVid int) int64 {
	n, _ := Engine.Where("vid = ?", iVid).Count(&NodeFieldRevision{Vid: iVid})
	return n
}

// GetNodeFieldRevisionCountViaLangcode Get NodeFieldRevision via Langcode
func GetNodeFieldRevisionCountViaLangcode(iLangcode string) int64 {
	n, _ := Engine.Where("langcode = ?", iLangcode).Count(&NodeFieldRevision{Langcode: iLangcode})
	return n
}

// GetNodeFieldRevisionCountViaTitle Get NodeFieldRevision via Title
func GetNodeFieldRevisionCountViaTitle(iTitle string) int64 {
	n, _ := Engine.Where("title = ?", iTitle).Count(&NodeFieldRevision{Title: iTitle})
	return n
}

// GetNodeFieldRevisionCountViaUid Get NodeFieldRevision via Uid
func GetNodeFieldRevisionCountViaUid(iUid int) int64 {
	n, _ := Engine.Where("uid = ?", iUid).Count(&NodeFieldRevision{Uid: iUid})
	return n
}

// GetNodeFieldRevisionCountViaStatus Get NodeFieldRevision via Status
func GetNodeFieldRevisionCountViaStatus(iStatus int) int64 {
	n, _ := Engine.Where("status = ?", iStatus).Count(&NodeFieldRevision{Status: iStatus})
	return n
}

// GetNodeFieldRevisionCountViaCreated Get NodeFieldRevision via Created
func GetNodeFieldRevisionCountViaCreated(iCreated int) int64 {
	n, _ := Engine.Where("created = ?", iCreated).Count(&NodeFieldRevision{Created: iCreated})
	return n
}

// GetNodeFieldRevisionCountViaChanged Get NodeFieldRevision via Changed
func GetNodeFieldRevisionCountViaChanged(iChanged int) int64 {
	n, _ := Engine.Where("changed = ?", iChanged).Count(&NodeFieldRevision{Changed: iChanged})
	return n
}

// GetNodeFieldRevisionCountViaPromote Get NodeFieldRevision via Promote
func GetNodeFieldRevisionCountViaPromote(iPromote int) int64 {
	n, _ := Engine.Where("promote = ?", iPromote).Count(&NodeFieldRevision{Promote: iPromote})
	return n
}

// GetNodeFieldRevisionCountViaSticky Get NodeFieldRevision via Sticky
func GetNodeFieldRevisionCountViaSticky(iSticky int) int64 {
	n, _ := Engine.Where("sticky = ?", iSticky).Count(&NodeFieldRevision{Sticky: iSticky})
	return n
}

// GetNodeFieldRevisionCountViaRevisionTranslationAffected Get NodeFieldRevision via RevisionTranslationAffected
func GetNodeFieldRevisionCountViaRevisionTranslationAffected(iRevisionTranslationAffected int) int64 {
	n, _ := Engine.Where("revision_translation_affected = ?", iRevisionTranslationAffected).Count(&NodeFieldRevision{RevisionTranslationAffected: iRevisionTranslationAffected})
	return n
}

// GetNodeFieldRevisionCountViaDefaultLangcode Get NodeFieldRevision via DefaultLangcode
func GetNodeFieldRevisionCountViaDefaultLangcode(iDefaultLangcode int) int64 {
	n, _ := Engine.Where("default_langcode = ?", iDefaultLangcode).Count(&NodeFieldRevision{DefaultLangcode: iDefaultLangcode})
	return n
}

// GetNodeFieldRevisionsByNidAndVidAndLangcode Get NodeFieldRevisions via NidAndVidAndLangcode
func GetNodeFieldRevisionsByNidAndVidAndLangcode(offset int, limit int, Nid_ int64, Vid_ int, Langcode_ string) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("nid = ? and vid = ? and langcode = ?", Nid_, Vid_, Langcode_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByNidAndVidAndTitle Get NodeFieldRevisions via NidAndVidAndTitle
func GetNodeFieldRevisionsByNidAndVidAndTitle(offset int, limit int, Nid_ int64, Vid_ int, Title_ string) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("nid = ? and vid = ? and title = ?", Nid_, Vid_, Title_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByNidAndVidAndUid Get NodeFieldRevisions via NidAndVidAndUid
func GetNodeFieldRevisionsByNidAndVidAndUid(offset int, limit int, Nid_ int64, Vid_ int, Uid_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("nid = ? and vid = ? and uid = ?", Nid_, Vid_, Uid_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByNidAndVidAndStatus Get NodeFieldRevisions via NidAndVidAndStatus
func GetNodeFieldRevisionsByNidAndVidAndStatus(offset int, limit int, Nid_ int64, Vid_ int, Status_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("nid = ? and vid = ? and status = ?", Nid_, Vid_, Status_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByNidAndVidAndCreated Get NodeFieldRevisions via NidAndVidAndCreated
func GetNodeFieldRevisionsByNidAndVidAndCreated(offset int, limit int, Nid_ int64, Vid_ int, Created_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("nid = ? and vid = ? and created = ?", Nid_, Vid_, Created_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByNidAndVidAndChanged Get NodeFieldRevisions via NidAndVidAndChanged
func GetNodeFieldRevisionsByNidAndVidAndChanged(offset int, limit int, Nid_ int64, Vid_ int, Changed_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("nid = ? and vid = ? and changed = ?", Nid_, Vid_, Changed_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByNidAndVidAndPromote Get NodeFieldRevisions via NidAndVidAndPromote
func GetNodeFieldRevisionsByNidAndVidAndPromote(offset int, limit int, Nid_ int64, Vid_ int, Promote_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("nid = ? and vid = ? and promote = ?", Nid_, Vid_, Promote_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByNidAndVidAndSticky Get NodeFieldRevisions via NidAndVidAndSticky
func GetNodeFieldRevisionsByNidAndVidAndSticky(offset int, limit int, Nid_ int64, Vid_ int, Sticky_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("nid = ? and vid = ? and sticky = ?", Nid_, Vid_, Sticky_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByNidAndVidAndRevisionTranslationAffected Get NodeFieldRevisions via NidAndVidAndRevisionTranslationAffected
func GetNodeFieldRevisionsByNidAndVidAndRevisionTranslationAffected(offset int, limit int, Nid_ int64, Vid_ int, RevisionTranslationAffected_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("nid = ? and vid = ? and revision_translation_affected = ?", Nid_, Vid_, RevisionTranslationAffected_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByNidAndVidAndDefaultLangcode Get NodeFieldRevisions via NidAndVidAndDefaultLangcode
func GetNodeFieldRevisionsByNidAndVidAndDefaultLangcode(offset int, limit int, Nid_ int64, Vid_ int, DefaultLangcode_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("nid = ? and vid = ? and default_langcode = ?", Nid_, Vid_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByNidAndLangcodeAndTitle Get NodeFieldRevisions via NidAndLangcodeAndTitle
func GetNodeFieldRevisionsByNidAndLangcodeAndTitle(offset int, limit int, Nid_ int64, Langcode_ string, Title_ string) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("nid = ? and langcode = ? and title = ?", Nid_, Langcode_, Title_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByNidAndLangcodeAndUid Get NodeFieldRevisions via NidAndLangcodeAndUid
func GetNodeFieldRevisionsByNidAndLangcodeAndUid(offset int, limit int, Nid_ int64, Langcode_ string, Uid_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("nid = ? and langcode = ? and uid = ?", Nid_, Langcode_, Uid_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByNidAndLangcodeAndStatus Get NodeFieldRevisions via NidAndLangcodeAndStatus
func GetNodeFieldRevisionsByNidAndLangcodeAndStatus(offset int, limit int, Nid_ int64, Langcode_ string, Status_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("nid = ? and langcode = ? and status = ?", Nid_, Langcode_, Status_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByNidAndLangcodeAndCreated Get NodeFieldRevisions via NidAndLangcodeAndCreated
func GetNodeFieldRevisionsByNidAndLangcodeAndCreated(offset int, limit int, Nid_ int64, Langcode_ string, Created_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("nid = ? and langcode = ? and created = ?", Nid_, Langcode_, Created_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByNidAndLangcodeAndChanged Get NodeFieldRevisions via NidAndLangcodeAndChanged
func GetNodeFieldRevisionsByNidAndLangcodeAndChanged(offset int, limit int, Nid_ int64, Langcode_ string, Changed_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("nid = ? and langcode = ? and changed = ?", Nid_, Langcode_, Changed_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByNidAndLangcodeAndPromote Get NodeFieldRevisions via NidAndLangcodeAndPromote
func GetNodeFieldRevisionsByNidAndLangcodeAndPromote(offset int, limit int, Nid_ int64, Langcode_ string, Promote_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("nid = ? and langcode = ? and promote = ?", Nid_, Langcode_, Promote_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByNidAndLangcodeAndSticky Get NodeFieldRevisions via NidAndLangcodeAndSticky
func GetNodeFieldRevisionsByNidAndLangcodeAndSticky(offset int, limit int, Nid_ int64, Langcode_ string, Sticky_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("nid = ? and langcode = ? and sticky = ?", Nid_, Langcode_, Sticky_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByNidAndLangcodeAndRevisionTranslationAffected Get NodeFieldRevisions via NidAndLangcodeAndRevisionTranslationAffected
func GetNodeFieldRevisionsByNidAndLangcodeAndRevisionTranslationAffected(offset int, limit int, Nid_ int64, Langcode_ string, RevisionTranslationAffected_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("nid = ? and langcode = ? and revision_translation_affected = ?", Nid_, Langcode_, RevisionTranslationAffected_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByNidAndLangcodeAndDefaultLangcode Get NodeFieldRevisions via NidAndLangcodeAndDefaultLangcode
func GetNodeFieldRevisionsByNidAndLangcodeAndDefaultLangcode(offset int, limit int, Nid_ int64, Langcode_ string, DefaultLangcode_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("nid = ? and langcode = ? and default_langcode = ?", Nid_, Langcode_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByNidAndTitleAndUid Get NodeFieldRevisions via NidAndTitleAndUid
func GetNodeFieldRevisionsByNidAndTitleAndUid(offset int, limit int, Nid_ int64, Title_ string, Uid_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("nid = ? and title = ? and uid = ?", Nid_, Title_, Uid_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByNidAndTitleAndStatus Get NodeFieldRevisions via NidAndTitleAndStatus
func GetNodeFieldRevisionsByNidAndTitleAndStatus(offset int, limit int, Nid_ int64, Title_ string, Status_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("nid = ? and title = ? and status = ?", Nid_, Title_, Status_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByNidAndTitleAndCreated Get NodeFieldRevisions via NidAndTitleAndCreated
func GetNodeFieldRevisionsByNidAndTitleAndCreated(offset int, limit int, Nid_ int64, Title_ string, Created_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("nid = ? and title = ? and created = ?", Nid_, Title_, Created_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByNidAndTitleAndChanged Get NodeFieldRevisions via NidAndTitleAndChanged
func GetNodeFieldRevisionsByNidAndTitleAndChanged(offset int, limit int, Nid_ int64, Title_ string, Changed_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("nid = ? and title = ? and changed = ?", Nid_, Title_, Changed_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByNidAndTitleAndPromote Get NodeFieldRevisions via NidAndTitleAndPromote
func GetNodeFieldRevisionsByNidAndTitleAndPromote(offset int, limit int, Nid_ int64, Title_ string, Promote_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("nid = ? and title = ? and promote = ?", Nid_, Title_, Promote_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByNidAndTitleAndSticky Get NodeFieldRevisions via NidAndTitleAndSticky
func GetNodeFieldRevisionsByNidAndTitleAndSticky(offset int, limit int, Nid_ int64, Title_ string, Sticky_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("nid = ? and title = ? and sticky = ?", Nid_, Title_, Sticky_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByNidAndTitleAndRevisionTranslationAffected Get NodeFieldRevisions via NidAndTitleAndRevisionTranslationAffected
func GetNodeFieldRevisionsByNidAndTitleAndRevisionTranslationAffected(offset int, limit int, Nid_ int64, Title_ string, RevisionTranslationAffected_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("nid = ? and title = ? and revision_translation_affected = ?", Nid_, Title_, RevisionTranslationAffected_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByNidAndTitleAndDefaultLangcode Get NodeFieldRevisions via NidAndTitleAndDefaultLangcode
func GetNodeFieldRevisionsByNidAndTitleAndDefaultLangcode(offset int, limit int, Nid_ int64, Title_ string, DefaultLangcode_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("nid = ? and title = ? and default_langcode = ?", Nid_, Title_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByNidAndUidAndStatus Get NodeFieldRevisions via NidAndUidAndStatus
func GetNodeFieldRevisionsByNidAndUidAndStatus(offset int, limit int, Nid_ int64, Uid_ int, Status_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("nid = ? and uid = ? and status = ?", Nid_, Uid_, Status_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByNidAndUidAndCreated Get NodeFieldRevisions via NidAndUidAndCreated
func GetNodeFieldRevisionsByNidAndUidAndCreated(offset int, limit int, Nid_ int64, Uid_ int, Created_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("nid = ? and uid = ? and created = ?", Nid_, Uid_, Created_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByNidAndUidAndChanged Get NodeFieldRevisions via NidAndUidAndChanged
func GetNodeFieldRevisionsByNidAndUidAndChanged(offset int, limit int, Nid_ int64, Uid_ int, Changed_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("nid = ? and uid = ? and changed = ?", Nid_, Uid_, Changed_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByNidAndUidAndPromote Get NodeFieldRevisions via NidAndUidAndPromote
func GetNodeFieldRevisionsByNidAndUidAndPromote(offset int, limit int, Nid_ int64, Uid_ int, Promote_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("nid = ? and uid = ? and promote = ?", Nid_, Uid_, Promote_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByNidAndUidAndSticky Get NodeFieldRevisions via NidAndUidAndSticky
func GetNodeFieldRevisionsByNidAndUidAndSticky(offset int, limit int, Nid_ int64, Uid_ int, Sticky_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("nid = ? and uid = ? and sticky = ?", Nid_, Uid_, Sticky_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByNidAndUidAndRevisionTranslationAffected Get NodeFieldRevisions via NidAndUidAndRevisionTranslationAffected
func GetNodeFieldRevisionsByNidAndUidAndRevisionTranslationAffected(offset int, limit int, Nid_ int64, Uid_ int, RevisionTranslationAffected_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("nid = ? and uid = ? and revision_translation_affected = ?", Nid_, Uid_, RevisionTranslationAffected_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByNidAndUidAndDefaultLangcode Get NodeFieldRevisions via NidAndUidAndDefaultLangcode
func GetNodeFieldRevisionsByNidAndUidAndDefaultLangcode(offset int, limit int, Nid_ int64, Uid_ int, DefaultLangcode_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("nid = ? and uid = ? and default_langcode = ?", Nid_, Uid_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByNidAndStatusAndCreated Get NodeFieldRevisions via NidAndStatusAndCreated
func GetNodeFieldRevisionsByNidAndStatusAndCreated(offset int, limit int, Nid_ int64, Status_ int, Created_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("nid = ? and status = ? and created = ?", Nid_, Status_, Created_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByNidAndStatusAndChanged Get NodeFieldRevisions via NidAndStatusAndChanged
func GetNodeFieldRevisionsByNidAndStatusAndChanged(offset int, limit int, Nid_ int64, Status_ int, Changed_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("nid = ? and status = ? and changed = ?", Nid_, Status_, Changed_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByNidAndStatusAndPromote Get NodeFieldRevisions via NidAndStatusAndPromote
func GetNodeFieldRevisionsByNidAndStatusAndPromote(offset int, limit int, Nid_ int64, Status_ int, Promote_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("nid = ? and status = ? and promote = ?", Nid_, Status_, Promote_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByNidAndStatusAndSticky Get NodeFieldRevisions via NidAndStatusAndSticky
func GetNodeFieldRevisionsByNidAndStatusAndSticky(offset int, limit int, Nid_ int64, Status_ int, Sticky_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("nid = ? and status = ? and sticky = ?", Nid_, Status_, Sticky_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByNidAndStatusAndRevisionTranslationAffected Get NodeFieldRevisions via NidAndStatusAndRevisionTranslationAffected
func GetNodeFieldRevisionsByNidAndStatusAndRevisionTranslationAffected(offset int, limit int, Nid_ int64, Status_ int, RevisionTranslationAffected_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("nid = ? and status = ? and revision_translation_affected = ?", Nid_, Status_, RevisionTranslationAffected_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByNidAndStatusAndDefaultLangcode Get NodeFieldRevisions via NidAndStatusAndDefaultLangcode
func GetNodeFieldRevisionsByNidAndStatusAndDefaultLangcode(offset int, limit int, Nid_ int64, Status_ int, DefaultLangcode_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("nid = ? and status = ? and default_langcode = ?", Nid_, Status_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByNidAndCreatedAndChanged Get NodeFieldRevisions via NidAndCreatedAndChanged
func GetNodeFieldRevisionsByNidAndCreatedAndChanged(offset int, limit int, Nid_ int64, Created_ int, Changed_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("nid = ? and created = ? and changed = ?", Nid_, Created_, Changed_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByNidAndCreatedAndPromote Get NodeFieldRevisions via NidAndCreatedAndPromote
func GetNodeFieldRevisionsByNidAndCreatedAndPromote(offset int, limit int, Nid_ int64, Created_ int, Promote_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("nid = ? and created = ? and promote = ?", Nid_, Created_, Promote_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByNidAndCreatedAndSticky Get NodeFieldRevisions via NidAndCreatedAndSticky
func GetNodeFieldRevisionsByNidAndCreatedAndSticky(offset int, limit int, Nid_ int64, Created_ int, Sticky_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("nid = ? and created = ? and sticky = ?", Nid_, Created_, Sticky_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByNidAndCreatedAndRevisionTranslationAffected Get NodeFieldRevisions via NidAndCreatedAndRevisionTranslationAffected
func GetNodeFieldRevisionsByNidAndCreatedAndRevisionTranslationAffected(offset int, limit int, Nid_ int64, Created_ int, RevisionTranslationAffected_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("nid = ? and created = ? and revision_translation_affected = ?", Nid_, Created_, RevisionTranslationAffected_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByNidAndCreatedAndDefaultLangcode Get NodeFieldRevisions via NidAndCreatedAndDefaultLangcode
func GetNodeFieldRevisionsByNidAndCreatedAndDefaultLangcode(offset int, limit int, Nid_ int64, Created_ int, DefaultLangcode_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("nid = ? and created = ? and default_langcode = ?", Nid_, Created_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByNidAndChangedAndPromote Get NodeFieldRevisions via NidAndChangedAndPromote
func GetNodeFieldRevisionsByNidAndChangedAndPromote(offset int, limit int, Nid_ int64, Changed_ int, Promote_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("nid = ? and changed = ? and promote = ?", Nid_, Changed_, Promote_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByNidAndChangedAndSticky Get NodeFieldRevisions via NidAndChangedAndSticky
func GetNodeFieldRevisionsByNidAndChangedAndSticky(offset int, limit int, Nid_ int64, Changed_ int, Sticky_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("nid = ? and changed = ? and sticky = ?", Nid_, Changed_, Sticky_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByNidAndChangedAndRevisionTranslationAffected Get NodeFieldRevisions via NidAndChangedAndRevisionTranslationAffected
func GetNodeFieldRevisionsByNidAndChangedAndRevisionTranslationAffected(offset int, limit int, Nid_ int64, Changed_ int, RevisionTranslationAffected_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("nid = ? and changed = ? and revision_translation_affected = ?", Nid_, Changed_, RevisionTranslationAffected_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByNidAndChangedAndDefaultLangcode Get NodeFieldRevisions via NidAndChangedAndDefaultLangcode
func GetNodeFieldRevisionsByNidAndChangedAndDefaultLangcode(offset int, limit int, Nid_ int64, Changed_ int, DefaultLangcode_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("nid = ? and changed = ? and default_langcode = ?", Nid_, Changed_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByNidAndPromoteAndSticky Get NodeFieldRevisions via NidAndPromoteAndSticky
func GetNodeFieldRevisionsByNidAndPromoteAndSticky(offset int, limit int, Nid_ int64, Promote_ int, Sticky_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("nid = ? and promote = ? and sticky = ?", Nid_, Promote_, Sticky_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByNidAndPromoteAndRevisionTranslationAffected Get NodeFieldRevisions via NidAndPromoteAndRevisionTranslationAffected
func GetNodeFieldRevisionsByNidAndPromoteAndRevisionTranslationAffected(offset int, limit int, Nid_ int64, Promote_ int, RevisionTranslationAffected_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("nid = ? and promote = ? and revision_translation_affected = ?", Nid_, Promote_, RevisionTranslationAffected_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByNidAndPromoteAndDefaultLangcode Get NodeFieldRevisions via NidAndPromoteAndDefaultLangcode
func GetNodeFieldRevisionsByNidAndPromoteAndDefaultLangcode(offset int, limit int, Nid_ int64, Promote_ int, DefaultLangcode_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("nid = ? and promote = ? and default_langcode = ?", Nid_, Promote_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByNidAndStickyAndRevisionTranslationAffected Get NodeFieldRevisions via NidAndStickyAndRevisionTranslationAffected
func GetNodeFieldRevisionsByNidAndStickyAndRevisionTranslationAffected(offset int, limit int, Nid_ int64, Sticky_ int, RevisionTranslationAffected_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("nid = ? and sticky = ? and revision_translation_affected = ?", Nid_, Sticky_, RevisionTranslationAffected_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByNidAndStickyAndDefaultLangcode Get NodeFieldRevisions via NidAndStickyAndDefaultLangcode
func GetNodeFieldRevisionsByNidAndStickyAndDefaultLangcode(offset int, limit int, Nid_ int64, Sticky_ int, DefaultLangcode_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("nid = ? and sticky = ? and default_langcode = ?", Nid_, Sticky_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByNidAndRevisionTranslationAffectedAndDefaultLangcode Get NodeFieldRevisions via NidAndRevisionTranslationAffectedAndDefaultLangcode
func GetNodeFieldRevisionsByNidAndRevisionTranslationAffectedAndDefaultLangcode(offset int, limit int, Nid_ int64, RevisionTranslationAffected_ int, DefaultLangcode_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("nid = ? and revision_translation_affected = ? and default_langcode = ?", Nid_, RevisionTranslationAffected_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByVidAndLangcodeAndTitle Get NodeFieldRevisions via VidAndLangcodeAndTitle
func GetNodeFieldRevisionsByVidAndLangcodeAndTitle(offset int, limit int, Vid_ int, Langcode_ string, Title_ string) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("vid = ? and langcode = ? and title = ?", Vid_, Langcode_, Title_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByVidAndLangcodeAndUid Get NodeFieldRevisions via VidAndLangcodeAndUid
func GetNodeFieldRevisionsByVidAndLangcodeAndUid(offset int, limit int, Vid_ int, Langcode_ string, Uid_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("vid = ? and langcode = ? and uid = ?", Vid_, Langcode_, Uid_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByVidAndLangcodeAndStatus Get NodeFieldRevisions via VidAndLangcodeAndStatus
func GetNodeFieldRevisionsByVidAndLangcodeAndStatus(offset int, limit int, Vid_ int, Langcode_ string, Status_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("vid = ? and langcode = ? and status = ?", Vid_, Langcode_, Status_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByVidAndLangcodeAndCreated Get NodeFieldRevisions via VidAndLangcodeAndCreated
func GetNodeFieldRevisionsByVidAndLangcodeAndCreated(offset int, limit int, Vid_ int, Langcode_ string, Created_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("vid = ? and langcode = ? and created = ?", Vid_, Langcode_, Created_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByVidAndLangcodeAndChanged Get NodeFieldRevisions via VidAndLangcodeAndChanged
func GetNodeFieldRevisionsByVidAndLangcodeAndChanged(offset int, limit int, Vid_ int, Langcode_ string, Changed_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("vid = ? and langcode = ? and changed = ?", Vid_, Langcode_, Changed_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByVidAndLangcodeAndPromote Get NodeFieldRevisions via VidAndLangcodeAndPromote
func GetNodeFieldRevisionsByVidAndLangcodeAndPromote(offset int, limit int, Vid_ int, Langcode_ string, Promote_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("vid = ? and langcode = ? and promote = ?", Vid_, Langcode_, Promote_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByVidAndLangcodeAndSticky Get NodeFieldRevisions via VidAndLangcodeAndSticky
func GetNodeFieldRevisionsByVidAndLangcodeAndSticky(offset int, limit int, Vid_ int, Langcode_ string, Sticky_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("vid = ? and langcode = ? and sticky = ?", Vid_, Langcode_, Sticky_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByVidAndLangcodeAndRevisionTranslationAffected Get NodeFieldRevisions via VidAndLangcodeAndRevisionTranslationAffected
func GetNodeFieldRevisionsByVidAndLangcodeAndRevisionTranslationAffected(offset int, limit int, Vid_ int, Langcode_ string, RevisionTranslationAffected_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("vid = ? and langcode = ? and revision_translation_affected = ?", Vid_, Langcode_, RevisionTranslationAffected_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByVidAndLangcodeAndDefaultLangcode Get NodeFieldRevisions via VidAndLangcodeAndDefaultLangcode
func GetNodeFieldRevisionsByVidAndLangcodeAndDefaultLangcode(offset int, limit int, Vid_ int, Langcode_ string, DefaultLangcode_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("vid = ? and langcode = ? and default_langcode = ?", Vid_, Langcode_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByVidAndTitleAndUid Get NodeFieldRevisions via VidAndTitleAndUid
func GetNodeFieldRevisionsByVidAndTitleAndUid(offset int, limit int, Vid_ int, Title_ string, Uid_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("vid = ? and title = ? and uid = ?", Vid_, Title_, Uid_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByVidAndTitleAndStatus Get NodeFieldRevisions via VidAndTitleAndStatus
func GetNodeFieldRevisionsByVidAndTitleAndStatus(offset int, limit int, Vid_ int, Title_ string, Status_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("vid = ? and title = ? and status = ?", Vid_, Title_, Status_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByVidAndTitleAndCreated Get NodeFieldRevisions via VidAndTitleAndCreated
func GetNodeFieldRevisionsByVidAndTitleAndCreated(offset int, limit int, Vid_ int, Title_ string, Created_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("vid = ? and title = ? and created = ?", Vid_, Title_, Created_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByVidAndTitleAndChanged Get NodeFieldRevisions via VidAndTitleAndChanged
func GetNodeFieldRevisionsByVidAndTitleAndChanged(offset int, limit int, Vid_ int, Title_ string, Changed_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("vid = ? and title = ? and changed = ?", Vid_, Title_, Changed_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByVidAndTitleAndPromote Get NodeFieldRevisions via VidAndTitleAndPromote
func GetNodeFieldRevisionsByVidAndTitleAndPromote(offset int, limit int, Vid_ int, Title_ string, Promote_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("vid = ? and title = ? and promote = ?", Vid_, Title_, Promote_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByVidAndTitleAndSticky Get NodeFieldRevisions via VidAndTitleAndSticky
func GetNodeFieldRevisionsByVidAndTitleAndSticky(offset int, limit int, Vid_ int, Title_ string, Sticky_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("vid = ? and title = ? and sticky = ?", Vid_, Title_, Sticky_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByVidAndTitleAndRevisionTranslationAffected Get NodeFieldRevisions via VidAndTitleAndRevisionTranslationAffected
func GetNodeFieldRevisionsByVidAndTitleAndRevisionTranslationAffected(offset int, limit int, Vid_ int, Title_ string, RevisionTranslationAffected_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("vid = ? and title = ? and revision_translation_affected = ?", Vid_, Title_, RevisionTranslationAffected_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByVidAndTitleAndDefaultLangcode Get NodeFieldRevisions via VidAndTitleAndDefaultLangcode
func GetNodeFieldRevisionsByVidAndTitleAndDefaultLangcode(offset int, limit int, Vid_ int, Title_ string, DefaultLangcode_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("vid = ? and title = ? and default_langcode = ?", Vid_, Title_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByVidAndUidAndStatus Get NodeFieldRevisions via VidAndUidAndStatus
func GetNodeFieldRevisionsByVidAndUidAndStatus(offset int, limit int, Vid_ int, Uid_ int, Status_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("vid = ? and uid = ? and status = ?", Vid_, Uid_, Status_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByVidAndUidAndCreated Get NodeFieldRevisions via VidAndUidAndCreated
func GetNodeFieldRevisionsByVidAndUidAndCreated(offset int, limit int, Vid_ int, Uid_ int, Created_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("vid = ? and uid = ? and created = ?", Vid_, Uid_, Created_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByVidAndUidAndChanged Get NodeFieldRevisions via VidAndUidAndChanged
func GetNodeFieldRevisionsByVidAndUidAndChanged(offset int, limit int, Vid_ int, Uid_ int, Changed_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("vid = ? and uid = ? and changed = ?", Vid_, Uid_, Changed_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByVidAndUidAndPromote Get NodeFieldRevisions via VidAndUidAndPromote
func GetNodeFieldRevisionsByVidAndUidAndPromote(offset int, limit int, Vid_ int, Uid_ int, Promote_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("vid = ? and uid = ? and promote = ?", Vid_, Uid_, Promote_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByVidAndUidAndSticky Get NodeFieldRevisions via VidAndUidAndSticky
func GetNodeFieldRevisionsByVidAndUidAndSticky(offset int, limit int, Vid_ int, Uid_ int, Sticky_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("vid = ? and uid = ? and sticky = ?", Vid_, Uid_, Sticky_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByVidAndUidAndRevisionTranslationAffected Get NodeFieldRevisions via VidAndUidAndRevisionTranslationAffected
func GetNodeFieldRevisionsByVidAndUidAndRevisionTranslationAffected(offset int, limit int, Vid_ int, Uid_ int, RevisionTranslationAffected_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("vid = ? and uid = ? and revision_translation_affected = ?", Vid_, Uid_, RevisionTranslationAffected_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByVidAndUidAndDefaultLangcode Get NodeFieldRevisions via VidAndUidAndDefaultLangcode
func GetNodeFieldRevisionsByVidAndUidAndDefaultLangcode(offset int, limit int, Vid_ int, Uid_ int, DefaultLangcode_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("vid = ? and uid = ? and default_langcode = ?", Vid_, Uid_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByVidAndStatusAndCreated Get NodeFieldRevisions via VidAndStatusAndCreated
func GetNodeFieldRevisionsByVidAndStatusAndCreated(offset int, limit int, Vid_ int, Status_ int, Created_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("vid = ? and status = ? and created = ?", Vid_, Status_, Created_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByVidAndStatusAndChanged Get NodeFieldRevisions via VidAndStatusAndChanged
func GetNodeFieldRevisionsByVidAndStatusAndChanged(offset int, limit int, Vid_ int, Status_ int, Changed_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("vid = ? and status = ? and changed = ?", Vid_, Status_, Changed_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByVidAndStatusAndPromote Get NodeFieldRevisions via VidAndStatusAndPromote
func GetNodeFieldRevisionsByVidAndStatusAndPromote(offset int, limit int, Vid_ int, Status_ int, Promote_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("vid = ? and status = ? and promote = ?", Vid_, Status_, Promote_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByVidAndStatusAndSticky Get NodeFieldRevisions via VidAndStatusAndSticky
func GetNodeFieldRevisionsByVidAndStatusAndSticky(offset int, limit int, Vid_ int, Status_ int, Sticky_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("vid = ? and status = ? and sticky = ?", Vid_, Status_, Sticky_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByVidAndStatusAndRevisionTranslationAffected Get NodeFieldRevisions via VidAndStatusAndRevisionTranslationAffected
func GetNodeFieldRevisionsByVidAndStatusAndRevisionTranslationAffected(offset int, limit int, Vid_ int, Status_ int, RevisionTranslationAffected_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("vid = ? and status = ? and revision_translation_affected = ?", Vid_, Status_, RevisionTranslationAffected_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByVidAndStatusAndDefaultLangcode Get NodeFieldRevisions via VidAndStatusAndDefaultLangcode
func GetNodeFieldRevisionsByVidAndStatusAndDefaultLangcode(offset int, limit int, Vid_ int, Status_ int, DefaultLangcode_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("vid = ? and status = ? and default_langcode = ?", Vid_, Status_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByVidAndCreatedAndChanged Get NodeFieldRevisions via VidAndCreatedAndChanged
func GetNodeFieldRevisionsByVidAndCreatedAndChanged(offset int, limit int, Vid_ int, Created_ int, Changed_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("vid = ? and created = ? and changed = ?", Vid_, Created_, Changed_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByVidAndCreatedAndPromote Get NodeFieldRevisions via VidAndCreatedAndPromote
func GetNodeFieldRevisionsByVidAndCreatedAndPromote(offset int, limit int, Vid_ int, Created_ int, Promote_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("vid = ? and created = ? and promote = ?", Vid_, Created_, Promote_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByVidAndCreatedAndSticky Get NodeFieldRevisions via VidAndCreatedAndSticky
func GetNodeFieldRevisionsByVidAndCreatedAndSticky(offset int, limit int, Vid_ int, Created_ int, Sticky_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("vid = ? and created = ? and sticky = ?", Vid_, Created_, Sticky_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByVidAndCreatedAndRevisionTranslationAffected Get NodeFieldRevisions via VidAndCreatedAndRevisionTranslationAffected
func GetNodeFieldRevisionsByVidAndCreatedAndRevisionTranslationAffected(offset int, limit int, Vid_ int, Created_ int, RevisionTranslationAffected_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("vid = ? and created = ? and revision_translation_affected = ?", Vid_, Created_, RevisionTranslationAffected_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByVidAndCreatedAndDefaultLangcode Get NodeFieldRevisions via VidAndCreatedAndDefaultLangcode
func GetNodeFieldRevisionsByVidAndCreatedAndDefaultLangcode(offset int, limit int, Vid_ int, Created_ int, DefaultLangcode_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("vid = ? and created = ? and default_langcode = ?", Vid_, Created_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByVidAndChangedAndPromote Get NodeFieldRevisions via VidAndChangedAndPromote
func GetNodeFieldRevisionsByVidAndChangedAndPromote(offset int, limit int, Vid_ int, Changed_ int, Promote_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("vid = ? and changed = ? and promote = ?", Vid_, Changed_, Promote_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByVidAndChangedAndSticky Get NodeFieldRevisions via VidAndChangedAndSticky
func GetNodeFieldRevisionsByVidAndChangedAndSticky(offset int, limit int, Vid_ int, Changed_ int, Sticky_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("vid = ? and changed = ? and sticky = ?", Vid_, Changed_, Sticky_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByVidAndChangedAndRevisionTranslationAffected Get NodeFieldRevisions via VidAndChangedAndRevisionTranslationAffected
func GetNodeFieldRevisionsByVidAndChangedAndRevisionTranslationAffected(offset int, limit int, Vid_ int, Changed_ int, RevisionTranslationAffected_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("vid = ? and changed = ? and revision_translation_affected = ?", Vid_, Changed_, RevisionTranslationAffected_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByVidAndChangedAndDefaultLangcode Get NodeFieldRevisions via VidAndChangedAndDefaultLangcode
func GetNodeFieldRevisionsByVidAndChangedAndDefaultLangcode(offset int, limit int, Vid_ int, Changed_ int, DefaultLangcode_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("vid = ? and changed = ? and default_langcode = ?", Vid_, Changed_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByVidAndPromoteAndSticky Get NodeFieldRevisions via VidAndPromoteAndSticky
func GetNodeFieldRevisionsByVidAndPromoteAndSticky(offset int, limit int, Vid_ int, Promote_ int, Sticky_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("vid = ? and promote = ? and sticky = ?", Vid_, Promote_, Sticky_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByVidAndPromoteAndRevisionTranslationAffected Get NodeFieldRevisions via VidAndPromoteAndRevisionTranslationAffected
func GetNodeFieldRevisionsByVidAndPromoteAndRevisionTranslationAffected(offset int, limit int, Vid_ int, Promote_ int, RevisionTranslationAffected_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("vid = ? and promote = ? and revision_translation_affected = ?", Vid_, Promote_, RevisionTranslationAffected_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByVidAndPromoteAndDefaultLangcode Get NodeFieldRevisions via VidAndPromoteAndDefaultLangcode
func GetNodeFieldRevisionsByVidAndPromoteAndDefaultLangcode(offset int, limit int, Vid_ int, Promote_ int, DefaultLangcode_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("vid = ? and promote = ? and default_langcode = ?", Vid_, Promote_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByVidAndStickyAndRevisionTranslationAffected Get NodeFieldRevisions via VidAndStickyAndRevisionTranslationAffected
func GetNodeFieldRevisionsByVidAndStickyAndRevisionTranslationAffected(offset int, limit int, Vid_ int, Sticky_ int, RevisionTranslationAffected_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("vid = ? and sticky = ? and revision_translation_affected = ?", Vid_, Sticky_, RevisionTranslationAffected_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByVidAndStickyAndDefaultLangcode Get NodeFieldRevisions via VidAndStickyAndDefaultLangcode
func GetNodeFieldRevisionsByVidAndStickyAndDefaultLangcode(offset int, limit int, Vid_ int, Sticky_ int, DefaultLangcode_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("vid = ? and sticky = ? and default_langcode = ?", Vid_, Sticky_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByVidAndRevisionTranslationAffectedAndDefaultLangcode Get NodeFieldRevisions via VidAndRevisionTranslationAffectedAndDefaultLangcode
func GetNodeFieldRevisionsByVidAndRevisionTranslationAffectedAndDefaultLangcode(offset int, limit int, Vid_ int, RevisionTranslationAffected_ int, DefaultLangcode_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("vid = ? and revision_translation_affected = ? and default_langcode = ?", Vid_, RevisionTranslationAffected_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByLangcodeAndTitleAndUid Get NodeFieldRevisions via LangcodeAndTitleAndUid
func GetNodeFieldRevisionsByLangcodeAndTitleAndUid(offset int, limit int, Langcode_ string, Title_ string, Uid_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("langcode = ? and title = ? and uid = ?", Langcode_, Title_, Uid_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByLangcodeAndTitleAndStatus Get NodeFieldRevisions via LangcodeAndTitleAndStatus
func GetNodeFieldRevisionsByLangcodeAndTitleAndStatus(offset int, limit int, Langcode_ string, Title_ string, Status_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("langcode = ? and title = ? and status = ?", Langcode_, Title_, Status_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByLangcodeAndTitleAndCreated Get NodeFieldRevisions via LangcodeAndTitleAndCreated
func GetNodeFieldRevisionsByLangcodeAndTitleAndCreated(offset int, limit int, Langcode_ string, Title_ string, Created_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("langcode = ? and title = ? and created = ?", Langcode_, Title_, Created_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByLangcodeAndTitleAndChanged Get NodeFieldRevisions via LangcodeAndTitleAndChanged
func GetNodeFieldRevisionsByLangcodeAndTitleAndChanged(offset int, limit int, Langcode_ string, Title_ string, Changed_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("langcode = ? and title = ? and changed = ?", Langcode_, Title_, Changed_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByLangcodeAndTitleAndPromote Get NodeFieldRevisions via LangcodeAndTitleAndPromote
func GetNodeFieldRevisionsByLangcodeAndTitleAndPromote(offset int, limit int, Langcode_ string, Title_ string, Promote_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("langcode = ? and title = ? and promote = ?", Langcode_, Title_, Promote_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByLangcodeAndTitleAndSticky Get NodeFieldRevisions via LangcodeAndTitleAndSticky
func GetNodeFieldRevisionsByLangcodeAndTitleAndSticky(offset int, limit int, Langcode_ string, Title_ string, Sticky_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("langcode = ? and title = ? and sticky = ?", Langcode_, Title_, Sticky_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByLangcodeAndTitleAndRevisionTranslationAffected Get NodeFieldRevisions via LangcodeAndTitleAndRevisionTranslationAffected
func GetNodeFieldRevisionsByLangcodeAndTitleAndRevisionTranslationAffected(offset int, limit int, Langcode_ string, Title_ string, RevisionTranslationAffected_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("langcode = ? and title = ? and revision_translation_affected = ?", Langcode_, Title_, RevisionTranslationAffected_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByLangcodeAndTitleAndDefaultLangcode Get NodeFieldRevisions via LangcodeAndTitleAndDefaultLangcode
func GetNodeFieldRevisionsByLangcodeAndTitleAndDefaultLangcode(offset int, limit int, Langcode_ string, Title_ string, DefaultLangcode_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("langcode = ? and title = ? and default_langcode = ?", Langcode_, Title_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByLangcodeAndUidAndStatus Get NodeFieldRevisions via LangcodeAndUidAndStatus
func GetNodeFieldRevisionsByLangcodeAndUidAndStatus(offset int, limit int, Langcode_ string, Uid_ int, Status_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("langcode = ? and uid = ? and status = ?", Langcode_, Uid_, Status_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByLangcodeAndUidAndCreated Get NodeFieldRevisions via LangcodeAndUidAndCreated
func GetNodeFieldRevisionsByLangcodeAndUidAndCreated(offset int, limit int, Langcode_ string, Uid_ int, Created_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("langcode = ? and uid = ? and created = ?", Langcode_, Uid_, Created_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByLangcodeAndUidAndChanged Get NodeFieldRevisions via LangcodeAndUidAndChanged
func GetNodeFieldRevisionsByLangcodeAndUidAndChanged(offset int, limit int, Langcode_ string, Uid_ int, Changed_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("langcode = ? and uid = ? and changed = ?", Langcode_, Uid_, Changed_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByLangcodeAndUidAndPromote Get NodeFieldRevisions via LangcodeAndUidAndPromote
func GetNodeFieldRevisionsByLangcodeAndUidAndPromote(offset int, limit int, Langcode_ string, Uid_ int, Promote_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("langcode = ? and uid = ? and promote = ?", Langcode_, Uid_, Promote_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByLangcodeAndUidAndSticky Get NodeFieldRevisions via LangcodeAndUidAndSticky
func GetNodeFieldRevisionsByLangcodeAndUidAndSticky(offset int, limit int, Langcode_ string, Uid_ int, Sticky_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("langcode = ? and uid = ? and sticky = ?", Langcode_, Uid_, Sticky_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByLangcodeAndUidAndRevisionTranslationAffected Get NodeFieldRevisions via LangcodeAndUidAndRevisionTranslationAffected
func GetNodeFieldRevisionsByLangcodeAndUidAndRevisionTranslationAffected(offset int, limit int, Langcode_ string, Uid_ int, RevisionTranslationAffected_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("langcode = ? and uid = ? and revision_translation_affected = ?", Langcode_, Uid_, RevisionTranslationAffected_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByLangcodeAndUidAndDefaultLangcode Get NodeFieldRevisions via LangcodeAndUidAndDefaultLangcode
func GetNodeFieldRevisionsByLangcodeAndUidAndDefaultLangcode(offset int, limit int, Langcode_ string, Uid_ int, DefaultLangcode_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("langcode = ? and uid = ? and default_langcode = ?", Langcode_, Uid_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByLangcodeAndStatusAndCreated Get NodeFieldRevisions via LangcodeAndStatusAndCreated
func GetNodeFieldRevisionsByLangcodeAndStatusAndCreated(offset int, limit int, Langcode_ string, Status_ int, Created_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("langcode = ? and status = ? and created = ?", Langcode_, Status_, Created_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByLangcodeAndStatusAndChanged Get NodeFieldRevisions via LangcodeAndStatusAndChanged
func GetNodeFieldRevisionsByLangcodeAndStatusAndChanged(offset int, limit int, Langcode_ string, Status_ int, Changed_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("langcode = ? and status = ? and changed = ?", Langcode_, Status_, Changed_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByLangcodeAndStatusAndPromote Get NodeFieldRevisions via LangcodeAndStatusAndPromote
func GetNodeFieldRevisionsByLangcodeAndStatusAndPromote(offset int, limit int, Langcode_ string, Status_ int, Promote_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("langcode = ? and status = ? and promote = ?", Langcode_, Status_, Promote_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByLangcodeAndStatusAndSticky Get NodeFieldRevisions via LangcodeAndStatusAndSticky
func GetNodeFieldRevisionsByLangcodeAndStatusAndSticky(offset int, limit int, Langcode_ string, Status_ int, Sticky_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("langcode = ? and status = ? and sticky = ?", Langcode_, Status_, Sticky_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByLangcodeAndStatusAndRevisionTranslationAffected Get NodeFieldRevisions via LangcodeAndStatusAndRevisionTranslationAffected
func GetNodeFieldRevisionsByLangcodeAndStatusAndRevisionTranslationAffected(offset int, limit int, Langcode_ string, Status_ int, RevisionTranslationAffected_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("langcode = ? and status = ? and revision_translation_affected = ?", Langcode_, Status_, RevisionTranslationAffected_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByLangcodeAndStatusAndDefaultLangcode Get NodeFieldRevisions via LangcodeAndStatusAndDefaultLangcode
func GetNodeFieldRevisionsByLangcodeAndStatusAndDefaultLangcode(offset int, limit int, Langcode_ string, Status_ int, DefaultLangcode_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("langcode = ? and status = ? and default_langcode = ?", Langcode_, Status_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByLangcodeAndCreatedAndChanged Get NodeFieldRevisions via LangcodeAndCreatedAndChanged
func GetNodeFieldRevisionsByLangcodeAndCreatedAndChanged(offset int, limit int, Langcode_ string, Created_ int, Changed_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("langcode = ? and created = ? and changed = ?", Langcode_, Created_, Changed_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByLangcodeAndCreatedAndPromote Get NodeFieldRevisions via LangcodeAndCreatedAndPromote
func GetNodeFieldRevisionsByLangcodeAndCreatedAndPromote(offset int, limit int, Langcode_ string, Created_ int, Promote_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("langcode = ? and created = ? and promote = ?", Langcode_, Created_, Promote_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByLangcodeAndCreatedAndSticky Get NodeFieldRevisions via LangcodeAndCreatedAndSticky
func GetNodeFieldRevisionsByLangcodeAndCreatedAndSticky(offset int, limit int, Langcode_ string, Created_ int, Sticky_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("langcode = ? and created = ? and sticky = ?", Langcode_, Created_, Sticky_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByLangcodeAndCreatedAndRevisionTranslationAffected Get NodeFieldRevisions via LangcodeAndCreatedAndRevisionTranslationAffected
func GetNodeFieldRevisionsByLangcodeAndCreatedAndRevisionTranslationAffected(offset int, limit int, Langcode_ string, Created_ int, RevisionTranslationAffected_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("langcode = ? and created = ? and revision_translation_affected = ?", Langcode_, Created_, RevisionTranslationAffected_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByLangcodeAndCreatedAndDefaultLangcode Get NodeFieldRevisions via LangcodeAndCreatedAndDefaultLangcode
func GetNodeFieldRevisionsByLangcodeAndCreatedAndDefaultLangcode(offset int, limit int, Langcode_ string, Created_ int, DefaultLangcode_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("langcode = ? and created = ? and default_langcode = ?", Langcode_, Created_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByLangcodeAndChangedAndPromote Get NodeFieldRevisions via LangcodeAndChangedAndPromote
func GetNodeFieldRevisionsByLangcodeAndChangedAndPromote(offset int, limit int, Langcode_ string, Changed_ int, Promote_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("langcode = ? and changed = ? and promote = ?", Langcode_, Changed_, Promote_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByLangcodeAndChangedAndSticky Get NodeFieldRevisions via LangcodeAndChangedAndSticky
func GetNodeFieldRevisionsByLangcodeAndChangedAndSticky(offset int, limit int, Langcode_ string, Changed_ int, Sticky_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("langcode = ? and changed = ? and sticky = ?", Langcode_, Changed_, Sticky_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByLangcodeAndChangedAndRevisionTranslationAffected Get NodeFieldRevisions via LangcodeAndChangedAndRevisionTranslationAffected
func GetNodeFieldRevisionsByLangcodeAndChangedAndRevisionTranslationAffected(offset int, limit int, Langcode_ string, Changed_ int, RevisionTranslationAffected_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("langcode = ? and changed = ? and revision_translation_affected = ?", Langcode_, Changed_, RevisionTranslationAffected_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByLangcodeAndChangedAndDefaultLangcode Get NodeFieldRevisions via LangcodeAndChangedAndDefaultLangcode
func GetNodeFieldRevisionsByLangcodeAndChangedAndDefaultLangcode(offset int, limit int, Langcode_ string, Changed_ int, DefaultLangcode_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("langcode = ? and changed = ? and default_langcode = ?", Langcode_, Changed_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByLangcodeAndPromoteAndSticky Get NodeFieldRevisions via LangcodeAndPromoteAndSticky
func GetNodeFieldRevisionsByLangcodeAndPromoteAndSticky(offset int, limit int, Langcode_ string, Promote_ int, Sticky_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("langcode = ? and promote = ? and sticky = ?", Langcode_, Promote_, Sticky_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByLangcodeAndPromoteAndRevisionTranslationAffected Get NodeFieldRevisions via LangcodeAndPromoteAndRevisionTranslationAffected
func GetNodeFieldRevisionsByLangcodeAndPromoteAndRevisionTranslationAffected(offset int, limit int, Langcode_ string, Promote_ int, RevisionTranslationAffected_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("langcode = ? and promote = ? and revision_translation_affected = ?", Langcode_, Promote_, RevisionTranslationAffected_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByLangcodeAndPromoteAndDefaultLangcode Get NodeFieldRevisions via LangcodeAndPromoteAndDefaultLangcode
func GetNodeFieldRevisionsByLangcodeAndPromoteAndDefaultLangcode(offset int, limit int, Langcode_ string, Promote_ int, DefaultLangcode_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("langcode = ? and promote = ? and default_langcode = ?", Langcode_, Promote_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByLangcodeAndStickyAndRevisionTranslationAffected Get NodeFieldRevisions via LangcodeAndStickyAndRevisionTranslationAffected
func GetNodeFieldRevisionsByLangcodeAndStickyAndRevisionTranslationAffected(offset int, limit int, Langcode_ string, Sticky_ int, RevisionTranslationAffected_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("langcode = ? and sticky = ? and revision_translation_affected = ?", Langcode_, Sticky_, RevisionTranslationAffected_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByLangcodeAndStickyAndDefaultLangcode Get NodeFieldRevisions via LangcodeAndStickyAndDefaultLangcode
func GetNodeFieldRevisionsByLangcodeAndStickyAndDefaultLangcode(offset int, limit int, Langcode_ string, Sticky_ int, DefaultLangcode_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("langcode = ? and sticky = ? and default_langcode = ?", Langcode_, Sticky_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByLangcodeAndRevisionTranslationAffectedAndDefaultLangcode Get NodeFieldRevisions via LangcodeAndRevisionTranslationAffectedAndDefaultLangcode
func GetNodeFieldRevisionsByLangcodeAndRevisionTranslationAffectedAndDefaultLangcode(offset int, limit int, Langcode_ string, RevisionTranslationAffected_ int, DefaultLangcode_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("langcode = ? and revision_translation_affected = ? and default_langcode = ?", Langcode_, RevisionTranslationAffected_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByTitleAndUidAndStatus Get NodeFieldRevisions via TitleAndUidAndStatus
func GetNodeFieldRevisionsByTitleAndUidAndStatus(offset int, limit int, Title_ string, Uid_ int, Status_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("title = ? and uid = ? and status = ?", Title_, Uid_, Status_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByTitleAndUidAndCreated Get NodeFieldRevisions via TitleAndUidAndCreated
func GetNodeFieldRevisionsByTitleAndUidAndCreated(offset int, limit int, Title_ string, Uid_ int, Created_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("title = ? and uid = ? and created = ?", Title_, Uid_, Created_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByTitleAndUidAndChanged Get NodeFieldRevisions via TitleAndUidAndChanged
func GetNodeFieldRevisionsByTitleAndUidAndChanged(offset int, limit int, Title_ string, Uid_ int, Changed_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("title = ? and uid = ? and changed = ?", Title_, Uid_, Changed_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByTitleAndUidAndPromote Get NodeFieldRevisions via TitleAndUidAndPromote
func GetNodeFieldRevisionsByTitleAndUidAndPromote(offset int, limit int, Title_ string, Uid_ int, Promote_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("title = ? and uid = ? and promote = ?", Title_, Uid_, Promote_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByTitleAndUidAndSticky Get NodeFieldRevisions via TitleAndUidAndSticky
func GetNodeFieldRevisionsByTitleAndUidAndSticky(offset int, limit int, Title_ string, Uid_ int, Sticky_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("title = ? and uid = ? and sticky = ?", Title_, Uid_, Sticky_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByTitleAndUidAndRevisionTranslationAffected Get NodeFieldRevisions via TitleAndUidAndRevisionTranslationAffected
func GetNodeFieldRevisionsByTitleAndUidAndRevisionTranslationAffected(offset int, limit int, Title_ string, Uid_ int, RevisionTranslationAffected_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("title = ? and uid = ? and revision_translation_affected = ?", Title_, Uid_, RevisionTranslationAffected_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByTitleAndUidAndDefaultLangcode Get NodeFieldRevisions via TitleAndUidAndDefaultLangcode
func GetNodeFieldRevisionsByTitleAndUidAndDefaultLangcode(offset int, limit int, Title_ string, Uid_ int, DefaultLangcode_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("title = ? and uid = ? and default_langcode = ?", Title_, Uid_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByTitleAndStatusAndCreated Get NodeFieldRevisions via TitleAndStatusAndCreated
func GetNodeFieldRevisionsByTitleAndStatusAndCreated(offset int, limit int, Title_ string, Status_ int, Created_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("title = ? and status = ? and created = ?", Title_, Status_, Created_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByTitleAndStatusAndChanged Get NodeFieldRevisions via TitleAndStatusAndChanged
func GetNodeFieldRevisionsByTitleAndStatusAndChanged(offset int, limit int, Title_ string, Status_ int, Changed_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("title = ? and status = ? and changed = ?", Title_, Status_, Changed_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByTitleAndStatusAndPromote Get NodeFieldRevisions via TitleAndStatusAndPromote
func GetNodeFieldRevisionsByTitleAndStatusAndPromote(offset int, limit int, Title_ string, Status_ int, Promote_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("title = ? and status = ? and promote = ?", Title_, Status_, Promote_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByTitleAndStatusAndSticky Get NodeFieldRevisions via TitleAndStatusAndSticky
func GetNodeFieldRevisionsByTitleAndStatusAndSticky(offset int, limit int, Title_ string, Status_ int, Sticky_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("title = ? and status = ? and sticky = ?", Title_, Status_, Sticky_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByTitleAndStatusAndRevisionTranslationAffected Get NodeFieldRevisions via TitleAndStatusAndRevisionTranslationAffected
func GetNodeFieldRevisionsByTitleAndStatusAndRevisionTranslationAffected(offset int, limit int, Title_ string, Status_ int, RevisionTranslationAffected_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("title = ? and status = ? and revision_translation_affected = ?", Title_, Status_, RevisionTranslationAffected_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByTitleAndStatusAndDefaultLangcode Get NodeFieldRevisions via TitleAndStatusAndDefaultLangcode
func GetNodeFieldRevisionsByTitleAndStatusAndDefaultLangcode(offset int, limit int, Title_ string, Status_ int, DefaultLangcode_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("title = ? and status = ? and default_langcode = ?", Title_, Status_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByTitleAndCreatedAndChanged Get NodeFieldRevisions via TitleAndCreatedAndChanged
func GetNodeFieldRevisionsByTitleAndCreatedAndChanged(offset int, limit int, Title_ string, Created_ int, Changed_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("title = ? and created = ? and changed = ?", Title_, Created_, Changed_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByTitleAndCreatedAndPromote Get NodeFieldRevisions via TitleAndCreatedAndPromote
func GetNodeFieldRevisionsByTitleAndCreatedAndPromote(offset int, limit int, Title_ string, Created_ int, Promote_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("title = ? and created = ? and promote = ?", Title_, Created_, Promote_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByTitleAndCreatedAndSticky Get NodeFieldRevisions via TitleAndCreatedAndSticky
func GetNodeFieldRevisionsByTitleAndCreatedAndSticky(offset int, limit int, Title_ string, Created_ int, Sticky_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("title = ? and created = ? and sticky = ?", Title_, Created_, Sticky_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByTitleAndCreatedAndRevisionTranslationAffected Get NodeFieldRevisions via TitleAndCreatedAndRevisionTranslationAffected
func GetNodeFieldRevisionsByTitleAndCreatedAndRevisionTranslationAffected(offset int, limit int, Title_ string, Created_ int, RevisionTranslationAffected_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("title = ? and created = ? and revision_translation_affected = ?", Title_, Created_, RevisionTranslationAffected_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByTitleAndCreatedAndDefaultLangcode Get NodeFieldRevisions via TitleAndCreatedAndDefaultLangcode
func GetNodeFieldRevisionsByTitleAndCreatedAndDefaultLangcode(offset int, limit int, Title_ string, Created_ int, DefaultLangcode_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("title = ? and created = ? and default_langcode = ?", Title_, Created_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByTitleAndChangedAndPromote Get NodeFieldRevisions via TitleAndChangedAndPromote
func GetNodeFieldRevisionsByTitleAndChangedAndPromote(offset int, limit int, Title_ string, Changed_ int, Promote_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("title = ? and changed = ? and promote = ?", Title_, Changed_, Promote_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByTitleAndChangedAndSticky Get NodeFieldRevisions via TitleAndChangedAndSticky
func GetNodeFieldRevisionsByTitleAndChangedAndSticky(offset int, limit int, Title_ string, Changed_ int, Sticky_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("title = ? and changed = ? and sticky = ?", Title_, Changed_, Sticky_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByTitleAndChangedAndRevisionTranslationAffected Get NodeFieldRevisions via TitleAndChangedAndRevisionTranslationAffected
func GetNodeFieldRevisionsByTitleAndChangedAndRevisionTranslationAffected(offset int, limit int, Title_ string, Changed_ int, RevisionTranslationAffected_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("title = ? and changed = ? and revision_translation_affected = ?", Title_, Changed_, RevisionTranslationAffected_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByTitleAndChangedAndDefaultLangcode Get NodeFieldRevisions via TitleAndChangedAndDefaultLangcode
func GetNodeFieldRevisionsByTitleAndChangedAndDefaultLangcode(offset int, limit int, Title_ string, Changed_ int, DefaultLangcode_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("title = ? and changed = ? and default_langcode = ?", Title_, Changed_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByTitleAndPromoteAndSticky Get NodeFieldRevisions via TitleAndPromoteAndSticky
func GetNodeFieldRevisionsByTitleAndPromoteAndSticky(offset int, limit int, Title_ string, Promote_ int, Sticky_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("title = ? and promote = ? and sticky = ?", Title_, Promote_, Sticky_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByTitleAndPromoteAndRevisionTranslationAffected Get NodeFieldRevisions via TitleAndPromoteAndRevisionTranslationAffected
func GetNodeFieldRevisionsByTitleAndPromoteAndRevisionTranslationAffected(offset int, limit int, Title_ string, Promote_ int, RevisionTranslationAffected_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("title = ? and promote = ? and revision_translation_affected = ?", Title_, Promote_, RevisionTranslationAffected_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByTitleAndPromoteAndDefaultLangcode Get NodeFieldRevisions via TitleAndPromoteAndDefaultLangcode
func GetNodeFieldRevisionsByTitleAndPromoteAndDefaultLangcode(offset int, limit int, Title_ string, Promote_ int, DefaultLangcode_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("title = ? and promote = ? and default_langcode = ?", Title_, Promote_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByTitleAndStickyAndRevisionTranslationAffected Get NodeFieldRevisions via TitleAndStickyAndRevisionTranslationAffected
func GetNodeFieldRevisionsByTitleAndStickyAndRevisionTranslationAffected(offset int, limit int, Title_ string, Sticky_ int, RevisionTranslationAffected_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("title = ? and sticky = ? and revision_translation_affected = ?", Title_, Sticky_, RevisionTranslationAffected_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByTitleAndStickyAndDefaultLangcode Get NodeFieldRevisions via TitleAndStickyAndDefaultLangcode
func GetNodeFieldRevisionsByTitleAndStickyAndDefaultLangcode(offset int, limit int, Title_ string, Sticky_ int, DefaultLangcode_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("title = ? and sticky = ? and default_langcode = ?", Title_, Sticky_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByTitleAndRevisionTranslationAffectedAndDefaultLangcode Get NodeFieldRevisions via TitleAndRevisionTranslationAffectedAndDefaultLangcode
func GetNodeFieldRevisionsByTitleAndRevisionTranslationAffectedAndDefaultLangcode(offset int, limit int, Title_ string, RevisionTranslationAffected_ int, DefaultLangcode_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("title = ? and revision_translation_affected = ? and default_langcode = ?", Title_, RevisionTranslationAffected_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByUidAndStatusAndCreated Get NodeFieldRevisions via UidAndStatusAndCreated
func GetNodeFieldRevisionsByUidAndStatusAndCreated(offset int, limit int, Uid_ int, Status_ int, Created_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("uid = ? and status = ? and created = ?", Uid_, Status_, Created_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByUidAndStatusAndChanged Get NodeFieldRevisions via UidAndStatusAndChanged
func GetNodeFieldRevisionsByUidAndStatusAndChanged(offset int, limit int, Uid_ int, Status_ int, Changed_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("uid = ? and status = ? and changed = ?", Uid_, Status_, Changed_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByUidAndStatusAndPromote Get NodeFieldRevisions via UidAndStatusAndPromote
func GetNodeFieldRevisionsByUidAndStatusAndPromote(offset int, limit int, Uid_ int, Status_ int, Promote_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("uid = ? and status = ? and promote = ?", Uid_, Status_, Promote_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByUidAndStatusAndSticky Get NodeFieldRevisions via UidAndStatusAndSticky
func GetNodeFieldRevisionsByUidAndStatusAndSticky(offset int, limit int, Uid_ int, Status_ int, Sticky_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("uid = ? and status = ? and sticky = ?", Uid_, Status_, Sticky_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByUidAndStatusAndRevisionTranslationAffected Get NodeFieldRevisions via UidAndStatusAndRevisionTranslationAffected
func GetNodeFieldRevisionsByUidAndStatusAndRevisionTranslationAffected(offset int, limit int, Uid_ int, Status_ int, RevisionTranslationAffected_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("uid = ? and status = ? and revision_translation_affected = ?", Uid_, Status_, RevisionTranslationAffected_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByUidAndStatusAndDefaultLangcode Get NodeFieldRevisions via UidAndStatusAndDefaultLangcode
func GetNodeFieldRevisionsByUidAndStatusAndDefaultLangcode(offset int, limit int, Uid_ int, Status_ int, DefaultLangcode_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("uid = ? and status = ? and default_langcode = ?", Uid_, Status_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByUidAndCreatedAndChanged Get NodeFieldRevisions via UidAndCreatedAndChanged
func GetNodeFieldRevisionsByUidAndCreatedAndChanged(offset int, limit int, Uid_ int, Created_ int, Changed_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("uid = ? and created = ? and changed = ?", Uid_, Created_, Changed_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByUidAndCreatedAndPromote Get NodeFieldRevisions via UidAndCreatedAndPromote
func GetNodeFieldRevisionsByUidAndCreatedAndPromote(offset int, limit int, Uid_ int, Created_ int, Promote_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("uid = ? and created = ? and promote = ?", Uid_, Created_, Promote_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByUidAndCreatedAndSticky Get NodeFieldRevisions via UidAndCreatedAndSticky
func GetNodeFieldRevisionsByUidAndCreatedAndSticky(offset int, limit int, Uid_ int, Created_ int, Sticky_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("uid = ? and created = ? and sticky = ?", Uid_, Created_, Sticky_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByUidAndCreatedAndRevisionTranslationAffected Get NodeFieldRevisions via UidAndCreatedAndRevisionTranslationAffected
func GetNodeFieldRevisionsByUidAndCreatedAndRevisionTranslationAffected(offset int, limit int, Uid_ int, Created_ int, RevisionTranslationAffected_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("uid = ? and created = ? and revision_translation_affected = ?", Uid_, Created_, RevisionTranslationAffected_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByUidAndCreatedAndDefaultLangcode Get NodeFieldRevisions via UidAndCreatedAndDefaultLangcode
func GetNodeFieldRevisionsByUidAndCreatedAndDefaultLangcode(offset int, limit int, Uid_ int, Created_ int, DefaultLangcode_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("uid = ? and created = ? and default_langcode = ?", Uid_, Created_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByUidAndChangedAndPromote Get NodeFieldRevisions via UidAndChangedAndPromote
func GetNodeFieldRevisionsByUidAndChangedAndPromote(offset int, limit int, Uid_ int, Changed_ int, Promote_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("uid = ? and changed = ? and promote = ?", Uid_, Changed_, Promote_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByUidAndChangedAndSticky Get NodeFieldRevisions via UidAndChangedAndSticky
func GetNodeFieldRevisionsByUidAndChangedAndSticky(offset int, limit int, Uid_ int, Changed_ int, Sticky_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("uid = ? and changed = ? and sticky = ?", Uid_, Changed_, Sticky_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByUidAndChangedAndRevisionTranslationAffected Get NodeFieldRevisions via UidAndChangedAndRevisionTranslationAffected
func GetNodeFieldRevisionsByUidAndChangedAndRevisionTranslationAffected(offset int, limit int, Uid_ int, Changed_ int, RevisionTranslationAffected_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("uid = ? and changed = ? and revision_translation_affected = ?", Uid_, Changed_, RevisionTranslationAffected_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByUidAndChangedAndDefaultLangcode Get NodeFieldRevisions via UidAndChangedAndDefaultLangcode
func GetNodeFieldRevisionsByUidAndChangedAndDefaultLangcode(offset int, limit int, Uid_ int, Changed_ int, DefaultLangcode_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("uid = ? and changed = ? and default_langcode = ?", Uid_, Changed_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByUidAndPromoteAndSticky Get NodeFieldRevisions via UidAndPromoteAndSticky
func GetNodeFieldRevisionsByUidAndPromoteAndSticky(offset int, limit int, Uid_ int, Promote_ int, Sticky_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("uid = ? and promote = ? and sticky = ?", Uid_, Promote_, Sticky_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByUidAndPromoteAndRevisionTranslationAffected Get NodeFieldRevisions via UidAndPromoteAndRevisionTranslationAffected
func GetNodeFieldRevisionsByUidAndPromoteAndRevisionTranslationAffected(offset int, limit int, Uid_ int, Promote_ int, RevisionTranslationAffected_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("uid = ? and promote = ? and revision_translation_affected = ?", Uid_, Promote_, RevisionTranslationAffected_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByUidAndPromoteAndDefaultLangcode Get NodeFieldRevisions via UidAndPromoteAndDefaultLangcode
func GetNodeFieldRevisionsByUidAndPromoteAndDefaultLangcode(offset int, limit int, Uid_ int, Promote_ int, DefaultLangcode_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("uid = ? and promote = ? and default_langcode = ?", Uid_, Promote_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByUidAndStickyAndRevisionTranslationAffected Get NodeFieldRevisions via UidAndStickyAndRevisionTranslationAffected
func GetNodeFieldRevisionsByUidAndStickyAndRevisionTranslationAffected(offset int, limit int, Uid_ int, Sticky_ int, RevisionTranslationAffected_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("uid = ? and sticky = ? and revision_translation_affected = ?", Uid_, Sticky_, RevisionTranslationAffected_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByUidAndStickyAndDefaultLangcode Get NodeFieldRevisions via UidAndStickyAndDefaultLangcode
func GetNodeFieldRevisionsByUidAndStickyAndDefaultLangcode(offset int, limit int, Uid_ int, Sticky_ int, DefaultLangcode_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("uid = ? and sticky = ? and default_langcode = ?", Uid_, Sticky_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByUidAndRevisionTranslationAffectedAndDefaultLangcode Get NodeFieldRevisions via UidAndRevisionTranslationAffectedAndDefaultLangcode
func GetNodeFieldRevisionsByUidAndRevisionTranslationAffectedAndDefaultLangcode(offset int, limit int, Uid_ int, RevisionTranslationAffected_ int, DefaultLangcode_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("uid = ? and revision_translation_affected = ? and default_langcode = ?", Uid_, RevisionTranslationAffected_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByStatusAndCreatedAndChanged Get NodeFieldRevisions via StatusAndCreatedAndChanged
func GetNodeFieldRevisionsByStatusAndCreatedAndChanged(offset int, limit int, Status_ int, Created_ int, Changed_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("status = ? and created = ? and changed = ?", Status_, Created_, Changed_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByStatusAndCreatedAndPromote Get NodeFieldRevisions via StatusAndCreatedAndPromote
func GetNodeFieldRevisionsByStatusAndCreatedAndPromote(offset int, limit int, Status_ int, Created_ int, Promote_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("status = ? and created = ? and promote = ?", Status_, Created_, Promote_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByStatusAndCreatedAndSticky Get NodeFieldRevisions via StatusAndCreatedAndSticky
func GetNodeFieldRevisionsByStatusAndCreatedAndSticky(offset int, limit int, Status_ int, Created_ int, Sticky_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("status = ? and created = ? and sticky = ?", Status_, Created_, Sticky_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByStatusAndCreatedAndRevisionTranslationAffected Get NodeFieldRevisions via StatusAndCreatedAndRevisionTranslationAffected
func GetNodeFieldRevisionsByStatusAndCreatedAndRevisionTranslationAffected(offset int, limit int, Status_ int, Created_ int, RevisionTranslationAffected_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("status = ? and created = ? and revision_translation_affected = ?", Status_, Created_, RevisionTranslationAffected_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByStatusAndCreatedAndDefaultLangcode Get NodeFieldRevisions via StatusAndCreatedAndDefaultLangcode
func GetNodeFieldRevisionsByStatusAndCreatedAndDefaultLangcode(offset int, limit int, Status_ int, Created_ int, DefaultLangcode_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("status = ? and created = ? and default_langcode = ?", Status_, Created_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByStatusAndChangedAndPromote Get NodeFieldRevisions via StatusAndChangedAndPromote
func GetNodeFieldRevisionsByStatusAndChangedAndPromote(offset int, limit int, Status_ int, Changed_ int, Promote_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("status = ? and changed = ? and promote = ?", Status_, Changed_, Promote_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByStatusAndChangedAndSticky Get NodeFieldRevisions via StatusAndChangedAndSticky
func GetNodeFieldRevisionsByStatusAndChangedAndSticky(offset int, limit int, Status_ int, Changed_ int, Sticky_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("status = ? and changed = ? and sticky = ?", Status_, Changed_, Sticky_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByStatusAndChangedAndRevisionTranslationAffected Get NodeFieldRevisions via StatusAndChangedAndRevisionTranslationAffected
func GetNodeFieldRevisionsByStatusAndChangedAndRevisionTranslationAffected(offset int, limit int, Status_ int, Changed_ int, RevisionTranslationAffected_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("status = ? and changed = ? and revision_translation_affected = ?", Status_, Changed_, RevisionTranslationAffected_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByStatusAndChangedAndDefaultLangcode Get NodeFieldRevisions via StatusAndChangedAndDefaultLangcode
func GetNodeFieldRevisionsByStatusAndChangedAndDefaultLangcode(offset int, limit int, Status_ int, Changed_ int, DefaultLangcode_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("status = ? and changed = ? and default_langcode = ?", Status_, Changed_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByStatusAndPromoteAndSticky Get NodeFieldRevisions via StatusAndPromoteAndSticky
func GetNodeFieldRevisionsByStatusAndPromoteAndSticky(offset int, limit int, Status_ int, Promote_ int, Sticky_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("status = ? and promote = ? and sticky = ?", Status_, Promote_, Sticky_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByStatusAndPromoteAndRevisionTranslationAffected Get NodeFieldRevisions via StatusAndPromoteAndRevisionTranslationAffected
func GetNodeFieldRevisionsByStatusAndPromoteAndRevisionTranslationAffected(offset int, limit int, Status_ int, Promote_ int, RevisionTranslationAffected_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("status = ? and promote = ? and revision_translation_affected = ?", Status_, Promote_, RevisionTranslationAffected_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByStatusAndPromoteAndDefaultLangcode Get NodeFieldRevisions via StatusAndPromoteAndDefaultLangcode
func GetNodeFieldRevisionsByStatusAndPromoteAndDefaultLangcode(offset int, limit int, Status_ int, Promote_ int, DefaultLangcode_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("status = ? and promote = ? and default_langcode = ?", Status_, Promote_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByStatusAndStickyAndRevisionTranslationAffected Get NodeFieldRevisions via StatusAndStickyAndRevisionTranslationAffected
func GetNodeFieldRevisionsByStatusAndStickyAndRevisionTranslationAffected(offset int, limit int, Status_ int, Sticky_ int, RevisionTranslationAffected_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("status = ? and sticky = ? and revision_translation_affected = ?", Status_, Sticky_, RevisionTranslationAffected_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByStatusAndStickyAndDefaultLangcode Get NodeFieldRevisions via StatusAndStickyAndDefaultLangcode
func GetNodeFieldRevisionsByStatusAndStickyAndDefaultLangcode(offset int, limit int, Status_ int, Sticky_ int, DefaultLangcode_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("status = ? and sticky = ? and default_langcode = ?", Status_, Sticky_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByStatusAndRevisionTranslationAffectedAndDefaultLangcode Get NodeFieldRevisions via StatusAndRevisionTranslationAffectedAndDefaultLangcode
func GetNodeFieldRevisionsByStatusAndRevisionTranslationAffectedAndDefaultLangcode(offset int, limit int, Status_ int, RevisionTranslationAffected_ int, DefaultLangcode_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("status = ? and revision_translation_affected = ? and default_langcode = ?", Status_, RevisionTranslationAffected_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByCreatedAndChangedAndPromote Get NodeFieldRevisions via CreatedAndChangedAndPromote
func GetNodeFieldRevisionsByCreatedAndChangedAndPromote(offset int, limit int, Created_ int, Changed_ int, Promote_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("created = ? and changed = ? and promote = ?", Created_, Changed_, Promote_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByCreatedAndChangedAndSticky Get NodeFieldRevisions via CreatedAndChangedAndSticky
func GetNodeFieldRevisionsByCreatedAndChangedAndSticky(offset int, limit int, Created_ int, Changed_ int, Sticky_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("created = ? and changed = ? and sticky = ?", Created_, Changed_, Sticky_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByCreatedAndChangedAndRevisionTranslationAffected Get NodeFieldRevisions via CreatedAndChangedAndRevisionTranslationAffected
func GetNodeFieldRevisionsByCreatedAndChangedAndRevisionTranslationAffected(offset int, limit int, Created_ int, Changed_ int, RevisionTranslationAffected_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("created = ? and changed = ? and revision_translation_affected = ?", Created_, Changed_, RevisionTranslationAffected_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByCreatedAndChangedAndDefaultLangcode Get NodeFieldRevisions via CreatedAndChangedAndDefaultLangcode
func GetNodeFieldRevisionsByCreatedAndChangedAndDefaultLangcode(offset int, limit int, Created_ int, Changed_ int, DefaultLangcode_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("created = ? and changed = ? and default_langcode = ?", Created_, Changed_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByCreatedAndPromoteAndSticky Get NodeFieldRevisions via CreatedAndPromoteAndSticky
func GetNodeFieldRevisionsByCreatedAndPromoteAndSticky(offset int, limit int, Created_ int, Promote_ int, Sticky_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("created = ? and promote = ? and sticky = ?", Created_, Promote_, Sticky_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByCreatedAndPromoteAndRevisionTranslationAffected Get NodeFieldRevisions via CreatedAndPromoteAndRevisionTranslationAffected
func GetNodeFieldRevisionsByCreatedAndPromoteAndRevisionTranslationAffected(offset int, limit int, Created_ int, Promote_ int, RevisionTranslationAffected_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("created = ? and promote = ? and revision_translation_affected = ?", Created_, Promote_, RevisionTranslationAffected_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByCreatedAndPromoteAndDefaultLangcode Get NodeFieldRevisions via CreatedAndPromoteAndDefaultLangcode
func GetNodeFieldRevisionsByCreatedAndPromoteAndDefaultLangcode(offset int, limit int, Created_ int, Promote_ int, DefaultLangcode_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("created = ? and promote = ? and default_langcode = ?", Created_, Promote_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByCreatedAndStickyAndRevisionTranslationAffected Get NodeFieldRevisions via CreatedAndStickyAndRevisionTranslationAffected
func GetNodeFieldRevisionsByCreatedAndStickyAndRevisionTranslationAffected(offset int, limit int, Created_ int, Sticky_ int, RevisionTranslationAffected_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("created = ? and sticky = ? and revision_translation_affected = ?", Created_, Sticky_, RevisionTranslationAffected_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByCreatedAndStickyAndDefaultLangcode Get NodeFieldRevisions via CreatedAndStickyAndDefaultLangcode
func GetNodeFieldRevisionsByCreatedAndStickyAndDefaultLangcode(offset int, limit int, Created_ int, Sticky_ int, DefaultLangcode_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("created = ? and sticky = ? and default_langcode = ?", Created_, Sticky_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByCreatedAndRevisionTranslationAffectedAndDefaultLangcode Get NodeFieldRevisions via CreatedAndRevisionTranslationAffectedAndDefaultLangcode
func GetNodeFieldRevisionsByCreatedAndRevisionTranslationAffectedAndDefaultLangcode(offset int, limit int, Created_ int, RevisionTranslationAffected_ int, DefaultLangcode_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("created = ? and revision_translation_affected = ? and default_langcode = ?", Created_, RevisionTranslationAffected_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByChangedAndPromoteAndSticky Get NodeFieldRevisions via ChangedAndPromoteAndSticky
func GetNodeFieldRevisionsByChangedAndPromoteAndSticky(offset int, limit int, Changed_ int, Promote_ int, Sticky_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("changed = ? and promote = ? and sticky = ?", Changed_, Promote_, Sticky_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByChangedAndPromoteAndRevisionTranslationAffected Get NodeFieldRevisions via ChangedAndPromoteAndRevisionTranslationAffected
func GetNodeFieldRevisionsByChangedAndPromoteAndRevisionTranslationAffected(offset int, limit int, Changed_ int, Promote_ int, RevisionTranslationAffected_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("changed = ? and promote = ? and revision_translation_affected = ?", Changed_, Promote_, RevisionTranslationAffected_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByChangedAndPromoteAndDefaultLangcode Get NodeFieldRevisions via ChangedAndPromoteAndDefaultLangcode
func GetNodeFieldRevisionsByChangedAndPromoteAndDefaultLangcode(offset int, limit int, Changed_ int, Promote_ int, DefaultLangcode_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("changed = ? and promote = ? and default_langcode = ?", Changed_, Promote_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByChangedAndStickyAndRevisionTranslationAffected Get NodeFieldRevisions via ChangedAndStickyAndRevisionTranslationAffected
func GetNodeFieldRevisionsByChangedAndStickyAndRevisionTranslationAffected(offset int, limit int, Changed_ int, Sticky_ int, RevisionTranslationAffected_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("changed = ? and sticky = ? and revision_translation_affected = ?", Changed_, Sticky_, RevisionTranslationAffected_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByChangedAndStickyAndDefaultLangcode Get NodeFieldRevisions via ChangedAndStickyAndDefaultLangcode
func GetNodeFieldRevisionsByChangedAndStickyAndDefaultLangcode(offset int, limit int, Changed_ int, Sticky_ int, DefaultLangcode_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("changed = ? and sticky = ? and default_langcode = ?", Changed_, Sticky_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByChangedAndRevisionTranslationAffectedAndDefaultLangcode Get NodeFieldRevisions via ChangedAndRevisionTranslationAffectedAndDefaultLangcode
func GetNodeFieldRevisionsByChangedAndRevisionTranslationAffectedAndDefaultLangcode(offset int, limit int, Changed_ int, RevisionTranslationAffected_ int, DefaultLangcode_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("changed = ? and revision_translation_affected = ? and default_langcode = ?", Changed_, RevisionTranslationAffected_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByPromoteAndStickyAndRevisionTranslationAffected Get NodeFieldRevisions via PromoteAndStickyAndRevisionTranslationAffected
func GetNodeFieldRevisionsByPromoteAndStickyAndRevisionTranslationAffected(offset int, limit int, Promote_ int, Sticky_ int, RevisionTranslationAffected_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("promote = ? and sticky = ? and revision_translation_affected = ?", Promote_, Sticky_, RevisionTranslationAffected_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByPromoteAndStickyAndDefaultLangcode Get NodeFieldRevisions via PromoteAndStickyAndDefaultLangcode
func GetNodeFieldRevisionsByPromoteAndStickyAndDefaultLangcode(offset int, limit int, Promote_ int, Sticky_ int, DefaultLangcode_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("promote = ? and sticky = ? and default_langcode = ?", Promote_, Sticky_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByPromoteAndRevisionTranslationAffectedAndDefaultLangcode Get NodeFieldRevisions via PromoteAndRevisionTranslationAffectedAndDefaultLangcode
func GetNodeFieldRevisionsByPromoteAndRevisionTranslationAffectedAndDefaultLangcode(offset int, limit int, Promote_ int, RevisionTranslationAffected_ int, DefaultLangcode_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("promote = ? and revision_translation_affected = ? and default_langcode = ?", Promote_, RevisionTranslationAffected_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByStickyAndRevisionTranslationAffectedAndDefaultLangcode Get NodeFieldRevisions via StickyAndRevisionTranslationAffectedAndDefaultLangcode
func GetNodeFieldRevisionsByStickyAndRevisionTranslationAffectedAndDefaultLangcode(offset int, limit int, Sticky_ int, RevisionTranslationAffected_ int, DefaultLangcode_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("sticky = ? and revision_translation_affected = ? and default_langcode = ?", Sticky_, RevisionTranslationAffected_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByNidAndVid Get NodeFieldRevisions via NidAndVid
func GetNodeFieldRevisionsByNidAndVid(offset int, limit int, Nid_ int64, Vid_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("nid = ? and vid = ?", Nid_, Vid_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByNidAndLangcode Get NodeFieldRevisions via NidAndLangcode
func GetNodeFieldRevisionsByNidAndLangcode(offset int, limit int, Nid_ int64, Langcode_ string) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("nid = ? and langcode = ?", Nid_, Langcode_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByNidAndTitle Get NodeFieldRevisions via NidAndTitle
func GetNodeFieldRevisionsByNidAndTitle(offset int, limit int, Nid_ int64, Title_ string) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("nid = ? and title = ?", Nid_, Title_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByNidAndUid Get NodeFieldRevisions via NidAndUid
func GetNodeFieldRevisionsByNidAndUid(offset int, limit int, Nid_ int64, Uid_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("nid = ? and uid = ?", Nid_, Uid_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByNidAndStatus Get NodeFieldRevisions via NidAndStatus
func GetNodeFieldRevisionsByNidAndStatus(offset int, limit int, Nid_ int64, Status_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("nid = ? and status = ?", Nid_, Status_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByNidAndCreated Get NodeFieldRevisions via NidAndCreated
func GetNodeFieldRevisionsByNidAndCreated(offset int, limit int, Nid_ int64, Created_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("nid = ? and created = ?", Nid_, Created_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByNidAndChanged Get NodeFieldRevisions via NidAndChanged
func GetNodeFieldRevisionsByNidAndChanged(offset int, limit int, Nid_ int64, Changed_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("nid = ? and changed = ?", Nid_, Changed_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByNidAndPromote Get NodeFieldRevisions via NidAndPromote
func GetNodeFieldRevisionsByNidAndPromote(offset int, limit int, Nid_ int64, Promote_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("nid = ? and promote = ?", Nid_, Promote_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByNidAndSticky Get NodeFieldRevisions via NidAndSticky
func GetNodeFieldRevisionsByNidAndSticky(offset int, limit int, Nid_ int64, Sticky_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("nid = ? and sticky = ?", Nid_, Sticky_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByNidAndRevisionTranslationAffected Get NodeFieldRevisions via NidAndRevisionTranslationAffected
func GetNodeFieldRevisionsByNidAndRevisionTranslationAffected(offset int, limit int, Nid_ int64, RevisionTranslationAffected_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("nid = ? and revision_translation_affected = ?", Nid_, RevisionTranslationAffected_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByNidAndDefaultLangcode Get NodeFieldRevisions via NidAndDefaultLangcode
func GetNodeFieldRevisionsByNidAndDefaultLangcode(offset int, limit int, Nid_ int64, DefaultLangcode_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("nid = ? and default_langcode = ?", Nid_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByVidAndLangcode Get NodeFieldRevisions via VidAndLangcode
func GetNodeFieldRevisionsByVidAndLangcode(offset int, limit int, Vid_ int, Langcode_ string) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("vid = ? and langcode = ?", Vid_, Langcode_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByVidAndTitle Get NodeFieldRevisions via VidAndTitle
func GetNodeFieldRevisionsByVidAndTitle(offset int, limit int, Vid_ int, Title_ string) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("vid = ? and title = ?", Vid_, Title_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByVidAndUid Get NodeFieldRevisions via VidAndUid
func GetNodeFieldRevisionsByVidAndUid(offset int, limit int, Vid_ int, Uid_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("vid = ? and uid = ?", Vid_, Uid_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByVidAndStatus Get NodeFieldRevisions via VidAndStatus
func GetNodeFieldRevisionsByVidAndStatus(offset int, limit int, Vid_ int, Status_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("vid = ? and status = ?", Vid_, Status_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByVidAndCreated Get NodeFieldRevisions via VidAndCreated
func GetNodeFieldRevisionsByVidAndCreated(offset int, limit int, Vid_ int, Created_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("vid = ? and created = ?", Vid_, Created_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByVidAndChanged Get NodeFieldRevisions via VidAndChanged
func GetNodeFieldRevisionsByVidAndChanged(offset int, limit int, Vid_ int, Changed_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("vid = ? and changed = ?", Vid_, Changed_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByVidAndPromote Get NodeFieldRevisions via VidAndPromote
func GetNodeFieldRevisionsByVidAndPromote(offset int, limit int, Vid_ int, Promote_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("vid = ? and promote = ?", Vid_, Promote_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByVidAndSticky Get NodeFieldRevisions via VidAndSticky
func GetNodeFieldRevisionsByVidAndSticky(offset int, limit int, Vid_ int, Sticky_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("vid = ? and sticky = ?", Vid_, Sticky_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByVidAndRevisionTranslationAffected Get NodeFieldRevisions via VidAndRevisionTranslationAffected
func GetNodeFieldRevisionsByVidAndRevisionTranslationAffected(offset int, limit int, Vid_ int, RevisionTranslationAffected_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("vid = ? and revision_translation_affected = ?", Vid_, RevisionTranslationAffected_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByVidAndDefaultLangcode Get NodeFieldRevisions via VidAndDefaultLangcode
func GetNodeFieldRevisionsByVidAndDefaultLangcode(offset int, limit int, Vid_ int, DefaultLangcode_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("vid = ? and default_langcode = ?", Vid_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByLangcodeAndTitle Get NodeFieldRevisions via LangcodeAndTitle
func GetNodeFieldRevisionsByLangcodeAndTitle(offset int, limit int, Langcode_ string, Title_ string) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("langcode = ? and title = ?", Langcode_, Title_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByLangcodeAndUid Get NodeFieldRevisions via LangcodeAndUid
func GetNodeFieldRevisionsByLangcodeAndUid(offset int, limit int, Langcode_ string, Uid_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("langcode = ? and uid = ?", Langcode_, Uid_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByLangcodeAndStatus Get NodeFieldRevisions via LangcodeAndStatus
func GetNodeFieldRevisionsByLangcodeAndStatus(offset int, limit int, Langcode_ string, Status_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("langcode = ? and status = ?", Langcode_, Status_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByLangcodeAndCreated Get NodeFieldRevisions via LangcodeAndCreated
func GetNodeFieldRevisionsByLangcodeAndCreated(offset int, limit int, Langcode_ string, Created_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("langcode = ? and created = ?", Langcode_, Created_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByLangcodeAndChanged Get NodeFieldRevisions via LangcodeAndChanged
func GetNodeFieldRevisionsByLangcodeAndChanged(offset int, limit int, Langcode_ string, Changed_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("langcode = ? and changed = ?", Langcode_, Changed_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByLangcodeAndPromote Get NodeFieldRevisions via LangcodeAndPromote
func GetNodeFieldRevisionsByLangcodeAndPromote(offset int, limit int, Langcode_ string, Promote_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("langcode = ? and promote = ?", Langcode_, Promote_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByLangcodeAndSticky Get NodeFieldRevisions via LangcodeAndSticky
func GetNodeFieldRevisionsByLangcodeAndSticky(offset int, limit int, Langcode_ string, Sticky_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("langcode = ? and sticky = ?", Langcode_, Sticky_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByLangcodeAndRevisionTranslationAffected Get NodeFieldRevisions via LangcodeAndRevisionTranslationAffected
func GetNodeFieldRevisionsByLangcodeAndRevisionTranslationAffected(offset int, limit int, Langcode_ string, RevisionTranslationAffected_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("langcode = ? and revision_translation_affected = ?", Langcode_, RevisionTranslationAffected_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByLangcodeAndDefaultLangcode Get NodeFieldRevisions via LangcodeAndDefaultLangcode
func GetNodeFieldRevisionsByLangcodeAndDefaultLangcode(offset int, limit int, Langcode_ string, DefaultLangcode_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("langcode = ? and default_langcode = ?", Langcode_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByTitleAndUid Get NodeFieldRevisions via TitleAndUid
func GetNodeFieldRevisionsByTitleAndUid(offset int, limit int, Title_ string, Uid_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("title = ? and uid = ?", Title_, Uid_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByTitleAndStatus Get NodeFieldRevisions via TitleAndStatus
func GetNodeFieldRevisionsByTitleAndStatus(offset int, limit int, Title_ string, Status_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("title = ? and status = ?", Title_, Status_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByTitleAndCreated Get NodeFieldRevisions via TitleAndCreated
func GetNodeFieldRevisionsByTitleAndCreated(offset int, limit int, Title_ string, Created_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("title = ? and created = ?", Title_, Created_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByTitleAndChanged Get NodeFieldRevisions via TitleAndChanged
func GetNodeFieldRevisionsByTitleAndChanged(offset int, limit int, Title_ string, Changed_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("title = ? and changed = ?", Title_, Changed_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByTitleAndPromote Get NodeFieldRevisions via TitleAndPromote
func GetNodeFieldRevisionsByTitleAndPromote(offset int, limit int, Title_ string, Promote_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("title = ? and promote = ?", Title_, Promote_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByTitleAndSticky Get NodeFieldRevisions via TitleAndSticky
func GetNodeFieldRevisionsByTitleAndSticky(offset int, limit int, Title_ string, Sticky_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("title = ? and sticky = ?", Title_, Sticky_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByTitleAndRevisionTranslationAffected Get NodeFieldRevisions via TitleAndRevisionTranslationAffected
func GetNodeFieldRevisionsByTitleAndRevisionTranslationAffected(offset int, limit int, Title_ string, RevisionTranslationAffected_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("title = ? and revision_translation_affected = ?", Title_, RevisionTranslationAffected_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByTitleAndDefaultLangcode Get NodeFieldRevisions via TitleAndDefaultLangcode
func GetNodeFieldRevisionsByTitleAndDefaultLangcode(offset int, limit int, Title_ string, DefaultLangcode_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("title = ? and default_langcode = ?", Title_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByUidAndStatus Get NodeFieldRevisions via UidAndStatus
func GetNodeFieldRevisionsByUidAndStatus(offset int, limit int, Uid_ int, Status_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("uid = ? and status = ?", Uid_, Status_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByUidAndCreated Get NodeFieldRevisions via UidAndCreated
func GetNodeFieldRevisionsByUidAndCreated(offset int, limit int, Uid_ int, Created_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("uid = ? and created = ?", Uid_, Created_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByUidAndChanged Get NodeFieldRevisions via UidAndChanged
func GetNodeFieldRevisionsByUidAndChanged(offset int, limit int, Uid_ int, Changed_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("uid = ? and changed = ?", Uid_, Changed_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByUidAndPromote Get NodeFieldRevisions via UidAndPromote
func GetNodeFieldRevisionsByUidAndPromote(offset int, limit int, Uid_ int, Promote_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("uid = ? and promote = ?", Uid_, Promote_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByUidAndSticky Get NodeFieldRevisions via UidAndSticky
func GetNodeFieldRevisionsByUidAndSticky(offset int, limit int, Uid_ int, Sticky_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("uid = ? and sticky = ?", Uid_, Sticky_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByUidAndRevisionTranslationAffected Get NodeFieldRevisions via UidAndRevisionTranslationAffected
func GetNodeFieldRevisionsByUidAndRevisionTranslationAffected(offset int, limit int, Uid_ int, RevisionTranslationAffected_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("uid = ? and revision_translation_affected = ?", Uid_, RevisionTranslationAffected_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByUidAndDefaultLangcode Get NodeFieldRevisions via UidAndDefaultLangcode
func GetNodeFieldRevisionsByUidAndDefaultLangcode(offset int, limit int, Uid_ int, DefaultLangcode_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("uid = ? and default_langcode = ?", Uid_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByStatusAndCreated Get NodeFieldRevisions via StatusAndCreated
func GetNodeFieldRevisionsByStatusAndCreated(offset int, limit int, Status_ int, Created_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("status = ? and created = ?", Status_, Created_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByStatusAndChanged Get NodeFieldRevisions via StatusAndChanged
func GetNodeFieldRevisionsByStatusAndChanged(offset int, limit int, Status_ int, Changed_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("status = ? and changed = ?", Status_, Changed_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByStatusAndPromote Get NodeFieldRevisions via StatusAndPromote
func GetNodeFieldRevisionsByStatusAndPromote(offset int, limit int, Status_ int, Promote_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("status = ? and promote = ?", Status_, Promote_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByStatusAndSticky Get NodeFieldRevisions via StatusAndSticky
func GetNodeFieldRevisionsByStatusAndSticky(offset int, limit int, Status_ int, Sticky_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("status = ? and sticky = ?", Status_, Sticky_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByStatusAndRevisionTranslationAffected Get NodeFieldRevisions via StatusAndRevisionTranslationAffected
func GetNodeFieldRevisionsByStatusAndRevisionTranslationAffected(offset int, limit int, Status_ int, RevisionTranslationAffected_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("status = ? and revision_translation_affected = ?", Status_, RevisionTranslationAffected_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByStatusAndDefaultLangcode Get NodeFieldRevisions via StatusAndDefaultLangcode
func GetNodeFieldRevisionsByStatusAndDefaultLangcode(offset int, limit int, Status_ int, DefaultLangcode_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("status = ? and default_langcode = ?", Status_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByCreatedAndChanged Get NodeFieldRevisions via CreatedAndChanged
func GetNodeFieldRevisionsByCreatedAndChanged(offset int, limit int, Created_ int, Changed_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("created = ? and changed = ?", Created_, Changed_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByCreatedAndPromote Get NodeFieldRevisions via CreatedAndPromote
func GetNodeFieldRevisionsByCreatedAndPromote(offset int, limit int, Created_ int, Promote_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("created = ? and promote = ?", Created_, Promote_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByCreatedAndSticky Get NodeFieldRevisions via CreatedAndSticky
func GetNodeFieldRevisionsByCreatedAndSticky(offset int, limit int, Created_ int, Sticky_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("created = ? and sticky = ?", Created_, Sticky_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByCreatedAndRevisionTranslationAffected Get NodeFieldRevisions via CreatedAndRevisionTranslationAffected
func GetNodeFieldRevisionsByCreatedAndRevisionTranslationAffected(offset int, limit int, Created_ int, RevisionTranslationAffected_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("created = ? and revision_translation_affected = ?", Created_, RevisionTranslationAffected_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByCreatedAndDefaultLangcode Get NodeFieldRevisions via CreatedAndDefaultLangcode
func GetNodeFieldRevisionsByCreatedAndDefaultLangcode(offset int, limit int, Created_ int, DefaultLangcode_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("created = ? and default_langcode = ?", Created_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByChangedAndPromote Get NodeFieldRevisions via ChangedAndPromote
func GetNodeFieldRevisionsByChangedAndPromote(offset int, limit int, Changed_ int, Promote_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("changed = ? and promote = ?", Changed_, Promote_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByChangedAndSticky Get NodeFieldRevisions via ChangedAndSticky
func GetNodeFieldRevisionsByChangedAndSticky(offset int, limit int, Changed_ int, Sticky_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("changed = ? and sticky = ?", Changed_, Sticky_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByChangedAndRevisionTranslationAffected Get NodeFieldRevisions via ChangedAndRevisionTranslationAffected
func GetNodeFieldRevisionsByChangedAndRevisionTranslationAffected(offset int, limit int, Changed_ int, RevisionTranslationAffected_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("changed = ? and revision_translation_affected = ?", Changed_, RevisionTranslationAffected_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByChangedAndDefaultLangcode Get NodeFieldRevisions via ChangedAndDefaultLangcode
func GetNodeFieldRevisionsByChangedAndDefaultLangcode(offset int, limit int, Changed_ int, DefaultLangcode_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("changed = ? and default_langcode = ?", Changed_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByPromoteAndSticky Get NodeFieldRevisions via PromoteAndSticky
func GetNodeFieldRevisionsByPromoteAndSticky(offset int, limit int, Promote_ int, Sticky_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("promote = ? and sticky = ?", Promote_, Sticky_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByPromoteAndRevisionTranslationAffected Get NodeFieldRevisions via PromoteAndRevisionTranslationAffected
func GetNodeFieldRevisionsByPromoteAndRevisionTranslationAffected(offset int, limit int, Promote_ int, RevisionTranslationAffected_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("promote = ? and revision_translation_affected = ?", Promote_, RevisionTranslationAffected_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByPromoteAndDefaultLangcode Get NodeFieldRevisions via PromoteAndDefaultLangcode
func GetNodeFieldRevisionsByPromoteAndDefaultLangcode(offset int, limit int, Promote_ int, DefaultLangcode_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("promote = ? and default_langcode = ?", Promote_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByStickyAndRevisionTranslationAffected Get NodeFieldRevisions via StickyAndRevisionTranslationAffected
func GetNodeFieldRevisionsByStickyAndRevisionTranslationAffected(offset int, limit int, Sticky_ int, RevisionTranslationAffected_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("sticky = ? and revision_translation_affected = ?", Sticky_, RevisionTranslationAffected_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByStickyAndDefaultLangcode Get NodeFieldRevisions via StickyAndDefaultLangcode
func GetNodeFieldRevisionsByStickyAndDefaultLangcode(offset int, limit int, Sticky_ int, DefaultLangcode_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("sticky = ? and default_langcode = ?", Sticky_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsByRevisionTranslationAffectedAndDefaultLangcode Get NodeFieldRevisions via RevisionTranslationAffectedAndDefaultLangcode
func GetNodeFieldRevisionsByRevisionTranslationAffectedAndDefaultLangcode(offset int, limit int, RevisionTranslationAffected_ int, DefaultLangcode_ int) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("revision_translation_affected = ? and default_langcode = ?", RevisionTranslationAffected_, DefaultLangcode_).Limit(limit, offset).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisions Get NodeFieldRevisions via field
func GetNodeFieldRevisions(offset int, limit int, field string) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Limit(limit, offset).Desc(field).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsViaNid Get NodeFieldRevisions via Nid
func GetNodeFieldRevisionsViaNid(offset int, limit int, Nid_ int64, field string) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("nid = ?", Nid_).Limit(limit, offset).Desc(field).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsViaVid Get NodeFieldRevisions via Vid
func GetNodeFieldRevisionsViaVid(offset int, limit int, Vid_ int, field string) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("vid = ?", Vid_).Limit(limit, offset).Desc(field).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsViaLangcode Get NodeFieldRevisions via Langcode
func GetNodeFieldRevisionsViaLangcode(offset int, limit int, Langcode_ string, field string) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("langcode = ?", Langcode_).Limit(limit, offset).Desc(field).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsViaTitle Get NodeFieldRevisions via Title
func GetNodeFieldRevisionsViaTitle(offset int, limit int, Title_ string, field string) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("title = ?", Title_).Limit(limit, offset).Desc(field).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsViaUid Get NodeFieldRevisions via Uid
func GetNodeFieldRevisionsViaUid(offset int, limit int, Uid_ int, field string) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("uid = ?", Uid_).Limit(limit, offset).Desc(field).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsViaStatus Get NodeFieldRevisions via Status
func GetNodeFieldRevisionsViaStatus(offset int, limit int, Status_ int, field string) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("status = ?", Status_).Limit(limit, offset).Desc(field).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsViaCreated Get NodeFieldRevisions via Created
func GetNodeFieldRevisionsViaCreated(offset int, limit int, Created_ int, field string) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("created = ?", Created_).Limit(limit, offset).Desc(field).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsViaChanged Get NodeFieldRevisions via Changed
func GetNodeFieldRevisionsViaChanged(offset int, limit int, Changed_ int, field string) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("changed = ?", Changed_).Limit(limit, offset).Desc(field).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsViaPromote Get NodeFieldRevisions via Promote
func GetNodeFieldRevisionsViaPromote(offset int, limit int, Promote_ int, field string) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("promote = ?", Promote_).Limit(limit, offset).Desc(field).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsViaSticky Get NodeFieldRevisions via Sticky
func GetNodeFieldRevisionsViaSticky(offset int, limit int, Sticky_ int, field string) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("sticky = ?", Sticky_).Limit(limit, offset).Desc(field).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsViaRevisionTranslationAffected Get NodeFieldRevisions via RevisionTranslationAffected
func GetNodeFieldRevisionsViaRevisionTranslationAffected(offset int, limit int, RevisionTranslationAffected_ int, field string) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("revision_translation_affected = ?", RevisionTranslationAffected_).Limit(limit, offset).Desc(field).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// GetNodeFieldRevisionsViaDefaultLangcode Get NodeFieldRevisions via DefaultLangcode
func GetNodeFieldRevisionsViaDefaultLangcode(offset int, limit int, DefaultLangcode_ int, field string) (*[]*NodeFieldRevision, error) {
	var _NodeFieldRevision = new([]*NodeFieldRevision)
	err := Engine.Table("node_field_revision").Where("default_langcode = ?", DefaultLangcode_).Limit(limit, offset).Desc(field).Find(_NodeFieldRevision)
	return _NodeFieldRevision, err
}

// HasNodeFieldRevisionViaNid Has NodeFieldRevision via Nid
func HasNodeFieldRevisionViaNid(iNid int64) bool {
	if has, err := Engine.Where("nid = ?", iNid).Get(new(NodeFieldRevision)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasNodeFieldRevisionViaVid Has NodeFieldRevision via Vid
func HasNodeFieldRevisionViaVid(iVid int) bool {
	if has, err := Engine.Where("vid = ?", iVid).Get(new(NodeFieldRevision)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasNodeFieldRevisionViaLangcode Has NodeFieldRevision via Langcode
func HasNodeFieldRevisionViaLangcode(iLangcode string) bool {
	if has, err := Engine.Where("langcode = ?", iLangcode).Get(new(NodeFieldRevision)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasNodeFieldRevisionViaTitle Has NodeFieldRevision via Title
func HasNodeFieldRevisionViaTitle(iTitle string) bool {
	if has, err := Engine.Where("title = ?", iTitle).Get(new(NodeFieldRevision)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasNodeFieldRevisionViaUid Has NodeFieldRevision via Uid
func HasNodeFieldRevisionViaUid(iUid int) bool {
	if has, err := Engine.Where("uid = ?", iUid).Get(new(NodeFieldRevision)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasNodeFieldRevisionViaStatus Has NodeFieldRevision via Status
func HasNodeFieldRevisionViaStatus(iStatus int) bool {
	if has, err := Engine.Where("status = ?", iStatus).Get(new(NodeFieldRevision)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasNodeFieldRevisionViaCreated Has NodeFieldRevision via Created
func HasNodeFieldRevisionViaCreated(iCreated int) bool {
	if has, err := Engine.Where("created = ?", iCreated).Get(new(NodeFieldRevision)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasNodeFieldRevisionViaChanged Has NodeFieldRevision via Changed
func HasNodeFieldRevisionViaChanged(iChanged int) bool {
	if has, err := Engine.Where("changed = ?", iChanged).Get(new(NodeFieldRevision)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasNodeFieldRevisionViaPromote Has NodeFieldRevision via Promote
func HasNodeFieldRevisionViaPromote(iPromote int) bool {
	if has, err := Engine.Where("promote = ?", iPromote).Get(new(NodeFieldRevision)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasNodeFieldRevisionViaSticky Has NodeFieldRevision via Sticky
func HasNodeFieldRevisionViaSticky(iSticky int) bool {
	if has, err := Engine.Where("sticky = ?", iSticky).Get(new(NodeFieldRevision)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasNodeFieldRevisionViaRevisionTranslationAffected Has NodeFieldRevision via RevisionTranslationAffected
func HasNodeFieldRevisionViaRevisionTranslationAffected(iRevisionTranslationAffected int) bool {
	if has, err := Engine.Where("revision_translation_affected = ?", iRevisionTranslationAffected).Get(new(NodeFieldRevision)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasNodeFieldRevisionViaDefaultLangcode Has NodeFieldRevision via DefaultLangcode
func HasNodeFieldRevisionViaDefaultLangcode(iDefaultLangcode int) bool {
	if has, err := Engine.Where("default_langcode = ?", iDefaultLangcode).Get(new(NodeFieldRevision)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// GetNodeFieldRevisionViaNid Get NodeFieldRevision via Nid
func GetNodeFieldRevisionViaNid(iNid int64) (*NodeFieldRevision, error) {
	var _NodeFieldRevision = &NodeFieldRevision{Nid: iNid}
	has, err := Engine.Get(_NodeFieldRevision)
	if has {
		return _NodeFieldRevision, err
	} else {
		return nil, err
	}
}

// GetNodeFieldRevisionViaVid Get NodeFieldRevision via Vid
func GetNodeFieldRevisionViaVid(iVid int) (*NodeFieldRevision, error) {
	var _NodeFieldRevision = &NodeFieldRevision{Vid: iVid}
	has, err := Engine.Get(_NodeFieldRevision)
	if has {
		return _NodeFieldRevision, err
	} else {
		return nil, err
	}
}

// GetNodeFieldRevisionViaLangcode Get NodeFieldRevision via Langcode
func GetNodeFieldRevisionViaLangcode(iLangcode string) (*NodeFieldRevision, error) {
	var _NodeFieldRevision = &NodeFieldRevision{Langcode: iLangcode}
	has, err := Engine.Get(_NodeFieldRevision)
	if has {
		return _NodeFieldRevision, err
	} else {
		return nil, err
	}
}

// GetNodeFieldRevisionViaTitle Get NodeFieldRevision via Title
func GetNodeFieldRevisionViaTitle(iTitle string) (*NodeFieldRevision, error) {
	var _NodeFieldRevision = &NodeFieldRevision{Title: iTitle}
	has, err := Engine.Get(_NodeFieldRevision)
	if has {
		return _NodeFieldRevision, err
	} else {
		return nil, err
	}
}

// GetNodeFieldRevisionViaUid Get NodeFieldRevision via Uid
func GetNodeFieldRevisionViaUid(iUid int) (*NodeFieldRevision, error) {
	var _NodeFieldRevision = &NodeFieldRevision{Uid: iUid}
	has, err := Engine.Get(_NodeFieldRevision)
	if has {
		return _NodeFieldRevision, err
	} else {
		return nil, err
	}
}

// GetNodeFieldRevisionViaStatus Get NodeFieldRevision via Status
func GetNodeFieldRevisionViaStatus(iStatus int) (*NodeFieldRevision, error) {
	var _NodeFieldRevision = &NodeFieldRevision{Status: iStatus}
	has, err := Engine.Get(_NodeFieldRevision)
	if has {
		return _NodeFieldRevision, err
	} else {
		return nil, err
	}
}

// GetNodeFieldRevisionViaCreated Get NodeFieldRevision via Created
func GetNodeFieldRevisionViaCreated(iCreated int) (*NodeFieldRevision, error) {
	var _NodeFieldRevision = &NodeFieldRevision{Created: iCreated}
	has, err := Engine.Get(_NodeFieldRevision)
	if has {
		return _NodeFieldRevision, err
	} else {
		return nil, err
	}
}

// GetNodeFieldRevisionViaChanged Get NodeFieldRevision via Changed
func GetNodeFieldRevisionViaChanged(iChanged int) (*NodeFieldRevision, error) {
	var _NodeFieldRevision = &NodeFieldRevision{Changed: iChanged}
	has, err := Engine.Get(_NodeFieldRevision)
	if has {
		return _NodeFieldRevision, err
	} else {
		return nil, err
	}
}

// GetNodeFieldRevisionViaPromote Get NodeFieldRevision via Promote
func GetNodeFieldRevisionViaPromote(iPromote int) (*NodeFieldRevision, error) {
	var _NodeFieldRevision = &NodeFieldRevision{Promote: iPromote}
	has, err := Engine.Get(_NodeFieldRevision)
	if has {
		return _NodeFieldRevision, err
	} else {
		return nil, err
	}
}

// GetNodeFieldRevisionViaSticky Get NodeFieldRevision via Sticky
func GetNodeFieldRevisionViaSticky(iSticky int) (*NodeFieldRevision, error) {
	var _NodeFieldRevision = &NodeFieldRevision{Sticky: iSticky}
	has, err := Engine.Get(_NodeFieldRevision)
	if has {
		return _NodeFieldRevision, err
	} else {
		return nil, err
	}
}

// GetNodeFieldRevisionViaRevisionTranslationAffected Get NodeFieldRevision via RevisionTranslationAffected
func GetNodeFieldRevisionViaRevisionTranslationAffected(iRevisionTranslationAffected int) (*NodeFieldRevision, error) {
	var _NodeFieldRevision = &NodeFieldRevision{RevisionTranslationAffected: iRevisionTranslationAffected}
	has, err := Engine.Get(_NodeFieldRevision)
	if has {
		return _NodeFieldRevision, err
	} else {
		return nil, err
	}
}

// GetNodeFieldRevisionViaDefaultLangcode Get NodeFieldRevision via DefaultLangcode
func GetNodeFieldRevisionViaDefaultLangcode(iDefaultLangcode int) (*NodeFieldRevision, error) {
	var _NodeFieldRevision = &NodeFieldRevision{DefaultLangcode: iDefaultLangcode}
	has, err := Engine.Get(_NodeFieldRevision)
	if has {
		return _NodeFieldRevision, err
	} else {
		return nil, err
	}
}

// SetNodeFieldRevisionViaNid Set NodeFieldRevision via Nid
func SetNodeFieldRevisionViaNid(iNid int64, node_field_revision *NodeFieldRevision) (int64, error) {
	node_field_revision.Nid = iNid
	return Engine.Insert(node_field_revision)
}

// SetNodeFieldRevisionViaVid Set NodeFieldRevision via Vid
func SetNodeFieldRevisionViaVid(iVid int, node_field_revision *NodeFieldRevision) (int64, error) {
	node_field_revision.Vid = iVid
	return Engine.Insert(node_field_revision)
}

// SetNodeFieldRevisionViaLangcode Set NodeFieldRevision via Langcode
func SetNodeFieldRevisionViaLangcode(iLangcode string, node_field_revision *NodeFieldRevision) (int64, error) {
	node_field_revision.Langcode = iLangcode
	return Engine.Insert(node_field_revision)
}

// SetNodeFieldRevisionViaTitle Set NodeFieldRevision via Title
func SetNodeFieldRevisionViaTitle(iTitle string, node_field_revision *NodeFieldRevision) (int64, error) {
	node_field_revision.Title = iTitle
	return Engine.Insert(node_field_revision)
}

// SetNodeFieldRevisionViaUid Set NodeFieldRevision via Uid
func SetNodeFieldRevisionViaUid(iUid int, node_field_revision *NodeFieldRevision) (int64, error) {
	node_field_revision.Uid = iUid
	return Engine.Insert(node_field_revision)
}

// SetNodeFieldRevisionViaStatus Set NodeFieldRevision via Status
func SetNodeFieldRevisionViaStatus(iStatus int, node_field_revision *NodeFieldRevision) (int64, error) {
	node_field_revision.Status = iStatus
	return Engine.Insert(node_field_revision)
}

// SetNodeFieldRevisionViaCreated Set NodeFieldRevision via Created
func SetNodeFieldRevisionViaCreated(iCreated int, node_field_revision *NodeFieldRevision) (int64, error) {
	node_field_revision.Created = iCreated
	return Engine.Insert(node_field_revision)
}

// SetNodeFieldRevisionViaChanged Set NodeFieldRevision via Changed
func SetNodeFieldRevisionViaChanged(iChanged int, node_field_revision *NodeFieldRevision) (int64, error) {
	node_field_revision.Changed = iChanged
	return Engine.Insert(node_field_revision)
}

// SetNodeFieldRevisionViaPromote Set NodeFieldRevision via Promote
func SetNodeFieldRevisionViaPromote(iPromote int, node_field_revision *NodeFieldRevision) (int64, error) {
	node_field_revision.Promote = iPromote
	return Engine.Insert(node_field_revision)
}

// SetNodeFieldRevisionViaSticky Set NodeFieldRevision via Sticky
func SetNodeFieldRevisionViaSticky(iSticky int, node_field_revision *NodeFieldRevision) (int64, error) {
	node_field_revision.Sticky = iSticky
	return Engine.Insert(node_field_revision)
}

// SetNodeFieldRevisionViaRevisionTranslationAffected Set NodeFieldRevision via RevisionTranslationAffected
func SetNodeFieldRevisionViaRevisionTranslationAffected(iRevisionTranslationAffected int, node_field_revision *NodeFieldRevision) (int64, error) {
	node_field_revision.RevisionTranslationAffected = iRevisionTranslationAffected
	return Engine.Insert(node_field_revision)
}

// SetNodeFieldRevisionViaDefaultLangcode Set NodeFieldRevision via DefaultLangcode
func SetNodeFieldRevisionViaDefaultLangcode(iDefaultLangcode int, node_field_revision *NodeFieldRevision) (int64, error) {
	node_field_revision.DefaultLangcode = iDefaultLangcode
	return Engine.Insert(node_field_revision)
}

// AddNodeFieldRevision Add NodeFieldRevision via all columns
func AddNodeFieldRevision(iNid int64, iVid int, iLangcode string, iTitle string, iUid int, iStatus int, iCreated int, iChanged int, iPromote int, iSticky int, iRevisionTranslationAffected int, iDefaultLangcode int) error {
	_NodeFieldRevision := &NodeFieldRevision{Nid: iNid, Vid: iVid, Langcode: iLangcode, Title: iTitle, Uid: iUid, Status: iStatus, Created: iCreated, Changed: iChanged, Promote: iPromote, Sticky: iSticky, RevisionTranslationAffected: iRevisionTranslationAffected, DefaultLangcode: iDefaultLangcode}
	if _, err := Engine.Insert(_NodeFieldRevision); err != nil {
		return err
	} else {
		return nil
	}
}

// PostNodeFieldRevision Post NodeFieldRevision via iNodeFieldRevision
func PostNodeFieldRevision(iNodeFieldRevision *NodeFieldRevision) (int64, error) {
	_, err := Engine.Insert(iNodeFieldRevision)
	return iNodeFieldRevision.Nid, err
}

// PutNodeFieldRevision Put NodeFieldRevision
func PutNodeFieldRevision(iNodeFieldRevision *NodeFieldRevision) (int64, error) {
	_, err := Engine.Id(iNodeFieldRevision.Nid).Update(iNodeFieldRevision)
	return iNodeFieldRevision.Nid, err
}

// PutNodeFieldRevisionViaNid Put NodeFieldRevision via Nid
func PutNodeFieldRevisionViaNid(Nid_ int64, iNodeFieldRevision *NodeFieldRevision) (int64, error) {
	row, err := Engine.Update(iNodeFieldRevision, &NodeFieldRevision{Nid: Nid_})
	return row, err
}

// PutNodeFieldRevisionViaVid Put NodeFieldRevision via Vid
func PutNodeFieldRevisionViaVid(Vid_ int, iNodeFieldRevision *NodeFieldRevision) (int64, error) {
	row, err := Engine.Update(iNodeFieldRevision, &NodeFieldRevision{Vid: Vid_})
	return row, err
}

// PutNodeFieldRevisionViaLangcode Put NodeFieldRevision via Langcode
func PutNodeFieldRevisionViaLangcode(Langcode_ string, iNodeFieldRevision *NodeFieldRevision) (int64, error) {
	row, err := Engine.Update(iNodeFieldRevision, &NodeFieldRevision{Langcode: Langcode_})
	return row, err
}

// PutNodeFieldRevisionViaTitle Put NodeFieldRevision via Title
func PutNodeFieldRevisionViaTitle(Title_ string, iNodeFieldRevision *NodeFieldRevision) (int64, error) {
	row, err := Engine.Update(iNodeFieldRevision, &NodeFieldRevision{Title: Title_})
	return row, err
}

// PutNodeFieldRevisionViaUid Put NodeFieldRevision via Uid
func PutNodeFieldRevisionViaUid(Uid_ int, iNodeFieldRevision *NodeFieldRevision) (int64, error) {
	row, err := Engine.Update(iNodeFieldRevision, &NodeFieldRevision{Uid: Uid_})
	return row, err
}

// PutNodeFieldRevisionViaStatus Put NodeFieldRevision via Status
func PutNodeFieldRevisionViaStatus(Status_ int, iNodeFieldRevision *NodeFieldRevision) (int64, error) {
	row, err := Engine.Update(iNodeFieldRevision, &NodeFieldRevision{Status: Status_})
	return row, err
}

// PutNodeFieldRevisionViaCreated Put NodeFieldRevision via Created
func PutNodeFieldRevisionViaCreated(Created_ int, iNodeFieldRevision *NodeFieldRevision) (int64, error) {
	row, err := Engine.Update(iNodeFieldRevision, &NodeFieldRevision{Created: Created_})
	return row, err
}

// PutNodeFieldRevisionViaChanged Put NodeFieldRevision via Changed
func PutNodeFieldRevisionViaChanged(Changed_ int, iNodeFieldRevision *NodeFieldRevision) (int64, error) {
	row, err := Engine.Update(iNodeFieldRevision, &NodeFieldRevision{Changed: Changed_})
	return row, err
}

// PutNodeFieldRevisionViaPromote Put NodeFieldRevision via Promote
func PutNodeFieldRevisionViaPromote(Promote_ int, iNodeFieldRevision *NodeFieldRevision) (int64, error) {
	row, err := Engine.Update(iNodeFieldRevision, &NodeFieldRevision{Promote: Promote_})
	return row, err
}

// PutNodeFieldRevisionViaSticky Put NodeFieldRevision via Sticky
func PutNodeFieldRevisionViaSticky(Sticky_ int, iNodeFieldRevision *NodeFieldRevision) (int64, error) {
	row, err := Engine.Update(iNodeFieldRevision, &NodeFieldRevision{Sticky: Sticky_})
	return row, err
}

// PutNodeFieldRevisionViaRevisionTranslationAffected Put NodeFieldRevision via RevisionTranslationAffected
func PutNodeFieldRevisionViaRevisionTranslationAffected(RevisionTranslationAffected_ int, iNodeFieldRevision *NodeFieldRevision) (int64, error) {
	row, err := Engine.Update(iNodeFieldRevision, &NodeFieldRevision{RevisionTranslationAffected: RevisionTranslationAffected_})
	return row, err
}

// PutNodeFieldRevisionViaDefaultLangcode Put NodeFieldRevision via DefaultLangcode
func PutNodeFieldRevisionViaDefaultLangcode(DefaultLangcode_ int, iNodeFieldRevision *NodeFieldRevision) (int64, error) {
	row, err := Engine.Update(iNodeFieldRevision, &NodeFieldRevision{DefaultLangcode: DefaultLangcode_})
	return row, err
}

// UpdateNodeFieldRevisionViaNid via map[string]interface{}{}
func UpdateNodeFieldRevisionViaNid(iNid int64, iNodeFieldRevisionMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(NodeFieldRevision)).Where("nid = ?", iNid).Update(iNodeFieldRevisionMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateNodeFieldRevisionViaVid via map[string]interface{}{}
func UpdateNodeFieldRevisionViaVid(iVid int, iNodeFieldRevisionMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(NodeFieldRevision)).Where("vid = ?", iVid).Update(iNodeFieldRevisionMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateNodeFieldRevisionViaLangcode via map[string]interface{}{}
func UpdateNodeFieldRevisionViaLangcode(iLangcode string, iNodeFieldRevisionMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(NodeFieldRevision)).Where("langcode = ?", iLangcode).Update(iNodeFieldRevisionMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateNodeFieldRevisionViaTitle via map[string]interface{}{}
func UpdateNodeFieldRevisionViaTitle(iTitle string, iNodeFieldRevisionMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(NodeFieldRevision)).Where("title = ?", iTitle).Update(iNodeFieldRevisionMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateNodeFieldRevisionViaUid via map[string]interface{}{}
func UpdateNodeFieldRevisionViaUid(iUid int, iNodeFieldRevisionMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(NodeFieldRevision)).Where("uid = ?", iUid).Update(iNodeFieldRevisionMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateNodeFieldRevisionViaStatus via map[string]interface{}{}
func UpdateNodeFieldRevisionViaStatus(iStatus int, iNodeFieldRevisionMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(NodeFieldRevision)).Where("status = ?", iStatus).Update(iNodeFieldRevisionMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateNodeFieldRevisionViaCreated via map[string]interface{}{}
func UpdateNodeFieldRevisionViaCreated(iCreated int, iNodeFieldRevisionMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(NodeFieldRevision)).Where("created = ?", iCreated).Update(iNodeFieldRevisionMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateNodeFieldRevisionViaChanged via map[string]interface{}{}
func UpdateNodeFieldRevisionViaChanged(iChanged int, iNodeFieldRevisionMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(NodeFieldRevision)).Where("changed = ?", iChanged).Update(iNodeFieldRevisionMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateNodeFieldRevisionViaPromote via map[string]interface{}{}
func UpdateNodeFieldRevisionViaPromote(iPromote int, iNodeFieldRevisionMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(NodeFieldRevision)).Where("promote = ?", iPromote).Update(iNodeFieldRevisionMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateNodeFieldRevisionViaSticky via map[string]interface{}{}
func UpdateNodeFieldRevisionViaSticky(iSticky int, iNodeFieldRevisionMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(NodeFieldRevision)).Where("sticky = ?", iSticky).Update(iNodeFieldRevisionMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateNodeFieldRevisionViaRevisionTranslationAffected via map[string]interface{}{}
func UpdateNodeFieldRevisionViaRevisionTranslationAffected(iRevisionTranslationAffected int, iNodeFieldRevisionMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(NodeFieldRevision)).Where("revision_translation_affected = ?", iRevisionTranslationAffected).Update(iNodeFieldRevisionMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateNodeFieldRevisionViaDefaultLangcode via map[string]interface{}{}
func UpdateNodeFieldRevisionViaDefaultLangcode(iDefaultLangcode int, iNodeFieldRevisionMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(NodeFieldRevision)).Where("default_langcode = ?", iDefaultLangcode).Update(iNodeFieldRevisionMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// DeleteNodeFieldRevision Delete NodeFieldRevision
func DeleteNodeFieldRevision(iNid int64) (int64, error) {
	row, err := Engine.Id(iNid).Delete(new(NodeFieldRevision))
	return row, err
}

// DeleteNodeFieldRevisionViaNid Delete NodeFieldRevision via Nid
func DeleteNodeFieldRevisionViaNid(iNid int64) (err error) {
	var has bool
	var _NodeFieldRevision = &NodeFieldRevision{Nid: iNid}
	if has, err = Engine.Get(_NodeFieldRevision); (has == true) && (err == nil) {
		if row, err := Engine.Where("nid = ?", iNid).Delete(new(NodeFieldRevision)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteNodeFieldRevisionViaVid Delete NodeFieldRevision via Vid
func DeleteNodeFieldRevisionViaVid(iVid int) (err error) {
	var has bool
	var _NodeFieldRevision = &NodeFieldRevision{Vid: iVid}
	if has, err = Engine.Get(_NodeFieldRevision); (has == true) && (err == nil) {
		if row, err := Engine.Where("vid = ?", iVid).Delete(new(NodeFieldRevision)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteNodeFieldRevisionViaLangcode Delete NodeFieldRevision via Langcode
func DeleteNodeFieldRevisionViaLangcode(iLangcode string) (err error) {
	var has bool
	var _NodeFieldRevision = &NodeFieldRevision{Langcode: iLangcode}
	if has, err = Engine.Get(_NodeFieldRevision); (has == true) && (err == nil) {
		if row, err := Engine.Where("langcode = ?", iLangcode).Delete(new(NodeFieldRevision)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteNodeFieldRevisionViaTitle Delete NodeFieldRevision via Title
func DeleteNodeFieldRevisionViaTitle(iTitle string) (err error) {
	var has bool
	var _NodeFieldRevision = &NodeFieldRevision{Title: iTitle}
	if has, err = Engine.Get(_NodeFieldRevision); (has == true) && (err == nil) {
		if row, err := Engine.Where("title = ?", iTitle).Delete(new(NodeFieldRevision)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteNodeFieldRevisionViaUid Delete NodeFieldRevision via Uid
func DeleteNodeFieldRevisionViaUid(iUid int) (err error) {
	var has bool
	var _NodeFieldRevision = &NodeFieldRevision{Uid: iUid}
	if has, err = Engine.Get(_NodeFieldRevision); (has == true) && (err == nil) {
		if row, err := Engine.Where("uid = ?", iUid).Delete(new(NodeFieldRevision)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteNodeFieldRevisionViaStatus Delete NodeFieldRevision via Status
func DeleteNodeFieldRevisionViaStatus(iStatus int) (err error) {
	var has bool
	var _NodeFieldRevision = &NodeFieldRevision{Status: iStatus}
	if has, err = Engine.Get(_NodeFieldRevision); (has == true) && (err == nil) {
		if row, err := Engine.Where("status = ?", iStatus).Delete(new(NodeFieldRevision)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteNodeFieldRevisionViaCreated Delete NodeFieldRevision via Created
func DeleteNodeFieldRevisionViaCreated(iCreated int) (err error) {
	var has bool
	var _NodeFieldRevision = &NodeFieldRevision{Created: iCreated}
	if has, err = Engine.Get(_NodeFieldRevision); (has == true) && (err == nil) {
		if row, err := Engine.Where("created = ?", iCreated).Delete(new(NodeFieldRevision)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteNodeFieldRevisionViaChanged Delete NodeFieldRevision via Changed
func DeleteNodeFieldRevisionViaChanged(iChanged int) (err error) {
	var has bool
	var _NodeFieldRevision = &NodeFieldRevision{Changed: iChanged}
	if has, err = Engine.Get(_NodeFieldRevision); (has == true) && (err == nil) {
		if row, err := Engine.Where("changed = ?", iChanged).Delete(new(NodeFieldRevision)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteNodeFieldRevisionViaPromote Delete NodeFieldRevision via Promote
func DeleteNodeFieldRevisionViaPromote(iPromote int) (err error) {
	var has bool
	var _NodeFieldRevision = &NodeFieldRevision{Promote: iPromote}
	if has, err = Engine.Get(_NodeFieldRevision); (has == true) && (err == nil) {
		if row, err := Engine.Where("promote = ?", iPromote).Delete(new(NodeFieldRevision)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteNodeFieldRevisionViaSticky Delete NodeFieldRevision via Sticky
func DeleteNodeFieldRevisionViaSticky(iSticky int) (err error) {
	var has bool
	var _NodeFieldRevision = &NodeFieldRevision{Sticky: iSticky}
	if has, err = Engine.Get(_NodeFieldRevision); (has == true) && (err == nil) {
		if row, err := Engine.Where("sticky = ?", iSticky).Delete(new(NodeFieldRevision)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteNodeFieldRevisionViaRevisionTranslationAffected Delete NodeFieldRevision via RevisionTranslationAffected
func DeleteNodeFieldRevisionViaRevisionTranslationAffected(iRevisionTranslationAffected int) (err error) {
	var has bool
	var _NodeFieldRevision = &NodeFieldRevision{RevisionTranslationAffected: iRevisionTranslationAffected}
	if has, err = Engine.Get(_NodeFieldRevision); (has == true) && (err == nil) {
		if row, err := Engine.Where("revision_translation_affected = ?", iRevisionTranslationAffected).Delete(new(NodeFieldRevision)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteNodeFieldRevisionViaDefaultLangcode Delete NodeFieldRevision via DefaultLangcode
func DeleteNodeFieldRevisionViaDefaultLangcode(iDefaultLangcode int) (err error) {
	var has bool
	var _NodeFieldRevision = &NodeFieldRevision{DefaultLangcode: iDefaultLangcode}
	if has, err = Engine.Get(_NodeFieldRevision); (has == true) && (err == nil) {
		if row, err := Engine.Where("default_langcode = ?", iDefaultLangcode).Delete(new(NodeFieldRevision)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}
