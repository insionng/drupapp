package handler

import (
	"drupapp/helper"
	"drupapp/model"
	"github.com/insionng/macross"
)

func GetTaxonomyIndexsCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()

	if offset > 0 {
		var m = map[string]interface{}{}
		_int64, _error := model.GetTaxonomyIndexsCount(offset, limit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		m["taxonomy_indexsCount"] = _int64
		return self.JSON(m)
	}
	herr.Message = "Can't get to the GetTaxonomyIndexsCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyIndexCountViaNidHandler(self *macross.Context) error {
	Nid_ := self.Args("nid").MustInt64()
	_int64 := model.GetTaxonomyIndexCountViaNid(Nid_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["taxonomy_indexCount"] = 0
	}
	m["taxonomy_indexCount"] = _int64
	return self.JSON(m)
}

func GetTaxonomyIndexCountViaTidHandler(self *macross.Context) error {
	Tid_ := self.Args("tid").MustInt()
	_int64 := model.GetTaxonomyIndexCountViaTid(Tid_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["taxonomy_indexCount"] = 0
	}
	m["taxonomy_indexCount"] = _int64
	return self.JSON(m)
}

func GetTaxonomyIndexCountViaStatusHandler(self *macross.Context) error {
	Status_ := self.Args("status").MustInt()
	_int64 := model.GetTaxonomyIndexCountViaStatus(Status_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["taxonomy_indexCount"] = 0
	}
	m["taxonomy_indexCount"] = _int64
	return self.JSON(m)
}

func GetTaxonomyIndexCountViaStickyHandler(self *macross.Context) error {
	Sticky_ := self.Args("sticky").MustInt()
	_int64 := model.GetTaxonomyIndexCountViaSticky(Sticky_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["taxonomy_indexCount"] = 0
	}
	m["taxonomy_indexCount"] = _int64
	return self.JSON(m)
}

func GetTaxonomyIndexCountViaCreatedHandler(self *macross.Context) error {
	Created_ := self.Args("created").MustInt()
	_int64 := model.GetTaxonomyIndexCountViaCreated(Created_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["taxonomy_indexCount"] = 0
	}
	m["taxonomy_indexCount"] = _int64
	return self.JSON(m)
}

func GetTaxonomyIndexsViaNidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iNid := self.Args("nid").MustInt64()
	if (offset > 0) && helper.IsHas(iNid) {
		_TaxonomyIndex, _error := model.GetTaxonomyIndexsViaNid(offset, limit, iNid, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyIndex)
	}
	herr.Message = "Can't get to the GetTaxonomyIndexsViaNid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyIndexsViaTidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iTid := self.Args("tid").MustInt()
	if (offset > 0) && helper.IsHas(iTid) {
		_TaxonomyIndex, _error := model.GetTaxonomyIndexsViaTid(offset, limit, iTid, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyIndex)
	}
	herr.Message = "Can't get to the GetTaxonomyIndexsViaTid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyIndexsViaStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iStatus := self.Args("status").MustInt()
	if (offset > 0) && helper.IsHas(iStatus) {
		_TaxonomyIndex, _error := model.GetTaxonomyIndexsViaStatus(offset, limit, iStatus, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyIndex)
	}
	herr.Message = "Can't get to the GetTaxonomyIndexsViaStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyIndexsViaStickyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iSticky := self.Args("sticky").MustInt()
	if (offset > 0) && helper.IsHas(iSticky) {
		_TaxonomyIndex, _error := model.GetTaxonomyIndexsViaSticky(offset, limit, iSticky, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyIndex)
	}
	herr.Message = "Can't get to the GetTaxonomyIndexsViaSticky's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyIndexsViaCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iCreated := self.Args("created").MustInt()
	if (offset > 0) && helper.IsHas(iCreated) {
		_TaxonomyIndex, _error := model.GetTaxonomyIndexsViaCreated(offset, limit, iCreated, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyIndex)
	}
	herr.Message = "Can't get to the GetTaxonomyIndexsViaCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyIndexsByNidAndTidAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iTid := self.Args("tid").MustInt()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iNid) {
		_TaxonomyIndex, _error := model.GetTaxonomyIndexsByNidAndTidAndStatus(offset, limit, iNid,iTid,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyIndex)
	}
	herr.Message = "Can't get to the GetTaxonomyIndexsByNidAndTidAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyIndexsByNidAndTidAndStickyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iTid := self.Args("tid").MustInt()
	iSticky := self.Args("sticky").MustInt()

	if helper.IsHas(iNid) {
		_TaxonomyIndex, _error := model.GetTaxonomyIndexsByNidAndTidAndSticky(offset, limit, iNid,iTid,iSticky)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyIndex)
	}
	herr.Message = "Can't get to the GetTaxonomyIndexsByNidAndTidAndSticky's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyIndexsByNidAndTidAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iTid := self.Args("tid").MustInt()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iNid) {
		_TaxonomyIndex, _error := model.GetTaxonomyIndexsByNidAndTidAndCreated(offset, limit, iNid,iTid,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyIndex)
	}
	herr.Message = "Can't get to the GetTaxonomyIndexsByNidAndTidAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyIndexsByNidAndStatusAndStickyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iStatus := self.Args("status").MustInt()
	iSticky := self.Args("sticky").MustInt()

	if helper.IsHas(iNid) {
		_TaxonomyIndex, _error := model.GetTaxonomyIndexsByNidAndStatusAndSticky(offset, limit, iNid,iStatus,iSticky)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyIndex)
	}
	herr.Message = "Can't get to the GetTaxonomyIndexsByNidAndStatusAndSticky's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyIndexsByNidAndStatusAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iStatus := self.Args("status").MustInt()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iNid) {
		_TaxonomyIndex, _error := model.GetTaxonomyIndexsByNidAndStatusAndCreated(offset, limit, iNid,iStatus,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyIndex)
	}
	herr.Message = "Can't get to the GetTaxonomyIndexsByNidAndStatusAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyIndexsByNidAndStickyAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iSticky := self.Args("sticky").MustInt()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iNid) {
		_TaxonomyIndex, _error := model.GetTaxonomyIndexsByNidAndStickyAndCreated(offset, limit, iNid,iSticky,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyIndex)
	}
	herr.Message = "Can't get to the GetTaxonomyIndexsByNidAndStickyAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyIndexsByTidAndStatusAndStickyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTid := self.Args("tid").MustInt()
	iStatus := self.Args("status").MustInt()
	iSticky := self.Args("sticky").MustInt()

	if helper.IsHas(iTid) {
		_TaxonomyIndex, _error := model.GetTaxonomyIndexsByTidAndStatusAndSticky(offset, limit, iTid,iStatus,iSticky)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyIndex)
	}
	herr.Message = "Can't get to the GetTaxonomyIndexsByTidAndStatusAndSticky's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyIndexsByTidAndStatusAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTid := self.Args("tid").MustInt()
	iStatus := self.Args("status").MustInt()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iTid) {
		_TaxonomyIndex, _error := model.GetTaxonomyIndexsByTidAndStatusAndCreated(offset, limit, iTid,iStatus,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyIndex)
	}
	herr.Message = "Can't get to the GetTaxonomyIndexsByTidAndStatusAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyIndexsByTidAndStickyAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTid := self.Args("tid").MustInt()
	iSticky := self.Args("sticky").MustInt()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iTid) {
		_TaxonomyIndex, _error := model.GetTaxonomyIndexsByTidAndStickyAndCreated(offset, limit, iTid,iSticky,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyIndex)
	}
	herr.Message = "Can't get to the GetTaxonomyIndexsByTidAndStickyAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyIndexsByStatusAndStickyAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iStatus := self.Args("status").MustInt()
	iSticky := self.Args("sticky").MustInt()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iStatus) {
		_TaxonomyIndex, _error := model.GetTaxonomyIndexsByStatusAndStickyAndCreated(offset, limit, iStatus,iSticky,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyIndex)
	}
	herr.Message = "Can't get to the GetTaxonomyIndexsByStatusAndStickyAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyIndexsByNidAndTidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iTid := self.Args("tid").MustInt()

	if helper.IsHas(iNid) {
		_TaxonomyIndex, _error := model.GetTaxonomyIndexsByNidAndTid(offset, limit, iNid,iTid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyIndex)
	}
	herr.Message = "Can't get to the GetTaxonomyIndexsByNidAndTid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyIndexsByNidAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iNid) {
		_TaxonomyIndex, _error := model.GetTaxonomyIndexsByNidAndStatus(offset, limit, iNid,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyIndex)
	}
	herr.Message = "Can't get to the GetTaxonomyIndexsByNidAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyIndexsByNidAndStickyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iSticky := self.Args("sticky").MustInt()

	if helper.IsHas(iNid) {
		_TaxonomyIndex, _error := model.GetTaxonomyIndexsByNidAndSticky(offset, limit, iNid,iSticky)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyIndex)
	}
	herr.Message = "Can't get to the GetTaxonomyIndexsByNidAndSticky's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyIndexsByNidAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iNid := self.Args("nid").MustInt64()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iNid) {
		_TaxonomyIndex, _error := model.GetTaxonomyIndexsByNidAndCreated(offset, limit, iNid,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyIndex)
	}
	herr.Message = "Can't get to the GetTaxonomyIndexsByNidAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyIndexsByTidAndStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTid := self.Args("tid").MustInt()
	iStatus := self.Args("status").MustInt()

	if helper.IsHas(iTid) {
		_TaxonomyIndex, _error := model.GetTaxonomyIndexsByTidAndStatus(offset, limit, iTid,iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyIndex)
	}
	herr.Message = "Can't get to the GetTaxonomyIndexsByTidAndStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyIndexsByTidAndStickyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTid := self.Args("tid").MustInt()
	iSticky := self.Args("sticky").MustInt()

	if helper.IsHas(iTid) {
		_TaxonomyIndex, _error := model.GetTaxonomyIndexsByTidAndSticky(offset, limit, iTid,iSticky)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyIndex)
	}
	herr.Message = "Can't get to the GetTaxonomyIndexsByTidAndSticky's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyIndexsByTidAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTid := self.Args("tid").MustInt()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iTid) {
		_TaxonomyIndex, _error := model.GetTaxonomyIndexsByTidAndCreated(offset, limit, iTid,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyIndex)
	}
	herr.Message = "Can't get to the GetTaxonomyIndexsByTidAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyIndexsByStatusAndStickyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iStatus := self.Args("status").MustInt()
	iSticky := self.Args("sticky").MustInt()

	if helper.IsHas(iStatus) {
		_TaxonomyIndex, _error := model.GetTaxonomyIndexsByStatusAndSticky(offset, limit, iStatus,iSticky)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyIndex)
	}
	herr.Message = "Can't get to the GetTaxonomyIndexsByStatusAndSticky's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyIndexsByStatusAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iStatus := self.Args("status").MustInt()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iStatus) {
		_TaxonomyIndex, _error := model.GetTaxonomyIndexsByStatusAndCreated(offset, limit, iStatus,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyIndex)
	}
	herr.Message = "Can't get to the GetTaxonomyIndexsByStatusAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyIndexsByStickyAndCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSticky := self.Args("sticky").MustInt()
	iCreated := self.Args("created").MustInt()

	if helper.IsHas(iSticky) {
		_TaxonomyIndex, _error := model.GetTaxonomyIndexsByStickyAndCreated(offset, limit, iSticky,iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyIndex)
	}
	herr.Message = "Can't get to the GetTaxonomyIndexsByStickyAndCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyIndexsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	if (offset > 0) && (len(field) > 0) {
		_TaxonomyIndex, _error := model.GetTaxonomyIndexs(offset, limit, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyIndex)
	}
	herr.Message = "Can't get to the GetTaxonomyIndexs' args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasTaxonomyIndexViaNidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iNid := self.Args("nid").MustInt64()
	if helper.IsHas(iNid) {
		_TaxonomyIndex := model.HasTaxonomyIndexViaNid(iNid)
		var m = map[string]interface{}{}
		m["taxonomy_index"] = _TaxonomyIndex
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasTaxonomyIndexViaNid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasTaxonomyIndexViaTidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iTid := self.Args("tid").MustInt()
	if helper.IsHas(iTid) {
		_TaxonomyIndex := model.HasTaxonomyIndexViaTid(iTid)
		var m = map[string]interface{}{}
		m["taxonomy_index"] = _TaxonomyIndex
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasTaxonomyIndexViaTid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasTaxonomyIndexViaStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iStatus := self.Args("status").MustInt()
	if helper.IsHas(iStatus) {
		_TaxonomyIndex := model.HasTaxonomyIndexViaStatus(iStatus)
		var m = map[string]interface{}{}
		m["taxonomy_index"] = _TaxonomyIndex
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasTaxonomyIndexViaStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasTaxonomyIndexViaStickyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iSticky := self.Args("sticky").MustInt()
	if helper.IsHas(iSticky) {
		_TaxonomyIndex := model.HasTaxonomyIndexViaSticky(iSticky)
		var m = map[string]interface{}{}
		m["taxonomy_index"] = _TaxonomyIndex
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasTaxonomyIndexViaSticky's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasTaxonomyIndexViaCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iCreated := self.Args("created").MustInt()
	if helper.IsHas(iCreated) {
		_TaxonomyIndex := model.HasTaxonomyIndexViaCreated(iCreated)
		var m = map[string]interface{}{}
		m["taxonomy_index"] = _TaxonomyIndex
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasTaxonomyIndexViaCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyIndexViaNidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iNid := self.Args("nid").MustInt64()
	if helper.IsHas(iNid) {
		_TaxonomyIndex, _error := model.GetTaxonomyIndexViaNid(iNid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyIndex)
	}
	herr.Message = "Can't get to the GetTaxonomyIndexViaNid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyIndexViaTidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iTid := self.Args("tid").MustInt()
	if helper.IsHas(iTid) {
		_TaxonomyIndex, _error := model.GetTaxonomyIndexViaTid(iTid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyIndex)
	}
	herr.Message = "Can't get to the GetTaxonomyIndexViaTid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyIndexViaStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iStatus := self.Args("status").MustInt()
	if helper.IsHas(iStatus) {
		_TaxonomyIndex, _error := model.GetTaxonomyIndexViaStatus(iStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyIndex)
	}
	herr.Message = "Can't get to the GetTaxonomyIndexViaStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyIndexViaStickyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iSticky := self.Args("sticky").MustInt()
	if helper.IsHas(iSticky) {
		_TaxonomyIndex, _error := model.GetTaxonomyIndexViaSticky(iSticky)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyIndex)
	}
	herr.Message = "Can't get to the GetTaxonomyIndexViaSticky's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyIndexViaCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iCreated := self.Args("created").MustInt()
	if helper.IsHas(iCreated) {
		_TaxonomyIndex, _error := model.GetTaxonomyIndexViaCreated(iCreated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyIndex)
	}
	herr.Message = "Can't get to the GetTaxonomyIndexViaCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetTaxonomyIndexViaNidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Nid_ := self.Args("nid").MustInt64()
	if helper.IsHas(Nid_) {
		var iTaxonomyIndex model.TaxonomyIndex
		self.Bind(&iTaxonomyIndex)
		_TaxonomyIndex, _error := model.SetTaxonomyIndexViaNid(Nid_, &iTaxonomyIndex)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyIndex)
	}
	herr.Message = "Can't get to the SetTaxonomyIndexViaNid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetTaxonomyIndexViaTidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Tid_ := self.Args("tid").MustInt()
	if helper.IsHas(Tid_) {
		var iTaxonomyIndex model.TaxonomyIndex
		self.Bind(&iTaxonomyIndex)
		_TaxonomyIndex, _error := model.SetTaxonomyIndexViaTid(Tid_, &iTaxonomyIndex)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyIndex)
	}
	herr.Message = "Can't get to the SetTaxonomyIndexViaTid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetTaxonomyIndexViaStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Status_ := self.Args("status").MustInt()
	if helper.IsHas(Status_) {
		var iTaxonomyIndex model.TaxonomyIndex
		self.Bind(&iTaxonomyIndex)
		_TaxonomyIndex, _error := model.SetTaxonomyIndexViaStatus(Status_, &iTaxonomyIndex)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyIndex)
	}
	herr.Message = "Can't get to the SetTaxonomyIndexViaStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetTaxonomyIndexViaStickyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Sticky_ := self.Args("sticky").MustInt()
	if helper.IsHas(Sticky_) {
		var iTaxonomyIndex model.TaxonomyIndex
		self.Bind(&iTaxonomyIndex)
		_TaxonomyIndex, _error := model.SetTaxonomyIndexViaSticky(Sticky_, &iTaxonomyIndex)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyIndex)
	}
	herr.Message = "Can't get to the SetTaxonomyIndexViaSticky's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetTaxonomyIndexViaCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Created_ := self.Args("created").MustInt()
	if helper.IsHas(Created_) {
		var iTaxonomyIndex model.TaxonomyIndex
		self.Bind(&iTaxonomyIndex)
		_TaxonomyIndex, _error := model.SetTaxonomyIndexViaCreated(Created_, &iTaxonomyIndex)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyIndex)
	}
	herr.Message = "Can't get to the SetTaxonomyIndexViaCreated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostAddTaxonomyIndexHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Nid_ := self.Args("nid").MustInt64()
	Tid_ := self.Args("tid").MustInt()
	Status_ := self.Args("status").MustInt()
	Sticky_ := self.Args("sticky").MustInt()
	Created_ := self.Args("created").MustInt()

	if helper.IsHas(Nid_) {
		_error := model.AddTaxonomyIndex(Nid_,Tid_,Status_,Sticky_,Created_)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		herr.Message = "StatusOK"
		herr.Status = macross.StatusOK
		return self.JSON(herr)
	}
	herr.Message = "Can't get to the AddTaxonomyIndex's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostTaxonomyIndexHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iTaxonomyIndex model.TaxonomyIndex
	self.Bind(&iTaxonomyIndex)
	_int64, _error := model.PostTaxonomyIndex(&iTaxonomyIndex)
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

func PutTaxonomyIndexHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iTaxonomyIndex model.TaxonomyIndex
	self.Bind(&iTaxonomyIndex)
	_int64, _error := model.PutTaxonomyIndex(&iTaxonomyIndex)
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

func PutTaxonomyIndexViaNidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Nid_ := self.Args("nid").MustInt64()
	var iTaxonomyIndex model.TaxonomyIndex
	self.Bind(&iTaxonomyIndex)
	_int64, _error := model.PutTaxonomyIndexViaNid(Nid_, &iTaxonomyIndex)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutTaxonomyIndexViaTidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Tid_ := self.Args("tid").MustInt()
	var iTaxonomyIndex model.TaxonomyIndex
	self.Bind(&iTaxonomyIndex)
	_int64, _error := model.PutTaxonomyIndexViaTid(Tid_, &iTaxonomyIndex)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutTaxonomyIndexViaStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Status_ := self.Args("status").MustInt()
	var iTaxonomyIndex model.TaxonomyIndex
	self.Bind(&iTaxonomyIndex)
	_int64, _error := model.PutTaxonomyIndexViaStatus(Status_, &iTaxonomyIndex)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutTaxonomyIndexViaStickyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Sticky_ := self.Args("sticky").MustInt()
	var iTaxonomyIndex model.TaxonomyIndex
	self.Bind(&iTaxonomyIndex)
	_int64, _error := model.PutTaxonomyIndexViaSticky(Sticky_, &iTaxonomyIndex)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutTaxonomyIndexViaCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Created_ := self.Args("created").MustInt()
	var iTaxonomyIndex model.TaxonomyIndex
	self.Bind(&iTaxonomyIndex)
	_int64, _error := model.PutTaxonomyIndexViaCreated(Created_, &iTaxonomyIndex)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateTaxonomyIndexViaNidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Nid_ := self.Args("nid").MustInt64()
	var iTaxonomyIndex model.TaxonomyIndex
	self.Bind(&iTaxonomyIndex)
	var iMap = helper.StructToMap(iTaxonomyIndex)
	_error := model.UpdateTaxonomyIndexViaNid(Nid_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateTaxonomyIndexViaTidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Tid_ := self.Args("tid").MustInt()
	var iTaxonomyIndex model.TaxonomyIndex
	self.Bind(&iTaxonomyIndex)
	var iMap = helper.StructToMap(iTaxonomyIndex)
	_error := model.UpdateTaxonomyIndexViaTid(Tid_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateTaxonomyIndexViaStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Status_ := self.Args("status").MustInt()
	var iTaxonomyIndex model.TaxonomyIndex
	self.Bind(&iTaxonomyIndex)
	var iMap = helper.StructToMap(iTaxonomyIndex)
	_error := model.UpdateTaxonomyIndexViaStatus(Status_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateTaxonomyIndexViaStickyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Sticky_ := self.Args("sticky").MustInt()
	var iTaxonomyIndex model.TaxonomyIndex
	self.Bind(&iTaxonomyIndex)
	var iMap = helper.StructToMap(iTaxonomyIndex)
	_error := model.UpdateTaxonomyIndexViaSticky(Sticky_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateTaxonomyIndexViaCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Created_ := self.Args("created").MustInt()
	var iTaxonomyIndex model.TaxonomyIndex
	self.Bind(&iTaxonomyIndex)
	var iMap = helper.StructToMap(iTaxonomyIndex)
	_error := model.UpdateTaxonomyIndexViaCreated(Created_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteTaxonomyIndexHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Nid_ := self.Args("nid").MustInt64()
	_int64, _error := model.DeleteTaxonomyIndex(Nid_)
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

func DeleteTaxonomyIndexViaNidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Nid_ := self.Args("nid").MustInt64()
	_error := model.DeleteTaxonomyIndexViaNid(Nid_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteTaxonomyIndexViaTidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Tid_ := self.Args("tid").MustInt()
	_error := model.DeleteTaxonomyIndexViaTid(Tid_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteTaxonomyIndexViaStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Status_ := self.Args("status").MustInt()
	_error := model.DeleteTaxonomyIndexViaStatus(Status_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteTaxonomyIndexViaStickyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Sticky_ := self.Args("sticky").MustInt()
	_error := model.DeleteTaxonomyIndexViaSticky(Sticky_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteTaxonomyIndexViaCreatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Created_ := self.Args("created").MustInt()
	_error := model.DeleteTaxonomyIndexViaCreated(Created_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}
