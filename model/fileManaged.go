package model

type FileManaged struct {
	Fid      int64  `xorm:"not null pk autoincr INT(10)"`
	Uuid     string `xorm:"not null unique VARCHAR(128)"`
	Langcode string `xorm:"not null VARCHAR(12)"`
	Uid      int    `xorm:"index INT(10)"`
	Filename string `xorm:"VARCHAR(255)"`
	Uri      string `xorm:"not null index VARCHAR(255)"`
	Filemime string `xorm:"VARCHAR(255)"`
	Filesize int64  `xorm:"BIGINT(20)"`
	Status   int    `xorm:"not null index TINYINT(4)"`
	Created  int    `xorm:"INT(11)"`
	Changed  int    `xorm:"not null index INT(11)"`
}

// GetFileManagedsCount FileManageds' Count
func GetFileManagedsCount(offset int, limit int) (int64, error) {
	total, err := Engine.Limit(limit, offset).Count(&FileManaged{})
	return total, err
}

// GetFileManagedCountViaFid Get FileManaged via Fid
func GetFileManagedCountViaFid(iFid int64) int64 {
	n, _ := Engine.Where("fid = ?", iFid).Count(&FileManaged{Fid: iFid})
	return n
}

// GetFileManagedCountViaUuid Get FileManaged via Uuid
func GetFileManagedCountViaUuid(iUuid string) int64 {
	n, _ := Engine.Where("uuid = ?", iUuid).Count(&FileManaged{Uuid: iUuid})
	return n
}

// GetFileManagedCountViaLangcode Get FileManaged via Langcode
func GetFileManagedCountViaLangcode(iLangcode string) int64 {
	n, _ := Engine.Where("langcode = ?", iLangcode).Count(&FileManaged{Langcode: iLangcode})
	return n
}

// GetFileManagedCountViaUid Get FileManaged via Uid
func GetFileManagedCountViaUid(iUid int) int64 {
	n, _ := Engine.Where("uid = ?", iUid).Count(&FileManaged{Uid: iUid})
	return n
}

// GetFileManagedCountViaFilename Get FileManaged via Filename
func GetFileManagedCountViaFilename(iFilename string) int64 {
	n, _ := Engine.Where("filename = ?", iFilename).Count(&FileManaged{Filename: iFilename})
	return n
}

// GetFileManagedCountViaUri Get FileManaged via Uri
func GetFileManagedCountViaUri(iUri string) int64 {
	n, _ := Engine.Where("uri = ?", iUri).Count(&FileManaged{Uri: iUri})
	return n
}

// GetFileManagedCountViaFilemime Get FileManaged via Filemime
func GetFileManagedCountViaFilemime(iFilemime string) int64 {
	n, _ := Engine.Where("filemime = ?", iFilemime).Count(&FileManaged{Filemime: iFilemime})
	return n
}

// GetFileManagedCountViaFilesize Get FileManaged via Filesize
func GetFileManagedCountViaFilesize(iFilesize int64) int64 {
	n, _ := Engine.Where("filesize = ?", iFilesize).Count(&FileManaged{Filesize: iFilesize})
	return n
}

// GetFileManagedCountViaStatus Get FileManaged via Status
func GetFileManagedCountViaStatus(iStatus int) int64 {
	n, _ := Engine.Where("status = ?", iStatus).Count(&FileManaged{Status: iStatus})
	return n
}

// GetFileManagedCountViaCreated Get FileManaged via Created
func GetFileManagedCountViaCreated(iCreated int) int64 {
	n, _ := Engine.Where("created = ?", iCreated).Count(&FileManaged{Created: iCreated})
	return n
}

// GetFileManagedCountViaChanged Get FileManaged via Changed
func GetFileManagedCountViaChanged(iChanged int) int64 {
	n, _ := Engine.Where("changed = ?", iChanged).Count(&FileManaged{Changed: iChanged})
	return n
}

// GetFileManagedsByFidAndUuidAndLangcode Get FileManageds via FidAndUuidAndLangcode
func GetFileManagedsByFidAndUuidAndLangcode(offset int, limit int, Fid_ int64, Uuid_ string, Langcode_ string) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("fid = ? and uuid = ? and langcode = ?", Fid_, Uuid_, Langcode_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByFidAndUuidAndUid Get FileManageds via FidAndUuidAndUid
func GetFileManagedsByFidAndUuidAndUid(offset int, limit int, Fid_ int64, Uuid_ string, Uid_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("fid = ? and uuid = ? and uid = ?", Fid_, Uuid_, Uid_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByFidAndUuidAndFilename Get FileManageds via FidAndUuidAndFilename
func GetFileManagedsByFidAndUuidAndFilename(offset int, limit int, Fid_ int64, Uuid_ string, Filename_ string) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("fid = ? and uuid = ? and filename = ?", Fid_, Uuid_, Filename_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByFidAndUuidAndUri Get FileManageds via FidAndUuidAndUri
func GetFileManagedsByFidAndUuidAndUri(offset int, limit int, Fid_ int64, Uuid_ string, Uri_ string) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("fid = ? and uuid = ? and uri = ?", Fid_, Uuid_, Uri_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByFidAndUuidAndFilemime Get FileManageds via FidAndUuidAndFilemime
func GetFileManagedsByFidAndUuidAndFilemime(offset int, limit int, Fid_ int64, Uuid_ string, Filemime_ string) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("fid = ? and uuid = ? and filemime = ?", Fid_, Uuid_, Filemime_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByFidAndUuidAndFilesize Get FileManageds via FidAndUuidAndFilesize
func GetFileManagedsByFidAndUuidAndFilesize(offset int, limit int, Fid_ int64, Uuid_ string, Filesize_ int64) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("fid = ? and uuid = ? and filesize = ?", Fid_, Uuid_, Filesize_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByFidAndUuidAndStatus Get FileManageds via FidAndUuidAndStatus
func GetFileManagedsByFidAndUuidAndStatus(offset int, limit int, Fid_ int64, Uuid_ string, Status_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("fid = ? and uuid = ? and status = ?", Fid_, Uuid_, Status_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByFidAndUuidAndCreated Get FileManageds via FidAndUuidAndCreated
func GetFileManagedsByFidAndUuidAndCreated(offset int, limit int, Fid_ int64, Uuid_ string, Created_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("fid = ? and uuid = ? and created = ?", Fid_, Uuid_, Created_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByFidAndUuidAndChanged Get FileManageds via FidAndUuidAndChanged
func GetFileManagedsByFidAndUuidAndChanged(offset int, limit int, Fid_ int64, Uuid_ string, Changed_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("fid = ? and uuid = ? and changed = ?", Fid_, Uuid_, Changed_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByFidAndLangcodeAndUid Get FileManageds via FidAndLangcodeAndUid
func GetFileManagedsByFidAndLangcodeAndUid(offset int, limit int, Fid_ int64, Langcode_ string, Uid_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("fid = ? and langcode = ? and uid = ?", Fid_, Langcode_, Uid_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByFidAndLangcodeAndFilename Get FileManageds via FidAndLangcodeAndFilename
func GetFileManagedsByFidAndLangcodeAndFilename(offset int, limit int, Fid_ int64, Langcode_ string, Filename_ string) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("fid = ? and langcode = ? and filename = ?", Fid_, Langcode_, Filename_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByFidAndLangcodeAndUri Get FileManageds via FidAndLangcodeAndUri
func GetFileManagedsByFidAndLangcodeAndUri(offset int, limit int, Fid_ int64, Langcode_ string, Uri_ string) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("fid = ? and langcode = ? and uri = ?", Fid_, Langcode_, Uri_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByFidAndLangcodeAndFilemime Get FileManageds via FidAndLangcodeAndFilemime
func GetFileManagedsByFidAndLangcodeAndFilemime(offset int, limit int, Fid_ int64, Langcode_ string, Filemime_ string) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("fid = ? and langcode = ? and filemime = ?", Fid_, Langcode_, Filemime_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByFidAndLangcodeAndFilesize Get FileManageds via FidAndLangcodeAndFilesize
func GetFileManagedsByFidAndLangcodeAndFilesize(offset int, limit int, Fid_ int64, Langcode_ string, Filesize_ int64) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("fid = ? and langcode = ? and filesize = ?", Fid_, Langcode_, Filesize_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByFidAndLangcodeAndStatus Get FileManageds via FidAndLangcodeAndStatus
func GetFileManagedsByFidAndLangcodeAndStatus(offset int, limit int, Fid_ int64, Langcode_ string, Status_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("fid = ? and langcode = ? and status = ?", Fid_, Langcode_, Status_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByFidAndLangcodeAndCreated Get FileManageds via FidAndLangcodeAndCreated
func GetFileManagedsByFidAndLangcodeAndCreated(offset int, limit int, Fid_ int64, Langcode_ string, Created_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("fid = ? and langcode = ? and created = ?", Fid_, Langcode_, Created_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByFidAndLangcodeAndChanged Get FileManageds via FidAndLangcodeAndChanged
func GetFileManagedsByFidAndLangcodeAndChanged(offset int, limit int, Fid_ int64, Langcode_ string, Changed_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("fid = ? and langcode = ? and changed = ?", Fid_, Langcode_, Changed_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByFidAndUidAndFilename Get FileManageds via FidAndUidAndFilename
func GetFileManagedsByFidAndUidAndFilename(offset int, limit int, Fid_ int64, Uid_ int, Filename_ string) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("fid = ? and uid = ? and filename = ?", Fid_, Uid_, Filename_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByFidAndUidAndUri Get FileManageds via FidAndUidAndUri
func GetFileManagedsByFidAndUidAndUri(offset int, limit int, Fid_ int64, Uid_ int, Uri_ string) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("fid = ? and uid = ? and uri = ?", Fid_, Uid_, Uri_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByFidAndUidAndFilemime Get FileManageds via FidAndUidAndFilemime
func GetFileManagedsByFidAndUidAndFilemime(offset int, limit int, Fid_ int64, Uid_ int, Filemime_ string) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("fid = ? and uid = ? and filemime = ?", Fid_, Uid_, Filemime_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByFidAndUidAndFilesize Get FileManageds via FidAndUidAndFilesize
func GetFileManagedsByFidAndUidAndFilesize(offset int, limit int, Fid_ int64, Uid_ int, Filesize_ int64) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("fid = ? and uid = ? and filesize = ?", Fid_, Uid_, Filesize_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByFidAndUidAndStatus Get FileManageds via FidAndUidAndStatus
func GetFileManagedsByFidAndUidAndStatus(offset int, limit int, Fid_ int64, Uid_ int, Status_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("fid = ? and uid = ? and status = ?", Fid_, Uid_, Status_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByFidAndUidAndCreated Get FileManageds via FidAndUidAndCreated
func GetFileManagedsByFidAndUidAndCreated(offset int, limit int, Fid_ int64, Uid_ int, Created_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("fid = ? and uid = ? and created = ?", Fid_, Uid_, Created_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByFidAndUidAndChanged Get FileManageds via FidAndUidAndChanged
func GetFileManagedsByFidAndUidAndChanged(offset int, limit int, Fid_ int64, Uid_ int, Changed_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("fid = ? and uid = ? and changed = ?", Fid_, Uid_, Changed_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByFidAndFilenameAndUri Get FileManageds via FidAndFilenameAndUri
func GetFileManagedsByFidAndFilenameAndUri(offset int, limit int, Fid_ int64, Filename_ string, Uri_ string) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("fid = ? and filename = ? and uri = ?", Fid_, Filename_, Uri_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByFidAndFilenameAndFilemime Get FileManageds via FidAndFilenameAndFilemime
func GetFileManagedsByFidAndFilenameAndFilemime(offset int, limit int, Fid_ int64, Filename_ string, Filemime_ string) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("fid = ? and filename = ? and filemime = ?", Fid_, Filename_, Filemime_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByFidAndFilenameAndFilesize Get FileManageds via FidAndFilenameAndFilesize
func GetFileManagedsByFidAndFilenameAndFilesize(offset int, limit int, Fid_ int64, Filename_ string, Filesize_ int64) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("fid = ? and filename = ? and filesize = ?", Fid_, Filename_, Filesize_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByFidAndFilenameAndStatus Get FileManageds via FidAndFilenameAndStatus
func GetFileManagedsByFidAndFilenameAndStatus(offset int, limit int, Fid_ int64, Filename_ string, Status_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("fid = ? and filename = ? and status = ?", Fid_, Filename_, Status_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByFidAndFilenameAndCreated Get FileManageds via FidAndFilenameAndCreated
func GetFileManagedsByFidAndFilenameAndCreated(offset int, limit int, Fid_ int64, Filename_ string, Created_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("fid = ? and filename = ? and created = ?", Fid_, Filename_, Created_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByFidAndFilenameAndChanged Get FileManageds via FidAndFilenameAndChanged
func GetFileManagedsByFidAndFilenameAndChanged(offset int, limit int, Fid_ int64, Filename_ string, Changed_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("fid = ? and filename = ? and changed = ?", Fid_, Filename_, Changed_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByFidAndUriAndFilemime Get FileManageds via FidAndUriAndFilemime
func GetFileManagedsByFidAndUriAndFilemime(offset int, limit int, Fid_ int64, Uri_ string, Filemime_ string) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("fid = ? and uri = ? and filemime = ?", Fid_, Uri_, Filemime_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByFidAndUriAndFilesize Get FileManageds via FidAndUriAndFilesize
func GetFileManagedsByFidAndUriAndFilesize(offset int, limit int, Fid_ int64, Uri_ string, Filesize_ int64) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("fid = ? and uri = ? and filesize = ?", Fid_, Uri_, Filesize_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByFidAndUriAndStatus Get FileManageds via FidAndUriAndStatus
func GetFileManagedsByFidAndUriAndStatus(offset int, limit int, Fid_ int64, Uri_ string, Status_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("fid = ? and uri = ? and status = ?", Fid_, Uri_, Status_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByFidAndUriAndCreated Get FileManageds via FidAndUriAndCreated
func GetFileManagedsByFidAndUriAndCreated(offset int, limit int, Fid_ int64, Uri_ string, Created_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("fid = ? and uri = ? and created = ?", Fid_, Uri_, Created_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByFidAndUriAndChanged Get FileManageds via FidAndUriAndChanged
func GetFileManagedsByFidAndUriAndChanged(offset int, limit int, Fid_ int64, Uri_ string, Changed_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("fid = ? and uri = ? and changed = ?", Fid_, Uri_, Changed_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByFidAndFilemimeAndFilesize Get FileManageds via FidAndFilemimeAndFilesize
func GetFileManagedsByFidAndFilemimeAndFilesize(offset int, limit int, Fid_ int64, Filemime_ string, Filesize_ int64) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("fid = ? and filemime = ? and filesize = ?", Fid_, Filemime_, Filesize_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByFidAndFilemimeAndStatus Get FileManageds via FidAndFilemimeAndStatus
func GetFileManagedsByFidAndFilemimeAndStatus(offset int, limit int, Fid_ int64, Filemime_ string, Status_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("fid = ? and filemime = ? and status = ?", Fid_, Filemime_, Status_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByFidAndFilemimeAndCreated Get FileManageds via FidAndFilemimeAndCreated
func GetFileManagedsByFidAndFilemimeAndCreated(offset int, limit int, Fid_ int64, Filemime_ string, Created_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("fid = ? and filemime = ? and created = ?", Fid_, Filemime_, Created_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByFidAndFilemimeAndChanged Get FileManageds via FidAndFilemimeAndChanged
func GetFileManagedsByFidAndFilemimeAndChanged(offset int, limit int, Fid_ int64, Filemime_ string, Changed_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("fid = ? and filemime = ? and changed = ?", Fid_, Filemime_, Changed_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByFidAndFilesizeAndStatus Get FileManageds via FidAndFilesizeAndStatus
func GetFileManagedsByFidAndFilesizeAndStatus(offset int, limit int, Fid_ int64, Filesize_ int64, Status_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("fid = ? and filesize = ? and status = ?", Fid_, Filesize_, Status_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByFidAndFilesizeAndCreated Get FileManageds via FidAndFilesizeAndCreated
func GetFileManagedsByFidAndFilesizeAndCreated(offset int, limit int, Fid_ int64, Filesize_ int64, Created_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("fid = ? and filesize = ? and created = ?", Fid_, Filesize_, Created_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByFidAndFilesizeAndChanged Get FileManageds via FidAndFilesizeAndChanged
func GetFileManagedsByFidAndFilesizeAndChanged(offset int, limit int, Fid_ int64, Filesize_ int64, Changed_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("fid = ? and filesize = ? and changed = ?", Fid_, Filesize_, Changed_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByFidAndStatusAndCreated Get FileManageds via FidAndStatusAndCreated
func GetFileManagedsByFidAndStatusAndCreated(offset int, limit int, Fid_ int64, Status_ int, Created_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("fid = ? and status = ? and created = ?", Fid_, Status_, Created_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByFidAndStatusAndChanged Get FileManageds via FidAndStatusAndChanged
func GetFileManagedsByFidAndStatusAndChanged(offset int, limit int, Fid_ int64, Status_ int, Changed_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("fid = ? and status = ? and changed = ?", Fid_, Status_, Changed_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByFidAndCreatedAndChanged Get FileManageds via FidAndCreatedAndChanged
func GetFileManagedsByFidAndCreatedAndChanged(offset int, limit int, Fid_ int64, Created_ int, Changed_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("fid = ? and created = ? and changed = ?", Fid_, Created_, Changed_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByUuidAndLangcodeAndUid Get FileManageds via UuidAndLangcodeAndUid
func GetFileManagedsByUuidAndLangcodeAndUid(offset int, limit int, Uuid_ string, Langcode_ string, Uid_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("uuid = ? and langcode = ? and uid = ?", Uuid_, Langcode_, Uid_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByUuidAndLangcodeAndFilename Get FileManageds via UuidAndLangcodeAndFilename
func GetFileManagedsByUuidAndLangcodeAndFilename(offset int, limit int, Uuid_ string, Langcode_ string, Filename_ string) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("uuid = ? and langcode = ? and filename = ?", Uuid_, Langcode_, Filename_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByUuidAndLangcodeAndUri Get FileManageds via UuidAndLangcodeAndUri
func GetFileManagedsByUuidAndLangcodeAndUri(offset int, limit int, Uuid_ string, Langcode_ string, Uri_ string) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("uuid = ? and langcode = ? and uri = ?", Uuid_, Langcode_, Uri_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByUuidAndLangcodeAndFilemime Get FileManageds via UuidAndLangcodeAndFilemime
func GetFileManagedsByUuidAndLangcodeAndFilemime(offset int, limit int, Uuid_ string, Langcode_ string, Filemime_ string) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("uuid = ? and langcode = ? and filemime = ?", Uuid_, Langcode_, Filemime_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByUuidAndLangcodeAndFilesize Get FileManageds via UuidAndLangcodeAndFilesize
func GetFileManagedsByUuidAndLangcodeAndFilesize(offset int, limit int, Uuid_ string, Langcode_ string, Filesize_ int64) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("uuid = ? and langcode = ? and filesize = ?", Uuid_, Langcode_, Filesize_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByUuidAndLangcodeAndStatus Get FileManageds via UuidAndLangcodeAndStatus
func GetFileManagedsByUuidAndLangcodeAndStatus(offset int, limit int, Uuid_ string, Langcode_ string, Status_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("uuid = ? and langcode = ? and status = ?", Uuid_, Langcode_, Status_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByUuidAndLangcodeAndCreated Get FileManageds via UuidAndLangcodeAndCreated
func GetFileManagedsByUuidAndLangcodeAndCreated(offset int, limit int, Uuid_ string, Langcode_ string, Created_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("uuid = ? and langcode = ? and created = ?", Uuid_, Langcode_, Created_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByUuidAndLangcodeAndChanged Get FileManageds via UuidAndLangcodeAndChanged
func GetFileManagedsByUuidAndLangcodeAndChanged(offset int, limit int, Uuid_ string, Langcode_ string, Changed_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("uuid = ? and langcode = ? and changed = ?", Uuid_, Langcode_, Changed_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByUuidAndUidAndFilename Get FileManageds via UuidAndUidAndFilename
func GetFileManagedsByUuidAndUidAndFilename(offset int, limit int, Uuid_ string, Uid_ int, Filename_ string) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("uuid = ? and uid = ? and filename = ?", Uuid_, Uid_, Filename_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByUuidAndUidAndUri Get FileManageds via UuidAndUidAndUri
func GetFileManagedsByUuidAndUidAndUri(offset int, limit int, Uuid_ string, Uid_ int, Uri_ string) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("uuid = ? and uid = ? and uri = ?", Uuid_, Uid_, Uri_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByUuidAndUidAndFilemime Get FileManageds via UuidAndUidAndFilemime
func GetFileManagedsByUuidAndUidAndFilemime(offset int, limit int, Uuid_ string, Uid_ int, Filemime_ string) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("uuid = ? and uid = ? and filemime = ?", Uuid_, Uid_, Filemime_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByUuidAndUidAndFilesize Get FileManageds via UuidAndUidAndFilesize
func GetFileManagedsByUuidAndUidAndFilesize(offset int, limit int, Uuid_ string, Uid_ int, Filesize_ int64) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("uuid = ? and uid = ? and filesize = ?", Uuid_, Uid_, Filesize_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByUuidAndUidAndStatus Get FileManageds via UuidAndUidAndStatus
func GetFileManagedsByUuidAndUidAndStatus(offset int, limit int, Uuid_ string, Uid_ int, Status_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("uuid = ? and uid = ? and status = ?", Uuid_, Uid_, Status_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByUuidAndUidAndCreated Get FileManageds via UuidAndUidAndCreated
func GetFileManagedsByUuidAndUidAndCreated(offset int, limit int, Uuid_ string, Uid_ int, Created_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("uuid = ? and uid = ? and created = ?", Uuid_, Uid_, Created_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByUuidAndUidAndChanged Get FileManageds via UuidAndUidAndChanged
func GetFileManagedsByUuidAndUidAndChanged(offset int, limit int, Uuid_ string, Uid_ int, Changed_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("uuid = ? and uid = ? and changed = ?", Uuid_, Uid_, Changed_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByUuidAndFilenameAndUri Get FileManageds via UuidAndFilenameAndUri
func GetFileManagedsByUuidAndFilenameAndUri(offset int, limit int, Uuid_ string, Filename_ string, Uri_ string) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("uuid = ? and filename = ? and uri = ?", Uuid_, Filename_, Uri_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByUuidAndFilenameAndFilemime Get FileManageds via UuidAndFilenameAndFilemime
func GetFileManagedsByUuidAndFilenameAndFilemime(offset int, limit int, Uuid_ string, Filename_ string, Filemime_ string) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("uuid = ? and filename = ? and filemime = ?", Uuid_, Filename_, Filemime_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByUuidAndFilenameAndFilesize Get FileManageds via UuidAndFilenameAndFilesize
func GetFileManagedsByUuidAndFilenameAndFilesize(offset int, limit int, Uuid_ string, Filename_ string, Filesize_ int64) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("uuid = ? and filename = ? and filesize = ?", Uuid_, Filename_, Filesize_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByUuidAndFilenameAndStatus Get FileManageds via UuidAndFilenameAndStatus
func GetFileManagedsByUuidAndFilenameAndStatus(offset int, limit int, Uuid_ string, Filename_ string, Status_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("uuid = ? and filename = ? and status = ?", Uuid_, Filename_, Status_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByUuidAndFilenameAndCreated Get FileManageds via UuidAndFilenameAndCreated
func GetFileManagedsByUuidAndFilenameAndCreated(offset int, limit int, Uuid_ string, Filename_ string, Created_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("uuid = ? and filename = ? and created = ?", Uuid_, Filename_, Created_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByUuidAndFilenameAndChanged Get FileManageds via UuidAndFilenameAndChanged
func GetFileManagedsByUuidAndFilenameAndChanged(offset int, limit int, Uuid_ string, Filename_ string, Changed_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("uuid = ? and filename = ? and changed = ?", Uuid_, Filename_, Changed_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByUuidAndUriAndFilemime Get FileManageds via UuidAndUriAndFilemime
func GetFileManagedsByUuidAndUriAndFilemime(offset int, limit int, Uuid_ string, Uri_ string, Filemime_ string) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("uuid = ? and uri = ? and filemime = ?", Uuid_, Uri_, Filemime_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByUuidAndUriAndFilesize Get FileManageds via UuidAndUriAndFilesize
func GetFileManagedsByUuidAndUriAndFilesize(offset int, limit int, Uuid_ string, Uri_ string, Filesize_ int64) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("uuid = ? and uri = ? and filesize = ?", Uuid_, Uri_, Filesize_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByUuidAndUriAndStatus Get FileManageds via UuidAndUriAndStatus
func GetFileManagedsByUuidAndUriAndStatus(offset int, limit int, Uuid_ string, Uri_ string, Status_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("uuid = ? and uri = ? and status = ?", Uuid_, Uri_, Status_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByUuidAndUriAndCreated Get FileManageds via UuidAndUriAndCreated
func GetFileManagedsByUuidAndUriAndCreated(offset int, limit int, Uuid_ string, Uri_ string, Created_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("uuid = ? and uri = ? and created = ?", Uuid_, Uri_, Created_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByUuidAndUriAndChanged Get FileManageds via UuidAndUriAndChanged
func GetFileManagedsByUuidAndUriAndChanged(offset int, limit int, Uuid_ string, Uri_ string, Changed_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("uuid = ? and uri = ? and changed = ?", Uuid_, Uri_, Changed_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByUuidAndFilemimeAndFilesize Get FileManageds via UuidAndFilemimeAndFilesize
func GetFileManagedsByUuidAndFilemimeAndFilesize(offset int, limit int, Uuid_ string, Filemime_ string, Filesize_ int64) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("uuid = ? and filemime = ? and filesize = ?", Uuid_, Filemime_, Filesize_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByUuidAndFilemimeAndStatus Get FileManageds via UuidAndFilemimeAndStatus
func GetFileManagedsByUuidAndFilemimeAndStatus(offset int, limit int, Uuid_ string, Filemime_ string, Status_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("uuid = ? and filemime = ? and status = ?", Uuid_, Filemime_, Status_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByUuidAndFilemimeAndCreated Get FileManageds via UuidAndFilemimeAndCreated
func GetFileManagedsByUuidAndFilemimeAndCreated(offset int, limit int, Uuid_ string, Filemime_ string, Created_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("uuid = ? and filemime = ? and created = ?", Uuid_, Filemime_, Created_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByUuidAndFilemimeAndChanged Get FileManageds via UuidAndFilemimeAndChanged
func GetFileManagedsByUuidAndFilemimeAndChanged(offset int, limit int, Uuid_ string, Filemime_ string, Changed_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("uuid = ? and filemime = ? and changed = ?", Uuid_, Filemime_, Changed_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByUuidAndFilesizeAndStatus Get FileManageds via UuidAndFilesizeAndStatus
func GetFileManagedsByUuidAndFilesizeAndStatus(offset int, limit int, Uuid_ string, Filesize_ int64, Status_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("uuid = ? and filesize = ? and status = ?", Uuid_, Filesize_, Status_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByUuidAndFilesizeAndCreated Get FileManageds via UuidAndFilesizeAndCreated
func GetFileManagedsByUuidAndFilesizeAndCreated(offset int, limit int, Uuid_ string, Filesize_ int64, Created_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("uuid = ? and filesize = ? and created = ?", Uuid_, Filesize_, Created_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByUuidAndFilesizeAndChanged Get FileManageds via UuidAndFilesizeAndChanged
func GetFileManagedsByUuidAndFilesizeAndChanged(offset int, limit int, Uuid_ string, Filesize_ int64, Changed_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("uuid = ? and filesize = ? and changed = ?", Uuid_, Filesize_, Changed_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByUuidAndStatusAndCreated Get FileManageds via UuidAndStatusAndCreated
func GetFileManagedsByUuidAndStatusAndCreated(offset int, limit int, Uuid_ string, Status_ int, Created_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("uuid = ? and status = ? and created = ?", Uuid_, Status_, Created_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByUuidAndStatusAndChanged Get FileManageds via UuidAndStatusAndChanged
func GetFileManagedsByUuidAndStatusAndChanged(offset int, limit int, Uuid_ string, Status_ int, Changed_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("uuid = ? and status = ? and changed = ?", Uuid_, Status_, Changed_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByUuidAndCreatedAndChanged Get FileManageds via UuidAndCreatedAndChanged
func GetFileManagedsByUuidAndCreatedAndChanged(offset int, limit int, Uuid_ string, Created_ int, Changed_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("uuid = ? and created = ? and changed = ?", Uuid_, Created_, Changed_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByLangcodeAndUidAndFilename Get FileManageds via LangcodeAndUidAndFilename
func GetFileManagedsByLangcodeAndUidAndFilename(offset int, limit int, Langcode_ string, Uid_ int, Filename_ string) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("langcode = ? and uid = ? and filename = ?", Langcode_, Uid_, Filename_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByLangcodeAndUidAndUri Get FileManageds via LangcodeAndUidAndUri
func GetFileManagedsByLangcodeAndUidAndUri(offset int, limit int, Langcode_ string, Uid_ int, Uri_ string) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("langcode = ? and uid = ? and uri = ?", Langcode_, Uid_, Uri_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByLangcodeAndUidAndFilemime Get FileManageds via LangcodeAndUidAndFilemime
func GetFileManagedsByLangcodeAndUidAndFilemime(offset int, limit int, Langcode_ string, Uid_ int, Filemime_ string) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("langcode = ? and uid = ? and filemime = ?", Langcode_, Uid_, Filemime_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByLangcodeAndUidAndFilesize Get FileManageds via LangcodeAndUidAndFilesize
func GetFileManagedsByLangcodeAndUidAndFilesize(offset int, limit int, Langcode_ string, Uid_ int, Filesize_ int64) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("langcode = ? and uid = ? and filesize = ?", Langcode_, Uid_, Filesize_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByLangcodeAndUidAndStatus Get FileManageds via LangcodeAndUidAndStatus
func GetFileManagedsByLangcodeAndUidAndStatus(offset int, limit int, Langcode_ string, Uid_ int, Status_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("langcode = ? and uid = ? and status = ?", Langcode_, Uid_, Status_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByLangcodeAndUidAndCreated Get FileManageds via LangcodeAndUidAndCreated
func GetFileManagedsByLangcodeAndUidAndCreated(offset int, limit int, Langcode_ string, Uid_ int, Created_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("langcode = ? and uid = ? and created = ?", Langcode_, Uid_, Created_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByLangcodeAndUidAndChanged Get FileManageds via LangcodeAndUidAndChanged
func GetFileManagedsByLangcodeAndUidAndChanged(offset int, limit int, Langcode_ string, Uid_ int, Changed_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("langcode = ? and uid = ? and changed = ?", Langcode_, Uid_, Changed_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByLangcodeAndFilenameAndUri Get FileManageds via LangcodeAndFilenameAndUri
func GetFileManagedsByLangcodeAndFilenameAndUri(offset int, limit int, Langcode_ string, Filename_ string, Uri_ string) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("langcode = ? and filename = ? and uri = ?", Langcode_, Filename_, Uri_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByLangcodeAndFilenameAndFilemime Get FileManageds via LangcodeAndFilenameAndFilemime
func GetFileManagedsByLangcodeAndFilenameAndFilemime(offset int, limit int, Langcode_ string, Filename_ string, Filemime_ string) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("langcode = ? and filename = ? and filemime = ?", Langcode_, Filename_, Filemime_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByLangcodeAndFilenameAndFilesize Get FileManageds via LangcodeAndFilenameAndFilesize
func GetFileManagedsByLangcodeAndFilenameAndFilesize(offset int, limit int, Langcode_ string, Filename_ string, Filesize_ int64) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("langcode = ? and filename = ? and filesize = ?", Langcode_, Filename_, Filesize_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByLangcodeAndFilenameAndStatus Get FileManageds via LangcodeAndFilenameAndStatus
func GetFileManagedsByLangcodeAndFilenameAndStatus(offset int, limit int, Langcode_ string, Filename_ string, Status_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("langcode = ? and filename = ? and status = ?", Langcode_, Filename_, Status_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByLangcodeAndFilenameAndCreated Get FileManageds via LangcodeAndFilenameAndCreated
func GetFileManagedsByLangcodeAndFilenameAndCreated(offset int, limit int, Langcode_ string, Filename_ string, Created_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("langcode = ? and filename = ? and created = ?", Langcode_, Filename_, Created_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByLangcodeAndFilenameAndChanged Get FileManageds via LangcodeAndFilenameAndChanged
func GetFileManagedsByLangcodeAndFilenameAndChanged(offset int, limit int, Langcode_ string, Filename_ string, Changed_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("langcode = ? and filename = ? and changed = ?", Langcode_, Filename_, Changed_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByLangcodeAndUriAndFilemime Get FileManageds via LangcodeAndUriAndFilemime
func GetFileManagedsByLangcodeAndUriAndFilemime(offset int, limit int, Langcode_ string, Uri_ string, Filemime_ string) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("langcode = ? and uri = ? and filemime = ?", Langcode_, Uri_, Filemime_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByLangcodeAndUriAndFilesize Get FileManageds via LangcodeAndUriAndFilesize
func GetFileManagedsByLangcodeAndUriAndFilesize(offset int, limit int, Langcode_ string, Uri_ string, Filesize_ int64) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("langcode = ? and uri = ? and filesize = ?", Langcode_, Uri_, Filesize_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByLangcodeAndUriAndStatus Get FileManageds via LangcodeAndUriAndStatus
func GetFileManagedsByLangcodeAndUriAndStatus(offset int, limit int, Langcode_ string, Uri_ string, Status_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("langcode = ? and uri = ? and status = ?", Langcode_, Uri_, Status_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByLangcodeAndUriAndCreated Get FileManageds via LangcodeAndUriAndCreated
func GetFileManagedsByLangcodeAndUriAndCreated(offset int, limit int, Langcode_ string, Uri_ string, Created_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("langcode = ? and uri = ? and created = ?", Langcode_, Uri_, Created_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByLangcodeAndUriAndChanged Get FileManageds via LangcodeAndUriAndChanged
func GetFileManagedsByLangcodeAndUriAndChanged(offset int, limit int, Langcode_ string, Uri_ string, Changed_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("langcode = ? and uri = ? and changed = ?", Langcode_, Uri_, Changed_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByLangcodeAndFilemimeAndFilesize Get FileManageds via LangcodeAndFilemimeAndFilesize
func GetFileManagedsByLangcodeAndFilemimeAndFilesize(offset int, limit int, Langcode_ string, Filemime_ string, Filesize_ int64) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("langcode = ? and filemime = ? and filesize = ?", Langcode_, Filemime_, Filesize_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByLangcodeAndFilemimeAndStatus Get FileManageds via LangcodeAndFilemimeAndStatus
func GetFileManagedsByLangcodeAndFilemimeAndStatus(offset int, limit int, Langcode_ string, Filemime_ string, Status_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("langcode = ? and filemime = ? and status = ?", Langcode_, Filemime_, Status_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByLangcodeAndFilemimeAndCreated Get FileManageds via LangcodeAndFilemimeAndCreated
func GetFileManagedsByLangcodeAndFilemimeAndCreated(offset int, limit int, Langcode_ string, Filemime_ string, Created_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("langcode = ? and filemime = ? and created = ?", Langcode_, Filemime_, Created_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByLangcodeAndFilemimeAndChanged Get FileManageds via LangcodeAndFilemimeAndChanged
func GetFileManagedsByLangcodeAndFilemimeAndChanged(offset int, limit int, Langcode_ string, Filemime_ string, Changed_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("langcode = ? and filemime = ? and changed = ?", Langcode_, Filemime_, Changed_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByLangcodeAndFilesizeAndStatus Get FileManageds via LangcodeAndFilesizeAndStatus
func GetFileManagedsByLangcodeAndFilesizeAndStatus(offset int, limit int, Langcode_ string, Filesize_ int64, Status_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("langcode = ? and filesize = ? and status = ?", Langcode_, Filesize_, Status_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByLangcodeAndFilesizeAndCreated Get FileManageds via LangcodeAndFilesizeAndCreated
func GetFileManagedsByLangcodeAndFilesizeAndCreated(offset int, limit int, Langcode_ string, Filesize_ int64, Created_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("langcode = ? and filesize = ? and created = ?", Langcode_, Filesize_, Created_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByLangcodeAndFilesizeAndChanged Get FileManageds via LangcodeAndFilesizeAndChanged
func GetFileManagedsByLangcodeAndFilesizeAndChanged(offset int, limit int, Langcode_ string, Filesize_ int64, Changed_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("langcode = ? and filesize = ? and changed = ?", Langcode_, Filesize_, Changed_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByLangcodeAndStatusAndCreated Get FileManageds via LangcodeAndStatusAndCreated
func GetFileManagedsByLangcodeAndStatusAndCreated(offset int, limit int, Langcode_ string, Status_ int, Created_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("langcode = ? and status = ? and created = ?", Langcode_, Status_, Created_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByLangcodeAndStatusAndChanged Get FileManageds via LangcodeAndStatusAndChanged
func GetFileManagedsByLangcodeAndStatusAndChanged(offset int, limit int, Langcode_ string, Status_ int, Changed_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("langcode = ? and status = ? and changed = ?", Langcode_, Status_, Changed_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByLangcodeAndCreatedAndChanged Get FileManageds via LangcodeAndCreatedAndChanged
func GetFileManagedsByLangcodeAndCreatedAndChanged(offset int, limit int, Langcode_ string, Created_ int, Changed_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("langcode = ? and created = ? and changed = ?", Langcode_, Created_, Changed_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByUidAndFilenameAndUri Get FileManageds via UidAndFilenameAndUri
func GetFileManagedsByUidAndFilenameAndUri(offset int, limit int, Uid_ int, Filename_ string, Uri_ string) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("uid = ? and filename = ? and uri = ?", Uid_, Filename_, Uri_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByUidAndFilenameAndFilemime Get FileManageds via UidAndFilenameAndFilemime
func GetFileManagedsByUidAndFilenameAndFilemime(offset int, limit int, Uid_ int, Filename_ string, Filemime_ string) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("uid = ? and filename = ? and filemime = ?", Uid_, Filename_, Filemime_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByUidAndFilenameAndFilesize Get FileManageds via UidAndFilenameAndFilesize
func GetFileManagedsByUidAndFilenameAndFilesize(offset int, limit int, Uid_ int, Filename_ string, Filesize_ int64) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("uid = ? and filename = ? and filesize = ?", Uid_, Filename_, Filesize_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByUidAndFilenameAndStatus Get FileManageds via UidAndFilenameAndStatus
func GetFileManagedsByUidAndFilenameAndStatus(offset int, limit int, Uid_ int, Filename_ string, Status_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("uid = ? and filename = ? and status = ?", Uid_, Filename_, Status_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByUidAndFilenameAndCreated Get FileManageds via UidAndFilenameAndCreated
func GetFileManagedsByUidAndFilenameAndCreated(offset int, limit int, Uid_ int, Filename_ string, Created_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("uid = ? and filename = ? and created = ?", Uid_, Filename_, Created_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByUidAndFilenameAndChanged Get FileManageds via UidAndFilenameAndChanged
func GetFileManagedsByUidAndFilenameAndChanged(offset int, limit int, Uid_ int, Filename_ string, Changed_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("uid = ? and filename = ? and changed = ?", Uid_, Filename_, Changed_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByUidAndUriAndFilemime Get FileManageds via UidAndUriAndFilemime
func GetFileManagedsByUidAndUriAndFilemime(offset int, limit int, Uid_ int, Uri_ string, Filemime_ string) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("uid = ? and uri = ? and filemime = ?", Uid_, Uri_, Filemime_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByUidAndUriAndFilesize Get FileManageds via UidAndUriAndFilesize
func GetFileManagedsByUidAndUriAndFilesize(offset int, limit int, Uid_ int, Uri_ string, Filesize_ int64) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("uid = ? and uri = ? and filesize = ?", Uid_, Uri_, Filesize_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByUidAndUriAndStatus Get FileManageds via UidAndUriAndStatus
func GetFileManagedsByUidAndUriAndStatus(offset int, limit int, Uid_ int, Uri_ string, Status_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("uid = ? and uri = ? and status = ?", Uid_, Uri_, Status_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByUidAndUriAndCreated Get FileManageds via UidAndUriAndCreated
func GetFileManagedsByUidAndUriAndCreated(offset int, limit int, Uid_ int, Uri_ string, Created_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("uid = ? and uri = ? and created = ?", Uid_, Uri_, Created_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByUidAndUriAndChanged Get FileManageds via UidAndUriAndChanged
func GetFileManagedsByUidAndUriAndChanged(offset int, limit int, Uid_ int, Uri_ string, Changed_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("uid = ? and uri = ? and changed = ?", Uid_, Uri_, Changed_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByUidAndFilemimeAndFilesize Get FileManageds via UidAndFilemimeAndFilesize
func GetFileManagedsByUidAndFilemimeAndFilesize(offset int, limit int, Uid_ int, Filemime_ string, Filesize_ int64) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("uid = ? and filemime = ? and filesize = ?", Uid_, Filemime_, Filesize_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByUidAndFilemimeAndStatus Get FileManageds via UidAndFilemimeAndStatus
func GetFileManagedsByUidAndFilemimeAndStatus(offset int, limit int, Uid_ int, Filemime_ string, Status_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("uid = ? and filemime = ? and status = ?", Uid_, Filemime_, Status_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByUidAndFilemimeAndCreated Get FileManageds via UidAndFilemimeAndCreated
func GetFileManagedsByUidAndFilemimeAndCreated(offset int, limit int, Uid_ int, Filemime_ string, Created_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("uid = ? and filemime = ? and created = ?", Uid_, Filemime_, Created_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByUidAndFilemimeAndChanged Get FileManageds via UidAndFilemimeAndChanged
func GetFileManagedsByUidAndFilemimeAndChanged(offset int, limit int, Uid_ int, Filemime_ string, Changed_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("uid = ? and filemime = ? and changed = ?", Uid_, Filemime_, Changed_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByUidAndFilesizeAndStatus Get FileManageds via UidAndFilesizeAndStatus
func GetFileManagedsByUidAndFilesizeAndStatus(offset int, limit int, Uid_ int, Filesize_ int64, Status_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("uid = ? and filesize = ? and status = ?", Uid_, Filesize_, Status_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByUidAndFilesizeAndCreated Get FileManageds via UidAndFilesizeAndCreated
func GetFileManagedsByUidAndFilesizeAndCreated(offset int, limit int, Uid_ int, Filesize_ int64, Created_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("uid = ? and filesize = ? and created = ?", Uid_, Filesize_, Created_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByUidAndFilesizeAndChanged Get FileManageds via UidAndFilesizeAndChanged
func GetFileManagedsByUidAndFilesizeAndChanged(offset int, limit int, Uid_ int, Filesize_ int64, Changed_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("uid = ? and filesize = ? and changed = ?", Uid_, Filesize_, Changed_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByUidAndStatusAndCreated Get FileManageds via UidAndStatusAndCreated
func GetFileManagedsByUidAndStatusAndCreated(offset int, limit int, Uid_ int, Status_ int, Created_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("uid = ? and status = ? and created = ?", Uid_, Status_, Created_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByUidAndStatusAndChanged Get FileManageds via UidAndStatusAndChanged
func GetFileManagedsByUidAndStatusAndChanged(offset int, limit int, Uid_ int, Status_ int, Changed_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("uid = ? and status = ? and changed = ?", Uid_, Status_, Changed_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByUidAndCreatedAndChanged Get FileManageds via UidAndCreatedAndChanged
func GetFileManagedsByUidAndCreatedAndChanged(offset int, limit int, Uid_ int, Created_ int, Changed_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("uid = ? and created = ? and changed = ?", Uid_, Created_, Changed_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByFilenameAndUriAndFilemime Get FileManageds via FilenameAndUriAndFilemime
func GetFileManagedsByFilenameAndUriAndFilemime(offset int, limit int, Filename_ string, Uri_ string, Filemime_ string) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("filename = ? and uri = ? and filemime = ?", Filename_, Uri_, Filemime_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByFilenameAndUriAndFilesize Get FileManageds via FilenameAndUriAndFilesize
func GetFileManagedsByFilenameAndUriAndFilesize(offset int, limit int, Filename_ string, Uri_ string, Filesize_ int64) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("filename = ? and uri = ? and filesize = ?", Filename_, Uri_, Filesize_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByFilenameAndUriAndStatus Get FileManageds via FilenameAndUriAndStatus
func GetFileManagedsByFilenameAndUriAndStatus(offset int, limit int, Filename_ string, Uri_ string, Status_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("filename = ? and uri = ? and status = ?", Filename_, Uri_, Status_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByFilenameAndUriAndCreated Get FileManageds via FilenameAndUriAndCreated
func GetFileManagedsByFilenameAndUriAndCreated(offset int, limit int, Filename_ string, Uri_ string, Created_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("filename = ? and uri = ? and created = ?", Filename_, Uri_, Created_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByFilenameAndUriAndChanged Get FileManageds via FilenameAndUriAndChanged
func GetFileManagedsByFilenameAndUriAndChanged(offset int, limit int, Filename_ string, Uri_ string, Changed_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("filename = ? and uri = ? and changed = ?", Filename_, Uri_, Changed_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByFilenameAndFilemimeAndFilesize Get FileManageds via FilenameAndFilemimeAndFilesize
func GetFileManagedsByFilenameAndFilemimeAndFilesize(offset int, limit int, Filename_ string, Filemime_ string, Filesize_ int64) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("filename = ? and filemime = ? and filesize = ?", Filename_, Filemime_, Filesize_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByFilenameAndFilemimeAndStatus Get FileManageds via FilenameAndFilemimeAndStatus
func GetFileManagedsByFilenameAndFilemimeAndStatus(offset int, limit int, Filename_ string, Filemime_ string, Status_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("filename = ? and filemime = ? and status = ?", Filename_, Filemime_, Status_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByFilenameAndFilemimeAndCreated Get FileManageds via FilenameAndFilemimeAndCreated
func GetFileManagedsByFilenameAndFilemimeAndCreated(offset int, limit int, Filename_ string, Filemime_ string, Created_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("filename = ? and filemime = ? and created = ?", Filename_, Filemime_, Created_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByFilenameAndFilemimeAndChanged Get FileManageds via FilenameAndFilemimeAndChanged
func GetFileManagedsByFilenameAndFilemimeAndChanged(offset int, limit int, Filename_ string, Filemime_ string, Changed_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("filename = ? and filemime = ? and changed = ?", Filename_, Filemime_, Changed_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByFilenameAndFilesizeAndStatus Get FileManageds via FilenameAndFilesizeAndStatus
func GetFileManagedsByFilenameAndFilesizeAndStatus(offset int, limit int, Filename_ string, Filesize_ int64, Status_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("filename = ? and filesize = ? and status = ?", Filename_, Filesize_, Status_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByFilenameAndFilesizeAndCreated Get FileManageds via FilenameAndFilesizeAndCreated
func GetFileManagedsByFilenameAndFilesizeAndCreated(offset int, limit int, Filename_ string, Filesize_ int64, Created_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("filename = ? and filesize = ? and created = ?", Filename_, Filesize_, Created_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByFilenameAndFilesizeAndChanged Get FileManageds via FilenameAndFilesizeAndChanged
func GetFileManagedsByFilenameAndFilesizeAndChanged(offset int, limit int, Filename_ string, Filesize_ int64, Changed_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("filename = ? and filesize = ? and changed = ?", Filename_, Filesize_, Changed_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByFilenameAndStatusAndCreated Get FileManageds via FilenameAndStatusAndCreated
func GetFileManagedsByFilenameAndStatusAndCreated(offset int, limit int, Filename_ string, Status_ int, Created_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("filename = ? and status = ? and created = ?", Filename_, Status_, Created_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByFilenameAndStatusAndChanged Get FileManageds via FilenameAndStatusAndChanged
func GetFileManagedsByFilenameAndStatusAndChanged(offset int, limit int, Filename_ string, Status_ int, Changed_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("filename = ? and status = ? and changed = ?", Filename_, Status_, Changed_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByFilenameAndCreatedAndChanged Get FileManageds via FilenameAndCreatedAndChanged
func GetFileManagedsByFilenameAndCreatedAndChanged(offset int, limit int, Filename_ string, Created_ int, Changed_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("filename = ? and created = ? and changed = ?", Filename_, Created_, Changed_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByUriAndFilemimeAndFilesize Get FileManageds via UriAndFilemimeAndFilesize
func GetFileManagedsByUriAndFilemimeAndFilesize(offset int, limit int, Uri_ string, Filemime_ string, Filesize_ int64) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("uri = ? and filemime = ? and filesize = ?", Uri_, Filemime_, Filesize_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByUriAndFilemimeAndStatus Get FileManageds via UriAndFilemimeAndStatus
func GetFileManagedsByUriAndFilemimeAndStatus(offset int, limit int, Uri_ string, Filemime_ string, Status_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("uri = ? and filemime = ? and status = ?", Uri_, Filemime_, Status_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByUriAndFilemimeAndCreated Get FileManageds via UriAndFilemimeAndCreated
func GetFileManagedsByUriAndFilemimeAndCreated(offset int, limit int, Uri_ string, Filemime_ string, Created_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("uri = ? and filemime = ? and created = ?", Uri_, Filemime_, Created_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByUriAndFilemimeAndChanged Get FileManageds via UriAndFilemimeAndChanged
func GetFileManagedsByUriAndFilemimeAndChanged(offset int, limit int, Uri_ string, Filemime_ string, Changed_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("uri = ? and filemime = ? and changed = ?", Uri_, Filemime_, Changed_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByUriAndFilesizeAndStatus Get FileManageds via UriAndFilesizeAndStatus
func GetFileManagedsByUriAndFilesizeAndStatus(offset int, limit int, Uri_ string, Filesize_ int64, Status_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("uri = ? and filesize = ? and status = ?", Uri_, Filesize_, Status_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByUriAndFilesizeAndCreated Get FileManageds via UriAndFilesizeAndCreated
func GetFileManagedsByUriAndFilesizeAndCreated(offset int, limit int, Uri_ string, Filesize_ int64, Created_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("uri = ? and filesize = ? and created = ?", Uri_, Filesize_, Created_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByUriAndFilesizeAndChanged Get FileManageds via UriAndFilesizeAndChanged
func GetFileManagedsByUriAndFilesizeAndChanged(offset int, limit int, Uri_ string, Filesize_ int64, Changed_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("uri = ? and filesize = ? and changed = ?", Uri_, Filesize_, Changed_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByUriAndStatusAndCreated Get FileManageds via UriAndStatusAndCreated
func GetFileManagedsByUriAndStatusAndCreated(offset int, limit int, Uri_ string, Status_ int, Created_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("uri = ? and status = ? and created = ?", Uri_, Status_, Created_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByUriAndStatusAndChanged Get FileManageds via UriAndStatusAndChanged
func GetFileManagedsByUriAndStatusAndChanged(offset int, limit int, Uri_ string, Status_ int, Changed_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("uri = ? and status = ? and changed = ?", Uri_, Status_, Changed_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByUriAndCreatedAndChanged Get FileManageds via UriAndCreatedAndChanged
func GetFileManagedsByUriAndCreatedAndChanged(offset int, limit int, Uri_ string, Created_ int, Changed_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("uri = ? and created = ? and changed = ?", Uri_, Created_, Changed_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByFilemimeAndFilesizeAndStatus Get FileManageds via FilemimeAndFilesizeAndStatus
func GetFileManagedsByFilemimeAndFilesizeAndStatus(offset int, limit int, Filemime_ string, Filesize_ int64, Status_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("filemime = ? and filesize = ? and status = ?", Filemime_, Filesize_, Status_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByFilemimeAndFilesizeAndCreated Get FileManageds via FilemimeAndFilesizeAndCreated
func GetFileManagedsByFilemimeAndFilesizeAndCreated(offset int, limit int, Filemime_ string, Filesize_ int64, Created_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("filemime = ? and filesize = ? and created = ?", Filemime_, Filesize_, Created_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByFilemimeAndFilesizeAndChanged Get FileManageds via FilemimeAndFilesizeAndChanged
func GetFileManagedsByFilemimeAndFilesizeAndChanged(offset int, limit int, Filemime_ string, Filesize_ int64, Changed_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("filemime = ? and filesize = ? and changed = ?", Filemime_, Filesize_, Changed_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByFilemimeAndStatusAndCreated Get FileManageds via FilemimeAndStatusAndCreated
func GetFileManagedsByFilemimeAndStatusAndCreated(offset int, limit int, Filemime_ string, Status_ int, Created_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("filemime = ? and status = ? and created = ?", Filemime_, Status_, Created_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByFilemimeAndStatusAndChanged Get FileManageds via FilemimeAndStatusAndChanged
func GetFileManagedsByFilemimeAndStatusAndChanged(offset int, limit int, Filemime_ string, Status_ int, Changed_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("filemime = ? and status = ? and changed = ?", Filemime_, Status_, Changed_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByFilemimeAndCreatedAndChanged Get FileManageds via FilemimeAndCreatedAndChanged
func GetFileManagedsByFilemimeAndCreatedAndChanged(offset int, limit int, Filemime_ string, Created_ int, Changed_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("filemime = ? and created = ? and changed = ?", Filemime_, Created_, Changed_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByFilesizeAndStatusAndCreated Get FileManageds via FilesizeAndStatusAndCreated
func GetFileManagedsByFilesizeAndStatusAndCreated(offset int, limit int, Filesize_ int64, Status_ int, Created_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("filesize = ? and status = ? and created = ?", Filesize_, Status_, Created_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByFilesizeAndStatusAndChanged Get FileManageds via FilesizeAndStatusAndChanged
func GetFileManagedsByFilesizeAndStatusAndChanged(offset int, limit int, Filesize_ int64, Status_ int, Changed_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("filesize = ? and status = ? and changed = ?", Filesize_, Status_, Changed_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByFilesizeAndCreatedAndChanged Get FileManageds via FilesizeAndCreatedAndChanged
func GetFileManagedsByFilesizeAndCreatedAndChanged(offset int, limit int, Filesize_ int64, Created_ int, Changed_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("filesize = ? and created = ? and changed = ?", Filesize_, Created_, Changed_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByStatusAndCreatedAndChanged Get FileManageds via StatusAndCreatedAndChanged
func GetFileManagedsByStatusAndCreatedAndChanged(offset int, limit int, Status_ int, Created_ int, Changed_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("status = ? and created = ? and changed = ?", Status_, Created_, Changed_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByFidAndUuid Get FileManageds via FidAndUuid
func GetFileManagedsByFidAndUuid(offset int, limit int, Fid_ int64, Uuid_ string) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("fid = ? and uuid = ?", Fid_, Uuid_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByFidAndLangcode Get FileManageds via FidAndLangcode
func GetFileManagedsByFidAndLangcode(offset int, limit int, Fid_ int64, Langcode_ string) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("fid = ? and langcode = ?", Fid_, Langcode_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByFidAndUid Get FileManageds via FidAndUid
func GetFileManagedsByFidAndUid(offset int, limit int, Fid_ int64, Uid_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("fid = ? and uid = ?", Fid_, Uid_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByFidAndFilename Get FileManageds via FidAndFilename
func GetFileManagedsByFidAndFilename(offset int, limit int, Fid_ int64, Filename_ string) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("fid = ? and filename = ?", Fid_, Filename_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByFidAndUri Get FileManageds via FidAndUri
func GetFileManagedsByFidAndUri(offset int, limit int, Fid_ int64, Uri_ string) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("fid = ? and uri = ?", Fid_, Uri_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByFidAndFilemime Get FileManageds via FidAndFilemime
func GetFileManagedsByFidAndFilemime(offset int, limit int, Fid_ int64, Filemime_ string) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("fid = ? and filemime = ?", Fid_, Filemime_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByFidAndFilesize Get FileManageds via FidAndFilesize
func GetFileManagedsByFidAndFilesize(offset int, limit int, Fid_ int64, Filesize_ int64) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("fid = ? and filesize = ?", Fid_, Filesize_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByFidAndStatus Get FileManageds via FidAndStatus
func GetFileManagedsByFidAndStatus(offset int, limit int, Fid_ int64, Status_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("fid = ? and status = ?", Fid_, Status_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByFidAndCreated Get FileManageds via FidAndCreated
func GetFileManagedsByFidAndCreated(offset int, limit int, Fid_ int64, Created_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("fid = ? and created = ?", Fid_, Created_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByFidAndChanged Get FileManageds via FidAndChanged
func GetFileManagedsByFidAndChanged(offset int, limit int, Fid_ int64, Changed_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("fid = ? and changed = ?", Fid_, Changed_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByUuidAndLangcode Get FileManageds via UuidAndLangcode
func GetFileManagedsByUuidAndLangcode(offset int, limit int, Uuid_ string, Langcode_ string) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("uuid = ? and langcode = ?", Uuid_, Langcode_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByUuidAndUid Get FileManageds via UuidAndUid
func GetFileManagedsByUuidAndUid(offset int, limit int, Uuid_ string, Uid_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("uuid = ? and uid = ?", Uuid_, Uid_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByUuidAndFilename Get FileManageds via UuidAndFilename
func GetFileManagedsByUuidAndFilename(offset int, limit int, Uuid_ string, Filename_ string) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("uuid = ? and filename = ?", Uuid_, Filename_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByUuidAndUri Get FileManageds via UuidAndUri
func GetFileManagedsByUuidAndUri(offset int, limit int, Uuid_ string, Uri_ string) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("uuid = ? and uri = ?", Uuid_, Uri_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByUuidAndFilemime Get FileManageds via UuidAndFilemime
func GetFileManagedsByUuidAndFilemime(offset int, limit int, Uuid_ string, Filemime_ string) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("uuid = ? and filemime = ?", Uuid_, Filemime_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByUuidAndFilesize Get FileManageds via UuidAndFilesize
func GetFileManagedsByUuidAndFilesize(offset int, limit int, Uuid_ string, Filesize_ int64) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("uuid = ? and filesize = ?", Uuid_, Filesize_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByUuidAndStatus Get FileManageds via UuidAndStatus
func GetFileManagedsByUuidAndStatus(offset int, limit int, Uuid_ string, Status_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("uuid = ? and status = ?", Uuid_, Status_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByUuidAndCreated Get FileManageds via UuidAndCreated
func GetFileManagedsByUuidAndCreated(offset int, limit int, Uuid_ string, Created_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("uuid = ? and created = ?", Uuid_, Created_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByUuidAndChanged Get FileManageds via UuidAndChanged
func GetFileManagedsByUuidAndChanged(offset int, limit int, Uuid_ string, Changed_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("uuid = ? and changed = ?", Uuid_, Changed_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByLangcodeAndUid Get FileManageds via LangcodeAndUid
func GetFileManagedsByLangcodeAndUid(offset int, limit int, Langcode_ string, Uid_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("langcode = ? and uid = ?", Langcode_, Uid_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByLangcodeAndFilename Get FileManageds via LangcodeAndFilename
func GetFileManagedsByLangcodeAndFilename(offset int, limit int, Langcode_ string, Filename_ string) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("langcode = ? and filename = ?", Langcode_, Filename_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByLangcodeAndUri Get FileManageds via LangcodeAndUri
func GetFileManagedsByLangcodeAndUri(offset int, limit int, Langcode_ string, Uri_ string) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("langcode = ? and uri = ?", Langcode_, Uri_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByLangcodeAndFilemime Get FileManageds via LangcodeAndFilemime
func GetFileManagedsByLangcodeAndFilemime(offset int, limit int, Langcode_ string, Filemime_ string) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("langcode = ? and filemime = ?", Langcode_, Filemime_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByLangcodeAndFilesize Get FileManageds via LangcodeAndFilesize
func GetFileManagedsByLangcodeAndFilesize(offset int, limit int, Langcode_ string, Filesize_ int64) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("langcode = ? and filesize = ?", Langcode_, Filesize_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByLangcodeAndStatus Get FileManageds via LangcodeAndStatus
func GetFileManagedsByLangcodeAndStatus(offset int, limit int, Langcode_ string, Status_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("langcode = ? and status = ?", Langcode_, Status_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByLangcodeAndCreated Get FileManageds via LangcodeAndCreated
func GetFileManagedsByLangcodeAndCreated(offset int, limit int, Langcode_ string, Created_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("langcode = ? and created = ?", Langcode_, Created_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByLangcodeAndChanged Get FileManageds via LangcodeAndChanged
func GetFileManagedsByLangcodeAndChanged(offset int, limit int, Langcode_ string, Changed_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("langcode = ? and changed = ?", Langcode_, Changed_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByUidAndFilename Get FileManageds via UidAndFilename
func GetFileManagedsByUidAndFilename(offset int, limit int, Uid_ int, Filename_ string) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("uid = ? and filename = ?", Uid_, Filename_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByUidAndUri Get FileManageds via UidAndUri
func GetFileManagedsByUidAndUri(offset int, limit int, Uid_ int, Uri_ string) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("uid = ? and uri = ?", Uid_, Uri_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByUidAndFilemime Get FileManageds via UidAndFilemime
func GetFileManagedsByUidAndFilemime(offset int, limit int, Uid_ int, Filemime_ string) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("uid = ? and filemime = ?", Uid_, Filemime_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByUidAndFilesize Get FileManageds via UidAndFilesize
func GetFileManagedsByUidAndFilesize(offset int, limit int, Uid_ int, Filesize_ int64) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("uid = ? and filesize = ?", Uid_, Filesize_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByUidAndStatus Get FileManageds via UidAndStatus
func GetFileManagedsByUidAndStatus(offset int, limit int, Uid_ int, Status_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("uid = ? and status = ?", Uid_, Status_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByUidAndCreated Get FileManageds via UidAndCreated
func GetFileManagedsByUidAndCreated(offset int, limit int, Uid_ int, Created_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("uid = ? and created = ?", Uid_, Created_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByUidAndChanged Get FileManageds via UidAndChanged
func GetFileManagedsByUidAndChanged(offset int, limit int, Uid_ int, Changed_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("uid = ? and changed = ?", Uid_, Changed_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByFilenameAndUri Get FileManageds via FilenameAndUri
func GetFileManagedsByFilenameAndUri(offset int, limit int, Filename_ string, Uri_ string) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("filename = ? and uri = ?", Filename_, Uri_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByFilenameAndFilemime Get FileManageds via FilenameAndFilemime
func GetFileManagedsByFilenameAndFilemime(offset int, limit int, Filename_ string, Filemime_ string) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("filename = ? and filemime = ?", Filename_, Filemime_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByFilenameAndFilesize Get FileManageds via FilenameAndFilesize
func GetFileManagedsByFilenameAndFilesize(offset int, limit int, Filename_ string, Filesize_ int64) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("filename = ? and filesize = ?", Filename_, Filesize_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByFilenameAndStatus Get FileManageds via FilenameAndStatus
func GetFileManagedsByFilenameAndStatus(offset int, limit int, Filename_ string, Status_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("filename = ? and status = ?", Filename_, Status_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByFilenameAndCreated Get FileManageds via FilenameAndCreated
func GetFileManagedsByFilenameAndCreated(offset int, limit int, Filename_ string, Created_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("filename = ? and created = ?", Filename_, Created_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByFilenameAndChanged Get FileManageds via FilenameAndChanged
func GetFileManagedsByFilenameAndChanged(offset int, limit int, Filename_ string, Changed_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("filename = ? and changed = ?", Filename_, Changed_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByUriAndFilemime Get FileManageds via UriAndFilemime
func GetFileManagedsByUriAndFilemime(offset int, limit int, Uri_ string, Filemime_ string) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("uri = ? and filemime = ?", Uri_, Filemime_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByUriAndFilesize Get FileManageds via UriAndFilesize
func GetFileManagedsByUriAndFilesize(offset int, limit int, Uri_ string, Filesize_ int64) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("uri = ? and filesize = ?", Uri_, Filesize_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByUriAndStatus Get FileManageds via UriAndStatus
func GetFileManagedsByUriAndStatus(offset int, limit int, Uri_ string, Status_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("uri = ? and status = ?", Uri_, Status_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByUriAndCreated Get FileManageds via UriAndCreated
func GetFileManagedsByUriAndCreated(offset int, limit int, Uri_ string, Created_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("uri = ? and created = ?", Uri_, Created_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByUriAndChanged Get FileManageds via UriAndChanged
func GetFileManagedsByUriAndChanged(offset int, limit int, Uri_ string, Changed_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("uri = ? and changed = ?", Uri_, Changed_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByFilemimeAndFilesize Get FileManageds via FilemimeAndFilesize
func GetFileManagedsByFilemimeAndFilesize(offset int, limit int, Filemime_ string, Filesize_ int64) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("filemime = ? and filesize = ?", Filemime_, Filesize_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByFilemimeAndStatus Get FileManageds via FilemimeAndStatus
func GetFileManagedsByFilemimeAndStatus(offset int, limit int, Filemime_ string, Status_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("filemime = ? and status = ?", Filemime_, Status_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByFilemimeAndCreated Get FileManageds via FilemimeAndCreated
func GetFileManagedsByFilemimeAndCreated(offset int, limit int, Filemime_ string, Created_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("filemime = ? and created = ?", Filemime_, Created_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByFilemimeAndChanged Get FileManageds via FilemimeAndChanged
func GetFileManagedsByFilemimeAndChanged(offset int, limit int, Filemime_ string, Changed_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("filemime = ? and changed = ?", Filemime_, Changed_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByFilesizeAndStatus Get FileManageds via FilesizeAndStatus
func GetFileManagedsByFilesizeAndStatus(offset int, limit int, Filesize_ int64, Status_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("filesize = ? and status = ?", Filesize_, Status_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByFilesizeAndCreated Get FileManageds via FilesizeAndCreated
func GetFileManagedsByFilesizeAndCreated(offset int, limit int, Filesize_ int64, Created_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("filesize = ? and created = ?", Filesize_, Created_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByFilesizeAndChanged Get FileManageds via FilesizeAndChanged
func GetFileManagedsByFilesizeAndChanged(offset int, limit int, Filesize_ int64, Changed_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("filesize = ? and changed = ?", Filesize_, Changed_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByStatusAndCreated Get FileManageds via StatusAndCreated
func GetFileManagedsByStatusAndCreated(offset int, limit int, Status_ int, Created_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("status = ? and created = ?", Status_, Created_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByStatusAndChanged Get FileManageds via StatusAndChanged
func GetFileManagedsByStatusAndChanged(offset int, limit int, Status_ int, Changed_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("status = ? and changed = ?", Status_, Changed_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsByCreatedAndChanged Get FileManageds via CreatedAndChanged
func GetFileManagedsByCreatedAndChanged(offset int, limit int, Created_ int, Changed_ int) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("created = ? and changed = ?", Created_, Changed_).Limit(limit, offset).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManageds Get FileManageds via field
func GetFileManageds(offset int, limit int, field string) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Limit(limit, offset).Desc(field).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsViaFid Get FileManageds via Fid
func GetFileManagedsViaFid(offset int, limit int, Fid_ int64, field string) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("fid = ?", Fid_).Limit(limit, offset).Desc(field).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsViaUuid Get FileManageds via Uuid
func GetFileManagedsViaUuid(offset int, limit int, Uuid_ string, field string) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("uuid = ?", Uuid_).Limit(limit, offset).Desc(field).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsViaLangcode Get FileManageds via Langcode
func GetFileManagedsViaLangcode(offset int, limit int, Langcode_ string, field string) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("langcode = ?", Langcode_).Limit(limit, offset).Desc(field).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsViaUid Get FileManageds via Uid
func GetFileManagedsViaUid(offset int, limit int, Uid_ int, field string) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("uid = ?", Uid_).Limit(limit, offset).Desc(field).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsViaFilename Get FileManageds via Filename
func GetFileManagedsViaFilename(offset int, limit int, Filename_ string, field string) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("filename = ?", Filename_).Limit(limit, offset).Desc(field).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsViaUri Get FileManageds via Uri
func GetFileManagedsViaUri(offset int, limit int, Uri_ string, field string) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("uri = ?", Uri_).Limit(limit, offset).Desc(field).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsViaFilemime Get FileManageds via Filemime
func GetFileManagedsViaFilemime(offset int, limit int, Filemime_ string, field string) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("filemime = ?", Filemime_).Limit(limit, offset).Desc(field).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsViaFilesize Get FileManageds via Filesize
func GetFileManagedsViaFilesize(offset int, limit int, Filesize_ int64, field string) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("filesize = ?", Filesize_).Limit(limit, offset).Desc(field).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsViaStatus Get FileManageds via Status
func GetFileManagedsViaStatus(offset int, limit int, Status_ int, field string) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("status = ?", Status_).Limit(limit, offset).Desc(field).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsViaCreated Get FileManageds via Created
func GetFileManagedsViaCreated(offset int, limit int, Created_ int, field string) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("created = ?", Created_).Limit(limit, offset).Desc(field).Find(_FileManaged)
	return _FileManaged, err
}

// GetFileManagedsViaChanged Get FileManageds via Changed
func GetFileManagedsViaChanged(offset int, limit int, Changed_ int, field string) (*[]*FileManaged, error) {
	var _FileManaged = new([]*FileManaged)
	err := Engine.Table("file_managed").Where("changed = ?", Changed_).Limit(limit, offset).Desc(field).Find(_FileManaged)
	return _FileManaged, err
}

// HasFileManagedViaFid Has FileManaged via Fid
func HasFileManagedViaFid(iFid int64) bool {
	if has, err := Engine.Where("fid = ?", iFid).Get(new(FileManaged)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasFileManagedViaUuid Has FileManaged via Uuid
func HasFileManagedViaUuid(iUuid string) bool {
	if has, err := Engine.Where("uuid = ?", iUuid).Get(new(FileManaged)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasFileManagedViaLangcode Has FileManaged via Langcode
func HasFileManagedViaLangcode(iLangcode string) bool {
	if has, err := Engine.Where("langcode = ?", iLangcode).Get(new(FileManaged)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasFileManagedViaUid Has FileManaged via Uid
func HasFileManagedViaUid(iUid int) bool {
	if has, err := Engine.Where("uid = ?", iUid).Get(new(FileManaged)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasFileManagedViaFilename Has FileManaged via Filename
func HasFileManagedViaFilename(iFilename string) bool {
	if has, err := Engine.Where("filename = ?", iFilename).Get(new(FileManaged)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasFileManagedViaUri Has FileManaged via Uri
func HasFileManagedViaUri(iUri string) bool {
	if has, err := Engine.Where("uri = ?", iUri).Get(new(FileManaged)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasFileManagedViaFilemime Has FileManaged via Filemime
func HasFileManagedViaFilemime(iFilemime string) bool {
	if has, err := Engine.Where("filemime = ?", iFilemime).Get(new(FileManaged)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasFileManagedViaFilesize Has FileManaged via Filesize
func HasFileManagedViaFilesize(iFilesize int64) bool {
	if has, err := Engine.Where("filesize = ?", iFilesize).Get(new(FileManaged)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasFileManagedViaStatus Has FileManaged via Status
func HasFileManagedViaStatus(iStatus int) bool {
	if has, err := Engine.Where("status = ?", iStatus).Get(new(FileManaged)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasFileManagedViaCreated Has FileManaged via Created
func HasFileManagedViaCreated(iCreated int) bool {
	if has, err := Engine.Where("created = ?", iCreated).Get(new(FileManaged)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasFileManagedViaChanged Has FileManaged via Changed
func HasFileManagedViaChanged(iChanged int) bool {
	if has, err := Engine.Where("changed = ?", iChanged).Get(new(FileManaged)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// GetFileManagedViaFid Get FileManaged via Fid
func GetFileManagedViaFid(iFid int64) (*FileManaged, error) {
	var _FileManaged = &FileManaged{Fid: iFid}
	has, err := Engine.Get(_FileManaged)
	if has {
		return _FileManaged, err
	} else {
		return nil, err
	}
}

// GetFileManagedViaUuid Get FileManaged via Uuid
func GetFileManagedViaUuid(iUuid string) (*FileManaged, error) {
	var _FileManaged = &FileManaged{Uuid: iUuid}
	has, err := Engine.Get(_FileManaged)
	if has {
		return _FileManaged, err
	} else {
		return nil, err
	}
}

// GetFileManagedViaLangcode Get FileManaged via Langcode
func GetFileManagedViaLangcode(iLangcode string) (*FileManaged, error) {
	var _FileManaged = &FileManaged{Langcode: iLangcode}
	has, err := Engine.Get(_FileManaged)
	if has {
		return _FileManaged, err
	} else {
		return nil, err
	}
}

// GetFileManagedViaUid Get FileManaged via Uid
func GetFileManagedViaUid(iUid int) (*FileManaged, error) {
	var _FileManaged = &FileManaged{Uid: iUid}
	has, err := Engine.Get(_FileManaged)
	if has {
		return _FileManaged, err
	} else {
		return nil, err
	}
}

// GetFileManagedViaFilename Get FileManaged via Filename
func GetFileManagedViaFilename(iFilename string) (*FileManaged, error) {
	var _FileManaged = &FileManaged{Filename: iFilename}
	has, err := Engine.Get(_FileManaged)
	if has {
		return _FileManaged, err
	} else {
		return nil, err
	}
}

// GetFileManagedViaUri Get FileManaged via Uri
func GetFileManagedViaUri(iUri string) (*FileManaged, error) {
	var _FileManaged = &FileManaged{Uri: iUri}
	has, err := Engine.Get(_FileManaged)
	if has {
		return _FileManaged, err
	} else {
		return nil, err
	}
}

// GetFileManagedViaFilemime Get FileManaged via Filemime
func GetFileManagedViaFilemime(iFilemime string) (*FileManaged, error) {
	var _FileManaged = &FileManaged{Filemime: iFilemime}
	has, err := Engine.Get(_FileManaged)
	if has {
		return _FileManaged, err
	} else {
		return nil, err
	}
}

// GetFileManagedViaFilesize Get FileManaged via Filesize
func GetFileManagedViaFilesize(iFilesize int64) (*FileManaged, error) {
	var _FileManaged = &FileManaged{Filesize: iFilesize}
	has, err := Engine.Get(_FileManaged)
	if has {
		return _FileManaged, err
	} else {
		return nil, err
	}
}

// GetFileManagedViaStatus Get FileManaged via Status
func GetFileManagedViaStatus(iStatus int) (*FileManaged, error) {
	var _FileManaged = &FileManaged{Status: iStatus}
	has, err := Engine.Get(_FileManaged)
	if has {
		return _FileManaged, err
	} else {
		return nil, err
	}
}

// GetFileManagedViaCreated Get FileManaged via Created
func GetFileManagedViaCreated(iCreated int) (*FileManaged, error) {
	var _FileManaged = &FileManaged{Created: iCreated}
	has, err := Engine.Get(_FileManaged)
	if has {
		return _FileManaged, err
	} else {
		return nil, err
	}
}

// GetFileManagedViaChanged Get FileManaged via Changed
func GetFileManagedViaChanged(iChanged int) (*FileManaged, error) {
	var _FileManaged = &FileManaged{Changed: iChanged}
	has, err := Engine.Get(_FileManaged)
	if has {
		return _FileManaged, err
	} else {
		return nil, err
	}
}

// SetFileManagedViaFid Set FileManaged via Fid
func SetFileManagedViaFid(iFid int64, file_managed *FileManaged) (int64, error) {
	file_managed.Fid = iFid
	return Engine.Insert(file_managed)
}

// SetFileManagedViaUuid Set FileManaged via Uuid
func SetFileManagedViaUuid(iUuid string, file_managed *FileManaged) (int64, error) {
	file_managed.Uuid = iUuid
	return Engine.Insert(file_managed)
}

// SetFileManagedViaLangcode Set FileManaged via Langcode
func SetFileManagedViaLangcode(iLangcode string, file_managed *FileManaged) (int64, error) {
	file_managed.Langcode = iLangcode
	return Engine.Insert(file_managed)
}

// SetFileManagedViaUid Set FileManaged via Uid
func SetFileManagedViaUid(iUid int, file_managed *FileManaged) (int64, error) {
	file_managed.Uid = iUid
	return Engine.Insert(file_managed)
}

// SetFileManagedViaFilename Set FileManaged via Filename
func SetFileManagedViaFilename(iFilename string, file_managed *FileManaged) (int64, error) {
	file_managed.Filename = iFilename
	return Engine.Insert(file_managed)
}

// SetFileManagedViaUri Set FileManaged via Uri
func SetFileManagedViaUri(iUri string, file_managed *FileManaged) (int64, error) {
	file_managed.Uri = iUri
	return Engine.Insert(file_managed)
}

// SetFileManagedViaFilemime Set FileManaged via Filemime
func SetFileManagedViaFilemime(iFilemime string, file_managed *FileManaged) (int64, error) {
	file_managed.Filemime = iFilemime
	return Engine.Insert(file_managed)
}

// SetFileManagedViaFilesize Set FileManaged via Filesize
func SetFileManagedViaFilesize(iFilesize int64, file_managed *FileManaged) (int64, error) {
	file_managed.Filesize = iFilesize
	return Engine.Insert(file_managed)
}

// SetFileManagedViaStatus Set FileManaged via Status
func SetFileManagedViaStatus(iStatus int, file_managed *FileManaged) (int64, error) {
	file_managed.Status = iStatus
	return Engine.Insert(file_managed)
}

// SetFileManagedViaCreated Set FileManaged via Created
func SetFileManagedViaCreated(iCreated int, file_managed *FileManaged) (int64, error) {
	file_managed.Created = iCreated
	return Engine.Insert(file_managed)
}

// SetFileManagedViaChanged Set FileManaged via Changed
func SetFileManagedViaChanged(iChanged int, file_managed *FileManaged) (int64, error) {
	file_managed.Changed = iChanged
	return Engine.Insert(file_managed)
}

// AddFileManaged Add FileManaged via all columns
func AddFileManaged(iFid int64, iUuid string, iLangcode string, iUid int, iFilename string, iUri string, iFilemime string, iFilesize int64, iStatus int, iCreated int, iChanged int) error {
	_FileManaged := &FileManaged{Fid: iFid, Uuid: iUuid, Langcode: iLangcode, Uid: iUid, Filename: iFilename, Uri: iUri, Filemime: iFilemime, Filesize: iFilesize, Status: iStatus, Created: iCreated, Changed: iChanged}
	if _, err := Engine.Insert(_FileManaged); err != nil {
		return err
	} else {
		return nil
	}
}

// PostFileManaged Post FileManaged via iFileManaged
func PostFileManaged(iFileManaged *FileManaged) (int64, error) {
	_, err := Engine.Insert(iFileManaged)
	return iFileManaged.Fid, err
}

// PutFileManaged Put FileManaged
func PutFileManaged(iFileManaged *FileManaged) (int64, error) {
	_, err := Engine.Id(iFileManaged.Fid).Update(iFileManaged)
	return iFileManaged.Fid, err
}

// PutFileManagedViaFid Put FileManaged via Fid
func PutFileManagedViaFid(Fid_ int64, iFileManaged *FileManaged) (int64, error) {
	row, err := Engine.Update(iFileManaged, &FileManaged{Fid: Fid_})
	return row, err
}

// PutFileManagedViaUuid Put FileManaged via Uuid
func PutFileManagedViaUuid(Uuid_ string, iFileManaged *FileManaged) (int64, error) {
	row, err := Engine.Update(iFileManaged, &FileManaged{Uuid: Uuid_})
	return row, err
}

// PutFileManagedViaLangcode Put FileManaged via Langcode
func PutFileManagedViaLangcode(Langcode_ string, iFileManaged *FileManaged) (int64, error) {
	row, err := Engine.Update(iFileManaged, &FileManaged{Langcode: Langcode_})
	return row, err
}

// PutFileManagedViaUid Put FileManaged via Uid
func PutFileManagedViaUid(Uid_ int, iFileManaged *FileManaged) (int64, error) {
	row, err := Engine.Update(iFileManaged, &FileManaged{Uid: Uid_})
	return row, err
}

// PutFileManagedViaFilename Put FileManaged via Filename
func PutFileManagedViaFilename(Filename_ string, iFileManaged *FileManaged) (int64, error) {
	row, err := Engine.Update(iFileManaged, &FileManaged{Filename: Filename_})
	return row, err
}

// PutFileManagedViaUri Put FileManaged via Uri
func PutFileManagedViaUri(Uri_ string, iFileManaged *FileManaged) (int64, error) {
	row, err := Engine.Update(iFileManaged, &FileManaged{Uri: Uri_})
	return row, err
}

// PutFileManagedViaFilemime Put FileManaged via Filemime
func PutFileManagedViaFilemime(Filemime_ string, iFileManaged *FileManaged) (int64, error) {
	row, err := Engine.Update(iFileManaged, &FileManaged{Filemime: Filemime_})
	return row, err
}

// PutFileManagedViaFilesize Put FileManaged via Filesize
func PutFileManagedViaFilesize(Filesize_ int64, iFileManaged *FileManaged) (int64, error) {
	row, err := Engine.Update(iFileManaged, &FileManaged{Filesize: Filesize_})
	return row, err
}

// PutFileManagedViaStatus Put FileManaged via Status
func PutFileManagedViaStatus(Status_ int, iFileManaged *FileManaged) (int64, error) {
	row, err := Engine.Update(iFileManaged, &FileManaged{Status: Status_})
	return row, err
}

// PutFileManagedViaCreated Put FileManaged via Created
func PutFileManagedViaCreated(Created_ int, iFileManaged *FileManaged) (int64, error) {
	row, err := Engine.Update(iFileManaged, &FileManaged{Created: Created_})
	return row, err
}

// PutFileManagedViaChanged Put FileManaged via Changed
func PutFileManagedViaChanged(Changed_ int, iFileManaged *FileManaged) (int64, error) {
	row, err := Engine.Update(iFileManaged, &FileManaged{Changed: Changed_})
	return row, err
}

// UpdateFileManagedViaFid via map[string]interface{}{}
func UpdateFileManagedViaFid(iFid int64, iFileManagedMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(FileManaged)).Where("fid = ?", iFid).Update(iFileManagedMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateFileManagedViaUuid via map[string]interface{}{}
func UpdateFileManagedViaUuid(iUuid string, iFileManagedMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(FileManaged)).Where("uuid = ?", iUuid).Update(iFileManagedMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateFileManagedViaLangcode via map[string]interface{}{}
func UpdateFileManagedViaLangcode(iLangcode string, iFileManagedMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(FileManaged)).Where("langcode = ?", iLangcode).Update(iFileManagedMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateFileManagedViaUid via map[string]interface{}{}
func UpdateFileManagedViaUid(iUid int, iFileManagedMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(FileManaged)).Where("uid = ?", iUid).Update(iFileManagedMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateFileManagedViaFilename via map[string]interface{}{}
func UpdateFileManagedViaFilename(iFilename string, iFileManagedMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(FileManaged)).Where("filename = ?", iFilename).Update(iFileManagedMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateFileManagedViaUri via map[string]interface{}{}
func UpdateFileManagedViaUri(iUri string, iFileManagedMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(FileManaged)).Where("uri = ?", iUri).Update(iFileManagedMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateFileManagedViaFilemime via map[string]interface{}{}
func UpdateFileManagedViaFilemime(iFilemime string, iFileManagedMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(FileManaged)).Where("filemime = ?", iFilemime).Update(iFileManagedMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateFileManagedViaFilesize via map[string]interface{}{}
func UpdateFileManagedViaFilesize(iFilesize int64, iFileManagedMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(FileManaged)).Where("filesize = ?", iFilesize).Update(iFileManagedMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateFileManagedViaStatus via map[string]interface{}{}
func UpdateFileManagedViaStatus(iStatus int, iFileManagedMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(FileManaged)).Where("status = ?", iStatus).Update(iFileManagedMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateFileManagedViaCreated via map[string]interface{}{}
func UpdateFileManagedViaCreated(iCreated int, iFileManagedMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(FileManaged)).Where("created = ?", iCreated).Update(iFileManagedMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateFileManagedViaChanged via map[string]interface{}{}
func UpdateFileManagedViaChanged(iChanged int, iFileManagedMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(FileManaged)).Where("changed = ?", iChanged).Update(iFileManagedMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// DeleteFileManaged Delete FileManaged
func DeleteFileManaged(iFid int64) (int64, error) {
	row, err := Engine.Id(iFid).Delete(new(FileManaged))
	return row, err
}

// DeleteFileManagedViaFid Delete FileManaged via Fid
func DeleteFileManagedViaFid(iFid int64) (err error) {
	var has bool
	var _FileManaged = &FileManaged{Fid: iFid}
	if has, err = Engine.Get(_FileManaged); (has == true) && (err == nil) {
		if row, err := Engine.Where("fid = ?", iFid).Delete(new(FileManaged)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteFileManagedViaUuid Delete FileManaged via Uuid
func DeleteFileManagedViaUuid(iUuid string) (err error) {
	var has bool
	var _FileManaged = &FileManaged{Uuid: iUuid}
	if has, err = Engine.Get(_FileManaged); (has == true) && (err == nil) {
		if row, err := Engine.Where("uuid = ?", iUuid).Delete(new(FileManaged)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteFileManagedViaLangcode Delete FileManaged via Langcode
func DeleteFileManagedViaLangcode(iLangcode string) (err error) {
	var has bool
	var _FileManaged = &FileManaged{Langcode: iLangcode}
	if has, err = Engine.Get(_FileManaged); (has == true) && (err == nil) {
		if row, err := Engine.Where("langcode = ?", iLangcode).Delete(new(FileManaged)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteFileManagedViaUid Delete FileManaged via Uid
func DeleteFileManagedViaUid(iUid int) (err error) {
	var has bool
	var _FileManaged = &FileManaged{Uid: iUid}
	if has, err = Engine.Get(_FileManaged); (has == true) && (err == nil) {
		if row, err := Engine.Where("uid = ?", iUid).Delete(new(FileManaged)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteFileManagedViaFilename Delete FileManaged via Filename
func DeleteFileManagedViaFilename(iFilename string) (err error) {
	var has bool
	var _FileManaged = &FileManaged{Filename: iFilename}
	if has, err = Engine.Get(_FileManaged); (has == true) && (err == nil) {
		if row, err := Engine.Where("filename = ?", iFilename).Delete(new(FileManaged)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteFileManagedViaUri Delete FileManaged via Uri
func DeleteFileManagedViaUri(iUri string) (err error) {
	var has bool
	var _FileManaged = &FileManaged{Uri: iUri}
	if has, err = Engine.Get(_FileManaged); (has == true) && (err == nil) {
		if row, err := Engine.Where("uri = ?", iUri).Delete(new(FileManaged)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteFileManagedViaFilemime Delete FileManaged via Filemime
func DeleteFileManagedViaFilemime(iFilemime string) (err error) {
	var has bool
	var _FileManaged = &FileManaged{Filemime: iFilemime}
	if has, err = Engine.Get(_FileManaged); (has == true) && (err == nil) {
		if row, err := Engine.Where("filemime = ?", iFilemime).Delete(new(FileManaged)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteFileManagedViaFilesize Delete FileManaged via Filesize
func DeleteFileManagedViaFilesize(iFilesize int64) (err error) {
	var has bool
	var _FileManaged = &FileManaged{Filesize: iFilesize}
	if has, err = Engine.Get(_FileManaged); (has == true) && (err == nil) {
		if row, err := Engine.Where("filesize = ?", iFilesize).Delete(new(FileManaged)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteFileManagedViaStatus Delete FileManaged via Status
func DeleteFileManagedViaStatus(iStatus int) (err error) {
	var has bool
	var _FileManaged = &FileManaged{Status: iStatus}
	if has, err = Engine.Get(_FileManaged); (has == true) && (err == nil) {
		if row, err := Engine.Where("status = ?", iStatus).Delete(new(FileManaged)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteFileManagedViaCreated Delete FileManaged via Created
func DeleteFileManagedViaCreated(iCreated int) (err error) {
	var has bool
	var _FileManaged = &FileManaged{Created: iCreated}
	if has, err = Engine.Get(_FileManaged); (has == true) && (err == nil) {
		if row, err := Engine.Where("created = ?", iCreated).Delete(new(FileManaged)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteFileManagedViaChanged Delete FileManaged via Changed
func DeleteFileManagedViaChanged(iChanged int) (err error) {
	var has bool
	var _FileManaged = &FileManaged{Changed: iChanged}
	if has, err = Engine.Get(_FileManaged); (has == true) && (err == nil) {
		if row, err := Engine.Where("changed = ?", iChanged).Delete(new(FileManaged)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}
