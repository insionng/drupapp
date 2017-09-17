package handler

import (
	"github.com/insionng/drupapp/helper"
	"github.com/insionng/drupapp/model"
	"github.com/insionng/macross"
)

func GetUrlAliasesCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()

	if offset > 0 {
		var m = map[string]interface{}{}
		_int64, _error := model.GetUrlAliasesCount(offset, limit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		m["url_aliassCount"] = _int64
		return self.JSON(m)
	}
	herr.Message = "Can't get to the GetUrlAliasesCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUrlAliasCountViaPidHandler(self *macross.Context) error {
	Pid_ := self.Args("pid").MustInt64()
	_int64 := model.GetUrlAliasCountViaPid(Pid_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["url_aliasCount"] = 0
	}
	m["url_aliasCount"] = _int64
	return self.JSON(m)
}

func GetUrlAliasCountViaSourceHandler(self *macross.Context) error {
	Source_ := self.Args("source").String()
	_int64 := model.GetUrlAliasCountViaSource(Source_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["url_aliasCount"] = 0
	}
	m["url_aliasCount"] = _int64
	return self.JSON(m)
}

func GetUrlAliasCountViaAliasHandler(self *macross.Context) error {
	Alias_ := self.Args("alias").String()
	_int64 := model.GetUrlAliasCountViaAlias(Alias_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["url_aliasCount"] = 0
	}
	m["url_aliasCount"] = _int64
	return self.JSON(m)
}

func GetUrlAliasCountViaLangcodeHandler(self *macross.Context) error {
	Langcode_ := self.Args("langcode").String()
	_int64 := model.GetUrlAliasCountViaLangcode(Langcode_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["url_aliasCount"] = 0
	}
	m["url_aliasCount"] = _int64
	return self.JSON(m)
}

func GetUrlAliasesViaPidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iPid := self.Args("pid").MustInt64()
	if (offset > 0) && helper.IsHas(iPid) {
		_UrlAlias, _error := model.GetUrlAliasesViaPid(offset, limit, iPid, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UrlAlias)
	}
	herr.Message = "Can't get to the GetUrlAliasesViaPid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUrlAliasesViaSourceHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iSource := self.Args("source").String()
	if (offset > 0) && helper.IsHas(iSource) {
		_UrlAlias, _error := model.GetUrlAliasesViaSource(offset, limit, iSource, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UrlAlias)
	}
	herr.Message = "Can't get to the GetUrlAliasesViaSource's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUrlAliasesViaAliasHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iAlias := self.Args("alias").String()
	if (offset > 0) && helper.IsHas(iAlias) {
		_UrlAlias, _error := model.GetUrlAliasesViaAlias(offset, limit, iAlias, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UrlAlias)
	}
	herr.Message = "Can't get to the GetUrlAliasesViaAlias's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUrlAliasesViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iLangcode := self.Args("langcode").String()
	if (offset > 0) && helper.IsHas(iLangcode) {
		_UrlAlias, _error := model.GetUrlAliasesViaLangcode(offset, limit, iLangcode, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UrlAlias)
	}
	herr.Message = "Can't get to the GetUrlAliasesViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUrlAliasesByPidAndSourceAndAliasHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPid := self.Args("pid").MustInt64()
	iSource := self.Args("source").String()
	iAlias := self.Args("alias").String()

	if helper.IsHas(iPid) {
		_UrlAlias, _error := model.GetUrlAliasesByPidAndSourceAndAlias(offset, limit, iPid,iSource,iAlias)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UrlAlias)
	}
	herr.Message = "Can't get to the GetUrlAliasesByPidAndSourceAndAlias's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUrlAliasesByPidAndSourceAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPid := self.Args("pid").MustInt64()
	iSource := self.Args("source").String()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iPid) {
		_UrlAlias, _error := model.GetUrlAliasesByPidAndSourceAndLangcode(offset, limit, iPid,iSource,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UrlAlias)
	}
	herr.Message = "Can't get to the GetUrlAliasesByPidAndSourceAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUrlAliasesByPidAndAliasAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPid := self.Args("pid").MustInt64()
	iAlias := self.Args("alias").String()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iPid) {
		_UrlAlias, _error := model.GetUrlAliasesByPidAndAliasAndLangcode(offset, limit, iPid,iAlias,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UrlAlias)
	}
	herr.Message = "Can't get to the GetUrlAliasesByPidAndAliasAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUrlAliasesBySourceAndAliasAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSource := self.Args("source").String()
	iAlias := self.Args("alias").String()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iSource) {
		_UrlAlias, _error := model.GetUrlAliasesBySourceAndAliasAndLangcode(offset, limit, iSource,iAlias,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UrlAlias)
	}
	herr.Message = "Can't get to the GetUrlAliasesBySourceAndAliasAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUrlAliasesByPidAndSourceHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPid := self.Args("pid").MustInt64()
	iSource := self.Args("source").String()

	if helper.IsHas(iPid) {
		_UrlAlias, _error := model.GetUrlAliasesByPidAndSource(offset, limit, iPid,iSource)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UrlAlias)
	}
	herr.Message = "Can't get to the GetUrlAliasesByPidAndSource's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUrlAliasesByPidAndAliasHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPid := self.Args("pid").MustInt64()
	iAlias := self.Args("alias").String()

	if helper.IsHas(iPid) {
		_UrlAlias, _error := model.GetUrlAliasesByPidAndAlias(offset, limit, iPid,iAlias)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UrlAlias)
	}
	herr.Message = "Can't get to the GetUrlAliasesByPidAndAlias's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUrlAliasesByPidAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPid := self.Args("pid").MustInt64()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iPid) {
		_UrlAlias, _error := model.GetUrlAliasesByPidAndLangcode(offset, limit, iPid,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UrlAlias)
	}
	herr.Message = "Can't get to the GetUrlAliasesByPidAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUrlAliasesBySourceAndAliasHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSource := self.Args("source").String()
	iAlias := self.Args("alias").String()

	if helper.IsHas(iSource) {
		_UrlAlias, _error := model.GetUrlAliasesBySourceAndAlias(offset, limit, iSource,iAlias)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UrlAlias)
	}
	herr.Message = "Can't get to the GetUrlAliasesBySourceAndAlias's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUrlAliasesBySourceAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSource := self.Args("source").String()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iSource) {
		_UrlAlias, _error := model.GetUrlAliasesBySourceAndLangcode(offset, limit, iSource,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UrlAlias)
	}
	herr.Message = "Can't get to the GetUrlAliasesBySourceAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUrlAliasesByAliasAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iAlias := self.Args("alias").String()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iAlias) {
		_UrlAlias, _error := model.GetUrlAliasesByAliasAndLangcode(offset, limit, iAlias,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UrlAlias)
	}
	herr.Message = "Can't get to the GetUrlAliasesByAliasAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUrlAliasesHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	if (offset > 0) && (len(field) > 0) {
		_UrlAlias, _error := model.GetUrlAliases(offset, limit, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UrlAlias)
	}
	herr.Message = "Can't get to the GetUrlAliases' args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasUrlAliasViaPidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iPid := self.Args("pid").MustInt64()
	if helper.IsHas(iPid) {
		_UrlAlias := model.HasUrlAliasViaPid(iPid)
		var m = map[string]interface{}{}
		m["url_alias"] = _UrlAlias
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasUrlAliasViaPid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasUrlAliasViaSourceHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iSource := self.Args("source").String()
	if helper.IsHas(iSource) {
		_UrlAlias := model.HasUrlAliasViaSource(iSource)
		var m = map[string]interface{}{}
		m["url_alias"] = _UrlAlias
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasUrlAliasViaSource's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasUrlAliasViaAliasHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iAlias := self.Args("alias").String()
	if helper.IsHas(iAlias) {
		_UrlAlias := model.HasUrlAliasViaAlias(iAlias)
		var m = map[string]interface{}{}
		m["url_alias"] = _UrlAlias
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasUrlAliasViaAlias's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasUrlAliasViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iLangcode := self.Args("langcode").String()
	if helper.IsHas(iLangcode) {
		_UrlAlias := model.HasUrlAliasViaLangcode(iLangcode)
		var m = map[string]interface{}{}
		m["url_alias"] = _UrlAlias
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasUrlAliasViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUrlAliasViaPidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iPid := self.Args("pid").MustInt64()
	if helper.IsHas(iPid) {
		_UrlAlias, _error := model.GetUrlAliasViaPid(iPid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UrlAlias)
	}
	herr.Message = "Can't get to the GetUrlAliasViaPid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUrlAliasViaSourceHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iSource := self.Args("source").String()
	if helper.IsHas(iSource) {
		_UrlAlias, _error := model.GetUrlAliasViaSource(iSource)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UrlAlias)
	}
	herr.Message = "Can't get to the GetUrlAliasViaSource's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUrlAliasViaAliasHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iAlias := self.Args("alias").String()
	if helper.IsHas(iAlias) {
		_UrlAlias, _error := model.GetUrlAliasViaAlias(iAlias)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UrlAlias)
	}
	herr.Message = "Can't get to the GetUrlAliasViaAlias's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUrlAliasViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iLangcode := self.Args("langcode").String()
	if helper.IsHas(iLangcode) {
		_UrlAlias, _error := model.GetUrlAliasViaLangcode(iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UrlAlias)
	}
	herr.Message = "Can't get to the GetUrlAliasViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetUrlAliasViaPidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Pid_ := self.Args("pid").MustInt64()
	if helper.IsHas(Pid_) {
		var iUrlAlias model.UrlAlias
		self.Bind(&iUrlAlias)
		_UrlAlias, _error := model.SetUrlAliasViaPid(Pid_, &iUrlAlias)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UrlAlias)
	}
	herr.Message = "Can't get to the SetUrlAliasViaPid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetUrlAliasViaSourceHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Source_ := self.Args("source").String()
	if helper.IsHas(Source_) {
		var iUrlAlias model.UrlAlias
		self.Bind(&iUrlAlias)
		_UrlAlias, _error := model.SetUrlAliasViaSource(Source_, &iUrlAlias)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UrlAlias)
	}
	herr.Message = "Can't get to the SetUrlAliasViaSource's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetUrlAliasViaAliasHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Alias_ := self.Args("alias").String()
	if helper.IsHas(Alias_) {
		var iUrlAlias model.UrlAlias
		self.Bind(&iUrlAlias)
		_UrlAlias, _error := model.SetUrlAliasViaAlias(Alias_, &iUrlAlias)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UrlAlias)
	}
	herr.Message = "Can't get to the SetUrlAliasViaAlias's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetUrlAliasViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	if helper.IsHas(Langcode_) {
		var iUrlAlias model.UrlAlias
		self.Bind(&iUrlAlias)
		_UrlAlias, _error := model.SetUrlAliasViaLangcode(Langcode_, &iUrlAlias)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_UrlAlias)
	}
	herr.Message = "Can't get to the SetUrlAliasViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostAddUrlAliasHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Pid_ := self.Args("pid").MustInt64()
	Source_ := self.Args("source").String()
	Alias_ := self.Args("alias").String()
	Langcode_ := self.Args("langcode").String()

	if helper.IsHas(Pid_) {
		_error := model.AddUrlAlias(Pid_,Source_,Alias_,Langcode_)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		herr.Message = "StatusOK"
		herr.Status = macross.StatusOK
		return self.JSON(herr)
	}
	herr.Message = "Can't get to the AddUrlAlias's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostUrlAliasHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iUrlAlias model.UrlAlias
	self.Bind(&iUrlAlias)
	_int64, _error := model.PostUrlAlias(&iUrlAlias)
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

func PutUrlAliasHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iUrlAlias model.UrlAlias
	self.Bind(&iUrlAlias)
	_int64, _error := model.PutUrlAlias(&iUrlAlias)
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

func PutUrlAliasViaPidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Pid_ := self.Args("pid").MustInt64()
	var iUrlAlias model.UrlAlias
	self.Bind(&iUrlAlias)
	_int64, _error := model.PutUrlAliasViaPid(Pid_, &iUrlAlias)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUrlAliasViaSourceHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Source_ := self.Args("source").String()
	var iUrlAlias model.UrlAlias
	self.Bind(&iUrlAlias)
	_int64, _error := model.PutUrlAliasViaSource(Source_, &iUrlAlias)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUrlAliasViaAliasHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Alias_ := self.Args("alias").String()
	var iUrlAlias model.UrlAlias
	self.Bind(&iUrlAlias)
	_int64, _error := model.PutUrlAliasViaAlias(Alias_, &iUrlAlias)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUrlAliasViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	var iUrlAlias model.UrlAlias
	self.Bind(&iUrlAlias)
	_int64, _error := model.PutUrlAliasViaLangcode(Langcode_, &iUrlAlias)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateUrlAliasViaPidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Pid_ := self.Args("pid").MustInt64()
	var iUrlAlias model.UrlAlias
	self.Bind(&iUrlAlias)
	var iMap = helper.StructToMap(iUrlAlias)
	_error := model.UpdateUrlAliasViaPid(Pid_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateUrlAliasViaSourceHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Source_ := self.Args("source").String()
	var iUrlAlias model.UrlAlias
	self.Bind(&iUrlAlias)
	var iMap = helper.StructToMap(iUrlAlias)
	_error := model.UpdateUrlAliasViaSource(Source_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateUrlAliasViaAliasHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Alias_ := self.Args("alias").String()
	var iUrlAlias model.UrlAlias
	self.Bind(&iUrlAlias)
	var iMap = helper.StructToMap(iUrlAlias)
	_error := model.UpdateUrlAliasViaAlias(Alias_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateUrlAliasViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	var iUrlAlias model.UrlAlias
	self.Bind(&iUrlAlias)
	var iMap = helper.StructToMap(iUrlAlias)
	_error := model.UpdateUrlAliasViaLangcode(Langcode_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteUrlAliasHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Pid_ := self.Args("pid").MustInt64()
	_int64, _error := model.DeleteUrlAlias(Pid_)
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

func DeleteUrlAliasViaPidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Pid_ := self.Args("pid").MustInt64()
	_error := model.DeleteUrlAliasViaPid(Pid_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteUrlAliasViaSourceHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Source_ := self.Args("source").String()
	_error := model.DeleteUrlAliasViaSource(Source_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteUrlAliasViaAliasHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Alias_ := self.Args("alias").String()
	_error := model.DeleteUrlAliasViaAlias(Alias_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteUrlAliasViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	_error := model.DeleteUrlAliasViaLangcode(Langcode_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}
