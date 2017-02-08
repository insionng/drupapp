package model

type CommentFieldData struct {
	Cid             int64  `xorm:"not null pk index(comment__num_new) index(comment__id__default_langcode__langcode) INT(10)"`
	CommentType     string `xorm:"not null index(comment__num_new) index(comment__entity_langcode) index VARCHAR(32)"`
	Langcode        string `xorm:"not null pk index(comment__id__default_langcode__langcode) VARCHAR(12)"`
	Pid             int    `xorm:"index(comment__status_pid) INT(10)"`
	EntityId        int    `xorm:"index(comment__entity_langcode) index(comment__num_new) INT(10)"`
	Subject         string `xorm:"VARCHAR(64)"`
	Uid             int    `xorm:"index INT(10)"`
	Name            string `xorm:"VARCHAR(60)"`
	Mail            string `xorm:"VARCHAR(254)"`
	Homepage        string `xorm:"VARCHAR(255)"`
	Hostname        string `xorm:"VARCHAR(128)"`
	Created         int    `xorm:"not null index index(comment__num_new) INT(11)"`
	Changed         int    `xorm:"INT(11)"`
	Status          int    `xorm:"index(comment__status_pid) index(comment__num_new) TINYINT(4)"`
	Thread          string `xorm:"not null index(comment__num_new) VARCHAR(255)"`
	EntityType      string `xorm:"index(comment__num_new) index(comment__entity_langcode) VARCHAR(32)"`
	FieldName       string `xorm:"VARCHAR(32)"`
	DefaultLangcode int    `xorm:"not null index(comment__entity_langcode) index(comment__id__default_langcode__langcode) TINYINT(4)"`
}

// GetCommentFieldDatasCount CommentFieldDatas' Count
func GetCommentFieldDatasCount(offset int, limit int) (int64, error) {
	total, err := Engine.Limit(limit, offset).Count(&CommentFieldData{})
	return total, err
}

// GetCommentFieldDataCountViaCid Get CommentFieldData via Cid
func GetCommentFieldDataCountViaCid(iCid int64) int64 {
	n, _ := Engine.Where("cid = ?", iCid).Count(&CommentFieldData{Cid: iCid})
	return n
}

// GetCommentFieldDataCountViaCommentType Get CommentFieldData via CommentType
func GetCommentFieldDataCountViaCommentType(iCommentType string) int64 {
	n, _ := Engine.Where("comment_type = ?", iCommentType).Count(&CommentFieldData{CommentType: iCommentType})
	return n
}

// GetCommentFieldDataCountViaLangcode Get CommentFieldData via Langcode
func GetCommentFieldDataCountViaLangcode(iLangcode string) int64 {
	n, _ := Engine.Where("langcode = ?", iLangcode).Count(&CommentFieldData{Langcode: iLangcode})
	return n
}

// GetCommentFieldDataCountViaPid Get CommentFieldData via Pid
func GetCommentFieldDataCountViaPid(iPid int) int64 {
	n, _ := Engine.Where("pid = ?", iPid).Count(&CommentFieldData{Pid: iPid})
	return n
}

// GetCommentFieldDataCountViaEntityId Get CommentFieldData via EntityId
func GetCommentFieldDataCountViaEntityId(iEntityId int) int64 {
	n, _ := Engine.Where("entity_id = ?", iEntityId).Count(&CommentFieldData{EntityId: iEntityId})
	return n
}

// GetCommentFieldDataCountViaSubject Get CommentFieldData via Subject
func GetCommentFieldDataCountViaSubject(iSubject string) int64 {
	n, _ := Engine.Where("subject = ?", iSubject).Count(&CommentFieldData{Subject: iSubject})
	return n
}

// GetCommentFieldDataCountViaUid Get CommentFieldData via Uid
func GetCommentFieldDataCountViaUid(iUid int) int64 {
	n, _ := Engine.Where("uid = ?", iUid).Count(&CommentFieldData{Uid: iUid})
	return n
}

// GetCommentFieldDataCountViaName Get CommentFieldData via Name
func GetCommentFieldDataCountViaName(iName string) int64 {
	n, _ := Engine.Where("name = ?", iName).Count(&CommentFieldData{Name: iName})
	return n
}

// GetCommentFieldDataCountViaMail Get CommentFieldData via Mail
func GetCommentFieldDataCountViaMail(iMail string) int64 {
	n, _ := Engine.Where("mail = ?", iMail).Count(&CommentFieldData{Mail: iMail})
	return n
}

// GetCommentFieldDataCountViaHomepage Get CommentFieldData via Homepage
func GetCommentFieldDataCountViaHomepage(iHomepage string) int64 {
	n, _ := Engine.Where("homepage = ?", iHomepage).Count(&CommentFieldData{Homepage: iHomepage})
	return n
}

// GetCommentFieldDataCountViaHostname Get CommentFieldData via Hostname
func GetCommentFieldDataCountViaHostname(iHostname string) int64 {
	n, _ := Engine.Where("hostname = ?", iHostname).Count(&CommentFieldData{Hostname: iHostname})
	return n
}

// GetCommentFieldDataCountViaCreated Get CommentFieldData via Created
func GetCommentFieldDataCountViaCreated(iCreated int) int64 {
	n, _ := Engine.Where("created = ?", iCreated).Count(&CommentFieldData{Created: iCreated})
	return n
}

// GetCommentFieldDataCountViaChanged Get CommentFieldData via Changed
func GetCommentFieldDataCountViaChanged(iChanged int) int64 {
	n, _ := Engine.Where("changed = ?", iChanged).Count(&CommentFieldData{Changed: iChanged})
	return n
}

// GetCommentFieldDataCountViaStatus Get CommentFieldData via Status
func GetCommentFieldDataCountViaStatus(iStatus int) int64 {
	n, _ := Engine.Where("status = ?", iStatus).Count(&CommentFieldData{Status: iStatus})
	return n
}

// GetCommentFieldDataCountViaThread Get CommentFieldData via Thread
func GetCommentFieldDataCountViaThread(iThread string) int64 {
	n, _ := Engine.Where("thread = ?", iThread).Count(&CommentFieldData{Thread: iThread})
	return n
}

// GetCommentFieldDataCountViaEntityType Get CommentFieldData via EntityType
func GetCommentFieldDataCountViaEntityType(iEntityType string) int64 {
	n, _ := Engine.Where("entity_type = ?", iEntityType).Count(&CommentFieldData{EntityType: iEntityType})
	return n
}

// GetCommentFieldDataCountViaFieldName Get CommentFieldData via FieldName
func GetCommentFieldDataCountViaFieldName(iFieldName string) int64 {
	n, _ := Engine.Where("field_name = ?", iFieldName).Count(&CommentFieldData{FieldName: iFieldName})
	return n
}

// GetCommentFieldDataCountViaDefaultLangcode Get CommentFieldData via DefaultLangcode
func GetCommentFieldDataCountViaDefaultLangcode(iDefaultLangcode int) int64 {
	n, _ := Engine.Where("default_langcode = ?", iDefaultLangcode).Count(&CommentFieldData{DefaultLangcode: iDefaultLangcode})
	return n
}

// GetCommentFieldDatasByCidAndCommentTypeAndLangcode Get CommentFieldDatas via CidAndCommentTypeAndLangcode
func GetCommentFieldDatasByCidAndCommentTypeAndLangcode(offset int, limit int, Cid_ int64, CommentType_ string, Langcode_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and comment_type = ? and langcode = ?", Cid_, CommentType_, Langcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndCommentTypeAndPid Get CommentFieldDatas via CidAndCommentTypeAndPid
func GetCommentFieldDatasByCidAndCommentTypeAndPid(offset int, limit int, Cid_ int64, CommentType_ string, Pid_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and comment_type = ? and pid = ?", Cid_, CommentType_, Pid_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndCommentTypeAndEntityId Get CommentFieldDatas via CidAndCommentTypeAndEntityId
func GetCommentFieldDatasByCidAndCommentTypeAndEntityId(offset int, limit int, Cid_ int64, CommentType_ string, EntityId_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and comment_type = ? and entity_id = ?", Cid_, CommentType_, EntityId_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndCommentTypeAndSubject Get CommentFieldDatas via CidAndCommentTypeAndSubject
func GetCommentFieldDatasByCidAndCommentTypeAndSubject(offset int, limit int, Cid_ int64, CommentType_ string, Subject_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and comment_type = ? and subject = ?", Cid_, CommentType_, Subject_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndCommentTypeAndUid Get CommentFieldDatas via CidAndCommentTypeAndUid
func GetCommentFieldDatasByCidAndCommentTypeAndUid(offset int, limit int, Cid_ int64, CommentType_ string, Uid_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and comment_type = ? and uid = ?", Cid_, CommentType_, Uid_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndCommentTypeAndName Get CommentFieldDatas via CidAndCommentTypeAndName
func GetCommentFieldDatasByCidAndCommentTypeAndName(offset int, limit int, Cid_ int64, CommentType_ string, Name_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and comment_type = ? and name = ?", Cid_, CommentType_, Name_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndCommentTypeAndMail Get CommentFieldDatas via CidAndCommentTypeAndMail
func GetCommentFieldDatasByCidAndCommentTypeAndMail(offset int, limit int, Cid_ int64, CommentType_ string, Mail_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and comment_type = ? and mail = ?", Cid_, CommentType_, Mail_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndCommentTypeAndHomepage Get CommentFieldDatas via CidAndCommentTypeAndHomepage
func GetCommentFieldDatasByCidAndCommentTypeAndHomepage(offset int, limit int, Cid_ int64, CommentType_ string, Homepage_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and comment_type = ? and homepage = ?", Cid_, CommentType_, Homepage_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndCommentTypeAndHostname Get CommentFieldDatas via CidAndCommentTypeAndHostname
func GetCommentFieldDatasByCidAndCommentTypeAndHostname(offset int, limit int, Cid_ int64, CommentType_ string, Hostname_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and comment_type = ? and hostname = ?", Cid_, CommentType_, Hostname_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndCommentTypeAndCreated Get CommentFieldDatas via CidAndCommentTypeAndCreated
func GetCommentFieldDatasByCidAndCommentTypeAndCreated(offset int, limit int, Cid_ int64, CommentType_ string, Created_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and comment_type = ? and created = ?", Cid_, CommentType_, Created_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndCommentTypeAndChanged Get CommentFieldDatas via CidAndCommentTypeAndChanged
func GetCommentFieldDatasByCidAndCommentTypeAndChanged(offset int, limit int, Cid_ int64, CommentType_ string, Changed_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and comment_type = ? and changed = ?", Cid_, CommentType_, Changed_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndCommentTypeAndStatus Get CommentFieldDatas via CidAndCommentTypeAndStatus
func GetCommentFieldDatasByCidAndCommentTypeAndStatus(offset int, limit int, Cid_ int64, CommentType_ string, Status_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and comment_type = ? and status = ?", Cid_, CommentType_, Status_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndCommentTypeAndThread Get CommentFieldDatas via CidAndCommentTypeAndThread
func GetCommentFieldDatasByCidAndCommentTypeAndThread(offset int, limit int, Cid_ int64, CommentType_ string, Thread_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and comment_type = ? and thread = ?", Cid_, CommentType_, Thread_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndCommentTypeAndEntityType Get CommentFieldDatas via CidAndCommentTypeAndEntityType
func GetCommentFieldDatasByCidAndCommentTypeAndEntityType(offset int, limit int, Cid_ int64, CommentType_ string, EntityType_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and comment_type = ? and entity_type = ?", Cid_, CommentType_, EntityType_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndCommentTypeAndFieldName Get CommentFieldDatas via CidAndCommentTypeAndFieldName
func GetCommentFieldDatasByCidAndCommentTypeAndFieldName(offset int, limit int, Cid_ int64, CommentType_ string, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and comment_type = ? and field_name = ?", Cid_, CommentType_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndCommentTypeAndDefaultLangcode Get CommentFieldDatas via CidAndCommentTypeAndDefaultLangcode
func GetCommentFieldDatasByCidAndCommentTypeAndDefaultLangcode(offset int, limit int, Cid_ int64, CommentType_ string, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and comment_type = ? and default_langcode = ?", Cid_, CommentType_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndLangcodeAndPid Get CommentFieldDatas via CidAndLangcodeAndPid
func GetCommentFieldDatasByCidAndLangcodeAndPid(offset int, limit int, Cid_ int64, Langcode_ string, Pid_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and langcode = ? and pid = ?", Cid_, Langcode_, Pid_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndLangcodeAndEntityId Get CommentFieldDatas via CidAndLangcodeAndEntityId
func GetCommentFieldDatasByCidAndLangcodeAndEntityId(offset int, limit int, Cid_ int64, Langcode_ string, EntityId_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and langcode = ? and entity_id = ?", Cid_, Langcode_, EntityId_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndLangcodeAndSubject Get CommentFieldDatas via CidAndLangcodeAndSubject
func GetCommentFieldDatasByCidAndLangcodeAndSubject(offset int, limit int, Cid_ int64, Langcode_ string, Subject_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and langcode = ? and subject = ?", Cid_, Langcode_, Subject_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndLangcodeAndUid Get CommentFieldDatas via CidAndLangcodeAndUid
func GetCommentFieldDatasByCidAndLangcodeAndUid(offset int, limit int, Cid_ int64, Langcode_ string, Uid_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and langcode = ? and uid = ?", Cid_, Langcode_, Uid_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndLangcodeAndName Get CommentFieldDatas via CidAndLangcodeAndName
func GetCommentFieldDatasByCidAndLangcodeAndName(offset int, limit int, Cid_ int64, Langcode_ string, Name_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and langcode = ? and name = ?", Cid_, Langcode_, Name_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndLangcodeAndMail Get CommentFieldDatas via CidAndLangcodeAndMail
func GetCommentFieldDatasByCidAndLangcodeAndMail(offset int, limit int, Cid_ int64, Langcode_ string, Mail_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and langcode = ? and mail = ?", Cid_, Langcode_, Mail_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndLangcodeAndHomepage Get CommentFieldDatas via CidAndLangcodeAndHomepage
func GetCommentFieldDatasByCidAndLangcodeAndHomepage(offset int, limit int, Cid_ int64, Langcode_ string, Homepage_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and langcode = ? and homepage = ?", Cid_, Langcode_, Homepage_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndLangcodeAndHostname Get CommentFieldDatas via CidAndLangcodeAndHostname
func GetCommentFieldDatasByCidAndLangcodeAndHostname(offset int, limit int, Cid_ int64, Langcode_ string, Hostname_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and langcode = ? and hostname = ?", Cid_, Langcode_, Hostname_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndLangcodeAndCreated Get CommentFieldDatas via CidAndLangcodeAndCreated
func GetCommentFieldDatasByCidAndLangcodeAndCreated(offset int, limit int, Cid_ int64, Langcode_ string, Created_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and langcode = ? and created = ?", Cid_, Langcode_, Created_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndLangcodeAndChanged Get CommentFieldDatas via CidAndLangcodeAndChanged
func GetCommentFieldDatasByCidAndLangcodeAndChanged(offset int, limit int, Cid_ int64, Langcode_ string, Changed_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and langcode = ? and changed = ?", Cid_, Langcode_, Changed_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndLangcodeAndStatus Get CommentFieldDatas via CidAndLangcodeAndStatus
func GetCommentFieldDatasByCidAndLangcodeAndStatus(offset int, limit int, Cid_ int64, Langcode_ string, Status_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and langcode = ? and status = ?", Cid_, Langcode_, Status_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndLangcodeAndThread Get CommentFieldDatas via CidAndLangcodeAndThread
func GetCommentFieldDatasByCidAndLangcodeAndThread(offset int, limit int, Cid_ int64, Langcode_ string, Thread_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and langcode = ? and thread = ?", Cid_, Langcode_, Thread_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndLangcodeAndEntityType Get CommentFieldDatas via CidAndLangcodeAndEntityType
func GetCommentFieldDatasByCidAndLangcodeAndEntityType(offset int, limit int, Cid_ int64, Langcode_ string, EntityType_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and langcode = ? and entity_type = ?", Cid_, Langcode_, EntityType_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndLangcodeAndFieldName Get CommentFieldDatas via CidAndLangcodeAndFieldName
func GetCommentFieldDatasByCidAndLangcodeAndFieldName(offset int, limit int, Cid_ int64, Langcode_ string, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and langcode = ? and field_name = ?", Cid_, Langcode_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndLangcodeAndDefaultLangcode Get CommentFieldDatas via CidAndLangcodeAndDefaultLangcode
func GetCommentFieldDatasByCidAndLangcodeAndDefaultLangcode(offset int, limit int, Cid_ int64, Langcode_ string, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and langcode = ? and default_langcode = ?", Cid_, Langcode_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndPidAndEntityId Get CommentFieldDatas via CidAndPidAndEntityId
func GetCommentFieldDatasByCidAndPidAndEntityId(offset int, limit int, Cid_ int64, Pid_ int, EntityId_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and pid = ? and entity_id = ?", Cid_, Pid_, EntityId_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndPidAndSubject Get CommentFieldDatas via CidAndPidAndSubject
func GetCommentFieldDatasByCidAndPidAndSubject(offset int, limit int, Cid_ int64, Pid_ int, Subject_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and pid = ? and subject = ?", Cid_, Pid_, Subject_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndPidAndUid Get CommentFieldDatas via CidAndPidAndUid
func GetCommentFieldDatasByCidAndPidAndUid(offset int, limit int, Cid_ int64, Pid_ int, Uid_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and pid = ? and uid = ?", Cid_, Pid_, Uid_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndPidAndName Get CommentFieldDatas via CidAndPidAndName
func GetCommentFieldDatasByCidAndPidAndName(offset int, limit int, Cid_ int64, Pid_ int, Name_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and pid = ? and name = ?", Cid_, Pid_, Name_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndPidAndMail Get CommentFieldDatas via CidAndPidAndMail
func GetCommentFieldDatasByCidAndPidAndMail(offset int, limit int, Cid_ int64, Pid_ int, Mail_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and pid = ? and mail = ?", Cid_, Pid_, Mail_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndPidAndHomepage Get CommentFieldDatas via CidAndPidAndHomepage
func GetCommentFieldDatasByCidAndPidAndHomepage(offset int, limit int, Cid_ int64, Pid_ int, Homepage_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and pid = ? and homepage = ?", Cid_, Pid_, Homepage_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndPidAndHostname Get CommentFieldDatas via CidAndPidAndHostname
func GetCommentFieldDatasByCidAndPidAndHostname(offset int, limit int, Cid_ int64, Pid_ int, Hostname_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and pid = ? and hostname = ?", Cid_, Pid_, Hostname_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndPidAndCreated Get CommentFieldDatas via CidAndPidAndCreated
func GetCommentFieldDatasByCidAndPidAndCreated(offset int, limit int, Cid_ int64, Pid_ int, Created_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and pid = ? and created = ?", Cid_, Pid_, Created_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndPidAndChanged Get CommentFieldDatas via CidAndPidAndChanged
func GetCommentFieldDatasByCidAndPidAndChanged(offset int, limit int, Cid_ int64, Pid_ int, Changed_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and pid = ? and changed = ?", Cid_, Pid_, Changed_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndPidAndStatus Get CommentFieldDatas via CidAndPidAndStatus
func GetCommentFieldDatasByCidAndPidAndStatus(offset int, limit int, Cid_ int64, Pid_ int, Status_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and pid = ? and status = ?", Cid_, Pid_, Status_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndPidAndThread Get CommentFieldDatas via CidAndPidAndThread
func GetCommentFieldDatasByCidAndPidAndThread(offset int, limit int, Cid_ int64, Pid_ int, Thread_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and pid = ? and thread = ?", Cid_, Pid_, Thread_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndPidAndEntityType Get CommentFieldDatas via CidAndPidAndEntityType
func GetCommentFieldDatasByCidAndPidAndEntityType(offset int, limit int, Cid_ int64, Pid_ int, EntityType_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and pid = ? and entity_type = ?", Cid_, Pid_, EntityType_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndPidAndFieldName Get CommentFieldDatas via CidAndPidAndFieldName
func GetCommentFieldDatasByCidAndPidAndFieldName(offset int, limit int, Cid_ int64, Pid_ int, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and pid = ? and field_name = ?", Cid_, Pid_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndPidAndDefaultLangcode Get CommentFieldDatas via CidAndPidAndDefaultLangcode
func GetCommentFieldDatasByCidAndPidAndDefaultLangcode(offset int, limit int, Cid_ int64, Pid_ int, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and pid = ? and default_langcode = ?", Cid_, Pid_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndEntityIdAndSubject Get CommentFieldDatas via CidAndEntityIdAndSubject
func GetCommentFieldDatasByCidAndEntityIdAndSubject(offset int, limit int, Cid_ int64, EntityId_ int, Subject_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and entity_id = ? and subject = ?", Cid_, EntityId_, Subject_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndEntityIdAndUid Get CommentFieldDatas via CidAndEntityIdAndUid
func GetCommentFieldDatasByCidAndEntityIdAndUid(offset int, limit int, Cid_ int64, EntityId_ int, Uid_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and entity_id = ? and uid = ?", Cid_, EntityId_, Uid_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndEntityIdAndName Get CommentFieldDatas via CidAndEntityIdAndName
func GetCommentFieldDatasByCidAndEntityIdAndName(offset int, limit int, Cid_ int64, EntityId_ int, Name_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and entity_id = ? and name = ?", Cid_, EntityId_, Name_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndEntityIdAndMail Get CommentFieldDatas via CidAndEntityIdAndMail
func GetCommentFieldDatasByCidAndEntityIdAndMail(offset int, limit int, Cid_ int64, EntityId_ int, Mail_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and entity_id = ? and mail = ?", Cid_, EntityId_, Mail_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndEntityIdAndHomepage Get CommentFieldDatas via CidAndEntityIdAndHomepage
func GetCommentFieldDatasByCidAndEntityIdAndHomepage(offset int, limit int, Cid_ int64, EntityId_ int, Homepage_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and entity_id = ? and homepage = ?", Cid_, EntityId_, Homepage_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndEntityIdAndHostname Get CommentFieldDatas via CidAndEntityIdAndHostname
func GetCommentFieldDatasByCidAndEntityIdAndHostname(offset int, limit int, Cid_ int64, EntityId_ int, Hostname_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and entity_id = ? and hostname = ?", Cid_, EntityId_, Hostname_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndEntityIdAndCreated Get CommentFieldDatas via CidAndEntityIdAndCreated
func GetCommentFieldDatasByCidAndEntityIdAndCreated(offset int, limit int, Cid_ int64, EntityId_ int, Created_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and entity_id = ? and created = ?", Cid_, EntityId_, Created_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndEntityIdAndChanged Get CommentFieldDatas via CidAndEntityIdAndChanged
func GetCommentFieldDatasByCidAndEntityIdAndChanged(offset int, limit int, Cid_ int64, EntityId_ int, Changed_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and entity_id = ? and changed = ?", Cid_, EntityId_, Changed_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndEntityIdAndStatus Get CommentFieldDatas via CidAndEntityIdAndStatus
func GetCommentFieldDatasByCidAndEntityIdAndStatus(offset int, limit int, Cid_ int64, EntityId_ int, Status_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and entity_id = ? and status = ?", Cid_, EntityId_, Status_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndEntityIdAndThread Get CommentFieldDatas via CidAndEntityIdAndThread
func GetCommentFieldDatasByCidAndEntityIdAndThread(offset int, limit int, Cid_ int64, EntityId_ int, Thread_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and entity_id = ? and thread = ?", Cid_, EntityId_, Thread_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndEntityIdAndEntityType Get CommentFieldDatas via CidAndEntityIdAndEntityType
func GetCommentFieldDatasByCidAndEntityIdAndEntityType(offset int, limit int, Cid_ int64, EntityId_ int, EntityType_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and entity_id = ? and entity_type = ?", Cid_, EntityId_, EntityType_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndEntityIdAndFieldName Get CommentFieldDatas via CidAndEntityIdAndFieldName
func GetCommentFieldDatasByCidAndEntityIdAndFieldName(offset int, limit int, Cid_ int64, EntityId_ int, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and entity_id = ? and field_name = ?", Cid_, EntityId_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndEntityIdAndDefaultLangcode Get CommentFieldDatas via CidAndEntityIdAndDefaultLangcode
func GetCommentFieldDatasByCidAndEntityIdAndDefaultLangcode(offset int, limit int, Cid_ int64, EntityId_ int, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and entity_id = ? and default_langcode = ?", Cid_, EntityId_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndSubjectAndUid Get CommentFieldDatas via CidAndSubjectAndUid
func GetCommentFieldDatasByCidAndSubjectAndUid(offset int, limit int, Cid_ int64, Subject_ string, Uid_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and subject = ? and uid = ?", Cid_, Subject_, Uid_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndSubjectAndName Get CommentFieldDatas via CidAndSubjectAndName
func GetCommentFieldDatasByCidAndSubjectAndName(offset int, limit int, Cid_ int64, Subject_ string, Name_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and subject = ? and name = ?", Cid_, Subject_, Name_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndSubjectAndMail Get CommentFieldDatas via CidAndSubjectAndMail
func GetCommentFieldDatasByCidAndSubjectAndMail(offset int, limit int, Cid_ int64, Subject_ string, Mail_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and subject = ? and mail = ?", Cid_, Subject_, Mail_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndSubjectAndHomepage Get CommentFieldDatas via CidAndSubjectAndHomepage
func GetCommentFieldDatasByCidAndSubjectAndHomepage(offset int, limit int, Cid_ int64, Subject_ string, Homepage_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and subject = ? and homepage = ?", Cid_, Subject_, Homepage_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndSubjectAndHostname Get CommentFieldDatas via CidAndSubjectAndHostname
func GetCommentFieldDatasByCidAndSubjectAndHostname(offset int, limit int, Cid_ int64, Subject_ string, Hostname_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and subject = ? and hostname = ?", Cid_, Subject_, Hostname_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndSubjectAndCreated Get CommentFieldDatas via CidAndSubjectAndCreated
func GetCommentFieldDatasByCidAndSubjectAndCreated(offset int, limit int, Cid_ int64, Subject_ string, Created_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and subject = ? and created = ?", Cid_, Subject_, Created_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndSubjectAndChanged Get CommentFieldDatas via CidAndSubjectAndChanged
func GetCommentFieldDatasByCidAndSubjectAndChanged(offset int, limit int, Cid_ int64, Subject_ string, Changed_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and subject = ? and changed = ?", Cid_, Subject_, Changed_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndSubjectAndStatus Get CommentFieldDatas via CidAndSubjectAndStatus
func GetCommentFieldDatasByCidAndSubjectAndStatus(offset int, limit int, Cid_ int64, Subject_ string, Status_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and subject = ? and status = ?", Cid_, Subject_, Status_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndSubjectAndThread Get CommentFieldDatas via CidAndSubjectAndThread
func GetCommentFieldDatasByCidAndSubjectAndThread(offset int, limit int, Cid_ int64, Subject_ string, Thread_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and subject = ? and thread = ?", Cid_, Subject_, Thread_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndSubjectAndEntityType Get CommentFieldDatas via CidAndSubjectAndEntityType
func GetCommentFieldDatasByCidAndSubjectAndEntityType(offset int, limit int, Cid_ int64, Subject_ string, EntityType_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and subject = ? and entity_type = ?", Cid_, Subject_, EntityType_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndSubjectAndFieldName Get CommentFieldDatas via CidAndSubjectAndFieldName
func GetCommentFieldDatasByCidAndSubjectAndFieldName(offset int, limit int, Cid_ int64, Subject_ string, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and subject = ? and field_name = ?", Cid_, Subject_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndSubjectAndDefaultLangcode Get CommentFieldDatas via CidAndSubjectAndDefaultLangcode
func GetCommentFieldDatasByCidAndSubjectAndDefaultLangcode(offset int, limit int, Cid_ int64, Subject_ string, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and subject = ? and default_langcode = ?", Cid_, Subject_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndUidAndName Get CommentFieldDatas via CidAndUidAndName
func GetCommentFieldDatasByCidAndUidAndName(offset int, limit int, Cid_ int64, Uid_ int, Name_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and uid = ? and name = ?", Cid_, Uid_, Name_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndUidAndMail Get CommentFieldDatas via CidAndUidAndMail
func GetCommentFieldDatasByCidAndUidAndMail(offset int, limit int, Cid_ int64, Uid_ int, Mail_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and uid = ? and mail = ?", Cid_, Uid_, Mail_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndUidAndHomepage Get CommentFieldDatas via CidAndUidAndHomepage
func GetCommentFieldDatasByCidAndUidAndHomepage(offset int, limit int, Cid_ int64, Uid_ int, Homepage_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and uid = ? and homepage = ?", Cid_, Uid_, Homepage_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndUidAndHostname Get CommentFieldDatas via CidAndUidAndHostname
func GetCommentFieldDatasByCidAndUidAndHostname(offset int, limit int, Cid_ int64, Uid_ int, Hostname_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and uid = ? and hostname = ?", Cid_, Uid_, Hostname_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndUidAndCreated Get CommentFieldDatas via CidAndUidAndCreated
func GetCommentFieldDatasByCidAndUidAndCreated(offset int, limit int, Cid_ int64, Uid_ int, Created_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and uid = ? and created = ?", Cid_, Uid_, Created_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndUidAndChanged Get CommentFieldDatas via CidAndUidAndChanged
func GetCommentFieldDatasByCidAndUidAndChanged(offset int, limit int, Cid_ int64, Uid_ int, Changed_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and uid = ? and changed = ?", Cid_, Uid_, Changed_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndUidAndStatus Get CommentFieldDatas via CidAndUidAndStatus
func GetCommentFieldDatasByCidAndUidAndStatus(offset int, limit int, Cid_ int64, Uid_ int, Status_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and uid = ? and status = ?", Cid_, Uid_, Status_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndUidAndThread Get CommentFieldDatas via CidAndUidAndThread
func GetCommentFieldDatasByCidAndUidAndThread(offset int, limit int, Cid_ int64, Uid_ int, Thread_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and uid = ? and thread = ?", Cid_, Uid_, Thread_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndUidAndEntityType Get CommentFieldDatas via CidAndUidAndEntityType
func GetCommentFieldDatasByCidAndUidAndEntityType(offset int, limit int, Cid_ int64, Uid_ int, EntityType_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and uid = ? and entity_type = ?", Cid_, Uid_, EntityType_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndUidAndFieldName Get CommentFieldDatas via CidAndUidAndFieldName
func GetCommentFieldDatasByCidAndUidAndFieldName(offset int, limit int, Cid_ int64, Uid_ int, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and uid = ? and field_name = ?", Cid_, Uid_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndUidAndDefaultLangcode Get CommentFieldDatas via CidAndUidAndDefaultLangcode
func GetCommentFieldDatasByCidAndUidAndDefaultLangcode(offset int, limit int, Cid_ int64, Uid_ int, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and uid = ? and default_langcode = ?", Cid_, Uid_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndNameAndMail Get CommentFieldDatas via CidAndNameAndMail
func GetCommentFieldDatasByCidAndNameAndMail(offset int, limit int, Cid_ int64, Name_ string, Mail_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and name = ? and mail = ?", Cid_, Name_, Mail_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndNameAndHomepage Get CommentFieldDatas via CidAndNameAndHomepage
func GetCommentFieldDatasByCidAndNameAndHomepage(offset int, limit int, Cid_ int64, Name_ string, Homepage_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and name = ? and homepage = ?", Cid_, Name_, Homepage_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndNameAndHostname Get CommentFieldDatas via CidAndNameAndHostname
func GetCommentFieldDatasByCidAndNameAndHostname(offset int, limit int, Cid_ int64, Name_ string, Hostname_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and name = ? and hostname = ?", Cid_, Name_, Hostname_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndNameAndCreated Get CommentFieldDatas via CidAndNameAndCreated
func GetCommentFieldDatasByCidAndNameAndCreated(offset int, limit int, Cid_ int64, Name_ string, Created_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and name = ? and created = ?", Cid_, Name_, Created_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndNameAndChanged Get CommentFieldDatas via CidAndNameAndChanged
func GetCommentFieldDatasByCidAndNameAndChanged(offset int, limit int, Cid_ int64, Name_ string, Changed_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and name = ? and changed = ?", Cid_, Name_, Changed_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndNameAndStatus Get CommentFieldDatas via CidAndNameAndStatus
func GetCommentFieldDatasByCidAndNameAndStatus(offset int, limit int, Cid_ int64, Name_ string, Status_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and name = ? and status = ?", Cid_, Name_, Status_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndNameAndThread Get CommentFieldDatas via CidAndNameAndThread
func GetCommentFieldDatasByCidAndNameAndThread(offset int, limit int, Cid_ int64, Name_ string, Thread_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and name = ? and thread = ?", Cid_, Name_, Thread_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndNameAndEntityType Get CommentFieldDatas via CidAndNameAndEntityType
func GetCommentFieldDatasByCidAndNameAndEntityType(offset int, limit int, Cid_ int64, Name_ string, EntityType_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and name = ? and entity_type = ?", Cid_, Name_, EntityType_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndNameAndFieldName Get CommentFieldDatas via CidAndNameAndFieldName
func GetCommentFieldDatasByCidAndNameAndFieldName(offset int, limit int, Cid_ int64, Name_ string, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and name = ? and field_name = ?", Cid_, Name_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndNameAndDefaultLangcode Get CommentFieldDatas via CidAndNameAndDefaultLangcode
func GetCommentFieldDatasByCidAndNameAndDefaultLangcode(offset int, limit int, Cid_ int64, Name_ string, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and name = ? and default_langcode = ?", Cid_, Name_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndMailAndHomepage Get CommentFieldDatas via CidAndMailAndHomepage
func GetCommentFieldDatasByCidAndMailAndHomepage(offset int, limit int, Cid_ int64, Mail_ string, Homepage_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and mail = ? and homepage = ?", Cid_, Mail_, Homepage_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndMailAndHostname Get CommentFieldDatas via CidAndMailAndHostname
func GetCommentFieldDatasByCidAndMailAndHostname(offset int, limit int, Cid_ int64, Mail_ string, Hostname_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and mail = ? and hostname = ?", Cid_, Mail_, Hostname_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndMailAndCreated Get CommentFieldDatas via CidAndMailAndCreated
func GetCommentFieldDatasByCidAndMailAndCreated(offset int, limit int, Cid_ int64, Mail_ string, Created_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and mail = ? and created = ?", Cid_, Mail_, Created_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndMailAndChanged Get CommentFieldDatas via CidAndMailAndChanged
func GetCommentFieldDatasByCidAndMailAndChanged(offset int, limit int, Cid_ int64, Mail_ string, Changed_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and mail = ? and changed = ?", Cid_, Mail_, Changed_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndMailAndStatus Get CommentFieldDatas via CidAndMailAndStatus
func GetCommentFieldDatasByCidAndMailAndStatus(offset int, limit int, Cid_ int64, Mail_ string, Status_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and mail = ? and status = ?", Cid_, Mail_, Status_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndMailAndThread Get CommentFieldDatas via CidAndMailAndThread
func GetCommentFieldDatasByCidAndMailAndThread(offset int, limit int, Cid_ int64, Mail_ string, Thread_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and mail = ? and thread = ?", Cid_, Mail_, Thread_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndMailAndEntityType Get CommentFieldDatas via CidAndMailAndEntityType
func GetCommentFieldDatasByCidAndMailAndEntityType(offset int, limit int, Cid_ int64, Mail_ string, EntityType_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and mail = ? and entity_type = ?", Cid_, Mail_, EntityType_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndMailAndFieldName Get CommentFieldDatas via CidAndMailAndFieldName
func GetCommentFieldDatasByCidAndMailAndFieldName(offset int, limit int, Cid_ int64, Mail_ string, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and mail = ? and field_name = ?", Cid_, Mail_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndMailAndDefaultLangcode Get CommentFieldDatas via CidAndMailAndDefaultLangcode
func GetCommentFieldDatasByCidAndMailAndDefaultLangcode(offset int, limit int, Cid_ int64, Mail_ string, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and mail = ? and default_langcode = ?", Cid_, Mail_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndHomepageAndHostname Get CommentFieldDatas via CidAndHomepageAndHostname
func GetCommentFieldDatasByCidAndHomepageAndHostname(offset int, limit int, Cid_ int64, Homepage_ string, Hostname_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and homepage = ? and hostname = ?", Cid_, Homepage_, Hostname_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndHomepageAndCreated Get CommentFieldDatas via CidAndHomepageAndCreated
func GetCommentFieldDatasByCidAndHomepageAndCreated(offset int, limit int, Cid_ int64, Homepage_ string, Created_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and homepage = ? and created = ?", Cid_, Homepage_, Created_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndHomepageAndChanged Get CommentFieldDatas via CidAndHomepageAndChanged
func GetCommentFieldDatasByCidAndHomepageAndChanged(offset int, limit int, Cid_ int64, Homepage_ string, Changed_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and homepage = ? and changed = ?", Cid_, Homepage_, Changed_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndHomepageAndStatus Get CommentFieldDatas via CidAndHomepageAndStatus
func GetCommentFieldDatasByCidAndHomepageAndStatus(offset int, limit int, Cid_ int64, Homepage_ string, Status_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and homepage = ? and status = ?", Cid_, Homepage_, Status_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndHomepageAndThread Get CommentFieldDatas via CidAndHomepageAndThread
func GetCommentFieldDatasByCidAndHomepageAndThread(offset int, limit int, Cid_ int64, Homepage_ string, Thread_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and homepage = ? and thread = ?", Cid_, Homepage_, Thread_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndHomepageAndEntityType Get CommentFieldDatas via CidAndHomepageAndEntityType
func GetCommentFieldDatasByCidAndHomepageAndEntityType(offset int, limit int, Cid_ int64, Homepage_ string, EntityType_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and homepage = ? and entity_type = ?", Cid_, Homepage_, EntityType_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndHomepageAndFieldName Get CommentFieldDatas via CidAndHomepageAndFieldName
func GetCommentFieldDatasByCidAndHomepageAndFieldName(offset int, limit int, Cid_ int64, Homepage_ string, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and homepage = ? and field_name = ?", Cid_, Homepage_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndHomepageAndDefaultLangcode Get CommentFieldDatas via CidAndHomepageAndDefaultLangcode
func GetCommentFieldDatasByCidAndHomepageAndDefaultLangcode(offset int, limit int, Cid_ int64, Homepage_ string, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and homepage = ? and default_langcode = ?", Cid_, Homepage_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndHostnameAndCreated Get CommentFieldDatas via CidAndHostnameAndCreated
func GetCommentFieldDatasByCidAndHostnameAndCreated(offset int, limit int, Cid_ int64, Hostname_ string, Created_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and hostname = ? and created = ?", Cid_, Hostname_, Created_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndHostnameAndChanged Get CommentFieldDatas via CidAndHostnameAndChanged
func GetCommentFieldDatasByCidAndHostnameAndChanged(offset int, limit int, Cid_ int64, Hostname_ string, Changed_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and hostname = ? and changed = ?", Cid_, Hostname_, Changed_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndHostnameAndStatus Get CommentFieldDatas via CidAndHostnameAndStatus
func GetCommentFieldDatasByCidAndHostnameAndStatus(offset int, limit int, Cid_ int64, Hostname_ string, Status_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and hostname = ? and status = ?", Cid_, Hostname_, Status_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndHostnameAndThread Get CommentFieldDatas via CidAndHostnameAndThread
func GetCommentFieldDatasByCidAndHostnameAndThread(offset int, limit int, Cid_ int64, Hostname_ string, Thread_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and hostname = ? and thread = ?", Cid_, Hostname_, Thread_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndHostnameAndEntityType Get CommentFieldDatas via CidAndHostnameAndEntityType
func GetCommentFieldDatasByCidAndHostnameAndEntityType(offset int, limit int, Cid_ int64, Hostname_ string, EntityType_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and hostname = ? and entity_type = ?", Cid_, Hostname_, EntityType_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndHostnameAndFieldName Get CommentFieldDatas via CidAndHostnameAndFieldName
func GetCommentFieldDatasByCidAndHostnameAndFieldName(offset int, limit int, Cid_ int64, Hostname_ string, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and hostname = ? and field_name = ?", Cid_, Hostname_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndHostnameAndDefaultLangcode Get CommentFieldDatas via CidAndHostnameAndDefaultLangcode
func GetCommentFieldDatasByCidAndHostnameAndDefaultLangcode(offset int, limit int, Cid_ int64, Hostname_ string, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and hostname = ? and default_langcode = ?", Cid_, Hostname_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndCreatedAndChanged Get CommentFieldDatas via CidAndCreatedAndChanged
func GetCommentFieldDatasByCidAndCreatedAndChanged(offset int, limit int, Cid_ int64, Created_ int, Changed_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and created = ? and changed = ?", Cid_, Created_, Changed_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndCreatedAndStatus Get CommentFieldDatas via CidAndCreatedAndStatus
func GetCommentFieldDatasByCidAndCreatedAndStatus(offset int, limit int, Cid_ int64, Created_ int, Status_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and created = ? and status = ?", Cid_, Created_, Status_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndCreatedAndThread Get CommentFieldDatas via CidAndCreatedAndThread
func GetCommentFieldDatasByCidAndCreatedAndThread(offset int, limit int, Cid_ int64, Created_ int, Thread_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and created = ? and thread = ?", Cid_, Created_, Thread_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndCreatedAndEntityType Get CommentFieldDatas via CidAndCreatedAndEntityType
func GetCommentFieldDatasByCidAndCreatedAndEntityType(offset int, limit int, Cid_ int64, Created_ int, EntityType_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and created = ? and entity_type = ?", Cid_, Created_, EntityType_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndCreatedAndFieldName Get CommentFieldDatas via CidAndCreatedAndFieldName
func GetCommentFieldDatasByCidAndCreatedAndFieldName(offset int, limit int, Cid_ int64, Created_ int, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and created = ? and field_name = ?", Cid_, Created_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndCreatedAndDefaultLangcode Get CommentFieldDatas via CidAndCreatedAndDefaultLangcode
func GetCommentFieldDatasByCidAndCreatedAndDefaultLangcode(offset int, limit int, Cid_ int64, Created_ int, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and created = ? and default_langcode = ?", Cid_, Created_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndChangedAndStatus Get CommentFieldDatas via CidAndChangedAndStatus
func GetCommentFieldDatasByCidAndChangedAndStatus(offset int, limit int, Cid_ int64, Changed_ int, Status_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and changed = ? and status = ?", Cid_, Changed_, Status_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndChangedAndThread Get CommentFieldDatas via CidAndChangedAndThread
func GetCommentFieldDatasByCidAndChangedAndThread(offset int, limit int, Cid_ int64, Changed_ int, Thread_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and changed = ? and thread = ?", Cid_, Changed_, Thread_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndChangedAndEntityType Get CommentFieldDatas via CidAndChangedAndEntityType
func GetCommentFieldDatasByCidAndChangedAndEntityType(offset int, limit int, Cid_ int64, Changed_ int, EntityType_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and changed = ? and entity_type = ?", Cid_, Changed_, EntityType_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndChangedAndFieldName Get CommentFieldDatas via CidAndChangedAndFieldName
func GetCommentFieldDatasByCidAndChangedAndFieldName(offset int, limit int, Cid_ int64, Changed_ int, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and changed = ? and field_name = ?", Cid_, Changed_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndChangedAndDefaultLangcode Get CommentFieldDatas via CidAndChangedAndDefaultLangcode
func GetCommentFieldDatasByCidAndChangedAndDefaultLangcode(offset int, limit int, Cid_ int64, Changed_ int, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and changed = ? and default_langcode = ?", Cid_, Changed_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndStatusAndThread Get CommentFieldDatas via CidAndStatusAndThread
func GetCommentFieldDatasByCidAndStatusAndThread(offset int, limit int, Cid_ int64, Status_ int, Thread_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and status = ? and thread = ?", Cid_, Status_, Thread_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndStatusAndEntityType Get CommentFieldDatas via CidAndStatusAndEntityType
func GetCommentFieldDatasByCidAndStatusAndEntityType(offset int, limit int, Cid_ int64, Status_ int, EntityType_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and status = ? and entity_type = ?", Cid_, Status_, EntityType_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndStatusAndFieldName Get CommentFieldDatas via CidAndStatusAndFieldName
func GetCommentFieldDatasByCidAndStatusAndFieldName(offset int, limit int, Cid_ int64, Status_ int, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and status = ? and field_name = ?", Cid_, Status_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndStatusAndDefaultLangcode Get CommentFieldDatas via CidAndStatusAndDefaultLangcode
func GetCommentFieldDatasByCidAndStatusAndDefaultLangcode(offset int, limit int, Cid_ int64, Status_ int, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and status = ? and default_langcode = ?", Cid_, Status_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndThreadAndEntityType Get CommentFieldDatas via CidAndThreadAndEntityType
func GetCommentFieldDatasByCidAndThreadAndEntityType(offset int, limit int, Cid_ int64, Thread_ string, EntityType_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and thread = ? and entity_type = ?", Cid_, Thread_, EntityType_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndThreadAndFieldName Get CommentFieldDatas via CidAndThreadAndFieldName
func GetCommentFieldDatasByCidAndThreadAndFieldName(offset int, limit int, Cid_ int64, Thread_ string, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and thread = ? and field_name = ?", Cid_, Thread_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndThreadAndDefaultLangcode Get CommentFieldDatas via CidAndThreadAndDefaultLangcode
func GetCommentFieldDatasByCidAndThreadAndDefaultLangcode(offset int, limit int, Cid_ int64, Thread_ string, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and thread = ? and default_langcode = ?", Cid_, Thread_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndEntityTypeAndFieldName Get CommentFieldDatas via CidAndEntityTypeAndFieldName
func GetCommentFieldDatasByCidAndEntityTypeAndFieldName(offset int, limit int, Cid_ int64, EntityType_ string, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and entity_type = ? and field_name = ?", Cid_, EntityType_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndEntityTypeAndDefaultLangcode Get CommentFieldDatas via CidAndEntityTypeAndDefaultLangcode
func GetCommentFieldDatasByCidAndEntityTypeAndDefaultLangcode(offset int, limit int, Cid_ int64, EntityType_ string, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and entity_type = ? and default_langcode = ?", Cid_, EntityType_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndFieldNameAndDefaultLangcode Get CommentFieldDatas via CidAndFieldNameAndDefaultLangcode
func GetCommentFieldDatasByCidAndFieldNameAndDefaultLangcode(offset int, limit int, Cid_ int64, FieldName_ string, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and field_name = ? and default_langcode = ?", Cid_, FieldName_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndLangcodeAndPid Get CommentFieldDatas via CommentTypeAndLangcodeAndPid
func GetCommentFieldDatasByCommentTypeAndLangcodeAndPid(offset int, limit int, CommentType_ string, Langcode_ string, Pid_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and langcode = ? and pid = ?", CommentType_, Langcode_, Pid_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndLangcodeAndEntityId Get CommentFieldDatas via CommentTypeAndLangcodeAndEntityId
func GetCommentFieldDatasByCommentTypeAndLangcodeAndEntityId(offset int, limit int, CommentType_ string, Langcode_ string, EntityId_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and langcode = ? and entity_id = ?", CommentType_, Langcode_, EntityId_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndLangcodeAndSubject Get CommentFieldDatas via CommentTypeAndLangcodeAndSubject
func GetCommentFieldDatasByCommentTypeAndLangcodeAndSubject(offset int, limit int, CommentType_ string, Langcode_ string, Subject_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and langcode = ? and subject = ?", CommentType_, Langcode_, Subject_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndLangcodeAndUid Get CommentFieldDatas via CommentTypeAndLangcodeAndUid
func GetCommentFieldDatasByCommentTypeAndLangcodeAndUid(offset int, limit int, CommentType_ string, Langcode_ string, Uid_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and langcode = ? and uid = ?", CommentType_, Langcode_, Uid_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndLangcodeAndName Get CommentFieldDatas via CommentTypeAndLangcodeAndName
func GetCommentFieldDatasByCommentTypeAndLangcodeAndName(offset int, limit int, CommentType_ string, Langcode_ string, Name_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and langcode = ? and name = ?", CommentType_, Langcode_, Name_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndLangcodeAndMail Get CommentFieldDatas via CommentTypeAndLangcodeAndMail
func GetCommentFieldDatasByCommentTypeAndLangcodeAndMail(offset int, limit int, CommentType_ string, Langcode_ string, Mail_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and langcode = ? and mail = ?", CommentType_, Langcode_, Mail_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndLangcodeAndHomepage Get CommentFieldDatas via CommentTypeAndLangcodeAndHomepage
func GetCommentFieldDatasByCommentTypeAndLangcodeAndHomepage(offset int, limit int, CommentType_ string, Langcode_ string, Homepage_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and langcode = ? and homepage = ?", CommentType_, Langcode_, Homepage_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndLangcodeAndHostname Get CommentFieldDatas via CommentTypeAndLangcodeAndHostname
func GetCommentFieldDatasByCommentTypeAndLangcodeAndHostname(offset int, limit int, CommentType_ string, Langcode_ string, Hostname_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and langcode = ? and hostname = ?", CommentType_, Langcode_, Hostname_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndLangcodeAndCreated Get CommentFieldDatas via CommentTypeAndLangcodeAndCreated
func GetCommentFieldDatasByCommentTypeAndLangcodeAndCreated(offset int, limit int, CommentType_ string, Langcode_ string, Created_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and langcode = ? and created = ?", CommentType_, Langcode_, Created_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndLangcodeAndChanged Get CommentFieldDatas via CommentTypeAndLangcodeAndChanged
func GetCommentFieldDatasByCommentTypeAndLangcodeAndChanged(offset int, limit int, CommentType_ string, Langcode_ string, Changed_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and langcode = ? and changed = ?", CommentType_, Langcode_, Changed_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndLangcodeAndStatus Get CommentFieldDatas via CommentTypeAndLangcodeAndStatus
func GetCommentFieldDatasByCommentTypeAndLangcodeAndStatus(offset int, limit int, CommentType_ string, Langcode_ string, Status_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and langcode = ? and status = ?", CommentType_, Langcode_, Status_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndLangcodeAndThread Get CommentFieldDatas via CommentTypeAndLangcodeAndThread
func GetCommentFieldDatasByCommentTypeAndLangcodeAndThread(offset int, limit int, CommentType_ string, Langcode_ string, Thread_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and langcode = ? and thread = ?", CommentType_, Langcode_, Thread_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndLangcodeAndEntityType Get CommentFieldDatas via CommentTypeAndLangcodeAndEntityType
func GetCommentFieldDatasByCommentTypeAndLangcodeAndEntityType(offset int, limit int, CommentType_ string, Langcode_ string, EntityType_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and langcode = ? and entity_type = ?", CommentType_, Langcode_, EntityType_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndLangcodeAndFieldName Get CommentFieldDatas via CommentTypeAndLangcodeAndFieldName
func GetCommentFieldDatasByCommentTypeAndLangcodeAndFieldName(offset int, limit int, CommentType_ string, Langcode_ string, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and langcode = ? and field_name = ?", CommentType_, Langcode_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndLangcodeAndDefaultLangcode Get CommentFieldDatas via CommentTypeAndLangcodeAndDefaultLangcode
func GetCommentFieldDatasByCommentTypeAndLangcodeAndDefaultLangcode(offset int, limit int, CommentType_ string, Langcode_ string, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and langcode = ? and default_langcode = ?", CommentType_, Langcode_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndPidAndEntityId Get CommentFieldDatas via CommentTypeAndPidAndEntityId
func GetCommentFieldDatasByCommentTypeAndPidAndEntityId(offset int, limit int, CommentType_ string, Pid_ int, EntityId_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and pid = ? and entity_id = ?", CommentType_, Pid_, EntityId_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndPidAndSubject Get CommentFieldDatas via CommentTypeAndPidAndSubject
func GetCommentFieldDatasByCommentTypeAndPidAndSubject(offset int, limit int, CommentType_ string, Pid_ int, Subject_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and pid = ? and subject = ?", CommentType_, Pid_, Subject_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndPidAndUid Get CommentFieldDatas via CommentTypeAndPidAndUid
func GetCommentFieldDatasByCommentTypeAndPidAndUid(offset int, limit int, CommentType_ string, Pid_ int, Uid_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and pid = ? and uid = ?", CommentType_, Pid_, Uid_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndPidAndName Get CommentFieldDatas via CommentTypeAndPidAndName
func GetCommentFieldDatasByCommentTypeAndPidAndName(offset int, limit int, CommentType_ string, Pid_ int, Name_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and pid = ? and name = ?", CommentType_, Pid_, Name_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndPidAndMail Get CommentFieldDatas via CommentTypeAndPidAndMail
func GetCommentFieldDatasByCommentTypeAndPidAndMail(offset int, limit int, CommentType_ string, Pid_ int, Mail_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and pid = ? and mail = ?", CommentType_, Pid_, Mail_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndPidAndHomepage Get CommentFieldDatas via CommentTypeAndPidAndHomepage
func GetCommentFieldDatasByCommentTypeAndPidAndHomepage(offset int, limit int, CommentType_ string, Pid_ int, Homepage_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and pid = ? and homepage = ?", CommentType_, Pid_, Homepage_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndPidAndHostname Get CommentFieldDatas via CommentTypeAndPidAndHostname
func GetCommentFieldDatasByCommentTypeAndPidAndHostname(offset int, limit int, CommentType_ string, Pid_ int, Hostname_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and pid = ? and hostname = ?", CommentType_, Pid_, Hostname_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndPidAndCreated Get CommentFieldDatas via CommentTypeAndPidAndCreated
func GetCommentFieldDatasByCommentTypeAndPidAndCreated(offset int, limit int, CommentType_ string, Pid_ int, Created_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and pid = ? and created = ?", CommentType_, Pid_, Created_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndPidAndChanged Get CommentFieldDatas via CommentTypeAndPidAndChanged
func GetCommentFieldDatasByCommentTypeAndPidAndChanged(offset int, limit int, CommentType_ string, Pid_ int, Changed_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and pid = ? and changed = ?", CommentType_, Pid_, Changed_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndPidAndStatus Get CommentFieldDatas via CommentTypeAndPidAndStatus
func GetCommentFieldDatasByCommentTypeAndPidAndStatus(offset int, limit int, CommentType_ string, Pid_ int, Status_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and pid = ? and status = ?", CommentType_, Pid_, Status_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndPidAndThread Get CommentFieldDatas via CommentTypeAndPidAndThread
func GetCommentFieldDatasByCommentTypeAndPidAndThread(offset int, limit int, CommentType_ string, Pid_ int, Thread_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and pid = ? and thread = ?", CommentType_, Pid_, Thread_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndPidAndEntityType Get CommentFieldDatas via CommentTypeAndPidAndEntityType
func GetCommentFieldDatasByCommentTypeAndPidAndEntityType(offset int, limit int, CommentType_ string, Pid_ int, EntityType_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and pid = ? and entity_type = ?", CommentType_, Pid_, EntityType_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndPidAndFieldName Get CommentFieldDatas via CommentTypeAndPidAndFieldName
func GetCommentFieldDatasByCommentTypeAndPidAndFieldName(offset int, limit int, CommentType_ string, Pid_ int, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and pid = ? and field_name = ?", CommentType_, Pid_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndPidAndDefaultLangcode Get CommentFieldDatas via CommentTypeAndPidAndDefaultLangcode
func GetCommentFieldDatasByCommentTypeAndPidAndDefaultLangcode(offset int, limit int, CommentType_ string, Pid_ int, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and pid = ? and default_langcode = ?", CommentType_, Pid_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndEntityIdAndSubject Get CommentFieldDatas via CommentTypeAndEntityIdAndSubject
func GetCommentFieldDatasByCommentTypeAndEntityIdAndSubject(offset int, limit int, CommentType_ string, EntityId_ int, Subject_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and entity_id = ? and subject = ?", CommentType_, EntityId_, Subject_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndEntityIdAndUid Get CommentFieldDatas via CommentTypeAndEntityIdAndUid
func GetCommentFieldDatasByCommentTypeAndEntityIdAndUid(offset int, limit int, CommentType_ string, EntityId_ int, Uid_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and entity_id = ? and uid = ?", CommentType_, EntityId_, Uid_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndEntityIdAndName Get CommentFieldDatas via CommentTypeAndEntityIdAndName
func GetCommentFieldDatasByCommentTypeAndEntityIdAndName(offset int, limit int, CommentType_ string, EntityId_ int, Name_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and entity_id = ? and name = ?", CommentType_, EntityId_, Name_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndEntityIdAndMail Get CommentFieldDatas via CommentTypeAndEntityIdAndMail
func GetCommentFieldDatasByCommentTypeAndEntityIdAndMail(offset int, limit int, CommentType_ string, EntityId_ int, Mail_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and entity_id = ? and mail = ?", CommentType_, EntityId_, Mail_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndEntityIdAndHomepage Get CommentFieldDatas via CommentTypeAndEntityIdAndHomepage
func GetCommentFieldDatasByCommentTypeAndEntityIdAndHomepage(offset int, limit int, CommentType_ string, EntityId_ int, Homepage_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and entity_id = ? and homepage = ?", CommentType_, EntityId_, Homepage_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndEntityIdAndHostname Get CommentFieldDatas via CommentTypeAndEntityIdAndHostname
func GetCommentFieldDatasByCommentTypeAndEntityIdAndHostname(offset int, limit int, CommentType_ string, EntityId_ int, Hostname_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and entity_id = ? and hostname = ?", CommentType_, EntityId_, Hostname_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndEntityIdAndCreated Get CommentFieldDatas via CommentTypeAndEntityIdAndCreated
func GetCommentFieldDatasByCommentTypeAndEntityIdAndCreated(offset int, limit int, CommentType_ string, EntityId_ int, Created_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and entity_id = ? and created = ?", CommentType_, EntityId_, Created_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndEntityIdAndChanged Get CommentFieldDatas via CommentTypeAndEntityIdAndChanged
func GetCommentFieldDatasByCommentTypeAndEntityIdAndChanged(offset int, limit int, CommentType_ string, EntityId_ int, Changed_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and entity_id = ? and changed = ?", CommentType_, EntityId_, Changed_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndEntityIdAndStatus Get CommentFieldDatas via CommentTypeAndEntityIdAndStatus
func GetCommentFieldDatasByCommentTypeAndEntityIdAndStatus(offset int, limit int, CommentType_ string, EntityId_ int, Status_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and entity_id = ? and status = ?", CommentType_, EntityId_, Status_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndEntityIdAndThread Get CommentFieldDatas via CommentTypeAndEntityIdAndThread
func GetCommentFieldDatasByCommentTypeAndEntityIdAndThread(offset int, limit int, CommentType_ string, EntityId_ int, Thread_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and entity_id = ? and thread = ?", CommentType_, EntityId_, Thread_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndEntityIdAndEntityType Get CommentFieldDatas via CommentTypeAndEntityIdAndEntityType
func GetCommentFieldDatasByCommentTypeAndEntityIdAndEntityType(offset int, limit int, CommentType_ string, EntityId_ int, EntityType_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and entity_id = ? and entity_type = ?", CommentType_, EntityId_, EntityType_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndEntityIdAndFieldName Get CommentFieldDatas via CommentTypeAndEntityIdAndFieldName
func GetCommentFieldDatasByCommentTypeAndEntityIdAndFieldName(offset int, limit int, CommentType_ string, EntityId_ int, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and entity_id = ? and field_name = ?", CommentType_, EntityId_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndEntityIdAndDefaultLangcode Get CommentFieldDatas via CommentTypeAndEntityIdAndDefaultLangcode
func GetCommentFieldDatasByCommentTypeAndEntityIdAndDefaultLangcode(offset int, limit int, CommentType_ string, EntityId_ int, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and entity_id = ? and default_langcode = ?", CommentType_, EntityId_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndSubjectAndUid Get CommentFieldDatas via CommentTypeAndSubjectAndUid
func GetCommentFieldDatasByCommentTypeAndSubjectAndUid(offset int, limit int, CommentType_ string, Subject_ string, Uid_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and subject = ? and uid = ?", CommentType_, Subject_, Uid_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndSubjectAndName Get CommentFieldDatas via CommentTypeAndSubjectAndName
func GetCommentFieldDatasByCommentTypeAndSubjectAndName(offset int, limit int, CommentType_ string, Subject_ string, Name_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and subject = ? and name = ?", CommentType_, Subject_, Name_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndSubjectAndMail Get CommentFieldDatas via CommentTypeAndSubjectAndMail
func GetCommentFieldDatasByCommentTypeAndSubjectAndMail(offset int, limit int, CommentType_ string, Subject_ string, Mail_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and subject = ? and mail = ?", CommentType_, Subject_, Mail_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndSubjectAndHomepage Get CommentFieldDatas via CommentTypeAndSubjectAndHomepage
func GetCommentFieldDatasByCommentTypeAndSubjectAndHomepage(offset int, limit int, CommentType_ string, Subject_ string, Homepage_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and subject = ? and homepage = ?", CommentType_, Subject_, Homepage_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndSubjectAndHostname Get CommentFieldDatas via CommentTypeAndSubjectAndHostname
func GetCommentFieldDatasByCommentTypeAndSubjectAndHostname(offset int, limit int, CommentType_ string, Subject_ string, Hostname_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and subject = ? and hostname = ?", CommentType_, Subject_, Hostname_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndSubjectAndCreated Get CommentFieldDatas via CommentTypeAndSubjectAndCreated
func GetCommentFieldDatasByCommentTypeAndSubjectAndCreated(offset int, limit int, CommentType_ string, Subject_ string, Created_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and subject = ? and created = ?", CommentType_, Subject_, Created_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndSubjectAndChanged Get CommentFieldDatas via CommentTypeAndSubjectAndChanged
func GetCommentFieldDatasByCommentTypeAndSubjectAndChanged(offset int, limit int, CommentType_ string, Subject_ string, Changed_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and subject = ? and changed = ?", CommentType_, Subject_, Changed_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndSubjectAndStatus Get CommentFieldDatas via CommentTypeAndSubjectAndStatus
func GetCommentFieldDatasByCommentTypeAndSubjectAndStatus(offset int, limit int, CommentType_ string, Subject_ string, Status_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and subject = ? and status = ?", CommentType_, Subject_, Status_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndSubjectAndThread Get CommentFieldDatas via CommentTypeAndSubjectAndThread
func GetCommentFieldDatasByCommentTypeAndSubjectAndThread(offset int, limit int, CommentType_ string, Subject_ string, Thread_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and subject = ? and thread = ?", CommentType_, Subject_, Thread_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndSubjectAndEntityType Get CommentFieldDatas via CommentTypeAndSubjectAndEntityType
func GetCommentFieldDatasByCommentTypeAndSubjectAndEntityType(offset int, limit int, CommentType_ string, Subject_ string, EntityType_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and subject = ? and entity_type = ?", CommentType_, Subject_, EntityType_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndSubjectAndFieldName Get CommentFieldDatas via CommentTypeAndSubjectAndFieldName
func GetCommentFieldDatasByCommentTypeAndSubjectAndFieldName(offset int, limit int, CommentType_ string, Subject_ string, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and subject = ? and field_name = ?", CommentType_, Subject_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndSubjectAndDefaultLangcode Get CommentFieldDatas via CommentTypeAndSubjectAndDefaultLangcode
func GetCommentFieldDatasByCommentTypeAndSubjectAndDefaultLangcode(offset int, limit int, CommentType_ string, Subject_ string, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and subject = ? and default_langcode = ?", CommentType_, Subject_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndUidAndName Get CommentFieldDatas via CommentTypeAndUidAndName
func GetCommentFieldDatasByCommentTypeAndUidAndName(offset int, limit int, CommentType_ string, Uid_ int, Name_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and uid = ? and name = ?", CommentType_, Uid_, Name_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndUidAndMail Get CommentFieldDatas via CommentTypeAndUidAndMail
func GetCommentFieldDatasByCommentTypeAndUidAndMail(offset int, limit int, CommentType_ string, Uid_ int, Mail_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and uid = ? and mail = ?", CommentType_, Uid_, Mail_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndUidAndHomepage Get CommentFieldDatas via CommentTypeAndUidAndHomepage
func GetCommentFieldDatasByCommentTypeAndUidAndHomepage(offset int, limit int, CommentType_ string, Uid_ int, Homepage_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and uid = ? and homepage = ?", CommentType_, Uid_, Homepage_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndUidAndHostname Get CommentFieldDatas via CommentTypeAndUidAndHostname
func GetCommentFieldDatasByCommentTypeAndUidAndHostname(offset int, limit int, CommentType_ string, Uid_ int, Hostname_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and uid = ? and hostname = ?", CommentType_, Uid_, Hostname_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndUidAndCreated Get CommentFieldDatas via CommentTypeAndUidAndCreated
func GetCommentFieldDatasByCommentTypeAndUidAndCreated(offset int, limit int, CommentType_ string, Uid_ int, Created_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and uid = ? and created = ?", CommentType_, Uid_, Created_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndUidAndChanged Get CommentFieldDatas via CommentTypeAndUidAndChanged
func GetCommentFieldDatasByCommentTypeAndUidAndChanged(offset int, limit int, CommentType_ string, Uid_ int, Changed_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and uid = ? and changed = ?", CommentType_, Uid_, Changed_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndUidAndStatus Get CommentFieldDatas via CommentTypeAndUidAndStatus
func GetCommentFieldDatasByCommentTypeAndUidAndStatus(offset int, limit int, CommentType_ string, Uid_ int, Status_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and uid = ? and status = ?", CommentType_, Uid_, Status_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndUidAndThread Get CommentFieldDatas via CommentTypeAndUidAndThread
func GetCommentFieldDatasByCommentTypeAndUidAndThread(offset int, limit int, CommentType_ string, Uid_ int, Thread_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and uid = ? and thread = ?", CommentType_, Uid_, Thread_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndUidAndEntityType Get CommentFieldDatas via CommentTypeAndUidAndEntityType
func GetCommentFieldDatasByCommentTypeAndUidAndEntityType(offset int, limit int, CommentType_ string, Uid_ int, EntityType_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and uid = ? and entity_type = ?", CommentType_, Uid_, EntityType_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndUidAndFieldName Get CommentFieldDatas via CommentTypeAndUidAndFieldName
func GetCommentFieldDatasByCommentTypeAndUidAndFieldName(offset int, limit int, CommentType_ string, Uid_ int, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and uid = ? and field_name = ?", CommentType_, Uid_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndUidAndDefaultLangcode Get CommentFieldDatas via CommentTypeAndUidAndDefaultLangcode
func GetCommentFieldDatasByCommentTypeAndUidAndDefaultLangcode(offset int, limit int, CommentType_ string, Uid_ int, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and uid = ? and default_langcode = ?", CommentType_, Uid_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndNameAndMail Get CommentFieldDatas via CommentTypeAndNameAndMail
func GetCommentFieldDatasByCommentTypeAndNameAndMail(offset int, limit int, CommentType_ string, Name_ string, Mail_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and name = ? and mail = ?", CommentType_, Name_, Mail_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndNameAndHomepage Get CommentFieldDatas via CommentTypeAndNameAndHomepage
func GetCommentFieldDatasByCommentTypeAndNameAndHomepage(offset int, limit int, CommentType_ string, Name_ string, Homepage_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and name = ? and homepage = ?", CommentType_, Name_, Homepage_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndNameAndHostname Get CommentFieldDatas via CommentTypeAndNameAndHostname
func GetCommentFieldDatasByCommentTypeAndNameAndHostname(offset int, limit int, CommentType_ string, Name_ string, Hostname_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and name = ? and hostname = ?", CommentType_, Name_, Hostname_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndNameAndCreated Get CommentFieldDatas via CommentTypeAndNameAndCreated
func GetCommentFieldDatasByCommentTypeAndNameAndCreated(offset int, limit int, CommentType_ string, Name_ string, Created_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and name = ? and created = ?", CommentType_, Name_, Created_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndNameAndChanged Get CommentFieldDatas via CommentTypeAndNameAndChanged
func GetCommentFieldDatasByCommentTypeAndNameAndChanged(offset int, limit int, CommentType_ string, Name_ string, Changed_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and name = ? and changed = ?", CommentType_, Name_, Changed_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndNameAndStatus Get CommentFieldDatas via CommentTypeAndNameAndStatus
func GetCommentFieldDatasByCommentTypeAndNameAndStatus(offset int, limit int, CommentType_ string, Name_ string, Status_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and name = ? and status = ?", CommentType_, Name_, Status_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndNameAndThread Get CommentFieldDatas via CommentTypeAndNameAndThread
func GetCommentFieldDatasByCommentTypeAndNameAndThread(offset int, limit int, CommentType_ string, Name_ string, Thread_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and name = ? and thread = ?", CommentType_, Name_, Thread_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndNameAndEntityType Get CommentFieldDatas via CommentTypeAndNameAndEntityType
func GetCommentFieldDatasByCommentTypeAndNameAndEntityType(offset int, limit int, CommentType_ string, Name_ string, EntityType_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and name = ? and entity_type = ?", CommentType_, Name_, EntityType_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndNameAndFieldName Get CommentFieldDatas via CommentTypeAndNameAndFieldName
func GetCommentFieldDatasByCommentTypeAndNameAndFieldName(offset int, limit int, CommentType_ string, Name_ string, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and name = ? and field_name = ?", CommentType_, Name_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndNameAndDefaultLangcode Get CommentFieldDatas via CommentTypeAndNameAndDefaultLangcode
func GetCommentFieldDatasByCommentTypeAndNameAndDefaultLangcode(offset int, limit int, CommentType_ string, Name_ string, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and name = ? and default_langcode = ?", CommentType_, Name_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndMailAndHomepage Get CommentFieldDatas via CommentTypeAndMailAndHomepage
func GetCommentFieldDatasByCommentTypeAndMailAndHomepage(offset int, limit int, CommentType_ string, Mail_ string, Homepage_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and mail = ? and homepage = ?", CommentType_, Mail_, Homepage_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndMailAndHostname Get CommentFieldDatas via CommentTypeAndMailAndHostname
func GetCommentFieldDatasByCommentTypeAndMailAndHostname(offset int, limit int, CommentType_ string, Mail_ string, Hostname_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and mail = ? and hostname = ?", CommentType_, Mail_, Hostname_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndMailAndCreated Get CommentFieldDatas via CommentTypeAndMailAndCreated
func GetCommentFieldDatasByCommentTypeAndMailAndCreated(offset int, limit int, CommentType_ string, Mail_ string, Created_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and mail = ? and created = ?", CommentType_, Mail_, Created_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndMailAndChanged Get CommentFieldDatas via CommentTypeAndMailAndChanged
func GetCommentFieldDatasByCommentTypeAndMailAndChanged(offset int, limit int, CommentType_ string, Mail_ string, Changed_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and mail = ? and changed = ?", CommentType_, Mail_, Changed_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndMailAndStatus Get CommentFieldDatas via CommentTypeAndMailAndStatus
func GetCommentFieldDatasByCommentTypeAndMailAndStatus(offset int, limit int, CommentType_ string, Mail_ string, Status_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and mail = ? and status = ?", CommentType_, Mail_, Status_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndMailAndThread Get CommentFieldDatas via CommentTypeAndMailAndThread
func GetCommentFieldDatasByCommentTypeAndMailAndThread(offset int, limit int, CommentType_ string, Mail_ string, Thread_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and mail = ? and thread = ?", CommentType_, Mail_, Thread_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndMailAndEntityType Get CommentFieldDatas via CommentTypeAndMailAndEntityType
func GetCommentFieldDatasByCommentTypeAndMailAndEntityType(offset int, limit int, CommentType_ string, Mail_ string, EntityType_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and mail = ? and entity_type = ?", CommentType_, Mail_, EntityType_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndMailAndFieldName Get CommentFieldDatas via CommentTypeAndMailAndFieldName
func GetCommentFieldDatasByCommentTypeAndMailAndFieldName(offset int, limit int, CommentType_ string, Mail_ string, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and mail = ? and field_name = ?", CommentType_, Mail_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndMailAndDefaultLangcode Get CommentFieldDatas via CommentTypeAndMailAndDefaultLangcode
func GetCommentFieldDatasByCommentTypeAndMailAndDefaultLangcode(offset int, limit int, CommentType_ string, Mail_ string, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and mail = ? and default_langcode = ?", CommentType_, Mail_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndHomepageAndHostname Get CommentFieldDatas via CommentTypeAndHomepageAndHostname
func GetCommentFieldDatasByCommentTypeAndHomepageAndHostname(offset int, limit int, CommentType_ string, Homepage_ string, Hostname_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and homepage = ? and hostname = ?", CommentType_, Homepage_, Hostname_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndHomepageAndCreated Get CommentFieldDatas via CommentTypeAndHomepageAndCreated
func GetCommentFieldDatasByCommentTypeAndHomepageAndCreated(offset int, limit int, CommentType_ string, Homepage_ string, Created_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and homepage = ? and created = ?", CommentType_, Homepage_, Created_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndHomepageAndChanged Get CommentFieldDatas via CommentTypeAndHomepageAndChanged
func GetCommentFieldDatasByCommentTypeAndHomepageAndChanged(offset int, limit int, CommentType_ string, Homepage_ string, Changed_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and homepage = ? and changed = ?", CommentType_, Homepage_, Changed_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndHomepageAndStatus Get CommentFieldDatas via CommentTypeAndHomepageAndStatus
func GetCommentFieldDatasByCommentTypeAndHomepageAndStatus(offset int, limit int, CommentType_ string, Homepage_ string, Status_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and homepage = ? and status = ?", CommentType_, Homepage_, Status_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndHomepageAndThread Get CommentFieldDatas via CommentTypeAndHomepageAndThread
func GetCommentFieldDatasByCommentTypeAndHomepageAndThread(offset int, limit int, CommentType_ string, Homepage_ string, Thread_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and homepage = ? and thread = ?", CommentType_, Homepage_, Thread_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndHomepageAndEntityType Get CommentFieldDatas via CommentTypeAndHomepageAndEntityType
func GetCommentFieldDatasByCommentTypeAndHomepageAndEntityType(offset int, limit int, CommentType_ string, Homepage_ string, EntityType_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and homepage = ? and entity_type = ?", CommentType_, Homepage_, EntityType_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndHomepageAndFieldName Get CommentFieldDatas via CommentTypeAndHomepageAndFieldName
func GetCommentFieldDatasByCommentTypeAndHomepageAndFieldName(offset int, limit int, CommentType_ string, Homepage_ string, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and homepage = ? and field_name = ?", CommentType_, Homepage_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndHomepageAndDefaultLangcode Get CommentFieldDatas via CommentTypeAndHomepageAndDefaultLangcode
func GetCommentFieldDatasByCommentTypeAndHomepageAndDefaultLangcode(offset int, limit int, CommentType_ string, Homepage_ string, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and homepage = ? and default_langcode = ?", CommentType_, Homepage_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndHostnameAndCreated Get CommentFieldDatas via CommentTypeAndHostnameAndCreated
func GetCommentFieldDatasByCommentTypeAndHostnameAndCreated(offset int, limit int, CommentType_ string, Hostname_ string, Created_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and hostname = ? and created = ?", CommentType_, Hostname_, Created_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndHostnameAndChanged Get CommentFieldDatas via CommentTypeAndHostnameAndChanged
func GetCommentFieldDatasByCommentTypeAndHostnameAndChanged(offset int, limit int, CommentType_ string, Hostname_ string, Changed_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and hostname = ? and changed = ?", CommentType_, Hostname_, Changed_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndHostnameAndStatus Get CommentFieldDatas via CommentTypeAndHostnameAndStatus
func GetCommentFieldDatasByCommentTypeAndHostnameAndStatus(offset int, limit int, CommentType_ string, Hostname_ string, Status_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and hostname = ? and status = ?", CommentType_, Hostname_, Status_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndHostnameAndThread Get CommentFieldDatas via CommentTypeAndHostnameAndThread
func GetCommentFieldDatasByCommentTypeAndHostnameAndThread(offset int, limit int, CommentType_ string, Hostname_ string, Thread_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and hostname = ? and thread = ?", CommentType_, Hostname_, Thread_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndHostnameAndEntityType Get CommentFieldDatas via CommentTypeAndHostnameAndEntityType
func GetCommentFieldDatasByCommentTypeAndHostnameAndEntityType(offset int, limit int, CommentType_ string, Hostname_ string, EntityType_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and hostname = ? and entity_type = ?", CommentType_, Hostname_, EntityType_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndHostnameAndFieldName Get CommentFieldDatas via CommentTypeAndHostnameAndFieldName
func GetCommentFieldDatasByCommentTypeAndHostnameAndFieldName(offset int, limit int, CommentType_ string, Hostname_ string, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and hostname = ? and field_name = ?", CommentType_, Hostname_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndHostnameAndDefaultLangcode Get CommentFieldDatas via CommentTypeAndHostnameAndDefaultLangcode
func GetCommentFieldDatasByCommentTypeAndHostnameAndDefaultLangcode(offset int, limit int, CommentType_ string, Hostname_ string, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and hostname = ? and default_langcode = ?", CommentType_, Hostname_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndCreatedAndChanged Get CommentFieldDatas via CommentTypeAndCreatedAndChanged
func GetCommentFieldDatasByCommentTypeAndCreatedAndChanged(offset int, limit int, CommentType_ string, Created_ int, Changed_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and created = ? and changed = ?", CommentType_, Created_, Changed_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndCreatedAndStatus Get CommentFieldDatas via CommentTypeAndCreatedAndStatus
func GetCommentFieldDatasByCommentTypeAndCreatedAndStatus(offset int, limit int, CommentType_ string, Created_ int, Status_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and created = ? and status = ?", CommentType_, Created_, Status_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndCreatedAndThread Get CommentFieldDatas via CommentTypeAndCreatedAndThread
func GetCommentFieldDatasByCommentTypeAndCreatedAndThread(offset int, limit int, CommentType_ string, Created_ int, Thread_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and created = ? and thread = ?", CommentType_, Created_, Thread_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndCreatedAndEntityType Get CommentFieldDatas via CommentTypeAndCreatedAndEntityType
func GetCommentFieldDatasByCommentTypeAndCreatedAndEntityType(offset int, limit int, CommentType_ string, Created_ int, EntityType_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and created = ? and entity_type = ?", CommentType_, Created_, EntityType_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndCreatedAndFieldName Get CommentFieldDatas via CommentTypeAndCreatedAndFieldName
func GetCommentFieldDatasByCommentTypeAndCreatedAndFieldName(offset int, limit int, CommentType_ string, Created_ int, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and created = ? and field_name = ?", CommentType_, Created_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndCreatedAndDefaultLangcode Get CommentFieldDatas via CommentTypeAndCreatedAndDefaultLangcode
func GetCommentFieldDatasByCommentTypeAndCreatedAndDefaultLangcode(offset int, limit int, CommentType_ string, Created_ int, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and created = ? and default_langcode = ?", CommentType_, Created_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndChangedAndStatus Get CommentFieldDatas via CommentTypeAndChangedAndStatus
func GetCommentFieldDatasByCommentTypeAndChangedAndStatus(offset int, limit int, CommentType_ string, Changed_ int, Status_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and changed = ? and status = ?", CommentType_, Changed_, Status_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndChangedAndThread Get CommentFieldDatas via CommentTypeAndChangedAndThread
func GetCommentFieldDatasByCommentTypeAndChangedAndThread(offset int, limit int, CommentType_ string, Changed_ int, Thread_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and changed = ? and thread = ?", CommentType_, Changed_, Thread_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndChangedAndEntityType Get CommentFieldDatas via CommentTypeAndChangedAndEntityType
func GetCommentFieldDatasByCommentTypeAndChangedAndEntityType(offset int, limit int, CommentType_ string, Changed_ int, EntityType_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and changed = ? and entity_type = ?", CommentType_, Changed_, EntityType_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndChangedAndFieldName Get CommentFieldDatas via CommentTypeAndChangedAndFieldName
func GetCommentFieldDatasByCommentTypeAndChangedAndFieldName(offset int, limit int, CommentType_ string, Changed_ int, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and changed = ? and field_name = ?", CommentType_, Changed_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndChangedAndDefaultLangcode Get CommentFieldDatas via CommentTypeAndChangedAndDefaultLangcode
func GetCommentFieldDatasByCommentTypeAndChangedAndDefaultLangcode(offset int, limit int, CommentType_ string, Changed_ int, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and changed = ? and default_langcode = ?", CommentType_, Changed_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndStatusAndThread Get CommentFieldDatas via CommentTypeAndStatusAndThread
func GetCommentFieldDatasByCommentTypeAndStatusAndThread(offset int, limit int, CommentType_ string, Status_ int, Thread_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and status = ? and thread = ?", CommentType_, Status_, Thread_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndStatusAndEntityType Get CommentFieldDatas via CommentTypeAndStatusAndEntityType
func GetCommentFieldDatasByCommentTypeAndStatusAndEntityType(offset int, limit int, CommentType_ string, Status_ int, EntityType_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and status = ? and entity_type = ?", CommentType_, Status_, EntityType_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndStatusAndFieldName Get CommentFieldDatas via CommentTypeAndStatusAndFieldName
func GetCommentFieldDatasByCommentTypeAndStatusAndFieldName(offset int, limit int, CommentType_ string, Status_ int, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and status = ? and field_name = ?", CommentType_, Status_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndStatusAndDefaultLangcode Get CommentFieldDatas via CommentTypeAndStatusAndDefaultLangcode
func GetCommentFieldDatasByCommentTypeAndStatusAndDefaultLangcode(offset int, limit int, CommentType_ string, Status_ int, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and status = ? and default_langcode = ?", CommentType_, Status_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndThreadAndEntityType Get CommentFieldDatas via CommentTypeAndThreadAndEntityType
func GetCommentFieldDatasByCommentTypeAndThreadAndEntityType(offset int, limit int, CommentType_ string, Thread_ string, EntityType_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and thread = ? and entity_type = ?", CommentType_, Thread_, EntityType_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndThreadAndFieldName Get CommentFieldDatas via CommentTypeAndThreadAndFieldName
func GetCommentFieldDatasByCommentTypeAndThreadAndFieldName(offset int, limit int, CommentType_ string, Thread_ string, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and thread = ? and field_name = ?", CommentType_, Thread_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndThreadAndDefaultLangcode Get CommentFieldDatas via CommentTypeAndThreadAndDefaultLangcode
func GetCommentFieldDatasByCommentTypeAndThreadAndDefaultLangcode(offset int, limit int, CommentType_ string, Thread_ string, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and thread = ? and default_langcode = ?", CommentType_, Thread_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndEntityTypeAndFieldName Get CommentFieldDatas via CommentTypeAndEntityTypeAndFieldName
func GetCommentFieldDatasByCommentTypeAndEntityTypeAndFieldName(offset int, limit int, CommentType_ string, EntityType_ string, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and entity_type = ? and field_name = ?", CommentType_, EntityType_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndEntityTypeAndDefaultLangcode Get CommentFieldDatas via CommentTypeAndEntityTypeAndDefaultLangcode
func GetCommentFieldDatasByCommentTypeAndEntityTypeAndDefaultLangcode(offset int, limit int, CommentType_ string, EntityType_ string, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and entity_type = ? and default_langcode = ?", CommentType_, EntityType_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndFieldNameAndDefaultLangcode Get CommentFieldDatas via CommentTypeAndFieldNameAndDefaultLangcode
func GetCommentFieldDatasByCommentTypeAndFieldNameAndDefaultLangcode(offset int, limit int, CommentType_ string, FieldName_ string, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and field_name = ? and default_langcode = ?", CommentType_, FieldName_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByLangcodeAndPidAndEntityId Get CommentFieldDatas via LangcodeAndPidAndEntityId
func GetCommentFieldDatasByLangcodeAndPidAndEntityId(offset int, limit int, Langcode_ string, Pid_ int, EntityId_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("langcode = ? and pid = ? and entity_id = ?", Langcode_, Pid_, EntityId_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByLangcodeAndPidAndSubject Get CommentFieldDatas via LangcodeAndPidAndSubject
func GetCommentFieldDatasByLangcodeAndPidAndSubject(offset int, limit int, Langcode_ string, Pid_ int, Subject_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("langcode = ? and pid = ? and subject = ?", Langcode_, Pid_, Subject_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByLangcodeAndPidAndUid Get CommentFieldDatas via LangcodeAndPidAndUid
func GetCommentFieldDatasByLangcodeAndPidAndUid(offset int, limit int, Langcode_ string, Pid_ int, Uid_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("langcode = ? and pid = ? and uid = ?", Langcode_, Pid_, Uid_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByLangcodeAndPidAndName Get CommentFieldDatas via LangcodeAndPidAndName
func GetCommentFieldDatasByLangcodeAndPidAndName(offset int, limit int, Langcode_ string, Pid_ int, Name_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("langcode = ? and pid = ? and name = ?", Langcode_, Pid_, Name_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByLangcodeAndPidAndMail Get CommentFieldDatas via LangcodeAndPidAndMail
func GetCommentFieldDatasByLangcodeAndPidAndMail(offset int, limit int, Langcode_ string, Pid_ int, Mail_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("langcode = ? and pid = ? and mail = ?", Langcode_, Pid_, Mail_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByLangcodeAndPidAndHomepage Get CommentFieldDatas via LangcodeAndPidAndHomepage
func GetCommentFieldDatasByLangcodeAndPidAndHomepage(offset int, limit int, Langcode_ string, Pid_ int, Homepage_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("langcode = ? and pid = ? and homepage = ?", Langcode_, Pid_, Homepage_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByLangcodeAndPidAndHostname Get CommentFieldDatas via LangcodeAndPidAndHostname
func GetCommentFieldDatasByLangcodeAndPidAndHostname(offset int, limit int, Langcode_ string, Pid_ int, Hostname_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("langcode = ? and pid = ? and hostname = ?", Langcode_, Pid_, Hostname_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByLangcodeAndPidAndCreated Get CommentFieldDatas via LangcodeAndPidAndCreated
func GetCommentFieldDatasByLangcodeAndPidAndCreated(offset int, limit int, Langcode_ string, Pid_ int, Created_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("langcode = ? and pid = ? and created = ?", Langcode_, Pid_, Created_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByLangcodeAndPidAndChanged Get CommentFieldDatas via LangcodeAndPidAndChanged
func GetCommentFieldDatasByLangcodeAndPidAndChanged(offset int, limit int, Langcode_ string, Pid_ int, Changed_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("langcode = ? and pid = ? and changed = ?", Langcode_, Pid_, Changed_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByLangcodeAndPidAndStatus Get CommentFieldDatas via LangcodeAndPidAndStatus
func GetCommentFieldDatasByLangcodeAndPidAndStatus(offset int, limit int, Langcode_ string, Pid_ int, Status_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("langcode = ? and pid = ? and status = ?", Langcode_, Pid_, Status_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByLangcodeAndPidAndThread Get CommentFieldDatas via LangcodeAndPidAndThread
func GetCommentFieldDatasByLangcodeAndPidAndThread(offset int, limit int, Langcode_ string, Pid_ int, Thread_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("langcode = ? and pid = ? and thread = ?", Langcode_, Pid_, Thread_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByLangcodeAndPidAndEntityType Get CommentFieldDatas via LangcodeAndPidAndEntityType
func GetCommentFieldDatasByLangcodeAndPidAndEntityType(offset int, limit int, Langcode_ string, Pid_ int, EntityType_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("langcode = ? and pid = ? and entity_type = ?", Langcode_, Pid_, EntityType_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByLangcodeAndPidAndFieldName Get CommentFieldDatas via LangcodeAndPidAndFieldName
func GetCommentFieldDatasByLangcodeAndPidAndFieldName(offset int, limit int, Langcode_ string, Pid_ int, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("langcode = ? and pid = ? and field_name = ?", Langcode_, Pid_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByLangcodeAndPidAndDefaultLangcode Get CommentFieldDatas via LangcodeAndPidAndDefaultLangcode
func GetCommentFieldDatasByLangcodeAndPidAndDefaultLangcode(offset int, limit int, Langcode_ string, Pid_ int, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("langcode = ? and pid = ? and default_langcode = ?", Langcode_, Pid_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByLangcodeAndEntityIdAndSubject Get CommentFieldDatas via LangcodeAndEntityIdAndSubject
func GetCommentFieldDatasByLangcodeAndEntityIdAndSubject(offset int, limit int, Langcode_ string, EntityId_ int, Subject_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("langcode = ? and entity_id = ? and subject = ?", Langcode_, EntityId_, Subject_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByLangcodeAndEntityIdAndUid Get CommentFieldDatas via LangcodeAndEntityIdAndUid
func GetCommentFieldDatasByLangcodeAndEntityIdAndUid(offset int, limit int, Langcode_ string, EntityId_ int, Uid_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("langcode = ? and entity_id = ? and uid = ?", Langcode_, EntityId_, Uid_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByLangcodeAndEntityIdAndName Get CommentFieldDatas via LangcodeAndEntityIdAndName
func GetCommentFieldDatasByLangcodeAndEntityIdAndName(offset int, limit int, Langcode_ string, EntityId_ int, Name_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("langcode = ? and entity_id = ? and name = ?", Langcode_, EntityId_, Name_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByLangcodeAndEntityIdAndMail Get CommentFieldDatas via LangcodeAndEntityIdAndMail
func GetCommentFieldDatasByLangcodeAndEntityIdAndMail(offset int, limit int, Langcode_ string, EntityId_ int, Mail_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("langcode = ? and entity_id = ? and mail = ?", Langcode_, EntityId_, Mail_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByLangcodeAndEntityIdAndHomepage Get CommentFieldDatas via LangcodeAndEntityIdAndHomepage
func GetCommentFieldDatasByLangcodeAndEntityIdAndHomepage(offset int, limit int, Langcode_ string, EntityId_ int, Homepage_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("langcode = ? and entity_id = ? and homepage = ?", Langcode_, EntityId_, Homepage_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByLangcodeAndEntityIdAndHostname Get CommentFieldDatas via LangcodeAndEntityIdAndHostname
func GetCommentFieldDatasByLangcodeAndEntityIdAndHostname(offset int, limit int, Langcode_ string, EntityId_ int, Hostname_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("langcode = ? and entity_id = ? and hostname = ?", Langcode_, EntityId_, Hostname_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByLangcodeAndEntityIdAndCreated Get CommentFieldDatas via LangcodeAndEntityIdAndCreated
func GetCommentFieldDatasByLangcodeAndEntityIdAndCreated(offset int, limit int, Langcode_ string, EntityId_ int, Created_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("langcode = ? and entity_id = ? and created = ?", Langcode_, EntityId_, Created_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByLangcodeAndEntityIdAndChanged Get CommentFieldDatas via LangcodeAndEntityIdAndChanged
func GetCommentFieldDatasByLangcodeAndEntityIdAndChanged(offset int, limit int, Langcode_ string, EntityId_ int, Changed_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("langcode = ? and entity_id = ? and changed = ?", Langcode_, EntityId_, Changed_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByLangcodeAndEntityIdAndStatus Get CommentFieldDatas via LangcodeAndEntityIdAndStatus
func GetCommentFieldDatasByLangcodeAndEntityIdAndStatus(offset int, limit int, Langcode_ string, EntityId_ int, Status_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("langcode = ? and entity_id = ? and status = ?", Langcode_, EntityId_, Status_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByLangcodeAndEntityIdAndThread Get CommentFieldDatas via LangcodeAndEntityIdAndThread
func GetCommentFieldDatasByLangcodeAndEntityIdAndThread(offset int, limit int, Langcode_ string, EntityId_ int, Thread_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("langcode = ? and entity_id = ? and thread = ?", Langcode_, EntityId_, Thread_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByLangcodeAndEntityIdAndEntityType Get CommentFieldDatas via LangcodeAndEntityIdAndEntityType
func GetCommentFieldDatasByLangcodeAndEntityIdAndEntityType(offset int, limit int, Langcode_ string, EntityId_ int, EntityType_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("langcode = ? and entity_id = ? and entity_type = ?", Langcode_, EntityId_, EntityType_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByLangcodeAndEntityIdAndFieldName Get CommentFieldDatas via LangcodeAndEntityIdAndFieldName
func GetCommentFieldDatasByLangcodeAndEntityIdAndFieldName(offset int, limit int, Langcode_ string, EntityId_ int, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("langcode = ? and entity_id = ? and field_name = ?", Langcode_, EntityId_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByLangcodeAndEntityIdAndDefaultLangcode Get CommentFieldDatas via LangcodeAndEntityIdAndDefaultLangcode
func GetCommentFieldDatasByLangcodeAndEntityIdAndDefaultLangcode(offset int, limit int, Langcode_ string, EntityId_ int, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("langcode = ? and entity_id = ? and default_langcode = ?", Langcode_, EntityId_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByLangcodeAndSubjectAndUid Get CommentFieldDatas via LangcodeAndSubjectAndUid
func GetCommentFieldDatasByLangcodeAndSubjectAndUid(offset int, limit int, Langcode_ string, Subject_ string, Uid_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("langcode = ? and subject = ? and uid = ?", Langcode_, Subject_, Uid_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByLangcodeAndSubjectAndName Get CommentFieldDatas via LangcodeAndSubjectAndName
func GetCommentFieldDatasByLangcodeAndSubjectAndName(offset int, limit int, Langcode_ string, Subject_ string, Name_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("langcode = ? and subject = ? and name = ?", Langcode_, Subject_, Name_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByLangcodeAndSubjectAndMail Get CommentFieldDatas via LangcodeAndSubjectAndMail
func GetCommentFieldDatasByLangcodeAndSubjectAndMail(offset int, limit int, Langcode_ string, Subject_ string, Mail_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("langcode = ? and subject = ? and mail = ?", Langcode_, Subject_, Mail_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByLangcodeAndSubjectAndHomepage Get CommentFieldDatas via LangcodeAndSubjectAndHomepage
func GetCommentFieldDatasByLangcodeAndSubjectAndHomepage(offset int, limit int, Langcode_ string, Subject_ string, Homepage_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("langcode = ? and subject = ? and homepage = ?", Langcode_, Subject_, Homepage_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByLangcodeAndSubjectAndHostname Get CommentFieldDatas via LangcodeAndSubjectAndHostname
func GetCommentFieldDatasByLangcodeAndSubjectAndHostname(offset int, limit int, Langcode_ string, Subject_ string, Hostname_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("langcode = ? and subject = ? and hostname = ?", Langcode_, Subject_, Hostname_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByLangcodeAndSubjectAndCreated Get CommentFieldDatas via LangcodeAndSubjectAndCreated
func GetCommentFieldDatasByLangcodeAndSubjectAndCreated(offset int, limit int, Langcode_ string, Subject_ string, Created_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("langcode = ? and subject = ? and created = ?", Langcode_, Subject_, Created_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByLangcodeAndSubjectAndChanged Get CommentFieldDatas via LangcodeAndSubjectAndChanged
func GetCommentFieldDatasByLangcodeAndSubjectAndChanged(offset int, limit int, Langcode_ string, Subject_ string, Changed_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("langcode = ? and subject = ? and changed = ?", Langcode_, Subject_, Changed_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByLangcodeAndSubjectAndStatus Get CommentFieldDatas via LangcodeAndSubjectAndStatus
func GetCommentFieldDatasByLangcodeAndSubjectAndStatus(offset int, limit int, Langcode_ string, Subject_ string, Status_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("langcode = ? and subject = ? and status = ?", Langcode_, Subject_, Status_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByLangcodeAndSubjectAndThread Get CommentFieldDatas via LangcodeAndSubjectAndThread
func GetCommentFieldDatasByLangcodeAndSubjectAndThread(offset int, limit int, Langcode_ string, Subject_ string, Thread_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("langcode = ? and subject = ? and thread = ?", Langcode_, Subject_, Thread_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByLangcodeAndSubjectAndEntityType Get CommentFieldDatas via LangcodeAndSubjectAndEntityType
func GetCommentFieldDatasByLangcodeAndSubjectAndEntityType(offset int, limit int, Langcode_ string, Subject_ string, EntityType_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("langcode = ? and subject = ? and entity_type = ?", Langcode_, Subject_, EntityType_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByLangcodeAndSubjectAndFieldName Get CommentFieldDatas via LangcodeAndSubjectAndFieldName
func GetCommentFieldDatasByLangcodeAndSubjectAndFieldName(offset int, limit int, Langcode_ string, Subject_ string, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("langcode = ? and subject = ? and field_name = ?", Langcode_, Subject_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByLangcodeAndSubjectAndDefaultLangcode Get CommentFieldDatas via LangcodeAndSubjectAndDefaultLangcode
func GetCommentFieldDatasByLangcodeAndSubjectAndDefaultLangcode(offset int, limit int, Langcode_ string, Subject_ string, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("langcode = ? and subject = ? and default_langcode = ?", Langcode_, Subject_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByLangcodeAndUidAndName Get CommentFieldDatas via LangcodeAndUidAndName
func GetCommentFieldDatasByLangcodeAndUidAndName(offset int, limit int, Langcode_ string, Uid_ int, Name_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("langcode = ? and uid = ? and name = ?", Langcode_, Uid_, Name_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByLangcodeAndUidAndMail Get CommentFieldDatas via LangcodeAndUidAndMail
func GetCommentFieldDatasByLangcodeAndUidAndMail(offset int, limit int, Langcode_ string, Uid_ int, Mail_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("langcode = ? and uid = ? and mail = ?", Langcode_, Uid_, Mail_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByLangcodeAndUidAndHomepage Get CommentFieldDatas via LangcodeAndUidAndHomepage
func GetCommentFieldDatasByLangcodeAndUidAndHomepage(offset int, limit int, Langcode_ string, Uid_ int, Homepage_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("langcode = ? and uid = ? and homepage = ?", Langcode_, Uid_, Homepage_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByLangcodeAndUidAndHostname Get CommentFieldDatas via LangcodeAndUidAndHostname
func GetCommentFieldDatasByLangcodeAndUidAndHostname(offset int, limit int, Langcode_ string, Uid_ int, Hostname_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("langcode = ? and uid = ? and hostname = ?", Langcode_, Uid_, Hostname_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByLangcodeAndUidAndCreated Get CommentFieldDatas via LangcodeAndUidAndCreated
func GetCommentFieldDatasByLangcodeAndUidAndCreated(offset int, limit int, Langcode_ string, Uid_ int, Created_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("langcode = ? and uid = ? and created = ?", Langcode_, Uid_, Created_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByLangcodeAndUidAndChanged Get CommentFieldDatas via LangcodeAndUidAndChanged
func GetCommentFieldDatasByLangcodeAndUidAndChanged(offset int, limit int, Langcode_ string, Uid_ int, Changed_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("langcode = ? and uid = ? and changed = ?", Langcode_, Uid_, Changed_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByLangcodeAndUidAndStatus Get CommentFieldDatas via LangcodeAndUidAndStatus
func GetCommentFieldDatasByLangcodeAndUidAndStatus(offset int, limit int, Langcode_ string, Uid_ int, Status_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("langcode = ? and uid = ? and status = ?", Langcode_, Uid_, Status_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByLangcodeAndUidAndThread Get CommentFieldDatas via LangcodeAndUidAndThread
func GetCommentFieldDatasByLangcodeAndUidAndThread(offset int, limit int, Langcode_ string, Uid_ int, Thread_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("langcode = ? and uid = ? and thread = ?", Langcode_, Uid_, Thread_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByLangcodeAndUidAndEntityType Get CommentFieldDatas via LangcodeAndUidAndEntityType
func GetCommentFieldDatasByLangcodeAndUidAndEntityType(offset int, limit int, Langcode_ string, Uid_ int, EntityType_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("langcode = ? and uid = ? and entity_type = ?", Langcode_, Uid_, EntityType_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByLangcodeAndUidAndFieldName Get CommentFieldDatas via LangcodeAndUidAndFieldName
func GetCommentFieldDatasByLangcodeAndUidAndFieldName(offset int, limit int, Langcode_ string, Uid_ int, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("langcode = ? and uid = ? and field_name = ?", Langcode_, Uid_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByLangcodeAndUidAndDefaultLangcode Get CommentFieldDatas via LangcodeAndUidAndDefaultLangcode
func GetCommentFieldDatasByLangcodeAndUidAndDefaultLangcode(offset int, limit int, Langcode_ string, Uid_ int, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("langcode = ? and uid = ? and default_langcode = ?", Langcode_, Uid_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByLangcodeAndNameAndMail Get CommentFieldDatas via LangcodeAndNameAndMail
func GetCommentFieldDatasByLangcodeAndNameAndMail(offset int, limit int, Langcode_ string, Name_ string, Mail_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("langcode = ? and name = ? and mail = ?", Langcode_, Name_, Mail_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByLangcodeAndNameAndHomepage Get CommentFieldDatas via LangcodeAndNameAndHomepage
func GetCommentFieldDatasByLangcodeAndNameAndHomepage(offset int, limit int, Langcode_ string, Name_ string, Homepage_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("langcode = ? and name = ? and homepage = ?", Langcode_, Name_, Homepage_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByLangcodeAndNameAndHostname Get CommentFieldDatas via LangcodeAndNameAndHostname
func GetCommentFieldDatasByLangcodeAndNameAndHostname(offset int, limit int, Langcode_ string, Name_ string, Hostname_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("langcode = ? and name = ? and hostname = ?", Langcode_, Name_, Hostname_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByLangcodeAndNameAndCreated Get CommentFieldDatas via LangcodeAndNameAndCreated
func GetCommentFieldDatasByLangcodeAndNameAndCreated(offset int, limit int, Langcode_ string, Name_ string, Created_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("langcode = ? and name = ? and created = ?", Langcode_, Name_, Created_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByLangcodeAndNameAndChanged Get CommentFieldDatas via LangcodeAndNameAndChanged
func GetCommentFieldDatasByLangcodeAndNameAndChanged(offset int, limit int, Langcode_ string, Name_ string, Changed_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("langcode = ? and name = ? and changed = ?", Langcode_, Name_, Changed_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByLangcodeAndNameAndStatus Get CommentFieldDatas via LangcodeAndNameAndStatus
func GetCommentFieldDatasByLangcodeAndNameAndStatus(offset int, limit int, Langcode_ string, Name_ string, Status_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("langcode = ? and name = ? and status = ?", Langcode_, Name_, Status_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByLangcodeAndNameAndThread Get CommentFieldDatas via LangcodeAndNameAndThread
func GetCommentFieldDatasByLangcodeAndNameAndThread(offset int, limit int, Langcode_ string, Name_ string, Thread_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("langcode = ? and name = ? and thread = ?", Langcode_, Name_, Thread_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByLangcodeAndNameAndEntityType Get CommentFieldDatas via LangcodeAndNameAndEntityType
func GetCommentFieldDatasByLangcodeAndNameAndEntityType(offset int, limit int, Langcode_ string, Name_ string, EntityType_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("langcode = ? and name = ? and entity_type = ?", Langcode_, Name_, EntityType_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByLangcodeAndNameAndFieldName Get CommentFieldDatas via LangcodeAndNameAndFieldName
func GetCommentFieldDatasByLangcodeAndNameAndFieldName(offset int, limit int, Langcode_ string, Name_ string, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("langcode = ? and name = ? and field_name = ?", Langcode_, Name_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByLangcodeAndNameAndDefaultLangcode Get CommentFieldDatas via LangcodeAndNameAndDefaultLangcode
func GetCommentFieldDatasByLangcodeAndNameAndDefaultLangcode(offset int, limit int, Langcode_ string, Name_ string, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("langcode = ? and name = ? and default_langcode = ?", Langcode_, Name_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByLangcodeAndMailAndHomepage Get CommentFieldDatas via LangcodeAndMailAndHomepage
func GetCommentFieldDatasByLangcodeAndMailAndHomepage(offset int, limit int, Langcode_ string, Mail_ string, Homepage_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("langcode = ? and mail = ? and homepage = ?", Langcode_, Mail_, Homepage_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByLangcodeAndMailAndHostname Get CommentFieldDatas via LangcodeAndMailAndHostname
func GetCommentFieldDatasByLangcodeAndMailAndHostname(offset int, limit int, Langcode_ string, Mail_ string, Hostname_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("langcode = ? and mail = ? and hostname = ?", Langcode_, Mail_, Hostname_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByLangcodeAndMailAndCreated Get CommentFieldDatas via LangcodeAndMailAndCreated
func GetCommentFieldDatasByLangcodeAndMailAndCreated(offset int, limit int, Langcode_ string, Mail_ string, Created_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("langcode = ? and mail = ? and created = ?", Langcode_, Mail_, Created_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByLangcodeAndMailAndChanged Get CommentFieldDatas via LangcodeAndMailAndChanged
func GetCommentFieldDatasByLangcodeAndMailAndChanged(offset int, limit int, Langcode_ string, Mail_ string, Changed_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("langcode = ? and mail = ? and changed = ?", Langcode_, Mail_, Changed_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByLangcodeAndMailAndStatus Get CommentFieldDatas via LangcodeAndMailAndStatus
func GetCommentFieldDatasByLangcodeAndMailAndStatus(offset int, limit int, Langcode_ string, Mail_ string, Status_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("langcode = ? and mail = ? and status = ?", Langcode_, Mail_, Status_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByLangcodeAndMailAndThread Get CommentFieldDatas via LangcodeAndMailAndThread
func GetCommentFieldDatasByLangcodeAndMailAndThread(offset int, limit int, Langcode_ string, Mail_ string, Thread_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("langcode = ? and mail = ? and thread = ?", Langcode_, Mail_, Thread_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByLangcodeAndMailAndEntityType Get CommentFieldDatas via LangcodeAndMailAndEntityType
func GetCommentFieldDatasByLangcodeAndMailAndEntityType(offset int, limit int, Langcode_ string, Mail_ string, EntityType_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("langcode = ? and mail = ? and entity_type = ?", Langcode_, Mail_, EntityType_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByLangcodeAndMailAndFieldName Get CommentFieldDatas via LangcodeAndMailAndFieldName
func GetCommentFieldDatasByLangcodeAndMailAndFieldName(offset int, limit int, Langcode_ string, Mail_ string, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("langcode = ? and mail = ? and field_name = ?", Langcode_, Mail_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByLangcodeAndMailAndDefaultLangcode Get CommentFieldDatas via LangcodeAndMailAndDefaultLangcode
func GetCommentFieldDatasByLangcodeAndMailAndDefaultLangcode(offset int, limit int, Langcode_ string, Mail_ string, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("langcode = ? and mail = ? and default_langcode = ?", Langcode_, Mail_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByLangcodeAndHomepageAndHostname Get CommentFieldDatas via LangcodeAndHomepageAndHostname
func GetCommentFieldDatasByLangcodeAndHomepageAndHostname(offset int, limit int, Langcode_ string, Homepage_ string, Hostname_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("langcode = ? and homepage = ? and hostname = ?", Langcode_, Homepage_, Hostname_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByLangcodeAndHomepageAndCreated Get CommentFieldDatas via LangcodeAndHomepageAndCreated
func GetCommentFieldDatasByLangcodeAndHomepageAndCreated(offset int, limit int, Langcode_ string, Homepage_ string, Created_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("langcode = ? and homepage = ? and created = ?", Langcode_, Homepage_, Created_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByLangcodeAndHomepageAndChanged Get CommentFieldDatas via LangcodeAndHomepageAndChanged
func GetCommentFieldDatasByLangcodeAndHomepageAndChanged(offset int, limit int, Langcode_ string, Homepage_ string, Changed_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("langcode = ? and homepage = ? and changed = ?", Langcode_, Homepage_, Changed_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByLangcodeAndHomepageAndStatus Get CommentFieldDatas via LangcodeAndHomepageAndStatus
func GetCommentFieldDatasByLangcodeAndHomepageAndStatus(offset int, limit int, Langcode_ string, Homepage_ string, Status_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("langcode = ? and homepage = ? and status = ?", Langcode_, Homepage_, Status_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByLangcodeAndHomepageAndThread Get CommentFieldDatas via LangcodeAndHomepageAndThread
func GetCommentFieldDatasByLangcodeAndHomepageAndThread(offset int, limit int, Langcode_ string, Homepage_ string, Thread_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("langcode = ? and homepage = ? and thread = ?", Langcode_, Homepage_, Thread_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByLangcodeAndHomepageAndEntityType Get CommentFieldDatas via LangcodeAndHomepageAndEntityType
func GetCommentFieldDatasByLangcodeAndHomepageAndEntityType(offset int, limit int, Langcode_ string, Homepage_ string, EntityType_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("langcode = ? and homepage = ? and entity_type = ?", Langcode_, Homepage_, EntityType_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByLangcodeAndHomepageAndFieldName Get CommentFieldDatas via LangcodeAndHomepageAndFieldName
func GetCommentFieldDatasByLangcodeAndHomepageAndFieldName(offset int, limit int, Langcode_ string, Homepage_ string, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("langcode = ? and homepage = ? and field_name = ?", Langcode_, Homepage_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByLangcodeAndHomepageAndDefaultLangcode Get CommentFieldDatas via LangcodeAndHomepageAndDefaultLangcode
func GetCommentFieldDatasByLangcodeAndHomepageAndDefaultLangcode(offset int, limit int, Langcode_ string, Homepage_ string, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("langcode = ? and homepage = ? and default_langcode = ?", Langcode_, Homepage_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByLangcodeAndHostnameAndCreated Get CommentFieldDatas via LangcodeAndHostnameAndCreated
func GetCommentFieldDatasByLangcodeAndHostnameAndCreated(offset int, limit int, Langcode_ string, Hostname_ string, Created_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("langcode = ? and hostname = ? and created = ?", Langcode_, Hostname_, Created_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByLangcodeAndHostnameAndChanged Get CommentFieldDatas via LangcodeAndHostnameAndChanged
func GetCommentFieldDatasByLangcodeAndHostnameAndChanged(offset int, limit int, Langcode_ string, Hostname_ string, Changed_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("langcode = ? and hostname = ? and changed = ?", Langcode_, Hostname_, Changed_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByLangcodeAndHostnameAndStatus Get CommentFieldDatas via LangcodeAndHostnameAndStatus
func GetCommentFieldDatasByLangcodeAndHostnameAndStatus(offset int, limit int, Langcode_ string, Hostname_ string, Status_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("langcode = ? and hostname = ? and status = ?", Langcode_, Hostname_, Status_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByLangcodeAndHostnameAndThread Get CommentFieldDatas via LangcodeAndHostnameAndThread
func GetCommentFieldDatasByLangcodeAndHostnameAndThread(offset int, limit int, Langcode_ string, Hostname_ string, Thread_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("langcode = ? and hostname = ? and thread = ?", Langcode_, Hostname_, Thread_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByLangcodeAndHostnameAndEntityType Get CommentFieldDatas via LangcodeAndHostnameAndEntityType
func GetCommentFieldDatasByLangcodeAndHostnameAndEntityType(offset int, limit int, Langcode_ string, Hostname_ string, EntityType_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("langcode = ? and hostname = ? and entity_type = ?", Langcode_, Hostname_, EntityType_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByLangcodeAndHostnameAndFieldName Get CommentFieldDatas via LangcodeAndHostnameAndFieldName
func GetCommentFieldDatasByLangcodeAndHostnameAndFieldName(offset int, limit int, Langcode_ string, Hostname_ string, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("langcode = ? and hostname = ? and field_name = ?", Langcode_, Hostname_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByLangcodeAndHostnameAndDefaultLangcode Get CommentFieldDatas via LangcodeAndHostnameAndDefaultLangcode
func GetCommentFieldDatasByLangcodeAndHostnameAndDefaultLangcode(offset int, limit int, Langcode_ string, Hostname_ string, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("langcode = ? and hostname = ? and default_langcode = ?", Langcode_, Hostname_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByLangcodeAndCreatedAndChanged Get CommentFieldDatas via LangcodeAndCreatedAndChanged
func GetCommentFieldDatasByLangcodeAndCreatedAndChanged(offset int, limit int, Langcode_ string, Created_ int, Changed_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("langcode = ? and created = ? and changed = ?", Langcode_, Created_, Changed_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByLangcodeAndCreatedAndStatus Get CommentFieldDatas via LangcodeAndCreatedAndStatus
func GetCommentFieldDatasByLangcodeAndCreatedAndStatus(offset int, limit int, Langcode_ string, Created_ int, Status_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("langcode = ? and created = ? and status = ?", Langcode_, Created_, Status_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByLangcodeAndCreatedAndThread Get CommentFieldDatas via LangcodeAndCreatedAndThread
func GetCommentFieldDatasByLangcodeAndCreatedAndThread(offset int, limit int, Langcode_ string, Created_ int, Thread_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("langcode = ? and created = ? and thread = ?", Langcode_, Created_, Thread_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByLangcodeAndCreatedAndEntityType Get CommentFieldDatas via LangcodeAndCreatedAndEntityType
func GetCommentFieldDatasByLangcodeAndCreatedAndEntityType(offset int, limit int, Langcode_ string, Created_ int, EntityType_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("langcode = ? and created = ? and entity_type = ?", Langcode_, Created_, EntityType_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByLangcodeAndCreatedAndFieldName Get CommentFieldDatas via LangcodeAndCreatedAndFieldName
func GetCommentFieldDatasByLangcodeAndCreatedAndFieldName(offset int, limit int, Langcode_ string, Created_ int, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("langcode = ? and created = ? and field_name = ?", Langcode_, Created_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByLangcodeAndCreatedAndDefaultLangcode Get CommentFieldDatas via LangcodeAndCreatedAndDefaultLangcode
func GetCommentFieldDatasByLangcodeAndCreatedAndDefaultLangcode(offset int, limit int, Langcode_ string, Created_ int, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("langcode = ? and created = ? and default_langcode = ?", Langcode_, Created_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByLangcodeAndChangedAndStatus Get CommentFieldDatas via LangcodeAndChangedAndStatus
func GetCommentFieldDatasByLangcodeAndChangedAndStatus(offset int, limit int, Langcode_ string, Changed_ int, Status_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("langcode = ? and changed = ? and status = ?", Langcode_, Changed_, Status_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByLangcodeAndChangedAndThread Get CommentFieldDatas via LangcodeAndChangedAndThread
func GetCommentFieldDatasByLangcodeAndChangedAndThread(offset int, limit int, Langcode_ string, Changed_ int, Thread_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("langcode = ? and changed = ? and thread = ?", Langcode_, Changed_, Thread_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByLangcodeAndChangedAndEntityType Get CommentFieldDatas via LangcodeAndChangedAndEntityType
func GetCommentFieldDatasByLangcodeAndChangedAndEntityType(offset int, limit int, Langcode_ string, Changed_ int, EntityType_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("langcode = ? and changed = ? and entity_type = ?", Langcode_, Changed_, EntityType_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByLangcodeAndChangedAndFieldName Get CommentFieldDatas via LangcodeAndChangedAndFieldName
func GetCommentFieldDatasByLangcodeAndChangedAndFieldName(offset int, limit int, Langcode_ string, Changed_ int, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("langcode = ? and changed = ? and field_name = ?", Langcode_, Changed_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByLangcodeAndChangedAndDefaultLangcode Get CommentFieldDatas via LangcodeAndChangedAndDefaultLangcode
func GetCommentFieldDatasByLangcodeAndChangedAndDefaultLangcode(offset int, limit int, Langcode_ string, Changed_ int, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("langcode = ? and changed = ? and default_langcode = ?", Langcode_, Changed_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByLangcodeAndStatusAndThread Get CommentFieldDatas via LangcodeAndStatusAndThread
func GetCommentFieldDatasByLangcodeAndStatusAndThread(offset int, limit int, Langcode_ string, Status_ int, Thread_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("langcode = ? and status = ? and thread = ?", Langcode_, Status_, Thread_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByLangcodeAndStatusAndEntityType Get CommentFieldDatas via LangcodeAndStatusAndEntityType
func GetCommentFieldDatasByLangcodeAndStatusAndEntityType(offset int, limit int, Langcode_ string, Status_ int, EntityType_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("langcode = ? and status = ? and entity_type = ?", Langcode_, Status_, EntityType_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByLangcodeAndStatusAndFieldName Get CommentFieldDatas via LangcodeAndStatusAndFieldName
func GetCommentFieldDatasByLangcodeAndStatusAndFieldName(offset int, limit int, Langcode_ string, Status_ int, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("langcode = ? and status = ? and field_name = ?", Langcode_, Status_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByLangcodeAndStatusAndDefaultLangcode Get CommentFieldDatas via LangcodeAndStatusAndDefaultLangcode
func GetCommentFieldDatasByLangcodeAndStatusAndDefaultLangcode(offset int, limit int, Langcode_ string, Status_ int, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("langcode = ? and status = ? and default_langcode = ?", Langcode_, Status_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByLangcodeAndThreadAndEntityType Get CommentFieldDatas via LangcodeAndThreadAndEntityType
func GetCommentFieldDatasByLangcodeAndThreadAndEntityType(offset int, limit int, Langcode_ string, Thread_ string, EntityType_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("langcode = ? and thread = ? and entity_type = ?", Langcode_, Thread_, EntityType_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByLangcodeAndThreadAndFieldName Get CommentFieldDatas via LangcodeAndThreadAndFieldName
func GetCommentFieldDatasByLangcodeAndThreadAndFieldName(offset int, limit int, Langcode_ string, Thread_ string, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("langcode = ? and thread = ? and field_name = ?", Langcode_, Thread_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByLangcodeAndThreadAndDefaultLangcode Get CommentFieldDatas via LangcodeAndThreadAndDefaultLangcode
func GetCommentFieldDatasByLangcodeAndThreadAndDefaultLangcode(offset int, limit int, Langcode_ string, Thread_ string, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("langcode = ? and thread = ? and default_langcode = ?", Langcode_, Thread_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByLangcodeAndEntityTypeAndFieldName Get CommentFieldDatas via LangcodeAndEntityTypeAndFieldName
func GetCommentFieldDatasByLangcodeAndEntityTypeAndFieldName(offset int, limit int, Langcode_ string, EntityType_ string, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("langcode = ? and entity_type = ? and field_name = ?", Langcode_, EntityType_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByLangcodeAndEntityTypeAndDefaultLangcode Get CommentFieldDatas via LangcodeAndEntityTypeAndDefaultLangcode
func GetCommentFieldDatasByLangcodeAndEntityTypeAndDefaultLangcode(offset int, limit int, Langcode_ string, EntityType_ string, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("langcode = ? and entity_type = ? and default_langcode = ?", Langcode_, EntityType_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByLangcodeAndFieldNameAndDefaultLangcode Get CommentFieldDatas via LangcodeAndFieldNameAndDefaultLangcode
func GetCommentFieldDatasByLangcodeAndFieldNameAndDefaultLangcode(offset int, limit int, Langcode_ string, FieldName_ string, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("langcode = ? and field_name = ? and default_langcode = ?", Langcode_, FieldName_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByPidAndEntityIdAndSubject Get CommentFieldDatas via PidAndEntityIdAndSubject
func GetCommentFieldDatasByPidAndEntityIdAndSubject(offset int, limit int, Pid_ int, EntityId_ int, Subject_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("pid = ? and entity_id = ? and subject = ?", Pid_, EntityId_, Subject_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByPidAndEntityIdAndUid Get CommentFieldDatas via PidAndEntityIdAndUid
func GetCommentFieldDatasByPidAndEntityIdAndUid(offset int, limit int, Pid_ int, EntityId_ int, Uid_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("pid = ? and entity_id = ? and uid = ?", Pid_, EntityId_, Uid_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByPidAndEntityIdAndName Get CommentFieldDatas via PidAndEntityIdAndName
func GetCommentFieldDatasByPidAndEntityIdAndName(offset int, limit int, Pid_ int, EntityId_ int, Name_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("pid = ? and entity_id = ? and name = ?", Pid_, EntityId_, Name_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByPidAndEntityIdAndMail Get CommentFieldDatas via PidAndEntityIdAndMail
func GetCommentFieldDatasByPidAndEntityIdAndMail(offset int, limit int, Pid_ int, EntityId_ int, Mail_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("pid = ? and entity_id = ? and mail = ?", Pid_, EntityId_, Mail_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByPidAndEntityIdAndHomepage Get CommentFieldDatas via PidAndEntityIdAndHomepage
func GetCommentFieldDatasByPidAndEntityIdAndHomepage(offset int, limit int, Pid_ int, EntityId_ int, Homepage_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("pid = ? and entity_id = ? and homepage = ?", Pid_, EntityId_, Homepage_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByPidAndEntityIdAndHostname Get CommentFieldDatas via PidAndEntityIdAndHostname
func GetCommentFieldDatasByPidAndEntityIdAndHostname(offset int, limit int, Pid_ int, EntityId_ int, Hostname_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("pid = ? and entity_id = ? and hostname = ?", Pid_, EntityId_, Hostname_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByPidAndEntityIdAndCreated Get CommentFieldDatas via PidAndEntityIdAndCreated
func GetCommentFieldDatasByPidAndEntityIdAndCreated(offset int, limit int, Pid_ int, EntityId_ int, Created_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("pid = ? and entity_id = ? and created = ?", Pid_, EntityId_, Created_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByPidAndEntityIdAndChanged Get CommentFieldDatas via PidAndEntityIdAndChanged
func GetCommentFieldDatasByPidAndEntityIdAndChanged(offset int, limit int, Pid_ int, EntityId_ int, Changed_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("pid = ? and entity_id = ? and changed = ?", Pid_, EntityId_, Changed_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByPidAndEntityIdAndStatus Get CommentFieldDatas via PidAndEntityIdAndStatus
func GetCommentFieldDatasByPidAndEntityIdAndStatus(offset int, limit int, Pid_ int, EntityId_ int, Status_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("pid = ? and entity_id = ? and status = ?", Pid_, EntityId_, Status_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByPidAndEntityIdAndThread Get CommentFieldDatas via PidAndEntityIdAndThread
func GetCommentFieldDatasByPidAndEntityIdAndThread(offset int, limit int, Pid_ int, EntityId_ int, Thread_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("pid = ? and entity_id = ? and thread = ?", Pid_, EntityId_, Thread_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByPidAndEntityIdAndEntityType Get CommentFieldDatas via PidAndEntityIdAndEntityType
func GetCommentFieldDatasByPidAndEntityIdAndEntityType(offset int, limit int, Pid_ int, EntityId_ int, EntityType_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("pid = ? and entity_id = ? and entity_type = ?", Pid_, EntityId_, EntityType_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByPidAndEntityIdAndFieldName Get CommentFieldDatas via PidAndEntityIdAndFieldName
func GetCommentFieldDatasByPidAndEntityIdAndFieldName(offset int, limit int, Pid_ int, EntityId_ int, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("pid = ? and entity_id = ? and field_name = ?", Pid_, EntityId_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByPidAndEntityIdAndDefaultLangcode Get CommentFieldDatas via PidAndEntityIdAndDefaultLangcode
func GetCommentFieldDatasByPidAndEntityIdAndDefaultLangcode(offset int, limit int, Pid_ int, EntityId_ int, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("pid = ? and entity_id = ? and default_langcode = ?", Pid_, EntityId_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByPidAndSubjectAndUid Get CommentFieldDatas via PidAndSubjectAndUid
func GetCommentFieldDatasByPidAndSubjectAndUid(offset int, limit int, Pid_ int, Subject_ string, Uid_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("pid = ? and subject = ? and uid = ?", Pid_, Subject_, Uid_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByPidAndSubjectAndName Get CommentFieldDatas via PidAndSubjectAndName
func GetCommentFieldDatasByPidAndSubjectAndName(offset int, limit int, Pid_ int, Subject_ string, Name_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("pid = ? and subject = ? and name = ?", Pid_, Subject_, Name_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByPidAndSubjectAndMail Get CommentFieldDatas via PidAndSubjectAndMail
func GetCommentFieldDatasByPidAndSubjectAndMail(offset int, limit int, Pid_ int, Subject_ string, Mail_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("pid = ? and subject = ? and mail = ?", Pid_, Subject_, Mail_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByPidAndSubjectAndHomepage Get CommentFieldDatas via PidAndSubjectAndHomepage
func GetCommentFieldDatasByPidAndSubjectAndHomepage(offset int, limit int, Pid_ int, Subject_ string, Homepage_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("pid = ? and subject = ? and homepage = ?", Pid_, Subject_, Homepage_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByPidAndSubjectAndHostname Get CommentFieldDatas via PidAndSubjectAndHostname
func GetCommentFieldDatasByPidAndSubjectAndHostname(offset int, limit int, Pid_ int, Subject_ string, Hostname_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("pid = ? and subject = ? and hostname = ?", Pid_, Subject_, Hostname_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByPidAndSubjectAndCreated Get CommentFieldDatas via PidAndSubjectAndCreated
func GetCommentFieldDatasByPidAndSubjectAndCreated(offset int, limit int, Pid_ int, Subject_ string, Created_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("pid = ? and subject = ? and created = ?", Pid_, Subject_, Created_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByPidAndSubjectAndChanged Get CommentFieldDatas via PidAndSubjectAndChanged
func GetCommentFieldDatasByPidAndSubjectAndChanged(offset int, limit int, Pid_ int, Subject_ string, Changed_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("pid = ? and subject = ? and changed = ?", Pid_, Subject_, Changed_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByPidAndSubjectAndStatus Get CommentFieldDatas via PidAndSubjectAndStatus
func GetCommentFieldDatasByPidAndSubjectAndStatus(offset int, limit int, Pid_ int, Subject_ string, Status_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("pid = ? and subject = ? and status = ?", Pid_, Subject_, Status_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByPidAndSubjectAndThread Get CommentFieldDatas via PidAndSubjectAndThread
func GetCommentFieldDatasByPidAndSubjectAndThread(offset int, limit int, Pid_ int, Subject_ string, Thread_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("pid = ? and subject = ? and thread = ?", Pid_, Subject_, Thread_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByPidAndSubjectAndEntityType Get CommentFieldDatas via PidAndSubjectAndEntityType
func GetCommentFieldDatasByPidAndSubjectAndEntityType(offset int, limit int, Pid_ int, Subject_ string, EntityType_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("pid = ? and subject = ? and entity_type = ?", Pid_, Subject_, EntityType_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByPidAndSubjectAndFieldName Get CommentFieldDatas via PidAndSubjectAndFieldName
func GetCommentFieldDatasByPidAndSubjectAndFieldName(offset int, limit int, Pid_ int, Subject_ string, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("pid = ? and subject = ? and field_name = ?", Pid_, Subject_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByPidAndSubjectAndDefaultLangcode Get CommentFieldDatas via PidAndSubjectAndDefaultLangcode
func GetCommentFieldDatasByPidAndSubjectAndDefaultLangcode(offset int, limit int, Pid_ int, Subject_ string, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("pid = ? and subject = ? and default_langcode = ?", Pid_, Subject_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByPidAndUidAndName Get CommentFieldDatas via PidAndUidAndName
func GetCommentFieldDatasByPidAndUidAndName(offset int, limit int, Pid_ int, Uid_ int, Name_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("pid = ? and uid = ? and name = ?", Pid_, Uid_, Name_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByPidAndUidAndMail Get CommentFieldDatas via PidAndUidAndMail
func GetCommentFieldDatasByPidAndUidAndMail(offset int, limit int, Pid_ int, Uid_ int, Mail_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("pid = ? and uid = ? and mail = ?", Pid_, Uid_, Mail_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByPidAndUidAndHomepage Get CommentFieldDatas via PidAndUidAndHomepage
func GetCommentFieldDatasByPidAndUidAndHomepage(offset int, limit int, Pid_ int, Uid_ int, Homepage_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("pid = ? and uid = ? and homepage = ?", Pid_, Uid_, Homepage_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByPidAndUidAndHostname Get CommentFieldDatas via PidAndUidAndHostname
func GetCommentFieldDatasByPidAndUidAndHostname(offset int, limit int, Pid_ int, Uid_ int, Hostname_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("pid = ? and uid = ? and hostname = ?", Pid_, Uid_, Hostname_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByPidAndUidAndCreated Get CommentFieldDatas via PidAndUidAndCreated
func GetCommentFieldDatasByPidAndUidAndCreated(offset int, limit int, Pid_ int, Uid_ int, Created_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("pid = ? and uid = ? and created = ?", Pid_, Uid_, Created_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByPidAndUidAndChanged Get CommentFieldDatas via PidAndUidAndChanged
func GetCommentFieldDatasByPidAndUidAndChanged(offset int, limit int, Pid_ int, Uid_ int, Changed_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("pid = ? and uid = ? and changed = ?", Pid_, Uid_, Changed_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByPidAndUidAndStatus Get CommentFieldDatas via PidAndUidAndStatus
func GetCommentFieldDatasByPidAndUidAndStatus(offset int, limit int, Pid_ int, Uid_ int, Status_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("pid = ? and uid = ? and status = ?", Pid_, Uid_, Status_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByPidAndUidAndThread Get CommentFieldDatas via PidAndUidAndThread
func GetCommentFieldDatasByPidAndUidAndThread(offset int, limit int, Pid_ int, Uid_ int, Thread_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("pid = ? and uid = ? and thread = ?", Pid_, Uid_, Thread_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByPidAndUidAndEntityType Get CommentFieldDatas via PidAndUidAndEntityType
func GetCommentFieldDatasByPidAndUidAndEntityType(offset int, limit int, Pid_ int, Uid_ int, EntityType_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("pid = ? and uid = ? and entity_type = ?", Pid_, Uid_, EntityType_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByPidAndUidAndFieldName Get CommentFieldDatas via PidAndUidAndFieldName
func GetCommentFieldDatasByPidAndUidAndFieldName(offset int, limit int, Pid_ int, Uid_ int, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("pid = ? and uid = ? and field_name = ?", Pid_, Uid_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByPidAndUidAndDefaultLangcode Get CommentFieldDatas via PidAndUidAndDefaultLangcode
func GetCommentFieldDatasByPidAndUidAndDefaultLangcode(offset int, limit int, Pid_ int, Uid_ int, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("pid = ? and uid = ? and default_langcode = ?", Pid_, Uid_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByPidAndNameAndMail Get CommentFieldDatas via PidAndNameAndMail
func GetCommentFieldDatasByPidAndNameAndMail(offset int, limit int, Pid_ int, Name_ string, Mail_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("pid = ? and name = ? and mail = ?", Pid_, Name_, Mail_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByPidAndNameAndHomepage Get CommentFieldDatas via PidAndNameAndHomepage
func GetCommentFieldDatasByPidAndNameAndHomepage(offset int, limit int, Pid_ int, Name_ string, Homepage_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("pid = ? and name = ? and homepage = ?", Pid_, Name_, Homepage_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByPidAndNameAndHostname Get CommentFieldDatas via PidAndNameAndHostname
func GetCommentFieldDatasByPidAndNameAndHostname(offset int, limit int, Pid_ int, Name_ string, Hostname_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("pid = ? and name = ? and hostname = ?", Pid_, Name_, Hostname_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByPidAndNameAndCreated Get CommentFieldDatas via PidAndNameAndCreated
func GetCommentFieldDatasByPidAndNameAndCreated(offset int, limit int, Pid_ int, Name_ string, Created_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("pid = ? and name = ? and created = ?", Pid_, Name_, Created_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByPidAndNameAndChanged Get CommentFieldDatas via PidAndNameAndChanged
func GetCommentFieldDatasByPidAndNameAndChanged(offset int, limit int, Pid_ int, Name_ string, Changed_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("pid = ? and name = ? and changed = ?", Pid_, Name_, Changed_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByPidAndNameAndStatus Get CommentFieldDatas via PidAndNameAndStatus
func GetCommentFieldDatasByPidAndNameAndStatus(offset int, limit int, Pid_ int, Name_ string, Status_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("pid = ? and name = ? and status = ?", Pid_, Name_, Status_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByPidAndNameAndThread Get CommentFieldDatas via PidAndNameAndThread
func GetCommentFieldDatasByPidAndNameAndThread(offset int, limit int, Pid_ int, Name_ string, Thread_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("pid = ? and name = ? and thread = ?", Pid_, Name_, Thread_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByPidAndNameAndEntityType Get CommentFieldDatas via PidAndNameAndEntityType
func GetCommentFieldDatasByPidAndNameAndEntityType(offset int, limit int, Pid_ int, Name_ string, EntityType_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("pid = ? and name = ? and entity_type = ?", Pid_, Name_, EntityType_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByPidAndNameAndFieldName Get CommentFieldDatas via PidAndNameAndFieldName
func GetCommentFieldDatasByPidAndNameAndFieldName(offset int, limit int, Pid_ int, Name_ string, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("pid = ? and name = ? and field_name = ?", Pid_, Name_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByPidAndNameAndDefaultLangcode Get CommentFieldDatas via PidAndNameAndDefaultLangcode
func GetCommentFieldDatasByPidAndNameAndDefaultLangcode(offset int, limit int, Pid_ int, Name_ string, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("pid = ? and name = ? and default_langcode = ?", Pid_, Name_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByPidAndMailAndHomepage Get CommentFieldDatas via PidAndMailAndHomepage
func GetCommentFieldDatasByPidAndMailAndHomepage(offset int, limit int, Pid_ int, Mail_ string, Homepage_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("pid = ? and mail = ? and homepage = ?", Pid_, Mail_, Homepage_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByPidAndMailAndHostname Get CommentFieldDatas via PidAndMailAndHostname
func GetCommentFieldDatasByPidAndMailAndHostname(offset int, limit int, Pid_ int, Mail_ string, Hostname_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("pid = ? and mail = ? and hostname = ?", Pid_, Mail_, Hostname_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByPidAndMailAndCreated Get CommentFieldDatas via PidAndMailAndCreated
func GetCommentFieldDatasByPidAndMailAndCreated(offset int, limit int, Pid_ int, Mail_ string, Created_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("pid = ? and mail = ? and created = ?", Pid_, Mail_, Created_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByPidAndMailAndChanged Get CommentFieldDatas via PidAndMailAndChanged
func GetCommentFieldDatasByPidAndMailAndChanged(offset int, limit int, Pid_ int, Mail_ string, Changed_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("pid = ? and mail = ? and changed = ?", Pid_, Mail_, Changed_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByPidAndMailAndStatus Get CommentFieldDatas via PidAndMailAndStatus
func GetCommentFieldDatasByPidAndMailAndStatus(offset int, limit int, Pid_ int, Mail_ string, Status_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("pid = ? and mail = ? and status = ?", Pid_, Mail_, Status_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByPidAndMailAndThread Get CommentFieldDatas via PidAndMailAndThread
func GetCommentFieldDatasByPidAndMailAndThread(offset int, limit int, Pid_ int, Mail_ string, Thread_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("pid = ? and mail = ? and thread = ?", Pid_, Mail_, Thread_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByPidAndMailAndEntityType Get CommentFieldDatas via PidAndMailAndEntityType
func GetCommentFieldDatasByPidAndMailAndEntityType(offset int, limit int, Pid_ int, Mail_ string, EntityType_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("pid = ? and mail = ? and entity_type = ?", Pid_, Mail_, EntityType_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByPidAndMailAndFieldName Get CommentFieldDatas via PidAndMailAndFieldName
func GetCommentFieldDatasByPidAndMailAndFieldName(offset int, limit int, Pid_ int, Mail_ string, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("pid = ? and mail = ? and field_name = ?", Pid_, Mail_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByPidAndMailAndDefaultLangcode Get CommentFieldDatas via PidAndMailAndDefaultLangcode
func GetCommentFieldDatasByPidAndMailAndDefaultLangcode(offset int, limit int, Pid_ int, Mail_ string, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("pid = ? and mail = ? and default_langcode = ?", Pid_, Mail_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByPidAndHomepageAndHostname Get CommentFieldDatas via PidAndHomepageAndHostname
func GetCommentFieldDatasByPidAndHomepageAndHostname(offset int, limit int, Pid_ int, Homepage_ string, Hostname_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("pid = ? and homepage = ? and hostname = ?", Pid_, Homepage_, Hostname_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByPidAndHomepageAndCreated Get CommentFieldDatas via PidAndHomepageAndCreated
func GetCommentFieldDatasByPidAndHomepageAndCreated(offset int, limit int, Pid_ int, Homepage_ string, Created_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("pid = ? and homepage = ? and created = ?", Pid_, Homepage_, Created_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByPidAndHomepageAndChanged Get CommentFieldDatas via PidAndHomepageAndChanged
func GetCommentFieldDatasByPidAndHomepageAndChanged(offset int, limit int, Pid_ int, Homepage_ string, Changed_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("pid = ? and homepage = ? and changed = ?", Pid_, Homepage_, Changed_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByPidAndHomepageAndStatus Get CommentFieldDatas via PidAndHomepageAndStatus
func GetCommentFieldDatasByPidAndHomepageAndStatus(offset int, limit int, Pid_ int, Homepage_ string, Status_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("pid = ? and homepage = ? and status = ?", Pid_, Homepage_, Status_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByPidAndHomepageAndThread Get CommentFieldDatas via PidAndHomepageAndThread
func GetCommentFieldDatasByPidAndHomepageAndThread(offset int, limit int, Pid_ int, Homepage_ string, Thread_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("pid = ? and homepage = ? and thread = ?", Pid_, Homepage_, Thread_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByPidAndHomepageAndEntityType Get CommentFieldDatas via PidAndHomepageAndEntityType
func GetCommentFieldDatasByPidAndHomepageAndEntityType(offset int, limit int, Pid_ int, Homepage_ string, EntityType_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("pid = ? and homepage = ? and entity_type = ?", Pid_, Homepage_, EntityType_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByPidAndHomepageAndFieldName Get CommentFieldDatas via PidAndHomepageAndFieldName
func GetCommentFieldDatasByPidAndHomepageAndFieldName(offset int, limit int, Pid_ int, Homepage_ string, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("pid = ? and homepage = ? and field_name = ?", Pid_, Homepage_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByPidAndHomepageAndDefaultLangcode Get CommentFieldDatas via PidAndHomepageAndDefaultLangcode
func GetCommentFieldDatasByPidAndHomepageAndDefaultLangcode(offset int, limit int, Pid_ int, Homepage_ string, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("pid = ? and homepage = ? and default_langcode = ?", Pid_, Homepage_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByPidAndHostnameAndCreated Get CommentFieldDatas via PidAndHostnameAndCreated
func GetCommentFieldDatasByPidAndHostnameAndCreated(offset int, limit int, Pid_ int, Hostname_ string, Created_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("pid = ? and hostname = ? and created = ?", Pid_, Hostname_, Created_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByPidAndHostnameAndChanged Get CommentFieldDatas via PidAndHostnameAndChanged
func GetCommentFieldDatasByPidAndHostnameAndChanged(offset int, limit int, Pid_ int, Hostname_ string, Changed_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("pid = ? and hostname = ? and changed = ?", Pid_, Hostname_, Changed_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByPidAndHostnameAndStatus Get CommentFieldDatas via PidAndHostnameAndStatus
func GetCommentFieldDatasByPidAndHostnameAndStatus(offset int, limit int, Pid_ int, Hostname_ string, Status_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("pid = ? and hostname = ? and status = ?", Pid_, Hostname_, Status_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByPidAndHostnameAndThread Get CommentFieldDatas via PidAndHostnameAndThread
func GetCommentFieldDatasByPidAndHostnameAndThread(offset int, limit int, Pid_ int, Hostname_ string, Thread_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("pid = ? and hostname = ? and thread = ?", Pid_, Hostname_, Thread_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByPidAndHostnameAndEntityType Get CommentFieldDatas via PidAndHostnameAndEntityType
func GetCommentFieldDatasByPidAndHostnameAndEntityType(offset int, limit int, Pid_ int, Hostname_ string, EntityType_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("pid = ? and hostname = ? and entity_type = ?", Pid_, Hostname_, EntityType_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByPidAndHostnameAndFieldName Get CommentFieldDatas via PidAndHostnameAndFieldName
func GetCommentFieldDatasByPidAndHostnameAndFieldName(offset int, limit int, Pid_ int, Hostname_ string, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("pid = ? and hostname = ? and field_name = ?", Pid_, Hostname_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByPidAndHostnameAndDefaultLangcode Get CommentFieldDatas via PidAndHostnameAndDefaultLangcode
func GetCommentFieldDatasByPidAndHostnameAndDefaultLangcode(offset int, limit int, Pid_ int, Hostname_ string, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("pid = ? and hostname = ? and default_langcode = ?", Pid_, Hostname_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByPidAndCreatedAndChanged Get CommentFieldDatas via PidAndCreatedAndChanged
func GetCommentFieldDatasByPidAndCreatedAndChanged(offset int, limit int, Pid_ int, Created_ int, Changed_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("pid = ? and created = ? and changed = ?", Pid_, Created_, Changed_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByPidAndCreatedAndStatus Get CommentFieldDatas via PidAndCreatedAndStatus
func GetCommentFieldDatasByPidAndCreatedAndStatus(offset int, limit int, Pid_ int, Created_ int, Status_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("pid = ? and created = ? and status = ?", Pid_, Created_, Status_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByPidAndCreatedAndThread Get CommentFieldDatas via PidAndCreatedAndThread
func GetCommentFieldDatasByPidAndCreatedAndThread(offset int, limit int, Pid_ int, Created_ int, Thread_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("pid = ? and created = ? and thread = ?", Pid_, Created_, Thread_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByPidAndCreatedAndEntityType Get CommentFieldDatas via PidAndCreatedAndEntityType
func GetCommentFieldDatasByPidAndCreatedAndEntityType(offset int, limit int, Pid_ int, Created_ int, EntityType_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("pid = ? and created = ? and entity_type = ?", Pid_, Created_, EntityType_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByPidAndCreatedAndFieldName Get CommentFieldDatas via PidAndCreatedAndFieldName
func GetCommentFieldDatasByPidAndCreatedAndFieldName(offset int, limit int, Pid_ int, Created_ int, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("pid = ? and created = ? and field_name = ?", Pid_, Created_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByPidAndCreatedAndDefaultLangcode Get CommentFieldDatas via PidAndCreatedAndDefaultLangcode
func GetCommentFieldDatasByPidAndCreatedAndDefaultLangcode(offset int, limit int, Pid_ int, Created_ int, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("pid = ? and created = ? and default_langcode = ?", Pid_, Created_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByPidAndChangedAndStatus Get CommentFieldDatas via PidAndChangedAndStatus
func GetCommentFieldDatasByPidAndChangedAndStatus(offset int, limit int, Pid_ int, Changed_ int, Status_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("pid = ? and changed = ? and status = ?", Pid_, Changed_, Status_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByPidAndChangedAndThread Get CommentFieldDatas via PidAndChangedAndThread
func GetCommentFieldDatasByPidAndChangedAndThread(offset int, limit int, Pid_ int, Changed_ int, Thread_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("pid = ? and changed = ? and thread = ?", Pid_, Changed_, Thread_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByPidAndChangedAndEntityType Get CommentFieldDatas via PidAndChangedAndEntityType
func GetCommentFieldDatasByPidAndChangedAndEntityType(offset int, limit int, Pid_ int, Changed_ int, EntityType_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("pid = ? and changed = ? and entity_type = ?", Pid_, Changed_, EntityType_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByPidAndChangedAndFieldName Get CommentFieldDatas via PidAndChangedAndFieldName
func GetCommentFieldDatasByPidAndChangedAndFieldName(offset int, limit int, Pid_ int, Changed_ int, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("pid = ? and changed = ? and field_name = ?", Pid_, Changed_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByPidAndChangedAndDefaultLangcode Get CommentFieldDatas via PidAndChangedAndDefaultLangcode
func GetCommentFieldDatasByPidAndChangedAndDefaultLangcode(offset int, limit int, Pid_ int, Changed_ int, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("pid = ? and changed = ? and default_langcode = ?", Pid_, Changed_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByPidAndStatusAndThread Get CommentFieldDatas via PidAndStatusAndThread
func GetCommentFieldDatasByPidAndStatusAndThread(offset int, limit int, Pid_ int, Status_ int, Thread_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("pid = ? and status = ? and thread = ?", Pid_, Status_, Thread_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByPidAndStatusAndEntityType Get CommentFieldDatas via PidAndStatusAndEntityType
func GetCommentFieldDatasByPidAndStatusAndEntityType(offset int, limit int, Pid_ int, Status_ int, EntityType_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("pid = ? and status = ? and entity_type = ?", Pid_, Status_, EntityType_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByPidAndStatusAndFieldName Get CommentFieldDatas via PidAndStatusAndFieldName
func GetCommentFieldDatasByPidAndStatusAndFieldName(offset int, limit int, Pid_ int, Status_ int, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("pid = ? and status = ? and field_name = ?", Pid_, Status_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByPidAndStatusAndDefaultLangcode Get CommentFieldDatas via PidAndStatusAndDefaultLangcode
func GetCommentFieldDatasByPidAndStatusAndDefaultLangcode(offset int, limit int, Pid_ int, Status_ int, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("pid = ? and status = ? and default_langcode = ?", Pid_, Status_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByPidAndThreadAndEntityType Get CommentFieldDatas via PidAndThreadAndEntityType
func GetCommentFieldDatasByPidAndThreadAndEntityType(offset int, limit int, Pid_ int, Thread_ string, EntityType_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("pid = ? and thread = ? and entity_type = ?", Pid_, Thread_, EntityType_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByPidAndThreadAndFieldName Get CommentFieldDatas via PidAndThreadAndFieldName
func GetCommentFieldDatasByPidAndThreadAndFieldName(offset int, limit int, Pid_ int, Thread_ string, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("pid = ? and thread = ? and field_name = ?", Pid_, Thread_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByPidAndThreadAndDefaultLangcode Get CommentFieldDatas via PidAndThreadAndDefaultLangcode
func GetCommentFieldDatasByPidAndThreadAndDefaultLangcode(offset int, limit int, Pid_ int, Thread_ string, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("pid = ? and thread = ? and default_langcode = ?", Pid_, Thread_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByPidAndEntityTypeAndFieldName Get CommentFieldDatas via PidAndEntityTypeAndFieldName
func GetCommentFieldDatasByPidAndEntityTypeAndFieldName(offset int, limit int, Pid_ int, EntityType_ string, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("pid = ? and entity_type = ? and field_name = ?", Pid_, EntityType_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByPidAndEntityTypeAndDefaultLangcode Get CommentFieldDatas via PidAndEntityTypeAndDefaultLangcode
func GetCommentFieldDatasByPidAndEntityTypeAndDefaultLangcode(offset int, limit int, Pid_ int, EntityType_ string, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("pid = ? and entity_type = ? and default_langcode = ?", Pid_, EntityType_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByPidAndFieldNameAndDefaultLangcode Get CommentFieldDatas via PidAndFieldNameAndDefaultLangcode
func GetCommentFieldDatasByPidAndFieldNameAndDefaultLangcode(offset int, limit int, Pid_ int, FieldName_ string, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("pid = ? and field_name = ? and default_langcode = ?", Pid_, FieldName_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByEntityIdAndSubjectAndUid Get CommentFieldDatas via EntityIdAndSubjectAndUid
func GetCommentFieldDatasByEntityIdAndSubjectAndUid(offset int, limit int, EntityId_ int, Subject_ string, Uid_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("entity_id = ? and subject = ? and uid = ?", EntityId_, Subject_, Uid_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByEntityIdAndSubjectAndName Get CommentFieldDatas via EntityIdAndSubjectAndName
func GetCommentFieldDatasByEntityIdAndSubjectAndName(offset int, limit int, EntityId_ int, Subject_ string, Name_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("entity_id = ? and subject = ? and name = ?", EntityId_, Subject_, Name_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByEntityIdAndSubjectAndMail Get CommentFieldDatas via EntityIdAndSubjectAndMail
func GetCommentFieldDatasByEntityIdAndSubjectAndMail(offset int, limit int, EntityId_ int, Subject_ string, Mail_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("entity_id = ? and subject = ? and mail = ?", EntityId_, Subject_, Mail_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByEntityIdAndSubjectAndHomepage Get CommentFieldDatas via EntityIdAndSubjectAndHomepage
func GetCommentFieldDatasByEntityIdAndSubjectAndHomepage(offset int, limit int, EntityId_ int, Subject_ string, Homepage_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("entity_id = ? and subject = ? and homepage = ?", EntityId_, Subject_, Homepage_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByEntityIdAndSubjectAndHostname Get CommentFieldDatas via EntityIdAndSubjectAndHostname
func GetCommentFieldDatasByEntityIdAndSubjectAndHostname(offset int, limit int, EntityId_ int, Subject_ string, Hostname_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("entity_id = ? and subject = ? and hostname = ?", EntityId_, Subject_, Hostname_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByEntityIdAndSubjectAndCreated Get CommentFieldDatas via EntityIdAndSubjectAndCreated
func GetCommentFieldDatasByEntityIdAndSubjectAndCreated(offset int, limit int, EntityId_ int, Subject_ string, Created_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("entity_id = ? and subject = ? and created = ?", EntityId_, Subject_, Created_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByEntityIdAndSubjectAndChanged Get CommentFieldDatas via EntityIdAndSubjectAndChanged
func GetCommentFieldDatasByEntityIdAndSubjectAndChanged(offset int, limit int, EntityId_ int, Subject_ string, Changed_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("entity_id = ? and subject = ? and changed = ?", EntityId_, Subject_, Changed_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByEntityIdAndSubjectAndStatus Get CommentFieldDatas via EntityIdAndSubjectAndStatus
func GetCommentFieldDatasByEntityIdAndSubjectAndStatus(offset int, limit int, EntityId_ int, Subject_ string, Status_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("entity_id = ? and subject = ? and status = ?", EntityId_, Subject_, Status_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByEntityIdAndSubjectAndThread Get CommentFieldDatas via EntityIdAndSubjectAndThread
func GetCommentFieldDatasByEntityIdAndSubjectAndThread(offset int, limit int, EntityId_ int, Subject_ string, Thread_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("entity_id = ? and subject = ? and thread = ?", EntityId_, Subject_, Thread_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByEntityIdAndSubjectAndEntityType Get CommentFieldDatas via EntityIdAndSubjectAndEntityType
func GetCommentFieldDatasByEntityIdAndSubjectAndEntityType(offset int, limit int, EntityId_ int, Subject_ string, EntityType_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("entity_id = ? and subject = ? and entity_type = ?", EntityId_, Subject_, EntityType_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByEntityIdAndSubjectAndFieldName Get CommentFieldDatas via EntityIdAndSubjectAndFieldName
func GetCommentFieldDatasByEntityIdAndSubjectAndFieldName(offset int, limit int, EntityId_ int, Subject_ string, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("entity_id = ? and subject = ? and field_name = ?", EntityId_, Subject_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByEntityIdAndSubjectAndDefaultLangcode Get CommentFieldDatas via EntityIdAndSubjectAndDefaultLangcode
func GetCommentFieldDatasByEntityIdAndSubjectAndDefaultLangcode(offset int, limit int, EntityId_ int, Subject_ string, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("entity_id = ? and subject = ? and default_langcode = ?", EntityId_, Subject_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByEntityIdAndUidAndName Get CommentFieldDatas via EntityIdAndUidAndName
func GetCommentFieldDatasByEntityIdAndUidAndName(offset int, limit int, EntityId_ int, Uid_ int, Name_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("entity_id = ? and uid = ? and name = ?", EntityId_, Uid_, Name_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByEntityIdAndUidAndMail Get CommentFieldDatas via EntityIdAndUidAndMail
func GetCommentFieldDatasByEntityIdAndUidAndMail(offset int, limit int, EntityId_ int, Uid_ int, Mail_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("entity_id = ? and uid = ? and mail = ?", EntityId_, Uid_, Mail_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByEntityIdAndUidAndHomepage Get CommentFieldDatas via EntityIdAndUidAndHomepage
func GetCommentFieldDatasByEntityIdAndUidAndHomepage(offset int, limit int, EntityId_ int, Uid_ int, Homepage_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("entity_id = ? and uid = ? and homepage = ?", EntityId_, Uid_, Homepage_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByEntityIdAndUidAndHostname Get CommentFieldDatas via EntityIdAndUidAndHostname
func GetCommentFieldDatasByEntityIdAndUidAndHostname(offset int, limit int, EntityId_ int, Uid_ int, Hostname_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("entity_id = ? and uid = ? and hostname = ?", EntityId_, Uid_, Hostname_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByEntityIdAndUidAndCreated Get CommentFieldDatas via EntityIdAndUidAndCreated
func GetCommentFieldDatasByEntityIdAndUidAndCreated(offset int, limit int, EntityId_ int, Uid_ int, Created_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("entity_id = ? and uid = ? and created = ?", EntityId_, Uid_, Created_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByEntityIdAndUidAndChanged Get CommentFieldDatas via EntityIdAndUidAndChanged
func GetCommentFieldDatasByEntityIdAndUidAndChanged(offset int, limit int, EntityId_ int, Uid_ int, Changed_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("entity_id = ? and uid = ? and changed = ?", EntityId_, Uid_, Changed_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByEntityIdAndUidAndStatus Get CommentFieldDatas via EntityIdAndUidAndStatus
func GetCommentFieldDatasByEntityIdAndUidAndStatus(offset int, limit int, EntityId_ int, Uid_ int, Status_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("entity_id = ? and uid = ? and status = ?", EntityId_, Uid_, Status_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByEntityIdAndUidAndThread Get CommentFieldDatas via EntityIdAndUidAndThread
func GetCommentFieldDatasByEntityIdAndUidAndThread(offset int, limit int, EntityId_ int, Uid_ int, Thread_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("entity_id = ? and uid = ? and thread = ?", EntityId_, Uid_, Thread_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByEntityIdAndUidAndEntityType Get CommentFieldDatas via EntityIdAndUidAndEntityType
func GetCommentFieldDatasByEntityIdAndUidAndEntityType(offset int, limit int, EntityId_ int, Uid_ int, EntityType_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("entity_id = ? and uid = ? and entity_type = ?", EntityId_, Uid_, EntityType_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByEntityIdAndUidAndFieldName Get CommentFieldDatas via EntityIdAndUidAndFieldName
func GetCommentFieldDatasByEntityIdAndUidAndFieldName(offset int, limit int, EntityId_ int, Uid_ int, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("entity_id = ? and uid = ? and field_name = ?", EntityId_, Uid_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByEntityIdAndUidAndDefaultLangcode Get CommentFieldDatas via EntityIdAndUidAndDefaultLangcode
func GetCommentFieldDatasByEntityIdAndUidAndDefaultLangcode(offset int, limit int, EntityId_ int, Uid_ int, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("entity_id = ? and uid = ? and default_langcode = ?", EntityId_, Uid_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByEntityIdAndNameAndMail Get CommentFieldDatas via EntityIdAndNameAndMail
func GetCommentFieldDatasByEntityIdAndNameAndMail(offset int, limit int, EntityId_ int, Name_ string, Mail_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("entity_id = ? and name = ? and mail = ?", EntityId_, Name_, Mail_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByEntityIdAndNameAndHomepage Get CommentFieldDatas via EntityIdAndNameAndHomepage
func GetCommentFieldDatasByEntityIdAndNameAndHomepage(offset int, limit int, EntityId_ int, Name_ string, Homepage_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("entity_id = ? and name = ? and homepage = ?", EntityId_, Name_, Homepage_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByEntityIdAndNameAndHostname Get CommentFieldDatas via EntityIdAndNameAndHostname
func GetCommentFieldDatasByEntityIdAndNameAndHostname(offset int, limit int, EntityId_ int, Name_ string, Hostname_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("entity_id = ? and name = ? and hostname = ?", EntityId_, Name_, Hostname_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByEntityIdAndNameAndCreated Get CommentFieldDatas via EntityIdAndNameAndCreated
func GetCommentFieldDatasByEntityIdAndNameAndCreated(offset int, limit int, EntityId_ int, Name_ string, Created_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("entity_id = ? and name = ? and created = ?", EntityId_, Name_, Created_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByEntityIdAndNameAndChanged Get CommentFieldDatas via EntityIdAndNameAndChanged
func GetCommentFieldDatasByEntityIdAndNameAndChanged(offset int, limit int, EntityId_ int, Name_ string, Changed_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("entity_id = ? and name = ? and changed = ?", EntityId_, Name_, Changed_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByEntityIdAndNameAndStatus Get CommentFieldDatas via EntityIdAndNameAndStatus
func GetCommentFieldDatasByEntityIdAndNameAndStatus(offset int, limit int, EntityId_ int, Name_ string, Status_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("entity_id = ? and name = ? and status = ?", EntityId_, Name_, Status_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByEntityIdAndNameAndThread Get CommentFieldDatas via EntityIdAndNameAndThread
func GetCommentFieldDatasByEntityIdAndNameAndThread(offset int, limit int, EntityId_ int, Name_ string, Thread_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("entity_id = ? and name = ? and thread = ?", EntityId_, Name_, Thread_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByEntityIdAndNameAndEntityType Get CommentFieldDatas via EntityIdAndNameAndEntityType
func GetCommentFieldDatasByEntityIdAndNameAndEntityType(offset int, limit int, EntityId_ int, Name_ string, EntityType_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("entity_id = ? and name = ? and entity_type = ?", EntityId_, Name_, EntityType_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByEntityIdAndNameAndFieldName Get CommentFieldDatas via EntityIdAndNameAndFieldName
func GetCommentFieldDatasByEntityIdAndNameAndFieldName(offset int, limit int, EntityId_ int, Name_ string, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("entity_id = ? and name = ? and field_name = ?", EntityId_, Name_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByEntityIdAndNameAndDefaultLangcode Get CommentFieldDatas via EntityIdAndNameAndDefaultLangcode
func GetCommentFieldDatasByEntityIdAndNameAndDefaultLangcode(offset int, limit int, EntityId_ int, Name_ string, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("entity_id = ? and name = ? and default_langcode = ?", EntityId_, Name_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByEntityIdAndMailAndHomepage Get CommentFieldDatas via EntityIdAndMailAndHomepage
func GetCommentFieldDatasByEntityIdAndMailAndHomepage(offset int, limit int, EntityId_ int, Mail_ string, Homepage_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("entity_id = ? and mail = ? and homepage = ?", EntityId_, Mail_, Homepage_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByEntityIdAndMailAndHostname Get CommentFieldDatas via EntityIdAndMailAndHostname
func GetCommentFieldDatasByEntityIdAndMailAndHostname(offset int, limit int, EntityId_ int, Mail_ string, Hostname_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("entity_id = ? and mail = ? and hostname = ?", EntityId_, Mail_, Hostname_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByEntityIdAndMailAndCreated Get CommentFieldDatas via EntityIdAndMailAndCreated
func GetCommentFieldDatasByEntityIdAndMailAndCreated(offset int, limit int, EntityId_ int, Mail_ string, Created_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("entity_id = ? and mail = ? and created = ?", EntityId_, Mail_, Created_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByEntityIdAndMailAndChanged Get CommentFieldDatas via EntityIdAndMailAndChanged
func GetCommentFieldDatasByEntityIdAndMailAndChanged(offset int, limit int, EntityId_ int, Mail_ string, Changed_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("entity_id = ? and mail = ? and changed = ?", EntityId_, Mail_, Changed_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByEntityIdAndMailAndStatus Get CommentFieldDatas via EntityIdAndMailAndStatus
func GetCommentFieldDatasByEntityIdAndMailAndStatus(offset int, limit int, EntityId_ int, Mail_ string, Status_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("entity_id = ? and mail = ? and status = ?", EntityId_, Mail_, Status_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByEntityIdAndMailAndThread Get CommentFieldDatas via EntityIdAndMailAndThread
func GetCommentFieldDatasByEntityIdAndMailAndThread(offset int, limit int, EntityId_ int, Mail_ string, Thread_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("entity_id = ? and mail = ? and thread = ?", EntityId_, Mail_, Thread_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByEntityIdAndMailAndEntityType Get CommentFieldDatas via EntityIdAndMailAndEntityType
func GetCommentFieldDatasByEntityIdAndMailAndEntityType(offset int, limit int, EntityId_ int, Mail_ string, EntityType_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("entity_id = ? and mail = ? and entity_type = ?", EntityId_, Mail_, EntityType_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByEntityIdAndMailAndFieldName Get CommentFieldDatas via EntityIdAndMailAndFieldName
func GetCommentFieldDatasByEntityIdAndMailAndFieldName(offset int, limit int, EntityId_ int, Mail_ string, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("entity_id = ? and mail = ? and field_name = ?", EntityId_, Mail_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByEntityIdAndMailAndDefaultLangcode Get CommentFieldDatas via EntityIdAndMailAndDefaultLangcode
func GetCommentFieldDatasByEntityIdAndMailAndDefaultLangcode(offset int, limit int, EntityId_ int, Mail_ string, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("entity_id = ? and mail = ? and default_langcode = ?", EntityId_, Mail_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByEntityIdAndHomepageAndHostname Get CommentFieldDatas via EntityIdAndHomepageAndHostname
func GetCommentFieldDatasByEntityIdAndHomepageAndHostname(offset int, limit int, EntityId_ int, Homepage_ string, Hostname_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("entity_id = ? and homepage = ? and hostname = ?", EntityId_, Homepage_, Hostname_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByEntityIdAndHomepageAndCreated Get CommentFieldDatas via EntityIdAndHomepageAndCreated
func GetCommentFieldDatasByEntityIdAndHomepageAndCreated(offset int, limit int, EntityId_ int, Homepage_ string, Created_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("entity_id = ? and homepage = ? and created = ?", EntityId_, Homepage_, Created_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByEntityIdAndHomepageAndChanged Get CommentFieldDatas via EntityIdAndHomepageAndChanged
func GetCommentFieldDatasByEntityIdAndHomepageAndChanged(offset int, limit int, EntityId_ int, Homepage_ string, Changed_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("entity_id = ? and homepage = ? and changed = ?", EntityId_, Homepage_, Changed_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByEntityIdAndHomepageAndStatus Get CommentFieldDatas via EntityIdAndHomepageAndStatus
func GetCommentFieldDatasByEntityIdAndHomepageAndStatus(offset int, limit int, EntityId_ int, Homepage_ string, Status_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("entity_id = ? and homepage = ? and status = ?", EntityId_, Homepage_, Status_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByEntityIdAndHomepageAndThread Get CommentFieldDatas via EntityIdAndHomepageAndThread
func GetCommentFieldDatasByEntityIdAndHomepageAndThread(offset int, limit int, EntityId_ int, Homepage_ string, Thread_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("entity_id = ? and homepage = ? and thread = ?", EntityId_, Homepage_, Thread_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByEntityIdAndHomepageAndEntityType Get CommentFieldDatas via EntityIdAndHomepageAndEntityType
func GetCommentFieldDatasByEntityIdAndHomepageAndEntityType(offset int, limit int, EntityId_ int, Homepage_ string, EntityType_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("entity_id = ? and homepage = ? and entity_type = ?", EntityId_, Homepage_, EntityType_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByEntityIdAndHomepageAndFieldName Get CommentFieldDatas via EntityIdAndHomepageAndFieldName
func GetCommentFieldDatasByEntityIdAndHomepageAndFieldName(offset int, limit int, EntityId_ int, Homepage_ string, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("entity_id = ? and homepage = ? and field_name = ?", EntityId_, Homepage_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByEntityIdAndHomepageAndDefaultLangcode Get CommentFieldDatas via EntityIdAndHomepageAndDefaultLangcode
func GetCommentFieldDatasByEntityIdAndHomepageAndDefaultLangcode(offset int, limit int, EntityId_ int, Homepage_ string, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("entity_id = ? and homepage = ? and default_langcode = ?", EntityId_, Homepage_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByEntityIdAndHostnameAndCreated Get CommentFieldDatas via EntityIdAndHostnameAndCreated
func GetCommentFieldDatasByEntityIdAndHostnameAndCreated(offset int, limit int, EntityId_ int, Hostname_ string, Created_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("entity_id = ? and hostname = ? and created = ?", EntityId_, Hostname_, Created_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByEntityIdAndHostnameAndChanged Get CommentFieldDatas via EntityIdAndHostnameAndChanged
func GetCommentFieldDatasByEntityIdAndHostnameAndChanged(offset int, limit int, EntityId_ int, Hostname_ string, Changed_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("entity_id = ? and hostname = ? and changed = ?", EntityId_, Hostname_, Changed_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByEntityIdAndHostnameAndStatus Get CommentFieldDatas via EntityIdAndHostnameAndStatus
func GetCommentFieldDatasByEntityIdAndHostnameAndStatus(offset int, limit int, EntityId_ int, Hostname_ string, Status_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("entity_id = ? and hostname = ? and status = ?", EntityId_, Hostname_, Status_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByEntityIdAndHostnameAndThread Get CommentFieldDatas via EntityIdAndHostnameAndThread
func GetCommentFieldDatasByEntityIdAndHostnameAndThread(offset int, limit int, EntityId_ int, Hostname_ string, Thread_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("entity_id = ? and hostname = ? and thread = ?", EntityId_, Hostname_, Thread_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByEntityIdAndHostnameAndEntityType Get CommentFieldDatas via EntityIdAndHostnameAndEntityType
func GetCommentFieldDatasByEntityIdAndHostnameAndEntityType(offset int, limit int, EntityId_ int, Hostname_ string, EntityType_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("entity_id = ? and hostname = ? and entity_type = ?", EntityId_, Hostname_, EntityType_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByEntityIdAndHostnameAndFieldName Get CommentFieldDatas via EntityIdAndHostnameAndFieldName
func GetCommentFieldDatasByEntityIdAndHostnameAndFieldName(offset int, limit int, EntityId_ int, Hostname_ string, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("entity_id = ? and hostname = ? and field_name = ?", EntityId_, Hostname_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByEntityIdAndHostnameAndDefaultLangcode Get CommentFieldDatas via EntityIdAndHostnameAndDefaultLangcode
func GetCommentFieldDatasByEntityIdAndHostnameAndDefaultLangcode(offset int, limit int, EntityId_ int, Hostname_ string, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("entity_id = ? and hostname = ? and default_langcode = ?", EntityId_, Hostname_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByEntityIdAndCreatedAndChanged Get CommentFieldDatas via EntityIdAndCreatedAndChanged
func GetCommentFieldDatasByEntityIdAndCreatedAndChanged(offset int, limit int, EntityId_ int, Created_ int, Changed_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("entity_id = ? and created = ? and changed = ?", EntityId_, Created_, Changed_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByEntityIdAndCreatedAndStatus Get CommentFieldDatas via EntityIdAndCreatedAndStatus
func GetCommentFieldDatasByEntityIdAndCreatedAndStatus(offset int, limit int, EntityId_ int, Created_ int, Status_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("entity_id = ? and created = ? and status = ?", EntityId_, Created_, Status_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByEntityIdAndCreatedAndThread Get CommentFieldDatas via EntityIdAndCreatedAndThread
func GetCommentFieldDatasByEntityIdAndCreatedAndThread(offset int, limit int, EntityId_ int, Created_ int, Thread_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("entity_id = ? and created = ? and thread = ?", EntityId_, Created_, Thread_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByEntityIdAndCreatedAndEntityType Get CommentFieldDatas via EntityIdAndCreatedAndEntityType
func GetCommentFieldDatasByEntityIdAndCreatedAndEntityType(offset int, limit int, EntityId_ int, Created_ int, EntityType_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("entity_id = ? and created = ? and entity_type = ?", EntityId_, Created_, EntityType_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByEntityIdAndCreatedAndFieldName Get CommentFieldDatas via EntityIdAndCreatedAndFieldName
func GetCommentFieldDatasByEntityIdAndCreatedAndFieldName(offset int, limit int, EntityId_ int, Created_ int, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("entity_id = ? and created = ? and field_name = ?", EntityId_, Created_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByEntityIdAndCreatedAndDefaultLangcode Get CommentFieldDatas via EntityIdAndCreatedAndDefaultLangcode
func GetCommentFieldDatasByEntityIdAndCreatedAndDefaultLangcode(offset int, limit int, EntityId_ int, Created_ int, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("entity_id = ? and created = ? and default_langcode = ?", EntityId_, Created_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByEntityIdAndChangedAndStatus Get CommentFieldDatas via EntityIdAndChangedAndStatus
func GetCommentFieldDatasByEntityIdAndChangedAndStatus(offset int, limit int, EntityId_ int, Changed_ int, Status_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("entity_id = ? and changed = ? and status = ?", EntityId_, Changed_, Status_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByEntityIdAndChangedAndThread Get CommentFieldDatas via EntityIdAndChangedAndThread
func GetCommentFieldDatasByEntityIdAndChangedAndThread(offset int, limit int, EntityId_ int, Changed_ int, Thread_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("entity_id = ? and changed = ? and thread = ?", EntityId_, Changed_, Thread_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByEntityIdAndChangedAndEntityType Get CommentFieldDatas via EntityIdAndChangedAndEntityType
func GetCommentFieldDatasByEntityIdAndChangedAndEntityType(offset int, limit int, EntityId_ int, Changed_ int, EntityType_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("entity_id = ? and changed = ? and entity_type = ?", EntityId_, Changed_, EntityType_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByEntityIdAndChangedAndFieldName Get CommentFieldDatas via EntityIdAndChangedAndFieldName
func GetCommentFieldDatasByEntityIdAndChangedAndFieldName(offset int, limit int, EntityId_ int, Changed_ int, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("entity_id = ? and changed = ? and field_name = ?", EntityId_, Changed_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByEntityIdAndChangedAndDefaultLangcode Get CommentFieldDatas via EntityIdAndChangedAndDefaultLangcode
func GetCommentFieldDatasByEntityIdAndChangedAndDefaultLangcode(offset int, limit int, EntityId_ int, Changed_ int, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("entity_id = ? and changed = ? and default_langcode = ?", EntityId_, Changed_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByEntityIdAndStatusAndThread Get CommentFieldDatas via EntityIdAndStatusAndThread
func GetCommentFieldDatasByEntityIdAndStatusAndThread(offset int, limit int, EntityId_ int, Status_ int, Thread_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("entity_id = ? and status = ? and thread = ?", EntityId_, Status_, Thread_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByEntityIdAndStatusAndEntityType Get CommentFieldDatas via EntityIdAndStatusAndEntityType
func GetCommentFieldDatasByEntityIdAndStatusAndEntityType(offset int, limit int, EntityId_ int, Status_ int, EntityType_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("entity_id = ? and status = ? and entity_type = ?", EntityId_, Status_, EntityType_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByEntityIdAndStatusAndFieldName Get CommentFieldDatas via EntityIdAndStatusAndFieldName
func GetCommentFieldDatasByEntityIdAndStatusAndFieldName(offset int, limit int, EntityId_ int, Status_ int, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("entity_id = ? and status = ? and field_name = ?", EntityId_, Status_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByEntityIdAndStatusAndDefaultLangcode Get CommentFieldDatas via EntityIdAndStatusAndDefaultLangcode
func GetCommentFieldDatasByEntityIdAndStatusAndDefaultLangcode(offset int, limit int, EntityId_ int, Status_ int, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("entity_id = ? and status = ? and default_langcode = ?", EntityId_, Status_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByEntityIdAndThreadAndEntityType Get CommentFieldDatas via EntityIdAndThreadAndEntityType
func GetCommentFieldDatasByEntityIdAndThreadAndEntityType(offset int, limit int, EntityId_ int, Thread_ string, EntityType_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("entity_id = ? and thread = ? and entity_type = ?", EntityId_, Thread_, EntityType_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByEntityIdAndThreadAndFieldName Get CommentFieldDatas via EntityIdAndThreadAndFieldName
func GetCommentFieldDatasByEntityIdAndThreadAndFieldName(offset int, limit int, EntityId_ int, Thread_ string, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("entity_id = ? and thread = ? and field_name = ?", EntityId_, Thread_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByEntityIdAndThreadAndDefaultLangcode Get CommentFieldDatas via EntityIdAndThreadAndDefaultLangcode
func GetCommentFieldDatasByEntityIdAndThreadAndDefaultLangcode(offset int, limit int, EntityId_ int, Thread_ string, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("entity_id = ? and thread = ? and default_langcode = ?", EntityId_, Thread_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByEntityIdAndEntityTypeAndFieldName Get CommentFieldDatas via EntityIdAndEntityTypeAndFieldName
func GetCommentFieldDatasByEntityIdAndEntityTypeAndFieldName(offset int, limit int, EntityId_ int, EntityType_ string, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("entity_id = ? and entity_type = ? and field_name = ?", EntityId_, EntityType_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByEntityIdAndEntityTypeAndDefaultLangcode Get CommentFieldDatas via EntityIdAndEntityTypeAndDefaultLangcode
func GetCommentFieldDatasByEntityIdAndEntityTypeAndDefaultLangcode(offset int, limit int, EntityId_ int, EntityType_ string, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("entity_id = ? and entity_type = ? and default_langcode = ?", EntityId_, EntityType_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByEntityIdAndFieldNameAndDefaultLangcode Get CommentFieldDatas via EntityIdAndFieldNameAndDefaultLangcode
func GetCommentFieldDatasByEntityIdAndFieldNameAndDefaultLangcode(offset int, limit int, EntityId_ int, FieldName_ string, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("entity_id = ? and field_name = ? and default_langcode = ?", EntityId_, FieldName_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasBySubjectAndUidAndName Get CommentFieldDatas via SubjectAndUidAndName
func GetCommentFieldDatasBySubjectAndUidAndName(offset int, limit int, Subject_ string, Uid_ int, Name_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("subject = ? and uid = ? and name = ?", Subject_, Uid_, Name_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasBySubjectAndUidAndMail Get CommentFieldDatas via SubjectAndUidAndMail
func GetCommentFieldDatasBySubjectAndUidAndMail(offset int, limit int, Subject_ string, Uid_ int, Mail_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("subject = ? and uid = ? and mail = ?", Subject_, Uid_, Mail_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasBySubjectAndUidAndHomepage Get CommentFieldDatas via SubjectAndUidAndHomepage
func GetCommentFieldDatasBySubjectAndUidAndHomepage(offset int, limit int, Subject_ string, Uid_ int, Homepage_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("subject = ? and uid = ? and homepage = ?", Subject_, Uid_, Homepage_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasBySubjectAndUidAndHostname Get CommentFieldDatas via SubjectAndUidAndHostname
func GetCommentFieldDatasBySubjectAndUidAndHostname(offset int, limit int, Subject_ string, Uid_ int, Hostname_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("subject = ? and uid = ? and hostname = ?", Subject_, Uid_, Hostname_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasBySubjectAndUidAndCreated Get CommentFieldDatas via SubjectAndUidAndCreated
func GetCommentFieldDatasBySubjectAndUidAndCreated(offset int, limit int, Subject_ string, Uid_ int, Created_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("subject = ? and uid = ? and created = ?", Subject_, Uid_, Created_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasBySubjectAndUidAndChanged Get CommentFieldDatas via SubjectAndUidAndChanged
func GetCommentFieldDatasBySubjectAndUidAndChanged(offset int, limit int, Subject_ string, Uid_ int, Changed_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("subject = ? and uid = ? and changed = ?", Subject_, Uid_, Changed_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasBySubjectAndUidAndStatus Get CommentFieldDatas via SubjectAndUidAndStatus
func GetCommentFieldDatasBySubjectAndUidAndStatus(offset int, limit int, Subject_ string, Uid_ int, Status_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("subject = ? and uid = ? and status = ?", Subject_, Uid_, Status_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasBySubjectAndUidAndThread Get CommentFieldDatas via SubjectAndUidAndThread
func GetCommentFieldDatasBySubjectAndUidAndThread(offset int, limit int, Subject_ string, Uid_ int, Thread_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("subject = ? and uid = ? and thread = ?", Subject_, Uid_, Thread_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasBySubjectAndUidAndEntityType Get CommentFieldDatas via SubjectAndUidAndEntityType
func GetCommentFieldDatasBySubjectAndUidAndEntityType(offset int, limit int, Subject_ string, Uid_ int, EntityType_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("subject = ? and uid = ? and entity_type = ?", Subject_, Uid_, EntityType_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasBySubjectAndUidAndFieldName Get CommentFieldDatas via SubjectAndUidAndFieldName
func GetCommentFieldDatasBySubjectAndUidAndFieldName(offset int, limit int, Subject_ string, Uid_ int, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("subject = ? and uid = ? and field_name = ?", Subject_, Uid_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasBySubjectAndUidAndDefaultLangcode Get CommentFieldDatas via SubjectAndUidAndDefaultLangcode
func GetCommentFieldDatasBySubjectAndUidAndDefaultLangcode(offset int, limit int, Subject_ string, Uid_ int, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("subject = ? and uid = ? and default_langcode = ?", Subject_, Uid_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasBySubjectAndNameAndMail Get CommentFieldDatas via SubjectAndNameAndMail
func GetCommentFieldDatasBySubjectAndNameAndMail(offset int, limit int, Subject_ string, Name_ string, Mail_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("subject = ? and name = ? and mail = ?", Subject_, Name_, Mail_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasBySubjectAndNameAndHomepage Get CommentFieldDatas via SubjectAndNameAndHomepage
func GetCommentFieldDatasBySubjectAndNameAndHomepage(offset int, limit int, Subject_ string, Name_ string, Homepage_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("subject = ? and name = ? and homepage = ?", Subject_, Name_, Homepage_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasBySubjectAndNameAndHostname Get CommentFieldDatas via SubjectAndNameAndHostname
func GetCommentFieldDatasBySubjectAndNameAndHostname(offset int, limit int, Subject_ string, Name_ string, Hostname_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("subject = ? and name = ? and hostname = ?", Subject_, Name_, Hostname_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasBySubjectAndNameAndCreated Get CommentFieldDatas via SubjectAndNameAndCreated
func GetCommentFieldDatasBySubjectAndNameAndCreated(offset int, limit int, Subject_ string, Name_ string, Created_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("subject = ? and name = ? and created = ?", Subject_, Name_, Created_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasBySubjectAndNameAndChanged Get CommentFieldDatas via SubjectAndNameAndChanged
func GetCommentFieldDatasBySubjectAndNameAndChanged(offset int, limit int, Subject_ string, Name_ string, Changed_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("subject = ? and name = ? and changed = ?", Subject_, Name_, Changed_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasBySubjectAndNameAndStatus Get CommentFieldDatas via SubjectAndNameAndStatus
func GetCommentFieldDatasBySubjectAndNameAndStatus(offset int, limit int, Subject_ string, Name_ string, Status_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("subject = ? and name = ? and status = ?", Subject_, Name_, Status_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasBySubjectAndNameAndThread Get CommentFieldDatas via SubjectAndNameAndThread
func GetCommentFieldDatasBySubjectAndNameAndThread(offset int, limit int, Subject_ string, Name_ string, Thread_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("subject = ? and name = ? and thread = ?", Subject_, Name_, Thread_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasBySubjectAndNameAndEntityType Get CommentFieldDatas via SubjectAndNameAndEntityType
func GetCommentFieldDatasBySubjectAndNameAndEntityType(offset int, limit int, Subject_ string, Name_ string, EntityType_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("subject = ? and name = ? and entity_type = ?", Subject_, Name_, EntityType_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasBySubjectAndNameAndFieldName Get CommentFieldDatas via SubjectAndNameAndFieldName
func GetCommentFieldDatasBySubjectAndNameAndFieldName(offset int, limit int, Subject_ string, Name_ string, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("subject = ? and name = ? and field_name = ?", Subject_, Name_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasBySubjectAndNameAndDefaultLangcode Get CommentFieldDatas via SubjectAndNameAndDefaultLangcode
func GetCommentFieldDatasBySubjectAndNameAndDefaultLangcode(offset int, limit int, Subject_ string, Name_ string, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("subject = ? and name = ? and default_langcode = ?", Subject_, Name_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasBySubjectAndMailAndHomepage Get CommentFieldDatas via SubjectAndMailAndHomepage
func GetCommentFieldDatasBySubjectAndMailAndHomepage(offset int, limit int, Subject_ string, Mail_ string, Homepage_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("subject = ? and mail = ? and homepage = ?", Subject_, Mail_, Homepage_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasBySubjectAndMailAndHostname Get CommentFieldDatas via SubjectAndMailAndHostname
func GetCommentFieldDatasBySubjectAndMailAndHostname(offset int, limit int, Subject_ string, Mail_ string, Hostname_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("subject = ? and mail = ? and hostname = ?", Subject_, Mail_, Hostname_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasBySubjectAndMailAndCreated Get CommentFieldDatas via SubjectAndMailAndCreated
func GetCommentFieldDatasBySubjectAndMailAndCreated(offset int, limit int, Subject_ string, Mail_ string, Created_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("subject = ? and mail = ? and created = ?", Subject_, Mail_, Created_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasBySubjectAndMailAndChanged Get CommentFieldDatas via SubjectAndMailAndChanged
func GetCommentFieldDatasBySubjectAndMailAndChanged(offset int, limit int, Subject_ string, Mail_ string, Changed_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("subject = ? and mail = ? and changed = ?", Subject_, Mail_, Changed_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasBySubjectAndMailAndStatus Get CommentFieldDatas via SubjectAndMailAndStatus
func GetCommentFieldDatasBySubjectAndMailAndStatus(offset int, limit int, Subject_ string, Mail_ string, Status_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("subject = ? and mail = ? and status = ?", Subject_, Mail_, Status_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasBySubjectAndMailAndThread Get CommentFieldDatas via SubjectAndMailAndThread
func GetCommentFieldDatasBySubjectAndMailAndThread(offset int, limit int, Subject_ string, Mail_ string, Thread_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("subject = ? and mail = ? and thread = ?", Subject_, Mail_, Thread_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasBySubjectAndMailAndEntityType Get CommentFieldDatas via SubjectAndMailAndEntityType
func GetCommentFieldDatasBySubjectAndMailAndEntityType(offset int, limit int, Subject_ string, Mail_ string, EntityType_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("subject = ? and mail = ? and entity_type = ?", Subject_, Mail_, EntityType_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasBySubjectAndMailAndFieldName Get CommentFieldDatas via SubjectAndMailAndFieldName
func GetCommentFieldDatasBySubjectAndMailAndFieldName(offset int, limit int, Subject_ string, Mail_ string, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("subject = ? and mail = ? and field_name = ?", Subject_, Mail_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasBySubjectAndMailAndDefaultLangcode Get CommentFieldDatas via SubjectAndMailAndDefaultLangcode
func GetCommentFieldDatasBySubjectAndMailAndDefaultLangcode(offset int, limit int, Subject_ string, Mail_ string, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("subject = ? and mail = ? and default_langcode = ?", Subject_, Mail_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasBySubjectAndHomepageAndHostname Get CommentFieldDatas via SubjectAndHomepageAndHostname
func GetCommentFieldDatasBySubjectAndHomepageAndHostname(offset int, limit int, Subject_ string, Homepage_ string, Hostname_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("subject = ? and homepage = ? and hostname = ?", Subject_, Homepage_, Hostname_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasBySubjectAndHomepageAndCreated Get CommentFieldDatas via SubjectAndHomepageAndCreated
func GetCommentFieldDatasBySubjectAndHomepageAndCreated(offset int, limit int, Subject_ string, Homepage_ string, Created_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("subject = ? and homepage = ? and created = ?", Subject_, Homepage_, Created_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasBySubjectAndHomepageAndChanged Get CommentFieldDatas via SubjectAndHomepageAndChanged
func GetCommentFieldDatasBySubjectAndHomepageAndChanged(offset int, limit int, Subject_ string, Homepage_ string, Changed_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("subject = ? and homepage = ? and changed = ?", Subject_, Homepage_, Changed_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasBySubjectAndHomepageAndStatus Get CommentFieldDatas via SubjectAndHomepageAndStatus
func GetCommentFieldDatasBySubjectAndHomepageAndStatus(offset int, limit int, Subject_ string, Homepage_ string, Status_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("subject = ? and homepage = ? and status = ?", Subject_, Homepage_, Status_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasBySubjectAndHomepageAndThread Get CommentFieldDatas via SubjectAndHomepageAndThread
func GetCommentFieldDatasBySubjectAndHomepageAndThread(offset int, limit int, Subject_ string, Homepage_ string, Thread_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("subject = ? and homepage = ? and thread = ?", Subject_, Homepage_, Thread_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasBySubjectAndHomepageAndEntityType Get CommentFieldDatas via SubjectAndHomepageAndEntityType
func GetCommentFieldDatasBySubjectAndHomepageAndEntityType(offset int, limit int, Subject_ string, Homepage_ string, EntityType_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("subject = ? and homepage = ? and entity_type = ?", Subject_, Homepage_, EntityType_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasBySubjectAndHomepageAndFieldName Get CommentFieldDatas via SubjectAndHomepageAndFieldName
func GetCommentFieldDatasBySubjectAndHomepageAndFieldName(offset int, limit int, Subject_ string, Homepage_ string, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("subject = ? and homepage = ? and field_name = ?", Subject_, Homepage_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasBySubjectAndHomepageAndDefaultLangcode Get CommentFieldDatas via SubjectAndHomepageAndDefaultLangcode
func GetCommentFieldDatasBySubjectAndHomepageAndDefaultLangcode(offset int, limit int, Subject_ string, Homepage_ string, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("subject = ? and homepage = ? and default_langcode = ?", Subject_, Homepage_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasBySubjectAndHostnameAndCreated Get CommentFieldDatas via SubjectAndHostnameAndCreated
func GetCommentFieldDatasBySubjectAndHostnameAndCreated(offset int, limit int, Subject_ string, Hostname_ string, Created_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("subject = ? and hostname = ? and created = ?", Subject_, Hostname_, Created_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasBySubjectAndHostnameAndChanged Get CommentFieldDatas via SubjectAndHostnameAndChanged
func GetCommentFieldDatasBySubjectAndHostnameAndChanged(offset int, limit int, Subject_ string, Hostname_ string, Changed_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("subject = ? and hostname = ? and changed = ?", Subject_, Hostname_, Changed_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasBySubjectAndHostnameAndStatus Get CommentFieldDatas via SubjectAndHostnameAndStatus
func GetCommentFieldDatasBySubjectAndHostnameAndStatus(offset int, limit int, Subject_ string, Hostname_ string, Status_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("subject = ? and hostname = ? and status = ?", Subject_, Hostname_, Status_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasBySubjectAndHostnameAndThread Get CommentFieldDatas via SubjectAndHostnameAndThread
func GetCommentFieldDatasBySubjectAndHostnameAndThread(offset int, limit int, Subject_ string, Hostname_ string, Thread_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("subject = ? and hostname = ? and thread = ?", Subject_, Hostname_, Thread_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasBySubjectAndHostnameAndEntityType Get CommentFieldDatas via SubjectAndHostnameAndEntityType
func GetCommentFieldDatasBySubjectAndHostnameAndEntityType(offset int, limit int, Subject_ string, Hostname_ string, EntityType_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("subject = ? and hostname = ? and entity_type = ?", Subject_, Hostname_, EntityType_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasBySubjectAndHostnameAndFieldName Get CommentFieldDatas via SubjectAndHostnameAndFieldName
func GetCommentFieldDatasBySubjectAndHostnameAndFieldName(offset int, limit int, Subject_ string, Hostname_ string, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("subject = ? and hostname = ? and field_name = ?", Subject_, Hostname_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasBySubjectAndHostnameAndDefaultLangcode Get CommentFieldDatas via SubjectAndHostnameAndDefaultLangcode
func GetCommentFieldDatasBySubjectAndHostnameAndDefaultLangcode(offset int, limit int, Subject_ string, Hostname_ string, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("subject = ? and hostname = ? and default_langcode = ?", Subject_, Hostname_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasBySubjectAndCreatedAndChanged Get CommentFieldDatas via SubjectAndCreatedAndChanged
func GetCommentFieldDatasBySubjectAndCreatedAndChanged(offset int, limit int, Subject_ string, Created_ int, Changed_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("subject = ? and created = ? and changed = ?", Subject_, Created_, Changed_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasBySubjectAndCreatedAndStatus Get CommentFieldDatas via SubjectAndCreatedAndStatus
func GetCommentFieldDatasBySubjectAndCreatedAndStatus(offset int, limit int, Subject_ string, Created_ int, Status_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("subject = ? and created = ? and status = ?", Subject_, Created_, Status_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasBySubjectAndCreatedAndThread Get CommentFieldDatas via SubjectAndCreatedAndThread
func GetCommentFieldDatasBySubjectAndCreatedAndThread(offset int, limit int, Subject_ string, Created_ int, Thread_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("subject = ? and created = ? and thread = ?", Subject_, Created_, Thread_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasBySubjectAndCreatedAndEntityType Get CommentFieldDatas via SubjectAndCreatedAndEntityType
func GetCommentFieldDatasBySubjectAndCreatedAndEntityType(offset int, limit int, Subject_ string, Created_ int, EntityType_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("subject = ? and created = ? and entity_type = ?", Subject_, Created_, EntityType_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasBySubjectAndCreatedAndFieldName Get CommentFieldDatas via SubjectAndCreatedAndFieldName
func GetCommentFieldDatasBySubjectAndCreatedAndFieldName(offset int, limit int, Subject_ string, Created_ int, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("subject = ? and created = ? and field_name = ?", Subject_, Created_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasBySubjectAndCreatedAndDefaultLangcode Get CommentFieldDatas via SubjectAndCreatedAndDefaultLangcode
func GetCommentFieldDatasBySubjectAndCreatedAndDefaultLangcode(offset int, limit int, Subject_ string, Created_ int, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("subject = ? and created = ? and default_langcode = ?", Subject_, Created_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasBySubjectAndChangedAndStatus Get CommentFieldDatas via SubjectAndChangedAndStatus
func GetCommentFieldDatasBySubjectAndChangedAndStatus(offset int, limit int, Subject_ string, Changed_ int, Status_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("subject = ? and changed = ? and status = ?", Subject_, Changed_, Status_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasBySubjectAndChangedAndThread Get CommentFieldDatas via SubjectAndChangedAndThread
func GetCommentFieldDatasBySubjectAndChangedAndThread(offset int, limit int, Subject_ string, Changed_ int, Thread_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("subject = ? and changed = ? and thread = ?", Subject_, Changed_, Thread_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasBySubjectAndChangedAndEntityType Get CommentFieldDatas via SubjectAndChangedAndEntityType
func GetCommentFieldDatasBySubjectAndChangedAndEntityType(offset int, limit int, Subject_ string, Changed_ int, EntityType_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("subject = ? and changed = ? and entity_type = ?", Subject_, Changed_, EntityType_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasBySubjectAndChangedAndFieldName Get CommentFieldDatas via SubjectAndChangedAndFieldName
func GetCommentFieldDatasBySubjectAndChangedAndFieldName(offset int, limit int, Subject_ string, Changed_ int, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("subject = ? and changed = ? and field_name = ?", Subject_, Changed_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasBySubjectAndChangedAndDefaultLangcode Get CommentFieldDatas via SubjectAndChangedAndDefaultLangcode
func GetCommentFieldDatasBySubjectAndChangedAndDefaultLangcode(offset int, limit int, Subject_ string, Changed_ int, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("subject = ? and changed = ? and default_langcode = ?", Subject_, Changed_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasBySubjectAndStatusAndThread Get CommentFieldDatas via SubjectAndStatusAndThread
func GetCommentFieldDatasBySubjectAndStatusAndThread(offset int, limit int, Subject_ string, Status_ int, Thread_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("subject = ? and status = ? and thread = ?", Subject_, Status_, Thread_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasBySubjectAndStatusAndEntityType Get CommentFieldDatas via SubjectAndStatusAndEntityType
func GetCommentFieldDatasBySubjectAndStatusAndEntityType(offset int, limit int, Subject_ string, Status_ int, EntityType_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("subject = ? and status = ? and entity_type = ?", Subject_, Status_, EntityType_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasBySubjectAndStatusAndFieldName Get CommentFieldDatas via SubjectAndStatusAndFieldName
func GetCommentFieldDatasBySubjectAndStatusAndFieldName(offset int, limit int, Subject_ string, Status_ int, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("subject = ? and status = ? and field_name = ?", Subject_, Status_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasBySubjectAndStatusAndDefaultLangcode Get CommentFieldDatas via SubjectAndStatusAndDefaultLangcode
func GetCommentFieldDatasBySubjectAndStatusAndDefaultLangcode(offset int, limit int, Subject_ string, Status_ int, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("subject = ? and status = ? and default_langcode = ?", Subject_, Status_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasBySubjectAndThreadAndEntityType Get CommentFieldDatas via SubjectAndThreadAndEntityType
func GetCommentFieldDatasBySubjectAndThreadAndEntityType(offset int, limit int, Subject_ string, Thread_ string, EntityType_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("subject = ? and thread = ? and entity_type = ?", Subject_, Thread_, EntityType_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasBySubjectAndThreadAndFieldName Get CommentFieldDatas via SubjectAndThreadAndFieldName
func GetCommentFieldDatasBySubjectAndThreadAndFieldName(offset int, limit int, Subject_ string, Thread_ string, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("subject = ? and thread = ? and field_name = ?", Subject_, Thread_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasBySubjectAndThreadAndDefaultLangcode Get CommentFieldDatas via SubjectAndThreadAndDefaultLangcode
func GetCommentFieldDatasBySubjectAndThreadAndDefaultLangcode(offset int, limit int, Subject_ string, Thread_ string, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("subject = ? and thread = ? and default_langcode = ?", Subject_, Thread_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasBySubjectAndEntityTypeAndFieldName Get CommentFieldDatas via SubjectAndEntityTypeAndFieldName
func GetCommentFieldDatasBySubjectAndEntityTypeAndFieldName(offset int, limit int, Subject_ string, EntityType_ string, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("subject = ? and entity_type = ? and field_name = ?", Subject_, EntityType_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasBySubjectAndEntityTypeAndDefaultLangcode Get CommentFieldDatas via SubjectAndEntityTypeAndDefaultLangcode
func GetCommentFieldDatasBySubjectAndEntityTypeAndDefaultLangcode(offset int, limit int, Subject_ string, EntityType_ string, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("subject = ? and entity_type = ? and default_langcode = ?", Subject_, EntityType_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasBySubjectAndFieldNameAndDefaultLangcode Get CommentFieldDatas via SubjectAndFieldNameAndDefaultLangcode
func GetCommentFieldDatasBySubjectAndFieldNameAndDefaultLangcode(offset int, limit int, Subject_ string, FieldName_ string, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("subject = ? and field_name = ? and default_langcode = ?", Subject_, FieldName_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByUidAndNameAndMail Get CommentFieldDatas via UidAndNameAndMail
func GetCommentFieldDatasByUidAndNameAndMail(offset int, limit int, Uid_ int, Name_ string, Mail_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("uid = ? and name = ? and mail = ?", Uid_, Name_, Mail_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByUidAndNameAndHomepage Get CommentFieldDatas via UidAndNameAndHomepage
func GetCommentFieldDatasByUidAndNameAndHomepage(offset int, limit int, Uid_ int, Name_ string, Homepage_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("uid = ? and name = ? and homepage = ?", Uid_, Name_, Homepage_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByUidAndNameAndHostname Get CommentFieldDatas via UidAndNameAndHostname
func GetCommentFieldDatasByUidAndNameAndHostname(offset int, limit int, Uid_ int, Name_ string, Hostname_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("uid = ? and name = ? and hostname = ?", Uid_, Name_, Hostname_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByUidAndNameAndCreated Get CommentFieldDatas via UidAndNameAndCreated
func GetCommentFieldDatasByUidAndNameAndCreated(offset int, limit int, Uid_ int, Name_ string, Created_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("uid = ? and name = ? and created = ?", Uid_, Name_, Created_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByUidAndNameAndChanged Get CommentFieldDatas via UidAndNameAndChanged
func GetCommentFieldDatasByUidAndNameAndChanged(offset int, limit int, Uid_ int, Name_ string, Changed_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("uid = ? and name = ? and changed = ?", Uid_, Name_, Changed_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByUidAndNameAndStatus Get CommentFieldDatas via UidAndNameAndStatus
func GetCommentFieldDatasByUidAndNameAndStatus(offset int, limit int, Uid_ int, Name_ string, Status_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("uid = ? and name = ? and status = ?", Uid_, Name_, Status_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByUidAndNameAndThread Get CommentFieldDatas via UidAndNameAndThread
func GetCommentFieldDatasByUidAndNameAndThread(offset int, limit int, Uid_ int, Name_ string, Thread_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("uid = ? and name = ? and thread = ?", Uid_, Name_, Thread_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByUidAndNameAndEntityType Get CommentFieldDatas via UidAndNameAndEntityType
func GetCommentFieldDatasByUidAndNameAndEntityType(offset int, limit int, Uid_ int, Name_ string, EntityType_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("uid = ? and name = ? and entity_type = ?", Uid_, Name_, EntityType_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByUidAndNameAndFieldName Get CommentFieldDatas via UidAndNameAndFieldName
func GetCommentFieldDatasByUidAndNameAndFieldName(offset int, limit int, Uid_ int, Name_ string, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("uid = ? and name = ? and field_name = ?", Uid_, Name_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByUidAndNameAndDefaultLangcode Get CommentFieldDatas via UidAndNameAndDefaultLangcode
func GetCommentFieldDatasByUidAndNameAndDefaultLangcode(offset int, limit int, Uid_ int, Name_ string, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("uid = ? and name = ? and default_langcode = ?", Uid_, Name_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByUidAndMailAndHomepage Get CommentFieldDatas via UidAndMailAndHomepage
func GetCommentFieldDatasByUidAndMailAndHomepage(offset int, limit int, Uid_ int, Mail_ string, Homepage_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("uid = ? and mail = ? and homepage = ?", Uid_, Mail_, Homepage_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByUidAndMailAndHostname Get CommentFieldDatas via UidAndMailAndHostname
func GetCommentFieldDatasByUidAndMailAndHostname(offset int, limit int, Uid_ int, Mail_ string, Hostname_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("uid = ? and mail = ? and hostname = ?", Uid_, Mail_, Hostname_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByUidAndMailAndCreated Get CommentFieldDatas via UidAndMailAndCreated
func GetCommentFieldDatasByUidAndMailAndCreated(offset int, limit int, Uid_ int, Mail_ string, Created_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("uid = ? and mail = ? and created = ?", Uid_, Mail_, Created_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByUidAndMailAndChanged Get CommentFieldDatas via UidAndMailAndChanged
func GetCommentFieldDatasByUidAndMailAndChanged(offset int, limit int, Uid_ int, Mail_ string, Changed_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("uid = ? and mail = ? and changed = ?", Uid_, Mail_, Changed_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByUidAndMailAndStatus Get CommentFieldDatas via UidAndMailAndStatus
func GetCommentFieldDatasByUidAndMailAndStatus(offset int, limit int, Uid_ int, Mail_ string, Status_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("uid = ? and mail = ? and status = ?", Uid_, Mail_, Status_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByUidAndMailAndThread Get CommentFieldDatas via UidAndMailAndThread
func GetCommentFieldDatasByUidAndMailAndThread(offset int, limit int, Uid_ int, Mail_ string, Thread_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("uid = ? and mail = ? and thread = ?", Uid_, Mail_, Thread_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByUidAndMailAndEntityType Get CommentFieldDatas via UidAndMailAndEntityType
func GetCommentFieldDatasByUidAndMailAndEntityType(offset int, limit int, Uid_ int, Mail_ string, EntityType_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("uid = ? and mail = ? and entity_type = ?", Uid_, Mail_, EntityType_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByUidAndMailAndFieldName Get CommentFieldDatas via UidAndMailAndFieldName
func GetCommentFieldDatasByUidAndMailAndFieldName(offset int, limit int, Uid_ int, Mail_ string, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("uid = ? and mail = ? and field_name = ?", Uid_, Mail_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByUidAndMailAndDefaultLangcode Get CommentFieldDatas via UidAndMailAndDefaultLangcode
func GetCommentFieldDatasByUidAndMailAndDefaultLangcode(offset int, limit int, Uid_ int, Mail_ string, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("uid = ? and mail = ? and default_langcode = ?", Uid_, Mail_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByUidAndHomepageAndHostname Get CommentFieldDatas via UidAndHomepageAndHostname
func GetCommentFieldDatasByUidAndHomepageAndHostname(offset int, limit int, Uid_ int, Homepage_ string, Hostname_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("uid = ? and homepage = ? and hostname = ?", Uid_, Homepage_, Hostname_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByUidAndHomepageAndCreated Get CommentFieldDatas via UidAndHomepageAndCreated
func GetCommentFieldDatasByUidAndHomepageAndCreated(offset int, limit int, Uid_ int, Homepage_ string, Created_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("uid = ? and homepage = ? and created = ?", Uid_, Homepage_, Created_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByUidAndHomepageAndChanged Get CommentFieldDatas via UidAndHomepageAndChanged
func GetCommentFieldDatasByUidAndHomepageAndChanged(offset int, limit int, Uid_ int, Homepage_ string, Changed_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("uid = ? and homepage = ? and changed = ?", Uid_, Homepage_, Changed_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByUidAndHomepageAndStatus Get CommentFieldDatas via UidAndHomepageAndStatus
func GetCommentFieldDatasByUidAndHomepageAndStatus(offset int, limit int, Uid_ int, Homepage_ string, Status_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("uid = ? and homepage = ? and status = ?", Uid_, Homepage_, Status_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByUidAndHomepageAndThread Get CommentFieldDatas via UidAndHomepageAndThread
func GetCommentFieldDatasByUidAndHomepageAndThread(offset int, limit int, Uid_ int, Homepage_ string, Thread_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("uid = ? and homepage = ? and thread = ?", Uid_, Homepage_, Thread_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByUidAndHomepageAndEntityType Get CommentFieldDatas via UidAndHomepageAndEntityType
func GetCommentFieldDatasByUidAndHomepageAndEntityType(offset int, limit int, Uid_ int, Homepage_ string, EntityType_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("uid = ? and homepage = ? and entity_type = ?", Uid_, Homepage_, EntityType_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByUidAndHomepageAndFieldName Get CommentFieldDatas via UidAndHomepageAndFieldName
func GetCommentFieldDatasByUidAndHomepageAndFieldName(offset int, limit int, Uid_ int, Homepage_ string, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("uid = ? and homepage = ? and field_name = ?", Uid_, Homepage_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByUidAndHomepageAndDefaultLangcode Get CommentFieldDatas via UidAndHomepageAndDefaultLangcode
func GetCommentFieldDatasByUidAndHomepageAndDefaultLangcode(offset int, limit int, Uid_ int, Homepage_ string, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("uid = ? and homepage = ? and default_langcode = ?", Uid_, Homepage_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByUidAndHostnameAndCreated Get CommentFieldDatas via UidAndHostnameAndCreated
func GetCommentFieldDatasByUidAndHostnameAndCreated(offset int, limit int, Uid_ int, Hostname_ string, Created_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("uid = ? and hostname = ? and created = ?", Uid_, Hostname_, Created_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByUidAndHostnameAndChanged Get CommentFieldDatas via UidAndHostnameAndChanged
func GetCommentFieldDatasByUidAndHostnameAndChanged(offset int, limit int, Uid_ int, Hostname_ string, Changed_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("uid = ? and hostname = ? and changed = ?", Uid_, Hostname_, Changed_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByUidAndHostnameAndStatus Get CommentFieldDatas via UidAndHostnameAndStatus
func GetCommentFieldDatasByUidAndHostnameAndStatus(offset int, limit int, Uid_ int, Hostname_ string, Status_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("uid = ? and hostname = ? and status = ?", Uid_, Hostname_, Status_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByUidAndHostnameAndThread Get CommentFieldDatas via UidAndHostnameAndThread
func GetCommentFieldDatasByUidAndHostnameAndThread(offset int, limit int, Uid_ int, Hostname_ string, Thread_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("uid = ? and hostname = ? and thread = ?", Uid_, Hostname_, Thread_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByUidAndHostnameAndEntityType Get CommentFieldDatas via UidAndHostnameAndEntityType
func GetCommentFieldDatasByUidAndHostnameAndEntityType(offset int, limit int, Uid_ int, Hostname_ string, EntityType_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("uid = ? and hostname = ? and entity_type = ?", Uid_, Hostname_, EntityType_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByUidAndHostnameAndFieldName Get CommentFieldDatas via UidAndHostnameAndFieldName
func GetCommentFieldDatasByUidAndHostnameAndFieldName(offset int, limit int, Uid_ int, Hostname_ string, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("uid = ? and hostname = ? and field_name = ?", Uid_, Hostname_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByUidAndHostnameAndDefaultLangcode Get CommentFieldDatas via UidAndHostnameAndDefaultLangcode
func GetCommentFieldDatasByUidAndHostnameAndDefaultLangcode(offset int, limit int, Uid_ int, Hostname_ string, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("uid = ? and hostname = ? and default_langcode = ?", Uid_, Hostname_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByUidAndCreatedAndChanged Get CommentFieldDatas via UidAndCreatedAndChanged
func GetCommentFieldDatasByUidAndCreatedAndChanged(offset int, limit int, Uid_ int, Created_ int, Changed_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("uid = ? and created = ? and changed = ?", Uid_, Created_, Changed_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByUidAndCreatedAndStatus Get CommentFieldDatas via UidAndCreatedAndStatus
func GetCommentFieldDatasByUidAndCreatedAndStatus(offset int, limit int, Uid_ int, Created_ int, Status_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("uid = ? and created = ? and status = ?", Uid_, Created_, Status_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByUidAndCreatedAndThread Get CommentFieldDatas via UidAndCreatedAndThread
func GetCommentFieldDatasByUidAndCreatedAndThread(offset int, limit int, Uid_ int, Created_ int, Thread_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("uid = ? and created = ? and thread = ?", Uid_, Created_, Thread_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByUidAndCreatedAndEntityType Get CommentFieldDatas via UidAndCreatedAndEntityType
func GetCommentFieldDatasByUidAndCreatedAndEntityType(offset int, limit int, Uid_ int, Created_ int, EntityType_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("uid = ? and created = ? and entity_type = ?", Uid_, Created_, EntityType_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByUidAndCreatedAndFieldName Get CommentFieldDatas via UidAndCreatedAndFieldName
func GetCommentFieldDatasByUidAndCreatedAndFieldName(offset int, limit int, Uid_ int, Created_ int, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("uid = ? and created = ? and field_name = ?", Uid_, Created_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByUidAndCreatedAndDefaultLangcode Get CommentFieldDatas via UidAndCreatedAndDefaultLangcode
func GetCommentFieldDatasByUidAndCreatedAndDefaultLangcode(offset int, limit int, Uid_ int, Created_ int, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("uid = ? and created = ? and default_langcode = ?", Uid_, Created_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByUidAndChangedAndStatus Get CommentFieldDatas via UidAndChangedAndStatus
func GetCommentFieldDatasByUidAndChangedAndStatus(offset int, limit int, Uid_ int, Changed_ int, Status_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("uid = ? and changed = ? and status = ?", Uid_, Changed_, Status_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByUidAndChangedAndThread Get CommentFieldDatas via UidAndChangedAndThread
func GetCommentFieldDatasByUidAndChangedAndThread(offset int, limit int, Uid_ int, Changed_ int, Thread_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("uid = ? and changed = ? and thread = ?", Uid_, Changed_, Thread_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByUidAndChangedAndEntityType Get CommentFieldDatas via UidAndChangedAndEntityType
func GetCommentFieldDatasByUidAndChangedAndEntityType(offset int, limit int, Uid_ int, Changed_ int, EntityType_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("uid = ? and changed = ? and entity_type = ?", Uid_, Changed_, EntityType_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByUidAndChangedAndFieldName Get CommentFieldDatas via UidAndChangedAndFieldName
func GetCommentFieldDatasByUidAndChangedAndFieldName(offset int, limit int, Uid_ int, Changed_ int, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("uid = ? and changed = ? and field_name = ?", Uid_, Changed_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByUidAndChangedAndDefaultLangcode Get CommentFieldDatas via UidAndChangedAndDefaultLangcode
func GetCommentFieldDatasByUidAndChangedAndDefaultLangcode(offset int, limit int, Uid_ int, Changed_ int, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("uid = ? and changed = ? and default_langcode = ?", Uid_, Changed_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByUidAndStatusAndThread Get CommentFieldDatas via UidAndStatusAndThread
func GetCommentFieldDatasByUidAndStatusAndThread(offset int, limit int, Uid_ int, Status_ int, Thread_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("uid = ? and status = ? and thread = ?", Uid_, Status_, Thread_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByUidAndStatusAndEntityType Get CommentFieldDatas via UidAndStatusAndEntityType
func GetCommentFieldDatasByUidAndStatusAndEntityType(offset int, limit int, Uid_ int, Status_ int, EntityType_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("uid = ? and status = ? and entity_type = ?", Uid_, Status_, EntityType_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByUidAndStatusAndFieldName Get CommentFieldDatas via UidAndStatusAndFieldName
func GetCommentFieldDatasByUidAndStatusAndFieldName(offset int, limit int, Uid_ int, Status_ int, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("uid = ? and status = ? and field_name = ?", Uid_, Status_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByUidAndStatusAndDefaultLangcode Get CommentFieldDatas via UidAndStatusAndDefaultLangcode
func GetCommentFieldDatasByUidAndStatusAndDefaultLangcode(offset int, limit int, Uid_ int, Status_ int, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("uid = ? and status = ? and default_langcode = ?", Uid_, Status_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByUidAndThreadAndEntityType Get CommentFieldDatas via UidAndThreadAndEntityType
func GetCommentFieldDatasByUidAndThreadAndEntityType(offset int, limit int, Uid_ int, Thread_ string, EntityType_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("uid = ? and thread = ? and entity_type = ?", Uid_, Thread_, EntityType_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByUidAndThreadAndFieldName Get CommentFieldDatas via UidAndThreadAndFieldName
func GetCommentFieldDatasByUidAndThreadAndFieldName(offset int, limit int, Uid_ int, Thread_ string, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("uid = ? and thread = ? and field_name = ?", Uid_, Thread_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByUidAndThreadAndDefaultLangcode Get CommentFieldDatas via UidAndThreadAndDefaultLangcode
func GetCommentFieldDatasByUidAndThreadAndDefaultLangcode(offset int, limit int, Uid_ int, Thread_ string, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("uid = ? and thread = ? and default_langcode = ?", Uid_, Thread_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByUidAndEntityTypeAndFieldName Get CommentFieldDatas via UidAndEntityTypeAndFieldName
func GetCommentFieldDatasByUidAndEntityTypeAndFieldName(offset int, limit int, Uid_ int, EntityType_ string, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("uid = ? and entity_type = ? and field_name = ?", Uid_, EntityType_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByUidAndEntityTypeAndDefaultLangcode Get CommentFieldDatas via UidAndEntityTypeAndDefaultLangcode
func GetCommentFieldDatasByUidAndEntityTypeAndDefaultLangcode(offset int, limit int, Uid_ int, EntityType_ string, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("uid = ? and entity_type = ? and default_langcode = ?", Uid_, EntityType_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByUidAndFieldNameAndDefaultLangcode Get CommentFieldDatas via UidAndFieldNameAndDefaultLangcode
func GetCommentFieldDatasByUidAndFieldNameAndDefaultLangcode(offset int, limit int, Uid_ int, FieldName_ string, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("uid = ? and field_name = ? and default_langcode = ?", Uid_, FieldName_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByNameAndMailAndHomepage Get CommentFieldDatas via NameAndMailAndHomepage
func GetCommentFieldDatasByNameAndMailAndHomepage(offset int, limit int, Name_ string, Mail_ string, Homepage_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("name = ? and mail = ? and homepage = ?", Name_, Mail_, Homepage_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByNameAndMailAndHostname Get CommentFieldDatas via NameAndMailAndHostname
func GetCommentFieldDatasByNameAndMailAndHostname(offset int, limit int, Name_ string, Mail_ string, Hostname_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("name = ? and mail = ? and hostname = ?", Name_, Mail_, Hostname_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByNameAndMailAndCreated Get CommentFieldDatas via NameAndMailAndCreated
func GetCommentFieldDatasByNameAndMailAndCreated(offset int, limit int, Name_ string, Mail_ string, Created_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("name = ? and mail = ? and created = ?", Name_, Mail_, Created_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByNameAndMailAndChanged Get CommentFieldDatas via NameAndMailAndChanged
func GetCommentFieldDatasByNameAndMailAndChanged(offset int, limit int, Name_ string, Mail_ string, Changed_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("name = ? and mail = ? and changed = ?", Name_, Mail_, Changed_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByNameAndMailAndStatus Get CommentFieldDatas via NameAndMailAndStatus
func GetCommentFieldDatasByNameAndMailAndStatus(offset int, limit int, Name_ string, Mail_ string, Status_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("name = ? and mail = ? and status = ?", Name_, Mail_, Status_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByNameAndMailAndThread Get CommentFieldDatas via NameAndMailAndThread
func GetCommentFieldDatasByNameAndMailAndThread(offset int, limit int, Name_ string, Mail_ string, Thread_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("name = ? and mail = ? and thread = ?", Name_, Mail_, Thread_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByNameAndMailAndEntityType Get CommentFieldDatas via NameAndMailAndEntityType
func GetCommentFieldDatasByNameAndMailAndEntityType(offset int, limit int, Name_ string, Mail_ string, EntityType_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("name = ? and mail = ? and entity_type = ?", Name_, Mail_, EntityType_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByNameAndMailAndFieldName Get CommentFieldDatas via NameAndMailAndFieldName
func GetCommentFieldDatasByNameAndMailAndFieldName(offset int, limit int, Name_ string, Mail_ string, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("name = ? and mail = ? and field_name = ?", Name_, Mail_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByNameAndMailAndDefaultLangcode Get CommentFieldDatas via NameAndMailAndDefaultLangcode
func GetCommentFieldDatasByNameAndMailAndDefaultLangcode(offset int, limit int, Name_ string, Mail_ string, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("name = ? and mail = ? and default_langcode = ?", Name_, Mail_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByNameAndHomepageAndHostname Get CommentFieldDatas via NameAndHomepageAndHostname
func GetCommentFieldDatasByNameAndHomepageAndHostname(offset int, limit int, Name_ string, Homepage_ string, Hostname_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("name = ? and homepage = ? and hostname = ?", Name_, Homepage_, Hostname_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByNameAndHomepageAndCreated Get CommentFieldDatas via NameAndHomepageAndCreated
func GetCommentFieldDatasByNameAndHomepageAndCreated(offset int, limit int, Name_ string, Homepage_ string, Created_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("name = ? and homepage = ? and created = ?", Name_, Homepage_, Created_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByNameAndHomepageAndChanged Get CommentFieldDatas via NameAndHomepageAndChanged
func GetCommentFieldDatasByNameAndHomepageAndChanged(offset int, limit int, Name_ string, Homepage_ string, Changed_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("name = ? and homepage = ? and changed = ?", Name_, Homepage_, Changed_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByNameAndHomepageAndStatus Get CommentFieldDatas via NameAndHomepageAndStatus
func GetCommentFieldDatasByNameAndHomepageAndStatus(offset int, limit int, Name_ string, Homepage_ string, Status_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("name = ? and homepage = ? and status = ?", Name_, Homepage_, Status_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByNameAndHomepageAndThread Get CommentFieldDatas via NameAndHomepageAndThread
func GetCommentFieldDatasByNameAndHomepageAndThread(offset int, limit int, Name_ string, Homepage_ string, Thread_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("name = ? and homepage = ? and thread = ?", Name_, Homepage_, Thread_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByNameAndHomepageAndEntityType Get CommentFieldDatas via NameAndHomepageAndEntityType
func GetCommentFieldDatasByNameAndHomepageAndEntityType(offset int, limit int, Name_ string, Homepage_ string, EntityType_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("name = ? and homepage = ? and entity_type = ?", Name_, Homepage_, EntityType_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByNameAndHomepageAndFieldName Get CommentFieldDatas via NameAndHomepageAndFieldName
func GetCommentFieldDatasByNameAndHomepageAndFieldName(offset int, limit int, Name_ string, Homepage_ string, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("name = ? and homepage = ? and field_name = ?", Name_, Homepage_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByNameAndHomepageAndDefaultLangcode Get CommentFieldDatas via NameAndHomepageAndDefaultLangcode
func GetCommentFieldDatasByNameAndHomepageAndDefaultLangcode(offset int, limit int, Name_ string, Homepage_ string, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("name = ? and homepage = ? and default_langcode = ?", Name_, Homepage_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByNameAndHostnameAndCreated Get CommentFieldDatas via NameAndHostnameAndCreated
func GetCommentFieldDatasByNameAndHostnameAndCreated(offset int, limit int, Name_ string, Hostname_ string, Created_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("name = ? and hostname = ? and created = ?", Name_, Hostname_, Created_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByNameAndHostnameAndChanged Get CommentFieldDatas via NameAndHostnameAndChanged
func GetCommentFieldDatasByNameAndHostnameAndChanged(offset int, limit int, Name_ string, Hostname_ string, Changed_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("name = ? and hostname = ? and changed = ?", Name_, Hostname_, Changed_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByNameAndHostnameAndStatus Get CommentFieldDatas via NameAndHostnameAndStatus
func GetCommentFieldDatasByNameAndHostnameAndStatus(offset int, limit int, Name_ string, Hostname_ string, Status_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("name = ? and hostname = ? and status = ?", Name_, Hostname_, Status_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByNameAndHostnameAndThread Get CommentFieldDatas via NameAndHostnameAndThread
func GetCommentFieldDatasByNameAndHostnameAndThread(offset int, limit int, Name_ string, Hostname_ string, Thread_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("name = ? and hostname = ? and thread = ?", Name_, Hostname_, Thread_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByNameAndHostnameAndEntityType Get CommentFieldDatas via NameAndHostnameAndEntityType
func GetCommentFieldDatasByNameAndHostnameAndEntityType(offset int, limit int, Name_ string, Hostname_ string, EntityType_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("name = ? and hostname = ? and entity_type = ?", Name_, Hostname_, EntityType_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByNameAndHostnameAndFieldName Get CommentFieldDatas via NameAndHostnameAndFieldName
func GetCommentFieldDatasByNameAndHostnameAndFieldName(offset int, limit int, Name_ string, Hostname_ string, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("name = ? and hostname = ? and field_name = ?", Name_, Hostname_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByNameAndHostnameAndDefaultLangcode Get CommentFieldDatas via NameAndHostnameAndDefaultLangcode
func GetCommentFieldDatasByNameAndHostnameAndDefaultLangcode(offset int, limit int, Name_ string, Hostname_ string, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("name = ? and hostname = ? and default_langcode = ?", Name_, Hostname_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByNameAndCreatedAndChanged Get CommentFieldDatas via NameAndCreatedAndChanged
func GetCommentFieldDatasByNameAndCreatedAndChanged(offset int, limit int, Name_ string, Created_ int, Changed_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("name = ? and created = ? and changed = ?", Name_, Created_, Changed_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByNameAndCreatedAndStatus Get CommentFieldDatas via NameAndCreatedAndStatus
func GetCommentFieldDatasByNameAndCreatedAndStatus(offset int, limit int, Name_ string, Created_ int, Status_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("name = ? and created = ? and status = ?", Name_, Created_, Status_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByNameAndCreatedAndThread Get CommentFieldDatas via NameAndCreatedAndThread
func GetCommentFieldDatasByNameAndCreatedAndThread(offset int, limit int, Name_ string, Created_ int, Thread_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("name = ? and created = ? and thread = ?", Name_, Created_, Thread_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByNameAndCreatedAndEntityType Get CommentFieldDatas via NameAndCreatedAndEntityType
func GetCommentFieldDatasByNameAndCreatedAndEntityType(offset int, limit int, Name_ string, Created_ int, EntityType_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("name = ? and created = ? and entity_type = ?", Name_, Created_, EntityType_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByNameAndCreatedAndFieldName Get CommentFieldDatas via NameAndCreatedAndFieldName
func GetCommentFieldDatasByNameAndCreatedAndFieldName(offset int, limit int, Name_ string, Created_ int, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("name = ? and created = ? and field_name = ?", Name_, Created_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByNameAndCreatedAndDefaultLangcode Get CommentFieldDatas via NameAndCreatedAndDefaultLangcode
func GetCommentFieldDatasByNameAndCreatedAndDefaultLangcode(offset int, limit int, Name_ string, Created_ int, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("name = ? and created = ? and default_langcode = ?", Name_, Created_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByNameAndChangedAndStatus Get CommentFieldDatas via NameAndChangedAndStatus
func GetCommentFieldDatasByNameAndChangedAndStatus(offset int, limit int, Name_ string, Changed_ int, Status_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("name = ? and changed = ? and status = ?", Name_, Changed_, Status_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByNameAndChangedAndThread Get CommentFieldDatas via NameAndChangedAndThread
func GetCommentFieldDatasByNameAndChangedAndThread(offset int, limit int, Name_ string, Changed_ int, Thread_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("name = ? and changed = ? and thread = ?", Name_, Changed_, Thread_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByNameAndChangedAndEntityType Get CommentFieldDatas via NameAndChangedAndEntityType
func GetCommentFieldDatasByNameAndChangedAndEntityType(offset int, limit int, Name_ string, Changed_ int, EntityType_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("name = ? and changed = ? and entity_type = ?", Name_, Changed_, EntityType_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByNameAndChangedAndFieldName Get CommentFieldDatas via NameAndChangedAndFieldName
func GetCommentFieldDatasByNameAndChangedAndFieldName(offset int, limit int, Name_ string, Changed_ int, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("name = ? and changed = ? and field_name = ?", Name_, Changed_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByNameAndChangedAndDefaultLangcode Get CommentFieldDatas via NameAndChangedAndDefaultLangcode
func GetCommentFieldDatasByNameAndChangedAndDefaultLangcode(offset int, limit int, Name_ string, Changed_ int, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("name = ? and changed = ? and default_langcode = ?", Name_, Changed_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByNameAndStatusAndThread Get CommentFieldDatas via NameAndStatusAndThread
func GetCommentFieldDatasByNameAndStatusAndThread(offset int, limit int, Name_ string, Status_ int, Thread_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("name = ? and status = ? and thread = ?", Name_, Status_, Thread_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByNameAndStatusAndEntityType Get CommentFieldDatas via NameAndStatusAndEntityType
func GetCommentFieldDatasByNameAndStatusAndEntityType(offset int, limit int, Name_ string, Status_ int, EntityType_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("name = ? and status = ? and entity_type = ?", Name_, Status_, EntityType_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByNameAndStatusAndFieldName Get CommentFieldDatas via NameAndStatusAndFieldName
func GetCommentFieldDatasByNameAndStatusAndFieldName(offset int, limit int, Name_ string, Status_ int, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("name = ? and status = ? and field_name = ?", Name_, Status_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByNameAndStatusAndDefaultLangcode Get CommentFieldDatas via NameAndStatusAndDefaultLangcode
func GetCommentFieldDatasByNameAndStatusAndDefaultLangcode(offset int, limit int, Name_ string, Status_ int, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("name = ? and status = ? and default_langcode = ?", Name_, Status_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByNameAndThreadAndEntityType Get CommentFieldDatas via NameAndThreadAndEntityType
func GetCommentFieldDatasByNameAndThreadAndEntityType(offset int, limit int, Name_ string, Thread_ string, EntityType_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("name = ? and thread = ? and entity_type = ?", Name_, Thread_, EntityType_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByNameAndThreadAndFieldName Get CommentFieldDatas via NameAndThreadAndFieldName
func GetCommentFieldDatasByNameAndThreadAndFieldName(offset int, limit int, Name_ string, Thread_ string, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("name = ? and thread = ? and field_name = ?", Name_, Thread_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByNameAndThreadAndDefaultLangcode Get CommentFieldDatas via NameAndThreadAndDefaultLangcode
func GetCommentFieldDatasByNameAndThreadAndDefaultLangcode(offset int, limit int, Name_ string, Thread_ string, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("name = ? and thread = ? and default_langcode = ?", Name_, Thread_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByNameAndEntityTypeAndFieldName Get CommentFieldDatas via NameAndEntityTypeAndFieldName
func GetCommentFieldDatasByNameAndEntityTypeAndFieldName(offset int, limit int, Name_ string, EntityType_ string, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("name = ? and entity_type = ? and field_name = ?", Name_, EntityType_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByNameAndEntityTypeAndDefaultLangcode Get CommentFieldDatas via NameAndEntityTypeAndDefaultLangcode
func GetCommentFieldDatasByNameAndEntityTypeAndDefaultLangcode(offset int, limit int, Name_ string, EntityType_ string, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("name = ? and entity_type = ? and default_langcode = ?", Name_, EntityType_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByNameAndFieldNameAndDefaultLangcode Get CommentFieldDatas via NameAndFieldNameAndDefaultLangcode
func GetCommentFieldDatasByNameAndFieldNameAndDefaultLangcode(offset int, limit int, Name_ string, FieldName_ string, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("name = ? and field_name = ? and default_langcode = ?", Name_, FieldName_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByMailAndHomepageAndHostname Get CommentFieldDatas via MailAndHomepageAndHostname
func GetCommentFieldDatasByMailAndHomepageAndHostname(offset int, limit int, Mail_ string, Homepage_ string, Hostname_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("mail = ? and homepage = ? and hostname = ?", Mail_, Homepage_, Hostname_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByMailAndHomepageAndCreated Get CommentFieldDatas via MailAndHomepageAndCreated
func GetCommentFieldDatasByMailAndHomepageAndCreated(offset int, limit int, Mail_ string, Homepage_ string, Created_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("mail = ? and homepage = ? and created = ?", Mail_, Homepage_, Created_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByMailAndHomepageAndChanged Get CommentFieldDatas via MailAndHomepageAndChanged
func GetCommentFieldDatasByMailAndHomepageAndChanged(offset int, limit int, Mail_ string, Homepage_ string, Changed_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("mail = ? and homepage = ? and changed = ?", Mail_, Homepage_, Changed_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByMailAndHomepageAndStatus Get CommentFieldDatas via MailAndHomepageAndStatus
func GetCommentFieldDatasByMailAndHomepageAndStatus(offset int, limit int, Mail_ string, Homepage_ string, Status_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("mail = ? and homepage = ? and status = ?", Mail_, Homepage_, Status_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByMailAndHomepageAndThread Get CommentFieldDatas via MailAndHomepageAndThread
func GetCommentFieldDatasByMailAndHomepageAndThread(offset int, limit int, Mail_ string, Homepage_ string, Thread_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("mail = ? and homepage = ? and thread = ?", Mail_, Homepage_, Thread_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByMailAndHomepageAndEntityType Get CommentFieldDatas via MailAndHomepageAndEntityType
func GetCommentFieldDatasByMailAndHomepageAndEntityType(offset int, limit int, Mail_ string, Homepage_ string, EntityType_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("mail = ? and homepage = ? and entity_type = ?", Mail_, Homepage_, EntityType_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByMailAndHomepageAndFieldName Get CommentFieldDatas via MailAndHomepageAndFieldName
func GetCommentFieldDatasByMailAndHomepageAndFieldName(offset int, limit int, Mail_ string, Homepage_ string, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("mail = ? and homepage = ? and field_name = ?", Mail_, Homepage_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByMailAndHomepageAndDefaultLangcode Get CommentFieldDatas via MailAndHomepageAndDefaultLangcode
func GetCommentFieldDatasByMailAndHomepageAndDefaultLangcode(offset int, limit int, Mail_ string, Homepage_ string, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("mail = ? and homepage = ? and default_langcode = ?", Mail_, Homepage_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByMailAndHostnameAndCreated Get CommentFieldDatas via MailAndHostnameAndCreated
func GetCommentFieldDatasByMailAndHostnameAndCreated(offset int, limit int, Mail_ string, Hostname_ string, Created_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("mail = ? and hostname = ? and created = ?", Mail_, Hostname_, Created_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByMailAndHostnameAndChanged Get CommentFieldDatas via MailAndHostnameAndChanged
func GetCommentFieldDatasByMailAndHostnameAndChanged(offset int, limit int, Mail_ string, Hostname_ string, Changed_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("mail = ? and hostname = ? and changed = ?", Mail_, Hostname_, Changed_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByMailAndHostnameAndStatus Get CommentFieldDatas via MailAndHostnameAndStatus
func GetCommentFieldDatasByMailAndHostnameAndStatus(offset int, limit int, Mail_ string, Hostname_ string, Status_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("mail = ? and hostname = ? and status = ?", Mail_, Hostname_, Status_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByMailAndHostnameAndThread Get CommentFieldDatas via MailAndHostnameAndThread
func GetCommentFieldDatasByMailAndHostnameAndThread(offset int, limit int, Mail_ string, Hostname_ string, Thread_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("mail = ? and hostname = ? and thread = ?", Mail_, Hostname_, Thread_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByMailAndHostnameAndEntityType Get CommentFieldDatas via MailAndHostnameAndEntityType
func GetCommentFieldDatasByMailAndHostnameAndEntityType(offset int, limit int, Mail_ string, Hostname_ string, EntityType_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("mail = ? and hostname = ? and entity_type = ?", Mail_, Hostname_, EntityType_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByMailAndHostnameAndFieldName Get CommentFieldDatas via MailAndHostnameAndFieldName
func GetCommentFieldDatasByMailAndHostnameAndFieldName(offset int, limit int, Mail_ string, Hostname_ string, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("mail = ? and hostname = ? and field_name = ?", Mail_, Hostname_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByMailAndHostnameAndDefaultLangcode Get CommentFieldDatas via MailAndHostnameAndDefaultLangcode
func GetCommentFieldDatasByMailAndHostnameAndDefaultLangcode(offset int, limit int, Mail_ string, Hostname_ string, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("mail = ? and hostname = ? and default_langcode = ?", Mail_, Hostname_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByMailAndCreatedAndChanged Get CommentFieldDatas via MailAndCreatedAndChanged
func GetCommentFieldDatasByMailAndCreatedAndChanged(offset int, limit int, Mail_ string, Created_ int, Changed_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("mail = ? and created = ? and changed = ?", Mail_, Created_, Changed_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByMailAndCreatedAndStatus Get CommentFieldDatas via MailAndCreatedAndStatus
func GetCommentFieldDatasByMailAndCreatedAndStatus(offset int, limit int, Mail_ string, Created_ int, Status_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("mail = ? and created = ? and status = ?", Mail_, Created_, Status_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByMailAndCreatedAndThread Get CommentFieldDatas via MailAndCreatedAndThread
func GetCommentFieldDatasByMailAndCreatedAndThread(offset int, limit int, Mail_ string, Created_ int, Thread_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("mail = ? and created = ? and thread = ?", Mail_, Created_, Thread_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByMailAndCreatedAndEntityType Get CommentFieldDatas via MailAndCreatedAndEntityType
func GetCommentFieldDatasByMailAndCreatedAndEntityType(offset int, limit int, Mail_ string, Created_ int, EntityType_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("mail = ? and created = ? and entity_type = ?", Mail_, Created_, EntityType_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByMailAndCreatedAndFieldName Get CommentFieldDatas via MailAndCreatedAndFieldName
func GetCommentFieldDatasByMailAndCreatedAndFieldName(offset int, limit int, Mail_ string, Created_ int, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("mail = ? and created = ? and field_name = ?", Mail_, Created_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByMailAndCreatedAndDefaultLangcode Get CommentFieldDatas via MailAndCreatedAndDefaultLangcode
func GetCommentFieldDatasByMailAndCreatedAndDefaultLangcode(offset int, limit int, Mail_ string, Created_ int, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("mail = ? and created = ? and default_langcode = ?", Mail_, Created_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByMailAndChangedAndStatus Get CommentFieldDatas via MailAndChangedAndStatus
func GetCommentFieldDatasByMailAndChangedAndStatus(offset int, limit int, Mail_ string, Changed_ int, Status_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("mail = ? and changed = ? and status = ?", Mail_, Changed_, Status_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByMailAndChangedAndThread Get CommentFieldDatas via MailAndChangedAndThread
func GetCommentFieldDatasByMailAndChangedAndThread(offset int, limit int, Mail_ string, Changed_ int, Thread_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("mail = ? and changed = ? and thread = ?", Mail_, Changed_, Thread_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByMailAndChangedAndEntityType Get CommentFieldDatas via MailAndChangedAndEntityType
func GetCommentFieldDatasByMailAndChangedAndEntityType(offset int, limit int, Mail_ string, Changed_ int, EntityType_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("mail = ? and changed = ? and entity_type = ?", Mail_, Changed_, EntityType_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByMailAndChangedAndFieldName Get CommentFieldDatas via MailAndChangedAndFieldName
func GetCommentFieldDatasByMailAndChangedAndFieldName(offset int, limit int, Mail_ string, Changed_ int, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("mail = ? and changed = ? and field_name = ?", Mail_, Changed_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByMailAndChangedAndDefaultLangcode Get CommentFieldDatas via MailAndChangedAndDefaultLangcode
func GetCommentFieldDatasByMailAndChangedAndDefaultLangcode(offset int, limit int, Mail_ string, Changed_ int, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("mail = ? and changed = ? and default_langcode = ?", Mail_, Changed_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByMailAndStatusAndThread Get CommentFieldDatas via MailAndStatusAndThread
func GetCommentFieldDatasByMailAndStatusAndThread(offset int, limit int, Mail_ string, Status_ int, Thread_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("mail = ? and status = ? and thread = ?", Mail_, Status_, Thread_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByMailAndStatusAndEntityType Get CommentFieldDatas via MailAndStatusAndEntityType
func GetCommentFieldDatasByMailAndStatusAndEntityType(offset int, limit int, Mail_ string, Status_ int, EntityType_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("mail = ? and status = ? and entity_type = ?", Mail_, Status_, EntityType_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByMailAndStatusAndFieldName Get CommentFieldDatas via MailAndStatusAndFieldName
func GetCommentFieldDatasByMailAndStatusAndFieldName(offset int, limit int, Mail_ string, Status_ int, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("mail = ? and status = ? and field_name = ?", Mail_, Status_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByMailAndStatusAndDefaultLangcode Get CommentFieldDatas via MailAndStatusAndDefaultLangcode
func GetCommentFieldDatasByMailAndStatusAndDefaultLangcode(offset int, limit int, Mail_ string, Status_ int, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("mail = ? and status = ? and default_langcode = ?", Mail_, Status_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByMailAndThreadAndEntityType Get CommentFieldDatas via MailAndThreadAndEntityType
func GetCommentFieldDatasByMailAndThreadAndEntityType(offset int, limit int, Mail_ string, Thread_ string, EntityType_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("mail = ? and thread = ? and entity_type = ?", Mail_, Thread_, EntityType_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByMailAndThreadAndFieldName Get CommentFieldDatas via MailAndThreadAndFieldName
func GetCommentFieldDatasByMailAndThreadAndFieldName(offset int, limit int, Mail_ string, Thread_ string, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("mail = ? and thread = ? and field_name = ?", Mail_, Thread_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByMailAndThreadAndDefaultLangcode Get CommentFieldDatas via MailAndThreadAndDefaultLangcode
func GetCommentFieldDatasByMailAndThreadAndDefaultLangcode(offset int, limit int, Mail_ string, Thread_ string, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("mail = ? and thread = ? and default_langcode = ?", Mail_, Thread_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByMailAndEntityTypeAndFieldName Get CommentFieldDatas via MailAndEntityTypeAndFieldName
func GetCommentFieldDatasByMailAndEntityTypeAndFieldName(offset int, limit int, Mail_ string, EntityType_ string, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("mail = ? and entity_type = ? and field_name = ?", Mail_, EntityType_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByMailAndEntityTypeAndDefaultLangcode Get CommentFieldDatas via MailAndEntityTypeAndDefaultLangcode
func GetCommentFieldDatasByMailAndEntityTypeAndDefaultLangcode(offset int, limit int, Mail_ string, EntityType_ string, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("mail = ? and entity_type = ? and default_langcode = ?", Mail_, EntityType_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByMailAndFieldNameAndDefaultLangcode Get CommentFieldDatas via MailAndFieldNameAndDefaultLangcode
func GetCommentFieldDatasByMailAndFieldNameAndDefaultLangcode(offset int, limit int, Mail_ string, FieldName_ string, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("mail = ? and field_name = ? and default_langcode = ?", Mail_, FieldName_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByHomepageAndHostnameAndCreated Get CommentFieldDatas via HomepageAndHostnameAndCreated
func GetCommentFieldDatasByHomepageAndHostnameAndCreated(offset int, limit int, Homepage_ string, Hostname_ string, Created_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("homepage = ? and hostname = ? and created = ?", Homepage_, Hostname_, Created_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByHomepageAndHostnameAndChanged Get CommentFieldDatas via HomepageAndHostnameAndChanged
func GetCommentFieldDatasByHomepageAndHostnameAndChanged(offset int, limit int, Homepage_ string, Hostname_ string, Changed_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("homepage = ? and hostname = ? and changed = ?", Homepage_, Hostname_, Changed_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByHomepageAndHostnameAndStatus Get CommentFieldDatas via HomepageAndHostnameAndStatus
func GetCommentFieldDatasByHomepageAndHostnameAndStatus(offset int, limit int, Homepage_ string, Hostname_ string, Status_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("homepage = ? and hostname = ? and status = ?", Homepage_, Hostname_, Status_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByHomepageAndHostnameAndThread Get CommentFieldDatas via HomepageAndHostnameAndThread
func GetCommentFieldDatasByHomepageAndHostnameAndThread(offset int, limit int, Homepage_ string, Hostname_ string, Thread_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("homepage = ? and hostname = ? and thread = ?", Homepage_, Hostname_, Thread_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByHomepageAndHostnameAndEntityType Get CommentFieldDatas via HomepageAndHostnameAndEntityType
func GetCommentFieldDatasByHomepageAndHostnameAndEntityType(offset int, limit int, Homepage_ string, Hostname_ string, EntityType_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("homepage = ? and hostname = ? and entity_type = ?", Homepage_, Hostname_, EntityType_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByHomepageAndHostnameAndFieldName Get CommentFieldDatas via HomepageAndHostnameAndFieldName
func GetCommentFieldDatasByHomepageAndHostnameAndFieldName(offset int, limit int, Homepage_ string, Hostname_ string, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("homepage = ? and hostname = ? and field_name = ?", Homepage_, Hostname_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByHomepageAndHostnameAndDefaultLangcode Get CommentFieldDatas via HomepageAndHostnameAndDefaultLangcode
func GetCommentFieldDatasByHomepageAndHostnameAndDefaultLangcode(offset int, limit int, Homepage_ string, Hostname_ string, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("homepage = ? and hostname = ? and default_langcode = ?", Homepage_, Hostname_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByHomepageAndCreatedAndChanged Get CommentFieldDatas via HomepageAndCreatedAndChanged
func GetCommentFieldDatasByHomepageAndCreatedAndChanged(offset int, limit int, Homepage_ string, Created_ int, Changed_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("homepage = ? and created = ? and changed = ?", Homepage_, Created_, Changed_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByHomepageAndCreatedAndStatus Get CommentFieldDatas via HomepageAndCreatedAndStatus
func GetCommentFieldDatasByHomepageAndCreatedAndStatus(offset int, limit int, Homepage_ string, Created_ int, Status_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("homepage = ? and created = ? and status = ?", Homepage_, Created_, Status_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByHomepageAndCreatedAndThread Get CommentFieldDatas via HomepageAndCreatedAndThread
func GetCommentFieldDatasByHomepageAndCreatedAndThread(offset int, limit int, Homepage_ string, Created_ int, Thread_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("homepage = ? and created = ? and thread = ?", Homepage_, Created_, Thread_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByHomepageAndCreatedAndEntityType Get CommentFieldDatas via HomepageAndCreatedAndEntityType
func GetCommentFieldDatasByHomepageAndCreatedAndEntityType(offset int, limit int, Homepage_ string, Created_ int, EntityType_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("homepage = ? and created = ? and entity_type = ?", Homepage_, Created_, EntityType_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByHomepageAndCreatedAndFieldName Get CommentFieldDatas via HomepageAndCreatedAndFieldName
func GetCommentFieldDatasByHomepageAndCreatedAndFieldName(offset int, limit int, Homepage_ string, Created_ int, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("homepage = ? and created = ? and field_name = ?", Homepage_, Created_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByHomepageAndCreatedAndDefaultLangcode Get CommentFieldDatas via HomepageAndCreatedAndDefaultLangcode
func GetCommentFieldDatasByHomepageAndCreatedAndDefaultLangcode(offset int, limit int, Homepage_ string, Created_ int, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("homepage = ? and created = ? and default_langcode = ?", Homepage_, Created_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByHomepageAndChangedAndStatus Get CommentFieldDatas via HomepageAndChangedAndStatus
func GetCommentFieldDatasByHomepageAndChangedAndStatus(offset int, limit int, Homepage_ string, Changed_ int, Status_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("homepage = ? and changed = ? and status = ?", Homepage_, Changed_, Status_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByHomepageAndChangedAndThread Get CommentFieldDatas via HomepageAndChangedAndThread
func GetCommentFieldDatasByHomepageAndChangedAndThread(offset int, limit int, Homepage_ string, Changed_ int, Thread_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("homepage = ? and changed = ? and thread = ?", Homepage_, Changed_, Thread_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByHomepageAndChangedAndEntityType Get CommentFieldDatas via HomepageAndChangedAndEntityType
func GetCommentFieldDatasByHomepageAndChangedAndEntityType(offset int, limit int, Homepage_ string, Changed_ int, EntityType_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("homepage = ? and changed = ? and entity_type = ?", Homepage_, Changed_, EntityType_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByHomepageAndChangedAndFieldName Get CommentFieldDatas via HomepageAndChangedAndFieldName
func GetCommentFieldDatasByHomepageAndChangedAndFieldName(offset int, limit int, Homepage_ string, Changed_ int, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("homepage = ? and changed = ? and field_name = ?", Homepage_, Changed_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByHomepageAndChangedAndDefaultLangcode Get CommentFieldDatas via HomepageAndChangedAndDefaultLangcode
func GetCommentFieldDatasByHomepageAndChangedAndDefaultLangcode(offset int, limit int, Homepage_ string, Changed_ int, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("homepage = ? and changed = ? and default_langcode = ?", Homepage_, Changed_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByHomepageAndStatusAndThread Get CommentFieldDatas via HomepageAndStatusAndThread
func GetCommentFieldDatasByHomepageAndStatusAndThread(offset int, limit int, Homepage_ string, Status_ int, Thread_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("homepage = ? and status = ? and thread = ?", Homepage_, Status_, Thread_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByHomepageAndStatusAndEntityType Get CommentFieldDatas via HomepageAndStatusAndEntityType
func GetCommentFieldDatasByHomepageAndStatusAndEntityType(offset int, limit int, Homepage_ string, Status_ int, EntityType_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("homepage = ? and status = ? and entity_type = ?", Homepage_, Status_, EntityType_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByHomepageAndStatusAndFieldName Get CommentFieldDatas via HomepageAndStatusAndFieldName
func GetCommentFieldDatasByHomepageAndStatusAndFieldName(offset int, limit int, Homepage_ string, Status_ int, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("homepage = ? and status = ? and field_name = ?", Homepage_, Status_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByHomepageAndStatusAndDefaultLangcode Get CommentFieldDatas via HomepageAndStatusAndDefaultLangcode
func GetCommentFieldDatasByHomepageAndStatusAndDefaultLangcode(offset int, limit int, Homepage_ string, Status_ int, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("homepage = ? and status = ? and default_langcode = ?", Homepage_, Status_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByHomepageAndThreadAndEntityType Get CommentFieldDatas via HomepageAndThreadAndEntityType
func GetCommentFieldDatasByHomepageAndThreadAndEntityType(offset int, limit int, Homepage_ string, Thread_ string, EntityType_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("homepage = ? and thread = ? and entity_type = ?", Homepage_, Thread_, EntityType_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByHomepageAndThreadAndFieldName Get CommentFieldDatas via HomepageAndThreadAndFieldName
func GetCommentFieldDatasByHomepageAndThreadAndFieldName(offset int, limit int, Homepage_ string, Thread_ string, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("homepage = ? and thread = ? and field_name = ?", Homepage_, Thread_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByHomepageAndThreadAndDefaultLangcode Get CommentFieldDatas via HomepageAndThreadAndDefaultLangcode
func GetCommentFieldDatasByHomepageAndThreadAndDefaultLangcode(offset int, limit int, Homepage_ string, Thread_ string, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("homepage = ? and thread = ? and default_langcode = ?", Homepage_, Thread_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByHomepageAndEntityTypeAndFieldName Get CommentFieldDatas via HomepageAndEntityTypeAndFieldName
func GetCommentFieldDatasByHomepageAndEntityTypeAndFieldName(offset int, limit int, Homepage_ string, EntityType_ string, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("homepage = ? and entity_type = ? and field_name = ?", Homepage_, EntityType_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByHomepageAndEntityTypeAndDefaultLangcode Get CommentFieldDatas via HomepageAndEntityTypeAndDefaultLangcode
func GetCommentFieldDatasByHomepageAndEntityTypeAndDefaultLangcode(offset int, limit int, Homepage_ string, EntityType_ string, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("homepage = ? and entity_type = ? and default_langcode = ?", Homepage_, EntityType_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByHomepageAndFieldNameAndDefaultLangcode Get CommentFieldDatas via HomepageAndFieldNameAndDefaultLangcode
func GetCommentFieldDatasByHomepageAndFieldNameAndDefaultLangcode(offset int, limit int, Homepage_ string, FieldName_ string, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("homepage = ? and field_name = ? and default_langcode = ?", Homepage_, FieldName_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByHostnameAndCreatedAndChanged Get CommentFieldDatas via HostnameAndCreatedAndChanged
func GetCommentFieldDatasByHostnameAndCreatedAndChanged(offset int, limit int, Hostname_ string, Created_ int, Changed_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("hostname = ? and created = ? and changed = ?", Hostname_, Created_, Changed_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByHostnameAndCreatedAndStatus Get CommentFieldDatas via HostnameAndCreatedAndStatus
func GetCommentFieldDatasByHostnameAndCreatedAndStatus(offset int, limit int, Hostname_ string, Created_ int, Status_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("hostname = ? and created = ? and status = ?", Hostname_, Created_, Status_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByHostnameAndCreatedAndThread Get CommentFieldDatas via HostnameAndCreatedAndThread
func GetCommentFieldDatasByHostnameAndCreatedAndThread(offset int, limit int, Hostname_ string, Created_ int, Thread_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("hostname = ? and created = ? and thread = ?", Hostname_, Created_, Thread_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByHostnameAndCreatedAndEntityType Get CommentFieldDatas via HostnameAndCreatedAndEntityType
func GetCommentFieldDatasByHostnameAndCreatedAndEntityType(offset int, limit int, Hostname_ string, Created_ int, EntityType_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("hostname = ? and created = ? and entity_type = ?", Hostname_, Created_, EntityType_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByHostnameAndCreatedAndFieldName Get CommentFieldDatas via HostnameAndCreatedAndFieldName
func GetCommentFieldDatasByHostnameAndCreatedAndFieldName(offset int, limit int, Hostname_ string, Created_ int, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("hostname = ? and created = ? and field_name = ?", Hostname_, Created_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByHostnameAndCreatedAndDefaultLangcode Get CommentFieldDatas via HostnameAndCreatedAndDefaultLangcode
func GetCommentFieldDatasByHostnameAndCreatedAndDefaultLangcode(offset int, limit int, Hostname_ string, Created_ int, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("hostname = ? and created = ? and default_langcode = ?", Hostname_, Created_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByHostnameAndChangedAndStatus Get CommentFieldDatas via HostnameAndChangedAndStatus
func GetCommentFieldDatasByHostnameAndChangedAndStatus(offset int, limit int, Hostname_ string, Changed_ int, Status_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("hostname = ? and changed = ? and status = ?", Hostname_, Changed_, Status_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByHostnameAndChangedAndThread Get CommentFieldDatas via HostnameAndChangedAndThread
func GetCommentFieldDatasByHostnameAndChangedAndThread(offset int, limit int, Hostname_ string, Changed_ int, Thread_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("hostname = ? and changed = ? and thread = ?", Hostname_, Changed_, Thread_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByHostnameAndChangedAndEntityType Get CommentFieldDatas via HostnameAndChangedAndEntityType
func GetCommentFieldDatasByHostnameAndChangedAndEntityType(offset int, limit int, Hostname_ string, Changed_ int, EntityType_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("hostname = ? and changed = ? and entity_type = ?", Hostname_, Changed_, EntityType_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByHostnameAndChangedAndFieldName Get CommentFieldDatas via HostnameAndChangedAndFieldName
func GetCommentFieldDatasByHostnameAndChangedAndFieldName(offset int, limit int, Hostname_ string, Changed_ int, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("hostname = ? and changed = ? and field_name = ?", Hostname_, Changed_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByHostnameAndChangedAndDefaultLangcode Get CommentFieldDatas via HostnameAndChangedAndDefaultLangcode
func GetCommentFieldDatasByHostnameAndChangedAndDefaultLangcode(offset int, limit int, Hostname_ string, Changed_ int, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("hostname = ? and changed = ? and default_langcode = ?", Hostname_, Changed_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByHostnameAndStatusAndThread Get CommentFieldDatas via HostnameAndStatusAndThread
func GetCommentFieldDatasByHostnameAndStatusAndThread(offset int, limit int, Hostname_ string, Status_ int, Thread_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("hostname = ? and status = ? and thread = ?", Hostname_, Status_, Thread_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByHostnameAndStatusAndEntityType Get CommentFieldDatas via HostnameAndStatusAndEntityType
func GetCommentFieldDatasByHostnameAndStatusAndEntityType(offset int, limit int, Hostname_ string, Status_ int, EntityType_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("hostname = ? and status = ? and entity_type = ?", Hostname_, Status_, EntityType_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByHostnameAndStatusAndFieldName Get CommentFieldDatas via HostnameAndStatusAndFieldName
func GetCommentFieldDatasByHostnameAndStatusAndFieldName(offset int, limit int, Hostname_ string, Status_ int, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("hostname = ? and status = ? and field_name = ?", Hostname_, Status_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByHostnameAndStatusAndDefaultLangcode Get CommentFieldDatas via HostnameAndStatusAndDefaultLangcode
func GetCommentFieldDatasByHostnameAndStatusAndDefaultLangcode(offset int, limit int, Hostname_ string, Status_ int, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("hostname = ? and status = ? and default_langcode = ?", Hostname_, Status_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByHostnameAndThreadAndEntityType Get CommentFieldDatas via HostnameAndThreadAndEntityType
func GetCommentFieldDatasByHostnameAndThreadAndEntityType(offset int, limit int, Hostname_ string, Thread_ string, EntityType_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("hostname = ? and thread = ? and entity_type = ?", Hostname_, Thread_, EntityType_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByHostnameAndThreadAndFieldName Get CommentFieldDatas via HostnameAndThreadAndFieldName
func GetCommentFieldDatasByHostnameAndThreadAndFieldName(offset int, limit int, Hostname_ string, Thread_ string, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("hostname = ? and thread = ? and field_name = ?", Hostname_, Thread_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByHostnameAndThreadAndDefaultLangcode Get CommentFieldDatas via HostnameAndThreadAndDefaultLangcode
func GetCommentFieldDatasByHostnameAndThreadAndDefaultLangcode(offset int, limit int, Hostname_ string, Thread_ string, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("hostname = ? and thread = ? and default_langcode = ?", Hostname_, Thread_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByHostnameAndEntityTypeAndFieldName Get CommentFieldDatas via HostnameAndEntityTypeAndFieldName
func GetCommentFieldDatasByHostnameAndEntityTypeAndFieldName(offset int, limit int, Hostname_ string, EntityType_ string, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("hostname = ? and entity_type = ? and field_name = ?", Hostname_, EntityType_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByHostnameAndEntityTypeAndDefaultLangcode Get CommentFieldDatas via HostnameAndEntityTypeAndDefaultLangcode
func GetCommentFieldDatasByHostnameAndEntityTypeAndDefaultLangcode(offset int, limit int, Hostname_ string, EntityType_ string, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("hostname = ? and entity_type = ? and default_langcode = ?", Hostname_, EntityType_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByHostnameAndFieldNameAndDefaultLangcode Get CommentFieldDatas via HostnameAndFieldNameAndDefaultLangcode
func GetCommentFieldDatasByHostnameAndFieldNameAndDefaultLangcode(offset int, limit int, Hostname_ string, FieldName_ string, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("hostname = ? and field_name = ? and default_langcode = ?", Hostname_, FieldName_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCreatedAndChangedAndStatus Get CommentFieldDatas via CreatedAndChangedAndStatus
func GetCommentFieldDatasByCreatedAndChangedAndStatus(offset int, limit int, Created_ int, Changed_ int, Status_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("created = ? and changed = ? and status = ?", Created_, Changed_, Status_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCreatedAndChangedAndThread Get CommentFieldDatas via CreatedAndChangedAndThread
func GetCommentFieldDatasByCreatedAndChangedAndThread(offset int, limit int, Created_ int, Changed_ int, Thread_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("created = ? and changed = ? and thread = ?", Created_, Changed_, Thread_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCreatedAndChangedAndEntityType Get CommentFieldDatas via CreatedAndChangedAndEntityType
func GetCommentFieldDatasByCreatedAndChangedAndEntityType(offset int, limit int, Created_ int, Changed_ int, EntityType_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("created = ? and changed = ? and entity_type = ?", Created_, Changed_, EntityType_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCreatedAndChangedAndFieldName Get CommentFieldDatas via CreatedAndChangedAndFieldName
func GetCommentFieldDatasByCreatedAndChangedAndFieldName(offset int, limit int, Created_ int, Changed_ int, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("created = ? and changed = ? and field_name = ?", Created_, Changed_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCreatedAndChangedAndDefaultLangcode Get CommentFieldDatas via CreatedAndChangedAndDefaultLangcode
func GetCommentFieldDatasByCreatedAndChangedAndDefaultLangcode(offset int, limit int, Created_ int, Changed_ int, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("created = ? and changed = ? and default_langcode = ?", Created_, Changed_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCreatedAndStatusAndThread Get CommentFieldDatas via CreatedAndStatusAndThread
func GetCommentFieldDatasByCreatedAndStatusAndThread(offset int, limit int, Created_ int, Status_ int, Thread_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("created = ? and status = ? and thread = ?", Created_, Status_, Thread_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCreatedAndStatusAndEntityType Get CommentFieldDatas via CreatedAndStatusAndEntityType
func GetCommentFieldDatasByCreatedAndStatusAndEntityType(offset int, limit int, Created_ int, Status_ int, EntityType_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("created = ? and status = ? and entity_type = ?", Created_, Status_, EntityType_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCreatedAndStatusAndFieldName Get CommentFieldDatas via CreatedAndStatusAndFieldName
func GetCommentFieldDatasByCreatedAndStatusAndFieldName(offset int, limit int, Created_ int, Status_ int, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("created = ? and status = ? and field_name = ?", Created_, Status_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCreatedAndStatusAndDefaultLangcode Get CommentFieldDatas via CreatedAndStatusAndDefaultLangcode
func GetCommentFieldDatasByCreatedAndStatusAndDefaultLangcode(offset int, limit int, Created_ int, Status_ int, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("created = ? and status = ? and default_langcode = ?", Created_, Status_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCreatedAndThreadAndEntityType Get CommentFieldDatas via CreatedAndThreadAndEntityType
func GetCommentFieldDatasByCreatedAndThreadAndEntityType(offset int, limit int, Created_ int, Thread_ string, EntityType_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("created = ? and thread = ? and entity_type = ?", Created_, Thread_, EntityType_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCreatedAndThreadAndFieldName Get CommentFieldDatas via CreatedAndThreadAndFieldName
func GetCommentFieldDatasByCreatedAndThreadAndFieldName(offset int, limit int, Created_ int, Thread_ string, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("created = ? and thread = ? and field_name = ?", Created_, Thread_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCreatedAndThreadAndDefaultLangcode Get CommentFieldDatas via CreatedAndThreadAndDefaultLangcode
func GetCommentFieldDatasByCreatedAndThreadAndDefaultLangcode(offset int, limit int, Created_ int, Thread_ string, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("created = ? and thread = ? and default_langcode = ?", Created_, Thread_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCreatedAndEntityTypeAndFieldName Get CommentFieldDatas via CreatedAndEntityTypeAndFieldName
func GetCommentFieldDatasByCreatedAndEntityTypeAndFieldName(offset int, limit int, Created_ int, EntityType_ string, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("created = ? and entity_type = ? and field_name = ?", Created_, EntityType_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCreatedAndEntityTypeAndDefaultLangcode Get CommentFieldDatas via CreatedAndEntityTypeAndDefaultLangcode
func GetCommentFieldDatasByCreatedAndEntityTypeAndDefaultLangcode(offset int, limit int, Created_ int, EntityType_ string, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("created = ? and entity_type = ? and default_langcode = ?", Created_, EntityType_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCreatedAndFieldNameAndDefaultLangcode Get CommentFieldDatas via CreatedAndFieldNameAndDefaultLangcode
func GetCommentFieldDatasByCreatedAndFieldNameAndDefaultLangcode(offset int, limit int, Created_ int, FieldName_ string, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("created = ? and field_name = ? and default_langcode = ?", Created_, FieldName_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByChangedAndStatusAndThread Get CommentFieldDatas via ChangedAndStatusAndThread
func GetCommentFieldDatasByChangedAndStatusAndThread(offset int, limit int, Changed_ int, Status_ int, Thread_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("changed = ? and status = ? and thread = ?", Changed_, Status_, Thread_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByChangedAndStatusAndEntityType Get CommentFieldDatas via ChangedAndStatusAndEntityType
func GetCommentFieldDatasByChangedAndStatusAndEntityType(offset int, limit int, Changed_ int, Status_ int, EntityType_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("changed = ? and status = ? and entity_type = ?", Changed_, Status_, EntityType_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByChangedAndStatusAndFieldName Get CommentFieldDatas via ChangedAndStatusAndFieldName
func GetCommentFieldDatasByChangedAndStatusAndFieldName(offset int, limit int, Changed_ int, Status_ int, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("changed = ? and status = ? and field_name = ?", Changed_, Status_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByChangedAndStatusAndDefaultLangcode Get CommentFieldDatas via ChangedAndStatusAndDefaultLangcode
func GetCommentFieldDatasByChangedAndStatusAndDefaultLangcode(offset int, limit int, Changed_ int, Status_ int, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("changed = ? and status = ? and default_langcode = ?", Changed_, Status_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByChangedAndThreadAndEntityType Get CommentFieldDatas via ChangedAndThreadAndEntityType
func GetCommentFieldDatasByChangedAndThreadAndEntityType(offset int, limit int, Changed_ int, Thread_ string, EntityType_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("changed = ? and thread = ? and entity_type = ?", Changed_, Thread_, EntityType_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByChangedAndThreadAndFieldName Get CommentFieldDatas via ChangedAndThreadAndFieldName
func GetCommentFieldDatasByChangedAndThreadAndFieldName(offset int, limit int, Changed_ int, Thread_ string, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("changed = ? and thread = ? and field_name = ?", Changed_, Thread_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByChangedAndThreadAndDefaultLangcode Get CommentFieldDatas via ChangedAndThreadAndDefaultLangcode
func GetCommentFieldDatasByChangedAndThreadAndDefaultLangcode(offset int, limit int, Changed_ int, Thread_ string, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("changed = ? and thread = ? and default_langcode = ?", Changed_, Thread_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByChangedAndEntityTypeAndFieldName Get CommentFieldDatas via ChangedAndEntityTypeAndFieldName
func GetCommentFieldDatasByChangedAndEntityTypeAndFieldName(offset int, limit int, Changed_ int, EntityType_ string, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("changed = ? and entity_type = ? and field_name = ?", Changed_, EntityType_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByChangedAndEntityTypeAndDefaultLangcode Get CommentFieldDatas via ChangedAndEntityTypeAndDefaultLangcode
func GetCommentFieldDatasByChangedAndEntityTypeAndDefaultLangcode(offset int, limit int, Changed_ int, EntityType_ string, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("changed = ? and entity_type = ? and default_langcode = ?", Changed_, EntityType_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByChangedAndFieldNameAndDefaultLangcode Get CommentFieldDatas via ChangedAndFieldNameAndDefaultLangcode
func GetCommentFieldDatasByChangedAndFieldNameAndDefaultLangcode(offset int, limit int, Changed_ int, FieldName_ string, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("changed = ? and field_name = ? and default_langcode = ?", Changed_, FieldName_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByStatusAndThreadAndEntityType Get CommentFieldDatas via StatusAndThreadAndEntityType
func GetCommentFieldDatasByStatusAndThreadAndEntityType(offset int, limit int, Status_ int, Thread_ string, EntityType_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("status = ? and thread = ? and entity_type = ?", Status_, Thread_, EntityType_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByStatusAndThreadAndFieldName Get CommentFieldDatas via StatusAndThreadAndFieldName
func GetCommentFieldDatasByStatusAndThreadAndFieldName(offset int, limit int, Status_ int, Thread_ string, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("status = ? and thread = ? and field_name = ?", Status_, Thread_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByStatusAndThreadAndDefaultLangcode Get CommentFieldDatas via StatusAndThreadAndDefaultLangcode
func GetCommentFieldDatasByStatusAndThreadAndDefaultLangcode(offset int, limit int, Status_ int, Thread_ string, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("status = ? and thread = ? and default_langcode = ?", Status_, Thread_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByStatusAndEntityTypeAndFieldName Get CommentFieldDatas via StatusAndEntityTypeAndFieldName
func GetCommentFieldDatasByStatusAndEntityTypeAndFieldName(offset int, limit int, Status_ int, EntityType_ string, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("status = ? and entity_type = ? and field_name = ?", Status_, EntityType_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByStatusAndEntityTypeAndDefaultLangcode Get CommentFieldDatas via StatusAndEntityTypeAndDefaultLangcode
func GetCommentFieldDatasByStatusAndEntityTypeAndDefaultLangcode(offset int, limit int, Status_ int, EntityType_ string, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("status = ? and entity_type = ? and default_langcode = ?", Status_, EntityType_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByStatusAndFieldNameAndDefaultLangcode Get CommentFieldDatas via StatusAndFieldNameAndDefaultLangcode
func GetCommentFieldDatasByStatusAndFieldNameAndDefaultLangcode(offset int, limit int, Status_ int, FieldName_ string, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("status = ? and field_name = ? and default_langcode = ?", Status_, FieldName_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByThreadAndEntityTypeAndFieldName Get CommentFieldDatas via ThreadAndEntityTypeAndFieldName
func GetCommentFieldDatasByThreadAndEntityTypeAndFieldName(offset int, limit int, Thread_ string, EntityType_ string, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("thread = ? and entity_type = ? and field_name = ?", Thread_, EntityType_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByThreadAndEntityTypeAndDefaultLangcode Get CommentFieldDatas via ThreadAndEntityTypeAndDefaultLangcode
func GetCommentFieldDatasByThreadAndEntityTypeAndDefaultLangcode(offset int, limit int, Thread_ string, EntityType_ string, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("thread = ? and entity_type = ? and default_langcode = ?", Thread_, EntityType_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByThreadAndFieldNameAndDefaultLangcode Get CommentFieldDatas via ThreadAndFieldNameAndDefaultLangcode
func GetCommentFieldDatasByThreadAndFieldNameAndDefaultLangcode(offset int, limit int, Thread_ string, FieldName_ string, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("thread = ? and field_name = ? and default_langcode = ?", Thread_, FieldName_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByEntityTypeAndFieldNameAndDefaultLangcode Get CommentFieldDatas via EntityTypeAndFieldNameAndDefaultLangcode
func GetCommentFieldDatasByEntityTypeAndFieldNameAndDefaultLangcode(offset int, limit int, EntityType_ string, FieldName_ string, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("entity_type = ? and field_name = ? and default_langcode = ?", EntityType_, FieldName_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndCommentType Get CommentFieldDatas via CidAndCommentType
func GetCommentFieldDatasByCidAndCommentType(offset int, limit int, Cid_ int64, CommentType_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and comment_type = ?", Cid_, CommentType_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndLangcode Get CommentFieldDatas via CidAndLangcode
func GetCommentFieldDatasByCidAndLangcode(offset int, limit int, Cid_ int64, Langcode_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and langcode = ?", Cid_, Langcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndPid Get CommentFieldDatas via CidAndPid
func GetCommentFieldDatasByCidAndPid(offset int, limit int, Cid_ int64, Pid_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and pid = ?", Cid_, Pid_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndEntityId Get CommentFieldDatas via CidAndEntityId
func GetCommentFieldDatasByCidAndEntityId(offset int, limit int, Cid_ int64, EntityId_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and entity_id = ?", Cid_, EntityId_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndSubject Get CommentFieldDatas via CidAndSubject
func GetCommentFieldDatasByCidAndSubject(offset int, limit int, Cid_ int64, Subject_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and subject = ?", Cid_, Subject_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndUid Get CommentFieldDatas via CidAndUid
func GetCommentFieldDatasByCidAndUid(offset int, limit int, Cid_ int64, Uid_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and uid = ?", Cid_, Uid_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndName Get CommentFieldDatas via CidAndName
func GetCommentFieldDatasByCidAndName(offset int, limit int, Cid_ int64, Name_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and name = ?", Cid_, Name_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndMail Get CommentFieldDatas via CidAndMail
func GetCommentFieldDatasByCidAndMail(offset int, limit int, Cid_ int64, Mail_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and mail = ?", Cid_, Mail_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndHomepage Get CommentFieldDatas via CidAndHomepage
func GetCommentFieldDatasByCidAndHomepage(offset int, limit int, Cid_ int64, Homepage_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and homepage = ?", Cid_, Homepage_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndHostname Get CommentFieldDatas via CidAndHostname
func GetCommentFieldDatasByCidAndHostname(offset int, limit int, Cid_ int64, Hostname_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and hostname = ?", Cid_, Hostname_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndCreated Get CommentFieldDatas via CidAndCreated
func GetCommentFieldDatasByCidAndCreated(offset int, limit int, Cid_ int64, Created_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and created = ?", Cid_, Created_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndChanged Get CommentFieldDatas via CidAndChanged
func GetCommentFieldDatasByCidAndChanged(offset int, limit int, Cid_ int64, Changed_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and changed = ?", Cid_, Changed_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndStatus Get CommentFieldDatas via CidAndStatus
func GetCommentFieldDatasByCidAndStatus(offset int, limit int, Cid_ int64, Status_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and status = ?", Cid_, Status_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndThread Get CommentFieldDatas via CidAndThread
func GetCommentFieldDatasByCidAndThread(offset int, limit int, Cid_ int64, Thread_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and thread = ?", Cid_, Thread_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndEntityType Get CommentFieldDatas via CidAndEntityType
func GetCommentFieldDatasByCidAndEntityType(offset int, limit int, Cid_ int64, EntityType_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and entity_type = ?", Cid_, EntityType_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndFieldName Get CommentFieldDatas via CidAndFieldName
func GetCommentFieldDatasByCidAndFieldName(offset int, limit int, Cid_ int64, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and field_name = ?", Cid_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCidAndDefaultLangcode Get CommentFieldDatas via CidAndDefaultLangcode
func GetCommentFieldDatasByCidAndDefaultLangcode(offset int, limit int, Cid_ int64, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ? and default_langcode = ?", Cid_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndLangcode Get CommentFieldDatas via CommentTypeAndLangcode
func GetCommentFieldDatasByCommentTypeAndLangcode(offset int, limit int, CommentType_ string, Langcode_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and langcode = ?", CommentType_, Langcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndPid Get CommentFieldDatas via CommentTypeAndPid
func GetCommentFieldDatasByCommentTypeAndPid(offset int, limit int, CommentType_ string, Pid_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and pid = ?", CommentType_, Pid_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndEntityId Get CommentFieldDatas via CommentTypeAndEntityId
func GetCommentFieldDatasByCommentTypeAndEntityId(offset int, limit int, CommentType_ string, EntityId_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and entity_id = ?", CommentType_, EntityId_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndSubject Get CommentFieldDatas via CommentTypeAndSubject
func GetCommentFieldDatasByCommentTypeAndSubject(offset int, limit int, CommentType_ string, Subject_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and subject = ?", CommentType_, Subject_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndUid Get CommentFieldDatas via CommentTypeAndUid
func GetCommentFieldDatasByCommentTypeAndUid(offset int, limit int, CommentType_ string, Uid_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and uid = ?", CommentType_, Uid_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndName Get CommentFieldDatas via CommentTypeAndName
func GetCommentFieldDatasByCommentTypeAndName(offset int, limit int, CommentType_ string, Name_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and name = ?", CommentType_, Name_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndMail Get CommentFieldDatas via CommentTypeAndMail
func GetCommentFieldDatasByCommentTypeAndMail(offset int, limit int, CommentType_ string, Mail_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and mail = ?", CommentType_, Mail_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndHomepage Get CommentFieldDatas via CommentTypeAndHomepage
func GetCommentFieldDatasByCommentTypeAndHomepage(offset int, limit int, CommentType_ string, Homepage_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and homepage = ?", CommentType_, Homepage_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndHostname Get CommentFieldDatas via CommentTypeAndHostname
func GetCommentFieldDatasByCommentTypeAndHostname(offset int, limit int, CommentType_ string, Hostname_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and hostname = ?", CommentType_, Hostname_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndCreated Get CommentFieldDatas via CommentTypeAndCreated
func GetCommentFieldDatasByCommentTypeAndCreated(offset int, limit int, CommentType_ string, Created_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and created = ?", CommentType_, Created_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndChanged Get CommentFieldDatas via CommentTypeAndChanged
func GetCommentFieldDatasByCommentTypeAndChanged(offset int, limit int, CommentType_ string, Changed_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and changed = ?", CommentType_, Changed_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndStatus Get CommentFieldDatas via CommentTypeAndStatus
func GetCommentFieldDatasByCommentTypeAndStatus(offset int, limit int, CommentType_ string, Status_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and status = ?", CommentType_, Status_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndThread Get CommentFieldDatas via CommentTypeAndThread
func GetCommentFieldDatasByCommentTypeAndThread(offset int, limit int, CommentType_ string, Thread_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and thread = ?", CommentType_, Thread_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndEntityType Get CommentFieldDatas via CommentTypeAndEntityType
func GetCommentFieldDatasByCommentTypeAndEntityType(offset int, limit int, CommentType_ string, EntityType_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and entity_type = ?", CommentType_, EntityType_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndFieldName Get CommentFieldDatas via CommentTypeAndFieldName
func GetCommentFieldDatasByCommentTypeAndFieldName(offset int, limit int, CommentType_ string, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and field_name = ?", CommentType_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCommentTypeAndDefaultLangcode Get CommentFieldDatas via CommentTypeAndDefaultLangcode
func GetCommentFieldDatasByCommentTypeAndDefaultLangcode(offset int, limit int, CommentType_ string, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ? and default_langcode = ?", CommentType_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByLangcodeAndPid Get CommentFieldDatas via LangcodeAndPid
func GetCommentFieldDatasByLangcodeAndPid(offset int, limit int, Langcode_ string, Pid_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("langcode = ? and pid = ?", Langcode_, Pid_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByLangcodeAndEntityId Get CommentFieldDatas via LangcodeAndEntityId
func GetCommentFieldDatasByLangcodeAndEntityId(offset int, limit int, Langcode_ string, EntityId_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("langcode = ? and entity_id = ?", Langcode_, EntityId_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByLangcodeAndSubject Get CommentFieldDatas via LangcodeAndSubject
func GetCommentFieldDatasByLangcodeAndSubject(offset int, limit int, Langcode_ string, Subject_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("langcode = ? and subject = ?", Langcode_, Subject_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByLangcodeAndUid Get CommentFieldDatas via LangcodeAndUid
func GetCommentFieldDatasByLangcodeAndUid(offset int, limit int, Langcode_ string, Uid_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("langcode = ? and uid = ?", Langcode_, Uid_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByLangcodeAndName Get CommentFieldDatas via LangcodeAndName
func GetCommentFieldDatasByLangcodeAndName(offset int, limit int, Langcode_ string, Name_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("langcode = ? and name = ?", Langcode_, Name_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByLangcodeAndMail Get CommentFieldDatas via LangcodeAndMail
func GetCommentFieldDatasByLangcodeAndMail(offset int, limit int, Langcode_ string, Mail_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("langcode = ? and mail = ?", Langcode_, Mail_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByLangcodeAndHomepage Get CommentFieldDatas via LangcodeAndHomepage
func GetCommentFieldDatasByLangcodeAndHomepage(offset int, limit int, Langcode_ string, Homepage_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("langcode = ? and homepage = ?", Langcode_, Homepage_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByLangcodeAndHostname Get CommentFieldDatas via LangcodeAndHostname
func GetCommentFieldDatasByLangcodeAndHostname(offset int, limit int, Langcode_ string, Hostname_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("langcode = ? and hostname = ?", Langcode_, Hostname_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByLangcodeAndCreated Get CommentFieldDatas via LangcodeAndCreated
func GetCommentFieldDatasByLangcodeAndCreated(offset int, limit int, Langcode_ string, Created_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("langcode = ? and created = ?", Langcode_, Created_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByLangcodeAndChanged Get CommentFieldDatas via LangcodeAndChanged
func GetCommentFieldDatasByLangcodeAndChanged(offset int, limit int, Langcode_ string, Changed_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("langcode = ? and changed = ?", Langcode_, Changed_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByLangcodeAndStatus Get CommentFieldDatas via LangcodeAndStatus
func GetCommentFieldDatasByLangcodeAndStatus(offset int, limit int, Langcode_ string, Status_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("langcode = ? and status = ?", Langcode_, Status_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByLangcodeAndThread Get CommentFieldDatas via LangcodeAndThread
func GetCommentFieldDatasByLangcodeAndThread(offset int, limit int, Langcode_ string, Thread_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("langcode = ? and thread = ?", Langcode_, Thread_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByLangcodeAndEntityType Get CommentFieldDatas via LangcodeAndEntityType
func GetCommentFieldDatasByLangcodeAndEntityType(offset int, limit int, Langcode_ string, EntityType_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("langcode = ? and entity_type = ?", Langcode_, EntityType_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByLangcodeAndFieldName Get CommentFieldDatas via LangcodeAndFieldName
func GetCommentFieldDatasByLangcodeAndFieldName(offset int, limit int, Langcode_ string, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("langcode = ? and field_name = ?", Langcode_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByLangcodeAndDefaultLangcode Get CommentFieldDatas via LangcodeAndDefaultLangcode
func GetCommentFieldDatasByLangcodeAndDefaultLangcode(offset int, limit int, Langcode_ string, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("langcode = ? and default_langcode = ?", Langcode_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByPidAndEntityId Get CommentFieldDatas via PidAndEntityId
func GetCommentFieldDatasByPidAndEntityId(offset int, limit int, Pid_ int, EntityId_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("pid = ? and entity_id = ?", Pid_, EntityId_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByPidAndSubject Get CommentFieldDatas via PidAndSubject
func GetCommentFieldDatasByPidAndSubject(offset int, limit int, Pid_ int, Subject_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("pid = ? and subject = ?", Pid_, Subject_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByPidAndUid Get CommentFieldDatas via PidAndUid
func GetCommentFieldDatasByPidAndUid(offset int, limit int, Pid_ int, Uid_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("pid = ? and uid = ?", Pid_, Uid_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByPidAndName Get CommentFieldDatas via PidAndName
func GetCommentFieldDatasByPidAndName(offset int, limit int, Pid_ int, Name_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("pid = ? and name = ?", Pid_, Name_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByPidAndMail Get CommentFieldDatas via PidAndMail
func GetCommentFieldDatasByPidAndMail(offset int, limit int, Pid_ int, Mail_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("pid = ? and mail = ?", Pid_, Mail_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByPidAndHomepage Get CommentFieldDatas via PidAndHomepage
func GetCommentFieldDatasByPidAndHomepage(offset int, limit int, Pid_ int, Homepage_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("pid = ? and homepage = ?", Pid_, Homepage_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByPidAndHostname Get CommentFieldDatas via PidAndHostname
func GetCommentFieldDatasByPidAndHostname(offset int, limit int, Pid_ int, Hostname_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("pid = ? and hostname = ?", Pid_, Hostname_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByPidAndCreated Get CommentFieldDatas via PidAndCreated
func GetCommentFieldDatasByPidAndCreated(offset int, limit int, Pid_ int, Created_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("pid = ? and created = ?", Pid_, Created_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByPidAndChanged Get CommentFieldDatas via PidAndChanged
func GetCommentFieldDatasByPidAndChanged(offset int, limit int, Pid_ int, Changed_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("pid = ? and changed = ?", Pid_, Changed_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByPidAndStatus Get CommentFieldDatas via PidAndStatus
func GetCommentFieldDatasByPidAndStatus(offset int, limit int, Pid_ int, Status_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("pid = ? and status = ?", Pid_, Status_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByPidAndThread Get CommentFieldDatas via PidAndThread
func GetCommentFieldDatasByPidAndThread(offset int, limit int, Pid_ int, Thread_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("pid = ? and thread = ?", Pid_, Thread_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByPidAndEntityType Get CommentFieldDatas via PidAndEntityType
func GetCommentFieldDatasByPidAndEntityType(offset int, limit int, Pid_ int, EntityType_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("pid = ? and entity_type = ?", Pid_, EntityType_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByPidAndFieldName Get CommentFieldDatas via PidAndFieldName
func GetCommentFieldDatasByPidAndFieldName(offset int, limit int, Pid_ int, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("pid = ? and field_name = ?", Pid_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByPidAndDefaultLangcode Get CommentFieldDatas via PidAndDefaultLangcode
func GetCommentFieldDatasByPidAndDefaultLangcode(offset int, limit int, Pid_ int, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("pid = ? and default_langcode = ?", Pid_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByEntityIdAndSubject Get CommentFieldDatas via EntityIdAndSubject
func GetCommentFieldDatasByEntityIdAndSubject(offset int, limit int, EntityId_ int, Subject_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("entity_id = ? and subject = ?", EntityId_, Subject_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByEntityIdAndUid Get CommentFieldDatas via EntityIdAndUid
func GetCommentFieldDatasByEntityIdAndUid(offset int, limit int, EntityId_ int, Uid_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("entity_id = ? and uid = ?", EntityId_, Uid_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByEntityIdAndName Get CommentFieldDatas via EntityIdAndName
func GetCommentFieldDatasByEntityIdAndName(offset int, limit int, EntityId_ int, Name_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("entity_id = ? and name = ?", EntityId_, Name_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByEntityIdAndMail Get CommentFieldDatas via EntityIdAndMail
func GetCommentFieldDatasByEntityIdAndMail(offset int, limit int, EntityId_ int, Mail_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("entity_id = ? and mail = ?", EntityId_, Mail_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByEntityIdAndHomepage Get CommentFieldDatas via EntityIdAndHomepage
func GetCommentFieldDatasByEntityIdAndHomepage(offset int, limit int, EntityId_ int, Homepage_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("entity_id = ? and homepage = ?", EntityId_, Homepage_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByEntityIdAndHostname Get CommentFieldDatas via EntityIdAndHostname
func GetCommentFieldDatasByEntityIdAndHostname(offset int, limit int, EntityId_ int, Hostname_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("entity_id = ? and hostname = ?", EntityId_, Hostname_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByEntityIdAndCreated Get CommentFieldDatas via EntityIdAndCreated
func GetCommentFieldDatasByEntityIdAndCreated(offset int, limit int, EntityId_ int, Created_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("entity_id = ? and created = ?", EntityId_, Created_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByEntityIdAndChanged Get CommentFieldDatas via EntityIdAndChanged
func GetCommentFieldDatasByEntityIdAndChanged(offset int, limit int, EntityId_ int, Changed_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("entity_id = ? and changed = ?", EntityId_, Changed_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByEntityIdAndStatus Get CommentFieldDatas via EntityIdAndStatus
func GetCommentFieldDatasByEntityIdAndStatus(offset int, limit int, EntityId_ int, Status_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("entity_id = ? and status = ?", EntityId_, Status_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByEntityIdAndThread Get CommentFieldDatas via EntityIdAndThread
func GetCommentFieldDatasByEntityIdAndThread(offset int, limit int, EntityId_ int, Thread_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("entity_id = ? and thread = ?", EntityId_, Thread_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByEntityIdAndEntityType Get CommentFieldDatas via EntityIdAndEntityType
func GetCommentFieldDatasByEntityIdAndEntityType(offset int, limit int, EntityId_ int, EntityType_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("entity_id = ? and entity_type = ?", EntityId_, EntityType_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByEntityIdAndFieldName Get CommentFieldDatas via EntityIdAndFieldName
func GetCommentFieldDatasByEntityIdAndFieldName(offset int, limit int, EntityId_ int, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("entity_id = ? and field_name = ?", EntityId_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByEntityIdAndDefaultLangcode Get CommentFieldDatas via EntityIdAndDefaultLangcode
func GetCommentFieldDatasByEntityIdAndDefaultLangcode(offset int, limit int, EntityId_ int, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("entity_id = ? and default_langcode = ?", EntityId_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasBySubjectAndUid Get CommentFieldDatas via SubjectAndUid
func GetCommentFieldDatasBySubjectAndUid(offset int, limit int, Subject_ string, Uid_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("subject = ? and uid = ?", Subject_, Uid_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasBySubjectAndName Get CommentFieldDatas via SubjectAndName
func GetCommentFieldDatasBySubjectAndName(offset int, limit int, Subject_ string, Name_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("subject = ? and name = ?", Subject_, Name_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasBySubjectAndMail Get CommentFieldDatas via SubjectAndMail
func GetCommentFieldDatasBySubjectAndMail(offset int, limit int, Subject_ string, Mail_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("subject = ? and mail = ?", Subject_, Mail_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasBySubjectAndHomepage Get CommentFieldDatas via SubjectAndHomepage
func GetCommentFieldDatasBySubjectAndHomepage(offset int, limit int, Subject_ string, Homepage_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("subject = ? and homepage = ?", Subject_, Homepage_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasBySubjectAndHostname Get CommentFieldDatas via SubjectAndHostname
func GetCommentFieldDatasBySubjectAndHostname(offset int, limit int, Subject_ string, Hostname_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("subject = ? and hostname = ?", Subject_, Hostname_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasBySubjectAndCreated Get CommentFieldDatas via SubjectAndCreated
func GetCommentFieldDatasBySubjectAndCreated(offset int, limit int, Subject_ string, Created_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("subject = ? and created = ?", Subject_, Created_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasBySubjectAndChanged Get CommentFieldDatas via SubjectAndChanged
func GetCommentFieldDatasBySubjectAndChanged(offset int, limit int, Subject_ string, Changed_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("subject = ? and changed = ?", Subject_, Changed_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasBySubjectAndStatus Get CommentFieldDatas via SubjectAndStatus
func GetCommentFieldDatasBySubjectAndStatus(offset int, limit int, Subject_ string, Status_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("subject = ? and status = ?", Subject_, Status_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasBySubjectAndThread Get CommentFieldDatas via SubjectAndThread
func GetCommentFieldDatasBySubjectAndThread(offset int, limit int, Subject_ string, Thread_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("subject = ? and thread = ?", Subject_, Thread_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasBySubjectAndEntityType Get CommentFieldDatas via SubjectAndEntityType
func GetCommentFieldDatasBySubjectAndEntityType(offset int, limit int, Subject_ string, EntityType_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("subject = ? and entity_type = ?", Subject_, EntityType_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasBySubjectAndFieldName Get CommentFieldDatas via SubjectAndFieldName
func GetCommentFieldDatasBySubjectAndFieldName(offset int, limit int, Subject_ string, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("subject = ? and field_name = ?", Subject_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasBySubjectAndDefaultLangcode Get CommentFieldDatas via SubjectAndDefaultLangcode
func GetCommentFieldDatasBySubjectAndDefaultLangcode(offset int, limit int, Subject_ string, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("subject = ? and default_langcode = ?", Subject_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByUidAndName Get CommentFieldDatas via UidAndName
func GetCommentFieldDatasByUidAndName(offset int, limit int, Uid_ int, Name_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("uid = ? and name = ?", Uid_, Name_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByUidAndMail Get CommentFieldDatas via UidAndMail
func GetCommentFieldDatasByUidAndMail(offset int, limit int, Uid_ int, Mail_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("uid = ? and mail = ?", Uid_, Mail_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByUidAndHomepage Get CommentFieldDatas via UidAndHomepage
func GetCommentFieldDatasByUidAndHomepage(offset int, limit int, Uid_ int, Homepage_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("uid = ? and homepage = ?", Uid_, Homepage_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByUidAndHostname Get CommentFieldDatas via UidAndHostname
func GetCommentFieldDatasByUidAndHostname(offset int, limit int, Uid_ int, Hostname_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("uid = ? and hostname = ?", Uid_, Hostname_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByUidAndCreated Get CommentFieldDatas via UidAndCreated
func GetCommentFieldDatasByUidAndCreated(offset int, limit int, Uid_ int, Created_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("uid = ? and created = ?", Uid_, Created_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByUidAndChanged Get CommentFieldDatas via UidAndChanged
func GetCommentFieldDatasByUidAndChanged(offset int, limit int, Uid_ int, Changed_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("uid = ? and changed = ?", Uid_, Changed_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByUidAndStatus Get CommentFieldDatas via UidAndStatus
func GetCommentFieldDatasByUidAndStatus(offset int, limit int, Uid_ int, Status_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("uid = ? and status = ?", Uid_, Status_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByUidAndThread Get CommentFieldDatas via UidAndThread
func GetCommentFieldDatasByUidAndThread(offset int, limit int, Uid_ int, Thread_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("uid = ? and thread = ?", Uid_, Thread_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByUidAndEntityType Get CommentFieldDatas via UidAndEntityType
func GetCommentFieldDatasByUidAndEntityType(offset int, limit int, Uid_ int, EntityType_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("uid = ? and entity_type = ?", Uid_, EntityType_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByUidAndFieldName Get CommentFieldDatas via UidAndFieldName
func GetCommentFieldDatasByUidAndFieldName(offset int, limit int, Uid_ int, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("uid = ? and field_name = ?", Uid_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByUidAndDefaultLangcode Get CommentFieldDatas via UidAndDefaultLangcode
func GetCommentFieldDatasByUidAndDefaultLangcode(offset int, limit int, Uid_ int, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("uid = ? and default_langcode = ?", Uid_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByNameAndMail Get CommentFieldDatas via NameAndMail
func GetCommentFieldDatasByNameAndMail(offset int, limit int, Name_ string, Mail_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("name = ? and mail = ?", Name_, Mail_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByNameAndHomepage Get CommentFieldDatas via NameAndHomepage
func GetCommentFieldDatasByNameAndHomepage(offset int, limit int, Name_ string, Homepage_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("name = ? and homepage = ?", Name_, Homepage_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByNameAndHostname Get CommentFieldDatas via NameAndHostname
func GetCommentFieldDatasByNameAndHostname(offset int, limit int, Name_ string, Hostname_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("name = ? and hostname = ?", Name_, Hostname_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByNameAndCreated Get CommentFieldDatas via NameAndCreated
func GetCommentFieldDatasByNameAndCreated(offset int, limit int, Name_ string, Created_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("name = ? and created = ?", Name_, Created_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByNameAndChanged Get CommentFieldDatas via NameAndChanged
func GetCommentFieldDatasByNameAndChanged(offset int, limit int, Name_ string, Changed_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("name = ? and changed = ?", Name_, Changed_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByNameAndStatus Get CommentFieldDatas via NameAndStatus
func GetCommentFieldDatasByNameAndStatus(offset int, limit int, Name_ string, Status_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("name = ? and status = ?", Name_, Status_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByNameAndThread Get CommentFieldDatas via NameAndThread
func GetCommentFieldDatasByNameAndThread(offset int, limit int, Name_ string, Thread_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("name = ? and thread = ?", Name_, Thread_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByNameAndEntityType Get CommentFieldDatas via NameAndEntityType
func GetCommentFieldDatasByNameAndEntityType(offset int, limit int, Name_ string, EntityType_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("name = ? and entity_type = ?", Name_, EntityType_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByNameAndFieldName Get CommentFieldDatas via NameAndFieldName
func GetCommentFieldDatasByNameAndFieldName(offset int, limit int, Name_ string, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("name = ? and field_name = ?", Name_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByNameAndDefaultLangcode Get CommentFieldDatas via NameAndDefaultLangcode
func GetCommentFieldDatasByNameAndDefaultLangcode(offset int, limit int, Name_ string, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("name = ? and default_langcode = ?", Name_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByMailAndHomepage Get CommentFieldDatas via MailAndHomepage
func GetCommentFieldDatasByMailAndHomepage(offset int, limit int, Mail_ string, Homepage_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("mail = ? and homepage = ?", Mail_, Homepage_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByMailAndHostname Get CommentFieldDatas via MailAndHostname
func GetCommentFieldDatasByMailAndHostname(offset int, limit int, Mail_ string, Hostname_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("mail = ? and hostname = ?", Mail_, Hostname_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByMailAndCreated Get CommentFieldDatas via MailAndCreated
func GetCommentFieldDatasByMailAndCreated(offset int, limit int, Mail_ string, Created_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("mail = ? and created = ?", Mail_, Created_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByMailAndChanged Get CommentFieldDatas via MailAndChanged
func GetCommentFieldDatasByMailAndChanged(offset int, limit int, Mail_ string, Changed_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("mail = ? and changed = ?", Mail_, Changed_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByMailAndStatus Get CommentFieldDatas via MailAndStatus
func GetCommentFieldDatasByMailAndStatus(offset int, limit int, Mail_ string, Status_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("mail = ? and status = ?", Mail_, Status_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByMailAndThread Get CommentFieldDatas via MailAndThread
func GetCommentFieldDatasByMailAndThread(offset int, limit int, Mail_ string, Thread_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("mail = ? and thread = ?", Mail_, Thread_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByMailAndEntityType Get CommentFieldDatas via MailAndEntityType
func GetCommentFieldDatasByMailAndEntityType(offset int, limit int, Mail_ string, EntityType_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("mail = ? and entity_type = ?", Mail_, EntityType_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByMailAndFieldName Get CommentFieldDatas via MailAndFieldName
func GetCommentFieldDatasByMailAndFieldName(offset int, limit int, Mail_ string, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("mail = ? and field_name = ?", Mail_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByMailAndDefaultLangcode Get CommentFieldDatas via MailAndDefaultLangcode
func GetCommentFieldDatasByMailAndDefaultLangcode(offset int, limit int, Mail_ string, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("mail = ? and default_langcode = ?", Mail_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByHomepageAndHostname Get CommentFieldDatas via HomepageAndHostname
func GetCommentFieldDatasByHomepageAndHostname(offset int, limit int, Homepage_ string, Hostname_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("homepage = ? and hostname = ?", Homepage_, Hostname_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByHomepageAndCreated Get CommentFieldDatas via HomepageAndCreated
func GetCommentFieldDatasByHomepageAndCreated(offset int, limit int, Homepage_ string, Created_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("homepage = ? and created = ?", Homepage_, Created_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByHomepageAndChanged Get CommentFieldDatas via HomepageAndChanged
func GetCommentFieldDatasByHomepageAndChanged(offset int, limit int, Homepage_ string, Changed_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("homepage = ? and changed = ?", Homepage_, Changed_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByHomepageAndStatus Get CommentFieldDatas via HomepageAndStatus
func GetCommentFieldDatasByHomepageAndStatus(offset int, limit int, Homepage_ string, Status_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("homepage = ? and status = ?", Homepage_, Status_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByHomepageAndThread Get CommentFieldDatas via HomepageAndThread
func GetCommentFieldDatasByHomepageAndThread(offset int, limit int, Homepage_ string, Thread_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("homepage = ? and thread = ?", Homepage_, Thread_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByHomepageAndEntityType Get CommentFieldDatas via HomepageAndEntityType
func GetCommentFieldDatasByHomepageAndEntityType(offset int, limit int, Homepage_ string, EntityType_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("homepage = ? and entity_type = ?", Homepage_, EntityType_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByHomepageAndFieldName Get CommentFieldDatas via HomepageAndFieldName
func GetCommentFieldDatasByHomepageAndFieldName(offset int, limit int, Homepage_ string, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("homepage = ? and field_name = ?", Homepage_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByHomepageAndDefaultLangcode Get CommentFieldDatas via HomepageAndDefaultLangcode
func GetCommentFieldDatasByHomepageAndDefaultLangcode(offset int, limit int, Homepage_ string, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("homepage = ? and default_langcode = ?", Homepage_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByHostnameAndCreated Get CommentFieldDatas via HostnameAndCreated
func GetCommentFieldDatasByHostnameAndCreated(offset int, limit int, Hostname_ string, Created_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("hostname = ? and created = ?", Hostname_, Created_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByHostnameAndChanged Get CommentFieldDatas via HostnameAndChanged
func GetCommentFieldDatasByHostnameAndChanged(offset int, limit int, Hostname_ string, Changed_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("hostname = ? and changed = ?", Hostname_, Changed_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByHostnameAndStatus Get CommentFieldDatas via HostnameAndStatus
func GetCommentFieldDatasByHostnameAndStatus(offset int, limit int, Hostname_ string, Status_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("hostname = ? and status = ?", Hostname_, Status_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByHostnameAndThread Get CommentFieldDatas via HostnameAndThread
func GetCommentFieldDatasByHostnameAndThread(offset int, limit int, Hostname_ string, Thread_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("hostname = ? and thread = ?", Hostname_, Thread_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByHostnameAndEntityType Get CommentFieldDatas via HostnameAndEntityType
func GetCommentFieldDatasByHostnameAndEntityType(offset int, limit int, Hostname_ string, EntityType_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("hostname = ? and entity_type = ?", Hostname_, EntityType_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByHostnameAndFieldName Get CommentFieldDatas via HostnameAndFieldName
func GetCommentFieldDatasByHostnameAndFieldName(offset int, limit int, Hostname_ string, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("hostname = ? and field_name = ?", Hostname_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByHostnameAndDefaultLangcode Get CommentFieldDatas via HostnameAndDefaultLangcode
func GetCommentFieldDatasByHostnameAndDefaultLangcode(offset int, limit int, Hostname_ string, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("hostname = ? and default_langcode = ?", Hostname_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCreatedAndChanged Get CommentFieldDatas via CreatedAndChanged
func GetCommentFieldDatasByCreatedAndChanged(offset int, limit int, Created_ int, Changed_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("created = ? and changed = ?", Created_, Changed_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCreatedAndStatus Get CommentFieldDatas via CreatedAndStatus
func GetCommentFieldDatasByCreatedAndStatus(offset int, limit int, Created_ int, Status_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("created = ? and status = ?", Created_, Status_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCreatedAndThread Get CommentFieldDatas via CreatedAndThread
func GetCommentFieldDatasByCreatedAndThread(offset int, limit int, Created_ int, Thread_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("created = ? and thread = ?", Created_, Thread_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCreatedAndEntityType Get CommentFieldDatas via CreatedAndEntityType
func GetCommentFieldDatasByCreatedAndEntityType(offset int, limit int, Created_ int, EntityType_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("created = ? and entity_type = ?", Created_, EntityType_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCreatedAndFieldName Get CommentFieldDatas via CreatedAndFieldName
func GetCommentFieldDatasByCreatedAndFieldName(offset int, limit int, Created_ int, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("created = ? and field_name = ?", Created_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByCreatedAndDefaultLangcode Get CommentFieldDatas via CreatedAndDefaultLangcode
func GetCommentFieldDatasByCreatedAndDefaultLangcode(offset int, limit int, Created_ int, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("created = ? and default_langcode = ?", Created_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByChangedAndStatus Get CommentFieldDatas via ChangedAndStatus
func GetCommentFieldDatasByChangedAndStatus(offset int, limit int, Changed_ int, Status_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("changed = ? and status = ?", Changed_, Status_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByChangedAndThread Get CommentFieldDatas via ChangedAndThread
func GetCommentFieldDatasByChangedAndThread(offset int, limit int, Changed_ int, Thread_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("changed = ? and thread = ?", Changed_, Thread_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByChangedAndEntityType Get CommentFieldDatas via ChangedAndEntityType
func GetCommentFieldDatasByChangedAndEntityType(offset int, limit int, Changed_ int, EntityType_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("changed = ? and entity_type = ?", Changed_, EntityType_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByChangedAndFieldName Get CommentFieldDatas via ChangedAndFieldName
func GetCommentFieldDatasByChangedAndFieldName(offset int, limit int, Changed_ int, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("changed = ? and field_name = ?", Changed_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByChangedAndDefaultLangcode Get CommentFieldDatas via ChangedAndDefaultLangcode
func GetCommentFieldDatasByChangedAndDefaultLangcode(offset int, limit int, Changed_ int, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("changed = ? and default_langcode = ?", Changed_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByStatusAndThread Get CommentFieldDatas via StatusAndThread
func GetCommentFieldDatasByStatusAndThread(offset int, limit int, Status_ int, Thread_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("status = ? and thread = ?", Status_, Thread_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByStatusAndEntityType Get CommentFieldDatas via StatusAndEntityType
func GetCommentFieldDatasByStatusAndEntityType(offset int, limit int, Status_ int, EntityType_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("status = ? and entity_type = ?", Status_, EntityType_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByStatusAndFieldName Get CommentFieldDatas via StatusAndFieldName
func GetCommentFieldDatasByStatusAndFieldName(offset int, limit int, Status_ int, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("status = ? and field_name = ?", Status_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByStatusAndDefaultLangcode Get CommentFieldDatas via StatusAndDefaultLangcode
func GetCommentFieldDatasByStatusAndDefaultLangcode(offset int, limit int, Status_ int, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("status = ? and default_langcode = ?", Status_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByThreadAndEntityType Get CommentFieldDatas via ThreadAndEntityType
func GetCommentFieldDatasByThreadAndEntityType(offset int, limit int, Thread_ string, EntityType_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("thread = ? and entity_type = ?", Thread_, EntityType_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByThreadAndFieldName Get CommentFieldDatas via ThreadAndFieldName
func GetCommentFieldDatasByThreadAndFieldName(offset int, limit int, Thread_ string, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("thread = ? and field_name = ?", Thread_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByThreadAndDefaultLangcode Get CommentFieldDatas via ThreadAndDefaultLangcode
func GetCommentFieldDatasByThreadAndDefaultLangcode(offset int, limit int, Thread_ string, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("thread = ? and default_langcode = ?", Thread_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByEntityTypeAndFieldName Get CommentFieldDatas via EntityTypeAndFieldName
func GetCommentFieldDatasByEntityTypeAndFieldName(offset int, limit int, EntityType_ string, FieldName_ string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("entity_type = ? and field_name = ?", EntityType_, FieldName_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByEntityTypeAndDefaultLangcode Get CommentFieldDatas via EntityTypeAndDefaultLangcode
func GetCommentFieldDatasByEntityTypeAndDefaultLangcode(offset int, limit int, EntityType_ string, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("entity_type = ? and default_langcode = ?", EntityType_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasByFieldNameAndDefaultLangcode Get CommentFieldDatas via FieldNameAndDefaultLangcode
func GetCommentFieldDatasByFieldNameAndDefaultLangcode(offset int, limit int, FieldName_ string, DefaultLangcode_ int) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("field_name = ? and default_langcode = ?", FieldName_, DefaultLangcode_).Limit(limit, offset).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatas Get CommentFieldDatas via field
func GetCommentFieldDatas(offset int, limit int, field string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Limit(limit, offset).Desc(field).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasViaCid Get CommentFieldDatas via Cid
func GetCommentFieldDatasViaCid(offset int, limit int, Cid_ int64, field string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("cid = ?", Cid_).Limit(limit, offset).Desc(field).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasViaCommentType Get CommentFieldDatas via CommentType
func GetCommentFieldDatasViaCommentType(offset int, limit int, CommentType_ string, field string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("comment_type = ?", CommentType_).Limit(limit, offset).Desc(field).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasViaLangcode Get CommentFieldDatas via Langcode
func GetCommentFieldDatasViaLangcode(offset int, limit int, Langcode_ string, field string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("langcode = ?", Langcode_).Limit(limit, offset).Desc(field).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasViaPid Get CommentFieldDatas via Pid
func GetCommentFieldDatasViaPid(offset int, limit int, Pid_ int, field string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("pid = ?", Pid_).Limit(limit, offset).Desc(field).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasViaEntityId Get CommentFieldDatas via EntityId
func GetCommentFieldDatasViaEntityId(offset int, limit int, EntityId_ int, field string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("entity_id = ?", EntityId_).Limit(limit, offset).Desc(field).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasViaSubject Get CommentFieldDatas via Subject
func GetCommentFieldDatasViaSubject(offset int, limit int, Subject_ string, field string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("subject = ?", Subject_).Limit(limit, offset).Desc(field).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasViaUid Get CommentFieldDatas via Uid
func GetCommentFieldDatasViaUid(offset int, limit int, Uid_ int, field string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("uid = ?", Uid_).Limit(limit, offset).Desc(field).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasViaName Get CommentFieldDatas via Name
func GetCommentFieldDatasViaName(offset int, limit int, Name_ string, field string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("name = ?", Name_).Limit(limit, offset).Desc(field).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasViaMail Get CommentFieldDatas via Mail
func GetCommentFieldDatasViaMail(offset int, limit int, Mail_ string, field string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("mail = ?", Mail_).Limit(limit, offset).Desc(field).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasViaHomepage Get CommentFieldDatas via Homepage
func GetCommentFieldDatasViaHomepage(offset int, limit int, Homepage_ string, field string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("homepage = ?", Homepage_).Limit(limit, offset).Desc(field).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasViaHostname Get CommentFieldDatas via Hostname
func GetCommentFieldDatasViaHostname(offset int, limit int, Hostname_ string, field string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("hostname = ?", Hostname_).Limit(limit, offset).Desc(field).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasViaCreated Get CommentFieldDatas via Created
func GetCommentFieldDatasViaCreated(offset int, limit int, Created_ int, field string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("created = ?", Created_).Limit(limit, offset).Desc(field).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasViaChanged Get CommentFieldDatas via Changed
func GetCommentFieldDatasViaChanged(offset int, limit int, Changed_ int, field string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("changed = ?", Changed_).Limit(limit, offset).Desc(field).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasViaStatus Get CommentFieldDatas via Status
func GetCommentFieldDatasViaStatus(offset int, limit int, Status_ int, field string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("status = ?", Status_).Limit(limit, offset).Desc(field).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasViaThread Get CommentFieldDatas via Thread
func GetCommentFieldDatasViaThread(offset int, limit int, Thread_ string, field string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("thread = ?", Thread_).Limit(limit, offset).Desc(field).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasViaEntityType Get CommentFieldDatas via EntityType
func GetCommentFieldDatasViaEntityType(offset int, limit int, EntityType_ string, field string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("entity_type = ?", EntityType_).Limit(limit, offset).Desc(field).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasViaFieldName Get CommentFieldDatas via FieldName
func GetCommentFieldDatasViaFieldName(offset int, limit int, FieldName_ string, field string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("field_name = ?", FieldName_).Limit(limit, offset).Desc(field).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// GetCommentFieldDatasViaDefaultLangcode Get CommentFieldDatas via DefaultLangcode
func GetCommentFieldDatasViaDefaultLangcode(offset int, limit int, DefaultLangcode_ int, field string) (*[]*CommentFieldData, error) {
	var _CommentFieldData = new([]*CommentFieldData)
	err := Engine.Table("comment_field_data").Where("default_langcode = ?", DefaultLangcode_).Limit(limit, offset).Desc(field).Find(_CommentFieldData)
	return _CommentFieldData, err
}

// HasCommentFieldDataViaCid Has CommentFieldData via Cid
func HasCommentFieldDataViaCid(iCid int64) bool {
	if has, err := Engine.Where("cid = ?", iCid).Get(new(CommentFieldData)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasCommentFieldDataViaCommentType Has CommentFieldData via CommentType
func HasCommentFieldDataViaCommentType(iCommentType string) bool {
	if has, err := Engine.Where("comment_type = ?", iCommentType).Get(new(CommentFieldData)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasCommentFieldDataViaLangcode Has CommentFieldData via Langcode
func HasCommentFieldDataViaLangcode(iLangcode string) bool {
	if has, err := Engine.Where("langcode = ?", iLangcode).Get(new(CommentFieldData)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasCommentFieldDataViaPid Has CommentFieldData via Pid
func HasCommentFieldDataViaPid(iPid int) bool {
	if has, err := Engine.Where("pid = ?", iPid).Get(new(CommentFieldData)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasCommentFieldDataViaEntityId Has CommentFieldData via EntityId
func HasCommentFieldDataViaEntityId(iEntityId int) bool {
	if has, err := Engine.Where("entity_id = ?", iEntityId).Get(new(CommentFieldData)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasCommentFieldDataViaSubject Has CommentFieldData via Subject
func HasCommentFieldDataViaSubject(iSubject string) bool {
	if has, err := Engine.Where("subject = ?", iSubject).Get(new(CommentFieldData)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasCommentFieldDataViaUid Has CommentFieldData via Uid
func HasCommentFieldDataViaUid(iUid int) bool {
	if has, err := Engine.Where("uid = ?", iUid).Get(new(CommentFieldData)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasCommentFieldDataViaName Has CommentFieldData via Name
func HasCommentFieldDataViaName(iName string) bool {
	if has, err := Engine.Where("name = ?", iName).Get(new(CommentFieldData)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasCommentFieldDataViaMail Has CommentFieldData via Mail
func HasCommentFieldDataViaMail(iMail string) bool {
	if has, err := Engine.Where("mail = ?", iMail).Get(new(CommentFieldData)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasCommentFieldDataViaHomepage Has CommentFieldData via Homepage
func HasCommentFieldDataViaHomepage(iHomepage string) bool {
	if has, err := Engine.Where("homepage = ?", iHomepage).Get(new(CommentFieldData)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasCommentFieldDataViaHostname Has CommentFieldData via Hostname
func HasCommentFieldDataViaHostname(iHostname string) bool {
	if has, err := Engine.Where("hostname = ?", iHostname).Get(new(CommentFieldData)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasCommentFieldDataViaCreated Has CommentFieldData via Created
func HasCommentFieldDataViaCreated(iCreated int) bool {
	if has, err := Engine.Where("created = ?", iCreated).Get(new(CommentFieldData)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasCommentFieldDataViaChanged Has CommentFieldData via Changed
func HasCommentFieldDataViaChanged(iChanged int) bool {
	if has, err := Engine.Where("changed = ?", iChanged).Get(new(CommentFieldData)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasCommentFieldDataViaStatus Has CommentFieldData via Status
func HasCommentFieldDataViaStatus(iStatus int) bool {
	if has, err := Engine.Where("status = ?", iStatus).Get(new(CommentFieldData)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasCommentFieldDataViaThread Has CommentFieldData via Thread
func HasCommentFieldDataViaThread(iThread string) bool {
	if has, err := Engine.Where("thread = ?", iThread).Get(new(CommentFieldData)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasCommentFieldDataViaEntityType Has CommentFieldData via EntityType
func HasCommentFieldDataViaEntityType(iEntityType string) bool {
	if has, err := Engine.Where("entity_type = ?", iEntityType).Get(new(CommentFieldData)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasCommentFieldDataViaFieldName Has CommentFieldData via FieldName
func HasCommentFieldDataViaFieldName(iFieldName string) bool {
	if has, err := Engine.Where("field_name = ?", iFieldName).Get(new(CommentFieldData)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasCommentFieldDataViaDefaultLangcode Has CommentFieldData via DefaultLangcode
func HasCommentFieldDataViaDefaultLangcode(iDefaultLangcode int) bool {
	if has, err := Engine.Where("default_langcode = ?", iDefaultLangcode).Get(new(CommentFieldData)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// GetCommentFieldDataViaCid Get CommentFieldData via Cid
func GetCommentFieldDataViaCid(iCid int64) (*CommentFieldData, error) {
	var _CommentFieldData = &CommentFieldData{Cid: iCid}
	has, err := Engine.Get(_CommentFieldData)
	if has {
		return _CommentFieldData, err
	} else {
		return nil, err
	}
}

// GetCommentFieldDataViaCommentType Get CommentFieldData via CommentType
func GetCommentFieldDataViaCommentType(iCommentType string) (*CommentFieldData, error) {
	var _CommentFieldData = &CommentFieldData{CommentType: iCommentType}
	has, err := Engine.Get(_CommentFieldData)
	if has {
		return _CommentFieldData, err
	} else {
		return nil, err
	}
}

// GetCommentFieldDataViaLangcode Get CommentFieldData via Langcode
func GetCommentFieldDataViaLangcode(iLangcode string) (*CommentFieldData, error) {
	var _CommentFieldData = &CommentFieldData{Langcode: iLangcode}
	has, err := Engine.Get(_CommentFieldData)
	if has {
		return _CommentFieldData, err
	} else {
		return nil, err
	}
}

// GetCommentFieldDataViaPid Get CommentFieldData via Pid
func GetCommentFieldDataViaPid(iPid int) (*CommentFieldData, error) {
	var _CommentFieldData = &CommentFieldData{Pid: iPid}
	has, err := Engine.Get(_CommentFieldData)
	if has {
		return _CommentFieldData, err
	} else {
		return nil, err
	}
}

// GetCommentFieldDataViaEntityId Get CommentFieldData via EntityId
func GetCommentFieldDataViaEntityId(iEntityId int) (*CommentFieldData, error) {
	var _CommentFieldData = &CommentFieldData{EntityId: iEntityId}
	has, err := Engine.Get(_CommentFieldData)
	if has {
		return _CommentFieldData, err
	} else {
		return nil, err
	}
}

// GetCommentFieldDataViaSubject Get CommentFieldData via Subject
func GetCommentFieldDataViaSubject(iSubject string) (*CommentFieldData, error) {
	var _CommentFieldData = &CommentFieldData{Subject: iSubject}
	has, err := Engine.Get(_CommentFieldData)
	if has {
		return _CommentFieldData, err
	} else {
		return nil, err
	}
}

// GetCommentFieldDataViaUid Get CommentFieldData via Uid
func GetCommentFieldDataViaUid(iUid int) (*CommentFieldData, error) {
	var _CommentFieldData = &CommentFieldData{Uid: iUid}
	has, err := Engine.Get(_CommentFieldData)
	if has {
		return _CommentFieldData, err
	} else {
		return nil, err
	}
}

// GetCommentFieldDataViaName Get CommentFieldData via Name
func GetCommentFieldDataViaName(iName string) (*CommentFieldData, error) {
	var _CommentFieldData = &CommentFieldData{Name: iName}
	has, err := Engine.Get(_CommentFieldData)
	if has {
		return _CommentFieldData, err
	} else {
		return nil, err
	}
}

// GetCommentFieldDataViaMail Get CommentFieldData via Mail
func GetCommentFieldDataViaMail(iMail string) (*CommentFieldData, error) {
	var _CommentFieldData = &CommentFieldData{Mail: iMail}
	has, err := Engine.Get(_CommentFieldData)
	if has {
		return _CommentFieldData, err
	} else {
		return nil, err
	}
}

// GetCommentFieldDataViaHomepage Get CommentFieldData via Homepage
func GetCommentFieldDataViaHomepage(iHomepage string) (*CommentFieldData, error) {
	var _CommentFieldData = &CommentFieldData{Homepage: iHomepage}
	has, err := Engine.Get(_CommentFieldData)
	if has {
		return _CommentFieldData, err
	} else {
		return nil, err
	}
}

// GetCommentFieldDataViaHostname Get CommentFieldData via Hostname
func GetCommentFieldDataViaHostname(iHostname string) (*CommentFieldData, error) {
	var _CommentFieldData = &CommentFieldData{Hostname: iHostname}
	has, err := Engine.Get(_CommentFieldData)
	if has {
		return _CommentFieldData, err
	} else {
		return nil, err
	}
}

// GetCommentFieldDataViaCreated Get CommentFieldData via Created
func GetCommentFieldDataViaCreated(iCreated int) (*CommentFieldData, error) {
	var _CommentFieldData = &CommentFieldData{Created: iCreated}
	has, err := Engine.Get(_CommentFieldData)
	if has {
		return _CommentFieldData, err
	} else {
		return nil, err
	}
}

// GetCommentFieldDataViaChanged Get CommentFieldData via Changed
func GetCommentFieldDataViaChanged(iChanged int) (*CommentFieldData, error) {
	var _CommentFieldData = &CommentFieldData{Changed: iChanged}
	has, err := Engine.Get(_CommentFieldData)
	if has {
		return _CommentFieldData, err
	} else {
		return nil, err
	}
}

// GetCommentFieldDataViaStatus Get CommentFieldData via Status
func GetCommentFieldDataViaStatus(iStatus int) (*CommentFieldData, error) {
	var _CommentFieldData = &CommentFieldData{Status: iStatus}
	has, err := Engine.Get(_CommentFieldData)
	if has {
		return _CommentFieldData, err
	} else {
		return nil, err
	}
}

// GetCommentFieldDataViaThread Get CommentFieldData via Thread
func GetCommentFieldDataViaThread(iThread string) (*CommentFieldData, error) {
	var _CommentFieldData = &CommentFieldData{Thread: iThread}
	has, err := Engine.Get(_CommentFieldData)
	if has {
		return _CommentFieldData, err
	} else {
		return nil, err
	}
}

// GetCommentFieldDataViaEntityType Get CommentFieldData via EntityType
func GetCommentFieldDataViaEntityType(iEntityType string) (*CommentFieldData, error) {
	var _CommentFieldData = &CommentFieldData{EntityType: iEntityType}
	has, err := Engine.Get(_CommentFieldData)
	if has {
		return _CommentFieldData, err
	} else {
		return nil, err
	}
}

// GetCommentFieldDataViaFieldName Get CommentFieldData via FieldName
func GetCommentFieldDataViaFieldName(iFieldName string) (*CommentFieldData, error) {
	var _CommentFieldData = &CommentFieldData{FieldName: iFieldName}
	has, err := Engine.Get(_CommentFieldData)
	if has {
		return _CommentFieldData, err
	} else {
		return nil, err
	}
}

// GetCommentFieldDataViaDefaultLangcode Get CommentFieldData via DefaultLangcode
func GetCommentFieldDataViaDefaultLangcode(iDefaultLangcode int) (*CommentFieldData, error) {
	var _CommentFieldData = &CommentFieldData{DefaultLangcode: iDefaultLangcode}
	has, err := Engine.Get(_CommentFieldData)
	if has {
		return _CommentFieldData, err
	} else {
		return nil, err
	}
}

// SetCommentFieldDataViaCid Set CommentFieldData via Cid
func SetCommentFieldDataViaCid(iCid int64, comment_field_data *CommentFieldData) (int64, error) {
	comment_field_data.Cid = iCid
	return Engine.Insert(comment_field_data)
}

// SetCommentFieldDataViaCommentType Set CommentFieldData via CommentType
func SetCommentFieldDataViaCommentType(iCommentType string, comment_field_data *CommentFieldData) (int64, error) {
	comment_field_data.CommentType = iCommentType
	return Engine.Insert(comment_field_data)
}

// SetCommentFieldDataViaLangcode Set CommentFieldData via Langcode
func SetCommentFieldDataViaLangcode(iLangcode string, comment_field_data *CommentFieldData) (int64, error) {
	comment_field_data.Langcode = iLangcode
	return Engine.Insert(comment_field_data)
}

// SetCommentFieldDataViaPid Set CommentFieldData via Pid
func SetCommentFieldDataViaPid(iPid int, comment_field_data *CommentFieldData) (int64, error) {
	comment_field_data.Pid = iPid
	return Engine.Insert(comment_field_data)
}

// SetCommentFieldDataViaEntityId Set CommentFieldData via EntityId
func SetCommentFieldDataViaEntityId(iEntityId int, comment_field_data *CommentFieldData) (int64, error) {
	comment_field_data.EntityId = iEntityId
	return Engine.Insert(comment_field_data)
}

// SetCommentFieldDataViaSubject Set CommentFieldData via Subject
func SetCommentFieldDataViaSubject(iSubject string, comment_field_data *CommentFieldData) (int64, error) {
	comment_field_data.Subject = iSubject
	return Engine.Insert(comment_field_data)
}

// SetCommentFieldDataViaUid Set CommentFieldData via Uid
func SetCommentFieldDataViaUid(iUid int, comment_field_data *CommentFieldData) (int64, error) {
	comment_field_data.Uid = iUid
	return Engine.Insert(comment_field_data)
}

// SetCommentFieldDataViaName Set CommentFieldData via Name
func SetCommentFieldDataViaName(iName string, comment_field_data *CommentFieldData) (int64, error) {
	comment_field_data.Name = iName
	return Engine.Insert(comment_field_data)
}

// SetCommentFieldDataViaMail Set CommentFieldData via Mail
func SetCommentFieldDataViaMail(iMail string, comment_field_data *CommentFieldData) (int64, error) {
	comment_field_data.Mail = iMail
	return Engine.Insert(comment_field_data)
}

// SetCommentFieldDataViaHomepage Set CommentFieldData via Homepage
func SetCommentFieldDataViaHomepage(iHomepage string, comment_field_data *CommentFieldData) (int64, error) {
	comment_field_data.Homepage = iHomepage
	return Engine.Insert(comment_field_data)
}

// SetCommentFieldDataViaHostname Set CommentFieldData via Hostname
func SetCommentFieldDataViaHostname(iHostname string, comment_field_data *CommentFieldData) (int64, error) {
	comment_field_data.Hostname = iHostname
	return Engine.Insert(comment_field_data)
}

// SetCommentFieldDataViaCreated Set CommentFieldData via Created
func SetCommentFieldDataViaCreated(iCreated int, comment_field_data *CommentFieldData) (int64, error) {
	comment_field_data.Created = iCreated
	return Engine.Insert(comment_field_data)
}

// SetCommentFieldDataViaChanged Set CommentFieldData via Changed
func SetCommentFieldDataViaChanged(iChanged int, comment_field_data *CommentFieldData) (int64, error) {
	comment_field_data.Changed = iChanged
	return Engine.Insert(comment_field_data)
}

// SetCommentFieldDataViaStatus Set CommentFieldData via Status
func SetCommentFieldDataViaStatus(iStatus int, comment_field_data *CommentFieldData) (int64, error) {
	comment_field_data.Status = iStatus
	return Engine.Insert(comment_field_data)
}

// SetCommentFieldDataViaThread Set CommentFieldData via Thread
func SetCommentFieldDataViaThread(iThread string, comment_field_data *CommentFieldData) (int64, error) {
	comment_field_data.Thread = iThread
	return Engine.Insert(comment_field_data)
}

// SetCommentFieldDataViaEntityType Set CommentFieldData via EntityType
func SetCommentFieldDataViaEntityType(iEntityType string, comment_field_data *CommentFieldData) (int64, error) {
	comment_field_data.EntityType = iEntityType
	return Engine.Insert(comment_field_data)
}

// SetCommentFieldDataViaFieldName Set CommentFieldData via FieldName
func SetCommentFieldDataViaFieldName(iFieldName string, comment_field_data *CommentFieldData) (int64, error) {
	comment_field_data.FieldName = iFieldName
	return Engine.Insert(comment_field_data)
}

// SetCommentFieldDataViaDefaultLangcode Set CommentFieldData via DefaultLangcode
func SetCommentFieldDataViaDefaultLangcode(iDefaultLangcode int, comment_field_data *CommentFieldData) (int64, error) {
	comment_field_data.DefaultLangcode = iDefaultLangcode
	return Engine.Insert(comment_field_data)
}

// AddCommentFieldData Add CommentFieldData via all columns
func AddCommentFieldData(iCid int64, iCommentType string, iLangcode string, iPid int, iEntityId int, iSubject string, iUid int, iName string, iMail string, iHomepage string, iHostname string, iCreated int, iChanged int, iStatus int, iThread string, iEntityType string, iFieldName string, iDefaultLangcode int) error {
	_CommentFieldData := &CommentFieldData{Cid: iCid, CommentType: iCommentType, Langcode: iLangcode, Pid: iPid, EntityId: iEntityId, Subject: iSubject, Uid: iUid, Name: iName, Mail: iMail, Homepage: iHomepage, Hostname: iHostname, Created: iCreated, Changed: iChanged, Status: iStatus, Thread: iThread, EntityType: iEntityType, FieldName: iFieldName, DefaultLangcode: iDefaultLangcode}
	if _, err := Engine.Insert(_CommentFieldData); err != nil {
		return err
	} else {
		return nil
	}
}

// PostCommentFieldData Post CommentFieldData via iCommentFieldData
func PostCommentFieldData(iCommentFieldData *CommentFieldData) (int64, error) {
	_, err := Engine.Insert(iCommentFieldData)
	return iCommentFieldData.Cid, err
}

// PutCommentFieldData Put CommentFieldData
func PutCommentFieldData(iCommentFieldData *CommentFieldData) (int64, error) {
	_, err := Engine.Id(iCommentFieldData.Cid).Update(iCommentFieldData)
	return iCommentFieldData.Cid, err
}

// PutCommentFieldDataViaCid Put CommentFieldData via Cid
func PutCommentFieldDataViaCid(Cid_ int64, iCommentFieldData *CommentFieldData) (int64, error) {
	row, err := Engine.Update(iCommentFieldData, &CommentFieldData{Cid: Cid_})
	return row, err
}

// PutCommentFieldDataViaCommentType Put CommentFieldData via CommentType
func PutCommentFieldDataViaCommentType(CommentType_ string, iCommentFieldData *CommentFieldData) (int64, error) {
	row, err := Engine.Update(iCommentFieldData, &CommentFieldData{CommentType: CommentType_})
	return row, err
}

// PutCommentFieldDataViaLangcode Put CommentFieldData via Langcode
func PutCommentFieldDataViaLangcode(Langcode_ string, iCommentFieldData *CommentFieldData) (int64, error) {
	row, err := Engine.Update(iCommentFieldData, &CommentFieldData{Langcode: Langcode_})
	return row, err
}

// PutCommentFieldDataViaPid Put CommentFieldData via Pid
func PutCommentFieldDataViaPid(Pid_ int, iCommentFieldData *CommentFieldData) (int64, error) {
	row, err := Engine.Update(iCommentFieldData, &CommentFieldData{Pid: Pid_})
	return row, err
}

// PutCommentFieldDataViaEntityId Put CommentFieldData via EntityId
func PutCommentFieldDataViaEntityId(EntityId_ int, iCommentFieldData *CommentFieldData) (int64, error) {
	row, err := Engine.Update(iCommentFieldData, &CommentFieldData{EntityId: EntityId_})
	return row, err
}

// PutCommentFieldDataViaSubject Put CommentFieldData via Subject
func PutCommentFieldDataViaSubject(Subject_ string, iCommentFieldData *CommentFieldData) (int64, error) {
	row, err := Engine.Update(iCommentFieldData, &CommentFieldData{Subject: Subject_})
	return row, err
}

// PutCommentFieldDataViaUid Put CommentFieldData via Uid
func PutCommentFieldDataViaUid(Uid_ int, iCommentFieldData *CommentFieldData) (int64, error) {
	row, err := Engine.Update(iCommentFieldData, &CommentFieldData{Uid: Uid_})
	return row, err
}

// PutCommentFieldDataViaName Put CommentFieldData via Name
func PutCommentFieldDataViaName(Name_ string, iCommentFieldData *CommentFieldData) (int64, error) {
	row, err := Engine.Update(iCommentFieldData, &CommentFieldData{Name: Name_})
	return row, err
}

// PutCommentFieldDataViaMail Put CommentFieldData via Mail
func PutCommentFieldDataViaMail(Mail_ string, iCommentFieldData *CommentFieldData) (int64, error) {
	row, err := Engine.Update(iCommentFieldData, &CommentFieldData{Mail: Mail_})
	return row, err
}

// PutCommentFieldDataViaHomepage Put CommentFieldData via Homepage
func PutCommentFieldDataViaHomepage(Homepage_ string, iCommentFieldData *CommentFieldData) (int64, error) {
	row, err := Engine.Update(iCommentFieldData, &CommentFieldData{Homepage: Homepage_})
	return row, err
}

// PutCommentFieldDataViaHostname Put CommentFieldData via Hostname
func PutCommentFieldDataViaHostname(Hostname_ string, iCommentFieldData *CommentFieldData) (int64, error) {
	row, err := Engine.Update(iCommentFieldData, &CommentFieldData{Hostname: Hostname_})
	return row, err
}

// PutCommentFieldDataViaCreated Put CommentFieldData via Created
func PutCommentFieldDataViaCreated(Created_ int, iCommentFieldData *CommentFieldData) (int64, error) {
	row, err := Engine.Update(iCommentFieldData, &CommentFieldData{Created: Created_})
	return row, err
}

// PutCommentFieldDataViaChanged Put CommentFieldData via Changed
func PutCommentFieldDataViaChanged(Changed_ int, iCommentFieldData *CommentFieldData) (int64, error) {
	row, err := Engine.Update(iCommentFieldData, &CommentFieldData{Changed: Changed_})
	return row, err
}

// PutCommentFieldDataViaStatus Put CommentFieldData via Status
func PutCommentFieldDataViaStatus(Status_ int, iCommentFieldData *CommentFieldData) (int64, error) {
	row, err := Engine.Update(iCommentFieldData, &CommentFieldData{Status: Status_})
	return row, err
}

// PutCommentFieldDataViaThread Put CommentFieldData via Thread
func PutCommentFieldDataViaThread(Thread_ string, iCommentFieldData *CommentFieldData) (int64, error) {
	row, err := Engine.Update(iCommentFieldData, &CommentFieldData{Thread: Thread_})
	return row, err
}

// PutCommentFieldDataViaEntityType Put CommentFieldData via EntityType
func PutCommentFieldDataViaEntityType(EntityType_ string, iCommentFieldData *CommentFieldData) (int64, error) {
	row, err := Engine.Update(iCommentFieldData, &CommentFieldData{EntityType: EntityType_})
	return row, err
}

// PutCommentFieldDataViaFieldName Put CommentFieldData via FieldName
func PutCommentFieldDataViaFieldName(FieldName_ string, iCommentFieldData *CommentFieldData) (int64, error) {
	row, err := Engine.Update(iCommentFieldData, &CommentFieldData{FieldName: FieldName_})
	return row, err
}

// PutCommentFieldDataViaDefaultLangcode Put CommentFieldData via DefaultLangcode
func PutCommentFieldDataViaDefaultLangcode(DefaultLangcode_ int, iCommentFieldData *CommentFieldData) (int64, error) {
	row, err := Engine.Update(iCommentFieldData, &CommentFieldData{DefaultLangcode: DefaultLangcode_})
	return row, err
}

// UpdateCommentFieldDataViaCid via map[string]interface{}{}
func UpdateCommentFieldDataViaCid(iCid int64, iCommentFieldDataMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(CommentFieldData)).Where("cid = ?", iCid).Update(iCommentFieldDataMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateCommentFieldDataViaCommentType via map[string]interface{}{}
func UpdateCommentFieldDataViaCommentType(iCommentType string, iCommentFieldDataMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(CommentFieldData)).Where("comment_type = ?", iCommentType).Update(iCommentFieldDataMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateCommentFieldDataViaLangcode via map[string]interface{}{}
func UpdateCommentFieldDataViaLangcode(iLangcode string, iCommentFieldDataMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(CommentFieldData)).Where("langcode = ?", iLangcode).Update(iCommentFieldDataMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateCommentFieldDataViaPid via map[string]interface{}{}
func UpdateCommentFieldDataViaPid(iPid int, iCommentFieldDataMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(CommentFieldData)).Where("pid = ?", iPid).Update(iCommentFieldDataMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateCommentFieldDataViaEntityId via map[string]interface{}{}
func UpdateCommentFieldDataViaEntityId(iEntityId int, iCommentFieldDataMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(CommentFieldData)).Where("entity_id = ?", iEntityId).Update(iCommentFieldDataMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateCommentFieldDataViaSubject via map[string]interface{}{}
func UpdateCommentFieldDataViaSubject(iSubject string, iCommentFieldDataMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(CommentFieldData)).Where("subject = ?", iSubject).Update(iCommentFieldDataMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateCommentFieldDataViaUid via map[string]interface{}{}
func UpdateCommentFieldDataViaUid(iUid int, iCommentFieldDataMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(CommentFieldData)).Where("uid = ?", iUid).Update(iCommentFieldDataMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateCommentFieldDataViaName via map[string]interface{}{}
func UpdateCommentFieldDataViaName(iName string, iCommentFieldDataMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(CommentFieldData)).Where("name = ?", iName).Update(iCommentFieldDataMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateCommentFieldDataViaMail via map[string]interface{}{}
func UpdateCommentFieldDataViaMail(iMail string, iCommentFieldDataMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(CommentFieldData)).Where("mail = ?", iMail).Update(iCommentFieldDataMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateCommentFieldDataViaHomepage via map[string]interface{}{}
func UpdateCommentFieldDataViaHomepage(iHomepage string, iCommentFieldDataMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(CommentFieldData)).Where("homepage = ?", iHomepage).Update(iCommentFieldDataMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateCommentFieldDataViaHostname via map[string]interface{}{}
func UpdateCommentFieldDataViaHostname(iHostname string, iCommentFieldDataMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(CommentFieldData)).Where("hostname = ?", iHostname).Update(iCommentFieldDataMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateCommentFieldDataViaCreated via map[string]interface{}{}
func UpdateCommentFieldDataViaCreated(iCreated int, iCommentFieldDataMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(CommentFieldData)).Where("created = ?", iCreated).Update(iCommentFieldDataMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateCommentFieldDataViaChanged via map[string]interface{}{}
func UpdateCommentFieldDataViaChanged(iChanged int, iCommentFieldDataMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(CommentFieldData)).Where("changed = ?", iChanged).Update(iCommentFieldDataMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateCommentFieldDataViaStatus via map[string]interface{}{}
func UpdateCommentFieldDataViaStatus(iStatus int, iCommentFieldDataMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(CommentFieldData)).Where("status = ?", iStatus).Update(iCommentFieldDataMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateCommentFieldDataViaThread via map[string]interface{}{}
func UpdateCommentFieldDataViaThread(iThread string, iCommentFieldDataMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(CommentFieldData)).Where("thread = ?", iThread).Update(iCommentFieldDataMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateCommentFieldDataViaEntityType via map[string]interface{}{}
func UpdateCommentFieldDataViaEntityType(iEntityType string, iCommentFieldDataMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(CommentFieldData)).Where("entity_type = ?", iEntityType).Update(iCommentFieldDataMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateCommentFieldDataViaFieldName via map[string]interface{}{}
func UpdateCommentFieldDataViaFieldName(iFieldName string, iCommentFieldDataMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(CommentFieldData)).Where("field_name = ?", iFieldName).Update(iCommentFieldDataMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateCommentFieldDataViaDefaultLangcode via map[string]interface{}{}
func UpdateCommentFieldDataViaDefaultLangcode(iDefaultLangcode int, iCommentFieldDataMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(CommentFieldData)).Where("default_langcode = ?", iDefaultLangcode).Update(iCommentFieldDataMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// DeleteCommentFieldData Delete CommentFieldData
func DeleteCommentFieldData(iCid int64) (int64, error) {
	row, err := Engine.Id(iCid).Delete(new(CommentFieldData))
	return row, err
}

// DeleteCommentFieldDataViaCid Delete CommentFieldData via Cid
func DeleteCommentFieldDataViaCid(iCid int64) (err error) {
	var has bool
	var _CommentFieldData = &CommentFieldData{Cid: iCid}
	if has, err = Engine.Get(_CommentFieldData); (has == true) && (err == nil) {
		if row, err := Engine.Where("cid = ?", iCid).Delete(new(CommentFieldData)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteCommentFieldDataViaCommentType Delete CommentFieldData via CommentType
func DeleteCommentFieldDataViaCommentType(iCommentType string) (err error) {
	var has bool
	var _CommentFieldData = &CommentFieldData{CommentType: iCommentType}
	if has, err = Engine.Get(_CommentFieldData); (has == true) && (err == nil) {
		if row, err := Engine.Where("comment_type = ?", iCommentType).Delete(new(CommentFieldData)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteCommentFieldDataViaLangcode Delete CommentFieldData via Langcode
func DeleteCommentFieldDataViaLangcode(iLangcode string) (err error) {
	var has bool
	var _CommentFieldData = &CommentFieldData{Langcode: iLangcode}
	if has, err = Engine.Get(_CommentFieldData); (has == true) && (err == nil) {
		if row, err := Engine.Where("langcode = ?", iLangcode).Delete(new(CommentFieldData)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteCommentFieldDataViaPid Delete CommentFieldData via Pid
func DeleteCommentFieldDataViaPid(iPid int) (err error) {
	var has bool
	var _CommentFieldData = &CommentFieldData{Pid: iPid}
	if has, err = Engine.Get(_CommentFieldData); (has == true) && (err == nil) {
		if row, err := Engine.Where("pid = ?", iPid).Delete(new(CommentFieldData)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteCommentFieldDataViaEntityId Delete CommentFieldData via EntityId
func DeleteCommentFieldDataViaEntityId(iEntityId int) (err error) {
	var has bool
	var _CommentFieldData = &CommentFieldData{EntityId: iEntityId}
	if has, err = Engine.Get(_CommentFieldData); (has == true) && (err == nil) {
		if row, err := Engine.Where("entity_id = ?", iEntityId).Delete(new(CommentFieldData)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteCommentFieldDataViaSubject Delete CommentFieldData via Subject
func DeleteCommentFieldDataViaSubject(iSubject string) (err error) {
	var has bool
	var _CommentFieldData = &CommentFieldData{Subject: iSubject}
	if has, err = Engine.Get(_CommentFieldData); (has == true) && (err == nil) {
		if row, err := Engine.Where("subject = ?", iSubject).Delete(new(CommentFieldData)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteCommentFieldDataViaUid Delete CommentFieldData via Uid
func DeleteCommentFieldDataViaUid(iUid int) (err error) {
	var has bool
	var _CommentFieldData = &CommentFieldData{Uid: iUid}
	if has, err = Engine.Get(_CommentFieldData); (has == true) && (err == nil) {
		if row, err := Engine.Where("uid = ?", iUid).Delete(new(CommentFieldData)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteCommentFieldDataViaName Delete CommentFieldData via Name
func DeleteCommentFieldDataViaName(iName string) (err error) {
	var has bool
	var _CommentFieldData = &CommentFieldData{Name: iName}
	if has, err = Engine.Get(_CommentFieldData); (has == true) && (err == nil) {
		if row, err := Engine.Where("name = ?", iName).Delete(new(CommentFieldData)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteCommentFieldDataViaMail Delete CommentFieldData via Mail
func DeleteCommentFieldDataViaMail(iMail string) (err error) {
	var has bool
	var _CommentFieldData = &CommentFieldData{Mail: iMail}
	if has, err = Engine.Get(_CommentFieldData); (has == true) && (err == nil) {
		if row, err := Engine.Where("mail = ?", iMail).Delete(new(CommentFieldData)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteCommentFieldDataViaHomepage Delete CommentFieldData via Homepage
func DeleteCommentFieldDataViaHomepage(iHomepage string) (err error) {
	var has bool
	var _CommentFieldData = &CommentFieldData{Homepage: iHomepage}
	if has, err = Engine.Get(_CommentFieldData); (has == true) && (err == nil) {
		if row, err := Engine.Where("homepage = ?", iHomepage).Delete(new(CommentFieldData)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteCommentFieldDataViaHostname Delete CommentFieldData via Hostname
func DeleteCommentFieldDataViaHostname(iHostname string) (err error) {
	var has bool
	var _CommentFieldData = &CommentFieldData{Hostname: iHostname}
	if has, err = Engine.Get(_CommentFieldData); (has == true) && (err == nil) {
		if row, err := Engine.Where("hostname = ?", iHostname).Delete(new(CommentFieldData)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteCommentFieldDataViaCreated Delete CommentFieldData via Created
func DeleteCommentFieldDataViaCreated(iCreated int) (err error) {
	var has bool
	var _CommentFieldData = &CommentFieldData{Created: iCreated}
	if has, err = Engine.Get(_CommentFieldData); (has == true) && (err == nil) {
		if row, err := Engine.Where("created = ?", iCreated).Delete(new(CommentFieldData)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteCommentFieldDataViaChanged Delete CommentFieldData via Changed
func DeleteCommentFieldDataViaChanged(iChanged int) (err error) {
	var has bool
	var _CommentFieldData = &CommentFieldData{Changed: iChanged}
	if has, err = Engine.Get(_CommentFieldData); (has == true) && (err == nil) {
		if row, err := Engine.Where("changed = ?", iChanged).Delete(new(CommentFieldData)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteCommentFieldDataViaStatus Delete CommentFieldData via Status
func DeleteCommentFieldDataViaStatus(iStatus int) (err error) {
	var has bool
	var _CommentFieldData = &CommentFieldData{Status: iStatus}
	if has, err = Engine.Get(_CommentFieldData); (has == true) && (err == nil) {
		if row, err := Engine.Where("status = ?", iStatus).Delete(new(CommentFieldData)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteCommentFieldDataViaThread Delete CommentFieldData via Thread
func DeleteCommentFieldDataViaThread(iThread string) (err error) {
	var has bool
	var _CommentFieldData = &CommentFieldData{Thread: iThread}
	if has, err = Engine.Get(_CommentFieldData); (has == true) && (err == nil) {
		if row, err := Engine.Where("thread = ?", iThread).Delete(new(CommentFieldData)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteCommentFieldDataViaEntityType Delete CommentFieldData via EntityType
func DeleteCommentFieldDataViaEntityType(iEntityType string) (err error) {
	var has bool
	var _CommentFieldData = &CommentFieldData{EntityType: iEntityType}
	if has, err = Engine.Get(_CommentFieldData); (has == true) && (err == nil) {
		if row, err := Engine.Where("entity_type = ?", iEntityType).Delete(new(CommentFieldData)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteCommentFieldDataViaFieldName Delete CommentFieldData via FieldName
func DeleteCommentFieldDataViaFieldName(iFieldName string) (err error) {
	var has bool
	var _CommentFieldData = &CommentFieldData{FieldName: iFieldName}
	if has, err = Engine.Get(_CommentFieldData); (has == true) && (err == nil) {
		if row, err := Engine.Where("field_name = ?", iFieldName).Delete(new(CommentFieldData)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteCommentFieldDataViaDefaultLangcode Delete CommentFieldData via DefaultLangcode
func DeleteCommentFieldDataViaDefaultLangcode(iDefaultLangcode int) (err error) {
	var has bool
	var _CommentFieldData = &CommentFieldData{DefaultLangcode: iDefaultLangcode}
	if has, err = Engine.Get(_CommentFieldData); (has == true) && (err == nil) {
		if row, err := Engine.Where("default_langcode = ?", iDefaultLangcode).Delete(new(CommentFieldData)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}
