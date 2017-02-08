package handler

import (
	"drupapp/helper"
	"drupapp/model"
	"github.com/insionng/macross"
)

func GetTaxonomyTermHierarchiesCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()

	if offset > 0 {
		var m = map[string]interface{}{}
		_int64, _error := model.GetTaxonomyTermHierarchiesCount(offset, limit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		m["taxonomy_term_hierarchysCount"] = _int64
		return self.JSON(m)
	}
	herr.Message = "Can't get to the GetTaxonomyTermHierarchiesCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermHierarchyCountViaTidHandler(self *macross.Context) error {
	Tid_ := self.Args("tid").MustInt64()
	_int64 := model.GetTaxonomyTermHierarchyCountViaTid(Tid_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["taxonomy_term_hierarchyCount"] = 0
	}
	m["taxonomy_term_hierarchyCount"] = _int64
	return self.JSON(m)
}

func GetTaxonomyTermHierarchyCountViaParentHandler(self *macross.Context) error {
	Parent_ := self.Args("parent").MustInt()
	_int64 := model.GetTaxonomyTermHierarchyCountViaParent(Parent_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["taxonomy_term_hierarchyCount"] = 0
	}
	m["taxonomy_term_hierarchyCount"] = _int64
	return self.JSON(m)
}

func GetTaxonomyTermHierarchiesViaTidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iTid := self.Args("tid").MustInt64()
	if (offset > 0) && helper.IsHas(iTid) {
		_TaxonomyTermHierarchy, _error := model.GetTaxonomyTermHierarchiesViaTid(offset, limit, iTid, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermHierarchy)
	}
	herr.Message = "Can't get to the GetTaxonomyTermHierarchiesViaTid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermHierarchiesViaParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iParent := self.Args("parent").MustInt()
	if (offset > 0) && helper.IsHas(iParent) {
		_TaxonomyTermHierarchy, _error := model.GetTaxonomyTermHierarchiesViaParent(offset, limit, iParent, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermHierarchy)
	}
	herr.Message = "Can't get to the GetTaxonomyTermHierarchiesViaParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermHierarchiesByTidAndParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTid := self.Args("tid").MustInt64()
	iParent := self.Args("parent").MustInt()

	if helper.IsHas(iTid) {
		_TaxonomyTermHierarchy, _error := model.GetTaxonomyTermHierarchiesByTidAndParent(offset, limit, iTid,iParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermHierarchy)
	}
	herr.Message = "Can't get to the GetTaxonomyTermHierarchiesByTidAndParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermHierarchiesHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	if (offset > 0) && (len(field) > 0) {
		_TaxonomyTermHierarchy, _error := model.GetTaxonomyTermHierarchies(offset, limit, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermHierarchy)
	}
	herr.Message = "Can't get to the GetTaxonomyTermHierarchies' args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasTaxonomyTermHierarchyViaTidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iTid := self.Args("tid").MustInt64()
	if helper.IsHas(iTid) {
		_TaxonomyTermHierarchy := model.HasTaxonomyTermHierarchyViaTid(iTid)
		var m = map[string]interface{}{}
		m["taxonomy_term_hierarchy"] = _TaxonomyTermHierarchy
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasTaxonomyTermHierarchyViaTid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasTaxonomyTermHierarchyViaParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iParent := self.Args("parent").MustInt()
	if helper.IsHas(iParent) {
		_TaxonomyTermHierarchy := model.HasTaxonomyTermHierarchyViaParent(iParent)
		var m = map[string]interface{}{}
		m["taxonomy_term_hierarchy"] = _TaxonomyTermHierarchy
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasTaxonomyTermHierarchyViaParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermHierarchyViaTidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iTid := self.Args("tid").MustInt64()
	if helper.IsHas(iTid) {
		_TaxonomyTermHierarchy, _error := model.GetTaxonomyTermHierarchyViaTid(iTid)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermHierarchy)
	}
	herr.Message = "Can't get to the GetTaxonomyTermHierarchyViaTid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTaxonomyTermHierarchyViaParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iParent := self.Args("parent").MustInt()
	if helper.IsHas(iParent) {
		_TaxonomyTermHierarchy, _error := model.GetTaxonomyTermHierarchyViaParent(iParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermHierarchy)
	}
	herr.Message = "Can't get to the GetTaxonomyTermHierarchyViaParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetTaxonomyTermHierarchyViaTidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Tid_ := self.Args("tid").MustInt64()
	if helper.IsHas(Tid_) {
		var iTaxonomyTermHierarchy model.TaxonomyTermHierarchy
		self.Bind(&iTaxonomyTermHierarchy)
		_TaxonomyTermHierarchy, _error := model.SetTaxonomyTermHierarchyViaTid(Tid_, &iTaxonomyTermHierarchy)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermHierarchy)
	}
	herr.Message = "Can't get to the SetTaxonomyTermHierarchyViaTid's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetTaxonomyTermHierarchyViaParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Parent_ := self.Args("parent").MustInt()
	if helper.IsHas(Parent_) {
		var iTaxonomyTermHierarchy model.TaxonomyTermHierarchy
		self.Bind(&iTaxonomyTermHierarchy)
		_TaxonomyTermHierarchy, _error := model.SetTaxonomyTermHierarchyViaParent(Parent_, &iTaxonomyTermHierarchy)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TaxonomyTermHierarchy)
	}
	herr.Message = "Can't get to the SetTaxonomyTermHierarchyViaParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostAddTaxonomyTermHierarchyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Tid_ := self.Args("tid").MustInt64()
	Parent_ := self.Args("parent").MustInt()

	if helper.IsHas(Tid_) {
		_error := model.AddTaxonomyTermHierarchy(Tid_,Parent_)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		herr.Message = "StatusOK"
		herr.Status = macross.StatusOK
		return self.JSON(herr)
	}
	herr.Message = "Can't get to the AddTaxonomyTermHierarchy's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostTaxonomyTermHierarchyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iTaxonomyTermHierarchy model.TaxonomyTermHierarchy
	self.Bind(&iTaxonomyTermHierarchy)
	_int64, _error := model.PostTaxonomyTermHierarchy(&iTaxonomyTermHierarchy)
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

func PutTaxonomyTermHierarchyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iTaxonomyTermHierarchy model.TaxonomyTermHierarchy
	self.Bind(&iTaxonomyTermHierarchy)
	_int64, _error := model.PutTaxonomyTermHierarchy(&iTaxonomyTermHierarchy)
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

func PutTaxonomyTermHierarchyViaTidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Tid_ := self.Args("tid").MustInt64()
	var iTaxonomyTermHierarchy model.TaxonomyTermHierarchy
	self.Bind(&iTaxonomyTermHierarchy)
	_int64, _error := model.PutTaxonomyTermHierarchyViaTid(Tid_, &iTaxonomyTermHierarchy)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutTaxonomyTermHierarchyViaParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Parent_ := self.Args("parent").MustInt()
	var iTaxonomyTermHierarchy model.TaxonomyTermHierarchy
	self.Bind(&iTaxonomyTermHierarchy)
	_int64, _error := model.PutTaxonomyTermHierarchyViaParent(Parent_, &iTaxonomyTermHierarchy)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateTaxonomyTermHierarchyViaTidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Tid_ := self.Args("tid").MustInt64()
	var iTaxonomyTermHierarchy model.TaxonomyTermHierarchy
	self.Bind(&iTaxonomyTermHierarchy)
	var iMap = helper.StructToMap(iTaxonomyTermHierarchy)
	_error := model.UpdateTaxonomyTermHierarchyViaTid(Tid_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateTaxonomyTermHierarchyViaParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Parent_ := self.Args("parent").MustInt()
	var iTaxonomyTermHierarchy model.TaxonomyTermHierarchy
	self.Bind(&iTaxonomyTermHierarchy)
	var iMap = helper.StructToMap(iTaxonomyTermHierarchy)
	_error := model.UpdateTaxonomyTermHierarchyViaParent(Parent_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteTaxonomyTermHierarchyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Tid_ := self.Args("tid").MustInt64()
	_int64, _error := model.DeleteTaxonomyTermHierarchy(Tid_)
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

func DeleteTaxonomyTermHierarchyViaTidHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Tid_ := self.Args("tid").MustInt64()
	_error := model.DeleteTaxonomyTermHierarchyViaTid(Tid_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteTaxonomyTermHierarchyViaParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Parent_ := self.Args("parent").MustInt()
	_error := model.DeleteTaxonomyTermHierarchyViaParent(Parent_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}
