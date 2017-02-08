package handler

import (
	"drupapp/helper"
	"drupapp/model"
	"github.com/insionng/macross"
)

func GetFileUsagesCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()

	if offset > 0 {
		var m = map[string]interface{}{}
		_int64, _error := model.GetFileUsagesCount(offset, limit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		m["file_usagesCount"] = _int64
		return self.JSON(m)
	}
	herr.Message = "Can't get to the GetFileUsagesCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileUsageCountViaFidHandler(self *macross.Context) error {
	Fid_ := self.Args("fid").MustInt64()
	_int64 := model.GetFileUsageCountViaFid(Fid_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["file_usageCount"] = 0
	}
	m["file_usageCount"] = _int64
	return self.JSON(m)
}

func GetFileUsageCountViaModuleHandler(self *macross.Context) error {
	Module_ := self.Args("module").String()
	_int64 := model.GetFileUsageCountViaModule(Module_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["file_usageCount"] = 0
	}
	m["file_usageCount"] = _int64
	return self.JSON(m)
}

func GetFileUsageCountViaTypeHandler(self *macross.Context) error {
	Type_ := self.Args("type").String()
	_int64 := model.GetFileUsageCountViaType(Type_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["file_usageCount"] = 0
	}
	m["file_usageCount"] = _int64
	return self.JSON(m)
}

func GetFileUsageCountViaIdHandler(self *macross.Context) error {
	Id_ := self.Args("id").String()
	_int64 := model.GetFileUsageCountViaId(Id_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["file_usageCount"] = 0
	}
	m["file_usageCount"] = _int64
	return self.JSON(m)
}

func GetFileUsageCountViaCountHandler(self *macross.Context) error {
	Count_ := self.Args("count").MustInt()
	_int64 := model.GetFileUsageCountViaCount(Count_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["file_usageCount"] = 0
	}
	m["file_usageCount"] = _int64
	return self.JSON(m)
}

func GetFileUsagesViaFidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iFid := self.Args("fid").MustInt64()
	if (offset > 0) && helper.IsHas(iFid) {
		_FileUsage, _error := model.GetFileUsagesViaFid(offset, limit, iFid, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileUsage)
	}
	herr.Message = "Can't get to the GetFileUsagesViaFid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileUsagesViaModuleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iModule := self.Args("module").String()
	if (offset > 0) && helper.IsHas(iModule) {
		_FileUsage, _error := model.GetFileUsagesViaModule(offset, limit, iModule, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileUsage)
	}
	herr.Message = "Can't get to the GetFileUsagesViaModule's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileUsagesViaTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iType := self.Args("type").String()
	if (offset > 0) && helper.IsHas(iType) {
		_FileUsage, _error := model.GetFileUsagesViaType(offset, limit, iType, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileUsage)
	}
	herr.Message = "Can't get to the GetFileUsagesViaType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileUsagesViaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iId := self.Args("id").String()
	if (offset > 0) && helper.IsHas(iId) {
		_FileUsage, _error := model.GetFileUsagesViaId(offset, limit, iId, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileUsage)
	}
	herr.Message = "Can't get to the GetFileUsagesViaId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileUsagesViaCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iCount := self.Args("count").MustInt()
	if (offset > 0) && helper.IsHas(iCount) {
		_FileUsage, _error := model.GetFileUsagesViaCount(offset, limit, iCount, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileUsage)
	}
	herr.Message = "Can't get to the GetFileUsagesViaCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileUsagesByFidAndModuleAndTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFid := self.Args("fid").MustInt64()
	iModule := self.Args("module").String()
	iType := self.Args("type").String()

	if helper.IsHas(iFid) {
		_FileUsage, _error := model.GetFileUsagesByFidAndModuleAndType(offset, limit, iFid,iModule,iType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileUsage)
	}
	herr.Message = "Can't get to the GetFileUsagesByFidAndModuleAndType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileUsagesByFidAndModuleAndIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFid := self.Args("fid").MustInt64()
	iModule := self.Args("module").String()
	iId := self.Args("id").String()

	if helper.IsHas(iFid) {
		_FileUsage, _error := model.GetFileUsagesByFidAndModuleAndId(offset, limit, iFid,iModule,iId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileUsage)
	}
	herr.Message = "Can't get to the GetFileUsagesByFidAndModuleAndId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileUsagesByFidAndModuleAndCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFid := self.Args("fid").MustInt64()
	iModule := self.Args("module").String()
	iCount := self.Args("count").MustInt()

	if helper.IsHas(iFid) {
		_FileUsage, _error := model.GetFileUsagesByFidAndModuleAndCount(offset, limit, iFid,iModule,iCount)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileUsage)
	}
	herr.Message = "Can't get to the GetFileUsagesByFidAndModuleAndCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileUsagesByFidAndTypeAndIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFid := self.Args("fid").MustInt64()
	iType := self.Args("type").String()
	iId := self.Args("id").String()

	if helper.IsHas(iFid) {
		_FileUsage, _error := model.GetFileUsagesByFidAndTypeAndId(offset, limit, iFid,iType,iId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileUsage)
	}
	herr.Message = "Can't get to the GetFileUsagesByFidAndTypeAndId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileUsagesByFidAndTypeAndCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFid := self.Args("fid").MustInt64()
	iType := self.Args("type").String()
	iCount := self.Args("count").MustInt()

	if helper.IsHas(iFid) {
		_FileUsage, _error := model.GetFileUsagesByFidAndTypeAndCount(offset, limit, iFid,iType,iCount)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileUsage)
	}
	herr.Message = "Can't get to the GetFileUsagesByFidAndTypeAndCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileUsagesByFidAndIdAndCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFid := self.Args("fid").MustInt64()
	iId := self.Args("id").String()
	iCount := self.Args("count").MustInt()

	if helper.IsHas(iFid) {
		_FileUsage, _error := model.GetFileUsagesByFidAndIdAndCount(offset, limit, iFid,iId,iCount)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileUsage)
	}
	herr.Message = "Can't get to the GetFileUsagesByFidAndIdAndCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileUsagesByModuleAndTypeAndIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iModule := self.Args("module").String()
	iType := self.Args("type").String()
	iId := self.Args("id").String()

	if helper.IsHas(iModule) {
		_FileUsage, _error := model.GetFileUsagesByModuleAndTypeAndId(offset, limit, iModule,iType,iId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileUsage)
	}
	herr.Message = "Can't get to the GetFileUsagesByModuleAndTypeAndId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileUsagesByModuleAndTypeAndCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iModule := self.Args("module").String()
	iType := self.Args("type").String()
	iCount := self.Args("count").MustInt()

	if helper.IsHas(iModule) {
		_FileUsage, _error := model.GetFileUsagesByModuleAndTypeAndCount(offset, limit, iModule,iType,iCount)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileUsage)
	}
	herr.Message = "Can't get to the GetFileUsagesByModuleAndTypeAndCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileUsagesByModuleAndIdAndCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iModule := self.Args("module").String()
	iId := self.Args("id").String()
	iCount := self.Args("count").MustInt()

	if helper.IsHas(iModule) {
		_FileUsage, _error := model.GetFileUsagesByModuleAndIdAndCount(offset, limit, iModule,iId,iCount)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileUsage)
	}
	herr.Message = "Can't get to the GetFileUsagesByModuleAndIdAndCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileUsagesByTypeAndIdAndCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iType := self.Args("type").String()
	iId := self.Args("id").String()
	iCount := self.Args("count").MustInt()

	if helper.IsHas(iType) {
		_FileUsage, _error := model.GetFileUsagesByTypeAndIdAndCount(offset, limit, iType,iId,iCount)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileUsage)
	}
	herr.Message = "Can't get to the GetFileUsagesByTypeAndIdAndCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileUsagesByFidAndModuleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFid := self.Args("fid").MustInt64()
	iModule := self.Args("module").String()

	if helper.IsHas(iFid) {
		_FileUsage, _error := model.GetFileUsagesByFidAndModule(offset, limit, iFid,iModule)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileUsage)
	}
	herr.Message = "Can't get to the GetFileUsagesByFidAndModule's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileUsagesByFidAndTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFid := self.Args("fid").MustInt64()
	iType := self.Args("type").String()

	if helper.IsHas(iFid) {
		_FileUsage, _error := model.GetFileUsagesByFidAndType(offset, limit, iFid,iType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileUsage)
	}
	herr.Message = "Can't get to the GetFileUsagesByFidAndType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileUsagesByFidAndIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFid := self.Args("fid").MustInt64()
	iId := self.Args("id").String()

	if helper.IsHas(iFid) {
		_FileUsage, _error := model.GetFileUsagesByFidAndId(offset, limit, iFid,iId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileUsage)
	}
	herr.Message = "Can't get to the GetFileUsagesByFidAndId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileUsagesByFidAndCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iFid := self.Args("fid").MustInt64()
	iCount := self.Args("count").MustInt()

	if helper.IsHas(iFid) {
		_FileUsage, _error := model.GetFileUsagesByFidAndCount(offset, limit, iFid,iCount)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileUsage)
	}
	herr.Message = "Can't get to the GetFileUsagesByFidAndCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileUsagesByModuleAndTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iModule := self.Args("module").String()
	iType := self.Args("type").String()

	if helper.IsHas(iModule) {
		_FileUsage, _error := model.GetFileUsagesByModuleAndType(offset, limit, iModule,iType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileUsage)
	}
	herr.Message = "Can't get to the GetFileUsagesByModuleAndType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileUsagesByModuleAndIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iModule := self.Args("module").String()
	iId := self.Args("id").String()

	if helper.IsHas(iModule) {
		_FileUsage, _error := model.GetFileUsagesByModuleAndId(offset, limit, iModule,iId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileUsage)
	}
	herr.Message = "Can't get to the GetFileUsagesByModuleAndId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileUsagesByModuleAndCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iModule := self.Args("module").String()
	iCount := self.Args("count").MustInt()

	if helper.IsHas(iModule) {
		_FileUsage, _error := model.GetFileUsagesByModuleAndCount(offset, limit, iModule,iCount)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileUsage)
	}
	herr.Message = "Can't get to the GetFileUsagesByModuleAndCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileUsagesByTypeAndIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iType := self.Args("type").String()
	iId := self.Args("id").String()

	if helper.IsHas(iType) {
		_FileUsage, _error := model.GetFileUsagesByTypeAndId(offset, limit, iType,iId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileUsage)
	}
	herr.Message = "Can't get to the GetFileUsagesByTypeAndId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileUsagesByTypeAndCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iType := self.Args("type").String()
	iCount := self.Args("count").MustInt()

	if helper.IsHas(iType) {
		_FileUsage, _error := model.GetFileUsagesByTypeAndCount(offset, limit, iType,iCount)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileUsage)
	}
	herr.Message = "Can't get to the GetFileUsagesByTypeAndCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileUsagesByIdAndCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").String()
	iCount := self.Args("count").MustInt()

	if helper.IsHas(iId) {
		_FileUsage, _error := model.GetFileUsagesByIdAndCount(offset, limit, iId,iCount)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileUsage)
	}
	herr.Message = "Can't get to the GetFileUsagesByIdAndCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileUsagesHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	if (offset > 0) && (len(field) > 0) {
		_FileUsage, _error := model.GetFileUsages(offset, limit, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileUsage)
	}
	herr.Message = "Can't get to the GetFileUsages' args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasFileUsageViaFidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iFid := self.Args("fid").MustInt64()
	if helper.IsHas(iFid) {
		_FileUsage := model.HasFileUsageViaFid(iFid)
		var m = map[string]interface{}{}
		m["file_usage"] = _FileUsage
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasFileUsageViaFid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasFileUsageViaModuleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iModule := self.Args("module").String()
	if helper.IsHas(iModule) {
		_FileUsage := model.HasFileUsageViaModule(iModule)
		var m = map[string]interface{}{}
		m["file_usage"] = _FileUsage
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasFileUsageViaModule's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasFileUsageViaTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iType := self.Args("type").String()
	if helper.IsHas(iType) {
		_FileUsage := model.HasFileUsageViaType(iType)
		var m = map[string]interface{}{}
		m["file_usage"] = _FileUsage
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasFileUsageViaType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasFileUsageViaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iId := self.Args("id").String()
	if helper.IsHas(iId) {
		_FileUsage := model.HasFileUsageViaId(iId)
		var m = map[string]interface{}{}
		m["file_usage"] = _FileUsage
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasFileUsageViaId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasFileUsageViaCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iCount := self.Args("count").MustInt()
	if helper.IsHas(iCount) {
		_FileUsage := model.HasFileUsageViaCount(iCount)
		var m = map[string]interface{}{}
		m["file_usage"] = _FileUsage
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasFileUsageViaCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileUsageViaFidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iFid := self.Args("fid").MustInt64()
	if helper.IsHas(iFid) {
		_FileUsage, _error := model.GetFileUsageViaFid(iFid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileUsage)
	}
	herr.Message = "Can't get to the GetFileUsageViaFid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileUsageViaModuleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iModule := self.Args("module").String()
	if helper.IsHas(iModule) {
		_FileUsage, _error := model.GetFileUsageViaModule(iModule)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileUsage)
	}
	herr.Message = "Can't get to the GetFileUsageViaModule's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileUsageViaTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iType := self.Args("type").String()
	if helper.IsHas(iType) {
		_FileUsage, _error := model.GetFileUsageViaType(iType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileUsage)
	}
	herr.Message = "Can't get to the GetFileUsageViaType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileUsageViaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iId := self.Args("id").String()
	if helper.IsHas(iId) {
		_FileUsage, _error := model.GetFileUsageViaId(iId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileUsage)
	}
	herr.Message = "Can't get to the GetFileUsageViaId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetFileUsageViaCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iCount := self.Args("count").MustInt()
	if helper.IsHas(iCount) {
		_FileUsage, _error := model.GetFileUsageViaCount(iCount)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileUsage)
	}
	herr.Message = "Can't get to the GetFileUsageViaCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetFileUsageViaFidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Fid_ := self.Args("fid").MustInt64()
	if helper.IsHas(Fid_) {
		var iFileUsage model.FileUsage
		self.Bind(&iFileUsage)
		_FileUsage, _error := model.SetFileUsageViaFid(Fid_, &iFileUsage)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileUsage)
	}
	herr.Message = "Can't get to the SetFileUsageViaFid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetFileUsageViaModuleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Module_ := self.Args("module").String()
	if helper.IsHas(Module_) {
		var iFileUsage model.FileUsage
		self.Bind(&iFileUsage)
		_FileUsage, _error := model.SetFileUsageViaModule(Module_, &iFileUsage)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileUsage)
	}
	herr.Message = "Can't get to the SetFileUsageViaModule's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetFileUsageViaTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Type_ := self.Args("type").String()
	if helper.IsHas(Type_) {
		var iFileUsage model.FileUsage
		self.Bind(&iFileUsage)
		_FileUsage, _error := model.SetFileUsageViaType(Type_, &iFileUsage)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileUsage)
	}
	herr.Message = "Can't get to the SetFileUsageViaType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetFileUsageViaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Id_ := self.Args("id").String()
	if helper.IsHas(Id_) {
		var iFileUsage model.FileUsage
		self.Bind(&iFileUsage)
		_FileUsage, _error := model.SetFileUsageViaId(Id_, &iFileUsage)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileUsage)
	}
	herr.Message = "Can't get to the SetFileUsageViaId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetFileUsageViaCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Count_ := self.Args("count").MustInt()
	if helper.IsHas(Count_) {
		var iFileUsage model.FileUsage
		self.Bind(&iFileUsage)
		_FileUsage, _error := model.SetFileUsageViaCount(Count_, &iFileUsage)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_FileUsage)
	}
	herr.Message = "Can't get to the SetFileUsageViaCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostAddFileUsageHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Fid_ := self.Args("fid").MustInt64()
	Module_ := self.Args("module").String()
	Type_ := self.Args("type").String()
	Id_ := self.Args("id").String()
	Count_ := self.Args("count").MustInt()

	if helper.IsHas(Fid_) {
		_error := model.AddFileUsage(Fid_,Module_,Type_,Id_,Count_)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		herr.Message = "StatusOK"
		herr.Status = macross.StatusOK
		return self.JSON(herr)
	}
	herr.Message = "Can't get to the AddFileUsage's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostFileUsageHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iFileUsage model.FileUsage
	self.Bind(&iFileUsage)
	_int64, _error := model.PostFileUsage(&iFileUsage)
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

func PutFileUsageHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iFileUsage model.FileUsage
	self.Bind(&iFileUsage)
	_int64, _error := model.PutFileUsage(&iFileUsage)
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

func PutFileUsageViaFidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Fid_ := self.Args("fid").MustInt64()
	var iFileUsage model.FileUsage
	self.Bind(&iFileUsage)
	_int64, _error := model.PutFileUsageViaFid(Fid_, &iFileUsage)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutFileUsageViaModuleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Module_ := self.Args("module").String()
	var iFileUsage model.FileUsage
	self.Bind(&iFileUsage)
	_int64, _error := model.PutFileUsageViaModule(Module_, &iFileUsage)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutFileUsageViaTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Type_ := self.Args("type").String()
	var iFileUsage model.FileUsage
	self.Bind(&iFileUsage)
	_int64, _error := model.PutFileUsageViaType(Type_, &iFileUsage)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutFileUsageViaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Id_ := self.Args("id").String()
	var iFileUsage model.FileUsage
	self.Bind(&iFileUsage)
	_int64, _error := model.PutFileUsageViaId(Id_, &iFileUsage)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutFileUsageViaCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Count_ := self.Args("count").MustInt()
	var iFileUsage model.FileUsage
	self.Bind(&iFileUsage)
	_int64, _error := model.PutFileUsageViaCount(Count_, &iFileUsage)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateFileUsageViaFidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Fid_ := self.Args("fid").MustInt64()
	var iFileUsage model.FileUsage
	self.Bind(&iFileUsage)
	var iMap = helper.StructToMap(iFileUsage)
	_error := model.UpdateFileUsageViaFid(Fid_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateFileUsageViaModuleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Module_ := self.Args("module").String()
	var iFileUsage model.FileUsage
	self.Bind(&iFileUsage)
	var iMap = helper.StructToMap(iFileUsage)
	_error := model.UpdateFileUsageViaModule(Module_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateFileUsageViaTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Type_ := self.Args("type").String()
	var iFileUsage model.FileUsage
	self.Bind(&iFileUsage)
	var iMap = helper.StructToMap(iFileUsage)
	_error := model.UpdateFileUsageViaType(Type_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateFileUsageViaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Id_ := self.Args("id").String()
	var iFileUsage model.FileUsage
	self.Bind(&iFileUsage)
	var iMap = helper.StructToMap(iFileUsage)
	_error := model.UpdateFileUsageViaId(Id_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateFileUsageViaCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Count_ := self.Args("count").MustInt()
	var iFileUsage model.FileUsage
	self.Bind(&iFileUsage)
	var iMap = helper.StructToMap(iFileUsage)
	_error := model.UpdateFileUsageViaCount(Count_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteFileUsageHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Fid_ := self.Args("fid").MustInt64()
	_int64, _error := model.DeleteFileUsage(Fid_)
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

func DeleteFileUsageViaFidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Fid_ := self.Args("fid").MustInt64()
	_error := model.DeleteFileUsageViaFid(Fid_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteFileUsageViaModuleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Module_ := self.Args("module").String()
	_error := model.DeleteFileUsageViaModule(Module_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteFileUsageViaTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Type_ := self.Args("type").String()
	_error := model.DeleteFileUsageViaType(Type_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteFileUsageViaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Id_ := self.Args("id").String()
	_error := model.DeleteFileUsageViaId(Id_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteFileUsageViaCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Count_ := self.Args("count").MustInt()
	_error := model.DeleteFileUsageViaCount(Count_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}
