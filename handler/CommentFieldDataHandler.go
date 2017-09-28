package handler

import (
	"github.com/insionng/drupapp/helper"
	"github.com/insionng/drupapp/model"
	"github.com/insionng/macross"
)

func GetCommentFieldDatasCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()

	if offset > 0 {
		var m = map[string]interface{}{}
		_int64, _error := model.GetCommentFieldDatasCount(offset, limit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		m["comment_field_datasCount"] = _int64
		return self.JSON(m)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDataCountViaCidHandler(self *macross.Context) error {
	Cid_ := self.Args("cid").MustInt64()
	_int64 := model.GetCommentFieldDataCountViaCid(Cid_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["comment_field_dataCount"] = 0
	}
	m["comment_field_dataCount"] = _int64
	return self.JSON(m)
}

func GetCommentFieldDataCountViaCommentTypeHandler(self *macross.Context) error {
	CommentType_ := self.Args("comment_type").String()
	_int64 := model.GetCommentFieldDataCountViaCommentType(CommentType_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["comment_field_dataCount"] = 0
	}
	m["comment_field_dataCount"] = _int64
	return self.JSON(m)
}

func GetCommentFieldDataCountViaLangcodeHandler(self *macross.Context) error {
	Langcode_ := self.Args("langcode").String()
	_int64 := model.GetCommentFieldDataCountViaLangcode(Langcode_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["comment_field_dataCount"] = 0
	}
	m["comment_field_dataCount"] = _int64
	return self.JSON(m)
}

func GetCommentFieldDataCountViaPidHandler(self *macross.Context) error {
	Pid_ := self.Args("pid").MustInt()
	_int64 := model.GetCommentFieldDataCountViaPid(Pid_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["comment_field_dataCount"] = 0
	}
	m["comment_field_dataCount"] = _int64
	return self.JSON(m)
}

func GetCommentFieldDataCountViaEntityIdHandler(self *macross.Context) error {
	EntityId_ := self.Args("entity_id").MustInt()
	_int64 := model.GetCommentFieldDataCountViaEntityId(EntityId_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["comment_field_dataCount"] = 0
	}
	m["comment_field_dataCount"] = _int64
	return self.JSON(m)
}

func GetCommentFieldDataCountViaSubjectHandler(self *macross.Context) error {
	Subject_ := self.Args("subject").String()
	_int64 := model.GetCommentFieldDataCountViaSubject(Subject_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["comment_field_dataCount"] = 0
	}
	m["comment_field_dataCount"] = _int64
	return self.JSON(m)
}

func GetCommentFieldDataCountViaUidHandler(self *macross.Context) error {
	Uid_ := self.Args("uid").MustInt()
	_int64 := model.GetCommentFieldDataCountViaUid(Uid_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["comment_field_dataCount"] = 0
	}
	m["comment_field_dataCount"] = _int64
	return self.JSON(m)
}

func GetCommentFieldDataCountViaNameHandler(self *macross.Context) error {
	Name_ := self.Args("name").String()
	_int64 := model.GetCommentFieldDataCountViaName(Name_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["comment_field_dataCount"] = 0
	}
	m["comment_field_dataCount"] = _int64
	return self.JSON(m)
}

func GetCommentFieldDataCountViaMailHandler(self *macross.Context) error {
	Mail_ := self.Args("mail").String()
	_int64 := model.GetCommentFieldDataCountViaMail(Mail_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["comment_field_dataCount"] = 0
	}
	m["comment_field_dataCount"] = _int64
	return self.JSON(m)
}

func GetCommentFieldDataCountViaHomepageHandler(self *macross.Context) error {
	Homepage_ := self.Args("homepage").String()
	_int64 := model.GetCommentFieldDataCountViaHomepage(Homepage_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["comment_field_dataCount"] = 0
	}
	m["comment_field_dataCount"] = _int64
	return self.JSON(m)
}

func GetCommentFieldDataCountViaHostnameHandler(self *macross.Context) error {
	Hostname_ := self.Args("hostname").String()
	_int64 := model.GetCommentFieldDataCountViaHostname(Hostname_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["comment_field_dataCount"] = 0
	}
	m["comment_field_dataCount"] = _int64
	return self.JSON(m)
}

func GetCommentFieldDataCountViaCreatedHandler(self *macross.Context) error {
	Created_ := self.Args("created").MustInt()
	_int64 := model.GetCommentFieldDataCountViaCreated(Created_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["comment_field_dataCount"] = 0
	}
	m["comment_field_dataCount"] = _int64
	return self.JSON(m)
}

func GetCommentFieldDataCountViaChangedHandler(self *macross.Context) error {
	Changed_ := self.Args("changed").MustInt()
	_int64 := model.GetCommentFieldDataCountViaChanged(Changed_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["comment_field_dataCount"] = 0
	}
	m["comment_field_dataCount"] = _int64
	return self.JSON(m)
}

func GetCommentFieldDataCountViaStatusHandler(self *macross.Context) error {
	Status_ := self.Args("status").MustInt()
	_int64 := model.GetCommentFieldDataCountViaStatus(Status_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["comment_field_dataCount"] = 0
	}
	m["comment_field_dataCount"] = _int64
	return self.JSON(m)
}

func GetCommentFieldDataCountViaThreadHandler(self *macross.Context) error {
	Thread_ := self.Args("thread").String()
	_int64 := model.GetCommentFieldDataCountViaThread(Thread_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["comment_field_dataCount"] = 0
	}
	m["comment_field_dataCount"] = _int64
	return self.JSON(m)
}

func GetCommentFieldDataCountViaEntityTypeHandler(self *macross.Context) error {
	EntityType_ := self.Args("entity_type").String()
	_int64 := model.GetCommentFieldDataCountViaEntityType(EntityType_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["comment_field_dataCount"] = 0
	}
	m["comment_field_dataCount"] = _int64
	return self.JSON(m)
}

func GetCommentFieldDataCountViaFieldNameHandler(self *macross.Context) error {
	FieldName_ := self.Args("field_name").String()
	_int64 := model.GetCommentFieldDataCountViaFieldName(FieldName_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["comment_field_dataCount"] = 0
	}
	m["comment_field_dataCount"] = _int64
	return self.JSON(m)
}

func GetCommentFieldDataCountViaDefaultLangcodeHandler(self *macross.Context) error {
	DefaultLangcode_ := self.Args("default_langcode").MustInt()
	_int64 := model.GetCommentFieldDataCountViaDefaultLangcode(DefaultLangcode_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["comment_field_dataCount"] = 0
	}
	m["comment_field_dataCount"] = _int64
	return self.JSON(m)
}

func GetCommentFieldDatasViaCidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iCid := self.Args("cid").MustInt64()
	if (offset > 0) && helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasViaCid(offset, limit, iCid, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasViaCid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasViaCommentTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iCommentType := self.Args("comment_type").String()
	if (offset > 0) && helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasViaCommentType(offset, limit, iCommentType, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasViaCommentType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iLangcode := self.Args("langcode").String()
	if (offset > 0) && helper.IsHas(iLangcode) {
		_CommentFieldData, _error := model.GetCommentFieldDatasViaLangcode(offset, limit, iLangcode, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasViaPidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iPid := self.Args("pid").MustInt()
	if (offset > 0) && helper.IsHas(iPid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasViaPid(offset, limit, iPid, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasViaPid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasViaEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iEntityId := self.Args("entity_id").MustInt()
	if (offset > 0) && helper.IsHas(iEntityId) {
		_CommentFieldData, _error := model.GetCommentFieldDatasViaEntityId(offset, limit, iEntityId, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasViaEntityId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasViaSubjectHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iSubject := self.Args("subject").String()
	if (offset > 0) && helper.IsHas(iSubject) {
		_CommentFieldData, _error := model.GetCommentFieldDatasViaSubject(offset, limit, iSubject, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasViaSubject's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasViaUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iUid := self.Args("uid").MustInt()
	if (offset > 0) && helper.IsHas(iUid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasViaUid(offset, limit, iUid, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasViaUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasViaNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iName := self.Args("name").String()
	if (offset > 0) && helper.IsHas(iName) {
		_CommentFieldData, _error := model.GetCommentFieldDatasViaName(offset, limit, iName, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasViaName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasViaMailHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iMail := self.Args("mail").String()
	if (offset > 0) && helper.IsHas(iMail) {
		_CommentFieldData, _error := model.GetCommentFieldDatasViaMail(offset, limit, iMail, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasViaMail's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasViaHomepageHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iHomepage := self.Args("homepage").String()
	if (offset > 0) && helper.IsHas(iHomepage) {
		_CommentFieldData, _error := model.GetCommentFieldDatasViaHomepage(offset, limit, iHomepage, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasViaHomepage's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasViaHostnameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iHostname := self.Args("hostname").String()
	if (offset > 0) && helper.IsHas(iHostname) {
		_CommentFieldData, _error := model.GetCommentFieldDatasViaHostname(offset, limit, iHostname, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasViaHostname's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasViaCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iCreated := self.Args("created").MustInt()
	if (offset > 0) && helper.IsHas(iCreated) {
		_CommentFieldData, _error := model.GetCommentFieldDatasViaCreated(offset, limit, iCreated, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasViaCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasViaChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iChanged := self.Args("changed").MustInt()
	if (offset > 0) && helper.IsHas(iChanged) {
		_CommentFieldData, _error := model.GetCommentFieldDatasViaChanged(offset, limit, iChanged, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasViaChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasViaStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iStatus := self.Args("status").MustInt()
	if (offset > 0) && helper.IsHas(iStatus) {
		_CommentFieldData, _error := model.GetCommentFieldDatasViaStatus(offset, limit, iStatus, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasViaStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasViaThreadHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iThread := self.Args("thread").String()
	if (offset > 0) && helper.IsHas(iThread) {
		_CommentFieldData, _error := model.GetCommentFieldDatasViaThread(offset, limit, iThread, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasViaThread's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasViaEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iEntityType := self.Args("entity_type").String()
	if (offset > 0) && helper.IsHas(iEntityType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasViaEntityType(offset, limit, iEntityType, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasViaEntityType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasViaFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iFieldName := self.Args("field_name").String()
	if (offset > 0) && helper.IsHas(iFieldName) {
		_CommentFieldData, _error := model.GetCommentFieldDatasViaFieldName(offset, limit, iFieldName, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasViaFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasViaDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()
	if (offset > 0) && helper.IsHas(iDefaultLangcode) {
		_CommentFieldData, _error := model.GetCommentFieldDatasViaDefaultLangcode(offset, limit, iDefaultLangcode, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasViaDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndCommentTypeAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iCommentType := self.Args("comment_type").String()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndCommentTypeAndLangcode(offset, limit, iCid,iCommentType,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndCommentTypeAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndCommentTypeAndPidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iCommentType := self.Args("comment_type").String()
	iPid := self.Args("pid").MustInt()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndCommentTypeAndPid(offset, limit, iCid,iCommentType,iPid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndCommentTypeAndPid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndCommentTypeAndEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iCommentType := self.Args("comment_type").String()
	iEntityId := self.Args("entity_id").MustInt()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndCommentTypeAndEntityId(offset, limit, iCid,iCommentType,iEntityId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndCommentTypeAndEntityId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndCommentTypeAndSubjectHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iCommentType := self.Args("comment_type").String()
	iSubject := self.Args("subject").String()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndCommentTypeAndSubject(offset, limit, iCid,iCommentType,iSubject)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndCommentTypeAndSubject's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndCommentTypeAndUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iCommentType := self.Args("comment_type").String()
	iUid := self.Args("uid").MustInt()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndCommentTypeAndUid(offset, limit, iCid,iCommentType,iUid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndCommentTypeAndUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndCommentTypeAndNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iCommentType := self.Args("comment_type").String()
	iName := self.Args("name").String()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndCommentTypeAndName(offset, limit, iCid,iCommentType,iName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndCommentTypeAndName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndCommentTypeAndMailHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iCommentType := self.Args("comment_type").String()
	iMail := self.Args("mail").String()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndCommentTypeAndMail(offset, limit, iCid,iCommentType,iMail)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndCommentTypeAndMail's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndCommentTypeAndHomepageHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iCommentType := self.Args("comment_type").String()
	iHomepage := self.Args("homepage").String()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndCommentTypeAndHomepage(offset, limit, iCid,iCommentType,iHomepage)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndCommentTypeAndHomepage's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndCommentTypeAndHostnameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iCommentType := self.Args("comment_type").String()
	iHostname := self.Args("hostname").String()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndCommentTypeAndHostname(offset, limit, iCid,iCommentType,iHostname)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndCommentTypeAndHostname's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndCommentTypeAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iCommentType := self.Args("comment_type").String()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndCommentTypeAndCreated(offset, limit, iCid,iCommentType,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndCommentTypeAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndCommentTypeAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iCommentType := self.Args("comment_type").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndCommentTypeAndChanged(offset, limit, iCid,iCommentType,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndCommentTypeAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndCommentTypeAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iCommentType := self.Args("comment_type").String()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndCommentTypeAndStatus(offset, limit, iCid,iCommentType,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndCommentTypeAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndCommentTypeAndThreadHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iCommentType := self.Args("comment_type").String()
	iThread := self.Args("thread").String()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndCommentTypeAndThread(offset, limit, iCid,iCommentType,iThread)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndCommentTypeAndThread's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndCommentTypeAndEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iCommentType := self.Args("comment_type").String()
	iEntityType := self.Args("entity_type").String()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndCommentTypeAndEntityType(offset, limit, iCid,iCommentType,iEntityType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndCommentTypeAndEntityType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndCommentTypeAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iCommentType := self.Args("comment_type").String()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndCommentTypeAndFieldName(offset, limit, iCid,iCommentType,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndCommentTypeAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndCommentTypeAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iCommentType := self.Args("comment_type").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndCommentTypeAndDefaultLangcode(offset, limit, iCid,iCommentType,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndCommentTypeAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndLangcodeAndPidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iLangcode := self.Args("langcode").String()
	iPid := self.Args("pid").MustInt()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndLangcodeAndPid(offset, limit, iCid,iLangcode,iPid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndLangcodeAndPid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndLangcodeAndEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iLangcode := self.Args("langcode").String()
	iEntityId := self.Args("entity_id").MustInt()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndLangcodeAndEntityId(offset, limit, iCid,iLangcode,iEntityId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndLangcodeAndEntityId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndLangcodeAndSubjectHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iLangcode := self.Args("langcode").String()
	iSubject := self.Args("subject").String()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndLangcodeAndSubject(offset, limit, iCid,iLangcode,iSubject)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndLangcodeAndSubject's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndLangcodeAndUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iLangcode := self.Args("langcode").String()
	iUid := self.Args("uid").MustInt()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndLangcodeAndUid(offset, limit, iCid,iLangcode,iUid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndLangcodeAndUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndLangcodeAndNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iLangcode := self.Args("langcode").String()
	iName := self.Args("name").String()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndLangcodeAndName(offset, limit, iCid,iLangcode,iName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndLangcodeAndName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndLangcodeAndMailHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iLangcode := self.Args("langcode").String()
	iMail := self.Args("mail").String()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndLangcodeAndMail(offset, limit, iCid,iLangcode,iMail)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndLangcodeAndMail's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndLangcodeAndHomepageHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iLangcode := self.Args("langcode").String()
	iHomepage := self.Args("homepage").String()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndLangcodeAndHomepage(offset, limit, iCid,iLangcode,iHomepage)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndLangcodeAndHomepage's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndLangcodeAndHostnameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iLangcode := self.Args("langcode").String()
	iHostname := self.Args("hostname").String()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndLangcodeAndHostname(offset, limit, iCid,iLangcode,iHostname)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndLangcodeAndHostname's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndLangcodeAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iLangcode := self.Args("langcode").String()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndLangcodeAndCreated(offset, limit, iCid,iLangcode,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndLangcodeAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndLangcodeAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iLangcode := self.Args("langcode").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndLangcodeAndChanged(offset, limit, iCid,iLangcode,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndLangcodeAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndLangcodeAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iLangcode := self.Args("langcode").String()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndLangcodeAndStatus(offset, limit, iCid,iLangcode,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndLangcodeAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndLangcodeAndThreadHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iLangcode := self.Args("langcode").String()
	iThread := self.Args("thread").String()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndLangcodeAndThread(offset, limit, iCid,iLangcode,iThread)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndLangcodeAndThread's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndLangcodeAndEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iLangcode := self.Args("langcode").String()
	iEntityType := self.Args("entity_type").String()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndLangcodeAndEntityType(offset, limit, iCid,iLangcode,iEntityType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndLangcodeAndEntityType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndLangcodeAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iLangcode := self.Args("langcode").String()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndLangcodeAndFieldName(offset, limit, iCid,iLangcode,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndLangcodeAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndLangcodeAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iLangcode := self.Args("langcode").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndLangcodeAndDefaultLangcode(offset, limit, iCid,iLangcode,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndLangcodeAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndPidAndEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iPid := self.Args("pid").MustInt()
	iEntityId := self.Args("entity_id").MustInt()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndPidAndEntityId(offset, limit, iCid,iPid,iEntityId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndPidAndEntityId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndPidAndSubjectHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iPid := self.Args("pid").MustInt()
	iSubject := self.Args("subject").String()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndPidAndSubject(offset, limit, iCid,iPid,iSubject)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndPidAndSubject's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndPidAndUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iPid := self.Args("pid").MustInt()
	iUid := self.Args("uid").MustInt()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndPidAndUid(offset, limit, iCid,iPid,iUid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndPidAndUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndPidAndNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iPid := self.Args("pid").MustInt()
	iName := self.Args("name").String()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndPidAndName(offset, limit, iCid,iPid,iName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndPidAndName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndPidAndMailHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iPid := self.Args("pid").MustInt()
	iMail := self.Args("mail").String()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndPidAndMail(offset, limit, iCid,iPid,iMail)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndPidAndMail's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndPidAndHomepageHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iPid := self.Args("pid").MustInt()
	iHomepage := self.Args("homepage").String()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndPidAndHomepage(offset, limit, iCid,iPid,iHomepage)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndPidAndHomepage's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndPidAndHostnameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iPid := self.Args("pid").MustInt()
	iHostname := self.Args("hostname").String()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndPidAndHostname(offset, limit, iCid,iPid,iHostname)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndPidAndHostname's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndPidAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iPid := self.Args("pid").MustInt()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndPidAndCreated(offset, limit, iCid,iPid,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndPidAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndPidAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iPid := self.Args("pid").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndPidAndChanged(offset, limit, iCid,iPid,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndPidAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndPidAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iPid := self.Args("pid").MustInt()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndPidAndStatus(offset, limit, iCid,iPid,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndPidAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndPidAndThreadHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iPid := self.Args("pid").MustInt()
	iThread := self.Args("thread").String()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndPidAndThread(offset, limit, iCid,iPid,iThread)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndPidAndThread's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndPidAndEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iPid := self.Args("pid").MustInt()
	iEntityType := self.Args("entity_type").String()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndPidAndEntityType(offset, limit, iCid,iPid,iEntityType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndPidAndEntityType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndPidAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iPid := self.Args("pid").MustInt()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndPidAndFieldName(offset, limit, iCid,iPid,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndPidAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndPidAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iPid := self.Args("pid").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndPidAndDefaultLangcode(offset, limit, iCid,iPid,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndPidAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndEntityIdAndSubjectHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iEntityId := self.Args("entity_id").MustInt()
	iSubject := self.Args("subject").String()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndEntityIdAndSubject(offset, limit, iCid,iEntityId,iSubject)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndEntityIdAndSubject's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndEntityIdAndUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iEntityId := self.Args("entity_id").MustInt()
	iUid := self.Args("uid").MustInt()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndEntityIdAndUid(offset, limit, iCid,iEntityId,iUid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndEntityIdAndUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndEntityIdAndNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iEntityId := self.Args("entity_id").MustInt()
	iName := self.Args("name").String()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndEntityIdAndName(offset, limit, iCid,iEntityId,iName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndEntityIdAndName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndEntityIdAndMailHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iEntityId := self.Args("entity_id").MustInt()
	iMail := self.Args("mail").String()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndEntityIdAndMail(offset, limit, iCid,iEntityId,iMail)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndEntityIdAndMail's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndEntityIdAndHomepageHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iEntityId := self.Args("entity_id").MustInt()
	iHomepage := self.Args("homepage").String()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndEntityIdAndHomepage(offset, limit, iCid,iEntityId,iHomepage)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndEntityIdAndHomepage's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndEntityIdAndHostnameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iEntityId := self.Args("entity_id").MustInt()
	iHostname := self.Args("hostname").String()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndEntityIdAndHostname(offset, limit, iCid,iEntityId,iHostname)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndEntityIdAndHostname's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndEntityIdAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iEntityId := self.Args("entity_id").MustInt()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndEntityIdAndCreated(offset, limit, iCid,iEntityId,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndEntityIdAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndEntityIdAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iEntityId := self.Args("entity_id").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndEntityIdAndChanged(offset, limit, iCid,iEntityId,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndEntityIdAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndEntityIdAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iEntityId := self.Args("entity_id").MustInt()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndEntityIdAndStatus(offset, limit, iCid,iEntityId,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndEntityIdAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndEntityIdAndThreadHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iEntityId := self.Args("entity_id").MustInt()
	iThread := self.Args("thread").String()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndEntityIdAndThread(offset, limit, iCid,iEntityId,iThread)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndEntityIdAndThread's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndEntityIdAndEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iEntityId := self.Args("entity_id").MustInt()
	iEntityType := self.Args("entity_type").String()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndEntityIdAndEntityType(offset, limit, iCid,iEntityId,iEntityType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndEntityIdAndEntityType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndEntityIdAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iEntityId := self.Args("entity_id").MustInt()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndEntityIdAndFieldName(offset, limit, iCid,iEntityId,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndEntityIdAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndEntityIdAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iEntityId := self.Args("entity_id").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndEntityIdAndDefaultLangcode(offset, limit, iCid,iEntityId,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndEntityIdAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndSubjectAndUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iSubject := self.Args("subject").String()
	iUid := self.Args("uid").MustInt()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndSubjectAndUid(offset, limit, iCid,iSubject,iUid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndSubjectAndUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndSubjectAndNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iSubject := self.Args("subject").String()
	iName := self.Args("name").String()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndSubjectAndName(offset, limit, iCid,iSubject,iName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndSubjectAndName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndSubjectAndMailHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iSubject := self.Args("subject").String()
	iMail := self.Args("mail").String()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndSubjectAndMail(offset, limit, iCid,iSubject,iMail)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndSubjectAndMail's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndSubjectAndHomepageHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iSubject := self.Args("subject").String()
	iHomepage := self.Args("homepage").String()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndSubjectAndHomepage(offset, limit, iCid,iSubject,iHomepage)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndSubjectAndHomepage's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndSubjectAndHostnameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iSubject := self.Args("subject").String()
	iHostname := self.Args("hostname").String()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndSubjectAndHostname(offset, limit, iCid,iSubject,iHostname)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndSubjectAndHostname's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndSubjectAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iSubject := self.Args("subject").String()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndSubjectAndCreated(offset, limit, iCid,iSubject,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndSubjectAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndSubjectAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iSubject := self.Args("subject").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndSubjectAndChanged(offset, limit, iCid,iSubject,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndSubjectAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndSubjectAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iSubject := self.Args("subject").String()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndSubjectAndStatus(offset, limit, iCid,iSubject,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndSubjectAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndSubjectAndThreadHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iSubject := self.Args("subject").String()
	iThread := self.Args("thread").String()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndSubjectAndThread(offset, limit, iCid,iSubject,iThread)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndSubjectAndThread's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndSubjectAndEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iSubject := self.Args("subject").String()
	iEntityType := self.Args("entity_type").String()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndSubjectAndEntityType(offset, limit, iCid,iSubject,iEntityType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndSubjectAndEntityType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndSubjectAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iSubject := self.Args("subject").String()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndSubjectAndFieldName(offset, limit, iCid,iSubject,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndSubjectAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndSubjectAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iSubject := self.Args("subject").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndSubjectAndDefaultLangcode(offset, limit, iCid,iSubject,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndSubjectAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndUidAndNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iUid := self.Args("uid").MustInt()
	iName := self.Args("name").String()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndUidAndName(offset, limit, iCid,iUid,iName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndUidAndName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndUidAndMailHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iUid := self.Args("uid").MustInt()
	iMail := self.Args("mail").String()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndUidAndMail(offset, limit, iCid,iUid,iMail)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndUidAndMail's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndUidAndHomepageHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iUid := self.Args("uid").MustInt()
	iHomepage := self.Args("homepage").String()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndUidAndHomepage(offset, limit, iCid,iUid,iHomepage)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndUidAndHomepage's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndUidAndHostnameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iUid := self.Args("uid").MustInt()
	iHostname := self.Args("hostname").String()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndUidAndHostname(offset, limit, iCid,iUid,iHostname)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndUidAndHostname's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndUidAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iUid := self.Args("uid").MustInt()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndUidAndCreated(offset, limit, iCid,iUid,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndUidAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndUidAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iUid := self.Args("uid").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndUidAndChanged(offset, limit, iCid,iUid,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndUidAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndUidAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iUid := self.Args("uid").MustInt()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndUidAndStatus(offset, limit, iCid,iUid,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndUidAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndUidAndThreadHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iUid := self.Args("uid").MustInt()
	iThread := self.Args("thread").String()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndUidAndThread(offset, limit, iCid,iUid,iThread)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndUidAndThread's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndUidAndEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iUid := self.Args("uid").MustInt()
	iEntityType := self.Args("entity_type").String()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndUidAndEntityType(offset, limit, iCid,iUid,iEntityType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndUidAndEntityType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndUidAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iUid := self.Args("uid").MustInt()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndUidAndFieldName(offset, limit, iCid,iUid,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndUidAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndUidAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iUid := self.Args("uid").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndUidAndDefaultLangcode(offset, limit, iCid,iUid,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndUidAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndNameAndMailHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iName := self.Args("name").String()
	iMail := self.Args("mail").String()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndNameAndMail(offset, limit, iCid,iName,iMail)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndNameAndMail's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndNameAndHomepageHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iName := self.Args("name").String()
	iHomepage := self.Args("homepage").String()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndNameAndHomepage(offset, limit, iCid,iName,iHomepage)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndNameAndHomepage's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndNameAndHostnameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iName := self.Args("name").String()
	iHostname := self.Args("hostname").String()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndNameAndHostname(offset, limit, iCid,iName,iHostname)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndNameAndHostname's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndNameAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iName := self.Args("name").String()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndNameAndCreated(offset, limit, iCid,iName,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndNameAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndNameAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iName := self.Args("name").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndNameAndChanged(offset, limit, iCid,iName,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndNameAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndNameAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iName := self.Args("name").String()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndNameAndStatus(offset, limit, iCid,iName,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndNameAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndNameAndThreadHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iName := self.Args("name").String()
	iThread := self.Args("thread").String()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndNameAndThread(offset, limit, iCid,iName,iThread)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndNameAndThread's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndNameAndEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iName := self.Args("name").String()
	iEntityType := self.Args("entity_type").String()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndNameAndEntityType(offset, limit, iCid,iName,iEntityType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndNameAndEntityType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndNameAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iName := self.Args("name").String()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndNameAndFieldName(offset, limit, iCid,iName,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndNameAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndNameAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iName := self.Args("name").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndNameAndDefaultLangcode(offset, limit, iCid,iName,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndNameAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndMailAndHomepageHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iMail := self.Args("mail").String()
	iHomepage := self.Args("homepage").String()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndMailAndHomepage(offset, limit, iCid,iMail,iHomepage)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndMailAndHomepage's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndMailAndHostnameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iMail := self.Args("mail").String()
	iHostname := self.Args("hostname").String()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndMailAndHostname(offset, limit, iCid,iMail,iHostname)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndMailAndHostname's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndMailAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iMail := self.Args("mail").String()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndMailAndCreated(offset, limit, iCid,iMail,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndMailAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndMailAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iMail := self.Args("mail").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndMailAndChanged(offset, limit, iCid,iMail,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndMailAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndMailAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iMail := self.Args("mail").String()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndMailAndStatus(offset, limit, iCid,iMail,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndMailAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndMailAndThreadHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iMail := self.Args("mail").String()
	iThread := self.Args("thread").String()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndMailAndThread(offset, limit, iCid,iMail,iThread)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndMailAndThread's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndMailAndEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iMail := self.Args("mail").String()
	iEntityType := self.Args("entity_type").String()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndMailAndEntityType(offset, limit, iCid,iMail,iEntityType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndMailAndEntityType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndMailAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iMail := self.Args("mail").String()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndMailAndFieldName(offset, limit, iCid,iMail,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndMailAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndMailAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iMail := self.Args("mail").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndMailAndDefaultLangcode(offset, limit, iCid,iMail,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndMailAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndHomepageAndHostnameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iHomepage := self.Args("homepage").String()
	iHostname := self.Args("hostname").String()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndHomepageAndHostname(offset, limit, iCid,iHomepage,iHostname)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndHomepageAndHostname's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndHomepageAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iHomepage := self.Args("homepage").String()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndHomepageAndCreated(offset, limit, iCid,iHomepage,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndHomepageAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndHomepageAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iHomepage := self.Args("homepage").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndHomepageAndChanged(offset, limit, iCid,iHomepage,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndHomepageAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndHomepageAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iHomepage := self.Args("homepage").String()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndHomepageAndStatus(offset, limit, iCid,iHomepage,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndHomepageAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndHomepageAndThreadHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iHomepage := self.Args("homepage").String()
	iThread := self.Args("thread").String()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndHomepageAndThread(offset, limit, iCid,iHomepage,iThread)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndHomepageAndThread's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndHomepageAndEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iHomepage := self.Args("homepage").String()
	iEntityType := self.Args("entity_type").String()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndHomepageAndEntityType(offset, limit, iCid,iHomepage,iEntityType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndHomepageAndEntityType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndHomepageAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iHomepage := self.Args("homepage").String()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndHomepageAndFieldName(offset, limit, iCid,iHomepage,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndHomepageAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndHomepageAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iHomepage := self.Args("homepage").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndHomepageAndDefaultLangcode(offset, limit, iCid,iHomepage,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndHomepageAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndHostnameAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iHostname := self.Args("hostname").String()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndHostnameAndCreated(offset, limit, iCid,iHostname,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndHostnameAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndHostnameAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iHostname := self.Args("hostname").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndHostnameAndChanged(offset, limit, iCid,iHostname,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndHostnameAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndHostnameAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iHostname := self.Args("hostname").String()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndHostnameAndStatus(offset, limit, iCid,iHostname,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndHostnameAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndHostnameAndThreadHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iHostname := self.Args("hostname").String()
	iThread := self.Args("thread").String()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndHostnameAndThread(offset, limit, iCid,iHostname,iThread)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndHostnameAndThread's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndHostnameAndEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iHostname := self.Args("hostname").String()
	iEntityType := self.Args("entity_type").String()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndHostnameAndEntityType(offset, limit, iCid,iHostname,iEntityType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndHostnameAndEntityType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndHostnameAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iHostname := self.Args("hostname").String()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndHostnameAndFieldName(offset, limit, iCid,iHostname,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndHostnameAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndHostnameAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iHostname := self.Args("hostname").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndHostnameAndDefaultLangcode(offset, limit, iCid,iHostname,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndHostnameAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndCreatedAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iCreated := self.Args("created").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndCreatedAndChanged(offset, limit, iCid,iCreated,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndCreatedAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndCreatedAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iCreated := self.Args("created").MustInt()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndCreatedAndStatus(offset, limit, iCid,iCreated,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndCreatedAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndCreatedAndThreadHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iCreated := self.Args("created").MustInt()
	iThread := self.Args("thread").String()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndCreatedAndThread(offset, limit, iCid,iCreated,iThread)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndCreatedAndThread's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndCreatedAndEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iCreated := self.Args("created").MustInt()
	iEntityType := self.Args("entity_type").String()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndCreatedAndEntityType(offset, limit, iCid,iCreated,iEntityType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndCreatedAndEntityType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndCreatedAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iCreated := self.Args("created").MustInt()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndCreatedAndFieldName(offset, limit, iCid,iCreated,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndCreatedAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndCreatedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iCreated := self.Args("created").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndCreatedAndDefaultLangcode(offset, limit, iCid,iCreated,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndCreatedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndChangedAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iChanged := self.Args("changed").MustInt()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndChangedAndStatus(offset, limit, iCid,iChanged,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndChangedAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndChangedAndThreadHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iChanged := self.Args("changed").MustInt()
	iThread := self.Args("thread").String()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndChangedAndThread(offset, limit, iCid,iChanged,iThread)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndChangedAndThread's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndChangedAndEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iChanged := self.Args("changed").MustInt()
	iEntityType := self.Args("entity_type").String()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndChangedAndEntityType(offset, limit, iCid,iChanged,iEntityType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndChangedAndEntityType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndChangedAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iChanged := self.Args("changed").MustInt()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndChangedAndFieldName(offset, limit, iCid,iChanged,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndChangedAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndChangedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iChanged := self.Args("changed").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndChangedAndDefaultLangcode(offset, limit, iCid,iChanged,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndChangedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndStatusAndThreadHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iStatus := self.Args("status").MustInt()
	iThread := self.Args("thread").String()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndStatusAndThread(offset, limit, iCid,iStatus,iThread)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndStatusAndThread's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndStatusAndEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iStatus := self.Args("status").MustInt()
	iEntityType := self.Args("entity_type").String()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndStatusAndEntityType(offset, limit, iCid,iStatus,iEntityType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndStatusAndEntityType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndStatusAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iStatus := self.Args("status").MustInt()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndStatusAndFieldName(offset, limit, iCid,iStatus,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndStatusAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndStatusAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iStatus := self.Args("status").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndStatusAndDefaultLangcode(offset, limit, iCid,iStatus,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndStatusAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndThreadAndEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iThread := self.Args("thread").String()
	iEntityType := self.Args("entity_type").String()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndThreadAndEntityType(offset, limit, iCid,iThread,iEntityType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndThreadAndEntityType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndThreadAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iThread := self.Args("thread").String()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndThreadAndFieldName(offset, limit, iCid,iThread,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndThreadAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndThreadAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iThread := self.Args("thread").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndThreadAndDefaultLangcode(offset, limit, iCid,iThread,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndThreadAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndEntityTypeAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iEntityType := self.Args("entity_type").String()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndEntityTypeAndFieldName(offset, limit, iCid,iEntityType,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndEntityTypeAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndEntityTypeAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iEntityType := self.Args("entity_type").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndEntityTypeAndDefaultLangcode(offset, limit, iCid,iEntityType,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndEntityTypeAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndFieldNameAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iFieldName := self.Args("field_name").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndFieldNameAndDefaultLangcode(offset, limit, iCid,iFieldName,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndFieldNameAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndLangcodeAndPidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iLangcode := self.Args("langcode").String()
	iPid := self.Args("pid").MustInt()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndLangcodeAndPid(offset, limit, iCommentType,iLangcode,iPid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndLangcodeAndPid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndLangcodeAndEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iLangcode := self.Args("langcode").String()
	iEntityId := self.Args("entity_id").MustInt()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndLangcodeAndEntityId(offset, limit, iCommentType,iLangcode,iEntityId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndLangcodeAndEntityId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndLangcodeAndSubjectHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iLangcode := self.Args("langcode").String()
	iSubject := self.Args("subject").String()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndLangcodeAndSubject(offset, limit, iCommentType,iLangcode,iSubject)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndLangcodeAndSubject's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndLangcodeAndUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iLangcode := self.Args("langcode").String()
	iUid := self.Args("uid").MustInt()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndLangcodeAndUid(offset, limit, iCommentType,iLangcode,iUid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndLangcodeAndUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndLangcodeAndNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iLangcode := self.Args("langcode").String()
	iName := self.Args("name").String()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndLangcodeAndName(offset, limit, iCommentType,iLangcode,iName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndLangcodeAndName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndLangcodeAndMailHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iLangcode := self.Args("langcode").String()
	iMail := self.Args("mail").String()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndLangcodeAndMail(offset, limit, iCommentType,iLangcode,iMail)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndLangcodeAndMail's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndLangcodeAndHomepageHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iLangcode := self.Args("langcode").String()
	iHomepage := self.Args("homepage").String()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndLangcodeAndHomepage(offset, limit, iCommentType,iLangcode,iHomepage)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndLangcodeAndHomepage's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndLangcodeAndHostnameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iLangcode := self.Args("langcode").String()
	iHostname := self.Args("hostname").String()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndLangcodeAndHostname(offset, limit, iCommentType,iLangcode,iHostname)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndLangcodeAndHostname's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndLangcodeAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iLangcode := self.Args("langcode").String()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndLangcodeAndCreated(offset, limit, iCommentType,iLangcode,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndLangcodeAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndLangcodeAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iLangcode := self.Args("langcode").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndLangcodeAndChanged(offset, limit, iCommentType,iLangcode,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndLangcodeAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndLangcodeAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iLangcode := self.Args("langcode").String()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndLangcodeAndStatus(offset, limit, iCommentType,iLangcode,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndLangcodeAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndLangcodeAndThreadHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iLangcode := self.Args("langcode").String()
	iThread := self.Args("thread").String()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndLangcodeAndThread(offset, limit, iCommentType,iLangcode,iThread)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndLangcodeAndThread's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndLangcodeAndEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iLangcode := self.Args("langcode").String()
	iEntityType := self.Args("entity_type").String()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndLangcodeAndEntityType(offset, limit, iCommentType,iLangcode,iEntityType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndLangcodeAndEntityType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndLangcodeAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iLangcode := self.Args("langcode").String()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndLangcodeAndFieldName(offset, limit, iCommentType,iLangcode,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndLangcodeAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndLangcodeAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iLangcode := self.Args("langcode").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndLangcodeAndDefaultLangcode(offset, limit, iCommentType,iLangcode,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndLangcodeAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndPidAndEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iPid := self.Args("pid").MustInt()
	iEntityId := self.Args("entity_id").MustInt()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndPidAndEntityId(offset, limit, iCommentType,iPid,iEntityId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndPidAndEntityId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndPidAndSubjectHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iPid := self.Args("pid").MustInt()
	iSubject := self.Args("subject").String()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndPidAndSubject(offset, limit, iCommentType,iPid,iSubject)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndPidAndSubject's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndPidAndUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iPid := self.Args("pid").MustInt()
	iUid := self.Args("uid").MustInt()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndPidAndUid(offset, limit, iCommentType,iPid,iUid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndPidAndUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndPidAndNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iPid := self.Args("pid").MustInt()
	iName := self.Args("name").String()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndPidAndName(offset, limit, iCommentType,iPid,iName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndPidAndName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndPidAndMailHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iPid := self.Args("pid").MustInt()
	iMail := self.Args("mail").String()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndPidAndMail(offset, limit, iCommentType,iPid,iMail)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndPidAndMail's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndPidAndHomepageHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iPid := self.Args("pid").MustInt()
	iHomepage := self.Args("homepage").String()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndPidAndHomepage(offset, limit, iCommentType,iPid,iHomepage)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndPidAndHomepage's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndPidAndHostnameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iPid := self.Args("pid").MustInt()
	iHostname := self.Args("hostname").String()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndPidAndHostname(offset, limit, iCommentType,iPid,iHostname)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndPidAndHostname's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndPidAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iPid := self.Args("pid").MustInt()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndPidAndCreated(offset, limit, iCommentType,iPid,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndPidAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndPidAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iPid := self.Args("pid").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndPidAndChanged(offset, limit, iCommentType,iPid,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndPidAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndPidAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iPid := self.Args("pid").MustInt()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndPidAndStatus(offset, limit, iCommentType,iPid,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndPidAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndPidAndThreadHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iPid := self.Args("pid").MustInt()
	iThread := self.Args("thread").String()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndPidAndThread(offset, limit, iCommentType,iPid,iThread)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndPidAndThread's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndPidAndEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iPid := self.Args("pid").MustInt()
	iEntityType := self.Args("entity_type").String()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndPidAndEntityType(offset, limit, iCommentType,iPid,iEntityType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndPidAndEntityType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndPidAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iPid := self.Args("pid").MustInt()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndPidAndFieldName(offset, limit, iCommentType,iPid,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndPidAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndPidAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iPid := self.Args("pid").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndPidAndDefaultLangcode(offset, limit, iCommentType,iPid,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndPidAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndEntityIdAndSubjectHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iEntityId := self.Args("entity_id").MustInt()
	iSubject := self.Args("subject").String()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndEntityIdAndSubject(offset, limit, iCommentType,iEntityId,iSubject)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndEntityIdAndSubject's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndEntityIdAndUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iEntityId := self.Args("entity_id").MustInt()
	iUid := self.Args("uid").MustInt()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndEntityIdAndUid(offset, limit, iCommentType,iEntityId,iUid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndEntityIdAndUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndEntityIdAndNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iEntityId := self.Args("entity_id").MustInt()
	iName := self.Args("name").String()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndEntityIdAndName(offset, limit, iCommentType,iEntityId,iName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndEntityIdAndName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndEntityIdAndMailHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iEntityId := self.Args("entity_id").MustInt()
	iMail := self.Args("mail").String()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndEntityIdAndMail(offset, limit, iCommentType,iEntityId,iMail)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndEntityIdAndMail's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndEntityIdAndHomepageHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iEntityId := self.Args("entity_id").MustInt()
	iHomepage := self.Args("homepage").String()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndEntityIdAndHomepage(offset, limit, iCommentType,iEntityId,iHomepage)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndEntityIdAndHomepage's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndEntityIdAndHostnameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iEntityId := self.Args("entity_id").MustInt()
	iHostname := self.Args("hostname").String()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndEntityIdAndHostname(offset, limit, iCommentType,iEntityId,iHostname)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndEntityIdAndHostname's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndEntityIdAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iEntityId := self.Args("entity_id").MustInt()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndEntityIdAndCreated(offset, limit, iCommentType,iEntityId,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndEntityIdAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndEntityIdAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iEntityId := self.Args("entity_id").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndEntityIdAndChanged(offset, limit, iCommentType,iEntityId,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndEntityIdAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndEntityIdAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iEntityId := self.Args("entity_id").MustInt()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndEntityIdAndStatus(offset, limit, iCommentType,iEntityId,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndEntityIdAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndEntityIdAndThreadHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iEntityId := self.Args("entity_id").MustInt()
	iThread := self.Args("thread").String()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndEntityIdAndThread(offset, limit, iCommentType,iEntityId,iThread)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndEntityIdAndThread's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndEntityIdAndEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iEntityId := self.Args("entity_id").MustInt()
	iEntityType := self.Args("entity_type").String()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndEntityIdAndEntityType(offset, limit, iCommentType,iEntityId,iEntityType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndEntityIdAndEntityType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndEntityIdAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iEntityId := self.Args("entity_id").MustInt()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndEntityIdAndFieldName(offset, limit, iCommentType,iEntityId,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndEntityIdAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndEntityIdAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iEntityId := self.Args("entity_id").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndEntityIdAndDefaultLangcode(offset, limit, iCommentType,iEntityId,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndEntityIdAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndSubjectAndUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iSubject := self.Args("subject").String()
	iUid := self.Args("uid").MustInt()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndSubjectAndUid(offset, limit, iCommentType,iSubject,iUid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndSubjectAndUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndSubjectAndNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iSubject := self.Args("subject").String()
	iName := self.Args("name").String()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndSubjectAndName(offset, limit, iCommentType,iSubject,iName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndSubjectAndName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndSubjectAndMailHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iSubject := self.Args("subject").String()
	iMail := self.Args("mail").String()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndSubjectAndMail(offset, limit, iCommentType,iSubject,iMail)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndSubjectAndMail's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndSubjectAndHomepageHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iSubject := self.Args("subject").String()
	iHomepage := self.Args("homepage").String()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndSubjectAndHomepage(offset, limit, iCommentType,iSubject,iHomepage)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndSubjectAndHomepage's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndSubjectAndHostnameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iSubject := self.Args("subject").String()
	iHostname := self.Args("hostname").String()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndSubjectAndHostname(offset, limit, iCommentType,iSubject,iHostname)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndSubjectAndHostname's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndSubjectAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iSubject := self.Args("subject").String()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndSubjectAndCreated(offset, limit, iCommentType,iSubject,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndSubjectAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndSubjectAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iSubject := self.Args("subject").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndSubjectAndChanged(offset, limit, iCommentType,iSubject,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndSubjectAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndSubjectAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iSubject := self.Args("subject").String()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndSubjectAndStatus(offset, limit, iCommentType,iSubject,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndSubjectAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndSubjectAndThreadHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iSubject := self.Args("subject").String()
	iThread := self.Args("thread").String()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndSubjectAndThread(offset, limit, iCommentType,iSubject,iThread)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndSubjectAndThread's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndSubjectAndEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iSubject := self.Args("subject").String()
	iEntityType := self.Args("entity_type").String()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndSubjectAndEntityType(offset, limit, iCommentType,iSubject,iEntityType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndSubjectAndEntityType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndSubjectAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iSubject := self.Args("subject").String()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndSubjectAndFieldName(offset, limit, iCommentType,iSubject,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndSubjectAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndSubjectAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iSubject := self.Args("subject").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndSubjectAndDefaultLangcode(offset, limit, iCommentType,iSubject,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndSubjectAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndUidAndNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iUid := self.Args("uid").MustInt()
	iName := self.Args("name").String()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndUidAndName(offset, limit, iCommentType,iUid,iName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndUidAndName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndUidAndMailHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iUid := self.Args("uid").MustInt()
	iMail := self.Args("mail").String()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndUidAndMail(offset, limit, iCommentType,iUid,iMail)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndUidAndMail's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndUidAndHomepageHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iUid := self.Args("uid").MustInt()
	iHomepage := self.Args("homepage").String()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndUidAndHomepage(offset, limit, iCommentType,iUid,iHomepage)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndUidAndHomepage's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndUidAndHostnameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iUid := self.Args("uid").MustInt()
	iHostname := self.Args("hostname").String()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndUidAndHostname(offset, limit, iCommentType,iUid,iHostname)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndUidAndHostname's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndUidAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iUid := self.Args("uid").MustInt()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndUidAndCreated(offset, limit, iCommentType,iUid,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndUidAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndUidAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iUid := self.Args("uid").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndUidAndChanged(offset, limit, iCommentType,iUid,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndUidAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndUidAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iUid := self.Args("uid").MustInt()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndUidAndStatus(offset, limit, iCommentType,iUid,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndUidAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndUidAndThreadHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iUid := self.Args("uid").MustInt()
	iThread := self.Args("thread").String()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndUidAndThread(offset, limit, iCommentType,iUid,iThread)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndUidAndThread's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndUidAndEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iUid := self.Args("uid").MustInt()
	iEntityType := self.Args("entity_type").String()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndUidAndEntityType(offset, limit, iCommentType,iUid,iEntityType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndUidAndEntityType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndUidAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iUid := self.Args("uid").MustInt()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndUidAndFieldName(offset, limit, iCommentType,iUid,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndUidAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndUidAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iUid := self.Args("uid").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndUidAndDefaultLangcode(offset, limit, iCommentType,iUid,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndUidAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndNameAndMailHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iName := self.Args("name").String()
	iMail := self.Args("mail").String()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndNameAndMail(offset, limit, iCommentType,iName,iMail)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndNameAndMail's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndNameAndHomepageHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iName := self.Args("name").String()
	iHomepage := self.Args("homepage").String()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndNameAndHomepage(offset, limit, iCommentType,iName,iHomepage)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndNameAndHomepage's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndNameAndHostnameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iName := self.Args("name").String()
	iHostname := self.Args("hostname").String()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndNameAndHostname(offset, limit, iCommentType,iName,iHostname)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndNameAndHostname's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndNameAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iName := self.Args("name").String()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndNameAndCreated(offset, limit, iCommentType,iName,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndNameAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndNameAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iName := self.Args("name").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndNameAndChanged(offset, limit, iCommentType,iName,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndNameAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndNameAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iName := self.Args("name").String()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndNameAndStatus(offset, limit, iCommentType,iName,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndNameAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndNameAndThreadHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iName := self.Args("name").String()
	iThread := self.Args("thread").String()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndNameAndThread(offset, limit, iCommentType,iName,iThread)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndNameAndThread's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndNameAndEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iName := self.Args("name").String()
	iEntityType := self.Args("entity_type").String()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndNameAndEntityType(offset, limit, iCommentType,iName,iEntityType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndNameAndEntityType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndNameAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iName := self.Args("name").String()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndNameAndFieldName(offset, limit, iCommentType,iName,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndNameAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndNameAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iName := self.Args("name").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndNameAndDefaultLangcode(offset, limit, iCommentType,iName,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndNameAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndMailAndHomepageHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iMail := self.Args("mail").String()
	iHomepage := self.Args("homepage").String()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndMailAndHomepage(offset, limit, iCommentType,iMail,iHomepage)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndMailAndHomepage's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndMailAndHostnameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iMail := self.Args("mail").String()
	iHostname := self.Args("hostname").String()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndMailAndHostname(offset, limit, iCommentType,iMail,iHostname)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndMailAndHostname's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndMailAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iMail := self.Args("mail").String()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndMailAndCreated(offset, limit, iCommentType,iMail,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndMailAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndMailAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iMail := self.Args("mail").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndMailAndChanged(offset, limit, iCommentType,iMail,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndMailAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndMailAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iMail := self.Args("mail").String()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndMailAndStatus(offset, limit, iCommentType,iMail,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndMailAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndMailAndThreadHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iMail := self.Args("mail").String()
	iThread := self.Args("thread").String()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndMailAndThread(offset, limit, iCommentType,iMail,iThread)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndMailAndThread's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndMailAndEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iMail := self.Args("mail").String()
	iEntityType := self.Args("entity_type").String()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndMailAndEntityType(offset, limit, iCommentType,iMail,iEntityType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndMailAndEntityType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndMailAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iMail := self.Args("mail").String()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndMailAndFieldName(offset, limit, iCommentType,iMail,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndMailAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndMailAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iMail := self.Args("mail").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndMailAndDefaultLangcode(offset, limit, iCommentType,iMail,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndMailAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndHomepageAndHostnameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iHomepage := self.Args("homepage").String()
	iHostname := self.Args("hostname").String()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndHomepageAndHostname(offset, limit, iCommentType,iHomepage,iHostname)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndHomepageAndHostname's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndHomepageAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iHomepage := self.Args("homepage").String()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndHomepageAndCreated(offset, limit, iCommentType,iHomepage,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndHomepageAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndHomepageAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iHomepage := self.Args("homepage").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndHomepageAndChanged(offset, limit, iCommentType,iHomepage,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndHomepageAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndHomepageAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iHomepage := self.Args("homepage").String()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndHomepageAndStatus(offset, limit, iCommentType,iHomepage,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndHomepageAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndHomepageAndThreadHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iHomepage := self.Args("homepage").String()
	iThread := self.Args("thread").String()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndHomepageAndThread(offset, limit, iCommentType,iHomepage,iThread)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndHomepageAndThread's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndHomepageAndEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iHomepage := self.Args("homepage").String()
	iEntityType := self.Args("entity_type").String()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndHomepageAndEntityType(offset, limit, iCommentType,iHomepage,iEntityType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndHomepageAndEntityType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndHomepageAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iHomepage := self.Args("homepage").String()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndHomepageAndFieldName(offset, limit, iCommentType,iHomepage,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndHomepageAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndHomepageAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iHomepage := self.Args("homepage").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndHomepageAndDefaultLangcode(offset, limit, iCommentType,iHomepage,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndHomepageAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndHostnameAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iHostname := self.Args("hostname").String()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndHostnameAndCreated(offset, limit, iCommentType,iHostname,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndHostnameAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndHostnameAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iHostname := self.Args("hostname").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndHostnameAndChanged(offset, limit, iCommentType,iHostname,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndHostnameAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndHostnameAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iHostname := self.Args("hostname").String()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndHostnameAndStatus(offset, limit, iCommentType,iHostname,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndHostnameAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndHostnameAndThreadHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iHostname := self.Args("hostname").String()
	iThread := self.Args("thread").String()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndHostnameAndThread(offset, limit, iCommentType,iHostname,iThread)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndHostnameAndThread's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndHostnameAndEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iHostname := self.Args("hostname").String()
	iEntityType := self.Args("entity_type").String()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndHostnameAndEntityType(offset, limit, iCommentType,iHostname,iEntityType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndHostnameAndEntityType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndHostnameAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iHostname := self.Args("hostname").String()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndHostnameAndFieldName(offset, limit, iCommentType,iHostname,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndHostnameAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndHostnameAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iHostname := self.Args("hostname").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndHostnameAndDefaultLangcode(offset, limit, iCommentType,iHostname,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndHostnameAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndCreatedAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iCreated := self.Args("created").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndCreatedAndChanged(offset, limit, iCommentType,iCreated,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndCreatedAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndCreatedAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iCreated := self.Args("created").MustInt()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndCreatedAndStatus(offset, limit, iCommentType,iCreated,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndCreatedAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndCreatedAndThreadHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iCreated := self.Args("created").MustInt()
	iThread := self.Args("thread").String()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndCreatedAndThread(offset, limit, iCommentType,iCreated,iThread)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndCreatedAndThread's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndCreatedAndEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iCreated := self.Args("created").MustInt()
	iEntityType := self.Args("entity_type").String()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndCreatedAndEntityType(offset, limit, iCommentType,iCreated,iEntityType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndCreatedAndEntityType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndCreatedAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iCreated := self.Args("created").MustInt()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndCreatedAndFieldName(offset, limit, iCommentType,iCreated,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndCreatedAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndCreatedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iCreated := self.Args("created").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndCreatedAndDefaultLangcode(offset, limit, iCommentType,iCreated,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndCreatedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndChangedAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iChanged := self.Args("changed").MustInt()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndChangedAndStatus(offset, limit, iCommentType,iChanged,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndChangedAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndChangedAndThreadHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iChanged := self.Args("changed").MustInt()
	iThread := self.Args("thread").String()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndChangedAndThread(offset, limit, iCommentType,iChanged,iThread)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndChangedAndThread's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndChangedAndEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iChanged := self.Args("changed").MustInt()
	iEntityType := self.Args("entity_type").String()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndChangedAndEntityType(offset, limit, iCommentType,iChanged,iEntityType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndChangedAndEntityType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndChangedAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iChanged := self.Args("changed").MustInt()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndChangedAndFieldName(offset, limit, iCommentType,iChanged,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndChangedAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndChangedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iChanged := self.Args("changed").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndChangedAndDefaultLangcode(offset, limit, iCommentType,iChanged,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndChangedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndStatusAndThreadHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iStatus := self.Args("status").MustInt()
	iThread := self.Args("thread").String()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndStatusAndThread(offset, limit, iCommentType,iStatus,iThread)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndStatusAndThread's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndStatusAndEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iStatus := self.Args("status").MustInt()
	iEntityType := self.Args("entity_type").String()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndStatusAndEntityType(offset, limit, iCommentType,iStatus,iEntityType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndStatusAndEntityType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndStatusAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iStatus := self.Args("status").MustInt()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndStatusAndFieldName(offset, limit, iCommentType,iStatus,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndStatusAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndStatusAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iStatus := self.Args("status").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndStatusAndDefaultLangcode(offset, limit, iCommentType,iStatus,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndStatusAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndThreadAndEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iThread := self.Args("thread").String()
	iEntityType := self.Args("entity_type").String()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndThreadAndEntityType(offset, limit, iCommentType,iThread,iEntityType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndThreadAndEntityType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndThreadAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iThread := self.Args("thread").String()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndThreadAndFieldName(offset, limit, iCommentType,iThread,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndThreadAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndThreadAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iThread := self.Args("thread").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndThreadAndDefaultLangcode(offset, limit, iCommentType,iThread,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndThreadAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndEntityTypeAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iEntityType := self.Args("entity_type").String()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndEntityTypeAndFieldName(offset, limit, iCommentType,iEntityType,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndEntityTypeAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndEntityTypeAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iEntityType := self.Args("entity_type").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndEntityTypeAndDefaultLangcode(offset, limit, iCommentType,iEntityType,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndEntityTypeAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndFieldNameAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iFieldName := self.Args("field_name").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndFieldNameAndDefaultLangcode(offset, limit, iCommentType,iFieldName,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndFieldNameAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByLangcodeAndPidAndEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iPid := self.Args("pid").MustInt()
	iEntityId := self.Args("entity_id").MustInt()

	if helper.IsHas(iLangcode) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByLangcodeAndPidAndEntityId(offset, limit, iLangcode,iPid,iEntityId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByLangcodeAndPidAndEntityId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByLangcodeAndPidAndSubjectHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iPid := self.Args("pid").MustInt()
	iSubject := self.Args("subject").String()

	if helper.IsHas(iLangcode) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByLangcodeAndPidAndSubject(offset, limit, iLangcode,iPid,iSubject)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByLangcodeAndPidAndSubject's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByLangcodeAndPidAndUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iPid := self.Args("pid").MustInt()
	iUid := self.Args("uid").MustInt()

	if helper.IsHas(iLangcode) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByLangcodeAndPidAndUid(offset, limit, iLangcode,iPid,iUid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByLangcodeAndPidAndUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByLangcodeAndPidAndNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iPid := self.Args("pid").MustInt()
	iName := self.Args("name").String()

	if helper.IsHas(iLangcode) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByLangcodeAndPidAndName(offset, limit, iLangcode,iPid,iName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByLangcodeAndPidAndName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByLangcodeAndPidAndMailHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iPid := self.Args("pid").MustInt()
	iMail := self.Args("mail").String()

	if helper.IsHas(iLangcode) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByLangcodeAndPidAndMail(offset, limit, iLangcode,iPid,iMail)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByLangcodeAndPidAndMail's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByLangcodeAndPidAndHomepageHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iPid := self.Args("pid").MustInt()
	iHomepage := self.Args("homepage").String()

	if helper.IsHas(iLangcode) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByLangcodeAndPidAndHomepage(offset, limit, iLangcode,iPid,iHomepage)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByLangcodeAndPidAndHomepage's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByLangcodeAndPidAndHostnameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iPid := self.Args("pid").MustInt()
	iHostname := self.Args("hostname").String()

	if helper.IsHas(iLangcode) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByLangcodeAndPidAndHostname(offset, limit, iLangcode,iPid,iHostname)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByLangcodeAndPidAndHostname's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByLangcodeAndPidAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iPid := self.Args("pid").MustInt()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iLangcode) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByLangcodeAndPidAndCreated(offset, limit, iLangcode,iPid,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByLangcodeAndPidAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByLangcodeAndPidAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iPid := self.Args("pid").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iLangcode) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByLangcodeAndPidAndChanged(offset, limit, iLangcode,iPid,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByLangcodeAndPidAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByLangcodeAndPidAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iPid := self.Args("pid").MustInt()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iLangcode) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByLangcodeAndPidAndStatus(offset, limit, iLangcode,iPid,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByLangcodeAndPidAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByLangcodeAndPidAndThreadHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iPid := self.Args("pid").MustInt()
	iThread := self.Args("thread").String()

	if helper.IsHas(iLangcode) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByLangcodeAndPidAndThread(offset, limit, iLangcode,iPid,iThread)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByLangcodeAndPidAndThread's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByLangcodeAndPidAndEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iPid := self.Args("pid").MustInt()
	iEntityType := self.Args("entity_type").String()

	if helper.IsHas(iLangcode) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByLangcodeAndPidAndEntityType(offset, limit, iLangcode,iPid,iEntityType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByLangcodeAndPidAndEntityType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByLangcodeAndPidAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iPid := self.Args("pid").MustInt()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iLangcode) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByLangcodeAndPidAndFieldName(offset, limit, iLangcode,iPid,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByLangcodeAndPidAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByLangcodeAndPidAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iPid := self.Args("pid").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iLangcode) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByLangcodeAndPidAndDefaultLangcode(offset, limit, iLangcode,iPid,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByLangcodeAndPidAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByLangcodeAndEntityIdAndSubjectHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iEntityId := self.Args("entity_id").MustInt()
	iSubject := self.Args("subject").String()

	if helper.IsHas(iLangcode) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByLangcodeAndEntityIdAndSubject(offset, limit, iLangcode,iEntityId,iSubject)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByLangcodeAndEntityIdAndSubject's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByLangcodeAndEntityIdAndUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iEntityId := self.Args("entity_id").MustInt()
	iUid := self.Args("uid").MustInt()

	if helper.IsHas(iLangcode) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByLangcodeAndEntityIdAndUid(offset, limit, iLangcode,iEntityId,iUid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByLangcodeAndEntityIdAndUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByLangcodeAndEntityIdAndNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iEntityId := self.Args("entity_id").MustInt()
	iName := self.Args("name").String()

	if helper.IsHas(iLangcode) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByLangcodeAndEntityIdAndName(offset, limit, iLangcode,iEntityId,iName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByLangcodeAndEntityIdAndName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByLangcodeAndEntityIdAndMailHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iEntityId := self.Args("entity_id").MustInt()
	iMail := self.Args("mail").String()

	if helper.IsHas(iLangcode) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByLangcodeAndEntityIdAndMail(offset, limit, iLangcode,iEntityId,iMail)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByLangcodeAndEntityIdAndMail's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByLangcodeAndEntityIdAndHomepageHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iEntityId := self.Args("entity_id").MustInt()
	iHomepage := self.Args("homepage").String()

	if helper.IsHas(iLangcode) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByLangcodeAndEntityIdAndHomepage(offset, limit, iLangcode,iEntityId,iHomepage)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByLangcodeAndEntityIdAndHomepage's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByLangcodeAndEntityIdAndHostnameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iEntityId := self.Args("entity_id").MustInt()
	iHostname := self.Args("hostname").String()

	if helper.IsHas(iLangcode) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByLangcodeAndEntityIdAndHostname(offset, limit, iLangcode,iEntityId,iHostname)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByLangcodeAndEntityIdAndHostname's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByLangcodeAndEntityIdAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iEntityId := self.Args("entity_id").MustInt()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iLangcode) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByLangcodeAndEntityIdAndCreated(offset, limit, iLangcode,iEntityId,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByLangcodeAndEntityIdAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByLangcodeAndEntityIdAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iEntityId := self.Args("entity_id").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iLangcode) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByLangcodeAndEntityIdAndChanged(offset, limit, iLangcode,iEntityId,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByLangcodeAndEntityIdAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByLangcodeAndEntityIdAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iEntityId := self.Args("entity_id").MustInt()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iLangcode) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByLangcodeAndEntityIdAndStatus(offset, limit, iLangcode,iEntityId,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByLangcodeAndEntityIdAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByLangcodeAndEntityIdAndThreadHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iEntityId := self.Args("entity_id").MustInt()
	iThread := self.Args("thread").String()

	if helper.IsHas(iLangcode) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByLangcodeAndEntityIdAndThread(offset, limit, iLangcode,iEntityId,iThread)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByLangcodeAndEntityIdAndThread's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByLangcodeAndEntityIdAndEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iEntityId := self.Args("entity_id").MustInt()
	iEntityType := self.Args("entity_type").String()

	if helper.IsHas(iLangcode) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByLangcodeAndEntityIdAndEntityType(offset, limit, iLangcode,iEntityId,iEntityType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByLangcodeAndEntityIdAndEntityType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByLangcodeAndEntityIdAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iEntityId := self.Args("entity_id").MustInt()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iLangcode) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByLangcodeAndEntityIdAndFieldName(offset, limit, iLangcode,iEntityId,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByLangcodeAndEntityIdAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByLangcodeAndEntityIdAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iEntityId := self.Args("entity_id").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iLangcode) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByLangcodeAndEntityIdAndDefaultLangcode(offset, limit, iLangcode,iEntityId,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByLangcodeAndEntityIdAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByLangcodeAndSubjectAndUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iSubject := self.Args("subject").String()
	iUid := self.Args("uid").MustInt()

	if helper.IsHas(iLangcode) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByLangcodeAndSubjectAndUid(offset, limit, iLangcode,iSubject,iUid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByLangcodeAndSubjectAndUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByLangcodeAndSubjectAndNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iSubject := self.Args("subject").String()
	iName := self.Args("name").String()

	if helper.IsHas(iLangcode) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByLangcodeAndSubjectAndName(offset, limit, iLangcode,iSubject,iName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByLangcodeAndSubjectAndName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByLangcodeAndSubjectAndMailHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iSubject := self.Args("subject").String()
	iMail := self.Args("mail").String()

	if helper.IsHas(iLangcode) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByLangcodeAndSubjectAndMail(offset, limit, iLangcode,iSubject,iMail)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByLangcodeAndSubjectAndMail's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByLangcodeAndSubjectAndHomepageHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iSubject := self.Args("subject").String()
	iHomepage := self.Args("homepage").String()

	if helper.IsHas(iLangcode) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByLangcodeAndSubjectAndHomepage(offset, limit, iLangcode,iSubject,iHomepage)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByLangcodeAndSubjectAndHomepage's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByLangcodeAndSubjectAndHostnameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iSubject := self.Args("subject").String()
	iHostname := self.Args("hostname").String()

	if helper.IsHas(iLangcode) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByLangcodeAndSubjectAndHostname(offset, limit, iLangcode,iSubject,iHostname)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByLangcodeAndSubjectAndHostname's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByLangcodeAndSubjectAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iSubject := self.Args("subject").String()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iLangcode) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByLangcodeAndSubjectAndCreated(offset, limit, iLangcode,iSubject,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByLangcodeAndSubjectAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByLangcodeAndSubjectAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iSubject := self.Args("subject").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iLangcode) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByLangcodeAndSubjectAndChanged(offset, limit, iLangcode,iSubject,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByLangcodeAndSubjectAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByLangcodeAndSubjectAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iSubject := self.Args("subject").String()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iLangcode) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByLangcodeAndSubjectAndStatus(offset, limit, iLangcode,iSubject,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByLangcodeAndSubjectAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByLangcodeAndSubjectAndThreadHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iSubject := self.Args("subject").String()
	iThread := self.Args("thread").String()

	if helper.IsHas(iLangcode) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByLangcodeAndSubjectAndThread(offset, limit, iLangcode,iSubject,iThread)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByLangcodeAndSubjectAndThread's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByLangcodeAndSubjectAndEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iSubject := self.Args("subject").String()
	iEntityType := self.Args("entity_type").String()

	if helper.IsHas(iLangcode) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByLangcodeAndSubjectAndEntityType(offset, limit, iLangcode,iSubject,iEntityType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByLangcodeAndSubjectAndEntityType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByLangcodeAndSubjectAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iSubject := self.Args("subject").String()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iLangcode) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByLangcodeAndSubjectAndFieldName(offset, limit, iLangcode,iSubject,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByLangcodeAndSubjectAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByLangcodeAndSubjectAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iSubject := self.Args("subject").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iLangcode) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByLangcodeAndSubjectAndDefaultLangcode(offset, limit, iLangcode,iSubject,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByLangcodeAndSubjectAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByLangcodeAndUidAndNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iUid := self.Args("uid").MustInt()
	iName := self.Args("name").String()

	if helper.IsHas(iLangcode) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByLangcodeAndUidAndName(offset, limit, iLangcode,iUid,iName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByLangcodeAndUidAndName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByLangcodeAndUidAndMailHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iUid := self.Args("uid").MustInt()
	iMail := self.Args("mail").String()

	if helper.IsHas(iLangcode) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByLangcodeAndUidAndMail(offset, limit, iLangcode,iUid,iMail)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByLangcodeAndUidAndMail's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByLangcodeAndUidAndHomepageHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iUid := self.Args("uid").MustInt()
	iHomepage := self.Args("homepage").String()

	if helper.IsHas(iLangcode) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByLangcodeAndUidAndHomepage(offset, limit, iLangcode,iUid,iHomepage)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByLangcodeAndUidAndHomepage's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByLangcodeAndUidAndHostnameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iUid := self.Args("uid").MustInt()
	iHostname := self.Args("hostname").String()

	if helper.IsHas(iLangcode) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByLangcodeAndUidAndHostname(offset, limit, iLangcode,iUid,iHostname)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByLangcodeAndUidAndHostname's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByLangcodeAndUidAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iUid := self.Args("uid").MustInt()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iLangcode) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByLangcodeAndUidAndCreated(offset, limit, iLangcode,iUid,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByLangcodeAndUidAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByLangcodeAndUidAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iUid := self.Args("uid").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iLangcode) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByLangcodeAndUidAndChanged(offset, limit, iLangcode,iUid,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByLangcodeAndUidAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByLangcodeAndUidAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iUid := self.Args("uid").MustInt()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iLangcode) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByLangcodeAndUidAndStatus(offset, limit, iLangcode,iUid,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByLangcodeAndUidAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByLangcodeAndUidAndThreadHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iUid := self.Args("uid").MustInt()
	iThread := self.Args("thread").String()

	if helper.IsHas(iLangcode) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByLangcodeAndUidAndThread(offset, limit, iLangcode,iUid,iThread)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByLangcodeAndUidAndThread's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByLangcodeAndUidAndEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iUid := self.Args("uid").MustInt()
	iEntityType := self.Args("entity_type").String()

	if helper.IsHas(iLangcode) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByLangcodeAndUidAndEntityType(offset, limit, iLangcode,iUid,iEntityType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByLangcodeAndUidAndEntityType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByLangcodeAndUidAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iUid := self.Args("uid").MustInt()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iLangcode) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByLangcodeAndUidAndFieldName(offset, limit, iLangcode,iUid,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByLangcodeAndUidAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByLangcodeAndUidAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iUid := self.Args("uid").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iLangcode) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByLangcodeAndUidAndDefaultLangcode(offset, limit, iLangcode,iUid,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByLangcodeAndUidAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByLangcodeAndNameAndMailHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iName := self.Args("name").String()
	iMail := self.Args("mail").String()

	if helper.IsHas(iLangcode) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByLangcodeAndNameAndMail(offset, limit, iLangcode,iName,iMail)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByLangcodeAndNameAndMail's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByLangcodeAndNameAndHomepageHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iName := self.Args("name").String()
	iHomepage := self.Args("homepage").String()

	if helper.IsHas(iLangcode) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByLangcodeAndNameAndHomepage(offset, limit, iLangcode,iName,iHomepage)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByLangcodeAndNameAndHomepage's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByLangcodeAndNameAndHostnameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iName := self.Args("name").String()
	iHostname := self.Args("hostname").String()

	if helper.IsHas(iLangcode) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByLangcodeAndNameAndHostname(offset, limit, iLangcode,iName,iHostname)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByLangcodeAndNameAndHostname's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByLangcodeAndNameAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iName := self.Args("name").String()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iLangcode) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByLangcodeAndNameAndCreated(offset, limit, iLangcode,iName,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByLangcodeAndNameAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByLangcodeAndNameAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iName := self.Args("name").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iLangcode) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByLangcodeAndNameAndChanged(offset, limit, iLangcode,iName,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByLangcodeAndNameAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByLangcodeAndNameAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iName := self.Args("name").String()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iLangcode) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByLangcodeAndNameAndStatus(offset, limit, iLangcode,iName,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByLangcodeAndNameAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByLangcodeAndNameAndThreadHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iName := self.Args("name").String()
	iThread := self.Args("thread").String()

	if helper.IsHas(iLangcode) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByLangcodeAndNameAndThread(offset, limit, iLangcode,iName,iThread)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByLangcodeAndNameAndThread's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByLangcodeAndNameAndEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iName := self.Args("name").String()
	iEntityType := self.Args("entity_type").String()

	if helper.IsHas(iLangcode) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByLangcodeAndNameAndEntityType(offset, limit, iLangcode,iName,iEntityType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByLangcodeAndNameAndEntityType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByLangcodeAndNameAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iName := self.Args("name").String()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iLangcode) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByLangcodeAndNameAndFieldName(offset, limit, iLangcode,iName,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByLangcodeAndNameAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByLangcodeAndNameAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iName := self.Args("name").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iLangcode) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByLangcodeAndNameAndDefaultLangcode(offset, limit, iLangcode,iName,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByLangcodeAndNameAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByLangcodeAndMailAndHomepageHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iMail := self.Args("mail").String()
	iHomepage := self.Args("homepage").String()

	if helper.IsHas(iLangcode) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByLangcodeAndMailAndHomepage(offset, limit, iLangcode,iMail,iHomepage)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByLangcodeAndMailAndHomepage's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByLangcodeAndMailAndHostnameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iMail := self.Args("mail").String()
	iHostname := self.Args("hostname").String()

	if helper.IsHas(iLangcode) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByLangcodeAndMailAndHostname(offset, limit, iLangcode,iMail,iHostname)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByLangcodeAndMailAndHostname's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByLangcodeAndMailAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iMail := self.Args("mail").String()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iLangcode) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByLangcodeAndMailAndCreated(offset, limit, iLangcode,iMail,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByLangcodeAndMailAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByLangcodeAndMailAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iMail := self.Args("mail").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iLangcode) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByLangcodeAndMailAndChanged(offset, limit, iLangcode,iMail,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByLangcodeAndMailAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByLangcodeAndMailAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iMail := self.Args("mail").String()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iLangcode) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByLangcodeAndMailAndStatus(offset, limit, iLangcode,iMail,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByLangcodeAndMailAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByLangcodeAndMailAndThreadHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iMail := self.Args("mail").String()
	iThread := self.Args("thread").String()

	if helper.IsHas(iLangcode) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByLangcodeAndMailAndThread(offset, limit, iLangcode,iMail,iThread)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByLangcodeAndMailAndThread's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByLangcodeAndMailAndEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iMail := self.Args("mail").String()
	iEntityType := self.Args("entity_type").String()

	if helper.IsHas(iLangcode) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByLangcodeAndMailAndEntityType(offset, limit, iLangcode,iMail,iEntityType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByLangcodeAndMailAndEntityType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByLangcodeAndMailAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iMail := self.Args("mail").String()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iLangcode) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByLangcodeAndMailAndFieldName(offset, limit, iLangcode,iMail,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByLangcodeAndMailAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByLangcodeAndMailAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iMail := self.Args("mail").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iLangcode) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByLangcodeAndMailAndDefaultLangcode(offset, limit, iLangcode,iMail,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByLangcodeAndMailAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByLangcodeAndHomepageAndHostnameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iHomepage := self.Args("homepage").String()
	iHostname := self.Args("hostname").String()

	if helper.IsHas(iLangcode) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByLangcodeAndHomepageAndHostname(offset, limit, iLangcode,iHomepage,iHostname)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByLangcodeAndHomepageAndHostname's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByLangcodeAndHomepageAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iHomepage := self.Args("homepage").String()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iLangcode) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByLangcodeAndHomepageAndCreated(offset, limit, iLangcode,iHomepage,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByLangcodeAndHomepageAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByLangcodeAndHomepageAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iHomepage := self.Args("homepage").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iLangcode) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByLangcodeAndHomepageAndChanged(offset, limit, iLangcode,iHomepage,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByLangcodeAndHomepageAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByLangcodeAndHomepageAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iHomepage := self.Args("homepage").String()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iLangcode) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByLangcodeAndHomepageAndStatus(offset, limit, iLangcode,iHomepage,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByLangcodeAndHomepageAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByLangcodeAndHomepageAndThreadHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iHomepage := self.Args("homepage").String()
	iThread := self.Args("thread").String()

	if helper.IsHas(iLangcode) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByLangcodeAndHomepageAndThread(offset, limit, iLangcode,iHomepage,iThread)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByLangcodeAndHomepageAndThread's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByLangcodeAndHomepageAndEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iHomepage := self.Args("homepage").String()
	iEntityType := self.Args("entity_type").String()

	if helper.IsHas(iLangcode) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByLangcodeAndHomepageAndEntityType(offset, limit, iLangcode,iHomepage,iEntityType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByLangcodeAndHomepageAndEntityType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByLangcodeAndHomepageAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iHomepage := self.Args("homepage").String()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iLangcode) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByLangcodeAndHomepageAndFieldName(offset, limit, iLangcode,iHomepage,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByLangcodeAndHomepageAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByLangcodeAndHomepageAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iHomepage := self.Args("homepage").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iLangcode) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByLangcodeAndHomepageAndDefaultLangcode(offset, limit, iLangcode,iHomepage,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByLangcodeAndHomepageAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByLangcodeAndHostnameAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iHostname := self.Args("hostname").String()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iLangcode) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByLangcodeAndHostnameAndCreated(offset, limit, iLangcode,iHostname,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByLangcodeAndHostnameAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByLangcodeAndHostnameAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iHostname := self.Args("hostname").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iLangcode) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByLangcodeAndHostnameAndChanged(offset, limit, iLangcode,iHostname,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByLangcodeAndHostnameAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByLangcodeAndHostnameAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iHostname := self.Args("hostname").String()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iLangcode) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByLangcodeAndHostnameAndStatus(offset, limit, iLangcode,iHostname,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByLangcodeAndHostnameAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByLangcodeAndHostnameAndThreadHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iHostname := self.Args("hostname").String()
	iThread := self.Args("thread").String()

	if helper.IsHas(iLangcode) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByLangcodeAndHostnameAndThread(offset, limit, iLangcode,iHostname,iThread)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByLangcodeAndHostnameAndThread's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByLangcodeAndHostnameAndEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iHostname := self.Args("hostname").String()
	iEntityType := self.Args("entity_type").String()

	if helper.IsHas(iLangcode) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByLangcodeAndHostnameAndEntityType(offset, limit, iLangcode,iHostname,iEntityType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByLangcodeAndHostnameAndEntityType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByLangcodeAndHostnameAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iHostname := self.Args("hostname").String()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iLangcode) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByLangcodeAndHostnameAndFieldName(offset, limit, iLangcode,iHostname,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByLangcodeAndHostnameAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByLangcodeAndHostnameAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iHostname := self.Args("hostname").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iLangcode) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByLangcodeAndHostnameAndDefaultLangcode(offset, limit, iLangcode,iHostname,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByLangcodeAndHostnameAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByLangcodeAndCreatedAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iCreated := self.Args("created").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iLangcode) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByLangcodeAndCreatedAndChanged(offset, limit, iLangcode,iCreated,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByLangcodeAndCreatedAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByLangcodeAndCreatedAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iCreated := self.Args("created").MustInt()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iLangcode) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByLangcodeAndCreatedAndStatus(offset, limit, iLangcode,iCreated,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByLangcodeAndCreatedAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByLangcodeAndCreatedAndThreadHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iCreated := self.Args("created").MustInt()
	iThread := self.Args("thread").String()

	if helper.IsHas(iLangcode) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByLangcodeAndCreatedAndThread(offset, limit, iLangcode,iCreated,iThread)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByLangcodeAndCreatedAndThread's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByLangcodeAndCreatedAndEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iCreated := self.Args("created").MustInt()
	iEntityType := self.Args("entity_type").String()

	if helper.IsHas(iLangcode) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByLangcodeAndCreatedAndEntityType(offset, limit, iLangcode,iCreated,iEntityType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByLangcodeAndCreatedAndEntityType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByLangcodeAndCreatedAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iCreated := self.Args("created").MustInt()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iLangcode) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByLangcodeAndCreatedAndFieldName(offset, limit, iLangcode,iCreated,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByLangcodeAndCreatedAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByLangcodeAndCreatedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iCreated := self.Args("created").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iLangcode) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByLangcodeAndCreatedAndDefaultLangcode(offset, limit, iLangcode,iCreated,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByLangcodeAndCreatedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByLangcodeAndChangedAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iChanged := self.Args("changed").MustInt()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iLangcode) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByLangcodeAndChangedAndStatus(offset, limit, iLangcode,iChanged,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByLangcodeAndChangedAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByLangcodeAndChangedAndThreadHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iChanged := self.Args("changed").MustInt()
	iThread := self.Args("thread").String()

	if helper.IsHas(iLangcode) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByLangcodeAndChangedAndThread(offset, limit, iLangcode,iChanged,iThread)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByLangcodeAndChangedAndThread's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByLangcodeAndChangedAndEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iChanged := self.Args("changed").MustInt()
	iEntityType := self.Args("entity_type").String()

	if helper.IsHas(iLangcode) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByLangcodeAndChangedAndEntityType(offset, limit, iLangcode,iChanged,iEntityType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByLangcodeAndChangedAndEntityType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByLangcodeAndChangedAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iChanged := self.Args("changed").MustInt()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iLangcode) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByLangcodeAndChangedAndFieldName(offset, limit, iLangcode,iChanged,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByLangcodeAndChangedAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByLangcodeAndChangedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iChanged := self.Args("changed").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iLangcode) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByLangcodeAndChangedAndDefaultLangcode(offset, limit, iLangcode,iChanged,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByLangcodeAndChangedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByLangcodeAndStatusAndThreadHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iStatus := self.Args("status").MustInt()
	iThread := self.Args("thread").String()

	if helper.IsHas(iLangcode) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByLangcodeAndStatusAndThread(offset, limit, iLangcode,iStatus,iThread)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByLangcodeAndStatusAndThread's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByLangcodeAndStatusAndEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iStatus := self.Args("status").MustInt()
	iEntityType := self.Args("entity_type").String()

	if helper.IsHas(iLangcode) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByLangcodeAndStatusAndEntityType(offset, limit, iLangcode,iStatus,iEntityType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByLangcodeAndStatusAndEntityType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByLangcodeAndStatusAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iStatus := self.Args("status").MustInt()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iLangcode) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByLangcodeAndStatusAndFieldName(offset, limit, iLangcode,iStatus,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByLangcodeAndStatusAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByLangcodeAndStatusAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iStatus := self.Args("status").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iLangcode) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByLangcodeAndStatusAndDefaultLangcode(offset, limit, iLangcode,iStatus,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByLangcodeAndStatusAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByLangcodeAndThreadAndEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iThread := self.Args("thread").String()
	iEntityType := self.Args("entity_type").String()

	if helper.IsHas(iLangcode) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByLangcodeAndThreadAndEntityType(offset, limit, iLangcode,iThread,iEntityType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByLangcodeAndThreadAndEntityType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByLangcodeAndThreadAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iThread := self.Args("thread").String()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iLangcode) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByLangcodeAndThreadAndFieldName(offset, limit, iLangcode,iThread,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByLangcodeAndThreadAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByLangcodeAndThreadAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iThread := self.Args("thread").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iLangcode) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByLangcodeAndThreadAndDefaultLangcode(offset, limit, iLangcode,iThread,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByLangcodeAndThreadAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByLangcodeAndEntityTypeAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iEntityType := self.Args("entity_type").String()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iLangcode) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByLangcodeAndEntityTypeAndFieldName(offset, limit, iLangcode,iEntityType,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByLangcodeAndEntityTypeAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByLangcodeAndEntityTypeAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iEntityType := self.Args("entity_type").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iLangcode) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByLangcodeAndEntityTypeAndDefaultLangcode(offset, limit, iLangcode,iEntityType,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByLangcodeAndEntityTypeAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByLangcodeAndFieldNameAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iFieldName := self.Args("field_name").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iLangcode) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByLangcodeAndFieldNameAndDefaultLangcode(offset, limit, iLangcode,iFieldName,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByLangcodeAndFieldNameAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByPidAndEntityIdAndSubjectHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPid := self.Args("pid").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iSubject := self.Args("subject").String()

	if helper.IsHas(iPid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByPidAndEntityIdAndSubject(offset, limit, iPid,iEntityId,iSubject)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByPidAndEntityIdAndSubject's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByPidAndEntityIdAndUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPid := self.Args("pid").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iUid := self.Args("uid").MustInt()

	if helper.IsHas(iPid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByPidAndEntityIdAndUid(offset, limit, iPid,iEntityId,iUid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByPidAndEntityIdAndUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByPidAndEntityIdAndNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPid := self.Args("pid").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iName := self.Args("name").String()

	if helper.IsHas(iPid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByPidAndEntityIdAndName(offset, limit, iPid,iEntityId,iName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByPidAndEntityIdAndName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByPidAndEntityIdAndMailHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPid := self.Args("pid").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iMail := self.Args("mail").String()

	if helper.IsHas(iPid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByPidAndEntityIdAndMail(offset, limit, iPid,iEntityId,iMail)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByPidAndEntityIdAndMail's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByPidAndEntityIdAndHomepageHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPid := self.Args("pid").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iHomepage := self.Args("homepage").String()

	if helper.IsHas(iPid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByPidAndEntityIdAndHomepage(offset, limit, iPid,iEntityId,iHomepage)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByPidAndEntityIdAndHomepage's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByPidAndEntityIdAndHostnameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPid := self.Args("pid").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iHostname := self.Args("hostname").String()

	if helper.IsHas(iPid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByPidAndEntityIdAndHostname(offset, limit, iPid,iEntityId,iHostname)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByPidAndEntityIdAndHostname's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByPidAndEntityIdAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPid := self.Args("pid").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iPid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByPidAndEntityIdAndCreated(offset, limit, iPid,iEntityId,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByPidAndEntityIdAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByPidAndEntityIdAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPid := self.Args("pid").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iPid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByPidAndEntityIdAndChanged(offset, limit, iPid,iEntityId,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByPidAndEntityIdAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByPidAndEntityIdAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPid := self.Args("pid").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iPid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByPidAndEntityIdAndStatus(offset, limit, iPid,iEntityId,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByPidAndEntityIdAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByPidAndEntityIdAndThreadHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPid := self.Args("pid").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iThread := self.Args("thread").String()

	if helper.IsHas(iPid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByPidAndEntityIdAndThread(offset, limit, iPid,iEntityId,iThread)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByPidAndEntityIdAndThread's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByPidAndEntityIdAndEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPid := self.Args("pid").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iEntityType := self.Args("entity_type").String()

	if helper.IsHas(iPid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByPidAndEntityIdAndEntityType(offset, limit, iPid,iEntityId,iEntityType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByPidAndEntityIdAndEntityType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByPidAndEntityIdAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPid := self.Args("pid").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iPid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByPidAndEntityIdAndFieldName(offset, limit, iPid,iEntityId,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByPidAndEntityIdAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByPidAndEntityIdAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPid := self.Args("pid").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iPid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByPidAndEntityIdAndDefaultLangcode(offset, limit, iPid,iEntityId,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByPidAndEntityIdAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByPidAndSubjectAndUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPid := self.Args("pid").MustInt()
	iSubject := self.Args("subject").String()
	iUid := self.Args("uid").MustInt()

	if helper.IsHas(iPid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByPidAndSubjectAndUid(offset, limit, iPid,iSubject,iUid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByPidAndSubjectAndUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByPidAndSubjectAndNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPid := self.Args("pid").MustInt()
	iSubject := self.Args("subject").String()
	iName := self.Args("name").String()

	if helper.IsHas(iPid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByPidAndSubjectAndName(offset, limit, iPid,iSubject,iName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByPidAndSubjectAndName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByPidAndSubjectAndMailHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPid := self.Args("pid").MustInt()
	iSubject := self.Args("subject").String()
	iMail := self.Args("mail").String()

	if helper.IsHas(iPid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByPidAndSubjectAndMail(offset, limit, iPid,iSubject,iMail)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByPidAndSubjectAndMail's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByPidAndSubjectAndHomepageHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPid := self.Args("pid").MustInt()
	iSubject := self.Args("subject").String()
	iHomepage := self.Args("homepage").String()

	if helper.IsHas(iPid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByPidAndSubjectAndHomepage(offset, limit, iPid,iSubject,iHomepage)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByPidAndSubjectAndHomepage's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByPidAndSubjectAndHostnameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPid := self.Args("pid").MustInt()
	iSubject := self.Args("subject").String()
	iHostname := self.Args("hostname").String()

	if helper.IsHas(iPid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByPidAndSubjectAndHostname(offset, limit, iPid,iSubject,iHostname)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByPidAndSubjectAndHostname's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByPidAndSubjectAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPid := self.Args("pid").MustInt()
	iSubject := self.Args("subject").String()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iPid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByPidAndSubjectAndCreated(offset, limit, iPid,iSubject,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByPidAndSubjectAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByPidAndSubjectAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPid := self.Args("pid").MustInt()
	iSubject := self.Args("subject").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iPid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByPidAndSubjectAndChanged(offset, limit, iPid,iSubject,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByPidAndSubjectAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByPidAndSubjectAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPid := self.Args("pid").MustInt()
	iSubject := self.Args("subject").String()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iPid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByPidAndSubjectAndStatus(offset, limit, iPid,iSubject,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByPidAndSubjectAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByPidAndSubjectAndThreadHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPid := self.Args("pid").MustInt()
	iSubject := self.Args("subject").String()
	iThread := self.Args("thread").String()

	if helper.IsHas(iPid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByPidAndSubjectAndThread(offset, limit, iPid,iSubject,iThread)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByPidAndSubjectAndThread's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByPidAndSubjectAndEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPid := self.Args("pid").MustInt()
	iSubject := self.Args("subject").String()
	iEntityType := self.Args("entity_type").String()

	if helper.IsHas(iPid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByPidAndSubjectAndEntityType(offset, limit, iPid,iSubject,iEntityType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByPidAndSubjectAndEntityType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByPidAndSubjectAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPid := self.Args("pid").MustInt()
	iSubject := self.Args("subject").String()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iPid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByPidAndSubjectAndFieldName(offset, limit, iPid,iSubject,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByPidAndSubjectAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByPidAndSubjectAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPid := self.Args("pid").MustInt()
	iSubject := self.Args("subject").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iPid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByPidAndSubjectAndDefaultLangcode(offset, limit, iPid,iSubject,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByPidAndSubjectAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByPidAndUidAndNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPid := self.Args("pid").MustInt()
	iUid := self.Args("uid").MustInt()
	iName := self.Args("name").String()

	if helper.IsHas(iPid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByPidAndUidAndName(offset, limit, iPid,iUid,iName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByPidAndUidAndName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByPidAndUidAndMailHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPid := self.Args("pid").MustInt()
	iUid := self.Args("uid").MustInt()
	iMail := self.Args("mail").String()

	if helper.IsHas(iPid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByPidAndUidAndMail(offset, limit, iPid,iUid,iMail)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByPidAndUidAndMail's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByPidAndUidAndHomepageHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPid := self.Args("pid").MustInt()
	iUid := self.Args("uid").MustInt()
	iHomepage := self.Args("homepage").String()

	if helper.IsHas(iPid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByPidAndUidAndHomepage(offset, limit, iPid,iUid,iHomepage)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByPidAndUidAndHomepage's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByPidAndUidAndHostnameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPid := self.Args("pid").MustInt()
	iUid := self.Args("uid").MustInt()
	iHostname := self.Args("hostname").String()

	if helper.IsHas(iPid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByPidAndUidAndHostname(offset, limit, iPid,iUid,iHostname)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByPidAndUidAndHostname's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByPidAndUidAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPid := self.Args("pid").MustInt()
	iUid := self.Args("uid").MustInt()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iPid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByPidAndUidAndCreated(offset, limit, iPid,iUid,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByPidAndUidAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByPidAndUidAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPid := self.Args("pid").MustInt()
	iUid := self.Args("uid").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iPid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByPidAndUidAndChanged(offset, limit, iPid,iUid,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByPidAndUidAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByPidAndUidAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPid := self.Args("pid").MustInt()
	iUid := self.Args("uid").MustInt()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iPid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByPidAndUidAndStatus(offset, limit, iPid,iUid,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByPidAndUidAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByPidAndUidAndThreadHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPid := self.Args("pid").MustInt()
	iUid := self.Args("uid").MustInt()
	iThread := self.Args("thread").String()

	if helper.IsHas(iPid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByPidAndUidAndThread(offset, limit, iPid,iUid,iThread)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByPidAndUidAndThread's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByPidAndUidAndEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPid := self.Args("pid").MustInt()
	iUid := self.Args("uid").MustInt()
	iEntityType := self.Args("entity_type").String()

	if helper.IsHas(iPid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByPidAndUidAndEntityType(offset, limit, iPid,iUid,iEntityType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByPidAndUidAndEntityType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByPidAndUidAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPid := self.Args("pid").MustInt()
	iUid := self.Args("uid").MustInt()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iPid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByPidAndUidAndFieldName(offset, limit, iPid,iUid,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByPidAndUidAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByPidAndUidAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPid := self.Args("pid").MustInt()
	iUid := self.Args("uid").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iPid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByPidAndUidAndDefaultLangcode(offset, limit, iPid,iUid,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByPidAndUidAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByPidAndNameAndMailHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPid := self.Args("pid").MustInt()
	iName := self.Args("name").String()
	iMail := self.Args("mail").String()

	if helper.IsHas(iPid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByPidAndNameAndMail(offset, limit, iPid,iName,iMail)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByPidAndNameAndMail's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByPidAndNameAndHomepageHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPid := self.Args("pid").MustInt()
	iName := self.Args("name").String()
	iHomepage := self.Args("homepage").String()

	if helper.IsHas(iPid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByPidAndNameAndHomepage(offset, limit, iPid,iName,iHomepage)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByPidAndNameAndHomepage's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByPidAndNameAndHostnameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPid := self.Args("pid").MustInt()
	iName := self.Args("name").String()
	iHostname := self.Args("hostname").String()

	if helper.IsHas(iPid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByPidAndNameAndHostname(offset, limit, iPid,iName,iHostname)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByPidAndNameAndHostname's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByPidAndNameAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPid := self.Args("pid").MustInt()
	iName := self.Args("name").String()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iPid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByPidAndNameAndCreated(offset, limit, iPid,iName,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByPidAndNameAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByPidAndNameAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPid := self.Args("pid").MustInt()
	iName := self.Args("name").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iPid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByPidAndNameAndChanged(offset, limit, iPid,iName,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByPidAndNameAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByPidAndNameAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPid := self.Args("pid").MustInt()
	iName := self.Args("name").String()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iPid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByPidAndNameAndStatus(offset, limit, iPid,iName,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByPidAndNameAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByPidAndNameAndThreadHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPid := self.Args("pid").MustInt()
	iName := self.Args("name").String()
	iThread := self.Args("thread").String()

	if helper.IsHas(iPid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByPidAndNameAndThread(offset, limit, iPid,iName,iThread)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByPidAndNameAndThread's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByPidAndNameAndEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPid := self.Args("pid").MustInt()
	iName := self.Args("name").String()
	iEntityType := self.Args("entity_type").String()

	if helper.IsHas(iPid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByPidAndNameAndEntityType(offset, limit, iPid,iName,iEntityType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByPidAndNameAndEntityType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByPidAndNameAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPid := self.Args("pid").MustInt()
	iName := self.Args("name").String()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iPid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByPidAndNameAndFieldName(offset, limit, iPid,iName,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByPidAndNameAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByPidAndNameAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPid := self.Args("pid").MustInt()
	iName := self.Args("name").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iPid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByPidAndNameAndDefaultLangcode(offset, limit, iPid,iName,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByPidAndNameAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByPidAndMailAndHomepageHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPid := self.Args("pid").MustInt()
	iMail := self.Args("mail").String()
	iHomepage := self.Args("homepage").String()

	if helper.IsHas(iPid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByPidAndMailAndHomepage(offset, limit, iPid,iMail,iHomepage)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByPidAndMailAndHomepage's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByPidAndMailAndHostnameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPid := self.Args("pid").MustInt()
	iMail := self.Args("mail").String()
	iHostname := self.Args("hostname").String()

	if helper.IsHas(iPid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByPidAndMailAndHostname(offset, limit, iPid,iMail,iHostname)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByPidAndMailAndHostname's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByPidAndMailAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPid := self.Args("pid").MustInt()
	iMail := self.Args("mail").String()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iPid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByPidAndMailAndCreated(offset, limit, iPid,iMail,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByPidAndMailAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByPidAndMailAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPid := self.Args("pid").MustInt()
	iMail := self.Args("mail").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iPid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByPidAndMailAndChanged(offset, limit, iPid,iMail,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByPidAndMailAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByPidAndMailAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPid := self.Args("pid").MustInt()
	iMail := self.Args("mail").String()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iPid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByPidAndMailAndStatus(offset, limit, iPid,iMail,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByPidAndMailAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByPidAndMailAndThreadHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPid := self.Args("pid").MustInt()
	iMail := self.Args("mail").String()
	iThread := self.Args("thread").String()

	if helper.IsHas(iPid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByPidAndMailAndThread(offset, limit, iPid,iMail,iThread)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByPidAndMailAndThread's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByPidAndMailAndEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPid := self.Args("pid").MustInt()
	iMail := self.Args("mail").String()
	iEntityType := self.Args("entity_type").String()

	if helper.IsHas(iPid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByPidAndMailAndEntityType(offset, limit, iPid,iMail,iEntityType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByPidAndMailAndEntityType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByPidAndMailAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPid := self.Args("pid").MustInt()
	iMail := self.Args("mail").String()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iPid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByPidAndMailAndFieldName(offset, limit, iPid,iMail,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByPidAndMailAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByPidAndMailAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPid := self.Args("pid").MustInt()
	iMail := self.Args("mail").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iPid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByPidAndMailAndDefaultLangcode(offset, limit, iPid,iMail,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByPidAndMailAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByPidAndHomepageAndHostnameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPid := self.Args("pid").MustInt()
	iHomepage := self.Args("homepage").String()
	iHostname := self.Args("hostname").String()

	if helper.IsHas(iPid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByPidAndHomepageAndHostname(offset, limit, iPid,iHomepage,iHostname)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByPidAndHomepageAndHostname's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByPidAndHomepageAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPid := self.Args("pid").MustInt()
	iHomepage := self.Args("homepage").String()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iPid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByPidAndHomepageAndCreated(offset, limit, iPid,iHomepage,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByPidAndHomepageAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByPidAndHomepageAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPid := self.Args("pid").MustInt()
	iHomepage := self.Args("homepage").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iPid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByPidAndHomepageAndChanged(offset, limit, iPid,iHomepage,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByPidAndHomepageAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByPidAndHomepageAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPid := self.Args("pid").MustInt()
	iHomepage := self.Args("homepage").String()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iPid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByPidAndHomepageAndStatus(offset, limit, iPid,iHomepage,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByPidAndHomepageAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByPidAndHomepageAndThreadHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPid := self.Args("pid").MustInt()
	iHomepage := self.Args("homepage").String()
	iThread := self.Args("thread").String()

	if helper.IsHas(iPid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByPidAndHomepageAndThread(offset, limit, iPid,iHomepage,iThread)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByPidAndHomepageAndThread's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByPidAndHomepageAndEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPid := self.Args("pid").MustInt()
	iHomepage := self.Args("homepage").String()
	iEntityType := self.Args("entity_type").String()

	if helper.IsHas(iPid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByPidAndHomepageAndEntityType(offset, limit, iPid,iHomepage,iEntityType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByPidAndHomepageAndEntityType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByPidAndHomepageAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPid := self.Args("pid").MustInt()
	iHomepage := self.Args("homepage").String()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iPid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByPidAndHomepageAndFieldName(offset, limit, iPid,iHomepage,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByPidAndHomepageAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByPidAndHomepageAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPid := self.Args("pid").MustInt()
	iHomepage := self.Args("homepage").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iPid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByPidAndHomepageAndDefaultLangcode(offset, limit, iPid,iHomepage,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByPidAndHomepageAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByPidAndHostnameAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPid := self.Args("pid").MustInt()
	iHostname := self.Args("hostname").String()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iPid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByPidAndHostnameAndCreated(offset, limit, iPid,iHostname,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByPidAndHostnameAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByPidAndHostnameAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPid := self.Args("pid").MustInt()
	iHostname := self.Args("hostname").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iPid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByPidAndHostnameAndChanged(offset, limit, iPid,iHostname,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByPidAndHostnameAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByPidAndHostnameAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPid := self.Args("pid").MustInt()
	iHostname := self.Args("hostname").String()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iPid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByPidAndHostnameAndStatus(offset, limit, iPid,iHostname,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByPidAndHostnameAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByPidAndHostnameAndThreadHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPid := self.Args("pid").MustInt()
	iHostname := self.Args("hostname").String()
	iThread := self.Args("thread").String()

	if helper.IsHas(iPid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByPidAndHostnameAndThread(offset, limit, iPid,iHostname,iThread)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByPidAndHostnameAndThread's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByPidAndHostnameAndEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPid := self.Args("pid").MustInt()
	iHostname := self.Args("hostname").String()
	iEntityType := self.Args("entity_type").String()

	if helper.IsHas(iPid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByPidAndHostnameAndEntityType(offset, limit, iPid,iHostname,iEntityType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByPidAndHostnameAndEntityType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByPidAndHostnameAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPid := self.Args("pid").MustInt()
	iHostname := self.Args("hostname").String()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iPid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByPidAndHostnameAndFieldName(offset, limit, iPid,iHostname,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByPidAndHostnameAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByPidAndHostnameAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPid := self.Args("pid").MustInt()
	iHostname := self.Args("hostname").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iPid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByPidAndHostnameAndDefaultLangcode(offset, limit, iPid,iHostname,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByPidAndHostnameAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByPidAndCreatedAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPid := self.Args("pid").MustInt()
	iCreated := self.Args("created").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iPid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByPidAndCreatedAndChanged(offset, limit, iPid,iCreated,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByPidAndCreatedAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByPidAndCreatedAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPid := self.Args("pid").MustInt()
	iCreated := self.Args("created").MustInt()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iPid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByPidAndCreatedAndStatus(offset, limit, iPid,iCreated,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByPidAndCreatedAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByPidAndCreatedAndThreadHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPid := self.Args("pid").MustInt()
	iCreated := self.Args("created").MustInt()
	iThread := self.Args("thread").String()

	if helper.IsHas(iPid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByPidAndCreatedAndThread(offset, limit, iPid,iCreated,iThread)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByPidAndCreatedAndThread's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByPidAndCreatedAndEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPid := self.Args("pid").MustInt()
	iCreated := self.Args("created").MustInt()
	iEntityType := self.Args("entity_type").String()

	if helper.IsHas(iPid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByPidAndCreatedAndEntityType(offset, limit, iPid,iCreated,iEntityType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByPidAndCreatedAndEntityType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByPidAndCreatedAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPid := self.Args("pid").MustInt()
	iCreated := self.Args("created").MustInt()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iPid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByPidAndCreatedAndFieldName(offset, limit, iPid,iCreated,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByPidAndCreatedAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByPidAndCreatedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPid := self.Args("pid").MustInt()
	iCreated := self.Args("created").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iPid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByPidAndCreatedAndDefaultLangcode(offset, limit, iPid,iCreated,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByPidAndCreatedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByPidAndChangedAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPid := self.Args("pid").MustInt()
	iChanged := self.Args("changed").MustInt()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iPid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByPidAndChangedAndStatus(offset, limit, iPid,iChanged,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByPidAndChangedAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByPidAndChangedAndThreadHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPid := self.Args("pid").MustInt()
	iChanged := self.Args("changed").MustInt()
	iThread := self.Args("thread").String()

	if helper.IsHas(iPid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByPidAndChangedAndThread(offset, limit, iPid,iChanged,iThread)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByPidAndChangedAndThread's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByPidAndChangedAndEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPid := self.Args("pid").MustInt()
	iChanged := self.Args("changed").MustInt()
	iEntityType := self.Args("entity_type").String()

	if helper.IsHas(iPid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByPidAndChangedAndEntityType(offset, limit, iPid,iChanged,iEntityType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByPidAndChangedAndEntityType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByPidAndChangedAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPid := self.Args("pid").MustInt()
	iChanged := self.Args("changed").MustInt()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iPid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByPidAndChangedAndFieldName(offset, limit, iPid,iChanged,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByPidAndChangedAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByPidAndChangedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPid := self.Args("pid").MustInt()
	iChanged := self.Args("changed").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iPid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByPidAndChangedAndDefaultLangcode(offset, limit, iPid,iChanged,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByPidAndChangedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByPidAndStatusAndThreadHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPid := self.Args("pid").MustInt()
	iStatus := self.Args("status").MustInt()
	iThread := self.Args("thread").String()

	if helper.IsHas(iPid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByPidAndStatusAndThread(offset, limit, iPid,iStatus,iThread)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByPidAndStatusAndThread's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByPidAndStatusAndEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPid := self.Args("pid").MustInt()
	iStatus := self.Args("status").MustInt()
	iEntityType := self.Args("entity_type").String()

	if helper.IsHas(iPid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByPidAndStatusAndEntityType(offset, limit, iPid,iStatus,iEntityType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByPidAndStatusAndEntityType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByPidAndStatusAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPid := self.Args("pid").MustInt()
	iStatus := self.Args("status").MustInt()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iPid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByPidAndStatusAndFieldName(offset, limit, iPid,iStatus,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByPidAndStatusAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByPidAndStatusAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPid := self.Args("pid").MustInt()
	iStatus := self.Args("status").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iPid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByPidAndStatusAndDefaultLangcode(offset, limit, iPid,iStatus,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByPidAndStatusAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByPidAndThreadAndEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPid := self.Args("pid").MustInt()
	iThread := self.Args("thread").String()
	iEntityType := self.Args("entity_type").String()

	if helper.IsHas(iPid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByPidAndThreadAndEntityType(offset, limit, iPid,iThread,iEntityType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByPidAndThreadAndEntityType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByPidAndThreadAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPid := self.Args("pid").MustInt()
	iThread := self.Args("thread").String()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iPid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByPidAndThreadAndFieldName(offset, limit, iPid,iThread,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByPidAndThreadAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByPidAndThreadAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPid := self.Args("pid").MustInt()
	iThread := self.Args("thread").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iPid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByPidAndThreadAndDefaultLangcode(offset, limit, iPid,iThread,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByPidAndThreadAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByPidAndEntityTypeAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPid := self.Args("pid").MustInt()
	iEntityType := self.Args("entity_type").String()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iPid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByPidAndEntityTypeAndFieldName(offset, limit, iPid,iEntityType,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByPidAndEntityTypeAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByPidAndEntityTypeAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPid := self.Args("pid").MustInt()
	iEntityType := self.Args("entity_type").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iPid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByPidAndEntityTypeAndDefaultLangcode(offset, limit, iPid,iEntityType,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByPidAndEntityTypeAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByPidAndFieldNameAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPid := self.Args("pid").MustInt()
	iFieldName := self.Args("field_name").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iPid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByPidAndFieldNameAndDefaultLangcode(offset, limit, iPid,iFieldName,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByPidAndFieldNameAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByEntityIdAndSubjectAndUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iSubject := self.Args("subject").String()
	iUid := self.Args("uid").MustInt()

	if helper.IsHas(iEntityId) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByEntityIdAndSubjectAndUid(offset, limit, iEntityId,iSubject,iUid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByEntityIdAndSubjectAndUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByEntityIdAndSubjectAndNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iSubject := self.Args("subject").String()
	iName := self.Args("name").String()

	if helper.IsHas(iEntityId) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByEntityIdAndSubjectAndName(offset, limit, iEntityId,iSubject,iName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByEntityIdAndSubjectAndName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByEntityIdAndSubjectAndMailHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iSubject := self.Args("subject").String()
	iMail := self.Args("mail").String()

	if helper.IsHas(iEntityId) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByEntityIdAndSubjectAndMail(offset, limit, iEntityId,iSubject,iMail)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByEntityIdAndSubjectAndMail's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByEntityIdAndSubjectAndHomepageHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iSubject := self.Args("subject").String()
	iHomepage := self.Args("homepage").String()

	if helper.IsHas(iEntityId) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByEntityIdAndSubjectAndHomepage(offset, limit, iEntityId,iSubject,iHomepage)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByEntityIdAndSubjectAndHomepage's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByEntityIdAndSubjectAndHostnameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iSubject := self.Args("subject").String()
	iHostname := self.Args("hostname").String()

	if helper.IsHas(iEntityId) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByEntityIdAndSubjectAndHostname(offset, limit, iEntityId,iSubject,iHostname)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByEntityIdAndSubjectAndHostname's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByEntityIdAndSubjectAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iSubject := self.Args("subject").String()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iEntityId) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByEntityIdAndSubjectAndCreated(offset, limit, iEntityId,iSubject,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByEntityIdAndSubjectAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByEntityIdAndSubjectAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iSubject := self.Args("subject").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iEntityId) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByEntityIdAndSubjectAndChanged(offset, limit, iEntityId,iSubject,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByEntityIdAndSubjectAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByEntityIdAndSubjectAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iSubject := self.Args("subject").String()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iEntityId) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByEntityIdAndSubjectAndStatus(offset, limit, iEntityId,iSubject,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByEntityIdAndSubjectAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByEntityIdAndSubjectAndThreadHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iSubject := self.Args("subject").String()
	iThread := self.Args("thread").String()

	if helper.IsHas(iEntityId) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByEntityIdAndSubjectAndThread(offset, limit, iEntityId,iSubject,iThread)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByEntityIdAndSubjectAndThread's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByEntityIdAndSubjectAndEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iSubject := self.Args("subject").String()
	iEntityType := self.Args("entity_type").String()

	if helper.IsHas(iEntityId) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByEntityIdAndSubjectAndEntityType(offset, limit, iEntityId,iSubject,iEntityType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByEntityIdAndSubjectAndEntityType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByEntityIdAndSubjectAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iSubject := self.Args("subject").String()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iEntityId) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByEntityIdAndSubjectAndFieldName(offset, limit, iEntityId,iSubject,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByEntityIdAndSubjectAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByEntityIdAndSubjectAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iSubject := self.Args("subject").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iEntityId) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByEntityIdAndSubjectAndDefaultLangcode(offset, limit, iEntityId,iSubject,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByEntityIdAndSubjectAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByEntityIdAndUidAndNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iUid := self.Args("uid").MustInt()
	iName := self.Args("name").String()

	if helper.IsHas(iEntityId) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByEntityIdAndUidAndName(offset, limit, iEntityId,iUid,iName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByEntityIdAndUidAndName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByEntityIdAndUidAndMailHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iUid := self.Args("uid").MustInt()
	iMail := self.Args("mail").String()

	if helper.IsHas(iEntityId) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByEntityIdAndUidAndMail(offset, limit, iEntityId,iUid,iMail)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByEntityIdAndUidAndMail's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByEntityIdAndUidAndHomepageHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iUid := self.Args("uid").MustInt()
	iHomepage := self.Args("homepage").String()

	if helper.IsHas(iEntityId) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByEntityIdAndUidAndHomepage(offset, limit, iEntityId,iUid,iHomepage)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByEntityIdAndUidAndHomepage's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByEntityIdAndUidAndHostnameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iUid := self.Args("uid").MustInt()
	iHostname := self.Args("hostname").String()

	if helper.IsHas(iEntityId) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByEntityIdAndUidAndHostname(offset, limit, iEntityId,iUid,iHostname)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByEntityIdAndUidAndHostname's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByEntityIdAndUidAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iUid := self.Args("uid").MustInt()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iEntityId) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByEntityIdAndUidAndCreated(offset, limit, iEntityId,iUid,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByEntityIdAndUidAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByEntityIdAndUidAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iUid := self.Args("uid").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iEntityId) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByEntityIdAndUidAndChanged(offset, limit, iEntityId,iUid,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByEntityIdAndUidAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByEntityIdAndUidAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iUid := self.Args("uid").MustInt()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iEntityId) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByEntityIdAndUidAndStatus(offset, limit, iEntityId,iUid,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByEntityIdAndUidAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByEntityIdAndUidAndThreadHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iUid := self.Args("uid").MustInt()
	iThread := self.Args("thread").String()

	if helper.IsHas(iEntityId) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByEntityIdAndUidAndThread(offset, limit, iEntityId,iUid,iThread)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByEntityIdAndUidAndThread's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByEntityIdAndUidAndEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iUid := self.Args("uid").MustInt()
	iEntityType := self.Args("entity_type").String()

	if helper.IsHas(iEntityId) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByEntityIdAndUidAndEntityType(offset, limit, iEntityId,iUid,iEntityType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByEntityIdAndUidAndEntityType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByEntityIdAndUidAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iUid := self.Args("uid").MustInt()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iEntityId) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByEntityIdAndUidAndFieldName(offset, limit, iEntityId,iUid,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByEntityIdAndUidAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByEntityIdAndUidAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iUid := self.Args("uid").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iEntityId) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByEntityIdAndUidAndDefaultLangcode(offset, limit, iEntityId,iUid,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByEntityIdAndUidAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByEntityIdAndNameAndMailHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iName := self.Args("name").String()
	iMail := self.Args("mail").String()

	if helper.IsHas(iEntityId) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByEntityIdAndNameAndMail(offset, limit, iEntityId,iName,iMail)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByEntityIdAndNameAndMail's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByEntityIdAndNameAndHomepageHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iName := self.Args("name").String()
	iHomepage := self.Args("homepage").String()

	if helper.IsHas(iEntityId) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByEntityIdAndNameAndHomepage(offset, limit, iEntityId,iName,iHomepage)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByEntityIdAndNameAndHomepage's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByEntityIdAndNameAndHostnameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iName := self.Args("name").String()
	iHostname := self.Args("hostname").String()

	if helper.IsHas(iEntityId) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByEntityIdAndNameAndHostname(offset, limit, iEntityId,iName,iHostname)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByEntityIdAndNameAndHostname's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByEntityIdAndNameAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iName := self.Args("name").String()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iEntityId) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByEntityIdAndNameAndCreated(offset, limit, iEntityId,iName,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByEntityIdAndNameAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByEntityIdAndNameAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iName := self.Args("name").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iEntityId) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByEntityIdAndNameAndChanged(offset, limit, iEntityId,iName,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByEntityIdAndNameAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByEntityIdAndNameAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iName := self.Args("name").String()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iEntityId) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByEntityIdAndNameAndStatus(offset, limit, iEntityId,iName,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByEntityIdAndNameAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByEntityIdAndNameAndThreadHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iName := self.Args("name").String()
	iThread := self.Args("thread").String()

	if helper.IsHas(iEntityId) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByEntityIdAndNameAndThread(offset, limit, iEntityId,iName,iThread)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByEntityIdAndNameAndThread's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByEntityIdAndNameAndEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iName := self.Args("name").String()
	iEntityType := self.Args("entity_type").String()

	if helper.IsHas(iEntityId) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByEntityIdAndNameAndEntityType(offset, limit, iEntityId,iName,iEntityType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByEntityIdAndNameAndEntityType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByEntityIdAndNameAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iName := self.Args("name").String()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iEntityId) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByEntityIdAndNameAndFieldName(offset, limit, iEntityId,iName,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByEntityIdAndNameAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByEntityIdAndNameAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iName := self.Args("name").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iEntityId) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByEntityIdAndNameAndDefaultLangcode(offset, limit, iEntityId,iName,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByEntityIdAndNameAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByEntityIdAndMailAndHomepageHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iMail := self.Args("mail").String()
	iHomepage := self.Args("homepage").String()

	if helper.IsHas(iEntityId) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByEntityIdAndMailAndHomepage(offset, limit, iEntityId,iMail,iHomepage)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByEntityIdAndMailAndHomepage's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByEntityIdAndMailAndHostnameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iMail := self.Args("mail").String()
	iHostname := self.Args("hostname").String()

	if helper.IsHas(iEntityId) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByEntityIdAndMailAndHostname(offset, limit, iEntityId,iMail,iHostname)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByEntityIdAndMailAndHostname's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByEntityIdAndMailAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iMail := self.Args("mail").String()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iEntityId) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByEntityIdAndMailAndCreated(offset, limit, iEntityId,iMail,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByEntityIdAndMailAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByEntityIdAndMailAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iMail := self.Args("mail").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iEntityId) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByEntityIdAndMailAndChanged(offset, limit, iEntityId,iMail,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByEntityIdAndMailAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByEntityIdAndMailAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iMail := self.Args("mail").String()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iEntityId) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByEntityIdAndMailAndStatus(offset, limit, iEntityId,iMail,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByEntityIdAndMailAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByEntityIdAndMailAndThreadHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iMail := self.Args("mail").String()
	iThread := self.Args("thread").String()

	if helper.IsHas(iEntityId) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByEntityIdAndMailAndThread(offset, limit, iEntityId,iMail,iThread)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByEntityIdAndMailAndThread's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByEntityIdAndMailAndEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iMail := self.Args("mail").String()
	iEntityType := self.Args("entity_type").String()

	if helper.IsHas(iEntityId) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByEntityIdAndMailAndEntityType(offset, limit, iEntityId,iMail,iEntityType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByEntityIdAndMailAndEntityType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByEntityIdAndMailAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iMail := self.Args("mail").String()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iEntityId) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByEntityIdAndMailAndFieldName(offset, limit, iEntityId,iMail,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByEntityIdAndMailAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByEntityIdAndMailAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iMail := self.Args("mail").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iEntityId) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByEntityIdAndMailAndDefaultLangcode(offset, limit, iEntityId,iMail,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByEntityIdAndMailAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByEntityIdAndHomepageAndHostnameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iHomepage := self.Args("homepage").String()
	iHostname := self.Args("hostname").String()

	if helper.IsHas(iEntityId) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByEntityIdAndHomepageAndHostname(offset, limit, iEntityId,iHomepage,iHostname)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByEntityIdAndHomepageAndHostname's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByEntityIdAndHomepageAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iHomepage := self.Args("homepage").String()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iEntityId) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByEntityIdAndHomepageAndCreated(offset, limit, iEntityId,iHomepage,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByEntityIdAndHomepageAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByEntityIdAndHomepageAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iHomepage := self.Args("homepage").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iEntityId) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByEntityIdAndHomepageAndChanged(offset, limit, iEntityId,iHomepage,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByEntityIdAndHomepageAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByEntityIdAndHomepageAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iHomepage := self.Args("homepage").String()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iEntityId) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByEntityIdAndHomepageAndStatus(offset, limit, iEntityId,iHomepage,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByEntityIdAndHomepageAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByEntityIdAndHomepageAndThreadHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iHomepage := self.Args("homepage").String()
	iThread := self.Args("thread").String()

	if helper.IsHas(iEntityId) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByEntityIdAndHomepageAndThread(offset, limit, iEntityId,iHomepage,iThread)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByEntityIdAndHomepageAndThread's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByEntityIdAndHomepageAndEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iHomepage := self.Args("homepage").String()
	iEntityType := self.Args("entity_type").String()

	if helper.IsHas(iEntityId) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByEntityIdAndHomepageAndEntityType(offset, limit, iEntityId,iHomepage,iEntityType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByEntityIdAndHomepageAndEntityType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByEntityIdAndHomepageAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iHomepage := self.Args("homepage").String()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iEntityId) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByEntityIdAndHomepageAndFieldName(offset, limit, iEntityId,iHomepage,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByEntityIdAndHomepageAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByEntityIdAndHomepageAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iHomepage := self.Args("homepage").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iEntityId) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByEntityIdAndHomepageAndDefaultLangcode(offset, limit, iEntityId,iHomepage,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByEntityIdAndHomepageAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByEntityIdAndHostnameAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iHostname := self.Args("hostname").String()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iEntityId) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByEntityIdAndHostnameAndCreated(offset, limit, iEntityId,iHostname,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByEntityIdAndHostnameAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByEntityIdAndHostnameAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iHostname := self.Args("hostname").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iEntityId) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByEntityIdAndHostnameAndChanged(offset, limit, iEntityId,iHostname,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByEntityIdAndHostnameAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByEntityIdAndHostnameAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iHostname := self.Args("hostname").String()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iEntityId) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByEntityIdAndHostnameAndStatus(offset, limit, iEntityId,iHostname,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByEntityIdAndHostnameAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByEntityIdAndHostnameAndThreadHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iHostname := self.Args("hostname").String()
	iThread := self.Args("thread").String()

	if helper.IsHas(iEntityId) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByEntityIdAndHostnameAndThread(offset, limit, iEntityId,iHostname,iThread)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByEntityIdAndHostnameAndThread's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByEntityIdAndHostnameAndEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iHostname := self.Args("hostname").String()
	iEntityType := self.Args("entity_type").String()

	if helper.IsHas(iEntityId) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByEntityIdAndHostnameAndEntityType(offset, limit, iEntityId,iHostname,iEntityType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByEntityIdAndHostnameAndEntityType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByEntityIdAndHostnameAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iHostname := self.Args("hostname").String()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iEntityId) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByEntityIdAndHostnameAndFieldName(offset, limit, iEntityId,iHostname,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByEntityIdAndHostnameAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByEntityIdAndHostnameAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iHostname := self.Args("hostname").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iEntityId) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByEntityIdAndHostnameAndDefaultLangcode(offset, limit, iEntityId,iHostname,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByEntityIdAndHostnameAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByEntityIdAndCreatedAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iCreated := self.Args("created").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iEntityId) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByEntityIdAndCreatedAndChanged(offset, limit, iEntityId,iCreated,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByEntityIdAndCreatedAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByEntityIdAndCreatedAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iCreated := self.Args("created").MustInt()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iEntityId) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByEntityIdAndCreatedAndStatus(offset, limit, iEntityId,iCreated,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByEntityIdAndCreatedAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByEntityIdAndCreatedAndThreadHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iCreated := self.Args("created").MustInt()
	iThread := self.Args("thread").String()

	if helper.IsHas(iEntityId) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByEntityIdAndCreatedAndThread(offset, limit, iEntityId,iCreated,iThread)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByEntityIdAndCreatedAndThread's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByEntityIdAndCreatedAndEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iCreated := self.Args("created").MustInt()
	iEntityType := self.Args("entity_type").String()

	if helper.IsHas(iEntityId) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByEntityIdAndCreatedAndEntityType(offset, limit, iEntityId,iCreated,iEntityType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByEntityIdAndCreatedAndEntityType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByEntityIdAndCreatedAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iCreated := self.Args("created").MustInt()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iEntityId) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByEntityIdAndCreatedAndFieldName(offset, limit, iEntityId,iCreated,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByEntityIdAndCreatedAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByEntityIdAndCreatedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iCreated := self.Args("created").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iEntityId) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByEntityIdAndCreatedAndDefaultLangcode(offset, limit, iEntityId,iCreated,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByEntityIdAndCreatedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByEntityIdAndChangedAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iChanged := self.Args("changed").MustInt()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iEntityId) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByEntityIdAndChangedAndStatus(offset, limit, iEntityId,iChanged,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByEntityIdAndChangedAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByEntityIdAndChangedAndThreadHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iChanged := self.Args("changed").MustInt()
	iThread := self.Args("thread").String()

	if helper.IsHas(iEntityId) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByEntityIdAndChangedAndThread(offset, limit, iEntityId,iChanged,iThread)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByEntityIdAndChangedAndThread's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByEntityIdAndChangedAndEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iChanged := self.Args("changed").MustInt()
	iEntityType := self.Args("entity_type").String()

	if helper.IsHas(iEntityId) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByEntityIdAndChangedAndEntityType(offset, limit, iEntityId,iChanged,iEntityType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByEntityIdAndChangedAndEntityType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByEntityIdAndChangedAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iChanged := self.Args("changed").MustInt()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iEntityId) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByEntityIdAndChangedAndFieldName(offset, limit, iEntityId,iChanged,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByEntityIdAndChangedAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByEntityIdAndChangedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iChanged := self.Args("changed").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iEntityId) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByEntityIdAndChangedAndDefaultLangcode(offset, limit, iEntityId,iChanged,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByEntityIdAndChangedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByEntityIdAndStatusAndThreadHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iStatus := self.Args("status").MustInt()
	iThread := self.Args("thread").String()

	if helper.IsHas(iEntityId) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByEntityIdAndStatusAndThread(offset, limit, iEntityId,iStatus,iThread)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByEntityIdAndStatusAndThread's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByEntityIdAndStatusAndEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iStatus := self.Args("status").MustInt()
	iEntityType := self.Args("entity_type").String()

	if helper.IsHas(iEntityId) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByEntityIdAndStatusAndEntityType(offset, limit, iEntityId,iStatus,iEntityType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByEntityIdAndStatusAndEntityType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByEntityIdAndStatusAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iStatus := self.Args("status").MustInt()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iEntityId) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByEntityIdAndStatusAndFieldName(offset, limit, iEntityId,iStatus,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByEntityIdAndStatusAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByEntityIdAndStatusAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iStatus := self.Args("status").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iEntityId) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByEntityIdAndStatusAndDefaultLangcode(offset, limit, iEntityId,iStatus,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByEntityIdAndStatusAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByEntityIdAndThreadAndEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iThread := self.Args("thread").String()
	iEntityType := self.Args("entity_type").String()

	if helper.IsHas(iEntityId) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByEntityIdAndThreadAndEntityType(offset, limit, iEntityId,iThread,iEntityType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByEntityIdAndThreadAndEntityType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByEntityIdAndThreadAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iThread := self.Args("thread").String()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iEntityId) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByEntityIdAndThreadAndFieldName(offset, limit, iEntityId,iThread,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByEntityIdAndThreadAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByEntityIdAndThreadAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iThread := self.Args("thread").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iEntityId) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByEntityIdAndThreadAndDefaultLangcode(offset, limit, iEntityId,iThread,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByEntityIdAndThreadAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByEntityIdAndEntityTypeAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iEntityType := self.Args("entity_type").String()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iEntityId) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByEntityIdAndEntityTypeAndFieldName(offset, limit, iEntityId,iEntityType,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByEntityIdAndEntityTypeAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByEntityIdAndEntityTypeAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iEntityType := self.Args("entity_type").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iEntityId) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByEntityIdAndEntityTypeAndDefaultLangcode(offset, limit, iEntityId,iEntityType,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByEntityIdAndEntityTypeAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByEntityIdAndFieldNameAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iFieldName := self.Args("field_name").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iEntityId) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByEntityIdAndFieldNameAndDefaultLangcode(offset, limit, iEntityId,iFieldName,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByEntityIdAndFieldNameAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasBySubjectAndUidAndNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSubject := self.Args("subject").String()
	iUid := self.Args("uid").MustInt()
	iName := self.Args("name").String()

	if helper.IsHas(iSubject) {
		_CommentFieldData, _error := model.GetCommentFieldDatasBySubjectAndUidAndName(offset, limit, iSubject,iUid,iName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasBySubjectAndUidAndName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasBySubjectAndUidAndMailHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSubject := self.Args("subject").String()
	iUid := self.Args("uid").MustInt()
	iMail := self.Args("mail").String()

	if helper.IsHas(iSubject) {
		_CommentFieldData, _error := model.GetCommentFieldDatasBySubjectAndUidAndMail(offset, limit, iSubject,iUid,iMail)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasBySubjectAndUidAndMail's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasBySubjectAndUidAndHomepageHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSubject := self.Args("subject").String()
	iUid := self.Args("uid").MustInt()
	iHomepage := self.Args("homepage").String()

	if helper.IsHas(iSubject) {
		_CommentFieldData, _error := model.GetCommentFieldDatasBySubjectAndUidAndHomepage(offset, limit, iSubject,iUid,iHomepage)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasBySubjectAndUidAndHomepage's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasBySubjectAndUidAndHostnameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSubject := self.Args("subject").String()
	iUid := self.Args("uid").MustInt()
	iHostname := self.Args("hostname").String()

	if helper.IsHas(iSubject) {
		_CommentFieldData, _error := model.GetCommentFieldDatasBySubjectAndUidAndHostname(offset, limit, iSubject,iUid,iHostname)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasBySubjectAndUidAndHostname's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasBySubjectAndUidAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSubject := self.Args("subject").String()
	iUid := self.Args("uid").MustInt()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iSubject) {
		_CommentFieldData, _error := model.GetCommentFieldDatasBySubjectAndUidAndCreated(offset, limit, iSubject,iUid,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasBySubjectAndUidAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasBySubjectAndUidAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSubject := self.Args("subject").String()
	iUid := self.Args("uid").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iSubject) {
		_CommentFieldData, _error := model.GetCommentFieldDatasBySubjectAndUidAndChanged(offset, limit, iSubject,iUid,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasBySubjectAndUidAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasBySubjectAndUidAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSubject := self.Args("subject").String()
	iUid := self.Args("uid").MustInt()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iSubject) {
		_CommentFieldData, _error := model.GetCommentFieldDatasBySubjectAndUidAndStatus(offset, limit, iSubject,iUid,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasBySubjectAndUidAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasBySubjectAndUidAndThreadHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSubject := self.Args("subject").String()
	iUid := self.Args("uid").MustInt()
	iThread := self.Args("thread").String()

	if helper.IsHas(iSubject) {
		_CommentFieldData, _error := model.GetCommentFieldDatasBySubjectAndUidAndThread(offset, limit, iSubject,iUid,iThread)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasBySubjectAndUidAndThread's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasBySubjectAndUidAndEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSubject := self.Args("subject").String()
	iUid := self.Args("uid").MustInt()
	iEntityType := self.Args("entity_type").String()

	if helper.IsHas(iSubject) {
		_CommentFieldData, _error := model.GetCommentFieldDatasBySubjectAndUidAndEntityType(offset, limit, iSubject,iUid,iEntityType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasBySubjectAndUidAndEntityType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasBySubjectAndUidAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSubject := self.Args("subject").String()
	iUid := self.Args("uid").MustInt()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iSubject) {
		_CommentFieldData, _error := model.GetCommentFieldDatasBySubjectAndUidAndFieldName(offset, limit, iSubject,iUid,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasBySubjectAndUidAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasBySubjectAndUidAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSubject := self.Args("subject").String()
	iUid := self.Args("uid").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iSubject) {
		_CommentFieldData, _error := model.GetCommentFieldDatasBySubjectAndUidAndDefaultLangcode(offset, limit, iSubject,iUid,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasBySubjectAndUidAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasBySubjectAndNameAndMailHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSubject := self.Args("subject").String()
	iName := self.Args("name").String()
	iMail := self.Args("mail").String()

	if helper.IsHas(iSubject) {
		_CommentFieldData, _error := model.GetCommentFieldDatasBySubjectAndNameAndMail(offset, limit, iSubject,iName,iMail)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasBySubjectAndNameAndMail's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasBySubjectAndNameAndHomepageHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSubject := self.Args("subject").String()
	iName := self.Args("name").String()
	iHomepage := self.Args("homepage").String()

	if helper.IsHas(iSubject) {
		_CommentFieldData, _error := model.GetCommentFieldDatasBySubjectAndNameAndHomepage(offset, limit, iSubject,iName,iHomepage)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasBySubjectAndNameAndHomepage's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasBySubjectAndNameAndHostnameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSubject := self.Args("subject").String()
	iName := self.Args("name").String()
	iHostname := self.Args("hostname").String()

	if helper.IsHas(iSubject) {
		_CommentFieldData, _error := model.GetCommentFieldDatasBySubjectAndNameAndHostname(offset, limit, iSubject,iName,iHostname)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasBySubjectAndNameAndHostname's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasBySubjectAndNameAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSubject := self.Args("subject").String()
	iName := self.Args("name").String()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iSubject) {
		_CommentFieldData, _error := model.GetCommentFieldDatasBySubjectAndNameAndCreated(offset, limit, iSubject,iName,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasBySubjectAndNameAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasBySubjectAndNameAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSubject := self.Args("subject").String()
	iName := self.Args("name").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iSubject) {
		_CommentFieldData, _error := model.GetCommentFieldDatasBySubjectAndNameAndChanged(offset, limit, iSubject,iName,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasBySubjectAndNameAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasBySubjectAndNameAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSubject := self.Args("subject").String()
	iName := self.Args("name").String()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iSubject) {
		_CommentFieldData, _error := model.GetCommentFieldDatasBySubjectAndNameAndStatus(offset, limit, iSubject,iName,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasBySubjectAndNameAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasBySubjectAndNameAndThreadHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSubject := self.Args("subject").String()
	iName := self.Args("name").String()
	iThread := self.Args("thread").String()

	if helper.IsHas(iSubject) {
		_CommentFieldData, _error := model.GetCommentFieldDatasBySubjectAndNameAndThread(offset, limit, iSubject,iName,iThread)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasBySubjectAndNameAndThread's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasBySubjectAndNameAndEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSubject := self.Args("subject").String()
	iName := self.Args("name").String()
	iEntityType := self.Args("entity_type").String()

	if helper.IsHas(iSubject) {
		_CommentFieldData, _error := model.GetCommentFieldDatasBySubjectAndNameAndEntityType(offset, limit, iSubject,iName,iEntityType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasBySubjectAndNameAndEntityType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasBySubjectAndNameAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSubject := self.Args("subject").String()
	iName := self.Args("name").String()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iSubject) {
		_CommentFieldData, _error := model.GetCommentFieldDatasBySubjectAndNameAndFieldName(offset, limit, iSubject,iName,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasBySubjectAndNameAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasBySubjectAndNameAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSubject := self.Args("subject").String()
	iName := self.Args("name").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iSubject) {
		_CommentFieldData, _error := model.GetCommentFieldDatasBySubjectAndNameAndDefaultLangcode(offset, limit, iSubject,iName,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasBySubjectAndNameAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasBySubjectAndMailAndHomepageHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSubject := self.Args("subject").String()
	iMail := self.Args("mail").String()
	iHomepage := self.Args("homepage").String()

	if helper.IsHas(iSubject) {
		_CommentFieldData, _error := model.GetCommentFieldDatasBySubjectAndMailAndHomepage(offset, limit, iSubject,iMail,iHomepage)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasBySubjectAndMailAndHomepage's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasBySubjectAndMailAndHostnameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSubject := self.Args("subject").String()
	iMail := self.Args("mail").String()
	iHostname := self.Args("hostname").String()

	if helper.IsHas(iSubject) {
		_CommentFieldData, _error := model.GetCommentFieldDatasBySubjectAndMailAndHostname(offset, limit, iSubject,iMail,iHostname)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasBySubjectAndMailAndHostname's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasBySubjectAndMailAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSubject := self.Args("subject").String()
	iMail := self.Args("mail").String()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iSubject) {
		_CommentFieldData, _error := model.GetCommentFieldDatasBySubjectAndMailAndCreated(offset, limit, iSubject,iMail,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasBySubjectAndMailAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasBySubjectAndMailAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSubject := self.Args("subject").String()
	iMail := self.Args("mail").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iSubject) {
		_CommentFieldData, _error := model.GetCommentFieldDatasBySubjectAndMailAndChanged(offset, limit, iSubject,iMail,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasBySubjectAndMailAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasBySubjectAndMailAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSubject := self.Args("subject").String()
	iMail := self.Args("mail").String()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iSubject) {
		_CommentFieldData, _error := model.GetCommentFieldDatasBySubjectAndMailAndStatus(offset, limit, iSubject,iMail,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasBySubjectAndMailAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasBySubjectAndMailAndThreadHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSubject := self.Args("subject").String()
	iMail := self.Args("mail").String()
	iThread := self.Args("thread").String()

	if helper.IsHas(iSubject) {
		_CommentFieldData, _error := model.GetCommentFieldDatasBySubjectAndMailAndThread(offset, limit, iSubject,iMail,iThread)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasBySubjectAndMailAndThread's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasBySubjectAndMailAndEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSubject := self.Args("subject").String()
	iMail := self.Args("mail").String()
	iEntityType := self.Args("entity_type").String()

	if helper.IsHas(iSubject) {
		_CommentFieldData, _error := model.GetCommentFieldDatasBySubjectAndMailAndEntityType(offset, limit, iSubject,iMail,iEntityType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasBySubjectAndMailAndEntityType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasBySubjectAndMailAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSubject := self.Args("subject").String()
	iMail := self.Args("mail").String()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iSubject) {
		_CommentFieldData, _error := model.GetCommentFieldDatasBySubjectAndMailAndFieldName(offset, limit, iSubject,iMail,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasBySubjectAndMailAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasBySubjectAndMailAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSubject := self.Args("subject").String()
	iMail := self.Args("mail").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iSubject) {
		_CommentFieldData, _error := model.GetCommentFieldDatasBySubjectAndMailAndDefaultLangcode(offset, limit, iSubject,iMail,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasBySubjectAndMailAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasBySubjectAndHomepageAndHostnameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSubject := self.Args("subject").String()
	iHomepage := self.Args("homepage").String()
	iHostname := self.Args("hostname").String()

	if helper.IsHas(iSubject) {
		_CommentFieldData, _error := model.GetCommentFieldDatasBySubjectAndHomepageAndHostname(offset, limit, iSubject,iHomepage,iHostname)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasBySubjectAndHomepageAndHostname's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasBySubjectAndHomepageAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSubject := self.Args("subject").String()
	iHomepage := self.Args("homepage").String()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iSubject) {
		_CommentFieldData, _error := model.GetCommentFieldDatasBySubjectAndHomepageAndCreated(offset, limit, iSubject,iHomepage,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasBySubjectAndHomepageAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasBySubjectAndHomepageAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSubject := self.Args("subject").String()
	iHomepage := self.Args("homepage").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iSubject) {
		_CommentFieldData, _error := model.GetCommentFieldDatasBySubjectAndHomepageAndChanged(offset, limit, iSubject,iHomepage,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasBySubjectAndHomepageAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasBySubjectAndHomepageAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSubject := self.Args("subject").String()
	iHomepage := self.Args("homepage").String()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iSubject) {
		_CommentFieldData, _error := model.GetCommentFieldDatasBySubjectAndHomepageAndStatus(offset, limit, iSubject,iHomepage,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasBySubjectAndHomepageAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasBySubjectAndHomepageAndThreadHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSubject := self.Args("subject").String()
	iHomepage := self.Args("homepage").String()
	iThread := self.Args("thread").String()

	if helper.IsHas(iSubject) {
		_CommentFieldData, _error := model.GetCommentFieldDatasBySubjectAndHomepageAndThread(offset, limit, iSubject,iHomepage,iThread)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasBySubjectAndHomepageAndThread's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasBySubjectAndHomepageAndEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSubject := self.Args("subject").String()
	iHomepage := self.Args("homepage").String()
	iEntityType := self.Args("entity_type").String()

	if helper.IsHas(iSubject) {
		_CommentFieldData, _error := model.GetCommentFieldDatasBySubjectAndHomepageAndEntityType(offset, limit, iSubject,iHomepage,iEntityType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasBySubjectAndHomepageAndEntityType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasBySubjectAndHomepageAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSubject := self.Args("subject").String()
	iHomepage := self.Args("homepage").String()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iSubject) {
		_CommentFieldData, _error := model.GetCommentFieldDatasBySubjectAndHomepageAndFieldName(offset, limit, iSubject,iHomepage,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasBySubjectAndHomepageAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasBySubjectAndHomepageAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSubject := self.Args("subject").String()
	iHomepage := self.Args("homepage").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iSubject) {
		_CommentFieldData, _error := model.GetCommentFieldDatasBySubjectAndHomepageAndDefaultLangcode(offset, limit, iSubject,iHomepage,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasBySubjectAndHomepageAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasBySubjectAndHostnameAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSubject := self.Args("subject").String()
	iHostname := self.Args("hostname").String()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iSubject) {
		_CommentFieldData, _error := model.GetCommentFieldDatasBySubjectAndHostnameAndCreated(offset, limit, iSubject,iHostname,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasBySubjectAndHostnameAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasBySubjectAndHostnameAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSubject := self.Args("subject").String()
	iHostname := self.Args("hostname").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iSubject) {
		_CommentFieldData, _error := model.GetCommentFieldDatasBySubjectAndHostnameAndChanged(offset, limit, iSubject,iHostname,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasBySubjectAndHostnameAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasBySubjectAndHostnameAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSubject := self.Args("subject").String()
	iHostname := self.Args("hostname").String()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iSubject) {
		_CommentFieldData, _error := model.GetCommentFieldDatasBySubjectAndHostnameAndStatus(offset, limit, iSubject,iHostname,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasBySubjectAndHostnameAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasBySubjectAndHostnameAndThreadHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSubject := self.Args("subject").String()
	iHostname := self.Args("hostname").String()
	iThread := self.Args("thread").String()

	if helper.IsHas(iSubject) {
		_CommentFieldData, _error := model.GetCommentFieldDatasBySubjectAndHostnameAndThread(offset, limit, iSubject,iHostname,iThread)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasBySubjectAndHostnameAndThread's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasBySubjectAndHostnameAndEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSubject := self.Args("subject").String()
	iHostname := self.Args("hostname").String()
	iEntityType := self.Args("entity_type").String()

	if helper.IsHas(iSubject) {
		_CommentFieldData, _error := model.GetCommentFieldDatasBySubjectAndHostnameAndEntityType(offset, limit, iSubject,iHostname,iEntityType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasBySubjectAndHostnameAndEntityType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasBySubjectAndHostnameAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSubject := self.Args("subject").String()
	iHostname := self.Args("hostname").String()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iSubject) {
		_CommentFieldData, _error := model.GetCommentFieldDatasBySubjectAndHostnameAndFieldName(offset, limit, iSubject,iHostname,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasBySubjectAndHostnameAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasBySubjectAndHostnameAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSubject := self.Args("subject").String()
	iHostname := self.Args("hostname").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iSubject) {
		_CommentFieldData, _error := model.GetCommentFieldDatasBySubjectAndHostnameAndDefaultLangcode(offset, limit, iSubject,iHostname,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasBySubjectAndHostnameAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasBySubjectAndCreatedAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSubject := self.Args("subject").String()
	iCreated := self.Args("created").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iSubject) {
		_CommentFieldData, _error := model.GetCommentFieldDatasBySubjectAndCreatedAndChanged(offset, limit, iSubject,iCreated,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasBySubjectAndCreatedAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasBySubjectAndCreatedAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSubject := self.Args("subject").String()
	iCreated := self.Args("created").MustInt()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iSubject) {
		_CommentFieldData, _error := model.GetCommentFieldDatasBySubjectAndCreatedAndStatus(offset, limit, iSubject,iCreated,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasBySubjectAndCreatedAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasBySubjectAndCreatedAndThreadHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSubject := self.Args("subject").String()
	iCreated := self.Args("created").MustInt()
	iThread := self.Args("thread").String()

	if helper.IsHas(iSubject) {
		_CommentFieldData, _error := model.GetCommentFieldDatasBySubjectAndCreatedAndThread(offset, limit, iSubject,iCreated,iThread)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasBySubjectAndCreatedAndThread's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasBySubjectAndCreatedAndEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSubject := self.Args("subject").String()
	iCreated := self.Args("created").MustInt()
	iEntityType := self.Args("entity_type").String()

	if helper.IsHas(iSubject) {
		_CommentFieldData, _error := model.GetCommentFieldDatasBySubjectAndCreatedAndEntityType(offset, limit, iSubject,iCreated,iEntityType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasBySubjectAndCreatedAndEntityType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasBySubjectAndCreatedAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSubject := self.Args("subject").String()
	iCreated := self.Args("created").MustInt()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iSubject) {
		_CommentFieldData, _error := model.GetCommentFieldDatasBySubjectAndCreatedAndFieldName(offset, limit, iSubject,iCreated,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasBySubjectAndCreatedAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasBySubjectAndCreatedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSubject := self.Args("subject").String()
	iCreated := self.Args("created").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iSubject) {
		_CommentFieldData, _error := model.GetCommentFieldDatasBySubjectAndCreatedAndDefaultLangcode(offset, limit, iSubject,iCreated,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasBySubjectAndCreatedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasBySubjectAndChangedAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSubject := self.Args("subject").String()
	iChanged := self.Args("changed").MustInt()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iSubject) {
		_CommentFieldData, _error := model.GetCommentFieldDatasBySubjectAndChangedAndStatus(offset, limit, iSubject,iChanged,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasBySubjectAndChangedAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasBySubjectAndChangedAndThreadHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSubject := self.Args("subject").String()
	iChanged := self.Args("changed").MustInt()
	iThread := self.Args("thread").String()

	if helper.IsHas(iSubject) {
		_CommentFieldData, _error := model.GetCommentFieldDatasBySubjectAndChangedAndThread(offset, limit, iSubject,iChanged,iThread)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasBySubjectAndChangedAndThread's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasBySubjectAndChangedAndEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSubject := self.Args("subject").String()
	iChanged := self.Args("changed").MustInt()
	iEntityType := self.Args("entity_type").String()

	if helper.IsHas(iSubject) {
		_CommentFieldData, _error := model.GetCommentFieldDatasBySubjectAndChangedAndEntityType(offset, limit, iSubject,iChanged,iEntityType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasBySubjectAndChangedAndEntityType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasBySubjectAndChangedAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSubject := self.Args("subject").String()
	iChanged := self.Args("changed").MustInt()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iSubject) {
		_CommentFieldData, _error := model.GetCommentFieldDatasBySubjectAndChangedAndFieldName(offset, limit, iSubject,iChanged,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasBySubjectAndChangedAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasBySubjectAndChangedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSubject := self.Args("subject").String()
	iChanged := self.Args("changed").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iSubject) {
		_CommentFieldData, _error := model.GetCommentFieldDatasBySubjectAndChangedAndDefaultLangcode(offset, limit, iSubject,iChanged,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasBySubjectAndChangedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasBySubjectAndStatusAndThreadHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSubject := self.Args("subject").String()
	iStatus := self.Args("status").MustInt()
	iThread := self.Args("thread").String()

	if helper.IsHas(iSubject) {
		_CommentFieldData, _error := model.GetCommentFieldDatasBySubjectAndStatusAndThread(offset, limit, iSubject,iStatus,iThread)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasBySubjectAndStatusAndThread's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasBySubjectAndStatusAndEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSubject := self.Args("subject").String()
	iStatus := self.Args("status").MustInt()
	iEntityType := self.Args("entity_type").String()

	if helper.IsHas(iSubject) {
		_CommentFieldData, _error := model.GetCommentFieldDatasBySubjectAndStatusAndEntityType(offset, limit, iSubject,iStatus,iEntityType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasBySubjectAndStatusAndEntityType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasBySubjectAndStatusAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSubject := self.Args("subject").String()
	iStatus := self.Args("status").MustInt()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iSubject) {
		_CommentFieldData, _error := model.GetCommentFieldDatasBySubjectAndStatusAndFieldName(offset, limit, iSubject,iStatus,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasBySubjectAndStatusAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasBySubjectAndStatusAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSubject := self.Args("subject").String()
	iStatus := self.Args("status").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iSubject) {
		_CommentFieldData, _error := model.GetCommentFieldDatasBySubjectAndStatusAndDefaultLangcode(offset, limit, iSubject,iStatus,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasBySubjectAndStatusAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasBySubjectAndThreadAndEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSubject := self.Args("subject").String()
	iThread := self.Args("thread").String()
	iEntityType := self.Args("entity_type").String()

	if helper.IsHas(iSubject) {
		_CommentFieldData, _error := model.GetCommentFieldDatasBySubjectAndThreadAndEntityType(offset, limit, iSubject,iThread,iEntityType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasBySubjectAndThreadAndEntityType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasBySubjectAndThreadAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSubject := self.Args("subject").String()
	iThread := self.Args("thread").String()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iSubject) {
		_CommentFieldData, _error := model.GetCommentFieldDatasBySubjectAndThreadAndFieldName(offset, limit, iSubject,iThread,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasBySubjectAndThreadAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasBySubjectAndThreadAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSubject := self.Args("subject").String()
	iThread := self.Args("thread").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iSubject) {
		_CommentFieldData, _error := model.GetCommentFieldDatasBySubjectAndThreadAndDefaultLangcode(offset, limit, iSubject,iThread,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasBySubjectAndThreadAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasBySubjectAndEntityTypeAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSubject := self.Args("subject").String()
	iEntityType := self.Args("entity_type").String()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iSubject) {
		_CommentFieldData, _error := model.GetCommentFieldDatasBySubjectAndEntityTypeAndFieldName(offset, limit, iSubject,iEntityType,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasBySubjectAndEntityTypeAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasBySubjectAndEntityTypeAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSubject := self.Args("subject").String()
	iEntityType := self.Args("entity_type").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iSubject) {
		_CommentFieldData, _error := model.GetCommentFieldDatasBySubjectAndEntityTypeAndDefaultLangcode(offset, limit, iSubject,iEntityType,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasBySubjectAndEntityTypeAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasBySubjectAndFieldNameAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSubject := self.Args("subject").String()
	iFieldName := self.Args("field_name").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iSubject) {
		_CommentFieldData, _error := model.GetCommentFieldDatasBySubjectAndFieldNameAndDefaultLangcode(offset, limit, iSubject,iFieldName,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasBySubjectAndFieldNameAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByUidAndNameAndMailHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iName := self.Args("name").String()
	iMail := self.Args("mail").String()

	if helper.IsHas(iUid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByUidAndNameAndMail(offset, limit, iUid,iName,iMail)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByUidAndNameAndMail's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByUidAndNameAndHomepageHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iName := self.Args("name").String()
	iHomepage := self.Args("homepage").String()

	if helper.IsHas(iUid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByUidAndNameAndHomepage(offset, limit, iUid,iName,iHomepage)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByUidAndNameAndHomepage's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByUidAndNameAndHostnameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iName := self.Args("name").String()
	iHostname := self.Args("hostname").String()

	if helper.IsHas(iUid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByUidAndNameAndHostname(offset, limit, iUid,iName,iHostname)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByUidAndNameAndHostname's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByUidAndNameAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iName := self.Args("name").String()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iUid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByUidAndNameAndCreated(offset, limit, iUid,iName,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByUidAndNameAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByUidAndNameAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iName := self.Args("name").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iUid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByUidAndNameAndChanged(offset, limit, iUid,iName,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByUidAndNameAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByUidAndNameAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iName := self.Args("name").String()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iUid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByUidAndNameAndStatus(offset, limit, iUid,iName,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByUidAndNameAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByUidAndNameAndThreadHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iName := self.Args("name").String()
	iThread := self.Args("thread").String()

	if helper.IsHas(iUid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByUidAndNameAndThread(offset, limit, iUid,iName,iThread)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByUidAndNameAndThread's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByUidAndNameAndEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iName := self.Args("name").String()
	iEntityType := self.Args("entity_type").String()

	if helper.IsHas(iUid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByUidAndNameAndEntityType(offset, limit, iUid,iName,iEntityType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByUidAndNameAndEntityType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByUidAndNameAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iName := self.Args("name").String()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iUid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByUidAndNameAndFieldName(offset, limit, iUid,iName,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByUidAndNameAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByUidAndNameAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iName := self.Args("name").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iUid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByUidAndNameAndDefaultLangcode(offset, limit, iUid,iName,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByUidAndNameAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByUidAndMailAndHomepageHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iMail := self.Args("mail").String()
	iHomepage := self.Args("homepage").String()

	if helper.IsHas(iUid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByUidAndMailAndHomepage(offset, limit, iUid,iMail,iHomepage)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByUidAndMailAndHomepage's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByUidAndMailAndHostnameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iMail := self.Args("mail").String()
	iHostname := self.Args("hostname").String()

	if helper.IsHas(iUid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByUidAndMailAndHostname(offset, limit, iUid,iMail,iHostname)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByUidAndMailAndHostname's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByUidAndMailAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iMail := self.Args("mail").String()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iUid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByUidAndMailAndCreated(offset, limit, iUid,iMail,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByUidAndMailAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByUidAndMailAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iMail := self.Args("mail").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iUid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByUidAndMailAndChanged(offset, limit, iUid,iMail,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByUidAndMailAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByUidAndMailAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iMail := self.Args("mail").String()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iUid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByUidAndMailAndStatus(offset, limit, iUid,iMail,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByUidAndMailAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByUidAndMailAndThreadHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iMail := self.Args("mail").String()
	iThread := self.Args("thread").String()

	if helper.IsHas(iUid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByUidAndMailAndThread(offset, limit, iUid,iMail,iThread)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByUidAndMailAndThread's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByUidAndMailAndEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iMail := self.Args("mail").String()
	iEntityType := self.Args("entity_type").String()

	if helper.IsHas(iUid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByUidAndMailAndEntityType(offset, limit, iUid,iMail,iEntityType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByUidAndMailAndEntityType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByUidAndMailAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iMail := self.Args("mail").String()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iUid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByUidAndMailAndFieldName(offset, limit, iUid,iMail,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByUidAndMailAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByUidAndMailAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iMail := self.Args("mail").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iUid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByUidAndMailAndDefaultLangcode(offset, limit, iUid,iMail,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByUidAndMailAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByUidAndHomepageAndHostnameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iHomepage := self.Args("homepage").String()
	iHostname := self.Args("hostname").String()

	if helper.IsHas(iUid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByUidAndHomepageAndHostname(offset, limit, iUid,iHomepage,iHostname)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByUidAndHomepageAndHostname's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByUidAndHomepageAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iHomepage := self.Args("homepage").String()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iUid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByUidAndHomepageAndCreated(offset, limit, iUid,iHomepage,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByUidAndHomepageAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByUidAndHomepageAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iHomepage := self.Args("homepage").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iUid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByUidAndHomepageAndChanged(offset, limit, iUid,iHomepage,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByUidAndHomepageAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByUidAndHomepageAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iHomepage := self.Args("homepage").String()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iUid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByUidAndHomepageAndStatus(offset, limit, iUid,iHomepage,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByUidAndHomepageAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByUidAndHomepageAndThreadHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iHomepage := self.Args("homepage").String()
	iThread := self.Args("thread").String()

	if helper.IsHas(iUid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByUidAndHomepageAndThread(offset, limit, iUid,iHomepage,iThread)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByUidAndHomepageAndThread's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByUidAndHomepageAndEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iHomepage := self.Args("homepage").String()
	iEntityType := self.Args("entity_type").String()

	if helper.IsHas(iUid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByUidAndHomepageAndEntityType(offset, limit, iUid,iHomepage,iEntityType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByUidAndHomepageAndEntityType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByUidAndHomepageAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iHomepage := self.Args("homepage").String()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iUid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByUidAndHomepageAndFieldName(offset, limit, iUid,iHomepage,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByUidAndHomepageAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByUidAndHomepageAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iHomepage := self.Args("homepage").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iUid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByUidAndHomepageAndDefaultLangcode(offset, limit, iUid,iHomepage,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByUidAndHomepageAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByUidAndHostnameAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iHostname := self.Args("hostname").String()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iUid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByUidAndHostnameAndCreated(offset, limit, iUid,iHostname,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByUidAndHostnameAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByUidAndHostnameAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iHostname := self.Args("hostname").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iUid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByUidAndHostnameAndChanged(offset, limit, iUid,iHostname,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByUidAndHostnameAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByUidAndHostnameAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iHostname := self.Args("hostname").String()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iUid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByUidAndHostnameAndStatus(offset, limit, iUid,iHostname,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByUidAndHostnameAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByUidAndHostnameAndThreadHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iHostname := self.Args("hostname").String()
	iThread := self.Args("thread").String()

	if helper.IsHas(iUid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByUidAndHostnameAndThread(offset, limit, iUid,iHostname,iThread)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByUidAndHostnameAndThread's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByUidAndHostnameAndEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iHostname := self.Args("hostname").String()
	iEntityType := self.Args("entity_type").String()

	if helper.IsHas(iUid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByUidAndHostnameAndEntityType(offset, limit, iUid,iHostname,iEntityType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByUidAndHostnameAndEntityType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByUidAndHostnameAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iHostname := self.Args("hostname").String()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iUid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByUidAndHostnameAndFieldName(offset, limit, iUid,iHostname,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByUidAndHostnameAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByUidAndHostnameAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iHostname := self.Args("hostname").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iUid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByUidAndHostnameAndDefaultLangcode(offset, limit, iUid,iHostname,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByUidAndHostnameAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByUidAndCreatedAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iCreated := self.Args("created").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iUid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByUidAndCreatedAndChanged(offset, limit, iUid,iCreated,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByUidAndCreatedAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByUidAndCreatedAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iCreated := self.Args("created").MustInt()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iUid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByUidAndCreatedAndStatus(offset, limit, iUid,iCreated,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByUidAndCreatedAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByUidAndCreatedAndThreadHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iCreated := self.Args("created").MustInt()
	iThread := self.Args("thread").String()

	if helper.IsHas(iUid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByUidAndCreatedAndThread(offset, limit, iUid,iCreated,iThread)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByUidAndCreatedAndThread's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByUidAndCreatedAndEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iCreated := self.Args("created").MustInt()
	iEntityType := self.Args("entity_type").String()

	if helper.IsHas(iUid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByUidAndCreatedAndEntityType(offset, limit, iUid,iCreated,iEntityType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByUidAndCreatedAndEntityType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByUidAndCreatedAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iCreated := self.Args("created").MustInt()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iUid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByUidAndCreatedAndFieldName(offset, limit, iUid,iCreated,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByUidAndCreatedAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByUidAndCreatedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iCreated := self.Args("created").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iUid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByUidAndCreatedAndDefaultLangcode(offset, limit, iUid,iCreated,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByUidAndCreatedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByUidAndChangedAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iChanged := self.Args("changed").MustInt()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iUid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByUidAndChangedAndStatus(offset, limit, iUid,iChanged,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByUidAndChangedAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByUidAndChangedAndThreadHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iChanged := self.Args("changed").MustInt()
	iThread := self.Args("thread").String()

	if helper.IsHas(iUid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByUidAndChangedAndThread(offset, limit, iUid,iChanged,iThread)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByUidAndChangedAndThread's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByUidAndChangedAndEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iChanged := self.Args("changed").MustInt()
	iEntityType := self.Args("entity_type").String()

	if helper.IsHas(iUid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByUidAndChangedAndEntityType(offset, limit, iUid,iChanged,iEntityType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByUidAndChangedAndEntityType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByUidAndChangedAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iChanged := self.Args("changed").MustInt()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iUid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByUidAndChangedAndFieldName(offset, limit, iUid,iChanged,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByUidAndChangedAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByUidAndChangedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iChanged := self.Args("changed").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iUid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByUidAndChangedAndDefaultLangcode(offset, limit, iUid,iChanged,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByUidAndChangedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByUidAndStatusAndThreadHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iStatus := self.Args("status").MustInt()
	iThread := self.Args("thread").String()

	if helper.IsHas(iUid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByUidAndStatusAndThread(offset, limit, iUid,iStatus,iThread)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByUidAndStatusAndThread's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByUidAndStatusAndEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iStatus := self.Args("status").MustInt()
	iEntityType := self.Args("entity_type").String()

	if helper.IsHas(iUid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByUidAndStatusAndEntityType(offset, limit, iUid,iStatus,iEntityType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByUidAndStatusAndEntityType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByUidAndStatusAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iStatus := self.Args("status").MustInt()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iUid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByUidAndStatusAndFieldName(offset, limit, iUid,iStatus,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByUidAndStatusAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByUidAndStatusAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iStatus := self.Args("status").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iUid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByUidAndStatusAndDefaultLangcode(offset, limit, iUid,iStatus,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByUidAndStatusAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByUidAndThreadAndEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iThread := self.Args("thread").String()
	iEntityType := self.Args("entity_type").String()

	if helper.IsHas(iUid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByUidAndThreadAndEntityType(offset, limit, iUid,iThread,iEntityType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByUidAndThreadAndEntityType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByUidAndThreadAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iThread := self.Args("thread").String()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iUid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByUidAndThreadAndFieldName(offset, limit, iUid,iThread,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByUidAndThreadAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByUidAndThreadAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iThread := self.Args("thread").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iUid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByUidAndThreadAndDefaultLangcode(offset, limit, iUid,iThread,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByUidAndThreadAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByUidAndEntityTypeAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iEntityType := self.Args("entity_type").String()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iUid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByUidAndEntityTypeAndFieldName(offset, limit, iUid,iEntityType,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByUidAndEntityTypeAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByUidAndEntityTypeAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iEntityType := self.Args("entity_type").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iUid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByUidAndEntityTypeAndDefaultLangcode(offset, limit, iUid,iEntityType,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByUidAndEntityTypeAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByUidAndFieldNameAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iFieldName := self.Args("field_name").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iUid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByUidAndFieldNameAndDefaultLangcode(offset, limit, iUid,iFieldName,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByUidAndFieldNameAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByNameAndMailAndHomepageHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iMail := self.Args("mail").String()
	iHomepage := self.Args("homepage").String()

	if helper.IsHas(iName) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByNameAndMailAndHomepage(offset, limit, iName,iMail,iHomepage)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByNameAndMailAndHomepage's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByNameAndMailAndHostnameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iMail := self.Args("mail").String()
	iHostname := self.Args("hostname").String()

	if helper.IsHas(iName) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByNameAndMailAndHostname(offset, limit, iName,iMail,iHostname)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByNameAndMailAndHostname's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByNameAndMailAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iMail := self.Args("mail").String()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iName) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByNameAndMailAndCreated(offset, limit, iName,iMail,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByNameAndMailAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByNameAndMailAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iMail := self.Args("mail").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iName) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByNameAndMailAndChanged(offset, limit, iName,iMail,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByNameAndMailAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByNameAndMailAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iMail := self.Args("mail").String()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iName) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByNameAndMailAndStatus(offset, limit, iName,iMail,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByNameAndMailAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByNameAndMailAndThreadHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iMail := self.Args("mail").String()
	iThread := self.Args("thread").String()

	if helper.IsHas(iName) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByNameAndMailAndThread(offset, limit, iName,iMail,iThread)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByNameAndMailAndThread's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByNameAndMailAndEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iMail := self.Args("mail").String()
	iEntityType := self.Args("entity_type").String()

	if helper.IsHas(iName) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByNameAndMailAndEntityType(offset, limit, iName,iMail,iEntityType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByNameAndMailAndEntityType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByNameAndMailAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iMail := self.Args("mail").String()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iName) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByNameAndMailAndFieldName(offset, limit, iName,iMail,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByNameAndMailAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByNameAndMailAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iMail := self.Args("mail").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iName) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByNameAndMailAndDefaultLangcode(offset, limit, iName,iMail,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByNameAndMailAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByNameAndHomepageAndHostnameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iHomepage := self.Args("homepage").String()
	iHostname := self.Args("hostname").String()

	if helper.IsHas(iName) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByNameAndHomepageAndHostname(offset, limit, iName,iHomepage,iHostname)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByNameAndHomepageAndHostname's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByNameAndHomepageAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iHomepage := self.Args("homepage").String()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iName) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByNameAndHomepageAndCreated(offset, limit, iName,iHomepage,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByNameAndHomepageAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByNameAndHomepageAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iHomepage := self.Args("homepage").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iName) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByNameAndHomepageAndChanged(offset, limit, iName,iHomepage,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByNameAndHomepageAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByNameAndHomepageAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iHomepage := self.Args("homepage").String()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iName) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByNameAndHomepageAndStatus(offset, limit, iName,iHomepage,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByNameAndHomepageAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByNameAndHomepageAndThreadHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iHomepage := self.Args("homepage").String()
	iThread := self.Args("thread").String()

	if helper.IsHas(iName) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByNameAndHomepageAndThread(offset, limit, iName,iHomepage,iThread)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByNameAndHomepageAndThread's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByNameAndHomepageAndEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iHomepage := self.Args("homepage").String()
	iEntityType := self.Args("entity_type").String()

	if helper.IsHas(iName) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByNameAndHomepageAndEntityType(offset, limit, iName,iHomepage,iEntityType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByNameAndHomepageAndEntityType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByNameAndHomepageAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iHomepage := self.Args("homepage").String()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iName) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByNameAndHomepageAndFieldName(offset, limit, iName,iHomepage,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByNameAndHomepageAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByNameAndHomepageAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iHomepage := self.Args("homepage").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iName) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByNameAndHomepageAndDefaultLangcode(offset, limit, iName,iHomepage,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByNameAndHomepageAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByNameAndHostnameAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iHostname := self.Args("hostname").String()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iName) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByNameAndHostnameAndCreated(offset, limit, iName,iHostname,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByNameAndHostnameAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByNameAndHostnameAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iHostname := self.Args("hostname").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iName) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByNameAndHostnameAndChanged(offset, limit, iName,iHostname,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByNameAndHostnameAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByNameAndHostnameAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iHostname := self.Args("hostname").String()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iName) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByNameAndHostnameAndStatus(offset, limit, iName,iHostname,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByNameAndHostnameAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByNameAndHostnameAndThreadHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iHostname := self.Args("hostname").String()
	iThread := self.Args("thread").String()

	if helper.IsHas(iName) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByNameAndHostnameAndThread(offset, limit, iName,iHostname,iThread)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByNameAndHostnameAndThread's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByNameAndHostnameAndEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iHostname := self.Args("hostname").String()
	iEntityType := self.Args("entity_type").String()

	if helper.IsHas(iName) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByNameAndHostnameAndEntityType(offset, limit, iName,iHostname,iEntityType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByNameAndHostnameAndEntityType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByNameAndHostnameAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iHostname := self.Args("hostname").String()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iName) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByNameAndHostnameAndFieldName(offset, limit, iName,iHostname,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByNameAndHostnameAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByNameAndHostnameAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iHostname := self.Args("hostname").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iName) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByNameAndHostnameAndDefaultLangcode(offset, limit, iName,iHostname,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByNameAndHostnameAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByNameAndCreatedAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iCreated := self.Args("created").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iName) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByNameAndCreatedAndChanged(offset, limit, iName,iCreated,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByNameAndCreatedAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByNameAndCreatedAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iCreated := self.Args("created").MustInt()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iName) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByNameAndCreatedAndStatus(offset, limit, iName,iCreated,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByNameAndCreatedAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByNameAndCreatedAndThreadHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iCreated := self.Args("created").MustInt()
	iThread := self.Args("thread").String()

	if helper.IsHas(iName) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByNameAndCreatedAndThread(offset, limit, iName,iCreated,iThread)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByNameAndCreatedAndThread's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByNameAndCreatedAndEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iCreated := self.Args("created").MustInt()
	iEntityType := self.Args("entity_type").String()

	if helper.IsHas(iName) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByNameAndCreatedAndEntityType(offset, limit, iName,iCreated,iEntityType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByNameAndCreatedAndEntityType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByNameAndCreatedAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iCreated := self.Args("created").MustInt()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iName) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByNameAndCreatedAndFieldName(offset, limit, iName,iCreated,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByNameAndCreatedAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByNameAndCreatedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iCreated := self.Args("created").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iName) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByNameAndCreatedAndDefaultLangcode(offset, limit, iName,iCreated,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByNameAndCreatedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByNameAndChangedAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iChanged := self.Args("changed").MustInt()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iName) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByNameAndChangedAndStatus(offset, limit, iName,iChanged,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByNameAndChangedAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByNameAndChangedAndThreadHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iChanged := self.Args("changed").MustInt()
	iThread := self.Args("thread").String()

	if helper.IsHas(iName) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByNameAndChangedAndThread(offset, limit, iName,iChanged,iThread)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByNameAndChangedAndThread's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByNameAndChangedAndEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iChanged := self.Args("changed").MustInt()
	iEntityType := self.Args("entity_type").String()

	if helper.IsHas(iName) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByNameAndChangedAndEntityType(offset, limit, iName,iChanged,iEntityType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByNameAndChangedAndEntityType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByNameAndChangedAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iChanged := self.Args("changed").MustInt()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iName) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByNameAndChangedAndFieldName(offset, limit, iName,iChanged,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByNameAndChangedAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByNameAndChangedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iChanged := self.Args("changed").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iName) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByNameAndChangedAndDefaultLangcode(offset, limit, iName,iChanged,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByNameAndChangedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByNameAndStatusAndThreadHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iStatus := self.Args("status").MustInt()
	iThread := self.Args("thread").String()

	if helper.IsHas(iName) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByNameAndStatusAndThread(offset, limit, iName,iStatus,iThread)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByNameAndStatusAndThread's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByNameAndStatusAndEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iStatus := self.Args("status").MustInt()
	iEntityType := self.Args("entity_type").String()

	if helper.IsHas(iName) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByNameAndStatusAndEntityType(offset, limit, iName,iStatus,iEntityType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByNameAndStatusAndEntityType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByNameAndStatusAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iStatus := self.Args("status").MustInt()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iName) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByNameAndStatusAndFieldName(offset, limit, iName,iStatus,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByNameAndStatusAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByNameAndStatusAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iStatus := self.Args("status").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iName) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByNameAndStatusAndDefaultLangcode(offset, limit, iName,iStatus,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByNameAndStatusAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByNameAndThreadAndEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iThread := self.Args("thread").String()
	iEntityType := self.Args("entity_type").String()

	if helper.IsHas(iName) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByNameAndThreadAndEntityType(offset, limit, iName,iThread,iEntityType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByNameAndThreadAndEntityType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByNameAndThreadAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iThread := self.Args("thread").String()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iName) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByNameAndThreadAndFieldName(offset, limit, iName,iThread,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByNameAndThreadAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByNameAndThreadAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iThread := self.Args("thread").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iName) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByNameAndThreadAndDefaultLangcode(offset, limit, iName,iThread,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByNameAndThreadAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByNameAndEntityTypeAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iEntityType := self.Args("entity_type").String()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iName) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByNameAndEntityTypeAndFieldName(offset, limit, iName,iEntityType,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByNameAndEntityTypeAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByNameAndEntityTypeAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iEntityType := self.Args("entity_type").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iName) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByNameAndEntityTypeAndDefaultLangcode(offset, limit, iName,iEntityType,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByNameAndEntityTypeAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByNameAndFieldNameAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iFieldName := self.Args("field_name").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iName) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByNameAndFieldNameAndDefaultLangcode(offset, limit, iName,iFieldName,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByNameAndFieldNameAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByMailAndHomepageAndHostnameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMail := self.Args("mail").String()
	iHomepage := self.Args("homepage").String()
	iHostname := self.Args("hostname").String()

	if helper.IsHas(iMail) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByMailAndHomepageAndHostname(offset, limit, iMail,iHomepage,iHostname)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByMailAndHomepageAndHostname's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByMailAndHomepageAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMail := self.Args("mail").String()
	iHomepage := self.Args("homepage").String()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iMail) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByMailAndHomepageAndCreated(offset, limit, iMail,iHomepage,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByMailAndHomepageAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByMailAndHomepageAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMail := self.Args("mail").String()
	iHomepage := self.Args("homepage").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iMail) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByMailAndHomepageAndChanged(offset, limit, iMail,iHomepage,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByMailAndHomepageAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByMailAndHomepageAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMail := self.Args("mail").String()
	iHomepage := self.Args("homepage").String()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iMail) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByMailAndHomepageAndStatus(offset, limit, iMail,iHomepage,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByMailAndHomepageAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByMailAndHomepageAndThreadHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMail := self.Args("mail").String()
	iHomepage := self.Args("homepage").String()
	iThread := self.Args("thread").String()

	if helper.IsHas(iMail) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByMailAndHomepageAndThread(offset, limit, iMail,iHomepage,iThread)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByMailAndHomepageAndThread's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByMailAndHomepageAndEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMail := self.Args("mail").String()
	iHomepage := self.Args("homepage").String()
	iEntityType := self.Args("entity_type").String()

	if helper.IsHas(iMail) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByMailAndHomepageAndEntityType(offset, limit, iMail,iHomepage,iEntityType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByMailAndHomepageAndEntityType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByMailAndHomepageAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMail := self.Args("mail").String()
	iHomepage := self.Args("homepage").String()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iMail) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByMailAndHomepageAndFieldName(offset, limit, iMail,iHomepage,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByMailAndHomepageAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByMailAndHomepageAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMail := self.Args("mail").String()
	iHomepage := self.Args("homepage").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iMail) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByMailAndHomepageAndDefaultLangcode(offset, limit, iMail,iHomepage,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByMailAndHomepageAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByMailAndHostnameAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMail := self.Args("mail").String()
	iHostname := self.Args("hostname").String()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iMail) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByMailAndHostnameAndCreated(offset, limit, iMail,iHostname,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByMailAndHostnameAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByMailAndHostnameAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMail := self.Args("mail").String()
	iHostname := self.Args("hostname").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iMail) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByMailAndHostnameAndChanged(offset, limit, iMail,iHostname,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByMailAndHostnameAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByMailAndHostnameAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMail := self.Args("mail").String()
	iHostname := self.Args("hostname").String()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iMail) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByMailAndHostnameAndStatus(offset, limit, iMail,iHostname,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByMailAndHostnameAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByMailAndHostnameAndThreadHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMail := self.Args("mail").String()
	iHostname := self.Args("hostname").String()
	iThread := self.Args("thread").String()

	if helper.IsHas(iMail) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByMailAndHostnameAndThread(offset, limit, iMail,iHostname,iThread)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByMailAndHostnameAndThread's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByMailAndHostnameAndEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMail := self.Args("mail").String()
	iHostname := self.Args("hostname").String()
	iEntityType := self.Args("entity_type").String()

	if helper.IsHas(iMail) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByMailAndHostnameAndEntityType(offset, limit, iMail,iHostname,iEntityType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByMailAndHostnameAndEntityType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByMailAndHostnameAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMail := self.Args("mail").String()
	iHostname := self.Args("hostname").String()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iMail) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByMailAndHostnameAndFieldName(offset, limit, iMail,iHostname,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByMailAndHostnameAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByMailAndHostnameAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMail := self.Args("mail").String()
	iHostname := self.Args("hostname").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iMail) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByMailAndHostnameAndDefaultLangcode(offset, limit, iMail,iHostname,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByMailAndHostnameAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByMailAndCreatedAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMail := self.Args("mail").String()
	iCreated := self.Args("created").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iMail) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByMailAndCreatedAndChanged(offset, limit, iMail,iCreated,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByMailAndCreatedAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByMailAndCreatedAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMail := self.Args("mail").String()
	iCreated := self.Args("created").MustInt()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iMail) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByMailAndCreatedAndStatus(offset, limit, iMail,iCreated,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByMailAndCreatedAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByMailAndCreatedAndThreadHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMail := self.Args("mail").String()
	iCreated := self.Args("created").MustInt()
	iThread := self.Args("thread").String()

	if helper.IsHas(iMail) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByMailAndCreatedAndThread(offset, limit, iMail,iCreated,iThread)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByMailAndCreatedAndThread's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByMailAndCreatedAndEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMail := self.Args("mail").String()
	iCreated := self.Args("created").MustInt()
	iEntityType := self.Args("entity_type").String()

	if helper.IsHas(iMail) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByMailAndCreatedAndEntityType(offset, limit, iMail,iCreated,iEntityType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByMailAndCreatedAndEntityType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByMailAndCreatedAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMail := self.Args("mail").String()
	iCreated := self.Args("created").MustInt()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iMail) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByMailAndCreatedAndFieldName(offset, limit, iMail,iCreated,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByMailAndCreatedAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByMailAndCreatedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMail := self.Args("mail").String()
	iCreated := self.Args("created").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iMail) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByMailAndCreatedAndDefaultLangcode(offset, limit, iMail,iCreated,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByMailAndCreatedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByMailAndChangedAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMail := self.Args("mail").String()
	iChanged := self.Args("changed").MustInt()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iMail) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByMailAndChangedAndStatus(offset, limit, iMail,iChanged,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByMailAndChangedAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByMailAndChangedAndThreadHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMail := self.Args("mail").String()
	iChanged := self.Args("changed").MustInt()
	iThread := self.Args("thread").String()

	if helper.IsHas(iMail) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByMailAndChangedAndThread(offset, limit, iMail,iChanged,iThread)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByMailAndChangedAndThread's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByMailAndChangedAndEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMail := self.Args("mail").String()
	iChanged := self.Args("changed").MustInt()
	iEntityType := self.Args("entity_type").String()

	if helper.IsHas(iMail) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByMailAndChangedAndEntityType(offset, limit, iMail,iChanged,iEntityType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByMailAndChangedAndEntityType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByMailAndChangedAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMail := self.Args("mail").String()
	iChanged := self.Args("changed").MustInt()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iMail) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByMailAndChangedAndFieldName(offset, limit, iMail,iChanged,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByMailAndChangedAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByMailAndChangedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMail := self.Args("mail").String()
	iChanged := self.Args("changed").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iMail) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByMailAndChangedAndDefaultLangcode(offset, limit, iMail,iChanged,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByMailAndChangedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByMailAndStatusAndThreadHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMail := self.Args("mail").String()
	iStatus := self.Args("status").MustInt()
	iThread := self.Args("thread").String()

	if helper.IsHas(iMail) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByMailAndStatusAndThread(offset, limit, iMail,iStatus,iThread)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByMailAndStatusAndThread's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByMailAndStatusAndEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMail := self.Args("mail").String()
	iStatus := self.Args("status").MustInt()
	iEntityType := self.Args("entity_type").String()

	if helper.IsHas(iMail) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByMailAndStatusAndEntityType(offset, limit, iMail,iStatus,iEntityType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByMailAndStatusAndEntityType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByMailAndStatusAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMail := self.Args("mail").String()
	iStatus := self.Args("status").MustInt()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iMail) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByMailAndStatusAndFieldName(offset, limit, iMail,iStatus,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByMailAndStatusAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByMailAndStatusAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMail := self.Args("mail").String()
	iStatus := self.Args("status").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iMail) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByMailAndStatusAndDefaultLangcode(offset, limit, iMail,iStatus,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByMailAndStatusAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByMailAndThreadAndEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMail := self.Args("mail").String()
	iThread := self.Args("thread").String()
	iEntityType := self.Args("entity_type").String()

	if helper.IsHas(iMail) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByMailAndThreadAndEntityType(offset, limit, iMail,iThread,iEntityType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByMailAndThreadAndEntityType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByMailAndThreadAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMail := self.Args("mail").String()
	iThread := self.Args("thread").String()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iMail) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByMailAndThreadAndFieldName(offset, limit, iMail,iThread,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByMailAndThreadAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByMailAndThreadAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMail := self.Args("mail").String()
	iThread := self.Args("thread").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iMail) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByMailAndThreadAndDefaultLangcode(offset, limit, iMail,iThread,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByMailAndThreadAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByMailAndEntityTypeAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMail := self.Args("mail").String()
	iEntityType := self.Args("entity_type").String()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iMail) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByMailAndEntityTypeAndFieldName(offset, limit, iMail,iEntityType,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByMailAndEntityTypeAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByMailAndEntityTypeAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMail := self.Args("mail").String()
	iEntityType := self.Args("entity_type").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iMail) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByMailAndEntityTypeAndDefaultLangcode(offset, limit, iMail,iEntityType,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByMailAndEntityTypeAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByMailAndFieldNameAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMail := self.Args("mail").String()
	iFieldName := self.Args("field_name").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iMail) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByMailAndFieldNameAndDefaultLangcode(offset, limit, iMail,iFieldName,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByMailAndFieldNameAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByHomepageAndHostnameAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iHomepage := self.Args("homepage").String()
	iHostname := self.Args("hostname").String()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iHomepage) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByHomepageAndHostnameAndCreated(offset, limit, iHomepage,iHostname,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByHomepageAndHostnameAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByHomepageAndHostnameAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iHomepage := self.Args("homepage").String()
	iHostname := self.Args("hostname").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iHomepage) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByHomepageAndHostnameAndChanged(offset, limit, iHomepage,iHostname,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByHomepageAndHostnameAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByHomepageAndHostnameAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iHomepage := self.Args("homepage").String()
	iHostname := self.Args("hostname").String()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iHomepage) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByHomepageAndHostnameAndStatus(offset, limit, iHomepage,iHostname,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByHomepageAndHostnameAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByHomepageAndHostnameAndThreadHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iHomepage := self.Args("homepage").String()
	iHostname := self.Args("hostname").String()
	iThread := self.Args("thread").String()

	if helper.IsHas(iHomepage) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByHomepageAndHostnameAndThread(offset, limit, iHomepage,iHostname,iThread)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByHomepageAndHostnameAndThread's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByHomepageAndHostnameAndEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iHomepage := self.Args("homepage").String()
	iHostname := self.Args("hostname").String()
	iEntityType := self.Args("entity_type").String()

	if helper.IsHas(iHomepage) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByHomepageAndHostnameAndEntityType(offset, limit, iHomepage,iHostname,iEntityType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByHomepageAndHostnameAndEntityType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByHomepageAndHostnameAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iHomepage := self.Args("homepage").String()
	iHostname := self.Args("hostname").String()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iHomepage) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByHomepageAndHostnameAndFieldName(offset, limit, iHomepage,iHostname,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByHomepageAndHostnameAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByHomepageAndHostnameAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iHomepage := self.Args("homepage").String()
	iHostname := self.Args("hostname").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iHomepage) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByHomepageAndHostnameAndDefaultLangcode(offset, limit, iHomepage,iHostname,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByHomepageAndHostnameAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByHomepageAndCreatedAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iHomepage := self.Args("homepage").String()
	iCreated := self.Args("created").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iHomepage) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByHomepageAndCreatedAndChanged(offset, limit, iHomepage,iCreated,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByHomepageAndCreatedAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByHomepageAndCreatedAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iHomepage := self.Args("homepage").String()
	iCreated := self.Args("created").MustInt()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iHomepage) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByHomepageAndCreatedAndStatus(offset, limit, iHomepage,iCreated,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByHomepageAndCreatedAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByHomepageAndCreatedAndThreadHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iHomepage := self.Args("homepage").String()
	iCreated := self.Args("created").MustInt()
	iThread := self.Args("thread").String()

	if helper.IsHas(iHomepage) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByHomepageAndCreatedAndThread(offset, limit, iHomepage,iCreated,iThread)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByHomepageAndCreatedAndThread's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByHomepageAndCreatedAndEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iHomepage := self.Args("homepage").String()
	iCreated := self.Args("created").MustInt()
	iEntityType := self.Args("entity_type").String()

	if helper.IsHas(iHomepage) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByHomepageAndCreatedAndEntityType(offset, limit, iHomepage,iCreated,iEntityType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByHomepageAndCreatedAndEntityType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByHomepageAndCreatedAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iHomepage := self.Args("homepage").String()
	iCreated := self.Args("created").MustInt()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iHomepage) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByHomepageAndCreatedAndFieldName(offset, limit, iHomepage,iCreated,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByHomepageAndCreatedAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByHomepageAndCreatedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iHomepage := self.Args("homepage").String()
	iCreated := self.Args("created").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iHomepage) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByHomepageAndCreatedAndDefaultLangcode(offset, limit, iHomepage,iCreated,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByHomepageAndCreatedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByHomepageAndChangedAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iHomepage := self.Args("homepage").String()
	iChanged := self.Args("changed").MustInt()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iHomepage) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByHomepageAndChangedAndStatus(offset, limit, iHomepage,iChanged,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByHomepageAndChangedAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByHomepageAndChangedAndThreadHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iHomepage := self.Args("homepage").String()
	iChanged := self.Args("changed").MustInt()
	iThread := self.Args("thread").String()

	if helper.IsHas(iHomepage) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByHomepageAndChangedAndThread(offset, limit, iHomepage,iChanged,iThread)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByHomepageAndChangedAndThread's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByHomepageAndChangedAndEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iHomepage := self.Args("homepage").String()
	iChanged := self.Args("changed").MustInt()
	iEntityType := self.Args("entity_type").String()

	if helper.IsHas(iHomepage) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByHomepageAndChangedAndEntityType(offset, limit, iHomepage,iChanged,iEntityType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByHomepageAndChangedAndEntityType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByHomepageAndChangedAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iHomepage := self.Args("homepage").String()
	iChanged := self.Args("changed").MustInt()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iHomepage) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByHomepageAndChangedAndFieldName(offset, limit, iHomepage,iChanged,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByHomepageAndChangedAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByHomepageAndChangedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iHomepage := self.Args("homepage").String()
	iChanged := self.Args("changed").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iHomepage) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByHomepageAndChangedAndDefaultLangcode(offset, limit, iHomepage,iChanged,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByHomepageAndChangedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByHomepageAndStatusAndThreadHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iHomepage := self.Args("homepage").String()
	iStatus := self.Args("status").MustInt()
	iThread := self.Args("thread").String()

	if helper.IsHas(iHomepage) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByHomepageAndStatusAndThread(offset, limit, iHomepage,iStatus,iThread)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByHomepageAndStatusAndThread's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByHomepageAndStatusAndEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iHomepage := self.Args("homepage").String()
	iStatus := self.Args("status").MustInt()
	iEntityType := self.Args("entity_type").String()

	if helper.IsHas(iHomepage) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByHomepageAndStatusAndEntityType(offset, limit, iHomepage,iStatus,iEntityType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByHomepageAndStatusAndEntityType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByHomepageAndStatusAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iHomepage := self.Args("homepage").String()
	iStatus := self.Args("status").MustInt()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iHomepage) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByHomepageAndStatusAndFieldName(offset, limit, iHomepage,iStatus,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByHomepageAndStatusAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByHomepageAndStatusAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iHomepage := self.Args("homepage").String()
	iStatus := self.Args("status").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iHomepage) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByHomepageAndStatusAndDefaultLangcode(offset, limit, iHomepage,iStatus,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByHomepageAndStatusAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByHomepageAndThreadAndEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iHomepage := self.Args("homepage").String()
	iThread := self.Args("thread").String()
	iEntityType := self.Args("entity_type").String()

	if helper.IsHas(iHomepage) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByHomepageAndThreadAndEntityType(offset, limit, iHomepage,iThread,iEntityType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByHomepageAndThreadAndEntityType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByHomepageAndThreadAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iHomepage := self.Args("homepage").String()
	iThread := self.Args("thread").String()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iHomepage) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByHomepageAndThreadAndFieldName(offset, limit, iHomepage,iThread,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByHomepageAndThreadAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByHomepageAndThreadAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iHomepage := self.Args("homepage").String()
	iThread := self.Args("thread").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iHomepage) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByHomepageAndThreadAndDefaultLangcode(offset, limit, iHomepage,iThread,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByHomepageAndThreadAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByHomepageAndEntityTypeAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iHomepage := self.Args("homepage").String()
	iEntityType := self.Args("entity_type").String()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iHomepage) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByHomepageAndEntityTypeAndFieldName(offset, limit, iHomepage,iEntityType,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByHomepageAndEntityTypeAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByHomepageAndEntityTypeAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iHomepage := self.Args("homepage").String()
	iEntityType := self.Args("entity_type").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iHomepage) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByHomepageAndEntityTypeAndDefaultLangcode(offset, limit, iHomepage,iEntityType,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByHomepageAndEntityTypeAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByHomepageAndFieldNameAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iHomepage := self.Args("homepage").String()
	iFieldName := self.Args("field_name").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iHomepage) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByHomepageAndFieldNameAndDefaultLangcode(offset, limit, iHomepage,iFieldName,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByHomepageAndFieldNameAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByHostnameAndCreatedAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iHostname := self.Args("hostname").String()
	iCreated := self.Args("created").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iHostname) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByHostnameAndCreatedAndChanged(offset, limit, iHostname,iCreated,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByHostnameAndCreatedAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByHostnameAndCreatedAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iHostname := self.Args("hostname").String()
	iCreated := self.Args("created").MustInt()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iHostname) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByHostnameAndCreatedAndStatus(offset, limit, iHostname,iCreated,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByHostnameAndCreatedAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByHostnameAndCreatedAndThreadHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iHostname := self.Args("hostname").String()
	iCreated := self.Args("created").MustInt()
	iThread := self.Args("thread").String()

	if helper.IsHas(iHostname) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByHostnameAndCreatedAndThread(offset, limit, iHostname,iCreated,iThread)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByHostnameAndCreatedAndThread's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByHostnameAndCreatedAndEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iHostname := self.Args("hostname").String()
	iCreated := self.Args("created").MustInt()
	iEntityType := self.Args("entity_type").String()

	if helper.IsHas(iHostname) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByHostnameAndCreatedAndEntityType(offset, limit, iHostname,iCreated,iEntityType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByHostnameAndCreatedAndEntityType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByHostnameAndCreatedAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iHostname := self.Args("hostname").String()
	iCreated := self.Args("created").MustInt()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iHostname) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByHostnameAndCreatedAndFieldName(offset, limit, iHostname,iCreated,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByHostnameAndCreatedAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByHostnameAndCreatedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iHostname := self.Args("hostname").String()
	iCreated := self.Args("created").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iHostname) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByHostnameAndCreatedAndDefaultLangcode(offset, limit, iHostname,iCreated,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByHostnameAndCreatedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByHostnameAndChangedAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iHostname := self.Args("hostname").String()
	iChanged := self.Args("changed").MustInt()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iHostname) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByHostnameAndChangedAndStatus(offset, limit, iHostname,iChanged,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByHostnameAndChangedAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByHostnameAndChangedAndThreadHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iHostname := self.Args("hostname").String()
	iChanged := self.Args("changed").MustInt()
	iThread := self.Args("thread").String()

	if helper.IsHas(iHostname) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByHostnameAndChangedAndThread(offset, limit, iHostname,iChanged,iThread)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByHostnameAndChangedAndThread's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByHostnameAndChangedAndEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iHostname := self.Args("hostname").String()
	iChanged := self.Args("changed").MustInt()
	iEntityType := self.Args("entity_type").String()

	if helper.IsHas(iHostname) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByHostnameAndChangedAndEntityType(offset, limit, iHostname,iChanged,iEntityType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByHostnameAndChangedAndEntityType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByHostnameAndChangedAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iHostname := self.Args("hostname").String()
	iChanged := self.Args("changed").MustInt()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iHostname) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByHostnameAndChangedAndFieldName(offset, limit, iHostname,iChanged,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByHostnameAndChangedAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByHostnameAndChangedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iHostname := self.Args("hostname").String()
	iChanged := self.Args("changed").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iHostname) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByHostnameAndChangedAndDefaultLangcode(offset, limit, iHostname,iChanged,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByHostnameAndChangedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByHostnameAndStatusAndThreadHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iHostname := self.Args("hostname").String()
	iStatus := self.Args("status").MustInt()
	iThread := self.Args("thread").String()

	if helper.IsHas(iHostname) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByHostnameAndStatusAndThread(offset, limit, iHostname,iStatus,iThread)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByHostnameAndStatusAndThread's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByHostnameAndStatusAndEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iHostname := self.Args("hostname").String()
	iStatus := self.Args("status").MustInt()
	iEntityType := self.Args("entity_type").String()

	if helper.IsHas(iHostname) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByHostnameAndStatusAndEntityType(offset, limit, iHostname,iStatus,iEntityType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByHostnameAndStatusAndEntityType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByHostnameAndStatusAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iHostname := self.Args("hostname").String()
	iStatus := self.Args("status").MustInt()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iHostname) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByHostnameAndStatusAndFieldName(offset, limit, iHostname,iStatus,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByHostnameAndStatusAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByHostnameAndStatusAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iHostname := self.Args("hostname").String()
	iStatus := self.Args("status").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iHostname) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByHostnameAndStatusAndDefaultLangcode(offset, limit, iHostname,iStatus,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByHostnameAndStatusAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByHostnameAndThreadAndEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iHostname := self.Args("hostname").String()
	iThread := self.Args("thread").String()
	iEntityType := self.Args("entity_type").String()

	if helper.IsHas(iHostname) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByHostnameAndThreadAndEntityType(offset, limit, iHostname,iThread,iEntityType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByHostnameAndThreadAndEntityType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByHostnameAndThreadAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iHostname := self.Args("hostname").String()
	iThread := self.Args("thread").String()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iHostname) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByHostnameAndThreadAndFieldName(offset, limit, iHostname,iThread,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByHostnameAndThreadAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByHostnameAndThreadAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iHostname := self.Args("hostname").String()
	iThread := self.Args("thread").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iHostname) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByHostnameAndThreadAndDefaultLangcode(offset, limit, iHostname,iThread,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByHostnameAndThreadAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByHostnameAndEntityTypeAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iHostname := self.Args("hostname").String()
	iEntityType := self.Args("entity_type").String()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iHostname) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByHostnameAndEntityTypeAndFieldName(offset, limit, iHostname,iEntityType,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByHostnameAndEntityTypeAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByHostnameAndEntityTypeAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iHostname := self.Args("hostname").String()
	iEntityType := self.Args("entity_type").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iHostname) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByHostnameAndEntityTypeAndDefaultLangcode(offset, limit, iHostname,iEntityType,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByHostnameAndEntityTypeAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByHostnameAndFieldNameAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iHostname := self.Args("hostname").String()
	iFieldName := self.Args("field_name").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iHostname) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByHostnameAndFieldNameAndDefaultLangcode(offset, limit, iHostname,iFieldName,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByHostnameAndFieldNameAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCreatedAndChangedAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCreated := self.Args("created").MustInt()
	iChanged := self.Args("changed").MustInt()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iCreated) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCreatedAndChangedAndStatus(offset, limit, iCreated,iChanged,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCreatedAndChangedAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCreatedAndChangedAndThreadHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCreated := self.Args("created").MustInt()
	iChanged := self.Args("changed").MustInt()
	iThread := self.Args("thread").String()

	if helper.IsHas(iCreated) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCreatedAndChangedAndThread(offset, limit, iCreated,iChanged,iThread)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCreatedAndChangedAndThread's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCreatedAndChangedAndEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCreated := self.Args("created").MustInt()
	iChanged := self.Args("changed").MustInt()
	iEntityType := self.Args("entity_type").String()

	if helper.IsHas(iCreated) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCreatedAndChangedAndEntityType(offset, limit, iCreated,iChanged,iEntityType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCreatedAndChangedAndEntityType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCreatedAndChangedAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCreated := self.Args("created").MustInt()
	iChanged := self.Args("changed").MustInt()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iCreated) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCreatedAndChangedAndFieldName(offset, limit, iCreated,iChanged,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCreatedAndChangedAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCreatedAndChangedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCreated := self.Args("created").MustInt()
	iChanged := self.Args("changed").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iCreated) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCreatedAndChangedAndDefaultLangcode(offset, limit, iCreated,iChanged,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCreatedAndChangedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCreatedAndStatusAndThreadHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCreated := self.Args("created").MustInt()
	iStatus := self.Args("status").MustInt()
	iThread := self.Args("thread").String()

	if helper.IsHas(iCreated) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCreatedAndStatusAndThread(offset, limit, iCreated,iStatus,iThread)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCreatedAndStatusAndThread's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCreatedAndStatusAndEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCreated := self.Args("created").MustInt()
	iStatus := self.Args("status").MustInt()
	iEntityType := self.Args("entity_type").String()

	if helper.IsHas(iCreated) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCreatedAndStatusAndEntityType(offset, limit, iCreated,iStatus,iEntityType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCreatedAndStatusAndEntityType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCreatedAndStatusAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCreated := self.Args("created").MustInt()
	iStatus := self.Args("status").MustInt()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iCreated) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCreatedAndStatusAndFieldName(offset, limit, iCreated,iStatus,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCreatedAndStatusAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCreatedAndStatusAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCreated := self.Args("created").MustInt()
	iStatus := self.Args("status").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iCreated) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCreatedAndStatusAndDefaultLangcode(offset, limit, iCreated,iStatus,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCreatedAndStatusAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCreatedAndThreadAndEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCreated := self.Args("created").MustInt()
	iThread := self.Args("thread").String()
	iEntityType := self.Args("entity_type").String()

	if helper.IsHas(iCreated) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCreatedAndThreadAndEntityType(offset, limit, iCreated,iThread,iEntityType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCreatedAndThreadAndEntityType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCreatedAndThreadAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCreated := self.Args("created").MustInt()
	iThread := self.Args("thread").String()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iCreated) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCreatedAndThreadAndFieldName(offset, limit, iCreated,iThread,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCreatedAndThreadAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCreatedAndThreadAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCreated := self.Args("created").MustInt()
	iThread := self.Args("thread").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iCreated) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCreatedAndThreadAndDefaultLangcode(offset, limit, iCreated,iThread,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCreatedAndThreadAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCreatedAndEntityTypeAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCreated := self.Args("created").MustInt()
	iEntityType := self.Args("entity_type").String()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iCreated) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCreatedAndEntityTypeAndFieldName(offset, limit, iCreated,iEntityType,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCreatedAndEntityTypeAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCreatedAndEntityTypeAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCreated := self.Args("created").MustInt()
	iEntityType := self.Args("entity_type").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iCreated) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCreatedAndEntityTypeAndDefaultLangcode(offset, limit, iCreated,iEntityType,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCreatedAndEntityTypeAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCreatedAndFieldNameAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCreated := self.Args("created").MustInt()
	iFieldName := self.Args("field_name").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iCreated) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCreatedAndFieldNameAndDefaultLangcode(offset, limit, iCreated,iFieldName,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCreatedAndFieldNameAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByChangedAndStatusAndThreadHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iChanged := self.Args("changed").MustInt()
	iStatus := self.Args("status").MustInt()
	iThread := self.Args("thread").String()

	if helper.IsHas(iChanged) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByChangedAndStatusAndThread(offset, limit, iChanged,iStatus,iThread)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByChangedAndStatusAndThread's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByChangedAndStatusAndEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iChanged := self.Args("changed").MustInt()
	iStatus := self.Args("status").MustInt()
	iEntityType := self.Args("entity_type").String()

	if helper.IsHas(iChanged) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByChangedAndStatusAndEntityType(offset, limit, iChanged,iStatus,iEntityType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByChangedAndStatusAndEntityType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByChangedAndStatusAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iChanged := self.Args("changed").MustInt()
	iStatus := self.Args("status").MustInt()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iChanged) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByChangedAndStatusAndFieldName(offset, limit, iChanged,iStatus,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByChangedAndStatusAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByChangedAndStatusAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iChanged := self.Args("changed").MustInt()
	iStatus := self.Args("status").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iChanged) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByChangedAndStatusAndDefaultLangcode(offset, limit, iChanged,iStatus,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByChangedAndStatusAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByChangedAndThreadAndEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iChanged := self.Args("changed").MustInt()
	iThread := self.Args("thread").String()
	iEntityType := self.Args("entity_type").String()

	if helper.IsHas(iChanged) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByChangedAndThreadAndEntityType(offset, limit, iChanged,iThread,iEntityType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByChangedAndThreadAndEntityType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByChangedAndThreadAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iChanged := self.Args("changed").MustInt()
	iThread := self.Args("thread").String()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iChanged) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByChangedAndThreadAndFieldName(offset, limit, iChanged,iThread,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByChangedAndThreadAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByChangedAndThreadAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iChanged := self.Args("changed").MustInt()
	iThread := self.Args("thread").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iChanged) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByChangedAndThreadAndDefaultLangcode(offset, limit, iChanged,iThread,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByChangedAndThreadAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByChangedAndEntityTypeAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iChanged := self.Args("changed").MustInt()
	iEntityType := self.Args("entity_type").String()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iChanged) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByChangedAndEntityTypeAndFieldName(offset, limit, iChanged,iEntityType,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByChangedAndEntityTypeAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByChangedAndEntityTypeAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iChanged := self.Args("changed").MustInt()
	iEntityType := self.Args("entity_type").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iChanged) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByChangedAndEntityTypeAndDefaultLangcode(offset, limit, iChanged,iEntityType,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByChangedAndEntityTypeAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByChangedAndFieldNameAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iChanged := self.Args("changed").MustInt()
	iFieldName := self.Args("field_name").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iChanged) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByChangedAndFieldNameAndDefaultLangcode(offset, limit, iChanged,iFieldName,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByChangedAndFieldNameAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByStatusAndThreadAndEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iStatus := self.Args("status").MustInt()
	iThread := self.Args("thread").String()
	iEntityType := self.Args("entity_type").String()

	if helper.IsHas(iStatus) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByStatusAndThreadAndEntityType(offset, limit, iStatus,iThread,iEntityType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByStatusAndThreadAndEntityType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByStatusAndThreadAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iStatus := self.Args("status").MustInt()
	iThread := self.Args("thread").String()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iStatus) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByStatusAndThreadAndFieldName(offset, limit, iStatus,iThread,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByStatusAndThreadAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByStatusAndThreadAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iStatus := self.Args("status").MustInt()
	iThread := self.Args("thread").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iStatus) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByStatusAndThreadAndDefaultLangcode(offset, limit, iStatus,iThread,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByStatusAndThreadAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByStatusAndEntityTypeAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iStatus := self.Args("status").MustInt()
	iEntityType := self.Args("entity_type").String()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iStatus) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByStatusAndEntityTypeAndFieldName(offset, limit, iStatus,iEntityType,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByStatusAndEntityTypeAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByStatusAndEntityTypeAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iStatus := self.Args("status").MustInt()
	iEntityType := self.Args("entity_type").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iStatus) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByStatusAndEntityTypeAndDefaultLangcode(offset, limit, iStatus,iEntityType,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByStatusAndEntityTypeAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByStatusAndFieldNameAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iStatus := self.Args("status").MustInt()
	iFieldName := self.Args("field_name").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iStatus) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByStatusAndFieldNameAndDefaultLangcode(offset, limit, iStatus,iFieldName,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByStatusAndFieldNameAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByThreadAndEntityTypeAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iThread := self.Args("thread").String()
	iEntityType := self.Args("entity_type").String()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iThread) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByThreadAndEntityTypeAndFieldName(offset, limit, iThread,iEntityType,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByThreadAndEntityTypeAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByThreadAndEntityTypeAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iThread := self.Args("thread").String()
	iEntityType := self.Args("entity_type").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iThread) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByThreadAndEntityTypeAndDefaultLangcode(offset, limit, iThread,iEntityType,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByThreadAndEntityTypeAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByThreadAndFieldNameAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iThread := self.Args("thread").String()
	iFieldName := self.Args("field_name").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iThread) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByThreadAndFieldNameAndDefaultLangcode(offset, limit, iThread,iFieldName,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByThreadAndFieldNameAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByEntityTypeAndFieldNameAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityType := self.Args("entity_type").String()
	iFieldName := self.Args("field_name").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iEntityType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByEntityTypeAndFieldNameAndDefaultLangcode(offset, limit, iEntityType,iFieldName,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByEntityTypeAndFieldNameAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndCommentTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iCommentType := self.Args("comment_type").String()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndCommentType(offset, limit, iCid,iCommentType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndCommentType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndLangcode(offset, limit, iCid,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndPidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iPid := self.Args("pid").MustInt()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndPid(offset, limit, iCid,iPid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndPid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iEntityId := self.Args("entity_id").MustInt()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndEntityId(offset, limit, iCid,iEntityId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndEntityId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndSubjectHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iSubject := self.Args("subject").String()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndSubject(offset, limit, iCid,iSubject)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndSubject's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iUid := self.Args("uid").MustInt()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndUid(offset, limit, iCid,iUid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iName := self.Args("name").String()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndName(offset, limit, iCid,iName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndMailHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iMail := self.Args("mail").String()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndMail(offset, limit, iCid,iMail)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndMail's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndHomepageHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iHomepage := self.Args("homepage").String()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndHomepage(offset, limit, iCid,iHomepage)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndHomepage's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndHostnameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iHostname := self.Args("hostname").String()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndHostname(offset, limit, iCid,iHostname)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndHostname's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndCreated(offset, limit, iCid,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndChanged(offset, limit, iCid,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndStatus(offset, limit, iCid,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndThreadHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iThread := self.Args("thread").String()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndThread(offset, limit, iCid,iThread)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndThread's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iEntityType := self.Args("entity_type").String()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndEntityType(offset, limit, iCid,iEntityType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndEntityType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndFieldName(offset, limit, iCid,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCidAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCid := self.Args("cid").MustInt64()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCidAndDefaultLangcode(offset, limit, iCid,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCidAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndLangcode(offset, limit, iCommentType,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndPidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iPid := self.Args("pid").MustInt()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndPid(offset, limit, iCommentType,iPid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndPid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iEntityId := self.Args("entity_id").MustInt()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndEntityId(offset, limit, iCommentType,iEntityId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndEntityId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndSubjectHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iSubject := self.Args("subject").String()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndSubject(offset, limit, iCommentType,iSubject)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndSubject's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iUid := self.Args("uid").MustInt()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndUid(offset, limit, iCommentType,iUid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iName := self.Args("name").String()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndName(offset, limit, iCommentType,iName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndMailHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iMail := self.Args("mail").String()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndMail(offset, limit, iCommentType,iMail)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndMail's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndHomepageHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iHomepage := self.Args("homepage").String()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndHomepage(offset, limit, iCommentType,iHomepage)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndHomepage's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndHostnameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iHostname := self.Args("hostname").String()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndHostname(offset, limit, iCommentType,iHostname)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndHostname's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndCreated(offset, limit, iCommentType,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndChanged(offset, limit, iCommentType,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndStatus(offset, limit, iCommentType,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndThreadHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iThread := self.Args("thread").String()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndThread(offset, limit, iCommentType,iThread)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndThread's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iEntityType := self.Args("entity_type").String()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndEntityType(offset, limit, iCommentType,iEntityType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndEntityType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndFieldName(offset, limit, iCommentType,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCommentTypeAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCommentTypeAndDefaultLangcode(offset, limit, iCommentType,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCommentTypeAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByLangcodeAndPidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iPid := self.Args("pid").MustInt()

	if helper.IsHas(iLangcode) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByLangcodeAndPid(offset, limit, iLangcode,iPid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByLangcodeAndPid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByLangcodeAndEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iEntityId := self.Args("entity_id").MustInt()

	if helper.IsHas(iLangcode) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByLangcodeAndEntityId(offset, limit, iLangcode,iEntityId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByLangcodeAndEntityId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByLangcodeAndSubjectHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iSubject := self.Args("subject").String()

	if helper.IsHas(iLangcode) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByLangcodeAndSubject(offset, limit, iLangcode,iSubject)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByLangcodeAndSubject's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByLangcodeAndUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iUid := self.Args("uid").MustInt()

	if helper.IsHas(iLangcode) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByLangcodeAndUid(offset, limit, iLangcode,iUid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByLangcodeAndUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByLangcodeAndNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iName := self.Args("name").String()

	if helper.IsHas(iLangcode) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByLangcodeAndName(offset, limit, iLangcode,iName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByLangcodeAndName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByLangcodeAndMailHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iMail := self.Args("mail").String()

	if helper.IsHas(iLangcode) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByLangcodeAndMail(offset, limit, iLangcode,iMail)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByLangcodeAndMail's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByLangcodeAndHomepageHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iHomepage := self.Args("homepage").String()

	if helper.IsHas(iLangcode) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByLangcodeAndHomepage(offset, limit, iLangcode,iHomepage)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByLangcodeAndHomepage's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByLangcodeAndHostnameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iHostname := self.Args("hostname").String()

	if helper.IsHas(iLangcode) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByLangcodeAndHostname(offset, limit, iLangcode,iHostname)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByLangcodeAndHostname's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByLangcodeAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iLangcode) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByLangcodeAndCreated(offset, limit, iLangcode,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByLangcodeAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByLangcodeAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iLangcode) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByLangcodeAndChanged(offset, limit, iLangcode,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByLangcodeAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByLangcodeAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iLangcode) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByLangcodeAndStatus(offset, limit, iLangcode,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByLangcodeAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByLangcodeAndThreadHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iThread := self.Args("thread").String()

	if helper.IsHas(iLangcode) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByLangcodeAndThread(offset, limit, iLangcode,iThread)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByLangcodeAndThread's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByLangcodeAndEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iEntityType := self.Args("entity_type").String()

	if helper.IsHas(iLangcode) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByLangcodeAndEntityType(offset, limit, iLangcode,iEntityType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByLangcodeAndEntityType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByLangcodeAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iLangcode) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByLangcodeAndFieldName(offset, limit, iLangcode,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByLangcodeAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByLangcodeAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iLangcode) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByLangcodeAndDefaultLangcode(offset, limit, iLangcode,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByLangcodeAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByPidAndEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPid := self.Args("pid").MustInt()
	iEntityId := self.Args("entity_id").MustInt()

	if helper.IsHas(iPid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByPidAndEntityId(offset, limit, iPid,iEntityId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByPidAndEntityId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByPidAndSubjectHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPid := self.Args("pid").MustInt()
	iSubject := self.Args("subject").String()

	if helper.IsHas(iPid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByPidAndSubject(offset, limit, iPid,iSubject)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByPidAndSubject's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByPidAndUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPid := self.Args("pid").MustInt()
	iUid := self.Args("uid").MustInt()

	if helper.IsHas(iPid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByPidAndUid(offset, limit, iPid,iUid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByPidAndUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByPidAndNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPid := self.Args("pid").MustInt()
	iName := self.Args("name").String()

	if helper.IsHas(iPid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByPidAndName(offset, limit, iPid,iName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByPidAndName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByPidAndMailHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPid := self.Args("pid").MustInt()
	iMail := self.Args("mail").String()

	if helper.IsHas(iPid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByPidAndMail(offset, limit, iPid,iMail)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByPidAndMail's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByPidAndHomepageHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPid := self.Args("pid").MustInt()
	iHomepage := self.Args("homepage").String()

	if helper.IsHas(iPid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByPidAndHomepage(offset, limit, iPid,iHomepage)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByPidAndHomepage's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByPidAndHostnameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPid := self.Args("pid").MustInt()
	iHostname := self.Args("hostname").String()

	if helper.IsHas(iPid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByPidAndHostname(offset, limit, iPid,iHostname)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByPidAndHostname's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByPidAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPid := self.Args("pid").MustInt()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iPid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByPidAndCreated(offset, limit, iPid,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByPidAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByPidAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPid := self.Args("pid").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iPid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByPidAndChanged(offset, limit, iPid,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByPidAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByPidAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPid := self.Args("pid").MustInt()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iPid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByPidAndStatus(offset, limit, iPid,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByPidAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByPidAndThreadHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPid := self.Args("pid").MustInt()
	iThread := self.Args("thread").String()

	if helper.IsHas(iPid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByPidAndThread(offset, limit, iPid,iThread)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByPidAndThread's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByPidAndEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPid := self.Args("pid").MustInt()
	iEntityType := self.Args("entity_type").String()

	if helper.IsHas(iPid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByPidAndEntityType(offset, limit, iPid,iEntityType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByPidAndEntityType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByPidAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPid := self.Args("pid").MustInt()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iPid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByPidAndFieldName(offset, limit, iPid,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByPidAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByPidAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPid := self.Args("pid").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iPid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByPidAndDefaultLangcode(offset, limit, iPid,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByPidAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByEntityIdAndSubjectHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iSubject := self.Args("subject").String()

	if helper.IsHas(iEntityId) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByEntityIdAndSubject(offset, limit, iEntityId,iSubject)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByEntityIdAndSubject's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByEntityIdAndUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iUid := self.Args("uid").MustInt()

	if helper.IsHas(iEntityId) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByEntityIdAndUid(offset, limit, iEntityId,iUid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByEntityIdAndUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByEntityIdAndNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iName := self.Args("name").String()

	if helper.IsHas(iEntityId) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByEntityIdAndName(offset, limit, iEntityId,iName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByEntityIdAndName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByEntityIdAndMailHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iMail := self.Args("mail").String()

	if helper.IsHas(iEntityId) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByEntityIdAndMail(offset, limit, iEntityId,iMail)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByEntityIdAndMail's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByEntityIdAndHomepageHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iHomepage := self.Args("homepage").String()

	if helper.IsHas(iEntityId) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByEntityIdAndHomepage(offset, limit, iEntityId,iHomepage)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByEntityIdAndHomepage's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByEntityIdAndHostnameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iHostname := self.Args("hostname").String()

	if helper.IsHas(iEntityId) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByEntityIdAndHostname(offset, limit, iEntityId,iHostname)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByEntityIdAndHostname's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByEntityIdAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iEntityId) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByEntityIdAndCreated(offset, limit, iEntityId,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByEntityIdAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByEntityIdAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iEntityId) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByEntityIdAndChanged(offset, limit, iEntityId,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByEntityIdAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByEntityIdAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iEntityId) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByEntityIdAndStatus(offset, limit, iEntityId,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByEntityIdAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByEntityIdAndThreadHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iThread := self.Args("thread").String()

	if helper.IsHas(iEntityId) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByEntityIdAndThread(offset, limit, iEntityId,iThread)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByEntityIdAndThread's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByEntityIdAndEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iEntityType := self.Args("entity_type").String()

	if helper.IsHas(iEntityId) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByEntityIdAndEntityType(offset, limit, iEntityId,iEntityType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByEntityIdAndEntityType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByEntityIdAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iEntityId) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByEntityIdAndFieldName(offset, limit, iEntityId,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByEntityIdAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByEntityIdAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityId := self.Args("entity_id").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iEntityId) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByEntityIdAndDefaultLangcode(offset, limit, iEntityId,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByEntityIdAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasBySubjectAndUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSubject := self.Args("subject").String()
	iUid := self.Args("uid").MustInt()

	if helper.IsHas(iSubject) {
		_CommentFieldData, _error := model.GetCommentFieldDatasBySubjectAndUid(offset, limit, iSubject,iUid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasBySubjectAndUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasBySubjectAndNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSubject := self.Args("subject").String()
	iName := self.Args("name").String()

	if helper.IsHas(iSubject) {
		_CommentFieldData, _error := model.GetCommentFieldDatasBySubjectAndName(offset, limit, iSubject,iName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasBySubjectAndName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasBySubjectAndMailHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSubject := self.Args("subject").String()
	iMail := self.Args("mail").String()

	if helper.IsHas(iSubject) {
		_CommentFieldData, _error := model.GetCommentFieldDatasBySubjectAndMail(offset, limit, iSubject,iMail)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasBySubjectAndMail's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasBySubjectAndHomepageHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSubject := self.Args("subject").String()
	iHomepage := self.Args("homepage").String()

	if helper.IsHas(iSubject) {
		_CommentFieldData, _error := model.GetCommentFieldDatasBySubjectAndHomepage(offset, limit, iSubject,iHomepage)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasBySubjectAndHomepage's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasBySubjectAndHostnameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSubject := self.Args("subject").String()
	iHostname := self.Args("hostname").String()

	if helper.IsHas(iSubject) {
		_CommentFieldData, _error := model.GetCommentFieldDatasBySubjectAndHostname(offset, limit, iSubject,iHostname)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasBySubjectAndHostname's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasBySubjectAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSubject := self.Args("subject").String()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iSubject) {
		_CommentFieldData, _error := model.GetCommentFieldDatasBySubjectAndCreated(offset, limit, iSubject,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasBySubjectAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasBySubjectAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSubject := self.Args("subject").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iSubject) {
		_CommentFieldData, _error := model.GetCommentFieldDatasBySubjectAndChanged(offset, limit, iSubject,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasBySubjectAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasBySubjectAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSubject := self.Args("subject").String()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iSubject) {
		_CommentFieldData, _error := model.GetCommentFieldDatasBySubjectAndStatus(offset, limit, iSubject,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasBySubjectAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasBySubjectAndThreadHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSubject := self.Args("subject").String()
	iThread := self.Args("thread").String()

	if helper.IsHas(iSubject) {
		_CommentFieldData, _error := model.GetCommentFieldDatasBySubjectAndThread(offset, limit, iSubject,iThread)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasBySubjectAndThread's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasBySubjectAndEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSubject := self.Args("subject").String()
	iEntityType := self.Args("entity_type").String()

	if helper.IsHas(iSubject) {
		_CommentFieldData, _error := model.GetCommentFieldDatasBySubjectAndEntityType(offset, limit, iSubject,iEntityType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasBySubjectAndEntityType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasBySubjectAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSubject := self.Args("subject").String()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iSubject) {
		_CommentFieldData, _error := model.GetCommentFieldDatasBySubjectAndFieldName(offset, limit, iSubject,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasBySubjectAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasBySubjectAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSubject := self.Args("subject").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iSubject) {
		_CommentFieldData, _error := model.GetCommentFieldDatasBySubjectAndDefaultLangcode(offset, limit, iSubject,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasBySubjectAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByUidAndNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iName := self.Args("name").String()

	if helper.IsHas(iUid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByUidAndName(offset, limit, iUid,iName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByUidAndName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByUidAndMailHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iMail := self.Args("mail").String()

	if helper.IsHas(iUid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByUidAndMail(offset, limit, iUid,iMail)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByUidAndMail's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByUidAndHomepageHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iHomepage := self.Args("homepage").String()

	if helper.IsHas(iUid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByUidAndHomepage(offset, limit, iUid,iHomepage)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByUidAndHomepage's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByUidAndHostnameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iHostname := self.Args("hostname").String()

	if helper.IsHas(iUid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByUidAndHostname(offset, limit, iUid,iHostname)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByUidAndHostname's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByUidAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iUid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByUidAndCreated(offset, limit, iUid,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByUidAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByUidAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iUid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByUidAndChanged(offset, limit, iUid,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByUidAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByUidAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iUid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByUidAndStatus(offset, limit, iUid,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByUidAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByUidAndThreadHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iThread := self.Args("thread").String()

	if helper.IsHas(iUid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByUidAndThread(offset, limit, iUid,iThread)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByUidAndThread's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByUidAndEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iEntityType := self.Args("entity_type").String()

	if helper.IsHas(iUid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByUidAndEntityType(offset, limit, iUid,iEntityType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByUidAndEntityType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByUidAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iUid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByUidAndFieldName(offset, limit, iUid,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByUidAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByUidAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUid := self.Args("uid").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iUid) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByUidAndDefaultLangcode(offset, limit, iUid,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByUidAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByNameAndMailHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iMail := self.Args("mail").String()

	if helper.IsHas(iName) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByNameAndMail(offset, limit, iName,iMail)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByNameAndMail's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByNameAndHomepageHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iHomepage := self.Args("homepage").String()

	if helper.IsHas(iName) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByNameAndHomepage(offset, limit, iName,iHomepage)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByNameAndHomepage's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByNameAndHostnameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iHostname := self.Args("hostname").String()

	if helper.IsHas(iName) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByNameAndHostname(offset, limit, iName,iHostname)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByNameAndHostname's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByNameAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iName) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByNameAndCreated(offset, limit, iName,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByNameAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByNameAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iName) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByNameAndChanged(offset, limit, iName,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByNameAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByNameAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iName) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByNameAndStatus(offset, limit, iName,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByNameAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByNameAndThreadHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iThread := self.Args("thread").String()

	if helper.IsHas(iName) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByNameAndThread(offset, limit, iName,iThread)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByNameAndThread's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByNameAndEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iEntityType := self.Args("entity_type").String()

	if helper.IsHas(iName) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByNameAndEntityType(offset, limit, iName,iEntityType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByNameAndEntityType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByNameAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iName) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByNameAndFieldName(offset, limit, iName,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByNameAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByNameAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iName) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByNameAndDefaultLangcode(offset, limit, iName,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByNameAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByMailAndHomepageHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMail := self.Args("mail").String()
	iHomepage := self.Args("homepage").String()

	if helper.IsHas(iMail) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByMailAndHomepage(offset, limit, iMail,iHomepage)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByMailAndHomepage's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByMailAndHostnameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMail := self.Args("mail").String()
	iHostname := self.Args("hostname").String()

	if helper.IsHas(iMail) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByMailAndHostname(offset, limit, iMail,iHostname)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByMailAndHostname's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByMailAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMail := self.Args("mail").String()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iMail) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByMailAndCreated(offset, limit, iMail,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByMailAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByMailAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMail := self.Args("mail").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iMail) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByMailAndChanged(offset, limit, iMail,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByMailAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByMailAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMail := self.Args("mail").String()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iMail) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByMailAndStatus(offset, limit, iMail,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByMailAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByMailAndThreadHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMail := self.Args("mail").String()
	iThread := self.Args("thread").String()

	if helper.IsHas(iMail) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByMailAndThread(offset, limit, iMail,iThread)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByMailAndThread's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByMailAndEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMail := self.Args("mail").String()
	iEntityType := self.Args("entity_type").String()

	if helper.IsHas(iMail) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByMailAndEntityType(offset, limit, iMail,iEntityType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByMailAndEntityType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByMailAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMail := self.Args("mail").String()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iMail) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByMailAndFieldName(offset, limit, iMail,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByMailAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByMailAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMail := self.Args("mail").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iMail) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByMailAndDefaultLangcode(offset, limit, iMail,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByMailAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByHomepageAndHostnameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iHomepage := self.Args("homepage").String()
	iHostname := self.Args("hostname").String()

	if helper.IsHas(iHomepage) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByHomepageAndHostname(offset, limit, iHomepage,iHostname)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByHomepageAndHostname's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByHomepageAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iHomepage := self.Args("homepage").String()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iHomepage) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByHomepageAndCreated(offset, limit, iHomepage,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByHomepageAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByHomepageAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iHomepage := self.Args("homepage").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iHomepage) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByHomepageAndChanged(offset, limit, iHomepage,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByHomepageAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByHomepageAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iHomepage := self.Args("homepage").String()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iHomepage) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByHomepageAndStatus(offset, limit, iHomepage,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByHomepageAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByHomepageAndThreadHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iHomepage := self.Args("homepage").String()
	iThread := self.Args("thread").String()

	if helper.IsHas(iHomepage) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByHomepageAndThread(offset, limit, iHomepage,iThread)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByHomepageAndThread's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByHomepageAndEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iHomepage := self.Args("homepage").String()
	iEntityType := self.Args("entity_type").String()

	if helper.IsHas(iHomepage) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByHomepageAndEntityType(offset, limit, iHomepage,iEntityType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByHomepageAndEntityType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByHomepageAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iHomepage := self.Args("homepage").String()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iHomepage) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByHomepageAndFieldName(offset, limit, iHomepage,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByHomepageAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByHomepageAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iHomepage := self.Args("homepage").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iHomepage) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByHomepageAndDefaultLangcode(offset, limit, iHomepage,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByHomepageAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByHostnameAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iHostname := self.Args("hostname").String()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iHostname) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByHostnameAndCreated(offset, limit, iHostname,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByHostnameAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByHostnameAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iHostname := self.Args("hostname").String()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iHostname) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByHostnameAndChanged(offset, limit, iHostname,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByHostnameAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByHostnameAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iHostname := self.Args("hostname").String()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iHostname) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByHostnameAndStatus(offset, limit, iHostname,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByHostnameAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByHostnameAndThreadHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iHostname := self.Args("hostname").String()
	iThread := self.Args("thread").String()

	if helper.IsHas(iHostname) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByHostnameAndThread(offset, limit, iHostname,iThread)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByHostnameAndThread's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByHostnameAndEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iHostname := self.Args("hostname").String()
	iEntityType := self.Args("entity_type").String()

	if helper.IsHas(iHostname) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByHostnameAndEntityType(offset, limit, iHostname,iEntityType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByHostnameAndEntityType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByHostnameAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iHostname := self.Args("hostname").String()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iHostname) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByHostnameAndFieldName(offset, limit, iHostname,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByHostnameAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByHostnameAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iHostname := self.Args("hostname").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iHostname) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByHostnameAndDefaultLangcode(offset, limit, iHostname,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByHostnameAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCreatedAndChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCreated := self.Args("created").MustInt()
	iChanged := self.Args("changed").MustInt()

	if helper.IsHas(iCreated) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCreatedAndChanged(offset, limit, iCreated,iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCreatedAndChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCreatedAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCreated := self.Args("created").MustInt()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iCreated) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCreatedAndStatus(offset, limit, iCreated,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCreatedAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCreatedAndThreadHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCreated := self.Args("created").MustInt()
	iThread := self.Args("thread").String()

	if helper.IsHas(iCreated) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCreatedAndThread(offset, limit, iCreated,iThread)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCreatedAndThread's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCreatedAndEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCreated := self.Args("created").MustInt()
	iEntityType := self.Args("entity_type").String()

	if helper.IsHas(iCreated) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCreatedAndEntityType(offset, limit, iCreated,iEntityType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCreatedAndEntityType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCreatedAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCreated := self.Args("created").MustInt()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iCreated) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCreatedAndFieldName(offset, limit, iCreated,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCreatedAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByCreatedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCreated := self.Args("created").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iCreated) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByCreatedAndDefaultLangcode(offset, limit, iCreated,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByCreatedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByChangedAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iChanged := self.Args("changed").MustInt()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iChanged) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByChangedAndStatus(offset, limit, iChanged,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByChangedAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByChangedAndThreadHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iChanged := self.Args("changed").MustInt()
	iThread := self.Args("thread").String()

	if helper.IsHas(iChanged) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByChangedAndThread(offset, limit, iChanged,iThread)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByChangedAndThread's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByChangedAndEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iChanged := self.Args("changed").MustInt()
	iEntityType := self.Args("entity_type").String()

	if helper.IsHas(iChanged) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByChangedAndEntityType(offset, limit, iChanged,iEntityType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByChangedAndEntityType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByChangedAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iChanged := self.Args("changed").MustInt()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iChanged) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByChangedAndFieldName(offset, limit, iChanged,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByChangedAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByChangedAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iChanged := self.Args("changed").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iChanged) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByChangedAndDefaultLangcode(offset, limit, iChanged,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByChangedAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByStatusAndThreadHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iStatus := self.Args("status").MustInt()
	iThread := self.Args("thread").String()

	if helper.IsHas(iStatus) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByStatusAndThread(offset, limit, iStatus,iThread)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByStatusAndThread's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByStatusAndEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iStatus := self.Args("status").MustInt()
	iEntityType := self.Args("entity_type").String()

	if helper.IsHas(iStatus) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByStatusAndEntityType(offset, limit, iStatus,iEntityType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByStatusAndEntityType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByStatusAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iStatus := self.Args("status").MustInt()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iStatus) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByStatusAndFieldName(offset, limit, iStatus,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByStatusAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByStatusAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iStatus := self.Args("status").MustInt()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iStatus) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByStatusAndDefaultLangcode(offset, limit, iStatus,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByStatusAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByThreadAndEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iThread := self.Args("thread").String()
	iEntityType := self.Args("entity_type").String()

	if helper.IsHas(iThread) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByThreadAndEntityType(offset, limit, iThread,iEntityType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByThreadAndEntityType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByThreadAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iThread := self.Args("thread").String()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iThread) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByThreadAndFieldName(offset, limit, iThread,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByThreadAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByThreadAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iThread := self.Args("thread").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iThread) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByThreadAndDefaultLangcode(offset, limit, iThread,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByThreadAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByEntityTypeAndFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityType := self.Args("entity_type").String()
	iFieldName := self.Args("field_name").String()

	if helper.IsHas(iEntityType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByEntityTypeAndFieldName(offset, limit, iEntityType,iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByEntityTypeAndFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByEntityTypeAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEntityType := self.Args("entity_type").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iEntityType) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByEntityTypeAndDefaultLangcode(offset, limit, iEntityType,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByEntityTypeAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasByFieldNameAndDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFieldName := self.Args("field_name").String()
	iDefaultLangcode := self.Args("default_langcode").MustInt()

	if helper.IsHas(iFieldName) {
		_CommentFieldData, _error := model.GetCommentFieldDatasByFieldNameAndDefaultLangcode(offset, limit, iFieldName,iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatasByFieldNameAndDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDatasHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	if (offset > 0) && (len(field) > 0) {
		_CommentFieldData, _error := model.GetCommentFieldDatas(offset, limit, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDatas' args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasCommentFieldDataViaCidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iCid := self.Args("cid").MustInt64()
	if helper.IsHas(iCid) {
		_CommentFieldData := model.HasCommentFieldDataViaCid(iCid)
		var m = map[string]interface{}{}
		m["comment_field_data"] = _CommentFieldData
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasCommentFieldDataViaCid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasCommentFieldDataViaCommentTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iCommentType := self.Args("comment_type").String()
	if helper.IsHas(iCommentType) {
		_CommentFieldData := model.HasCommentFieldDataViaCommentType(iCommentType)
		var m = map[string]interface{}{}
		m["comment_field_data"] = _CommentFieldData
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasCommentFieldDataViaCommentType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasCommentFieldDataViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iLangcode := self.Args("langcode").String()
	if helper.IsHas(iLangcode) {
		_CommentFieldData := model.HasCommentFieldDataViaLangcode(iLangcode)
		var m = map[string]interface{}{}
		m["comment_field_data"] = _CommentFieldData
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasCommentFieldDataViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasCommentFieldDataViaPidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iPid := self.Args("pid").MustInt()
	if helper.IsHas(iPid) {
		_CommentFieldData := model.HasCommentFieldDataViaPid(iPid)
		var m = map[string]interface{}{}
		m["comment_field_data"] = _CommentFieldData
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasCommentFieldDataViaPid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasCommentFieldDataViaEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iEntityId := self.Args("entity_id").MustInt()
	if helper.IsHas(iEntityId) {
		_CommentFieldData := model.HasCommentFieldDataViaEntityId(iEntityId)
		var m = map[string]interface{}{}
		m["comment_field_data"] = _CommentFieldData
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasCommentFieldDataViaEntityId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasCommentFieldDataViaSubjectHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iSubject := self.Args("subject").String()
	if helper.IsHas(iSubject) {
		_CommentFieldData := model.HasCommentFieldDataViaSubject(iSubject)
		var m = map[string]interface{}{}
		m["comment_field_data"] = _CommentFieldData
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasCommentFieldDataViaSubject's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasCommentFieldDataViaUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iUid := self.Args("uid").MustInt()
	if helper.IsHas(iUid) {
		_CommentFieldData := model.HasCommentFieldDataViaUid(iUid)
		var m = map[string]interface{}{}
		m["comment_field_data"] = _CommentFieldData
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasCommentFieldDataViaUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasCommentFieldDataViaNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iName := self.Args("name").String()
	if helper.IsHas(iName) {
		_CommentFieldData := model.HasCommentFieldDataViaName(iName)
		var m = map[string]interface{}{}
		m["comment_field_data"] = _CommentFieldData
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasCommentFieldDataViaName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasCommentFieldDataViaMailHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iMail := self.Args("mail").String()
	if helper.IsHas(iMail) {
		_CommentFieldData := model.HasCommentFieldDataViaMail(iMail)
		var m = map[string]interface{}{}
		m["comment_field_data"] = _CommentFieldData
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasCommentFieldDataViaMail's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasCommentFieldDataViaHomepageHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iHomepage := self.Args("homepage").String()
	if helper.IsHas(iHomepage) {
		_CommentFieldData := model.HasCommentFieldDataViaHomepage(iHomepage)
		var m = map[string]interface{}{}
		m["comment_field_data"] = _CommentFieldData
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasCommentFieldDataViaHomepage's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasCommentFieldDataViaHostnameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iHostname := self.Args("hostname").String()
	if helper.IsHas(iHostname) {
		_CommentFieldData := model.HasCommentFieldDataViaHostname(iHostname)
		var m = map[string]interface{}{}
		m["comment_field_data"] = _CommentFieldData
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasCommentFieldDataViaHostname's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasCommentFieldDataViaCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iCreated := self.Args("created").MustInt()
	if helper.IsHas(iCreated) {
		_CommentFieldData := model.HasCommentFieldDataViaCreated(iCreated)
		var m = map[string]interface{}{}
		m["comment_field_data"] = _CommentFieldData
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasCommentFieldDataViaCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasCommentFieldDataViaChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iChanged := self.Args("changed").MustInt()
	if helper.IsHas(iChanged) {
		_CommentFieldData := model.HasCommentFieldDataViaChanged(iChanged)
		var m = map[string]interface{}{}
		m["comment_field_data"] = _CommentFieldData
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasCommentFieldDataViaChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasCommentFieldDataViaStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iStatus := self.Args("status").MustInt()
	if helper.IsHas(iStatus) {
		_CommentFieldData := model.HasCommentFieldDataViaStatus(iStatus)
		var m = map[string]interface{}{}
		m["comment_field_data"] = _CommentFieldData
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasCommentFieldDataViaStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasCommentFieldDataViaThreadHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iThread := self.Args("thread").String()
	if helper.IsHas(iThread) {
		_CommentFieldData := model.HasCommentFieldDataViaThread(iThread)
		var m = map[string]interface{}{}
		m["comment_field_data"] = _CommentFieldData
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasCommentFieldDataViaThread's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasCommentFieldDataViaEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iEntityType := self.Args("entity_type").String()
	if helper.IsHas(iEntityType) {
		_CommentFieldData := model.HasCommentFieldDataViaEntityType(iEntityType)
		var m = map[string]interface{}{}
		m["comment_field_data"] = _CommentFieldData
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasCommentFieldDataViaEntityType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasCommentFieldDataViaFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iFieldName := self.Args("field_name").String()
	if helper.IsHas(iFieldName) {
		_CommentFieldData := model.HasCommentFieldDataViaFieldName(iFieldName)
		var m = map[string]interface{}{}
		m["comment_field_data"] = _CommentFieldData
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasCommentFieldDataViaFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasCommentFieldDataViaDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iDefaultLangcode := self.Args("default_langcode").MustInt()
	if helper.IsHas(iDefaultLangcode) {
		_CommentFieldData := model.HasCommentFieldDataViaDefaultLangcode(iDefaultLangcode)
		var m = map[string]interface{}{}
		m["comment_field_data"] = _CommentFieldData
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasCommentFieldDataViaDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDataViaCidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iCid := self.Args("cid").MustInt64()
	if helper.IsHas(iCid) {
		_CommentFieldData, _error := model.GetCommentFieldDataViaCid(iCid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDataViaCid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDataViaCommentTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iCommentType := self.Args("comment_type").String()
	if helper.IsHas(iCommentType) {
		_CommentFieldData, _error := model.GetCommentFieldDataViaCommentType(iCommentType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDataViaCommentType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDataViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iLangcode := self.Args("langcode").String()
	if helper.IsHas(iLangcode) {
		_CommentFieldData, _error := model.GetCommentFieldDataViaLangcode(iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDataViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDataViaPidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iPid := self.Args("pid").MustInt()
	if helper.IsHas(iPid) {
		_CommentFieldData, _error := model.GetCommentFieldDataViaPid(iPid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDataViaPid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDataViaEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iEntityId := self.Args("entity_id").MustInt()
	if helper.IsHas(iEntityId) {
		_CommentFieldData, _error := model.GetCommentFieldDataViaEntityId(iEntityId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDataViaEntityId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDataViaSubjectHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iSubject := self.Args("subject").String()
	if helper.IsHas(iSubject) {
		_CommentFieldData, _error := model.GetCommentFieldDataViaSubject(iSubject)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDataViaSubject's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDataViaUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iUid := self.Args("uid").MustInt()
	if helper.IsHas(iUid) {
		_CommentFieldData, _error := model.GetCommentFieldDataViaUid(iUid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDataViaUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDataViaNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iName := self.Args("name").String()
	if helper.IsHas(iName) {
		_CommentFieldData, _error := model.GetCommentFieldDataViaName(iName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDataViaName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDataViaMailHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iMail := self.Args("mail").String()
	if helper.IsHas(iMail) {
		_CommentFieldData, _error := model.GetCommentFieldDataViaMail(iMail)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDataViaMail's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDataViaHomepageHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iHomepage := self.Args("homepage").String()
	if helper.IsHas(iHomepage) {
		_CommentFieldData, _error := model.GetCommentFieldDataViaHomepage(iHomepage)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDataViaHomepage's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDataViaHostnameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iHostname := self.Args("hostname").String()
	if helper.IsHas(iHostname) {
		_CommentFieldData, _error := model.GetCommentFieldDataViaHostname(iHostname)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDataViaHostname's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDataViaCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iCreated := self.Args("created").MustInt()
	if helper.IsHas(iCreated) {
		_CommentFieldData, _error := model.GetCommentFieldDataViaCreated(iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDataViaCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDataViaChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iChanged := self.Args("changed").MustInt()
	if helper.IsHas(iChanged) {
		_CommentFieldData, _error := model.GetCommentFieldDataViaChanged(iChanged)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDataViaChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDataViaStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iStatus := self.Args("status").MustInt()
	if helper.IsHas(iStatus) {
		_CommentFieldData, _error := model.GetCommentFieldDataViaStatus(iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDataViaStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDataViaThreadHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iThread := self.Args("thread").String()
	if helper.IsHas(iThread) {
		_CommentFieldData, _error := model.GetCommentFieldDataViaThread(iThread)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDataViaThread's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDataViaEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iEntityType := self.Args("entity_type").String()
	if helper.IsHas(iEntityType) {
		_CommentFieldData, _error := model.GetCommentFieldDataViaEntityType(iEntityType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDataViaEntityType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDataViaFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iFieldName := self.Args("field_name").String()
	if helper.IsHas(iFieldName) {
		_CommentFieldData, _error := model.GetCommentFieldDataViaFieldName(iFieldName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDataViaFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentFieldDataViaDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iDefaultLangcode := self.Args("default_langcode").MustInt()
	if helper.IsHas(iDefaultLangcode) {
		_CommentFieldData, _error := model.GetCommentFieldDataViaDefaultLangcode(iDefaultLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the GetCommentFieldDataViaDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetCommentFieldDataViaCidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Cid_ := self.Args("cid").MustInt64()
	if helper.IsHas(Cid_) {
		var iCommentFieldData model.CommentFieldData
		self.Bind(&iCommentFieldData)
		_CommentFieldData, _error := model.SetCommentFieldDataViaCid(Cid_, &iCommentFieldData)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the SetCommentFieldDataViaCid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetCommentFieldDataViaCommentTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	CommentType_ := self.Args("comment_type").String()
	if helper.IsHas(CommentType_) {
		var iCommentFieldData model.CommentFieldData
		self.Bind(&iCommentFieldData)
		_CommentFieldData, _error := model.SetCommentFieldDataViaCommentType(CommentType_, &iCommentFieldData)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the SetCommentFieldDataViaCommentType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetCommentFieldDataViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	if helper.IsHas(Langcode_) {
		var iCommentFieldData model.CommentFieldData
		self.Bind(&iCommentFieldData)
		_CommentFieldData, _error := model.SetCommentFieldDataViaLangcode(Langcode_, &iCommentFieldData)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the SetCommentFieldDataViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetCommentFieldDataViaPidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Pid_ := self.Args("pid").MustInt()
	if helper.IsHas(Pid_) {
		var iCommentFieldData model.CommentFieldData
		self.Bind(&iCommentFieldData)
		_CommentFieldData, _error := model.SetCommentFieldDataViaPid(Pid_, &iCommentFieldData)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the SetCommentFieldDataViaPid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetCommentFieldDataViaEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	EntityId_ := self.Args("entity_id").MustInt()
	if helper.IsHas(EntityId_) {
		var iCommentFieldData model.CommentFieldData
		self.Bind(&iCommentFieldData)
		_CommentFieldData, _error := model.SetCommentFieldDataViaEntityId(EntityId_, &iCommentFieldData)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the SetCommentFieldDataViaEntityId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetCommentFieldDataViaSubjectHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Subject_ := self.Args("subject").String()
	if helper.IsHas(Subject_) {
		var iCommentFieldData model.CommentFieldData
		self.Bind(&iCommentFieldData)
		_CommentFieldData, _error := model.SetCommentFieldDataViaSubject(Subject_, &iCommentFieldData)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the SetCommentFieldDataViaSubject's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetCommentFieldDataViaUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Uid_ := self.Args("uid").MustInt()
	if helper.IsHas(Uid_) {
		var iCommentFieldData model.CommentFieldData
		self.Bind(&iCommentFieldData)
		_CommentFieldData, _error := model.SetCommentFieldDataViaUid(Uid_, &iCommentFieldData)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the SetCommentFieldDataViaUid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetCommentFieldDataViaNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Name_ := self.Args("name").String()
	if helper.IsHas(Name_) {
		var iCommentFieldData model.CommentFieldData
		self.Bind(&iCommentFieldData)
		_CommentFieldData, _error := model.SetCommentFieldDataViaName(Name_, &iCommentFieldData)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the SetCommentFieldDataViaName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetCommentFieldDataViaMailHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Mail_ := self.Args("mail").String()
	if helper.IsHas(Mail_) {
		var iCommentFieldData model.CommentFieldData
		self.Bind(&iCommentFieldData)
		_CommentFieldData, _error := model.SetCommentFieldDataViaMail(Mail_, &iCommentFieldData)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the SetCommentFieldDataViaMail's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetCommentFieldDataViaHomepageHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Homepage_ := self.Args("homepage").String()
	if helper.IsHas(Homepage_) {
		var iCommentFieldData model.CommentFieldData
		self.Bind(&iCommentFieldData)
		_CommentFieldData, _error := model.SetCommentFieldDataViaHomepage(Homepage_, &iCommentFieldData)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the SetCommentFieldDataViaHomepage's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetCommentFieldDataViaHostnameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Hostname_ := self.Args("hostname").String()
	if helper.IsHas(Hostname_) {
		var iCommentFieldData model.CommentFieldData
		self.Bind(&iCommentFieldData)
		_CommentFieldData, _error := model.SetCommentFieldDataViaHostname(Hostname_, &iCommentFieldData)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the SetCommentFieldDataViaHostname's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetCommentFieldDataViaCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Created_ := self.Args("created").MustInt()
	if helper.IsHas(Created_) {
		var iCommentFieldData model.CommentFieldData
		self.Bind(&iCommentFieldData)
		_CommentFieldData, _error := model.SetCommentFieldDataViaCreated(Created_, &iCommentFieldData)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the SetCommentFieldDataViaCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetCommentFieldDataViaChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Changed_ := self.Args("changed").MustInt()
	if helper.IsHas(Changed_) {
		var iCommentFieldData model.CommentFieldData
		self.Bind(&iCommentFieldData)
		_CommentFieldData, _error := model.SetCommentFieldDataViaChanged(Changed_, &iCommentFieldData)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the SetCommentFieldDataViaChanged's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetCommentFieldDataViaStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Status_ := self.Args("status").MustInt()
	if helper.IsHas(Status_) {
		var iCommentFieldData model.CommentFieldData
		self.Bind(&iCommentFieldData)
		_CommentFieldData, _error := model.SetCommentFieldDataViaStatus(Status_, &iCommentFieldData)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the SetCommentFieldDataViaStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetCommentFieldDataViaThreadHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Thread_ := self.Args("thread").String()
	if helper.IsHas(Thread_) {
		var iCommentFieldData model.CommentFieldData
		self.Bind(&iCommentFieldData)
		_CommentFieldData, _error := model.SetCommentFieldDataViaThread(Thread_, &iCommentFieldData)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the SetCommentFieldDataViaThread's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetCommentFieldDataViaEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	EntityType_ := self.Args("entity_type").String()
	if helper.IsHas(EntityType_) {
		var iCommentFieldData model.CommentFieldData
		self.Bind(&iCommentFieldData)
		_CommentFieldData, _error := model.SetCommentFieldDataViaEntityType(EntityType_, &iCommentFieldData)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the SetCommentFieldDataViaEntityType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetCommentFieldDataViaFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	FieldName_ := self.Args("field_name").String()
	if helper.IsHas(FieldName_) {
		var iCommentFieldData model.CommentFieldData
		self.Bind(&iCommentFieldData)
		_CommentFieldData, _error := model.SetCommentFieldDataViaFieldName(FieldName_, &iCommentFieldData)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the SetCommentFieldDataViaFieldName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetCommentFieldDataViaDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	DefaultLangcode_ := self.Args("default_langcode").MustInt()
	if helper.IsHas(DefaultLangcode_) {
		var iCommentFieldData model.CommentFieldData
		self.Bind(&iCommentFieldData)
		_CommentFieldData, _error := model.SetCommentFieldDataViaDefaultLangcode(DefaultLangcode_, &iCommentFieldData)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_CommentFieldData)
	}
	herr.Message = "Can't get to the SetCommentFieldDataViaDefaultLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostAddCommentFieldDataHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Cid_ := self.Args("cid").MustInt64()
	CommentType_ := self.Args("comment_type").String()
	Langcode_ := self.Args("langcode").String()
	Pid_ := self.Args("pid").MustInt()
	EntityId_ := self.Args("entity_id").MustInt()
	Subject_ := self.Args("subject").String()
	Uid_ := self.Args("uid").MustInt()
	Name_ := self.Args("name").String()
	Mail_ := self.Args("mail").String()
	Homepage_ := self.Args("homepage").String()
	Hostname_ := self.Args("hostname").String()
	Created_ := self.Args("created").MustInt()
	Changed_ := self.Args("changed").MustInt()
	Status_ := self.Args("status").MustInt()
	Thread_ := self.Args("thread").String()
	EntityType_ := self.Args("entity_type").String()
	FieldName_ := self.Args("field_name").String()
	DefaultLangcode_ := self.Args("default_langcode").MustInt()

	if helper.IsHas(Cid_) {
		_error := model.AddCommentFieldData(Cid_,CommentType_,Langcode_,Pid_,EntityId_,Subject_,Uid_,Name_,Mail_,Homepage_,Hostname_,Created_,Changed_,Status_,Thread_,EntityType_,FieldName_,DefaultLangcode_)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		herr.Message = "StatusOK"
		herr.Status = macross.StatusOK
		return self.JSON(herr)
	}
	herr.Message = "Can't get to the AddCommentFieldData's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostCommentFieldDataHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iCommentFieldData model.CommentFieldData
	self.Bind(&iCommentFieldData)
	_int64, _error := model.PostCommentFieldData(&iCommentFieldData)
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

func PutCommentFieldDataHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iCommentFieldData model.CommentFieldData
	self.Bind(&iCommentFieldData)
	_int64, _error := model.PutCommentFieldData(&iCommentFieldData)
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

func PutCommentFieldDataViaCidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Cid_ := self.Args("cid").MustInt64()
	var iCommentFieldData model.CommentFieldData
	self.Bind(&iCommentFieldData)
	_int64, _error := model.PutCommentFieldDataViaCid(Cid_, &iCommentFieldData)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutCommentFieldDataViaCommentTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	CommentType_ := self.Args("comment_type").String()
	var iCommentFieldData model.CommentFieldData
	self.Bind(&iCommentFieldData)
	_int64, _error := model.PutCommentFieldDataViaCommentType(CommentType_, &iCommentFieldData)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutCommentFieldDataViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	var iCommentFieldData model.CommentFieldData
	self.Bind(&iCommentFieldData)
	_int64, _error := model.PutCommentFieldDataViaLangcode(Langcode_, &iCommentFieldData)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutCommentFieldDataViaPidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Pid_ := self.Args("pid").MustInt()
	var iCommentFieldData model.CommentFieldData
	self.Bind(&iCommentFieldData)
	_int64, _error := model.PutCommentFieldDataViaPid(Pid_, &iCommentFieldData)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutCommentFieldDataViaEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	EntityId_ := self.Args("entity_id").MustInt()
	var iCommentFieldData model.CommentFieldData
	self.Bind(&iCommentFieldData)
	_int64, _error := model.PutCommentFieldDataViaEntityId(EntityId_, &iCommentFieldData)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutCommentFieldDataViaSubjectHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Subject_ := self.Args("subject").String()
	var iCommentFieldData model.CommentFieldData
	self.Bind(&iCommentFieldData)
	_int64, _error := model.PutCommentFieldDataViaSubject(Subject_, &iCommentFieldData)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutCommentFieldDataViaUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Uid_ := self.Args("uid").MustInt()
	var iCommentFieldData model.CommentFieldData
	self.Bind(&iCommentFieldData)
	_int64, _error := model.PutCommentFieldDataViaUid(Uid_, &iCommentFieldData)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutCommentFieldDataViaNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Name_ := self.Args("name").String()
	var iCommentFieldData model.CommentFieldData
	self.Bind(&iCommentFieldData)
	_int64, _error := model.PutCommentFieldDataViaName(Name_, &iCommentFieldData)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutCommentFieldDataViaMailHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Mail_ := self.Args("mail").String()
	var iCommentFieldData model.CommentFieldData
	self.Bind(&iCommentFieldData)
	_int64, _error := model.PutCommentFieldDataViaMail(Mail_, &iCommentFieldData)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutCommentFieldDataViaHomepageHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Homepage_ := self.Args("homepage").String()
	var iCommentFieldData model.CommentFieldData
	self.Bind(&iCommentFieldData)
	_int64, _error := model.PutCommentFieldDataViaHomepage(Homepage_, &iCommentFieldData)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutCommentFieldDataViaHostnameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Hostname_ := self.Args("hostname").String()
	var iCommentFieldData model.CommentFieldData
	self.Bind(&iCommentFieldData)
	_int64, _error := model.PutCommentFieldDataViaHostname(Hostname_, &iCommentFieldData)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutCommentFieldDataViaCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Created_ := self.Args("created").MustInt()
	var iCommentFieldData model.CommentFieldData
	self.Bind(&iCommentFieldData)
	_int64, _error := model.PutCommentFieldDataViaCreated(Created_, &iCommentFieldData)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutCommentFieldDataViaChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Changed_ := self.Args("changed").MustInt()
	var iCommentFieldData model.CommentFieldData
	self.Bind(&iCommentFieldData)
	_int64, _error := model.PutCommentFieldDataViaChanged(Changed_, &iCommentFieldData)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutCommentFieldDataViaStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Status_ := self.Args("status").MustInt()
	var iCommentFieldData model.CommentFieldData
	self.Bind(&iCommentFieldData)
	_int64, _error := model.PutCommentFieldDataViaStatus(Status_, &iCommentFieldData)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutCommentFieldDataViaThreadHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Thread_ := self.Args("thread").String()
	var iCommentFieldData model.CommentFieldData
	self.Bind(&iCommentFieldData)
	_int64, _error := model.PutCommentFieldDataViaThread(Thread_, &iCommentFieldData)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutCommentFieldDataViaEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	EntityType_ := self.Args("entity_type").String()
	var iCommentFieldData model.CommentFieldData
	self.Bind(&iCommentFieldData)
	_int64, _error := model.PutCommentFieldDataViaEntityType(EntityType_, &iCommentFieldData)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutCommentFieldDataViaFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	FieldName_ := self.Args("field_name").String()
	var iCommentFieldData model.CommentFieldData
	self.Bind(&iCommentFieldData)
	_int64, _error := model.PutCommentFieldDataViaFieldName(FieldName_, &iCommentFieldData)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutCommentFieldDataViaDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	DefaultLangcode_ := self.Args("default_langcode").MustInt()
	var iCommentFieldData model.CommentFieldData
	self.Bind(&iCommentFieldData)
	_int64, _error := model.PutCommentFieldDataViaDefaultLangcode(DefaultLangcode_, &iCommentFieldData)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateCommentFieldDataViaCidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Cid_ := self.Args("cid").MustInt64()
	var iCommentFieldData model.CommentFieldData
	self.Bind(&iCommentFieldData)
	var iMap = helper.StructToMap(iCommentFieldData)
	_error := model.UpdateCommentFieldDataViaCid(Cid_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateCommentFieldDataViaCommentTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	CommentType_ := self.Args("comment_type").String()
	var iCommentFieldData model.CommentFieldData
	self.Bind(&iCommentFieldData)
	var iMap = helper.StructToMap(iCommentFieldData)
	_error := model.UpdateCommentFieldDataViaCommentType(CommentType_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateCommentFieldDataViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	var iCommentFieldData model.CommentFieldData
	self.Bind(&iCommentFieldData)
	var iMap = helper.StructToMap(iCommentFieldData)
	_error := model.UpdateCommentFieldDataViaLangcode(Langcode_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateCommentFieldDataViaPidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Pid_ := self.Args("pid").MustInt()
	var iCommentFieldData model.CommentFieldData
	self.Bind(&iCommentFieldData)
	var iMap = helper.StructToMap(iCommentFieldData)
	_error := model.UpdateCommentFieldDataViaPid(Pid_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateCommentFieldDataViaEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	EntityId_ := self.Args("entity_id").MustInt()
	var iCommentFieldData model.CommentFieldData
	self.Bind(&iCommentFieldData)
	var iMap = helper.StructToMap(iCommentFieldData)
	_error := model.UpdateCommentFieldDataViaEntityId(EntityId_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateCommentFieldDataViaSubjectHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Subject_ := self.Args("subject").String()
	var iCommentFieldData model.CommentFieldData
	self.Bind(&iCommentFieldData)
	var iMap = helper.StructToMap(iCommentFieldData)
	_error := model.UpdateCommentFieldDataViaSubject(Subject_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateCommentFieldDataViaUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Uid_ := self.Args("uid").MustInt()
	var iCommentFieldData model.CommentFieldData
	self.Bind(&iCommentFieldData)
	var iMap = helper.StructToMap(iCommentFieldData)
	_error := model.UpdateCommentFieldDataViaUid(Uid_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateCommentFieldDataViaNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Name_ := self.Args("name").String()
	var iCommentFieldData model.CommentFieldData
	self.Bind(&iCommentFieldData)
	var iMap = helper.StructToMap(iCommentFieldData)
	_error := model.UpdateCommentFieldDataViaName(Name_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateCommentFieldDataViaMailHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Mail_ := self.Args("mail").String()
	var iCommentFieldData model.CommentFieldData
	self.Bind(&iCommentFieldData)
	var iMap = helper.StructToMap(iCommentFieldData)
	_error := model.UpdateCommentFieldDataViaMail(Mail_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateCommentFieldDataViaHomepageHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Homepage_ := self.Args("homepage").String()
	var iCommentFieldData model.CommentFieldData
	self.Bind(&iCommentFieldData)
	var iMap = helper.StructToMap(iCommentFieldData)
	_error := model.UpdateCommentFieldDataViaHomepage(Homepage_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateCommentFieldDataViaHostnameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Hostname_ := self.Args("hostname").String()
	var iCommentFieldData model.CommentFieldData
	self.Bind(&iCommentFieldData)
	var iMap = helper.StructToMap(iCommentFieldData)
	_error := model.UpdateCommentFieldDataViaHostname(Hostname_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateCommentFieldDataViaCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Created_ := self.Args("created").MustInt()
	var iCommentFieldData model.CommentFieldData
	self.Bind(&iCommentFieldData)
	var iMap = helper.StructToMap(iCommentFieldData)
	_error := model.UpdateCommentFieldDataViaCreated(Created_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateCommentFieldDataViaChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Changed_ := self.Args("changed").MustInt()
	var iCommentFieldData model.CommentFieldData
	self.Bind(&iCommentFieldData)
	var iMap = helper.StructToMap(iCommentFieldData)
	_error := model.UpdateCommentFieldDataViaChanged(Changed_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateCommentFieldDataViaStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Status_ := self.Args("status").MustInt()
	var iCommentFieldData model.CommentFieldData
	self.Bind(&iCommentFieldData)
	var iMap = helper.StructToMap(iCommentFieldData)
	_error := model.UpdateCommentFieldDataViaStatus(Status_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateCommentFieldDataViaThreadHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Thread_ := self.Args("thread").String()
	var iCommentFieldData model.CommentFieldData
	self.Bind(&iCommentFieldData)
	var iMap = helper.StructToMap(iCommentFieldData)
	_error := model.UpdateCommentFieldDataViaThread(Thread_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateCommentFieldDataViaEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	EntityType_ := self.Args("entity_type").String()
	var iCommentFieldData model.CommentFieldData
	self.Bind(&iCommentFieldData)
	var iMap = helper.StructToMap(iCommentFieldData)
	_error := model.UpdateCommentFieldDataViaEntityType(EntityType_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateCommentFieldDataViaFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	FieldName_ := self.Args("field_name").String()
	var iCommentFieldData model.CommentFieldData
	self.Bind(&iCommentFieldData)
	var iMap = helper.StructToMap(iCommentFieldData)
	_error := model.UpdateCommentFieldDataViaFieldName(FieldName_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateCommentFieldDataViaDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	DefaultLangcode_ := self.Args("default_langcode").MustInt()
	var iCommentFieldData model.CommentFieldData
	self.Bind(&iCommentFieldData)
	var iMap = helper.StructToMap(iCommentFieldData)
	_error := model.UpdateCommentFieldDataViaDefaultLangcode(DefaultLangcode_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteCommentFieldDataHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Cid_ := self.Args("cid").MustInt64()
	_int64, _error := model.DeleteCommentFieldData(Cid_)
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

func DeleteCommentFieldDataViaCidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Cid_ := self.Args("cid").MustInt64()
	_error := model.DeleteCommentFieldDataViaCid(Cid_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteCommentFieldDataViaCommentTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	CommentType_ := self.Args("comment_type").String()
	_error := model.DeleteCommentFieldDataViaCommentType(CommentType_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteCommentFieldDataViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	_error := model.DeleteCommentFieldDataViaLangcode(Langcode_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteCommentFieldDataViaPidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Pid_ := self.Args("pid").MustInt()
	_error := model.DeleteCommentFieldDataViaPid(Pid_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteCommentFieldDataViaEntityIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	EntityId_ := self.Args("entity_id").MustInt()
	_error := model.DeleteCommentFieldDataViaEntityId(EntityId_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteCommentFieldDataViaSubjectHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Subject_ := self.Args("subject").String()
	_error := model.DeleteCommentFieldDataViaSubject(Subject_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteCommentFieldDataViaUidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Uid_ := self.Args("uid").MustInt()
	_error := model.DeleteCommentFieldDataViaUid(Uid_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteCommentFieldDataViaNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Name_ := self.Args("name").String()
	_error := model.DeleteCommentFieldDataViaName(Name_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteCommentFieldDataViaMailHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Mail_ := self.Args("mail").String()
	_error := model.DeleteCommentFieldDataViaMail(Mail_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteCommentFieldDataViaHomepageHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Homepage_ := self.Args("homepage").String()
	_error := model.DeleteCommentFieldDataViaHomepage(Homepage_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteCommentFieldDataViaHostnameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Hostname_ := self.Args("hostname").String()
	_error := model.DeleteCommentFieldDataViaHostname(Hostname_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteCommentFieldDataViaCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Created_ := self.Args("created").MustInt()
	_error := model.DeleteCommentFieldDataViaCreated(Created_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteCommentFieldDataViaChangedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Changed_ := self.Args("changed").MustInt()
	_error := model.DeleteCommentFieldDataViaChanged(Changed_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteCommentFieldDataViaStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Status_ := self.Args("status").MustInt()
	_error := model.DeleteCommentFieldDataViaStatus(Status_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteCommentFieldDataViaThreadHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Thread_ := self.Args("thread").String()
	_error := model.DeleteCommentFieldDataViaThread(Thread_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteCommentFieldDataViaEntityTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	EntityType_ := self.Args("entity_type").String()
	_error := model.DeleteCommentFieldDataViaEntityType(EntityType_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteCommentFieldDataViaFieldNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	FieldName_ := self.Args("field_name").String()
	_error := model.DeleteCommentFieldDataViaFieldName(FieldName_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteCommentFieldDataViaDefaultLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	DefaultLangcode_ := self.Args("default_langcode").MustInt()
	_error := model.DeleteCommentFieldDataViaDefaultLangcode(DefaultLangcode_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}
