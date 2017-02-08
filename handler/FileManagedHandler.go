package handler

import (
	"drupapp/helper"
	"drupapp/model"
	"github.com/insionng/macross"
)

func GetFileManagedsCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()

	if offset > 0 {
		var m = map[string]interface{}{}
		_int64, _error := model.GetFileManagedsCount(offset, limit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		m["file_managedsCount"] = _int64
		return self.JSON(m)
	}
	herr.Message = "Can't get to the GetFileManagedsCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedCountViaFidHandler(self *macross.Context) error {
	Fid_ := self.Args("fid").MustInt64()
	_int64 := model.GetFileManagedCountViaFid(Fid_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["file_managedCount"] = 0
	}
	m["file_managedCount"] = _int64
	return self.JSON(m)
}

func GetFileManagedCountViaUuidHandler(self *macross.Context) error {
	Uuid_ := self.Args("uuid").String()
	_int64 := model.GetFileManagedCountViaUuid(Uuid_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["file_managedCount"] = 0
	}
	m["file_managedCount"] = _int64
	return self.JSON(m)
}

func GetFileManagedCountViaLangcodeHandler(self *macross.Context) error {
	Langcode_ := self.Args("langcode").String()
	_int64 := model.GetFileManagedCountViaLangcode(Langcode_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["file_managedCount"] = 0
	}
	m["file_managedCount"] = _int64
	return self.JSON(m)
}

func GetFileManagedCountViaUidHandler(self *macross.Context) error {
	Uid_ := self.Args("uid").MustInt()
	_int64 := model.GetFileManagedCountViaUid(Uid_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["file_managedCount"] = 0
	}
	m["file_managedCount"] = _int64
	return self.JSON(m)
}

func GetFileManagedCountViaFilenameHandler(self *macross.Context) error {
	Filename_ := self.Args("filename").String()
	_int64 := model.GetFileManagedCountViaFilename(Filename_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["file_managedCount"] = 0
	}
	m["file_managedCount"] = _int64
	return self.JSON(m)
}

func GetFileManagedCountViaUriHandler(self *macross.Context) error {
	Uri_ := self.Args("uri").String()
	_int64 := model.GetFileManagedCountViaUri(Uri_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["file_managedCount"] = 0
	}
	m["file_managedCount"] = _int64
	return self.JSON(m)
}

func GetFileManagedCountViaFilemimeHandler(self *macross.Context) error {
	Filemime_ := self.Args("filemime").String()
	_int64 := model.GetFileManagedCountViaFilemime(Filemime_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["file_managedCount"] = 0
	}
	m["file_managedCount"] = _int64
	return self.JSON(m)
}

func GetFileManagedCountViaFilesizeHandler(self *macross.Context) error {
	Filesize_ := self.Args("filesize").MustInt64()
	_int64 := model.GetFileManagedCountViaFilesize(Filesize_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["file_managedCount"] = 0
	}
	m["file_managedCount"] = _int64
	return self.JSON(m)
}

func GetFileManagedCountViaStatusHandler(self *macross.Context) error {
	Status_ := self.Args("status").MustInt()
	_int64 := model.GetFileManagedCountViaStatus(Status_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["file_managedCount"] = 0
	}
	m["file_managedCount"] = _int64
	return self.JSON(m)
}

func GetFileManagedCountViaCreatedHandler(self *macross.Context) error {
	Created_ := self.Args("created").MustInt()
	_int64 := model.GetFileManagedCountViaCreated(Created_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["file_managedCount"] = 0
	}
	m["file_managedCount"] = _int64
	return self.JSON(m)
}

func GetFileManagedCountViaChangedHandler(self *macross.Context) error {
	Changed_ := self.Args("changed").MustInt()
	_int64 := model.GetFileManagedCountViaChanged(Changed_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["file_managedCount"] = 0
	}
	m["file_managedCount"] = _int64
	return self.JSON(m)
}

func GetFileManagedsViaFidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iFid := self.Args("fid").MustInt64()
	if (offset > 0) && helper.IsHas(iFid) {
		_FileManaged, _error := model.GetFileManagedsViaFid(offset, limit, iFid, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsViaFid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsViaUuidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iUuid := self.Args("uuid").String()
	if (offset > 0) && helper.IsHas(iUuid) {
		_FileManaged, _error := model.GetFileManagedsViaUuid(offset, limit, iUuid, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsViaUuid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iLangcode := self.Args("langcode").String()
	if (offset > 0) && helper.IsHas(iLangcode) {
		_FileManaged, _error := model.GetFileManagedsViaLangcode(offset, limit, iLangcode, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsViaUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iUid := self.Args("uid").MustInt()
	if (offset > 0) && helper.IsHas(iUid) {
		_FileManaged, _error := model.GetFileManagedsViaUid(offset, limit, iUid, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsViaUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsViaFilenameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iFilename := self.Args("filename").String()
	if (offset > 0) && helper.IsHas(iFilename) {
		_FileManaged, _error := model.GetFileManagedsViaFilename(offset, limit, iFilename, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsViaFilename's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsViaUriHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iUri := self.Args("uri").String()
	if (offset > 0) && helper.IsHas(iUri) {
		_FileManaged, _error := model.GetFileManagedsViaUri(offset, limit, iUri, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsViaUri's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsViaFilemimeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iFilemime := self.Args("filemime").String()
	if (offset > 0) && helper.IsHas(iFilemime) {
		_FileManaged, _error := model.GetFileManagedsViaFilemime(offset, limit, iFilemime, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsViaFilemime's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsViaFilesizeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iFilesize := self.Args("filesize").MustInt64()
	if (offset > 0) && helper.IsHas(iFilesize) {
		_FileManaged, _error := model.GetFileManagedsViaFilesize(offset, limit, iFilesize, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsViaFilesize's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsViaStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iStatus := self.Args("status").MustInt()
	if (offset > 0) && helper.IsHas(iStatus) {
		_FileManaged, _error := model.GetFileManagedsViaStatus(offset, limit, iStatus, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsViaStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsViaCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iCreated := self.Args("created").MustInt()
	if (offset > 0) && helper.IsHas(iCreated) {
		_FileManaged, _error := model.GetFileManagedsViaCreated(offset, limit, iCreated, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsViaCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsViaChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iChanged := self.Args("changed").MustInt()
	if (offset > 0) && helper.IsHas(iChanged) {
		_FileManaged, _error := model.GetFileManagedsViaChanged(offset, limit, iChanged, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsViaChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByFidAndUuidAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFid := self.Args("fid").MustInt64()
	iUuid := self.Args("uuid").String()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iFid) {
		_FileManaged, _error := model.GetFileManagedsByFidAndUuidAndLangcode(offset, limit, iFid,iUuid,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByFidAndUuidAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByFidAndUuidAndUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFid := self.Args("fid").MustInt64()
	iUuid := self.Args("uuid").String()
	iUid := self.Args("uid").MustInt()

	if helper.IsHas(iFid) {
		_FileManaged, _error := model.GetFileManagedsByFidAndUuidAndUid(offset, limit, iFid,iUuid,iUid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByFidAndUuidAndUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByFidAndUuidAndFilenameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFid := self.Args("fid").MustInt64()
	iUuid := self.Args("uuid").String()
	iFilename := self.Args("filename").String()

	if helper.IsHas(iFid) {
		_FileManaged, _error := model.GetFileManagedsByFidAndUuidAndFilename(offset, limit, iFid,iUuid,iFilename)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByFidAndUuidAndFilename's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByFidAndUuidAndUriHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFid := self.Args("fid").MustInt64()
	iUuid := self.Args("uuid").String()
	iUri := self.Args("uri").String()

	if helper.IsHas(iFid) {
		_FileManaged, _error := model.GetFileManagedsByFidAndUuidAndUri(offset, limit, iFid,iUuid,iUri)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByFidAndUuidAndUri's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByFidAndUuidAndFilemimeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFid := self.Args("fid").MustInt64()
	iUuid := self.Args("uuid").String()
	iFilemime := self.Args("filemime").String()

	if helper.IsHas(iFid) {
		_FileManaged, _error := model.GetFileManagedsByFidAndUuidAndFilemime(offset, limit, iFid,iUuid,iFilemime)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByFidAndUuidAndFilemime's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByFidAndUuidAndFilesizeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFid := self.Args("fid").MustInt64()
	iUuid := self.Args("uuid").String()
	iFilesize := self.Args("filesize").MustInt64()

	if helper.IsHas(iFid) {
		_FileManaged, _error := model.GetFileManagedsByFidAndUuidAndFilesize(offset, limit, iFid,iUuid,iFilesize)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByFidAndUuidAndFilesize's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByFidAndUuidAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFid := self.Args("fid").MustInt64()
	iUuid := self.Args("uuid").String()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iFid) {
		_FileManaged, _error := model.GetFileManagedsByFidAndUuidAndStatus(offset, limit, iFid,iUuid,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByFidAndUuidAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByFidAndUuidAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFid := self.Args("fid").MustInt64()
	iUuid := self.Args("uuid").String()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iFid) {
		_FileManaged, _error := model.GetFileManagedsByFidAndUuidAndCreated(offset, limit, iFid,iUuid,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByFidAndUuidAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByFidAndUuidAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFid := self.Args("fid").MustInt64()
	iUuid := self.Args("uuid").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iFid) {
		_FileManaged, _error := model.GetFileManagedsByFidAndUuidAndChanged(offset, limit, iFid,iUuid,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByFidAndUuidAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByFidAndLangcodeAndUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFid := self.Args("fid").MustInt64()
	iLangcode := self.Args("langcode").String()
	iUid := self.Args("uid").MustInt()

	if helper.IsHas(iFid) {
		_FileManaged, _error := model.GetFileManagedsByFidAndLangcodeAndUid(offset, limit, iFid,iLangcode,iUid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByFidAndLangcodeAndUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByFidAndLangcodeAndFilenameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFid := self.Args("fid").MustInt64()
	iLangcode := self.Args("langcode").String()
	iFilename := self.Args("filename").String()

	if helper.IsHas(iFid) {
		_FileManaged, _error := model.GetFileManagedsByFidAndLangcodeAndFilename(offset, limit, iFid,iLangcode,iFilename)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByFidAndLangcodeAndFilename's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByFidAndLangcodeAndUriHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFid := self.Args("fid").MustInt64()
	iLangcode := self.Args("langcode").String()
	iUri := self.Args("uri").String()

	if helper.IsHas(iFid) {
		_FileManaged, _error := model.GetFileManagedsByFidAndLangcodeAndUri(offset, limit, iFid,iLangcode,iUri)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByFidAndLangcodeAndUri's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByFidAndLangcodeAndFilemimeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFid := self.Args("fid").MustInt64()
	iLangcode := self.Args("langcode").String()
	iFilemime := self.Args("filemime").String()

	if helper.IsHas(iFid) {
		_FileManaged, _error := model.GetFileManagedsByFidAndLangcodeAndFilemime(offset, limit, iFid,iLangcode,iFilemime)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByFidAndLangcodeAndFilemime's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByFidAndLangcodeAndFilesizeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFid := self.Args("fid").MustInt64()
	iLangcode := self.Args("langcode").String()
	iFilesize := self.Args("filesize").MustInt64()

	if helper.IsHas(iFid) {
		_FileManaged, _error := model.GetFileManagedsByFidAndLangcodeAndFilesize(offset, limit, iFid,iLangcode,iFilesize)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByFidAndLangcodeAndFilesize's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByFidAndLangcodeAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFid := self.Args("fid").MustInt64()
	iLangcode := self.Args("langcode").String()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iFid) {
		_FileManaged, _error := model.GetFileManagedsByFidAndLangcodeAndStatus(offset, limit, iFid,iLangcode,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByFidAndLangcodeAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByFidAndLangcodeAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFid := self.Args("fid").MustInt64()
	iLangcode := self.Args("langcode").String()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iFid) {
		_FileManaged, _error := model.GetFileManagedsByFidAndLangcodeAndCreated(offset, limit, iFid,iLangcode,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByFidAndLangcodeAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByFidAndLangcodeAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFid := self.Args("fid").MustInt64()
	iLangcode := self.Args("langcode").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iFid) {
		_FileManaged, _error := model.GetFileManagedsByFidAndLangcodeAndChanged(offset, limit, iFid,iLangcode,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByFidAndLangcodeAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByFidAndUidAndFilenameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFid := self.Args("fid").MustInt64()
	iUid := self.Args("uid").MustInt()
	iFilename := self.Args("filename").String()

	if helper.IsHas(iFid) {
		_FileManaged, _error := model.GetFileManagedsByFidAndUidAndFilename(offset, limit, iFid,iUid,iFilename)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByFidAndUidAndFilename's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByFidAndUidAndUriHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFid := self.Args("fid").MustInt64()
	iUid := self.Args("uid").MustInt()
	iUri := self.Args("uri").String()

	if helper.IsHas(iFid) {
		_FileManaged, _error := model.GetFileManagedsByFidAndUidAndUri(offset, limit, iFid,iUid,iUri)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByFidAndUidAndUri's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByFidAndUidAndFilemimeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFid := self.Args("fid").MustInt64()
	iUid := self.Args("uid").MustInt()
	iFilemime := self.Args("filemime").String()

	if helper.IsHas(iFid) {
		_FileManaged, _error := model.GetFileManagedsByFidAndUidAndFilemime(offset, limit, iFid,iUid,iFilemime)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByFidAndUidAndFilemime's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByFidAndUidAndFilesizeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFid := self.Args("fid").MustInt64()
	iUid := self.Args("uid").MustInt()
	iFilesize := self.Args("filesize").MustInt64()

	if helper.IsHas(iFid) {
		_FileManaged, _error := model.GetFileManagedsByFidAndUidAndFilesize(offset, limit, iFid,iUid,iFilesize)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByFidAndUidAndFilesize's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByFidAndUidAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFid := self.Args("fid").MustInt64()
	iUid := self.Args("uid").MustInt()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iFid) {
		_FileManaged, _error := model.GetFileManagedsByFidAndUidAndStatus(offset, limit, iFid,iUid,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByFidAndUidAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByFidAndUidAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFid := self.Args("fid").MustInt64()
	iUid := self.Args("uid").MustInt()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iFid) {
		_FileManaged, _error := model.GetFileManagedsByFidAndUidAndCreated(offset, limit, iFid,iUid,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByFidAndUidAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByFidAndUidAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFid := self.Args("fid").MustInt64()
	iUid := self.Args("uid").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iFid) {
		_FileManaged, _error := model.GetFileManagedsByFidAndUidAndChanged(offset, limit, iFid,iUid,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByFidAndUidAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByFidAndFilenameAndUriHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFid := self.Args("fid").MustInt64()
	iFilename := self.Args("filename").String()
	iUri := self.Args("uri").String()

	if helper.IsHas(iFid) {
		_FileManaged, _error := model.GetFileManagedsByFidAndFilenameAndUri(offset, limit, iFid,iFilename,iUri)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByFidAndFilenameAndUri's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByFidAndFilenameAndFilemimeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFid := self.Args("fid").MustInt64()
	iFilename := self.Args("filename").String()
	iFilemime := self.Args("filemime").String()

	if helper.IsHas(iFid) {
		_FileManaged, _error := model.GetFileManagedsByFidAndFilenameAndFilemime(offset, limit, iFid,iFilename,iFilemime)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByFidAndFilenameAndFilemime's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByFidAndFilenameAndFilesizeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFid := self.Args("fid").MustInt64()
	iFilename := self.Args("filename").String()
	iFilesize := self.Args("filesize").MustInt64()

	if helper.IsHas(iFid) {
		_FileManaged, _error := model.GetFileManagedsByFidAndFilenameAndFilesize(offset, limit, iFid,iFilename,iFilesize)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByFidAndFilenameAndFilesize's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByFidAndFilenameAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFid := self.Args("fid").MustInt64()
	iFilename := self.Args("filename").String()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iFid) {
		_FileManaged, _error := model.GetFileManagedsByFidAndFilenameAndStatus(offset, limit, iFid,iFilename,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByFidAndFilenameAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByFidAndFilenameAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFid := self.Args("fid").MustInt64()
	iFilename := self.Args("filename").String()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iFid) {
		_FileManaged, _error := model.GetFileManagedsByFidAndFilenameAndCreated(offset, limit, iFid,iFilename,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByFidAndFilenameAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByFidAndFilenameAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFid := self.Args("fid").MustInt64()
	iFilename := self.Args("filename").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iFid) {
		_FileManaged, _error := model.GetFileManagedsByFidAndFilenameAndChanged(offset, limit, iFid,iFilename,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByFidAndFilenameAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByFidAndUriAndFilemimeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFid := self.Args("fid").MustInt64()
	iUri := self.Args("uri").String()
	iFilemime := self.Args("filemime").String()

	if helper.IsHas(iFid) {
		_FileManaged, _error := model.GetFileManagedsByFidAndUriAndFilemime(offset, limit, iFid,iUri,iFilemime)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByFidAndUriAndFilemime's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByFidAndUriAndFilesizeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFid := self.Args("fid").MustInt64()
	iUri := self.Args("uri").String()
	iFilesize := self.Args("filesize").MustInt64()

	if helper.IsHas(iFid) {
		_FileManaged, _error := model.GetFileManagedsByFidAndUriAndFilesize(offset, limit, iFid,iUri,iFilesize)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByFidAndUriAndFilesize's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByFidAndUriAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFid := self.Args("fid").MustInt64()
	iUri := self.Args("uri").String()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iFid) {
		_FileManaged, _error := model.GetFileManagedsByFidAndUriAndStatus(offset, limit, iFid,iUri,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByFidAndUriAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByFidAndUriAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFid := self.Args("fid").MustInt64()
	iUri := self.Args("uri").String()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iFid) {
		_FileManaged, _error := model.GetFileManagedsByFidAndUriAndCreated(offset, limit, iFid,iUri,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByFidAndUriAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByFidAndUriAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFid := self.Args("fid").MustInt64()
	iUri := self.Args("uri").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iFid) {
		_FileManaged, _error := model.GetFileManagedsByFidAndUriAndChanged(offset, limit, iFid,iUri,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByFidAndUriAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByFidAndFilemimeAndFilesizeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFid := self.Args("fid").MustInt64()
	iFilemime := self.Args("filemime").String()
	iFilesize := self.Args("filesize").MustInt64()

	if helper.IsHas(iFid) {
		_FileManaged, _error := model.GetFileManagedsByFidAndFilemimeAndFilesize(offset, limit, iFid,iFilemime,iFilesize)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByFidAndFilemimeAndFilesize's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByFidAndFilemimeAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFid := self.Args("fid").MustInt64()
	iFilemime := self.Args("filemime").String()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iFid) {
		_FileManaged, _error := model.GetFileManagedsByFidAndFilemimeAndStatus(offset, limit, iFid,iFilemime,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByFidAndFilemimeAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByFidAndFilemimeAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFid := self.Args("fid").MustInt64()
	iFilemime := self.Args("filemime").String()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iFid) {
		_FileManaged, _error := model.GetFileManagedsByFidAndFilemimeAndCreated(offset, limit, iFid,iFilemime,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByFidAndFilemimeAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByFidAndFilemimeAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFid := self.Args("fid").MustInt64()
	iFilemime := self.Args("filemime").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iFid) {
		_FileManaged, _error := model.GetFileManagedsByFidAndFilemimeAndChanged(offset, limit, iFid,iFilemime,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByFidAndFilemimeAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByFidAndFilesizeAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFid := self.Args("fid").MustInt64()
	iFilesize := self.Args("filesize").MustInt64()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iFid) {
		_FileManaged, _error := model.GetFileManagedsByFidAndFilesizeAndStatus(offset, limit, iFid,iFilesize,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByFidAndFilesizeAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByFidAndFilesizeAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFid := self.Args("fid").MustInt64()
	iFilesize := self.Args("filesize").MustInt64()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iFid) {
		_FileManaged, _error := model.GetFileManagedsByFidAndFilesizeAndCreated(offset, limit, iFid,iFilesize,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByFidAndFilesizeAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByFidAndFilesizeAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFid := self.Args("fid").MustInt64()
	iFilesize := self.Args("filesize").MustInt64()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iFid) {
		_FileManaged, _error := model.GetFileManagedsByFidAndFilesizeAndChanged(offset, limit, iFid,iFilesize,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByFidAndFilesizeAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByFidAndStatusAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFid := self.Args("fid").MustInt64()
	iStatus := self.Args("status").MustInt()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iFid) {
		_FileManaged, _error := model.GetFileManagedsByFidAndStatusAndCreated(offset, limit, iFid,iStatus,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByFidAndStatusAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByFidAndStatusAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFid := self.Args("fid").MustInt64()
	iStatus := self.Args("status").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iFid) {
		_FileManaged, _error := model.GetFileManagedsByFidAndStatusAndChanged(offset, limit, iFid,iStatus,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByFidAndStatusAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByFidAndCreatedAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFid := self.Args("fid").MustInt64()
	iCreated := self.Args("created").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iFid) {
		_FileManaged, _error := model.GetFileManagedsByFidAndCreatedAndChanged(offset, limit, iFid,iCreated,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByFidAndCreatedAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByUuidAndLangcodeAndUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUuid := self.Args("uuid").String()
	iLangcode := self.Args("langcode").String()
	iUid := self.Args("uid").MustInt()

	if helper.IsHas(iUuid) {
		_FileManaged, _error := model.GetFileManagedsByUuidAndLangcodeAndUid(offset, limit, iUuid,iLangcode,iUid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByUuidAndLangcodeAndUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByUuidAndLangcodeAndFilenameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUuid := self.Args("uuid").String()
	iLangcode := self.Args("langcode").String()
	iFilename := self.Args("filename").String()

	if helper.IsHas(iUuid) {
		_FileManaged, _error := model.GetFileManagedsByUuidAndLangcodeAndFilename(offset, limit, iUuid,iLangcode,iFilename)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByUuidAndLangcodeAndFilename's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByUuidAndLangcodeAndUriHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUuid := self.Args("uuid").String()
	iLangcode := self.Args("langcode").String()
	iUri := self.Args("uri").String()

	if helper.IsHas(iUuid) {
		_FileManaged, _error := model.GetFileManagedsByUuidAndLangcodeAndUri(offset, limit, iUuid,iLangcode,iUri)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByUuidAndLangcodeAndUri's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByUuidAndLangcodeAndFilemimeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUuid := self.Args("uuid").String()
	iLangcode := self.Args("langcode").String()
	iFilemime := self.Args("filemime").String()

	if helper.IsHas(iUuid) {
		_FileManaged, _error := model.GetFileManagedsByUuidAndLangcodeAndFilemime(offset, limit, iUuid,iLangcode,iFilemime)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByUuidAndLangcodeAndFilemime's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByUuidAndLangcodeAndFilesizeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUuid := self.Args("uuid").String()
	iLangcode := self.Args("langcode").String()
	iFilesize := self.Args("filesize").MustInt64()

	if helper.IsHas(iUuid) {
		_FileManaged, _error := model.GetFileManagedsByUuidAndLangcodeAndFilesize(offset, limit, iUuid,iLangcode,iFilesize)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByUuidAndLangcodeAndFilesize's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByUuidAndLangcodeAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUuid := self.Args("uuid").String()
	iLangcode := self.Args("langcode").String()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iUuid) {
		_FileManaged, _error := model.GetFileManagedsByUuidAndLangcodeAndStatus(offset, limit, iUuid,iLangcode,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByUuidAndLangcodeAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByUuidAndLangcodeAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUuid := self.Args("uuid").String()
	iLangcode := self.Args("langcode").String()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iUuid) {
		_FileManaged, _error := model.GetFileManagedsByUuidAndLangcodeAndCreated(offset, limit, iUuid,iLangcode,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByUuidAndLangcodeAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByUuidAndLangcodeAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUuid := self.Args("uuid").String()
	iLangcode := self.Args("langcode").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iUuid) {
		_FileManaged, _error := model.GetFileManagedsByUuidAndLangcodeAndChanged(offset, limit, iUuid,iLangcode,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByUuidAndLangcodeAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByUuidAndUidAndFilenameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUuid := self.Args("uuid").String()
	iUid := self.Args("uid").MustInt()
	iFilename := self.Args("filename").String()

	if helper.IsHas(iUuid) {
		_FileManaged, _error := model.GetFileManagedsByUuidAndUidAndFilename(offset, limit, iUuid,iUid,iFilename)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByUuidAndUidAndFilename's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByUuidAndUidAndUriHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUuid := self.Args("uuid").String()
	iUid := self.Args("uid").MustInt()
	iUri := self.Args("uri").String()

	if helper.IsHas(iUuid) {
		_FileManaged, _error := model.GetFileManagedsByUuidAndUidAndUri(offset, limit, iUuid,iUid,iUri)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByUuidAndUidAndUri's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByUuidAndUidAndFilemimeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUuid := self.Args("uuid").String()
	iUid := self.Args("uid").MustInt()
	iFilemime := self.Args("filemime").String()

	if helper.IsHas(iUuid) {
		_FileManaged, _error := model.GetFileManagedsByUuidAndUidAndFilemime(offset, limit, iUuid,iUid,iFilemime)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByUuidAndUidAndFilemime's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByUuidAndUidAndFilesizeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUuid := self.Args("uuid").String()
	iUid := self.Args("uid").MustInt()
	iFilesize := self.Args("filesize").MustInt64()

	if helper.IsHas(iUuid) {
		_FileManaged, _error := model.GetFileManagedsByUuidAndUidAndFilesize(offset, limit, iUuid,iUid,iFilesize)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByUuidAndUidAndFilesize's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByUuidAndUidAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUuid := self.Args("uuid").String()
	iUid := self.Args("uid").MustInt()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iUuid) {
		_FileManaged, _error := model.GetFileManagedsByUuidAndUidAndStatus(offset, limit, iUuid,iUid,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByUuidAndUidAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByUuidAndUidAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUuid := self.Args("uuid").String()
	iUid := self.Args("uid").MustInt()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iUuid) {
		_FileManaged, _error := model.GetFileManagedsByUuidAndUidAndCreated(offset, limit, iUuid,iUid,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByUuidAndUidAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByUuidAndUidAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUuid := self.Args("uuid").String()
	iUid := self.Args("uid").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iUuid) {
		_FileManaged, _error := model.GetFileManagedsByUuidAndUidAndChanged(offset, limit, iUuid,iUid,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByUuidAndUidAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByUuidAndFilenameAndUriHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUuid := self.Args("uuid").String()
	iFilename := self.Args("filename").String()
	iUri := self.Args("uri").String()

	if helper.IsHas(iUuid) {
		_FileManaged, _error := model.GetFileManagedsByUuidAndFilenameAndUri(offset, limit, iUuid,iFilename,iUri)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByUuidAndFilenameAndUri's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByUuidAndFilenameAndFilemimeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUuid := self.Args("uuid").String()
	iFilename := self.Args("filename").String()
	iFilemime := self.Args("filemime").String()

	if helper.IsHas(iUuid) {
		_FileManaged, _error := model.GetFileManagedsByUuidAndFilenameAndFilemime(offset, limit, iUuid,iFilename,iFilemime)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByUuidAndFilenameAndFilemime's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByUuidAndFilenameAndFilesizeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUuid := self.Args("uuid").String()
	iFilename := self.Args("filename").String()
	iFilesize := self.Args("filesize").MustInt64()

	if helper.IsHas(iUuid) {
		_FileManaged, _error := model.GetFileManagedsByUuidAndFilenameAndFilesize(offset, limit, iUuid,iFilename,iFilesize)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByUuidAndFilenameAndFilesize's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByUuidAndFilenameAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUuid := self.Args("uuid").String()
	iFilename := self.Args("filename").String()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iUuid) {
		_FileManaged, _error := model.GetFileManagedsByUuidAndFilenameAndStatus(offset, limit, iUuid,iFilename,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByUuidAndFilenameAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByUuidAndFilenameAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUuid := self.Args("uuid").String()
	iFilename := self.Args("filename").String()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iUuid) {
		_FileManaged, _error := model.GetFileManagedsByUuidAndFilenameAndCreated(offset, limit, iUuid,iFilename,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByUuidAndFilenameAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByUuidAndFilenameAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUuid := self.Args("uuid").String()
	iFilename := self.Args("filename").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iUuid) {
		_FileManaged, _error := model.GetFileManagedsByUuidAndFilenameAndChanged(offset, limit, iUuid,iFilename,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByUuidAndFilenameAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByUuidAndUriAndFilemimeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUuid := self.Args("uuid").String()
	iUri := self.Args("uri").String()
	iFilemime := self.Args("filemime").String()

	if helper.IsHas(iUuid) {
		_FileManaged, _error := model.GetFileManagedsByUuidAndUriAndFilemime(offset, limit, iUuid,iUri,iFilemime)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByUuidAndUriAndFilemime's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByUuidAndUriAndFilesizeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUuid := self.Args("uuid").String()
	iUri := self.Args("uri").String()
	iFilesize := self.Args("filesize").MustInt64()

	if helper.IsHas(iUuid) {
		_FileManaged, _error := model.GetFileManagedsByUuidAndUriAndFilesize(offset, limit, iUuid,iUri,iFilesize)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByUuidAndUriAndFilesize's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByUuidAndUriAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUuid := self.Args("uuid").String()
	iUri := self.Args("uri").String()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iUuid) {
		_FileManaged, _error := model.GetFileManagedsByUuidAndUriAndStatus(offset, limit, iUuid,iUri,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByUuidAndUriAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByUuidAndUriAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUuid := self.Args("uuid").String()
	iUri := self.Args("uri").String()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iUuid) {
		_FileManaged, _error := model.GetFileManagedsByUuidAndUriAndCreated(offset, limit, iUuid,iUri,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByUuidAndUriAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByUuidAndUriAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUuid := self.Args("uuid").String()
	iUri := self.Args("uri").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iUuid) {
		_FileManaged, _error := model.GetFileManagedsByUuidAndUriAndChanged(offset, limit, iUuid,iUri,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByUuidAndUriAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByUuidAndFilemimeAndFilesizeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUuid := self.Args("uuid").String()
	iFilemime := self.Args("filemime").String()
	iFilesize := self.Args("filesize").MustInt64()

	if helper.IsHas(iUuid) {
		_FileManaged, _error := model.GetFileManagedsByUuidAndFilemimeAndFilesize(offset, limit, iUuid,iFilemime,iFilesize)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByUuidAndFilemimeAndFilesize's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByUuidAndFilemimeAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUuid := self.Args("uuid").String()
	iFilemime := self.Args("filemime").String()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iUuid) {
		_FileManaged, _error := model.GetFileManagedsByUuidAndFilemimeAndStatus(offset, limit, iUuid,iFilemime,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByUuidAndFilemimeAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByUuidAndFilemimeAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUuid := self.Args("uuid").String()
	iFilemime := self.Args("filemime").String()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iUuid) {
		_FileManaged, _error := model.GetFileManagedsByUuidAndFilemimeAndCreated(offset, limit, iUuid,iFilemime,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByUuidAndFilemimeAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByUuidAndFilemimeAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUuid := self.Args("uuid").String()
	iFilemime := self.Args("filemime").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iUuid) {
		_FileManaged, _error := model.GetFileManagedsByUuidAndFilemimeAndChanged(offset, limit, iUuid,iFilemime,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByUuidAndFilemimeAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByUuidAndFilesizeAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUuid := self.Args("uuid").String()
	iFilesize := self.Args("filesize").MustInt64()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iUuid) {
		_FileManaged, _error := model.GetFileManagedsByUuidAndFilesizeAndStatus(offset, limit, iUuid,iFilesize,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByUuidAndFilesizeAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByUuidAndFilesizeAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUuid := self.Args("uuid").String()
	iFilesize := self.Args("filesize").MustInt64()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iUuid) {
		_FileManaged, _error := model.GetFileManagedsByUuidAndFilesizeAndCreated(offset, limit, iUuid,iFilesize,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByUuidAndFilesizeAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByUuidAndFilesizeAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUuid := self.Args("uuid").String()
	iFilesize := self.Args("filesize").MustInt64()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iUuid) {
		_FileManaged, _error := model.GetFileManagedsByUuidAndFilesizeAndChanged(offset, limit, iUuid,iFilesize,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByUuidAndFilesizeAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByUuidAndStatusAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUuid := self.Args("uuid").String()
	iStatus := self.Args("status").MustInt()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iUuid) {
		_FileManaged, _error := model.GetFileManagedsByUuidAndStatusAndCreated(offset, limit, iUuid,iStatus,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByUuidAndStatusAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByUuidAndStatusAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUuid := self.Args("uuid").String()
	iStatus := self.Args("status").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iUuid) {
		_FileManaged, _error := model.GetFileManagedsByUuidAndStatusAndChanged(offset, limit, iUuid,iStatus,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByUuidAndStatusAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByUuidAndCreatedAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUuid := self.Args("uuid").String()
	iCreated := self.Args("created").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iUuid) {
		_FileManaged, _error := model.GetFileManagedsByUuidAndCreatedAndChanged(offset, limit, iUuid,iCreated,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByUuidAndCreatedAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByLangcodeAndUidAndFilenameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iUid := self.Args("uid").MustInt()
	iFilename := self.Args("filename").String()

	if helper.IsHas(iLangcode) {
		_FileManaged, _error := model.GetFileManagedsByLangcodeAndUidAndFilename(offset, limit, iLangcode,iUid,iFilename)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByLangcodeAndUidAndFilename's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByLangcodeAndUidAndUriHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iUid := self.Args("uid").MustInt()
	iUri := self.Args("uri").String()

	if helper.IsHas(iLangcode) {
		_FileManaged, _error := model.GetFileManagedsByLangcodeAndUidAndUri(offset, limit, iLangcode,iUid,iUri)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByLangcodeAndUidAndUri's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByLangcodeAndUidAndFilemimeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iUid := self.Args("uid").MustInt()
	iFilemime := self.Args("filemime").String()

	if helper.IsHas(iLangcode) {
		_FileManaged, _error := model.GetFileManagedsByLangcodeAndUidAndFilemime(offset, limit, iLangcode,iUid,iFilemime)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByLangcodeAndUidAndFilemime's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByLangcodeAndUidAndFilesizeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iUid := self.Args("uid").MustInt()
	iFilesize := self.Args("filesize").MustInt64()

	if helper.IsHas(iLangcode) {
		_FileManaged, _error := model.GetFileManagedsByLangcodeAndUidAndFilesize(offset, limit, iLangcode,iUid,iFilesize)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByLangcodeAndUidAndFilesize's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByLangcodeAndUidAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iUid := self.Args("uid").MustInt()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iLangcode) {
		_FileManaged, _error := model.GetFileManagedsByLangcodeAndUidAndStatus(offset, limit, iLangcode,iUid,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByLangcodeAndUidAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByLangcodeAndUidAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iUid := self.Args("uid").MustInt()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iLangcode) {
		_FileManaged, _error := model.GetFileManagedsByLangcodeAndUidAndCreated(offset, limit, iLangcode,iUid,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByLangcodeAndUidAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByLangcodeAndUidAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iUid := self.Args("uid").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iLangcode) {
		_FileManaged, _error := model.GetFileManagedsByLangcodeAndUidAndChanged(offset, limit, iLangcode,iUid,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByLangcodeAndUidAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByLangcodeAndFilenameAndUriHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iFilename := self.Args("filename").String()
	iUri := self.Args("uri").String()

	if helper.IsHas(iLangcode) {
		_FileManaged, _error := model.GetFileManagedsByLangcodeAndFilenameAndUri(offset, limit, iLangcode,iFilename,iUri)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByLangcodeAndFilenameAndUri's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByLangcodeAndFilenameAndFilemimeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iFilename := self.Args("filename").String()
	iFilemime := self.Args("filemime").String()

	if helper.IsHas(iLangcode) {
		_FileManaged, _error := model.GetFileManagedsByLangcodeAndFilenameAndFilemime(offset, limit, iLangcode,iFilename,iFilemime)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByLangcodeAndFilenameAndFilemime's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByLangcodeAndFilenameAndFilesizeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iFilename := self.Args("filename").String()
	iFilesize := self.Args("filesize").MustInt64()

	if helper.IsHas(iLangcode) {
		_FileManaged, _error := model.GetFileManagedsByLangcodeAndFilenameAndFilesize(offset, limit, iLangcode,iFilename,iFilesize)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByLangcodeAndFilenameAndFilesize's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByLangcodeAndFilenameAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iFilename := self.Args("filename").String()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iLangcode) {
		_FileManaged, _error := model.GetFileManagedsByLangcodeAndFilenameAndStatus(offset, limit, iLangcode,iFilename,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByLangcodeAndFilenameAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByLangcodeAndFilenameAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iFilename := self.Args("filename").String()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iLangcode) {
		_FileManaged, _error := model.GetFileManagedsByLangcodeAndFilenameAndCreated(offset, limit, iLangcode,iFilename,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByLangcodeAndFilenameAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByLangcodeAndFilenameAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iFilename := self.Args("filename").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iLangcode) {
		_FileManaged, _error := model.GetFileManagedsByLangcodeAndFilenameAndChanged(offset, limit, iLangcode,iFilename,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByLangcodeAndFilenameAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByLangcodeAndUriAndFilemimeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iUri := self.Args("uri").String()
	iFilemime := self.Args("filemime").String()

	if helper.IsHas(iLangcode) {
		_FileManaged, _error := model.GetFileManagedsByLangcodeAndUriAndFilemime(offset, limit, iLangcode,iUri,iFilemime)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByLangcodeAndUriAndFilemime's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByLangcodeAndUriAndFilesizeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iUri := self.Args("uri").String()
	iFilesize := self.Args("filesize").MustInt64()

	if helper.IsHas(iLangcode) {
		_FileManaged, _error := model.GetFileManagedsByLangcodeAndUriAndFilesize(offset, limit, iLangcode,iUri,iFilesize)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByLangcodeAndUriAndFilesize's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByLangcodeAndUriAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iUri := self.Args("uri").String()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iLangcode) {
		_FileManaged, _error := model.GetFileManagedsByLangcodeAndUriAndStatus(offset, limit, iLangcode,iUri,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByLangcodeAndUriAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByLangcodeAndUriAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iUri := self.Args("uri").String()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iLangcode) {
		_FileManaged, _error := model.GetFileManagedsByLangcodeAndUriAndCreated(offset, limit, iLangcode,iUri,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByLangcodeAndUriAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByLangcodeAndUriAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iUri := self.Args("uri").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iLangcode) {
		_FileManaged, _error := model.GetFileManagedsByLangcodeAndUriAndChanged(offset, limit, iLangcode,iUri,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByLangcodeAndUriAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByLangcodeAndFilemimeAndFilesizeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iFilemime := self.Args("filemime").String()
	iFilesize := self.Args("filesize").MustInt64()

	if helper.IsHas(iLangcode) {
		_FileManaged, _error := model.GetFileManagedsByLangcodeAndFilemimeAndFilesize(offset, limit, iLangcode,iFilemime,iFilesize)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByLangcodeAndFilemimeAndFilesize's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByLangcodeAndFilemimeAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iFilemime := self.Args("filemime").String()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iLangcode) {
		_FileManaged, _error := model.GetFileManagedsByLangcodeAndFilemimeAndStatus(offset, limit, iLangcode,iFilemime,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByLangcodeAndFilemimeAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByLangcodeAndFilemimeAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iFilemime := self.Args("filemime").String()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iLangcode) {
		_FileManaged, _error := model.GetFileManagedsByLangcodeAndFilemimeAndCreated(offset, limit, iLangcode,iFilemime,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByLangcodeAndFilemimeAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByLangcodeAndFilemimeAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iFilemime := self.Args("filemime").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iLangcode) {
		_FileManaged, _error := model.GetFileManagedsByLangcodeAndFilemimeAndChanged(offset, limit, iLangcode,iFilemime,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByLangcodeAndFilemimeAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByLangcodeAndFilesizeAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iFilesize := self.Args("filesize").MustInt64()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iLangcode) {
		_FileManaged, _error := model.GetFileManagedsByLangcodeAndFilesizeAndStatus(offset, limit, iLangcode,iFilesize,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByLangcodeAndFilesizeAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByLangcodeAndFilesizeAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iFilesize := self.Args("filesize").MustInt64()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iLangcode) {
		_FileManaged, _error := model.GetFileManagedsByLangcodeAndFilesizeAndCreated(offset, limit, iLangcode,iFilesize,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByLangcodeAndFilesizeAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByLangcodeAndFilesizeAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iFilesize := self.Args("filesize").MustInt64()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iLangcode) {
		_FileManaged, _error := model.GetFileManagedsByLangcodeAndFilesizeAndChanged(offset, limit, iLangcode,iFilesize,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByLangcodeAndFilesizeAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByLangcodeAndStatusAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iStatus := self.Args("status").MustInt()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iLangcode) {
		_FileManaged, _error := model.GetFileManagedsByLangcodeAndStatusAndCreated(offset, limit, iLangcode,iStatus,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByLangcodeAndStatusAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByLangcodeAndStatusAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iStatus := self.Args("status").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iLangcode) {
		_FileManaged, _error := model.GetFileManagedsByLangcodeAndStatusAndChanged(offset, limit, iLangcode,iStatus,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByLangcodeAndStatusAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByLangcodeAndCreatedAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iCreated := self.Args("created").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iLangcode) {
		_FileManaged, _error := model.GetFileManagedsByLangcodeAndCreatedAndChanged(offset, limit, iLangcode,iCreated,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByLangcodeAndCreatedAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByUidAndFilenameAndUriHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iFilename := self.Args("filename").String()
	iUri := self.Args("uri").String()

	if helper.IsHas(iUid) {
		_FileManaged, _error := model.GetFileManagedsByUidAndFilenameAndUri(offset, limit, iUid,iFilename,iUri)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByUidAndFilenameAndUri's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByUidAndFilenameAndFilemimeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iFilename := self.Args("filename").String()
	iFilemime := self.Args("filemime").String()

	if helper.IsHas(iUid) {
		_FileManaged, _error := model.GetFileManagedsByUidAndFilenameAndFilemime(offset, limit, iUid,iFilename,iFilemime)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByUidAndFilenameAndFilemime's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByUidAndFilenameAndFilesizeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iFilename := self.Args("filename").String()
	iFilesize := self.Args("filesize").MustInt64()

	if helper.IsHas(iUid) {
		_FileManaged, _error := model.GetFileManagedsByUidAndFilenameAndFilesize(offset, limit, iUid,iFilename,iFilesize)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByUidAndFilenameAndFilesize's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByUidAndFilenameAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iFilename := self.Args("filename").String()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iUid) {
		_FileManaged, _error := model.GetFileManagedsByUidAndFilenameAndStatus(offset, limit, iUid,iFilename,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByUidAndFilenameAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByUidAndFilenameAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iFilename := self.Args("filename").String()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iUid) {
		_FileManaged, _error := model.GetFileManagedsByUidAndFilenameAndCreated(offset, limit, iUid,iFilename,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByUidAndFilenameAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByUidAndFilenameAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iFilename := self.Args("filename").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iUid) {
		_FileManaged, _error := model.GetFileManagedsByUidAndFilenameAndChanged(offset, limit, iUid,iFilename,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByUidAndFilenameAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByUidAndUriAndFilemimeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iUri := self.Args("uri").String()
	iFilemime := self.Args("filemime").String()

	if helper.IsHas(iUid) {
		_FileManaged, _error := model.GetFileManagedsByUidAndUriAndFilemime(offset, limit, iUid,iUri,iFilemime)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByUidAndUriAndFilemime's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByUidAndUriAndFilesizeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iUri := self.Args("uri").String()
	iFilesize := self.Args("filesize").MustInt64()

	if helper.IsHas(iUid) {
		_FileManaged, _error := model.GetFileManagedsByUidAndUriAndFilesize(offset, limit, iUid,iUri,iFilesize)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByUidAndUriAndFilesize's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByUidAndUriAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iUri := self.Args("uri").String()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iUid) {
		_FileManaged, _error := model.GetFileManagedsByUidAndUriAndStatus(offset, limit, iUid,iUri,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByUidAndUriAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByUidAndUriAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iUri := self.Args("uri").String()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iUid) {
		_FileManaged, _error := model.GetFileManagedsByUidAndUriAndCreated(offset, limit, iUid,iUri,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByUidAndUriAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByUidAndUriAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iUri := self.Args("uri").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iUid) {
		_FileManaged, _error := model.GetFileManagedsByUidAndUriAndChanged(offset, limit, iUid,iUri,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByUidAndUriAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByUidAndFilemimeAndFilesizeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iFilemime := self.Args("filemime").String()
	iFilesize := self.Args("filesize").MustInt64()

	if helper.IsHas(iUid) {
		_FileManaged, _error := model.GetFileManagedsByUidAndFilemimeAndFilesize(offset, limit, iUid,iFilemime,iFilesize)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByUidAndFilemimeAndFilesize's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByUidAndFilemimeAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iFilemime := self.Args("filemime").String()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iUid) {
		_FileManaged, _error := model.GetFileManagedsByUidAndFilemimeAndStatus(offset, limit, iUid,iFilemime,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByUidAndFilemimeAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByUidAndFilemimeAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iFilemime := self.Args("filemime").String()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iUid) {
		_FileManaged, _error := model.GetFileManagedsByUidAndFilemimeAndCreated(offset, limit, iUid,iFilemime,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByUidAndFilemimeAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByUidAndFilemimeAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iFilemime := self.Args("filemime").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iUid) {
		_FileManaged, _error := model.GetFileManagedsByUidAndFilemimeAndChanged(offset, limit, iUid,iFilemime,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByUidAndFilemimeAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByUidAndFilesizeAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iFilesize := self.Args("filesize").MustInt64()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iUid) {
		_FileManaged, _error := model.GetFileManagedsByUidAndFilesizeAndStatus(offset, limit, iUid,iFilesize,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByUidAndFilesizeAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByUidAndFilesizeAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iFilesize := self.Args("filesize").MustInt64()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iUid) {
		_FileManaged, _error := model.GetFileManagedsByUidAndFilesizeAndCreated(offset, limit, iUid,iFilesize,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByUidAndFilesizeAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByUidAndFilesizeAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iFilesize := self.Args("filesize").MustInt64()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iUid) {
		_FileManaged, _error := model.GetFileManagedsByUidAndFilesizeAndChanged(offset, limit, iUid,iFilesize,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByUidAndFilesizeAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByUidAndStatusAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iStatus := self.Args("status").MustInt()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iUid) {
		_FileManaged, _error := model.GetFileManagedsByUidAndStatusAndCreated(offset, limit, iUid,iStatus,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByUidAndStatusAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByUidAndStatusAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iStatus := self.Args("status").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iUid) {
		_FileManaged, _error := model.GetFileManagedsByUidAndStatusAndChanged(offset, limit, iUid,iStatus,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByUidAndStatusAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByUidAndCreatedAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iCreated := self.Args("created").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iUid) {
		_FileManaged, _error := model.GetFileManagedsByUidAndCreatedAndChanged(offset, limit, iUid,iCreated,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByUidAndCreatedAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByFilenameAndUriAndFilemimeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFilename := self.Args("filename").String()
	iUri := self.Args("uri").String()
	iFilemime := self.Args("filemime").String()

	if helper.IsHas(iFilename) {
		_FileManaged, _error := model.GetFileManagedsByFilenameAndUriAndFilemime(offset, limit, iFilename,iUri,iFilemime)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByFilenameAndUriAndFilemime's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByFilenameAndUriAndFilesizeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFilename := self.Args("filename").String()
	iUri := self.Args("uri").String()
	iFilesize := self.Args("filesize").MustInt64()

	if helper.IsHas(iFilename) {
		_FileManaged, _error := model.GetFileManagedsByFilenameAndUriAndFilesize(offset, limit, iFilename,iUri,iFilesize)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByFilenameAndUriAndFilesize's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByFilenameAndUriAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFilename := self.Args("filename").String()
	iUri := self.Args("uri").String()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iFilename) {
		_FileManaged, _error := model.GetFileManagedsByFilenameAndUriAndStatus(offset, limit, iFilename,iUri,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByFilenameAndUriAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByFilenameAndUriAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFilename := self.Args("filename").String()
	iUri := self.Args("uri").String()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iFilename) {
		_FileManaged, _error := model.GetFileManagedsByFilenameAndUriAndCreated(offset, limit, iFilename,iUri,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByFilenameAndUriAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByFilenameAndUriAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFilename := self.Args("filename").String()
	iUri := self.Args("uri").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iFilename) {
		_FileManaged, _error := model.GetFileManagedsByFilenameAndUriAndChanged(offset, limit, iFilename,iUri,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByFilenameAndUriAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByFilenameAndFilemimeAndFilesizeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFilename := self.Args("filename").String()
	iFilemime := self.Args("filemime").String()
	iFilesize := self.Args("filesize").MustInt64()

	if helper.IsHas(iFilename) {
		_FileManaged, _error := model.GetFileManagedsByFilenameAndFilemimeAndFilesize(offset, limit, iFilename,iFilemime,iFilesize)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByFilenameAndFilemimeAndFilesize's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByFilenameAndFilemimeAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFilename := self.Args("filename").String()
	iFilemime := self.Args("filemime").String()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iFilename) {
		_FileManaged, _error := model.GetFileManagedsByFilenameAndFilemimeAndStatus(offset, limit, iFilename,iFilemime,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByFilenameAndFilemimeAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByFilenameAndFilemimeAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFilename := self.Args("filename").String()
	iFilemime := self.Args("filemime").String()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iFilename) {
		_FileManaged, _error := model.GetFileManagedsByFilenameAndFilemimeAndCreated(offset, limit, iFilename,iFilemime,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByFilenameAndFilemimeAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByFilenameAndFilemimeAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFilename := self.Args("filename").String()
	iFilemime := self.Args("filemime").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iFilename) {
		_FileManaged, _error := model.GetFileManagedsByFilenameAndFilemimeAndChanged(offset, limit, iFilename,iFilemime,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByFilenameAndFilemimeAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByFilenameAndFilesizeAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFilename := self.Args("filename").String()
	iFilesize := self.Args("filesize").MustInt64()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iFilename) {
		_FileManaged, _error := model.GetFileManagedsByFilenameAndFilesizeAndStatus(offset, limit, iFilename,iFilesize,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByFilenameAndFilesizeAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByFilenameAndFilesizeAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFilename := self.Args("filename").String()
	iFilesize := self.Args("filesize").MustInt64()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iFilename) {
		_FileManaged, _error := model.GetFileManagedsByFilenameAndFilesizeAndCreated(offset, limit, iFilename,iFilesize,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByFilenameAndFilesizeAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByFilenameAndFilesizeAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFilename := self.Args("filename").String()
	iFilesize := self.Args("filesize").MustInt64()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iFilename) {
		_FileManaged, _error := model.GetFileManagedsByFilenameAndFilesizeAndChanged(offset, limit, iFilename,iFilesize,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByFilenameAndFilesizeAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByFilenameAndStatusAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFilename := self.Args("filename").String()
	iStatus := self.Args("status").MustInt()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iFilename) {
		_FileManaged, _error := model.GetFileManagedsByFilenameAndStatusAndCreated(offset, limit, iFilename,iStatus,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByFilenameAndStatusAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByFilenameAndStatusAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFilename := self.Args("filename").String()
	iStatus := self.Args("status").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iFilename) {
		_FileManaged, _error := model.GetFileManagedsByFilenameAndStatusAndChanged(offset, limit, iFilename,iStatus,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByFilenameAndStatusAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByFilenameAndCreatedAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFilename := self.Args("filename").String()
	iCreated := self.Args("created").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iFilename) {
		_FileManaged, _error := model.GetFileManagedsByFilenameAndCreatedAndChanged(offset, limit, iFilename,iCreated,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByFilenameAndCreatedAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByUriAndFilemimeAndFilesizeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUri := self.Args("uri").String()
	iFilemime := self.Args("filemime").String()
	iFilesize := self.Args("filesize").MustInt64()

	if helper.IsHas(iUri) {
		_FileManaged, _error := model.GetFileManagedsByUriAndFilemimeAndFilesize(offset, limit, iUri,iFilemime,iFilesize)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByUriAndFilemimeAndFilesize's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByUriAndFilemimeAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUri := self.Args("uri").String()
	iFilemime := self.Args("filemime").String()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iUri) {
		_FileManaged, _error := model.GetFileManagedsByUriAndFilemimeAndStatus(offset, limit, iUri,iFilemime,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByUriAndFilemimeAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByUriAndFilemimeAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUri := self.Args("uri").String()
	iFilemime := self.Args("filemime").String()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iUri) {
		_FileManaged, _error := model.GetFileManagedsByUriAndFilemimeAndCreated(offset, limit, iUri,iFilemime,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByUriAndFilemimeAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByUriAndFilemimeAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUri := self.Args("uri").String()
	iFilemime := self.Args("filemime").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iUri) {
		_FileManaged, _error := model.GetFileManagedsByUriAndFilemimeAndChanged(offset, limit, iUri,iFilemime,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByUriAndFilemimeAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByUriAndFilesizeAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUri := self.Args("uri").String()
	iFilesize := self.Args("filesize").MustInt64()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iUri) {
		_FileManaged, _error := model.GetFileManagedsByUriAndFilesizeAndStatus(offset, limit, iUri,iFilesize,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByUriAndFilesizeAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByUriAndFilesizeAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUri := self.Args("uri").String()
	iFilesize := self.Args("filesize").MustInt64()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iUri) {
		_FileManaged, _error := model.GetFileManagedsByUriAndFilesizeAndCreated(offset, limit, iUri,iFilesize,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByUriAndFilesizeAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByUriAndFilesizeAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUri := self.Args("uri").String()
	iFilesize := self.Args("filesize").MustInt64()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iUri) {
		_FileManaged, _error := model.GetFileManagedsByUriAndFilesizeAndChanged(offset, limit, iUri,iFilesize,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByUriAndFilesizeAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByUriAndStatusAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUri := self.Args("uri").String()
	iStatus := self.Args("status").MustInt()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iUri) {
		_FileManaged, _error := model.GetFileManagedsByUriAndStatusAndCreated(offset, limit, iUri,iStatus,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByUriAndStatusAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByUriAndStatusAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUri := self.Args("uri").String()
	iStatus := self.Args("status").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iUri) {
		_FileManaged, _error := model.GetFileManagedsByUriAndStatusAndChanged(offset, limit, iUri,iStatus,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByUriAndStatusAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByUriAndCreatedAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUri := self.Args("uri").String()
	iCreated := self.Args("created").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iUri) {
		_FileManaged, _error := model.GetFileManagedsByUriAndCreatedAndChanged(offset, limit, iUri,iCreated,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByUriAndCreatedAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByFilemimeAndFilesizeAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFilemime := self.Args("filemime").String()
	iFilesize := self.Args("filesize").MustInt64()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iFilemime) {
		_FileManaged, _error := model.GetFileManagedsByFilemimeAndFilesizeAndStatus(offset, limit, iFilemime,iFilesize,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByFilemimeAndFilesizeAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByFilemimeAndFilesizeAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFilemime := self.Args("filemime").String()
	iFilesize := self.Args("filesize").MustInt64()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iFilemime) {
		_FileManaged, _error := model.GetFileManagedsByFilemimeAndFilesizeAndCreated(offset, limit, iFilemime,iFilesize,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByFilemimeAndFilesizeAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByFilemimeAndFilesizeAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFilemime := self.Args("filemime").String()
	iFilesize := self.Args("filesize").MustInt64()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iFilemime) {
		_FileManaged, _error := model.GetFileManagedsByFilemimeAndFilesizeAndChanged(offset, limit, iFilemime,iFilesize,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByFilemimeAndFilesizeAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByFilemimeAndStatusAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFilemime := self.Args("filemime").String()
	iStatus := self.Args("status").MustInt()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iFilemime) {
		_FileManaged, _error := model.GetFileManagedsByFilemimeAndStatusAndCreated(offset, limit, iFilemime,iStatus,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByFilemimeAndStatusAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByFilemimeAndStatusAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFilemime := self.Args("filemime").String()
	iStatus := self.Args("status").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iFilemime) {
		_FileManaged, _error := model.GetFileManagedsByFilemimeAndStatusAndChanged(offset, limit, iFilemime,iStatus,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByFilemimeAndStatusAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByFilemimeAndCreatedAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFilemime := self.Args("filemime").String()
	iCreated := self.Args("created").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iFilemime) {
		_FileManaged, _error := model.GetFileManagedsByFilemimeAndCreatedAndChanged(offset, limit, iFilemime,iCreated,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByFilemimeAndCreatedAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByFilesizeAndStatusAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFilesize := self.Args("filesize").MustInt64()
	iStatus := self.Args("status").MustInt()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iFilesize) {
		_FileManaged, _error := model.GetFileManagedsByFilesizeAndStatusAndCreated(offset, limit, iFilesize,iStatus,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByFilesizeAndStatusAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByFilesizeAndStatusAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFilesize := self.Args("filesize").MustInt64()
	iStatus := self.Args("status").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iFilesize) {
		_FileManaged, _error := model.GetFileManagedsByFilesizeAndStatusAndChanged(offset, limit, iFilesize,iStatus,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByFilesizeAndStatusAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByFilesizeAndCreatedAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFilesize := self.Args("filesize").MustInt64()
	iCreated := self.Args("created").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iFilesize) {
		_FileManaged, _error := model.GetFileManagedsByFilesizeAndCreatedAndChanged(offset, limit, iFilesize,iCreated,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByFilesizeAndCreatedAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByStatusAndCreatedAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iStatus := self.Args("status").MustInt()
	iCreated := self.Args("created").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iStatus) {
		_FileManaged, _error := model.GetFileManagedsByStatusAndCreatedAndChanged(offset, limit, iStatus,iCreated,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByStatusAndCreatedAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByFidAndUuidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFid := self.Args("fid").MustInt64()
	iUuid := self.Args("uuid").String()

	if helper.IsHas(iFid) {
		_FileManaged, _error := model.GetFileManagedsByFidAndUuid(offset, limit, iFid,iUuid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByFidAndUuid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByFidAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFid := self.Args("fid").MustInt64()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iFid) {
		_FileManaged, _error := model.GetFileManagedsByFidAndLangcode(offset, limit, iFid,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByFidAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByFidAndUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFid := self.Args("fid").MustInt64()
	iUid := self.Args("uid").MustInt()

	if helper.IsHas(iFid) {
		_FileManaged, _error := model.GetFileManagedsByFidAndUid(offset, limit, iFid,iUid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByFidAndUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByFidAndFilenameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFid := self.Args("fid").MustInt64()
	iFilename := self.Args("filename").String()

	if helper.IsHas(iFid) {
		_FileManaged, _error := model.GetFileManagedsByFidAndFilename(offset, limit, iFid,iFilename)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByFidAndFilename's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByFidAndUriHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFid := self.Args("fid").MustInt64()
	iUri := self.Args("uri").String()

	if helper.IsHas(iFid) {
		_FileManaged, _error := model.GetFileManagedsByFidAndUri(offset, limit, iFid,iUri)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByFidAndUri's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByFidAndFilemimeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFid := self.Args("fid").MustInt64()
	iFilemime := self.Args("filemime").String()

	if helper.IsHas(iFid) {
		_FileManaged, _error := model.GetFileManagedsByFidAndFilemime(offset, limit, iFid,iFilemime)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByFidAndFilemime's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByFidAndFilesizeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFid := self.Args("fid").MustInt64()
	iFilesize := self.Args("filesize").MustInt64()

	if helper.IsHas(iFid) {
		_FileManaged, _error := model.GetFileManagedsByFidAndFilesize(offset, limit, iFid,iFilesize)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByFidAndFilesize's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByFidAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFid := self.Args("fid").MustInt64()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iFid) {
		_FileManaged, _error := model.GetFileManagedsByFidAndStatus(offset, limit, iFid,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByFidAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByFidAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFid := self.Args("fid").MustInt64()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iFid) {
		_FileManaged, _error := model.GetFileManagedsByFidAndCreated(offset, limit, iFid,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByFidAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByFidAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFid := self.Args("fid").MustInt64()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iFid) {
		_FileManaged, _error := model.GetFileManagedsByFidAndChanged(offset, limit, iFid,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByFidAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByUuidAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUuid := self.Args("uuid").String()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iUuid) {
		_FileManaged, _error := model.GetFileManagedsByUuidAndLangcode(offset, limit, iUuid,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByUuidAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByUuidAndUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUuid := self.Args("uuid").String()
	iUid := self.Args("uid").MustInt()

	if helper.IsHas(iUuid) {
		_FileManaged, _error := model.GetFileManagedsByUuidAndUid(offset, limit, iUuid,iUid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByUuidAndUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByUuidAndFilenameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUuid := self.Args("uuid").String()
	iFilename := self.Args("filename").String()

	if helper.IsHas(iUuid) {
		_FileManaged, _error := model.GetFileManagedsByUuidAndFilename(offset, limit, iUuid,iFilename)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByUuidAndFilename's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByUuidAndUriHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUuid := self.Args("uuid").String()
	iUri := self.Args("uri").String()

	if helper.IsHas(iUuid) {
		_FileManaged, _error := model.GetFileManagedsByUuidAndUri(offset, limit, iUuid,iUri)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByUuidAndUri's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByUuidAndFilemimeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUuid := self.Args("uuid").String()
	iFilemime := self.Args("filemime").String()

	if helper.IsHas(iUuid) {
		_FileManaged, _error := model.GetFileManagedsByUuidAndFilemime(offset, limit, iUuid,iFilemime)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByUuidAndFilemime's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByUuidAndFilesizeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUuid := self.Args("uuid").String()
	iFilesize := self.Args("filesize").MustInt64()

	if helper.IsHas(iUuid) {
		_FileManaged, _error := model.GetFileManagedsByUuidAndFilesize(offset, limit, iUuid,iFilesize)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByUuidAndFilesize's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByUuidAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUuid := self.Args("uuid").String()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iUuid) {
		_FileManaged, _error := model.GetFileManagedsByUuidAndStatus(offset, limit, iUuid,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByUuidAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByUuidAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUuid := self.Args("uuid").String()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iUuid) {
		_FileManaged, _error := model.GetFileManagedsByUuidAndCreated(offset, limit, iUuid,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByUuidAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByUuidAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUuid := self.Args("uuid").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iUuid) {
		_FileManaged, _error := model.GetFileManagedsByUuidAndChanged(offset, limit, iUuid,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByUuidAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByLangcodeAndUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iUid := self.Args("uid").MustInt()

	if helper.IsHas(iLangcode) {
		_FileManaged, _error := model.GetFileManagedsByLangcodeAndUid(offset, limit, iLangcode,iUid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByLangcodeAndUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByLangcodeAndFilenameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iFilename := self.Args("filename").String()

	if helper.IsHas(iLangcode) {
		_FileManaged, _error := model.GetFileManagedsByLangcodeAndFilename(offset, limit, iLangcode,iFilename)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByLangcodeAndFilename's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByLangcodeAndUriHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iUri := self.Args("uri").String()

	if helper.IsHas(iLangcode) {
		_FileManaged, _error := model.GetFileManagedsByLangcodeAndUri(offset, limit, iLangcode,iUri)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByLangcodeAndUri's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByLangcodeAndFilemimeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iFilemime := self.Args("filemime").String()

	if helper.IsHas(iLangcode) {
		_FileManaged, _error := model.GetFileManagedsByLangcodeAndFilemime(offset, limit, iLangcode,iFilemime)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByLangcodeAndFilemime's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByLangcodeAndFilesizeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iFilesize := self.Args("filesize").MustInt64()

	if helper.IsHas(iLangcode) {
		_FileManaged, _error := model.GetFileManagedsByLangcodeAndFilesize(offset, limit, iLangcode,iFilesize)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByLangcodeAndFilesize's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByLangcodeAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iLangcode) {
		_FileManaged, _error := model.GetFileManagedsByLangcodeAndStatus(offset, limit, iLangcode,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByLangcodeAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByLangcodeAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iLangcode) {
		_FileManaged, _error := model.GetFileManagedsByLangcodeAndCreated(offset, limit, iLangcode,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByLangcodeAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByLangcodeAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iLangcode) {
		_FileManaged, _error := model.GetFileManagedsByLangcodeAndChanged(offset, limit, iLangcode,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByLangcodeAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByUidAndFilenameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iFilename := self.Args("filename").String()

	if helper.IsHas(iUid) {
		_FileManaged, _error := model.GetFileManagedsByUidAndFilename(offset, limit, iUid,iFilename)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByUidAndFilename's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByUidAndUriHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iUri := self.Args("uri").String()

	if helper.IsHas(iUid) {
		_FileManaged, _error := model.GetFileManagedsByUidAndUri(offset, limit, iUid,iUri)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByUidAndUri's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByUidAndFilemimeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iFilemime := self.Args("filemime").String()

	if helper.IsHas(iUid) {
		_FileManaged, _error := model.GetFileManagedsByUidAndFilemime(offset, limit, iUid,iFilemime)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByUidAndFilemime's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByUidAndFilesizeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iFilesize := self.Args("filesize").MustInt64()

	if helper.IsHas(iUid) {
		_FileManaged, _error := model.GetFileManagedsByUidAndFilesize(offset, limit, iUid,iFilesize)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByUidAndFilesize's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByUidAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iUid) {
		_FileManaged, _error := model.GetFileManagedsByUidAndStatus(offset, limit, iUid,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByUidAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByUidAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iUid) {
		_FileManaged, _error := model.GetFileManagedsByUidAndCreated(offset, limit, iUid,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByUidAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByUidAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iUid) {
		_FileManaged, _error := model.GetFileManagedsByUidAndChanged(offset, limit, iUid,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByUidAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByFilenameAndUriHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFilename := self.Args("filename").String()
	iUri := self.Args("uri").String()

	if helper.IsHas(iFilename) {
		_FileManaged, _error := model.GetFileManagedsByFilenameAndUri(offset, limit, iFilename,iUri)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByFilenameAndUri's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByFilenameAndFilemimeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFilename := self.Args("filename").String()
	iFilemime := self.Args("filemime").String()

	if helper.IsHas(iFilename) {
		_FileManaged, _error := model.GetFileManagedsByFilenameAndFilemime(offset, limit, iFilename,iFilemime)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByFilenameAndFilemime's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByFilenameAndFilesizeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFilename := self.Args("filename").String()
	iFilesize := self.Args("filesize").MustInt64()

	if helper.IsHas(iFilename) {
		_FileManaged, _error := model.GetFileManagedsByFilenameAndFilesize(offset, limit, iFilename,iFilesize)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByFilenameAndFilesize's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByFilenameAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFilename := self.Args("filename").String()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iFilename) {
		_FileManaged, _error := model.GetFileManagedsByFilenameAndStatus(offset, limit, iFilename,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByFilenameAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByFilenameAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFilename := self.Args("filename").String()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iFilename) {
		_FileManaged, _error := model.GetFileManagedsByFilenameAndCreated(offset, limit, iFilename,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByFilenameAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByFilenameAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFilename := self.Args("filename").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iFilename) {
		_FileManaged, _error := model.GetFileManagedsByFilenameAndChanged(offset, limit, iFilename,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByFilenameAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByUriAndFilemimeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUri := self.Args("uri").String()
	iFilemime := self.Args("filemime").String()

	if helper.IsHas(iUri) {
		_FileManaged, _error := model.GetFileManagedsByUriAndFilemime(offset, limit, iUri,iFilemime)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByUriAndFilemime's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByUriAndFilesizeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUri := self.Args("uri").String()
	iFilesize := self.Args("filesize").MustInt64()

	if helper.IsHas(iUri) {
		_FileManaged, _error := model.GetFileManagedsByUriAndFilesize(offset, limit, iUri,iFilesize)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByUriAndFilesize's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByUriAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUri := self.Args("uri").String()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iUri) {
		_FileManaged, _error := model.GetFileManagedsByUriAndStatus(offset, limit, iUri,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByUriAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByUriAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUri := self.Args("uri").String()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iUri) {
		_FileManaged, _error := model.GetFileManagedsByUriAndCreated(offset, limit, iUri,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByUriAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByUriAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUri := self.Args("uri").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iUri) {
		_FileManaged, _error := model.GetFileManagedsByUriAndChanged(offset, limit, iUri,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByUriAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByFilemimeAndFilesizeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFilemime := self.Args("filemime").String()
	iFilesize := self.Args("filesize").MustInt64()

	if helper.IsHas(iFilemime) {
		_FileManaged, _error := model.GetFileManagedsByFilemimeAndFilesize(offset, limit, iFilemime,iFilesize)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByFilemimeAndFilesize's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByFilemimeAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFilemime := self.Args("filemime").String()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iFilemime) {
		_FileManaged, _error := model.GetFileManagedsByFilemimeAndStatus(offset, limit, iFilemime,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByFilemimeAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByFilemimeAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFilemime := self.Args("filemime").String()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iFilemime) {
		_FileManaged, _error := model.GetFileManagedsByFilemimeAndCreated(offset, limit, iFilemime,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByFilemimeAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByFilemimeAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFilemime := self.Args("filemime").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iFilemime) {
		_FileManaged, _error := model.GetFileManagedsByFilemimeAndChanged(offset, limit, iFilemime,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByFilemimeAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByFilesizeAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFilesize := self.Args("filesize").MustInt64()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iFilesize) {
		_FileManaged, _error := model.GetFileManagedsByFilesizeAndStatus(offset, limit, iFilesize,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByFilesizeAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByFilesizeAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFilesize := self.Args("filesize").MustInt64()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iFilesize) {
		_FileManaged, _error := model.GetFileManagedsByFilesizeAndCreated(offset, limit, iFilesize,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByFilesizeAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByFilesizeAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFilesize := self.Args("filesize").MustInt64()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iFilesize) {
		_FileManaged, _error := model.GetFileManagedsByFilesizeAndChanged(offset, limit, iFilesize,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByFilesizeAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByStatusAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iStatus := self.Args("status").MustInt()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iStatus) {
		_FileManaged, _error := model.GetFileManagedsByStatusAndCreated(offset, limit, iStatus,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByStatusAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByStatusAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iStatus := self.Args("status").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iStatus) {
		_FileManaged, _error := model.GetFileManagedsByStatusAndChanged(offset, limit, iStatus,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByStatusAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsByCreatedAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCreated := self.Args("created").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iCreated) {
		_FileManaged, _error := model.GetFileManagedsByCreatedAndChanged(offset, limit, iCreated,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedsByCreatedAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	if (offset > 0) && (len(field) > 0) {
		_FileManaged, _error := model.GetFileManageds(offset, limit, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManageds' args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasFileManagedViaFidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iFid := self.Args("fid").MustInt64()
	if helper.IsHas(iFid) {
		_FileManaged := model.HasFileManagedViaFid(iFid)
		var m = map[string]interface{}{}
		m["file_managed"] = _FileManaged
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasFileManagedViaFid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasFileManagedViaUuidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iUuid := self.Args("uuid").String()
	if helper.IsHas(iUuid) {
		_FileManaged := model.HasFileManagedViaUuid(iUuid)
		var m = map[string]interface{}{}
		m["file_managed"] = _FileManaged
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasFileManagedViaUuid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasFileManagedViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iLangcode := self.Args("langcode").String()
	if helper.IsHas(iLangcode) {
		_FileManaged := model.HasFileManagedViaLangcode(iLangcode)
		var m = map[string]interface{}{}
		m["file_managed"] = _FileManaged
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasFileManagedViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasFileManagedViaUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iUid := self.Args("uid").MustInt()
	if helper.IsHas(iUid) {
		_FileManaged := model.HasFileManagedViaUid(iUid)
		var m = map[string]interface{}{}
		m["file_managed"] = _FileManaged
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasFileManagedViaUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasFileManagedViaFilenameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iFilename := self.Args("filename").String()
	if helper.IsHas(iFilename) {
		_FileManaged := model.HasFileManagedViaFilename(iFilename)
		var m = map[string]interface{}{}
		m["file_managed"] = _FileManaged
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasFileManagedViaFilename's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasFileManagedViaUriHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iUri := self.Args("uri").String()
	if helper.IsHas(iUri) {
		_FileManaged := model.HasFileManagedViaUri(iUri)
		var m = map[string]interface{}{}
		m["file_managed"] = _FileManaged
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasFileManagedViaUri's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasFileManagedViaFilemimeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iFilemime := self.Args("filemime").String()
	if helper.IsHas(iFilemime) {
		_FileManaged := model.HasFileManagedViaFilemime(iFilemime)
		var m = map[string]interface{}{}
		m["file_managed"] = _FileManaged
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasFileManagedViaFilemime's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasFileManagedViaFilesizeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iFilesize := self.Args("filesize").MustInt64()
	if helper.IsHas(iFilesize) {
		_FileManaged := model.HasFileManagedViaFilesize(iFilesize)
		var m = map[string]interface{}{}
		m["file_managed"] = _FileManaged
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasFileManagedViaFilesize's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasFileManagedViaStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iStatus := self.Args("status").MustInt()
	if helper.IsHas(iStatus) {
		_FileManaged := model.HasFileManagedViaStatus(iStatus)
		var m = map[string]interface{}{}
		m["file_managed"] = _FileManaged
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasFileManagedViaStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasFileManagedViaCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iCreated := self.Args("created").MustInt()
	if helper.IsHas(iCreated) {
		_FileManaged := model.HasFileManagedViaCreated(iCreated)
		var m = map[string]interface{}{}
		m["file_managed"] = _FileManaged
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasFileManagedViaCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasFileManagedViaChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iChanged := self.Args("changed").MustInt()
	if helper.IsHas(iChanged) {
		_FileManaged := model.HasFileManagedViaChanged(iChanged)
		var m = map[string]interface{}{}
		m["file_managed"] = _FileManaged
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasFileManagedViaChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedViaFidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iFid := self.Args("fid").MustInt64()
	if helper.IsHas(iFid) {
		_FileManaged, _error := model.GetFileManagedViaFid(iFid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedViaFid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedViaUuidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iUuid := self.Args("uuid").String()
	if helper.IsHas(iUuid) {
		_FileManaged, _error := model.GetFileManagedViaUuid(iUuid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedViaUuid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iLangcode := self.Args("langcode").String()
	if helper.IsHas(iLangcode) {
		_FileManaged, _error := model.GetFileManagedViaLangcode(iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedViaUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iUid := self.Args("uid").MustInt()
	if helper.IsHas(iUid) {
		_FileManaged, _error := model.GetFileManagedViaUid(iUid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedViaUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedViaFilenameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iFilename := self.Args("filename").String()
	if helper.IsHas(iFilename) {
		_FileManaged, _error := model.GetFileManagedViaFilename(iFilename)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedViaFilename's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedViaUriHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iUri := self.Args("uri").String()
	if helper.IsHas(iUri) {
		_FileManaged, _error := model.GetFileManagedViaUri(iUri)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedViaUri's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedViaFilemimeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iFilemime := self.Args("filemime").String()
	if helper.IsHas(iFilemime) {
		_FileManaged, _error := model.GetFileManagedViaFilemime(iFilemime)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedViaFilemime's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedViaFilesizeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iFilesize := self.Args("filesize").MustInt64()
	if helper.IsHas(iFilesize) {
		_FileManaged, _error := model.GetFileManagedViaFilesize(iFilesize)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedViaFilesize's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedViaStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iStatus := self.Args("status").MustInt()
	if helper.IsHas(iStatus) {
		_FileManaged, _error := model.GetFileManagedViaStatus(iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedViaStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedViaCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iCreated := self.Args("created").MustInt()
	if helper.IsHas(iCreated) {
		_FileManaged, _error := model.GetFileManagedViaCreated(iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedViaCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileManagedViaChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iChanged := self.Args("changed").MustInt()
	if helper.IsHas(iChanged) {
		_FileManaged, _error := model.GetFileManagedViaChanged(iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the GetFileManagedViaChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetFileManagedViaFidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Fid_ := self.Args("fid").MustInt64()
	if helper.IsHas(Fid_) {
		var iFileManaged model.FileManaged
		self.Bind(&iFileManaged)
		_FileManaged, _error := model.SetFileManagedViaFid(Fid_, &iFileManaged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the SetFileManagedViaFid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetFileManagedViaUuidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Uuid_ := self.Args("uuid").String()
	if helper.IsHas(Uuid_) {
		var iFileManaged model.FileManaged
		self.Bind(&iFileManaged)
		_FileManaged, _error := model.SetFileManagedViaUuid(Uuid_, &iFileManaged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the SetFileManagedViaUuid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetFileManagedViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	if helper.IsHas(Langcode_) {
		var iFileManaged model.FileManaged
		self.Bind(&iFileManaged)
		_FileManaged, _error := model.SetFileManagedViaLangcode(Langcode_, &iFileManaged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the SetFileManagedViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetFileManagedViaUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Uid_ := self.Args("uid").MustInt()
	if helper.IsHas(Uid_) {
		var iFileManaged model.FileManaged
		self.Bind(&iFileManaged)
		_FileManaged, _error := model.SetFileManagedViaUid(Uid_, &iFileManaged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the SetFileManagedViaUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetFileManagedViaFilenameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Filename_ := self.Args("filename").String()
	if helper.IsHas(Filename_) {
		var iFileManaged model.FileManaged
		self.Bind(&iFileManaged)
		_FileManaged, _error := model.SetFileManagedViaFilename(Filename_, &iFileManaged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the SetFileManagedViaFilename's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetFileManagedViaUriHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Uri_ := self.Args("uri").String()
	if helper.IsHas(Uri_) {
		var iFileManaged model.FileManaged
		self.Bind(&iFileManaged)
		_FileManaged, _error := model.SetFileManagedViaUri(Uri_, &iFileManaged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the SetFileManagedViaUri's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetFileManagedViaFilemimeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Filemime_ := self.Args("filemime").String()
	if helper.IsHas(Filemime_) {
		var iFileManaged model.FileManaged
		self.Bind(&iFileManaged)
		_FileManaged, _error := model.SetFileManagedViaFilemime(Filemime_, &iFileManaged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the SetFileManagedViaFilemime's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetFileManagedViaFilesizeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Filesize_ := self.Args("filesize").MustInt64()
	if helper.IsHas(Filesize_) {
		var iFileManaged model.FileManaged
		self.Bind(&iFileManaged)
		_FileManaged, _error := model.SetFileManagedViaFilesize(Filesize_, &iFileManaged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the SetFileManagedViaFilesize's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetFileManagedViaStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Status_ := self.Args("status").MustInt()
	if helper.IsHas(Status_) {
		var iFileManaged model.FileManaged
		self.Bind(&iFileManaged)
		_FileManaged, _error := model.SetFileManagedViaStatus(Status_, &iFileManaged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the SetFileManagedViaStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetFileManagedViaCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Created_ := self.Args("created").MustInt()
	if helper.IsHas(Created_) {
		var iFileManaged model.FileManaged
		self.Bind(&iFileManaged)
		_FileManaged, _error := model.SetFileManagedViaCreated(Created_, &iFileManaged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the SetFileManagedViaCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetFileManagedViaChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Changed_ := self.Args("changed").MustInt()
	if helper.IsHas(Changed_) {
		var iFileManaged model.FileManaged
		self.Bind(&iFileManaged)
		_FileManaged, _error := model.SetFileManagedViaChanged(Changed_, &iFileManaged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileManaged)
	}
	herr.Message = "Can't get to the SetFileManagedViaChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostAddFileManagedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Fid_ := self.Args("fid").MustInt64()
	Uuid_ := self.Args("uuid").String()
	Langcode_ := self.Args("langcode").String()
	Uid_ := self.Args("uid").MustInt()
	Filename_ := self.Args("filename").String()
	Uri_ := self.Args("uri").String()
	Filemime_ := self.Args("filemime").String()
	Filesize_ := self.Args("filesize").MustInt64()
	Status_ := self.Args("status").MustInt()
	Created_ := self.Args("created").MustInt()
	Changed_ := self.Args("changed").MustInt()

	if helper.IsHas(Fid_) {
		_error := model.AddFileManaged(Fid_,Uuid_,Langcode_,Uid_,Filename_,Uri_,Filemime_,Filesize_,Status_,Created_,Changed_)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		herr.Message = "StatusOK"
		herr.Status = macross.StatusOK
		return self.JSON(herr)
	}
	herr.Message = "Can't get to the AddFileManaged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostFileManagedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iFileManaged model.FileManaged
	self.Bind(&iFileManaged)
	_int64, _error := model.PostFileManaged(&iFileManaged)
	if (helper.IsHas(_int64)) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	} else {
		var m = map[string]interface{}{}
		m["created"] = _int64
		return self.JSON(m, macross.StatusCreated)
	}
	return self.JSON(herr)
}

func PutFileManagedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iFileManaged model.FileManaged
	self.Bind(&iFileManaged)
	_int64, _error := model.PutFileManaged(&iFileManaged)
	if (helper.IsHas(_int64)) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	} else {
		var m = map[string]interface{}{}
		m["updated"] = _int64
		return self.JSON(m)
	}
	return self.JSON(herr)
}

func PutFileManagedViaFidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Fid_ := self.Args("fid").MustInt64()
	var iFileManaged model.FileManaged
	self.Bind(&iFileManaged)
	_int64, _error := model.PutFileManagedViaFid(Fid_, &iFileManaged)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutFileManagedViaUuidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Uuid_ := self.Args("uuid").String()
	var iFileManaged model.FileManaged
	self.Bind(&iFileManaged)
	_int64, _error := model.PutFileManagedViaUuid(Uuid_, &iFileManaged)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutFileManagedViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	var iFileManaged model.FileManaged
	self.Bind(&iFileManaged)
	_int64, _error := model.PutFileManagedViaLangcode(Langcode_, &iFileManaged)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutFileManagedViaUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Uid_ := self.Args("uid").MustInt()
	var iFileManaged model.FileManaged
	self.Bind(&iFileManaged)
	_int64, _error := model.PutFileManagedViaUid(Uid_, &iFileManaged)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutFileManagedViaFilenameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Filename_ := self.Args("filename").String()
	var iFileManaged model.FileManaged
	self.Bind(&iFileManaged)
	_int64, _error := model.PutFileManagedViaFilename(Filename_, &iFileManaged)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutFileManagedViaUriHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Uri_ := self.Args("uri").String()
	var iFileManaged model.FileManaged
	self.Bind(&iFileManaged)
	_int64, _error := model.PutFileManagedViaUri(Uri_, &iFileManaged)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutFileManagedViaFilemimeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Filemime_ := self.Args("filemime").String()
	var iFileManaged model.FileManaged
	self.Bind(&iFileManaged)
	_int64, _error := model.PutFileManagedViaFilemime(Filemime_, &iFileManaged)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutFileManagedViaFilesizeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Filesize_ := self.Args("filesize").MustInt64()
	var iFileManaged model.FileManaged
	self.Bind(&iFileManaged)
	_int64, _error := model.PutFileManagedViaFilesize(Filesize_, &iFileManaged)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutFileManagedViaStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Status_ := self.Args("status").MustInt()
	var iFileManaged model.FileManaged
	self.Bind(&iFileManaged)
	_int64, _error := model.PutFileManagedViaStatus(Status_, &iFileManaged)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutFileManagedViaCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Created_ := self.Args("created").MustInt()
	var iFileManaged model.FileManaged
	self.Bind(&iFileManaged)
	_int64, _error := model.PutFileManagedViaCreated(Created_, &iFileManaged)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutFileManagedViaChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Changed_ := self.Args("changed").MustInt()
	var iFileManaged model.FileManaged
	self.Bind(&iFileManaged)
	_int64, _error := model.PutFileManagedViaChanged(Changed_, &iFileManaged)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateFileManagedViaFidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Fid_ := self.Args("fid").MustInt64()
	var iFileManaged model.FileManaged
	self.Bind(&iFileManaged)
	var iMap = helper.StructToMap(iFileManaged)
	_error := model.UpdateFileManagedViaFid(Fid_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateFileManagedViaUuidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Uuid_ := self.Args("uuid").String()
	var iFileManaged model.FileManaged
	self.Bind(&iFileManaged)
	var iMap = helper.StructToMap(iFileManaged)
	_error := model.UpdateFileManagedViaUuid(Uuid_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateFileManagedViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	var iFileManaged model.FileManaged
	self.Bind(&iFileManaged)
	var iMap = helper.StructToMap(iFileManaged)
	_error := model.UpdateFileManagedViaLangcode(Langcode_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateFileManagedViaUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Uid_ := self.Args("uid").MustInt()
	var iFileManaged model.FileManaged
	self.Bind(&iFileManaged)
	var iMap = helper.StructToMap(iFileManaged)
	_error := model.UpdateFileManagedViaUid(Uid_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateFileManagedViaFilenameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Filename_ := self.Args("filename").String()
	var iFileManaged model.FileManaged
	self.Bind(&iFileManaged)
	var iMap = helper.StructToMap(iFileManaged)
	_error := model.UpdateFileManagedViaFilename(Filename_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateFileManagedViaUriHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Uri_ := self.Args("uri").String()
	var iFileManaged model.FileManaged
	self.Bind(&iFileManaged)
	var iMap = helper.StructToMap(iFileManaged)
	_error := model.UpdateFileManagedViaUri(Uri_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateFileManagedViaFilemimeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Filemime_ := self.Args("filemime").String()
	var iFileManaged model.FileManaged
	self.Bind(&iFileManaged)
	var iMap = helper.StructToMap(iFileManaged)
	_error := model.UpdateFileManagedViaFilemime(Filemime_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateFileManagedViaFilesizeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Filesize_ := self.Args("filesize").MustInt64()
	var iFileManaged model.FileManaged
	self.Bind(&iFileManaged)
	var iMap = helper.StructToMap(iFileManaged)
	_error := model.UpdateFileManagedViaFilesize(Filesize_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateFileManagedViaStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Status_ := self.Args("status").MustInt()
	var iFileManaged model.FileManaged
	self.Bind(&iFileManaged)
	var iMap = helper.StructToMap(iFileManaged)
	_error := model.UpdateFileManagedViaStatus(Status_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateFileManagedViaCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Created_ := self.Args("created").MustInt()
	var iFileManaged model.FileManaged
	self.Bind(&iFileManaged)
	var iMap = helper.StructToMap(iFileManaged)
	_error := model.UpdateFileManagedViaCreated(Created_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateFileManagedViaChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Changed_ := self.Args("changed").MustInt()
	var iFileManaged model.FileManaged
	self.Bind(&iFileManaged)
	var iMap = helper.StructToMap(iFileManaged)
	_error := model.UpdateFileManagedViaChanged(Changed_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteFileManagedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Fid_ := self.Args("fid").MustInt64()
	_int64, _error := model.DeleteFileManaged(Fid_)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	} else {
		var m = map[string]interface{}{}
		m["deleted"] = _int64
		return self.JSON(m)
	}
	return self.JSON(herr)
}

func DeleteFileManagedViaFidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Fid_ := self.Args("fid").MustInt64()
	_error := model.DeleteFileManagedViaFid(Fid_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteFileManagedViaUuidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Uuid_ := self.Args("uuid").String()
	_error := model.DeleteFileManagedViaUuid(Uuid_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteFileManagedViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	_error := model.DeleteFileManagedViaLangcode(Langcode_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteFileManagedViaUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Uid_ := self.Args("uid").MustInt()
	_error := model.DeleteFileManagedViaUid(Uid_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteFileManagedViaFilenameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Filename_ := self.Args("filename").String()
	_error := model.DeleteFileManagedViaFilename(Filename_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteFileManagedViaUriHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Uri_ := self.Args("uri").String()
	_error := model.DeleteFileManagedViaUri(Uri_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteFileManagedViaFilemimeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Filemime_ := self.Args("filemime").String()
	_error := model.DeleteFileManagedViaFilemime(Filemime_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteFileManagedViaFilesizeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Filesize_ := self.Args("filesize").MustInt64()
	_error := model.DeleteFileManagedViaFilesize(Filesize_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteFileManagedViaStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Status_ := self.Args("status").MustInt()
	_error := model.DeleteFileManagedViaStatus(Status_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteFileManagedViaCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Created_ := self.Args("created").MustInt()
	_error := model.DeleteFileManagedViaCreated(Created_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteFileManagedViaChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Changed_ := self.Args("changed").MustInt()
	_error := model.DeleteFileManagedViaChanged(Changed_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}
