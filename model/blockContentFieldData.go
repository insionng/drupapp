package model

type BlockContentFieldData struct {
	Id                          int64  `xorm:"not null pk index(block_content__id__default_langcode__langcode) INT(10)"`
	RevisionId                  int    `xorm:"not null index INT(10)"`
	Type                        string `xorm:"not null index VARCHAR(32)"`
	Langcode                    string `xorm:"not null pk index(block_content__id__default_langcode__langcode) VARCHAR(12)"`
	Info                        string `xorm:"VARCHAR(255)"`
	Changed                     int    `xorm:"INT(11)"`
	RevisionCreated             int    `xorm:"INT(11)"`
	RevisionUser                int    `xorm:"index INT(10)"`
	RevisionTranslationAffected int    `xorm:"TINYINT(4)"`
	DefaultLangcode             int    `xorm:"not null index(block_content__id__default_langcode__langcode) TINYINT(4)"`
}

// GetBlockContentFieldDatasCount BlockContentFieldDatas' Count
func GetBlockContentFieldDatasCount(offset int, limit int) (int64, error) {
	total, err := Engine.Limit(limit, offset).Count(&BlockContentFieldData{})
	return total, err
}

// GetBlockContentFieldDataCountViaId Get BlockContentFieldData via Id
func GetBlockContentFieldDataCountViaId(iId int64) int64 {
	n, _ := Engine.Where("id = ?", iId).Count(&BlockContentFieldData{Id: iId})
	return n
}

// GetBlockContentFieldDataCountViaRevisionId Get BlockContentFieldData via RevisionId
func GetBlockContentFieldDataCountViaRevisionId(iRevisionId int) int64 {
	n, _ := Engine.Where("revision_id = ?", iRevisionId).Count(&BlockContentFieldData{RevisionId: iRevisionId})
	return n
}

// GetBlockContentFieldDataCountViaType Get BlockContentFieldData via Type
func GetBlockContentFieldDataCountViaType(iType string) int64 {
	n, _ := Engine.Where("type = ?", iType).Count(&BlockContentFieldData{Type: iType})
	return n
}

// GetBlockContentFieldDataCountViaLangcode Get BlockContentFieldData via Langcode
func GetBlockContentFieldDataCountViaLangcode(iLangcode string) int64 {
	n, _ := Engine.Where("langcode = ?", iLangcode).Count(&BlockContentFieldData{Langcode: iLangcode})
	return n
}

// GetBlockContentFieldDataCountViaInfo Get BlockContentFieldData via Info
func GetBlockContentFieldDataCountViaInfo(iInfo string) int64 {
	n, _ := Engine.Where("info = ?", iInfo).Count(&BlockContentFieldData{Info: iInfo})
	return n
}

// GetBlockContentFieldDataCountViaChanged Get BlockContentFieldData via Changed
func GetBlockContentFieldDataCountViaChanged(iChanged int) int64 {
	n, _ := Engine.Where("changed = ?", iChanged).Count(&BlockContentFieldData{Changed: iChanged})
	return n
}

// GetBlockContentFieldDataCountViaRevisionCreated Get BlockContentFieldData via RevisionCreated
func GetBlockContentFieldDataCountViaRevisionCreated(iRevisionCreated int) int64 {
	n, _ := Engine.Where("revision_created = ?", iRevisionCreated).Count(&BlockContentFieldData{RevisionCreated: iRevisionCreated})
	return n
}

// GetBlockContentFieldDataCountViaRevisionUser Get BlockContentFieldData via RevisionUser
func GetBlockContentFieldDataCountViaRevisionUser(iRevisionUser int) int64 {
	n, _ := Engine.Where("revision_user = ?", iRevisionUser).Count(&BlockContentFieldData{RevisionUser: iRevisionUser})
	return n
}

// GetBlockContentFieldDataCountViaRevisionTranslationAffected Get BlockContentFieldData via RevisionTranslationAffected
func GetBlockContentFieldDataCountViaRevisionTranslationAffected(iRevisionTranslationAffected int) int64 {
	n, _ := Engine.Where("revision_translation_affected = ?", iRevisionTranslationAffected).Count(&BlockContentFieldData{RevisionTranslationAffected: iRevisionTranslationAffected})
	return n
}

// GetBlockContentFieldDataCountViaDefaultLangcode Get BlockContentFieldData via DefaultLangcode
func GetBlockContentFieldDataCountViaDefaultLangcode(iDefaultLangcode int) int64 {
	n, _ := Engine.Where("default_langcode = ?", iDefaultLangcode).Count(&BlockContentFieldData{DefaultLangcode: iDefaultLangcode})
	return n
}

// GetBlockContentFieldDatasByIdAndRevisionIdAndType Get BlockContentFieldDatas via IdAndRevisionIdAndType
func GetBlockContentFieldDatasByIdAndRevisionIdAndType(offset int, limit int, Id_ int64, RevisionId_ int, Type_ string) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("id = ? and revision_id = ? and type = ?", Id_, RevisionId_, Type_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByIdAndRevisionIdAndLangcode Get BlockContentFieldDatas via IdAndRevisionIdAndLangcode
func GetBlockContentFieldDatasByIdAndRevisionIdAndLangcode(offset int, limit int, Id_ int64, RevisionId_ int, Langcode_ string) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("id = ? and revision_id = ? and langcode = ?", Id_, RevisionId_, Langcode_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByIdAndRevisionIdAndInfo Get BlockContentFieldDatas via IdAndRevisionIdAndInfo
func GetBlockContentFieldDatasByIdAndRevisionIdAndInfo(offset int, limit int, Id_ int64, RevisionId_ int, Info_ string) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("id = ? and revision_id = ? and info = ?", Id_, RevisionId_, Info_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByIdAndRevisionIdAndChanged Get BlockContentFieldDatas via IdAndRevisionIdAndChanged
func GetBlockContentFieldDatasByIdAndRevisionIdAndChanged(offset int, limit int, Id_ int64, RevisionId_ int, Changed_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("id = ? and revision_id = ? and changed = ?", Id_, RevisionId_, Changed_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByIdAndRevisionIdAndRevisionCreated Get BlockContentFieldDatas via IdAndRevisionIdAndRevisionCreated
func GetBlockContentFieldDatasByIdAndRevisionIdAndRevisionCreated(offset int, limit int, Id_ int64, RevisionId_ int, RevisionCreated_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("id = ? and revision_id = ? and revision_created = ?", Id_, RevisionId_, RevisionCreated_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByIdAndRevisionIdAndRevisionUser Get BlockContentFieldDatas via IdAndRevisionIdAndRevisionUser
func GetBlockContentFieldDatasByIdAndRevisionIdAndRevisionUser(offset int, limit int, Id_ int64, RevisionId_ int, RevisionUser_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("id = ? and revision_id = ? and revision_user = ?", Id_, RevisionId_, RevisionUser_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByIdAndRevisionIdAndRevisionTranslationAffected Get BlockContentFieldDatas via IdAndRevisionIdAndRevisionTranslationAffected
func GetBlockContentFieldDatasByIdAndRevisionIdAndRevisionTranslationAffected(offset int, limit int, Id_ int64, RevisionId_ int, RevisionTranslationAffected_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("id = ? and revision_id = ? and revision_translation_affected = ?", Id_, RevisionId_, RevisionTranslationAffected_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByIdAndRevisionIdAndDefaultLangcode Get BlockContentFieldDatas via IdAndRevisionIdAndDefaultLangcode
func GetBlockContentFieldDatasByIdAndRevisionIdAndDefaultLangcode(offset int, limit int, Id_ int64, RevisionId_ int, DefaultLangcode_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("id = ? and revision_id = ? and default_langcode = ?", Id_, RevisionId_, DefaultLangcode_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByIdAndTypeAndLangcode Get BlockContentFieldDatas via IdAndTypeAndLangcode
func GetBlockContentFieldDatasByIdAndTypeAndLangcode(offset int, limit int, Id_ int64, Type_ string, Langcode_ string) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("id = ? and type = ? and langcode = ?", Id_, Type_, Langcode_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByIdAndTypeAndInfo Get BlockContentFieldDatas via IdAndTypeAndInfo
func GetBlockContentFieldDatasByIdAndTypeAndInfo(offset int, limit int, Id_ int64, Type_ string, Info_ string) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("id = ? and type = ? and info = ?", Id_, Type_, Info_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByIdAndTypeAndChanged Get BlockContentFieldDatas via IdAndTypeAndChanged
func GetBlockContentFieldDatasByIdAndTypeAndChanged(offset int, limit int, Id_ int64, Type_ string, Changed_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("id = ? and type = ? and changed = ?", Id_, Type_, Changed_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByIdAndTypeAndRevisionCreated Get BlockContentFieldDatas via IdAndTypeAndRevisionCreated
func GetBlockContentFieldDatasByIdAndTypeAndRevisionCreated(offset int, limit int, Id_ int64, Type_ string, RevisionCreated_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("id = ? and type = ? and revision_created = ?", Id_, Type_, RevisionCreated_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByIdAndTypeAndRevisionUser Get BlockContentFieldDatas via IdAndTypeAndRevisionUser
func GetBlockContentFieldDatasByIdAndTypeAndRevisionUser(offset int, limit int, Id_ int64, Type_ string, RevisionUser_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("id = ? and type = ? and revision_user = ?", Id_, Type_, RevisionUser_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByIdAndTypeAndRevisionTranslationAffected Get BlockContentFieldDatas via IdAndTypeAndRevisionTranslationAffected
func GetBlockContentFieldDatasByIdAndTypeAndRevisionTranslationAffected(offset int, limit int, Id_ int64, Type_ string, RevisionTranslationAffected_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("id = ? and type = ? and revision_translation_affected = ?", Id_, Type_, RevisionTranslationAffected_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByIdAndTypeAndDefaultLangcode Get BlockContentFieldDatas via IdAndTypeAndDefaultLangcode
func GetBlockContentFieldDatasByIdAndTypeAndDefaultLangcode(offset int, limit int, Id_ int64, Type_ string, DefaultLangcode_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("id = ? and type = ? and default_langcode = ?", Id_, Type_, DefaultLangcode_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByIdAndLangcodeAndInfo Get BlockContentFieldDatas via IdAndLangcodeAndInfo
func GetBlockContentFieldDatasByIdAndLangcodeAndInfo(offset int, limit int, Id_ int64, Langcode_ string, Info_ string) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("id = ? and langcode = ? and info = ?", Id_, Langcode_, Info_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByIdAndLangcodeAndChanged Get BlockContentFieldDatas via IdAndLangcodeAndChanged
func GetBlockContentFieldDatasByIdAndLangcodeAndChanged(offset int, limit int, Id_ int64, Langcode_ string, Changed_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("id = ? and langcode = ? and changed = ?", Id_, Langcode_, Changed_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByIdAndLangcodeAndRevisionCreated Get BlockContentFieldDatas via IdAndLangcodeAndRevisionCreated
func GetBlockContentFieldDatasByIdAndLangcodeAndRevisionCreated(offset int, limit int, Id_ int64, Langcode_ string, RevisionCreated_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("id = ? and langcode = ? and revision_created = ?", Id_, Langcode_, RevisionCreated_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByIdAndLangcodeAndRevisionUser Get BlockContentFieldDatas via IdAndLangcodeAndRevisionUser
func GetBlockContentFieldDatasByIdAndLangcodeAndRevisionUser(offset int, limit int, Id_ int64, Langcode_ string, RevisionUser_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("id = ? and langcode = ? and revision_user = ?", Id_, Langcode_, RevisionUser_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByIdAndLangcodeAndRevisionTranslationAffected Get BlockContentFieldDatas via IdAndLangcodeAndRevisionTranslationAffected
func GetBlockContentFieldDatasByIdAndLangcodeAndRevisionTranslationAffected(offset int, limit int, Id_ int64, Langcode_ string, RevisionTranslationAffected_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("id = ? and langcode = ? and revision_translation_affected = ?", Id_, Langcode_, RevisionTranslationAffected_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByIdAndLangcodeAndDefaultLangcode Get BlockContentFieldDatas via IdAndLangcodeAndDefaultLangcode
func GetBlockContentFieldDatasByIdAndLangcodeAndDefaultLangcode(offset int, limit int, Id_ int64, Langcode_ string, DefaultLangcode_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("id = ? and langcode = ? and default_langcode = ?", Id_, Langcode_, DefaultLangcode_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByIdAndInfoAndChanged Get BlockContentFieldDatas via IdAndInfoAndChanged
func GetBlockContentFieldDatasByIdAndInfoAndChanged(offset int, limit int, Id_ int64, Info_ string, Changed_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("id = ? and info = ? and changed = ?", Id_, Info_, Changed_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByIdAndInfoAndRevisionCreated Get BlockContentFieldDatas via IdAndInfoAndRevisionCreated
func GetBlockContentFieldDatasByIdAndInfoAndRevisionCreated(offset int, limit int, Id_ int64, Info_ string, RevisionCreated_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("id = ? and info = ? and revision_created = ?", Id_, Info_, RevisionCreated_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByIdAndInfoAndRevisionUser Get BlockContentFieldDatas via IdAndInfoAndRevisionUser
func GetBlockContentFieldDatasByIdAndInfoAndRevisionUser(offset int, limit int, Id_ int64, Info_ string, RevisionUser_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("id = ? and info = ? and revision_user = ?", Id_, Info_, RevisionUser_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByIdAndInfoAndRevisionTranslationAffected Get BlockContentFieldDatas via IdAndInfoAndRevisionTranslationAffected
func GetBlockContentFieldDatasByIdAndInfoAndRevisionTranslationAffected(offset int, limit int, Id_ int64, Info_ string, RevisionTranslationAffected_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("id = ? and info = ? and revision_translation_affected = ?", Id_, Info_, RevisionTranslationAffected_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByIdAndInfoAndDefaultLangcode Get BlockContentFieldDatas via IdAndInfoAndDefaultLangcode
func GetBlockContentFieldDatasByIdAndInfoAndDefaultLangcode(offset int, limit int, Id_ int64, Info_ string, DefaultLangcode_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("id = ? and info = ? and default_langcode = ?", Id_, Info_, DefaultLangcode_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByIdAndChangedAndRevisionCreated Get BlockContentFieldDatas via IdAndChangedAndRevisionCreated
func GetBlockContentFieldDatasByIdAndChangedAndRevisionCreated(offset int, limit int, Id_ int64, Changed_ int, RevisionCreated_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("id = ? and changed = ? and revision_created = ?", Id_, Changed_, RevisionCreated_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByIdAndChangedAndRevisionUser Get BlockContentFieldDatas via IdAndChangedAndRevisionUser
func GetBlockContentFieldDatasByIdAndChangedAndRevisionUser(offset int, limit int, Id_ int64, Changed_ int, RevisionUser_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("id = ? and changed = ? and revision_user = ?", Id_, Changed_, RevisionUser_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByIdAndChangedAndRevisionTranslationAffected Get BlockContentFieldDatas via IdAndChangedAndRevisionTranslationAffected
func GetBlockContentFieldDatasByIdAndChangedAndRevisionTranslationAffected(offset int, limit int, Id_ int64, Changed_ int, RevisionTranslationAffected_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("id = ? and changed = ? and revision_translation_affected = ?", Id_, Changed_, RevisionTranslationAffected_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByIdAndChangedAndDefaultLangcode Get BlockContentFieldDatas via IdAndChangedAndDefaultLangcode
func GetBlockContentFieldDatasByIdAndChangedAndDefaultLangcode(offset int, limit int, Id_ int64, Changed_ int, DefaultLangcode_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("id = ? and changed = ? and default_langcode = ?", Id_, Changed_, DefaultLangcode_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByIdAndRevisionCreatedAndRevisionUser Get BlockContentFieldDatas via IdAndRevisionCreatedAndRevisionUser
func GetBlockContentFieldDatasByIdAndRevisionCreatedAndRevisionUser(offset int, limit int, Id_ int64, RevisionCreated_ int, RevisionUser_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("id = ? and revision_created = ? and revision_user = ?", Id_, RevisionCreated_, RevisionUser_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByIdAndRevisionCreatedAndRevisionTranslationAffected Get BlockContentFieldDatas via IdAndRevisionCreatedAndRevisionTranslationAffected
func GetBlockContentFieldDatasByIdAndRevisionCreatedAndRevisionTranslationAffected(offset int, limit int, Id_ int64, RevisionCreated_ int, RevisionTranslationAffected_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("id = ? and revision_created = ? and revision_translation_affected = ?", Id_, RevisionCreated_, RevisionTranslationAffected_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByIdAndRevisionCreatedAndDefaultLangcode Get BlockContentFieldDatas via IdAndRevisionCreatedAndDefaultLangcode
func GetBlockContentFieldDatasByIdAndRevisionCreatedAndDefaultLangcode(offset int, limit int, Id_ int64, RevisionCreated_ int, DefaultLangcode_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("id = ? and revision_created = ? and default_langcode = ?", Id_, RevisionCreated_, DefaultLangcode_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByIdAndRevisionUserAndRevisionTranslationAffected Get BlockContentFieldDatas via IdAndRevisionUserAndRevisionTranslationAffected
func GetBlockContentFieldDatasByIdAndRevisionUserAndRevisionTranslationAffected(offset int, limit int, Id_ int64, RevisionUser_ int, RevisionTranslationAffected_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("id = ? and revision_user = ? and revision_translation_affected = ?", Id_, RevisionUser_, RevisionTranslationAffected_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByIdAndRevisionUserAndDefaultLangcode Get BlockContentFieldDatas via IdAndRevisionUserAndDefaultLangcode
func GetBlockContentFieldDatasByIdAndRevisionUserAndDefaultLangcode(offset int, limit int, Id_ int64, RevisionUser_ int, DefaultLangcode_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("id = ? and revision_user = ? and default_langcode = ?", Id_, RevisionUser_, DefaultLangcode_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByIdAndRevisionTranslationAffectedAndDefaultLangcode Get BlockContentFieldDatas via IdAndRevisionTranslationAffectedAndDefaultLangcode
func GetBlockContentFieldDatasByIdAndRevisionTranslationAffectedAndDefaultLangcode(offset int, limit int, Id_ int64, RevisionTranslationAffected_ int, DefaultLangcode_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("id = ? and revision_translation_affected = ? and default_langcode = ?", Id_, RevisionTranslationAffected_, DefaultLangcode_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByRevisionIdAndTypeAndLangcode Get BlockContentFieldDatas via RevisionIdAndTypeAndLangcode
func GetBlockContentFieldDatasByRevisionIdAndTypeAndLangcode(offset int, limit int, RevisionId_ int, Type_ string, Langcode_ string) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("revision_id = ? and type = ? and langcode = ?", RevisionId_, Type_, Langcode_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByRevisionIdAndTypeAndInfo Get BlockContentFieldDatas via RevisionIdAndTypeAndInfo
func GetBlockContentFieldDatasByRevisionIdAndTypeAndInfo(offset int, limit int, RevisionId_ int, Type_ string, Info_ string) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("revision_id = ? and type = ? and info = ?", RevisionId_, Type_, Info_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByRevisionIdAndTypeAndChanged Get BlockContentFieldDatas via RevisionIdAndTypeAndChanged
func GetBlockContentFieldDatasByRevisionIdAndTypeAndChanged(offset int, limit int, RevisionId_ int, Type_ string, Changed_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("revision_id = ? and type = ? and changed = ?", RevisionId_, Type_, Changed_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByRevisionIdAndTypeAndRevisionCreated Get BlockContentFieldDatas via RevisionIdAndTypeAndRevisionCreated
func GetBlockContentFieldDatasByRevisionIdAndTypeAndRevisionCreated(offset int, limit int, RevisionId_ int, Type_ string, RevisionCreated_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("revision_id = ? and type = ? and revision_created = ?", RevisionId_, Type_, RevisionCreated_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByRevisionIdAndTypeAndRevisionUser Get BlockContentFieldDatas via RevisionIdAndTypeAndRevisionUser
func GetBlockContentFieldDatasByRevisionIdAndTypeAndRevisionUser(offset int, limit int, RevisionId_ int, Type_ string, RevisionUser_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("revision_id = ? and type = ? and revision_user = ?", RevisionId_, Type_, RevisionUser_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByRevisionIdAndTypeAndRevisionTranslationAffected Get BlockContentFieldDatas via RevisionIdAndTypeAndRevisionTranslationAffected
func GetBlockContentFieldDatasByRevisionIdAndTypeAndRevisionTranslationAffected(offset int, limit int, RevisionId_ int, Type_ string, RevisionTranslationAffected_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("revision_id = ? and type = ? and revision_translation_affected = ?", RevisionId_, Type_, RevisionTranslationAffected_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByRevisionIdAndTypeAndDefaultLangcode Get BlockContentFieldDatas via RevisionIdAndTypeAndDefaultLangcode
func GetBlockContentFieldDatasByRevisionIdAndTypeAndDefaultLangcode(offset int, limit int, RevisionId_ int, Type_ string, DefaultLangcode_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("revision_id = ? and type = ? and default_langcode = ?", RevisionId_, Type_, DefaultLangcode_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByRevisionIdAndLangcodeAndInfo Get BlockContentFieldDatas via RevisionIdAndLangcodeAndInfo
func GetBlockContentFieldDatasByRevisionIdAndLangcodeAndInfo(offset int, limit int, RevisionId_ int, Langcode_ string, Info_ string) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("revision_id = ? and langcode = ? and info = ?", RevisionId_, Langcode_, Info_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByRevisionIdAndLangcodeAndChanged Get BlockContentFieldDatas via RevisionIdAndLangcodeAndChanged
func GetBlockContentFieldDatasByRevisionIdAndLangcodeAndChanged(offset int, limit int, RevisionId_ int, Langcode_ string, Changed_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("revision_id = ? and langcode = ? and changed = ?", RevisionId_, Langcode_, Changed_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByRevisionIdAndLangcodeAndRevisionCreated Get BlockContentFieldDatas via RevisionIdAndLangcodeAndRevisionCreated
func GetBlockContentFieldDatasByRevisionIdAndLangcodeAndRevisionCreated(offset int, limit int, RevisionId_ int, Langcode_ string, RevisionCreated_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("revision_id = ? and langcode = ? and revision_created = ?", RevisionId_, Langcode_, RevisionCreated_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByRevisionIdAndLangcodeAndRevisionUser Get BlockContentFieldDatas via RevisionIdAndLangcodeAndRevisionUser
func GetBlockContentFieldDatasByRevisionIdAndLangcodeAndRevisionUser(offset int, limit int, RevisionId_ int, Langcode_ string, RevisionUser_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("revision_id = ? and langcode = ? and revision_user = ?", RevisionId_, Langcode_, RevisionUser_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByRevisionIdAndLangcodeAndRevisionTranslationAffected Get BlockContentFieldDatas via RevisionIdAndLangcodeAndRevisionTranslationAffected
func GetBlockContentFieldDatasByRevisionIdAndLangcodeAndRevisionTranslationAffected(offset int, limit int, RevisionId_ int, Langcode_ string, RevisionTranslationAffected_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("revision_id = ? and langcode = ? and revision_translation_affected = ?", RevisionId_, Langcode_, RevisionTranslationAffected_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByRevisionIdAndLangcodeAndDefaultLangcode Get BlockContentFieldDatas via RevisionIdAndLangcodeAndDefaultLangcode
func GetBlockContentFieldDatasByRevisionIdAndLangcodeAndDefaultLangcode(offset int, limit int, RevisionId_ int, Langcode_ string, DefaultLangcode_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("revision_id = ? and langcode = ? and default_langcode = ?", RevisionId_, Langcode_, DefaultLangcode_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByRevisionIdAndInfoAndChanged Get BlockContentFieldDatas via RevisionIdAndInfoAndChanged
func GetBlockContentFieldDatasByRevisionIdAndInfoAndChanged(offset int, limit int, RevisionId_ int, Info_ string, Changed_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("revision_id = ? and info = ? and changed = ?", RevisionId_, Info_, Changed_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByRevisionIdAndInfoAndRevisionCreated Get BlockContentFieldDatas via RevisionIdAndInfoAndRevisionCreated
func GetBlockContentFieldDatasByRevisionIdAndInfoAndRevisionCreated(offset int, limit int, RevisionId_ int, Info_ string, RevisionCreated_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("revision_id = ? and info = ? and revision_created = ?", RevisionId_, Info_, RevisionCreated_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByRevisionIdAndInfoAndRevisionUser Get BlockContentFieldDatas via RevisionIdAndInfoAndRevisionUser
func GetBlockContentFieldDatasByRevisionIdAndInfoAndRevisionUser(offset int, limit int, RevisionId_ int, Info_ string, RevisionUser_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("revision_id = ? and info = ? and revision_user = ?", RevisionId_, Info_, RevisionUser_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByRevisionIdAndInfoAndRevisionTranslationAffected Get BlockContentFieldDatas via RevisionIdAndInfoAndRevisionTranslationAffected
func GetBlockContentFieldDatasByRevisionIdAndInfoAndRevisionTranslationAffected(offset int, limit int, RevisionId_ int, Info_ string, RevisionTranslationAffected_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("revision_id = ? and info = ? and revision_translation_affected = ?", RevisionId_, Info_, RevisionTranslationAffected_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByRevisionIdAndInfoAndDefaultLangcode Get BlockContentFieldDatas via RevisionIdAndInfoAndDefaultLangcode
func GetBlockContentFieldDatasByRevisionIdAndInfoAndDefaultLangcode(offset int, limit int, RevisionId_ int, Info_ string, DefaultLangcode_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("revision_id = ? and info = ? and default_langcode = ?", RevisionId_, Info_, DefaultLangcode_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByRevisionIdAndChangedAndRevisionCreated Get BlockContentFieldDatas via RevisionIdAndChangedAndRevisionCreated
func GetBlockContentFieldDatasByRevisionIdAndChangedAndRevisionCreated(offset int, limit int, RevisionId_ int, Changed_ int, RevisionCreated_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("revision_id = ? and changed = ? and revision_created = ?", RevisionId_, Changed_, RevisionCreated_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByRevisionIdAndChangedAndRevisionUser Get BlockContentFieldDatas via RevisionIdAndChangedAndRevisionUser
func GetBlockContentFieldDatasByRevisionIdAndChangedAndRevisionUser(offset int, limit int, RevisionId_ int, Changed_ int, RevisionUser_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("revision_id = ? and changed = ? and revision_user = ?", RevisionId_, Changed_, RevisionUser_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByRevisionIdAndChangedAndRevisionTranslationAffected Get BlockContentFieldDatas via RevisionIdAndChangedAndRevisionTranslationAffected
func GetBlockContentFieldDatasByRevisionIdAndChangedAndRevisionTranslationAffected(offset int, limit int, RevisionId_ int, Changed_ int, RevisionTranslationAffected_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("revision_id = ? and changed = ? and revision_translation_affected = ?", RevisionId_, Changed_, RevisionTranslationAffected_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByRevisionIdAndChangedAndDefaultLangcode Get BlockContentFieldDatas via RevisionIdAndChangedAndDefaultLangcode
func GetBlockContentFieldDatasByRevisionIdAndChangedAndDefaultLangcode(offset int, limit int, RevisionId_ int, Changed_ int, DefaultLangcode_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("revision_id = ? and changed = ? and default_langcode = ?", RevisionId_, Changed_, DefaultLangcode_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByRevisionIdAndRevisionCreatedAndRevisionUser Get BlockContentFieldDatas via RevisionIdAndRevisionCreatedAndRevisionUser
func GetBlockContentFieldDatasByRevisionIdAndRevisionCreatedAndRevisionUser(offset int, limit int, RevisionId_ int, RevisionCreated_ int, RevisionUser_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("revision_id = ? and revision_created = ? and revision_user = ?", RevisionId_, RevisionCreated_, RevisionUser_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByRevisionIdAndRevisionCreatedAndRevisionTranslationAffected Get BlockContentFieldDatas via RevisionIdAndRevisionCreatedAndRevisionTranslationAffected
func GetBlockContentFieldDatasByRevisionIdAndRevisionCreatedAndRevisionTranslationAffected(offset int, limit int, RevisionId_ int, RevisionCreated_ int, RevisionTranslationAffected_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("revision_id = ? and revision_created = ? and revision_translation_affected = ?", RevisionId_, RevisionCreated_, RevisionTranslationAffected_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByRevisionIdAndRevisionCreatedAndDefaultLangcode Get BlockContentFieldDatas via RevisionIdAndRevisionCreatedAndDefaultLangcode
func GetBlockContentFieldDatasByRevisionIdAndRevisionCreatedAndDefaultLangcode(offset int, limit int, RevisionId_ int, RevisionCreated_ int, DefaultLangcode_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("revision_id = ? and revision_created = ? and default_langcode = ?", RevisionId_, RevisionCreated_, DefaultLangcode_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByRevisionIdAndRevisionUserAndRevisionTranslationAffected Get BlockContentFieldDatas via RevisionIdAndRevisionUserAndRevisionTranslationAffected
func GetBlockContentFieldDatasByRevisionIdAndRevisionUserAndRevisionTranslationAffected(offset int, limit int, RevisionId_ int, RevisionUser_ int, RevisionTranslationAffected_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("revision_id = ? and revision_user = ? and revision_translation_affected = ?", RevisionId_, RevisionUser_, RevisionTranslationAffected_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByRevisionIdAndRevisionUserAndDefaultLangcode Get BlockContentFieldDatas via RevisionIdAndRevisionUserAndDefaultLangcode
func GetBlockContentFieldDatasByRevisionIdAndRevisionUserAndDefaultLangcode(offset int, limit int, RevisionId_ int, RevisionUser_ int, DefaultLangcode_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("revision_id = ? and revision_user = ? and default_langcode = ?", RevisionId_, RevisionUser_, DefaultLangcode_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByRevisionIdAndRevisionTranslationAffectedAndDefaultLangcode Get BlockContentFieldDatas via RevisionIdAndRevisionTranslationAffectedAndDefaultLangcode
func GetBlockContentFieldDatasByRevisionIdAndRevisionTranslationAffectedAndDefaultLangcode(offset int, limit int, RevisionId_ int, RevisionTranslationAffected_ int, DefaultLangcode_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("revision_id = ? and revision_translation_affected = ? and default_langcode = ?", RevisionId_, RevisionTranslationAffected_, DefaultLangcode_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByTypeAndLangcodeAndInfo Get BlockContentFieldDatas via TypeAndLangcodeAndInfo
func GetBlockContentFieldDatasByTypeAndLangcodeAndInfo(offset int, limit int, Type_ string, Langcode_ string, Info_ string) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("type = ? and langcode = ? and info = ?", Type_, Langcode_, Info_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByTypeAndLangcodeAndChanged Get BlockContentFieldDatas via TypeAndLangcodeAndChanged
func GetBlockContentFieldDatasByTypeAndLangcodeAndChanged(offset int, limit int, Type_ string, Langcode_ string, Changed_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("type = ? and langcode = ? and changed = ?", Type_, Langcode_, Changed_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByTypeAndLangcodeAndRevisionCreated Get BlockContentFieldDatas via TypeAndLangcodeAndRevisionCreated
func GetBlockContentFieldDatasByTypeAndLangcodeAndRevisionCreated(offset int, limit int, Type_ string, Langcode_ string, RevisionCreated_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("type = ? and langcode = ? and revision_created = ?", Type_, Langcode_, RevisionCreated_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByTypeAndLangcodeAndRevisionUser Get BlockContentFieldDatas via TypeAndLangcodeAndRevisionUser
func GetBlockContentFieldDatasByTypeAndLangcodeAndRevisionUser(offset int, limit int, Type_ string, Langcode_ string, RevisionUser_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("type = ? and langcode = ? and revision_user = ?", Type_, Langcode_, RevisionUser_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByTypeAndLangcodeAndRevisionTranslationAffected Get BlockContentFieldDatas via TypeAndLangcodeAndRevisionTranslationAffected
func GetBlockContentFieldDatasByTypeAndLangcodeAndRevisionTranslationAffected(offset int, limit int, Type_ string, Langcode_ string, RevisionTranslationAffected_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("type = ? and langcode = ? and revision_translation_affected = ?", Type_, Langcode_, RevisionTranslationAffected_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByTypeAndLangcodeAndDefaultLangcode Get BlockContentFieldDatas via TypeAndLangcodeAndDefaultLangcode
func GetBlockContentFieldDatasByTypeAndLangcodeAndDefaultLangcode(offset int, limit int, Type_ string, Langcode_ string, DefaultLangcode_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("type = ? and langcode = ? and default_langcode = ?", Type_, Langcode_, DefaultLangcode_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByTypeAndInfoAndChanged Get BlockContentFieldDatas via TypeAndInfoAndChanged
func GetBlockContentFieldDatasByTypeAndInfoAndChanged(offset int, limit int, Type_ string, Info_ string, Changed_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("type = ? and info = ? and changed = ?", Type_, Info_, Changed_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByTypeAndInfoAndRevisionCreated Get BlockContentFieldDatas via TypeAndInfoAndRevisionCreated
func GetBlockContentFieldDatasByTypeAndInfoAndRevisionCreated(offset int, limit int, Type_ string, Info_ string, RevisionCreated_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("type = ? and info = ? and revision_created = ?", Type_, Info_, RevisionCreated_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByTypeAndInfoAndRevisionUser Get BlockContentFieldDatas via TypeAndInfoAndRevisionUser
func GetBlockContentFieldDatasByTypeAndInfoAndRevisionUser(offset int, limit int, Type_ string, Info_ string, RevisionUser_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("type = ? and info = ? and revision_user = ?", Type_, Info_, RevisionUser_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByTypeAndInfoAndRevisionTranslationAffected Get BlockContentFieldDatas via TypeAndInfoAndRevisionTranslationAffected
func GetBlockContentFieldDatasByTypeAndInfoAndRevisionTranslationAffected(offset int, limit int, Type_ string, Info_ string, RevisionTranslationAffected_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("type = ? and info = ? and revision_translation_affected = ?", Type_, Info_, RevisionTranslationAffected_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByTypeAndInfoAndDefaultLangcode Get BlockContentFieldDatas via TypeAndInfoAndDefaultLangcode
func GetBlockContentFieldDatasByTypeAndInfoAndDefaultLangcode(offset int, limit int, Type_ string, Info_ string, DefaultLangcode_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("type = ? and info = ? and default_langcode = ?", Type_, Info_, DefaultLangcode_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByTypeAndChangedAndRevisionCreated Get BlockContentFieldDatas via TypeAndChangedAndRevisionCreated
func GetBlockContentFieldDatasByTypeAndChangedAndRevisionCreated(offset int, limit int, Type_ string, Changed_ int, RevisionCreated_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("type = ? and changed = ? and revision_created = ?", Type_, Changed_, RevisionCreated_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByTypeAndChangedAndRevisionUser Get BlockContentFieldDatas via TypeAndChangedAndRevisionUser
func GetBlockContentFieldDatasByTypeAndChangedAndRevisionUser(offset int, limit int, Type_ string, Changed_ int, RevisionUser_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("type = ? and changed = ? and revision_user = ?", Type_, Changed_, RevisionUser_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByTypeAndChangedAndRevisionTranslationAffected Get BlockContentFieldDatas via TypeAndChangedAndRevisionTranslationAffected
func GetBlockContentFieldDatasByTypeAndChangedAndRevisionTranslationAffected(offset int, limit int, Type_ string, Changed_ int, RevisionTranslationAffected_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("type = ? and changed = ? and revision_translation_affected = ?", Type_, Changed_, RevisionTranslationAffected_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByTypeAndChangedAndDefaultLangcode Get BlockContentFieldDatas via TypeAndChangedAndDefaultLangcode
func GetBlockContentFieldDatasByTypeAndChangedAndDefaultLangcode(offset int, limit int, Type_ string, Changed_ int, DefaultLangcode_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("type = ? and changed = ? and default_langcode = ?", Type_, Changed_, DefaultLangcode_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByTypeAndRevisionCreatedAndRevisionUser Get BlockContentFieldDatas via TypeAndRevisionCreatedAndRevisionUser
func GetBlockContentFieldDatasByTypeAndRevisionCreatedAndRevisionUser(offset int, limit int, Type_ string, RevisionCreated_ int, RevisionUser_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("type = ? and revision_created = ? and revision_user = ?", Type_, RevisionCreated_, RevisionUser_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByTypeAndRevisionCreatedAndRevisionTranslationAffected Get BlockContentFieldDatas via TypeAndRevisionCreatedAndRevisionTranslationAffected
func GetBlockContentFieldDatasByTypeAndRevisionCreatedAndRevisionTranslationAffected(offset int, limit int, Type_ string, RevisionCreated_ int, RevisionTranslationAffected_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("type = ? and revision_created = ? and revision_translation_affected = ?", Type_, RevisionCreated_, RevisionTranslationAffected_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByTypeAndRevisionCreatedAndDefaultLangcode Get BlockContentFieldDatas via TypeAndRevisionCreatedAndDefaultLangcode
func GetBlockContentFieldDatasByTypeAndRevisionCreatedAndDefaultLangcode(offset int, limit int, Type_ string, RevisionCreated_ int, DefaultLangcode_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("type = ? and revision_created = ? and default_langcode = ?", Type_, RevisionCreated_, DefaultLangcode_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByTypeAndRevisionUserAndRevisionTranslationAffected Get BlockContentFieldDatas via TypeAndRevisionUserAndRevisionTranslationAffected
func GetBlockContentFieldDatasByTypeAndRevisionUserAndRevisionTranslationAffected(offset int, limit int, Type_ string, RevisionUser_ int, RevisionTranslationAffected_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("type = ? and revision_user = ? and revision_translation_affected = ?", Type_, RevisionUser_, RevisionTranslationAffected_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByTypeAndRevisionUserAndDefaultLangcode Get BlockContentFieldDatas via TypeAndRevisionUserAndDefaultLangcode
func GetBlockContentFieldDatasByTypeAndRevisionUserAndDefaultLangcode(offset int, limit int, Type_ string, RevisionUser_ int, DefaultLangcode_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("type = ? and revision_user = ? and default_langcode = ?", Type_, RevisionUser_, DefaultLangcode_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByTypeAndRevisionTranslationAffectedAndDefaultLangcode Get BlockContentFieldDatas via TypeAndRevisionTranslationAffectedAndDefaultLangcode
func GetBlockContentFieldDatasByTypeAndRevisionTranslationAffectedAndDefaultLangcode(offset int, limit int, Type_ string, RevisionTranslationAffected_ int, DefaultLangcode_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("type = ? and revision_translation_affected = ? and default_langcode = ?", Type_, RevisionTranslationAffected_, DefaultLangcode_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByLangcodeAndInfoAndChanged Get BlockContentFieldDatas via LangcodeAndInfoAndChanged
func GetBlockContentFieldDatasByLangcodeAndInfoAndChanged(offset int, limit int, Langcode_ string, Info_ string, Changed_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("langcode = ? and info = ? and changed = ?", Langcode_, Info_, Changed_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByLangcodeAndInfoAndRevisionCreated Get BlockContentFieldDatas via LangcodeAndInfoAndRevisionCreated
func GetBlockContentFieldDatasByLangcodeAndInfoAndRevisionCreated(offset int, limit int, Langcode_ string, Info_ string, RevisionCreated_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("langcode = ? and info = ? and revision_created = ?", Langcode_, Info_, RevisionCreated_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByLangcodeAndInfoAndRevisionUser Get BlockContentFieldDatas via LangcodeAndInfoAndRevisionUser
func GetBlockContentFieldDatasByLangcodeAndInfoAndRevisionUser(offset int, limit int, Langcode_ string, Info_ string, RevisionUser_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("langcode = ? and info = ? and revision_user = ?", Langcode_, Info_, RevisionUser_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByLangcodeAndInfoAndRevisionTranslationAffected Get BlockContentFieldDatas via LangcodeAndInfoAndRevisionTranslationAffected
func GetBlockContentFieldDatasByLangcodeAndInfoAndRevisionTranslationAffected(offset int, limit int, Langcode_ string, Info_ string, RevisionTranslationAffected_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("langcode = ? and info = ? and revision_translation_affected = ?", Langcode_, Info_, RevisionTranslationAffected_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByLangcodeAndInfoAndDefaultLangcode Get BlockContentFieldDatas via LangcodeAndInfoAndDefaultLangcode
func GetBlockContentFieldDatasByLangcodeAndInfoAndDefaultLangcode(offset int, limit int, Langcode_ string, Info_ string, DefaultLangcode_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("langcode = ? and info = ? and default_langcode = ?", Langcode_, Info_, DefaultLangcode_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByLangcodeAndChangedAndRevisionCreated Get BlockContentFieldDatas via LangcodeAndChangedAndRevisionCreated
func GetBlockContentFieldDatasByLangcodeAndChangedAndRevisionCreated(offset int, limit int, Langcode_ string, Changed_ int, RevisionCreated_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("langcode = ? and changed = ? and revision_created = ?", Langcode_, Changed_, RevisionCreated_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByLangcodeAndChangedAndRevisionUser Get BlockContentFieldDatas via LangcodeAndChangedAndRevisionUser
func GetBlockContentFieldDatasByLangcodeAndChangedAndRevisionUser(offset int, limit int, Langcode_ string, Changed_ int, RevisionUser_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("langcode = ? and changed = ? and revision_user = ?", Langcode_, Changed_, RevisionUser_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByLangcodeAndChangedAndRevisionTranslationAffected Get BlockContentFieldDatas via LangcodeAndChangedAndRevisionTranslationAffected
func GetBlockContentFieldDatasByLangcodeAndChangedAndRevisionTranslationAffected(offset int, limit int, Langcode_ string, Changed_ int, RevisionTranslationAffected_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("langcode = ? and changed = ? and revision_translation_affected = ?", Langcode_, Changed_, RevisionTranslationAffected_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByLangcodeAndChangedAndDefaultLangcode Get BlockContentFieldDatas via LangcodeAndChangedAndDefaultLangcode
func GetBlockContentFieldDatasByLangcodeAndChangedAndDefaultLangcode(offset int, limit int, Langcode_ string, Changed_ int, DefaultLangcode_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("langcode = ? and changed = ? and default_langcode = ?", Langcode_, Changed_, DefaultLangcode_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByLangcodeAndRevisionCreatedAndRevisionUser Get BlockContentFieldDatas via LangcodeAndRevisionCreatedAndRevisionUser
func GetBlockContentFieldDatasByLangcodeAndRevisionCreatedAndRevisionUser(offset int, limit int, Langcode_ string, RevisionCreated_ int, RevisionUser_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("langcode = ? and revision_created = ? and revision_user = ?", Langcode_, RevisionCreated_, RevisionUser_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByLangcodeAndRevisionCreatedAndRevisionTranslationAffected Get BlockContentFieldDatas via LangcodeAndRevisionCreatedAndRevisionTranslationAffected
func GetBlockContentFieldDatasByLangcodeAndRevisionCreatedAndRevisionTranslationAffected(offset int, limit int, Langcode_ string, RevisionCreated_ int, RevisionTranslationAffected_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("langcode = ? and revision_created = ? and revision_translation_affected = ?", Langcode_, RevisionCreated_, RevisionTranslationAffected_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByLangcodeAndRevisionCreatedAndDefaultLangcode Get BlockContentFieldDatas via LangcodeAndRevisionCreatedAndDefaultLangcode
func GetBlockContentFieldDatasByLangcodeAndRevisionCreatedAndDefaultLangcode(offset int, limit int, Langcode_ string, RevisionCreated_ int, DefaultLangcode_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("langcode = ? and revision_created = ? and default_langcode = ?", Langcode_, RevisionCreated_, DefaultLangcode_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByLangcodeAndRevisionUserAndRevisionTranslationAffected Get BlockContentFieldDatas via LangcodeAndRevisionUserAndRevisionTranslationAffected
func GetBlockContentFieldDatasByLangcodeAndRevisionUserAndRevisionTranslationAffected(offset int, limit int, Langcode_ string, RevisionUser_ int, RevisionTranslationAffected_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("langcode = ? and revision_user = ? and revision_translation_affected = ?", Langcode_, RevisionUser_, RevisionTranslationAffected_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByLangcodeAndRevisionUserAndDefaultLangcode Get BlockContentFieldDatas via LangcodeAndRevisionUserAndDefaultLangcode
func GetBlockContentFieldDatasByLangcodeAndRevisionUserAndDefaultLangcode(offset int, limit int, Langcode_ string, RevisionUser_ int, DefaultLangcode_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("langcode = ? and revision_user = ? and default_langcode = ?", Langcode_, RevisionUser_, DefaultLangcode_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByLangcodeAndRevisionTranslationAffectedAndDefaultLangcode Get BlockContentFieldDatas via LangcodeAndRevisionTranslationAffectedAndDefaultLangcode
func GetBlockContentFieldDatasByLangcodeAndRevisionTranslationAffectedAndDefaultLangcode(offset int, limit int, Langcode_ string, RevisionTranslationAffected_ int, DefaultLangcode_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("langcode = ? and revision_translation_affected = ? and default_langcode = ?", Langcode_, RevisionTranslationAffected_, DefaultLangcode_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByInfoAndChangedAndRevisionCreated Get BlockContentFieldDatas via InfoAndChangedAndRevisionCreated
func GetBlockContentFieldDatasByInfoAndChangedAndRevisionCreated(offset int, limit int, Info_ string, Changed_ int, RevisionCreated_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("info = ? and changed = ? and revision_created = ?", Info_, Changed_, RevisionCreated_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByInfoAndChangedAndRevisionUser Get BlockContentFieldDatas via InfoAndChangedAndRevisionUser
func GetBlockContentFieldDatasByInfoAndChangedAndRevisionUser(offset int, limit int, Info_ string, Changed_ int, RevisionUser_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("info = ? and changed = ? and revision_user = ?", Info_, Changed_, RevisionUser_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByInfoAndChangedAndRevisionTranslationAffected Get BlockContentFieldDatas via InfoAndChangedAndRevisionTranslationAffected
func GetBlockContentFieldDatasByInfoAndChangedAndRevisionTranslationAffected(offset int, limit int, Info_ string, Changed_ int, RevisionTranslationAffected_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("info = ? and changed = ? and revision_translation_affected = ?", Info_, Changed_, RevisionTranslationAffected_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByInfoAndChangedAndDefaultLangcode Get BlockContentFieldDatas via InfoAndChangedAndDefaultLangcode
func GetBlockContentFieldDatasByInfoAndChangedAndDefaultLangcode(offset int, limit int, Info_ string, Changed_ int, DefaultLangcode_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("info = ? and changed = ? and default_langcode = ?", Info_, Changed_, DefaultLangcode_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByInfoAndRevisionCreatedAndRevisionUser Get BlockContentFieldDatas via InfoAndRevisionCreatedAndRevisionUser
func GetBlockContentFieldDatasByInfoAndRevisionCreatedAndRevisionUser(offset int, limit int, Info_ string, RevisionCreated_ int, RevisionUser_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("info = ? and revision_created = ? and revision_user = ?", Info_, RevisionCreated_, RevisionUser_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByInfoAndRevisionCreatedAndRevisionTranslationAffected Get BlockContentFieldDatas via InfoAndRevisionCreatedAndRevisionTranslationAffected
func GetBlockContentFieldDatasByInfoAndRevisionCreatedAndRevisionTranslationAffected(offset int, limit int, Info_ string, RevisionCreated_ int, RevisionTranslationAffected_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("info = ? and revision_created = ? and revision_translation_affected = ?", Info_, RevisionCreated_, RevisionTranslationAffected_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByInfoAndRevisionCreatedAndDefaultLangcode Get BlockContentFieldDatas via InfoAndRevisionCreatedAndDefaultLangcode
func GetBlockContentFieldDatasByInfoAndRevisionCreatedAndDefaultLangcode(offset int, limit int, Info_ string, RevisionCreated_ int, DefaultLangcode_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("info = ? and revision_created = ? and default_langcode = ?", Info_, RevisionCreated_, DefaultLangcode_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByInfoAndRevisionUserAndRevisionTranslationAffected Get BlockContentFieldDatas via InfoAndRevisionUserAndRevisionTranslationAffected
func GetBlockContentFieldDatasByInfoAndRevisionUserAndRevisionTranslationAffected(offset int, limit int, Info_ string, RevisionUser_ int, RevisionTranslationAffected_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("info = ? and revision_user = ? and revision_translation_affected = ?", Info_, RevisionUser_, RevisionTranslationAffected_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByInfoAndRevisionUserAndDefaultLangcode Get BlockContentFieldDatas via InfoAndRevisionUserAndDefaultLangcode
func GetBlockContentFieldDatasByInfoAndRevisionUserAndDefaultLangcode(offset int, limit int, Info_ string, RevisionUser_ int, DefaultLangcode_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("info = ? and revision_user = ? and default_langcode = ?", Info_, RevisionUser_, DefaultLangcode_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByInfoAndRevisionTranslationAffectedAndDefaultLangcode Get BlockContentFieldDatas via InfoAndRevisionTranslationAffectedAndDefaultLangcode
func GetBlockContentFieldDatasByInfoAndRevisionTranslationAffectedAndDefaultLangcode(offset int, limit int, Info_ string, RevisionTranslationAffected_ int, DefaultLangcode_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("info = ? and revision_translation_affected = ? and default_langcode = ?", Info_, RevisionTranslationAffected_, DefaultLangcode_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByChangedAndRevisionCreatedAndRevisionUser Get BlockContentFieldDatas via ChangedAndRevisionCreatedAndRevisionUser
func GetBlockContentFieldDatasByChangedAndRevisionCreatedAndRevisionUser(offset int, limit int, Changed_ int, RevisionCreated_ int, RevisionUser_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("changed = ? and revision_created = ? and revision_user = ?", Changed_, RevisionCreated_, RevisionUser_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByChangedAndRevisionCreatedAndRevisionTranslationAffected Get BlockContentFieldDatas via ChangedAndRevisionCreatedAndRevisionTranslationAffected
func GetBlockContentFieldDatasByChangedAndRevisionCreatedAndRevisionTranslationAffected(offset int, limit int, Changed_ int, RevisionCreated_ int, RevisionTranslationAffected_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("changed = ? and revision_created = ? and revision_translation_affected = ?", Changed_, RevisionCreated_, RevisionTranslationAffected_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByChangedAndRevisionCreatedAndDefaultLangcode Get BlockContentFieldDatas via ChangedAndRevisionCreatedAndDefaultLangcode
func GetBlockContentFieldDatasByChangedAndRevisionCreatedAndDefaultLangcode(offset int, limit int, Changed_ int, RevisionCreated_ int, DefaultLangcode_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("changed = ? and revision_created = ? and default_langcode = ?", Changed_, RevisionCreated_, DefaultLangcode_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByChangedAndRevisionUserAndRevisionTranslationAffected Get BlockContentFieldDatas via ChangedAndRevisionUserAndRevisionTranslationAffected
func GetBlockContentFieldDatasByChangedAndRevisionUserAndRevisionTranslationAffected(offset int, limit int, Changed_ int, RevisionUser_ int, RevisionTranslationAffected_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("changed = ? and revision_user = ? and revision_translation_affected = ?", Changed_, RevisionUser_, RevisionTranslationAffected_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByChangedAndRevisionUserAndDefaultLangcode Get BlockContentFieldDatas via ChangedAndRevisionUserAndDefaultLangcode
func GetBlockContentFieldDatasByChangedAndRevisionUserAndDefaultLangcode(offset int, limit int, Changed_ int, RevisionUser_ int, DefaultLangcode_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("changed = ? and revision_user = ? and default_langcode = ?", Changed_, RevisionUser_, DefaultLangcode_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByChangedAndRevisionTranslationAffectedAndDefaultLangcode Get BlockContentFieldDatas via ChangedAndRevisionTranslationAffectedAndDefaultLangcode
func GetBlockContentFieldDatasByChangedAndRevisionTranslationAffectedAndDefaultLangcode(offset int, limit int, Changed_ int, RevisionTranslationAffected_ int, DefaultLangcode_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("changed = ? and revision_translation_affected = ? and default_langcode = ?", Changed_, RevisionTranslationAffected_, DefaultLangcode_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByRevisionCreatedAndRevisionUserAndRevisionTranslationAffected Get BlockContentFieldDatas via RevisionCreatedAndRevisionUserAndRevisionTranslationAffected
func GetBlockContentFieldDatasByRevisionCreatedAndRevisionUserAndRevisionTranslationAffected(offset int, limit int, RevisionCreated_ int, RevisionUser_ int, RevisionTranslationAffected_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("revision_created = ? and revision_user = ? and revision_translation_affected = ?", RevisionCreated_, RevisionUser_, RevisionTranslationAffected_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByRevisionCreatedAndRevisionUserAndDefaultLangcode Get BlockContentFieldDatas via RevisionCreatedAndRevisionUserAndDefaultLangcode
func GetBlockContentFieldDatasByRevisionCreatedAndRevisionUserAndDefaultLangcode(offset int, limit int, RevisionCreated_ int, RevisionUser_ int, DefaultLangcode_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("revision_created = ? and revision_user = ? and default_langcode = ?", RevisionCreated_, RevisionUser_, DefaultLangcode_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByRevisionCreatedAndRevisionTranslationAffectedAndDefaultLangcode Get BlockContentFieldDatas via RevisionCreatedAndRevisionTranslationAffectedAndDefaultLangcode
func GetBlockContentFieldDatasByRevisionCreatedAndRevisionTranslationAffectedAndDefaultLangcode(offset int, limit int, RevisionCreated_ int, RevisionTranslationAffected_ int, DefaultLangcode_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("revision_created = ? and revision_translation_affected = ? and default_langcode = ?", RevisionCreated_, RevisionTranslationAffected_, DefaultLangcode_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByRevisionUserAndRevisionTranslationAffectedAndDefaultLangcode Get BlockContentFieldDatas via RevisionUserAndRevisionTranslationAffectedAndDefaultLangcode
func GetBlockContentFieldDatasByRevisionUserAndRevisionTranslationAffectedAndDefaultLangcode(offset int, limit int, RevisionUser_ int, RevisionTranslationAffected_ int, DefaultLangcode_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("revision_user = ? and revision_translation_affected = ? and default_langcode = ?", RevisionUser_, RevisionTranslationAffected_, DefaultLangcode_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByIdAndRevisionId Get BlockContentFieldDatas via IdAndRevisionId
func GetBlockContentFieldDatasByIdAndRevisionId(offset int, limit int, Id_ int64, RevisionId_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("id = ? and revision_id = ?", Id_, RevisionId_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByIdAndType Get BlockContentFieldDatas via IdAndType
func GetBlockContentFieldDatasByIdAndType(offset int, limit int, Id_ int64, Type_ string) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("id = ? and type = ?", Id_, Type_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByIdAndLangcode Get BlockContentFieldDatas via IdAndLangcode
func GetBlockContentFieldDatasByIdAndLangcode(offset int, limit int, Id_ int64, Langcode_ string) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("id = ? and langcode = ?", Id_, Langcode_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByIdAndInfo Get BlockContentFieldDatas via IdAndInfo
func GetBlockContentFieldDatasByIdAndInfo(offset int, limit int, Id_ int64, Info_ string) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("id = ? and info = ?", Id_, Info_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByIdAndChanged Get BlockContentFieldDatas via IdAndChanged
func GetBlockContentFieldDatasByIdAndChanged(offset int, limit int, Id_ int64, Changed_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("id = ? and changed = ?", Id_, Changed_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByIdAndRevisionCreated Get BlockContentFieldDatas via IdAndRevisionCreated
func GetBlockContentFieldDatasByIdAndRevisionCreated(offset int, limit int, Id_ int64, RevisionCreated_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("id = ? and revision_created = ?", Id_, RevisionCreated_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByIdAndRevisionUser Get BlockContentFieldDatas via IdAndRevisionUser
func GetBlockContentFieldDatasByIdAndRevisionUser(offset int, limit int, Id_ int64, RevisionUser_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("id = ? and revision_user = ?", Id_, RevisionUser_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByIdAndRevisionTranslationAffected Get BlockContentFieldDatas via IdAndRevisionTranslationAffected
func GetBlockContentFieldDatasByIdAndRevisionTranslationAffected(offset int, limit int, Id_ int64, RevisionTranslationAffected_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("id = ? and revision_translation_affected = ?", Id_, RevisionTranslationAffected_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByIdAndDefaultLangcode Get BlockContentFieldDatas via IdAndDefaultLangcode
func GetBlockContentFieldDatasByIdAndDefaultLangcode(offset int, limit int, Id_ int64, DefaultLangcode_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("id = ? and default_langcode = ?", Id_, DefaultLangcode_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByRevisionIdAndType Get BlockContentFieldDatas via RevisionIdAndType
func GetBlockContentFieldDatasByRevisionIdAndType(offset int, limit int, RevisionId_ int, Type_ string) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("revision_id = ? and type = ?", RevisionId_, Type_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByRevisionIdAndLangcode Get BlockContentFieldDatas via RevisionIdAndLangcode
func GetBlockContentFieldDatasByRevisionIdAndLangcode(offset int, limit int, RevisionId_ int, Langcode_ string) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("revision_id = ? and langcode = ?", RevisionId_, Langcode_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByRevisionIdAndInfo Get BlockContentFieldDatas via RevisionIdAndInfo
func GetBlockContentFieldDatasByRevisionIdAndInfo(offset int, limit int, RevisionId_ int, Info_ string) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("revision_id = ? and info = ?", RevisionId_, Info_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByRevisionIdAndChanged Get BlockContentFieldDatas via RevisionIdAndChanged
func GetBlockContentFieldDatasByRevisionIdAndChanged(offset int, limit int, RevisionId_ int, Changed_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("revision_id = ? and changed = ?", RevisionId_, Changed_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByRevisionIdAndRevisionCreated Get BlockContentFieldDatas via RevisionIdAndRevisionCreated
func GetBlockContentFieldDatasByRevisionIdAndRevisionCreated(offset int, limit int, RevisionId_ int, RevisionCreated_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("revision_id = ? and revision_created = ?", RevisionId_, RevisionCreated_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByRevisionIdAndRevisionUser Get BlockContentFieldDatas via RevisionIdAndRevisionUser
func GetBlockContentFieldDatasByRevisionIdAndRevisionUser(offset int, limit int, RevisionId_ int, RevisionUser_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("revision_id = ? and revision_user = ?", RevisionId_, RevisionUser_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByRevisionIdAndRevisionTranslationAffected Get BlockContentFieldDatas via RevisionIdAndRevisionTranslationAffected
func GetBlockContentFieldDatasByRevisionIdAndRevisionTranslationAffected(offset int, limit int, RevisionId_ int, RevisionTranslationAffected_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("revision_id = ? and revision_translation_affected = ?", RevisionId_, RevisionTranslationAffected_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByRevisionIdAndDefaultLangcode Get BlockContentFieldDatas via RevisionIdAndDefaultLangcode
func GetBlockContentFieldDatasByRevisionIdAndDefaultLangcode(offset int, limit int, RevisionId_ int, DefaultLangcode_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("revision_id = ? and default_langcode = ?", RevisionId_, DefaultLangcode_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByTypeAndLangcode Get BlockContentFieldDatas via TypeAndLangcode
func GetBlockContentFieldDatasByTypeAndLangcode(offset int, limit int, Type_ string, Langcode_ string) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("type = ? and langcode = ?", Type_, Langcode_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByTypeAndInfo Get BlockContentFieldDatas via TypeAndInfo
func GetBlockContentFieldDatasByTypeAndInfo(offset int, limit int, Type_ string, Info_ string) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("type = ? and info = ?", Type_, Info_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByTypeAndChanged Get BlockContentFieldDatas via TypeAndChanged
func GetBlockContentFieldDatasByTypeAndChanged(offset int, limit int, Type_ string, Changed_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("type = ? and changed = ?", Type_, Changed_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByTypeAndRevisionCreated Get BlockContentFieldDatas via TypeAndRevisionCreated
func GetBlockContentFieldDatasByTypeAndRevisionCreated(offset int, limit int, Type_ string, RevisionCreated_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("type = ? and revision_created = ?", Type_, RevisionCreated_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByTypeAndRevisionUser Get BlockContentFieldDatas via TypeAndRevisionUser
func GetBlockContentFieldDatasByTypeAndRevisionUser(offset int, limit int, Type_ string, RevisionUser_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("type = ? and revision_user = ?", Type_, RevisionUser_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByTypeAndRevisionTranslationAffected Get BlockContentFieldDatas via TypeAndRevisionTranslationAffected
func GetBlockContentFieldDatasByTypeAndRevisionTranslationAffected(offset int, limit int, Type_ string, RevisionTranslationAffected_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("type = ? and revision_translation_affected = ?", Type_, RevisionTranslationAffected_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByTypeAndDefaultLangcode Get BlockContentFieldDatas via TypeAndDefaultLangcode
func GetBlockContentFieldDatasByTypeAndDefaultLangcode(offset int, limit int, Type_ string, DefaultLangcode_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("type = ? and default_langcode = ?", Type_, DefaultLangcode_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByLangcodeAndInfo Get BlockContentFieldDatas via LangcodeAndInfo
func GetBlockContentFieldDatasByLangcodeAndInfo(offset int, limit int, Langcode_ string, Info_ string) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("langcode = ? and info = ?", Langcode_, Info_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByLangcodeAndChanged Get BlockContentFieldDatas via LangcodeAndChanged
func GetBlockContentFieldDatasByLangcodeAndChanged(offset int, limit int, Langcode_ string, Changed_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("langcode = ? and changed = ?", Langcode_, Changed_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByLangcodeAndRevisionCreated Get BlockContentFieldDatas via LangcodeAndRevisionCreated
func GetBlockContentFieldDatasByLangcodeAndRevisionCreated(offset int, limit int, Langcode_ string, RevisionCreated_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("langcode = ? and revision_created = ?", Langcode_, RevisionCreated_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByLangcodeAndRevisionUser Get BlockContentFieldDatas via LangcodeAndRevisionUser
func GetBlockContentFieldDatasByLangcodeAndRevisionUser(offset int, limit int, Langcode_ string, RevisionUser_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("langcode = ? and revision_user = ?", Langcode_, RevisionUser_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByLangcodeAndRevisionTranslationAffected Get BlockContentFieldDatas via LangcodeAndRevisionTranslationAffected
func GetBlockContentFieldDatasByLangcodeAndRevisionTranslationAffected(offset int, limit int, Langcode_ string, RevisionTranslationAffected_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("langcode = ? and revision_translation_affected = ?", Langcode_, RevisionTranslationAffected_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByLangcodeAndDefaultLangcode Get BlockContentFieldDatas via LangcodeAndDefaultLangcode
func GetBlockContentFieldDatasByLangcodeAndDefaultLangcode(offset int, limit int, Langcode_ string, DefaultLangcode_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("langcode = ? and default_langcode = ?", Langcode_, DefaultLangcode_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByInfoAndChanged Get BlockContentFieldDatas via InfoAndChanged
func GetBlockContentFieldDatasByInfoAndChanged(offset int, limit int, Info_ string, Changed_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("info = ? and changed = ?", Info_, Changed_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByInfoAndRevisionCreated Get BlockContentFieldDatas via InfoAndRevisionCreated
func GetBlockContentFieldDatasByInfoAndRevisionCreated(offset int, limit int, Info_ string, RevisionCreated_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("info = ? and revision_created = ?", Info_, RevisionCreated_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByInfoAndRevisionUser Get BlockContentFieldDatas via InfoAndRevisionUser
func GetBlockContentFieldDatasByInfoAndRevisionUser(offset int, limit int, Info_ string, RevisionUser_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("info = ? and revision_user = ?", Info_, RevisionUser_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByInfoAndRevisionTranslationAffected Get BlockContentFieldDatas via InfoAndRevisionTranslationAffected
func GetBlockContentFieldDatasByInfoAndRevisionTranslationAffected(offset int, limit int, Info_ string, RevisionTranslationAffected_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("info = ? and revision_translation_affected = ?", Info_, RevisionTranslationAffected_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByInfoAndDefaultLangcode Get BlockContentFieldDatas via InfoAndDefaultLangcode
func GetBlockContentFieldDatasByInfoAndDefaultLangcode(offset int, limit int, Info_ string, DefaultLangcode_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("info = ? and default_langcode = ?", Info_, DefaultLangcode_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByChangedAndRevisionCreated Get BlockContentFieldDatas via ChangedAndRevisionCreated
func GetBlockContentFieldDatasByChangedAndRevisionCreated(offset int, limit int, Changed_ int, RevisionCreated_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("changed = ? and revision_created = ?", Changed_, RevisionCreated_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByChangedAndRevisionUser Get BlockContentFieldDatas via ChangedAndRevisionUser
func GetBlockContentFieldDatasByChangedAndRevisionUser(offset int, limit int, Changed_ int, RevisionUser_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("changed = ? and revision_user = ?", Changed_, RevisionUser_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByChangedAndRevisionTranslationAffected Get BlockContentFieldDatas via ChangedAndRevisionTranslationAffected
func GetBlockContentFieldDatasByChangedAndRevisionTranslationAffected(offset int, limit int, Changed_ int, RevisionTranslationAffected_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("changed = ? and revision_translation_affected = ?", Changed_, RevisionTranslationAffected_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByChangedAndDefaultLangcode Get BlockContentFieldDatas via ChangedAndDefaultLangcode
func GetBlockContentFieldDatasByChangedAndDefaultLangcode(offset int, limit int, Changed_ int, DefaultLangcode_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("changed = ? and default_langcode = ?", Changed_, DefaultLangcode_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByRevisionCreatedAndRevisionUser Get BlockContentFieldDatas via RevisionCreatedAndRevisionUser
func GetBlockContentFieldDatasByRevisionCreatedAndRevisionUser(offset int, limit int, RevisionCreated_ int, RevisionUser_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("revision_created = ? and revision_user = ?", RevisionCreated_, RevisionUser_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByRevisionCreatedAndRevisionTranslationAffected Get BlockContentFieldDatas via RevisionCreatedAndRevisionTranslationAffected
func GetBlockContentFieldDatasByRevisionCreatedAndRevisionTranslationAffected(offset int, limit int, RevisionCreated_ int, RevisionTranslationAffected_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("revision_created = ? and revision_translation_affected = ?", RevisionCreated_, RevisionTranslationAffected_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByRevisionCreatedAndDefaultLangcode Get BlockContentFieldDatas via RevisionCreatedAndDefaultLangcode
func GetBlockContentFieldDatasByRevisionCreatedAndDefaultLangcode(offset int, limit int, RevisionCreated_ int, DefaultLangcode_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("revision_created = ? and default_langcode = ?", RevisionCreated_, DefaultLangcode_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByRevisionUserAndRevisionTranslationAffected Get BlockContentFieldDatas via RevisionUserAndRevisionTranslationAffected
func GetBlockContentFieldDatasByRevisionUserAndRevisionTranslationAffected(offset int, limit int, RevisionUser_ int, RevisionTranslationAffected_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("revision_user = ? and revision_translation_affected = ?", RevisionUser_, RevisionTranslationAffected_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByRevisionUserAndDefaultLangcode Get BlockContentFieldDatas via RevisionUserAndDefaultLangcode
func GetBlockContentFieldDatasByRevisionUserAndDefaultLangcode(offset int, limit int, RevisionUser_ int, DefaultLangcode_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("revision_user = ? and default_langcode = ?", RevisionUser_, DefaultLangcode_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasByRevisionTranslationAffectedAndDefaultLangcode Get BlockContentFieldDatas via RevisionTranslationAffectedAndDefaultLangcode
func GetBlockContentFieldDatasByRevisionTranslationAffectedAndDefaultLangcode(offset int, limit int, RevisionTranslationAffected_ int, DefaultLangcode_ int) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("revision_translation_affected = ? and default_langcode = ?", RevisionTranslationAffected_, DefaultLangcode_).Limit(limit, offset).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatas Get BlockContentFieldDatas via field
func GetBlockContentFieldDatas(offset int, limit int, field string) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Limit(limit, offset).Desc(field).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasViaId Get BlockContentFieldDatas via Id
func GetBlockContentFieldDatasViaId(offset int, limit int, Id_ int64, field string) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("id = ?", Id_).Limit(limit, offset).Desc(field).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasViaRevisionId Get BlockContentFieldDatas via RevisionId
func GetBlockContentFieldDatasViaRevisionId(offset int, limit int, RevisionId_ int, field string) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("revision_id = ?", RevisionId_).Limit(limit, offset).Desc(field).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasViaType Get BlockContentFieldDatas via Type
func GetBlockContentFieldDatasViaType(offset int, limit int, Type_ string, field string) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("type = ?", Type_).Limit(limit, offset).Desc(field).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasViaLangcode Get BlockContentFieldDatas via Langcode
func GetBlockContentFieldDatasViaLangcode(offset int, limit int, Langcode_ string, field string) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("langcode = ?", Langcode_).Limit(limit, offset).Desc(field).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasViaInfo Get BlockContentFieldDatas via Info
func GetBlockContentFieldDatasViaInfo(offset int, limit int, Info_ string, field string) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("info = ?", Info_).Limit(limit, offset).Desc(field).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasViaChanged Get BlockContentFieldDatas via Changed
func GetBlockContentFieldDatasViaChanged(offset int, limit int, Changed_ int, field string) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("changed = ?", Changed_).Limit(limit, offset).Desc(field).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasViaRevisionCreated Get BlockContentFieldDatas via RevisionCreated
func GetBlockContentFieldDatasViaRevisionCreated(offset int, limit int, RevisionCreated_ int, field string) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("revision_created = ?", RevisionCreated_).Limit(limit, offset).Desc(field).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasViaRevisionUser Get BlockContentFieldDatas via RevisionUser
func GetBlockContentFieldDatasViaRevisionUser(offset int, limit int, RevisionUser_ int, field string) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("revision_user = ?", RevisionUser_).Limit(limit, offset).Desc(field).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasViaRevisionTranslationAffected Get BlockContentFieldDatas via RevisionTranslationAffected
func GetBlockContentFieldDatasViaRevisionTranslationAffected(offset int, limit int, RevisionTranslationAffected_ int, field string) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("revision_translation_affected = ?", RevisionTranslationAffected_).Limit(limit, offset).Desc(field).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// GetBlockContentFieldDatasViaDefaultLangcode Get BlockContentFieldDatas via DefaultLangcode
func GetBlockContentFieldDatasViaDefaultLangcode(offset int, limit int, DefaultLangcode_ int, field string) (*[]*BlockContentFieldData, error) {
	var _BlockContentFieldData = new([]*BlockContentFieldData)
	err := Engine.Table("block_content_field_data").Where("default_langcode = ?", DefaultLangcode_).Limit(limit, offset).Desc(field).Find(_BlockContentFieldData)
	return _BlockContentFieldData, err
}

// HasBlockContentFieldDataViaId Has BlockContentFieldData via Id
func HasBlockContentFieldDataViaId(iId int64) bool {
	if has, err := Engine.Where("id = ?", iId).Get(new(BlockContentFieldData)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasBlockContentFieldDataViaRevisionId Has BlockContentFieldData via RevisionId
func HasBlockContentFieldDataViaRevisionId(iRevisionId int) bool {
	if has, err := Engine.Where("revision_id = ?", iRevisionId).Get(new(BlockContentFieldData)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasBlockContentFieldDataViaType Has BlockContentFieldData via Type
func HasBlockContentFieldDataViaType(iType string) bool {
	if has, err := Engine.Where("type = ?", iType).Get(new(BlockContentFieldData)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasBlockContentFieldDataViaLangcode Has BlockContentFieldData via Langcode
func HasBlockContentFieldDataViaLangcode(iLangcode string) bool {
	if has, err := Engine.Where("langcode = ?", iLangcode).Get(new(BlockContentFieldData)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasBlockContentFieldDataViaInfo Has BlockContentFieldData via Info
func HasBlockContentFieldDataViaInfo(iInfo string) bool {
	if has, err := Engine.Where("info = ?", iInfo).Get(new(BlockContentFieldData)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasBlockContentFieldDataViaChanged Has BlockContentFieldData via Changed
func HasBlockContentFieldDataViaChanged(iChanged int) bool {
	if has, err := Engine.Where("changed = ?", iChanged).Get(new(BlockContentFieldData)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasBlockContentFieldDataViaRevisionCreated Has BlockContentFieldData via RevisionCreated
func HasBlockContentFieldDataViaRevisionCreated(iRevisionCreated int) bool {
	if has, err := Engine.Where("revision_created = ?", iRevisionCreated).Get(new(BlockContentFieldData)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasBlockContentFieldDataViaRevisionUser Has BlockContentFieldData via RevisionUser
func HasBlockContentFieldDataViaRevisionUser(iRevisionUser int) bool {
	if has, err := Engine.Where("revision_user = ?", iRevisionUser).Get(new(BlockContentFieldData)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasBlockContentFieldDataViaRevisionTranslationAffected Has BlockContentFieldData via RevisionTranslationAffected
func HasBlockContentFieldDataViaRevisionTranslationAffected(iRevisionTranslationAffected int) bool {
	if has, err := Engine.Where("revision_translation_affected = ?", iRevisionTranslationAffected).Get(new(BlockContentFieldData)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasBlockContentFieldDataViaDefaultLangcode Has BlockContentFieldData via DefaultLangcode
func HasBlockContentFieldDataViaDefaultLangcode(iDefaultLangcode int) bool {
	if has, err := Engine.Where("default_langcode = ?", iDefaultLangcode).Get(new(BlockContentFieldData)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// GetBlockContentFieldDataViaId Get BlockContentFieldData via Id
func GetBlockContentFieldDataViaId(iId int64) (*BlockContentFieldData, error) {
	var _BlockContentFieldData = &BlockContentFieldData{Id: iId}
	has, err := Engine.Get(_BlockContentFieldData)
	if has {
		return _BlockContentFieldData, err
	} else {
		return nil, err
	}
}

// GetBlockContentFieldDataViaRevisionId Get BlockContentFieldData via RevisionId
func GetBlockContentFieldDataViaRevisionId(iRevisionId int) (*BlockContentFieldData, error) {
	var _BlockContentFieldData = &BlockContentFieldData{RevisionId: iRevisionId}
	has, err := Engine.Get(_BlockContentFieldData)
	if has {
		return _BlockContentFieldData, err
	} else {
		return nil, err
	}
}

// GetBlockContentFieldDataViaType Get BlockContentFieldData via Type
func GetBlockContentFieldDataViaType(iType string) (*BlockContentFieldData, error) {
	var _BlockContentFieldData = &BlockContentFieldData{Type: iType}
	has, err := Engine.Get(_BlockContentFieldData)
	if has {
		return _BlockContentFieldData, err
	} else {
		return nil, err
	}
}

// GetBlockContentFieldDataViaLangcode Get BlockContentFieldData via Langcode
func GetBlockContentFieldDataViaLangcode(iLangcode string) (*BlockContentFieldData, error) {
	var _BlockContentFieldData = &BlockContentFieldData{Langcode: iLangcode}
	has, err := Engine.Get(_BlockContentFieldData)
	if has {
		return _BlockContentFieldData, err
	} else {
		return nil, err
	}
}

// GetBlockContentFieldDataViaInfo Get BlockContentFieldData via Info
func GetBlockContentFieldDataViaInfo(iInfo string) (*BlockContentFieldData, error) {
	var _BlockContentFieldData = &BlockContentFieldData{Info: iInfo}
	has, err := Engine.Get(_BlockContentFieldData)
	if has {
		return _BlockContentFieldData, err
	} else {
		return nil, err
	}
}

// GetBlockContentFieldDataViaChanged Get BlockContentFieldData via Changed
func GetBlockContentFieldDataViaChanged(iChanged int) (*BlockContentFieldData, error) {
	var _BlockContentFieldData = &BlockContentFieldData{Changed: iChanged}
	has, err := Engine.Get(_BlockContentFieldData)
	if has {
		return _BlockContentFieldData, err
	} else {
		return nil, err
	}
}

// GetBlockContentFieldDataViaRevisionCreated Get BlockContentFieldData via RevisionCreated
func GetBlockContentFieldDataViaRevisionCreated(iRevisionCreated int) (*BlockContentFieldData, error) {
	var _BlockContentFieldData = &BlockContentFieldData{RevisionCreated: iRevisionCreated}
	has, err := Engine.Get(_BlockContentFieldData)
	if has {
		return _BlockContentFieldData, err
	} else {
		return nil, err
	}
}

// GetBlockContentFieldDataViaRevisionUser Get BlockContentFieldData via RevisionUser
func GetBlockContentFieldDataViaRevisionUser(iRevisionUser int) (*BlockContentFieldData, error) {
	var _BlockContentFieldData = &BlockContentFieldData{RevisionUser: iRevisionUser}
	has, err := Engine.Get(_BlockContentFieldData)
	if has {
		return _BlockContentFieldData, err
	} else {
		return nil, err
	}
}

// GetBlockContentFieldDataViaRevisionTranslationAffected Get BlockContentFieldData via RevisionTranslationAffected
func GetBlockContentFieldDataViaRevisionTranslationAffected(iRevisionTranslationAffected int) (*BlockContentFieldData, error) {
	var _BlockContentFieldData = &BlockContentFieldData{RevisionTranslationAffected: iRevisionTranslationAffected}
	has, err := Engine.Get(_BlockContentFieldData)
	if has {
		return _BlockContentFieldData, err
	} else {
		return nil, err
	}
}

// GetBlockContentFieldDataViaDefaultLangcode Get BlockContentFieldData via DefaultLangcode
func GetBlockContentFieldDataViaDefaultLangcode(iDefaultLangcode int) (*BlockContentFieldData, error) {
	var _BlockContentFieldData = &BlockContentFieldData{DefaultLangcode: iDefaultLangcode}
	has, err := Engine.Get(_BlockContentFieldData)
	if has {
		return _BlockContentFieldData, err
	} else {
		return nil, err
	}
}

// SetBlockContentFieldDataViaId Set BlockContentFieldData via Id
func SetBlockContentFieldDataViaId(iId int64, block_content_field_data *BlockContentFieldData) (int64, error) {
	block_content_field_data.Id = iId
	return Engine.Insert(block_content_field_data)
}

// SetBlockContentFieldDataViaRevisionId Set BlockContentFieldData via RevisionId
func SetBlockContentFieldDataViaRevisionId(iRevisionId int, block_content_field_data *BlockContentFieldData) (int64, error) {
	block_content_field_data.RevisionId = iRevisionId
	return Engine.Insert(block_content_field_data)
}

// SetBlockContentFieldDataViaType Set BlockContentFieldData via Type
func SetBlockContentFieldDataViaType(iType string, block_content_field_data *BlockContentFieldData) (int64, error) {
	block_content_field_data.Type = iType
	return Engine.Insert(block_content_field_data)
}

// SetBlockContentFieldDataViaLangcode Set BlockContentFieldData via Langcode
func SetBlockContentFieldDataViaLangcode(iLangcode string, block_content_field_data *BlockContentFieldData) (int64, error) {
	block_content_field_data.Langcode = iLangcode
	return Engine.Insert(block_content_field_data)
}

// SetBlockContentFieldDataViaInfo Set BlockContentFieldData via Info
func SetBlockContentFieldDataViaInfo(iInfo string, block_content_field_data *BlockContentFieldData) (int64, error) {
	block_content_field_data.Info = iInfo
	return Engine.Insert(block_content_field_data)
}

// SetBlockContentFieldDataViaChanged Set BlockContentFieldData via Changed
func SetBlockContentFieldDataViaChanged(iChanged int, block_content_field_data *BlockContentFieldData) (int64, error) {
	block_content_field_data.Changed = iChanged
	return Engine.Insert(block_content_field_data)
}

// SetBlockContentFieldDataViaRevisionCreated Set BlockContentFieldData via RevisionCreated
func SetBlockContentFieldDataViaRevisionCreated(iRevisionCreated int, block_content_field_data *BlockContentFieldData) (int64, error) {
	block_content_field_data.RevisionCreated = iRevisionCreated
	return Engine.Insert(block_content_field_data)
}

// SetBlockContentFieldDataViaRevisionUser Set BlockContentFieldData via RevisionUser
func SetBlockContentFieldDataViaRevisionUser(iRevisionUser int, block_content_field_data *BlockContentFieldData) (int64, error) {
	block_content_field_data.RevisionUser = iRevisionUser
	return Engine.Insert(block_content_field_data)
}

// SetBlockContentFieldDataViaRevisionTranslationAffected Set BlockContentFieldData via RevisionTranslationAffected
func SetBlockContentFieldDataViaRevisionTranslationAffected(iRevisionTranslationAffected int, block_content_field_data *BlockContentFieldData) (int64, error) {
	block_content_field_data.RevisionTranslationAffected = iRevisionTranslationAffected
	return Engine.Insert(block_content_field_data)
}

// SetBlockContentFieldDataViaDefaultLangcode Set BlockContentFieldData via DefaultLangcode
func SetBlockContentFieldDataViaDefaultLangcode(iDefaultLangcode int, block_content_field_data *BlockContentFieldData) (int64, error) {
	block_content_field_data.DefaultLangcode = iDefaultLangcode
	return Engine.Insert(block_content_field_data)
}

// AddBlockContentFieldData Add BlockContentFieldData via all columns
func AddBlockContentFieldData(iId int64, iRevisionId int, iType string, iLangcode string, iInfo string, iChanged int, iRevisionCreated int, iRevisionUser int, iRevisionTranslationAffected int, iDefaultLangcode int) error {
	_BlockContentFieldData := &BlockContentFieldData{Id: iId, RevisionId: iRevisionId, Type: iType, Langcode: iLangcode, Info: iInfo, Changed: iChanged, RevisionCreated: iRevisionCreated, RevisionUser: iRevisionUser, RevisionTranslationAffected: iRevisionTranslationAffected, DefaultLangcode: iDefaultLangcode}
	if _, err := Engine.Insert(_BlockContentFieldData); err != nil {
		return err
	} else {
		return nil
	}
}

// PostBlockContentFieldData Post BlockContentFieldData via iBlockContentFieldData
func PostBlockContentFieldData(iBlockContentFieldData *BlockContentFieldData) (int64, error) {
	_, err := Engine.Insert(iBlockContentFieldData)
	return iBlockContentFieldData.Id, err
}

// PutBlockContentFieldData Put BlockContentFieldData
func PutBlockContentFieldData(iBlockContentFieldData *BlockContentFieldData) (int64, error) {
	_, err := Engine.Id(iBlockContentFieldData.Id).Update(iBlockContentFieldData)
	return iBlockContentFieldData.Id, err
}

// PutBlockContentFieldDataViaId Put BlockContentFieldData via Id
func PutBlockContentFieldDataViaId(Id_ int64, iBlockContentFieldData *BlockContentFieldData) (int64, error) {
	row, err := Engine.Update(iBlockContentFieldData, &BlockContentFieldData{Id: Id_})
	return row, err
}

// PutBlockContentFieldDataViaRevisionId Put BlockContentFieldData via RevisionId
func PutBlockContentFieldDataViaRevisionId(RevisionId_ int, iBlockContentFieldData *BlockContentFieldData) (int64, error) {
	row, err := Engine.Update(iBlockContentFieldData, &BlockContentFieldData{RevisionId: RevisionId_})
	return row, err
}

// PutBlockContentFieldDataViaType Put BlockContentFieldData via Type
func PutBlockContentFieldDataViaType(Type_ string, iBlockContentFieldData *BlockContentFieldData) (int64, error) {
	row, err := Engine.Update(iBlockContentFieldData, &BlockContentFieldData{Type: Type_})
	return row, err
}

// PutBlockContentFieldDataViaLangcode Put BlockContentFieldData via Langcode
func PutBlockContentFieldDataViaLangcode(Langcode_ string, iBlockContentFieldData *BlockContentFieldData) (int64, error) {
	row, err := Engine.Update(iBlockContentFieldData, &BlockContentFieldData{Langcode: Langcode_})
	return row, err
}

// PutBlockContentFieldDataViaInfo Put BlockContentFieldData via Info
func PutBlockContentFieldDataViaInfo(Info_ string, iBlockContentFieldData *BlockContentFieldData) (int64, error) {
	row, err := Engine.Update(iBlockContentFieldData, &BlockContentFieldData{Info: Info_})
	return row, err
}

// PutBlockContentFieldDataViaChanged Put BlockContentFieldData via Changed
func PutBlockContentFieldDataViaChanged(Changed_ int, iBlockContentFieldData *BlockContentFieldData) (int64, error) {
	row, err := Engine.Update(iBlockContentFieldData, &BlockContentFieldData{Changed: Changed_})
	return row, err
}

// PutBlockContentFieldDataViaRevisionCreated Put BlockContentFieldData via RevisionCreated
func PutBlockContentFieldDataViaRevisionCreated(RevisionCreated_ int, iBlockContentFieldData *BlockContentFieldData) (int64, error) {
	row, err := Engine.Update(iBlockContentFieldData, &BlockContentFieldData{RevisionCreated: RevisionCreated_})
	return row, err
}

// PutBlockContentFieldDataViaRevisionUser Put BlockContentFieldData via RevisionUser
func PutBlockContentFieldDataViaRevisionUser(RevisionUser_ int, iBlockContentFieldData *BlockContentFieldData) (int64, error) {
	row, err := Engine.Update(iBlockContentFieldData, &BlockContentFieldData{RevisionUser: RevisionUser_})
	return row, err
}

// PutBlockContentFieldDataViaRevisionTranslationAffected Put BlockContentFieldData via RevisionTranslationAffected
func PutBlockContentFieldDataViaRevisionTranslationAffected(RevisionTranslationAffected_ int, iBlockContentFieldData *BlockContentFieldData) (int64, error) {
	row, err := Engine.Update(iBlockContentFieldData, &BlockContentFieldData{RevisionTranslationAffected: RevisionTranslationAffected_})
	return row, err
}

// PutBlockContentFieldDataViaDefaultLangcode Put BlockContentFieldData via DefaultLangcode
func PutBlockContentFieldDataViaDefaultLangcode(DefaultLangcode_ int, iBlockContentFieldData *BlockContentFieldData) (int64, error) {
	row, err := Engine.Update(iBlockContentFieldData, &BlockContentFieldData{DefaultLangcode: DefaultLangcode_})
	return row, err
}

// UpdateBlockContentFieldDataViaId via map[string]interface{}{}
func UpdateBlockContentFieldDataViaId(iId int64, iBlockContentFieldDataMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(BlockContentFieldData)).Where("id = ?", iId).Update(iBlockContentFieldDataMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateBlockContentFieldDataViaRevisionId via map[string]interface{}{}
func UpdateBlockContentFieldDataViaRevisionId(iRevisionId int, iBlockContentFieldDataMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(BlockContentFieldData)).Where("revision_id = ?", iRevisionId).Update(iBlockContentFieldDataMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateBlockContentFieldDataViaType via map[string]interface{}{}
func UpdateBlockContentFieldDataViaType(iType string, iBlockContentFieldDataMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(BlockContentFieldData)).Where("type = ?", iType).Update(iBlockContentFieldDataMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateBlockContentFieldDataViaLangcode via map[string]interface{}{}
func UpdateBlockContentFieldDataViaLangcode(iLangcode string, iBlockContentFieldDataMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(BlockContentFieldData)).Where("langcode = ?", iLangcode).Update(iBlockContentFieldDataMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateBlockContentFieldDataViaInfo via map[string]interface{}{}
func UpdateBlockContentFieldDataViaInfo(iInfo string, iBlockContentFieldDataMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(BlockContentFieldData)).Where("info = ?", iInfo).Update(iBlockContentFieldDataMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateBlockContentFieldDataViaChanged via map[string]interface{}{}
func UpdateBlockContentFieldDataViaChanged(iChanged int, iBlockContentFieldDataMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(BlockContentFieldData)).Where("changed = ?", iChanged).Update(iBlockContentFieldDataMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateBlockContentFieldDataViaRevisionCreated via map[string]interface{}{}
func UpdateBlockContentFieldDataViaRevisionCreated(iRevisionCreated int, iBlockContentFieldDataMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(BlockContentFieldData)).Where("revision_created = ?", iRevisionCreated).Update(iBlockContentFieldDataMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateBlockContentFieldDataViaRevisionUser via map[string]interface{}{}
func UpdateBlockContentFieldDataViaRevisionUser(iRevisionUser int, iBlockContentFieldDataMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(BlockContentFieldData)).Where("revision_user = ?", iRevisionUser).Update(iBlockContentFieldDataMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateBlockContentFieldDataViaRevisionTranslationAffected via map[string]interface{}{}
func UpdateBlockContentFieldDataViaRevisionTranslationAffected(iRevisionTranslationAffected int, iBlockContentFieldDataMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(BlockContentFieldData)).Where("revision_translation_affected = ?", iRevisionTranslationAffected).Update(iBlockContentFieldDataMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateBlockContentFieldDataViaDefaultLangcode via map[string]interface{}{}
func UpdateBlockContentFieldDataViaDefaultLangcode(iDefaultLangcode int, iBlockContentFieldDataMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(BlockContentFieldData)).Where("default_langcode = ?", iDefaultLangcode).Update(iBlockContentFieldDataMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// DeleteBlockContentFieldData Delete BlockContentFieldData
func DeleteBlockContentFieldData(iId int64) (int64, error) {
	row, err := Engine.Id(iId).Delete(new(BlockContentFieldData))
	return row, err
}

// DeleteBlockContentFieldDataViaId Delete BlockContentFieldData via Id
func DeleteBlockContentFieldDataViaId(iId int64) (err error) {
	var has bool
	var _BlockContentFieldData = &BlockContentFieldData{Id: iId}
	if has, err = Engine.Get(_BlockContentFieldData); (has == true) && (err == nil) {
		if row, err := Engine.Where("id = ?", iId).Delete(new(BlockContentFieldData)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteBlockContentFieldDataViaRevisionId Delete BlockContentFieldData via RevisionId
func DeleteBlockContentFieldDataViaRevisionId(iRevisionId int) (err error) {
	var has bool
	var _BlockContentFieldData = &BlockContentFieldData{RevisionId: iRevisionId}
	if has, err = Engine.Get(_BlockContentFieldData); (has == true) && (err == nil) {
		if row, err := Engine.Where("revision_id = ?", iRevisionId).Delete(new(BlockContentFieldData)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteBlockContentFieldDataViaType Delete BlockContentFieldData via Type
func DeleteBlockContentFieldDataViaType(iType string) (err error) {
	var has bool
	var _BlockContentFieldData = &BlockContentFieldData{Type: iType}
	if has, err = Engine.Get(_BlockContentFieldData); (has == true) && (err == nil) {
		if row, err := Engine.Where("type = ?", iType).Delete(new(BlockContentFieldData)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteBlockContentFieldDataViaLangcode Delete BlockContentFieldData via Langcode
func DeleteBlockContentFieldDataViaLangcode(iLangcode string) (err error) {
	var has bool
	var _BlockContentFieldData = &BlockContentFieldData{Langcode: iLangcode}
	if has, err = Engine.Get(_BlockContentFieldData); (has == true) && (err == nil) {
		if row, err := Engine.Where("langcode = ?", iLangcode).Delete(new(BlockContentFieldData)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteBlockContentFieldDataViaInfo Delete BlockContentFieldData via Info
func DeleteBlockContentFieldDataViaInfo(iInfo string) (err error) {
	var has bool
	var _BlockContentFieldData = &BlockContentFieldData{Info: iInfo}
	if has, err = Engine.Get(_BlockContentFieldData); (has == true) && (err == nil) {
		if row, err := Engine.Where("info = ?", iInfo).Delete(new(BlockContentFieldData)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteBlockContentFieldDataViaChanged Delete BlockContentFieldData via Changed
func DeleteBlockContentFieldDataViaChanged(iChanged int) (err error) {
	var has bool
	var _BlockContentFieldData = &BlockContentFieldData{Changed: iChanged}
	if has, err = Engine.Get(_BlockContentFieldData); (has == true) && (err == nil) {
		if row, err := Engine.Where("changed = ?", iChanged).Delete(new(BlockContentFieldData)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteBlockContentFieldDataViaRevisionCreated Delete BlockContentFieldData via RevisionCreated
func DeleteBlockContentFieldDataViaRevisionCreated(iRevisionCreated int) (err error) {
	var has bool
	var _BlockContentFieldData = &BlockContentFieldData{RevisionCreated: iRevisionCreated}
	if has, err = Engine.Get(_BlockContentFieldData); (has == true) && (err == nil) {
		if row, err := Engine.Where("revision_created = ?", iRevisionCreated).Delete(new(BlockContentFieldData)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteBlockContentFieldDataViaRevisionUser Delete BlockContentFieldData via RevisionUser
func DeleteBlockContentFieldDataViaRevisionUser(iRevisionUser int) (err error) {
	var has bool
	var _BlockContentFieldData = &BlockContentFieldData{RevisionUser: iRevisionUser}
	if has, err = Engine.Get(_BlockContentFieldData); (has == true) && (err == nil) {
		if row, err := Engine.Where("revision_user = ?", iRevisionUser).Delete(new(BlockContentFieldData)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteBlockContentFieldDataViaRevisionTranslationAffected Delete BlockContentFieldData via RevisionTranslationAffected
func DeleteBlockContentFieldDataViaRevisionTranslationAffected(iRevisionTranslationAffected int) (err error) {
	var has bool
	var _BlockContentFieldData = &BlockContentFieldData{RevisionTranslationAffected: iRevisionTranslationAffected}
	if has, err = Engine.Get(_BlockContentFieldData); (has == true) && (err == nil) {
		if row, err := Engine.Where("revision_translation_affected = ?", iRevisionTranslationAffected).Delete(new(BlockContentFieldData)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteBlockContentFieldDataViaDefaultLangcode Delete BlockContentFieldData via DefaultLangcode
func DeleteBlockContentFieldDataViaDefaultLangcode(iDefaultLangcode int) (err error) {
	var has bool
	var _BlockContentFieldData = &BlockContentFieldData{DefaultLangcode: iDefaultLangcode}
	if has, err = Engine.Get(_BlockContentFieldData); (has == true) && (err == nil) {
		if row, err := Engine.Where("default_langcode = ?", iDefaultLangcode).Delete(new(BlockContentFieldData)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}
