package model

type BlockContentFieldRevision struct {
	Id                          int64  `xorm:"not null index(block_content__id__default_langcode__langcode) INT(10)"`
	RevisionId                  int    `xorm:"not null pk INT(10)"`
	Langcode                    string `xorm:"not null pk index(block_content__id__default_langcode__langcode) VARCHAR(12)"`
	Info                        string `xorm:"VARCHAR(255)"`
	Changed                     int    `xorm:"INT(11)"`
	RevisionCreated             int    `xorm:"INT(11)"`
	RevisionUser                int    `xorm:"index INT(10)"`
	RevisionTranslationAffected int    `xorm:"TINYINT(4)"`
	DefaultLangcode             int    `xorm:"not null index(block_content__id__default_langcode__langcode) TINYINT(4)"`
}

// GetBlockContentFieldRevisionsCount BlockContentFieldRevisions' Count
func GetBlockContentFieldRevisionsCount(offset int, limit int) (int64, error) {
	total, err := Engine.Limit(limit, offset).Count(&BlockContentFieldRevision{})
	return total, err
}

// GetBlockContentFieldRevisionCountViaId Get BlockContentFieldRevision via Id
func GetBlockContentFieldRevisionCountViaId(iId int64) int64 {
	n, _ := Engine.Where("id = ?", iId).Count(&BlockContentFieldRevision{Id: iId})
	return n
}

// GetBlockContentFieldRevisionCountViaRevisionId Get BlockContentFieldRevision via RevisionId
func GetBlockContentFieldRevisionCountViaRevisionId(iRevisionId int) int64 {
	n, _ := Engine.Where("revision_id = ?", iRevisionId).Count(&BlockContentFieldRevision{RevisionId: iRevisionId})
	return n
}

// GetBlockContentFieldRevisionCountViaLangcode Get BlockContentFieldRevision via Langcode
func GetBlockContentFieldRevisionCountViaLangcode(iLangcode string) int64 {
	n, _ := Engine.Where("langcode = ?", iLangcode).Count(&BlockContentFieldRevision{Langcode: iLangcode})
	return n
}

// GetBlockContentFieldRevisionCountViaInfo Get BlockContentFieldRevision via Info
func GetBlockContentFieldRevisionCountViaInfo(iInfo string) int64 {
	n, _ := Engine.Where("info = ?", iInfo).Count(&BlockContentFieldRevision{Info: iInfo})
	return n
}

// GetBlockContentFieldRevisionCountViaChanged Get BlockContentFieldRevision via Changed
func GetBlockContentFieldRevisionCountViaChanged(iChanged int) int64 {
	n, _ := Engine.Where("changed = ?", iChanged).Count(&BlockContentFieldRevision{Changed: iChanged})
	return n
}

// GetBlockContentFieldRevisionCountViaRevisionCreated Get BlockContentFieldRevision via RevisionCreated
func GetBlockContentFieldRevisionCountViaRevisionCreated(iRevisionCreated int) int64 {
	n, _ := Engine.Where("revision_created = ?", iRevisionCreated).Count(&BlockContentFieldRevision{RevisionCreated: iRevisionCreated})
	return n
}

// GetBlockContentFieldRevisionCountViaRevisionUser Get BlockContentFieldRevision via RevisionUser
func GetBlockContentFieldRevisionCountViaRevisionUser(iRevisionUser int) int64 {
	n, _ := Engine.Where("revision_user = ?", iRevisionUser).Count(&BlockContentFieldRevision{RevisionUser: iRevisionUser})
	return n
}

// GetBlockContentFieldRevisionCountViaRevisionTranslationAffected Get BlockContentFieldRevision via RevisionTranslationAffected
func GetBlockContentFieldRevisionCountViaRevisionTranslationAffected(iRevisionTranslationAffected int) int64 {
	n, _ := Engine.Where("revision_translation_affected = ?", iRevisionTranslationAffected).Count(&BlockContentFieldRevision{RevisionTranslationAffected: iRevisionTranslationAffected})
	return n
}

// GetBlockContentFieldRevisionCountViaDefaultLangcode Get BlockContentFieldRevision via DefaultLangcode
func GetBlockContentFieldRevisionCountViaDefaultLangcode(iDefaultLangcode int) int64 {
	n, _ := Engine.Where("default_langcode = ?", iDefaultLangcode).Count(&BlockContentFieldRevision{DefaultLangcode: iDefaultLangcode})
	return n
}

// GetBlockContentFieldRevisionsByIdAndRevisionIdAndLangcode Get BlockContentFieldRevisions via IdAndRevisionIdAndLangcode
func GetBlockContentFieldRevisionsByIdAndRevisionIdAndLangcode(offset int, limit int, Id_ int64, RevisionId_ int, Langcode_ string) (*[]*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = new([]*BlockContentFieldRevision)
	err := Engine.Table("block_content_field_revision").Where("id = ? and revision_id = ? and langcode = ?", Id_, RevisionId_, Langcode_).Limit(limit, offset).Find(_BlockContentFieldRevision)
	return _BlockContentFieldRevision, err
}

// GetBlockContentFieldRevisionsByIdAndRevisionIdAndInfo Get BlockContentFieldRevisions via IdAndRevisionIdAndInfo
func GetBlockContentFieldRevisionsByIdAndRevisionIdAndInfo(offset int, limit int, Id_ int64, RevisionId_ int, Info_ string) (*[]*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = new([]*BlockContentFieldRevision)
	err := Engine.Table("block_content_field_revision").Where("id = ? and revision_id = ? and info = ?", Id_, RevisionId_, Info_).Limit(limit, offset).Find(_BlockContentFieldRevision)
	return _BlockContentFieldRevision, err
}

// GetBlockContentFieldRevisionsByIdAndRevisionIdAndChanged Get BlockContentFieldRevisions via IdAndRevisionIdAndChanged
func GetBlockContentFieldRevisionsByIdAndRevisionIdAndChanged(offset int, limit int, Id_ int64, RevisionId_ int, Changed_ int) (*[]*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = new([]*BlockContentFieldRevision)
	err := Engine.Table("block_content_field_revision").Where("id = ? and revision_id = ? and changed = ?", Id_, RevisionId_, Changed_).Limit(limit, offset).Find(_BlockContentFieldRevision)
	return _BlockContentFieldRevision, err
}

// GetBlockContentFieldRevisionsByIdAndRevisionIdAndRevisionCreated Get BlockContentFieldRevisions via IdAndRevisionIdAndRevisionCreated
func GetBlockContentFieldRevisionsByIdAndRevisionIdAndRevisionCreated(offset int, limit int, Id_ int64, RevisionId_ int, RevisionCreated_ int) (*[]*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = new([]*BlockContentFieldRevision)
	err := Engine.Table("block_content_field_revision").Where("id = ? and revision_id = ? and revision_created = ?", Id_, RevisionId_, RevisionCreated_).Limit(limit, offset).Find(_BlockContentFieldRevision)
	return _BlockContentFieldRevision, err
}

// GetBlockContentFieldRevisionsByIdAndRevisionIdAndRevisionUser Get BlockContentFieldRevisions via IdAndRevisionIdAndRevisionUser
func GetBlockContentFieldRevisionsByIdAndRevisionIdAndRevisionUser(offset int, limit int, Id_ int64, RevisionId_ int, RevisionUser_ int) (*[]*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = new([]*BlockContentFieldRevision)
	err := Engine.Table("block_content_field_revision").Where("id = ? and revision_id = ? and revision_user = ?", Id_, RevisionId_, RevisionUser_).Limit(limit, offset).Find(_BlockContentFieldRevision)
	return _BlockContentFieldRevision, err
}

// GetBlockContentFieldRevisionsByIdAndRevisionIdAndRevisionTranslationAffected Get BlockContentFieldRevisions via IdAndRevisionIdAndRevisionTranslationAffected
func GetBlockContentFieldRevisionsByIdAndRevisionIdAndRevisionTranslationAffected(offset int, limit int, Id_ int64, RevisionId_ int, RevisionTranslationAffected_ int) (*[]*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = new([]*BlockContentFieldRevision)
	err := Engine.Table("block_content_field_revision").Where("id = ? and revision_id = ? and revision_translation_affected = ?", Id_, RevisionId_, RevisionTranslationAffected_).Limit(limit, offset).Find(_BlockContentFieldRevision)
	return _BlockContentFieldRevision, err
}

// GetBlockContentFieldRevisionsByIdAndRevisionIdAndDefaultLangcode Get BlockContentFieldRevisions via IdAndRevisionIdAndDefaultLangcode
func GetBlockContentFieldRevisionsByIdAndRevisionIdAndDefaultLangcode(offset int, limit int, Id_ int64, RevisionId_ int, DefaultLangcode_ int) (*[]*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = new([]*BlockContentFieldRevision)
	err := Engine.Table("block_content_field_revision").Where("id = ? and revision_id = ? and default_langcode = ?", Id_, RevisionId_, DefaultLangcode_).Limit(limit, offset).Find(_BlockContentFieldRevision)
	return _BlockContentFieldRevision, err
}

// GetBlockContentFieldRevisionsByIdAndLangcodeAndInfo Get BlockContentFieldRevisions via IdAndLangcodeAndInfo
func GetBlockContentFieldRevisionsByIdAndLangcodeAndInfo(offset int, limit int, Id_ int64, Langcode_ string, Info_ string) (*[]*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = new([]*BlockContentFieldRevision)
	err := Engine.Table("block_content_field_revision").Where("id = ? and langcode = ? and info = ?", Id_, Langcode_, Info_).Limit(limit, offset).Find(_BlockContentFieldRevision)
	return _BlockContentFieldRevision, err
}

// GetBlockContentFieldRevisionsByIdAndLangcodeAndChanged Get BlockContentFieldRevisions via IdAndLangcodeAndChanged
func GetBlockContentFieldRevisionsByIdAndLangcodeAndChanged(offset int, limit int, Id_ int64, Langcode_ string, Changed_ int) (*[]*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = new([]*BlockContentFieldRevision)
	err := Engine.Table("block_content_field_revision").Where("id = ? and langcode = ? and changed = ?", Id_, Langcode_, Changed_).Limit(limit, offset).Find(_BlockContentFieldRevision)
	return _BlockContentFieldRevision, err
}

// GetBlockContentFieldRevisionsByIdAndLangcodeAndRevisionCreated Get BlockContentFieldRevisions via IdAndLangcodeAndRevisionCreated
func GetBlockContentFieldRevisionsByIdAndLangcodeAndRevisionCreated(offset int, limit int, Id_ int64, Langcode_ string, RevisionCreated_ int) (*[]*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = new([]*BlockContentFieldRevision)
	err := Engine.Table("block_content_field_revision").Where("id = ? and langcode = ? and revision_created = ?", Id_, Langcode_, RevisionCreated_).Limit(limit, offset).Find(_BlockContentFieldRevision)
	return _BlockContentFieldRevision, err
}

// GetBlockContentFieldRevisionsByIdAndLangcodeAndRevisionUser Get BlockContentFieldRevisions via IdAndLangcodeAndRevisionUser
func GetBlockContentFieldRevisionsByIdAndLangcodeAndRevisionUser(offset int, limit int, Id_ int64, Langcode_ string, RevisionUser_ int) (*[]*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = new([]*BlockContentFieldRevision)
	err := Engine.Table("block_content_field_revision").Where("id = ? and langcode = ? and revision_user = ?", Id_, Langcode_, RevisionUser_).Limit(limit, offset).Find(_BlockContentFieldRevision)
	return _BlockContentFieldRevision, err
}

// GetBlockContentFieldRevisionsByIdAndLangcodeAndRevisionTranslationAffected Get BlockContentFieldRevisions via IdAndLangcodeAndRevisionTranslationAffected
func GetBlockContentFieldRevisionsByIdAndLangcodeAndRevisionTranslationAffected(offset int, limit int, Id_ int64, Langcode_ string, RevisionTranslationAffected_ int) (*[]*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = new([]*BlockContentFieldRevision)
	err := Engine.Table("block_content_field_revision").Where("id = ? and langcode = ? and revision_translation_affected = ?", Id_, Langcode_, RevisionTranslationAffected_).Limit(limit, offset).Find(_BlockContentFieldRevision)
	return _BlockContentFieldRevision, err
}

// GetBlockContentFieldRevisionsByIdAndLangcodeAndDefaultLangcode Get BlockContentFieldRevisions via IdAndLangcodeAndDefaultLangcode
func GetBlockContentFieldRevisionsByIdAndLangcodeAndDefaultLangcode(offset int, limit int, Id_ int64, Langcode_ string, DefaultLangcode_ int) (*[]*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = new([]*BlockContentFieldRevision)
	err := Engine.Table("block_content_field_revision").Where("id = ? and langcode = ? and default_langcode = ?", Id_, Langcode_, DefaultLangcode_).Limit(limit, offset).Find(_BlockContentFieldRevision)
	return _BlockContentFieldRevision, err
}

// GetBlockContentFieldRevisionsByIdAndInfoAndChanged Get BlockContentFieldRevisions via IdAndInfoAndChanged
func GetBlockContentFieldRevisionsByIdAndInfoAndChanged(offset int, limit int, Id_ int64, Info_ string, Changed_ int) (*[]*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = new([]*BlockContentFieldRevision)
	err := Engine.Table("block_content_field_revision").Where("id = ? and info = ? and changed = ?", Id_, Info_, Changed_).Limit(limit, offset).Find(_BlockContentFieldRevision)
	return _BlockContentFieldRevision, err
}

// GetBlockContentFieldRevisionsByIdAndInfoAndRevisionCreated Get BlockContentFieldRevisions via IdAndInfoAndRevisionCreated
func GetBlockContentFieldRevisionsByIdAndInfoAndRevisionCreated(offset int, limit int, Id_ int64, Info_ string, RevisionCreated_ int) (*[]*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = new([]*BlockContentFieldRevision)
	err := Engine.Table("block_content_field_revision").Where("id = ? and info = ? and revision_created = ?", Id_, Info_, RevisionCreated_).Limit(limit, offset).Find(_BlockContentFieldRevision)
	return _BlockContentFieldRevision, err
}

// GetBlockContentFieldRevisionsByIdAndInfoAndRevisionUser Get BlockContentFieldRevisions via IdAndInfoAndRevisionUser
func GetBlockContentFieldRevisionsByIdAndInfoAndRevisionUser(offset int, limit int, Id_ int64, Info_ string, RevisionUser_ int) (*[]*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = new([]*BlockContentFieldRevision)
	err := Engine.Table("block_content_field_revision").Where("id = ? and info = ? and revision_user = ?", Id_, Info_, RevisionUser_).Limit(limit, offset).Find(_BlockContentFieldRevision)
	return _BlockContentFieldRevision, err
}

// GetBlockContentFieldRevisionsByIdAndInfoAndRevisionTranslationAffected Get BlockContentFieldRevisions via IdAndInfoAndRevisionTranslationAffected
func GetBlockContentFieldRevisionsByIdAndInfoAndRevisionTranslationAffected(offset int, limit int, Id_ int64, Info_ string, RevisionTranslationAffected_ int) (*[]*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = new([]*BlockContentFieldRevision)
	err := Engine.Table("block_content_field_revision").Where("id = ? and info = ? and revision_translation_affected = ?", Id_, Info_, RevisionTranslationAffected_).Limit(limit, offset).Find(_BlockContentFieldRevision)
	return _BlockContentFieldRevision, err
}

// GetBlockContentFieldRevisionsByIdAndInfoAndDefaultLangcode Get BlockContentFieldRevisions via IdAndInfoAndDefaultLangcode
func GetBlockContentFieldRevisionsByIdAndInfoAndDefaultLangcode(offset int, limit int, Id_ int64, Info_ string, DefaultLangcode_ int) (*[]*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = new([]*BlockContentFieldRevision)
	err := Engine.Table("block_content_field_revision").Where("id = ? and info = ? and default_langcode = ?", Id_, Info_, DefaultLangcode_).Limit(limit, offset).Find(_BlockContentFieldRevision)
	return _BlockContentFieldRevision, err
}

// GetBlockContentFieldRevisionsByIdAndChangedAndRevisionCreated Get BlockContentFieldRevisions via IdAndChangedAndRevisionCreated
func GetBlockContentFieldRevisionsByIdAndChangedAndRevisionCreated(offset int, limit int, Id_ int64, Changed_ int, RevisionCreated_ int) (*[]*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = new([]*BlockContentFieldRevision)
	err := Engine.Table("block_content_field_revision").Where("id = ? and changed = ? and revision_created = ?", Id_, Changed_, RevisionCreated_).Limit(limit, offset).Find(_BlockContentFieldRevision)
	return _BlockContentFieldRevision, err
}

// GetBlockContentFieldRevisionsByIdAndChangedAndRevisionUser Get BlockContentFieldRevisions via IdAndChangedAndRevisionUser
func GetBlockContentFieldRevisionsByIdAndChangedAndRevisionUser(offset int, limit int, Id_ int64, Changed_ int, RevisionUser_ int) (*[]*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = new([]*BlockContentFieldRevision)
	err := Engine.Table("block_content_field_revision").Where("id = ? and changed = ? and revision_user = ?", Id_, Changed_, RevisionUser_).Limit(limit, offset).Find(_BlockContentFieldRevision)
	return _BlockContentFieldRevision, err
}

// GetBlockContentFieldRevisionsByIdAndChangedAndRevisionTranslationAffected Get BlockContentFieldRevisions via IdAndChangedAndRevisionTranslationAffected
func GetBlockContentFieldRevisionsByIdAndChangedAndRevisionTranslationAffected(offset int, limit int, Id_ int64, Changed_ int, RevisionTranslationAffected_ int) (*[]*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = new([]*BlockContentFieldRevision)
	err := Engine.Table("block_content_field_revision").Where("id = ? and changed = ? and revision_translation_affected = ?", Id_, Changed_, RevisionTranslationAffected_).Limit(limit, offset).Find(_BlockContentFieldRevision)
	return _BlockContentFieldRevision, err
}

// GetBlockContentFieldRevisionsByIdAndChangedAndDefaultLangcode Get BlockContentFieldRevisions via IdAndChangedAndDefaultLangcode
func GetBlockContentFieldRevisionsByIdAndChangedAndDefaultLangcode(offset int, limit int, Id_ int64, Changed_ int, DefaultLangcode_ int) (*[]*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = new([]*BlockContentFieldRevision)
	err := Engine.Table("block_content_field_revision").Where("id = ? and changed = ? and default_langcode = ?", Id_, Changed_, DefaultLangcode_).Limit(limit, offset).Find(_BlockContentFieldRevision)
	return _BlockContentFieldRevision, err
}

// GetBlockContentFieldRevisionsByIdAndRevisionCreatedAndRevisionUser Get BlockContentFieldRevisions via IdAndRevisionCreatedAndRevisionUser
func GetBlockContentFieldRevisionsByIdAndRevisionCreatedAndRevisionUser(offset int, limit int, Id_ int64, RevisionCreated_ int, RevisionUser_ int) (*[]*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = new([]*BlockContentFieldRevision)
	err := Engine.Table("block_content_field_revision").Where("id = ? and revision_created = ? and revision_user = ?", Id_, RevisionCreated_, RevisionUser_).Limit(limit, offset).Find(_BlockContentFieldRevision)
	return _BlockContentFieldRevision, err
}

// GetBlockContentFieldRevisionsByIdAndRevisionCreatedAndRevisionTranslationAffected Get BlockContentFieldRevisions via IdAndRevisionCreatedAndRevisionTranslationAffected
func GetBlockContentFieldRevisionsByIdAndRevisionCreatedAndRevisionTranslationAffected(offset int, limit int, Id_ int64, RevisionCreated_ int, RevisionTranslationAffected_ int) (*[]*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = new([]*BlockContentFieldRevision)
	err := Engine.Table("block_content_field_revision").Where("id = ? and revision_created = ? and revision_translation_affected = ?", Id_, RevisionCreated_, RevisionTranslationAffected_).Limit(limit, offset).Find(_BlockContentFieldRevision)
	return _BlockContentFieldRevision, err
}

// GetBlockContentFieldRevisionsByIdAndRevisionCreatedAndDefaultLangcode Get BlockContentFieldRevisions via IdAndRevisionCreatedAndDefaultLangcode
func GetBlockContentFieldRevisionsByIdAndRevisionCreatedAndDefaultLangcode(offset int, limit int, Id_ int64, RevisionCreated_ int, DefaultLangcode_ int) (*[]*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = new([]*BlockContentFieldRevision)
	err := Engine.Table("block_content_field_revision").Where("id = ? and revision_created = ? and default_langcode = ?", Id_, RevisionCreated_, DefaultLangcode_).Limit(limit, offset).Find(_BlockContentFieldRevision)
	return _BlockContentFieldRevision, err
}

// GetBlockContentFieldRevisionsByIdAndRevisionUserAndRevisionTranslationAffected Get BlockContentFieldRevisions via IdAndRevisionUserAndRevisionTranslationAffected
func GetBlockContentFieldRevisionsByIdAndRevisionUserAndRevisionTranslationAffected(offset int, limit int, Id_ int64, RevisionUser_ int, RevisionTranslationAffected_ int) (*[]*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = new([]*BlockContentFieldRevision)
	err := Engine.Table("block_content_field_revision").Where("id = ? and revision_user = ? and revision_translation_affected = ?", Id_, RevisionUser_, RevisionTranslationAffected_).Limit(limit, offset).Find(_BlockContentFieldRevision)
	return _BlockContentFieldRevision, err
}

// GetBlockContentFieldRevisionsByIdAndRevisionUserAndDefaultLangcode Get BlockContentFieldRevisions via IdAndRevisionUserAndDefaultLangcode
func GetBlockContentFieldRevisionsByIdAndRevisionUserAndDefaultLangcode(offset int, limit int, Id_ int64, RevisionUser_ int, DefaultLangcode_ int) (*[]*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = new([]*BlockContentFieldRevision)
	err := Engine.Table("block_content_field_revision").Where("id = ? and revision_user = ? and default_langcode = ?", Id_, RevisionUser_, DefaultLangcode_).Limit(limit, offset).Find(_BlockContentFieldRevision)
	return _BlockContentFieldRevision, err
}

// GetBlockContentFieldRevisionsByIdAndRevisionTranslationAffectedAndDefaultLangcode Get BlockContentFieldRevisions via IdAndRevisionTranslationAffectedAndDefaultLangcode
func GetBlockContentFieldRevisionsByIdAndRevisionTranslationAffectedAndDefaultLangcode(offset int, limit int, Id_ int64, RevisionTranslationAffected_ int, DefaultLangcode_ int) (*[]*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = new([]*BlockContentFieldRevision)
	err := Engine.Table("block_content_field_revision").Where("id = ? and revision_translation_affected = ? and default_langcode = ?", Id_, RevisionTranslationAffected_, DefaultLangcode_).Limit(limit, offset).Find(_BlockContentFieldRevision)
	return _BlockContentFieldRevision, err
}

// GetBlockContentFieldRevisionsByRevisionIdAndLangcodeAndInfo Get BlockContentFieldRevisions via RevisionIdAndLangcodeAndInfo
func GetBlockContentFieldRevisionsByRevisionIdAndLangcodeAndInfo(offset int, limit int, RevisionId_ int, Langcode_ string, Info_ string) (*[]*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = new([]*BlockContentFieldRevision)
	err := Engine.Table("block_content_field_revision").Where("revision_id = ? and langcode = ? and info = ?", RevisionId_, Langcode_, Info_).Limit(limit, offset).Find(_BlockContentFieldRevision)
	return _BlockContentFieldRevision, err
}

// GetBlockContentFieldRevisionsByRevisionIdAndLangcodeAndChanged Get BlockContentFieldRevisions via RevisionIdAndLangcodeAndChanged
func GetBlockContentFieldRevisionsByRevisionIdAndLangcodeAndChanged(offset int, limit int, RevisionId_ int, Langcode_ string, Changed_ int) (*[]*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = new([]*BlockContentFieldRevision)
	err := Engine.Table("block_content_field_revision").Where("revision_id = ? and langcode = ? and changed = ?", RevisionId_, Langcode_, Changed_).Limit(limit, offset).Find(_BlockContentFieldRevision)
	return _BlockContentFieldRevision, err
}

// GetBlockContentFieldRevisionsByRevisionIdAndLangcodeAndRevisionCreated Get BlockContentFieldRevisions via RevisionIdAndLangcodeAndRevisionCreated
func GetBlockContentFieldRevisionsByRevisionIdAndLangcodeAndRevisionCreated(offset int, limit int, RevisionId_ int, Langcode_ string, RevisionCreated_ int) (*[]*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = new([]*BlockContentFieldRevision)
	err := Engine.Table("block_content_field_revision").Where("revision_id = ? and langcode = ? and revision_created = ?", RevisionId_, Langcode_, RevisionCreated_).Limit(limit, offset).Find(_BlockContentFieldRevision)
	return _BlockContentFieldRevision, err
}

// GetBlockContentFieldRevisionsByRevisionIdAndLangcodeAndRevisionUser Get BlockContentFieldRevisions via RevisionIdAndLangcodeAndRevisionUser
func GetBlockContentFieldRevisionsByRevisionIdAndLangcodeAndRevisionUser(offset int, limit int, RevisionId_ int, Langcode_ string, RevisionUser_ int) (*[]*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = new([]*BlockContentFieldRevision)
	err := Engine.Table("block_content_field_revision").Where("revision_id = ? and langcode = ? and revision_user = ?", RevisionId_, Langcode_, RevisionUser_).Limit(limit, offset).Find(_BlockContentFieldRevision)
	return _BlockContentFieldRevision, err
}

// GetBlockContentFieldRevisionsByRevisionIdAndLangcodeAndRevisionTranslationAffected Get BlockContentFieldRevisions via RevisionIdAndLangcodeAndRevisionTranslationAffected
func GetBlockContentFieldRevisionsByRevisionIdAndLangcodeAndRevisionTranslationAffected(offset int, limit int, RevisionId_ int, Langcode_ string, RevisionTranslationAffected_ int) (*[]*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = new([]*BlockContentFieldRevision)
	err := Engine.Table("block_content_field_revision").Where("revision_id = ? and langcode = ? and revision_translation_affected = ?", RevisionId_, Langcode_, RevisionTranslationAffected_).Limit(limit, offset).Find(_BlockContentFieldRevision)
	return _BlockContentFieldRevision, err
}

// GetBlockContentFieldRevisionsByRevisionIdAndLangcodeAndDefaultLangcode Get BlockContentFieldRevisions via RevisionIdAndLangcodeAndDefaultLangcode
func GetBlockContentFieldRevisionsByRevisionIdAndLangcodeAndDefaultLangcode(offset int, limit int, RevisionId_ int, Langcode_ string, DefaultLangcode_ int) (*[]*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = new([]*BlockContentFieldRevision)
	err := Engine.Table("block_content_field_revision").Where("revision_id = ? and langcode = ? and default_langcode = ?", RevisionId_, Langcode_, DefaultLangcode_).Limit(limit, offset).Find(_BlockContentFieldRevision)
	return _BlockContentFieldRevision, err
}

// GetBlockContentFieldRevisionsByRevisionIdAndInfoAndChanged Get BlockContentFieldRevisions via RevisionIdAndInfoAndChanged
func GetBlockContentFieldRevisionsByRevisionIdAndInfoAndChanged(offset int, limit int, RevisionId_ int, Info_ string, Changed_ int) (*[]*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = new([]*BlockContentFieldRevision)
	err := Engine.Table("block_content_field_revision").Where("revision_id = ? and info = ? and changed = ?", RevisionId_, Info_, Changed_).Limit(limit, offset).Find(_BlockContentFieldRevision)
	return _BlockContentFieldRevision, err
}

// GetBlockContentFieldRevisionsByRevisionIdAndInfoAndRevisionCreated Get BlockContentFieldRevisions via RevisionIdAndInfoAndRevisionCreated
func GetBlockContentFieldRevisionsByRevisionIdAndInfoAndRevisionCreated(offset int, limit int, RevisionId_ int, Info_ string, RevisionCreated_ int) (*[]*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = new([]*BlockContentFieldRevision)
	err := Engine.Table("block_content_field_revision").Where("revision_id = ? and info = ? and revision_created = ?", RevisionId_, Info_, RevisionCreated_).Limit(limit, offset).Find(_BlockContentFieldRevision)
	return _BlockContentFieldRevision, err
}

// GetBlockContentFieldRevisionsByRevisionIdAndInfoAndRevisionUser Get BlockContentFieldRevisions via RevisionIdAndInfoAndRevisionUser
func GetBlockContentFieldRevisionsByRevisionIdAndInfoAndRevisionUser(offset int, limit int, RevisionId_ int, Info_ string, RevisionUser_ int) (*[]*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = new([]*BlockContentFieldRevision)
	err := Engine.Table("block_content_field_revision").Where("revision_id = ? and info = ? and revision_user = ?", RevisionId_, Info_, RevisionUser_).Limit(limit, offset).Find(_BlockContentFieldRevision)
	return _BlockContentFieldRevision, err
}

// GetBlockContentFieldRevisionsByRevisionIdAndInfoAndRevisionTranslationAffected Get BlockContentFieldRevisions via RevisionIdAndInfoAndRevisionTranslationAffected
func GetBlockContentFieldRevisionsByRevisionIdAndInfoAndRevisionTranslationAffected(offset int, limit int, RevisionId_ int, Info_ string, RevisionTranslationAffected_ int) (*[]*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = new([]*BlockContentFieldRevision)
	err := Engine.Table("block_content_field_revision").Where("revision_id = ? and info = ? and revision_translation_affected = ?", RevisionId_, Info_, RevisionTranslationAffected_).Limit(limit, offset).Find(_BlockContentFieldRevision)
	return _BlockContentFieldRevision, err
}

// GetBlockContentFieldRevisionsByRevisionIdAndInfoAndDefaultLangcode Get BlockContentFieldRevisions via RevisionIdAndInfoAndDefaultLangcode
func GetBlockContentFieldRevisionsByRevisionIdAndInfoAndDefaultLangcode(offset int, limit int, RevisionId_ int, Info_ string, DefaultLangcode_ int) (*[]*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = new([]*BlockContentFieldRevision)
	err := Engine.Table("block_content_field_revision").Where("revision_id = ? and info = ? and default_langcode = ?", RevisionId_, Info_, DefaultLangcode_).Limit(limit, offset).Find(_BlockContentFieldRevision)
	return _BlockContentFieldRevision, err
}

// GetBlockContentFieldRevisionsByRevisionIdAndChangedAndRevisionCreated Get BlockContentFieldRevisions via RevisionIdAndChangedAndRevisionCreated
func GetBlockContentFieldRevisionsByRevisionIdAndChangedAndRevisionCreated(offset int, limit int, RevisionId_ int, Changed_ int, RevisionCreated_ int) (*[]*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = new([]*BlockContentFieldRevision)
	err := Engine.Table("block_content_field_revision").Where("revision_id = ? and changed = ? and revision_created = ?", RevisionId_, Changed_, RevisionCreated_).Limit(limit, offset).Find(_BlockContentFieldRevision)
	return _BlockContentFieldRevision, err
}

// GetBlockContentFieldRevisionsByRevisionIdAndChangedAndRevisionUser Get BlockContentFieldRevisions via RevisionIdAndChangedAndRevisionUser
func GetBlockContentFieldRevisionsByRevisionIdAndChangedAndRevisionUser(offset int, limit int, RevisionId_ int, Changed_ int, RevisionUser_ int) (*[]*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = new([]*BlockContentFieldRevision)
	err := Engine.Table("block_content_field_revision").Where("revision_id = ? and changed = ? and revision_user = ?", RevisionId_, Changed_, RevisionUser_).Limit(limit, offset).Find(_BlockContentFieldRevision)
	return _BlockContentFieldRevision, err
}

// GetBlockContentFieldRevisionsByRevisionIdAndChangedAndRevisionTranslationAffected Get BlockContentFieldRevisions via RevisionIdAndChangedAndRevisionTranslationAffected
func GetBlockContentFieldRevisionsByRevisionIdAndChangedAndRevisionTranslationAffected(offset int, limit int, RevisionId_ int, Changed_ int, RevisionTranslationAffected_ int) (*[]*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = new([]*BlockContentFieldRevision)
	err := Engine.Table("block_content_field_revision").Where("revision_id = ? and changed = ? and revision_translation_affected = ?", RevisionId_, Changed_, RevisionTranslationAffected_).Limit(limit, offset).Find(_BlockContentFieldRevision)
	return _BlockContentFieldRevision, err
}

// GetBlockContentFieldRevisionsByRevisionIdAndChangedAndDefaultLangcode Get BlockContentFieldRevisions via RevisionIdAndChangedAndDefaultLangcode
func GetBlockContentFieldRevisionsByRevisionIdAndChangedAndDefaultLangcode(offset int, limit int, RevisionId_ int, Changed_ int, DefaultLangcode_ int) (*[]*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = new([]*BlockContentFieldRevision)
	err := Engine.Table("block_content_field_revision").Where("revision_id = ? and changed = ? and default_langcode = ?", RevisionId_, Changed_, DefaultLangcode_).Limit(limit, offset).Find(_BlockContentFieldRevision)
	return _BlockContentFieldRevision, err
}

// GetBlockContentFieldRevisionsByRevisionIdAndRevisionCreatedAndRevisionUser Get BlockContentFieldRevisions via RevisionIdAndRevisionCreatedAndRevisionUser
func GetBlockContentFieldRevisionsByRevisionIdAndRevisionCreatedAndRevisionUser(offset int, limit int, RevisionId_ int, RevisionCreated_ int, RevisionUser_ int) (*[]*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = new([]*BlockContentFieldRevision)
	err := Engine.Table("block_content_field_revision").Where("revision_id = ? and revision_created = ? and revision_user = ?", RevisionId_, RevisionCreated_, RevisionUser_).Limit(limit, offset).Find(_BlockContentFieldRevision)
	return _BlockContentFieldRevision, err
}

// GetBlockContentFieldRevisionsByRevisionIdAndRevisionCreatedAndRevisionTranslationAffected Get BlockContentFieldRevisions via RevisionIdAndRevisionCreatedAndRevisionTranslationAffected
func GetBlockContentFieldRevisionsByRevisionIdAndRevisionCreatedAndRevisionTranslationAffected(offset int, limit int, RevisionId_ int, RevisionCreated_ int, RevisionTranslationAffected_ int) (*[]*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = new([]*BlockContentFieldRevision)
	err := Engine.Table("block_content_field_revision").Where("revision_id = ? and revision_created = ? and revision_translation_affected = ?", RevisionId_, RevisionCreated_, RevisionTranslationAffected_).Limit(limit, offset).Find(_BlockContentFieldRevision)
	return _BlockContentFieldRevision, err
}

// GetBlockContentFieldRevisionsByRevisionIdAndRevisionCreatedAndDefaultLangcode Get BlockContentFieldRevisions via RevisionIdAndRevisionCreatedAndDefaultLangcode
func GetBlockContentFieldRevisionsByRevisionIdAndRevisionCreatedAndDefaultLangcode(offset int, limit int, RevisionId_ int, RevisionCreated_ int, DefaultLangcode_ int) (*[]*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = new([]*BlockContentFieldRevision)
	err := Engine.Table("block_content_field_revision").Where("revision_id = ? and revision_created = ? and default_langcode = ?", RevisionId_, RevisionCreated_, DefaultLangcode_).Limit(limit, offset).Find(_BlockContentFieldRevision)
	return _BlockContentFieldRevision, err
}

// GetBlockContentFieldRevisionsByRevisionIdAndRevisionUserAndRevisionTranslationAffected Get BlockContentFieldRevisions via RevisionIdAndRevisionUserAndRevisionTranslationAffected
func GetBlockContentFieldRevisionsByRevisionIdAndRevisionUserAndRevisionTranslationAffected(offset int, limit int, RevisionId_ int, RevisionUser_ int, RevisionTranslationAffected_ int) (*[]*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = new([]*BlockContentFieldRevision)
	err := Engine.Table("block_content_field_revision").Where("revision_id = ? and revision_user = ? and revision_translation_affected = ?", RevisionId_, RevisionUser_, RevisionTranslationAffected_).Limit(limit, offset).Find(_BlockContentFieldRevision)
	return _BlockContentFieldRevision, err
}

// GetBlockContentFieldRevisionsByRevisionIdAndRevisionUserAndDefaultLangcode Get BlockContentFieldRevisions via RevisionIdAndRevisionUserAndDefaultLangcode
func GetBlockContentFieldRevisionsByRevisionIdAndRevisionUserAndDefaultLangcode(offset int, limit int, RevisionId_ int, RevisionUser_ int, DefaultLangcode_ int) (*[]*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = new([]*BlockContentFieldRevision)
	err := Engine.Table("block_content_field_revision").Where("revision_id = ? and revision_user = ? and default_langcode = ?", RevisionId_, RevisionUser_, DefaultLangcode_).Limit(limit, offset).Find(_BlockContentFieldRevision)
	return _BlockContentFieldRevision, err
}

// GetBlockContentFieldRevisionsByRevisionIdAndRevisionTranslationAffectedAndDefaultLangcode Get BlockContentFieldRevisions via RevisionIdAndRevisionTranslationAffectedAndDefaultLangcode
func GetBlockContentFieldRevisionsByRevisionIdAndRevisionTranslationAffectedAndDefaultLangcode(offset int, limit int, RevisionId_ int, RevisionTranslationAffected_ int, DefaultLangcode_ int) (*[]*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = new([]*BlockContentFieldRevision)
	err := Engine.Table("block_content_field_revision").Where("revision_id = ? and revision_translation_affected = ? and default_langcode = ?", RevisionId_, RevisionTranslationAffected_, DefaultLangcode_).Limit(limit, offset).Find(_BlockContentFieldRevision)
	return _BlockContentFieldRevision, err
}

// GetBlockContentFieldRevisionsByLangcodeAndInfoAndChanged Get BlockContentFieldRevisions via LangcodeAndInfoAndChanged
func GetBlockContentFieldRevisionsByLangcodeAndInfoAndChanged(offset int, limit int, Langcode_ string, Info_ string, Changed_ int) (*[]*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = new([]*BlockContentFieldRevision)
	err := Engine.Table("block_content_field_revision").Where("langcode = ? and info = ? and changed = ?", Langcode_, Info_, Changed_).Limit(limit, offset).Find(_BlockContentFieldRevision)
	return _BlockContentFieldRevision, err
}

// GetBlockContentFieldRevisionsByLangcodeAndInfoAndRevisionCreated Get BlockContentFieldRevisions via LangcodeAndInfoAndRevisionCreated
func GetBlockContentFieldRevisionsByLangcodeAndInfoAndRevisionCreated(offset int, limit int, Langcode_ string, Info_ string, RevisionCreated_ int) (*[]*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = new([]*BlockContentFieldRevision)
	err := Engine.Table("block_content_field_revision").Where("langcode = ? and info = ? and revision_created = ?", Langcode_, Info_, RevisionCreated_).Limit(limit, offset).Find(_BlockContentFieldRevision)
	return _BlockContentFieldRevision, err
}

// GetBlockContentFieldRevisionsByLangcodeAndInfoAndRevisionUser Get BlockContentFieldRevisions via LangcodeAndInfoAndRevisionUser
func GetBlockContentFieldRevisionsByLangcodeAndInfoAndRevisionUser(offset int, limit int, Langcode_ string, Info_ string, RevisionUser_ int) (*[]*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = new([]*BlockContentFieldRevision)
	err := Engine.Table("block_content_field_revision").Where("langcode = ? and info = ? and revision_user = ?", Langcode_, Info_, RevisionUser_).Limit(limit, offset).Find(_BlockContentFieldRevision)
	return _BlockContentFieldRevision, err
}

// GetBlockContentFieldRevisionsByLangcodeAndInfoAndRevisionTranslationAffected Get BlockContentFieldRevisions via LangcodeAndInfoAndRevisionTranslationAffected
func GetBlockContentFieldRevisionsByLangcodeAndInfoAndRevisionTranslationAffected(offset int, limit int, Langcode_ string, Info_ string, RevisionTranslationAffected_ int) (*[]*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = new([]*BlockContentFieldRevision)
	err := Engine.Table("block_content_field_revision").Where("langcode = ? and info = ? and revision_translation_affected = ?", Langcode_, Info_, RevisionTranslationAffected_).Limit(limit, offset).Find(_BlockContentFieldRevision)
	return _BlockContentFieldRevision, err
}

// GetBlockContentFieldRevisionsByLangcodeAndInfoAndDefaultLangcode Get BlockContentFieldRevisions via LangcodeAndInfoAndDefaultLangcode
func GetBlockContentFieldRevisionsByLangcodeAndInfoAndDefaultLangcode(offset int, limit int, Langcode_ string, Info_ string, DefaultLangcode_ int) (*[]*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = new([]*BlockContentFieldRevision)
	err := Engine.Table("block_content_field_revision").Where("langcode = ? and info = ? and default_langcode = ?", Langcode_, Info_, DefaultLangcode_).Limit(limit, offset).Find(_BlockContentFieldRevision)
	return _BlockContentFieldRevision, err
}

// GetBlockContentFieldRevisionsByLangcodeAndChangedAndRevisionCreated Get BlockContentFieldRevisions via LangcodeAndChangedAndRevisionCreated
func GetBlockContentFieldRevisionsByLangcodeAndChangedAndRevisionCreated(offset int, limit int, Langcode_ string, Changed_ int, RevisionCreated_ int) (*[]*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = new([]*BlockContentFieldRevision)
	err := Engine.Table("block_content_field_revision").Where("langcode = ? and changed = ? and revision_created = ?", Langcode_, Changed_, RevisionCreated_).Limit(limit, offset).Find(_BlockContentFieldRevision)
	return _BlockContentFieldRevision, err
}

// GetBlockContentFieldRevisionsByLangcodeAndChangedAndRevisionUser Get BlockContentFieldRevisions via LangcodeAndChangedAndRevisionUser
func GetBlockContentFieldRevisionsByLangcodeAndChangedAndRevisionUser(offset int, limit int, Langcode_ string, Changed_ int, RevisionUser_ int) (*[]*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = new([]*BlockContentFieldRevision)
	err := Engine.Table("block_content_field_revision").Where("langcode = ? and changed = ? and revision_user = ?", Langcode_, Changed_, RevisionUser_).Limit(limit, offset).Find(_BlockContentFieldRevision)
	return _BlockContentFieldRevision, err
}

// GetBlockContentFieldRevisionsByLangcodeAndChangedAndRevisionTranslationAffected Get BlockContentFieldRevisions via LangcodeAndChangedAndRevisionTranslationAffected
func GetBlockContentFieldRevisionsByLangcodeAndChangedAndRevisionTranslationAffected(offset int, limit int, Langcode_ string, Changed_ int, RevisionTranslationAffected_ int) (*[]*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = new([]*BlockContentFieldRevision)
	err := Engine.Table("block_content_field_revision").Where("langcode = ? and changed = ? and revision_translation_affected = ?", Langcode_, Changed_, RevisionTranslationAffected_).Limit(limit, offset).Find(_BlockContentFieldRevision)
	return _BlockContentFieldRevision, err
}

// GetBlockContentFieldRevisionsByLangcodeAndChangedAndDefaultLangcode Get BlockContentFieldRevisions via LangcodeAndChangedAndDefaultLangcode
func GetBlockContentFieldRevisionsByLangcodeAndChangedAndDefaultLangcode(offset int, limit int, Langcode_ string, Changed_ int, DefaultLangcode_ int) (*[]*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = new([]*BlockContentFieldRevision)
	err := Engine.Table("block_content_field_revision").Where("langcode = ? and changed = ? and default_langcode = ?", Langcode_, Changed_, DefaultLangcode_).Limit(limit, offset).Find(_BlockContentFieldRevision)
	return _BlockContentFieldRevision, err
}

// GetBlockContentFieldRevisionsByLangcodeAndRevisionCreatedAndRevisionUser Get BlockContentFieldRevisions via LangcodeAndRevisionCreatedAndRevisionUser
func GetBlockContentFieldRevisionsByLangcodeAndRevisionCreatedAndRevisionUser(offset int, limit int, Langcode_ string, RevisionCreated_ int, RevisionUser_ int) (*[]*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = new([]*BlockContentFieldRevision)
	err := Engine.Table("block_content_field_revision").Where("langcode = ? and revision_created = ? and revision_user = ?", Langcode_, RevisionCreated_, RevisionUser_).Limit(limit, offset).Find(_BlockContentFieldRevision)
	return _BlockContentFieldRevision, err
}

// GetBlockContentFieldRevisionsByLangcodeAndRevisionCreatedAndRevisionTranslationAffected Get BlockContentFieldRevisions via LangcodeAndRevisionCreatedAndRevisionTranslationAffected
func GetBlockContentFieldRevisionsByLangcodeAndRevisionCreatedAndRevisionTranslationAffected(offset int, limit int, Langcode_ string, RevisionCreated_ int, RevisionTranslationAffected_ int) (*[]*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = new([]*BlockContentFieldRevision)
	err := Engine.Table("block_content_field_revision").Where("langcode = ? and revision_created = ? and revision_translation_affected = ?", Langcode_, RevisionCreated_, RevisionTranslationAffected_).Limit(limit, offset).Find(_BlockContentFieldRevision)
	return _BlockContentFieldRevision, err
}

// GetBlockContentFieldRevisionsByLangcodeAndRevisionCreatedAndDefaultLangcode Get BlockContentFieldRevisions via LangcodeAndRevisionCreatedAndDefaultLangcode
func GetBlockContentFieldRevisionsByLangcodeAndRevisionCreatedAndDefaultLangcode(offset int, limit int, Langcode_ string, RevisionCreated_ int, DefaultLangcode_ int) (*[]*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = new([]*BlockContentFieldRevision)
	err := Engine.Table("block_content_field_revision").Where("langcode = ? and revision_created = ? and default_langcode = ?", Langcode_, RevisionCreated_, DefaultLangcode_).Limit(limit, offset).Find(_BlockContentFieldRevision)
	return _BlockContentFieldRevision, err
}

// GetBlockContentFieldRevisionsByLangcodeAndRevisionUserAndRevisionTranslationAffected Get BlockContentFieldRevisions via LangcodeAndRevisionUserAndRevisionTranslationAffected
func GetBlockContentFieldRevisionsByLangcodeAndRevisionUserAndRevisionTranslationAffected(offset int, limit int, Langcode_ string, RevisionUser_ int, RevisionTranslationAffected_ int) (*[]*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = new([]*BlockContentFieldRevision)
	err := Engine.Table("block_content_field_revision").Where("langcode = ? and revision_user = ? and revision_translation_affected = ?", Langcode_, RevisionUser_, RevisionTranslationAffected_).Limit(limit, offset).Find(_BlockContentFieldRevision)
	return _BlockContentFieldRevision, err
}

// GetBlockContentFieldRevisionsByLangcodeAndRevisionUserAndDefaultLangcode Get BlockContentFieldRevisions via LangcodeAndRevisionUserAndDefaultLangcode
func GetBlockContentFieldRevisionsByLangcodeAndRevisionUserAndDefaultLangcode(offset int, limit int, Langcode_ string, RevisionUser_ int, DefaultLangcode_ int) (*[]*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = new([]*BlockContentFieldRevision)
	err := Engine.Table("block_content_field_revision").Where("langcode = ? and revision_user = ? and default_langcode = ?", Langcode_, RevisionUser_, DefaultLangcode_).Limit(limit, offset).Find(_BlockContentFieldRevision)
	return _BlockContentFieldRevision, err
}

// GetBlockContentFieldRevisionsByLangcodeAndRevisionTranslationAffectedAndDefaultLangcode Get BlockContentFieldRevisions via LangcodeAndRevisionTranslationAffectedAndDefaultLangcode
func GetBlockContentFieldRevisionsByLangcodeAndRevisionTranslationAffectedAndDefaultLangcode(offset int, limit int, Langcode_ string, RevisionTranslationAffected_ int, DefaultLangcode_ int) (*[]*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = new([]*BlockContentFieldRevision)
	err := Engine.Table("block_content_field_revision").Where("langcode = ? and revision_translation_affected = ? and default_langcode = ?", Langcode_, RevisionTranslationAffected_, DefaultLangcode_).Limit(limit, offset).Find(_BlockContentFieldRevision)
	return _BlockContentFieldRevision, err
}

// GetBlockContentFieldRevisionsByInfoAndChangedAndRevisionCreated Get BlockContentFieldRevisions via InfoAndChangedAndRevisionCreated
func GetBlockContentFieldRevisionsByInfoAndChangedAndRevisionCreated(offset int, limit int, Info_ string, Changed_ int, RevisionCreated_ int) (*[]*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = new([]*BlockContentFieldRevision)
	err := Engine.Table("block_content_field_revision").Where("info = ? and changed = ? and revision_created = ?", Info_, Changed_, RevisionCreated_).Limit(limit, offset).Find(_BlockContentFieldRevision)
	return _BlockContentFieldRevision, err
}

// GetBlockContentFieldRevisionsByInfoAndChangedAndRevisionUser Get BlockContentFieldRevisions via InfoAndChangedAndRevisionUser
func GetBlockContentFieldRevisionsByInfoAndChangedAndRevisionUser(offset int, limit int, Info_ string, Changed_ int, RevisionUser_ int) (*[]*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = new([]*BlockContentFieldRevision)
	err := Engine.Table("block_content_field_revision").Where("info = ? and changed = ? and revision_user = ?", Info_, Changed_, RevisionUser_).Limit(limit, offset).Find(_BlockContentFieldRevision)
	return _BlockContentFieldRevision, err
}

// GetBlockContentFieldRevisionsByInfoAndChangedAndRevisionTranslationAffected Get BlockContentFieldRevisions via InfoAndChangedAndRevisionTranslationAffected
func GetBlockContentFieldRevisionsByInfoAndChangedAndRevisionTranslationAffected(offset int, limit int, Info_ string, Changed_ int, RevisionTranslationAffected_ int) (*[]*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = new([]*BlockContentFieldRevision)
	err := Engine.Table("block_content_field_revision").Where("info = ? and changed = ? and revision_translation_affected = ?", Info_, Changed_, RevisionTranslationAffected_).Limit(limit, offset).Find(_BlockContentFieldRevision)
	return _BlockContentFieldRevision, err
}

// GetBlockContentFieldRevisionsByInfoAndChangedAndDefaultLangcode Get BlockContentFieldRevisions via InfoAndChangedAndDefaultLangcode
func GetBlockContentFieldRevisionsByInfoAndChangedAndDefaultLangcode(offset int, limit int, Info_ string, Changed_ int, DefaultLangcode_ int) (*[]*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = new([]*BlockContentFieldRevision)
	err := Engine.Table("block_content_field_revision").Where("info = ? and changed = ? and default_langcode = ?", Info_, Changed_, DefaultLangcode_).Limit(limit, offset).Find(_BlockContentFieldRevision)
	return _BlockContentFieldRevision, err
}

// GetBlockContentFieldRevisionsByInfoAndRevisionCreatedAndRevisionUser Get BlockContentFieldRevisions via InfoAndRevisionCreatedAndRevisionUser
func GetBlockContentFieldRevisionsByInfoAndRevisionCreatedAndRevisionUser(offset int, limit int, Info_ string, RevisionCreated_ int, RevisionUser_ int) (*[]*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = new([]*BlockContentFieldRevision)
	err := Engine.Table("block_content_field_revision").Where("info = ? and revision_created = ? and revision_user = ?", Info_, RevisionCreated_, RevisionUser_).Limit(limit, offset).Find(_BlockContentFieldRevision)
	return _BlockContentFieldRevision, err
}

// GetBlockContentFieldRevisionsByInfoAndRevisionCreatedAndRevisionTranslationAffected Get BlockContentFieldRevisions via InfoAndRevisionCreatedAndRevisionTranslationAffected
func GetBlockContentFieldRevisionsByInfoAndRevisionCreatedAndRevisionTranslationAffected(offset int, limit int, Info_ string, RevisionCreated_ int, RevisionTranslationAffected_ int) (*[]*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = new([]*BlockContentFieldRevision)
	err := Engine.Table("block_content_field_revision").Where("info = ? and revision_created = ? and revision_translation_affected = ?", Info_, RevisionCreated_, RevisionTranslationAffected_).Limit(limit, offset).Find(_BlockContentFieldRevision)
	return _BlockContentFieldRevision, err
}

// GetBlockContentFieldRevisionsByInfoAndRevisionCreatedAndDefaultLangcode Get BlockContentFieldRevisions via InfoAndRevisionCreatedAndDefaultLangcode
func GetBlockContentFieldRevisionsByInfoAndRevisionCreatedAndDefaultLangcode(offset int, limit int, Info_ string, RevisionCreated_ int, DefaultLangcode_ int) (*[]*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = new([]*BlockContentFieldRevision)
	err := Engine.Table("block_content_field_revision").Where("info = ? and revision_created = ? and default_langcode = ?", Info_, RevisionCreated_, DefaultLangcode_).Limit(limit, offset).Find(_BlockContentFieldRevision)
	return _BlockContentFieldRevision, err
}

// GetBlockContentFieldRevisionsByInfoAndRevisionUserAndRevisionTranslationAffected Get BlockContentFieldRevisions via InfoAndRevisionUserAndRevisionTranslationAffected
func GetBlockContentFieldRevisionsByInfoAndRevisionUserAndRevisionTranslationAffected(offset int, limit int, Info_ string, RevisionUser_ int, RevisionTranslationAffected_ int) (*[]*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = new([]*BlockContentFieldRevision)
	err := Engine.Table("block_content_field_revision").Where("info = ? and revision_user = ? and revision_translation_affected = ?", Info_, RevisionUser_, RevisionTranslationAffected_).Limit(limit, offset).Find(_BlockContentFieldRevision)
	return _BlockContentFieldRevision, err
}

// GetBlockContentFieldRevisionsByInfoAndRevisionUserAndDefaultLangcode Get BlockContentFieldRevisions via InfoAndRevisionUserAndDefaultLangcode
func GetBlockContentFieldRevisionsByInfoAndRevisionUserAndDefaultLangcode(offset int, limit int, Info_ string, RevisionUser_ int, DefaultLangcode_ int) (*[]*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = new([]*BlockContentFieldRevision)
	err := Engine.Table("block_content_field_revision").Where("info = ? and revision_user = ? and default_langcode = ?", Info_, RevisionUser_, DefaultLangcode_).Limit(limit, offset).Find(_BlockContentFieldRevision)
	return _BlockContentFieldRevision, err
}

// GetBlockContentFieldRevisionsByInfoAndRevisionTranslationAffectedAndDefaultLangcode Get BlockContentFieldRevisions via InfoAndRevisionTranslationAffectedAndDefaultLangcode
func GetBlockContentFieldRevisionsByInfoAndRevisionTranslationAffectedAndDefaultLangcode(offset int, limit int, Info_ string, RevisionTranslationAffected_ int, DefaultLangcode_ int) (*[]*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = new([]*BlockContentFieldRevision)
	err := Engine.Table("block_content_field_revision").Where("info = ? and revision_translation_affected = ? and default_langcode = ?", Info_, RevisionTranslationAffected_, DefaultLangcode_).Limit(limit, offset).Find(_BlockContentFieldRevision)
	return _BlockContentFieldRevision, err
}

// GetBlockContentFieldRevisionsByChangedAndRevisionCreatedAndRevisionUser Get BlockContentFieldRevisions via ChangedAndRevisionCreatedAndRevisionUser
func GetBlockContentFieldRevisionsByChangedAndRevisionCreatedAndRevisionUser(offset int, limit int, Changed_ int, RevisionCreated_ int, RevisionUser_ int) (*[]*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = new([]*BlockContentFieldRevision)
	err := Engine.Table("block_content_field_revision").Where("changed = ? and revision_created = ? and revision_user = ?", Changed_, RevisionCreated_, RevisionUser_).Limit(limit, offset).Find(_BlockContentFieldRevision)
	return _BlockContentFieldRevision, err
}

// GetBlockContentFieldRevisionsByChangedAndRevisionCreatedAndRevisionTranslationAffected Get BlockContentFieldRevisions via ChangedAndRevisionCreatedAndRevisionTranslationAffected
func GetBlockContentFieldRevisionsByChangedAndRevisionCreatedAndRevisionTranslationAffected(offset int, limit int, Changed_ int, RevisionCreated_ int, RevisionTranslationAffected_ int) (*[]*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = new([]*BlockContentFieldRevision)
	err := Engine.Table("block_content_field_revision").Where("changed = ? and revision_created = ? and revision_translation_affected = ?", Changed_, RevisionCreated_, RevisionTranslationAffected_).Limit(limit, offset).Find(_BlockContentFieldRevision)
	return _BlockContentFieldRevision, err
}

// GetBlockContentFieldRevisionsByChangedAndRevisionCreatedAndDefaultLangcode Get BlockContentFieldRevisions via ChangedAndRevisionCreatedAndDefaultLangcode
func GetBlockContentFieldRevisionsByChangedAndRevisionCreatedAndDefaultLangcode(offset int, limit int, Changed_ int, RevisionCreated_ int, DefaultLangcode_ int) (*[]*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = new([]*BlockContentFieldRevision)
	err := Engine.Table("block_content_field_revision").Where("changed = ? and revision_created = ? and default_langcode = ?", Changed_, RevisionCreated_, DefaultLangcode_).Limit(limit, offset).Find(_BlockContentFieldRevision)
	return _BlockContentFieldRevision, err
}

// GetBlockContentFieldRevisionsByChangedAndRevisionUserAndRevisionTranslationAffected Get BlockContentFieldRevisions via ChangedAndRevisionUserAndRevisionTranslationAffected
func GetBlockContentFieldRevisionsByChangedAndRevisionUserAndRevisionTranslationAffected(offset int, limit int, Changed_ int, RevisionUser_ int, RevisionTranslationAffected_ int) (*[]*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = new([]*BlockContentFieldRevision)
	err := Engine.Table("block_content_field_revision").Where("changed = ? and revision_user = ? and revision_translation_affected = ?", Changed_, RevisionUser_, RevisionTranslationAffected_).Limit(limit, offset).Find(_BlockContentFieldRevision)
	return _BlockContentFieldRevision, err
}

// GetBlockContentFieldRevisionsByChangedAndRevisionUserAndDefaultLangcode Get BlockContentFieldRevisions via ChangedAndRevisionUserAndDefaultLangcode
func GetBlockContentFieldRevisionsByChangedAndRevisionUserAndDefaultLangcode(offset int, limit int, Changed_ int, RevisionUser_ int, DefaultLangcode_ int) (*[]*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = new([]*BlockContentFieldRevision)
	err := Engine.Table("block_content_field_revision").Where("changed = ? and revision_user = ? and default_langcode = ?", Changed_, RevisionUser_, DefaultLangcode_).Limit(limit, offset).Find(_BlockContentFieldRevision)
	return _BlockContentFieldRevision, err
}

// GetBlockContentFieldRevisionsByChangedAndRevisionTranslationAffectedAndDefaultLangcode Get BlockContentFieldRevisions via ChangedAndRevisionTranslationAffectedAndDefaultLangcode
func GetBlockContentFieldRevisionsByChangedAndRevisionTranslationAffectedAndDefaultLangcode(offset int, limit int, Changed_ int, RevisionTranslationAffected_ int, DefaultLangcode_ int) (*[]*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = new([]*BlockContentFieldRevision)
	err := Engine.Table("block_content_field_revision").Where("changed = ? and revision_translation_affected = ? and default_langcode = ?", Changed_, RevisionTranslationAffected_, DefaultLangcode_).Limit(limit, offset).Find(_BlockContentFieldRevision)
	return _BlockContentFieldRevision, err
}

// GetBlockContentFieldRevisionsByRevisionCreatedAndRevisionUserAndRevisionTranslationAffected Get BlockContentFieldRevisions via RevisionCreatedAndRevisionUserAndRevisionTranslationAffected
func GetBlockContentFieldRevisionsByRevisionCreatedAndRevisionUserAndRevisionTranslationAffected(offset int, limit int, RevisionCreated_ int, RevisionUser_ int, RevisionTranslationAffected_ int) (*[]*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = new([]*BlockContentFieldRevision)
	err := Engine.Table("block_content_field_revision").Where("revision_created = ? and revision_user = ? and revision_translation_affected = ?", RevisionCreated_, RevisionUser_, RevisionTranslationAffected_).Limit(limit, offset).Find(_BlockContentFieldRevision)
	return _BlockContentFieldRevision, err
}

// GetBlockContentFieldRevisionsByRevisionCreatedAndRevisionUserAndDefaultLangcode Get BlockContentFieldRevisions via RevisionCreatedAndRevisionUserAndDefaultLangcode
func GetBlockContentFieldRevisionsByRevisionCreatedAndRevisionUserAndDefaultLangcode(offset int, limit int, RevisionCreated_ int, RevisionUser_ int, DefaultLangcode_ int) (*[]*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = new([]*BlockContentFieldRevision)
	err := Engine.Table("block_content_field_revision").Where("revision_created = ? and revision_user = ? and default_langcode = ?", RevisionCreated_, RevisionUser_, DefaultLangcode_).Limit(limit, offset).Find(_BlockContentFieldRevision)
	return _BlockContentFieldRevision, err
}

// GetBlockContentFieldRevisionsByRevisionCreatedAndRevisionTranslationAffectedAndDefaultLangcode Get BlockContentFieldRevisions via RevisionCreatedAndRevisionTranslationAffectedAndDefaultLangcode
func GetBlockContentFieldRevisionsByRevisionCreatedAndRevisionTranslationAffectedAndDefaultLangcode(offset int, limit int, RevisionCreated_ int, RevisionTranslationAffected_ int, DefaultLangcode_ int) (*[]*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = new([]*BlockContentFieldRevision)
	err := Engine.Table("block_content_field_revision").Where("revision_created = ? and revision_translation_affected = ? and default_langcode = ?", RevisionCreated_, RevisionTranslationAffected_, DefaultLangcode_).Limit(limit, offset).Find(_BlockContentFieldRevision)
	return _BlockContentFieldRevision, err
}

// GetBlockContentFieldRevisionsByRevisionUserAndRevisionTranslationAffectedAndDefaultLangcode Get BlockContentFieldRevisions via RevisionUserAndRevisionTranslationAffectedAndDefaultLangcode
func GetBlockContentFieldRevisionsByRevisionUserAndRevisionTranslationAffectedAndDefaultLangcode(offset int, limit int, RevisionUser_ int, RevisionTranslationAffected_ int, DefaultLangcode_ int) (*[]*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = new([]*BlockContentFieldRevision)
	err := Engine.Table("block_content_field_revision").Where("revision_user = ? and revision_translation_affected = ? and default_langcode = ?", RevisionUser_, RevisionTranslationAffected_, DefaultLangcode_).Limit(limit, offset).Find(_BlockContentFieldRevision)
	return _BlockContentFieldRevision, err
}

// GetBlockContentFieldRevisionsByIdAndRevisionId Get BlockContentFieldRevisions via IdAndRevisionId
func GetBlockContentFieldRevisionsByIdAndRevisionId(offset int, limit int, Id_ int64, RevisionId_ int) (*[]*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = new([]*BlockContentFieldRevision)
	err := Engine.Table("block_content_field_revision").Where("id = ? and revision_id = ?", Id_, RevisionId_).Limit(limit, offset).Find(_BlockContentFieldRevision)
	return _BlockContentFieldRevision, err
}

// GetBlockContentFieldRevisionsByIdAndLangcode Get BlockContentFieldRevisions via IdAndLangcode
func GetBlockContentFieldRevisionsByIdAndLangcode(offset int, limit int, Id_ int64, Langcode_ string) (*[]*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = new([]*BlockContentFieldRevision)
	err := Engine.Table("block_content_field_revision").Where("id = ? and langcode = ?", Id_, Langcode_).Limit(limit, offset).Find(_BlockContentFieldRevision)
	return _BlockContentFieldRevision, err
}

// GetBlockContentFieldRevisionsByIdAndInfo Get BlockContentFieldRevisions via IdAndInfo
func GetBlockContentFieldRevisionsByIdAndInfo(offset int, limit int, Id_ int64, Info_ string) (*[]*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = new([]*BlockContentFieldRevision)
	err := Engine.Table("block_content_field_revision").Where("id = ? and info = ?", Id_, Info_).Limit(limit, offset).Find(_BlockContentFieldRevision)
	return _BlockContentFieldRevision, err
}

// GetBlockContentFieldRevisionsByIdAndChanged Get BlockContentFieldRevisions via IdAndChanged
func GetBlockContentFieldRevisionsByIdAndChanged(offset int, limit int, Id_ int64, Changed_ int) (*[]*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = new([]*BlockContentFieldRevision)
	err := Engine.Table("block_content_field_revision").Where("id = ? and changed = ?", Id_, Changed_).Limit(limit, offset).Find(_BlockContentFieldRevision)
	return _BlockContentFieldRevision, err
}

// GetBlockContentFieldRevisionsByIdAndRevisionCreated Get BlockContentFieldRevisions via IdAndRevisionCreated
func GetBlockContentFieldRevisionsByIdAndRevisionCreated(offset int, limit int, Id_ int64, RevisionCreated_ int) (*[]*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = new([]*BlockContentFieldRevision)
	err := Engine.Table("block_content_field_revision").Where("id = ? and revision_created = ?", Id_, RevisionCreated_).Limit(limit, offset).Find(_BlockContentFieldRevision)
	return _BlockContentFieldRevision, err
}

// GetBlockContentFieldRevisionsByIdAndRevisionUser Get BlockContentFieldRevisions via IdAndRevisionUser
func GetBlockContentFieldRevisionsByIdAndRevisionUser(offset int, limit int, Id_ int64, RevisionUser_ int) (*[]*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = new([]*BlockContentFieldRevision)
	err := Engine.Table("block_content_field_revision").Where("id = ? and revision_user = ?", Id_, RevisionUser_).Limit(limit, offset).Find(_BlockContentFieldRevision)
	return _BlockContentFieldRevision, err
}

// GetBlockContentFieldRevisionsByIdAndRevisionTranslationAffected Get BlockContentFieldRevisions via IdAndRevisionTranslationAffected
func GetBlockContentFieldRevisionsByIdAndRevisionTranslationAffected(offset int, limit int, Id_ int64, RevisionTranslationAffected_ int) (*[]*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = new([]*BlockContentFieldRevision)
	err := Engine.Table("block_content_field_revision").Where("id = ? and revision_translation_affected = ?", Id_, RevisionTranslationAffected_).Limit(limit, offset).Find(_BlockContentFieldRevision)
	return _BlockContentFieldRevision, err
}

// GetBlockContentFieldRevisionsByIdAndDefaultLangcode Get BlockContentFieldRevisions via IdAndDefaultLangcode
func GetBlockContentFieldRevisionsByIdAndDefaultLangcode(offset int, limit int, Id_ int64, DefaultLangcode_ int) (*[]*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = new([]*BlockContentFieldRevision)
	err := Engine.Table("block_content_field_revision").Where("id = ? and default_langcode = ?", Id_, DefaultLangcode_).Limit(limit, offset).Find(_BlockContentFieldRevision)
	return _BlockContentFieldRevision, err
}

// GetBlockContentFieldRevisionsByRevisionIdAndLangcode Get BlockContentFieldRevisions via RevisionIdAndLangcode
func GetBlockContentFieldRevisionsByRevisionIdAndLangcode(offset int, limit int, RevisionId_ int, Langcode_ string) (*[]*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = new([]*BlockContentFieldRevision)
	err := Engine.Table("block_content_field_revision").Where("revision_id = ? and langcode = ?", RevisionId_, Langcode_).Limit(limit, offset).Find(_BlockContentFieldRevision)
	return _BlockContentFieldRevision, err
}

// GetBlockContentFieldRevisionsByRevisionIdAndInfo Get BlockContentFieldRevisions via RevisionIdAndInfo
func GetBlockContentFieldRevisionsByRevisionIdAndInfo(offset int, limit int, RevisionId_ int, Info_ string) (*[]*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = new([]*BlockContentFieldRevision)
	err := Engine.Table("block_content_field_revision").Where("revision_id = ? and info = ?", RevisionId_, Info_).Limit(limit, offset).Find(_BlockContentFieldRevision)
	return _BlockContentFieldRevision, err
}

// GetBlockContentFieldRevisionsByRevisionIdAndChanged Get BlockContentFieldRevisions via RevisionIdAndChanged
func GetBlockContentFieldRevisionsByRevisionIdAndChanged(offset int, limit int, RevisionId_ int, Changed_ int) (*[]*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = new([]*BlockContentFieldRevision)
	err := Engine.Table("block_content_field_revision").Where("revision_id = ? and changed = ?", RevisionId_, Changed_).Limit(limit, offset).Find(_BlockContentFieldRevision)
	return _BlockContentFieldRevision, err
}

// GetBlockContentFieldRevisionsByRevisionIdAndRevisionCreated Get BlockContentFieldRevisions via RevisionIdAndRevisionCreated
func GetBlockContentFieldRevisionsByRevisionIdAndRevisionCreated(offset int, limit int, RevisionId_ int, RevisionCreated_ int) (*[]*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = new([]*BlockContentFieldRevision)
	err := Engine.Table("block_content_field_revision").Where("revision_id = ? and revision_created = ?", RevisionId_, RevisionCreated_).Limit(limit, offset).Find(_BlockContentFieldRevision)
	return _BlockContentFieldRevision, err
}

// GetBlockContentFieldRevisionsByRevisionIdAndRevisionUser Get BlockContentFieldRevisions via RevisionIdAndRevisionUser
func GetBlockContentFieldRevisionsByRevisionIdAndRevisionUser(offset int, limit int, RevisionId_ int, RevisionUser_ int) (*[]*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = new([]*BlockContentFieldRevision)
	err := Engine.Table("block_content_field_revision").Where("revision_id = ? and revision_user = ?", RevisionId_, RevisionUser_).Limit(limit, offset).Find(_BlockContentFieldRevision)
	return _BlockContentFieldRevision, err
}

// GetBlockContentFieldRevisionsByRevisionIdAndRevisionTranslationAffected Get BlockContentFieldRevisions via RevisionIdAndRevisionTranslationAffected
func GetBlockContentFieldRevisionsByRevisionIdAndRevisionTranslationAffected(offset int, limit int, RevisionId_ int, RevisionTranslationAffected_ int) (*[]*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = new([]*BlockContentFieldRevision)
	err := Engine.Table("block_content_field_revision").Where("revision_id = ? and revision_translation_affected = ?", RevisionId_, RevisionTranslationAffected_).Limit(limit, offset).Find(_BlockContentFieldRevision)
	return _BlockContentFieldRevision, err
}

// GetBlockContentFieldRevisionsByRevisionIdAndDefaultLangcode Get BlockContentFieldRevisions via RevisionIdAndDefaultLangcode
func GetBlockContentFieldRevisionsByRevisionIdAndDefaultLangcode(offset int, limit int, RevisionId_ int, DefaultLangcode_ int) (*[]*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = new([]*BlockContentFieldRevision)
	err := Engine.Table("block_content_field_revision").Where("revision_id = ? and default_langcode = ?", RevisionId_, DefaultLangcode_).Limit(limit, offset).Find(_BlockContentFieldRevision)
	return _BlockContentFieldRevision, err
}

// GetBlockContentFieldRevisionsByLangcodeAndInfo Get BlockContentFieldRevisions via LangcodeAndInfo
func GetBlockContentFieldRevisionsByLangcodeAndInfo(offset int, limit int, Langcode_ string, Info_ string) (*[]*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = new([]*BlockContentFieldRevision)
	err := Engine.Table("block_content_field_revision").Where("langcode = ? and info = ?", Langcode_, Info_).Limit(limit, offset).Find(_BlockContentFieldRevision)
	return _BlockContentFieldRevision, err
}

// GetBlockContentFieldRevisionsByLangcodeAndChanged Get BlockContentFieldRevisions via LangcodeAndChanged
func GetBlockContentFieldRevisionsByLangcodeAndChanged(offset int, limit int, Langcode_ string, Changed_ int) (*[]*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = new([]*BlockContentFieldRevision)
	err := Engine.Table("block_content_field_revision").Where("langcode = ? and changed = ?", Langcode_, Changed_).Limit(limit, offset).Find(_BlockContentFieldRevision)
	return _BlockContentFieldRevision, err
}

// GetBlockContentFieldRevisionsByLangcodeAndRevisionCreated Get BlockContentFieldRevisions via LangcodeAndRevisionCreated
func GetBlockContentFieldRevisionsByLangcodeAndRevisionCreated(offset int, limit int, Langcode_ string, RevisionCreated_ int) (*[]*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = new([]*BlockContentFieldRevision)
	err := Engine.Table("block_content_field_revision").Where("langcode = ? and revision_created = ?", Langcode_, RevisionCreated_).Limit(limit, offset).Find(_BlockContentFieldRevision)
	return _BlockContentFieldRevision, err
}

// GetBlockContentFieldRevisionsByLangcodeAndRevisionUser Get BlockContentFieldRevisions via LangcodeAndRevisionUser
func GetBlockContentFieldRevisionsByLangcodeAndRevisionUser(offset int, limit int, Langcode_ string, RevisionUser_ int) (*[]*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = new([]*BlockContentFieldRevision)
	err := Engine.Table("block_content_field_revision").Where("langcode = ? and revision_user = ?", Langcode_, RevisionUser_).Limit(limit, offset).Find(_BlockContentFieldRevision)
	return _BlockContentFieldRevision, err
}

// GetBlockContentFieldRevisionsByLangcodeAndRevisionTranslationAffected Get BlockContentFieldRevisions via LangcodeAndRevisionTranslationAffected
func GetBlockContentFieldRevisionsByLangcodeAndRevisionTranslationAffected(offset int, limit int, Langcode_ string, RevisionTranslationAffected_ int) (*[]*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = new([]*BlockContentFieldRevision)
	err := Engine.Table("block_content_field_revision").Where("langcode = ? and revision_translation_affected = ?", Langcode_, RevisionTranslationAffected_).Limit(limit, offset).Find(_BlockContentFieldRevision)
	return _BlockContentFieldRevision, err
}

// GetBlockContentFieldRevisionsByLangcodeAndDefaultLangcode Get BlockContentFieldRevisions via LangcodeAndDefaultLangcode
func GetBlockContentFieldRevisionsByLangcodeAndDefaultLangcode(offset int, limit int, Langcode_ string, DefaultLangcode_ int) (*[]*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = new([]*BlockContentFieldRevision)
	err := Engine.Table("block_content_field_revision").Where("langcode = ? and default_langcode = ?", Langcode_, DefaultLangcode_).Limit(limit, offset).Find(_BlockContentFieldRevision)
	return _BlockContentFieldRevision, err
}

// GetBlockContentFieldRevisionsByInfoAndChanged Get BlockContentFieldRevisions via InfoAndChanged
func GetBlockContentFieldRevisionsByInfoAndChanged(offset int, limit int, Info_ string, Changed_ int) (*[]*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = new([]*BlockContentFieldRevision)
	err := Engine.Table("block_content_field_revision").Where("info = ? and changed = ?", Info_, Changed_).Limit(limit, offset).Find(_BlockContentFieldRevision)
	return _BlockContentFieldRevision, err
}

// GetBlockContentFieldRevisionsByInfoAndRevisionCreated Get BlockContentFieldRevisions via InfoAndRevisionCreated
func GetBlockContentFieldRevisionsByInfoAndRevisionCreated(offset int, limit int, Info_ string, RevisionCreated_ int) (*[]*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = new([]*BlockContentFieldRevision)
	err := Engine.Table("block_content_field_revision").Where("info = ? and revision_created = ?", Info_, RevisionCreated_).Limit(limit, offset).Find(_BlockContentFieldRevision)
	return _BlockContentFieldRevision, err
}

// GetBlockContentFieldRevisionsByInfoAndRevisionUser Get BlockContentFieldRevisions via InfoAndRevisionUser
func GetBlockContentFieldRevisionsByInfoAndRevisionUser(offset int, limit int, Info_ string, RevisionUser_ int) (*[]*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = new([]*BlockContentFieldRevision)
	err := Engine.Table("block_content_field_revision").Where("info = ? and revision_user = ?", Info_, RevisionUser_).Limit(limit, offset).Find(_BlockContentFieldRevision)
	return _BlockContentFieldRevision, err
}

// GetBlockContentFieldRevisionsByInfoAndRevisionTranslationAffected Get BlockContentFieldRevisions via InfoAndRevisionTranslationAffected
func GetBlockContentFieldRevisionsByInfoAndRevisionTranslationAffected(offset int, limit int, Info_ string, RevisionTranslationAffected_ int) (*[]*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = new([]*BlockContentFieldRevision)
	err := Engine.Table("block_content_field_revision").Where("info = ? and revision_translation_affected = ?", Info_, RevisionTranslationAffected_).Limit(limit, offset).Find(_BlockContentFieldRevision)
	return _BlockContentFieldRevision, err
}

// GetBlockContentFieldRevisionsByInfoAndDefaultLangcode Get BlockContentFieldRevisions via InfoAndDefaultLangcode
func GetBlockContentFieldRevisionsByInfoAndDefaultLangcode(offset int, limit int, Info_ string, DefaultLangcode_ int) (*[]*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = new([]*BlockContentFieldRevision)
	err := Engine.Table("block_content_field_revision").Where("info = ? and default_langcode = ?", Info_, DefaultLangcode_).Limit(limit, offset).Find(_BlockContentFieldRevision)
	return _BlockContentFieldRevision, err
}

// GetBlockContentFieldRevisionsByChangedAndRevisionCreated Get BlockContentFieldRevisions via ChangedAndRevisionCreated
func GetBlockContentFieldRevisionsByChangedAndRevisionCreated(offset int, limit int, Changed_ int, RevisionCreated_ int) (*[]*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = new([]*BlockContentFieldRevision)
	err := Engine.Table("block_content_field_revision").Where("changed = ? and revision_created = ?", Changed_, RevisionCreated_).Limit(limit, offset).Find(_BlockContentFieldRevision)
	return _BlockContentFieldRevision, err
}

// GetBlockContentFieldRevisionsByChangedAndRevisionUser Get BlockContentFieldRevisions via ChangedAndRevisionUser
func GetBlockContentFieldRevisionsByChangedAndRevisionUser(offset int, limit int, Changed_ int, RevisionUser_ int) (*[]*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = new([]*BlockContentFieldRevision)
	err := Engine.Table("block_content_field_revision").Where("changed = ? and revision_user = ?", Changed_, RevisionUser_).Limit(limit, offset).Find(_BlockContentFieldRevision)
	return _BlockContentFieldRevision, err
}

// GetBlockContentFieldRevisionsByChangedAndRevisionTranslationAffected Get BlockContentFieldRevisions via ChangedAndRevisionTranslationAffected
func GetBlockContentFieldRevisionsByChangedAndRevisionTranslationAffected(offset int, limit int, Changed_ int, RevisionTranslationAffected_ int) (*[]*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = new([]*BlockContentFieldRevision)
	err := Engine.Table("block_content_field_revision").Where("changed = ? and revision_translation_affected = ?", Changed_, RevisionTranslationAffected_).Limit(limit, offset).Find(_BlockContentFieldRevision)
	return _BlockContentFieldRevision, err
}

// GetBlockContentFieldRevisionsByChangedAndDefaultLangcode Get BlockContentFieldRevisions via ChangedAndDefaultLangcode
func GetBlockContentFieldRevisionsByChangedAndDefaultLangcode(offset int, limit int, Changed_ int, DefaultLangcode_ int) (*[]*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = new([]*BlockContentFieldRevision)
	err := Engine.Table("block_content_field_revision").Where("changed = ? and default_langcode = ?", Changed_, DefaultLangcode_).Limit(limit, offset).Find(_BlockContentFieldRevision)
	return _BlockContentFieldRevision, err
}

// GetBlockContentFieldRevisionsByRevisionCreatedAndRevisionUser Get BlockContentFieldRevisions via RevisionCreatedAndRevisionUser
func GetBlockContentFieldRevisionsByRevisionCreatedAndRevisionUser(offset int, limit int, RevisionCreated_ int, RevisionUser_ int) (*[]*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = new([]*BlockContentFieldRevision)
	err := Engine.Table("block_content_field_revision").Where("revision_created = ? and revision_user = ?", RevisionCreated_, RevisionUser_).Limit(limit, offset).Find(_BlockContentFieldRevision)
	return _BlockContentFieldRevision, err
}

// GetBlockContentFieldRevisionsByRevisionCreatedAndRevisionTranslationAffected Get BlockContentFieldRevisions via RevisionCreatedAndRevisionTranslationAffected
func GetBlockContentFieldRevisionsByRevisionCreatedAndRevisionTranslationAffected(offset int, limit int, RevisionCreated_ int, RevisionTranslationAffected_ int) (*[]*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = new([]*BlockContentFieldRevision)
	err := Engine.Table("block_content_field_revision").Where("revision_created = ? and revision_translation_affected = ?", RevisionCreated_, RevisionTranslationAffected_).Limit(limit, offset).Find(_BlockContentFieldRevision)
	return _BlockContentFieldRevision, err
}

// GetBlockContentFieldRevisionsByRevisionCreatedAndDefaultLangcode Get BlockContentFieldRevisions via RevisionCreatedAndDefaultLangcode
func GetBlockContentFieldRevisionsByRevisionCreatedAndDefaultLangcode(offset int, limit int, RevisionCreated_ int, DefaultLangcode_ int) (*[]*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = new([]*BlockContentFieldRevision)
	err := Engine.Table("block_content_field_revision").Where("revision_created = ? and default_langcode = ?", RevisionCreated_, DefaultLangcode_).Limit(limit, offset).Find(_BlockContentFieldRevision)
	return _BlockContentFieldRevision, err
}

// GetBlockContentFieldRevisionsByRevisionUserAndRevisionTranslationAffected Get BlockContentFieldRevisions via RevisionUserAndRevisionTranslationAffected
func GetBlockContentFieldRevisionsByRevisionUserAndRevisionTranslationAffected(offset int, limit int, RevisionUser_ int, RevisionTranslationAffected_ int) (*[]*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = new([]*BlockContentFieldRevision)
	err := Engine.Table("block_content_field_revision").Where("revision_user = ? and revision_translation_affected = ?", RevisionUser_, RevisionTranslationAffected_).Limit(limit, offset).Find(_BlockContentFieldRevision)
	return _BlockContentFieldRevision, err
}

// GetBlockContentFieldRevisionsByRevisionUserAndDefaultLangcode Get BlockContentFieldRevisions via RevisionUserAndDefaultLangcode
func GetBlockContentFieldRevisionsByRevisionUserAndDefaultLangcode(offset int, limit int, RevisionUser_ int, DefaultLangcode_ int) (*[]*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = new([]*BlockContentFieldRevision)
	err := Engine.Table("block_content_field_revision").Where("revision_user = ? and default_langcode = ?", RevisionUser_, DefaultLangcode_).Limit(limit, offset).Find(_BlockContentFieldRevision)
	return _BlockContentFieldRevision, err
}

// GetBlockContentFieldRevisionsByRevisionTranslationAffectedAndDefaultLangcode Get BlockContentFieldRevisions via RevisionTranslationAffectedAndDefaultLangcode
func GetBlockContentFieldRevisionsByRevisionTranslationAffectedAndDefaultLangcode(offset int, limit int, RevisionTranslationAffected_ int, DefaultLangcode_ int) (*[]*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = new([]*BlockContentFieldRevision)
	err := Engine.Table("block_content_field_revision").Where("revision_translation_affected = ? and default_langcode = ?", RevisionTranslationAffected_, DefaultLangcode_).Limit(limit, offset).Find(_BlockContentFieldRevision)
	return _BlockContentFieldRevision, err
}

// GetBlockContentFieldRevisions Get BlockContentFieldRevisions via field
func GetBlockContentFieldRevisions(offset int, limit int, field string) (*[]*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = new([]*BlockContentFieldRevision)
	err := Engine.Table("block_content_field_revision").Limit(limit, offset).Desc(field).Find(_BlockContentFieldRevision)
	return _BlockContentFieldRevision, err
}

// GetBlockContentFieldRevisionsViaId Get BlockContentFieldRevisions via Id
func GetBlockContentFieldRevisionsViaId(offset int, limit int, Id_ int64, field string) (*[]*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = new([]*BlockContentFieldRevision)
	err := Engine.Table("block_content_field_revision").Where("id = ?", Id_).Limit(limit, offset).Desc(field).Find(_BlockContentFieldRevision)
	return _BlockContentFieldRevision, err
}

// GetBlockContentFieldRevisionsViaRevisionId Get BlockContentFieldRevisions via RevisionId
func GetBlockContentFieldRevisionsViaRevisionId(offset int, limit int, RevisionId_ int, field string) (*[]*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = new([]*BlockContentFieldRevision)
	err := Engine.Table("block_content_field_revision").Where("revision_id = ?", RevisionId_).Limit(limit, offset).Desc(field).Find(_BlockContentFieldRevision)
	return _BlockContentFieldRevision, err
}

// GetBlockContentFieldRevisionsViaLangcode Get BlockContentFieldRevisions via Langcode
func GetBlockContentFieldRevisionsViaLangcode(offset int, limit int, Langcode_ string, field string) (*[]*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = new([]*BlockContentFieldRevision)
	err := Engine.Table("block_content_field_revision").Where("langcode = ?", Langcode_).Limit(limit, offset).Desc(field).Find(_BlockContentFieldRevision)
	return _BlockContentFieldRevision, err
}

// GetBlockContentFieldRevisionsViaInfo Get BlockContentFieldRevisions via Info
func GetBlockContentFieldRevisionsViaInfo(offset int, limit int, Info_ string, field string) (*[]*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = new([]*BlockContentFieldRevision)
	err := Engine.Table("block_content_field_revision").Where("info = ?", Info_).Limit(limit, offset).Desc(field).Find(_BlockContentFieldRevision)
	return _BlockContentFieldRevision, err
}

// GetBlockContentFieldRevisionsViaChanged Get BlockContentFieldRevisions via Changed
func GetBlockContentFieldRevisionsViaChanged(offset int, limit int, Changed_ int, field string) (*[]*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = new([]*BlockContentFieldRevision)
	err := Engine.Table("block_content_field_revision").Where("changed = ?", Changed_).Limit(limit, offset).Desc(field).Find(_BlockContentFieldRevision)
	return _BlockContentFieldRevision, err
}

// GetBlockContentFieldRevisionsViaRevisionCreated Get BlockContentFieldRevisions via RevisionCreated
func GetBlockContentFieldRevisionsViaRevisionCreated(offset int, limit int, RevisionCreated_ int, field string) (*[]*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = new([]*BlockContentFieldRevision)
	err := Engine.Table("block_content_field_revision").Where("revision_created = ?", RevisionCreated_).Limit(limit, offset).Desc(field).Find(_BlockContentFieldRevision)
	return _BlockContentFieldRevision, err
}

// GetBlockContentFieldRevisionsViaRevisionUser Get BlockContentFieldRevisions via RevisionUser
func GetBlockContentFieldRevisionsViaRevisionUser(offset int, limit int, RevisionUser_ int, field string) (*[]*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = new([]*BlockContentFieldRevision)
	err := Engine.Table("block_content_field_revision").Where("revision_user = ?", RevisionUser_).Limit(limit, offset).Desc(field).Find(_BlockContentFieldRevision)
	return _BlockContentFieldRevision, err
}

// GetBlockContentFieldRevisionsViaRevisionTranslationAffected Get BlockContentFieldRevisions via RevisionTranslationAffected
func GetBlockContentFieldRevisionsViaRevisionTranslationAffected(offset int, limit int, RevisionTranslationAffected_ int, field string) (*[]*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = new([]*BlockContentFieldRevision)
	err := Engine.Table("block_content_field_revision").Where("revision_translation_affected = ?", RevisionTranslationAffected_).Limit(limit, offset).Desc(field).Find(_BlockContentFieldRevision)
	return _BlockContentFieldRevision, err
}

// GetBlockContentFieldRevisionsViaDefaultLangcode Get BlockContentFieldRevisions via DefaultLangcode
func GetBlockContentFieldRevisionsViaDefaultLangcode(offset int, limit int, DefaultLangcode_ int, field string) (*[]*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = new([]*BlockContentFieldRevision)
	err := Engine.Table("block_content_field_revision").Where("default_langcode = ?", DefaultLangcode_).Limit(limit, offset).Desc(field).Find(_BlockContentFieldRevision)
	return _BlockContentFieldRevision, err
}

// HasBlockContentFieldRevisionViaId Has BlockContentFieldRevision via Id
func HasBlockContentFieldRevisionViaId(iId int64) bool {
	if has, err := Engine.Where("id = ?", iId).Get(new(BlockContentFieldRevision)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasBlockContentFieldRevisionViaRevisionId Has BlockContentFieldRevision via RevisionId
func HasBlockContentFieldRevisionViaRevisionId(iRevisionId int) bool {
	if has, err := Engine.Where("revision_id = ?", iRevisionId).Get(new(BlockContentFieldRevision)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasBlockContentFieldRevisionViaLangcode Has BlockContentFieldRevision via Langcode
func HasBlockContentFieldRevisionViaLangcode(iLangcode string) bool {
	if has, err := Engine.Where("langcode = ?", iLangcode).Get(new(BlockContentFieldRevision)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasBlockContentFieldRevisionViaInfo Has BlockContentFieldRevision via Info
func HasBlockContentFieldRevisionViaInfo(iInfo string) bool {
	if has, err := Engine.Where("info = ?", iInfo).Get(new(BlockContentFieldRevision)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasBlockContentFieldRevisionViaChanged Has BlockContentFieldRevision via Changed
func HasBlockContentFieldRevisionViaChanged(iChanged int) bool {
	if has, err := Engine.Where("changed = ?", iChanged).Get(new(BlockContentFieldRevision)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasBlockContentFieldRevisionViaRevisionCreated Has BlockContentFieldRevision via RevisionCreated
func HasBlockContentFieldRevisionViaRevisionCreated(iRevisionCreated int) bool {
	if has, err := Engine.Where("revision_created = ?", iRevisionCreated).Get(new(BlockContentFieldRevision)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasBlockContentFieldRevisionViaRevisionUser Has BlockContentFieldRevision via RevisionUser
func HasBlockContentFieldRevisionViaRevisionUser(iRevisionUser int) bool {
	if has, err := Engine.Where("revision_user = ?", iRevisionUser).Get(new(BlockContentFieldRevision)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasBlockContentFieldRevisionViaRevisionTranslationAffected Has BlockContentFieldRevision via RevisionTranslationAffected
func HasBlockContentFieldRevisionViaRevisionTranslationAffected(iRevisionTranslationAffected int) bool {
	if has, err := Engine.Where("revision_translation_affected = ?", iRevisionTranslationAffected).Get(new(BlockContentFieldRevision)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasBlockContentFieldRevisionViaDefaultLangcode Has BlockContentFieldRevision via DefaultLangcode
func HasBlockContentFieldRevisionViaDefaultLangcode(iDefaultLangcode int) bool {
	if has, err := Engine.Where("default_langcode = ?", iDefaultLangcode).Get(new(BlockContentFieldRevision)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// GetBlockContentFieldRevisionViaId Get BlockContentFieldRevision via Id
func GetBlockContentFieldRevisionViaId(iId int64) (*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = &BlockContentFieldRevision{Id: iId}
	has, err := Engine.Get(_BlockContentFieldRevision)
	if has {
		return _BlockContentFieldRevision, err
	} else {
		return nil, err
	}
}

// GetBlockContentFieldRevisionViaRevisionId Get BlockContentFieldRevision via RevisionId
func GetBlockContentFieldRevisionViaRevisionId(iRevisionId int) (*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = &BlockContentFieldRevision{RevisionId: iRevisionId}
	has, err := Engine.Get(_BlockContentFieldRevision)
	if has {
		return _BlockContentFieldRevision, err
	} else {
		return nil, err
	}
}

// GetBlockContentFieldRevisionViaLangcode Get BlockContentFieldRevision via Langcode
func GetBlockContentFieldRevisionViaLangcode(iLangcode string) (*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = &BlockContentFieldRevision{Langcode: iLangcode}
	has, err := Engine.Get(_BlockContentFieldRevision)
	if has {
		return _BlockContentFieldRevision, err
	} else {
		return nil, err
	}
}

// GetBlockContentFieldRevisionViaInfo Get BlockContentFieldRevision via Info
func GetBlockContentFieldRevisionViaInfo(iInfo string) (*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = &BlockContentFieldRevision{Info: iInfo}
	has, err := Engine.Get(_BlockContentFieldRevision)
	if has {
		return _BlockContentFieldRevision, err
	} else {
		return nil, err
	}
}

// GetBlockContentFieldRevisionViaChanged Get BlockContentFieldRevision via Changed
func GetBlockContentFieldRevisionViaChanged(iChanged int) (*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = &BlockContentFieldRevision{Changed: iChanged}
	has, err := Engine.Get(_BlockContentFieldRevision)
	if has {
		return _BlockContentFieldRevision, err
	} else {
		return nil, err
	}
}

// GetBlockContentFieldRevisionViaRevisionCreated Get BlockContentFieldRevision via RevisionCreated
func GetBlockContentFieldRevisionViaRevisionCreated(iRevisionCreated int) (*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = &BlockContentFieldRevision{RevisionCreated: iRevisionCreated}
	has, err := Engine.Get(_BlockContentFieldRevision)
	if has {
		return _BlockContentFieldRevision, err
	} else {
		return nil, err
	}
}

// GetBlockContentFieldRevisionViaRevisionUser Get BlockContentFieldRevision via RevisionUser
func GetBlockContentFieldRevisionViaRevisionUser(iRevisionUser int) (*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = &BlockContentFieldRevision{RevisionUser: iRevisionUser}
	has, err := Engine.Get(_BlockContentFieldRevision)
	if has {
		return _BlockContentFieldRevision, err
	} else {
		return nil, err
	}
}

// GetBlockContentFieldRevisionViaRevisionTranslationAffected Get BlockContentFieldRevision via RevisionTranslationAffected
func GetBlockContentFieldRevisionViaRevisionTranslationAffected(iRevisionTranslationAffected int) (*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = &BlockContentFieldRevision{RevisionTranslationAffected: iRevisionTranslationAffected}
	has, err := Engine.Get(_BlockContentFieldRevision)
	if has {
		return _BlockContentFieldRevision, err
	} else {
		return nil, err
	}
}

// GetBlockContentFieldRevisionViaDefaultLangcode Get BlockContentFieldRevision via DefaultLangcode
func GetBlockContentFieldRevisionViaDefaultLangcode(iDefaultLangcode int) (*BlockContentFieldRevision, error) {
	var _BlockContentFieldRevision = &BlockContentFieldRevision{DefaultLangcode: iDefaultLangcode}
	has, err := Engine.Get(_BlockContentFieldRevision)
	if has {
		return _BlockContentFieldRevision, err
	} else {
		return nil, err
	}
}

// SetBlockContentFieldRevisionViaId Set BlockContentFieldRevision via Id
func SetBlockContentFieldRevisionViaId(iId int64, block_content_field_revision *BlockContentFieldRevision) (int64, error) {
	block_content_field_revision.Id = iId
	return Engine.Insert(block_content_field_revision)
}

// SetBlockContentFieldRevisionViaRevisionId Set BlockContentFieldRevision via RevisionId
func SetBlockContentFieldRevisionViaRevisionId(iRevisionId int, block_content_field_revision *BlockContentFieldRevision) (int64, error) {
	block_content_field_revision.RevisionId = iRevisionId
	return Engine.Insert(block_content_field_revision)
}

// SetBlockContentFieldRevisionViaLangcode Set BlockContentFieldRevision via Langcode
func SetBlockContentFieldRevisionViaLangcode(iLangcode string, block_content_field_revision *BlockContentFieldRevision) (int64, error) {
	block_content_field_revision.Langcode = iLangcode
	return Engine.Insert(block_content_field_revision)
}

// SetBlockContentFieldRevisionViaInfo Set BlockContentFieldRevision via Info
func SetBlockContentFieldRevisionViaInfo(iInfo string, block_content_field_revision *BlockContentFieldRevision) (int64, error) {
	block_content_field_revision.Info = iInfo
	return Engine.Insert(block_content_field_revision)
}

// SetBlockContentFieldRevisionViaChanged Set BlockContentFieldRevision via Changed
func SetBlockContentFieldRevisionViaChanged(iChanged int, block_content_field_revision *BlockContentFieldRevision) (int64, error) {
	block_content_field_revision.Changed = iChanged
	return Engine.Insert(block_content_field_revision)
}

// SetBlockContentFieldRevisionViaRevisionCreated Set BlockContentFieldRevision via RevisionCreated
func SetBlockContentFieldRevisionViaRevisionCreated(iRevisionCreated int, block_content_field_revision *BlockContentFieldRevision) (int64, error) {
	block_content_field_revision.RevisionCreated = iRevisionCreated
	return Engine.Insert(block_content_field_revision)
}

// SetBlockContentFieldRevisionViaRevisionUser Set BlockContentFieldRevision via RevisionUser
func SetBlockContentFieldRevisionViaRevisionUser(iRevisionUser int, block_content_field_revision *BlockContentFieldRevision) (int64, error) {
	block_content_field_revision.RevisionUser = iRevisionUser
	return Engine.Insert(block_content_field_revision)
}

// SetBlockContentFieldRevisionViaRevisionTranslationAffected Set BlockContentFieldRevision via RevisionTranslationAffected
func SetBlockContentFieldRevisionViaRevisionTranslationAffected(iRevisionTranslationAffected int, block_content_field_revision *BlockContentFieldRevision) (int64, error) {
	block_content_field_revision.RevisionTranslationAffected = iRevisionTranslationAffected
	return Engine.Insert(block_content_field_revision)
}

// SetBlockContentFieldRevisionViaDefaultLangcode Set BlockContentFieldRevision via DefaultLangcode
func SetBlockContentFieldRevisionViaDefaultLangcode(iDefaultLangcode int, block_content_field_revision *BlockContentFieldRevision) (int64, error) {
	block_content_field_revision.DefaultLangcode = iDefaultLangcode
	return Engine.Insert(block_content_field_revision)
}

// AddBlockContentFieldRevision Add BlockContentFieldRevision via all columns
func AddBlockContentFieldRevision(iId int64, iRevisionId int, iLangcode string, iInfo string, iChanged int, iRevisionCreated int, iRevisionUser int, iRevisionTranslationAffected int, iDefaultLangcode int) error {
	_BlockContentFieldRevision := &BlockContentFieldRevision{Id: iId, RevisionId: iRevisionId, Langcode: iLangcode, Info: iInfo, Changed: iChanged, RevisionCreated: iRevisionCreated, RevisionUser: iRevisionUser, RevisionTranslationAffected: iRevisionTranslationAffected, DefaultLangcode: iDefaultLangcode}
	if _, err := Engine.Insert(_BlockContentFieldRevision); err != nil {
		return err
	} else {
		return nil
	}
}

// PostBlockContentFieldRevision Post BlockContentFieldRevision via iBlockContentFieldRevision
func PostBlockContentFieldRevision(iBlockContentFieldRevision *BlockContentFieldRevision) (int64, error) {
	_, err := Engine.Insert(iBlockContentFieldRevision)
	return iBlockContentFieldRevision.Id, err
}

// PutBlockContentFieldRevision Put BlockContentFieldRevision
func PutBlockContentFieldRevision(iBlockContentFieldRevision *BlockContentFieldRevision) (int64, error) {
	_, err := Engine.Id(iBlockContentFieldRevision.Id).Update(iBlockContentFieldRevision)
	return iBlockContentFieldRevision.Id, err
}

// PutBlockContentFieldRevisionViaId Put BlockContentFieldRevision via Id
func PutBlockContentFieldRevisionViaId(Id_ int64, iBlockContentFieldRevision *BlockContentFieldRevision) (int64, error) {
	row, err := Engine.Update(iBlockContentFieldRevision, &BlockContentFieldRevision{Id: Id_})
	return row, err
}

// PutBlockContentFieldRevisionViaRevisionId Put BlockContentFieldRevision via RevisionId
func PutBlockContentFieldRevisionViaRevisionId(RevisionId_ int, iBlockContentFieldRevision *BlockContentFieldRevision) (int64, error) {
	row, err := Engine.Update(iBlockContentFieldRevision, &BlockContentFieldRevision{RevisionId: RevisionId_})
	return row, err
}

// PutBlockContentFieldRevisionViaLangcode Put BlockContentFieldRevision via Langcode
func PutBlockContentFieldRevisionViaLangcode(Langcode_ string, iBlockContentFieldRevision *BlockContentFieldRevision) (int64, error) {
	row, err := Engine.Update(iBlockContentFieldRevision, &BlockContentFieldRevision{Langcode: Langcode_})
	return row, err
}

// PutBlockContentFieldRevisionViaInfo Put BlockContentFieldRevision via Info
func PutBlockContentFieldRevisionViaInfo(Info_ string, iBlockContentFieldRevision *BlockContentFieldRevision) (int64, error) {
	row, err := Engine.Update(iBlockContentFieldRevision, &BlockContentFieldRevision{Info: Info_})
	return row, err
}

// PutBlockContentFieldRevisionViaChanged Put BlockContentFieldRevision via Changed
func PutBlockContentFieldRevisionViaChanged(Changed_ int, iBlockContentFieldRevision *BlockContentFieldRevision) (int64, error) {
	row, err := Engine.Update(iBlockContentFieldRevision, &BlockContentFieldRevision{Changed: Changed_})
	return row, err
}

// PutBlockContentFieldRevisionViaRevisionCreated Put BlockContentFieldRevision via RevisionCreated
func PutBlockContentFieldRevisionViaRevisionCreated(RevisionCreated_ int, iBlockContentFieldRevision *BlockContentFieldRevision) (int64, error) {
	row, err := Engine.Update(iBlockContentFieldRevision, &BlockContentFieldRevision{RevisionCreated: RevisionCreated_})
	return row, err
}

// PutBlockContentFieldRevisionViaRevisionUser Put BlockContentFieldRevision via RevisionUser
func PutBlockContentFieldRevisionViaRevisionUser(RevisionUser_ int, iBlockContentFieldRevision *BlockContentFieldRevision) (int64, error) {
	row, err := Engine.Update(iBlockContentFieldRevision, &BlockContentFieldRevision{RevisionUser: RevisionUser_})
	return row, err
}

// PutBlockContentFieldRevisionViaRevisionTranslationAffected Put BlockContentFieldRevision via RevisionTranslationAffected
func PutBlockContentFieldRevisionViaRevisionTranslationAffected(RevisionTranslationAffected_ int, iBlockContentFieldRevision *BlockContentFieldRevision) (int64, error) {
	row, err := Engine.Update(iBlockContentFieldRevision, &BlockContentFieldRevision{RevisionTranslationAffected: RevisionTranslationAffected_})
	return row, err
}

// PutBlockContentFieldRevisionViaDefaultLangcode Put BlockContentFieldRevision via DefaultLangcode
func PutBlockContentFieldRevisionViaDefaultLangcode(DefaultLangcode_ int, iBlockContentFieldRevision *BlockContentFieldRevision) (int64, error) {
	row, err := Engine.Update(iBlockContentFieldRevision, &BlockContentFieldRevision{DefaultLangcode: DefaultLangcode_})
	return row, err
}

// UpdateBlockContentFieldRevisionViaId via map[string]interface{}{}
func UpdateBlockContentFieldRevisionViaId(iId int64, iBlockContentFieldRevisionMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(BlockContentFieldRevision)).Where("id = ?", iId).Update(iBlockContentFieldRevisionMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateBlockContentFieldRevisionViaRevisionId via map[string]interface{}{}
func UpdateBlockContentFieldRevisionViaRevisionId(iRevisionId int, iBlockContentFieldRevisionMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(BlockContentFieldRevision)).Where("revision_id = ?", iRevisionId).Update(iBlockContentFieldRevisionMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateBlockContentFieldRevisionViaLangcode via map[string]interface{}{}
func UpdateBlockContentFieldRevisionViaLangcode(iLangcode string, iBlockContentFieldRevisionMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(BlockContentFieldRevision)).Where("langcode = ?", iLangcode).Update(iBlockContentFieldRevisionMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateBlockContentFieldRevisionViaInfo via map[string]interface{}{}
func UpdateBlockContentFieldRevisionViaInfo(iInfo string, iBlockContentFieldRevisionMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(BlockContentFieldRevision)).Where("info = ?", iInfo).Update(iBlockContentFieldRevisionMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateBlockContentFieldRevisionViaChanged via map[string]interface{}{}
func UpdateBlockContentFieldRevisionViaChanged(iChanged int, iBlockContentFieldRevisionMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(BlockContentFieldRevision)).Where("changed = ?", iChanged).Update(iBlockContentFieldRevisionMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateBlockContentFieldRevisionViaRevisionCreated via map[string]interface{}{}
func UpdateBlockContentFieldRevisionViaRevisionCreated(iRevisionCreated int, iBlockContentFieldRevisionMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(BlockContentFieldRevision)).Where("revision_created = ?", iRevisionCreated).Update(iBlockContentFieldRevisionMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateBlockContentFieldRevisionViaRevisionUser via map[string]interface{}{}
func UpdateBlockContentFieldRevisionViaRevisionUser(iRevisionUser int, iBlockContentFieldRevisionMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(BlockContentFieldRevision)).Where("revision_user = ?", iRevisionUser).Update(iBlockContentFieldRevisionMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateBlockContentFieldRevisionViaRevisionTranslationAffected via map[string]interface{}{}
func UpdateBlockContentFieldRevisionViaRevisionTranslationAffected(iRevisionTranslationAffected int, iBlockContentFieldRevisionMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(BlockContentFieldRevision)).Where("revision_translation_affected = ?", iRevisionTranslationAffected).Update(iBlockContentFieldRevisionMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateBlockContentFieldRevisionViaDefaultLangcode via map[string]interface{}{}
func UpdateBlockContentFieldRevisionViaDefaultLangcode(iDefaultLangcode int, iBlockContentFieldRevisionMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(BlockContentFieldRevision)).Where("default_langcode = ?", iDefaultLangcode).Update(iBlockContentFieldRevisionMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// DeleteBlockContentFieldRevision Delete BlockContentFieldRevision
func DeleteBlockContentFieldRevision(iId int64) (int64, error) {
	row, err := Engine.Id(iId).Delete(new(BlockContentFieldRevision))
	return row, err
}

// DeleteBlockContentFieldRevisionViaId Delete BlockContentFieldRevision via Id
func DeleteBlockContentFieldRevisionViaId(iId int64) (err error) {
	var has bool
	var _BlockContentFieldRevision = &BlockContentFieldRevision{Id: iId}
	if has, err = Engine.Get(_BlockContentFieldRevision); (has == true) && (err == nil) {
		if row, err := Engine.Where("id = ?", iId).Delete(new(BlockContentFieldRevision)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteBlockContentFieldRevisionViaRevisionId Delete BlockContentFieldRevision via RevisionId
func DeleteBlockContentFieldRevisionViaRevisionId(iRevisionId int) (err error) {
	var has bool
	var _BlockContentFieldRevision = &BlockContentFieldRevision{RevisionId: iRevisionId}
	if has, err = Engine.Get(_BlockContentFieldRevision); (has == true) && (err == nil) {
		if row, err := Engine.Where("revision_id = ?", iRevisionId).Delete(new(BlockContentFieldRevision)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteBlockContentFieldRevisionViaLangcode Delete BlockContentFieldRevision via Langcode
func DeleteBlockContentFieldRevisionViaLangcode(iLangcode string) (err error) {
	var has bool
	var _BlockContentFieldRevision = &BlockContentFieldRevision{Langcode: iLangcode}
	if has, err = Engine.Get(_BlockContentFieldRevision); (has == true) && (err == nil) {
		if row, err := Engine.Where("langcode = ?", iLangcode).Delete(new(BlockContentFieldRevision)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteBlockContentFieldRevisionViaInfo Delete BlockContentFieldRevision via Info
func DeleteBlockContentFieldRevisionViaInfo(iInfo string) (err error) {
	var has bool
	var _BlockContentFieldRevision = &BlockContentFieldRevision{Info: iInfo}
	if has, err = Engine.Get(_BlockContentFieldRevision); (has == true) && (err == nil) {
		if row, err := Engine.Where("info = ?", iInfo).Delete(new(BlockContentFieldRevision)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteBlockContentFieldRevisionViaChanged Delete BlockContentFieldRevision via Changed
func DeleteBlockContentFieldRevisionViaChanged(iChanged int) (err error) {
	var has bool
	var _BlockContentFieldRevision = &BlockContentFieldRevision{Changed: iChanged}
	if has, err = Engine.Get(_BlockContentFieldRevision); (has == true) && (err == nil) {
		if row, err := Engine.Where("changed = ?", iChanged).Delete(new(BlockContentFieldRevision)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteBlockContentFieldRevisionViaRevisionCreated Delete BlockContentFieldRevision via RevisionCreated
func DeleteBlockContentFieldRevisionViaRevisionCreated(iRevisionCreated int) (err error) {
	var has bool
	var _BlockContentFieldRevision = &BlockContentFieldRevision{RevisionCreated: iRevisionCreated}
	if has, err = Engine.Get(_BlockContentFieldRevision); (has == true) && (err == nil) {
		if row, err := Engine.Where("revision_created = ?", iRevisionCreated).Delete(new(BlockContentFieldRevision)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteBlockContentFieldRevisionViaRevisionUser Delete BlockContentFieldRevision via RevisionUser
func DeleteBlockContentFieldRevisionViaRevisionUser(iRevisionUser int) (err error) {
	var has bool
	var _BlockContentFieldRevision = &BlockContentFieldRevision{RevisionUser: iRevisionUser}
	if has, err = Engine.Get(_BlockContentFieldRevision); (has == true) && (err == nil) {
		if row, err := Engine.Where("revision_user = ?", iRevisionUser).Delete(new(BlockContentFieldRevision)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteBlockContentFieldRevisionViaRevisionTranslationAffected Delete BlockContentFieldRevision via RevisionTranslationAffected
func DeleteBlockContentFieldRevisionViaRevisionTranslationAffected(iRevisionTranslationAffected int) (err error) {
	var has bool
	var _BlockContentFieldRevision = &BlockContentFieldRevision{RevisionTranslationAffected: iRevisionTranslationAffected}
	if has, err = Engine.Get(_BlockContentFieldRevision); (has == true) && (err == nil) {
		if row, err := Engine.Where("revision_translation_affected = ?", iRevisionTranslationAffected).Delete(new(BlockContentFieldRevision)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteBlockContentFieldRevisionViaDefaultLangcode Delete BlockContentFieldRevision via DefaultLangcode
func DeleteBlockContentFieldRevisionViaDefaultLangcode(iDefaultLangcode int) (err error) {
	var has bool
	var _BlockContentFieldRevision = &BlockContentFieldRevision{DefaultLangcode: iDefaultLangcode}
	if has, err = Engine.Get(_BlockContentFieldRevision); (has == true) && (err == nil) {
		if row, err := Engine.Where("default_langcode = ?", iDefaultLangcode).Delete(new(BlockContentFieldRevision)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}
