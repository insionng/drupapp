package handler

import (
	"drupapp/helper"
	"drupapp/model"
	"github.com/insionng/macross"
)

func GetSearchDatasetsCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()

	if offset > 0 {
		var m = map[string]interface{}{}
		_int64, _error := model.GetSearchDatasetsCount(offset, limit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		m["search_datasetsCount"] = _int64
		return self.JSON(m)
	}
	herr.Message = "Can't get to the GetSearchDatasetsCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSearchDatasetCountViaSidHandler(self *macross.Context) error {
	Sid_ := self.Args("sid").MustInt64()
	_int64 := model.GetSearchDatasetCountViaSid(Sid_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["search_datasetCount"] = 0
	}
	m["search_datasetCount"] = _int64
	return self.JSON(m)
}

func GetSearchDatasetCountViaLangcodeHandler(self *macross.Context) error {
	Langcode_ := self.Args("langcode").String()
	_int64 := model.GetSearchDatasetCountViaLangcode(Langcode_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["search_datasetCount"] = 0
	}
	m["search_datasetCount"] = _int64
	return self.JSON(m)
}

func GetSearchDatasetCountViaTypeHandler(self *macross.Context) error {
	Type_ := self.Args("type").String()
	_int64 := model.GetSearchDatasetCountViaType(Type_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["search_datasetCount"] = 0
	}
	m["search_datasetCount"] = _int64
	return self.JSON(m)
}

func GetSearchDatasetCountViaDataHandler(self *macross.Context) error {
	Data_ := self.Args("data").String()
	_int64 := model.GetSearchDatasetCountViaData(Data_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["search_datasetCount"] = 0
	}
	m["search_datasetCount"] = _int64
	return self.JSON(m)
}

func GetSearchDatasetCountViaReindexHandler(self *macross.Context) error {
	Reindex_ := self.Args("reindex").MustInt()
	_int64 := model.GetSearchDatasetCountViaReindex(Reindex_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["search_datasetCount"] = 0
	}
	m["search_datasetCount"] = _int64
	return self.JSON(m)
}

func GetSearchDatasetsViaSidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iSid := self.Args("sid").MustInt64()
	if (offset > 0) && helper.IsHas(iSid) {
		_SearchDataset, _error := model.GetSearchDatasetsViaSid(offset, limit, iSid, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_SearchDataset)
	}
	herr.Message = "Can't get to the GetSearchDatasetsViaSid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSearchDatasetsViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iLangcode := self.Args("langcode").String()
	if (offset > 0) && helper.IsHas(iLangcode) {
		_SearchDataset, _error := model.GetSearchDatasetsViaLangcode(offset, limit, iLangcode, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_SearchDataset)
	}
	herr.Message = "Can't get to the GetSearchDatasetsViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSearchDatasetsViaTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iType := self.Args("type").String()
	if (offset > 0) && helper.IsHas(iType) {
		_SearchDataset, _error := model.GetSearchDatasetsViaType(offset, limit, iType, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_SearchDataset)
	}
	herr.Message = "Can't get to the GetSearchDatasetsViaType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSearchDatasetsViaDataHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iData := self.Args("data").String()
	if (offset > 0) && helper.IsHas(iData) {
		_SearchDataset, _error := model.GetSearchDatasetsViaData(offset, limit, iData, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_SearchDataset)
	}
	herr.Message = "Can't get to the GetSearchDatasetsViaData's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSearchDatasetsViaReindexHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iReindex := self.Args("reindex").MustInt()
	if (offset > 0) && helper.IsHas(iReindex) {
		_SearchDataset, _error := model.GetSearchDatasetsViaReindex(offset, limit, iReindex, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_SearchDataset)
	}
	herr.Message = "Can't get to the GetSearchDatasetsViaReindex's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSearchDatasetsBySidAndLangcodeAndTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSid := self.Args("sid").MustInt64()
	iLangcode := self.Args("langcode").String()
	iType := self.Args("type").String()

	if helper.IsHas(iSid) {
		_SearchDataset, _error := model.GetSearchDatasetsBySidAndLangcodeAndType(offset, limit, iSid,iLangcode,iType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_SearchDataset)
	}
	herr.Message = "Can't get to the GetSearchDatasetsBySidAndLangcodeAndType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSearchDatasetsBySidAndLangcodeAndDataHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSid := self.Args("sid").MustInt64()
	iLangcode := self.Args("langcode").String()
	iData := self.Args("data").String()

	if helper.IsHas(iSid) {
		_SearchDataset, _error := model.GetSearchDatasetsBySidAndLangcodeAndData(offset, limit, iSid,iLangcode,iData)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_SearchDataset)
	}
	herr.Message = "Can't get to the GetSearchDatasetsBySidAndLangcodeAndData's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSearchDatasetsBySidAndLangcodeAndReindexHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSid := self.Args("sid").MustInt64()
	iLangcode := self.Args("langcode").String()
	iReindex := self.Args("reindex").MustInt()

	if helper.IsHas(iSid) {
		_SearchDataset, _error := model.GetSearchDatasetsBySidAndLangcodeAndReindex(offset, limit, iSid,iLangcode,iReindex)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_SearchDataset)
	}
	herr.Message = "Can't get to the GetSearchDatasetsBySidAndLangcodeAndReindex's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSearchDatasetsBySidAndTypeAndDataHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSid := self.Args("sid").MustInt64()
	iType := self.Args("type").String()
	iData := self.Args("data").String()

	if helper.IsHas(iSid) {
		_SearchDataset, _error := model.GetSearchDatasetsBySidAndTypeAndData(offset, limit, iSid,iType,iData)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_SearchDataset)
	}
	herr.Message = "Can't get to the GetSearchDatasetsBySidAndTypeAndData's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSearchDatasetsBySidAndTypeAndReindexHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSid := self.Args("sid").MustInt64()
	iType := self.Args("type").String()
	iReindex := self.Args("reindex").MustInt()

	if helper.IsHas(iSid) {
		_SearchDataset, _error := model.GetSearchDatasetsBySidAndTypeAndReindex(offset, limit, iSid,iType,iReindex)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_SearchDataset)
	}
	herr.Message = "Can't get to the GetSearchDatasetsBySidAndTypeAndReindex's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSearchDatasetsBySidAndDataAndReindexHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSid := self.Args("sid").MustInt64()
	iData := self.Args("data").String()
	iReindex := self.Args("reindex").MustInt()

	if helper.IsHas(iSid) {
		_SearchDataset, _error := model.GetSearchDatasetsBySidAndDataAndReindex(offset, limit, iSid,iData,iReindex)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_SearchDataset)
	}
	herr.Message = "Can't get to the GetSearchDatasetsBySidAndDataAndReindex's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSearchDatasetsByLangcodeAndTypeAndDataHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iType := self.Args("type").String()
	iData := self.Args("data").String()

	if helper.IsHas(iLangcode) {
		_SearchDataset, _error := model.GetSearchDatasetsByLangcodeAndTypeAndData(offset, limit, iLangcode,iType,iData)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_SearchDataset)
	}
	herr.Message = "Can't get to the GetSearchDatasetsByLangcodeAndTypeAndData's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSearchDatasetsByLangcodeAndTypeAndReindexHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iType := self.Args("type").String()
	iReindex := self.Args("reindex").MustInt()

	if helper.IsHas(iLangcode) {
		_SearchDataset, _error := model.GetSearchDatasetsByLangcodeAndTypeAndReindex(offset, limit, iLangcode,iType,iReindex)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_SearchDataset)
	}
	herr.Message = "Can't get to the GetSearchDatasetsByLangcodeAndTypeAndReindex's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSearchDatasetsByLangcodeAndDataAndReindexHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iData := self.Args("data").String()
	iReindex := self.Args("reindex").MustInt()

	if helper.IsHas(iLangcode) {
		_SearchDataset, _error := model.GetSearchDatasetsByLangcodeAndDataAndReindex(offset, limit, iLangcode,iData,iReindex)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_SearchDataset)
	}
	herr.Message = "Can't get to the GetSearchDatasetsByLangcodeAndDataAndReindex's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSearchDatasetsByTypeAndDataAndReindexHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iType := self.Args("type").String()
	iData := self.Args("data").String()
	iReindex := self.Args("reindex").MustInt()

	if helper.IsHas(iType) {
		_SearchDataset, _error := model.GetSearchDatasetsByTypeAndDataAndReindex(offset, limit, iType,iData,iReindex)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_SearchDataset)
	}
	herr.Message = "Can't get to the GetSearchDatasetsByTypeAndDataAndReindex's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSearchDatasetsBySidAndLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSid := self.Args("sid").MustInt64()
	iLangcode := self.Args("langcode").String()

	if helper.IsHas(iSid) {
		_SearchDataset, _error := model.GetSearchDatasetsBySidAndLangcode(offset, limit, iSid,iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_SearchDataset)
	}
	herr.Message = "Can't get to the GetSearchDatasetsBySidAndLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSearchDatasetsBySidAndTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSid := self.Args("sid").MustInt64()
	iType := self.Args("type").String()

	if helper.IsHas(iSid) {
		_SearchDataset, _error := model.GetSearchDatasetsBySidAndType(offset, limit, iSid,iType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_SearchDataset)
	}
	herr.Message = "Can't get to the GetSearchDatasetsBySidAndType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSearchDatasetsBySidAndDataHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSid := self.Args("sid").MustInt64()
	iData := self.Args("data").String()

	if helper.IsHas(iSid) {
		_SearchDataset, _error := model.GetSearchDatasetsBySidAndData(offset, limit, iSid,iData)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_SearchDataset)
	}
	herr.Message = "Can't get to the GetSearchDatasetsBySidAndData's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSearchDatasetsBySidAndReindexHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSid := self.Args("sid").MustInt64()
	iReindex := self.Args("reindex").MustInt()

	if helper.IsHas(iSid) {
		_SearchDataset, _error := model.GetSearchDatasetsBySidAndReindex(offset, limit, iSid,iReindex)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_SearchDataset)
	}
	herr.Message = "Can't get to the GetSearchDatasetsBySidAndReindex's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSearchDatasetsByLangcodeAndTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iType := self.Args("type").String()

	if helper.IsHas(iLangcode) {
		_SearchDataset, _error := model.GetSearchDatasetsByLangcodeAndType(offset, limit, iLangcode,iType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_SearchDataset)
	}
	herr.Message = "Can't get to the GetSearchDatasetsByLangcodeAndType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSearchDatasetsByLangcodeAndDataHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iData := self.Args("data").String()

	if helper.IsHas(iLangcode) {
		_SearchDataset, _error := model.GetSearchDatasetsByLangcodeAndData(offset, limit, iLangcode,iData)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_SearchDataset)
	}
	herr.Message = "Can't get to the GetSearchDatasetsByLangcodeAndData's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSearchDatasetsByLangcodeAndReindexHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLangcode := self.Args("langcode").String()
	iReindex := self.Args("reindex").MustInt()

	if helper.IsHas(iLangcode) {
		_SearchDataset, _error := model.GetSearchDatasetsByLangcodeAndReindex(offset, limit, iLangcode,iReindex)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_SearchDataset)
	}
	herr.Message = "Can't get to the GetSearchDatasetsByLangcodeAndReindex's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSearchDatasetsByTypeAndDataHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iType := self.Args("type").String()
	iData := self.Args("data").String()

	if helper.IsHas(iType) {
		_SearchDataset, _error := model.GetSearchDatasetsByTypeAndData(offset, limit, iType,iData)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_SearchDataset)
	}
	herr.Message = "Can't get to the GetSearchDatasetsByTypeAndData's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSearchDatasetsByTypeAndReindexHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iType := self.Args("type").String()
	iReindex := self.Args("reindex").MustInt()

	if helper.IsHas(iType) {
		_SearchDataset, _error := model.GetSearchDatasetsByTypeAndReindex(offset, limit, iType,iReindex)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_SearchDataset)
	}
	herr.Message = "Can't get to the GetSearchDatasetsByTypeAndReindex's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSearchDatasetsByDataAndReindexHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iData := self.Args("data").String()
	iReindex := self.Args("reindex").MustInt()

	if helper.IsHas(iData) {
		_SearchDataset, _error := model.GetSearchDatasetsByDataAndReindex(offset, limit, iData,iReindex)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_SearchDataset)
	}
	herr.Message = "Can't get to the GetSearchDatasetsByDataAndReindex's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSearchDatasetsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	if (offset > 0) && (len(field) > 0) {
		_SearchDataset, _error := model.GetSearchDatasets(offset, limit, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_SearchDataset)
	}
	herr.Message = "Can't get to the GetSearchDatasets' args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasSearchDatasetViaSidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iSid := self.Args("sid").MustInt64()
	if helper.IsHas(iSid) {
		_SearchDataset := model.HasSearchDatasetViaSid(iSid)
		var m = map[string]interface{}{}
		m["search_dataset"] = _SearchDataset
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasSearchDatasetViaSid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasSearchDatasetViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iLangcode := self.Args("langcode").String()
	if helper.IsHas(iLangcode) {
		_SearchDataset := model.HasSearchDatasetViaLangcode(iLangcode)
		var m = map[string]interface{}{}
		m["search_dataset"] = _SearchDataset
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasSearchDatasetViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasSearchDatasetViaTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iType := self.Args("type").String()
	if helper.IsHas(iType) {
		_SearchDataset := model.HasSearchDatasetViaType(iType)
		var m = map[string]interface{}{}
		m["search_dataset"] = _SearchDataset
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasSearchDatasetViaType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasSearchDatasetViaDataHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iData := self.Args("data").String()
	if helper.IsHas(iData) {
		_SearchDataset := model.HasSearchDatasetViaData(iData)
		var m = map[string]interface{}{}
		m["search_dataset"] = _SearchDataset
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasSearchDatasetViaData's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasSearchDatasetViaReindexHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iReindex := self.Args("reindex").MustInt()
	if helper.IsHas(iReindex) {
		_SearchDataset := model.HasSearchDatasetViaReindex(iReindex)
		var m = map[string]interface{}{}
		m["search_dataset"] = _SearchDataset
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasSearchDatasetViaReindex's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSearchDatasetViaSidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iSid := self.Args("sid").MustInt64()
	if helper.IsHas(iSid) {
		_SearchDataset, _error := model.GetSearchDatasetViaSid(iSid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_SearchDataset)
	}
	herr.Message = "Can't get to the GetSearchDatasetViaSid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSearchDatasetViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iLangcode := self.Args("langcode").String()
	if helper.IsHas(iLangcode) {
		_SearchDataset, _error := model.GetSearchDatasetViaLangcode(iLangcode)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_SearchDataset)
	}
	herr.Message = "Can't get to the GetSearchDatasetViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSearchDatasetViaTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iType := self.Args("type").String()
	if helper.IsHas(iType) {
		_SearchDataset, _error := model.GetSearchDatasetViaType(iType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_SearchDataset)
	}
	herr.Message = "Can't get to the GetSearchDatasetViaType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSearchDatasetViaDataHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iData := self.Args("data").String()
	if helper.IsHas(iData) {
		_SearchDataset, _error := model.GetSearchDatasetViaData(iData)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_SearchDataset)
	}
	herr.Message = "Can't get to the GetSearchDatasetViaData's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSearchDatasetViaReindexHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iReindex := self.Args("reindex").MustInt()
	if helper.IsHas(iReindex) {
		_SearchDataset, _error := model.GetSearchDatasetViaReindex(iReindex)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_SearchDataset)
	}
	herr.Message = "Can't get to the GetSearchDatasetViaReindex's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetSearchDatasetViaSidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Sid_ := self.Args("sid").MustInt64()
	if helper.IsHas(Sid_) {
		var iSearchDataset model.SearchDataset
		self.Bind(&iSearchDataset)
		_SearchDataset, _error := model.SetSearchDatasetViaSid(Sid_, &iSearchDataset)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_SearchDataset)
	}
	herr.Message = "Can't get to the SetSearchDatasetViaSid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetSearchDatasetViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	if helper.IsHas(Langcode_) {
		var iSearchDataset model.SearchDataset
		self.Bind(&iSearchDataset)
		_SearchDataset, _error := model.SetSearchDatasetViaLangcode(Langcode_, &iSearchDataset)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_SearchDataset)
	}
	herr.Message = "Can't get to the SetSearchDatasetViaLangcode's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetSearchDatasetViaTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Type_ := self.Args("type").String()
	if helper.IsHas(Type_) {
		var iSearchDataset model.SearchDataset
		self.Bind(&iSearchDataset)
		_SearchDataset, _error := model.SetSearchDatasetViaType(Type_, &iSearchDataset)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_SearchDataset)
	}
	herr.Message = "Can't get to the SetSearchDatasetViaType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetSearchDatasetViaDataHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Data_ := self.Args("data").String()
	if helper.IsHas(Data_) {
		var iSearchDataset model.SearchDataset
		self.Bind(&iSearchDataset)
		_SearchDataset, _error := model.SetSearchDatasetViaData(Data_, &iSearchDataset)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_SearchDataset)
	}
	herr.Message = "Can't get to the SetSearchDatasetViaData's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetSearchDatasetViaReindexHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Reindex_ := self.Args("reindex").MustInt()
	if helper.IsHas(Reindex_) {
		var iSearchDataset model.SearchDataset
		self.Bind(&iSearchDataset)
		_SearchDataset, _error := model.SetSearchDatasetViaReindex(Reindex_, &iSearchDataset)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_SearchDataset)
	}
	herr.Message = "Can't get to the SetSearchDatasetViaReindex's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostAddSearchDatasetHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Sid_ := self.Args("sid").MustInt64()
	Langcode_ := self.Args("langcode").String()
	Type_ := self.Args("type").String()
	Data_ := self.Args("data").String()
	Reindex_ := self.Args("reindex").MustInt()

	if helper.IsHas(Sid_) {
		_error := model.AddSearchDataset(Sid_,Langcode_,Type_,Data_,Reindex_)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		herr.Message = "StatusOK"
		herr.Status = macross.StatusOK
		return self.JSON(herr)
	}
	herr.Message = "Can't get to the AddSearchDataset's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSearchDatasetHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iSearchDataset model.SearchDataset
	self.Bind(&iSearchDataset)
	_int64, _error := model.PostSearchDataset(&iSearchDataset)
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

func PutSearchDatasetHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iSearchDataset model.SearchDataset
	self.Bind(&iSearchDataset)
	_int64, _error := model.PutSearchDataset(&iSearchDataset)
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

func PutSearchDatasetViaSidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Sid_ := self.Args("sid").MustInt64()
	var iSearchDataset model.SearchDataset
	self.Bind(&iSearchDataset)
	_int64, _error := model.PutSearchDatasetViaSid(Sid_, &iSearchDataset)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutSearchDatasetViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	var iSearchDataset model.SearchDataset
	self.Bind(&iSearchDataset)
	_int64, _error := model.PutSearchDatasetViaLangcode(Langcode_, &iSearchDataset)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutSearchDatasetViaTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Type_ := self.Args("type").String()
	var iSearchDataset model.SearchDataset
	self.Bind(&iSearchDataset)
	_int64, _error := model.PutSearchDatasetViaType(Type_, &iSearchDataset)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutSearchDatasetViaDataHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Data_ := self.Args("data").String()
	var iSearchDataset model.SearchDataset
	self.Bind(&iSearchDataset)
	_int64, _error := model.PutSearchDatasetViaData(Data_, &iSearchDataset)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutSearchDatasetViaReindexHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Reindex_ := self.Args("reindex").MustInt()
	var iSearchDataset model.SearchDataset
	self.Bind(&iSearchDataset)
	_int64, _error := model.PutSearchDatasetViaReindex(Reindex_, &iSearchDataset)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateSearchDatasetViaSidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Sid_ := self.Args("sid").MustInt64()
	var iSearchDataset model.SearchDataset
	self.Bind(&iSearchDataset)
	var iMap = helper.StructToMap(iSearchDataset)
	_error := model.UpdateSearchDatasetViaSid(Sid_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateSearchDatasetViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	var iSearchDataset model.SearchDataset
	self.Bind(&iSearchDataset)
	var iMap = helper.StructToMap(iSearchDataset)
	_error := model.UpdateSearchDatasetViaLangcode(Langcode_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateSearchDatasetViaTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Type_ := self.Args("type").String()
	var iSearchDataset model.SearchDataset
	self.Bind(&iSearchDataset)
	var iMap = helper.StructToMap(iSearchDataset)
	_error := model.UpdateSearchDatasetViaType(Type_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateSearchDatasetViaDataHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Data_ := self.Args("data").String()
	var iSearchDataset model.SearchDataset
	self.Bind(&iSearchDataset)
	var iMap = helper.StructToMap(iSearchDataset)
	_error := model.UpdateSearchDatasetViaData(Data_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateSearchDatasetViaReindexHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Reindex_ := self.Args("reindex").MustInt()
	var iSearchDataset model.SearchDataset
	self.Bind(&iSearchDataset)
	var iMap = helper.StructToMap(iSearchDataset)
	_error := model.UpdateSearchDatasetViaReindex(Reindex_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteSearchDatasetHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Sid_ := self.Args("sid").MustInt64()
	_int64, _error := model.DeleteSearchDataset(Sid_)
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

func DeleteSearchDatasetViaSidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Sid_ := self.Args("sid").MustInt64()
	_error := model.DeleteSearchDatasetViaSid(Sid_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteSearchDatasetViaLangcodeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Langcode_ := self.Args("langcode").String()
	_error := model.DeleteSearchDatasetViaLangcode(Langcode_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteSearchDatasetViaTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Type_ := self.Args("type").String()
	_error := model.DeleteSearchDatasetViaType(Type_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteSearchDatasetViaDataHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Data_ := self.Args("data").String()
	_error := model.DeleteSearchDatasetViaData(Data_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteSearchDatasetViaReindexHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Reindex_ := self.Args("reindex").MustInt()
	_error := model.DeleteSearchDatasetViaReindex(Reindex_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}
