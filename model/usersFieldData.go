package model

type UsersFieldData struct {
	Uid                    int64  `xorm:"not null pk index(user__id__default_langcode__langcode) INT(10)"`
	Langcode               string `xorm:"not null pk index(user__id__default_langcode__langcode) unique(user__name) VARCHAR(12)"`
	PreferredLangcode      string `xorm:"VARCHAR(12)"`
	PreferredAdminLangcode string `xorm:"VARCHAR(12)"`
	Name                   string `xorm:"not null unique(user__name) VARCHAR(60)"`
	Pass                   string `xorm:"VARCHAR(255)"`
	Mail                   string `xorm:"index VARCHAR(254)"`
	Timezone               string `xorm:"VARCHAR(32)"`
	Status                 int    `xorm:"TINYINT(4)"`
	Created                int    `xorm:"not null index INT(11)"`
	Changed                int    `xorm:"INT(11)"`
	Access                 int    `xorm:"not null index INT(11)"`
	Login                  int    `xorm:"INT(11)"`
	Init                   string `xorm:"VARCHAR(254)"`
	DefaultLangcode        int    `xorm:"not null index(user__id__default_langcode__langcode) TINYINT(4)"`
}

// GetUsersFieldDatasCount UsersFieldDatas' Count
func GetUsersFieldDatasCount(offset int, limit int) (int64, error) {
	total, err := Engine.Limit(limit, offset).Count(&UsersFieldData{})
	return total, err
}

// GetUsersFieldDataCountViaUid Get UsersFieldData via Uid
func GetUsersFieldDataCountViaUid(iUid int64) int64 {
	n, _ := Engine.Where("uid = ?", iUid).Count(&UsersFieldData{Uid: iUid})
	return n
}

// GetUsersFieldDataCountViaLangcode Get UsersFieldData via Langcode
func GetUsersFieldDataCountViaLangcode(iLangcode string) int64 {
	n, _ := Engine.Where("langcode = ?", iLangcode).Count(&UsersFieldData{Langcode: iLangcode})
	return n
}

// GetUsersFieldDataCountViaPreferredLangcode Get UsersFieldData via PreferredLangcode
func GetUsersFieldDataCountViaPreferredLangcode(iPreferredLangcode string) int64 {
	n, _ := Engine.Where("preferred_langcode = ?", iPreferredLangcode).Count(&UsersFieldData{PreferredLangcode: iPreferredLangcode})
	return n
}

// GetUsersFieldDataCountViaPreferredAdminLangcode Get UsersFieldData via PreferredAdminLangcode
func GetUsersFieldDataCountViaPreferredAdminLangcode(iPreferredAdminLangcode string) int64 {
	n, _ := Engine.Where("preferred_admin_langcode = ?", iPreferredAdminLangcode).Count(&UsersFieldData{PreferredAdminLangcode: iPreferredAdminLangcode})
	return n
}

// GetUsersFieldDataCountViaName Get UsersFieldData via Name
func GetUsersFieldDataCountViaName(iName string) int64 {
	n, _ := Engine.Where("name = ?", iName).Count(&UsersFieldData{Name: iName})
	return n
}

// GetUsersFieldDataCountViaPass Get UsersFieldData via Pass
func GetUsersFieldDataCountViaPass(iPass string) int64 {
	n, _ := Engine.Where("pass = ?", iPass).Count(&UsersFieldData{Pass: iPass})
	return n
}

// GetUsersFieldDataCountViaMail Get UsersFieldData via Mail
func GetUsersFieldDataCountViaMail(iMail string) int64 {
	n, _ := Engine.Where("mail = ?", iMail).Count(&UsersFieldData{Mail: iMail})
	return n
}

// GetUsersFieldDataCountViaTimezone Get UsersFieldData via Timezone
func GetUsersFieldDataCountViaTimezone(iTimezone string) int64 {
	n, _ := Engine.Where("timezone = ?", iTimezone).Count(&UsersFieldData{Timezone: iTimezone})
	return n
}

// GetUsersFieldDataCountViaStatus Get UsersFieldData via Status
func GetUsersFieldDataCountViaStatus(iStatus int) int64 {
	n, _ := Engine.Where("status = ?", iStatus).Count(&UsersFieldData{Status: iStatus})
	return n
}

// GetUsersFieldDataCountViaCreated Get UsersFieldData via Created
func GetUsersFieldDataCountViaCreated(iCreated int) int64 {
	n, _ := Engine.Where("created = ?", iCreated).Count(&UsersFieldData{Created: iCreated})
	return n
}

// GetUsersFieldDataCountViaChanged Get UsersFieldData via Changed
func GetUsersFieldDataCountViaChanged(iChanged int) int64 {
	n, _ := Engine.Where("changed = ?", iChanged).Count(&UsersFieldData{Changed: iChanged})
	return n
}

// GetUsersFieldDataCountViaAccess Get UsersFieldData via Access
func GetUsersFieldDataCountViaAccess(iAccess int) int64 {
	n, _ := Engine.Where("access = ?", iAccess).Count(&UsersFieldData{Access: iAccess})
	return n
}

// GetUsersFieldDataCountViaLogin Get UsersFieldData via Login
func GetUsersFieldDataCountViaLogin(iLogin int) int64 {
	n, _ := Engine.Where("login = ?", iLogin).Count(&UsersFieldData{Login: iLogin})
	return n
}

// GetUsersFieldDataCountViaInit Get UsersFieldData via Init
func GetUsersFieldDataCountViaInit(iInit string) int64 {
	n, _ := Engine.Where("init = ?", iInit).Count(&UsersFieldData{Init: iInit})
	return n
}

// GetUsersFieldDataCountViaDefaultLangcode Get UsersFieldData via DefaultLangcode
func GetUsersFieldDataCountViaDefaultLangcode(iDefaultLangcode int) int64 {
	n, _ := Engine.Where("default_langcode = ?", iDefaultLangcode).Count(&UsersFieldData{DefaultLangcode: iDefaultLangcode})
	return n
}

// GetUsersFieldDatasByUidAndLangcodeAndPreferredLangcode Get UsersFieldDatas via UidAndLangcodeAndPreferredLangcode
func GetUsersFieldDatasByUidAndLangcodeAndPreferredLangcode(offset int, limit int, Uid_ int64, Langcode_ string, PreferredLangcode_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("uid = ? and langcode = ? and preferred_langcode = ?", Uid_, Langcode_, PreferredLangcode_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByUidAndLangcodeAndPreferredAdminLangcode Get UsersFieldDatas via UidAndLangcodeAndPreferredAdminLangcode
func GetUsersFieldDatasByUidAndLangcodeAndPreferredAdminLangcode(offset int, limit int, Uid_ int64, Langcode_ string, PreferredAdminLangcode_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("uid = ? and langcode = ? and preferred_admin_langcode = ?", Uid_, Langcode_, PreferredAdminLangcode_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByUidAndLangcodeAndName Get UsersFieldDatas via UidAndLangcodeAndName
func GetUsersFieldDatasByUidAndLangcodeAndName(offset int, limit int, Uid_ int64, Langcode_ string, Name_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("uid = ? and langcode = ? and name = ?", Uid_, Langcode_, Name_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByUidAndLangcodeAndPass Get UsersFieldDatas via UidAndLangcodeAndPass
func GetUsersFieldDatasByUidAndLangcodeAndPass(offset int, limit int, Uid_ int64, Langcode_ string, Pass_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("uid = ? and langcode = ? and pass = ?", Uid_, Langcode_, Pass_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByUidAndLangcodeAndMail Get UsersFieldDatas via UidAndLangcodeAndMail
func GetUsersFieldDatasByUidAndLangcodeAndMail(offset int, limit int, Uid_ int64, Langcode_ string, Mail_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("uid = ? and langcode = ? and mail = ?", Uid_, Langcode_, Mail_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByUidAndLangcodeAndTimezone Get UsersFieldDatas via UidAndLangcodeAndTimezone
func GetUsersFieldDatasByUidAndLangcodeAndTimezone(offset int, limit int, Uid_ int64, Langcode_ string, Timezone_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("uid = ? and langcode = ? and timezone = ?", Uid_, Langcode_, Timezone_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByUidAndLangcodeAndStatus Get UsersFieldDatas via UidAndLangcodeAndStatus
func GetUsersFieldDatasByUidAndLangcodeAndStatus(offset int, limit int, Uid_ int64, Langcode_ string, Status_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("uid = ? and langcode = ? and status = ?", Uid_, Langcode_, Status_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByUidAndLangcodeAndCreated Get UsersFieldDatas via UidAndLangcodeAndCreated
func GetUsersFieldDatasByUidAndLangcodeAndCreated(offset int, limit int, Uid_ int64, Langcode_ string, Created_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("uid = ? and langcode = ? and created = ?", Uid_, Langcode_, Created_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByUidAndLangcodeAndChanged Get UsersFieldDatas via UidAndLangcodeAndChanged
func GetUsersFieldDatasByUidAndLangcodeAndChanged(offset int, limit int, Uid_ int64, Langcode_ string, Changed_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("uid = ? and langcode = ? and changed = ?", Uid_, Langcode_, Changed_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByUidAndLangcodeAndAccess Get UsersFieldDatas via UidAndLangcodeAndAccess
func GetUsersFieldDatasByUidAndLangcodeAndAccess(offset int, limit int, Uid_ int64, Langcode_ string, Access_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("uid = ? and langcode = ? and access = ?", Uid_, Langcode_, Access_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByUidAndLangcodeAndLogin Get UsersFieldDatas via UidAndLangcodeAndLogin
func GetUsersFieldDatasByUidAndLangcodeAndLogin(offset int, limit int, Uid_ int64, Langcode_ string, Login_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("uid = ? and langcode = ? and login = ?", Uid_, Langcode_, Login_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByUidAndLangcodeAndInit Get UsersFieldDatas via UidAndLangcodeAndInit
func GetUsersFieldDatasByUidAndLangcodeAndInit(offset int, limit int, Uid_ int64, Langcode_ string, Init_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("uid = ? and langcode = ? and init = ?", Uid_, Langcode_, Init_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByUidAndLangcodeAndDefaultLangcode Get UsersFieldDatas via UidAndLangcodeAndDefaultLangcode
func GetUsersFieldDatasByUidAndLangcodeAndDefaultLangcode(offset int, limit int, Uid_ int64, Langcode_ string, DefaultLangcode_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("uid = ? and langcode = ? and default_langcode = ?", Uid_, Langcode_, DefaultLangcode_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByUidAndPreferredLangcodeAndPreferredAdminLangcode Get UsersFieldDatas via UidAndPreferredLangcodeAndPreferredAdminLangcode
func GetUsersFieldDatasByUidAndPreferredLangcodeAndPreferredAdminLangcode(offset int, limit int, Uid_ int64, PreferredLangcode_ string, PreferredAdminLangcode_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("uid = ? and preferred_langcode = ? and preferred_admin_langcode = ?", Uid_, PreferredLangcode_, PreferredAdminLangcode_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByUidAndPreferredLangcodeAndName Get UsersFieldDatas via UidAndPreferredLangcodeAndName
func GetUsersFieldDatasByUidAndPreferredLangcodeAndName(offset int, limit int, Uid_ int64, PreferredLangcode_ string, Name_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("uid = ? and preferred_langcode = ? and name = ?", Uid_, PreferredLangcode_, Name_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByUidAndPreferredLangcodeAndPass Get UsersFieldDatas via UidAndPreferredLangcodeAndPass
func GetUsersFieldDatasByUidAndPreferredLangcodeAndPass(offset int, limit int, Uid_ int64, PreferredLangcode_ string, Pass_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("uid = ? and preferred_langcode = ? and pass = ?", Uid_, PreferredLangcode_, Pass_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByUidAndPreferredLangcodeAndMail Get UsersFieldDatas via UidAndPreferredLangcodeAndMail
func GetUsersFieldDatasByUidAndPreferredLangcodeAndMail(offset int, limit int, Uid_ int64, PreferredLangcode_ string, Mail_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("uid = ? and preferred_langcode = ? and mail = ?", Uid_, PreferredLangcode_, Mail_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByUidAndPreferredLangcodeAndTimezone Get UsersFieldDatas via UidAndPreferredLangcodeAndTimezone
func GetUsersFieldDatasByUidAndPreferredLangcodeAndTimezone(offset int, limit int, Uid_ int64, PreferredLangcode_ string, Timezone_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("uid = ? and preferred_langcode = ? and timezone = ?", Uid_, PreferredLangcode_, Timezone_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByUidAndPreferredLangcodeAndStatus Get UsersFieldDatas via UidAndPreferredLangcodeAndStatus
func GetUsersFieldDatasByUidAndPreferredLangcodeAndStatus(offset int, limit int, Uid_ int64, PreferredLangcode_ string, Status_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("uid = ? and preferred_langcode = ? and status = ?", Uid_, PreferredLangcode_, Status_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByUidAndPreferredLangcodeAndCreated Get UsersFieldDatas via UidAndPreferredLangcodeAndCreated
func GetUsersFieldDatasByUidAndPreferredLangcodeAndCreated(offset int, limit int, Uid_ int64, PreferredLangcode_ string, Created_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("uid = ? and preferred_langcode = ? and created = ?", Uid_, PreferredLangcode_, Created_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByUidAndPreferredLangcodeAndChanged Get UsersFieldDatas via UidAndPreferredLangcodeAndChanged
func GetUsersFieldDatasByUidAndPreferredLangcodeAndChanged(offset int, limit int, Uid_ int64, PreferredLangcode_ string, Changed_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("uid = ? and preferred_langcode = ? and changed = ?", Uid_, PreferredLangcode_, Changed_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByUidAndPreferredLangcodeAndAccess Get UsersFieldDatas via UidAndPreferredLangcodeAndAccess
func GetUsersFieldDatasByUidAndPreferredLangcodeAndAccess(offset int, limit int, Uid_ int64, PreferredLangcode_ string, Access_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("uid = ? and preferred_langcode = ? and access = ?", Uid_, PreferredLangcode_, Access_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByUidAndPreferredLangcodeAndLogin Get UsersFieldDatas via UidAndPreferredLangcodeAndLogin
func GetUsersFieldDatasByUidAndPreferredLangcodeAndLogin(offset int, limit int, Uid_ int64, PreferredLangcode_ string, Login_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("uid = ? and preferred_langcode = ? and login = ?", Uid_, PreferredLangcode_, Login_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByUidAndPreferredLangcodeAndInit Get UsersFieldDatas via UidAndPreferredLangcodeAndInit
func GetUsersFieldDatasByUidAndPreferredLangcodeAndInit(offset int, limit int, Uid_ int64, PreferredLangcode_ string, Init_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("uid = ? and preferred_langcode = ? and init = ?", Uid_, PreferredLangcode_, Init_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByUidAndPreferredLangcodeAndDefaultLangcode Get UsersFieldDatas via UidAndPreferredLangcodeAndDefaultLangcode
func GetUsersFieldDatasByUidAndPreferredLangcodeAndDefaultLangcode(offset int, limit int, Uid_ int64, PreferredLangcode_ string, DefaultLangcode_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("uid = ? and preferred_langcode = ? and default_langcode = ?", Uid_, PreferredLangcode_, DefaultLangcode_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByUidAndPreferredAdminLangcodeAndName Get UsersFieldDatas via UidAndPreferredAdminLangcodeAndName
func GetUsersFieldDatasByUidAndPreferredAdminLangcodeAndName(offset int, limit int, Uid_ int64, PreferredAdminLangcode_ string, Name_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("uid = ? and preferred_admin_langcode = ? and name = ?", Uid_, PreferredAdminLangcode_, Name_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByUidAndPreferredAdminLangcodeAndPass Get UsersFieldDatas via UidAndPreferredAdminLangcodeAndPass
func GetUsersFieldDatasByUidAndPreferredAdminLangcodeAndPass(offset int, limit int, Uid_ int64, PreferredAdminLangcode_ string, Pass_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("uid = ? and preferred_admin_langcode = ? and pass = ?", Uid_, PreferredAdminLangcode_, Pass_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByUidAndPreferredAdminLangcodeAndMail Get UsersFieldDatas via UidAndPreferredAdminLangcodeAndMail
func GetUsersFieldDatasByUidAndPreferredAdminLangcodeAndMail(offset int, limit int, Uid_ int64, PreferredAdminLangcode_ string, Mail_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("uid = ? and preferred_admin_langcode = ? and mail = ?", Uid_, PreferredAdminLangcode_, Mail_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByUidAndPreferredAdminLangcodeAndTimezone Get UsersFieldDatas via UidAndPreferredAdminLangcodeAndTimezone
func GetUsersFieldDatasByUidAndPreferredAdminLangcodeAndTimezone(offset int, limit int, Uid_ int64, PreferredAdminLangcode_ string, Timezone_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("uid = ? and preferred_admin_langcode = ? and timezone = ?", Uid_, PreferredAdminLangcode_, Timezone_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByUidAndPreferredAdminLangcodeAndStatus Get UsersFieldDatas via UidAndPreferredAdminLangcodeAndStatus
func GetUsersFieldDatasByUidAndPreferredAdminLangcodeAndStatus(offset int, limit int, Uid_ int64, PreferredAdminLangcode_ string, Status_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("uid = ? and preferred_admin_langcode = ? and status = ?", Uid_, PreferredAdminLangcode_, Status_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByUidAndPreferredAdminLangcodeAndCreated Get UsersFieldDatas via UidAndPreferredAdminLangcodeAndCreated
func GetUsersFieldDatasByUidAndPreferredAdminLangcodeAndCreated(offset int, limit int, Uid_ int64, PreferredAdminLangcode_ string, Created_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("uid = ? and preferred_admin_langcode = ? and created = ?", Uid_, PreferredAdminLangcode_, Created_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByUidAndPreferredAdminLangcodeAndChanged Get UsersFieldDatas via UidAndPreferredAdminLangcodeAndChanged
func GetUsersFieldDatasByUidAndPreferredAdminLangcodeAndChanged(offset int, limit int, Uid_ int64, PreferredAdminLangcode_ string, Changed_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("uid = ? and preferred_admin_langcode = ? and changed = ?", Uid_, PreferredAdminLangcode_, Changed_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByUidAndPreferredAdminLangcodeAndAccess Get UsersFieldDatas via UidAndPreferredAdminLangcodeAndAccess
func GetUsersFieldDatasByUidAndPreferredAdminLangcodeAndAccess(offset int, limit int, Uid_ int64, PreferredAdminLangcode_ string, Access_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("uid = ? and preferred_admin_langcode = ? and access = ?", Uid_, PreferredAdminLangcode_, Access_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByUidAndPreferredAdminLangcodeAndLogin Get UsersFieldDatas via UidAndPreferredAdminLangcodeAndLogin
func GetUsersFieldDatasByUidAndPreferredAdminLangcodeAndLogin(offset int, limit int, Uid_ int64, PreferredAdminLangcode_ string, Login_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("uid = ? and preferred_admin_langcode = ? and login = ?", Uid_, PreferredAdminLangcode_, Login_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByUidAndPreferredAdminLangcodeAndInit Get UsersFieldDatas via UidAndPreferredAdminLangcodeAndInit
func GetUsersFieldDatasByUidAndPreferredAdminLangcodeAndInit(offset int, limit int, Uid_ int64, PreferredAdminLangcode_ string, Init_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("uid = ? and preferred_admin_langcode = ? and init = ?", Uid_, PreferredAdminLangcode_, Init_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByUidAndPreferredAdminLangcodeAndDefaultLangcode Get UsersFieldDatas via UidAndPreferredAdminLangcodeAndDefaultLangcode
func GetUsersFieldDatasByUidAndPreferredAdminLangcodeAndDefaultLangcode(offset int, limit int, Uid_ int64, PreferredAdminLangcode_ string, DefaultLangcode_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("uid = ? and preferred_admin_langcode = ? and default_langcode = ?", Uid_, PreferredAdminLangcode_, DefaultLangcode_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByUidAndNameAndPass Get UsersFieldDatas via UidAndNameAndPass
func GetUsersFieldDatasByUidAndNameAndPass(offset int, limit int, Uid_ int64, Name_ string, Pass_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("uid = ? and name = ? and pass = ?", Uid_, Name_, Pass_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByUidAndNameAndMail Get UsersFieldDatas via UidAndNameAndMail
func GetUsersFieldDatasByUidAndNameAndMail(offset int, limit int, Uid_ int64, Name_ string, Mail_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("uid = ? and name = ? and mail = ?", Uid_, Name_, Mail_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByUidAndNameAndTimezone Get UsersFieldDatas via UidAndNameAndTimezone
func GetUsersFieldDatasByUidAndNameAndTimezone(offset int, limit int, Uid_ int64, Name_ string, Timezone_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("uid = ? and name = ? and timezone = ?", Uid_, Name_, Timezone_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByUidAndNameAndStatus Get UsersFieldDatas via UidAndNameAndStatus
func GetUsersFieldDatasByUidAndNameAndStatus(offset int, limit int, Uid_ int64, Name_ string, Status_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("uid = ? and name = ? and status = ?", Uid_, Name_, Status_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByUidAndNameAndCreated Get UsersFieldDatas via UidAndNameAndCreated
func GetUsersFieldDatasByUidAndNameAndCreated(offset int, limit int, Uid_ int64, Name_ string, Created_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("uid = ? and name = ? and created = ?", Uid_, Name_, Created_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByUidAndNameAndChanged Get UsersFieldDatas via UidAndNameAndChanged
func GetUsersFieldDatasByUidAndNameAndChanged(offset int, limit int, Uid_ int64, Name_ string, Changed_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("uid = ? and name = ? and changed = ?", Uid_, Name_, Changed_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByUidAndNameAndAccess Get UsersFieldDatas via UidAndNameAndAccess
func GetUsersFieldDatasByUidAndNameAndAccess(offset int, limit int, Uid_ int64, Name_ string, Access_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("uid = ? and name = ? and access = ?", Uid_, Name_, Access_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByUidAndNameAndLogin Get UsersFieldDatas via UidAndNameAndLogin
func GetUsersFieldDatasByUidAndNameAndLogin(offset int, limit int, Uid_ int64, Name_ string, Login_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("uid = ? and name = ? and login = ?", Uid_, Name_, Login_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByUidAndNameAndInit Get UsersFieldDatas via UidAndNameAndInit
func GetUsersFieldDatasByUidAndNameAndInit(offset int, limit int, Uid_ int64, Name_ string, Init_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("uid = ? and name = ? and init = ?", Uid_, Name_, Init_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByUidAndNameAndDefaultLangcode Get UsersFieldDatas via UidAndNameAndDefaultLangcode
func GetUsersFieldDatasByUidAndNameAndDefaultLangcode(offset int, limit int, Uid_ int64, Name_ string, DefaultLangcode_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("uid = ? and name = ? and default_langcode = ?", Uid_, Name_, DefaultLangcode_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByUidAndPassAndMail Get UsersFieldDatas via UidAndPassAndMail
func GetUsersFieldDatasByUidAndPassAndMail(offset int, limit int, Uid_ int64, Pass_ string, Mail_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("uid = ? and pass = ? and mail = ?", Uid_, Pass_, Mail_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByUidAndPassAndTimezone Get UsersFieldDatas via UidAndPassAndTimezone
func GetUsersFieldDatasByUidAndPassAndTimezone(offset int, limit int, Uid_ int64, Pass_ string, Timezone_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("uid = ? and pass = ? and timezone = ?", Uid_, Pass_, Timezone_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByUidAndPassAndStatus Get UsersFieldDatas via UidAndPassAndStatus
func GetUsersFieldDatasByUidAndPassAndStatus(offset int, limit int, Uid_ int64, Pass_ string, Status_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("uid = ? and pass = ? and status = ?", Uid_, Pass_, Status_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByUidAndPassAndCreated Get UsersFieldDatas via UidAndPassAndCreated
func GetUsersFieldDatasByUidAndPassAndCreated(offset int, limit int, Uid_ int64, Pass_ string, Created_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("uid = ? and pass = ? and created = ?", Uid_, Pass_, Created_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByUidAndPassAndChanged Get UsersFieldDatas via UidAndPassAndChanged
func GetUsersFieldDatasByUidAndPassAndChanged(offset int, limit int, Uid_ int64, Pass_ string, Changed_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("uid = ? and pass = ? and changed = ?", Uid_, Pass_, Changed_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByUidAndPassAndAccess Get UsersFieldDatas via UidAndPassAndAccess
func GetUsersFieldDatasByUidAndPassAndAccess(offset int, limit int, Uid_ int64, Pass_ string, Access_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("uid = ? and pass = ? and access = ?", Uid_, Pass_, Access_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByUidAndPassAndLogin Get UsersFieldDatas via UidAndPassAndLogin
func GetUsersFieldDatasByUidAndPassAndLogin(offset int, limit int, Uid_ int64, Pass_ string, Login_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("uid = ? and pass = ? and login = ?", Uid_, Pass_, Login_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByUidAndPassAndInit Get UsersFieldDatas via UidAndPassAndInit
func GetUsersFieldDatasByUidAndPassAndInit(offset int, limit int, Uid_ int64, Pass_ string, Init_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("uid = ? and pass = ? and init = ?", Uid_, Pass_, Init_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByUidAndPassAndDefaultLangcode Get UsersFieldDatas via UidAndPassAndDefaultLangcode
func GetUsersFieldDatasByUidAndPassAndDefaultLangcode(offset int, limit int, Uid_ int64, Pass_ string, DefaultLangcode_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("uid = ? and pass = ? and default_langcode = ?", Uid_, Pass_, DefaultLangcode_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByUidAndMailAndTimezone Get UsersFieldDatas via UidAndMailAndTimezone
func GetUsersFieldDatasByUidAndMailAndTimezone(offset int, limit int, Uid_ int64, Mail_ string, Timezone_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("uid = ? and mail = ? and timezone = ?", Uid_, Mail_, Timezone_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByUidAndMailAndStatus Get UsersFieldDatas via UidAndMailAndStatus
func GetUsersFieldDatasByUidAndMailAndStatus(offset int, limit int, Uid_ int64, Mail_ string, Status_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("uid = ? and mail = ? and status = ?", Uid_, Mail_, Status_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByUidAndMailAndCreated Get UsersFieldDatas via UidAndMailAndCreated
func GetUsersFieldDatasByUidAndMailAndCreated(offset int, limit int, Uid_ int64, Mail_ string, Created_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("uid = ? and mail = ? and created = ?", Uid_, Mail_, Created_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByUidAndMailAndChanged Get UsersFieldDatas via UidAndMailAndChanged
func GetUsersFieldDatasByUidAndMailAndChanged(offset int, limit int, Uid_ int64, Mail_ string, Changed_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("uid = ? and mail = ? and changed = ?", Uid_, Mail_, Changed_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByUidAndMailAndAccess Get UsersFieldDatas via UidAndMailAndAccess
func GetUsersFieldDatasByUidAndMailAndAccess(offset int, limit int, Uid_ int64, Mail_ string, Access_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("uid = ? and mail = ? and access = ?", Uid_, Mail_, Access_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByUidAndMailAndLogin Get UsersFieldDatas via UidAndMailAndLogin
func GetUsersFieldDatasByUidAndMailAndLogin(offset int, limit int, Uid_ int64, Mail_ string, Login_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("uid = ? and mail = ? and login = ?", Uid_, Mail_, Login_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByUidAndMailAndInit Get UsersFieldDatas via UidAndMailAndInit
func GetUsersFieldDatasByUidAndMailAndInit(offset int, limit int, Uid_ int64, Mail_ string, Init_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("uid = ? and mail = ? and init = ?", Uid_, Mail_, Init_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByUidAndMailAndDefaultLangcode Get UsersFieldDatas via UidAndMailAndDefaultLangcode
func GetUsersFieldDatasByUidAndMailAndDefaultLangcode(offset int, limit int, Uid_ int64, Mail_ string, DefaultLangcode_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("uid = ? and mail = ? and default_langcode = ?", Uid_, Mail_, DefaultLangcode_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByUidAndTimezoneAndStatus Get UsersFieldDatas via UidAndTimezoneAndStatus
func GetUsersFieldDatasByUidAndTimezoneAndStatus(offset int, limit int, Uid_ int64, Timezone_ string, Status_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("uid = ? and timezone = ? and status = ?", Uid_, Timezone_, Status_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByUidAndTimezoneAndCreated Get UsersFieldDatas via UidAndTimezoneAndCreated
func GetUsersFieldDatasByUidAndTimezoneAndCreated(offset int, limit int, Uid_ int64, Timezone_ string, Created_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("uid = ? and timezone = ? and created = ?", Uid_, Timezone_, Created_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByUidAndTimezoneAndChanged Get UsersFieldDatas via UidAndTimezoneAndChanged
func GetUsersFieldDatasByUidAndTimezoneAndChanged(offset int, limit int, Uid_ int64, Timezone_ string, Changed_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("uid = ? and timezone = ? and changed = ?", Uid_, Timezone_, Changed_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByUidAndTimezoneAndAccess Get UsersFieldDatas via UidAndTimezoneAndAccess
func GetUsersFieldDatasByUidAndTimezoneAndAccess(offset int, limit int, Uid_ int64, Timezone_ string, Access_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("uid = ? and timezone = ? and access = ?", Uid_, Timezone_, Access_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByUidAndTimezoneAndLogin Get UsersFieldDatas via UidAndTimezoneAndLogin
func GetUsersFieldDatasByUidAndTimezoneAndLogin(offset int, limit int, Uid_ int64, Timezone_ string, Login_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("uid = ? and timezone = ? and login = ?", Uid_, Timezone_, Login_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByUidAndTimezoneAndInit Get UsersFieldDatas via UidAndTimezoneAndInit
func GetUsersFieldDatasByUidAndTimezoneAndInit(offset int, limit int, Uid_ int64, Timezone_ string, Init_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("uid = ? and timezone = ? and init = ?", Uid_, Timezone_, Init_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByUidAndTimezoneAndDefaultLangcode Get UsersFieldDatas via UidAndTimezoneAndDefaultLangcode
func GetUsersFieldDatasByUidAndTimezoneAndDefaultLangcode(offset int, limit int, Uid_ int64, Timezone_ string, DefaultLangcode_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("uid = ? and timezone = ? and default_langcode = ?", Uid_, Timezone_, DefaultLangcode_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByUidAndStatusAndCreated Get UsersFieldDatas via UidAndStatusAndCreated
func GetUsersFieldDatasByUidAndStatusAndCreated(offset int, limit int, Uid_ int64, Status_ int, Created_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("uid = ? and status = ? and created = ?", Uid_, Status_, Created_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByUidAndStatusAndChanged Get UsersFieldDatas via UidAndStatusAndChanged
func GetUsersFieldDatasByUidAndStatusAndChanged(offset int, limit int, Uid_ int64, Status_ int, Changed_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("uid = ? and status = ? and changed = ?", Uid_, Status_, Changed_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByUidAndStatusAndAccess Get UsersFieldDatas via UidAndStatusAndAccess
func GetUsersFieldDatasByUidAndStatusAndAccess(offset int, limit int, Uid_ int64, Status_ int, Access_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("uid = ? and status = ? and access = ?", Uid_, Status_, Access_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByUidAndStatusAndLogin Get UsersFieldDatas via UidAndStatusAndLogin
func GetUsersFieldDatasByUidAndStatusAndLogin(offset int, limit int, Uid_ int64, Status_ int, Login_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("uid = ? and status = ? and login = ?", Uid_, Status_, Login_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByUidAndStatusAndInit Get UsersFieldDatas via UidAndStatusAndInit
func GetUsersFieldDatasByUidAndStatusAndInit(offset int, limit int, Uid_ int64, Status_ int, Init_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("uid = ? and status = ? and init = ?", Uid_, Status_, Init_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByUidAndStatusAndDefaultLangcode Get UsersFieldDatas via UidAndStatusAndDefaultLangcode
func GetUsersFieldDatasByUidAndStatusAndDefaultLangcode(offset int, limit int, Uid_ int64, Status_ int, DefaultLangcode_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("uid = ? and status = ? and default_langcode = ?", Uid_, Status_, DefaultLangcode_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByUidAndCreatedAndChanged Get UsersFieldDatas via UidAndCreatedAndChanged
func GetUsersFieldDatasByUidAndCreatedAndChanged(offset int, limit int, Uid_ int64, Created_ int, Changed_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("uid = ? and created = ? and changed = ?", Uid_, Created_, Changed_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByUidAndCreatedAndAccess Get UsersFieldDatas via UidAndCreatedAndAccess
func GetUsersFieldDatasByUidAndCreatedAndAccess(offset int, limit int, Uid_ int64, Created_ int, Access_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("uid = ? and created = ? and access = ?", Uid_, Created_, Access_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByUidAndCreatedAndLogin Get UsersFieldDatas via UidAndCreatedAndLogin
func GetUsersFieldDatasByUidAndCreatedAndLogin(offset int, limit int, Uid_ int64, Created_ int, Login_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("uid = ? and created = ? and login = ?", Uid_, Created_, Login_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByUidAndCreatedAndInit Get UsersFieldDatas via UidAndCreatedAndInit
func GetUsersFieldDatasByUidAndCreatedAndInit(offset int, limit int, Uid_ int64, Created_ int, Init_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("uid = ? and created = ? and init = ?", Uid_, Created_, Init_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByUidAndCreatedAndDefaultLangcode Get UsersFieldDatas via UidAndCreatedAndDefaultLangcode
func GetUsersFieldDatasByUidAndCreatedAndDefaultLangcode(offset int, limit int, Uid_ int64, Created_ int, DefaultLangcode_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("uid = ? and created = ? and default_langcode = ?", Uid_, Created_, DefaultLangcode_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByUidAndChangedAndAccess Get UsersFieldDatas via UidAndChangedAndAccess
func GetUsersFieldDatasByUidAndChangedAndAccess(offset int, limit int, Uid_ int64, Changed_ int, Access_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("uid = ? and changed = ? and access = ?", Uid_, Changed_, Access_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByUidAndChangedAndLogin Get UsersFieldDatas via UidAndChangedAndLogin
func GetUsersFieldDatasByUidAndChangedAndLogin(offset int, limit int, Uid_ int64, Changed_ int, Login_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("uid = ? and changed = ? and login = ?", Uid_, Changed_, Login_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByUidAndChangedAndInit Get UsersFieldDatas via UidAndChangedAndInit
func GetUsersFieldDatasByUidAndChangedAndInit(offset int, limit int, Uid_ int64, Changed_ int, Init_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("uid = ? and changed = ? and init = ?", Uid_, Changed_, Init_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByUidAndChangedAndDefaultLangcode Get UsersFieldDatas via UidAndChangedAndDefaultLangcode
func GetUsersFieldDatasByUidAndChangedAndDefaultLangcode(offset int, limit int, Uid_ int64, Changed_ int, DefaultLangcode_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("uid = ? and changed = ? and default_langcode = ?", Uid_, Changed_, DefaultLangcode_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByUidAndAccessAndLogin Get UsersFieldDatas via UidAndAccessAndLogin
func GetUsersFieldDatasByUidAndAccessAndLogin(offset int, limit int, Uid_ int64, Access_ int, Login_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("uid = ? and access = ? and login = ?", Uid_, Access_, Login_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByUidAndAccessAndInit Get UsersFieldDatas via UidAndAccessAndInit
func GetUsersFieldDatasByUidAndAccessAndInit(offset int, limit int, Uid_ int64, Access_ int, Init_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("uid = ? and access = ? and init = ?", Uid_, Access_, Init_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByUidAndAccessAndDefaultLangcode Get UsersFieldDatas via UidAndAccessAndDefaultLangcode
func GetUsersFieldDatasByUidAndAccessAndDefaultLangcode(offset int, limit int, Uid_ int64, Access_ int, DefaultLangcode_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("uid = ? and access = ? and default_langcode = ?", Uid_, Access_, DefaultLangcode_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByUidAndLoginAndInit Get UsersFieldDatas via UidAndLoginAndInit
func GetUsersFieldDatasByUidAndLoginAndInit(offset int, limit int, Uid_ int64, Login_ int, Init_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("uid = ? and login = ? and init = ?", Uid_, Login_, Init_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByUidAndLoginAndDefaultLangcode Get UsersFieldDatas via UidAndLoginAndDefaultLangcode
func GetUsersFieldDatasByUidAndLoginAndDefaultLangcode(offset int, limit int, Uid_ int64, Login_ int, DefaultLangcode_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("uid = ? and login = ? and default_langcode = ?", Uid_, Login_, DefaultLangcode_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByUidAndInitAndDefaultLangcode Get UsersFieldDatas via UidAndInitAndDefaultLangcode
func GetUsersFieldDatasByUidAndInitAndDefaultLangcode(offset int, limit int, Uid_ int64, Init_ string, DefaultLangcode_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("uid = ? and init = ? and default_langcode = ?", Uid_, Init_, DefaultLangcode_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByLangcodeAndPreferredLangcodeAndPreferredAdminLangcode Get UsersFieldDatas via LangcodeAndPreferredLangcodeAndPreferredAdminLangcode
func GetUsersFieldDatasByLangcodeAndPreferredLangcodeAndPreferredAdminLangcode(offset int, limit int, Langcode_ string, PreferredLangcode_ string, PreferredAdminLangcode_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("langcode = ? and preferred_langcode = ? and preferred_admin_langcode = ?", Langcode_, PreferredLangcode_, PreferredAdminLangcode_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByLangcodeAndPreferredLangcodeAndName Get UsersFieldDatas via LangcodeAndPreferredLangcodeAndName
func GetUsersFieldDatasByLangcodeAndPreferredLangcodeAndName(offset int, limit int, Langcode_ string, PreferredLangcode_ string, Name_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("langcode = ? and preferred_langcode = ? and name = ?", Langcode_, PreferredLangcode_, Name_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByLangcodeAndPreferredLangcodeAndPass Get UsersFieldDatas via LangcodeAndPreferredLangcodeAndPass
func GetUsersFieldDatasByLangcodeAndPreferredLangcodeAndPass(offset int, limit int, Langcode_ string, PreferredLangcode_ string, Pass_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("langcode = ? and preferred_langcode = ? and pass = ?", Langcode_, PreferredLangcode_, Pass_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByLangcodeAndPreferredLangcodeAndMail Get UsersFieldDatas via LangcodeAndPreferredLangcodeAndMail
func GetUsersFieldDatasByLangcodeAndPreferredLangcodeAndMail(offset int, limit int, Langcode_ string, PreferredLangcode_ string, Mail_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("langcode = ? and preferred_langcode = ? and mail = ?", Langcode_, PreferredLangcode_, Mail_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByLangcodeAndPreferredLangcodeAndTimezone Get UsersFieldDatas via LangcodeAndPreferredLangcodeAndTimezone
func GetUsersFieldDatasByLangcodeAndPreferredLangcodeAndTimezone(offset int, limit int, Langcode_ string, PreferredLangcode_ string, Timezone_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("langcode = ? and preferred_langcode = ? and timezone = ?", Langcode_, PreferredLangcode_, Timezone_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByLangcodeAndPreferredLangcodeAndStatus Get UsersFieldDatas via LangcodeAndPreferredLangcodeAndStatus
func GetUsersFieldDatasByLangcodeAndPreferredLangcodeAndStatus(offset int, limit int, Langcode_ string, PreferredLangcode_ string, Status_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("langcode = ? and preferred_langcode = ? and status = ?", Langcode_, PreferredLangcode_, Status_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByLangcodeAndPreferredLangcodeAndCreated Get UsersFieldDatas via LangcodeAndPreferredLangcodeAndCreated
func GetUsersFieldDatasByLangcodeAndPreferredLangcodeAndCreated(offset int, limit int, Langcode_ string, PreferredLangcode_ string, Created_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("langcode = ? and preferred_langcode = ? and created = ?", Langcode_, PreferredLangcode_, Created_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByLangcodeAndPreferredLangcodeAndChanged Get UsersFieldDatas via LangcodeAndPreferredLangcodeAndChanged
func GetUsersFieldDatasByLangcodeAndPreferredLangcodeAndChanged(offset int, limit int, Langcode_ string, PreferredLangcode_ string, Changed_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("langcode = ? and preferred_langcode = ? and changed = ?", Langcode_, PreferredLangcode_, Changed_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByLangcodeAndPreferredLangcodeAndAccess Get UsersFieldDatas via LangcodeAndPreferredLangcodeAndAccess
func GetUsersFieldDatasByLangcodeAndPreferredLangcodeAndAccess(offset int, limit int, Langcode_ string, PreferredLangcode_ string, Access_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("langcode = ? and preferred_langcode = ? and access = ?", Langcode_, PreferredLangcode_, Access_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByLangcodeAndPreferredLangcodeAndLogin Get UsersFieldDatas via LangcodeAndPreferredLangcodeAndLogin
func GetUsersFieldDatasByLangcodeAndPreferredLangcodeAndLogin(offset int, limit int, Langcode_ string, PreferredLangcode_ string, Login_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("langcode = ? and preferred_langcode = ? and login = ?", Langcode_, PreferredLangcode_, Login_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByLangcodeAndPreferredLangcodeAndInit Get UsersFieldDatas via LangcodeAndPreferredLangcodeAndInit
func GetUsersFieldDatasByLangcodeAndPreferredLangcodeAndInit(offset int, limit int, Langcode_ string, PreferredLangcode_ string, Init_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("langcode = ? and preferred_langcode = ? and init = ?", Langcode_, PreferredLangcode_, Init_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByLangcodeAndPreferredLangcodeAndDefaultLangcode Get UsersFieldDatas via LangcodeAndPreferredLangcodeAndDefaultLangcode
func GetUsersFieldDatasByLangcodeAndPreferredLangcodeAndDefaultLangcode(offset int, limit int, Langcode_ string, PreferredLangcode_ string, DefaultLangcode_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("langcode = ? and preferred_langcode = ? and default_langcode = ?", Langcode_, PreferredLangcode_, DefaultLangcode_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByLangcodeAndPreferredAdminLangcodeAndName Get UsersFieldDatas via LangcodeAndPreferredAdminLangcodeAndName
func GetUsersFieldDatasByLangcodeAndPreferredAdminLangcodeAndName(offset int, limit int, Langcode_ string, PreferredAdminLangcode_ string, Name_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("langcode = ? and preferred_admin_langcode = ? and name = ?", Langcode_, PreferredAdminLangcode_, Name_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByLangcodeAndPreferredAdminLangcodeAndPass Get UsersFieldDatas via LangcodeAndPreferredAdminLangcodeAndPass
func GetUsersFieldDatasByLangcodeAndPreferredAdminLangcodeAndPass(offset int, limit int, Langcode_ string, PreferredAdminLangcode_ string, Pass_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("langcode = ? and preferred_admin_langcode = ? and pass = ?", Langcode_, PreferredAdminLangcode_, Pass_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByLangcodeAndPreferredAdminLangcodeAndMail Get UsersFieldDatas via LangcodeAndPreferredAdminLangcodeAndMail
func GetUsersFieldDatasByLangcodeAndPreferredAdminLangcodeAndMail(offset int, limit int, Langcode_ string, PreferredAdminLangcode_ string, Mail_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("langcode = ? and preferred_admin_langcode = ? and mail = ?", Langcode_, PreferredAdminLangcode_, Mail_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByLangcodeAndPreferredAdminLangcodeAndTimezone Get UsersFieldDatas via LangcodeAndPreferredAdminLangcodeAndTimezone
func GetUsersFieldDatasByLangcodeAndPreferredAdminLangcodeAndTimezone(offset int, limit int, Langcode_ string, PreferredAdminLangcode_ string, Timezone_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("langcode = ? and preferred_admin_langcode = ? and timezone = ?", Langcode_, PreferredAdminLangcode_, Timezone_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByLangcodeAndPreferredAdminLangcodeAndStatus Get UsersFieldDatas via LangcodeAndPreferredAdminLangcodeAndStatus
func GetUsersFieldDatasByLangcodeAndPreferredAdminLangcodeAndStatus(offset int, limit int, Langcode_ string, PreferredAdminLangcode_ string, Status_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("langcode = ? and preferred_admin_langcode = ? and status = ?", Langcode_, PreferredAdminLangcode_, Status_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByLangcodeAndPreferredAdminLangcodeAndCreated Get UsersFieldDatas via LangcodeAndPreferredAdminLangcodeAndCreated
func GetUsersFieldDatasByLangcodeAndPreferredAdminLangcodeAndCreated(offset int, limit int, Langcode_ string, PreferredAdminLangcode_ string, Created_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("langcode = ? and preferred_admin_langcode = ? and created = ?", Langcode_, PreferredAdminLangcode_, Created_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByLangcodeAndPreferredAdminLangcodeAndChanged Get UsersFieldDatas via LangcodeAndPreferredAdminLangcodeAndChanged
func GetUsersFieldDatasByLangcodeAndPreferredAdminLangcodeAndChanged(offset int, limit int, Langcode_ string, PreferredAdminLangcode_ string, Changed_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("langcode = ? and preferred_admin_langcode = ? and changed = ?", Langcode_, PreferredAdminLangcode_, Changed_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByLangcodeAndPreferredAdminLangcodeAndAccess Get UsersFieldDatas via LangcodeAndPreferredAdminLangcodeAndAccess
func GetUsersFieldDatasByLangcodeAndPreferredAdminLangcodeAndAccess(offset int, limit int, Langcode_ string, PreferredAdminLangcode_ string, Access_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("langcode = ? and preferred_admin_langcode = ? and access = ?", Langcode_, PreferredAdminLangcode_, Access_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByLangcodeAndPreferredAdminLangcodeAndLogin Get UsersFieldDatas via LangcodeAndPreferredAdminLangcodeAndLogin
func GetUsersFieldDatasByLangcodeAndPreferredAdminLangcodeAndLogin(offset int, limit int, Langcode_ string, PreferredAdminLangcode_ string, Login_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("langcode = ? and preferred_admin_langcode = ? and login = ?", Langcode_, PreferredAdminLangcode_, Login_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByLangcodeAndPreferredAdminLangcodeAndInit Get UsersFieldDatas via LangcodeAndPreferredAdminLangcodeAndInit
func GetUsersFieldDatasByLangcodeAndPreferredAdminLangcodeAndInit(offset int, limit int, Langcode_ string, PreferredAdminLangcode_ string, Init_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("langcode = ? and preferred_admin_langcode = ? and init = ?", Langcode_, PreferredAdminLangcode_, Init_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByLangcodeAndPreferredAdminLangcodeAndDefaultLangcode Get UsersFieldDatas via LangcodeAndPreferredAdminLangcodeAndDefaultLangcode
func GetUsersFieldDatasByLangcodeAndPreferredAdminLangcodeAndDefaultLangcode(offset int, limit int, Langcode_ string, PreferredAdminLangcode_ string, DefaultLangcode_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("langcode = ? and preferred_admin_langcode = ? and default_langcode = ?", Langcode_, PreferredAdminLangcode_, DefaultLangcode_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByLangcodeAndNameAndPass Get UsersFieldDatas via LangcodeAndNameAndPass
func GetUsersFieldDatasByLangcodeAndNameAndPass(offset int, limit int, Langcode_ string, Name_ string, Pass_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("langcode = ? and name = ? and pass = ?", Langcode_, Name_, Pass_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByLangcodeAndNameAndMail Get UsersFieldDatas via LangcodeAndNameAndMail
func GetUsersFieldDatasByLangcodeAndNameAndMail(offset int, limit int, Langcode_ string, Name_ string, Mail_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("langcode = ? and name = ? and mail = ?", Langcode_, Name_, Mail_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByLangcodeAndNameAndTimezone Get UsersFieldDatas via LangcodeAndNameAndTimezone
func GetUsersFieldDatasByLangcodeAndNameAndTimezone(offset int, limit int, Langcode_ string, Name_ string, Timezone_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("langcode = ? and name = ? and timezone = ?", Langcode_, Name_, Timezone_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByLangcodeAndNameAndStatus Get UsersFieldDatas via LangcodeAndNameAndStatus
func GetUsersFieldDatasByLangcodeAndNameAndStatus(offset int, limit int, Langcode_ string, Name_ string, Status_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("langcode = ? and name = ? and status = ?", Langcode_, Name_, Status_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByLangcodeAndNameAndCreated Get UsersFieldDatas via LangcodeAndNameAndCreated
func GetUsersFieldDatasByLangcodeAndNameAndCreated(offset int, limit int, Langcode_ string, Name_ string, Created_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("langcode = ? and name = ? and created = ?", Langcode_, Name_, Created_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByLangcodeAndNameAndChanged Get UsersFieldDatas via LangcodeAndNameAndChanged
func GetUsersFieldDatasByLangcodeAndNameAndChanged(offset int, limit int, Langcode_ string, Name_ string, Changed_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("langcode = ? and name = ? and changed = ?", Langcode_, Name_, Changed_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByLangcodeAndNameAndAccess Get UsersFieldDatas via LangcodeAndNameAndAccess
func GetUsersFieldDatasByLangcodeAndNameAndAccess(offset int, limit int, Langcode_ string, Name_ string, Access_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("langcode = ? and name = ? and access = ?", Langcode_, Name_, Access_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByLangcodeAndNameAndLogin Get UsersFieldDatas via LangcodeAndNameAndLogin
func GetUsersFieldDatasByLangcodeAndNameAndLogin(offset int, limit int, Langcode_ string, Name_ string, Login_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("langcode = ? and name = ? and login = ?", Langcode_, Name_, Login_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByLangcodeAndNameAndInit Get UsersFieldDatas via LangcodeAndNameAndInit
func GetUsersFieldDatasByLangcodeAndNameAndInit(offset int, limit int, Langcode_ string, Name_ string, Init_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("langcode = ? and name = ? and init = ?", Langcode_, Name_, Init_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByLangcodeAndNameAndDefaultLangcode Get UsersFieldDatas via LangcodeAndNameAndDefaultLangcode
func GetUsersFieldDatasByLangcodeAndNameAndDefaultLangcode(offset int, limit int, Langcode_ string, Name_ string, DefaultLangcode_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("langcode = ? and name = ? and default_langcode = ?", Langcode_, Name_, DefaultLangcode_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByLangcodeAndPassAndMail Get UsersFieldDatas via LangcodeAndPassAndMail
func GetUsersFieldDatasByLangcodeAndPassAndMail(offset int, limit int, Langcode_ string, Pass_ string, Mail_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("langcode = ? and pass = ? and mail = ?", Langcode_, Pass_, Mail_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByLangcodeAndPassAndTimezone Get UsersFieldDatas via LangcodeAndPassAndTimezone
func GetUsersFieldDatasByLangcodeAndPassAndTimezone(offset int, limit int, Langcode_ string, Pass_ string, Timezone_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("langcode = ? and pass = ? and timezone = ?", Langcode_, Pass_, Timezone_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByLangcodeAndPassAndStatus Get UsersFieldDatas via LangcodeAndPassAndStatus
func GetUsersFieldDatasByLangcodeAndPassAndStatus(offset int, limit int, Langcode_ string, Pass_ string, Status_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("langcode = ? and pass = ? and status = ?", Langcode_, Pass_, Status_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByLangcodeAndPassAndCreated Get UsersFieldDatas via LangcodeAndPassAndCreated
func GetUsersFieldDatasByLangcodeAndPassAndCreated(offset int, limit int, Langcode_ string, Pass_ string, Created_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("langcode = ? and pass = ? and created = ?", Langcode_, Pass_, Created_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByLangcodeAndPassAndChanged Get UsersFieldDatas via LangcodeAndPassAndChanged
func GetUsersFieldDatasByLangcodeAndPassAndChanged(offset int, limit int, Langcode_ string, Pass_ string, Changed_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("langcode = ? and pass = ? and changed = ?", Langcode_, Pass_, Changed_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByLangcodeAndPassAndAccess Get UsersFieldDatas via LangcodeAndPassAndAccess
func GetUsersFieldDatasByLangcodeAndPassAndAccess(offset int, limit int, Langcode_ string, Pass_ string, Access_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("langcode = ? and pass = ? and access = ?", Langcode_, Pass_, Access_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByLangcodeAndPassAndLogin Get UsersFieldDatas via LangcodeAndPassAndLogin
func GetUsersFieldDatasByLangcodeAndPassAndLogin(offset int, limit int, Langcode_ string, Pass_ string, Login_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("langcode = ? and pass = ? and login = ?", Langcode_, Pass_, Login_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByLangcodeAndPassAndInit Get UsersFieldDatas via LangcodeAndPassAndInit
func GetUsersFieldDatasByLangcodeAndPassAndInit(offset int, limit int, Langcode_ string, Pass_ string, Init_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("langcode = ? and pass = ? and init = ?", Langcode_, Pass_, Init_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByLangcodeAndPassAndDefaultLangcode Get UsersFieldDatas via LangcodeAndPassAndDefaultLangcode
func GetUsersFieldDatasByLangcodeAndPassAndDefaultLangcode(offset int, limit int, Langcode_ string, Pass_ string, DefaultLangcode_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("langcode = ? and pass = ? and default_langcode = ?", Langcode_, Pass_, DefaultLangcode_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByLangcodeAndMailAndTimezone Get UsersFieldDatas via LangcodeAndMailAndTimezone
func GetUsersFieldDatasByLangcodeAndMailAndTimezone(offset int, limit int, Langcode_ string, Mail_ string, Timezone_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("langcode = ? and mail = ? and timezone = ?", Langcode_, Mail_, Timezone_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByLangcodeAndMailAndStatus Get UsersFieldDatas via LangcodeAndMailAndStatus
func GetUsersFieldDatasByLangcodeAndMailAndStatus(offset int, limit int, Langcode_ string, Mail_ string, Status_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("langcode = ? and mail = ? and status = ?", Langcode_, Mail_, Status_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByLangcodeAndMailAndCreated Get UsersFieldDatas via LangcodeAndMailAndCreated
func GetUsersFieldDatasByLangcodeAndMailAndCreated(offset int, limit int, Langcode_ string, Mail_ string, Created_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("langcode = ? and mail = ? and created = ?", Langcode_, Mail_, Created_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByLangcodeAndMailAndChanged Get UsersFieldDatas via LangcodeAndMailAndChanged
func GetUsersFieldDatasByLangcodeAndMailAndChanged(offset int, limit int, Langcode_ string, Mail_ string, Changed_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("langcode = ? and mail = ? and changed = ?", Langcode_, Mail_, Changed_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByLangcodeAndMailAndAccess Get UsersFieldDatas via LangcodeAndMailAndAccess
func GetUsersFieldDatasByLangcodeAndMailAndAccess(offset int, limit int, Langcode_ string, Mail_ string, Access_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("langcode = ? and mail = ? and access = ?", Langcode_, Mail_, Access_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByLangcodeAndMailAndLogin Get UsersFieldDatas via LangcodeAndMailAndLogin
func GetUsersFieldDatasByLangcodeAndMailAndLogin(offset int, limit int, Langcode_ string, Mail_ string, Login_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("langcode = ? and mail = ? and login = ?", Langcode_, Mail_, Login_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByLangcodeAndMailAndInit Get UsersFieldDatas via LangcodeAndMailAndInit
func GetUsersFieldDatasByLangcodeAndMailAndInit(offset int, limit int, Langcode_ string, Mail_ string, Init_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("langcode = ? and mail = ? and init = ?", Langcode_, Mail_, Init_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByLangcodeAndMailAndDefaultLangcode Get UsersFieldDatas via LangcodeAndMailAndDefaultLangcode
func GetUsersFieldDatasByLangcodeAndMailAndDefaultLangcode(offset int, limit int, Langcode_ string, Mail_ string, DefaultLangcode_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("langcode = ? and mail = ? and default_langcode = ?", Langcode_, Mail_, DefaultLangcode_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByLangcodeAndTimezoneAndStatus Get UsersFieldDatas via LangcodeAndTimezoneAndStatus
func GetUsersFieldDatasByLangcodeAndTimezoneAndStatus(offset int, limit int, Langcode_ string, Timezone_ string, Status_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("langcode = ? and timezone = ? and status = ?", Langcode_, Timezone_, Status_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByLangcodeAndTimezoneAndCreated Get UsersFieldDatas via LangcodeAndTimezoneAndCreated
func GetUsersFieldDatasByLangcodeAndTimezoneAndCreated(offset int, limit int, Langcode_ string, Timezone_ string, Created_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("langcode = ? and timezone = ? and created = ?", Langcode_, Timezone_, Created_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByLangcodeAndTimezoneAndChanged Get UsersFieldDatas via LangcodeAndTimezoneAndChanged
func GetUsersFieldDatasByLangcodeAndTimezoneAndChanged(offset int, limit int, Langcode_ string, Timezone_ string, Changed_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("langcode = ? and timezone = ? and changed = ?", Langcode_, Timezone_, Changed_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByLangcodeAndTimezoneAndAccess Get UsersFieldDatas via LangcodeAndTimezoneAndAccess
func GetUsersFieldDatasByLangcodeAndTimezoneAndAccess(offset int, limit int, Langcode_ string, Timezone_ string, Access_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("langcode = ? and timezone = ? and access = ?", Langcode_, Timezone_, Access_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByLangcodeAndTimezoneAndLogin Get UsersFieldDatas via LangcodeAndTimezoneAndLogin
func GetUsersFieldDatasByLangcodeAndTimezoneAndLogin(offset int, limit int, Langcode_ string, Timezone_ string, Login_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("langcode = ? and timezone = ? and login = ?", Langcode_, Timezone_, Login_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByLangcodeAndTimezoneAndInit Get UsersFieldDatas via LangcodeAndTimezoneAndInit
func GetUsersFieldDatasByLangcodeAndTimezoneAndInit(offset int, limit int, Langcode_ string, Timezone_ string, Init_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("langcode = ? and timezone = ? and init = ?", Langcode_, Timezone_, Init_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByLangcodeAndTimezoneAndDefaultLangcode Get UsersFieldDatas via LangcodeAndTimezoneAndDefaultLangcode
func GetUsersFieldDatasByLangcodeAndTimezoneAndDefaultLangcode(offset int, limit int, Langcode_ string, Timezone_ string, DefaultLangcode_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("langcode = ? and timezone = ? and default_langcode = ?", Langcode_, Timezone_, DefaultLangcode_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByLangcodeAndStatusAndCreated Get UsersFieldDatas via LangcodeAndStatusAndCreated
func GetUsersFieldDatasByLangcodeAndStatusAndCreated(offset int, limit int, Langcode_ string, Status_ int, Created_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("langcode = ? and status = ? and created = ?", Langcode_, Status_, Created_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByLangcodeAndStatusAndChanged Get UsersFieldDatas via LangcodeAndStatusAndChanged
func GetUsersFieldDatasByLangcodeAndStatusAndChanged(offset int, limit int, Langcode_ string, Status_ int, Changed_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("langcode = ? and status = ? and changed = ?", Langcode_, Status_, Changed_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByLangcodeAndStatusAndAccess Get UsersFieldDatas via LangcodeAndStatusAndAccess
func GetUsersFieldDatasByLangcodeAndStatusAndAccess(offset int, limit int, Langcode_ string, Status_ int, Access_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("langcode = ? and status = ? and access = ?", Langcode_, Status_, Access_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByLangcodeAndStatusAndLogin Get UsersFieldDatas via LangcodeAndStatusAndLogin
func GetUsersFieldDatasByLangcodeAndStatusAndLogin(offset int, limit int, Langcode_ string, Status_ int, Login_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("langcode = ? and status = ? and login = ?", Langcode_, Status_, Login_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByLangcodeAndStatusAndInit Get UsersFieldDatas via LangcodeAndStatusAndInit
func GetUsersFieldDatasByLangcodeAndStatusAndInit(offset int, limit int, Langcode_ string, Status_ int, Init_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("langcode = ? and status = ? and init = ?", Langcode_, Status_, Init_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByLangcodeAndStatusAndDefaultLangcode Get UsersFieldDatas via LangcodeAndStatusAndDefaultLangcode
func GetUsersFieldDatasByLangcodeAndStatusAndDefaultLangcode(offset int, limit int, Langcode_ string, Status_ int, DefaultLangcode_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("langcode = ? and status = ? and default_langcode = ?", Langcode_, Status_, DefaultLangcode_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByLangcodeAndCreatedAndChanged Get UsersFieldDatas via LangcodeAndCreatedAndChanged
func GetUsersFieldDatasByLangcodeAndCreatedAndChanged(offset int, limit int, Langcode_ string, Created_ int, Changed_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("langcode = ? and created = ? and changed = ?", Langcode_, Created_, Changed_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByLangcodeAndCreatedAndAccess Get UsersFieldDatas via LangcodeAndCreatedAndAccess
func GetUsersFieldDatasByLangcodeAndCreatedAndAccess(offset int, limit int, Langcode_ string, Created_ int, Access_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("langcode = ? and created = ? and access = ?", Langcode_, Created_, Access_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByLangcodeAndCreatedAndLogin Get UsersFieldDatas via LangcodeAndCreatedAndLogin
func GetUsersFieldDatasByLangcodeAndCreatedAndLogin(offset int, limit int, Langcode_ string, Created_ int, Login_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("langcode = ? and created = ? and login = ?", Langcode_, Created_, Login_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByLangcodeAndCreatedAndInit Get UsersFieldDatas via LangcodeAndCreatedAndInit
func GetUsersFieldDatasByLangcodeAndCreatedAndInit(offset int, limit int, Langcode_ string, Created_ int, Init_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("langcode = ? and created = ? and init = ?", Langcode_, Created_, Init_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByLangcodeAndCreatedAndDefaultLangcode Get UsersFieldDatas via LangcodeAndCreatedAndDefaultLangcode
func GetUsersFieldDatasByLangcodeAndCreatedAndDefaultLangcode(offset int, limit int, Langcode_ string, Created_ int, DefaultLangcode_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("langcode = ? and created = ? and default_langcode = ?", Langcode_, Created_, DefaultLangcode_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByLangcodeAndChangedAndAccess Get UsersFieldDatas via LangcodeAndChangedAndAccess
func GetUsersFieldDatasByLangcodeAndChangedAndAccess(offset int, limit int, Langcode_ string, Changed_ int, Access_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("langcode = ? and changed = ? and access = ?", Langcode_, Changed_, Access_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByLangcodeAndChangedAndLogin Get UsersFieldDatas via LangcodeAndChangedAndLogin
func GetUsersFieldDatasByLangcodeAndChangedAndLogin(offset int, limit int, Langcode_ string, Changed_ int, Login_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("langcode = ? and changed = ? and login = ?", Langcode_, Changed_, Login_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByLangcodeAndChangedAndInit Get UsersFieldDatas via LangcodeAndChangedAndInit
func GetUsersFieldDatasByLangcodeAndChangedAndInit(offset int, limit int, Langcode_ string, Changed_ int, Init_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("langcode = ? and changed = ? and init = ?", Langcode_, Changed_, Init_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByLangcodeAndChangedAndDefaultLangcode Get UsersFieldDatas via LangcodeAndChangedAndDefaultLangcode
func GetUsersFieldDatasByLangcodeAndChangedAndDefaultLangcode(offset int, limit int, Langcode_ string, Changed_ int, DefaultLangcode_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("langcode = ? and changed = ? and default_langcode = ?", Langcode_, Changed_, DefaultLangcode_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByLangcodeAndAccessAndLogin Get UsersFieldDatas via LangcodeAndAccessAndLogin
func GetUsersFieldDatasByLangcodeAndAccessAndLogin(offset int, limit int, Langcode_ string, Access_ int, Login_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("langcode = ? and access = ? and login = ?", Langcode_, Access_, Login_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByLangcodeAndAccessAndInit Get UsersFieldDatas via LangcodeAndAccessAndInit
func GetUsersFieldDatasByLangcodeAndAccessAndInit(offset int, limit int, Langcode_ string, Access_ int, Init_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("langcode = ? and access = ? and init = ?", Langcode_, Access_, Init_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByLangcodeAndAccessAndDefaultLangcode Get UsersFieldDatas via LangcodeAndAccessAndDefaultLangcode
func GetUsersFieldDatasByLangcodeAndAccessAndDefaultLangcode(offset int, limit int, Langcode_ string, Access_ int, DefaultLangcode_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("langcode = ? and access = ? and default_langcode = ?", Langcode_, Access_, DefaultLangcode_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByLangcodeAndLoginAndInit Get UsersFieldDatas via LangcodeAndLoginAndInit
func GetUsersFieldDatasByLangcodeAndLoginAndInit(offset int, limit int, Langcode_ string, Login_ int, Init_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("langcode = ? and login = ? and init = ?", Langcode_, Login_, Init_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByLangcodeAndLoginAndDefaultLangcode Get UsersFieldDatas via LangcodeAndLoginAndDefaultLangcode
func GetUsersFieldDatasByLangcodeAndLoginAndDefaultLangcode(offset int, limit int, Langcode_ string, Login_ int, DefaultLangcode_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("langcode = ? and login = ? and default_langcode = ?", Langcode_, Login_, DefaultLangcode_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByLangcodeAndInitAndDefaultLangcode Get UsersFieldDatas via LangcodeAndInitAndDefaultLangcode
func GetUsersFieldDatasByLangcodeAndInitAndDefaultLangcode(offset int, limit int, Langcode_ string, Init_ string, DefaultLangcode_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("langcode = ? and init = ? and default_langcode = ?", Langcode_, Init_, DefaultLangcode_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredLangcodeAndPreferredAdminLangcodeAndName Get UsersFieldDatas via PreferredLangcodeAndPreferredAdminLangcodeAndName
func GetUsersFieldDatasByPreferredLangcodeAndPreferredAdminLangcodeAndName(offset int, limit int, PreferredLangcode_ string, PreferredAdminLangcode_ string, Name_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_langcode = ? and preferred_admin_langcode = ? and name = ?", PreferredLangcode_, PreferredAdminLangcode_, Name_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredLangcodeAndPreferredAdminLangcodeAndPass Get UsersFieldDatas via PreferredLangcodeAndPreferredAdminLangcodeAndPass
func GetUsersFieldDatasByPreferredLangcodeAndPreferredAdminLangcodeAndPass(offset int, limit int, PreferredLangcode_ string, PreferredAdminLangcode_ string, Pass_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_langcode = ? and preferred_admin_langcode = ? and pass = ?", PreferredLangcode_, PreferredAdminLangcode_, Pass_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredLangcodeAndPreferredAdminLangcodeAndMail Get UsersFieldDatas via PreferredLangcodeAndPreferredAdminLangcodeAndMail
func GetUsersFieldDatasByPreferredLangcodeAndPreferredAdminLangcodeAndMail(offset int, limit int, PreferredLangcode_ string, PreferredAdminLangcode_ string, Mail_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_langcode = ? and preferred_admin_langcode = ? and mail = ?", PreferredLangcode_, PreferredAdminLangcode_, Mail_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredLangcodeAndPreferredAdminLangcodeAndTimezone Get UsersFieldDatas via PreferredLangcodeAndPreferredAdminLangcodeAndTimezone
func GetUsersFieldDatasByPreferredLangcodeAndPreferredAdminLangcodeAndTimezone(offset int, limit int, PreferredLangcode_ string, PreferredAdminLangcode_ string, Timezone_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_langcode = ? and preferred_admin_langcode = ? and timezone = ?", PreferredLangcode_, PreferredAdminLangcode_, Timezone_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredLangcodeAndPreferredAdminLangcodeAndStatus Get UsersFieldDatas via PreferredLangcodeAndPreferredAdminLangcodeAndStatus
func GetUsersFieldDatasByPreferredLangcodeAndPreferredAdminLangcodeAndStatus(offset int, limit int, PreferredLangcode_ string, PreferredAdminLangcode_ string, Status_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_langcode = ? and preferred_admin_langcode = ? and status = ?", PreferredLangcode_, PreferredAdminLangcode_, Status_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredLangcodeAndPreferredAdminLangcodeAndCreated Get UsersFieldDatas via PreferredLangcodeAndPreferredAdminLangcodeAndCreated
func GetUsersFieldDatasByPreferredLangcodeAndPreferredAdminLangcodeAndCreated(offset int, limit int, PreferredLangcode_ string, PreferredAdminLangcode_ string, Created_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_langcode = ? and preferred_admin_langcode = ? and created = ?", PreferredLangcode_, PreferredAdminLangcode_, Created_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredLangcodeAndPreferredAdminLangcodeAndChanged Get UsersFieldDatas via PreferredLangcodeAndPreferredAdminLangcodeAndChanged
func GetUsersFieldDatasByPreferredLangcodeAndPreferredAdminLangcodeAndChanged(offset int, limit int, PreferredLangcode_ string, PreferredAdminLangcode_ string, Changed_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_langcode = ? and preferred_admin_langcode = ? and changed = ?", PreferredLangcode_, PreferredAdminLangcode_, Changed_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredLangcodeAndPreferredAdminLangcodeAndAccess Get UsersFieldDatas via PreferredLangcodeAndPreferredAdminLangcodeAndAccess
func GetUsersFieldDatasByPreferredLangcodeAndPreferredAdminLangcodeAndAccess(offset int, limit int, PreferredLangcode_ string, PreferredAdminLangcode_ string, Access_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_langcode = ? and preferred_admin_langcode = ? and access = ?", PreferredLangcode_, PreferredAdminLangcode_, Access_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredLangcodeAndPreferredAdminLangcodeAndLogin Get UsersFieldDatas via PreferredLangcodeAndPreferredAdminLangcodeAndLogin
func GetUsersFieldDatasByPreferredLangcodeAndPreferredAdminLangcodeAndLogin(offset int, limit int, PreferredLangcode_ string, PreferredAdminLangcode_ string, Login_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_langcode = ? and preferred_admin_langcode = ? and login = ?", PreferredLangcode_, PreferredAdminLangcode_, Login_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredLangcodeAndPreferredAdminLangcodeAndInit Get UsersFieldDatas via PreferredLangcodeAndPreferredAdminLangcodeAndInit
func GetUsersFieldDatasByPreferredLangcodeAndPreferredAdminLangcodeAndInit(offset int, limit int, PreferredLangcode_ string, PreferredAdminLangcode_ string, Init_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_langcode = ? and preferred_admin_langcode = ? and init = ?", PreferredLangcode_, PreferredAdminLangcode_, Init_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredLangcodeAndPreferredAdminLangcodeAndDefaultLangcode Get UsersFieldDatas via PreferredLangcodeAndPreferredAdminLangcodeAndDefaultLangcode
func GetUsersFieldDatasByPreferredLangcodeAndPreferredAdminLangcodeAndDefaultLangcode(offset int, limit int, PreferredLangcode_ string, PreferredAdminLangcode_ string, DefaultLangcode_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_langcode = ? and preferred_admin_langcode = ? and default_langcode = ?", PreferredLangcode_, PreferredAdminLangcode_, DefaultLangcode_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredLangcodeAndNameAndPass Get UsersFieldDatas via PreferredLangcodeAndNameAndPass
func GetUsersFieldDatasByPreferredLangcodeAndNameAndPass(offset int, limit int, PreferredLangcode_ string, Name_ string, Pass_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_langcode = ? and name = ? and pass = ?", PreferredLangcode_, Name_, Pass_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredLangcodeAndNameAndMail Get UsersFieldDatas via PreferredLangcodeAndNameAndMail
func GetUsersFieldDatasByPreferredLangcodeAndNameAndMail(offset int, limit int, PreferredLangcode_ string, Name_ string, Mail_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_langcode = ? and name = ? and mail = ?", PreferredLangcode_, Name_, Mail_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredLangcodeAndNameAndTimezone Get UsersFieldDatas via PreferredLangcodeAndNameAndTimezone
func GetUsersFieldDatasByPreferredLangcodeAndNameAndTimezone(offset int, limit int, PreferredLangcode_ string, Name_ string, Timezone_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_langcode = ? and name = ? and timezone = ?", PreferredLangcode_, Name_, Timezone_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredLangcodeAndNameAndStatus Get UsersFieldDatas via PreferredLangcodeAndNameAndStatus
func GetUsersFieldDatasByPreferredLangcodeAndNameAndStatus(offset int, limit int, PreferredLangcode_ string, Name_ string, Status_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_langcode = ? and name = ? and status = ?", PreferredLangcode_, Name_, Status_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredLangcodeAndNameAndCreated Get UsersFieldDatas via PreferredLangcodeAndNameAndCreated
func GetUsersFieldDatasByPreferredLangcodeAndNameAndCreated(offset int, limit int, PreferredLangcode_ string, Name_ string, Created_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_langcode = ? and name = ? and created = ?", PreferredLangcode_, Name_, Created_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredLangcodeAndNameAndChanged Get UsersFieldDatas via PreferredLangcodeAndNameAndChanged
func GetUsersFieldDatasByPreferredLangcodeAndNameAndChanged(offset int, limit int, PreferredLangcode_ string, Name_ string, Changed_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_langcode = ? and name = ? and changed = ?", PreferredLangcode_, Name_, Changed_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredLangcodeAndNameAndAccess Get UsersFieldDatas via PreferredLangcodeAndNameAndAccess
func GetUsersFieldDatasByPreferredLangcodeAndNameAndAccess(offset int, limit int, PreferredLangcode_ string, Name_ string, Access_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_langcode = ? and name = ? and access = ?", PreferredLangcode_, Name_, Access_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredLangcodeAndNameAndLogin Get UsersFieldDatas via PreferredLangcodeAndNameAndLogin
func GetUsersFieldDatasByPreferredLangcodeAndNameAndLogin(offset int, limit int, PreferredLangcode_ string, Name_ string, Login_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_langcode = ? and name = ? and login = ?", PreferredLangcode_, Name_, Login_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredLangcodeAndNameAndInit Get UsersFieldDatas via PreferredLangcodeAndNameAndInit
func GetUsersFieldDatasByPreferredLangcodeAndNameAndInit(offset int, limit int, PreferredLangcode_ string, Name_ string, Init_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_langcode = ? and name = ? and init = ?", PreferredLangcode_, Name_, Init_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredLangcodeAndNameAndDefaultLangcode Get UsersFieldDatas via PreferredLangcodeAndNameAndDefaultLangcode
func GetUsersFieldDatasByPreferredLangcodeAndNameAndDefaultLangcode(offset int, limit int, PreferredLangcode_ string, Name_ string, DefaultLangcode_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_langcode = ? and name = ? and default_langcode = ?", PreferredLangcode_, Name_, DefaultLangcode_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredLangcodeAndPassAndMail Get UsersFieldDatas via PreferredLangcodeAndPassAndMail
func GetUsersFieldDatasByPreferredLangcodeAndPassAndMail(offset int, limit int, PreferredLangcode_ string, Pass_ string, Mail_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_langcode = ? and pass = ? and mail = ?", PreferredLangcode_, Pass_, Mail_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredLangcodeAndPassAndTimezone Get UsersFieldDatas via PreferredLangcodeAndPassAndTimezone
func GetUsersFieldDatasByPreferredLangcodeAndPassAndTimezone(offset int, limit int, PreferredLangcode_ string, Pass_ string, Timezone_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_langcode = ? and pass = ? and timezone = ?", PreferredLangcode_, Pass_, Timezone_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredLangcodeAndPassAndStatus Get UsersFieldDatas via PreferredLangcodeAndPassAndStatus
func GetUsersFieldDatasByPreferredLangcodeAndPassAndStatus(offset int, limit int, PreferredLangcode_ string, Pass_ string, Status_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_langcode = ? and pass = ? and status = ?", PreferredLangcode_, Pass_, Status_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredLangcodeAndPassAndCreated Get UsersFieldDatas via PreferredLangcodeAndPassAndCreated
func GetUsersFieldDatasByPreferredLangcodeAndPassAndCreated(offset int, limit int, PreferredLangcode_ string, Pass_ string, Created_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_langcode = ? and pass = ? and created = ?", PreferredLangcode_, Pass_, Created_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredLangcodeAndPassAndChanged Get UsersFieldDatas via PreferredLangcodeAndPassAndChanged
func GetUsersFieldDatasByPreferredLangcodeAndPassAndChanged(offset int, limit int, PreferredLangcode_ string, Pass_ string, Changed_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_langcode = ? and pass = ? and changed = ?", PreferredLangcode_, Pass_, Changed_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredLangcodeAndPassAndAccess Get UsersFieldDatas via PreferredLangcodeAndPassAndAccess
func GetUsersFieldDatasByPreferredLangcodeAndPassAndAccess(offset int, limit int, PreferredLangcode_ string, Pass_ string, Access_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_langcode = ? and pass = ? and access = ?", PreferredLangcode_, Pass_, Access_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredLangcodeAndPassAndLogin Get UsersFieldDatas via PreferredLangcodeAndPassAndLogin
func GetUsersFieldDatasByPreferredLangcodeAndPassAndLogin(offset int, limit int, PreferredLangcode_ string, Pass_ string, Login_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_langcode = ? and pass = ? and login = ?", PreferredLangcode_, Pass_, Login_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredLangcodeAndPassAndInit Get UsersFieldDatas via PreferredLangcodeAndPassAndInit
func GetUsersFieldDatasByPreferredLangcodeAndPassAndInit(offset int, limit int, PreferredLangcode_ string, Pass_ string, Init_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_langcode = ? and pass = ? and init = ?", PreferredLangcode_, Pass_, Init_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredLangcodeAndPassAndDefaultLangcode Get UsersFieldDatas via PreferredLangcodeAndPassAndDefaultLangcode
func GetUsersFieldDatasByPreferredLangcodeAndPassAndDefaultLangcode(offset int, limit int, PreferredLangcode_ string, Pass_ string, DefaultLangcode_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_langcode = ? and pass = ? and default_langcode = ?", PreferredLangcode_, Pass_, DefaultLangcode_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredLangcodeAndMailAndTimezone Get UsersFieldDatas via PreferredLangcodeAndMailAndTimezone
func GetUsersFieldDatasByPreferredLangcodeAndMailAndTimezone(offset int, limit int, PreferredLangcode_ string, Mail_ string, Timezone_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_langcode = ? and mail = ? and timezone = ?", PreferredLangcode_, Mail_, Timezone_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredLangcodeAndMailAndStatus Get UsersFieldDatas via PreferredLangcodeAndMailAndStatus
func GetUsersFieldDatasByPreferredLangcodeAndMailAndStatus(offset int, limit int, PreferredLangcode_ string, Mail_ string, Status_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_langcode = ? and mail = ? and status = ?", PreferredLangcode_, Mail_, Status_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredLangcodeAndMailAndCreated Get UsersFieldDatas via PreferredLangcodeAndMailAndCreated
func GetUsersFieldDatasByPreferredLangcodeAndMailAndCreated(offset int, limit int, PreferredLangcode_ string, Mail_ string, Created_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_langcode = ? and mail = ? and created = ?", PreferredLangcode_, Mail_, Created_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredLangcodeAndMailAndChanged Get UsersFieldDatas via PreferredLangcodeAndMailAndChanged
func GetUsersFieldDatasByPreferredLangcodeAndMailAndChanged(offset int, limit int, PreferredLangcode_ string, Mail_ string, Changed_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_langcode = ? and mail = ? and changed = ?", PreferredLangcode_, Mail_, Changed_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredLangcodeAndMailAndAccess Get UsersFieldDatas via PreferredLangcodeAndMailAndAccess
func GetUsersFieldDatasByPreferredLangcodeAndMailAndAccess(offset int, limit int, PreferredLangcode_ string, Mail_ string, Access_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_langcode = ? and mail = ? and access = ?", PreferredLangcode_, Mail_, Access_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredLangcodeAndMailAndLogin Get UsersFieldDatas via PreferredLangcodeAndMailAndLogin
func GetUsersFieldDatasByPreferredLangcodeAndMailAndLogin(offset int, limit int, PreferredLangcode_ string, Mail_ string, Login_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_langcode = ? and mail = ? and login = ?", PreferredLangcode_, Mail_, Login_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredLangcodeAndMailAndInit Get UsersFieldDatas via PreferredLangcodeAndMailAndInit
func GetUsersFieldDatasByPreferredLangcodeAndMailAndInit(offset int, limit int, PreferredLangcode_ string, Mail_ string, Init_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_langcode = ? and mail = ? and init = ?", PreferredLangcode_, Mail_, Init_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredLangcodeAndMailAndDefaultLangcode Get UsersFieldDatas via PreferredLangcodeAndMailAndDefaultLangcode
func GetUsersFieldDatasByPreferredLangcodeAndMailAndDefaultLangcode(offset int, limit int, PreferredLangcode_ string, Mail_ string, DefaultLangcode_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_langcode = ? and mail = ? and default_langcode = ?", PreferredLangcode_, Mail_, DefaultLangcode_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredLangcodeAndTimezoneAndStatus Get UsersFieldDatas via PreferredLangcodeAndTimezoneAndStatus
func GetUsersFieldDatasByPreferredLangcodeAndTimezoneAndStatus(offset int, limit int, PreferredLangcode_ string, Timezone_ string, Status_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_langcode = ? and timezone = ? and status = ?", PreferredLangcode_, Timezone_, Status_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredLangcodeAndTimezoneAndCreated Get UsersFieldDatas via PreferredLangcodeAndTimezoneAndCreated
func GetUsersFieldDatasByPreferredLangcodeAndTimezoneAndCreated(offset int, limit int, PreferredLangcode_ string, Timezone_ string, Created_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_langcode = ? and timezone = ? and created = ?", PreferredLangcode_, Timezone_, Created_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredLangcodeAndTimezoneAndChanged Get UsersFieldDatas via PreferredLangcodeAndTimezoneAndChanged
func GetUsersFieldDatasByPreferredLangcodeAndTimezoneAndChanged(offset int, limit int, PreferredLangcode_ string, Timezone_ string, Changed_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_langcode = ? and timezone = ? and changed = ?", PreferredLangcode_, Timezone_, Changed_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredLangcodeAndTimezoneAndAccess Get UsersFieldDatas via PreferredLangcodeAndTimezoneAndAccess
func GetUsersFieldDatasByPreferredLangcodeAndTimezoneAndAccess(offset int, limit int, PreferredLangcode_ string, Timezone_ string, Access_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_langcode = ? and timezone = ? and access = ?", PreferredLangcode_, Timezone_, Access_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredLangcodeAndTimezoneAndLogin Get UsersFieldDatas via PreferredLangcodeAndTimezoneAndLogin
func GetUsersFieldDatasByPreferredLangcodeAndTimezoneAndLogin(offset int, limit int, PreferredLangcode_ string, Timezone_ string, Login_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_langcode = ? and timezone = ? and login = ?", PreferredLangcode_, Timezone_, Login_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredLangcodeAndTimezoneAndInit Get UsersFieldDatas via PreferredLangcodeAndTimezoneAndInit
func GetUsersFieldDatasByPreferredLangcodeAndTimezoneAndInit(offset int, limit int, PreferredLangcode_ string, Timezone_ string, Init_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_langcode = ? and timezone = ? and init = ?", PreferredLangcode_, Timezone_, Init_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredLangcodeAndTimezoneAndDefaultLangcode Get UsersFieldDatas via PreferredLangcodeAndTimezoneAndDefaultLangcode
func GetUsersFieldDatasByPreferredLangcodeAndTimezoneAndDefaultLangcode(offset int, limit int, PreferredLangcode_ string, Timezone_ string, DefaultLangcode_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_langcode = ? and timezone = ? and default_langcode = ?", PreferredLangcode_, Timezone_, DefaultLangcode_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredLangcodeAndStatusAndCreated Get UsersFieldDatas via PreferredLangcodeAndStatusAndCreated
func GetUsersFieldDatasByPreferredLangcodeAndStatusAndCreated(offset int, limit int, PreferredLangcode_ string, Status_ int, Created_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_langcode = ? and status = ? and created = ?", PreferredLangcode_, Status_, Created_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredLangcodeAndStatusAndChanged Get UsersFieldDatas via PreferredLangcodeAndStatusAndChanged
func GetUsersFieldDatasByPreferredLangcodeAndStatusAndChanged(offset int, limit int, PreferredLangcode_ string, Status_ int, Changed_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_langcode = ? and status = ? and changed = ?", PreferredLangcode_, Status_, Changed_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredLangcodeAndStatusAndAccess Get UsersFieldDatas via PreferredLangcodeAndStatusAndAccess
func GetUsersFieldDatasByPreferredLangcodeAndStatusAndAccess(offset int, limit int, PreferredLangcode_ string, Status_ int, Access_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_langcode = ? and status = ? and access = ?", PreferredLangcode_, Status_, Access_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredLangcodeAndStatusAndLogin Get UsersFieldDatas via PreferredLangcodeAndStatusAndLogin
func GetUsersFieldDatasByPreferredLangcodeAndStatusAndLogin(offset int, limit int, PreferredLangcode_ string, Status_ int, Login_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_langcode = ? and status = ? and login = ?", PreferredLangcode_, Status_, Login_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredLangcodeAndStatusAndInit Get UsersFieldDatas via PreferredLangcodeAndStatusAndInit
func GetUsersFieldDatasByPreferredLangcodeAndStatusAndInit(offset int, limit int, PreferredLangcode_ string, Status_ int, Init_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_langcode = ? and status = ? and init = ?", PreferredLangcode_, Status_, Init_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredLangcodeAndStatusAndDefaultLangcode Get UsersFieldDatas via PreferredLangcodeAndStatusAndDefaultLangcode
func GetUsersFieldDatasByPreferredLangcodeAndStatusAndDefaultLangcode(offset int, limit int, PreferredLangcode_ string, Status_ int, DefaultLangcode_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_langcode = ? and status = ? and default_langcode = ?", PreferredLangcode_, Status_, DefaultLangcode_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredLangcodeAndCreatedAndChanged Get UsersFieldDatas via PreferredLangcodeAndCreatedAndChanged
func GetUsersFieldDatasByPreferredLangcodeAndCreatedAndChanged(offset int, limit int, PreferredLangcode_ string, Created_ int, Changed_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_langcode = ? and created = ? and changed = ?", PreferredLangcode_, Created_, Changed_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredLangcodeAndCreatedAndAccess Get UsersFieldDatas via PreferredLangcodeAndCreatedAndAccess
func GetUsersFieldDatasByPreferredLangcodeAndCreatedAndAccess(offset int, limit int, PreferredLangcode_ string, Created_ int, Access_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_langcode = ? and created = ? and access = ?", PreferredLangcode_, Created_, Access_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredLangcodeAndCreatedAndLogin Get UsersFieldDatas via PreferredLangcodeAndCreatedAndLogin
func GetUsersFieldDatasByPreferredLangcodeAndCreatedAndLogin(offset int, limit int, PreferredLangcode_ string, Created_ int, Login_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_langcode = ? and created = ? and login = ?", PreferredLangcode_, Created_, Login_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredLangcodeAndCreatedAndInit Get UsersFieldDatas via PreferredLangcodeAndCreatedAndInit
func GetUsersFieldDatasByPreferredLangcodeAndCreatedAndInit(offset int, limit int, PreferredLangcode_ string, Created_ int, Init_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_langcode = ? and created = ? and init = ?", PreferredLangcode_, Created_, Init_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredLangcodeAndCreatedAndDefaultLangcode Get UsersFieldDatas via PreferredLangcodeAndCreatedAndDefaultLangcode
func GetUsersFieldDatasByPreferredLangcodeAndCreatedAndDefaultLangcode(offset int, limit int, PreferredLangcode_ string, Created_ int, DefaultLangcode_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_langcode = ? and created = ? and default_langcode = ?", PreferredLangcode_, Created_, DefaultLangcode_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredLangcodeAndChangedAndAccess Get UsersFieldDatas via PreferredLangcodeAndChangedAndAccess
func GetUsersFieldDatasByPreferredLangcodeAndChangedAndAccess(offset int, limit int, PreferredLangcode_ string, Changed_ int, Access_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_langcode = ? and changed = ? and access = ?", PreferredLangcode_, Changed_, Access_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredLangcodeAndChangedAndLogin Get UsersFieldDatas via PreferredLangcodeAndChangedAndLogin
func GetUsersFieldDatasByPreferredLangcodeAndChangedAndLogin(offset int, limit int, PreferredLangcode_ string, Changed_ int, Login_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_langcode = ? and changed = ? and login = ?", PreferredLangcode_, Changed_, Login_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredLangcodeAndChangedAndInit Get UsersFieldDatas via PreferredLangcodeAndChangedAndInit
func GetUsersFieldDatasByPreferredLangcodeAndChangedAndInit(offset int, limit int, PreferredLangcode_ string, Changed_ int, Init_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_langcode = ? and changed = ? and init = ?", PreferredLangcode_, Changed_, Init_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredLangcodeAndChangedAndDefaultLangcode Get UsersFieldDatas via PreferredLangcodeAndChangedAndDefaultLangcode
func GetUsersFieldDatasByPreferredLangcodeAndChangedAndDefaultLangcode(offset int, limit int, PreferredLangcode_ string, Changed_ int, DefaultLangcode_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_langcode = ? and changed = ? and default_langcode = ?", PreferredLangcode_, Changed_, DefaultLangcode_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredLangcodeAndAccessAndLogin Get UsersFieldDatas via PreferredLangcodeAndAccessAndLogin
func GetUsersFieldDatasByPreferredLangcodeAndAccessAndLogin(offset int, limit int, PreferredLangcode_ string, Access_ int, Login_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_langcode = ? and access = ? and login = ?", PreferredLangcode_, Access_, Login_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredLangcodeAndAccessAndInit Get UsersFieldDatas via PreferredLangcodeAndAccessAndInit
func GetUsersFieldDatasByPreferredLangcodeAndAccessAndInit(offset int, limit int, PreferredLangcode_ string, Access_ int, Init_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_langcode = ? and access = ? and init = ?", PreferredLangcode_, Access_, Init_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredLangcodeAndAccessAndDefaultLangcode Get UsersFieldDatas via PreferredLangcodeAndAccessAndDefaultLangcode
func GetUsersFieldDatasByPreferredLangcodeAndAccessAndDefaultLangcode(offset int, limit int, PreferredLangcode_ string, Access_ int, DefaultLangcode_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_langcode = ? and access = ? and default_langcode = ?", PreferredLangcode_, Access_, DefaultLangcode_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredLangcodeAndLoginAndInit Get UsersFieldDatas via PreferredLangcodeAndLoginAndInit
func GetUsersFieldDatasByPreferredLangcodeAndLoginAndInit(offset int, limit int, PreferredLangcode_ string, Login_ int, Init_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_langcode = ? and login = ? and init = ?", PreferredLangcode_, Login_, Init_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredLangcodeAndLoginAndDefaultLangcode Get UsersFieldDatas via PreferredLangcodeAndLoginAndDefaultLangcode
func GetUsersFieldDatasByPreferredLangcodeAndLoginAndDefaultLangcode(offset int, limit int, PreferredLangcode_ string, Login_ int, DefaultLangcode_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_langcode = ? and login = ? and default_langcode = ?", PreferredLangcode_, Login_, DefaultLangcode_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredLangcodeAndInitAndDefaultLangcode Get UsersFieldDatas via PreferredLangcodeAndInitAndDefaultLangcode
func GetUsersFieldDatasByPreferredLangcodeAndInitAndDefaultLangcode(offset int, limit int, PreferredLangcode_ string, Init_ string, DefaultLangcode_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_langcode = ? and init = ? and default_langcode = ?", PreferredLangcode_, Init_, DefaultLangcode_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredAdminLangcodeAndNameAndPass Get UsersFieldDatas via PreferredAdminLangcodeAndNameAndPass
func GetUsersFieldDatasByPreferredAdminLangcodeAndNameAndPass(offset int, limit int, PreferredAdminLangcode_ string, Name_ string, Pass_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_admin_langcode = ? and name = ? and pass = ?", PreferredAdminLangcode_, Name_, Pass_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredAdminLangcodeAndNameAndMail Get UsersFieldDatas via PreferredAdminLangcodeAndNameAndMail
func GetUsersFieldDatasByPreferredAdminLangcodeAndNameAndMail(offset int, limit int, PreferredAdminLangcode_ string, Name_ string, Mail_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_admin_langcode = ? and name = ? and mail = ?", PreferredAdminLangcode_, Name_, Mail_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredAdminLangcodeAndNameAndTimezone Get UsersFieldDatas via PreferredAdminLangcodeAndNameAndTimezone
func GetUsersFieldDatasByPreferredAdminLangcodeAndNameAndTimezone(offset int, limit int, PreferredAdminLangcode_ string, Name_ string, Timezone_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_admin_langcode = ? and name = ? and timezone = ?", PreferredAdminLangcode_, Name_, Timezone_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredAdminLangcodeAndNameAndStatus Get UsersFieldDatas via PreferredAdminLangcodeAndNameAndStatus
func GetUsersFieldDatasByPreferredAdminLangcodeAndNameAndStatus(offset int, limit int, PreferredAdminLangcode_ string, Name_ string, Status_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_admin_langcode = ? and name = ? and status = ?", PreferredAdminLangcode_, Name_, Status_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredAdminLangcodeAndNameAndCreated Get UsersFieldDatas via PreferredAdminLangcodeAndNameAndCreated
func GetUsersFieldDatasByPreferredAdminLangcodeAndNameAndCreated(offset int, limit int, PreferredAdminLangcode_ string, Name_ string, Created_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_admin_langcode = ? and name = ? and created = ?", PreferredAdminLangcode_, Name_, Created_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredAdminLangcodeAndNameAndChanged Get UsersFieldDatas via PreferredAdminLangcodeAndNameAndChanged
func GetUsersFieldDatasByPreferredAdminLangcodeAndNameAndChanged(offset int, limit int, PreferredAdminLangcode_ string, Name_ string, Changed_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_admin_langcode = ? and name = ? and changed = ?", PreferredAdminLangcode_, Name_, Changed_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredAdminLangcodeAndNameAndAccess Get UsersFieldDatas via PreferredAdminLangcodeAndNameAndAccess
func GetUsersFieldDatasByPreferredAdminLangcodeAndNameAndAccess(offset int, limit int, PreferredAdminLangcode_ string, Name_ string, Access_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_admin_langcode = ? and name = ? and access = ?", PreferredAdminLangcode_, Name_, Access_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredAdminLangcodeAndNameAndLogin Get UsersFieldDatas via PreferredAdminLangcodeAndNameAndLogin
func GetUsersFieldDatasByPreferredAdminLangcodeAndNameAndLogin(offset int, limit int, PreferredAdminLangcode_ string, Name_ string, Login_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_admin_langcode = ? and name = ? and login = ?", PreferredAdminLangcode_, Name_, Login_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredAdminLangcodeAndNameAndInit Get UsersFieldDatas via PreferredAdminLangcodeAndNameAndInit
func GetUsersFieldDatasByPreferredAdminLangcodeAndNameAndInit(offset int, limit int, PreferredAdminLangcode_ string, Name_ string, Init_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_admin_langcode = ? and name = ? and init = ?", PreferredAdminLangcode_, Name_, Init_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredAdminLangcodeAndNameAndDefaultLangcode Get UsersFieldDatas via PreferredAdminLangcodeAndNameAndDefaultLangcode
func GetUsersFieldDatasByPreferredAdminLangcodeAndNameAndDefaultLangcode(offset int, limit int, PreferredAdminLangcode_ string, Name_ string, DefaultLangcode_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_admin_langcode = ? and name = ? and default_langcode = ?", PreferredAdminLangcode_, Name_, DefaultLangcode_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredAdminLangcodeAndPassAndMail Get UsersFieldDatas via PreferredAdminLangcodeAndPassAndMail
func GetUsersFieldDatasByPreferredAdminLangcodeAndPassAndMail(offset int, limit int, PreferredAdminLangcode_ string, Pass_ string, Mail_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_admin_langcode = ? and pass = ? and mail = ?", PreferredAdminLangcode_, Pass_, Mail_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredAdminLangcodeAndPassAndTimezone Get UsersFieldDatas via PreferredAdminLangcodeAndPassAndTimezone
func GetUsersFieldDatasByPreferredAdminLangcodeAndPassAndTimezone(offset int, limit int, PreferredAdminLangcode_ string, Pass_ string, Timezone_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_admin_langcode = ? and pass = ? and timezone = ?", PreferredAdminLangcode_, Pass_, Timezone_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredAdminLangcodeAndPassAndStatus Get UsersFieldDatas via PreferredAdminLangcodeAndPassAndStatus
func GetUsersFieldDatasByPreferredAdminLangcodeAndPassAndStatus(offset int, limit int, PreferredAdminLangcode_ string, Pass_ string, Status_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_admin_langcode = ? and pass = ? and status = ?", PreferredAdminLangcode_, Pass_, Status_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredAdminLangcodeAndPassAndCreated Get UsersFieldDatas via PreferredAdminLangcodeAndPassAndCreated
func GetUsersFieldDatasByPreferredAdminLangcodeAndPassAndCreated(offset int, limit int, PreferredAdminLangcode_ string, Pass_ string, Created_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_admin_langcode = ? and pass = ? and created = ?", PreferredAdminLangcode_, Pass_, Created_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredAdminLangcodeAndPassAndChanged Get UsersFieldDatas via PreferredAdminLangcodeAndPassAndChanged
func GetUsersFieldDatasByPreferredAdminLangcodeAndPassAndChanged(offset int, limit int, PreferredAdminLangcode_ string, Pass_ string, Changed_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_admin_langcode = ? and pass = ? and changed = ?", PreferredAdminLangcode_, Pass_, Changed_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredAdminLangcodeAndPassAndAccess Get UsersFieldDatas via PreferredAdminLangcodeAndPassAndAccess
func GetUsersFieldDatasByPreferredAdminLangcodeAndPassAndAccess(offset int, limit int, PreferredAdminLangcode_ string, Pass_ string, Access_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_admin_langcode = ? and pass = ? and access = ?", PreferredAdminLangcode_, Pass_, Access_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredAdminLangcodeAndPassAndLogin Get UsersFieldDatas via PreferredAdminLangcodeAndPassAndLogin
func GetUsersFieldDatasByPreferredAdminLangcodeAndPassAndLogin(offset int, limit int, PreferredAdminLangcode_ string, Pass_ string, Login_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_admin_langcode = ? and pass = ? and login = ?", PreferredAdminLangcode_, Pass_, Login_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredAdminLangcodeAndPassAndInit Get UsersFieldDatas via PreferredAdminLangcodeAndPassAndInit
func GetUsersFieldDatasByPreferredAdminLangcodeAndPassAndInit(offset int, limit int, PreferredAdminLangcode_ string, Pass_ string, Init_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_admin_langcode = ? and pass = ? and init = ?", PreferredAdminLangcode_, Pass_, Init_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredAdminLangcodeAndPassAndDefaultLangcode Get UsersFieldDatas via PreferredAdminLangcodeAndPassAndDefaultLangcode
func GetUsersFieldDatasByPreferredAdminLangcodeAndPassAndDefaultLangcode(offset int, limit int, PreferredAdminLangcode_ string, Pass_ string, DefaultLangcode_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_admin_langcode = ? and pass = ? and default_langcode = ?", PreferredAdminLangcode_, Pass_, DefaultLangcode_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredAdminLangcodeAndMailAndTimezone Get UsersFieldDatas via PreferredAdminLangcodeAndMailAndTimezone
func GetUsersFieldDatasByPreferredAdminLangcodeAndMailAndTimezone(offset int, limit int, PreferredAdminLangcode_ string, Mail_ string, Timezone_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_admin_langcode = ? and mail = ? and timezone = ?", PreferredAdminLangcode_, Mail_, Timezone_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredAdminLangcodeAndMailAndStatus Get UsersFieldDatas via PreferredAdminLangcodeAndMailAndStatus
func GetUsersFieldDatasByPreferredAdminLangcodeAndMailAndStatus(offset int, limit int, PreferredAdminLangcode_ string, Mail_ string, Status_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_admin_langcode = ? and mail = ? and status = ?", PreferredAdminLangcode_, Mail_, Status_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredAdminLangcodeAndMailAndCreated Get UsersFieldDatas via PreferredAdminLangcodeAndMailAndCreated
func GetUsersFieldDatasByPreferredAdminLangcodeAndMailAndCreated(offset int, limit int, PreferredAdminLangcode_ string, Mail_ string, Created_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_admin_langcode = ? and mail = ? and created = ?", PreferredAdminLangcode_, Mail_, Created_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredAdminLangcodeAndMailAndChanged Get UsersFieldDatas via PreferredAdminLangcodeAndMailAndChanged
func GetUsersFieldDatasByPreferredAdminLangcodeAndMailAndChanged(offset int, limit int, PreferredAdminLangcode_ string, Mail_ string, Changed_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_admin_langcode = ? and mail = ? and changed = ?", PreferredAdminLangcode_, Mail_, Changed_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredAdminLangcodeAndMailAndAccess Get UsersFieldDatas via PreferredAdminLangcodeAndMailAndAccess
func GetUsersFieldDatasByPreferredAdminLangcodeAndMailAndAccess(offset int, limit int, PreferredAdminLangcode_ string, Mail_ string, Access_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_admin_langcode = ? and mail = ? and access = ?", PreferredAdminLangcode_, Mail_, Access_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredAdminLangcodeAndMailAndLogin Get UsersFieldDatas via PreferredAdminLangcodeAndMailAndLogin
func GetUsersFieldDatasByPreferredAdminLangcodeAndMailAndLogin(offset int, limit int, PreferredAdminLangcode_ string, Mail_ string, Login_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_admin_langcode = ? and mail = ? and login = ?", PreferredAdminLangcode_, Mail_, Login_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredAdminLangcodeAndMailAndInit Get UsersFieldDatas via PreferredAdminLangcodeAndMailAndInit
func GetUsersFieldDatasByPreferredAdminLangcodeAndMailAndInit(offset int, limit int, PreferredAdminLangcode_ string, Mail_ string, Init_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_admin_langcode = ? and mail = ? and init = ?", PreferredAdminLangcode_, Mail_, Init_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredAdminLangcodeAndMailAndDefaultLangcode Get UsersFieldDatas via PreferredAdminLangcodeAndMailAndDefaultLangcode
func GetUsersFieldDatasByPreferredAdminLangcodeAndMailAndDefaultLangcode(offset int, limit int, PreferredAdminLangcode_ string, Mail_ string, DefaultLangcode_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_admin_langcode = ? and mail = ? and default_langcode = ?", PreferredAdminLangcode_, Mail_, DefaultLangcode_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredAdminLangcodeAndTimezoneAndStatus Get UsersFieldDatas via PreferredAdminLangcodeAndTimezoneAndStatus
func GetUsersFieldDatasByPreferredAdminLangcodeAndTimezoneAndStatus(offset int, limit int, PreferredAdminLangcode_ string, Timezone_ string, Status_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_admin_langcode = ? and timezone = ? and status = ?", PreferredAdminLangcode_, Timezone_, Status_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredAdminLangcodeAndTimezoneAndCreated Get UsersFieldDatas via PreferredAdminLangcodeAndTimezoneAndCreated
func GetUsersFieldDatasByPreferredAdminLangcodeAndTimezoneAndCreated(offset int, limit int, PreferredAdminLangcode_ string, Timezone_ string, Created_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_admin_langcode = ? and timezone = ? and created = ?", PreferredAdminLangcode_, Timezone_, Created_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredAdminLangcodeAndTimezoneAndChanged Get UsersFieldDatas via PreferredAdminLangcodeAndTimezoneAndChanged
func GetUsersFieldDatasByPreferredAdminLangcodeAndTimezoneAndChanged(offset int, limit int, PreferredAdminLangcode_ string, Timezone_ string, Changed_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_admin_langcode = ? and timezone = ? and changed = ?", PreferredAdminLangcode_, Timezone_, Changed_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredAdminLangcodeAndTimezoneAndAccess Get UsersFieldDatas via PreferredAdminLangcodeAndTimezoneAndAccess
func GetUsersFieldDatasByPreferredAdminLangcodeAndTimezoneAndAccess(offset int, limit int, PreferredAdminLangcode_ string, Timezone_ string, Access_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_admin_langcode = ? and timezone = ? and access = ?", PreferredAdminLangcode_, Timezone_, Access_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredAdminLangcodeAndTimezoneAndLogin Get UsersFieldDatas via PreferredAdminLangcodeAndTimezoneAndLogin
func GetUsersFieldDatasByPreferredAdminLangcodeAndTimezoneAndLogin(offset int, limit int, PreferredAdminLangcode_ string, Timezone_ string, Login_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_admin_langcode = ? and timezone = ? and login = ?", PreferredAdminLangcode_, Timezone_, Login_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredAdminLangcodeAndTimezoneAndInit Get UsersFieldDatas via PreferredAdminLangcodeAndTimezoneAndInit
func GetUsersFieldDatasByPreferredAdminLangcodeAndTimezoneAndInit(offset int, limit int, PreferredAdminLangcode_ string, Timezone_ string, Init_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_admin_langcode = ? and timezone = ? and init = ?", PreferredAdminLangcode_, Timezone_, Init_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredAdminLangcodeAndTimezoneAndDefaultLangcode Get UsersFieldDatas via PreferredAdminLangcodeAndTimezoneAndDefaultLangcode
func GetUsersFieldDatasByPreferredAdminLangcodeAndTimezoneAndDefaultLangcode(offset int, limit int, PreferredAdminLangcode_ string, Timezone_ string, DefaultLangcode_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_admin_langcode = ? and timezone = ? and default_langcode = ?", PreferredAdminLangcode_, Timezone_, DefaultLangcode_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredAdminLangcodeAndStatusAndCreated Get UsersFieldDatas via PreferredAdminLangcodeAndStatusAndCreated
func GetUsersFieldDatasByPreferredAdminLangcodeAndStatusAndCreated(offset int, limit int, PreferredAdminLangcode_ string, Status_ int, Created_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_admin_langcode = ? and status = ? and created = ?", PreferredAdminLangcode_, Status_, Created_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredAdminLangcodeAndStatusAndChanged Get UsersFieldDatas via PreferredAdminLangcodeAndStatusAndChanged
func GetUsersFieldDatasByPreferredAdminLangcodeAndStatusAndChanged(offset int, limit int, PreferredAdminLangcode_ string, Status_ int, Changed_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_admin_langcode = ? and status = ? and changed = ?", PreferredAdminLangcode_, Status_, Changed_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredAdminLangcodeAndStatusAndAccess Get UsersFieldDatas via PreferredAdminLangcodeAndStatusAndAccess
func GetUsersFieldDatasByPreferredAdminLangcodeAndStatusAndAccess(offset int, limit int, PreferredAdminLangcode_ string, Status_ int, Access_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_admin_langcode = ? and status = ? and access = ?", PreferredAdminLangcode_, Status_, Access_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredAdminLangcodeAndStatusAndLogin Get UsersFieldDatas via PreferredAdminLangcodeAndStatusAndLogin
func GetUsersFieldDatasByPreferredAdminLangcodeAndStatusAndLogin(offset int, limit int, PreferredAdminLangcode_ string, Status_ int, Login_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_admin_langcode = ? and status = ? and login = ?", PreferredAdminLangcode_, Status_, Login_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredAdminLangcodeAndStatusAndInit Get UsersFieldDatas via PreferredAdminLangcodeAndStatusAndInit
func GetUsersFieldDatasByPreferredAdminLangcodeAndStatusAndInit(offset int, limit int, PreferredAdminLangcode_ string, Status_ int, Init_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_admin_langcode = ? and status = ? and init = ?", PreferredAdminLangcode_, Status_, Init_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredAdminLangcodeAndStatusAndDefaultLangcode Get UsersFieldDatas via PreferredAdminLangcodeAndStatusAndDefaultLangcode
func GetUsersFieldDatasByPreferredAdminLangcodeAndStatusAndDefaultLangcode(offset int, limit int, PreferredAdminLangcode_ string, Status_ int, DefaultLangcode_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_admin_langcode = ? and status = ? and default_langcode = ?", PreferredAdminLangcode_, Status_, DefaultLangcode_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredAdminLangcodeAndCreatedAndChanged Get UsersFieldDatas via PreferredAdminLangcodeAndCreatedAndChanged
func GetUsersFieldDatasByPreferredAdminLangcodeAndCreatedAndChanged(offset int, limit int, PreferredAdminLangcode_ string, Created_ int, Changed_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_admin_langcode = ? and created = ? and changed = ?", PreferredAdminLangcode_, Created_, Changed_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredAdminLangcodeAndCreatedAndAccess Get UsersFieldDatas via PreferredAdminLangcodeAndCreatedAndAccess
func GetUsersFieldDatasByPreferredAdminLangcodeAndCreatedAndAccess(offset int, limit int, PreferredAdminLangcode_ string, Created_ int, Access_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_admin_langcode = ? and created = ? and access = ?", PreferredAdminLangcode_, Created_, Access_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredAdminLangcodeAndCreatedAndLogin Get UsersFieldDatas via PreferredAdminLangcodeAndCreatedAndLogin
func GetUsersFieldDatasByPreferredAdminLangcodeAndCreatedAndLogin(offset int, limit int, PreferredAdminLangcode_ string, Created_ int, Login_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_admin_langcode = ? and created = ? and login = ?", PreferredAdminLangcode_, Created_, Login_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredAdminLangcodeAndCreatedAndInit Get UsersFieldDatas via PreferredAdminLangcodeAndCreatedAndInit
func GetUsersFieldDatasByPreferredAdminLangcodeAndCreatedAndInit(offset int, limit int, PreferredAdminLangcode_ string, Created_ int, Init_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_admin_langcode = ? and created = ? and init = ?", PreferredAdminLangcode_, Created_, Init_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredAdminLangcodeAndCreatedAndDefaultLangcode Get UsersFieldDatas via PreferredAdminLangcodeAndCreatedAndDefaultLangcode
func GetUsersFieldDatasByPreferredAdminLangcodeAndCreatedAndDefaultLangcode(offset int, limit int, PreferredAdminLangcode_ string, Created_ int, DefaultLangcode_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_admin_langcode = ? and created = ? and default_langcode = ?", PreferredAdminLangcode_, Created_, DefaultLangcode_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredAdminLangcodeAndChangedAndAccess Get UsersFieldDatas via PreferredAdminLangcodeAndChangedAndAccess
func GetUsersFieldDatasByPreferredAdminLangcodeAndChangedAndAccess(offset int, limit int, PreferredAdminLangcode_ string, Changed_ int, Access_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_admin_langcode = ? and changed = ? and access = ?", PreferredAdminLangcode_, Changed_, Access_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredAdminLangcodeAndChangedAndLogin Get UsersFieldDatas via PreferredAdminLangcodeAndChangedAndLogin
func GetUsersFieldDatasByPreferredAdminLangcodeAndChangedAndLogin(offset int, limit int, PreferredAdminLangcode_ string, Changed_ int, Login_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_admin_langcode = ? and changed = ? and login = ?", PreferredAdminLangcode_, Changed_, Login_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredAdminLangcodeAndChangedAndInit Get UsersFieldDatas via PreferredAdminLangcodeAndChangedAndInit
func GetUsersFieldDatasByPreferredAdminLangcodeAndChangedAndInit(offset int, limit int, PreferredAdminLangcode_ string, Changed_ int, Init_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_admin_langcode = ? and changed = ? and init = ?", PreferredAdminLangcode_, Changed_, Init_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredAdminLangcodeAndChangedAndDefaultLangcode Get UsersFieldDatas via PreferredAdminLangcodeAndChangedAndDefaultLangcode
func GetUsersFieldDatasByPreferredAdminLangcodeAndChangedAndDefaultLangcode(offset int, limit int, PreferredAdminLangcode_ string, Changed_ int, DefaultLangcode_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_admin_langcode = ? and changed = ? and default_langcode = ?", PreferredAdminLangcode_, Changed_, DefaultLangcode_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredAdminLangcodeAndAccessAndLogin Get UsersFieldDatas via PreferredAdminLangcodeAndAccessAndLogin
func GetUsersFieldDatasByPreferredAdminLangcodeAndAccessAndLogin(offset int, limit int, PreferredAdminLangcode_ string, Access_ int, Login_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_admin_langcode = ? and access = ? and login = ?", PreferredAdminLangcode_, Access_, Login_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredAdminLangcodeAndAccessAndInit Get UsersFieldDatas via PreferredAdminLangcodeAndAccessAndInit
func GetUsersFieldDatasByPreferredAdminLangcodeAndAccessAndInit(offset int, limit int, PreferredAdminLangcode_ string, Access_ int, Init_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_admin_langcode = ? and access = ? and init = ?", PreferredAdminLangcode_, Access_, Init_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredAdminLangcodeAndAccessAndDefaultLangcode Get UsersFieldDatas via PreferredAdminLangcodeAndAccessAndDefaultLangcode
func GetUsersFieldDatasByPreferredAdminLangcodeAndAccessAndDefaultLangcode(offset int, limit int, PreferredAdminLangcode_ string, Access_ int, DefaultLangcode_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_admin_langcode = ? and access = ? and default_langcode = ?", PreferredAdminLangcode_, Access_, DefaultLangcode_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredAdminLangcodeAndLoginAndInit Get UsersFieldDatas via PreferredAdminLangcodeAndLoginAndInit
func GetUsersFieldDatasByPreferredAdminLangcodeAndLoginAndInit(offset int, limit int, PreferredAdminLangcode_ string, Login_ int, Init_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_admin_langcode = ? and login = ? and init = ?", PreferredAdminLangcode_, Login_, Init_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredAdminLangcodeAndLoginAndDefaultLangcode Get UsersFieldDatas via PreferredAdminLangcodeAndLoginAndDefaultLangcode
func GetUsersFieldDatasByPreferredAdminLangcodeAndLoginAndDefaultLangcode(offset int, limit int, PreferredAdminLangcode_ string, Login_ int, DefaultLangcode_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_admin_langcode = ? and login = ? and default_langcode = ?", PreferredAdminLangcode_, Login_, DefaultLangcode_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredAdminLangcodeAndInitAndDefaultLangcode Get UsersFieldDatas via PreferredAdminLangcodeAndInitAndDefaultLangcode
func GetUsersFieldDatasByPreferredAdminLangcodeAndInitAndDefaultLangcode(offset int, limit int, PreferredAdminLangcode_ string, Init_ string, DefaultLangcode_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_admin_langcode = ? and init = ? and default_langcode = ?", PreferredAdminLangcode_, Init_, DefaultLangcode_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByNameAndPassAndMail Get UsersFieldDatas via NameAndPassAndMail
func GetUsersFieldDatasByNameAndPassAndMail(offset int, limit int, Name_ string, Pass_ string, Mail_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("name = ? and pass = ? and mail = ?", Name_, Pass_, Mail_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByNameAndPassAndTimezone Get UsersFieldDatas via NameAndPassAndTimezone
func GetUsersFieldDatasByNameAndPassAndTimezone(offset int, limit int, Name_ string, Pass_ string, Timezone_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("name = ? and pass = ? and timezone = ?", Name_, Pass_, Timezone_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByNameAndPassAndStatus Get UsersFieldDatas via NameAndPassAndStatus
func GetUsersFieldDatasByNameAndPassAndStatus(offset int, limit int, Name_ string, Pass_ string, Status_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("name = ? and pass = ? and status = ?", Name_, Pass_, Status_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByNameAndPassAndCreated Get UsersFieldDatas via NameAndPassAndCreated
func GetUsersFieldDatasByNameAndPassAndCreated(offset int, limit int, Name_ string, Pass_ string, Created_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("name = ? and pass = ? and created = ?", Name_, Pass_, Created_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByNameAndPassAndChanged Get UsersFieldDatas via NameAndPassAndChanged
func GetUsersFieldDatasByNameAndPassAndChanged(offset int, limit int, Name_ string, Pass_ string, Changed_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("name = ? and pass = ? and changed = ?", Name_, Pass_, Changed_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByNameAndPassAndAccess Get UsersFieldDatas via NameAndPassAndAccess
func GetUsersFieldDatasByNameAndPassAndAccess(offset int, limit int, Name_ string, Pass_ string, Access_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("name = ? and pass = ? and access = ?", Name_, Pass_, Access_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByNameAndPassAndLogin Get UsersFieldDatas via NameAndPassAndLogin
func GetUsersFieldDatasByNameAndPassAndLogin(offset int, limit int, Name_ string, Pass_ string, Login_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("name = ? and pass = ? and login = ?", Name_, Pass_, Login_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByNameAndPassAndInit Get UsersFieldDatas via NameAndPassAndInit
func GetUsersFieldDatasByNameAndPassAndInit(offset int, limit int, Name_ string, Pass_ string, Init_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("name = ? and pass = ? and init = ?", Name_, Pass_, Init_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByNameAndPassAndDefaultLangcode Get UsersFieldDatas via NameAndPassAndDefaultLangcode
func GetUsersFieldDatasByNameAndPassAndDefaultLangcode(offset int, limit int, Name_ string, Pass_ string, DefaultLangcode_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("name = ? and pass = ? and default_langcode = ?", Name_, Pass_, DefaultLangcode_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByNameAndMailAndTimezone Get UsersFieldDatas via NameAndMailAndTimezone
func GetUsersFieldDatasByNameAndMailAndTimezone(offset int, limit int, Name_ string, Mail_ string, Timezone_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("name = ? and mail = ? and timezone = ?", Name_, Mail_, Timezone_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByNameAndMailAndStatus Get UsersFieldDatas via NameAndMailAndStatus
func GetUsersFieldDatasByNameAndMailAndStatus(offset int, limit int, Name_ string, Mail_ string, Status_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("name = ? and mail = ? and status = ?", Name_, Mail_, Status_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByNameAndMailAndCreated Get UsersFieldDatas via NameAndMailAndCreated
func GetUsersFieldDatasByNameAndMailAndCreated(offset int, limit int, Name_ string, Mail_ string, Created_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("name = ? and mail = ? and created = ?", Name_, Mail_, Created_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByNameAndMailAndChanged Get UsersFieldDatas via NameAndMailAndChanged
func GetUsersFieldDatasByNameAndMailAndChanged(offset int, limit int, Name_ string, Mail_ string, Changed_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("name = ? and mail = ? and changed = ?", Name_, Mail_, Changed_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByNameAndMailAndAccess Get UsersFieldDatas via NameAndMailAndAccess
func GetUsersFieldDatasByNameAndMailAndAccess(offset int, limit int, Name_ string, Mail_ string, Access_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("name = ? and mail = ? and access = ?", Name_, Mail_, Access_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByNameAndMailAndLogin Get UsersFieldDatas via NameAndMailAndLogin
func GetUsersFieldDatasByNameAndMailAndLogin(offset int, limit int, Name_ string, Mail_ string, Login_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("name = ? and mail = ? and login = ?", Name_, Mail_, Login_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByNameAndMailAndInit Get UsersFieldDatas via NameAndMailAndInit
func GetUsersFieldDatasByNameAndMailAndInit(offset int, limit int, Name_ string, Mail_ string, Init_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("name = ? and mail = ? and init = ?", Name_, Mail_, Init_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByNameAndMailAndDefaultLangcode Get UsersFieldDatas via NameAndMailAndDefaultLangcode
func GetUsersFieldDatasByNameAndMailAndDefaultLangcode(offset int, limit int, Name_ string, Mail_ string, DefaultLangcode_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("name = ? and mail = ? and default_langcode = ?", Name_, Mail_, DefaultLangcode_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByNameAndTimezoneAndStatus Get UsersFieldDatas via NameAndTimezoneAndStatus
func GetUsersFieldDatasByNameAndTimezoneAndStatus(offset int, limit int, Name_ string, Timezone_ string, Status_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("name = ? and timezone = ? and status = ?", Name_, Timezone_, Status_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByNameAndTimezoneAndCreated Get UsersFieldDatas via NameAndTimezoneAndCreated
func GetUsersFieldDatasByNameAndTimezoneAndCreated(offset int, limit int, Name_ string, Timezone_ string, Created_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("name = ? and timezone = ? and created = ?", Name_, Timezone_, Created_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByNameAndTimezoneAndChanged Get UsersFieldDatas via NameAndTimezoneAndChanged
func GetUsersFieldDatasByNameAndTimezoneAndChanged(offset int, limit int, Name_ string, Timezone_ string, Changed_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("name = ? and timezone = ? and changed = ?", Name_, Timezone_, Changed_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByNameAndTimezoneAndAccess Get UsersFieldDatas via NameAndTimezoneAndAccess
func GetUsersFieldDatasByNameAndTimezoneAndAccess(offset int, limit int, Name_ string, Timezone_ string, Access_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("name = ? and timezone = ? and access = ?", Name_, Timezone_, Access_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByNameAndTimezoneAndLogin Get UsersFieldDatas via NameAndTimezoneAndLogin
func GetUsersFieldDatasByNameAndTimezoneAndLogin(offset int, limit int, Name_ string, Timezone_ string, Login_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("name = ? and timezone = ? and login = ?", Name_, Timezone_, Login_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByNameAndTimezoneAndInit Get UsersFieldDatas via NameAndTimezoneAndInit
func GetUsersFieldDatasByNameAndTimezoneAndInit(offset int, limit int, Name_ string, Timezone_ string, Init_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("name = ? and timezone = ? and init = ?", Name_, Timezone_, Init_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByNameAndTimezoneAndDefaultLangcode Get UsersFieldDatas via NameAndTimezoneAndDefaultLangcode
func GetUsersFieldDatasByNameAndTimezoneAndDefaultLangcode(offset int, limit int, Name_ string, Timezone_ string, DefaultLangcode_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("name = ? and timezone = ? and default_langcode = ?", Name_, Timezone_, DefaultLangcode_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByNameAndStatusAndCreated Get UsersFieldDatas via NameAndStatusAndCreated
func GetUsersFieldDatasByNameAndStatusAndCreated(offset int, limit int, Name_ string, Status_ int, Created_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("name = ? and status = ? and created = ?", Name_, Status_, Created_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByNameAndStatusAndChanged Get UsersFieldDatas via NameAndStatusAndChanged
func GetUsersFieldDatasByNameAndStatusAndChanged(offset int, limit int, Name_ string, Status_ int, Changed_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("name = ? and status = ? and changed = ?", Name_, Status_, Changed_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByNameAndStatusAndAccess Get UsersFieldDatas via NameAndStatusAndAccess
func GetUsersFieldDatasByNameAndStatusAndAccess(offset int, limit int, Name_ string, Status_ int, Access_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("name = ? and status = ? and access = ?", Name_, Status_, Access_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByNameAndStatusAndLogin Get UsersFieldDatas via NameAndStatusAndLogin
func GetUsersFieldDatasByNameAndStatusAndLogin(offset int, limit int, Name_ string, Status_ int, Login_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("name = ? and status = ? and login = ?", Name_, Status_, Login_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByNameAndStatusAndInit Get UsersFieldDatas via NameAndStatusAndInit
func GetUsersFieldDatasByNameAndStatusAndInit(offset int, limit int, Name_ string, Status_ int, Init_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("name = ? and status = ? and init = ?", Name_, Status_, Init_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByNameAndStatusAndDefaultLangcode Get UsersFieldDatas via NameAndStatusAndDefaultLangcode
func GetUsersFieldDatasByNameAndStatusAndDefaultLangcode(offset int, limit int, Name_ string, Status_ int, DefaultLangcode_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("name = ? and status = ? and default_langcode = ?", Name_, Status_, DefaultLangcode_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByNameAndCreatedAndChanged Get UsersFieldDatas via NameAndCreatedAndChanged
func GetUsersFieldDatasByNameAndCreatedAndChanged(offset int, limit int, Name_ string, Created_ int, Changed_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("name = ? and created = ? and changed = ?", Name_, Created_, Changed_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByNameAndCreatedAndAccess Get UsersFieldDatas via NameAndCreatedAndAccess
func GetUsersFieldDatasByNameAndCreatedAndAccess(offset int, limit int, Name_ string, Created_ int, Access_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("name = ? and created = ? and access = ?", Name_, Created_, Access_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByNameAndCreatedAndLogin Get UsersFieldDatas via NameAndCreatedAndLogin
func GetUsersFieldDatasByNameAndCreatedAndLogin(offset int, limit int, Name_ string, Created_ int, Login_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("name = ? and created = ? and login = ?", Name_, Created_, Login_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByNameAndCreatedAndInit Get UsersFieldDatas via NameAndCreatedAndInit
func GetUsersFieldDatasByNameAndCreatedAndInit(offset int, limit int, Name_ string, Created_ int, Init_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("name = ? and created = ? and init = ?", Name_, Created_, Init_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByNameAndCreatedAndDefaultLangcode Get UsersFieldDatas via NameAndCreatedAndDefaultLangcode
func GetUsersFieldDatasByNameAndCreatedAndDefaultLangcode(offset int, limit int, Name_ string, Created_ int, DefaultLangcode_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("name = ? and created = ? and default_langcode = ?", Name_, Created_, DefaultLangcode_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByNameAndChangedAndAccess Get UsersFieldDatas via NameAndChangedAndAccess
func GetUsersFieldDatasByNameAndChangedAndAccess(offset int, limit int, Name_ string, Changed_ int, Access_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("name = ? and changed = ? and access = ?", Name_, Changed_, Access_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByNameAndChangedAndLogin Get UsersFieldDatas via NameAndChangedAndLogin
func GetUsersFieldDatasByNameAndChangedAndLogin(offset int, limit int, Name_ string, Changed_ int, Login_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("name = ? and changed = ? and login = ?", Name_, Changed_, Login_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByNameAndChangedAndInit Get UsersFieldDatas via NameAndChangedAndInit
func GetUsersFieldDatasByNameAndChangedAndInit(offset int, limit int, Name_ string, Changed_ int, Init_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("name = ? and changed = ? and init = ?", Name_, Changed_, Init_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByNameAndChangedAndDefaultLangcode Get UsersFieldDatas via NameAndChangedAndDefaultLangcode
func GetUsersFieldDatasByNameAndChangedAndDefaultLangcode(offset int, limit int, Name_ string, Changed_ int, DefaultLangcode_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("name = ? and changed = ? and default_langcode = ?", Name_, Changed_, DefaultLangcode_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByNameAndAccessAndLogin Get UsersFieldDatas via NameAndAccessAndLogin
func GetUsersFieldDatasByNameAndAccessAndLogin(offset int, limit int, Name_ string, Access_ int, Login_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("name = ? and access = ? and login = ?", Name_, Access_, Login_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByNameAndAccessAndInit Get UsersFieldDatas via NameAndAccessAndInit
func GetUsersFieldDatasByNameAndAccessAndInit(offset int, limit int, Name_ string, Access_ int, Init_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("name = ? and access = ? and init = ?", Name_, Access_, Init_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByNameAndAccessAndDefaultLangcode Get UsersFieldDatas via NameAndAccessAndDefaultLangcode
func GetUsersFieldDatasByNameAndAccessAndDefaultLangcode(offset int, limit int, Name_ string, Access_ int, DefaultLangcode_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("name = ? and access = ? and default_langcode = ?", Name_, Access_, DefaultLangcode_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByNameAndLoginAndInit Get UsersFieldDatas via NameAndLoginAndInit
func GetUsersFieldDatasByNameAndLoginAndInit(offset int, limit int, Name_ string, Login_ int, Init_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("name = ? and login = ? and init = ?", Name_, Login_, Init_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByNameAndLoginAndDefaultLangcode Get UsersFieldDatas via NameAndLoginAndDefaultLangcode
func GetUsersFieldDatasByNameAndLoginAndDefaultLangcode(offset int, limit int, Name_ string, Login_ int, DefaultLangcode_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("name = ? and login = ? and default_langcode = ?", Name_, Login_, DefaultLangcode_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByNameAndInitAndDefaultLangcode Get UsersFieldDatas via NameAndInitAndDefaultLangcode
func GetUsersFieldDatasByNameAndInitAndDefaultLangcode(offset int, limit int, Name_ string, Init_ string, DefaultLangcode_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("name = ? and init = ? and default_langcode = ?", Name_, Init_, DefaultLangcode_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPassAndMailAndTimezone Get UsersFieldDatas via PassAndMailAndTimezone
func GetUsersFieldDatasByPassAndMailAndTimezone(offset int, limit int, Pass_ string, Mail_ string, Timezone_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("pass = ? and mail = ? and timezone = ?", Pass_, Mail_, Timezone_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPassAndMailAndStatus Get UsersFieldDatas via PassAndMailAndStatus
func GetUsersFieldDatasByPassAndMailAndStatus(offset int, limit int, Pass_ string, Mail_ string, Status_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("pass = ? and mail = ? and status = ?", Pass_, Mail_, Status_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPassAndMailAndCreated Get UsersFieldDatas via PassAndMailAndCreated
func GetUsersFieldDatasByPassAndMailAndCreated(offset int, limit int, Pass_ string, Mail_ string, Created_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("pass = ? and mail = ? and created = ?", Pass_, Mail_, Created_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPassAndMailAndChanged Get UsersFieldDatas via PassAndMailAndChanged
func GetUsersFieldDatasByPassAndMailAndChanged(offset int, limit int, Pass_ string, Mail_ string, Changed_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("pass = ? and mail = ? and changed = ?", Pass_, Mail_, Changed_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPassAndMailAndAccess Get UsersFieldDatas via PassAndMailAndAccess
func GetUsersFieldDatasByPassAndMailAndAccess(offset int, limit int, Pass_ string, Mail_ string, Access_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("pass = ? and mail = ? and access = ?", Pass_, Mail_, Access_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPassAndMailAndLogin Get UsersFieldDatas via PassAndMailAndLogin
func GetUsersFieldDatasByPassAndMailAndLogin(offset int, limit int, Pass_ string, Mail_ string, Login_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("pass = ? and mail = ? and login = ?", Pass_, Mail_, Login_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPassAndMailAndInit Get UsersFieldDatas via PassAndMailAndInit
func GetUsersFieldDatasByPassAndMailAndInit(offset int, limit int, Pass_ string, Mail_ string, Init_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("pass = ? and mail = ? and init = ?", Pass_, Mail_, Init_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPassAndMailAndDefaultLangcode Get UsersFieldDatas via PassAndMailAndDefaultLangcode
func GetUsersFieldDatasByPassAndMailAndDefaultLangcode(offset int, limit int, Pass_ string, Mail_ string, DefaultLangcode_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("pass = ? and mail = ? and default_langcode = ?", Pass_, Mail_, DefaultLangcode_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPassAndTimezoneAndStatus Get UsersFieldDatas via PassAndTimezoneAndStatus
func GetUsersFieldDatasByPassAndTimezoneAndStatus(offset int, limit int, Pass_ string, Timezone_ string, Status_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("pass = ? and timezone = ? and status = ?", Pass_, Timezone_, Status_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPassAndTimezoneAndCreated Get UsersFieldDatas via PassAndTimezoneAndCreated
func GetUsersFieldDatasByPassAndTimezoneAndCreated(offset int, limit int, Pass_ string, Timezone_ string, Created_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("pass = ? and timezone = ? and created = ?", Pass_, Timezone_, Created_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPassAndTimezoneAndChanged Get UsersFieldDatas via PassAndTimezoneAndChanged
func GetUsersFieldDatasByPassAndTimezoneAndChanged(offset int, limit int, Pass_ string, Timezone_ string, Changed_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("pass = ? and timezone = ? and changed = ?", Pass_, Timezone_, Changed_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPassAndTimezoneAndAccess Get UsersFieldDatas via PassAndTimezoneAndAccess
func GetUsersFieldDatasByPassAndTimezoneAndAccess(offset int, limit int, Pass_ string, Timezone_ string, Access_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("pass = ? and timezone = ? and access = ?", Pass_, Timezone_, Access_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPassAndTimezoneAndLogin Get UsersFieldDatas via PassAndTimezoneAndLogin
func GetUsersFieldDatasByPassAndTimezoneAndLogin(offset int, limit int, Pass_ string, Timezone_ string, Login_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("pass = ? and timezone = ? and login = ?", Pass_, Timezone_, Login_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPassAndTimezoneAndInit Get UsersFieldDatas via PassAndTimezoneAndInit
func GetUsersFieldDatasByPassAndTimezoneAndInit(offset int, limit int, Pass_ string, Timezone_ string, Init_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("pass = ? and timezone = ? and init = ?", Pass_, Timezone_, Init_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPassAndTimezoneAndDefaultLangcode Get UsersFieldDatas via PassAndTimezoneAndDefaultLangcode
func GetUsersFieldDatasByPassAndTimezoneAndDefaultLangcode(offset int, limit int, Pass_ string, Timezone_ string, DefaultLangcode_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("pass = ? and timezone = ? and default_langcode = ?", Pass_, Timezone_, DefaultLangcode_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPassAndStatusAndCreated Get UsersFieldDatas via PassAndStatusAndCreated
func GetUsersFieldDatasByPassAndStatusAndCreated(offset int, limit int, Pass_ string, Status_ int, Created_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("pass = ? and status = ? and created = ?", Pass_, Status_, Created_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPassAndStatusAndChanged Get UsersFieldDatas via PassAndStatusAndChanged
func GetUsersFieldDatasByPassAndStatusAndChanged(offset int, limit int, Pass_ string, Status_ int, Changed_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("pass = ? and status = ? and changed = ?", Pass_, Status_, Changed_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPassAndStatusAndAccess Get UsersFieldDatas via PassAndStatusAndAccess
func GetUsersFieldDatasByPassAndStatusAndAccess(offset int, limit int, Pass_ string, Status_ int, Access_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("pass = ? and status = ? and access = ?", Pass_, Status_, Access_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPassAndStatusAndLogin Get UsersFieldDatas via PassAndStatusAndLogin
func GetUsersFieldDatasByPassAndStatusAndLogin(offset int, limit int, Pass_ string, Status_ int, Login_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("pass = ? and status = ? and login = ?", Pass_, Status_, Login_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPassAndStatusAndInit Get UsersFieldDatas via PassAndStatusAndInit
func GetUsersFieldDatasByPassAndStatusAndInit(offset int, limit int, Pass_ string, Status_ int, Init_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("pass = ? and status = ? and init = ?", Pass_, Status_, Init_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPassAndStatusAndDefaultLangcode Get UsersFieldDatas via PassAndStatusAndDefaultLangcode
func GetUsersFieldDatasByPassAndStatusAndDefaultLangcode(offset int, limit int, Pass_ string, Status_ int, DefaultLangcode_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("pass = ? and status = ? and default_langcode = ?", Pass_, Status_, DefaultLangcode_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPassAndCreatedAndChanged Get UsersFieldDatas via PassAndCreatedAndChanged
func GetUsersFieldDatasByPassAndCreatedAndChanged(offset int, limit int, Pass_ string, Created_ int, Changed_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("pass = ? and created = ? and changed = ?", Pass_, Created_, Changed_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPassAndCreatedAndAccess Get UsersFieldDatas via PassAndCreatedAndAccess
func GetUsersFieldDatasByPassAndCreatedAndAccess(offset int, limit int, Pass_ string, Created_ int, Access_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("pass = ? and created = ? and access = ?", Pass_, Created_, Access_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPassAndCreatedAndLogin Get UsersFieldDatas via PassAndCreatedAndLogin
func GetUsersFieldDatasByPassAndCreatedAndLogin(offset int, limit int, Pass_ string, Created_ int, Login_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("pass = ? and created = ? and login = ?", Pass_, Created_, Login_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPassAndCreatedAndInit Get UsersFieldDatas via PassAndCreatedAndInit
func GetUsersFieldDatasByPassAndCreatedAndInit(offset int, limit int, Pass_ string, Created_ int, Init_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("pass = ? and created = ? and init = ?", Pass_, Created_, Init_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPassAndCreatedAndDefaultLangcode Get UsersFieldDatas via PassAndCreatedAndDefaultLangcode
func GetUsersFieldDatasByPassAndCreatedAndDefaultLangcode(offset int, limit int, Pass_ string, Created_ int, DefaultLangcode_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("pass = ? and created = ? and default_langcode = ?", Pass_, Created_, DefaultLangcode_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPassAndChangedAndAccess Get UsersFieldDatas via PassAndChangedAndAccess
func GetUsersFieldDatasByPassAndChangedAndAccess(offset int, limit int, Pass_ string, Changed_ int, Access_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("pass = ? and changed = ? and access = ?", Pass_, Changed_, Access_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPassAndChangedAndLogin Get UsersFieldDatas via PassAndChangedAndLogin
func GetUsersFieldDatasByPassAndChangedAndLogin(offset int, limit int, Pass_ string, Changed_ int, Login_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("pass = ? and changed = ? and login = ?", Pass_, Changed_, Login_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPassAndChangedAndInit Get UsersFieldDatas via PassAndChangedAndInit
func GetUsersFieldDatasByPassAndChangedAndInit(offset int, limit int, Pass_ string, Changed_ int, Init_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("pass = ? and changed = ? and init = ?", Pass_, Changed_, Init_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPassAndChangedAndDefaultLangcode Get UsersFieldDatas via PassAndChangedAndDefaultLangcode
func GetUsersFieldDatasByPassAndChangedAndDefaultLangcode(offset int, limit int, Pass_ string, Changed_ int, DefaultLangcode_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("pass = ? and changed = ? and default_langcode = ?", Pass_, Changed_, DefaultLangcode_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPassAndAccessAndLogin Get UsersFieldDatas via PassAndAccessAndLogin
func GetUsersFieldDatasByPassAndAccessAndLogin(offset int, limit int, Pass_ string, Access_ int, Login_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("pass = ? and access = ? and login = ?", Pass_, Access_, Login_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPassAndAccessAndInit Get UsersFieldDatas via PassAndAccessAndInit
func GetUsersFieldDatasByPassAndAccessAndInit(offset int, limit int, Pass_ string, Access_ int, Init_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("pass = ? and access = ? and init = ?", Pass_, Access_, Init_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPassAndAccessAndDefaultLangcode Get UsersFieldDatas via PassAndAccessAndDefaultLangcode
func GetUsersFieldDatasByPassAndAccessAndDefaultLangcode(offset int, limit int, Pass_ string, Access_ int, DefaultLangcode_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("pass = ? and access = ? and default_langcode = ?", Pass_, Access_, DefaultLangcode_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPassAndLoginAndInit Get UsersFieldDatas via PassAndLoginAndInit
func GetUsersFieldDatasByPassAndLoginAndInit(offset int, limit int, Pass_ string, Login_ int, Init_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("pass = ? and login = ? and init = ?", Pass_, Login_, Init_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPassAndLoginAndDefaultLangcode Get UsersFieldDatas via PassAndLoginAndDefaultLangcode
func GetUsersFieldDatasByPassAndLoginAndDefaultLangcode(offset int, limit int, Pass_ string, Login_ int, DefaultLangcode_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("pass = ? and login = ? and default_langcode = ?", Pass_, Login_, DefaultLangcode_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPassAndInitAndDefaultLangcode Get UsersFieldDatas via PassAndInitAndDefaultLangcode
func GetUsersFieldDatasByPassAndInitAndDefaultLangcode(offset int, limit int, Pass_ string, Init_ string, DefaultLangcode_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("pass = ? and init = ? and default_langcode = ?", Pass_, Init_, DefaultLangcode_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByMailAndTimezoneAndStatus Get UsersFieldDatas via MailAndTimezoneAndStatus
func GetUsersFieldDatasByMailAndTimezoneAndStatus(offset int, limit int, Mail_ string, Timezone_ string, Status_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("mail = ? and timezone = ? and status = ?", Mail_, Timezone_, Status_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByMailAndTimezoneAndCreated Get UsersFieldDatas via MailAndTimezoneAndCreated
func GetUsersFieldDatasByMailAndTimezoneAndCreated(offset int, limit int, Mail_ string, Timezone_ string, Created_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("mail = ? and timezone = ? and created = ?", Mail_, Timezone_, Created_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByMailAndTimezoneAndChanged Get UsersFieldDatas via MailAndTimezoneAndChanged
func GetUsersFieldDatasByMailAndTimezoneAndChanged(offset int, limit int, Mail_ string, Timezone_ string, Changed_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("mail = ? and timezone = ? and changed = ?", Mail_, Timezone_, Changed_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByMailAndTimezoneAndAccess Get UsersFieldDatas via MailAndTimezoneAndAccess
func GetUsersFieldDatasByMailAndTimezoneAndAccess(offset int, limit int, Mail_ string, Timezone_ string, Access_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("mail = ? and timezone = ? and access = ?", Mail_, Timezone_, Access_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByMailAndTimezoneAndLogin Get UsersFieldDatas via MailAndTimezoneAndLogin
func GetUsersFieldDatasByMailAndTimezoneAndLogin(offset int, limit int, Mail_ string, Timezone_ string, Login_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("mail = ? and timezone = ? and login = ?", Mail_, Timezone_, Login_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByMailAndTimezoneAndInit Get UsersFieldDatas via MailAndTimezoneAndInit
func GetUsersFieldDatasByMailAndTimezoneAndInit(offset int, limit int, Mail_ string, Timezone_ string, Init_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("mail = ? and timezone = ? and init = ?", Mail_, Timezone_, Init_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByMailAndTimezoneAndDefaultLangcode Get UsersFieldDatas via MailAndTimezoneAndDefaultLangcode
func GetUsersFieldDatasByMailAndTimezoneAndDefaultLangcode(offset int, limit int, Mail_ string, Timezone_ string, DefaultLangcode_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("mail = ? and timezone = ? and default_langcode = ?", Mail_, Timezone_, DefaultLangcode_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByMailAndStatusAndCreated Get UsersFieldDatas via MailAndStatusAndCreated
func GetUsersFieldDatasByMailAndStatusAndCreated(offset int, limit int, Mail_ string, Status_ int, Created_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("mail = ? and status = ? and created = ?", Mail_, Status_, Created_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByMailAndStatusAndChanged Get UsersFieldDatas via MailAndStatusAndChanged
func GetUsersFieldDatasByMailAndStatusAndChanged(offset int, limit int, Mail_ string, Status_ int, Changed_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("mail = ? and status = ? and changed = ?", Mail_, Status_, Changed_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByMailAndStatusAndAccess Get UsersFieldDatas via MailAndStatusAndAccess
func GetUsersFieldDatasByMailAndStatusAndAccess(offset int, limit int, Mail_ string, Status_ int, Access_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("mail = ? and status = ? and access = ?", Mail_, Status_, Access_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByMailAndStatusAndLogin Get UsersFieldDatas via MailAndStatusAndLogin
func GetUsersFieldDatasByMailAndStatusAndLogin(offset int, limit int, Mail_ string, Status_ int, Login_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("mail = ? and status = ? and login = ?", Mail_, Status_, Login_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByMailAndStatusAndInit Get UsersFieldDatas via MailAndStatusAndInit
func GetUsersFieldDatasByMailAndStatusAndInit(offset int, limit int, Mail_ string, Status_ int, Init_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("mail = ? and status = ? and init = ?", Mail_, Status_, Init_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByMailAndStatusAndDefaultLangcode Get UsersFieldDatas via MailAndStatusAndDefaultLangcode
func GetUsersFieldDatasByMailAndStatusAndDefaultLangcode(offset int, limit int, Mail_ string, Status_ int, DefaultLangcode_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("mail = ? and status = ? and default_langcode = ?", Mail_, Status_, DefaultLangcode_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByMailAndCreatedAndChanged Get UsersFieldDatas via MailAndCreatedAndChanged
func GetUsersFieldDatasByMailAndCreatedAndChanged(offset int, limit int, Mail_ string, Created_ int, Changed_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("mail = ? and created = ? and changed = ?", Mail_, Created_, Changed_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByMailAndCreatedAndAccess Get UsersFieldDatas via MailAndCreatedAndAccess
func GetUsersFieldDatasByMailAndCreatedAndAccess(offset int, limit int, Mail_ string, Created_ int, Access_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("mail = ? and created = ? and access = ?", Mail_, Created_, Access_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByMailAndCreatedAndLogin Get UsersFieldDatas via MailAndCreatedAndLogin
func GetUsersFieldDatasByMailAndCreatedAndLogin(offset int, limit int, Mail_ string, Created_ int, Login_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("mail = ? and created = ? and login = ?", Mail_, Created_, Login_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByMailAndCreatedAndInit Get UsersFieldDatas via MailAndCreatedAndInit
func GetUsersFieldDatasByMailAndCreatedAndInit(offset int, limit int, Mail_ string, Created_ int, Init_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("mail = ? and created = ? and init = ?", Mail_, Created_, Init_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByMailAndCreatedAndDefaultLangcode Get UsersFieldDatas via MailAndCreatedAndDefaultLangcode
func GetUsersFieldDatasByMailAndCreatedAndDefaultLangcode(offset int, limit int, Mail_ string, Created_ int, DefaultLangcode_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("mail = ? and created = ? and default_langcode = ?", Mail_, Created_, DefaultLangcode_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByMailAndChangedAndAccess Get UsersFieldDatas via MailAndChangedAndAccess
func GetUsersFieldDatasByMailAndChangedAndAccess(offset int, limit int, Mail_ string, Changed_ int, Access_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("mail = ? and changed = ? and access = ?", Mail_, Changed_, Access_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByMailAndChangedAndLogin Get UsersFieldDatas via MailAndChangedAndLogin
func GetUsersFieldDatasByMailAndChangedAndLogin(offset int, limit int, Mail_ string, Changed_ int, Login_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("mail = ? and changed = ? and login = ?", Mail_, Changed_, Login_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByMailAndChangedAndInit Get UsersFieldDatas via MailAndChangedAndInit
func GetUsersFieldDatasByMailAndChangedAndInit(offset int, limit int, Mail_ string, Changed_ int, Init_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("mail = ? and changed = ? and init = ?", Mail_, Changed_, Init_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByMailAndChangedAndDefaultLangcode Get UsersFieldDatas via MailAndChangedAndDefaultLangcode
func GetUsersFieldDatasByMailAndChangedAndDefaultLangcode(offset int, limit int, Mail_ string, Changed_ int, DefaultLangcode_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("mail = ? and changed = ? and default_langcode = ?", Mail_, Changed_, DefaultLangcode_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByMailAndAccessAndLogin Get UsersFieldDatas via MailAndAccessAndLogin
func GetUsersFieldDatasByMailAndAccessAndLogin(offset int, limit int, Mail_ string, Access_ int, Login_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("mail = ? and access = ? and login = ?", Mail_, Access_, Login_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByMailAndAccessAndInit Get UsersFieldDatas via MailAndAccessAndInit
func GetUsersFieldDatasByMailAndAccessAndInit(offset int, limit int, Mail_ string, Access_ int, Init_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("mail = ? and access = ? and init = ?", Mail_, Access_, Init_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByMailAndAccessAndDefaultLangcode Get UsersFieldDatas via MailAndAccessAndDefaultLangcode
func GetUsersFieldDatasByMailAndAccessAndDefaultLangcode(offset int, limit int, Mail_ string, Access_ int, DefaultLangcode_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("mail = ? and access = ? and default_langcode = ?", Mail_, Access_, DefaultLangcode_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByMailAndLoginAndInit Get UsersFieldDatas via MailAndLoginAndInit
func GetUsersFieldDatasByMailAndLoginAndInit(offset int, limit int, Mail_ string, Login_ int, Init_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("mail = ? and login = ? and init = ?", Mail_, Login_, Init_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByMailAndLoginAndDefaultLangcode Get UsersFieldDatas via MailAndLoginAndDefaultLangcode
func GetUsersFieldDatasByMailAndLoginAndDefaultLangcode(offset int, limit int, Mail_ string, Login_ int, DefaultLangcode_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("mail = ? and login = ? and default_langcode = ?", Mail_, Login_, DefaultLangcode_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByMailAndInitAndDefaultLangcode Get UsersFieldDatas via MailAndInitAndDefaultLangcode
func GetUsersFieldDatasByMailAndInitAndDefaultLangcode(offset int, limit int, Mail_ string, Init_ string, DefaultLangcode_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("mail = ? and init = ? and default_langcode = ?", Mail_, Init_, DefaultLangcode_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByTimezoneAndStatusAndCreated Get UsersFieldDatas via TimezoneAndStatusAndCreated
func GetUsersFieldDatasByTimezoneAndStatusAndCreated(offset int, limit int, Timezone_ string, Status_ int, Created_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("timezone = ? and status = ? and created = ?", Timezone_, Status_, Created_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByTimezoneAndStatusAndChanged Get UsersFieldDatas via TimezoneAndStatusAndChanged
func GetUsersFieldDatasByTimezoneAndStatusAndChanged(offset int, limit int, Timezone_ string, Status_ int, Changed_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("timezone = ? and status = ? and changed = ?", Timezone_, Status_, Changed_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByTimezoneAndStatusAndAccess Get UsersFieldDatas via TimezoneAndStatusAndAccess
func GetUsersFieldDatasByTimezoneAndStatusAndAccess(offset int, limit int, Timezone_ string, Status_ int, Access_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("timezone = ? and status = ? and access = ?", Timezone_, Status_, Access_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByTimezoneAndStatusAndLogin Get UsersFieldDatas via TimezoneAndStatusAndLogin
func GetUsersFieldDatasByTimezoneAndStatusAndLogin(offset int, limit int, Timezone_ string, Status_ int, Login_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("timezone = ? and status = ? and login = ?", Timezone_, Status_, Login_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByTimezoneAndStatusAndInit Get UsersFieldDatas via TimezoneAndStatusAndInit
func GetUsersFieldDatasByTimezoneAndStatusAndInit(offset int, limit int, Timezone_ string, Status_ int, Init_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("timezone = ? and status = ? and init = ?", Timezone_, Status_, Init_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByTimezoneAndStatusAndDefaultLangcode Get UsersFieldDatas via TimezoneAndStatusAndDefaultLangcode
func GetUsersFieldDatasByTimezoneAndStatusAndDefaultLangcode(offset int, limit int, Timezone_ string, Status_ int, DefaultLangcode_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("timezone = ? and status = ? and default_langcode = ?", Timezone_, Status_, DefaultLangcode_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByTimezoneAndCreatedAndChanged Get UsersFieldDatas via TimezoneAndCreatedAndChanged
func GetUsersFieldDatasByTimezoneAndCreatedAndChanged(offset int, limit int, Timezone_ string, Created_ int, Changed_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("timezone = ? and created = ? and changed = ?", Timezone_, Created_, Changed_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByTimezoneAndCreatedAndAccess Get UsersFieldDatas via TimezoneAndCreatedAndAccess
func GetUsersFieldDatasByTimezoneAndCreatedAndAccess(offset int, limit int, Timezone_ string, Created_ int, Access_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("timezone = ? and created = ? and access = ?", Timezone_, Created_, Access_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByTimezoneAndCreatedAndLogin Get UsersFieldDatas via TimezoneAndCreatedAndLogin
func GetUsersFieldDatasByTimezoneAndCreatedAndLogin(offset int, limit int, Timezone_ string, Created_ int, Login_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("timezone = ? and created = ? and login = ?", Timezone_, Created_, Login_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByTimezoneAndCreatedAndInit Get UsersFieldDatas via TimezoneAndCreatedAndInit
func GetUsersFieldDatasByTimezoneAndCreatedAndInit(offset int, limit int, Timezone_ string, Created_ int, Init_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("timezone = ? and created = ? and init = ?", Timezone_, Created_, Init_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByTimezoneAndCreatedAndDefaultLangcode Get UsersFieldDatas via TimezoneAndCreatedAndDefaultLangcode
func GetUsersFieldDatasByTimezoneAndCreatedAndDefaultLangcode(offset int, limit int, Timezone_ string, Created_ int, DefaultLangcode_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("timezone = ? and created = ? and default_langcode = ?", Timezone_, Created_, DefaultLangcode_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByTimezoneAndChangedAndAccess Get UsersFieldDatas via TimezoneAndChangedAndAccess
func GetUsersFieldDatasByTimezoneAndChangedAndAccess(offset int, limit int, Timezone_ string, Changed_ int, Access_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("timezone = ? and changed = ? and access = ?", Timezone_, Changed_, Access_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByTimezoneAndChangedAndLogin Get UsersFieldDatas via TimezoneAndChangedAndLogin
func GetUsersFieldDatasByTimezoneAndChangedAndLogin(offset int, limit int, Timezone_ string, Changed_ int, Login_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("timezone = ? and changed = ? and login = ?", Timezone_, Changed_, Login_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByTimezoneAndChangedAndInit Get UsersFieldDatas via TimezoneAndChangedAndInit
func GetUsersFieldDatasByTimezoneAndChangedAndInit(offset int, limit int, Timezone_ string, Changed_ int, Init_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("timezone = ? and changed = ? and init = ?", Timezone_, Changed_, Init_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByTimezoneAndChangedAndDefaultLangcode Get UsersFieldDatas via TimezoneAndChangedAndDefaultLangcode
func GetUsersFieldDatasByTimezoneAndChangedAndDefaultLangcode(offset int, limit int, Timezone_ string, Changed_ int, DefaultLangcode_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("timezone = ? and changed = ? and default_langcode = ?", Timezone_, Changed_, DefaultLangcode_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByTimezoneAndAccessAndLogin Get UsersFieldDatas via TimezoneAndAccessAndLogin
func GetUsersFieldDatasByTimezoneAndAccessAndLogin(offset int, limit int, Timezone_ string, Access_ int, Login_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("timezone = ? and access = ? and login = ?", Timezone_, Access_, Login_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByTimezoneAndAccessAndInit Get UsersFieldDatas via TimezoneAndAccessAndInit
func GetUsersFieldDatasByTimezoneAndAccessAndInit(offset int, limit int, Timezone_ string, Access_ int, Init_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("timezone = ? and access = ? and init = ?", Timezone_, Access_, Init_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByTimezoneAndAccessAndDefaultLangcode Get UsersFieldDatas via TimezoneAndAccessAndDefaultLangcode
func GetUsersFieldDatasByTimezoneAndAccessAndDefaultLangcode(offset int, limit int, Timezone_ string, Access_ int, DefaultLangcode_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("timezone = ? and access = ? and default_langcode = ?", Timezone_, Access_, DefaultLangcode_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByTimezoneAndLoginAndInit Get UsersFieldDatas via TimezoneAndLoginAndInit
func GetUsersFieldDatasByTimezoneAndLoginAndInit(offset int, limit int, Timezone_ string, Login_ int, Init_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("timezone = ? and login = ? and init = ?", Timezone_, Login_, Init_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByTimezoneAndLoginAndDefaultLangcode Get UsersFieldDatas via TimezoneAndLoginAndDefaultLangcode
func GetUsersFieldDatasByTimezoneAndLoginAndDefaultLangcode(offset int, limit int, Timezone_ string, Login_ int, DefaultLangcode_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("timezone = ? and login = ? and default_langcode = ?", Timezone_, Login_, DefaultLangcode_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByTimezoneAndInitAndDefaultLangcode Get UsersFieldDatas via TimezoneAndInitAndDefaultLangcode
func GetUsersFieldDatasByTimezoneAndInitAndDefaultLangcode(offset int, limit int, Timezone_ string, Init_ string, DefaultLangcode_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("timezone = ? and init = ? and default_langcode = ?", Timezone_, Init_, DefaultLangcode_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByStatusAndCreatedAndChanged Get UsersFieldDatas via StatusAndCreatedAndChanged
func GetUsersFieldDatasByStatusAndCreatedAndChanged(offset int, limit int, Status_ int, Created_ int, Changed_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("status = ? and created = ? and changed = ?", Status_, Created_, Changed_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByStatusAndCreatedAndAccess Get UsersFieldDatas via StatusAndCreatedAndAccess
func GetUsersFieldDatasByStatusAndCreatedAndAccess(offset int, limit int, Status_ int, Created_ int, Access_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("status = ? and created = ? and access = ?", Status_, Created_, Access_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByStatusAndCreatedAndLogin Get UsersFieldDatas via StatusAndCreatedAndLogin
func GetUsersFieldDatasByStatusAndCreatedAndLogin(offset int, limit int, Status_ int, Created_ int, Login_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("status = ? and created = ? and login = ?", Status_, Created_, Login_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByStatusAndCreatedAndInit Get UsersFieldDatas via StatusAndCreatedAndInit
func GetUsersFieldDatasByStatusAndCreatedAndInit(offset int, limit int, Status_ int, Created_ int, Init_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("status = ? and created = ? and init = ?", Status_, Created_, Init_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByStatusAndCreatedAndDefaultLangcode Get UsersFieldDatas via StatusAndCreatedAndDefaultLangcode
func GetUsersFieldDatasByStatusAndCreatedAndDefaultLangcode(offset int, limit int, Status_ int, Created_ int, DefaultLangcode_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("status = ? and created = ? and default_langcode = ?", Status_, Created_, DefaultLangcode_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByStatusAndChangedAndAccess Get UsersFieldDatas via StatusAndChangedAndAccess
func GetUsersFieldDatasByStatusAndChangedAndAccess(offset int, limit int, Status_ int, Changed_ int, Access_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("status = ? and changed = ? and access = ?", Status_, Changed_, Access_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByStatusAndChangedAndLogin Get UsersFieldDatas via StatusAndChangedAndLogin
func GetUsersFieldDatasByStatusAndChangedAndLogin(offset int, limit int, Status_ int, Changed_ int, Login_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("status = ? and changed = ? and login = ?", Status_, Changed_, Login_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByStatusAndChangedAndInit Get UsersFieldDatas via StatusAndChangedAndInit
func GetUsersFieldDatasByStatusAndChangedAndInit(offset int, limit int, Status_ int, Changed_ int, Init_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("status = ? and changed = ? and init = ?", Status_, Changed_, Init_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByStatusAndChangedAndDefaultLangcode Get UsersFieldDatas via StatusAndChangedAndDefaultLangcode
func GetUsersFieldDatasByStatusAndChangedAndDefaultLangcode(offset int, limit int, Status_ int, Changed_ int, DefaultLangcode_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("status = ? and changed = ? and default_langcode = ?", Status_, Changed_, DefaultLangcode_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByStatusAndAccessAndLogin Get UsersFieldDatas via StatusAndAccessAndLogin
func GetUsersFieldDatasByStatusAndAccessAndLogin(offset int, limit int, Status_ int, Access_ int, Login_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("status = ? and access = ? and login = ?", Status_, Access_, Login_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByStatusAndAccessAndInit Get UsersFieldDatas via StatusAndAccessAndInit
func GetUsersFieldDatasByStatusAndAccessAndInit(offset int, limit int, Status_ int, Access_ int, Init_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("status = ? and access = ? and init = ?", Status_, Access_, Init_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByStatusAndAccessAndDefaultLangcode Get UsersFieldDatas via StatusAndAccessAndDefaultLangcode
func GetUsersFieldDatasByStatusAndAccessAndDefaultLangcode(offset int, limit int, Status_ int, Access_ int, DefaultLangcode_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("status = ? and access = ? and default_langcode = ?", Status_, Access_, DefaultLangcode_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByStatusAndLoginAndInit Get UsersFieldDatas via StatusAndLoginAndInit
func GetUsersFieldDatasByStatusAndLoginAndInit(offset int, limit int, Status_ int, Login_ int, Init_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("status = ? and login = ? and init = ?", Status_, Login_, Init_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByStatusAndLoginAndDefaultLangcode Get UsersFieldDatas via StatusAndLoginAndDefaultLangcode
func GetUsersFieldDatasByStatusAndLoginAndDefaultLangcode(offset int, limit int, Status_ int, Login_ int, DefaultLangcode_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("status = ? and login = ? and default_langcode = ?", Status_, Login_, DefaultLangcode_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByStatusAndInitAndDefaultLangcode Get UsersFieldDatas via StatusAndInitAndDefaultLangcode
func GetUsersFieldDatasByStatusAndInitAndDefaultLangcode(offset int, limit int, Status_ int, Init_ string, DefaultLangcode_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("status = ? and init = ? and default_langcode = ?", Status_, Init_, DefaultLangcode_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByCreatedAndChangedAndAccess Get UsersFieldDatas via CreatedAndChangedAndAccess
func GetUsersFieldDatasByCreatedAndChangedAndAccess(offset int, limit int, Created_ int, Changed_ int, Access_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("created = ? and changed = ? and access = ?", Created_, Changed_, Access_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByCreatedAndChangedAndLogin Get UsersFieldDatas via CreatedAndChangedAndLogin
func GetUsersFieldDatasByCreatedAndChangedAndLogin(offset int, limit int, Created_ int, Changed_ int, Login_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("created = ? and changed = ? and login = ?", Created_, Changed_, Login_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByCreatedAndChangedAndInit Get UsersFieldDatas via CreatedAndChangedAndInit
func GetUsersFieldDatasByCreatedAndChangedAndInit(offset int, limit int, Created_ int, Changed_ int, Init_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("created = ? and changed = ? and init = ?", Created_, Changed_, Init_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByCreatedAndChangedAndDefaultLangcode Get UsersFieldDatas via CreatedAndChangedAndDefaultLangcode
func GetUsersFieldDatasByCreatedAndChangedAndDefaultLangcode(offset int, limit int, Created_ int, Changed_ int, DefaultLangcode_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("created = ? and changed = ? and default_langcode = ?", Created_, Changed_, DefaultLangcode_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByCreatedAndAccessAndLogin Get UsersFieldDatas via CreatedAndAccessAndLogin
func GetUsersFieldDatasByCreatedAndAccessAndLogin(offset int, limit int, Created_ int, Access_ int, Login_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("created = ? and access = ? and login = ?", Created_, Access_, Login_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByCreatedAndAccessAndInit Get UsersFieldDatas via CreatedAndAccessAndInit
func GetUsersFieldDatasByCreatedAndAccessAndInit(offset int, limit int, Created_ int, Access_ int, Init_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("created = ? and access = ? and init = ?", Created_, Access_, Init_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByCreatedAndAccessAndDefaultLangcode Get UsersFieldDatas via CreatedAndAccessAndDefaultLangcode
func GetUsersFieldDatasByCreatedAndAccessAndDefaultLangcode(offset int, limit int, Created_ int, Access_ int, DefaultLangcode_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("created = ? and access = ? and default_langcode = ?", Created_, Access_, DefaultLangcode_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByCreatedAndLoginAndInit Get UsersFieldDatas via CreatedAndLoginAndInit
func GetUsersFieldDatasByCreatedAndLoginAndInit(offset int, limit int, Created_ int, Login_ int, Init_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("created = ? and login = ? and init = ?", Created_, Login_, Init_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByCreatedAndLoginAndDefaultLangcode Get UsersFieldDatas via CreatedAndLoginAndDefaultLangcode
func GetUsersFieldDatasByCreatedAndLoginAndDefaultLangcode(offset int, limit int, Created_ int, Login_ int, DefaultLangcode_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("created = ? and login = ? and default_langcode = ?", Created_, Login_, DefaultLangcode_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByCreatedAndInitAndDefaultLangcode Get UsersFieldDatas via CreatedAndInitAndDefaultLangcode
func GetUsersFieldDatasByCreatedAndInitAndDefaultLangcode(offset int, limit int, Created_ int, Init_ string, DefaultLangcode_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("created = ? and init = ? and default_langcode = ?", Created_, Init_, DefaultLangcode_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByChangedAndAccessAndLogin Get UsersFieldDatas via ChangedAndAccessAndLogin
func GetUsersFieldDatasByChangedAndAccessAndLogin(offset int, limit int, Changed_ int, Access_ int, Login_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("changed = ? and access = ? and login = ?", Changed_, Access_, Login_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByChangedAndAccessAndInit Get UsersFieldDatas via ChangedAndAccessAndInit
func GetUsersFieldDatasByChangedAndAccessAndInit(offset int, limit int, Changed_ int, Access_ int, Init_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("changed = ? and access = ? and init = ?", Changed_, Access_, Init_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByChangedAndAccessAndDefaultLangcode Get UsersFieldDatas via ChangedAndAccessAndDefaultLangcode
func GetUsersFieldDatasByChangedAndAccessAndDefaultLangcode(offset int, limit int, Changed_ int, Access_ int, DefaultLangcode_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("changed = ? and access = ? and default_langcode = ?", Changed_, Access_, DefaultLangcode_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByChangedAndLoginAndInit Get UsersFieldDatas via ChangedAndLoginAndInit
func GetUsersFieldDatasByChangedAndLoginAndInit(offset int, limit int, Changed_ int, Login_ int, Init_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("changed = ? and login = ? and init = ?", Changed_, Login_, Init_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByChangedAndLoginAndDefaultLangcode Get UsersFieldDatas via ChangedAndLoginAndDefaultLangcode
func GetUsersFieldDatasByChangedAndLoginAndDefaultLangcode(offset int, limit int, Changed_ int, Login_ int, DefaultLangcode_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("changed = ? and login = ? and default_langcode = ?", Changed_, Login_, DefaultLangcode_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByChangedAndInitAndDefaultLangcode Get UsersFieldDatas via ChangedAndInitAndDefaultLangcode
func GetUsersFieldDatasByChangedAndInitAndDefaultLangcode(offset int, limit int, Changed_ int, Init_ string, DefaultLangcode_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("changed = ? and init = ? and default_langcode = ?", Changed_, Init_, DefaultLangcode_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByAccessAndLoginAndInit Get UsersFieldDatas via AccessAndLoginAndInit
func GetUsersFieldDatasByAccessAndLoginAndInit(offset int, limit int, Access_ int, Login_ int, Init_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("access = ? and login = ? and init = ?", Access_, Login_, Init_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByAccessAndLoginAndDefaultLangcode Get UsersFieldDatas via AccessAndLoginAndDefaultLangcode
func GetUsersFieldDatasByAccessAndLoginAndDefaultLangcode(offset int, limit int, Access_ int, Login_ int, DefaultLangcode_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("access = ? and login = ? and default_langcode = ?", Access_, Login_, DefaultLangcode_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByAccessAndInitAndDefaultLangcode Get UsersFieldDatas via AccessAndInitAndDefaultLangcode
func GetUsersFieldDatasByAccessAndInitAndDefaultLangcode(offset int, limit int, Access_ int, Init_ string, DefaultLangcode_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("access = ? and init = ? and default_langcode = ?", Access_, Init_, DefaultLangcode_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByLoginAndInitAndDefaultLangcode Get UsersFieldDatas via LoginAndInitAndDefaultLangcode
func GetUsersFieldDatasByLoginAndInitAndDefaultLangcode(offset int, limit int, Login_ int, Init_ string, DefaultLangcode_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("login = ? and init = ? and default_langcode = ?", Login_, Init_, DefaultLangcode_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByUidAndLangcode Get UsersFieldDatas via UidAndLangcode
func GetUsersFieldDatasByUidAndLangcode(offset int, limit int, Uid_ int64, Langcode_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("uid = ? and langcode = ?", Uid_, Langcode_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByUidAndPreferredLangcode Get UsersFieldDatas via UidAndPreferredLangcode
func GetUsersFieldDatasByUidAndPreferredLangcode(offset int, limit int, Uid_ int64, PreferredLangcode_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("uid = ? and preferred_langcode = ?", Uid_, PreferredLangcode_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByUidAndPreferredAdminLangcode Get UsersFieldDatas via UidAndPreferredAdminLangcode
func GetUsersFieldDatasByUidAndPreferredAdminLangcode(offset int, limit int, Uid_ int64, PreferredAdminLangcode_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("uid = ? and preferred_admin_langcode = ?", Uid_, PreferredAdminLangcode_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByUidAndName Get UsersFieldDatas via UidAndName
func GetUsersFieldDatasByUidAndName(offset int, limit int, Uid_ int64, Name_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("uid = ? and name = ?", Uid_, Name_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByUidAndPass Get UsersFieldDatas via UidAndPass
func GetUsersFieldDatasByUidAndPass(offset int, limit int, Uid_ int64, Pass_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("uid = ? and pass = ?", Uid_, Pass_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByUidAndMail Get UsersFieldDatas via UidAndMail
func GetUsersFieldDatasByUidAndMail(offset int, limit int, Uid_ int64, Mail_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("uid = ? and mail = ?", Uid_, Mail_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByUidAndTimezone Get UsersFieldDatas via UidAndTimezone
func GetUsersFieldDatasByUidAndTimezone(offset int, limit int, Uid_ int64, Timezone_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("uid = ? and timezone = ?", Uid_, Timezone_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByUidAndStatus Get UsersFieldDatas via UidAndStatus
func GetUsersFieldDatasByUidAndStatus(offset int, limit int, Uid_ int64, Status_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("uid = ? and status = ?", Uid_, Status_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByUidAndCreated Get UsersFieldDatas via UidAndCreated
func GetUsersFieldDatasByUidAndCreated(offset int, limit int, Uid_ int64, Created_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("uid = ? and created = ?", Uid_, Created_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByUidAndChanged Get UsersFieldDatas via UidAndChanged
func GetUsersFieldDatasByUidAndChanged(offset int, limit int, Uid_ int64, Changed_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("uid = ? and changed = ?", Uid_, Changed_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByUidAndAccess Get UsersFieldDatas via UidAndAccess
func GetUsersFieldDatasByUidAndAccess(offset int, limit int, Uid_ int64, Access_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("uid = ? and access = ?", Uid_, Access_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByUidAndLogin Get UsersFieldDatas via UidAndLogin
func GetUsersFieldDatasByUidAndLogin(offset int, limit int, Uid_ int64, Login_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("uid = ? and login = ?", Uid_, Login_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByUidAndInit Get UsersFieldDatas via UidAndInit
func GetUsersFieldDatasByUidAndInit(offset int, limit int, Uid_ int64, Init_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("uid = ? and init = ?", Uid_, Init_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByUidAndDefaultLangcode Get UsersFieldDatas via UidAndDefaultLangcode
func GetUsersFieldDatasByUidAndDefaultLangcode(offset int, limit int, Uid_ int64, DefaultLangcode_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("uid = ? and default_langcode = ?", Uid_, DefaultLangcode_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByLangcodeAndPreferredLangcode Get UsersFieldDatas via LangcodeAndPreferredLangcode
func GetUsersFieldDatasByLangcodeAndPreferredLangcode(offset int, limit int, Langcode_ string, PreferredLangcode_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("langcode = ? and preferred_langcode = ?", Langcode_, PreferredLangcode_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByLangcodeAndPreferredAdminLangcode Get UsersFieldDatas via LangcodeAndPreferredAdminLangcode
func GetUsersFieldDatasByLangcodeAndPreferredAdminLangcode(offset int, limit int, Langcode_ string, PreferredAdminLangcode_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("langcode = ? and preferred_admin_langcode = ?", Langcode_, PreferredAdminLangcode_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByLangcodeAndName Get UsersFieldDatas via LangcodeAndName
func GetUsersFieldDatasByLangcodeAndName(offset int, limit int, Langcode_ string, Name_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("langcode = ? and name = ?", Langcode_, Name_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByLangcodeAndPass Get UsersFieldDatas via LangcodeAndPass
func GetUsersFieldDatasByLangcodeAndPass(offset int, limit int, Langcode_ string, Pass_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("langcode = ? and pass = ?", Langcode_, Pass_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByLangcodeAndMail Get UsersFieldDatas via LangcodeAndMail
func GetUsersFieldDatasByLangcodeAndMail(offset int, limit int, Langcode_ string, Mail_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("langcode = ? and mail = ?", Langcode_, Mail_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByLangcodeAndTimezone Get UsersFieldDatas via LangcodeAndTimezone
func GetUsersFieldDatasByLangcodeAndTimezone(offset int, limit int, Langcode_ string, Timezone_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("langcode = ? and timezone = ?", Langcode_, Timezone_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByLangcodeAndStatus Get UsersFieldDatas via LangcodeAndStatus
func GetUsersFieldDatasByLangcodeAndStatus(offset int, limit int, Langcode_ string, Status_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("langcode = ? and status = ?", Langcode_, Status_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByLangcodeAndCreated Get UsersFieldDatas via LangcodeAndCreated
func GetUsersFieldDatasByLangcodeAndCreated(offset int, limit int, Langcode_ string, Created_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("langcode = ? and created = ?", Langcode_, Created_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByLangcodeAndChanged Get UsersFieldDatas via LangcodeAndChanged
func GetUsersFieldDatasByLangcodeAndChanged(offset int, limit int, Langcode_ string, Changed_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("langcode = ? and changed = ?", Langcode_, Changed_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByLangcodeAndAccess Get UsersFieldDatas via LangcodeAndAccess
func GetUsersFieldDatasByLangcodeAndAccess(offset int, limit int, Langcode_ string, Access_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("langcode = ? and access = ?", Langcode_, Access_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByLangcodeAndLogin Get UsersFieldDatas via LangcodeAndLogin
func GetUsersFieldDatasByLangcodeAndLogin(offset int, limit int, Langcode_ string, Login_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("langcode = ? and login = ?", Langcode_, Login_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByLangcodeAndInit Get UsersFieldDatas via LangcodeAndInit
func GetUsersFieldDatasByLangcodeAndInit(offset int, limit int, Langcode_ string, Init_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("langcode = ? and init = ?", Langcode_, Init_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByLangcodeAndDefaultLangcode Get UsersFieldDatas via LangcodeAndDefaultLangcode
func GetUsersFieldDatasByLangcodeAndDefaultLangcode(offset int, limit int, Langcode_ string, DefaultLangcode_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("langcode = ? and default_langcode = ?", Langcode_, DefaultLangcode_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredLangcodeAndPreferredAdminLangcode Get UsersFieldDatas via PreferredLangcodeAndPreferredAdminLangcode
func GetUsersFieldDatasByPreferredLangcodeAndPreferredAdminLangcode(offset int, limit int, PreferredLangcode_ string, PreferredAdminLangcode_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_langcode = ? and preferred_admin_langcode = ?", PreferredLangcode_, PreferredAdminLangcode_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredLangcodeAndName Get UsersFieldDatas via PreferredLangcodeAndName
func GetUsersFieldDatasByPreferredLangcodeAndName(offset int, limit int, PreferredLangcode_ string, Name_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_langcode = ? and name = ?", PreferredLangcode_, Name_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredLangcodeAndPass Get UsersFieldDatas via PreferredLangcodeAndPass
func GetUsersFieldDatasByPreferredLangcodeAndPass(offset int, limit int, PreferredLangcode_ string, Pass_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_langcode = ? and pass = ?", PreferredLangcode_, Pass_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredLangcodeAndMail Get UsersFieldDatas via PreferredLangcodeAndMail
func GetUsersFieldDatasByPreferredLangcodeAndMail(offset int, limit int, PreferredLangcode_ string, Mail_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_langcode = ? and mail = ?", PreferredLangcode_, Mail_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredLangcodeAndTimezone Get UsersFieldDatas via PreferredLangcodeAndTimezone
func GetUsersFieldDatasByPreferredLangcodeAndTimezone(offset int, limit int, PreferredLangcode_ string, Timezone_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_langcode = ? and timezone = ?", PreferredLangcode_, Timezone_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredLangcodeAndStatus Get UsersFieldDatas via PreferredLangcodeAndStatus
func GetUsersFieldDatasByPreferredLangcodeAndStatus(offset int, limit int, PreferredLangcode_ string, Status_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_langcode = ? and status = ?", PreferredLangcode_, Status_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredLangcodeAndCreated Get UsersFieldDatas via PreferredLangcodeAndCreated
func GetUsersFieldDatasByPreferredLangcodeAndCreated(offset int, limit int, PreferredLangcode_ string, Created_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_langcode = ? and created = ?", PreferredLangcode_, Created_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredLangcodeAndChanged Get UsersFieldDatas via PreferredLangcodeAndChanged
func GetUsersFieldDatasByPreferredLangcodeAndChanged(offset int, limit int, PreferredLangcode_ string, Changed_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_langcode = ? and changed = ?", PreferredLangcode_, Changed_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredLangcodeAndAccess Get UsersFieldDatas via PreferredLangcodeAndAccess
func GetUsersFieldDatasByPreferredLangcodeAndAccess(offset int, limit int, PreferredLangcode_ string, Access_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_langcode = ? and access = ?", PreferredLangcode_, Access_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredLangcodeAndLogin Get UsersFieldDatas via PreferredLangcodeAndLogin
func GetUsersFieldDatasByPreferredLangcodeAndLogin(offset int, limit int, PreferredLangcode_ string, Login_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_langcode = ? and login = ?", PreferredLangcode_, Login_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredLangcodeAndInit Get UsersFieldDatas via PreferredLangcodeAndInit
func GetUsersFieldDatasByPreferredLangcodeAndInit(offset int, limit int, PreferredLangcode_ string, Init_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_langcode = ? and init = ?", PreferredLangcode_, Init_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredLangcodeAndDefaultLangcode Get UsersFieldDatas via PreferredLangcodeAndDefaultLangcode
func GetUsersFieldDatasByPreferredLangcodeAndDefaultLangcode(offset int, limit int, PreferredLangcode_ string, DefaultLangcode_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_langcode = ? and default_langcode = ?", PreferredLangcode_, DefaultLangcode_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredAdminLangcodeAndName Get UsersFieldDatas via PreferredAdminLangcodeAndName
func GetUsersFieldDatasByPreferredAdminLangcodeAndName(offset int, limit int, PreferredAdminLangcode_ string, Name_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_admin_langcode = ? and name = ?", PreferredAdminLangcode_, Name_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredAdminLangcodeAndPass Get UsersFieldDatas via PreferredAdminLangcodeAndPass
func GetUsersFieldDatasByPreferredAdminLangcodeAndPass(offset int, limit int, PreferredAdminLangcode_ string, Pass_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_admin_langcode = ? and pass = ?", PreferredAdminLangcode_, Pass_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredAdminLangcodeAndMail Get UsersFieldDatas via PreferredAdminLangcodeAndMail
func GetUsersFieldDatasByPreferredAdminLangcodeAndMail(offset int, limit int, PreferredAdminLangcode_ string, Mail_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_admin_langcode = ? and mail = ?", PreferredAdminLangcode_, Mail_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredAdminLangcodeAndTimezone Get UsersFieldDatas via PreferredAdminLangcodeAndTimezone
func GetUsersFieldDatasByPreferredAdminLangcodeAndTimezone(offset int, limit int, PreferredAdminLangcode_ string, Timezone_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_admin_langcode = ? and timezone = ?", PreferredAdminLangcode_, Timezone_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredAdminLangcodeAndStatus Get UsersFieldDatas via PreferredAdminLangcodeAndStatus
func GetUsersFieldDatasByPreferredAdminLangcodeAndStatus(offset int, limit int, PreferredAdminLangcode_ string, Status_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_admin_langcode = ? and status = ?", PreferredAdminLangcode_, Status_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredAdminLangcodeAndCreated Get UsersFieldDatas via PreferredAdminLangcodeAndCreated
func GetUsersFieldDatasByPreferredAdminLangcodeAndCreated(offset int, limit int, PreferredAdminLangcode_ string, Created_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_admin_langcode = ? and created = ?", PreferredAdminLangcode_, Created_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredAdminLangcodeAndChanged Get UsersFieldDatas via PreferredAdminLangcodeAndChanged
func GetUsersFieldDatasByPreferredAdminLangcodeAndChanged(offset int, limit int, PreferredAdminLangcode_ string, Changed_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_admin_langcode = ? and changed = ?", PreferredAdminLangcode_, Changed_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredAdminLangcodeAndAccess Get UsersFieldDatas via PreferredAdminLangcodeAndAccess
func GetUsersFieldDatasByPreferredAdminLangcodeAndAccess(offset int, limit int, PreferredAdminLangcode_ string, Access_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_admin_langcode = ? and access = ?", PreferredAdminLangcode_, Access_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredAdminLangcodeAndLogin Get UsersFieldDatas via PreferredAdminLangcodeAndLogin
func GetUsersFieldDatasByPreferredAdminLangcodeAndLogin(offset int, limit int, PreferredAdminLangcode_ string, Login_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_admin_langcode = ? and login = ?", PreferredAdminLangcode_, Login_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredAdminLangcodeAndInit Get UsersFieldDatas via PreferredAdminLangcodeAndInit
func GetUsersFieldDatasByPreferredAdminLangcodeAndInit(offset int, limit int, PreferredAdminLangcode_ string, Init_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_admin_langcode = ? and init = ?", PreferredAdminLangcode_, Init_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPreferredAdminLangcodeAndDefaultLangcode Get UsersFieldDatas via PreferredAdminLangcodeAndDefaultLangcode
func GetUsersFieldDatasByPreferredAdminLangcodeAndDefaultLangcode(offset int, limit int, PreferredAdminLangcode_ string, DefaultLangcode_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_admin_langcode = ? and default_langcode = ?", PreferredAdminLangcode_, DefaultLangcode_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByNameAndPass Get UsersFieldDatas via NameAndPass
func GetUsersFieldDatasByNameAndPass(offset int, limit int, Name_ string, Pass_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("name = ? and pass = ?", Name_, Pass_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByNameAndMail Get UsersFieldDatas via NameAndMail
func GetUsersFieldDatasByNameAndMail(offset int, limit int, Name_ string, Mail_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("name = ? and mail = ?", Name_, Mail_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByNameAndTimezone Get UsersFieldDatas via NameAndTimezone
func GetUsersFieldDatasByNameAndTimezone(offset int, limit int, Name_ string, Timezone_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("name = ? and timezone = ?", Name_, Timezone_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByNameAndStatus Get UsersFieldDatas via NameAndStatus
func GetUsersFieldDatasByNameAndStatus(offset int, limit int, Name_ string, Status_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("name = ? and status = ?", Name_, Status_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByNameAndCreated Get UsersFieldDatas via NameAndCreated
func GetUsersFieldDatasByNameAndCreated(offset int, limit int, Name_ string, Created_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("name = ? and created = ?", Name_, Created_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByNameAndChanged Get UsersFieldDatas via NameAndChanged
func GetUsersFieldDatasByNameAndChanged(offset int, limit int, Name_ string, Changed_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("name = ? and changed = ?", Name_, Changed_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByNameAndAccess Get UsersFieldDatas via NameAndAccess
func GetUsersFieldDatasByNameAndAccess(offset int, limit int, Name_ string, Access_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("name = ? and access = ?", Name_, Access_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByNameAndLogin Get UsersFieldDatas via NameAndLogin
func GetUsersFieldDatasByNameAndLogin(offset int, limit int, Name_ string, Login_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("name = ? and login = ?", Name_, Login_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByNameAndInit Get UsersFieldDatas via NameAndInit
func GetUsersFieldDatasByNameAndInit(offset int, limit int, Name_ string, Init_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("name = ? and init = ?", Name_, Init_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByNameAndDefaultLangcode Get UsersFieldDatas via NameAndDefaultLangcode
func GetUsersFieldDatasByNameAndDefaultLangcode(offset int, limit int, Name_ string, DefaultLangcode_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("name = ? and default_langcode = ?", Name_, DefaultLangcode_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPassAndMail Get UsersFieldDatas via PassAndMail
func GetUsersFieldDatasByPassAndMail(offset int, limit int, Pass_ string, Mail_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("pass = ? and mail = ?", Pass_, Mail_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPassAndTimezone Get UsersFieldDatas via PassAndTimezone
func GetUsersFieldDatasByPassAndTimezone(offset int, limit int, Pass_ string, Timezone_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("pass = ? and timezone = ?", Pass_, Timezone_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPassAndStatus Get UsersFieldDatas via PassAndStatus
func GetUsersFieldDatasByPassAndStatus(offset int, limit int, Pass_ string, Status_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("pass = ? and status = ?", Pass_, Status_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPassAndCreated Get UsersFieldDatas via PassAndCreated
func GetUsersFieldDatasByPassAndCreated(offset int, limit int, Pass_ string, Created_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("pass = ? and created = ?", Pass_, Created_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPassAndChanged Get UsersFieldDatas via PassAndChanged
func GetUsersFieldDatasByPassAndChanged(offset int, limit int, Pass_ string, Changed_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("pass = ? and changed = ?", Pass_, Changed_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPassAndAccess Get UsersFieldDatas via PassAndAccess
func GetUsersFieldDatasByPassAndAccess(offset int, limit int, Pass_ string, Access_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("pass = ? and access = ?", Pass_, Access_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPassAndLogin Get UsersFieldDatas via PassAndLogin
func GetUsersFieldDatasByPassAndLogin(offset int, limit int, Pass_ string, Login_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("pass = ? and login = ?", Pass_, Login_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPassAndInit Get UsersFieldDatas via PassAndInit
func GetUsersFieldDatasByPassAndInit(offset int, limit int, Pass_ string, Init_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("pass = ? and init = ?", Pass_, Init_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByPassAndDefaultLangcode Get UsersFieldDatas via PassAndDefaultLangcode
func GetUsersFieldDatasByPassAndDefaultLangcode(offset int, limit int, Pass_ string, DefaultLangcode_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("pass = ? and default_langcode = ?", Pass_, DefaultLangcode_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByMailAndTimezone Get UsersFieldDatas via MailAndTimezone
func GetUsersFieldDatasByMailAndTimezone(offset int, limit int, Mail_ string, Timezone_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("mail = ? and timezone = ?", Mail_, Timezone_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByMailAndStatus Get UsersFieldDatas via MailAndStatus
func GetUsersFieldDatasByMailAndStatus(offset int, limit int, Mail_ string, Status_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("mail = ? and status = ?", Mail_, Status_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByMailAndCreated Get UsersFieldDatas via MailAndCreated
func GetUsersFieldDatasByMailAndCreated(offset int, limit int, Mail_ string, Created_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("mail = ? and created = ?", Mail_, Created_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByMailAndChanged Get UsersFieldDatas via MailAndChanged
func GetUsersFieldDatasByMailAndChanged(offset int, limit int, Mail_ string, Changed_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("mail = ? and changed = ?", Mail_, Changed_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByMailAndAccess Get UsersFieldDatas via MailAndAccess
func GetUsersFieldDatasByMailAndAccess(offset int, limit int, Mail_ string, Access_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("mail = ? and access = ?", Mail_, Access_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByMailAndLogin Get UsersFieldDatas via MailAndLogin
func GetUsersFieldDatasByMailAndLogin(offset int, limit int, Mail_ string, Login_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("mail = ? and login = ?", Mail_, Login_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByMailAndInit Get UsersFieldDatas via MailAndInit
func GetUsersFieldDatasByMailAndInit(offset int, limit int, Mail_ string, Init_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("mail = ? and init = ?", Mail_, Init_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByMailAndDefaultLangcode Get UsersFieldDatas via MailAndDefaultLangcode
func GetUsersFieldDatasByMailAndDefaultLangcode(offset int, limit int, Mail_ string, DefaultLangcode_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("mail = ? and default_langcode = ?", Mail_, DefaultLangcode_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByTimezoneAndStatus Get UsersFieldDatas via TimezoneAndStatus
func GetUsersFieldDatasByTimezoneAndStatus(offset int, limit int, Timezone_ string, Status_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("timezone = ? and status = ?", Timezone_, Status_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByTimezoneAndCreated Get UsersFieldDatas via TimezoneAndCreated
func GetUsersFieldDatasByTimezoneAndCreated(offset int, limit int, Timezone_ string, Created_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("timezone = ? and created = ?", Timezone_, Created_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByTimezoneAndChanged Get UsersFieldDatas via TimezoneAndChanged
func GetUsersFieldDatasByTimezoneAndChanged(offset int, limit int, Timezone_ string, Changed_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("timezone = ? and changed = ?", Timezone_, Changed_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByTimezoneAndAccess Get UsersFieldDatas via TimezoneAndAccess
func GetUsersFieldDatasByTimezoneAndAccess(offset int, limit int, Timezone_ string, Access_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("timezone = ? and access = ?", Timezone_, Access_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByTimezoneAndLogin Get UsersFieldDatas via TimezoneAndLogin
func GetUsersFieldDatasByTimezoneAndLogin(offset int, limit int, Timezone_ string, Login_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("timezone = ? and login = ?", Timezone_, Login_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByTimezoneAndInit Get UsersFieldDatas via TimezoneAndInit
func GetUsersFieldDatasByTimezoneAndInit(offset int, limit int, Timezone_ string, Init_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("timezone = ? and init = ?", Timezone_, Init_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByTimezoneAndDefaultLangcode Get UsersFieldDatas via TimezoneAndDefaultLangcode
func GetUsersFieldDatasByTimezoneAndDefaultLangcode(offset int, limit int, Timezone_ string, DefaultLangcode_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("timezone = ? and default_langcode = ?", Timezone_, DefaultLangcode_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByStatusAndCreated Get UsersFieldDatas via StatusAndCreated
func GetUsersFieldDatasByStatusAndCreated(offset int, limit int, Status_ int, Created_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("status = ? and created = ?", Status_, Created_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByStatusAndChanged Get UsersFieldDatas via StatusAndChanged
func GetUsersFieldDatasByStatusAndChanged(offset int, limit int, Status_ int, Changed_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("status = ? and changed = ?", Status_, Changed_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByStatusAndAccess Get UsersFieldDatas via StatusAndAccess
func GetUsersFieldDatasByStatusAndAccess(offset int, limit int, Status_ int, Access_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("status = ? and access = ?", Status_, Access_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByStatusAndLogin Get UsersFieldDatas via StatusAndLogin
func GetUsersFieldDatasByStatusAndLogin(offset int, limit int, Status_ int, Login_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("status = ? and login = ?", Status_, Login_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByStatusAndInit Get UsersFieldDatas via StatusAndInit
func GetUsersFieldDatasByStatusAndInit(offset int, limit int, Status_ int, Init_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("status = ? and init = ?", Status_, Init_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByStatusAndDefaultLangcode Get UsersFieldDatas via StatusAndDefaultLangcode
func GetUsersFieldDatasByStatusAndDefaultLangcode(offset int, limit int, Status_ int, DefaultLangcode_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("status = ? and default_langcode = ?", Status_, DefaultLangcode_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByCreatedAndChanged Get UsersFieldDatas via CreatedAndChanged
func GetUsersFieldDatasByCreatedAndChanged(offset int, limit int, Created_ int, Changed_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("created = ? and changed = ?", Created_, Changed_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByCreatedAndAccess Get UsersFieldDatas via CreatedAndAccess
func GetUsersFieldDatasByCreatedAndAccess(offset int, limit int, Created_ int, Access_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("created = ? and access = ?", Created_, Access_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByCreatedAndLogin Get UsersFieldDatas via CreatedAndLogin
func GetUsersFieldDatasByCreatedAndLogin(offset int, limit int, Created_ int, Login_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("created = ? and login = ?", Created_, Login_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByCreatedAndInit Get UsersFieldDatas via CreatedAndInit
func GetUsersFieldDatasByCreatedAndInit(offset int, limit int, Created_ int, Init_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("created = ? and init = ?", Created_, Init_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByCreatedAndDefaultLangcode Get UsersFieldDatas via CreatedAndDefaultLangcode
func GetUsersFieldDatasByCreatedAndDefaultLangcode(offset int, limit int, Created_ int, DefaultLangcode_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("created = ? and default_langcode = ?", Created_, DefaultLangcode_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByChangedAndAccess Get UsersFieldDatas via ChangedAndAccess
func GetUsersFieldDatasByChangedAndAccess(offset int, limit int, Changed_ int, Access_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("changed = ? and access = ?", Changed_, Access_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByChangedAndLogin Get UsersFieldDatas via ChangedAndLogin
func GetUsersFieldDatasByChangedAndLogin(offset int, limit int, Changed_ int, Login_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("changed = ? and login = ?", Changed_, Login_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByChangedAndInit Get UsersFieldDatas via ChangedAndInit
func GetUsersFieldDatasByChangedAndInit(offset int, limit int, Changed_ int, Init_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("changed = ? and init = ?", Changed_, Init_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByChangedAndDefaultLangcode Get UsersFieldDatas via ChangedAndDefaultLangcode
func GetUsersFieldDatasByChangedAndDefaultLangcode(offset int, limit int, Changed_ int, DefaultLangcode_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("changed = ? and default_langcode = ?", Changed_, DefaultLangcode_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByAccessAndLogin Get UsersFieldDatas via AccessAndLogin
func GetUsersFieldDatasByAccessAndLogin(offset int, limit int, Access_ int, Login_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("access = ? and login = ?", Access_, Login_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByAccessAndInit Get UsersFieldDatas via AccessAndInit
func GetUsersFieldDatasByAccessAndInit(offset int, limit int, Access_ int, Init_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("access = ? and init = ?", Access_, Init_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByAccessAndDefaultLangcode Get UsersFieldDatas via AccessAndDefaultLangcode
func GetUsersFieldDatasByAccessAndDefaultLangcode(offset int, limit int, Access_ int, DefaultLangcode_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("access = ? and default_langcode = ?", Access_, DefaultLangcode_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByLoginAndInit Get UsersFieldDatas via LoginAndInit
func GetUsersFieldDatasByLoginAndInit(offset int, limit int, Login_ int, Init_ string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("login = ? and init = ?", Login_, Init_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByLoginAndDefaultLangcode Get UsersFieldDatas via LoginAndDefaultLangcode
func GetUsersFieldDatasByLoginAndDefaultLangcode(offset int, limit int, Login_ int, DefaultLangcode_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("login = ? and default_langcode = ?", Login_, DefaultLangcode_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasByInitAndDefaultLangcode Get UsersFieldDatas via InitAndDefaultLangcode
func GetUsersFieldDatasByInitAndDefaultLangcode(offset int, limit int, Init_ string, DefaultLangcode_ int) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("init = ? and default_langcode = ?", Init_, DefaultLangcode_).Limit(limit, offset).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatas Get UsersFieldDatas via field
func GetUsersFieldDatas(offset int, limit int, field string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Limit(limit, offset).Desc(field).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasViaUid Get UsersFieldDatas via Uid
func GetUsersFieldDatasViaUid(offset int, limit int, Uid_ int64, field string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("uid = ?", Uid_).Limit(limit, offset).Desc(field).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasViaLangcode Get UsersFieldDatas via Langcode
func GetUsersFieldDatasViaLangcode(offset int, limit int, Langcode_ string, field string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("langcode = ?", Langcode_).Limit(limit, offset).Desc(field).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasViaPreferredLangcode Get UsersFieldDatas via PreferredLangcode
func GetUsersFieldDatasViaPreferredLangcode(offset int, limit int, PreferredLangcode_ string, field string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_langcode = ?", PreferredLangcode_).Limit(limit, offset).Desc(field).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasViaPreferredAdminLangcode Get UsersFieldDatas via PreferredAdminLangcode
func GetUsersFieldDatasViaPreferredAdminLangcode(offset int, limit int, PreferredAdminLangcode_ string, field string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("preferred_admin_langcode = ?", PreferredAdminLangcode_).Limit(limit, offset).Desc(field).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasViaName Get UsersFieldDatas via Name
func GetUsersFieldDatasViaName(offset int, limit int, Name_ string, field string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("name = ?", Name_).Limit(limit, offset).Desc(field).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasViaPass Get UsersFieldDatas via Pass
func GetUsersFieldDatasViaPass(offset int, limit int, Pass_ string, field string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("pass = ?", Pass_).Limit(limit, offset).Desc(field).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasViaMail Get UsersFieldDatas via Mail
func GetUsersFieldDatasViaMail(offset int, limit int, Mail_ string, field string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("mail = ?", Mail_).Limit(limit, offset).Desc(field).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasViaTimezone Get UsersFieldDatas via Timezone
func GetUsersFieldDatasViaTimezone(offset int, limit int, Timezone_ string, field string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("timezone = ?", Timezone_).Limit(limit, offset).Desc(field).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasViaStatus Get UsersFieldDatas via Status
func GetUsersFieldDatasViaStatus(offset int, limit int, Status_ int, field string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("status = ?", Status_).Limit(limit, offset).Desc(field).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasViaCreated Get UsersFieldDatas via Created
func GetUsersFieldDatasViaCreated(offset int, limit int, Created_ int, field string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("created = ?", Created_).Limit(limit, offset).Desc(field).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasViaChanged Get UsersFieldDatas via Changed
func GetUsersFieldDatasViaChanged(offset int, limit int, Changed_ int, field string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("changed = ?", Changed_).Limit(limit, offset).Desc(field).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasViaAccess Get UsersFieldDatas via Access
func GetUsersFieldDatasViaAccess(offset int, limit int, Access_ int, field string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("access = ?", Access_).Limit(limit, offset).Desc(field).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasViaLogin Get UsersFieldDatas via Login
func GetUsersFieldDatasViaLogin(offset int, limit int, Login_ int, field string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("login = ?", Login_).Limit(limit, offset).Desc(field).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasViaInit Get UsersFieldDatas via Init
func GetUsersFieldDatasViaInit(offset int, limit int, Init_ string, field string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("init = ?", Init_).Limit(limit, offset).Desc(field).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// GetUsersFieldDatasViaDefaultLangcode Get UsersFieldDatas via DefaultLangcode
func GetUsersFieldDatasViaDefaultLangcode(offset int, limit int, DefaultLangcode_ int, field string) (*[]*UsersFieldData, error) {
	var _UsersFieldData = new([]*UsersFieldData)
	err := Engine.Table("users_field_data").Where("default_langcode = ?", DefaultLangcode_).Limit(limit, offset).Desc(field).Find(_UsersFieldData)
	return _UsersFieldData, err
}

// HasUsersFieldDataViaUid Has UsersFieldData via Uid
func HasUsersFieldDataViaUid(iUid int64) bool {
	if has, err := Engine.Where("uid = ?", iUid).Get(new(UsersFieldData)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasUsersFieldDataViaLangcode Has UsersFieldData via Langcode
func HasUsersFieldDataViaLangcode(iLangcode string) bool {
	if has, err := Engine.Where("langcode = ?", iLangcode).Get(new(UsersFieldData)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasUsersFieldDataViaPreferredLangcode Has UsersFieldData via PreferredLangcode
func HasUsersFieldDataViaPreferredLangcode(iPreferredLangcode string) bool {
	if has, err := Engine.Where("preferred_langcode = ?", iPreferredLangcode).Get(new(UsersFieldData)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasUsersFieldDataViaPreferredAdminLangcode Has UsersFieldData via PreferredAdminLangcode
func HasUsersFieldDataViaPreferredAdminLangcode(iPreferredAdminLangcode string) bool {
	if has, err := Engine.Where("preferred_admin_langcode = ?", iPreferredAdminLangcode).Get(new(UsersFieldData)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasUsersFieldDataViaName Has UsersFieldData via Name
func HasUsersFieldDataViaName(iName string) bool {
	if has, err := Engine.Where("name = ?", iName).Get(new(UsersFieldData)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasUsersFieldDataViaPass Has UsersFieldData via Pass
func HasUsersFieldDataViaPass(iPass string) bool {
	if has, err := Engine.Where("pass = ?", iPass).Get(new(UsersFieldData)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasUsersFieldDataViaMail Has UsersFieldData via Mail
func HasUsersFieldDataViaMail(iMail string) bool {
	if has, err := Engine.Where("mail = ?", iMail).Get(new(UsersFieldData)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasUsersFieldDataViaTimezone Has UsersFieldData via Timezone
func HasUsersFieldDataViaTimezone(iTimezone string) bool {
	if has, err := Engine.Where("timezone = ?", iTimezone).Get(new(UsersFieldData)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasUsersFieldDataViaStatus Has UsersFieldData via Status
func HasUsersFieldDataViaStatus(iStatus int) bool {
	if has, err := Engine.Where("status = ?", iStatus).Get(new(UsersFieldData)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasUsersFieldDataViaCreated Has UsersFieldData via Created
func HasUsersFieldDataViaCreated(iCreated int) bool {
	if has, err := Engine.Where("created = ?", iCreated).Get(new(UsersFieldData)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasUsersFieldDataViaChanged Has UsersFieldData via Changed
func HasUsersFieldDataViaChanged(iChanged int) bool {
	if has, err := Engine.Where("changed = ?", iChanged).Get(new(UsersFieldData)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasUsersFieldDataViaAccess Has UsersFieldData via Access
func HasUsersFieldDataViaAccess(iAccess int) bool {
	if has, err := Engine.Where("access = ?", iAccess).Get(new(UsersFieldData)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasUsersFieldDataViaLogin Has UsersFieldData via Login
func HasUsersFieldDataViaLogin(iLogin int) bool {
	if has, err := Engine.Where("login = ?", iLogin).Get(new(UsersFieldData)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasUsersFieldDataViaInit Has UsersFieldData via Init
func HasUsersFieldDataViaInit(iInit string) bool {
	if has, err := Engine.Where("init = ?", iInit).Get(new(UsersFieldData)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasUsersFieldDataViaDefaultLangcode Has UsersFieldData via DefaultLangcode
func HasUsersFieldDataViaDefaultLangcode(iDefaultLangcode int) bool {
	if has, err := Engine.Where("default_langcode = ?", iDefaultLangcode).Get(new(UsersFieldData)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// GetUsersFieldDataViaUid Get UsersFieldData via Uid
func GetUsersFieldDataViaUid(iUid int64) (*UsersFieldData, error) {
	var _UsersFieldData = &UsersFieldData{Uid: iUid}
	has, err := Engine.Get(_UsersFieldData)
	if has {
		return _UsersFieldData, err
	} else {
		return nil, err
	}
}

// GetUsersFieldDataViaLangcode Get UsersFieldData via Langcode
func GetUsersFieldDataViaLangcode(iLangcode string) (*UsersFieldData, error) {
	var _UsersFieldData = &UsersFieldData{Langcode: iLangcode}
	has, err := Engine.Get(_UsersFieldData)
	if has {
		return _UsersFieldData, err
	} else {
		return nil, err
	}
}

// GetUsersFieldDataViaPreferredLangcode Get UsersFieldData via PreferredLangcode
func GetUsersFieldDataViaPreferredLangcode(iPreferredLangcode string) (*UsersFieldData, error) {
	var _UsersFieldData = &UsersFieldData{PreferredLangcode: iPreferredLangcode}
	has, err := Engine.Get(_UsersFieldData)
	if has {
		return _UsersFieldData, err
	} else {
		return nil, err
	}
}

// GetUsersFieldDataViaPreferredAdminLangcode Get UsersFieldData via PreferredAdminLangcode
func GetUsersFieldDataViaPreferredAdminLangcode(iPreferredAdminLangcode string) (*UsersFieldData, error) {
	var _UsersFieldData = &UsersFieldData{PreferredAdminLangcode: iPreferredAdminLangcode}
	has, err := Engine.Get(_UsersFieldData)
	if has {
		return _UsersFieldData, err
	} else {
		return nil, err
	}
}

// GetUsersFieldDataViaName Get UsersFieldData via Name
func GetUsersFieldDataViaName(iName string) (*UsersFieldData, error) {
	var _UsersFieldData = &UsersFieldData{Name: iName}
	has, err := Engine.Get(_UsersFieldData)
	if has {
		return _UsersFieldData, err
	} else {
		return nil, err
	}
}

// GetUsersFieldDataViaPass Get UsersFieldData via Pass
func GetUsersFieldDataViaPass(iPass string) (*UsersFieldData, error) {
	var _UsersFieldData = &UsersFieldData{Pass: iPass}
	has, err := Engine.Get(_UsersFieldData)
	if has {
		return _UsersFieldData, err
	} else {
		return nil, err
	}
}

// GetUsersFieldDataViaMail Get UsersFieldData via Mail
func GetUsersFieldDataViaMail(iMail string) (*UsersFieldData, error) {
	var _UsersFieldData = &UsersFieldData{Mail: iMail}
	has, err := Engine.Get(_UsersFieldData)
	if has {
		return _UsersFieldData, err
	} else {
		return nil, err
	}
}

// GetUsersFieldDataViaTimezone Get UsersFieldData via Timezone
func GetUsersFieldDataViaTimezone(iTimezone string) (*UsersFieldData, error) {
	var _UsersFieldData = &UsersFieldData{Timezone: iTimezone}
	has, err := Engine.Get(_UsersFieldData)
	if has {
		return _UsersFieldData, err
	} else {
		return nil, err
	}
}

// GetUsersFieldDataViaStatus Get UsersFieldData via Status
func GetUsersFieldDataViaStatus(iStatus int) (*UsersFieldData, error) {
	var _UsersFieldData = &UsersFieldData{Status: iStatus}
	has, err := Engine.Get(_UsersFieldData)
	if has {
		return _UsersFieldData, err
	} else {
		return nil, err
	}
}

// GetUsersFieldDataViaCreated Get UsersFieldData via Created
func GetUsersFieldDataViaCreated(iCreated int) (*UsersFieldData, error) {
	var _UsersFieldData = &UsersFieldData{Created: iCreated}
	has, err := Engine.Get(_UsersFieldData)
	if has {
		return _UsersFieldData, err
	} else {
		return nil, err
	}
}

// GetUsersFieldDataViaChanged Get UsersFieldData via Changed
func GetUsersFieldDataViaChanged(iChanged int) (*UsersFieldData, error) {
	var _UsersFieldData = &UsersFieldData{Changed: iChanged}
	has, err := Engine.Get(_UsersFieldData)
	if has {
		return _UsersFieldData, err
	} else {
		return nil, err
	}
}

// GetUsersFieldDataViaAccess Get UsersFieldData via Access
func GetUsersFieldDataViaAccess(iAccess int) (*UsersFieldData, error) {
	var _UsersFieldData = &UsersFieldData{Access: iAccess}
	has, err := Engine.Get(_UsersFieldData)
	if has {
		return _UsersFieldData, err
	} else {
		return nil, err
	}
}

// GetUsersFieldDataViaLogin Get UsersFieldData via Login
func GetUsersFieldDataViaLogin(iLogin int) (*UsersFieldData, error) {
	var _UsersFieldData = &UsersFieldData{Login: iLogin}
	has, err := Engine.Get(_UsersFieldData)
	if has {
		return _UsersFieldData, err
	} else {
		return nil, err
	}
}

// GetUsersFieldDataViaInit Get UsersFieldData via Init
func GetUsersFieldDataViaInit(iInit string) (*UsersFieldData, error) {
	var _UsersFieldData = &UsersFieldData{Init: iInit}
	has, err := Engine.Get(_UsersFieldData)
	if has {
		return _UsersFieldData, err
	} else {
		return nil, err
	}
}

// GetUsersFieldDataViaDefaultLangcode Get UsersFieldData via DefaultLangcode
func GetUsersFieldDataViaDefaultLangcode(iDefaultLangcode int) (*UsersFieldData, error) {
	var _UsersFieldData = &UsersFieldData{DefaultLangcode: iDefaultLangcode}
	has, err := Engine.Get(_UsersFieldData)
	if has {
		return _UsersFieldData, err
	} else {
		return nil, err
	}
}

// SetUsersFieldDataViaUid Set UsersFieldData via Uid
func SetUsersFieldDataViaUid(iUid int64, users_field_data *UsersFieldData) (int64, error) {
	users_field_data.Uid = iUid
	return Engine.Insert(users_field_data)
}

// SetUsersFieldDataViaLangcode Set UsersFieldData via Langcode
func SetUsersFieldDataViaLangcode(iLangcode string, users_field_data *UsersFieldData) (int64, error) {
	users_field_data.Langcode = iLangcode
	return Engine.Insert(users_field_data)
}

// SetUsersFieldDataViaPreferredLangcode Set UsersFieldData via PreferredLangcode
func SetUsersFieldDataViaPreferredLangcode(iPreferredLangcode string, users_field_data *UsersFieldData) (int64, error) {
	users_field_data.PreferredLangcode = iPreferredLangcode
	return Engine.Insert(users_field_data)
}

// SetUsersFieldDataViaPreferredAdminLangcode Set UsersFieldData via PreferredAdminLangcode
func SetUsersFieldDataViaPreferredAdminLangcode(iPreferredAdminLangcode string, users_field_data *UsersFieldData) (int64, error) {
	users_field_data.PreferredAdminLangcode = iPreferredAdminLangcode
	return Engine.Insert(users_field_data)
}

// SetUsersFieldDataViaName Set UsersFieldData via Name
func SetUsersFieldDataViaName(iName string, users_field_data *UsersFieldData) (int64, error) {
	users_field_data.Name = iName
	return Engine.Insert(users_field_data)
}

// SetUsersFieldDataViaPass Set UsersFieldData via Pass
func SetUsersFieldDataViaPass(iPass string, users_field_data *UsersFieldData) (int64, error) {
	users_field_data.Pass = iPass
	return Engine.Insert(users_field_data)
}

// SetUsersFieldDataViaMail Set UsersFieldData via Mail
func SetUsersFieldDataViaMail(iMail string, users_field_data *UsersFieldData) (int64, error) {
	users_field_data.Mail = iMail
	return Engine.Insert(users_field_data)
}

// SetUsersFieldDataViaTimezone Set UsersFieldData via Timezone
func SetUsersFieldDataViaTimezone(iTimezone string, users_field_data *UsersFieldData) (int64, error) {
	users_field_data.Timezone = iTimezone
	return Engine.Insert(users_field_data)
}

// SetUsersFieldDataViaStatus Set UsersFieldData via Status
func SetUsersFieldDataViaStatus(iStatus int, users_field_data *UsersFieldData) (int64, error) {
	users_field_data.Status = iStatus
	return Engine.Insert(users_field_data)
}

// SetUsersFieldDataViaCreated Set UsersFieldData via Created
func SetUsersFieldDataViaCreated(iCreated int, users_field_data *UsersFieldData) (int64, error) {
	users_field_data.Created = iCreated
	return Engine.Insert(users_field_data)
}

// SetUsersFieldDataViaChanged Set UsersFieldData via Changed
func SetUsersFieldDataViaChanged(iChanged int, users_field_data *UsersFieldData) (int64, error) {
	users_field_data.Changed = iChanged
	return Engine.Insert(users_field_data)
}

// SetUsersFieldDataViaAccess Set UsersFieldData via Access
func SetUsersFieldDataViaAccess(iAccess int, users_field_data *UsersFieldData) (int64, error) {
	users_field_data.Access = iAccess
	return Engine.Insert(users_field_data)
}

// SetUsersFieldDataViaLogin Set UsersFieldData via Login
func SetUsersFieldDataViaLogin(iLogin int, users_field_data *UsersFieldData) (int64, error) {
	users_field_data.Login = iLogin
	return Engine.Insert(users_field_data)
}

// SetUsersFieldDataViaInit Set UsersFieldData via Init
func SetUsersFieldDataViaInit(iInit string, users_field_data *UsersFieldData) (int64, error) {
	users_field_data.Init = iInit
	return Engine.Insert(users_field_data)
}

// SetUsersFieldDataViaDefaultLangcode Set UsersFieldData via DefaultLangcode
func SetUsersFieldDataViaDefaultLangcode(iDefaultLangcode int, users_field_data *UsersFieldData) (int64, error) {
	users_field_data.DefaultLangcode = iDefaultLangcode
	return Engine.Insert(users_field_data)
}

// AddUsersFieldData Add UsersFieldData via all columns
func AddUsersFieldData(iUid int64, iLangcode string, iPreferredLangcode string, iPreferredAdminLangcode string, iName string, iPass string, iMail string, iTimezone string, iStatus int, iCreated int, iChanged int, iAccess int, iLogin int, iInit string, iDefaultLangcode int) error {
	_UsersFieldData := &UsersFieldData{Uid: iUid, Langcode: iLangcode, PreferredLangcode: iPreferredLangcode, PreferredAdminLangcode: iPreferredAdminLangcode, Name: iName, Pass: iPass, Mail: iMail, Timezone: iTimezone, Status: iStatus, Created: iCreated, Changed: iChanged, Access: iAccess, Login: iLogin, Init: iInit, DefaultLangcode: iDefaultLangcode}
	if _, err := Engine.Insert(_UsersFieldData); err != nil {
		return err
	} else {
		return nil
	}
}

// PostUsersFieldData Post UsersFieldData via iUsersFieldData
func PostUsersFieldData(iUsersFieldData *UsersFieldData) (int64, error) {
	_, err := Engine.Insert(iUsersFieldData)
	return iUsersFieldData.Uid, err
}

// PutUsersFieldData Put UsersFieldData
func PutUsersFieldData(iUsersFieldData *UsersFieldData) (int64, error) {
	_, err := Engine.Id(iUsersFieldData.Uid).Update(iUsersFieldData)
	return iUsersFieldData.Uid, err
}

// PutUsersFieldDataViaUid Put UsersFieldData via Uid
func PutUsersFieldDataViaUid(Uid_ int64, iUsersFieldData *UsersFieldData) (int64, error) {
	row, err := Engine.Update(iUsersFieldData, &UsersFieldData{Uid: Uid_})
	return row, err
}

// PutUsersFieldDataViaLangcode Put UsersFieldData via Langcode
func PutUsersFieldDataViaLangcode(Langcode_ string, iUsersFieldData *UsersFieldData) (int64, error) {
	row, err := Engine.Update(iUsersFieldData, &UsersFieldData{Langcode: Langcode_})
	return row, err
}

// PutUsersFieldDataViaPreferredLangcode Put UsersFieldData via PreferredLangcode
func PutUsersFieldDataViaPreferredLangcode(PreferredLangcode_ string, iUsersFieldData *UsersFieldData) (int64, error) {
	row, err := Engine.Update(iUsersFieldData, &UsersFieldData{PreferredLangcode: PreferredLangcode_})
	return row, err
}

// PutUsersFieldDataViaPreferredAdminLangcode Put UsersFieldData via PreferredAdminLangcode
func PutUsersFieldDataViaPreferredAdminLangcode(PreferredAdminLangcode_ string, iUsersFieldData *UsersFieldData) (int64, error) {
	row, err := Engine.Update(iUsersFieldData, &UsersFieldData{PreferredAdminLangcode: PreferredAdminLangcode_})
	return row, err
}

// PutUsersFieldDataViaName Put UsersFieldData via Name
func PutUsersFieldDataViaName(Name_ string, iUsersFieldData *UsersFieldData) (int64, error) {
	row, err := Engine.Update(iUsersFieldData, &UsersFieldData{Name: Name_})
	return row, err
}

// PutUsersFieldDataViaPass Put UsersFieldData via Pass
func PutUsersFieldDataViaPass(Pass_ string, iUsersFieldData *UsersFieldData) (int64, error) {
	row, err := Engine.Update(iUsersFieldData, &UsersFieldData{Pass: Pass_})
	return row, err
}

// PutUsersFieldDataViaMail Put UsersFieldData via Mail
func PutUsersFieldDataViaMail(Mail_ string, iUsersFieldData *UsersFieldData) (int64, error) {
	row, err := Engine.Update(iUsersFieldData, &UsersFieldData{Mail: Mail_})
	return row, err
}

// PutUsersFieldDataViaTimezone Put UsersFieldData via Timezone
func PutUsersFieldDataViaTimezone(Timezone_ string, iUsersFieldData *UsersFieldData) (int64, error) {
	row, err := Engine.Update(iUsersFieldData, &UsersFieldData{Timezone: Timezone_})
	return row, err
}

// PutUsersFieldDataViaStatus Put UsersFieldData via Status
func PutUsersFieldDataViaStatus(Status_ int, iUsersFieldData *UsersFieldData) (int64, error) {
	row, err := Engine.Update(iUsersFieldData, &UsersFieldData{Status: Status_})
	return row, err
}

// PutUsersFieldDataViaCreated Put UsersFieldData via Created
func PutUsersFieldDataViaCreated(Created_ int, iUsersFieldData *UsersFieldData) (int64, error) {
	row, err := Engine.Update(iUsersFieldData, &UsersFieldData{Created: Created_})
	return row, err
}

// PutUsersFieldDataViaChanged Put UsersFieldData via Changed
func PutUsersFieldDataViaChanged(Changed_ int, iUsersFieldData *UsersFieldData) (int64, error) {
	row, err := Engine.Update(iUsersFieldData, &UsersFieldData{Changed: Changed_})
	return row, err
}

// PutUsersFieldDataViaAccess Put UsersFieldData via Access
func PutUsersFieldDataViaAccess(Access_ int, iUsersFieldData *UsersFieldData) (int64, error) {
	row, err := Engine.Update(iUsersFieldData, &UsersFieldData{Access: Access_})
	return row, err
}

// PutUsersFieldDataViaLogin Put UsersFieldData via Login
func PutUsersFieldDataViaLogin(Login_ int, iUsersFieldData *UsersFieldData) (int64, error) {
	row, err := Engine.Update(iUsersFieldData, &UsersFieldData{Login: Login_})
	return row, err
}

// PutUsersFieldDataViaInit Put UsersFieldData via Init
func PutUsersFieldDataViaInit(Init_ string, iUsersFieldData *UsersFieldData) (int64, error) {
	row, err := Engine.Update(iUsersFieldData, &UsersFieldData{Init: Init_})
	return row, err
}

// PutUsersFieldDataViaDefaultLangcode Put UsersFieldData via DefaultLangcode
func PutUsersFieldDataViaDefaultLangcode(DefaultLangcode_ int, iUsersFieldData *UsersFieldData) (int64, error) {
	row, err := Engine.Update(iUsersFieldData, &UsersFieldData{DefaultLangcode: DefaultLangcode_})
	return row, err
}

// UpdateUsersFieldDataViaUid via map[string]interface{}{}
func UpdateUsersFieldDataViaUid(iUid int64, iUsersFieldDataMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(UsersFieldData)).Where("uid = ?", iUid).Update(iUsersFieldDataMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateUsersFieldDataViaLangcode via map[string]interface{}{}
func UpdateUsersFieldDataViaLangcode(iLangcode string, iUsersFieldDataMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(UsersFieldData)).Where("langcode = ?", iLangcode).Update(iUsersFieldDataMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateUsersFieldDataViaPreferredLangcode via map[string]interface{}{}
func UpdateUsersFieldDataViaPreferredLangcode(iPreferredLangcode string, iUsersFieldDataMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(UsersFieldData)).Where("preferred_langcode = ?", iPreferredLangcode).Update(iUsersFieldDataMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateUsersFieldDataViaPreferredAdminLangcode via map[string]interface{}{}
func UpdateUsersFieldDataViaPreferredAdminLangcode(iPreferredAdminLangcode string, iUsersFieldDataMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(UsersFieldData)).Where("preferred_admin_langcode = ?", iPreferredAdminLangcode).Update(iUsersFieldDataMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateUsersFieldDataViaName via map[string]interface{}{}
func UpdateUsersFieldDataViaName(iName string, iUsersFieldDataMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(UsersFieldData)).Where("name = ?", iName).Update(iUsersFieldDataMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateUsersFieldDataViaPass via map[string]interface{}{}
func UpdateUsersFieldDataViaPass(iPass string, iUsersFieldDataMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(UsersFieldData)).Where("pass = ?", iPass).Update(iUsersFieldDataMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateUsersFieldDataViaMail via map[string]interface{}{}
func UpdateUsersFieldDataViaMail(iMail string, iUsersFieldDataMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(UsersFieldData)).Where("mail = ?", iMail).Update(iUsersFieldDataMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateUsersFieldDataViaTimezone via map[string]interface{}{}
func UpdateUsersFieldDataViaTimezone(iTimezone string, iUsersFieldDataMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(UsersFieldData)).Where("timezone = ?", iTimezone).Update(iUsersFieldDataMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateUsersFieldDataViaStatus via map[string]interface{}{}
func UpdateUsersFieldDataViaStatus(iStatus int, iUsersFieldDataMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(UsersFieldData)).Where("status = ?", iStatus).Update(iUsersFieldDataMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateUsersFieldDataViaCreated via map[string]interface{}{}
func UpdateUsersFieldDataViaCreated(iCreated int, iUsersFieldDataMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(UsersFieldData)).Where("created = ?", iCreated).Update(iUsersFieldDataMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateUsersFieldDataViaChanged via map[string]interface{}{}
func UpdateUsersFieldDataViaChanged(iChanged int, iUsersFieldDataMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(UsersFieldData)).Where("changed = ?", iChanged).Update(iUsersFieldDataMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateUsersFieldDataViaAccess via map[string]interface{}{}
func UpdateUsersFieldDataViaAccess(iAccess int, iUsersFieldDataMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(UsersFieldData)).Where("access = ?", iAccess).Update(iUsersFieldDataMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateUsersFieldDataViaLogin via map[string]interface{}{}
func UpdateUsersFieldDataViaLogin(iLogin int, iUsersFieldDataMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(UsersFieldData)).Where("login = ?", iLogin).Update(iUsersFieldDataMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateUsersFieldDataViaInit via map[string]interface{}{}
func UpdateUsersFieldDataViaInit(iInit string, iUsersFieldDataMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(UsersFieldData)).Where("init = ?", iInit).Update(iUsersFieldDataMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateUsersFieldDataViaDefaultLangcode via map[string]interface{}{}
func UpdateUsersFieldDataViaDefaultLangcode(iDefaultLangcode int, iUsersFieldDataMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(UsersFieldData)).Where("default_langcode = ?", iDefaultLangcode).Update(iUsersFieldDataMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// DeleteUsersFieldData Delete UsersFieldData
func DeleteUsersFieldData(iUid int64) (int64, error) {
	row, err := Engine.Id(iUid).Delete(new(UsersFieldData))
	return row, err
}

// DeleteUsersFieldDataViaUid Delete UsersFieldData via Uid
func DeleteUsersFieldDataViaUid(iUid int64) (err error) {
	var has bool
	var _UsersFieldData = &UsersFieldData{Uid: iUid}
	if has, err = Engine.Get(_UsersFieldData); (has == true) && (err == nil) {
		if row, err := Engine.Where("uid = ?", iUid).Delete(new(UsersFieldData)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteUsersFieldDataViaLangcode Delete UsersFieldData via Langcode
func DeleteUsersFieldDataViaLangcode(iLangcode string) (err error) {
	var has bool
	var _UsersFieldData = &UsersFieldData{Langcode: iLangcode}
	if has, err = Engine.Get(_UsersFieldData); (has == true) && (err == nil) {
		if row, err := Engine.Where("langcode = ?", iLangcode).Delete(new(UsersFieldData)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteUsersFieldDataViaPreferredLangcode Delete UsersFieldData via PreferredLangcode
func DeleteUsersFieldDataViaPreferredLangcode(iPreferredLangcode string) (err error) {
	var has bool
	var _UsersFieldData = &UsersFieldData{PreferredLangcode: iPreferredLangcode}
	if has, err = Engine.Get(_UsersFieldData); (has == true) && (err == nil) {
		if row, err := Engine.Where("preferred_langcode = ?", iPreferredLangcode).Delete(new(UsersFieldData)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteUsersFieldDataViaPreferredAdminLangcode Delete UsersFieldData via PreferredAdminLangcode
func DeleteUsersFieldDataViaPreferredAdminLangcode(iPreferredAdminLangcode string) (err error) {
	var has bool
	var _UsersFieldData = &UsersFieldData{PreferredAdminLangcode: iPreferredAdminLangcode}
	if has, err = Engine.Get(_UsersFieldData); (has == true) && (err == nil) {
		if row, err := Engine.Where("preferred_admin_langcode = ?", iPreferredAdminLangcode).Delete(new(UsersFieldData)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteUsersFieldDataViaName Delete UsersFieldData via Name
func DeleteUsersFieldDataViaName(iName string) (err error) {
	var has bool
	var _UsersFieldData = &UsersFieldData{Name: iName}
	if has, err = Engine.Get(_UsersFieldData); (has == true) && (err == nil) {
		if row, err := Engine.Where("name = ?", iName).Delete(new(UsersFieldData)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteUsersFieldDataViaPass Delete UsersFieldData via Pass
func DeleteUsersFieldDataViaPass(iPass string) (err error) {
	var has bool
	var _UsersFieldData = &UsersFieldData{Pass: iPass}
	if has, err = Engine.Get(_UsersFieldData); (has == true) && (err == nil) {
		if row, err := Engine.Where("pass = ?", iPass).Delete(new(UsersFieldData)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteUsersFieldDataViaMail Delete UsersFieldData via Mail
func DeleteUsersFieldDataViaMail(iMail string) (err error) {
	var has bool
	var _UsersFieldData = &UsersFieldData{Mail: iMail}
	if has, err = Engine.Get(_UsersFieldData); (has == true) && (err == nil) {
		if row, err := Engine.Where("mail = ?", iMail).Delete(new(UsersFieldData)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteUsersFieldDataViaTimezone Delete UsersFieldData via Timezone
func DeleteUsersFieldDataViaTimezone(iTimezone string) (err error) {
	var has bool
	var _UsersFieldData = &UsersFieldData{Timezone: iTimezone}
	if has, err = Engine.Get(_UsersFieldData); (has == true) && (err == nil) {
		if row, err := Engine.Where("timezone = ?", iTimezone).Delete(new(UsersFieldData)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteUsersFieldDataViaStatus Delete UsersFieldData via Status
func DeleteUsersFieldDataViaStatus(iStatus int) (err error) {
	var has bool
	var _UsersFieldData = &UsersFieldData{Status: iStatus}
	if has, err = Engine.Get(_UsersFieldData); (has == true) && (err == nil) {
		if row, err := Engine.Where("status = ?", iStatus).Delete(new(UsersFieldData)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteUsersFieldDataViaCreated Delete UsersFieldData via Created
func DeleteUsersFieldDataViaCreated(iCreated int) (err error) {
	var has bool
	var _UsersFieldData = &UsersFieldData{Created: iCreated}
	if has, err = Engine.Get(_UsersFieldData); (has == true) && (err == nil) {
		if row, err := Engine.Where("created = ?", iCreated).Delete(new(UsersFieldData)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteUsersFieldDataViaChanged Delete UsersFieldData via Changed
func DeleteUsersFieldDataViaChanged(iChanged int) (err error) {
	var has bool
	var _UsersFieldData = &UsersFieldData{Changed: iChanged}
	if has, err = Engine.Get(_UsersFieldData); (has == true) && (err == nil) {
		if row, err := Engine.Where("changed = ?", iChanged).Delete(new(UsersFieldData)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteUsersFieldDataViaAccess Delete UsersFieldData via Access
func DeleteUsersFieldDataViaAccess(iAccess int) (err error) {
	var has bool
	var _UsersFieldData = &UsersFieldData{Access: iAccess}
	if has, err = Engine.Get(_UsersFieldData); (has == true) && (err == nil) {
		if row, err := Engine.Where("access = ?", iAccess).Delete(new(UsersFieldData)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteUsersFieldDataViaLogin Delete UsersFieldData via Login
func DeleteUsersFieldDataViaLogin(iLogin int) (err error) {
	var has bool
	var _UsersFieldData = &UsersFieldData{Login: iLogin}
	if has, err = Engine.Get(_UsersFieldData); (has == true) && (err == nil) {
		if row, err := Engine.Where("login = ?", iLogin).Delete(new(UsersFieldData)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteUsersFieldDataViaInit Delete UsersFieldData via Init
func DeleteUsersFieldDataViaInit(iInit string) (err error) {
	var has bool
	var _UsersFieldData = &UsersFieldData{Init: iInit}
	if has, err = Engine.Get(_UsersFieldData); (has == true) && (err == nil) {
		if row, err := Engine.Where("init = ?", iInit).Delete(new(UsersFieldData)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteUsersFieldDataViaDefaultLangcode Delete UsersFieldData via DefaultLangcode
func DeleteUsersFieldDataViaDefaultLangcode(iDefaultLangcode int) (err error) {
	var has bool
	var _UsersFieldData = &UsersFieldData{DefaultLangcode: iDefaultLangcode}
	if has, err = Engine.Get(_UsersFieldData); (has == true) && (err == nil) {
		if row, err := Engine.Where("default_langcode = ?", iDefaultLangcode).Delete(new(UsersFieldData)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}
